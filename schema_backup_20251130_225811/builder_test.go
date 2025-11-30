package schema

import (
	"context"
	"testing"

	"github.com/niiniyare/ruun/pkg/condition"
	"github.com/stretchr/testify/suite"
)

// =============================================================================
// Test Suite Definition
// =============================================================================

// BuilderTestSuite provides comprehensive testing for the schema builder,
// covering all builder methods, Foundation features, and integration scenarios.
type BuilderTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *BuilderTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(BuilderTestSuite))
}

// =============================================================================
// Test Fixtures & Helpers
// =============================================================================

// testFixtures provides reusable test data
type testFixtures struct{}

func (f testFixtures) selectOptions() []FieldOption {
	return []FieldOption{
		{Value: "option1", Label: "Option 1"},
		{Value: "option2", Label: "Option 2"},
		{Value: "option3", Label: "Option 3"},
	}
}

func (f testFixtures) lineItemTemplate() []Field {
	return []Field{
		{Name: "product", Type: FieldText, Label: "Product", Required: true},
		{Name: "quantity", Type: FieldNumber, Label: "Quantity", Required: true},
		{Name: "unit_price", Type: FieldNumber, Label: "Unit Price", Required: true},
		{Name: "line_total", Type: FieldNumber, Label: "Line Total", Readonly: true},
	}
}

func (f testFixtures) visibilityCondition(field, value string) *condition.ConditionGroup {
	return &condition.ConditionGroup{
		ID:          "visibility_condition",
		Conjunction: condition.ConjunctionAnd,
		Children: []any{
			&condition.ConditionRule{
				ID: "field_check",
				Left: condition.Expression{
					Type:  condition.ValueTypeField,
					Field: field,
				},
				Op:    condition.OpEqual,
				Right: value,
			},
		},
	}
}

var fixtures = testFixtures{}

// assertFieldExists verifies a field exists with expected properties
func (s *BuilderTestSuite) assertFieldExists(schema *Schema, name string, fieldType FieldType) Field {
	field, exists := schema.GetField(name)
	s.Require().True(exists, "Field %s should exist", name)
	s.Require().Equal(fieldType, field.Type, "Field %s should have type %s", name, fieldType)
	return *field
}

// assertFieldCount verifies the schema has expected number of fields
func (s *BuilderTestSuite) assertFieldCount(schema *Schema, count int) {
	s.Require().Len(schema.Fields, count, "Schema should have %d fields", count)
}

// assertActionExists verifies an action exists with expected properties
func (s *BuilderTestSuite) assertActionExists(schema *Schema, id string, actionType ActionType) Action {
	for _, action := range schema.Actions {
		if action.ID == id {
			s.Require().Equal(actionType, action.Type, "Action %s should have type %s", id, actionType)
			return action
		}
	}
	s.Require().Fail("Action %s should exist", id)
	return Action{}
}

// =============================================================================
// 1. Core Builder Creation & Configuration
// =============================================================================

func (s *BuilderTestSuite) TestCoreBuilder() {
	s.Run("NewSchemaBuilder initializes correctly", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")

		s.Require().NotNil(builder)
		s.Require().NotNil(builder.schema)
		s.Require().NotNil(builder.mixinSupport)
		s.Require().NotNil(builder.ruleEngine)
		s.Require().Equal("test-schema", builder.schema.ID)
		s.Require().Equal(TypeForm, builder.schema.Type)
		s.Require().Equal("Test Schema", builder.schema.Title)
	})

	s.Run("Build creates valid schema", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("name", "Name", true)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema)
		s.Require().Equal("test-schema", schema.ID)
	})

	s.Run("MustBuild returns schema on success", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("name", "Name", true)

		schema := builder.MustBuild(s.ctx)

		s.Require().NotNil(schema)
		s.Require().Equal("test-schema", schema.ID)
	})

	s.Run("Build fails with invalid schema", func() {
		builder := NewSchemaBuilder("", TypeForm, "")

		_, err := builder.Build(s.ctx)

		s.Require().Error(err)
	})
}

