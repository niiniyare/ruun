package schema

// Behavior defines technology-agnostic request/interaction behavior.
// This replaces HTMX, ActionHTMX, FieldHTMX with a unified approach
// that can adapt to any frontend framework (HTMX, Alpine, Datastar, etc.)
type Behavior struct {
	// Request configuration
	Method  string            `json:"method,omitempty"`  // HTTP method: GET, POST, PUT, DELETE, PATCH
	URL     string            `json:"url,omitempty"`     // Target URL for the request
	Target  string            `json:"target,omitempty"`  // CSS selector for response target
	Swap    SwapStrategy      `json:"swap,omitempty"`    // How to swap the response content
	Trigger *Trigger          `json:"trigger,omitempty"` // Event that triggers the behavior
	Headers map[string]string `json:"headers,omitempty"` // Additional request headers
	Params  map[string]string `json:"params,omitempty"`  // Query/body parameters

	// Loading states
	Indicator string `json:"indicator,omitempty"` // CSS selector for loading indicator
	Disabled  string `json:"disabled,omitempty"`  // CSS selector for elements to disable

	// Timing control
	Debounce int `json:"debounce,omitempty"` // Debounce delay in milliseconds
	Throttle int `json:"throttle,omitempty"` // Throttle interval in milliseconds
	Timeout  int `json:"timeout,omitempty"`  // Request timeout in milliseconds

	// User confirmation
	Confirm string `json:"confirm,omitempty"` // Confirmation message before action

	// Browser history
	PushURL   bool   `json:"pushUrl,omitempty"`   // Push URL to browser history
	ReplaceURL bool   `json:"replaceUrl,omitempty"` // Replace current URL in history
	Select    string `json:"select,omitempty"`    // CSS selector to pick from response

	// Retry configuration
	Retry      int `json:"retry,omitempty"`      // Number of retry attempts
	RetryDelay int `json:"retryDelay,omitempty"` // Delay between retries in ms
}

// Trigger defines when a behavior should be activated
type Trigger struct {
	Event    string `json:"event,omitempty"`    // DOM event: click, change, submit, input, etc.
	Target   string `json:"target,omitempty"`   // CSS selector for event target
	Filter   string `json:"filter,omitempty"`   // Condition expression to filter events
	Delay    int    `json:"delay,omitempty"`    // Delay before triggering in ms
	Once     bool   `json:"once,omitempty"`     // Trigger only once
	Changed  bool   `json:"changed,omitempty"`  // Only if value has changed
	Throttle int    `json:"throttle,omitempty"` // Throttle trigger rate
	From     string `json:"from,omitempty"`     // Listen for events from another element
}

// SwapStrategy defines how content should be swapped
type SwapStrategy string

const (
	SwapInnerHTML   SwapStrategy = "innerHTML"   // Replace inner content
	SwapOuterHTML   SwapStrategy = "outerHTML"   // Replace entire element
	SwapBeforeBegin SwapStrategy = "beforebegin" // Insert before element
	SwapAfterBegin  SwapStrategy = "afterbegin"  // Insert at start of element
	SwapBeforeEnd   SwapStrategy = "beforeend"   // Insert at end of element
	SwapAfterEnd    SwapStrategy = "afterend"    // Insert after element
	SwapDelete      SwapStrategy = "delete"      // Delete the element
	SwapNone        SwapStrategy = "none"        // No swap, just trigger events
)

// String returns the string representation of SwapStrategy
func (s SwapStrategy) String() string {
	return string(s)
}

// IsValid checks if the swap strategy is valid
func (s SwapStrategy) IsValid() bool {
	switch s {
	case SwapInnerHTML, SwapOuterHTML, SwapBeforeBegin, SwapAfterBegin,
		SwapBeforeEnd, SwapAfterEnd, SwapDelete, SwapNone:
		return true
	default:
		return s == ""
	}
}