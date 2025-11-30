package schema

import (
	"context"
	"time"
)

// ==============================================================================
// CONSOLIDATED SCHEMA PACKAGE INTERFACES
// ==============================================================================
//
// This file consolidates ALL interfaces for the schema package to eliminate
// duplicates and provide a central location for all interface definitions.
//
// CONSOLIDATION SUMMARY:
// =====================
//
// VALIDATION (consolidated ~6 interfaces into 1-2):
// - Validator interface: ValidateSchema, ValidateData, ValidateField, AddRule, GetRule
// - ValidatorFunc, AsyncValidatorFunc types: moved from types.go
//
// RENDERING (consolidated 2 into 1):
// - Renderer interface: Render, RenderField, Format
//
// REGISTRY & STORAGE (3 interfaces):
// - SchemaRegistry interface: Register, Get, List, Delete, Exists
// - StorageBackend interface: low-level byte storage operations
// - CacheBackend interface: caching with TTL support
//
// EVENTS (2 interfaces):
// - EventHandler type: function type for event handling
// - EventDispatcher interface: Dispatch, Subscribe, Unsubscribe, GetHandlers
//
// STATE MANAGEMENT (1 interface):
// - StateManager interface: SaveState, LoadState, DeleteState, ListStates
//
// ENRICHMENT (3 interfaces):
// - User interface: GetID, GetTenantID, GetPermissions, HasPermission, etc.
// - TenantProvider interface: GetCustomization
// - Enricher interface: Enrich, EnrichField, EnrichWithLocale
//
// BEHAVIOR (3 interfaces):
// - BehaviorAdapter interface: Execute, Validate, GetMetadata
// - ConditionalEngine interface: Evaluate
// - BusinessRulesEngine interface: ApplyRules
//
// DATA ACCESS (1 interface):
// - Database interface: Exists (for uniqueness checks)
//
// PARSING (1 interface):
// - ParserPlugin interface: Name, CanParse, Parse
//
// LOGGING (1 interface):
// - Logger interface: Warn, Error, Debug
//
// ACCESSORS (3 interfaces for mocking/testing):
// - SchemaAccessor interface: GetID, GetType, GetFields, etc.
// - FieldAccessor interface: GetName, GetType, GetValidation, GetRuntime, etc.
// - RepeatableFieldInfo interface: GetName, GetMinItems, GetTemplate, etc.
//
// ERRORS (1 interface):
// - SchemaError interface: Code, Type, Field, Details, WithField, WithDetail
//
// REMOVED DUPLICATES FROM:
// - doc.go: SchemaValidator, SchemaRenderer, SchemaRegistry → consolidated
// - validator.go: FieldInterface, SchemaInterface, etc. → consolidated into FieldAccessor/SchemaAccessor
// - enricher.go: User, TenantProvider, Logger, Enricher → moved here
// - runtime.go: ConditionalEngine → moved here
// - registry.go: StorageBackend, CacheBackend → moved here (with correct signatures)
// - repeatable.go: RepeatableFieldInfo → moved here
// - errors.go: SchemaError → moved here
// - types.go: ValidatorFunc, AsyncValidatorFunc, EventHandler → moved here
//
// ==============================================================================

// ==============================================================================
// VALIDATION INTERFACES
// ==============================================================================

// Validator provides schema and field validation capabilities
type Validator interface {
	// ValidateSchema validates an entire schema with data
	ValidateSchema(ctx context.Context, schema Schema, data map[string]any) *ValidationResult
	// ValidateField validates a single field with its value
	ValidateField(ctx context.Context, field Field, value any) *ValidationResult
	// ValidateData validates form data against a schema
	ValidateData(ctx context.Context, schema *Schema, data map[string]any) error
	// AddRule adds a custom validation rule
	AddRule(name string, rule ValidatorFunc) error
	// GetRule retrieves a validation rule by name
	GetRule(name string) (ValidatorFunc, bool)
}

// ValidatorFunc represents a validation function
type ValidatorFunc func(ctx context.Context, value any, params map[string]any) error

// AsyncValidatorFunc represents an async validation function
type AsyncValidatorFunc func(ctx context.Context, value any, params map[string]any) error

// ==============================================================================
// RENDERING INTERFACE
// ==============================================================================

// Renderer handles schema rendering to various formats
type Renderer interface {
	// Render renders a complete schema
	Render(ctx context.Context, schema *Schema, data map[string]any) (string, error)
	// RenderField renders a single field
	RenderField(ctx context.Context, field *Field, value any) (string, error)
	// Format formats a field value for display
	Format(ctx context.Context, field *Field, value any) (string, error)
}

// ==============================================================================
// REGISTRY & STORAGE INTERFACES
// ==============================================================================

