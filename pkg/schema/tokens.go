package schema

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// TokenReference represents a reference to another token using TailwindCSS-style syntax
// No braces needed - just dot notation like "colors.primary" or "spacing.md"
type TokenReference string

// IsReference checks if the value is a token reference (contains dots indicating a path).
func (t TokenReference) IsReference() bool {
	s := string(t)
	// A token reference should contain at least one dot and not be an empty string
	// Examples: "colors.primary", "spacing.md", "typography.font-size-base"
	// Exclude CSS values like "2.25rem", "#ffffff", "hsl(...)", "rgb(...)", etc.
	if len(s) == 0 {
		return false
	}
	// Not a reference if it starts with common CSS value patterns
	if strings.HasPrefix(s, "#") || strings.HasPrefix(s, "hsl") || strings.HasPrefix(s, "rgb") {
		return false
	}
	// Not a reference if it looks like a CSS unit (number + unit)
	if strings.HasSuffix(s, "rem") || strings.HasSuffix(s, "px") || strings.HasSuffix(s, "em") ||
		strings.HasSuffix(s, "%") || strings.HasSuffix(s, "vh") || strings.HasSuffix(s, "vw") {
		return false
	}
	// Not a reference if it's a pure number (with or without decimal)
	if matched, _ := regexp.MatchString(`^\d+(\.\d+)?$`, s); matched {
		return false
	}
	// Not a reference if it contains spaces (CSS values often have spaces)
	if strings.Contains(s, " ") {
		return false
	}
	// Must contain dots to be a valid token path
	if !strings.Contains(s, ".") {
		return false
	}
	// Should have at least 2 parts when split by dots (e.g., "colors.primary")
	parts := strings.Split(s, ".")
	return len(parts) >= 2
}

// Path extracts the token path from a reference.
// For "colors.primary", it returns "colors.primary".
func (t TokenReference) Path() string {
	if !t.IsReference() {
		return string(t)
	}
	return string(t)
}

// String returns the string representation of the token reference.
func (t TokenReference) String() string {
	return string(t)
}

// Validate checks if the token reference is well-formed.
func (t TokenReference) Validate() error {
	s := string(t)
	// If it's clearly a CSS value (starts with # or contains spaces), it's valid
	if strings.HasPrefix(s, "#") || strings.Contains(s, " ") ||
		strings.HasPrefix(s, "hsl") || strings.HasPrefix(s, "rgb") {
		return nil
	}
	// If it ends with CSS units, it's valid
	if strings.HasSuffix(s, "rem") || strings.HasSuffix(s, "px") || strings.HasSuffix(s, "em") ||
		strings.HasSuffix(s, "%") || strings.HasSuffix(s, "vh") || strings.HasSuffix(s, "vw") {
		return nil
	}
	// If it's a pure number, it's valid
	if matched, _ := regexp.MatchString(`^\d+(\.\d+)?$`, s); matched {
		return nil
	}
	// If we get here, it's either a valid token reference or an invalid one
	// Token references must contain at least one dot
	if !strings.Contains(s, ".") {
		return NewValidationError("invalid_token_path", "token path must contain at least one dot")
	}
	path := t.Path()
	if path == "" {
		return NewValidationError("empty_token_path", "empty token path in reference")
	}
	// Token paths must contain at least one dot
	if !strings.Contains(path, ".") {
		return NewValidationError("invalid_token_path", fmt.Sprintf("invalid token path format: %s (must contain at least one dot)", path))
	}
	// Check for valid characters (alphanumeric, dots, underscores, hyphens)
	for _, char := range path {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') || char == '.' || char == '_' || char == '-') {
			return NewValidationError("invalid_token_char", fmt.Sprintf("invalid character '%c' in token path: %s", char, path))
		}
	}
	return nil
}

// DesignTokens is the root container for all design tokens.
// It follows a three-tier architecture for maximum flexibility and maintainability.
//
// Architecture:
//   - Primitives: Raw values (hsl(220, 50%, 50%))
//   - Semantic: Functional meaning (colors.primary, spacing.md)
//   - Components: Component-specific styles (semantic.interactive.default)
//
// Breaking Changes: The structure may evolve to add new token categories.
// Always use token references to ensure forward compatibility.
type DesignTokens struct {
	// Primitives contains raw design values - the foundation of the system
	Primitives *PrimitiveTokens `json:"primitives"`
	// Semantic contains functional token assignments
	Semantic *SemanticTokens `json:"semantic"`
	// Components contains component-specific token assignments
	Components *ComponentTokens `json:"components"`
}

