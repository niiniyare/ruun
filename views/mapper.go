package views

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// FieldMapper converts pure schema field definitions to FormField props
// This maintains separation between schema (business logic) and views (UI)
type FieldMapper struct {
	schema *schema.Schema
	locale string
}

// NewFieldMapper creates a new field mapper with schema context and locale
func NewFieldMapper(s *schema.Schema, locale string) *FieldMapper {
	return &FieldMapper{
		schema: s,
		locale: locale,
	}
}

// ConvertField converts a schema.Field to molecules.FormFieldProps
// This is the main conversion function that handles all field types
func (m *FieldMapper) ConvertField(
	field *schema.Field,
	value any,
	errors []string,
	touched, dirty bool,
) (molecules.FormFieldProps, error) {
	if field == nil {
		return molecules.FormFieldProps{}, fmt.Errorf("field cannot be nil")
	}

	// Start with basic props
	props := molecules.FormFieldProps{
		// Core identification
		Type:  m.mapFieldType(field.Type),
		ID:    m.generateFieldID(field.Name),
		Name:  field.Name,
		Label: m.getLocalizedText(field.Label, field.I18n, "label"),
		Value: m.convertValueToString(value),

		// Display & UX
		Placeholder: m.getLocalizedText(field.Placeholder, field.I18n, "placeholder"),
		HelpText:    m.getLocalizedText(field.Help, field.I18n, "help"),
		Description: m.getLocalizedText(field.Description, field.I18n, "description"),
		Tooltip:     m.getLocalizedText(field.Tooltip, field.I18n, "tooltip"),
		Icon:        field.Icon,
		ErrorText:   m.formatErrors(errors),

		// State
		Required: field.Required,
		Disabled: field.Disabled,
		Readonly: field.Readonly,

		// Validation from schema
		ValidateOnBlur:  true,  // Default to true for good UX
		ValidateOnInput: false, // Don't be too aggressive

		// Styling
		GenerateID: true,
		Locale:     m.locale,
	}

	// Map validation rules
	if err := m.mapValidationRules(field, &props); err != nil {
		return props, fmt.Errorf("failed to map validation rules: %w", err)
	}

	// Map field-type-specific properties
	if err := m.mapFieldTypeSpecific(field, &props); err != nil {
		return props, fmt.Errorf("failed to map field-specific properties: %w", err)
	}

	// Map options for selection fields
	if err := m.mapFieldOptions(field, &props); err != nil {
		return props, fmt.Errorf("failed to map field options: %w", err)
	}

	// Map HTMX attributes
	m.mapHTMXAttributes(field, &props)

	// Map Alpine.js attributes
	m.mapAlpineAttributes(field, &props)

	// Map styling and layout
	m.mapStyling(field, &props)

	// Apply defaults and normalization
	props.NormalizeDefaults()

	return props, nil
}

// ConvertAction converts a schema.Action to ButtonProps
func (m *FieldMapper) ConvertAction(action *schema.Action, enabled bool) ButtonProps {
	if action == nil {
		return ButtonProps{}
	}

	// Map action type to button type
	buttonType := "button"
	switch action.Type {
	case schema.ActionSubmit:
		buttonType = "submit"
	case schema.ActionReset:
		buttonType = "reset"
	case schema.ActionLink:
		buttonType = "button" // Links are still buttons in forms
	}

	// Map variant
	variant := "default"
	if action.Variant != "" {
		variant = action.Variant
	}

	return ButtonProps{
		Type:    buttonType,
		Text:    action.Text,
		Variant: variant,
		Enabled: enabled && !action.Disabled,
		ID:      action.ID,
		Class:   m.buildActionClasses(action),
	}
}

// Field type mapping
func (m *FieldMapper) mapFieldType(fieldType schema.FieldType) molecules.FormFieldType {
	mappings := map[schema.FieldType]molecules.FormFieldType{
		// Basic text inputs
		schema.FieldText:     molecules.FormFieldText,
		schema.FieldEmail:    molecules.FormFieldEmail,
		schema.FieldPassword: molecules.FormFieldPassword,
		schema.FieldNumber:   molecules.FormFieldNumber,
		schema.FieldPhone:    molecules.FormFieldTel,
		schema.FieldURL:      molecules.FormFieldURL,
		schema.FieldHidden:   molecules.FormFieldText, // Handle via props.Hidden

		// Date and time
		schema.FieldDate:      molecules.FormFieldDate,
		schema.FieldTime:      molecules.FormFieldTime,
		schema.FieldDateTime:  molecules.FormFieldDateTime,
		schema.FieldDateRange: molecules.FormFieldDateRange,

		// Text content
		schema.FieldTextarea: molecules.FormFieldTextarea,

		// Selection
		schema.FieldSelect:      molecules.FormFieldSelect,
		schema.FieldMultiSelect: molecules.FormFieldMultiSelect,
		schema.FieldRadio:       molecules.FormFieldRadio,
		schema.FieldCheckbox:    molecules.FormFieldCheckbox,
		schema.FieldCheckboxes:  molecules.FormFieldCheckboxGroup,

		// Specialized
		schema.FieldAutoComplete: molecules.FormFieldAutoComplete,
		schema.FieldColor:        molecules.FormFieldColor,
		schema.FieldFile:         molecules.FormFieldFile,
		schema.FieldSlider:       molecules.FormFieldRange,
		schema.FieldTags:         molecules.FormFieldTags,
	}

	if mapped, exists := mappings[fieldType]; exists {
		return mapped
	}

	// Default fallback
	return molecules.FormFieldText
}

