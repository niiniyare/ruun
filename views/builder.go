package views

import (
	"context"
	"fmt"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/pkg/schema/runtime"
)

// RuntimeBuilder provides a clean, fluent interface for building runtime instances
// with sensible defaults and easy customization options.
type RuntimeBuilder struct {
	schema       *schema.Schema
	registry     *ComponentRegistry
	renderer     *TemplateRenderer
	stateManager *StateManager
	config       *RuntimeConfig
	initialData  map[string]any
	locale       string
}

// RuntimeConfig holds view-specific configuration options
type RuntimeConfig struct {
	EnableValidation    bool
	EnableConditionals  bool
	EnableEventTracking bool
	DefaultLocale       string
	ThemeVariant        string
	ComponentPrefix     string
}

// NewRuntimeBuilder creates a new builder with sensible defaults
func NewRuntimeBuilder(s *schema.Schema) *RuntimeBuilder {
	return &RuntimeBuilder{
		schema:      s,
		registry:    NewComponentRegistry(),
		config:      defaultRuntimeConfig(),
		initialData: make(map[string]any),
		locale:      "en",
	}
}

// WithDefaultComponents registers all built-in Templ components
func (b *RuntimeBuilder) WithDefaultComponents() *RuntimeBuilder {
	b.registry.RegisterDefaults()
	return b
}

// WithCustomRegistry replaces the component registry with a custom one
func (b *RuntimeBuilder) WithCustomRegistry(registry *ComponentRegistry) *RuntimeBuilder {
	b.registry = registry
	return b
}

// WithConfig sets custom configuration options
func (b *RuntimeBuilder) WithConfig(config *RuntimeConfig) *RuntimeBuilder {
	b.config = config
	return b
}

// WithInitialData sets initial form values
func (b *RuntimeBuilder) WithInitialData(data map[string]any) *RuntimeBuilder {
	for k, v := range data {
		b.initialData[k] = v
	}
	return b
}

// WithLocale sets the default locale for rendering
func (b *RuntimeBuilder) WithLocale(locale string) *RuntimeBuilder {
	b.locale = locale
	return b
}

// Build creates a ready-to-use runtime instance
func (b *RuntimeBuilder) Build() (*Runtime, error) {
	// Initialize components if not done
	if !b.registry.hasDefaults {
		b.registry.RegisterDefaults()
	}

	// Create field mapper with schema context
	mapper := NewFieldMapper(b.schema, b.locale)

	// Create template renderer with registry and mapper
	b.renderer = &TemplateRenderer{
		registry: b.registry,
		mapper:   mapper,
		schema:   b.schema,
	}

	// Create state manager
	b.stateManager = NewStateManager(b.schema, b.initialData)

	// Create schema runtime using existing runtime builder
	runtimeBuilder := runtime.NewRuntimeBuilder(b.schema).
		WithRenderer(b.renderer).
		WithLocale(b.locale)

	// Add initial data if provided
	if len(b.initialData) > 0 {
		runtimeBuilder = runtimeBuilder.WithInitialData(b.initialData)
	}

	// Build the schema runtime
	schemaRuntime, err := runtimeBuilder.Build(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create schema runtime: %w", err)
	}

	// Return our views runtime wrapper
	return &Runtime{
		schema:        b.schema,
		schemaRuntime: schemaRuntime,
		renderer:      b.renderer,
		stateManager:  b.stateManager,
		config:        b.config,
	}, nil
}

// Runtime wraps the schema runtime with views-specific functionality
type Runtime struct {
	schema        *schema.Schema
	schemaRuntime *runtime.Runtime
	renderer      *TemplateRenderer
	stateManager  *StateManager
	config        *RuntimeConfig
}

// RenderForm renders the complete form using the schema
func (r *Runtime) RenderForm(ctx context.Context) (string, error) {
	// Get current state from state manager
	state := r.stateManager.GetAllValues()
	errors := r.stateManager.GetAllErrors()

	// Use the schema runtime to render
	return r.renderer.RenderForm(ctx, r.schema, state, errors)
}

// RenderField renders a single field with current state
func (r *Runtime) RenderField(ctx context.Context, fieldName string) (string, error) {
	field, exists := r.schema.GetField(fieldName)
	if !exists {
		return "", fmt.Errorf("field %s not found", fieldName)
	}

	value := r.stateManager.GetValue(fieldName)
	errors := r.stateManager.GetFieldErrors(fieldName)
	touched := r.stateManager.IsFieldTouched(fieldName)
	dirty := r.stateManager.IsFieldDirty(fieldName)

	return r.renderer.RenderField(ctx, &field, value, errors, touched, dirty)
}

// SetFieldValue updates a field value and triggers validation
func (r *Runtime) SetFieldValue(fieldName string, value any) error {
	return r.stateManager.SetValue(fieldName, value)
}

// GetState returns the current form state
func (r *Runtime) GetState() map[string]any {
	return r.stateManager.GetAllValues()
}

// GetErrors returns current validation errors
func (r *Runtime) GetErrors() map[string][]string {
	return r.stateManager.GetAllErrors()
}

// Validate runs validation on all fields
func (r *Runtime) Validate(ctx context.Context) error {
	return r.stateManager.ValidateAll(ctx)
}

// defaultRuntimeConfig returns sensible default configuration
func defaultRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		EnableValidation:    true,
		EnableConditionals:  true,
		EnableEventTracking: false,
		DefaultLocale:       "en",
		ThemeVariant:        "default",
		ComponentPrefix:     "form",
	}
}
