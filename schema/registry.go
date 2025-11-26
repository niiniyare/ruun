// pkg/schema/registry_v2.go
package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Registry provides centralized schema management with caching and versioning
type Registry struct {
	storage   StorageBackend
	cache     CacheBackend
	parser    *Parser
	validator schemaValidator
	events    *RegistryEventBus
	metrics   *RegistryMetrics
	config    *RegistryConfig

	mu sync.RWMutex
}


// DefaultRegistryConfig returns default configuration
func DefaultRegistryConfig() *RegistryConfig {
	return &RegistryConfig{
		EnableStorage:          true,
		StorageType:            "memory",
		EnableMemoryCache:      true,
		MemoryCacheTTL:         5 * time.Minute,
		MaxMemoryCacheSize:     1000,
		EnableDistributedCache: false,
		DistributedCacheTTL:    1 * time.Hour,
		ValidateOnStore:        true,
		ValidateOnLoad:         true,
		EnableVersioning:       true,
		MaxVersions:            10,
		EnableEvents:           true,
		EnableMetrics:          true,
	}
}

// StorageBackend defines storage operations
type StorageBackend interface {
	// Basic CRUD
	Get(ctx context.Context, id string) ([]byte, error)
	Set(ctx context.Context, id string, data []byte) error
	Delete(ctx context.Context, id string) error
	Exists(ctx context.Context, id string) (bool, error)

	// Listing
	List(ctx context.Context, filter *StorageFilter) ([]string, error)

	// Versioning
	GetVersion(ctx context.Context, id, version string) ([]byte, error)
	ListVersions(ctx context.Context, id string) ([]string, error)

	// Batch operations
	GetBatch(ctx context.Context, ids []string) (map[string][]byte, error)
	SetBatch(ctx context.Context, items map[string][]byte) error

	// Metadata
	GetMetadata(ctx context.Context, id string) (*StorageMetadata, error)

	// Health
	Health(ctx context.Context) error
}

// CacheBackend defines caching operations
type CacheBackend interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
	Clear(ctx context.Context) error
	Health(ctx context.Context) error
}

// StorageFilter defines filtering criteria


// NewRegistry creates a new schema registry
func NewRegistry(config *RegistryConfig) (*Registry, error) {
	if config == nil {
		config = DefaultRegistryConfig()
	}

	registry := &Registry{
		parser:  NewParser(),
		config:  config,
		events:  NewRegistryEventBus(),
		metrics: NewRegistryMetrics(),
	}

	// Initialize storage
	storage, err := createStorageBackend(config)
	if err != nil {
		return nil, fmt.Errorf("create storage backend: %w", err)
	}
	registry.storage = storage

	// Initialize cache if enabled
	if config.EnableMemoryCache || config.EnableDistributedCache {
		cache, err := createCacheBackend(config)
		if err != nil {
			return nil, fmt.Errorf("create cache backend: %w", err)
		}
		registry.cache = cache
	}

	// Initialize validator
	registry.validator = *newSchemaValidator()

	return registry, nil
}

// Register stores a new schema
func (r *Registry) Register(ctx context.Context, schema *Schema) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	startTime := time.Now()
	defer func() {
		if r.config.EnableMetrics {
			r.metrics.RecordOperation("register", time.Since(startTime))
		}
	}()

	// Validate schema if enabled
	if r.config.ValidateOnStore {
		if err := schema.Validate(ctx); err != nil {
			return fmt.Errorf("schema validation failed: %w", err)
		}
	}

	// Check if schema exists
	exists, err := r.storage.Exists(ctx, schema.ID)
	if err != nil {
		return fmt.Errorf("check existence: %w", err)
	}
	if exists {
		return fmt.Errorf("schema already exists: %s", schema.ID)
	}

	// Serialize schema
	data, err := json.Marshal(schema)
	if err != nil {
		return fmt.Errorf("marshal schema: %w", err)
	}

	// Store in backend
	if err := r.storage.Set(ctx, schema.ID, data); err != nil {
		return fmt.Errorf("store schema: %w", err)
	}

	// Cache if enabled
	if r.cache != nil {
		cacheKey := r.buildCacheKey(schema.ID, "")
		if err := r.cache.Set(ctx, cacheKey, data, r.config.MemoryCacheTTL); err != nil {
			// Log but don't fail
			fmt.Printf("cache set failed: %v\n", err)
		}
	}

	// Emit event
	if r.config.EnableEvents {
		r.events.Emit(&RegistryEvent{
			Type:      EventSchemaRegistered,
			SchemaID:  schema.ID,
			Timestamp: time.Now(),
		})
	}

	// Update metrics
	if r.config.EnableMetrics {
		r.metrics.IncrementSchemaCount()
	}

	return nil
}

