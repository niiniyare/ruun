package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/utils"
)

// ModalSchema represents a schema-driven modal configuration
type ModalSchema struct {
	// Basic metadata
	ID          string `json:"id" yaml:"id"`
	Version     string `json:"version" yaml:"version"`
	Name        string `json:"name" yaml:"name"`
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Modal configuration
	Type     string `json:"type" yaml:"type"` // "dialog", "drawer", "popover", "fullscreen", "bottomsheet"
	Size     string `json:"size,omitempty" yaml:"size,omitempty"`
	Position string `json:"position,omitempty" yaml:"position,omitempty"`
	Variant  string `json:"variant,omitempty" yaml:"variant,omitempty"`

	// Layout and behavior
	Layout   ModalLayout   `json:"layout" yaml:"layout"`
	Behavior ModalBehavior `json:"behavior" yaml:"behavior"`

	// Content sections
	Header *ModalHeaderSchema `json:"header,omitempty" yaml:"header,omitempty"`
	Body   *ModalBodySchema   `json:"body,omitempty" yaml:"body,omitempty"`
	Footer *ModalFooterSchema `json:"footer,omitempty" yaml:"footer,omitempty"`

	// Multi-step workflow
	Steps    []ModalStepSchema `json:"steps,omitempty" yaml:"steps,omitempty"`
	Workflow *WorkflowSchema   `json:"workflow,omitempty" yaml:"workflow,omitempty"`

	// Actions and validation
	Actions    []ModalActionSchema `json:"actions,omitempty" yaml:"actions,omitempty"`
	Validation *ValidationSchema   `json:"validation,omitempty" yaml:"validation,omitempty"`

	// Data and integration
	DataSources []DataSourceSchema `json:"dataSources,omitempty" yaml:"dataSources,omitempty"`
	FormSchema  *schema.FormSchema `json:"formSchema,omitempty" yaml:"formSchema,omitempty"`

	// Business logic integration
	Permissions []PermissionRule        `json:"permissions,omitempty" yaml:"permissions,omitempty"`
	Events      map[string]EventHandler `json:"events,omitempty" yaml:"events,omitempty"`

	// Styling and theming
	Theme *ModalThemeSchema `json:"theme,omitempty" yaml:"theme,omitempty"`

	// Context and metadata
	Context  map[string]any    `json:"context,omitempty" yaml:"context,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// Schema version and compatibility
	SchemaVersion string    `json:"schemaVersion" yaml:"schemaVersion"`
	CreatedAt     time.Time `json:"createdAt" yaml:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt" yaml:"updatedAt"`
	CreatedBy     string    `json:"createdBy,omitempty" yaml:"createdBy,omitempty"`
}

// ModalLayout defines layout configuration
type ModalLayout struct {
	// Responsive design
	Mobile  LayoutConfig `json:"mobile,omitempty" yaml:"mobile,omitempty"`
	Tablet  LayoutConfig `json:"tablet,omitempty" yaml:"tablet,omitempty"`
	Desktop LayoutConfig `json:"desktop,omitempty" yaml:"desktop,omitempty"`

	// Content layout
	ContentFlow string `json:"contentFlow,omitempty" yaml:"contentFlow,omitempty"` // "vertical", "horizontal", "grid"
	Spacing     string `json:"spacing,omitempty" yaml:"spacing,omitempty"`         // "compact", "normal", "spacious"
	Alignment   string `json:"alignment,omitempty" yaml:"alignment,omitempty"`     // "left", "center", "right"

	// Advanced layout
	Sections []LayoutSection `json:"sections,omitempty" yaml:"sections,omitempty"`
}

// LayoutConfig defines responsive layout configuration
type LayoutConfig struct {
	Type       string   `json:"type,omitempty" yaml:"type,omitempty"`         // Override modal type for this breakpoint
	Size       string   `json:"size,omitempty" yaml:"size,omitempty"`         // Override size
	Position   string   `json:"position,omitempty" yaml:"position,omitempty"` // Override position
	Fullscreen bool     `json:"fullscreen,omitempty" yaml:"fullscreen,omitempty"`
	Hidden     []string `json:"hidden,omitempty" yaml:"hidden,omitempty"` // Hide sections on this breakpoint
}

// LayoutSection defines custom layout sections
type LayoutSection struct {
	ID        string            `json:"id" yaml:"id"`
	Type      string            `json:"type" yaml:"type"` // "header", "body", "footer", "sidebar", "custom"
	Order     int               `json:"order,omitempty" yaml:"order,omitempty"`
	Width     string            `json:"width,omitempty" yaml:"width,omitempty"`
	Height    string            `json:"height,omitempty" yaml:"height,omitempty"`
	Condition string            `json:"condition,omitempty" yaml:"condition,omitempty"` // Show/hide condition
	Content   *SectionContent   `json:"content,omitempty" yaml:"content,omitempty"`
	Style     map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
}

// SectionContent defines content for layout sections
type SectionContent struct {
	Type      string         `json:"type" yaml:"type"`                         // "form", "text", "html", "component"
	Source    string         `json:"source,omitempty" yaml:"source,omitempty"` // URL, template path, etc.
	Data      map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	Template  string         `json:"template,omitempty" yaml:"template,omitempty"`
	Component string         `json:"component,omitempty" yaml:"component,omitempty"`
}

// ModalBehavior defines behavioral configuration
type ModalBehavior struct {
	// Core behavior
	Dismissible    bool `json:"dismissible" yaml:"dismissible"`
	CloseOnOverlay bool `json:"closeOnOverlay" yaml:"closeOnOverlay"`
	CloseOnEscape  bool `json:"closeOnEscape" yaml:"closeOnEscape"`
	PersistOnClose bool `json:"persistOnClose,omitempty" yaml:"persistOnClose,omitempty"`

	// Focus management
	AutoFocus    bool `json:"autoFocus,omitempty" yaml:"autoFocus,omitempty"`
	TrapFocus    bool `json:"trapFocus,omitempty" yaml:"trapFocus,omitempty"`
	RestoreFocus bool `json:"restoreFocus,omitempty" yaml:"restoreFocus,omitempty"`

	// Animation and transitions
	Animation AnimationConfig `json:"animation,omitempty" yaml:"animation,omitempty"`

	// Auto-save and data persistence
	AutoSave         bool `json:"autoSave,omitempty" yaml:"autoSave,omitempty"`
	AutoSaveInterval int  `json:"autoSaveInterval,omitempty" yaml:"autoSaveInterval,omitempty"`
	VersionControl   bool `json:"versionControl,omitempty" yaml:"versionControl,omitempty"`

	// Confirmation and warnings
	ConfirmOnClose bool `json:"confirmOnClose,omitempty" yaml:"confirmOnClose,omitempty"`
	DirtyWarning   bool `json:"dirtyWarning,omitempty" yaml:"dirtyWarning,omitempty"`

	// Loading and error handling
	LoadingStrategy string              `json:"loadingStrategy,omitempty" yaml:"loadingStrategy,omitempty"` // "eager", "lazy", "on-demand"
	ErrorHandling   ErrorHandlingConfig `json:"errorHandling,omitempty" yaml:"errorHandling,omitempty"`

	// Modal stack and interaction
	StackBehavior StackBehaviorConfig `json:"stackBehavior,omitempty" yaml:"stackBehavior,omitempty"`

	// Responsive behavior
	ResponsiveBehavior map[string]BehaviorOverride `json:"responsiveBehavior,omitempty" yaml:"responsiveBehavior,omitempty"`
}

// AnimationConfig defines animation configuration
type AnimationConfig struct {
	Enabled           bool              `json:"enabled" yaml:"enabled"`
	Duration          int               `json:"duration,omitempty" yaml:"duration,omitempty"` // milliseconds
	Easing            string            `json:"easing,omitempty" yaml:"easing,omitempty"`
	EnterAnimation    string            `json:"enterAnimation,omitempty" yaml:"enterAnimation,omitempty"`
	LeaveAnimation    string            `json:"leaveAnimation,omitempty" yaml:"leaveAnimation,omitempty"`
	CustomTransitions map[string]string `json:"customTransitions,omitempty" yaml:"customTransitions,omitempty"`
}

// ErrorHandlingConfig defines error handling configuration
type ErrorHandlingConfig struct {
	Strategy       string            `json:"strategy" yaml:"strategy"` // "inline", "toast", "modal", "silent"
	RetryEnabled   bool              `json:"retryEnabled,omitempty" yaml:"retryEnabled,omitempty"`
	MaxRetries     int               `json:"maxRetries,omitempty" yaml:"maxRetries,omitempty"`
	FallbackAction string            `json:"fallbackAction,omitempty" yaml:"fallbackAction,omitempty"`
	CustomHandlers map[string]string `json:"customHandlers,omitempty" yaml:"customHandlers,omitempty"`
}

