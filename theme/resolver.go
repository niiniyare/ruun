package theme

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/dgraph-io/ristretto"
)

// Resolver resolves token references to their final values with caching,
// circular reference detection, and multi-level resolution support.
type Resolver struct {
	cache          *ristretto.Cache
	cacheTTL       time.Duration
	maxDepth       int
	strictMode     bool
	mu             sync.RWMutex
	resolutionPath []string
}

// ResolverConfig configures the token resolver behavior.
type ResolverConfig struct {
	MaxDepth      int
	StrictMode    bool
	EnableCaching bool
	CacheTTL      time.Duration
	CacheSize     int64
	CacheBuffers  int64
}

// DefaultResolverConfig returns production-ready resolver configuration.
func DefaultResolverConfig() *ResolverConfig {
	return &ResolverConfig{
		MaxDepth:      10,
		StrictMode:    false,
		EnableCaching: true,
		CacheTTL:      5 * time.Minute,
		CacheSize:     10 << 20, // 10MB
		CacheBuffers:  64,
	}
}

// ResolvedToken contains the resolved value with metadata.
type ResolvedToken struct {
	Value       string
	OriginalRef string
	Path        []string
	Depth       int
	FromCache   bool
	Duration    time.Duration
}

// NewResolver creates a new token resolver with the given configuration.
func NewResolver(config *ResolverConfig) (*Resolver, error) {
	if config == nil {
		config = DefaultResolverConfig()
	}

	var cache *ristretto.Cache
	var err error

	if config.EnableCaching {
		cache, err = ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e4,
			MaxCost:     config.CacheSize,
			BufferItems: config.CacheBuffers,
			Metrics:     true,
		})
		if err != nil {
			return nil, WrapError(ErrCodeResolution, "failed to create cache", err)
		}
	}

	return &Resolver{
		cache:      cache,
		cacheTTL:   config.CacheTTL,
		maxDepth:   config.MaxDepth,
		strictMode: config.StrictMode,
	}, nil
}

// Resolve resolves a token reference to its final value.
func (r *Resolver) Resolve(ctx context.Context, ref TokenReference, theme *Theme, tenantID string, darkMode bool) (*ResolvedToken, error) {
	start := time.Now()

	if err := ref.Validate(); err != nil {
		return nil, WrapError(ErrCodeValidation, "invalid token reference", err)
	}

	if !ref.IsReference() {
		return &ResolvedToken{
			Value:       string(ref),
			OriginalRef: string(ref),
			Path:        []string{},
			Depth:       0,
			FromCache:   false,
			Duration:    time.Since(start),
		}, nil
	}

	cacheKey := r.buildCacheKey(string(ref), theme.ID, tenantID, darkMode)
	if cached, found := r.getFromCache(cacheKey); found {
		cached.Duration = time.Since(start)
		return cached, nil
	}

	r.mu.Lock()
	r.resolutionPath = make([]string, 0, r.maxDepth)
	r.mu.Unlock()

	tokens := theme.Tokens
	if darkMode && theme.DarkMode != nil && theme.DarkMode.Enabled && theme.DarkMode.DarkTokens != nil {
		tokens = r.mergeTokens(tokens, theme.DarkMode.DarkTokens)
	}

	value, path, depth, err := r.resolveRecursive(ctx, string(ref), tokens, 0)
	if err != nil {
		return nil, err
	}

	result := &ResolvedToken{
		Value:       value,
		OriginalRef: string(ref),
		Path:        path,
		Depth:       depth,
		FromCache:   false,
		Duration:    time.Since(start),
	}

	r.setCache(cacheKey, result)

	return result, nil
}

// ResolveAll resolves all token references in the theme to their final values.
func (r *Resolver) ResolveAll(ctx context.Context, theme *Theme, tenantID string, darkMode bool) (*Tokens, error) {
	resolved := theme.Tokens.Clone()

	if err := r.resolveTokenCategory(ctx, resolved.Primitives, theme, tenantID, darkMode); err != nil {
		return nil, WrapError(ErrCodeResolution, "failed to resolve primitives", err)
	}

	if err := r.resolveSemanticTokens(ctx, resolved.Semantic, theme, tenantID, darkMode); err != nil {
		return nil, WrapError(ErrCodeResolution, "failed to resolve semantic tokens", err)
	}

	if err := r.resolveComponentTokens(ctx, resolved.Components, theme, tenantID, darkMode); err != nil {
		return nil, WrapError(ErrCodeResolution, "failed to resolve component tokens", err)
	}

	return resolved, nil
}

