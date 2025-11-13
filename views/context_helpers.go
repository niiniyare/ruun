package views

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/theme"
	"github.com/niiniyare/ruun/views/validation"
)

// ThemeContextKeys holds context key constants for theme-related data
type ThemeContextKeys struct {
	ThemeID    string
	DarkMode   string
	Tenant     string
	UserID     string
	Session    string
	Request    string
	FormState  string
	Validation string
}

// DefaultThemeContextKeys provides default context keys
var DefaultThemeContextKeys = ThemeContextKeys{
	ThemeID:    "ruun.theme.id",
	DarkMode:   "ruun.theme.darkMode",
	Tenant:     "ruun.theme.tenant",
	UserID:     "ruun.theme.userId",
	Session:    "ruun.theme.session",
	Request:    "ruun.theme.request",
	FormState:  "ruun.form.state",
	Validation: "ruun.validation.context",
}

// ThemeContextHelper provides comprehensive theme context management and resolution
// Integrates with theme coordination manager, renderer, and validation orchestrator
type ThemeContextHelper struct {
	themeManager    *theme.ThemeCoordinationManager
	renderer        *TemplateRenderer
	orchestrator    *validation.UIValidationOrchestrator
	contextKeys     ThemeContextKeys
	defaultTheme    string
	defaultDarkMode bool
}

// ThemeContextOptions holds configuration for theme context extraction
type ThemeContextOptions struct {
	// Theme extraction
	ThemeHeader           string   // HTTP header for theme ID (e.g., "X-Theme-ID")
	DarkModeHeader        string   // HTTP header for dark mode (e.g., "X-Dark-Mode")
	TenantHeader          string   // HTTP header for tenant (e.g., "X-Tenant-ID")
	ThemeQueryParam       string   // URL query param for theme (e.g., "theme")
	DarkModeQueryParam    string   // URL query param for dark mode (e.g., "dark")
	
	// Cookie settings
	ThemeCookie           string   // Cookie name for theme preference
	DarkModeCookie        string   // Cookie name for dark mode preference
	CookieMaxAge          int      // Cookie max age in seconds
	CookieSecure          bool     // Cookie secure flag
	CookieSameSite        http.SameSite // Cookie SameSite policy
	
	// Session integration
	SessionThemeKey       string   // Session key for theme
	SessionDarkModeKey    string   // Session key for dark mode
	SessionTenantKey      string   // Session key for tenant
	
	// Auto-detection
	AutoDetectDarkMode    bool     // Auto-detect dark mode from user agent/headers
	AutoDetectTheme       bool     // Auto-detect theme from user preferences
	
	// Fallbacks
	DefaultTheme          string   // Default theme if none specified
	DefaultDarkMode       bool     // Default dark mode if none specified
	FallbackStrategies    []string // Fallback strategies ("header", "cookie", "session", "default")
}

