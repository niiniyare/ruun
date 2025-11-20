package theme

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ComponentIntegrator provides utilities for integrating themes with components
type ComponentIntegrator struct {
	runtime      *Runtime
	resolver     *TokenResolver
	classBuilder *ClassBuilder
	config       *IntegrationConfig
	cache        *IntegrationCache
}

// IntegrationConfig configures component-theme integration
type IntegrationConfig struct {
	EnableClassGeneration  bool     `json:"enableClassGeneration"`
	EnableUtilityClasses   bool     `json:"enableUtilityClasses"`
	ComponentPrefix        string   `json:"componentPrefix"`
	UtilityPrefix          string   `json:"utilityPrefix"`
	EnableVariantClasses   bool     `json:"enableVariantClasses"`
	EnableStateClasses     bool     `json:"enableStateClasses"`
	EnableResponsiveClasses bool    `json:"enableResponsiveClasses"`
	SupportedComponents    []string `json:"supportedComponents"`
}

// IntegrationCache caches component integration results
type IntegrationCache struct {
	componentClasses map[string]string
	utilityClasses   map[string]string
	variantClasses   map[string]map[string]string
	stateClasses     map[string]map[string]string
	mu               sync.RWMutex
}

// ClassBuilder builds CSS classes for components
type ClassBuilder struct {
	config       *IntegrationConfig
	tokenResolver *TokenResolver
	currentTheme *schema.Theme
}

// ComponentStyle represents styling for a component
type ComponentStyle struct {
	Component string                 `json:"component"`
	Variant   string                 `json:"variant,omitempty"`
	State     string                 `json:"state,omitempty"`
	Classes   []string               `json:"classes"`
	Styles    map[string]string      `json:"styles"`
	Variables map[string]string      `json:"variables"`
	Tokens    map[string]any `json:"tokens"`
}

// StyleResult represents the result of style generation
type StyleResult struct {
	ClassName   string            `json:"className"`
	InlineCSS   string            `json:"inlineCSS,omitempty"`
	Variables   map[string]string `json:"variables,omitempty"`
	Tokens      []string          `json:"tokens,omitempty"`
	Generated   bool              `json:"generated"`
	Cached      bool              `json:"cached"`
}

// ComponentStyleConfig configures component styling
type ComponentStyleConfig struct {
	Component    string                 `json:"component"`
	Variant      string                 `json:"variant,omitempty"`
	Size         string                 `json:"size,omitempty"`
	State        string                 `json:"state,omitempty"`
	CustomTokens map[string]string      `json:"customTokens,omitempty"`
	Modifiers    []string               `json:"modifiers,omitempty"`
	Context      map[string]any `json:"context,omitempty"`
}

// NewComponentIntegrator creates a new component integrator
func NewComponentIntegrator(runtime *Runtime, resolver *TokenResolver, config *IntegrationConfig) *ComponentIntegrator {
	if config == nil {
		config = DefaultIntegrationConfig()
	}

	return &ComponentIntegrator{
		runtime:      runtime,
		resolver:     resolver,
		config:       config,
		cache:        NewIntegrationCache(),
		classBuilder: NewClassBuilder(config, resolver),
	}
}

// DefaultIntegrationConfig returns default integration configuration
func DefaultIntegrationConfig() *IntegrationConfig {
	return &IntegrationConfig{
		EnableClassGeneration:   true,
		EnableUtilityClasses:    true,
		ComponentPrefix:         "",
		UtilityPrefix:           "",
		EnableVariantClasses:    true,
		EnableStateClasses:      true,
		EnableResponsiveClasses: true,
		SupportedComponents: []string{
			"button", "input", "card", "modal", "form", "table", "navigation",
		},
	}
}

// NewIntegrationCache creates a new integration cache
func NewIntegrationCache() *IntegrationCache {
	return &IntegrationCache{
		componentClasses: make(map[string]string),
		utilityClasses:   make(map[string]string),
		variantClasses:   make(map[string]map[string]string),
		stateClasses:     make(map[string]map[string]string),
	}
}

// NewClassBuilder creates a new class builder
func NewClassBuilder(config *IntegrationConfig, resolver *TokenResolver) *ClassBuilder {
	return &ClassBuilder{
		config:        config,
		tokenResolver: resolver,
	}
}

