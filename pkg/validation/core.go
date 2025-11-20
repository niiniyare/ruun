package validation

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ValidationLevel defines the severity of validation
type ValidationLevel string

const (
	ValidationLevelWarn   ValidationLevel = "warn"
	ValidationLevelError  ValidationLevel = "error"
	ValidationLevelStrict ValidationLevel = "strict"
)

// ValidationResult represents the outcome of a validation operation
type ValidationResult struct {
	Valid        bool                       `json:"valid"`
	Level        ValidationLevel           `json:"level"`
	Errors       []ValidationError         `json:"errors,omitempty"`
	Warnings     []ValidationWarning       `json:"warnings,omitempty"`
	Performance  *PerformanceMetrics       `json:"performance,omitempty"`
	Accessibility *AccessibilityResult     `json:"accessibility,omitempty"`
	Theme        *ThemeValidationResult    `json:"theme,omitempty"`
	Metadata     map[string]interface{}    `json:"metadata,omitempty"`
	Timestamp    time.Time                 `json:"timestamp"`
}

// ValidationError represents a validation error
type ValidationError struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Field      string                 `json:"field,omitempty"`
	Component  string                 `json:"component,omitempty"`
	Level      ValidationLevel        `json:"level"`
	Source     string                 `json:"source"`
	Location   *SourceLocation        `json:"location,omitempty"`
	Suggestion string                 `json:"suggestion,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
	Timestamp  time.Time             `json:"timestamp"`
}

// ValidationWarning represents a validation warning
type ValidationWarning struct {
	Code      string                 `json:"code"`
	Message   string                 `json:"message"`
	Field     string                 `json:"field,omitempty"`
	Component string                 `json:"component,omitempty"`
	Source    string                 `json:"source"`
	Location  *SourceLocation        `json:"location,omitempty"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

// SourceLocation represents the location where validation occurred
type SourceLocation struct {
	File   string `json:"file,omitempty"`
	Line   int    `json:"line,omitempty"`
	Column int    `json:"column,omitempty"`
	Path   string `json:"path,omitempty"` // JSONPath or component path
}

// ValidationContext provides context for validation operations
type ValidationContext struct {
	ctx       context.Context
	Level     ValidationLevel           `json:"level"`
	Component string                   `json:"component,omitempty"`
	Source    string                   `json:"source"`
	Metadata  map[string]interface{}   `json:"metadata,omitempty"`
	Rules     []ValidationRule         `json:"rules,omitempty"`
	Options   ValidationOptions        `json:"options"`
}

// ValidationOptions configure validation behavior
type ValidationOptions struct {
	StrictMode       bool     `json:"strictMode"`
	SkipWarnings     bool     `json:"skipWarnings"`
	EnableMetrics    bool     `json:"enableMetrics"`
	EnableA11y       bool     `json:"enableA11y"`
	EnableTheme      bool     `json:"enableTheme"`
	EnablePerformance bool    `json:"enablePerformance"`
	CustomRules      []string `json:"customRules,omitempty"`
	ExcludeRules     []string `json:"excludeRules,omitempty"`
	Timeout          time.Duration `json:"timeout,omitempty"`
}

// ValidationRule defines a validation rule
type ValidationRule struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Category    string                 `json:"category"` // component, schema, accessibility, performance, theme
	Level       ValidationLevel        `json:"level"`
	Enabled     bool                  `json:"enabled"`
	Config      map[string]interface{} `json:"config,omitempty"`
	Validator   Validator             `json:"-"`
}

// Validator interface for custom validation logic
type Validator interface {
	Validate(ctx *ValidationContext, value interface{}) *ValidationResult
	GetRule() ValidationRule
}

// ValidationEngine is the core validation engine
type ValidationEngine struct {
	rules       map[string]ValidationRule
	validators  map[string]Validator
	options     ValidationOptions
	metrics     *MetricsCollector
	a11y        *AccessibilityValidator
	theme       *ThemeValidator
	performance *PerformanceValidator
}

// NewValidationEngine creates a new validation engine
func NewValidationEngine(opts ValidationOptions) *ValidationEngine {
	engine := &ValidationEngine{
		rules:      make(map[string]ValidationRule),
		validators: make(map[string]Validator),
		options:    opts,
	}

	if opts.EnableMetrics {
		engine.metrics = NewMetricsCollector()
	}
	if opts.EnableA11y {
		engine.a11y = NewAccessibilityValidator()
	}
	if opts.EnableTheme {
		engine.theme = NewThemeValidator()
	}
	if opts.EnablePerformance {
		engine.performance = NewPerformanceValidator()
	}

	// Register built-in rules
	engine.registerBuiltinRules()

	return engine
}

