package schema

import (
	"context"
	"slices"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
)

// Config defines form configuration for submission
type Config struct {
	Action   string            `json:"action,omitempty"`   // Form action URL
	Method   string            `json:"method,omitempty"`   // HTTP method (GET, POST, etc.)
	Encoding string            `json:"encoding,omitempty"` // Form encoding type
	Timeout  int               `json:"timeout,omitempty"`  // Request timeout in seconds
	Headers  map[string]string `json:"headers,omitempty"`  // Additional headers
}

// Validator interface for validating schemas and fields
type Validator interface {
	ValidateSchema(ctx context.Context, schema *Schema) error
	ValidateField(ctx context.Context, field *Field, value any) error
	ValidateData(ctx context.Context, schema *Schema, data map[string]any) error
}

// ValidationCallback is a callback function for async validation
type ValidationCallback func(fieldName string, errors []string)

// Security defines security and access control configuration
type Security struct {
	CSRF          *CSRF       `json:"csrf,omitempty"`                        // CSRF protection
	RateLimit     *RateLimit  `json:"rateLimit,omitempty"`                   // Rate limiting
	Encryption    *Encryption `json:"encryption,omitempty"`                  // Field encryption
	Sanitization  bool        `json:"sanitization,omitempty"`                // HTML sanitization
	Origins       []string    `json:"origins,omitempty" validate:"dive,url"` // Allowed origins
	ContentType   []string    `json:"contentType,omitempty"`                 // Allowed content types
	MaxFileSize   int64       `json:"maxFileSize,omitempty"`                 // Max upload size in bytes
	AllowedExts   []string    `json:"allowedExts,omitempty"`                 // Allowed file extensions
	XSSProtection bool        `json:"xssProtection,omitempty"`               // XSS protection
	SQLInjection  bool        `json:"sqlInjection,omitempty"`                // SQL injection protection
}

// CSRF defines CSRF protection configuration
type CSRF struct {
	Enabled     bool   `json:"enabled"`                                                     // Enable CSRF protection
	FieldName   string `json:"fieldName,omitempty" example:"_csrf"`                         // Form field name
	HeaderName  string `json:"headerName,omitempty" example:"X-CSRF-Token"`                 // Header name
	CookieName  string `json:"cookieName,omitempty" example:"csrf_token"`                   // Cookie name
	TokenLength int    `json:"tokenLength,omitempty" validate:"min=16,max=64" example:"32"` // Token length
	Expiry      int    `json:"expiry,omitempty"`                                            // Token expiry in seconds
}

// RateLimit defines rate limiting configuration
type RateLimit struct {
	Enabled       bool          `json:"enabled"`                                    // Enable rate limiting
	MaxRequests   int           `json:"maxRequests" validate:"min=1" example:"100"` // Max requests
	Window        time.Duration `json:"window" example:"3600"`                      // Time window (seconds)
	ByIP          bool          `json:"byIp,omitempty"`                             // Rate limit by IP
	ByUser        bool          `json:"byUser,omitempty"`                           // Rate limit by user
	ByTenant      bool          `json:"byTenant,omitempty"`                         // Rate limit by tenant
	Burst         int           `json:"burst,omitempty"`                            // Burst capacity
	BlockDuration time.Duration `json:"blockDuration,omitempty"`                    // Block duration after limit
}

// Encryption defines field encryption configuration
type Encryption struct {
	Enabled     bool     `json:"enabled"`                          // Enable encryption
	Fields      []string `json:"fields" validate:"dive,fieldname"` // Fields to encrypt
	Algorithm   string   `json:"algorithm,omitempty" validate:"oneof=AES-256-GCM ChaCha20-Poly1305" example:"AES-256-GCM"`
	KeyRotation bool     `json:"keyRotation,omitempty"` // Enable key rotation
	KeyVersion  int      `json:"keyVersion,omitempty"`  // Current key version
}

// Tenant defines multi-tenancy configuration
type Tenant struct {
	Enabled    bool     `json:"enabled"`                                                          // Enable tenant isolation
	Field      string   `json:"field,omitempty" validate:"fieldname" example:"tenant_id"`         // Tenant ID field
	Isolation  string   `json:"isolation" validate:"oneof=strict shared hybrid" example:"strict"` // Isolation level
	Inherit    bool     `json:"inherit,omitempty"`                                                // Inherit from parent context
	AllowCross bool     `json:"allowCross,omitempty"`                                             // Allow cross-tenant access
	Whitelist  []string `json:"whitelist,omitempty" validate:"dive,uuid"`                         // Allowed tenant IDs
	Blacklist  []string `json:"blacklist,omitempty" validate:"dive,uuid"`                         // Blocked tenant IDs
	Validation bool     `json:"validation,omitempty"`                                             // Validate tenant access
}

