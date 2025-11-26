package schema

import (
	"fmt"
	"regexp"
)

// UnifiedValidation combines field-level and cross-field validation in a single structure
type UnifiedValidation struct {
	// Field-level validation (from FieldValidation)
	// String validation
	MinLength *int   `json:"minLength,omitempty"`
	MaxLength *int   `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Format    string `json:"format,omitempty"`

	// Number validation
	Min      *float64 `json:"min,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Step     *float64 `json:"step,omitempty"`
	Integer  bool     `json:"integer,omitempty"`
	Positive bool     `json:"positive,omitempty"`

	// Array validation
	MinItems *int `json:"minItems,omitempty"`
	MaxItems *int `json:"maxItems,omitempty"`
	Unique   bool `json:"unique,omitempty"`

	// File validation
	MaxFileSize  int64    `json:"maxFileSize,omitempty"`
	AcceptedTypes []string `json:"acceptedTypes,omitempty"`

	// Cross-field validation (from Validation)
	Rules      []ValidationRule `json:"rules,omitempty" validate:"dive"`
	Mode       ValidationMode   `json:"mode,omitempty"`
	Async      bool             `json:"async,omitempty"`
	DebounceMS int              `json:"debounceMs,omitempty"`

	// Unified features
	Required       bool                `json:"required,omitempty"`
	Messages       *ValidationMessages `json:"messages,omitempty"`
	CustomFunction string              `json:"customFunction,omitempty"`
	Custom         string              `json:"custom,omitempty"`       // Backward compatibility alias for CustomFunction
	Condition      string              `json:"condition,omitempty"`    // When to apply validation
}

// ValidationMode defines when validation should occur
type ValidationMode string

const (
	ValidationModeOnChange ValidationMode = "onChange"
	ValidationModeOnBlur   ValidationMode = "onBlur"
	ValidationModeOnSubmit ValidationMode = "onSubmit"
	ValidationModeManual   ValidationMode = "manual"
)

// IsValid validates the ValidationMode enum
func (vm ValidationMode) IsValid() bool {
	switch vm {
	case ValidationModeOnChange, ValidationModeOnBlur, ValidationModeOnSubmit, ValidationModeManual:
		return true
	default:
		return false
	}
}

// HasFieldValidation returns true if field-level validation rules are set
func (v *UnifiedValidation) HasFieldValidation() bool {
	if v == nil {
		return false
	}
	return v.MinLength != nil || v.MaxLength != nil || v.Pattern != "" ||
		v.Min != nil || v.Max != nil || v.MinItems != nil || v.MaxItems != nil ||
		v.MaxFileSize > 0 || len(v.AcceptedTypes) > 0
}

// HasCrossFieldValidation returns true if cross-field validation rules are set
func (v *UnifiedValidation) HasCrossFieldValidation() bool {
	if v == nil {
		return false
	}
	return len(v.Rules) > 0
}

// HasCustomValidation returns true if custom validation is configured
func (v *UnifiedValidation) HasCustomValidation() bool {
	if v == nil {
		return false
	}
	return v.CustomFunction != "" || v.Custom != ""
}

// IsEmpty returns true if no validation rules are configured
func (v *UnifiedValidation) IsEmpty() bool {
	if v == nil {
		return true
	}
	return !v.HasFieldValidation() && !v.HasCrossFieldValidation() && !v.HasCustomValidation() && !v.Required
}

// GetEffectiveMode returns the effective validation mode with fallback
func (v *UnifiedValidation) GetEffectiveMode() ValidationMode {
	if v == nil || v.Mode == "" {
		return ValidationModeOnBlur // Default
	}
	if !v.Mode.IsValid() {
		return ValidationModeOnBlur // Fallback for invalid modes
	}
	return v.Mode
}

// ValidateStringValue validates a string value against field validation rules
func (v *UnifiedValidation) ValidateStringValue(value string) []string {
	if v == nil {
		return nil
	}

	var errors []string

	// Required validation
	if v.Required && value == "" {
		message := "This field is required"
		if v.Messages != nil && v.Messages.Required != "" {
			message = v.Messages.Required
		}
		errors = append(errors, message)
		return errors // If required and empty, don't check other rules
	}

	// Skip other validations if value is empty and not required
	if value == "" {
		return errors
	}

	// Length validation
	if v.MinLength != nil && len(value) < *v.MinLength {
		message := fmt.Sprintf("Must be at least %d characters", *v.MinLength)
		if v.Messages != nil && v.Messages.MinLength != "" {
			message = v.Messages.MinLength
		}
		errors = append(errors, message)
	}

	if v.MaxLength != nil && len(value) > *v.MaxLength {
		message := fmt.Sprintf("Must be no more than %d characters", *v.MaxLength)
		if v.Messages != nil && v.Messages.MaxLength != "" {
			message = v.Messages.MaxLength
		}
		errors = append(errors, message)
	}

	// Pattern validation
	if v.Pattern != "" {
		if matched, err := regexp.MatchString(v.Pattern, value); err != nil {
			errors = append(errors, "Invalid pattern configuration")
		} else if !matched {
			message := "Invalid format"
			if v.Messages != nil && v.Messages.Pattern != "" {
				message = v.Messages.Pattern
			}
			errors = append(errors, message)
		}
	}

	// Format validation (basic implementations)
	if v.Format != "" {
		if !v.validateFormat(value) {
			message := fmt.Sprintf("Invalid %s format", v.Format)
			errors = append(errors, message)
		}
	}

	return errors
}

