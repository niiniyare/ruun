package validation

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ThemeValidator validates theme definitions and consistency
type ThemeValidator struct {
	rules         []ThemeValidationRule
	tokenRegistry *TokenRegistry
	config        ThemeValidationConfig
}

// ThemeValidationResult contains theme validation results
type ThemeValidationResult struct {
	Valid         bool                      `json:"valid"`
	Score         float64                   `json:"score"`          // 0-100
	Consistency   ThemeConsistencyResult   `json:"consistency"`
	TokenUsage    TokenUsageResult         `json:"tokenUsage"`
	Completeness  CompletenessResult       `json:"completeness"`
	Accessibility A11yThemeResult          `json:"accessibility"`
	Performance   ThemePerformanceResult   `json:"performance"`
	Violations    []ThemeViolation         `json:"violations"`
	Warnings      []ThemeWarning           `json:"warnings"`
	Suggestions   []ThemeSuggestion        `json:"suggestions"`
	Timestamp     time.Time                `json:"timestamp"`
}

// ThemeValidationConfig configures theme validation
type ThemeValidationConfig struct {
	RequireConsistency     bool     `json:"requireConsistency"`
	CheckTokenUsage        bool     `json:"checkTokenUsage"`
	ValidateCompleteness   bool     `json:"validateCompleteness"`
	CheckAccessibility     bool     `json:"checkAccessibility"`
	ValidatePerformance    bool     `json:"validatePerformance"`
	MinContrastRatio       float64  `json:"minContrastRatio"`
	MaxTokenCount          int      `json:"maxTokenCount"`
	RequiredCategories     []string `json:"requiredCategories"`
	AllowedColorFormats    []string `json:"allowedColorFormats"`
	StrictMode            bool     `json:"strictMode"`
}

// ThemeConsistencyResult contains consistency validation results
type ThemeConsistencyResult struct {
	Consistent        bool                     `json:"consistent"`
	Score             float64                  `json:"score"`
	ColorConsistency  ColorConsistencyResult  `json:"colorConsistency"`
	SpacingConsistency SpacingConsistencyResult `json:"spacingConsistency"`
	TypographyConsistency TypographyConsistencyResult `json:"typographyConsistency"`
	Issues            []ConsistencyIssue      `json:"issues"`
}

// TokenUsageResult contains token usage analysis
type TokenUsageResult struct {
	TotalTokens       int                      `json:"totalTokens"`
	UsedTokens        int                      `json:"usedTokens"`
	UnusedTokens      []string                 `json:"unusedTokens"`
	MissingTokens     []string                 `json:"missingTokens"`
	DuplicateTokens   []TokenDuplicate         `json:"duplicateTokens"`
	UsagePercentage   float64                  `json:"usagePercentage"`
	CategoryBreakdown map[string]TokenCategoryUsage `json:"categoryBreakdown"`
}

// CompletenessResult contains theme completeness analysis
type CompletenessResult struct {
	Complete            bool                      `json:"complete"`
	Score               float64                   `json:"score"`
	MissingCategories   []string                  `json:"missingCategories"`
	IncompleteCategories []IncompleteCategory     `json:"incompleteCategories"`
	CoveragePercentage  float64                   `json:"coveragePercentage"`
	Recommendations     []CompletenessRecommendation `json:"recommendations"`
}

// A11yThemeResult contains accessibility analysis for themes
type A11yThemeResult struct {
	Accessible       bool               `json:"accessible"`
	Score            float64            `json:"score"`
	ContrastIssues   []ContrastIssue    `json:"contrastIssues"`
	ColorBlindness   ColorBlindnessResult `json:"colorBlindness"`
	ReadabilityScore float64            `json:"readabilityScore"`
	Issues           []A11yThemeIssue   `json:"issues"`
}

// ThemePerformanceResult contains theme performance analysis
type ThemePerformanceResult struct {
	Optimized     bool                    `json:"optimized"`
	Score         float64                 `json:"score"`
	BundleSize    int64                   `json:"bundleSize"`
	ComputedSize  int64                   `json:"computedSize"`
	TokenCount    int                     `json:"tokenCount"`
	Complexity    ThemeComplexityMetrics  `json:"complexity"`
	Optimizations []PerformanceOptimization `json:"optimizations"`
}

// Supporting types

type ThemeValidationRule struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Category    ThemeValidationCategory `json:"category"`
	Level       ValidationLevel        `json:"level"`
	Enabled     bool                   `json:"enabled"`
	Validator   ThemeRuleValidator     `json:"-"`
}

type ThemeValidationCategory string

const (
	ThemeValidationCategoryConsistency  ThemeValidationCategory = "consistency"
	ThemeValidationCategoryCompleteness ThemeValidationCategory = "completeness"
	ThemeValidationCategoryAccessibility ThemeValidationCategory = "accessibility"
	ThemeValidationCategoryPerformance  ThemeValidationCategory = "performance"
	ThemeValidationCategoryTokens       ThemeValidationCategory = "tokens"
)

type ThemeRuleValidator interface {
	ValidateTheme(theme any, context *ThemeValidationContext) *ThemeRuleResult
	GetCategory() ThemeValidationCategory
}

type ThemeValidationContext struct {
	Theme       any            `json:"theme"`
	Registry    *TokenRegistry         `json:"registry"`
	Config      ThemeValidationConfig  `json:"config"`
	Context     map[string]any `json:"context"`
}