// Notification defines a workflow notification
type Notification struct {
	Event    string   `json:"event" validate:"required" example:"approval_required"` // Trigger event
	Template string   `json:"template" validate:"required"`                          // Notification template
	To       []string `json:"to" validate:"required,dive"`                           // Recipient roles/users
	CC       []string `json:"cc,omitempty" validate:"dive"`                          // CC recipients
	BCC      []string `json:"bcc,omitempty" validate:"dive"`                         // BCC recipients
	Method   string   `json:"method" validate:"oneof=email sms push in-app webhook" example:"email"`
	Priority string   `json:"priority,omitempty" validate:"oneof=low normal high urgent"`
}

// ApprovalConfig defines approval requirements
type ApprovalConfig struct {
	Required        bool             `json:"required"`                      // Approval required
	MinApprovals    int              `json:"minApprovals" validate:"min=1"` // Min approvals needed
	Approvers       []ApproverConfig `json:"approvers" validate:"dive"`     // Approver configuration
	Sequential      bool             `json:"sequential,omitempty"`          // Sequential approval
	AllowDelegation bool             `json:"allowDelegation,omitempty"`     // Allow delegation
	Timeout         *time.Duration   `json:"timeout,omitempty"`             // Approval timeout
	EscalationTo    string           `json:"escalationTo,omitempty"`        // Escalation target
	RemindAfter     *time.Duration   `json:"remindAfter,omitempty"`         // Send reminder after
	AutoApprove     bool             `json:"autoApprove,omitempty"`         // Auto-approve on timeout
}

// ApproverConfig defines who can approve
type ApproverConfig struct {
	Type  string `json:"type" validate:"oneof=role user group dynamic manager" example:"role"`
	Value string `json:"value" validate:"required"` // Role/User/Group ID
	Order int    `json:"order,omitempty"`           // Approval order
	Level int    `json:"level,omitempty"`           // Approval level
}

// Validation defines cross-field validation rules
type Validation struct {
	Rules      []ValidationRule `json:"rules" validate:"dive"`                                                     // Validation rules
	Mode       string           `json:"mode,omitempty" validate:"oneof=onChange onBlur onSubmit" example:"onBlur"` // When to validate
	Async      bool             `json:"async,omitempty"`                                                           // Async validation
	DebounceMS int              `json:"debounceMs,omitempty"`                                                      // Debounce delay
}

// ValidationRule defines a cross-field validation rule
type ValidationRule struct {
	ID        string                    `json:"id" validate:"required"`
	Type      string                    `json:"type" validate:"required" example:"compare"` // Rule type
	Fields    []string                  `json:"fields" validate:"required,dive,fieldname"`  // Fields involved
	Condition *condition.ConditionGroup `json:"condition,omitempty"`                        // When to apply
	Message   string                    `json:"message" validate:"required"`                // Error message
	Validator string                    `json:"validator,omitempty"`                        // Custom validator function
	Severity  string                    `json:"severity,omitempty" validate:"oneof=error warning info" example:"error"`
	Priority  int                       `json:"priority,omitempty"` // Validation priority
}

