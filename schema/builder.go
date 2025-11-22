// pkg/schema/builder_v2.go
package schema

import (
	"context"
	"fmt"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
)

// SchemaBuilder provides fluent API for building schemas
type SchemaBuilder struct {
	schema    *Schema
	evaluator *condition.Evaluator
	mixins    *MixinRegistry
	errors    []error
}

// NewSchemaBuilder creates a new schema builder
func NewSchemaBuilder(id string, schemaType Type, title string) *SchemaBuilder {
	now := time.Now()
	return &SchemaBuilder{
		schema: &Schema{
			ID:          id,
			Type:        schemaType,
			Version:     "1.0.0",
			Title:       title,
			Fields:      []Field{},
			Actions:     []Action{},
			fieldIndex:  make(map[string]int),
			actionIndex: make(map[string]int),
			Meta: &Meta{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		mixins: NewMixinRegistry(),
		errors: []error{},
	}
}

// Core Configuration Methods

func (b *SchemaBuilder) WithDescription(desc string) *SchemaBuilder {
	b.schema.Description = desc
	return b
}

func (b *SchemaBuilder) WithVersion(version string) *SchemaBuilder {
	b.schema.Version = version
	return b
}

func (b *SchemaBuilder) WithCategory(category string) *SchemaBuilder {
	b.schema.Category = category
	return b
}

func (b *SchemaBuilder) WithModule(module string) *SchemaBuilder {
	b.schema.Module = module
	return b
}

func (b *SchemaBuilder) WithTags(tags ...string) *SchemaBuilder {
	b.schema.Tags = append(b.schema.Tags, tags...)
	return b
}

// Evaluator Configuration

func (b *SchemaBuilder) WithEvaluator(evaluator *condition.Evaluator) *SchemaBuilder {
	b.evaluator = evaluator
	return b
}

// Field Management

func (b *SchemaBuilder) AddField(field Field) *SchemaBuilder {
	if b.evaluator != nil {
		field.SetEvaluator(b.evaluator)
	}

	// Check for duplicates
	if _, exists := b.schema.fieldIndex[field.Name]; exists {
		b.errors = append(b.errors, fmt.Errorf("duplicate field: %s", field.Name))
		return b
	}

	b.schema.Fields = append(b.schema.Fields, field)
	b.schema.fieldIndex[field.Name] = len(b.schema.Fields) - 1
	return b
}

func (b *SchemaBuilder) AddFields(fields ...Field) *SchemaBuilder {
	for _, field := range fields {
		b.AddField(field)
	}
	return b
}

// Action Management

func (b *SchemaBuilder) AddAction(action Action) *SchemaBuilder {
	// Check for duplicates
	if _, exists := b.schema.actionIndex[action.ID]; exists {
		b.errors = append(b.errors, fmt.Errorf("duplicate action: %s", action.ID))
		return b
	}

	b.schema.Actions = append(b.schema.Actions, action)
	b.schema.actionIndex[action.ID] = len(b.schema.Actions) - 1
	return b
}

func (b *SchemaBuilder) AddActions(actions ...Action) *SchemaBuilder {
	for _, action := range actions {
		b.AddAction(action)
	}
	return b
}

// Layout Configuration

func (b *SchemaBuilder) WithLayout(layout *Layout) *SchemaBuilder {
	if b.evaluator != nil {
		layout.SetEvaluator(b.evaluator)
	}
	b.schema.Layout = layout
	return b
}

// Configuration Objects

func (b *SchemaBuilder) WithConfig(config *Config) *SchemaBuilder {
	b.schema.Config = config
	return b
}

func (b *SchemaBuilder) WithSecurity(security *Security) *SchemaBuilder {
	b.schema.Security = security
	return b
}

func (b *SchemaBuilder) WithTenant(field, isolation string) *SchemaBuilder {
	b.schema.Tenant = &Tenant{
		Enabled:   true,
		Field:     field,
		Isolation: isolation,
	}
	return b
}

func (b *SchemaBuilder) WithWorkflow(workflow *Workflow) *SchemaBuilder {
	b.schema.Workflow = workflow
	return b
}

func (b *SchemaBuilder) WithValidation(validation *Validation) *SchemaBuilder {
	b.schema.Validation = validation
	return b
}

func (b *SchemaBuilder) WithEvents(events *Events) *SchemaBuilder {
	b.schema.Events = events
	return b
}

func (b *SchemaBuilder) WithI18n(defaultLocale string, supported ...string) *SchemaBuilder {
	b.schema.I18n = &I18n{
		Enabled:          true,
		DefaultLocale:    defaultLocale,
		SupportedLocales: supported,
	}
	return b
}

// Framework Integration

func (b *SchemaBuilder) WithHTMX(post, target string) *SchemaBuilder {
	b.schema.HTMX = &HTMX{
		Enabled: true,
		Post:    post,
		Target:  target,
		Swap:    "innerHTML",
	}
	return b
}

func (b *SchemaBuilder) WithAlpine(xData string) *SchemaBuilder {
	b.schema.Alpine = &Alpine{
		Enabled: true,
		XData:   xData,
	}
	return b
}

// Mixin Support

func (b *SchemaBuilder) WithMixin(mixinID string) *SchemaBuilder {
	if err := b.mixins.ApplyMixinWithEvaluator(b.schema, mixinID, "", b.evaluator); err != nil {
		b.errors = append(b.errors, fmt.Errorf("mixin %s: %w", mixinID, err))
	}
	return b
}

func (b *SchemaBuilder) WithCustomMixin(mixin *Mixin) *SchemaBuilder {
	if err := b.mixins.Register(mixin); err != nil {
		b.errors = append(b.errors, err)
		return b
	}
	return b.WithMixin(mixin.ID)
}

// Build Methods

// Build validates and returns the schema
func (b *SchemaBuilder) Build(ctx context.Context) (*Schema, error) {
	// Check for build errors
	if len(b.errors) > 0 {
		return nil, fmt.Errorf("build errors: %v", b.errors)
	}

	// Set evaluator on schema
	if b.evaluator != nil {
		b.schema.SetEvaluator(b.evaluator)
	}

	// Rebuild indexes
	b.schema.buildIndexes()

	// Validate
	if err := b.schema.Validate(ctx); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// Update metadata
	b.schema.Meta.UpdatedAt = time.Now()

	return b.schema, nil
}

// MustBuild builds or panics
func (b *SchemaBuilder) MustBuild(ctx context.Context) *Schema {
	schema, err := b.Build(ctx)
	if err != nil {
		panic(fmt.Sprintf("schema build failed: %v", err))
	}
	return schema
}

// BuildWithoutValidation builds without validation (use with caution)
func (b *SchemaBuilder) BuildWithoutValidation() *Schema {
	if b.evaluator != nil {
		b.schema.SetEvaluator(b.evaluator)
	}
	b.schema.buildIndexes()
	b.schema.Meta.UpdatedAt = time.Now()
	return b.schema
}

// Helper Methods for Common Patterns

// AddTextField adds a text field
func (b *SchemaBuilder) AddTextField(name, label string, required bool) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldText,
		Label:    label,
		Required: required,
	})
}

