// Package validation provides comprehensive validation mechanisms for atomic design systems
//
// This package includes:
//   - Component validation (props, types, composition)
//   - Schema validation for forms, themes, and configurations  
//   - Accessibility validation (ARIA, contrast, keyboard nav)
//   - Performance validation (bundle size, render time)
//   - Theme validation (token usage, consistency)
//   - Runtime validation for user inputs and API calls
//   - Testing framework for validation systems
//   - Build process integration
//
// Example usage:
//
//	// Create a validation engine
//	engine := validation.NewValidationEngine(validation.ValidationOptions{
//		StrictMode: true,
//		EnableA11y: true,
//		EnablePerformance: true,
//		EnableTheme: true,
//	})
//
//	// Validate a component
//	component := &validation.ComponentInstance{
//		Type: "Button",
//		Props: map[string]interface{}{
//			"variant": "primary",
//			"size": "md",
//			"children": "Click me",
//		},
//	}
//
//	ctx := validation.NewValidationContext(context.Background(), validation.ValidationLevelError, "component")
//	result := engine.ValidateComponent(component, "Button")
//
//	if !result.IsValid() {
//		for _, err := range result.Errors {
//			fmt.Printf("Validation error: %s\n", err.Message)
//		}
//	}
//
//	// Set up build integration
//	buildIntegration := validation.NewBuildIntegration()
//	buildIntegration.SetupGitHooks()
//	buildIntegration.StartFileWatching()
//
//	// Add validation pipeline
//	pipeline := validation.BuildPipeline{
//		Name: "validation-pipeline",
//		Stages: []validation.BuildStage{
//			{
//				Name: "component-validation",
//				Type: validation.BuildStageTypeValidation,
//				Commands: []validation.BuildCommand{
//					{
//						Name: "validate-components",
//						Command: "validation-cli",
//						Args: []string{"validate", "--type", "component"},
//					},
//				},
//			},
//		},
//	}
//	buildIntegration.AddPipeline(pipeline)
//
package validation

import (
	"context"
	"fmt"
	"time"
)

// Version of the validation package
const Version = "1.0.0"

// ValidationSuite provides a unified interface for all validation functionality
type ValidationSuite struct {
	engine         *ValidationEngine
	componentValidator *ComponentValidator
	schemaValidator    *SchemaValidator
	a11yValidator      *AccessibilityValidator
	performanceValidator *PerformanceValidator
	themeValidator     *ThemeValidator
	runtimeValidator   *RuntimeValidator
	testFramework      *TestFramework
	buildIntegration   *BuildIntegration
	config            ValidationSuiteConfig
}

// ValidationSuiteConfig configures the validation suite
type ValidationSuiteConfig struct {
	// Core validation options
	ValidationOptions ValidationOptions `json:"validationOptions"`
	
	// Component validation
	EnableComponentValidation bool `json:"enableComponentValidation"`
	ComponentSchemas         map[string]*ComponentSchema `json:"componentSchemas,omitempty"`
	
	// Schema validation
	EnableSchemaValidation bool `json:"enableSchemaValidation"`
	SchemaFormats         []string `json:"schemaFormats,omitempty"`
	
	// Accessibility validation
	EnableA11yValidation bool `json:"enableA11yValidation"`
	A11yConfig          A11yConfig `json:"a11yConfig"`
	
	// Performance validation
	EnablePerformanceValidation bool `json:"enablePerformanceValidation"`
	PerformanceConfig          PerformanceConfig `json:"performanceConfig"`
	
	// Theme validation
	EnableThemeValidation bool `json:"enableThemeValidation"`
	ThemeConfig          ThemeValidationConfig `json:"themeConfig"`
	
	// Runtime validation
	EnableRuntimeValidation bool `json:"enableRuntimeValidation"`
	RuntimeConfig          RuntimeValidationConfig `json:"runtimeConfig"`
	
	// Testing
	EnableTesting bool `json:"enableTesting"`
	TestConfig   TestFrameworkConfig `json:"testConfig"`
	
	// Build integration
	EnableBuildIntegration bool `json:"enableBuildIntegration"`
	BuildConfig           BuildIntegrationConfig `json:"buildConfig"`
	
	// Global settings
	CacheEnabled   bool          `json:"cacheEnabled"`
	CacheTTL      time.Duration `json:"cacheTTL"`
	LogLevel      string        `json:"logLevel"` // debug, info, warn, error
	ReportFormats []string      `json:"reportFormats"` // json, html, junit
	OutputPath    string        `json:"outputPath"`
}

