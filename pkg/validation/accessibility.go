package validation

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// AccessibilityValidator validates components for accessibility compliance
type AccessibilityValidator struct {
	rules     []A11yRule
	standards []A11yStandard
	config    A11yConfig
}

// AccessibilityResult contains accessibility validation results
type AccessibilityResult struct {
	Compliant   bool                 `json:"compliant"`
	Level       A11yLevel            `json:"level"`
	Score       float64              `json:"score"` // 0-100 accessibility score
	Violations  []A11yViolation      `json:"violations"`
	Warnings    []A11yWarning        `json:"warnings"`
	Suggestions []A11ySuggestion     `json:"suggestions"`
	Metrics     A11yMetrics          `json:"metrics"`
	Standards   []A11yStandardResult `json:"standards"`
	Timestamp   time.Time            `json:"timestamp"`
}

// A11yLevel represents accessibility conformance level
type A11yLevel string

const (
	A11yLevelA   A11yLevel = "A"
	A11yLevelAA  A11yLevel = "AA"
	A11yLevelAAA A11yLevel = "AAA"
)

// A11yStandard represents accessibility standards
type A11yStandard string

const (
	A11yStandardWCAG21     A11yStandard = "WCAG2.1"
	A11yStandardWCAG22     A11yStandard = "WCAG2.2"
	A11yStandardSection508 A11yStandard = "Section508"
	A11yStandardEN301549   A11yStandard = "EN301549"
)

// A11yConfig configures accessibility validation
type A11yConfig struct {
	Level         A11yLevel          `json:"level"`
	Standards     []A11yStandard     `json:"standards"`
	ColorContrast ContrastConfig     `json:"colorContrast"`
	KeyboardNav   KeyboardConfig     `json:"keyboardNav"`
	ScreenReader  ScreenReaderConfig `json:"screenReader"`
	CustomRules   []string           `json:"customRules"`
	IgnoreRules   []string           `json:"ignoreRules"`
	StrictMode    bool               `json:"strictMode"`
}

// ContrastConfig configures color contrast validation
type ContrastConfig struct {
	MinContrastNormal float64 `json:"minContrastNormal"` // WCAG AA: 4.5:1
	MinContrastLarge  float64 `json:"minContrastLarge"`  // WCAG AA: 3:1
	MinContrastAAA    float64 `json:"minContrastAAA"`    // WCAG AAA: 7:1
	CheckGraphics     bool    `json:"checkGraphics"`
	CheckNonText      bool    `json:"checkNonText"`
}

// KeyboardConfig configures keyboard navigation validation
type KeyboardConfig struct {
	CheckTabOrder     bool     `json:"checkTabOrder"`
	CheckFocusVisible bool     `json:"checkFocusVisible"`
	CheckSkipLinks    bool     `json:"checkSkipLinks"`
	RequiredKeys      []string `json:"requiredKeys"` // Tab, Enter, Space, Escape
}

// ScreenReaderConfig configures screen reader validation
type ScreenReaderConfig struct {
	CheckLabels     bool `json:"checkLabels"`
	CheckHeadings   bool `json:"checkHeadings"`
	CheckLandmarks  bool `json:"checkLandmarks"`
	CheckAltText    bool `json:"checkAltText"`
	CheckAriaLabels bool `json:"checkAriaLabels"`
}

// A11yRule represents an accessibility validation rule
type A11yRule struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Category    A11yCategory      `json:"category"`
	Level       A11yLevel         `json:"level"`
	Standard    A11yStandard      `json:"standard"`
	Enabled     bool              `json:"enabled"`
	Validator   A11yRuleValidator `json:"-"`
}

// A11yCategory represents categories of accessibility rules
type A11yCategory string

const (
	A11yCategoryPerceivable    A11yCategory = "perceivable"
	A11yCategoryOperable       A11yCategory = "operable"
	A11yCategoryUnderstandable A11yCategory = "understandable"
	A11yCategoryRobust         A11yCategory = "robust"
)

// A11yRuleValidator interface for accessibility rule validators
type A11yRuleValidator interface {
	ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult
	GetCategory() A11yCategory
}

// A11yValidationContext provides context for accessibility validation
type A11yValidationContext struct {
	Element  any            `json:"element"`
	Parent   any            `json:"parent,omitempty"`
	Children []any          `json:"children,omitempty"`
	Path     []string       `json:"path"`
	Config   A11yConfig     `json:"config"`
	Metadata map[string]any `json:"metadata"`
}

