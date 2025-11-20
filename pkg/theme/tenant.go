package theme

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// TenantThemeManager provides multi-tenant theme management capabilities
type TenantThemeManager struct {
	globalManager  *Runtime
	tenantRuntimes map[string]*Runtime
	tenantConfigs  map[string]*TenantConfig
	tenantStorage  TenantStorage
	config         *MultiTenantConfig
	mu             sync.RWMutex
}

// MultiTenantConfig configures multi-tenant theme behavior
type MultiTenantConfig struct {
	EnableTenantIsolation    bool          `json:"enableTenantIsolation"`
	EnableTenantInheritance  bool          `json:"enableTenantInheritance"`
	DefaultTenantTheme       string        `json:"defaultTenantTheme"`
	GlobalFallbackTheme      string        `json:"globalFallbackTheme"`
	TenantCacheSize          int           `json:"tenantCacheSize"`
	TenantCacheTTL           time.Duration `json:"tenantCacheTTL"`
	EnableTenantCustomization bool         `json:"enableTenantCustomization"`
	MaxTenantsPerNode        int           `json:"maxTenantsPerNode"`
	TenantConfigPath         string        `json:"tenantConfigPath"`
}

// TenantConfig represents theme configuration for a specific tenant
type TenantConfig struct {
	TenantID            string                 `json:"tenantId"`
	ThemeID             string                 `json:"themeId"`
	CustomThemes        []string               `json:"customThemes,omitempty"`
	AllowedThemes       []string               `json:"allowedThemes,omitempty"`
	BlockedThemes       []string               `json:"blockedThemes,omitempty"`
	DefaultTheme        string                 `json:"defaultTheme"`
	EnableDarkMode      bool                   `json:"enableDarkMode"`
	CustomTokens        map[string]string      `json:"customTokens,omitempty"`
	BrandingOverrides   *BrandingOverrides     `json:"brandingOverrides,omitempty"`
	FeatureFlags        map[string]bool        `json:"featureFlags,omitempty"`
	Metadata            map[string]any `json:"metadata,omitempty"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
	Active              bool                   `json:"active"`
}

// BrandingOverrides allows tenant-specific branding customizations
type BrandingOverrides struct {
	PrimaryColor     string            `json:"primaryColor,omitempty"`
	SecondaryColor   string            `json:"secondaryColor,omitempty"`
	AccentColor      string            `json:"accentColor,omitempty"`
	LogoURL          string            `json:"logoUrl,omitempty"`
	FaviconURL       string            `json:"faviconUrl,omitempty"`
	FontFamily       string            `json:"fontFamily,omitempty"`
	CustomCSS        string            `json:"customCss,omitempty"`
	ColorPalette     map[string]string `json:"colorPalette,omitempty"`
	TypographyScale  map[string]string `json:"typographyScale,omitempty"`
	SpacingScale     map[string]string `json:"spacingScale,omitempty"`
}

// TenantStorage interface for persisting tenant theme configurations
type TenantStorage interface {
	SaveTenantConfig(tenantID string, config *TenantConfig) error
	LoadTenantConfig(tenantID string) (*TenantConfig, error)
	DeleteTenantConfig(tenantID string) error
	ListTenantConfigs() ([]TenantConfig, error)
	UpdateTenantConfig(tenantID string, updates map[string]any) error
}

// TenantThemeContext provides context for tenant-specific theme operations
type TenantThemeContext struct {
	TenantID     string                 `json:"tenantId"`
	UserID       string                 `json:"userId,omitempty"`
	ThemeID      string                 `json:"themeId"`
	DarkMode     bool                   `json:"darkMode"`
	CustomTokens map[string]string      `json:"customTokens,omitempty"`
	Permissions  *TenantPermissions     `json:"permissions,omitempty"`
	Metadata     map[string]any `json:"metadata,omitempty"`
}

// TenantPermissions defines what theme operations a tenant can perform
type TenantPermissions struct {
	CanSwitchThemes      bool     `json:"canSwitchThemes"`
	CanCustomizeTokens   bool     `json:"canCustomizeTokens"`
	CanUploadThemes      bool     `json:"canUploadThemes"`
	CanCreateThemes      bool     `json:"canCreateThemes"`
	AllowedThemes        []string `json:"allowedThemes,omitempty"`
	MaxCustomThemes      int      `json:"maxCustomThemes"`
	MaxTokenOverrides    int      `json:"maxTokenOverrides"`
}

// TenantThemeResult represents the result of tenant theme operations
type TenantThemeResult struct {
	Success        bool              `json:"success"`
	TenantID       string            `json:"tenantId"`
	ThemeID        string            `json:"themeId"`
	CSS            string            `json:"css"`
	Variables      map[string]string `json:"variables"`
	Classes        []string          `json:"classes"`
	BrandingCSS    string            `json:"brandingCss,omitempty"`
	CustomCSS      string            `json:"customCss,omitempty"`
	Error          string            `json:"error,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
}

