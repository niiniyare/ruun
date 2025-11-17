package schema

import (
	"sync"
	"time"
)

// Theme represents a complete theme configuration for the UI system.
// A theme is essentially a configured set of design tokens with metadata.
//
// Design Philosophy:
//   - Themes are immutable once created (use ThemeManager for customization)
//   - All styling flows through design tokens (single source of truth)
//   - Supports runtime customization via ThemeManager
//   - Multi-tenant capable with TenantThemeManager
//
// Breaking Changes: The Theme structure is designed to be stable, but new
// optional fields may be added for enhanced functionality.
type Theme struct {
	// Identity
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty" validate:"semver"`
	Author      string `json:"author,omitempty"`
	// Design tokens - the single source of truth for all styling
	Tokens *DesignTokens `json:"tokens" validate:"required"`
	// Dark mode configuration
	DarkMode *DarkModeConfig `json:"darkMode,omitempty"`
	// Accessibility configuration
	Accessibility *AccessibilityConfig `json:"accessibility,omitempty"`
	// Metadata
	Meta *ThemeMeta `json:"meta,omitempty"`
	// Custom extensions (advanced users only)
	// Breaking Changes: CustomCSS/CustomJS may be deprecated in favor of
	// a more structured extension system.
	CustomCSS string `json:"customCSS,omitempty"`
	CustomJS  string `json:"customJS,omitempty"`
	// Internal fields (not serialized)
	createdAt time.Time
	updatedAt time.Time
}

// DarkModeConfig defines dark mode behavior and color overrides.
type DarkModeConfig struct {
	// Enabled indicates if dark mode is available
	Enabled bool `json:"enabled"`
	// Default indicates if dark mode should be the default
	Default bool `json:"default,omitempty"`
	// Strategy defines how dark mode is detected/applied
	// Valid values: "class", "media", "auto"
	Strategy string `json:"strategy,omitempty" validate:"oneof=class media auto"`
	// DarkTokens are token overrides specifically for dark mode
	// These override the base tokens when dark mode is active
	DarkTokens *DesignTokens `json:"darkTokens,omitempty"`
}

// AccessibilityConfig defines accessibility settings and preferences.
// These settings help meet WCAG 2.1 Level AA compliance.
type AccessibilityConfig struct {
	// ARIA
	AutoARIA        bool   `json:"autoAria,omitempty"` // Auto-generate ARIA attributes
	AriaLive        string `json:"ariaLive,omitempty" validate:"oneof=off polite assertive"`
	AriaDescribedBy bool   `json:"ariaDescribedBy,omitempty"` // Auto-link descriptions
	// Keyboard navigation
	KeyboardNav        bool `json:"keyboardNav,omitempty"`        // Enable enhanced keyboard nav
	FocusIndicator     bool `json:"focusIndicator,omitempty"`     // Enhanced focus indicators
	SkipLinks          bool `json:"skipLinks,omitempty"`          // Add skip navigation links
	TabIndexManagement bool `json:"tabIndexManagement,omitempty"` // Automatic tabindex management
	// Screen reader optimizations
	ScreenReaderOnly  bool `json:"screenReaderOnly,omitempty"`  // Screen reader enhancements
	LiveAnnouncements bool `json:"liveAnnouncements,omitempty"` // Announce dynamic changes
	// Contrast and visibility
	HighContrast     bool    `json:"highContrast,omitempty"`                             // High contrast mode
	MinContrastRatio float64 `json:"minContrastRatio,omitempty" validate:"gte=0,lte=21"` // WCAG ratio
	// Focus styling
	FocusOutlineColor string `json:"focusOutlineColor,omitempty"`
	FocusOutlineWidth string `json:"focusOutlineWidth,omitempty"`
	// Motion preferences
	ReducedMotion bool `json:"reducedMotion,omitempty"` // Respect prefers-reduced-motion
}

// ThemeMeta contains theme metadata and organizational information.
type ThemeMeta struct {
	Tags       []string       `json:"tags,omitempty"`
	License    string         `json:"license,omitempty"`
	Repository string         `json:"repository,omitempty" validate:"url"`
	Homepage   string         `json:"homepage,omitempty" validate:"url"`
	Preview    string         `json:"preview,omitempty" validate:"url"` // Preview image URL
	CustomData map[string]any `json:"customData,omitempty"`
	CreatedAt  time.Time      `json:"createdAt,omitempty"`
	UpdatedAt  time.Time      `json:"updatedAt,omitempty"`
}

