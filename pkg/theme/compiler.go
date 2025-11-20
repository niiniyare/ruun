package theme

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// EnhancedCompiler provides improved CSS compilation with better token support
type EnhancedCompiler struct {
	config        *CompilerConfig
	tokenResolver *TokenResolver
	cache         *CompilerCache
	mu            sync.RWMutex
}

// CompilerConfig configures CSS compilation behavior
type CompilerConfig struct {
	// Output options
	Minify              bool   `json:"minify"`
	SourceMaps          bool   `json:"sourceMaps"`
	OutputFormat        string `json:"outputFormat"` // "css", "scss", "less"
	
	// Token options
	UseCustomProperties bool   `json:"useCustomProperties"` // Generate CSS custom properties
	UseTailwindClasses  bool   `json:"useTailwindClasses"`  // Generate Tailwind-like utility classes
	UseComponentClasses bool   `json:"useComponentClasses"` // Generate component-specific classes
	
	// Optimization options
	PurgeUnusedCSS      bool     `json:"purgeUnusedCSS"`
	OptimizeVariables   bool     `json:"optimizeVariables"`
	MergeSelectors      bool     `json:"mergeSelectors"`
	RemoveDuplicates    bool     `json:"removeDuplicates"`
	UsedComponents      []string `json:"usedComponents,omitempty"`
	
	// Advanced options
	EnableDarkMode      bool `json:"enableDarkMode"`
	EnableResponsive    bool `json:"enableResponsive"`
	EnableAnimations    bool `json:"enableAnimations"`
	EnablePrint         bool `json:"enablePrint"`
	SelectorPrefix      string `json:"selectorPrefix,omitempty"`
	VariablePrefix      string `json:"variablePrefix,omitempty"`
	
	// Performance options
	LazyLoad            bool `json:"lazyLoad"`
	ChunkByComponent    bool `json:"chunkByComponent"`
	EnableCaching       bool `json:"enableCaching"`
}

// CompilerCache provides caching for compiled CSS
type CompilerCache struct {
	compiled     map[string]*CachedCompilation
	chunks       map[string]*CSSChunk
	mu           sync.RWMutex
	lastCleanup  time.Time
}

// CachedCompilation represents cached compilation results
type CachedCompilation struct {
	ThemeID     string            `json:"themeId"`
	Config      *CompilerConfig   `json:"config"`
	CSS         string            `json:"css"`
	SourceMap   string            `json:"sourceMap,omitempty"`
	Variables   map[string]string `json:"variables"`
	Classes     map[string]string `json:"classes"`
	Chunks      []string          `json:"chunks,omitempty"`
	Size        int64             `json:"size"`
	Hash        string            `json:"hash"`
	Timestamp   time.Time         `json:"timestamp"`
}

// CSSChunk represents a chunk of compiled CSS
type CSSChunk struct {
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	CSS         string    `json:"css"`
	Dependencies []string `json:"dependencies,omitempty"`
	Size        int64     `json:"size"`
	Hash        string    `json:"hash"`
	Timestamp   time.Time `json:"timestamp"`
}

// CompilationResult provides detailed compilation results
type CompilationResult struct {
	Success       bool              `json:"success"`
	CSS           string            `json:"css"`
	SourceMap     string            `json:"sourceMap,omitempty"`
	Variables     map[string]string `json:"variables"`
	Classes       map[string]string `json:"classes"`
	Chunks        []*CSSChunk       `json:"chunks,omitempty"`
	Metrics       *CompilationMetrics `json:"metrics"`
	Warnings      []CompilationWarning `json:"warnings,omitempty"`
	Error         string            `json:"error,omitempty"`
}

// CompilationMetrics provides compilation performance metrics
type CompilationMetrics struct {
	CompilationTime   time.Duration `json:"compilationTime"`
	OutputSize        int64         `json:"outputSize"`
	CompressedSize    int64         `json:"compressedSize,omitempty"`
	VariableCount     int           `json:"variableCount"`
	ClassCount        int           `json:"classCount"`
	SelectorCount     int           `json:"selectorCount"`
	RuleCount         int           `json:"ruleCount"`
	CacheHits         int           `json:"cacheHits"`
	CacheMisses       int           `json:"cacheMisses"`
}

// CompilationWarning represents compilation warnings
type CompilationWarning struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Line    int    `json:"line,omitempty"`
	Column  int    `json:"column,omitempty"`
	Source  string `json:"source,omitempty"`
}

// NewEnhancedCompiler creates a new enhanced CSS compiler
func NewEnhancedCompiler(tokenResolver *TokenResolver, config *CompilerConfig) *EnhancedCompiler {
	if config == nil {
		config = DefaultCompilerConfig()
	}

	return &EnhancedCompiler{
		config:        config,
		tokenResolver: tokenResolver,
		cache:         NewCompilerCache(),
	}
}

