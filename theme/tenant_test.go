package theme

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// MockStorage is a mock for the Storage interface
type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetTheme(ctx context.Context, id string) (*Theme, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Theme), args.Error(1)
}

func (m *MockStorage) SaveTheme(ctx context.Context, theme *Theme) error {
	args := m.Called(ctx, theme)
	return args.Error(0)
}

func (m *MockStorage) DeleteTheme(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockStorage) ListThemes(ctx context.Context) ([]*Theme, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*Theme), args.Error(1)
}

func (m *MockStorage) GetTenant(ctx context.Context, tenantID string) (*TenantConfig, error) {
	args := m.Called(ctx, tenantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*TenantConfig), args.Error(1)
}

func (m *MockStorage) SaveTenant(ctx context.Context, config *TenantConfig) error {
	args := m.Called(ctx, config)
	return args.Error(0)
}

func (m *MockStorage) DeleteTenant(ctx context.Context, tenantID string) error {
	args := m.Called(ctx, tenantID)
	return args.Error(0)
}

func (m *MockStorage) ListTenants(ctx context.Context) ([]*TenantConfig, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*TenantConfig), args.Error(1)
}

type TenantManagerTestSuite struct {
	suite.Suite
	manager       *Manager
	mockStorage   *MockStorage
	tenantManager *TenantManager
	ctx           context.Context
}

func (s *TenantManagerTestSuite) SetupTest() {
	s.mockStorage = new(MockStorage)

	managerConfig := DefaultManagerConfig()
	// Disable caching in tests for predictability, unless testing cache behavior
	managerConfig.EnableCaching = false
	managerConfig.CompilerConfig.EnableCaching = false
	managerConfig.ResolverConfig.EnableCaching = false

	var err error
	s.manager, err = NewManager(managerConfig, s.mockStorage, nil)
	require.NoError(s.T(), err)

	s.tenantManager = NewTenantManager(s.manager, s.mockStorage)
	s.ctx = context.Background()
}

func TestTenantManagerTestSuite(t *testing.T) {
	suite.Run(t, new(TenantManagerTestSuite))
}

func newTestTheme(id string) *Theme {
	return &Theme{
		ID:   id,
		Name: id,
		Tokens: &Tokens{
			Primitives: &PrimitiveTokens{
				Colors: map[string]string{
					"primary":   "#0000FF",
					"secondary": "#00FF00",
					"text":      "#000000",
				},
			},
			Semantic: &SemanticTokens{
				Colors: map[string]string{
					"button.background": "primitives.colors.primary",
					"button.text":       "primitives.colors.text",
				},
			},
		},
		CustomCSS: ".button { background-color: var(--awo-semantic-color-button-background); color: var(--awo-semantic-color-button-text); }",
	}
}

