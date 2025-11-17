package schema

import (
	"context"
	"time"
)

// ValidationCallback is called when async validation completes
type ValidationCallback func(fieldName string, errors []string)

// RuntimeRenderer defines the interface for rendering schema components
// UI frameworks must implement this interface to render forms
type RuntimeRenderer interface {
	// Core rendering methods
	RenderField(ctx context.Context, field *Field, value any, errors []string, touched, dirty bool) (string, error)
	RenderForm(ctx context.Context, schema *Schema, state map[string]any, errors map[string][]string) (string, error)
	RenderAction(ctx context.Context, action *Action, enabled bool) (string, error)

	// Layout rendering
	RenderLayout(ctx context.Context, layout *Layout, fields []*Field, state map[string]any) (string, error)
	RenderSection(ctx context.Context, section *Section, fields []*Field, state map[string]any) (string, error)
	RenderTab(ctx context.Context, tab *Tab, fields []*Field, state map[string]any) (string, error)
	RenderStep(ctx context.Context, step *Step, fields []*Field, state map[string]any) (string, error)
	RenderGroup(ctx context.Context, group *Group, fields []*Field, state map[string]any) (string, error)

	// Error rendering
	RenderErrors(ctx context.Context, errors map[string][]string) (string, error)
}

// RuntimeValidator defines the interface for runtime validation
// Validates data in real-time during user interaction
type RuntimeValidator interface {
	// Field validation
	ValidateField(ctx context.Context, field *Field, value any, allData map[string]any) []string
	ValidateAllFields(ctx context.Context, schema *Schema, data map[string]any) map[string][]string

	// Action validation - validates if action can be performed
	ValidateAction(ctx context.Context, action *Action, formState map[string]any) error

	// Workflow validation - validates workflow state transitions
	ValidateWorkflow(ctx context.Context, workflow *Workflow, currentStage string, formState map[string]any) error

	// Business rules validation - validates business logic
	ValidateBusinessRules(ctx context.Context, rules []ValidationRule, allValues map[string]any) map[string][]string

	// Async validation for expensive operations (API calls, database lookups)
	ValidateFieldAsync(ctx context.Context, field *Field, value any, callback ValidationCallback) error
}

// FieldConditionalResult contains the result of conditional evaluation for a field
type FieldConditionalResult struct {
	FieldName string // Name of the field
	Visible   bool   // Whether field should be visible
	Editable  bool   // Whether field should be editable
	Required  bool   // Whether field should be required
	Reason    string // Reason for the state (for debugging)
}

// RuntimeLayoutConditionalResult contains the result of conditional evaluation for layouts
type RuntimeLayoutConditionalResult struct {
	Changed bool                       // Whether any conditions changed
	Updates []*FieldConditionalResult // Field updates to apply
}

// ConditionalEvaluationResult contains the result of evaluating all conditions
type ConditionalEvaluationResult struct {
	Changed bool                       // Whether any conditions changed state
	Updates []*FieldConditionalResult // List of field updates to apply
}

// RuntimeConditionalEngine defines the interface for evaluating conditional logic
// Determines field visibility, editability, and requirements based on form state
type RuntimeConditionalEngine interface {
	// Field conditional evaluation
	EvaluateFieldVisibility(ctx context.Context, field *Field, data map[string]any) (bool, error)
	EvaluateFieldRequired(ctx context.Context, field *Field, data map[string]any) (bool, error)
	EvaluateFieldEditable(ctx context.Context, field *Field, data map[string]any) (bool, error)

	// Action conditional evaluation - determines if actions should be available
	EvaluateActionConditions(ctx context.Context, action *Action, allValues map[string]any) (bool, error)

	// Layout conditional evaluation - evaluates conditions for layout components
	EvaluateLayoutConditions(ctx context.Context, layout *Layout, allValues map[string]any) (*RuntimeLayoutConditionalResult, error)

	// Workflow conditional evaluation - evaluates workflow transitions
	EvaluateWorkflowConditions(ctx context.Context, workflow *Workflow, allValues map[string]any) (*RuntimeLayoutConditionalResult, error)

	// Batch evaluations for performance - evaluates all field conditions at once
	EvaluateAllConditions(ctx context.Context, schema *Schema, data map[string]any) (*ConditionalEvaluationResult, error)
}

// StateSnapshot represents a point-in-time snapshot of form state
type StateSnapshot struct {
	Values      map[string]any      `json:"values"`      // Current field values
	Initial     map[string]any      `json:"initial"`     // Initial values for dirty tracking
	Touched     map[string]bool     `json:"touched"`     // Fields user has interacted with
	Dirty       map[string]bool     `json:"dirty"`       // Fields that changed from initial
	Errors      map[string][]string `json:"errors"`      // Validation errors per field
	Timestamp   time.Time           `json:"timestamp"`   // When snapshot was taken
	Initialized time.Time           `json:"initialized"` // When state was initialized
}