// NewValidationSuite creates a comprehensive validation suite
func NewValidationSuite(config ...ValidationSuiteConfig) *ValidationSuite {
	var cfg ValidationSuiteConfig
	if len(config) > 0 {
		cfg = config[0]
	} else {
		cfg = DefaultValidationSuiteConfig()
	}

	suite := &ValidationSuite{
		config: cfg,
	}

	// Initialize core validation engine
	if cfg.EnableComponentValidation || cfg.EnableSchemaValidation {
		suite.engine = NewValidationEngine(cfg.ValidationOptions)
	}

	// Initialize specialized validators
	if cfg.EnableComponentValidation {
		suite.componentValidator = NewComponentValidator()
		// Register component schemas if provided
		for name, schema := range cfg.ComponentSchemas {
			suite.componentValidator.RegisterComponentSchema(schema)
			_ = name // Use the name for registration
		}
	}

	if cfg.EnableSchemaValidation {
		suite.schemaValidator = NewSchemaValidator()
	}

	if cfg.EnableA11yValidation {
		suite.a11yValidator = NewAccessibilityValidator()
		suite.a11yValidator.config = cfg.A11yConfig
	}

	if cfg.EnablePerformanceValidation {
		suite.performanceValidator = NewPerformanceValidator()
		suite.performanceValidator.config = cfg.PerformanceConfig
	}

	if cfg.EnableThemeValidation {
		suite.themeValidator = NewThemeValidator()
		suite.themeValidator.config = cfg.ThemeConfig
	}

	if cfg.EnableRuntimeValidation {
		suite.runtimeValidator = NewRuntimeValidator()
		suite.runtimeValidator.config = cfg.RuntimeConfig
	}

	if cfg.EnableTesting {
		suite.testFramework = NewTestFramework()
		suite.testFramework.config = cfg.TestConfig
	}

	if cfg.EnableBuildIntegration {
		suite.buildIntegration = NewBuildIntegration()
		suite.buildIntegration.config = cfg.BuildConfig
	}

	return suite
}

// DefaultValidationSuiteConfig returns a default configuration
func DefaultValidationSuiteConfig() ValidationSuiteConfig {
	return ValidationSuiteConfig{
		ValidationOptions: ValidationOptions{
			StrictMode:        false,
			EnableMetrics:     true,
			EnableA11y:        true,
			EnableTheme:       true,
			EnablePerformance: true,
			Timeout:          time.Minute * 5,
		},
		EnableComponentValidation:   true,
		EnableSchemaValidation:      true,
		EnableA11yValidation:        true,
		EnablePerformanceValidation: true,
		EnableThemeValidation:       true,
		EnableRuntimeValidation:     true,
		EnableTesting:              true,
		EnableBuildIntegration:      false, // Disabled by default
		A11yConfig: A11yConfig{
			Level:     A11yLevelAA,
			Standards: []A11yStandard{A11yStandardWCAG21},
			ColorContrast: ContrastConfig{
				MinContrastNormal: 4.5,
				MinContrastLarge:  3.0,
				MinContrastAAA:    7.0,
			},
		},
		PerformanceConfig: PerformanceConfig{
			EnableBundleSize:   true,
			EnableRenderTime:   true,
			EnableMemoryUsage:  true,
			EnableComplexity:   true,
			Thresholds: PerformanceThresholds{
				MaxBundleSize:  5 * 1024 * 1024, // 5MB
				MaxRenderTime:  time.Millisecond * 100,
				MaxMemoryUsage: 100 * 1024 * 1024, // 100MB
				MaxComplexity:  50,
			},
		},
		ThemeConfig: ThemeValidationConfig{
			RequireConsistency:  true,
			CheckTokenUsage:     true,
			ValidateCompleteness: true,
			CheckAccessibility:  true,
			ValidatePerformance: true,
			MinContrastRatio:   4.5,
			MaxTokenCount:      500,
			RequiredCategories: []string{"colors", "spacing", "typography"},
		},
		RuntimeConfig: RuntimeValidationConfig{
			EnableInputValidation: true,
			EnableAPIValidation:   true,
			EnableSanitization:    true,
			EnableRateLimit:      true,
			EnableCaching:        true,
			ValidationTimeout:    time.Second * 5,
		},
		TestConfig: TestFrameworkConfig{
			EnableParallelExecution: true,
			EnableBenchmarking:      true,
			EnableCoverage:          true,
			TestTimeout:            time.Minute * 5,
			ReportFormat:           "json",
		},
		BuildConfig: BuildIntegrationConfig{
			EnableFileWatching:   true,
			EnablePreCommitHooks: true,
			EnableCIPipeline:     true,
			EnableCaching:        true,
			WatchDirectories:     []string{"./src", "./components", "./schemas"},
			ValidationTimeout:    time.Minute * 10,
		},
		CacheEnabled:  true,
		CacheTTL:     time.Hour,
		LogLevel:     "info",
		ReportFormats: []string{"json", "html"},
		OutputPath:   "./validation-reports",
	}
}

