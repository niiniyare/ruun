package theme

import (
	"bytes"
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

// Compiler compiles design tokens into CSS with optimization and caching.
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
	EnableCaching    bool
	EnableMinify     bool
	EnableSourceMap  bool
	EnableVariables  bool
	Prefix           string
	CacheTTL         time.Duration
	CacheSize        int64
	IncludeComments  bool
	GenerateUtilities bool
}

// DefaultCompilerConfig returns production-ready compiler configuration.
func DefaultCompilerConfig() *CompilerConfig {
	return &CompilerConfig{
		EnableCaching:     true,
		EnableMinify:      true,
		EnableSourceMap:   false,
		EnableVariables:   true,
		Prefix:            "--awo",
		CacheTTL:          10 * time.Minute,
		CacheSize:         20 << 20, // 20MB
		IncludeComments:   false,
		GenerateUtilities: false,
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

// Compile compiles design tokens into CSS.
func (c *Compiler) Compile(ctx context.Context, tokens *Tokens, theme *Theme) (string, error) {
	start := time.Now()

	cacheKey := c.buildCacheKey(theme.ID)
	if cached, found := c.getFromCache(cacheKey); found {
		c.recordStats(time.Since(start), true)
		return cached, nil
	}

	var buf bytes.Buffer

	if c.config.IncludeComments {
		buf.WriteString(fmt.Sprintf("/* Theme: %s */\n", theme.Name))
		buf.WriteString(fmt.Sprintf("/* Generated: %s */\n\n", time.Now().Format(time.RFC3339)))
	}

	if c.config.EnableVariables {
		cssVars := c.generateCSSVariables(tokens)
		buf.WriteString(":root {\n")
		buf.WriteString(cssVars)
		buf.WriteString("}\n\n")

		if theme.DarkMode != nil && theme.DarkMode.Enabled && theme.DarkMode.DarkTokens != nil {
			darkVars := c.generateCSSVariables(theme.DarkMode.DarkTokens)
			
			selector := c.getDarkModeSelector(theme.DarkMode.Strategy)
			buf.WriteString(fmt.Sprintf("%s {\n", selector))
			buf.WriteString(darkVars)
			buf.WriteString("}\n\n")
		}
	}

	if c.config.GenerateUtilities {
		utilities := c.generateUtilityClasses(tokens)
		buf.WriteString(utilities)
		buf.WriteString("\n")
	}

	if theme.CustomCSS != "" {
		if c.config.IncludeComments {
			buf.WriteString("/* Custom CSS */\n")
		}
		buf.WriteString(theme.CustomCSS)
		buf.WriteString("\n")
	}

	if theme.Accessibility != nil {
		accessibilityCSS := c.generateAccessibilityCSS(theme.Accessibility)
		buf.WriteString(accessibilityCSS)
		buf.WriteString("\n")
	}

	css := buf.String()

	if c.config.EnableMinify {
		css = c.minifyCSS(css)
	}

	c.setCache(cacheKey, css)
	c.recordStats(time.Since(start), false)

	return css, nil
}

// generateCSSVariables converts design tokens to CSS custom properties.
func (c *Compiler) generateCSSVariables(tokens *Tokens) string {
	var buf bytes.Buffer

	if tokens.Primitives != nil {
		c.writeCategoryVariables(&buf, "color", tokens.Primitives.Colors)
		c.writeCategoryVariables(&buf, "spacing", tokens.Primitives.Spacing)
		c.writeCategoryVariables(&buf, "radius", tokens.Primitives.Radius)
		c.writeCategoryVariables(&buf, "font", tokens.Primitives.Typography)
		c.writeCategoryVariables(&buf, "border", tokens.Primitives.Borders)
		c.writeCategoryVariables(&buf, "shadow", tokens.Primitives.Shadows)
		c.writeCategoryVariables(&buf, "effect", tokens.Primitives.Effects)
		c.writeCategoryVariables(&buf, "animation", tokens.Primitives.Animation)
		c.writeCategoryVariables(&buf, "z", tokens.Primitives.ZIndex)
		c.writeCategoryVariables(&buf, "breakpoint", tokens.Primitives.Breakpoints)
	}

	if tokens.Semantic != nil {
		c.writeCategoryVariables(&buf, "semantic-color", tokens.Semantic.Colors)
		c.writeCategoryVariables(&buf, "semantic-spacing", tokens.Semantic.Spacing)
		c.writeCategoryVariables(&buf, "semantic-font", tokens.Semantic.Typography)
		c.writeCategoryVariables(&buf, "interactive", tokens.Semantic.Interactive)
	}

	return buf.String()
}

// writeCategoryVariables writes CSS variables for a token category.
func (c *Compiler) writeCategoryVariables(buf *bytes.Buffer, prefix string, tokens map[string]string) {
	if tokens == nil || len(tokens) == 0 {
		return
	}

	keys := make([]string, 0, len(tokens))
	for key := range tokens {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		varName := c.makeVariableName(prefix, key)
		value := tokens[key]
		buf.WriteString(fmt.Sprintf("  %s: %s;\n", varName, value))
	}
}

// makeVariableName creates a CSS variable name with prefix.
func (c *Compiler) makeVariableName(prefix, key string) string {
	normalized := strings.ReplaceAll(key, "_", "-")
	normalized = strings.ReplaceAll(normalized, ".", "-")
	return fmt.Sprintf("%s-%s-%s", c.config.Prefix, prefix, normalized)
}

// getDarkModeSelector returns the appropriate CSS selector for dark mode.
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

// generateUtilityClasses generates utility CSS classes from tokens.
func (c *Compiler) generateUtilityClasses(tokens *Tokens) string {
	var buf bytes.Buffer

	if c.config.IncludeComments {
		buf.WriteString("/* Utility Classes */\n")
	}

	if tokens.Primitives != nil && tokens.Primitives.Spacing != nil {
		for key, value := range tokens.Primitives.Spacing {
			className := strings.ReplaceAll(key, ".", "-")
			buf.WriteString(fmt.Sprintf(".m-%s { margin: %s; }\n", className, value))
			buf.WriteString(fmt.Sprintf(".p-%s { padding: %s; }\n", className, value))
		}
	}

	if tokens.Primitives != nil && tokens.Primitives.Colors != nil {
		for key, value := range tokens.Primitives.Colors {
			className := strings.ReplaceAll(key, ".", "-")
			buf.WriteString(fmt.Sprintf(".text-%s { color: %s; }\n", className, value))
			buf.WriteString(fmt.Sprintf(".bg-%s { background-color: %s; }\n", className, value))
		}
	}

	return buf.String()
}

// generateAccessibilityCSS generates accessibility-related CSS.
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

// minifyCSS performs basic CSS minification.
func (c *Compiler) minifyCSS(css string) string {
	css = regexp.MustCompile(`/\*.*?\*/`).ReplaceAllString(css, "")
	css = regexp.MustCompile(`\s+`).ReplaceAllString(css, " ")
	css = regexp.MustCompile(`\s*{\s*`).ReplaceAllString(css, "{")
	css = regexp.MustCompile(`\s*}\s*`).ReplaceAllString(css, "}")
	css = regexp.MustCompile(`\s*:\s*`).ReplaceAllString(css, ":")
	css = regexp.MustCompile(`\s*;\s*`).ReplaceAllString(css, ";")
	css = regexp.MustCompile(`;\s*}`).ReplaceAllString(css, "}")
	css = strings.TrimSpace(css)
	return css
}

// buildCacheKey builds a cache key for compilation.
func (c *Compiler) buildCacheKey(themeID string) string {
	return fmt.Sprintf("compiled:%s", themeID)
}

// getFromCache retrieves compiled CSS from cache.
func (c *Compiler) getFromCache(key string) (string, bool) {
	if c.cache == nil {
		return "", false
	}

	value, found := c.cache.Get(key)
	if !found {
		return "", false
	}

	css, ok := value.(string)
	if !ok {
		return "", false
	}

	return css, true
}

// setCache stores compiled CSS in cache.
func (c *Compiler) setCache(key, css string) {
	if c.cache == nil {
		return
	}

	c.cache.SetWithTTL(key, css, int64(len(css)), c.cacheTTL)
}

// recordStats records compilation statistics.
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

// GetStats returns compiler performance statistics.
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

// InvalidateCache clears the compiler cache.
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
