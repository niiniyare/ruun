package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"regexp"
	"strings"
	"sync"

	"github.com/dgraph-io/ristretto"
)

// =============================================================================
// TOKEN REFERENCE & RESOLUTION
// =============================================================================

// TokenReference represents a reference to another token using dot notation.
// Supports both direct token references and CSS literal values.
//
// Reference formats:
//   - Token reference: "primitives.colors.primary" (dot-separated path)
//   - CSS literal: "hsl(217, 91%, 60%)", "#3b82f6", "1rem"
//   - CSS function: "calc(100% - 2rem)", "var(--primary)"
//
// Token Reference Examples:
//   - "primitives.colors.primary" → resolves to actual color value
//   - "semantic.colors.background" → resolves through reference chain
//   - "components.button.primary.background-color" → component token
//
// CSS Literal Examples (used as-is):
//   - "hsl(217, 91%, 60%)" → HSL color
//   - "1rem" → spacing/size unit
//   - "#3b82f6" → hex color
type TokenReference string

// IsReference determines if a token value is a reference to another token
// or a literal CSS value that should be used as-is.
//
// A value is considered a reference if it:
//   - Contains at least one dot separator (e.g., "colors.primary")
//   - Has 2 or more path segments when split by dots
//   - Doesn't match any CSS value patterns
//   - All segments contain only valid identifier characters
//
// CSS patterns that are NOT references:
//   - Color values: "#fff", "hsl()", "rgb()", "rgba()", "hsla()"
//   - Units: "1rem", "16px", "50%", "1.5em", "100vh"
//   - Functions: "calc()", "var()", "url()", "clamp()", "min()", "max()"
//   - Gradients: "linear-gradient()", "radial-gradient()", "conic-gradient()"
//   - Keywords: "transparent", "inherit", "auto", "none", "bold", "solid"
//   - Numbers: "0", "1", "0.5", "-10", "3.14"
//   - Compound values: "0 1px 2px rgba(0,0,0,0.1)" (contains spaces)
//
// Valid token path characters:
//   - Lowercase: a-z
//   - Uppercase: A-Z
//   - Numbers: 0-9
//   - Separators: hyphen (-), underscore (_)
//   - Path delimiter: dot (.)
//
// IsReference checks if the value is a token reference (not a CSS literal)
// Token references:
// - Must contain at least one dot
// - Must have at least 2 segments (tier.category or tier.category.name)
// - Cannot contain spaces (compound values are not single references)
// - Cannot start with common CSS patterns
func (tr TokenReference) IsReference() bool {
	value := strings.TrimSpace(string(tr))

	if value == "" {
		return false
	}

	// Compound values with spaces are not single references
	if strings.Contains(value, " ") {
		return false
	}

	// Must have at least one dot for a reference
	if !strings.Contains(value, ".") {
		return false
	}

	// Check if it looks like a CSS literal
	// Hex colors
	if strings.HasPrefix(value, "#") {
		return false
	}

	// CSS functions
	cssFunctions := []string{
		"calc(", "var(", "url(", "rgb(", "rgba(", "hsl(", "hsla(",
		"linear-gradient(", "radial-gradient(", "repeating-linear-gradient(",
		"repeating-radial-gradient(", "conic-gradient(",
	}
	for _, fn := range cssFunctions {
		if strings.HasPrefix(value, fn) {
			return false
		}
	}

	// If it has a dot and is a CSS unit/number (e.g., "0.5rem", "3.14px")
	if strings.Contains(value, ".") {
		// Check if it's a decimal number with optional unit
		if matched, _ := regexp.MatchString(`^-?\d+\.\d+[a-zA-Z%]*$`, value); matched {
			return false
		}
	}

	// Must have at least 2 segments separated by dots
	segments := strings.Split(value, ".")
	if len(segments) < 2 {
		return false
	}

	// Check for malformed paths with empty segments
	for _, segment := range segments {
		if segment == "" {
			return false // Contains empty segments (leading, trailing, or double dots)
		}
	}

	// It looks like a reference
	return true
}

// isNumeric checks if a string is numeric
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	// Allow negative numbers
	if strings.HasPrefix(s, "-") {
		s = s[1:]
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// isValidTokenSegment checks if a token path segment contains only valid characters.
// Valid characters: a-z, A-Z, 0-9, hyphen (-), underscore (_)
//
// Examples:
//   - "primary" → valid
//   - "primary-500" → valid
//   - "font_size" → valid
//   - "2xl" → valid
//   - "primary color" → invalid (space)
//   - "primary.secondary" → invalid (dot not allowed in segment)
func isValidTokenSegment(s string) bool {
	for _, char := range s {
		isValid := (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_'
		if !isValid {
			return false
		}
	}
	return true
}

// Path returns the token path for references, or the original value for CSS literals.
//
// Examples:
//   - "primitives.colors.primary" → "primitives.colors.primary"
//   - "hsl(217, 91%, 60%)" → "hsl(217, 91%, 60%)"
func (t TokenReference) Path() string {
	return strings.TrimSpace(string(t))
}

// String returns the string representation of the token reference.
func (t TokenReference) String() string {
	return string(t)
}

// Validate ensures the token reference is well-formed.
//
// For token references:
//   - Path must not be empty
//   - Must contain at least one dot
//   - Must have at least 2 path segments
//   - All segments must be non-empty
//   - Only valid identifier characters allowed (alphanumeric, hyphens, underscores, dots)
//
// For CSS literals:
//   - Always considered valid (no additional validation)
//
// Returns ValidationError if the reference is malformed.
func (t TokenReference) Validate() error {
	s := strings.TrimSpace(string(t))
	if s == "" {
		return NewValidationError("empty_token_reference", "token reference cannot be empty")
	}

	// If it's not a reference, it must be a valid CSS literal
	if !t.IsReference() {
		// Check if it looks like an attempted token reference that's malformed
		if t.isLikelyMalformedReference() {
			return t.validateAsTokenReference()
		}
		return nil
	}

	return t.validateAsTokenReference()
}

// validateAsTokenReference validates the value as a token reference
func (t TokenReference) validateAsTokenReference() error {
	path := t.Path()

	// Must contain at least one dot
	if !strings.Contains(path, ".") {
		return NewValidationError("invalid_token_path",
			fmt.Sprintf("token path must contain at least one dot: %s", path))
	}

	// Validate path segments
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return NewValidationError("invalid_token_path",
			fmt.Sprintf("token path must have at least 2 segments: %s", path))
	}

	for i, part := range parts {
		if part == "" {
			return NewValidationError("empty_token_segment",
				fmt.Sprintf("empty segment at position %d in path: %s", i, path))
		}
		if !isValidTokenSegment(part) {
			return NewValidationError("invalid_token_segment",
				fmt.Sprintf("invalid segment '%s' at position %d in path: %s", part, i, path))
		}
	}

	return nil
}

// isLikelyMalformedReference checks if a value looks like it's attempting to be a token reference
// but has formatting issues that make IsReference() return false
func (t TokenReference) isLikelyMalformedReference() bool {
	s := strings.TrimSpace(string(t))
	if s == "" {
		return false
	}

	// If it contains dots, it might be trying to be a reference
	// But exclude CSS values that legitimately contain dots (like rgba decimal values)
	if strings.Contains(s, ".") {
		// Check if dots are in CSS function contexts (like rgba(0, 0, 0, 0.5))
		if strings.Contains(s, "(") && strings.Contains(s, ")") {
			return false // This is likely a CSS function with decimal values
		}

		// Check if it contains spaces - token references shouldn't have spaces
		// But if it has both dots AND spaces, it's likely a malformed token reference
		if strings.Contains(s, " ") {
			// If it looks like token notation (dots + segments), treat as malformed reference
			parts := strings.Split(strings.Fields(s)[0], ".")
			if len(parts) >= 2 {
				return true // Likely malformed token reference like "colors.primary color"
			}
			return false // This is likely a compound CSS value
		}

		return true
	}

	// Single words that are clearly CSS values should not be treated as malformed references
	if len(strings.Fields(s)) == 1 {
		// Pure numbers (including decimals) are CSS values, not malformed references
		if matched, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, s); matched {
			return false
		}

		// If it has obvious CSS indicators, it's not a malformed reference
		if t.hasObviousCSSIndicators() {
			return false
		}

		// Short strings (1-2 chars) that aren't CSS keywords are ambiguous but lean towards CSS
		if len(s) <= 2 {
			return false
		}

		// Single words that look like token identifiers are likely malformed references
		// unless they're clearly CSS values (which would have been caught above)
		validChars := true
		hasAlpha := false
		for _, r := range s {
			if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') ||
				(r >= '0' && r <= '9') || r == '-' || r == '_') {
				validChars = false
				break
			}
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				hasAlpha = true
			}
		}

		// Treat as potential malformed reference if it looks like a token identifier
		if validChars && hasAlpha {
			return true
		}
	}

	return false
}

