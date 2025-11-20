package theme

import (
	"context"
	"fmt"
	"time"
)

// Package version and metadata
const (
	// Version of the theme package
	Version = "2.0.0"

	// PackageName is the name of this package
	PackageName = "theme"

	// Description provides a brief description of the package
	Description = "Enterprise-grade theme management system for multi-tenant applications"

	// Author of the package
	Author = "Awo ERP Team"
)

// ThemePackageInfo provides information about the theme package
type ThemePackageInfo struct {
	Version     string   `json:"version"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Features    []string `json:"features"`
	Components  []string `json:"components"`
}

// GetPackageInfo returns information about the theme package
func GetPackageInfo() *ThemePackageInfo {
	return &ThemePackageInfo{
		Version:     Version,
		Name:        PackageName,
		Description: Description,
		Author:      Author,
		Features: []string{
			"Manager-based lifecycle",
			"Conditional theming",
			"Token resolution with caching",
			"CSS compilation",
			"Multi-tenant isolation",
			"Observer pattern",
			"Theme validation",
			"Alpine.js integration",
			"Dark mode support",
			"User preferences",
		},
		Components: []string{
			"Manager",
			"TenantManager",
			"Resolver",
			"Compiler",
			"Validator",
			"ThemeSwitcher",
			"Observer",
		},
	}
}

// ThemeSystemBundle groups core theme system components with new architecture
type ThemeSystemBundle struct {
	Manager       *Manager
	TenantManager *TenantManager
	Resolver      *Resolver
	Compiler      *Compiler
	Validator     *Validator
	Switcher      *ThemeSwitcher
	Observer      Observer
}

// MultiTenantBundle extends ThemeSystemBundle with multi-tenant features
type MultiTenantBundle struct {
	*ThemeSystemBundle
	TenantStorage TenantStorage
	UserStorage   UserPreferenceStorage
}

// NewBasicSetup creates a basic theme setup for simple applications
func NewBasicSetup(ctx context.Context, storage Storage) (*Manager, error) {
	config := DefaultManagerConfig()
	return NewManager(config, storage, nil)
}

// NewAdvancedSetup creates an advanced theme setup with all features
func NewAdvancedSetup(ctx context.Context, storage Storage, tenantStorage TenantStorage, evaluator ConditionEvaluator) (*ThemeSystemBundle, error) {
	// Create manager
	managerConfig := DefaultManagerConfig()
	manager, err := NewManager(managerConfig, storage, evaluator)
	if err != nil {
		return nil, fmt.Errorf("failed to create manager: %w", err)
	}

	// Create resolver
	resolverConfig := DefaultResolverConfig()
	resolver, err := NewResolver(resolverConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create resolver: %w", err)
	}

	// Create compiler
	compiler, err := NewCompiler(managerConfig.CompilerConfig)
	if err != nil {
		return nil, NewError("compiler_config_err", "compiler can not be initialzed")
	}
	// Create validator
	validatorConfig := DefaultValidatorConfig()
	validator := NewValidator(validatorConfig)

	// Create tenant manager
	tenantManager := NewTenantManager(manager, tenantStorage)

	// Create switcher (optional)
	userStorage := NewSimpleUserPreferenceStorage()
	switcherConfig := DefaultSwitcherConfig()
	switcher := NewThemeSwitcher(manager, tenantManager, userStorage, switcherConfig)

	return &ThemeSystemBundle{
		Manager:       manager,
		TenantManager: tenantManager,
		Resolver:      resolver,
		Compiler:      compiler,
		Validator:     validator,
		Switcher:      switcher,
	}, nil
}

// NewMultiTenantSetup creates a multi-tenant theme setup
func NewMultiTenantSetup(ctx context.Context, storage Storage, tenantStorage TenantStorage, evaluator ConditionEvaluator) (*MultiTenantBundle, error) {
	// Create basic setup
	basic, err := NewAdvancedSetup(ctx, storage, tenantStorage, evaluator)
	if err != nil {
		return nil, err
	}

	// Create user preference storage
	userStorage := NewSimpleUserPreferenceStorage()

	return &MultiTenantBundle{
		ThemeSystemBundle: basic,
		TenantStorage:     tenantStorage,
		UserStorage:       userStorage,
	}, nil
}

// GetSystemStats returns comprehensive statistics about the theme system
func (tsb *ThemeSystemBundle) GetSystemStats() map[string]any {
	stats := make(map[string]any)

	// Manager stats
	if tsb.Manager != nil {
		managerStats := tsb.Manager.GetStats()
		stats["manager"] = map[string]any{
			"themesRegistered": managerStats.ThemeCount,
		}
	}

	// Resolver stats
	if tsb.Resolver != nil {
		resolverStats := tsb.Resolver.GetStats()
		stats["resolver"] = map[string]any{
			"cacheHits":   resolverStats.CacheHits,
			"cacheMisses": resolverStats.CacheMisses,
			"hitRate":     resolverStats.CacheHitRate,
		}
	}

	// Compiler stats
	if tsb.Compiler != nil {
		compilerStats := tsb.Compiler.GetStats()
		stats["compiler"] = compilerStats
	}

	return stats
}

// GetTenantStats returns statistics about multi-tenant usage
func (mtb *MultiTenantBundle) GetTenantStats(ctx context.Context) (map[string]any, error) {
	stats := mtb.GetSystemStats()

	if mtb.TenantManager != nil {
		tenants, err := mtb.TenantManager.ListTenants(ctx)
		if err != nil {
			return stats, err
		}

		activeTenants := 0
		for _, tenant := range tenants {
			if tenant.Active {
				activeTenants++
			}
		}

		stats["tenants"] = map[string]any{
			"total":  len(tenants),
			"active": activeTenants,
		}
	}

	return stats, nil
}

// Utility functions for common operations

// SwitchThemeForUser is a convenience function to switch themes for a user
func SwitchThemeForUser(ctx context.Context, switcher *ThemeSwitcher, themeID, userID string) (*ThemeSwitchResult, error) {
	return switcher.SwitchTheme(ctx, themeID, userID, "")
}

// SwitchThemeForTenant is a convenience function to switch themes for a tenant
func SwitchThemeForTenant(ctx context.Context, switcher *ThemeSwitcher, tenantID, themeID, userID string) (*ThemeSwitchResult, error) {
	return switcher.SwitchTheme(ctx, themeID, userID, tenantID)
}

// ToggleDarkModeForUser is a convenience function to toggle dark mode for a user
func ToggleDarkModeForUser(ctx context.Context, switcher *ThemeSwitcher, userID string) (*ThemeSwitchResult, error) {
	return switcher.ToggleDarkMode(ctx, userID, "")
}

// ValidateAndRegisterTheme validates and registers a theme in one operation
func ValidateAndRegisterTheme(ctx context.Context, validator *Validator, manager *Manager, theme *Theme) (*ValidationResult, error) {
	// Validate first
	validationResult := validator.Validate(theme)

	// Don't register if validation failed
	if !validationResult.Valid {
		return validationResult, fmt.Errorf("theme validation failed with %d issues", len(validationResult.Issues))
	}

	// Register theme
	if err := manager.RegisterTheme(ctx, theme); err != nil {
		return validationResult, fmt.Errorf("failed to register theme: %w", err)
	}

	return validationResult, nil
}

// CompileAndValidateTheme compiles and validates a theme
func CompileAndValidateTheme(ctx context.Context, compiler *Compiler, validator *Validator, tokens *Tokens, theme *Theme) (string, *ValidationResult, error) {
	// Compile first
	css, err := compiler.Compile(ctx, tokens, theme)
	if err != nil {
		return "", nil, fmt.Errorf("compilation failed: %w", err)
	}

	// Validate
	validationResult := validator.Validate(theme)

	return css, validationResult, nil
}

// ThemeBundle represents a complete theme bundle with metadata
type ThemeBundle struct {
	Theme      *Theme            `json:"theme"`
	CSS        string            `json:"css"`
	Validation *ValidationResult `json:"validation"`
	Metadata   map[string]any    `json:"metadata"`
}

// GenerateThemeBundle creates a comprehensive theme bundle
func GenerateThemeBundle(ctx context.Context, bundle *ThemeSystemBundle, themeID string) (*ThemeBundle, error) {
	// Get compiled theme
	compiled, err := bundle.Manager.GetTheme(ctx, themeID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get theme: %w", err)
	}

	// Validate theme
	validation := bundle.Validator.Validate(compiled.Theme)

	return &ThemeBundle{
		Theme:      compiled.Theme,
		CSS:        compiled.CSS,
		Validation: validation,
		Metadata: map[string]any{
			"packageVersion": Version,
			"generated":      time.Now(),
			"bundleSize":     len(compiled.CSS),
			"valid":          validation.Valid,
			"score":          validation.Score,
		},
	}, nil
}

// ExportThemeBundle exports a theme bundle for a tenant
func ExportTenantThemeBundle(ctx context.Context, bundle *MultiTenantBundle, tenantID, themeID string) (*ThemeBundle, error) {
	// Get tenant theme
	compiled, err := bundle.TenantManager.GetTenantTheme(ctx, tenantID, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant theme: %w", err)
	}

	// Validate theme
	validation := bundle.Validator.Validate(compiled.Theme)

	// Get tenant config
	tenantConfig, err := bundle.TenantManager.GetTenantConfig(ctx, tenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant config: %w", err)
	}

	return &ThemeBundle{
		Theme:      compiled.Theme,
		CSS:        compiled.CSS,
		Validation: validation,
		Metadata: map[string]any{
			"packageVersion": Version,
			"generated":      time.Now(),
			"tenantId":       tenantID,
			"tenantName":     tenantConfig.TenantID,
			"bundleSize":     len(compiled.CSS),
			"valid":          validation.Valid,
			"score":          validation.Score,
		},
	}, nil
}

// RegisterDefaultTheme registers the default theme
func RegisterDefaultTheme(ctx context.Context, manager *Manager) error {
	defaultTheme := GetDefaultTheme()
	return manager.RegisterTheme(ctx, defaultTheme)
}

// PreloadThemes preloads multiple themes for performance
func PreloadThemes(ctx context.Context, manager *Manager, themeIDs []string) error {
	for _, themeID := range themeIDs {
		_, err := manager.GetTheme(ctx, themeID, nil)
		if err != nil {
			return fmt.Errorf("failed to preload theme %s: %w", themeID, err)
		}
	}
	return nil
}

// CloseBundle closes all components in the bundle
func (tsb *ThemeSystemBundle) Close() error {
	var errs []error

	if tsb.Manager != nil {
		if err := tsb.Manager.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if tsb.Resolver != nil {
		tsb.Resolver.Close()
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors closing bundle: %v", errs)
	}

	return nil
}

// CloseMultiTenantBundle closes all components in the multi-tenant bundle
func (mtb *MultiTenantBundle) Close() error {
	return mtb.ThemeSystemBundle.Close()
}

// QuickStart provides a quick start function for common use cases
func QuickStart(ctx context.Context) (*ThemeSystemBundle, error) {
	// Use memory storage for quick start
	storage := NewMemoryStorage()

	// Create advanced setup
	bundle, err := NewAdvancedSetup(ctx, storage, nil, nil)
	if err != nil {
		return nil, err
	}

	// Register default theme
	if err := RegisterDefaultTheme(ctx, bundle.Manager); err != nil {
		return nil, fmt.Errorf("failed to register default theme: %w", err)
	}

	return bundle, nil
}

// QuickStartMultiTenant provides a quick start for multi-tenant applications
func QuickStartMultiTenant(ctx context.Context) (*MultiTenantBundle, error) {
	// Use memory storage for quick start
	storage := NewMemoryStorage()
	tenantStorage := NewSimpleTenantStorage()

	// Create multi-tenant setup
	bundle, err := NewMultiTenantSetup(ctx, storage, tenantStorage, nil)
	if err != nil {
		return nil, err
	}

	// Register default theme
	if err := RegisterDefaultTheme(ctx, bundle.Manager); err != nil {
		return nil, fmt.Errorf("failed to register default theme: %w", err)
	}

	return bundle, nil
}

// Helper functions for creating HTTP handlers

// ThemeHandlerHelper provides helper functions for HTTP handlers
type ThemeHandlerHelper struct {
	bundle *ThemeSystemBundle
}

// NewThemeHandlerHelper creates a new theme handler helper
func NewThemeHandlerHelper(bundle *ThemeSystemBundle) *ThemeHandlerHelper {
	return &ThemeHandlerHelper{bundle: bundle}
}

// GetThemeCSS returns CSS for a theme
func (h *ThemeHandlerHelper) GetThemeCSS(ctx context.Context, themeID string, evalData map[string]any) (string, error) {
	compiled, err := h.bundle.Manager.GetTheme(ctx, themeID, evalData)
	if err != nil {
		return "", err
	}
	return compiled.CSS, nil
}

// GetTenantThemeCSS returns CSS for a tenant theme
func (h *ThemeHandlerHelper) GetTenantThemeCSS(ctx context.Context, tenantID string, evalData map[string]any) (string, error) {
	compiled, err := h.bundle.TenantManager.GetTenantTheme(ctx, tenantID, evalData)
	if err != nil {
		return "", err
	}
	return compiled.CSS, nil
}

// SwitchUserTheme switches theme for a user
func (h *ThemeHandlerHelper) SwitchUserTheme(ctx context.Context, userID, themeID string) (*ThemeSwitchResult, error) {
	if h.bundle.Switcher == nil {
		return nil, fmt.Errorf("theme switcher not configured")
	}
	return h.bundle.Switcher.SwitchTheme(ctx, themeID, userID, "")
}

// ToggleUserDarkMode toggles dark mode for a user
func (h *ThemeHandlerHelper) ToggleUserDarkMode(ctx context.Context, userID string) (*ThemeSwitchResult, error) {
	if h.bundle.Switcher == nil {
		return nil, fmt.Errorf("theme switcher not configured")
	}
	return h.bundle.Switcher.ToggleDarkMode(ctx, userID, "")
}

// GetAlpineJSComponent returns Alpine.js component for a tenant
func (h *ThemeHandlerHelper) GetAlpineJSComponent(ctx context.Context, tenantID string) (string, error) {
	if h.bundle.Switcher == nil {
		return "", fmt.Errorf("theme switcher not configured")
	}
	return h.bundle.Switcher.GenerateAlpineJSComponent(ctx, tenantID)
}