// DefaultCompilerConfig returns default compiler configuration
func DefaultCompilerConfig() *CompilerConfig {
	return &CompilerConfig{
		Minify:              true,
		SourceMaps:          false,
		OutputFormat:        "css",
		UseCustomProperties: true,
		UseTailwindClasses:  true,
		UseComponentClasses: true,
		PurgeUnusedCSS:      false,
		OptimizeVariables:   true,
		MergeSelectors:      true,
		RemoveDuplicates:    true,
		EnableDarkMode:      true,
		EnableResponsive:    true,
		EnableAnimations:    true,
		EnablePrint:        false,
		VariablePrefix:     "--",
		LazyLoad:           false,
		ChunkByComponent:   false,
		EnableCaching:      true,
	}
}

// NewCompilerCache creates a new compiler cache
func NewCompilerCache() *CompilerCache {
	return &CompilerCache{
		compiled:    make(map[string]*CachedCompilation),
		chunks:      make(map[string]*CSSChunk),
		lastCleanup: time.Now(),
	}
}

// CompileTheme compiles a theme to CSS with enhanced features
func (ec *EnhancedCompiler) CompileTheme(theme *schema.Theme) (*CompilationResult, error) {
	startTime := time.Now()

	// Check cache first
	if ec.config.EnableCaching {
		if cached := ec.getFromCache(theme.ID); cached != nil {
			return &CompilationResult{
				Success:   true,
				CSS:       cached.CSS,
				SourceMap: cached.SourceMap,
				Variables: cached.Variables,
				Classes:   cached.Classes,
				Chunks:    ec.getCachedChunks(cached.Chunks),
				Metrics: &CompilationMetrics{
					CompilationTime: time.Since(startTime),
					OutputSize:      cached.Size,
					CacheHits:       1,
				},
			}, nil
		}
	}

	result := &CompilationResult{
		Success:   false,
		Variables: make(map[string]string),
		Classes:   make(map[string]string),
		Warnings:  []CompilationWarning{},
		Metrics:   &CompilationMetrics{},
	}

	// Start compilation
	compiler := &cssCompiler{
		config:        ec.config,
		tokenResolver: ec.tokenResolver,
		theme:         theme,
		result:        result,
	}

	css, err := compiler.compile()
	if err != nil {
		result.Error = err.Error()
		return result, err
	}

	result.CSS = css
	result.Success = true

	// Calculate metrics
	result.Metrics.CompilationTime = time.Since(startTime)
	result.Metrics.OutputSize = int64(len(css))
	result.Metrics.VariableCount = len(result.Variables)
	result.Metrics.ClassCount = len(result.Classes)

	// Cache result if enabled
	if ec.config.EnableCaching {
		ec.addToCache(theme.ID, result)
	}

	return result, nil
}

// cssCompiler handles the actual CSS compilation process
type cssCompiler struct {
	config        *CompilerConfig
	tokenResolver *TokenResolver
	theme         *schema.Theme
	result        *CompilationResult
	css           strings.Builder
	variables     map[string]string
	classes       map[string]string
}

// compile performs the CSS compilation
func (c *cssCompiler) compile() (string, error) {
	c.variables = make(map[string]string)
	c.classes = make(map[string]string)

	// Write header comment
	c.writeHeader()

	// Generate CSS custom properties (variables)
	if c.config.UseCustomProperties {
		if err := c.generateCustomProperties(); err != nil {
			return "", fmt.Errorf("failed to generate custom properties: %w", err)
		}
	}

	// Generate utility classes
	if c.config.UseTailwindClasses {
		if err := c.generateUtilityClasses(); err != nil {
			return "", fmt.Errorf("failed to generate utility classes: %w", err)
		}
	}

	// Generate component classes
	if c.config.UseComponentClasses {
		if err := c.generateComponentClasses(); err != nil {
			return "", fmt.Errorf("failed to generate component classes: %w", err)
		}
	}

	// Generate dark mode styles
	if c.config.EnableDarkMode && c.theme.DarkMode != nil && c.theme.DarkMode.Enabled {
		if err := c.generateDarkModeStyles(); err != nil {
			return "", fmt.Errorf("failed to generate dark mode styles: %w", err)
		}
	}

	// Generate responsive styles
	if c.config.EnableResponsive {
		if err := c.generateResponsiveStyles(); err != nil {
			return "", fmt.Errorf("failed to generate responsive styles: %w", err)
		}
	}

	// Generate animation styles
	if c.config.EnableAnimations {
		if err := c.generateAnimationStyles(); err != nil {
			return "", fmt.Errorf("failed to generate animation styles: %w", err)
		}
	}

	// Generate print styles
	if c.config.EnablePrint {
		if err := c.generatePrintStyles(); err != nil {
			return "", fmt.Errorf("failed to generate print styles: %w", err)
		}
	}

	// Store variables and classes in result
	c.result.Variables = c.variables
	c.result.Classes = c.classes

	cssOutput := c.css.String()

	// Post-process CSS
	if c.config.OptimizeVariables {
		cssOutput = c.optimizeVariables(cssOutput)
	}

	if c.config.MergeSelectors {
		cssOutput = c.mergeSelectors(cssOutput)
	}

	if c.config.RemoveDuplicates {
		cssOutput = c.removeDuplicates(cssOutput)
	}

	if c.config.Minify {
		cssOutput = c.minifyCSS(cssOutput)
	}

	return cssOutput, nil
}

