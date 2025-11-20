package theme

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// Validator validates themes for correctness, completeness, and best practices.
type Validator struct {
	config *ValidatorConfig
	rules  []ValidationRule
	mu     sync.RWMutex
	stats  *ValidatorStats
}

// ValidatorConfig configures the validator behavior.
type ValidatorConfig struct {
	StrictMode           bool
	CheckAccessibility   bool
	CheckPerformance     bool
	CheckBestPractices   bool
	MaxTokenDepth        int
	MaxThemeSize         int64
	RequiredCategories   []string
	CustomRules          []ValidationRule
}

// DefaultValidatorConfig returns production-ready validator configuration.
func DefaultValidatorConfig() *ValidatorConfig {
	return &ValidatorConfig{
		StrictMode:          false,
		CheckAccessibility:  true,
		CheckPerformance:    true,
		CheckBestPractices:  true,
		MaxTokenDepth:       5,
		MaxThemeSize:        10 << 20, // 10MB
		RequiredCategories:  []string{"colors", "spacing", "typography"},
		CustomRules:         []ValidationRule{},
	}
}

// ValidationRule is a custom validation rule function.
type ValidationRule func(*Theme) []ValidationIssue

// NewValidator creates a new theme validator with the given configuration.
func NewValidator(config *ValidatorConfig) *Validator {
	if config == nil {
		config = DefaultValidatorConfig()
	}

	v := &Validator{
		config: config,
		rules:  make([]ValidationRule, 0),
		stats:  &ValidatorStats{},
	}

	v.registerDefaultRules()

	for _, rule := range config.CustomRules {
		v.rules = append(v.rules, rule)
	}

	return v
}

// Validate validates a theme and returns the validation result.
func (v *Validator) Validate(theme *Theme) *ValidationResult {
	if theme == nil {
		return &ValidationResult{
			Valid:  false,
			Issues: []ValidationIssue{{Severity: "error", Message: "theme cannot be nil"}},
		}
	}

	v.mu.Lock()
	v.stats.TotalValidations++
	v.mu.Unlock()

	result := &ValidationResult{
		Valid:  true,
		Issues: make([]ValidationIssue, 0),
	}

	if err := theme.Validate(); err != nil {
		result.Valid = false
		result.Issues = append(result.Issues, ValidationIssue{
			Severity: "error",
			Message:  fmt.Sprintf("theme validation failed: %s", err.Error()),
			Path:     "theme",
		})
	}

	for _, rule := range v.rules {
		issues := rule(theme)
		result.Issues = append(result.Issues, issues...)
	}

	for _, issue := range result.Issues {
		if issue.Severity == "error" {
			result.Valid = false
			v.mu.Lock()
			v.stats.FailedValidations++
			v.mu.Unlock()
			break
		}
	}

	return result
}

// registerDefaultRules registers the default validation rules.
func (v *Validator) registerDefaultRules() {
	v.rules = append(v.rules, v.validateRequiredCategories)
	v.rules = append(v.rules, v.validateTokenReferences)
	v.rules = append(v.rules, v.validateColorContrast)
	v.rules = append(v.rules, v.validateTokenDepth)
	v.rules = append(v.rules, v.validateNamingConventions)
	
	if v.config.CheckAccessibility {
		v.rules = append(v.rules, v.validateAccessibility)
	}
	if v.config.CheckPerformance {
		v.rules = append(v.rules, v.validatePerformance)
	}
	if v.config.CheckBestPractices {
		v.rules = append(v.rules, v.validateBestPractices)
	}
}

// validateRequiredCategories checks that required token categories are present.
func (v *Validator) validateRequiredCategories(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Tokens == nil || theme.Tokens.Primitives == nil {
		issues = append(issues, ValidationIssue{
			Severity: "error",
			Message:  "theme must have primitive tokens",
			Path:     "tokens.primitives",
		})
		return issues
	}

	primitives := theme.Tokens.Primitives
	categoryMap := map[string]map[string]string{
		"colors":     primitives.Colors,
		"spacing":    primitives.Spacing,
		"typography": primitives.Typography,
	}

	for _, required := range v.config.RequiredCategories {
		tokens, exists := categoryMap[required]
		if !exists || tokens == nil || len(tokens) == 0 {
			issues = append(issues, ValidationIssue{
				Severity: "error",
				Message:  fmt.Sprintf("required token category '%s' is missing or empty", required),
				Path:     fmt.Sprintf("tokens.primitives.%s", required),
			})
		}
	}

	return issues
}

// validateTokenReferences validates that all token references resolve correctly.
func (v *Validator) validateTokenReferences(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Tokens == nil {
		return issues
	}

	if theme.Tokens.Semantic != nil {
		issues = append(issues, v.validateSemanticReferences(theme.Tokens)...)
	}

	if theme.Tokens.Components != nil {
		issues = append(issues, v.validateComponentReferences(theme.Tokens)...)
	}

	return issues
}

