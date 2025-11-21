package theme

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
)

// TokenReference represents a design token value that may be either:
// - A direct CSS literal value (e.g., "#FF0000", "16px", "red", "bold")
// - A reference to another token (e.g., "primitives.colors.blue-500")
//
// Examples:
//   - CSS Literals: "#FF0000", "16px", "2rem", "solid", "rgb(255, 0, 0)"
//   - Token References: "primitives.colors.primary", "semantic.spacing.lg"
type TokenReference string

// String returns the string representation of the token reference.
func (tr TokenReference) String() string {
	return string(tr)
}

// IsReference determines if this token value is a reference to another token
// rather than a direct CSS literal value.
//
// A value is considered a reference if it:
//   - Is not empty and contains no spaces
//   - Contains at least one dot (.)
//   - Is not a CSS literal (colors, functions, units, keywords)
//   - Has at least 2 segments when split by dots
//   - All segments are non-empty
//
// Examples:
//   - Returns true: "primitives.colors.blue", "semantic.spacing.md"
//   - Returns false: "#FF0000", "16px", "red", "calc(100% - 20px)"
func (tr TokenReference) IsReference() bool {
	value := strings.TrimSpace(string(tr))

	// Empty values or values with spaces cannot be references
	if value == "" || strings.Contains(value, " ") {
		return false
	}

	// References must contain at least one dot to separate segments
	if !strings.Contains(value, ".") {
		return false
	}

	// If this looks like a CSS literal, it's not a reference
	// Example: "0.5" or "1.5rem" should not be treated as references
	if tr.isCSSLiteral(value) {
		return false
	}

	// Validate the path structure
	segments := strings.Split(value, ".")
	if len(segments) < 2 {
		return false
	}

	// All segments must be non-empty
	// Example: "primitives..colors" is invalid
	return !slices.Contains(segments, "")
}

// isCSSLiteral determines if a value is a direct CSS value rather than a token reference.
//
// CSS literals include:
//   - Hex colors: #FFF, #FF0000, #RGBA
//   - CSS functions: calc(), var(), rgb(), rgba(), hsl(), hsla(), url(), gradients
//   - Numeric values: 16px, 2rem, 1.5em, 100%, 0, -5px
//   - Named colors: red, blue, transparent, currentColor
//   - CSS keywords: inherit, initial, unset, none, auto, solid, bold, italic, normal
//
// Returns true if the value is a CSS literal, false if it might be a token reference.
func (tr TokenReference) isCSSLiteral(value string) bool {
	// Hex colors: #FFF, #FF0000, #FFFFFF00
	if strings.HasPrefix(value, "#") {
		return true
	}

	// CSS functions: calc(), var(), rgb(), url(), gradients, etc.
	cssFunctions := []string{
		"calc(", "var(", "url(",
		"rgb(", "rgba(", "hsl(", "hsla(",
		"linear-gradient(", "radial-gradient(", "conic-gradient(",
		"repeating-linear-gradient(", "repeating-radial-gradient(",
	}
	for _, fn := range cssFunctions {
		if strings.HasPrefix(value, fn) {
			return true
		}
	}

	// Numeric values with optional units: 16px, 2rem, 1.5em, 100%, 0, -5px
	// Pattern: optional minus, digits, optional decimal part, optional unit
	// Examples: 16px, 2.5rem, -10px, 0, 100%
	numericPattern := `^-?\d+(\.\d+)?[a-zA-Z%]*$`
	if matched, _ := regexp.MatchString(numericPattern, value); matched {
		return true
	}

	// CSS named colors (common subset)
	// Note: This is not exhaustive but covers the most common cases
	namedColors := map[string]bool{
		"transparent":  true,
		"currentcolor": true,
		"black":        true,
		"white":        true,
		"red":          true,
		"blue":         true,
		"green":        true,
		"yellow":       true,
		"orange":       true,
		"purple":       true,
		"pink":         true,
		"gray":         true,
		"grey":         true,
		"brown":        true,
		"cyan":         true,
		"magenta":      true,
	}
	if namedColors[strings.ToLower(value)] {
		return true
	}

	// CSS keywords that are commonly used as values
	keywords := []string{
		"inherit", "initial", "unset", "revert", "revert-layer",
		"none", "auto", "normal",
		"solid", "dashed", "dotted", "double",
		"bold", "bolder", "lighter",
		"italic", "oblique",
		"uppercase", "lowercase", "capitalize",
		"underline", "overline", "line-through",
		"left", "right", "center", "justify",
		"block", "inline", "inline-block", "flex", "grid",
		"relative", "absolute", "fixed", "sticky",
		"hidden", "visible", "scroll",
	}
	lowerValue := strings.ToLower(value)
	return slices.Contains(keywords, lowerValue)
}

