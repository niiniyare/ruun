# Registry

## Overview

Registry manages schema storage, caching, and retrieval.

## Interface

```go
type Registry interface {
    Get(ctx context.Context, id string) (*Schema, error)
    Set(ctx context.Context, schema *Schema) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context) ([]*Schema, error)
    Exists(ctx context.Context, id string) (bool, error)
}
```

## Create Registry

```go
// File-based
registry := registry.NewFilesystemRegistry("./schemas")

// PostgreSQL
registry := registry.NewPostgresRegistry(db)

// With Redis cache
cached := registry.NewCachedRegistry(
    registry.NewPostgresRegistry(db),
    redisClient,
    1*time.Hour,
)
```

## Usage in Fiber Handler

```go
func HandleForm(c *fiber.Ctx) error {
    schemaID := c.Params("id")
    
    // Load schema
    schema, err := registry.Get(c.Context(), schemaID)
    if err != nil {
        return c.Status(404).SendString("Schema not found")
    }
    
    // Render
    return views.FormPage(schema, nil).Render(
        c.Context(), 
        c.Response().BodyWriter(),
    )
}
```

## Caching Strategy

```
Request
  ↓
Memory Cache (fastest, 1-5 min TTL)
  ↓ if miss
Redis Cache (fast, 1 hour TTL)
  ↓ if miss
PostgreSQL (source of truth)
  ↓
Cache in Redis + Memory
  ↓
Return schema
```

## Advanced Implementation

### Thread-Safe Operations

The registry uses `sync.RWMutex` for concurrent access protection:

```go
type Registry struct {
    storage     Storage
    cache       Storage              // Optional Redis cache
    memoryCache map[string]*cacheEntry
    mu          sync.RWMutex        // Protects memoryCache access
    parser      *Parser
}

// Thread-safe read with read lock
func (r *Registry) Get(ctx context.Context, id string) (*Schema, error) {
    r.mu.RLock()
    if entry, exists := r.memoryCache[id]; exists && time.Now().Before(entry.expiresAt) {
        r.mu.RUnlock()
        return entry.schema, nil
    }
    r.mu.RUnlock()
    // ... continue with cache miss logic
}

// Thread-safe write with write lock
func (r *Registry) cacheInMemory(id string, schema *Schema) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.memoryCache[id] = &cacheEntry{
        schema:    schema,
        expiresAt: time.Now().Add(5 * time.Minute),
    }
}
```

### Multi-Level Caching Architecture

```go
// Three-tier caching with automatic fallback
func (r *Registry) Get(ctx context.Context, id string) (*Schema, error) {
    // L1: Memory Cache (1ms) - Hot data in process memory
    r.mu.RLock()
    if entry, exists := r.memoryCache[id]; exists && time.Now().Before(entry.expiresAt) {
        r.mu.RUnlock()
        return entry.schema, nil // Fastest path
    }
    r.mu.RUnlock()

    // L2: Redis Cache (5ms) - Warm data in distributed cache
    if r.cache != nil {
        if data, err := r.cache.Get(ctx, id); err == nil {
            schema, err := r.parser.Parse(ctx, data)
            if err == nil {
                r.cacheInMemory(id, schema) // Promote to L1
                return schema, nil
            }
        }
    }

    // L3: Storage (50ms) - Cold data from persistent storage
    data, err := r.storage.Get(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("schema not found: %s", id)
    }
    
    schema, err := r.parser.Parse(ctx, data)
    if err != nil {
        return nil, fmt.Errorf("failed to parse schema: %w", err)
    }

    // Populate all cache layers (async for non-blocking)
    if r.cache != nil {
        go func() {
            r.cache.Set(context.Background(), id, data) // Populate L2
        }()
    }
    r.cacheInMemory(id, schema) // Populate L1
    
    return schema, nil
}
```

### Cache Management & Cleanup

```go
// Automatic cache cleanup with background task
func (r *Registry) StartCleanupTask() {
    ticker := time.NewTicker(1 * time.Minute)
    go func() {
        for range ticker.C {
            r.CleanupExpired()
        }
    }()
}

func (r *Registry) CleanupExpired() {
    r.mu.Lock()
    defer r.mu.Unlock()
    now := time.Now()
    for id, entry := range r.memoryCache {
        if now.After(entry.expiresAt) {
            delete(r.memoryCache, id)
        }
    }
}

// Configurable TTL per cache layer
type RegistryConfig struct {
    Storage     Storage
    Cache       Storage              // Optional Redis cache
    MemoryTTL   time.Duration       // Default: 5 minutes
    CacheTTL    time.Duration       // Default: 1 hour  
    EnableCache bool
}
```

### Schema Validation & Versioning

```go
// Validation before storage
func (r *Registry) Set(ctx context.Context, id string, schema *Schema) error {
    // Validate schema first
    if err := schema.Validate(ctx); err != nil {
        return fmt.Errorf("invalid schema: %w", err)
    }
    // ... storage logic
}

// Version-aware retrieval
func (r *Registry) GetVersion(ctx context.Context, id, version string) (*Schema, error) {
    versionedID := fmt.Sprintf("%s@%s", id, version)
    return r.Get(ctx, versionedID)
}

// Filtering capabilities
func (r *Registry) List(ctx context.Context, filter map[string]any) ([]*Schema, error) {
    ids, err := r.storage.List(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to list schema IDs: %w", err)
    }
    
    var schemas []*Schema
    for _, id := range ids {
        schema, err := r.Get(ctx, id) // Uses caching
        if err != nil {
            continue // Skip invalid schemas
        }
        
        if filter != nil && !r.matchesFilter(schema, filter) {
            continue // Apply filters
        }
        
        schemas = append(schemas, schema)
    }
    return schemas, nil
}
```

### Performance Characteristics

| Operation | L1 Cache | L2 Cache | Storage | Notes |
|-----------|----------|----------|---------|-------|
| Hot Read  | ~1ms     | -        | -       | Memory hit |
| Warm Read | ~1ms     | ~5ms     | -       | Redis hit + memory promotion |
| Cold Read | ~1ms     | ~5ms     | ~50ms   | Full cache miss |
| Write     | ~1ms     | ~5ms*    | ~50ms   | *Async cache update |

**Cache Hit Ratios** (typical production):
- L1 Memory: 85-95% for active schemas
- L2 Redis: 5-10% for recently used schemas  
- L3 Storage: 1-5% for new/rarely used schemas

---

[← Back](09-storage-interface.md) | [Next: Enricher →](11-enricher.md)