// StackBehaviorConfig defines modal stack behavior
type StackBehaviorConfig struct {
	Enabled          bool   `json:"enabled" yaml:"enabled"`
	MaxDepth         int    `json:"maxDepth,omitempty" yaml:"maxDepth,omitempty"`
	Strategy         string `json:"strategy,omitempty" yaml:"strategy,omitempty"`                 // "stack", "replace", "queue"
	BackdropBehavior string `json:"backdropBehavior,omitempty" yaml:"backdropBehavior,omitempty"` // "single", "layered", "none"
}

// BehaviorOverride defines responsive behavior overrides
type BehaviorOverride struct {
	CloseOnOverlay  *bool  `json:"closeOnOverlay,omitempty" yaml:"closeOnOverlay,omitempty"`
	TrapFocus       *bool  `json:"trapFocus,omitempty" yaml:"trapFocus,omitempty"`
	AutoSave        *bool  `json:"autoSave,omitempty" yaml:"autoSave,omitempty"`
	LoadingStrategy string `json:"loadingStrategy,omitempty" yaml:"loadingStrategy,omitempty"`
}

// Modal section schemas

// ModalHeaderSchema defines header configuration
type ModalHeaderSchema struct {
	Visible     bool   `json:"visible" yaml:"visible"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty" yaml:"subtitle,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Icon        string `json:"icon,omitempty" yaml:"icon,omitempty"`

	// Header actions
	Actions []HeaderAction `json:"actions,omitempty" yaml:"actions,omitempty"`

	// Progress indicator
	ShowProgress bool   `json:"showProgress,omitempty" yaml:"showProgress,omitempty"`
	ProgressType string `json:"progressType,omitempty" yaml:"progressType,omitempty"` // "bar", "steps", "circular"

	// Customization
	Template  string            `json:"template,omitempty" yaml:"template,omitempty"`
	Component string            `json:"component,omitempty" yaml:"component,omitempty"`
	Style     map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
}

// HeaderAction defines header-specific actions
type HeaderAction struct {
	Type      string `json:"type" yaml:"type"` // "close", "minimize", "help", "custom"
	Icon      string `json:"icon,omitempty" yaml:"icon,omitempty"`
	Text      string `json:"text,omitempty" yaml:"text,omitempty"`
	Action    string `json:"action,omitempty" yaml:"action,omitempty"`
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`
}

// ModalBodySchema defines body configuration
type ModalBodySchema struct {
	Scrollable bool   `json:"scrollable" yaml:"scrollable"`
	Padding    string `json:"padding,omitempty" yaml:"padding,omitempty"`

	// Content configuration
	ContentType   string `json:"contentType" yaml:"contentType"` // "form", "steps", "search", "upload", "custom"
	ContentSource string `json:"contentSource,omitempty" yaml:"contentSource,omitempty"`

	// Form configuration
	FormConfig *FormConfig `json:"formConfig,omitempty" yaml:"formConfig,omitempty"`

	// Multi-step configuration
	StepConfig *StepConfig `json:"stepConfig,omitempty" yaml:"stepConfig,omitempty"`

	// Search configuration
	SearchConfig *SearchConfig `json:"searchConfig,omitempty" yaml:"searchConfig,omitempty"`

	// Upload configuration
	UploadConfig *UploadConfig `json:"uploadConfig,omitempty" yaml:"uploadConfig,omitempty"`

	// Custom content
	CustomContent []ContentBlock `json:"customContent,omitempty" yaml:"customContent,omitempty"`

	// Styling
	Style map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
}

// FormConfig defines form-specific configuration
type FormConfig struct {
	SchemaRef      string             `json:"schemaRef,omitempty" yaml:"schemaRef,omitempty"`
	InlineSchema   *schema.FormSchema `json:"inlineSchema,omitempty" yaml:"inlineSchema,omitempty"`
	Layout         string             `json:"layout,omitempty" yaml:"layout,omitempty"` // "vertical", "horizontal", "grid"
	Columns        int                `json:"columns,omitempty" yaml:"columns,omitempty"`
	Validation     ValidationConfig   `json:"validation" yaml:"validation"`
	SubmitBehavior SubmitConfig       `json:"submitBehavior" yaml:"submitBehavior"`
}

// StepConfig defines multi-step configuration
type StepConfig struct {
	ShowNavigation bool   `json:"showNavigation" yaml:"showNavigation"`
	LinearProgress bool   `json:"linearProgress" yaml:"linearProgress"`
	NavigationType string `json:"navigationType,omitempty" yaml:"navigationType,omitempty"` // "buttons", "tabs", "progress"
	StepValidation bool   `json:"stepValidation" yaml:"stepValidation"`
	SaveOnStep     bool   `json:"saveOnStep,omitempty" yaml:"saveOnStep,omitempty"`
}

// SearchConfig defines search configuration
type SearchConfig struct {
	Endpoint       string          `json:"endpoint" yaml:"endpoint"`
	Placeholder    string          `json:"placeholder,omitempty" yaml:"placeholder,omitempty"`
	MinCharacters  int             `json:"minCharacters,omitempty" yaml:"minCharacters,omitempty"`
	DebounceMs     int             `json:"debounceMs,omitempty" yaml:"debounceMs,omitempty"`
	MaxResults     int             `json:"maxResults,omitempty" yaml:"maxResults,omitempty"`
	Filters        []SearchFilter  `json:"filters,omitempty" yaml:"filters,omitempty"`
	ResultTemplate string          `json:"resultTemplate,omitempty" yaml:"resultTemplate,omitempty"`
	SelectBehavior SelectionConfig `json:"selectBehavior" yaml:"selectBehavior"`
}

// SearchFilter defines search filters
type SearchFilter struct {
	Name         string         `json:"name" yaml:"name"`
	Type         string         `json:"type" yaml:"type"` // "text", "select", "date", "range"
	Label        string         `json:"label,omitempty" yaml:"label,omitempty"`
	Options      []FilterOption `json:"options,omitempty" yaml:"options,omitempty"`
	DefaultValue any            `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`
}

// FilterOption defines filter option
type FilterOption struct {
	Value string `json:"value" yaml:"value"`
	Label string `json:"label" yaml:"label"`
	Icon  string `json:"icon,omitempty" yaml:"icon,omitempty"`
}

// SelectionConfig defines selection behavior
type SelectionConfig struct {
	MultiSelect      bool   `json:"multiSelect" yaml:"multiSelect"`
	MaxSelections    int    `json:"maxSelections,omitempty" yaml:"maxSelections,omitempty"`
	SelectionMode    string `json:"selectionMode,omitempty" yaml:"selectionMode,omitempty"` // "click", "checkbox", "radio"
	ConfirmSelection bool   `json:"confirmSelection,omitempty" yaml:"confirmSelection,omitempty"`
}

// UploadConfig defines file upload configuration
type UploadConfig struct {
	Endpoint        string             `json:"endpoint" yaml:"endpoint"`
	MaxFiles        int                `json:"maxFiles,omitempty" yaml:"maxFiles,omitempty"`
	MaxFileSize     int                `json:"maxFileSize,omitempty" yaml:"maxFileSize,omitempty"` // MB
	AcceptedTypes   []string           `json:"acceptedTypes,omitempty" yaml:"acceptedTypes,omitempty"`
	UploadMode      string             `json:"uploadMode,omitempty" yaml:"uploadMode,omitempty"` // "immediate", "batch", "manual"
	ShowProgress    bool               `json:"showProgress" yaml:"showProgress"`
	PreviewEnabled  bool               `json:"previewEnabled,omitempty" yaml:"previewEnabled,omitempty"`
	ValidationRules []UploadValidation `json:"validationRules,omitempty" yaml:"validationRules,omitempty"`
}

// UploadValidation defines upload validation rules
type UploadValidation struct {
	Type          string `json:"type" yaml:"type"` // "size", "type", "dimension", "custom"
	Value         any    `json:"value,omitempty" yaml:"value,omitempty"`
	Message       string `json:"message,omitempty" yaml:"message,omitempty"`
	ErrorStrategy string `json:"errorStrategy,omitempty" yaml:"errorStrategy,omitempty"` // "block", "warn", "skip"
}

// ContentBlock defines custom content blocks
type ContentBlock struct {
	ID        string            `json:"id" yaml:"id"`
	Type      string            `json:"type" yaml:"type"` // "text", "html", "component", "template"
	Content   string            `json:"content,omitempty" yaml:"content,omitempty"`
	Data      map[string]any    `json:"data,omitempty" yaml:"data,omitempty"`
	Condition string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	Order     int               `json:"order,omitempty" yaml:"order,omitempty"`
	Style     map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
	Events    map[string]string `json:"events,omitempty" yaml:"events,omitempty"`
}

