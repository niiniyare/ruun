package validation

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// SchemaValidator validates form schemas, themes, and configurations
type SchemaValidator struct {
	rules      []SchemaValidationRule
	engine     *ValidationEngine
	reporter   *ErrorReporter
}

// SchemaValidationRule defines a rule for schema validation
type SchemaValidationRule struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Category    SchemaCategory         `json:"category"`
	Level       ValidationLevel        `json:"level"`
	Enabled     bool                  `json:"enabled"`
	Validator   SchemaRuleValidator   `json:"-"`
}

// SchemaCategory represents different types of schemas that can be validated
type SchemaCategory string

const (
	SchemaCategoryForm   SchemaCategory = "form"
	SchemaCategoryField  SchemaCategory = "field"
	SchemaCategoryTheme  SchemaCategory = "theme"
	SchemaCategoryLayout SchemaCategory = "layout"
	SchemaCategoryConfig SchemaCategory = "config"
)

// SchemaRuleValidator interface for schema validation rules
type SchemaRuleValidator interface {
	ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult
	GetCategory() SchemaCategory
}

// SchemaValidationContext provides context for schema validation
type SchemaValidationContext struct {
	Path     []string               `json:"path"`
	Schema   interface{}            `json:"schema"`
	Parent   interface{}            `json:"parent,omitempty"`
	Root     interface{}            `json:"root"`
	Metadata map[string]interface{} `json:"metadata"`
	Level    ValidationLevel        `json:"level"`
}

// SchemaValidationResult represents the result of schema validation
type SchemaValidationResult struct {
	Valid       bool                   `json:"valid"`
	Errors      []SchemaError          `json:"errors,omitempty"`
	Warnings    []SchemaWarning        `json:"warnings,omitempty"`
	Suggestions []SchemaSuggestion     `json:"suggestions,omitempty"`
	Metadata    map[string]interface{} `json:"metadata"`
	Performance *ValidationMetrics     `json:"performance,omitempty"`
	Timestamp   time.Time             `json:"timestamp"`
}

// SchemaError represents a schema validation error
type SchemaError struct {
	Code        string         `json:"code"`
	Message     string         `json:"message"`
	Path        string         `json:"path"`
	Value       interface{}    `json:"value,omitempty"`
	Expected    interface{}    `json:"expected,omitempty"`
	Actual      interface{}    `json:"actual,omitempty"`
	Level       ValidationLevel `json:"level"`
	Category    SchemaCategory `json:"category"`
	Rule        string         `json:"rule"`
	Location    *SourceLocation `json:"location,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"`
	Suggestion  string         `json:"suggestion,omitempty"`
}

// SchemaWarning represents a schema validation warning
type SchemaWarning struct {
	Code       string                 `json:"code"`
	Message    string                 `json:"message"`
	Path       string                 `json:"path"`
	Value      interface{}            `json:"value,omitempty"`
	Category   SchemaCategory         `json:"category"`
	Rule       string                 `json:"rule"`
	Location   *SourceLocation        `json:"location,omitempty"`
	Context    map[string]interface{} `json:"context,omitempty"`
	Suggestion string                 `json:"suggestion,omitempty"`
}

// SchemaSuggestion represents a suggestion for schema improvement
type SchemaSuggestion struct {
	Code        string                 `json:"code"`
	Message     string                 `json:"message"`
	Path        string                 `json:"path"`
	Suggestion  string                 `json:"suggestion"`
	AutoFix     bool                  `json:"autoFix"`
	Category    SchemaCategory         `json:"category"`
	Rule        string                 `json:"rule"`
	Details     map[string]interface{} `json:"details,omitempty"`
}

// ValidationMetrics contains performance metrics for validation
type ValidationMetrics struct {
	Duration    time.Duration `json:"duration"`
	RulesRun    int          `json:"rulesRun"`
	ErrorsFound int          `json:"errorsFound"`
	WarningsFound int        `json:"warningsFound"`
	Memory      int64        `json:"memory,omitempty"`
}

