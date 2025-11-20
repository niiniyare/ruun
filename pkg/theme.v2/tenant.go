package theme

import (
	"context"
	"sync"
	"time"
)

// TenantManager manages tenant-specific theme configurations and isolation.
type TenantManager struct {
	manager *Manager
	tenants map[string]*TenantConfig
	mu      sync.RWMutex
}

// TenantConfig contains tenant-specific theme configuration.
type TenantConfig struct {
	TenantID     string
	DefaultTheme string
	AllowedThemes []string
	Branding     *BrandingOverrides
	Features     map[string]bool
	CustomData   map[string]interface{}
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// BrandingOverrides contains tenant-specific branding customizations.
type BrandingOverrides struct {
	PrimaryColor   string
	SecondaryColor string
	Logo           string
	Favicon        string
	CompanyName    string
	CustomCSS      string
	CustomTokens   map[string]string
}

// NewTenantManager creates a new tenant manager.
func NewTenantManager(manager *Manager) *TenantManager {
	return &TenantManager{
		manager: manager,
		tenants: make(map[string]*TenantConfig),
	}
}

// ConfigureTenant configures a tenant with specific theme settings.
func (tm *TenantManager) ConfigureTenant(ctx context.Context, config *TenantConfig) error {
	if config == nil {
		return NewError(ErrCodeValidation, "tenant config cannot be nil")
	}
	if config.TenantID == "" {
		return NewError(ErrCodeValidation, "tenant ID cannot be empty")
	}

	if config.DefaultTheme != "" {
		_, err := tm.manager.loadTheme(ctx, config.DefaultTheme)
		if err != nil {
			return WrapError(ErrCodeNotFound,
				"default theme not found", err)
		}
	}

	now := time.Now()
	config.UpdatedAt = now
	
	tm.mu.Lock()
	if _, exists := tm.tenants[config.TenantID]; !exists {
		config.CreatedAt = now
	}
	tm.tenants[config.TenantID] = config
	tm.mu.Unlock()

	return nil
}

// GetTenantConfig retrieves tenant configuration.
func (tm *TenantManager) GetTenantConfig(ctx context.Context, tenantID string) (*TenantConfig, error) {
	tm.mu.RLock()
	config, exists := tm.tenants[tenantID]
	tm.mu.RUnlock()

	if !exists {
		return nil, NewErrorf(ErrCodeNotFound, "tenant not found: %s", tenantID)
	}

	return config, nil
}

// GetTenantTheme retrieves and compiles the theme for a tenant with branding overrides.
func (tm *TenantManager) GetTenantTheme(ctx context.Context, tenantID string, evalData map[string]any) (*CompiledTheme, error) {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	themeID := config.DefaultTheme
	if themeID == "" {
		return nil, NewError(ErrCodeValidation, "tenant has no default theme")
	}

	ctxWithTenant := WithTenant(ctx, tenantID)
	
	compiled, err := tm.manager.GetTheme(ctxWithTenant, themeID, evalData)
	if err != nil {
		return nil, err
	}

	if config.Branding != nil {
		compiled.Theme = tm.applyBrandingOverrides(compiled.Theme, config.Branding)
		
		if len(config.Branding.CustomTokens) > 0 {
			tm.applyCustomTokens(compiled.ResolvedTokens, config.Branding.CustomTokens)
		}

		if config.Branding.CustomCSS != "" {
			compiled.CSS += "\n" + config.Branding.CustomCSS
		}
	}

	return compiled, nil
}

// applyBrandingOverrides applies branding customizations to a theme.
func (tm *TenantManager) applyBrandingOverrides(theme *Theme, branding *BrandingOverrides) *Theme {
	cloned := theme.Clone()

	if branding.PrimaryColor != "" && cloned.Tokens.Primitives != nil && cloned.Tokens.Primitives.Colors != nil {
		cloned.Tokens.Primitives.Colors["primary"] = branding.PrimaryColor
	}

	if branding.SecondaryColor != "" && cloned.Tokens.Primitives != nil && cloned.Tokens.Primitives.Colors != nil {
		cloned.Tokens.Primitives.Colors["secondary"] = branding.SecondaryColor
	}

	return cloned
}

// applyCustomTokens applies tenant-specific token overrides.
func (tm *TenantManager) applyCustomTokens(tokens *Tokens, customTokens map[string]string) {
	if tokens == nil || len(customTokens) == 0 {
		return
	}

	for path, value := range customTokens {
		tm.manager.setTokenValue(tokens, path, value)
	}
}

// ListTenantThemes returns all themes accessible to a tenant.
func (tm *TenantManager) ListTenantThemes(ctx context.Context, tenantID string) ([]*Theme, error) {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	allThemes, err := tm.manager.ListThemes(ctx)
	if err != nil {
		return nil, err
	}

	if len(config.AllowedThemes) == 0 {
		return allThemes, nil
	}

	allowed := make(map[string]bool)
	for _, themeID := range config.AllowedThemes {
		allowed[themeID] = true
	}

	filtered := make([]*Theme, 0)
	for _, theme := range allThemes {
		if allowed[theme.ID] {
			filtered = append(filtered, theme)
		}
	}

	return filtered, nil
}

// SetTenantTheme sets the default theme for a tenant.
func (tm *TenantManager) SetTenantTheme(ctx context.Context, tenantID, themeID string) error {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return err
	}

	_, err = tm.manager.loadTheme(ctx, themeID)
	if err != nil {
		return WrapError(ErrCodeNotFound, "theme not found", err)
	}

	if len(config.AllowedThemes) > 0 {
		allowed := false
		for _, allowed_id := range config.AllowedThemes {
			if allowed_id == themeID {
				allowed = true
				break
			}
		}
		if !allowed {
			return NewErrorf(ErrCodeValidation, "theme %s not allowed for tenant %s", themeID, tenantID)
		}
	}

	tm.mu.Lock()
	config.DefaultTheme = themeID
	config.UpdatedAt = time.Now()
	tm.mu.Unlock()

	tm.manager.InvalidateCache(themeID, tenantID)

	return nil
}

