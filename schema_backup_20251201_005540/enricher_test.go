package schema

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type EnricherTestSuite struct {
	suite.Suite
	enricher *DefaultEnricher
	ctx      context.Context
	logger   *MockLogger
}

func (suite *EnricherTestSuite) SetupTest() {
	suite.logger = NewMockLogger()
	suite.enricher = NewEnricherWithLogger(suite.logger)
	suite.ctx = context.Background()
}

func TestEnricherTestSuite(t *testing.T) {
	suite.Run(t, new(EnricherTestSuite))
}

// Mock implementations for testing

// MockLogger implements the Logger interface for testing
type MockLogger struct {
	WarnCalls  []LogCall
	ErrorCalls []LogCall
	DebugCalls []LogCall
}

type LogCall struct {
	Message string
	Fields  map[string]any
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		WarnCalls:  make([]LogCall, 0),
		ErrorCalls: make([]LogCall, 0),
		DebugCalls: make([]LogCall, 0),
	}
}

func (l *MockLogger) Warn(ctx context.Context, msg string, fields map[string]any) {
	l.WarnCalls = append(l.WarnCalls, LogCall{Message: msg, Fields: fields})
}

func (l *MockLogger) Error(ctx context.Context, msg string, fields map[string]any) {
	l.ErrorCalls = append(l.ErrorCalls, LogCall{Message: msg, Fields: fields})
}

func (l *MockLogger) Debug(ctx context.Context, msg string, fields map[string]any) {
	l.DebugCalls = append(l.DebugCalls, LogCall{Message: msg, Fields: fields})
}

func (l *MockLogger) Reset() {
	l.WarnCalls = make([]LogCall, 0)
	l.ErrorCalls = make([]LogCall, 0)
	l.DebugCalls = make([]LogCall, 0)
}

// MockTenantProvider implements TenantProvider for testing
type MockTenantProvider struct {
	customizations map[string]*TenantCustomization
	shouldError    bool
}

func NewMockTenantProvider() *MockTenantProvider {
	return &MockTenantProvider{
		customizations: make(map[string]*TenantCustomization),
		shouldError:    false,
	}
}

func (m *MockTenantProvider) GetCustomization(ctx context.Context, schemaID, tenantID string) (*TenantCustomization, error) {
	if m.shouldError {
		return nil, fmt.Errorf("mock error: failed to get customization")
	}

	key := schemaID + ":" + tenantID
	customization, exists := m.customizations[key]
	if !exists {
		return nil, nil
	}
	return customization, nil
}

func (m *MockTenantProvider) SetCustomization(schemaID, tenantID string, customization *TenantCustomization) {
	key := schemaID + ":" + tenantID
	m.customizations[key] = customization
}

func (m *MockTenantProvider) SetError(shouldError bool) {
	m.shouldError = shouldError
}

// Test Enricher Creation

func (suite *EnricherTestSuite) TestNewEnricher() {
	enricher := NewEnricher()
	suite.Require().NotNil(enricher)
	suite.Require().Nil(enricher.tenantProvider)
	suite.Require().NotNil(enricher.logger) // Should have noop logger
}

func (suite *EnricherTestSuite) TestNewEnricherWithLogger() {
	logger := NewMockLogger()
	enricher := NewEnricherWithLogger(logger)
	suite.Require().NotNil(enricher)
	suite.Require().Equal(logger, enricher.logger)
}

func (suite *EnricherTestSuite) TestNewEnricherWithNilLogger() {
	enricher := NewEnricherWithLogger(nil)
	suite.Require().NotNil(enricher)
	suite.Require().NotNil(enricher.logger) // Should fallback to noop logger
}

func (suite *EnricherTestSuite) TestSetTenantProvider() {
	provider := NewMockTenantProvider()
	suite.enricher.SetTenantProvider(provider)
	suite.Require().Equal(provider, suite.enricher.tenantProvider)
}

func (suite *EnricherTestSuite) TestSetLogger() {
	newLogger := NewMockLogger()
	suite.enricher.SetLogger(newLogger)
	suite.Require().Equal(newLogger, suite.enricher.logger)
}

