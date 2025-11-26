// pkg/schema/field_v2.go
package schema

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Field represents a form input element
type Field struct {
	// Identity
	Name  string    `json:"name"`
	Type  FieldType `json:"type"`
	Label string    `json:"label"`

	// Display
	Description string `json:"description,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Help        string `json:"help,omitempty"`
	Tooltip     string `json:"tooltip,omitempty"`
	Icon        string `json:"icon,omitempty"`

	// State
	Required bool `json:"required,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Readonly bool `json:"readonly,omitempty"`
	Hidden   bool `json:"hidden,omitempty"`

	// Values
	Value   any           `json:"value,omitempty"`
	Default any           `json:"default,omitempty"`
	Options []FieldOption `json:"options,omitempty"`

	// Behavior
	Validation   *FieldValidation  `json:"validation,omitempty"`
	Transform    *FieldTransform   `json:"transform,omitempty"`
	Conditional  *Conditional `json:"conditional,omitempty"`
	DataSource   *DataSource       `json:"dataSource,omitempty"`
	Dependencies []string          `json:"dependencies,omitempty"`

	// Layout & Styling
	Layout *FieldLayout   `json:"layout,omitempty"`
	Style  *Style    `json:"style,omitempty"`
	Config map[string]any `json:"config,omitempty"`

	// Permissions
	RequirePermission string   `json:"requirePermission,omitempty"`
	RequireRoles      []string `json:"requireRoles,omitempty"`

	// Framework Integration
	Behavior *Behavior `json:"behavior,omitempty"`
	Binding  *Binding  `json:"binding,omitempty"`

	// Internationalization
	I18n *FieldI18n `json:"i18n,omitempty"`

	// Runtime (populated by enricher)
	Runtime *FieldRuntime `json:"runtime,omitempty"`

	// Internal
	evaluator *condition.Evaluator `json:"-"`
}

// FieldType defines all supported input and display field types
type FieldType string