// ErrorReporter provides detailed error reporting with suggestions
type ErrorReporter struct {
	format     ReportFormat
	suggestions bool
	context    bool
}

// ReportFormat defines how errors are formatted
type ReportFormat string

const (
	ReportFormatJSON     ReportFormat = "json"
	ReportFormatHuman    ReportFormat = "human"
	ReportFormatMarkdown ReportFormat = "markdown"
	ReportFormatJUnit    ReportFormat = "junit"
)

// NewSchemaValidator creates a new schema validator
func NewSchemaValidator() *SchemaValidator {
	validator := &SchemaValidator{
		rules:    make([]SchemaValidationRule, 0),
		reporter: NewErrorReporter(ReportFormatHuman, true, true),
	}
	
	// Register built-in validation rules
	validator.registerBuiltinRules()
	
	return validator
}

// NewErrorReporter creates a new error reporter
func NewErrorReporter(format ReportFormat, suggestions, context bool) *ErrorReporter {
	return &ErrorReporter{
		format:      format,
		suggestions: suggestions,
		context:     context,
	}
}

// ValidateSchema validates a schema with detailed error reporting
func (sv *SchemaValidator) ValidateSchema(schemaData interface{}) *SchemaValidationResult {
	start := time.Now()
	
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: start,
		Metadata:  make(map[string]interface{}),
	}

	context := &SchemaValidationContext{
		Path:     []string{},
		Schema:   schemaData,
		Root:     schemaData,
		Metadata: make(map[string]interface{}),
		Level:    ValidationLevelError,
	}

	rulesRun := 0
	
	// Run all enabled validation rules
	for _, rule := range sv.rules {
		if !rule.Enabled {
			continue
		}

		ruleResult := rule.Validator.ValidateSchema(schemaData, context)
		if ruleResult != nil {
			rulesRun++
			result = sv.mergeResults(result, ruleResult)
		}
	}

	// Add performance metrics
	result.Performance = &ValidationMetrics{
		Duration:      time.Since(start),
		RulesRun:      rulesRun,
		ErrorsFound:   len(result.Errors),
		WarningsFound: len(result.Warnings),
	}

	return result
}

// ValidateFormSchema validates a form schema
func (sv *SchemaValidator) ValidateFormSchema(formSchema *schema.Schema) *SchemaValidationResult {
	return sv.ValidateSchema(formSchema)
}

// ValidateThemeSchema validates a theme schema
func (sv *SchemaValidator) ValidateThemeSchema(themeSchema *schema.Theme) *SchemaValidationResult {
	return sv.ValidateSchema(themeSchema)
}

// ValidateFieldSchema validates a field schema
func (sv *SchemaValidator) ValidateFieldSchema(fieldSchema *schema.Field) *SchemaValidationResult {
	return sv.ValidateSchema(fieldSchema)
}

// AddRule adds a validation rule
func (sv *SchemaValidator) AddRule(rule SchemaValidationRule) {
	sv.rules = append(sv.rules, rule)
}

// DisableRule disables a validation rule
func (sv *SchemaValidator) DisableRule(ruleID string) {
	for i := range sv.rules {
		if sv.rules[i].ID == ruleID {
			sv.rules[i].Enabled = false
			break
		}
	}
}

// EnableRule enables a validation rule
func (sv *SchemaValidator) EnableRule(ruleID string) {
	for i := range sv.rules {
		if sv.rules[i].ID == ruleID {
			sv.rules[i].Enabled = true
			break
		}
	}
}

// GenerateReport generates a formatted validation report
func (sv *SchemaValidator) GenerateReport(result *SchemaValidationResult, format ReportFormat) (string, error) {
	reporter := NewErrorReporter(format, true, true)
	return reporter.GenerateReport(result)
}

