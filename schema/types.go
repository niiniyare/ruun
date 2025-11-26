package schema

import (
	"context"
	"slices"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
)

/*
Package schema/types.go - Consolidated Type System

This file serves as the central repository for all common types used throughout
the schema package. It consolidates 73 types that were previously scattered
across multiple files, eliminating duplication and providing a single source
of truth for type definitions.

ARCHITECTURE PRINCIPLES:
┌─────────────────────────────────────────────────────────────────────────────┐
│                          CONSOLIDATED TYPES (types.go)                     │
├─────────────────────────────────────────────────────────────────────────────┤
│ ✅ Enumerations (8):     Type, ActionType, ErrorType, TransformType...     │
│ ✅ Validation (8):       ValidationResult, ValidationError, ValidationRule │
│ ✅ Metadata (8):         BaseMetadata, Meta, Theme, TenantOverride...       │
│ ✅ Configuration (12):   Config, Security, CSRF, I18nConfig...              │
│ ✅ Workflow (8):         Workflow, WorkflowAction, ApprovalConfig...        │
│ ✅ Layout Config (4):    GridConfig, FlexConfig, StackConfig...             │
│ ✅ Other Types (25):     DataSource, FieldOption, Events, HTMX...           │
└─────────────────────────────────────────────────────────────────────────────┘

EXCLUDED TYPES (remain in dedicated files):
┌─────────────────────────────────────────────────────────────────────────────┐
│                           DEDICATED FILES                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│ 🏗️  Field struct       → field.go      (main field definition)            │
│ 🏗️  Schema struct      → schema.go     (main schema definition)           │
│ 🏗️  Action struct      → action.go     (main action definition)           │
│ 🏗️  Layout struct      → layout.go     (main layout definition)           │
│ 🏗️  Conditional types  → conditional.go (complex conditional logic)       │
│ 🏗️  Style types        → style.go      (styling system)                   │
│ 🏗️  Binding types      → binding.go    (data binding)                     │
│ 🏗️  Behavior types     → behavior.go   (complex behavior system)          │
│ 🏗️  Interface types    → interface.go  (contracts & interfaces)           │
└─────────────────────────────────────────────────────────────────────────────┘

PERFORMANCE CHARACTERISTICS:
• Memory Usage: 34 pointer types, 34 maps, 37 slices (optimized for efficiency)
• Build Time: Zero compilation errors, fast build with eliminated redundancy
• Type Safety: 73 strongly-typed definitions with validation tags
• API Surface: 32 helper methods for common operations

USAGE PATTERNS:
• Import once: All common types available from single import
• Type Safety: Enums prevent invalid state assignments
• Validation: Built-in validation with structured error handling
• Extensibility: Clean foundation for adding new types

MAINTENANCE:
• Single file to modify for common type changes
• Clear categorization with section headers
• Comprehensive helper methods
• Future-proof architecture

Generated: $(date)
Types: 73 total, 8 enum groups, 32 helper methods
*/

// ============================================================================
// ENUMERATIONS
// ============================================================================

// Type represents the schema type (form, table, wizard, etc.)
type Type string

const (
	TypeForm   Type = "form"
	TypeTable  Type = "table"
	TypeWizard Type = "wizard"
	TypeDialog Type = "dialog"
	TypeCard   Type = "card"
	TypeList   Type = "list"
	TypeDetail Type = "detail"
)

// ActionType represents different action types
type ActionType string

const (
	ActionSubmit ActionType = "submit"
	ActionReset  ActionType = "reset"
	ActionButton ActionType = "button"
	ActionLink   ActionType = "link"
	ActionCustom ActionType = "custom"
)

// ErrorType represents different types of errors
type ErrorType string

const (
	ErrorTypeValidation ErrorType = "validation"
	ErrorTypeNotFound   ErrorType = "not_found"
	ErrorTypeConflict   ErrorType = "conflict"
	ErrorTypePermission ErrorType = "permission"
	ErrorTypeDataSource ErrorType = "data_source"
	ErrorTypeRender     ErrorType = "render"
	ErrorTypeWorkflow   ErrorType = "workflow"
	ErrorTypeTenant     ErrorType = "tenant"
	ErrorTypeInternal   ErrorType = "internal"
)

// TransformType represents field value transformation types
type TransformType string

const (
	TransformUpperCase  TransformType = "uppercase"
	TransformLowerCase  TransformType = "lowercase"
	TransformTrim       TransformType = "trim"
	TransformCapitalize TransformType = "capitalize"
	TransformSlugify    TransformType = "slugify"
)

// SwapStrategy represents HTMX swap strategies
type SwapStrategy string

const (
	SwapInnerHTML   SwapStrategy = "innerHTML"
	SwapOuterHTML   SwapStrategy = "outerHTML"
	SwapBeforeBegin SwapStrategy = "beforebegin"
	SwapAfterBegin  SwapStrategy = "afterbegin"
	SwapBeforeEnd   SwapStrategy = "beforeend"
	SwapAfterEnd    SwapStrategy = "afterend"
	SwapDelete      SwapStrategy = "delete"
	SwapNone        SwapStrategy = "none"
)