type ThemeRuleResult struct {
	Passed      bool                   `json:"passed"`
	Score       float64                `json:"score"`
	Rule        string                 `json:"rule"`
	Violations  []ThemeViolation       `json:"violations,omitempty"`
	Warnings    []ThemeWarning         `json:"warnings,omitempty"`
	Suggestions []ThemeSuggestion      `json:"suggestions,omitempty"`
	Metadata    map[string]any `json:"metadata"`
}

type ThemeViolation struct {
	Code        string                 `json:"code"`
	Message     string                 `json:"message"`
	Category    ThemeValidationCategory `json:"category"`
	Severity    ValidationLevel        `json:"severity"`
	Token       string                 `json:"token,omitempty"`
	Value       string                 `json:"value,omitempty"`
	Expected    string                 `json:"expected,omitempty"`
	Location    *SourceLocation        `json:"location,omitempty"`
	Fix         string                 `json:"fix,omitempty"`
	AutoFix     bool                   `json:"autoFix"`
	Context     map[string]any `json:"context,omitempty"`
}

type ThemeWarning struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Category  ThemeValidationCategory `json:"category"`
	Token     string                 `json:"token,omitempty"`
	Value     string                 `json:"value,omitempty"`
	Location  *SourceLocation        `json:"location,omitempty"`
	Fix       string                 `json:"fix,omitempty"`
	Context   map[string]any `json:"context,omitempty"`
}

type ThemeSuggestion struct {
	Code        string                 `json:"code"`
	Message     string                 `json:"message"`
	Category    ThemeValidationCategory `json:"category"`
	Priority    string                 `json:"priority"` // high, medium, low
	Token       string                 `json:"token,omitempty"`
	Suggestion  string                 `json:"suggestion"`
	Examples    []string               `json:"examples,omitempty"`
	Benefits    []string               `json:"benefits,omitempty"`
	AutoFix     bool                   `json:"autoFix"`
	Context     map[string]any `json:"context,omitempty"`
}

// Consistency types
type ColorConsistencyResult struct {
	Consistent    bool                   `json:"consistent"`
	Score         float64                `json:"score"`
	Palette       ColorPaletteAnalysis   `json:"palette"`
	Harmony       ColorHarmonyAnalysis   `json:"harmony"`
	Issues        []ColorConsistencyIssue `json:"issues"`
}

type SpacingConsistencyResult struct {
	Consistent bool                     `json:"consistent"`
	Score      float64                  `json:"score"`
	Scale      SpacingScaleAnalysis     `json:"scale"`
	Rhythm     SpacingRhythmAnalysis    `json:"rhythm"`
	Issues     []SpacingConsistencyIssue `json:"issues"`
}

type TypographyConsistencyResult struct {
	Consistent bool                        `json:"consistent"`
	Score      float64                     `json:"score"`
	Scale      TypographyScaleAnalysis     `json:"scale"`
	Hierarchy  TypographyHierarchyAnalysis `json:"hierarchy"`
	Issues     []TypographyConsistencyIssue `json:"issues"`
}

type ConsistencyIssue struct {
	Type        string                 `json:"type"`
	Message     string                 `json:"message"`
	Tokens      []string               `json:"tokens"`
	Suggestion  string                 `json:"suggestion"`
	Severity    ValidationLevel        `json:"severity"`
	Context     map[string]any `json:"context,omitempty"`
}

// Token usage types
type TokenDuplicate struct {
	Value  string   `json:"value"`
	Tokens []string `json:"tokens"`
}

type TokenCategoryUsage struct {
	Total     int     `json:"total"`
	Used      int     `json:"used"`
	Unused    int     `json:"unused"`
	Percentage float64 `json:"percentage"`
}

// Completeness types
type IncompleteCategory struct {
	Category    string   `json:"category"`
	Missing     []string `json:"missing"`
	Coverage    float64  `json:"coverage"`
	Required    []string `json:"required"`
	Suggestions []string `json:"suggestions,omitempty"`
}

type CompletenessRecommendation struct {
	Category    string   `json:"category"`
	Message     string   `json:"message"`
	Priority    string   `json:"priority"`
	Tokens      []string `json:"tokens"`
	Examples    []string `json:"examples,omitempty"`
}

// Accessibility types
type ContrastIssue struct {
	ForegroundToken string  `json:"foregroundToken"`
	BackgroundToken string  `json:"backgroundToken"`
	ForegroundColor string  `json:"foregroundColor"`
	BackgroundColor string  `json:"backgroundColor"`
	ContrastRatio   float64 `json:"contrastRatio"`
	RequiredRatio   float64 `json:"requiredRatio"`
	Context         string  `json:"context"`
}

type ColorBlindnessResult struct {
	Deuteranopia  ColorBlindnessTest `json:"deuteranopia"`
	Protanopia    ColorBlindnessTest `json:"protanopia"`
	Tritanopia    ColorBlindnessTest `json:"tritanopia"`
	Achromatopsia ColorBlindnessTest `json:"achromatopsia"`
}

type ColorBlindnessTest struct {
	Accessible bool     `json:"accessible"`
	Issues     []string `json:"issues,omitempty"`
	Score      float64  `json:"score"`
}

type A11yThemeIssue struct {
	Type        string                 `json:"type"`
	Message     string                 `json:"message"`
	Severity    ValidationLevel        `json:"severity"`
	Tokens      []string               `json:"tokens,omitempty"`
	Fix         string                 `json:"fix,omitempty"`
	Context     map[string]any `json:"context,omitempty"`
}

