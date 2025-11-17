# Atom Components Refactoring Strategy

## Executive Summary

This document outlines a comprehensive refactoring strategy for the atomic component system to create a cohesive, reusable component library that leverages Templ's full capabilities while maintaining TailwindCSS compatibility and supporting runtime customization for multi-tenant environments.

## 1. Design Token System Architecture

### 1.1 Hybrid Token Approach

The new design token system combines Tailwind's semantic classes with CSS variables for runtime customization:

```css
/* views/styles/tokens.css */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  :root {
    /* Color tokens using HSL for runtime manipulation */
    --primary: 222.2 47.4% 11.2%;
    --primary-foreground: 210 40% 98%;
    --secondary: 210 40% 96.1%;
    --secondary-foreground: 222.2 47.4% 11.2%;
    
    /* Spacing tokens */
    --spacing-xs: 0.25rem;
    --spacing-sm: 0.5rem;
    --spacing-md: 1rem;
    --spacing-lg: 1.5rem;
    --spacing-xl: 2rem;
    
    /* Radius tokens */
    --radius-sm: 0.125rem;
    --radius-md: 0.375rem;
    --radius-lg: 0.5rem;
    --radius-full: 9999px;
    
    /* Animation tokens */
    --animation-fast: 150ms;
    --animation-normal: 300ms;
    --animation-slow: 500ms;
  }
  
  /* Tenant-specific overrides */
  [data-tenant="enterprise"] {
    --primary: 217 91% 60%;
    --spacing-md: 1.25rem;
  }
}

@layer components {
  /* Semantic component classes */
  .btn {
    @apply inline-flex items-center justify-center rounded-md font-medium transition-colors;
    @apply focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2;
    @apply disabled:pointer-events-none disabled:opacity-50;
  }
  
  .btn-primary {
    background-color: hsl(var(--primary));
    color: hsl(var(--primary-foreground));
    @apply hover:opacity-90;
  }
  
  .btn-secondary {
    background-color: hsl(var(--secondary));
    color: hsl(var(--secondary-foreground));
    @apply hover:opacity-90;
  }
  
  /* Size variants using tokens */
  .btn-sm {
    height: 2rem;
    padding-left: var(--spacing-sm);
    padding-right: var(--spacing-sm);
    font-size: 0.875rem;
  }
  
  .btn-md {
    height: 2.5rem;
    padding-left: var(--spacing-md);
    padding-right: var(--spacing-md);
    font-size: 0.875rem;
  }
}
```

### 1.2 Tailwind Configuration

```javascript
// tailwind.config.js
module.exports = {
  content: ["./views/**/*.{templ,go}"],
  theme: {
    extend: {
      colors: {
        // Map CSS variables to Tailwind utilities
        primary: 'hsl(var(--primary) / <alpha-value>)',
        'primary-foreground': 'hsl(var(--primary-foreground) / <alpha-value>)',
        secondary: 'hsl(var(--secondary) / <alpha-value>)',
        'secondary-foreground': 'hsl(var(--secondary-foreground) / <alpha-value>)',
      },
      spacing: {
        'xs': 'var(--spacing-xs)',
        'sm': 'var(--spacing-sm)',
        'md': 'var(--spacing-md)',
        'lg': 'var(--spacing-lg)',
        'xl': 'var(--spacing-xl)',
      },
      borderRadius: {
        'sm': 'var(--radius-sm)',
        'md': 'var(--radius-md)',
        'lg': 'var(--radius-lg)',
        'full': 'var(--radius-full)',
      },
    },
  },
  plugins: [],
}
```

## 2. Component Composition Architecture

### 2.1 Base Component Pattern

Create a set of primitive components that serve as building blocks:

```go
// views/components/atoms/primitives.templ
package atoms

// Base button primitive with minimal styling
templ ButtonBase(attrs templ.Attributes, children ...templ.Component) {
    <button {...attrs} class="btn">
        for _, child := range children {
            @child
        }
    </button>
}

// Flexible icon component
templ Icon(name string, size string, attrs ...templ.Attributes) {
    <i 
        class={fmt.Sprintf("icon icon-%s icon-%s", name, size)}
        {...mergeAttributes(attrs...)}
    />
}

// Text label with size variants
templ Label(text string, size string, attrs ...templ.Attributes) {
    <span 
        class={fmt.Sprintf("label label-%s", size)}
        {...mergeAttributes(attrs...)}
    >
        {text}
    </span>
}

// Loading spinner
templ Spinner(size string) {
    <svg 
        class={fmt.Sprintf("animate-spin spinner-%s", size)} 
        xmlns="http://www.w3.org/2000/svg" 
        fill="none" 
        viewBox="0 0 24 24"
    >
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
    </svg>
}
```

### 2.2 Composition Pattern