// ValidateAll performs comprehensive validation across all enabled validators
func (suite *ValidationSuite) ValidateAll(ctx context.Context, data interface{}) (*ComprehensiveValidationResult, error) {
	result := &ComprehensiveValidationResult{
		Overall:   true,
		StartTime: time.Now(),
		Results:   make(map[string]*ValidationResult),
		Metadata:  make(map[string]interface{}),
	}

	// Component validation
	if suite.config.EnableComponentValidation && suite.componentValidator != nil {
		if component, ok := data.(*ComponentInstance); ok {
			compResult := suite.componentValidator.ValidateComponent(component)
			result.Results["component"] = compResult
			if !compResult.Valid {
				result.Overall = false
			}
		}
	}

	// Schema validation
	if suite.config.EnableSchemaValidation && suite.schemaValidator != nil {
		schemaResult := suite.schemaValidator.ValidateSchema(data)
		result.Results["schema"] = suite.convertSchemaValidationResult(schemaResult)
		if !schemaResult.Valid {
			result.Overall = false
		}
	}

	// Accessibility validation
	if suite.config.EnableA11yValidation && suite.a11yValidator != nil {
		validationCtx := NewValidationContext(ctx, ValidationLevelError, "comprehensive")
		a11yResult := suite.a11yValidator.ValidateAccessibility(validationCtx, data)
		result.Results["accessibility"] = suite.convertA11yValidationResult(a11yResult)
		if !a11yResult.Compliant {
			result.Overall = false
		}
	}

	// Performance validation
	if suite.config.EnablePerformanceValidation && suite.performanceValidator != nil {
		validationCtx := NewValidationContext(ctx, ValidationLevelError, "comprehensive")
		perfResult := suite.performanceValidator.ValidatePerformance(validationCtx, data)
		result.Results["performance"] = suite.convertPerformanceValidationResult(perfResult)
		if !perfResult.MeetsThresholds {
			result.Overall = false
		}
	}

	// Theme validation
	if suite.config.EnableThemeValidation && suite.themeValidator != nil {
		validationCtx := NewValidationContext(ctx, ValidationLevelError, "comprehensive")
		themeResult := suite.themeValidator.ValidateTheme(validationCtx, data)
		result.Results["theme"] = suite.convertThemeValidationResult(themeResult)
		if !themeResult.Valid {
			result.Overall = false
		}
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)

	return result, nil
}

// ComprehensiveValidationResult contains results from all validators
type ComprehensiveValidationResult struct {
	Overall   bool                            `json:"overall"`
	Results   map[string]*ValidationResult    `json:"results"`
	StartTime time.Time                      `json:"startTime"`
	EndTime   time.Time                      `json:"endTime"`
	Duration  time.Duration                  `json:"duration"`
	Metadata  map[string]interface{}         `json:"metadata"`
}

// ValidateComponent validates a component using the component validator
func (suite *ValidationSuite) ValidateComponent(component *ComponentInstance) (*ValidationResult, error) {
	if !suite.config.EnableComponentValidation || suite.componentValidator == nil {
		return nil, fmt.Errorf("component validation is not enabled")
	}

	result := suite.componentValidator.ValidateComponent(component)
	return result, nil
}