// writeHeader writes the CSS header comment
func (c *cssCompiler) writeHeader() {
	c.css.WriteString(fmt.Sprintf("/*\n"))
	c.css.WriteString(fmt.Sprintf(" * Theme: %s\n", c.theme.Name))
	c.css.WriteString(fmt.Sprintf(" * Version: %s\n", c.theme.Version))
	if c.theme.Description != "" {
		c.css.WriteString(fmt.Sprintf(" * Description: %s\n", c.theme.Description))
	}
	if c.theme.Author != "" {
		c.css.WriteString(fmt.Sprintf(" * Author: %s\n", c.theme.Author))
	}
	c.css.WriteString(fmt.Sprintf(" * Generated: %s\n", time.Now().Format(time.RFC3339)))
	c.css.WriteString(" */\n\n")
}

// generateCustomProperties generates CSS custom properties from tokens
func (c *cssCompiler) generateCustomProperties() error {
	if c.theme.Tokens == nil {
		return nil
	}

	c.css.WriteString(":root {\n")
	c.css.WriteString("  /* Theme Variables */\n")

	// Generate color variables
	if err := c.generateColorVariables(); err != nil {
		return err
	}

	// Generate spacing variables
	if err := c.generateSpacingVariables(); err != nil {
		return err
	}

	// Generate typography variables
	if err := c.generateTypographyVariables(); err != nil {
		return err
	}

	// Generate other variables
	if err := c.generateOtherVariables(); err != nil {
		return err
	}

	c.css.WriteString("}\n\n")

	return nil
}

// generateColorVariables generates color CSS variables
func (c *cssCompiler) generateColorVariables() error {
	if c.theme.Tokens.Semantic == nil || c.theme.Tokens.Semantic.Colors == nil {
		return nil
	}

	colors := c.theme.Tokens.Semantic.Colors
	c.css.WriteString("  /* Colors */\n")

	// Background colors
	if colors.Background != nil {
		c.writeVariable("color-background", colors.Background.Default.String())
		c.writeVariable("color-background-subtle", colors.Background.Subtle.String())
		c.writeVariable("color-background-emphasis", colors.Background.Emphasis.String())
		c.writeVariable("color-background-overlay", colors.Background.Overlay.String())
	}

	// Text colors
	if colors.Text != nil {
		c.writeVariable("color-text", colors.Text.Default.String())
		c.writeVariable("color-text-subtle", colors.Text.Subtle.String())
		c.writeVariable("color-text-emphasis", colors.Text.Emphasis.String())
		c.writeVariable("color-text-accent", colors.Text.Accent.String())
	}

	// Border colors
	if colors.Border != nil {
		c.writeVariable("color-border", colors.Border.Default.String())
		c.writeVariable("color-border-subtle", colors.Border.Subtle.String())
		c.writeVariable("color-border-emphasis", colors.Border.Emphasis.String())
	}

	// Interactive colors
	if colors.Interactive != nil {
		c.writeVariable("color-primary", colors.Interactive.Primary.String())
		c.writeVariable("color-secondary", colors.Interactive.Secondary.String())
		c.writeVariable("color-accent", colors.Interactive.Accent.String())
		c.writeVariable("color-hover", colors.Interactive.Hover.String())
		c.writeVariable("color-active", colors.Interactive.Active.String())
		c.writeVariable("color-focus", colors.Interactive.Focus.String())
		c.writeVariable("color-disabled", colors.Interactive.Disabled.String())
	}

	// Feedback colors
	if colors.Feedback != nil {
		c.writeVariable("color-success", colors.Feedback.Success.String())
		c.writeVariable("color-warning", colors.Feedback.Warning.String())
		c.writeVariable("color-error", colors.Feedback.Error.String())
		c.writeVariable("color-info", colors.Feedback.Info.String())
	}

	return nil
}

// generateSpacingVariables generates spacing CSS variables
func (c *cssCompiler) generateSpacingVariables() error {
	if c.theme.Tokens.Semantic == nil || c.theme.Tokens.Semantic.Spacing == nil {
		return nil
	}

	spacing := c.theme.Tokens.Semantic.Spacing
	c.css.WriteString("  /* Spacing */\n")

	// Component spacing
	if spacing.Component != nil {
		c.writeVariable("spacing-tight", spacing.Component.Tight.String())
		c.writeVariable("spacing-default", spacing.Component.Default.String())
		c.writeVariable("spacing-loose", spacing.Component.Loose.String())
	}

	// Layout spacing
	if spacing.Layout != nil {
		c.writeVariable("spacing-section", spacing.Layout.Section.String())
		c.writeVariable("spacing-page", spacing.Layout.Page.String())
	}

	return nil
}

