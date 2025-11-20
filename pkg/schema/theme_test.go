package schema

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/niiniyare/ruun/pkg/logger"
	"github.com/stretchr/testify/suite"
)

// =============================================================================
// THEME BASICS TEST SUITE
// =============================================================================

// ThemeBasicsTestSuite tests basic theme structure and operations
type ThemeBasicsTestSuite struct {
	suite.Suite
	theme *Theme
}

func TestThemeBasicsTestSuite(t *testing.T) {
	suite.Run(t, new(ThemeBasicsTestSuite))
}

func (s *ThemeBasicsTestSuite) SetupTest() {
	s.theme = &Theme{
		ID:          "test-theme",
		Name:        "Test Theme",
		Description: "A theme for testing",
		Version:     "1.0.0",
		Author:      "Test Author",
		Tokens:      GetDefaultTokens(),
	}
}

func (s *ThemeBasicsTestSuite) TestValidate_ValidTheme() {
	err := s.theme.Validate()
	s.NoError(err)
}

func (s *ThemeBasicsTestSuite) TestValidate_MissingID() {
	s.theme.ID = ""
	err := s.theme.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_theme_id")
}

func (s *ThemeBasicsTestSuite) TestValidate_MissingName() {
	s.theme.Name = ""
	err := s.theme.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_theme_name")
}

func (s *ThemeBasicsTestSuite) TestValidate_NilTokens() {
	s.theme.Tokens = nil
	err := s.theme.Validate()
	s.Error(err)
	s.Contains(err.Error(), "nil_tokens")
}

func (s *ThemeBasicsTestSuite) TestValidate_InvalidTokens() {
	s.theme.Tokens = &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "semantic.colors.primary", // Invalid
			},
		},
	}
	err := s.theme.Validate()
	s.Error(err)
}

func (s *ThemeBasicsTestSuite) TestValidate_CustomCSSTooBig() {
	s.theme.CustomCSS = string(make([]byte, 2<<20)) // 2MB
	err := s.theme.Validate()
	s.Error(err)
	s.Contains(err.Error(), "custom_css_too_large")
}

func (s *ThemeBasicsTestSuite) TestValidate_CustomJSTooBig() {
	s.theme.CustomJS = string(make([]byte, 2<<20)) // 2MB
	err := s.theme.Validate()
	s.Error(err)
	s.Contains(err.Error(), "custom_js_too_large")
}

func (s *ThemeBasicsTestSuite) TestClone() {
	cloned := s.theme.Clone()
	s.NotNil(cloned)
	s.NotSame(s.theme, cloned)
	s.Equal(s.theme.ID, cloned.ID)
	s.NotSame(s.theme.Tokens, cloned.Tokens)

	// Modify cloned - should not affect original
	cloned.Name = "Modified Name"
	s.Equal("Test Theme", s.theme.Name)
	s.Equal("Modified Name", cloned.Name)
}

func (s *ThemeBasicsTestSuite) TestClone_NilTheme() {
	var theme *Theme
	cloned := theme.Clone()
	s.Nil(cloned)
}

func (s *ThemeBasicsTestSuite) TestToJSON() {
	jsonData, err := s.theme.ToJSON()
	s.NoError(err)
	s.NotEmpty(jsonData)
	s.Contains(string(jsonData), `"test-theme"`)
	s.Contains(string(jsonData), `"Test Theme"`)
}

func (s *ThemeBasicsTestSuite) TestFromJSON() {
	jsonData, err := s.theme.ToJSON()
	s.Require().NoError(err)

	parsed, err := ThemeFromJSON(jsonData)
	s.NoError(err)
	s.NotNil(parsed)
	s.Equal(s.theme.ID, parsed.ID)
	s.Equal(s.theme.Name, parsed.Name)
}

func (s *ThemeBasicsTestSuite) TestFromJSON_Invalid() {
	invalidJSON := []byte(`{"invalid": "json"`)
	_, err := ThemeFromJSON(invalidJSON)
	s.Error(err)
}

func (s *ThemeBasicsTestSuite) TestToYAML() {
	yamlData, err := s.theme.ToYAML()
	s.NoError(err)
	s.NotEmpty(yamlData)
	s.Contains(string(yamlData), "test-theme")
	s.Contains(string(yamlData), "Test Theme")
}

