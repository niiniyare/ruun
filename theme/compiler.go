package theme

import (
	"bytes"
	"context"
	"fmt"
	"maps"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

// Compiler compiles design tokens into optimized CSS with caching and advanced features.
type Compiler struct {
	cache      *ristretto.Cache
	cacheTTL   time.Duration
	mu         sync.RWMutex
	config     *CompilerConfig
	stats      *CompilerStats
	statsMutex sync.Mutex
}

// CompilerConfig configures the CSS compiler behavior.
type CompilerConfig struct {
	EnableCaching         bool
	EnableMinify          bool
	EnableSourceMap       bool
	EnableVariables       bool
	Prefix                string
	CacheTTL              time.Duration
	CacheSize             int64
	IncludeComments       bool
	GenerateUtilities     bool
	GenerateCSSProperties bool
	GenerateThemeClasses  bool
	FlattenNestedTokens   bool
	OptimizeOutput        bool // New: Enable CSS optimizations
	GenerateColorSchemes  bool // New: Generate color scheme variations
	IncludeBaseStyles     bool // New: Include base CSS reset/normalization
}

// DefaultCompilerConfig returns production-ready compiler configuration.
func DefaultCompilerConfig() *CompilerConfig {
	return &CompilerConfig{
		EnableCaching:         true,
		EnableMinify:          true,
		EnableSourceMap:       false,
		EnableVariables:       true,
		Prefix:                "",
		CacheTTL:              10 * time.Minute,
		CacheSize:             20 << 20, // 20MB
		IncludeComments:       false,
		GenerateUtilities:     false,
		GenerateCSSProperties: true,
		GenerateThemeClasses:  true,
		FlattenNestedTokens:   true,
		OptimizeOutput:        true,
		GenerateColorSchemes:  false,
		IncludeBaseStyles:     false,
	}
}

// NewCompiler creates a new CSS compiler with the given configuration.
func NewCompiler(config *CompilerConfig) (*Compiler, error) {
	if config == nil {
		config = DefaultCompilerConfig()
	}

	var cache *ristretto.Cache
	var err error

	if config.EnableCaching {
		cache, err = ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e4,
			MaxCost:     config.CacheSize,
			BufferItems: 64,
			Metrics:     true,
		})
		if err != nil {
			return nil, WrapError(ErrCodeCompilation, "failed to create cache", err)
		}
	}

	return &Compiler{
		cache:    cache,
		cacheTTL: config.CacheTTL,
		config:   config,
		stats:    &CompilerStats{},
	}, nil
}

// Compile compiles design tokens into optimized CSS.
func (c *Compiler) Compile(ctx context.Context, tokens *Tokens, theme *Theme) (string, error) {
	start := time.Now()

	// Validate inputs
	if err := c.validateInputs(tokens, theme); err != nil {
		return "", err
	}

	// Check cache
	cacheKey := c.buildCacheKey(theme.ID)
	if cached, found := c.getFromCache(cacheKey); found {
		c.recordStats(time.Since(start), true)
		return cached, nil
	}

	// Generate CSS
	css, err := c.generateCSS(ctx, tokens, theme)
	if err != nil {
		return "", err
	}

	// Optimize and cache
	if c.config.OptimizeOutput {
		css = c.optimizeCSS(css)
	}

	c.setCache(cacheKey, css)
	c.recordStats(time.Since(start), false)

	return css, nil
}

// validateInputs validates compiler inputs
func (c *Compiler) validateInputs(tokens *Tokens, theme *Theme) error {
	if tokens == nil {
		return NewError(ErrCodeValidation, "tokens cannot be nil")
	}
	if theme == nil {
		return NewError(ErrCodeValidation, "theme cannot be nil")
	}
	if err := tokens.Validate(); err != nil {
		return WrapError(ErrCodeValidation, "invalid tokens", err)
	}
	return nil
}