// Events defines lifecycle event handlers
type Events struct {
	// Form lifecycle
	OnMount   string `json:"onMount,omitempty" validate:"js_function"`   // Form mounted
	OnUnmount string `json:"onUnmount,omitempty" validate:"js_function"` // Form unmounted
	// Submission lifecycle
	OnSubmit        string `json:"onSubmit,omitempty" validate:"js_function"`        // Submit triggered
	BeforeSubmit    string `json:"beforeSubmit,omitempty" validate:"js_function"`    // Before submit
	AfterSubmit     string `json:"afterSubmit,omitempty" validate:"js_function"`     // After submit
	OnSubmitSuccess string `json:"onSubmitSuccess,omitempty" validate:"js_function"` // Submit success
	OnSubmitError   string `json:"onSubmitError,omitempty" validate:"js_function"`   // Submit error
	// Form state
	OnReset    string `json:"onReset,omitempty" validate:"js_function"`    // Form reset
	OnValidate string `json:"onValidate,omitempty" validate:"js_function"` // Validation
	OnChange   string `json:"onChange,omitempty" validate:"js_function"`   // Any field changed
	OnDirty    string `json:"onDirty,omitempty" validate:"js_function"`    // Form becomes dirty
	OnPristine string `json:"onPristine,omitempty" validate:"js_function"` // Form becomes pristine
	// Data events
	OnLoad  string `json:"onLoad,omitempty" validate:"js_function"`  // Data loaded
	OnError string `json:"onError,omitempty" validate:"js_function"` // Error occurred
	// Field events
	OnFieldChange   string `json:"onFieldChange,omitempty" validate:"js_function"`   // Individual field changed
	OnFieldFocus    string `json:"onFieldFocus,omitempty" validate:"js_function"`    // Field focused
	OnFieldBlur     string `json:"onFieldBlur,omitempty" validate:"js_function"`     // Field blurred
	OnFieldValidate string `json:"onFieldValidate,omitempty" validate:"js_function"` // Field validated
}

// HTMX defines HTMX integration for the entire form
type HTMX struct {
	Enabled   bool              `json:"enabled"` // Enable HTMX
	Get       string            `json:"get,omitempty" validate:"url"`
	Post      string            `json:"post,omitempty" validate:"url"`
	Put       string            `json:"put,omitempty" validate:"url"`
	Delete    string            `json:"delete,omitempty" validate:"url"`
	Patch     string            `json:"patch,omitempty" validate:"url"`
	Trigger   string            `json:"trigger,omitempty"`                        // Trigger event
	Target    string            `json:"target,omitempty" validate:"css_selector"` // Update target
	Swap      string            `json:"swap,omitempty" validate:"oneof=innerHTML outerHTML beforebegin afterbegin beforeend afterend delete none"`
	Select    string            `json:"select,omitempty" validate:"css_selector"`    // Response selector
	Indicator string            `json:"indicator,omitempty" validate:"css_selector"` // Loading indicator
	PushURL   string            `json:"pushUrl,omitempty" validate:"url"`            // Push to history
	Headers   map[string]string `json:"headers,omitempty"`                           // Request headers
	Vals      string            `json:"vals,omitempty"`                              // Additional values
	Confirm   string            `json:"confirm,omitempty"`                           // Confirmation prompt
	Boost     bool              `json:"boost,omitempty"`                             // Boost links
	Sync      string            `json:"sync,omitempty"`                              // Sync specification
	Validate  bool              `json:"validate,omitempty"`                          // Validate before request
	Timeout   int               `json:"timeout,omitempty"`                           // Request timeout in ms
	Retry     int               `json:"retry,omitempty"`                             // Number of retries
}

// Alpine defines Alpine.js integration for the entire form
type Alpine struct {
	Enabled     bool   `json:"enabled"`                                     // Enable Alpine.js
	XData       string `json:"xData,omitempty" validate:"js_object"`        // Component data
	XInit       string `json:"xInit,omitempty" validate:"js_function"`      // Initialization
	XShow       string `json:"xShow,omitempty" validate:"js_expression"`    // Show/hide
	XIf         string `json:"xIf,omitempty" validate:"js_expression"`      // Conditional
	XModel      string `json:"xModel,omitempty" validate:"js_variable"`     // Two-way binding
	XBind       string `json:"xBind,omitempty" validate:"js_object"`        // Attribute bindings
	XOn         string `json:"xOn,omitempty" validate:"js_object"`          // Event handlers
	XText       string `json:"xText,omitempty" validate:"js_expression"`    // Text content
	XHTML       string `json:"xHtml,omitempty" validate:"js_expression"`    // HTML content
	XRef        string `json:"xRef,omitempty" validate:"js_variable"`       // Reference
	XCloak      bool   `json:"xCloak,omitempty"`                            // Cloak until ready
	XTransition string `json:"xTransition,omitempty"`                       // Transition effect
	XTeleport   string `json:"xTeleport,omitempty" validate:"css_selector"` // Teleport target
	XFor        string `json:"xFor,omitempty" validate:"js_expression"`     // Loop directive
	XEffect     string `json:"xEffect,omitempty" validate:"js_function"`    // Side effects
}