// ThemeOverrides represents customizations that can be applied to a base theme.
// This enables runtime theme customization without modifying the original theme.
type ThemeOverrides struct {
	// Token overrides - specific token values to override
	TokenOverrides map[string]string `json:"tokenOverrides,omitempty"`
	// Component customizations
	ComponentOverrides map[string]any `json:"componentOverrides,omitempty"`
	// Custom CSS to inject
	CustomCSS string `json:"customCSS,omitempty"`
	// Conditional overrides - apply changes based on context
	ConditionalOverrides map[string]*ConditionalOverride `json:"conditionalOverrides,omitempty"`
}

// ConditionalOverride defines overrides that are applied under specific conditions.
type ConditionalOverride struct {
	// Condition under which to apply the override
	Condition string `json:"condition" validate:"required"` // JS expression or predefined condition
	// Token overrides to apply when condition is true
	TokenOverrides map[string]string `json:"tokenOverrides,omitempty"`
	// Component overrides to apply when condition is true
	ComponentOverrides map[string]any `json:"componentOverrides,omitempty"`
	// Priority for conflict resolution (higher wins)
	Priority int `json:"priority,omitempty"`
}

// ThemeManager provides comprehensive theme management capabilities.
// It handles theme registration, customization, inheritance, and multi-tenant support.
//
// Core Features:
//   - Theme registration and lookup
//   - Runtime theme customization via overrides
//   - Theme inheritance and composition
//   - Multi-tenant theme isolation
//   - Caching and performance optimization
//   - Validation and error handling
//
// Usage:
//   manager := NewThemeManager()
//   manager.RegisterTheme(myTheme)
//   customized := manager.ApplyOverrides("base-theme", overrides)
type ThemeManager struct {
	// Core storage
	themes map[string]*Theme
	// Overrides storage
	overrides map[string]*ThemeOverrides
	// Theme inheritance
	inheritance map[string]string // child -> parent
	// Caching
	cache map[string]*Theme
	// Thread safety
	mu sync.RWMutex
	// Configuration
	config *ThemeManagerConfig
}

// ThemeManagerConfig configures the behavior of a ThemeManager.
type ThemeManagerConfig struct {
	// EnableCaching controls whether resolved themes are cached
	EnableCaching bool `json:"enableCaching"`
	// MaxCacheSize sets the maximum number of themes to cache
	MaxCacheSize int `json:"maxCacheSize"`
	// ValidateOnRegister controls whether themes are validated when registered
	ValidateOnRegister bool `json:"validateOnRegister"`
	// AllowOverrides controls whether theme overrides are allowed
	AllowOverrides bool `json:"allowOverrides"`
	// DefaultTheme is the fallback theme ID when a requested theme is not found
	DefaultTheme string `json:"defaultTheme,omitempty"`
}

// NewThemeManager creates a new ThemeManager with default configuration.
func NewThemeManager() *ThemeManager {
	return &ThemeManager{
		themes:      make(map[string]*Theme),
		overrides:   make(map[string]*ThemeOverrides),
		inheritance: make(map[string]string),
		cache:       make(map[string]*Theme),
		config: &ThemeManagerConfig{
			EnableCaching:      true,
			MaxCacheSize:       100,
			ValidateOnRegister: true,
			AllowOverrides:     true,
		},
	}
}

// RegisterTheme registers a theme with the manager.
func (tm *ThemeManager) RegisterTheme(theme *Theme) error {
	if theme == nil {
		return NewValidationError("nil_theme", "theme cannot be nil")
	}

	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Validate theme if configured
	if tm.config.ValidateOnRegister {
		if err := tm.validateTheme(theme); err != nil {
			return err
		}
	}

	// Update timestamps
	now := time.Now()
	theme.createdAt = now
	theme.updatedAt = now

	// Store theme
	tm.themes[theme.ID] = theme

	// Clear cache since we have a new theme
	if tm.config.EnableCaching {
		tm.clearCache()
	}

	return nil
}