// generateCSS generates the complete CSS output
func (c *Compiler) generateCSS(_ context.Context, tokens *Tokens, theme *Theme) (string, error) {
	var buf bytes.Buffer

	// Add header comments
	c.writeHeader(&buf, theme)

	// Generate base styles if enabled
	if c.config.IncludeBaseStyles {
		c.writeBaseStyles(&buf)
	}

	// Generate theme scope
	c.writeThemeScope(&buf, tokens, theme)

	// Generate dark mode
	if theme.DarkMode != nil && theme.DarkMode.Enabled {
		c.writeDarkMode(&buf, theme)
	}

	// Generate utilities
	if c.config.GenerateUtilities {
		c.writeUtilityClasses(&buf, tokens)
	}

	// Generate color schemes
	if c.config.GenerateColorSchemes {
		c.writeColorSchemes(&buf, tokens, theme)
	}

	// Add custom CSS and accessibility
	c.writeThemeExtensions(&buf, theme)

	return buf.String(), nil
}

// writeHeader writes CSS header with metadata
func (c *Compiler) writeHeader(buf *bytes.Buffer, theme *Theme) {
	if !c.config.IncludeComments {
		return
	}

	fmt.Fprintf(buf, "/* Theme: %s (%s) */\n", theme.Name, theme.ID)
	fmt.Fprintf(buf, "/* Generated: %s */\n", time.Now().Format(time.RFC3339))
	if theme.Description != "" {
		fmt.Fprintf(buf, "/* Description: %s */\n", theme.Description)
	}
	if theme.Version != "" {
		fmt.Fprintf(buf, "/* Version: %s */\n", theme.Version)
	}
	buf.WriteString("\n")
}

// writeBaseStyles writes base CSS reset/normalization
func (c *Compiler) writeBaseStyles(buf *bytes.Buffer) {
	if c.config.IncludeComments {
		buf.WriteString("/* Base Styles */\n")
	}

	buf.WriteString(`
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

html {
  line-height: 1.5;
  -webkit-text-size-adjust: 100%;
}

body {
  font-family: var(--font-sans, system-ui, sans-serif);
  line-height: inherit;
}

img, svg, video, canvas, audio, iframe, embed, object {
  display: block;
  max-width: 100%;
}

button, input, optgroup, select, textarea {
  font-family: inherit;
  font-size: 100%;
  line-height: inherit;
}
`)
}

// writeThemeScope writes the main theme scope with variables and properties
func (c *Compiler) writeThemeScope(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	// Open theme class wrapper
	if c.config.GenerateThemeClasses {
		themeClass := c.makeThemeClassName(theme.ID)
		fmt.Fprintf(buf, ".%s {\n", themeClass)
	} else {
		buf.WriteString(":root {\n")
	}

	// Generate CSS variables
	if c.config.EnableVariables {
		cssVars := c.generateCSSVariables(tokens)
		buf.WriteString(cssVars)
	}

	// Generate CSS properties
	if c.config.GenerateCSSProperties {
		cssProps := c.generateCSSProperties(tokens, theme)
		buf.WriteString(cssProps)
	}

	// Close theme scope
	buf.WriteString("}\n\n")
}

// writeDarkMode writes dark mode styles
func (c *Compiler) writeDarkMode(buf *bytes.Buffer, theme *Theme) {
	if theme.DarkMode == nil || !theme.DarkMode.Enabled || theme.DarkMode.DarkTokens == nil {
		return
	}

	selector := c.getDarkModeSelector(theme.DarkMode.Strategy)

	if c.config.IncludeComments {
		fmt.Fprintf(buf, "/* Dark Mode (%s) */\n", theme.DarkMode.Strategy)
	}

	// Determine selector based on theme class strategy
	var darkSelector string
	if c.config.GenerateThemeClasses {
		themeClass := c.makeThemeClassName(theme.ID)
		darkSelector = fmt.Sprintf(".%s%s", themeClass, selector)
	} else {
		darkSelector = selector
	}

	fmt.Fprintf(buf, "%s {\n", darkSelector)

	// Generate dark mode variables
	if c.config.EnableVariables {
		darkVars := c.generateCSSVariables(theme.DarkMode.DarkTokens)
		buf.WriteString(darkVars)
	}

	// Generate dark mode CSS properties
	if c.config.GenerateCSSProperties {
		darkProps := c.generateDarkModeCSSProperties(theme.DarkMode.DarkTokens, theme)
		buf.WriteString(darkProps)
	}

	buf.WriteString("}\n\n")
}

