package views

import (
	"sync"
	"time"
)

// ValidationEvent represents a validation state change event
type ValidationEvent struct {
	FieldName string            `json:"fieldName"`
	EventType ValidationEventType `json:"eventType"`
	State     ValidationState   `json:"state"`
	Errors    []string          `json:"errors,omitempty"`
	Timestamp time.Time         `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// ValidationEventType defines types of validation events
type ValidationEventType string

const (
	ValidationEventStart   ValidationEventType = "validation_start"
	ValidationEventSuccess ValidationEventType = "validation_success"
	ValidationEventError   ValidationEventType = "validation_error"
	ValidationEventCancel  ValidationEventType = "validation_cancel"
	ValidationEventClear   ValidationEventType = "validation_clear"
)

// ValidationEventListener handles validation events
type ValidationEventListener func(event ValidationEvent)

// ValidationEventBus propagates validation state changes to UI components
// Provides real-time feedback to FormField components and other UI elements
type ValidationEventBus struct {
	// Field-specific listeners
	listeners map[string][]ValidationEventListener
	
	// Global listeners (listen to all fields)
	globalListeners []ValidationEventListener
	
	// Event history for debugging
	eventHistory []ValidationEvent
	maxHistory   int
	
	// Thread safety
	mu sync.RWMutex
}

// NewValidationEventBus creates a new validation event bus
func NewValidationEventBus() *ValidationEventBus {
	return &ValidationEventBus{
		listeners:       make(map[string][]ValidationEventListener),
		globalListeners: make([]ValidationEventListener, 0),
		eventHistory:    make([]ValidationEvent, 0),
		maxHistory:      100, // Keep last 100 events for debugging
	}
}

// AddListener adds a listener for specific field events
func (b *ValidationEventBus) AddListener(fieldName string, listener ValidationEventListener) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if fieldName == "*" {
		// Global listener - listens to all field events
		b.globalListeners = append(b.globalListeners, listener)
	} else {
		// Field-specific listener
		if b.listeners[fieldName] == nil {
			b.listeners[fieldName] = make([]ValidationEventListener, 0)
		}
		b.listeners[fieldName] = append(b.listeners[fieldName], listener)
	}
}

// RemoveListener removes a listener for a specific field
// Note: This removes all listeners for the field - in production you'd want listener IDs
func (b *ValidationEventBus) RemoveListener(fieldName string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	if fieldName == "*" {
		b.globalListeners = make([]ValidationEventListener, 0)
	} else {
		delete(b.listeners, fieldName)
	}
}

// EmitValidationStart emits a validation start event
func (b *ValidationEventBus) EmitValidationStart(fieldName string) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: ValidationEventStart,
		State:     ValidationStateValidating,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"trigger": "user_input",
		},
	}
	
	b.emitEvent(event)
}

// EmitValidationSuccess emits a validation success event
func (b *ValidationEventBus) EmitValidationSuccess(fieldName string) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: ValidationEventSuccess,
		State:     ValidationStateValid,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"result": "valid",
		},
	}
	
	b.emitEvent(event)
}

// EmitValidationError emits a validation error event
func (b *ValidationEventBus) EmitValidationError(fieldName string, errors []string) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: ValidationEventError,
		State:     ValidationStateInvalid,
		Errors:    errors,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"errorCount": len(errors),
			"result":     "invalid",
		},
	}
	
	b.emitEvent(event)
}

// EmitValidationCancel emits a validation cancel event
func (b *ValidationEventBus) EmitValidationCancel(fieldName string) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: ValidationEventCancel,
		State:     ValidationStateIdle,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"reason": "debounce_cancel",
		},
	}
	
	b.emitEvent(event)
}

// EmitValidationClear emits a validation clear event
func (b *ValidationEventBus) EmitValidationClear(fieldName string) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: ValidationEventClear,
		State:     ValidationStateIdle,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"action": "clear_errors",
		},
	}
	
	b.emitEvent(event)
}

// emitEvent sends an event to all relevant listeners
func (b *ValidationEventBus) emitEvent(event ValidationEvent) {
	b.mu.Lock()
	
	// Add to event history
	b.addToHistory(event)
	
	// Get listeners while holding lock
	fieldListeners := make([]ValidationEventListener, 0)
	if listeners, exists := b.listeners[event.FieldName]; exists {
		fieldListeners = make([]ValidationEventListener, len(listeners))
		copy(fieldListeners, listeners)
	}
	
	globalListeners := make([]ValidationEventListener, len(b.globalListeners))
	copy(globalListeners, b.globalListeners)
	
	b.mu.Unlock()
	
	// Emit to field-specific listeners
	for _, listener := range fieldListeners {
		go func(l ValidationEventListener) {
			defer func() {
				if r := recover(); r != nil {
					// Log error in production
					// For now, silently continue
				}
			}()
			l(event)
		}(listener)
	}
	
	// Emit to global listeners
	for _, listener := range globalListeners {
		go func(l ValidationEventListener) {
			defer func() {
				if r := recover(); r != nil {
					// Log error in production
					// For now, silently continue
				}
			}()
			l(event)
		}(listener)
	}
}

// addToHistory adds event to history with size limit
func (b *ValidationEventBus) addToHistory(event ValidationEvent) {
	b.eventHistory = append(b.eventHistory, event)
	
	// Keep only the last maxHistory events
	if len(b.eventHistory) > b.maxHistory {
		b.eventHistory = b.eventHistory[len(b.eventHistory)-b.maxHistory:]
	}
}

// GetEventHistory returns recent validation events for debugging
func (b *ValidationEventBus) GetEventHistory() []ValidationEvent {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	// Return a copy to prevent external modification
	history := make([]ValidationEvent, len(b.eventHistory))
	copy(history, b.eventHistory)
	
	return history
}

// GetEventsForField returns recent events for a specific field
func (b *ValidationEventBus) GetEventsForField(fieldName string) []ValidationEvent {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	fieldEvents := make([]ValidationEvent, 0)
	for _, event := range b.eventHistory {
		if event.FieldName == fieldName {
			fieldEvents = append(fieldEvents, event)
		}
	}
	
	return fieldEvents
}

// ClearHistory clears the event history
func (b *ValidationEventBus) ClearHistory() {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	b.eventHistory = make([]ValidationEvent, 0)
}

// ConnectToFormField connects validation events to FormField component updates
// This provides the bridge between validation events and UI component state
func (b *ValidationEventBus) ConnectToFormField(fieldName string, updateFunc func(ValidationEvent)) {
	b.AddListener(fieldName, func(event ValidationEvent) {
		// Only call updateFunc for events that affect the UI
		switch event.EventType {
		case ValidationEventStart, ValidationEventSuccess, ValidationEventError, ValidationEventClear:
			updateFunc(event)
		case ValidationEventCancel:
			// Cancel events might not need UI updates unless showing loading states
			updateFunc(event)
		}
	})
}

// ConnectToGlobalValidationState connects to global validation state changes
// Useful for form-level validation indicators
func (b *ValidationEventBus) ConnectToGlobalValidationState(updateFunc func(ValidationEvent)) {
	b.AddListener("*", updateFunc)
}

// GetListenerCount returns the number of listeners (for debugging)
func (b *ValidationEventBus) GetListenerCount() map[string]int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	counts := make(map[string]int)
	
	for fieldName, listeners := range b.listeners {
		counts[fieldName] = len(listeners)
	}
	
	counts["*"] = len(b.globalListeners)
	
	return counts
}

// HasListenersForField checks if there are any listeners for a specific field
func (b *ValidationEventBus) HasListenersForField(fieldName string) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	if listeners, exists := b.listeners[fieldName]; exists && len(listeners) > 0 {
		return true
	}
	
	return len(b.globalListeners) > 0
}

// EmitCustomEvent emits a custom validation event with metadata
func (b *ValidationEventBus) EmitCustomEvent(fieldName string, eventType ValidationEventType, state ValidationState, metadata map[string]interface{}) {
	event := ValidationEvent{
		FieldName: fieldName,
		EventType: eventType,
		State:     state,
		Timestamp: time.Now(),
		Metadata:  metadata,
	}
	
	b.emitEvent(event)
}

// GetValidationSummary returns a summary of current validation states across all fields
func (b *ValidationEventBus) GetValidationSummary() map[string]ValidationState {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	summary := make(map[string]ValidationState)
	
	// Get the latest state for each field from event history
	fieldStates := make(map[string]ValidationEvent)
	
	for _, event := range b.eventHistory {
		// Keep only the most recent event per field
		if existing, exists := fieldStates[event.FieldName]; !exists || event.Timestamp.After(existing.Timestamp) {
			fieldStates[event.FieldName] = event
		}
	}
	
	// Extract states
	for fieldName, event := range fieldStates {
		summary[fieldName] = event.State
	}
	
	return summary
}