// Test Basic Enrichment

func (suite *EnricherTestSuite) TestEnrichBasic() {
	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{"read", "write"},
		Roles:       []string{"user"},
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:     "username",
				Type:     FieldText,
				Label:    "Username",
				Required: true,
			},
			{
				Name:  "email",
				Type:  FieldEmail,
				Label: "Email",
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)

	// Verify original schema is unchanged
	suite.Require().Nil(schema.Fields[0].Runtime)
	suite.Require().Nil(schema.Fields[1].Runtime)

	// Check enriched schema has runtime data
	for i := range enriched.Fields {
		suite.Require().NotNil(enriched.Fields[i].Runtime)
		suite.Require().True(enriched.Fields[i].Runtime.Visible)
		suite.Require().True(enriched.Fields[i].Runtime.Editable)
	}

	// Verify debug logging occurred
	suite.Require().Greater(len(suite.logger.DebugCalls), 0)
	lastDebug := suite.logger.DebugCalls[len(suite.logger.DebugCalls)-1]
	suite.Require().Equal("schema enrichment completed successfully", lastDebug.Message)
}

func (suite *EnricherTestSuite) TestEnrichReturnsNewInstance() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	original := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Original",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, original, user)
	suite.Require().NoError(err)

	// Modify enriched schema
	enriched.Title = "Modified"
	enriched.Fields[0].Label = "Modified Label"

	// Original should be unchanged
	suite.Require().Equal("Original", original.Title)
	suite.Require().Equal("Field 1", original.Fields[0].Label)
}

// Test Permission-Based Enrichment

func (suite *EnricherTestSuite) TestEnrichWithPermissions() {
	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{"read"},
		Roles:       []string{"user"},
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:              "public_field",
				Type:              FieldText,
				Label:             "Public Field",
				RequirePermission: "",
			},
			{
				Name:              "admin_field",
				Type:              FieldText,
				Label:             "Admin Field",
				RequirePermission: "admin",
			},
			{
				Name:         "role_field",
				Type:         FieldText,
				Label:        "Role Field",
				RequireRoles: []string{"admin", "manager"},
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Public field should be visible
	suite.Require().True(enriched.Fields[0].Runtime.Visible)
	suite.Require().True(enriched.Fields[0].Runtime.Editable)

	// Admin field should not be visible
	suite.Require().False(enriched.Fields[1].Runtime.Visible)
	suite.Require().False(enriched.Fields[1].Runtime.Editable)
	suite.Require().Equal("permission_required", enriched.Fields[1].Runtime.Reason)

	// Role field should not be visible
	suite.Require().False(enriched.Fields[2].Runtime.Visible)
	suite.Require().False(enriched.Fields[2].Runtime.Editable)
	suite.Require().Equal("role_required", enriched.Fields[2].Runtime.Reason)
}

func (suite *EnricherTestSuite) TestEnrichWithMultipleRoles() {
	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{},
		Roles:       []string{"manager"},
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:         "manager_or_admin_field",
				Type:         FieldText,
				Label:        "Manager/Admin Field",
				RequireRoles: []string{"admin", "manager"}, // User has manager, so should be visible
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Field should be visible (user has manager role)
	suite.Require().True(enriched.Fields[0].Runtime.Visible)
	suite.Require().True(enriched.Fields[0].Runtime.Editable)
}

// Test Dynamic Defaults

