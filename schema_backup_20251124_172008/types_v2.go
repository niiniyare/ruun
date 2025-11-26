package schema


// Binding represents Alpine.js data binding configuration
type Binding struct {
	// Core Alpine directives
	Data        string            `json:"x_data,omitempty"`
	Init        string            `json:"x_init,omitempty"`
	Show        string            `json:"x_show,omitempty"`
	Text        string            `json:"x_text,omitempty"`
	HTML        string            `json:"x_html,omitempty"`
	Model       string            `json:"x_model,omitempty"`
	ModelOpts   []string          `json:"x_model_options,omitempty"`
	
	// Conditional rendering
	If          string            `json:"x_if,omitempty"`
	For         string            `json:"x_for,omitempty"`
	ForKey      string            `json:"x_for_key,omitempty"`
	
	// Attribute binding
	Bind        map[string]string `json:"x_bind,omitempty"`
	
	// Event handlers
	On          map[string]string `json:"x_on,omitempty"`
	
	// Transitions
	Transition  TransitionConfig  `json:"x_transition,omitempty"`
	
	// Advanced features
	Effect      string            `json:"x_effect,omitempty"`
	Ignore      bool              `json:"x_ignore,omitempty"`
	Cloak       bool              `json:"x_cloak,omitempty"`
	Teleport    string            `json:"x_teleport,omitempty"`
	Intersect   string            `json:"x_intersect,omitempty"`
	Persist     []string          `json:"x_persist,omitempty"`
	
	// Component communication
	Ref         string            `json:"x_ref,omitempty"`
	Refs        map[string]string `json:"x_refs,omitempty"`
	Id          string            `json:"x_id,omitempty"`
}

// TransitionConfig represents Alpine.js transition configuration
type TransitionConfig struct {
	Enter       string `json:"enter,omitempty"`
	EnterStart  string `json:"enter_start,omitempty"`
	EnterEnd    string `json:"enter_end,omitempty"`
	Leave       string `json:"leave,omitempty"`
	LeaveStart  string `json:"leave_start,omitempty"`
	LeaveEnd    string `json:"leave_end,omitempty"`
	Duration    string `json:"duration,omitempty"`
}

// DOMEvents represents unified DOM event handlers
type DOMEvents struct {
	// Standard DOM events
	Click       string `json:"on_click,omitempty"`
	DblClick    string `json:"on_dblclick,omitempty"`
	MouseDown   string `json:"on_mousedown,omitempty"`
	MouseUp     string `json:"on_mouseup,omitempty"`
	MouseEnter  string `json:"on_mouseenter,omitempty"`
	MouseLeave  string `json:"on_mouseleave,omitempty"`
	MouseOver   string `json:"on_mouseover,omitempty"`
	MouseOut    string `json:"on_mouseout,omitempty"`
	MouseMove   string `json:"on_mousemove,omitempty"`
	
	// Keyboard events
	KeyDown     string `json:"on_keydown,omitempty"`
	KeyUp       string `json:"on_keyup,omitempty"`
	KeyPress    string `json:"on_keypress,omitempty"`
	
	// Form events
	Submit      string `json:"on_submit,omitempty"`
	Change      string `json:"on_change,omitempty"`
	Input       string `json:"on_input,omitempty"`
	Focus       string `json:"on_focus,omitempty"`
	Blur        string `json:"on_blur,omitempty"`
	Select      string `json:"on_select,omitempty"`
	
	// Touch events
	TouchStart  string `json:"on_touchstart,omitempty"`
	TouchEnd    string `json:"on_touchend,omitempty"`
	TouchMove   string `json:"on_touchmove,omitempty"`
	TouchCancel string `json:"on_touchcancel,omitempty"`
	
	// Drag events
	DragStart   string `json:"on_dragstart,omitempty"`
	DragEnd     string `json:"on_dragend,omitempty"`
	DragEnter   string `json:"on_dragenter,omitempty"`
	DragLeave   string `json:"on_dragleave,omitempty"`
	DragOver    string `json:"on_dragover,omitempty"`
	Drop        string `json:"on_drop,omitempty"`
	
	// Media events
	Play        string `json:"on_play,omitempty"`
	Pause       string `json:"on_pause,omitempty"`
	Ended       string `json:"on_ended,omitempty"`
	
	// Window events
	Load        string `json:"on_load,omitempty"`
	Unload      string `json:"on_unload,omitempty"`
	Resize      string `json:"on_resize,omitempty"`
	Scroll      string `json:"on_scroll,omitempty"`
	
	// Custom events
	Custom      map[string]string `json:"custom_events,omitempty"`
}