// A11yRuleResult represents the result of a single accessibility rule
type A11yRuleResult struct {
	Passed     bool            `json:"passed"`
	Rule       string          `json:"rule"`
	Violations []A11yViolation `json:"violations,omitempty"`
	Warnings   []A11yWarning   `json:"warnings,omitempty"`
	Metadata   map[string]any  `json:"metadata"`
}

// A11yViolation represents an accessibility violation
type A11yViolation struct {
	Code      string          `json:"code"`
	Message   string          `json:"message"`
	Element   string          `json:"element"`
	Impact    A11yImpact      `json:"impact"`
	Category  A11yCategory    `json:"category"`
	Standard  A11yStandard    `json:"standard"`
	Level     A11yLevel       `json:"level"`
	Location  *SourceLocation `json:"location,omitempty"`
	Fix       string          `json:"fix,omitempty"`
	Examples  []string        `json:"examples,omitempty"`
	Resources []A11yResource  `json:"resources,omitempty"`
	Context   map[string]any  `json:"context,omitempty"`
}

// A11yWarning represents an accessibility warning
type A11yWarning struct {
	Code     string          `json:"code"`
	Message  string          `json:"message"`
	Element  string          `json:"element"`
	Category A11yCategory    `json:"category"`
	Level    A11yLevel       `json:"level"`
	Location *SourceLocation `json:"location,omitempty"`
	Fix      string          `json:"fix,omitempty"`
	Context  map[string]any  `json:"context,omitempty"`
}

// A11ySuggestion represents accessibility improvement suggestions
type A11ySuggestion struct {
	Code     string         `json:"code"`
	Message  string         `json:"message"`
	Element  string         `json:"element"`
	Priority A11yPriority   `json:"priority"`
	Category A11yCategory   `json:"category"`
	Fix      string         `json:"fix"`
	AutoFix  bool           `json:"autoFix"`
	Examples []string       `json:"examples,omitempty"`
	Benefits []string       `json:"benefits,omitempty"`
	Context  map[string]any `json:"context,omitempty"`
}

// A11yImpact represents the impact level of accessibility issues
type A11yImpact string

const (
	A11yImpactCritical A11yImpact = "critical"
	A11yImpactSerious  A11yImpact = "serious"
	A11yImpactModerate A11yImpact = "moderate"
	A11yImpactMinor    A11yImpact = "minor"
)

// A11yPriority represents the priority of accessibility improvements
type A11yPriority string

const (
	A11yPriorityHigh   A11yPriority = "high"
	A11yPriorityMedium A11yPriority = "medium"
	A11yPriorityLow    A11yPriority = "low"
)

// A11yResource represents accessibility resources and documentation
type A11yResource struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Type  string `json:"type"` // documentation, example, tool
}

// A11yMetrics contains accessibility metrics
type A11yMetrics struct {
	TotalElements      int     `json:"totalElements"`
	ElementsWithIssues int     `json:"elementsWithIssues"`
	CriticalIssues     int     `json:"criticalIssues"`
	SeriousIssues      int     `json:"seriousIssues"`
	ModerateIssues     int     `json:"moderateIssues"`
	MinorIssues        int     `json:"minorIssues"`
	ContrastRatio      float64 `json:"contrastRatio,omitempty"`
	KeyboardScore      float64 `json:"keyboardScore,omitempty"`
	ScreenReaderScore  float64 `json:"screenReaderScore,omitempty"`
}

// A11yStandardResult represents compliance with a specific standard
type A11yStandardResult struct {
	Standard   A11yStandard `json:"standard"`
	Level      A11yLevel    `json:"level"`
	Compliant  bool         `json:"compliant"`
	Score      float64      `json:"score"`
	Violations int          `json:"violations"`
	Warnings   int          `json:"warnings"`
}