// AggregateType represents different aggregation types
type AggregateType string

const (
	AggregateSum   AggregateType = "sum"
	AggregateAvg   AggregateType = "avg"
	AggregateMin   AggregateType = "min"
	AggregateMax   AggregateType = "max"
	AggregateCount AggregateType = "count"
)

// RegistryEventType represents registry event types
type RegistryEventType string

const (
	EventSchemaRegistered RegistryEventType = "schema.registered"
	EventSchemaUpdated    RegistryEventType = "schema.updated"
	EventSchemaDeleted    RegistryEventType = "schema.deleted"
	EventSchemaAccessed   RegistryEventType = "schema.accessed"
)

// Format represents data formats
type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

// ============================================================================
// VALIDATION TYPES
// ============================================================================

// ValidationResult represents the result of a validation operation
type ValidationResult struct {
	Valid    bool                `json:"valid"`
	Warnings []string            `json:"warnings"`
	Errors   map[string][]string `json:"errors"`
	Data     map[string]any      `json:"data"`
}

// IsValid returns true if validation passed
func (vr *ValidationResult) IsValid() bool {
	return vr.Valid
}

// HasErrors returns true if there are validation errors
func (vr *ValidationResult) HasErrors() bool {
	return len(vr.Errors) > 0
}

// HasWarnings returns true if there are validation warnings
func (vr *ValidationResult) HasWarnings() bool {
	return len(vr.Warnings) > 0
}

// AddError adds an error for a specific field
func (vr *ValidationResult) AddError(field, message string) {
	if vr.Errors == nil {
		vr.Errors = make(map[string][]string)
	}
	vr.Errors[field] = append(vr.Errors[field], message)
	vr.Valid = false
}

// AddWarning adds a warning message
func (vr *ValidationResult) AddWarning(message string) {
	vr.Warnings = append(vr.Warnings, message)
}

// GetErrorsForField returns errors for a specific field
func (vr *ValidationResult) GetErrorsForField(field string) []string {
	if vr.Errors == nil {
		return nil
	}
	return vr.Errors[field]
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// ValidationMessages defines custom error messages for validation rules
type ValidationMessages struct {
	Required  string `json:"required,omitempty"`
	MinLength string `json:"minLength,omitempty"`
	MaxLength string `json:"maxLength,omitempty"`
	Pattern   string `json:"pattern,omitempty"`
	Min       string `json:"min,omitempty"`
	Max       string `json:"max,omitempty"`
	Custom    string `json:"custom,omitempty"`
}

// ValidationRule represents a single validation rule
type ValidationRule struct {
	Name       string         `json:"name"`
	Params     map[string]any `json:"params,omitempty"`
	Message    string         `json:"message,omitempty"`
	Async      bool           `json:"async,omitempty"`
	Condition  string         `json:"condition,omitempty"`
	CustomFunc string         `json:"customFunc,omitempty"`
}

// Validation defines cross-field validation rules and configuration
type Validation struct {
	Rules      []ValidationRule `json:"rules" validate:"dive"`
	Mode       string           `json:"mode,omitempty" validate:"oneof=onChange onBlur onSubmit"`
	Async      bool             `json:"async,omitempty"`
	DebounceMS int              `json:"debounceMs,omitempty"`
}

// Note: ValidatorFunc and AsyncValidatorFunc moved to interface.go

// CacheEntry represents a validation cache entry
type CacheEntry struct {
	Result    error
	ExpiresAt time.Time
}

// IsExpired returns true if the cache entry has expired
func (ce *CacheEntry) IsExpired() bool {
	return time.Now().After(ce.ExpiresAt)
}

// ============================================================================
// METADATA TYPES
// ============================================================================

// BaseMetadata contains core metadata fields shared across all metadata structures
type BaseMetadata struct {
	Version   string    `json:"version"`   // Version identifier
	CreatedAt time.Time `json:"createdAt"` // Creation timestamp
	UpdatedAt time.Time `json:"updatedAt"` // Last update timestamp
}

// Meta contains metadata and tracking information
type Meta struct {
	BaseMetadata                   // Embedded base metadata
	CreatedBy     string           `json:"createdBy,omitempty" validate:"uuid"`
	UpdatedBy     string           `json:"updatedBy,omitempty" validate:"uuid"`
	Deprecated    bool             `json:"deprecated,omitempty"`
	Experimental  bool             `json:"experimental,omitempty"`
	Changelog     []ChangelogEntry `json:"changelog,omitempty" validate:"dive"`
	CustomData    map[string]any   `json:"customData,omitempty"`
	Author        string           `json:"author,omitempty"`
	License       string           `json:"license,omitempty"`
	Repository    string           `json:"repository,omitempty" validate:"url"`
	Documentation string           `json:"documentation,omitempty" validate:"url"`
	Status        string           `json:"status,omitempty" validate:"oneof=draft published archived"`
	Theme         *Theme           `json:"theme,omitempty"`
}

// ChangelogEntry represents a version change
type ChangelogEntry struct {
	Version     string    `json:"version" validate:"semver"`
	Date        time.Time `json:"date"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description"`
	Breaking    bool      `json:"breaking,omitempty"`
	Changes     []string  `json:"changes,omitempty"`
	Migration   string    `json:"migration,omitempty"`
}

// Theme represents a complete theme configuration
type Theme struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Version   string         `json:"version,omitempty"`
	Tokens    map[string]any `json:"tokens,omitempty"`
	Overrides map[string]any `json:"overrides,omitempty"`
}