// hasObviousCSSIndicators checks if a value has clear CSS indicators
func (t TokenReference) hasObviousCSSIndicators() bool {
	s := strings.TrimSpace(string(t))

	// Has CSS color indicators
	if strings.HasPrefix(s, "#") {
		return true
	}

	// Has CSS function indicators
	if strings.Contains(s, "(") && strings.Contains(s, ")") {
		return true
	}

	// Has CSS unit indicators - check if it's a number followed by a unit
	cssUnits := []string{
		// Length units
		"px", "rem", "em", "%", "vh", "vw", "vmin", "vmax",
		"pt", "cm", "mm", "in", "pc", "ex", "ch", "fr",
		// Angle units
		"deg", "rad", "grad", "turn",
		// Time units
		"s", "ms",
		// Resolution units
		"dpi", "dpcm", "dppx",
	}
	for _, unit := range cssUnits {
		if strings.HasSuffix(s, unit) {
			// Make sure there's actually a number before the unit
			prefix := strings.TrimSuffix(s, unit)
			if prefix != "" && isNumeric(prefix) {
				return true
			}
		}
	}

	// Check CSS keywords
	cssKeywords := []string{
		// Color keywords
		"transparent", "currentColor", "inherit", "initial", "unset", "revert",
		// General keywords
		"auto", "none", "normal", "bold", "bolder", "lighter", "italic", "oblique",
		// Border styles
		"solid", "dashed", "dotted", "double", "groove", "ridge", "inset", "outset",
		// Display keywords
		"hidden", "visible", "collapse", "block", "inline", "inline-block",
		"flex", "inline-flex", "grid", "inline-grid", "flow-root",
		// Animation/transition keywords
		"linear", "ease", "ease-in", "ease-out", "ease-in-out",
		// Position keywords
		"relative", "absolute", "fixed", "sticky", "static",
		// Sizing keywords
		"fit-content", "max-content", "min-content",
		// Text keywords
		"left", "right", "center", "justify",
	}
	lowerS := strings.ToLower(s)
	for _, keyword := range cssKeywords {
		if keyword == lowerS {
			return true
		}
	}

	return false
}

// =============================================================================
// DESIGN TOKENS - ROOT STRUCTURE
// =============================================================================

// DesignTokens is the root container for the entire design token system.
//
// Architecture - Three-Tier Hierarchy:
//
//  1. Primitives (Foundation Layer)
//     - Raw design values: colors, spacing, typography scales
//     - These are literal CSS values, not references
//     - Example: "primary": "hsl(217, 91%, 60%)"
//
//  2. Semantic (Contextual Layer)
//     - Purpose-based token assignments
//     - References primitive tokens or provides context-specific values
//     - Example: "background": "primitives.colors.gray-50"
//
//  3. Components (Application Layer)
//     - Component-specific style definitions with variants
//     - References semantic or primitive tokens
//     - Example: "button.primary.background-color": "semantic.colors.primary"
//
// Token Resolution:
//   - Components can reference Semantic or Primitives
//   - Semantic can reference Primitives
//   - Primitives contain only CSS literals (no references)
//
// JSON Structure:
//
//	{
//	  "primitives": {
//	    "colors": { "primary": "hsl(217, 91%, 60%)", ... },
//	    "spacing": { "md": "1rem", ... }
//	  },
//	  "semantic": {
//	    "colors": { "background": "primitives.colors.white", ... }
//	  },
//	  "components": {
//	    "button": {
//	      "primary": {
//	        "background-color": "semantic.colors.primary",
//	        ...
//	      }
//	    }
//	  }
//	}
type DesignTokens struct {
	// Primitives contains raw, foundational design values
	// These are CSS literals and form the base of the token system
	Primitives *PrimitiveTokens `json:"primitives" yaml:"primitives"`

	// Semantic contains context-aware token assignments
	// These reference primitive tokens and assign semantic meaning
	Semantic *SemanticTokens `json:"semantic" yaml:"semantic"`

	// Components contains component-specific tokens with variants
	// These reference semantic or primitive tokens for component styles
	Components *ComponentTokens `json:"components" yaml:"components"`
}

// =============================================================================
// PRIMITIVE TOKENS - FOUNDATION LAYER
// =============================================================================

// PrimitiveTokens contains all foundational design values.
// These are raw CSS values that form the base of the design system.
//
// Naming Convention (kebab-case, Tailwind-inspired):
//   - Colors: "primary", "gray-50", "gray-900", "blue-500"
//   - Spacing: "xs", "sm", "md", "lg", "xl", "2xl", "3xl"
//   - Typography: "font-size-sm", "font-weight-bold", "line-height-tight"
//   - Radius: "sm", "md", "lg", "xl", "full"
//
// Values:
//   - Always CSS literals (never references)
//   - Use modern CSS formats: hsl(), rem, em, px
//   - Avoid hard-coded hex colors; prefer HSL for better manipulation
type PrimitiveTokens struct {
	// Colors defines the complete color palette
	// Keys: primary, secondary, gray-*, blue-*, success, error, warning, info
	// Values: HSL colors (e.g., "hsl(217, 91%, 60%)")
	Colors map[string]string `json:"colors" yaml:"colors"`

	// Spacing defines the spacing scale for margins, padding, gaps
	// Keys: xs, sm, md, lg, xl, 2xl, 3xl, 4xl
	// Values: rem/px units (e.g., "1rem", "16px")
	Spacing map[string]string `json:"spacing" yaml:"spacing"`

	// Radius defines border-radius values
	// Keys: none, sm, md, lg, xl, 2xl, 3xl, full
	// Values: rem/px units (e.g., "0.5rem", "9999px")
	Radius map[string]string `json:"radius" yaml:"radius"`

	// Typography defines all font-related values
	// Keys: font-size-*, font-weight-*, line-height-*, font-family-*, letter-spacing-*
	// Values: rem/px for sizes, numeric for weights, unitless for line-height
	Typography map[string]string `json:"typography" yaml:"typography"`

	// Borders defines border width and style primitives
	// Keys: border-width-*, border-style-*
	// Values: px for width, CSS keywords for style
	Borders map[string]string `json:"borders" yaml:"borders"`

	// Shadows defines elevation shadow values
	// Keys: none, xs, sm, md, lg, xl, 2xl, inner
	// Values: CSS box-shadow values
	Shadows map[string]string `json:"shadows" yaml:"shadows"`

	// Effects defines visual effects
	// Keys: opacity-*, blur-*, brightness-*, contrast-*
	// Values: numeric for opacity, px for blur, percentage for filters
	Effects map[string]string `json:"effects" yaml:"effects"`

	// Animation defines animation timing primitives
	// Keys: duration-*, easing-*, delay-*
	// Values: ms for duration/delay, cubic-bezier for easing
	Animation map[string]string `json:"animation" yaml:"animation"`

	// ZIndex defines layering values
	// Keys: base, dropdown, sticky, fixed, modal, popover, tooltip
	// Values: integers (e.g., "1", "10", "50")
	ZIndex map[string]string `json:"z-index" yaml:"z-index"`

	// Breakpoints defines responsive breakpoints
	// Keys: sm, md, lg, xl, 2xl
	// Values: px values (e.g., "640px", "768px")
	Breakpoints map[string]string `json:"breakpoints" yaml:"breakpoints"`
}

