// Package theme provides comprehensive theme integration utilities for the Ruun design system.
//
// This package offers runtime theme management, compilation, validation, testing, and multi-tenant
// support for theme-based applications. It seamlessly integrates with the existing schema.ThemeAPI
// and provides enhanced capabilities for production applications.
//
// # Key Features
//
//   - Runtime theme switching without page reloads
//   - Advanced CSS compilation with token resolution
//   - Comprehensive theme validation and testing
//   - Multi-tenant theme isolation and customization
//   - Component integration utilities
//   - Alpine.js integration for frontend theme switching
//   - Caching and performance optimizations
//   - Dark mode support
//   - Responsive theme utilities
//
// # Architecture
//
// The package is organized into several key components:
//
//   - Runtime: Core theme runtime with caching and performance features
//   - Compiler: Enhanced CSS compilation with token support
//   - Resolver: Token reference resolution with validation
//   - Validator: Comprehensive theme validation and testing
//   - Switcher: Theme switching service with persistence
//   - Integration: Component-theme integration utilities
//   - Tenant: Multi-tenant theme management
//
// # Basic Usage
//
// Create a theme runtime:
//
//	api := ruun.NewThemeAPI("./themes")
//	runtime := theme.NewRuntime(api, nil)
//
//	// Switch to a theme
//	err := runtime.SetTheme("my-theme")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get compiled CSS
//	css, err := runtime.GetThemeCSS()
//	if err != nil {
//		log.Fatal(err)
//	}
//
// # Advanced Usage
//
// Theme switching with persistence:
//
//	// Create theme switcher
//	storage := theme.NewSimpleThemeStorage()
//	switcher := theme.NewThemeSwitcher(runtime, storage, nil)
//
//	// Switch theme for a user
//	result, err := switcher.SwitchTheme("new-theme", "user123")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Generate Alpine.js component
//	alpineComponent := switcher.GenerateAlpineJSComponent()
//
// Theme validation and testing:
//
//	// Create validator
//	validator := theme.NewThemeValidator()
//	
//	// Validate theme
//	result, err := validator.ValidateTheme(myTheme)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Run comprehensive tests
//	tester := theme.NewThemeTester(validator, runtime, nil)
//	suite := tester.GetDefaultTestSuite()
//	testResult, err := tester.RunTestSuite(ctx, myTheme, suite)
//
// Enhanced CSS compilation:
//
//	// Create enhanced compiler
//	resolver := theme.NewTokenResolver(api, cache)
//	compiler := theme.NewEnhancedCompiler(resolver, nil)
//
//	// Compile theme with advanced features
//	result, err := compiler.CompileTheme(myTheme)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Get compiled CSS with utilities and components
//	css := result.CSS
//	variables := result.Variables
//	classes := result.Classes
//
// Component integration:
//
//	// Create component integrator
//	integrator := theme.NewComponentIntegrator(runtime, resolver, nil)
//
//	// Get component classes
//	buttonClasses := integrator.GetButtonClasses("primary", "lg")
//	inputClasses := integrator.GetInputClasses("error")
//
//	// Generate component styles
//	config := &theme.ComponentStyleConfig{
//		Component: "button",
//		Variant:   "primary",
//		Size:      "md",
//	}
//	styleResult, err := integrator.GenerateComponentStyle(ctx, config)
//
// Multi-tenant support:
//
//	// Create tenant manager
//	tenantStorage := theme.NewSimpleTenantStorage()
//	tenantManager := theme.NewTenantThemeManager(runtime, tenantStorage, nil)
//
//	// Configure tenant
//	tenantConfig := &theme.TenantConfig{
//		TenantID:       "tenant123",
//		DefaultTheme:   "corporate",
//		EnableDarkMode: true,
//		BrandingOverrides: &theme.BrandingOverrides{
//			PrimaryColor: "#ff6b35",
//			LogoURL:      "/tenant/logo.png",
//		},
//	}
//	err := tenantManager.ConfigureTenant("tenant123", tenantConfig)
//
//	// Switch tenant theme
//	ctx := &theme.TenantThemeContext{
//		TenantID: "tenant123",
//		UserID:   "user456",
//		ThemeID:  "corporate",
//		DarkMode: false,
//	}
//	result, err := tenantManager.SwitchTenantTheme(ctx)
//
// # Frontend Integration
//
// The package provides seamless integration with Alpine.js for frontend theme switching:
//
//	<!-- Theme switcher component -->
//	<div x-data="themeSwitcher()">
//		<!-- Theme selector -->
//		<select x-model="currentTheme" x-on:change="switchTheme($event.target.value)">
//			<template x-for="theme in availableThemes">
//				<option :value="theme.id" x-text="theme.name"></option>
//			</template>
//		</select>
//
//		<!-- Dark mode toggle -->
//		<button x-on:click="toggleDarkMode()" 
//				x-text="darkMode ? 'Light Mode' : 'Dark Mode'">
//		</button>
//
//		<!-- Loading state -->
//		<div x-show="isLoading">Switching theme...</div>
//
//		<!-- Error display -->
//		<div x-show="error" x-text="error"></div>
//	</div>
//
// The Alpine.js component provides:
//   - Reactive theme switching
//   - Dark mode toggle
//   - Loading states
//   - Error handling
//   - Preference persistence
//   - System theme detection
//
// # Performance Considerations
//
// The theme package is designed for high performance:
//
//   - Intelligent caching at multiple levels
//   - Lazy loading of theme resources
//   - CSS minification and optimization
//   - Token resolution caching
//   - Compilation result caching
//   - Multi-level cache invalidation
//
// # Security
//
// Security considerations include:
//
//   - Token validation to prevent injection
//   - Tenant isolation for multi-tenant environments
//   - Permission-based theme access
//   - Sanitization of custom CSS
//   - Validation of theme uploads
//
// # Error Handling
//
// The package provides comprehensive error handling:
//
//   - Structured error types with context
//   - Graceful fallbacks to default themes
//   - Validation errors with suggestions
//   - Recovery mechanisms for failed operations
//   - Detailed error reporting for debugging
//
// # Extending the Package
//
// The package is designed for extensibility:
//
//   - Interface-based design for easy mocking and testing
//   - Plugin architecture for custom validators
//   - Configurable compilation pipeline
//   - Custom storage backends
//   - Theme observer pattern for notifications
//
// # Best Practices
//
//   - Always validate themes before deployment
//   - Use caching in production environments
//   - Implement proper error handling and fallbacks
//   - Test theme switching thoroughly
//   - Monitor performance metrics
//   - Keep themes organized and documented
//   - Use semantic versioning for theme updates
//   - Implement proper tenant isolation
//
// For more detailed examples and documentation, see the examples directory
// and the comprehensive test suite.
package theme