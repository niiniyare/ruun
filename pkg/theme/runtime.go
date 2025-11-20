package theme

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// Runtime provides theme-aware component utilities with caching and performance optimizations
type Runtime struct {
	api           *ThemeAPIInterface
	resolver      *TokenResolver
	cache         *RuntimeCache
	config        *RuntimeConfig
	currentTheme  *schema.Theme
	darkMode      bool
	tenantID      string
	mu            sync.RWMutex
}

// RuntimeConfig configures the theme runtime behavior
type RuntimeConfig struct {
	EnableCaching     bool          `json:"enableCaching"`
	CacheTTL          time.Duration `json:"cacheTTL"`
	EnableDarkMode    bool          `json:"enableDarkMode"`
	PreloadThemes     []string      `json:"preloadThemes"`
	FallbackTheme     string        `json:"fallbackTheme"`
	AutoDarkMode      bool          `json:"autoDarkMode"`
	ValidateTokens    bool          `json:"validateTokens"`
	MaxCacheSize      int           `json:"maxCacheSize"`
	EnableHotReload   bool          `json:"enableHotReload"`
}

// RuntimeCache provides fast access to resolved theme values
type RuntimeCache struct {
	tokenValues map[string]string
	classNames  map[string]string
	compiled    map[string]*CompiledTheme
	mu          sync.RWMutex
	lastUpdate  time.Time
}

// CompiledTheme represents a theme compiled for runtime use
type CompiledTheme struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	CSS         string                 `json:"css"`
	Variables   map[string]string      `json:"variables"`
	Classes     map[string]string      `json:"classes"`
	Components  map[string]interface{} `json:"components"`
	Timestamp   time.Time              `json:"timestamp"`
	Hash        string                 `json:"hash"`
	DarkMode    *CompiledTheme         `json:"darkMode,omitempty"`
}

// ThemeAPIInterface abstracts the theme API for testing
type ThemeAPIInterface interface {
	GetTheme(themeID string) (*schema.Theme, error)
	CompileTheme(theme *schema.Theme) (string, error)
	GetCompiledCSS(themeID string) (string, error)
	ListThemes() []*schema.Theme
}

// NewRuntime creates a new theme runtime with the specified configuration
func NewRuntime(api ThemeAPIInterface, config *RuntimeConfig) *Runtime {
	if config == nil {
		config = DefaultRuntimeConfig()
	}

	cache := &RuntimeCache{
		tokenValues: make(map[string]string),
		classNames:  make(map[string]string),
		compiled:    make(map[string]*CompiledTheme),
		lastUpdate:  time.Now(),
	}

	runtime := &Runtime{
		api:      &api,
		cache:    cache,
		config:   config,
		resolver: NewTokenResolver(api, cache),
	}

	// Preload themes if configured
	if len(config.PreloadThemes) > 0 {
		go runtime.preloadThemes()
	}

	return runtime
}

// DefaultRuntimeConfig returns default configuration
func DefaultRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		EnableCaching:   true,
		CacheTTL:        30 * time.Minute,
		EnableDarkMode:  true,
		AutoDarkMode:    true,
		ValidateTokens:  true,
		MaxCacheSize:    1000,
		EnableHotReload: true,
		FallbackTheme:   "default",
	}
}

// SetTheme switches to a new theme
func (r *Runtime) SetTheme(themeID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	theme, err := (*r.api).GetTheme(themeID)
	if err != nil {
		// Try fallback theme
		if r.config.FallbackTheme != "" && r.config.FallbackTheme != themeID {
			theme, err = (*r.api).GetTheme(r.config.FallbackTheme)
			if err != nil {
				return fmt.Errorf("failed to load theme %s and fallback: %w", themeID, err)
			}
		} else {
			return fmt.Errorf("failed to load theme %s: %w", themeID, err)
		}
	}

	r.currentTheme = theme
	
	// Clear cache when switching themes
	if r.config.EnableCaching {
		r.clearCache()
	}

	// Compile theme for runtime
	if err := r.compileCurrentTheme(); err != nil {
		return fmt.Errorf("failed to compile theme: %w", err)
	}

	return nil
}

