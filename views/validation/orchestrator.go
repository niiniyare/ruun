package views

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ValidationTrigger defines when validation should occur
type ValidationTrigger int

const (
	ValidationTriggerImmediate ValidationTrigger = iota // Immediate validation (for required fields)
	ValidationTriggerDebounced                          // Debounced validation (for patterns, async checks)
	ValidationTriggerOnBlur                             // On-blur validation
	ValidationTriggerOnSubmit                           // On-submit validation only
)

// UIValidationOrchestrator coordinates validation timing and UI state
// REUSES existing schema validation infrastructure completely
type UIValidationOrchestrator struct {
	// REUSE existing schema validation (from pkg/schema/validate/)
	schema    *schema.Schema
	validator schema.RuntimeValidator // Existing validator interface

	// NEW: UI-specific timing and state control
	debouncer *FieldDebouncer
	stateMgr  *ValidationStateManager
	eventBus  *ValidationEventBus

	// Configuration
	config *ValidationConfig

	// Thread safety
	mu sync.RWMutex
}

// ValidationConfig holds UI validation configuration
type ValidationConfig struct {
	// Timing configuration
	ImmediateFields    []string      // Fields that validate immediately
	DebounceDelay      time.Duration // Default debounce delay
	OnBlurDelay        time.Duration // Delay for onBlur validation
	AsyncTimeoutDelay  time.Duration // Timeout for async validation

	// Behavior
	ValidateOnChange   bool // Enable change-based validation
	ValidateOnBlur     bool // Enable blur-based validation
	ShowErrorsOnlyOnBlur bool // Only show errors after blur
	ClearErrorsOnChange bool // Clear errors immediately on change

	// Client-side validation
	EnableClientRules  bool // Enable client-side rule extraction
}

// NewUIValidationOrchestrator creates a new UI validation orchestrator
func NewUIValidationOrchestrator(
	schema *schema.Schema,
	validator schema.RuntimeValidator,
	stateMgr *ValidationStateManager,
) *UIValidationOrchestrator {
	config := &ValidationConfig{
		DebounceDelay:        300 * time.Millisecond,
		OnBlurDelay:          100 * time.Millisecond,
		AsyncTimeoutDelay:    5 * time.Second,
		ValidateOnChange:     true,
		ValidateOnBlur:       true,
		ShowErrorsOnlyOnBlur: false,
		ClearErrorsOnChange:  true,
		EnableClientRules:    true,
	}

	orchestrator := &UIValidationOrchestrator{
		schema:    schema,
		validator: validator,
		stateMgr:  stateMgr,
		config:    config,
		debouncer: NewFieldDebouncer(),
		eventBus:  NewValidationEventBus(),
	}

	// Connect event bus to state manager for real-time updates
	orchestrator.connectEventHandlers()

	return orchestrator
}

// WithConfig sets custom validation configuration
func (o *UIValidationOrchestrator) WithConfig(config *ValidationConfig) *UIValidationOrchestrator {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.config = config
	return o
}

// ValidateFieldWithUI orchestrates validation with UI timing control
// This is the main entry point for UI-driven validation
func (o *UIValidationOrchestrator) ValidateFieldWithUI(
	ctx context.Context,
	fieldName string,
	value interface{},
	trigger ValidationTrigger,
) error {
	o.mu.RLock()
	schema := o.schema
	validator := o.validator
	config := o.config
	o.mu.RUnlock()

	if validator == nil {
		return fmt.Errorf("no validator configured")
	}

	// REUSE existing field validation from schema
	field, exists := schema.GetField(fieldName)
	if !exists {
		return fmt.Errorf("field %s not found", fieldName)
	}

	// Set UI state: validating
	o.stateMgr.SetValidationState(fieldName, ValidationStateValidating)
	o.eventBus.EmitValidationStart(fieldName)

	// Use existing schema validation based on trigger
	switch trigger {
	case ValidationTriggerImmediate:
		return o.validateImmediate(ctx, field, value, config)
	case ValidationTriggerDebounced:
		return o.validateDebounced(ctx, field, value, config)
	case ValidationTriggerOnBlur:
		return o.validateOnBlur(ctx, field, value, config)
	case ValidationTriggerOnSubmit:
		return o.validateOnSubmit(ctx, field, value, config)
	default:
		return o.validateImmediate(ctx, field, value, config)
	}
}

