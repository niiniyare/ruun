// pkg/schema/runtime_v2.go
package schema

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Renderer converts schema to HTML/templates
type Renderer interface {
	Render(ctx context.Context, schema *Schema, data map[string]any) (string, error)
	RenderField(ctx context.Context, field *Field, value any) (string, error)
}

// Runtime executes schema with state management
type Runtime struct {
	schema   *Schema
	state    *State
	handlers *EventHandlers

	// Pluggable components
	validator   Validator
	renderer    Renderer
	conditional ConditionalEngine

	// Configuration
	config *RuntimeConfig

	// Metadata
	createdAt time.Time
	updatedAt time.Time

	mu sync.RWMutex
}

// RuntimeConfig configures runtime behavior
type RuntimeConfig struct {
	ValidateOnChange bool
	ValidateOnBlur   bool
	ValidateOnSubmit bool
	EnableDebounce   bool
	DebounceDelay    time.Duration
	EnableTracking   bool
}

// DefaultRuntimeConfig returns default configuration
func DefaultRuntimeConfig() *RuntimeConfig {
	return &RuntimeConfig{
		ValidateOnChange: false,
		ValidateOnBlur:   true,
		ValidateOnSubmit: true,
		EnableDebounce:   true,
		DebounceDelay:    300 * time.Millisecond,
		EnableTracking:   false,
	}
}

// State manages form state
type State struct {
	values  map[string]any
	errors  map[string][]string
	touched map[string]bool
	dirty   map[string]bool
	initial map[string]any

	mu sync.RWMutex
}

// NewState creates a new state
func NewState() *State {
	return &State{
		values:  make(map[string]any),
		errors:  make(map[string][]string),
		touched: make(map[string]bool),
		dirty:   make(map[string]bool),
		initial: make(map[string]any),
	}
}

// GetValue gets a field value
func (s *State) GetValue(name string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.values[name]
	return val, ok
}

// SetValue sets a field value
func (s *State) SetValue(name string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[name] = value

	// Mark as dirty if different from initial
	if initial, ok := s.initial[name]; !ok || initial != value {
		s.dirty[name] = true
	}
}

// GetAllValues returns all values
func (s *State) GetAllValues() map[string]any {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make(map[string]any, len(s.values))
	for k, v := range s.values {
		values[k] = v
	}
	return values
}

// GetErrors gets field errors
func (s *State) GetErrors(name string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.errors[name]
}

// SetErrors sets field errors
func (s *State) SetErrors(name string, errors []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(errors) > 0 {
		s.errors[name] = errors
	} else {
		delete(s.errors, name)
	}
}

// GetAllErrors returns all errors
func (s *State) GetAllErrors() map[string][]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	errors := make(map[string][]string, len(s.errors))
	for k, v := range s.errors {
		errors[k] = append([]string{}, v...)
	}
	return errors
}

// IsTouched checks if field was touched
func (s *State) IsTouched(name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.touched[name]
}

// SetTouched marks field as touched
func (s *State) SetTouched(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.touched[name] = true
}

// IsDirty checks if field is dirty
func (s *State) IsDirty(name string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.dirty[name]
}

// IsValid checks if form is valid
func (s *State) IsValid() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.errors) == 0
}

// Reset resets state to initial
func (s *State) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.values = make(map[string]any)
	for k, v := range s.initial {
		s.values[k] = v
	}

	s.errors = make(map[string][]string)
	s.touched = make(map[string]bool)
	s.dirty = make(map[string]bool)
}

// EventHandlers manages event callbacks
type EventHandlers struct {
	onChange []func(ctx context.Context, field string, value any) error
	onBlur   []func(ctx context.Context, field string) error
	onFocus  []func(ctx context.Context, field string) error
	onSubmit []func(ctx context.Context) error
	onReset  []func(ctx context.Context) error

	mu sync.RWMutex
}

// NewEventHandlers creates event handlers
func NewEventHandlers() *EventHandlers {
	return &EventHandlers{}
}

// OnChange registers change handler
func (h *EventHandlers) OnChange(handler func(ctx context.Context, field string, value any) error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.onChange = append(h.onChange, handler)
}

// TriggerChange triggers change handlers
func (h *EventHandlers) TriggerChange(ctx context.Context, field string, value any) error {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, handler := range h.onChange {
		if err := handler(ctx, field, value); err != nil {
			return err
		}
	}
	return nil
}