// validateSemanticReferences validates semantic token references.
func (v *Validator) validateSemanticReferences(tokens *Tokens) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if tokens.Semantic == nil {
		return issues
	}

	categories := map[string]map[string]string{
		"colors":      tokens.Semantic.Colors,
		"spacing":     tokens.Semantic.Spacing,
		"typography":  tokens.Semantic.Typography,
		"interactive": tokens.Semantic.Interactive,
	}

	for category, tokenMap := range categories {
		if tokenMap == nil {
			continue
		}
		for key, value := range tokenMap {
			ref := TokenReference(value)
			if ref.IsReference() {
				if err := ref.Validate(); err != nil {
					issues = append(issues, ValidationIssue{
						Severity: "error",
						Message:  fmt.Sprintf("invalid reference: %s", err.Error()),
						Path:     fmt.Sprintf("tokens.semantic.%s.%s", category, key),
					})
				}
			}
		}
	}

	return issues
}

// validateComponentReferences validates component token references.
func (v *Validator) validateComponentReferences(tokens *Tokens) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if tokens.Components == nil || len(*tokens.Components) == 0 {
		return issues
	}

	for component, variants := range *tokens.Components {
		for variant, properties := range variants {
			for property, value := range properties {
				parts := strings.Fields(value)
				for _, part := range parts {
					ref := TokenReference(part)
					if ref.IsReference() {
						if err := ref.Validate(); err != nil {
							issues = append(issues, ValidationIssue{
								Severity: "error",
								Message:  fmt.Sprintf("invalid reference: %s", err.Error()),
								Path:     fmt.Sprintf("tokens.components.%s.%s.%s", component, variant, property),
							})
						}
					}
				}
			}
		}
	}

	return issues
}

// validateColorContrast checks color contrast ratios for accessibility.
func (v *Validator) validateColorContrast(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Accessibility != nil && theme.Accessibility.MinContrastRatio > 0 {
		if theme.Accessibility.MinContrastRatio < 4.5 {
			issues = append(issues, ValidationIssue{
				Severity: "warning",
				Message:  "minimum contrast ratio should be at least 4.5:1 for WCAG AA compliance",
				Path:     "accessibility.minContrastRatio",
			})
		}
	}

	return issues
}

// validateTokenDepth checks that token reference chains don't exceed maximum depth.
func (v *Validator) validateTokenDepth(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Tokens == nil || theme.Tokens.Semantic == nil {
		return issues
	}

	maxAllowedDepth := v.config.MaxTokenDepth

	categories := map[string]map[string]string{
		"colors":      theme.Tokens.Semantic.Colors,
		"spacing":     theme.Tokens.Semantic.Spacing,
		"typography":  theme.Tokens.Semantic.Typography,
		"interactive": theme.Tokens.Semantic.Interactive,
	}

	for category, tokenMap := range categories {
		if tokenMap == nil {
			continue
		}
		for key, value := range tokenMap {
			ref := TokenReference(value)
			if ref.IsReference() {
				depth := v.calculateTokenDepth(value, theme.Tokens, 0, make(map[string]bool))
				if depth > maxAllowedDepth {
					issues = append(issues, ValidationIssue{
						Severity: "warning",
						Message:  fmt.Sprintf("token reference chain exceeds recommended depth of %d (found: %d)", maxAllowedDepth, depth),
						Path:     fmt.Sprintf("tokens.semantic.%s.%s", category, key),
					})
				}
			}
		}
	}

	return issues
}

// calculateTokenDepth calculates the depth of a token reference chain.
func (v *Validator) calculateTokenDepth(refStr string, tokens *Tokens, currentDepth int, visited map[string]bool) int {
	if currentDepth >= 20 {
		return currentDepth
	}

	if visited[refStr] {
		return currentDepth
	}
	visited[refStr] = true

	segments := strings.Split(refStr, ".")
	if len(segments) < 2 {
		return currentDepth
	}

	category := segments[0]
	var value string

	switch category {
	case "primitives":
		return currentDepth
	case "semantic":
		if tokens.Semantic == nil {
			return currentDepth
		}
		subcategory := segments[1]
		key := strings.Join(segments[2:], ".")
		
		var tokenMap map[string]string
		switch subcategory {
		case "colors":
			tokenMap = tokens.Semantic.Colors
		case "spacing":
			tokenMap = tokens.Semantic.Spacing
		case "typography":
			tokenMap = tokens.Semantic.Typography
		case "interactive":
			tokenMap = tokens.Semantic.Interactive
		}
		
		if tokenMap != nil {
			value = tokenMap[key]
		}
	}

	if value == "" {
		return currentDepth
	}

	ref := TokenReference(value)
	if !ref.IsReference() {
		return currentDepth + 1
	}

	return v.calculateTokenDepth(value, tokens, currentDepth+1, visited)
}

