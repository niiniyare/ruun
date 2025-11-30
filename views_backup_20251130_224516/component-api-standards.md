# AWO ERP Component API Standards

## Overview

This document defines the API standards and patterns for all templ components in the AWO ERP system. These standards ensure consistency, type safety, and maintainability across the component library.

## Core Principles

1. **Props-First Architecture**: All component behavior is controlled through typed props
2. **Zero Business Logic**: Components are pure presentation layers
3. **Type Safety**: Use Go's type system for compile-time guarantees
4. **Composition Over Configuration**: Build complex UIs from simple, composable parts
5. **Semantic Classes**: Use Basecoat CSS utilities for consistent theming

## Component API Patterns

### 1. Props Structure

Every component must have a typed props struct:

```go
type ButtonProps struct {
    // Core properties
    Label    string
    Variant  ButtonVariant
    Size     ButtonSize
    Disabled bool
    Loading  bool
    Type     string // "button" | "submit" | "reset"
    
    // HTMX/Alpine attributes
    Attrs    templ.Attributes
    
    // Event handlers
    OnClick  string // Alpine.js expression
    
    // Layout override (use sparingly)
    Layout   string // For utils.TwMerge
}
```

### 2. Variant Types

Use typed enums for all variants:

```go
type ButtonVariant string
const (
    ButtonDefault     ButtonVariant = "default"
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonOutline     ButtonVariant = "outline"
    ButtonDestructive ButtonVariant = "destructive"
    ButtonGhost       ButtonVariant = "ghost"
)

type ButtonSize string
const (
    ButtonXS ButtonSize = "xs"
    ButtonSM ButtonSize = "sm"
    ButtonMD ButtonSize = "md"
    ButtonLG ButtonSize = "lg"
    ButtonXL ButtonSize = "xl"
)
```

### 3. Class Composition

Use `templ.Classes()` with `templ.KV()` for conditional classes:

```templ
templ Button(props ButtonProps) {
    <button
        type={ props.Type }
        disabled?={ props.Disabled || props.Loading }
        @click={ props.OnClick }
        class={ templ.Classes(
            // Base classes (always applied)
            "btn transition-colors focus-visible:outline-ring",
            
            // Variant classes
            templ.KV("btn-default", props.Variant == ButtonDefault),
            templ.KV("btn-primary", props.Variant == ButtonPrimary),
            templ.KV("btn-secondary", props.Variant == ButtonSecondary),
            templ.KV("btn-outline", props.Variant == ButtonOutline),
            templ.KV("btn-destructive", props.Variant == ButtonDestructive),
            templ.KV("btn-ghost", props.Variant == ButtonGhost),
            
            // Size classes
            templ.KV("btn-xs", props.Size == ButtonXS),
            templ.KV("btn-sm", props.Size == ButtonSM),
            templ.KV("btn-md", props.Size == ButtonMD),
            templ.KV("btn-lg", props.Size == ButtonLG),
            templ.KV("btn-xl", props.Size == ButtonXL),
            
            // State classes
            templ.KV("opacity-50 cursor-not-allowed", props.Disabled),
            
            // Layout override (if needed)
            utils.TwMerge("", props.Layout),
        ) }
        { props.Attrs... }
    >
        if props.Loading {
            @Spinner(SpinnerProps{Size: "sm", ClassName: "mr-2"})
        }
        { props.Label }
    </button>
}
```

### 4. Attribute Spreading

Use `templ.Attributes` for HTMX and ARIA attributes:

```go
// In handler/controller
buttonProps := ButtonProps{
    Label:   "Delete Invoice",
    Variant: ButtonDestructive,
    Attrs: templ.Attributes{
        "hx-delete":  "/api/invoices/123",
        "hx-confirm": "Are you sure?",
        "hx-target":  "#invoice-list",
        "aria-label": "Delete invoice #123",
    },
}
```

### 5. Component Slots

