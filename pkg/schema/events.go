package schema

import (
	"context"
	"time"
)

// EventType represents the type of form event
type EventType string

const (
	EventChange EventType = "change" // Field value changed
	EventBlur   EventType = "blur"   // Field lost focus
	EventFocus  EventType = "focus"  // Field gained focus
	EventSubmit EventType = "submit" // Form submitted
	EventReset  EventType = "reset"  // Form reset
	EventInit   EventType = "init"   // Form initialized
)

// EventCallback represents a function that handles form events
type EventCallback func(ctx context.Context, event *Event) error

// Event represents a form event
type Event struct {
	Type      EventType      `json:"type"`                // Event type
	FieldName string         `json:"fieldName,omitempty"` // Field that triggered the event
	Value     any            `json:"value,omitempty"`     // Field value (for change events)
	OldValue  any            `json:"oldValue,omitempty"`  // Previous value (for change events)
	Timestamp time.Time      `json:"timestamp"`           // When the event occurred
	Data      map[string]any `json:"data,omitempty"`      // Additional event data
	Metadata  map[string]any `json:"metadata,omitempty"`  // Event metadata
}

// EventHandlerConfig configures event handler behavior
type EventHandlerConfig struct {
	ValidationTiming ValidationTiming `json:"validationTiming"` // When to validate
	DebounceDelay    time.Duration    `json:"debounceDelay"`    // Debounce delay for events
	EnableTracking   bool             `json:"enableTracking"`   // Enable event tracking
	MaxEvents        int              `json:"maxEvents"`        // Maximum events to track
}

// DebouncedConfig configures debounced event handling
type DebouncedConfig struct {
	ChangeDelay time.Duration `json:"changeDelay"` // Delay for change events
	BlurDelay   time.Duration `json:"blurDelay"`   // Delay for blur events
	FocusDelay  time.Duration `json:"focusDelay"`  // Delay for focus events
}

// Default configurations

// DefaultEventHandlerConfig returns default event handler configuration
func DefaultEventHandlerConfig() *EventHandlerConfig {
	return &EventHandlerConfig{
		ValidationTiming: ValidateOnBlur,
		DebounceDelay:    300 * time.Millisecond,
		EnableTracking:   false,
		MaxEvents:        1000,
	}
}

// DefaultDebouncedConfig returns default debounced configuration
func DefaultDebouncedConfig() *DebouncedConfig {
	return &DebouncedConfig{
		ChangeDelay: 300 * time.Millisecond,
		BlurDelay:   100 * time.Millisecond,
		FocusDelay:  50 * time.Millisecond,
	}
}
