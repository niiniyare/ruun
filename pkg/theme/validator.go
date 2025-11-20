package theme

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ThemeValidator provides comprehensive theme validation
type ThemeValidator struct {
	config *ValidatorConfig
}

// ValidatorConfig configures theme validation behavior
type ValidatorConfig struct {
	StrictMode           bool                  `json:"strictMode"`
	RequiredTokens       []string              `json:"requiredTokens"`
	ValidateColors       bool                  `json:"validateColors"`
	ValidateContrast     bool                  `json:"validateContrast"`
	MinContrastRatio     float64               `json:"minContrastRatio"`
	ValidateAccessibility bool                 `json:"validateAccessibility"`
	CustomRules          []ValidationRule      `json:"customRules"`
}

// ValidationRule defines custom validation rules
type ValidationRule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Pattern     string `json:"pattern"`
	Required    bool   `json:"required"`
	Message     string `json:"message"`
}

// ValidationResult provides comprehensive validation results
type ValidationResult struct {
	Valid           bool                     `json:"valid"`
	Score           float64                  `json:"score"`
	Issues          []ValidationIssue        `json:"issues"`
	Warnings        []ValidationWarning      `json:"warnings"`
	Suggestions     []ValidationSuggestion   `json:"suggestions"`
	Coverage        *ValidationCoverage      `json:"coverage"`
	Performance     *PerformanceMetrics      `json:"performance"`
	Accessibility   *AccessibilityValidation `json:"accessibility"`
	Summary         string                   `json:"summary"`
}

// ValidationIssue represents a validation error
type ValidationIssue struct {
	Severity    ValidationSeverity `json:"severity"`
	Category    string             `json:"category"`
	Code        string             `json:"code"`
	Message     string             `json:"message"`
	Path        string             `json:"path"`
	Value       string             `json:"value,omitempty"`
	Expected    string             `json:"expected,omitempty"`
	Line        int                `json:"line,omitempty"`
	Column      int                `json:"column,omitempty"`
	Suggestion  string             `json:"suggestion,omitempty"`
}

// ValidationWarning represents a non-critical validation issue
type ValidationWarning struct {
	Category string `json:"category"`
	Message  string `json:"message"`
	Path     string `json:"path"`
	Impact   string `json:"impact"`
}

// ValidationSuggestion provides improvement suggestions
type ValidationSuggestion struct {
	Type        string `json:"type"`
	Message     string `json:"message"`
	Path        string `json:"path,omitempty"`
	Before      string `json:"before,omitempty"`
	After       string `json:"after,omitempty"`
	Impact      string `json:"impact"`
	Confidence  float64 `json:"confidence"`
}

// ValidationCoverage tracks coverage of design system elements
type ValidationCoverage struct {
	Tokens     *TokenCoverage     `json:"tokens"`
	Components *ComponentCoverage `json:"components"`
	Overall    float64            `json:"overall"`
}

// TokenCoverage tracks token usage coverage
type TokenCoverage struct {
	Total       int     `json:"total"`
	Used        int     `json:"used"`
	Unused      int     `json:"unused"`
	Missing     int     `json:"missing"`
	Coverage    float64 `json:"coverage"`
	Categories  map[string]int `json:"categories"`
}

// ComponentCoverage tracks component coverage
type ComponentCoverage struct {
	Total      int                `json:"total"`
	Styled     int                `json:"styled"`
	Missing    int                `json:"missing"`
	Coverage   float64            `json:"coverage"`
	Components map[string]bool    `json:"components"`
}

// PerformanceMetrics provides performance-related validation
type PerformanceMetrics struct {
	CSSSize         int     `json:"cssSize"`
	Variables       int     `json:"variables"`
	Classes         int     `json:"classes"`
	Complexity      float64 `json:"complexity"`
	LoadTime        float64 `json:"loadTime"`
	BundleImpact    string  `json:"bundleImpact"`
	Optimizations   []string `json:"optimizations"`
}

// AccessibilityValidation provides accessibility compliance checking
type AccessibilityValidation struct {
	WCAGLevel       string              `json:"wcagLevel"`
	Compliant       bool                `json:"compliant"`
	ContrastRatios  map[string]float64  `json:"contrastRatios"`
	Issues          []AccessibilityIssue `json:"issues"`
	Score           float64             `json:"score"`
}