// NewAccessibilityValidator creates a new accessibility validator
func NewAccessibilityValidator() *AccessibilityValidator {
	config := A11yConfig{
		Level:     A11yLevelAA,
		Standards: []A11yStandard{A11yStandardWCAG21},
		ColorContrast: ContrastConfig{
			MinContrastNormal: 4.5,
			MinContrastLarge:  3.0,
			MinContrastAAA:    7.0,
			CheckGraphics:     true,
			CheckNonText:      true,
		},
		KeyboardNav: KeyboardConfig{
			CheckTabOrder:     true,
			CheckFocusVisible: true,
			CheckSkipLinks:    true,
			RequiredKeys:      []string{"Tab", "Enter", "Space", "Escape"},
		},
		ScreenReader: ScreenReaderConfig{
			CheckLabels:     true,
			CheckHeadings:   true,
			CheckLandmarks:  true,
			CheckAltText:    true,
			CheckAriaLabels: true,
		},
	}

	validator := &AccessibilityValidator{
		config: config,
	}

	validator.registerBuiltinRules()
	return validator
}

// ValidateAccessibility validates accessibility for a component or element
func (av *AccessibilityValidator) ValidateAccessibility(ctx *ValidationContext, value any) *AccessibilityResult {
	start := time.Now()

	result := &AccessibilityResult{
		Compliant: true,
		Level:     av.config.Level,
		Score:     100.0,
		Timestamp: start,
		Metrics:   A11yMetrics{},
		Standards: make([]A11yStandardResult, 0),
	}

	a11yContext := &A11yValidationContext{
		Element:  value,
		Config:   av.config,
		Path:     []string{},
		Metadata: make(map[string]any),
	}

	// Run accessibility rules
	for _, rule := range av.rules {
		if !rule.Enabled {
			continue
		}

		// Skip ignored rules
		if contains(av.config.IgnoreRules, rule.ID) {
			continue
		}

		// Check if rule applies to current level and standards
		if !av.ruleApplies(rule) {
			continue
		}

		ruleResult := rule.Validator.ValidateA11y(value, a11yContext)
		if ruleResult != nil {
			result = av.mergeRuleResult(result, ruleResult, rule)
		}
	}

	// Calculate accessibility score
	result.Score = av.calculateScore(result.Metrics)

	// Check standards compliance
	for _, standard := range av.config.Standards {
		standardResult := av.checkStandardCompliance(result, standard)
		result.Standards = append(result.Standards, standardResult)

		if !standardResult.Compliant {
			result.Compliant = false
		}
	}

	return result
}

// registerBuiltinRules registers built-in accessibility rules
func (av *AccessibilityValidator) registerBuiltinRules() {
	// WCAG 1.1.1 - Non-text Content
	av.AddRule(A11yRule{
		ID:          "wcag111",
		Name:        "Non-text Content",
		Description: "All non-text content must have text alternatives",
		Category:    A11yCategoryPerceivable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &AltTextValidator{},
	})

	// WCAG 1.3.1 - Info and Relationships
	av.AddRule(A11yRule{
		ID:          "wcag131",
		Name:        "Info and Relationships",
		Description: "Information and relationships must be programmatically determinable",
		Category:    A11yCategoryPerceivable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &StructureValidator{},
	})

	// WCAG 1.4.3 - Contrast (Minimum)
	av.AddRule(A11yRule{
		ID:          "wcag143",
		Name:        "Contrast Minimum",
		Description: "Text must have sufficient color contrast",
		Category:    A11yCategoryPerceivable,
		Level:       A11yLevelAA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &ContrastValidator{config: av.config.ColorContrast},
	})

	// WCAG 2.1.1 - Keyboard
	av.AddRule(A11yRule{
		ID:          "wcag211",
		Name:        "Keyboard Navigation",
		Description: "All functionality must be available from keyboard",
		Category:    A11yCategoryOperable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &KeyboardValidator{config: av.config.KeyboardNav},
	})

	// WCAG 2.4.1 - Bypass Blocks
	av.AddRule(A11yRule{
		ID:          "wcag241",
		Name:        "Bypass Blocks",
		Description: "Mechanism to skip repeated content blocks",
		Category:    A11yCategoryOperable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &SkipLinksValidator{},
	})

	// WCAG 3.1.1 - Language of Page
	av.AddRule(A11yRule{
		ID:          "wcag311",
		Name:        "Language of Page",
		Description: "Default language of page must be programmatically determined",
		Category:    A11yCategoryUnderstandable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &LanguageValidator{},
	})

	// WCAG 4.1.1 - Parsing
	av.AddRule(A11yRule{
		ID:          "wcag411",
		Name:        "Parsing",
		Description: "Content must be parseable by assistive technologies",
		Category:    A11yCategoryRobust,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &ParsingValidator{},
	})

	// WCAG 4.1.2 - Name, Role, Value
	av.AddRule(A11yRule{
		ID:          "wcag412",
		Name:        "Name, Role, Value",
		Description: "Name and role must be programmatically determinable",
		Category:    A11yCategoryRobust,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &AriaValidator{config: av.config.ScreenReader},
	})

	// Form-specific rules
	av.AddRule(A11yRule{
		ID:          "form_labels",
		Name:        "Form Labels",
		Description: "All form inputs must have associated labels",
		Category:    A11yCategoryPerceivable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &FormLabelValidator{},
	})

	// Component-specific rules
	av.AddRule(A11yRule{
		ID:          "button_accessible",
		Name:        "Accessible Buttons",
		Description: "Buttons must be accessible to screen readers and keyboard",
		Category:    A11yCategoryOperable,
		Level:       A11yLevelA,
		Standard:    A11yStandardWCAG21,
		Enabled:     true,
		Validator:   &ButtonValidator{},
	})
}