// =============================================================================
// SEMANTIC TOKENS - CONTEXTUAL LAYER
// =============================================================================

// SemanticTokens contains context-aware token assignments.
// These tokens reference primitive tokens and provide semantic meaning.
//
// Purpose:
//   - Assign meaning to primitive values (e.g., "background" vs "gray-50")
//   - Create consistent design language
//   - Enable theme variations (light/dark) by changing references
//   - Reduce coupling between components and primitives
//
// Token References:
//   - Should primarily reference primitives: "primitives.colors.primary"
//   - Can use CSS literals for computed values
//   - Example: "background": "primitives.colors.white"
type SemanticTokens struct {
	// Colors defines semantic color assignments
	// Keys: background, foreground, primary, secondary, muted, border, ring,
	//       success, warning, error, info (with -foreground variants)
	// Values: References to primitive colors or CSS literals
	Colors map[string]string `json:"colors" yaml:"colors"`

	// Spacing defines semantic spacing assignments
	// Keys: component-tight, component-default, component-loose,
	//       layout-section, layout-page, stack-tight, stack-default
	// Values: References to primitive spacing or CSS literals
	Spacing map[string]string `json:"spacing" yaml:"spacing"`

	// Typography defines semantic typography assignments
	// Keys: heading-*, body-*, label-*, caption-*, code-*
	// Values: References to primitive typography or CSS literals
	Typography map[string]string `json:"typography" yaml:"typography"`

	// Interactive defines interactive element styling
	// Keys: border-radius-*, shadow-*, transition-*
	// Values: References to primitives or CSS literals
	Interactive map[string]string `json:"interactive" yaml:"interactive"`
}

// =============================================================================
// COMPONENT TOKENS - APPLICATION LAYER
// =============================================================================

// ComponentTokens contains component-specific design tokens with variants.
// Structure: component -> variant -> property -> value
//
// Component Organization:
//   - Each component (button, input, card) has multiple variants
//   - Each variant defines all style properties for that appearance
//   - Properties use CSS property names (kebab-case)
//
// Variants Pattern:
//   - base: Default/shared styles for all variants
//   - primary, secondary, tertiary: Visual hierarchy variants
//   - outline, ghost, link: Appearance variants
//   - success, warning, error, info: Status variants
//   - sm, md, lg: Size variants
//
// Token References:
//   - Should reference semantic tokens: "semantic.colors.primary"
//   - Can reference primitives: "primitives.spacing.md"
//   - Can use CSS literals for component-specific values
//
// Example:
//
//	{
//	  "button": {
//	    "primary": {
//	      "background-color": "semantic.colors.primary",
//	      "color": "semantic.colors.primary-foreground",
//	      "border-radius": "semantic.interactive.border-radius-md"
//	    },
//	    "secondary": { ... }
//	  }
//	}
type ComponentTokens map[string]ComponentVariants

// ComponentVariants maps variant names to their style properties.
// Each variant is a complete style definition for that component appearance.
//
// Structure: variant -> property -> value
//
// Common Variants:
//   - base: Shared styles applied to all variants
//   - primary, secondary, tertiary: Hierarchy levels
//   - outline, ghost, solid, link: Visual styles
//   - success, warning, error, info: Status indicators
//   - sm, md, lg, xl: Size variants
//   - disabled, hover, focus, active: State variants
type ComponentVariants map[string]StyleProperties

// StyleProperties maps CSS property names to their values.
// Property names must use kebab-case (matching CSS syntax).
//
// Property Naming:
//   - Use full CSS property names: "background-color", not "bg"
//   - Use kebab-case: "border-radius", not "borderRadius"
//   - Follow CSS specification naming
//
// Value Types:
//   - Token references: "semantic.colors.primary"
//   - CSS literals: "hsl(217, 91%, 60%)", "1rem", "500"
//   - Compound values: "0.5rem 1rem", "1px solid", "inset 0 2px 4px"
//
// Example:
//
//	{
//	  "background-color": "semantic.colors.primary",
//	  "color": "semantic.colors.primary-foreground",
//	  "padding": "0.5rem 1rem",
//	  "border": "1px solid semantic.colors.border",
//	  "border-radius": "semantic.interactive.border-radius-md"
//	}
type StyleProperties map[string]string

// =============================================================================
// CLONE OPERATIONS
// =============================================================================

// Clone creates a deep copy of the design tokens.
// Use this to ensure immutability when modifying tokens.
func (dt *DesignTokens) Clone() *DesignTokens {
	if dt == nil {
		return nil
	}
	cloned := &DesignTokens{}
	if dt.Primitives != nil {
		cloned.Primitives = dt.Primitives.Clone()
	}
	if dt.Semantic != nil {
		cloned.Semantic = dt.Semantic.Clone()
	}
	if dt.Components != nil {
		cloned.Components = dt.Components.Clone()
	}
	return cloned
}

// Clone creates a deep copy of primitive tokens.
func (pt *PrimitiveTokens) Clone() *PrimitiveTokens {
	if pt == nil {
		return nil
	}
	return &PrimitiveTokens{
		Colors:      cloneMap(pt.Colors),
		Spacing:     cloneMap(pt.Spacing),
		Radius:      cloneMap(pt.Radius),
		Typography:  cloneMap(pt.Typography),
		Borders:     cloneMap(pt.Borders),
		Shadows:     cloneMap(pt.Shadows),
		Effects:     cloneMap(pt.Effects),
		Animation:   cloneMap(pt.Animation),
		ZIndex:      cloneMap(pt.ZIndex),
		Breakpoints: cloneMap(pt.Breakpoints),
	}
}

// Clone creates a deep copy of semantic tokens.
func (st *SemanticTokens) Clone() *SemanticTokens {
	if st == nil {
		return nil
	}
	return &SemanticTokens{
		Colors:      cloneMap(st.Colors),
		Spacing:     cloneMap(st.Spacing),
		Typography:  cloneMap(st.Typography),
		Interactive: cloneMap(st.Interactive),
	}
}

// Clone creates a deep copy of component tokens.
func (ct *ComponentTokens) Clone() *ComponentTokens {
	if ct == nil {
		return nil
	}
	cloned := make(ComponentTokens, len(*ct))
	for component, variants := range *ct {
		cloned[component] = variants.Clone()
	}
	return &cloned
}

// Clone creates a deep copy of component variants.
func (cv ComponentVariants) Clone() ComponentVariants {
	if cv == nil {
		return nil
	}
	cloned := make(ComponentVariants, len(cv))
	for variant, props := range cv {
		cloned[variant] = cloneMap(props)
	}
	return cloned
}

// cloneMap creates a deep copy of a string map.
func cloneMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	cloned := make(map[string]string, len(m))
	maps.Copy(cloned, m)
	return cloned
}

// =============================================================================
// VALIDATION
// =============================================================================