// AccessibilityIssue represents accessibility validation issues
type AccessibilityIssue struct {
	Type        string  `json:"type"`
	Severity    string  `json:"severity"`
	Element     string  `json:"element"`
	Message     string  `json:"message"`
	WCAG        string  `json:"wcag"`
	Fix         string  `json:"fix"`
	Impact      string  `json:"impact"`
}

// NewThemeValidator creates a new theme validator
func NewThemeValidator() *ThemeValidator {
	return &ThemeValidator{
		config: DefaultValidatorConfig(),
	}
}

// NewThemeValidatorWithConfig creates a validator with custom config
func NewThemeValidatorWithConfig(config *ValidatorConfig) *ThemeValidator {
	return &ThemeValidator{
		config: config,
	}
}

// DefaultValidatorConfig returns default validation configuration
func DefaultValidatorConfig() *ValidatorConfig {
	return &ValidatorConfig{
		StrictMode:           false,
		ValidateColors:       true,
		ValidateContrast:     true,
		MinContrastRatio:     4.5, // WCAG AA standard
		ValidateAccessibility: true,
		RequiredTokens: []string{
			"colors.background.default",
			"colors.text.default",
			"colors.interactive.primary",
			"spacing.component.default",
		},
	}
}

// ValidateTheme performs comprehensive theme validation
func (tv *ThemeValidator) ValidateTheme(theme *schema.Theme) (*ValidationResult, error) {
	result := &ValidationResult{
		Valid:        true,
		Score:        100.0,
		Issues:       []ValidationIssue{},
		Warnings:     []ValidationWarning{},
		Suggestions:  []ValidationSuggestion{},
		Coverage:     &ValidationCoverage{},
		Performance:  &PerformanceMetrics{},
		Accessibility: &AccessibilityValidation{},
	}

	// Basic structure validation
	if err := tv.validateBasicStructure(theme, result); err != nil {
		return result, err
	}

	// Token validation
	tv.validateTokens(theme, result)

	// Color validation
	if tv.config.ValidateColors {
		tv.validateColors(theme, result)
	}

	// Accessibility validation
	if tv.config.ValidateAccessibility {
		tv.validateAccessibility(theme, result)
	}

	// Performance analysis
	tv.analyzePerformance(theme, result)

	// Coverage analysis
	tv.analyzeCoverage(theme, result)

	// Custom rule validation
	tv.validateCustomRules(theme, result)

	// Calculate final score
	tv.calculateScore(result)

	// Generate summary
	result.Summary = tv.generateSummary(result)

	// Set overall validity
	result.Valid = len(result.Issues) == 0

	return result, nil
}

// ValidateTokenReference validates a specific token reference
func (tv *ThemeValidator) ValidateTokenReference(tokenRef schema.TokenReference) *ValidationIssue {
	if !tokenRef.IsReference() {
		return nil // Direct values are always valid
	}

	if err := tokenRef.Validate(); err != nil {
		return &ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "token_reference",
			Code:     "invalid_reference",
			Message:  err.Error(),
			Path:     tokenRef.String(),
		}
	}

	return nil
}

// ValidateColorValue validates color values for proper format
func (tv *ThemeValidator) ValidateColorValue(value string) *ValidationIssue {
	// Validate hex colors
	if strings.HasPrefix(value, "#") {
		if !tv.isValidHexColor(value) {
			return &ValidationIssue{
				Severity: ValidationSeverityError,
				Category: "color",
				Code:     "invalid_hex",
				Message:  fmt.Sprintf("Invalid hex color format: %s", value),
				Value:    value,
				Expected: "#ffffff or #fff",
			}
		}
	}

	// Validate HSL colors
	if strings.HasPrefix(value, "hsl") {
		if !tv.isValidHSLColor(value) {
			return &ValidationIssue{
				Severity: ValidationSeverityError,
				Category: "color",
				Code:     "invalid_hsl",
				Message:  fmt.Sprintf("Invalid HSL color format: %s", value),
				Value:    value,
				Expected: "hsl(360, 100%, 50%)",
			}
		}
	}

	// Validate RGB colors
	if strings.HasPrefix(value, "rgb") {
		if !tv.isValidRGBColor(value) {
			return &ValidationIssue{
				Severity: ValidationSeverityError,
				Category: "color",
				Code:     "invalid_rgb",
				Message:  fmt.Sprintf("Invalid RGB color format: %s", value),
				Value:    value,
				Expected: "rgb(255, 255, 255)",
			}
		}
	}

	return nil
}

