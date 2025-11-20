package theme

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/niiniyare/ruun/pkg/schema"
)

// TokenResolver handles resolution of token references to actual values
type TokenResolver struct {
	api       ThemeAPIInterface
	cache     *RuntimeCache
	registry  *ResolverRegistry
	mu        sync.RWMutex
}

// ResolverRegistry provides token resolution with caching
type ResolverRegistry struct {
	primitives map[string]interface{}
	semantic   map[string]interface{}
	components map[string]interface{}
	mu         sync.RWMutex
}

// ResolvedToken represents a fully resolved token with metadata
type ResolvedToken struct {
	Value       string                 `json:"value"`
	Path        string                 `json:"path"`
	Category    TokenCategory          `json:"category"`
	Type        TokenType              `json:"type"`
	Fallback    string                 `json:"fallback,omitempty"`
	DarkValue   string                 `json:"darkValue,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Resolved    bool                   `json:"resolved"`
	Error       string                 `json:"error,omitempty"`
}

// TokenCategory represents the category of a design token
type TokenCategory string

const (
	TokenCategoryPrimitive TokenCategory = "primitive"
	TokenCategorySemantic  TokenCategory = "semantic"
	TokenCategoryComponent TokenCategory = "component"
	TokenCategoryCustom    TokenCategory = "custom"
)

// TokenType represents the type of token value
type TokenType string

const (
	TokenTypeColor       TokenType = "color"
	TokenTypeSpacing     TokenType = "spacing"
	TokenTypeTypography  TokenType = "typography"
	TokenTypeBorder      TokenType = "border"
	TokenTypeShadow      TokenType = "shadow"
	TokenTypeAnimation   TokenType = "animation"
	TokenTypeSize        TokenType = "size"
	TokenTypeZIndex      TokenType = "zindex"
	TokenTypeReference   TokenType = "reference"
	TokenTypeLiteral     TokenType = "literal"
)

// NewTokenResolver creates a new token resolver
func NewTokenResolver(api ThemeAPIInterface, cache *RuntimeCache) *TokenResolver {
	return &TokenResolver{
		api:   api,
		cache: cache,
		registry: &ResolverRegistry{
			primitives: make(map[string]interface{}),
			semantic:   make(map[string]interface{}),
			components: make(map[string]interface{}),
		},
	}
}

// ResolveToken resolves a token reference to its final value
func (tr *TokenResolver) ResolveToken(ctx context.Context, tokenRef schema.TokenReference, theme *schema.Theme, darkMode bool) (string, error) {
	if !tokenRef.IsReference() {
		// Direct value, return as-is
		return tokenRef.String(), nil
	}

	// Check cache first
	cacheKey := tr.buildCacheKey(tokenRef.Path(), theme.ID, darkMode)
	if tr.cache != nil {
		tr.cache.mu.RLock()
		if cached, exists := tr.cache.tokenValues[cacheKey]; exists {
			tr.cache.mu.RUnlock()
			return cached, nil
		}
		tr.cache.mu.RUnlock()
	}

	// Resolve token
	resolved, err := tr.resolveTokenPath(ctx, tokenRef.Path(), theme, darkMode)
	if err != nil {
		return "", fmt.Errorf("failed to resolve token %s: %w", tokenRef.Path(), err)
	}

	// Cache result
	if tr.cache != nil {
		tr.cache.mu.Lock()
		tr.cache.tokenValues[cacheKey] = resolved.Value
		tr.cache.mu.Unlock()
	}

	return resolved.Value, nil
}

// ResolveTokenWithMetadata resolves a token and returns full metadata
func (tr *TokenResolver) ResolveTokenWithMetadata(ctx context.Context, tokenRef schema.TokenReference, theme *schema.Theme, darkMode bool) (*ResolvedToken, error) {
	if !tokenRef.IsReference() {
		return &ResolvedToken{
			Value:    tokenRef.String(),
			Path:     tokenRef.String(),
			Category: TokenCategoryCustom,
			Type:     tr.inferTokenType(tokenRef.String()),
			Resolved: true,
		}, nil
	}

	return tr.resolveTokenPath(ctx, tokenRef.Path(), theme, darkMode)
}

// ResolveBatchTokens resolves multiple tokens efficiently
func (tr *TokenResolver) ResolveBatchTokens(ctx context.Context, tokens []schema.TokenReference, theme *schema.Theme, darkMode bool) (map[string]string, error) {
	results := make(map[string]string)
	
	for _, token := range tokens {
		value, err := tr.ResolveToken(ctx, token, theme, darkMode)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve token %s: %w", token.String(), err)
		}
		results[token.String()] = value
	}
	
	return results, nil
}

// GetTokenInfo returns information about a token without resolving it
func (tr *TokenResolver) GetTokenInfo(tokenPath string) *TokenInfo {
	parts := strings.Split(tokenPath, ".")
	if len(parts) < 2 {
		return &TokenInfo{
			Path:     tokenPath,
			Category: TokenCategoryCustom,
			Type:     TokenTypeLiteral,
			Valid:    false,
		}
	}

	category := tr.categorizeToken(parts[0])
	tokenType := tr.inferTokenTypeFromPath(parts)

	return &TokenInfo{
		Path:      tokenPath,
		Category:  category,
		Type:      tokenType,
		Valid:     tr.validateTokenPath(tokenPath),
		Segment:   parts[0],
		Name:      strings.Join(parts[1:], "."),
	}
}

// TokenInfo provides metadata about a token
type TokenInfo struct {
	Path     string        `json:"path"`
	Category TokenCategory `json:"category"`
	Type     TokenType     `json:"type"`
	Valid    bool          `json:"valid"`
	Segment  string        `json:"segment"`
	Name     string        `json:"name"`
}

// ValidateTokenReferences validates all token references in a theme
func (tr *TokenResolver) ValidateTokenReferences(ctx context.Context, theme *schema.Theme) (*TokenValidationReport, error) {
	report := &TokenValidationReport{
		ThemeID:     theme.ID,
		Valid:       true,
		Errors:      []TokenValidationError{},
		Warnings:    []TokenValidationWarning{},
		References:  []string{},
		Resolved:    0,
		Unresolved:  0,
	}

	// Extract all token references from the theme
	references := tr.extractTokenReferences(theme)
	report.References = references

	// Validate each reference
	for _, ref := range references {
		tokenRef := schema.TokenReference(ref)
		_, err := tr.ResolveToken(ctx, tokenRef, theme, false)
		if err != nil {
			report.Valid = false
			report.Unresolved++
			report.Errors = append(report.Errors, TokenValidationError{
				TokenPath: ref,
				Error:     err.Error(),
				Severity:  ValidationSeverityError,
			})
		} else {
			report.Resolved++
		}
	}

	return report, nil
}

// TokenValidationReport provides validation results for theme tokens
type TokenValidationReport struct {
	ThemeID    string                   `json:"themeId"`
	Valid      bool                     `json:"valid"`
	Errors     []TokenValidationError   `json:"errors"`
	Warnings   []TokenValidationWarning `json:"warnings"`
	References []string                 `json:"references"`
	Resolved   int                      `json:"resolved"`
	Unresolved int                      `json:"unresolved"`
}

// TokenValidationError represents a token validation error
type TokenValidationError struct {
	TokenPath string              `json:"tokenPath"`
	Error     string              `json:"error"`
	Severity  ValidationSeverity  `json:"severity"`
}

// TokenValidationWarning represents a token validation warning
type TokenValidationWarning struct {
	TokenPath string `json:"tokenPath"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

// ValidationSeverity indicates the severity of a validation issue
type ValidationSeverity string

const (
	ValidationSeverityError   ValidationSeverity = "error"
	ValidationSeverityWarning ValidationSeverity = "warning"
	ValidationSeverityInfo    ValidationSeverity = "info"
)

// RegisterTokens registers tokens for resolution
func (tr *TokenResolver) RegisterTokens(category string, tokens interface{}) {
	tr.registry.mu.Lock()
	defer tr.registry.mu.Unlock()

	switch category {
	case "primitives":
		tr.registry.primitives[category] = tokens
	case "semantic":
		tr.registry.semantic[category] = tokens
	case "components":
		tr.registry.components[category] = tokens
	}
}

// Private helper methods

func (tr *TokenResolver) resolveTokenPath(ctx context.Context, path string, theme *schema.Theme, darkMode bool) (*ResolvedToken, error) {
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid token path: %s", path)
	}

	category := tr.categorizeToken(parts[0])
	resolved := &ResolvedToken{
		Path:     path,
		Category: category,
		Type:     tr.inferTokenTypeFromPath(parts),
		Resolved: false,
	}

	// Navigate through the token structure
	value, err := tr.navigateTokenPath(parts, theme, darkMode)
	if err != nil {
		resolved.Error = err.Error()
		return resolved, err
	}

	resolved.Value = value
	resolved.Resolved = true

	// Recursively resolve if this value is also a reference
	if strings.Contains(value, ".") && tr.looksLikeReference(value) {
		nestedRef := schema.TokenReference(value)
		if nestedRef.IsReference() {
			nestedValue, err := tr.ResolveToken(ctx, nestedRef, theme, darkMode)
			if err != nil {
				resolved.Error = fmt.Sprintf("nested resolution failed: %s", err.Error())
				return resolved, err
			}
			resolved.Value = nestedValue
		}
	}

	return resolved, nil
}