// writeUtilityClasses generates utility CSS classes
func (c *Compiler) writeUtilityClasses(buf *bytes.Buffer, tokens *Tokens) {
	if c.config.IncludeComments {
		buf.WriteString("/* Utility Classes */\n")
	}

	if tokens.Primitives != nil {
		// Spacing utilities
		if tokens.Primitives.Spacing != nil {
			for key, value := range tokens.Primitives.Spacing {
				className := strings.ReplaceAll(key, ".", "-")
				fmt.Fprintf(buf, ".m-%s { margin: %s; }\n", className, value)
				fmt.Fprintf(buf, ".p-%s { padding: %s; }\n", className, value)
				fmt.Fprintf(buf, ".mt-%s { margin-top: %s; }\n", className, value)
				fmt.Fprintf(buf, ".mb-%s { margin-bottom: %s; }\n", className, value)
			}
		}

		// Color utilities
		if tokens.Primitives.Colors != nil {
			for key, value := range tokens.Primitives.Colors {
				className := strings.ReplaceAll(key, ".", "-")
				fmt.Fprintf(buf, ".text-%s { color: %s; }\n", className, value)
				fmt.Fprintf(buf, ".bg-%s { background-color: %s; }\n", className, value)
				fmt.Fprintf(buf, ".border-%s { border-color: %s; }\n", className, value)
			}
		}

		// Radius utilities
		if tokens.Primitives.Radius != nil {
			for key, value := range tokens.Primitives.Radius {
				className := strings.ReplaceAll(key, ".", "-")
				fmt.Fprintf(buf, ".rounded-%s { border-radius: %s; }\n", className, value)
			}
		}
	}
	buf.WriteString("\n")
}

// writeColorSchemes generates additional color scheme variations
func (c *Compiler) writeColorSchemes(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	if !c.config.GenerateColorSchemes || tokens.Primitives == nil || tokens.Primitives.Colors == nil {
		return
	}

	if c.config.IncludeComments {
		buf.WriteString("/* Color Schemes */\n")
	}

	// Generate high contrast scheme
	buf.WriteString("@media (prefers-contrast: high) {\n")
	if c.config.GenerateThemeClasses {
		themeClass := c.makeThemeClassName(theme.ID)
		fmt.Fprintf(buf, "  .%s {\n", themeClass)
	} else {
		buf.WriteString("  :root {\n")
	}

	// Enhance contrast by adjusting colors
	if primary, exists := tokens.Primitives.Colors["primary.600"]; exists {
		fmt.Fprintf(buf, "    --color-primary-500: %s;\n", primary)
	}
	if foreground, exists := tokens.Primitives.Colors["gray.950"]; exists {
		fmt.Fprintf(buf, "    --color-foreground: %s;\n", foreground)
	}

	if c.config.GenerateThemeClasses {
		buf.WriteString("  }\n")
	} else {
		buf.WriteString("  }\n")
	}
	buf.WriteString("}\n\n")
}

// writeThemeExtensions writes custom CSS and accessibility styles
func (c *Compiler) writeThemeExtensions(buf *bytes.Buffer, theme *Theme) {
	// Custom CSS
	if theme.CustomCSS != "" {
		if c.config.IncludeComments {
			buf.WriteString("/* Custom CSS */\n")
		}
		buf.WriteString(theme.CustomCSS)
		buf.WriteString("\n")
	}

	// Accessibility
	if theme.Accessibility != nil {
		accessibilityCSS := c.generateAccessibilityCSS(theme.Accessibility)
		buf.WriteString(accessibilityCSS)
		buf.WriteString("\n")
	}
}

// generateCSSVariables converts design tokens to CSS custom properties
func (c *Compiler) generateCSSVariables(tokens *Tokens) string {
	var buf bytes.Buffer

	if tokens.Primitives != nil {
		c.writePrimitiveVariables(&buf, tokens.Primitives)
	}

	if tokens.Semantic != nil {
		c.writeSemanticVariables(&buf, tokens.Semantic)
	}

	return buf.String()
}