// PrimitiveTokens contains the raw design values that form the foundation
// of the design system. These values should rarely change and represent
// the atomic design decisions.
type PrimitiveTokens struct {
	Colors     *ColorPrimitives      `json:"colors"`
	Spacing    *SpacingScale         `json:"spacing"`
	Typography *TypographyPrimitives `json:"typography"`
	Borders    *BorderPrimitives     `json:"borders"`
	Shadows    *ShadowScale          `json:"shadows"`
	Animations *AnimationPrimitives  `json:"animations"`
	Sizes      *SizeScale            `json:"sizes"`
	ZIndex     *ZIndexScale          `json:"zIndex"`
}

// ColorPrimitives defines the raw color palette.
// Uses HSL format for better manipulation and theming.
type ColorPrimitives struct {
	// Neutral colors - foundation for text, borders, and backgrounds
	Gray *GrayScale `json:"gray"`
	// Brand colors - full scales for primary, secondary, accent
	Blue   *ColorScale `json:"blue"`   // Primary brand color
	Purple *ColorScale `json:"purple"` // Secondary brand color
	Cyan   *ColorScale `json:"cyan"`   // Accent color
	// Semantic colors - full scales for feedback states
	Green  *ColorScale `json:"green"`  // Success states
	Red    *ColorScale `json:"red"`    // Error states
	Yellow *ColorScale `json:"yellow"` // Warning states
	Sky    *ColorScale `json:"sky"`    // Info states
	// Pure colors
	White string `json:"white"`
	Black string `json:"black"`
}

// ColorScale represents a complete color scale from 50 (lightest) to 900 (darkest).
// Follows industry standards for predictable color relationships.
type ColorScale struct {
	Scale50  string `json:"50"`
	Scale100 string `json:"100"`
	Scale200 string `json:"200"`
	Scale300 string `json:"300"`
	Scale400 string `json:"400"`
	Scale500 string `json:"500"` // Base color
	Scale600 string `json:"600"`
	Scale700 string `json:"700"`
	Scale800 string `json:"800"`
	Scale900 string `json:"900"`
}

// GrayScale represents the neutral color scale used throughout the system.
type GrayScale struct {
	Scale50  string `json:"50"`
	Scale100 string `json:"100"`
	Scale200 string `json:"200"`
	Scale300 string `json:"300"`
	Scale400 string `json:"400"`
	Scale500 string `json:"500"`
	Scale600 string `json:"600"`
	Scale700 string `json:"700"`
	Scale800 string `json:"800"`
	Scale900 string `json:"900"`
}

// SpacingScale defines the spacing scale used for margins, padding, and gaps.
type SpacingScale struct {
	None string `json:"0"`   // 0
	XS   string `json:"xs"`  // 0.25rem
	SM   string `json:"sm"`  // 0.5rem
	MD   string `json:"md"`  // 1rem - Base spacing unit
	LG   string `json:"lg"`  // 1.5rem
	XL   string `json:"xl"`  // 2rem
	XXL  string `json:"2xl"` // 3rem
	XXXL string `json:"3xl"` // 4rem
	Huge string `json:"4xl"` // 6rem
}

// TypographyPrimitives defines the raw typography values.
type TypographyPrimitives struct {
	FontSizes     *FontSizeScale      `json:"fontSizes"`
	FontWeights   *FontWeightScale    `json:"fontWeights"`
	LineHeights   *LineHeightScale    `json:"lineHeights"`
	FontFamilies  *FontFamilyScale    `json:"fontFamilies"`
	LetterSpacing *LetterSpacingScale `json:"letterSpacing"`
}

// FontSizeScale defines the typography size scale.
type FontSizeScale struct {
	XS   string `json:"xs"`   // 0.75rem
	SM   string `json:"sm"`   // 0.875rem
	Base string `json:"base"` // 1rem
	LG   string `json:"lg"`   // 1.125rem
	XL   string `json:"xl"`   // 1.25rem
	XXL  string `json:"2xl"`  // 1.5rem
	XXXL string `json:"3xl"`  // 1.875rem
	Huge string `json:"4xl"`  // 2.25rem
}