// StateStats provides statistics about the current state
type StateStats struct {
	FieldCount        int       `json:"fieldCount"`        // Total number of fields
	TouchedCount      int       `json:"touchedCount"`      // Number of touched fields
	DirtyCount        int       `json:"dirtyCount"`        // Number of dirty fields
	ErrorCount        int       `json:"errorCount"`        // Total number of validation errors
	IsValid           bool      `json:"isValid"`           // Whether form is currently valid
	InitializedAt     time.Time `json:"initializedAt"`     // When state was initialized
	LastUpdated       time.Time `json:"lastUpdated"`       // Last state modification time
	HasUnsavedChanges bool      `json:"hasUnsavedChanges"` // Whether there are unsaved changes
}

// RuntimeStateManager defines the interface for managing form state
// Tracks field values, touched state, dirty state, and validation errors
type RuntimeStateManager interface {
	// State initialization
	Initialize(schema *Schema, initialData map[string]any) error

	// Value management
	GetValue(fieldName string) (any, bool)
	SetValue(fieldName string, value any) error
	GetAll() map[string]any

	// State tracking
	IsTouched(fieldName string) bool
	Touch(fieldName string)
	IsDirty(fieldName string) bool
	IsAnyDirty() bool

	// Error management
	GetErrors(fieldName string) []string
	SetErrors(fieldName string, errors []string)
	GetAllErrors() map[string][]string
	IsValid() bool

	// State operations
	Reset()
	CreateSnapshot() any // Returns StateSnapshot but typed as any to avoid import cycles
	RestoreSnapshot(snapshot any)
}

// EventStats provides statistics about runtime events
type EventStats struct {
	TotalEvents     int                `json:"totalEvents"`     // Total number of events processed
	EventsByType    map[string]int     `json:"eventsByType"`    // Count by event type
	EventsByField   map[string]int     `json:"eventsByField"`   // Count by field name
	LastEventType   string             `json:"lastEventType"`   // Type of last event
	LastEventField  string             `json:"lastEventField"`  // Field of last event
	LastEventTime   time.Time          `json:"lastEventTime"`   // Time of last event
	ProcessingTimes map[string][]int64 `json:"processingTimes"` // Processing times by event type (microseconds)
}

// RuntimeStats provides comprehensive statistics about runtime performance
type RuntimeStats struct {
	// Field statistics
	FieldCount    int `json:"fieldCount"`    // Total number of fields in schema
	VisibleFields int `json:"visibleFields"` // Number of currently visible fields
	TouchedFields int `json:"touchedFields"` // Number of fields user has touched
	DirtyFields   int `json:"dirtyFields"`   // Number of fields that have changed
	ErrorCount    int `json:"errorCount"`    // Total number of validation errors

	// State information
	IsValid       bool      `json:"isValid"`       // Whether form is currently valid
	InitializedAt time.Time `json:"initializedAt"` // When runtime was initialized
	LastActivity  time.Time `json:"lastActivity"`  // Time of last user activity
	Uptime        time.Duration `json:"uptime"`    // How long runtime has been active

	// Performance metrics
	EventStats *EventStats `json:"eventStats,omitempty"` // Event processing statistics
}

// RuntimeConfig defines configuration options for the runtime
type RuntimeConfig struct {
	// Feature toggles
	EnableConditionals  bool `json:"enableConditionals"`  // Enable conditional field logic
	EnableValidation    bool `json:"enableValidation"`    // Enable real-time validation
	EnableEventTracking bool `json:"enableEventTracking"` // Track events for analytics

	// Validation configuration
	ValidationTiming ValidationTiming `json:"validationTiming"` // When to trigger validation
	EnableDebounce   bool             `json:"enableDebounce"`   // Debounce validation calls
	DebounceDelay    time.Duration    `json:"debounceDelay"`    // Delay before validation

	// State management
	MaxStateSnapshots int `json:"maxStateSnapshots"` // Maximum number of state snapshots to keep
}

// ValidationTiming defines when validation should occur
type ValidationTiming string

const (
	ValidateOnChange ValidationTiming = "change" // Validate immediately on field change
	ValidateOnBlur   ValidationTiming = "blur"   // Validate when field loses focus
	ValidateOnSubmit ValidationTiming = "submit" // Validate only on form submission
	ValidateNever    ValidationTiming = "never"  // No automatic validation
)

// DefaultRuntimeConfig returns a default runtime configuration
func DefaultRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		EnableConditionals:  true,
		EnableValidation:    true,
		EnableEventTracking: false,
		ValidationTiming:    ValidateOnBlur,
		EnableDebounce:      true,
		DebounceDelay:       300 * time.Millisecond,
		MaxStateSnapshots:   10,
	}
}