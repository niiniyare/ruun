package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/validation"
)

// DefaultFieldRenderer provides basic HTML rendering for demo purposes
type DefaultFieldRenderer struct{}

// SupportsType implements ComponentRenderer interface
func (r *DefaultFieldRenderer) SupportsType(fieldType schema.FieldType) bool {
	// DefaultFieldRenderer supports all basic field types
	supportedTypes := []schema.FieldType{
		schema.FieldText,
		schema.FieldEmail,
		schema.FieldNumber,
		schema.FieldSelect,
		schema.FieldCheckbox,
		schema.FieldTextarea,
	}

	for _, supported := range supportedTypes {
		if fieldType == supported {
			return true
		}
	}
	return false
}

// Render implements FieldRenderer interface
func (r *DefaultFieldRenderer) Render(ctx context.Context, props molecules.FormFieldProps) (string, error) {
	var inputHTML string
	var containerClasses []string

	// Add validation state classes
	containerClasses = append(containerClasses, "form-field")
	if props.ValidationState != validation.ValidationStateIdle {
		containerClasses = append(containerClasses, fmt.Sprintf("validation-state-%s", props.ValidationState.String()))
	}

	// Generate input based on type
	switch props.Type {
	case "text", "email":
		inputHTML = r.renderTextInput(props)
	case "number":
		inputHTML = r.renderNumberInput(props)
	case "select":
		inputHTML = r.renderSelectInput(props)
	case "checkbox":
		inputHTML = r.renderCheckboxInput(props)
	case "textarea":
		inputHTML = r.renderTextareaInput(props)
	default:
		inputHTML = r.renderTextInput(props)
	}

	// Render error messages
	errorHTML := ""
	if props.ErrorText != "" {
		errorHTML = fmt.Sprintf(`<div class="field-error">%s</div>`, props.ErrorText)
	}

	// Render help text
	helpHTML := ""
	if props.HelpText != "" {
		helpHTML = fmt.Sprintf(`<div class="field-help">%s</div>`, props.HelpText)
	}

	// Special handling for checkbox
	if props.Type == "checkbox" {
		return fmt.Sprintf(`
			<div class="%s" data-field="%s">
				%s
				%s
				%s
				<div class="validation-indicator"></div>
			</div>
		`, strings.Join(containerClasses, " "), props.Name, inputHTML, helpHTML, errorHTML), nil
	}

	// Regular field layout
	requiredClass := ""
	if props.Required {
		requiredClass = "required"
	}

	return fmt.Sprintf(`
		<div class="%s" data-field="%s">
			<label for="%s" class="field-label %s">%s</label>
			<div class="field-container">
				%s
				<div class="validation-indicator"></div>
			</div>
			%s
			%s
		</div>
	`, strings.Join(containerClasses, " "), props.Name, props.ID, requiredClass,
		props.Label, inputHTML, helpHTML, errorHTML), nil
}

func (r *DefaultFieldRenderer) renderTextInput(props molecules.FormFieldProps) string {
	inputType := "text"
	if props.Type == "email" {
		inputType = "email"
	}

	classes := []string{"field-input"}
	if props.ValidationState == validation.ValidationStateValid {
		classes = append(classes, "is-valid")
	} else if props.ValidationState == validation.ValidationStateInvalid {
		classes = append(classes, "is-invalid")
	}

	attributes := []string{
		fmt.Sprintf(`type="%s"`, inputType),
		fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")),
		fmt.Sprintf(`id="%s"`, props.ID),
		fmt.Sprintf(`name="%s"`, props.Name),
		fmt.Sprintf(`value="%s"`, props.Value),
	}

	if props.Placeholder != "" {
		attributes = append(attributes, fmt.Sprintf(`placeholder="%s"`, props.Placeholder))
	}

	if props.Required {
		attributes = append(attributes, "required")
	}

	if props.Disabled {
		attributes = append(attributes, "disabled")
	}

	if props.Readonly {
		attributes = append(attributes, "readonly")
	}

	// Add HTMX attributes for real-time validation
	attributes = append(attributes,
		`hx-post="/validate-field"`,
		`hx-trigger="input changed delay:500ms, blur"`,
		fmt.Sprintf(`hx-vals='{"field": "%s"}'`, props.Name),
		fmt.Sprintf(`hx-target="[data-field='%s']"`, props.Name),
		`hx-swap="outerHTML"`,
	)

	return fmt.Sprintf(`<input %s />`, strings.Join(attributes, " "))
}

