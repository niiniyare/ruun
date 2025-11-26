# Design Tokens

**Schema-driven design system tokens for consistent component styling**

## Overview

Design tokens provide the foundation for all component styling in our schema-driven system. They create a single source of truth that flows from the backend schema definitions to frontend components.

## Token Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    DESIGN TOKEN FLOW                       │
├─────────────────────────────────────────────────────────────┤
│ Schema Theme → Design Tokens → CSS Variables → Components  │
│                                                             │
│ 1. Backend Theme Definition                                 │
│    └── Go structs with semantic naming                     │
│                                                             │
│ 2. Token Generation                                         │
│    └── Automatic CSS custom property generation            │
│                                                             │
│ 3. Component Consumption                                    │
│    └── Components use var() references                     │
└─────────────────────────────────────────────────────────────┘
```

## Schema Token System

### Theme Definition in Go

```go
// Backend theme definition with comprehensive token system
theme := &schema.Theme{
    ID:   "corporate-design-system",
    Name: "Corporate Design System",
    Tokens: &schema.DesignTokens{
        // Semantic color system
        Colors: map[string]string{
            // Primary palette
            "primary.50":   "#eff6ff",
            "primary.100":  "#dbeafe", 
            "primary.500":  "#3b82f6",
            "primary.900":  "#1e3a8a",
            
            // Semantic colors
            "background":     "var(--white)",
            "foreground":     "var(--gray-950)",
            "muted":          "var(--gray-100)",
            "muted-foreground": "var(--gray-500)",
            
            // Component-specific semantics
            "card":           "var(--white)",
            "card-foreground": "var(--gray-950)",
            "input":          "var(--gray-200)",
            "ring":           "var(--primary-500)",
            
            // State colors
            "destructive":    "var(--error-500)",
            "warning":        "var(--warning-500)",
            "success":        "var(--success-500)",
        },
        
        // Typography system
        Typography: map[string]string{
            "font-family-sans": "Inter, system-ui, sans-serif",
            "font-family-mono": "JetBrains Mono, monospace",
            
            "font-size-xs":   "0.75rem",
            "font-size-sm":   "0.875rem", 
            "font-size-base": "1rem",
            "font-size-lg":   "1.125rem",
            "font-size-xl":   "1.25rem",
            
            "line-height-tight":  "1.25",
            "line-height-normal": "1.5",
            "line-height-relaxed": "1.625",
        },
        
        // Spacing system (based on 0.25rem increments)
        Spacing: map[string]string{
            "space-0":  "0",
            "space-px": "1px",
            "space-1":  "0.25rem",   // 4px
            "space-2":  "0.5rem",    // 8px
            "space-3":  "0.75rem",   // 12px
            "space-4":  "1rem",      // 16px
            "space-6":  "1.5rem",    // 24px
            "space-8":  "2rem",      // 32px
            "space-10": "2.5rem",    // 40px
            "space-12": "3rem",      // 48px
        },
        
        // Border radius system
        BorderRadius: map[string]string{
            "radius-none": "0",
            "radius-sm":   "0.125rem", // 2px
            "radius-base": "0.25rem",  // 4px
            "radius-md":   "0.375rem", // 6px
            "radius-lg":   "0.5rem",   // 8px
            "radius-xl":   "0.75rem",  // 12px
            "radius-full": "9999px",
        },
        
        // Shadow system
        Shadows: map[string]string{
            "shadow-sm":   "0 1px 2px 0 rgb(0 0 0 / 0.05)",
            "shadow-base": "0 1px 3px 0 rgb(0 0 0 / 0.1)",
            "shadow-md":   "0 4px 6px -1px rgb(0 0 0 / 0.1)",
            "shadow-lg":   "0 10px 15px -3px rgb(0 0 0 / 0.1)",
            "shadow-xl":   "0 20px 25px -5px rgb(0 0 0 / 0.1)",
        },
    },
}
```

### Generated CSS Custom Properties

The schema system automatically generates CSS custom properties:

```css
/* Auto-generated from schema theme definition */
:root {
  /* === SEMANTIC COLOR SYSTEM === */
  
  /* Background colors */
  --background: var(--white);
  --foreground: var(--gray-950);
  --card: var(--white);
  --card-foreground: var(--gray-950);
  --muted: var(--gray-100);
  --muted-foreground: var(--gray-500);
  
  /* Interactive colors */
  --primary: var(--primary-900);
  --primary-foreground: var(--primary-50);
  --secondary: var(--gray-100);
  --secondary-foreground: var(--gray-900);
  
  /* State colors */
  --destructive: var(--error-500);
  --destructive-foreground: var(--error-50);
  --warning: var(--warning-500);
  --success: var(--success-500);
  
  /* Form elements */
  --input: var(--gray-200);
  --input-foreground: var(--gray-950);
  --ring: var(--primary-500);
  --border: var(--gray-200);
  
  /* === TYPOGRAPHY SYSTEM === */
  --font-sans: Inter, system-ui, sans-serif;
  --font-mono: JetBrains Mono, monospace;
  
  --font-size-xs: 0.75rem;
  --font-size-sm: 0.875rem;
  --font-size-base: 1rem;
  --font-size-lg: 1.125rem;
  
  /* === SPACING SYSTEM === */
  --space-1: 0.25rem;
  --space-2: 0.5rem;
  --space-3: 0.75rem;
  --space-4: 1rem;
  --space-6: 1.5rem;
  --space-8: 2rem;
  
  /* === BORDER RADIUS === */
  --radius-sm: 0.125rem;
  --radius-md: 0.375rem;
  --radius-lg: 0.5rem;
  
  /* === SHADOWS === */
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1);
}

