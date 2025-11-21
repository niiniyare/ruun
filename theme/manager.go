package theme

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// ConditionEvaluator evaluates boolean expressions with runtime context.
// This interface allows plugging in external condition evaluation packages.
type ConditionEvaluator interface {
	Evaluate(ctx context.Context, expression string, data map[string]any) (bool, error)
}

// Manager manages theme lifecycle, compilation, and runtime switching.
type Manager struct {
	storage   Storage
	resolver  *Resolver
	compiler  *Compiler
	validator *Validator
	evaluator ConditionEvaluator
	
	themes map[string]*Theme
	mu     sync.RWMutex
	
	config *ManagerConfig
}

// ManagerConfig configures the theme manager behavior.
type ManagerConfig struct {
	EnableCaching      bool
	EnableValidation   bool
	EnableConditionals bool
	PreloadThemes      bool
	AutoCompile        bool
	StrictMode         bool
	MaxThemes          int
	
	ResolverConfig  *ResolverConfig
	CompilerConfig  *CompilerConfig
	ValidatorConfig *ValidatorConfig
}

// DefaultManagerConfig returns production-ready manager configuration.
func DefaultManagerConfig() *ManagerConfig {
	return &ManagerConfig{
		EnableCaching:      true,
		EnableValidation:   true,
		EnableConditionals: true,
		PreloadThemes:      false,
		AutoCompile:        true,
		StrictMode:         false,
		MaxThemes:          100,
		ResolverConfig:     DefaultResolverConfig(),
		CompilerConfig:     DefaultCompilerConfig(),
		ValidatorConfig:    DefaultValidatorConfig(),
	}
}

// NewManager creates a new theme manager with the given configuration.
func NewManager(config *ManagerConfig, storage Storage, evaluator ConditionEvaluator) (*Manager, error) {
	if config == nil {
		config = DefaultManagerConfig()
	}
	if storage == nil {
		storage = NewMemoryStorage()
	}

	resolver, err := NewResolver(config.ResolverConfig)
	if err != nil {
		return nil, WrapError(ErrCodeResolution, "failed to create resolver", err)
	}

	compiler, err := NewCompiler(config.CompilerConfig)
	if err != nil {
		return nil, WrapError(ErrCodeCompilation, "failed to create compiler", err)
	}

	validator := NewValidator(config.ValidatorConfig)

	m := &Manager{
		storage:   storage,
		resolver:  resolver,
		compiler:  compiler,
		validator: validator,
		evaluator: evaluator,
		themes:    make(map[string]*Theme),
		config:    config,
	}

	if config.PreloadThemes {
		if err := m.preloadThemes(context.Background()); err != nil {
			return nil, WrapError(ErrCodeStorage, "failed to preload themes", err)
		}
	}

	return m, nil
}

// RegisterTheme registers a new theme with validation.
func (m *Manager) RegisterTheme(ctx context.Context, theme *Theme) error {
	if theme == nil {
		return NewError(ErrCodeValidation, "theme cannot be nil")
	}

	if m.config.EnableValidation {
		result := m.validator.Validate(theme)
		if !result.Valid {
			return NewErrorf(ErrCodeValidation, "theme validation failed: %d issues", len(result.Issues))
		}
	}

	m.mu.Lock()
	if len(m.themes) >= m.config.MaxThemes {
		m.mu.Unlock()
		return NewErrorf(ErrCodeValidation, "maximum theme limit reached: %d", m.config.MaxThemes)
	}
	m.mu.Unlock()

	now := time.Now()
	theme.SetCreatedAt(now)
	theme.SetUpdatedAt(now)

	if err := m.storage.SaveTheme(ctx, theme); err != nil {
		return WrapError(ErrCodeStorage, "failed to save theme", err)
	}

	m.mu.Lock()
	m.themes[theme.ID] = theme
	m.mu.Unlock()

	return nil
}