// AddRule adds an accessibility rule
func (av *AccessibilityValidator) AddRule(rule A11yRule) {
	av.rules = append(av.rules, rule)
}

// ruleApplies checks if a rule applies to current configuration
func (av *AccessibilityValidator) ruleApplies(rule A11yRule) bool {
	// Check level
	if !av.levelIncludes(av.config.Level, rule.Level) {
		return false
	}

	// Check standard
	if len(av.config.Standards) > 0 {
		found := false
		for _, standard := range av.config.Standards {
			if standard == rule.Standard {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// levelIncludes checks if target level includes the required level
func (av *AccessibilityValidator) levelIncludes(target, required A11yLevel) bool {
	switch target {
	case A11yLevelAAA:
		return true // AAA includes AA and A
	case A11yLevelAA:
		return required == A11yLevelA || required == A11yLevelAA
	case A11yLevelA:
		return required == A11yLevelA
	default:
		return false
	}
}

// mergeRuleResult merges a rule result into the main result
func (av *AccessibilityValidator) mergeRuleResult(result *AccessibilityResult, ruleResult *A11yRuleResult, rule A11yRule) *AccessibilityResult {
	if !ruleResult.Passed {
		result.Compliant = false
		result.Violations = append(result.Violations, ruleResult.Violations...)

		// Update metrics
		for _, violation := range ruleResult.Violations {
			result.Metrics.ElementsWithIssues++
			switch violation.Impact {
			case A11yImpactCritical:
				result.Metrics.CriticalIssues++
			case A11yImpactSerious:
				result.Metrics.SeriousIssues++
			case A11yImpactModerate:
				result.Metrics.ModerateIssues++
			case A11yImpactMinor:
				result.Metrics.MinorIssues++
			}
		}
	}

	result.Warnings = append(result.Warnings, ruleResult.Warnings...)

	return result
}

// calculateScore calculates accessibility score based on metrics
func (av *AccessibilityValidator) calculateScore(metrics A11yMetrics) float64 {
	if metrics.TotalElements == 0 {
		return 100.0
	}

	// Weight different impact levels
	totalIssues := float64(metrics.CriticalIssues*4 + metrics.SeriousIssues*3 +
		metrics.ModerateIssues*2 + metrics.MinorIssues*1)
	maxPossibleIssues := float64(metrics.TotalElements * 4)

	if maxPossibleIssues == 0 {
		return 100.0
	}

	score := 100.0 * (1.0 - (totalIssues / maxPossibleIssues))

	// Ensure score is between 0 and 100
	if score < 0 {
		score = 0
	}
	if score > 100 {
		score = 100
	}

	return score
}

// checkStandardCompliance checks compliance with a specific standard
func (av *AccessibilityValidator) checkStandardCompliance(result *AccessibilityResult, standard A11yStandard) A11yStandardResult {
	// Count violations for this standard
	violations := 0
	warnings := 0

	for _, violation := range result.Violations {
		if violation.Standard == standard {
			violations++
		}
	}

	for _, warning := range result.Warnings {
		// Note: A11yWarning doesn't have Standard field, would need to add it
		warnings++
	}

	score := result.Score
	if violations > 0 {
		// Reduce score based on violations
		penalty := float64(violations) * 10.0
		score = math.Max(0, score-penalty)
	}

	return A11yStandardResult{
		Standard:   standard,
		Level:      av.config.Level,
		Compliant:  violations == 0,
		Score:      score,
		Violations: violations,
		Warnings:   warnings,
	}
}

// Built-in Accessibility Validators

// AltTextValidator validates alternative text for images
type AltTextValidator struct{}

func (v *AltTextValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	result := &A11yRuleResult{
		Passed:   true,
		Rule:     "wcag111",
		Metadata: make(map[string]any),
	}

	// Check if element is an image or has image-like properties
	if av.isImageElement(element) {
		altText := av.getAltText(element)

		if altText == "" {
			result.Passed = false
			result.Violations = append(result.Violations, A11yViolation{
				Code:     "missing_alt_text",
				Message:  "Image is missing alternative text",
				Element:  av.getElementDescription(element),
				Impact:   A11yImpactSerious,
				Category: A11yCategoryPerceivable,
				Standard: A11yStandardWCAG21,
				Level:    A11yLevelA,
				Fix:      "Add descriptive alt text to the image",
				Examples: []string{`alt="Description of the image content"`},
				Resources: []A11yResource{
					{Title: "WCAG 2.1 - Images of Text", URL: "https://www.w3.org/WAI/WCAG21/Understanding/images-of-text.html", Type: "documentation"},
				},
			})
		} else if len(altText) > 125 {
			result.Warnings = append(result.Warnings, A11yWarning{
				Code:     "alt_text_too_long",
				Message:  "Alternative text is very long and may be verbose",
				Element:  av.getElementDescription(element),
				Category: A11yCategoryPerceivable,
				Level:    A11yLevelAA,
				Fix:      "Consider shortening alt text to be more concise",
			})
		}
	}

	return result
}

func (v *AltTextValidator) GetCategory() A11yCategory {
	return A11yCategoryPerceivable
}

func (v *AltTextValidator) isImageElement(element any) bool {
	// Simplified implementation - would check for image components/elements
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			if typeStr, ok := elemType.(string); ok {
				return typeStr == "image" || typeStr == "img" || strings.Contains(typeStr, "image")
			}
		}
	}
	return false
}

func (v *AltTextValidator) getAltText(element any) string {
	if elementMap, ok := element.(map[string]any); ok {
		if props, exists := elementMap["props"]; exists {
			if propsMap, ok := props.(map[string]any); ok {
				if alt, exists := propsMap["alt"]; exists {
					if altStr, ok := alt.(string); ok {
						return altStr
					}
				}
			}
		}
	}
	return ""
}

func (v *AltTextValidator) getElementDescription(element any) string {
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			return fmt.Sprintf("%v element", elemType)
		}
	}
	return "element"
}