// validateFormat validates common formats
func (v *UnifiedValidation) validateFormat(value string) bool {
	switch v.Format {
	case "email":
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		return emailRegex.MatchString(value)
	case "url":
		urlRegex := regexp.MustCompile(`^https?://[^\s]+$`)
		return urlRegex.MatchString(value)
	case "phone":
		phoneRegex := regexp.MustCompile(`^\+?[\d\s\-\(\)]{10,}$`)
		return phoneRegex.MatchString(value)
	default:
		return true // Unknown format, assume valid
	}
}

// ValidateNumberValue validates a numeric value against validation rules
func (v *UnifiedValidation) ValidateNumberValue(value float64) []string {
	if v == nil {
		return nil
	}

	var errors []string

	// Min/Max validation
	if v.Min != nil && value < *v.Min {
		message := fmt.Sprintf("Must be at least %g", *v.Min)
		if v.Messages != nil && v.Messages.Min != "" {
			message = v.Messages.Min
		}
		errors = append(errors, message)
	}

	if v.Max != nil && value > *v.Max {
		message := fmt.Sprintf("Must be no more than %g", *v.Max)
		if v.Messages != nil && v.Messages.Max != "" {
			message = v.Messages.Max
		}
		errors = append(errors, message)
	}

	// Integer validation
	if v.Integer && value != float64(int64(value)) {
		errors = append(errors, "Must be a whole number")
	}

	// Positive validation
	if v.Positive && value <= 0 {
		errors = append(errors, "Must be a positive number")
	}

	// Step validation
	if v.Step != nil && *v.Step > 0 {
		remainder := value
		if v.Min != nil {
			remainder = value - *v.Min
		}
		if remainder != 0 && int64(remainder/ *v.Step) != int64(remainder/ *v.Step) {
			errors = append(errors, fmt.Sprintf("Must be a multiple of %g", *v.Step))
		}
	}

	return errors
}

// ValidateArrayValue validates an array value against validation rules
func (v *UnifiedValidation) ValidateArrayValue(values []any) []string {
	if v == nil {
		return nil
	}

	var errors []string
	length := len(values)

	// MinItems validation
	if v.MinItems != nil && length < *v.MinItems {
		message := fmt.Sprintf("Must have at least %d items", *v.MinItems)
		errors = append(errors, message)
	}

	// MaxItems validation
	if v.MaxItems != nil && length > *v.MaxItems {
		message := fmt.Sprintf("Must have no more than %d items", *v.MaxItems)
		errors = append(errors, message)
	}

	// Unique validation
	if v.Unique && length > 1 {
		seen := make(map[string]bool)
		for _, value := range values {
			key := fmt.Sprintf("%v", value)
			if seen[key] {
				errors = append(errors, "All items must be unique")
				break
			}
			seen[key] = true
		}
	}

	return errors
}

// Merge combines two UnifiedValidation instances, with 'other' taking precedence
func (v *UnifiedValidation) Merge(other *UnifiedValidation) *UnifiedValidation {
	if v == nil {
		return other
	}
	if other == nil {
		return v
	}

	result := *v // Copy base

	// Merge field validation
	if other.MinLength != nil {
		result.MinLength = other.MinLength
	}
	if other.MaxLength != nil {
		result.MaxLength = other.MaxLength
	}
	if other.Pattern != "" {
		result.Pattern = other.Pattern
	}
	if other.Format != "" {
		result.Format = other.Format
	}
	if other.Min != nil {
		result.Min = other.Min
	}
	if other.Max != nil {
		result.Max = other.Max
	}
	if other.Step != nil {
		result.Step = other.Step
	}
	if other.MinItems != nil {
		result.MinItems = other.MinItems
	}
	if other.MaxItems != nil {
		result.MaxItems = other.MaxItems
	}

	// Merge cross-field validation
	if len(other.Rules) > 0 {
		result.Rules = other.Rules
	}
	if other.Mode != "" {
		result.Mode = other.Mode
	}
	if other.CustomFunction != "" {
		result.CustomFunction = other.CustomFunction
	}

	// Merge messages
	if other.Messages != nil {
		result.Messages = other.Messages
	}

	return &result
}

// GetCustom returns the custom validation function for backward compatibility
func (v *UnifiedValidation) GetCustom() string {
	if v == nil {
		return ""
	}
	// Return the first non-empty custom validation
	if v.Custom != "" {
		return v.Custom
	}
	return v.CustomFunction
}

// SetCustom sets the custom validation function for backward compatibility  
func (v *UnifiedValidation) SetCustom(custom string) {
	if v != nil {
		v.Custom = custom
		v.CustomFunction = custom // Keep both in sync
	}
}