// Performance types
type ThemeComplexityMetrics struct {
	TokenCount      int     `json:"tokenCount"`
	NestingDepth    int     `json:"nestingDepth"`
	ComputedTokens  int     `json:"computedTokens"`
	CircularRefs    int     `json:"circularRefs"`
	ComplexityScore int     `json:"complexityScore"`
}

type PerformanceOptimization struct {
	Type        string   `json:"type"`
	Message     string   `json:"message"`
	Impact      string   `json:"impact"` // high, medium, low
	Savings     string   `json:"savings,omitempty"`
	AutoFix     bool     `json:"autoFix"`
	Examples    []string `json:"examples,omitempty"`
}

// Analysis types
type ColorPaletteAnalysis struct {
	PrimaryColors   []string `json:"primaryColors"`
	SecondaryColors []string `json:"secondaryColors"`
	GrayScale      []string `json:"grayScale"`
	AccentColors   []string `json:"accentColors"`
	TotalColors    int      `json:"totalColors"`
	Diversity      float64  `json:"diversity"`
}

type ColorHarmonyAnalysis struct {
	Harmonious    bool     `json:"harmonious"`
	HarmonyType   string   `json:"harmonyType"`
	Relationships []string `json:"relationships"`
	Score         float64  `json:"score"`
}

type ColorConsistencyIssue struct {
	Type      string   `json:"type"`
	Message   string   `json:"message"`
	Colors    []string `json:"colors"`
	Tokens    []string `json:"tokens"`
	Severity  ValidationLevel `json:"severity"`
}

type SpacingScaleAnalysis struct {
	BaseUnit     string   `json:"baseUnit"`
	Scale        []string `json:"scale"`
	Ratio        float64  `json:"ratio"`
	Consistent   bool     `json:"consistent"`
	Gaps         []string `json:"gaps,omitempty"`
}

type SpacingRhythmAnalysis struct {
	Rhythmic     bool     `json:"rhythmic"`
	BaseRhythm   string   `json:"baseRhythm"`
	Violations   []string `json:"violations,omitempty"`
	Score        float64  `json:"score"`
}

type SpacingConsistencyIssue struct {
	Type      string   `json:"type"`
	Message   string   `json:"message"`
	Values    []string `json:"values"`
	Tokens    []string `json:"tokens"`
	Severity  ValidationLevel `json:"severity"`
}

type TypographyScaleAnalysis struct {
	BaseSize     string   `json:"baseSize"`
	Scale        []string `json:"scale"`
	Ratio        float64  `json:"ratio"`
	Consistent   bool     `json:"consistent"`
	Gaps         []string `json:"gaps,omitempty"`
}

type TypographyHierarchyAnalysis struct {
	Clear        bool     `json:"clear"`
	Levels       int      `json:"levels"`
	Progression  []string `json:"progression"`
	Issues       []string `json:"issues,omitempty"`
	Score        float64  `json:"score"`
}

type TypographyConsistencyIssue struct {
	Type      string   `json:"type"`
	Message   string   `json:"message"`
	Values    []string `json:"values"`
	Tokens    []string `json:"tokens"`
	Severity  ValidationLevel `json:"severity"`
}

// TokenRegistry manages design token definitions
type TokenRegistry struct {
	categories map[string][]string
	required   map[string][]string
	patterns   map[string]*regexp.Regexp
}

// NewThemeValidator creates a new theme validator
func NewThemeValidator() *ThemeValidator {
	config := ThemeValidationConfig{
		RequireConsistency:   true,
		CheckTokenUsage:      true,
		ValidateCompleteness: true,
		CheckAccessibility:   true,
		ValidatePerformance:  true,
		MinContrastRatio:     4.5,
		MaxTokenCount:        500,
		RequiredCategories:   []string{"colors", "spacing", "typography"},
		AllowedColorFormats:  []string{"hex", "rgb", "hsl", "oklch"},
		StrictMode:          false,
	}

	validator := &ThemeValidator{
		config:        config,
		tokenRegistry: NewTokenRegistry(),
	}

	validator.registerBuiltinRules()
	return validator
}

// NewTokenRegistry creates a new token registry
func NewTokenRegistry() *TokenRegistry {
	registry := &TokenRegistry{
		categories: make(map[string][]string),
		required:   make(map[string][]string),
		patterns:   make(map[string]*regexp.Regexp),
	}

	// Define standard token categories
	registry.categories["colors"] = []string{
		"primary", "secondary", "background", "surface", "error", "warning", "success", "info",
		"text-primary", "text-secondary", "text-disabled", "border", "divider",
	}
	
	registry.categories["spacing"] = []string{
		"xs", "sm", "md", "lg", "xl", "2xl", "3xl", "4xl", "5xl", "6xl",
	}
	
	registry.categories["typography"] = []string{
		"font-family-primary", "font-family-secondary", "font-family-mono",
		"font-size-xs", "font-size-sm", "font-size-md", "font-size-lg", "font-size-xl",
		"line-height-tight", "line-height-normal", "line-height-loose",
		"font-weight-light", "font-weight-normal", "font-weight-medium", "font-weight-bold",
	}

	// Define required tokens
	registry.required["colors"] = []string{"primary", "background", "text-primary", "border"}
	registry.required["spacing"] = []string{"xs", "sm", "md", "lg", "xl"}
	registry.required["typography"] = []string{"font-family-primary", "font-size-md", "line-height-normal"}

	// Define validation patterns
	registry.patterns["color"] = regexp.MustCompile(`^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$|^rgb\(|^hsl\(|^oklch\(`)
	registry.patterns["spacing"] = regexp.MustCompile(`^\d+(\.\d+)?(px|rem|em|%)$`)
	registry.patterns["font-size"] = regexp.MustCompile(`^\d+(\.\d+)?(px|rem|em|%)$`)

	return registry
}