// validateImmediate performs immediate validation using existing schema validation
func (o *UIValidationOrchestrator) validateImmediate(
	ctx context.Context,
	field schema.Field,
	value interface{},
	config *ValidationConfig,
) error {
	// REUSE: Use existing schema RuntimeValidator interface
	allData := o.stateMgr.GetAllValues()
	errors := o.validator.ValidateField(ctx, &field, value, allData)

	fieldName := field.Name

	if len(errors) > 0 {
		// Handle validation errors
		o.handleValidationErrors(fieldName, errors)
		
		// Emit validation error event
		o.eventBus.EmitValidationError(fieldName, errors)
		
		// Update UI state
		o.stateMgr.SetValidationState(fieldName, ValidationStateInvalid)
		
		return fmt.Errorf("validation failed: %v", errors)
	}

	// Validation passed
	o.handleValidationSuccess(fieldName)
	o.eventBus.EmitValidationSuccess(fieldName)
	o.stateMgr.SetValidationState(fieldName, ValidationStateValid)
	
	return nil
}

// validateDebounced performs debounced validation for better UX
func (o *UIValidationOrchestrator) validateDebounced(
	ctx context.Context,
	field schema.Field,
	value interface{},
	config *ValidationConfig,
) error {
	fieldName := field.Name

	// Cancel any pending validation for this field
	o.debouncer.CancelValidation(fieldName)

	// Schedule debounced validation
	validationFunc := func() {
		if err := o.validateImmediate(ctx, field, value, config); err != nil {
			// Error already handled in validateImmediate
			return
		}
	}

	o.debouncer.ScheduleValidation(fieldName, config.DebounceDelay, validationFunc)
	
	// Clear errors immediately if configured
	if config.ClearErrorsOnChange {
		o.stateMgr.ClearFieldErrors(fieldName)
		o.stateMgr.SetValidationState(fieldName, ValidationStateIdle)
	}

	return nil
}

// validateOnBlur performs validation on field blur event
func (o *UIValidationOrchestrator) validateOnBlur(
	ctx context.Context,
	field schema.Field,
	value interface{},
	config *ValidationConfig,
) error {
	// Short delay for onBlur validation to improve UX
	time.Sleep(config.OnBlurDelay)
	
	return o.validateImmediate(ctx, field, value, config)
}

// validateOnSubmit performs validation only on form submission
func (o *UIValidationOrchestrator) validateOnSubmit(
	ctx context.Context,
	field schema.Field,
	value interface{},
	config *ValidationConfig,
) error {
	// Only validate during form submission - no immediate feedback
	o.stateMgr.SetValidationState(field.Name, ValidationStateIdle)
	return nil
}

// ValidateAllFieldsWithUI validates all fields using UI orchestration
func (o *UIValidationOrchestrator) ValidateAllFieldsWithUI(ctx context.Context) map[string][]string {
	o.mu.RLock()
	schema := o.schema
	validator := o.validator
	o.mu.RUnlock()

	if validator == nil {
		return make(map[string][]string)
	}

	// REUSE: Use existing schema RuntimeValidator.ValidateAllFields
	allData := o.stateMgr.GetAllValues()
	allErrors := validator.ValidateAllFields(ctx, schema, allData)

	// Update UI state for all fields
	for _, field := range schema.Fields {
		fieldName := field.Name
		
		if errors, hasErrors := allErrors[fieldName]; hasErrors && len(errors) > 0 {
			o.handleValidationErrors(fieldName, errors)
			o.stateMgr.SetValidationState(fieldName, ValidationStateInvalid)
			o.eventBus.EmitValidationError(fieldName, errors)
		} else {
			o.handleValidationSuccess(fieldName)
			o.stateMgr.SetValidationState(fieldName, ValidationStateValid)
			o.eventBus.EmitValidationSuccess(fieldName)
		}
	}

	return allErrors
}