// BaseMetadata contains core metadata fields shared across all metadata structures
type BaseMetadata struct {
	Version   string    `json:"version"`   // Version identifier
	CreatedAt time.Time `json:"createdAt"` // Creation timestamp
	UpdatedAt time.Time `json:"updatedAt"` // Last update timestamp
}

// Meta contains metadata and tracking information
type Meta struct {
	BaseMetadata                   // Embedded base metadata
	CreatedBy     string           `json:"createdBy,omitempty" validate:"uuid"`                        // Creator user ID
	UpdatedBy     string           `json:"updatedBy,omitempty" validate:"uuid"`                        // Last updater user ID
	Deprecated    bool             `json:"deprecated,omitempty"`                                       // Schema is deprecated
	Experimental  bool             `json:"experimental,omitempty"`                                     // Experimental feature
	Changelog     []ChangelogEntry `json:"changelog,omitempty" validate:"dive"`                        // Version history
	CustomData    map[string]any   `json:"customData,omitempty"`                                       // Custom metadata
	Author        string           `json:"author,omitempty"`                                           // Schema author
	License       string           `json:"license,omitempty"`                                          // License information
	Repository    string           `json:"repository,omitempty" validate:"url"`                        // Source repository
	Documentation string           `json:"documentation,omitempty" validate:"url"`                     // Documentation URL
	Status        string           `json:"status,omitempty" validate:"oneof=draft published archived"` // Schema status
	Theme         *Theme           `json:"theme,omitempty"`                                            // Theme configuration
}

// ChangelogEntry represents a version change
type ChangelogEntry struct {
	Version     string    `json:"version" validate:"semver"`
	Date        time.Time `json:"date"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description"`
	Breaking    bool      `json:"breaking,omitempty"`  // Breaking change
	Changes     []string  `json:"changes,omitempty"`   // List of changes
	Migration   string    `json:"migration,omitempty"` // Migration guide
}

// Theme represents a complete theme configuration
type Theme struct {
	ID        string                 `json:"id"`                  // Theme identifier
	Name      string                 `json:"name"`                // Human-readable name
	Version   string                 `json:"version,omitempty"`   // Theme version
	Tokens    map[string]interface{} `json:"tokens,omitempty"`    // Design tokens
	Overrides map[string]interface{} `json:"overrides,omitempty"` // Theme overrides
}

// ============================================================================
// Modern Framework Integration Types (from types_v2.go)
// ============================================================================

// // Binding represents Alpine.js data binding configuration
//
//	type Binding struct {
//		// Core Alpine directives
//		Data      string   `json:"x_data,omitempty"`
//		Init      string   `json:"x_init,omitempty"`
//		Show      string   `json:"x_show,omitempty"`
//		Text      string   `json:"x_text,omitempty"`
//		HTML      string   `json:"x_html,omitempty"`
//		Model     string   `json:"x_model,omitempty"`
//		ModelOpts []string `json:"x_model_options,omitempty"`
//
//		// Conditional rendering
//		If     string `json:"x_if,omitempty"`
//		For    string `json:"x_for,omitempty"`
//		ForKey string `json:"x_for_key,omitempty"`
//
//		// Attribute binding
//		Bind map[string]string `json:"x_bind,omitempty"`
//
//		// Event handlers
//		On map[string]string `json:"x_on,omitempty"`
//
//		// Transitions
//		Transition TransitionConfig `json:"x_transition,omitempty"`
//
//		// Advanced features
//		Effect    string   `json:"x_effect,omitempty"`
//		Ignore    bool     `json:"x_ignore,omitempty"`
//		Cloak     bool     `json:"x_cloak,omitempty"`
//		Teleport  string   `json:"x_teleport,omitempty"`
//		Intersect string   `json:"x_intersect,omitempty"`
//		Persist   []string `json:"x_persist,omitempty"`
//
//		// Component communication
//		Ref  string            `json:"x_ref,omitempty"`
//		Refs map[string]string `json:"x_refs,omitempty"`
//		Id   string            `json:"x_id,omitempty"`
//	}
//
// TransitionConfig represents Alpine.js transition configuration
type TransitionConfig struct {
	Enter      string `json:"enter,omitempty"`
	EnterStart string `json:"enter_start,omitempty"`
	EnterEnd   string `json:"enter_end,omitempty"`
	Leave      string `json:"leave,omitempty"`
	LeaveStart string `json:"leave_start,omitempty"`
	LeaveEnd   string `json:"leave_end,omitempty"`
	Duration   string `json:"duration,omitempty"`
}