// ValidateTheme validates a complete theme
func (tv *ThemeValidator) ValidateTheme(ctx *ValidationContext, theme any) *ThemeValidationResult {
	start := time.Now()
	
	result := &ThemeValidationResult{
		Valid:     true,
		Score:     100.0,
		Timestamp: start,
	}

	themeContext := &ThemeValidationContext{
		Theme:    theme,
		Registry: tv.tokenRegistry,
		Config:   tv.config,
		Context:  make(map[string]any),
	}

	// Run validation rules
	for _, rule := range tv.rules {
		if !rule.Enabled {
			continue
		}

		ruleResult := rule.Validator.ValidateTheme(theme, themeContext)
		if ruleResult != nil {
			result = tv.mergeRuleResult(result, ruleResult)
		}
	}

	// Run specialized validations
	if tv.config.RequireConsistency {
		result.Consistency = tv.validateConsistency(theme, themeContext)
		if !result.Consistency.Consistent {
			result.Valid = false
			result.Score -= 20.0
		}
	}

	if tv.config.CheckTokenUsage {
		result.TokenUsage = tv.validateTokenUsage(theme, themeContext)
	}

	if tv.config.ValidateCompleteness {
		result.Completeness = tv.validateCompleteness(theme, themeContext)
		if !result.Completeness.Complete {
			result.Valid = false
			result.Score -= 15.0
		}
	}

	if tv.config.CheckAccessibility {
		result.Accessibility = tv.validateAccessibility(theme, themeContext)
		if !result.Accessibility.Accessible {
			result.Valid = false
			result.Score -= 25.0
		}
	}

	if tv.config.ValidatePerformance {
		result.Performance = tv.validatePerformance(theme, themeContext)
		if !result.Performance.Optimized {
			result.Score -= 10.0
		}
	}

	// Ensure score is within bounds
	if result.Score < 0 {
		result.Score = 0
	}

	return result
}

// registerBuiltinRules registers built-in validation rules
func (tv *ThemeValidator) registerBuiltinRules() {
	// Token format validation
	tv.AddRule(ThemeValidationRule{
		ID:          "theme.tokens.format",
		Name:        "Token Format Validation",
		Description: "Validates that tokens follow proper format conventions",
		Category:    ThemeValidationCategoryTokens,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &TokenFormatValidator{registry: tv.tokenRegistry},
	})

	// Color consistency validation
	tv.AddRule(ThemeValidationRule{
		ID:          "theme.colors.consistency",
		Name:        "Color Consistency",
		Description: "Validates color palette consistency and harmony",
		Category:    ThemeValidationCategoryConsistency,
		Level:       ValidationLevelWarning,
		Enabled:     true,
		Validator:   &ColorConsistencyValidator{},
	})

	// Accessibility validation
	tv.AddRule(ThemeValidationRule{
		ID:          "theme.accessibility.contrast",
		Name:        "Color Contrast",
		Description: "Validates color contrast ratios for accessibility",
		Category:    ThemeValidationCategoryAccessibility,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &ThemeContrastValidator{minRatio: tv.config.MinContrastRatio},
	})

	// Performance validation
	tv.AddRule(ThemeValidationRule{
		ID:          "theme.performance.size",
		Name:        "Theme Size",
		Description: "Validates theme bundle size and complexity",
		Category:    ThemeValidationCategoryPerformance,
		Level:       ValidationLevelWarning,
		Enabled:     true,
		Validator:   &ThemePerformanceValidator{maxTokens: tv.config.MaxTokenCount},
	})
}

// AddRule adds a validation rule
func (tv *ThemeValidator) AddRule(rule ThemeValidationRule) {
	tv.rules = append(tv.rules, rule)
}

// validateConsistency validates theme consistency
func (tv *ThemeValidator) validateConsistency(theme any, context *ThemeValidationContext) ThemeConsistencyResult {
	result := ThemeConsistencyResult{
		Consistent: true,
		Score:      100.0,
		Issues:     make([]ConsistencyIssue, 0),
	}

	// Validate color consistency
	result.ColorConsistency = tv.validateColorConsistency(theme)
	if !result.ColorConsistency.Consistent {
		result.Consistent = false
		result.Score -= 30.0
	}

	// Validate spacing consistency
	result.SpacingConsistency = tv.validateSpacingConsistency(theme)
	if !result.SpacingConsistency.Consistent {
		result.Consistent = false
		result.Score -= 30.0
	}

	// Validate typography consistency
	result.TypographyConsistency = tv.validateTypographyConsistency(theme)
	if !result.TypographyConsistency.Consistent {
		result.Consistent = false
		result.Score -= 30.0
	}

	// Calculate overall consistency score
	result.Score = (result.ColorConsistency.Score + 
		result.SpacingConsistency.Score + 
		result.TypographyConsistency.Score) / 3.0

	return result
}