```go
// views/components/atoms/button_v2.templ
package atoms

type ButtonProps struct {
    ID          string
    Variant     ButtonVariant
    Size        ButtonSize
    Icon        string
    IconPos     IconPosition
    Loading     bool
    Disabled    bool
    Type        string
    AriaLabel   string
    
    // Runtime customization
    Theme       *ThemeOverride
    
    // Event handlers
    OnClick     string
    HXPost      string
    HXTarget    string
    AlpineData  string
}

// Composable button using primitives
templ Button(props ButtonProps, children ...templ.Component) {
    @ButtonBase(templ.Attributes{
        "id": props.ID,
        "type": getType(props.Type),
        "disabled": props.Disabled,
        "aria-label": props.AriaLabel,
        "class": buttonClasses(props),
        "onclick": props.OnClick,
        "hx-post": props.HXPost,
        "hx-target": props.HXTarget,
        "x-data": props.AlpineData,
    }) {
        if props.Loading {
            @Spinner(string(props.Size))
        } else {
            if props.Icon != "" && props.IconPos == IconLeft {
                @Icon(props.Icon, string(props.Size))
            }
            
            if len(children) > 0 {
                for _, child := range children {
                    @child
                }
            }
            
            if props.Icon != "" && props.IconPos == IconRight {
                @Icon(props.Icon, string(props.Size))
            }
        }
    }
}

// Helper function for class generation
func buttonClasses(props ButtonProps) string {
    classes := []string{"btn"}
    
    // Apply semantic classes
    classes = append(classes, fmt.Sprintf("btn-%s", props.Variant))
    classes = append(classes, fmt.Sprintf("btn-%s", props.Size))
    
    // Apply runtime theme overrides if provided
    if props.Theme != nil {
        classes = append(classes, props.Theme.Classes...)
    }
    
    return strings.Join(classes, " ")
}

// Higher-order component factory
templ PrimaryButton(text string, attrs ...templ.Attributes) {
    @Button(ButtonProps{
        Variant: ButtonPrimary,
        Size: ButtonMedium,
    }, mergeAttributes(attrs...)) {
        @Label(text, "md")
    }
}
```

### 2.3 Advanced Composition with Scripts

```go
// views/components/atoms/interactive.templ
package atoms

// Reusable client-side behavior
script rippleEffect() {
    document.querySelectorAll('.ripple').forEach(button => {
        button.addEventListener('click', function(e) {
            const ripple = document.createElement('span');
            const rect = this.getBoundingClientRect();
            const size = Math.max(rect.width, rect.height);
            const x = e.clientX - rect.left - size / 2;
            const y = e.clientY - rect.top - size / 2;
            
            ripple.style.width = ripple.style.height = size + 'px';
            ripple.style.left = x + 'px';
            ripple.style.top = y + 'px';
            ripple.classList.add('ripple-effect');
            
            this.appendChild(ripple);
            setTimeout(() => ripple.remove(), 600);
        });
    });
}

// Interactive button with ripple effect
templ InteractiveButton(props ButtonProps, children ...templ.Component) {
    @Button(ButtonProps{
        ...props,
        class: props.class + " ripple relative overflow-hidden",
    }, children...) 
    
    @rippleEffect()
}
```

## 3. Conditional Rendering Framework

### 3.1 Backend-Driven Rendering

```go
// views/components/atoms/conditional.go
package atoms

import (
    "github.com/ruun/schema"
    "github.com/ruun/auth"
)

// RenderContext provides information for conditional rendering
type RenderContext struct {
    User      *auth.User
    Tenant    *auth.Tenant
    Features  map[string]bool
    Device    DeviceType
}

// ComponentRenderer decides what to render based on context
type ComponentRenderer struct {
    ctx RenderContext
}

func NewComponentRenderer(ctx RenderContext) *ComponentRenderer {
    return &ComponentRenderer{ctx: ctx}
}

// ConditionalComponent interface for components that support conditional rendering
type ConditionalComponent interface {
    ShouldRender(ctx RenderContext) bool
    GetPermission() string
    GetFeatureFlag() string
}
```

```go
// views/components/atoms/field.templ
package atoms

// Conditional field rendering based on schema and context
templ ConditionalField(field schema.Field, ctx RenderContext) {
    // Backend decides visibility
    if !field.IsVisible(ctx.User, ctx.Tenant) {
        // Don't send any HTML for invisible fields
        return
    }
    
    // Check feature flags
    if field.FeatureFlag != "" && !ctx.Features[field.FeatureFlag] {
        return
    }
    
    // Render based on permissions
    switch {
    case !field.CanRead(ctx.User):
        // No output
    case !field.CanWrite(ctx.User):
        @ReadOnlyField(field)
    case field.IsRequired(ctx.Tenant):
        @RequiredField(field)
    default:
        @EditableField(field)
    }
}

// Composable field components
templ ReadOnlyField(field schema.Field) {
    <div class="field field-readonly">
        @Label(field.Label, "sm", templ.Attributes{"for": field.ID})
        <div class="field-value">{field.Value}</div>
    </div>
}

templ RequiredField(field schema.Field) {
    <div class="field field-required">
        @Label(field.Label, "sm", templ.Attributes{"for": field.ID})
        @Input(InputProps{
            ID:       field.ID,
            Name:     field.Name,
            Value:    field.Value,
            Required: true,
            Pattern:  field.ValidationPattern,
        })
        <span class="required-indicator">*</span>
    </div>
}
```

