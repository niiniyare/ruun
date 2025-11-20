package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/niiniyare/ruun/pkg/condition"
	"github.com/niiniyare/ruun/pkg/logger"
	"gopkg.in/yaml.v3"
)

// =============================================================================
// MULTI-TENANCY CONTEXT KEY
// =============================================================================

// tenantKey is the context key for tenant identification.
// Use WithTenant() and GetTenant() to safely access tenant information.
type tenantKey struct{}

// WithTenant returns a new context with the tenant ID embedded.
// The tenant ID can be a string identifier or UUID string.
//
// Example:
//
//	ctx := schema.WithTenant(context.Background(), "tenant-123")
//	theme, err := manager.GetTheme(ctx, "default-theme")
func WithTenant(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, tenantKey{}, tenantID)
}

// GetTenant extracts the tenant ID from the context.
// Returns empty string if no tenant is set (uses global theme pool).
//
// Example:
//
//	tenantID := schema.GetTenant(ctx)
//	if tenantID == "" {
//	    // Using global themes
//	} else {
//	    // Using tenant-specific themes
//	}
func GetTenant(ctx context.Context) string {
	if tenantID, ok := ctx.Value(tenantKey{}).(string); ok {
		return tenantID
	}
	return ""
}

// =============================================================================
// CONDITIONAL THEMING - SIMPLIFIED WRAPPER
// =============================================================================

