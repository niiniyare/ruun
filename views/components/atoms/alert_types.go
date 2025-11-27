package atoms

import (
	"strings"
)

// AlertVariant defines the visual style and semantic meaning of an alert
type AlertVariant string

const (
	AlertDefault     AlertVariant = "default"
	AlertDestructive AlertVariant = "destructive"
	AlertSuccess     AlertVariant = "success"
	AlertWarning     AlertVariant = "warning"
	AlertInfo        AlertVariant = "info"
)

// AlertProps defines all properties for an alert component
type AlertProps struct {
	// Content
	Title       string // Alert title/heading
	Message     string // Alert message content
	Description string // Additional description text
	Icon        string // Optional icon name
	
	// Styling
	Variant   AlertVariant // Visual style and semantic meaning
	ClassName string       // Additional CSS classes
	
	// Behavior
	Dismissible bool   // Whether alert can be dismissed
	AutoDismiss int    // Auto-dismiss after N seconds (0 = no auto-dismiss)
	OnDismiss   string // JavaScript handler for dismiss action
	
	// Actions
	Actions []AlertAction // Action buttons in the alert
	
	// Accessibility
	Role            string // ARIA role (default: "alert")
	AriaLabel       string // aria-label for accessibility
	AriaDescribedBy string // aria-describedby reference
	AriaLive        string // aria-live for dynamic alerts ("polite", "assertive")
}

// AlertAction represents an action button within an alert
type AlertAction struct {
	Label     string        // Button text
	Variant   ButtonVariant // Button style variant
	OnClick   string        // JavaScript click handler
	HxPost    string        // HTMX post URL
	HxGet     string        // HTMX get URL
	HxTarget  string        // HTMX target selector
	HxSwap    string        // HTMX swap strategy
	ClassName string        // Additional CSS classes
}

// GetAlertClasses generates the appropriate CSS classes for an alert
func GetAlertClasses(variant AlertVariant) string {
	baseClass := "alert"
	
	if variant != AlertDefault {
		return baseClass + "-" + string(variant)
	}
	
	return baseClass
}

// GetDefaultAlertProps returns alert props with sensible defaults
func GetDefaultAlertProps() AlertProps {
	return AlertProps{
		Variant:      AlertDefault,
		Dismissible:  false,
		AutoDismiss:  0,
		Role:         "alert",
		AriaLive:     "polite",
		ClassName:    "",
		Actions:      []AlertAction{},
	}
}

// WithVariant returns a copy of props with the specified variant
func (p AlertProps) WithVariant(variant AlertVariant) AlertProps {
	p.Variant = variant
	// Set appropriate aria-live for different variants
	switch variant {
	case AlertDestructive:
		p.AriaLive = "assertive"
	default:
		p.AriaLive = "polite"
	}
	return p
}

// WithIcon returns a copy of props with an icon
func (p AlertProps) WithIcon(iconName string) AlertProps {
	p.Icon = iconName
	return p
}

// WithDismissible returns a copy of props configured as dismissible
func (p AlertProps) WithDismissible(onDismiss string) AlertProps {
	p.Dismissible = true
	p.OnDismiss = onDismiss
	return p
}

// WithAutoDismiss returns a copy of props with auto-dismiss configuration
func (p AlertProps) WithAutoDismiss(seconds int) AlertProps {
	p.AutoDismiss = seconds
	return p
}

// WithActions returns a copy of props with action buttons
func (p AlertProps) WithActions(actions ...AlertAction) AlertProps {
	p.Actions = actions
	return p
}

// AddAction adds an action to the alert and returns the modified props
func (p AlertProps) AddAction(action AlertAction) AlertProps {
	p.Actions = append(p.Actions, action)
	return p
}

// HasIcon returns true if the alert has an icon
func (p AlertProps) HasIcon() bool {
	return p.Icon != ""
}

// HasTitle returns true if the alert has a title
func (p AlertProps) HasTitle() bool {
	return p.Title != ""
}

// HasDescription returns true if the alert has a description
func (p AlertProps) HasDescription() bool {
	return p.Description != ""
}

// HasActions returns true if the alert has action buttons
func (p AlertProps) HasActions() bool {
	return len(p.Actions) > 0
}

// GetDefaultIcon returns the default icon for each variant
func (p AlertProps) GetDefaultIcon() string {
	if p.Icon != "" {
		return p.Icon
	}
	
	// Return default icons based on variant
	switch p.Variant {
	case AlertDestructive:
		return "alert-circle"
	case AlertSuccess:
		return "check-circle"
	case AlertWarning:
		return "alert-triangle"
	case AlertInfo:
		return "info"
	default:
		return "info"
	}
}

// GetAriaLabel returns the appropriate aria-label value
func (p AlertProps) GetAriaLabel() string {
	if p.AriaLabel != "" {
		return p.AriaLabel
	}
	
	// Generate aria-label based on variant
	switch p.Variant {
	case AlertDestructive:
		return "Error alert"
	case AlertSuccess:
		return "Success alert"
	case AlertWarning:
		return "Warning alert"
	case AlertInfo:
		return "Information alert"
	default:
		return "Alert"
	}
}

// NewAlertAction creates a new alert action with the specified properties
func NewAlertAction(label string) AlertAction {
	return AlertAction{
		Label:   label,
		Variant: ButtonSecondary, // Default to secondary for alert actions
	}
}

// WithVariant returns a copy of the action with the specified button variant
func (a AlertAction) WithVariant(variant ButtonVariant) AlertAction {
	a.Variant = variant
	return a
}

// WithHtmx returns a copy of the action with HTMX configuration
func (a AlertAction) WithHtmx(method, url, target, swap string) AlertAction {
	switch method {
	case "post":
		a.HxPost = url
	case "get":
		a.HxGet = url
	}
	a.HxTarget = target
	a.HxSwap = swap
	return a
}