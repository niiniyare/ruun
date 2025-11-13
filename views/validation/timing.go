package validation

import (
	"context"
	"sync"
	"time"
)

// FieldDebouncer manages debounced validation timing for better UX
// Prevents excessive validation calls while user is typing
type FieldDebouncer struct {
	// Active timers for each field
	timers map[string]*time.Timer
	
	// Cancellation contexts for each field
	contexts map[string]context.CancelFunc
	
	// Configuration
	defaultDelay time.Duration
	
	// Thread safety
	mu sync.RWMutex
}

// NewFieldDebouncer creates a new field debouncer with default timing
func NewFieldDebouncer() *FieldDebouncer {
	return &FieldDebouncer{
		timers:       make(map[string]*time.Timer),
		contexts:     make(map[string]context.CancelFunc),
		defaultDelay: 300 * time.Millisecond,
	}
}

// WithDefaultDelay sets the default debounce delay
func (d *FieldDebouncer) WithDefaultDelay(delay time.Duration) *FieldDebouncer {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.defaultDelay = delay
	return d
}

// ScheduleValidation schedules a debounced validation for a field
// Cancels any existing validation timer for the same field
func (d *FieldDebouncer) ScheduleValidation(fieldName string, delay time.Duration, validationFunc func()) {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	// Use default delay if none specified
	if delay <= 0 {
		delay = d.defaultDelay
	}
	
	// Cancel existing timer for this field
	d.cancelFieldUnsafe(fieldName)
	
	// Create new context for this validation
	ctx, cancel := context.WithCancel(context.Background())
	d.contexts[fieldName] = cancel
	
	// Create timer for debounced execution
	timer := time.AfterFunc(delay, func() {
		// Check if context was cancelled before executing
		select {
		case <-ctx.Done():
			return // Validation was cancelled
		default:
			// Execute validation
			validationFunc()
			
			// Clean up after execution
			d.mu.Lock()
			delete(d.timers, fieldName)
			delete(d.contexts, fieldName)
			d.mu.Unlock()
		}
	})
	
	d.timers[fieldName] = timer
}

// CancelValidation cancels any pending validation for a specific field
func (d *FieldDebouncer) CancelValidation(fieldName string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	d.cancelFieldUnsafe(fieldName)
}

// cancelFieldUnsafe cancels validation for a field (must hold lock)
func (d *FieldDebouncer) cancelFieldUnsafe(fieldName string) {
	// Cancel timer
	if timer, exists := d.timers[fieldName]; exists {
		timer.Stop()
		delete(d.timers, fieldName)
	}
	
	// Cancel context
	if cancel, exists := d.contexts[fieldName]; exists {
		cancel()
		delete(d.contexts, fieldName)
	}
}

// CancelAllValidations cancels all pending validations
func (d *FieldDebouncer) CancelAllValidations() {
	d.mu.Lock()
	defer d.mu.Unlock()
	
	// Cancel all timers
	for fieldName := range d.timers {
		d.cancelFieldUnsafe(fieldName)
	}
}

// IsValidationPending checks if validation is pending for a field
func (d *FieldDebouncer) IsValidationPending(fieldName string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	_, exists := d.timers[fieldName]
	return exists
}

// GetPendingFields returns list of fields with pending validations
func (d *FieldDebouncer) GetPendingFields() []string {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	fields := make([]string, 0, len(d.timers))
	for fieldName := range d.timers {
		fields = append(fields, fieldName)
	}
	
	return fields
}

// GetPendingCount returns the number of fields with pending validations
func (d *FieldDebouncer) GetPendingCount() int {
	d.mu.RLock()
	defer d.mu.RUnlock()
	
	return len(d.timers)
}

// StateManager interface to avoid circular dependency
type StateManager interface {
	SetFieldErrors(fieldName string, errors []string)
	ClearFieldErrors(fieldName string)
	SetFieldTouched(fieldName string, touched bool)
	GetAllValues() map[string]any
	GetFieldErrors(fieldName string) []string
	IsFieldTouched(fieldName string) bool
	// Additional methods needed for validation orchestration
	GetValue(fieldName string) any
	SetValue(fieldName string, value any) error
	IsFieldDirty(fieldName string) bool
	SetFieldDirty(fieldName string, dirty bool)
}

// ValidationSummary provides a comprehensive view of validation states
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

// ValidationStateManager manages validation states with timing information
// EXTENDS the existing StateManager with validation-specific state tracking
type ValidationStateManager struct {
	// Enhanced validation states
	validationStates map[string]ValidationState
	validationTimes  map[string]time.Time
	pendingValidations map[string]context.CancelFunc
	
	// Integration with existing state manager
	stateManager StateManager
	
	// Thread safety
	mu sync.RWMutex
}

// ValidationState is defined in parent views package (views/state.go)
// We use the same type from the parent package to avoid duplication