// resolveRecursive performs recursive token resolution with circular reference detection.
func (r *Resolver) resolveRecursive(ctx context.Context, refStr string, tokens *Tokens, depth int) (string, []string, int, error) {
	if depth >= r.maxDepth {
		return "", nil, depth, NewErrorf(ErrCodeResolution,
			"maximum resolution depth (%d) exceeded for token: %s", r.maxDepth, refStr)
	}

	r.mu.Lock()
	for _, visited := range r.resolutionPath {
		if visited == refStr {
			r.mu.Unlock()
			return "", nil, depth, NewErrorf(ErrCodeCircularRef,
				"circular reference detected: %s", strings.Join(append(r.resolutionPath, refStr), " → "))
		}
	}
	r.resolutionPath = append(r.resolutionPath, refStr)
	r.mu.Unlock()

	defer func() {
		r.mu.Lock()
		if len(r.resolutionPath) > 0 {
			r.resolutionPath = r.resolutionPath[:len(r.resolutionPath)-1]
		}
		r.mu.Unlock()
	}()

	value, err := r.getTokenValue(refStr, tokens)
	if err != nil {
		if r.strictMode {
			return "", nil, depth, err
		}
		return refStr, []string{refStr}, depth, nil
	}

	ref := TokenReference(value)
	if !ref.IsReference() {
		return value, []string{refStr}, depth, nil
	}

	resolvedValue, path, finalDepth, err := r.resolveRecursive(ctx, value, tokens, depth+1)
	if err != nil {
		return "", nil, finalDepth, err
	}

	fullPath := append([]string{refStr}, path...)
	return resolvedValue, fullPath, finalDepth + 1, nil
}

// getTokenValue retrieves the token value from the token hierarchy.
func (r *Resolver) getTokenValue(path string, tokens *Tokens) (string, error) {
	segments := strings.Split(path, ".")
	if len(segments) < 2 {
		return "", NewErrorf(ErrCodeResolution, "invalid token path: %s", path)
	}

	category := segments[0]
	subcategory := segments[1]

	switch category {
	case "primitives":
		return r.getPrimitiveValue(subcategory, segments[2:], tokens.Primitives)
	case "semantic":
		return r.getSemanticValue(subcategory, segments[2:], tokens.Semantic)
	case "components":
		return r.getComponentValue(segments[1:], tokens.Components)
	default:
		return "", NewErrorf(ErrCodeResolution, "unknown token category: %s", category)
	}
}

// getPrimitiveValue retrieves a primitive token value.
func (r *Resolver) getPrimitiveValue(category string, path []string, primitives *PrimitiveTokens) (string, error) {
	if primitives == nil {
		return "", NewErrorf(ErrCodeNotFound, "primitives not defined")
	}

	if len(path) == 0 {
		return "", NewErrorf(ErrCodeResolution, "incomplete primitive path")
	}

	key := strings.Join(path, ".")

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
	case "zindex":
		tokenMap = primitives.ZIndex
	case "breakpoints":
		tokenMap = primitives.Breakpoints
	default:
		return "", NewErrorf(ErrCodeResolution, "unknown primitive category: %s", category)
	}

	if tokenMap == nil {
		return "", NewErrorf(ErrCodeNotFound, "primitive category %s not defined", category)
	}

	value, ok := tokenMap[key]
	if !ok {
		return "", NewErrorf(ErrCodeNotFound, "primitive token not found: %s.%s", category, key)
	}

	return value, nil
}

