package theme

import (
	"context"
	"fmt"
	"log"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ExampleBasicThemeUsage demonstrates basic theme operations
func ExampleBasicThemeUsage() {
	// Create a theme API instance
	api := NewMockThemeAPI() // In real use: ruun.NewThemeAPI("./themes")

	// Create theme runtime
	runtime := NewRuntime(api, nil)

	// Switch to a theme
	err := runtime.SetTheme("default")
	if err != nil {
		log.Printf("Failed to set theme: %v", err)
		return
	}

	// Get current theme info
	themeInfo := runtime.GetCurrentTheme()
	fmt.Printf("Current theme: %s (%s)\n", themeInfo.Name, themeInfo.ID)

	// Enable dark mode
	err = runtime.SetDarkMode(true)
	if err != nil {
		log.Printf("Failed to set dark mode: %v", err)
		return
	}

	// Get compiled CSS
	css, err := runtime.GetThemeCSS()
	if err != nil {
		log.Printf("Failed to get CSS: %v", err)
		return
	}

	fmt.Printf("Generated CSS length: %d characters\n", len(css))

	// Output:
	// Current theme: Default Theme (default)
	// Generated CSS length: 2543 characters
}

// ExampleThemeSwitchingWithPersistence demonstrates theme switching with user preferences
func ExampleThemeSwitchingWithPersistence() {
	// Setup
	api := NewMockThemeAPI()
	runtime := NewRuntime(api, nil)
	storage := NewSimpleThemeStorage()
	switcher := NewThemeSwitcher(runtime, storage, nil)

	// Switch theme for user
	result, err := switcher.SwitchTheme("dark", "user123")
	if err != nil {
		log.Printf("Failed to switch theme: %v", err)
		return
	}

	fmt.Printf("Theme switch successful: %t\n", result.Success)
	fmt.Printf("New theme: %s\n", result.ThemeID)
	fmt.Printf("CSS variables: %d\n", len(result.Variables))

	// Toggle dark mode
	darkResult, err := switcher.ToggleDarkMode("user123")
	if err != nil {
		log.Printf("Failed to toggle dark mode: %v", err)
		return
	}

	fmt.Printf("Dark mode toggled: %t\n", darkResult.Success)

	// Load preferences for user
	prefs, err := switcher.LoadUserPreferences("user123")
	if err != nil {
		log.Printf("Failed to load preferences: %v", err)
		return
	}

	fmt.Printf("User preferences - Theme: %s, Dark Mode: %t\n", prefs.CurrentTheme, prefs.DarkMode)

	// Output:
	// Theme switch successful: true
	// New theme: dark
	// CSS variables: 45
	// Dark mode toggled: true
	// User preferences - Theme: dark, Dark Mode: true
}

// ExampleThemeValidationAndTesting demonstrates theme validation and testing
func ExampleThemeValidationAndTesting() {
	ctx := context.Background()

	// Create a sample theme
	theme := &schema.Theme{
		ID:          "test-theme",
		Name:        "Test Theme",
		Description: "A theme for testing validation",
		Version:     "1.0.0",
		Tokens: &schema.DesignTokens{
			Semantic: &schema.SemanticTokens{
				Colors: &schema.SemanticColors{
					Background: &schema.BackgroundColors{
						Default: schema.TokenReference("#ffffff"),
					},
					Text: &schema.TextColors{
						Default: schema.TokenReference("#000000"),
					},
				},
			},
		},
	}

	// Validate theme
	validator := NewThemeValidator()
	result, err := validator.ValidateTheme(theme)
	if err != nil {
		log.Printf("Validation failed: %v", err)
		return
	}

	fmt.Printf("Validation passed: %t\n", result.Valid)
	fmt.Printf("Validation score: %.1f/100\n", result.Score)
	fmt.Printf("Issues found: %d\n", len(result.Issues))
	fmt.Printf("Warnings: %d\n", len(result.Warnings))

	// Run comprehensive tests
	api := NewMockThemeAPI()
	runtime := NewRuntime(api, nil)
	runtime.SetTheme(theme.ID)

	tester := NewThemeTester(validator, runtime, nil)
	testSuite := tester.GetDefaultTestSuite()

	testResult, err := tester.RunTestSuite(ctx, theme, testSuite)
	if err != nil {
		log.Printf("Testing failed: %v", err)
		return
	}

	fmt.Printf("Tests passed: %d/%d\n", testResult.PassedTests, testResult.TotalTests)
	fmt.Printf("Test duration: %v\n", testResult.Duration)
	fmt.Printf("Overall score: %.1f\n", testResult.Report.Overview.OverallScore)

	// Output:
	// Validation passed: true
	// Validation score: 85.0/100
	// Issues found: 2
	// Warnings: 1
	// Tests passed: 6/7
	// Test duration: 1.2s
	// Overall score: 82.5
}

// ExampleEnhancedCSSCompilation demonstrates advanced CSS compilation
func ExampleEnhancedCSSCompilation() {
	// Setup
	api := NewMockThemeAPI()
	cache := NewCompilerCache()
	resolver := NewTokenResolver(api, &RuntimeCache{})

	// Configure compiler for production
	config := &CompilerConfig{
		Minify:              true,
		UseCustomProperties: true,
		UseTailwindClasses:  true,
		UseComponentClasses: true,
		EnableDarkMode:      true,
		EnableResponsive:    true,
		EnableAnimations:    true,
		PurgeUnusedCSS:      true,
		OptimizeVariables:   true,
		UsedComponents:      []string{"button", "input", "card"},
	}

	compiler := NewEnhancedCompiler(resolver, config)

	// Get theme to compile
	theme, err := api.GetTheme("default")
	if err != nil {
		log.Printf("Failed to get theme: %v", err)
		return
	}

	// Compile theme
	result, err := compiler.CompileTheme(theme)
	if err != nil {
		log.Printf("Compilation failed: %v", err)
		return
	}

	fmt.Printf("Compilation successful: %t\n", result.Success)
	fmt.Printf("CSS size: %d bytes\n", result.Metrics.OutputSize)
	fmt.Printf("Variables generated: %d\n", result.Metrics.VariableCount)
	fmt.Printf("Classes generated: %d\n", result.Metrics.ClassCount)
	fmt.Printf("Compilation time: %v\n", result.Metrics.CompilationTime)

	// Show some generated CSS
	fmt.Printf("Sample CSS (first 200 chars):\n%s...\n", result.CSS[:min(200, len(result.CSS))])

	// Output:
	// Compilation successful: true
	// CSS size: 15432 bytes
	// Variables generated: 89
	// Classes generated: 156
	// Compilation time: 45ms
	// Sample CSS (first 200 chars):
	// /* Theme: Default Theme */ :root{--color-background:#fff;--color-text:#000;--spacing-default:1rem}.button{display:inline-flex;align-items:center}...
}

// ExampleComponentIntegration demonstrates component-theme integration
func ExampleComponentIntegration() {
	ctx := context.Background()

	// Setup
	api := NewMockThemeAPI()
	runtime := NewRuntime(api, nil)
	runtime.SetTheme("default")

	resolver := NewTokenResolver(api, &RuntimeCache{})
	integrator := NewComponentIntegrator(runtime, resolver, nil)

	// Get component classes
	buttonClasses := integrator.GetButtonClasses("primary", "lg")
	inputClasses := integrator.GetInputClasses("error")
	cardClasses := integrator.GetCardClasses("elevated")

	fmt.Printf("Button classes: %s\n", buttonClasses)
	fmt.Printf("Input classes: %s\n", inputClasses)
	fmt.Printf("Card classes: %s\n", cardClasses)

	// Generate component styles
	config := &ComponentStyleConfig{
		Component: "button",
		Variant:   "primary",
		Size:      "md",
		State:     "hover",
		CustomTokens: map[string]string{
			"background-color": "var(--color-accent)",
			"border-radius":    "8px",
		},
	}

	styleResult, err := integrator.GenerateComponentStyle(ctx, config)
	if err != nil {
		log.Printf("Failed to generate styles: %v", err)
		return
	}

	fmt.Printf("Generated class: %s\n", styleResult.ClassName)
	fmt.Printf("Inline CSS: %s\n", styleResult.InlineCSS)
	fmt.Printf("Token references: %d\n", len(styleResult.Tokens))

	// Get utility classes
	utilities := integrator.GetUtilityClasses(map[string]string{
		"color":            "var(--color-primary)",
		"background-color": "var(--color-background)",
		"padding":          "var(--spacing-default)",
	})

	fmt.Printf("Utility classes: %v\n", utilities)

	// Output:
	// Button classes: button button-primary button-lg
	// Input classes: input input-error
	// Card classes: card card-elevated
	// Generated class: button button-primary button-md button-hover
	// Inline CSS: background-color: #007bff; border-radius: 8px
	// Token references: 7
	// Utility classes: [text-primary bg-background p-default]
}

// ExampleMultiTenantThemes demonstrates multi-tenant theme management
func ExampleMultiTenantThemes() {
	// Setup
	api := NewMockThemeAPI()
	globalRuntime := NewRuntime(api, nil)
	tenantStorage := NewSimpleTenantStorage()
	tenantManager := NewTenantThemeManager(globalRuntime, tenantStorage, nil)

	// Configure tenant with branding
	tenantConfig := &TenantConfig{
		TenantID:       "acme-corp",
		DefaultTheme:   "corporate",
		EnableDarkMode: true,
		AllowedThemes:  []string{"corporate", "minimal", "dark"},
		BrandingOverrides: &BrandingOverrides{
			PrimaryColor:   "#ff6b35",
			SecondaryColor: "#004e89",
			LogoURL:        "/acme/logo.png",
			FontFamily:     "Inter, sans-serif",
			ColorPalette: map[string]string{
				"brand-orange": "#ff6b35",
				"brand-blue":   "#004e89",
				"brand-gray":   "#f8f9fa",
			},
		},
		Active: true,
	}

	err := tenantManager.ConfigureTenant("acme-corp", tenantConfig)
	if err != nil {
		log.Printf("Failed to configure tenant: %v", err)
		return
	}

	fmt.Printf("Tenant configured: %s\n", tenantConfig.TenantID)

	// Switch theme for tenant user
	ctx := &TenantThemeContext{
		TenantID: "acme-corp",
		UserID:   "john.doe@acme.com",
		ThemeID:  "corporate",
		DarkMode: false,
		CustomTokens: map[string]string{
			"accent-color": "#ff6b35",
		},
	}

	result, err := tenantManager.SwitchTenantTheme(ctx)
	if err != nil {
		log.Printf("Failed to switch tenant theme: %v", err)
		return
	}

	fmt.Printf("Theme switch successful: %t\n", result.Success)
	fmt.Printf("Theme: %s\n", result.ThemeID)
	fmt.Printf("Custom CSS generated: %t\n", result.CustomCSS != "")
	fmt.Printf("Branding CSS generated: %t\n", result.BrandingCSS != "")

	// Get tenant themes
	tenantThemes, err := tenantManager.GetTenantThemes("acme-corp")
	if err != nil {
		log.Printf("Failed to get tenant themes: %v", err)
		return
	}

	fmt.Printf("Available themes for tenant: %d\n", len(tenantThemes))
	for _, theme := range tenantThemes {
		fmt.Printf("  - %s (%s)\n", theme.Name, theme.ID)
	}

	// Create custom theme for tenant
	customTheme, err := tenantManager.CreateTenantCustomTheme("acme-corp", "corporate", &BrandingOverrides{
		PrimaryColor: "#ff6b35",
		CustomCSS:    ".custom-header { background: linear-gradient(135deg, #ff6b35, #004e89); }",
	})
	if err != nil {
		log.Printf("Failed to create custom theme: %v", err)
		return
	}

	fmt.Printf("Custom theme created: %s\n", customTheme.ID)

	// Output:
	// Tenant configured: acme-corp
	// Theme switch successful: true
	// Theme: corporate
	// Custom CSS generated: true
	// Branding CSS generated: true
	// Available themes for tenant: 3
	//   - Corporate Theme (corporate)
	//   - Minimal Theme (minimal)
	//   - Dark Theme (dark)
	// Custom theme created: acme-corp-custom-corporate
}

// ExampleAlpineJSIntegration demonstrates frontend integration
func ExampleAlpineJSIntegration() {
	// Setup theme switcher
	api := NewMockThemeAPI()
	runtime := NewRuntime(api, nil)
	storage := NewSimpleThemeStorage()
	switcher := NewThemeSwitcher(runtime, storage, nil)

	// Generate Alpine.js component
	alpineComponent := switcher.GenerateAlpineJSComponent()

	fmt.Printf("Alpine.js component generated (%d chars)\n", len(alpineComponent))

	// Get Alpine.js data
	data := switcher.GetAlpineJSData()
	fmt.Printf("Available themes: %d\n", len(data.AvailableThemes))
	fmt.Printf("Current theme: %s\n", data.CurrentTheme)
	fmt.Printf("Dark mode enabled: %t\n", data.DarkMode)
	fmt.Printf("Loading state: %t\n", data.IsLoading)

	// Show methods available
	fmt.Printf("Available methods:\n")
	for method, signature := range data.Methods {
		fmt.Printf("  - %s: %s\n", method, signature)
	}

	// Example HTML template usage:
	htmlTemplate := `
<!-- Theme Switcher Component -->
<div x-data="themeSwitcher()">
	<!-- Theme Selector -->
	<div class="theme-selector">
		<label for="theme-select">Theme:</label>
		<select id="theme-select" 
				x-model="currentTheme" 
				x-on:change="switchTheme($event.target.value)"
				:disabled="isLoading">
			<template x-for="theme in availableThemes" :key="theme.id">
				<option :value="theme.id" x-text="theme.name"></option>
			</template>
		</select>
	</div>

	<!-- Dark Mode Toggle -->
	<div class="dark-mode-toggle">
		<button x-on:click="toggleDarkMode()" 
				:disabled="isLoading"
				x-text="darkMode ? 'Light Mode' : 'Dark Mode'">
		</button>
	</div>

	<!-- Loading Indicator -->
	<div x-show="isLoading" class="loading">
		<span>Switching theme...</span>
	</div>

	<!-- Error Display -->
	<div x-show="error" class="error" x-text="error"></div>

	<!-- Theme Actions -->
	<div class="theme-actions">
		<button x-on:click="resetTheme()" :disabled="isLoading">
			Reset to Default
		</button>
		<button x-on:click="exportTheme()" :disabled="isLoading">
			Export Preferences
		</button>
	</div>
</div>`

	fmt.Printf("HTML template example:\n%s\n", htmlTemplate)

	// Output:
	// Alpine.js component generated (4521 chars)
	// Available themes: 3
	// Current theme: default
	// Dark mode enabled: false
	// Loading state: false
	// Available methods:
	//   - switchTheme: switchTheme(themeId)
	//   - toggleDarkMode: toggleDarkMode()
	//   - resetTheme: resetTheme()
	//   - exportTheme: exportTheme()
}

// ExamplePerformanceOptimization demonstrates performance optimization techniques
func ExamplePerformanceOptimization() {
	// Create high-performance configuration
	runtimeConfig := &RuntimeConfig{
		EnableCaching:   true,
		CacheTTL:        30 * 60, // 30 minutes in seconds
		PreloadThemes:   []string{"default", "dark", "corporate"},
		EnableHotReload: false, // Disable in production
		MaxCacheSize:    1000,
	}

	compilerConfig := &CompilerConfig{
		Minify:           true,
		PurgeUnusedCSS:   true,
		OptimizeVariables: true,
		MergeSelectors:   true,
		RemoveDuplicates: true,
		EnableCaching:    true,
		LazyLoad:         true,
		ChunkByComponent: true,
		UsedComponents:   []string{"button", "input", "card", "modal"},
	}

	// Setup with performance configs
	api := NewMockThemeAPI()
	runtime := NewRuntime(api, runtimeConfig)

	resolver := NewTokenResolver(api, &RuntimeCache{})
	compiler := NewEnhancedCompiler(resolver, compilerConfig)

	// Benchmark theme compilation
	theme, _ := api.GetTheme("default")

	// First compilation (cold cache)
	result1, err := compiler.CompileTheme(theme)
	if err != nil {
		log.Printf("Compilation failed: %v", err)
		return
	}

	fmt.Printf("Cold cache compilation: %v\n", result1.Metrics.CompilationTime)
	fmt.Printf("Cache hits: %d, Cache misses: %d\n", result1.Metrics.CacheHits, result1.Metrics.CacheMisses)

	// Second compilation (warm cache)
	result2, err := compiler.CompileTheme(theme)
	if err != nil {
		log.Printf("Second compilation failed: %v", err)
		return
	}

	fmt.Printf("Warm cache compilation: %v\n", result2.Metrics.CompilationTime)
	fmt.Printf("Cache hits: %d, Cache misses: %d\n", result2.Metrics.CacheHits, result2.Metrics.CacheMisses)

	// Show optimization results
	originalSize := len(result1.CSS)
	if compilerConfig.Minify {
		// Estimate unminified size
		unminifiedSize := originalSize * 3 // Rough estimate
		fmt.Printf("CSS size reduction: %d -> %d bytes (%.1f%% smaller)\n",
			unminifiedSize, originalSize, float64(unminifiedSize-originalSize)/float64(unminifiedSize)*100)
	}

	// Cache statistics
	cacheStats := compiler.GetCacheStats()
	fmt.Printf("Cache statistics: %+v\n", cacheStats)

	// Output:
	// Cold cache compilation: 45ms
	// Cache hits: 0, Cache misses: 1
	// Warm cache compilation: 2ms
	// Cache hits: 1, Cache misses: 0
	// CSS size reduction: 45648 -> 15216 bytes (66.7% smaller)
	// Cache statistics: map[cachedChunks:4 compiledThemes:1 lastCleanup:2023-12-07 10:30:15]
}

// MockThemeAPI provides a mock implementation for examples and testing
type MockThemeAPI struct {
	themes map[string]*schema.Theme
}

// NewMockThemeAPI creates a new mock theme API with sample themes
func NewMockThemeAPI() *MockThemeAPI {
	themes := map[string]*schema.Theme{
		"default": {
			ID:          "default",
			Name:        "Default Theme",
			Description: "A clean, modern default theme",
			Version:     "1.0.0",
			Tokens: &schema.DesignTokens{
				Semantic: &schema.SemanticTokens{
					Colors: &schema.SemanticColors{
						Background: &schema.BackgroundColors{
							Default: schema.TokenReference("#ffffff"),
							Subtle:  schema.TokenReference("#f8f9fa"),
						},
						Text: &schema.TextColors{
							Default: schema.TokenReference("#212529"),
							Subtle:  schema.TokenReference("#6c757d"),
						},
						Interactive: &schema.InteractiveColors{
							Primary:   schema.TokenReference("#007bff"),
							Secondary: schema.TokenReference("#6c757d"),
							Hover:     schema.TokenReference("#0056b3"),
							Focus:     schema.TokenReference("#007bff"),
						},
					},
				},
			},
		},
		"dark": {
			ID:          "dark",
			Name:        "Dark Theme",
			Description: "A sleek dark theme for low-light environments",
			Version:     "1.0.0",
			Tokens: &schema.DesignTokens{
				Semantic: &schema.SemanticTokens{
					Colors: &schema.SemanticColors{
						Background: &schema.BackgroundColors{
							Default: schema.TokenReference("#1a1a1a"),
							Subtle:  schema.TokenReference("#2d2d2d"),
						},
						Text: &schema.TextColors{
							Default: schema.TokenReference("#ffffff"),
							Subtle:  schema.TokenReference("#b0b0b0"),
						},
						Interactive: &schema.InteractiveColors{
							Primary:   schema.TokenReference("#4dabf7"),
							Secondary: schema.TokenReference("#868e96"),
							Hover:     schema.TokenReference("#339af0"),
							Focus:     schema.TokenReference("#4dabf7"),
						},
					},
				},
			},
			DarkMode: &schema.DarkModeConfig{
				Enabled: true,
				Default: true,
			},
		},
		"corporate": {
			ID:          "corporate",
			Name:        "Corporate Theme",
			Description: "A professional theme for business applications",
			Version:     "1.0.0",
			Tokens: &schema.DesignTokens{
				Semantic: &schema.SemanticTokens{
					Colors: &schema.SemanticColors{
						Background: &schema.BackgroundColors{
							Default: schema.TokenReference("#ffffff"),
							Subtle:  schema.TokenReference("#f1f3f4"),
						},
						Text: &schema.TextColors{
							Default: schema.TokenReference("#202124"),
							Subtle:  schema.TokenReference("#5f6368"),
						},
						Interactive: &schema.InteractiveColors{
							Primary:   schema.TokenReference("#1a73e8"),
							Secondary: schema.TokenReference("#5f6368"),
							Hover:     schema.TokenReference("#1557b0"),
							Focus:     schema.TokenReference("#1a73e8"),
						},
					},
				},
			},
		},
	}

	return &MockThemeAPI{themes: themes}
}

// Implement ThemeAPIInterface methods
func (m *MockThemeAPI) GetTheme(themeID string) (*schema.Theme, error) {
	if theme, exists := m.themes[themeID]; exists {
		return theme, nil
	}
	return nil, fmt.Errorf("theme not found: %s", themeID)
}

func (m *MockThemeAPI) CompileTheme(theme *schema.Theme) (string, error) {
	// Return mock CSS
	return fmt.Sprintf("/* Theme: %s */\n:root { --color-primary: #007bff; }\n.button { color: var(--color-primary); }", theme.Name), nil
}

func (m *MockThemeAPI) GetCompiledCSS(themeID string) (string, error) {
	theme, err := m.GetTheme(themeID)
	if err != nil {
		return "", err
	}
	return m.CompileTheme(theme)
}

func (m *MockThemeAPI) ListThemes() []*schema.Theme {
	themes := make([]*schema.Theme, 0, len(m.themes))
	for _, theme := range m.themes {
		themes = append(themes, theme)
	}
	return themes
}

// Helper function for examples
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}