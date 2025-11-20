package theme

import (
	"fmt"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// Package version and metadata
const (
	// Version of the theme package
	Version = "1.0.0"
	
	// PackageName is the name of this package
	PackageName = "theme"
	
	// Description provides a brief description of the package
	Description = "Comprehensive theme integration utilities for the Ruun design system"
	
	// Author of the package
	Author = "Ruun Design System Team"
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
			"Runtime theme management",
			"Enhanced CSS compilation",
			"Theme validation and testing",
			"Multi-tenant support",
			"Component integration",
			"Alpine.js integration",
			"Performance optimization",
			"Dark mode support",
			"Responsive utilities",
			"Token resolution",
		},
		Components: []string{
			"Runtime",
			"EnhancedCompiler",
			"TokenResolver",
			"ThemeValidator",
			"ThemeTester",
			"ThemeSwitcher",
			"ComponentIntegrator",
			"TenantThemeManager",
		},
	}
}

// DefaultConfigurations provides default configurations for all components
type DefaultConfigurations struct {
	Runtime     *RuntimeConfig
	Compiler    *CompilerConfig
	Integration *IntegrationConfig
	Validation  *ValidatorConfig
	Testing     *TestConfig
	Switcher    *SwitcherConfig
	MultiTenant *MultiTenantConfig
}

// GetDefaultConfigurations returns default configurations for all components
func GetDefaultConfigurations() *DefaultConfigurations {
	return &DefaultConfigurations{
		Runtime:     DefaultRuntimeConfig(),
		Compiler:    DefaultCompilerConfig(),
		Integration: DefaultIntegrationConfig(),
		Validation:  DefaultValidatorConfig(),
		Testing:     DefaultTestConfig(),
		Switcher:    DefaultSwitcherConfig(),
		MultiTenant: DefaultMultiTenantConfig(),
	}
}

// Quick setup functions for common use cases

// NewBasicThemeSetup creates a basic theme setup for simple applications
func NewBasicThemeSetup(api ThemeAPIInterface) (*Runtime, error) {
	runtime := NewRuntime(api, DefaultRuntimeConfig())
	return runtime, nil
}

// NewAdvancedThemeSetup creates an advanced theme setup with all features
func NewAdvancedThemeSetup(api ThemeAPIInterface) (*ThemeSystemBundle, error) {
	// Create runtime
	runtime := NewRuntime(api, DefaultRuntimeConfig())
	
	// Create resolver
	cache := &RuntimeCache{
		tokenValues: make(map[string]string),
		classNames:  make(map[string]string),
		compiled:    make(map[string]*CompiledTheme),
	}
	resolver := NewTokenResolver(api, cache)
	
	// Create compiler
	compiler := NewEnhancedCompiler(resolver, DefaultCompilerConfig())
	
	// Create validator and tester
	validator := NewThemeValidator()
	tester := NewThemeTester(validator, runtime, DefaultTestConfig())
	
	// Create switcher
	storage := NewSimpleThemeStorage()
	switcher := NewThemeSwitcher(runtime, storage, DefaultSwitcherConfig())
	
	// Create integrator
	integrator := NewComponentIntegrator(runtime, resolver, DefaultIntegrationConfig())
	
	return &ThemeSystemBundle{
		Runtime:    runtime,
		Compiler:   compiler,
		Resolver:   resolver,
		Validator:  validator,
		Tester:     tester,
		Switcher:   switcher,
		Integrator: integrator,
	}, nil
}

// NewMultiTenantThemeSetup creates a multi-tenant theme setup
func NewMultiTenantThemeSetup(api ThemeAPIInterface) (*MultiTenantThemeBundle, error) {
	// Create basic setup
	basic, err := NewAdvancedThemeSetup(api)
	if err != nil {
		return nil, err
	}
	
	// Create tenant components
	tenantStorage := NewSimpleTenantStorage()
	tenantManager := NewTenantThemeManager(basic.Runtime, tenantStorage, DefaultMultiTenantConfig())
	
	return &MultiTenantThemeBundle{
		ThemeSystemBundle: *basic,
		TenantManager:     tenantManager,
		TenantStorage:     tenantStorage,
	}, nil
}

// Bundle types for organized component groups

// ThemeSystemBundle groups core theme system components
type ThemeSystemBundle struct {
	Runtime    *Runtime
	Compiler   *EnhancedCompiler
	Resolver   *TokenResolver
	Validator  *ThemeValidator
	Tester     *ThemeTester
	Switcher   *ThemeSwitcher
	Integrator *ComponentIntegrator
}