// writePrimitiveVariables writes CSS variables for primitive tokens
func (c *Compiler) writePrimitiveVariables(buf *bytes.Buffer, primitives *PrimitiveTokens) {
	categories := []struct {
		prefix  string
		tokens  map[string]string
		flatten bool
	}{
		{"color", primitives.Colors, c.config.FlattenNestedTokens},
		{"spacing", primitives.Spacing, false},
		{"radius", primitives.Radius, false},
		{"font", primitives.Typography, false},
		{"border", primitives.Borders, false},
		{"shadow", primitives.Shadows, false},
		{"effect", primitives.Effects, false},
		{"animation", primitives.Animation, false},
		{"z", primitives.ZIndex, false},
		{"breakpoint", primitives.Breakpoints, false},
	}

	for _, category := range categories {
		if category.tokens != nil {
			if category.flatten {
				c.writeFlattenedVariables(buf, category.prefix, category.tokens)
			} else {
				c.writeCategoryVariables(buf, category.prefix, category.tokens)
			}
		}
	}
}

// writeSemanticVariables writes CSS variables for semantic tokens
func (c *Compiler) writeSemanticVariables(buf *bytes.Buffer, semantic *SemanticTokens) {
	categories := map[string]map[string]string{
		"semantic-color":   semantic.Colors,
		"semantic-spacing": semantic.Spacing,
		"semantic-font":    semantic.Typography,
		"interactive":      semantic.Interactive,
	}

	for prefix, tokens := range categories {
		if tokens != nil {
			c.writeCategoryVariables(buf, prefix, tokens)
		}
	}
}

// writeFlattenedVariables handles nested token objects
func (c *Compiler) writeFlattenedVariables(buf *bytes.Buffer, prefix string, tokens map[string]string) {
	keys := c.sortedKeys(tokens)
	for _, key := range keys {
		varName := c.makeVariableName(prefix, key)
		fmt.Fprintf(buf, "  %s: %s;\n", varName, tokens[key])
	}
}

// writeCategoryVariables writes CSS variables for a token category
func (c *Compiler) writeCategoryVariables(buf *bytes.Buffer, prefix string, tokens map[string]string) {
	if len(tokens) == 0 {
		return
	}

	keys := c.sortedKeys(tokens)
	for _, key := range keys {
		varName := c.makeVariableName(prefix, key)
		fmt.Fprintf(buf, "  %s: %s;\n", varName, tokens[key])
	}
}

// generateCSSProperties generates CSS custom properties from theme configuration
func (c *Compiler) generateCSSProperties(tokens *Tokens, theme *Theme) string {
	var buf bytes.Buffer

	buf.WriteString("  /* CSS Properties */\n")

	// Generate core properties
	c.generateCoreProperties(&buf, tokens, theme)

	// Generate chart colors
	c.generateChartColors(&buf, tokens, theme)

	// Generate theme-specific properties
	c.generateThemeSpecificProperties(&buf, theme)

	return buf.String()
}

// generateCoreProperties generates core CSS properties
func (c *Compiler) generateCoreProperties(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	// Border radius
	if radius := c.getPreferredRadius(tokens, theme); radius != "" {
		fmt.Fprintf(buf, "  --radius: %s;\n", radius)
	}

	// Font families
	c.generateFontProperties(buf, tokens, theme)

	// Letter spacing
	if spacing := c.getPreferredLetterSpacing(tokens, theme); spacing != "" {
		fmt.Fprintf(buf, "  --tracking-normal: %s;\n", spacing)
	}

	// Shadow variables
	c.generateShadowProperties(buf, tokens, theme)
}

// generateFontProperties generates font-related CSS properties
func (c *Compiler) generateFontProperties(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	if tokens.Primitives == nil || tokens.Primitives.Typography == nil {
		return
	}

	preferredFont := c.getThemePreference(theme, "preferredFontFamily", "")

	if preferredFont != "" {
		fmt.Fprintf(buf, "  --font-sans: %s;\n", preferredFont)
		fmt.Fprintf(buf, "  --font-serif: %s;\n", preferredFont)
		buf.WriteString("  --font-mono: monospace;\n")
	} else {
		if font, exists := tokens.Primitives.Typography["fontFamily.sans"]; exists {
			fmt.Fprintf(buf, "  --font-sans: %s;\n", font)
		}
		if font, exists := tokens.Primitives.Typography["fontFamily.serif"]; exists {
			fmt.Fprintf(buf, "  --font-serif: %s;\n", font)
		}
		if font, exists := tokens.Primitives.Typography["fontFamily.mono"]; exists {
			fmt.Fprintf(buf, "  --font-mono: %s;\n", font)
		}
	}
}

