// pkg/schema/action_v2.go
package schema

import (
	"context"
	"fmt"

	"github.com/niiniyare/ruun/pkg/condition"
)

// Action represents a button or actionable element
type Action struct {
	// Identity
	ID   string     `json:"id"`
	Type ActionType `json:"type"`
	Text string     `json:"text"`

	// Appearance
	Variant  string `json:"variant,omitempty"`
	Size     string `json:"size,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Position string `json:"position,omitempty"`

	// State
	Loading  bool `json:"loading,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Hidden   bool `json:"hidden,omitempty"`

	// Behavior
	Config      *ActionConfig      `json:"config,omitempty"`
	Confirm     *ActionConfirm     `json:"confirm,omitempty"`
	Conditional *ActionConditional `json:"conditional,omitempty"`
	Permissions *ActionPermissions `json:"permissions,omitempty"`

	// Framework Integration
	HTMX   *ActionHTMX   `json:"htmx,omitempty"`
	Alpine *ActionAlpine `json:"alpine,omitempty"`

	// Internal
	evaluator *condition.Evaluator `json:"-"`
}

// ActionType defines the type of action
type ActionType string

const (
	ActionSubmit ActionType = "submit"
	ActionReset  ActionType = "reset"
	ActionButton ActionType = "button"
	ActionLink   ActionType = "link"
	ActionCustom ActionType = "custom"
)

// ActionConfig holds action configuration
type ActionConfig struct {
	URL             string            `json:"url,omitempty"`
	Target          string            `json:"target,omitempty"`
	Handler         string            `json:"handler,omitempty"`
	Params          map[string]string `json:"params,omitempty"`
	Debounce        int               `json:"debounce,omitempty"`
	Throttle        int               `json:"throttle,omitempty"`
	PreventDefault  bool              `json:"preventDefault,omitempty"`
	StopPropagation bool              `json:"stopPropagation,omitempty"`
	SuccessMessage  string            `json:"successMessage,omitempty"`
	ErrorMessage    string            `json:"errorMessage,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty"`
}

// ActionConfirm defines confirmation dialog
type ActionConfirm struct {
	Enabled bool   `json:"enabled"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message"`
	Confirm string `json:"confirm,omitempty"`
	Cancel  string `json:"cancel,omitempty"`
	Variant string `json:"variant,omitempty"`
	Icon    string `json:"icon,omitempty"`
}

// ActionConditional defines conditional visibility
type ActionConditional struct {
	Show    *condition.ConditionGroup `json:"show,omitempty"`
	Hide    *condition.ConditionGroup `json:"hide,omitempty"`
	Enable  *condition.ConditionGroup `json:"enable,omitempty"`
	Disable *condition.ConditionGroup `json:"disable,omitempty"`
}

// ActionPermissions controls action access
type ActionPermissions struct {
	View     []string `json:"view,omitempty"`
	Execute  []string `json:"execute,omitempty"`
	Required []string `json:"required,omitempty"`
}