// validateNamingConventions checks that token names follow conventions.
func (v *Validator) validateNamingConventions(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Tokens == nil || theme.Tokens.Primitives == nil {
		return issues
	}

	kebabCase := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*(\.[a-z0-9]+(-[a-z0-9]+)*)*$`)

	primitives := theme.Tokens.Primitives
	categories := map[string]map[string]string{
		"colors":      primitives.Colors,
		"spacing":     primitives.Spacing,
		"radius":      primitives.Radius,
		"typography":  primitives.Typography,
		"borders":     primitives.Borders,
		"shadows":     primitives.Shadows,
		"effects":     primitives.Effects,
		"animation":   primitives.Animation,
		"zindex":      primitives.ZIndex,
		"breakpoints": primitives.Breakpoints,
	}

	for category, tokens := range categories {
		if tokens == nil {
			continue
		}
		for key := range tokens {
			if !kebabCase.MatchString(key) {
				issues = append(issues, ValidationIssue{
					Severity: "warning",
					Message:  "token name should use kebab-case (lowercase with hyphens)",
					Path:     fmt.Sprintf("tokens.primitives.%s.%s", category, key),
				})
			}
		}
	}

	return issues
}

// validateAccessibility checks accessibility configuration.
func (v *Validator) validateAccessibility(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Accessibility == nil {
		issues = append(issues, ValidationIssue{
			Severity: "info",
			Message:  "no accessibility configuration defined",
			Path:     "accessibility",
		})
		return issues
	}

	if !theme.Accessibility.FocusIndicator {
		issues = append(issues, ValidationIssue{
			Severity: "warning",
			Message:  "focus indicators should be enabled for keyboard navigation",
			Path:     "accessibility.focusIndicator",
		})
	}

	if !theme.Accessibility.KeyboardNav {
		issues = append(issues, ValidationIssue{
			Severity: "warning",
			Message:  "keyboard navigation should be enabled for accessibility",
			Path:     "accessibility.keyboardNav",
		})
	}

	return issues
}

// validatePerformance checks for performance issues.
func (v *Validator) validatePerformance(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if len(theme.CustomCSS) > 100*1024 {
		issues = append(issues, ValidationIssue{
			Severity: "warning",
			Message:  "custom CSS exceeds 100KB, consider splitting or optimizing",
			Path:     "customCSS",
		})
	}

	if len(theme.CustomJS) > 100*1024 {
		issues = append(issues, ValidationIssue{
			Severity: "warning",
			Message:  "custom JS exceeds 100KB, consider splitting or optimizing",
			Path:     "customJS",
		})
	}

	if theme.Tokens != nil && theme.Tokens.Primitives != nil {
		totalTokens := 0
		if theme.Tokens.Primitives.Colors != nil {
			totalTokens += len(theme.Tokens.Primitives.Colors)
		}
		if theme.Tokens.Primitives.Spacing != nil {
			totalTokens += len(theme.Tokens.Primitives.Spacing)
		}

		if totalTokens > 500 {
			issues = append(issues, ValidationIssue{
				Severity: "info",
				Message:  fmt.Sprintf("large number of primitive tokens (%d), consider consolidation", totalTokens),
				Path:     "tokens.primitives",
			})
		}
	}

	return issues
}

// validateBestPractices checks for best practice violations.
func (v *Validator) validateBestPractices(theme *Theme) []ValidationIssue {
	issues := make([]ValidationIssue, 0)

	if theme.Description == "" {
		issues = append(issues, ValidationIssue{
			Severity: "info",
			Message:  "theme should have a description",
			Path:     "description",
		})
	}

	if theme.Version == "" {
		issues = append(issues, ValidationIssue{
			Severity: "info",
			Message:  "theme should have a version",
			Path:     "version",
		})
	}

	if theme.Tokens != nil && theme.Tokens.Semantic == nil {
		issues = append(issues, ValidationIssue{
			Severity: "warning",
			Message:  "theme should define semantic tokens for better maintainability",
			Path:     "tokens.semantic",
		})
	}

	if theme.DarkMode == nil {
		issues = append(issues, ValidationIssue{
			Severity: "info",
			Message:  "consider adding dark mode support",
			Path:     "darkMode",
		})
	}

	return issues
}

// AddRule adds a custom validation rule.
func (v *Validator) AddRule(rule ValidationRule) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.rules = append(v.rules, rule)
}

// GetStats returns validator statistics.
func (v *Validator) GetStats() *ValidatorStats {
	v.mu.RLock()
	defer v.mu.RUnlock()

	stats := *v.stats
	return &stats
}

// ValidationResult contains the result of theme validation.
type ValidationResult struct {
	Valid  bool
	Issues []ValidationIssue
}

// ValidationIssue represents a single validation issue.
type ValidationIssue struct {
	Severity string
	Message  string
	Path     string
}

// ValidatorStats contains validator performance statistics.
type ValidatorStats struct {
	TotalValidations  uint64
	FailedValidations uint64
}