func (tr *TokenResolver) navigateTokenPath(parts []string, theme *schema.Theme, darkMode bool) (string, error) {
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid token path")
	}

	segment := parts[0]
	remainder := parts[1:]

	switch segment {
	case "colors":
		return tr.resolveColorToken(remainder, theme, darkMode)
	case "spacing":
		return tr.resolveSpacingToken(remainder, theme)
	case "typography":
		return tr.resolveTypographyToken(remainder, theme)
	case "border", "borders":
		return tr.resolveBorderToken(remainder, theme)
	case "shadow", "shadows":
		return tr.resolveShadowToken(remainder, theme)
	case "animation", "animations":
		return tr.resolveAnimationToken(remainder, theme)
	case "size", "sizes":
		return tr.resolveSizeToken(remainder, theme)
	case "zindex", "z-index":
		return tr.resolveZIndexToken(remainder, theme)
	default:
		return tr.resolveComponentToken(segment, remainder, theme)
	}
}

func (tr *TokenResolver) resolveColorToken(path []string, theme *schema.Theme, darkMode bool) (string, error) {
	if theme.Tokens == nil || theme.Tokens.Semantic == nil || theme.Tokens.Semantic.Colors == nil {
		return "", fmt.Errorf("no color tokens available")
	}

	colors := theme.Tokens.Semantic.Colors

	// Check for dark mode overrides
	if darkMode && theme.DarkMode != nil && theme.DarkMode.DarkTokens != nil &&
		theme.DarkMode.DarkTokens.Semantic != nil && theme.DarkMode.DarkTokens.Semantic.Colors != nil {
		colors = theme.DarkMode.DarkTokens.Semantic.Colors
	}

	if len(path) == 0 {
		return "", fmt.Errorf("empty color token path")
	}

	switch path[0] {
	case "background":
		if colors.Background == nil {
			return "", fmt.Errorf("no background colors available")
		}
		return tr.resolveBackgroundColor(path[1:], colors.Background)
	case "text":
		if colors.Text == nil {
			return "", fmt.Errorf("no text colors available")
		}
		return tr.resolveTextColor(path[1:], colors.Text)
	case "border":
		if colors.Border == nil {
			return "", fmt.Errorf("no border colors available")
		}
		return tr.resolveBorderColor(path[1:], colors.Border)
	case "interactive":
		if colors.Interactive == nil {
			return "", fmt.Errorf("no interactive colors available")
		}
		return tr.resolveInteractiveColor(path[1:], colors.Interactive)
	case "feedback":
		if colors.Feedback == nil {
			return "", fmt.Errorf("no feedback colors available")
		}
		return tr.resolveFeedbackColor(path[1:], colors.Feedback)
	default:
		return "", fmt.Errorf("unknown color category: %s", path[0])
	}
}

