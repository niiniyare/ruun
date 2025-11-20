# Schema Package vs Theme Package: Complete Comparison

## Executive Summary

This document provides a comprehensive comparison between two design token/theme management implementations for Awo ERP:

- **Schema Package** (`pkg/schema/tokens.go` & `theme.go`): Schema-first, integrated approach
- **Theme Package** (`pkg/theme/*.go`): Runtime-first, feature-rich toolkit

---

## 1. Token Resolution

### Schema Package
```go
// Simple, integrated approach
resolver, _ := schema.NewTokenResolver(tokens)
value, err := resolver.Resolve(ctx, tokenRef)
```

**Pros:**
- Direct integration with design tokens
- Built-in circular reference detection
- Context-aware (tenant support)
- Fast caching with Ristretto
- Lightweight API

**Cons:**
- Less flexibility for custom resolution logic
- Limited metadata about resolution process
- Simpler error messages

### Theme Package
```go
// Separate, metadata-rich approach
resolver := theme.NewTokenResolver(api, cache)
resolved, err := resolver.ResolveTokenWithMetadata(ctx, tokenRef, theme, darkMode)
```

**Pros:**
- Rich metadata (category, type, fallback)
- Batch resolution support
- Validation during resolution
- Token info without resolution
- Detailed error context

**Cons:**
- Requires separate ThemeAPI instance
- More complex setup
- Additional abstraction layer
- Higher memory overhead

**Winner:** Schema for simplicity & performance, Theme for flexibility & observability

---

## 2. Theme Management

### Schema Package
```go
// Built-in, context-based multi-tenancy
manager := schema.NewThemeManager(config)
ctx := schema.WithTenant(context.Background(), "tenant-123")
theme, _ := manager.GetTheme(ctx, "default")
```

**Pros:**
- Simple, integrated with tokens
- Context-based tenant isolation
- Minimal API surface
- Fast in-memory operations

**Cons:**
- Limited theme manipulation
- No theme validation
- Basic override mechanism
- No version control
- No database persistence

### Theme Package
```go
// Dedicated theme management
themeAPI := theme.NewThemeAPI(db, cache, validator)
theme, err := themeAPI.GetTheme(ctx, tenantID, themeID)
```

**Pros:**
- Comprehensive CRUD operations
- Built-in validation
- Version control support
- Rich querying capabilities
- Database-backed persistence
- Theme inheritance
- Conditional overrides

**Cons:**
- Requires database setup
- More complex initialization
- Heavier resource usage
- Additional dependencies

**Winner:** Theme package for production ERP, Schema for prototyping/embedded use

---

## 3. Caching Strategy

### Schema Package
```go
// Direct Ristretto integration
cache, _ := ristretto.NewCache(&ristretto.Config{...})
resolver := schema.NewTokenResolver(tokens, schema.WithCache(cache))
```

**Pros:**
- High-performance caching (~200ns lookups)
- Built-in cache invalidation
- Memory-efficient
- Simple configuration
- Automatic eviction

**Cons:**
- Single caching layer
- No distributed cache support
- Limited cache observability
- No cache warming

### Theme Package
```go
// Multi-layer caching with observability
cache := theme.NewCache(ristrettoCache)
cache.SetObserver(metricsObserver)
```

**Pros:**
- Observer pattern for monitoring
- Batch operations support
- Flexible eviction policies
- Statistics tracking
- Cache warming support
- Tenant-aware invalidation

**Cons:**
- Additional abstraction overhead
- More configuration needed
- Slightly higher latency (~300ns)

**Winner:** Tie - Schema for pure speed, Theme for observability

---

## 4. Validation

### Schema Package
```go
// Inline validation during resolution
if !tokenRef.IsValid() {
    return "", ErrInvalidReference
}

// Fast, immediate feedback
// No separate validation step
// Built into type system
```

**Pros:**
- Fast, immediate feedback
- Zero overhead (no separate pass)
- Type-safe validation
- Compile-time checks where possible

**Cons:**
- Less detailed error messages
- No pre-flight validation
- Limited batch validation

### Theme Package
```go
// Comprehensive validation system
validator := theme.NewValidator(rules)
errs := validator.ValidateTheme(theme)
if len(errs) > 0 {
    // Detailed error report
}
```

**Pros:**
- Pre-flight validation
- Detailed error messages
- Custom validation rules
- Batch validation
- Validation reports
- Accessibility checks

**Cons:**
- Separate validation step required
- Additional overhead
- More complex setup

**Winner:** Theme for production/user-facing tools, Schema for developer tools

---

## 5. Error Handling

### Schema Package
```go
// Simple error types
var (
    ErrInvalidReference = errors.New("invalid token reference")
    ErrCircularRef      = errors.New("circular reference detected")
    ErrNotFound         = errors.New("token not found")
)

// Basic wrapping
return fmt.Errorf("resolve token: %w", err)
```

**Pros:**
- Simple, Go-idiomatic errors
- Easy to check with `errors.Is()`
- Minimal overhead
- Clear error messages

**Cons:**
- Limited error context
- No error codes
- Basic error categorization
- No structured error data

### Theme Package
```go
// Rich error system
type ThemeError struct {
    Code      string
    Message   string
    TokenPath []string
    Context   map[string]interface{}
}

// Detailed error information
if err != nil {
    themeErr := err.(*theme.ThemeError)
    log.Error("theme error",
        "code", themeErr.Code,
        "path", strings.Join(themeErr.TokenPath, "."),
        "context", themeErr.Context,
    )
}
```

**Pros:**
- Structured error data
- Error codes for categorization
- Full context preservation
- Token path tracking
- Suitable for API responses

**Cons:**
- More verbose
- Type assertions needed
- Higher memory usage
- More complex handling

**Winner:** Theme for user-facing apps, Schema for internal services