// GetTheme retrieves a theme by ID, applying any registered overrides.
func (tm *ThemeManager) GetTheme(themeID string) (*Theme, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	// Check cache first
	if tm.config.EnableCaching {
		if cached, exists := tm.cache[themeID]; exists {
			return tm.cloneTheme(cached), nil
		}
	}

	// Get base theme
	baseTheme, exists := tm.themes[themeID]
	if !exists {
		// Try default theme if configured
		if tm.config.DefaultTheme != "" && tm.config.DefaultTheme != themeID {
			return tm.GetTheme(tm.config.DefaultTheme)
		}
		return nil, NewValidationError("theme_not_found", "theme not found: "+themeID)
	}

	// Apply inheritance if configured
	resolvedTheme := tm.resolveInheritance(baseTheme)

	// Apply overrides if they exist
	if overrides, hasOverrides := tm.overrides[themeID]; hasOverrides && tm.config.AllowOverrides {
		resolvedTheme = tm.applyOverridesToTheme(resolvedTheme, overrides)
	}

	// Cache the resolved theme
	if tm.config.EnableCaching {
		tm.cache[themeID] = tm.cloneTheme(resolvedTheme)
	}

	return resolvedTheme, nil
}

// ListThemes returns all registered themes.
func (tm *ThemeManager) ListThemes() []*Theme {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	themes := make([]*Theme, 0, len(tm.themes))
	for _, theme := range tm.themes {
		themes = append(themes, tm.cloneTheme(theme))
	}
	return themes
}

// UnregisterTheme removes a theme from the manager.
func (tm *ThemeManager) UnregisterTheme(themeID string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.themes[themeID]; !exists {
		return NewValidationError("theme_not_found", "theme not found: "+themeID)
	}

	// Remove theme
	delete(tm.themes, themeID)
	// Remove any overrides
	delete(tm.overrides, themeID)
	// Remove inheritance relationships
	delete(tm.inheritance, themeID)
	// Clear from cache
	delete(tm.cache, themeID)

	return nil
}

// ApplyOverrides applies runtime overrides to a theme without modifying the original.
func (tm *ThemeManager) ApplyOverrides(baseTheme *Theme, overrides *ThemeOverrides) *Theme {
	if overrides == nil {
		return tm.cloneTheme(baseTheme)
	}

	return tm.applyOverridesToTheme(baseTheme, overrides)
}

// RegisterOverrides registers persistent overrides for a theme.
func (tm *ThemeManager) RegisterOverrides(themeID string, overrides *ThemeOverrides) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if !tm.config.AllowOverrides {
		return NewValidationError("overrides_disabled", "theme overrides are disabled")
	}

	tm.overrides[themeID] = overrides

	// Clear cache for this theme
	delete(tm.cache, themeID)

	return nil
}

// SetInheritance establishes a parent-child relationship between themes.
func (tm *ThemeManager) SetInheritance(childThemeID, parentThemeID string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Validate both themes exist
	if _, exists := tm.themes[childThemeID]; !exists {
		return NewValidationError("child_theme_not_found", "child theme not found: "+childThemeID)
	}
	if _, exists := tm.themes[parentThemeID]; !exists {
		return NewValidationError("parent_theme_not_found", "parent theme not found: "+parentThemeID)
	}

	// Check for circular inheritance
	if tm.wouldCreateCircle(childThemeID, parentThemeID) {
		return NewValidationError("circular_inheritance", "inheritance would create circular dependency")
	}

	tm.inheritance[childThemeID] = parentThemeID

	// Clear cache for affected themes
	tm.clearInheritanceCache(childThemeID)

	return nil
}

// Private helper methods

func (tm *ThemeManager) validateTheme(theme *Theme) error {
	if theme.ID == "" {
		return NewValidationError("empty_theme_id", "theme ID cannot be empty")
	}
	if theme.Name == "" {
		return NewValidationError("empty_theme_name", "theme name cannot be empty")
	}
	if theme.Tokens == nil {
		return NewValidationError("missing_tokens", "theme must have tokens")
	}
	return nil
}

func (tm *ThemeManager) cloneTheme(theme *Theme) *Theme {
	// Deep clone theme to prevent mutations
	cloned := *theme
	// TODO: Deep clone tokens and other nested structures
	return &cloned
}

func (tm *ThemeManager) clearCache() {
	for k := range tm.cache {
		delete(tm.cache, k)
	}
}

func (tm *ThemeManager) clearInheritanceCache(themeID string) {
	// Clear cache for theme and all its children
	delete(tm.cache, themeID)
	for child, parent := range tm.inheritance {
		if parent == themeID {
			tm.clearInheritanceCache(child)
		}
	}
}