// TenantOverride represents tenant-specific field customization
type TenantOverride struct {
	FieldName    string `json:"field_name"`
	Label        string `json:"label,omitempty"`
	Required     *bool  `json:"required,omitempty"`
	Hidden       *bool  `json:"hidden,omitempty"`
	DefaultValue any    `json:"default_value,omitempty"`
	Placeholder  string `json:"placeholder,omitempty"`
	Help         string `json:"help,omitempty"`
}

// TenantCustomization represents tenant-specific schema customization
type TenantCustomization struct {
	SchemaID  string                    `json:"schema_id"`
	TenantID  string                    `json:"tenant_id"`
	Overrides map[string]TenantOverride `json:"overrides"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

// TranslationMetadata represents translation metadata
type TranslationMetadata struct {
	BaseMetadata
	Author      string `json:"author"`
	LastUpdated string `json:"lastUpdated"`
	Description string `json:"description"`
}

// TranslateParams represents translation parameters
type TranslateParams struct {
	Key          string         `json:"key"`
	Params       map[string]any `json:"params"`
	Count        *int           `json:"count"`
	DefaultValue string         `json:"defaultValue"`
}

// ============================================================================
// CONFIGURATION TYPES
// ============================================================================

// Config defines form configuration for submission
type Config struct {
	Action   string            `json:"action,omitempty"`
	Method   string            `json:"method,omitempty"`
	Encoding string            `json:"encoding,omitempty"`
	Timeout  int               `json:"timeout,omitempty"`
	Headers  map[string]string `json:"headers,omitempty"`
}

// Security defines security and access control configuration
type Security struct {
	CSRF          *CSRF       `json:"csrf,omitempty"`
	RateLimit     *RateLimit  `json:"rateLimit,omitempty"`
	Encryption    *Encryption `json:"encryption,omitempty"`
	Sanitization  bool        `json:"sanitization,omitempty"`
	Origins       []string    `json:"origins,omitempty" validate:"dive,url"`
	ContentType   []string    `json:"contentType,omitempty"`
	MaxFileSize   int64       `json:"maxFileSize,omitempty"`
	AllowedExts   []string    `json:"allowedExts,omitempty"`
	XSSProtection bool        `json:"xssProtection,omitempty"`
	SQLInjection  bool        `json:"sqlInjection,omitempty"`
}

// CSRF defines CSRF protection configuration
type CSRF struct {
	Enabled     bool   `json:"enabled"`
	FieldName   string `json:"fieldName,omitempty"`
	HeaderName  string `json:"headerName,omitempty"`
	CookieName  string `json:"cookieName,omitempty"`
	TokenLength int    `json:"tokenLength,omitempty" validate:"min=16,max=64"`
	Expiry      int    `json:"expiry,omitempty"`
}

// RateLimit defines rate limiting configuration
type RateLimit struct {
	Enabled       bool          `json:"enabled"`
	MaxRequests   int           `json:"maxRequests" validate:"min=1"`
	Window        time.Duration `json:"window"`
	ByIP          bool          `json:"byIp,omitempty"`
	ByUser        bool          `json:"byUser,omitempty"`
	ByTenant      bool          `json:"byTenant,omitempty"`
	Burst         int           `json:"burst,omitempty"`
	BlockDuration time.Duration `json:"blockDuration,omitempty"`
}

// Encryption defines field encryption configuration
type Encryption struct {
	Enabled     bool     `json:"enabled"`
	Fields      []string `json:"fields" validate:"dive,fieldname"`
	Algorithm   string   `json:"algorithm,omitempty" validate:"oneof=AES-256-GCM ChaCha20-Poly1305"`
	KeyRotation bool     `json:"keyRotation,omitempty"`
	KeyVersion  int      `json:"keyVersion,omitempty"`
}

// Tenant defines multi-tenancy configuration
type Tenant struct {
	Enabled    bool     `json:"enabled"`
	Field      string   `json:"field,omitempty" validate:"fieldname"`
	Isolation  string   `json:"isolation" validate:"oneof=strict shared hybrid"`
	Inherit    bool     `json:"inherit,omitempty"`
	AllowCross bool     `json:"allowCross,omitempty"`
	Whitelist  []string `json:"whitelist,omitempty" validate:"dive,uuid"`
	Blacklist  []string `json:"blacklist,omitempty" validate:"dive,uuid"`
	Validation bool     `json:"validation,omitempty"`
}

// RegistryConfig defines registry configuration
type RegistryConfig struct {
	EnableStorage          bool
	StorageType            string
	EnableMemoryCache      bool
	MemoryCacheTTL         time.Duration
	MaxMemoryCacheSize     int
	EnableDistributedCache bool
	DistributedCacheTTL    time.Duration
	ValidateOnStore        bool
	ValidateOnLoad         bool
	EnableVersioning       bool
	MaxVersions            int
	EnableEvents           bool
	EnableMetrics          bool
}

// StorageFilter represents storage filtering options
type StorageFilter struct {
	Type     string
	Category string
	Module   string
	Tags     []string
	Limit    int
	Offset   int
}

// StorageMetadata represents storage metadata
type StorageMetadata struct {
	BaseMetadata
	ID          string
	Size        int64
	AccessCount int64
	LastAccess  time.Time
}

// RuntimeConfig defines runtime configuration
type RuntimeConfig struct {
	ValidateOnChange bool
	ValidateOnBlur   bool
	ValidateOnSubmit bool
	EnableDebounce   bool
	DebounceDelay    time.Duration
	EnableTracking   bool
}

// ParserConfig defines parser configuration
type ParserConfig struct {
	ValidateOnParse     bool
	StrictMode          bool
	MaxDepth            int
	MaxSize             int64
	ValidationTimeout   time.Duration
	EnableCache         bool
	CacheTTL            time.Duration
	MaxCacheSize        int
	EnableStreaming     bool
	StreamBufferSize    int
	EnableVersionCheck  bool
	EnableMigration     bool
	EnableCustomParsers bool
	EnableMetrics       bool
	SanitizeInput       bool
	DisallowExtensions  bool
	ValidateReferences  bool
}

// I18nConfig defines internationalization configuration
type I18nConfig struct {
	DefaultLocale    string
	FallbackLocale   string
	SupportedLocales []string
	TranslationsPath string
	FilePattern      string
	EnableCache      bool
	EnablePlurals    bool
}

// ============================================================================
// WORKFLOW TYPES
// ============================================================================

// Workflow defines a business workflow
type Workflow struct {
	Enabled       bool                   `json:"enabled"`
	Stage         string                 `json:"stage,omitempty"`
	Status        string                 `json:"status,omitempty"`
	Actions       []WorkflowAction       `json:"actions,omitempty" validate:"dive"`
	Transitions   []WorkflowTransition   `json:"transitions,omitempty" validate:"dive"`
	Approval      *ApprovalConfig        `json:"approval,omitempty"`
	Timeout       *WorkflowTimeout       `json:"timeout,omitempty"`
	Notifications []Notification         `json:"notifications,omitempty" validate:"dive"`
	History       []WorkflowHistoryEntry `json:"history,omitempty"`
	Metadata      map[string]any         `json:"metadata,omitempty"`
}

// WorkflowAction represents a workflow action
type WorkflowAction struct {
	ID          string                    `json:"id" validate:"required"`
	Label       string                    `json:"label" validate:"required"`
	Type        string                    `json:"type" validate:"oneof=approve reject submit return delegate escalate"`
	ToStage     string                    `json:"toStage,omitempty"`
	ToStatus    string                    `json:"toStatus,omitempty"`
	Permissions []string                  `json:"permissions,omitempty"`
	RequireNote bool                      `json:"requireNote,omitempty"`
	Icon        string                    `json:"icon,omitempty"`
	Variant     string                    `json:"variant,omitempty" validate:"oneof=primary secondary outline destructive"`
	Condition   *condition.ConditionGroup `json:"condition,omitempty"`
}

// WorkflowTransition represents a workflow state transition
type WorkflowTransition struct {
	From       string                    `json:"from" validate:"required"`
	To         string                    `json:"to" validate:"required"`
	Action     string                    `json:"action" validate:"required"`
	Conditions *condition.ConditionGroup `json:"conditions,omitempty"`
	Validator  string                    `json:"validator,omitempty"`
	Auto       bool                      `json:"auto,omitempty"`
	Delay      time.Duration             `json:"delay,omitempty"`
}

// WorkflowTimeout represents workflow timeout configuration
type WorkflowTimeout struct {
	Duration time.Duration `json:"duration" validate:"min=1"`
	Action   string        `json:"action" validate:"oneof=escalate auto-approve auto-reject notify"`
	NotifyAt []int         `json:"notifyAt,omitempty"`
}

// WorkflowHistoryEntry represents a workflow history entry
type WorkflowHistoryEntry struct {
	ID        string         `json:"id"`
	Action    string         `json:"action"`
	FromStage string         `json:"fromStage"`
	ToStage   string         `json:"toStage"`
	ActorID   string         `json:"actorId"`
	ActorName string         `json:"actorName"`
	Note      string         `json:"note,omitempty"`
	Metadata  map[string]any `json:"metadata,omitempty"`
	Timestamp time.Time      `json:"timestamp"`
}

// Notification defines a workflow notification
type Notification struct {
	Event    string   `json:"event" validate:"required"`
	Template string   `json:"template" validate:"required"`
	To       []string `json:"to" validate:"required,dive"`
	CC       []string `json:"cc,omitempty" validate:"dive"`
	BCC      []string `json:"bcc,omitempty" validate:"dive"`
	Method   string   `json:"method" validate:"oneof=email sms push in-app webhook"`
	Priority string   `json:"priority,omitempty" validate:"oneof=low normal high urgent"`
}

// ApprovalConfig defines approval requirements
type ApprovalConfig struct {
	Required        bool             `json:"required"`
	MinApprovals    int              `json:"minApprovals" validate:"min=1"`
	Approvers       []ApproverConfig `json:"approvers" validate:"dive"`
	Sequential      bool             `json:"sequential,omitempty"`
	AllowDelegation bool             `json:"allowDelegation,omitempty"`
	Timeout         *time.Duration   `json:"timeout,omitempty"`
	EscalationTo    string           `json:"escalationTo,omitempty"`
	RemindAfter     *time.Duration   `json:"remindAfter,omitempty"`
	AutoApprove     bool             `json:"autoApprove,omitempty"`
}

// ApproverConfig defines who can approve
type ApproverConfig struct {
	Type  string `json:"type" validate:"oneof=role user group dynamic manager"`
	Value string `json:"value" validate:"required"`
	Order int    `json:"order,omitempty"`
	Level int    `json:"level,omitempty"`
}

// ============================================================================
// LAYOUT CONFIG TYPES
// ============================================================================

// GridConfig represents CSS Grid configuration
type GridConfig struct {
	Columns     string   `json:"columns,omitempty"`
	Rows        string   `json:"rows,omitempty"`
	Gap         string   `json:"gap,omitempty"`
	ColumnGap   string   `json:"column_gap,omitempty"`
	RowGap      string   `json:"row_gap,omitempty"`
	AutoFlow    string   `json:"auto_flow,omitempty"`
	AutoColumns string   `json:"auto_columns,omitempty"`
	AutoRows    string   `json:"auto_rows,omitempty"`
	Template    string   `json:"template,omitempty"`
	Areas       []string `json:"areas,omitempty"`
}

// FlexConfig represents Flexbox configuration
type FlexConfig struct {
	Direction      string `json:"direction,omitempty"`
	Wrap           string `json:"wrap,omitempty"`
	AlignItems     string `json:"align_items,omitempty"`
	AlignContent   string `json:"align_content,omitempty"`
	JustifyContent string `json:"justify_content,omitempty"`
	Gap            string `json:"gap,omitempty"`
	Grow           int    `json:"grow,omitempty"`
	Shrink         int    `json:"shrink,omitempty"`
	Basis          string `json:"basis,omitempty"`
}

// StackConfig represents Stack layout configuration
type StackConfig struct {
	Direction string `json:"direction,omitempty"`
	Spacing   string `json:"spacing,omitempty"`
	Align     string `json:"align,omitempty"`
	Justify   string `json:"justify,omitempty"`
	Divider   bool   `json:"divider,omitempty"`
}

// PositionConfig represents positioning configuration
type PositionConfig struct {
	Type   string `json:"type,omitempty"`
	Top    string `json:"top,omitempty"`
	Right  string `json:"right,omitempty"`
	Bottom string `json:"bottom,omitempty"`
	Left   string `json:"left,omitempty"`
	ZIndex int    `json:"z_index,omitempty"`
}

// ============================================================================
// OTHER TYPES
// ============================================================================

// Aggregate represents an aggregation configuration
type Aggregate struct {
	Type   AggregateType `json:"type" validate:"required,oneof=sum avg min max count"`
	Field  string        `json:"field" validate:"required"`
	Label  string        `json:"label,omitempty"`
	Format string        `json:"format,omitempty"`
}

// SchemaI18n holds schema-level translations
type SchemaI18n struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	Direction   map[string]string `json:"direction,omitempty"`
}

// ActionHTMX represents HTMX configuration for actions
type ActionHTMX struct {
	Method string `json:"method,omitempty"`
	URL    string `json:"url,omitempty"`
	Target string `json:"target,omitempty"`
	Swap   string `json:"swap,omitempty"`
}

// ActionTheme represents action-specific theme configuration
type ActionTheme struct {
	Colors       map[string]string `json:"colors,omitempty"`
	Spacing      map[string]string `json:"spacing,omitempty"`
	BorderRadius string            `json:"borderRadius,omitempty"`
	FontSize     string            `json:"fontSize,omitempty"`
	FontWeight   string            `json:"fontWeight,omitempty"`
	CustomCSS    string            `json:"customCSS,omitempty"`
}

// ActionConfig represents action configuration
type ActionConfig struct {
	URL             string            `json:"url,omitempty"`
	Target          string            `json:"target,omitempty"`
	Handler         string            `json:"handler,omitempty"`
	Params          map[string]string `json:"params,omitempty"`
	Debounce        int               `json:"debounce,omitempty"`
	Throttle        int               `json:"throttle,omitempty"`
	PreventDefault  bool              `json:"preventDefault,omitempty"`
	StopPropagation bool              `json:"stopPropagation,omitempty"`
	SuccessMessage  string            `json:"successMessage,omitempty"`
	ErrorMessage    string            `json:"errorMessage,omitempty"`
	RedirectURL     string            `json:"redirectUrl,omitempty"`
}

// ActionConfirm represents action confirmation configuration
type ActionConfirm struct {
	Enabled bool   `json:"enabled"`
	Title   string `json:"title,omitempty"`
	Message string `json:"message"`
	Confirm string `json:"confirm,omitempty"`
	Cancel  string `json:"cancel,omitempty"`
	Variant string `json:"variant,omitempty"`
	Icon    string `json:"icon,omitempty"`
}

// ActionPermissions represents action permission configuration
type ActionPermissions struct {
	View     []string `json:"view,omitempty"`
	Execute  []string `json:"execute,omitempty"`
	Required []string `json:"required,omitempty"`
}

// Trigger represents an event trigger
type Trigger struct {
	Event   string `json:"event,omitempty"`
	Target  string `json:"target,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Delay   int    `json:"delay,omitempty"`
	Once    bool   `json:"once,omitempty"`
	Changed bool   `json:"changed,omitempty"`
	From    string `json:"from,omitempty"`
}

