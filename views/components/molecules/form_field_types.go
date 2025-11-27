package molecules

import (
	"fmt"
	"strings"
	"github.com/your-org/your-project/views/components/atoms"
)

// FormFieldProps defines all properties for a form field molecule
// Combines Label + Input + Help Text + Error Message
type FormFieldProps struct {
	// Field identification
	Name string // Field name (used for input name and ID generation)
	ID   string // Override ID (optional, defaults to name)
	
	// Label properties
	Label    string // Field label text
	Required bool   // Required field indicator
	Optional bool   // Optional field indicator
	
	// Input properties
	InputType    atoms.InputType // Type of input
	Value        string          // Current value
	Placeholder  string          // Placeholder text
	Disabled     bool            // Disabled state
	ReadOnly     bool            // Read-only state
	
	// Validation constraints
	MinLength int    // Minimum length
	MaxLength int    // Maximum length
	Min       string // Minimum value (numbers/dates)
	Max       string // Maximum value (numbers/dates)
	Step      string // Step value (numbers)
	Pattern   string // Validation pattern
	
	// Help and validation
	HelpText     string   // Help text below the field
	ErrorMessage string   // Error message
	ShowError    bool     // Whether to show error state
	Warnings     []string // Warning messages (non-blocking)
	
	// Layout
	Orientation string // "vertical" (default) or "horizontal"
	ClassName   string // Additional CSS classes
	
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
	
	// Accessibility
	AriaLabel       string // aria-label override
	AriaDescribedBy string // Additional aria-describedby references
}

// GetFormFieldClasses generates CSS classes for the form field container
func GetFormFieldClasses(orientation string, hasError bool) string {
	var classes []string
	
	classes = append(classes, "field")
	
	if orientation == "horizontal" {
		classes = append(classes, "field-horizontal")
	}
	
	if hasError {
		classes = append(classes, "field-error")
	}
	
	return strings.Join(classes, " ")
}

// GetDefaultFormFieldProps returns form field props with sensible defaults
func GetDefaultFormFieldProps() FormFieldProps {
	return FormFieldProps{
		InputType:   atoms.InputTypeText,
		Required:    false,
		Optional:    false,
		Disabled:    false,
		ReadOnly:    false,
		ShowError:   false,
		Orientation: "vertical",
		ClassName:   "",
		Warnings:    []string{},
	}
}

// WithName returns a copy of props with the specified name
func (p FormFieldProps) WithName(name string) FormFieldProps {
	p.Name = name
	return p
}

// WithLabel returns a copy of props with the specified label
func (p FormFieldProps) WithLabel(label string) FormFieldProps {
	p.Label = label
	return p
}

// WithInputType returns a copy of props with the specified input type
func (p FormFieldProps) WithInputType(inputType atoms.InputType) FormFieldProps {
	p.InputType = inputType
	return p
}

// WithValue returns a copy of props with the specified value
func (p FormFieldProps) WithValue(value string) FormFieldProps {
	p.Value = value
	return p
}

// WithPlaceholder returns a copy of props with the specified placeholder
func (p FormFieldProps) WithPlaceholder(placeholder string) FormFieldProps {
	p.Placeholder = placeholder
	return p
}

// WithRequired returns a copy of props marked as required
func (p FormFieldProps) WithRequired(required bool) FormFieldProps {
	p.Required = required
	if required {
		p.Optional = false
	}
	return p
}

// WithOptional returns a copy of props marked as optional
func (p FormFieldProps) WithOptional(optional bool) FormFieldProps {
	p.Optional = optional
	if optional {
		p.Required = false
	}
	return p
}

// WithHelpText returns a copy of props with help text
func (p FormFieldProps) WithHelpText(helpText string) FormFieldProps {
	p.HelpText = helpText
	return p
}

// WithError returns a copy of props with error configuration
func (p FormFieldProps) WithError(errorMessage string) FormFieldProps {
	p.ErrorMessage = errorMessage
	p.ShowError = errorMessage != ""
	return p
}

// WithWarning returns a copy of props with a warning message added
func (p FormFieldProps) WithWarning(warning string) FormFieldProps {
	if warning != "" {
		p.Warnings = append(p.Warnings, warning)
	}
	return p
}