// registerBuiltinRules registers built-in validation rules
func (sv *SchemaValidator) registerBuiltinRules() {
	// Form schema validation rules
	sv.AddRule(SchemaValidationRule{
		ID:          "form.schema.structure",
		Name:        "Form Schema Structure",
		Description: "Validates basic form schema structure",
		Category:    SchemaCategoryForm,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &FormStructureValidator{},
	})

	sv.AddRule(SchemaValidationRule{
		ID:          "form.schema.fields",
		Name:        "Form Fields Validation",
		Description: "Validates form fields array and structure",
		Category:    SchemaCategoryForm,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &FormFieldsValidator{},
	})

	sv.AddRule(SchemaValidationRule{
		ID:          "form.schema.required",
		Name:        "Required Fields",
		Description: "Validates required fields configuration",
		Category:    SchemaCategoryForm,
		Level:       ValidationLevelWarning,
		Enabled:     true,
		Validator:   &RequiredFieldsValidator{},
	})

	// Field validation rules
	sv.AddRule(SchemaValidationRule{
		ID:          "field.schema.type",
		Name:        "Field Type Validation",
		Description: "Validates field type and configuration",
		Category:    SchemaCategoryField,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &FieldTypeValidator{},
	})

	sv.AddRule(SchemaValidationRule{
		ID:          "field.schema.validation",
		Name:        "Field Validation Rules",
		Description: "Validates field validation configuration",
		Category:    SchemaCategoryField,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &FieldValidationValidator{},
	})

	// Theme validation rules
	sv.AddRule(SchemaValidationRule{
		ID:          "theme.schema.tokens",
		Name:        "Theme Token Validation",
		Description: "Validates theme design tokens",
		Category:    SchemaCategoryTheme,
		Level:       ValidationLevelError,
		Enabled:     true,
		Validator:   &ThemeTokenValidator{},
	})

	sv.AddRule(SchemaValidationRule{
		ID:          "theme.schema.consistency",
		Name:        "Theme Consistency",
		Description: "Validates theme consistency and completeness",
		Category:    SchemaCategoryTheme,
		Level:       ValidationLevelWarning,
		Enabled:     true,
		Validator:   &ThemeConsistencyValidator{},
	})
}

// mergeResults merges schema validation results
func (sv *SchemaValidator) mergeResults(base, addition *SchemaValidationResult) *SchemaValidationResult {
	if !addition.Valid {
		base.Valid = false
	}

	base.Errors = append(base.Errors, addition.Errors...)
	base.Warnings = append(base.Warnings, addition.Warnings...)
	base.Suggestions = append(base.Suggestions, addition.Suggestions...)

	// Merge metadata
	for k, v := range addition.Metadata {
		base.Metadata[k] = v
	}

	return base
}

// Built-in Schema Validators

// FormStructureValidator validates basic form schema structure
type FormStructureValidator struct{}

func (v *FormStructureValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	formSchema, ok := schema.(*schema.Schema)
	if !ok {
		// Try to convert from interface
		if schemaMap, ok := schema.(map[string]interface{}); ok {
			// Basic structure validation for map representation
			requiredFields := []string{"id", "type", "title"}
			for _, field := range requiredFields {
				if _, exists := schemaMap[field]; !exists {
					result.Errors = append(result.Errors, SchemaError{
						Code:     "form.structure.missing_field",
						Message:  fmt.Sprintf("Required field '%s' is missing", field),
						Path:     strings.Join(context.Path, "."),
						Expected: field,
						Level:    ValidationLevelError,
						Category: SchemaCategoryForm,
						Rule:     "form.schema.structure",
					})
					result.Valid = false
				}
			}
		} else {
			result.Errors = append(result.Errors, SchemaError{
				Code:     "form.structure.invalid_type",
				Message:  "Schema must be a valid form schema object",
				Path:     strings.Join(context.Path, "."),
				Expected: "*schema.Schema or map[string]interface{}",
				Actual:   reflect.TypeOf(schema).String(),
				Level:    ValidationLevelError,
				Category: SchemaCategoryForm,
				Rule:     "form.schema.structure",
			})
			result.Valid = false
		}
		return result
	}

	// Validate required fields
	if formSchema.ID == "" {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "form.structure.missing_id",
			Message:    "Form schema must have a valid ID",
			Path:       strings.Join(append(context.Path, "id"), "."),
			Level:      ValidationLevelError,
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.structure",
			Suggestion: "Add a unique ID to the form schema",
		})
		result.Valid = false
	}

	if formSchema.Type == "" {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "form.structure.missing_type",
			Message:    "Form schema must have a valid type",
			Path:       strings.Join(append(context.Path, "type"), "."),
			Level:      ValidationLevelError,
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.structure",
			Suggestion: "Add a type (e.g., 'form', 'object') to the schema",
		})
		result.Valid = false
	}

	if formSchema.Title == "" {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "form.structure.missing_title",
			Message:    "Form schema should have a title for better user experience",
			Path:       strings.Join(append(context.Path, "title"), "."),
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.structure",
			Suggestion: "Add a descriptive title to the form schema",
		})
	}

	return result
}