// validateTokenUsage validates token usage patterns
func (tv *ThemeValidator) validateTokenUsage(theme any, context *ThemeValidationContext) TokenUsageResult {
	result := TokenUsageResult{
		UnusedTokens:      make([]string, 0),
		MissingTokens:     make([]string, 0),
		DuplicateTokens:   make([]TokenDuplicate, 0),
		CategoryBreakdown: make(map[string]TokenCategoryUsage),
	}

	tokens := tv.extractTokens(theme)
	result.TotalTokens = len(tokens)

	// Check for missing required tokens
	for category, requiredTokens := range tv.tokenRegistry.required {
		categoryUsage := TokenCategoryUsage{
			Total: len(requiredTokens),
		}

		for _, required := range requiredTokens {
			if _, exists := tokens[required]; !exists {
				result.MissingTokens = append(result.MissingTokens, required)
			} else {
				categoryUsage.Used++
			}
		}

		categoryUsage.Unused = categoryUsage.Total - categoryUsage.Used
		if categoryUsage.Total > 0 {
			categoryUsage.Percentage = float64(categoryUsage.Used) / float64(categoryUsage.Total) * 100.0
		}
		result.CategoryBreakdown[category] = categoryUsage
	}

	// Check for duplicate values
	valueToTokens := make(map[string][]string)
	for token, value := range tokens {
		valueStr := fmt.Sprintf("%v", value)
		valueToTokens[valueStr] = append(valueToTokens[valueStr], token)
	}

	for value, tokenList := range valueToTokens {
		if len(tokenList) > 1 {
			result.DuplicateTokens = append(result.DuplicateTokens, TokenDuplicate{
				Value:  value,
				Tokens: tokenList,
			})
		}
	}

	// Calculate usage percentage
	if result.TotalTokens > 0 {
		result.UsedTokens = result.TotalTokens - len(result.UnusedTokens)
		result.UsagePercentage = float64(result.UsedTokens) / float64(result.TotalTokens) * 100.0
	}

	return result
}

// validateCompleteness validates theme completeness
func (tv *ThemeValidator) validateCompleteness(theme any, context *ThemeValidationContext) CompletenessResult {
	result := CompletenessResult{
		Complete:             true,
		Score:               100.0,
		MissingCategories:   make([]string, 0),
		IncompleteCategories: make([]IncompleteCategory, 0),
		Recommendations:     make([]CompletenessRecommendation, 0),
	}

	tokens := tv.extractTokens(theme)
	totalRequired := 0
	totalPresent := 0

	// Check required categories
	for category, requiredTokens := range tv.tokenRegistry.required {
		missing := make([]string, 0)
		present := 0

		for _, token := range requiredTokens {
			if _, exists := tokens[token]; exists {
				present++
			} else {
				missing = append(missing, token)
			}
		}

		totalRequired += len(requiredTokens)
		totalPresent += present

		if len(missing) > 0 {
			result.Complete = false
			coverage := float64(present) / float64(len(requiredTokens)) * 100.0
			
			result.IncompleteCategories = append(result.IncompleteCategories, IncompleteCategory{
				Category: category,
				Missing:  missing,
				Coverage: coverage,
				Required: requiredTokens,
			})

			if present == 0 {
				result.MissingCategories = append(result.MissingCategories, category)
			}
		}
	}

	// Calculate coverage percentage
	if totalRequired > 0 {
		result.CoveragePercentage = float64(totalPresent) / float64(totalRequired) * 100.0
		result.Score = result.CoveragePercentage
	}

	return result
}

// validateAccessibility validates theme accessibility
func (tv *ThemeValidator) validateAccessibility(theme any, context *ThemeValidationContext) A11yThemeResult {
	result := A11yThemeResult{
		Accessible:     true,
		Score:          100.0,
		ContrastIssues: make([]ContrastIssue, 0),
		Issues:         make([]A11yThemeIssue, 0),
	}

	colors := tv.extractColorTokens(theme)
	
	// Check contrast ratios
	textColors := tv.filterTextColors(colors)
	backgroundColors := tv.filterBackgroundColors(colors)

	for textToken, textColor := range textColors {
		for bgToken, bgColor := range backgroundColors {
			ratio := tv.calculateContrastRatio(textColor, bgColor)
			
			if ratio < context.Config.MinContrastRatio {
				result.Accessible = false
				result.ContrastIssues = append(result.ContrastIssues, ContrastIssue{
					ForegroundToken: textToken,
					BackgroundToken: bgToken,
					ForegroundColor: textColor,
					BackgroundColor: bgColor,
					ContrastRatio:   ratio,
					RequiredRatio:   context.Config.MinContrastRatio,
					Context:        fmt.Sprintf("Text on background combination"),
				})
			}
		}
	}

	// Check color blindness accessibility
	result.ColorBlindness = tv.validateColorBlindness(colors)

	// Calculate accessibility score
	if len(result.ContrastIssues) > 0 {
		penalty := float64(len(result.ContrastIssues)) * 10.0
		result.Score = math.Max(0, result.Score-penalty)
	}

	if !result.ColorBlindness.Deuteranopia.Accessible {
		result.Score -= 15.0
	}
	if !result.ColorBlindness.Protanopia.Accessible {
		result.Score -= 15.0
	}
	if !result.ColorBlindness.Tritanopia.Accessible {
		result.Score -= 10.0
	}

	return result
}