// getSemanticValue retrieves a semantic token value.
func (r *Resolver) getSemanticValue(category string, path []string, semantic *SemanticTokens) (string, error) {
	if semantic == nil {
		return "", NewErrorf(ErrCodeNotFound, "semantic tokens not defined")
	}

	if len(path) == 0 {
		return "", NewErrorf(ErrCodeResolution, "incomplete semantic path")
	}

	key := strings.Join(path, ".")

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
		return "", NewErrorf(ErrCodeResolution, "unknown semantic category: %s", category)
	}

	if tokenMap == nil {
		return "", NewErrorf(ErrCodeNotFound, "semantic category %s not defined", category)
	}

	value, ok := tokenMap[key]
	if !ok {
		return "", NewErrorf(ErrCodeNotFound, "semantic token not found: %s.%s", category, key)
	}

	return value, nil
}

// getComponentValue retrieves a component token value.
func (r *Resolver) getComponentValue(path []string, components *ComponentTokens) (string, error) {
	if components == nil || len(*components) == 0 {
		return "", NewErrorf(ErrCodeNotFound, "component tokens not defined")
	}

	if len(path) < 3 {
		return "", NewErrorf(ErrCodeResolution, "incomplete component path: needs component.variant.property")
	}

	componentName := path[0]
	variantName := path[1]
	propertyName := strings.Join(path[2:], ".")

	variants, ok := (*components)[componentName]
	if !ok {
		return "", NewErrorf(ErrCodeNotFound, "component not found: %s", componentName)
	}

	properties, ok := variants[variantName]
	if !ok {
		return "", NewErrorf(ErrCodeNotFound, "variant not found: %s.%s", componentName, variantName)
	}

	value, ok := properties[propertyName]
	if !ok {
		return "", NewErrorf(ErrCodeNotFound, "property not found: %s.%s.%s", componentName, variantName, propertyName)
	}

	return value, nil
}

// resolveTokenCategory resolves all tokens in a primitive token category.
func (r *Resolver) resolveTokenCategory(ctx context.Context, primitives *PrimitiveTokens, theme *Theme, tenantID string, darkMode bool) error {
	if primitives == nil {
		return nil
	}

	categories := map[string]*map[string]string{
		"colors":      &primitives.Colors,
		"spacing":     &primitives.Spacing,
		"radius":      &primitives.Radius,
		"typography":  &primitives.Typography,
		"borders":     &primitives.Borders,
		"shadows":     &primitives.Shadows,
		"effects":     &primitives.Effects,
		"animation":   &primitives.Animation,
		"zindex":      &primitives.ZIndex,
		"breakpoints": &primitives.Breakpoints,
	}

	for _, tokenMap := range categories {
		if *tokenMap == nil {
			continue
		}
		for key, value := range *tokenMap {
			ref := TokenReference(value)
			if ref.IsReference() {
				resolved, err := r.Resolve(ctx, ref, theme, tenantID, darkMode)
				if err != nil {
					if r.strictMode {
						return err
					}
					continue
				}
				(*tokenMap)[key] = resolved.Value
			}
		}
	}

	return nil
}

// resolveSemanticTokens resolves all semantic tokens.
func (r *Resolver) resolveSemanticTokens(ctx context.Context, semantic *SemanticTokens, theme *Theme, tenantID string, darkMode bool) error {
	if semantic == nil {
		return nil
	}

	categories := map[string]*map[string]string{
		"colors":      &semantic.Colors,
		"spacing":     &semantic.Spacing,
		"typography":  &semantic.Typography,
		"interactive": &semantic.Interactive,
	}

	for _, tokenMap := range categories {
		if *tokenMap == nil {
			continue
		}
		for key, value := range *tokenMap {
			ref := TokenReference(value)
			if ref.IsReference() {
				resolved, err := r.Resolve(ctx, ref, theme, tenantID, darkMode)
				if err != nil {
					if r.strictMode {
						return err
					}
					continue
				}
				(*tokenMap)[key] = resolved.Value
			}
		}
	}

	return nil
}