// AddEmailField adds an email field
func (b *SchemaBuilder) AddEmailField(name, label string, required bool) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldEmail,
		Label:    label,
		Required: required,
		Validation: &FieldValidation{
			Format: "email",
		},
	})
}

// AddPasswordField adds a password field
func (b *SchemaBuilder) AddPasswordField(name, label string, required bool, minLength int) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldPassword,
		Label:    label,
		Required: required,
		Validation: &FieldValidation{
			MinLength: &minLength,
		},
	})
}

// AddNumberField adds a number field
func (b *SchemaBuilder) AddNumberField(name, label string, required bool, min, max *float64) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldNumber,
		Label:    label,
		Required: required,
		Validation: &FieldValidation{
			Min: min,
			Max: max,
		},
	})
}

// AddSelectField adds a select field
func (b *SchemaBuilder) AddSelectField(name, label string, required bool, options []FieldOption) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldSelect,
		Label:    label,
		Required: required,
		Options:  options,
	})
}

// AddTextareaField adds a textarea field
func (b *SchemaBuilder) AddTextareaField(name, label string, required bool, rows int) *SchemaBuilder {
	field := Field{
		Name:     name,
		Type:     FieldTextarea,
		Label:    label,
		Required: required,
	}
	if rows > 0 {
		field.Config = map[string]any{"rows": rows}
	}
	return b.AddField(field)
}

// AddCheckboxField adds a checkbox field
func (b *SchemaBuilder) AddCheckboxField(name, label string, defaultValue bool) *SchemaBuilder {
	return b.AddField(Field{
		Name:    name,
		Type:    FieldCheckbox,
		Label:   label,
		Default: defaultValue,
	})
}

// AddDateField adds a date field
func (b *SchemaBuilder) AddDateField(name, label string, required bool) *SchemaBuilder {
	return b.AddField(Field{
		Name:     name,
		Type:     FieldDate,
		Label:    label,
		Required: required,
	})
}

// AddSubmitButton adds a submit button
func (b *SchemaBuilder) AddSubmitButton(text string) *SchemaBuilder {
	return b.AddAction(Action{
		ID:      "submit",
		Type:    ActionSubmit,
		Text:    text,
		Variant: "primary",
		Size:    "md",
	})
}

// AddResetButton adds a reset button
func (b *SchemaBuilder) AddResetButton(text string) *SchemaBuilder {
	return b.AddAction(Action{
		ID:      "reset",
		Type:    ActionReset,
		Text:    text,
		Variant: "secondary",
		Size:    "md",
	})
}

// AddCancelButton adds a cancel button
func (b *SchemaBuilder) AddCancelButton(text string) *SchemaBuilder {
	return b.AddAction(Action{
		ID:      "cancel",
		Type:    ActionButton,
		Text:    text,
		Variant: "outline",
		Size:    "md",
	})
}

// Validation Helpers

// HasErrors checks if builder has any errors
func (b *SchemaBuilder) HasErrors() bool {
	return len(b.errors) > 0
}

// GetErrors returns all builder errors
func (b *SchemaBuilder) GetErrors() []error {
	return b.errors
}

// ClearErrors clears all builder errors
func (b *SchemaBuilder) ClearErrors() *SchemaBuilder {
	b.errors = []error{}
	return b
}
