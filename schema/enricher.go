package schema

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// EnrichOption is a functional option for enrichment configuration
type EnrichOption func(*EnrichConfig)

// EnrichConfig holds configuration for schema enrichment operations
type EnrichConfig struct {
	User           uuid.UUID
	Environment    string
	Features       map[string]bool
	FieldFilter    func(*Field) bool
	DepthLimit     int
	SkipValidation bool
	Extensions     map[string]any
}

// User represents a user in the system for enrichment purposes
type User interface {
	GetID() string
	GetTenantID() string
	GetPermissions() []string
	GetRoles() []string
	HasPermission(permission string) bool
	HasRole(role string) bool
	GetPreferredLocale() string
}

// TenantOverride represents tenant-specific customizations for a field
type TenantOverride struct {
	FieldName    string `json:"field_name"`
	Label        string `json:"label,omitempty"`
	Required     *bool  `json:"required,omitempty"`
	Hidden       *bool  `json:"hidden,omitempty"`
	DefaultValue any    `json:"default_value,omitempty"`
	Placeholder  string `json:"placeholder,omitempty"`
	Help         string `json:"help,omitempty"`
}

// TenantCustomization holds all tenant-specific overrides for a schema
type TenantCustomization struct {
	SchemaID  string                    `json:"schema_id"`
	TenantID  string                    `json:"tenant_id"`
	Overrides map[string]TenantOverride `json:"overrides"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

// TenantProvider interface for retrieving tenant customizations
type TenantProvider interface {
	GetCustomization(ctx context.Context, schemaID, tenantID string) (*TenantCustomization, error)
}

// Logger interface for structured logging in the enricher
type Logger interface {
	Warn(ctx context.Context, msg string, fields map[string]any)
	Error(ctx context.Context, msg string, fields map[string]any)
	Debug(ctx context.Context, msg string, fields map[string]any)
}

// Enricher interface defines the contract for schema enrichment
type Enricher interface {
	// Enrich enriches a schema with runtime data based on user context
	// Returns a new enriched schema instance, leaving the original unchanged
	Enrich(ctx context.Context, schema *Schema, user User) (*Schema, error)

	// EnrichField enriches a single field with runtime data
	EnrichField(ctx context.Context, field *Field, user User, data map[string]any) error

	// EnrichWithLocale enriches a schema with localization support
	EnrichWithLocale(ctx context.Context, schema *Schema, user User, locale string) (*Schema, error)

	// SetTenantProvider sets the tenant customization provider
	SetTenantProvider(provider TenantProvider)
}

// DefaultEnricher is the production-ready implementation of the Enricher interface
type DefaultEnricher struct {
	tenantProvider TenantProvider
	logger         Logger
}

// BasicUser is a simple implementation of the User interface
type BasicUser struct {
	ID              string
	TenantID        string
	Permissions     []string
	Roles           []string
	PreferredLocale string
}

// NewEnricher creates a new DefaultEnricher with default settings
func NewEnricher() *DefaultEnricher {
	return &DefaultEnricher{
		logger: &noopLogger{},
	}
}

// NewEnricherWithLogger creates a new DefaultEnricher with a custom logger
func NewEnricherWithLogger(logger Logger) *DefaultEnricher {
	if logger == nil {
		logger = &noopLogger{}
	}
	return &DefaultEnricher{
		logger: logger,
	}
}

// SetTenantProvider sets the tenant customization provider
func (e *DefaultEnricher) SetTenantProvider(provider TenantProvider) {
	e.tenantProvider = provider
}

// SetLogger sets the logger for the enricher
func (e *DefaultEnricher) SetLogger(logger Logger) {
	if logger != nil {
		e.logger = logger
	}
}

// Enrich enriches a schema with runtime data based on user context
// This method always returns a new schema instance, leaving the original unchanged
func (e *DefaultEnricher) Enrich(ctx context.Context, schema *Schema, user User) (*Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}
	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	// Clone the schema to avoid modifying the original
	enriched, err := schema.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone schema: %w", err)
	}

	// Get tenant customizations if provider is available
	var customization *TenantCustomization
	if e.tenantProvider != nil {
		customization, err = e.tenantProvider.GetCustomization(ctx, enriched.ID, user.GetTenantID())
		if err != nil {
			e.logger.Warn(ctx, "failed to retrieve tenant customization", map[string]any{
				"schema_id": enriched.ID,
				"tenant_id": user.GetTenantID(),
				"error":     err.Error(),
			})
			// Continue with enrichment even if customization retrieval fails
		}
	}

	// Create data context for dynamic defaults and enrichment
	dataContext := e.buildDataContext(user)

	// Enrich each field in the schema
	for i := range enriched.Fields {
		field := &enriched.Fields[i]

		// Apply tenant customizations first (these take precedence)
		if customization != nil {
			e.applyTenantCustomization(field, customization)
		}

		// Enrich the field with runtime data
		if err := e.EnrichField(ctx, field, user, dataContext); err != nil {
			return nil, fmt.Errorf("failed to enrich field '%s': %w", field.Name, err)
		}
	}

	// Enrich actions with permission checks
	for i := range enriched.Actions {
		action := &enriched.Actions[i]
		e.enrichAction(ctx, action, user)
	}

	// Validate the enriched schema
	if err := enriched.Validate(ctx); err != nil {
		return nil, fmt.Errorf("enriched schema validation failed: %w", err)
	}

	e.logger.Debug(ctx, "schema enrichment completed successfully", map[string]any{
		"schema_id":    enriched.ID,
		"tenant_id":    user.GetTenantID(),
		"field_count":  len(enriched.Fields),
		"action_count": len(enriched.Actions),
	})

	return enriched, nil
}

// EnrichWithOptions enriches a schema using functional options for advanced configuration
// This provides more flexibility than the standard Enrich method
func (e *DefaultEnricher) EnrichWithOptions(ctx context.Context, schema *Schema, opts ...EnrichOption) (*Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}

	// Build configuration from options
	config := e.buildEnrichConfig(opts...)

	// Extract user from configuration
	user, err := e.extractUserFromConfig(config)
	if err != nil {
		return nil, fmt.Errorf("invalid user configuration: %w", err)
	}

	// Clone the schema to avoid modifying the original
	enriched, err := schema.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone schema: %w", err)
	}

	// Get tenant customizations if provider is available
	var customization *TenantCustomization
	if e.tenantProvider != nil {
		customization, err = e.tenantProvider.GetCustomization(ctx, enriched.ID, user.GetTenantID())
		if err != nil {
			e.logger.Warn(ctx, "failed to retrieve tenant customization", map[string]any{
				"schema_id": enriched.ID,
				"tenant_id": user.GetTenantID(),
				"error":     err.Error(),
			})
			// Continue with enrichment
		}
	}

	// Create data context with configuration extensions
	dataContext := e.buildDataContextWithConfig(user, config)

	// Determine which fields to process based on filter
	fieldsToProcess := enriched.Fields
	if config.FieldFilter != nil {
		fieldsToProcess = e.filterFields(enriched.Fields, config.FieldFilter)
	}

	// Enrich each field
	for i := range fieldsToProcess {
		field := &fieldsToProcess[i]

		// Apply tenant customizations
		if customization != nil {
			e.applyTenantCustomization(field, customization)
		}

		// Enrich the field
		if err := e.EnrichField(ctx, field, user, dataContext); err != nil {
			if !config.SkipValidation {
				return nil, fmt.Errorf("failed to enrich field '%s': %w", field.Name, err)
			}
			e.logger.Warn(ctx, "field enrichment failed but continuing due to SkipValidation", map[string]any{
				"field_name": field.Name,
				"error":      err.Error(),
			})
		}
	}

	// Enrich actions
	for i := range enriched.Actions {
		action := &enriched.Actions[i]
		e.enrichAction(ctx, action, user)
	}

	// Validate enriched schema if not skipped
	if !config.SkipValidation {
		if err := enriched.Validate(ctx); err != nil {
			return nil, fmt.Errorf("enriched schema validation failed: %w", err)
		}
	}

	return enriched, nil
}

// EnrichWithLocale enriches a schema with localization support
// Returns a new localized and enriched schema instance
func (e *DefaultEnricher) EnrichWithLocale(ctx context.Context, schema *Schema, user User, locale string) (*Schema, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}
	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	// Use user's preferred locale if none specified
	if locale == "" {
		locale = user.GetPreferredLocale()
	}
	if locale == "" {
		locale = "en" // Default to English
	}

	// Validate the locale is available
	if !T_HasLocale(locale) {
		originalLocale := locale
		locale = T_DetectLocale([]string{locale, "en"})
		e.logger.Debug(ctx, "requested locale not available, using fallback", map[string]any{
			"requested_locale": originalLocale,
			"fallback_locale":  locale,
		})
	}

	// Clone the schema to avoid modifying the original
	localizedSchema, err := schema.Clone()
	if err != nil {
		return nil, fmt.Errorf("failed to clone schema: %w", err)
	}

	// Apply localization to all fields using embedded translations
	for i := range localizedSchema.Fields {
		field := &localizedSchema.Fields[i]

		// Apply embedded translations to field labels and texts
		field.Label = field.GetLocalizedLabel(locale)
		field.Placeholder = field.GetLocalizedPlaceholder(locale)
		field.Help = field.GetLocalizedHelp(locale)
		field.Description = field.GetLocalizedDescription(locale)
	}

	// Apply localization to actions
	for i := range localizedSchema.Actions {
		action := &localizedSchema.Actions[i]
		action.Text = T_Action(locale, action.ID)
	}

	// Now enrich with user context
	enrichedSchema, err := e.Enrich(ctx, localizedSchema, user)
	if err != nil {
		return nil, fmt.Errorf("enrichment failed after localization: %w", err)
	}

	return enrichedSchema, nil
}

// EnrichField enriches a single field with runtime data
func (e *DefaultEnricher) EnrichField(ctx context.Context, field *Field, user User, data map[string]any) error {
	if field == nil {
		return fmt.Errorf("field cannot be nil")
	}
	if user == nil {
		return fmt.Errorf("user cannot be nil")
	}

	// Initialize runtime if not exists
	if field.Runtime == nil {
		field.Runtime = &FieldRuntime{}
	}

	// Apply enrichment in order of precedence
	e.applyPermissions(field, user)
	e.applyDynamicDefaults(field, data)
	e.applyTenantIsolation(field, user)

	return nil
}

// applyPermissions applies permission-based field visibility and editability restrictions
func (e *DefaultEnricher) applyPermissions(field *Field, user User) {
	// Initialize with default visibility and editability
	field.Runtime.Visible = true
	field.Runtime.Editable = !field.Readonly

	// Check permission-based restrictions first
	if field.RequirePermission != "" {
		if !user.HasPermission(field.RequirePermission) {
			field.Runtime.Visible = false
			field.Runtime.Editable = false
			field.Runtime.Reason = "permission_required"
			return
		}
	}

	// Check role-based restrictions
	if len(field.RequireRoles) > 0 {
		if !e.hasAnyRole(user, field.RequireRoles) {
			field.Runtime.Visible = false
			field.Runtime.Editable = false
			field.Runtime.Reason = "role_required"
			return
		}
	}
}

// applyDynamicDefaults applies dynamic default values based on context
func (e *DefaultEnricher) applyDynamicDefaults(field *Field, data map[string]any) {
	// Only apply defaults if field has no current value
	if field.Value != nil {
		return
	}

	// Apply static default value if exists
	if field.Default != nil {
		field.Value = field.Default
		return
	}

	// Apply conventional dynamic defaults based on field name patterns
	switch field.Name {
	case "created_by", "user_id", "author_id":
		if userID, ok := data["user_id"].(string); ok {
			field.Value = userID
		}
	case "tenant_id", "organization_id":
		if tenantID, ok := data["tenant_id"].(string); ok {
			field.Value = tenantID
		}
	case "created_at", "timestamp":
		if timestamp, ok := data["timestamp"].(time.Time); ok {
			field.Value = timestamp.Format(time.RFC3339)
		}
	case "updated_at":
		field.Value = time.Now().Format(time.RFC3339)
	}
}

// applyTenantIsolation applies tenant isolation for multi-tenant fields
func (e *DefaultEnricher) applyTenantIsolation(field *Field, user User) {
	// Automatically configure tenant_id field for tenant isolation
	if field.Name == "tenant_id" {
		field.Value = user.GetTenantID()
		field.Hidden = true
		field.Readonly = true
	}
}

// applyTenantCustomization applies tenant-specific customizations to a field
func (e *DefaultEnricher) applyTenantCustomization(field *Field, customization *TenantCustomization) {
	override, exists := customization.Overrides[field.Name]
	if !exists {
		return
	}

	// Apply overrides in a specific order
	if override.Label != "" {
		field.Label = override.Label
	}

	if override.Required != nil {
		field.Required = *override.Required
	}

	if override.Hidden != nil {
		field.Hidden = *override.Hidden
	}

	if override.DefaultValue != nil {
		field.Default = override.DefaultValue
	}

	if override.Placeholder != "" {
		field.Placeholder = override.Placeholder
	}

	if override.Help != "" {
		field.Help = override.Help
	}
}

// enrichAction enriches an action with user permission checks
func (e *DefaultEnricher) enrichAction(ctx context.Context, action *Action, user User) {
	if len(action.Permissions) == 0 {
		return
	}

	// Check permissions - unified approach
	canExecute := e.hasAnyPermission(user, action.Permissions)
	action.Disabled = !canExecute
	action.Hidden = !canExecute
}

// buildDataContext creates a basic data context for enrichment
func (e *DefaultEnricher) buildDataContext(user User) map[string]any {
	return map[string]any{
		"user_id":   user.GetID(),
		"tenant_id": user.GetTenantID(),
		"timestamp": time.Now(),
	}
}

// buildDataContextWithConfig creates a data context with configuration extensions
func (e *DefaultEnricher) buildDataContextWithConfig(user User, config *EnrichConfig) map[string]any {
	dataContext := map[string]any{
		"user_id":     user.GetID(),
		"tenant_id":   user.GetTenantID(),
		"timestamp":   time.Now(),
		"environment": config.Environment,
		"features":    config.Features,
	}

	// Add extensions to data context, avoiding core key collisions
	for key, value := range config.Extensions {
		if key != "user" && key != "user_id" && key != "tenant_id" {
			dataContext[key] = value
		}
	}

	return dataContext
}

// buildEnrichConfig builds enrichment configuration from functional options
func (e *DefaultEnricher) buildEnrichConfig(opts ...EnrichOption) *EnrichConfig {
	config := &EnrichConfig{
		User:           uuid.Nil,
		Environment:    "production",
		Features:       make(map[string]bool),
		FieldFilter:    nil,
		DepthLimit:     10,
		SkipValidation: false,
		Extensions:     make(map[string]any),
	}

	// Apply all options to build configuration
	for _, opt := range opts {
		opt(config)
	}

	return config
}

// extractUserFromConfig extracts user from enrichment configuration
func (e *DefaultEnricher) extractUserFromConfig(config *EnrichConfig) (User, error) {
	// Try to get full user interface from extensions first
	if userInterface, exists := config.Extensions["user"]; exists {
		if user, ok := userInterface.(User); ok {
			return user, nil
		}
		return nil, fmt.Errorf("invalid user type in extensions, must implement User interface")
	}

	// Create BasicUser from UUID if available
	if config.User != uuid.Nil {
		return &BasicUser{
			ID:              config.User.String(),
			TenantID:        e.extractTenantFromConfig(config),
			Permissions:     e.extractPermissionsFromConfig(config),
			Roles:           e.extractRolesFromConfig(config),
			PreferredLocale: e.extractLocaleFromConfig(config),
		}, nil
	}

	return nil, fmt.Errorf("no user information provided in enrichment configuration")
}

// extractTenantFromConfig extracts tenant ID from enrichment config
func (e *DefaultEnricher) extractTenantFromConfig(config *EnrichConfig) string {
	if tenantInterface, exists := config.Extensions["tenant_id"]; exists {
		if tenantID, ok := tenantInterface.(string); ok {
			return tenantID
		}
	}
	if tenantInterface, exists := config.Extensions["tenant"]; exists {
		if tenantID, ok := tenantInterface.(string); ok {
			return tenantID
		}
	}
	return "default"
}

// extractPermissionsFromConfig extracts permissions from enrichment config
func (e *DefaultEnricher) extractPermissionsFromConfig(config *EnrichConfig) []string {
	if permInterface, exists := config.Extensions["permissions"]; exists {
		if permissions, ok := permInterface.([]string); ok {
			return permissions
		}
		// Handle []any conversion
		if permissions, ok := permInterface.([]any); ok {
			result := make([]string, 0, len(permissions))
			for _, perm := range permissions {
				if permStr, ok := perm.(string); ok {
					result = append(result, permStr)
				}
			}
			return result
		}
	}
	return []string{}
}

// extractRolesFromConfig extracts roles from enrichment config
func (e *DefaultEnricher) extractRolesFromConfig(config *EnrichConfig) []string {
	if roleInterface, exists := config.Extensions["roles"]; exists {
		if roles, ok := roleInterface.([]string); ok {
			return roles
		}
		// Handle []any conversion
		if roles, ok := roleInterface.([]any); ok {
			result := make([]string, 0, len(roles))
			for _, role := range roles {
				if roleStr, ok := role.(string); ok {
					result = append(result, roleStr)
				}
			}
			return result
		}
	}
	return []string{}
}

// extractLocaleFromConfig extracts locale from enrichment config
func (e *DefaultEnricher) extractLocaleFromConfig(config *EnrichConfig) string {
	if localeInterface, exists := config.Extensions["locale"]; exists {
		if locale, ok := localeInterface.(string); ok {
			return locale
		}
	}
	return "en"
}

// filterFields filters fields based on the provided filter function
func (e *DefaultEnricher) filterFields(fields []Field, filter func(*Field) bool) []Field {
	filtered := make([]Field, 0, len(fields))
	for i := range fields {
		if filter(&fields[i]) {
			filtered = append(filtered, fields[i])
		}
	}
	return filtered
}

// hasAnyPermission checks if user has any of the required permissions
func (e *DefaultEnricher) hasAnyPermission(user User, required []string) bool {
	if len(required) == 0 {
		return true
	}

	userPerms := user.GetPermissions()
	permSet := make(map[string]bool, len(userPerms))
	for _, perm := range userPerms {
		permSet[perm] = true
	}

	for _, req := range required {
		if permSet[req] {
			return true
		}
	}
	return false
}

// hasAnyRole checks if user has any of the required roles
func (e *DefaultEnricher) hasAnyRole(user User, required []string) bool {
	if len(required) == 0 {
		return true
	}

	userRoles := user.GetRoles()
	roleSet := make(map[string]bool, len(userRoles))
	for _, role := range userRoles {
		roleSet[role] = true
	}

	for _, req := range required {
		if roleSet[req] {
			return true
		}
	}
	return false
}

// BasicUser interface implementation

// GetID returns the user ID
func (u *BasicUser) GetID() string {
	return u.ID
}

// GetTenantID returns the tenant ID
func (u *BasicUser) GetTenantID() string {
	return u.TenantID
}

// GetPermissions returns user permissions
func (u *BasicUser) GetPermissions() []string {
	return u.Permissions
}

// GetRoles returns user roles
func (u *BasicUser) GetRoles() []string {
	return u.Roles
}

// GetPreferredLocale returns user's preferred locale
func (u *BasicUser) GetPreferredLocale() string {
	if u.PreferredLocale != "" {
		return u.PreferredLocale
	}
	return "en"
}

// HasPermission checks if user has specific permission
func (u *BasicUser) HasPermission(permission string) bool {
	for _, perm := range u.Permissions {
		if perm == permission {
			return true
		}
	}
	return false
}

// HasRole checks if user has specific role
func (u *BasicUser) HasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

// Functional options for EnrichWithOptions

// WithUser creates an EnrichOption that sets the full user in extensions
func WithUser(user User) EnrichOption {
	return func(config *EnrichConfig) {
		config.Extensions["user"] = user
	}
}

// WithUserID creates an EnrichOption that sets the user ID
func WithUserID(userID uuid.UUID) EnrichOption {
	return func(config *EnrichConfig) {
		config.User = userID
	}
}

// WithTenantID creates an EnrichOption that sets the tenant ID
func WithTenantID(tenantID string) EnrichOption {
	return func(config *EnrichConfig) {
		config.Extensions["tenant_id"] = tenantID
	}
}

// WithPermissions creates an EnrichOption that sets permissions
func WithPermissions(permissions []string) EnrichOption {
	return func(config *EnrichConfig) {
		config.Extensions["permissions"] = permissions
	}
}

// WithRoles creates an EnrichOption that sets roles
func WithRoles(roles []string) EnrichOption {
	return func(config *EnrichConfig) {
		config.Extensions["roles"] = roles
	}
}

// WithFieldFilter creates an EnrichOption that sets a field filter
func WithFieldFilter(filter func(*Field) bool) EnrichOption {
	return func(config *EnrichConfig) {
		config.FieldFilter = filter
	}
}

// WithEnvironment creates an EnrichOption that sets the environment
func WithEnvironment(env string) EnrichOption {
	return func(config *EnrichConfig) {
		config.Environment = env
	}
}

// WithFeature creates an EnrichOption that enables a feature flag
func WithFeature(name string, enabled bool) EnrichOption {
	return func(config *EnrichConfig) {
		if config.Features == nil {
			config.Features = make(map[string]bool)
		}
		config.Features[name] = enabled
	}
}

// WithSkipValidation creates an EnrichOption that skips schema validation
func WithSkipValidation(skip bool) EnrichOption {
	return func(config *EnrichConfig) {
		config.SkipValidation = skip
	}
}

// WithExtension creates an EnrichOption that adds an arbitrary extension value
func WithExtension(key string, value any) EnrichOption {
	return func(config *EnrichConfig) {
		if config.Extensions == nil {
			config.Extensions = make(map[string]any)
		}
		config.Extensions[key] = value
	}
}

// WithLocale creates an EnrichOption that sets the preferred locale
func WithLocale(locale string) EnrichOption {
	return func(config *EnrichConfig) {
		config.Extensions["locale"] = locale
	}
}

// WithDepthLimit creates an EnrichOption that sets the depth limit for nested enrichment
func WithDepthLimit(limit int) EnrichOption {
	return func(config *EnrichConfig) {
		config.DepthLimit = limit
	}
}

// noopLogger is a no-op logger implementation for when no logger is provided
type noopLogger struct{}

func (l *noopLogger) Warn(ctx context.Context, msg string, fields map[string]any)  {}
func (l *noopLogger) Error(ctx context.Context, msg string, fields map[string]any) {}
func (l *noopLogger) Debug(ctx context.Context, msg string, fields map[string]any) {}