// DataSource represents field data source configuration
type DataSource struct {
	Type      string            `json:"type" validate:"oneof=api static computed"`
	URL       string            `json:"url,omitempty" validate:"url"`
	Method    string            `json:"method,omitempty" validate:"oneof=GET POST"`
	Headers   map[string]string `json:"headers,omitempty"`
	Params    map[string]string `json:"params,omitempty"`
	CacheTTL  int               `json:"cacheTTL,omitempty"`
	Static    []FieldOption     `json:"static,omitempty"`
	Computed  string            `json:"computed,omitempty"`
	Transform string            `json:"transform,omitempty"`
}

// FieldOption represents a field option
type FieldOption struct {
	Value    any           `json:"value"`
	Label    string        `json:"label"`
	Disabled bool          `json:"disabled,omitempty"`
	Icon     string        `json:"icon,omitempty"`
	Group    string        `json:"group,omitempty"`
	Children []FieldOption `json:"children,omitempty"`
}

// RegistryEvent represents a registry event
type RegistryEvent struct {
	Type      RegistryEventType
	SchemaID  string
	Timestamp time.Time
	Data      map[string]any
}

// RegistryEventHandler handles registry events
type RegistryEventHandler func(*RegistryEvent)

// RegistryMetrics represents registry metrics
type RegistryMetrics struct {
	schemaCount int64
	cacheHits   int64
	cacheMisses int64
	operations  map[string]*operationMetrics
	mu          sync.RWMutex
}