// generateShadowProperties generates shadow CSS properties
func (c *Compiler) generateShadowProperties(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	if tokens.Primitives == nil || tokens.Primitives.Shadows == nil {
		return
	}

	shadowStyle := c.getThemePreference(theme, "shadowStyle", "")

	for key, value := range tokens.Primitives.Shadows {
		if shadowStyle == "" || strings.Contains(key, shadowStyle) || shadowStyle == "all" {
			varName := fmt.Sprintf("--shadow-%s", strings.ReplaceAll(key, ".", "-"))
			fmt.Fprintf(buf, "  %s: %s;\n", varName, value)
		}
	}
}

// generateChartColors generates chart color variables
func (c *Compiler) generateChartColors(buf *bytes.Buffer, tokens *Tokens, theme *Theme) {
	chartColors := map[string]string{
		"chart-1": "oklch(0.83 0.13 160.91)",
		"chart-2": "oklch(0.62 0.19 259.81)",
		"chart-3": "oklch(0.61 0.22 292.72)",
		"chart-4": "oklch(0.77 0.16 70.08)",
		"chart-5": "oklch(0.70 0.15 162.48)",
	}

	// Override with theme-specific chart colors
	if themeCharts := c.getThemeMap(theme, "chartColors"); themeCharts != nil {
		maps.Copy(chartColors, themeCharts)
	}

	// Use primary color from tokens
	if tokens.Primitives != nil && tokens.Primitives.Colors != nil {
		if primary := c.getPrimaryColor(tokens.Primitives.Colors); primary != "" {
			chartColors["chart-1"] = primary
		}
	}

	buf.WriteString("  /* Chart Colors */\n")
	for key, value := range chartColors {
		fmt.Fprintf(buf, "  --%s: %s;\n", key, value)
	}
}

// generateThemeSpecificProperties generates theme-specific CSS properties
func (c *Compiler) generateThemeSpecificProperties(buf *bytes.Buffer, theme *Theme) {
	if theme.Metadata == nil || theme.Metadata.CustomData == nil {
		return
	}

	buf.WriteString("  /* Theme Specific Properties */\n")

	// Common theme properties
	properties := map[string]string{
		"transition":   "transition-default",
		"backdropBlur": "backdrop-blur",
		"hoverScale":   "hover-scale",
	}

	for themeKey, cssProp := range properties {
		if value := c.getThemePreference(theme, themeKey, ""); value != "" {
			fmt.Fprintf(buf, "  --%s: %s;\n", cssProp, value)
		}
	}

	// Custom properties
	if customProps := c.getThemeMap(theme, "customProperties"); customProps != nil {
		for prop, value := range customProps {
			fmt.Fprintf(buf, "  --%s: %s;\n", prop, value)
		}
	}
}

// generateDarkModeCSSProperties generates dark mode specific CSS properties
func (c *Compiler) generateDarkModeCSSProperties(darkTokens *Tokens, theme *Theme) string {
	if darkTokens == nil {
		return ""
	}

	var buf bytes.Buffer
	buf.WriteString("  /* Dark Mode Properties */\n")

	// Dark mode chart colors
	if darkCharts := c.getThemeMap(theme, "darkChartColors"); darkCharts != nil {
		for key, value := range darkCharts {
			buf.WriteString(fmt.Sprintf("  --%s: %s;\n", key, value))
		}
	}

	// Use dark mode primary color
	if darkTokens.Primitives != nil && darkTokens.Primitives.Colors != nil {
		if primary := c.getDarkPrimaryColor(darkTokens.Primitives.Colors); primary != "" {
			buf.WriteString(fmt.Sprintf("  --chart-1: %s;\n", primary))
		}
	}

	// Dark mode specific properties
	if darkProps := c.getThemeMap(theme, "darkModeProperties"); darkProps != nil {
		for prop, value := range darkProps {
			buf.WriteString(fmt.Sprintf("  --%s: %s;\n", prop, value))
		}
	}

	return buf.String()
}

