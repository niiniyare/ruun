package views

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/tokens"
	"github.com/niiniyare/ruun/views/validation"
)

// FieldMapper converts pure schema field definitions to FormField props
// This maintains separation between schema (business logic) and views (UI)
type FieldMapper struct {
	schema        *schema.Schema
	locale        string
	theme         *schema.Theme // Design tokens for styling
	tokenRegistry TokenRegistry // Token resolution registry
}

// NewFieldMapper creates a new field mapper with schema context and locale
func NewFieldMapper(s *schema.Schema, locale string) *FieldMapper {
	return &FieldMapper{
		schema: s,
		locale: locale,
		theme:  nil, // Set via WithTheme
	}
}

// WithTheme sets the theme for design token integration
func (m *FieldMapper) WithTheme(theme *schema.Theme) *FieldMapper {
	m.theme = theme
	return m
}

// WithTokenRegistry sets the token registry for token resolution
func (m *FieldMapper) WithTokenRegistry(registry TokenRegistry) *FieldMapper {
	m.tokenRegistry = registry
	return m
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

	// Apply design tokens to styling
	m.applyDesignTokens(field, &props)

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

// Design token integration methods

// applyDesignTokens applies theme-based design tokens to form field props
func (m *FieldMapper) applyDesignTokens(field *schema.Field, props *molecules.FormFieldProps) {
	if m.theme == nil || m.theme.Tokens == nil {
		return // No theme available
	}

	tokens := m.theme.Tokens

	// Apply component-level token overrides if available
	if tokens.Components != nil {
		m.applyComponentTokens(field, props, tokens.Components)
	}

	// Apply semantic token mappings
	if tokens.Semantic != nil {
		m.applySemanticTokens(field, props, tokens.Semantic)
	}

	// Apply field-specific styling from schema
	m.applyFieldStyling(field, props)
}

// applyComponentTokens applies component-specific design tokens
func (m *FieldMapper) applyComponentTokens(field *schema.Field, props *molecules.FormFieldProps, components *schema.ComponentTokens) {
	if components == nil {
		return
	}

	// Apply form component tokens
	if components.Form != nil {
		formTokens := components.Form

		// Apply border radius from form tokens
		if formTokens.BorderRadius != "" {
			props.Rounded = m.resolveTokenReference(string(formTokens.BorderRadius))
		}
	}

	// Apply input component tokens
	if components.Input != nil {
		inputTokens := components.Input

		// Apply border radius from input tokens
		if inputTokens.BorderRadius != "" {
			props.Rounded = m.resolveTokenReference(string(inputTokens.BorderRadius))
		}

		// Apply input styling from tokens
		var inputClasses []string

		// Background and focus states
		if inputTokens.Background != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("bg-%s", m.resolveTokenReference(string(inputTokens.Background))))
		}
		if inputTokens.BackgroundFocus != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("focus:bg-%s", m.resolveTokenReference(string(inputTokens.BackgroundFocus))))
		}

		// Border and focus states
		if inputTokens.Border != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("border-%s", m.resolveTokenReference(string(inputTokens.Border))))
		}
		if inputTokens.BorderFocus != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("focus:border-%s", m.resolveTokenReference(string(inputTokens.BorderFocus))))
		}

		// Error border for invalid state
		if inputTokens.BorderError != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("invalid:border-%s", m.resolveTokenReference(string(inputTokens.BorderError))))
		}

		// Text color and placeholder
		if inputTokens.Color != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("text-%s", m.resolveTokenReference(string(inputTokens.Color))))
		}
		if inputTokens.Placeholder != "" {
			inputClasses = append(inputClasses, fmt.Sprintf("placeholder:%s", m.resolveTokenReference(string(inputTokens.Placeholder))))
		}

		// Apply input classes
		if len(inputClasses) > 0 {
			tokenClasses := strings.Join(inputClasses, " ")
			if props.InputClass != "" {
				props.InputClass += " " + tokenClasses
			} else {
				props.InputClass = tokenClasses
			}
		}
	}
}

// applySemanticTokens applies semantic design tokens for consistent theming
func (m *FieldMapper) applySemanticTokens(field *schema.Field, props *molecules.FormFieldProps, semantic *schema.SemanticTokens) {
	if semantic == nil {
		return
	}

	// Apply color semantics
	if semantic.Colors != nil {
		// Apply interactive states
		if semantic.Colors.Interactive != nil && semantic.Colors.Interactive.Primary != nil {
			var classes []string

			// Focus states (primary interactive colors)
			if semantic.Colors.Interactive.Primary.Focus != "" {
				classes = append(classes, fmt.Sprintf("focus:border-%s", m.resolveTokenReference(string(semantic.Colors.Interactive.Primary.Focus))))
			}

			// Hover states for interactive fields
			if m.isInteractiveField(field.Type) && semantic.Colors.Interactive.Primary.Hover != "" {
				classes = append(classes, fmt.Sprintf("hover:border-%s", m.resolveTokenReference(string(semantic.Colors.Interactive.Primary.Hover))))
			}

			// Add semantic classes to input
			if len(classes) > 0 {
				semanticClasses := strings.Join(classes, " ")
				if props.InputClass != "" {
					props.InputClass += " " + semanticClasses
				} else {
					props.InputClass = semanticClasses
				}
			}
		}

		// Apply feedback colors
		if semantic.Colors.Feedback != nil && semantic.Colors.Feedback.Error != nil {
			// Error states
			if semantic.Colors.Feedback.Error.Default != "" {
				errorClass := fmt.Sprintf("invalid:border-%s", m.resolveTokenReference(string(semantic.Colors.Feedback.Error.Default)))
				if props.ErrorClass != "" {
					props.ErrorClass += " " + errorClass
				} else {
					props.ErrorClass = errorClass
				}
			}
		}
	}

	// Apply spacing tokens
	if semantic.Spacing != nil && semantic.Spacing.Component != nil {
		// Apply consistent component spacing (use default spacing)
		if semantic.Spacing.Component.Default != "" {
			gapClass := fmt.Sprintf("gap-%s", m.resolveTokenReference(string(semantic.Spacing.Component.Default)))
			if props.WrapperClass != "" {
				props.WrapperClass += " " + gapClass
			} else {
				props.WrapperClass = gapClass
			}
		}
	}
}

// applyFieldStyling applies field-specific styling from schema.Field.Style
func (m *FieldMapper) applyFieldStyling(field *schema.Field, props *molecules.FormFieldProps) {
	// Field styling is already handled in mapStyling method
	// This method can be used for additional token-based field styling

	// Apply field-level theme overrides if present
	if field.Config != nil {
		// Check for theme override configuration
		if themeOverride, exists := field.Config["theme"]; exists {
			if themeMap, ok := themeOverride.(map[string]any); ok {
				m.applyThemeOverrides(themeMap, props)
			}
		}
	}
}