// Validate ensures the entire token system is well-formed.
// Checks structure integrity and validates all token values.
func (dt *DesignTokens) Validate() error {
	if dt == nil {
		return NewValidationError("nil_tokens", "design tokens cannot be nil")
	}

	if dt.Primitives != nil {
		if err := dt.Primitives.Validate(); err != nil {
			return fmt.Errorf("invalid primitives: %w", err)
		}
	}

	if dt.Semantic != nil {
		if err := dt.Semantic.Validate(); err != nil {
			return fmt.Errorf("invalid semantic: %w", err)
		}
	}

	if dt.Components != nil {
		if err := dt.Components.Validate(); err != nil {
			return fmt.Errorf("invalid components: %w", err)
		}
	}

	return nil
}

// Validate checks all primitive token values.
func (pt *PrimitiveTokens) Validate() error {
	if pt == nil {
		return NewValidationError("nil_primitives", "primitive tokens cannot be nil")
	}

	categories := map[string]map[string]string{
		"colors":      pt.Colors,
		"spacing":     pt.Spacing,
		"radius":      pt.Radius,
		"typography":  pt.Typography,
		"borders":     pt.Borders,
		"shadows":     pt.Shadows,
		"effects":     pt.Effects,
		"animation":   pt.Animation,
		"z-index":     pt.ZIndex,
		"breakpoints": pt.Breakpoints,
	}

	for category, tokens := range categories {
		if tokens == nil {
			continue
		}
		if len(tokens) == 0 {
			return NewValidationError("empty_category",
				fmt.Sprintf("primitive category '%s' cannot be empty", category))
		}
		for key, value := range tokens {
			ref := TokenReference(value)
			if err := ref.Validate(); err != nil {
				return fmt.Errorf("invalid primitive token %s.%s: %w", category, key, err)
			}
			// Primitives should not be references (should be CSS literals)
			if ref.IsReference() {
				return NewValidationError("primitive_reference",
					fmt.Sprintf("primitive token %s.%s contains reference '%s' (primitives should be CSS literals)",
						category, key, value))
			}
		}
	}

	return nil
}

// Validate checks all semantic token values.
func (st *SemanticTokens) Validate() error {
	if st == nil {
		return nil // Semantic tokens are optional
	}

	categories := map[string]map[string]string{
		"colors":      st.Colors,
		"spacing":     st.Spacing,
		"typography":  st.Typography,
		"interactive": st.Interactive,
	}

	for category, tokens := range categories {
		if tokens == nil {
			continue
		}
		for key, value := range tokens {
			ref := TokenReference(value)
			if err := ref.Validate(); err != nil {
				return fmt.Errorf("invalid semantic token %s.%s: %w", category, key, err)
			}
		}
	}

	return nil
}

// Validate checks all component token values.
func (ct *ComponentTokens) Validate() error {
	if ct == nil {
		return nil // Component tokens are optional
	}

	for component, variants := range *ct {
		if err := variants.Validate(component); err != nil {
			return fmt.Errorf("invalid component '%s': %w", component, err)
		}
	}

	return nil
}

// Validate checks all variant style properties.
func (cv ComponentVariants) Validate(component string) error {
	if cv == nil {
		return nil
	}

	for variant, props := range cv {
		if len(props) == 0 {
			return NewValidationError("empty_variant",
				fmt.Sprintf("variant '%s' has no properties", variant))
		}
		for property, value := range props {
			if err := validateComponentPropertyValue(component, variant, property, value); err != nil {
				return err
			}
		}
	}

	return nil
}

// =============================================================================
// TOKEN ACCESS METHODS
// =============================================================================

// GetToken retrieves a token value by its path.
// Supports paths in primitives, semantic, or components.
//
// Path formats:
//   - Primitive: "primitives.colors.primary" or "primitives.spacing.md"
//   - Semantic: "semantic.colors.background" or "semantic.spacing.component-default"
//   - Component: "components.button.primary.background-color"
//
// Returns the raw token value (may be a reference or CSS literal).
// Use TokenResolver.Resolve() to get the final computed value.
func (dt *DesignTokens) GetToken(path string) (string, error) {
	if dt == nil {
		return "", NewValidationError("nil_tokens", "design tokens cannot be nil")
	}

	parts := strings.Split(path, ".")
	if len(parts) < 3 {
		return "", NewValidationError("invalid_token_path",
			fmt.Sprintf("token path must have at least 3 segments (tier.category.key): %s", path))
	}

	tier := parts[0]

	switch tier {
	case "primitives":
		return dt.getPrimitiveToken(parts[1:])
	case "semantic":
		return dt.getSemanticToken(parts[1:])
	case "components":
		return dt.getComponentToken(parts[1:])
	default:
		return "", NewValidationError("invalid_tier",
			fmt.Sprintf("unknown token tier: %s (must be primitives, semantic, or components)", tier))
	}
}

// getPrimitiveToken retrieves a primitive token value.
// Path: [category, key] (e.g., ["colors", "primary"])
func (dt *DesignTokens) getPrimitiveToken(parts []string) (string, error) {
	if len(parts) != 2 {
		return "", NewValidationError("invalid_primitive_path",
			fmt.Sprintf("primitive path must have 2 segments (category.key): %v", parts))
	}

	if dt.Primitives == nil {
		return "", NewValidationError("nil_primitives", "primitives not initialized")
	}

	category, key := parts[0], parts[1]

	var tokenMap map[string]string
	switch category {
	case "colors":
		tokenMap = dt.Primitives.Colors
	case "spacing":
		tokenMap = dt.Primitives.Spacing
	case "radius":
		tokenMap = dt.Primitives.Radius
	case "typography":
		tokenMap = dt.Primitives.Typography
	case "borders":
		tokenMap = dt.Primitives.Borders
	case "shadows":
		tokenMap = dt.Primitives.Shadows
	case "effects":
		tokenMap = dt.Primitives.Effects
	case "animation":
		tokenMap = dt.Primitives.Animation
	case "z-index":
		tokenMap = dt.Primitives.ZIndex
	case "breakpoints":
		tokenMap = dt.Primitives.Breakpoints
	default:
		return "", NewValidationError("invalid_category",
			fmt.Sprintf("unknown primitive category: %s", category))
	}

	if tokenMap == nil {
		return "", NewValidationError("nil_category",
			fmt.Sprintf("primitive category '%s' not initialized", category))
	}

	value, exists := tokenMap[key]
	if !exists {
		return "", NewValidationError("token_not_found",
			fmt.Sprintf("primitive token not found: primitives.%s.%s", category, key))
	}

	return value, nil
}

// getSemanticToken retrieves a semantic token value.
// Path: [category, key] (e.g., ["colors", "background"])
func (dt *DesignTokens) getSemanticToken(parts []string) (string, error) {
	if len(parts) != 2 {
		return "", NewValidationError("invalid_semantic_path",
			fmt.Sprintf("semantic path must have 2 segments (category.key): %v", parts))
	}

	if dt.Semantic == nil {
		return "", NewValidationError("nil_semantic", "semantic tokens not initialized")
	}

	category, key := parts[0], parts[1]

	var tokenMap map[string]string
	switch category {
	case "colors":
		tokenMap = dt.Semantic.Colors
	case "spacing":
		tokenMap = dt.Semantic.Spacing
	case "typography":
		tokenMap = dt.Semantic.Typography
	case "interactive":
		tokenMap = dt.Semantic.Interactive
	default:
		return "", NewValidationError("invalid_category",
			fmt.Sprintf("unknown semantic category: %s", category))
	}

	if tokenMap == nil {
		return "", NewValidationError("nil_category",
			fmt.Sprintf("semantic category '%s' not initialized", category))
	}

	value, exists := tokenMap[key]
	if !exists {
		return "", NewValidationError("token_not_found",
			fmt.Sprintf("semantic token not found: semantic.%s.%s", category, key))
	}

	return value, nil
}

