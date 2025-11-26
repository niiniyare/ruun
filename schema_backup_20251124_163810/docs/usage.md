# Schema Package Usage Guide

**Comprehensive guide to building enterprise UIs with the schema-driven system**

## Table of Contents

1. [Package Overview](#package-overview)
2. [Architecture Components](#architecture-components)
3. [Quick Start](#quick-start)
4. [Core Usage Patterns](#core-usage-patterns)
5. [Advanced Integration](#advanced-integration)
6. [Performance Optimization](#performance-optimization)
7. [Production Best Practices](#production-best-practices)
8. [Troubleshooting](#troubleshooting)

---

## Package Overview

### What is the Schema Package?

The schema package is a comprehensive **JSON-driven UI system** that enables backend developers to define complete user interfaces without writing frontend code. It provides:

- **Thread-safe schema definitions** with comprehensive validation
- **Multi-level caching architecture** (Memory → Redis → Storage)
- **Business rules engine** for dynamic behavior
- **Theme system** with design tokens
- **Runtime state management** with dependency injection
- **Component integration** with atomic design principles

### Key Benefits

✅ **Backend-driven UI**: Define forms and components in JSON  
✅ **Type Safety**: Full Go type checking and validation  
✅ **Performance**: Multi-level caching reduces load times from 50ms to 1ms  
✅ **Enterprise Features**: Multi-tenancy, workflows, permissions, i18n  
✅ **Frontend Agnostic**: Works with HTMX, Alpine.js, React, Vue, etc.  

---

## Architecture Components

### Core Types

```go
// Main schema structure - thread-safe and fully configurable
type Schema struct {
    // Identity
    ID          string `json:"id"`
    Type        Type   `json:"type"`
    Title       string `json:"title"`
    
    // Core UI structure
    Fields      []Field  `json:"fields"`
    Actions     []Action `json:"actions"`
    Layout      *Layout  `json:"layout"`
    
    // Enterprise features
    Security    *Security   `json:"security"`
    Workflow    *Workflow   `json:"workflow"`
    Validation  *Validation `json:"validation"`
    
    // Framework integration
    HTMX        *HTMX   `json:"htmx"`
    Alpine      *Alpine `json:"alpine"`
}

// Registry with multi-level caching
type Registry struct {
    storage     Storage              // Backend storage
    cache       Storage              // Redis cache (optional)
    memoryCache map[string]*cacheEntry // L1 cache
}

// Runtime with builder pattern
type Runtime struct {
    schema      *Schema
    state       *RuntimeState
    renderer    RuntimeRenderer      // Dependency injection
    validator   RuntimeValidator
    conditional RuntimeConditionalEngine
}
```

### Performance Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    SCHEMA PACKAGE ARCHITECTURE                 │
├─────────────────────────────────────────────────────────────────┤
│  Request → Registry (Multi-level Cache) → Runtime → Components │
│                                                                 │
│  L1: Memory Cache (1ms)                                        │
│  ├── Thread-safe concurrent access                             │
│  └── Configurable TTL (default: 5 minutes)                     │
│                                                                 │
│  L2: Redis Cache (5ms)                                         │
│  ├── Distributed caching for scalability                       │
│  └── Configurable TTL (default: 1 hour)                        │
│                                                                 │
│  L3: Storage Backend (50ms)                                    │
│  ├── File system, PostgreSQL, S3, etc.                        │
│  └── Source of truth for schema definitions                    │
└─────────────────────────────────────────────────────────────────┘
```

---

## Quick Start

### 1. Basic Schema Creation

```go
package main

import (
    "context"
    "log"
    "github.com/niiniyare/ruun/pkg/schema"
)

func main() {
    // Create a simple form schema
    userForm := &schema.Schema{
        ID:    "user-registration",
        Type:  schema.TypeForm,
        Title: "User Registration",
        Config: &schema.Config{
            Action:   "/api/users",
            Method:   "POST",
            Target:   "#main-content",
            Encoding: "application/json",
        },
        Fields: []schema.Field{
            {
                Name:        "first_name",
                Type:        schema.FieldTypeText,
                Label:       "First Name",
                Placeholder: "Enter your first name",
                Required:    true,
                Validation: &schema.FieldValidation{
                    MinLength: 2,
                    MaxLength: 50,
                },
            },
            {
                Name:        "email",
                Type:        schema.FieldTypeEmail,
                Label:       "Email Address",
                Placeholder: "user@example.com",
                Required:    true,
            },
            {
                Name: "role",
                Type: schema.FieldTypeSelect,
                Label: "User Role",
                Options: []schema.FieldOption{
                    {Value: "user", Label: "Regular User"},
                    {Value: "admin", Label: "Administrator"},
                },
                Default: "user",
            },
        },
        Actions: []schema.Action{
            {
                ID:      "submit",
                Type:    schema.ActionSubmit,
                Text:    "Create User",
                Variant: "primary",
            },
            {
                ID:      "cancel",
                Type:    schema.ActionButton,
                Text:    "Cancel",
                Variant: "secondary",
                HTMX: &schema.HTMXConfig{
                    Get:    "/users",
                    Target: "#main-content",
                },
            },
        },
    }
    
    log.Printf("Created schema: %+v", userForm)
}
```

### 2. Registry Setup with Caching

```go
package main

import (
    "context"
    "time"
    "github.com/niiniyare/ruun/pkg/schema"
)

func setupRegistry() *schema.Registry {
    // Configure storage backend
    storage := schema.NewFilesystemStorage("./schemas")
    
    // Optional: Add Redis cache for distributed systems
    redisCache := schema.NewRedisStorage("redis://localhost:6379")
    
    // Create registry with multi-level caching
    registry := schema.NewRegistry(schema.RegistryConfig{
        Storage:     storage,
        Cache:       redisCache,
        MemoryTTL:   5 * time.Minute,
        CacheTTL:    1 * time.Hour,
        EnableCache: true,
    })
    
    return registry
}

func main() {
    ctx := context.Background()
    registry := setupRegistry()
    
    // Store schema
    userForm := createUserFormSchema()
    err := registry.Set(ctx, userForm)
    if err != nil {
        log.Fatal(err)
    }
    
    // Retrieve schema (uses caching automatically)
    retrieved, err := registry.Get(ctx, "user-registration")
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Retrieved schema: %s", retrieved.Title)
}
```

### 3. Runtime Integration

```go
package main

import (
    "context"
    "github.com/niiniyare/ruun/pkg/schema"
)

func createRuntime(enrichedSchema *schema.Schema) *schema.Runtime {
    // Use builder pattern for runtime creation
    runtime, err := schema.NewRuntimeBuilder(enrichedSchema).
        WithRenderer(&CustomRenderer{}).
        WithValidator(&CustomValidator{}).
        WithConditionalEngine(&CustomConditionalEngine{}).
        WithLocale("en-US").
        WithInitialData(map[string]any{
            "user_type": "customer",
            "country":   "US",
        }).
        WithEventHandler(schema.EventTypeSubmit, func(ctx context.Context, event schema.Event) error {
            log.Printf("Form submitted: %+v", event.Data)
            return nil
        }).
        Build()
        
    if err != nil {
        log.Fatal(err)
    }
    
    return runtime
}

// Custom implementations (dependency injection)
type CustomRenderer struct{}

func (r *CustomRenderer) RenderField(ctx context.Context, field *schema.Field) (string, error) {
    // Implement field rendering logic
    return fmt.Sprintf(`<input name="%s" type="%s" />`, field.Name, field.Type), nil
}

type CustomValidator struct{}

func (v *CustomValidator) ValidateField(ctx context.Context, field *schema.Field, value any) []string {
    // Implement validation logic
    return nil
}

type CustomConditionalEngine struct{}

func (c *CustomConditionalEngine) EvaluateCondition(ctx context.Context, condition string, data map[string]any) (bool, error) {
    // Implement conditional logic evaluation
    return true, nil
}
```

---

## Core Usage Patterns

### Builder Pattern for Schema Creation

```go
// Fluent schema building
schema := schema.NewBuilder("product-form", schema.TypeForm, "Create Product").
    WithConfig("/api/products", "POST").
    AddTextField("name", "Product Name", true).
    AddTextField("description", "Description", false).
    AddNumberField("price", "Price", true).
    AddSelectField("category", "Category", []schema.FieldOption{
        {Value: "electronics", Label: "Electronics"},
        {Value: "clothing", Label: "Clothing"},
        {Value: "books", Label: "Books"},
    }, true).
    WithValidation(&schema.Validation{
        Rules: []schema.ValidationRule{
            {
                Field:   "price",
                Type:    "min_value",
                Value:   0.01,
                Message: "Price must be greater than 0",
            },
        },
    }).
    WithSecurity(&schema.Security{
        CSRFProtection: true,
        RateLimit: &schema.RateLimit{
            Requests: 5,
            Window:   "1m",
        },
    }).
    AddAction("submit", "Create Product", schema.ActionSubmit, "primary").
    AddAction("draft", "Save Draft", schema.ActionButton, "secondary").
    Build(ctx)
```

### Business Rules Integration

```go
// Define business rules for dynamic behavior
func setupBusinessRules(schema *schema.Schema) {
    // Conditional field visibility
    visibilityRule := &schema.BusinessRule{
        ID:          "international-shipping",
        Name:        "International Shipping Fields",
        Type:        schema.RuleTypeFieldVisibility,
        Priority:    10,
        Enabled:     true,
        Condition:   condition.NewCondition("country", condition.OpNotEqual, "US"),
        Actions: []schema.BusinessRuleAction{
            {
                Type:   schema.ActionShowField,
                Target: "customs_info",
                Config: map[string]any{
                    "required": true,
                    "help":     "Required for international shipments",
                },
            },
        },
    }
    
    // Dynamic validation rules
    validationRule := &schema.BusinessRule{
        ID:       "premium-validation",
        Name:     "Premium Customer Validation",
        Type:     schema.RuleTypeFieldValidation,
        Priority: 5,
        Enabled:  true,
        Condition: condition.And(
            condition.NewCondition("customer_type", condition.OpEqual, "premium"),
            condition.NewCondition("order_total", condition.OpGreaterThan, 1000),
        ),
        Actions: []schema.BusinessRuleAction{
            {
                Type:   schema.ActionRequireField,
                Target: "approval_code",
                Config: map[string]any{
                    "message": "Approval code required for premium orders over $1000",
                },
            },
        },
    }
    
    // Apply rules to schema
    schema.AddBusinessRule(visibilityRule)
    schema.AddBusinessRule(validationRule)
}
```

### Theme and Design Token Integration

```go
// Create theme with design tokens
func createCorporateTheme() *schema.Theme {
    return &schema.Theme{
        ID:          "corporate-v2",
        Name:        "Corporate Theme v2",
        Description: "Updated corporate branding with accessibility improvements",
        Version:     "2.1.0",
        Tokens: &schema.DesignTokens{
            Colors: map[string]string{
                "primary.50":   "#eff6ff",
                "primary.500":  "#3b82f6",
                "primary.900":  "#1e3a8a",
                "gray.50":      "#f9fafb",
                "gray.500":     "#6b7280",
                "gray.900":     "#111827",
            },
            Typography: map[string]string{
                "font.family.sans": "Inter, system-ui, sans-serif",
                "font.size.sm":     "0.875rem",
                "font.size.base":   "1rem",
                "font.size.lg":     "1.125rem",
            },
            Spacing: map[string]string{
                "space.1": "0.25rem",
                "space.4": "1rem",
                "space.6": "1.5rem",
                "space.8": "2rem",
            },
            BorderRadius: map[string]string{
                "radius.sm": "0.125rem",
                "radius.md": "0.375rem",
                "radius.lg": "0.5rem",
            },
        },
        DarkMode: &schema.DarkModeConfig{
            Enabled:      true,
            AutoDetect:   true,
            ColorOverrides: map[string]string{
                "background": "#1f2937",
                "foreground": "#f9fafb",
            },
        },
        Accessibility: &schema.AccessibilityConfig{
            HighContrast:    true,
            ReducedMotion:   true,
            FocusIndicators: true,
        },
    }
}

// Apply theme to schema
func applyTheme(schema *schema.Schema, theme *schema.Theme) {
    // Theme manager handles token resolution and application
    themeManager := schema.NewThemeManager()
    themeManager.RegisterTheme(theme)
    
    // Apply theme to schema
    schema.ApplyTheme(ctx, theme.ID)
}
```

### Multi-Tenant Usage

```go
// Multi-tenant schema management
func setupMultiTenantSchemas() {
    tenantManager := schema.NewTenantManager()
    
    // Register tenants with specific configurations
    tenantManager.RegisterTenant("company-a", &schema.TenantConfig{
        ThemeID:        "corporate-blue",
        DefaultLocale:  "en-US",
        Features: []string{"advanced-validation", "workflow", "analytics"},
        Branding: &schema.TenantBranding{
            Logo:       "/assets/company-a-logo.png",
            PrimaryColor: "#1e40af",
        },
    })
    
    tenantManager.RegisterTenant("company-b", &schema.TenantConfig{
        ThemeID:       "corporate-green", 
        DefaultLocale: "es-ES",
        Features: []string{"basic-validation"},
        Branding: &schema.TenantBranding{
            Logo:        "/assets/company-b-logo.png",
            PrimaryColor: "#059669",
        },
    })
    
    // Get tenant-specific schema
    tenantSchema := schema.ForTenant(ctx, "company-a")
    
    // Apply tenant-specific customizations
    tenantSchema.ApplyTenantBranding("company-a")
    tenantSchema.ApplyTenantBusinessRules("company-a")
}
```

---

## Advanced Integration

### Fiber HTTP Handler Integration

```go
package handlers

import (
    "context"
    "github.com/gofiber/fiber/v2"
    "github.com/niiniyare/ruun/pkg/schema"
    "github.com/niiniyare/ruun/views/components/organisms"
)

// Schema-driven form handler
func SchemaFormHandler(registry *schema.Registry) fiber.Handler {
    return func(c *fiber.Ctx) error {
        schemaID := c.Params("schemaID")
        
        // Get schema from registry (cached)
        schema, err := registry.Get(c.Context(), schemaID)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{
                "error": "Schema not found",
            })
        }
        
        // Enrich schema with user context
        enrichedSchema := schema.WithContext(c.Context()).
            WithUser(getCurrentUser(c)).
            WithTenant(getCurrentTenant(c))
        
        // Apply business rules
        enrichedSchema.ApplyBusinessRules(c.Context())
        
        // Create runtime
        runtime := schema.NewRuntimeBuilder(enrichedSchema).
            WithRenderer(&TemplRenderer{}).
            WithValidator(&FiberValidator{}).
            Build()
        
        // Render form component
        return organisms.SchemaForm(organisms.SchemaFormProps{
            Schema:  enrichedSchema,
            Runtime: runtime,
            Action:  schema.Config.Action,
        }).Render(c.Context(), c.Response().BodyWriter())
    }
}

// Schema validation handler
func SchemaValidationHandler(registry *schema.Registry) fiber.Handler {
    return func(c *fiber.Ctx) error {
        schemaID := c.Params("schemaID")
        fieldName := c.Params("fieldName")
        
        schema, err := registry.Get(c.Context(), schemaID)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Schema not found"})
        }
        
        // Get field value from request
        value := c.FormValue(fieldName)
        
        // Validate field using schema rules
        errors := schema.ValidateField(c.Context(), fieldName, value)
        
        if len(errors) > 0 {
            return c.JSON(fiber.Map{
                "field":  fieldName,
                "errors": errors,
                "valid":  false,
            })
        }
        
        return c.JSON(fiber.Map{
            "field": fieldName,
            "valid": true,
        })
    }
}
```

### Component Integration with HTMX and Alpine.js

```go
// Templ component with schema integration
templ SchemaFormField(field *schema.Field, runtime *schema.Runtime) {
    <div 
        class="form-field"
        x-data="{
            value: @js(field.Value),
            errors: @js(runtime.GetFieldErrors(field.Name)),
            touched: false,
            valid: true
        }"
        x-show={ field.Runtime.Visible ? "true" : "false" }
        x-transition
    >
        <label 
            for={ field.Name }
            class="text-[var(--foreground)] font-medium"
        >
            { field.Label }
            if field.Required {
                <span class="text-[var(--destructive)]">*</span>
            }
        </label>
        
        switch field.Type {
        case schema.FieldTypeText, schema.FieldTypeEmail:
            <input
                type={ string(field.Type) }
                name={ field.Name }
                id={ field.Name }
                class="
                    w-full p-[var(--space-3)] 
                    border-[var(--border)] rounded-[var(--radius-md)]
                    bg-[var(--input)] text-[var(--input-foreground)]
                    focus:ring-2 focus:ring-[var(--ring)]
                "
                placeholder={ field.Placeholder }
                required?={ field.Required }
                disabled?={ field.Disabled }
                readonly?={ field.Readonly }
                value={ fmt.Sprintf("%v", field.Value) }
                x-model="value"
                x-on:blur="touched = true"
                hx-post={ "/api/validate/" + field.Name }
                hx-target={ "#" + field.Name + "-errors" }
                hx-trigger="blur changed delay:300ms"
            />
            
        case schema.FieldTypeSelect:
            <select
                name={ field.Name }
                id={ field.Name }
                class="
                    w-full p-[var(--space-3)]
                    border-[var(--border)] rounded-[var(--radius-md)]
                    bg-[var(--input)] text-[var(--input-foreground)]
                "
                required?={ field.Required }
                disabled?={ field.Disabled }
                x-model="value"
            >
                if field.Placeholder != "" {
                    <option value="">{ field.Placeholder }</option>
                }
                for _, option := range field.Options {
                    <option 
                        value={ option.Value }
                        selected?={ option.Value == fmt.Sprintf("%v", field.Value) }
                    >
                        { option.Label }
                    </option>
                }
            </select>
        }
        
        // Error display
        <div 
            id={ field.Name + "-errors" }
            x-show="errors.length > 0"
            x-transition
            class="mt-[var(--space-1)] text-[var(--destructive)] text-[var(--font-size-sm)]"
        >
            <template x-for="error in errors" :key="error">
                <div x-text="error"></div>
            </template>
        </div>
        
        // Help text
        if field.Help != "" {
            <div class="mt-[var(--space-1)] text-[var(--muted-foreground)] text-[var(--font-size-sm)]">
                { field.Help }
            </div>
        }
    </div>
}
```

### Workflow Integration

```go
// Workflow-enabled schema
func createWorkflowSchema() *schema.Schema {
    return schema.NewBuilder("purchase-order", schema.TypeWorkflow, "Purchase Order").
        WithWorkflow(&schema.Workflow{
            ID:   "purchase-approval",
            Name: "Purchase Order Approval",
            Steps: []schema.WorkflowStep{
                {
                    ID:          "draft",
                    Name:        "Draft",
                    Description: "Initial draft creation",
                    Actions:     []string{"save_draft", "submit_for_review"},
                    EditableFields: []string{"*"}, // All fields editable
                },
                {
                    ID:          "review",
                    Name:        "Under Review",
                    Description: "Manager review",
                    Actions:     []string{"approve", "reject", "request_changes"},
                    EditableFields: []string{}, // No fields editable
                    RequiredApprovers: []string{"manager"},
                },
                {
                    ID:          "approved",
                    Name:        "Approved",
                    Description: "Ready for processing",
                    Actions:     []string{"process"},
                    EditableFields: []string{}, // Read-only
                    FinalStep:   false,
                },
                {
                    ID:          "processed",
                    Name:        "Processed",
                    Description: "Order completed",
                    Actions:     []string{},
                    EditableFields: []string{},
                    FinalStep:   true,
                },
            },
        }).
        AddTextField("supplier", "Supplier", true).
        AddNumberField("amount", "Amount", true).
        AddTextAreaField("justification", "Justification", true).
        Build(ctx)
}

// Workflow state management
func handleWorkflowAction(runtime *schema.Runtime, action string, userID string) error {
    workflow := runtime.GetSchema().Workflow
    currentStep := runtime.GetCurrentWorkflowStep()
    
    switch action {
    case "submit_for_review":
        if currentStep.ID != "draft" {
            return errors.New("invalid transition")
        }
        return runtime.TransitionToStep("review", userID, "Submitted for review")
        
    case "approve":
        if currentStep.ID != "review" {
            return errors.New("invalid transition")
        }
        
        // Check approver permissions
        if !runtime.UserCanApprove(userID) {
            return errors.New("insufficient permissions")
        }
        
        return runtime.TransitionToStep("approved", userID, "Approved by manager")
        
    case "reject":
        return runtime.TransitionToStep("draft", userID, "Rejected - requires changes")
    }
    
    return errors.New("unknown action")
}
```

---

## Performance Optimization

### Caching Strategy

```go
// Optimized registry configuration for production
func setupProductionRegistry() *schema.Registry {
    // File system storage with optimized read paths
    storage := schema.NewFilesystemStorage("./schemas", schema.FilesystemConfig{
        CacheManifest: true,
        ReadBufferSize: 64 * 1024, // 64KB buffer
        EnableGzip:     true,
    })
    
    // Redis cluster for distributed caching
    redisCache := schema.NewRedisStorage("redis://redis-cluster:6379", schema.RedisConfig{
        PoolSize:           10,
        MaxRetries:         3,
        MinIdleConns:      5,
        ConnMaxIdleTime:   5 * time.Minute,
        EnableCompression: true,
    })
    
    // Optimized cache configuration
    return schema.NewRegistry(schema.RegistryConfig{
        Storage:       storage,
        Cache:        redisCache,
        MemoryTTL:    10 * time.Minute,  // Longer memory cache
        CacheTTL:     4 * time.Hour,     // Longer Redis cache
        EnableCache:  true,
        // Performance optimizations
        PreloadSchemas:    true,  // Preload frequently used schemas
        EnableMetrics:     true,  // Performance monitoring
        MaxMemoryEntries: 1000,   // Limit memory usage
    })
}

// Performance monitoring
func monitorSchemaPerformance(registry *schema.Registry) {
    metrics := registry.GetMetrics()
    
    log.Printf("Schema Registry Metrics:")
    log.Printf("Memory Cache Hit Rate: %.2f%%", metrics.MemoryCacheHitRate)
    log.Printf("Redis Cache Hit Rate: %.2f%%", metrics.RedisCacheHitRate)
    log.Printf("Average Response Time: %v", metrics.AvgResponseTime)
    log.Printf("Total Requests: %d", metrics.TotalRequests)
}
```

### Schema Preloading

```go
// Preload frequently used schemas for better performance
func preloadCriticalSchemas(registry *schema.Registry) error {
    criticalSchemas := []string{
        "user-login",
        "user-registration", 
        "order-form",
        "payment-form",
        "dashboard-layout",
    }
    
    ctx := context.Background()
    for _, schemaID := range criticalSchemas {
        _, err := registry.Get(ctx, schemaID)
        if err != nil {
            log.Printf("Warning: Failed to preload schema %s: %v", schemaID, err)
            continue
        }
        log.Printf("Preloaded schema: %s", schemaID)
    }
    
    return nil
}
```

### Concurrent Schema Processing

```go
// Process multiple schemas concurrently
func processSchemaConcurrently(registry *schema.Registry, schemaIDs []string) map[string]*schema.Schema {
    results := make(map[string]*schema.Schema)
    resultsMu := sync.Mutex{}
    
    var wg sync.WaitGroup
    semaphore := make(chan struct{}, 5) // Limit concurrent goroutines
    
    for _, schemaID := range schemaIDs {
        wg.Add(1)
        go func(id string) {
            defer wg.Done()
            semaphore <- struct{}{} // Acquire
            defer func() { <-semaphore }() // Release
            
            schema, err := registry.Get(context.Background(), id)
            if err != nil {
                log.Printf("Error loading schema %s: %v", id, err)
                return
            }
            
            // Apply business rules concurrently
            schema.ApplyBusinessRules(context.Background())
            
            resultsMu.Lock()
            results[id] = schema
            resultsMu.Unlock()
        }(schemaID)
    }
    
    wg.Wait()
    return results
}
```

---

## Production Best Practices

### Security Implementation

```go
// Production-ready security configuration
func setupSecureSchema() *schema.Schema {
    return schema.NewBuilder("secure-form", schema.TypeForm, "Secure Form").
        WithSecurity(&schema.Security{
            // CSRF Protection
            CSRFProtection: true,
            CSRFTokenName: "_csrf_token",
            
            // Rate Limiting
            RateLimit: &schema.RateLimit{
                Requests:    10,
                Window:     "1m",
                KeyFunc:    "ip+user", // Rate limit by IP + User
                Enabled:    true,
            },
            
            // Input Sanitization
            Sanitization: &schema.Sanitization{
                EnableHTML:     false,
                EnableSQL:      false, 
                AllowedTags:   []string{},
                StripScripts:  true,
            },
            
            // Field Encryption
            Encryption: &schema.Encryption{
                EncryptedFields: []string{"ssn", "credit_card", "bank_account"},
                Algorithm:       "AES-256-GCM",
                KeyRotation:     true,
            },
            
            // Audit Logging
            AuditLog: &schema.AuditConfig{
                Enabled:      true,
                LogLevel:     "INFO",
                LogFields:    []string{"user_id", "action", "timestamp", "ip"},
                RetentionDays: 90,
            },
        }).
        Build(ctx)
}

// Security middleware for Fiber
func SecurityMiddleware(registry *schema.Registry) fiber.Handler {
    return func(c *fiber.Ctx) error {
        schemaID := c.Params("schemaID")
        
        schema, err := registry.Get(c.Context(), schemaID)
        if err != nil {
            return err
        }
        
        if schema.Security != nil {
            // Check rate limits
            if schema.Security.RateLimit.Enabled {
                if !checkRateLimit(c, schema.Security.RateLimit) {
                    return c.Status(429).JSON(fiber.Map{
                        "error": "Rate limit exceeded",
                    })
                }
            }
            
            // Validate CSRF token
            if schema.Security.CSRFProtection {
                if !validateCSRFToken(c, schema.Security.CSRFTokenName) {
                    return c.Status(403).JSON(fiber.Map{
                        "error": "Invalid CSRF token",
                    })
                }
            }
            
            // Log security events
            if schema.Security.AuditLog.Enabled {
                logSecurityEvent(c, schema, "form_access")
            }
        }
        
        return c.Next()
    }
}
```

### Error Handling and Monitoring

```go
// Production error handling
func setupErrorHandling(registry *schema.Registry) {
    // Global error handler
    registry.SetErrorHandler(func(ctx context.Context, err error) {
        // Log error with context
        logger := log.WithFields(log.Fields{
            "error":     err.Error(),
            "timestamp": time.Now(),
            "request_id": ctx.Value("request_id"),
            "user_id":    ctx.Value("user_id"),
        })
        
        // Different handling based on error type
        switch e := err.(type) {
        case *schema.ValidationError:
            logger.Info("Schema validation failed")
            
        case *schema.PermissionError:
            logger.Warn("Permission denied")
            
        case *schema.NotFoundError:
            logger.Debug("Schema not found")
            
        default:
            logger.Error("Unexpected schema error")
            
            // Report to error tracking service
            reportError(ctx, err)
        }
    })
    
    // Health check endpoint
    registry.EnableHealthCheck("/health/schema", func() error {
        // Check registry health
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()
        
        // Test basic operations
        _, err := registry.List(ctx)
        return err
    })
}

// Metrics collection
func setupMetricsCollection(registry *schema.Registry) {
    // Enable Prometheus metrics
    registry.EnablePrometheusMetrics()
    
    // Custom metrics
    registry.AddMetric("schema_load_duration", prometheus.HistogramOpts{
        Name: "schema_load_duration_seconds",
        Help: "Time taken to load schemas",
        Buckets: []float64{0.001, 0.005, 0.01, 0.05, 0.1, 0.5, 1.0},
    })
    
    registry.AddMetric("business_rules_execution", prometheus.CounterOpts{
        Name: "business_rules_executed_total", 
        Help: "Total number of business rules executed",
    })
}
```

### Testing Strategy

```go
// Integration test example
func TestSchemaIntegration(t *testing.T) {
    // Setup test registry with in-memory storage
    storage := schema.NewMemoryStorage()
    registry := schema.NewRegistry(schema.RegistryConfig{
        Storage:     storage,
        MemoryTTL:   1 * time.Minute,
        EnableCache: false, // Disable caching for tests
    })
    
    // Create test schema
    testSchema := &schema.Schema{
        ID:    "test-form",
        Type:  schema.TypeForm,
        Title: "Test Form",
        Fields: []schema.Field{
            {
                Name:     "email",
                Type:     schema.FieldTypeEmail,
                Label:    "Email",
                Required: true,
            },
        },
    }
    
    ctx := context.Background()
    
    // Test storage
    err := registry.Set(ctx, testSchema)
    assert.NoError(t, err)
    
    // Test retrieval
    retrieved, err := registry.Get(ctx, "test-form")
    assert.NoError(t, err)
    assert.Equal(t, testSchema.Title, retrieved.Title)
    
    // Test validation
    errors := retrieved.ValidateField(ctx, "email", "invalid-email")
    assert.NotEmpty(t, errors)
    
    errors = retrieved.ValidateField(ctx, "email", "valid@example.com")
    assert.Empty(t, errors)
}

// Performance test
func BenchmarkSchemaLoad(b *testing.B) {
    registry := setupTestRegistry()
    ctx := context.Background()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := registry.Get(ctx, "test-schema")
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

---

## Troubleshooting

### Common Issues and Solutions

#### 1. Schema Not Found Errors

```go
// Problem: Schema not found in registry
// Solution: Check schema ID and storage configuration

func debugSchemaNotFound(registry *schema.Registry, schemaID string) {
    ctx := context.Background()
    
    // Check if schema exists
    exists, err := registry.Exists(ctx, schemaID)
    if err != nil {
        log.Printf("Error checking schema existence: %v", err)
        return
    }
    
    if !exists {
        // List available schemas
        schemas, err := registry.List(ctx)
        if err != nil {
            log.Printf("Error listing schemas: %v", err)
            return
        }
        
        log.Printf("Schema '%s' not found. Available schemas: %v", schemaID, schemas)
    } else {
        log.Printf("Schema exists but failed to load. Check storage configuration.")
    }
}
```

#### 2. Performance Issues

```go
// Problem: Slow schema loading
// Solution: Enable caching and optimize storage

func optimizePerformance(registry *schema.Registry) {
    metrics := registry.GetMetrics()
    
    if metrics.AvgResponseTime > 100*time.Millisecond {
        log.Printf("Performance issue detected:")
        
        if metrics.MemoryCacheHitRate < 0.8 {
            log.Printf("Low memory cache hit rate: %.2f%%", metrics.MemoryCacheHitRate)
            log.Printf("Solution: Increase memory cache TTL or size")
        }
        
        if metrics.RedisCacheHitRate < 0.9 {
            log.Printf("Low Redis cache hit rate: %.2f%%", metrics.RedisCacheHitRate)
            log.Printf("Solution: Increase Redis cache TTL or check Redis connectivity")
        }
        
        log.Printf("Consider preloading frequently used schemas")
    }
}
```

#### 3. Business Rules Not Working

```go
// Problem: Business rules not being applied
// Solution: Debug rule evaluation

func debugBusinessRules(schema *schema.Schema, data map[string]any) {
    for _, rule := range schema.GetBusinessRules() {
        log.Printf("Evaluating rule: %s", rule.Name)
        
        if rule.Condition == nil {
            log.Printf("Rule has no condition - will always execute")
            continue
        }
        
        result, err := rule.Condition.Evaluate(data)
        if err != nil {
            log.Printf("Error evaluating rule condition: %v", err)
            continue
        }
        
        log.Printf("Rule '%s' condition result: %v", rule.Name, result)
        
        if result {
            log.Printf("Rule actions will be executed: %v", rule.Actions)
        }
    }
}
```

### Debugging Tools

```go
// Schema debug utility
func debugSchema(schema *schema.Schema) {
    log.Printf("Schema Debug Information:")
    log.Printf("ID: %s", schema.ID)
    log.Printf("Type: %s", schema.Type)
    log.Printf("Field Count: %d", len(schema.Fields))
    log.Printf("Action Count: %d", len(schema.Actions))
    log.Printf("Business Rules: %d", len(schema.GetBusinessRules()))
    
    // Validate schema structure
    if err := schema.Validate(); err != nil {
        log.Printf("Schema validation errors: %v", err)
    }
    
    // Check for common issues
    for _, field := range schema.Fields {
        if field.Name == "" {
            log.Printf("WARNING: Field with empty name found")
        }
        
        if field.Type == "" {
            log.Printf("WARNING: Field '%s' has no type", field.Name)
        }
    }
}

// Runtime debug utility  
func debugRuntime(runtime *schema.Runtime) {
    log.Printf("Runtime Debug Information:")
    log.Printf("Schema ID: %s", runtime.GetSchema().ID)
    log.Printf("Current Locale: %s", runtime.GetLocale())
    log.Printf("State Keys: %v", runtime.GetStateKeys())
    
    // Check dependency injection
    if runtime.GetRenderer() == nil {
        log.Printf("WARNING: No renderer configured")
    }
    
    if runtime.GetValidator() == nil {
        log.Printf("WARNING: No validator configured")
    }
}
```

---

## Related Documentation

- **[Schema System Documentation](./schema/)** - Comprehensive schema system docs
- **[Design System Documentation](./styles/)** - Design tokens and styling
- **[Component Integration Guide](./components/)** - UI component integration  
- **[API Reference](../../../cmd/examples/)** - Example implementations

---

This usage guide provides comprehensive coverage of the schema package implementation, from basic usage to advanced production patterns. The package's architecture supports enterprise-scale applications with performance, security, and maintainability as core principles.