func (s *BuilderTestSuite) TestBuilderMetadata() {
	testCases := []struct {
		name     string
		setup    func(*SchemaBuilder)
		validate func(*Schema)
	}{
		{
			name: "WithVersion",
			setup: func(b *SchemaBuilder) {
				b.WithVersion("2.1.0")
			},
			validate: func(schema *Schema) {
				s.Require().Equal("2.1.0", schema.Version)
			},
		},
		{
			name: "WithDescription",
			setup: func(b *SchemaBuilder) {
				b.WithDescription("Test form description")
			},
			validate: func(schema *Schema) {
				s.Require().Equal("Test form description", schema.Description)
			},
		},
		{
			name: "WithCategory",
			setup: func(b *SchemaBuilder) {
				b.WithCategory("finance")
			},
			validate: func(schema *Schema) {
				s.Require().Equal("finance", schema.Category)
			},
		},
		{
			name: "WithModule",
			setup: func(b *SchemaBuilder) {
				b.WithModule("accounting")
			},
			validate: func(schema *Schema) {
				s.Require().Equal("accounting", schema.Module)
			},
		},
		{
			name: "WithTags",
			setup: func(b *SchemaBuilder) {
				b.WithTags("form", "test", "example")
			},
			validate: func(schema *Schema) {
				s.Require().Contains(schema.Tags, "form")
				s.Require().Contains(schema.Tags, "test")
				s.Require().Contains(schema.Tags, "example")
			},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
			tc.setup(builder)

			schema, err := builder.Build(s.ctx)

			s.Require().NoError(err)
			tc.validate(schema)
		})
	}
}

// =============================================================================
// 2. Field Builder Methods
// =============================================================================

func (s *BuilderTestSuite) TestFieldBuilderMethods() {
	testCases := []struct {
		name      string
		setup     func(*SchemaBuilder)
		fieldName string
		fieldType FieldType
		validate  func(Field)
	}{
		{
			name: "AddTextField",
			setup: func(b *SchemaBuilder) {
				b.AddTextField("username", "Username", true)
			},
			fieldName: "username",
			fieldType: FieldText,
			validate: func(f Field) {
				s.Require().True(f.Required)
				s.Require().Equal("Username", f.Label)
			},
		},
		{
			name: "AddEmailField",
			setup: func(b *SchemaBuilder) {
				b.AddEmailField("email", "Email Address", true)
			},
			fieldName: "email",
			fieldType: FieldEmail,
			validate: func(f Field) {
				s.Require().True(f.Required)
			},
		},
		{
			name: "AddPasswordField",
			setup: func(b *SchemaBuilder) {
				b.AddPasswordField("password", "Password", true, 8)
			},
			fieldName: "password",
			fieldType: FieldPassword,
			validate: func(f Field) {
				s.Require().True(f.Required)
			},
		},
		{
			name: "AddNumberField",
			setup: func(b *SchemaBuilder) {
				min, max := float64(0), float64(100)
				b.AddNumberField("age", "Age", false, &min, &max)
			},
			fieldName: "age",
			fieldType: FieldNumber,
			validate: func(f Field) {
				s.Require().False(f.Required)
			},
		},
		{
			name: "AddSelectField",
			setup: func(b *SchemaBuilder) {
				b.AddSelectField("category", "Category", false, fixtures.selectOptions())
			},
			fieldName: "category",
			fieldType: FieldSelect,
			validate: func(f Field) {
				s.Require().Len(f.Options, 3)
			},
		},
		{
			name: "AddTextareaField",
			setup: func(b *SchemaBuilder) {
				b.AddTextareaField("description", "Description", false, 5)
			},
			fieldName: "description",
			fieldType: FieldTextarea,
			validate:  func(f Field) {},
		},
		{
			name: "AddDateField",
			setup: func(b *SchemaBuilder) {
				b.AddDateField("birthdate", "Birth Date", false)
			},
			fieldName: "birthdate",
			fieldType: FieldDate,
			validate:  func(f Field) {},
		},
		{
			name: "AddCheckboxField",
			setup: func(b *SchemaBuilder) {
				b.AddCheckboxField("terms", "Accept Terms", true)
			},
			fieldName: "terms",
			fieldType: FieldCheckbox,
			validate:  func(f Field) {},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
			tc.setup(builder)

			schema, err := builder.Build(s.ctx)

			s.Require().NoError(err)
			field := s.assertFieldExists(schema, tc.fieldName, tc.fieldType)
			tc.validate(field)
		})
	}
}

