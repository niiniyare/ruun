package schema

import "strings"

// SwapStrategy represents HTMX swap strategies for DOM manipulation.
// This is the SINGLE definition - remove duplicates from other files.
type SwapStrategy string

const (
	SwapInnerHTML   SwapStrategy = "innerHTML"
	SwapOuterHTML   SwapStrategy = "outerHTML"
	SwapBeforeBegin SwapStrategy = "beforebegin"
	SwapAfterBegin  SwapStrategy = "afterbegin"
	SwapBeforeEnd   SwapStrategy = "beforeend"
	SwapAfterEnd    SwapStrategy = "afterend"
	SwapDelete      SwapStrategy = "delete"
	SwapNone        SwapStrategy = "none"
)

// String returns the string representation of SwapStrategy
func (ss SwapStrategy) String() string {
	return string(ss)
}

// IsValid returns true if the SwapStrategy is valid
func (ss SwapStrategy) IsValid() bool {
	switch ss {
	case SwapInnerHTML, SwapOuterHTML, SwapBeforeBegin, SwapAfterBegin,
		SwapBeforeEnd, SwapAfterEnd, SwapDelete, SwapNone:
		return true
	default:
		return false
	}
}

// Trigger represents an event trigger configuration.
// This is the SINGLE definition - remove duplicates from other files.
type Trigger struct {
	Event   string `json:"event,omitempty"`   // Event to listen for (click, change, etc.)
	Target  string `json:"target,omitempty"`  // Target selector for the event
	Filter  string `json:"filter,omitempty"`  // CSS selector to filter events
	Delay   int    `json:"delay,omitempty"`   // Delay before triggering (ms)
	Once    bool   `json:"once,omitempty"`    // Trigger only once
	Changed bool   `json:"changed,omitempty"` // Only trigger if value changed
	From    string `json:"from,omitempty"`    // Previous value condition
}

// IsEmpty returns true if the trigger has no configuration
func (t *Trigger) IsEmpty() bool {
	return t == nil || (t.Event == "" && t.Target == "" && t.Filter == "" &&
		t.Delay == 0 && !t.Once && !t.Changed && t.From == "")
}

// Behavior defines technology-agnostic request/interaction behavior.
// Replaces: HTMX, ActionHTMX, FieldHTMX types.
// This is the SINGLE definition for all behavior-related functionality.
type Behavior struct {
	// HTTP Configuration
	Method  string            `json:"method,omitempty"`  // HTTP method (GET, POST, PUT, DELETE, PATCH)
	URL     string            `json:"url,omitempty"`     // Request URL or endpoint
	Headers map[string]string `json:"headers,omitempty"` // Additional HTTP headers
	Params  map[string]string `json:"params,omitempty"`  // URL parameters or form data

	// DOM Configuration
	Target    string       `json:"target,omitempty"`    // Target element selector for response
	Swap      SwapStrategy `json:"swap,omitempty"`      // How to swap content into target
	Select    string       `json:"select,omitempty"`    // CSS selector to select from response
	Indicator string       `json:"indicator,omitempty"` // Loading indicator selector

	// Event Configuration
	Trigger *Trigger `json:"trigger,omitempty"` // Event trigger configuration

	// Timing & Performance
	Debounce int `json:"debounce,omitempty"` // Debounce delay in milliseconds
	Throttle int `json:"throttle,omitempty"` // Throttle delay in milliseconds
	Timeout  int `json:"timeout,omitempty"`  // Request timeout in milliseconds

	// State & Control
	Disabled string `json:"disabled,omitempty"` // Element to disable during request
	Confirm  string `json:"confirm,omitempty"`  // Confirmation message before request
	PushURL  bool   `json:"pushUrl,omitempty"`  // Whether to push URL to browser history

	// HTMX-specific (for backward compatibility)
	Enabled  bool   `json:"enabled,omitempty"`  // Whether behavior is enabled
	Vals     string `json:"vals,omitempty"`     // Additional values to include
	Boost    bool   `json:"boost,omitempty"`    // Use HTMX boost mode
	Sync     string `json:"sync,omitempty"`     // Synchronization strategy
	Validate bool   `json:"validate,omitempty"` // Whether to validate before request
	Retry    int    `json:"retry,omitempty"`    // Number of retry attempts
}

// IsEmpty returns true if the behavior has no configuration
func (b *Behavior) IsEmpty() bool {
	return b == nil || (b.Method == "" && b.URL == "" && len(b.Headers) == 0 &&
		len(b.Params) == 0 && b.Target == "" && b.Swap == "" && b.Select == "" &&
		b.Indicator == "" && b.Trigger.IsEmpty() && b.Debounce == 0 &&
		b.Throttle == 0 && b.Timeout == 0 && b.Disabled == "" &&
		b.Confirm == "" && !b.PushURL && !b.Enabled)
}

