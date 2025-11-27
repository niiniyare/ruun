package molecules

import (
	"strings"
	"github.com/your-org/your-project/views/components/atoms"
	"github.com/a-h/templ"
)

// CardProps defines all properties for a card component
type CardProps struct {
	// Header content
	Title       string           // Card title
	Description string           // Card description/subtitle
	Badge       *atoms.BadgeProps // Optional badge in header
	
	// Content
	Content templ.Component // Main card content
	
	// Footer actions
	Actions []CardAction // Action buttons in footer
	
	// Styling & Layout
	ClassName string // Additional CSS classes
	
	// Border variants
	BorderTop    bool   // Top border accent
	BorderColor  string // Border accent color class
	
	// Interactive behavior
	Clickable bool   // Whether entire card is clickable
	Href      string // URL if card is a link
	OnClick   string // JavaScript click handler
	HxPost    string // HTMX post URL
	HxGet     string // HTMX get URL
	HxTarget  string // HTMX target selector
	HxSwap    string // HTMX swap strategy
	
	// Accessibility
	AriaLabel       string // aria-label for accessibility
	AriaDescribedBy string // aria-describedby reference
	Role            string // ARIA role override
	TabIndex        string // tabindex for focusable cards
}

// CardAction represents an action button in the card footer
type CardAction struct {
	Label     string               // Button text
	Variant   atoms.ButtonVariant  // Button style
	Size      atoms.ButtonSize     // Button size
	Icon      string               // Optional icon
	OnClick   string               // JavaScript click handler
	HxPost    string               // HTMX post URL
	HxGet     string               // HTMX get URL
	HxTarget  string               // HTMX target selector
	HxSwap    string               // HTMX swap strategy
	ClassName string               // Additional CSS classes
}

// GetCardClasses generates the appropriate CSS classes for a card
func GetCardClasses(borderTop bool, clickable bool) string {
	var classes []string
	
	classes = append(classes, "card")
	
	if borderTop {
		classes = append(classes, "card-accent-border")
	}
	
	if clickable {
		classes = append(classes, "card-clickable")
	}
	
	return strings.Join(classes, " ")
}

// GetDefaultCardProps returns card props with sensible defaults
func GetDefaultCardProps() CardProps {
	return CardProps{
		Clickable:   false,
		BorderTop:   false,
		ClassName:   "",
		Actions:     []CardAction{},
	}
}

// WithTitle returns a copy of props with the specified title
func (p CardProps) WithTitle(title string) CardProps {
	p.Title = title
	return p
}

// WithDescription returns a copy of props with the specified description
func (p CardProps) WithDescription(description string) CardProps {
	p.Description = description
	return p
}

// WithBadge returns a copy of props with a badge in the header
func (p CardProps) WithBadge(badge atoms.BadgeProps) CardProps {
	p.Badge = &badge
	return p
}

// WithContent returns a copy of props with the specified content
func (p CardProps) WithContent(content templ.Component) CardProps {
	p.Content = content
	return p
}

// WithBorderTop returns a copy of props with top border accent
func (p CardProps) WithBorderTop(color string) CardProps {
	p.BorderTop = true
	p.BorderColor = color
	return p
}

// WithClickable returns a copy of props configured as clickable
func (p CardProps) WithClickable(href string) CardProps {
	p.Clickable = true
	p.Href = href
	return p
}

// WithOnClick returns a copy of props with click handler
func (p CardProps) WithOnClick(onClick string) CardProps {
	p.Clickable = true
	p.OnClick = onClick
	return p
}

// WithHtmx returns a copy of props with HTMX configuration
func (p CardProps) WithHtmx(method, url, target, swap string) CardProps {
	p.Clickable = true
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

// WithActions returns a copy of props with action buttons
func (p CardProps) WithActions(actions ...CardAction) CardProps {
	p.Actions = actions
	return p
}

// AddAction adds an action to the card and returns the modified props
func (p CardProps) AddAction(action CardAction) CardProps {
	p.Actions = append(p.Actions, action)
	return p
}

// HasHeader returns true if the card has header content
func (p CardProps) HasHeader() bool {
	return p.Title != "" || p.Description != "" || p.Badge != nil || len(p.Actions) > 0
}

// HasTitle returns true if the card has a title
func (p CardProps) HasTitle() bool {
	return p.Title != ""
}

// HasDescription returns true if the card has a description
func (p CardProps) HasDescription() bool {
	return p.Description != ""
}

// HasBadge returns true if the card has a badge
func (p CardProps) HasBadge() bool {
	return p.Badge != nil
}

// HasActions returns true if the card has action buttons
func (p CardProps) HasActions() bool {
	return len(p.Actions) > 0
}

// HasFooter returns true if the card has footer content
func (p CardProps) HasFooter() bool {
	return p.HasActions()
}

// IsClickable returns true if the card has click behavior
func (p CardProps) IsClickable() bool {
	return p.Clickable
}

// GetElementTag returns the appropriate HTML element tag
func (p CardProps) GetElementTag() string {
	if p.Href != "" {
		return "a"
	}
	if p.IsClickable() {
		return "button"
	}
	return "div"
}

// GetAriaLabel returns the appropriate aria-label value
func (p CardProps) GetAriaLabel() string {
	if p.AriaLabel != "" {
		return p.AriaLabel
	}
	
	// Generate aria-label for clickable cards
	if p.IsClickable() && p.Title != "" {
		return "Open " + p.Title
	}
	
	return ""
}

// GetTabIndex returns the appropriate tabindex value
func (p CardProps) GetTabIndex() string {
	if p.TabIndex != "" {
		return p.TabIndex
	}
	
	// Clickable cards should be focusable
	if p.IsClickable() {
		return "0"
	}
	
	return ""
}

// NewCardAction creates a new card action with default properties
func NewCardAction(label string) CardAction {
	return CardAction{
		Label:   label,
		Variant: atoms.ButtonSecondary,
		Size:    atoms.ButtonSizeSm,
	}
}

// WithVariant returns a copy of the action with the specified variant
func (a CardAction) WithVariant(variant atoms.ButtonVariant) CardAction {
	a.Variant = variant
	return a
}

// WithSize returns a copy of the action with the specified size
func (a CardAction) WithSize(size atoms.ButtonSize) CardAction {
	a.Size = size
	return a
}

// WithIcon returns a copy of the action with an icon
func (a CardAction) WithIcon(icon string) CardAction {
	a.Icon = icon
	return a
}

// WithHtmx returns a copy of the action with HTMX configuration
func (a CardAction) WithHtmx(method, url, target, swap string) CardAction {
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

// ToButtonProps converts card action to button props
func (a CardAction) ToButtonProps() atoms.ButtonProps {
	buttonProps := atoms.GetDefaultButtonProps()
	buttonProps.Label = a.Label
	buttonProps.Variant = a.Variant
	buttonProps.Size = a.Size
	buttonProps.Icon = a.Icon
	buttonProps.OnClick = a.OnClick
	buttonProps.HxPost = a.HxPost
	buttonProps.HxGet = a.HxGet
	buttonProps.HxTarget = a.HxTarget
	buttonProps.HxSwap = a.HxSwap
	buttonProps.ClassName = a.ClassName
	
	return buttonProps
}