func (suite *EnricherTestSuite) TestEnrichDynamicDefaults() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "created_by", Type: FieldText, Label: "Created By"},
			{Name: "user_id", Type: FieldText, Label: "User ID"},
			{Name: "author_id", Type: FieldText, Label: "Author ID"},
			{Name: "tenant_id", Type: FieldText, Label: "Tenant ID"},
			{Name: "organization_id", Type: FieldText, Label: "Organization ID"},
			{Name: "created_at", Type: FieldDateTime, Label: "Created At"},
			{Name: "timestamp", Type: FieldDateTime, Label: "Timestamp"},
			{Name: "updated_at", Type: FieldDateTime, Label: "Updated At"},
			{Name: "static_default", Type: FieldText, Label: "Static Default", Default: "static_value"},
			{Name: "existing_value", Type: FieldText, Label: "Existing Value", Value: "existing"},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Check user-related defaults
	suite.Require().Equal("user123", enriched.Fields[0].Value) // created_by
	suite.Require().Equal("user123", enriched.Fields[1].Value) // user_id
	suite.Require().Equal("user123", enriched.Fields[2].Value) // author_id

	// Check tenant-related defaults
	suite.Require().Equal("tenant456", enriched.Fields[3].Value) // tenant_id (also hidden/readonly)
	suite.Require().True(enriched.Fields[3].Hidden)
	suite.Require().True(enriched.Fields[3].Readonly)
	suite.Require().Equal("tenant456", enriched.Fields[4].Value) // organization_id

	// Check timestamp defaults
	suite.Require().NotNil(enriched.Fields[5].Value) // created_at
	suite.Require().NotNil(enriched.Fields[6].Value) // timestamp
	suite.Require().NotNil(enriched.Fields[7].Value) // updated_at

	// Static default should be applied
	suite.Require().Equal("static_value", enriched.Fields[8].Value)

	// Existing value should not be overridden
	suite.Require().Equal("existing", enriched.Fields[9].Value)
}

// Test Tenant Customization

func (suite *EnricherTestSuite) TestEnrichWithTenantCustomization() {
	provider := NewMockTenantProvider()
	suite.enricher.SetTenantProvider(provider)

	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	customization := &TenantCustomization{
		SchemaID: "test-schema",
		TenantID: "tenant456",
		Overrides: map[string]TenantOverride{
			"username": {
				FieldName:   "username",
				Label:       "Custom Username Label",
				Required:    boolPtr(false),
				Placeholder: "Enter custom username",
				Help:        "Custom help text",
			},
			"email": {
				FieldName:    "email",
				Hidden:       boolPtr(true),
				DefaultValue: "default@tenant.com",
			},
		},
	}
	provider.SetCustomization("test-schema", "tenant456", customization)

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:        "username",
				Type:        FieldText,
				Label:       "Original Username",
				Required:    true,
				Placeholder: "Original placeholder",
			},
			{
				Name:  "email",
				Type:  FieldEmail,
				Label: "Email",
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Check username customization
	suite.Require().Equal("Custom Username Label", enriched.Fields[0].Label)
	suite.Require().False(enriched.Fields[0].Required)
	suite.Require().Equal("Enter custom username", enriched.Fields[0].Placeholder)
	suite.Require().Equal("Custom help text", enriched.Fields[0].Help)

	// Check email customization
	suite.Require().True(enriched.Fields[1].Hidden)
	suite.Require().Equal("default@tenant.com", enriched.Fields[1].Default)

	// Original schema should be unchanged
	suite.Require().Equal("Original Username", schema.Fields[0].Label)
	suite.Require().True(schema.Fields[0].Required)
}

func (suite *EnricherTestSuite) TestTenantCustomizationNotFound() {
	provider := NewMockTenantProvider()
	suite.enricher.SetTenantProvider(provider)

	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	// No customization set for this schema/tenant
	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)

	// Should still enrich successfully with defaults
	suite.Require().NotNil(enriched.Fields[0].Runtime)
}

func (suite *EnricherTestSuite) TestTenantProviderError() {
	provider := NewMockTenantProvider()
	provider.SetError(true)
	suite.enricher.SetTenantProvider(provider)

	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	// Should succeed even if provider errors
	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)

	// Should have logged a warning
	suite.Require().Greater(len(suite.logger.WarnCalls), 0)
	lastWarn := suite.logger.WarnCalls[len(suite.logger.WarnCalls)-1]
	suite.Require().Equal("failed to retrieve tenant customization", lastWarn.Message)
}

// Test Action Enrichment