// NewValidationStateManager creates a validation state manager that extends StateManager
func NewValidationStateManager(stateManager StateManager) *ValidationStateManager {
	return &ValidationStateManager{
		validationStates:   make(map[string]ValidationState),
		validationTimes:    make(map[string]time.Time),
		pendingValidations: make(map[string]context.CancelFunc),
		stateManager:       stateManager,
	}
}

// SetValidationState sets the validation state for a field
func (vsm *ValidationStateManager) SetValidationState(fieldName string, state ValidationState) {
	vsm.mu.Lock()
	defer vsm.mu.Unlock()
	
	vsm.validationStates[fieldName] = state
	vsm.validationTimes[fieldName] = time.Now()
}

// GetValidationState returns the current validation state for a field
func (vsm *ValidationStateManager) GetValidationState(fieldName string) ValidationState {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	if state, exists := vsm.validationStates[fieldName]; exists {
		return state
	}
	return ValidationStateIdle
}

// SetValidationTime sets the last validation time for a field
func (vsm *ValidationStateManager) SetValidationTime(fieldName string, t time.Time) {
	vsm.mu.Lock()
	defer vsm.mu.Unlock()
	
	vsm.validationTimes[fieldName] = t
}

// GetValidationTime returns the last validation time for a field
func (vsm *ValidationStateManager) GetValidationTime(fieldName string) time.Time {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	return vsm.validationTimes[fieldName]
}

// HasFieldsInState checks if any fields are in the specified validation state
func (vsm *ValidationStateManager) HasFieldsInState(state ValidationState) bool {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	for _, fieldState := range vsm.validationStates {
		if fieldState == state {
			return true
		}
	}
	return false
}

// GetFieldsInState returns list of fields in the specified validation state
func (vsm *ValidationStateManager) GetFieldsInState(state ValidationState) []string {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	fields := make([]string, 0)
	for fieldName, fieldState := range vsm.validationStates {
		if fieldState == state {
			fields = append(fields, fieldName)
		}
	}
	return fields
}

// GetAllValidationStates returns all current validation states
func (vsm *ValidationStateManager) GetAllValidationStates() map[string]ValidationState {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	states := make(map[string]ValidationState)
	for fieldName, state := range vsm.validationStates {
		states[fieldName] = state
	}
	return states
}

// ClearValidationState clears validation state for a field
func (vsm *ValidationStateManager) ClearValidationState(fieldName string) {
	vsm.mu.Lock()
	defer vsm.mu.Unlock()
	
	delete(vsm.validationStates, fieldName)
	delete(vsm.validationTimes, fieldName)
}

// ClearAllValidationStates clears all validation states
func (vsm *ValidationStateManager) ClearAllValidationStates() {
	vsm.mu.Lock()
	defer vsm.mu.Unlock()
	
	vsm.validationStates = make(map[string]ValidationState)
	vsm.validationTimes = make(map[string]time.Time)
}

// Bridge methods to existing StateManager

// SetFieldErrors delegates to StateManager
func (vsm *ValidationStateManager) SetFieldErrors(fieldName string, errors []string) {
	if vsm.stateManager != nil {
		vsm.stateManager.SetFieldErrors(fieldName, errors)
	}
}

// ClearFieldErrors delegates to StateManager
func (vsm *ValidationStateManager) ClearFieldErrors(fieldName string) {
	if vsm.stateManager != nil {
		vsm.stateManager.ClearFieldErrors(fieldName)
	}
}

// SetFieldTouched delegates to StateManager
func (vsm *ValidationStateManager) SetFieldTouched(fieldName string, touched bool) {
	if vsm.stateManager != nil {
		vsm.stateManager.SetFieldTouched(fieldName, touched)
	}
}

// GetAllValues delegates to StateManager
func (vsm *ValidationStateManager) GetAllValues() map[string]any {
	if vsm.stateManager != nil {
		return vsm.stateManager.GetAllValues()
	}
	return make(map[string]any)
}

// GetFieldErrors delegates to StateManager
func (vsm *ValidationStateManager) GetFieldErrors(fieldName string) []string {
	if vsm.stateManager != nil {
		return vsm.stateManager.GetFieldErrors(fieldName)
	}
	return []string{}
}

// IsFieldTouched delegates to StateManager
func (vsm *ValidationStateManager) IsFieldTouched(fieldName string) bool {
	if vsm.stateManager != nil {
		return vsm.stateManager.IsFieldTouched(fieldName)
	}
	return false
}

// GetValidationSummary returns a comprehensive validation summary
func (vsm *ValidationStateManager) GetValidationSummary() ValidationSummary {
	vsm.mu.RLock()
	defer vsm.mu.RUnlock()
	
	summary := ValidationSummary{
		States:         make(map[string]ValidationState),
		LastValidated:  make(map[string]time.Time),
		TotalFields:    len(vsm.validationStates),
	}
	
	// Copy current states
	for fieldName, state := range vsm.validationStates {
		summary.States[fieldName] = state
		
		// Count states
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
	for fieldName, t := range vsm.validationTimes {
		summary.LastValidated[fieldName] = t
	}
	
	return summary
}

