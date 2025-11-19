package molecules

import (
	"fmt"

	"github.com/niiniyare/ruun/views/components//utils"
	"github.com/niiniyare/ruun/views/components/atoms"
)

// FormField validation state constants
const (
	ValidationStateIdle       = "idle"
	ValidationStateValidating = "validating"
	ValidationStateValid      = "valid"
	ValidationStateInvalid    = "invalid"
	ValidationStateWarning    = "warning"
)

// FormFieldProps represents the molecule-level form field component
// This integrates atoms with schema field definitions and provides
// comprehensive form functionality with compiled theme classes
type FormFieldProps struct {
	// Basic properties
	ID          string `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Placeholder string `json:"placeholder,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	Readonly    bool   `json:"readonly,omitempty"`

	// Enhanced display
	Description string `json:"description,omitempty"` // Rich text description
	HelpText    string `json:"helpText,omitempty"`    // Helper text below field
	Tooltip     string `json:"tooltip,omitempty"`     // Tooltip on hover
	Icon        string `json:"icon,omitempty"`        // Leading icon

	// Validation state integration
	ValidationState    string   `json:"validationState,omitempty"`    // idle, validating, valid, invalid, warning
	ValidationLoading  bool     `json:"validationLoading,omitempty"`  // Async validation in progress
	ValidationMessages []string `json:"validationMessages,omitempty"` // Current validation messages
	OnValidate         string   `json:"onValidate,omitempty"`         // HTMX validation endpoint
	ValidationDebounce int      `json:"validationDebounce,omitempty"` // Debounce timing in ms

	// Molecule-specific styling
	Size       string `json:"size,omitempty"`       // sm, md, lg, xl
	Variant    string `json:"variant,omitempty"`    // default, outline, filled, ghost
	FullWidth  bool   `json:"fullWidth,omitempty"`  // Expand to container width
	ClassName  string `json:"className,omitempty"`  // Additional CSS classes
	LabelClass string `json:"labelClass,omitempty"` // Label-specific classes
	InputClass string `json:"inputClass,omitempty"` // Input-specific classes

	// Options for select/radio/checkbox types
	Options []SelectOption `json:"options,omitempty"` // From molecules/form.go

	// Accessibility
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`

	// Framework integration
	HTMXProps   map[string]string `json:"htmxProps,omitempty"`   // HTMX attributes
	AlpineProps map[string]string `json:"alpineProps,omitempty"` // Alpine.js bindings
}

// SelectOption represents an option for select fields (from form.go)
type SelectOption struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Disabled    bool   `json:"disabled,omitempty"`
	Selected    bool   `json:"selected,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Group       string `json:"group,omitempty"`
	Meta        string `json:"meta,omitempty"`
}

// getFormFieldClasses returns compiled theme classes for the form field
func getFormFieldClasses(props FormFieldProps) string {
	return utils.TwMerge(
		"form-field",
		fmt.Sprintf("form-field-%s", utils.If(props.Size != "", props.Size, "md")),
		fmt.Sprintf("form-field-%s", utils.If(props.Variant != "", props.Variant, "default")),
		utils.If(props.FullWidth, "form-field-full-width"),
		utils.If(props.Required, "form-field-required"),
		utils.If(props.Disabled, "form-field-disabled"),
		utils.If(props.Readonly, "form-field-readonly"),
		utils.If(props.ValidationLoading, "form-field-validating"),
		getValidationStateClasses(props.ValidationState),
		props.ClassName,
	)
}

// getLabelClasses returns compiled theme classes for the label
func getLabelClasses(props FormFieldProps) string {
	return utils.TwMerge(
		"form-field-label",
		fmt.Sprintf("form-field-label-%s", utils.If(props.Size != "", props.Size, "md")),
		utils.If(props.Required, "form-field-label-required"),
		utils.If(props.Disabled, "form-field-label-disabled"),
		props.LabelClass,
	)
}