// GetHTTPMethod returns the HTTP method, defaulting to GET
func (b *Behavior) GetHTTPMethod() string {
	if b == nil || b.Method == "" {
		return "GET"
	}
	return strings.ToUpper(b.Method)
}

// HasTarget returns true if a target element is specified
func (b *Behavior) HasTarget() bool {
	return b != nil && b.Target != ""
}

// HasTrigger returns true if a trigger is configured
func (b *Behavior) HasTrigger() bool {
	return b != nil && !b.Trigger.IsEmpty()
}

// IsDebounced returns true if debouncing is enabled
func (b *Behavior) IsDebounced() bool {
	return b != nil && b.Debounce > 0
}

// IsThrottled returns true if throttling is enabled
func (b *Behavior) IsThrottled() bool {
	return b != nil && b.Throttle > 0
}

// HasTimeout returns true if a timeout is configured
func (b *Behavior) HasTimeout() bool {
	return b != nil && b.Timeout > 0
}

// RequiresConfirmation returns true if confirmation is required
func (b *Behavior) RequiresConfirmation() bool {
	return b != nil && b.Confirm != ""
}

// Clone returns a deep copy of the Behavior
func (b *Behavior) Clone() *Behavior {
	if b == nil {
		return nil
	}

	clone := &Behavior{
		Method:    b.Method,
		URL:       b.URL,
		Target:    b.Target,
		Swap:      b.Swap,
		Select:    b.Select,
		Indicator: b.Indicator,
		Debounce:  b.Debounce,
		Throttle:  b.Throttle,
		Timeout:   b.Timeout,
		Disabled:  b.Disabled,
		Confirm:   b.Confirm,
		PushURL:   b.PushURL,
		Enabled:   b.Enabled,
		Vals:      b.Vals,
		Boost:     b.Boost,
		Sync:      b.Sync,
		Validate:  b.Validate,
		Retry:     b.Retry,
	}

	// Clone headers map
	if b.Headers != nil {
		clone.Headers = make(map[string]string, len(b.Headers))
		for k, v := range b.Headers {
			clone.Headers[k] = v
		}
	}

	// Clone params map
	if b.Params != nil {
		clone.Params = make(map[string]string, len(b.Params))
		for k, v := range b.Params {
			clone.Params[k] = v
		}
	}

	// Clone trigger
	if b.Trigger != nil && !b.Trigger.IsEmpty() {
		clone.Trigger = &Trigger{
			Event:   b.Trigger.Event,
			Target:  b.Trigger.Target,
			Filter:  b.Trigger.Filter,
			Delay:   b.Trigger.Delay,
			Once:    b.Trigger.Once,
			Changed: b.Trigger.Changed,
			From:    b.Trigger.From,
		}
	}

	return clone
}

// Merge combines this behavior with another, with the other taking precedence
func (b *Behavior) Merge(other *Behavior) *Behavior {
	if b == nil {
		return other.Clone()
	}
	if other == nil {
		return b.Clone()
	}

	merged := b.Clone()

	// Override with non-empty values from other
	if other.Method != "" {
		merged.Method = other.Method
	}
	if other.URL != "" {
		merged.URL = other.URL
	}
	if other.Target != "" {
		merged.Target = other.Target
	}
	if other.Swap != "" {
		merged.Swap = other.Swap
	}
	if other.Select != "" {
		merged.Select = other.Select
	}
	if other.Indicator != "" {
		merged.Indicator = other.Indicator
	}
	if other.Debounce > 0 {
		merged.Debounce = other.Debounce
	}
	if other.Throttle > 0 {
		merged.Throttle = other.Throttle
	}
	if other.Timeout > 0 {
		merged.Timeout = other.Timeout
	}
	if other.Disabled != "" {
		merged.Disabled = other.Disabled
	}
	if other.Confirm != "" {
		merged.Confirm = other.Confirm
	}
	if other.PushURL {
		merged.PushURL = other.PushURL
	}
	if other.Enabled {
		merged.Enabled = other.Enabled
	}
	if other.Vals != "" {
		merged.Vals = other.Vals
	}
	if other.Boost {
		merged.Boost = other.Boost
	}
	if other.Sync != "" {
		merged.Sync = other.Sync
	}
	if other.Validate {
		merged.Validate = other.Validate
	}
	if other.Retry > 0 {
		merged.Retry = other.Retry
	}

	// Merge headers
	if other.Headers != nil {
		if merged.Headers == nil {
			merged.Headers = make(map[string]string)
		}
		for k, v := range other.Headers {
			merged.Headers[k] = v
		}
	}

	// Merge params
	if other.Params != nil {
		if merged.Params == nil {
			merged.Params = make(map[string]string)
		}
		for k, v := range other.Params {
			merged.Params[k] = v
		}
	}

	// Merge trigger
	if other.Trigger != nil && !other.Trigger.IsEmpty() {
		merged.Trigger = other.Trigger
	}

	return merged
}