// Validate checks if the token reference is valid according to the token path rules.
//
// Rules:
//   - Non-empty value
//   - If it's a reference (contains dots), it must have:
//   - At least one dot
//   - At least 2 segments
//   - All segments non-empty and valid (alphanumeric, dash, underscore only)
//
// Examples:
//   - Valid: "primitives.colors.blue", "#FF0000", "16px"
//   - Invalid: "", "primitives.", "colors..blue", "colors.blue space"
func (tr TokenReference) Validate() error {
	value := strings.TrimSpace(string(tr))

	// Token references cannot be empty
	if value == "" {
		return NewError(ErrCodeValidation, "token reference cannot be empty")
	}

	// If this is not a reference (it's a CSS literal), it's valid
	if !tr.IsReference() {
		return nil
	}

	// Validate the reference path structure
	path := value

	// References must contain at least one dot
	if !strings.Contains(path, ".") {
		return NewErrorf(ErrCodeValidation, "token path must contain at least one dot: %s", path)
	}

	// Split the path into segments and validate
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return NewErrorf(ErrCodeValidation, "token path must have at least 2 segments: %s", path)
	}

	// Validate each segment
	for i, part := range parts {
		if part == "" {
			return NewErrorf(ErrCodeValidation, "empty segment at position %d in path: %s", i, path)
		}
		if !isValidTokenSegment(part) {
			return NewErrorf(ErrCodeValidation, "invalid segment '%s' at position %d in path: %s", part, i, path)
		}
	}

	return nil
}

// isValidTokenSegment checks if a token path segment contains only valid characters.
//
// Valid characters are:
//   - Lowercase letters (a-z)
//   - Uppercase letters (A-Z)
//   - Digits (0-9)
//   - Hyphens (-)
//   - Underscores (_)
//
// Examples:
//   - Valid: "blue-500", "primary_color", "spacing16", "md"
//   - Invalid: "blue 500", "primary.color", "spacing@16"
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

// Tokens represents the complete design token system organized in three tiers:
//
//  1. Primitives: Raw CSS values (colors, spacing, typography, etc.)
//  2. Semantic: Contextual aliases that reference primitives (e.g., "text-primary")
//  3. Components: Component-specific styling with variants
//
// The three-tier system follows design token best practices:
//   - Primitives provide the foundation (e.g., "blue-500": "#3B82F6")
//   - Semantic tokens add meaning (e.g., "primary": "primitives.colors.blue-500")
//   - Components apply tokens to UI elements (e.g., Button.primary)
type Tokens struct {
	Primitives *PrimitiveTokens `json:"primitives" yaml:"primitives"`
	Semantic   *SemanticTokens  `json:"semantic" yaml:"semantic"`
	Components *ComponentTokens `json:"components" yaml:"components"`
}

// Validate ensures all token tiers are valid and consistent.
//
// It performs validation on:
//   - Primitive tokens (must be CSS literals, no references allowed)
//   - Semantic tokens (can reference primitives or other semantic tokens)
//   - Component tokens (can reference any token tier)
//
// Returns an error if any token tier fails validation.
func (t *Tokens) Validate() error {
	if t == nil {
		return NewError(ErrCodeValidation, "tokens cannot be nil")
	}

	// Validate primitives (foundational CSS values)
	if t.Primitives != nil {
		if err := t.Primitives.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid primitives", err)
		}
	}

	// Validate semantic tokens (contextual aliases)
	if t.Semantic != nil {
		if err := t.Semantic.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid semantic tokens", err)
		}
	}

	// Validate component tokens (component-specific styling)
	if t.Components != nil {
		if err := t.Components.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid component tokens", err)
		}
	}

	return nil
}