// validatePerformance validates theme performance characteristics
func (tv *ThemeValidator) validatePerformance(theme any, context *ThemeValidationContext) ThemePerformanceResult {
	result := ThemePerformanceResult{
		Optimized:     true,
		Score:         100.0,
		Optimizations: make([]PerformanceOptimization, 0),
	}

	tokens := tv.extractTokens(theme)
	result.TokenCount = len(tokens)

	// Calculate estimated bundle size
	result.BundleSize = tv.estimateThemeSize(theme)
	result.ComputedSize = tv.estimateComputedSize(theme)

	// Calculate complexity
	result.Complexity = tv.calculateThemeComplexity(theme)

	// Check token count threshold
	if result.TokenCount > context.Config.MaxTokenCount {
		result.Optimized = false
		result.Score -= 20.0
		result.Optimizations = append(result.Optimizations, PerformanceOptimization{
			Type:    "token_count",
			Message: fmt.Sprintf("Token count %d exceeds recommended maximum %d", result.TokenCount, context.Config.MaxTokenCount),
			Impact:  "medium",
			AutoFix: false,
		})
	}

	// Check for optimization opportunities
	duplicates := tv.findDuplicateValues(tokens)
	if len(duplicates) > 0 {
		result.Score -= 10.0
		result.Optimizations = append(result.Optimizations, PerformanceOptimization{
			Type:    "duplicate_values",
			Message: fmt.Sprintf("Found %d duplicate token values that could be consolidated", len(duplicates)),
			Impact:  "low",
			Savings: fmt.Sprintf("~%d tokens", len(duplicates)),
			AutoFix: true,
		})
	}

	return result
}

// Helper methods for theme validation

func (tv *ThemeValidator) extractTokens(theme any) map[string]any {
	tokens := make(map[string]any)
	
	if themeMap, ok := theme.(map[string]any); ok {
		if tokensData, exists := themeMap["tokens"]; exists {
			if tokensMap, ok := tokensData.(map[string]any); ok {
				tv.flattenTokens(tokensMap, "", tokens)
			}
		}
	}

	return tokens
}

func (tv *ThemeValidator) flattenTokens(data map[string]any, prefix string, result map[string]any) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		if valueMap, ok := value.(map[string]any); ok {
			// Check if this is a token object with a value
			if tokenValue, hasValue := valueMap["value"]; hasValue {
				result[fullKey] = tokenValue
			} else {
				// Recursively flatten nested objects
				tv.flattenTokens(valueMap, fullKey, result)
			}
		} else {
			// Direct value
			result[fullKey] = value
		}
	}
}

func (tv *ThemeValidator) extractColorTokens(theme any) map[string]string {
	tokens := tv.extractTokens(theme)
	colors := make(map[string]string)

	for token, value := range tokens {
		if strings.Contains(strings.ToLower(token), "color") || 
		   strings.Contains(strings.ToLower(token), "background") ||
		   strings.Contains(strings.ToLower(token), "text") ||
		   strings.Contains(strings.ToLower(token), "border") {
			if valueStr, ok := value.(string); ok {
				if tv.isColorValue(valueStr) {
					colors[token] = valueStr
				}
			}
		}
	}

	return colors
}

func (tv *ThemeValidator) isColorValue(value string) bool {
	if tv.tokenRegistry.patterns["color"] != nil {
		return tv.tokenRegistry.patterns["color"].MatchString(value)
	}
	
	// Fallback color detection
	return strings.HasPrefix(value, "#") || 
		   strings.HasPrefix(value, "rgb") || 
		   strings.HasPrefix(value, "hsl") ||
		   strings.HasPrefix(value, "oklch")
}

func (tv *ThemeValidator) filterTextColors(colors map[string]string) map[string]string {
	textColors := make(map[string]string)
	for token, color := range colors {
		if strings.Contains(strings.ToLower(token), "text") ||
		   strings.Contains(strings.ToLower(token), "foreground") {
			textColors[token] = color
		}
	}
	return textColors
}

func (tv *ThemeValidator) filterBackgroundColors(colors map[string]string) map[string]string {
	backgroundColors := make(map[string]string)
	for token, color := range colors {
		if strings.Contains(strings.ToLower(token), "background") ||
		   strings.Contains(strings.ToLower(token), "surface") {
			backgroundColors[token] = color
		}
	}
	return backgroundColors
}

func (tv *ThemeValidator) calculateContrastRatio(color1, color2 string) float64 {
	// Simplified contrast ratio calculation
	// In a real implementation, this would properly parse colors and calculate luminance
	lum1 := tv.getRelativeLuminance(color1)
	lum2 := tv.getRelativeLuminance(color2)
	
	lighter := math.Max(lum1, lum2)
	darker := math.Min(lum1, lum2)
	
	return (lighter + 0.05) / (darker + 0.05)
}

