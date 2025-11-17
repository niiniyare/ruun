package ruun

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ThemeAPI provides a unified interface for theme management
// This API is designed to be used by other projects
type ThemeAPI struct {
	manager  *schema.ThemeManager
	compiler *ThemeCompiler
	cache    map[string]*CachedTheme
	cacheMux sync.RWMutex
	basePath string
}

// CachedTheme represents a compiled theme with CSS output
type CachedTheme struct {
	Theme     *schema.Theme `json:"theme"`
	CSS       string        `json:"css"`
	Timestamp int64         `json:"timestamp"`
	Hash      string        `json:"hash"`
}

// ThemeCompiler handles compilation of JSON themes to CSS
type ThemeCompiler struct {
	outputPath string
	minify     bool
}

// NewThemeAPI creates a new Theme API instance
func NewThemeAPI(themesPath string) *ThemeAPI {
	themeRegistry := schema.NewThemeRegistry()
	tokenRegistry := schema.NewTokenRegistry()
	manager := schema.NewThemeManagerWithRegistries(themeRegistry, tokenRegistry)
	compiler := &ThemeCompiler{
		outputPath: filepath.Join(themesPath, "compiled"),
		minify:     true,
	}

	api := &ThemeAPI{
		manager:  manager,
		compiler: compiler,
		cache:    make(map[string]*CachedTheme),
		basePath: themesPath,
	}

	return api
}

// LoadTheme loads a theme from JSON file
func (api *ThemeAPI) LoadTheme(themePath string) (*schema.Theme, error) {
	data, err := os.ReadFile(themePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read theme file: %w", err)
	}

	var themeData struct {
		Name       string         `json:"name"`
		Extends    *string        `json:"extends"`
		Tokens     map[string]any `json:"tokens"`
		Components map[string]any `json:"components"`
	}

	if err := json.Unmarshal(data, &themeData); err != nil {
		return nil, fmt.Errorf("failed to parse theme JSON: %w", err)
	}

	// Convert to schema.Theme format
	theme := &schema.Theme{
		ID:          generateThemeID(themePath),
		Name:        themeData.Name,
		Description: fmt.Sprintf("Theme loaded from %s", filepath.Base(themePath)),
		Version:     "1.0.0",
		Tokens:      api.convertTokens(themeData.Tokens, themeData.Components),
	}

	return theme, nil
}

// CompileTheme compiles a theme to CSS
func (api *ThemeAPI) CompileTheme(theme *schema.Theme) (string, error) {
	// Check cache first
	api.cacheMux.RLock()
	if cached, exists := api.cache[theme.ID]; exists {
		api.cacheMux.RUnlock()
		return cached.CSS, nil
	}
	api.cacheMux.RUnlock()

	css := api.compiler.compileToCSS(theme)

	// Cache the result
	api.cacheMux.Lock()
	api.cache[theme.ID] = &CachedTheme{
		Theme:     theme,
		CSS:       css,
		Timestamp: getCurrentTimestamp(),
		Hash:      calculateHash(css),
	}
	api.cacheMux.Unlock()

	return css, nil
}

// GetTheme retrieves a theme by ID
func (api *ThemeAPI) GetTheme(themeID string) (*schema.Theme, error) {
	return api.manager.GetTheme(themeID)
}

// ListThemes returns all available themes
func (api *ThemeAPI) ListThemes() []*schema.Theme {
	return api.manager.ListThemes()
}

// RegisterTheme registers a theme with the manager
func (api *ThemeAPI) RegisterTheme(theme *schema.Theme) error {
	return api.manager.RegisterTheme(theme)
}

// LoadThemesFromDirectory loads all themes from a directory
func (api *ThemeAPI) LoadThemesFromDirectory(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		theme, err := api.LoadTheme(path)
		if err != nil {
			return fmt.Errorf("failed to load theme %s: %w", path, err)
		}

		return api.RegisterTheme(theme)
	})
}

