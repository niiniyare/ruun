# Design Token System - Final Architecture

## Executive Decision: Hybrid CSS Variables + TailwindCSS Approach

After analyzing all requirements, I recommend a **Hybrid Architecture** that combines the best of both worlds:

1. **CSS Variables with HSL values** for runtime customization
2. **TailwindCSS semantic classes** for developer experience  
3. **Go backend generation** from schema system
4. **Component-level overrides** through props
5. **Utils integration** for conflict resolution

## 1. Token Implementation Strategy (FINAL)

### 1.1 Three-Layer Token System

```css
/* views/styles/design-tokens.css */

/* Layer 1: Base Tokens (from schema) */
:root {
  /* Color Primitives - HSL for manipulation */
  --color-primary-h: 222.2;
  --color-primary-s: 47.4%;
  --color-primary-l: 11.2%;
  --primary: var(--color-primary-h) var(--color-primary-s) var(--color-primary-l);
  --primary-foreground: 210 40% 98%;
  
  /* Semantic Colors */
  --success: 142 71% 45%;
  --warning: 38 92% 50%;
  --error: 0 84% 60%;
  
  /* Spacing Tokens */
  --spacing-xs: 0.25rem;
  --spacing-sm: 0.5rem;
  --spacing-md: 1rem;
  --spacing-lg: 1.5rem;
  --spacing-xl: 2rem;
  
  /* Typography */
  --font-size-xs: 0.75rem;
  --font-size-sm: 0.875rem;
  --font-size-base: 1rem;
  --font-size-lg: 1.125rem;
  
  /* Radius */
  --radius-sm: 0.125rem;
  --radius-md: 0.375rem;
  --radius-lg: 0.5rem;
  
  /* Shadow Tokens */
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1);
}

/* Layer 2: Component Tokens */
:root {
  /* Button Tokens */
  --btn-primary-bg: var(--primary);
  --btn-primary-fg: var(--primary-foreground);
  --btn-primary-border: var(--primary);
  
  --btn-secondary-bg: var(--secondary);
  --btn-secondary-fg: var(--secondary-foreground);
  
  /* Input Tokens */
  --input-bg: var(--background);
  --input-border: var(--border);
  --input-focus-ring: var(--ring);
  
  /* Card Tokens */
  --card-bg: var(--card);
  --card-border: var(--border);
}

/* Layer 3: TailwindCSS Integration */
@layer components {
  .btn {
    @apply inline-flex items-center justify-center rounded-md font-medium;
    @apply focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2;
    @apply disabled:pointer-events-none disabled:opacity-50;
  }
  
  .btn-primary {
    background-color: hsl(var(--btn-primary-bg));
    color: hsl(var(--btn-primary-fg));
    border-color: hsl(var(--btn-primary-border));
  }
  
  .btn-primary:hover {
    background-color: hsl(var(--color-primary-h) var(--color-primary-s) calc(var(--color-primary-l) - 5%));
  }
}

/* Layer 4: Tenant Overrides */
[data-tenant="enterprise"] {
  --primary: 217 91% 60%;
  --spacing-md: 1.25rem;
}

[data-tenant="startup"] {
  --primary: 142 71% 45%;
  --radius-md: 0.75rem;
}
```

### 1.2 TailwindCSS Configuration

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
        success: 'hsl(var(--success) / <alpha-value>)',
        warning: 'hsl(var(--warning) / <alpha-value>)',
        error: 'hsl(var(--error) / <alpha-value>)',
      },
      spacing: {
        'xs': 'var(--spacing-xs)',
        'sm': 'var(--spacing-sm)', 
        'md': 'var(--spacing-md)',
        'lg': 'var(--spacing-lg)',
        'xl': 'var(--spacing-xl)',
      },
      fontSize: {
        'xs': 'var(--font-size-xs)',
        'sm': 'var(--font-size-sm)',
        'base': 'var(--font-size-base)',
        'lg': 'var(--font-size-lg)',
      },
      borderRadius: {
        'sm': 'var(--radius-sm)',
        'md': 'var(--radius-md)',
        'lg': 'var(--radius-lg)',
      },
      boxShadow: {
        'sm': 'var(--shadow-sm)',
        'md': 'var(--shadow-md)',
      }
    },
  },
  plugins: [
    // Custom plugin for component tokens
    function({ addComponents }) {
      addComponents({
        '.btn-primary': {
          backgroundColor: 'hsl(var(--btn-primary-bg))',
          color: 'hsl(var(--btn-primary-fg))',
          borderColor: 'hsl(var(--btn-primary-border))',
        },
        // ... other component classes
      })
    }
  ],
}
```

## 2. Schema Integration Pattern (FINAL)

### 2.1 Backend Token Generation

```go
// pkg/theme/generator.go
package theme