// Get retrieves a schema by ID
func (r *Registry) Get(ctx context.Context, id string) (*Schema, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	startTime := time.Now()
	defer func() {
		if r.config.EnableMetrics {
			r.metrics.RecordOperation("get", time.Since(startTime))
		}
	}()

	// Try cache first
	if r.cache != nil {
		cacheKey := r.buildCacheKey(id, "")
		if data, err := r.cache.Get(ctx, cacheKey); err == nil {
			if r.config.EnableMetrics {
				r.metrics.IncrementCacheHit()
			}
			return r.parser.Parse(ctx, data)
		}
		if r.config.EnableMetrics {
			r.metrics.IncrementCacheMiss()
		}
	}

	// Load from storage
	data, err := r.storage.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("storage get: %w", err)
	}

	// Parse schema
	schema, err := r.parser.Parse(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("parse schema: %w", err)
	}

	// Validate if enabled
	if r.config.ValidateOnLoad {
		if err := schema.Validate(ctx); err != nil {
			return nil, fmt.Errorf("schema validation failed: %w", err)
		}
	}

	// Update cache
	if r.cache != nil {
		cacheKey := r.buildCacheKey(id, "")
		if err := r.cache.Set(ctx, cacheKey, data, r.config.MemoryCacheTTL); err != nil {
			// Log but don't fail
			fmt.Printf("cache set failed: %v\n", err)
		}
	}

	// Emit event
	if r.config.EnableEvents {
		r.events.Emit(&RegistryEvent{
			Type:      EventSchemaAccessed,
			SchemaID:  id,
			Timestamp: time.Now(),
		})
	}

	return schema, nil
}

// Update updates an existing schema
func (r *Registry) Update(ctx context.Context, schema *Schema) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	startTime := time.Now()
	defer func() {
		if r.config.EnableMetrics {
			r.metrics.RecordOperation("update", time.Since(startTime))
		}
	}()

	// Check if exists
	exists, err := r.storage.Exists(ctx, schema.ID)
	if err != nil {
		return fmt.Errorf("check existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("schema not found: %s", schema.ID)
	}

	// Validate if enabled
	if r.config.ValidateOnStore {
		if err := schema.Validate(ctx); err != nil {
			return fmt.Errorf("schema validation failed: %w", err)
		}
	}

	// Serialize
	data, err := json.Marshal(schema)
	if err != nil {
		return fmt.Errorf("marshal schema: %w", err)
	}

	// Store
	if err := r.storage.Set(ctx, schema.ID, data); err != nil {
		return fmt.Errorf("store schema: %w", err)
	}

	// Invalidate cache
	if r.cache != nil {
		cacheKey := r.buildCacheKey(schema.ID, "")
		if err := r.cache.Delete(ctx, cacheKey); err != nil {
			// Log but don't fail
			fmt.Printf("cache delete failed: %v\n", err)
		}
	}

	// Emit event
	if r.config.EnableEvents {
		r.events.Emit(&RegistryEvent{
			Type:      EventSchemaUpdated,
			SchemaID:  schema.ID,
			Timestamp: time.Now(),
		})
	}

	return nil
}

// Delete removes a schema
func (r *Registry) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	startTime := time.Now()
	defer func() {
		if r.config.EnableMetrics {
			r.metrics.RecordOperation("delete", time.Since(startTime))
		}
	}()

	// Delete from storage
	if err := r.storage.Delete(ctx, id); err != nil {
		return fmt.Errorf("storage delete: %w", err)
	}

	// Invalidate cache
	if r.cache != nil {
		cacheKey := r.buildCacheKey(id, "")
		if err := r.cache.Delete(ctx, cacheKey); err != nil {
			// Log but don't fail
			fmt.Printf("cache delete failed: %v\n", err)
		}
	}

	// Emit event
	if r.config.EnableEvents {
		r.events.Emit(&RegistryEvent{
			Type:      EventSchemaDeleted,
			SchemaID:  id,
			Timestamp: time.Now(),
		})
	}

	// Update metrics
	if r.config.EnableMetrics {
		r.metrics.DecrementSchemaCount()
	}

	return nil
}