/* === DARK MODE TOKENS === */
[data-theme="dark"] {
  --background: var(--gray-950);
  --foreground: var(--gray-50);
  --card: var(--gray-900);
  --card-foreground: var(--gray-50);
  --muted: var(--gray-800);
  --muted-foreground: var(--gray-400);
  --primary: var(--primary-400);
  --border: var(--gray-800);
  --input: var(--gray-800);
}
```

## Component Token Usage

### ✅ Correct Token Usage

```go
// Button component using schema design tokens
func buttonClasses(variant ButtonVariant) string {
    baseClasses := []string{
        // Layout and structure
        "inline-flex",
        "items-center", 
        "justify-center",
        "whitespace-nowrap",
        
        // Design tokens for consistent styling
        "rounded-[var(--radius-md)]",
        "text-[var(--font-size-sm)]",
        "font-medium",
        "transition-colors",
        
        // Accessibility and interaction
        "focus-visible:outline-none",
        "focus-visible:ring-2",
        "focus-visible:ring-[var(--ring)]",
        "disabled:pointer-events-none",
        "disabled:opacity-50",
    }
    
    // Variant-specific token usage
    switch variant {
    case ButtonPrimary:
        baseClasses = append(baseClasses,
            "bg-[var(--primary)]",
            "text-[var(--primary-foreground)]",
            "hover:bg-[var(--primary)]/90",
        )
    case ButtonSecondary:
        baseClasses = append(baseClasses,
            "bg-[var(--secondary)]",
            "text-[var(--secondary-foreground)]", 
            "hover:bg-[var(--secondary)]/80",
        )
    case ButtonDestructive:
        baseClasses = append(baseClasses,
            "bg-[var(--destructive)]",
            "text-[var(--destructive-foreground)]",
            "hover:bg-[var(--destructive)]/90",
        )
    }
    
    return strings.Join(baseClasses, " ")
}
```

### ❌ Incorrect Hardcoded Classes

```go
// WRONG: Hardcoded Tailwind classes
func buttonClassesWrong(variant ButtonVariant) string {
    baseClasses := []string{
        "inline-flex",
        "items-center",
        "rounded-md",     // Should be "rounded-[var(--radius-md)]"
        "text-sm",        // Should be "text-[var(--font-size-sm)]"
        "bg-blue-500",    // Should be "bg-[var(--primary)]"
        "text-white",     // Should be "text-[var(--primary-foreground)]"
    }
    
    return strings.Join(baseClasses, " ")
}
```

## Token Categories

### 1. Semantic Tokens

**Purpose**: Provide meaning-based token names that adapt to different themes and contexts.

```css
/* Semantic tokens for consistent meaning */
--primary: var(--blue-600);
--secondary: var(--gray-100);
--destructive: var(--red-500);
--warning: var(--yellow-500);
--success: var(--green-500);

--background: var(--white);
--foreground: var(--gray-950);
--muted: var(--gray-100);
--border: var(--gray-200);
```

**Usage in Components**:
```go
// Always use semantic tokens for component styling
"bg-[var(--primary)]"           // ✅ Good - semantic meaning
"text-[var(--foreground)]"      // ✅ Good - adapts to theme
"border-[var(--border)]"        // ✅ Good - consistent borders
```

### 2. Component Tokens

**Purpose**: Specific tokens for individual component styling needs.

```css
/* Component-specific design tokens */
--component-button-height-sm: 2rem;
--component-button-height-md: 2.5rem;
--component-button-height-lg: 3rem;

--component-input-height: 2.5rem;
--component-input-padding: 0.75rem;

--component-card-padding: 1.5rem;
--component-card-gap: 1rem;
```

**Usage in Components**:
```go
// Size variants using component tokens
switch size {
case ButtonSizeSM:
    classes = append(classes, "h-[var(--component-button-height-sm)]")
case ButtonSizeMD:
    classes = append(classes, "h-[var(--component-button-height-md)]")
case ButtonSizeLG:
    classes = append(classes, "h-[var(--component-button-height-lg)]")
}
```

### 3. Primitive Tokens

**Purpose**: Raw color, spacing, and typography values that form the foundation.

```css
/* Primitive color palette */
--white: 0 0% 100%;
--gray-50: 210 40% 98%;
--gray-100: 210 40% 96%;
--gray-500: 215 16% 47%;
--gray-900: 222 84% 5%;

