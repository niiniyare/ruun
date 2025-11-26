// pkg/schema/action_v2.go
package schema

import (
	"context"

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

	// Unified Configuration
	Binding      []string         `json:"binding,omitempty"`
	Conditional  []string         `json:"conditional,omitempty"`
	Permissions  []string         `json:"permissions,omitempty"`
	Config       map[string]any   `json:"config,omitempty"`
	
	// Unified styling (replaces old Style []string and ActionTheme)
	Style        *Style             `json:"style,omitempty"`
	ActionConfig *ActionConfig      `json:"action_config,omitempty"`

	// Behavior and interaction (replaces HTMX/ActionHTMX)
	Behavior     *Behavior          `json:"behavior,omitempty"` // Replaces old []string behavior
	
	// Legacy Support
	Confirm     *ActionConfirm     `json:"confirm,omitempty"`

	// Internal
	evaluator *condition.Evaluator `json:"-"`
}

// Missing type definitions for compatibility


// ActionHTMX is a type alias for backward compatibility
// Use Behavior instead for new code
type ActionHTMX = Behavior

// ActionTheme defines theme overrides for this action
















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

	if len(a.Conditional) == 0 {
		return true, nil
	}

	// Check conditional expressions
	for _, cond := range a.Conditional {
		if cond == "hidden" {
			return false, nil
		}
	}

	return true, nil
}

// IsEnabled checks if action can be executed
func (a *Action) IsEnabled(ctx context.Context, data map[string]any) (bool, error) {
	if a.Disabled || a.Loading {
		return false, nil
	}

	if len(a.Conditional) == 0 {
		return true, nil
	}

	// Check conditional expressions
	for _, cond := range a.Conditional {
		if cond == "disabled" {
			return false, nil
		}
	}

	return true, nil
}

// CanView checks if user can view action
func (a *Action) CanView(userRoles []string) bool {
	if len(a.Permissions) == 0 {
		return true
	}

	for _, role := range userRoles {
		for _, required := range a.Permissions {
			if role == required {
				return true
			}
		}
	}
	return false
}

// CanExecute checks if user can execute action
func (a *Action) CanExecute(userRoles []string) bool {
	if len(a.Permissions) == 0 {
		return true
	}

	for _, role := range userRoles {
		for _, required := range a.Permissions {
			if role == required {
				return true
			}
		}
	}
	return false
}

// HasPermission checks if action requires permission
func (a *Action) HasPermission(permission string) bool {
	if len(a.Permissions) == 0 {
		return false
	}

	for _, perm := range a.Permissions {
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
	// Try new Behavior first (unified approach)
	if a.Behavior != nil && a.Behavior.URL != "" {
		return a.Behavior.URL
	}
	
	// Try new Behavior field
	if a.Behavior != nil && a.Behavior.URL != "" {
		return a.Behavior.URL
	}
	
	// Fall back to Config
	if a.Type == ActionLink || a.Config != nil {
		if url, ok := a.Config["url"].(string); ok {
			return url
		}
	}
	return ""
}

// GetHTTPMethod returns HTTP method
func (a *Action) GetHTTPMethod() string {
	// Try new Behavior field first
	if a.Behavior != nil && a.Behavior.Method != "" {
		return a.Behavior.Method
	}
	
	// Try Config
	if a.Config != nil {
		if method, ok := a.Config["method"].(string); ok && method != "" {
			return method
		}
	}
	
	// Default based on action type
	if a.Type == ActionSubmit {
		return "POST"
	}
	return "GET"
}

// ShouldDebounce checks if action should be debounced
func (a *Action) ShouldDebounce() bool {
	if a.Config != nil {
		if debounce, ok := a.Config["debounce"].(int); ok && debounce > 0 {
			return true
		}
	}
	return false
}

// GetDebounceDelay returns debounce delay
func (a *Action) GetDebounceDelay() int {
	if a.Config != nil {
		if debounce, ok := a.Config["debounce"].(int); ok {
			return debounce
		}
	}
	return 0
}

// ShouldThrottle checks if action should be throttled
func (a *Action) ShouldThrottle() bool {
	if a.Config != nil {
		if throttle, ok := a.Config["throttle"].(int); ok && throttle > 0 {
			return true
		}
	}
	return false
}

// GetThrottleDelay returns throttle delay
func (a *Action) GetThrottleDelay() int {
	if a.Config != nil {
		if throttle, ok := a.Config["throttle"].(int); ok {
			return throttle
		}
	}
	return 0
}

// GetIconPosition returns the icon position (defaults to "left")
func (a *Action) GetIconPosition() string {
	if a.Position != "" {
		return a.Position
	}
	return "left"
}

// HasIcon returns true if the action has an icon
func (a *Action) HasIcon() bool {
	return a.Icon != ""
}

// HasConfirmation returns true if the action has confirmation enabled
func (a *Action) HasConfirmation() bool {
	return a.Confirm != nil && a.Confirm.Enabled
}

// RequiresPermission checks if the action requires a specific permission
func (a *Action) RequiresPermission(permission string) bool {
	for _, perm := range a.Permissions {
		if perm == permission {
			return true
		}
	}
	return false
}


// GetConfigMap returns the action config map, initializing it if nil
func (a *Action) GetConfigMap() map[string]any {
	if a.Config == nil {
		a.Config = make(map[string]any)
	}
	return a.Config
}

// HasBehavior returns true if the action has behavior configuration
func (a *Action) HasBehavior() bool {
	return a != nil && a.Behavior != nil && !a.Behavior.IsEmpty()
}



// GetVariantClass returns CSS class based on variant
func (a *Action) GetVariantClass() string {
	switch a.Variant {
	case "primary":
		return "action-primary"
	case "secondary":
		return "action-secondary"
	case "outline":
		return "action-outline"
	case "ghost":
		return "action-ghost"
	case "destructive":
		return "action-destructive"
	default:
		return "action-primary" // Default fallback
	}
}

// GetSizeClass returns CSS class based on size
func (a *Action) GetSizeClass() string {
	switch a.Size {
	case "sm":
		return "action-sm"
	case "md":
		return "action-md"
	case "lg":
		return "action-lg"
	case "xl":
		return "action-xl"
	default:
		return "action-md" // Default fallback
	}
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
		if a.Config == nil {
			collector.AddValidationError(a.ID, "missing_url",
				"link action requires URL")
		} else if url, ok := a.Config["url"].(string); !ok || url == "" {
			collector.AddValidationError(a.ID, "missing_url",
				"link action requires URL")
		}
	}

	// Validate custom actions have handler
	if a.Type == ActionCustom {
		if a.Config == nil {
			collector.AddValidationError(a.ID, "missing_handler",
				"custom action requires handler")
		} else if handler, ok := a.Config["handler"].(string); !ok || handler == "" {
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

func (b *ActionBuilder) WithConfig(config map[string]any) *ActionBuilder {
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

func (b *ActionBuilder) WithConditional(conditional []string) *ActionBuilder {
	b.action.Conditional = conditional
	return b
}

func (b *ActionBuilder) WithCondition(condition []string) *ActionBuilder {
	b.action.Conditional = condition
	return b
}

func (b *ActionBuilder) WithPermissions(permissions []string) *ActionBuilder {
	b.action.Permissions = permissions
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
		WithConfig(map[string]any{
			"url": url,
		})
}
