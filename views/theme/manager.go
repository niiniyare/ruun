package theme

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/tokens"
	"github.com/niiniyare/ruun/views/validation"
)

// ThemeCoordinationManager provides comprehensive theme coordination for multi-tenant validation UI
// Integrates with the three-tier token architecture: Primitives → Semantic → Components
type ThemeCoordinationManager struct {
	// Core theme management
	activeThemes    map[string]*ActiveTheme // Context-based active themes
	tokenResolver   *TokenResolver          // Three-tier token resolution
	darkModeManager *DarkModeManager        // Dark mode detection and management

	// Multi-tenant support
	tenantResolver *TenantResolver // Tenant-specific theme resolution

	// Event system
	eventBus        *ThemeEventBus        // Theme change events
	changeListeners []ThemeChangeListener // Change event listeners

	// Configuration
	config *ThemeManagerConfig

	// Thread safety
	mu sync.RWMutex
}

// ActiveTheme represents an active theme context
type ActiveTheme struct {
	ID               string               // Theme identifier
	Tenant           string               // Tenant/organization context
	DarkMode         bool                 // Dark mode state
	ValidationTokens *schema.DesignTokens // Resolved validation tokens
	LastUpdated      time.Time            // Last update timestamp
	Context          map[string]any       // Additional context data
}

// TokenResolver handles three-tier token resolution with validation state awareness
type TokenResolver struct {
	baseTokens       *schema.DesignTokens         // Base design tokens
	validationTokens *schema.DesignTokens         // Validation-specific tokens
	overrideCache    map[string]map[string]string // Cached token overrides

	// Dark mode token mappings
	darkModeOverrides map[string]map[string]string

	mu sync.RWMutex
}

// DarkModeManager handles dark mode detection and synchronization
type DarkModeManager struct {
	// Detection strategies
	systemDetection bool            // Use system preference detection
	userPreferences map[string]bool // User-specific dark mode preferences
	tenantDefaults  map[string]bool // Tenant default dark mode settings

	// Auto-detection
	autoDetectionEnabled bool
	detectionInterval    time.Duration

	mu sync.RWMutex
}

// TenantResolver provides multi-tenant theme resolution
type TenantResolver struct {
	tenantThemes    map[string]string            // Tenant -> Theme ID mapping
	tenantOverrides map[string]map[string]string // Tenant-specific token overrides
	defaultTheme    string                       // Default theme for unknown tenants

	mu sync.RWMutex
}

// ThemeEventBus handles theme change events
type ThemeEventBus struct {
	listeners map[string][]func(event ThemeEvent)
	mu        sync.RWMutex
}

// ThemeEvent represents a theme change event
type ThemeEvent struct {
	Type      string         // Event type (change, toggle, etc.)
	ThemeID   string         // Theme identifier
	Context   string         // Context (tenant, user, etc.)
	DarkMode  bool           // Dark mode state
	Timestamp time.Time      // Event timestamp
	Data      map[string]any // Additional event data
}

// ThemeChangeListener interface for theme change callbacks
type ThemeChangeListener func(themeID string, darkMode bool)

// ThemeManagerConfig holds theme manager configuration
type ThemeManagerConfig struct {
	// Default settings
	DefaultTheme    string // Default theme ID
	DefaultDarkMode bool   // Default dark mode state

	// Auto-detection
	EnableSystemDetection bool          // Enable system dark mode detection
	DetectionInterval     time.Duration // Dark mode detection interval

	// Multi-tenant support
	EnableMultiTenant  bool   // Enable multi-tenant themes
	TenantContextField string // Field name for tenant context

	// Performance
	EnableTokenCaching bool          // Enable token resolution caching
	CacheExpiration    time.Duration // Token cache expiration

	// Validation integration
	ValidateTokenContrast bool // Validate color contrast for accessibility
	EnableThemeEvents     bool // Enable theme change events
}