--blue-500: 217 91% 60%;
--blue-600: 221 83% 53%;

/* Primitive spacing scale */
--space-1: 0.25rem;
--space-2: 0.5rem;
--space-3: 0.75rem;
--space-4: 1rem;
```

**Usage**: Primitives are referenced by semantic tokens, not directly by components.

## Multi-Tenant Token Support

### Tenant-Specific Themes

```go
// Tenant-specific theme customization
func applyTenantTheme(schema *schema.Schema, tenantID string) {
    tenantManager := schema.NewTenantManager()
    
    // Get tenant-specific theme overrides
    tenantTheme := tenantManager.GetTenantTheme(tenantID)
    
    // Apply tenant customizations
    if tenantTheme != nil {
        schema.ApplyTheme(ctx, tenantTheme.ID)
        
        // Inject tenant-specific CSS variables
        tenantTokens := tenantTheme.GenerateTokenOverrides()
        schema.InjectCustomTokens(tenantTokens)
    }
}
```

### Dynamic Token Injection

```go
// Dynamic theme token generation for multi-tenant support
templ TenantThemeTokens() {
    if theme := getTenantTheme(ctx); theme != nil {
        <style>
            :root {
                for token, value := range theme.GetTokenOverrides() {
                    --{ token }: { value };
                }
                
                /* Tenant branding */
                --tenant-primary: { theme.BrandingColors.Primary };
                --tenant-logo-url: url("{ theme.BrandingAssets.Logo }");
            }
        </style>
    }
}
```

## Dark Mode Token System

```css
/* Automatic dark mode token switching */
[data-theme="dark"] {
  /* Invert semantic meanings for dark mode */
  --background: var(--gray-950);
  --foreground: var(--gray-50);
  --card: var(--gray-900);
  --card-foreground: var(--gray-50);
  --muted: var(--gray-800);
  --muted-foreground: var(--gray-400);
  
  /* Adjust interactive colors for dark mode */
  --primary: var(--blue-400);
  --border: var(--gray-800);
  --input: var(--gray-800);
  --ring: var(--blue-400);
}
```

## Token Migration Checklist

### From Hardcoded to Schema Tokens

- [ ] **Replace color classes**: `bg-blue-500` → `bg-[var(--primary)]`
- [ ] **Replace spacing**: `p-4` → `p-[var(--space-4)]`
- [ ] **Replace typography**: `text-sm` → `text-[var(--font-size-sm)]`
- [ ] **Replace borders**: `rounded-md` → `rounded-[var(--radius-md)]`
- [ ] **Replace shadows**: `shadow-md` → `shadow-[var(--shadow-md)]`

### Validation Tools

```bash
# Check for hardcoded classes in components
grep -r "bg-blue\|text-red\|p-[0-9]\|m-[0-9]" views/components/

# Verify token usage
grep -r "var(--" views/components/ | wc -l

# Schema token validation
go run ./cmd/validate-tokens --check-usage
```

## Performance Considerations

### Token Resolution Performance

```go
// Efficient token resolution with caching
type TokenResolver struct {
    cache map[string]string
    mu    sync.RWMutex
}

func (r *TokenResolver) ResolveToken(token string, theme *schema.Theme) string {
    r.mu.RLock()
    cached, exists := r.cache[token]
    r.mu.RUnlock()
    
    if exists {
        return cached
    }
    
    // Resolve token through theme hierarchy
    resolved := theme.ResolveToken(token)
    
    r.mu.Lock()
    r.cache[token] = resolved
    r.mu.Unlock()
    
    return resolved
}
```

### CSS Generation Optimization

```go
// Optimize CSS generation for production
func generateOptimizedCSS(theme *schema.Theme) string {
    var css strings.Builder
    
    // Only generate CSS for tokens actually used by components
    usedTokens := analyzeComponentTokenUsage()
    
    css.WriteString(":root {\n")
    for _, token := range usedTokens {
        if value := theme.Tokens.Get(token); value != "" {
            css.WriteString(fmt.Sprintf("  --%s: %s;\n", token, value))
        }
    }
    css.WriteString("}\n")
    
    return css.String()
}
```

## Related Documentation

- **[Schema Theme System](../schema/07-layout.md)** - Backend theme management
- **[Design System Styles](../styles/)** - Complete design system documentation
- **[Component Integration](./07-integration.md)** - Schema-component integration patterns

---

Design tokens provide the foundation for consistent, maintainable, and theme-able component systems. By following the schema-driven token approach, components automatically inherit proper styling that adapts to different themes, tenants, and accessibility requirements.