// SchemaRegistry manages schema registration and retrieval
type SchemaRegistry interface {
	// Register stores a schema
	Register(ctx context.Context, schema *Schema) error
	// Get retrieves a schema by ID
	Get(ctx context.Context, id string) (*Schema, error)
	// List returns schemas matching filter criteria
	List(ctx context.Context, filter ListFilter) ([]*Schema, error)
	// Delete removes a schema
	Delete(ctx context.Context, id string) error
	// Exists checks if a schema exists
	Exists(ctx context.Context, id string) (bool, error)
}

// StorageBackend defines interface for persistent schema storage (low-level byte storage)
type StorageBackend interface {
	// Basic CRUD operations
	Get(ctx context.Context, id string) ([]byte, error)
	Set(ctx context.Context, id string, data []byte) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)
	// Listing operations
	List(ctx context.Context, filter *StorageFilter) ([]string, error)
	// Versioning operations
	GetVersion(ctx context.Context, id, version string) ([]byte, error)
	ListVersions(ctx context.Context, id string) ([]string, error)
	// Batch operations
	GetBatch(ctx context.Context, ids []string) (map[string][]byte, error)
	SetBatch(ctx context.Context, items map[string][]byte) error
	// Metadata operations
	GetMetadata(ctx context.Context, id string) (*StorageMetadata, error)
	// Health check
	Health(ctx context.Context) error
}

// CacheBackend defines interface for schema caching (low-level byte caching)
type CacheBackend interface {
	// Get retrieves cached data
	Get(ctx context.Context, key string) ([]byte, error)
	// Set stores data in cache with TTL
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	// Delete removes data from cache
	Delete(ctx context.Context, key string) error
	// Clear removes all cached data
	Clear(ctx context.Context) error
	// Health checks cache backend health
	Health(ctx context.Context) error
}

// ==============================================================================
// EVENTS INTERFACE
// ==============================================================================

// EventHandler represents an event handler function
type EventHandler func(ctx context.Context, event Event) error

// EventDispatcher manages event handling and notifications
type EventDispatcher interface {
	// Dispatch sends an event to all registered handlers
	Dispatch(ctx context.Context, event *Event) error
	// Subscribe registers an event handler for specific event types
	Subscribe(eventType string, handler EventHandler) error
	// Unsubscribe removes an event handler
	Unsubscribe(eventType string, handler EventHandler) error
	// GetHandlers returns all handlers for an event type
	GetHandlers(eventType string) []EventHandler
}

// ==============================================================================
// STATE MANAGEMENT INTERFACE
// ==============================================================================

// StateManager handles persistent state storage for forms/schemas
type StateManager interface {
	// SaveState persists form state
	SaveState(ctx context.Context, sessionID, schemaID string, state map[string]any) error
	// LoadState retrieves saved form state
	LoadState(ctx context.Context, sessionID, schemaID string) (map[string]any, error)
	// DeleteState removes saved state
	DeleteState(ctx context.Context, sessionID, schemaID string) error
	// ListStates returns all saved states for a session
	ListStates(ctx context.Context, sessionID string) ([]string, error)
}

// ==============================================================================
// ENRICHMENT INTERFACES
// ==============================================================================

// User represents a user in the system for enrichment purposes
type User interface {
	// GetID returns the user ID
	GetID() string
	// GetTenantID returns the tenant ID
	GetTenantID() string
	// GetPermissions returns user permissions
	GetPermissions() []string
	// GetRoles returns user roles
	GetRoles() []string
	// HasPermission checks if user has specific permission
	HasPermission(permission string) bool
	// HasRole checks if user has specific role
	HasRole(role string) bool
	// GetPreferredLocale returns user's preferred locale
	GetPreferredLocale() string
}

// TenantProvider retrieves tenant-specific customizations
type TenantProvider interface {
	// GetCustomization retrieves tenant customizations for a schema
	GetCustomization(ctx context.Context, schemaID, tenantID string) (*TenantCustomization, error)
}

// Enricher enriches schemas with runtime data and user context
type Enricher interface {
	// Enrich enriches a schema with runtime data based on user context
	Enrich(ctx context.Context, schema *Schema, user User) (*Schema, error)
	// EnrichField enriches a single field with runtime data
	EnrichField(ctx context.Context, field *Field, user User, data map[string]any) error
	// EnrichWithLocale enriches a schema with localization support
	EnrichWithLocale(ctx context.Context, schema *Schema, user User, locale string) (*Schema, error)
	// SetTenantProvider sets the tenant customization provider
	SetTenantProvider(provider TenantProvider)
}

// ==============================================================================
// BEHAVIOR INTERFACES
// ==============================================================================

// BehaviorResult contains the result of behavior execution
type BehaviorResult struct {
	Success bool           `json:"success"`
	Data    map[string]any `json:"data,omitempty"`
	Error   string         `json:"error,omitempty"`
	Status  int            `json:"status,omitempty"`
}