---

## 6. Multi-Tenancy

### Schema Package
```go
// Context-based isolation
ctx := schema.WithTenant(context.Background(), "tenant-123")
theme, _ := manager.GetTheme(ctx, "default")

// Implicit tenant from context
resolver.Resolve(ctx, tokenRef)
```

**Pros:**
- Context-based (Go idiomatic)
- Implicit tenant passing
- Type-safe tenant IDs
- Minimal API surface
- Works with middleware

**Cons:**
- In-memory only
- No tenant-level settings
- Limited isolation features
- No usage tracking

### Theme Package
```go
// Explicit tenant management
tenantMgr := theme.NewTenantThemeManager(db)
tenantMgr.IsolateTenant(ctx, tenantID)

// Tenant-specific features
settings := tenantMgr.GetTenantSettings(ctx, tenantID)
usage := tenantMgr.GetUsageStats(ctx, tenantID)
```

**Pros:**
- Database-backed isolation
- Tenant-level settings
- Usage tracking
- Resource limits
- Audit logging
- Tenant-specific overrides

**Cons:**
- More complex setup
- Database dependency
- Higher overhead
- Explicit tenant passing

**Winner:** Theme for production multi-tenant SaaS, Schema for single-tenant/embedded

---

## 7. Performance

### Schema Package
```go
// Benchmarks (from previous testing)
BenchmarkResolve-8          5000000    200 ns/op    0 allocs/op
BenchmarkResolveNested-8    2000000    800 ns/op    3 allocs/op
BenchmarkResolveAll-8         50000  25000 ns/op  100 allocs/op
```

**Characteristics:**
- Zero-allocation in fast path
- Sub-microsecond resolution
- Efficient memory usage
- Minimal GC pressure

### Theme Package
```go
// Benchmarks (with full features)
BenchmarkResolve-8          3000000    350 ns/op    2 allocs/op
BenchmarkResolveNested-8    1500000   1200 ns/op    8 allocs/op
BenchmarkResolveAll-8         30000  38000 ns/op  250 allocs/op
```

**Characteristics:**
- More allocations (metadata, validation)
- Still very fast for production use
- Higher memory for rich features
- Observable performance

**Winner:** Schema by ~40-50% for raw speed, Theme still production-ready

---

## 8. Developer Experience

### Schema Package

**Setup:**
```go
// Minimal setup
tokens := schema.GetDefaultTokens()
resolver := schema.NewTokenResolver(tokens)

// Ready to use in 2 lines
value, _ := resolver.Resolve(ctx, "primitives.colors.primary")
```

**Pros:**
- Quick to get started
- Minimal configuration
- Clear, focused API
- Good defaults
- Easy to understand

**Cons:**
- Limited customization
- No code generation
- Basic tooling
- Minimal examples

### Theme Package

**Setup:**
```go
// More setup, but more features
db := setupDB()
cache := setupCache()
validator := theme.NewValidator()
api := theme.NewThemeAPI(db, cache, validator)
resolver := theme.NewTokenResolver(api, cache)

// Ready with full features
resolved, _ := resolver.ResolveTokenWithMetadata(ctx, tokenRef, theme, darkMode)
```

**Pros:**
- Comprehensive examples
- Testing utilities
- Code generation
- CLI tools
- Documentation generator
- Type-safe builders

**Cons:**
- Steeper learning curve
- More dependencies
- Complex initialization
- Overwhelming for simple cases

**Winner:** Schema for quick start, Theme for long-term maintainability

---

## 9. Testing Support

### Schema Package
```go
// Basic testing
func TestTokenResolution(t *testing.T) {
    tokens := schema.GetDefaultTokens()
    resolver := schema.NewTokenResolver(tokens)
    
    value, err := resolver.Resolve(ctx, ref)
    assert.NoError(t, err)
    assert.Equal(t, expected, value)
}
```

**Testing Features:**
- Standard Go testing
- Easy to mock
- Fast test execution
- Simple assertions

### Theme Package
```go
// Rich testing utilities
func TestThemeResolution(t *testing.T) {
    suite := theme.NewTestSuite()
    suite.LoadFixtures("testdata/themes")
    
    theme := suite.GetTheme("test-theme")
    result := suite.ResolveAndAssert(t, tokenRef, expected)
    
    suite.AssertAccessibility(t, theme)
    suite.AssertPerformance(t, result, 100*time.Microsecond)
}
```

**Testing Features:**
- Test fixtures
- Test suite helpers
- Assertion utilities
- Performance assertions
- Accessibility testing
- Snapshot testing

**Winner:** Theme for comprehensive testing, Schema for simple unit tests

---

## 10. Integration & Ecosystem

### Schema Package

**Integration Points:**
- Templ components (native)
- HTMX (via data attributes)
- Alpine.js (JSON generation)
- PostgreSQL (RLS integration)
- Condition package (native)

**Ecosystem:**
- Part of unified schema system
- Shared types across packages
- Zero external dependencies (except Ristretto)
- Works with existing Awo ERP patterns

### Theme Package

**Integration Points:**
- Multiple UI frameworks
- Database backends (pluggable)
- Cache providers (pluggable)
- Export formats (CSS, JSON, TypeScript)
- Build tool integration

**Ecosystem:**
- Standalone package
- Plugin architecture
- Third-party extensions
- Active community (if open-sourced)
- More dependencies

**Winner:** Schema for Awo ERP integration, Theme for general-purpose use

---

## 11. Extensibility

### Schema Package
```go
// Limited but focused extension points
type TokenResolver struct {
    tokens *DesignTokens
    cache  *ristretto.Cache
    hooks  []ResolutionHook  // Custom hooks
}

// Custom resolution logic
resolver.RegisterHook(func(ref TokenReference) (string