func (v *FormStructureValidator) GetCategory() SchemaCategory {
	return SchemaCategoryForm
}

// FormFieldsValidator validates form fields
type FormFieldsValidator struct{}

func (v *FormFieldsValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	formSchema, ok := schema.(*schema.Schema)
	if !ok {
		return result // Skip if not a form schema
	}

	// Validate fields array
	if len(formSchema.Fields) == 0 {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "form.fields.empty",
			Message:    "Form schema has no fields defined",
			Path:       strings.Join(append(context.Path, "fields"), "."),
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.fields",
			Suggestion: "Add field definitions to make the form functional",
		})
	}

	// Validate each field
	for i, field := range formSchema.Fields {
		fieldContext := &SchemaValidationContext{
			Path:     append(context.Path, "fields", fmt.Sprintf("[%d]", i)),
			Schema:   field,
			Parent:   formSchema,
			Root:     context.Root,
			Metadata: context.Metadata,
			Level:    context.Level,
		}

		fieldResult := v.validateField(field, fieldContext)
		result = v.mergeResults(result, fieldResult)
	}

	return result
}

func (v *FormFieldsValidator) validateField(field *schema.Field, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	// Validate required field properties
	if field.Name == "" {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "field.missing_name",
			Message:    "Field must have a name",
			Path:       strings.Join(append(context.Path, "name"), "."),
			Level:      ValidationLevelError,
			Category:   SchemaCategoryField,
			Rule:       "form.schema.fields",
			Suggestion: "Add a unique name to the field",
		})
		result.Valid = false
	}

	if field.Type == "" {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "field.missing_type",
			Message:    "Field must have a type",
			Path:       strings.Join(append(context.Path, "type"), "."),
			Level:      ValidationLevelError,
			Category:   SchemaCategoryField,
			Rule:       "form.schema.fields",
			Suggestion: "Add a valid field type (text, email, number, etc.)",
		})
		result.Valid = false
	}

	return result
}

func (v *FormFieldsValidator) mergeResults(base, addition *SchemaValidationResult) *SchemaValidationResult {
	if !addition.Valid {
		base.Valid = false
	}
	base.Errors = append(base.Errors, addition.Errors...)
	base.Warnings = append(base.Warnings, addition.Warnings...)
	base.Suggestions = append(base.Suggestions, addition.Suggestions...)
	return base
}

func (v *FormFieldsValidator) GetCategory() SchemaCategory {
	return SchemaCategoryForm
}

// RequiredFieldsValidator validates required field configuration
type RequiredFieldsValidator struct{}

func (v *RequiredFieldsValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	formSchema, ok := schema.(*schema.Schema)
	if !ok {
		return result
	}

	requiredCount := 0
	for _, field := range formSchema.Fields {
		if field.Required {
			requiredCount++
		}
	}

	// Warn if no required fields
	if requiredCount == 0 && len(formSchema.Fields) > 0 {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "form.required.none",
			Message:    "Form has no required fields which may impact data quality",
			Path:       strings.Join(context.Path, "."),
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.required",
			Suggestion: "Consider making some important fields required",
		})
	}

	// Warn if all fields are required (might hurt UX)
	if requiredCount == len(formSchema.Fields) && len(formSchema.Fields) > 3 {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "form.required.all",
			Message:    "All fields are required which may negatively impact user experience",
			Path:       strings.Join(context.Path, "."),
			Category:   SchemaCategoryForm,
			Rule:       "form.schema.required",
			Suggestion: "Consider making some fields optional to improve user experience",
		})
	}

	return result
}

