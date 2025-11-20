package organisms

import (
	"fmt"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/pkg/utils"
)

// SchemaFormBuilder provides a fluent interface for building forms from schema definitions
type SchemaFormBuilder struct {
	props EnhancedFormProps
}

// NewSchemaFormBuilder creates a new builder for schema-driven forms
func NewSchemaFormBuilder(schema *schema.Schema) *SchemaFormBuilder {
	builder := &SchemaFormBuilder{
		props: EnhancedFormProps{
			Schema:             schema,
			Layout:             FormLayoutVertical,
			Size:               FormSizeMD,
			ValidationStrategy: "realtime",
			AutoSave:           false,
			AutoSaveInterval:   30,
			ShowProgress:       false,
			ShowSaveState:      true,
			Locale:             "en",
			Method:             "POST",
			HXSwap:             "innerHTML",
			ContentType:        "application/json",
		},
	}

	// Initialize from schema
	if schema != nil {
		builder.initializeFromSchema(schema)
	}

	return builder
}

// initializeFromSchema populates form properties from schema definition
func (b *SchemaFormBuilder) initializeFromSchema(s *schema.Schema) {
	b.props.ID = s.ID
	b.props.Title = s.Title
	b.props.Description = s.Description

	if s.Config != nil {
		b.props.SubmitURL = s.Config.Action
		b.props.Method = utils.IfElse(s.Config.Method != "", s.Config.Method, "POST")
		b.props.HXTarget = s.Config.Target
		b.props.ContentType = s.Config.Encoding
		b.props.Headers = s.Config.Headers
	}

	if s.Layout != nil {
		b.props.Layout = mapSchemaLayoutToFormLayout(s.Layout.Type)
		b.props.GridCols = s.Layout.Columns
	}

	if s.Validation != nil {
		b.props.ValidationStrategy = mapSchemaValidationStrategy(s.Validation)
		if s.Validation.ServerURL != "" {
			b.props.ValidationURL = s.Validation.ServerURL
		}
	}

	// Convert schema sections to form sections
	if len(s.Children) > 0 {
		b.props.Sections = b.convertSchemaChildrenToSections(s.Children)
		b.props.ShowProgress = len(s.Children) > 1
	}

	// Security configuration
	if s.Security != nil && s.Security.CSRF != nil && s.Security.CSRF.Enabled {
		b.props.CSRFField = s.Security.CSRF.FieldName
	}

	// HTMX configuration
	if s.HTMX != nil {
		b.props.HXTarget = utils.IfElse(s.HTMX.Target != "", s.HTMX.Target, b.props.HXTarget)
		b.props.HXSwap = utils.IfElse(s.HTMX.Swap != "", s.HTMX.Swap, b.props.HXSwap)
		b.props.HXTrigger = s.HTMX.Trigger
		b.props.HXIndicator = s.HTMX.Indicator
	}
}

// Fluent interface methods

func (b *SchemaFormBuilder) WithID(id string) *SchemaFormBuilder {
	b.props.ID = id
	return b
}

func (b *SchemaFormBuilder) WithTitle(title string) *SchemaFormBuilder {
	b.props.Title = title
	return b
}

func (b *SchemaFormBuilder) WithDescription(description string) *SchemaFormBuilder {
	b.props.Description = description
	return b
}

func (b *SchemaFormBuilder) WithLayout(layout FormLayout) *SchemaFormBuilder {
	b.props.Layout = layout
	return b
}

func (b *SchemaFormBuilder) WithGridColumns(cols int) *SchemaFormBuilder {
	b.props.GridCols = cols
	if cols > 1 {
		b.props.Layout = FormLayoutGrid
	}
	return b
}

func (b *SchemaFormBuilder) WithSize(size FormSize) *SchemaFormBuilder {
	b.props.Size = size
	return b
}

func (b *SchemaFormBuilder) WithValidationStrategy(strategy string) *SchemaFormBuilder {
	b.props.ValidationStrategy = strategy
	return b
}

func (b *SchemaFormBuilder) WithAutoSave(enabled bool, interval int) *SchemaFormBuilder {
	b.props.AutoSave = enabled
	if interval > 0 {
		b.props.AutoSaveInterval = interval
	}
	return b
}

func (b *SchemaFormBuilder) WithProgress(show bool) *SchemaFormBuilder {
	b.props.ShowProgress = show
	return b
}

func (b *SchemaFormBuilder) WithSaveState(show bool) *SchemaFormBuilder {
	b.props.ShowSaveState = show
	return b
}

func (b *SchemaFormBuilder) WithReadOnly(readOnly bool) *SchemaFormBuilder {
	b.props.ReadOnly = readOnly
	return b
}