// generateTypographyVariables generates typography CSS variables
func (c *cssCompiler) generateTypographyVariables() error {
	if c.theme.Tokens.Semantic == nil || c.theme.Tokens.Semantic.Typography == nil {
		return nil
	}

	typography := c.theme.Tokens.Semantic.Typography
	c.css.WriteString("  /* Typography */\n")

	// Heading typography
	if typography.Headings != nil {
		c.writeVariable("font-size-heading", typography.Headings.FontSize.String())
		c.writeVariable("font-weight-heading", typography.Headings.FontWeight.String())
		c.writeVariable("line-height-heading", typography.Headings.LineHeight.String())
	}

	// Body typography
	if typography.Body != nil {
		c.writeVariable("font-size-body", typography.Body.FontSize.String())
		c.writeVariable("font-weight-body", typography.Body.FontWeight.String())
		c.writeVariable("line-height-body", typography.Body.LineHeight.String())
	}

	// Label typography
	if typography.Labels != nil {
		c.writeVariable("font-size-label", typography.Labels.FontSize.String())
		c.writeVariable("font-weight-label", typography.Labels.FontWeight.String())
	}

	// Caption typography
	if typography.Captions != nil {
		c.writeVariable("font-size-caption", typography.Captions.FontSize.String())
		c.writeVariable("font-weight-caption", typography.Captions.FontWeight.String())
	}

	// Code typography
	if typography.Code != nil {
		c.writeVariable("font-size-code", typography.Code.FontSize.String())
		c.writeVariable("font-family-code", typography.Code.FontFamily.String())
	}

	return nil
}

// generateOtherVariables generates other CSS variables
func (c *cssCompiler) generateOtherVariables() error {
	if c.theme.Tokens.Semantic == nil || c.theme.Tokens.Semantic.Interactive == nil {
		return nil
	}

	interactive := c.theme.Tokens.Semantic.Interactive
	c.css.WriteString("  /* Interactive */\n")

	// Border radius
	if interactive.BorderRadius != nil {
		c.writeVariable("radius-small", interactive.BorderRadius.Small.String())
		c.writeVariable("radius-medium", interactive.BorderRadius.Medium.String())
		c.writeVariable("radius-large", interactive.BorderRadius.Large.String())
	}

	// Shadows
	if interactive.Shadow != nil {
		c.writeVariable("shadow-small", interactive.Shadow.Small.String())
		c.writeVariable("shadow-medium", interactive.Shadow.Medium.String())
		c.writeVariable("shadow-large", interactive.Shadow.Large.String())
	}

	return nil
}

// generateUtilityClasses generates Tailwind-like utility classes
func (c *cssCompiler) generateUtilityClasses() error {
	c.css.WriteString("/* Utility Classes */\n")

	// Color utilities
	c.generateColorUtilities()

	// Spacing utilities
	c.generateSpacingUtilities()

	// Typography utilities
	c.generateTypographyUtilities()

	// Border utilities
	c.generateBorderUtilities()

	c.css.WriteString("\n")
	return nil
}

// generateColorUtilities generates color utility classes
func (c *cssCompiler) generateColorUtilities() {
	c.css.WriteString("/* Color Utilities */\n")

	colors := []struct {
		name     string
		variable string
	}{
		{"text-primary", "var(--color-text)"},
		{"text-subtle", "var(--color-text-subtle)"},
		{"text-emphasis", "var(--color-text-emphasis)"},
		{"bg-primary", "var(--color-background)"},
		{"bg-subtle", "var(--color-background-subtle)"},
		{"bg-emphasis", "var(--color-background-emphasis)"},
		{"border-default", "var(--color-border)"},
		{"border-subtle", "var(--color-border-subtle)"},
		{"border-emphasis", "var(--color-border-emphasis)"},
	}

	for _, color := range colors {
		if strings.HasPrefix(color.name, "text-") {
			c.css.WriteString(fmt.Sprintf(".%s { color: %s; }\n", color.name, color.variable))
		} else if strings.HasPrefix(color.name, "bg-") {
			c.css.WriteString(fmt.Sprintf(".%s { background-color: %s; }\n", color.name, color.variable))
		} else if strings.HasPrefix(color.name, "border-") {
			c.css.WriteString(fmt.Sprintf(".%s { border-color: %s; }\n", color.name, color.variable))
		}
		c.classes[color.name] = color.name
	}
}

// generateSpacingUtilities generates spacing utility classes
func (c *cssCompiler) generateSpacingUtilities() {
	c.css.WriteString("/* Spacing Utilities */\n")

	spacings := []struct {
		name     string
		variable string
	}{
		{"p-tight", "var(--spacing-tight)"},
		{"p-default", "var(--spacing-default)"},
		{"p-loose", "var(--spacing-loose)"},
		{"m-tight", "var(--spacing-tight)"},
		{"m-default", "var(--spacing-default)"},
		{"m-loose", "var(--spacing-loose)"},
	}

	for _, spacing := range spacings {
		if strings.HasPrefix(spacing.name, "p-") {
			c.css.WriteString(fmt.Sprintf(".%s { padding: %s; }\n", spacing.name, spacing.variable))
		} else if strings.HasPrefix(spacing.name, "m-") {
			c.css.WriteString(fmt.Sprintf(".%s { margin: %s; }\n", spacing.name, spacing.variable))
		}
		c.classes[spacing.name] = spacing.name
	}
}