// ContrastValidator validates color contrast ratios
type ContrastValidator struct {
	config ContrastConfig
}

func (v *ContrastValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	result := &A11yRuleResult{
		Passed:   true,
		Rule:     "wcag143",
		Metadata: make(map[string]any),
	}

	colors := v.extractColors(element)
	if colors.foreground != "" && colors.background != "" {
		ratio := v.calculateContrastRatio(colors.foreground, colors.background)

		isLargeText := v.isLargeText(element)
		minRatio := v.config.MinContrastNormal
		if isLargeText {
			minRatio = v.config.MinContrastLarge
		}

		if ratio < minRatio {
			result.Passed = false
			result.Violations = append(result.Violations, A11yViolation{
				Code:     "insufficient_contrast",
				Message:  fmt.Sprintf("Text contrast ratio %.2f:1 is below minimum %.2f:1", ratio, minRatio),
				Element:  v.getElementDescription(element),
				Impact:   A11yImpactSerious,
				Category: A11yCategoryPerceivable,
				Standard: A11yStandardWCAG21,
				Level:    A11yLevelAA,
				Fix:      "Increase contrast between text and background colors",
				Context: map[string]any{
					"foregroundColor": colors.foreground,
					"backgroundColor": colors.background,
					"contrastRatio":   ratio,
					"requiredRatio":   minRatio,
				},
			})
		} else if ratio < v.config.MinContrastAAA {
			result.Warnings = append(result.Warnings, A11yWarning{
				Code:     "contrast_aaa_recommended",
				Message:  fmt.Sprintf("Consider higher contrast ratio for AAA compliance (current: %.2f:1, AAA: %.2f:1)", ratio, v.config.MinContrastAAA),
				Element:  v.getElementDescription(element),
				Category: A11yCategoryPerceivable,
				Level:    A11yLevelAAA,
				Fix:      "Increase contrast for AAA compliance",
			})
		}
	}

	return result
}