const (
	// =========================================================================
	// Basic Input Fields
	// =========================================================================

	// Text inputs
	FieldText     FieldType = "text"     // Single-line text input
	FieldEmail    FieldType = "email"    // Email input with validation
	FieldPassword FieldType = "password" // Password input (masked)
	FieldNumber   FieldType = "number"   // Numeric input
	FieldPhone    FieldType = "phone"    // Phone number input
	FieldURL      FieldType = "url"      // URL input with validation
	FieldHidden   FieldType = "hidden"   // Hidden input field

	// =========================================================================
	// Date & Time Fields
	// =========================================================================

	FieldDate      FieldType = "date"      // Date picker
	FieldTime      FieldType = "time"      // Time picker
	FieldDateTime  FieldType = "datetime"  // Date and time picker
	FieldDateRange FieldType = "daterange" // Date range picker
	FieldMonth     FieldType = "month"     // Month picker (YYYY-MM)
	FieldYear      FieldType = "year"      // Year picker
	FieldQuarter   FieldType = "quarter"   // Quarter picker (Q1-Q4)

	// =========================================================================
	// Rich Content Fields
	// =========================================================================

	FieldTextarea FieldType = "textarea" // Multi-line text area
	FieldRichText FieldType = "richtext" // WYSIWYG rich text editor
	FieldCode     FieldType = "code"     // Code editor with syntax highlighting
	FieldJSON     FieldType = "json"     // JSON editor with validation

	// =========================================================================
	// Selection Fields
	// =========================================================================

	// Basic selection
	FieldSelect      FieldType = "select"      // Single selection dropdown
	FieldMultiSelect FieldType = "multiselect" // Multiple selection dropdown
	FieldRadio       FieldType = "radio"       // Radio button group
	FieldCheckbox    FieldType = "checkbox"    // Single toggle checkbox
	FieldCheckboxes  FieldType = "checkboxes"  // Multiple checkbox group

	// Advanced selection
	FieldTreeSelect FieldType = "treeselect" // Hierarchical tree selection
	FieldCascader   FieldType = "cascader"   // Cascading dropdown selection
	FieldTransfer   FieldType = "transfer"   // Dual-pane transfer list

	// =========================================================================
	// Interactive Controls
	// =========================================================================

	FieldSwitch FieldType = "switch" // Toggle switch
	FieldSlider FieldType = "slider" // Range slider
	FieldRating FieldType = "rating" // Star rating input
	FieldColor  FieldType = "color"  // Color picker

	// =========================================================================
	// File & Media Upload Fields
	// =========================================================================

	FieldFile      FieldType = "file"      // Generic file upload
	FieldImage     FieldType = "image"     // Image upload with preview
	FieldVideo     FieldType = "video"     // Video file upload
	FieldAudio     FieldType = "audio"     // Audio file upload
	FieldSignature FieldType = "signature" // Digital signature pad

	// =========================================================================
	// Specialized Input Fields
	// =========================================================================

	FieldCurrency     FieldType = "currency"     // Currency input with formatting
	FieldTags         FieldType = "tags"         // Tag input with auto-complete
	FieldLocation     FieldType = "location"     // Address/location picker
	FieldRelation     FieldType = "relation"     // Foreign key relationship selector
	FieldAutoComplete FieldType = "autocomplete" // Auto-complete search input
	FieldIconPicker   FieldType = "icon-picker"  // Icon selector
	FieldFormula      FieldType = "formula"      // Computed/calculated field

	// =========================================================================
	// Display & Read-only Fields
	// =========================================================================

	FieldDisplay FieldType = "display" // Read-only text display
	FieldDivider FieldType = "divider" // Visual separator line
	FieldHTML    FieldType = "html"    // Raw HTML content display
	FieldStatic  FieldType = "static"  // Static text display

	// =========================================================================
	// Layout & Container Fields
	// =========================================================================

	FieldGroup    FieldType = "group"    // Field grouping container
	FieldFieldset FieldType = "fieldset" // HTML fieldset container
	FieldTabs     FieldType = "tabs"     // Tabbed interface container
	FieldPanel    FieldType = "panel"    // Panel container with header
	FieldCollapse FieldType = "collapse" // Collapsible section container

	// =========================================================================
	// Collection & Repeating Fields
	// =========================================================================

	FieldRepeatable    FieldType = "repeatable"     // Repeatable field groups
	FieldTableRepeater FieldType = "table_repeater" // Table-style repeatable fields
)

// FieldValidation defines validation rules
type FieldValidation struct {
	// String validation
	MinLength *int   `json:"minLength,omitempty"`
	MaxLength *int   `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Format    string `json:"format,omitempty"`

	// Number validation
	Min      *float64 `json:"min,omitempty"`
	Max      *float64 `json:"max,omitempty"`
	Step     *float64 `json:"step,omitempty"`
	Integer  bool     `json:"integer,omitempty"`
	Positive bool     `json:"positive,omitempty"`

	// Array validation
	MinItems *int `json:"minItems,omitempty"`
	MaxItems *int `json:"maxItems,omitempty"`
	Unique   bool `json:"unique,omitempty"`

	// File validation
	MaxSize     int64    `json:"maxSize,omitempty"`
	AllowedMIME []string `json:"allowedMime,omitempty"`

	// Custom
	Custom   string              `json:"custom,omitempty"`
	Messages *ValidationMessages `json:"messages,omitempty"`
}

// DataSource defines where to fetch dynamic options


// FieldTransform defines value transformation
type FieldTransform struct {
	Type   TransformType  `json:"type"`
	Params map[string]any `json:"params,omitempty"`
}



// FieldLayout controls positioning
type FieldLayout struct {
	Row     int    `json:"row,omitempty"`
	Column  int    `json:"column,omitempty"`
	ColSpan int    `json:"colSpan,omitempty"`
	Width   string `json:"width,omitempty"`
	Order   int    `json:"order,omitempty"`
}



// FieldRuntime holds runtime state from enricher
type FieldRuntime struct {
	Visible  bool   `json:"visible"`
	Editable bool   `json:"editable"`
	Reason   string `json:"reason,omitempty"`
}

// IsVisible returns whether the field is visible
func (fr *FieldRuntime) IsVisible() bool {
	return fr != nil && fr.Visible
}

// IsEditable returns whether the field is editable
func (fr *FieldRuntime) IsEditable() bool {
	return fr != nil && fr.Editable
}

// FieldI18n holds translations
type FieldI18n struct {
	Label       map[string]string `json:"label,omitempty"`
	Placeholder map[string]string `json:"placeholder,omitempty"`
	Help        map[string]string `json:"help,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	Tooltip     map[string]string `json:"tooltip,omitempty"`
}