// List returns schemas matching filter
func (r *Registry) List(ctx context.Context, filter *StorageFilter) ([]*Schema, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	startTime := time.Now()
	defer func() {
		if r.config.EnableMetrics {
			r.metrics.RecordOperation("list", time.Since(startTime))
		}
	}()

	// Get IDs from storage
	ids, err := r.storage.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("storage list: %w", err)
	}

	// Load schemas
	schemas := make([]*Schema, 0, len(ids))
	for _, id := range ids {
		schema, err := r.Get(ctx, id)
		if err != nil {
			// Log error but continue
			fmt.Printf("load schema %s: %v\n", id, err)
			continue
		}
		schemas = append(schemas, schema)
	}

	return schemas, nil
}

// GetVersion retrieves a specific version
func (r *Registry) GetVersion(ctx context.Context, id, version string) (*Schema, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.config.EnableVersioning {
		return nil, fmt.Errorf("versioning not enabled")
	}

	// Try cache
	if r.cache != nil {
		cacheKey := r.buildCacheKey(id, version)
		if data, err := r.cache.Get(ctx, cacheKey); err == nil {
			if r.config.EnableMetrics {
				r.metrics.IncrementCacheHit()
			}
			return r.parser.Parse(ctx, data)
		}
		if r.config.EnableMetrics {
			r.metrics.IncrementCacheMiss()
		}
	}

	// Load from storage
	data, err := r.storage.GetVersion(ctx, id, version)
	if err != nil {
		return nil, fmt.Errorf("storage get version: %w", err)
	}

	// Parse
	schema, err := r.parser.Parse(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("parse schema: %w", err)
	}

	// Cache
	if r.cache != nil {
		cacheKey := r.buildCacheKey(id, version)
		if err := r.cache.Set(ctx, cacheKey, data, r.config.DistributedCacheTTL); err != nil {
			fmt.Printf("cache set failed: %v\n", err)
		}
	}

	return schema, nil
}

// ListVersions returns all versions of a schema
func (r *Registry) ListVersions(ctx context.Context, id string) ([]string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.config.EnableVersioning {
		return nil, fmt.Errorf("versioning not enabled")
	}

	return r.storage.ListVersions(ctx, id)
}

// Exists checks if schema exists
func (r *Registry) Exists(ctx context.Context, id string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.storage.Exists(ctx, id)
}

// GetMetadata returns schema metadata
func (r *Registry) GetMetadata(ctx context.Context, id string) (*StorageMetadata, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.storage.GetMetadata(ctx, id)
}

// Health checks registry health
func (r *Registry) Health(ctx context.Context) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Check storage
	if err := r.storage.Health(ctx); err != nil {
		return fmt.Errorf("storage health: %w", err)
	}

	// Check cache
	if r.cache != nil {
		if err := r.cache.Health(ctx); err != nil {
			return fmt.Errorf("cache health: %w", err)
		}
	}

	return nil
}

// ClearCache clears all cached schemas
func (r *Registry) ClearCache(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cache == nil {
		return nil
	}

	return r.cache.Clear(ctx)
}

// InvalidateCache invalidates cache for specific schema
func (r *Registry) InvalidateCache(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cache == nil {
		return nil
	}

	cacheKey := r.buildCacheKey(id, "")
	return r.cache.Delete(ctx, cacheKey)
}

// GetMetrics returns registry metrics
func (r *Registry) GetMetrics() *RegistryMetrics {
	return r.metrics
}

// Subscribe subscribes to registry events
func (r *Registry) Subscribe(handler RegistryEventHandler) {
	if r.config.EnableEvents {
		r.events.Subscribe(handler)
	}
}

// buildCacheKey builds cache key
func (r *Registry) buildCacheKey(id, version string) string {
	if version == "" {
		return fmt.Sprintf("schema:%s", id)
	}
	return fmt.Sprintf("schema:%s:v:%s", id, version)
}

// Helper functions to create backends
func createStorageBackend(config *RegistryConfig) (StorageBackend, error) {
	switch config.StorageType {
	case "memory":
		return NewMemoryStorage(), nil
	case "file":
		return NewFileStorage("./schemas")
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", config.StorageType)
	}
}

func createCacheBackend(config *RegistryConfig) (CacheBackend, error) {
	if config.EnableMemoryCache {
		return NewMemoryCache(config.MaxMemoryCacheSize), nil
	}
	return nil, fmt.Errorf("no cache backend configured")
}
