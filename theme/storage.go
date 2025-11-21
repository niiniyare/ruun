package theme

import (
	"context"
	"sync"
)

// Storage is a minimal interface for theme persistence.
// Similar to io.Reader/Writer pattern, this allows users to implement
// their own backends (database, filesystem, CDN, etc.).
type Storage interface {
	// GetTheme retrieves a theme by ID.
	GetTheme(ctx context.Context, themeID string) (*Theme, error)
	
	// SaveTheme stores a theme.
	SaveTheme(ctx context.Context, theme *Theme) error
	
	// DeleteTheme removes a theme.
	DeleteTheme(ctx context.Context, themeID string) error
	
	// ListThemes returns all stored themes.
	ListThemes(ctx context.Context) ([]*Theme, error)
}

// MemoryStorage is an in-memory implementation of Storage for development and testing.
type MemoryStorage struct {
	themes map[string]*Theme
	mu     sync.RWMutex
}

// NewMemoryStorage creates a new in-memory storage backend.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		themes: make(map[string]*Theme),
	}
}

// GetTheme retrieves a theme from memory.
func (s *MemoryStorage) GetTheme(ctx context.Context, themeID string) (*Theme, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	theme, exists := s.themes[themeID]
	if !exists {
		return nil, NewErrorf(ErrCodeNotFound, "theme not found: %s", themeID)
	}

	return theme.Clone(), nil
}

// SaveTheme stores a theme in memory.
func (s *MemoryStorage) SaveTheme(ctx context.Context, theme *Theme) error {
	if theme == nil {
		return NewError(ErrCodeValidation, "theme cannot be nil")
	}
	if theme.ID == "" {
		return NewError(ErrCodeValidation, "theme ID cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.themes[theme.ID] = theme.Clone()
	return nil
}

// DeleteTheme removes a theme from memory.
func (s *MemoryStorage) DeleteTheme(ctx context.Context, themeID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.themes[themeID]; !exists {
		return NewErrorf(ErrCodeNotFound, "theme not found: %s", themeID)
	}

	delete(s.themes, themeID)
	return nil
}

// ListThemes returns all themes from memory.
func (s *MemoryStorage) ListThemes(ctx context.Context) ([]*Theme, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	themes := make([]*Theme, 0, len(s.themes))
	for _, theme := range s.themes {
		themes = append(themes, theme.Clone())
	}

	return themes, nil
}

// TenantStorage is an interface for tenant configuration persistence.
type TenantStorage interface {
	GetTenant(ctx context.Context, tenantID string) (*TenantConfig, error)
	SaveTenant(ctx context.Context, config *TenantConfig) error
	DeleteTenant(ctx context.Context, tenantID string) error
	ListTenants(ctx context.Context) ([]*TenantConfig, error)
}

// SimpleTenantStorage is an in-memory implementation of TenantStorage.
type SimpleTenantStorage struct {
	tenants map[string]*TenantConfig
	mu      sync.RWMutex
}

// NewSimpleTenantStorage creates a new in-memory tenant storage.
func NewSimpleTenantStorage() *SimpleTenantStorage {
	return &SimpleTenantStorage{
		tenants: make(map[string]*TenantConfig),
	}
}

// GetTenant retrieves a tenant from memory.
func (s *SimpleTenantStorage) GetTenant(ctx context.Context, tenantID string) (*TenantConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	config, exists := s.tenants[tenantID]
	if !exists {
		return nil, NewErrorf(ErrCodeNotFound, "tenant not found: %s", tenantID)
	}

	return config, nil
}

// SaveTenant stores a tenant in memory.
func (s *SimpleTenantStorage) SaveTenant(ctx context.Context, config *TenantConfig) error {
	if config == nil {
		return NewError(ErrCodeValidation, "tenant config cannot be nil")
	}
	if config.TenantID == "" {
		return NewError(ErrCodeValidation, "tenant ID cannot be empty")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.tenants[config.TenantID] = config
	return nil
}

// DeleteTenant removes a tenant from memory.
func (s *SimpleTenantStorage) DeleteTenant(ctx context.Context, tenantID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tenants[tenantID]; !exists {
		return NewErrorf(ErrCodeNotFound, "tenant not found: %s", tenantID)
	}

	delete(s.tenants, tenantID)
	return nil
}

// ListTenants returns all tenants from memory.
func (s *SimpleTenantStorage) ListTenants(ctx context.Context) ([]*TenantConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	configs := make([]*TenantConfig, 0, len(s.tenants))
	for _, config := range s.tenants {
		configs = append(configs, config)
	}

	return configs, nil
}