// GetTheme retrieves and optionally compiles a theme with conditional overrides.
func (m *Manager) GetTheme(ctx context.Context, themeID string, evalData map[string]any) (*CompiledTheme, error) {
	theme, err := m.loadTheme(ctx, themeID)
	if err != nil {
		return nil, err
	}

	tenantID := getTenantID(ctx)
	darkMode := getDarkMode(ctx)

	processedTheme := theme.Clone()

	if m.config.EnableConditionals && len(theme.Conditions) > 0 && m.evaluator != nil {
		processedTheme, err = m.applyConditionalOverrides(ctx, processedTheme, evalData)
		if err != nil {
			return nil, WrapError(ErrCodeCondition, "failed to apply conditional overrides", err)
		}
	}

	var resolvedTokens *Tokens
	if m.config.AutoCompile {
		resolvedTokens, err = m.resolver.ResolveAll(ctx, processedTheme, tenantID, darkMode)
		if err != nil {
			return nil, WrapError(ErrCodeResolution, "failed to resolve tokens", err)
		}
	} else {
		resolvedTokens = processedTheme.Tokens
	}

	css, err := m.compiler.Compile(ctx, resolvedTokens, processedTheme)
	if err != nil {
		return nil, WrapError(ErrCodeCompilation, "failed to compile CSS", err)
	}

	return &CompiledTheme{
		Theme:          processedTheme,
		ResolvedTokens: resolvedTokens,
		CSS:            css,
		DarkMode:       darkMode,
		TenantID:       tenantID,
		CompiledAt:     time.Now(),
	}, nil
}

// UpdateTheme updates an existing theme.
func (m *Manager) UpdateTheme(ctx context.Context, theme *Theme) error {
	if theme == nil {
		return NewError(ErrCodeValidation, "theme cannot be nil")
	}

	existing, err := m.loadTheme(ctx, theme.ID)
	if err != nil {
		return err
	}

	if m.config.EnableValidation {
		result := m.validator.Validate(theme)
		if !result.Valid {
			return NewErrorf(ErrCodeValidation, "theme validation failed: %d issues", len(result.Issues))
		}
	}

	theme.SetCreatedAt(existing.CreatedAt())
	theme.SetUpdatedAt(time.Now())

	if err := m.storage.SaveTheme(ctx, theme); err != nil {
		return WrapError(ErrCodeStorage, "failed to update theme", err)
	}

	m.mu.Lock()
	m.themes[theme.ID] = theme
	m.mu.Unlock()

	m.resolver.InvalidateCache(theme.ID, "")

	return nil
}

// DeleteTheme deletes a theme.
func (m *Manager) DeleteTheme(ctx context.Context, themeID string) error {
	if err := m.storage.DeleteTheme(ctx, themeID); err != nil {
		return WrapError(ErrCodeStorage, "failed to delete theme", err)
	}

	m.mu.Lock()
	delete(m.themes, themeID)
	m.mu.Unlock()

	m.resolver.InvalidateCache(themeID, "")

	return nil
}

// ListThemes returns all registered themes.
func (m *Manager) ListThemes(ctx context.Context) ([]*Theme, error) {
	themes, err := m.storage.ListThemes(ctx)
	if err != nil {
		return nil, WrapError(ErrCodeStorage, "failed to list themes", err)
	}

	m.mu.Lock()
	for _, theme := range themes {
		m.themes[theme.ID] = theme
	}
	m.mu.Unlock()

	return themes, nil
}

// ResolveToken resolves a single token reference.
func (m *Manager) ResolveToken(ctx context.Context, themeID string, ref TokenReference) (*ResolvedToken, error) {
	theme, err := m.loadTheme(ctx, themeID)
	if err != nil {
		return nil, err
	}

	tenantID := getTenantID(ctx)
	darkMode := getDarkMode(ctx)

	return m.resolver.Resolve(ctx, ref, theme, tenantID, darkMode)
}

// InvalidateCache invalidates all caches for a theme or tenant.
func (m *Manager) InvalidateCache(themeID, tenantID string) {
	m.resolver.InvalidateCache(themeID, tenantID)
	if m.compiler.cache != nil {
		if tenantID != "" {
			m.compiler.cache.Clear()
		}
	}
}