// NewThemeCoordinationManager creates a new theme coordination manager
func NewThemeCoordinationManager() *ThemeCoordinationManager {
	config := &ThemeManagerConfig{
		DefaultTheme:          "default",
		DefaultDarkMode:       false,
		EnableSystemDetection: true,
		DetectionInterval:     5 * time.Second,
		EnableMultiTenant:     true,
		TenantContextField:    "tenant",
		EnableTokenCaching:    true,
		CacheExpiration:       5 * time.Minute,
		ValidateTokenContrast: true,
		EnableThemeEvents:     true,
	}

	manager := &ThemeCoordinationManager{
		activeThemes:    make(map[string]*ActiveTheme),
		tokenResolver:   NewTokenResolver(),
		darkModeManager: NewDarkModeManager(),
		tenantResolver:  NewTenantResolver(),
		eventBus:        NewThemeEventBus(),
		changeListeners: make([]ThemeChangeListener, 0),
		config:          config,
	}

	// Initialize with validation tokens
	manager.tokenResolver.LoadValidationTokens(tokens.ValidationTokens)

	// Start auto-detection if enabled
	if config.EnableSystemDetection {
		manager.darkModeManager.StartAutoDetection(config.DetectionInterval)
	}

	return manager
}

// ═══════════════════════════════════════════════════════════════════════════
// THEME MANAGER INTERFACE IMPLEMENTATION
// ═══════════════════════════════════════════════════════════════════════════

// GetActiveTheme returns the active theme for a given context
func (m *ThemeCoordinationManager) GetActiveTheme(context string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if activeTheme, exists := m.activeThemes[context]; exists {
		return activeTheme.ID
	}

	// Check tenant-specific theme
	if m.config.EnableMultiTenant {
		if themeID := m.tenantResolver.GetThemeForTenant(context); themeID != "" {
			return themeID
		}
	}

	return m.config.DefaultTheme
}

// IsDarkModeEnabled checks if dark mode is enabled for a given context
func (m *ThemeCoordinationManager) IsDarkModeEnabled(context string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if activeTheme, exists := m.activeThemes[context]; exists {
		return activeTheme.DarkMode
	}

	// Check dark mode manager
	return m.darkModeManager.IsDarkModeEnabled(context)
}

// GetThemeTokens returns theme tokens for a specific theme and validation state
func (m *ThemeCoordinationManager) GetThemeTokens(themeID string, validationState validation.ValidationState) map[string]string {
	// Get base validation tokens for the state
	stateStr := m.validationStateToString(validationState)
	baseTokens := m.getValidationTokensForState(stateStr)

	// Apply theme-specific overrides
	themeTokens := m.tokenResolver.ResolveThemeTokens(themeID, baseTokens)

	// Apply tenant-specific overrides if enabled
	if m.config.EnableMultiTenant {
		tenantOverrides := m.tenantResolver.GetTenantTokenOverrides(themeID)
		for key, value := range tenantOverrides {
			themeTokens[key] = value
		}
	}

	return themeTokens
}

// SubscribeToThemeChanges subscribes to theme change events
func (m *ThemeCoordinationManager) SubscribeToThemeChanges(callback func(themeID string, darkMode bool)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.changeListeners = append(m.changeListeners, callback)
}

// ═══════════════════════════════════════════════════════════════════════════
// ADVANCED THEME COORDINATION METHODS
// ═══════════════════════════════════════════════════════════════════════════

// SetActiveTheme sets the active theme for a context with comprehensive coordination
func (m *ThemeCoordinationManager) SetActiveTheme(
	ctx context.Context,
	contextName string,
	themeID string,
	options *ThemeOptions,
) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Validate theme exists
	if !m.validateThemeExists(themeID) {
		return fmt.Errorf("theme %s not found", themeID)
	}

	// Resolve options
	opts := m.resolveThemeOptions(options)

	// Create active theme
	activeTheme := &ActiveTheme{
		ID:               themeID,
		Tenant:           opts.Tenant,
		DarkMode:         opts.DarkMode,
		ValidationTokens: m.tokenResolver.GetValidationTokens(themeID, opts.DarkMode),
		LastUpdated:      time.Now(),
		Context:          opts.Context,
	}

	// Store active theme
	m.activeThemes[contextName] = activeTheme

	// Emit theme change event
	if m.config.EnableThemeEvents {
		event := ThemeEvent{
			Type:      "theme_change",
			ThemeID:   themeID,
			Context:   contextName,
			DarkMode:  opts.DarkMode,
			Timestamp: time.Now(),
			Data: map[string]any{
				"tenant": opts.Tenant,
				"source": "manual",
			},
		}
		m.eventBus.Emit("theme_change", event)
	}

	// Notify listeners
	m.notifyThemeChange(themeID, opts.DarkMode)

	return nil
}