### 3.2 Dynamic Component Trees

```go
// views/components/atoms/dynamic.templ
package atoms

// Dynamic form generation based on schema
templ DynamicForm(schema schema.FormSchema, ctx RenderContext) {
    <form class="dynamic-form" hx-post={schema.Action}>
        for _, section := range schema.Sections {
            if section.IsVisible(ctx) {
                @FormSection(section, ctx)
            }
        }
        
        if schema.ShowActions(ctx) {
            @FormActions(schema.Actions, ctx)
        }
    </form>
}

templ FormSection(section schema.Section, ctx RenderContext) {
    <fieldset class="form-section">
        if section.Title != "" {
            <legend class="section-title">{section.Title}</legend>
        }
        
        <div class={sectionLayoutClasses(section.Layout)}>
            for _, field := range section.Fields {
                @ConditionalField(field, ctx)
            }
        </div>
    </fieldset>
}

// Helper for responsive layouts
func sectionLayoutClasses(layout schema.Layout) string {
    switch layout {
    case schema.LayoutGrid:
        return "grid grid-cols-1 md:grid-cols-2 gap-4"
    case schema.LayoutStack:
        return "flex flex-col space-y-4"
    default:
        return "space-y-4"
    }
}
```

## 4. Runtime Theme Customization

### 4.1 Theme Manager

```go
// pkg/theme/manager.go
package theme

import (
    "fmt"
    "sync"
)

type Theme struct {
    ID          string
    Name        string
    Tokens      map[string]string
    Components  map[string]ComponentTheme
}

type ComponentTheme struct {
    Classes []string
    Styles  map[string]string
}

type ThemeManager struct {
    mu       sync.RWMutex
    themes   map[string]*Theme
    defaults *Theme
}

func (tm *ThemeManager) GetTenantTheme(tenantID string) *Theme {
    tm.mu.RLock()
    defer tm.mu.RUnlock()
    
    if theme, exists := tm.themes[tenantID]; exists {
        return theme
    }
    return tm.defaults
}

// Generate CSS for a tenant's theme
func (tm *ThemeManager) GenerateCSS(tenantID string) string {
    theme := tm.GetTenantTheme(tenantID)
    
    var css strings.Builder
    css.WriteString(fmt.Sprintf("[data-tenant=\"%s\"] {\n", tenantID))
    
    for token, value := range theme.Tokens {
        css.WriteString(fmt.Sprintf("  --%s: %s;\n", token, value))
    }
    
    css.WriteString("}\n")
    
    return css.String()
}
```

### 4.2 Theme Application Component

```go
// views/components/theme.templ
package components

import (
    "github.com/ruun/pkg/theme"
)

// Apply tenant theme at runtime
templ ApplyTheme(tenantID string, manager *theme.ThemeManager) {
    <style>
        {templ.Raw(manager.GenerateCSS(tenantID))}
    </style>
    <div data-tenant={tenantID} class="theme-root">
        {children...}
    </div>
}

// Theme preview for admin interface
templ ThemePreview(theme *theme.Theme) {
    <div class="theme-preview" data-theme-id={theme.ID}>
        <style>
            .theme-preview[data-theme-id="{theme.ID}"] {
                {templ.Raw(generatePreviewCSS(theme))}
            }
        </style>
        
        <h3>{theme.Name}</h3>
        <div class="component-samples">
            @Button(ButtonProps{Variant: ButtonPrimary}, Label("Primary Button", "md"))
            @Button(ButtonProps{Variant: ButtonSecondary}, Label("Secondary Button", "md"))
            @Input(InputProps{Placeholder: "Sample input"})
        </div>
    </div>
}
```

## 5. Templ-Specific Patterns and Optimizations

### 5.1 Component Factories

