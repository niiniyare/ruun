package molecules

// import (
// 	"crypto/rand"
// 	"errors"
// 	"fmt"
// 	"html"
// 	"math/big"
// 	"regexp"
// 	"strconv"
// 	"strings"
//
// 	"github.com/niiniyare/ruun/views/components/atoms"
// )
//
// // FormFieldType defines the input type for the form field
// type FormFieldType string
//
// const (
// 	// Basic input types
// 	FormFieldText     FormFieldType = "text"
// 	FormFieldEmail    FormFieldType = "email"
// 	FormFieldPassword FormFieldType = "password"
// 	FormFieldNumber   FormFieldType = "number"
// 	FormFieldSearch   FormFieldType = "search"
// 	FormFieldURL      FormFieldType = "url"
// 	FormFieldTel      FormFieldType = "tel"
//
// 	// Selection types
// 	FormFieldRadio         FormFieldType = "radio"
// 	FormFieldCheckbox      FormFieldType = "checkbox"
// 	FormFieldCheckboxGroup FormFieldType = "checkbox-group"
// 	FormFieldSelect        FormFieldType = "select"
// 	FormFieldMultiSelect   FormFieldType = "multi-select"
// 	FormFieldAutoComplete  FormFieldType = "autocomplete"
//
// 	// Date and time types
// 	FormFieldDate      FormFieldType = "date"
// 	FormFieldTime      FormFieldType = "time"
// 	FormFieldDateTime  FormFieldType = "datetime"
// 	FormFieldDateRange FormFieldType = "date-range"
//
// 	// Complex input types
// 	FormFieldTextarea FormFieldType = "textarea"
// 	FormFieldTags     FormFieldType = "tags"
// 	FormFieldFile     FormFieldType = "file"
// 	FormFieldRange    FormFieldType = "range"
// 	FormFieldColor    FormFieldType = "color"
// )
//
// // FormFieldSize defines the size variants for form fields
// type FormFieldSize string
//
// const (
// 	FormFieldSizeSM FormFieldSize = "sm"
// 	FormFieldSizeMD FormFieldSize = "md"
// 	FormFieldSizeLG FormFieldSize = "lg"
// 	FormFieldSizeXL FormFieldSize = "xl"
// )
//
// // FormFieldVariant defines the visual variant
// type FormFieldVariant string
//
// const (
// 	FormFieldVariantDefault FormFieldVariant = "default"
// 	FormFieldVariantOutline FormFieldVariant = "outline"
// 	FormFieldVariantGhost   FormFieldVariant = "ghost"
// 	FormFieldVariantFilled  FormFieldVariant = "filled"
// )
//
// // Token keys for validation states and field styling
// const (
// 	TokenFieldBackground   = "components.field.background"
// 	TokenFieldBorder       = "components.field.border"
// 	TokenFieldText         = "components.field.text"
// 	TokenFieldLabel        = "components.field.label"
// 	TokenFieldPlaceholder  = "components.field.placeholder"
// 	TokenFieldHelp         = "components.field.help"
// 	TokenValidationLoading = "components.field.validation.loading"
// 	TokenValidationSuccess = "components.field.validation.success"
// 	TokenValidationError   = "components.field.validation.error"
// 	TokenValidationWarning = "components.field.validation.warning"
// 	TokenFieldRadius       = "components.field.radius"
// 	TokenFieldSpacing      = "components.field.spacing"
// 	TokenFieldShadow       = "components.field.shadow"
// )
//
// // SelectOption represents an option for select fields
// type SelectOption struct {
// 	Value       string
// 	Label       string
// 	Description string // Added for richer options
// 	Disabled    bool
// 	Selected    bool
// 	Icon        string // Icon identifier
// 	Group       string // For grouped options
// 	Meta        string // Additional metadata
// }
//
// ValidationRule defines a validation rule
type ValidationRule struct {
	Type    string // required, minLength, maxLength, pattern, custom, etc.
	Value   any
	Message string
}
//
// // FormFieldProps defines the comprehensive properties for the FormField component
// type FormFieldProps struct {
// 	// Core identification
// 	Type    FormFieldType
// 	Size    FormFieldSize
// 	Variant FormFieldVariant
// 	ID      string
// 	Name    string
// 	Label   string
// 	Value   string
//
// 	// Display & UX
// 	Placeholder    string
// 	HelpText       string
// 	ErrorText      string
// 	SuccessText    string
// 	WarningText    string
// 	Icon           string // Leading icon
// 	IconPosition   string // "left" or "right"
// 	Tooltip        string
// 	Description    string // Rich description (supports markdown)
// 	HideLabel      bool   // Accessible label that's visually hidden
// 	LabelPosition  string // "top", "left", "right", "inline"
// 	CharacterCount bool   // Show character count for text inputs
//
// 	// State
// 	Required  bool
// 	Disabled  bool
// 	Readonly  bool
// 	Loading   bool
// 	AutoFocus bool
//
// 	// Validation
// 	ValidationRules []ValidationRule
// 	ValidateOnBlur  bool
// 	ValidateOnInput bool
// 	DebounceMs      int // Debounce validation
//
// 	// NEW: Validation state integration
// 	ValidationState    ValidationState        `json:"validationState,omitempty"`    // idle, validating, valid, invalid, warning
// 	ValidationLoading  bool                   `json:"validationLoading,omitempty"`  // Async validation in progress
// 	ClientRules        *ClientValidationRules `json:"clientRules,omitempty"`        // For immediate client-side validation
// 	OnValidate         string                 `json:"onValidate,omitempty"`         // HTMX validation endpoint
// 	ValidationDebounce int                    `json:"validationDebounce,omitempty"` // Debounce timing in ms
// 	Errors             []string               `json:"errors,omitempty"`             // Current validation errors
//
// 	// Input constraints
// 	MinLength    int
// 	MaxLength    int
// 	Min          string
// 	Max          string
// 	Step         string
// 	Pattern      string
// 	Autocomplete string
//
// 	// Select/Multi-select specific
// 	Options      []SelectOption
// 	Multiple     bool
// 	Searchable   bool
// 	Clearable    bool
// 	Creatable    bool // Allow creating new options
// 	GroupOptions bool // Enable option grouping
//
// 	// Multi-value fields (checkbox groups, multi-select)
// 	Values []string
//
// 	// Layout for checkbox/radio groups
// 	Inline      bool
// 	Columns     int
// 	ColumnsMD   int    // Responsive columns for medium screens
// 	ColumnsLG   int    // Responsive columns for large screens
// 	Orientation string // "horizontal" or "vertical"
// 	ItemSpacing string // spacing class
//
// 	// Textarea specific
// 	Rows       int
// 	Cols       int
// 	Resizable  bool
// 	AutoResize bool // Auto-grow textarea
//
// 	// Date/Time specific
// 	ShowCalendar  bool
// 	Format24      bool
// 	StartDate     string
// 	EndDate       string
// 	MinDate       string
// 	MaxDate       string
// 	DisabledDates []string // ISO date strings
// 	DateFormat    string   // Display format
//
// 	// Autocomplete specific
// 	SearchURL    string
// 	MinChars     int
// 	MaxResults   int
// 	Debounce     int
// 	ShowClear    bool
// 	ShowIcon     bool
// 	FreeForm     bool // Allow non-matching values
// 	RemoteSearch bool
//
// 	// Tags specific
// 	TagsEditable    bool
// 	MaxTags         int
// 	TagValidator    string // Alpine.js expression
// 	TagsOptions     []SelectOption
// 	AllowDuplicates bool
//
// 	// File upload specific
// 	Accept        string // File types
// 	MaxFileSize   int64  // In bytes
// 	MaxFiles      int
// 	ShowPreview   bool
// 	DropZone      bool
// 	MultipleFiles bool
//
// 	// Range/Slider specific
// 	ShowValue     bool
// 	ShowMinMax    bool
// 	Marks         map[int]string // Value labels
// 	RangeMinValue string
// 	RangeMaxValue string
//
// 	// Styling
// 	Class        string
// 	LabelClass   string
// 	InputClass   string
// 	WrapperClass string
// 	ErrorClass   string
// 	HelperClass  string
// 	FullWidth    bool
// 	Rounded      string // "none", "sm", "md", "lg", "full"
//
// 	// HTMX attributes
// 	HXPost      string
// 	HXGet       string
// 	HXPut       string
// 	HXPatch     string
// 	HXDelete    string
// 	HXTarget    string
// 	HXSwap      string
// 	HXTrigger   string
// 	HXIndicator string
// 	HXConfirm   string
// 	HXInclude   string
// 	HXHeaders   string
// 	HXSync      string
// 	HXValidate  string // Custom HTMX validation endpoint
//
// 	// Alpine.js attributes
// 	AlpineModel  string
// 	AlpineChange string
// 	AlpineBlur   string
// 	AlpineFocus  string
// 	AlpineInput  string
// 	AlpineClick  string
// 	AlpineInit   string
// 	AlpineShow   string
// 	AlpineBind   string
//
// 	// Accessibility
// 	AriaLabel       string
// 	AriaDescribedBy string
// 	AriaRequired    string
// 	AriaInvalid     string
// 	TabIndex        int
// 	Role            string
//
// 	// Advanced features
// 	Mask         string // Input mask pattern
// 	Transform    string // "uppercase", "lowercase", "capitalize"
// 	CopyButton   bool   // Add copy to clipboard button
// 	ShowPassword bool   // Toggle for password visibility
// 	GenerateID   bool   // Auto-generate ID if not provided
// 	Prefix       string // Text prefix
// 	Suffix       string // Text suffix
// 	PrefixIcon   string
// 	SuffixIcon   string
//
// 	// Internationalization
// 	Locale string
// 	RTL    bool
//
// 	// NEW: Token-based theming (replaces legacy CSS classes)
// 	Tokens         map[string]string `json:"tokens,omitempty"`         // Resolved token values
// 	ThemeID        string            `json:"themeID,omitempty"`        // Active theme identifier
// 	TokenOverrides map[string]string `json:"tokenOverrides,omitempty"` // Tenant-specific token overrides
// 	DarkMode       bool              `json:"darkMode,omitempty"`       // Dark mode preference
//
// 	// Conditional rendering
// 	Condition    string   // Condition.Condition expression for conditional display
// 	Dependencies []string // Field IDs this field depends on
// }
//
// // ValidationState represents the state of validation.
// type ValidationState string
//
// // Define the possible validation states.
// const (
// 	Idle       ValidationState = "idle"
// 	Validating ValidationState = "validating"
// 	Valid      ValidationState = "valid"
// 	Invalid    ValidationState = "invalid"
// 	Warning    ValidationState = "warning"
// )
//
// // String returns the string representation of ValidationState
// func (v ValidationState) String() string {
// 	return string(v)
// }
//
// // ClientValidationRules represents validation rules that can be executed client-side
// type ClientValidationRules struct {
// 	FieldName    string                 `json:"fieldName"`
// 	Required     bool                   `json:"required"`
// 	MinLength    *int                   `json:"minLength,omitempty"`
// 	MaxLength    *int                   `json:"maxLength,omitempty"`
// 	Min          *float64               `json:"min,omitempty"`
// 	Max          *float64               `json:"max,omitempty"`
// 	Step         *float64               `json:"step,omitempty"`
// 	Pattern      string                 `json:"pattern,omitempty"`
// 	PatternFlags string                 `json:"patternFlags,omitempty"`
// 	Format       string                 `json:"format,omitempty"` // email, url, phone, etc.
// 	CustomRules  []ClientValidationRule `json:"customRules,omitempty"`
// }
//
// // ClientValidationRule represents a single client-side validation rule
// type ClientValidationRule struct {
// 	Type       string `json:"type"`                 // Type of validation
// 	Value      any    `json:"value"`                // Rule value/parameter
// 	Message    string `json:"message"`              // Error message
// 	Expression string `json:"expression,omitempty"` // JS expression for custom rules
// 	Async      bool   `json:"async"`                // Whether this rule requires async validation
// }
//
// // Validation and utility methods
//
// // Validate performs server-side validation based on rules
// func (p FormFieldProps) Validate(value string) error {
// 	// Required validation
// 	if p.Required && strings.TrimSpace(value) == "" {
// 		if p.Label != "" {
// 			return fmt.Errorf("%s is required", p.Label)
// 		}
// 		return errors.New("this field is required")
// 	}
//
// 	// Skip further validation if value is empty and not required
// 	if value == "" {
// 		return nil
// 	}
//
// 	// MinLength validation
// 	if p.MinLength > 0 && len(value) < p.MinLength {
// 		return fmt.Errorf("minimum length is %d characters", p.MinLength)
// 	}
//
// 	// MaxLength validation
// 	if p.MaxLength > 0 && len(value) > p.MaxLength {
// 		return fmt.Errorf("maximum length is %d characters", p.MaxLength)
// 	}
//
// 	// Pattern validation
// 	if p.Pattern != "" {
// 		matched, err := regexp.MatchString(p.Pattern, value)
// 		if err != nil {
// 			return fmt.Errorf("invalid pattern: %v", err)
// 		}
// 		if !matched {
// 			return errors.New("value does not match required format")
// 		}
// 	}
//
// 	// Type-specific validation
// 	switch p.Type {
// 	case FormFieldEmail:
// 		if !isValidEmail(value) {
// 			return errors.New("invalid email address")
// 		}
// 	case FormFieldURL:
// 		if !isValidURL(value) {
// 			return errors.New("invalid URL")
// 		}
// 	case FormFieldNumber:
// 		if _, err := strconv.ParseFloat(value, 64); err != nil {
// 			return errors.New("invalid number")
// 		}
// 	}
//
// 	// Custom validation rules
// 	for _, rule := range p.ValidationRules {
// 		if err := validateRule(rule, value); err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
//
// // ValidateMultiple validates multiple values for multi-select fields
// func (p FormFieldProps) ValidateMultiple(values []string) error {
// 	if p.Required && len(values) == 0 {
// 		return errors.New("at least one option must be selected")
// 	}
//
// 	// Validate each value
// 	for _, value := range values {
// 		if err := p.Validate(value); err != nil {
// 			return err
// 		}
// 	}
//
// 	return nil
// }
//
// // NormalizeDefaults sets default values for optional properties
// func (p *FormFieldProps) NormalizeDefaults() {
// 	// Size defaults
// 	if p.Size == "" {
// 		p.Size = FormFieldSizeMD
// 	}
//
// 	// Variant defaults
// 	if p.Variant == "" {
// 		p.Variant = FormFieldVariantDefault
// 	}
//
// 	// Generate ID if requested
// 	if p.GenerateID && p.ID == "" {
// 		p.ID = generateFieldID(p.Name, p.Type)
// 	}
//
// 	// Columns defaults
// 	if p.Columns == 0 && (p.Type == FormFieldCheckboxGroup || p.Type == FormFieldRadio) {
// 		p.Columns = 1
// 	}
//
// 	// Autocomplete defaults
// 	if p.Type == FormFieldAutoComplete {
// 		if p.MinChars == 0 {
// 			p.MinChars = 3
// 		}
// 		if p.MaxResults == 0 {
// 			p.MaxResults = 10
// 		}
// 		if p.Debounce == 0 {
// 			p.Debounce = 300
// 		}
// 	}
//
// 	// Textarea defaults
// 	if p.Type == FormFieldTextarea && p.Rows == 0 {
// 		p.Rows = 3
// 	}
//
// 	// Date format defaults
// 	if p.Type == FormFieldDate && p.DateFormat == "" {
// 		p.DateFormat = "YYYY-MM-DD"
// 	}
//
// 	// Rounded defaults
// 	if p.Rounded == "" {
// 		p.Rounded = "md"
// 	}
//
// 	// ARIA defaults
// 	if p.AriaDescribedBy == "" && p.ID != "" {
// 		if p.ErrorText != "" {
// 			p.AriaDescribedBy = p.ID + "-error"
// 		} else if p.HelpText != "" {
// 			p.AriaDescribedBy = p.ID + "-help"
// 		}
// 	}
//
// 	// Tab index
// 	if p.Disabled && p.TabIndex == 0 {
// 		p.TabIndex = -1
// 	}
// }
//
// // GetSelectedValue returns the selected value from options
// func (p FormFieldProps) GetSelectedValue() string {
// 	for _, opt := range p.Options {
// 		if opt.Selected {
// 			return opt.Value
// 		}
// 	}
// 	return p.Value
// }
//
// // GetSelectedValues returns all selected values for multi-select
// func (p FormFieldProps) GetSelectedValues() []string {
// 	if p.Values != nil {
// 		return p.Values
// 	}
//
// 	var selected []string
// 	for _, opt := range p.Options {
// 		if opt.Selected {
// 			selected = append(selected, opt.Value)
// 		}
// 	}
// 	return selected
// }
//
// // SanitizeAlpineExpression validates Alpine.js expressions for XSS
// func (p FormFieldProps) SanitizeAlpineExpression(expr string) (string, error) {
// 	// Basic XSS prevention
// 	dangerous := []string{"<script", "javascript:", "onerror=", "onload="}
// 	lowerExpr := strings.ToLower(expr)
//
// 	for _, pattern := range dangerous {
// 		if strings.Contains(lowerExpr, pattern) {
// 			return "", fmt.Errorf("potentially dangerous expression: %s", pattern)
// 		}
// 	}
//
// 	return html.EscapeString(expr), nil
// }
//
// // Helper functions
//
// func isValidEmail(email string) bool {
// 	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
// 	matched, _ := regexp.MatchString(pattern, email)
// 	return matched
// }
//
// func isValidURL(url string) bool {
// 	pattern := `^https?://[^\s/$.?#].[^\s]*$`
// 	matched, _ := regexp.MatchString(pattern, url)
// 	return matched
// }
//
// func validateRule(rule ValidationRule, value string) error {
// 	switch rule.Type {
// 	case "minLength":
// 		if minLen, ok := rule.Value.(int); ok && len(value) < minLen {
// 			if rule.Message != "" {
// 				return errors.New(rule.Message)
// 			}
// 			return fmt.Errorf("minimum length is %d", minLen)
// 		}
// 	case "maxLength":
// 		if maxLen, ok := rule.Value.(int); ok && len(value) > maxLen {
// 			if rule.Message != "" {
// 				return errors.New(rule.Message)
// 			}
// 			return fmt.Errorf("maximum length is %d", maxLen)
// 		}
// 	case "pattern":
// 		if pattern, ok := rule.Value.(string); ok {
// 			matched, err := regexp.MatchString(pattern, value)
// 			if err != nil || !matched {
// 				if rule.Message != "" {
// 					return errors.New(rule.Message)
// 				}
// 				return errors.New("value does not match required format")
// 			}
// 		}
// 	}
// 	return nil
// }
//
// func generateFieldID(name string, fieldType FormFieldType) string {
// 	if name == "" {
// 		return fmt.Sprintf("field-%s-%d", fieldType, generateRandomInt())
// 	}
// 	return fmt.Sprintf("field-%s", strings.ReplaceAll(name, " ", "-"))
// }
//
// // generateRandomInt returns a cryptographically secure random integer between 0 and 999999.
// func generateRandomInt() int {
// 	// Define upper limit (exclusive)
// 	max := big.NewInt(1_000_000) // e.g., generate numbers in [0, 999999]
// 	n, err := rand.Int(rand.Reader, max)
// 	if err != nil {
// 		// Handle error gracefully, or fallback to pseudorandom
// 		panic("failed to generate secure random int: " + err.Error())
// 	}
// 	return int(n.Int64())
// }
//
// // Class generation functions
//
// func formFieldClasses(props FormFieldProps) string {
// 	var classes []string
//
// 	// Base container classes
// 	classes = append(classes, "form-field-wrapper")
//
// 	// Spacing
// 	if !props.Inline {
// 		classes = append(classes, "space-y-2")
// 	}
//
// 	// Full width
// 	if props.FullWidth {
// 		classes = append(classes, "w-full")
// 	}
//
// 	// RTL support
// 	if props.RTL {
// 		classes = append(classes, "rtl")
// 	}
//
// 	// Custom wrapper class
// 	if props.WrapperClass != "" {
// 		classes = append(classes, props.WrapperClass)
// 	}
//
// 	// Custom class
// 	if props.Class != "" {
// 		classes = append(classes, props.Class)
// 	}
//
// 	// Conditional display
// 	if props.Condition != "" {
// 		classes = append(classes, "conditional-field")
// 	}
//
// 	return strings.Join(classes, " ")
// }
//
// func labelClasses(props FormFieldProps) string {
// 	var classes []string
//
// 	// Base classes
// 	classes = append(classes, "block", "text-sm", "font-medium", "leading-none")
//
// 	// Size variants
// 	switch props.Size {
// 	case FormFieldSizeSM:
// 		classes = append(classes, "text-xs")
// 	case FormFieldSizeLG, FormFieldSizeXL:
// 		classes = append(classes, "text-base")
// 	}
//
// 	// State colors
// 	switch {
// 	case props.ErrorText != "":
// 		classes = append(classes, "text-destructive")
// 	case props.SuccessText != "":
// 		classes = append(classes, "text-success")
// 	case props.WarningText != "":
// 		classes = append(classes, "text-warning")
// 	case props.Disabled:
// 		classes = append(classes, "text-muted-foreground", "opacity-50")
// 	default:
// 		classes = append(classes, "text-foreground")
// 	}
//
// 	// Hide label visually but keep for screen readers
// 	if props.HideLabel {
// 		classes = append(classes, "sr-only")
// 	}
//
// 	// Label position
// 	switch props.LabelPosition {
// 	case "left", "right":
// 		classes = append(classes, "inline-block", "mr-2")
// 	case "inline":
// 		classes = append(classes, "inline", "mr-2")
// 	}
//
// 	// Custom label class
// 	if props.LabelClass != "" {
// 		classes = append(classes, props.LabelClass)
// 	}
//
// 	return strings.Join(classes, " ")
// }
//
// func helperTextClasses(props FormFieldProps, textType string) string {
// 	var classes []string
//
// 	// Base classes
// 	classes = append(classes, "text-xs", "mt-1.5")
//
// 	// Type-specific colors
// 	switch textType {
// 	case "error":
// 		classes = append(classes, "text-destructive")
// 		if props.ErrorClass != "" {
// 			classes = append(classes, props.ErrorClass)
// 		}
// 	case "success":
// 		classes = append(classes, "text-success")
// 	case "warning":
// 		classes = append(classes, "text-warning")
// 	default:
// 		classes = append(classes, "text-muted-foreground")
// 		if props.HelperClass != "" {
// 			classes = append(classes, props.HelperClass)
// 		}
// 	}
//
// 	return strings.Join(classes, " ")
// }
//
// func inputWrapperClasses(props FormFieldProps) string {
// 	var classes []string
//
// 	classes = append(classes, "relative")
//
// 	// Add spacing for icons
// 	if props.Icon != "" || props.PrefixIcon != "" {
// 		classes = append(classes, "has-prefix-icon")
// 	}
// 	if props.SuffixIcon != "" || props.CopyButton || props.ShowPassword {
// 		classes = append(classes, "has-suffix-icon")
// 	}
//
// 	return strings.Join(classes, " ")
// }
//
// func roundedClasses(rounded string) string {
// 	switch rounded {
// 	case "none":
// 		return "rounded-none"
// 	case "sm":
// 		return "rounded-sm"
// 	case "md":
// 		return "rounded-md"
// 	case "lg":
// 		return "rounded-lg"
// 	case "full":
// 		return "rounded-full"
// 	default:
// 		return "rounded-md"
// 	}
// }
//
// // Conversion functions for atom components
//
// func convertToRadioOptions(options []SelectOption) []atoms.RadioOption {
// 	if len(options) == 0 {
// 		return []atoms.RadioOption{}
// 	}
//
// 	radioOptions := make([]atoms.RadioOption, 0, len(options))
// 	for _, opt := range options {
// 		radioOptions = append(radioOptions, atoms.RadioOption{
// 			Value:       opt.Value,
// 			Label:       opt.Label,
// 			Description: opt.Description,
// 			Disabled:    opt.Disabled,
// 			Icon:        opt.Icon,
// 		})
// 	}
// 	return radioOptions
// }
//
// func convertToCheckboxOptions(options []SelectOption) []atoms.CheckboxOption {
// 	if len(options) == 0 {
// 		return []atoms.CheckboxOption{}
// 	}
//
// 	checkboxOptions := make([]atoms.CheckboxOption, 0, len(options))
// 	for _, opt := range options {
// 		checkboxOptions = append(checkboxOptions, atoms.CheckboxOption{
// 			Value:       opt.Value,
// 			Label:       opt.Label,
// 			Description: opt.Description,
// 			Disabled:    opt.Disabled,
// 			Checked:     opt.Selected,
// 			Icon:        opt.Icon,
// 		})
// 	}
// 	return checkboxOptions
// }
//
// func convertToSelectOptions(options []SelectOption) []atoms.SelectOption {
// 	if len(options) == 0 {
// 		return []atoms.SelectOption{}
// 	}
//
// 	selectOptions := make([]atoms.SelectOption, 0, len(options))
// 	for _, opt := range options {
// 		selectOptions = append(selectOptions, atoms.SelectOption{
// 			Value:       opt.Value,
// 			Label:       opt.Label,
// 			Description: opt.Description,
// 			Disabled:    opt.Disabled,
// 			Selected:    opt.Selected,
// 			Icon:        opt.Icon,
// 			Group:       opt.Group,
// 		})
// 	}
// 	return selectOptions
// }
//
// func convertToAutoCompleteOptions(options []SelectOption) []atoms.AutoCompleteOption {
// 	if len(options) == 0 {
// 		return []atoms.AutoCompleteOption{}
// 	}
//
// 	autoCompleteOptions := make([]atoms.AutoCompleteOption, 0, len(options))
// 	for _, opt := range options {
// 		autoCompleteOptions = append(autoCompleteOptions, atoms.AutoCompleteOption{
// 			Value:       opt.Value,
// 			Label:       opt.Label,
// 			Description: opt.Description,
// 			Disabled:    opt.Disabled,
// 			Icon:        opt.Icon,
// 			Meta:        opt.Meta,
// 		})
// 	}
// 	return autoCompleteOptions
// }
//
// func convertSelectOptionsToTagProps(options []SelectOption) []atoms.TagProps {
// 	if len(options) == 0 {
// 		return []atoms.TagProps{}
// 	}
//
// 	tags := make([]atoms.TagProps, 0, len(options))
// 	for _, option := range options {
// 		tags = append(tags, atoms.TagProps{
// 			Text:     option.Label,
// 			Value:    option.Value,
// 			Selected: option.Selected,
// 			Disabled: option.Disabled,
// 		})
// 	}
// 	return tags
// }
//
// // Additional helper functions
//
// func getGridColumnsClass(cols, md, lg int) string {
// 	var classes []string
//
// 	if cols > 0 {
// 		classes = append(classes, fmt.Sprintf("grid-cols-%d", cols))
// 	}
// 	if md > 0 {
// 		classes = append(classes, fmt.Sprintf("md:grid-cols-%d", md))
// 	}
// 	if lg > 0 {
// 		classes = append(classes, fmt.Sprintf("lg:grid-cols-%d", lg))
// 	}
//
// 	if len(classes) == 0 {
// 		return "grid-cols-1"
// 	}
//
// 	return strings.Join(classes, " ")
// }
//
// func buildHTMXAttributes(props FormFieldProps) map[string]string {
// 	attrs := make(map[string]string)
//
// 	if props.HXPost != "" {
// 		attrs["hx-post"] = props.HXPost
// 	}
// 	if props.HXGet != "" {
// 		attrs["hx-get"] = props.HXGet
// 	}
// 	if props.HXPut != "" {
// 		attrs["hx-put"] = props.HXPut
// 	}
// 	if props.HXPatch != "" {
// 		attrs["hx-patch"] = props.HXPatch
// 	}
// 	if props.HXDelete != "" {
// 		attrs["hx-delete"] = props.HXDelete
// 	}
// 	if props.HXTarget != "" {
// 		attrs["hx-target"] = props.HXTarget
// 	}
// 	if props.HXSwap != "" {
// 		attrs["hx-swap"] = props.HXSwap
// 	}
// 	if props.HXTrigger != "" {
// 		attrs["hx-trigger"] = props.HXTrigger
// 	}
// 	if props.HXIndicator != "" {
// 		attrs["hx-indicator"] = props.HXIndicator
// 	}
// 	if props.HXConfirm != "" {
// 		attrs["hx-confirm"] = props.HXConfirm
// 	}
// 	if props.HXInclude != "" {
// 		attrs["hx-include"] = props.HXInclude
// 	}
// 	if props.HXHeaders != "" {
// 		attrs["hx-headers"] = props.HXHeaders
// 	}
// 	if props.HXSync != "" {
// 		attrs["hx-sync"] = props.HXSync
// 	}
// 	if props.HXValidate != "" {
// 		attrs["hx-validate"] = props.HXValidate
// 	}
//
// 	return attrs
// }
//
// func buildAlpineAttributes(props FormFieldProps) map[string]string {
// 	attrs := make(map[string]string)
//
// 	if props.AlpineModel != "" {
// 		attrs["x-model"] = props.AlpineModel
// 	}
// 	if props.AlpineChange != "" {
// 		attrs["x-on:change"] = props.AlpineChange
// 	}
// 	if props.AlpineBlur != "" {
// 		attrs["x-on:blur"] = props.AlpineBlur
// 	}
// 	if props.AlpineFocus != "" {
// 		attrs["x-on:focus"] = props.AlpineFocus
// 	}
// 	if props.AlpineInput != "" {
// 		attrs["x-on:input"] = props.AlpineInput
// 	}
// 	if props.AlpineClick != "" {
// 		attrs["x-on:click"] = props.AlpineClick
// 	}
// 	if props.AlpineInit != "" {
// 		attrs["x-init"] = props.AlpineInit
// 	}
// 	if props.AlpineShow != "" {
// 		attrs["x-show"] = props.AlpineShow
// 	}
// 	if props.AlpineBind != "" {
// 		attrs["x-bind"] = props.AlpineBind
// 	}
//
// 	return attrs
// }
//
// // groupOptionsByName groups options by their Group field
// func groupOptionsByName(options []SelectOption) map[string][]SelectOption {
// 	if len(options) == 0 {
// 		return make(map[string][]SelectOption)
// 	}
//
// 	groups := make(map[string][]SelectOption)
// 	for _, option := range options {
// 		group := option.Group
// 		if group == "" {
// 			group = ""
// 		}
// 		groups[group] = append(groups[group], option)
// 	}
// 	return groups
// }