func (s *BuilderTestSuite) TestFieldBuilderPattern() {
	s.Run("NewField with full configuration", func() {
		field := NewField("test_field", FieldText).
			WithLabel("Test Label").
			WithDescription("Test description").
			WithPlaceholder("Enter text here").
			WithHelp("This is help text").
			Required().
			WithDefault("default value").
			Build()

		s.Require().Equal("test_field", field.Name)
		s.Require().Equal(FieldText, field.Type)
		s.Require().Equal("Test Label", field.Label)
		s.Require().Equal("Test description", field.Description)
		s.Require().Equal("Enter text here", field.Placeholder)
		s.Require().Equal("This is help text", field.Help)
		s.Require().True(field.Required)
		s.Require().Equal("default value", field.Default)
	})

	s.Run("NewField with Disabled and Readonly", func() {
		field := NewField("test_field", FieldText).
			Disabled().
			Readonly().
			Build()

		s.Require().True(field.Disabled)
		s.Require().True(field.Readonly)
	})

	s.Run("NewField with Options", func() {
		field := NewField("test_field", FieldSelect).
			WithOptions(fixtures.selectOptions()).
			Build()

		s.Require().Len(field.Options, 3)
	})

	s.Run("NewField with Validation", func() {
		minLen, maxLen := 5, 100
		validation := &FieldValidation{
			MinLength: &minLen,
			MaxLength: &maxLen,
		}

		field := NewField("test_field", FieldText).
			WithValidation(validation).
			Build()

		s.Require().NotNil(field.Validation)
		s.Require().Equal(5, *field.Validation.MinLength)
		s.Require().Equal(100, *field.Validation.MaxLength)
	})
}

func (s *BuilderTestSuite) TestConditionalFields() {
	s.Run("Field with visibility condition", func() {
		builder := NewSchemaBuilder("test-conditional", TypeForm, "Test Conditional")
		builder.AddSelectField("account_type", "Account Type", true, []FieldOption{
			{Value: "personal", Label: "Personal"},
			{Value: "business", Label: "Business"},
		})

		businessField := Field{
			Name:     "company_name",
			Type:     FieldText,
			Label:    "Company Name",
			Required: true,
			Conditional: &Conditional{
				Show: fixtures.visibilityCondition("account_type", "business"),
			},
		}
		builder.AddField(businessField)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		field, exists := schema.GetField("company_name")
		s.Require().True(exists)
		s.Require().NotNil(field.Conditional)
		s.Require().NotNil(field.Conditional.Show)
	})
}

// =============================================================================
// 3. Action Builder Methods
// =============================================================================

func (s *BuilderTestSuite) TestActionBuilderMethods() {
	s.Run("AddSubmitButton", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddSubmitButton("Submit Form")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.assertActionExists(schema, "submit", ActionSubmit)
		s.Require().Equal("Submit Form", schema.Actions[0].Text)
	})

	s.Run("AddResetButton", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddResetButton("Reset Form")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.assertActionExists(schema, "reset", ActionReset)
	})

	s.Run("AddButton", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddButton("custom-btn", "Custom Button", "secondary")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.assertActionExists(schema, "custom-btn", ActionButton)
	})

	s.Run("AddActionWithConfig", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		customAction := Action{
			ID:      "advanced-action",
			Type:    ActionButton,
			Text:    "Advanced Action",
			Variant: "outline",
			Size:    "lg",
		}
		builder.AddActionWithConfig(customAction)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		action := s.assertActionExists(schema, "advanced-action", ActionButton)
		s.Require().Equal("outline", action.Variant)
		s.Require().Equal("lg", action.Size)
	})
}