import (
    "fmt"
    "strings"
    "github.com/ruun/pkg/schema"
)

type TokenGenerator struct {
    schema *schema.ThemeSchema
}

func NewTokenGenerator(schema *schema.ThemeSchema) *TokenGenerator {
    return &TokenGenerator{schema: schema}
}

// Generate CSS variables from schema
func (tg *TokenGenerator) GenerateCSS() string {
    var css strings.Builder
    
    // Base tokens
    css.WriteString(":root {\n")
    for token, value := range tg.schema.BaseTokens {
        css.WriteString(fmt.Sprintf("  --%s: %s;\n", token, value))
    }
    css.WriteString("}\n\n")
    
    // Tenant-specific overrides
    for tenantID, theme := range tg.schema.TenantThemes {
        css.WriteString(fmt.Sprintf("[data-tenant=\"%s\"] {\n", tenantID))
        for token, value := range theme.Overrides {
            css.WriteString(fmt.Sprintf("  --%s: %s;\n", token, value))
        }
        css.WriteString("}\n\n")
    }
    
    return css.String()
}

// Generate Go constants for type safety
func (tg *TokenGenerator) GenerateGoTokens() string {
    var go_code strings.Builder
    
    go_code.WriteString("package tokens\n\n")
    go_code.WriteString("// Auto-generated design tokens\n")
    go_code.WriteString("const (\n")
    
    for token := range tg.schema.BaseTokens {
        constName := strings.ReplaceAll(strings.ToUpper(token), "-", "_")
        go_code.WriteString(fmt.Sprintf("    Token%s = \"--%s\"\n", constName, token))
    }
    
    go_code.WriteString(")\n")
    return go_code.String()
}
```

### 2.2 Runtime Theme Application

```go
// views/components/theme/provider.templ
package theme

import (
    "views/components/utils"
    "github.com/ruun/pkg/theme"
)

type ThemeProviderProps struct {
    TenantID    string
    UserTheme   string  // light, dark, auto
    CustomCSS   string  // Additional custom overrides
}

templ ThemeProvider(props ThemeProviderProps, generator *theme.TokenGenerator) {
    // Generate base CSS
    <style id="design-tokens">
        { templ.Raw(generator.GenerateCSS()) }
        
        if props.CustomCSS != "" {
            { templ.Raw(props.CustomCSS) }
        }
    </style>
    
    // Apply theme attributes
    <div 
        data-tenant={ props.TenantID }
        data-theme={ props.UserTheme }
        class="theme-root"
    >
        { children... }
    </div>
}

// Theme toggle component with Alpine.js
templ ThemeToggle() {
    <div x-data="{ theme: localStorage.theme || 'light' }">
        @atoms.Button(atoms.ButtonProps{
            Variant: atoms.ButtonGhost,
            XOn: map[string]string{
                "click": "theme = theme === 'light' ? 'dark' : 'light'; localStorage.theme = theme; document.documentElement.setAttribute('data-theme', theme)",
            },
            XBind: map[string]string{
                "aria-label": "theme === 'light' ? 'Switch to dark mode' : 'Switch to light mode'",
            },
        }) {
            @atoms.Icon(atoms.IconProps{
                XBind: map[string]string{
                    "name": "theme === 'light' ? 'moon' : 'sun'",
                },
                Size: "md",
            })
        }
    </div>
}
```

## 3. Atomic Design Token Flow (FINAL)

### 3.1 Token Resolution Hierarchy

```
Schema Tokens → CSS Variables → TailwindCSS → Component Props → Local Overrides
```

### 3.2 Component Token Usage

```go
// views/components/atoms/button_with_tokens.templ
package atoms

import (
    "views/components/utils"
    "views/tokens"  // Auto-generated from schema
)

type ButtonTokenProps struct {
    // Allow component-level token overrides
    CustomTokens    map[string]string  // e.g., {"--btn-primary-bg": "hsl(120 100% 50%)"}
    ThemeVariant    string             // Override tenant theme
}

templ ButtonWithTokens(props ButtonProps, tokenProps ButtonTokenProps) {
    // Generate inline styles for custom tokens
    inlineStyles := buildInlineStyles(tokenProps.CustomTokens)
    
    // Build classes with token awareness
    classes := utils.TwMerge(
        "btn",
        fmt.Sprintf("btn-%s", props.Variant),
        fmt.Sprintf("btn-%s", props.Size),
        utils.If(tokenProps.ThemeVariant != "", fmt.Sprintf("theme-%s", tokenProps.ThemeVariant)),
        props.ClassName,
    )
    
    <button 
        class={ classes }
        style={ inlineStyles }
        data-token-variant={ tokenProps.ThemeVariant }
        { buildButtonAttributes(props)... }
    >
        { children... }
    </button>
}