// GetValidationStateForField returns the current validation state for a field
func (o *UIValidationOrchestrator) GetValidationStateForField(fieldName string) ValidationState {
	return o.stateMgr.GetValidationState(fieldName)
}

// IsValidationInProgress checks if any field is currently being validated
func (o *UIValidationOrchestrator) IsValidationInProgress() bool {
	return o.stateMgr.HasFieldsInState(ValidationStateValidating)
}

// GetPendingValidations returns fields with pending validations
func (o *UIValidationOrchestrator) GetPendingValidations() []string {
	return o.debouncer.GetPendingFields()
}

// connectEventHandlers connects the event bus to the state manager
func (o *UIValidationOrchestrator) connectEventHandlers() {
	// Connect validation events to state updates
	o.eventBus.AddListener("*", func(event ValidationEvent) {
		// Update validation timestamps
		o.stateMgr.SetValidationTime(event.FieldName, event.Timestamp)
		
		// Additional event-driven state updates can be added here
	})
}

// handleValidationErrors processes validation errors from schema validation
func (o *UIValidationOrchestrator) handleValidationErrors(fieldName string, errors []string) {
	// Update state manager with errors
	o.stateMgr.SetFieldErrors(fieldName, errors)
	
	// Mark field as touched since we validated it
	o.stateMgr.SetFieldTouched(fieldName, true)
}

// handleValidationSuccess processes successful validation
func (o *UIValidationOrchestrator) handleValidationSuccess(fieldName string) {
	// Clear any existing errors
	o.stateMgr.ClearFieldErrors(fieldName)
	
	// Mark field as touched
	o.stateMgr.SetFieldTouched(fieldName, true)
}

// OnFieldChange handles field value changes with appropriate validation trigger
func (o *UIValidationOrchestrator) OnFieldChange(
	ctx context.Context,
	fieldName string,
	newValue interface{},
) error {
	o.mu.RLock()
	config := o.config
	o.mu.RUnlock()

	if !config.ValidateOnChange {
		return nil // Validation on change disabled
	}

	// Determine validation trigger based on field and configuration
	trigger := o.determineValidationTrigger(fieldName)
	
	return o.ValidateFieldWithUI(ctx, fieldName, newValue, trigger)
}

// OnFieldBlur handles field blur events
func (o *UIValidationOrchestrator) OnFieldBlur(
	ctx context.Context,
	fieldName string,
	value interface{},
) error {
	o.mu.RLock()
	config := o.config
	o.mu.RUnlock()

	if !config.ValidateOnBlur {
		return nil // Validation on blur disabled
	}

	return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerOnBlur)
}

// determineValidationTrigger determines the appropriate validation trigger for a field
func (o *UIValidationOrchestrator) determineValidationTrigger(fieldName string) ValidationTrigger {
	// Check if field should validate immediately
	for _, immediatField := range o.config.ImmediateFields {
		if immediatField == fieldName {
			return ValidationTriggerImmediate
		}
	}

	// Get field from schema to check validation complexity
	field, exists := o.schema.GetField(fieldName)
	if !exists {
		return ValidationTriggerDebounced // Safe default
	}

	// Use immediate validation for simple checks
	if field.Validation != nil {
		// Required fields - immediate feedback
		if field.Required {
			return ValidationTriggerImmediate
		}
		
		// Complex validation (async, business rules) - debounced
		if field.Validation.Unique || field.Validation.Custom != "" {
			return ValidationTriggerDebounced
		}
	}

	// Default to debounced for better UX
	return ValidationTriggerDebounced
}

// Cleanup cancels all pending validations
func (o *UIValidationOrchestrator) Cleanup() {
	o.debouncer.CancelAllValidations()
}