// RuntimeBuilder builds runtime instances
type RuntimeBuilder struct {
	schema      *Schema
	config      *RuntimeConfig
	validator   Validator
	renderer    Renderer
	conditional ConditionalEngine
	initialData map[string]any
}

// NewRuntime starts building a runtime
func NewRuntime(schema *Schema) *RuntimeBuilder {
	return &RuntimeBuilder{
		schema:      schema,
		config:      DefaultRuntimeConfig(),
		initialData: make(map[string]any),
	}
}

func (b *RuntimeBuilder) WithConfig(config *RuntimeConfig) *RuntimeBuilder {
	b.config = config
	return b
}

func (b *RuntimeBuilder) WithValidator(validator Validator) *RuntimeBuilder {
	b.validator = validator
	return b
}

func (b *RuntimeBuilder) WithRenderer(renderer Renderer) *RuntimeBuilder {
	b.renderer = renderer
	return b
}

func (b *RuntimeBuilder) WithConditionalEngine(engine ConditionalEngine) *RuntimeBuilder {
	b.conditional = engine
	return b
}

func (b *RuntimeBuilder) WithInitialData(data map[string]any) *RuntimeBuilder {
	for k, v := range data {
		b.initialData[k] = v
	}
	return b
}

func (b *RuntimeBuilder) Build() (*Runtime, error) {
	if b.schema == nil {
		return nil, fmt.Errorf("schema is required")
	}

	state := NewState()
	for k, v := range b.initialData {
		state.values[k] = v
		state.initial[k] = v
	}

	return &Runtime{
		schema:      b.schema,
		state:       state,
		handlers:    NewEventHandlers(),
		validator:   b.validator,
		renderer:    b.renderer,
		conditional: b.conditional,
		config:      b.config,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (b *RuntimeBuilder) MustBuild() *Runtime {
	runtime, err := b.Build()
	if err != nil {
		panic(err)
	}
	return runtime
}

// Runtime Methods

// HandleChange handles field change
func (r *Runtime) HandleChange(ctx context.Context, field string, value any) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Update state
	r.state.SetValue(field, value)
	r.updatedAt = time.Now()

	// Trigger handlers
	if err := r.handlers.TriggerChange(ctx, field, value); err != nil {
		return err
	}

	// Validate if enabled
	if r.config.ValidateOnChange && r.validator != nil {
		if err := r.validateField(ctx, field); err != nil {
			return err
		}
	}

	// Update conditionals
	if r.conditional != nil {
		if err := r.applyConditionals(ctx); err != nil {
			return err
		}
	}

	return nil
}

// HandleBlur handles field blur
func (r *Runtime) HandleBlur(ctx context.Context, field string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.state.SetTouched(field)
	r.updatedAt = time.Now()

	if r.config.ValidateOnBlur && r.validator != nil {
		return r.validateField(ctx, field)
	}

	return nil
}

// HandleSubmit handles form submission
func (r *Runtime) HandleSubmit(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.config.ValidateOnSubmit && r.validator != nil {
		if err := r.validateAll(ctx); err != nil {
			return err
		}
	}

	if !r.state.IsValid() {
		return fmt.Errorf("form has validation errors")
	}

	return nil
}

// GetState returns current state
func (r *Runtime) GetState() *State {
	return r.state
}

// GetSchema returns schema
func (r *Runtime) GetSchema() *Schema {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.schema
}

// validateField validates a single field
func (r *Runtime) validateField(ctx context.Context, fieldName string) error {
	field, ok := r.schema.GetField(fieldName)
	if !ok {
		return fmt.Errorf("field not found: %s", fieldName)
	}

	value, _ := r.state.GetValue(fieldName)

	if err := r.validator.ValidateField(ctx, field, value); err != nil {
		r.state.SetErrors(fieldName, []string{err.Error()})
		return err
	}

	return nil
}

// validateAll validates all fields
func (r *Runtime) validateAll(ctx context.Context) error {
	if err := r.validator.ValidateSchema(ctx, r.schema); err != nil {
		return err
	}

	data := r.state.GetAllValues()
	return r.validator.ValidateData(ctx, r.schema, data)
}

// applyConditionals applies conditional logic
func (r *Runtime) applyConditionals(ctx context.Context) error {
	// Implementation depends on ConditionalEngine interface
	return nil
}

// ConditionalEngine interface for conditional logic
type ConditionalEngine interface {
	Evaluate(ctx context.Context, schema *Schema, data map[string]any) error
}