// DOMEvents represents unified DOM event handlers
type DOMEvents struct {
	// Standard DOM events
	Click      string `json:"on_click,omitempty"`
	DblClick   string `json:"on_dblclick,omitempty"`
	MouseDown  string `json:"on_mousedown,omitempty"`
	MouseUp    string `json:"on_mouseup,omitempty"`
	MouseEnter string `json:"on_mouseenter,omitempty"`
	MouseLeave string `json:"on_mouseleave,omitempty"`
	MouseOver  string `json:"on_mouseover,omitempty"`
	MouseOut   string `json:"on_mouseout,omitempty"`
	MouseMove  string `json:"on_mousemove,omitempty"`

	// Keyboard events
	KeyDown  string `json:"on_keydown,omitempty"`
	KeyUp    string `json:"on_keyup,omitempty"`
	KeyPress string `json:"on_keypress,omitempty"`

	// Form events
	Submit string `json:"on_submit,omitempty"`
	Change string `json:"on_change,omitempty"`
	Input  string `json:"on_input,omitempty"`
	Focus  string `json:"on_focus,omitempty"`
	Blur   string `json:"on_blur,omitempty"`
	Select string `json:"on_select,omitempty"`

	// Touch events
	TouchStart  string `json:"on_touchstart,omitempty"`
	TouchEnd    string `json:"on_touchend,omitempty"`
	TouchMove   string `json:"on_touchmove,omitempty"`
	TouchCancel string `json:"on_touchcancel,omitempty"`

	// Drag events
	DragStart string `json:"on_dragstart,omitempty"`
	DragEnd   string `json:"on_dragend,omitempty"`
	DragEnter string `json:"on_dragenter,omitempty"`
	DragLeave string `json:"on_dragleave,omitempty"`
	DragOver  string `json:"on_dragover,omitempty"`
	Drop      string `json:"on_drop,omitempty"`

	// Media events
	Play  string `json:"on_play,omitempty"`
	Pause string `json:"on_pause,omitempty"`
	Ended string `json:"on_ended,omitempty"`

	// Window events
	Load   string `json:"on_load,omitempty"`
	Unload string `json:"on_unload,omitempty"`
	Resize string `json:"on_resize,omitempty"`
	Scroll string `json:"on_scroll,omitempty"`

	// Custom events
	Custom map[string]string `json:"custom_events,omitempty"`
}

// ConditionalV2 represents unified conditional rendering
// Note: Named ConditionalV2 to avoid conflict with existing Conditional types
type Conditional struct {
	// Simple conditions
	If     string `json:"if,omitempty"`
	Unless string `json:"unless,omitempty"`

	// Multiple conditions
	Switch  string            `json:"switch,omitempty"`
	Cases   []ConditionalCase `json:"cases,omitempty"`
	Default interface{}       `json:"default,omitempty"`

	// Visibility conditions
	Show string `json:"show,omitempty"`
	Hide string `json:"hide,omitempty"`

	// State-based conditions
	Loading bool `json:"loading,omitempty"`
	Error   bool `json:"error,omitempty"`
	Success bool `json:"success,omitempty"`

	// Data conditions
	Empty     string `json:"empty,omitempty"`
	NotEmpty  string `json:"not_empty,omitempty"`
	Equals    string `json:"equals,omitempty"`
	NotEquals string `json:"not_equals,omitempty"`

	// Numeric conditions
	GreaterThan string `json:"greater_than,omitempty"`
	LessThan    string `json:"less_than,omitempty"`
	Between     []int  `json:"between,omitempty"`

	// Collection conditions
	In       []string `json:"in,omitempty"`
	NotIn    []string `json:"not_in,omitempty"`
	Contains string   `json:"contains,omitempty"`

	// Permission conditions
	Can    string `json:"can,omitempty"`
	Cannot string `json:"cannot,omitempty"`
	Role   string `json:"role,omitempty"`
}

// ConditionalCase represents a single case in a switch statement
type ConditionalCase struct {
	When string      `json:"when"`
	Then interface{} `json:"then"`
}

// ============================================================================
// Layout Configuration Types
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
// Helper methods for enterprise features
// IsCSRFEnabled checks if CSRF protection is enabled
func (s *Security) IsCSRFEnabled() bool {
	return s.CSRF != nil && s.CSRF.Enabled
}