// Clone creates a deep copy of the entire token system.
//
// This is useful for:
//   - Theme inheritance (copy base theme, then override)
//   - Immutable operations (modify copy without affecting original)
//   - Multi-tenant scenarios (separate token instances per tenant)
//
// Returns nil if the source is nil.
func (t *Tokens) Clone() *Tokens {
	if t == nil {
		return nil
	}

	cloned := &Tokens{}

	if t.Primitives != nil {
		cloned.Primitives = t.Primitives.Clone()
	}
	if t.Semantic != nil {
		cloned.Semantic = t.Semantic.Clone()
	}
	if t.Components != nil {
		cloned.Components = t.Components.Clone()
	}

	return cloned
}

// PrimitiveTokens contains the foundational CSS values for the design system.
//
// These are the raw, direct CSS values that form the base of the token system.
// Primitives MUST NOT reference other tokens - they must be CSS literals only.
//
// Categories:
//   - Colors: Hex colors, rgb/rgba, named colors (e.g., "#3B82F6", "rgb(59, 130, 246)")
//   - Spacing: Size values for margins, padding (e.g., "16px", "2rem")
//   - Radius: Border radius values (e.g., "4px", "0.5rem")
//   - Typography: Font families, sizes, weights (e.g., "Inter", "16px", "bold")
//   - Borders: Border widths and styles (e.g., "1px", "solid")
//   - Shadows: Box shadow definitions (e.g., "0 2px 4px rgba(0,0,0,0.1)")
//   - Effects: Visual effects like opacity, blur (e.g., "0.8", "4px")
//   - Animation: Timing functions, durations (e.g., "ease-in-out", "200ms")
//   - ZIndex: Stacking order values (e.g., "1", "10", "100")
//   - Breakpoints: Media query breakpoints (e.g., "768px", "1024px")
//
// Example:
//
//	primitives:
//	  colors:
//	    blue-500: "#3B82F6"
//	    blue-600: "#2563EB"
//	  spacing:
//	    sm: "8px"
//	    md: "16px"
//	    lg: "24px"
type PrimitiveTokens struct {
	Colors      map[string]string `json:"colors" yaml:"colors"`
	Spacing     map[string]string `json:"spacing" yaml:"spacing"`
	Radius      map[string]string `json:"radius" yaml:"radius"`
	Typography  map[string]string `json:"typography" yaml:"typography"`
	Borders     map[string]string `json:"borders" yaml:"borders"`
	Shadows     map[string]string `json:"shadows" yaml:"shadows"`
	Effects     map[string]string `json:"effects" yaml:"effects"`
	Animation   map[string]string `json:"animation" yaml:"animation"`
	ZIndex      map[string]string `json:"zindex" yaml:"zindex"`
	Breakpoints map[string]string `json:"breakpoints" yaml:"breakpoints"`
}

// Validate ensures all primitive tokens are valid CSS literals (no references allowed).
//
// Primitives are the foundation of the design system and must contain only
// direct CSS values. References are not allowed at this level.
//
// For each category, this validates:
//   - Token values are not empty
//   - Token values are valid CSS literals (colors, sizes, keywords, etc.)
//   - Token values do NOT reference other tokens
//
// Returns an error if any primitive token is invalid or contains a reference.
func (pt *PrimitiveTokens) Validate() error {
	if pt == nil {
		return NewError(ErrCodeValidation, "primitive tokens cannot be nil")
	}

	// Map of all primitive token categories for validation
	categories := map[string]map[string]string{
		"colors":      pt.Colors,
		"spacing":     pt.Spacing,
		"radius":      pt.Radius,
		"typography":  pt.Typography,
		"borders":     pt.Borders,
		"shadows":     pt.Shadows,
		"effects":     pt.Effects,
		"animation":   pt.Animation,
		"zindex":      pt.ZIndex,
		"breakpoints": pt.Breakpoints,
	}

	// Validate each category
	for category, tokens := range categories {
		if tokens == nil {
			continue
		}

		// Validate each token in the category
		for key, value := range tokens {
			ref := TokenReference(value)

			// Check if the token value is valid (not empty, well-formed)
			if err := ref.Validate(); err != nil {
				return WrapError(ErrCodeValidation,
					fmt.Sprintf("invalid primitive token %s.%s", category, key), err)
			}

			// Primitives MUST NOT reference other tokens - they must be CSS literals
			// This ensures the foundation is stable and doesn't have circular dependencies
			if ref.IsReference() {
				return NewErrorf(ErrCodeValidation,
					"primitive token %s.%s contains reference '%s' (must be CSS literal)",
					category, key, value)
			}
		}
	}

	return nil
}