// Helper methods for token integration

// resolveTokenReference resolves a token reference to its actual value
func (m *FieldMapper) resolveTokenReference(tokenRef string) string {
	if m.theme == nil || m.theme.Tokens == nil {
		return tokenRef
	}

	// Simple token reference resolution
	// In a full implementation, this would traverse the token hierarchy
	// For now, return the reference as a CSS custom property
	if strings.HasPrefix(tokenRef, "{") && strings.HasSuffix(tokenRef, "}") {
		// Convert {semantic.spacing.md} to var(--semantic-spacing-md)
		cleanRef := strings.Trim(tokenRef, "{}")
		cssVar := strings.ReplaceAll(cleanRef, ".", "-")
		return fmt.Sprintf("var(--%s)", cssVar)
	}

	return tokenRef
}

// getFieldTypeClass returns type-specific CSS class based on field type
func (m *FieldMapper) getFieldTypeClass(fieldType schema.FieldType) string {
	// Map field types to basic CSS classes
	// In a full implementation, these would be configurable via tokens
	typeClasses := map[schema.FieldType]string{
		schema.FieldText:         "input-text",
		schema.FieldEmail:        "input-email",
		schema.FieldPassword:     "input-password",
		schema.FieldNumber:       "input-number",
		schema.FieldSelect:       "input-select",
		schema.FieldMultiSelect:  "input-multi-select",
		schema.FieldTextarea:     "input-textarea",
		schema.FieldCheckbox:     "input-checkbox",
		schema.FieldRadio:        "input-radio",
		schema.FieldAutoComplete: "input-autocomplete",
		schema.FieldFile:         "input-file",
		schema.FieldDate:         "input-date",
		schema.FieldTime:         "input-time",
		schema.FieldDateTime:     "input-datetime",
		schema.FieldSlider:       "input-slider",
		schema.FieldColor:        "input-color",
		schema.FieldTags:         "input-tags",
	}

	if class, exists := typeClasses[fieldType]; exists {
		return class
	}

	// Fallback to general input class
	return "input-default"
}

// isInteractiveField checks if a field type supports interactive states
func (m *FieldMapper) isInteractiveField(fieldType schema.FieldType) bool {
	interactiveTypes := map[schema.FieldType]bool{
		schema.FieldText:         true,
		schema.FieldEmail:        true,
		schema.FieldPassword:     true,
		schema.FieldNumber:       true,
		schema.FieldSelect:       true,
		schema.FieldMultiSelect:  true,
		schema.FieldTextarea:     true,
		schema.FieldAutoComplete: true,
		schema.FieldSlider:       true,
		schema.FieldCheckbox:     true,
		schema.FieldRadio:        true,
	}

	return interactiveTypes[fieldType]
}