func (tm *ThemeManager) resolveInheritance(theme *Theme) *Theme {
	parentID, hasParent := tm.inheritance[theme.ID]
	if !hasParent {
		return theme
	}

	parent, exists := tm.themes[parentID]
	if !exists {
		return theme
	}

	// Recursively resolve parent
	resolvedParent := tm.resolveInheritance(parent)

	// Merge parent with child (child takes precedence)
	return tm.mergeThemes(resolvedParent, theme)
}

func (tm *ThemeManager) mergeThemes(parent, child *Theme) *Theme {
	// Simple merge - child overrides parent
	merged := tm.cloneTheme(parent)
	merged.ID = child.ID
	merged.Name = child.Name
	if child.Description != "" {
		merged.Description = child.Description
	}
	if child.Version != "" {
		merged.Version = child.Version
	}
	if child.Author != "" {
		merged.Author = child.Author
	}
	if child.Tokens != nil {
		merged.Tokens = child.Tokens
	}
	if child.DarkMode != nil {
		merged.DarkMode = child.DarkMode
	}
	if child.Accessibility != nil {
		merged.Accessibility = child.Accessibility
	}
	if child.Meta != nil {
		merged.Meta = child.Meta
	}
	if child.CustomCSS != "" {
		merged.CustomCSS = child.CustomCSS
	}
	if child.CustomJS != "" {
		merged.CustomJS = child.CustomJS
	}
	return merged
}

func (tm *ThemeManager) applyOverridesToTheme(theme *Theme, overrides *ThemeOverrides) *Theme {
	customized := tm.cloneTheme(theme)
	
	// Apply token overrides
	if len(overrides.TokenOverrides) > 0 {
		// TODO: Deep apply token overrides
		// This would require token resolution and application
	}

	// Apply custom CSS
	if overrides.CustomCSS != "" {
		if customized.CustomCSS != "" {
			customized.CustomCSS += "\n" + overrides.CustomCSS
		} else {
			customized.CustomCSS = overrides.CustomCSS
		}
	}

	return customized
}

func (tm *ThemeManager) wouldCreateCircle(childID, parentID string) bool {
	current := parentID
	visited := make(map[string]bool)

	for {
		if visited[current] {
			return true // Found a cycle
		}
		if current == childID {
			return true // Would create direct cycle
		}

		visited[current] = true
		next, hasParent := tm.inheritance[current]
		if !hasParent {
			break
		}
		current = next
	}

	return false
}

// Legacy compatibility structures and functions
// Kept for backward compatibility with existing theme.go usage

type ThemeRegistry struct {
	manager *ThemeManager
}

func NewThemeRegistry() *ThemeRegistry {
	return &ThemeRegistry{
		manager: NewThemeManager(),
	}
}

// TokenRegistry wrapper for compatibility
type TokenRegistry struct {
	registry TokenRegistryInterface
}

func NewTokenRegistry() *TokenRegistry {
	return &TokenRegistry{
		registry: NewSimpleTokenRegistry(),
	}
}

// NewThemeManagerWithRegistries with legacy signature for compatibility
func NewThemeManagerWithRegistries(themeRegistry *ThemeRegistry, tokenRegistry *TokenRegistry) *ThemeManager {
	if themeRegistry != nil {
		return themeRegistry.manager
	}
	return NewThemeManager()
}

// TenantThemeManager provides multi-tenant theme management.
type TenantThemeManager struct {
	tenantManagers map[string]*ThemeManager
	globalManager  *ThemeManager
	mu             sync.RWMutex
}

// NewTenantThemeManager creates a new tenant-aware theme manager.
func NewTenantThemeManager() *TenantThemeManager {
	return &TenantThemeManager{
		tenantManagers: make(map[string]*ThemeManager),
		globalManager:  NewThemeManager(),
	}
}

// GetManagerForTenant returns a theme manager for a specific tenant.
func (ttm *TenantThemeManager) GetManagerForTenant(tenantID string) *ThemeManager {
	ttm.mu.RLock()
	manager, exists := ttm.tenantManagers[tenantID]
	ttm.mu.RUnlock()

	if exists {
		return manager
	}

	ttm.mu.Lock()
	defer ttm.mu.Unlock()

	// Double-check after acquiring write lock
	if manager, exists := ttm.tenantManagers[tenantID]; exists {
		return manager
	}

	// Create new manager for tenant
	manager = NewThemeManager()
	ttm.tenantManagers[tenantID] = manager
	return manager
}

// GetGlobalManager returns the global theme manager.
func (ttm *TenantThemeManager) GetGlobalManager() *ThemeManager {
	return ttm.globalManager
}