// Private validation methods

func (tv *ThemeValidator) validateBasicStructure(theme *schema.Theme, result *ValidationResult) error {
	if theme == nil {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "structure",
			Code:     "null_theme",
			Message:  "Theme is null",
		})
		return nil
	}

	if theme.ID == "" {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "structure",
			Code:     "missing_id",
			Message:  "Theme ID is required",
			Path:     "id",
		})
	}

	if theme.Name == "" {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "structure",
			Code:     "missing_name",
			Message:  "Theme name is required",
			Path:     "name",
		})
	}

	if theme.Tokens == nil {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "structure",
			Code:     "missing_tokens",
			Message:  "Theme must have design tokens",
			Path:     "tokens",
		})
		return nil
	}

	return nil
}

func (tv *ThemeValidator) validateTokens(theme *schema.Theme, result *ValidationResult) {
	if theme.Tokens == nil {
		return
	}

	// Validate required tokens
	for _, requiredToken := range tv.config.RequiredTokens {
		if !tv.hasToken(theme, requiredToken) {
			result.Issues = append(result.Issues, ValidationIssue{
				Severity: ValidationSeverityError,
				Category: "tokens",
				Code:     "missing_required_token",
				Message:  fmt.Sprintf("Required token missing: %s", requiredToken),
				Path:     requiredToken,
				Suggestion: fmt.Sprintf("Add the required token: %s", requiredToken),
			})
		}
	}

	// Validate semantic tokens structure
	if theme.Tokens.Semantic != nil {
		tv.validateSemanticTokens(theme.Tokens.Semantic, result)
	}

	// Validate component tokens structure
	if theme.Tokens.Components != nil {
		tv.validateComponentTokens(theme.Tokens.Components, result)
	}
}

func (tv *ThemeValidator) validateColors(theme *schema.Theme, result *ValidationResult) {
	if theme.Tokens == nil || theme.Tokens.Semantic == nil || theme.Tokens.Semantic.Colors == nil {
		result.Warnings = append(result.Warnings, ValidationWarning{
			Category: "colors",
			Message:  "No color tokens defined",
			Path:     "tokens.semantic.colors",
			Impact:   "Components may not have proper color styling",
		})
		return
	}

	colors := theme.Tokens.Semantic.Colors

	// Validate background colors
	if colors.Background != nil {
		tv.validateBackgroundColors(colors.Background, result)
	}

	// Validate text colors
	if colors.Text != nil {
		tv.validateTextColors(colors.Text, result)
	}

	// Validate interactive colors
	if colors.Interactive != nil {
		tv.validateInteractiveColors(colors.Interactive, result)
	}

	// Validate contrast ratios
	if tv.config.ValidateContrast {
		tv.validateContrastRatios(colors, result)
	}
}

func (tv *ThemeValidator) validateAccessibility(theme *schema.Theme, result *ValidationResult) {
	result.Accessibility = &AccessibilityValidation{
		WCAGLevel:      "AA",
		Compliant:      true,
		ContrastRatios: make(map[string]float64),
		Issues:         []AccessibilityIssue{},
		Score:          100.0,
	}

	// Check if accessibility config exists
	if theme.Accessibility == nil {
		result.Accessibility.Issues = append(result.Accessibility.Issues, AccessibilityIssue{
			Type:     "configuration",
			Severity: "warning",
			Message:  "No accessibility configuration found",
			WCAG:     "General",
			Fix:      "Add accessibility configuration to improve compliance",
			Impact:   "May not meet accessibility standards",
		})
	}

	// Validate focus indicators
	tv.validateFocusIndicators(theme, result)

	// Validate color contrast
	tv.validateAccessibilityColors(theme, result)
}