// Core Methods

// SetEvaluator sets the condition evaluator
func (f *Field) SetEvaluator(evaluator *condition.Evaluator) {
	f.evaluator = evaluator
}

// GetEvaluator returns the condition evaluator
func (f *Field) GetEvaluator() *condition.Evaluator {
	return f.evaluator
}

// Validate validates field configuration
func (f *Field) Validate(ctx context.Context) error {
	validator := newFieldValidator()
	return validator.Validate(ctx, f)
}

// IsVisible checks if field should be displayed
func (f *Field) IsVisible(ctx context.Context, data map[string]any) (bool, error) {
	if f.Hidden {
		return false, nil
	}

	if f.Conditional == nil {
		return true, nil
	}

	// Check Hide condition first
	if f.Conditional.Hide != nil {
		if f.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		hide, err := f.evaluator.Evaluate(ctx, f.Conditional.Hide, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate hide condition: %w", err)
		}
		if hide {
			return false, nil
		}
	}

	// Check Show condition
	if f.Conditional.Show != nil {
		if f.evaluator == nil {
			return true, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		show, err := f.evaluator.Evaluate(ctx, f.Conditional.Show, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate show condition: %w", err)
		}
		return show, nil
	}

	return true, nil
}

// IsRequired checks if field is required
func (f *Field) IsRequired(ctx context.Context, data map[string]any) (bool, error) {
	if f.Required {
		return true, nil
	}

	if f.Conditional != nil && f.Conditional.Required != nil {
		if f.evaluator == nil {
			return false, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		required, err := f.evaluator.Evaluate(ctx, f.Conditional.Required, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate required condition: %w", err)
		}
		return required, nil
	}

	return false, nil
}

// IsDisabled checks if field is disabled
func (f *Field) IsDisabled(ctx context.Context, data map[string]any) (bool, error) {
	if f.Disabled {
		return true, nil
	}

	if f.Conditional != nil && f.Conditional.Disabled != nil {
		if f.evaluator == nil {
			return false, nil
		}
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		disabled, err := f.evaluator.Evaluate(ctx, f.Conditional.Disabled, evalCtx)
		if err != nil {
			return false, fmt.Errorf("evaluate disabled condition: %w", err)
		}
		return disabled, nil
	}

	return false, nil
}

// ValidateValue validates a field value
func (f *Field) ValidateValue(ctx context.Context, value any) error {
	validator := newFieldValueValidator()
	return validator.Validate(ctx, f, value)
}

// ApplyTransform applies transformation to value
func (f *Field) ApplyTransform(value any) (any, error) {
	if f.Transform == nil {
		return value, nil
	}

	str, ok := value.(string)
	if !ok {
		return value, nil
	}

	transformer := newFieldTransformer()
	return transformer.Transform(f.Transform, str)
}

// GetDefaultValue returns the default value
func (f *Field) GetDefaultValue(ctx context.Context) (any, error) {
	if f.Default == nil {
		return f.getTypeDefaultValue(), nil
	}
	return f.Default, nil
}

// getTypeDefaultValue returns type-specific defaults
func (f *Field) getTypeDefaultValue() any {
	switch f.Type {
	case FieldCheckbox, FieldSwitch:
		return false
	case FieldMultiSelect, FieldCheckboxes, FieldTags:
		return []string{}
	case FieldNumber, FieldCurrency, FieldSlider, FieldRating:
		return 0
	case FieldDate:
		return time.Now().Format("2006-01-02")
	case FieldDateTime:
		return time.Now().Format(time.RFC3339)
	case FieldTime:
		return time.Now().Format("15:04:05")
	default:
		return ""
	}
}

// Localization Methods

// GetLocalizedLabel returns localized label
func (f *Field) GetLocalizedLabel(locale string) string {
	if f.I18n != nil && f.I18n.Label != nil {
		if label, ok := f.I18n.Label[locale]; ok && label != "" {
			return label
		}
	}
	return T_Field(locale, f.Name)
}

// GetLocalizedPlaceholder returns localized placeholder
func (f *Field) GetLocalizedPlaceholder(locale string) string {
	if f.I18n != nil && f.I18n.Placeholder != nil {
		if placeholder, ok := f.I18n.Placeholder[locale]; ok && placeholder != "" {
			return placeholder
		}
	}
	key := f.Name + "Placeholder"
	translated := T_Field(locale, key)
	if translated == humanizeFieldID(key) {
		return f.Placeholder
	}
	return translated
}

// GetLocalizedHelp returns localized help text
func (f *Field) GetLocalizedHelp(locale string) string {
	if f.I18n != nil && f.I18n.Help != nil {
		if help, ok := f.I18n.Help[locale]; ok && help != "" {
			return help
		}
	}
	key := f.Name + "Help"
	translated := T_Field(locale, key)
	if translated == humanizeFieldID(key) {
		return f.Help
	}
	return translated
}

// GetLocalizedDescription returns localized description
func (f *Field) GetLocalizedDescription(locale string) string {
	if f.I18n != nil && f.I18n.Description != nil {
		if desc, ok := f.I18n.Description[locale]; ok && desc != "" {
			return desc
		}
	}
	key := f.Name + "Description"
	translated := T_Field(locale, key)
	if translated == humanizeFieldID(key) {
		return f.Description
	}
	return translated
}

// GetLocalizedTooltip returns localized tooltip text
func (f *Field) GetLocalizedTooltip(locale string) string {
	if f.I18n != nil && f.I18n.Tooltip != nil {
		if tooltip, ok := f.I18n.Tooltip[locale]; ok && tooltip != "" {
			return tooltip
		}
	}
	key := f.Name + "Tooltip"
	translated := T_Field(locale, key)
	if translated == humanizeFieldID(key) {
		return f.Tooltip
	}
	return translated
}

// Type Checking Methods

// IsSelectionType checks if field is a selection type
func (f *Field) IsSelectionType() bool {
	switch f.Type {
	case FieldSelect, FieldMultiSelect, FieldRadio, FieldCheckboxes:
		return true
	}
	return false
}

// IsFileType checks if field handles files
func (f *Field) IsFileType() bool {
	switch f.Type {
	case FieldFile, FieldImage, FieldSignature:
		return true
	}
	return false
}

// IsNumericType checks if field is numeric
func (f *Field) IsNumericType() bool {
	switch f.Type {
	case FieldNumber, FieldCurrency, FieldSlider, FieldRating:
		return true
	}
	return false
}

// IsDateTimeType checks if field is date/time related
func (f *Field) IsDateTimeType() bool {
	switch f.Type {
	case FieldDate, FieldTime, FieldDateTime:
		return true
	}
	return false
}

// IsLayoutType checks if field is a layout/container type
func (f *Field) IsLayoutType() bool {
	switch f.Type {
	case FieldGroup, FieldFieldset, FieldTabs, FieldPanel, FieldCollapse:
		return true
	}
	return false
}

// GetOptionLabel returns label for an option value
func (f *Field) GetOptionLabel(value string) string {
	for _, option := range f.Options {
		if option.Value == value {
			return option.Label
		}
	}
	return value
}

// GetLabel returns the field label or a humanized version of the name
func (f *Field) GetLabel() string {
	if f.Label != "" {
		return f.Label
	}
	return humanizeFieldID(f.Name)
}

// GetPlaceholder returns the field placeholder
func (f *Field) GetPlaceholder() string {
	return f.Placeholder
}

// GetHelp returns the field help text
func (f *Field) GetHelp() string {
	return f.Help
}

// TextConfig returns text-specific configuration
func (f *Field) TextConfig() map[string]any {
	if f.Config == nil {
		return make(map[string]any)
	}
	
	config := make(map[string]any)
	textFields := []string{"maxLength", "minLength", "autocomplete", "prefix", "suffix"}
	
	for _, field := range textFields {
		if val, ok := f.Config[field]; ok {
			config[field] = val
		}
	}
	
	return config
}

// SelectConfig returns select-specific configuration
func (f *Field) SelectConfig() map[string]any {
	if f.Config == nil {
		return make(map[string]any)
	}
	
	config := make(map[string]any)
	selectFields := []string{"multiple", "searchable", "clearable", "createable", "maxSelections", "placeholder"}
	
	for _, field := range selectFields {
		if val, ok := f.Config[field]; ok {
			config[field] = val
		}
	}
	
	return config
}

// FileConfig returns file-specific configuration
func (f *Field) FileConfig() map[string]any {
	if f.Config == nil {
		return make(map[string]any)
	}
	
	config := make(map[string]any)
	fileFields := []string{"accept", "maxSize", "multiple", "previewMode", "autoUpload"}
	
	for _, field := range fileFields {
		if val, ok := f.Config[field]; ok {
			config[field] = val
		}
	}
	
	return config
}

// RelationConfig returns relation-specific configuration
func (f *Field) RelationConfig() map[string]any {
	if f.Config == nil {
		return make(map[string]any)
	}
	
	config := make(map[string]any)
	relationFields := []string{"entity", "displayField", "valueField", "searchField", "targetSchema", "multiple", "createNew"}
	
	for _, field := range relationFields {
		if val, ok := f.Config[field]; ok {
			config[field] = val
		}
	}
	
	return config
}

// fieldValidator validates field configuration
type fieldValidator struct{}

func newFieldValidator() *fieldValidator {
	return &fieldValidator{}
}

func (v *fieldValidator) Validate(ctx context.Context, f *Field) error {
	collector := NewErrorCollector()

	if f.Name == "" {
		collector.AddError(ErrInvalidFieldName)
	}
	if f.Type == "" {
		collector.AddError(ErrInvalidFieldType)
	}
	if f.Label == "" {
		collector.AddError(ErrInvalidFieldLabel)
	}

	// Validate validation rules
	if f.Validation != nil {
		if err := v.validateValidationRules(f); err != nil {
			collector.AddError(err.(SchemaError))
		}
	}

	// Validate selection fields have options
	if f.IsSelectionType() && len(f.Options) == 0 && f.DataSource == nil {
		collector.AddValidationError(f.Name, "missing_options",
			"selection field requires options or dataSource")
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

func (v *fieldValidator) validateValidationRules(f *Field) error {
	val := f.Validation

	if val.Min != nil && val.Max != nil && *val.Min > *val.Max {
		return NewValidationError("invalid_range",
			"min cannot be greater than max").WithField(f.Name)
	}

	if val.MinLength != nil && val.MaxLength != nil && *val.MinLength > *val.MaxLength {
		return NewValidationError("invalid_length_range",
			"minLength cannot be greater than maxLength").WithField(f.Name)
	}

	if val.Pattern != "" {
		if _, err := regexp.Compile(val.Pattern); err != nil {
			return NewValidationError("invalid_pattern",
				"invalid regex pattern").WithField(f.Name)
		}
	}

	return nil
}

// fieldValueValidator validates field values
type fieldValueValidator struct{}

func newFieldValueValidator() *fieldValueValidator {
	return &fieldValueValidator{}
}

func (v *fieldValueValidator) Validate(ctx context.Context, f *Field, value any) error {
	// Required check
	if f.Required && isEmpty(value) {
		return v.createError(f, "required", "%s is required", f.Label)
	}

	if isEmpty(value) {
		return nil
	}

	if f.Validation == nil {
		return nil
	}

	// Type-specific validation
	switch f.Type {
	case FieldText, FieldTextarea, FieldPassword:
		return v.validateString(f, value)
	case FieldNumber, FieldCurrency:
		return v.validateNumber(f, value)
	case FieldEmail:
		return v.validateEmail(f, value)
	case FieldURL:
		return v.validateURL(f, value)
	case FieldPhone:
		return v.validatePhone(f, value)
	}

	return nil
}

func (v *fieldValueValidator) validateString(f *Field, value any) error {
	str, ok := value.(string)
	if !ok {
		return v.createError(f, "invalid_type", "must be a string")
	}

	val := f.Validation

	if val.MinLength != nil && len(str) < *val.MinLength {
		return v.createError(f, "min_length",
			"%s must be at least %d characters", f.Label, *val.MinLength)
	}

	if val.MaxLength != nil && len(str) > *val.MaxLength {
		return v.createError(f, "max_length",
			"%s must be at most %d characters", f.Label, *val.MaxLength)
	}

	if val.Pattern != "" {
		matched, _ := regexp.MatchString(val.Pattern, str)
		if !matched {
			msg := "invalid format"
			if val.Messages != nil && val.Messages.Pattern != "" {
				msg = val.Messages.Pattern
			}
			return v.createError(f, "pattern", "%s", msg)
		}
	}

	return nil
}

func (v *fieldValueValidator) validateNumber(f *Field, value any) error {
	num, ok := toFloat64(value)
	if !ok {
		return v.createError(f, "invalid_type", "must be a number")
	}

	val := f.Validation

	if val.Min != nil && num < *val.Min {
		return v.createError(f, "min_value",
			"%s must be at least %v", f.Label, *val.Min)
	}

	if val.Max != nil && num > *val.Max {
		return v.createError(f, "max_value",
			"%s must be at most %v", f.Label, *val.Max)
	}

	if val.Integer && num != float64(int(num)) {
		return v.createError(f, "not_integer", "%s must be an integer", f.Label)
	}

	return nil
}

func (v *fieldValueValidator) validateEmail(f *Field, value any) error {
	str, ok := value.(string)
	if !ok {
		return v.createError(f, "invalid_type", "must be a string")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, str)
	if !matched {
		return v.createError(f, "invalid_email", "must be a valid email address")
	}

	return nil
}

func (v *fieldValueValidator) validateURL(f *Field, value any) error {
	str, ok := value.(string)
	if !ok {
		return v.createError(f, "invalid_type", "must be a string")
	}

	urlRegex := `^https?://[^\s/$.?#].[^\s]*$`
	matched, _ := regexp.MatchString(urlRegex, str)
	if !matched {
		return v.createError(f, "invalid_url", "must be a valid URL")
	}

	return nil
}

func (v *fieldValueValidator) validatePhone(f *Field, value any) error {
	str, ok := value.(string)
	if !ok {
		return v.createError(f, "invalid_type", "must be a string")
	}

	// Basic phone validation
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(str, "")
	if len(cleaned) < 7 || len(cleaned) > 15 {
		return v.createError(f, "invalid_phone", "must be a valid phone number")
	}

	return nil
}

func (v *fieldValueValidator) createError(f *Field, code, format string, args ...any) SchemaError {
	message := fmt.Sprintf(format, args...)
	return NewValidationError(code, message).WithField(f.Name)
}

// fieldTransformer handles value transformations
type fieldTransformer struct{}

func newFieldTransformer() *fieldTransformer {
	return &fieldTransformer{}
}

func (t *fieldTransformer) Transform(transform *FieldTransform, value string) (any, error) {
	switch transform.Type {
	case TransformUpperCase:
		return strings.ToUpper(value), nil
	case TransformLowerCase:
		return strings.ToLower(value), nil
	case TransformTrim:
		return strings.TrimSpace(value), nil
	case TransformCapitalize:
		if len(value) == 0 {
			return value, nil
		}
		return strings.ToUpper(string(value[0])) + strings.ToLower(value[1:]), nil
	case TransformSlugify:
		return t.slugify(value), nil
	default:
		return value, fmt.Errorf("unknown transform type: %s", transform.Type)
	}
}

func (t *fieldTransformer) slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	reg := regexp.MustCompile("[^a-z0-9-]+")
	s = reg.ReplaceAllString(s, "")
	reg = regexp.MustCompile("-+")
	s = reg.ReplaceAllString(s, "-")
	return strings.Trim(s, "-")
}

// FieldBuilder provides fluent API for building fields
type FieldBuilder struct {
	field     Field
	evaluator *condition.Evaluator
}

// NewField starts building a field
func NewField(name string, fieldType FieldType) *FieldBuilder {
	return &FieldBuilder{
		field: Field{
			Name: name,
			Type: fieldType,
		},
	}
}

func (b *FieldBuilder) WithLabel(label string) *FieldBuilder {
	b.field.Label = label
	return b
}

func (b *FieldBuilder) WithDescription(desc string) *FieldBuilder {
	b.field.Description = desc
	return b
}

func (b *FieldBuilder) WithPlaceholder(placeholder string) *FieldBuilder {
	b.field.Placeholder = placeholder
	return b
}

func (b *FieldBuilder) WithHelp(help string) *FieldBuilder {
	b.field.Help = help
	return b
}

func (b *FieldBuilder) Required() *FieldBuilder {
	b.field.Required = true
	return b
}

func (b *FieldBuilder) Disabled() *FieldBuilder {
	b.field.Disabled = true
	return b
}

func (b *FieldBuilder) Readonly() *FieldBuilder {
	b.field.Readonly = true
	return b
}

func (b *FieldBuilder) Hidden() *FieldBuilder {
	b.field.Hidden = true
	return b
}

func (b *FieldBuilder) WithDefault(value any) *FieldBuilder {
	b.field.Default = value
	return b
}

func (b *FieldBuilder) WithValue(value any) *FieldBuilder {
	b.field.Value = value
	return b
}

func (b *FieldBuilder) WithOptions(options []FieldOption) *FieldBuilder {
	b.field.Options = options
	return b
}

func (b *FieldBuilder) WithValidation(validation *FieldValidation) *FieldBuilder {
	b.field.Validation = validation
	return b
}

func (b *FieldBuilder) WithTransform(transform *FieldTransform) *FieldBuilder {
	b.field.Transform = transform
	return b
}

func (b *FieldBuilder) WithConditional(conditional *Conditional) *FieldBuilder {
	b.field.Conditional = conditional
	return b
}

func (b *FieldBuilder) WithDataSource(source *DataSource) *FieldBuilder {
	b.field.DataSource = source
	return b
}

func (b *FieldBuilder) WithDependencies(deps ...string) *FieldBuilder {
	b.field.Dependencies = append(b.field.Dependencies, deps...)
	return b
}

func (b *FieldBuilder) WithLayout(layout *FieldLayout) *FieldBuilder {
	b.field.Layout = layout
	return b
}

func (b *FieldBuilder) WithStyle(style *Style) *FieldBuilder {
	b.field.Style = style
	return b
}

func (b *FieldBuilder) WithConfig(config map[string]any) *FieldBuilder {
	b.field.Config = config
	return b
}

func (b *FieldBuilder) WithPermission(permission string) *FieldBuilder {
	b.field.RequirePermission = permission
	return b
}

func (b *FieldBuilder) WithRoles(roles ...string) *FieldBuilder {
	b.field.RequireRoles = roles
	return b
}

func (b *FieldBuilder) WithEvaluator(evaluator *condition.Evaluator) *FieldBuilder {
	b.evaluator = evaluator
	return b
}

func (b *FieldBuilder) Build() Field {
	if b.evaluator != nil {
		b.field.SetEvaluator(b.evaluator)
	}
	return b.field
}

// Utility functions

func isEmpty(value any) bool {
	if value == nil {
		return true
	}
	switch v := value.(type) {
	case string:
		return strings.TrimSpace(v) == ""
	case []any:
		return len(v) == 0
	case []string:
		return len(v) == 0
	default:
		return false
	}
}

func toFloat64(value any) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case string:
		// Try to parse string as float
		return 0, false
	default:
		return 0, false
	}
}

func humanizeFieldID(id string) string {
	// Convert snake_case or camelCase to Title Case
	words := regexp.MustCompile(`[_\s]+|(?=[A-Z])`).Split(id, -1)

	// Create a title caser
	titleCaser := cases.Title(language.English)

	for i, word := range words {
		if len(word) > 0 {
			words[i] = titleCaser.String(strings.ToLower(word))
		}
	}
	return strings.Join(words, " ")
}