// =============================================================================
// 4. Layout & Configuration
// =============================================================================

func (s *BuilderTestSuite) TestLayoutConfiguration() {
	s.Run("WithLayout Grid", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		layout := &Layout{
			Type:    LayoutGrid,
			Columns: 2,
			Gap:     "md",
		}
		builder.WithLayout(layout)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Layout)
		s.Require().Equal(LayoutGrid, schema.Layout.Type)
		s.Require().Equal(2, schema.Layout.Columns)
	})

	s.Run("WithConfig", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		config := &Config{
			Method:   "POST",
			Action:   "/api/users",
			Encoding: "application/json",
		}
		builder.WithConfig(config)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Config)
		s.Require().Equal("POST", schema.Config.Method)
		s.Require().Equal("/api/users", schema.Config.Action)
	})

	s.Run("WithI18n", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithI18n("en", "en", "es", "fr", "de")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.I18n)
		s.Require().True(schema.I18n.Enabled)
		s.Require().Equal("en", schema.I18n.DefaultLocale)
		s.Require().Contains(schema.I18n.SupportedLocales, "es")
		s.Require().Contains(schema.I18n.SupportedLocales, "fr")
	})
}

func (s *BuilderTestSuite) TestLayoutHelperFunctions() {
	s.Run("NewGridLayout", func() {
		layout := NewGridLayout(3)

		s.Require().Equal(LayoutGrid, layout.Type)
		s.Require().Equal(3, layout.Columns)
		s.Require().Equal("1rem", layout.Gap)
	})

	s.Run("NewTabLayout", func() {
		tabs := []Tab{
			{ID: "tab1", Label: "Tab 1", Fields: []string{"field1"}},
			{ID: "tab2", Label: "Tab 2", Fields: []string{"field2"}},
		}
		layout := NewTabLayout(tabs)

		s.Require().Equal(LayoutTabs, layout.Type)
		s.Require().Len(layout.Tabs, 2)
	})

	s.Run("NewStepLayout", func() {
		steps := []Step{
			{ID: "step1", Title: "Step 1", Order: 1, Fields: []string{"field1"}},
			{ID: "step2", Title: "Step 2", Order: 2, Fields: []string{"field2"}},
		}
		layout := NewStepLayout(steps)

		s.Require().Equal(LayoutSteps, layout.Type)
		s.Require().Len(layout.Steps, 2)
	})

	s.Run("NewSimpleConfig", func() {
		config := NewSimpleConfig("/api/submit", "POST")

		s.Require().Equal("POST", config.Method)
		s.Require().Equal("/api/submit", config.Action)
		s.Require().Equal("application/json", config.Encoding)
	})
}

// =============================================================================
// 5. Security Configuration
// =============================================================================

func (s *BuilderTestSuite) TestSecurityConfiguration() {
	s.Run("WithCSRF", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithCSRF()

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Security)
		s.Require().NotNil(schema.Security.CSRF)
		s.Require().True(schema.Security.CSRF.Enabled)
	})

	s.Run("WithSecurity full config", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		security := &Security{
			CSRF: &CSRF{
				Enabled:    true,
				FieldName:  "_csrf_token",
				HeaderName: "X-CSRF-Token",
			},
		}
		builder.WithSecurity(security)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Security)
		s.Require().NotNil(schema.Security.CSRF)
		s.Require().Equal("_csrf_token", schema.Security.CSRF.FieldName)
	})

	s.Run("WithRateLimit", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithRateLimit(100, 60)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Security)
		s.Require().NotNil(schema.Security.RateLimit)
		s.Require().True(schema.Security.RateLimit.Enabled)
		s.Require().Equal(100, schema.Security.RateLimit.MaxRequests)
	})

	s.Run("WithTenant", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithTenant("organization_id", "strict")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Tenant)
		s.Require().True(schema.Tenant.Enabled)
		s.Require().Equal("organization_id", schema.Tenant.Field)
		s.Require().Equal("strict", schema.Tenant.Isolation)
	})
}