// ModalFooterSchema defines footer configuration
type ModalFooterSchema struct {
	Visible bool   `json:"visible" yaml:"visible"`
	Layout  string `json:"layout,omitempty" yaml:"layout,omitempty"` // "left", "right", "center", "between", "around"

	// Action groups
	LeftActions   []string `json:"leftActions,omitempty" yaml:"leftActions,omitempty"` // Action IDs
	CenterActions []string `json:"centerActions,omitempty" yaml:"centerActions,omitempty"`
	RightActions  []string `json:"rightActions,omitempty" yaml:"rightActions,omitempty"`

	// Step navigation (for multi-step modals)
	ShowStepNav  bool   `json:"showStepNav,omitempty" yaml:"showStepNav,omitempty"`
	StepNavStyle string `json:"stepNavStyle,omitempty" yaml:"stepNavStyle,omitempty"` // "buttons", "text", "minimal"

	// Customization
	Template string            `json:"template,omitempty" yaml:"template,omitempty"`
	Style    map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
}

// Workflow and step schemas

// WorkflowSchema defines workflow configuration
type WorkflowSchema struct {
	ID        string `json:"id" yaml:"id"`
	Type      string `json:"type,omitempty" yaml:"type,omitempty"` // "linear", "conditional", "parallel"
	StartStep string `json:"startStep,omitempty" yaml:"startStep,omitempty"`

	// Workflow behavior
	AllowSkipping  bool `json:"allowSkipping,omitempty" yaml:"allowSkipping,omitempty"`
	SaveProgress   bool `json:"saveProgress,omitempty" yaml:"saveProgress,omitempty"`
	ResetOnRestart bool `json:"resetOnRestart,omitempty" yaml:"resetOnRestart,omitempty"`

	// Conditional logic
	Rules     []WorkflowRule `json:"rules,omitempty" yaml:"rules,omitempty"`
	Variables map[string]any `json:"variables,omitempty" yaml:"variables,omitempty"`

	// Integration
	Hooks       map[string]string `json:"hooks,omitempty" yaml:"hooks,omitempty"`
	ExternalAPI *APIConfig        `json:"externalApi,omitempty" yaml:"externalApi,omitempty"`
}

// WorkflowRule defines workflow conditional rules
type WorkflowRule struct {
	ID        string `json:"id" yaml:"id"`
	Condition string `json:"condition" yaml:"condition"`
	Action    string `json:"action" yaml:"action"` // "goto", "skip", "show", "hide", "validate"
	Target    string `json:"target,omitempty" yaml:"target,omitempty"`
	Value     any    `json:"value,omitempty" yaml:"value,omitempty"`
	Priority  int    `json:"priority,omitempty" yaml:"priority,omitempty"`
}

// ModalStepSchema defines individual step configuration
type ModalStepSchema struct {
	ID          string `json:"id" yaml:"id"`
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Icon        string `json:"icon,omitempty" yaml:"icon,omitempty"`

	// Step behavior
	Required   bool `json:"required" yaml:"required"`
	Skippable  bool `json:"skippable,omitempty" yaml:"skippable,omitempty"`
	Repeatable bool `json:"repeatable,omitempty" yaml:"repeatable,omitempty"`

	// Content
	ContentType string             `json:"contentType" yaml:"contentType"` // "form", "info", "confirmation", "custom"
	Content     *StepContent       `json:"content,omitempty" yaml:"content,omitempty"`
	FormSchema  *schema.FormSchema `json:"formSchema,omitempty" yaml:"formSchema,omitempty"`

	// Validation and navigation
	Validation *StepValidation `json:"validation,omitempty" yaml:"validation,omitempty"`
	Navigation *StepNavigation `json:"navigation,omitempty" yaml:"navigation,omitempty"`

	// Conditional behavior
	ShowIf    string `json:"showIf,omitempty" yaml:"showIf,omitempty"`
	EnabledIf string `json:"enabledIf,omitempty" yaml:"enabledIf,omitempty"`

	// Styling and layout
	Style       map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
	CustomClass string            `json:"customClass,omitempty" yaml:"customClass,omitempty"`
}

// StepContent defines step content configuration
type StepContent struct {
	Template   string         `json:"template,omitempty" yaml:"template,omitempty"`
	Component  string         `json:"component,omitempty" yaml:"component,omitempty"`
	HTML       string         `json:"html,omitempty" yaml:"html,omitempty"`
	Data       map[string]any `json:"data,omitempty" yaml:"data,omitempty"`
	DataSource string         `json:"dataSource,omitempty" yaml:"dataSource,omitempty"`
}

// StepValidation defines step-specific validation
type StepValidation struct {
	ValidateOn      string           `json:"validateOn,omitempty" yaml:"validateOn,omitempty"` // "next", "always", "blur", "submit"
	Rules           []ValidationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
	CustomValidator string           `json:"customValidator,omitempty" yaml:"customValidator,omitempty"`
	ErrorMessage    string           `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`
	Required        []string         `json:"required,omitempty" yaml:"required,omitempty"` // Required fields
}

// StepNavigation defines step navigation behavior
type StepNavigation struct {
	NextStep         string `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`
	PrevStep         string `json:"prevStep,omitempty" yaml:"prevStep,omitempty"`
	NextCondition    string `json:"nextCondition,omitempty" yaml:"nextCondition,omitempty"`
	PrevCondition    string `json:"prevCondition,omitempty" yaml:"prevCondition,omitempty"`
	NextLabel        string `json:"nextLabel,omitempty" yaml:"nextLabel,omitempty"`
	PrevLabel        string `json:"prevLabel,omitempty" yaml:"prevLabel,omitempty"`
	AutoAdvance      bool   `json:"autoAdvance,omitempty" yaml:"autoAdvance,omitempty"`
	AutoAdvanceDelay int    `json:"autoAdvanceDelay,omitempty" yaml:"autoAdvanceDelay,omitempty"`
}

// Action and validation schemas