// Conditional represents unified conditional rendering
type Conditional struct {
	// Simple conditions
	If          string `json:"if,omitempty"`
	Unless      string `json:"unless,omitempty"`
	
	// Multiple conditions
	Switch      string            `json:"switch,omitempty"`
	Cases       []ConditionalCase `json:"cases,omitempty"`
	Default     interface{}       `json:"default,omitempty"`
	
	// Visibility conditions
	Show        string `json:"show,omitempty"`
	Hide        string `json:"hide,omitempty"`
	
	// State-based conditions
	Loading     bool   `json:"loading,omitempty"`
	Error       bool   `json:"error,omitempty"`
	Success     bool   `json:"success,omitempty"`
	
	// Data conditions
	Empty       string `json:"empty,omitempty"`
	NotEmpty    string `json:"not_empty,omitempty"`
	Equals      string `json:"equals,omitempty"`
	NotEquals   string `json:"not_equals,omitempty"`
	
	// Numeric conditions
	GreaterThan string `json:"greater_than,omitempty"`
	LessThan    string `json:"less_than,omitempty"`
	Between     []int  `json:"between,omitempty"`
	
	// Collection conditions
	In          []string `json:"in,omitempty"`
	NotIn       []string `json:"not_in,omitempty"`
	Contains    string   `json:"contains,omitempty"`
	
	// Permission conditions
	Can         string `json:"can,omitempty"`
	Cannot      string `json:"cannot,omitempty"`
	Role        string `json:"role,omitempty"`
}

// ConditionalCase represents a single case in a switch statement
type ConditionalCase struct {
	When  string      `json:"when"`
	Then  interface{} `json:"then"`
}



// GridConfig represents CSS Grid configuration
type GridConfig struct {
	Columns     string            `json:"columns,omitempty"`
	Rows        string            `json:"rows,omitempty"`
	Gap         string            `json:"gap,omitempty"`
	ColumnGap   string            `json:"column_gap,omitempty"`
	RowGap      string            `json:"row_gap,omitempty"`
	AutoFlow    string            `json:"auto_flow,omitempty"`
	AutoColumns string            `json:"auto_columns,omitempty"`
	AutoRows    string            `json:"auto_rows,omitempty"`
	Template    string            `json:"template,omitempty"`
	Areas       []string          `json:"areas,omitempty"`
}

// FlexConfig represents Flexbox configuration
type FlexConfig struct {
	Direction   string            `json:"direction,omitempty"`
	Wrap        string            `json:"wrap,omitempty"`
	AlignItems  string            `json:"align_items,omitempty"`
	AlignContent string           `json:"align_content,omitempty"`
	JustifyContent string         `json:"justify_content,omitempty"`
	Gap         string            `json:"gap,omitempty"`
	Grow        int               `json:"grow,omitempty"`
	Shrink      int               `json:"shrink,omitempty"`
	Basis       string            `json:"basis,omitempty"`
}

// StackConfig represents Stack layout configuration
type StackConfig struct {
	Direction   string            `json:"direction,omitempty"`
	Spacing     string            `json:"spacing,omitempty"`
	Align       string            `json:"align,omitempty"`
	Justify     string            `json:"justify,omitempty"`
	Divider     bool              `json:"divider,omitempty"`
}

// PositionConfig represents positioning configuration
type PositionConfig struct {
	Type        string            `json:"type,omitempty"`
	Top         string            `json:"top,omitempty"`
	Right       string            `json:"right,omitempty"`
	Bottom      string            `json:"bottom,omitempty"`
	Left        string            `json:"left,omitempty"`
	ZIndex      int               `json:"z_index,omitempty"`
}