func (v *RequiredFieldsValidator) GetCategory() SchemaCategory {
	return SchemaCategoryForm
}

// FieldTypeValidator validates field types
type FieldTypeValidator struct{}

func (v *FieldTypeValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	field, ok := schema.(*schema.Field)
	if !ok {
		return result
	}

	// Validate field type
	validTypes := []schema.FieldType{
		schema.FieldText, schema.FieldEmail, schema.FieldPassword,
		schema.FieldNumber, schema.FieldDate, schema.FieldSelect,
		schema.FieldCheckbox, schema.FieldTextarea, schema.FieldFile,
	}

	typeValid := false
	for _, validType := range validTypes {
		if field.Type == validType {
			typeValid = true
			break
		}
	}

	if !typeValid {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "field.type.invalid",
			Message:    fmt.Sprintf("Invalid field type: %s", field.Type),
			Path:       strings.Join(append(context.Path, "type"), "."),
			Actual:     string(field.Type),
			Expected:   validTypes,
			Level:      ValidationLevelError,
			Category:   SchemaCategoryField,
			Rule:       "field.schema.type",
			Suggestion: "Use a valid field type from the supported list",
		})
		result.Valid = false
	}

	return result
}

func (v *FieldTypeValidator) GetCategory() SchemaCategory {
	return SchemaCategoryField
}

// FieldValidationValidator validates field validation rules
type FieldValidationValidator struct{}

func (v *FieldValidationValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	field, ok := schema.(*schema.Field)
	if !ok {
		return result
	}

	if field.Validation == nil {
		return result
	}

	validation := field.Validation

	// Validate min/max length for string types
	if field.Type == schema.FieldText || field.Type == schema.FieldTextarea {
		if validation.MinLength != nil && validation.MaxLength != nil {
			if *validation.MinLength > *validation.MaxLength {
				result.Errors = append(result.Errors, SchemaError{
					Code:       "field.validation.length_conflict",
					Message:    "MinLength cannot be greater than MaxLength",
					Path:       strings.Join(append(context.Path, "validation"), "."),
					Level:      ValidationLevelError,
					Category:   SchemaCategoryField,
					Rule:       "field.schema.validation",
					Suggestion: "Ensure minLength is less than or equal to maxLength",
				})
				result.Valid = false
			}
		}
	}

	// Validate min/max for number types
	if field.Type == schema.FieldNumber {
		if validation.Min != nil && validation.Max != nil {
			if *validation.Min > *validation.Max {
				result.Errors = append(result.Errors, SchemaError{
					Code:       "field.validation.range_conflict",
					Message:    "Min value cannot be greater than Max value",
					Path:       strings.Join(append(context.Path, "validation"), "."),
					Level:      ValidationLevelError,
					Category:   SchemaCategoryField,
					Rule:       "field.schema.validation",
					Suggestion: "Ensure min is less than or equal to max",
				})
				result.Valid = false
			}
		}
	}

	return result
}

func (v *FieldValidationValidator) GetCategory() SchemaCategory {
	return SchemaCategoryField
}

// ThemeTokenValidator validates theme design tokens
type ThemeTokenValidator struct{}

func (v *ThemeTokenValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	theme, ok := schema.(*schema.Theme)
	if !ok {
		return result
	}

	if theme.Tokens == nil {
		result.Errors = append(result.Errors, SchemaError{
			Code:       "theme.tokens.missing",
			Message:    "Theme must have design tokens defined",
			Path:       strings.Join(append(context.Path, "tokens"), "."),
			Level:      ValidationLevelError,
			Category:   SchemaCategoryTheme,
			Rule:       "theme.schema.tokens",
			Suggestion: "Define design tokens for colors, typography, spacing, etc.",
		})
		result.Valid = false
		return result
	}

	// Validate that essential token categories are present
	if theme.Tokens.Semantic == nil {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "theme.tokens.semantic_missing",
			Message:    "Theme should have semantic tokens for better maintainability",
			Path:       strings.Join(append(context.Path, "tokens", "semantic"), "."),
			Category:   SchemaCategoryTheme,
			Rule:       "theme.schema.tokens",
			Suggestion: "Add semantic color tokens (primary, secondary, background, etc.)",
		})
	}

	return result
}