// GenerateComponentStyle generates styles for a component
func (ci *ComponentIntegrator) GenerateComponentStyle(ctx context.Context, config *ComponentStyleConfig) (*StyleResult, error) {
	// Check cache first
	cacheKey := ci.buildCacheKey(config)
	if cached := ci.getCachedStyle(cacheKey); cached != nil {
		return cached, nil
	}

	// Generate new style
	result, err := ci.generateStyle(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to generate component style: %w", err)
	}

	// Cache result
	ci.cacheStyle(cacheKey, result)

	return result, nil
}

// GetComponentClasses returns CSS classes for a component
func (ci *ComponentIntegrator) GetComponentClasses(component, variant string) string {
	ci.cache.mu.RLock()
	defer ci.cache.mu.RUnlock()

	// Base component class
	baseClass := ci.getBaseComponentClass(component)
	classes := []string{baseClass}

	// Add variant class
	if variant != "" {
		variantClass := ci.getVariantClass(component, variant)
		if variantClass != "" {
			classes = append(classes, variantClass)
		}
	}

	return strings.Join(classes, " ")
}

// GetUtilityClasses returns utility classes for styling properties
func (ci *ComponentIntegrator) GetUtilityClasses(properties map[string]string) []string {
	var classes []string

	for property, value := range properties {
		utilityClass := ci.generateUtilityClass(property, value)
		if utilityClass != "" {
			classes = append(classes, utilityClass)
		}
	}

	return classes
}

// ResolveComponentTokens resolves component tokens for a component
func (ci *ComponentIntegrator) ResolveComponentTokens(ctx context.Context, component string) (map[string]string, error) {
	theme := ci.runtime.GetCurrentTheme()
	if theme == nil {
		return nil, fmt.Errorf("no current theme set")
	}

	// Get current theme
	currentTheme, err := ci.runtime.api.GetTheme(theme.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get current theme: %w", err)
	}

	tokens := make(map[string]string)

	// Get component tokens
	if currentTheme.Tokens != nil && currentTheme.Tokens.Components != nil {
		componentTokens := ci.getComponentTokens(currentTheme.Tokens.Components, component)
		
		for tokenPath, tokenRef := range componentTokens {
			resolved, err := ci.resolver.ResolveToken(ctx, tokenRef, currentTheme, false)
			if err != nil {
				// Log warning but continue
				continue
			}
			tokens[tokenPath] = resolved
		}
	}

	return tokens, nil
}

// GenerateInlineStyles generates inline CSS styles for a component
func (ci *ComponentIntegrator) GenerateInlineStyles(ctx context.Context, config *ComponentStyleConfig) (string, error) {
	tokens, err := ci.ResolveComponentTokens(ctx, config.Component)
	if err != nil {
		return "", err
	}

	var styles []string

	// Apply component-specific styles
	componentStyles := ci.getComponentStyles(config.Component, config.Variant)
	for property, value := range componentStyles {
		// Resolve token references in values
		resolvedValue := ci.resolveTokenInValue(ctx, value, tokens)
		styles = append(styles, fmt.Sprintf("%s: %s", property, resolvedValue))
	}

	// Apply custom tokens
	for property, value := range config.CustomTokens {
		resolvedValue := ci.resolveTokenInValue(ctx, value, tokens)
		styles = append(styles, fmt.Sprintf("%s: %s", property, resolvedValue))
	}

	return strings.Join(styles, "; "), nil
}

// GetComponentStateClasses returns state-specific classes
func (ci *ComponentIntegrator) GetComponentStateClasses(component, state string) string {
	if !ci.config.EnableStateClasses {
		return ""
	}

	ci.cache.mu.RLock()
	defer ci.cache.mu.RUnlock()

	if stateClasses, exists := ci.cache.stateClasses[component]; exists {
		if stateClass, exists := stateClasses[state]; exists {
			return stateClass
		}
	}

	// Generate state class
	stateClass := ci.generateStateClass(component, state)
	
	// Cache result
	ci.cache.mu.RUnlock()
	ci.cache.mu.Lock()
	if ci.cache.stateClasses[component] == nil {
		ci.cache.stateClasses[component] = make(map[string]string)
	}
	ci.cache.stateClasses[component][state] = stateClass
	ci.cache.mu.Unlock()
	ci.cache.mu.RLock()

	return stateClass
}