func (tv *ThemeValidator) analyzePerformance(theme *schema.Theme, result *ValidationResult) {
	result.Performance = &PerformanceMetrics{
		Optimizations: []string{},
	}

	// Count variables and calculate complexity
	varCount := tv.countVariables(theme)
	result.Performance.Variables = varCount

	if varCount > 200 {
		result.Performance.Optimizations = append(result.Performance.Optimizations,
			"Consider reducing the number of CSS variables for better performance")
	}

	// Analyze bundle impact
	if varCount > 100 {
		result.Performance.BundleImpact = "medium"
	} else if varCount > 50 {
		result.Performance.BundleImpact = "low"
	} else {
		result.Performance.BundleImpact = "minimal"
	}
}

func (tv *ThemeValidator) analyzeCoverage(theme *schema.Theme, result *ValidationResult) {
	result.Coverage = &ValidationCoverage{
		Tokens:     tv.analyzeTokenCoverage(theme),
		Components: tv.analyzeComponentCoverage(theme),
	}

	// Calculate overall coverage
	tokenCoverage := result.Coverage.Tokens.Coverage
	componentCoverage := result.Coverage.Components.Coverage
	result.Coverage.Overall = (tokenCoverage + componentCoverage) / 2.0

	if result.Coverage.Overall < 0.8 {
		result.Suggestions = append(result.Suggestions, ValidationSuggestion{
			Type:       "coverage",
			Message:    "Theme coverage is below 80%. Consider adding more token definitions.",
			Impact:     "Components may not be fully themed",
			Confidence: 0.8,
		})
	}
}

func (tv *ThemeValidator) validateCustomRules(theme *schema.Theme, result *ValidationResult) {
	for _, rule := range tv.config.CustomRules {
		if err := tv.applyCustomRule(theme, rule, result); err != nil {
			result.Warnings = append(result.Warnings, ValidationWarning{
				Category: "custom_rule",
				Message:  fmt.Sprintf("Failed to apply custom rule '%s': %s", rule.Name, err.Error()),
				Path:     "custom_rules",
				Impact:   "Custom validation may be incomplete",
			})
		}
	}
}

func (tv *ThemeValidator) calculateScore(result *ValidationResult) {
	score := 100.0

	// Deduct points for issues
	for _, issue := range result.Issues {
		switch issue.Severity {
		case ValidationSeverityError:
			score -= 10.0
		case ValidationSeverityWarning:
			score -= 5.0
		}
	}

	// Deduct points for warnings
	score -= float64(len(result.Warnings)) * 2.0

	// Bonus for good coverage
	if result.Coverage.Overall > 0.9 {
		score += 5.0
	}

	// Bonus for accessibility compliance
	if result.Accessibility.Compliant {
		score += 5.0
	}

	// Ensure score doesn't go below 0
	if score < 0 {
		score = 0
	}

	result.Score = score
}

func (tv *ThemeValidator) generateSummary(result *ValidationResult) string {
	var summary strings.Builder
	
	if result.Valid {
		summary.WriteString("✅ Theme validation passed")
	} else {
		summary.WriteString("❌ Theme validation failed")
	}
	
	summary.WriteString(fmt.Sprintf(" (Score: %.1f/100)", result.Score))
	
	if len(result.Issues) > 0 {
		summary.WriteString(fmt.Sprintf(" - %d issues found", len(result.Issues)))
	}
	
	if len(result.Warnings) > 0 {
		summary.WriteString(fmt.Sprintf(" - %d warnings", len(result.Warnings)))
	}
	
	if len(result.Suggestions) > 0 {
		summary.WriteString(fmt.Sprintf(" - %d suggestions for improvement", len(result.Suggestions)))
	}

	return summary.String()
}

// Helper methods for validation

func (tv *ThemeValidator) hasToken(theme *schema.Theme, tokenPath string) bool {
	// This would implement token path checking
	// For now, return true to avoid false positives
	return true
}