// NewTenantThemeManager creates a new multi-tenant theme manager
func NewTenantThemeManager(globalManager *Runtime, storage TenantStorage, config *MultiTenantConfig) *TenantThemeManager {
	if config == nil {
		config = DefaultMultiTenantConfig()
	}

	return &TenantThemeManager{
		globalManager:  globalManager,
		tenantRuntimes: make(map[string]*Runtime),
		tenantConfigs:  make(map[string]*TenantConfig),
		tenantStorage:  storage,
		config:         config,
	}
}

// DefaultMultiTenantConfig returns default multi-tenant configuration
func DefaultMultiTenantConfig() *MultiTenantConfig {
	return &MultiTenantConfig{
		EnableTenantIsolation:     true,
		EnableTenantInheritance:   true,
		DefaultTenantTheme:        "default",
		GlobalFallbackTheme:       "default",
		TenantCacheSize:           100,
		TenantCacheTTL:            30 * time.Minute,
		EnableTenantCustomization: true,
		MaxTenantsPerNode:         1000,
	}
}

// GetTenantRuntime returns a theme runtime for a specific tenant
func (ttm *TenantThemeManager) GetTenantRuntime(tenantID string) (*Runtime, error) {
	ttm.mu.RLock()
	runtime, exists := ttm.tenantRuntimes[tenantID]
	ttm.mu.RUnlock()

	if exists {
		return runtime, nil
	}

	// Create new runtime for tenant
	return ttm.createTenantRuntime(tenantID)
}

// SwitchTenantTheme switches theme for a specific tenant
func (ttm *TenantThemeManager) SwitchTenantTheme(ctx *TenantThemeContext) (*TenantThemeResult, error) {
	// Validate tenant permissions
	if err := ttm.validateTenantPermissions(ctx); err != nil {
		return &TenantThemeResult{
			Success:  false,
			TenantID: ctx.TenantID,
			Error:    err.Error(),
		}, err
	}

	// Get tenant runtime
	runtime, err := ttm.GetTenantRuntime(ctx.TenantID)
	if err != nil {
		return &TenantThemeResult{
			Success:  false,
			TenantID: ctx.TenantID,
			Error:    err.Error(),
		}, err
	}

	// Switch theme
	if err := runtime.SetTheme(ctx.ThemeID); err != nil {
		return &TenantThemeResult{
			Success:  false,
			TenantID: ctx.TenantID,
			ThemeID:  ctx.ThemeID,
			Error:    err.Error(),
		}, err
	}

	// Set dark mode if specified
	if err := runtime.SetDarkMode(ctx.DarkMode); err != nil {
		// Don't fail on dark mode errors
	}

	// Get compiled CSS
	css, err := runtime.GetThemeCSS()
	if err != nil {
		return &TenantThemeResult{
			Success:  false,
			TenantID: ctx.TenantID,
			ThemeID:  ctx.ThemeID,
			Error:    err.Error(),
		}, err
	}

	// Apply tenant customizations
	brandingCSS := ttm.generateBrandingCSS(ctx.TenantID)
	customCSS := ttm.generateCustomCSS(ctx.TenantID, ctx.CustomTokens)

	// Update tenant configuration
	if err := ttm.updateTenantTheme(ctx.TenantID, ctx.ThemeID, ctx.DarkMode); err != nil {
		// Log warning but don't fail
	}

	return &TenantThemeResult{
		Success:     true,
		TenantID:    ctx.TenantID,
		ThemeID:     ctx.ThemeID,
		CSS:         css,
		BrandingCSS: brandingCSS,
		CustomCSS:   customCSS,
		Variables:   ttm.extractVariables(css),
		Classes:     ttm.extractClasses(css),
		Metadata: map[string]any{
			"timestamp":   time.Now(),
			"darkMode":    ctx.DarkMode,
			"customized":  len(ctx.CustomTokens) > 0,
		},
	}, nil
}