// Backward compatibility type aliases
type FieldValidation = UnifiedValidation
type Validation = UnifiedValidation

// Legacy validation construction for backward compatibility
func NewFieldValidation() *FieldValidation {
	return &UnifiedValidation{}
}

func NewValidation() *Validation {
	return &UnifiedValidation{}
}

// Helper functions for common validation scenarios

// RequiredValidation creates validation that only requires a value
func RequiredValidation() *UnifiedValidation {
	return &UnifiedValidation{
		Required: true,
	}
}

// StringLengthValidation creates string length validation
func StringLengthValidation(minLength, maxLength int) *UnifiedValidation {
	return &UnifiedValidation{
		MinLength: &minLength,
		MaxLength: &maxLength,
	}
}

// EmailValidation creates email format validation
func EmailValidation() *UnifiedValidation {
	return &UnifiedValidation{
		Format: "email",
	}
}

// NumberRangeValidation creates number range validation
func NumberRangeValidation(min, max float64) *UnifiedValidation {
	return &UnifiedValidation{
		Min: &min,
		Max: &max,
	}
}

// ValidationBuilder provides fluent API for building validations
type ValidationBuilder struct {
	validation *UnifiedValidation
	*BuilderMixin
}

// NewValidationBuilder creates a new validation builder
func NewValidationBuilder() *ValidationBuilder {
	return &ValidationBuilder{
		validation:   &UnifiedValidation{},
		BuilderMixin: NewBuilderMixin(),
	}
}

// Required sets the field as required
func (b *ValidationBuilder) Required() *ValidationBuilder {
	b.validation.Required = true
	return b
}

// MinLength sets minimum string length
func (b *ValidationBuilder) MinLength(min int) *ValidationBuilder {
	if min < 0 {
		b.Context.AddErrorf("MinLength must be non-negative, got %d", min)
		return b
	}
	b.validation.MinLength = &min
	return b
}

// MaxLength sets maximum string length
func (b *ValidationBuilder) MaxLength(max int) *ValidationBuilder {
	if max < 0 {
		b.Context.AddErrorf("MaxLength must be non-negative, got %d", max)
		return b
	}
	if b.validation.MinLength != nil && max < *b.validation.MinLength {
		b.Context.AddErrorf("MaxLength (%d) cannot be less than MinLength (%d)", max, *b.validation.MinLength)
		return b
	}
	b.validation.MaxLength = &max
	return b
}

// Pattern sets regex pattern validation
func (b *ValidationBuilder) Pattern(pattern string) *ValidationBuilder {
	if pattern == "" {
		b.Context.AddError(fmt.Errorf("Pattern cannot be empty"))
		return b
	}
	// Validate the regex pattern
	if _, err := regexp.Compile(pattern); err != nil {
		b.Context.AddErrorf("Invalid regex pattern: %v", err)
		return b
	}
	b.validation.Pattern = pattern
	return b
}

// Email sets email format validation
func (b *ValidationBuilder) Email() *ValidationBuilder {
	b.validation.Format = "email"
	return b
}

// URL sets URL format validation
func (b *ValidationBuilder) URL() *ValidationBuilder {
	b.validation.Format = "url"
	return b
}

// NumberRange sets min and max number validation
func (b *ValidationBuilder) NumberRange(min, max float64) *ValidationBuilder {
	if min > max {
		b.Context.AddErrorf("Min (%g) cannot be greater than Max (%g)", min, max)
		return b
	}
	b.validation.Min = &min
	b.validation.Max = &max
	return b
}

// OnChange sets validation to trigger on change
func (b *ValidationBuilder) OnChange() *ValidationBuilder {
	b.validation.Mode = ValidationModeOnChange
	return b
}

// OnBlur sets validation to trigger on blur
func (b *ValidationBuilder) OnBlur() *ValidationBuilder {
	b.validation.Mode = ValidationModeOnBlur
	return b
}

// OnSubmit sets validation to trigger on submit
func (b *ValidationBuilder) OnSubmit() *ValidationBuilder {
	b.validation.Mode = ValidationModeOnSubmit
	return b
}

// Build implements BaseBuilder interface
func (b *ValidationBuilder) Build() (*UnifiedValidation, error) {
	if err := b.CheckBuild("Validation"); err != nil {
		return nil, err
	}
	return b.validation, nil
}

// MustBuild builds and panics on error
func (b *ValidationBuilder) MustBuild() *UnifiedValidation {
	result, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("ValidationBuilder.MustBuild() failed: %v", err))
	}
	return result
}

// Implement remaining BaseBuilder interface methods
func (b *ValidationBuilder) HasErrors() bool {
	return b.Context.HasErrors()
}

func (b *ValidationBuilder) GetErrors() []error {
	return b.Context.GetErrors()
}

func (b *ValidationBuilder) AddError(err error) BaseBuilder[UnifiedValidation] {
	b.Context.AddError(err)
	return b
}

func (b *ValidationBuilder) ClearErrors() BaseBuilder[UnifiedValidation] {
	b.Context.ClearErrors()
	return b
}