// GetResponsiveClasses returns responsive utility classes
func (ci *ComponentIntegrator) GetResponsiveClasses(breakpoint, property, value string) string {
	if !ci.config.EnableResponsiveClasses {
		return ""
	}

	return fmt.Sprintf("%s:%s", breakpoint, ci.generateUtilityClass(property, value))
}

// UpdateThemeContext updates the theme context for component integration
func (ci *ComponentIntegrator) UpdateThemeContext(theme *schema.Theme) error {
	ci.classBuilder.currentTheme = theme

	// Clear cache when theme changes
	ci.clearCache()

	return nil
}

// Private helper methods

func (ci *ComponentIntegrator) generateStyle(ctx context.Context, config *ComponentStyleConfig) (*StyleResult, error) {
	result := &StyleResult{
		Variables: make(map[string]string),
		Tokens:    []string{},
		Generated: true,
		Cached:    false,
	}

	// Build class name
	className := ci.buildClassName(config)
	result.ClassName = className

	// Generate inline CSS if needed
	if config.CustomTokens != nil || len(config.Modifiers) > 0 {
		inlineCSS, err := ci.GenerateInlineStyles(ctx, config)
		if err != nil {
			return nil, err
		}
		result.InlineCSS = inlineCSS
	}

	// Get component tokens
	tokens, err := ci.ResolveComponentTokens(ctx, config.Component)
	if err != nil {
		return nil, err
	}
	result.Variables = tokens

	// Extract token references
	for tokenPath := range tokens {
		result.Tokens = append(result.Tokens, tokenPath)
	}

	return result, nil
}

func (ci *ComponentIntegrator) buildClassName(config *ComponentStyleConfig) string {
	var classes []string

	// Base component class
	baseClass := ci.getBaseComponentClass(config.Component)
	classes = append(classes, baseClass)

	// Variant class
	if config.Variant != "" {
		variantClass := ci.getVariantClass(config.Component, config.Variant)
		if variantClass != "" {
			classes = append(classes, variantClass)
		}
	}

	// Size class
	if config.Size != "" {
		sizeClass := ci.getSizeClass(config.Component, config.Size)
		if sizeClass != "" {
			classes = append(classes, sizeClass)
		}
	}

	// State class
	if config.State != "" {
		stateClass := ci.getStateClass(config.Component, config.State)
		if stateClass != "" {
			classes = append(classes, stateClass)
		}
	}

	// Modifier classes
	for _, modifier := range config.Modifiers {
		modifierClass := ci.getModifierClass(config.Component, modifier)
		if modifierClass != "" {
			classes = append(classes, modifierClass)
		}
	}

	return strings.Join(classes, " ")
}

func (ci *ComponentIntegrator) getBaseComponentClass(component string) string {
	if ci.config.ComponentPrefix != "" {
		return fmt.Sprintf("%s-%s", ci.config.ComponentPrefix, component)
	}
	return component
}

func (ci *ComponentIntegrator) getVariantClass(component, variant string) string {
	if ci.config.EnableVariantClasses {
		return fmt.Sprintf("%s-%s", ci.getBaseComponentClass(component), variant)
	}
	return ""
}

func (ci *ComponentIntegrator) getSizeClass(component, size string) string {
	return fmt.Sprintf("%s-%s", ci.getBaseComponentClass(component), size)
}

func (ci *ComponentIntegrator) getStateClass(component, state string) string {
	return fmt.Sprintf("%s-%s", ci.getBaseComponentClass(component), state)
}

func (ci *ComponentIntegrator) getModifierClass(component, modifier string) string {
	return fmt.Sprintf("%s-%s", ci.getBaseComponentClass(component), modifier)
}