// =============================================================================
// 6. Behavior Configuration (HTMX, Alpine, Binding)
// =============================================================================

func (s *BuilderTestSuite) TestBehaviorConfiguration() {
	s.Run("WithHTMX", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithHTMX("/api/submit", "#result")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Behavior)
	})

	s.Run("WithBinding", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		binding := &Binding{
			Data: "{ count: 0, increment() { this.count++ } }",
			Show: "count > 0",
			On:   map[string]string{"click": "increment()"},
		}
		builder.WithBinding(binding)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Binding)
		s.Require().Equal("{ count: 0, increment() { this.count++ } }", schema.Binding.Data)
		s.Require().True(schema.Binding.HasVisibility())
	})

	s.Run("WithAlpine deprecated compatibility", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithAlpine("{ count: 0, increment() { this.count++ } }")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema.Binding)
		s.Require().Equal("{ count: 0, increment() { this.count++ } }", schema.Binding.Data)
	})
}

// =============================================================================
// 7. Mixin Support
// =============================================================================

func (s *BuilderTestSuite) TestMixinSupport() {
	s.Run("Built-in audit mixin", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithMixin("audit_fields")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)

		// Verify audit fields exist
		fieldNames := make(map[string]bool)
		for _, field := range schema.Fields {
			fieldNames[field.Name] = true
		}
		s.Require().True(fieldNames["created_at"], "Should have created_at")
		s.Require().True(fieldNames["updated_at"], "Should have updated_at")
		s.Require().True(fieldNames["created_by"], "Should have created_by")
		s.Require().True(fieldNames["updated_by"], "Should have updated_by")

		// Check metadata
		s.Require().NotNil(schema.Meta)
		appliedMixins, exists := schema.Meta.CustomData["applied_mixins"]
		s.Require().True(exists)
		s.Require().Contains(appliedMixins, "audit_fields")
	})

	s.Run("Custom mixin", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		customMixin := &Mixin{
			ID:          "custom_test",
			Name:        "Custom Test Mixin",
			Description: "A test mixin",
			Fields: []Field{
				{Name: "custom_field", Type: FieldText, Label: "Custom Field"},
			},
			Actions: []Action{
				{ID: "custom_action", Type: ActionButton, Text: "Custom Action"},
			},
		}
		builder.WithCustomMixin(customMixin)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.assertFieldExists(schema, "custom_field", FieldText)
		s.assertActionExists(schema, "custom_action", ActionButton)
	})

	s.Run("Non-existent mixin gracefully handled", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithMixin("non_existent_mixin")

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema)
	})

	s.Run("GetMixinRegistry accessor", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")

		registry := builder.GetMixinRegistry()

		s.Require().NotNil(registry)
		s.Require().IsType(&MixinRegistry{}, registry)
	})
}

// =============================================================================
// 8. Repeatable Field Support
// =============================================================================

func (s *BuilderTestSuite) TestRepeatableFieldSupport() {
	s.Run("WithRepeatable for line items", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		repeatableField := &RepeatableField{
			Field: Field{
				Name:  "line_items",
				Type:  FieldRepeatable,
				Label: "Line Items",
			},
			Template: fixtures.lineItemTemplate(),
			MinItems: 1,
			MaxItems: 50,
		}
		builder.WithRepeatable(repeatableField)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.assertFieldExists(schema, "line_items", FieldRepeatable)
	})
}

// =============================================================================
// 9. Validation Support
// =============================================================================

