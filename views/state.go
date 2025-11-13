package views

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ValidationState represents the validation state of a field (for UI orchestration)
type ValidationState int

const (
	ValidationStateIdle       ValidationState = iota // No validation activity
	ValidationStateValidating                        // Currently validating
	ValidationStateValid                             // Validation passed
	ValidationStateInvalid                           // Validation failed
	ValidationStateWarning                           // Validation warning
)

// String returns string representation of validation state
func (vs ValidationState) String() string {
	switch vs {
	case ValidationStateIdle:
		return "idle"
	case ValidationStateValidating:
		return "validating"
	case ValidationStateValid:
		return "valid"
	case ValidationStateInvalid:
		return "invalid"
	case ValidationStateWarning:
		return "warning"
	default:
		return "unknown"
	}
}

// StateManager manages form state including values, errors, and validation state
// This provides thread-safe state management for form interactions
// EXTENDED with UI validation orchestration state tracking
type StateManager struct {
	mu sync.RWMutex

	schema      *schema.Schema
	values      map[string]any
	errors      map[string][]string
	touched     map[string]bool
	dirty       map[string]bool
	initialData map[string]any
	
	// NEW: UI validation orchestration state tracking
	validationStates map[string]ValidationState // Tracks validation UI states
	validationTimes  map[string]time.Time       // Last validation timestamp
}

// NewStateManager creates a new state manager with initial data
func NewStateManager(s *schema.Schema, initialData map[string]any) *StateManager {
	if initialData == nil {
		initialData = make(map[string]any)
	}

	sm := &StateManager{
		schema:      s,
		values:      make(map[string]any),
		errors:      make(map[string][]string),
		touched:     make(map[string]bool),
		dirty:       make(map[string]bool),
		initialData: make(map[string]any),
		
		// NEW: Initialize validation orchestration state
		validationStates: make(map[string]ValidationState),
		validationTimes:  make(map[string]time.Time),
	}

	// Copy initial data
	for k, v := range initialData {
		sm.initialData[k] = v
		sm.values[k] = v
	}

	// Set default values for fields that don't have initial data
	sm.initializeDefaultValues()

	return sm
}

// GetAllValues returns all current form values
func (sm *StateManager) GetAllValues() map[string]any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	result := make(map[string]any, len(sm.values))
	for k, v := range sm.values {
		result[k] = v
	}
	return result
}

// GetAllErrors returns all current validation errors
func (sm *StateManager) GetAllErrors() map[string][]string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	result := make(map[string][]string, len(sm.errors))
	for k, v := range sm.errors {
		if len(v) > 0 {
			result[k] = make([]string, len(v))
			copy(result[k], v)
		}
	}
	return result
}

// GetValue returns the current value for a specific field
func (sm *StateManager) GetValue(fieldName string) any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return sm.values[fieldName]
}

// GetFieldErrors returns validation errors for a specific field
func (sm *StateManager) GetFieldErrors(fieldName string) []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	if errors, exists := sm.errors[fieldName]; exists {
		result := make([]string, len(errors))
		copy(result, errors)
		return result
	}
	return []string{}
}

// IsFieldTouched returns true if the user has interacted with the field
func (sm *StateManager) IsFieldTouched(fieldName string) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return sm.touched[fieldName]
}

// IsFieldDirty returns true if the field value differs from its initial value
func (sm *StateManager) IsFieldDirty(fieldName string) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return sm.dirty[fieldName]
}

// SetFieldDirty manually sets the dirty state for a field
func (sm *StateManager) SetFieldDirty(fieldName string, dirty bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.dirty[fieldName] = dirty
}

// SetValue updates a field value and recalculates dirty state
func (sm *StateManager) SetValue(fieldName string, value any) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Validate field exists in schema
	field, exists := sm.schema.GetField(fieldName)
	if !exists {
		return fmt.Errorf("field %s not found in schema", fieldName)
	}

	// Store the value
	sm.values[fieldName] = value

	// Mark as touched
	sm.touched[fieldName] = true

	// Calculate dirty state
	initialValue, hasInitial := sm.initialData[fieldName]
	sm.dirty[fieldName] = hasInitial && !sm.valuesEqual(initialValue, value)

	// Clear previous errors for this field
	delete(sm.errors, fieldName)

	// Validate the field immediately
	ctx := context.Background() // TODO: Pass context through if needed
	if err := sm.validateFieldInternal(ctx, &field, value); err != nil {
		sm.setFieldErrors(fieldName, []string{err.Error()})
	}

	return nil
}

