package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// ComponentRenderer defines the interface for rendering schema fields
type ComponentRenderer interface {
	Render(ctx context.Context, props molecules.FormFieldProps) (string, error)
	SupportsType(fieldType schema.FieldType) bool
}

// ComponentConfig holds configuration for a component renderer
type ComponentConfig struct {
	Renderer    ComponentRenderer
	Priority    int  // Higher priority takes precedence
	Fallback    bool // Can be used as fallback for unsupported types
	Description string
}

// ComponentRegistry manages the mapping between schema field types and Templ components
type ComponentRegistry struct {
	mappings    map[schema.FieldType]ComponentConfig
	fallbacks   []ComponentConfig
	hasDefaults bool
}

// NewComponentRegistry creates a new component registry
func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		mappings:  make(map[schema.FieldType]ComponentConfig),
		fallbacks: make([]ComponentConfig, 0),
	}
}

// Register adds a component renderer for a specific field type
func (r *ComponentRegistry) Register(fieldType schema.FieldType, renderer ComponentRenderer) error {
	return r.RegisterWithConfig(fieldType, ComponentConfig{
		Renderer:    renderer,
		Priority:    100, // Default priority
		Description: fmt.Sprintf("Renderer for %s fields", fieldType),
	})
}

// RegisterWithConfig adds a component renderer with custom configuration
func (r *ComponentRegistry) RegisterWithConfig(fieldType schema.FieldType, config ComponentConfig) error {
	if config.Renderer == nil {
		return fmt.Errorf("renderer cannot be nil")
	}

	// Check if existing registration has lower priority
	if existing, exists := r.mappings[fieldType]; exists {
		if existing.Priority >= config.Priority {
			return fmt.Errorf("field type %s already registered with higher priority", fieldType)
		}
	}

	r.mappings[fieldType] = config

	// Add to fallbacks if marked as fallback
	if config.Fallback {
		r.fallbacks = append(r.fallbacks, config)
	}

	return nil
}

// RegisterDefaults registers all built-in Templ component mappings
func (r *ComponentRegistry) RegisterDefaults() {
	if r.hasDefaults {
		return
	}

	// Create the FormField renderer (our main renderer) - not used directly in mapping

	// Basic text inputs - all use FormField with different types
	basicMappings := map[schema.FieldType]molecules.FormFieldType{
		schema.FieldText:     molecules.FormFieldText,
		schema.FieldEmail:    molecules.FormFieldEmail,
		schema.FieldPassword: molecules.FormFieldPassword,
		schema.FieldNumber:   molecules.FormFieldNumber,
		schema.FieldPhone:    molecules.FormFieldTel,
		schema.FieldURL:      molecules.FormFieldURL,
		schema.FieldHidden:   molecules.FormFieldText, // Hidden handled via props
	}

	for schemaType, formFieldType := range basicMappings {
		r.mappings[schemaType] = ComponentConfig{
			Renderer:    &TypedFormFieldRenderer{FormFieldType: formFieldType},
			Priority:    100,
			Description: fmt.Sprintf("Basic input for %s", schemaType),
		}
	}

	// Date and time fields
	dateMappings := map[schema.FieldType]molecules.FormFieldType{
		schema.FieldDate:      molecules.FormFieldDate,
		schema.FieldTime:      molecules.FormFieldTime,
		schema.FieldDateTime:  molecules.FormFieldDateTime,
		schema.FieldDateRange: molecules.FormFieldDateRange,
	}

	for schemaType, formFieldType := range dateMappings {
		r.mappings[schemaType] = ComponentConfig{
			Renderer:    &TypedFormFieldRenderer{FormFieldType: formFieldType},
			Priority:    100,
			Description: fmt.Sprintf("Date/time input for %s", schemaType),
		}
	}

	// Selection fields
	selectionMappings := map[schema.FieldType]molecules.FormFieldType{
		schema.FieldSelect:      molecules.FormFieldSelect,
		schema.FieldMultiSelect: molecules.FormFieldMultiSelect,
		schema.FieldRadio:       molecules.FormFieldRadio,
		schema.FieldCheckbox:    molecules.FormFieldCheckbox,
		schema.FieldCheckboxes:  molecules.FormFieldCheckboxGroup,
	}

	for schemaType, formFieldType := range selectionMappings {
		r.mappings[schemaType] = ComponentConfig{
			Renderer:    &TypedFormFieldRenderer{FormFieldType: formFieldType},
			Priority:    100,
			Description: fmt.Sprintf("Selection input for %s", schemaType),
		}
	}

	// Text content fields
	textMappings := map[schema.FieldType]molecules.FormFieldType{
		schema.FieldTextarea: molecules.FormFieldTextarea,
	}

	for schemaType, formFieldType := range textMappings {
		r.mappings[schemaType] = ComponentConfig{
			Renderer:    &TypedFormFieldRenderer{FormFieldType: formFieldType},
			Priority:    100,
			Description: fmt.Sprintf("Text content input for %s", schemaType),
		}
	}

	// Specialized fields
	specializedMappings := map[schema.FieldType]molecules.FormFieldType{
		schema.FieldAutoComplete: molecules.FormFieldAutoComplete,
		schema.FieldColor:        molecules.FormFieldColor,
		schema.FieldFile:         molecules.FormFieldFile,
		schema.FieldSlider:       molecules.FormFieldRange,
		schema.FieldTags:         molecules.FormFieldTags,
	}

	for schemaType, formFieldType := range specializedMappings {
		r.mappings[schemaType] = ComponentConfig{
			Renderer:    &TypedFormFieldRenderer{FormFieldType: formFieldType},
			Priority:    100,
			Description: fmt.Sprintf("Specialized input for %s", schemaType),
		}
	}

	// Register fallback renderer for any unmapped types
	r.RegisterWithConfig(schema.FieldText, ComponentConfig{
		Renderer:    &TypedFormFieldRenderer{FormFieldType: molecules.FormFieldText},
		Priority:    1, // Low priority fallback
		Fallback:    true,
		Description: "Fallback text renderer for unsupported field types",
	})

	r.hasDefaults = true
}