// NewValidationContext creates a validation context
func NewValidationContext(ctx context.Context, level ValidationLevel, source string) *ValidationContext {
	return &ValidationContext{
		ctx:      ctx,
		Level:    level,
		Source:   source,
		Metadata: make(map[string]interface{}),
		Options:  ValidationOptions{},
	}
}

// AddRule registers a validation rule
func (e *ValidationEngine) AddRule(rule ValidationRule, validator Validator) {
	e.rules[rule.ID] = rule
	e.validators[rule.ID] = validator
}

// Validate performs validation on the given value
func (e *ValidationEngine) Validate(ctx *ValidationContext, value interface{}) *ValidationResult {
	start := time.Now()
	result := &ValidationResult{
		Valid:     true,
		Level:     ctx.Level,
		Timestamp: start,
		Metadata:  make(map[string]interface{}),
	}

	// Run applicable validation rules
	for ruleID, rule := range e.rules {
		if !rule.Enabled {
			continue
		}

		// Skip excluded rules
		if contains(e.options.ExcludeRules, ruleID) {
			continue
		}

		// Check custom rules filter
		if len(e.options.CustomRules) > 0 && !contains(e.options.CustomRules, ruleID) {
			continue
		}

		validator := e.validators[ruleID]
		if validator == nil {
			continue
		}

		ruleResult := validator.Validate(ctx, value)
		if ruleResult != nil {
			result = e.mergeResults(result, ruleResult)
		}
	}

	// Run specialized validators if enabled
	if e.options.EnableA11y && e.a11y != nil {
		a11yResult := e.a11y.ValidateAccessibility(ctx, value)
		result.Accessibility = a11yResult
		if !a11yResult.Compliant {
			result.Valid = false
		}
	}

	if e.options.EnableTheme && e.theme != nil {
		themeResult := e.theme.ValidateTheme(ctx, value)
		result.Theme = themeResult
		if !themeResult.Valid {
			result.Valid = false
		}
	}

	if e.options.EnablePerformance && e.performance != nil {
		perfResult := e.performance.ValidatePerformance(ctx, value)
		result.Performance = perfResult
		if !perfResult.MeetsThresholds {
			result.Valid = false
		}
	}

	return result
}

// mergeResults combines validation results
func (e *ValidationEngine) mergeResults(base, addition *ValidationResult) *ValidationResult {
	if !addition.Valid {
		base.Valid = false
	}

	base.Errors = append(base.Errors, addition.Errors...)
	base.Warnings = append(base.Warnings, addition.Warnings...)

	// Merge metadata
	for k, v := range addition.Metadata {
		base.Metadata[k] = v
	}

	return base
}

// ValidateComponent validates a UI component
func (e *ValidationEngine) ValidateComponent(component interface{}, componentType string) *ValidationResult {
	ctx := NewValidationContext(context.Background(), e.getDefaultLevel(), "component")
	ctx.Component = componentType
	return e.Validate(ctx, component)
}

// ValidateSchema validates a schema definition
func (e *ValidationEngine) ValidateSchema(schema interface{}) *ValidationResult {
	ctx := NewValidationContext(context.Background(), e.getDefaultLevel(), "schema")
	return e.Validate(ctx, schema)
}

// ValidateRuntime validates runtime data
func (e *ValidationEngine) ValidateRuntime(data interface{}) *ValidationResult {
	ctx := NewValidationContext(context.Background(), e.getDefaultLevel(), "runtime")
	return e.Validate(ctx, data)
}

// GetValidationReport generates a comprehensive validation report
func (e *ValidationEngine) GetValidationReport() *ValidationReport {
	return &ValidationReport{
		Engine:    e,
		Timestamp: time.Now(),
		Rules:     e.listRules(),
		Stats:     e.getStats(),
	}
}

// ValidationReport contains comprehensive validation information
type ValidationReport struct {
	Engine    *ValidationEngine    `json:"-"`
	Timestamp time.Time           `json:"timestamp"`
	Rules     []ValidationRule    `json:"rules"`
	Stats     ValidationStats     `json:"stats"`
}

// ValidationStats contains validation statistics
type ValidationStats struct {
	TotalRules    int                    `json:"totalRules"`
	EnabledRules  int                    `json:"enabledRules"`
	Categories    map[string]int         `json:"categories"`
	Levels        map[ValidationLevel]int `json:"levels"`
}