// Clone creates a deep copy of the primitive tokens.
//
// Returns nil if the source is nil.
func (pt *PrimitiveTokens) Clone() *PrimitiveTokens {
	if pt == nil {
		return nil
	}

	return &PrimitiveTokens{
		Colors:      cloneStringMap(pt.Colors),
		Spacing:     cloneStringMap(pt.Spacing),
		Radius:      cloneStringMap(pt.Radius),
		Typography:  cloneStringMap(pt.Typography),
		Borders:     cloneStringMap(pt.Borders),
		Shadows:     cloneStringMap(pt.Shadows),
		Effects:     cloneStringMap(pt.Effects),
		Animation:   cloneStringMap(pt.Animation),
		ZIndex:      cloneStringMap(pt.ZIndex),
		Breakpoints: cloneStringMap(pt.Breakpoints),
	}
}

// SemanticTokens provide contextual, meaningful names for design decisions.
//
// These tokens reference primitives (or other semantic tokens) to create a
// semantic layer that expresses intent rather than implementation.
//
// Semantic tokens CAN reference other tokens using dot notation.
//
// Categories:
//   - Colors: Purpose-based colors (e.g., "text-primary", "bg-surface", "border-error")
//   - Spacing: Contextual spacing (e.g., "button-padding", "card-gap")
//   - Typography: Text styling by purpose (e.g., "heading-1", "body-text")
//   - Interactive: Interactive element states (e.g., "hover", "active", "disabled")
//
// Example:
//
//	semantic:
//	  colors:
//	    text-primary: "primitives.colors.gray-900"
//	    text-secondary: "primitives.colors.gray-600"
//	    bg-primary: "primitives.colors.blue-500"
//	  spacing:
//	    button-padding: "primitives.spacing.md"
type SemanticTokens struct {
	Colors      map[string]string `json:"colors" yaml:"colors"`
	Spacing     map[string]string `json:"spacing" yaml:"spacing"`
	Typography  map[string]string `json:"typography" yaml:"typography"`
	Interactive map[string]string `json:"interactive" yaml:"interactive"`
}

// Validate ensures all semantic tokens are valid and well-formed.
//
// Semantic tokens can reference:
//   - Primitive tokens (e.g., "primitives.colors.blue-500")
//   - Other semantic tokens (e.g., "semantic.colors.primary")
//   - Direct CSS literals (e.g., "#FF0000", "16px")
//
// This validates:
//   - Token values are not empty
//   - Token references are well-formed (if they are references)
//   - Token paths follow the correct format
//
// Returns an error if any semantic token is invalid.
func (st *SemanticTokens) Validate() error {
	if st == nil {
		return nil
	}

	// Map of all semantic token categories for validation
	categories := map[string]map[string]string{
		"colors":      st.Colors,
		"spacing":     st.Spacing,
		"typography":  st.Typography,
		"interactive": st.Interactive,
	}

	// Validate each category
	for category, tokens := range categories {
		if tokens == nil {
			continue
		}

		// Validate each token in the category
		for key, value := range tokens {
			ref := TokenReference(value)

			// Check if the token value is valid (not empty, well-formed)
			if err := ref.Validate(); err != nil {
				return WrapError(ErrCodeValidation,
					fmt.Sprintf("invalid semantic token %s.%s", category, key), err)
			}

			// Note: Semantic tokens CAN be references or CSS literals
			// Both are valid at this level
		}
	}

	return nil
}

// Clone creates a deep copy of the semantic tokens.
//
// Returns nil if the source is nil.
func (st *SemanticTokens) Clone() *SemanticTokens {
	if st == nil {
		return nil
	}

	return &SemanticTokens{
		Colors:      cloneStringMap(st.Colors),
		Spacing:     cloneStringMap(st.Spacing),
		Typography:  cloneStringMap(st.Typography),
		Interactive: cloneStringMap(st.Interactive),
	}
}