// SetDarkMode toggles dark mode
func (r *Runtime) SetDarkMode(enabled bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.config.EnableDarkMode {
		return fmt.Errorf("dark mode is disabled")
	}

	r.darkMode = enabled
	
	// Clear cache when toggling dark mode
	if r.config.EnableCaching {
		r.clearCache()
	}

	return nil
}

// SetTenant sets the current tenant for multi-tenant support
func (r *Runtime) SetTenant(tenantID string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.tenantID = tenantID
}

// ResolveToken resolves a token reference to its final value
func (r *Runtime) ResolveToken(ctx context.Context, tokenRef schema.TokenReference) (string, error) {
	return r.resolver.ResolveToken(ctx, tokenRef, r.currentTheme, r.darkMode)
}

// GetCSSVariable returns a CSS custom property for a token
func (r *Runtime) GetCSSVariable(tokenPath string) string {
	// Convert token path to CSS variable name
	// "colors.primary" -> "--color-primary"
	parts := strings.Split(tokenPath, ".")
	if len(parts) >= 2 {
		category := parts[0]
		name := strings.Join(parts[1:], "-")
		return fmt.Sprintf("--color-%s", name)
	}
	return fmt.Sprintf("--%s", strings.ReplaceAll(tokenPath, ".", "-"))
}

// GetCompiledClass returns a compiled CSS class name for a component variant
func (r *Runtime) GetCompiledClass(component, variant string) string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.config.EnableCaching {
		key := fmt.Sprintf("%s-%s", component, variant)
		if cached, exists := r.cache.classNames[key]; exists {
			return cached
		}
	}

	// Generate class name: "button-primary", "input-error", etc.
	className := fmt.Sprintf("%s-%s", component, variant)

	if r.config.EnableCaching {
		r.cache.mu.Lock()
		r.cache.classNames[fmt.Sprintf("%s-%s", component, variant)] = className
		r.cache.mu.Unlock()
	}

	return className
}

// GetComponentCSS returns the compiled CSS for a component
func (r *Runtime) GetComponentCSS(componentName string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.currentTheme == nil {
		return "", fmt.Errorf("no theme set")
	}

	compiled := r.getCompiledTheme(r.currentTheme.ID)
	if compiled == nil {
		return "", fmt.Errorf("theme not compiled")
	}

	if css, exists := compiled.Components[componentName]; exists {
		if cssStr, ok := css.(string); ok {
			return cssStr, nil
		}
	}

	return "", fmt.Errorf("component CSS not found: %s", componentName)
}

// GetThemeCSS returns the complete CSS for the current theme
func (r *Runtime) GetThemeCSS() (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.currentTheme == nil {
		return "", fmt.Errorf("no theme set")
	}

	compiled := r.getCompiledTheme(r.currentTheme.ID)
	if compiled == nil {
		return "", fmt.Errorf("theme not compiled")
	}

	css := compiled.CSS

	// Append dark mode CSS if enabled
	if r.darkMode && compiled.DarkMode != nil {
		css += "\n@media (prefers-color-scheme: dark) {\n"
		css += compiled.DarkMode.CSS
		css += "\n}"
	}

	return css, nil
}

// GetCurrentTheme returns the current theme information
func (r *Runtime) GetCurrentTheme() *ThemeInfo {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.currentTheme == nil {
		return nil
	}

	return &ThemeInfo{
		ID:          r.currentTheme.ID,
		Name:        r.currentTheme.Name,
		Description: r.currentTheme.Description,
		Version:     r.currentTheme.Version,
		DarkMode:    r.darkMode,
		TenantID:    r.tenantID,
	}
}

// ThemeInfo provides basic theme information
type ThemeInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	DarkMode    bool   `json:"darkMode"`
	TenantID    string `json:"tenantId"`
}

// ValidateTheme validates a theme configuration
func (r *Runtime) ValidateTheme(themeID string) (*ValidationResult, error) {
	theme, err := (*r.api).GetTheme(themeID)
	if err != nil {
		return nil, err
	}

	validator := NewThemeValidator()
	return validator.ValidateTheme(theme)
}

