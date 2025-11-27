package atoms

import (
	"fmt"
	"strings"
)

// ButtonVariant defines the visual style of a button
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "primary"
	ButtonSecondary   ButtonVariant = "secondary"
	ButtonOutline     ButtonVariant = "outline"
	ButtonGhost       ButtonVariant = "ghost"
	ButtonLink        ButtonVariant = "link"
	ButtonDestructive ButtonVariant = "destructive"
)

// ButtonSize defines the size of a button
type ButtonSize string

const (
	ButtonSizeSm ButtonSize = "sm"
	ButtonSizeMd ButtonSize = "md"
	ButtonSizeLg ButtonSize = "lg"
)

// ButtonType defines whether the button is text, icon-only, or has both
type ButtonType string

const (
	ButtonTypeDefault ButtonType = "default" // text with optional icon
	ButtonTypeIcon    ButtonType = "icon"    // icon only
)

// ButtonProps defines all properties for a button component
type ButtonProps struct {
	// Content
	Label string // Button text (optional for icon buttons)
	Icon  string // Icon name (optional)
	
	// Styling
	Variant   ButtonVariant // Visual style
	Size      ButtonSize    // Size variant
	Type      ButtonType    // Icon-only vs text button
	Disabled  bool          // Disabled state
	Loading   bool          // Loading state with spinner
	ClassName string        // Additional CSS classes
	
	// HTML attributes
	ID       string // HTML id attribute
	HTMLType string // button type: "button", "submit", "reset"
	
	// Interactivity
	OnClick  string // JavaScript click handler
	HxPost   string // HTMX post URL
	HxGet    string // HTMX get URL
	HxTarget string // HTMX target selector
	HxSwap   string // HTMX swap strategy
	
	// Accessibility
	AriaLabel       string // aria-label for accessibility
	AriaPressed     string // aria-pressed for toggle buttons
	AriaExpanded    string // aria-expanded for dropdown triggers
	AriaDescribedBy string // aria-describedby reference
	Role            string // ARIA role override
	TabIndex        string // tabindex attribute
}

// GetButtonClasses generates the appropriate CSS classes for a button
func GetButtonClasses(variant ButtonVariant, size ButtonSize, buttonType ButtonType, disabled, loading bool) string {
	var classes []string
	
	// Base class with variant and type
	baseClass := "btn"
	if size != ButtonSizeMd {
		baseClass += "-" + string(size)
	}
	if buttonType == ButtonTypeIcon {
		baseClass += "-icon"
	}
	if variant != ButtonPrimary {
		baseClass += "-" + string(variant)
	}
	
	classes = append(classes, baseClass)
	
	// State classes
	if disabled {
		classes = append(classes, "disabled")
	}
	if loading {
		classes = append(classes, "loading")
	}
	
	return strings.Join(classes, " ")
}

// GetDefaultButtonProps returns button props with sensible defaults
func GetDefaultButtonProps() ButtonProps {
	return ButtonProps{
		Variant:   ButtonPrimary,
		Size:      ButtonSizeMd,
		Type:      ButtonTypeDefault,
		HTMLType:  "button",
		Disabled:  false,
		Loading:   false,
		ClassName: "",
	}
}

// WithVariant returns a copy of props with the specified variant
func (p ButtonProps) WithVariant(variant ButtonVariant) ButtonProps {
	p.Variant = variant
	return p
}

// WithSize returns a copy of props with the specified size
func (p ButtonProps) WithSize(size ButtonSize) ButtonProps {
	p.Size = size
	return p
}

// WithIcon returns a copy of props configured as an icon-only button
func (p ButtonProps) WithIcon(iconName string) ButtonProps {
	p.Icon = iconName
	p.Type = ButtonTypeIcon
	p.Label = "" // Clear label for icon-only buttons
	return p
}

// WithHtmx returns a copy of props with HTMX configuration
func (p ButtonProps) WithHtmx(method, url, target, swap string) ButtonProps {
	switch method {
	case "post":
		p.HxPost = url
	case "get":
		p.HxGet = url
	}
	p.HxTarget = target
	p.HxSwap = swap
	return p
}

// IsIconOnly returns true if this is an icon-only button
func (p ButtonProps) IsIconOnly() bool {
	return p.Type == ButtonTypeIcon
}

// HasIcon returns true if the button has an icon
func (p ButtonProps) HasIcon() bool {
	return p.Icon != ""
}

// GetAriaLabel returns the appropriate aria-label value
func (p ButtonProps) GetAriaLabel() string {
	if p.AriaLabel != "" {
		return p.AriaLabel
	}
	// For icon-only buttons, try to generate a meaningful label
	if p.IsIconOnly() && p.Icon != "" {
		return strings.Title(strings.ReplaceAll(p.Icon, "-", " "))
	}
	return ""
}

// Validate checks if the button props are valid
func (p ButtonProps) Validate() error {
	// Icon-only buttons must have an icon
	if p.IsIconOnly() && p.Icon == "" {
		return fmt.Errorf("icon-only buttons must specify an Icon")
	}
	
	// Regular buttons should have either a label or icon
	if p.Type == ButtonTypeDefault && p.Label == "" && p.Icon == "" {
		return fmt.Errorf("buttons must have either a Label or Icon")
	}
	
	return nil
}