func (s *BuilderTestSuite) TestValidationSupport() {
	s.Run("WithCustomValidator", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithCustomValidator("custom_test", func(ctx context.Context, value any, params map[string]any) error {
			if value == "invalid" {
				return NewValidationError("custom_error", "custom validation failed")
			}
			return nil
		})

		validator, exists := builder.GetValidator("custom_test")

		s.Require().True(exists, "custom validator should exist")
		s.Require().NotNil(validator)
	})

	s.Run("GetValidator for non-existent validator", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")

		validator, exists := builder.GetValidator("non_existent")

		s.Require().False(exists)
		s.Require().Nil(validator)
	})

	s.Run("Schema validation on build", func() {
		builder := NewSchemaBuilder("test-validation", TypeForm, "Test Validation")
		builder.AddTextField("username", "Username", true)
		builder.AddEmailField("email", "Email", true)

		schema, err := builder.Build(s.ctx)
		s.Require().NoError(err)

		err = schema.Validate(s.ctx)
		s.Require().NoError(err)
	})
}

// =============================================================================
// 10. Business Rules Support
// =============================================================================

func (s *BuilderTestSuite) TestBusinessRulesSupport() {
	s.Run("WithBusinessRule", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("status", "Status", true)
		builder.AddTextField("conditional_field", "Conditional Field", false)

		rule := &BusinessRule{
			ID:      "visibility_rule",
			Name:    "Visibility Rule",
			Type:    RuleTypeFieldVisibility,
			Enabled: true,
			Actions: []BusinessRuleAction{
				{Type: ActionShowField, Target: "conditional_field"},
			},
		}
		builder.WithBusinessRule(rule)

		ruleEngine := builder.GetBusinessRuleEngine()
		s.Require().NotNil(ruleEngine)

		retrievedRule, exists := ruleEngine.GetRule("visibility_rule")
		s.Require().True(exists)
		s.Require().Equal("Visibility Rule", retrievedRule.Name)
	})

	s.Run("WithBusinessRuleBuilder", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("amount", "Amount", true)
		builder.AddTextField("tax", "Tax", false)

		ruleBuilder := NewBusinessRule("calculation_rule", "Tax Calculation", RuleTypeDataCalculation).
			WithDescription("Calculate tax based on amount").
			WithAction(ActionCalculate, "tax", nil)
		builder.WithBusinessRuleBuilder(ruleBuilder)

		ruleEngine := builder.GetBusinessRuleEngine()
		rule, exists := ruleEngine.GetRule("calculation_rule")
		s.Require().True(exists)
		s.Require().Equal("Tax Calculation", rule.Name)
		s.Require().Equal("Calculate tax based on amount", rule.Description)
	})

	s.Run("ApplyBusinessRules", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("field1", "Field 1", false)
		builder.AddTextField("field2", "Field 2", false)

		rule := &BusinessRule{
			ID:      "hide_rule",
			Name:    "Hide Rule",
			Type:    RuleTypeFieldVisibility,
			Enabled: true,
			Actions: []BusinessRuleAction{
				{Type: ActionHideField, Target: "field2"},
			},
		}
		builder.WithBusinessRule(rule)

		modifiedSchema, err := builder.ApplyBusinessRules(s.ctx, map[string]any{})

		s.Require().NoError(err)
		for _, field := range modifiedSchema.Fields {
			if field.Name == "field2" {
				s.Require().True(field.Hidden, "Field2 should be hidden")
			}
		}
	})

	s.Run("BuildWithRules conditional visibility", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.AddTextField("name", "Name", true)
		builder.AddTextField("optional_field", "Optional Field", false)

		ruleCondition := fixtures.visibilityCondition("name", "admin")
		rule := &BusinessRule{
			ID:        "admin_rule",
			Name:      "Admin Rule",
			Type:      RuleTypeFieldVisibility,
			Enabled:   true,
			Condition: ruleCondition,
			Actions: []BusinessRuleAction{
				{Type: ActionShowField, Target: "optional_field"},
			},
		}
		builder.WithBusinessRule(rule)

		// Build for admin user
		adminSchema, err := builder.BuildWithRules(s.ctx, map[string]any{"name": "admin"})
		s.Require().NoError(err)
		for _, field := range adminSchema.Fields {
			if field.Name == "optional_field" {
				s.Require().False(field.Hidden, "Optional field should be visible for admin")
			}
		}

		// Build for regular user
		userSchema, err := builder.BuildWithRules(s.ctx, map[string]any{"name": "user"})
		s.Require().NoError(err)
		s.Require().NotNil(userSchema)
	})

	s.Run("Invalid business rule not added", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		invalidRuleBuilder := NewBusinessRule("", "Invalid Rule", RuleTypeFieldVisibility)
		builder.WithBusinessRuleBuilder(invalidRuleBuilder)

		ruleEngine := builder.GetBusinessRuleEngine()
		rules := ruleEngine.ListRules()
		s.Require().Len(rules, 0, "Invalid rule should not be added")
	})

	s.Run("GetBusinessRuleEngine accessor", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")

		ruleEngine := builder.GetBusinessRuleEngine()

		s.Require().NotNil(ruleEngine)
		s.Require().IsType(&BusinessRuleEngine{}, ruleEngine)
	})
}