// Validation rules mapping
func (m *FieldMapper) mapValidationRules(field *schema.Field, props *molecules.FormFieldProps) error {
	if field.Validation == nil {
		return nil
	}

	v := field.Validation

	// String constraints
	if v.MinLength != nil {
		props.MinLength = *v.MinLength
	}
	if v.MaxLength != nil {
		props.MaxLength = *v.MaxLength
	}
	if v.Pattern != "" {
		props.Pattern = v.Pattern
	}

	// Number constraints
	if v.Min != nil {
		props.Min = fmt.Sprintf("%.2f", *v.Min)
	}
	if v.Max != nil {
		props.Max = fmt.Sprintf("%.2f", *v.Max)
	}
	if v.Step != nil {
		props.Step = fmt.Sprintf("%.2f", *v.Step)
	}

	// Array constraints (for multi-select, checkboxes)
	if v.MinItems != nil && field.Type == schema.FieldMultiSelect {
		// FormFieldProps doesn't have MinItems directly, but we can use validation rules
		rule := molecules.ValidationRule{
			Type:    "minItems",
			Value:   *v.MinItems,
			Message: fmt.Sprintf("Select at least %d items", *v.MinItems),
		}
		if v.Messages != nil && v.Messages.Required != "" {
			rule.Message = v.Messages.Required
		}
		props.ValidationRules = append(props.ValidationRules, rule)
	}

	// File validation
	if v.File != nil {
		props.MaxFileSize = v.File.MaxSize
		if len(v.File.Accept) > 0 {
			props.Accept = strings.Join(v.File.Accept, ",")
		}
		props.MaxFiles = v.File.MaxFiles
	}

	// Custom validation rules
	if v.Custom != "" {
		props.ValidationRules = append(props.ValidationRules, molecules.ValidationRule{
			Type:    "custom",
			Value:   v.Custom,
			Message: "Custom validation failed",
		})
	}

	return nil
}

// Field-type-specific properties
func (m *FieldMapper) mapFieldTypeSpecific(field *schema.Field, props *molecules.FormFieldProps) error {
	switch field.Type {
	case schema.FieldTextarea:
		m.mapTextareaProps(field, props)
	case schema.FieldSelect, schema.FieldMultiSelect:
		m.mapSelectProps(field, props)
	case schema.FieldCheckboxes, schema.FieldRadio:
		m.mapGroupProps(field, props)
	case schema.FieldAutoComplete:
		m.mapAutocompleteProps(field, props)
	case schema.FieldFile, schema.FieldImage:
		m.mapFileProps(field, props)
	case schema.FieldDate, schema.FieldTime, schema.FieldDateTime:
		m.mapDateTimeProps(field, props)
	case schema.FieldSlider:
		m.mapSliderProps(field, props)
	case schema.FieldTags:
		m.mapTagsProps(field, props)
	case schema.FieldHidden:
		// Hidden fields should not be rendered normally
		props.Type = molecules.FormFieldText
	}

	return nil
}

// Textarea-specific properties
func (m *FieldMapper) mapTextareaProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["rows"]; ok {
		if rows, ok := config.(int); ok {
			props.Rows = rows
		} else if rowsStr, ok := config.(string); ok {
			if rows, err := strconv.Atoi(rowsStr); err == nil {
				props.Rows = rows
			}
		}
	}

	if config, ok := field.Config["autoResize"]; ok {
		if autoResize, ok := config.(bool); ok {
			props.AutoResize = autoResize
		}
	}

	if config, ok := field.Config["resizable"]; ok {
		if resizable, ok := config.(bool); ok {
			props.Resizable = resizable
		}
	}
}