func (ci *ComponentIntegrator) generateUtilityClass(property, value string) string {
	// Map CSS properties to utility class patterns
	propertyMap := map[string]string{
		"color":            "text",
		"background-color": "bg",
		"border-color":     "border",
		"padding":          "p",
		"margin":           "m",
		"font-size":        "text",
		"font-weight":      "font",
		"border-radius":    "rounded",
		"box-shadow":       "shadow",
	}

	if prefix, exists := propertyMap[property]; exists {
		// Convert value to utility class name
		classValue := ci.valueToClassName(value)
		if ci.config.UtilityPrefix != "" {
			return fmt.Sprintf("%s-%s-%s", ci.config.UtilityPrefix, prefix, classValue)
		}
		return fmt.Sprintf("%s-%s", prefix, classValue)
	}

	return ""
}

func (ci *ComponentIntegrator) generateStateClass(component, state string) string {
	stateMap := map[string]string{
		"hover":    "hover",
		"focus":    "focus",
		"active":   "active",
		"disabled": "disabled",
		"loading":  "loading",
		"error":    "error",
		"success":  "success",
	}

	if stateClass, exists := stateMap[state]; exists {
		return fmt.Sprintf("%s-%s", ci.getBaseComponentClass(component), stateClass)
	}

	return ""
}

func (ci *ComponentIntegrator) valueToClassName(value string) string {
	// Convert CSS values to class-friendly names
	value = strings.ReplaceAll(value, "var(--", "")
	value = strings.ReplaceAll(value, ")", "")
	value = strings.ReplaceAll(value, ".", "-")
	value = strings.ReplaceAll(value, " ", "-")
	return strings.ToLower(value)
}

func (ci *ComponentIntegrator) getComponentTokens(componentTokens *schema.ComponentTokens, component string) map[string]schema.TokenReference {
	tokens := make(map[string]schema.TokenReference)

	switch component {
	case "button":
		if componentTokens.Button != nil {
			tokens["background"] = componentTokens.Button.Background
			tokens["border"] = componentTokens.Button.Border
			tokens["text"] = componentTokens.Button.Text
			tokens["borderRadius"] = componentTokens.Button.BorderRadius
			tokens["padding"] = componentTokens.Button.Padding
			tokens["fontSize"] = componentTokens.Button.FontSize
			tokens["fontWeight"] = componentTokens.Button.FontWeight
		}
	case "input":
		if componentTokens.Input != nil {
			tokens["background"] = componentTokens.Input.Background
			tokens["border"] = componentTokens.Input.Border
			tokens["text"] = componentTokens.Input.Text
			tokens["placeholder"] = componentTokens.Input.Placeholder
			tokens["borderRadius"] = componentTokens.Input.BorderRadius
			tokens["padding"] = componentTokens.Input.Padding
			tokens["fontSize"] = componentTokens.Input.FontSize
		}
	case "card":
		if componentTokens.Card != nil {
			tokens["background"] = componentTokens.Card.Background
			tokens["border"] = componentTokens.Card.Border
			tokens["shadow"] = componentTokens.Card.Shadow
			tokens["borderRadius"] = componentTokens.Card.BorderRadius
			tokens["padding"] = componentTokens.Card.Padding
		}
	case "modal":
		if componentTokens.Modal != nil {
			tokens["background"] = componentTokens.Modal.Background
			tokens["overlay"] = componentTokens.Modal.Overlay
			tokens["border"] = componentTokens.Modal.Border
			tokens["shadow"] = componentTokens.Modal.Shadow
			tokens["borderRadius"] = componentTokens.Modal.BorderRadius
			tokens["padding"] = componentTokens.Modal.Padding
		}
	case "table":
		if componentTokens.Table != nil {
			tokens["background"] = componentTokens.Table.Background
			tokens["alternateBg"] = componentTokens.Table.AlternateBg
			tokens["border"] = componentTokens.Table.Border
			tokens["headerBg"] = componentTokens.Table.HeaderBg
			tokens["headerText"] = componentTokens.Table.HeaderText
			tokens["cellPadding"] = componentTokens.Table.CellPadding
		}
	case "navigation":
		if componentTokens.Navigation != nil {
			tokens["background"] = componentTokens.Navigation.Background
			tokens["border"] = componentTokens.Navigation.Border
			tokens["linkColor"] = componentTokens.Navigation.LinkColor
			tokens["activeColor"] = componentTokens.Navigation.ActiveColor
			tokens["padding"] = componentTokens.Navigation.Padding
		}
	}

	return tokens
}