func (tr *TokenResolver) resolveBackgroundColor(path []string, bg *schema.BackgroundColors) (string, error) {
	if len(path) == 0 || path[0] == "default" {
		return bg.Default.String(), nil
	}
	
	switch path[0] {
	case "subtle":
		return bg.Subtle.String(), nil
	case "emphasis":
		return bg.Emphasis.String(), nil
	case "overlay":
		return bg.Overlay.String(), nil
	default:
		return "", fmt.Errorf("unknown background color: %s", path[0])
	}
}

func (tr *TokenResolver) resolveTextColor(path []string, text *schema.TextColors) (string, error) {
	if len(path) == 0 || path[0] == "default" {
		return text.Default.String(), nil
	}
	
	switch path[0] {
	case "subtle":
		return text.Subtle.String(), nil
	case "emphasis":
		return text.Emphasis.String(), nil
	case "accent":
		return text.Accent.String(), nil
	default:
		return "", fmt.Errorf("unknown text color: %s", path[0])
	}
}

func (tr *TokenResolver) resolveBorderColor(path []string, border *schema.BorderColors) (string, error) {
	if len(path) == 0 || path[0] == "default" {
		return border.Default.String(), nil
	}
	
	switch path[0] {
	case "subtle":
		return border.Subtle.String(), nil
	case "emphasis":
		return border.Emphasis.String(), nil
	default:
		return "", fmt.Errorf("unknown border color: %s", path[0])
	}
}

func (tr *TokenResolver) resolveInteractiveColor(path []string, interactive *schema.InteractiveColors) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("interactive color requires a variant")
	}
	
	switch path[0] {
	case "primary":
		return interactive.Primary.String(), nil
	case "secondary":
		return interactive.Secondary.String(), nil
	case "accent":
		return interactive.Accent.String(), nil
	case "hover":
		return interactive.Hover.String(), nil
	case "active":
		return interactive.Active.String(), nil
	case "focus":
		return interactive.Focus.String(), nil
	case "disabled":
		return interactive.Disabled.String(), nil
	default:
		return "", fmt.Errorf("unknown interactive color: %s", path[0])
	}
}