// IsRateLimitEnabled checks if rate limiting is enabled
func (s *Security) IsRateLimitEnabled() bool {
	return s.RateLimit != nil && s.RateLimit.Enabled
}

// IsEncryptionEnabled checks if encryption is enabled
func (s *Security) IsEncryptionEnabled() bool {
	return s.Encryption != nil && s.Encryption.Enabled
}

// ShouldEncryptField checks if a specific field should be encrypted
func (s *Security) ShouldEncryptField(fieldName string) bool {
	if !s.IsEncryptionEnabled() {
		return false
	}
	return slices.Contains(s.Encryption.Fields, fieldName)
}

// IsTenantEnabled checks if multi-tenancy is enabled
func (t *Tenant) IsTenantEnabled() bool {
	return t != nil && t.Enabled
}

// IsWorkflowEnabled checks if workflow is enabled
func (w *Workflow) IsWorkflowEnabled() bool {
	return w != nil && w.Enabled
}

// GetCurrentStage returns the current workflow stage
func (w *Workflow) GetCurrentStage() string {
	if w == nil {
		return ""
	}
	return w.Stage
}

// GetCurrentStatus returns the current workflow status
func (w *Workflow) GetCurrentStatus() string {
	if w == nil {
		return ""
	}
	return w.Status
}

// GetAvailableActions returns actions available in current stage
func (w *Workflow) GetAvailableActions(ctx context.Context, data map[string]any, evaluator *condition.Evaluator) []WorkflowAction {
	if w == nil || !w.Enabled {
		return []WorkflowAction{}
	}
	available := make([]WorkflowAction, 0)
	for _, action := range w.Actions {
		// Check if action has conditions
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

// SchemaI18n holds schema-level translations
type SchemaI18n struct {
	Title       map[string]string `json:"title,omitempty"`
	Description map[string]string `json:"description,omitempty"`
	Direction   map[string]string `json:"direction,omitempty"`
}

// IsRTLForSchemaI18n checks if a locale requires right-to-left text direction
func IsRTLForSchemaI18n(schemaI18n *SchemaI18n, locale string) bool {
	// Check if direction is explicitly set in SchemaI18n
	if schemaI18n != nil && schemaI18n.Direction != nil {
		if direction, ok := schemaI18n.Direction[locale]; ok {
			return direction == "rtl"
		}
	}

	// Fall back to default RTL language detection
	rtlLanguages := map[string]bool{
		"ar": true, // Arabic
		"he": true, // Hebrew
		"fa": true, // Persian/Farsi
		"ur": true, // Urdu
	}
	return rtlLanguages[locale]
}

// IsI18nEnabled checks if internationalization is enabled
func (i *I18nManager) IsI18nEnabled() bool {
	return i != nil && i.config != nil && i.config.DefaultLocale != ""
}

// GetLocale returns the default locale or fallback
func (i *I18nManager) GetLocale() string {
	if i == nil || i.config == nil {
		return "en-US"
	}
	if i.config.DefaultLocale != "" {
		return i.config.DefaultLocale
	}
	return "en-US"
}

// IsHTMXEnabled checks if HTMX is enabled
func (h *HTMX) IsHTMXEnabled() bool {
	return h != nil && h.Enabled
}

// IsAlpineEnabled checks if Alpine.js is enabled
func (a *Alpine) IsAlpineEnabled() bool {
	return a != nil && a.Enabled
}

// GetHTTPMethod returns the appropriate HTTP method for HTMX
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

// GetURL returns the appropriate URL for HTMX based on method
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

// AddChangelog adds a new changelog entry
func (m *Meta) AddChangelog(entry ChangelogEntry) {
	if m.Changelog == nil {
		m.Changelog = []ChangelogEntry{}
	}
	m.Changelog = append(m.Changelog, entry)
}

// GetLatestVersion returns the latest version from changelog
func (m *Meta) GetLatestVersion() string {
	if m == nil || len(m.Changelog) == 0 {
		return ""
	}
	return m.Changelog[len(m.Changelog)-1].Version
}

// IsDeprecated checks if the schema is deprecated
func (m *Meta) IsDeprecated() bool {
	return m != nil && m.Deprecated
}

// IsExperimental checks if the schema is experimental
func (m *Meta) IsExperimental() bool {
	return m != nil && m.Experimental
}