// MultiTenantThemeBundle extends ThemeSystemBundle with multi-tenant support
type MultiTenantThemeBundle struct {
	ThemeSystemBundle
	TenantManager *TenantThemeManager
	TenantStorage TenantStorage
}

// GetSystemStats returns comprehensive statistics about the theme system
func (tsb *ThemeSystemBundle) GetSystemStats() map[string]any {
	stats := make(map[string]any)
	
	// Runtime stats
	if tsb.Runtime != nil {
		runtimeStats := tsb.Runtime.GetCacheStats()
		stats["runtime"] = runtimeStats
	}
	
	// Compiler stats
	if tsb.Compiler != nil {
		compilerStats := tsb.Compiler.GetCacheStats()
		stats["compiler"] = compilerStats
	}
	
	// Integration stats
	if tsb.Integrator != nil {
		integrationStats := tsb.Integrator.GetIntegrationStats()
		stats["integration"] = integrationStats
	}
	
	return stats
}

// GetTenantStats returns statistics about multi-tenant usage
func (mtb *MultiTenantThemeBundle) GetTenantStats() map[string]any {
	stats := mtb.GetSystemStats()
	
	if mtb.TenantManager != nil {
		tenants, _ := mtb.TenantManager.ListTenants()
		stats["tenants"] = map[string]any{
			"total":  len(tenants),
			"active": countActiveTenants(tenants),
		}
	}
	
	return stats
}

// Helper function
func countActiveTenants(tenants []TenantConfig) int {
	count := 0
	for _, tenant := range tenants {
		if tenant.Active {
			count++
		}
	}
	return count
}

// Utility functions for common operations

// SwitchThemeForUser is a convenience function to switch themes for a user
func SwitchThemeForUser(switcher *ThemeSwitcher, themeID, userID string) (*ThemeSwitchResult, error) {
	return switcher.SwitchTheme(themeID, userID)
}

// SwitchThemeForTenant is a convenience function to switch themes for a tenant
func SwitchThemeForTenant(tenantManager *TenantThemeManager, tenantID, themeID, userID string) (*TenantThemeResult, error) {
	ctx := &TenantThemeContext{
		TenantID: tenantID,
		UserID:   userID,
		ThemeID:  themeID,
		DarkMode: false,
	}
	return tenantManager.SwitchTenantTheme(ctx)
}

// ValidateAndCompileTheme validates and compiles a theme in one operation
func ValidateAndCompileTheme(validator *ThemeValidator, compiler *EnhancedCompiler, theme *schema.Theme) (*ValidationResult, *CompilationResult, error) {
	// Validate first
	validationResult, err := validator.ValidateTheme(theme)
	if err != nil {
		return nil, nil, err
	}
	
	// Don't compile if validation failed
	if !validationResult.Valid {
		return validationResult, nil, fmt.Errorf("theme validation failed with %d issues", len(validationResult.Issues))
	}
	
	// Compile theme
	compilationResult, err := compiler.CompileTheme(theme)
	if err != nil {
		return validationResult, nil, err
	}
	
	return validationResult, compilationResult, nil
}

// GenerateThemeBundle generates a complete theme bundle with CSS, validation, and metadata
type ThemeBundle struct {
	Theme            *schema.Theme        `json:"theme"`
	CSS              string               `json:"css"`
	ValidationResult *ValidationResult    `json:"validation"`
	CompilationResult *CompilationResult  `json:"compilation"`
	Metadata         map[string]any `json:"metadata"`
}

// GenerateCompleteThemeBundle creates a comprehensive theme bundle
func GenerateCompleteThemeBundle(bundle *ThemeSystemBundle, theme *schema.Theme) (*ThemeBundle, error) {
	// Validate and compile
	validation, compilation, err := ValidateAndCompileTheme(bundle.Validator, bundle.Compiler, theme)
	if err != nil {
		return nil, err
	}
	
	return &ThemeBundle{
		Theme:             theme,
		CSS:               compilation.CSS,
		ValidationResult:  validation,
		CompilationResult: compilation,
		Metadata: map[string]any{
			"packageVersion": Version,
			"generated":      time.Now(),
			"bundleSize":     len(compilation.CSS),
			"variables":      len(compilation.Variables),
			"classes":        len(compilation.Classes),
		},
	}, nil
}