// FontWeightScale defines font weights.
type FontWeightScale struct {
	Thin       string `json:"thin"`       // 100
	ExtraLight string `json:"extraLight"` // 200
	Light      string `json:"light"`      // 300
	Normal     string `json:"normal"`     // 400
	Medium     string `json:"medium"`     // 500
	SemiBold   string `json:"semiBold"`   // 600
	Bold       string `json:"bold"`       // 700
	ExtraBold  string `json:"extraBold"`  // 800
	Black      string `json:"black"`      // 900
}

// LineHeightScale defines line heights.
type LineHeightScale struct {
	None    string `json:"none"`    // 1
	Tight   string `json:"tight"`   // 1.25
	Snug    string `json:"snug"`    // 1.375
	Normal  string `json:"normal"`  // 1.5
	Relaxed string `json:"relaxed"` // 1.625
	Loose   string `json:"loose"`   // 2
}

// FontFamilyScale defines font families.
type FontFamilyScale struct {
	Sans  string `json:"sans"`  // System UI fonts
	Serif string `json:"serif"` // Serif fonts
	Mono  string `json:"mono"`  // Monospace fonts
}

// LetterSpacingScale defines letter spacing values.
type LetterSpacingScale struct {
	Tighter string `json:"tighter"` // -0.05em
	Tight   string `json:"tight"`   // -0.025em
	Normal  string `json:"normal"`  // 0em
	Wide    string `json:"wide"`    // 0.025em
	Wider   string `json:"wider"`   // 0.05em
	Widest  string `json:"widest"`  // 0.1em
}

// BorderPrimitives defines border-related primitives.
type BorderPrimitives struct {
	Width  *BorderWidthScale  `json:"width"`
	Radius *BorderRadiusScale `json:"radius"`
	Style  *BorderStyleScale  `json:"style"`
}

// BorderWidthScale defines border widths.
type BorderWidthScale struct {
	None   string `json:"0"`      // 0px
	Thin   string `json:"thin"`   // 1px
	Medium string `json:"medium"` // 2px
	Thick  string `json:"thick"`  // 4px
}

// BorderRadiusScale defines border radius values.
type BorderRadiusScale struct {
	None string `json:"none"` // 0px
	SM   string `json:"sm"`   // 0.125rem
	Base string `json:"base"` // 0.25rem
	MD   string `json:"md"`   // 0.375rem
	LG   string `json:"lg"`   // 0.5rem
	XL   string `json:"xl"`   // 0.75rem
	XXL  string `json:"2xl"`  // 1rem
	XXXL string `json:"3xl"`  // 1.5rem
	Full string `json:"full"` // 9999px
}

// BorderStyleScale defines border styles.
type BorderStyleScale struct {
	Solid  string `json:"solid"`
	Dashed string `json:"dashed"`
	Dotted string `json:"dotted"`
	Double string `json:"double"`
	None   string `json:"none"`
}

// ShadowScale defines shadow values.
type ShadowScale struct {
	None  string `json:"none"`  // none
	SM    string `json:"sm"`    // Small shadow
	Base  string `json:"base"`  // Default shadow
	MD    string `json:"md"`    // Medium shadow
	LG    string `json:"lg"`    // Large shadow
	XL    string `json:"xl"`    // Extra large shadow
	XXL   string `json:"2xl"`   // 2x large shadow
	Inner string `json:"inner"` // Inner shadow
}

// AnimationPrimitives defines animation-related primitives.
type AnimationPrimitives struct {
	Duration *AnimationDurationScale `json:"duration"`
	Easing   *AnimationEasingScale   `json:"easing"`
}

// AnimationDurationScale defines animation durations.
type AnimationDurationScale struct {
	Fast   string `json:"fast"`   // 150ms
	Normal string `json:"normal"` // 200ms
	Slow   string `json:"slow"`   // 300ms
	Slower string `json:"slower"` // 500ms
}

// AnimationEasingScale defines animation easing functions.
type AnimationEasingScale struct {
	Linear    string `json:"linear"`
	EaseIn    string `json:"easeIn"`
	EaseOut   string `json:"easeOut"`
	EaseInOut string `json:"easeInOut"`
}

// SizeScale defines sizing values.
type SizeScale struct {
	XS   string `json:"xs"`   // 1rem
	SM   string `json:"sm"`   // 2rem
	Base string `json:"base"` // 3rem
	LG   string `json:"lg"`   // 4rem
	XL   string `json:"xl"`   // 6rem
	XXL  string `json:"2xl"`  // 8rem
	XXXL string `json:"3xl"`  // 12rem
}