// ToggleDarkMode toggles dark mode for a specific context
func (m *ThemeCoordinationManager) ToggleDarkMode(
	ctx context.Context,
	contextName string,
) (bool, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	activeTheme, exists := m.activeThemes[contextName]
	if !exists {
		return false, fmt.Errorf("no active theme for context %s", contextName)
	}

	// Toggle dark mode
	activeTheme.DarkMode = !activeTheme.DarkMode
	activeTheme.LastUpdated = time.Now()

	// Update validation tokens for new dark mode state
	activeTheme.ValidationTokens = m.tokenResolver.GetValidationTokens(activeTheme.ID, activeTheme.DarkMode)

	// Update dark mode manager
	m.darkModeManager.SetDarkMode(contextName, activeTheme.DarkMode)

	// Emit dark mode toggle event
	if m.config.EnableThemeEvents {
		event := ThemeEvent{
			Type:      "dark_mode_toggle",
			ThemeID:   activeTheme.ID,
			Context:   contextName,
			DarkMode:  activeTheme.DarkMode,
			Timestamp: time.Now(),
			Data: map[string]any{
				"previous_state": !activeTheme.DarkMode,
			},
		}
		m.eventBus.Emit("dark_mode_toggle", event)
	}

	// Notify listeners
	m.notifyThemeChange(activeTheme.ID, activeTheme.DarkMode)

	return activeTheme.DarkMode, nil
}

// GetValidationTokensForState returns validation tokens with full theme coordination
func (m *ThemeCoordinationManager) GetValidationTokensForState(
	contextName string,
	validationState validation.ValidationState,
) map[string]string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Get active theme
	activeTheme := m.getActiveThemeForContext(contextName)

	// Get validation state tokens
	stateTokens := m.GetThemeTokens(activeTheme.ID, validationState)

	// Apply dark mode overrides if needed
	if activeTheme.DarkMode {
		darkTokens := m.getDarkModeValidationTokens(activeTheme.ID, validationState)
		for key, value := range darkTokens {
			stateTokens[key] = value
		}
	}

	return stateTokens
}

// RegisterTenant registers a tenant with specific theme configuration
func (m *ThemeCoordinationManager) RegisterTenant(
	tenantID string,
	config *TenantThemeConfig,
) error {
	return m.tenantResolver.RegisterTenant(tenantID, config)
}

// ═══════════════════════════════════════════════════════════════════════════
// HELPER METHODS
// ═══════════════════════════════════════════════════════════════════════════

// ThemeOptions holds options for theme setting
type ThemeOptions struct {
	DarkMode bool           // Dark mode preference
	Tenant   string         // Tenant context
	Context  map[string]any // Additional context
}

// TenantThemeConfig holds tenant-specific theme configuration
type TenantThemeConfig struct {
	DefaultTheme    string            // Default theme for tenant
	DefaultDarkMode bool              // Default dark mode for tenant
	TokenOverrides  map[string]string // Tenant-specific token overrides
	AllowedThemes   []string          // Allowed themes for tenant
}

// resolveThemeOptions resolves theme options with defaults
func (m *ThemeCoordinationManager) resolveThemeOptions(options *ThemeOptions) *ThemeOptions {
	if options == nil {
		return &ThemeOptions{
			DarkMode: m.config.DefaultDarkMode,
			Tenant:   "",
			Context:  make(map[string]any),
		}
	}

	if options.Context == nil {
		options.Context = make(map[string]any)
	}

	return options
}