func (v *ThemeTokenValidator) GetCategory() SchemaCategory {
	return SchemaCategoryTheme
}

// ThemeConsistencyValidator validates theme consistency
type ThemeConsistencyValidator struct{}

func (v *ThemeConsistencyValidator) ValidateSchema(schema interface{}, context *SchemaValidationContext) *SchemaValidationResult {
	result := &SchemaValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	theme, ok := schema.(*schema.Theme)
	if !ok {
		return result
	}

	// Check for theme completeness
	if theme.Name == "" {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "theme.consistency.missing_name",
			Message:    "Theme should have a descriptive name",
			Path:       strings.Join(append(context.Path, "name"), "."),
			Category:   SchemaCategoryTheme,
			Rule:       "theme.schema.consistency",
			Suggestion: "Add a descriptive name to identify the theme",
		})
	}

	if theme.Version == "" {
		result.Warnings = append(result.Warnings, SchemaWarning{
			Code:       "theme.consistency.missing_version",
			Message:    "Theme should have a version for tracking changes",
			Path:       strings.Join(append(context.Path, "version"), "."),
			Category:   SchemaCategoryTheme,
			Rule:       "theme.schema.consistency",
			Suggestion: "Add semantic version number (e.g., '1.0.0') to the theme",
		})
	}

	return result
}

func (v *ThemeConsistencyValidator) GetCategory() SchemaCategory {
	return SchemaCategoryTheme
}

// ErrorReporter methods

// GenerateReport generates a formatted error report
func (er *ErrorReporter) GenerateReport(result *SchemaValidationResult) (string, error) {
	switch er.format {
	case ReportFormatJSON:
		return er.generateJSONReport(result)
	case ReportFormatHuman:
		return er.generateHumanReport(result)
	case ReportFormatMarkdown:
		return er.generateMarkdownReport(result)
	case ReportFormatJUnit:
		return er.generateJUnitReport(result)
	default:
		return er.generateHumanReport(result)
	}
}

func (er *ErrorReporter) generateJSONReport(result *SchemaValidationResult) (string, error) {
	data, err := json.MarshalIndent(result, "", "  ")
	return string(data), err
}

func (er *ErrorReporter) generateHumanReport(result *SchemaValidationResult) (string, error) {
	var report strings.Builder
	
	report.WriteString("Schema Validation Report\n")
	report.WriteString("========================\n\n")
	
	if result.Valid {
		report.WriteString("✅ Schema is valid\n\n")
	} else {
		report.WriteString("❌ Schema validation failed\n\n")
	}
	
	if len(result.Errors) > 0 {
		report.WriteString("Errors:\n")
		report.WriteString("-------\n")
		for _, err := range result.Errors {
			report.WriteString(fmt.Sprintf("• %s: %s\n", err.Path, err.Message))
			if er.suggestions && err.Suggestion != "" {
				report.WriteString(fmt.Sprintf("  💡 %s\n", err.Suggestion))
			}
			report.WriteString("\n")
		}
	}
	
	if len(result.Warnings) > 0 {
		report.WriteString("Warnings:\n")
		report.WriteString("---------\n")
		for _, warn := range result.Warnings {
			report.WriteString(fmt.Sprintf("• %s: %s\n", warn.Path, warn.Message))
			if er.suggestions && warn.Suggestion != "" {
				report.WriteString(fmt.Sprintf("  💡 %s\n", warn.Suggestion))
			}
			report.WriteString("\n")
		}
	}
	
	if result.Performance != nil {
		report.WriteString("Performance:\n")
		report.WriteString("------------\n")
		report.WriteString(fmt.Sprintf("Duration: %v\n", result.Performance.Duration))
		report.WriteString(fmt.Sprintf("Rules Run: %d\n", result.Performance.RulesRun))
		report.WriteString(fmt.Sprintf("Errors: %d, Warnings: %d\n\n", result.Performance.ErrorsFound, result.Performance.WarningsFound))
	}
	
	return report.String(), nil
}