// ZIndexScale defines z-index values.
type ZIndexScale struct {
	Base    string `json:"base"`    // 1
	Overlay string `json:"overlay"` // 10
	Modal   string `json:"modal"`   // 50
	Tooltip string `json:"tooltip"` // 100
}

// SemanticTokens contains semantic token assignments
type SemanticTokens struct {
	Colors      *SemanticColors      `json:"colors"`
	Spacing     *SemanticSpacing     `json:"spacing"`
	Typography  *SemanticTypography  `json:"typography"`
	Interactive *SemanticInteractive `json:"interactive"`
}

type SemanticColors struct {
	Background  *BackgroundColors  `json:"background"`
	Text        *TextColors        `json:"text"`
	Border      *BorderColors      `json:"border"`
	Interactive *InteractiveColors `json:"interactive"`
	Feedback    *FeedbackColors    `json:"feedback"`
}

type BackgroundColors struct {
	Default  TokenReference `json:"default"`
	Subtle   TokenReference `json:"subtle"`
	Emphasis TokenReference `json:"emphasis"`
	Overlay  TokenReference `json:"overlay"`
}

type TextColors struct {
	Default  TokenReference `json:"default"`
	Subtle   TokenReference `json:"subtle"`
	Emphasis TokenReference `json:"emphasis"`
	Accent   TokenReference `json:"accent"`
}

type BorderColors struct {
	Default  TokenReference `json:"default"`
	Subtle   TokenReference `json:"subtle"`
	Emphasis TokenReference `json:"emphasis"`
}

type InteractiveColors struct {
	Primary   TokenReference `json:"primary"`
	Secondary TokenReference `json:"secondary"`
	Accent    TokenReference `json:"accent"`
	Hover     TokenReference `json:"hover"`
	Active    TokenReference `json:"active"`
	Focus     TokenReference `json:"focus"`
	Disabled  TokenReference `json:"disabled"`
}

type FeedbackColors struct {
	Success TokenReference `json:"success"`
	Warning TokenReference `json:"warning"`
	Error   TokenReference `json:"error"`
	Info    TokenReference `json:"info"`
}

type SemanticSpacing struct {
	Component *ComponentSpacing `json:"component"`
	Layout    *LayoutSpacing    `json:"layout"`
}

type ComponentSpacing struct {
	Tight   TokenReference `json:"tight"`
	Default TokenReference `json:"default"`
	Loose   TokenReference `json:"loose"`
}

type LayoutSpacing struct {
	Section TokenReference `json:"section"`
	Page    TokenReference `json:"page"`
}

type SemanticTypography struct {
	Headings *HeadingTokens `json:"headings"`
	Body     *BodyTokens    `json:"body"`
	Labels   *LabelTokens   `json:"labels"`
	Captions *CaptionTokens `json:"captions"`
	Code     *CodeTokens    `json:"code"`
}

type HeadingTokens struct {
	FontSize   TokenReference `json:"fontSize"`
	FontWeight TokenReference `json:"fontWeight"`
	LineHeight TokenReference `json:"lineHeight"`
}

type BodyTokens struct {
	FontSize   TokenReference `json:"fontSize"`
	FontWeight TokenReference `json:"fontWeight"`
	LineHeight TokenReference `json:"lineHeight"`
}

type LabelTokens struct {
	FontSize   TokenReference `json:"fontSize"`
	FontWeight TokenReference `json:"fontWeight"`
}

type CaptionTokens struct {
	FontSize   TokenReference `json:"fontSize"`
	FontWeight TokenReference `json:"fontWeight"`
}

type CodeTokens struct {
	FontSize   TokenReference `json:"fontSize"`
	FontFamily TokenReference `json:"fontFamily"`
}

type SemanticInteractive struct {
	BorderRadius *InteractiveBorderRadius `json:"borderRadius"`
	Shadow       *InteractiveShadow       `json:"shadow"`
}

type InteractiveBorderRadius struct {
	Small  TokenReference `json:"small"`
	Medium TokenReference `json:"medium"`
	Large  TokenReference `json:"large"`
}

type InteractiveShadow struct {
	Small  TokenReference `json:"small"`
	Medium TokenReference `json:"medium"`
	Large  TokenReference `json:"large"`
}

// ComponentTokens contains component-specific tokens
type ComponentTokens struct {
	Button     *ButtonTokens     `json:"button"`
	Input      *InputTokens      `json:"input"`
	Card       *CardTokens       `json:"card"`
	Modal      *ModalTokens      `json:"modal"`
	Form       *FormTokens       `json:"form"`
	Table      *TableTokens      `json:"table"`
	Navigation *NavigationTokens `json:"navigation"`
}