// generateTypographyUtilities generates typography utility classes
func (c *cssCompiler) generateTypographyUtilities() {
	c.css.WriteString("/* Typography Utilities */\n")

	typography := []struct {
		name       string
		properties map[string]string
	}{
		{"text-heading", map[string]string{
			"font-size":   "var(--font-size-heading)",
			"font-weight": "var(--font-weight-heading)",
			"line-height": "var(--line-height-heading)",
		}},
		{"text-body", map[string]string{
			"font-size":   "var(--font-size-body)",
			"font-weight": "var(--font-weight-body)",
			"line-height": "var(--line-height-body)",
		}},
		{"text-caption", map[string]string{
			"font-size":   "var(--font-size-caption)",
			"font-weight": "var(--font-weight-caption)",
		}},
		{"text-code", map[string]string{
			"font-size":   "var(--font-size-code)",
			"font-family": "var(--font-family-code)",
		}},
	}

	for _, typo := range typography {
		c.css.WriteString(fmt.Sprintf(".%s {\n", typo.name))
		for property, value := range typo.properties {
			c.css.WriteString(fmt.Sprintf("  %s: %s;\n", property, value))
		}
		c.css.WriteString("}\n")
		c.classes[typo.name] = typo.name
	}
}

// generateBorderUtilities generates border utility classes
func (c *cssCompiler) generateBorderUtilities() {
	c.css.WriteString("/* Border Utilities */\n")

	borders := []struct {
		name     string
		variable string
	}{
		{"rounded-small", "var(--radius-small)"},
		{"rounded-medium", "var(--radius-medium)"},
		{"rounded-large", "var(--radius-large)"},
		{"shadow-small", "var(--shadow-small)"},
		{"shadow-medium", "var(--shadow-medium)"},
		{"shadow-large", "var(--shadow-large)"},
	}

	for _, border := range borders {
		if strings.HasPrefix(border.name, "rounded-") {
			c.css.WriteString(fmt.Sprintf(".%s { border-radius: %s; }\n", border.name, border.variable))
		} else if strings.HasPrefix(border.name, "shadow-") {
			c.css.WriteString(fmt.Sprintf(".%s { box-shadow: %s; }\n", border.name, border.variable))
		}
		c.classes[border.name] = border.name
	}
}

// generateComponentClasses generates component-specific CSS classes
func (c *cssCompiler) generateComponentClasses() error {
	c.css.WriteString("/* Component Classes */\n")

	if c.theme.Tokens.Components == nil {
		return nil
	}

	// Generate button classes
	if c.theme.Tokens.Components.Button != nil || c.shouldIncludeComponent("button") {
		c.generateButtonClasses()
	}

	// Generate input classes
	if c.theme.Tokens.Components.Input != nil || c.shouldIncludeComponent("input") {
		c.generateInputClasses()
	}

	// Generate card classes
	if c.theme.Tokens.Components.Card != nil || c.shouldIncludeComponent("card") {
		c.generateCardClasses()
	}

	// Generate modal classes
	if c.theme.Tokens.Components.Modal != nil || c.shouldIncludeComponent("modal") {
		c.generateModalClasses()
	}

	c.css.WriteString("\n")
	return nil
}

// shouldIncludeComponent checks if a component should be included based on configuration
func (c *cssCompiler) shouldIncludeComponent(component string) bool {
	if len(c.config.UsedComponents) == 0 {
		return true // Include all if not specified
	}

	for _, used := range c.config.UsedComponents {
		if used == component {
			return true
		}
	}
	return false
}