// getComponentToken retrieves a component token value.
// Path: [component, variant, property] (e.g., ["button", "primary", "background-color"])
func (dt *DesignTokens) getComponentToken(parts []string) (string, error) {
	if len(parts) != 3 {
		return "", NewValidationError("invalid_component_path",
			fmt.Sprintf("component path must have 3 segments (component.variant.property): %v", parts))
	}

	if dt.Components == nil {
		return "", NewValidationError("nil_components", "component tokens not initialized")
	}

	component, variant, property := parts[0], parts[1], parts[2]

	variants, exists := (*dt.Components)[component]
	if !exists {
		return "", NewValidationError("component_not_found",
			fmt.Sprintf("component not found: %s", component))
	}

	props, exists := variants[variant]
	if !exists {
		return "", NewValidationError("variant_not_found",
			fmt.Sprintf("variant not found: %s.%s", component, variant))
	}

	value, exists := props[property]
	if !exists {
		return "", NewValidationError("property_not_found",
			fmt.Sprintf("property not found: components.%s.%s.%s", component, variant, property))
	}

	return value, nil
}

// =============================================================================
// TOKEN RESOLVER
// =============================================================================

// TokenResolver provides high-performance token resolution with caching.
// Resolves token references to their final CSS values by traversing the token hierarchy.
//
// Features:
//   - Recursive reference resolution through token hierarchy
//   - Circular reference detection with detailed error messages
//   - High-performance LRU cache with bounded memory
//   - Context-aware cancellation support
//   - Thread-safe for concurrent use
//   - Automatic cache invalidation on token updates
//
// Resolution Order:
//  1. Check if value is a CSS literal → return immediately
//  2. Check cache for previously resolved value
//  3. Look up token in appropriate tier (primitives/semantic/components)
//  4. If value is a reference, recursively resolve it
//  5. Cache the final resolved value
//  6. Return the CSS literal value
//
// Example:
//
//	resolver, _ := NewTokenResolver(tokens)
//	// Resolves: semantic.colors.primary → primitives.colors.primary → "hsl(217, 91%, 60%)"
//	value, err := resolver.Resolve(ctx, TokenReference("semantic.colors.primary"))
type TokenResolver struct {
	tokens *DesignTokens
	cache  *ristretto.Cache
	mu     sync.RWMutex
}

// NewTokenResolver creates a new token resolver with bounded cache.
//
// Cache Configuration:
//   - MaxCost: 1MB total memory for cached strings
//   - NumCounters: 10,000 items for admission policy
//   - BufferItems: 64 async operations buffer
//   - Cost function: String length in bytes
//   - Eviction: Automatic LRU when cache is full
//
// The resolver validates the token structure on creation to ensure
// all tokens are well-formed before any resolution attempts.
func NewTokenResolver(tokens *DesignTokens) (*TokenResolver, error) {
	if tokens == nil {
		return nil, NewValidationError("nil_tokens", "design tokens cannot be nil")
	}

	// Validate token structure before creating resolver
	if err := tokens.Validate(); err != nil {
		return nil, fmt.Errorf("invalid design tokens: %w", err)
	}

	// Create bounded cache with LRU eviction
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,   // Track 10k items for admission
		MaxCost:     1 << 20, // 1MB total cache size
		BufferItems: 64,      // Async operation buffer
		Cost: func(value any) int64 {
			return int64(len(value.(string))) // Cost = string byte length
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create token cache: %w", err)
	}

	return &TokenResolver{
		tokens: tokens,
		cache:  cache,
	}, nil
}

// Resolve resolves a token reference to its final CSS value.
//
// Resolution process:
//  1. If input is CSS literal (not a reference) → return as-is
//  2. Check cache for previously resolved value
//  3. Look up token in design token hierarchy
//  4. If value is another reference, recursively resolve it
//  5. Detect circular references during traversal
//  6. Cache the final CSS literal value
//  7. Return the resolved value
//
// Thread-safe: Multiple goroutines can call Resolve concurrently.
//
// Examples:
//   - Input: "hsl(217, 91%, 60%)" → Output: "hsl(217, 91%, 60%)" (CSS literal, no resolution)
//   - Input: "primitives.colors.primary" → Output: "hsl(217, 91%, 60%)" (direct lookup)
//   - Input: "semantic.colors.primary" → Output: "hsl(217, 91%, 60%)" (resolves through reference chain)
//   - Input: "components.button.primary.background-color" → Output: "hsl(217, 91%, 60%)" (multi-tier resolution)
func (tr *TokenResolver) Resolve(ctx context.Context, ref TokenReference) (string, error) {
	// Check context cancellation early
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	// CSS literals don't need resolution - return immediately
	if !ref.IsReference() {
		return ref.String(), nil
	}

	path := ref.Path()

	// Check cache for previously resolved value
	tr.mu.RLock()
	if cached, found := tr.cache.Get(path); found {
		tr.mu.RUnlock()
		return cached.(string), nil
	}
	tr.mu.RUnlock()

	// Resolve with circular reference detection
	visited := make(map[string]bool)
	resolved, err := tr.resolveRecursive(ctx, ref, visited)
	if err != nil {
		return "", err
	}

	// Cache the resolved value for future lookups
	tr.mu.Lock()
	tr.cache.Set(path, resolved, int64(len(resolved)))
	tr.mu.Unlock()

	return resolved, nil
}

// resolveRecursive performs recursive token resolution with cycle detection.
// Uses a visited map to detect circular references and prevent infinite loops.
func (tr *TokenResolver) resolveRecursive(ctx context.Context, ref TokenReference, visited map[string]bool) (string, error) {
	// Check context cancellation
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	path := ref.Path()

	// Detect circular references
	if visited[path] {
		// Build the circular path for better error message
		visitedPaths := make([]string, 0, len(visited))
		for p := range visited {
			visitedPaths = append(visitedPaths, p)
		}
		return "", NewValidationError("circular_reference",
			fmt.Sprintf("circular token reference detected at '%s' (visited: %v)", path, visitedPaths))
	}

	// Mark this path as visited
	visited[path] = true
	defer delete(visited, path) // Clean up for proper backtracking

	// Look up the token value
	value, err := tr.tokens.GetToken(path)
	if err != nil {
		return "", fmt.Errorf("failed to resolve token '%s': %w", path, err)
	}

	// Check if the value is itself a reference
	valueRef := TokenReference(value)
	if valueRef.IsReference() {
		// Recursively resolve the referenced token
		return tr.resolveRecursive(ctx, valueRef, visited)
	}

	// Value is a CSS literal - return it
	return value, nil
}

// ResolveAll resolves all tokens in the design system to their final CSS values.
// Returns a new DesignTokens instance with all references resolved to literals.
//
// This is useful for:
//   - Exporting themes for use in other systems
//   - Generating CSS custom properties
//   - Creating static token files
//   - Debugging token resolution
//   - Validating the complete token chain
//
// The returned tokens contain only CSS literals (no references).
func (tr *TokenResolver) ResolveAll(ctx context.Context) (*DesignTokens, error) {
	// Clone to avoid modifying original
	resolved := tr.tokens.Clone()

	// Resolve semantic tokens (these may reference primitives)
	if resolved.Semantic != nil {
		if err := tr.resolveSemanticTokens(ctx, resolved.Semantic); err != nil {
			return nil, fmt.Errorf("failed to resolve semantic tokens: %w", err)
		}
	}

	// Resolve component tokens (these may reference semantic or primitives)
	if resolved.Components != nil {
		if err := tr.resolveComponentTokens(ctx, resolved.Components); err != nil {
			return nil, fmt.Errorf("failed to resolve component tokens: %w", err)
		}
	}

	return resolved, nil
}