func (ci *ComponentIntegrator) getComponentStyles(component, variant string) map[string]string {
	styles := make(map[string]string)

	// This would return CSS properties for specific components and variants
	// For now, return empty map
	return styles
}

func (ci *ComponentIntegrator) resolveTokenInValue(ctx context.Context, value string, tokens map[string]string) string {
	// Look for token references in the value
	if strings.Contains(value, "var(--") {
		// Extract variable name
		start := strings.Index(value, "var(--") + 6
		end := strings.Index(value[start:], ")")
		if end != -1 {
			varName := value[start : start+end]
			if resolvedValue, exists := tokens[varName]; exists {
				return strings.ReplaceAll(value, fmt.Sprintf("var(--%s)", varName), resolvedValue)
			}
		}
	}

	return value
}

func (ci *ComponentIntegrator) buildCacheKey(config *ComponentStyleConfig) string {
	key := fmt.Sprintf("%s-%s-%s-%s", config.Component, config.Variant, config.Size, config.State)
	
	// Add modifiers
	if len(config.Modifiers) > 0 {
		key += "-" + strings.Join(config.Modifiers, "-")
	}
	
	return key
}

func (ci *ComponentIntegrator) getCachedStyle(key string) *StyleResult {
	// This would implement cache lookup
	return nil
}

func (ci *ComponentIntegrator) cacheStyle(key string, result *StyleResult) {
	// This would implement cache storage
	result.Cached = true
}

func (ci *ComponentIntegrator) clearCache() {
	ci.cache.mu.Lock()
	defer ci.cache.mu.Unlock()
	
	ci.cache.componentClasses = make(map[string]string)
	ci.cache.utilityClasses = make(map[string]string)
	ci.cache.variantClasses = make(map[string]map[string]string)
	ci.cache.stateClasses = make(map[string]map[string]string)
}

// Public utility methods for component templates

// GetButtonClasses returns classes for button components
func (ci *ComponentIntegrator) GetButtonClasses(variant, size string) string {
	classes := []string{"button"}
	
	if variant != "" {
		classes = append(classes, fmt.Sprintf("button-%s", variant))
	}
	
	if size != "" {
		classes = append(classes, fmt.Sprintf("button-%s", size))
	}
	
	return strings.Join(classes, " ")
}

// GetInputClasses returns classes for input components
func (ci *ComponentIntegrator) GetInputClasses(state string) string {
	classes := []string{"input"}
	
	if state != "" {
		classes = append(classes, fmt.Sprintf("input-%s", state))
	}
	
	return strings.Join(classes, " ")
}

// GetCardClasses returns classes for card components
func (ci *ComponentIntegrator) GetCardClasses(variant string) string {
	classes := []string{"card"}
	
	if variant != "" {
		classes = append(classes, fmt.Sprintf("card-%s", variant))
	}
	
	return strings.Join(classes, " ")
}

// GetModalClasses returns classes for modal components
func (ci *ComponentIntegrator) GetModalClasses(size string) string {
	classes := []string{"modal"}
	
	if size != "" {
		classes = append(classes, fmt.Sprintf("modal-%s", size))
	}
	
	return strings.Join(classes, " ")
}

// IsComponentSupported checks if a component is supported
func (ci *ComponentIntegrator) IsComponentSupported(component string) bool {
	for _, supported := range ci.config.SupportedComponents {
		if supported == component {
			return true
		}
	}
	return false
}

// GetSupportedComponents returns list of supported components
func (ci *ComponentIntegrator) GetSupportedComponents() []string {
	return ci.config.SupportedComponents
}

// GetIntegrationStats returns integration statistics
func (ci *ComponentIntegrator) GetIntegrationStats() map[string]any {
	ci.cache.mu.RLock()
	defer ci.cache.mu.RUnlock()

	return map[string]any{
		"componentClasses":    len(ci.cache.componentClasses),
		"utilityClasses":      len(ci.cache.utilityClasses),
		"variantClasses":      len(ci.cache.variantClasses),
		"stateClasses":        len(ci.cache.stateClasses),
		"supportedComponents": len(ci.config.SupportedComponents),
	}
}