// applyThemeOverrides applies field-level theme overrides
func (m *FieldMapper) applyThemeOverrides(themeOverride map[string]any, props *molecules.FormFieldProps) {
	// Apply class overrides
	if class, exists := themeOverride["class"]; exists {
		if classStr, ok := class.(string); ok {
			if props.Class != "" {
				props.Class += " " + classStr
			} else {
				props.Class = classStr
			}
		}
	}

	// Apply input class overrides
	if inputClass, exists := themeOverride["inputClass"]; exists {
		if inputClassStr, ok := inputClass.(string); ok {
			if props.InputClass != "" {
				props.InputClass += " " + inputClassStr
			} else {
				props.InputClass = inputClassStr
			}
		}
	}

	// Apply label class overrides
	if labelClass, exists := themeOverride["labelClass"]; exists {
		if labelClassStr, ok := labelClass.(string); ok {
			if props.LabelClass != "" {
				props.LabelClass += " " + labelClassStr
			} else {
				props.LabelClass = labelClassStr
			}
		}
	}

	// Apply rounded overrides
	if rounded, exists := themeOverride["rounded"]; exists {
		if roundedStr, ok := rounded.(string); ok {
			props.Rounded = roundedStr
		}
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// Client Validation Rule Extraction (EXTENSION)
// ═══════════════════════════════════════════════════════════════════════════

// Use ClientValidationRules and ClientValidationRule from molecules package to avoid type conflicts

// ExtractClientValidationRules extracts validation rules that can be validated client-side
// REUSES existing schema validation configuration to generate client rules
func (m *FieldMapper) ExtractClientValidationRules(field *schema.Field) molecules.ClientValidationRules {
	rules := molecules.ClientValidationRules{
		FieldName: field.Name,
		Required:  field.Required,
	}

	// Extract rules that can be validated client-side
	if field.Validation != nil {
		v := field.Validation

		// String length constraints
		if v.MinLength != nil {
			rules.MinLength = v.MinLength
		}
		if v.MaxLength != nil {
			rules.MaxLength = v.MaxLength
		}

		// Numeric constraints
		if v.Min != nil {
			rules.Min = v.Min
		}
		if v.Max != nil {
			rules.Max = v.Max
		}
		if v.Step != nil {
			rules.Step = v.Step
		}

		// Pattern validation
		if v.Pattern != "" {
			rules.Pattern = v.Pattern
		}

		// Format validation (can be handled client-side)
		if v.Format != "" {
			switch v.Format {
			case "email", "url", "phone", "date", "time", "datetime", "color":
				rules.Format = v.Format
			}
		}

		// Extract client-validatable custom rules
		if v.Custom != "" {
			clientRule := m.extractClientCustomRule(v.Custom, v.Messages)
			if clientRule != nil {
				rules.CustomRules = append(rules.CustomRules, *clientRule)
			}
		}
	}

	// Extract field-type specific client rules
	m.extractFieldTypeClientRules(field, &rules)

	return rules
}

// extractClientCustomRule attempts to extract client-side executable custom validation
func (m *FieldMapper) extractClientCustomRule(customExpr string, messages any) *molecules.ClientValidationRule {
	// For complex expressions, mark as async (server-side validation)
	return &molecules.ClientValidationRule{
		Type:    "async_custom",
		Value:   customExpr,
		Message: m.getCustomRuleMessage(messages, "custom"),
		Async:   true,
	}
}

// extractFieldTypeClientRules extracts field-type-specific client validation rules
func (m *FieldMapper) extractFieldTypeClientRules(field *schema.Field, rules *molecules.ClientValidationRules) {
	switch field.Type {
	case schema.FieldEmail:
		rules.CustomRules = append(rules.CustomRules, molecules.ClientValidationRule{
			Type:    "email_format",
			Message: "Please enter a valid email address",
			Async:   false,
		})

	case schema.FieldURL:
		rules.CustomRules = append(rules.CustomRules, molecules.ClientValidationRule{
			Type:    "url_format",
			Message: "Please enter a valid URL",
			Async:   false,
		})

	case schema.FieldPhone:
		rules.CustomRules = append(rules.CustomRules, molecules.ClientValidationRule{
			Type:    "phone_format",
			Message: "Please enter a valid phone number",
			Async:   false,
		})
	}

	// Add uniqueness validation if specified
	if field.Validation != nil && field.Validation.Unique {
		rules.CustomRules = append(rules.CustomRules, molecules.ClientValidationRule{
			Type:    "uniqueness",
			Message: "This value must be unique",
			Async:   true, // Uniqueness validation requires database check
		})
	}
}

// getCustomRuleMessage extracts the appropriate error message for custom rules
func (m *FieldMapper) getCustomRuleMessage(messages any, ruleType string) string {
	if messages == nil {
		return "Validation failed"
	}

	// Fallback to generic validation message
	return "Please enter a valid value"
}

// GetClientValidationConfig returns client validation configuration for a field
func (m *FieldMapper) GetClientValidationConfig(field *schema.Field) map[string]any {
	rules := m.ExtractClientValidationRules(field)

	config := map[string]any{
		"fieldName":  rules.FieldName,
		"required":   rules.Required,
		"rules":      rules,
		"realTime":   true, // Enable real-time client validation
		"debounceMs": 300,  // Default debounce timing
	}

	return config
}

// HasClientValidationRules checks if a field has any client-side validation rules
func (m *FieldMapper) HasClientValidationRules(field *schema.Field) bool {
	rules := m.ExtractClientValidationRules(field)

	return rules.Required ||
		rules.MinLength != nil ||
		rules.MaxLength != nil ||
		rules.Min != nil ||
		rules.Max != nil ||
		rules.Pattern != "" ||
		rules.Format != "" ||
		len(rules.CustomRules) > 0
}

// ═══════════════════════════════════════════════════════════════════════════
// TOKEN RESOLUTION METHODS FOR THEME-AWARE VALIDATION
// ═══════════════════════════════════════════════════════════════════════════

// TokenRegistry interface for token resolution
type TokenRegistry interface {
	ResolveToken(ctx context.Context, tokenRef schema.TokenReference) (string, error)
	ResolveTokens(ctx context.Context, tokenRefs map[string]schema.TokenReference) (map[string]string, error)
	GetCachedTokens(cacheKey string) (map[string]string, bool)
	SetCachedTokens(cacheKey string, tokens map[string]string, expiry time.Duration)
}

// DefaultTokenRegistry provides a default implementation
type DefaultTokenRegistry struct {
	baseTokens      *schema.DesignTokens
	tokenCache      map[string]map[string]string
	cacheExpiry     map[string]time.Time
	cacheMutex      sync.RWMutex
	cacheExpiration time.Duration
}

// NewDefaultTokenRegistry creates a new default token registry
func NewDefaultTokenRegistry() *DefaultTokenRegistry {
	return &DefaultTokenRegistry{
		baseTokens:      tokens.ValidationTokens,
		tokenCache:      make(map[string]map[string]string),
		cacheExpiry:     make(map[string]time.Time),
		cacheExpiration: 5 * time.Minute,
	}
}

// ResolveToken resolves a single token reference to its value
func (r *DefaultTokenRegistry) ResolveToken(ctx context.Context, tokenRef schema.TokenReference) (string, error) {
	// Parse token reference (e.g., "components.field.background.default")
	tokenPath := string(tokenRef)

	// Try to resolve from validation tokens
	if value := r.resolveFromDesignTokens(tokenPath); value != "" {
		return value, nil
	}

	// Return empty if not found (caller should handle fallback)
	return "", fmt.Errorf("token not found: %s", tokenPath)
}

// ResolveTokens resolves multiple token references
func (r *DefaultTokenRegistry) ResolveTokens(ctx context.Context, tokenRefs map[string]schema.TokenReference) (map[string]string, error) {
	resolved := make(map[string]string)

	for key, tokenRef := range tokenRefs {
		value, err := r.ResolveToken(ctx, tokenRef)
		if err == nil {
			resolved[key] = value
		}
		// Skip tokens that can't be resolved (fallbacks will be used)
	}

	return resolved, nil
}

// GetCachedTokens retrieves tokens from cache
func (r *DefaultTokenRegistry) GetCachedTokens(cacheKey string) (map[string]string, bool) {
	r.cacheMutex.RLock()
	defer r.cacheMutex.RUnlock()

	// Check if cache entry exists and is not expired
	if expiry, exists := r.cacheExpiry[cacheKey]; exists {
		if time.Now().Before(expiry) {
			if tokens, found := r.tokenCache[cacheKey]; found {
				return tokens, true
			}
		}
	}

	return nil, false
}

// SetCachedTokens stores tokens in cache
func (r *DefaultTokenRegistry) SetCachedTokens(cacheKey string, tokens map[string]string, expiry time.Duration) {
	r.cacheMutex.Lock()
	defer r.cacheMutex.Unlock()

	r.tokenCache[cacheKey] = tokens
	r.cacheExpiry[cacheKey] = time.Now().Add(expiry)
}

// resolveFromDesignTokens resolves a token from the design token system
func (r *DefaultTokenRegistry) resolveFromDesignTokens(tokenPath string) string {
	if r.baseTokens == nil {
		return ""
	}

	parts := strings.Split(tokenPath, ".")
	if len(parts) < 2 {
		return ""
	}

	switch parts[0] {
	case "primitives":
		return r.resolvePrimitiveToken(parts[1:])
	case "semantic":
		return r.resolveSemanticToken(parts[1:])
	case "components":
		return r.resolveComponentToken(parts[1:])
	default:
		return ""
	}
}

// resolvePrimitiveToken resolves primitive tokens
func (r *DefaultTokenRegistry) resolvePrimitiveToken(parts []string) string {
	if r.baseTokens.Primitives == nil || len(parts) < 2 {
		return ""
	}

	switch parts[0] {
	case "colors":
		return r.resolvePrimitiveColor(parts[1:])
	case "typography":
		return r.resolvePrimitiveTypography(parts[1:])
	case "spacing":
		return r.resolvePrimitiveSpacing(parts[1:])
	case "border":
		return r.resolvePrimitiveBorder(parts[1:])
	case "animation":
		return r.resolvePrimitiveAnimation(parts[1:])
	case "shadow":
		return r.resolvePrimitiveShadow(parts[1:])
	}

	return ""
}

// resolvePrimitiveColor resolves primitive color tokens
func (r *DefaultTokenRegistry) resolvePrimitiveColor(parts []string) string {
	if r.baseTokens.Primitives.Colors == nil || len(parts) < 2 {
		return ""
	}

	colorGroup := parts[0]
	colorKey := parts[1]

	switch colorGroup {
	case "gray":
		if r.baseTokens.Primitives.Colors.Gray != nil {
			if value := getColorFromGrayScale(r.baseTokens.Primitives.Colors.Gray, colorKey); value != "" {
				return value
			}
		}
	case "blue":
		if r.baseTokens.Primitives.Colors.Blue != nil {
			if value := getColorFromScale(r.baseTokens.Primitives.Colors.Blue, colorKey); value != "" {
				return value
			}
		}
	case "green":
		if r.baseTokens.Primitives.Colors.Green != nil {
			if value := getColorFromScale(r.baseTokens.Primitives.Colors.Green, colorKey); value != "" {
				return value
			}
		}
	case "red":
		if r.baseTokens.Primitives.Colors.Red != nil {
			if value := getColorFromScale(r.baseTokens.Primitives.Colors.Red, colorKey); value != "" {
				return value
			}
		}
	case "yellow":
		if r.baseTokens.Primitives.Colors.Yellow != nil {
			if value := getColorFromScale(r.baseTokens.Primitives.Colors.Yellow, colorKey); value != "" {
				return value
			}
		}
	}

	return ""
}

// resolvePrimitiveTypography resolves primitive typography tokens
func (r *DefaultTokenRegistry) resolvePrimitiveTypography(parts []string) string {
	if r.baseTokens.Primitives.Typography == nil || len(parts) < 2 {
		return ""
	}

	typoGroup := parts[0]
	typoKey := parts[1]

	switch typoGroup {
	case "fontSizes":
		if r.baseTokens.Primitives.Typography.FontSizes != nil {
			if value := getFontSizeFromScale(r.baseTokens.Primitives.Typography.FontSizes, typoKey); value != "" {
				return value
			}
		}
	case "fontWeights":
		if r.baseTokens.Primitives.Typography.FontWeights != nil {
			if value := getFontWeightFromScale(r.baseTokens.Primitives.Typography.FontWeights, typoKey); value != "" {
				return value
			}
		}
	case "lineHeights":
		if r.baseTokens.Primitives.Typography.LineHeights != nil {
			if value := getLineHeightFromScale(r.baseTokens.Primitives.Typography.LineHeights, typoKey); value != "" {
				return value
			}
		}
	}

	return ""
}

// resolvePrimitiveSpacing resolves primitive spacing tokens
func (r *DefaultTokenRegistry) resolvePrimitiveSpacing(parts []string) string {
	if r.baseTokens.Primitives.Spacing == nil || len(parts) < 2 {
		return ""
	}

	if parts[0] == "spacing" && r.baseTokens.Primitives.Spacing != nil {
		if value := getSpacingFromScale(r.baseTokens.Primitives.Spacing, parts[1]); value != "" {
			return value
		}
	}

	return ""
}

// resolvePrimitiveBorder resolves primitive border tokens
func (r *DefaultTokenRegistry) resolvePrimitiveBorder(parts []string) string {
	if r.baseTokens.Primitives.Borders == nil || len(parts) < 2 {
		return ""
	}

	borderGroup := parts[0]
	borderKey := parts[1]

	switch borderGroup {
	case "width":
		if r.baseTokens.Primitives.Borders.Width != nil {
			if value := getBorderWidthFromScale(r.baseTokens.Primitives.Borders.Width, borderKey); value != "" {
				return value
			}
		}
	case "radius":
		if r.baseTokens.Primitives.Borders.Radius != nil {
			if value := getBorderRadiusFromScale(r.baseTokens.Primitives.Borders.Radius, borderKey); value != "" {
				return value
			}
		}
	}

	return ""
}

// resolvePrimitiveAnimation resolves primitive animation tokens
func (r *DefaultTokenRegistry) resolvePrimitiveAnimation(parts []string) string {
	if r.baseTokens.Primitives.Animations == nil || len(parts) < 2 {
		return ""
	}

	animGroup := parts[0]
	animKey := parts[1]

	switch animGroup {
	case "duration":
		if r.baseTokens.Primitives.Animations.Duration != nil {
			return getDurationFromScale(r.baseTokens.Primitives.Animations.Duration, animKey)
		}
	case "easing":
		if r.baseTokens.Primitives.Animations.Easing != nil {
			return getEasingFromScale(r.baseTokens.Primitives.Animations.Easing, animKey)
		}
	}

	return ""
}

// resolvePrimitiveShadow resolves primitive shadow tokens
func (r *DefaultTokenRegistry) resolvePrimitiveShadow(parts []string) string {
	if r.baseTokens.Primitives.Shadows == nil || len(parts) < 1 {
		return ""
	}

	// Shadow tokens don't have nested groups, they're direct scale values
	shadowKey := parts[0]
	return getShadowFromScale(r.baseTokens.Primitives.Shadows, shadowKey)
}

// resolveSemanticToken resolves semantic tokens (with potential primitive references)
func (r *DefaultTokenRegistry) resolveSemanticToken(parts []string) string {
	if r.baseTokens.Semantic == nil || len(parts) < 2 {
		return ""
	}

	switch parts[0] {
	case "colors":
		return r.resolveSemanticColor(parts[1:])
	case "typography":
		return r.resolveSemanticTypography(parts[1:])
	case "spacing":
		return r.resolveSemanticSpacing(parts[1:])
	case "animation":
		return r.resolveSemanticInteractive(parts[1:])
	}

	return ""
}

// resolveSemanticColor resolves semantic color tokens
func (r *DefaultTokenRegistry) resolveSemanticColor(parts []string) string {
	if r.baseTokens.Semantic.Colors == nil || len(parts) < 2 {
		return ""
	}

	colorGroup := parts[0]
	colorKey := parts[1]

	var tokenRef string

	switch colorGroup {
	case "background":
		if r.baseTokens.Semantic.Colors.Background != nil {
			tokenRef = string(getBackgroundColorFromStruct(r.baseTokens.Semantic.Colors.Background, colorKey))
		}
	case "text":
		if r.baseTokens.Semantic.Colors.Text != nil {
			tokenRef = string(getTextColorFromStruct(r.baseTokens.Semantic.Colors.Text, colorKey))
		}
	case "border":
		if r.baseTokens.Semantic.Colors.Border != nil {
			tokenRef = string(getBorderColorFromStruct(r.baseTokens.Semantic.Colors.Border, colorKey))
		}
	case "interactive":
		if r.baseTokens.Semantic.Colors.Interactive != nil {
			tokenRef = string(getInteractiveColorFromStruct(r.baseTokens.Semantic.Colors.Interactive, colorKey))
		}
	case "feedback":
		if r.baseTokens.Semantic.Colors.Feedback != nil {
			tokenRef = string(getFeedbackColorFromStruct(r.baseTokens.Semantic.Colors.Feedback, colorKey))
		}
	}

	// If token reference found, resolve it recursively
	if tokenRef != "" {
		return r.resolveFromDesignTokens(tokenRef)
	}

	return ""
}

// resolveSemanticTypography resolves semantic typography tokens
func (r *DefaultTokenRegistry) resolveSemanticTypography(parts []string) string {
	if r.baseTokens.Semantic.Typography == nil || len(parts) < 2 {
		return ""
	}

	typoGroup := parts[0]
	typoKey := parts[1]

	var tokenRef string

	switch typoGroup {
	case "body":
		if r.baseTokens.Semantic.Typography.Body != nil {
			tokenRef = string(getBodyTypographyFromStruct(r.baseTokens.Semantic.Typography.Body, typoKey))
		}
	case "labels":
		if r.baseTokens.Semantic.Typography.Labels != nil {
			tokenRef = string(getLabelTypographyFromStruct(r.baseTokens.Semantic.Typography.Labels, typoKey))
		}
	case "captions":
		if r.baseTokens.Semantic.Typography.Captions != nil {
			tokenRef = string(getCaptionTypographyFromStruct(r.baseTokens.Semantic.Typography.Captions, typoKey))
		}
	}

	if tokenRef != "" {
		return r.resolveFromDesignTokens(tokenRef)
	}

	return ""
}

// resolveSemanticSpacing resolves semantic spacing tokens
func (r *DefaultTokenRegistry) resolveSemanticSpacing(parts []string) string {
	if r.baseTokens.Semantic.Spacing == nil || len(parts) < 2 {
		return ""
	}

	if parts[0] == "component" && r.baseTokens.Semantic.Spacing.Component != nil {
		tokenRef := string(getComponentSpacingFromStruct(r.baseTokens.Semantic.Spacing.Component, parts[1]))
		if tokenRef != "" {
			return r.resolveFromDesignTokens(tokenRef)
		}
	}

	return ""
}

// resolveSemanticInteractive resolves semantic interactive tokens
func (r *DefaultTokenRegistry) resolveSemanticInteractive(parts []string) string {
	if r.baseTokens.Semantic.Interactive == nil || len(parts) < 2 {
		return ""
	}

	// Handle interactive tokens like borderRadius, shadow
	switch parts[0] {
	case "borderRadius":
		if r.baseTokens.Semantic.Interactive.BorderRadius != nil {
			tokenRef := string(getInteractiveBorderRadiusFromStruct(r.baseTokens.Semantic.Interactive.BorderRadius, parts[1]))
			if tokenRef != "" {
				return r.resolveFromDesignTokens(tokenRef)
			}
		}
	case "shadow":
		if r.baseTokens.Semantic.Interactive.Shadow != nil {
			tokenRef := string(getInteractiveShadowFromStruct(r.baseTokens.Semantic.Interactive.Shadow, parts[1]))
			if tokenRef != "" {
				return r.resolveFromDesignTokens(tokenRef)
			}
		}
	}

	return ""
}

// resolveComponentToken resolves component tokens
func (r *DefaultTokenRegistry) resolveComponentToken(parts []string) string {
	if r.baseTokens.Components == nil || len(parts) < 2 {
		return ""
	}

	switch parts[0] {
	case "input":
		return r.resolveInputComponentToken(parts[1:])
	case "button":
		return r.resolveButtonComponentToken(parts[1:])
	case "form":
		return r.resolveFormComponentToken(parts[1:])
	}

	return ""
}

// resolveInputComponentToken resolves input component tokens
func (r *DefaultTokenRegistry) resolveInputComponentToken(parts []string) string {
	if r.baseTokens.Components.Input == nil || len(parts) == 0 {
		return ""
	}

	// Return input token references - these are TokenReferences that need resolution
	switch parts[0] {
	case "background":
		return string(r.baseTokens.Components.Input.Background)
	case "backgroundFocus":
		return string(r.baseTokens.Components.Input.BackgroundFocus)
	case "border":
		return string(r.baseTokens.Components.Input.Border)
	case "borderFocus":
		return string(r.baseTokens.Components.Input.BorderFocus)
	case "borderError":
		return string(r.baseTokens.Components.Input.BorderError)
	case "borderRadius":
		return string(r.baseTokens.Components.Input.BorderRadius)
	case "padding":
		return string(r.baseTokens.Components.Input.Padding)
	case "color":
		return string(r.baseTokens.Components.Input.Color)
	case "placeholder":
		return string(r.baseTokens.Components.Input.Placeholder)
	}
	return ""
}

// resolveButtonComponentToken resolves button component tokens
func (r *DefaultTokenRegistry) resolveButtonComponentToken(parts []string) string {
	if r.baseTokens.Components.Button == nil || len(parts) < 2 {
		return ""
	}

	// Handle button variants like primary.background, secondary.hover, etc.
	variant := parts[0]
	prop := parts[1]

	switch variant {
	case "primary":
		if r.baseTokens.Components.Button.Primary != nil {
			return string(getButtonTokenFromVariant(r.baseTokens.Components.Button.Primary, prop))
		}
	case "secondary":
		if r.baseTokens.Components.Button.Secondary != nil {
			return string(getButtonTokenFromVariant(r.baseTokens.Components.Button.Secondary, prop))
		}
	}
	return ""
}

// resolveFormComponentToken resolves form component tokens
func (r *DefaultTokenRegistry) resolveFormComponentToken(parts []string) string {
	if r.baseTokens.Components.Form == nil || len(parts) == 0 {
		return ""
	}

	// Form tokens would be handled here when defined in schema
	return ""
}

// ResolveFieldTokens resolves tokens for a field with validation state and dark mode support
// This is the main method called by the renderer and FormField component
func (m *FieldMapper) ResolveFieldTokens(
	ctx context.Context,
	props molecules.FormFieldProps,
	darkMode bool,
) map[string]string {
	// Build cache key for performance
	cacheKey := m.buildTokenCacheKey(props.Name, props.ValidationState.String(), props.ThemeID, darkMode)

	// Check cache first
	if m.tokenRegistry != nil {
		if cachedTokens, found := m.tokenRegistry.GetCachedTokens(cacheKey); found {
			return cachedTokens
		}
	}

	// Build base token references
	baseTokenRefs := m.buildBaseTokenReferences(props)

	// Apply validation-specific token refs
	validationTokenRefs := m.buildValidationTokenReferences(props.ValidationState)
	for key, ref := range validationTokenRefs {
		baseTokenRefs[key] = ref
	}

	// Merge in explicit props.Tokens overrides (highest precedence)
	for key, value := range props.Tokens {
		baseTokenRefs[key] = schema.TokenReference(value)
	}

	// Resolve tokens with token registry
	var resolved map[string]string
	if m.tokenRegistry != nil {
		resolved, _ = m.tokenRegistry.ResolveTokens(ctx, baseTokenRefs)
	} else {
		// Fallback resolution without registry
		resolved = m.fallbackTokenResolution(baseTokenRefs)
	}

	// Apply dark mode overrides if requested
	if darkMode {
		darkOverrides := m.getDarkModeTokenOverrides(props.ValidationState)
		for key, override := range darkOverrides {
			if m.tokenRegistry != nil {
				if resolvedValue, err := m.tokenRegistry.ResolveToken(ctx, schema.TokenReference(override)); err == nil {
					resolved[key] = resolvedValue
				}
			} else if fallback := m.fallbackTokenValue(override); fallback != "" {
				resolved[key] = fallback
			}
		}
	}

	// Apply tenant-specific overrides from props.TokenOverrides
	for key, value := range props.TokenOverrides {
		resolved[key] = value
	}

	// Cache resolved tokens
	if m.tokenRegistry != nil {
		m.tokenRegistry.SetCachedTokens(cacheKey, resolved, 5*time.Minute)
	}

	return resolved
}

// buildTokenCacheKey builds a cache key for token resolution
func (m *FieldMapper) buildTokenCacheKey(fieldName, validationState, themeID string, darkMode bool) string {
	return fmt.Sprintf("tokens:%s:%s:%s:%t", fieldName, validationState, themeID, darkMode)
}

// buildBaseTokenReferences builds base token references for a field
func (m *FieldMapper) buildBaseTokenReferences(props molecules.FormFieldProps) map[string]schema.TokenReference {
	baseTokens := map[string]schema.TokenReference{
		molecules.TokenFieldBackground:  "components.field.background.default",
		molecules.TokenFieldBorder:      "components.field.border.color.default",
		molecules.TokenFieldText:        "components.field.text.default",
		molecules.TokenFieldLabel:       "components.field.label.color.default",
		molecules.TokenFieldPlaceholder: "components.field.text.placeholder",
		molecules.TokenFieldHelp:        "components.field.help.color",
		molecules.TokenFieldRadius:      "components.field.border.radius",
		molecules.TokenFieldSpacing:     "components.field.spacing.padding",
		molecules.TokenFieldShadow:      "components.field.shadow.default",
	}

	// Add field state-specific tokens
	if props.Disabled {
		baseTokens[molecules.TokenFieldBackground] = "components.field.background.disabled"
		baseTokens[molecules.TokenFieldText] = "components.field.text.disabled"
		baseTokens[molecules.TokenFieldLabel] = "components.field.label.color.disabled"
	}

	if props.Readonly {
		baseTokens[molecules.TokenFieldBackground] = "components.field.background.readonly"
		baseTokens[molecules.TokenFieldText] = "components.field.text.readonly"
	}

	return baseTokens
}

// buildValidationTokenReferences builds validation-specific token references
func (m *FieldMapper) buildValidationTokenReferences(state validation.ValidationState) map[string]schema.TokenReference {
	validationTokens := make(map[string]schema.TokenReference)

	switch state {
	case validation.ValidationStateValidating:
		validationTokens[molecules.TokenFieldBorder] = "components.field.border.color.validating"
		validationTokens[molecules.TokenValidationLoading] = "components.field.validation.loading.color"
		validationTokens[molecules.TokenValidationLoading+".icon"] = "components.field.validation.loading.icon"
		validationTokens[molecules.TokenValidationLoading+".message"] = "components.field.validation.loading.message"

	case validation.ValidationStateValid:
		validationTokens[molecules.TokenFieldBorder] = "components.field.border.color.valid"
		validationTokens[molecules.TokenValidationSuccess] = "components.field.validation.success.color"
		validationTokens[molecules.TokenValidationSuccess+".icon"] = "components.field.validation.success.icon"
		validationTokens[molecules.TokenValidationSuccess+".message"] = "components.field.validation.success.message"

	case validation.ValidationStateInvalid:
		validationTokens[molecules.TokenFieldBorder] = "components.field.border.color.invalid"
		validationTokens[molecules.TokenValidationError] = "components.field.validation.error.color"
		validationTokens[molecules.TokenValidationError+".icon"] = "components.field.validation.error.icon"
		validationTokens[molecules.TokenValidationError+".message"] = "components.field.validation.error.message"

	case validation.ValidationStateWarning:
		validationTokens[molecules.TokenFieldBorder] = "components.field.border.color.warning"
		validationTokens[molecules.TokenValidationWarning] = "components.field.validation.warning.color"
		validationTokens[molecules.TokenValidationWarning+".icon"] = "components.field.validation.warning.icon"
		validationTokens[molecules.TokenValidationWarning+".message"] = "components.field.validation.warning.message"
	}

	return validationTokens
}

// getDarkModeTokenOverrides returns dark mode token overrides
func (m *FieldMapper) getDarkModeTokenOverrides(state validation.ValidationState) map[string]string {
	// Return empty map for now - dark mode overrides would be implemented
	// based on the specific design system requirements
	return make(map[string]string)
}

// validationStateToString converts validation state to string
func (m *FieldMapper) validationStateToString(state validation.ValidationState) string {
	switch state {
	case validation.ValidationStateValidating:
		return "validating"
	case validation.ValidationStateValid:
		return "valid"
	case validation.ValidationStateInvalid:
		return "invalid"
	case validation.ValidationStateWarning:
		return "warning"
	default:
		return "idle"
	}
}

// fallbackTokenResolution provides fallback when no token registry is available
func (m *FieldMapper) fallbackTokenResolution(tokenRefs map[string]schema.TokenReference) map[string]string {
	resolved := make(map[string]string)

	// Provide basic fallback values
	fallbackValues := map[string]string{
		molecules.TokenFieldBackground:                "#ffffff",
		molecules.TokenFieldBorder:                    "#e5e7eb",
		molecules.TokenFieldText:                      "#111827",
		molecules.TokenFieldLabel:                     "#111827",
		molecules.TokenFieldPlaceholder:               "#9ca3af",
		molecules.TokenFieldHelp:                      "#6b7280",
		molecules.TokenFieldRadius:                    "0.375rem",
		molecules.TokenFieldSpacing:                   "0.75rem",
		molecules.TokenFieldShadow:                    "none",
		molecules.TokenValidationLoading:              "#f59e0b",
		molecules.TokenValidationSuccess:              "#16a34a",
		molecules.TokenValidationError:                "#dc2626",
		molecules.TokenValidationWarning:              "#d97706",
		molecules.TokenValidationLoading + ".icon":    "⟳",
		molecules.TokenValidationSuccess + ".icon":    "✓",
		molecules.TokenValidationError + ".icon":      "✗",
		molecules.TokenValidationWarning + ".icon":    "⚠",
		molecules.TokenValidationLoading + ".message": "Validating...",
		molecules.TokenValidationSuccess + ".message": "Valid",
		molecules.TokenValidationError + ".message":   "Invalid",
		molecules.TokenValidationWarning + ".message": "Warning",
	}

	for key := range tokenRefs {
		if fallback, exists := fallbackValues[key]; exists {
			resolved[key] = fallback
		}
	}

	return resolved
}

// fallbackTokenValue provides a fallback value for a single token
func (m *FieldMapper) fallbackTokenValue(tokenRef string) string {
	fallbackMap := map[string]string{
		"semantic.colors.background.primary": "#111827",
		"semantic.colors.text.primary":       "#f9fafb",
		"semantic.colors.border.primary":     "#374151",
		"semantic.colors.validation.loading": "#fbbf24",
		"semantic.colors.validation.success": "#22c55e",
		"semantic.colors.validation.error":   "#ef4444",
		"semantic.colors.validation.warning": "#f59e0b",
	}

	if fallback, exists := fallbackMap[tokenRef]; exists {
		return fallback
	}

	return ""
}

// AddTokenRegistryToFieldMapper is a helper to initialize the token registry
var defaultTokenRegistry *DefaultTokenRegistry

// GetDefaultTokenRegistry returns the default token registry (singleton)
func GetDefaultTokenRegistry() *DefaultTokenRegistry {
	if defaultTokenRegistry == nil {
		defaultTokenRegistry = NewDefaultTokenRegistry()
	}
	return defaultTokenRegistry
}

// Helper functions for accessing token scale structs

// getColorFromGrayScale extracts a color value from a GrayScale struct by key
func getColorFromGrayScale(scale *schema.GrayScale, key string) string {
	switch key {
	case "50":
		return scale.Scale50
	case "100":
		return scale.Scale100
	case "200":
		return scale.Scale200
	case "300":
		return scale.Scale300
	case "400":
		return scale.Scale400
	case "500":
		return scale.Scale500
	case "600":
		return scale.Scale600
	case "700":
		return scale.Scale700
	case "800":
		return scale.Scale800
	case "900":
		return scale.Scale900
	default:
		return ""
	}
}

// getColorFromScale extracts a color value from a ColorScale struct by key
func getColorFromScale(scale *schema.ColorScale, key string) string {
	switch key {
	case "50":
		return scale.Scale50
	case "100":
		return scale.Scale100
	case "200":
		return scale.Scale200
	case "300":
		return scale.Scale300
	case "400":
		return scale.Scale400
	case "500":
		return scale.Scale500
	case "600":
		return scale.Scale600
	case "700":
		return scale.Scale700
	case "800":
		return scale.Scale800
	case "900":
		return scale.Scale900
	default:
		return ""
	}
}

// getFontSizeFromScale extracts a font size value from a FontSizeScale struct by key
func getFontSizeFromScale(scale *schema.FontSizeScale, key string) string {
	switch key {
	case "xs":
		return scale.XS
	case "sm":
		return scale.SM
	case "base":
		return scale.Base
	case "lg":
		return scale.LG
	case "xl":
		return scale.XL
	case "2xl":
		return scale.XXL
	case "3xl":
		return scale.XXXL
	case "4xl":
		return scale.Huge
	default:
		return ""
	}
}

// getFontWeightFromScale extracts a font weight value from a FontWeightScale struct by key
func getFontWeightFromScale(scale *schema.FontWeightScale, key string) string {
	switch key {
	case "thin":
		return scale.Thin
	case "light":
		return scale.Light
	case "normal":
		return scale.Normal
	case "medium":
		return scale.Medium
	case "semibold":
		return scale.Semibold
	case "bold":
		return scale.Bold
	case "extrabold":
		return scale.Extrabold
	default:
		return ""
	}
}

// getLineHeightFromScale extracts a line height value from a LineHeightScale struct by key
func getLineHeightFromScale(scale *schema.LineHeightScale, key string) string {
	switch key {
	case "tight":
		return scale.Tight
	case "normal":
		return scale.Normal
	case "relaxed":
		return scale.Relaxed
	case "loose":
		return scale.Loose
	default:
		return ""
	}
}

// getSpacingFromScale extracts a spacing value from a SpacingScale struct by key
func getSpacingFromScale(scale *schema.SpacingScale, key string) string {
	switch key {
	case "none", "0":
		return scale.None
	case "xs":
		return scale.XS
	case "sm":
		return scale.SM
	case "md":
		return scale.MD
	case "lg":
		return scale.LG
	case "xl":
		return scale.XL
	case "2xl":
		return scale.XXL
	case "3xl":
		return scale.XXXL
	case "4xl":
		return scale.Huge
	default:
		return ""
	}
}

// getBorderWidthFromScale extracts a border width value from a BorderWidthScale struct by key
func getBorderWidthFromScale(scale *schema.BorderWidthScale, key string) string {
	switch key {
	case "none":
		return scale.None
	case "thin":
		return scale.Thin
	case "medium":
		return scale.Medium
	case "thick":
		return scale.Thick
	default:
		return ""
	}
}

// getBorderRadiusFromScale extracts a border radius value from a BorderRadiusScale struct by key
func getBorderRadiusFromScale(scale *schema.BorderRadiusScale, key string) string {
	switch key {
	case "none":
		return scale.None
	case "sm":
		return scale.SM
	case "md":
		return scale.MD
	case "lg":
		return scale.LG
	case "xl":
		return scale.XL
	case "full":
		return scale.Full
	default:
		return ""
	}
}

// getDurationFromScale safely gets duration from animation duration scale
func getDurationFromScale(scale *schema.AnimationDurationScale, key string) string {
	if scale == nil {
		return ""
	}
	switch key {
	case "fast":
		return scale.Fast
	case "normal":
		return scale.Normal
	case "slow":
		return scale.Slow
	default:
		return ""
	}
}

// getEasingFromScale safely gets easing from animation easing scale
func getEasingFromScale(scale *schema.AnimationEasingScale, key string) string {
	if scale == nil {
		return ""
	}
	switch key {
	case "linear":
		return scale.Linear
	case "easeIn":
		return scale.EaseIn
	case "easeOut":
		return scale.EaseOut
	case "easeInOut":
		return scale.EaseInOut
	default:
		return ""
	}
}

// getShadowFromScale safely gets shadow from shadow scale
func getShadowFromScale(scale *schema.ShadowScale, key string) string {
	if scale == nil {
		return ""
	}
	switch key {
	case "none":
		return scale.None
	case "sm":
		return scale.SM
	case "md":
		return scale.MD
	case "lg":
		return scale.LG
	case "xl":
		return scale.XL
	case "2xl":
		return scale.XXL
	case "inner":
		return scale.Inner
	default:
		return ""
	}
}

// getBackgroundColorFromStruct safely gets background color from struct
func getBackgroundColorFromStruct(bg *schema.BackgroundColors, key string) schema.TokenReference {
	if bg == nil {
		return ""
	}
	switch key {
	case "default":
		return bg.Default
	case "subtle":
		return bg.Subtle
	case "emphasis":
		return bg.Emphasis
	case "overlay":
		return bg.Overlay
	default:
		return ""
	}
}

// getTextColorFromStruct safely gets text color from struct
func getTextColorFromStruct(text *schema.TextColors, key string) schema.TokenReference {
	if text == nil {
		return ""
	}
	switch key {
	case "default":
		return text.Default
	case "subtle":
		return text.Subtle
	case "disabled":
		return text.Disabled
	case "inverted":
		return text.Inverted
	case "link":
		return text.Link
	case "linkHover":
		return text.LinkHover
	default:
		return ""
	}
}

// getBorderColorFromStruct safely gets border color from struct
func getBorderColorFromStruct(border *schema.BorderColors, key string) schema.TokenReference {
	if border == nil {
		return ""
	}
	switch key {
	case "default":
		return border.Default
	case "focus":
		return border.Focus
	case "strong":
		return border.Strong
	case "subtle":
		return border.Subtle
	default:
		return ""
	}
}

// getInteractiveColorFromStruct safely gets interactive color from struct
func getInteractiveColorFromStruct(interactive *schema.InteractiveColors, key string) schema.TokenReference {
	if interactive == nil {
		return ""
	}
	// Interactive colors have nested structures like primary.default, secondary.hover, etc.
	// For now, return first match or empty
	if interactive.Primary != nil {
		switch key {
		case "primary.default":
			return interactive.Primary.Default
		case "primary.hover":
			return interactive.Primary.Hover
		case "primary.active":
			return interactive.Primary.Active
		case "primary.focus":
			return interactive.Primary.Focus
		}
	}
	if interactive.Secondary != nil {
		switch key {
		case "secondary.default":
			return interactive.Secondary.Default
		case "secondary.hover":
			return interactive.Secondary.Hover
		case "secondary.active":
			return interactive.Secondary.Active
		case "secondary.focus":
			return interactive.Secondary.Focus
		}
	}
	return ""
}

// getFeedbackColorFromStruct safely gets feedback color from struct
func getFeedbackColorFromStruct(feedback *schema.FeedbackColors, key string) schema.TokenReference {
	if feedback == nil {
		return ""
	}
	// Handle feedback colors like success.default, error.strong, etc.
	if feedback.Success != nil {
		switch key {
		case "success.default":
			return feedback.Success.Default
		case "success.subtle":
			return feedback.Success.Subtle
		case "success.strong":
			return feedback.Success.Strong
		}
	}
	if feedback.Error != nil {
		switch key {
		case "error.default":
			return feedback.Error.Default
		case "error.subtle":
			return feedback.Error.Subtle
		case "error.strong":
			return feedback.Error.Strong
		}
	}
	if feedback.Warning != nil {
		switch key {
		case "warning.default":
			return feedback.Warning.Default
		case "warning.subtle":
			return feedback.Warning.Subtle
		case "warning.strong":
			return feedback.Warning.Strong
		}
	}
	if feedback.Info != nil {
		switch key {
		case "info.default":
			return feedback.Info.Default
		case "info.subtle":
			return feedback.Info.Subtle
		case "info.strong":
			return feedback.Info.Strong
		}
	}
	return ""
}

// getBodyTypographyFromStruct safely gets body typography from struct
func getBodyTypographyFromStruct(body *schema.BodyTokens, key string) schema.TokenReference {
	if body == nil {
		return ""
	}
	// Return font size reference for simplicity, could be enhanced for other props
	switch key {
	case "large":
		if body.Large != nil {
			return body.Large.FontSize
		}
	case "default":
		if body.Default != nil {
			return body.Default.FontSize
		}
	case "small":
		if body.Small != nil {
			return body.Small.FontSize
		}
	}
	return ""
}

// getLabelTypographyFromStruct safely gets label typography from struct
func getLabelTypographyFromStruct(labels *schema.LabelTokens, key string) schema.TokenReference {
	if labels == nil {
		return ""
	}
	switch key {
	case "large":
		if labels.Large != nil {
			return labels.Large.FontSize
		}
	case "default":
		if labels.Default != nil {
			return labels.Default.FontSize
		}
	case "small":
		if labels.Small != nil {
			return labels.Small.FontSize
		}
	}
	return ""
}

// getCaptionTypographyFromStruct safely gets caption typography from struct
func getCaptionTypographyFromStruct(captions *schema.CaptionTokens, key string) schema.TokenReference {
	if captions == nil {
		return ""
	}
	switch key {
	case "default":
		if captions.Default != nil {
			return captions.Default.FontSize
		}
	case "small":
		if captions.Small != nil {
			return captions.Small.FontSize
		}
	}
	return ""
}

// getComponentSpacingFromStruct safely gets component spacing from struct
func getComponentSpacingFromStruct(component *schema.ComponentSpacing, key string) schema.TokenReference {
	if component == nil {
		return ""
	}
	switch key {
	case "tight":
		return component.Tight
	case "default":
		return component.Default
	case "loose":
		return component.Loose
	default:
		return ""
	}
}

// getInteractiveBorderRadiusFromStruct safely gets interactive border radius from struct
func getInteractiveBorderRadiusFromStruct(radius *schema.InteractiveBorderRadius, key string) schema.TokenReference {
	if radius == nil {
		return ""
	}
	switch key {
	case "small":
		return radius.Small
	case "default":
		return radius.Default
	case "large":
		return radius.Large
	default:
		return ""
	}
}

// getInteractiveShadowFromStruct safely gets interactive shadow from struct
func getInteractiveShadowFromStruct(shadow *schema.InteractiveShadow, key string) schema.TokenReference {
	if shadow == nil {
		return ""
	}
	switch key {
	case "default":
		return shadow.Default
	case "hover":
		return shadow.Hover
	case "focus":
		return shadow.Focus
	default:
		return ""
	}
}

// getButtonTokenFromVariant safely gets button token from variant struct
func getButtonTokenFromVariant(variant *schema.ButtonVariantTokens, prop string) schema.TokenReference {
	if variant == nil {
		return ""
	}
	switch prop {
	case "background":
		return variant.Background
	case "backgroundHover":
		return variant.BackgroundHover
	case "backgroundActive":
		return variant.BackgroundActive
	case "color":
		return variant.Color
	case "colorHover":
		return variant.ColorHover
	case "border":
		return variant.Border
	case "borderHover":
		return variant.BorderHover
	case "borderRadius":
		return variant.BorderRadius
	case "padding":
		return variant.Padding
	case "fontWeight":
		return variant.FontWeight
	case "shadow":
		return variant.Shadow
	case "shadowHover":
		return variant.ShadowHover
	default:
		return ""
	}
}