// SetFieldTouched marks a field as touched
func (sm *StateManager) SetFieldTouched(fieldName string, touched bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.touched[fieldName] = touched
}

// SetFieldErrors sets validation errors for a field
func (sm *StateManager) SetFieldErrors(fieldName string, errors []string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.setFieldErrors(fieldName, errors)
}

// ClearFieldErrors removes all errors for a field
func (sm *StateManager) ClearFieldErrors(fieldName string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	delete(sm.errors, fieldName)
}

// ClearAllErrors removes all validation errors
func (sm *StateManager) ClearAllErrors() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.errors = make(map[string][]string)
}

// ValidateAll validates all fields and returns the first error encountered
func (sm *StateManager) ValidateAll(ctx context.Context) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Clear all existing errors
	sm.errors = make(map[string][]string)

	var firstError error

	// Validate each field in the schema
	for _, field := range sm.schema.Fields {
		value := sm.values[field.Name]
		if err := sm.validateFieldInternal(ctx, &field, value); err != nil {
			sm.setFieldErrors(field.Name, []string{err.Error()})
			if firstError == nil {
				firstError = err
			}
		}
	}

	return firstError
}

// ValidateField validates a single field
func (sm *StateManager) ValidateField(ctx context.Context, fieldName string) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	field, exists := sm.schema.GetField(fieldName)
	if !exists {
		return fmt.Errorf("field %s not found in schema", fieldName)
	}

	value := sm.values[fieldName]
	if err := sm.validateFieldInternal(ctx, &field, value); err != nil {
		sm.setFieldErrors(fieldName, []string{err.Error()})
		return err
	}

	// Clear errors if validation passed
	delete(sm.errors, fieldName)
	return nil
}

// IsValid returns true if there are no validation errors
func (sm *StateManager) IsValid() bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, errors := range sm.errors {
		if len(errors) > 0 {
			return false
		}
	}
	return true
}

// IsDirty returns true if any field has been modified
func (sm *StateManager) IsDirty() bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	for _, dirty := range sm.dirty {
		if dirty {
			return true
		}
	}
	return false
}

// Reset resets all form state to initial values
func (sm *StateManager) Reset() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Reset values to initial data
	sm.values = make(map[string]any)
	for k, v := range sm.initialData {
		sm.values[k] = v
	}

	// Re-initialize default values
	sm.initializeDefaultValuesUnsafe()

	// Clear all state flags
	sm.errors = make(map[string][]string)
	sm.touched = make(map[string]bool)
	sm.dirty = make(map[string]bool)
}

// GetFormState returns a comprehensive state object
func (sm *StateManager) GetFormState() FormState {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return FormState{
		Values:  sm.GetAllValues(),
		Errors:  sm.GetAllErrors(),
		Touched: sm.copyBoolMap(sm.touched),
		Dirty:   sm.copyBoolMap(sm.dirty),
		Valid:   sm.isValidUnsafe(),
	}
}

// FormState represents the complete form state
type FormState struct {
	Values  map[string]any      `json:"values"`
	Errors  map[string][]string `json:"errors"`
	Touched map[string]bool     `json:"touched"`
	Dirty   map[string]bool     `json:"dirty"`
	Valid   bool                `json:"valid"`
}

// Private helper methods

// validateFieldInternal validates a field without acquiring locks
func (sm *StateManager) validateFieldInternal(ctx context.Context, field *schema.Field, value any) error {
	if field == nil {
		return fmt.Errorf("field cannot be nil")
	}

	// Use the field's built-in validation
	return field.ValidateValue(ctx, value)
}

// setFieldErrors sets errors without acquiring locks
func (sm *StateManager) setFieldErrors(fieldName string, errors []string) {
	if len(errors) > 0 {
		sm.errors[fieldName] = errors
	} else {
		delete(sm.errors, fieldName)
	}
}