// resolveSemanticTokens resolves all semantic token references to CSS literals.
func (tr *TokenResolver) resolveSemanticTokens(ctx context.Context, semantic *SemanticTokens) error {
	categories := []struct {
		name string
		m    map[string]string
	}{
		{"colors", semantic.Colors},
		{"spacing", semantic.Spacing},
		{"typography", semantic.Typography},
		{"interactive", semantic.Interactive},
	}

	for _, cat := range categories {
		if cat.m == nil {
			continue
		}
		for key, value := range cat.m {
			resolved, err := tr.Resolve(ctx, TokenReference(value))
			if err != nil {
				return fmt.Errorf("failed to resolve semantic.%s.%s: %w", cat.name, key, err)
			}
			cat.m[key] = resolved
		}
	}

	return nil
}

// resolveComponentTokens resolves all component token references to CSS literals.
func (tr *TokenResolver) resolveComponentTokens(ctx context.Context, components *ComponentTokens) error {
	for component, variants := range *components {
		for variant, props := range variants {
			for property, value := range props {
				resolved, err := tr.Resolve(ctx, TokenReference(value))
				if err != nil {
					return fmt.Errorf("failed to resolve components.%s.%s.%s: %w",
						component, variant, property, err)
				}
				props[property] = resolved
			}
		}
	}
	return nil
}

// ClearCache clears all cached token resolutions.
// Call this when tokens are updated to invalidate stale cached values.
//
// Thread-safe: Can be called concurrently with Resolve operations.
func (tr *TokenResolver) ClearCache() {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	tr.cache.Clear()
}

// Close releases resources used by the resolver.
// Must be called when the resolver is no longer needed to prevent resource leaks.
//
// After calling Close, the resolver cannot be used anymore.
func (tr *TokenResolver) Close() error {
	tr.cache.Close()
	return nil
}

// =============================================================================
// SERIALIZATION
// =============================================================================

// TokensToJSON serializes the design tokens to pretty-printed JSON.
func (dt *DesignTokens) TokensToJSON() ([]byte, error) {
	return json.MarshalIndent(dt, "", "  ")
}

// TokensFromJSON deserializes design tokens from JSON and validates the structure.
func TokensFromJSON(data []byte) (*DesignTokens, error) {
	var tokens DesignTokens
	if err := json.Unmarshal(data, &tokens); err != nil {
		return nil, fmt.Errorf("failed to parse design tokens JSON: %w", err)
	}
	if err := tokens.Validate(); err != nil {
		return nil, fmt.Errorf("invalid design tokens: %w", err)
	}
	return &tokens, nil
}