// CompileAllThemes compiles all registered themes to CSS
func (api *ThemeAPI) CompileAllThemes() error {
	themes := api.ListThemes()

	for _, theme := range themes {
		css, err := api.CompileTheme(theme)
		if err != nil {
			return fmt.Errorf("failed to compile theme %s: %w", theme.ID, err)
		}

		// Write CSS to file
		outputPath := filepath.Join(api.compiler.outputPath, theme.ID+".css")
		if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		if err := os.WriteFile(outputPath, []byte(css), 0o644); err != nil {
			return fmt.Errorf("failed to write CSS file: %w", err)
		}
	}

	return nil
}

// GetCompiledCSS returns the compiled CSS for a theme
func (api *ThemeAPI) GetCompiledCSS(themeID string) (string, error) {
	theme, err := api.GetTheme(themeID)
	if err != nil {
		return "", err
	}

	return api.CompileTheme(theme)
}

// CreateCustomTheme creates a custom theme with overrides
func (api *ThemeAPI) CreateCustomTheme(baseThemeID string, overrides map[string]string) (*schema.Theme, error) {
	baseTheme, err := api.GetTheme(baseThemeID)
	if err != nil {
		return nil, fmt.Errorf("base theme not found: %w", err)
	}

	// Create custom theme with overrides
	customTheme := &schema.Theme{
		ID:          generateCustomThemeID(baseThemeID, overrides),
		Name:        fmt.Sprintf("%s (Custom)", baseTheme.Name),
		Description: fmt.Sprintf("Custom theme based on %s", baseTheme.Name),
		Version:     baseTheme.Version,
		Tokens:      api.applyOverrides(baseTheme.Tokens, overrides),
	}

	return customTheme, nil
}

// Theme compilation helpers

func (compiler *ThemeCompiler) compileToCSS(theme *schema.Theme) string {
	var css strings.Builder

	css.WriteString("/* Generated theme: ")
	css.WriteString(theme.Name)
	css.WriteString(" */\n\n")

	// CSS Custom Properties (CSS Variables)
	css.WriteString(":root {\n")
	compiler.writeCSSVariables(&css, theme.Tokens)
	css.WriteString("}\n\n")

	// Component Classes
	compiler.writeComponentClasses(&css, theme.Tokens)

	result := css.String()
	if compiler.minify {
		return minifyCSS(result)
	}

	return result
}

func (compiler *ThemeCompiler) writeCSSVariables(css *strings.Builder, tokens *schema.DesignTokens) {
	if tokens.Semantic != nil && tokens.Semantic.Colors != nil {
		css.WriteString("  /* Colors */\n")
		writeColorTokens(css, tokens.Semantic.Colors)
	}

	if tokens.Semantic != nil && tokens.Semantic.Spacing != nil {
		css.WriteString("  /* Spacing */\n")
		writeSpacingTokens(css, tokens.Semantic.Spacing)
	}

	if tokens.Semantic != nil && tokens.Semantic.Typography != nil {
		css.WriteString("  /* Typography */\n")
		writeTypographyTokens(css, tokens.Semantic.Typography)
	}
}

