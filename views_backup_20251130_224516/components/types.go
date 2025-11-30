package components

import (
	"strings"

	"github.com/a-h/templ"
)

// Common type definitions and utilities shared across all component levels

// Size represents a size variant that can be used across different components
type Size string

const (
	SizeXs      Size = "xs"
	SizeSm      Size = "sm"
	SizeDefault Size = "default"
	SizeMd      Size = "md"
	SizeLg      Size = "lg"
	SizeXl      Size = "xl"
)

// Variant represents a visual style variant
type Variant string

const (
	VariantDefault     Variant = "default"
	VariantPrimary     Variant = "primary"
	VariantSecondary   Variant = "secondary"
	VariantOutline     Variant = "outline"
	VariantGhost       Variant = "ghost"
	VariantLink        Variant = "link"
	VariantDestructive Variant = "destructive"
	VariantSuccess     Variant = "success"
	VariantWarning     Variant = "warning"
	VariantInfo        Variant = "info"
)

// ButtonVariant represents button-specific style variations
type ButtonVariant string

const (
	ButtonDefault     ButtonVariant = "default"
	ButtonPrimary     ButtonVariant = "primary" 
	ButtonSecondary   ButtonVariant = "secondary"
	ButtonOutline     ButtonVariant = "outline"
	ButtonDestructive ButtonVariant = "destructive"
	ButtonGhost       ButtonVariant = "ghost"
	ButtonLink        ButtonVariant = "link"
)

// BadgeVariant represents badge style variations
type BadgeVariant string

const (
	BadgeDefault     BadgeVariant = "default"
	BadgePrimary     BadgeVariant = "primary"
	BadgeSecondary   BadgeVariant = "secondary"
	BadgeSuccess     BadgeVariant = "success"
	BadgeWarning     BadgeVariant = "warning"
	BadgeError       BadgeVariant = "error"
	BadgeDestructive BadgeVariant = "destructive"
	BadgeInfo        BadgeVariant = "info"
)

// AlertVariant represents alert style variations
type AlertVariant string

const (
	AlertDefault     AlertVariant = "default"
	AlertSuccess     AlertVariant = "success"
	AlertWarning     AlertVariant = "warning"
	AlertError       AlertVariant = "error"
	AlertDestructive AlertVariant = "destructive"
	AlertInfo        AlertVariant = "info"
)

// ComponentState represents common component states
type ComponentState struct {
	Disabled bool
	Loading  bool
	Error    bool
	Required bool
	ReadOnly bool
	Active   bool
	Selected bool
	Focused  bool
	Hovered  bool
}

// HTMXConfig represents HTMX-specific configuration
type HTMXConfig struct {
	Post    string // hx-post URL
	Get     string // hx-get URL
	Put     string // hx-put URL
	Patch   string // hx-patch URL
	Delete  string // hx-delete URL
	Target  string // hx-target selector
	Swap    string // hx-swap strategy
	Trigger string // hx-trigger event
	Confirm string // hx-confirm message
	Headers string // hx-headers JSON
	Vals    string // hx-vals JSON
	Include string // hx-include selector
	Boost   bool   // hx-boost
	PushURL string // hx-push-url
}

// AccessibilityProps represents ARIA and accessibility attributes
type AccessibilityProps struct {
	AriaLabel       string // aria-label
	AriaDescribedBy string // aria-describedby
	AriaLabelledBy  string // aria-labelledby
	AriaExpanded    string // aria-expanded
	AriaPressed     string // aria-pressed
	AriaSelected    string // aria-selected
	AriaChecked     string // aria-checked
	AriaDisabled    string // aria-disabled
	AriaHidden      string // aria-hidden
	AriaInvalid     string // aria-invalid
	AriaRequired    string // aria-required
	AriaReadOnly    string // aria-readonly
	AriaLive        string // aria-live
	AriaAtomic      string // aria-atomic
	Role            string // role attribute
	TabIndex        string // tabindex
}

// EventHandlers represents JavaScript event handlers
type EventHandlers struct {
	OnClick     string // onclick
	OnChange    string // onchange
	OnInput     string // oninput
	OnFocus     string // onfocus
	OnBlur      string // onblur
	OnKeyDown   string // onkeydown
	OnKeyUp     string // onkeyup
	OnKeyPress  string // onkeypress
	OnSubmit    string // onsubmit
	OnLoad      string // onload
	OnError     string // onerror
	OnMouseOver string // onmouseover
	OnMouseOut  string // onmouseout
}

// BaseProps represents properties common to all components
type BaseProps struct {
	ID        string             // HTML id attribute
	ClassName string             // Additional CSS classes
	Style     string             // Inline styles
	DataAttrs map[string]string  // data-* attributes
	State     ComponentState     // Component state
	HTMX      HTMXConfig         // HTMX configuration
	A11y      AccessibilityProps // Accessibility properties
	Events    EventHandlers      // Event handlers
}

