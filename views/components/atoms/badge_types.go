package atoms

import (
	"strings"
)

// BadgeVariant defines the visual style of a badge
type BadgeVariant string

const (
	BadgePrimary     BadgeVariant = "primary"
	BadgeSecondary   BadgeVariant = "secondary"
	BadgeDestructive BadgeVariant = "destructive"
	BadgeOutline     BadgeVariant = "outline"
)

// BadgeProps defines all properties for a badge component
type BadgeProps struct {
	// Content
	Label string // Badge text content
	Icon  string // Optional icon name
	
	// Styling
	Variant   BadgeVariant // Visual style variant
	ClassName string       // Additional CSS classes
	
	// Link behavior (if badge should be clickable)
	Href     string // URL for link badges
	OnClick  string // JavaScript click handler
	HxPost   string // HTMX post URL
	HxGet    string // HTMX get URL
	HxTarget string // HTMX target selector
	HxSwap   string // HTMX swap strategy
	
	// Accessibility
	AriaLabel       string // aria-label for accessibility
	AriaDescribedBy string // aria-describedby reference
	Role            string // ARIA role override
	TabIndex        string // tabindex for focusable badges
}

// GetBadgeClasses generates the appropriate CSS classes for a badge
func GetBadgeClasses(variant BadgeVariant) string {
	baseClass := "badge"
	
	if variant != BadgePrimary {
		return baseClass + "-" + string(variant)
	}
	
	return baseClass
}

// GetDefaultBadgeProps returns badge props with sensible defaults
func GetDefaultBadgeProps() BadgeProps {
	return BadgeProps{
		Variant:   BadgePrimary,
		ClassName: "",
	}
}

// WithVariant returns a copy of props with the specified variant
func (p BadgeProps) WithVariant(variant BadgeVariant) BadgeProps {
	p.Variant = variant
	return p
}

// WithIcon returns a copy of props with an icon
func (p BadgeProps) WithIcon(iconName string) BadgeProps {
	p.Icon = iconName
	return p
}

// WithLink returns a copy of props configured as a clickable link
func (p BadgeProps) WithLink(href string) BadgeProps {
	p.Href = href
	return p
}

// WithHtmx returns a copy of props with HTMX configuration
func (p BadgeProps) WithHtmx(method, url, target, swap string) BadgeProps {
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

// IsClickable returns true if the badge has click behavior
func (p BadgeProps) IsClickable() bool {
	return p.Href != "" || p.OnClick != "" || p.HxPost != "" || p.HxGet != ""
}

// HasIcon returns true if the badge has an icon
func (p BadgeProps) HasIcon() bool {
	return p.Icon != ""
}

// GetElementTag returns the appropriate HTML element tag for the badge
func (p BadgeProps) GetElementTag() string {
	if p.Href != "" {
		return "a"
	}
	if p.IsClickable() {
		return "button"
	}
	return "span"
}

// GetAriaLabel returns the appropriate aria-label value
func (p BadgeProps) GetAriaLabel() string {
	if p.AriaLabel != "" {
		return p.AriaLabel
	}
	return ""
}