// ConfigureTenant configures theme settings for a tenant
func (ttm *TenantThemeManager) ConfigureTenant(tenantID string, config *TenantConfig) error {
	ttm.mu.Lock()
	defer ttm.mu.Unlock()

	// Validate configuration
	if err := ttm.validateTenantConfig(config); err != nil {
		return fmt.Errorf("invalid tenant config: %w", err)
	}

	// Update timestamps
	now := time.Now()
	config.UpdatedAt = now
	if config.CreatedAt.IsZero() {
		config.CreatedAt = now
	}

	// Store configuration
	ttm.tenantConfigs[tenantID] = config

	// Persist to storage
	if ttm.tenantStorage != nil {
		if err := ttm.tenantStorage.SaveTenantConfig(tenantID, config); err != nil {
			return fmt.Errorf("failed to save tenant config: %w", err)
		}
	}

	// Clear existing runtime to force recreation with new config
	delete(ttm.tenantRuntimes, tenantID)

	return nil
}

// GetTenantConfig returns configuration for a tenant
func (ttm *TenantThemeManager) GetTenantConfig(tenantID string) (*TenantConfig, error) {
	ttm.mu.RLock()
	config, exists := ttm.tenantConfigs[tenantID]
	ttm.mu.RUnlock()

	if exists {
		return config, nil
	}

	// Load from storage
	if ttm.tenantStorage != nil {
		config, err := ttm.tenantStorage.LoadTenantConfig(tenantID)
		if err != nil {
			// Return default config if not found
			return ttm.createDefaultTenantConfig(tenantID), nil
		}

		// Cache the config
		ttm.mu.Lock()
		ttm.tenantConfigs[tenantID] = config
		ttm.mu.Unlock()

		return config, nil
	}

	return ttm.createDefaultTenantConfig(tenantID), nil
}

// RemoveTenant removes a tenant and cleans up resources
func (ttm *TenantThemeManager) RemoveTenant(tenantID string) error {
	ttm.mu.Lock()
	defer ttm.mu.Unlock()

	// Remove from memory
	delete(ttm.tenantRuntimes, tenantID)
	delete(ttm.tenantConfigs, tenantID)

	// Remove from storage
	if ttm.tenantStorage != nil {
		if err := ttm.tenantStorage.DeleteTenantConfig(tenantID); err != nil {
			return fmt.Errorf("failed to delete tenant config: %w", err)
		}
	}

	return nil
}

// ListTenants returns all configured tenants
func (ttm *TenantThemeManager) ListTenants() ([]TenantConfig, error) {
	if ttm.tenantStorage != nil {
		return ttm.tenantStorage.ListTenantConfigs()
	}

	ttm.mu.RLock()
	defer ttm.mu.RUnlock()

	configs := make([]TenantConfig, 0, len(ttm.tenantConfigs))
	for _, config := range ttm.tenantConfigs {
		configs = append(configs, *config)
	}

	return configs, nil
}

// GetTenantThemes returns available themes for a tenant
func (ttm *TenantThemeManager) GetTenantThemes(tenantID string) ([]*ThemeInfo, error) {
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return nil, err
	}

	allThemes := ttm.globalManager.GetAvailableThemes()
	var tenantThemes []*ThemeInfo

	// Filter based on tenant configuration
	for _, theme := range allThemes {
		if ttm.isThemeAllowedForTenant(theme.ID, config) {
			tenantThemes = append(tenantThemes, theme)
		}
	}

	return tenantThemes, nil
}

// ValidateTenantAccess validates if a tenant can access a specific theme
func (ttm *TenantThemeManager) ValidateTenantAccess(tenantID, themeID string) error {
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return err
	}

	if !ttm.isThemeAllowedForTenant(themeID, config) {
		return fmt.Errorf("tenant %s is not allowed to access theme %s", tenantID, themeID)
	}

	return nil
}