// BehaviorBuilder provides a fluent API for building Behavior configurations
type BehaviorBuilder struct {
	behavior *Behavior
}

// NewBehaviorBuilder creates a new BehaviorBuilder
func NewBehaviorBuilder() *BehaviorBuilder {
	return &BehaviorBuilder{
		behavior: &Behavior{},
	}
}

// Method sets the HTTP method
func (bb *BehaviorBuilder) Method(method string) *BehaviorBuilder {
	bb.behavior.Method = method
	return bb
}

// URL sets the request URL
func (bb *BehaviorBuilder) URL(url string) *BehaviorBuilder {
	bb.behavior.URL = url
	return bb
}

// Target sets the target element selector
func (bb *BehaviorBuilder) Target(target string) *BehaviorBuilder {
	bb.behavior.Target = target
	return bb
}

// Swap sets the swap strategy
func (bb *BehaviorBuilder) Swap(swap SwapStrategy) *BehaviorBuilder {
	bb.behavior.Swap = swap
	return bb
}

// Trigger sets the event trigger
func (bb *BehaviorBuilder) Trigger(event string) *BehaviorBuilder {
	if bb.behavior.Trigger == nil {
		bb.behavior.Trigger = &Trigger{}
	}
	bb.behavior.Trigger.Event = event
	return bb
}

// TriggerWithTarget sets the event trigger with target
func (bb *BehaviorBuilder) TriggerWithTarget(event, target string) *BehaviorBuilder {
	if bb.behavior.Trigger == nil {
		bb.behavior.Trigger = &Trigger{}
	}
	bb.behavior.Trigger.Event = event
	bb.behavior.Trigger.Target = target
	return bb
}

// Debounce sets the debounce delay
func (bb *BehaviorBuilder) Debounce(delay int) *BehaviorBuilder {
	bb.behavior.Debounce = delay
	return bb
}

// Throttle sets the throttle delay
func (bb *BehaviorBuilder) Throttle(delay int) *BehaviorBuilder {
	bb.behavior.Throttle = delay
	return bb
}

// Timeout sets the request timeout
func (bb *BehaviorBuilder) Timeout(timeout int) *BehaviorBuilder {
	bb.behavior.Timeout = timeout
	return bb
}

// Confirm sets the confirmation message
func (bb *BehaviorBuilder) Confirm(message string) *BehaviorBuilder {
	bb.behavior.Confirm = message
	return bb
}

// Header adds an HTTP header
func (bb *BehaviorBuilder) Header(name, value string) *BehaviorBuilder {
	if bb.behavior.Headers == nil {
		bb.behavior.Headers = make(map[string]string)
	}
	bb.behavior.Headers[name] = value
	return bb
}

// Param adds a parameter
func (bb *BehaviorBuilder) Param(name, value string) *BehaviorBuilder {
	if bb.behavior.Params == nil {
		bb.behavior.Params = make(map[string]string)
	}
	bb.behavior.Params[name] = value
	return bb
}

// Build returns the built Behavior
func (bb *BehaviorBuilder) Build() *Behavior {
	return bb.behavior.Clone()
}

// Common behavior constructors for convenience

// HTTPBehavior creates a basic HTTP request behavior
func HTTPBehavior(method, url string) *Behavior {
	return &Behavior{
		Method: method,
		URL:    url,
	}
}

// HTMXBehavior creates an HTMX-style behavior with target and swap
func HTMXBehavior(method, url, target string, swap SwapStrategy) *Behavior {
	return &Behavior{
		Method:  method,
		URL:     url,
		Target:  target,
		Swap:    swap,
		Enabled: true,
	}
}

// FormSubmitBehavior creates a form submission behavior
func FormSubmitBehavior(url string) *Behavior {
	return &Behavior{
		Method:  "POST",
		URL:     url,
		Trigger: &Trigger{Event: "submit"},
		Enabled: true,
	}
}

// ClickBehavior creates a click-triggered behavior
func ClickBehavior(url, target string) *Behavior {
	return &Behavior{
		Method:  "GET",
		URL:     url,
		Target:  target,
		Swap:    SwapInnerHTML,
		Trigger: &Trigger{Event: "click"},
		Enabled: true,
	}
}

// SearchBehavior creates a search input behavior with debouncing
func SearchBehavior(url, target string, debounce int) *Behavior {
	return &Behavior{
		Method:   "GET",
		URL:      url,
		Target:   target,
		Swap:     SwapInnerHTML,
		Trigger:  &Trigger{Event: "keyup", Changed: true},
		Debounce: debounce,
		Enabled:  true,
	}
}