// WithConstraints returns a copy of props with validation constraints
func (p FormFieldProps) WithConstraints(minLength, maxLength int, pattern string) FormFieldProps {
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

// WithNumberConstraints returns a copy of props with number constraints
func (p FormFieldProps) WithNumberConstraints(min, max, step string) FormFieldProps {
	p.Min = min
	p.Max = max
	p.Step = step
	return p
}

// WithHorizontalLayout returns a copy of props with horizontal layout
func (p FormFieldProps) WithHorizontalLayout() FormFieldProps {
	p.Orientation = "horizontal"
	return p
}

// WithHtmx returns a copy of props with HTMX configuration
func (p FormFieldProps) WithHtmx(method, url, target, trigger, swap string) FormFieldProps {
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

// GetFieldID returns the effective field ID
func (p FormFieldProps) GetFieldID() string {
	if p.ID != "" {
		return p.ID
	}
	return p.Name
}

// GetHelpTextID returns the ID for the help text element
func (p FormFieldProps) GetHelpTextID() string {
	if p.HelpText != "" {
		return p.GetFieldID() + "-help"
	}
	return ""
}

// GetErrorID returns the ID for the error message element
func (p FormFieldProps) GetErrorID() string {
	if p.ShowError && p.ErrorMessage != "" {
		return p.GetFieldID() + "-error"
	}
	return ""
}

// GetWarningIDs returns IDs for warning message elements
func (p FormFieldProps) GetWarningIDs() []string {
	if len(p.Warnings) == 0 {
		return []string{}
	}
	
	var ids []string
	baseID := p.GetFieldID()
	for i := range p.Warnings {
		ids = append(ids, fmt.Sprintf("%s-warning-%d", baseID, i))
	}
	return ids
}

// GetAriaDescribedBy returns the complete aria-describedby value
func (p FormFieldProps) GetAriaDescribedBy() string {
	var refs []string
	
	// Add custom references
	if p.AriaDescribedBy != "" {
		refs = append(refs, p.AriaDescribedBy)
	}
	
	// Add help text reference
	if helpID := p.GetHelpTextID(); helpID != "" {
		refs = append(refs, helpID)
	}
	
	// Add error reference
	if errorID := p.GetErrorID(); errorID != "" {
		refs = append(refs, errorID)
	}
	
	// Add warning references
	refs = append(refs, p.GetWarningIDs()...)
	
	return strings.Join(refs, " ")
}

// ToLabelProps converts field props to label props
func (p FormFieldProps) ToLabelProps() atoms.LabelProps {
	labelProps := atoms.GetDefaultLabelProps()
	labelProps.Text = p.Label
	labelProps.For = p.GetFieldID()
	labelProps.Required = p.Required
	labelProps.Optional = p.Optional
	
	if p.HelpText != "" {
		labelProps = labelProps.WithHelpText(p.HelpText)
	}
	
	return labelProps
}

// ToInputProps converts field props to input props
func (p FormFieldProps) ToInputProps() atoms.InputProps {
	inputProps := atoms.GetDefaultInputProps()
	inputProps.Name = p.Name
	inputProps.ID = p.GetFieldID()
	inputProps.Type = p.InputType
	inputProps.Value = p.Value
	inputProps.Placeholder = p.Placeholder
	inputProps.Required = p.Required
	inputProps.Disabled = p.Disabled
	inputProps.ReadOnly = p.ReadOnly
	inputProps.Error = p.ShowError
	inputProps.ErrorMessage = p.ErrorMessage
	inputProps.AriaInvalid = p.ShowError
	
	// Set constraints
	inputProps.MinLength = p.MinLength
	inputProps.MaxLength = p.MaxLength
	inputProps.Min = p.Min
	inputProps.Max = p.Max
	inputProps.Step = p.Step
	inputProps.Pattern = p.Pattern
	
	// Set event handlers
	inputProps.OnChange = p.OnChange
	inputProps.OnInput = p.OnInput
	inputProps.OnFocus = p.OnFocus
	inputProps.OnBlur = p.OnBlur
	
	// Set HTMX
	inputProps.HxPost = p.HxPost
	inputProps.HxGet = p.HxGet
	inputProps.HxTarget = p.HxTarget
	inputProps.HxTrigger = p.HxTrigger
	inputProps.HxSwap = p.HxSwap
	
	// Set aria-describedby
	inputProps.AriaDescribedBy = p.GetAriaDescribedBy()
	
	// Override aria-label if provided
	if p.AriaLabel != "" {
		inputProps.AriaLabel = p.AriaLabel
	}
	
	return inputProps
}

// HasHelpText returns true if the field has help text
func (p FormFieldProps) HasHelpText() bool {
	return p.HelpText != ""
}

// HasError returns true if the field has an error
func (p FormFieldProps) HasError() bool {
	return p.ShowError
}

// HasWarnings returns true if the field has warnings
func (p FormFieldProps) HasWarnings() bool {
	return len(p.Warnings) > 0
}

// IsHorizontal returns true if the field uses horizontal layout
func (p FormFieldProps) IsHorizontal() bool {
	return p.Orientation == "horizontal"
}

// Validate checks if the form field props are valid
func (p FormFieldProps) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("form field must have a Name")
	}
	
	if p.Label == "" {
		return fmt.Errorf("form field must have a Label")
	}
	
	// Validate the input props
	inputProps := p.ToInputProps()
	if err := inputProps.Validate(); err != nil {
		return fmt.Errorf("invalid input props: %w", err)
	}
	
	return nil
}