// ComponentTokens define styling for specific UI components with variants.
//
// Structure: Component -> Variant -> Property -> Value
//
// Components can reference any token (primitives, semantic, or other components).
//
// Example:
//
//	components:
//	  Button:
//	    primary:
//	      background: "semantic.colors.bg-primary"
//	      color: "semantic.colors.text-on-primary"
//	      padding: "semantic.spacing.button-padding"
//	    secondary:
//	      background: "semantic.colors.bg-secondary"
//	      color: "semantic.colors.text-on-secondary"
//	  Card:
//	    default:
//	      background: "#FFFFFF"
//	      padding: "primitives.spacing.lg"
//	      borderRadius: "primitives.radius.md"
type ComponentTokens map[string]ComponentVariants

// ComponentVariants maps variant names to their style properties.
//
// Each component can have multiple variants (e.g., "primary", "secondary", "outline").
// Each variant contains a set of CSS properties and their values.
type ComponentVariants map[string]StyleProperties

// StyleProperties maps CSS property names to their values.
//
// Values can be:
//   - Token references (e.g., "semantic.colors.primary")
//   - Direct CSS literals (e.g., "#FF0000", "16px", "bold")
type StyleProperties map[string]string

// Validate ensures all component tokens are valid and well-structured.
//
// For each component, this validates:
//   - Component names are not empty
//   - All variants are valid (have names and properties)
//   - All properties within variants are valid (non-empty names and values)
//   - All token references are well-formed
//
// Returns an error if any component structure is invalid.
func (ct *ComponentTokens) Validate() error {
	if ct == nil {
		return nil
	}

	// Validate each component
	for componentName, variants := range *ct {
		if componentName == "" {
			return NewError(ErrCodeValidation, "component name cannot be empty")
		}

		// Validate the component's variants
		if err := variants.Validate(componentName); err != nil {
			return WrapError(ErrCodeValidation,
				fmt.Sprintf("invalid component '%s'", componentName), err)
		}
	}

	return nil
}

// Validate ensures all variants in a component are valid and well-structured.
//
// For each variant, this validates:
//   - Variant names are not empty
//   - Variants have at least one property
//   - All property names are not empty
//   - All property values are not empty
//   - All token references (if used) are well-formed
//
// Parameters:
//   - component: The name of the component (for error messages)
//
// Returns an error if any variant structure is invalid.
func (cv ComponentVariants) Validate(component string) error {
	if cv == nil {
		return nil
	}

	// Validate each variant
	for variant, props := range cv {
		// Variant names cannot be empty
		if variant == "" {
			return NewErrorf(ErrCodeValidation,
				"variant name cannot be empty in component '%s'", component)
		}

		// Variants must have at least one property
		if len(props) == 0 {
			return NewErrorf(ErrCodeValidation,
				"variant '%s' in component '%s' has no properties", variant, component)
		}

		// Validate each property in the variant
		for property, value := range props {
			// Property names cannot be empty
			if property == "" {
				return NewErrorf(ErrCodeValidation,
					"property name cannot be empty in component '%s', variant '%s'",
					component, variant)
			}

			// Property values cannot be empty
			if strings.TrimSpace(value) == "" {
				return NewErrorf(ErrCodeValidation,
					"empty value in component '%s', variant '%s', property '%s'",
					component, variant, property)
			}

			// Validate token references (if the value is a reference)
			ref := TokenReference(value)
			if err := ref.Validate(); err != nil {
				return WrapError(ErrCodeValidation,
					fmt.Sprintf("invalid token in component '%s', variant '%s', property '%s'",
						component, variant, property), err)
			}
		}
	}

	return nil
}

// Clone creates a deep copy of the component tokens.
//
// Returns nil if the source is nil.
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

// Clone creates a deep copy of the component variants.
//
// Returns nil if the source is nil.
func (cv ComponentVariants) Clone() ComponentVariants {
	if cv == nil {
		return nil
	}

	cloned := make(ComponentVariants, len(cv))
	for variant, props := range cv {
		cloned[variant] = cloneStringMap(props)
	}

	return cloned
}

// cloneStringMap creates a deep copy of a string map.
//
// This is a utility function used by all Clone methods to ensure
// complete deep copies of nested string maps.
//
// Returns nil if the source map is nil.
func cloneStringMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}

	cloned := make(map[string]string, len(m))
	maps.Copy(cloned, m)

	return cloned
}
