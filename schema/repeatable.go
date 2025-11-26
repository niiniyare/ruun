package schema

import (
	"context"
	"fmt"
	"strings"

	"github.com/niiniyare/ruun/pkg/condition"
)

// RepeatableField represents a field that can have multiple instances
type RepeatableField struct {
	Field // Embed base field

	// Template
	Template []Field `json:"template"`

	// Constraints
	MinItems int `json:"minItems,omitempty"`
	MaxItems int `json:"maxItems,omitempty"`

	// Display
	ItemLabel   string `json:"itemLabel,omitempty"`
	AddText     string `json:"addText,omitempty"`
	RemoveText  string `json:"removeText,omitempty"`
	ShowIndex   bool   `json:"showIndex,omitempty"`
	Sortable    bool   `json:"sortable,omitempty"`
	Collapsible bool   `json:"collapsible,omitempty"`

	// Defaults
	DefaultItem map[string]any `json:"defaultItem,omitempty"`

	// Table mode (for table-style display)
	TableMode *TableModeConfig `json:"tableMode,omitempty"`
}

// TableModeConfig configures table-style display
type TableModeConfig struct {
	Enabled     bool                 `json:"enabled"`
	Editable    bool                 `json:"editable,omitempty"`
	Selectable  bool                 `json:"selectable,omitempty"`
	BulkActions []Action             `json:"bulkActions,omitempty"`
	RowActions  []Action             `json:"rowActions,omitempty"`
	Footer      bool                 `json:"footer,omitempty"`
	Aggregates  map[string]Aggregate `json:"aggregates,omitempty"`
}

// Aggregate defines column aggregation

// AggregateType defines aggregation functions

// Core Methods

// SetEvaluator sets evaluator on repeatable field and all template fields
func (rf *RepeatableField) SetEvaluator(evaluator *condition.Evaluator) {
	rf.Field.SetEvaluator(evaluator)
	for i := range rf.Template {
		rf.Template[i].SetEvaluator(evaluator)
	}
}

// Validate validates repeatable field configuration
func (rf *RepeatableField) Validate(ctx context.Context) error {
	validator := newRepeatableValidator()
	return validator.Validate(ctx, rf)
}

// ValidateItems validates all items in the repeatable field
func (rf *RepeatableField) ValidateItems(ctx context.Context, items []map[string]any) error {
	validator := newRepeatableItemsValidator()
	return validator.ValidateItems(ctx, rf, items)
}

// GetItemLabel generates label for an item
func (rf *RepeatableField) GetItemLabel(item map[string]any, index int) string {
	if rf.ItemLabel == "" {
		return fmt.Sprintf("Item %d", index+1)
	}

	generator := newLabelGenerator()
	return generator.Generate(rf.ItemLabel, item, index)
}

// CreateDefaultItem creates a new item with default values
func (rf *RepeatableField) CreateDefaultItem(ctx context.Context) map[string]any {
	factory := newItemFactory()
	return factory.CreateDefault(ctx, rf)
}

// GetAddButtonText returns text for add button
func (rf *RepeatableField) GetAddButtonText() string {
	if rf.AddText != "" {
		return rf.AddText
	}
	return "Add Item"
}

// GetRemoveButtonText returns text for remove button
func (rf *RepeatableField) GetRemoveButtonText() string {
	if rf.RemoveText != "" {
		return rf.RemoveText
	}
	return "Remove"
}

// CanAddItem checks if more items can be added
func (rf *RepeatableField) CanAddItem(currentCount int) bool {
	if rf.MaxItems <= 0 {
		return true
	}
	return currentCount < rf.MaxItems
}

// CanRemoveItem checks if items can be removed
func (rf *RepeatableField) CanRemoveItem(currentCount int) bool {
	return currentCount > rf.MinItems
}

// repeatableValidator validates repeatable field configuration
type repeatableValidator struct{}

func newRepeatableValidator() *repeatableValidator {
	return &repeatableValidator{}
}