// SetBranding sets branding overrides for a tenant.
func (tm *TenantManager) SetBranding(ctx context.Context, tenantID string, branding *BrandingOverrides) error {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return err
	}

	tm.mu.Lock()
	config.Branding = branding
	config.UpdatedAt = time.Now()
	tm.mu.Unlock()

	tm.manager.InvalidateCache("", tenantID)

	return nil
}

// EnableFeature enables a feature flag for a tenant.
func (tm *TenantManager) EnableFeature(ctx context.Context, tenantID, feature string) error {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return err
	}

	tm.mu.Lock()
	if config.Features == nil {
		config.Features = make(map[string]bool)
	}
	config.Features[feature] = true
	config.UpdatedAt = time.Now()
	tm.mu.Unlock()

	return nil
}

// DisableFeature disables a feature flag for a tenant.
func (tm *TenantManager) DisableFeature(ctx context.Context, tenantID, feature string) error {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return err
	}

	tm.mu.Lock()
	if config.Features == nil {
		config.Features = make(map[string]bool)
	}
	config.Features[feature] = false
	config.UpdatedAt = time.Now()
	tm.mu.Unlock()

	return nil
}

// IsFeatureEnabled checks if a feature is enabled for a tenant.
func (tm *TenantManager) IsFeatureEnabled(ctx context.Context, tenantID, feature string) bool {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return false
	}

	tm.mu.RLock()
	defer tm.mu.RUnlock()

	if config.Features == nil {
		return false
	}

	enabled, exists := config.Features[feature]
	return exists && enabled
}

// DeleteTenant removes tenant configuration.
func (tm *TenantManager) DeleteTenant(ctx context.Context, tenantID string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.tenants[tenantID]; !exists {
		return NewErrorf(ErrCodeNotFound, "tenant not found: %s", tenantID)
	}

	delete(tm.tenants, tenantID)
	tm.manager.InvalidateCache("", tenantID)

	return nil
}

// ListTenants returns all configured tenants.
func (tm *TenantManager) ListTenants(ctx context.Context) ([]*TenantConfig, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	tenants := make([]*TenantConfig, 0, len(tm.tenants))
	for _, config := range tm.tenants {
		tenants = append(tenants, config)
	}

	return tenants, nil
}

// GetTenantStats returns usage statistics for a tenant.
func (tm *TenantManager) GetTenantStats(ctx context.Context, tenantID string) (*TenantStats, error) {
	config, err := tm.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return nil, err
	}

	return &TenantStats{
		TenantID:      config.TenantID,
		DefaultTheme:  config.DefaultTheme,
		ThemeCount:    len(config.AllowedThemes),
		FeatureCount:  len(config.Features),
		HasBranding:   config.Branding != nil,
		LastUpdated:   config.UpdatedAt,
	}, nil
}

// TenantStats contains tenant usage statistics.
type TenantStats struct {
	TenantID     string
	DefaultTheme string
	ThemeCount   int
	FeatureCount int
	HasBranding  bool
	LastUpdated  time.Time
}