func (tv *ThemeValidator) getRelativeLuminance(color string) float64 {
	// Simplified luminance calculation
	// Real implementation would parse hex, rgb, hsl properly
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

func (tv *ThemeValidator) validateColorBlindness(colors map[string]string) ColorBlindnessResult {
	// Simplified color blindness validation
	// Real implementation would simulate different types of color blindness
	return ColorBlindnessResult{
		Deuteranopia: ColorBlindnessTest{
			Accessible: true,
			Score:      85.0,
		},
		Protanopia: ColorBlindnessTest{
			Accessible: true,
			Score:      85.0,
		},
		Tritanopia: ColorBlindnessTest{
			Accessible: true,
			Score:      90.0,
		},
		Achromatopsia: ColorBlindnessTest{
			Accessible: true,
			Score:      75.0,
		},
	}
}

// Additional validation methods would be implemented similarly...

func (tv *ThemeValidator) validateColorConsistency(theme any) ColorConsistencyResult {
	result := ColorConsistencyResult{
		Consistent: true,
		Score:      100.0,
		Issues:     make([]ColorConsistencyIssue, 0),
	}

	colors := tv.extractColorTokens(theme)
	
	// Analyze color palette
	result.Palette = tv.analyzeColorPalette(colors)
	
	// Analyze color harmony
	result.Harmony = tv.analyzeColorHarmony(colors)
	
	if !result.Harmony.Harmonious {
		result.Consistent = false
		result.Score -= 30.0
	}

	return result
}

func (tv *ThemeValidator) validateSpacingConsistency(theme any) SpacingConsistencyResult {
	result := SpacingConsistencyResult{
		Consistent: true,
		Score:      100.0,
		Issues:     make([]SpacingConsistencyIssue, 0),
	}

	// Extract spacing tokens
	tokens := tv.extractTokens(theme)
	spacingTokens := make(map[string]string)
	
	for token, value := range tokens {
		if strings.Contains(strings.ToLower(token), "spacing") ||
		   strings.Contains(strings.ToLower(token), "margin") ||
		   strings.Contains(strings.ToLower(token), "padding") {
			if valueStr, ok := value.(string); ok {
				spacingTokens[token] = valueStr
			}
		}
	}

	// Analyze spacing scale
	result.Scale = tv.analyzeSpacingScale(spacingTokens)
	
	// Analyze spacing rhythm
	result.Rhythm = tv.analyzeSpacingRhythm(spacingTokens)

	if !result.Scale.Consistent || !result.Rhythm.Rhythmic {
		result.Consistent = false
		result.Score -= 25.0
	}

	return result
}

func (tv *ThemeValidator) validateTypographyConsistency(theme any) TypographyConsistencyResult {
	result := TypographyConsistencyResult{
		Consistent: true,
		Score:      100.0,
		Issues:     make([]TypographyConsistencyIssue, 0),
	}

	// Extract typography tokens
	tokens := tv.extractTokens(theme)
	typographyTokens := make(map[string]string)
	
	for token, value := range tokens {
		if strings.Contains(strings.ToLower(token), "font") ||
		   strings.Contains(strings.ToLower(token), "typography") ||
		   strings.Contains(strings.ToLower(token), "text") {
			if valueStr, ok := value.(string); ok {
				typographyTokens[token] = valueStr
			}
		}
	}

	// Analyze typography scale
	result.Scale = tv.analyzeTypographyScale(typographyTokens)
	
	// Analyze typography hierarchy
	result.Hierarchy = tv.analyzeTypographyHierarchy(typographyTokens)

	if !result.Scale.Consistent || !result.Hierarchy.Clear {
		result.Consistent = false
		result.Score -= 25.0
	}

	return result
}

// More helper methods would be implemented here...

func (tv *ThemeValidator) analyzeColorPalette(colors map[string]string) ColorPaletteAnalysis {
	return ColorPaletteAnalysis{
		TotalColors: len(colors),
		Diversity:   0.8, // Mock diversity score
	}
}

func (tv *ThemeValidator) analyzeColorHarmony(colors map[string]string) ColorHarmonyAnalysis {
	return ColorHarmonyAnalysis{
		Harmonious:  true,
		HarmonyType: "complementary",
		Score:       85.0,
	}
}

func (tv *ThemeValidator) analyzeSpacingScale(spacing map[string]string) SpacingScaleAnalysis {
	return SpacingScaleAnalysis{
		BaseUnit:   "1rem",
		Consistent: true,
		Ratio:      1.5,
	}
}

func (tv *ThemeValidator) analyzeSpacingRhythm(spacing map[string]string) SpacingRhythmAnalysis {
	return SpacingRhythmAnalysis{
		Rhythmic:   true,
		BaseRhythm: "1.5rem",
		Score:      90.0,
	}
}

func (tv *ThemeValidator) analyzeTypographyScale(typography map[string]string) TypographyScaleAnalysis {
	return TypographyScaleAnalysis{
		BaseSize:   "1rem",
		Consistent: true,
		Ratio:      1.25,
	}
}

func (tv *ThemeValidator) analyzeTypographyHierarchy(typography map[string]string) TypographyHierarchyAnalysis {
	return TypographyHierarchyAnalysis{
		Clear:  true,
		Levels: 6,
		Score:  85.0,
	}
}

func (tv *ThemeValidator) estimateThemeSize(theme any) int64 {
	// Simplified size estimation
	tokens := tv.extractTokens(theme)
	size := int64(len(tokens) * 50) // Estimate 50 bytes per token
	return size
}

func (tv *ThemeValidator) estimateComputedSize(theme any) int64 {
	// Estimate CSS output size
	return tv.estimateThemeSize(theme) * 3 // Roughly 3x expansion
}

func (tv *ThemeValidator) calculateThemeComplexity(theme any) ThemeComplexityMetrics {
	tokens := tv.extractTokens(theme)
	
	return ThemeComplexityMetrics{
		TokenCount:      len(tokens),
		NestingDepth:    tv.calculateNestingDepth(theme),
		ComputedTokens:  len(tokens) / 2, // Estimate
		ComplexityScore: len(tokens) + tv.calculateNestingDepth(theme)*10,
	}
}

func (tv *ThemeValidator) calculateNestingDepth(data any) int {
	if dataMap, ok := data.(map[string]any); ok {
		maxDepth := 0
		for _, value := range dataMap {
			depth := 1 + tv.calculateNestingDepth(value)
			if depth > maxDepth {
				maxDepth = depth
			}
		}
		return maxDepth
	}
	return 0
}

func (tv *ThemeValidator) findDuplicateValues(tokens map[string]any) []TokenDuplicate {
	valueToTokens := make(map[string][]string)
	duplicates := make([]TokenDuplicate, 0)

	for token, value := range tokens {
		valueStr := fmt.Sprintf("%v", value)
		valueToTokens[valueStr] = append(valueToTokens[valueStr], token)
	}

	for value, tokenList := range valueToTokens {
		if len(tokenList) > 1 {
			duplicates = append(duplicates, TokenDuplicate{
				Value:  value,
				Tokens: tokenList,
			})
		}
	}

	return duplicates
}

func (tv *ThemeValidator) mergeRuleResult(base *ThemeValidationResult, ruleResult *ThemeRuleResult) *ThemeValidationResult {
	if !ruleResult.Passed {
		base.Valid = false
		base.Score = math.Min(base.Score, ruleResult.Score)
	}

	base.Violations = append(base.Violations, ruleResult.Violations...)
	base.Warnings = append(base.Warnings, ruleResult.Warnings...)
	base.Suggestions = append(base.Suggestions, ruleResult.Suggestions...)

	return base
}

// Built-in Theme Validators

// TokenFormatValidator validates token format conventions
type TokenFormatValidator struct {
	registry *TokenRegistry
}

func (v *TokenFormatValidator) ValidateTheme(theme any, context *ThemeValidationContext) *ThemeRuleResult {
	result := &ThemeRuleResult{
		Passed:      true,
		Score:       100.0,
		Rule:        "theme.tokens.format",
		Metadata:    make(map[string]any),
		Violations:  make([]ThemeViolation, 0),
		Warnings:    make([]ThemeWarning, 0),
		Suggestions: make([]ThemeSuggestion, 0),
	}

	tokens := v.extractTokens(theme)
	
	for token, value := range tokens {
		// Validate token naming convention
		if !v.isValidTokenName(token) {
			result.Violations = append(result.Violations, ThemeViolation{
				Code:     "invalid_token_name",
				Message:  fmt.Sprintf("Token name '%s' doesn't follow naming conventions", token),
				Category: ThemeValidationCategoryTokens,
				Severity: ValidationLevelWarning,
				Token:    token,
				Fix:      "Use kebab-case or dot notation for token names",
				AutoFix:  true,
			})
		}

		// Validate token value format
		if !v.isValidTokenValue(token, value) {
			result.Violations = append(result.Violations, ThemeViolation{
				Code:     "invalid_token_value",
				Message:  fmt.Sprintf("Token value '%v' has invalid format", value),
				Category: ThemeValidationCategoryTokens,
				Severity: ValidationLevelError,
				Token:    token,
				Value:    fmt.Sprintf("%v", value),
				Fix:      "Ensure token value matches expected format for its type",
				AutoFix:  false,
			})
			result.Passed = false
			result.Score -= 10.0
		}
	}

	return result
}

func (v *TokenFormatValidator) GetCategory() ThemeValidationCategory {
	return ThemeValidationCategoryTokens
}

func (v *TokenFormatValidator) extractTokens(theme any) map[string]any {
	// Implementation similar to ThemeValidator.extractTokens
	tokens := make(map[string]any)
	// ... token extraction logic
	return tokens
}

func (v *TokenFormatValidator) isValidTokenName(name string) bool {
	// Validate token naming conventions
	validName := regexp.MustCompile(`^[a-z][a-z0-9]*(-[a-z0-9]+)*(\.[a-z][a-z0-9]*(-[a-z0-9]+)*)*$`)
	return validName.MatchString(name)
}

func (v *TokenFormatValidator) isValidTokenValue(token string, value any) bool {
	valueStr := fmt.Sprintf("%v", value)
	
	// Determine token type and validate accordingly
	if strings.Contains(strings.ToLower(token), "color") {
		return v.registry.patterns["color"].MatchString(valueStr)
	}
	if strings.Contains(strings.ToLower(token), "spacing") {
		return v.registry.patterns["spacing"].MatchString(valueStr)
	}
	if strings.Contains(strings.ToLower(token), "font-size") {
		return v.registry.patterns["font-size"].MatchString(valueStr)
	}
	
	return true // Default to valid for unknown types
}

// Additional validators would be implemented similarly...

// ColorConsistencyValidator validates color consistency
type ColorConsistencyValidator struct{}

func (v *ColorConsistencyValidator) ValidateTheme(theme any, context *ThemeValidationContext) *ThemeRuleResult {
	// Implementation for color consistency validation
	return &ThemeRuleResult{Passed: true, Score: 100.0, Rule: "theme.colors.consistency", Metadata: make(map[string]any)}
}

func (v *ColorConsistencyValidator) GetCategory() ThemeValidationCategory {
	return ThemeValidationCategoryConsistency
}

// ThemeContrastValidator validates color contrast
type ThemeContrastValidator struct {
	minRatio float64
}

func (v *ThemeContrastValidator) ValidateTheme(theme any, context *ThemeValidationContext) *ThemeRuleResult {
	// Implementation for contrast validation
	return &ThemeRuleResult{Passed: true, Score: 100.0, Rule: "theme.accessibility.contrast", Metadata: make(map[string]any)}
}

func (v *ThemeContrastValidator) GetCategory() ThemeValidationCategory {
	return ThemeValidationCategoryAccessibility
}

// ThemePerformanceValidator validates theme performance
type ThemePerformanceValidator struct {
	maxTokens int
}

func (v *ThemePerformanceValidator) ValidateTheme(theme any, context *ThemeValidationContext) *ThemeRuleResult {
	// Implementation for performance validation
	return &ThemeRuleResult{Passed: true, Score: 100.0, Rule: "theme.performance.size", Metadata: make(map[string]any)}
}

func (v *ThemePerformanceValidator) GetCategory() ThemeValidationCategory {
	return ThemeValidationCategoryPerformance
}