func (r *DefaultFieldRenderer) renderNumberInput(props molecules.FormFieldProps) string {
	classes := []string{"field-input"}
	if props.ValidationState == validation.ValidationStateValid {
		classes = append(classes, "is-valid")
	} else if props.ValidationState == validation.ValidationStateInvalid {
		classes = append(classes, "is-invalid")
	}

	attributes := []string{
		`type="number"`,
		fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")),
		fmt.Sprintf(`id="%s"`, props.ID),
		fmt.Sprintf(`name="%s"`, props.Name),
		fmt.Sprintf(`value="%s"`, props.Value),
	}

	if props.Placeholder != "" {
		attributes = append(attributes, fmt.Sprintf(`placeholder="%s"`, props.Placeholder))
	}

	if props.Required {
		attributes = append(attributes, "required")
	}

	// Add HTMX attributes for real-time validation
	attributes = append(attributes,
		`hx-post="/validate-field"`,
		`hx-trigger="input changed delay:500ms, blur"`,
		fmt.Sprintf(`hx-vals='{"field": "%s"}'`, props.Name),
		fmt.Sprintf(`hx-target="[data-field='%s']"`, props.Name),
		`hx-swap="outerHTML"`,
	)

	return fmt.Sprintf(`<input %s />`, strings.Join(attributes, " "))
}

func (r *DefaultFieldRenderer) renderSelectInput(props molecules.FormFieldProps) string {
	classes := []string{"field-input", "field-select"}
	if props.ValidationState == validation.ValidationStateValid {
		classes = append(classes, "is-valid")
	} else if props.ValidationState == validation.ValidationStateInvalid {
		classes = append(classes, "is-invalid")
	}

	attributes := []string{
		fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")),
		fmt.Sprintf(`id="%s"`, props.ID),
		fmt.Sprintf(`name="%s"`, props.Name),
	}

	if props.Required {
		attributes = append(attributes, "required")
	}

	// Add HTMX attributes for real-time validation
	attributes = append(attributes,
		`hx-post="/validate-field"`,
		`hx-trigger="change"`,
		fmt.Sprintf(`hx-vals='{"field": "%s"}'`, props.Name),
		fmt.Sprintf(`hx-target="[data-field='%s']"`, props.Name),
		`hx-swap="outerHTML"`,
	)

	options := fmt.Sprintf(`<option value="">Select %s</option>`, props.Label)
	for _, option := range props.Options {
		selected := ""
		if option.Value == props.Value {
			selected = " selected"
		}
		options += fmt.Sprintf(`<option value="%s"%s>%s</option>`,
			option.Value, selected, option.Label)
	}

	return fmt.Sprintf(`<select %s>%s</select>`, strings.Join(attributes, " "), options)
}

func (r *DefaultFieldRenderer) renderCheckboxInput(props molecules.FormFieldProps) string {
	classes := []string{"field-checkbox"}

	attributes := []string{
		`type="checkbox"`,
		fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")),
		fmt.Sprintf(`id="%s"`, props.ID),
		fmt.Sprintf(`name="%s"`, props.Name),
		`value="true"`,
	}

	if props.Value == "true" || props.Value == "on" {
		attributes = append(attributes, "checked")
	}

	if props.Required {
		attributes = append(attributes, "required")
	}

	// Add HTMX attributes for real-time validation
	attributes = append(attributes,
		`hx-post="/validate-field"`,
		`hx-trigger="change"`,
		fmt.Sprintf(`hx-vals='{"field": "%s"}'`, props.Name),
		fmt.Sprintf(`hx-target="[data-field='%s']"`, props.Name),
		`hx-swap="outerHTML"`,
	)

	requiredIndicator := ""
	if props.Required {
		requiredIndicator = ` <span style="color: var(--semantic-feedback-error-default);">*</span>`
	}

	return fmt.Sprintf(`
		<div class="field-checkbox-container">
			<input %s />
			<label for="%s" class="field-label">%s%s</label>
		</div>
	`, strings.Join(attributes, " "), props.ID, props.Label, requiredIndicator)
}

func (r *DefaultFieldRenderer) renderTextareaInput(props molecules.FormFieldProps) string {
	classes := []string{"field-input"}
	if props.ValidationState == validation.ValidationStateValid {
		classes = append(classes, "is-valid")
	} else if props.ValidationState == validation.ValidationStateInvalid {
		classes = append(classes, "is-invalid")
	}

	attributes := []string{
		fmt.Sprintf(`class="%s"`, strings.Join(classes, " ")),
		fmt.Sprintf(`id="%s"`, props.ID),
		fmt.Sprintf(`name="%s"`, props.Name),
		`rows="4"`,
	}

	if props.Placeholder != "" {
		attributes = append(attributes, fmt.Sprintf(`placeholder="%s"`, props.Placeholder))
	}

	if props.Required {
		attributes = append(attributes, "required")
	}

	// Add HTMX attributes for real-time validation
	attributes = append(attributes,
		`hx-post="/validate-field"`,
		`hx-trigger="input changed delay:500ms, blur"`,
		fmt.Sprintf(`hx-vals='{"field": "%s"}'`, props.Name),
		fmt.Sprintf(`hx-target="[data-field='%s']"`, props.Name),
		`hx-swap="outerHTML"`,
	)

	return fmt.Sprintf(`<textarea %s>%s</textarea>`, strings.Join(attributes, " "), props.Value)
}