// initializeDefaultValues sets default values for all fields
func (sm *StateManager) initializeDefaultValues() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.initializeDefaultValuesUnsafe()
}

// initializeDefaultValuesUnsafe sets default values without acquiring locks
func (sm *StateManager) initializeDefaultValuesUnsafe() {
	ctx := context.Background()

	for _, field := range sm.schema.Fields {
		// Skip if value already exists
		if _, exists := sm.values[field.Name]; exists {
			continue
		}

		// Get default value for the field
		defaultValue, err := field.GetDefaultValue(ctx)
		if err == nil && defaultValue != nil {
			sm.values[field.Name] = defaultValue
		} else {
			// Set type-appropriate default
			sm.values[field.Name] = sm.getTypeDefault(field.Type)
		}
	}
}

// getTypeDefault returns a default value for the field type
func (sm *StateManager) getTypeDefault(fieldType schema.FieldType) any {
	switch fieldType {
	case schema.FieldCheckbox, schema.FieldSwitch:
		return false
	case schema.FieldMultiSelect, schema.FieldCheckboxes, schema.FieldTags:
		return []string{}
	case schema.FieldNumber, schema.FieldCurrency, schema.FieldSlider, schema.FieldRating:
		return 0
	default:
		return ""
	}
}

// valuesEqual compares two values for equality
func (sm *StateManager) valuesEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	// Handle string arrays (common for multi-select)
	if aSlice, aOk := a.([]string); aOk {
		if bSlice, bOk := b.([]string); bOk {
			return sm.stringSlicesEqual(aSlice, bSlice)
		}
	}

	// Handle any arrays
	if aSlice, aOk := a.([]any); aOk {
		if bSlice, bOk := b.([]any); bOk {
			return sm.interfaceSlicesEqual(aSlice, bSlice)
		}
	}

	// Basic value comparison
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

// stringSlicesEqual compares string slices for equality
func (sm *StateManager) stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// interfaceSlicesEqual compares any slices for equality
func (sm *StateManager) interfaceSlicesEqual(a, b []any) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if fmt.Sprintf("%v", v) != fmt.Sprintf("%v", b[i]) {
			return false
		}
	}
	return true
}

// copyBoolMap creates a copy of a boolean map
func (sm *StateManager) copyBoolMap(original map[string]bool) map[string]bool {
	result := make(map[string]bool, len(original))
	for k, v := range original {
		result[k] = v
	}
	return result
}

// isValidUnsafe checks validity without acquiring locks
func (sm *StateManager) isValidUnsafe() bool {
	for _, errors := range sm.errors {
		if len(errors) > 0 {
			return false
		}
	}
	return true
}

// UpdateInitialData updates the initial data reference point
// This is useful after successful form submission to reset the "dirty" baseline
func (sm *StateManager) UpdateInitialData(newInitialData map[string]any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Clear existing initial data
	sm.initialData = make(map[string]any)

	// Set new initial data
	if newInitialData != nil {
		for k, v := range newInitialData {
			sm.initialData[k] = v
		}
	}

	// Recalculate dirty state for all fields
	for fieldName := range sm.values {
		initialValue, hasInitial := sm.initialData[fieldName]
		currentValue := sm.values[fieldName]
		sm.dirty[fieldName] = hasInitial && !sm.valuesEqual(initialValue, currentValue)
	}
}

// GetChangedValues returns only the values that have changed from initial
func (sm *StateManager) GetChangedValues() map[string]any {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	result := make(map[string]any)
	for fieldName, isDirty := range sm.dirty {
		if isDirty {
			result[fieldName] = sm.values[fieldName]
		}
	}
	return result
}

// SetInitialValue sets the initial value for a field (used for dynamic forms)
func (sm *StateManager) SetInitialValue(fieldName string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.initialData[fieldName] = value

	// If the field doesn't have a current value, set it to the initial value
	if _, exists := sm.values[fieldName]; !exists {
		sm.values[fieldName] = value
	}

	// Recalculate dirty state
	currentValue := sm.values[fieldName]
	sm.dirty[fieldName] = !sm.valuesEqual(value, currentValue)
}