// =============================================================================
// 11. Helper Functions
// =============================================================================

func (s *BuilderTestSuite) TestOptionHelperFunctions() {
	s.Run("CreateOption", func() {
		option := CreateOption("value1", "Label 1")

		s.Require().Equal("value1", option.Value)
		s.Require().Equal("Label 1", option.Label)
	})

	s.Run("CreateOptionWithIcon", func() {
		option := CreateOptionWithIcon("value2", "Label 2", "icon-name")

		s.Require().Equal("value2", option.Value)
		s.Require().Equal("Label 2", option.Label)
		s.Require().Equal("icon-name", option.Icon)
	})

	s.Run("CreateGroupedOption", func() {
		option := CreateGroupedOption("group1", "Group 1", "finance")

		s.Require().Equal("group1", option.Value)
		s.Require().Equal("Group 1", option.Label)
		s.Require().Equal("finance", option.Group)
	})
}

// =============================================================================
// 12. Integration Tests
// =============================================================================

func (s *BuilderTestSuite) TestFullInvoiceFormIntegration() {
	s.Run("Complete invoice form with all Foundation features", func() {
		builder := NewSchemaBuilder("invoice-form", TypeForm, "Invoice Form")

		// Apply audit mixin
		builder.WithMixin("audit_fields")

		// Add invoice fields
		builder.AddTextField("invoice_number", "Invoice Number", true).
			AddTextField("customer_name", "Customer Name", true).
			AddNumberField("total", "Total Amount", true, nil, nil)

		// Add line items
		lineItemField := &RepeatableField{
			Field: Field{
				Name:  "line_items",
				Type:  FieldRepeatable,
				Label: "Line Items",
			},
			Template: fixtures.lineItemTemplate(),
			MinItems: 1,
			MaxItems: 100,
		}
		builder.WithRepeatable(lineItemField)

		// Add custom validator
		builder.WithCustomValidator("invoice_format", func(ctx context.Context, value any, params map[string]any) error {
			str, ok := value.(string)
			if !ok {
				return NewValidationError("invalid_type", "invoice number must be string")
			}
			if len(str) < 3 || str[:3] != "INV" {
				return NewValidationError("invalid_format", "invoice number must start with INV")
			}
			return nil
		})

		// Add business rule
		calcRule, buildErr := NewBusinessRule("total_calculation", "Total Calculation", RuleTypeDataCalculation).
			WithDescription("Calculate total from line items").
			WithAction(ActionCalculate, "total", nil).
			Build()
		s.Require().NoError(buildErr)
		builder.WithBusinessRule(calcRule)

		// Add security
		builder.WithCSRF()
		builder.WithRateLimit(10, 60)

		// Add submit button
		builder.AddSubmitButton("Create Invoice")

		// Build the complete schema
		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().Equal("invoice-form", schema.ID)
		s.Require().Equal("Invoice Form", schema.Title)

		// Verify field count (audit fields + custom fields + repeatable)
		s.Require().GreaterOrEqual(len(schema.Fields), 7)

		// Verify security
		s.Require().NotNil(schema.Security)
		s.Require().NotNil(schema.Security.CSRF)
		s.Require().True(schema.Security.CSRF.Enabled)
		s.Require().NotNil(schema.Security.RateLimit)
		s.Require().True(schema.Security.RateLimit.Enabled)

		// Verify actions
		s.Require().Len(schema.Actions, 1)
		s.Require().Equal("submit", schema.Actions[0].ID)

		// Verify business rule
		ruleEngine := builder.GetBusinessRuleEngine()
		rule, exists := ruleEngine.GetRule("total_calculation")
		s.Require().True(exists)
		s.Require().Equal("Total Calculation", rule.Name)
	})
}