// Helper methods for theme configuration
func (c *Compiler) getPreferredRadius(tokens *Tokens, theme *Theme) string {
	preferred := c.getThemePreference(theme, "preferredRadius", "md")
	if tokens.Primitives != nil && tokens.Primitives.Radius != nil {
		if radius, exists := tokens.Primitives.Radius[preferred]; exists {
			return radius
		}
		if radius, exists := tokens.Primitives.Radius["md"]; exists {
			return radius
		}
	}
	return "0.5rem"
}

func (c *Compiler) getPreferredLetterSpacing(tokens *Tokens, theme *Theme) string {
	preferred := c.getThemePreference(theme, "letterSpacing", "")
	if preferred != "" {
		return preferred
	}
	if tokens.Primitives != nil && tokens.Primitives.Typography != nil {
		if spacing, exists := tokens.Primitives.Typography["letterSpacing.normal"]; exists {
			return spacing
		}
	}
	return "0.025em"
}

func (c *Compiler) getPrimaryColor(colors map[string]string) string {
	if color, exists := colors["primary.500"]; exists {
		return color
	}
	if color, exists := colors["primary.400"]; exists {
		return color
	}
	return ""
}

func (c *Compiler) getDarkPrimaryColor(colors map[string]string) string {
	if color, exists := colors["primary.400"]; exists {
		return color
	}
	if color, exists := colors["primary.300"]; exists {
		return color
	}
	return ""
}

func (c *Compiler) getThemePreference(theme *Theme, key, defaultValue string) string {
	if theme.Metadata == nil || theme.Metadata.CustomData == nil {
		return defaultValue
	}
	if value, ok := theme.Metadata.CustomData[key].(string); ok {
		return value
	}
	return defaultValue
}

func (c *Compiler) getThemeMap(theme *Theme, key string) map[string]string {
	if theme.Metadata == nil || theme.Metadata.CustomData == nil {
		return nil
	}
	if data, ok := theme.Metadata.CustomData[key].(map[string]any); ok {
		result := make(map[string]string)
		for k, v := range data {
			if str, ok := v.(string); ok {
				result[k] = str
			}
		}
		return result
	}
	return nil
}

// Utility methods
func (c *Compiler) sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func (c *Compiler) makeVariableName(prefix, key string) string {
	normalized := strings.ReplaceAll(key, "_", "-")
	normalized = strings.ReplaceAll(normalized, ".", "-")
	return fmt.Sprintf("%s-%s-%s", c.config.Prefix, prefix, normalized)
}

func (c *Compiler) makeThemeClassName(themeID string) string {
	return fmt.Sprintf("theme-%s", strings.ToLower(themeID))
}

func (c *Compiler) getDarkModeSelector(strategy string) string {
	switch strategy {
	case "class":
		return ".dark"
	case "media":
		return "@media (prefers-color-scheme: dark)"
	case "auto":
		return ".dark, @media (prefers-color-scheme: dark)"
	default:
		return ".dark"
	}
}

// generateAccessibilityCSS generates accessibility-related CSS
func (c *Compiler) generateAccessibilityCSS(config *AccessibilityConfig) string {
	var buf bytes.Buffer

	if c.config.IncludeComments {
		buf.WriteString("/* Accessibility */\n")
	}

	if config.FocusIndicator {
		outlineColor := config.FocusOutlineColor
		if outlineColor == "" {
			outlineColor = "#0066cc"
		}
		outlineWidth := config.FocusOutlineWidth
		if outlineWidth == "" {
			outlineWidth = "2px"
		}

		buf.WriteString("*:focus {\n")
		buf.WriteString(fmt.Sprintf("  outline: %s solid %s;\n", outlineWidth, outlineColor))
		buf.WriteString("  outline-offset: 2px;\n")
		buf.WriteString("}\n\n")

		buf.WriteString("*:focus:not(:focus-visible) {\n")
		buf.WriteString("  outline: none;\n")
		buf.WriteString("}\n\n")
	}

	if config.ReducedMotion {
		buf.WriteString("@media (prefers-reduced-motion: reduce) {\n")
		buf.WriteString("  *,\n")
		buf.WriteString("  *::before,\n")
		buf.WriteString("  *::after {\n")
		buf.WriteString("    animation-duration: 0.01ms !important;\n")
		buf.WriteString("    animation-iteration-count: 1 !important;\n")
		buf.WriteString("    transition-duration: 0.01ms !important;\n")
		buf.WriteString("  }\n")
		buf.WriteString("}\n\n")
	}

	if config.HighContrast {
		buf.WriteString("@media (prefers-contrast: high) {\n")
		buf.WriteString("  * {\n")
		buf.WriteString("    border-width: 2px;\n")
		buf.WriteString("  }\n")
		buf.WriteString("}\n\n")
	}

	return buf.String()
}