// Resolve finds the appropriate component renderer for a field
func (r *ComponentRegistry) Resolve(field *schema.Field) (ComponentRenderer, error) {
	if field == nil {
		return nil, fmt.Errorf("field cannot be nil")
	}

	// Try direct mapping first
	if config, exists := r.mappings[field.Type]; exists {
		return config.Renderer, nil
	}

	// Try fallback renderers
	for _, fallback := range r.fallbacks {
		if fallback.Renderer.SupportsType(field.Type) {
			return fallback.Renderer, nil
		}
	}

	return nil, fmt.Errorf("no renderer found for field type: %s", field.Type)
}

// GetSupportedTypes returns all field types supported by the registry
func (r *ComponentRegistry) GetSupportedTypes() []schema.FieldType {
	types := make([]schema.FieldType, 0, len(r.mappings))
	for fieldType := range r.mappings {
		types = append(types, fieldType)
	}
	return types
}

// GetRendererInfo returns information about registered renderers
func (r *ComponentRegistry) GetRendererInfo() map[schema.FieldType]string {
	info := make(map[schema.FieldType]string)
	for fieldType, config := range r.mappings {
		info[fieldType] = config.Description
	}
	return info
}

// FormFieldRenderer is the default renderer that uses the FormField molecule
type FormFieldRenderer struct{}

// Render renders a field using the FormField molecule
func (r *FormFieldRenderer) Render(ctx context.Context, props molecules.FormFieldProps) (string, error) {
	// Use templ to render the FormField component
	// Note: In a real implementation, you'd use templ's rendering system
	// For now, we return a placeholder
	return renderFormField(props), nil
}

// SupportsType checks if this renderer supports the field type
func (r *FormFieldRenderer) SupportsType(fieldType schema.FieldType) bool {
	// FormFieldRenderer is our universal renderer
	return true
}

// TypedFormFieldRenderer renders FormField with a specific type
type TypedFormFieldRenderer struct {
	FormFieldType molecules.FormFieldType
}

// Render renders using FormField with the specified type
func (r *TypedFormFieldRenderer) Render(ctx context.Context, props molecules.FormFieldProps) (string, error) {
	// Override the type with our specific type
	props.Type = r.FormFieldType
	return renderFormField(props), nil
}

// SupportsType always returns true since this is a typed renderer
func (r *TypedFormFieldRenderer) SupportsType(fieldType schema.FieldType) bool {
	return true
}

// renderFormField uses actual Templ rendering to render the FormField molecule
func renderFormField(props molecules.FormFieldProps) string {
	// Use templ to render the FormField component
	var buf strings.Builder
	err := molecules.FormField(props).Render(context.Background(), &buf)
	if err != nil {
		// Fallback to error display in case of rendering failure
		return fmt.Sprintf(`<div class="form-field-error" data-error="render-failed">Failed to render field: %s</div>`, err.Error())
	}
	return buf.String()
}

// getHTMLInputType maps FormFieldType to HTML input type (deprecated)
// This function is kept for backward compatibility but should not be used
// now that we're using real Templ components that handle type mapping internally
func getHTMLInputType(fieldType molecules.FormFieldType) string {
	mappings := map[molecules.FormFieldType]string{
		molecules.FormFieldText:     "text",
		molecules.FormFieldEmail:    "email",
		molecules.FormFieldPassword: "password",
		molecules.FormFieldNumber:   "number",
		molecules.FormFieldTel:      "tel",
		molecules.FormFieldURL:      "url",
		molecules.FormFieldDate:     "date",
		molecules.FormFieldTime:     "time",
		molecules.FormFieldDateTime: "datetime-local",
	}

	if htmlType, exists := mappings[fieldType]; exists {
		return htmlType
	}
	return "text" // Default fallback
}

// RegisterComponent is a convenience function for external packages
var DefaultRegistry = NewComponentRegistry()

// RegisterComponent registers a component with the default registry
func RegisterComponent(fieldType schema.FieldType, renderer ComponentRenderer) error {
	return DefaultRegistry.Register(fieldType, renderer)
}

// GetDefaultRegistry returns the default component registry
func GetDefaultRegistry() *ComponentRegistry {
	if !DefaultRegistry.hasDefaults {
		DefaultRegistry.RegisterDefaults()
	}
	return DefaultRegistry
}