func (v *repeatableValidator) Validate(ctx context.Context, rf *RepeatableField) error {
	collector := NewErrorCollector()

	// Validate base field
	if err := rf.Field.Validate(ctx); err != nil {
		collector.AddError(err.(SchemaError))
	}

	// Validate template
	if len(rf.Template) == 0 {
		collector.AddValidationError(rf.Name, "empty_template",
			"repeatable field must have at least one template field")
	}

	// Validate template fields
	templateFieldNames := make(map[string]bool)
	for i, field := range rf.Template {
		if err := field.Validate(ctx); err != nil {
			collector.AddFieldError(fmt.Sprintf("Template[%d]", i), err.(SchemaError))
		}

		if templateFieldNames[field.Name] {
			collector.AddValidationError(fmt.Sprintf("Template[%d]", i), "duplicate_field",
				fmt.Sprintf("duplicate field name: %s", field.Name))
		}
		templateFieldNames[field.Name] = true
	}

	// Validate constraints
	if rf.MinItems < 0 {
		collector.AddValidationError(rf.Name, "invalid_min_items",
			"minItems cannot be negative")
	}

	if rf.MaxItems > 0 && rf.MinItems > rf.MaxItems {
		collector.AddValidationError(rf.Name, "invalid_item_range",
			"minItems cannot be greater than maxItems")
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

// repeatableItemsValidator validates items in a repeatable field
type repeatableItemsValidator struct{}

func newRepeatableItemsValidator() *repeatableItemsValidator {
	return &repeatableItemsValidator{}
}

func (v *repeatableItemsValidator) ValidateItems(ctx context.Context, rf *RepeatableField, items []map[string]any) error {
	collector := NewErrorCollector()

	// Validate count constraints
	if rf.MinItems > 0 && len(items) < rf.MinItems {
		collector.AddValidationError(rf.Name, "min_items",
			fmt.Sprintf("minimum %d items required, got %d", rf.MinItems, len(items)))
	}

	if rf.MaxItems > 0 && len(items) > rf.MaxItems {
		collector.AddValidationError(rf.Name, "max_items",
			fmt.Sprintf("maximum %d items allowed, got %d", rf.MaxItems, len(items)))
	}

	// Validate each item
	for itemIndex, item := range items {
		if err := v.validateSingleItem(ctx, rf, item, itemIndex); err != nil {
			collector.AddError(err.(SchemaError))
		}
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

func (v *repeatableItemsValidator) validateSingleItem(ctx context.Context, rf *RepeatableField, item map[string]any, itemIndex int) error {
	collector := NewErrorCollector()

	for _, templateField := range rf.Template {
		value, exists := item[templateField.Name]

		// Check required fields
		if templateField.Required && (!exists || isEmpty(value)) {
			fieldPath := fmt.Sprintf("%s[%d].%s", rf.Name, itemIndex, templateField.Name)
			collector.AddValidationError(fieldPath, "required",
				fmt.Sprintf("%s is required", templateField.Label))
			continue
		}

		// Validate field value if exists
		if exists && !isEmpty(value) {
			if err := templateField.ValidateValue(ctx, value); err != nil {
				fieldPath := fmt.Sprintf("%s[%d].%s", rf.Name, itemIndex, templateField.Name)
				if schemaErr, ok := err.(SchemaError); ok {
					collector.AddFieldError(fieldPath, schemaErr)
				} else {
					collector.AddValidationError(fieldPath, "invalid_value", err.Error())
				}
			}
		}
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

// labelGenerator generates item labels
type labelGenerator struct{}

func newLabelGenerator() *labelGenerator {
	return &labelGenerator{}
}

func (g *labelGenerator) Generate(template string, item map[string]any, index int) string {
	label := template

	// Replace {index} with 1-based index
	label = strings.ReplaceAll(label, "{index}", fmt.Sprintf("%d", index+1))
	label = strings.ReplaceAll(label, "{index0}", fmt.Sprintf("%d", index))

	// Replace field references like {field_name}
	for fieldName, value := range item {
		placeholder := fmt.Sprintf("{%s}", fieldName)
		if value != nil {
			label = strings.ReplaceAll(label, placeholder, fmt.Sprintf("%v", value))
		}
	}

	return label
}

// itemFactory creates new items
type itemFactory struct{}

func newItemFactory() *itemFactory {
	return &itemFactory{}
}

func (f *itemFactory) CreateDefault(ctx context.Context, rf *RepeatableField) map[string]any {
	item := make(map[string]any)

	// Apply template field defaults
	for _, field := range rf.Template {
		if field.Default != nil {
			item[field.Name] = field.Default
		} else {
			defaultValue, _ := field.GetDefaultValue(ctx)
			item[field.Name] = defaultValue
		}
	}

	// Apply repeatable field default item
	for key, value := range rf.DefaultItem {
		item[key] = value
	}

	return item
}

// Operations Methods for RepeatableField

// AddItem adds a new item to the collection
func (rf *RepeatableField) AddItem(ctx context.Context, items []map[string]any, item map[string]any) ([]map[string]any, error) {
	// Check max items constraint
	if !rf.CanAddItem(len(items)) {
		return items, NewValidationError("max_items_exceeded",
			fmt.Sprintf("maximum %d items allowed", rf.MaxItems))
	}

	// Validate the new item
	validator := newRepeatableItemsValidator()
	if err := validator.validateSingleItem(ctx, rf, item, len(items)); err != nil {
		return items, err
	}

	return append(items, item), nil
}

// RemoveItem removes an item at the given index
func (rf *RepeatableField) RemoveItem(ctx context.Context, items []map[string]any, index int) ([]map[string]any, error) {
	if index < 0 || index >= len(items) {
		return items, NewValidationError("invalid_index",
			fmt.Sprintf("index %d out of range", index))
	}

	// Check min items constraint
	if !rf.CanRemoveItem(len(items)) {
		return items, NewValidationError("min_items_required",
			fmt.Sprintf("minimum %d items required", rf.MinItems))
	}

	return append(items[:index], items[index+1:]...), nil
}

// UpdateItem updates an item at the given index
func (rf *RepeatableField) UpdateItem(ctx context.Context, items []map[string]any, index int, item map[string]any) ([]map[string]any, error) {
	if index < 0 || index >= len(items) {
		return items, NewValidationError("invalid_index",
			fmt.Sprintf("index %d out of range", index))
	}

	// Validate the updated item
	validator := newRepeatableItemsValidator()
	if err := validator.validateSingleItem(ctx, rf, item, index); err != nil {
		return items, err
	}

	items[index] = item
	return items, nil
}

// MoveItem moves an item from one index to another
func (rf *RepeatableField) MoveItem(items []map[string]any, fromIndex, toIndex int) ([]map[string]any, error) {
	if fromIndex < 0 || fromIndex >= len(items) {
		return items, NewValidationError("invalid_index",
			fmt.Sprintf("from index %d out of range", fromIndex))
	}

	if toIndex < 0 || toIndex >= len(items) {
		return items, NewValidationError("invalid_index",
			fmt.Sprintf("to index %d out of range", toIndex))
	}

	if fromIndex == toIndex {
		return items, nil
	}

	item := items[fromIndex]
	items = append(items[:fromIndex], items[fromIndex+1:]...)

	if toIndex > fromIndex {
		toIndex--
	}

	result := make([]map[string]any, len(items)+1)
	copy(result[:toIndex], items[:toIndex])
	result[toIndex] = item
	copy(result[toIndex+1:], items[toIndex:])

	return result, nil
}

// DuplicateItem duplicates an item at the given index
func (rf *RepeatableField) DuplicateItem(items []map[string]any, index int) ([]map[string]any, error) {
	if index < 0 || index >= len(items) {
		return items, NewValidationError("invalid_index",
			fmt.Sprintf("index %d out of range", index))
	}

	// Check max items constraint
	if !rf.CanAddItem(len(items)) {
		return items, NewValidationError("max_items_exceeded",
			fmt.Sprintf("maximum %d items allowed", rf.MaxItems))
	}

	// Deep copy the item
	newItem := make(map[string]any)
	for key, value := range items[index] {
		newItem[key] = value
	}

	// Insert after current item
	result := make([]map[string]any, len(items)+1)
	copy(result[:index+1], items[:index+1])
	result[index+1] = newItem
	copy(result[index+2:], items[index+1:])

	return result, nil
}

// ClearItems removes all items
func (rf *RepeatableField) ClearItems() ([]map[string]any, error) {
	if rf.MinItems > 0 {
		return nil, NewValidationError("min_items_required",
			fmt.Sprintf("minimum %d items required", rf.MinItems))
	}
	return []map[string]any{}, nil
}

// CalculateAggregates calculates aggregate values
func (rf *RepeatableField) CalculateAggregates(items []map[string]any, aggregates map[string]Aggregate) (map[string]any, error) {
	calculator := newAggregateCalculator()
	return calculator.Calculate(items, aggregates)
}


// aggregateCalculator calculates aggregate values
type aggregateCalculator struct{}

func newAggregateCalculator() *aggregateCalculator {
	return &aggregateCalculator{}
}

func (c *aggregateCalculator) Calculate(items []map[string]any, aggregates map[string]Aggregate) (map[string]any, error) {
	results := make(map[string]any)

	for name, agg := range aggregates {
		switch agg.Type {
		case AggregateSum:
			results[name] = c.calculateSum(items, agg.Field)
		case AggregateAvg:
			results[name] = c.calculateAvg(items, agg.Field)
		case AggregateCount:
			results[name] = len(items)
		case AggregateMin:
			results[name] = c.calculateMin(items, agg.Field)
		case AggregateMax:
			results[name] = c.calculateMax(items, agg.Field)
		default:
			return nil, fmt.Errorf("unknown aggregate type: %s", agg.Type)
		}
	}

	return results, nil
}

func (c *aggregateCalculator) calculateSum(items []map[string]any, field string) float64 {
	var sum float64
	for _, item := range items {
		if val, ok := item[field]; ok {
			if num, ok := toFloat64(val); ok {
				sum += num
			}
		}
	}
	return sum
}

func (c *aggregateCalculator) calculateAvg(items []map[string]any, field string) float64 {
	if len(items) == 0 {
		return 0
	}
	sum := c.calculateSum(items, field)
	return sum / float64(len(items))
}

func (c *aggregateCalculator) calculateMin(items []map[string]any, field string) *float64 {
	var min *float64
	for _, item := range items {
		if val, ok := item[field]; ok {
			if num, ok := toFloat64(val); ok {
				if min == nil || num < *min {
					temp := num
					min = &temp
				}
			}
		}
	}
	return min
}

func (c *aggregateCalculator) calculateMax(items []map[string]any, field string) *float64 {
	var max *float64
	for _, item := range items {
		if val, ok := item[field]; ok {
			if num, ok := toFloat64(val); ok {
				if max == nil || num > *max {
					temp := num
					max = &temp
				}
			}
		}
	}
	return max
}

// RepeatableBuilder provides fluent API for building repeatable fields
type RepeatableBuilder struct {
	field     RepeatableField
	evaluator *condition.Evaluator
}

// NewRepeatableField starts building a repeatable field
func NewRepeatableField(name, label string) *RepeatableBuilder {
	return &RepeatableBuilder{
		field: RepeatableField{
			Field: Field{
				Name:  name,
				Type:  FieldRepeatable,
				Label: label,
			},
			Template:    []Field{},
			MinItems:    0,
			MaxItems:    0,
			ShowIndex:   true,
			Sortable:    true,
			Collapsible: false,
			AddText:     "Add Item",
			RemoveText:  "Remove",
		},
	}
}

func (b *RepeatableBuilder) WithTemplate(fields ...Field) *RepeatableBuilder {
	b.field.Template = append(b.field.Template, fields...)
	return b
}

func (b *RepeatableBuilder) WithMinItems(min int) *RepeatableBuilder {
	b.field.MinItems = min
	return b
}

func (b *RepeatableBuilder) WithMaxItems(max int) *RepeatableBuilder {
	b.field.MaxItems = max
	return b
}

func (b *RepeatableBuilder) WithItemLabel(label string) *RepeatableBuilder {
	b.field.ItemLabel = label
	return b
}

func (b *RepeatableBuilder) WithButtonTexts(addText, removeText string) *RepeatableBuilder {
	b.field.AddText = addText
	b.field.RemoveText = removeText
	return b
}

func (b *RepeatableBuilder) WithDefaultItem(defaults map[string]any) *RepeatableBuilder {
	b.field.DefaultItem = defaults
	return b
}

func (b *RepeatableBuilder) ShowIndex(show bool) *RepeatableBuilder {
	b.field.ShowIndex = show
	return b
}

func (b *RepeatableBuilder) Sortable(sortable bool) *RepeatableBuilder {
	b.field.Sortable = sortable
	return b
}

func (b *RepeatableBuilder) Collapsible(collapsible bool) *RepeatableBuilder {
	b.field.Collapsible = collapsible
	return b
}

func (b *RepeatableBuilder) WithTableMode(config *TableModeConfig) *RepeatableBuilder {
	b.field.TableMode = config
	return b
}

func (b *RepeatableBuilder) WithEvaluator(evaluator *condition.Evaluator) *RepeatableBuilder {
	b.evaluator = evaluator
	return b
}

func (b *RepeatableBuilder) Required() *RepeatableBuilder {
	b.field.Required = true
	return b
}

func (b *RepeatableBuilder) Disabled() *RepeatableBuilder {
	b.field.Disabled = true
	return b
}

func (b *RepeatableBuilder) Hidden() *RepeatableBuilder {
	b.field.Hidden = true
	return b
}

func (b *RepeatableBuilder) WithDescription(desc string) *RepeatableBuilder {
	b.field.Description = desc
	return b
}

func (b *RepeatableBuilder) WithHelp(help string) *RepeatableBuilder {
	b.field.Help = help
	return b
}

func (b *RepeatableBuilder) Build(ctx context.Context) (*RepeatableField, error) {
	if b.evaluator != nil {
		b.field.SetEvaluator(b.evaluator)
	}

	if err := b.field.Validate(ctx); err != nil {
		return nil, err
	}

	return &b.field, nil
}

func (b *RepeatableBuilder) MustBuild(ctx context.Context) *RepeatableField {
	field, err := b.Build(ctx)
	if err != nil {
		panic(fmt.Sprintf("repeatable field build failed: %v", err))
	}
	return field
}

// Preset Repeatable Fields

// NewInvoiceLineItems creates invoice line items repeatable field
func NewInvoiceLineItems() *RepeatableBuilder {
	return NewRepeatableField("line_items", "Line Items").
		WithTemplate(
			Field{
				Name:     "product_id",
				Type:     FieldSelect,
				Label:    "Product",
				Required: true,
			},
			Field{
				Name:  "description",
				Type:  FieldText,
				Label: "Description",
			},
			Field{
				Name:     "quantity",
				Type:     FieldNumber,
				Label:    "Quantity",
				Required: true,
				Default:  1,
				Validation: &FieldValidation{
					Min: floatPtr(0.01),
				},
			},
			Field{
				Name:     "unit_price",
				Type:     FieldCurrency,
				Label:    "Unit Price",
				Required: true,
				Validation: &FieldValidation{
					Min: floatPtr(0),
				},
			},
			Field{
				Name:     "line_total",
				Type:     FieldCurrency,
				Label:    "Total",
				Readonly: true,
			},
		).
		WithMinItems(1).
		WithMaxItems(100).
		WithItemLabel("Line {index}").
		Sortable(true)
}

// NewTaskList creates task list repeatable field
func NewTaskList() *RepeatableBuilder {
	return NewRepeatableField("tasks", "Tasks").
		WithTemplate(
			Field{
				Name:     "title",
				Type:     FieldText,
				Label:    "Task",
				Required: true,
			},
			Field{
				Name:    "completed",
				Type:    FieldCheckbox,
				Label:   "Done",
				Default: false,
			},
			Field{
				Name:  "assigned_to",
				Type:  FieldSelect,
				Label: "Assigned To",
			},
			Field{
				Name:  "due_date",
				Type:  FieldDate,
				Label: "Due Date",
			},
		).
		WithItemLabel("{title}").
		Sortable(true).
		Collapsible(true)
}

// NewAddressList creates address list repeatable field
func NewAddressList() *RepeatableBuilder {
	return NewRepeatableField("addresses", "Addresses").
		WithTemplate(
			Field{
				Name:     "street",
				Type:     FieldText,
				Label:    "Street Address",
				Required: true,
			},
			Field{
				Name:     "city",
				Type:     FieldText,
				Label:    "City",
				Required: true,
			},
			Field{
				Name:  "state",
				Type:  FieldText,
				Label: "State/Province",
			},
			Field{
				Name:     "postal_code",
				Type:     FieldText,
				Label:    "Postal Code",
				Required: true,
			},
			Field{
				Name:     "country",
				Type:     FieldSelect,
				Label:    "Country",
				Required: true,
			},
			Field{
				Name:  "is_primary",
				Type:  FieldCheckbox,
				Label: "Primary Address",
			},
		).
		WithItemLabel("{city}, {country}").
		WithMinItems(1).
		WithMaxItems(5).
		Collapsible(true)
}

// NewContactsList creates contacts list repeatable field
func NewContactsList() *RepeatableBuilder {
	return NewRepeatableField("contacts", "Contacts").
		WithTemplate(
			Field{
				Name:     "name",
				Type:     FieldText,
				Label:    "Name",
				Required: true,
			},
			Field{
				Name:     "email",
				Type:     FieldEmail,
				Label:    "Email",
				Required: true,
			},
			Field{
				Name:  "phone",
				Type:  FieldPhone,
				Label: "Phone",
			},
			Field{
				Name:  "role",
				Type:  FieldText,
				Label: "Role",
			},
		).
		WithItemLabel("{name}").
		WithMinItems(1).
		Sortable(true)
}

// NewExpensesList creates expenses list repeatable field
func NewExpensesList() *RepeatableBuilder {
	return NewRepeatableField("expenses", "Expenses").
		WithTemplate(
			Field{
				Name:     "date",
				Type:     FieldDate,
				Label:    "Date",
				Required: true,
			},
			Field{
				Name:     "category",
				Type:     FieldSelect,
				Label:    "Category",
				Required: true,
			},
			Field{
				Name:     "description",
				Type:     FieldText,
				Label:    "Description",
				Required: true,
			},
			Field{
				Name:     "amount",
				Type:     FieldCurrency,
				Label:    "Amount",
				Required: true,
				Validation: &FieldValidation{
					Min: floatPtr(0),
				},
			},
			Field{
				Name:  "receipt",
				Type:  FieldFile,
				Label: "Receipt",
			},
		).
		WithItemLabel("{date} - {description}").
		WithTableMode(&TableModeConfig{
			Enabled:  true,
			Editable: true,
			Footer:   true,
			Aggregates: map[string]Aggregate{
				"total": {
					Type:  AggregateSum,
					Field: "amount",
					Label: "Total",
				},
			},
		}).
		Sortable(true)
}

// NewPhonesList creates phone numbers list repeatable field
func NewPhonesList() *RepeatableBuilder {
	return NewRepeatableField("phones", "Phone Numbers").
		WithTemplate(
			Field{
				Name:     "type",
				Type:     FieldSelect,
				Label:    "Type",
				Required: true,
				Options: []FieldOption{
					{Value: "mobile", Label: "Mobile"},
					{Value: "home", Label: "Home"},
					{Value: "work", Label: "Work"},
					{Value: "other", Label: "Other"},
				},
			},
			Field{
				Name:     "number",
				Type:     FieldPhone,
				Label:    "Number",
				Required: true,
			},
			Field{
				Name:  "is_primary",
				Type:  FieldCheckbox,
				Label: "Primary",
			},
		).
		WithItemLabel("{type}: {number}").
		WithMinItems(1).
		WithMaxItems(5)
}

// NewDocumentsList creates documents list repeatable field
func NewDocumentsList() *RepeatableBuilder {
	return NewRepeatableField("documents", "Documents").
		WithTemplate(
			Field{
				Name:     "title",
				Type:     FieldText,
				Label:    "Title",
				Required: true,
			},
			Field{
				Name:  "description",
				Type:  FieldTextarea,
				Label: "Description",
			},
			Field{
				Name:     "file",
				Type:     FieldFile,
				Label:    "File",
				Required: true,
			},
			Field{
				Name:  "category",
				Type:  FieldSelect,
				Label: "Category",
			},
			Field{
				Name:  "upload_date",
				Type:  FieldDate,
				Label: "Upload Date",
			},
		).
		WithItemLabel("{title}").
		WithMaxItems(20).
		Sortable(true).
		Collapsible(true)
}

// Utility functions

func floatPtr(f float64) *float64 {
	return &f
}

// RepeatableFieldInfo provides read-only information about repeatable field
type RepeatableFieldInfo interface {
	GetName() string
	GetMinItems() int
	GetMaxItems() int
	GetTemplate() []Field
	CanAddItem(currentCount int) bool
	CanRemoveItem(currentCount int) bool
}

// Implement interface
func (rf *RepeatableField) GetName() string {
	return rf.Name
}

func (rf *RepeatableField) GetMinItems() int {
	return rf.MinItems
}

func (rf *RepeatableField) GetMaxItems() int {
	return rf.MaxItems
}

func (rf *RepeatableField) GetTemplate() []Field {
	fields := make([]Field, len(rf.Template))
	copy(fields, rf.Template)
	return fields
}