// GetStats returns manager performance statistics.
func (m *Manager) GetStats() *ManagerStats {
	m.mu.RLock()
	themeCount := len(m.themes)
	m.mu.RUnlock()

	return &ManagerStats{
		ThemeCount:     themeCount,
		ResolverStats:  m.resolver.GetStats(),
		CompilerStats:  m.compiler.GetStats(),
		ValidatorStats: m.validator.GetStats(),
	}
}

// loadTheme loads a theme from memory or storage.
func (m *Manager) loadTheme(ctx context.Context, themeID string) (*Theme, error) {
	m.mu.RLock()
	theme, exists := m.themes[themeID]
	m.mu.RUnlock()

	if exists {
		return theme, nil
	}

	theme, err := m.storage.GetTheme(ctx, themeID)
	if err != nil {
		return nil, WrapError(ErrCodeNotFound, fmt.Sprintf("theme not found: %s", themeID), err)
	}

	m.mu.Lock()
	m.themes[themeID] = theme
	m.mu.Unlock()

	return theme, nil
}

// applyConditionalOverrides applies conditional theme overrides based on runtime context.
func (m *Manager) applyConditionalOverrides(ctx context.Context, theme *Theme, evalData map[string]any) (*Theme, error) {
	if len(theme.Conditions) == 0 || m.evaluator == nil {
		return theme, nil
	}

	activeConditions := make([]*Condition, 0, len(theme.Conditions))

	for _, condition := range theme.Conditions {
		matched, err := m.evaluator.Evaluate(ctx, condition.Expression, evalData)
		if err != nil {
			if m.config.StrictMode {
				return nil, WrapError(ErrCodeCondition,
					fmt.Sprintf("failed to evaluate condition %s", condition.ID), err)
			}
			continue
		}

		if matched {
			activeConditions = append(activeConditions, condition)
		}
	}

	if len(activeConditions) == 0 {
		return theme, nil
	}

	sort.Slice(activeConditions, func(i, j int) bool {
		return activeConditions[i].Priority > activeConditions[j].Priority
	})

	processedTheme := theme.Clone()

	for _, condition := range activeConditions {
		if err := m.applyOverrides(processedTheme, condition.Overrides); err != nil {
			if m.config.StrictMode {
				return nil, WrapError(ErrCodeCondition,
					fmt.Sprintf("failed to apply overrides for condition %s", condition.ID), err)
			}
		}
	}

	return processedTheme, nil
}

// applyOverrides applies token overrides to a theme.
func (m *Manager) applyOverrides(theme *Theme, overrides map[string]string) error {
	for path, value := range overrides {
		if err := m.setTokenValue(theme.Tokens, path, value); err != nil {
			return err
		}
	}
	return nil
}

// setTokenValue sets a token value at the given path.
func (m *Manager) setTokenValue(tokens *Tokens, path, value string) error {
	segments := strings.Split(path, ".")
	if len(segments) < 2 {
		return NewErrorf(ErrCodeValidation, "invalid token path: %s", path)
	}

	category := segments[0]
	
	switch category {
	case "primitives":
		return m.setPrimitiveValue(tokens.Primitives, segments[1:], value)
	case "semantic":
		return m.setSemanticValue(tokens.Semantic, segments[1:], value)
	case "components":
		return m.setComponentValue(tokens.Components, segments[1:], value)
	default:
		return NewErrorf(ErrCodeValidation, "unknown token category: %s", category)
	}
}

// setPrimitiveValue sets a primitive token value.
func (m *Manager) setPrimitiveValue(primitives *PrimitiveTokens, path []string, value string) error {
	if primitives == nil || len(path) < 2 {
		return NewError(ErrCodeValidation, "invalid primitive path")
	}

	subcategory := path[0]
	key := strings.Join(path[1:], ".")

	var tokenMap map[string]string
	switch subcategory {
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
	case "zindex":
		tokenMap = primitives.ZIndex
	case "breakpoints":
		tokenMap = primitives.Breakpoints
	default:
		return NewErrorf(ErrCodeValidation, "unknown primitive subcategory: %s", subcategory)
	}

	if tokenMap == nil {
		return NewErrorf(ErrCodeValidation, "primitive subcategory not initialized: %s", subcategory)
	}

	tokenMap[key] = value
	return nil
}

