package atoms

import (
	"fmt"
	"strings"
)

// InputType defines the HTML input type
type InputType string

const (
	InputTypeText         InputType = "text"
	InputTypeEmail        InputType = "email"
	InputTypePassword     InputType = "password"
	InputTypeNumber       InputType = "number"
	InputTypeFile         InputType = "file"
	InputTypeTel          InputType = "tel"
	InputTypeUrl          InputType = "url"
	InputTypeSearch       InputType = "search"
	InputTypeDate         InputType = "date"
	InputTypeDatetimeLocal InputType = "datetime-local"
	InputTypeMonth        InputType = "month"
	InputTypeWeek         InputType = "week"
	InputTypeTime         InputType = "time"
)

// InputProps defines all properties for an input component
type InputProps struct {
	// Core attributes
	Name        string    // Input name attribute
	ID          string    // Input id attribute
	Type        InputType // Input type
	Value       string    // Current value
	Placeholder string    // Placeholder text
	
	// Validation & State
	Required     bool   // Required field
	Disabled     bool   // Disabled state
	ReadOnly     bool   // Read-only state
	Error        bool   // Error state
	ErrorMessage string // Error message text
	
	// Size & Layout
	ClassName string // Additional CSS classes
	
	// Constraints
	MinLength int    // Minimum length (text inputs)
	MaxLength int    // Maximum length (text inputs)
	Min       string // Minimum value (number/date inputs)
	Max       string // Maximum value (number/date inputs)
	Step      string // Step value (number inputs)
	Pattern   string // Validation pattern (regex)
	
	// Accessibility
	AriaLabel       string // aria-label for accessibility
	AriaDescribedBy string // aria-describedby reference
	AriaInvalid     bool   // aria-invalid state
	
	// Interactivity
	OnChange string // JavaScript change handler
	OnInput  string // JavaScript input handler
	OnFocus  string // JavaScript focus handler
	OnBlur   string // JavaScript blur handler
	
	// HTMX
	HxPost    string // HTMX post URL
	HxGet     string // HTMX get URL
	HxTarget  string // HTMX target selector
	HxTrigger string // HTMX trigger event
	HxSwap    string // HTMX swap strategy
}

// GetInputClasses generates the appropriate CSS classes for an input
func GetInputClasses(inputType InputType, hasError, disabled, readOnly bool) string {
	var classes []string
	
	// Base class - all input types use .input class from Basecoat
	classes = append(classes, "input")
	
	// State classes
	if hasError {
		classes = append(classes, "error")
	}
	if disabled {
		classes = append(classes, "disabled")
	}
	if readOnly {
		classes = append(classes, "readonly")
	}
	
	return strings.Join(classes, " ")
}

// GetDefaultInputProps returns input props with sensible defaults
func GetDefaultInputProps() InputProps {
	return InputProps{
		Type:        InputTypeText,
		Required:    false,
		Disabled:    false,
		ReadOnly:    false,
		Error:       false,
		AriaInvalid: false,
		ClassName:   "",
	}
}

// WithType returns a copy of props with the specified input type
func (p InputProps) WithType(inputType InputType) InputProps {
	p.Type = inputType
	return p
}

// WithValue returns a copy of props with the specified value
func (p InputProps) WithValue(value string) InputProps {
	p.Value = value
	return p
}

// WithPlaceholder returns a copy of props with the specified placeholder
func (p InputProps) WithPlaceholder(placeholder string) InputProps {
	p.Placeholder = placeholder
	return p
}

// WithError returns a copy of props configured with error state
func (p InputProps) WithError(errorMessage string) InputProps {
	p.Error = true
	p.ErrorMessage = errorMessage
	p.AriaInvalid = true
	return p
}

// WithRequired returns a copy of props marked as required
func (p InputProps) WithRequired(required bool) InputProps {
	p.Required = required
	return p
}

// WithDisabled returns a copy of props with disabled state
func (p InputProps) WithDisabled(disabled bool) InputProps {
	p.Disabled = disabled
	return p
}

// WithConstraints returns a copy of props with validation constraints
func (p InputProps) WithConstraints(minLength, maxLength int, pattern string) InputProps {
	if minLength > 0 {
		p.MinLength = minLength
	}
	if maxLength > 0 {
		p.MaxLength = maxLength
	}
	if pattern != "" {
		p.Pattern = pattern
	}
	return p
}

// WithNumberConstraints returns a copy of props with number validation constraints
func (p InputProps) WithNumberConstraints(min, max, step string) InputProps {
	p.Min = min
	p.Max = max
	p.Step = step
	return p
}

// WithHtmx returns a copy of props with HTMX configuration
func (p InputProps) WithHtmx(method, url, target, trigger, swap string) InputProps {
	switch method {
	case "post":
		p.HxPost = url
	case "get":
		p.HxGet = url
	}
	p.HxTarget = target
	p.HxTrigger = trigger
	p.HxSwap = swap
	return p
}

// GetAriaDescribedBy returns the appropriate aria-describedby value
func (p InputProps) GetAriaDescribedBy() string {
	var refs []string
	
	if p.AriaDescribedBy != "" {
		refs = append(refs, p.AriaDescribedBy)
	}
	
	// Add error message reference if there's an error
	if p.Error && p.ErrorMessage != "" && p.ID != "" {
		refs = append(refs, p.ID+"-error")
	}
	
	return strings.Join(refs, " ")
}

// GetErrorID returns the ID for the error message element
func (p InputProps) GetErrorID() string {
	if p.ID != "" && p.Error && p.ErrorMessage != "" {
		return p.ID + "-error"
	}
	return ""
}

// Validate checks if the input props are valid
func (p InputProps) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("input must have a Name")
	}
	
	// Number inputs should have numeric constraints
	if p.Type == InputTypeNumber {
		// Min/Max validation could be added here if needed
	}
	
	// File inputs have different validation rules
	if p.Type == InputTypeFile {
		// File-specific validation could be added here
	}
	
	return nil
}

// IsNumericInput returns true if this is a numeric input type
func (p InputProps) IsNumericInput() bool {
	return p.Type == InputTypeNumber
}

// IsDateInput returns true if this is a date-related input type
func (p InputProps) IsDateInput() bool {
	return p.Type == InputTypeDate || 
		   p.Type == InputTypeDatetimeLocal || 
		   p.Type == InputTypeMonth || 
		   p.Type == InputTypeWeek || 
		   p.Type == InputTypeTime
}

// IsFileInput returns true if this is a file input
func (p InputProps) IsFileInput() bool {
	return p.Type == InputTypeFile
}