// CreateTenantCustomTheme creates a custom theme for a tenant
func (ttm *TenantThemeManager) CreateTenantCustomTheme(tenantID, baseThemeID string, overrides *BrandingOverrides) (*schema.Theme, error) {
	// Validate tenant permissions
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return nil, err
	}

	if !ttm.config.EnableTenantCustomization {
		return nil, fmt.Errorf("tenant customization is disabled")
	}

	// Check custom theme limits
	if len(config.CustomThemes) >= 5 { // Default limit
		return nil, fmt.Errorf("tenant has reached custom theme limit")
	}

	// Get base theme
	baseTheme, err := ttm.globalManager.GetTheme(baseThemeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get base theme: %w", err)
	}

	// Create custom theme
	customTheme := ttm.applyBrandingOverrides(baseTheme, overrides)
	customTheme.ID = fmt.Sprintf("%s-custom-%s", tenantID, baseThemeID)
	customTheme.Name = fmt.Sprintf("%s (Custom)", baseTheme.Name)

	// Register custom theme
	runtime, err := ttm.GetTenantRuntime(tenantID)
	if err != nil {
		return nil, err
	}

	if err := (*runtime.api).RegisterTheme(customTheme); err != nil {
		return nil, fmt.Errorf("failed to register custom theme: %w", err)
	}

	// Update tenant config
	config.CustomThemes = append(config.CustomThemes, customTheme.ID)
	if err := ttm.ConfigureTenant(tenantID, config); err != nil {
		return nil, fmt.Errorf("failed to update tenant config: %w", err)
	}

	return customTheme, nil
}

// Private helper methods

func (ttm *TenantThemeManager) createTenantRuntime(tenantID string) (*Runtime, error) {
	ttm.mu.Lock()
	defer ttm.mu.Unlock()

	// Check if runtime already exists (double-check pattern)
	if runtime, exists := ttm.tenantRuntimes[tenantID]; exists {
		return runtime, nil
	}

	// Check tenant limit
	if len(ttm.tenantRuntimes) >= ttm.config.MaxTenantsPerNode {
		return nil, fmt.Errorf("maximum tenants per node reached")
	}

	// Get tenant configuration
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return nil, err
	}

	// Create runtime config based on tenant settings
	runtimeConfig := &RuntimeConfig{
		EnableCaching:   true,
		EnableDarkMode:  config.EnableDarkMode,
		FallbackTheme:   config.DefaultTheme,
		ValidateTokens:  true,
	}

	// Create new runtime for tenant
	runtime := NewRuntime(*ttm.globalManager.api, runtimeConfig)
	runtime.SetTenant(tenantID)

	// Set default theme
	if config.ThemeID != "" {
		runtime.SetTheme(config.ThemeID)
	} else if config.DefaultTheme != "" {
		runtime.SetTheme(config.DefaultTheme)
	} else {
		runtime.SetTheme(ttm.config.DefaultTenantTheme)
	}

	ttm.tenantRuntimes[tenantID] = runtime
	return runtime, nil
}

func (ttm *TenantThemeManager) validateTenantPermissions(ctx *TenantThemeContext) error {
	config, err := ttm.GetTenantConfig(ctx.TenantID)
	if err != nil {
		return err
	}

	if !config.Active {
		return fmt.Errorf("tenant %s is not active", ctx.TenantID)
	}

	// Check theme access
	return ttm.ValidateTenantAccess(ctx.TenantID, ctx.ThemeID)
}

func (ttm *TenantThemeManager) validateTenantConfig(config *TenantConfig) error {
	if config.TenantID == "" {
		return fmt.Errorf("tenant ID is required")
	}

	// Validate themes exist
	if config.ThemeID != "" {
		if _, err := ttm.globalManager.GetTheme(config.ThemeID); err != nil {
			return fmt.Errorf("invalid theme ID: %s", config.ThemeID)
		}
	}

	return nil
}

func (ttm *TenantThemeManager) isThemeAllowedForTenant(themeID string, config *TenantConfig) bool {
	// Check blocked themes first
	for _, blocked := range config.BlockedThemes {
		if blocked == themeID {
			return false
		}
	}

	// If there are allowed themes, check if theme is in the list
	if len(config.AllowedThemes) > 0 {
		for _, allowed := range config.AllowedThemes {
			if allowed == themeID {
				return true
			}
		}
		return false
	}

	// If no specific restrictions, allow all
	return true
}