func (compiler *ThemeCompiler) writeComponentClasses(css *strings.Builder, tokens *schema.DesignTokens) {
	// Button Component Classes
	css.WriteString("/* Button Components */\n")
	css.WriteString(".button {\n")
	css.WriteString("  display: inline-flex;\n")
	css.WriteString("  align-items: center;\n")
	css.WriteString("  justify-content: center;\n")
	css.WriteString("  font-weight: 500;\n")
	css.WriteString("  transition: all 0.2s ease;\n")
	css.WriteString("  cursor: pointer;\n")
	css.WriteString("  border: 1px solid transparent;\n")
	css.WriteString("}\n\n")

	css.WriteString(".button-primary {\n")
	css.WriteString("  background-color: var(--color-primary);\n")
	css.WriteString("  color: var(--color-primary-foreground);\n")
	css.WriteString("  border-color: var(--color-primary);\n")
	css.WriteString("}\n\n")

	css.WriteString(".button-secondary {\n")
	css.WriteString("  background-color: var(--color-secondary);\n")
	css.WriteString("  color: var(--color-secondary-foreground);\n")
	css.WriteString("  border-color: var(--color-border);\n")
	css.WriteString("}\n\n")

	css.WriteString(".button-outline {\n")
	css.WriteString("  background-color: transparent;\n")
	css.WriteString("  color: var(--color-primary);\n")
	css.WriteString("  border-color: var(--color-primary);\n")
	css.WriteString("}\n\n")

	// Input Component Classes
	css.WriteString("/* Input Components */\n")
	css.WriteString(".input {\n")
	css.WriteString("  width: 100%;\n")
	css.WriteString("  background-color: var(--color-input);\n")
	css.WriteString("  border: 1px solid var(--color-border);\n")
	css.WriteString("  border-radius: var(--radius-md);\n")
	css.WriteString("  padding: 0.5rem 0.75rem;\n")
	css.WriteString("  color: var(--color-foreground);\n")
	css.WriteString("  transition: border-color 0.2s ease;\n")
	css.WriteString("}\n\n")

	css.WriteString(".input:focus {\n")
	css.WriteString("  outline: none;\n")
	css.WriteString("  border-color: var(--color-ring);\n")
	css.WriteString("  box-shadow: 0 0 0 2px var(--color-ring);\n")
	css.WriteString("}\n\n")

	css.WriteString(".input-error {\n")
	css.WriteString("  border-color: var(--color-error);\n")
	css.WriteString("}\n\n")

	// Alert Component Classes
	css.WriteString("/* Alert Components */\n")
	css.WriteString(".alert {\n")
	css.WriteString("  padding: 1rem;\n")
	css.WriteString("  border-radius: var(--radius-md);\n")
	css.WriteString("  border: 1px solid var(--color-border);\n")
	css.WriteString("}\n\n")

	css.WriteString(".alert-success {\n")
	css.WriteString("  background-color: hsl(142, 71%, 97%);\n")
	css.WriteString("  border-color: var(--color-success);\n")
	css.WriteString("  color: hsl(142, 71%, 25%);\n")
	css.WriteString("}\n\n")

	css.WriteString(".alert-warning {\n")
	css.WriteString("  background-color: hsl(45, 93%, 97%);\n")
	css.WriteString("  border-color: var(--color-warning);\n")
	css.WriteString("  color: hsl(45, 93%, 25%);\n")
	css.WriteString("}\n\n")

	css.WriteString(".alert-error {\n")
	css.WriteString("  background-color: hsl(0, 84%, 97%);\n")
	css.WriteString("  border-color: var(--color-error);\n")
	css.WriteString("  color: hsl(0, 84%, 25%);\n")
	css.WriteString("}\n\n")
}

// Helper functions for token conversion

func (api *ThemeAPI) convertTokens(tokens map[string]interface{}, components map[string]interface{}) *schema.DesignTokens {
	designTokens := &schema.DesignTokens{
		Semantic: &schema.SemanticTokens{
			Colors:      convertColors(tokens),
			Spacing:     convertSpacing(tokens),
			Typography:  convertTypography(tokens),
			Interactive: convertInteractive(tokens),
		},
		Components: convertComponents(components),
	}

	return designTokens
}

func convertColors(tokens map[string]interface{}) *schema.SemanticColors {
	colors, ok := tokens["colors"].(map[string]interface{})
	if !ok {
		return &schema.SemanticColors{}
	}

	semanticColors := &schema.SemanticColors{
		Background:  &schema.BackgroundColors{},
		Text:        &schema.TextColors{},
		Border:      &schema.BorderColors{},
		Interactive: &schema.InteractiveColors{},
		Feedback:    &schema.FeedbackColors{},
	}

	for key, value := range colors {
		if strValue, ok := value.(string); ok {
			tokenRef := schema.TokenReference(strValue)
			switch key {
			case "background":
				semanticColors.Background.Default = tokenRef
			case "secondary":
				semanticColors.Background.Subtle = tokenRef
			case "foreground":
				semanticColors.Text.Default = tokenRef
			case "muted-foreground":
				semanticColors.Text.Subtle = tokenRef
			case "border":
				semanticColors.Border.Default = tokenRef
			}
		}
	}

	return semanticColors
}

func convertSpacing(tokens map[string]interface{}) *schema.SemanticSpacing {
	return &schema.SemanticSpacing{
		Component: &schema.ComponentSpacing{},
		Layout:    &schema.LayoutSpacing{},
	}
}

