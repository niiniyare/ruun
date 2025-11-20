// Package theme provides a production-ready, enterprise-grade theme management system
// for multi-tenant applications with comprehensive token resolution, CSS compilation,
// validation, and runtime management capabilities.
//
// # Architecture
//
// The package is organized into logical layers:
//
//  ┌─────────────────────────────────────────────────────────────┐
//  │                    Application Layer                        │
//  │  (Schema Enricher, Runtime Builder, UI Components)         │
//  └─────────────────────────────────────────────────────────────┘
//                              ▼
//  ┌─────────────────────────────────────────────────────────────┐
//  │                    Theme Package (pkg/theme)                │
//  │                                                             │
//  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │
//  │  │   Tokens     │  │    Theme     │  │   Manager    │    │
//  │  │  Structure   │  │  Definition  │  │  & Runtime   │    │
//  │  └──────────────┘  └──────────────┘  └──────────────┘    │
//  │                                                             │
//  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │
//  │  │   Resolver   │  │   Compiler   │  │  Validator   │    │
//  │  │  & Cache     │  │     CSS      │  │  & Testing   │    │
//  │  └──────────────┘  └──────────────┘  └──────────────┘    │
//  │                                                             │
//  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐    │
//  │  │    Tenant    │  │   Storage    │  │   Observers  │    │
//  │  │  Management  │  │  Interface   │  │  & Events    │    │
//  │  └──────────────┘  └──────────────┘  └──────────────┘    │
//  └─────────────────────────────────────────────────────────────┘
//                              ▼
//  ┌─────────────────────────────────────────────────────────────┐
//  │                  External Dependencies                      │
//  │  (logger, condition evaluator, ristretto cache)            │
//  └─────────────────────────────────────────────────────────────┘
//
// # Core Concepts
//
// ## Design Tokens
//
// Design tokens are the atomic visual design decisions that define a theme.
// The package uses a three-tier token hierarchy:
//
//  1. Primitives: Raw values (colors: "#3b82f6", spacing: "1rem")
//  2. Semantic: Contextual assignments (background: "primitives.colors.white")
//  3. Components: Component-specific styles (button.primary.background)
//
// ## Token Resolution
//
// Token references are resolved recursively through the hierarchy:
//
//  components.button.primary.background
//    → semantic.colors.primary
//      → primitives.colors.blue-600
//        → "#3b82f6"
//
// Resolution includes:
//   - Circular reference detection
//   - Multi-level caching
//   - Context-aware evaluation (tenant, dark mode)
//   - Validation and error reporting
//
// ## Theme Management
//
// The ThemeManager provides centralized theme lifecycle management:
//   - Theme registration and storage
//   - Runtime switching and compilation
//   - Multi-tenant isolation
//   - Conditional overrides
//   - Performance optimization
//
// ## Multi-Tenancy
//
// Complete tenant isolation with per-tenant:
//   - Theme selection and customization
//   - Branding overrides
//   - Feature flags
//   - Storage and caching
//   - Access control
//
// # Usage Examples
//
// ## Basic Theme Management
//
//  // Create manager
//  config := theme.DefaultManagerConfig()
//  manager := theme.NewManager(config)
//
//  // Register theme
//  myTheme := &theme.Theme{
//      ID:   "corporate",
//      Name: "Corporate Theme",
//      Tokens: theme.GetDefaultTokens(),
//  }
//  err := manager.RegisterTheme(ctx, myTheme)
//
//  // Get and compile theme
//  compiled, err := manager.GetTheme(ctx, "corporate", nil)
//  css := compiled.CSS
//
// ## Token Resolution
//
//  // Create resolver
//  resolver := theme.NewResolver(manager, cache)
//
//  // Resolve token
//  value, err := resolver.Resolve(ctx, "semantic.colors.primary")
//  // Returns: "#3b82f6"
//
// ## Multi-Tenant Usage
//
//  // Create tenant context
//  ctx := theme.WithTenant(context.Background(), "acme-corp")
//
//  // Configure tenant
//  tenantConfig := &theme.TenantConfig{
//      DefaultTheme: "corporate",
//      Branding: &theme.BrandingOverrides{
//          PrimaryColor: "#ff6b35",
//      },
//  }
//  err := manager.ConfigureTenant(ctx, tenantConfig)
//
//  // Get tenant theme
//  compiled, err := manager.GetTheme(ctx, "corporate", nil)
//
// ## Conditional Theming
//
//  // Add conditional override
//  condition := &theme.Condition{
//      ID:         "dark-hours",
//      Expression: "time.hour >= 18 || time.hour < 6",
//      Priority:   100,
//      Overrides: map[string]string{
//          "semantic.colors.background": "primitives.colors.gray-900",
//      },
//  }
//  theme.Conditions = append(theme.Conditions, condition)
//
//  // Evaluate with context
//  evalData := map[string]any{
//      "time": map[string]any{"hour": 20},
//  }
//  compiled, err := manager.GetTheme(ctx, "corporate", evalData)
//
// ## Validation
//
//  // Validate theme
//  validator := theme.NewValidator()
//  result, err := validator.Validate(myTheme)
//  if !result.Valid {
//      for _, issue := range result.Issues {
//          log.Printf("Issue: %s", issue.Message)
//      }
//  }
//
// # Performance Considerations
//
// The package is optimized for production use:
//
//   - Multi-level caching (token resolution, compiled themes, tenant configs)
//   - Lazy loading and preloading strategies
//   - Efficient memory usage with bounded caches
//   - Concurrent-safe operations with fine-grained locking
//   - CSS minification and optimization
//   - Token resolution memoization
//
// # Thread Safety
//
// All components are thread-safe and can be used concurrently:
//   - Manager supports concurrent theme operations
//   - Resolver handles parallel token resolution
//   - Compiler enables concurrent compilation
//   - Caches use proper synchronization
//
// # Extensibility
//
// The package provides extension points:
//   - Custom storage backends (Storage interface)
//   - Custom validators (ValidatorFunc)
//   - Theme observers (Observer interface)
//   - Condition evaluators (pluggable)
//   - Custom token tiers (extensible structure)
//
// # Best Practices
//
//  1. Always validate themes before deployment
//  2. Use caching in production environments
//  3. Implement proper tenant isolation
//  4. Monitor cache hit ratios
//  5. Preload frequently used themes
//  6. Use semantic tokens over primitives in components
//  7. Leverage conditional theming for dynamic UX
//  8. Keep token hierarchies shallow (max 3-4 levels)
//  9. Version themes for controlled updates
//  10. Test themes across all supported tenants
//
// # Security
//
// The package includes security measures:
//   - Input validation for all theme data
//   - Sanitization of custom CSS/JS
//   - Tenant isolation enforcement
//   - Expression evaluation sandboxing
//   - Size limits for uploaded themes
//   - Access control integration points
//
// # Monitoring
//
// Built-in observability:
//   - Cache statistics and hit rates
//   - Compilation metrics
//   - Resolution performance
//   - Validation results
//   - Tenant usage patterns
//
// Package version: 1.0.0
package theme
}
}
}
}
}
}
}
}