// generateButtonClasses generates button component classes
func (c *cssCompiler) generateButtonClasses() {
	c.css.WriteString("/* Button Component */\n")

	// Base button class
	c.css.WriteString(".button {\n")
	c.css.WriteString("  display: inline-flex;\n")
	c.css.WriteString("  align-items: center;\n")
	c.css.WriteString("  justify-content: center;\n")
	c.css.WriteString("  padding: var(--spacing-default);\n")
	c.css.WriteString("  border-radius: var(--radius-medium);\n")
	c.css.WriteString("  font-size: var(--font-size-body);\n")
	c.css.WriteString("  font-weight: var(--font-weight-label);\n")
	c.css.WriteString("  border: 1px solid transparent;\n")
	c.css.WriteString("  cursor: pointer;\n")
	c.css.WriteString("  transition: all 0.2s ease;\n")
	c.css.WriteString("  text-decoration: none;\n")
	c.css.WriteString("}\n")
	c.classes["button"] = "button"

	// Button variants
	variants := []struct {
		name       string
		properties map[string]string
	}{
		{"button-primary", map[string]string{
			"background-color": "var(--color-primary)",
			"color":           "var(--color-background)",
			"border-color":    "var(--color-primary)",
		}},
		{"button-secondary", map[string]string{
			"background-color": "var(--color-background)",
			"color":           "var(--color-text)",
			"border-color":    "var(--color-border)",
		}},
		{"button-outline", map[string]string{
			"background-color": "transparent",
			"color":           "var(--color-primary)",
			"border-color":    "var(--color-primary)",
		}},
		{"button-destructive", map[string]string{
			"background-color": "var(--color-error)",
			"color":           "var(--color-background)",
			"border-color":    "var(--color-error)",
		}},
		{"button-ghost", map[string]string{
			"background-color": "transparent",
			"color":           "var(--color-text)",
			"border-color":    "transparent",
		}},
	}

	for _, variant := range variants {
		c.css.WriteString(fmt.Sprintf(".%s {\n", variant.name))
		for property, value := range variant.properties {
			c.css.WriteString(fmt.Sprintf("  %s: %s;\n", property, value))
		}
		c.css.WriteString("}\n")

		// Hover states
		c.css.WriteString(fmt.Sprintf(".%s:hover {\n", variant.name))
		c.css.WriteString("  opacity: 0.9;\n")
		c.css.WriteString("  transform: translateY(-1px);\n")
		c.css.WriteString("}\n")

		// Focus states
		c.css.WriteString(fmt.Sprintf(".%s:focus {\n", variant.name))
		c.css.WriteString("  outline: none;\n")
		c.css.WriteString("  box-shadow: 0 0 0 2px var(--color-focus);\n")
		c.css.WriteString("}\n")

		// Disabled states
		c.css.WriteString(fmt.Sprintf(".%s:disabled {\n", variant.name))
		c.css.WriteString("  opacity: 0.5;\n")
		c.css.WriteString("  cursor: not-allowed;\n")
		c.css.WriteString("  transform: none;\n")
		c.css.WriteString("}\n")

		c.classes[variant.name] = variant.name
	}

	// Button sizes
	sizes := []struct {
		name       string
		properties map[string]string
	}{
		{"button-xs", map[string]string{
			"padding":    "var(--spacing-tight)",
			"font-size":  "0.75rem",
		}},
		{"button-sm", map[string]string{
			"padding":    "var(--spacing-tight)",
			"font-size":  "0.875rem",
		}},
		{"button-md", map[string]string{
			"padding":    "var(--spacing-default)",
			"font-size":  "var(--font-size-body)",
		}},
		{"button-lg", map[string]string{
			"padding":    "var(--spacing-loose)",
			"font-size":  "1.125rem",
		}},
		{"button-xl", map[string]string{
			"padding":    "var(--spacing-loose)",
			"font-size":  "1.25rem",
		}},
	}

	for _, size := range sizes {
		c.css.WriteString(fmt.Sprintf(".%s {\n", size.name))
		for property, value := range size.properties {
			c.css.WriteString(fmt.Sprintf("  %s: %s;\n", property, value))
		}
		c.css.WriteString("}\n")
		c.classes[size.name] = size.name
	}
}