// GetAvailableThemes returns all available themes
func (r *Runtime) GetAvailableThemes() []*ThemeInfo {
	themes := (*r.api).ListThemes()
	result := make([]*ThemeInfo, len(themes))
	
	for i, theme := range themes {
		result[i] = &ThemeInfo{
			ID:          theme.ID,
			Name:        theme.Name,
			Description: theme.Description,
			Version:     theme.Version,
		}
	}
	
	return result
}

// ClearCache clears the runtime cache
func (r *Runtime) ClearCache() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.clearCache()
}

// Private helper methods

func (r *Runtime) clearCache() {
	r.cache.mu.Lock()
	defer r.cache.mu.Unlock()
	
	r.cache.tokenValues = make(map[string]string)
	r.cache.classNames = make(map[string]string)
	r.cache.lastUpdate = time.Now()
}

func (r *Runtime) compileCurrentTheme() error {
	if r.currentTheme == nil {
		return fmt.Errorf("no theme to compile")
	}

	css, err := (*r.api).CompileTheme(r.currentTheme)
	if err != nil {
		return err
	}

	compiled := &CompiledTheme{
		ID:        r.currentTheme.ID,
		Name:      r.currentTheme.Name,
		CSS:       css,
		Variables: r.extractCSSVariables(css),
		Classes:   r.extractCSSClasses(css),
		Timestamp: time.Now(),
		Hash:      r.calculateHash(css),
	}

	// Compile dark mode if enabled
	if r.config.EnableDarkMode && r.currentTheme.DarkMode != nil && r.currentTheme.DarkMode.Enabled {
		darkCSS := r.compileDarkMode(r.currentTheme)
		compiled.DarkMode = &CompiledTheme{
			ID:        r.currentTheme.ID + "-dark",
			CSS:       darkCSS,
			Variables: r.extractCSSVariables(darkCSS),
			Classes:   r.extractCSSClasses(darkCSS),
			Timestamp: time.Now(),
		}
	}

	// Cache compiled theme
	if r.config.EnableCaching {
		r.cache.mu.Lock()
		r.cache.compiled[r.currentTheme.ID] = compiled
		r.cache.mu.Unlock()
	}

	return nil
}

func (r *Runtime) getCompiledTheme(themeID string) *CompiledTheme {
	if !r.config.EnableCaching {
		return nil
	}

	r.cache.mu.RLock()
	defer r.cache.mu.RUnlock()
	return r.cache.compiled[themeID]
}

func (r *Runtime) preloadThemes() {
	for _, themeID := range r.config.PreloadThemes {
		if err := r.SetTheme(themeID); err == nil {
			// Theme preloaded and compiled
			continue
		}
	}
}

func (r *Runtime) extractCSSVariables(css string) map[string]string {
	variables := make(map[string]string)
	// Simple extraction - in a real implementation, use a CSS parser
	lines := strings.Split(css, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "--") && strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(strings.TrimSuffix(parts[1], ";"))
				variables[name] = value
			}
		}
	}
	return variables
}

func (r *Runtime) extractCSSClasses(css string) map[string]string {
	classes := make(map[string]string)
	// Simple extraction - in a real implementation, use a CSS parser
	lines := strings.Split(css, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ".") && strings.Contains(line, "{") {
			className := strings.TrimSpace(strings.Split(line, "{")[0])
			className = strings.TrimPrefix(className, ".")
			classes[className] = className
		}
	}
	return classes
}

func (r *Runtime) compileDarkMode(theme *schema.Theme) string {
	// Simple dark mode compilation - override variables
	css := ":root.dark {\n"
	if theme.DarkMode != nil && theme.DarkMode.DarkTokens != nil {
		// Add dark mode variables
		css += "  /* Dark mode overrides */\n"
		// This would be implemented with proper token resolution
	}
	css += "}\n"
	return css
}

func (r *Runtime) calculateHash(content string) string {
	// Simple hash - in a real implementation, use crypto/sha256
	return fmt.Sprintf("hash-%d", len(content))
}