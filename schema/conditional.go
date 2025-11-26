package schema

import (
	"fmt"
	"github.com/niiniyare/ruun/pkg/condition"
)

// Conditional defines visibility and state conditions for any component.
// Used by Field, Action, Layout, LayoutBlock, Section, Tab, Step, etc.
type Conditional struct {
	// Visibility conditions
	Show *condition.ConditionGroup `json:"show,omitempty"`
	Hide *condition.ConditionGroup `json:"hide,omitempty"`

	// State conditions
	Required *condition.ConditionGroup `json:"required,omitempty"`
	Disabled *condition.ConditionGroup `json:"disabled,omitempty"`
	Readonly *condition.ConditionGroup `json:"readonly,omitempty"`

	// Validation conditions
	Validate *condition.ConditionGroup `json:"validate,omitempty"`
}

// IsEmpty returns true if the conditional has no conditions set
func (c *Conditional) IsEmpty() bool {
	if c == nil {
		return true
	}
	return c.Show == nil && c.Hide == nil && c.Required == nil &&
		c.Disabled == nil && c.Readonly == nil && c.Validate == nil
}

// HasVisibility returns true if visibility conditions are set
func (c *Conditional) HasVisibility() bool {
	return c != nil && (c.Show != nil || c.Hide != nil)
}

// HasStateConditions returns true if state conditions are set
func (c *Conditional) HasStateConditions() bool {
	return c != nil && (c.Required != nil || c.Disabled != nil || c.Readonly != nil)
}

// ConditionalBuilder for fluent construction
type ConditionalBuilder struct {
	cond   Conditional
	*BuilderMixin
}

// NewConditional creates a new ConditionalBuilder
func NewConditional() *ConditionalBuilder {
	return &ConditionalBuilder{
		BuilderMixin: NewBuilderMixin(),
	}
}

// ShowWhen sets the show condition
func (b *ConditionalBuilder) ShowWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Show = cg
	return b
}

// HideWhen sets the hide condition
func (b *ConditionalBuilder) HideWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Hide = cg
	return b
}

// RequiredWhen sets the required condition
func (b *ConditionalBuilder) RequiredWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Required = cg
	return b
}

// DisabledWhen sets the disabled condition
func (b *ConditionalBuilder) DisabledWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Disabled = cg
	return b
}

// ReadonlyWhen sets the readonly condition
func (b *ConditionalBuilder) ReadonlyWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Readonly = cg
	return b
}

// ValidateWhen sets the validate condition
func (b *ConditionalBuilder) ValidateWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
	b.cond.Validate = cg
	return b
}

// Build returns the constructed Conditional
func (b *ConditionalBuilder) Build() (*Conditional, error) {
	// Use the unified builder validation
	if err := b.CheckBuild("Conditional"); err != nil {
		return nil, err
	}
	
	// Validate that at least one condition is set
	if b.cond.IsEmpty() {
		b.Context.AddError(fmt.Errorf("conditional must have at least one condition set"))
		return nil, b.Context.CombinedError()
	}
	
	return &b.cond, nil
}

// MustBuild builds and panics on error (for backward compatibility)
func (b *ConditionalBuilder) MustBuild() *Conditional {
	result, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("ConditionalBuilder.MustBuild() failed: %v", err))
	}
	return result
}

// HasErrors implements BaseBuilder interface
func (b *ConditionalBuilder) HasErrors() bool {
	return b.Context.HasErrors()
}

// GetErrors implements BaseBuilder interface  
func (b *ConditionalBuilder) GetErrors() []error {
	return b.Context.GetErrors()
}

// AddError implements BaseBuilder interface
func (b *ConditionalBuilder) AddError(err error) BaseBuilder[Conditional] {
	b.Context.AddError(err)
	return b
}

// ClearErrors implements BaseBuilder interface
func (b *ConditionalBuilder) ClearErrors() BaseBuilder[Conditional] {
	b.Context.ClearErrors()
	return b
}