func (s *ThemeBasicsTestSuite) TestFromYAML() {
	yamlData, err := s.theme.ToYAML()
	s.Require().NoError(err)

	parsed, err := ThemeFromYAML(yamlData)
	s.NoError(err)
	s.NotNil(parsed)
	s.Equal(s.theme.ID, parsed.ID)
}

// =============================================================================
// DARK MODE TEST SUITE
// =============================================================================

// DarkModeTestSuite tests dark mode configuration
type DarkModeTestSuite struct {
	suite.Suite
}

func TestDarkModeTestSuite(t *testing.T) {
	suite.Run(t, new(DarkModeTestSuite))
}

func (s *DarkModeTestSuite) TestValidate_ValidConfig() {
	config := &DarkModeConfig{
		Enabled:  true,
		Default:  false,
		Strategy: "auto",
		DarkTokens: &DesignTokens{
			Semantic: &SemanticTokens{
				Colors: map[string]string{
					"background": "primitives.colors.gray-900",
				},
			},
		},
	}

	err := config.Validate()
	s.NoError(err)
}

func (s *DarkModeTestSuite) TestValidate_InvalidStrategy() {
	config := &DarkModeConfig{
		Enabled:  true,
		Strategy: "invalid",
	}

	err := config.Validate()
	s.Error(err)
	s.Contains(err.Error(), "invalid_dark_mode_strategy")
}

func (s *DarkModeTestSuite) TestValidate_ValidStrategies() {
	strategies := []string{"class", "media", "auto"}
	for _, strategy := range strategies {
		s.Run(strategy, func() {
			config := &DarkModeConfig{
				Enabled:  true,
				Strategy: strategy,
			}
			err := config.Validate()
			s.NoError(err)
		})
	}
}

func (s *DarkModeTestSuite) TestClone() {
	original := &DarkModeConfig{
		Enabled:  true,
		Default:  true,
		Strategy: "auto",
		DarkTokens: &DesignTokens{
			Primitives: &PrimitiveTokens{
				Colors: map[string]string{
					"background": "#000000",
				},
			},
		},
	}

	cloned := original.Clone()
	s.NotNil(cloned)
	s.NotSame(original, cloned)
	s.NotSame(original.DarkTokens, cloned.DarkTokens)

	// Modify cloned
	cloned.Enabled = false
	s.True(original.Enabled)
	s.False(cloned.Enabled)
}

// =============================================================================
// ACCESSIBILITY TEST SUITE
// =============================================================================

// AccessibilityTestSuite tests accessibility configuration
type AccessibilityTestSuite struct {
	suite.Suite
}

func TestAccessibilityTestSuite(t *testing.T) {
	suite.Run(t, new(AccessibilityTestSuite))
}

func (s *AccessibilityTestSuite) TestValidate_ValidConfig() {
	config := &AccessibilityConfig{
		HighContrast:          true,
		MinContrastRatio:      4.5,
		FocusIndicator:        true,
		FocusOutlineColor:     "#0066cc",
		FocusOutlineWidth:     "2px",
		KeyboardNav:           true,
		ReducedMotion:         false,
		ScreenReaderOptimized: true,
		AriaLive:              "polite",
	}

	err := config.Validate()
	s.NoError(err)
}

func (s *AccessibilityTestSuite) TestValidate_InvalidContrastRatio() {
	testCases := []float64{-1, 22, 100}
	for _, ratio := range testCases {
		s.Run(fmt.Sprintf("ratio_%.0f", ratio), func() {
			config := &AccessibilityConfig{
				MinContrastRatio: ratio,
			}
			err := config.Validate()
			s.Error(err)
			s.Contains(err.Error(), "invalid_contrast_ratio")
		})
	}
}

func (s *AccessibilityTestSuite) TestValidate_InvalidAriaLive() {
	config := &AccessibilityConfig{
		AriaLive: "invalid",
	}
	err := config.Validate()
	s.Error(err)
	s.Contains(err.Error(), "invalid_aria_live")
}

func (s *AccessibilityTestSuite) TestValidate_ValidAriaLive() {
	validValues := []string{"off", "polite", "assertive"}
	for _, value := range validValues {
		s.Run(value, func() {
			config := &AccessibilityConfig{
				AriaLive: value,
			}
			err := config.Validate()
			s.NoError(err)
		})
	}
}

