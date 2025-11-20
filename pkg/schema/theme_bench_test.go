package schema

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/niiniyare/ruun/pkg/logger"
)

// =============================================================================
// TOKEN RESOLUTION BENCHMARKS
// =============================================================================

// BenchmarkTokenResolver_Resolve benchmarks token resolution
func BenchmarkTokenResolver_Resolve(b *testing.B) {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	if err != nil {
		b.Fatal(err)
	}
	defer resolver.Close()

	ctx := context.Background()

	benchmarks := []struct {
		name string
		ref  TokenReference
	}{
		{"CSSLiteral", TokenReference("#3b82f6")},
		{"DirectPrimitive", TokenReference("primitives.colors.primary")},
		{"SemanticToPrimitive", TokenReference("semantic.colors.primary")},
		{"ComponentToSemanticToPrimitive", TokenReference("components.button.primary.background-color")},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err := resolver.Resolve(ctx, bm.ref)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// BenchmarkTokenResolver_ResolveWithCache benchmarks cached resolution
func BenchmarkTokenResolver_ResolveWithCache(b *testing.B) {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	if err != nil {
		b.Fatal(err)
	}
	defer resolver.Close()

	ctx := context.Background()
	ref := TokenReference("semantic.colors.primary")

	// Warm up cache
	_, _ = resolver.Resolve(ctx, ref)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := resolver.Resolve(ctx, ref)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTokenResolver_ResolveWithoutCache benchmarks resolution without cache
func BenchmarkTokenResolver_ResolveWithoutCache(b *testing.B) {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	if err != nil {
		b.Fatal(err)
	}
	defer resolver.Close()

	ctx := context.Background()
	ref := TokenReference("semantic.colors.primary")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clear cache before each resolve
		resolver.ClearCache()
		_, err := resolver.Resolve(ctx, ref)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTokenResolver_ResolveConcurrent benchmarks concurrent resolution
func BenchmarkTokenResolver_ResolveConcurrent(b *testing.B) {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	if err != nil {
		b.Fatal(err)
	}
	defer resolver.Close()

	ctx := context.Background()
	ref := TokenReference("semantic.colors.primary")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := resolver.Resolve(ctx, ref)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkTokenResolver_ResolveAll benchmarks resolving all tokens
func BenchmarkTokenResolver_ResolveAll(b *testing.B) {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	if err != nil {
		b.Fatal(err)
	}
	defer resolver.Close()

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := resolver.ResolveAll(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// =============================================================================
// TOKEN REFERENCE BENCHMARKS
// =============================================================================

// BenchmarkTokenReference_IsReference benchmarks reference detection
func BenchmarkTokenReference_IsReference(b *testing.B) {
	testCases := []struct {
		name string
		ref  TokenReference
	}{
		{"ValidReference", TokenReference("primitives.colors.primary")},
		{"CSSLiteral_Hex", TokenReference("#3b82f6")},
		{"CSSLiteral_HSL", TokenReference("hsl(217, 91%, 60%)")},
		{"CSSLiteral_Unit", TokenReference("1rem")},
		{"CSSLiteral_Calc", TokenReference("calc(100% - 2rem)")},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = tc.ref.IsReference()
			}
		})
	}
}

// BenchmarkTokenReference_Validate benchmarks validation
func BenchmarkTokenReference_Validate(b *testing.B) {
	testCases := []struct {
		name string
		ref  TokenReference
	}{
		{"ValidReference", TokenReference("primitives.colors.primary")},
		{"CSSLiteral", TokenReference("#3b82f6")},
		{"LongPath", TokenReference("components.button.primary.background-color.hover.active")},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = tc.ref.Validate()
			}
		})
	}
}

// =============================================================================
// DESIGN TOKENS BENCHMARKS
// =============================================================================

// BenchmarkDesignTokens_Validate benchmarks token validation
func BenchmarkDesignTokens_Validate(b *testing.B) {
	tokens := GetDefaultTokens()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tokens.Validate()
	}
}

// BenchmarkDesignTokens_Clone benchmarks deep cloning
func BenchmarkDesignTokens_Clone(b *testing.B) {
	tokens := GetDefaultTokens()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tokens.Clone()
	}
}

// BenchmarkDesignTokens_GetToken benchmarks token retrieval
func BenchmarkDesignTokens_GetToken(b *testing.B) {
	tokens := GetDefaultTokens()

	testCases := []struct {
		name string
		path string
	}{
		{"Primitive", "primitives.colors.primary"},
		{"Semantic", "semantic.colors.background"},
		{"Component", "components.button.primary.background-color"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = tokens.GetToken(tc.path)
			}
		})
	}
}