// getValidationStateClasses returns compiled theme classes for validation state
func getValidationStateClasses(state string) string {
	switch state {
	case ValidationStateValidating:
		return "form-field-state-validating"
	case ValidationStateValid:
		return "form-field-state-valid"
	case ValidationStateInvalid:
		return "form-field-state-invalid"
	case ValidationStateWarning:
		return "form-field-state-warning"
	default:
		return "form-field-state-idle"
	}
}

// getValidationMessageClasses returns compiled theme classes for validation messages
func getValidationMessageClasses(state string) string {
	switch state {
	case ValidationStateInvalid:
		return "validation-message validation-message-error"
	case ValidationStateWarning:
		return "validation-message validation-message-warning"
	case ValidationStateValid:
		return "validation-message validation-message-success"
	default:
		return "validation-message validation-message-info"
	}
}

// buildInputProps creates input props from FormField props
func buildInputProps(props FormFieldProps) atoms.InputProps {
	// Determine input state based on validation
	inputState := atoms.InputStateDefault
	switch props.ValidationState {
	case ValidationStateInvalid:
		inputState = atoms.InputStateError
	case ValidationStateValid:
		inputState = atoms.InputStateSuccess
	case ValidationStateWarning:
		inputState = atoms.InputStateWarning
	}

	// Convert size
	inputSize := atoms.InputSizeMD
	switch props.Size {
	case "sm":
		inputSize = atoms.InputSizeSM
	case "lg":
		inputSize = atoms.InputSizeLG
	case "xl":
		inputSize = atoms.InputSizeXL
	}

	// Convert variant
	inputVariant := atoms.InputVariantDefault
	switch props.Variant {
	case "outline":
		inputVariant = atoms.InputVariantOutline
	case "filled":
		inputVariant = atoms.InputVariantFilled
	case "ghost":
		inputVariant = atoms.InputVariantGhost
	}

	return atoms.InputProps{
		ID:          props.ID,
		Name:        props.Name,
		Type:        props.Type,
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Size:        inputSize,
		Variant:     inputVariant,
		State:       inputState,
		Required:    props.Required,
		Disabled:    props.Disabled,
		Readonly:    props.Readonly,
		FullWidth:   props.FullWidth,
		ClassName:   props.InputClass,
		Icon:        props.Icon,
		AriaLabel:   props.AriaLabel,
	}
}

// buildCheckboxProps creates checkbox props from FormField props
func buildCheckboxProps(props FormFieldProps, option SelectOption) atoms.CheckboxProps {
	checkboxSize := atoms.CheckboxSizeMD
	switch props.Size {
	case "sm":
		checkboxSize = atoms.CheckboxSizeSM
	case "lg":
		checkboxSize = atoms.CheckboxSizeLG
	}

	checkboxState := atoms.CheckboxStateDefault
	if props.ValidationState == ValidationStateInvalid {
		checkboxState = atoms.CheckboxStateError
	}

	return atoms.CheckboxProps{
		ID:       fmt.Sprintf("%s-%s", props.ID, option.Value),
		Name:     props.Name,
		Value:    option.Value,
		Label:    option.Label,
		Checked:  option.Selected,
		Disabled: props.Disabled || option.Disabled,
		Size:     checkboxSize,
		State:    checkboxState,
		Icon:     option.Icon,
	}
}

// buildRadioProps creates radio props from FormField props
func buildRadioProps(props FormFieldProps, option SelectOption) atoms.RadioProps {
	radioSize := atoms.RadioSizeMD
	switch props.Size {
	case "sm":
		radioSize = atoms.RadioSizeSM
	case "lg":
		radioSize = atoms.RadioSizeLG
	}

	radioState := atoms.RadioStateDefault
	if props.ValidationState == ValidationStateInvalid {
		radioState = atoms.RadioStateError
	}

	return atoms.RadioProps{
		ID:       fmt.Sprintf("%s-%s", props.ID, option.Value),
		Name:     props.Name,
		Value:    option.Value,
		Label:    option.Label,
		Selected: option.Selected,
		Disabled: props.Disabled || option.Disabled,
		Size:     radioSize,
		State:    radioState,
		Icon:     option.Icon,
	}
}