// NewThemeContextHelper creates a new theme context helper
func NewThemeContextHelper(
	themeManager *theme.ThemeCoordinationManager,
	renderer *TemplateRenderer,
	orchestrator *validation.UIValidationOrchestrator,
) *ThemeContextHelper {
	return &ThemeContextHelper{
		themeManager:    themeManager,
		renderer:        renderer,
		orchestrator:    orchestrator,
		contextKeys:     DefaultThemeContextKeys,
		defaultTheme:    "default",
		defaultDarkMode: false,
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// CONTEXT EXTRACTION AND COORDINATION
// ═══════════════════════════════════════════════════════════════════════════

// ExtractThemeContext extracts comprehensive theme context from HTTP request
// This is the main entry point for theme context resolution
func (h *ThemeContextHelper) ExtractThemeContext(
	ctx context.Context,
	r *http.Request,
	options *ThemeContextOptions,
) (context.Context, map[string]any, error) {
	if options == nil {
		options = h.getDefaultOptions()
	}
	
	themeContext := make(map[string]any)
	
	// Extract theme ID using fallback strategy
	themeID := h.extractThemeID(r, options)
	themeContext["themeId"] = themeID
	
	// Extract dark mode preference
	darkMode := h.extractDarkMode(r, options)
	themeContext["darkMode"] = darkMode
	
	// Extract tenant context
	tenant := h.extractTenant(r, options)
	if tenant != "" {
		themeContext["tenant"] = tenant
	}
	
	// Extract user ID if available
	if userID := h.extractUserID(r); userID != "" {
		themeContext["userId"] = userID
	}
	
	// Auto-detect preferences if enabled
	if options.AutoDetectDarkMode {
		autoDetectedDarkMode := h.autoDetectDarkMode(r)
		if autoDetectedDarkMode != darkMode {
			themeContext["autoDetectedDarkMode"] = autoDetectedDarkMode
		}
	}
	
	// Set theme context in HTTP context
	enrichedCtx := h.setThemeContextInContext(ctx, themeContext)
	
	// Coordinate with theme manager
	if h.themeManager != nil {
		err := h.coordinateThemeState(enrichedCtx, themeID, darkMode, tenant)
		if err != nil {
			return enrichedCtx, themeContext, fmt.Errorf("theme coordination failed: %w", err)
		}
	}
	
	return enrichedCtx, themeContext, nil
}

// GetThemeContextFromContext extracts theme context from Go context
func (h *ThemeContextHelper) GetThemeContextFromContext(ctx context.Context) map[string]any {
	themeContext := make(map[string]any)
	
	// Extract theme ID
	if themeID := ctx.Value(h.contextKeys.ThemeID); themeID != nil {
		themeContext["themeId"] = themeID
	}
	
	// Extract dark mode
	if darkMode := ctx.Value(h.contextKeys.DarkMode); darkMode != nil {
		themeContext["darkMode"] = darkMode
	}
	
	// Extract tenant
	if tenant := ctx.Value(h.contextKeys.Tenant); tenant != nil {
		themeContext["tenant"] = tenant
	}
	
	// Extract user ID
	if userID := ctx.Value(h.contextKeys.UserID); userID != nil {
		themeContext["userId"] = userID
	}
	
	return themeContext
}

// CreateThemeAwareFieldContext creates enhanced context for field rendering with theme awareness
func (h *ThemeContextHelper) CreateThemeAwareFieldContext(
	ctx context.Context,
	field *schema.Field,
	formState map[string]any,
) (context.Context, map[string]any, error) {
	// Get base theme context
	themeContext := h.GetThemeContextFromContext(ctx)
	
	// Add field-specific context
	fieldContext := make(map[string]any)
	for key, value := range themeContext {
		fieldContext[key] = value
	}
	
	// Add validation context
	if h.orchestrator != nil {
		validationContext := h.orchestrator.GetContextForValidation(field)
		for key, value := range validationContext {
			fieldContext[key] = value
		}
	}
	
	// Add form state context
	if formState != nil {
		fieldContext["formState"] = formState
	}
	
	// Set field context in Go context
	enrichedCtx := context.WithValue(ctx, h.contextKeys.FormState, fieldContext)
	
	return enrichedCtx, fieldContext, nil
}

// ResolveFieldTokensWithContext resolves field tokens using full context coordination
func (h *ThemeContextHelper) ResolveFieldTokensWithContext(
	ctx context.Context,
	field *schema.Field,
	props molecules.FormFieldProps,
) (map[string]string, error) {
	if h.renderer == nil || h.renderer.mapper == nil {
		return make(map[string]string), fmt.Errorf("renderer or mapper not available")
	}
	
	// Get theme context
	themeContext := h.GetThemeContextFromContext(ctx)
	
	// Extract dark mode from context
	darkMode := false
	if dm, exists := themeContext["darkMode"]; exists {
		if dark, ok := dm.(bool); ok {
			darkMode = dark
		}
	}
	
	// Resolve tokens using mapper's comprehensive resolution
	resolvedTokens := h.renderer.mapper.ResolveFieldTokens(ctx, props, darkMode)
	
	return resolvedTokens, nil
}

// ═══════════════════════════════════════════════════════════════════════════
// MIDDLEWARE AND INTEGRATION HELPERS
// ═══════════════════════════════════════════════════════════════════════════

// ThemeMiddleware creates HTTP middleware for automatic theme context extraction
func (h *ThemeContextHelper) ThemeMiddleware(options *ThemeContextOptions) func(http.Handler) http.Handler {
	if options == nil {
		options = h.getDefaultOptions()
	}
	
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract theme context
			ctx, themeContext, err := h.ExtractThemeContext(r.Context(), r, options)
			if err != nil {
				// Log error but continue with defaults
				ctx = h.setDefaultThemeContext(r.Context())
			}
			
			// Set theme cookies if they've changed
			h.updateThemeCookies(w, themeContext, options)
			
			// Continue with enriched context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// FormRenderingHelper creates a helper for theme-aware form field rendering
func (h *ThemeContextHelper) FormRenderingHelper(ctx context.Context) *FormRenderingHelper {
	return &FormRenderingHelper{
		contextHelper: h,
		context:       ctx,
		themeContext:  h.GetThemeContextFromContext(ctx),
	}
}

// ValidationHelper creates a helper for theme-aware validation operations
func (h *ThemeContextHelper) ValidationHelper(ctx context.Context) *ValidationHelper {
	return &ValidationHelper{
		contextHelper: h,
		context:       ctx,
		themeContext:  h.GetThemeContextFromContext(ctx),
		orchestrator:  h.orchestrator,
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// EXTRACTION METHODS
// ═══════════════════════════════════════════════════════════════════════════

// extractThemeID extracts theme ID using fallback strategy
func (h *ThemeContextHelper) extractThemeID(r *http.Request, options *ThemeContextOptions) string {
	for _, strategy := range options.FallbackStrategies {
		switch strategy {
		case "header":
			if themeID := r.Header.Get(options.ThemeHeader); themeID != "" {
				return themeID
			}
		case "query":
			if themeID := r.URL.Query().Get(options.ThemeQueryParam); themeID != "" {
				return themeID
			}
		case "cookie":
			if cookie, err := r.Cookie(options.ThemeCookie); err == nil {
				return cookie.Value
			}
		case "session":
			if themeID := h.extractFromSession(r, options.SessionThemeKey); themeID != "" {
				return themeID
			}
		}
	}
	return h.defaultTheme
}

// extractDarkMode extracts dark mode preference using fallback strategy
func (h *ThemeContextHelper) extractDarkMode(r *http.Request, options *ThemeContextOptions) bool {
	for _, strategy := range options.FallbackStrategies {
		switch strategy {
		case "header":
			if darkModeStr := r.Header.Get(options.DarkModeHeader); darkModeStr != "" {
				if darkMode, err := strconv.ParseBool(darkModeStr); err == nil {
					return darkMode
				}
			}
		case "query":
			if darkModeStr := r.URL.Query().Get(options.DarkModeQueryParam); darkModeStr != "" {
				return darkModeStr == "true" || darkModeStr == "1"
			}
		case "cookie":
			if cookie, err := r.Cookie(options.DarkModeCookie); err == nil {
				if darkMode, err := strconv.ParseBool(cookie.Value); err == nil {
					return darkMode
				}
			}
		case "session":
			if darkModeStr := h.extractFromSession(r, options.SessionDarkModeKey); darkModeStr != "" {
				if darkMode, err := strconv.ParseBool(darkModeStr); err == nil {
					return darkMode
				}
			}
		}
	}
	return h.defaultDarkMode
}

// extractTenant extracts tenant context
func (h *ThemeContextHelper) extractTenant(r *http.Request, options *ThemeContextOptions) string {
	// Try header first
	if tenant := r.Header.Get(options.TenantHeader); tenant != "" {
		return tenant
	}
	
	// Try session
	if tenant := h.extractFromSession(r, options.SessionTenantKey); tenant != "" {
		return tenant
	}
	
	// Try to extract from subdomain
	if tenant := h.extractTenantFromHost(r.Host); tenant != "" {
		return tenant
	}
	
	return ""
}

// extractUserID extracts user ID from request context (placeholder implementation)
func (h *ThemeContextHelper) extractUserID(r *http.Request) string {
	// This would integrate with your authentication system
	// For now, check for common user ID sources
	if userID := r.Header.Get("X-User-ID"); userID != "" {
		return userID
	}
	return ""
}

// autoDetectDarkMode attempts to auto-detect dark mode from browser/system hints
func (h *ThemeContextHelper) autoDetectDarkMode(r *http.Request) bool {
	// Check for prefers-color-scheme in Accept header (some browsers send this)
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "prefers-color-scheme=dark") {
		return true
	}
	
	// Check for SEC-CH-Prefers-Color-Scheme header (Client Hints)
	if colorScheme := r.Header.Get("SEC-CH-Prefers-Color-Scheme"); colorScheme == "dark" {
		return true
	}
	
	return false
}

// ═══════════════════════════════════════════════════════════════════════════
// HELPER CLASSES
// ═══════════════════════════════════════════════════════════════════════════

// FormRenderingHelper provides theme-aware form rendering utilities
type FormRenderingHelper struct {
	contextHelper *ThemeContextHelper
	context       context.Context
	themeContext  map[string]any
}

// RenderFieldWithTheme renders a field with complete theme context
func (h *FormRenderingHelper) RenderFieldWithTheme(
	field *schema.Field,
	value any,
	errors []string,
	touched, dirty bool,
) (string, error) {
	if h.contextHelper.renderer == nil {
		return "", fmt.Errorf("renderer not available")
	}
	
	// Create field context
	fieldCtx, _, err := h.contextHelper.CreateThemeAwareFieldContext(h.context, field, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create field context: %w", err)
	}
	
	// Use enhanced renderer method - renderer is already *TemplateRenderer
	if h.contextHelper.renderer != nil {
		// TODO:INTEGRATION_FIX - RenderFieldWithThemeContext method needs to be implemented
		// For now, use standard RenderField method
		return h.contextHelper.renderer.RenderField(fieldCtx, field, value, errors, touched, dirty)
	}
	
	// Fallback error if no renderer available
	return "", fmt.Errorf("no renderer available")
}

// GetFieldTokens gets resolved tokens for a field
func (h *FormRenderingHelper) GetFieldTokens(field *schema.Field, props molecules.FormFieldProps) (map[string]string, error) {
	return h.contextHelper.ResolveFieldTokensWithContext(h.context, field, props)
}

// ValidationHelper provides theme-aware validation utilities
type ValidationHelper struct {
	contextHelper *ThemeContextHelper
	context       context.Context
	themeContext  map[string]any
	orchestrator  *validation.UIValidationOrchestrator
}

// ValidateFieldWithTheme validates a field with theme context
func (h *ValidationHelper) ValidateFieldWithTheme(
	fieldName string,
	value any,
	trigger validation.ValidationTrigger,
) error {
	if h.orchestrator == nil {
		return fmt.Errorf("orchestrator not available")
	}
	
	// TODO:INTEGRATION_FIX - Theme-aware validation integration needs proper interface
	// For now, use standard validation method
	return h.orchestrator.ValidateFieldWithUI(h.context, fieldName, value, trigger)
}

// ═══════════════════════════════════════════════════════════════════════════
// UTILITY METHODS
// ═══════════════════════════════════════════════════════════════════════════

// getDefaultOptions returns default theme context options
func (h *ThemeContextHelper) getDefaultOptions() *ThemeContextOptions {
	return &ThemeContextOptions{
		ThemeHeader:         "X-Theme-ID",
		DarkModeHeader:      "X-Dark-Mode",
		TenantHeader:        "X-Tenant-ID",
		ThemeQueryParam:     "theme",
		DarkModeQueryParam:  "dark",
		ThemeCookie:         "ruun_theme",
		DarkModeCookie:      "ruun_dark_mode",
		CookieMaxAge:        86400 * 30, // 30 days
		CookieSecure:        true,
		CookieSameSite:      http.SameSiteStrictMode,
		SessionThemeKey:     "theme",
		SessionDarkModeKey:  "dark_mode",
		SessionTenantKey:    "tenant",
		AutoDetectDarkMode:  true,
		AutoDetectTheme:     false,
		DefaultTheme:        h.defaultTheme,
		DefaultDarkMode:     h.defaultDarkMode,
		FallbackStrategies:  []string{"header", "query", "cookie", "session", "default"},
	}
}

// setThemeContextInContext sets theme context in Go context
func (h *ThemeContextHelper) setThemeContextInContext(ctx context.Context, themeContext map[string]any) context.Context {
	enrichedCtx := ctx
	
	if themeID, exists := themeContext["themeId"]; exists {
		enrichedCtx = context.WithValue(enrichedCtx, h.contextKeys.ThemeID, themeID)
	}
	
	if darkMode, exists := themeContext["darkMode"]; exists {
		enrichedCtx = context.WithValue(enrichedCtx, h.contextKeys.DarkMode, darkMode)
	}
	
	if tenant, exists := themeContext["tenant"]; exists {
		enrichedCtx = context.WithValue(enrichedCtx, h.contextKeys.Tenant, tenant)
	}
	
	if userID, exists := themeContext["userId"]; exists {
		enrichedCtx = context.WithValue(enrichedCtx, h.contextKeys.UserID, userID)
	}
	
	return enrichedCtx
}

// setDefaultThemeContext sets default theme context
func (h *ThemeContextHelper) setDefaultThemeContext(ctx context.Context) context.Context {
	themeContext := map[string]any{
		"themeId":  h.defaultTheme,
		"darkMode": h.defaultDarkMode,
	}
	return h.setThemeContextInContext(ctx, themeContext)
}

// coordinateThemeState coordinates theme state with theme manager
func (h *ThemeContextHelper) coordinateThemeState(ctx context.Context, themeID string, darkMode bool, tenant string) error {
	if h.themeManager == nil {
		return nil
	}
	
	options := &theme.ThemeOptions{
		DarkMode: darkMode,
		Tenant:   tenant,
		Context: map[string]any{
			"requestTime": time.Now(),
		},
	}
	
	return h.themeManager.SetActiveTheme(ctx, "request", themeID, options)
}

// updateThemeCookies updates theme-related cookies if values have changed
func (h *ThemeContextHelper) updateThemeCookies(w http.ResponseWriter, themeContext map[string]any, options *ThemeContextOptions) {
	// Set theme cookie
	if themeID, exists := themeContext["themeId"]; exists {
		if theme, ok := themeID.(string); ok && theme != "" {
			http.SetCookie(w, &http.Cookie{
				Name:     options.ThemeCookie,
				Value:    theme,
				MaxAge:   options.CookieMaxAge,
				Secure:   options.CookieSecure,
				SameSite: options.CookieSameSite,
				Path:     "/",
			})
		}
	}
	
	// Set dark mode cookie
	if darkMode, exists := themeContext["darkMode"]; exists {
		if dark, ok := darkMode.(bool); ok {
			http.SetCookie(w, &http.Cookie{
				Name:     options.DarkModeCookie,
				Value:    strconv.FormatBool(dark),
				MaxAge:   options.CookieMaxAge,
				Secure:   options.CookieSecure,
				SameSite: options.CookieSameSite,
				Path:     "/",
			})
		}
	}
}

// extractFromSession extracts value from session (placeholder implementation)
func (h *ThemeContextHelper) extractFromSession(r *http.Request, key string) string {
	// This would integrate with your session management system
	// For now, return empty string
	return ""
}

// extractTenantFromHost extracts tenant from host/subdomain
func (h *ThemeContextHelper) extractTenantFromHost(host string) string {
	// Extract tenant from subdomain (e.g., "tenant.example.com" -> "tenant")
	parts := strings.Split(host, ".")
	if len(parts) > 2 {
		return parts[0]
	}
	return ""
}