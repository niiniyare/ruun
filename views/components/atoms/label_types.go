package atoms

import (
	"strings"
)

// LabelProps defines all properties for a label component
type LabelProps struct {
	// Content
	Text string // Label text content
	
	// Associations
	For string // Associated form control ID (HTML for attribute)
	
	// State indicators
	Required bool // Shows required indicator (*)
	Optional bool // Shows optional indicator
	
	// Styling
	ClassName string // Additional CSS classes
	
	// Accessibility
	ID              string // HTML id attribute
	AriaLabel       string // aria-label for accessibility
	AriaDescribedBy string // aria-describedby reference
	
	// Help text
	HelpText    string // Additional help text
	HelpTextID  string // ID for help text element
}

// GetLabelClasses generates the appropriate CSS classes for a label
func GetLabelClasses() string {
	return "label"
}

// GetDefaultLabelProps returns label props with sensible defaults
func GetDefaultLabelProps() LabelProps {
	return LabelProps{
		Required:  false,
		Optional:  false,
		ClassName: "",
	}
}

// WithText returns a copy of props with the specified text
func (p LabelProps) WithText(text string) LabelProps {
	p.Text = text
	return p
}

// WithFor returns a copy of props with the specified form control association
func (p LabelProps) WithFor(controlID string) LabelProps {
	p.For = controlID
	return p
}

// WithRequired returns a copy of props marked as required
func (p LabelProps) WithRequired(required bool) LabelProps {
	p.Required = required
	// Clear optional if setting required
	if required {
		p.Optional = false
	}
	return p
}

// WithOptional returns a copy of props marked as optional
func (p LabelProps) WithOptional(optional bool) LabelProps {
	p.Optional = optional
	// Clear required if setting optional
	if optional {
		p.Required = false
	}
	return p
}

// WithHelpText returns a copy of props with help text
func (p LabelProps) WithHelpText(helpText string) LabelProps {
	p.HelpText = helpText
	// Generate help text ID if not provided
	if p.HelpTextID == "" && p.For != "" {
		p.HelpTextID = p.For + "-help"
	}
	return p
}

// HasHelpText returns true if the label has help text
func (p LabelProps) HasHelpText() bool {
	return p.HelpText != ""
}

// GetRequiredIndicator returns the required indicator text
func (p LabelProps) GetRequiredIndicator() string {
	if p.Required {
		return "*"
	}
	return ""
}

// GetOptionalIndicator returns the optional indicator text
func (p LabelProps) GetOptionalIndicator() string {
	if p.Optional {
		return "(optional)"
	}
	return ""
}

// GetDisplayText returns the complete display text with indicators
func (p LabelProps) GetDisplayText() string {
	parts := []string{p.Text}
	
	if p.Required {
		parts = append(parts, "*")
	} else if p.Optional {
		parts = append(parts, "(optional)")
	}
	
	return strings.Join(parts, " ")
}

// GetAriaLabel returns the appropriate aria-label value
func (p LabelProps) GetAriaLabel() string {
	if p.AriaLabel != "" {
		return p.AriaLabel
	}
	return ""
}

// GetHelpTextID returns the ID for the help text element
func (p LabelProps) GetHelpTextID() string {
	if p.HelpTextID != "" {
		return p.HelpTextID
	}
	if p.For != "" && p.HasHelpText() {
		return p.For + "-help"
	}
	return ""
}