package schema

import "time"

// Event represents a schema lifecycle or interaction event.
// This is the SINGLE definition - remove duplicates from other files.
type Event struct {
	Type      string         `json:"type"`
	Source    string         `json:"source"`
	Timestamp time.Time      `json:"timestamp"`
	Data      map[string]any `json:"data,omitempty"`
}

func NewEvent(eventType, source string) Event {
	return Event{
		Type:      eventType,
		Source:    source,
		Timestamp: time.Now(),
		Data:      make(map[string]any),
	}
}

func (e Event) WithData(key string, value any) Event {
	if e.Data == nil {
		e.Data = make(map[string]any)
	}
	e.Data[key] = value
	return e
}

// Events configures event handlers for schemas/fields/actions.
type Events struct {
	// Lifecycle
	OnInit    string `json:"onInit,omitempty"`
	OnMount   string `json:"onMount,omitempty"`
	OnUnmount string `json:"onUnmount,omitempty"`
	OnDestroy string `json:"onDestroy,omitempty"`

	// Form submission
	OnSubmit        string `json:"onSubmit,omitempty"`
	BeforeSubmit    string `json:"beforeSubmit,omitempty"`
	AfterSubmit     string `json:"afterSubmit,omitempty"`
	OnSubmitSuccess string `json:"onSubmitSuccess,omitempty"`
	OnSubmitError   string `json:"onSubmitError,omitempty"`

	// Field interactions
	OnChange   string `json:"onChange,omitempty"`
	OnInput    string `json:"onInput,omitempty"`
	OnFocus    string `json:"onFocus,omitempty"`
	OnBlur     string `json:"onBlur,omitempty"`
	OnValidate string `json:"onValidate,omitempty"`

	// Data operations
	OnLoad    string `json:"onLoad,omitempty"`
	OnSuccess string `json:"onSuccess,omitempty"`
	OnError   string `json:"onError,omitempty"`

	// State changes
	OnReset    string `json:"onReset,omitempty"`
	OnDirty    string `json:"onDirty,omitempty"`
	OnPristine string `json:"onPristine,omitempty"`

	// Custom handlers
	Custom map[string]string `json:"custom,omitempty"`
}

func (e *Events) IsEmpty() bool {
	if e == nil {
		return true
	}
	return e.OnInit == "" && e.OnMount == "" && e.OnSubmit == "" &&
		e.OnChange == "" && e.OnLoad == "" && len(e.Custom) == 0
}

// BehaviorMetadata describes a behavior capability.
// This is the SINGLE definition - remove duplicates.
type BehaviorMetadata struct {
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	Version     string               `json:"version,omitempty"`
	Inputs      map[string]FieldType `json:"inputs,omitempty"`
	Outputs     map[string]FieldType `json:"outputs,omitempty"`
	Tags        []string             `json:"tags,omitempty"`
}

// Common event type constants
const (
	EventTypeFieldChange  = "field:change"
	EventTypeFieldFocus   = "field:focus"
	EventTypeFieldBlur    = "field:blur"
	EventTypeFormSubmit   = "form:submit"
	EventTypeFormReset    = "form:reset"
	EventTypeFormValidate = "form:validate"
	EventTypeDataLoad     = "data:load"
	EventTypeDataSave     = "data:save"
	EventTypeStateChange  = "state:change"
)