// ModalActionSchema defines modal action configuration
type ModalActionSchema struct {
	ID       string `json:"id" yaml:"id"`
	Type     string `json:"type" yaml:"type"` // "primary", "secondary", "destructive", "submit", "cancel"
	Text     string `json:"text" yaml:"text"`
	Icon     string `json:"icon,omitempty" yaml:"icon,omitempty"`
	Position string `json:"position,omitempty" yaml:"position,omitempty"` // "left", "center", "right"
	Priority int    `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Behavior
	Action       ActionBehavior      `json:"action" yaml:"action"`
	Confirmation *ConfirmationConfig `json:"confirmation,omitempty" yaml:"confirmation,omitempty"`

	// State and conditions
	EnabledIf    string `json:"enabledIf,omitempty" yaml:"enabledIf,omitempty"`
	VisibleIf    string `json:"visibleIf,omitempty" yaml:"visibleIf,omitempty"`
	LoadingState string `json:"loadingState,omitempty" yaml:"loadingState,omitempty"`

	// Validation dependencies
	RequiredFields []string `json:"requiredFields,omitempty" yaml:"requiredFields,omitempty"`
	ValidateForm   bool     `json:"validateForm,omitempty" yaml:"validateForm,omitempty"`

	// Styling
	Variant     string            `json:"variant,omitempty" yaml:"variant,omitempty"`
	Size        string            `json:"size,omitempty" yaml:"size,omitempty"`
	FullWidth   bool              `json:"fullWidth,omitempty" yaml:"fullWidth,omitempty"`
	CustomClass string            `json:"customClass,omitempty" yaml:"customClass,omitempty"`
	Style       map[string]string `json:"style,omitempty" yaml:"style,omitempty"`
}

// ActionBehavior defines action behavior configuration
type ActionBehavior struct {
	Type string `json:"type" yaml:"type"` // "submit", "close", "navigate", "custom", "api"

	// Submit behavior
	SubmitURL    string `json:"submitUrl,omitempty" yaml:"submitUrl,omitempty"`
	SubmitMethod string `json:"submitMethod,omitempty" yaml:"submitMethod,omitempty"`

	// Navigation behavior
	NavigateStep string `json:"navigateStep,omitempty" yaml:"navigateStep,omitempty"`
	NavigateURL  string `json:"navigateUrl,omitempty" yaml:"navigateUrl,omitempty"`
	CloseModal   bool   `json:"closeModal,omitempty" yaml:"closeModal,omitempty"`

	// API behavior
	APICall *APIConfig `json:"apiCall,omitempty" yaml:"apiCall,omitempty"`

	// Custom behavior
	CustomHandler string `json:"customHandler,omitempty" yaml:"customHandler,omitempty"`
	JavaScript    string `json:"javascript,omitempty" yaml:"javascript,omitempty"`

	// Success/error handling
	OnSuccess *ActionResult `json:"onSuccess,omitempty" yaml:"onSuccess,omitempty"`
	OnError   *ActionResult `json:"onError,omitempty" yaml:"onError,omitempty"`

	// Data handling
	DataMapping   map[string]string `json:"dataMapping,omitempty" yaml:"dataMapping,omitempty"`
	TransformData string            `json:"transformData,omitempty" yaml:"transformData,omitempty"`
}

// ActionResult defines action result handling
type ActionResult struct {
	Action        string `json:"action" yaml:"action"` // "close", "navigate", "show_message", "refresh", "custom"
	Message       string `json:"message,omitempty" yaml:"message,omitempty"`
	NavigateURL   string `json:"navigateUrl,omitempty" yaml:"navigateUrl,omitempty"`
	CloseModal    bool   `json:"closeModal,omitempty" yaml:"closeModal,omitempty"`
	RefreshData   bool   `json:"refreshData,omitempty" yaml:"refreshData,omitempty"`
	CustomHandler string `json:"customHandler,omitempty" yaml:"customHandler,omitempty"`
}

// ConfirmationConfig defines action confirmation
type ConfirmationConfig struct {
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	Title       string `json:"title,omitempty" yaml:"title,omitempty"`
	Message     string `json:"message" yaml:"message"`
	ConfirmText string `json:"confirmText,omitempty" yaml:"confirmText,omitempty"`
	CancelText  string `json:"cancelText,omitempty" yaml:"cancelText,omitempty"`
	Type        string `json:"type,omitempty" yaml:"type,omitempty"` // "dialog", "inline", "tooltip"
	Icon        string `json:"icon,omitempty" yaml:"icon,omitempty"`
	Variant     string `json:"variant,omitempty" yaml:"variant,omitempty"`
}

// ValidationSchema defines validation configuration
type ValidationSchema struct {
	Strategy         string                  `json:"strategy" yaml:"strategy"` // "realtime", "onblur", "onsubmit"
	Rules            []GlobalValidationRule  `json:"rules,omitempty" yaml:"rules,omitempty"`
	ServerValidation *ServerValidationConfig `json:"serverValidation,omitempty" yaml:"serverValidation,omitempty"`
	CustomValidators map[string]string       `json:"customValidators,omitempty" yaml:"customValidators,omitempty"`
	ErrorDisplay     ErrorDisplayConfig      `json:"errorDisplay" yaml:"errorDisplay"`
}

// GlobalValidationRule defines global validation rules
type GlobalValidationRule struct {
	ID        string   `json:"id" yaml:"id"`
	Fields    []string `json:"fields" yaml:"fields"`
	Type      string   `json:"type" yaml:"type"` // "required", "pattern", "min", "max", "custom"
	Value     any      `json:"value,omitempty" yaml:"value,omitempty"`
	Message   string   `json:"message" yaml:"message"`
	Condition string   `json:"condition,omitempty" yaml:"condition,omitempty"`
	Priority  int      `json:"priority,omitempty" yaml:"priority,omitempty"`
}

// ServerValidationConfig defines server-side validation
type ServerValidationConfig struct {
	Enabled        bool              `json:"enabled" yaml:"enabled"`
	Endpoint       string            `json:"endpoint" yaml:"endpoint"`
	Method         string            `json:"method,omitempty" yaml:"method,omitempty"`
	DebounceMs     int               `json:"debounceMs,omitempty" yaml:"debounceMs,omitempty"`
	ValidateFields []string          `json:"validateFields,omitempty" yaml:"validateFields,omitempty"`
	Headers        map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// ErrorDisplayConfig defines error display configuration
type ErrorDisplayConfig struct {
	Strategy    string `json:"strategy" yaml:"strategy"`                     // "inline", "toast", "summary", "modal"
	Position    string `json:"position,omitempty" yaml:"position,omitempty"` // "top", "bottom", "field"
	Style       string `json:"style,omitempty" yaml:"style,omitempty"`       // "default", "compact", "detailed"
	AutoClear   bool   `json:"autoClear,omitempty" yaml:"autoClear,omitempty"`
	ClearDelay  int    `json:"clearDelay,omitempty" yaml:"clearDelay,omitempty"`
	GroupErrors bool   `json:"groupErrors,omitempty" yaml:"groupErrors,omitempty"`
}

// ValidationConfig defines form validation configuration
type ValidationConfig struct {
	Enabled        bool             `json:"enabled" yaml:"enabled"`
	Strategy       string           `json:"strategy" yaml:"strategy"`
	Rules          []ValidationRule `json:"rules,omitempty" yaml:"rules,omitempty"`
	ServerEndpoint string           `json:"serverEndpoint,omitempty" yaml:"serverEndpoint,omitempty"`
	RealtimeFields []string         `json:"realtimeFields,omitempty" yaml:"realtimeFields,omitempty"`
}

// SubmitConfig defines form submission configuration
type SubmitConfig struct {
	Endpoint      string `json:"endpoint" yaml:"endpoint"`
	Method        string `json:"method,omitempty" yaml:"method,omitempty"`
	ValidateFirst bool   `json:"validateFirst" yaml:"validateFirst"`
	ConfirmSubmit bool   `json:"confirmSubmit,omitempty" yaml:"confirmSubmit,omitempty"`
	SuccessAction string `json:"successAction,omitempty" yaml:"successAction,omitempty"` // "close", "redirect", "refresh"
	ErrorAction   string `json:"errorAction,omitempty" yaml:"errorAction,omitempty"`     // "stay", "close", "retry"
	DataTransform string `json:"dataTransform,omitempty" yaml:"dataTransform,omitempty"`
}

// ValidationRule defines individual validation rules
type ValidationRule struct {
	Field     string `json:"field" yaml:"field"`
	Type      string `json:"type" yaml:"type"`
	Value     any    `json:"value,omitempty" yaml:"value,omitempty"`
	Message   string `json:"message" yaml:"message"`
	Condition string `json:"condition,omitempty" yaml:"condition,omitempty"`
}

// Data and integration schemas

// DataSourceSchema defines data source configuration
type DataSourceSchema struct {
	ID   string `json:"id" yaml:"id"`
	Type string `json:"type" yaml:"type"` // "api", "static", "function", "storage"

	// API data source
	API *APIConfig `json:"api,omitempty" yaml:"api,omitempty"`

	// Static data source
	StaticData map[string]any `json:"staticData,omitempty" yaml:"staticData,omitempty"`

	// Function data source
	Function string `json:"function,omitempty" yaml:"function,omitempty"`

	// Storage data source
	StorageKey  string `json:"storageKey,omitempty" yaml:"storageKey,omitempty"`
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"` // "local", "session", "cookie"

	// Caching and refresh
	CacheEnabled    bool `json:"cacheEnabled,omitempty" yaml:"cacheEnabled,omitempty"`
	CacheDuration   int  `json:"cacheDuration,omitempty" yaml:"cacheDuration,omitempty"`
	RefreshOnOpen   bool `json:"refreshOnOpen,omitempty" yaml:"refreshOnOpen,omitempty"`
	RefreshInterval int  `json:"refreshInterval,omitempty" yaml:"refreshInterval,omitempty"`

	// Data transformation
	Transform string       `json:"transform,omitempty" yaml:"transform,omitempty"`
	Filter    string       `json:"filter,omitempty" yaml:"filter,omitempty"`
	Sort      []SortConfig `json:"sort,omitempty" yaml:"sort,omitempty"`

	// Error handling
	FallbackData  map[string]any `json:"fallbackData,omitempty" yaml:"fallbackData,omitempty"`
	ErrorStrategy string         `json:"errorStrategy,omitempty" yaml:"errorStrategy,omitempty"`
}

// APIConfig defines API configuration
type APIConfig struct {
	URL            string            `json:"url" yaml:"url"`
	Method         string            `json:"method,omitempty" yaml:"method,omitempty"`
	Headers        map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	QueryParams    map[string]string `json:"queryParams,omitempty" yaml:"queryParams,omitempty"`
	Body           map[string]any    `json:"body,omitempty" yaml:"body,omitempty"`
	Timeout        int               `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	RetryCount     int               `json:"retryCount,omitempty" yaml:"retryCount,omitempty"`
	Authentication *AuthConfig       `json:"authentication,omitempty" yaml:"authentication,omitempty"`
}

// AuthConfig defines authentication configuration
type AuthConfig struct {
	Type          string            `json:"type" yaml:"type"` // "bearer", "basic", "api_key", "custom"
	Token         string            `json:"token,omitempty" yaml:"token,omitempty"`
	Username      string            `json:"username,omitempty" yaml:"username,omitempty"`
	Password      string            `json:"password,omitempty" yaml:"password,omitempty"`
	APIKey        string            `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
	APIKeyHeader  string            `json:"apiKeyHeader,omitempty" yaml:"apiKeyHeader,omitempty"`
	CustomHeaders map[string]string `json:"customHeaders,omitempty" yaml:"customHeaders,omitempty"`
}