// ValidateSchema validates a schema using the schema validator
func (suite *ValidationSuite) ValidateSchema(schema interface{}) (*SchemaValidationResult, error) {
	if !suite.config.EnableSchemaValidation || suite.schemaValidator == nil {
		return nil, fmt.Errorf("schema validation is not enabled")
	}

	return suite.schemaValidator.ValidateSchema(schema), nil
}

// ValidateAccessibility validates accessibility using the a11y validator
func (suite *ValidationSuite) ValidateAccessibility(ctx context.Context, element interface{}) (*AccessibilityResult, error) {
	if !suite.config.EnableA11yValidation || suite.a11yValidator == nil {
		return nil, fmt.Errorf("accessibility validation is not enabled")
	}

	validationCtx := NewValidationContext(ctx, ValidationLevelError, "accessibility")
	return suite.a11yValidator.ValidateAccessibility(validationCtx, element), nil
}

// ValidatePerformance validates performance characteristics
func (suite *ValidationSuite) ValidatePerformance(ctx context.Context, data interface{}) (*PerformanceMetrics, error) {
	if !suite.config.EnablePerformanceValidation || suite.performanceValidator == nil {
		return nil, fmt.Errorf("performance validation is not enabled")
	}

	validationCtx := NewValidationContext(ctx, ValidationLevelError, "performance")
	return suite.performanceValidator.ValidatePerformance(validationCtx, data), nil
}

// ValidateTheme validates theme definitions
func (suite *ValidationSuite) ValidateTheme(ctx context.Context, theme interface{}) (*ThemeValidationResult, error) {
	if !suite.config.EnableThemeValidation || suite.themeValidator == nil {
		return nil, fmt.Errorf("theme validation is not enabled")
	}

	validationCtx := NewValidationContext(ctx, ValidationLevelError, "theme")
	return suite.themeValidator.ValidateTheme(validationCtx, theme), nil
}

// ValidateRuntime validates runtime data
func (suite *ValidationSuite) ValidateRuntime(ctx context.Context, data interface{}, schema interface{}) (*RuntimeValidationResponse, error) {
	if !suite.config.EnableRuntimeValidation || suite.runtimeValidator == nil {
		return nil, fmt.Errorf("runtime validation is not enabled")
	}

	return suite.runtimeValidator.ValidateInput(ctx, data, schema), nil
}

// RunTests executes the test suite
func (suite *ValidationSuite) RunTests(ctx context.Context) (*TestSummary, error) {
	if !suite.config.EnableTesting || suite.testFramework == nil {
		return nil, fmt.Errorf("testing is not enabled")
	}

	return suite.testFramework.RunTests(ctx)
}

// SetupBuildIntegration sets up build process integration
func (suite *ValidationSuite) SetupBuildIntegration() error {
	if !suite.config.EnableBuildIntegration || suite.buildIntegration == nil {
		return fmt.Errorf("build integration is not enabled")
	}

	// Setup Git hooks
	if err := suite.buildIntegration.SetupGitHooks(); err != nil {
		return fmt.Errorf("failed to setup Git hooks: %w", err)
	}

	// Start file watching
	if err := suite.buildIntegration.StartFileWatching(); err != nil {
		return fmt.Errorf("failed to start file watching: %w", err)
	}

	return nil
}

// AddComponentSchema registers a component schema for validation
func (suite *ValidationSuite) AddComponentSchema(name string, schema *ComponentSchema) error {
	if suite.componentValidator == nil {
		return fmt.Errorf("component validator is not initialized")
	}

	schema.Name = name
	suite.componentValidator.RegisterComponentSchema(schema)
	return nil
}

// AddBuildPipeline adds a build pipeline
func (suite *ValidationSuite) AddBuildPipeline(pipeline BuildPipeline) error {
	if suite.buildIntegration == nil {
		return fmt.Errorf("build integration is not enabled")
	}

	suite.buildIntegration.AddPipeline(pipeline)
	return nil
}

// AddTestSuite adds a test suite
func (suite *ValidationSuite) AddTestSuite(testSuite TestSuite) error {
	if suite.testFramework == nil {
		return fmt.Errorf("test framework is not enabled")
	}

	suite.testFramework.AddTestSuite(testSuite)
	return nil
}