func (suite *EnricherTestSuite) TestEnrichActions() {
	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{"read", "write"},
		Roles:       []string{"user"},
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Actions: []Action{
			{
				ID:   "public-action",
				Type: ActionButton,
				Text: "Public Action",
			},
			{
				ID:          "admin-action",
				Type:        ActionButton,
				Text:        "Admin Action",
				Permissions: []string{"admin"},
			},
			{
				ID:          "user-action",
				Type:        ActionButton,
				Text:        "User Action",
				Permissions: []string{"read", "write"},
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Public action should remain unchanged
	suite.Require().False(enriched.Actions[0].Hidden)
	suite.Require().False(enriched.Actions[0].Disabled)

	// Admin action should be hidden and disabled
	suite.Require().True(enriched.Actions[1].Hidden)
	suite.Require().True(enriched.Actions[1].Disabled)

	// User action should be visible and enabled
	suite.Require().False(enriched.Actions[2].Hidden)
	suite.Require().False(enriched.Actions[2].Disabled)
}

func (suite *EnricherTestSuite) TestEnrichActionsViewOnly() {
	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{"read"}, // Only read permission
		Roles:       []string{"user"},
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Actions: []Action{
			{
				ID:          "readonly-action",
				Type:        ActionButton,
				Text:        "Read-Only Action",
				Permissions: []string{"read", "write"},
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Action should be visible but disabled
	suite.Require().False(enriched.Actions[0].Hidden)
	suite.Require().True(enriched.Actions[0].Disabled)
}

// Test Enrichment with Options

func (suite *EnricherTestSuite) TestEnrichWithOptions() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
			{Name: "field2", Type: FieldText, Label: "Field 2"},
		},
	}

	enriched, err := suite.enricher.EnrichWithOptions(
		suite.ctx,
		schema,
		WithUser(user),
		WithEnvironment("development"),
		WithFeature("beta_feature", true),
	)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)
}

func (suite *EnricherTestSuite) TestEnrichWithFieldFilter() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "include_me", Type: FieldText, Label: "Include"},
			{Name: "exclude_me", Type: FieldText, Label: "Exclude"},
		},
	}

	filter := func(f *Field) bool {
		return f.Name == "include_me"
	}

	enriched, err := suite.enricher.EnrichWithOptions(
		suite.ctx,
		schema,
		WithUser(user),
		WithFieldFilter(filter),
	)
	suite.Require().NoError(err)

	// Only filtered field should have runtime data
	suite.Require().NotNil(enriched.Fields[0].Runtime) // include_me
	suite.Require().Nil(enriched.Fields[1].Runtime)    // exclude_me (not processed)
}

func (suite *EnricherTestSuite) TestEnrichWithSkipValidation() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	// Create an invalid schema that would normally fail validation
	schema := &Schema{
		ID:    "", // Invalid: empty ID
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	// Without skip validation, should fail
	_, err := suite.enricher.EnrichWithOptions(suite.ctx, schema, WithUser(user))
	suite.Require().Error(err)

	// With skip validation, should succeed
	enriched, err := suite.enricher.EnrichWithOptions(
		suite.ctx,
		schema,
		WithUser(user),
		WithSkipValidation(true),
	)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)
}

func (suite *EnricherTestSuite) TestEnrichWithExtensions() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	enriched, err := suite.enricher.EnrichWithOptions(
		suite.ctx,
		schema,
		WithUser(user),
		WithExtension("custom_key", "custom_value"),
		WithExtension("request_id", "req-123"),
	)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)
}

// Test Localization

func (suite *EnricherTestSuite) TestEnrichWithLocale() {
	user := &BasicUser{
		ID:              "user123",
		TenantID:        "tenant456",
		PreferredLocale: "es",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:  "username",
				Type:  FieldText,
				Label: "Username",
			},
		},
	}

	// This test assumes localization functions are available
	// The actual behavior depends on your localization implementation
	enriched, err := suite.enricher.EnrichWithLocale(suite.ctx, schema, user, "es")
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)
}

func (suite *EnricherTestSuite) TestEnrichWithLocaleUsesUserPreference() {
	user := &BasicUser{
		ID:              "user123",
		TenantID:        "tenant456",
		PreferredLocale: "fr",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{Name: "field1", Type: FieldText, Label: "Field 1"},
		},
	}

	// Empty locale should use user's preferred locale
	enriched, err := suite.enricher.EnrichWithLocale(suite.ctx, schema, user, "")
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)
}