func (ttm *TenantThemeManager) createDefaultTenantConfig(tenantID string) *TenantConfig {
	return &TenantConfig{
		TenantID:       tenantID,
		DefaultTheme:   ttm.config.DefaultTenantTheme,
		EnableDarkMode: true,
		Active:         true,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (ttm *TenantThemeManager) updateTenantTheme(tenantID, themeID string, darkMode bool) error {
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return err
	}

	config.ThemeID = themeID
	config.UpdatedAt = time.Now()

	return ttm.ConfigureTenant(tenantID, config)
}

func (ttm *TenantThemeManager) generateBrandingCSS(tenantID string) string {
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil || config.BrandingOverrides == nil {
		return ""
	}

	var css strings.Builder
	css.WriteString("/* Tenant Branding Overrides */\n")

	overrides := config.BrandingOverrides

	// Color overrides
	if overrides.PrimaryColor != "" {
		css.WriteString(fmt.Sprintf(":root { --color-primary: %s; }\n", overrides.PrimaryColor))
	}
	if overrides.SecondaryColor != "" {
		css.WriteString(fmt.Sprintf(":root { --color-secondary: %s; }\n", overrides.SecondaryColor))
	}
	if overrides.AccentColor != "" {
		css.WriteString(fmt.Sprintf(":root { --color-accent: %s; }\n", overrides.AccentColor))
	}

	// Typography overrides
	if overrides.FontFamily != "" {
		css.WriteString(fmt.Sprintf(":root { --font-family-base: %s; }\n", overrides.FontFamily))
	}

	// Custom color palette
	for name, color := range overrides.ColorPalette {
		css.WriteString(fmt.Sprintf(":root { --color-%s: %s; }\n", name, color))
	}

	return css.String()
}

func (ttm *TenantThemeManager) generateCustomCSS(tenantID string, customTokens map[string]string) string {
	config, err := ttm.GetTenantConfig(tenantID)
	if err != nil {
		return ""
	}

	var css strings.Builder

	// Add custom CSS from config
	if config.BrandingOverrides != nil && config.BrandingOverrides.CustomCSS != "" {
		css.WriteString("/* Tenant Custom CSS */\n")
		css.WriteString(config.BrandingOverrides.CustomCSS)
		css.WriteString("\n")
	}

	// Add custom token overrides
	if len(customTokens) > 0 {
		css.WriteString("/* Custom Token Overrides */\n")
		css.WriteString(":root {\n")
		for token, value := range customTokens {
			css.WriteString(fmt.Sprintf("  --%s: %s;\n", token, value))
		}
		css.WriteString("}\n")
	}

	return css.String()
}

func (ttm *TenantThemeManager) applyBrandingOverrides(baseTheme *schema.Theme, overrides *BrandingOverrides) *schema.Theme {
	// Deep clone the base theme
	customTheme := *baseTheme

	// Apply branding overrides to create a new custom theme
	// This would involve modifying the theme tokens based on the overrides

	return &customTheme
}

func (ttm *TenantThemeManager) extractVariables(css string) map[string]string {
	// Extract CSS variables from CSS string
	variables := make(map[string]string)
	// This would be implemented properly with CSS parsing
	return variables
}

func (ttm *TenantThemeManager) extractClasses(css string) []string {
	// Extract CSS classes from CSS string
	var classes []string
	// This would be implemented properly with CSS parsing
	return classes
}

// SimpleTenantStorage provides a simple in-memory storage implementation
type SimpleTenantStorage struct {
	configs map[string]*TenantConfig
	mu      sync.RWMutex
}

// NewSimpleTenantStorage creates a new simple tenant storage
func NewSimpleTenantStorage() *SimpleTenantStorage {
	return &SimpleTenantStorage{
		configs: make(map[string]*TenantConfig),
	}
}

// SaveTenantConfig saves tenant configuration
func (s *SimpleTenantStorage) SaveTenantConfig(tenantID string, config *TenantConfig) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.configs[tenantID] = config
	return nil
}

// LoadTenantConfig loads tenant configuration
func (s *SimpleTenantStorage) LoadTenantConfig(tenantID string) (*TenantConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if config, exists := s.configs[tenantID]; exists {
		return config, nil
	}

	return nil, fmt.Errorf("tenant config not found: %s", tenantID)
}

// DeleteTenantConfig deletes tenant configuration
func (s *SimpleTenantStorage) DeleteTenantConfig(tenantID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.configs, tenantID)
	return nil
}

// ListTenantConfigs lists all tenant configurations
func (s *SimpleTenantStorage) ListTenantConfigs() ([]TenantConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	configs := make([]TenantConfig, 0, len(s.configs))
	for _, config := range s.configs {
		configs = append(configs, *config)
	}

	return configs, nil
}

// UpdateTenantConfig updates tenant configuration
func (s *SimpleTenantStorage) UpdateTenantConfig(tenantID string, updates map[string]any) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	config, exists := s.configs[tenantID]
	if !exists {
		return fmt.Errorf("tenant config not found: %s", tenantID)
	}

	// Apply updates (simplified implementation)
	if themeID, ok := updates["themeId"].(string); ok {
		config.ThemeID = themeID
	}
	if darkMode, ok := updates["enableDarkMode"].(bool); ok {
		config.EnableDarkMode = darkMode
	}

	config.UpdatedAt = time.Now()
	return nil
}