package schema

// Behavior defines technology-agnostic request/interaction behavior.
// Replaces HTMX, ActionHTMX, FieldHTMX
type Behavior struct {
	Method    string            `json:"method,omitempty"`
	URL       string            `json:"url,omitempty"`
	Target    string            `json:"target,omitempty"`
	Swap      SwapStrategy      `json:"swap,omitempty"`
	Trigger   *Trigger          `json:"trigger,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Params    map[string]string `json:"params,omitempty"`
	Indicator string            `json:"indicator,omitempty"`
	Disabled  string            `json:"disabled,omitempty"`
	Debounce  int               `json:"debounce,omitempty"`
	Throttle  int               `json:"throttle,omitempty"`
	Timeout   int               `json:"timeout,omitempty"`
	Confirm   string            `json:"confirm,omitempty"`
	PushURL   bool              `json:"pushUrl,omitempty"`
	Select    string            `json:"select,omitempty"`
}

type Trigger struct {
	Event   string `json:"event,omitempty"`
	Target  string `json:"target,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Delay   int    `json:"delay,omitempty"`
	Once    bool   `json:"once,omitempty"`
	Changed bool   `json:"changed,omitempty"`
	From    string `json:"from,omitempty"`
}

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