// Select-specific properties
func (m *FieldMapper) mapSelectProps(field *schema.Field, props *molecules.FormFieldProps) {
	if field.Type == schema.FieldMultiSelect {
		props.Multiple = true
	}

	// Get config from field.Config
	if config, ok := field.Config["searchable"]; ok {
		if searchable, ok := config.(bool); ok {
			props.Searchable = searchable
		}
	}

	if config, ok := field.Config["clearable"]; ok {
		if clearable, ok := config.(bool); ok {
			props.Clearable = clearable
		}
	}

	if config, ok := field.Config["creatable"]; ok {
		if creatable, ok := config.(bool); ok {
			props.Creatable = creatable
		}
	}
}

// Group properties (radio, checkbox groups)
func (m *FieldMapper) mapGroupProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["inline"]; ok {
		if inline, ok := config.(bool); ok {
			props.Inline = inline
		}
	}

	if config, ok := field.Config["columns"]; ok {
		if columns, ok := config.(int); ok {
			props.Columns = columns
		}
	}
}

// Autocomplete-specific properties
func (m *FieldMapper) mapAutocompleteProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["searchURL"]; ok {
		if searchURL, ok := config.(string); ok {
			props.SearchURL = searchURL
		}
	}

	if config, ok := field.Config["minChars"]; ok {
		if minChars, ok := config.(int); ok {
			props.MinChars = minChars
		}
	}

	if config, ok := field.Config["maxResults"]; ok {
		if maxResults, ok := config.(int); ok {
			props.MaxResults = maxResults
		}
	}

	if config, ok := field.Config["debounce"]; ok {
		if debounce, ok := config.(int); ok {
			props.Debounce = debounce
		}
	}
}

// File upload properties
func (m *FieldMapper) mapFileProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["multiple"]; ok {
		if multiple, ok := config.(bool); ok {
			props.MultipleFiles = multiple
		}
	}

	if config, ok := field.Config["showPreview"]; ok {
		if showPreview, ok := config.(bool); ok {
			props.ShowPreview = showPreview
		}
	}

	if config, ok := field.Config["dropZone"]; ok {
		if dropZone, ok := config.(bool); ok {
			props.DropZone = dropZone
		}
	}
}

// Date/Time properties
func (m *FieldMapper) mapDateTimeProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["format24"]; ok {
		if format24, ok := config.(bool); ok {
			props.Format24 = format24
		}
	}

	if config, ok := field.Config["showCalendar"]; ok {
		if showCalendar, ok := config.(bool); ok {
			props.ShowCalendar = showCalendar
		}
	}

	if config, ok := field.Config["minDate"]; ok {
		if minDate, ok := config.(string); ok {
			props.MinDate = minDate
		}
	}

	if config, ok := field.Config["maxDate"]; ok {
		if maxDate, ok := config.(string); ok {
			props.MaxDate = maxDate
		}
	}
}

// Slider properties
func (m *FieldMapper) mapSliderProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["showValue"]; ok {
		if showValue, ok := config.(bool); ok {
			props.ShowValue = showValue
		}
	}

	if config, ok := field.Config["showMinMax"]; ok {
		if showMinMax, ok := config.(bool); ok {
			props.ShowMinMax = showMinMax
		}
	}
}

// Tags properties
func (m *FieldMapper) mapTagsProps(field *schema.Field, props *molecules.FormFieldProps) {
	if config, ok := field.Config["maxTags"]; ok {
		if maxTags, ok := config.(int); ok {
			props.MaxTags = maxTags
		}
	}

	if config, ok := field.Config["editable"]; ok {
		if editable, ok := config.(bool); ok {
			props.TagsEditable = editable
		}
	}
}

// Map field options for selection fields
func (m *FieldMapper) mapFieldOptions(field *schema.Field, props *molecules.FormFieldProps) error {
	if len(field.Options) == 0 {
		return nil
	}

	// Convert schema.Option to molecules.SelectOption
	options := make([]molecules.SelectOption, 0, len(field.Options))
	for _, opt := range field.Options {
		selectOpt := molecules.SelectOption{
			Value:       opt.Value,
			Label:       opt.Label,
			Description: opt.Description,
			Disabled:    opt.Disabled,
			Selected:    opt.Selected,
			Icon:        opt.Icon,
			Group:       opt.Group,
			Meta:        fmt.Sprintf("%v", opt.Meta),
		}
		options = append(options, selectOpt)
	}

	props.Options = options

	// For multi-select or checkbox groups, also populate Values array
	if field.Type == schema.FieldMultiSelect || field.Type == schema.FieldCheckboxes {
		var selectedValues []string
		for _, opt := range options {
			if opt.Selected {
				selectedValues = append(selectedValues, opt.Value)
			}
		}
		props.Values = selectedValues
	}

	return nil
}