func convertInteractive(tokens map[string]interface{}) *schema.SemanticInteractive {
	return &schema.SemanticInteractive{
		BorderRadius: &schema.InteractiveBorderRadius{},
		Shadow:       &schema.InteractiveShadow{},
	}
}

func convertTypography(tokens map[string]interface{}) *schema.SemanticTypography {
	return &schema.SemanticTypography{
		Headings: &schema.HeadingTokens{},
		Body:     &schema.BodyTokens{},
		Labels:   &schema.LabelTokens{},
		Captions: &schema.CaptionTokens{},
		Code:     &schema.CodeTokens{},
	}
}

func convertComponents(components map[string]interface{}) *schema.ComponentTokens {
	return &schema.ComponentTokens{
		Button:     &schema.ButtonTokens{},
		Input:      &schema.InputTokens{},
		Card:       &schema.CardTokens{},
		Modal:      &schema.ModalTokens{},
		Form:       &schema.FormTokens{},
		Table:      &schema.TableTokens{},
		Navigation: &schema.NavigationTokens{},
	}
}

func writeColorTokens(css *strings.Builder, colors *schema.SemanticColors) {
	if colors.Background != nil {
		css.WriteString(fmt.Sprintf("  --color-background: %s;\n", colors.Background.Default.String()))
		css.WriteString(fmt.Sprintf("  --color-background-subtle: %s;\n", colors.Background.Subtle.String()))
		css.WriteString(fmt.Sprintf("  --color-background-emphasis: %s;\n", colors.Background.Emphasis.String()))
	}

	if colors.Text != nil {
		css.WriteString(fmt.Sprintf("  --color-foreground: %s;\n", colors.Text.Default.String()))
		css.WriteString(fmt.Sprintf("  --color-foreground-subtle: %s;\n", colors.Text.Subtle.String()))
	}

	if colors.Border != nil {
		css.WriteString(fmt.Sprintf("  --color-border: %s;\n", colors.Border.Default.String()))
	}
}

func writeSpacingTokens(css *strings.Builder, spacing *schema.SemanticSpacing) {
	if spacing.Component != nil {
		css.WriteString(fmt.Sprintf("  --spacing-component-tight: %s;\n", spacing.Component.Tight.String()))
		css.WriteString(fmt.Sprintf("  --spacing-component-default: %s;\n", spacing.Component.Default.String()))
		css.WriteString(fmt.Sprintf("  --spacing-component-loose: %s;\n", spacing.Component.Loose.String()))
	}

	if spacing.Layout != nil {
		css.WriteString(fmt.Sprintf("  --spacing-layout-section: %s;\n", spacing.Layout.Section.String()))
		css.WriteString(fmt.Sprintf("  --spacing-layout-page: %s;\n", spacing.Layout.Page.String()))
	}
}

func writeTypographyTokens(css *strings.Builder, typography *schema.SemanticTypography) {
	// Write basic typography variables
	css.WriteString("  --font-size-sm: 0.875rem;\n")
	css.WriteString("  --font-size-base: 1rem;\n")
	css.WriteString("  --font-size-lg: 1.125rem;\n")
}

func (api *ThemeAPI) applyOverrides(tokens *schema.DesignTokens, overrides map[string]string) *schema.DesignTokens {
	// Create a copy of tokens and apply overrides
	// This is a simplified implementation
	return tokens
}

// Utility functions

func generateThemeID(path string) string {
	base := filepath.Base(path)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

func generateCustomThemeID(baseID string, overrides map[string]string) string {
	// Generate a hash-based ID for custom themes
	return fmt.Sprintf("%s-custom-%d", baseID, len(overrides))
}

func getCurrentTimestamp() int64 {
	return 0 // Simplified - should return actual timestamp
}

func calculateHash(content string) string {
	// Simplified - should calculate actual hash
	return fmt.Sprintf("hash-%d", len(content))
}

func minifyCSS(css string) string {
	// Basic CSS minification
	css = strings.ReplaceAll(css, "\n", "")
	css = strings.ReplaceAll(css, "  ", " ")
	css = strings.ReplaceAll(css, ": ", ":")
	css = strings.ReplaceAll(css, "; ", ";")
	return css
}