func buildInlineStyles(customTokens map[string]string) string {
    if len(customTokens) == 0 {
        return ""
    }
    
    var styles []string
    for token, value := range customTokens {
        styles = append(styles, fmt.Sprintf("%s: %s", token, value))
    }
    
    return strings.Join(styles, "; ")
}
```

### 3.3 Molecule-Level Token Composition

```go
// views/components/molecules/form_field_with_tokens.templ
package molecules

type FormFieldTokenProps struct {
    InputTokens map[string]string
    LabelTokens map[string]string
    ErrorTokens map[string]string
}

templ FormFieldWithTokens(props FormFieldProps, tokenProps FormFieldTokenProps) {
    <div class="form-field" data-field-type={ string(props.Type) }>
        @atoms.LabelWithTokens(atoms.LabelProps{
            For: props.ID,
            Required: props.Required,
        }, atoms.LabelTokenProps{
            CustomTokens: tokenProps.LabelTokens,
        }) {
            { props.Label }
        }
        
        @atoms.InputWithTokens(atoms.InputProps{
            ID: props.ID,
            Type: props.Type,
            Value: props.Value,
            Invalid: props.ErrorMessage != "",
        }, atoms.InputTokenProps{
            CustomTokens: tokenProps.InputTokens,
        })
        
        if props.ErrorMessage != "" {
            @atoms.TextWithTokens(atoms.TextProps{
                Variant: atoms.TextError,
                Size: atoms.TextSM,
            }, atoms.TextTokenProps{
                CustomTokens: tokenProps.ErrorTokens,
            }) {
                { props.ErrorMessage }
            }
        }
    </div>
}
```

## 4. Runtime Customization Approach (FINAL)

### 4.1 Programmatic Theme Changes

```go
// views/components/theme/controls.templ
package theme

templ ThemeCustomizer() {
    <div 
        x-data="{
            customTokens: {},
            
            updateToken(token, value) {
                this.customTokens[token] = value;
                document.documentElement.style.setProperty(token, value);
            },
            
            resetTheme() {
                Object.keys(this.customTokens).forEach(token => {
                    document.documentElement.style.removeProperty(token);
                });
                this.customTokens = {};
            },
            
            exportTheme() {
                return JSON.stringify(this.customTokens, null, 2);
            },
            
            importTheme(themeJson) {
                try {
                    const theme = JSON.parse(themeJson);
                    Object.entries(theme).forEach(([token, value]) => {
                        this.updateToken(token, value);
                    });
                } catch(e) {
                    console.error('Invalid theme JSON:', e);
                }
            }
        }"
        class="theme-customizer"
    >
        <!-- Color Picker for Primary -->
        <div class="control-group">
            <label>Primary Color</label>
            <input 
                type="color" 
                x-on:input="updateToken('--primary', `${$event.target.value.slice(1)} 100% 50%`)"
            />
        </div>
        
        <!-- Spacing Control -->
        <div class="control-group">
            <label>Base Spacing</label>
            <input 
                type="range" 
                min="0.25" 
                max="2" 
                step="0.25"
                x-on:input="updateToken('--spacing-md', `${$event.target.value}rem`)"
            />
        </div>
        
        <!-- Export/Import -->
        <div class="control-group">
            @atoms.Button(atoms.ButtonProps{
                XOn: map[string]string{"click": "navigator.clipboard.writeText(exportTheme())"},
            }) { "Export Theme" }
            
            @atoms.Button(atoms.ButtonProps{
                XOn: map[string]string{"click": "importTheme(prompt('Paste theme JSON:'))"},
            }) { "Import Theme" }
        </div>
    </div>
}
```

### 4.2 Server-Side Theme Generation

```go
// handlers/theme.go
func HandleThemeGeneration(w http.ResponseWriter, r *http.Request) {
    tenantID := getTenantID(r)
    
    // Get theme from schema
    themeSchema := getThemeSchema(tenantID)
    generator := theme.NewTokenGenerator(themeSchema)
    
    // Generate CSS
    css := generator.GenerateCSS()
    
    // Set proper headers for CSS
    w.Header().Set("Content-Type", "text/css")
    w.Header().Set("Cache-Control", "public, max-age=3600")
    
    fmt.Fprint(w, css)
}