func (tr *TokenResolver) resolveFeedbackColor(path []string, feedback *schema.FeedbackColors) (string, error) {
	if len(path) == 0 {
		return "", fmt.Errorf("feedback color requires a type")
	}
	
	switch path[0] {
	case "success":
		return feedback.Success.String(), nil
	case "warning":
		return feedback.Warning.String(), nil
	case "error":
		return feedback.Error.String(), nil
	case "info":
		return feedback.Info.String(), nil
	default:
		return "", fmt.Errorf("unknown feedback color: %s", path[0])
	}
}

func (tr *TokenResolver) resolveSpacingToken(path []string, theme *schema.Theme) (string, error) {
	if theme.Tokens == nil || theme.Tokens.Semantic == nil || theme.Tokens.Semantic.Spacing == nil {
		return "", fmt.Errorf("no spacing tokens available")
	}
	
	// For now, return a default spacing value
	// This would be properly implemented with the spacing structure
	return "1rem", nil
}

func (tr *TokenResolver) resolveTypographyToken(path []string, theme *schema.Theme) (string, error) {
	// Implement typography token resolution
	return "1rem", nil
}

func (tr *TokenResolver) resolveBorderToken(path []string, theme *schema.Theme) (string, error) {
	// Implement border token resolution
	return "1px solid #ccc", nil
}

func (tr *TokenResolver) resolveShadowToken(path []string, theme *schema.Theme) (string, error) {
	// Implement shadow token resolution
	return "0 1px 3px rgba(0,0,0,0.1)", nil
}

func (tr *TokenResolver) resolveAnimationToken(path []string, theme *schema.Theme) (string, error) {
	// Implement animation token resolution
	return "0.2s ease", nil
}

func (tr *TokenResolver) resolveSizeToken(path []string, theme *schema.Theme) (string, error) {
	// Implement size token resolution
	return "1rem", nil
}

func (tr *TokenResolver) resolveZIndexToken(path []string, theme *schema.Theme) (string, error) {
	// Implement z-index token resolution
	return "1", nil
}

func (tr *TokenResolver) resolveComponentToken(component string, path []string, theme *schema.Theme) (string, error) {
	// Implement component-specific token resolution
	return "", fmt.Errorf("component tokens not yet implemented for: %s", component)
}

func (tr *TokenResolver) buildCacheKey(path, themeID string, darkMode bool) string {
	if darkMode {
		return fmt.Sprintf("%s:%s:dark", themeID, path)
	}
	return fmt.Sprintf("%s:%s", themeID, path)
}

func (tr *TokenResolver) categorizeToken(segment string) TokenCategory {
	primitiveSegments := map[string]bool{
		"colors": true, "spacing": true, "typography": true, "border": true, "borders": true,
		"shadow": true, "shadows": true, "animation": true, "animations": true, "size": true, "sizes": true,
		"zindex": true, "z-index": true,
	}
	
	if primitiveSegments[segment] {
		return TokenCategorySemantic
	}
	
	return TokenCategoryComponent
}

func (tr *TokenResolver) inferTokenType(value string) TokenType {
	if strings.HasPrefix(value, "#") || strings.HasPrefix(value, "hsl") || strings.HasPrefix(value, "rgb") {
		return TokenTypeColor
	}
	if strings.HasSuffix(value, "rem") || strings.HasSuffix(value, "px") || strings.HasSuffix(value, "em") {
		return TokenTypeSpacing
	}
	if strings.Contains(value, ".") && tr.looksLikeReference(value) {
		return TokenTypeReference
	}
	return TokenTypeLiteral
}

func (tr *TokenResolver) inferTokenTypeFromPath(parts []string) TokenType {
	if len(parts) == 0 {
		return TokenTypeLiteral
	}
	
	switch parts[0] {
	case "colors":
		return TokenTypeColor
	case "spacing":
		return TokenTypeSpacing
	case "typography":
		return TokenTypeTypography
	case "border", "borders":
		return TokenTypeBorder
	case "shadow", "shadows":
		return TokenTypeShadow
	case "animation", "animations":
		return TokenTypeAnimation
	case "size", "sizes":
		return TokenTypeSize
	case "zindex", "z-index":
		return TokenTypeZIndex
	default:
		return TokenTypeReference
	}
}

func (tr *TokenResolver) validateTokenPath(path string) bool {
	parts := strings.Split(path, ".")
	return len(parts) >= 2
}

func (tr *TokenResolver) looksLikeReference(value string) bool {
	return strings.Contains(value, ".") && !strings.HasPrefix(value, "#") &&
		!strings.HasPrefix(value, "hsl") && !strings.HasPrefix(value, "rgb") &&
		!strings.Contains(value, " ")
}

func (tr *TokenResolver) extractTokenReferences(theme *schema.Theme) []string {
	var references []string
	// This would extract all token references from the theme structure
	// For now, return an empty slice
	return references
}