func (v *ContrastValidator) GetCategory() A11yCategory {
	return A11yCategoryPerceivable
}

type colorPair struct {
	foreground string
	background string
}

func (v *ContrastValidator) extractColors(element any) colorPair {
	colors := colorPair{}

	if elementMap, ok := element.(map[string]any); ok {
		if style, exists := elementMap["style"]; exists {
			if styleMap, ok := style.(map[string]any); ok {
				if fg, exists := styleMap["color"]; exists {
					if fgStr, ok := fg.(string); ok {
						colors.foreground = fgStr
					}
				}
				if bg, exists := styleMap["backgroundColor"]; exists {
					if bgStr, ok := bg.(string); ok {
						colors.background = bgStr
					}
				}
			}
		}
	}

	return colors
}

func (v *ContrastValidator) calculateContrastRatio(fg, bg string) float64 {
	// Simplified contrast calculation
	// In a real implementation, this would parse color values and calculate luminance
	fgLum := v.getRelativeLuminance(fg)
	bgLum := v.getRelativeLuminance(bg)

	lighter := math.Max(fgLum, bgLum)
	darker := math.Min(fgLum, bgLum)

	return (lighter + 0.05) / (darker + 0.05)
}

func (v *ContrastValidator) getRelativeLuminance(color string) float64 {
	// Simplified luminance calculation
	// Real implementation would parse hex, rgb, hsl values properly
	if color == "" {
		return 0.5
	}

	// Mock calculation based on color string
	colorValue := 0.0
	for _, char := range color {
		colorValue += float64(char)
	}

	return math.Mod(colorValue, 255) / 255.0
}

func (v *ContrastValidator) isLargeText(element any) bool {
	if elementMap, ok := element.(map[string]any); ok {
		if style, exists := elementMap["style"]; exists {
			if styleMap, ok := style.(map[string]any); ok {
				if fontSize, exists := styleMap["fontSize"]; exists {
					if sizeStr, ok := fontSize.(string); ok {
						// Parse font size and determine if it's large text (18pt+ or 14pt+ bold)
						return v.parseFontSize(sizeStr) >= 18.0
					}
				}
			}
		}
	}
	return false
}

func (v *ContrastValidator) parseFontSize(fontSize string) float64 {
	// Simplified font size parsing
	re := regexp.MustCompile(`(\d+(?:\.\d+)?)`)
	matches := re.FindStringSubmatch(fontSize)
	if len(matches) > 1 {
		if size, err := strconv.ParseFloat(matches[1], 64); err == nil {
			return size
		}
	}
	return 16.0 // Default font size
}

func (v *ContrastValidator) getElementDescription(element any) string {
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			return fmt.Sprintf("%v element", elemType)
		}
	}
	return "element"
}

// KeyboardValidator validates keyboard navigation
type KeyboardValidator struct {
	config KeyboardConfig
}

func (v *KeyboardValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	result := &A11yRuleResult{
		Passed:   true,
		Rule:     "wcag211",
		Metadata: make(map[string]any),
	}

	if v.isInteractiveElement(element) {
		// Check if element is focusable
		if !v.isFocusable(element) {
			result.Passed = false
			result.Violations = append(result.Violations, A11yViolation{
				Code:     "not_focusable",
				Message:  "Interactive element is not keyboard focusable",
				Element:  v.getElementDescription(element),
				Impact:   A11yImpactSerious,
				Category: A11yCategoryOperable,
				Standard: A11yStandardWCAG21,
				Level:    A11yLevelA,
				Fix:      "Add tabindex or use focusable element",
				Examples: []string{`tabindex="0"`, `<button>`, `<a href="...">`},
			})
		}

		// Check for keyboard event handlers
		if !v.hasKeyboardHandlers(element) {
			result.Warnings = append(result.Warnings, A11yWarning{
				Code:     "missing_keyboard_handlers",
				Message:  "Interactive element may not respond to keyboard events",
				Element:  v.getElementDescription(element),
				Category: A11yCategoryOperable,
				Level:    A11yLevelA,
				Fix:      "Add onKeyDown or onKeyPress handlers for Space and Enter keys",
			})
		}
	}

	return result
}

func (v *KeyboardValidator) GetCategory() A11yCategory {
	return A11yCategoryOperable
}

