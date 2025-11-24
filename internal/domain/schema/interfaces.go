package schema

import (
	"context"
	"io"
)

// Renderer defines the contract for rendering schemas into output formats.
// Implementations transform schema definitions into specific formats like HTML,
// JSON, or custom templating languages.
type Renderer interface {
	// Render transforms a schema into the target output format.
	// Returns an error if the schema cannot be rendered.
	Render(ctx context.Context, schema Schema, writer io.Writer) error

	// RenderField renders a single field with its current value.
	// Useful for partial updates and field-level rendering.
	RenderField(ctx context.Context, field Field, value interface{}, writer io.Writer) error

	// SupportsFormat checks if this renderer can handle the specified format.
	SupportsFormat(format string) bool
}

// StateManager handles the persistence and retrieval of schema states.
// This includes form data, field values, and UI state across sessions.
type StateManager interface {
	// SaveState persists the current state of a schema instance.
	// The state includes field values, validation results, and UI state.
	SaveState(ctx context.Context, schemaID string, state map[string]interface{}) error

	// LoadState retrieves a previously saved state for a schema instance.
	// Returns nil if no state exists for the given schemaID.
	LoadState(ctx context.Context, schemaID string) (map[string]interface{}, error)

	// DeleteState removes the saved state for a schema instance.
	DeleteState(ctx context.Context, schemaID string) error

	// ListStates returns all saved states matching the given criteria.
	// Useful for managing multiple form instances or sessions.
	ListStates(ctx context.Context, criteria map[string]interface{}) ([]string, error)
}

// EventDispatcher manages event routing and handling for schema interactions.
// Events include field changes, validations, submissions, and custom actions.
type EventDispatcher interface {
	// Dispatch sends an event to all registered handlers.
	// Returns an error if any handler fails critically.
	Dispatch(ctx context.Context, event Event) error

	// Subscribe registers a handler for specific event types.
	// The handler will be called whenever matching events are dispatched.
	Subscribe(eventType string, handler EventHandler) error

	// Unsubscribe removes a handler from the event dispatcher.
	Unsubscribe(eventType string, handler EventHandler) error

	// GetHandlers returns all handlers registered for a specific event type.
	GetHandlers(eventType string) []EventHandler
}

// BehaviorAdapter defines custom behavior implementations for schemas.
// This allows extending schema functionality with domain-specific logic.
type BehaviorAdapter interface {
	// Execute runs the behavior with the provided context and arguments.
	// Returns the behavior result and any error encountered.
	Execute(ctx context.Context, args map[string]interface{}) (interface{}, error)

	// Validate checks if the behavior can be executed with the given arguments.
	// Returns validation errors without executing the behavior.
	Validate(args map[string]interface{}) error

	// GetMetadata returns information about the behavior's requirements and capabilities.
	// Useful for UI generation and documentation.
	GetMetadata() BehaviorMetadata
}

// Validator defines the contract for field and schema validation.
// Implementations provide various validation strategies and rules.
type Validator interface {
	// ValidateField checks a single field value against its validation rules.
	// Returns a ValidationResult containing any errors or warnings.
	ValidateField(ctx context.Context, field Field, value interface{}) ValidationResult

	// ValidateSchema validates an entire schema instance.
	// This includes field validations and cross-field dependencies.
	ValidateSchema(ctx context.Context, schema Schema, values map[string]interface{}) ValidationResult

	// AddRule registers a custom validation rule.
	// Rules can be reused across multiple fields and schemas.
	AddRule(name string, rule ValidationRule) error

	// GetRule retrieves a registered validation rule by name.
	GetRule(name string) (ValidationRule, bool)
}

// Event represents a schema-related event that can be dispatched and handled.
type Event struct {
	Type      string                 // Event type identifier
	Source    string                 // Component or field that triggered the event
	Timestamp int64                  // Unix timestamp when event occurred
	Data      map[string]interface{} // Event-specific data payload
}

// EventHandler processes events dispatched by the EventDispatcher.
type EventHandler func(ctx context.Context, event Event) error

// BehaviorMetadata describes a behavior's capabilities and requirements.
type BehaviorMetadata struct {
	Name        string                 // Unique behavior identifier
	Description string                 // Human-readable description
	Version     string                 // Behavior version for compatibility
	Inputs      map[string]FieldType   // Required input parameters
	Outputs     map[string]FieldType   // Expected output values
	Tags        []string               // Categorization tags
}

// ValidationResult contains the outcome of a validation operation.
type ValidationResult struct {
	Valid    bool              // Whether validation passed
	Errors   []ValidationError // List of validation errors
	Warnings []string          // Non-critical validation warnings
}

// ValidationError represents a specific validation failure.
type ValidationError struct {
	Field   string // Field that failed validation
	Rule    string // Validation rule that failed
	Message string // Human-readable error message
	Code    string // Machine-readable error code
}

// ValidationRule defines a reusable validation check.
type ValidationRule func(value interface{}, params map[string]interface{}) error