// ActionHTMX defines HTMX behavior
type ActionHTMX struct {
	Method    string            `json:"method,omitempty"`
	URL       string            `json:"url,omitempty"`
	Target    string            `json:"target,omitempty"`
	Swap      string            `json:"swap,omitempty"`
	Trigger   string            `json:"trigger,omitempty"`
	Indicator string            `json:"indicator,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
}

// ActionAlpine defines Alpine.js bindings
type ActionAlpine struct {
	XOn   string `json:"xOn,omitempty"`
	XBind string `json:"xBind,omitempty"`
	XShow string `json:"xShow,omitempty"`
}

// Core Methods

// SetEvaluator sets the condition evaluator
func (a *Action) SetEvaluator(evaluator *condition.Evaluator) {
	a.evaluator = evaluator
}

// GetEvaluator returns the condition evaluator
func (a *Action) GetEvaluator() *condition.Evaluator {
	return a.evaluator
}

// Validate validates action configuration
func (a *Action) Validate(ctx context.Context) error {
	validator := newActionValidator()
	return validator.Validate(ctx, a)
}

// IsVisible checks if action should be displayed
func (a *Action) IsVisible(ctx context.Context, data map[string]any) (bool, error) {
	if a.Hidden {
		return false, nil
	}

	if a.Conditional == nil {
		return true, nil
	}

	// Check Hide condition first
	if a.Conditional.Hide != nil {
		if a.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		hide, err := a.evaluator.Evaluate(ctx, a.Conditional.Hide, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate hide condition: %w", err)
		}
		if hide {
			return false, nil
		}
	}

	// Check Show condition
	if a.Conditional.Show != nil {
		if a.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		show, err := a.evaluator.Evaluate(ctx, a.Conditional.Show, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate show condition: %w", err)
		}
		return show, nil
	}

	return true, nil
}

// IsEnabled checks if action can be executed
func (a *Action) IsEnabled(ctx context.Context, data map[string]any) (bool, error) {
	if a.Disabled || a.Loading {
		return false, nil
	}

	if a.Conditional == nil {
		return true, nil
	}

	// Check Disable condition first
	if a.Conditional.Disable != nil {
		if a.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		disable, err := a.evaluator.Evaluate(ctx, a.Conditional.Disable, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate disable condition: %w", err)
		}
		if disable {
			return false, nil
		}
	}

	// Check Enable condition
	if a.Conditional.Enable != nil {
		if a.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		enable, err := a.evaluator.Evaluate(ctx, a.Conditional.Enable, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate enable condition: %w", err)
		}
		return enable, nil
	}

	return true, nil
}

// CanView checks if user can view action
func (a *Action) CanView(userRoles []string) bool {
	if a.Permissions == nil || len(a.Permissions.View) == 0 {
		return true
	}

	for _, role := range userRoles {
		for _, required := range a.Permissions.View {
			if role == required {
				return true
			}
		}
	}
	return false
}

// CanExecute checks if user can execute action
func (a *Action) CanExecute(userRoles []string) bool {
	if a.Permissions == nil || len(a.Permissions.Execute) == 0 {
		return true
	}

	for _, role := range userRoles {
		for _, required := range a.Permissions.Execute {
			if role == required {
				return true
			}
		}
	}
	return false
}

// HasPermission checks if action requires permission
func (a *Action) HasPermission(permission string) bool {
	if a.Permissions == nil {
		return false
	}

	for _, perm := range a.Permissions.Required {
		if perm == permission {
			return true
		}
	}
	return false
}

// RequiresConfirmation checks if action needs confirmation
func (a *Action) RequiresConfirmation() bool {
	return a.Confirm != nil && a.Confirm.Enabled
}

// Type Checking Methods

// IsSubmitAction checks if this is a submit action
func (a *Action) IsSubmitAction() bool {
	return a.Type == ActionSubmit
}

// IsResetAction checks if this is a reset action
func (a *Action) IsResetAction() bool {
	return a.Type == ActionReset
}

// IsLinkAction checks if this is a link action
func (a *Action) IsLinkAction() bool {
	return a.Type == ActionLink
}

// IsCustomAction checks if this is a custom action
func (a *Action) IsCustomAction() bool {
	return a.Type == ActionCustom
}

// Configuration Methods

// GetURL returns action URL
func (a *Action) GetURL() string {
	if a.Type == ActionLink && a.Config != nil {
		return a.Config.URL
	}
	if a.HTMX != nil {
		return a.HTMX.URL
	}
	return ""
}

// GetHTTPMethod returns HTTP method
func (a *Action) GetHTTPMethod() string {
	if a.HTMX != nil && a.HTMX.Method != "" {
		return a.HTMX.Method
	}
	if a.Type == ActionSubmit {
		return "POST"
	}
	return "GET"
}

// ShouldDebounce checks if action should be debounced
func (a *Action) ShouldDebounce() bool {
	return a.Config != nil && a.Config.Debounce > 0
}

// GetDebounceDelay returns debounce delay
func (a *Action) GetDebounceDelay() int {
	if a.Config != nil {
		return a.Config.Debounce
	}
	return 0
}

// ShouldThrottle checks if action should be throttled
func (a *Action) ShouldThrottle() bool {
	return a.Config != nil && a.Config.Throttle > 0
}

// GetThrottleDelay returns throttle delay
func (a *Action) GetThrottleDelay() int {
	if a.Config != nil {
		return a.Config.Throttle
	}
	return 0
}

// State Management

// SetLoading sets loading state
func (a *Action) SetLoading(loading bool) {
	a.Loading = loading
}

// SetDisabled sets disabled state
func (a *Action) SetDisabled(disabled bool) {
	a.Disabled = disabled
}

// SetHidden sets hidden state
func (a *Action) SetHidden(hidden bool) {
	a.Hidden = hidden
}

// Clone creates a copy of the action
func (a *Action) Clone() *Action {
	clone := *a
	clone.evaluator = a.evaluator
	return &clone
}

// actionValidator validates action configuration
type actionValidator struct{}

func newActionValidator() *actionValidator {
	return &actionValidator{}
}

func (v *actionValidator) Validate(ctx context.Context, a *Action) error {
	collector := NewErrorCollector()

	if a.ID == "" {
		collector.AddValidationError("", "action_id", "action ID is required")
	}
	if a.Text == "" {
		collector.AddValidationError(a.ID, "action_text", "action text is required")
	}
	if a.Type == "" {
		a.Type = ActionButton // Default
	}

	// Validate link actions have URL
	if a.Type == ActionLink {
		if a.Config == nil || a.Config.URL == "" {
			collector.AddValidationError(a.ID, "missing_url",
				"link action requires URL")
		}
	}

	// Validate custom actions have handler
	if a.Type == ActionCustom {
		if a.Config == nil || a.Config.Handler == "" {
			collector.AddValidationError(a.ID, "missing_handler",
				"custom action requires handler")
		}
	}

	// Validate confirmation
	if a.Confirm != nil && a.Confirm.Enabled && a.Confirm.Message == "" {
		collector.AddValidationError(a.ID, "missing_confirm_message",
			"confirmation requires message")
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

// ActionBuilder provides fluent API for building actions
type ActionBuilder struct {
	action    Action
	evaluator *condition.Evaluator
}

// NewAction starts building an action
func NewAction(id string, actionType ActionType, text string) *ActionBuilder {
	return &ActionBuilder{
		action: Action{
			ID:   id,
			Type: actionType,
			Text: text,
		},
	}
}

func (b *ActionBuilder) WithVariant(variant string) *ActionBuilder {
	b.action.Variant = variant
	return b
}

func (b *ActionBuilder) WithSize(size string) *ActionBuilder {
	b.action.Size = size
	return b
}

func (b *ActionBuilder) WithIcon(icon, position string) *ActionBuilder {
	b.action.Icon = icon
	b.action.Position = position
	return b
}

func (b *ActionBuilder) Disabled() *ActionBuilder {
	b.action.Disabled = true
	return b
}

func (b *ActionBuilder) Hidden() *ActionBuilder {
	b.action.Hidden = true
	return b
}

func (b *ActionBuilder) WithConfig(config *ActionConfig) *ActionBuilder {
	b.action.Config = config
	return b
}

func (b *ActionBuilder) WithConfirmation(message, title string) *ActionBuilder {
	b.action.Confirm = &ActionConfirm{
		Enabled: true,
		Title:   title,
		Message: message,
	}
	return b
}

func (b *ActionBuilder) WithConditional(conditional *ActionConditional) *ActionBuilder {
	b.action.Conditional = conditional
	return b
}

func (b *ActionBuilder) WithPermissions(permissions *ActionPermissions) *ActionBuilder {
	b.action.Permissions = permissions
	return b
}

func (b *ActionBuilder) WithHTMX(htmx *ActionHTMX) *ActionBuilder {
	b.action.HTMX = htmx
	return b
}

func (b *ActionBuilder) WithAlpine(alpine *ActionAlpine) *ActionBuilder {
	b.action.Alpine = alpine
	return b
}

func (b *ActionBuilder) WithEvaluator(evaluator *condition.Evaluator) *ActionBuilder {
	b.evaluator = evaluator
	return b
}

func (b *ActionBuilder) Build() Action {
	if b.evaluator != nil {
		b.action.SetEvaluator(b.evaluator)
	}
	return b.action
}

// Convenience Constructors

// NewSubmitAction creates a submit action
func NewSubmitAction(id, text string) *ActionBuilder {
	return NewAction(id, ActionSubmit, text).WithVariant("primary")
}

// NewResetAction creates a reset action
func NewResetAction(id, text string) *ActionBuilder {
	return NewAction(id, ActionReset, text).WithVariant("secondary")
}

// NewCancelAction creates a cancel action
func NewCancelAction(id, text string) *ActionBuilder {
	return NewAction(id, ActionButton, text).WithVariant("outline")
}

// NewDeleteAction creates a delete action with confirmation
func NewDeleteAction(id, text string) *ActionBuilder {
	return NewAction(id, ActionButton, text).
		WithVariant("destructive").
		WithConfirmation("Are you sure you want to delete this?", "Confirm Deletion")
}

// NewLinkAction creates a link action
func NewLinkAction(id, text, url string) *ActionBuilder {
	return NewAction(id, ActionLink, text).
		WithConfig(&ActionConfig{
			URL: url,
		})
}