// Test Error Handling

func (suite *EnricherTestSuite) TestEnrichNilSchema() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	_, err := suite.enricher.Enrich(suite.ctx, nil, user)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "schema cannot be nil")
}

func (suite *EnricherTestSuite) TestEnrichNilUser() {
	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
	}

	_, err := suite.enricher.Enrich(suite.ctx, schema, nil)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "user cannot be nil")
}

func (suite *EnricherTestSuite) TestEnrichWithOptionsNilSchema() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	_, err := suite.enricher.EnrichWithOptions(suite.ctx, nil, WithUser(user))
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "schema cannot be nil")
}

func (suite *EnricherTestSuite) TestEnrichWithOptionsNoUser() {
	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
	}

	_, err := suite.enricher.EnrichWithOptions(suite.ctx, schema)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "no user information provided")
}

func (suite *EnricherTestSuite) TestEnrichFieldNilField() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}
	data := map[string]any{}

	err := suite.enricher.EnrichField(suite.ctx, nil, user, data)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "field cannot be nil")
}

func (suite *EnricherTestSuite) TestEnrichFieldNilUser() {
	field := &Field{
		Name:  "test_field",
		Type:  FieldText,
		Label: "Test",
	}
	data := map[string]any{}

	err := suite.enricher.EnrichField(suite.ctx, field, nil, data)
	suite.Require().Error(err)
	suite.Require().Contains(err.Error(), "user cannot be nil")
}

// Test BasicUser Implementation

func (suite *EnricherTestSuite) TestBasicUserInterface() {
	user := &BasicUser{
		ID:              "user123",
		TenantID:        "tenant456",
		Permissions:     []string{"read", "write", "admin"},
		Roles:           []string{"user", "manager"},
		PreferredLocale: "en-US",
	}

	suite.Require().Equal("user123", user.GetID())
	suite.Require().Equal("tenant456", user.GetTenantID())
	suite.Require().Equal([]string{"read", "write", "admin"}, user.GetPermissions())
	suite.Require().Equal([]string{"user", "manager"}, user.GetRoles())
	suite.Require().Equal("en-US", user.GetPreferredLocale())

	// Test permission checks
	suite.Require().True(user.HasPermission("read"))
	suite.Require().True(user.HasPermission("write"))
	suite.Require().True(user.HasPermission("admin"))
	suite.Require().False(user.HasPermission("super_admin"))

	// Test role checks
	suite.Require().True(user.HasRole("user"))
	suite.Require().True(user.HasRole("manager"))
	suite.Require().False(user.HasRole("admin"))
}

func (suite *EnricherTestSuite) TestBasicUserDefaultLocale() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	// Should default to "en"
	suite.Require().Equal("en", user.GetPreferredLocale())
}

// Test Readonly Fields

func (suite *EnricherTestSuite) TestEnrichReadonlyField() {
	user := &BasicUser{
		ID:       "user123",
		TenantID: "tenant456",
	}

	schema := &Schema{
		ID:    "test-schema",
		Type:  TypeForm,
		Title: "Test Schema",
		Fields: []Field{
			{
				Name:     "readonly_field",
				Type:     FieldText,
				Label:    "Readonly Field",
				Readonly: true,
			},
			{
				Name:  "editable_field",
				Type:  FieldText,
				Label: "Editable Field",
			},
		},
	}

	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)

	// Readonly field should be visible but not editable
	suite.Require().True(enriched.Fields[0].Runtime.Visible)
	suite.Require().False(enriched.Fields[0].Runtime.Editable)

	// Editable field should be visible and editable
	suite.Require().True(enriched.Fields[1].Runtime.Visible)
	suite.Require().True(enriched.Fields[1].Runtime.Editable)
}

// Comprehensive Integration Test