// =============================================================================
// THEME CONDITIONS TEST SUITE
// =============================================================================

// ThemeConditionsTestSuite tests conditional theme overrides
type ThemeConditionsTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestThemeConditionsTestSuite(t *testing.T) {
	suite.Run(t, new(ThemeConditionsTestSuite))
}

func (s *ThemeConditionsTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *ThemeConditionsTestSuite) TestThemeCondition_Validate() {
	condition := &ThemeCondition{
		ID:         "test-condition",
		Expression: "user.role == 'admin'",
		Priority:   100,
		TokenOverrides: map[string]string{
			"semantic.colors.background": "primitives.colors.gray-900",
		},
		Description: "Admin users get dark background",
	}

	err := condition.Validate()
	s.NoError(err)
}

func (s *ThemeConditionsTestSuite) TestThemeCondition_Validate_MissingID() {
	condition := &ThemeCondition{
		Expression: "user.role == 'admin'",
		TokenOverrides: map[string]string{
			"semantic.colors.background": "primitives.colors.gray-900",
		},
	}

	err := condition.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_condition_id")
}

func (s *ThemeConditionsTestSuite) TestThemeCondition_Validate_MissingExpression() {
	condition := &ThemeCondition{
		ID: "test",
		TokenOverrides: map[string]string{
			"semantic.colors.background": "primitives.colors.gray-900",
		},
	}

	err := condition.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_expression")
}

func (s *ThemeConditionsTestSuite) TestThemeCondition_Validate_EmptyOverrides() {
	condition := &ThemeCondition{
		ID:             "test",
		Expression:     "user.role == 'admin'",
		TokenOverrides: map[string]string{},
	}

	err := condition.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_overrides")
}

func (s *ThemeConditionsTestSuite) TestThemeCondition_Clone() {
	original := &ThemeCondition{
		ID:         "test",
		Expression: "user.role == 'admin'",
		Priority:   100,
		TokenOverrides: map[string]string{
			"semantic.colors.background": "primitives.colors.gray-900",
		},
	}

	cloned := original.Clone()
	s.NotNil(cloned)
	s.NotSame(original, cloned)

	// Modify cloned - should not affect original
	cloned.Priority = 200
	cloned.TokenOverrides["new.token"] = "new.value"
	s.Equal(100, original.Priority)
	s.Equal(200, cloned.Priority)
	s.NotContains(original.TokenOverrides, "new.token")
	s.Contains(cloned.TokenOverrides, "new.token")
}

// =============================================================================
// THEME MANAGER TEST SUITE
// =============================================================================

// ThemeManagerTestSuite tests theme manager operations
type ThemeManagerTestSuite struct {
	suite.Suite
	manager *ThemeManager
	ctx     context.Context
}

func TestThemeManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ThemeManagerTestSuite))
}

func (s *ThemeManagerTestSuite) SetupTest() {
	config := DefaultThemeManagerConfig()
	config.ValidateOnRegister = true
	config.EnableCaching = true
	config.Logger = logger.WithFields(logger.Fields{"test": "theme_manager"})

	var err error
	s.manager, err = NewThemeManager(config)
	s.Require().NoError(err)

	s.ctx = context.Background()
}

func (s *ThemeManagerTestSuite) TearDownTest() {
	if s.manager != nil {
		s.manager.Close()
	}
}

func (s *ThemeManagerTestSuite) TestNewThemeManager_NilConfig() {
	manager, err := NewThemeManager(nil)
	s.NoError(err)
	s.NotNil(manager)
	defer manager.Close()
}