// setSemanticValue sets a semantic token value.
func (m *Manager) setSemanticValue(semantic *SemanticTokens, path []string, value string) error {
	if semantic == nil || len(path) < 2 {
		return NewError(ErrCodeValidation, "invalid semantic path")
	}

	subcategory := path[0]
	key := strings.Join(path[1:], ".")

	var tokenMap map[string]string
	switch subcategory {
	case "colors":
		tokenMap = semantic.Colors
	case "spacing":
		tokenMap = semantic.Spacing
	case "typography":
		tokenMap = semantic.Typography
	case "interactive":
		tokenMap = semantic.Interactive
	default:
		return NewErrorf(ErrCodeValidation, "unknown semantic subcategory: %s", subcategory)
	}

	if tokenMap == nil {
		return NewErrorf(ErrCodeValidation, "semantic subcategory not initialized: %s", subcategory)
	}

	tokenMap[key] = value
	return nil
}

// setComponentValue sets a component token value.
func (m *Manager) setComponentValue(components *ComponentTokens, path []string, value string) error {
	if components == nil || len(path) < 3 {
		return NewError(ErrCodeValidation, "invalid component path: needs component.variant.property")
	}

	componentName := path[0]
	variantName := path[1]
	propertyName := strings.Join(path[2:], ".")

	if (*components)[componentName] == nil {
		(*components)[componentName] = make(ComponentVariants)
	}
	if (*components)[componentName][variantName] == nil {
		(*components)[componentName][variantName] = make(StyleProperties)
	}

	(*components)[componentName][variantName][propertyName] = value
	return nil
}

// preloadThemes loads all themes from storage into memory.
func (m *Manager) preloadThemes(ctx context.Context) error {
	themes, err := m.storage.ListThemes(ctx)
	if err != nil {
		return err
	}

	m.mu.Lock()
	for _, theme := range themes {
		m.themes[theme.ID] = theme
	}
	m.mu.Unlock()

	return nil
}

// Close closes the manager and releases resources.
func (m *Manager) Close() error {
	m.resolver.Close()
	if m.compiler.cache != nil {
		m.compiler.cache.Close()
	}
	return nil
}

// CompiledTheme represents a fully compiled theme ready for use.
type CompiledTheme struct {
	Theme          *Theme
	ResolvedTokens *Tokens
	CSS            string
	DarkMode       bool
	TenantID       string
	CompiledAt     time.Time
}

// ManagerStats contains manager performance statistics.
type ManagerStats struct {
	ThemeCount     int
	ResolverStats  *ResolverStats
	CompilerStats  *CompilerStats
	ValidatorStats *ValidatorStats
}

// Context keys for theme management.
type contextKey string

const (
	contextKeyTenant   contextKey = "theme:tenant"
	contextKeyDarkMode contextKey = "theme:darkmode"
)

// WithTenant returns a new context with the tenant ID.
func WithTenant(ctx context.Context, tenantID string) context.Context {
	return context.WithValue(ctx, contextKeyTenant, tenantID)
}

// WithDarkMode returns a new context with dark mode enabled.
func WithDarkMode(ctx context.Context, enabled bool) context.Context {
	return context.WithValue(ctx, contextKeyDarkMode, enabled)
}

// getTenantID retrieves the tenant ID from context.
func getTenantID(ctx context.Context) string {
	if tenantID, ok := ctx.Value(contextKeyTenant).(string); ok {
		return tenantID
	}
	return ""
}

// getDarkMode retrieves the dark mode setting from context.
func getDarkMode(ctx context.Context) bool {
	if darkMode, ok := ctx.Value(contextKeyDarkMode).(bool); ok {
		return darkMode
	}
	return false
}