func (s *TenantManagerTestSuite) TestConfigureTenant() {
	s.T().Run("new tenant", func(t *testing.T) {
		s.SetupTest()
		config := &TenantConfig{
			TenantID:     "tenant-1",
			DefaultTheme: "dark",
		}

		s.mockStorage.On("GetTheme", s.ctx, "dark").Return(newTestTheme("dark"), nil).Once()
		s.mockStorage.On("GetTenant", s.ctx, "tenant-1").Return(nil, NewError(ErrCodeNotFound, "not found")).Once()
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.Equal(t, "tenant-1", savedConfig.TenantID)
			require.NotZero(t, savedConfig.CreatedAt)
			require.NotZero(t, savedConfig.UpdatedAt)
		})

		err := s.tenantManager.ConfigureTenant(s.ctx, config)
		require.NoError(t, err)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("update existing tenant", func(t *testing.T) {
		s.SetupTest()
		existingCreatedAt := time.Now().Add(-24 * time.Hour)
		existingConfig := &TenantConfig{
			TenantID:  "tenant-1",
			CreatedAt: existingCreatedAt,
		}
		config := &TenantConfig{
			TenantID:     "tenant-1",
			DefaultTheme: "light",
		}

		s.mockStorage.On("GetTheme", s.ctx, "light").Return(newTestTheme("light"), nil).Once()
		s.mockStorage.On("GetTenant", s.ctx, "tenant-1").Return(existingConfig, nil).Once()
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.Equal(t, "tenant-1", savedConfig.TenantID)
			require.Equal(t, existingCreatedAt, savedConfig.CreatedAt)
			require.True(t, savedConfig.UpdatedAt.After(existingCreatedAt))
		})

		err := s.tenantManager.ConfigureTenant(s.ctx, config)
		require.NoError(t, err)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("invalid default theme", func(t *testing.T) {
		s.SetupTest()
		config := &TenantConfig{
			TenantID:     "tenant-1",
			DefaultTheme: "non-existent",
		}

		s.mockStorage.On("GetTheme", s.ctx, "non-existent").Return(nil, NewError(ErrCodeNotFound, "not found")).Once()

		err := s.tenantManager.ConfigureTenant(s.ctx, config)
		require.Error(t, err)
		var e *Error
		require.ErrorAs(t, err, &e)
		require.Equal(t, ErrCodeNotFound, e.Code)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("nil config", func(t *testing.T) {
		s.SetupTest()
		err := s.tenantManager.ConfigureTenant(s.ctx, nil)
		require.Error(t, err)
		var e *Error
		require.ErrorAs(t, err, &e)
		require.Equal(t, ErrCodeValidation, e.Code)
	})

	s.T().Run("empty tenant id", func(t *testing.T) {
		s.SetupTest()
		err := s.tenantManager.ConfigureTenant(s.ctx, &TenantConfig{TenantID: ""})
		require.Error(t, err)
		var e *Error
		require.ErrorAs(t, err, &e)
		require.Equal(t, ErrCodeValidation, e.Code)
	})
}

func (s *TenantManagerTestSuite) TestGetTenantConfig() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		config := &TenantConfig{TenantID: "tenant-1"}
		s.mockStorage.On("GetTenant", s.ctx, "tenant-1").Return(config, nil).Once()

		retrieved, err := s.tenantManager.GetTenantConfig(s.ctx, "tenant-1")
		require.NoError(t, err)
		require.Equal(t, config, retrieved)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("not found", func(t *testing.T) {
		s.SetupTest()
		s.mockStorage.On("GetTenant", s.ctx, "tenant-1").Return(nil, NewError(ErrCodeNotFound, "not found")).Once()

		_, err := s.tenantManager.GetTenantConfig(s.ctx, "tenant-1")
		require.Error(t, err)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestGetTenantTheme() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		themeID := "default"
		config := &TenantConfig{
			TenantID:     tenantID,
			DefaultTheme: themeID,
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("GetTheme", mock.Anything, themeID).Return(newTestTheme(themeID), nil).Once()

		compiledTheme, err := s.tenantManager.GetTenantTheme(s.ctx, tenantID, nil)
		require.NoError(t, err)
		require.NotNil(t, compiledTheme)
		require.Equal(t, themeID, compiledTheme.Theme.ID)
		require.Contains(t, compiledTheme.CSS, "--awo-primitives-color-primary:#0000FF")
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("with branding overrides", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		themeID := "default"
		config := &TenantConfig{
			TenantID:     tenantID,
			DefaultTheme: themeID,
			Branding: &BrandingOverrides{
				PrimaryColor:   "#FF0000",
				SecondaryColor: "#FFFF00",
				CustomCSS:      ".custom { color: red; }",
			},
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("GetTheme", mock.Anything, themeID).Return(newTestTheme(themeID), nil).Once()

		compiledTheme, err := s.tenantManager.GetTenantTheme(s.ctx, tenantID, nil)
		require.NoError(t, err)
		require.NotNil(t, compiledTheme)
		require.Equal(t, themeID, compiledTheme.Theme.ID)
		require.Contains(t, compiledTheme.CSS, "--awo-primitives-color-primary:#FF0000")
		require.Contains(t, compiledTheme.CSS, ".custom { color: red; }")
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("with custom tokens", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		themeID := "default"
		config := &TenantConfig{
			TenantID:     tenantID,
			DefaultTheme: themeID,
			Branding: &BrandingOverrides{
				CustomTokens: map[string]string{
					"primitives.colors.primary": "#0000AA",
				},
			},
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("GetTheme", mock.Anything, themeID).Return(newTestTheme(themeID), nil).Once()

		compiledTheme, err := s.tenantManager.GetTenantTheme(s.ctx, tenantID, nil)
		require.NoError(t, err)
		require.NotNil(t, compiledTheme)
		require.Contains(t, compiledTheme.CSS, "--awo-primitives-color-primary:#0000AA")
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("tenant has no default theme", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		config := &TenantConfig{TenantID: tenantID}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()

		_, err := s.tenantManager.GetTenantTheme(s.ctx, tenantID, nil)
		require.Error(t, err)
		var e *Error
		require.ErrorAs(t, err, &e)
		require.Equal(t, ErrCodeValidation, e.Code)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestListTenantThemes() {
	s.T().Run("all themes allowed", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		config := &TenantConfig{TenantID: tenantID, AllowedThemes: []string{}}
		themes := []*Theme{newTestTheme("dark"), newTestTheme("light")}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("ListThemes", s.ctx).Return(themes, nil).Once()

		result, err := s.tenantManager.ListTenantThemes(s.ctx, tenantID)
		require.NoError(t, err)
		require.Len(t, result, 2)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("subset of themes allowed", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		config := &TenantConfig{TenantID: tenantID, AllowedThemes: []string{"dark"}}
		themes := []*Theme{newTestTheme("dark"), newTestTheme("light")}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("ListThemes", s.ctx).Return(themes, nil).Once()

		result, err := s.tenantManager.ListTenantThemes(s.ctx, tenantID)
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, "dark", result[0].ID)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestSetTenantTheme() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		themeID := "new-theme"
		config := &TenantConfig{
			TenantID:     tenantID,
			DefaultTheme: "old-theme",
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("GetTheme", s.ctx, themeID).Return(newTestTheme(themeID), nil).Once()
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.Equal(t, themeID, savedConfig.DefaultTheme)
		})

		err := s.tenantManager.SetTenantTheme(s.ctx, tenantID, themeID)
		require.NoError(t, err)
		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("theme not allowed", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		themeID := "not-allowed-theme"
		config := &TenantConfig{
			TenantID:      tenantID,
			DefaultTheme:  "old-theme",
			AllowedThemes: []string{"theme1", "theme2"},
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("GetTheme", s.ctx, themeID).Return(newTestTheme(themeID), nil).Once()

		err := s.tenantManager.SetTenantTheme(s.ctx, tenantID, themeID)
		require.Error(t, err)
		var e *Error
		require.ErrorAs(t, err, &e)
		require.Equal(t, ErrCodeValidation, e.Code)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestSetBranding() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		branding := &BrandingOverrides{PrimaryColor: "#123456"}
		config := &TenantConfig{TenantID: tenantID}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil).Once()
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.Equal(t, branding, savedConfig.Branding)
			require.NotZero(t, savedConfig.UpdatedAt)
		})

		err := s.tenantManager.SetBranding(s.ctx, tenantID, branding)
		require.NoError(t, err)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestFeatures() {
	tenantID := "tenant-1"
	feature := "new-feature"

	s.T().Run("enable and check feature", func(t *testing.T) {
		s.SetupTest()
		config := &TenantConfig{TenantID: tenantID}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil) // Called by Enable and IsEnabled
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.True(t, savedConfig.Features[feature])
		})

		err := s.tenantManager.EnableFeature(s.ctx, tenantID, feature)
		require.NoError(t, err)

		enabled := s.tenantManager.IsFeatureEnabled(s.ctx, tenantID, feature)
		require.True(t, enabled)

		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("disable feature", func(t *testing.T) {
		s.SetupTest()
		config := &TenantConfig{
			TenantID: tenantID,
			Features: map[string]bool{feature: true},
		}

		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(config, nil)
		s.mockStorage.On("SaveTenant", s.ctx, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
			savedConfig := args.Get(1).(*TenantConfig)
			require.False(t, savedConfig.Features[feature])
		})

		err := s.tenantManager.DisableFeature(s.ctx, tenantID, feature)
		require.NoError(t, err)

		enabled := s.tenantManager.IsFeatureEnabled(s.ctx, tenantID, feature)
		require.False(t, enabled)

		s.mockStorage.AssertExpectations(t)
	})

	s.T().Run("is feature enabled on non-existent tenant", func(t *testing.T) {
		s.SetupTest()
		s.mockStorage.On("GetTenant", s.ctx, tenantID).Return(nil, errors.New("not found"))
		enabled := s.tenantManager.IsFeatureEnabled(s.ctx, tenantID, feature)
		require.False(t, enabled)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestDeleteTenant() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		tenantID := "tenant-1"
		s.mockStorage.On("DeleteTenant", s.ctx, tenantID).Return(nil).Once()

		err := s.tenantManager.DeleteTenant(s.ctx, tenantID)
		require.NoError(t, err)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestListTenants() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		configs := []*TenantConfig{{TenantID: "t1"}, {TenantID: "t2"}}
		s.mockStorage.On("ListTenants", s.ctx).Return(configs, nil).Once()

		result, err := s.tenantManager.ListTenants(s.ctx)
		require.NoError(t, err)
		require.Equal(t, configs, result)
		s.mockStorage.AssertExpectations(t)
	})
}

func (s *TenantManagerTestSuite) TestGetTenantStats() {
	s.T().Run("success", func(t *testing.T) {
		s.SetupTest()
		now := time.Now()
		config := &TenantConfig{
			TenantID:      "t1",
			DefaultTheme:  "dark",
			AllowedThemes: []string{"dark", "light"},
			Features:      map[string]bool{"f1": true},
			Branding:      &BrandingOverrides{},
			UpdatedAt:     now,
		}
		s.mockStorage.On("GetTenant", s.ctx, "t1").Return(config, nil).Once()

		stats, err := s.tenantManager.GetTenantStats(s.ctx, "t1")
		require.NoError(t, err)
		require.NotNil(t, stats)
		require.Equal(t, "t1", stats.TenantID)
		require.Equal(t, "dark", stats.DefaultTheme)
		require.Equal(t, 2, stats.ThemeCount)
		require.Equal(t, 1, stats.FeatureCount)
		require.True(t, stats.HasBranding)
		require.Equal(t, now, stats.LastUpdated)
		s.mockStorage.AssertExpectations(t)
	})
}