func (b *SchemaFormBuilder) WithDebug(debug bool) *SchemaFormBuilder {
	b.props.Debug = debug
	return b
}

func (b *SchemaFormBuilder) WithLocale(locale string) *SchemaFormBuilder {
	b.props.Locale = locale
	return b
}

func (b *SchemaFormBuilder) WithSubmitURL(url string) *SchemaFormBuilder {
	b.props.SubmitURL = url
	return b
}

func (b *SchemaFormBuilder) WithValidationURL(url string) *SchemaFormBuilder {
	b.props.ValidationURL = url
	return b
}

func (b *SchemaFormBuilder) WithAutoSaveURL(url string) *SchemaFormBuilder {
	b.props.AutoSaveURL = url
	return b
}

func (b *SchemaFormBuilder) WithMethod(method string) *SchemaFormBuilder {
	b.props.Method = method
	return b
}

func (b *SchemaFormBuilder) WithHTMX(target, swap, trigger, indicator string) *SchemaFormBuilder {
	b.props.HXTarget = target
	b.props.HXSwap = swap
	b.props.HXTrigger = trigger
	b.props.HXIndicator = indicator
	return b
}

func (b *SchemaFormBuilder) WithCSRF(token, field string) *SchemaFormBuilder {
	b.props.CSRFToken = token
	b.props.CSRFField = field
	return b
}

func (b *SchemaFormBuilder) WithTheme(themeID string, tokenOverrides map[string]string, darkMode bool) *SchemaFormBuilder {
	b.props.ThemeID = themeID
	b.props.TokenOverrides = tokenOverrides
	b.props.DarkMode = darkMode
	return b
}

func (b *SchemaFormBuilder) WithInitialData(data map[string]any) *SchemaFormBuilder {
	b.props.InitialData = data
	return b
}

func (b *SchemaFormBuilder) WithClass(class string) *SchemaFormBuilder {
	b.props.Class = class
	return b
}

func (b *SchemaFormBuilder) WithEventHandlers(onSubmit, onValidate, onFieldChange, onAutoSave, onReset string) *SchemaFormBuilder {
	b.props.OnSubmit = onSubmit
	b.props.OnValidate = onValidate
	b.props.OnFieldChange = onFieldChange
	b.props.OnAutoSave = onAutoSave
	b.props.OnReset = onReset
	return b
}

func (b *SchemaFormBuilder) AddAction(action EnhancedFormAction) *SchemaFormBuilder {
	b.props.Actions = append(b.props.Actions, action)
	return b
}

func (b *SchemaFormBuilder) AddSubmitAction(text string, variant atoms.ButtonVariant) *SchemaFormBuilder {
	return b.AddAction(EnhancedFormAction{
		ID:       "submit",
		Type:     "submit",
		Text:     text,
		Variant:  variant,
		Position: "right",
	})
}

func (b *SchemaFormBuilder) AddCancelAction(text string) *SchemaFormBuilder {
	return b.AddAction(EnhancedFormAction{
		ID:       "cancel",
		Type:     "cancel",
		Text:     text,
		Variant:  atoms.ButtonOutline,
		Position: "right",
	})
}

func (b *SchemaFormBuilder) AddSaveDraftAction(text string) *SchemaFormBuilder {
	return b.AddAction(EnhancedFormAction{
		ID:       "save-draft",
		Type:     "save-draft",
		Text:     text,
		Variant:  atoms.ButtonGhost,
		Position: "left",
	})
}

func (b *SchemaFormBuilder) AddCustomAction(id, actionType, text, position string, variant atoms.ButtonVariant, onClick string) *SchemaFormBuilder {
	return b.AddAction(EnhancedFormAction{
		ID:       id,
		Type:     actionType,
		Text:     text,
		Variant:  variant,
		Position: position,
		OnClick:  onClick,
	})
}

// Build methods

func (b *SchemaFormBuilder) Build() EnhancedFormProps {
	// Add default actions if none specified
	if len(b.props.Actions) == 0 {
		b.AddCancelAction("Cancel")
		b.AddSubmitAction("Submit", atoms.ButtonPrimary)
	}

	return b.props
}

func (b *SchemaFormBuilder) BuildAsTemplate() string {
	props := b.Build()
	return fmt.Sprintf("@EnhancedForm(%+v)", props)
}

// Schema conversion helper methods

