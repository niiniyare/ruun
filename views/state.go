package views

import (
	"context"
	"fmt"
	"sync"

	"github.com/niiniyare/ruun/pkg/schema"
)

// StateManager manages form state including values, errors, and validation state
// This provides thread-safe state management for form interactions
type StateManager struct {
	mu sync.RWMutex

	schema      *schema.Schema
	values      map[string]any
	errors      map[string][]string
	touched     map[string]bool
	dirty       map[string]bool
	initialData map[string]any
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

	// Handle interface{} arrays
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

// interfaceSlicesEqual compares interface{} slices for equality
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
