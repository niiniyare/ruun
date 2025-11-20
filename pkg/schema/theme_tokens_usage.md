# Theme & Tokens Usage Guide

Complete guide for using the Awo ERP theme and design token system with Atomic Design components.

## Table of Contents

1. [Overview](#overview)
2. [Atomic Design Integration](#atomic-design-integration)
3. [Token Architecture](#token-architecture)
4. [Theme Structure](#theme-structure)
5. [Getting Started](#getting-started)
6. [Component Integration](#component-integration)
7. [CSS Compilation](#css-compilation)
8. [Workflows](#workflows)
9. [Best Practices](#best-practices)
10. [Advanced Patterns](#advanced-patterns)

---

## Overview

The Awo ERP theme system provides an enterprise-grade solution that bridges design tokens with Atomic Design components:

```
Design Tokens → CSS Classes → Component Props → UI Components
     ↓              ↓              ↓                ↓
  JSON/YAML    button-primary   ButtonProps     Button Atom
```

### Why This Architecture?

**Design Tokens** provide the single source of truth for design decisions.  
**Atomic Design** provides the component hierarchy for building interfaces.  
**Together** they create a scalable, maintainable design system.

```
┌─────────────────────────────────────────┐
│          DESIGN TOKENS                  │
│  primitives → semantic → components     │
└──────────────┬──────────────────────────┘
               │ compiles to
               ▼
┌─────────────────────────────────────────┐
│          CSS CLASSES                    │
│  .button, .button-primary, .button-md   │
└──────────────┬──────────────────────────┘
               │ used in
               ▼
┌─────────────────────────────────────────┐
│      ATOMIC COMPONENTS (templ)          │
│  Atoms → Molecules → Organisms          │
│  Templates → Pages                      │
└─────────────────────────────────────────┘
```

---

## Atomic Design Integration

### The Atomic Hierarchy

```
Atoms (Basic Building Blocks)
  ↓
Molecules (Simple Combinations)
  ↓
Organisms (Complex Sections)
  ↓
Templates (Page Layouts)
  ↓
Pages (Specific Instances)
```

### 1. Atoms (Tokens → Classes → Props)

**Atoms** are the most basic UI elements. Each atom uses design tokens through compiled CSS classes.

**Example: Button Atom**

```go
// Token Definition (components/button.json)
{
  "button": {
    "primary": {
      "background-color": "semantic.colors.primary",
      "color": "semantic.colors.primary-foreground",
      "border-radius": "semantic.interactive.border-radius-md",
      "padding": "primitives.spacing.sm primitives.spacing.md"
    }
  }
}

// Compiled CSS Class
.button-primary {
  background-color: hsl(217, 91%, 60%);
  color: #ffffff;
  border-radius: 0.5rem;
  padding: 0.75rem 1rem;
}

// Component Props
type ButtonProps struct {
  Variant ButtonVariant  // Maps to .button-{variant}
  Size    ButtonSize     // Maps to .button-{size}
}

// templ Component (Pure Presentation)
templ Button(props ButtonProps) {
  <button class={ buttonClasses(props) }>
    { props.Text }
  </button>
}
```

**Token Flow:**
```txt
Token Path → CSS Class → Component Prop → Rendered HTML
components.button.primary.background-color
  → .button-primary { background-color: ... }
    → ButtonVariant: "primary"
      → <button class="button button-primary">
```


### 2. Molecules (Atoms + Composition)

**Molecules** combine atoms into simple functional groups.

**Example: Search Input Molecule**

```go
// Uses Token-based Atoms
type SearchInputProps struct {
  Placeholder string
  ButtonText  string
}

templ SearchInput(props SearchInputProps) {
  <div class="search-input">
    // Input atom uses .input-base classes from tokens
    @atoms.Input(atoms.InputProps{
      Type:        "text",
      Placeholder: props.Placeholder,
      ClassName:   "search-input__field",
    })
    
    // Button atom uses .button-primary classes from tokens
    @atoms.Button(atoms.ButtonProps{
      Variant: atoms.ButtonPrimary,
      Text:    props.ButtonText,
      Type:    "submit",
    })
  </div>
}
```

**Token Cascade:**
```
Molecule: SearchInput
  ├─ Atom: Input
  │   └─ Uses: .input-base (from tokens)
  └─ Atom: Button
      └─ Uses: .button-primary (from tokens)
```

### 3. Organisms (Molecules + Complex Logic)

**Organisms** are complex UI sections composed of molecules and atoms.

**Example: Data Table Organism**

```go
type DataTableProps struct {
  Headers []string
  Rows    [][]string
  Actions bool
}

templ DataTable(props DataTableProps) {
  <div class="data-table">
    // Table uses .table, .table-header, etc. from tokens
    <table class="table">
      <thead class="table-header">
        <tr>
          for _, header := range props.Headers {
            <th class="table-header-cell">{ header }</th>
          }
        </tr>
      </thead>
      <tbody>
        for _, row := range props.Rows {
          <tr class="table-row">
            for _, cell := range row {
              <td class="table-cell">{ cell }</td>
            }
            if props.Actions {
              <td>
                // Molecule: Action buttons
                @molecules.ActionButtons()
              </td>
            }
          }
        </tr>
      </tbody>
    </table>
  </div>
}
```

**Token Cascade:**
```
Organism: DataTable
  ├─ Uses: .table, .table-header (from tokens)
  └─ Molecule: ActionButtons
      ├─ Atom: EditButton (uses .button-outline)
      └─ Atom: DeleteButton (uses .button-destructive)
```

### 4. Templates (Layout Structure)

**Templates** define page layouts using organisms.

**Example: Dashboard Template**

```go
type DashboardTemplateProps struct {
  Title   string
  Content templ.Component
  Sidebar templ.Component
}

templ DashboardTemplate(props DashboardTemplateProps) {
  <!DOCTYPE html>
  <html>
    <head>
      <link rel="stylesheet" href="/css/theme.css">
    </head>
    <body class="dashboard">
      // Organism: Navigation
      @organisms.Navigation()
      
      <div class="dashboard-layout">
        // Organism: Sidebar
        <aside class="dashboard-sidebar">
          @props.Sidebar
        </aside>
        
        // Content area
        <main class="dashboard-main">
          <h1 class="dashboard-title">{ props.Title }</h1>
          @props.Content
        </main>
      </div>
    </body>
  </html>
}
```

### 5. Pages (Specific Instances)

**Pages** are specific instances of templates with real data.

**Example: User List Page**

```go
func UserListPage(users []User) templ.Component {
  return templates.DashboardTemplate(
    templates.DashboardTemplateProps{
      Title: "Users",
      Sidebar: organisms.UserSidebar(),
      Content: organisms.DataTable(organisms.DataTableProps{
        Headers: []string{"Name", "Email", "Role"},
        Rows:    formatUsers(users),
        Actions: true,
      }),
    },
  )
}
```

---

## Token Architecture

### Three-Tier Token System

```
1. PRIMITIVES (Raw Values)
   "gray-500": "hsl(210, 10%, 58%)"
   ↓ referenced by

2. SEMANTIC (Contextual)
   "text-default": "primitives.colors.gray-500"
   ↓ referenced by

3. COMPONENTS (Application)
   "button.primary.color": "semantic.colors.text-default"
   ↓ compiles to

4. CSS CLASSES
   .button-primary { color: hsl(210, 10%, 58%); }
   ↓ used in

5. COMPONENT PROPS
   ButtonVariant: "primary"
   ↓ renders to

6. HTML
   <button class="button button-primary">
```

### Token-to-Component Mapping

**Token Definition:**
```json
{
  "components": {
    "button": {
      "primary": {
        "background-color": "semantic.colors.primary",
        "color": "semantic.colors.primary-foreground",
        "padding": "primitives.spacing.sm primitives.spacing.md",
        "border-radius": "semantic.interactive.border-radius-md",
        "font-weight": "primitives.typography.font-weight-medium"
      }
    }
  }
}
```

**Compiled CSS:**
```css
.button-primary {
  background-color: hsl(217, 91%, 60%);
  color: #ffffff;
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  font-weight: 500;
}
```

**Component Usage:**
```go
@atoms.Button(atoms.ButtonProps{
  Variant: atoms.ButtonPrimary, // → class="button button-primary"
})
```

---

## Theme Structure

### Option 1: Multi-File Theme (Recommended for Large Themes)

Organize theme files by concern for better maintainability:

```
themes/
  acme-brand/
    theme.json              # Theme metadata & configuration
    tokens.json             # Primitive tokens
    semantic.json           # Semantic tokens
    components/             # Component tokens
      button.json
      input.json
      badge.json
      card.json
      table.json
      modal.json
      form.json
      navigation.json
```

#### File: `theme.json` (Metadata)

```json
{
  "id": "acme-brand",
  "name": "Acme Brand Theme",
  "description": "Official Acme Corporation brand theme",
  "version": "1.0.0",
  "author": "Acme Design Team",
  
  "darkMode": {
    "enabled": true,
    "default": false,
    "strategy": "auto"
  },
  
  "accessibility": {
    "minContrastRatio": 4.5,
    "focusIndicator": true,
    "keyboardNav": true
  },
  
  "conditions": [
    {
      "id": "admin-highlight",
      "expression": "user.role == 'admin'",
      "priority": 100,
      "tokenOverrides": {
        "semantic.colors.primary": "primitives.colors.brand-accent"
      }
    }
  ],
  
  "meta": {
    "tags": ["brand", "corporate", "light", "dark"],
    "license": "proprietary"
  }
}
```

#### File: `tokens.json` (Primitives)

```json
{
  "primitives": {
    "colors": {
      "brand-primary": "hsl(210, 100%, 45%)",
      "brand-secondary": "hsl(340, 82%, 52%)",
      "brand-accent": "hsl(45, 100%, 51%)",
      
      "gray-50": "hsl(210, 20%, 98%)",
      "gray-100": "hsl(210, 20%, 95%)",
      "gray-200": "hsl(210, 16%, 93%)",
      "gray-300": "hsl(210, 14%, 89%)",
      "gray-400": "hsl(210, 12%, 78%)",
      "gray-500": "hsl(210, 10%, 64%)",
      "gray-600": "hsl(210, 10%, 48%)",
      "gray-700": "hsl(210, 12%, 36%)",
      "gray-800": "hsl(210, 16%, 24%)",
      "gray-900": "hsl(210, 20%, 14%)",
      
      "white": "#ffffff",
      "black": "#000000",
      
      "success": "hsl(142, 71%, 45%)",
      "warning": "hsl(45, 93%, 47%)",
      "error": "hsl(0, 84%, 60%)",
      "info": "hsl(200, 100%, 56%)"
    },
    
    "spacing": {
      "xs": "0.5rem",
      "sm": "0.75rem",
      "md": "1rem",
      "lg": "1.5rem",
      "xl": "2rem",
      "2xl": "3rem",
      "3xl": "4rem"
    },
    
    "radius": {
      "none": "0",
      "sm": "0.25rem",
      "md": "0.5rem",
      "lg": "0.75rem",
      "xl": "1rem",
      "full": "9999px"
    },
    
    "typography": {
      "font-size-xs": "0.75rem",
      "font-size-sm": "0.875rem",
      "font-size-base": "1rem",
      "font-size-lg": "1.125rem",
      "font-size-xl": "1.25rem",
      "font-size-2xl": "1.5rem",
      "font-size-3xl": "1.875rem",
      
      "font-weight-normal": "400",
      "font-weight-medium": "500",
      "font-weight-semibold": "600",
      "font-weight-bold": "700",
      
      "line-height-tight": "1.25",
      "line-height-normal": "1.5",
      "line-height-relaxed": "1.75",
      
      "font-family-sans": "'Inter', -apple-system, sans-serif",
      "font-family-mono": "'Fira Code', monospace"
    },
    
    "shadows": {
      "none": "none",
      "xs": "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
      "sm": "0 2px 4px 0 rgba(0, 0, 0, 0.06)",
      "md": "0 4px 6px -1px rgba(0, 0, 0, 0.1)",
      "lg": "0 10px 15px -3px rgba(0, 0, 0, 0.1)",
      "xl": "0 20px 25px -5px rgba(0, 0, 0, 0.1)",
      "2xl": "0 25px 50px -12px rgba(0, 0, 0, 0.25)"
    },
    
    "borders": {
      "border-width-thin": "1px",
      "border-width-medium": "2px",
      "border-width-thick": "4px",
      "border-style-solid": "solid",
      "border-style-dashed": "dashed"
    },
    
    "animation": {
      "duration-fast": "150ms",
      "duration-normal": "200ms",
      "duration-slow": "300ms",
      "easing-in-out": "cubic-bezier(0.4, 0, 0.2, 1)",
      "easing-spring": "cubic-bezier(0.68, -0.55, 0.265, 1.55)"
    },
    
    "z-index": {
      "base": "1",
      "dropdown": "10",
      "sticky": "20",
      "modal": "50",
      "tooltip": "100"
    }
  }
}
```

#### File: `semantic.json` (Semantic Tokens)

```json
{
  "semantic": {
    "colors": {
      "background": "primitives.colors.white",
      "background-subtle": "primitives.colors.gray-50",
      "background-muted": "primitives.colors.gray-100",
      
      "foreground": "primitives.colors.gray-900",
      "foreground-subtle": "primitives.colors.gray-600",
      "foreground-muted": "primitives.colors.gray-500",
      
      "primary": "primitives.colors.brand-primary",
      "primary-foreground": "primitives.colors.white",
      "secondary": "primitives.colors.brand-secondary",
      "secondary-foreground": "primitives.colors.white",
      "accent": "primitives.colors.brand-accent",
      "accent-foreground": "primitives.colors.white",
      
      "border": "primitives.colors.gray-300",
      "border-subtle": "primitives.colors.gray-200",
      "border-emphasis": "primitives.colors.gray-400",
      
      "success": "primitives.colors.success",
      "success-foreground": "primitives.colors.white",
      "warning": "primitives.colors.warning",
      "warning-foreground": "primitives.colors.white",
      "error": "primitives.colors.error",
      "error-foreground": "primitives.colors.white",
      "info": "primitives.colors.info",
      "info-foreground": "primitives.colors.white"
    },
    
    "spacing": {
      "component-tight": "primitives.spacing.xs",
      "component-default": "primitives.spacing.sm",
      "component-loose": "primitives.spacing.md",
      "layout-section": "primitives.spacing.2xl",
      "layout-page": "primitives.spacing.3xl",
      "stack-tight": "primitives.spacing.xs",
      "stack-default": "primitives.spacing.md",
      "stack-loose": "primitives.spacing.lg"
    },
    
    "typography": {
      "heading-font-size": "primitives.typography.font-size-2xl",
      "heading-font-weight": "primitives.typography.font-weight-bold",
      "heading-line-height": "primitives.typography.line-height-tight",
      
      "body-font-size": "primitives.typography.font-size-base",
      "body-font-weight": "primitives.typography.font-weight-normal",
      "body-line-height": "primitives.typography.line-height-normal",
      "body-font-family": "primitives.typography.font-family-sans",
      
      "label-font-size": "primitives.typography.font-size-sm",
      "label-font-weight": "primitives.typography.font-weight-medium",
      
      "caption-font-size": "primitives.typography.font-size-xs",
      "caption-font-weight": "primitives.typography.font-weight-normal",
      
      "code-font-size": "primitives.typography.font-size-sm",
      "code-font-family": "primitives.typography.font-family-mono"
    },
    
    "interactive": {
      "border-radius-sm": "primitives.radius.sm",
      "border-radius-md": "primitives.radius.md",
      "border-radius-lg": "primitives.radius.lg",
      
      "shadow-sm": "primitives.shadows.sm",
      "shadow-md": "primitives.shadows.md",
      "shadow-lg": "primitives.shadows.lg",
      
      "transition-fast": "primitives.animation.duration-fast",
      "transition-normal": "primitives.animation.duration-normal",
      "transition-slow": "primitives.animation.duration-slow"
    }
  }
}
```

#### File: `components/button.json`

```json
{
  "components": {
    "button": {
      "base": {
        "font-size": "semantic.typography.body-font-size",
        "font-weight": "primitives.typography.font-weight-medium",
        "line-height": "primitives.typography.line-height-normal",
        "border-width": "0",
        "cursor": "pointer",
        "transition": "all semantic.interactive.transition-fast primitives.animation.easing-in-out"
      },
      
      "primary": {
        "background-color": "semantic.colors.primary",
        "color": "semantic.colors.primary-foreground",
        "border": "none"
      },
      
      "secondary": {
        "background-color": "semantic.colors.secondary",
        "color": "semantic.colors.secondary-foreground",
        "border": "none"
      },
      
      "outline": {
        "background-color": "transparent",
        "color": "semantic.colors.primary",
        "border": "primitives.borders.border-width-thin solid semantic.colors.primary"
      },
      
      "destructive": {
        "background-color": "semantic.colors.error",
        "color": "semantic.colors.error-foreground",
        "border": "none"
      },
      
      "ghost": {
        "background-color": "transparent",
        "color": "semantic.colors.primary",
        "border": "none"
      },
      
      "link": {
        "background-color": "transparent",
        "color": "semantic.colors.primary",
        "border": "none",
        "text-decoration": "underline"
      },
      
      "xs": {
        "padding": "0.25rem 0.5rem",
        "border-radius": "primitives.radius.sm",
        "font-size": "primitives.typography.font-size-xs"
      },
      
      "sm": {
        "padding": "0.5rem 0.75rem",
        "border-radius": "semantic.interactive.border-radius-sm",
        "font-size": "primitives.typography.font-size-sm"
      },
      
      "md": {
        "padding": "primitives.spacing.sm primitives.spacing.md",
        "border-radius": "semantic.interactive.border-radius-md",
        "font-size": "semantic.typography.body-font-size"
      },
      
      "lg": {
        "padding": "0.75rem 1.5rem",
        "border-radius": "semantic.interactive.border-radius-md",
        "font-size": "primitives.typography.font-size-lg"
      },
      
      "xl": {
        "padding": "primitives.spacing.md primitives.spacing.xl",
        "border-radius": "semantic.interactive.border-radius-lg",
        "font-size": "primitives.typography.font-size-xl"
      },
      
      "disabled": {
        "opacity": "0.5",
        "cursor": "not-allowed"
      },
      
      "loading": {
        "opacity": "0.7",
        "cursor": "wait"
      }
    }
  }
}
```

#### File: `components/input.json`

```json
{
  "components": {
    "input": {
      "base": {
        "background-color": "semantic.colors.background",
        "border": "primitives.borders.border-width-thin solid semantic.colors.border",
        "border-radius": "semantic.interactive.border-radius-md",
        "padding": "primitives.spacing.sm",
        "color": "semantic.colors.foreground",
        "font-size": "semantic.typography.body-font-size",
        "line-height": "semantic.typography.body-line-height",
        "transition": "all semantic.interactive.transition-fast"
      },
      
      "focus": {
        "border-color": "semantic.colors.primary",
        "outline": "2px solid semantic.colors.primary",
        "outline-offset": "2px"
      },
      
      "error": {
        "border-color": "semantic.colors.error"
      },
      
      "disabled": {
        "opacity": "0.5",
        "cursor": "not-allowed",
        "background-color": "semantic.colors.background-muted"
      }
    }
  }
}
```

#### File: `components/badge.json`

```json
{
  "components": {
    "badge": {
      "base": {
        "display": "inline-flex",
        "align-items": "center",
        "padding": "0.125rem primitives.spacing.xs",
        "border-radius": "primitives.radius.sm",
        "font-size": "semantic.typography.caption-font-size",
        "font-weight": "primitives.typography.font-weight-medium",
        "line-height": "1"
      },
      
      "default": {
        "background-color": "semantic.colors.background-muted",
        "color": "semantic.colors.foreground-muted"
      },
      
      "primary": {
        "background-color": "semantic.colors.primary",
        "color": "semantic.colors.primary-foreground"
      },
      
      "success": {
        "background-color": "semantic.colors.success",
        "color": "semantic.colors.success-foreground"
      },
      
      "warning": {
        "background-color": "semantic.colors.warning",
        "color": "semantic.colors.warning-foreground"
      },
      
      "error": {
        "background-color": "semantic.colors.error",
        "color": "semantic.colors.error-foreground"
      },
      
      "outline": {
        "background-color": "transparent",
        "border": "primitives.borders.border-width-thin solid semantic.colors.border",
        "color": "semantic.colors.foreground"
      }
    }
  }
}
```

#### File: `components/card.json`

```json
{
  "components": {
    "card": {
      "base": {
        "background-color": "semantic.colors.background",
        "border": "primitives.borders.border-width-thin solid semantic.colors.border-subtle",
        "border-radius": "semantic.interactive.border-radius-lg",
        "padding": "primitives.spacing.lg",
        "box-shadow": "semantic.interactive.shadow-sm"
      },
      
      "elevated": {
        "box-shadow": "semantic.interactive.shadow-md"
      },
      
      "header": {
        "padding-bottom": "primitives.spacing.md",
        "border-bottom": "primitives.borders.border-width-thin solid semantic.colors.border-subtle",
        "margin-bottom": "primitives.spacing.md"
      },
      
      "footer": {
        "padding-top": "primitives.spacing.md",
        "border-top": "primitives.borders.border-width-thin solid semantic.colors.border-subtle",
        "margin-top": "primitives.spacing.md"
      }
    }
  }
}
```

### Theme Composition Utility

**File: `pkg/theme/loader.go`**

```go
package theme

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    
    "github.com/niiniyare/ruun/internal/schema"
    "gopkg.in/yaml.v3"
)

// ThemeLoader handles loading and composing multi-file themes
type ThemeLoader struct {
    basePath string
}

// NewThemeLoader creates a new theme loader
func NewThemeLoader(basePath string) *ThemeLoader {
    return &ThemeLoader{basePath: basePath}
}

// LoadTheme loads and composes a theme from multiple files
func (tl *ThemeLoader) LoadTheme(themeName string) (*schema.Theme, error) {
    themeDir := filepath.Join(tl.basePath, themeName)
    
    // Load metadata
    metadata, err := tl.loadMetadata(themeDir)
    if err != nil {
        return nil, fmt.Errorf("failed to load metadata: %w", err)
    }
    
    // Load tokens
    tokens, err := tl.loadTokens(themeDir)
    if err != nil {
        return nil, fmt.Errorf("failed to load tokens: %w", err)
    }
    
    // Compose theme
    theme := &schema.Theme{
        ID:            metadata.ID,
        Name:          metadata.Name,
        Description:   metadata.Description,
        Version:       metadata.Version,
        Author:        metadata.Author,
        Tokens:        tokens,
        DarkMode:      metadata.DarkMode,
        Accessibility: metadata.Accessibility,
        Conditions:    metadata.Conditions,
        Meta:          metadata.Meta,
    }
    
    return theme, nil
}

// ThemeMetadata represents theme.json structure
type ThemeMetadata struct {
    ID            string                      `json:"id"`
    Name          string                      `json:"name"`
    Description   string                      `json:"description"`
    Version       string                      `json:"version"`
    Author        string                      `json:"author"`
    DarkMode      *schema.DarkModeConfig      `json:"darkMode,omitempty"`
    Accessibility *schema.AccessibilityConfig `json:"accessibility,omitempty"`
    Conditions    []*schema.ThemeCondition    `json:"conditions,omitempty"`
    Meta          *schema.ThemeMeta           `json:"meta,omitempty"`
}

// loadMetadata loads theme.json
func (tl *ThemeLoader) loadMetadata(themeDir string) (*ThemeMetadata, error) {
    path := filepath.Join(themeDir, "theme.json")
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    
    var metadata ThemeMetadata
    if err := json.Unmarshal(data, &metadata); err != nil {
        return nil, err
    }
    
    return &metadata, nil
}

// loadTokens loads and composes all token files
func (tl *ThemeLoader) loadTokens(themeDir string) (*schema.DesignTokens, error) {
    tokens := &schema.DesignTokens{
        Primitives: &schema.PrimitiveTokens{},
        Semantic:   &schema.SemanticTokens{},
        Components: &schema.ComponentTokens{},
    }
    
    // Load primitives (tokens.json)
    if err := tl.loadPrimitives(themeDir, tokens); err != nil {
        return nil, err
    }
    
    // Load semantic (semantic.json)
    if err := tl.loadSemantic(themeDir, tokens); err != nil {
        return nil, err
    }
    
    // Load components (components/*.json)
    if err := tl.loadComponents(themeDir, tokens); err != nil {
        return nil, err
    }
    
    return tokens, nil
}

// loadPrimitives loads tokens.json
func (tl *ThemeLoader) loadPrimitives(themeDir string, tokens *schema.DesignTokens) error {
    path := filepath.Join(themeDir, "tokens.json")
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    
    var temp struct {
        Primitives *schema.PrimitiveTokens `json:"primitives"`
    }
    
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }
    
    tokens.Primitives = temp.Primitives
    return nil
}

// loadSemantic loads semantic.json
func (tl *ThemeLoader) loadSemantic(themeDir string, tokens *schema.DesignTokens) error {
    path := filepath.Join(themeDir, "semantic.json")
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    
    var temp struct {
        Semantic *schema.SemanticTokens `json:"semantic"`
    }
    
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }
    
    tokens.Semantic = temp.Semantic
    return nil
}

// loadComponents loads all component/*.json files
func (tl *ThemeLoader) loadComponents(themeDir string, tokens *schema.DesignTokens) error {
    componentsDir := filepath.Join(themeDir, "components")
    
    files, err := filepath.Glob(filepath.Join(componentsDir, "*.json"))
    if err != nil {
        return err
    }
    
    for _, file := range files {
        if err := tl.loadComponentFile(file, tokens); err != nil {
            return fmt.Errorf("failed to load %s: %w", file, err)
        }
    }
    
    return nil
}

// loadComponentFile loads a single component file
func (tl *ThemeLoader) loadComponentFile(path string, tokens *schema.DesignTokens) error {
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    
    var temp struct {
        Components schema.ComponentTokens `json:"components"`
    }
    
    if err := json.Unmarshal(data, &temp); err != nil {
        return err
    }
    
    // Merge into existing components
    for component, variants := range temp.Components {
        (*tokens.Components)[component] = variants
    }
    
    return nil
}

// LoadThemeYAML loads theme from YAML files
func (tl *ThemeLoader) LoadThemeYAML(themeName string) (*schema.Theme, error) {
    // Similar implementation but using yaml.Unmarshal
    // ... (implementation follows same pattern as JSON)
    return nil, nil
}
```

**Usage:**

```go
package main

import (
    "log"
    "github.com/niiniyare/ruun/pkg/theme"
    "github.com/niiniyare/ruun/internal/schema"
)

func main() {
    // Load multi-file theme
    loader := theme.NewThemeLoader("themes")
    acmeTheme, err := loader.LoadTheme("acme-brand")
    if err != nil {
        log.Fatal(err)
    }
    
    // Register with theme manager
    manager, _ := schema.NewThemeManager(schema.DefaultThemeManagerConfig())
    err = manager.RegisterTheme(ctx, acmeTheme)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Option 2: Single-File Theme (Simpler for Small Themes)

For smaller themes, you can still use a single file:

```
themes/
  simple-theme.json
  another-theme.yaml
```

Just use the standard `ThemeFromJSON()` or `ThemeFromYAML()`:

```go
theme, err := schema.ThemeFromJSON(data)
```

---

## Getting Started

### 1. Initialize Theme System

```go
package main

import (
    "context"
    "log"
    
    "github.com/niiniyare/ruun/internal/schema"
    "github.com/niiniyare/ruun/pkg/theme"
)

func main() {
    // Initialize theme manager
    config := schema.DefaultThemeManagerConfig()
    config.EnableCaching = true
    config.ValidateOnRegister = true
    
    manager, err := schema.NewThemeManager(config)
    if err != nil {
        log.Fatal(err)
    }
    defer manager.Close()
    
    // Load and register themes
    loader := theme.NewThemeLoader("themes")
    
    // Load default theme
    defaultTheme, err := loader.LoadTheme("acme-brand")
    if err != nil {
        log.Fatal(err)
    }
    
    ctx := context.Background()
    err = manager.RegisterTheme(ctx, defaultTheme)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Println("Theme system initialized")
}
```

### 2. Generate CSS from Theme

```go
// pkg/theme/compiler.go
package theme

import (
    "context"
    "fmt"
    "strings"
    
    "github.com/niiniyare/ruun/internal/schema"
)

// CSSCompiler compiles themes to CSS
type CSSCompiler struct {
    resolver *schema.TokenResolver
}

// NewCSSCompiler creates a new CSS compiler
func NewCSSCompiler(theme *schema.Theme) (*CSSCompiler, error) {
    resolver, err := schema.NewTokenResolver(theme.Tokens)
    if err != nil {
        return nil, err
    }
    
    return &CSSCompiler{
        resolver: resolver,
    }, nil
}

// Compile compiles theme to CSS
func (cc *CSSCompiler) Compile(ctx context.Context) (string, error) {
    var css strings.Builder
    
    // Root variables
    css.WriteString(":root {\n")
    
    // Compile primitives
    if err := cc.compilePrimitives(&css); err != nil {
        return "", err
    }
    
    // Compile semantic
    if err := cc.compileSemantic(ctx, &css); err != nil {
        return "", err
    }
    
    css.WriteString("}\n\n")
    
    // Compile component classes
    if err := cc.compileComponents(ctx, &css); err != nil {
        return "", err
    }
    
    return css.String(), nil
}

// compilePrimitives compiles primitive tokens to CSS variables
func (cc *CSSCompiler) compilePrimitives(css *strings.Builder) error {
    css.WriteString("  /* Primitive Tokens */\n")
    
    primitives := cc.resolver.tokens.Primitives
    
    // Colors
    for key, value := range primitives.Colors {
        css.WriteString(fmt.Sprintf("  --color-%s: %s;\n", key, value))
    }
    
    // Spacing
    for key, value := range primitives.Spacing {
        css.WriteString(fmt.Sprintf("  --spacing-%s: %s;\n", key, value))
    }
    
    // Typography
    for key, value := range primitives.Typography {
        css.WriteString(fmt.Sprintf("  --%s: %s;\n", key, value))
    }
    
    // Shadows
    for key, value := range primitives.Shadows {
        css.WriteString(fmt.Sprintf("  --shadow-%s: %s;\n", key, value))
    }
    
    // Radius
    for key, value := range primitives.Radius {
        css.WriteString(fmt.Sprintf("  --radius-%s: %s;\n", key, value))
    }
    
    css.WriteString("\n")
    return nil
}

// compileSemantic compiles semantic tokens to CSS variables
func (cc *CSSCompiler) compileSemantic(ctx context.Context, css *strings.Builder) error {
    css.WriteString("  /* Semantic Tokens */\n")
    
    semantic := cc.resolver.tokens.Semantic
    
    // Resolve and write semantic colors
    for key, value := range semantic.Colors {
        ref := schema.TokenReference(value)
        resolved, err := cc.resolver.Resolve(ctx, ref)
        if err != nil {
            return err
        }
        css.WriteString(fmt.Sprintf("  --semantic-%s: %s;\n", key, resolved))
    }
    
    css.WriteString("\n")
    return nil
}

// compileComponents compiles component tokens to CSS classes
func (cc *CSSCompiler) compileComponents(ctx context.Context, css *strings.Builder) error {
    components := cc.resolver.tokens.Components
    
    for component, variants := range *components {
        css.WriteString(fmt.Sprintf("/* %s Component */\n", strings.Title(component)))
        
        // Base class
        css.WriteString(fmt.Sprintf(".%s {\n", component))
        css.WriteString("  /* Base styles applied to all variants */\n")
        
        // Check if "base" variant exists
        if baseProps, hasBase := variants["base"]; hasBase {
            for prop, value := range baseProps {
                resolved, err := cc.resolveValue(ctx, value)
                if err != nil {
                    return err
                }
                css.WriteString(fmt.Sprintf("  %s: %s;\n", prop, resolved))
            }
        }
        css.WriteString("}\n\n")
        
        // Variant classes
        for variant, props := range variants {
            if variant == "base" {
                continue // Already handled
            }
            
            className := fmt.Sprintf("%s-%s", component, variant)
            css.WriteString(fmt.Sprintf(".%s {\n", className))
            
            for prop, value := range props {
                resolved, err := cc.resolveValue(ctx, value)
                if err != nil {
                    return err
                }
                css.WriteString(fmt.Sprintf("  %s: %s;\n", prop, resolved))
            }
            
            css.WriteString("}\n\n")
        }
    }
    
    return nil
}

// resolveValue resolves a token reference or returns literal
func (cc *CSSCompiler) resolveValue(ctx context.Context, value string) (string, error) {
    ref := schema.TokenReference(value)
    return cc.resolver.Resolve(ctx, ref)
}

// Close releases resources
func (cc *CSSCompiler) Close() error {
    return cc.resolver.Close()
}
```

**Usage:**

```go
func GenerateThemeCSS(theme *schema.Theme) (string, error) {
    compiler, err := theme.NewCSSCompiler(theme)
    if err != nil {
        return "", err
    }
    defer compiler.Close()
    
    ctx := context.Background()
    css, err := compiler.Compile(ctx)
    if err != nil {
        return "", err
    }
    
    return css, nil
}
```

---

## Component Integration

### Atom: Button Component

**File: `views/atoms/button.templ`**

```go
package atoms

import (
    "fmt"
    "github.com/niiniyare/ruun/pkg/utils"
)

type ButtonVariant string

const (
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonOutline     ButtonVariant = "outline"
    ButtonDestructive ButtonVariant = "destructive"
    ButtonGhost       ButtonVariant = "ghost"
    ButtonLink        ButtonVariant = "link"
)

type ButtonSize string

const (
    ButtonSizeXS ButtonSize = "xs"
    ButtonSizeSM ButtonSize = "sm"
    ButtonSizeMD ButtonSize = "md"
    ButtonSizeLG ButtonSize = "lg"
    ButtonSizeXL ButtonSize = "xl"
)

type ButtonProps struct {
    Variant   ButtonVariant
    Size      ButtonSize
    Text      string
    Icon      string
    IconLeft  string
    IconRight string
    Type      string
    Disabled  bool
    Loading   bool
    ClassName string
    
    // HTMX attributes
    HXPost    string
    HXGet     string
    HXTarget  string
    HXSwap    string
    
    // Alpine.js
    AlpineClick string
}

// buttonClasses generates CSS classes from compiled theme
func buttonClasses(props ButtonProps) string {
    return utils.TwMerge(
        // Base class from theme: .button
        "button",
        
        // Variant class from theme: .button-{variant}
        fmt.Sprintf("button-%s", props.Variant),
        
        // Size class from theme: .button-{size}
        fmt.Sprintf("button-%s", props.Size),
        
        // State classes from theme
        utils.If(props.Loading, "button-loading"),
        utils.If(props.Disabled, "button-disabled"),
        
        // Custom classes
        props.ClassName,
    )
}

templ Button(props ButtonProps) {
    <button
        type={ utils.IfElse(props.Type != "", props.Type, "button") }
        class={ buttonClasses(props) }
        if props.Disabled {
            disabled
        }
        if props.HXPost != "" {
            hx-post={ props.HXPost }
        }
        if props.HXGet != "" {
            hx-get={ props.HXGet }
        }
        if props.HXTarget != "" {
            hx-target={ props.HXTarget }
        }
        if props.HXSwap != "" {
            hx-swap={ props.HXSwap }
        }
        if props.AlpineClick != "" {
            x-on:click={ props.AlpineClick }
        }
    >
        if props.Loading {
            <svg class="spinner" viewBox="0 0 24 24">
                <circle cx="12" cy="12" r="10" stroke="currentColor"></circle>
            </svg>
        }
        
        if !props.Loading && props.IconLeft != "" {
            @Icon(IconProps{Name: props.IconLeft, Size: "sm"})
        }
        
        if props.Text != "" {
            { props.Text }
        }
        
        if !props.Loading && props.IconRight != "" {
            @Icon(IconProps{Name: props.IconRight, Size: "sm"})
        }
    </button>
}
```

### Molecule: Form Field

**File: `views/molecules/form_field.templ`**

```go
package molecules

import (
    "github.com/niiniyare/ruun/views/atoms"
)

type FormFieldProps struct {
    Label       string
    Name        string
    Type        string
    Placeholder string
    Required    bool
    Error       string
    HelpText    string
}

templ FormField(props FormFieldProps) {
    <div class="form-field">
        // Label atom (uses semantic typography tokens)
        <label for={ props.Name } class="form-label">
            { props.Label }
            if props.Required {
                <span class="form-required">*</span>
            }
        </label>
        
        // Input atom (uses .input-base from theme)
        @atoms.Input(atoms.InputProps{
            ID:          props.Name,
            Name:        props.Name,
            Type:        props.Type,
            Placeholder: props.Placeholder,
            Required:    props.Required,
            ClassName:   utils.If(props.Error != "", "input-error"),
        })
        
        // Help text (uses semantic typography)
        if props.HelpText != "" {
            <p class="form-help-text">{ props.HelpText }</p>
        }
        
        // Error message (uses semantic colors)
        if props.Error != "" {
            <p class="form-error-text">{ props.Error }</p>
        }
    </div>
}
```

### Organism: Login Form

**File: `views/organisms/login_form.templ`**

```go
package organisms

import (
    "github.com/niiniyare/ruun/views/atoms"
    "github.com/niiniyare/ruun/views/molecules"
)

type LoginFormProps struct {
    Action string
    Error  string
}

templ LoginForm(props LoginFormProps) {
    <div class="card card-elevated login-form">
        <div class="card-header">
            <h2 class="card-title">Login</h2>
        </div>
        
        <form 
            hx-post={ props.Action }
            hx-target="#login-response"
            hx-swap="outerHTML"
        >
            // Email field (molecule)
            @molecules.FormField(molecules.FormFieldProps{
                Label:       "Email",
                Name:        "email",
                Type:        "email",
                Placeholder: "you@example.com",
                Required:    true,
            })
            
            // Password field (molecule)
            @molecules.FormField(molecules.FormFieldProps{
                Label:    "Password",
                Name:     "password",
                Type:     "password",
                Required: true,
            })
            
            // Error display
            if props.Error != "" {
                <div class="alert alert-error" role="alert">
                    { props.Error }
                </div>
            }
            
            // Submit button (atom)
            @atoms.Button(atoms.ButtonProps{
                Variant: atoms.ButtonPrimary,
                Size:    atoms.ButtonSizeMD,
                Text:    "Sign In",
                Type:    "submit",
            })
        </form>
        
        <div id="login-response"></div>
    </div>
}
```

---

## CSS Compilation

### Generated CSS Output

From the theme tokens, the compiler generates:

**1. CSS Custom Properties (Variables)**

```css
:root {
  /* Primitives */
  --color-brand-primary: hsl(210, 100%, 45%);
  --color-gray-50: hsl(210, 20%, 98%);
  --color-gray-900: hsl(210, 20%, 14%);
  --spacing-sm: 0.75rem;
  --spacing-md: 1rem;
  --radius-md: 0.5rem;
  --font-size-base: 1rem;
  --font-weight-medium: 500;
  --shadow-sm: 0 2px 4px 0 rgba(0, 0, 0, 0.06);
  
  /* Semantic */
  --semantic-background: hsl(0, 0%, 100%);
  --semantic-foreground: hsl(210, 20%, 14%);
  --semantic-primary: hsl(210, 100%, 45%);
  --semantic-primary-foreground: #ffffff;
  --semantic-border: hsl(210, 14%, 89%);
}
```

**2. Component Classes**

```css
/* Button Component */
.button {
  /* Base styles applied to all variants */
  font-size: 1rem;
  font-weight: 500;
  line-height: 1.5;
  border-width: 0;
  cursor: pointer;
  transition: all 150ms cubic-bezier(0.4, 0, 0.2, 1);
}

.button-primary {
  background-color: hsl(210, 100%, 45%);
  color: #ffffff;
  border: none;
}

.button-secondary {
  background-color: hsl(340, 82%, 52%);
  color: #ffffff;
  border: none;
}

.button-outline {
  background-color: transparent;
  color: hsl(210, 100%, 45%);
  border: 1px solid hsl(210, 100%, 45%);
}

.button-destructive {
  background-color: hsl(0, 84%, 60%);
  color: #ffffff;
  border: none;
}

.button-ghost {
  background-color: transparent;
  color: hsl(210, 100%, 45%);
  border: none;
}

.button-xs {
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
}

.button-sm {
  padding: 0.5rem 0.75rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
}

.button-md {
  padding: 0.75rem 1rem;
  border-radius: 0.5rem;
  font-size: 1rem;
}

.button-lg {
  padding: 0.75rem 1.5rem;
  border-radius: 0.5rem;
  font-size: 1.125rem;
}

.button-xl {
  padding: 1rem 2rem;
  border-radius: 0.75rem;
  font-size: 1.25rem;
}

.button-disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.button-loading {
  opacity: 0.7;
  cursor: wait;
}

/* Input Component */
.input {
  background-color: hsl(0, 0%, 100%);
  border: 1px solid hsl(210, 14%, 89%);
  border-radius: 0.5rem;
  padding: 0.75rem;
  color: hsl(210, 20%, 14%);
  font-size: 1rem;
  line-height: 1.5;
  transition: all 150ms;
}

.input:focus {
  border-color: hsl(210, 100%, 45%);
  outline: 2px solid hsl(210, 100%, 45%);
  outline-offset: 2px;
}

.input-error {
  border-color: hsl(0, 84%, 60%);
}

.input:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  background-color: hsl(210, 20%, 95%);
}

/* Badge Component */
.badge {
  display: inline-flex;
  align-items: center;
  padding: 0.125rem 0.5rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  font-weight: 500;
  line-height: 1;
}

.badge-default {
  background-color: hsl(210, 20%, 95%);
  color: hsl(210, 10%, 64%);
}

.badge-primary {
  background-color: hsl(210, 100%, 45%);
  color: #ffffff;
}

.badge-success {
  background-color: hsl(142, 71%, 45%);
  color: #ffffff;
}

.badge-error {
  background-color: hsl(0, 84%, 60%);
  color: #ffffff;
}

/* Card Component */
.card {
  background-color: hsl(0, 0%, 100%);
  border: 1px solid hsl(210, 16%, 93%);
  border-radius: 0.75rem;
  padding: 1.5rem;
  box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.06);
}

.card-elevated {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}
```

**3. Dark Mode**

```css
.dark,
@media (prefers-color-scheme: dark) {
  :root {
    --semantic-background: hsl(210, 20%, 14%);
    --semantic-foreground: hsl(210, 20%, 98%);
    --semantic-border: hsl(210, 12%, 36%);
  }
  
  .button-primary {
    background-color: hsl(210, 100%, 55%);
  }
  
  .card {
    background-color: hsl(210, 20%, 14%);
    border-color: hsl(210, 12%, 36%);
  }
}
```

### HTTP Handler for CSS

```go
package handlers

import (
    "context"
    "net/http"
    
    "github.com/niiniyare/ruun/internal/schema"
    "github.com/niiniyare/ruun/pkg/theme"
)

type ThemeHandler struct {
    manager *schema.ThemeManager
}

func NewThemeHandler(manager *schema.ThemeManager) *ThemeHandler {
    return &ThemeHandler{manager: manager}
}

// ServeCSS generates and serves theme CSS
func (h *ThemeHandler) ServeCSS(w http.ResponseWriter, r *http.Request) {
    // Extract tenant from request
    tenantID := r.Header.Get("X-Tenant-ID")
    
    // Create context
    ctx := schema.WithTenant(r.Context(), tenantID)
    
    // Get theme
    themeObj, err := h.manager.GetTheme(ctx, "default", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Compile to CSS
    compiler, err := theme.NewCSSCompiler(themeObj)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer compiler.Close()
    
    css, err := compiler.Compile(context.Background())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    // Serve CSS with caching headers
    w.Header().Set("Content-Type", "text/css")
    w.Header().Set("Cache-Control", "public, max-age=3600")
    w.Write([]byte(css))
}
```

**Route Setup:**

```go
func SetupRoutes(mux *http.ServeMux, manager *schema.ThemeManager) {
    themeHandler := handlers.NewThemeHandler(manager)
    mux.HandleFunc("/css/theme.css", themeHandler.ServeCSS)
}
```

**HTML Template:**

```html
<!DOCTYPE html>
<html>
<head>
    <link rel="stylesheet" href="/css/theme.css">
</head>
<body>
    <!-- Components automatically use theme classes -->
</body>
</html>
```

---

## Workflows

### Complete Application Workflow

```go
package main

import (
    "context"
    "log"
    "net/http"
    
    "github.com/niiniyare/ruun/internal/schema"
    "github.com/niiniyare/ruun/pkg/theme"
    "github.com/niiniyare/ruun/handlers"
)

func main() {
    // 1. Initialize theme manager
    config := schema.DefaultThemeManagerConfig()
    config.EnableCaching = true
    manager, err := schema.NewThemeManager(config)
    if err != nil {
        log.Fatal(err)
    }
    defer manager.Close()
    
    // 2. Load themes
    loader := theme.NewThemeLoader("themes")
    defaultTheme, err := loader.LoadTheme("acme-brand")
    if err != nil {
        log.Fatal(err)
    }
    
    // 3. Register default theme
    ctx := context.Background()
    err = manager.RegisterTheme(ctx, defaultTheme)
    if err != nil {
        log.Fatal(err)
    }
    
    // 4. Setup HTTP routes
    mux := http.NewServeMux()
    
    // Theme CSS endpoint
    themeHandler := handlers.NewThemeHandler(manager)
    mux.HandleFunc("/css/theme.css", themeHandler.ServeCSS)
    
    // Application routes
    mux.HandleFunc("/", handlers.HomePage)
    mux.HandleFunc("/login", handlers.LoginPage)
    
    // 5. Start server
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

---

## Best Practices

### 1. Theme Organization

✅ **DO:** Organize by concern
```
themes/
  acme-brand/
    theme.json       # Metadata
    tokens.json      # Primitives
    semantic.json    # Semantic
    components/      # Components
```

✅ **DO:** Use semantic naming
```json
{
  "semantic": {
    "colors": {
      "background": "...",
      "foreground": "...",
      "primary": "..."
    }
  }
}
```

❌ **DON'T:** Mix concerns
```
themes/
  everything.json  # Everything in one file (hard to maintain)
```

### 2. Component Classes

✅ **DO:** Use BEM-like naming
```css
.button { }
.button-primary { }
.button-sm { }
.button-loading { }
```

✅ **DO:** Keep components atomic
```go
// Atom: Single purpose
templ Button(props ButtonProps) { }

// Molecule: Composition
templ FormField() {
  @Label()
  @Input()
  @ErrorMessage()
}
```

❌ **DON'T:** Create monolithic components
```go
// Bad: Too much responsibility
templ FormWithEverything() { }
```

### 3. Token References

✅ **DO:** Reference down the hierarchy
```json
{
  "components": {
    "button": {
      "primary": {
        "background-color": "semantic.colors.primary"
      }
    }
  }
}
```

❌ **DON'T:** Create circular references
```json
{
  "semantic": {
    "primary": "components.button.primary.background-color"
  }
}
```

---

## Advanced Patterns

### 1. Dynamic Theme Switching

```go
// Middleware to detect and apply theme
func ThemeMiddleware(manager *schema.ThemeManager) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            tenantID := r.Header.Get("X-Tenant-ID")
            userRole := r.Header.Get("X-User-Role")
            
            // Create context with theme info
            ctx := schema.WithTenant(r.Context(), tenantID)
            ctx = context.WithValue(ctx, "userRole", userRole)
            
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}
```

### 2. Component Theme Variants

```go
// Variant factory for themed components
func ThemedButton(themeName string) func(ButtonProps) templ.Component {
    return func(props ButtonProps) templ.Component {
        // Add theme prefix to classes
        props.ClassName = fmt.Sprintf("theme-%s %s", themeName, props.ClassName)
        return Button(props)
    }
}

// Usage
acmeButton := ThemedButton("acme")
@acmeButton(ButtonProps{Variant: ButtonPrimary})
```

### 3. Real-time Theme Updates

```go
// SSE endpoint for theme updates
func (h *ThemeHandler) StreamThemeUpdates(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    
    // Send initial theme
    theme, _ := h.manager.GetTheme(r.Context(), "default", nil)
    css, _ := generateCSS(theme)
    fmt.Fprintf(w, "data: %s\n\n", css)
    
    // Listen for updates
    updates := make(chan string)
    // ... subscribe to theme updates
    
    for update := range updates {
        fmt.Fprintf(w, "data: %s\n\n", update)
        w.(http.Flusher).Flush()
    }
}
```

---

## Summary

This guide covered:

✅ **Atomic Design Integration**: Atoms → Molecules → Organisms → Templates → Pages  
✅ **Token-to-Component Flow**: Tokens → CSS → Props → Components  
✅ **Multi-File Themes**: Organized by primitives, semantic, components  
✅ **CSS Compilation**: Automatic generation from tokens  
✅ **templ Components**: Pure presentation with token-based styling  
✅ **Complete Workflows**: From theme loading to rendering  

**Key Principles:**
1. **Tokens define styles** (single source of truth)
2. **CSS classes apply tokens** (compiled once)
3. **Components use classes** (through props)
4. **Atomic Design organizes** (clear hierarchy)
5. **Multi-tenancy isolates** (per-tenant themes)

For more information:
- [schema Documentation: ](./docs/schema/) 
- [Component Library: ](./docs/components/README.md)
- [Theme Examples: ](../../views/style/themes/default.json) 