// validateThemeExists validates that a theme exists
func (m *ThemeCoordinationManager) validateThemeExists(themeID string) bool {
	// In a full implementation, this would check against a theme registry
	// For now, accept any theme ID
	return themeID != ""
}

// getActiveThemeForContext returns the active theme for a context
func (m *ThemeCoordinationManager) getActiveThemeForContext(contextName string) *ActiveTheme {
	if activeTheme, exists := m.activeThemes[contextName]; exists {
		return activeTheme
	}

	// Return default active theme
	return &ActiveTheme{
		ID:       m.config.DefaultTheme,
		DarkMode: m.config.DefaultDarkMode,
		Context:  make(map[string]any),
	}
}

// getDarkModeValidationTokens returns dark mode validation tokens
func (m *ThemeCoordinationManager) getDarkModeValidationTokens(
	themeID string,
	validationState validation.ValidationState,
) map[string]string {
	stateStr := m.validationStateToString(validationState)
	return m.getDarkModeTokensForState(stateStr)
}

// validationStateToString converts validation state to string
func (m *ThemeCoordinationManager) validationStateToString(state validation.ValidationState) string {
	switch state {
	case validation.ValidationStateValidating:
		return "validating"
	case validation.ValidationStateValid:
		return "valid"
	case validation.ValidationStateInvalid:
		return "invalid"
	case validation.ValidationStateWarning:
		return "warning"
	default:
		return "idle"
	}
}

// getValidationTokensForState returns validation tokens for a specific state
func (m *ThemeCoordinationManager) getValidationTokensForState(stateStr string) map[string]string {
	// Return basic token map for validation states
	baseTokens := map[string]string{
		"background": "#ffffff",
		"border":     "#e5e7eb",
		"text":       "#111827",
	}

	switch stateStr {
	case "validating":
		baseTokens["border"] = "#f59e0b"
		baseTokens["icon"] = "⟳"
	case "valid":
		baseTokens["border"] = "#16a34a"
		baseTokens["icon"] = "✓"
	case "invalid":
		baseTokens["border"] = "#dc2626"
		baseTokens["icon"] = "✗"
	case "warning":
		baseTokens["border"] = "#d97706"
		baseTokens["icon"] = "⚠"
	}

	return baseTokens
}

// getDarkModeTokensForState returns dark mode validation tokens for a specific state
func (m *ThemeCoordinationManager) getDarkModeTokensForState(stateStr string) map[string]string {
	// Return dark mode token overrides
	darkTokens := map[string]string{
		"background": "#111827",
		"border":     "#374151",
		"text":       "#f9fafb",
	}

	switch stateStr {
	case "validating":
		darkTokens["border"] = "#fbbf24"
	case "valid":
		darkTokens["border"] = "#22c55e"
	case "invalid":
		darkTokens["border"] = "#ef4444"
	case "warning":
		darkTokens["border"] = "#f59e0b"
	}

	return darkTokens
}