// GenerateReport generates a comprehensive validation report
func (suite *ValidationSuite) GenerateReport() (string, error) {
	// Implementation would combine reports from all validators
	report := map[string]interface{}{
		"version":   Version,
		"timestamp": time.Now(),
		"config":    suite.config,
	}

	// Add reports from each validator if available
	if suite.engine != nil {
		report["engine"] = suite.engine.GetValidationReport()
	}

	// Convert to JSON
	data, err := json.MarshalIndent(report, "", "  ")
	return string(data), err
}

// Cleanup performs cleanup operations
func (suite *ValidationSuite) Cleanup() error {
	// Stop file watching
	if suite.buildIntegration != nil {
		suite.buildIntegration.StopFileWatching()
	}

	return nil
}

// Helper methods for result conversion

func (suite *ValidationSuite) convertSchemaValidationResult(result *SchemaValidationResult) *ValidationResult {
	validationResult := &ValidationResult{
		Valid:     result.Valid,
		Level:     ValidationLevelError,
		Timestamp: result.Timestamp,
		Metadata:  result.Metadata,
	}

	for _, err := range result.Errors {
		validationResult.Errors = append(validationResult.Errors, ValidationError{
			Code:      err.Code,
			Message:   err.Message,
			Field:     err.Path,
			Level:     err.Level,
			Source:    "schema",
			Timestamp: time.Now(),
		})
	}

	return validationResult
}

func (suite *ValidationSuite) convertA11yValidationResult(result *AccessibilityResult) *ValidationResult {
	validationResult := &ValidationResult{
		Valid:     result.Compliant,
		Level:     ValidationLevel(result.Level),
		Timestamp: result.Timestamp,
		Metadata:  map[string]interface{}{"score": result.Score},
	}

	for _, violation := range result.Violations {
		validationResult.Errors = append(validationResult.Errors, ValidationError{
			Code:      violation.Code,
			Message:   violation.Message,
			Field:     violation.Element,
			Level:     ValidationLevel(violation.Level),
			Source:    "accessibility",
			Timestamp: time.Now(),
		})
	}

	return validationResult
}

func (suite *ValidationSuite) convertPerformanceValidationResult(result *PerformanceMetrics) *ValidationResult {
	return &ValidationResult{
		Valid:     result.MeetsThresholds,
		Level:     ValidationLevelError,
		Timestamp: result.Timestamp,
		Metadata:  map[string]interface{}{"score": result.Score, "grade": result.Grade},
	}
}

func (suite *ValidationSuite) convertThemeValidationResult(result *ThemeValidationResult) *ValidationResult {
	validationResult := &ValidationResult{
		Valid:     result.Valid,
		Level:     ValidationLevelError,
		Timestamp: result.Timestamp,
		Metadata:  map[string]interface{}{"score": result.Score},
	}

	for _, violation := range result.Violations {
		validationResult.Errors = append(validationResult.Errors, ValidationError{
			Code:      violation.Code,
			Message:   violation.Message,
			Field:     violation.Token,
			Level:     violation.Severity,
			Source:    "theme",
			Timestamp: time.Now(),
		})
	}

	return validationResult
}

// Package-level convenience functions

// ValidateComponent is a package-level convenience function
func ValidateComponent(component *ComponentInstance) (*ValidationResult, error) {
	suite := NewValidationSuite()
	return suite.ValidateComponent(component)
}

// ValidateSchema is a package-level convenience function
func ValidateSchema(schema interface{}) (*SchemaValidationResult, error) {
	suite := NewValidationSuite()
	return suite.ValidateSchema(schema)
}

// ValidateAccessibility is a package-level convenience function
func ValidateAccessibility(ctx context.Context, element interface{}) (*AccessibilityResult, error) {
	suite := NewValidationSuite()
	return suite.ValidateAccessibility(ctx, element)
}

// ValidateTheme is a package-level convenience function
func ValidateTheme(ctx context.Context, theme interface{}) (*ThemeValidationResult, error) {
	suite := NewValidationSuite()
	return suite.ValidateTheme(ctx, theme)
}

// QuickValidate performs a quick validation with default settings
func QuickValidate(ctx context.Context, data interface{}) (*ComprehensiveValidationResult, error) {
	config := DefaultValidationSuiteConfig()
	config.EnableBuildIntegration = false // Disable for quick validation
	
	suite := NewValidationSuite(config)
	return suite.ValidateAll(ctx, data)
}