Use `templ.Component` for flexible composition:

```go
type CardProps struct {
    Title       string
    Description string
    Header      templ.Component // Optional slot
    Footer      templ.Component // Optional slot
    Children    templ.Component // Main content slot
}

templ Card(props CardProps) {
    <div class="card">
        if props.Header != nil {
            <div class="card-header">
                { props.Header... }
            </div>
        } else if props.Title != "" {
            <div class="card-header">
                <h3 class="card-title">{ props.Title }</h3>
                if props.Description != "" {
                    <p class="card-description">{ props.Description }</p>
                }
            </div>
        }
        
        <div class="card-content">
            { props.Children... }
        </div>
        
        if props.Footer != nil {
            <div class="card-footer">
                { props.Footer... }
            </div>
        }
    </div>
}
```

### 6. Layout Conflicts

Use `utils.TwMerge` when accepting layout overrides:

```go
type FlexProps struct {
    Direction string // "row" | "col"
    Gap       string // "1" | "2" | "4" | "8"
    Layout    string // Override for special cases
    Children  templ.Component
}

templ Flex(props FlexProps) {
    <div class={ utils.TwMerge(
        fmt.Sprintf("flex flex-%s gap-%s", props.Direction, props.Gap),
        props.Layout,
    ) }>
        { props.Children... }
    </div>
}
```

## Basecoat CSS Classes

Always use Basecoat semantic utility classes:

### Colors
- `btn`, `btn-primary`, `btn-secondary`, `btn-outline`, `btn-destructive`, `btn-ghost`
- `card`, `card-header`, `card-content`, `card-footer`
- `input`, `input-error`
- `text-foreground`, `text-muted`, `text-primary`, `text-destructive`
- `bg-background`, `bg-primary`, `bg-secondary`, `bg-muted`
- `border-border`, `border-primary`, `border-destructive`

### Layout
- `container`, `prose`
- `flex`, `grid`, `inline-flex`
- `gap-{1,2,4,8}`, `space-x-{1,2,4,8}`, `space-y-{1,2,4,8}`

### Typography
- `text-xs`, `text-sm`, `text-base`, `text-lg`, `text-xl`, `text-2xl`
- `font-normal`, `font-medium`, `font-semibold`, `font-bold`

### Interactive
- `hover:bg-muted`, `focus:ring-2`, `focus:ring-primary`
- `disabled:opacity-50`, `disabled:cursor-not-allowed`
- `transition-colors`, `transition-all`

## Component Examples

### Form Field

```go
type FormFieldProps struct {
    Label       string
    Name        string
    Type        string
    Value       string
    Placeholder string
    Required    bool
    Disabled    bool
    Error       string
    HelpText    string
    Attrs       templ.Attributes
}

templ FormField(props FormFieldProps) {
    <div class="space-y-2">
        <label 
            for={ props.Name }
            class={ templ.Classes(
                "text-sm font-medium",
                templ.KV("text-destructive", props.Error != ""),
            ) }
        >
            { props.Label }
            if props.Required {
                <span class="text-destructive ml-1">*</span>
            }
        </label>
        
        <input
            id={ props.Name }
            name={ props.Name }
            type={ props.Type }
            value={ props.Value }
            placeholder={ props.Placeholder }
            disabled?={ props.Disabled }
            required?={ props.Required }
            class={ templ.Classes(
                "input w-full",
                templ.KV("input-error", props.Error != ""),
            ) }
            aria-invalid={ fmt.Sprint(props.Error != "") }
            aria-describedby={ props.Name + "-error" }
            { props.Attrs... }
        />
        
        if props.Error != "" {
            <p id={ props.Name + "-error" } class="text-sm text-destructive">
                { props.Error }
            </p>
        } else if props.HelpText != "" {
            <p class="text-sm text-muted">
                { props.HelpText }
            </p>
        }
    </div>
}
```

### Data Table