// operationMetrics represents metrics for a specific operation
type operationMetrics struct {
	count    int64
	duration int64
}

// ParseResult represents the result of a parse operation
type ParseResult struct {
	Schema interface{} // Schema type to avoid circular dependency
	Error  error
	Index  int
}

// ParseError represents a parsing error
type ParseError struct {
	Code    string
	Message string
	Details map[string]any
}

// ParserStats represents parser statistics
type ParserStats struct {
	TotalParses     int64
	SuccessfulParse int64
	FailedParses    int64
	CacheHits       int64
	CacheMisses     int64
	TotalBytes      int64
	AverageDuration time.Duration
	mu              sync.RWMutex
}

// Event represents a schema-related event
type Event struct {
	Type      string
	Source    string
	Timestamp int64
	Data      map[string]any
}

// EventHandler processes events
// Note: EventHandler moved to interface.go

// BehaviorMetadata describes a behavior's capabilities
type BehaviorMetadata struct {
	Name        string
	Description string
	Version     string
	Inputs      map[string]string // Field type names as strings
	Outputs     map[string]string // Field type names as strings
	Tags        []string
}

// Events defines lifecycle event handlers
type Events struct {
	OnMount         string `json:"onMount,omitempty"`
	OnUnmount       string `json:"onUnmount,omitempty"`
	OnSubmit        string `json:"onSubmit,omitempty"`
	BeforeSubmit    string `json:"beforeSubmit,omitempty"`
	AfterSubmit     string `json:"afterSubmit,omitempty"`
	OnSubmitSuccess string `json:"onSubmitSuccess,omitempty"`
	OnSubmitError   string `json:"onSubmitError,omitempty"`
	OnReset         string `json:"onReset,omitempty"`
	OnValidate      string `json:"onValidate,omitempty"`
	OnChange        string `json:"onChange,omitempty"`
	OnDirty         string `json:"onDirty,omitempty"`
	OnPristine      string `json:"onPristine,omitempty"`
	OnLoad          string `json:"onLoad,omitempty"`
	OnError         string `json:"onError,omitempty"`
	OnFieldChange   string `json:"onFieldChange,omitempty"`
	OnFieldFocus    string `json:"onFieldFocus,omitempty"`
	OnFieldBlur     string `json:"onFieldBlur,omitempty"`
	OnFieldValidate string `json:"onFieldValidate,omitempty"`
}