// ═══════════════════════════════════════════════════════════════════════════
// UI Validation Orchestration Methods (EXTENSION)
// ═══════════════════════════════════════════════════════════════════════════

// SetValidationState sets the validation UI state for a field
func (sm *StateManager) SetValidationState(fieldName string, state ValidationState) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	sm.validationStates[fieldName] = state
	sm.validationTimes[fieldName] = time.Now()
}

// GetValidationState returns the current validation state for a field
func (sm *StateManager) GetValidationState(fieldName string) ValidationState {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	if state, exists := sm.validationStates[fieldName]; exists {
		return state
	}
	return ValidationStateIdle
}

// SetValidationTime sets the last validation time for a field
func (sm *StateManager) SetValidationTime(fieldName string, t time.Time) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	sm.validationTimes[fieldName] = t
}

// GetValidationTime returns the last validation time for a field
func (sm *StateManager) GetValidationTime(fieldName string) time.Time {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	return sm.validationTimes[fieldName]
}

// HasFieldsInState checks if any fields are in the specified validation state
func (sm *StateManager) HasFieldsInState(state ValidationState) bool {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	for _, fieldState := range sm.validationStates {
		if fieldState == state {
			return true
		}
	}
	return false
}

// GetFieldsInState returns list of fields in the specified validation state
func (sm *StateManager) GetFieldsInState(state ValidationState) []string {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	fields := make([]string, 0)
	for fieldName, fieldState := range sm.validationStates {
		if fieldState == state {
			fields = append(fields, fieldName)
		}
	}
	return fields
}

// GetAllValidationStates returns all current validation states
func (sm *StateManager) GetAllValidationStates() map[string]ValidationState {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	states := make(map[string]ValidationState)
	for fieldName, state := range sm.validationStates {
		states[fieldName] = state
	}
	return states
}

// ClearValidationState clears validation state for a field
func (sm *StateManager) ClearValidationState(fieldName string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	delete(sm.validationStates, fieldName)
	delete(sm.validationTimes, fieldName)
}

// ClearAllValidationStates clears all validation states
func (sm *StateManager) ClearAllValidationStates() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	
	sm.validationStates = make(map[string]ValidationState)
	sm.validationTimes = make(map[string]time.Time)
}

// IsValidationInProgress checks if any field is currently being validated
func (sm *StateManager) IsValidationInProgress() bool {
	return sm.HasFieldsInState(ValidationStateValidating)
}

// GetValidationSummary returns a comprehensive validation state summary
type ValidationSummary struct {
	States            map[string]ValidationState `json:"states"`
	LastValidated     map[string]time.Time       `json:"lastValidated"`
	TotalFields       int                        `json:"totalFields"`
	ValidFields       int                        `json:"validFields"`
	InvalidFields     int                        `json:"invalidFields"`
	ValidatingFields  int                        `json:"validatingFields"`
	WarningFields     int                        `json:"warningFields"`
}

// IsFormValid returns true if all fields are valid (no invalid or validating fields)
func (vs ValidationSummary) IsFormValid() bool {
	return vs.InvalidFields == 0 && vs.ValidatingFields == 0
}

// IsFormValidating returns true if any fields are currently validating
func (vs ValidationSummary) IsFormValidating() bool {
	return vs.ValidatingFields > 0
}

// GetValidationSummary returns a comprehensive validation summary
func (sm *StateManager) GetValidationSummary() ValidationSummary {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	
	summary := ValidationSummary{
		States:         make(map[string]ValidationState),
		LastValidated:  make(map[string]time.Time),
		TotalFields:    len(sm.validationStates),
	}
	
	// Copy current states and count them
	for fieldName, state := range sm.validationStates {
		summary.States[fieldName] = state
		
		switch state {
		case ValidationStateValid:
			summary.ValidFields++
		case ValidationStateInvalid:
			summary.InvalidFields++
		case ValidationStateValidating:
			summary.ValidatingFields++
		case ValidationStateWarning:
			summary.WarningFields++
		}
	}
	
	// Copy validation times
	for fieldName, t := range sm.validationTimes {
		summary.LastValidated[fieldName] = t
	}
	
	return summary
}