```go
type DataTableProps struct {
    Columns     []Column
    Rows        []map[string]any
    EmptyText   string
    Attrs       templ.Attributes
}

type Column struct {
    Key      string
    Label    string
    Sortable bool
    Format   func(any) string
}

templ DataTable(props DataTableProps) {
    <div class="table-container" { props.Attrs... }>
        <table class="table">
            <thead>
                <tr>
                    for _, col := range props.Columns {
                        <th class={ templ.Classes(
                            "table-header",
                            templ.KV("cursor-pointer hover:bg-muted", col.Sortable),
                        ) }>
                            { col.Label }
                            if col.Sortable {
                                @Icon(IconProps{Name: "arrow-up-down", Size: "sm"})
                            }
                        </th>
                    }
                </tr>
            </thead>
            <tbody>
                if len(props.Rows) == 0 {
                    <tr>
                        <td colspan={ fmt.Sprint(len(props.Columns)) } class="table-cell text-center text-muted">
                            { props.EmptyText }
                        </td>
                    </tr>
                } else {
                    for _, row := range props.Rows {
                        <tr class="table-row">
                            for _, col := range props.Columns {
                                <td class="table-cell">
                                    if col.Format != nil {
                                        { col.Format(row[col.Key]) }
                                    } else {
                                        { fmt.Sprint(row[col.Key]) }
                                    }
                                </td>
                            }
                        </tr>
                    }
                }
            </tbody>
        </table>
    </div>
}
```

## Anti-Patterns to Avoid

### ❌ DON'T: Use ClassName Props

```go
// WRONG
type ButtonProps struct {
    ClassName string // Avoid this
}

// CORRECT
type ButtonProps struct {
    Layout string // For specific layout overrides only
}
```

### ❌ DON'T: Build Classes Dynamically

```templ
// WRONG
<div class={ fmt.Sprintf("bg-%s text-%s", color, size) }>

// CORRECT
<div class={ templ.Classes(
    templ.KV("bg-primary text-white", color == "primary"),
    templ.KV("bg-secondary text-black", color == "secondary"),
) }>
```

### ❌ DON'T: Mix Concerns

```go
// WRONG
type ButtonProps struct {
    UserID    int    // Business data
    CanEdit   bool   // Computed permission
    APIRoute  string // Dynamic URL
}

// CORRECT
type ButtonProps struct {
    Label    string
    Disabled bool
    Attrs    templ.Attributes // Pass URLs via attrs
}
```

## Checklist for New Components

- [ ] Props struct with typed fields
- [ ] Variant/Size enums (if applicable)
- [ ] Use `templ.Classes()` with `templ.KV()`
- [ ] Support `templ.Attributes` spreading
- [ ] Use `templ.Component` for slots
- [ ] Apply `utils.TwMerge` for layout conflicts
- [ ] Use Basecoat semantic classes
- [ ] No business logic in component
- [ ] No dynamic class/URL construction
- [ ] Proper ARIA attributes
- [ ] Alpine.js directives (if interactive)
- [ ] HTMX attributes via props.Attrs
- [ ] TypeScript-like naming conventions

## Migration Guide

When refactoring existing components:

1. Replace string concatenation with `templ.Classes()`
2. Convert className props to Layout props with `utils.TwMerge`
3. Move HTMX/ARIA to `templ.Attributes`
4. Extract variants to typed enums
5. Ensure all classes use Basecoat utilities
6. Remove any business logic or data fetching

## Testing Components

Components should be tested for:

1. **Prop Variations**: All variant/size combinations render correctly
2. **Conditional Rendering**: States appear/disappear as expected
3. **Attribute Spreading**: HTMX/ARIA attributes apply correctly
4. **Slot Composition**: Child components render in correct positions
5. **Accessibility**: ARIA attributes and keyboard navigation work

Remember: Components are pure presentation. They receive computed props and render UI. All logic happens before the component boundary.