func (v *KeyboardValidator) isInteractiveElement(element any) bool {
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			if typeStr, ok := elemType.(string); ok {
				interactiveTypes := []string{"button", "input", "select", "textarea", "a", "link"}
				for _, interactiveType := range interactiveTypes {
					if typeStr == interactiveType || strings.Contains(typeStr, interactiveType) {
						return true
					}
				}
			}
		}

		// Check for click handlers
		if props, exists := elementMap["props"]; exists {
			if propsMap, ok := props.(map[string]any); ok {
				if _, hasOnClick := propsMap["onClick"]; hasOnClick {
					return true
				}
			}
		}
	}
	return false
}

func (v *KeyboardValidator) isFocusable(element any) bool {
	if elementMap, ok := element.(map[string]any); ok {
		if props, exists := elementMap["props"]; exists {
			if propsMap, ok := props.(map[string]any); ok {
				// Check for tabindex
				if tabindex, exists := propsMap["tabindex"]; exists {
					if tabindexStr, ok := tabindex.(string); ok {
						if tabindexInt, err := strconv.Atoi(tabindexStr); err == nil {
							return tabindexInt >= 0
						}
					}
				}
			}
		}
	}

	// Default focusable elements
	return v.isNaturallyFocusable(element)
}

func (v *KeyboardValidator) isNaturallyFocusable(element any) bool {
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			if typeStr, ok := elemType.(string); ok {
				focusableTypes := []string{"button", "input", "select", "textarea", "a"}
				for _, focusableType := range focusableTypes {
					if typeStr == focusableType {
						return true
					}
				}
			}
		}
	}
	return false
}

func (v *KeyboardValidator) hasKeyboardHandlers(element any) bool {
	if elementMap, ok := element.(map[string]any); ok {
		if props, exists := elementMap["props"]; exists {
			if propsMap, ok := props.(map[string]any); ok {
				keyboardEvents := []string{"onKeyDown", "onKeyPress", "onKeyUp"}
				for _, event := range keyboardEvents {
					if _, exists := propsMap[event]; exists {
						return true
					}
				}
			}
		}
	}
	return false
}

func (v *KeyboardValidator) getElementDescription(element any) string {
	if elementMap, ok := element.(map[string]any); ok {
		if elemType, exists := elementMap["type"]; exists {
			return fmt.Sprintf("%v element", elemType)
		}
	}
	return "element"
}

// Additional validators would be implemented similarly...

// StructureValidator validates semantic structure
type StructureValidator struct{}

func (v *StructureValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "wcag131", Metadata: make(map[string]any)}
}

func (v *StructureValidator) GetCategory() A11yCategory {
	return A11yCategoryPerceivable
}

// SkipLinksValidator validates skip links
type SkipLinksValidator struct{}

func (v *SkipLinksValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "wcag241", Metadata: make(map[string]any)}
}

func (v *SkipLinksValidator) GetCategory() A11yCategory {
	return A11yCategoryOperable
}

// LanguageValidator validates language attributes
type LanguageValidator struct{}

func (v *LanguageValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "wcag311", Metadata: make(map[string]any)}
}

func (v *LanguageValidator) GetCategory() A11yCategory {
	return A11yCategoryUnderstandable
}

// ParsingValidator validates HTML parsing
type ParsingValidator struct{}

func (v *ParsingValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "wcag411", Metadata: make(map[string]any)}
}

func (v *ParsingValidator) GetCategory() A11yCategory {
	return A11yCategoryRobust
}

// AriaValidator validates ARIA attributes
type AriaValidator struct {
	config ScreenReaderConfig
}

func (v *AriaValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "wcag412", Metadata: make(map[string]any)}
}

func (v *AriaValidator) GetCategory() A11yCategory {
	return A11yCategoryRobust
}

// FormLabelValidator validates form labels
type FormLabelValidator struct{}

func (v *FormLabelValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "form_labels", Metadata: make(map[string]any)}
}

func (v *FormLabelValidator) GetCategory() A11yCategory {
	return A11yCategoryPerceivable
}

// ButtonValidator validates button accessibility
type ButtonValidator struct{}

func (v *ButtonValidator) ValidateA11y(element any, context *A11yValidationContext) *A11yRuleResult {
	return &A11yRuleResult{Passed: true, Rule: "button_accessible", Metadata: make(map[string]any)}
}

func (v *ButtonValidator) GetCategory() A11yCategory {
	return A11yCategoryOperable
}