func (suite *EnricherTestSuite) TestComprehensiveEnrichment() {
	// Setup
	provider := NewMockTenantProvider()
	suite.enricher.SetTenantProvider(provider)

	user := &BasicUser{
		ID:          "user123",
		TenantID:    "tenant456",
		Permissions: []string{"read", "write", "user_management"},
		Roles:       []string{"manager"},
	}

	customization := &TenantCustomization{
		SchemaID: "comprehensive-schema",
		TenantID: "tenant456",
		Overrides: map[string]TenantOverride{
			"company_name": {
				FieldName:    "company_name",
				Label:        "Organization Name",
				Required:     boolPtr(true),
				DefaultValue: "Default Org",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	provider.SetCustomization("comprehensive-schema", "tenant456", customization)

	schema := &Schema{
		ID:    "comprehensive-schema",
		Type:  TypeForm,
		Title: "Comprehensive Test Schema",
		Fields: []Field{
			{Name: "created_by", Type: FieldText, Label: "Created By"},
			{Name: "tenant_id", Type: FieldText, Label: "Tenant ID"},
			{Name: "admin_notes", Type: FieldTextarea, Label: "Admin Notes", RequirePermission: "admin"},
			{Name: "manager_field", Type: FieldText, Label: "Manager Field", RequireRoles: []string{"manager", "admin"}},
			{Name: "company_name", Type: FieldText, Label: "Company Name", Required: false, Default: "Original Default"},
			{Name: "readonly_field", Type: FieldText, Label: "Readonly", Readonly: true},
		},
		Actions: []Action{
			{ID: "submit", Type: ActionSubmit, Text: "Submit"},
			{ID: "delete", Type: ActionButton, Text: "Delete", Permissions: []string{"admin"}},
			{ID: "approve", Type: ActionButton, Text: "Approve", Permissions: []string{"manager"}},
		},
	}

	// Execute
	enriched, err := suite.enricher.Enrich(suite.ctx, schema, user)
	suite.Require().NoError(err)
	suite.Require().NotNil(enriched)

	// Verify field enrichment
	suite.Require().Equal("user123", enriched.Fields[0].Value)                      // created_by
	suite.Require().Equal("tenant456", enriched.Fields[1].Value)                    // tenant_id
	suite.Require().True(enriched.Fields[1].Hidden)                                 // tenant_id hidden
	suite.Require().True(enriched.Fields[1].Readonly)                               // tenant_id readonly
	suite.Require().False(enriched.Fields[2].Runtime.Visible)                       // admin_notes not visible
	suite.Require().Equal("permission_required", enriched.Fields[2].Runtime.Reason) // admin_notes reason
	suite.Require().True(enriched.Fields[3].Runtime.Visible)                        // manager_field visible
	suite.Require().Equal("Organization Name", enriched.Fields[4].Label)            // tenant customization applied
	suite.Require().True(enriched.Fields[4].Required)                               // tenant customization applied
	suite.Require().Equal("Default Org", enriched.Fields[4].Default)                // tenant customization applied
	suite.Require().True(enriched.Fields[5].Runtime.Visible)                        // readonly visible
	suite.Require().False(enriched.Fields[5].Runtime.Editable)                      // readonly not editable

	// Verify action enrichment
	suite.Require().Len(enriched.Actions, 3)
	suite.Require().False(enriched.Actions[0].Hidden)   // submit visible
	suite.Require().False(enriched.Actions[0].Disabled) // submit enabled
	suite.Require().False(enriched.Actions[1].Hidden)   // delete visible (no view permission set)
	suite.Require().True(enriched.Actions[1].Disabled)  // delete disabled (no admin permission)
	suite.Require().False(enriched.Actions[2].Hidden)   // approve visible (has manager role)
	suite.Require().False(enriched.Actions[2].Disabled) // approve enabled (has manager role)

	// Verify logging
	suite.Require().Greater(len(suite.logger.DebugCalls), 0)

	// Verify original unchanged
	suite.Require().Equal("Company Name", schema.Fields[4].Label)
	suite.Require().False(schema.Fields[4].Required)
}

// Helper Functions

func boolPtr(b bool) *bool {
	return &b
}