func (b *SchemaFormBuilder) convertSchemaChildrenToSections(children []schema.Schema) []EnhancedFormSection {
	sections := make([]EnhancedFormSection, 0, len(children))

	for _, child := range children {
		section := EnhancedFormSection{
			ID:          child.ID,
			Name:        child.ID,
			Title:       child.Title,
			Description: child.Description,
			Fields:      b.convertSchemaFieldsToFormFields(child.Fields),
			Order:       len(sections) + 1,
		}

		// Configure section layout
		if child.Layout != nil {
			section.Layout = mapSchemaLayoutToFormLayout(child.Layout.Type)
			section.Columns = child.Layout.Columns
			section.Collapsible = child.Layout.Collapsible
			section.Collapsed = child.Layout.Collapsed
		}

		sections = append(sections, section)
	}

	return sections
}

func (b *SchemaFormBuilder) convertSchemaFieldsToFormFields(fields []schema.Field) []molecules.FormFieldProps {
	formFields := make([]molecules.FormFieldProps, 0, len(fields))

	for _, field := range fields {
		if field.Hidden {
			continue
		}

		formField := b.convertSchemaFieldToFormField(field)
		formFields = append(formFields, formField)
	}

	return formFields
}

func (b *SchemaFormBuilder) convertSchemaFieldToFormField(field schema.Field) molecules.FormFieldProps {
	formField := molecules.FormFieldProps{
		ID:          field.Name,
		Name:        field.Name,
		Label:       field.GetLocalizedLabel(b.props.Locale),
		Type:        mapSchemaFieldTypeToFormFieldType(field.Type),
		Value:       getFieldValueFromData(field.Name, b.props.InitialData),
		Placeholder: field.GetLocalizedPlaceholder(b.props.Locale),
		HelpText:    field.Description,
		Tooltip:     field.Tooltip,
		Icon:        field.Icon,
		Required:    field.Required,
		Disabled:    b.props.ReadOnly || field.Disabled,
		Readonly:    field.Readonly,
		FullWidth:   true,
	}

	// Apply validation rules
	if field.Validation != nil {
		b.applySchemaValidationToFormField(&formField, field.Validation)
	}

	// Apply field constraints
	b.applySchemaConstraintsToFormField(&formField, field)

	// Apply styling and layout
	if field.Style != nil {
		formField.Class = field.Style.Class
	}

	if field.Layout != nil {
		formField.FullWidth = field.Layout.FullWidth
	}

	// Handle field options
	if len(field.Options) > 0 {
		formField.Options = convertSchemaOptionsToFormOptions(field.Options)

		// Set default selection
		for i, option := range field.Options {
			if option.Default {
				formField.Options[i].Selected = true
			}
		}
	}

	// Apply conditional logic
	if field.Conditional != nil && len(field.Conditional.Rules) > 0 {
		formField.Condition = b.buildConditionalExpression(field.Conditional.Rules)
		// Track dependencies
		deps := b.extractConditionalDependencies(field.Conditional.Rules)
		formField.Dependencies = deps
	}

	// Add HTMX integration
	b.applyHTMXIntegrationToFormField(&formField, field)

	// Add Alpine.js integration
	b.applyAlpineIntegrationToFormField(&formField, field)

	// Apply theme integration
	formField.ThemeID = b.props.ThemeID
	formField.TokenOverrides = b.props.TokenOverrides
	formField.DarkMode = b.props.DarkMode

	return formField
}

// Schema mapping functions

func mapSchemaLayoutToFormLayout(layoutType schema.LayoutType) FormLayout {
	switch layoutType {
	case schema.LayoutGrid:
		return FormLayoutGrid
	case schema.LayoutHorizontal:
		return FormLayoutHorizontal
	case schema.LayoutInline:
		return FormLayoutInline
	default:
		return FormLayoutVertical
	}
}

func mapSchemaValidationStrategy(validation *schema.Validation) string {
	if validation.RealTime {
		return "realtime"
	} else if validation.OnBlur {
		return "onblur"
	}
	return "onsubmit"
}

// Field enhancement functions

func (b *SchemaFormBuilder) applySchemaValidationToFormField(formField *molecules.FormFieldProps, validation *schema.FieldValidation) {
	formField.ValidationState = molecules.ValidationStateIdle

	// Add validation rules
	if validation.Required {
		formField.Required = true
	}

	if validation.MinLength > 0 {
		formField.MinLength = validation.MinLength
	}

	if validation.MaxLength > 0 {
		formField.MaxLength = validation.MaxLength
	}

	if validation.Pattern != "" {
		formField.Pattern = validation.Pattern
	}

	// Set validation strategy
	switch b.props.ValidationStrategy {
	case "realtime":
		formField.OnValidate = b.props.ValidationURL
		formField.ValidationDebounce = 300
	case "onblur":
		// Will be handled in Alpine.js
	}

	// Custom validation rules
	if len(validation.Custom) > 0 {
		rules := make([]molecules.ValidationRule, 0, len(validation.Custom))
		for _, customRule := range validation.Custom {
			rules = append(rules, molecules.ValidationRule{
				Type:    customRule.Type,
				Value:   customRule.Value,
				Message: customRule.Message,
			})
		}
		formField.ValidationRules = rules
	}
}