func (s *BuilderTestSuite) TestFluentInterfaceChaining() {
	s.Run("All Foundation methods support chaining", func() {
		schema, err := NewSchemaBuilder("fluent-test", TypeForm, "Fluent Test").
			WithDescription("Testing fluent interface").
			WithVersion("1.0.0").
			WithTags("test", "fluent").
			WithMixin("audit_fields").
			WithCustomValidator("test", func(ctx context.Context, value any, params map[string]any) error {
				return nil
			}).
			WithBusinessRule(&BusinessRule{
				ID:      "test_rule",
				Name:    "Test Rule",
				Type:    RuleTypeFieldVisibility,
				Enabled: true,
				Actions: []BusinessRuleAction{
					{Type: ActionShowField, Target: "test_field"},
				},
			}).
			AddTextField("test_field", "Test Field", false).
			AddSubmitButton("Submit").
			Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema)
		s.Require().Equal("fluent-test", schema.ID)
		s.Require().Equal("Testing fluent interface", schema.Description)
		s.Require().Equal("1.0.0", schema.Version)
		s.Require().Len(schema.Fields, 5) // audit(4) + test(1)
		s.Require().Len(schema.Actions, 1)
	})
}

// =============================================================================
// 13. Error Handling & Edge Cases
// =============================================================================

func (s *BuilderTestSuite) TestErrorHandling() {
	s.Run("Invalid business rule fails at build time", func() {
		invalidRule, err := NewBusinessRule("", "", RuleTypeFieldVisibility).Build()

		s.Require().Error(err)
		s.Require().Nil(invalidRule)
	})

	s.Run("Builder handles nil evaluator gracefully", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		builder.WithEvaluator(nil)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema)
	})
}

// =============================================================================
// 14. Additional Builder Methods
// =============================================================================

func (s *BuilderTestSuite) TestAdditionalBuilderMethods() {
	s.Run("WithEvaluator", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		evaluator := condition.NewEvaluator(nil, condition.DefaultEvalOptions())
		builder.WithEvaluator(evaluator)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		s.Require().NotNil(schema)
	})

	s.Run("AddFieldWithConfig", func() {
		builder := NewSchemaBuilder("test-schema", TypeForm, "Test Schema")
		customField := Field{
			Name:     "custom_field",
			Type:     FieldText,
			Label:    "Custom Field",
			Required: true,
			Help:     "This is a custom field",
		}
		builder.AddFieldWithConfig(customField)

		schema, err := builder.Build(s.ctx)

		s.Require().NoError(err)
		field, exists := schema.GetField("custom_field")
		s.Require().True(exists)
		s.Require().Equal("Custom Field", field.Label)
		s.Require().True(field.Required)
		s.Require().Equal("This is a custom field", field.Help)
	})
}