// optimizeCSS performs CSS optimizations including minification
func (c *Compiler) optimizeCSS(css string) string {
	if c.config.EnableMinify {
		css = c.minifyCSS(css)
	}

	// Remove duplicate rules (basic deduplication)
	css = c.deduplicateCSS(css)

	return css
}

// minifyCSS performs basic CSS minification
func (c *Compiler) minifyCSS(css string) string {
	// Remove comments
	css = regexp.MustCompile(`/\*.*?\*/`).ReplaceAllString(css, "")

	// Compress whitespace
	css = regexp.MustCompile(`\s+`).ReplaceAllString(css, " ")
	css = regexp.MustCompile(`\s*{\s*`).ReplaceAllString(css, "{")
	css = regexp.MustCompile(`\s*}\s*`).ReplaceAllString(css, "}")
	css = regexp.MustCompile(`\s*:\s*`).ReplaceAllString(css, ":")
	css = regexp.MustCompile(`\s*;\s*`).ReplaceAllString(css, ";")
	css = regexp.MustCompile(`;\s*}`).ReplaceAllString(css, "}")

	return strings.TrimSpace(css)
}

// deduplicateCSS removes duplicate CSS rules (basic implementation)
func (c *Compiler) deduplicateCSS(css string) string {
	// This is a basic implementation - in production you might want a more sophisticated approach
	lines := strings.Split(css, "\n")
	seen := make(map[string]bool)
	var result []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || seen[trimmed] {
			continue
		}
		seen[trimmed] = true
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

// Cache methods
func (c *Compiler) buildCacheKey(themeID string) string {
	return fmt.Sprintf("compiled:%s", themeID)
}

func (c *Compiler) getFromCache(key string) (string, bool) {
	if c.cache == nil {
		return "", false
	}
	value, found := c.cache.Get(key)
	if !found {
		return "", false
	}
	css, ok := value.(string)
	return css, ok
}

func (c *Compiler) setCache(key, css string) {
	if c.cache == nil {
		return
	}
	c.cache.SetWithTTL(key, css, int64(len(css)), c.cacheTTL)
}

// Stats methods
func (c *Compiler) recordStats(duration time.Duration, fromCache bool) {
	c.statsMutex.Lock()
	defer c.statsMutex.Unlock()

	c.stats.TotalCompilations++
	c.stats.TotalDuration += duration

	if fromCache {
		c.stats.CacheHits++
	} else {
		c.stats.CacheMisses++
	}

	if duration > c.stats.MaxDuration {
		c.stats.MaxDuration = duration
	}
	if c.stats.MinDuration == 0 || duration < c.stats.MinDuration {
		c.stats.MinDuration = duration
	}
}

func (c *Compiler) GetStats() *CompilerStats {
	c.statsMutex.Lock()
	defer c.statsMutex.Unlock()

	stats := *c.stats

	if c.cache != nil {
		metrics := c.cache.Metrics
		stats.CacheHits = metrics.Hits()
		stats.CacheMisses = metrics.Misses()
		stats.CacheHitRate = metrics.Ratio()
	}

	return &stats
}

func (c *Compiler) InvalidateCache() {
	if c.cache != nil {
		c.cache.Clear()
	}
}

// CompilerStats contains compiler performance statistics.
type CompilerStats struct {
	TotalCompilations uint64
	TotalDuration     time.Duration
	MaxDuration       time.Duration
	MinDuration       time.Duration
	CacheHits         uint64
	CacheMisses       uint64
	CacheHitRate      float64
}