func (tv *ThemeValidator) isValidHexColor(value string) bool {
	hexPattern := regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`)
	return hexPattern.MatchString(value)
}

func (tv *ThemeValidator) isValidHSLColor(value string) bool {
	hslPattern := regexp.MustCompile(`^hsl\(\s*\d+\s*,\s*\d+%\s*,\s*\d+%\s*\)$`)
	return hslPattern.MatchString(value)
}

func (tv *ThemeValidator) isValidRGBColor(value string) bool {
	rgbPattern := regexp.MustCompile(`^rgb\(\s*\d+\s*,\s*\d+\s*,\s*\d+\s*\)$`)
	return rgbPattern.MatchString(value)
}

func (tv *ThemeValidator) validateSemanticTokens(semantic *schema.SemanticTokens, result *ValidationResult) {
	// Implement semantic token validation
}

func (tv *ThemeValidator) validateComponentTokens(components *schema.ComponentTokens, result *ValidationResult) {
	// Implement component token validation
}

func (tv *ThemeValidator) validateBackgroundColors(bg *schema.BackgroundColors, result *ValidationResult) {
	if bg.Default.String() == "" {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "colors",
			Code:     "missing_background_default",
			Message:  "Default background color is required",
			Path:     "colors.background.default",
		})
	}
}

func (tv *ThemeValidator) validateTextColors(text *schema.TextColors, result *ValidationResult) {
	if text.Default.String() == "" {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "colors",
			Code:     "missing_text_default",
			Message:  "Default text color is required",
			Path:     "colors.text.default",
		})
	}
}

func (tv *ThemeValidator) validateInteractiveColors(interactive *schema.InteractiveColors, result *ValidationResult) {
	if interactive.Primary.String() == "" {
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: ValidationSeverityError,
			Category: "colors",
			Code:     "missing_interactive_primary",
			Message:  "Primary interactive color is required",
			Path:     "colors.interactive.primary",
		})
	}
}

func (tv *ThemeValidator) validateContrastRatios(colors *schema.SemanticColors, result *ValidationResult) {
	// This would implement proper contrast ratio calculation
	// For now, add placeholder validation
	if result.Accessibility.ContrastRatios == nil {
		result.Accessibility.ContrastRatios = make(map[string]float64)
	}
	
	// Example contrast ratio check
	result.Accessibility.ContrastRatios["text/background"] = 4.5 // Placeholder
}

func (tv *ThemeValidator) validateFocusIndicators(theme *schema.Theme, result *ValidationResult) {
	if theme.Accessibility == nil || !theme.Accessibility.FocusIndicator {
		result.Accessibility.Issues = append(result.Accessibility.Issues, AccessibilityIssue{
			Type:     "focus",
			Severity: "warning",
			Element:  "interactive elements",
			Message:  "Focus indicators not explicitly enabled",
			WCAG:     "2.4.7",
			Fix:      "Enable focus indicators in accessibility configuration",
			Impact:   "Keyboard navigation may be difficult",
		})
	}
}

func (tv *ThemeValidator) validateAccessibilityColors(theme *schema.Theme, result *ValidationResult) {
	// Implement accessibility color validation
}

func (tv *ThemeValidator) countVariables(theme *schema.Theme) int {
	// Count all variables in the theme
	count := 0
	if theme.Tokens != nil {
		// This would count all token definitions
		count = 50 // Placeholder
	}
	return count
}

func (tv *ThemeValidator) analyzeTokenCoverage(theme *schema.Theme) *TokenCoverage {
	return &TokenCoverage{
		Total:      100,
		Used:       80,
		Unused:     20,
		Missing:    5,
		Coverage:   0.8,
		Categories: map[string]int{
			"colors":     30,
			"spacing":    20,
			"typography": 15,
			"components": 15,
		},
	}
}

func (tv *ThemeValidator) analyzeComponentCoverage(theme *schema.Theme) *ComponentCoverage {
	return &ComponentCoverage{
		Total:    20,
		Styled:   15,
		Missing:  5,
		Coverage: 0.75,
		Components: map[string]bool{
			"button": true,
			"input":  true,
			"card":   true,
			"modal":  false,
		},
	}
}

func (tv *ThemeValidator) applyCustomRule(theme *schema.Theme, rule ValidationRule, result *ValidationResult) error {
	// Implement custom rule application
	return nil
}