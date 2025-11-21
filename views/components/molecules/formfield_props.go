package molecules

import (
	"fmt"
	"github.com/niiniyare/ruun/pkg/utils"
	"github.com/niiniyare/ruun/views/components/atoms"
)

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
		fmt.Sprintf("form-field-%s", string(props.Size)),
		fmt.Sprintf("form-field-%s", string(props.Variant)),
		utils.If(props.FullWidth, "form-field-full-width"),
		utils.If(props.Required, "form-field-required"),
		utils.If(props.Disabled, "form-field-disabled"),
		utils.If(props.Readonly, "form-field-readonly"),
		utils.If(props.ValidationLoading, "form-field-validating"),
		getValidationStateClasses(string(props.ValidationState)),
		props.Class,
	)
}

// getLabelClasses returns compiled theme classes for the label
func getLabelClasses(props FormFieldProps) string {
	return utils.TwMerge(
		"form-field-label",
		fmt.Sprintf("form-field-label-%s", string(props.Size)),
		utils.If(props.Required, "form-field-label-required"),
		utils.If(props.Disabled, "form-field-label-disabled"),
		props.LabelClass,
	)
}

// getValidationStateClasses returns compiled theme classes for validation state
func getValidationStateClasses(state string) string {
	switch state {
	case "validating":
		return "form-field-state-validating"
	case "valid":
		return "form-field-state-valid"
	case "invalid":
		return "form-field-state-invalid"
	case "warning":
		return "form-field-state-warning"
	default:
		return "form-field-state-idle"
	}
}

// getFormFieldValidationMessageClasses returns compiled theme classes for validation messages
func getFormFieldValidationMessageClasses(state string) string {
	switch state {
	case "invalid", "error":
		return "validation-message validation-message-error"
	case "warning":
		return "validation-message validation-message-warning"
	case "valid", "success":
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
	case FormFieldSizeSM:
		inputSize = atoms.InputSizeSM
	case FormFieldSizeLG:
		inputSize = atoms.InputSizeLG
	case FormFieldSizeXL:
		inputSize = atoms.InputSizeXL
	}

	return atoms.InputProps{
		ID:          props.ID,
		Name:        props.Name,
		Type:        atoms.InputType(string(props.Type)),
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Size:        inputSize,
		State:       inputState,
		Required:    props.Required,
		Disabled:    props.Disabled,
		Readonly:    props.Readonly,
		ClassName:   utils.TwMerge(utils.If(props.FullWidth, "w-full"), props.InputClass),
	}
}

// buildCheckboxProps creates checkbox props from FormField props
func buildCheckboxProps(props FormFieldProps, option SelectOption) atoms.CheckboxProps {
	checkboxSize := atoms.CheckboxSizeMD
	switch props.Size {
	case FormFieldSizeSM:
		checkboxSize = atoms.CheckboxSizeSM
	case FormFieldSizeLG:
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

// buildRadioGroupProps creates radio props from FormField props
func buildRadioGroupProps(props FormFieldProps) atoms.RadioProps {
	radioSize := atoms.RadioSizeMD
	switch props.Size {
	case FormFieldSizeSM:
		radioSize = atoms.RadioSizeSM
	case FormFieldSizeLG:
		radioSize = atoms.RadioSizeLG
	}

	radioState := atoms.RadioStateDefault
	if props.ValidationState == ValidationStateInvalid {
		radioState = atoms.RadioStateError
	}

	// Convert SelectOptions to RadioOptions - need to get options from props
	var radioOptions []atoms.RadioOption
	for _, option := range props.Options {
		radioOptions = append(radioOptions, atoms.RadioOption{
			Value:       option.Value,
			Label:       option.Label,
			Description: option.Description,
			Disabled:    option.Disabled,
			Icon:        option.Icon,
		})
	}

	return atoms.RadioProps{
		ID:       props.ID,
		Name:     props.Name,
		Value:    props.Value,
		Options:  radioOptions,
		Size:     radioSize,
		State:    radioState,
		Required: props.Required,
		Disabled: props.Disabled,
		Readonly: props.Readonly,
		Class:    props.InputClass,
	}
}