// ClassBuilder helps build CSS class strings
type ClassBuilder struct {
	classes []string
}

// NewClassBuilder creates a new class builder
func NewClassBuilder(baseClasses ...string) *ClassBuilder {
	return &ClassBuilder{
		classes: append([]string{}, baseClasses...),
	}
}

// Add adds classes unconditionally
func (cb *ClassBuilder) Add(classes ...string) *ClassBuilder {
	cb.classes = append(cb.classes, classes...)
	return cb
}

// AddIf adds classes conditionally
func (cb *ClassBuilder) AddIf(condition bool, classes ...string) *ClassBuilder {
	if condition {
		cb.classes = append(cb.classes, classes...)
	}
	return cb
}

// AddVariant adds a variant class with the specified base
func (cb *ClassBuilder) AddVariant(base string, variant Variant) *ClassBuilder {
	if variant != VariantDefault && variant != VariantPrimary {
		cb.classes = append(cb.classes, base+"-"+string(variant))
	} else {
		cb.classes = append(cb.classes, base)
	}
	return cb
}

// AddSize adds a size class with the specified base
func (cb *ClassBuilder) AddSize(base string, size Size) *ClassBuilder {
	if size != SizeMd && size != SizeDefault {
		cb.classes = append(cb.classes, base+"-"+string(size))
	} else {
		cb.classes = append(cb.classes, base)
	}
	return cb
}

// AddState adds state-based classes
func (cb *ClassBuilder) AddState(state ComponentState) *ClassBuilder {
	cb.AddIf(state.Disabled, "disabled")
	cb.AddIf(state.Loading, "loading")
	cb.AddIf(state.Error, "error")
	cb.AddIf(state.Required, "required")
	cb.AddIf(state.ReadOnly, "readonly")
	cb.AddIf(state.Active, "active")
	cb.AddIf(state.Selected, "selected")
	cb.AddIf(state.Focused, "focused")
	cb.AddIf(state.Hovered, "hovered")
	return cb
}

// Build returns the final class string
func (cb *ClassBuilder) Build() string {
	// Remove empty classes and duplicates
	seen := make(map[string]bool)
	var result []string

	for _, class := range cb.classes {
		class = strings.TrimSpace(class)
		if class != "" && !seen[class] {
			result = append(result, class)
			seen[class] = true
		}
	}

	return strings.Join(result, " ")
}

// LayoutOrientation represents layout orientation options
type LayoutOrientation string

const (
	OrientationVertical   LayoutOrientation = "vertical"
	OrientationHorizontal LayoutOrientation = "horizontal"
)

// Alignment represents alignment options
type Alignment string

const (
	AlignmentStart   Alignment = "start"
	AlignmentCenter  Alignment = "center"
	AlignmentEnd     Alignment = "end"
	AlignmentStretch Alignment = "stretch"
)

// Spacing represents spacing scale
type Spacing string

const (
	SpacingNone Spacing = "none"
	SpacingXS   Spacing = "xs"
	SpacingSM   Spacing = "sm"
	SpacingMD   Spacing = "md"
	SpacingLG   Spacing = "lg"
	SpacingXL   Spacing = "xl"
)

// SlotComponent represents a component slot for composition
type SlotComponent struct {
	Name      string          // Slot name
	Component templ.Component // Component to render
	Props     interface{}     // Component props
}

// HasHTMX returns true if any HTMX attributes are configured
func (h HTMXConfig) HasHTMX() bool {
	return h.Post != "" || h.Get != "" || h.Put != "" || h.Patch != "" || h.Delete != ""
}

// GetMethod returns the HTTP method being used
func (h HTMXConfig) GetMethod() string {
	if h.Post != "" {
		return "POST"
	}
	if h.Get != "" {
		return "GET"
	}
	if h.Put != "" {
		return "PUT"
	}
	if h.Patch != "" {
		return "PATCH"
	}
	if h.Delete != "" {
		return "DELETE"
	}
	return ""
}

// GetURL returns the URL for the configured method
func (h HTMXConfig) GetURL() string {
	if h.Post != "" {
		return h.Post
	}
	if h.Get != "" {
		return h.Get
	}
	if h.Put != "" {
		return h.Put
	}
	if h.Patch != "" {
		return h.Patch
	}
	if h.Delete != "" {
		return h.Delete
	}
	return ""
}