func (s *ThemeManagerTestSuite) TestRegisterTheme_Global() {
	theme := &Theme{
		ID:     "global-theme",
		Name:   "Global Theme",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.NoError(err)

	// Should be able to retrieve it
	retrieved, err := s.manager.GetTheme(s.ctx, "global-theme", nil)
	s.NoError(err)
	s.NotNil(retrieved)
	s.Equal("global-theme", retrieved.ID)
}

func (s *ThemeManagerTestSuite) TestRegisterTheme_TenantSpecific() {
	tenantCtx := WithTenant(s.ctx, "tenant-123")

	theme := &Theme{
		ID:     "tenant-theme",
		Name:   "Tenant Theme",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(tenantCtx, theme)
	s.NoError(err)

	// Should be retrievable with tenant context
	retrieved, err := s.manager.GetTheme(tenantCtx, "tenant-theme", nil)
	s.NoError(err)
	s.NotNil(retrieved)
	s.Equal("tenant-theme", retrieved.ID)

	// Should NOT be retrievable without tenant context
	_, err = s.manager.GetTheme(s.ctx, "tenant-theme", nil)
	s.Error(err)
	s.Contains(err.Error(), "theme_not_found")
}

func (s *ThemeManagerTestSuite) TestRegisterTheme_NilTheme() {
	err := s.manager.RegisterTheme(s.ctx, nil)
	s.Error(err)
	s.Contains(err.Error(), "nil_theme")
}

func (s *ThemeManagerTestSuite) TestRegisterTheme_InvalidTheme() {
	theme := &Theme{
		ID:     "invalid",
		Name:   "Invalid Theme",
		Tokens: nil, // Invalid: nil tokens
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Error(err)
}

func (s *ThemeManagerTestSuite) TestRegisterTheme_SetsTimestamps() {
	theme := &Theme{
		ID:     "timestamp-test",
		Name:   "Timestamp Test",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	retrieved, err := s.manager.GetTheme(s.ctx, "timestamp-test", nil)
	s.Require().NoError(err)

	s.False(retrieved.createdAt.IsZero())
	s.False(retrieved.updatedAt.IsZero())
}

func (s *ThemeManagerTestSuite) TestGetTheme_NotFound() {
	_, err := s.manager.GetTheme(s.ctx, "nonexistent", nil)
	s.Error(err)
	s.Contains(err.Error(), "theme_not_found")
}

func (s *ThemeManagerTestSuite) TestGetTheme_DefaultFallback() {
	// Register default theme
	defaultTheme := &Theme{
		ID:     "default",
		Name:   "Default Theme",
		Tokens: GetDefaultTokens(),
	}
	err := s.manager.RegisterTheme(s.ctx, defaultTheme)
	s.Require().NoError(err)

	// Configure default
	s.manager.config.DefaultThemeID = "default"

	// Request nonexistent theme - should fallback to default
	retrieved, err := s.manager.GetTheme(s.ctx, "nonexistent", nil)
	s.NoError(err)
	s.NotNil(retrieved)
	s.Equal("default", retrieved.ID)
}

func (s *ThemeManagerTestSuite) TestGetTheme_WithConditionalOverrides() {
	// Register theme with conditions
	theme := &Theme{
		ID:     "conditional-theme",
		Name:   "Conditional Theme",
		Tokens: GetDefaultTokens(),
		Conditions: []*ThemeCondition{
			{
				ID:         "admin-dark",
				Expression: "user.role == 'admin'",
				Priority:   100,
				TokenOverrides: map[string]string{
					"semantic.colors.background": "#000000",
				},
			},
		},
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	// Get theme with evaluation data matching condition
	evalData := map[string]any{
		"user": map[string]any{
			"role": "admin",
		},
	}

	retrieved, err := s.manager.GetTheme(s.ctx, "conditional-theme", evalData)
	s.NoError(err)
	s.NotNil(retrieved)

	// Background color should be overridden
	s.Equal("#000000", retrieved.Tokens.Semantic.Colors["background"])
}

func (s *ThemeManagerTestSuite) TestGetTheme_ConditionPriority() {
	// Register theme with multiple conditions
	theme := &Theme{
		ID:     "priority-test",
		Name:   "Priority Test",
		Tokens: GetDefaultTokens(),
		Conditions: []*ThemeCondition{
			{
				ID:         "low-priority",
				Expression: "true", // Always true
				Priority:   10,
				TokenOverrides: map[string]string{
					"semantic.colors.background": "#111111",
				},
			},
			{
				ID:         "high-priority",
				Expression: "true", // Always true
				Priority:   100,
				TokenOverrides: map[string]string{
					"semantic.colors.background": "#222222",
				},
			},
		},
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	retrieved, err := s.manager.GetTheme(s.ctx, "priority-test", map[string]any{})
	s.NoError(err)

	// High priority should win
	s.Equal("#222222", retrieved.Tokens.Semantic.Colors["background"])
}

func (s *ThemeManagerTestSuite) TestListThemes_Global() {
	// Register multiple global themes
	for i := range 3 {
		theme := &Theme{
			ID:     fmt.Sprintf("theme-%d", i),
			Name:   fmt.Sprintf("Theme %d", i),
			Tokens: GetDefaultTokens(),
		}
		err := s.manager.RegisterTheme(s.ctx, theme)
		s.Require().NoError(err)
	}

	themes, err := s.manager.ListThemes(s.ctx)
	s.NoError(err)
	s.Len(themes, 3)
}

func (s *ThemeManagerTestSuite) TestListThemes_Tenant() {
	tenantCtx := WithTenant(s.ctx, "tenant-123")

	// Register tenant themes
	for i := range 2 {
		theme := &Theme{
			ID:     fmt.Sprintf("tenant-theme-%d", i),
			Name:   fmt.Sprintf("Tenant Theme %d", i),
			Tokens: GetDefaultTokens(),
		}
		err := s.manager.RegisterTheme(tenantCtx, theme)
		s.Require().NoError(err)
	}

	// Register global themes
	globalTheme := &Theme{
		ID:     "global-theme",
		Name:   "Global Theme",
		Tokens: GetDefaultTokens(),
	}
	err := s.manager.RegisterTheme(s.ctx, globalTheme)
	s.Require().NoError(err)

	// List with tenant context - should include both tenant and global
	themes, err := s.manager.ListThemes(tenantCtx)
	s.NoError(err)
	s.GreaterOrEqual(len(themes), 3)
}

func (s *ThemeManagerTestSuite) TestUnregisterTheme() {
	theme := &Theme{
		ID:     "unregister-test",
		Name:   "Unregister Test",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	// Unregister
	err = s.manager.UnregisterTheme(s.ctx, "unregister-test")
	s.NoError(err)

	// Should not be retrievable
	_, err = s.manager.GetTheme(s.ctx, "unregister-test", nil)
	s.Error(err)
	s.Contains(err.Error(), "theme_not_found")
}

func (s *ThemeManagerTestSuite) TestUnregisterTheme_NotFound() {
	err := s.manager.UnregisterTheme(s.ctx, "nonexistent")
	s.Error(err)
	s.Contains(err.Error(), "theme_not_found")
}

// =============================================================================
// THEME CACHING TEST SUITE
// =============================================================================

// ThemeCachingTestSuite tests theme caching behavior
type ThemeCachingTestSuite struct {
	suite.Suite
	manager *ThemeManager
	ctx     context.Context
}

func TestThemeCachingTestSuite(t *testing.T) {
	suite.Run(t, new(ThemeCachingTestSuite))
}

func (s *ThemeCachingTestSuite) SetupTest() {
	config := DefaultThemeManagerConfig()
	config.EnableCaching = true

	var err error
	s.manager, err = NewThemeManager(config)
	s.Require().NoError(err)

	s.ctx = context.Background()
}

func (s *ThemeCachingTestSuite) TearDownTest() {
	if s.manager != nil {
		s.manager.Close()
	}
}

func (s *ThemeCachingTestSuite) TestCacheHit() {
	theme := &Theme{
		ID:     "cache-test",
		Name:   "Cache Test",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	// First retrieval - cache miss
	start1 := time.Now()
	theme1, err := s.manager.GetTheme(s.ctx, "cache-test", nil)
	duration1 := time.Since(start1)
	s.NoError(err)
	s.NotNil(theme1)

	// Second retrieval - cache hit (should be faster)
	start2 := time.Now()
	theme2, err := s.manager.GetTheme(s.ctx, "cache-test", nil)
	duration2 := time.Since(start2)
	s.NoError(err)
	s.NotNil(theme2)

	// Second should generally be faster due to cache
	// (Though not guaranteed on all systems, so we just verify it works)
	s.Less(duration2, duration1*2) // At least not slower
}

func (s *ThemeCachingTestSuite) TestCacheInvalidation() {
	theme := &Theme{
		ID:     "invalidation-test",
		Name:   "Invalidation Test",
		Tokens: GetDefaultTokens(),
	}

	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	// Retrieve to populate cache
	_, err = s.manager.GetTheme(s.ctx, "invalidation-test", nil)
	s.NoError(err)

	// Re-register (should invalidate cache)
	theme.Name = "Updated Name"
	err = s.manager.RegisterTheme(s.ctx, theme)
	s.NoError(err)

	// Should get updated theme
	retrieved, err := s.manager.GetTheme(s.ctx, "invalidation-test", nil)
	s.NoError(err)
	s.Equal("Updated Name", retrieved.Name)
}

func (s *ThemeCachingTestSuite) TestCacheWithDifferentTenants() {
	theme := &Theme{
		ID:     "multi-tenant-cache",
		Name:   "Multi-Tenant Cache",
		Tokens: GetDefaultTokens(),
	}

	// Register for tenant1
	tenant1Ctx := WithTenant(s.ctx, "tenant-1")
	err := s.manager.RegisterTheme(tenant1Ctx, theme)
	s.Require().NoError(err)

	// Register for tenant2 with different name
	theme2 := theme.Clone()
	theme2.Name = "Tenant 2 Theme"
	tenant2Ctx := WithTenant(s.ctx, "tenant-2")
	err = s.manager.RegisterTheme(tenant2Ctx, theme2)
	s.Require().NoError(err)

	// Retrieve from tenant1
	retrieved1, err := s.manager.GetTheme(tenant1Ctx, "multi-tenant-cache", nil)
	s.NoError(err)
	s.Equal("Multi-Tenant Cache", retrieved1.Name)

	// Retrieve from tenant2
	retrieved2, err := s.manager.GetTheme(tenant2Ctx, "multi-tenant-cache", nil)
	s.NoError(err)
	s.Equal("Tenant 2 Theme", retrieved2.Name)
}

// =============================================================================
// THEME CONCURRENCY TEST SUITE
// =============================================================================

// ThemeConcurrencyTestSuite tests concurrent theme operations
type ThemeConcurrencyTestSuite struct {
	suite.Suite
	manager *ThemeManager
	ctx     context.Context
}

func TestThemeConcurrencyTestSuite(t *testing.T) {
	suite.Run(t, new(ThemeConcurrencyTestSuite))
}

func (s *ThemeConcurrencyTestSuite) SetupTest() {
	var err error
	s.manager, err = NewThemeManager(DefaultThemeManagerConfig())
	s.Require().NoError(err)

	s.ctx = context.Background()
}

func (s *ThemeConcurrencyTestSuite) TearDownTest() {
	if s.manager != nil {
		s.manager.Close()
	}
}

func (s *ThemeConcurrencyTestSuite) TestConcurrentRegistration() {
	const numGoroutines = 50
	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)

	for i := range numGoroutines {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			theme := &Theme{
				ID:     fmt.Sprintf("concurrent-theme-%d", id),
				Name:   fmt.Sprintf("Concurrent Theme %d", id),
				Tokens: GetDefaultTokens(),
			}
			if err := s.manager.RegisterTheme(s.ctx, theme); err != nil {
				errors <- err
			}
		}(i)
	}

	wg.Wait()
	close(errors)

	// Check for errors
	for err := range errors {
		s.Fail("Concurrent registration error: %v", err)
	}

	// Verify all themes registered
	themes, err := s.manager.ListThemes(s.ctx)
	s.NoError(err)
	s.GreaterOrEqual(len(themes), numGoroutines)
}

func (s *ThemeConcurrencyTestSuite) TestConcurrentRetrieval() {
	// Register a theme first
	theme := &Theme{
		ID:     "concurrent-get",
		Name:   "Concurrent Get",
		Tokens: GetDefaultTokens(),
	}
	err := s.manager.RegisterTheme(s.ctx, theme)
	s.Require().NoError(err)

	const numGoroutines = 100
	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)

	for range numGoroutines {
		wg.Go(func() {
			retrieved, err := s.manager.GetTheme(s.ctx, "concurrent-get", nil)
			if err != nil {
				errors <- err
				return
			}
			if retrieved.ID != "concurrent-get" {
				errors <- fmt.Errorf("wrong theme retrieved: %s", retrieved.ID)
			}
		})
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		close(errors)
		for err := range errors {
			s.Fail("Concurrent retrieval error: %v", err)
		}
	case <-time.After(5 * time.Second):
		s.Fail("Concurrent retrieval timed out")
	}
}

func (s *ThemeConcurrencyTestSuite) TestConcurrentMultiTenantOperations() {
	const numTenants = 10
	const numThemesPerTenant = 5
	var wg sync.WaitGroup
	errors := make(chan error, numTenants*numThemesPerTenant)

	for tenant := range numTenants {
		for themeNum := range numThemesPerTenant {
			wg.Add(1)
			go func(tenantID int, themeID int) {
				defer wg.Done()

				tenantCtx := WithTenant(s.ctx, fmt.Sprintf("tenant-%d", tenantID))
				theme := &Theme{
					ID:     fmt.Sprintf("theme-%d-%d", tenantID, themeID),
					Name:   fmt.Sprintf("Theme %d-%d", tenantID, themeID),
					Tokens: GetDefaultTokens(),
				}

				if err := s.manager.RegisterTheme(tenantCtx, theme); err != nil {
					errors <- err
					return
				}

				// Immediately try to retrieve
				retrieved, err := s.manager.GetTheme(tenantCtx, theme.ID, nil)
				if err != nil {
					errors <- err
					return
				}
				if retrieved.ID != theme.ID {
					errors <- fmt.Errorf("retrieved wrong theme: expected %s, got %s", theme.ID, retrieved.ID)
				}
			}(tenant, themeNum)
		}
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		close(errors)
		for err := range errors {
			s.Fail("Multi-tenant concurrent error: %v", err)
		}
	case <-time.After(10 * time.Second):
		s.Fail("Multi-tenant concurrent operations timed out")
	}
}

// =============================================================================
// MULTI-TENANCY TEST SUITE
// =============================================================================

// MultiTenancyTestSuite tests multi-tenant functionality
type MultiTenancyTestSuite struct {
	suite.Suite
	manager *ThemeManager
	ctx     context.Context
}

func TestMultiTenancyTestSuite(t *testing.T) {
	suite.Run(t, new(MultiTenancyTestSuite))
}

func (s *MultiTenancyTestSuite) SetupTest() {
	var err error
	s.manager, err = NewThemeManager(DefaultThemeManagerConfig())
	s.Require().NoError(err)

	s.ctx = context.Background()
}

func (s *MultiTenancyTestSuite) TearDownTest() {
	if s.manager != nil {
		s.manager.Close()
	}
}

func (s *MultiTenancyTestSuite) TestWithTenant() {
	tenantCtx := WithTenant(s.ctx, "test-tenant")
	tenantID := GetTenant(tenantCtx)
	s.Equal("test-tenant", tenantID)
}

func (s *MultiTenancyTestSuite) TestGetTenant_NoTenant() {
	tenantID := GetTenant(s.ctx)
	s.Empty(tenantID)
}

func (s *MultiTenancyTestSuite) TestTenantIsolation() {
	// Register theme for tenant1
	tenant1Ctx := WithTenant(s.ctx, "tenant-1")
	theme1 := &Theme{
		ID:     "isolated-theme",
		Name:   "Tenant 1 Theme",
		Tokens: GetDefaultTokens(),
	}
	err := s.manager.RegisterTheme(tenant1Ctx, theme1)
	s.Require().NoError(err)

	// Register different theme with same ID for tenant2
	tenant2Ctx := WithTenant(s.ctx, "tenant-2")
	theme2 := &Theme{
		ID:     "isolated-theme",
		Name:   "Tenant 2 Theme",
		Tokens: GetDefaultTokens(),
	}
	err = s.manager.RegisterTheme(tenant2Ctx, theme2)
	s.Require().NoError(err)

	// Retrieve from tenant1
	retrieved1, err := s.manager.GetTheme(tenant1Ctx, "isolated-theme", nil)
	s.NoError(err)
	s.Equal("Tenant 1 Theme", retrieved1.Name)

	// Retrieve from tenant2
	retrieved2, err := s.manager.GetTheme(tenant2Ctx, "isolated-theme", nil)
	s.NoError(err)
	s.Equal("Tenant 2 Theme", retrieved2.Name)
}

func (s *MultiTenancyTestSuite) TestTenantCanAccessGlobalThemes() {
	// Register global theme
	globalTheme := &Theme{
		ID:     "global-accessible",
		Name:   "Global Theme",
		Tokens: GetDefaultTokens(),
	}
	err := s.manager.RegisterTheme(s.ctx, globalTheme)
	s.Require().NoError(err)

	// Tenant should be able to access global theme
	tenantCtx := WithTenant(s.ctx, "tenant-123")
	retrieved, err := s.manager.GetTheme(tenantCtx, "global-accessible", nil)
	s.NoError(err)
	s.Equal("Global Theme", retrieved.Name)
}

// =============================================================================
// DEFAULT THEME TEST SUITE
// =============================================================================

// DefaultThemeTestSuite tests the default theme
type DefaultThemeTestSuite struct {
	suite.Suite
	theme *Theme
}

func TestDefaultThemeTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultThemeTestSuite))
}

func (s *DefaultThemeTestSuite) SetupTest() {
	s.theme = GetDefaultTheme()
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_Valid() {
	err := s.theme.Validate()
	s.NoError(err)
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_HasRequiredFields() {
	s.NotEmpty(s.theme.ID)
	s.NotEmpty(s.theme.Name)
	s.NotNil(s.theme.Tokens)
	s.NotNil(s.theme.DarkMode)
	s.NotNil(s.theme.Accessibility)
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_DarkModeEnabled() {
	s.True(s.theme.DarkMode.Enabled)
	s.NotNil(s.theme.DarkMode.DarkTokens)
	s.NotEmpty(s.theme.DarkMode.DarkTokens.Semantic.Colors)
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_AccessibilityConfigured() {
	s.True(s.theme.Accessibility.FocusIndicator)
	s.True(s.theme.Accessibility.KeyboardNav)
	s.Greater(s.theme.Accessibility.MinContrastRatio, 0.0)
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_HasConditions() {
	s.NotEmpty(s.theme.Conditions)
	for _, cond := range s.theme.Conditions {
		err := cond.Validate()
		s.NoError(err)
	}
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_Serializable() {
	// Should be able to serialize to JSON
	jsonData, err := s.theme.ToJSON()
	s.NoError(err)
	s.NotEmpty(jsonData)

	// Should be able to deserialize
	parsed, err := ThemeFromJSON(jsonData)
	s.NoError(err)
	s.NotNil(parsed)
	s.Equal(s.theme.ID, parsed.ID)
}

func (s *DefaultThemeTestSuite) TestDefaultTheme_Serializable_YAML() {
	// Should be able to serialize to YAML
	yamlData, err := s.theme.ToYAML()
	s.NoError(err)
	s.NotEmpty(yamlData)

	// Should be able to deserialize
	parsed, err := ThemeFromYAML(yamlData)
	s.NoError(err)
	s.NotNil(parsed)
	s.Equal(s.theme.ID, parsed.ID)
}

// =============================================================================
// CONTEXT KEY TEST SUITE
// =============================================================================

// ContextKeyTestSuite tests context key functionality
type ContextKeyTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestContextKeyTestSuite(t *testing.T) {
	suite.Run(t, new(ContextKeyTestSuite))
}

func (s *ContextKeyTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *ContextKeyTestSuite) TestWithTenant_String() {
	ctx := WithTenant(s.ctx, "tenant-123")
	tenantID := GetTenant(ctx)
	s.Equal("tenant-123", tenantID)
}

func (s *ContextKeyTestSuite) TestWithTenant_UUID() {
	uuid := "550e8400-e29b-41d4-a716-446655440000"
	ctx := WithTenant(s.ctx, uuid)
	tenantID := GetTenant(ctx)
	s.Equal(uuid, tenantID)
}

func (s *ContextKeyTestSuite) TestWithTenant_Empty() {
	ctx := WithTenant(s.ctx, "")
	tenantID := GetTenant(ctx)
	s.Empty(tenantID)
}

func (s *ContextKeyTestSuite) TestGetTenant_NoValue() {
	tenantID := GetTenant(s.ctx)
	s.Empty(tenantID)
}

func (s *ContextKeyTestSuite) TestGetTenant_WrongType() {
	// Put wrong type in context (shouldn't happen in practice)
	ctx := context.WithValue(s.ctx, tenantKey{}, 12345)
	tenantID := GetTenant(ctx)
	s.Empty(tenantID) // Should return empty, not panic
}