// generateInputClasses generates input component classes
func (c *cssCompiler) generateInputClasses() {
	c.css.WriteString("/* Input Component */\n")

	// Base input class
	c.css.WriteString(".input {\n")
	c.css.WriteString("  width: 100%;\n")
	c.css.WriteString("  padding: var(--spacing-default);\n")
	c.css.WriteString("  border: 1px solid var(--color-border);\n")
	c.css.WriteString("  border-radius: var(--radius-medium);\n")
	c.css.WriteString("  background-color: var(--color-background);\n")
	c.css.WriteString("  color: var(--color-text);\n")
	c.css.WriteString("  font-size: var(--font-size-body);\n")
	c.css.WriteString("  transition: border-color 0.2s ease, box-shadow 0.2s ease;\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".input:focus {\n")
	c.css.WriteString("  outline: none;\n")
	c.css.WriteString("  border-color: var(--color-focus);\n")
	c.css.WriteString("  box-shadow: 0 0 0 1px var(--color-focus);\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".input:disabled {\n")
	c.css.WriteString("  background-color: var(--color-background-subtle);\n")
	c.css.WriteString("  color: var(--color-text-subtle);\n")
	c.css.WriteString("  cursor: not-allowed;\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".input-error {\n")
	c.css.WriteString("  border-color: var(--color-error);\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".input-success {\n")
	c.css.WriteString("  border-color: var(--color-success);\n")
	c.css.WriteString("}\n")

	c.classes["input"] = "input"
	c.classes["input-error"] = "input-error"
	c.classes["input-success"] = "input-success"
}

// generateCardClasses generates card component classes
func (c *cssCompiler) generateCardClasses() {
	c.css.WriteString("/* Card Component */\n")

	c.css.WriteString(".card {\n")
	c.css.WriteString("  background-color: var(--color-background);\n")
	c.css.WriteString("  border: 1px solid var(--color-border);\n")
	c.css.WriteString("  border-radius: var(--radius-medium);\n")
	c.css.WriteString("  box-shadow: var(--shadow-small);\n")
	c.css.WriteString("  padding: var(--spacing-loose);\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".card-header {\n")
	c.css.WriteString("  padding-bottom: var(--spacing-default);\n")
	c.css.WriteString("  border-bottom: 1px solid var(--color-border-subtle);\n")
	c.css.WriteString("  margin-bottom: var(--spacing-default);\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".card-footer {\n")
	c.css.WriteString("  padding-top: var(--spacing-default);\n")
	c.css.WriteString("  border-top: 1px solid var(--color-border-subtle);\n")
	c.css.WriteString("  margin-top: var(--spacing-default);\n")
	c.css.WriteString("}\n")

	c.classes["card"] = "card"
	c.classes["card-header"] = "card-header"
	c.classes["card-footer"] = "card-footer"
}

// generateModalClasses generates modal component classes
func (c *cssCompiler) generateModalClasses() {
	c.css.WriteString("/* Modal Component */\n")

	c.css.WriteString(".modal-overlay {\n")
	c.css.WriteString("  position: fixed;\n")
	c.css.WriteString("  top: 0;\n")
	c.css.WriteString("  left: 0;\n")
	c.css.WriteString("  right: 0;\n")
	c.css.WriteString("  bottom: 0;\n")
	c.css.WriteString("  background-color: var(--color-background-overlay);\n")
	c.css.WriteString("  backdrop-filter: blur(4px);\n")
	c.css.WriteString("  z-index: 1000;\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".modal {\n")
	c.css.WriteString("  position: fixed;\n")
	c.css.WriteString("  top: 50%;\n")
	c.css.WriteString("  left: 50%;\n")
	c.css.WriteString("  transform: translate(-50%, -50%);\n")
	c.css.WriteString("  background-color: var(--color-background);\n")
	c.css.WriteString("  border-radius: var(--radius-large);\n")
	c.css.WriteString("  box-shadow: var(--shadow-large);\n")
	c.css.WriteString("  padding: var(--spacing-loose);\n")
	c.css.WriteString("  max-width: 90vw;\n")
	c.css.WriteString("  max-height: 90vh;\n")
	c.css.WriteString("  overflow: auto;\n")
	c.css.WriteString("  z-index: 1001;\n")
	c.css.WriteString("}\n")

	c.classes["modal-overlay"] = "modal-overlay"
	c.classes["modal"] = "modal"
}

// generateDarkModeStyles generates dark mode CSS
func (c *cssCompiler) generateDarkModeStyles() error {
	if c.theme.DarkMode == nil || c.theme.DarkMode.DarkTokens == nil {
		return nil
	}

	c.css.WriteString("/* Dark Mode */\n")

	strategy := c.theme.DarkMode.Strategy
	switch strategy {
	case "class":
		c.css.WriteString(".dark {\n")
	case "media":
		c.css.WriteString("@media (prefers-color-scheme: dark) {\n")
		c.css.WriteString("  :root {\n")
	default:
		c.css.WriteString("@media (prefers-color-scheme: dark) {\n")
		c.css.WriteString("  :root {\n")
	}

	// Generate dark mode color variables
	darkColors := c.theme.DarkMode.DarkTokens.Semantic.Colors
	if darkColors != nil {
		if darkColors.Background != nil {
			c.writeVariable("color-background", darkColors.Background.Default.String())
			c.writeVariable("color-background-subtle", darkColors.Background.Subtle.String())
			c.writeVariable("color-background-emphasis", darkColors.Background.Emphasis.String())
		}

		if darkColors.Text != nil {
			c.writeVariable("color-text", darkColors.Text.Default.String())
			c.writeVariable("color-text-subtle", darkColors.Text.Subtle.String())
			c.writeVariable("color-text-emphasis", darkColors.Text.Emphasis.String())
		}
	}

	if strategy == "class" {
		c.css.WriteString("}\n")
	} else {
		c.css.WriteString("  }\n")
		c.css.WriteString("}\n")
	}

	c.css.WriteString("\n")
	return nil
}

// generateResponsiveStyles generates responsive CSS
func (c *cssCompiler) generateResponsiveStyles() error {
	c.css.WriteString("/* Responsive Styles */\n")

	breakpoints := map[string]string{
		"sm": "640px",
		"md": "768px",
		"lg": "1024px",
		"xl": "1280px",
	}

	for name, size := range breakpoints {
		c.css.WriteString(fmt.Sprintf("@media (min-width: %s) {\n", size))
		c.css.WriteString(fmt.Sprintf("  .%s\\:hidden { display: none; }\n", name))
		c.css.WriteString(fmt.Sprintf("  .%s\\:block { display: block; }\n", name))
		c.css.WriteString(fmt.Sprintf("  .%s\\:flex { display: flex; }\n", name))
		c.css.WriteString("}\n")
	}

	c.css.WriteString("\n")
	return nil
}

// generateAnimationStyles generates animation CSS
func (c *cssCompiler) generateAnimationStyles() error {
	c.css.WriteString("/* Animations */\n")

	// Transition utilities
	c.css.WriteString(".transition-all { transition: all 0.2s ease; }\n")
	c.css.WriteString(".transition-colors { transition: color 0.2s ease, background-color 0.2s ease, border-color 0.2s ease; }\n")
	c.css.WriteString(".transition-transform { transition: transform 0.2s ease; }\n")

	// Animation classes
	c.css.WriteString("@keyframes fadeIn {\n")
	c.css.WriteString("  from { opacity: 0; }\n")
	c.css.WriteString("  to { opacity: 1; }\n")
	c.css.WriteString("}\n")

	c.css.WriteString("@keyframes slideIn {\n")
	c.css.WriteString("  from { transform: translateY(-10px); opacity: 0; }\n")
	c.css.WriteString("  to { transform: translateY(0); opacity: 1; }\n")
	c.css.WriteString("}\n")

	c.css.WriteString(".animate-fadeIn { animation: fadeIn 0.3s ease; }\n")
	c.css.WriteString(".animate-slideIn { animation: slideIn 0.3s ease; }\n")

	c.classes["transition-all"] = "transition-all"
	c.classes["transition-colors"] = "transition-colors"
	c.classes["animate-fadeIn"] = "animate-fadeIn"
	c.classes["animate-slideIn"] = "animate-slideIn"

	c.css.WriteString("\n")
	return nil
}

// generatePrintStyles generates print CSS
func (c *cssCompiler) generatePrintStyles() error {
	c.css.WriteString("/* Print Styles */\n")

	c.css.WriteString("@media print {\n")
	c.css.WriteString("  * {\n")
	c.css.WriteString("    background: white !important;\n")
	c.css.WriteString("    color: black !important;\n")
	c.css.WriteString("    box-shadow: none !important;\n")
	c.css.WriteString("  }\n")
	c.css.WriteString("  .no-print { display: none !important; }\n")
	c.css.WriteString("}\n")

	c.css.WriteString("\n")
	return nil
}

// writeVariable writes a CSS variable with proper formatting
func (c *cssCompiler) writeVariable(name, value string) {
	variableName := fmt.Sprintf("%s%s", c.config.VariablePrefix, name)
	c.css.WriteString(fmt.Sprintf("  %s: %s;\n", variableName, value))
	c.variables[variableName] = value
}

// Post-processing methods

func (c *cssCompiler) optimizeVariables(css string) string {
	// Implement variable optimization
	return css
}

func (c *cssCompiler) mergeSelectors(css string) string {
	// Implement selector merging
	return css
}

func (c *cssCompiler) removeDuplicates(css string) string {
	// Implement duplicate removal
	return css
}

func (c *cssCompiler) minifyCSS(css string) string {
	// Basic minification
	css = strings.ReplaceAll(css, "\n", "")
	css = strings.ReplaceAll(css, "  ", " ")
	css = strings.ReplaceAll(css, ": ", ":")
	css = strings.ReplaceAll(css, "; ", ";")
	css = strings.ReplaceAll(css, " {", "{")
	css = strings.ReplaceAll(css, "{ ", "{")
	css = strings.ReplaceAll(css, " }", "}")
	css = strings.ReplaceAll(css, "} ", "}")
	return css
}

// Cache management methods

func (ec *EnhancedCompiler) getFromCache(themeID string) *CachedCompilation {
	if !ec.config.EnableCaching {
		return nil
	}

	ec.cache.mu.RLock()
	defer ec.cache.mu.RUnlock()

	cached, exists := ec.cache.compiled[themeID]
	if !exists {
		return nil
	}

	// Check if cache is still valid (24 hours)
	if time.Since(cached.Timestamp) > 24*time.Hour {
		return nil
	}

	return cached
}

func (ec *EnhancedCompiler) addToCache(themeID string, result *CompilationResult) {
	if !ec.config.EnableCaching {
		return
	}

	cached := &CachedCompilation{
		ThemeID:   themeID,
		Config:    ec.config,
		CSS:       result.CSS,
		SourceMap: result.SourceMap,
		Variables: result.Variables,
		Classes:   result.Classes,
		Size:      result.Metrics.OutputSize,
		Hash:      ec.calculateHash(result.CSS),
		Timestamp: time.Now(),
	}

	ec.cache.mu.Lock()
	defer ec.cache.mu.Unlock()

	ec.cache.compiled[themeID] = cached

	// Clean up old entries
	if time.Since(ec.cache.lastCleanup) > time.Hour {
		ec.cleanupCache()
		ec.cache.lastCleanup = time.Now()
	}
}

func (ec *EnhancedCompiler) getCachedChunks(chunkNames []string) []*CSSChunk {
	if !ec.config.EnableCaching || len(chunkNames) == 0 {
		return nil
	}

	ec.cache.mu.RLock()
	defer ec.cache.mu.RUnlock()

	chunks := make([]*CSSChunk, 0, len(chunkNames))
	for _, name := range chunkNames {
		if chunk, exists := ec.cache.chunks[name]; exists {
			chunks = append(chunks, chunk)
		}
	}

	return chunks
}

func (ec *EnhancedCompiler) cleanupCache() {
	cutoff := time.Now().Add(-24 * time.Hour)
	
	for id, cached := range ec.cache.compiled {
		if cached.Timestamp.Before(cutoff) {
			delete(ec.cache.compiled, id)
		}
	}

	for name, chunk := range ec.cache.chunks {
		if chunk.Timestamp.Before(cutoff) {
			delete(ec.cache.chunks, name)
		}
	}
}

func (ec *EnhancedCompiler) calculateHash(content string) string {
	// Simple hash calculation - in production use crypto/sha256
	return fmt.Sprintf("hash-%d", len(content))
}

// GetCacheStats returns cache statistics
func (ec *EnhancedCompiler) GetCacheStats() map[string]any {
	ec.cache.mu.RLock()
	defer ec.cache.mu.RUnlock()

	return map[string]any{
		"compiledThemes": len(ec.cache.compiled),
		"cachedChunks":   len(ec.cache.chunks),
		"lastCleanup":    ec.cache.lastCleanup,
	}
}