// =============================================================================
// THEME MANAGER BENCHMARKS
// =============================================================================

// BenchmarkThemeManager_RegisterTheme benchmarks theme registration
func BenchmarkThemeManager_RegisterTheme(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()
	baseTheme := GetDefaultTheme()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		theme := baseTheme.Clone()
		theme.ID = fmt.Sprintf("bench-theme-%d", i)
		err := manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkThemeManager_GetTheme benchmarks theme retrieval
func BenchmarkThemeManager_GetTheme(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()
	theme := GetDefaultTheme()
	theme.ID = "bench-theme"

	err = manager.RegisterTheme(ctx, theme)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := manager.GetTheme(ctx, "bench-theme", nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkThemeManager_GetThemeWithCache benchmarks cached theme retrieval
func BenchmarkThemeManager_GetThemeWithCache(b *testing.B) {
	config := DefaultThemeManagerConfig()
	config.EnableCaching = true
	manager, err := NewThemeManager(config)
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()
	theme := GetDefaultTheme()
	theme.ID = "cache-bench-theme"

	err = manager.RegisterTheme(ctx, theme)
	if err != nil {
		b.Fatal(err)
	}

	// Warm up cache
	_, _ = manager.GetTheme(ctx, "cache-bench-theme", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := manager.GetTheme(ctx, "cache-bench-theme", nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkThemeManager_GetThemeWithConditions benchmarks conditional theme retrieval
func BenchmarkThemeManager_GetThemeWithConditions(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()
	theme := GetDefaultTheme()
	theme.ID = "condition-bench-theme"
	theme.Conditions = []*ThemeCondition{
		{
			ID:         "test-condition",
			Expression: "user.role == 'admin'",
			Priority:   100,
			TokenOverrides: map[string]string{
				"semantic.colors.background": "#000000",
			},
		},
	}

	err = manager.RegisterTheme(ctx, theme)
	if err != nil {
		b.Fatal(err)
	}

	evalData := map[string]any{
		"user": map[string]any{"role": "admin"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := manager.GetTheme(ctx, "condition-bench-theme", evalData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkThemeManager_GetThemeConcurrent benchmarks concurrent theme retrieval
func BenchmarkThemeManager_GetThemeConcurrent(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()
	theme := GetDefaultTheme()
	theme.ID = "concurrent-bench-theme"

	err = manager.RegisterTheme(ctx, theme)
	if err != nil {
		b.Fatal(err)
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := manager.GetTheme(ctx, "concurrent-bench-theme", nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkThemeManager_MultiTenantGetTheme benchmarks multi-tenant theme retrieval
func BenchmarkThemeManager_MultiTenantGetTheme(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	// Register themes for 10 tenants
	for i := 0; i < 10; i++ {
		tenantID := fmt.Sprintf("tenant-%d", i)
		ctx := WithTenant(context.Background(), tenantID)
		theme := GetDefaultTheme()
		theme.ID = "tenant-theme"
		err := manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tenantID := fmt.Sprintf("tenant-%d", i%10)
		ctx := WithTenant(context.Background(), tenantID)
		_, err := manager.GetTheme(ctx, "tenant-theme", nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkThemeManager_ListThemes benchmarks listing themes
func BenchmarkThemeManager_ListThemes(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	ctx := context.Background()

	// Register 100 themes
	for i := 0; i < 100; i++ {
		theme := GetDefaultTheme()
		theme.ID = fmt.Sprintf("list-bench-theme-%d", i)
		err := manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := manager.ListThemes(ctx)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// =============================================================================
// THEME OPERATIONS BENCHMARKS
// =============================================================================

// BenchmarkTheme_Validate benchmarks theme validation
func BenchmarkTheme_Validate(b *testing.B) {
	theme := GetDefaultTheme()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = theme.Validate()
	}
}

// BenchmarkTheme_Clone benchmarks theme cloning
func BenchmarkTheme_Clone(b *testing.B) {
	theme := GetDefaultTheme()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = theme.Clone()
	}
}

// BenchmarkTheme_ToJSON benchmarks JSON serialization
func BenchmarkTheme_ToJSON(b *testing.B) {
	theme := GetDefaultTheme()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := theme.ToJSON()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTheme_FromJSON benchmarks JSON deserialization
func BenchmarkTheme_FromJSON(b *testing.B) {
	theme := GetDefaultTheme()
	jsonData, err := theme.ToJSON()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ThemeFromJSON(jsonData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTheme_ToYAML benchmarks YAML serialization
func BenchmarkTheme_ToYAML(b *testing.B) {
	theme := GetDefaultTheme()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := theme.ToYAML()
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkTheme_FromYAML benchmarks YAML deserialization
func BenchmarkTheme_FromYAML(b *testing.B) {
	theme := GetDefaultTheme()
	yamlData, err := theme.ToYAML()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ThemeFromYAML(yamlData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// =============================================================================
// CONDITION EVALUATION BENCHMARKS
// =============================================================================

// BenchmarkThemeCondition_Validate benchmarks condition validation
func BenchmarkThemeCondition_Validate(b *testing.B) {
	condition := &ThemeCondition{
		ID:         "bench-condition",
		Expression: "user.role == 'admin'",
		Priority:   100,
		TokenOverrides: map[string]string{
			"semantic.colors.background": "#000000",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = condition.Validate()
	}
}

// BenchmarkConditionEvaluator_Evaluate benchmarks expression evaluation
func BenchmarkConditionEvaluator_Evaluate(b *testing.B) {
	evaluator, err := newConditionEvaluator()
	if err != nil {
		b.Fatal(err)
	}
	defer evaluator.close()

	ctx := context.Background()
	expr := "user.role == 'admin'"
	evalData := map[string]any{
		"user": map[string]any{"role": "admin"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := evaluator.evaluate(ctx, expr, evalData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// =============================================================================
// MEMORY ALLOCATION BENCHMARKS
// =============================================================================

// BenchmarkDefaultTokens_Allocation benchmarks default token creation
func BenchmarkDefaultTokens_Allocation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetDefaultTokens()
	}
}

// BenchmarkDefaultTheme_Allocation benchmarks default theme creation
func BenchmarkDefaultTheme_Allocation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GetDefaultTheme()
	}
}

// BenchmarkThemeManager_Allocation benchmarks theme manager creation
func BenchmarkThemeManager_Allocation(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		manager, err := NewThemeManager(DefaultThemeManagerConfig())
		if err != nil {
			b.Fatal(err)
		}
		manager.Close()
	}
}

// =============================================================================
// COMPREHENSIVE WORKFLOW BENCHMARKS
// =============================================================================

// BenchmarkCompleteWorkflow benchmarks a complete theme workflow
func BenchmarkCompleteWorkflow(b *testing.B) {
	config := DefaultThemeManagerConfig()
	config.Logger = logger.WithFields(logger.Fields{"benchmark": true})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create manager
		manager, err := NewThemeManager(config)
		if err != nil {
			b.Fatal(err)
		}

		// Register theme
		theme := GetDefaultTheme()
		theme.ID = fmt.Sprintf("workflow-theme-%d", i)
		ctx := context.Background()
		err = manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}

		// Retrieve theme
		_, err = manager.GetTheme(ctx, theme.ID, nil)
		if err != nil {
			b.Fatal(err)
		}

		// Clean up
		manager.Close()
	}
}

// BenchmarkMultiTenantWorkflow benchmarks multi-tenant workflow
func BenchmarkMultiTenantWorkflow(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tenantID := fmt.Sprintf("tenant-%d", i%10)
		ctx := WithTenant(context.Background(), tenantID)

		theme := GetDefaultTheme()
		theme.ID = fmt.Sprintf("workflow-theme-%d", i)

		// Register
		err := manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}

		// Retrieve
		_, err = manager.GetTheme(ctx, theme.ID, nil)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkConcurrentWorkflow benchmarks concurrent operations
func BenchmarkConcurrentWorkflow(b *testing.B) {
	manager, err := NewThemeManager(DefaultThemeManagerConfig())
	if err != nil {
		b.Fatal(err)
	}
	defer manager.Close()

	// Pre-register some themes
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		theme := GetDefaultTheme()
		theme.ID = fmt.Sprintf("concurrent-theme-%d", i)
		err := manager.RegisterTheme(ctx, theme)
		if err != nil {
			b.Fatal(err)
		}
	}

	var counter int
	var mu sync.Mutex

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			mu.Lock()
			i := counter
			counter++
			mu.Unlock()

			themeID := fmt.Sprintf("concurrent-theme-%d", i%10)
			_, err := manager.GetTheme(ctx, themeID, nil)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