// SortConfig defines sorting configuration
type SortConfig struct {
	Field     string `json:"field" yaml:"field"`
	Direction string `json:"direction" yaml:"direction"` // "asc", "desc"
	Priority  int    `json:"priority,omitempty" yaml:"priority,omitempty"`
}

// Permission and event schemas

// PermissionRule defines permission rules
type PermissionRule struct {
	ID             string   `json:"id" yaml:"id"`
	Action         string   `json:"action" yaml:"action"` // "view", "edit", "submit", "delete", "custom"
	Resource       string   `json:"resource,omitempty" yaml:"resource,omitempty"`
	Condition      string   `json:"condition,omitempty" yaml:"condition,omitempty"`
	Required       []string `json:"required" yaml:"required"`                                 // Required permissions/roles
	FallbackAction string   `json:"fallbackAction,omitempty" yaml:"fallbackAction,omitempty"` // "hide", "disable", "redirect"
	ErrorMessage   string   `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`
}

// EventHandler defines event handling configuration
type EventHandler struct {
	Type         string         `json:"type" yaml:"type"` // "javascript", "api", "navigation", "custom"
	Handler      string         `json:"handler" yaml:"handler"`
	Async        bool           `json:"async,omitempty" yaml:"async,omitempty"`
	Parameters   map[string]any `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	Condition    string         `json:"condition,omitempty" yaml:"condition,omitempty"`
	ErrorHandler string         `json:"errorHandler,omitempty" yaml:"errorHandler,omitempty"`
}

// Theme and styling schemas

// ModalThemeSchema defines theme configuration
type ModalThemeSchema struct {
	ID      string `json:"id,omitempty" yaml:"id,omitempty"`
	Extends string `json:"extends,omitempty" yaml:"extends,omitempty"` // Base theme to extend

	// Color scheme
	Colors map[string]string `json:"colors,omitempty" yaml:"colors,omitempty"`

	// Typography
	Typography TypographyConfig `json:"typography,omitempty" yaml:"typography,omitempty"`

	// Spacing and sizing
	Spacing map[string]string `json:"spacing,omitempty" yaml:"spacing,omitempty"`
	Sizing  map[string]string `json:"sizing,omitempty" yaml:"sizing,omitempty"`

	// Border and shadow
	Borders map[string]string `json:"borders,omitempty" yaml:"borders,omitempty"`
	Shadows map[string]string `json:"shadows,omitempty" yaml:"shadows,omitempty"`

	// Animation
	Animations map[string]string `json:"animations,omitempty" yaml:"animations,omitempty"`

	// Component-specific styles
	Components map[string]ComponentTheme `json:"components,omitempty" yaml:"components,omitempty"`

	// Responsive overrides
	Responsive map[string]ResponsiveTheme `json:"responsive,omitempty" yaml:"responsive,omitempty"`

	// CSS variables and custom properties
	CSSVariables map[string]string `json:"cssVariables,omitempty" yaml:"cssVariables,omitempty"`
	CustomCSS    string            `json:"customCss,omitempty" yaml:"customCss,omitempty"`
}

// TypographyConfig defines typography configuration
type TypographyConfig struct {
	FontFamily    string            `json:"fontFamily,omitempty" yaml:"fontFamily,omitempty"`
	FontSize      map[string]string `json:"fontSize,omitempty" yaml:"fontSize,omitempty"`
	FontWeight    map[string]string `json:"fontWeight,omitempty" yaml:"fontWeight,omitempty"`
	LineHeight    map[string]string `json:"lineHeight,omitempty" yaml:"lineHeight,omitempty"`
	LetterSpacing map[string]string `json:"letterSpacing,omitempty" yaml:"letterSpacing,omitempty"`
}

// ComponentTheme defines component-specific theme
type ComponentTheme struct {
	BaseClass string            `json:"baseClass,omitempty" yaml:"baseClass,omitempty"`
	Variants  map[string]string `json:"variants,omitempty" yaml:"variants,omitempty"`
	States    map[string]string `json:"states,omitempty" yaml:"states,omitempty"`
	CustomCSS string            `json:"customCss,omitempty" yaml:"customCss,omitempty"`
}

// ResponsiveTheme defines responsive theme overrides
type ResponsiveTheme struct {
	Colors     map[string]string         `json:"colors,omitempty" yaml:"colors,omitempty"`
	Spacing    map[string]string         `json:"spacing,omitempty" yaml:"spacing,omitempty"`
	Typography TypographyConfig          `json:"typography,omitempty" yaml:"typography,omitempty"`
	Components map[string]ComponentTheme `json:"components,omitempty" yaml:"components,omitempty"`
}

// Schema builder and utility functions

// ModalSchemaBuilder provides a fluent interface for building modal schemas
type ModalSchemaBuilder struct {
	schema *ModalSchema
}

// NewModalSchemaBuilder creates a new modal schema builder
func NewModalSchemaBuilder() *ModalSchemaBuilder {
	return &ModalSchemaBuilder{
		schema: &ModalSchema{
			SchemaVersion: "1.0.0",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			Layout:        ModalLayout{},
			Behavior:      ModalBehavior{},
			Context:       make(map[string]any),
			Metadata:      make(map[string]string),
		},
	}
}

// Basic configuration methods

func (b *ModalSchemaBuilder) ID(id string) *ModalSchemaBuilder {
	b.schema.ID = id
	return b
}

func (b *ModalSchemaBuilder) Name(name string) *ModalSchemaBuilder {
	b.schema.Name = name
	return b
}

func (b *ModalSchemaBuilder) Title(title string) *ModalSchemaBuilder {
	b.schema.Title = title
	return b
}

func (b *ModalSchemaBuilder) Description(description string) *ModalSchemaBuilder {
	b.schema.Description = description
	return b
}

func (b *ModalSchemaBuilder) Type(modalType string) *ModalSchemaBuilder {
	b.schema.Type = modalType
	return b
}

func (b *ModalSchemaBuilder) Size(size string) *ModalSchemaBuilder {
	b.schema.Size = size
	return b
}

func (b *ModalSchemaBuilder) Position(position string) *ModalSchemaBuilder {
	b.schema.Position = position
	return b
}

func (b *ModalSchemaBuilder) Variant(variant string) *ModalSchemaBuilder {
	b.schema.Variant = variant
	return b
}

// Layout configuration

func (b *ModalSchemaBuilder) WithResponsiveLayout(mobile, tablet, desktop LayoutConfig) *ModalSchemaBuilder {
	b.schema.Layout.Mobile = mobile
	b.schema.Layout.Tablet = tablet
	b.schema.Layout.Desktop = desktop
	return b
}

func (b *ModalSchemaBuilder) ContentFlow(flow string) *ModalSchemaBuilder {
	b.schema.Layout.ContentFlow = flow
	return b
}

func (b *ModalSchemaBuilder) Spacing(spacing string) *ModalSchemaBuilder {
	b.schema.Layout.Spacing = spacing
	return b
}

// Behavior configuration

func (b *ModalSchemaBuilder) Dismissible(dismissible bool) *ModalSchemaBuilder {
	b.schema.Behavior.Dismissible = dismissible
	return b
}

func (b *ModalSchemaBuilder) CloseOnOverlay(closeOnOverlay bool) *ModalSchemaBuilder {
	b.schema.Behavior.CloseOnOverlay = closeOnOverlay
	return b
}

func (b *ModalSchemaBuilder) CloseOnEscape(closeOnEscape bool) *ModalSchemaBuilder {
	b.schema.Behavior.CloseOnEscape = closeOnEscape
	return b
}

func (b *ModalSchemaBuilder) AutoFocus(autoFocus bool) *ModalSchemaBuilder {
	b.schema.Behavior.AutoFocus = autoFocus
	return b
}

func (b *ModalSchemaBuilder) TrapFocus(trapFocus bool) *ModalSchemaBuilder {
	b.schema.Behavior.TrapFocus = trapFocus
	return b
}

func (b *ModalSchemaBuilder) WithAnimation(config AnimationConfig) *ModalSchemaBuilder {
	b.schema.Behavior.Animation = config
	return b
}

func (b *ModalSchemaBuilder) AutoSave(enabled bool, interval int) *ModalSchemaBuilder {
	b.schema.Behavior.AutoSave = enabled
	b.schema.Behavior.AutoSaveInterval = interval
	return b
}

func (b *ModalSchemaBuilder) ErrorHandling(config ErrorHandlingConfig) *ModalSchemaBuilder {
	b.schema.Behavior.ErrorHandling = config
	return b
}

func (b *ModalSchemaBuilder) StackBehavior(config StackBehaviorConfig) *ModalSchemaBuilder {
	b.schema.Behavior.StackBehavior = config
	return b
}

// Section configuration

func (b *ModalSchemaBuilder) WithHeader(header ModalHeaderSchema) *ModalSchemaBuilder {
	b.schema.Header = &header
	return b
}

func (b *ModalSchemaBuilder) WithBody(body ModalBodySchema) *ModalSchemaBuilder {
	b.schema.Body = &body
	return b
}

func (b *ModalSchemaBuilder) WithFooter(footer ModalFooterSchema) *ModalSchemaBuilder {
	b.schema.Footer = &footer
	return b
}

// Multi-step configuration

func (b *ModalSchemaBuilder) AddStep(step ModalStepSchema) *ModalSchemaBuilder {
	b.schema.Steps = append(b.schema.Steps, step)
	return b
}

func (b *ModalSchemaBuilder) WithWorkflow(workflow WorkflowSchema) *ModalSchemaBuilder {
	b.schema.Workflow = &workflow
	return b
}

// Action configuration

func (b *ModalSchemaBuilder) AddAction(action ModalActionSchema) *ModalSchemaBuilder {
	b.schema.Actions = append(b.schema.Actions, action)
	return b
}

// Convenience methods for common actions

func (b *ModalSchemaBuilder) AddPrimaryAction(text, actionType string) *ModalSchemaBuilder {
	action := ModalActionSchema{
		ID:       "primary",
		Type:     "primary",
		Text:     text,
		Position: "right",
		Priority: 1,
		Action: ActionBehavior{
			Type: actionType,
		},
		Variant: "primary",
	}
	return b.AddAction(action)
}

func (b *ModalSchemaBuilder) AddSecondaryAction(text, actionType string) *ModalSchemaBuilder {
	action := ModalActionSchema{
		ID:       "secondary",
		Type:     "secondary",
		Text:     text,
		Position: "right",
		Priority: 2,
		Action: ActionBehavior{
			Type: actionType,
		},
		Variant: "secondary",
	}
	return b.AddAction(action)
}

func (b *ModalSchemaBuilder) AddCancelAction(text string) *ModalSchemaBuilder {
	action := ModalActionSchema{
		ID:       "cancel",
		Type:     "cancel",
		Text:     text,
		Position: "right",
		Priority: 10,
		Action: ActionBehavior{
			Type:       "close",
			CloseModal: true,
		},
		Variant: "outline",
	}
	return b.AddAction(action)
}

func (b *ModalSchemaBuilder) AddDestructiveAction(text, actionType string) *ModalSchemaBuilder {
	action := ModalActionSchema{
		ID:       "destructive",
		Type:     "destructive",
		Text:     text,
		Position: "right",
		Priority: 1,
		Action: ActionBehavior{
			Type: actionType,
		},
		Variant: "destructive",
		Confirmation: &ConfirmationConfig{
			Enabled: true,
			Message: "Are you sure you want to proceed? This action cannot be undone.",
		},
	}
	return b.AddAction(action)
}

// Validation configuration

func (b *ModalSchemaBuilder) WithValidation(validation ValidationSchema) *ModalSchemaBuilder {
	b.schema.Validation = &validation
	return b
}

// Data source configuration

func (b *ModalSchemaBuilder) AddDataSource(source DataSourceSchema) *ModalSchemaBuilder {
	b.schema.DataSources = append(b.schema.DataSources, source)
	return b
}

func (b *ModalSchemaBuilder) WithFormSchema(formSchema *schema.FormSchema) *ModalSchemaBuilder {
	b.schema.FormSchema = formSchema
	return b
}

// Permission configuration

func (b *ModalSchemaBuilder) AddPermission(rule PermissionRule) *ModalSchemaBuilder {
	b.schema.Permissions = append(b.schema.Permissions, rule)
	return b
}

// Event configuration

func (b *ModalSchemaBuilder) OnEvent(eventName string, handler EventHandler) *ModalSchemaBuilder {
	if b.schema.Events == nil {
		b.schema.Events = make(map[string]EventHandler)
	}
	b.schema.Events[eventName] = handler
	return b
}

// Theme configuration

func (b *ModalSchemaBuilder) WithTheme(theme ModalThemeSchema) *ModalSchemaBuilder {
	b.schema.Theme = &theme
	return b
}

// Metadata and context

func (b *ModalSchemaBuilder) SetContext(key string, value any) *ModalSchemaBuilder {
	b.schema.Context[key] = value
	return b
}

func (b *ModalSchemaBuilder) SetMetadata(key, value string) *ModalSchemaBuilder {
	b.schema.Metadata[key] = value
	return b
}

func (b *ModalSchemaBuilder) CreatedBy(creator string) *ModalSchemaBuilder {
	b.schema.CreatedBy = creator
	return b
}

// Build method

func (b *ModalSchemaBuilder) Build() *ModalSchema {
	b.schema.UpdatedAt = time.Now()
	return b.schema
}

// Schema converter functions

// ConvertSchemaToProps converts a ModalSchema to EnhancedModalProps
func ConvertSchemaToProps(ctx context.Context, schema *ModalSchema, data map[string]any) (*EnhancedModalProps, error) {
	if schema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}

	props := &EnhancedModalProps{
		ID:          schema.ID,
		Name:        schema.Name,
		Title:       schema.Title,
		Subtitle:    "", // Could be derived from schema
		Description: schema.Description,

		// Convert type
		Type:     convertSchemaTypeToEnhancedType(schema.Type),
		Size:     convertSchemaSize(schema.Size),
		Position: convertSchemaPosition(schema.Position),
		Variant:  convertSchemaVariant(schema.Variant),

		// Convert behavior
		Dismissible:    schema.Behavior.Dismissible,
		CloseOnOverlay: schema.Behavior.CloseOnOverlay,
		CloseOnEscape:  schema.Behavior.CloseOnEscape,
		PersistOnClose: schema.Behavior.PersistOnClose,
		AutoFocus:      schema.Behavior.AutoFocus,
		TrapFocus:      schema.Behavior.TrapFocus,
		RestoreFocus:   schema.Behavior.RestoreFocus,

		// Convert layout
		Backdrop:     true,  // Default
		BlurBackdrop: false, // Could be configurable
		Centered:     true,  // Default for dialogs
		Scrollable:   schema.Body != nil && schema.Body.Scrollable,

		// Convert structure
		Header:       schema.Header != nil && schema.Header.Visible,
		Footer:       schema.Footer != nil && schema.Footer.Visible,
		ShowProgress: schema.Header != nil && schema.Header.ShowProgress,

		// Convert auto-save
		AutoSave:         schema.Behavior.AutoSave,
		AutoSaveInterval: schema.Behavior.AutoSaveInterval,

		// Convert animation
		AnimationDuration: schema.Behavior.Animation.Duration,

		// Form data
		FormData:       data,
		FormValidation: schema.Validation != nil && schema.Validation.Strategy != "",

		// Convert form schema if present
		FormSchema: schema.FormSchema,

		// Default context
		Context: ctx,
	}

	// Convert header configuration
	if schema.Header != nil {
		if schema.Header.Icon != "" {
			props.Icon = schema.Header.Icon
		}
		if schema.Header.Title != "" && props.Title == "" {
			props.Title = schema.Header.Title
		}
	}

	// Convert steps if present
	if len(schema.Steps) > 0 {
		props.Steps = convertSchemaStepsToEnhancedSteps(schema.Steps)
		props.CurrentStep = 1
		props.ShowStepNav = schema.Body != nil && schema.Body.StepConfig != nil && schema.Body.StepConfig.ShowNavigation
		props.LinearProgress = schema.Body != nil && schema.Body.StepConfig != nil && schema.Body.StepConfig.LinearProgress
	}

	// Convert actions
	if len(schema.Actions) > 0 {
		actions, primaryAction, secondaryAction, destructiveAction := convertSchemaActionsToEnhancedActions(schema.Actions)
		props.Actions = actions
		props.PrimaryAction = primaryAction
		props.SecondaryAction = secondaryAction
		props.DestructiveAction = destructiveAction
	}

	// Convert body configuration
	if schema.Body != nil {
		convertBodyConfiguration(schema.Body, props)
	}

	// Convert validation configuration
	if schema.Validation != nil {
		convertValidationConfiguration(schema.Validation, props)
	}

	// Convert permissions
	if len(schema.Permissions) > 0 {
		props.RequiredPermissions = extractRequiredPermissions(schema.Permissions)
		props.SecureMode = true
	}

	// Convert theme
	if schema.Theme != nil {
		convertThemeConfiguration(schema.Theme, props)
	}

	return props, nil
}

// Helper conversion functions

func convertSchemaTypeToEnhancedType(schemaType string) EnhancedModalType {
	switch strings.ToLower(schemaType) {
	case "dialog":
		return EnhancedModalDialog
	case "drawer":
		return EnhancedModalDrawer
	case "popover":
		return EnhancedModalPopover
	case "fullscreen":
		return EnhancedModalFullscreen
	case "bottomsheet":
		return EnhancedModalBottomSheet
	default:
		return EnhancedModalDialog
	}
}

func convertSchemaSize(size string) EnhancedModalSize {
	switch strings.ToLower(size) {
	case "xs":
		return EnhancedModalSizeXS
	case "sm":
		return EnhancedModalSizeSM
	case "md":
		return EnhancedModalSizeMD
	case "lg":
		return EnhancedModalSizeLG
	case "xl":
		return EnhancedModalSizeXL
	case "2xl":
		return EnhancedModalSize2XL
	case "3xl":
		return EnhancedModalSize3XL
	case "full":
		return EnhancedModalSizeFull
	case "auto":
		return EnhancedModalSizeAuto
	default:
		return EnhancedModalSizeMD
	}
}

func convertSchemaPosition(position string) EnhancedModalPosition {
	switch strings.ToLower(position) {
	case "center":
		return EnhancedModalPositionCenter
	case "top":
		return EnhancedModalPositionTop
	case "right":
		return EnhancedModalPositionRight
	case "bottom":
		return EnhancedModalPositionBottom
	case "left":
		return EnhancedModalPositionLeft
	default:
		return EnhancedModalPositionCenter
	}
}

func convertSchemaVariant(variant string) EnhancedModalVariant {
	switch strings.ToLower(variant) {
	case "destructive":
		return EnhancedModalDestructive
	case "success":
		return EnhancedModalSuccess
	case "warning":
		return EnhancedModalWarning
	case "info":
		return EnhancedModalInfo
	default:
		return EnhancedModalDefault
	}
}

func convertSchemaStepsToEnhancedSteps(steps []ModalStepSchema) []EnhancedModalStep {
	enhancedSteps := make([]EnhancedModalStep, len(steps))
	for i, step := range steps {
		enhancedSteps[i] = EnhancedModalStep{
			ID:          step.ID,
			Title:       step.Title,
			Description: step.Description,
			Icon:        step.Icon,
			Required:    step.Required,
			Skippable:   step.Skippable,
			Content:     step.FormSchema,
			ShowIf:      step.ShowIf,
		}

		// Convert step validation
		if step.Validation != nil {
			enhancedSteps[i].ValidateOn = step.Validation.ValidateOn
			enhancedSteps[i].Validator = step.Validation.CustomValidator
		}

		// Convert step navigation
		if step.Navigation != nil {
			enhancedSteps[i].NextStep = step.Navigation.NextStep
			enhancedSteps[i].PrevStep = step.Navigation.PrevStep
		}
	}
	return enhancedSteps
}

func convertSchemaActionsToEnhancedActions(actions []ModalActionSchema) ([]EnhancedModalAction, *EnhancedModalAction, *EnhancedModalAction, *EnhancedModalAction) {
	enhancedActions := make([]EnhancedModalAction, len(actions))
	var primaryAction, secondaryAction, destructiveAction *EnhancedModalAction

	for i, action := range actions {
		enhancedAction := EnhancedModalAction{
			ID:        action.ID,
			Text:      action.Text,
			Type:      action.Type,
			Variant:   convertActionVariant(action.Variant),
			Size:      convertActionSize(action.Size),
			Icon:      action.Icon,
			Position:  action.Position,
			Priority:  action.Priority,
			AutoClose: action.Action.CloseModal,
			EnabledIf: action.EnabledIf,
			VisibleIf: action.VisibleIf,
			FullWidth: action.FullWidth,
			Class:     action.CustomClass,
		}

		// Convert action behavior
		switch action.Action.Type {
		case "submit":
			enhancedAction.HXPost = action.Action.SubmitURL
		case "api":
			if action.Action.APICall != nil {
				enhancedAction.HXPost = action.Action.APICall.URL
			}
		case "custom":
			enhancedAction.AlpineClick = action.Action.CustomHandler
		}

		// Convert confirmation
		if action.Confirmation != nil && action.Confirmation.Enabled {
			enhancedAction.ConfirmText = action.Confirmation.Message
		}

		enhancedActions[i] = enhancedAction

		// Set convenience action pointers
		switch action.Type {
		case "primary":
			primaryAction = &enhancedActions[i]
		case "secondary":
			secondaryAction = &enhancedActions[i]
		case "destructive":
			destructiveAction = &enhancedActions[i]
		}
	}

	return enhancedActions, primaryAction, secondaryAction, destructiveAction
}

func convertActionVariant(variant string) atoms.ButtonVariant {
	switch variant {
	case "primary":
		return atoms.ButtonPrimary
	case "secondary":
		return atoms.ButtonSecondary
	case "outline":
		return atoms.ButtonOutline
	case "destructive":
		return atoms.ButtonDestructive
	case "ghost":
		return atoms.ButtonGhost
	case "link":
		return atoms.ButtonLink
	default:
		return atoms.ButtonPrimary
	}
}

func convertActionSize(size string) atoms.ButtonSize {
	switch size {
	case "xs":
		return atoms.ButtonSizeXS
	case "sm":
		return atoms.ButtonSizeSM
	case "lg":
		return atoms.ButtonSizeLG
	case "xl":
		return atoms.ButtonSizeXL
	default:
		return atoms.ButtonSizeMD
	}
}

func convertBodyConfiguration(body *ModalBodySchema, props *EnhancedModalProps) {
	// Convert search configuration
	if body.SearchConfig != nil {
		props.Searchable = true
		props.SearchURL = body.SearchConfig.Endpoint
		props.MultiSelect = body.SearchConfig.SelectBehavior.MultiSelect
		props.DebounceMs = body.SearchConfig.DebounceMs
	}

	// Convert upload configuration
	if body.UploadConfig != nil {
		props.FileUpload = true
		props.UploadURL = body.UploadConfig.Endpoint
		props.UploadMultiple = body.UploadConfig.MaxFiles != 1
		props.UploadMaxFiles = body.UploadConfig.MaxFiles
		props.UploadMaxSize = body.UploadConfig.MaxFileSize
		if len(body.UploadConfig.AcceptedTypes) > 0 {
			props.UploadAccept = strings.Join(body.UploadConfig.AcceptedTypes, ",")
		}
		props.ShowProgress = body.UploadConfig.ShowProgress
	}
}

func convertValidationConfiguration(validation *ValidationSchema, props *EnhancedModalProps) {
	// Set validation strategy
	switch validation.Strategy {
	case "realtime":
		// Set realtime validation
	case "onblur":
		// Set blur validation
	case "onsubmit":
		// Set submit validation (default)
	}

	// Convert server validation
	if validation.ServerValidation != nil && validation.ServerValidation.Enabled {
		props.FormValidation = true
		// Set validation URL if available
	}
}

func extractRequiredPermissions(rules []PermissionRule) []string {
	var permissions []string
	for _, rule := range rules {
		permissions = append(permissions, rule.Required...)
	}
	return utils.RemoveDuplicates(permissions)
}

func convertThemeConfiguration(theme *ModalThemeSchema, props *EnhancedModalProps) {
	if theme.ID != "" {
		props.ThemeID = theme.ID
	}

	if len(theme.CSSVariables) > 0 {
		props.TokenOverrides = theme.CSSVariables
	}
}

// Predefined schema templates

// CreateConfirmationModalSchema creates a schema for confirmation modals
func CreateConfirmationModalSchema(id, title, message string, destructive bool) *ModalSchema {
	variant := "warning"
	if destructive {
		variant = "destructive"
	}

	builder := NewModalSchemaBuilder().
		ID(id).
		Name("confirmation").
		Title(title).
		Type("dialog").
		Size("sm").
		Variant(variant).
		Dismissible(true).
		CloseOnEscape(true).
		CloseOnOverlay(false)

	// Add header
	header := ModalHeaderSchema{
		Visible:     true,
		Title:       title,
		Description: message,
		Icon:        "alert-triangle",
	}
	builder = builder.WithHeader(header)

	// Add body
	body := ModalBodySchema{
		ContentType: "custom",
		CustomContent: []ContentBlock{
			{
				ID:      "message",
				Type:    "text",
				Content: message,
			},
		},
	}
	builder = builder.WithBody(body)

	// Add actions
	builder = builder.AddCancelAction("Cancel")
	if destructive {
		builder = builder.AddDestructiveAction("Delete", "submit")
	} else {
		builder = builder.AddPrimaryAction("Confirm", "submit")
	}

	return builder.Build()
}

// CreateFormModalSchema creates a schema for form modals
func CreateFormModalSchema(id, title string, formSchema *schema.FormSchema) *ModalSchema {
	builder := NewModalSchemaBuilder().
		ID(id).
		Name("form").
		Title(title).
		Type("dialog").
		Size("lg").
		Dismissible(true).
		CloseOnEscape(true).
		AutoSave(true, 30).
		TrapFocus(true).
		WithFormSchema(formSchema)

	// Add header
	header := ModalHeaderSchema{
		Visible: true,
		Title:   title,
	}
	builder = builder.WithHeader(header)

	// Add body with form configuration
	body := ModalBodySchema{
		ContentType: "form",
		Scrollable:  true,
		FormConfig: &FormConfig{
			InlineSchema: formSchema,
			Layout:       "vertical",
			Validation: ValidationConfig{
				Enabled:  true,
				Strategy: "onblur",
			},
			SubmitBehavior: SubmitConfig{
				ValidateFirst: true,
				SuccessAction: "close",
			},
		},
	}
	builder = builder.WithBody(body)

	// Add validation
	validation := ValidationSchema{
		Strategy: "onblur",
		ErrorDisplay: ErrorDisplayConfig{
			Strategy:    "inline",
			AutoClear:   true,
			ClearDelay:  5000,
			GroupErrors: false,
		},
	}
	builder = builder.WithValidation(validation)

	// Add actions
	builder = builder.
		AddCancelAction("Cancel").
		AddPrimaryAction("Save", "submit")

	return builder.Build()
}

// CreateSearchModalSchema creates a schema for search modals
func CreateSearchModalSchema(id, title, endpoint string, multiSelect bool) *ModalSchema {
	builder := NewModalSchemaBuilder().
		ID(id).
		Name("search").
		Title(title).
		Type("dialog").
		Size("lg").
		Dismissible(true).
		CloseOnEscape(true)

	// Add header
	header := ModalHeaderSchema{
		Visible: true,
		Title:   title,
	}
	builder = builder.WithHeader(header)

	// Add body with search configuration
	body := ModalBodySchema{
		ContentType: "search",
		Scrollable:  true,
		SearchConfig: &SearchConfig{
			Endpoint:      endpoint,
			Placeholder:   "Search...",
			MinCharacters: 2,
			DebounceMs:    300,
			MaxResults:    50,
			SelectBehavior: SelectionConfig{
				MultiSelect:      multiSelect,
				SelectionMode:    "click",
				ConfirmSelection: true,
			},
		},
	}
	builder = builder.WithBody(body)

	// Add actions
	builder = builder.
		AddCancelAction("Cancel").
		AddPrimaryAction("Select", "submit")

	return builder.Build()
}

// CreateUploadModalSchema creates a schema for file upload modals
func CreateUploadModalSchema(id, title, endpoint string, maxFiles int, acceptedTypes []string) *ModalSchema {
	builder := NewModalSchemaBuilder().
		ID(id).
		Name("upload").
		Title(title).
		Type("dialog").
		Size("lg").
		Dismissible(true).
		CloseOnEscape(true)

	// Add header
	header := ModalHeaderSchema{
		Visible: true,
		Title:   title,
	}
	builder = builder.WithHeader(header)

	// Add body with upload configuration
	body := ModalBodySchema{
		ContentType: "upload",
		UploadConfig: &UploadConfig{
			Endpoint:       endpoint,
			MaxFiles:       maxFiles,
			MaxFileSize:    10, // 10MB
			AcceptedTypes:  acceptedTypes,
			UploadMode:     "immediate",
			ShowProgress:   true,
			PreviewEnabled: true,
		},
	}
	builder = builder.WithBody(body)

	// Add actions
	builder = builder.
		AddCancelAction("Cancel").
		AddPrimaryAction("Upload", "submit")

	return builder.Build()
}

// CreateMultiStepModalSchema creates a schema for multi-step modals
func CreateMultiStepModalSchema(id, title string, steps []ModalStepSchema) *ModalSchema {
	builder := NewModalSchemaBuilder().
		ID(id).
		Name("multistep").
		Title(title).
		Type("dialog").
		Size("lg").
		Dismissible(true).
		CloseOnEscape(true).
		AutoSave(true, 30)

	// Add header with progress
	header := ModalHeaderSchema{
		Visible:      true,
		Title:        title,
		ShowProgress: true,
		ProgressType: "bar",
	}
	builder = builder.WithHeader(header)

	// Add body with step configuration
	body := ModalBodySchema{
		ContentType: "steps",
		Scrollable:  true,
		StepConfig: &StepConfig{
			ShowNavigation: true,
			LinearProgress: true,
			NavigationType: "buttons",
			StepValidation: true,
			SaveOnStep:     true,
		},
	}
	builder = builder.WithBody(body)

	// Add steps
	for _, step := range steps {
		builder = builder.AddStep(step)
	}

	// Add validation
	validation := ValidationSchema{
		Strategy: "onblur",
		ErrorDisplay: ErrorDisplayConfig{
			Strategy:    "inline",
			AutoClear:   true,
			GroupErrors: false,
		},
	}
	builder = builder.WithValidation(validation)

	// Add actions (step navigation will be handled automatically)
	builder = builder.
		AddCancelAction("Cancel").
		AddPrimaryAction("Complete", "submit")

	return builder.Build()
}

// Schema validation functions

// ValidateModalSchema validates a modal schema for correctness
func ValidateModalSchema(schema *ModalSchema) []string {
	var errors []string

	// Basic validation
	if schema.ID == "" {
		errors = append(errors, "schema ID is required")
	}

	if schema.Type == "" {
		errors = append(errors, "modal type is required")
	}

	// Validate steps if present
	if len(schema.Steps) > 0 {
		if err := validateSteps(schema.Steps); err != nil {
			errors = append(errors, fmt.Sprintf("step validation error: %s", err.Error()))
		}
	}

	// Validate actions
	if err := validateActions(schema.Actions); err != nil {
		errors = append(errors, fmt.Sprintf("action validation error: %s", err.Error()))
	}

	// Validate data sources
	if err := validateDataSources(schema.DataSources); err != nil {
		errors = append(errors, fmt.Sprintf("data source validation error: %s", err.Error()))
	}

	return errors
}

func validateSteps(steps []ModalStepSchema) error {
	stepIDs := make(map[string]bool)

	for i, step := range steps {
		if step.ID == "" {
			return fmt.Errorf("step %d: ID is required", i)
		}

		if stepIDs[step.ID] {
			return fmt.Errorf("step %d: duplicate step ID '%s'", i, step.ID)
		}
		stepIDs[step.ID] = true

		if step.Title == "" {
			return fmt.Errorf("step %d: title is required", i)
		}

		if step.ContentType == "" {
			return fmt.Errorf("step %d: content type is required", i)
		}
	}

	return nil
}

func validateActions(actions []ModalActionSchema) error {
	actionIDs := make(map[string]bool)

	for i, action := range actions {
		if action.ID == "" {
			return fmt.Errorf("action %d: ID is required", i)
		}

		if actionIDs[action.ID] {
			return fmt.Errorf("action %d: duplicate action ID '%s'", i, action.ID)
		}
		actionIDs[action.ID] = true

		if action.Text == "" {
			return fmt.Errorf("action %d: text is required", i)
		}

		if action.Action.Type == "" {
			return fmt.Errorf("action %d: action type is required", i)
		}
	}

	return nil
}

func validateDataSources(sources []DataSourceSchema) error {
	sourceIDs := make(map[string]bool)

	for i, source := range sources {
		if source.ID == "" {
			return fmt.Errorf("data source %d: ID is required", i)
		}

		if sourceIDs[source.ID] {
			return fmt.Errorf("data source %d: duplicate source ID '%s'", i, source.ID)
		}
		sourceIDs[source.ID] = true

		if source.Type == "" {
			return fmt.Errorf("data source %d: type is required", i)
		}

		// Validate type-specific requirements
		switch source.Type {
		case "api":
			if source.API == nil || source.API.URL == "" {
				return fmt.Errorf("data source %d: API URL is required for API type", i)
			}
		case "static":
			if source.StaticData == nil {
				return fmt.Errorf("data source %d: static data is required for static type", i)
			}
		case "function":
			if source.Function == "" {
				return fmt.Errorf("data source %d: function name is required for function type", i)
			}
		case "storage":
			if source.StorageKey == "" {
				return fmt.Errorf("data source %d: storage key is required for storage type", i)
			}
		}
	}

	return nil
}

// Schema serialization functions

// ToJSON serializes the modal schema to JSON
func (s *ModalSchema) ToJSON() ([]byte, error) {
	return json.MarshalIndent(s, "", "  ")
}

// FromJSON deserializes the modal schema from JSON
func (s *ModalSchema) FromJSON(data []byte) error {
	return json.Unmarshal(data, s)
}

// ToJSONString serializes the modal schema to JSON string
func (s *ModalSchema) ToJSONString() (string, error) {
	data, err := s.ToJSON()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSONString deserializes the modal schema from JSON string
func (s *ModalSchema) FromJSONString(jsonStr string) error {
	return s.FromJSON([]byte(jsonStr))
}