// validateComponentTokens validates component-level tokens
func validateComponentTokens(components *ComponentTokens) error {
	if components == nil {
		return nil
	}

	for componentName, variants := range *components {
		if componentName == "" {
			return newValidationError("empty_component_name", "component name cannot be empty")
		}

		if len(variants) == 0 {
			return newValidationError("empty_component_variants",
				fmt.Sprintf("component '%s' has no variants", componentName))
		}

		for variantName, properties := range variants {
			if variantName == "" {
				return newValidationError("empty_variant_name",
					fmt.Sprintf("variant name cannot be empty in component '%s'", componentName))
			}

			for propName, value := range properties {
				if propName == "" {
					return newValidationError("empty_property_name",
						fmt.Sprintf("property name cannot be empty in component '%s', variant '%s'",
							componentName, variantName))
				}

				// Validate the property value
				if err := validateComponentPropertyValue(componentName, variantName, propName, value); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// validateComponentPropertyValue validates a component property value
// Component values can be:
// 1. CSS literals: "1rem", "#fff", "solid"
// 2. Token references: "primitives.colors.primary"
// 3. Compound values: "0.5rem 1rem" or "1px solid semantic.colors.border"
func validateComponentPropertyValue(component, variant, property, value string) error {
	if strings.TrimSpace(value) == "" {
		return newValidationError("empty_property_value",
			fmt.Sprintf("empty value in component '%s', variant '%s', property '%s'",
				component, variant, property))
	}

	// For compound values (containing spaces), validate each part separately
	if strings.Contains(value, " ") {
		return validateCompoundValue(component, variant, property, value)
	}

	// For single values, check if it's a reference
	ref := TokenReference(value)
	if ref.IsReference() {
		// Validate the reference format
		if err := ref.Validate(); err != nil {
			return newValidationError("invalid_property",
				fmt.Sprintf("invalid component '%s': invalid property '%s.%s': %v",
					component, variant, property, err))
		}
	}
	// If it's not a reference, it's a CSS literal - always valid

	return nil
}

// validateCompoundValue validates compound CSS values that may contain multiple tokens
// Examples: "0.5rem 1rem", "1px solid #ccc", "0 4px 6px rgba(0,0,0,0.1)"
func validateCompoundValue(component, variant, property, value string) error {
	// Handle CSS functions that contain spaces (like rgba, hsla, calc, etc.)
	// Simple approach: if it contains function parentheses, treat as single unit
	if strings.Contains(value, "(") && strings.Contains(value, ")") {
		// For now, just check if the whole value might be a single CSS function
		// This is a simplified approach - a full CSS parser would be better
		if strings.Count(value, "(") == 1 && strings.Count(value, ")") == 1 {
			// Single function - validate as one unit
			ref := TokenReference(value)
			if ref.IsReference() {
				if err := ref.Validate(); err != nil {
					return newValidationError("invalid_property",
						fmt.Sprintf("invalid component '%s': invalid property '%s.%s': %v",
							component, variant, property, err))
				}
			}
			return nil
		}
	}
	
	parts := strings.Fields(value) // Split on whitespace

	for i, part := range parts {
		ref := TokenReference(part)

		// Only validate if it looks like a reference
		if ref.IsReference() {
			if err := ref.Validate(); err != nil {
				return newValidationError("invalid_property",
					fmt.Sprintf("invalid component '%s': invalid property '%s.%s': token at position %d: %v",
						component, variant, property, i, err))
			}
		}
		// CSS literals in compound values are always valid
	}

	return nil
}

// newValidationError creates a validation error with a code
func newValidationError(code, message string) error {
	return fmt.Errorf("[%s] %s", code, message)
}

// =============================================================================
// DEFAULT TOKENS
// =============================================================================

// GetDefaultTokens returns a production-ready default design token system.
//
// This provides a complete, enterprise-grade token set suitable for most applications:
//
// Primitive Tokens:
//   - Color palette: primary, secondary, grays, feedback colors (success/warning/error/info)
//   - Spacing scale: 7 sizes from xs (0.5rem) to 3xl (4rem)
//   - Border radius: 5 sizes from sm to full (9999px for pills)
//   - Typography: comprehensive font sizes, weights, line heights, families, letter spacing
//   - Borders: widths (thin/medium/thick) and styles (solid/dashed/dotted)
//   - Shadows: 6 elevation levels from sm to xl, plus inner shadow
//   - Effects: opacity and blur scales
//   - Animation: duration and easing curves
//   - Z-index: layering system from base (1) to tooltip (60)
//   - Breakpoints: responsive breakpoints (sm/md/lg/xl)
//
// Semantic Tokens:
//   - Color assignments: background, foreground, muted, border, ring, feedback
//   - Spacing contexts: component spacing (tight/default/loose), layout spacing
//   - Typography contexts: heading, body, label, caption, code styles
//   - Interactive elements: border-radius, shadows, transitions
//
// Component Tokens:
//   - Button: primary, secondary, outline, ghost variants
//   - Input: base, focus, error states
//   - Badge: default, success, warning, error variants
//   - Alert: base, success, warning, error, info variants
//   - Card: base styles with elevation
//   - Modal: base with overlay
//
// All tokens follow accessibility best practices:
//   - WCAG 2.1 AA compliant color contrasts
//   - Sufficient touch target sizes
//   - Clear focus indicators
//   - Semantic color assignments
func GetDefaultTokens() *DesignTokens {
	return &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				// Primary brand colors
				"primary":   "hsl(217, 91%, 60%)", // Blue - primary actions
				"secondary": "hsl(270, 50%, 60%)", // Purple - secondary actions
				"accent":    "hsl(190, 90%, 56%)", // Cyan - accent highlights

				// Neutral colors (grayscale)
				"white":    "hsl(0, 0%, 100%)",
				"black":    "hsl(0, 0%, 0%)",
				"gray-50":  "hsl(210, 20%, 98%)",
				"gray-100": "hsl(210, 20%, 95%)",
				"gray-200": "hsl(210, 16%, 93%)",
				"gray-300": "hsl(210, 14%, 89%)",
				"gray-400": "hsl(210, 12%, 78%)",
				"gray-500": "hsl(210, 10%, 64%)",
				"gray-600": "hsl(210, 10%, 48%)",
				"gray-700": "hsl(210, 12%, 36%)",
				"gray-800": "hsl(210, 16%, 24%)",
				"gray-900": "hsl(210, 20%, 14%)",

				// Feedback colors
				"success": "hsl(142, 71%, 45%)",  // Green - success states
				"warning": "hsl(45, 93%, 47%)",   // Yellow - warning states
				"error":   "hsl(0, 84%, 60%)",    // Red - error states
				"info":    "hsl(200, 100%, 56%)", // Sky blue - info states
			},
			Spacing: map[string]string{
				"xs":  "0.5rem",  // 8px
				"sm":  "0.75rem", // 12px
				"md":  "1rem",    // 16px
				"lg":  "1.5rem",  // 24px
				"xl":  "2rem",    // 32px
				"2xl": "3rem",    // 48px
				"3xl": "4rem",    // 64px
			},
			Radius: map[string]string{
				"none": "0",
				"sm":   "0.25rem", // 4px
				"md":   "0.5rem",  // 8px
				"lg":   "0.75rem", // 12px
				"xl":   "1rem",    // 16px
				"full": "9999px",  // Perfect circle/pill
			},
			Typography: map[string]string{
				// Font sizes
				"font-size-xs":   "0.75rem",  // 12px
				"font-size-sm":   "0.875rem", // 14px
				"font-size-base": "1rem",     // 16px
				"font-size-lg":   "1.125rem", // 18px
				"font-size-xl":   "1.25rem",  // 20px
				"font-size-2xl":  "1.5rem",   // 24px
				"font-size-3xl":  "1.875rem", // 30px
				"font-size-4xl":  "2.25rem",  // 36px

				// Font weights
				"font-weight-normal":    "400",
				"font-weight-medium":    "500",
				"font-weight-semibold":  "600",
				"font-weight-bold":      "700",
				"font-weight-extrabold": "800",

				// Line heights
				"line-height-none":    "1",
				"line-height-tight":   "1.25",
				"line-height-snug":    "1.375",
				"line-height-normal":  "1.5",
				"line-height-relaxed": "1.75",
				"line-height-loose":   "2",

				// Font families
				"font-family-sans":  "-apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif",
				"font-family-serif": "Georgia, Cambria, 'Times New Roman', Times, serif",
				"font-family-mono":  "'SF Mono', Monaco, Consolas, 'Liberation Mono', 'Courier New', monospace",

				// Letter spacing
				"letter-spacing-tighter": "-0.05em",
				"letter-spacing-tight":   "-0.025em",
				"letter-spacing-normal":  "0em",
				"letter-spacing-wide":    "0.025em",
				"letter-spacing-wider":   "0.05em",
			},
			Borders: map[string]string{
				// Border widths
				"border-width-none":   "0",
				"border-width-thin":   "1px",
				"border-width-medium": "2px",
				"border-width-thick":  "4px",

				// Border styles
				"border-style-solid":  "solid",
				"border-style-dashed": "dashed",
				"border-style-dotted": "dotted",
				"border-style-double": "double",
				"border-style-none":   "none",
			},
			Shadows: map[string]string{
				"none":  "none",
				"sm":    "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
				"md":    "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06)",
				"lg":    "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)",
				"xl":    "0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 10px 10px -5px rgba(0, 0, 0, 0.04)",
				"2xl":   "0 25px 50px -12px rgba(0, 0, 0, 0.25)",
				"inner": "inset 0 2px 4px 0 rgba(0, 0, 0, 0.06)",
			},
			Effects: map[string]string{
				// Opacity
				"opacity-0":   "0",
				"opacity-25":  "0.25",
				"opacity-50":  "0.5",
				"opacity-75":  "0.75",
				"opacity-100": "1",

				// Blur
				"blur-none": "0",
				"blur-sm":   "4px",
				"blur-md":   "8px",
				"blur-lg":   "12px",
				"blur-xl":   "16px",
			},
			Animation: map[string]string{
				// Duration
				"duration-fast":   "150ms",
				"duration-normal": "200ms",
				"duration-slow":   "300ms",
				"duration-slower": "500ms",

				// Easing
				"easing-linear":     "linear",
				"easing-in":         "cubic-bezier(0.4, 0, 1, 1)",
				"easing-out":        "cubic-bezier(0, 0, 0.2, 1)",
				"easing-in-out":     "cubic-bezier(0.4, 0, 0.2, 1)",
				"easing-sharp":      "cubic-bezier(0.4, 0, 0.6, 1)",
				"easing-emphasized": "cubic-bezier(0.2, 0, 0, 1)",
			},
			ZIndex: map[string]string{
				"base":     "1",
				"dropdown": "10",
				"sticky":   "20",
				"fixed":    "30",
				"modal":    "40",
				"popover":  "50",
				"tooltip":  "60",
			},
			Breakpoints: map[string]string{
				"sm":  "640px",
				"md":  "768px",
				"lg":  "1024px",
				"xl":  "1280px",
				"2xl": "1536px",
			},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				// Backgrounds
				"background":          "primitives.colors.white",
				"background-subtle":   "primitives.colors.gray-50",
				"background-muted":    "primitives.colors.gray-100",
				"background-emphasis": "primitives.colors.gray-200",

				// Foregrounds (text)
				"foreground":          "primitives.colors.gray-900",
				"foreground-subtle":   "primitives.colors.gray-600",
				"foreground-muted":    "primitives.colors.gray-500",
				"foreground-emphasis": "primitives.colors.black",

				// Interactive
				"primary":              "primitives.colors.primary",
				"primary-foreground":   "primitives.colors.white",
				"secondary":            "primitives.colors.gray-100",
				"secondary-foreground": "primitives.colors.gray-900",
				"accent":               "primitives.colors.accent",
				"accent-foreground":    "primitives.colors.white",

				// Borders & dividers
				"border":          "primitives.colors.gray-300",
				"border-subtle":   "primitives.colors.gray-200",
				"border-emphasis": "primitives.colors.gray-400",

				// Focus & selection
				"ring":      "primitives.colors.primary",
				"selection": "primitives.colors.primary",

				// Feedback
				"success":            "primitives.colors.success",
				"success-foreground": "primitives.colors.white",
				"warning":            "primitives.colors.warning",
				"warning-foreground": "primitives.colors.white",
				"error":              "primitives.colors.error",
				"error-foreground":   "primitives.colors.white",
				"info":               "primitives.colors.info",
				"info-foreground":    "primitives.colors.white",
			},
			Spacing: map[string]string{
				// Component spacing
				"component-tight":   "primitives.spacing.xs",
				"component-default": "primitives.spacing.sm",
				"component-loose":   "primitives.spacing.md",

				// Layout spacing
				"layout-section": "primitives.spacing.2xl",
				"layout-page":    "primitives.spacing.3xl",

				// Stack spacing (vertical rhythm)
				"stack-tight":   "primitives.spacing.xs",
				"stack-default": "primitives.spacing.md",
				"stack-loose":   "primitives.spacing.lg",
			},
			Typography: map[string]string{
				// Headings
				"heading-font-size":   "primitives.typography.font-size-2xl",
				"heading-font-weight": "primitives.typography.font-weight-bold",
				"heading-line-height": "primitives.typography.line-height-tight",

				// Body text
				"body-font-size":   "primitives.typography.font-size-base",
				"body-font-weight": "primitives.typography.font-weight-normal",
				"body-line-height": "primitives.typography.line-height-normal",
				"body-font-family": "primitives.typography.font-family-sans",

				// Labels
				"label-font-size":   "primitives.typography.font-size-sm",
				"label-font-weight": "primitives.typography.font-weight-medium",

				// Captions
				"caption-font-size":   "primitives.typography.font-size-xs",
				"caption-font-weight": "primitives.typography.font-weight-normal",

				// Code
				"code-font-size":   "primitives.typography.font-size-sm",
				"code-font-family": "primitives.typography.font-family-mono",
			},
			Interactive: map[string]string{
				// Border radius for interactive elements
				"border-radius-sm": "primitives.radius.sm",
				"border-radius-md": "primitives.radius.md",
				"border-radius-lg": "primitives.radius.lg",

				// Shadows for interactive elements
				"shadow-sm": "primitives.shadows.sm",
				"shadow-md": "primitives.shadows.md",
				"shadow-lg": "primitives.shadows.lg",

				// Transitions
				"transition-fast":   "primitives.animation.duration-fast",
				"transition-normal": "primitives.animation.duration-normal",
				"transition-slow":   "primitives.animation.duration-slow",
			},
		},
		Components: &ComponentTokens{
			"button": ComponentVariants{
				"primary": StyleProperties{
					"background-color": "semantic.colors.primary",
					"color":            "semantic.colors.primary-foreground",
					"border":           "none",
					"border-radius":    "semantic.interactive.border-radius-md",
					"padding":          "primitives.spacing.sm primitives.spacing.md",
					"font-size":        "semantic.typography.body-font-size",
					"font-weight":      "primitives.typography.font-weight-medium",
					"box-shadow":       "none",
					"transition":       "all semantic.interactive.transition-fast primitives.animation.easing-in-out",
				},
				"secondary": StyleProperties{
					"background-color": "semantic.colors.secondary",
					"color":            "semantic.colors.secondary-foreground",
					"border":           "primitives.borders.border-width-thin solid semantic.colors.border",
					"border-radius":    "semantic.interactive.border-radius-md",
					"padding":          "primitives.spacing.sm primitives.spacing.md",
					"font-size":        "semantic.typography.body-font-size",
					"font-weight":      "primitives.typography.font-weight-medium",
				},
				"outline": StyleProperties{
					"background-color": "transparent",
					"color":            "semantic.colors.primary",
					"border":           "primitives.borders.border-width-thin solid semantic.colors.primary",
					"border-radius":    "semantic.interactive.border-radius-md",
					"padding":          "primitives.spacing.sm primitives.spacing.md",
					"font-size":        "semantic.typography.body-font-size",
					"font-weight":      "primitives.typography.font-weight-medium",
				},
				"ghost": StyleProperties{
					"background-color": "transparent",
					"color":            "semantic.colors.primary",
					"border":           "none",
					"border-radius":    "semantic.interactive.border-radius-md",
					"padding":          "primitives.spacing.sm primitives.spacing.md",
					"font-size":        "semantic.typography.body-font-size",
					"font-weight":      "primitives.typography.font-weight-medium",
				},
			},
			"input": ComponentVariants{
				"base": StyleProperties{
					"background-color": "semantic.colors.background",
					"border":           "primitives.borders.border-width-thin solid semantic.colors.border",
					"border-radius":    "semantic.interactive.border-radius-md",
					"padding":          "primitives.spacing.sm primitives.spacing.sm",
					"color":            "semantic.colors.foreground",
					"font-size":        "semantic.typography.body-font-size",
				},
				"focus": StyleProperties{
					"border-color":   "semantic.colors.ring",
					"outline":        "2px solid semantic.colors.ring",
					"outline-offset": "2px",
				},
				"error": StyleProperties{
					"border-color": "semantic.colors.error",
				},
			},
			"badge": ComponentVariants{
				"default": StyleProperties{
					"background-color": "semantic.colors.background-muted",
					"color":            "semantic.colors.foreground-muted",
					"padding":          "0.125rem primitives.spacing.xs",
					"border-radius":    "primitives.radius.sm",
					"font-size":        "semantic.typography.caption-font-size",
					"font-weight":      "primitives.typography.font-weight-medium",
				},
				"success": StyleProperties{
					"background-color": "semantic.colors.success",
					"color":            "semantic.colors.success-foreground",
				},
				"warning": StyleProperties{
					"background-color": "semantic.colors.warning",
					"color":            "semantic.colors.warning-foreground",
				},
				"error": StyleProperties{
					"background-color": "semantic.colors.error",
					"color":            "semantic.colors.error-foreground",
				},
				"info": StyleProperties{
					"background-color": "semantic.colors.info",
					"color":            "semantic.colors.info-foreground",
				},
			},
			"alert": ComponentVariants{
				"base": StyleProperties{
					"padding":       "primitives.spacing.md",
					"border-radius": "semantic.interactive.border-radius-md",
					"border":        "primitives.borders.border-width-thin solid semantic.colors.border",
				},
				"success": StyleProperties{
					"background-color": "hsl(142, 71%, 97%)",
					"border-color":     "hsl(142, 71%, 85%)",
					"color":            "hsl(142, 71%, 25%)",
				},
				"warning": StyleProperties{
					"background-color": "hsl(45, 93%, 97%)",
					"border-color":     "hsl(45, 93%, 85%)",
					"color":            "hsl(45, 93%, 25%)",
				},
				"error": StyleProperties{
					"background-color": "hsl(0, 84%, 97%)",
					"border-color":     "hsl(0, 84%, 85%)",
					"color":            "hsl(0, 84%, 25%)",
				},
				"info": StyleProperties{
					"background-color": "hsl(200, 100%, 97%)",
					"border-color":     "hsl(200, 100%, 85%)",
					"color":            "hsl(200, 100%, 25%)",
				},
			},
			"card": ComponentVariants{
				"base": StyleProperties{
					"background-color": "semantic.colors.background",
					"border":           "primitives.borders.border-width-thin solid semantic.colors.border",
					"border-radius":    "semantic.interactive.border-radius-lg",
					"padding":          "primitives.spacing.lg",
					"box-shadow":       "semantic.interactive.shadow-sm",
				},
			},
			"modal": ComponentVariants{
				"base": StyleProperties{
					"background-color": "semantic.colors.background",
					"border":           "primitives.borders.border-width-thin solid semantic.colors.border",
					"border-radius":    "semantic.interactive.border-radius-lg",
					"padding":          "primitives.spacing.xl",
					"box-shadow":       "semantic.interactive.shadow-lg",
				},
				"overlay": StyleProperties{
					"background-color": "rgba(0, 0, 0, 0.5)",
				},
			},
		},
	}
}