// resolveComponentTokens resolves all component tokens.
func (r *Resolver) resolveComponentTokens(ctx context.Context, components *ComponentTokens, theme *Theme, tenantID string, darkMode bool) error {
	if components == nil || len(*components) == 0 {
		return nil
	}

	for componentName, variants := range *components {
		for variantName, properties := range variants {
			for propertyName, value := range properties {
				if strings.Contains(value, " ") {
					resolved, err := r.resolveCompoundValue(ctx, value, theme, tenantID, darkMode)
					if err != nil {
						if r.strictMode {
							return WrapError(ErrCodeResolution,
								fmt.Sprintf("failed to resolve %s.%s.%s", componentName, variantName, propertyName), err)
						}
						continue
					}
					properties[propertyName] = resolved
				} else {
					ref := TokenReference(value)
					if ref.IsReference() {
						resolved, err := r.Resolve(ctx, ref, theme, tenantID, darkMode)
						if err != nil {
							if r.strictMode {
								return WrapError(ErrCodeResolution,
									fmt.Sprintf("failed to resolve %s.%s.%s", componentName, variantName, propertyName), err)
							}
							continue
						}
						properties[propertyName] = resolved.Value
					}
				}
			}
		}
	}

	return nil
}

// resolveCompoundValue resolves compound CSS values that may contain multiple token references.
// Example: "1rem 2rem" or "1px solid semantic.colors.border"
func (r *Resolver) resolveCompoundValue(ctx context.Context, value string, theme *Theme, tenantID string, darkMode bool) (string, error) {
	if !strings.Contains(value, " ") {
		// Fast path for single values
		ref := TokenReference(value)
		if ref.IsReference() {
			result, err := r.Resolve(ctx, ref, theme, tenantID, darkMode)
			if err != nil {
				return "", err
			}
			return result.Value, nil
		}
		return value, nil
	}

	var sb strings.Builder
	sb.Grow(len(value)) // Pre-allocate

	parts := strings.Fields(value)
	for i, part := range parts {
		if i > 0 {
			sb.WriteByte(' ')
		}

		ref := TokenReference(part)
		if ref.IsReference() {
			result, err := r.Resolve(ctx, ref, theme, tenantID, darkMode)
			if err != nil {
				return "", err
			}
			sb.WriteString(result.Value)
		} else {
			sb.WriteString(part)
		}
	}

	return sb.String(), nil
}

// mergeTokens merges two token sets with the second overriding the first.
func (r *Resolver) mergeTokens(base, override *Tokens) *Tokens {
	merged := base.Clone()

	if override.Primitives != nil {
		merged.Primitives = r.mergePrimitives(merged.Primitives, override.Primitives)
	}
	if override.Semantic != nil {
		merged.Semantic = r.mergeSemantic(merged.Semantic, override.Semantic)
	}
	if override.Components != nil {
		merged.Components = r.mergeComponents(merged.Components, override.Components)
	}

	return merged
}

// mergePrimitives merges primitive tokens.
func (r *Resolver) mergePrimitives(base, override *PrimitiveTokens) *PrimitiveTokens {
	if base == nil {
		return override.Clone()
	}
	if override == nil {
		return base
	}

	merged := base.Clone()
	if override.Colors != nil {
		for k, v := range override.Colors {
			merged.Colors[k] = v
		}
	}
	if override.Spacing != nil {
		for k, v := range override.Spacing {
			merged.Spacing[k] = v
		}
	}
	if override.Radius != nil {
		for k, v := range override.Radius {
			merged.Radius[k] = v
		}
	}
	if override.Typography != nil {
		for k, v := range override.Typography {
			merged.Typography[k] = v
		}
	}
	if override.Borders != nil {
		for k, v := range override.Borders {
			merged.Borders[k] = v
		}
	}
	if override.Shadows != nil {
		for k, v := range override.Shadows {
			merged.Shadows[k] = v
		}
	}
	if override.Effects != nil {
		for k, v := range override.Effects {
			merged.Effects[k] = v
		}
	}
	if override.Animation != nil {
		for k, v := range override.Animation {
			merged.Animation[k] = v
		}
	}
	if override.ZIndex != nil {
		for k, v := range override.ZIndex {
			merged.ZIndex[k] = v
		}
	}
	if override.Breakpoints != nil {
		for k, v := range override.Breakpoints {
			merged.Breakpoints[k] = v
		}
	}

	return merged
}