// Map HTMX attributes
func (m *FieldMapper) mapHTMXAttributes(field *schema.Field, props *molecules.FormFieldProps) {
	if field.HTMX == nil {
		return
	}

	// Note: FieldHTMX in schema has different structure than expected HTMX props
	// We'll map common attributes here, but this may need adjustment based on actual schema.FieldHTMX structure
	if method, ok := field.Config["htmx.method"]; ok {
		if methodStr, ok := method.(string); ok {
			switch strings.ToUpper(methodStr) {
			case "POST":
				props.HXPost = field.Config["htmx.url"].(string)
			case "GET":
				props.HXGet = field.Config["htmx.url"].(string)
			case "PUT":
				props.HXPut = field.Config["htmx.url"].(string)
			case "PATCH":
				props.HXPatch = field.Config["htmx.url"].(string)
			case "DELETE":
				props.HXDelete = field.Config["htmx.url"].(string)
			}
		}
	}

	if target, ok := field.Config["htmx.target"]; ok {
		if targetStr, ok := target.(string); ok {
			props.HXTarget = targetStr
		}
	}

	if swap, ok := field.Config["htmx.swap"]; ok {
		if swapStr, ok := swap.(string); ok {
			props.HXSwap = swapStr
		}
	}
}

// Map Alpine.js attributes
func (m *FieldMapper) mapAlpineAttributes(field *schema.Field, props *molecules.FormFieldProps) {
	if field.Alpine == nil {
		return
	}

	if field.Alpine.XModel != "" {
		props.AlpineModel = field.Alpine.XModel
	}
	if field.Alpine.XBind != "" {
		props.AlpineBind = field.Alpine.XBind
	}
	if field.Alpine.XShow != "" {
		props.AlpineShow = field.Alpine.XShow
	}
	// Map other Alpine attributes as needed
}

// Map styling and layout
func (m *FieldMapper) mapStyling(field *schema.Field, props *molecules.FormFieldProps) {
	if field.Style == nil {
		return
	}

	if field.Style.Classes != "" {
		props.Class = field.Style.Classes
	}
	if field.Style.LabelClass != "" {
		props.LabelClass = field.Style.LabelClass
	}
	if field.Style.InputClass != "" {
		props.InputClass = field.Style.InputClass
	}
	if field.Style.ErrorClass != "" {
		props.ErrorClass = field.Style.ErrorClass
	}
}

// Helper methods

func (m *FieldMapper) generateFieldID(name string) string {
	if name == "" {
		return ""
	}
	// Convert field name to valid HTML ID
	id := strings.ReplaceAll(name, " ", "-")
	id = strings.ReplaceAll(id, "_", "-")
	return strings.ToLower(id)
}

func (m *FieldMapper) convertValueToString(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func (m *FieldMapper) formatErrors(errors []string) string {
	if len(errors) == 0 {
		return ""
	}
	return strings.Join(errors, "; ")
}

func (m *FieldMapper) getLocalizedText(defaultText string, i18n *schema.FieldI18n, textType string) string {
	// If no i18n or locale is default, return original
	if i18n == nil || m.locale == "" || m.locale == "en" {
		return defaultText
	}

	// Try to get localized text based on type
	switch textType {
	case "label":
		if i18n.Label != nil {
			if localized, ok := i18n.Label[m.locale]; ok && localized != "" {
				return localized
			}
		}
	case "placeholder":
		if i18n.Placeholder != nil {
			if localized, ok := i18n.Placeholder[m.locale]; ok && localized != "" {
				return localized
			}
		}
	case "help":
		if i18n.Help != nil {
			if localized, ok := i18n.Help[m.locale]; ok && localized != "" {
				return localized
			}
		}
	case "description":
		if i18n.Description != nil {
			if localized, ok := i18n.Description[m.locale]; ok && localized != "" {
				return localized
			}
		}
	case "tooltip":
		if i18n.Tooltip != nil {
			if localized, ok := i18n.Tooltip[m.locale]; ok && localized != "" {
				return localized
			}
		}
	}

	// Fallback to default text
	return defaultText
}

func (m *FieldMapper) buildActionClasses(action *schema.Action) string {
	var classes []string

	// Size classes
	switch action.Size {
	case "sm":
		classes = append(classes, "btn-sm")
	case "lg":
		classes = append(classes, "btn-lg")
	case "xl":
		classes = append(classes, "btn-xl")
	default:
		classes = append(classes, "btn-md")
	}

	// State classes
	if action.Loading {
		classes = append(classes, "btn-loading")
	}

	return strings.Join(classes, " ")
}