// Merge merges another HTMX config, with the other config taking precedence
func (h HTMXConfig) Merge(other HTMXConfig) HTMXConfig {
	result := h

	if other.Post != "" {
		result.Post = other.Post
	}
	if other.Get != "" {
		result.Get = other.Get
	}
	if other.Put != "" {
		result.Put = other.Put
	}
	if other.Patch != "" {
		result.Patch = other.Patch
	}
	if other.Delete != "" {
		result.Delete = other.Delete
	}
	if other.Target != "" {
		result.Target = other.Target
	}
	if other.Swap != "" {
		result.Swap = other.Swap
	}
	if other.Trigger != "" {
		result.Trigger = other.Trigger
	}
	if other.Confirm != "" {
		result.Confirm = other.Confirm
	}
	if other.Headers != "" {
		result.Headers = other.Headers
	}
	if other.Vals != "" {
		result.Vals = other.Vals
	}
	if other.Include != "" {
		result.Include = other.Include
	}
	if other.Boost {
		result.Boost = other.Boost
	}
	if other.PushURL != "" {
		result.PushURL = other.PushURL
	}

	return result
}

// HasAccessibilityAttrs returns true if any accessibility attributes are set
func (a AccessibilityProps) HasAccessibilityAttrs() bool {
	return a.AriaLabel != "" || a.AriaDescribedBy != "" || a.AriaLabelledBy != "" ||
		a.Role != "" || a.TabIndex != ""
}

// HasEventHandlers returns true if any event handlers are configured
func (e EventHandlers) HasEventHandlers() bool {
	return e.OnClick != "" || e.OnChange != "" || e.OnInput != "" || e.OnFocus != "" ||
		e.OnBlur != "" || e.OnSubmit != ""
}

// GetDataAttrsString converts data attributes map to HTML attribute string
func GetDataAttrsString(dataAttrs map[string]string) string {
	if len(dataAttrs) == 0 {
		return ""
	}

	var attrs []string
	for key, value := range dataAttrs {
		attrs = append(attrs, `data-`+key+`="`+value+`"`)
	}

	return strings.Join(attrs, " ")
}

// ComponentInterface represents the interface all components should implement
type ComponentInterface interface {
	// Validate validates the component props
	Validate() error

	// GetClasses returns the CSS classes for the component
	GetClasses() string

	// GetBaseProps returns the base props for the component
	GetBaseProps() BaseProps
}

// Icon represents icon configuration used across components
type Icon struct {
	Name      string // Icon name/identifier
	Size      Size   // Icon size
	ClassName string // Additional CSS classes
}

// Status represents generic status states
type Status string

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
	StatusDraft    Status = "draft"
	StatusArchived Status = "archived"
	StatusDeleted  Status = "deleted"
)

// InputType represents HTML input types
type InputType string

const (
	InputText     InputType = "text"
	InputEmail    InputType = "email"
	InputPassword InputType = "password"
	InputNumber   InputType = "number"
	InputTel      InputType = "tel"
	InputUrl      InputType = "url"
	InputSearch   InputType = "search"
	InputDate     InputType = "date"
	InputTime     InputType = "time"
	InputDateTime InputType = "datetime-local"
	InputMonth    InputType = "month"
	InputWeek     InputType = "week"
	InputColor    InputType = "color"
	InputFile     InputType = "file"
	InputHidden   InputType = "hidden"
)

// Position represents positioning options
type Position string

const (
	PositionTop         Position = "top"
	PositionRight       Position = "right"
	PositionBottom      Position = "bottom"
	PositionLeft        Position = "left"
	PositionTopLeft     Position = "top-left"
	PositionTopRight    Position = "top-right"
	PositionBottomLeft  Position = "bottom-left"
	PositionBottomRight Position = "bottom-right"
	PositionCenter      Position = "center"
)

// TrendDirection represents directional trends
type TrendDirection string

const (
	TrendUp   TrendDirection = "up"
	TrendDown TrendDirection = "down"
	TrendFlat TrendDirection = "flat"
	TrendNone TrendDirection = "none"
)

// ValidationMessage represents a validation message with severity
type ValidationMessage struct {
	Message  string             // Message text
	Severity ValidationSeverity // Message severity
	Field    string             // Associated field (optional)
}

// ValidationSeverity represents the severity of a validation message
type ValidationSeverity string

const (
	SeverityError   ValidationSeverity = "error"
	SeverityWarning ValidationSeverity = "warning"
	SeverityInfo    ValidationSeverity = "info"
	SeveritySuccess ValidationSeverity = "success"
)

// WithIcon provides icon configuration for components
type WithIcon struct {
	Icon    templ.Component // Icon component to render
	IconPos Position        // Position of icon relative to content
}

// WithSlots provides slot-based composition for complex components
type WithSlots struct {
	Header   templ.Component // Header slot content
	Footer   templ.Component // Footer slot content
	Leading  templ.Component // Leading content (left side)
	Trailing templ.Component // Trailing content (right side)
}
