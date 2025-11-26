// Package schema provides a JSON-driven UI schema system for building enterprise forms and components.
// Backend developers can define complete UIs by writing JSON files - no frontend code required.
//
// Example usage:
//
//	schema := schema.NewSchema("user-form", schema.TypeForm, "Create User")
//	schema.AddField(Field{Name: "email", Type: schema.FieldEmail, Label: "Email"})
//	schema.AddAction(Action{ID: "submit", Type: schema.ActionSubmit, Text: "Save"})
//
// pkg/schema/schema_v2.go
package schema

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/niiniyare/ruun/pkg/condition"
)

// Schema represents an immutable UI schema definition
// Once built, schemas should not be modified - create new versions instead
type Schema struct {
	// Identity
	ID          string `json:"id"`
	Type        Type   `json:"type"`
	Version     string `json:"version"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`

	// Structure
	Fields  []Field  `json:"fields"`
	Actions []Action `json:"actions"`
	Layout  *Layout  `json:"layout,omitempty"`

	// Configuration
	Config     *Config      `json:"config,omitempty"`
	Security   *Security    `json:"security,omitempty"`
	Tenant     *Tenant      `json:"tenant,omitempty"`
	Workflow   *Workflow    `json:"workflow,omitempty"`
	Validation *Validation  `json:"validation,omitempty"`
	Events     *Events      `json:"events,omitempty"`
	I18n       *I18nManager `json:"i18n,omitempty"`

	// Framework Integration
	Behavior *Behavior `json:"htmx,omitempty"`   // Backward compatibility with 'htmx' json tag
	Binding  *Binding  `json:"alpine,omitempty"` // Backward compatibility with 'alpine' json tag

	// Metadata
	Meta     *Meta    `json:"meta,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	Category string   `json:"category,omitempty"`
	Module   string   `json:"module,omitempty"`

	// Internal (optimized lookups)
	fieldIndex  map[string]int       `json:"-"`
	actionIndex map[string]int       `json:"-"`
	evaluator   *condition.Evaluator `json:"-"`
}

// NewSchema creates a new schema with basic properties
func NewSchema(id string, schemaType Type, title string) *Schema {
	return &Schema{
		ID:          id,
		Type:        schemaType,
		Title:       title,
		Version:     "1.0.0",
		Fields:      []Field{},
		Actions:     []Action{},
		evaluator:   &condition.Evaluator{},
		fieldIndex:  make(map[string]int),
		actionIndex: make(map[string]int),
	}
}

// AddField adds a field to the schema
func (s *Schema) AddField(field Field) {
	s.Fields = append(s.Fields, field)
	if s.fieldIndex == nil {
		s.fieldIndex = make(map[string]int)
	}
	s.fieldIndex[field.Name] = len(s.Fields) - 1
}

// Type defines what kind of UI element this schema represents
type Type string

const (
	TypeForm      Type = "form"      // Complete form with fields and submission
	TypeComponent Type = "component" // Reusable UI component
	TypeLayout    Type = "layout"    // Layout container only
	TypeWorkflow  Type = "workflow"  // Multi-step workflow
	TypeTheme     Type = "theme"     // Theme/styling configuration
	TypePage      Type = "page"      // Full page layout
)

// )
// SchemaInfo provides read-only access to schema metadata
type SchemaInfo interface {
	GetID() string
	GetType() Type
	GetVersion() string
	GetTitle() string
	GetFieldCount() int
	GetActionCount() int
}

// SchemaReader provides read-only access to schema data
type SchemaReader interface {
	SchemaInfo

	// Field access
	GetField(name string) (*Field, bool)
	GetFields() []Field
	GetFieldNames() []string
	HasField(name string) bool

	// Action access
	GetAction(id string) (*Action, bool)
	GetActions() []Action
	GetActionIDs() []string
	HasAction(id string) bool

	// Query methods
	GetVisibleFields(ctx context.Context, data map[string]any) []Field
	GetRequiredFields(ctx context.Context, data map[string]any) []Field
	GetFieldsByType(fieldType FieldType) []Field
}

// Implement SchemaInfo interface
func (s *Schema) GetID() string       { return s.ID }
func (s *Schema) GetType() Type       { return s.Type }
func (s *Schema) GetVersion() string  { return s.Version }
func (s *Schema) GetTitle() string    { return s.Title }
func (s *Schema) GetFieldCount() int  { return len(s.Fields) }
func (s *Schema) GetActionCount() int { return len(s.Actions) }

// Implement SchemaReader interface
func (s *Schema) GetField(name string) (*Field, bool) {
	if idx, ok := s.fieldIndex[name]; ok {
		field := s.Fields[idx]
		return &field, true
	}
	return nil, false
}

func (s *Schema) GetFields() []Field {
	// Return copy to prevent external modification
	fields := make([]Field, len(s.Fields))
	copy(fields, s.Fields)
	return fields
}

func (s *Schema) GetFieldNames() []string {
	names := make([]string, len(s.Fields))
	for i, field := range s.Fields {
		names[i] = field.Name
	}
	return names
}

func (s *Schema) HasField(name string) bool {
	_, ok := s.fieldIndex[name]
	return ok
}

func (s *Schema) GetAction(id string) (*Action, bool) {
	if idx, ok := s.actionIndex[id]; ok {
		action := s.Actions[idx]
		return &action, true
	}
	return nil, false
}

func (s *Schema) GetActions() []Action {
	actions := make([]Action, len(s.Actions))
	copy(actions, s.Actions)
	return actions
}

func (s *Schema) GetActionIDs() []string {
	ids := make([]string, len(s.Actions))
	for i, action := range s.Actions {
		ids[i] = action.ID
	}
	return ids
}

func (s *Schema) HasAction(id string) bool {
	_, ok := s.actionIndex[id]
	return ok
}

func (s *Schema) GetVisibleFields(ctx context.Context, data map[string]any) []Field {
	visible := make([]Field, 0, len(s.Fields))
	for _, field := range s.Fields {
		if isVisible, _ := field.IsVisible(ctx, data); isVisible {
			visible = append(visible, field)
		}
	}
	return visible
}

func (s *Schema) GetRequiredFields(ctx context.Context, data map[string]any) []Field {
	required := make([]Field, 0)
	for _, field := range s.Fields {
		if isRequired, _ := field.IsRequired(ctx, data); isRequired {
			if isVisible, _ := field.IsVisible(ctx, data); isVisible {
				required = append(required, field)
			}
		}
	}
	return required
}

func (s *Schema) GetFieldsByType(fieldType FieldType) []Field {
	fields := make([]Field, 0)
	for _, field := range s.Fields {
		if field.Type == fieldType {
			fields = append(fields, field)
		}
	}
	return fields
}

// Validate performs comprehensive schema validation
func (s *Schema) Validate(ctx context.Context) error {
	validator := newSchemaValidator()
	return validator.Validate(ctx, s)
}

// Clone creates a deep copy of the schema
func (s *Schema) Clone() (*Schema, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return nil, fmt.Errorf("clone marshal: %w", err)
	}

	var clone Schema
	if err := json.Unmarshal(data, &clone); err != nil {
		return nil, fmt.Errorf("clone unmarshal: %w", err)
	}

	clone.buildIndexes()
	clone.evaluator = s.evaluator

	return &clone, nil
}

// SetEvaluator sets the condition evaluator
func (s *Schema) SetEvaluator(evaluator *condition.Evaluator) {
	s.evaluator = evaluator
	for i := range s.Fields {
		s.Fields[i].SetEvaluator(evaluator)
	}
}

// GetEvaluator returns the condition evaluator
func (s *Schema) GetEvaluator() *condition.Evaluator {
	return s.evaluator
}

// buildIndexes builds internal lookup indexes
func (s *Schema) buildIndexes() {
	s.fieldIndex = make(map[string]int, len(s.Fields))
	for i, field := range s.Fields {
		s.fieldIndex[field.Name] = i
	}

	s.actionIndex = make(map[string]int, len(s.Actions))
	for i, action := range s.Actions {
		s.actionIndex[action.ID] = i
	}
}

// ToJSON converts schema to JSON
func (s *Schema) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// FromJSON creates schema from JSON
func FromJSON(data []byte) (*Schema, error) {
	var schema Schema
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, err
	}
	schema.buildIndexes()
	return &schema, nil
}

// schemaValidator handles schema validation
type schemaValidator struct{}

func newSchemaValidator() *schemaValidator {
	return &schemaValidator{}
}

func (v *schemaValidator) Validate(ctx context.Context, s *Schema) error {
	collector := NewErrorCollector()

	// Basic validation
	if s.ID == "" {
		collector.AddError(ErrInvalidSchemaID)
	}
	if s.Type == "" {
		collector.AddError(ErrInvalidSchemaType)
	}
	if s.Title == "" {
		collector.AddError(ErrInvalidSchemaTitle)
	}

	// Validate fields
	fieldNames := make(map[string]bool)
	for i, field := range s.Fields {
		if err := field.Validate(ctx); err != nil {
			collector.AddFieldError(fmt.Sprintf("Fields[%d]", i), err.(SchemaError))
		}

		if fieldNames[field.Name] {
			collector.AddError(ErrFieldDuplicate.WithField(field.Name))
		}
		fieldNames[field.Name] = true
	}

	// Validate actions
	actionIDs := make(map[string]bool)
	for i, action := range s.Actions {
		if err := action.Validate(ctx); err != nil {
			collector.AddFieldError(fmt.Sprintf("Actions[%d]", i), err.(SchemaError))
		}

		if actionIDs[action.ID] {
			collector.AddValidationError(action.ID, "duplicate_action",
				fmt.Sprintf("duplicate action ID: %s", action.ID))
		}
		actionIDs[action.ID] = true
	}

	// Validate dependencies
	for _, field := range s.Fields {
		for _, dep := range field.Dependencies {
			if !fieldNames[dep] {
				collector.AddValidationError(field.Name, "invalid_dependency",
					fmt.Sprintf("depends on non-existent field: %s", dep))
			}
		}
	}

	// Check circular dependencies
	if err := v.checkCircularDependencies(s); err != nil {
		collector.AddError(err.(SchemaError))
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

func (v *schemaValidator) checkCircularDependencies(s *Schema) error {
	visited := make(map[string]bool)
	stack := make(map[string]bool)

	for _, field := range s.Fields {
		if err := v.checkCycle(s, field.Name, visited, stack); err != nil {
			return err
		}
	}
	return nil
}

func (v *schemaValidator) checkCycle(s *Schema, fieldName string, visited, stack map[string]bool) error {
	if stack[fieldName] {
		return NewValidationError("circular_dependency",
			"circular dependency detected: "+fieldName).WithField(fieldName)
	}
	if visited[fieldName] {
		return nil
	}

	visited[fieldName] = true
	stack[fieldName] = true

	if field, ok := s.GetField(fieldName); ok {
		for _, dep := range field.Dependencies {
			if err := v.checkCycle(s, dep, visited, stack); err != nil {
				return err
			}
		}
	}

	stack[fieldName] = false
	return nil
}

// ApplySchemaI18nLocalization applies internationalization to schema elements
func (s *Schema) ApplySchemaI18nLocalization(schemaI18n *SchemaI18n, locale string) (*Schema, error) {
	if schemaI18n == nil {
		return s, nil
	}
	
	// Create a copy of the schema to modify
	localized := *s
	
	// Apply localized title
	if title, ok := schemaI18n.Title[locale]; ok && title != "" {
		localized.Title = title
	}
	
	// Apply localized description
	if desc, ok := schemaI18n.Description[locale]; ok && desc != "" {
		localized.Description = desc
	}
	
	return &localized, nil
}

// IsRTL checks if the schema should be rendered in right-to-left mode
func (s *Schema) IsRTL(locale string) bool {
	if s.I18n != nil && s.I18n.currentLocale != "" {
		// Check if the current locale is RTL
		rtlLanguages := map[string]bool{
			"ar": true, // Arabic
			"he": true, // Hebrew
			"fa": true, // Persian/Farsi
			"ur": true, // Urdu
		}
		return rtlLanguages[s.I18n.currentLocale] || rtlLanguages[locale]
	}
	return false
}

// buildFieldMap builds a map of field names to fields for quick lookup
func (s *Schema) buildFieldMap() map[string]*Field {
	fieldMap := make(map[string]*Field)
	for i := range s.Fields {
		fieldMap[s.Fields[i].Name] = &s.Fields[i]
	}
	return fieldMap
}