// notifyThemeChange notifies all registered listeners of theme changes
func (m *ThemeCoordinationManager) notifyThemeChange(themeID string, darkMode bool) {
	for _, listener := range m.changeListeners {
		// Run listeners in goroutines to avoid blocking
		go listener(themeID, darkMode)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// COMPONENT CONSTRUCTORS
// ═══════════════════════════════════════════════════════════════════════════

// NewTokenResolver creates a new token resolver
func NewTokenResolver() *TokenResolver {
	return &TokenResolver{
		overrideCache:     make(map[string]map[string]string),
		darkModeOverrides: make(map[string]map[string]string),
	}
}

// NewDarkModeManager creates a new dark mode manager
func NewDarkModeManager() *DarkModeManager {
	return &DarkModeManager{
		systemDetection:      true,
		userPreferences:      make(map[string]bool),
		tenantDefaults:       make(map[string]bool),
		autoDetectionEnabled: false,
		detectionInterval:    5 * time.Second,
	}
}

// NewTenantResolver creates a new tenant resolver
func NewTenantResolver() *TenantResolver {
	return &TenantResolver{
		tenantThemes:    make(map[string]string),
		tenantOverrides: make(map[string]map[string]string),
		defaultTheme:    "default",
	}
}

// NewThemeEventBus creates a new theme event bus
func NewThemeEventBus() *ThemeEventBus {
	return &ThemeEventBus{
		listeners: make(map[string][]func(event ThemeEvent)),
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// TOKEN RESOLVER METHODS
// ═══════════════════════════════════════════════════════════════════════════

// LoadValidationTokens loads validation tokens into the resolver
func (r *TokenResolver) LoadValidationTokens(validationTokens *schema.DesignTokens) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.validationTokens = validationTokens
}

// ResolveThemeTokens resolves tokens for a specific theme
func (r *TokenResolver) ResolveThemeTokens(themeID string, baseTokens map[string]string) map[string]string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Start with base tokens
	resolved := make(map[string]string)
	for key, value := range baseTokens {
		resolved[key] = value
	}

	// Apply cached overrides if available
	if overrides, exists := r.overrideCache[themeID]; exists {
		for key, value := range overrides {
			resolved[key] = value
		}
	}

	return resolved
}

// GetValidationTokens returns validation tokens for a theme and dark mode state
func (r *TokenResolver) GetValidationTokens(themeID string, darkMode bool) *schema.DesignTokens {
	// In a full implementation, this would resolve theme-specific validation tokens
	// For now, return the base validation tokens
	return tokens.ValidationTokens
}

// ═══════════════════════════════════════════════════════════════════════════
// DARK MODE MANAGER METHODS
// ═══════════════════════════════════════════════════════════════════════════

// IsDarkModeEnabled checks if dark mode is enabled for a context
func (d *DarkModeManager) IsDarkModeEnabled(context string) bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	// Check user preferences first
	if darkMode, exists := d.userPreferences[context]; exists {
		return darkMode
	}

	// Check tenant defaults
	if darkMode, exists := d.tenantDefaults[context]; exists {
		return darkMode
	}

	// Default to false
	return false
}

// SetDarkMode sets dark mode preference for a context
func (d *DarkModeManager) SetDarkMode(context string, darkMode bool) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.userPreferences[context] = darkMode
}

// StartAutoDetection starts automatic dark mode detection
func (d *DarkModeManager) StartAutoDetection(interval time.Duration) {
	d.mu.Lock()
	d.autoDetectionEnabled = true
	d.detectionInterval = interval
	d.mu.Unlock()

	// In a full implementation, this would start a goroutine for system detection
}

// ═══════════════════════════════════════════════════════════════════════════
// TENANT RESOLVER METHODS
// ═══════════════════════════════════════════════════════════════════════════

// GetThemeForTenant returns the theme ID for a tenant
func (t *TenantResolver) GetThemeForTenant(tenant string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if theme, exists := t.tenantThemes[tenant]; exists {
		return theme
	}

	return t.defaultTheme
}

// GetTenantTokenOverrides returns token overrides for a tenant
func (t *TenantResolver) GetTenantTokenOverrides(tenant string) map[string]string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if overrides, exists := t.tenantOverrides[tenant]; exists {
		return overrides
	}

	return make(map[string]string)
}

// RegisterTenant registers a tenant with theme configuration
func (t *TenantResolver) RegisterTenant(tenantID string, config *TenantThemeConfig) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tenantThemes[tenantID] = config.DefaultTheme
	t.tenantOverrides[tenantID] = config.TokenOverrides

	return nil
}

// ═══════════════════════════════════════════════════════════════════════════
// THEME EVENT BUS METHODS
// ═══════════════════════════════════════════════════════════════════════════

// Emit emits a theme event
func (e *ThemeEventBus) Emit(eventType string, event ThemeEvent) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if listeners, exists := e.listeners[eventType]; exists {
		for _, listener := range listeners {
			go listener(event)
		}
	}
}

// Subscribe subscribes to theme events
func (e *ThemeEventBus) Subscribe(eventType string, listener func(event ThemeEvent)) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.listeners[eventType]; !exists {
		e.listeners[eventType] = make([]func(event ThemeEvent), 0)
	}

	e.listeners[eventType] = append(e.listeners[eventType], listener)
}