// mergeSemantic merges semantic tokens.
func (r *Resolver) mergeSemantic(base, override *SemanticTokens) *SemanticTokens {
	if base == nil {
		return override.Clone()
	}
	if override == nil {
		return base
	}

	merged := base.Clone()
	if override.Colors != nil {
		for k, v := range override.Colors {
			merged.Colors[k] = v
		}
	}
	if override.Spacing != nil {
		for k, v := range override.Spacing {
			merged.Spacing[k] = v
		}
	}
	if override.Typography != nil {
		for k, v := range override.Typography {
			merged.Typography[k] = v
		}
	}
	if override.Interactive != nil {
		for k, v := range override.Interactive {
			merged.Interactive[k] = v
		}
	}

	return merged
}

// mergeComponents merges component tokens.
func (r *Resolver) mergeComponents(base, override *ComponentTokens) *ComponentTokens {
	if base == nil {
		return override.Clone()
	}
	if override == nil {
		return base
	}

	merged := base.Clone()
	for component, variants := range *override {
		if _, exists := (*merged)[component]; !exists {
			(*merged)[component] = variants.Clone()
		} else {
			for variant, properties := range variants {
				if _, exists := (*merged)[component][variant]; !exists {
					(*merged)[component][variant] = cloneStringMap(properties)
				} else {
					for prop, value := range properties {
						(*merged)[component][variant][prop] = value
					}
				}
			}
		}
	}

	return merged
}

// buildCacheKey builds a cache key for token resolution.
func (r *Resolver) buildCacheKey(ref, themeID, tenantID string, darkMode bool) string {
	key := fmt.Sprintf("%s:%s:%s", themeID, tenantID, ref)
	if darkMode {
		key += ":dark"
	}
	return key
}

// getFromCache retrieves a resolved token from cache.
func (r *Resolver) getFromCache(key string) (*ResolvedToken, bool) {
	if r.cache == nil {
		return nil, false
	}

	value, found := r.cache.Get(key)
	if !found {
		return nil, false
	}

	result, ok := value.(*ResolvedToken)
	if !ok {
		return nil, false
	}

	resultCopy := *result
	resultCopy.FromCache = true
	return &resultCopy, true
}

// setCache stores a resolved token in cache.
func (r *Resolver) setCache(key string, result *ResolvedToken) {
	if r.cache == nil {
		return
	}

	r.cache.SetWithTTL(key, result, 1, r.cacheTTL)
}

// InvalidateCache invalidates resolver cache for a theme or tenant.
func (r *Resolver) InvalidateCache(themeID, tenantID string) {
	if r.cache == nil {
		return
	}

	r.cache.Clear()
}

// GetStats returns resolver statistics.
func (r *Resolver) GetStats() *ResolverStats {
	stats := &ResolverStats{
		MaxDepth:       r.maxDepth,
		StrictMode:     r.strictMode,
		CachingEnabled: r.cache != nil,
	}

	if r.cache != nil {
		metrics := r.cache.Metrics
		stats.CacheHits = metrics.Hits()
		stats.CacheMisses = metrics.Misses()
		stats.CacheHitRate = metrics.Ratio()
	}

	return stats
}

// ResolverStats contains resolver performance statistics.
type ResolverStats struct {
	CacheHits      uint64
	CacheMisses    uint64
	CacheHitRate   float64
	MaxDepth       int
	StrictMode     bool
	CachingEnabled bool
}

// Close closes the resolver and releases resources.
func (r *Resolver) Close() {
	if r.cache != nil {
		r.cache.Close()
	}
}

// func (r *Resolver) Resolve(ctx context.Context, ref TokenReference, theme *Theme, tenantID string, darkMode bool) (*ResolvedToken, error) {
// 	start := time.Now()
//
// 	result, err := r.resolveInternal(ctx, ref, theme, tenantID, darkMode)
//
// 	// Debug logging
// 	if r.config.DebugMode && r.config.DebugLogger != nil {
// 		r.config.DebugLogger.LogResolution(string(ref), result, time.Since(start))
// 	}
//
// 	return result, err
// }