```go
// views/components/atoms/factories.templ
package atoms

// Factory for creating typed components
type ComponentFactory struct {
    defaults map[string]interface{}
}

func NewComponentFactory() *ComponentFactory {
    return &ComponentFactory{
        defaults: make(map[string]interface{}),
    }
}

// Create button with defaults
func (cf *ComponentFactory) Button(overrides ...ButtonProps) templ.Component {
    props := ButtonProps{
        Variant: ButtonPrimary,
        Size:    ButtonMedium,
        Type:    "button",
    }
    
    if len(overrides) > 0 {
        // Merge overrides
        props = mergeProps(props, overrides[0])
    }
    
    return Button(props)
}

// Batch component creation
templ ButtonGroup(buttons []ButtonConfig) {
    <div class="btn-group" role="group">
        for i, config := range buttons {
            @Button(ButtonProps{
                ID:      fmt.Sprintf("btn-%d", i),
                Variant: config.Variant,
                Size:    config.Size,
                OnClick: config.Handler,
            }) {
                @Label(config.Text, string(config.Size))
            }
        }
    </div>
}
```

### 5.2 Performance Optimizations

```go
// views/components/atoms/optimized.templ
package atoms

// Lazy loading for heavy components
templ LazyComponent(componentID string, loader templ.Component) {
    <div 
        id={componentID}
        hx-get={fmt.Sprintf("/components/load/%s", componentID)}
        hx-trigger="intersect once"
        class="lazy-component"
    >
        @loader
    </div>
}

// Memoized component rendering
var componentCache = &sync.Map{}

func MemoizedComponent(key string, generator func() templ.Component) templ.Component {
    if cached, exists := componentCache.Load(key); exists {
        return cached.(templ.Component)
    }
    
    component := generator()
    componentCache.Store(key, component)
    return component
}

// Streaming large lists
templ StreamedList(items []Item, batchSize int) {
    <div class="streamed-list" hx-ext="stream">
        for i := 0; i < len(items); i += batchSize {
            end := i + batchSize
            if end > len(items) {
                end = len(items)
            }
            
            @ListBatch(items[i:end], i/batchSize)
        }
    </div>
}
```

### 5.3 Advanced Attribute Handling

```go
// views/components/atoms/attributes.go
package atoms

import "github.com/a-h/templ"

// Merge multiple attribute sets
func mergeAttributes(attrs ...templ.Attributes) templ.Attributes {
    result := make(templ.Attributes)
    
    for _, attr := range attrs {
        for k, v := range attr {
            if k == "class" {
                // Merge classes
                if existing, ok := result[k]; ok {
                    result[k] = existing.(string) + " " + v.(string)
                } else {
                    result[k] = v
                }
            } else {
                result[k] = v
            }
        }
    }
    
    return result
}

// Conditional attributes
func conditionalAttrs(condition bool, attrs templ.Attributes) templ.Attributes {
    if condition {
        return attrs
    }
    return templ.Attributes{}
}

// Data attribute builder
func dataAttributes(data map[string]string) templ.Attributes {
    attrs := make(templ.Attributes)
    for k, v := range data {
        attrs[fmt.Sprintf("data-%s", k)] = v
    }
    return attrs
}
```

## 6. Migration Guide

### 6.1 Step-by-Step Migration

1. **Update Tailwind Configuration**
   ```bash
   # Update tailwind.config.js with new token mappings
   npm run build:css
   ```

2. **Refactor Existing Components**
   ```go
   // Before
   class="px-[var(--space-4)] rounded-[var(--radius-md)]"
   
   // After  
   class="px-4 rounded-md"
   ```

3. **Implement Base Primitives**
   - Create `primitives.templ`
   - Update existing components to use primitives
   - Test composition patterns

4. **Add Theme Support**
   ```go
   // In layout.templ
   @ApplyTheme(tenant.ID, themeManager) {
       @content
   }
   ```

5. **Update Component Usage**
   ```go
   // Before
   @Button("Click me", "primary", "medium")
   
   // After
   @Button(ButtonProps{
       Variant: ButtonPrimary,
       Size: ButtonMedium,
   }) {
       @Label("Click me", "md")
   }
   ```

### 6.2 Backwards Compatibility

Create adapter components for gradual migration:

```go
// views/components/atoms/adapters.templ

// Legacy button adapter
templ LegacyButton(text string, variant string, size string) {
    @Button(ButtonProps{
        Variant: parseVariant(variant),
        Size: parseSize(size),
    }) {
        @Label(text, size)
    }
}
```

## 7. Best Practices

1. **Component Naming**
   - Use clear, descriptive names
   - Follow `ComponentVariant` pattern
   - Group related components

2. **Props Design**
   - Required fields first
   - Sensible defaults
   - Avoid prop drilling

3. **Performance**
   - Minimize component depth
   - Use lazy loading for heavy components
   - Implement proper memoization

4. **Testing**
   - Unit test component logic
   - Visual regression testing
   - Accessibility testing

5. **Documentation**
   - Document all props
   - Provide usage examples
   - Maintain component catalog

## Conclusion

This refactoring strategy provides a robust foundation for a modern, maintainable component system that fully leverages Templ's capabilities while supporting enterprise requirements like multi-tenancy and runtime customization. The hybrid approach to design tokens ensures both compile-time optimization and runtime flexibility.