func (b *SchemaFormBuilder) applySchemaConstraintsToFormField(formField *molecules.FormFieldProps, field schema.Field) {
	// Apply type-specific constraints
	switch field.Type {
	case schema.FieldNumber:
		if field.Validation != nil {
			if field.Validation.Min != nil {
				formField.Min = fmt.Sprintf("%v", field.Validation.Min)
			}
			if field.Validation.Max != nil {
				formField.Max = fmt.Sprintf("%v", field.Validation.Max)
			}
			if field.Validation.Step != nil {
				formField.Step = fmt.Sprintf("%v", field.Validation.Step)
			}
		}
	case schema.FieldText:
		if field.Config != nil {
			if rows, ok := field.Config["rows"].(int); ok {
				formField.Rows = rows
			}
			if cols, ok := field.Config["cols"].(int); ok {
				formField.Cols = cols
			}
			if resizable, ok := field.Config["resizable"].(bool); ok {
				formField.Resizable = resizable
			}
		}
	case schema.FieldFile:
		if field.Config != nil {
			if accept, ok := field.Config["accept"].(string); ok {
				formField.Accept = accept
			}
			if maxSize, ok := field.Config["maxSize"].(int64); ok {
				formField.MaxFileSize = maxSize
			}
			if multiple, ok := field.Config["multiple"].(bool); ok {
				formField.MultipleFiles = multiple
			}
		}
	}

	// Apply autocomplete
	if field.Config != nil {
		if autocomplete, ok := field.Config["autocomplete"].(string); ok {
			formField.Autocomplete = autocomplete
		}
	}
}

func (b *SchemaFormBuilder) applyHTMXIntegrationToFormField(formField *molecules.FormFieldProps, field schema.Field) {
	// Real-time validation
	if b.props.ValidationStrategy == "realtime" && b.props.ValidationURL != "" {
		formField.HXPost = b.props.ValidationURL
		formField.HXTarget = fmt.Sprintf("#field-%s-feedback", field.Name)
		formField.HXTrigger = "change, keyup delay:300ms"
	}

	// Data source integration for dynamic options
	if field.DataSource != nil && field.DataSource.URL != "" {
		formField.HXGet = field.DataSource.URL
		formField.HXTrigger = "load, search"
	}

	// Field-specific HTMX configuration
	if field.HTMX != nil {
		if field.HTMX.Get != "" {
			formField.HXGet = field.HTMX.Get
		}
		if field.HTMX.Post != "" {
			formField.HXPost = field.HTMX.Post
		}
		if field.HTMX.Target != "" {
			formField.HXTarget = field.HTMX.Target
		}
		if field.HTMX.Swap != "" {
			formField.HXSwap = field.HTMX.Swap
		}
		if field.HTMX.Trigger != "" {
			formField.HXTrigger = field.HTMX.Trigger
		}
	}
}

func (b *SchemaFormBuilder) applyAlpineIntegrationToFormField(formField *molecules.FormFieldProps, field schema.Field) {
	// Basic Alpine model binding
	formField.AlpineModel = fmt.Sprintf("formData.%s", field.Name)

	// Field change handling
	formField.AlpineChange = fmt.Sprintf("updateField('%s', $event.target.value)", field.Name)

	// Validation strategy integration
	switch b.props.ValidationStrategy {
	case "onblur":
		formField.AlpineBlur = fmt.Sprintf("validateField('%s')", field.Name)
	case "realtime":
		formField.AlpineInput = fmt.Sprintf("updateField('%s', $event.target.value)", field.Name)
	}

	// Field-specific Alpine.js configuration
	if field.Alpine != nil {
		if field.Alpine.Model != "" {
			formField.AlpineModel = field.Alpine.Model
		}
		if field.Alpine.Change != "" {
			formField.AlpineChange = field.Alpine.Change
		}
		if field.Alpine.Blur != "" {
			formField.AlpineBlur = field.Alpine.Blur
		}
		if field.Alpine.Focus != "" {
			formField.AlpineFocus = field.Alpine.Focus
		}
		if field.Alpine.Show != "" {
			formField.AlpineShow = field.Alpine.Show
		}
		if field.Alpine.Init != "" {
			formField.AlpineInit = field.Alpine.Init
		}
	}
}