// Helper methods

func (e *ValidationEngine) getDefaultLevel() ValidationLevel {
	if e.options.StrictMode {
		return ValidationLevelStrict
	}
	return ValidationLevelError
}

func (e *ValidationEngine) listRules() []ValidationRule {
	rules := make([]ValidationRule, 0, len(e.rules))
	for _, rule := range e.rules {
		rules = append(rules, rule)
	}
	return rules
}

func (e *ValidationEngine) getStats() ValidationStats {
	stats := ValidationStats{
		TotalRules:   len(e.rules),
		EnabledRules: 0,
		Categories:   make(map[string]int),
		Levels:       make(map[ValidationLevel]int),
	}

	for _, rule := range e.rules {
		if rule.Enabled {
			stats.EnabledRules++
		}
		stats.Categories[rule.Category]++
		stats.Levels[rule.Level]++
	}

	return stats
}

func (e *ValidationEngine) registerBuiltinRules() {
	// Component validation rules
	e.AddRule(ValidationRule{
		ID:          "component.props.required",
		Name:        "Required Props",
		Description: "Validates that required props are provided",
		Category:    "component",
		Level:       ValidationLevelError,
		Enabled:     true,
	}, &RequiredPropsValidator{})

	e.AddRule(ValidationRule{
		ID:          "component.props.types",
		Name:        "Prop Types",
		Description: "Validates that props have correct types",
		Category:    "component",
		Level:       ValidationLevelError,
		Enabled:     true,
	}, &PropTypesValidator{})

	// Schema validation rules
	e.AddRule(ValidationRule{
		ID:          "schema.structure",
		Name:        "Schema Structure",
		Description: "Validates schema structure and format",
		Category:    "schema",
		Level:       ValidationLevelError,
		Enabled:     true,
	}, &SchemaStructureValidator{})

	// More built-in rules will be added in specific validator files
}

// Utility functions

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Error implements the error interface for ValidationError
func (e ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("[%s] %s: %s", e.Code, e.Field, e.Message)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewValidationError creates a new validation error
func NewValidationError(code, message, field string, level ValidationLevel) ValidationError {
	return ValidationError{
		Code:      code,
		Message:   message,
		Field:     field,
		Level:     level,
		Timestamp: time.Now(),
	}
}

// NewValidationWarning creates a new validation warning
func NewValidationWarning(code, message, field string) ValidationWarning {
	return ValidationWarning{
		Code:    code,
		Message: message,
		Field:   field,
	}
}

// IsValid returns true if the validation result is valid
func (r *ValidationResult) IsValid() bool {
	return r.Valid && len(r.Errors) == 0
}

// HasErrors returns true if there are validation errors
func (r *ValidationResult) HasErrors() bool {
	return len(r.Errors) > 0
}

// HasWarnings returns true if there are validation warnings
func (r *ValidationResult) HasWarnings() bool {
	return len(r.Warnings) > 0
}

// GetErrorsByCode returns errors filtered by code
func (r *ValidationResult) GetErrorsByCode(code string) []ValidationError {
	var errors []ValidationError
	for _, err := range r.Errors {
		if err.Code == code {
			errors = append(errors, err)
		}
	}
	return errors
}

// GetErrorsByField returns errors filtered by field
func (r *ValidationResult) GetErrorsByField(field string) []ValidationError {
	var errors []ValidationError
	for _, err := range r.Errors {
		if err.Field == field {
			errors = append(errors, err)
		}
	}
	return errors
}

// AddMetadata adds metadata to the validation context
func (ctx *ValidationContext) AddMetadata(key string, value interface{}) {
	if ctx.Metadata == nil {
		ctx.Metadata = make(map[string]interface{})
	}
	ctx.Metadata[key] = value
}

// GetMetadata retrieves metadata from the validation context
func (ctx *ValidationContext) GetMetadata(key string) interface{} {
	if ctx.Metadata == nil {
		return nil
	}
	return ctx.Metadata[key]
}

// WithComponent creates a new context with a component set
func (ctx *ValidationContext) WithComponent(component string) *ValidationContext {
	newCtx := *ctx
	newCtx.Component = component
	return &newCtx
}

// WithLevel creates a new context with a validation level set
func (ctx *ValidationContext) WithLevel(level ValidationLevel) *ValidationContext {
	newCtx := *ctx
	newCtx.Level = level
	return &newCtx
}