// HTMX endpoint for live theme updates
func HandleThemeUpdate(w http.ResponseWriter, r *http.Request) {
    tenantID := getTenantID(r)
    
    // Parse form data for token updates
    var updates map[string]string
    json.NewDecoder(r.Body).Decode(&updates)
    
    // Update theme in database
    updateTenantTheme(tenantID, updates)
    
    // Return updated CSS fragment
    generator := theme.NewTokenGenerator(getThemeSchema(tenantID))
    css := generator.GenerateCSS()
    
    w.Header().Set("Content-Type", "text/css")
    fmt.Fprint(w, css)
}
```

## 5. Developer Experience (FINAL)

### 5.1 Component Token API

```go
// Fluent API for token usage
@atoms.Button(atoms.NewButton().
    Variant(atoms.ButtonPrimary).
    WithTokens(map[string]string{
        "--btn-primary-bg": "hsl(120 100% 40%)",
    }).
    Build()) { "Custom Green Button" }

// Semantic token usage
@atoms.Button(atoms.ButtonProps{
    Variant: atoms.ButtonPrimary,
    TokenVariant: "success",  // Maps to predefined token set
}) { "Success Button" }

// Component-specific token overrides
@molecules.FormField(molecules.FormFieldProps{
    Label: "Email",
    Type: atoms.InputEmail,
}, molecules.FormFieldTokenProps{
    InputTokens: map[string]string{
        "--input-focus-ring": "hsl(var(--success))",
    },
}) 
```

### 5.2 Token Documentation Generator

```go
// pkg/theme/docs.go
func GenerateTokenDocumentation(schema *schema.ThemeSchema) string {
    var docs strings.Builder
    
    docs.WriteString("# Design Token Reference\n\n")
    
    // Base tokens
    docs.WriteString("## Base Tokens\n\n")
    docs.WriteString("| Token | Default Value | Description |\n")
    docs.WriteString("|-------|---------------|-------------|\n")
    
    for token, value := range schema.BaseTokens {
        description := schema.TokenDescriptions[token]
        docs.WriteString(fmt.Sprintf("| `%s` | `%s` | %s |\n", token, value, description))
    }
    
    // Component tokens
    docs.WriteString("\n## Component Tokens\n\n")
    for component, tokens := range schema.ComponentTokens {
        docs.WriteString(fmt.Sprintf("### %s\n\n", component))
        for token, value := range tokens {
            docs.WriteString(fmt.Sprintf("- `%s`: `%s`\n", token, value))
        }
        docs.WriteString("\n")
    }
    
    return docs.String()
}
```

## 6. Migration Strategy (FINAL)

### 6.1 Phase 1: Foundation
```bash
# 1. Generate base token CSS from schema
go run cmd/generate-tokens/main.go

# 2. Update TailwindCSS config
npm run build:css

# 3. Create theme provider component
# views/components/theme/provider.templ
```

### 6.2 Phase 2: Component Migration
```go
// 1. Update button component
// Before
class="bg-blue-500 text-white"

// After
class="btn btn-primary"  // Uses tokens automatically

// 2. Add token props to existing components
@atoms.Button(atoms.ButtonProps{
    Variant: atoms.ButtonPrimary,  // Uses tokens
    ClassName: "custom-class",     // Additional classes merged with utils.TwMerge
})
```

### 6.3 Phase 3: Advanced Features
```go
// 1. Add token customization to admin panel
@theme.ThemeCustomizer()

// 2. Implement tenant theme switching
@theme.ThemeProvider(theme.ThemeProviderProps{
    TenantID: tenant.ID,
})

// 3. Add runtime theme controls
@theme.ThemeToggle()
```

## 7. Performance Considerations (FINAL)

### 7.1 Optimization Strategy
- **CSS Variables**: Minimal runtime performance impact
- **TailwindCSS**: Compile-time optimization for static classes
- **Caching**: Server-side CSS generation with proper cache headers
- **Lazy Loading**: Load tenant themes on-demand
- **Utils Integration**: TwMerge optimizes class conflicts

### 7.2 Bundle Size
```css
/* Generated CSS is minimal and scoped */
:root { --primary: 222 47% 11%; }
[data-tenant="ent"] { --primary: 217 91% 60%; }

.btn-primary { 
    background-color: hsl(var(--primary));
    /* Compiled Tailwind utilities */
}
```

## Final Implementation Decision

**Recommendation: Hybrid CSS Variables + TailwindCSS Architecture**

This approach provides:
- ✅ Runtime tenant customization via CSS variables
- ✅ Excellent developer experience via TailwindCSS
- ✅ Schema integration through Go backend generation
- ✅ Atomic design compliance with token flow
- ✅ Performance optimization through compilation
- ✅ Utils integration for class conflict resolution
- ✅ Component-level token overrides
- ✅ Type safety through generated Go constants

The system bridges schema-driven backend tokens with atomic frontend components while maintaining performance and flexibility for multi-tenant customization.