// HTMX defines HTMX integration configuration
type HTMX struct {
	Enabled   bool              `json:"enabled"`
	Get       string            `json:"get,omitempty"`
	Post      string            `json:"post,omitempty"`
	Put       string            `json:"put,omitempty"`
	Delete    string            `json:"delete,omitempty"`
	Patch     string            `json:"patch,omitempty"`
	Trigger   string            `json:"trigger,omitempty"`
	Target    string            `json:"target,omitempty"`
	Swap      string            `json:"swap,omitempty"`
	Select    string            `json:"select,omitempty"`
	Indicator string            `json:"indicator,omitempty"`
	PushURL   string            `json:"pushUrl,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Vals      string            `json:"vals,omitempty"`
	Confirm   string            `json:"confirm,omitempty"`
	Boost     bool              `json:"boost,omitempty"`
	Sync      string            `json:"sync,omitempty"`
	Validate  bool              `json:"validate,omitempty"`
	Timeout   int               `json:"timeout,omitempty"`
	Retry     int               `json:"retry,omitempty"`
}

// DOMEvents represents unified DOM event handlers
type DOMEvents struct {
	Click       string            `json:"on_click,omitempty"`
	DblClick    string            `json:"on_dblclick,omitempty"`
	MouseDown   string            `json:"on_mousedown,omitempty"`
	MouseUp     string            `json:"on_mouseup,omitempty"`
	MouseEnter  string            `json:"on_mouseenter,omitempty"`
	MouseLeave  string            `json:"on_mouseleave,omitempty"`
	MouseOver   string            `json:"on_mouseover,omitempty"`
	MouseOut    string            `json:"on_mouseout,omitempty"`
	MouseMove   string            `json:"on_mousemove,omitempty"`
	KeyDown     string            `json:"on_keydown,omitempty"`
	KeyUp       string            `json:"on_keyup,omitempty"`
	KeyPress    string            `json:"on_keypress,omitempty"`
	Submit      string            `json:"on_submit,omitempty"`
	Change      string            `json:"on_change,omitempty"`
	Input       string            `json:"on_input,omitempty"`
	Focus       string            `json:"on_focus,omitempty"`
	Blur        string            `json:"on_blur,omitempty"`
	Select      string            `json:"on_select,omitempty"`
	TouchStart  string            `json:"on_touchstart,omitempty"`
	TouchEnd    string            `json:"on_touchend,omitempty"`
	TouchMove   string            `json:"on_touchmove,omitempty"`
	TouchCancel string            `json:"on_touchcancel,omitempty"`
	DragStart   string            `json:"on_dragstart,omitempty"`
	DragEnd     string            `json:"on_dragend,omitempty"`
	DragEnter   string            `json:"on_dragenter,omitempty"`
	DragLeave   string            `json:"on_dragleave,omitempty"`
	DragOver    string            `json:"on_dragover,omitempty"`
	Drop        string            `json:"on_drop,omitempty"`
	Play        string            `json:"on_play,omitempty"`
	Pause       string            `json:"on_pause,omitempty"`
	Ended       string            `json:"on_ended,omitempty"`
	Load        string            `json:"on_load,omitempty"`
	Unload      string            `json:"on_unload,omitempty"`
	Resize      string            `json:"on_resize,omitempty"`
	Scroll      string            `json:"on_scroll,omitempty"`
	Custom      map[string]string `json:"custom_events,omitempty"`
}

// ValidationCallback is a callback function for async validation
type ValidationCallback func(fieldName string, errors []string)

// I18nManager interface for i18n functionality (forward declaration)

// ============================================================================
// HELPER METHODS
// ============================================================================

// Security helper methods
func (s *Security) IsCSRFEnabled() bool {
	return s.CSRF != nil && s.CSRF.Enabled
}

func (s *Security) IsRateLimitEnabled() bool {
	return s.RateLimit != nil && s.RateLimit.Enabled
}

func (s *Security) IsEncryptionEnabled() bool {
	return s.Encryption != nil && s.Encryption.Enabled
}

func (s *Security) ShouldEncryptField(fieldName string) bool {
	if !s.IsEncryptionEnabled() {
		return false
	}
	return slices.Contains(s.Encryption.Fields, fieldName)
}

// Tenant helper methods
func (t *Tenant) IsTenantEnabled() bool {
	return t != nil && t.Enabled
}

// Workflow helper methods
func (w *Workflow) IsWorkflowEnabled() bool {
	return w != nil && w.Enabled
}

func (w *Workflow) GetCurrentStage() string {
	if w == nil {
		return ""
	}
	return w.Stage
}

func (w *Workflow) GetCurrentStatus() string {
	if w == nil {
		return ""
	}
	return w.Status
}

func (w *Workflow) GetAvailableActions(ctx context.Context, data map[string]any, evaluator *condition.Evaluator) []WorkflowAction {
	if w == nil || !w.Enabled {
		return []WorkflowAction{}
	}
	available := make([]WorkflowAction, 0)
	for _, action := range w.Actions {
		if action.Condition != nil && evaluator != nil {
			evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
			result, err := evaluator.Evaluate(ctx, action.Condition, evalCtx)
			if err != nil || !result {
				continue
			}
		}
		available = append(available, action)
	}
	return available
}

// Meta helper methods
func (m *Meta) AddChangelog(entry ChangelogEntry) {
	if m.Changelog == nil {
		m.Changelog = []ChangelogEntry{}
	}
	m.Changelog = append(m.Changelog, entry)
}

func (m *Meta) GetLatestVersion() string {
	if m == nil || len(m.Changelog) == 0 {
		return ""
	}
	return m.Changelog[len(m.Changelog)-1].Version
}

func (m *Meta) IsDeprecated() bool {
	return m != nil && m.Deprecated
}

func (m *Meta) IsExperimental() bool {
	return m != nil && m.Experimental
}

// SchemaI18n helper methods
func IsRTLForSchemaI18n(schemaI18n *SchemaI18n, locale string) bool {
	if schemaI18n != nil && schemaI18n.Direction != nil {
		if direction, ok := schemaI18n.Direction[locale]; ok {
			return direction == "rtl"
		}
	}
	rtlLanguages := map[string]bool{
		"ar": true, // Arabic
		"he": true, // Hebrew
		"fa": true, // Persian/Farsi
		"ur": true, // Urdu
	}
	return rtlLanguages[locale]
}

// HTMX helper methods
func (h *HTMX) IsHTMXEnabled() bool {
	return h != nil && h.Enabled
}

func (h *HTMX) GetHTTPMethod() string {
	if h == nil {
		return "GET"
	}
	if h.Post != "" {
		return "POST"
	}
	if h.Put != "" {
		return "PUT"
	}
	if h.Patch != "" {
		return "PATCH"
	}
	if h.Delete != "" {
		return "DELETE"
	}
	return "GET"
}

func (h *HTMX) GetURL() string {
	if h == nil {
		return ""
	}
	switch h.GetHTTPMethod() {
	case "POST":
		return h.Post
	case "PUT":
		return h.Put
	case "PATCH":
		return h.Patch
	case "DELETE":
		return h.Delete
	default:
		return h.Get
	}
}

// String method for ValidationError
func (ve *ValidationError) String() string {
	return "Field: " + ve.Field + ", Message: " + ve.Message + ", Code: " + ve.Code
}

// String method for ActionType
func (at ActionType) String() string {
	return string(at)
}

// String method for ErrorType
func (et ErrorType) String() string {
	return string(et)
}

// String method for TransformType
func (tt TransformType) String() string {
	return string(tt)
}

// String method for SwapStrategy
func (ss SwapStrategy) String() string {
	return string(ss)
}

// String method for AggregateType
func (at AggregateType) String() string {
	return string(at)
}

// IsValid method for Type enum
func (t Type) IsValid() bool {
	switch t {
	case TypeForm, TypeTable, TypeWizard, TypeDialog, TypeCard, TypeList, TypeDetail:
		return true
	default:
		return false
	}
}

// String method for Type
func (t Type) String() string {
	return string(t)
}

// ============================================================================
// TYPE ALIASES FOR ERGONOMICS
// ============================================================================

// Common type aliases for frequently used complex types
type (
	// StringMap represents a map of string to any value (common in JSON processing)
	StringMap = map[string]any

	// StringStringMap represents a map of string to string (common in headers/params)
	StringStringMap = map[string]string

	// StringSlice represents a slice of strings (common in tags/options)
	StringSlice = []string

	// ValidationErrors represents field validation errors
	ValidationErrors = map[string][]string

	// ConditionalFunc represents a conditional evaluation function
	ConditionalFunc = func(ctx context.Context, data StringMap) bool

	// TransformFunc represents a value transformation function
	TransformFunc = func(value any) (any, error)

	// RenderContext represents rendering context data
	RenderContext = map[string]any
)