func (er *ErrorReporter) generateMarkdownReport(result *SchemaValidationResult) (string, error) {
	var report strings.Builder
	
	report.WriteString("# Schema Validation Report\n\n")
	
	if result.Valid {
		report.WriteString("✅ **Schema is valid**\n\n")
	} else {
		report.WriteString("❌ **Schema validation failed**\n\n")
	}
	
	if len(result.Errors) > 0 {
		report.WriteString("## Errors\n\n")
		for _, err := range result.Errors {
			report.WriteString(fmt.Sprintf("- **%s**: %s\n", err.Path, err.Message))
			if er.suggestions && err.Suggestion != "" {
				report.WriteString(fmt.Sprintf("  > 💡 %s\n", err.Suggestion))
			}
			report.WriteString("\n")
		}
	}
	
	if len(result.Warnings) > 0 {
		report.WriteString("## Warnings\n\n")
		for _, warn := range result.Warnings {
			report.WriteString(fmt.Sprintf("- **%s**: %s\n", warn.Path, warn.Message))
			if er.suggestions && warn.Suggestion != "" {
				report.WriteString(fmt.Sprintf("  > 💡 %s\n", warn.Suggestion))
			}
			report.WriteString("\n")
		}
	}
	
	return report.String(), nil
}

func (er *ErrorReporter) generateJUnitReport(result *SchemaValidationResult) (string, error) {
	// Simplified JUnit XML format
	var report strings.Builder
	
	report.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	report.WriteString("\n")
	report.WriteString(fmt.Sprintf(`<testsuite name="SchemaValidation" tests="1" failures="%d" errors="0" time="%.3f">`,
		len(result.Errors), result.Performance.Duration.Seconds()))
	report.WriteString("\n")
	
	if result.Valid {
		report.WriteString(`  <testcase name="SchemaValidation" classname="validation"/>`)
	} else {
		report.WriteString(`  <testcase name="SchemaValidation" classname="validation">`)
		report.WriteString("\n")
		for _, err := range result.Errors {
			report.WriteString(fmt.Sprintf(`    <failure message="%s" type="%s">%s</failure>`,
				err.Message, err.Code, err.Path))
			report.WriteString("\n")
		}
		report.WriteString(`  </testcase>`)
	}
	
	report.WriteString("\n")
	report.WriteString("</testsuite>")
	
	return report.String(), nil
}

// SchemaStructureValidator validates basic schema structure
type SchemaStructureValidator struct{}

func (v *SchemaStructureValidator) Validate(ctx *ValidationContext, value interface{}) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]interface{}),
	}

	// This is a simplified implementation that would delegate to SchemaValidator
	schemaValidator := NewSchemaValidator()
	schemaResult := schemaValidator.ValidateSchema(value)
	
	// Convert SchemaValidationResult to ValidationResult
	result.Valid = schemaResult.Valid
	
	for _, err := range schemaResult.Errors {
		result.Errors = append(result.Errors, ValidationError{
			Code:      err.Code,
			Message:   err.Message,
			Field:     err.Path,
			Level:     err.Level,
			Source:    "schema",
			Timestamp: time.Now(),
		})
	}
	
	for _, warn := range schemaResult.Warnings {
		result.Warnings = append(result.Warnings, ValidationWarning{
			Code:    warn.Code,
			Message: warn.Message,
			Field:   warn.Path,
			Source:  "schema",
		})
	}

	return result
}

func (v *SchemaStructureValidator) GetRule() ValidationRule {
	return ValidationRule{
		ID:          "schema.structure",
		Name:        "Schema Structure",
		Description: "Validates schema structure and format",
		Category:    "schema",
		Level:       ValidationLevelError,
		Enabled:     true,
	}
}