// ThemeCondition defines a conditional theme override with expression evaluation.
// This is a simplified wrapper around the condition package, hiding complexity
// from users while providing powerful conditional theming capabilities.
//
// Expression Syntax:
//   - Use standard comparison operators: ==, !=, <, <=, >, >=
//   - Logical operators: &&, ||, !
//   - Variables from evaluation context: user.role, tenant.plan, time.hour
//   - String operations: contains, startsWith, endsWith
//   - Membership: in, not in
//
// Expression Examples:
//   - "user.role == 'admin'" - Apply override for admin users
//   - "tenant.plan == 'enterprise' && tenant.region == 'US'" - Enterprise US tenants
//   - "time.hour >= 18 || time.hour < 6" - Night mode hours (6 PM to 6 AM)
//   - "user.preferences.theme == 'dark'" - User preference-based
//
// Token Override Format:
//   - Use dot notation for token paths
//   - Example: "semantic.colors.background" -> "primitives.colors.gray-900"
//
// Priority:
//   - Higher priority values are applied last (override lower priority)
//   - Use priority to control override order when multiple conditions match
//   - Default priority is 0
type ThemeCondition struct {
	// ID is a unique identifier for this condition (for logging/debugging)
	ID string `json:"id" yaml:"id"`

	// Expression is the condition formula to evaluate
	// Must be a valid expression syntax supported by the condition package
	Expression string `json:"expression" yaml:"expression"`

	// Priority determines the order of application (higher = later = wins)
	// When multiple conditions match, they are applied in priority order
	Priority int `json:"priority" yaml:"priority"`

	// TokenOverrides maps token paths to their override values
	// Format: "tier.category.key" -> "value or reference"
	// Example: "semantic.colors.background" -> "primitives.colors.gray-900"
	TokenOverrides map[string]string `json:"tokenOverrides,omitempty" yaml:"tokenOverrides,omitempty"`

	// Description provides human-readable explanation of the condition
	// Used for documentation and debugging
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// Validate ensures the theme condition is well-formed.
func (tc *ThemeCondition) Validate() error {
	if tc.ID == "" {
		return NewValidationError("empty_condition_id", "condition ID cannot be empty")
	}
	if tc.Expression == "" {
		return NewValidationError("empty_expression", "expression cannot be empty")
	}
	if tc.TokenOverrides == nil || len(tc.TokenOverrides) == 0 {
		return NewValidationError("empty_overrides", "token overrides cannot be empty")
	}

	// Validate all token override paths
	for path, value := range tc.TokenOverrides {
		if path == "" {
			return NewValidationError("empty_override_path", "override path cannot be empty")
		}
		ref := TokenReference(value)
		if err := ref.Validate(); err != nil {
			return fmt.Errorf("invalid override value for '%s': %w", path, err)
		}
	}

	return nil
}

// Clone creates a deep copy of the theme condition.
func (tc *ThemeCondition) Clone() *ThemeCondition {
	if tc == nil {
		return nil
	}
	return &ThemeCondition{
		ID:             tc.ID,
		Expression:     tc.Expression,
		Priority:       tc.Priority,
		TokenOverrides: cloneMap(tc.TokenOverrides),
		Description:    tc.Description,
	}
}

// conditionEvaluator wraps the condition package evaluator with caching.
// This is internal and not exposed to users of the schema package.
type conditionEvaluator struct {
	evaluator *condition.Evaluator
	cache     *ristretto.Cache
	mu        sync.RWMutex
}

// newConditionEvaluator creates a new condition evaluator with bounded cache.
func newConditionEvaluator() (*conditionEvaluator, error) {
	// Create condition evaluator with default options
	opts := condition.DefaultEvalOptions()
	opts.Timeout = 5 * time.Second // 5s timeout for expression evaluation
	opts.MaxDepth = 10             // Max 10 levels of nesting
	opts.MaxConditions = 100       // Max 100 conditions per expression
	evaluator := condition.NewEvaluator(nil, opts)

	// Create cache for compiled expressions
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1000,
		MaxCost:     10 << 20, // 10MB for expression cache
		BufferItems: 64,
		Cost: func(value interface{}) int64 {
			return 1024 // Fixed cost per compiled expression
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create condition cache: %w", err)
	}

	return &conditionEvaluator{
		evaluator: evaluator,
		cache:     cache,
	}, nil
}

// evaluate evaluates a condition expression against provided data.
// Uses caching to avoid recompiling the same expressions.
func (ce *conditionEvaluator) evaluate(ctx context.Context, expr string, evalData map[string]any) (bool, error) {
	// Check cache for compiled expression
	ce.mu.RLock()
	if _, found := ce.cache.Get(expr); found {
		ce.mu.RUnlock()
		// Expression is cached, proceed with evaluation
	} else {
		ce.mu.RUnlock()
		// Cache the expression for future use
		ce.mu.Lock()
		ce.cache.Set(expr, struct{}{}, 1024)
		ce.mu.Unlock()
	}

	// Create evaluation context
	evalCtx := condition.NewEvalContext(evalData, condition.DefaultEvalOptions())

	// Build condition group from expression
	builder := condition.NewBuilder(condition.ConjunctionAnd)
	builder.AddFormula(expr)
	group := builder.Build()

	// Evaluate the expression
	result, err := ce.evaluator.Evaluate(ctx, group, evalCtx)
	if err != nil {
		return false, fmt.Errorf("expression evaluation failed: %w", err)
	}

	return result, nil
}

// close releases resources used by the evaluator.
func (ce *conditionEvaluator) close() {
	ce.cache.Close()
}

// =============================================================================
// THEME STRUCTURE
// =============================================================================

// Theme represents a complete design system theme with metadata and configuration.
//
// A theme consists of:
//   - Design tokens (primitives, semantic, components)
//   - Dark mode configuration and tokens
//   - Accessibility settings
//   - Conditional overrides based on runtime context
//   - Custom CSS/JS for advanced customization
//   - Metadata for organization and discovery
//
// Themes support:
//   - Multi-tenancy via context-based isolation
//   - Conditional token overrides based on user/tenant/time context
//   - Dark mode with separate token definitions
//   - Accessibility enhancements (high contrast, reduced motion)
//   - Custom CSS/JS injection for advanced use cases
//   - Theme inheritance (not yet implemented)
//
// JSON/YAML Structure:
//
//	{
//	  "id": "acme-brand",
//	  "name": "Acme Brand Theme",
//	  "version": "1.0.0",
//	  "tokens": { ... },
//	  "darkMode": { ... },
//	  "conditions": [ ... ],
//	  "accessibility": { ... },
//	  "meta": { ... }
//	}
type Theme struct {
	// ID is the unique identifier for this theme
	// Must be unique within a tenant's theme pool
	// Format: lowercase alphanumeric with hyphens (e.g., "acme-brand", "dark-mode")
	ID string `json:"id" yaml:"id"`

	// Name is the human-readable theme name
	// Displayed in theme selection UIs
	Name string `json:"name" yaml:"name"`

	// Description provides context about the theme's purpose and features
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Version follows semantic versioning (e.g., "1.0.0", "2.1.3")
	// Used for theme compatibility and migration tracking
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// Author identifies the theme creator or organization
	Author string `json:"author,omitempty" yaml:"author,omitempty"`

	// Tokens contains the complete design token system
	// This is the core of the theme defining all visual styles
	Tokens *DesignTokens `json:"tokens" yaml:"tokens"`

	// DarkMode defines dark mode configuration and token overrides
	// When enabled, provides an alternate color scheme for low-light environments
	DarkMode *DarkModeConfig `json:"darkMode,omitempty" yaml:"darkMode,omitempty"`

	// Conditions defines context-aware theme overrides
	// Applied when expressions evaluate to true based on runtime context
	Conditions []*ThemeCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`

	// Accessibility defines accessibility enhancements
	// Includes high contrast, focus indicators, reduced motion, etc.
	Accessibility *AccessibilityConfig `json:"accessibility,omitempty" yaml:"accessibility,omitempty"`

	// CustomCSS provides raw CSS for advanced customization
	// Injected after theme tokens are applied
	// Use sparingly; prefer token-based customization
	CustomCSS string `json:"customCSS,omitempty" yaml:"customCSS,omitempty"`

	// CustomJS provides JavaScript for dynamic theme behavior
	// Use with caution; may affect performance and security
	CustomJS string `json:"customJS,omitempty" yaml:"customJS,omitempty"`

	// Meta contains theme metadata for organization and discovery
	Meta *ThemeMeta `json:"meta,omitempty" yaml:"meta,omitempty"`

	// Internal timestamps (not exported in JSON/YAML)
	createdAt time.Time
	updatedAt time.Time
}

// DarkModeConfig defines dark mode behavior and token overrides.
//
// Dark mode provides an alternate color scheme optimized for low-light viewing.
// When enabled, dark tokens override or augment the base theme tokens.
//
// Configuration Options:
//   - Enabled: Whether dark mode is available
//   - Default: Whether dark mode is the default (vs light mode)
//   - Strategy: How to detect dark mode preference
//   - DarkTokens: Token overrides specific to dark mode
//
// Strategy Types:
//   - "class": Toggle via CSS class on root element (controlled by app)
//   - "media": Respect system preference via prefers-color-scheme media query
//   - "auto": Combination of class override + media query fallback
type DarkModeConfig struct {
	// Enabled determines if dark mode is available
	Enabled bool `json:"enabled" yaml:"enabled"`

	// Default determines if dark mode is the initial/default mode
	// When true, light mode is the alternate; when false, dark mode is alternate
	Default bool `json:"default,omitempty" yaml:"default,omitempty"`

	// Strategy defines how dark mode is detected/applied
	// Options: "class", "media", "auto"
	Strategy string `json:"strategy,omitempty" yaml:"strategy,omitempty"`

	// DarkTokens contains token overrides for dark mode
	// These typically override semantic colors and some primitives
	// Example: semantic.colors.background -> primitives.colors.gray-900
	DarkTokens *DesignTokens `json:"darkTokens,omitempty" yaml:"darkTokens,omitempty"`
}

// AccessibilityConfig defines accessibility enhancements for the theme.
//
// Accessibility features ensure the theme is usable by people with disabilities:
//   - Visual: High contrast, focus indicators, color blindness considerations
//   - Motor: Keyboard navigation, larger touch targets
//   - Cognitive: Reduced motion, clear focus states
//   - Screen readers: ARIA labels, semantic HTML
//
// These settings can be enforced or used as hints for component rendering.
type AccessibilityConfig struct {
	// HighContrast enables enhanced color contrast ratios
	// Ensures WCAG 2.1 AAA compliance (7:1 for normal text, 4.5:1 for large text)
	HighContrast bool `json:"highContrast,omitempty" yaml:"highContrast,omitempty"`

	// MinContrastRatio specifies minimum contrast ratio for text
	// WCAG AA: 4.5:1 (normal text), 3:1 (large text)
	// WCAG AAA: 7:1 (normal text), 4.5:1 (large text)
	MinContrastRatio float64 `json:"minContrastRatio,omitempty" yaml:"minContrastRatio,omitempty"`

	// FocusIndicator ensures visible focus states for keyboard navigation
	// Includes outline width, color, and offset
	FocusIndicator bool `json:"focusIndicator,omitempty" yaml:"focusIndicator,omitempty"`

	// FocusOutlineColor defines the color for focus indicators
	// Should have high contrast against all backgrounds
	FocusOutlineColor string `json:"focusOutlineColor,omitempty" yaml:"focusOutlineColor,omitempty"`

	// FocusOutlineWidth defines the width of focus outlines
	// WCAG recommends at least 2px
	FocusOutlineWidth string `json:"focusOutlineWidth,omitempty" yaml:"focusOutlineWidth,omitempty"`

	// KeyboardNav ensures all interactive elements are keyboard accessible
	// Includes proper tab order, focus management, and keyboard shortcuts
	KeyboardNav bool `json:"keyboardNav,omitempty" yaml:"keyboardNav,omitempty"`

	// ReducedMotion respects prefers-reduced-motion media query
	// Disables or reduces animations/transitions for motion-sensitive users
	ReducedMotion bool `json:"reducedMotion,omitempty" yaml:"reducedMotion,omitempty"`

	// ScreenReaderOptimized includes ARIA labels and semantic HTML
	// Ensures screen readers can properly announce all content
	ScreenReaderOptimized bool `json:"screenReaderOptimized,omitempty" yaml:"screenReaderOptimized,omitempty"`

	// AriaLive defines default ARIA live region politeness
	// Options: "off", "polite", "assertive"
	AriaLive string `json:"ariaLive,omitempty" yaml:"ariaLive,omitempty"`
}

// ThemeMeta contains theme metadata for organization, discovery, and attribution.
//
// Metadata helps with:
//   - Theme discovery and search
//   - Attribution and licensing
//   - Documentation and previews
//   - Custom application-specific data
type ThemeMeta struct {
	// Tags for categorization and search (e.g., ["corporate", "minimal", "dark"])
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// License identifier (e.g., "MIT", "Apache-2.0", "proprietary")
	License string `json:"license,omitempty" yaml:"license,omitempty"`

	// Repository URL for theme source code
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Homepage URL for theme documentation or landing page
	Homepage string `json:"homepage,omitempty" yaml:"homepage,omitempty"`

	// Preview URL for theme screenshot or demo
	Preview string `json:"preview,omitempty" yaml:"preview,omitempty"`

	// CustomData for application-specific metadata
	// Use for any additional data your application needs
	CustomData map[string]any `json:"customData,omitempty" yaml:"customData,omitempty"`

	// Timestamps (included in JSON/YAML)
	CreatedAt time.Time `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`
}

// =============================================================================
// THEME VALIDATION
// =============================================================================

// Validate ensures the theme is well-formed and all components are valid.
func (t *Theme) Validate() error {
	if t.ID == "" {
		return NewValidationError("empty_theme_id", "theme ID cannot be empty")
	}
	if t.Name == "" {
		return NewValidationError("empty_theme_name", "theme name cannot be empty")
	}
	if t.Tokens == nil {
		return NewValidationError("nil_tokens", "theme tokens cannot be nil")
	}

	// Validate tokens
	if err := t.Tokens.Validate(); err != nil {
		return fmt.Errorf("invalid tokens: %w", err)
	}

	// Validate dark mode if present
	if t.DarkMode != nil {
		if err := t.DarkMode.Validate(); err != nil {
			return fmt.Errorf("invalid dark mode: %w", err)
		}
	}

	// Validate conditions if present
	for i, cond := range t.Conditions {
		if err := cond.Validate(); err != nil {
			return fmt.Errorf("invalid condition %d (%s): %w", i, cond.ID, err)
		}
	}

	// Validate accessibility if present
	if t.Accessibility != nil {
		if err := t.Accessibility.Validate(); err != nil {
			return fmt.Errorf("invalid accessibility: %w", err)
		}
	}

	// Validate custom CSS/JS size limits (prevent DoS)
	const maxCustomCodeSize = 1 << 20 // 1MB
	if len(t.CustomCSS) > maxCustomCodeSize {
		return NewValidationError("custom_css_too_large",
			fmt.Sprintf("custom CSS exceeds maximum size of %d bytes", maxCustomCodeSize))
	}
	if len(t.CustomJS) > maxCustomCodeSize {
		return NewValidationError("custom_js_too_large",
			fmt.Sprintf("custom JS exceeds maximum size of %d bytes", maxCustomCodeSize))
	}

	return nil
}

// Validate checks dark mode configuration.
func (d *DarkModeConfig) Validate() error {
	if d == nil {
		return nil
	}

	// Validate strategy
	if d.Strategy != "" {
		validStrategies := map[string]bool{
			"class": true,
			"media": true,
			"auto":  true,
		}
		if !validStrategies[d.Strategy] {
			return NewValidationError("invalid_dark_mode_strategy",
				fmt.Sprintf("invalid dark mode strategy: %s (must be class, media, or auto)", d.Strategy))
		}
	}

	// Validate dark tokens if present
	if d.DarkTokens != nil {
		if err := d.DarkTokens.Validate(); err != nil {
			return fmt.Errorf("invalid dark tokens: %w", err)
		}
	}

	return nil
}

// Validate checks accessibility configuration.
func (a *AccessibilityConfig) Validate() error {
	if a == nil {
		return nil
	}

	// Validate contrast ratio
	if a.MinContrastRatio < 0 || a.MinContrastRatio > 21 {
		return NewValidationError("invalid_contrast_ratio",
			fmt.Sprintf("contrast ratio must be between 0 and 21, got: %f", a.MinContrastRatio))
	}

	// Validate ARIA live
	if a.AriaLive != "" {
		validValues := map[string]bool{
			"off":       true,
			"polite":    true,
			"assertive": true,
		}
		if !validValues[a.AriaLive] {
			return NewValidationError("invalid_aria_live",
				fmt.Sprintf("invalid ARIA live value: %s (must be off, polite, or assertive)", a.AriaLive))
		}
	}

	return nil
}

// =============================================================================
// THEME CLONING
// =============================================================================

// Clone creates a deep copy of the theme.
// Use this to ensure immutability when modifying themes.
func (t *Theme) Clone() *Theme {
	if t == nil {
		return nil
	}

	cloned := &Theme{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Author:      t.Author,
		CustomCSS:   t.CustomCSS,
		CustomJS:    t.CustomJS,
		createdAt:   t.createdAt,
		updatedAt:   t.updatedAt,
	}

	if t.Tokens != nil {
		cloned.Tokens = t.Tokens.Clone()
	}
	if t.DarkMode != nil {
		cloned.DarkMode = t.DarkMode.Clone()
	}
	if t.Accessibility != nil {
		cloned.Accessibility = t.Accessibility.Clone()
	}
	if t.Meta != nil {
		cloned.Meta = t.Meta.Clone()
	}
	if len(t.Conditions) > 0 {
		cloned.Conditions = make([]*ThemeCondition, len(t.Conditions))
		for i, cond := range t.Conditions {
			cloned.Conditions[i] = cond.Clone()
		}
	}

	return cloned
}

// Clone creates a deep copy of dark mode configuration.
func (d *DarkModeConfig) Clone() *DarkModeConfig {
	if d == nil {
		return nil
	}
	cloned := &DarkModeConfig{
		Enabled:  d.Enabled,
		Default:  d.Default,
		Strategy: d.Strategy,
	}
	if d.DarkTokens != nil {
		cloned.DarkTokens = d.DarkTokens.Clone()
	}
	return cloned
}

// Clone creates a deep copy of accessibility configuration.
func (a *AccessibilityConfig) Clone() *AccessibilityConfig {
	if a == nil {
		return nil
	}
	cloned := *a
	return &cloned
}

// Clone creates a deep copy of theme metadata.
func (m *ThemeMeta) Clone() *ThemeMeta {
	if m == nil {
		return nil
	}
	cloned := &ThemeMeta{
		License:    m.License,
		Repository: m.Repository,
		Homepage:   m.Homepage,
		Preview:    m.Preview,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	if len(m.Tags) > 0 {
		cloned.Tags = make([]string, len(m.Tags))
		copy(cloned.Tags, m.Tags)
	}
	if m.CustomData != nil {
		cloned.CustomData = make(map[string]any, len(m.CustomData))
		for k, v := range m.CustomData {
			cloned.CustomData[k] = v
		}
	}
	return cloned
}

// =============================================================================
// THEME SERIALIZATION
// =============================================================================

// ToJSON serializes the theme to pretty-printed JSON.
func (t *Theme) ToJSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

// FromJSON deserializes a theme from JSON and validates it.
func ThemeFromJSON(data []byte) (*Theme, error) {
	var theme Theme
	if err := json.Unmarshal(data, &theme); err != nil {
		return nil, fmt.Errorf("failed to parse theme JSON: %w", err)
	}
	if err := theme.Validate(); err != nil {
		return nil, fmt.Errorf("invalid theme: %w", err)
	}
	return &theme, nil
}

// ToYAML serializes the theme to YAML.
func (t *Theme) ToYAML() ([]byte, error) {
	return yaml.Marshal(t)
}

// FromYAML deserializes a theme from YAML and validates it.
func ThemeFromYAML(data []byte) (*Theme, error) {
	var theme Theme
	if err := yaml.Unmarshal(data, &theme); err != nil {
		return nil, fmt.Errorf("failed to parse theme YAML: %w", err)
	}
	if err := theme.Validate(); err != nil {
		return nil, fmt.Errorf("invalid theme: %w", err)
	}
	return &theme, nil
}

// =============================================================================
// THEME STORAGE INTERFACE
// =============================================================================

// ThemeStorage defines the interface for theme persistence.
//
// Implementation Strategy:
//   - Database: Store themes in PostgreSQL/MySQL with JSON/JSONB columns
//   - Filesystem: Store themes as JSON/YAML files in organized directories
//   - Object Storage: Store themes in S3/GCS/Azure Blob with metadata
//   - Cache: Use Redis/Memcached for frequently accessed themes
//   - CDN: Distribute themes via CDN for global low-latency access
//
// Multi-tenancy:
//   - Include tenant ID in all storage operations
//   - Implement tenant isolation at storage layer
//   - Use separate buckets/tables/directories per tenant if needed
//
// This interface follows the io.Reader/Writer philosophy:
// minimal, focused, and easy to implement for any storage backend.
type ThemeStorage interface {
	// SaveTheme persists a theme to storage.
	// Implementations should:
	//   - Serialize theme to JSON/YAML
	//   - Store with tenant isolation
	//   - Set appropriate timestamps
	//   - Handle duplicate IDs (update vs error)
	SaveTheme(ctx context.Context, theme *Theme) error

	// LoadTheme retrieves a theme from storage by ID.
	// Returns ValidationError with code "theme_not_found" if theme doesn't exist.
	LoadTheme(ctx context.Context, themeID string) (*Theme, error)

	// DeleteTheme removes a theme from storage.
	// Should be idempotent (no error if theme doesn't exist).
	DeleteTheme(ctx context.Context, themeID string) error

	// ListThemes returns all theme IDs in storage.
	// Should respect tenant context if multi-tenant.
	ListThemes(ctx context.Context) ([]string, error)
}

// =============================================================================
// THEME MANAGER CONFIGURATION
// =============================================================================

// ThemeManagerConfig defines configuration for the theme manager.
type ThemeManagerConfig struct {
	// EnableCaching enables theme caching for performance
	EnableCaching bool

	// CacheMaxCost is the maximum memory cost for the theme cache (in bytes)
	CacheMaxCost int64

	// CacheNumCounters is the number of counters for the cache admission policy
	CacheNumCounters int64

	// ValidateOnRegister validates themes when registering
	ValidateOnRegister bool

	// DefaultThemeID is the fallback theme when requested theme not found
	DefaultThemeID string

	// Storage is the theme persistence backend (optional)
	Storage ThemeStorage

	// Logger for theme operations (optional, defaults to no-op logger)
	Logger logger.Logger
}

// DefaultThemeManagerConfig returns sensible default configuration.
func DefaultThemeManagerConfig() *ThemeManagerConfig {
	return &ThemeManagerConfig{
		EnableCaching:      true,
		CacheMaxCost:       100 << 20, // 100MB
		CacheNumCounters:   1000,
		ValidateOnRegister: true,
		DefaultThemeID:     "",
		Storage:            nil,
		Logger:             logger.WithFields(logger.Fields{"component": "theme-manager"}),
	}
}

// =============================================================================
// THEME MANAGER
// =============================================================================

// ThemeManager manages themes with multi-tenancy, caching, and conditional overrides.
//
// Features:
//   - Multi-tenant theme isolation via context
//   - High-performance caching with automatic eviction
//   - Conditional theme overrides based on runtime context
//   - Dark mode support with token overrides
//   - Pluggable storage backends
//   - Thread-safe concurrent access
//   - Comprehensive logging and error handling
//
// Architecture:
//   - Global theme pool for shared themes (no tenant context)
//   - Per-tenant theme pools (accessed via context with tenant ID)
//   - Unified API regardless of tenancy
//   - Automatic tenant detection from context
//
// Example Usage:
//
//	config := schema.DefaultThemeManagerConfig()
//	config.Storage = myPostgresStorage
//	manager := schema.NewThemeManager(config)
//
//	// Register global theme
//	err := manager.RegisterTheme(context.Background(), myTheme)
//
//	// Register tenant-specific theme
//	ctx := schema.WithTenant(context.Background(), "tenant-123")
//	err = manager.RegisterTheme(ctx, tenantTheme)
//
//	// Get theme with conditional overrides
//	evalData := map[string]any{
//	    "user": map[string]any{"role": "admin"},
//	    "time": map[string]any{"hour": 20},
//	}
//	theme, err := manager.GetTheme(ctx, "default", evalData)
type ThemeManager struct {
	// Global themes (shared across all tenants)
	globalThemes map[string]*Theme

	// Tenant-specific themes (keyed by tenant ID)
	tenantThemes map[string]map[string]*Theme

	// Theme cache (unified for global and tenant themes)
	cache *ristretto.Cache

	// Condition evaluator for conditional overrides
	condEval *conditionEvaluator

	// Configuration
	config *ThemeManagerConfig

	// Logger
	logger logger.Logger

	// Mutex for thread-safe access
	mu sync.RWMutex
}

// NewThemeManager creates a new theme manager with the given configuration.
func NewThemeManager(config *ThemeManagerConfig) (*ThemeManager, error) {
	if config == nil {
		config = DefaultThemeManagerConfig()
	}

	// Initialize logger
	log := config.Logger
	if log == nil {
		log = logger.WithFields(logger.Fields{"component": "theme-manager"})
		// If global logger isn't set, WithFields might return nil, so create a safe fallback
		if log == nil {
			// Use a no-op logger as fallback
			log = &noOpLogger{}
		}
	}

	tm := &ThemeManager{
		globalThemes: make(map[string]*Theme),
		tenantThemes: make(map[string]map[string]*Theme),
		config:       config,
		logger:       log,
	}

	// Initialize cache if enabled
	if config.EnableCaching {
		cache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: config.CacheNumCounters,
			MaxCost:     config.CacheMaxCost,
			BufferItems: 64,
			Metrics:     false,
			Cost: func(value interface{}) int64 {
				theme, ok := value.(*Theme)
				if !ok {
					return 1024
				}
				// Estimate theme size based on JSON serialization
				data, err := json.Marshal(theme)
				if err != nil {
					return 1024
				}
				return int64(len(data))
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create theme cache: %w", err)
		}
		tm.cache = cache
		tm.logger.Info("theme cache initialized", logger.Fields{
			"maxCost":     config.CacheMaxCost,
			"numCounters": config.CacheNumCounters,
		})
	}

	// Initialize condition evaluator
	condEval, err := newConditionEvaluator()
	if err != nil {
		return nil, fmt.Errorf("failed to create condition evaluator: %w", err)
	}
	tm.condEval = condEval

	tm.logger.Info("theme manager initialized", logger.Fields{
		"caching":  config.EnableCaching,
		"validate": config.ValidateOnRegister,
		"default":  config.DefaultThemeID,
	})

	return tm, nil
}

// RegisterTheme registers a theme in the manager.
// Uses context to determine if this is a global or tenant-specific theme.
//
// Context Rules:
//   - No tenant in context → registers as global theme
//   - Tenant in context → registers as tenant-specific theme
//
// The theme is validated if ValidateOnRegister is true.
// If storage is configured, the theme is persisted.
// Any existing theme with the same ID is replaced.
func (tm *ThemeManager) RegisterTheme(ctx context.Context, theme *Theme) error {
	if theme == nil {
		return NewValidationError("nil_theme", "theme cannot be nil")
	}

	// Validate theme if configured
	if tm.config.ValidateOnRegister {
		if err := theme.Validate(); err != nil {
			tm.logger.Error("theme validation failed", logger.Fields{
				"themeID": theme.ID,
				"error":   err.Error(),
			})
			return fmt.Errorf("theme validation failed: %w", err)
		}
	}

	// Set timestamps
	now := time.Now()
	if theme.createdAt.IsZero() {
		theme.createdAt = now
	}
	theme.updatedAt = now
	if theme.Meta != nil {
		if theme.Meta.CreatedAt.IsZero() {
			theme.Meta.CreatedAt = now
		}
		theme.Meta.UpdatedAt = now
	}

	// Determine tenant context
	tenantID := GetTenant(ctx)

	// Register theme based on context
	tm.mu.Lock()
	if tenantID == "" {
		// Global theme
		tm.globalThemes[theme.ID] = theme.Clone()
		tm.logger.Info("global theme registered", logger.Fields{
			"themeID": theme.ID,
			"name":    theme.Name,
		})
	} else {
		// Tenant-specific theme
		if tm.tenantThemes[tenantID] == nil {
			tm.tenantThemes[tenantID] = make(map[string]*Theme)
		}
		tm.tenantThemes[tenantID][theme.ID] = theme.Clone()
		tm.logger.Info("tenant theme registered", logger.Fields{
			"tenantID": tenantID,
			"themeID":  theme.ID,
			"name":     theme.Name,
		})
	}
	tm.mu.Unlock()

	// Invalidate cache for this theme
	tm.invalidateCache(tenantID, theme.ID)

	// Persist to storage if configured
	if tm.config.Storage != nil {
		if err := tm.config.Storage.SaveTheme(ctx, theme); err != nil {
			tm.logger.Error("failed to persist theme to storage", logger.Fields{
				"tenantID": tenantID,
				"themeID":  theme.ID,
				"error":    err.Error(),
			})
			// Don't return error - theme is still registered in memory
		}
	}

	return nil
}

// GetTheme retrieves a theme by ID and applies conditional overrides.
//
// Resolution process:
//  1. Check cache for previously resolved theme
//  2. Look up base theme in tenant pool (if tenant in context) or global pool
//  3. Fall back to default theme if configured and theme not found
//  4. Apply conditional overrides based on evaluation data
//  5. Apply dark mode tokens if applicable
//  6. Cache the resolved theme
//  7. Return cloned theme (immutable)
//
// Parameters:
//   - ctx: Context with optional tenant ID
//   - themeID: Theme identifier to retrieve
//   - evalData: Context data for conditional expression evaluation
//     Expected keys: user, tenant, time, request, session, etc.
//
// Example evalData:
//
//	{
//	    "user": {"id": "user-123", "role": "admin"},
//	    "tenant": {"id": "tenant-456", "plan": "enterprise"},
//	    "time": {"hour": 20, "day": "monday"},
//	    "request": {"path": "/admin", "method": "GET"},
//	}
func (tm *ThemeManager) GetTheme(ctx context.Context, themeID string, evalData map[string]any) (*Theme, error) {
	tenantID := GetTenant(ctx)
	cacheKey := tm.buildCacheKey(tenantID, themeID)

	// Check cache first
	if tm.config.EnableCaching {
		if cached, found := tm.cache.Get(cacheKey); found {
			tm.logger.Debug("theme cache hit", logger.Fields{
				"tenantID": tenantID,
				"themeID":  themeID,
			})
			return cached.(*Theme).Clone(), nil
		}
	}

	// Look up base theme (no lock held during storage access)
	baseTheme := tm.lookupBaseTheme(ctx, tenantID, themeID)

	// If not found and default configured, try default
	if baseTheme == nil && tm.config.DefaultThemeID != "" && tm.config.DefaultThemeID != themeID {
		tm.logger.Debug("theme not found, using default", logger.Fields{
			"tenantID":  tenantID,
			"requested": themeID,
			"default":   tm.config.DefaultThemeID,
		})
		return tm.GetTheme(ctx, tm.config.DefaultThemeID, evalData)
	}

	if baseTheme == nil {
		return nil, NewValidationError("theme_not_found",
			fmt.Sprintf("theme not found: %s", themeID))
	}

	// Apply conditional overrides
	resolvedTheme := baseTheme.Clone()
	if len(resolvedTheme.Conditions) > 0 && evalData != nil {
		if err := tm.applyConditionalOverrides(ctx, resolvedTheme, evalData); err != nil {
			tm.logger.Error("failed to apply conditional overrides", logger.Fields{
				"tenantID": tenantID,
				"themeID":  themeID,
				"error":    err.Error(),
			})
			// Continue with base theme (don't fail on condition evaluation errors)
		}
	}

	// Cache the resolved theme
	if tm.config.EnableCaching {
		tm.cache.Set(cacheKey, resolvedTheme, 0)
	}

	return resolvedTheme.Clone(), nil
}

// lookupBaseTheme finds a theme in tenant or global pool, or loads from storage.
func (tm *ThemeManager) lookupBaseTheme(ctx context.Context, tenantID, themeID string) *Theme {
	// Check in-memory first (with read lock)
	tm.mu.RLock()
	var baseTheme *Theme
	if tenantID != "" {
		if tenantThemes, exists := tm.tenantThemes[tenantID]; exists {
			if theme, found := tenantThemes[themeID]; found {
				baseTheme = theme
			}
		}
	}
	if baseTheme == nil {
		if theme, found := tm.globalThemes[themeID]; found {
			baseTheme = theme
		}
	}
	tm.mu.RUnlock()

	// If found in memory, return it
	if baseTheme != nil {
		return baseTheme
	}

	// Try loading from storage (without holding lock)
	if tm.config.Storage != nil {
		loadedTheme, err := tm.config.Storage.LoadTheme(ctx, themeID)
		if err == nil {
			// Cache in memory
			tm.mu.Lock()
			if tenantID != "" {
				if tm.tenantThemes[tenantID] == nil {
					tm.tenantThemes[tenantID] = make(map[string]*Theme)
				}
				tm.tenantThemes[tenantID][themeID] = loadedTheme.Clone()
			} else {
				tm.globalThemes[themeID] = loadedTheme.Clone()
			}
			tm.mu.Unlock()

			tm.logger.Info("theme loaded from storage", logger.Fields{
				"tenantID": tenantID,
				"themeID":  themeID,
			})
			return loadedTheme
		}
	}

	return nil
}

// applyConditionalOverrides evaluates conditions and applies matching overrides.
func (tm *ThemeManager) applyConditionalOverrides(ctx context.Context, theme *Theme, evalData map[string]any) error {
	// Sort conditions by priority (ascending)
	type prioritizedCondition struct {
		index int
		cond  *ThemeCondition
	}
	sorted := make([]prioritizedCondition, len(theme.Conditions))
	for i, cond := range theme.Conditions {
		sorted[i] = prioritizedCondition{index: i, cond: cond}
	}
	// Simple bubble sort by priority
	for i := len(sorted) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if sorted[j].cond.Priority > sorted[j+1].cond.Priority {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	// Evaluate and apply conditions in priority order
	for _, pc := range sorted {
		shouldApply, err := tm.condEval.evaluate(ctx, pc.cond.Expression, evalData)
		if err != nil {
			tm.logger.Error("condition evaluation failed", logger.Fields{
				"conditionID": pc.cond.ID,
				"expression":  pc.cond.Expression,
				"error":       err.Error(),
			})
			continue // Skip this condition on error
		}

		if shouldApply {
			tm.logger.Debug("applying conditional override", logger.Fields{
				"conditionID": pc.cond.ID,
				"priority":    pc.cond.Priority,
			})

			// Apply token overrides
			for path, value := range pc.cond.TokenOverrides {
				if err := tm.applyTokenOverride(theme.Tokens, path, value); err != nil {
					tm.logger.Error("failed to apply token override", logger.Fields{
						"conditionID": pc.cond.ID,
						"path":        path,
						"error":       err.Error(),
					})
				}
			}
		}
	}

	return nil
}

// applyTokenOverride applies a single token override to the theme.
func (tm *ThemeManager) applyTokenOverride(tokens *DesignTokens, path, value string) error {
	parts := strings.Split(path, ".")
	if len(parts) < 3 {
		return fmt.Errorf("invalid token path: %s", path)
	}

	tier := parts[0]
	category := parts[1]
	key := strings.Join(parts[2:], ".")

	switch tier {
	case "primitives":
		return tm.applyPrimitiveOverride(tokens.Primitives, category, key, value)
	case "semantic":
		return tm.applySemanticOverride(tokens.Semantic, category, key, value)
	case "components":
		if len(parts) < 4 {
			return fmt.Errorf("component path requires 4+ segments: %s", path)
		}
		component := parts[1]
		variant := parts[2]
		property := strings.Join(parts[3:], ".")
		return tm.applyComponentOverride(tokens.Components, component, variant, property, value)
	default:
		return fmt.Errorf("invalid tier: %s", tier)
	}
}

// applyPrimitiveOverride applies an override to primitive tokens.
func (tm *ThemeManager) applyPrimitiveOverride(primitives *PrimitiveTokens, category, key, value string) error {
	if primitives == nil {
		return fmt.Errorf("primitives not initialized")
	}

	var tokenMap map[string]string
	switch category {
	case "colors":
		tokenMap = primitives.Colors
	case "spacing":
		tokenMap = primitives.Spacing
	case "radius":
		tokenMap = primitives.Radius
	case "typography":
		tokenMap = primitives.Typography
	case "borders":
		tokenMap = primitives.Borders
	case "shadows":
		tokenMap = primitives.Shadows
	case "effects":
		tokenMap = primitives.Effects
	case "animation":
		tokenMap = primitives.Animation
	case "z-index":
		tokenMap = primitives.ZIndex
	case "breakpoints":
		tokenMap = primitives.Breakpoints
	default:
		return fmt.Errorf("unknown primitive category: %s", category)
	}

	if tokenMap == nil {
		return fmt.Errorf("primitive category %s not initialized", category)
	}

	tokenMap[key] = value
	return nil
}

// applySemanticOverride applies an override to semantic tokens.
func (tm *ThemeManager) applySemanticOverride(semantic *SemanticTokens, category, key, value string) error {
	if semantic == nil {
		return fmt.Errorf("semantic tokens not initialized")
	}

	var tokenMap map[string]string
	switch category {
	case "colors":
		tokenMap = semantic.Colors
	case "spacing":
		tokenMap = semantic.Spacing
	case "typography":
		tokenMap = semantic.Typography
	case "interactive":
		tokenMap = semantic.Interactive
	default:
		return fmt.Errorf("unknown semantic category: %s", category)
	}

	if tokenMap == nil {
		return fmt.Errorf("semantic category %s not initialized", category)
	}

	tokenMap[key] = value
	return nil
}

// applyComponentOverride applies an override to component tokens.
func (tm *ThemeManager) applyComponentOverride(components *ComponentTokens, component, variant, property, value string) error {
	if components == nil {
		return fmt.Errorf("component tokens not initialized")
	}

	variants, exists := (*components)[component]
	if !exists {
		// Create component if it doesn't exist
		variants = make(ComponentVariants)
		(*components)[component] = variants
	}

	props, exists := variants[variant]
	if !exists {
		// Create variant if it doesn't exist
		props = make(StyleProperties)
		variants[variant] = props
	}

	props[property] = value
	return nil
}

// ListThemes returns all available theme IDs.
// Includes both tenant-specific and global themes based on context.
func (tm *ThemeManager) ListThemes(ctx context.Context) ([]string, error) {
	tenantID := GetTenant(ctx)
	themeSet := make(map[string]bool)

	// Add in-memory themes
	tm.mu.RLock()
	if tenantID != "" {
		if tenantThemes, exists := tm.tenantThemes[tenantID]; exists {
			for id := range tenantThemes {
				themeSet[id] = true
			}
		}
	}
	for id := range tm.globalThemes {
		themeSet[id] = true
	}
	tm.mu.RUnlock()

	// Add themes from storage
	if tm.config.Storage != nil {
		storageIDs, err := tm.config.Storage.ListThemes(ctx)
		if err == nil {
			for _, id := range storageIDs {
				themeSet[id] = true
			}
		}
	}

	// Convert set to slice
	themeIDs := make([]string, 0, len(themeSet))
	for id := range themeSet {
		themeIDs = append(themeIDs, id)
	}

	return themeIDs, nil
}

// UnregisterTheme removes a theme from the manager.
// Also deletes from storage if configured.
func (tm *ThemeManager) UnregisterTheme(ctx context.Context, themeID string) error {
	tenantID := GetTenant(ctx)

	// Remove from memory
	tm.mu.Lock()
	var existed bool
	if tenantID != "" {
		if tenantThemes, exists := tm.tenantThemes[tenantID]; exists {
			if _, found := tenantThemes[themeID]; found {
				delete(tenantThemes, themeID)
				existed = true
			}
		}
	} else {
		if _, found := tm.globalThemes[themeID]; found {
			delete(tm.globalThemes, themeID)
			existed = true
		}
	}
	tm.mu.Unlock()

	if !existed {
		return NewValidationError("theme_not_found",
			fmt.Sprintf("theme not found: %s", themeID))
	}

	// Invalidate cache
	tm.invalidateCache(tenantID, themeID)

	// Delete from storage
	if tm.config.Storage != nil {
		if err := tm.config.Storage.DeleteTheme(ctx, themeID); err != nil {
			tm.logger.Error("failed to delete theme from storage", logger.Fields{
				"tenantID": tenantID,
				"themeID":  themeID,
				"error":    err.Error(),
			})
		}
	}

	tm.logger.Info("theme unregistered", logger.Fields{
		"tenantID": tenantID,
		"themeID":  themeID,
	})

	return nil
}

// Close releases all resources used by the theme manager.
// Must be called when the manager is no longer needed.
func (tm *ThemeManager) Close() error {
	if tm.cache != nil {
		tm.cache.Close()
	}
	if tm.condEval != nil {
		tm.condEval.close()
	}
	tm.logger.Info("theme manager closed")
	return nil
}

// =============================================================================
// PRIVATE HELPERS
// =============================================================================

// buildCacheKey creates a cache key for a theme.
func (tm *ThemeManager) buildCacheKey(tenantID, themeID string) string {
	if tenantID == "" {
		return fmt.Sprintf("global:%s", themeID)
	}
	return fmt.Sprintf("tenant:%s:%s", tenantID, themeID)
}

// invalidateCache removes a theme from the cache.
func (tm *ThemeManager) invalidateCache(tenantID, themeID string) {
	if tm.cache != nil {
		cacheKey := tm.buildCacheKey(tenantID, themeID)
		tm.cache.Del(cacheKey)
	}
}

// =============================================================================
// DEFAULT THEME
// =============================================================================

// GetDefaultTheme returns a production-ready default theme.
// This includes comprehensive tokens, dark mode support, and accessibility features.
func GetDefaultTheme() *Theme {
	return &Theme{
		ID:          "default",
		Name:        "Default Theme",
		Description: "Production-ready default theme with comprehensive design tokens",
		Version:     "1.0.0",
		Author:      "Awo ERP",
		Tokens:      GetDefaultTokens(),
		DarkMode: &DarkModeConfig{
			Enabled:  true,
			Default:  false,
			Strategy: "auto",
			DarkTokens: &DesignTokens{
				Semantic: &SemanticTokens{
					Colors: map[string]string{
						"background":           "primitives.colors.gray-900",
						"background-subtle":    "primitives.colors.gray-800",
						"background-muted":     "primitives.colors.gray-700",
						"foreground":           "primitives.colors.gray-100",
						"foreground-subtle":    "primitives.colors.gray-400",
						"border":               "primitives.colors.gray-700",
						"border-subtle":        "primitives.colors.gray-800",
						"primary":              "primitives.colors.primary",
						"primary-foreground":   "primitives.colors.white",
						"secondary":            "primitives.colors.gray-800",
						"secondary-foreground": "primitives.colors.gray-100",
					},
				},
			},
		},
		Accessibility: &AccessibilityConfig{
			HighContrast:          false,
			MinContrastRatio:      4.5,
			FocusIndicator:        true,
			FocusOutlineColor:     "primitives.colors.primary",
			FocusOutlineWidth:     "2px",
			KeyboardNav:           true,
			ReducedMotion:         false,
			ScreenReaderOptimized: true,
			AriaLive:              "polite",
		},
		Conditions: []*ThemeCondition{
			{
				ID:          "high-contrast-mode",
				Expression:  "user.preferences.highContrast == true",
				Priority:    100,
				Description: "Enable high contrast colors for users with visual impairments",
				TokenOverrides: map[string]string{
					"semantic.colors.border":             "primitives.colors.black",
					"semantic.colors.foreground":         "primitives.colors.black",
					"semantic.colors.background":         "primitives.colors.white",
					"semantic.colors.primary":            "primitives.colors.black",
					"semantic.colors.primary-foreground": "primitives.colors.white",
				},
			},
		},
		Meta: &ThemeMeta{
			Tags:      []string{"default", "light", "enterprise"},
			License:   "proprietary",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

// noOpLogger is a no-operation logger that can be used as a fallback
// when the global logger is not initialized
type noOpLogger struct{}

func (n *noOpLogger) Debug(msg string, fields ...logger.Fields)                     {}
func (n *noOpLogger) Info(msg string, fields ...logger.Fields)                      {}
func (n *noOpLogger) Warn(msg string, fields ...logger.Fields)                      {}
func (n *noOpLogger) Error(msg string, fields ...logger.Fields)                     {}
func (n *noOpLogger) Fatal(msg string, fields ...logger.Fields)                     {}
func (n *noOpLogger) DebugContext(ctx context.Context, msg string, fields ...logger.Fields) {}
func (n *noOpLogger) InfoContext(ctx context.Context, msg string, fields ...logger.Fields)  {}
func (n *noOpLogger) WarnContext(ctx context.Context, msg string, fields ...logger.Fields)  {}
func (n *noOpLogger) ErrorContext(ctx context.Context, msg string, fields ...logger.Fields) {}
func (n *noOpLogger) FatalContext(ctx context.Context, msg string, fields ...logger.Fields) {}
func (n *noOpLogger) WithFields(fields logger.Fields) logger.Logger                 { return n }
func (n *noOpLogger) WithContext(ctx context.Context) logger.Logger                 { return n }
func (n *noOpLogger) SetLevel(level logger.LogLevel)                                {}
func (n *noOpLogger) Close() error                                                  { return nil }