type ButtonTokens struct {
	Background   TokenReference `json:"background"`
	Border       TokenReference `json:"border"`
	Text         TokenReference `json:"text"`
	BorderRadius TokenReference `json:"borderRadius"`
	Padding      TokenReference `json:"padding"`
	FontSize     TokenReference `json:"fontSize"`
	FontWeight   TokenReference `json:"fontWeight"`
}

type InputTokens struct {
	Background   TokenReference `json:"background"`
	Border       TokenReference `json:"border"`
	Text         TokenReference `json:"text"`
	Placeholder  TokenReference `json:"placeholder"`
	BorderRadius TokenReference `json:"borderRadius"`
	Padding      TokenReference `json:"padding"`
	FontSize     TokenReference `json:"fontSize"`
}

type CardTokens struct {
	Background   TokenReference `json:"background"`
	Border       TokenReference `json:"border"`
	Shadow       TokenReference `json:"shadow"`
	BorderRadius TokenReference `json:"borderRadius"`
	Padding      TokenReference `json:"padding"`
}

type ModalTokens struct {
	Background   TokenReference `json:"background"`
	Overlay      TokenReference `json:"overlay"`
	Border       TokenReference `json:"border"`
	Shadow       TokenReference `json:"shadow"`
	BorderRadius TokenReference `json:"borderRadius"`
	Padding      TokenReference `json:"padding"`
}

type FormTokens struct {
	Spacing      TokenReference `json:"spacing"`
	LabelSpacing TokenReference `json:"labelSpacing"`
}

type TableTokens struct {
	Background  TokenReference `json:"background"`
	AlternateBg TokenReference `json:"alternateBg"`
	Border      TokenReference `json:"border"`
	HeaderBg    TokenReference `json:"headerBg"`
	HeaderText  TokenReference `json:"headerText"`
	CellPadding TokenReference `json:"cellPadding"`
}

type NavigationTokens struct {
	Background  TokenReference `json:"background"`
	Border      TokenReference `json:"border"`
	LinkColor   TokenReference `json:"linkColor"`
	ActiveColor TokenReference `json:"activeColor"`
	Padding     TokenReference `json:"padding"`
}

// TokenRegistryInterface defines the token resolution interface
type TokenRegistryInterface interface {
	RegisterTokens(path string, tokens interface{})
	ResolveToken(ctx context.Context, reference TokenReference) (string, error)
}

// SimpleTokenRegistry is a basic implementation
type SimpleTokenRegistry struct {
	tokens map[string]interface{}
	mu     sync.RWMutex
}

// NewSimpleTokenRegistry creates a new token registry
func NewSimpleTokenRegistry() *SimpleTokenRegistry {
	return &SimpleTokenRegistry{
		tokens: make(map[string]interface{}),
	}
}

// RegisterTokens registers a set of tokens
func (tr *SimpleTokenRegistry) RegisterTokens(path string, tokens interface{}) {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	tr.tokens[path] = tokens
}

// ResolveToken resolves a token reference to its value
func (tr *SimpleTokenRegistry) ResolveToken(ctx context.Context, reference TokenReference) (string, error) {
	if !reference.IsReference() {
		return reference.String(), nil
	}

	path := reference.Path()
	parts := strings.Split(path, ".")

	tr.mu.RLock()
	defer tr.mu.RUnlock()

	// Navigate through the token structure
	current := tr.tokens
	for i, part := range parts {
		if i == len(parts)-1 {
			// Last part - get the value
			if val, ok := current[part]; ok {
				if str, ok := val.(string); ok {
					return str, nil
				}
			}
		} else {
			// Navigate deeper
			if val, ok := current[part]; ok {
				if nested, ok := val.(map[string]interface{}); ok {
					current = nested
				} else {
					break
				}
			} else {
				break
			}
		}
	}

	return "", fmt.Errorf("token not found: %s", path)
}

// GetSpacing returns spacing value by key
func GetSpacing(key string) string {
	spacingMap := map[string]string{
		"xs": "0.5rem",
		"sm": "0.75rem",
		"md": "1rem",
		"lg": "1.5rem",
		"xl": "2rem",
	}
	if val, ok := spacingMap[key]; ok {
		return val
	}
	return "1rem" // Default
}