// Conditional logic functions

func (b *SchemaFormBuilder) buildConditionalExpression(rules []schema.ConditionalRule) string {
	if len(rules) == 0 {
		return ""
	}

	expressions := make([]string, 0, len(rules))

	for _, rule := range rules {
		expr := b.buildSingleConditionalExpression(rule)
		if expr != "" {
			expressions = append(expressions, expr)
		}
	}

	if len(expressions) == 0 {
		return ""
	}

	// Combine with AND logic by default
	return strings.Join(expressions, " && ")
}

func (b *SchemaFormBuilder) buildSingleConditionalExpression(rule schema.ConditionalRule) string {
	fieldRef := fmt.Sprintf("formData.%s", rule.Field)

	switch rule.Operator {
	case "eq", "equals":
		return fmt.Sprintf("%s === '%s'", fieldRef, rule.Value)
	case "ne", "not_equals":
		return fmt.Sprintf("%s !== '%s'", fieldRef, rule.Value)
	case "gt", "greater_than":
		return fmt.Sprintf("%s > %v", fieldRef, rule.Value)
	case "lt", "less_than":
		return fmt.Sprintf("%s < %v", fieldRef, rule.Value)
	case "gte", "greater_than_equals":
		return fmt.Sprintf("%s >= %v", fieldRef, rule.Value)
	case "lte", "less_than_equals":
		return fmt.Sprintf("%s <= %v", fieldRef, rule.Value)
	case "in", "contains":
		if values, ok := rule.Value.([]string); ok {
			quotedValues := make([]string, len(values))
			for i, v := range values {
				quotedValues[i] = fmt.Sprintf("'%s'", v)
			}
			return fmt.Sprintf("[%s].includes(%s)", strings.Join(quotedValues, ", "), fieldRef)
		}
	case "not_in", "not_contains":
		if values, ok := rule.Value.([]string); ok {
			quotedValues := make([]string, len(values))
			for i, v := range values {
				quotedValues[i] = fmt.Sprintf("'%s'", v)
			}
			return fmt.Sprintf("![%s].includes(%s)", strings.Join(quotedValues, ", "), fieldRef)
		}
	case "empty":
		return fmt.Sprintf("!%s || %s === ''", fieldRef, fieldRef)
	case "not_empty":
		return fmt.Sprintf("%s && %s !== ''", fieldRef, fieldRef)
	}

	return ""
}

func (b *SchemaFormBuilder) extractConditionalDependencies(rules []schema.ConditionalRule) []string {
	deps := make([]string, 0, len(rules))
	seen := make(map[string]bool)

	for _, rule := range rules {
		if !seen[rule.Field] {
			deps = append(deps, rule.Field)
			seen[rule.Field] = true
		}
	}

	return deps
}

// Convenience factory methods

// CreateCRUDForm creates a standard CRUD form with common actions
func CreateCRUDForm(schema *schema.Schema) *SchemaFormBuilder {
	return NewSchemaFormBuilder(schema).
		WithValidationStrategy("realtime").
		WithAutoSave(true, 30).
		WithSaveState(true).
		AddSaveDraftAction("Save Draft").
		AddCancelAction("Cancel").
		AddSubmitAction("Save", atoms.ButtonPrimary)
}

// CreateWizardForm creates a multi-step wizard form
func CreateWizardForm(schema *schema.Schema) *SchemaFormBuilder {
	return NewSchemaFormBuilder(schema).
		WithProgress(true).
		WithValidationStrategy("onblur").
		WithAutoSave(false, 0).
		AddCancelAction("Cancel").
		AddSubmitAction("Complete", atoms.ButtonPrimary)
}

// CreateReadOnlyForm creates a read-only form for viewing data
func CreateReadOnlyForm(schema *schema.Schema) *SchemaFormBuilder {
	return NewSchemaFormBuilder(schema).
		WithReadOnly(true).
		WithValidationStrategy("none").
		WithAutoSave(false, 0).
		WithSaveState(false).
		AddCustomAction("close", "button", "Close", "right", atoms.ButtonOutline, "closeForm()")
}

// CreateSearchForm creates a search form with real-time filtering
func CreateSearchForm(schema *schema.Schema) *SchemaFormBuilder {
	return NewSchemaFormBuilder(schema).
		WithValidationStrategy("realtime").
		WithAutoSave(false, 0).
		WithSaveState(false).
		AddCustomAction("clear", "reset", "Clear", "left", atoms.ButtonGhost, "resetForm()").
		AddSubmitAction("Search", atoms.ButtonPrimary)
}