// BehaviorAdapter executes custom behaviors and business logic
type BehaviorAdapter interface {
	// Execute runs the behavior with given context and data
	Execute(ctx context.Context, behavior *Behavior, data map[string]any) (*BehaviorResult, error)
	// Validate validates behavior configuration
	Validate(ctx context.Context, behavior *Behavior) error
	// GetMetadata returns behavior metadata
	GetMetadata() *BehaviorMetadata
}

// ConditionalEngine evaluates conditional expressions for dynamic behavior
type ConditionalEngine interface {
	// Evaluate applies conditional logic to schema based on current data
	Evaluate(ctx context.Context, schema *Schema, data map[string]any) error
}

// BusinessRulesEngine applies business rules validation
type BusinessRulesEngine interface {
	// ApplyRules applies business rules to schema and data
	ApplyRules(ctx context.Context, schema *Schema, data map[string]any) (*Schema, error)
}

// ==============================================================================
// DATA INTERFACE
// ==============================================================================

// Database provides database operations for validation (uniqueness checks, etc.)
type Database interface {
	// Exists checks if a value exists in the specified table/column
	Exists(ctx context.Context, table, column string, value any) (bool, error)
}

// ==============================================================================
// PARSER INTERFACE
// ==============================================================================

// ParserPlugin enables parsing schemas from various formats
type ParserPlugin interface {
	// Name returns the parser name
	Name() string
	// CanParse checks if this plugin can parse the data
	CanParse(data []byte) bool
	// Parse parses content into a schema
	Parse(ctx context.Context, data []byte) (*Schema, error)
}

// ==============================================================================
// LOGGING INTERFACE
// ==============================================================================

// Logger provides structured logging for schema operations
type Logger interface {
	// Warn logs a warning message with structured fields
	Warn(ctx context.Context, msg string, fields map[string]any)
	// Error logs an error message with structured fields
	Error(ctx context.Context, msg string, fields map[string]any)
	// Debug logs a debug message with structured fields
	Debug(ctx context.Context, msg string, fields map[string]any)
}

// ==============================================================================
// ACCESSOR INTERFACES (for mocking/testing)
// ==============================================================================

// SchemaAccessor provides access to schema properties (useful for mocking)
type SchemaAccessor interface {
	// GetID returns schema ID
	GetID() string
	// GetType returns schema type
	GetType() string
	// GetTitle returns schema title
	GetTitle() string
	// GetVersion returns schema version
	GetVersion() string
	// GetFields returns all fields
	GetFields() []FieldAccessor
	// GetValidation returns validation config
	GetValidation() any
}

// FieldAccessor provides access to field properties (useful for mocking)
type FieldAccessor interface {
	// GetName returns field name
	GetName() string
	// GetType returns field type
	GetType() FieldType
	// GetRequired returns if field is required
	GetRequired() bool
	// GetValidation returns field validation config
	GetValidation() *FieldValidation
	// GetOptions returns field options
	GetOptions() []FieldOption
	// GetConfig returns field configuration
	GetConfig() map[string]any
	// GetRuntime returns runtime state (nil if not enriched)
	GetRuntime() *FieldRuntime
}

// RepeatableFieldInfo provides information about repeatable field capabilities
type RepeatableFieldInfo interface {
	// GetName returns the field name
	GetName() string
	// GetMinItems returns minimum number of items
	GetMinItems() int
	// GetMaxItems returns maximum number of items
	GetMaxItems() int
	// GetTemplate returns the template field configuration
	GetTemplate() []Field
	// CanAddItem checks if more items can be added
	CanAddItem(currentCount int) bool
	// CanRemoveItem checks if items can be removed
	CanRemoveItem(currentCount int) bool
}

// ==============================================================================
// ERROR INTERFACE
// ==============================================================================

// SchemaError is the base interface for all schema-related errors
type SchemaError interface {
	error
	// Code returns the error code
	Code() string
	// Type returns the error type
	Type() ErrorType
	// Field returns the field name (if field-specific)
	Field() string
	// Details returns additional error details
	Details() map[string]any
	// WithField returns a copy of the error with field information
	WithField(field string) SchemaError
	// WithDetail returns a copy of the error with additional detail
	WithDetail(key string, value any) SchemaError
}

// ==============================================================================
// HELPER TYPES FOR INTERFACES
// ==============================================================================

// ListFilter defines criteria for filtering schema lists
type ListFilter struct {
	// Type filters by schema type
	Type string
	// Tags filters by schema tags
	Tags []string
	// TenantID filters by tenant
	TenantID string
	// Limit limits number of results
	Limit int
	// Offset for pagination
	Offset int
}

// Note: StorageFilter and StorageMetadata are defined in types.go
