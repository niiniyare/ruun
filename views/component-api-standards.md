# Component API Standards & Patterns

> Reference guide for implementing templ components with Basecoat UI
> 
> Generated for the UI Component Refactoring project

---

## Core Principles

### 1. Type-Safe Variants (No Magic Strings)

```go
// ❌ BAD - Magic strings, no type safety
templ Button(variant string) {
    <button class={ "btn btn-" + variant }>
        { children... }
    </button>
}

// ✅ GOOD - Type-safe variants
type ButtonVariant string

const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonOutline   ButtonVariant = "outline"
    ButtonGhost     ButtonVariant = "ghost"
    ButtonLink      ButtonVariant = "link"
)

templ Button(variant ButtonVariant) {
    <button class="btn" data-variant={ string(variant) }>
        { children... }
    </button>
}
```

### 2. Data Attributes for State (Not Utility Classes)

```go
// ❌ BAD - Class manipulation for state
templ Button(loading bool) {
    <button class={ templ.KV("opacity-50 cursor-wait", loading), "btn" }>
        { children... }
    </button>
}

// ✅ GOOD - Data attributes for state
templ Button(loading bool) {
    <button 
        class="btn" 
        data-loading={ strconv.FormatBool(loading) }
        aria-busy={ strconv.FormatBool(loading) }
    >
        { children... }
    </button>
}

// CSS handles the styling:
// .btn[data-loading="true"] { opacity: 0.5; cursor: wait; }
```

### 3. Props Struct for Complex Components

```go
// ❌ BAD - Too many parameters
templ Input(name, id, placeholder, value, inputType string, required, disabled, readonly bool, size Size) {
    // ...
}

// ✅ GOOD - Props struct
type InputProps struct {
    Name        string
    ID          string
    Placeholder string
    Value       string
    Type        string // "text", "email", "password", etc.
    Required    bool
    Disabled    bool
    Readonly    bool
    Size        Size
}

// Default values helper
func DefaultInputProps() InputProps {
    return InputProps{
        Type: "text",
        Size: SizeMd,
    }
}

templ Input(props InputProps) {
    <input
        class="input"
        type={ props.Type }
        name={ props.Name }
        id={ props.ID }
        placeholder={ props.Placeholder }
        value={ props.Value }
        required?={ props.Required }
        disabled?={ props.Disabled }
        readonly?={ props.Readonly }
        data-size={ string(props.Size) }
    />
}
```

### 4. Semantic Classes Only

```go
// ❌ BAD - Utility class soup
templ Card() {
    <div class="bg-white rounded-lg shadow-md p-6 border border-gray-200 hover:shadow-lg transition-shadow">
        { children... }
    </div>
}

// ✅ GOOD - Semantic Basecoat class
templ Card() {
    <div class="card">
        { children... }
    </div>
}

// Additional structure with semantic parts
templ CardWithHeader(title string) {
    <div class="card">
        <div class="card-header">
            <h3 class="card-title">{ title }</h3>
        </div>
        <div class="card-body">
            { children... }
        </div>
    </div>
}
```

---

## Common Types (Shared Across Components)

```go
// views/components/types.go

package components

// Size represents component sizes
type Size string

const (
    SizeXs Size = "xs"
    SizeSm Size = "sm"
    SizeMd Size = "md"
    SizeLg Size = "lg"
    SizeXl Size = "xl"
)

// Intent represents semantic meaning/color
type Intent string

const (
    IntentDefault Intent = "default"
    IntentPrimary Intent = "primary"
    IntentSuccess Intent = "success"
    IntentWarning Intent = "warning"
    IntentDanger  Intent = "danger"
    IntentInfo    Intent = "info"
)

// Position for tooltips, dropdowns, etc.
type Position string

const (
    PositionTop    Position = "top"
    PositionBottom Position = "bottom"
    PositionLeft   Position = "left"
    PositionRight  Position = "right"
)
```

---

## Atom Patterns

### Button

```go
// views/components/atoms/button.templ

package atoms

import "strconv"

type ButtonVariant string

const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonOutline   ButtonVariant = "outline"
    ButtonGhost     ButtonVariant = "ghost"
    ButtonLink      ButtonVariant = "link"
    ButtonDestructive ButtonVariant = "destructive"
)

type ButtonSize string

const (
    ButtonSm   ButtonSize = "sm"
    ButtonMd   ButtonSize = "md"
    ButtonLg   ButtonSize = "lg"
    ButtonIcon ButtonSize = "icon"
)

type ButtonProps struct {
    Variant  ButtonVariant
    Size     ButtonSize
    Type     string // "button", "submit", "reset"
    Disabled bool
    Loading  bool
}

func DefaultButtonProps() ButtonProps {
    return ButtonProps{
        Variant: ButtonPrimary,
        Size:    ButtonMd,
        Type:    "button",
    }
}

// Basic button with children
templ Button(props ButtonProps) {
    <button
        class="btn"
        type={ props.Type }
        data-variant={ string(props.Variant) }
        data-size={ string(props.Size) }
        data-loading={ strconv.FormatBool(props.Loading) }
        disabled?={ props.Disabled || props.Loading }
        aria-busy?={ props.Loading }
    >
        if props.Loading {
            <span class="btn-spinner" aria-hidden="true"></span>
        }
        <span class="btn-label">
            { children... }
        </span>
    </button>
}

// Icon-only button
templ IconButton(props ButtonProps, icon string, label string) {
    <button
        class="btn"
        type={ props.Type }
        data-variant={ string(props.Variant) }
        data-size="icon"
        disabled?={ props.Disabled }
        aria-label={ label }
    >
        @Icon(icon)
    </button>
}

// Button with HTMX
templ HTMXButton(props ButtonProps, hx HtmxAttrs) {
    <button
        class="btn"
        type={ props.Type }
        data-variant={ string(props.Variant) }
        data-size={ string(props.Size) }
        disabled?={ props.Disabled }
        hx-post?={ hx.Post }
        hx-get?={ hx.Get }
        hx-target?={ hx.Target }
        hx-swap?={ hx.Swap }
        hx-indicator?={ hx.Indicator }
    >
        { children... }
    </button>
}
```

### Input

```go
// views/components/atoms/input.templ

package atoms

type InputType string

const (
    InputText     InputType = "text"
    InputEmail    InputType = "email"
    InputPassword InputType = "password"
    InputNumber   InputType = "number"
    InputTel      InputType = "tel"
    InputUrl      InputType = "url"
    InputSearch   InputType = "search"
)

type InputProps struct {
    Name        string
    ID          string
    Type        InputType
    Placeholder string
    Value       string
    Required    bool
    Disabled    bool
    Readonly    bool
    Error       bool
    Size        Size
    // HTMX
    HxPost    string
    HxTrigger string
    HxTarget  string
}

func DefaultInputProps() InputProps {
    return InputProps{
        Type: InputText,
        Size: SizeMd,
    }
}

templ Input(props InputProps) {
    <input
        class="input"
        type={ string(props.Type) }
        name={ props.Name }
        id={ props.ID }
        placeholder={ props.Placeholder }
        value={ props.Value }
        required?={ props.Required }
        disabled?={ props.Disabled }
        readonly?={ props.Readonly }
        data-size={ string(props.Size) }
        data-error={ boolStr(props.Error) }
        aria-invalid?={ props.Error }
        hx-post?={ props.HxPost }
        hx-trigger?={ props.HxTrigger }
        hx-target?={ props.HxTarget }
    />
}
```

### Label

```go
// views/components/atoms/label.templ

package atoms

type LabelProps struct {
    For      string
    Required bool
}

templ Label(props LabelProps) {
    <label 
        class="label"
        for={ props.For }
        data-required={ boolStr(props.Required) }
    >
        { children... }
        if props.Required {
            <span class="label-required" aria-hidden="true">*</span>
        }
    </label>
}
```

### Badge

```go
// views/components/atoms/badge.templ

package atoms

type BadgeVariant string

const (
    BadgeDefault   BadgeVariant = "default"
    BadgePrimary   BadgeVariant = "primary"
    BadgeSecondary BadgeVariant = "secondary"
    BadgeSuccess   BadgeVariant = "success"
    BadgeWarning   BadgeVariant = "warning"
    BadgeDanger    BadgeVariant = "danger"
    BadgeOutline   BadgeVariant = "outline"
)

templ Badge(variant BadgeVariant) {
    <span class="badge" data-variant={ string(variant) }>
        { children... }
    </span>
}
```

### Alert

```go
// views/components/atoms/alert.templ

package atoms

type AlertVariant string

const (
    AlertInfo    AlertVariant = "info"
    AlertSuccess AlertVariant = "success"
    AlertWarning AlertVariant = "warning"
    AlertError   AlertVariant = "error"
)

type AlertProps struct {
    Variant     AlertVariant
    Dismissible bool
    Title       string
}

templ Alert(props AlertProps) {
    <div 
        class="alert"
        role="alert"
        data-variant={ string(props.Variant) }
        data-dismissible={ boolStr(props.Dismissible) }
    >
        <div class="alert-icon">
            @alertIcon(props.Variant)
        </div>
        <div class="alert-content">
            if props.Title != "" {
                <h4 class="alert-title">{ props.Title }</h4>
            }
            <div class="alert-description">
                { children... }
            </div>
        </div>
        if props.Dismissible {
            <button class="alert-dismiss" aria-label="Dismiss">
                @Icon("x")
            </button>
        }
    </div>
}

templ alertIcon(variant AlertVariant) {
    switch variant {
        case AlertSuccess:
            @Icon("check-circle")
        case AlertWarning:
            @Icon("alert-triangle")
        case AlertError:
            @Icon("alert-circle")
        default:
            @Icon("info")
    }
}
```

---

## Molecule Patterns

### Form Field (Composes Label + Input + Error)

```go
// views/components/molecules/form_field.templ

package molecules

import "myapp/views/components/atoms"

type FormFieldProps struct {
    Name        string
    Label       string
    Type        atoms.InputType
    Placeholder string
    Value       string
    Error       string
    Help        string
    Required    bool
    Disabled    bool
}

templ FormField(props FormFieldProps) {
    <div class="field" data-error={ boolStr(props.Error != "") }>
        @atoms.Label(atoms.LabelProps{
            For:      props.Name,
            Required: props.Required,
        }) {
            { props.Label }
        }
        
        @atoms.Input(atoms.InputProps{
            Name:        props.Name,
            ID:          props.Name,
            Type:        props.Type,
            Placeholder: props.Placeholder,
            Value:       props.Value,
            Required:    props.Required,
            Disabled:    props.Disabled,
            Error:       props.Error != "",
        })
        
        if props.Help != "" && props.Error == "" {
            <p class="field-help">{ props.Help }</p>
        }
        
        if props.Error != "" {
            <p class="field-error" role="alert">{ props.Error }</p>
        }
    </div>
}

// With validation via HTMX
templ FormFieldWithValidation(props FormFieldProps, validationUrl string) {
    <div class="field" data-error={ boolStr(props.Error != "") }>
        @atoms.Label(atoms.LabelProps{
            For:      props.Name,
            Required: props.Required,
        }) {
            { props.Label }
        }
        
        @atoms.Input(atoms.InputProps{
            Name:        props.Name,
            ID:          props.Name,
            Type:        props.Type,
            Placeholder: props.Placeholder,
            Value:       props.Value,
            Required:    props.Required,
            Disabled:    props.Disabled,
            Error:       props.Error != "",
            HxPost:      validationUrl,
            HxTrigger:   "blur changed",
            HxTarget:    "#" + props.Name + "-error",
        })
        
        <div id={ props.Name + "-error" }>
            if props.Error != "" {
                <p class="field-error" role="alert">{ props.Error }</p>
            }
        </div>
    </div>
}
```

### Card

```go
// views/components/molecules/card.templ

package molecules

type CardProps struct {
    Hoverable bool
    Clickable bool
}

templ Card(props CardProps) {
    <div 
        class="card"
        data-hoverable={ boolStr(props.Hoverable) }
        data-clickable={ boolStr(props.Clickable) }
    >
        { children... }
    </div>
}

templ CardHeader() {
    <div class="card-header">
        { children... }
    </div>
}

templ CardTitle() {
    <h3 class="card-title">
        { children... }
    </h3>
}

templ CardDescription() {
    <p class="card-description">
        { children... }
    </p>
}

templ CardBody() {
    <div class="card-body">
        { children... }
    </div>
}

templ CardFooter() {
    <div class="card-footer">
        { children... }
    </div>
}
```

### Dropdown Menu

```go
// views/components/molecules/dropdown_menu.templ

package molecules

import "myapp/views/components/atoms"

type DropdownProps struct {
    ID        string
    Align     Position // left, right
    Placement Position // top, bottom
}

// Dropdown container - uses Alpine.js for state
templ Dropdown(props DropdownProps) {
    <div 
        class="dropdown"
        data-align={ string(props.Align) }
        data-placement={ string(props.Placement) }
        x-data="{ open: false }"
        @click.outside="open = false"
        @keydown.escape.window="open = false"
    >
        { children... }
    </div>
}

templ DropdownTrigger() {
    <div @click="open = !open" class="dropdown-trigger">
        { children... }
    </div>
}

templ DropdownContent() {
    <div 
        class="dropdown-content"
        x-show="open"
        x-transition:enter="dropdown-enter"
        x-transition:leave="dropdown-leave"
        role="menu"
    >
        { children... }
    </div>
}

templ DropdownItem(href string) {
    if href != "" {
        <a href={ templ.SafeURL(href) } class="dropdown-item" role="menuitem">
            { children... }
        </a>
    } else {
        <button class="dropdown-item" role="menuitem">
            { children... }
        </button>
    }
}

templ DropdownSeparator() {
    <div class="dropdown-separator" role="separator"></div>
}
```

---

## Organism Patterns

### Form (Complex with HTMX)

```go
// views/components/organisms/form.templ

package organisms

import "myapp/views/components/atoms"

type FormProps struct {
    ID     string
    Action string
    Method string // "post", "get"
}

type FormSubmitProps struct {
    Action    string
    Target    string
    Swap      string
    Indicator string
}

templ Form(props FormProps) {
    <form
        id={ props.ID }
        class="form"
        action={ props.Action }
        method={ props.Method }
    >
        { children... }
    </form>
}

// HTMX-enhanced form
templ HTMXForm(props FormSubmitProps) {
    <form
        class="form"
        hx-post={ props.Action }
        hx-target={ props.Target }
        hx-swap={ props.Swap }
        hx-indicator={ props.Indicator }
    >
        { children... }
    </form>
}

templ FormSection(title string) {
    <fieldset class="form-section">
        if title != "" {
            <legend class="form-section-title">{ title }</legend>
        }
        <div class="form-section-content">
            { children... }
        </div>
    </fieldset>
}

templ FormActions() {
    <div class="form-actions">
        { children... }
    </div>
}
```

### DataTable

```go
// views/components/organisms/datatable.templ

package organisms

type Column struct {
    Key      string
    Label    string
    Sortable bool
    Width    string
}

type DataTableProps struct {
    ID          string
    Columns     []Column
    Selectable  bool
    Hoverable   bool
    Striped     bool
    // HTMX
    SortUrl     string
    PaginateUrl string
}

templ DataTable(props DataTableProps) {
    <div 
        class="datatable"
        id={ props.ID }
        data-selectable={ boolStr(props.Selectable) }
    >
        <table class="table" data-hoverable={ boolStr(props.Hoverable) } data-striped={ boolStr(props.Striped) }>
            <thead class="table-header">
                <tr>
                    if props.Selectable {
                        <th class="table-cell-select">
                            <input type="checkbox" class="checkbox" aria-label="Select all"/>
                        </th>
                    }
                    for _, col := range props.Columns {
                        @tableHeaderCell(col, props.SortUrl)
                    }
                </tr>
            </thead>
            <tbody class="table-body" id={ props.ID + "-body" }>
                { children... }
            </tbody>
        </table>
    </div>
}

templ tableHeaderCell(col Column, sortUrl string) {
    <th 
        class="table-header-cell"
        data-sortable={ boolStr(col.Sortable) }
        if col.Width != "" {
            style={ "width: " + col.Width }
        }
    >
        if col.Sortable && sortUrl != "" {
            <button 
                class="table-sort-button"
                hx-get={ sortUrl + "?sort=" + col.Key }
                hx-target="closest tbody"
            >
                { col.Label }
                <span class="table-sort-icon"></span>
            </button>
        } else {
            { col.Label }
        }
    </th>
}

templ TableRow(selectable bool) {
    <tr class="table-row">
        if selectable {
            <td class="table-cell-select">
                <input type="checkbox" class="checkbox" aria-label="Select row"/>
            </td>
        }
        { children... }
    </tr>
}

templ TableCell() {
    <td class="table-cell">
        { children... }
    </td>
}
```

---

## Template Patterns

### Base Layout

```go
// views/components/templates/base_layout.templ

package templates

type PageMeta struct {
    Title       string
    Description string
    OGImage     string
}

templ BaseLayout(meta PageMeta) {
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="UTF-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
            <title>{ meta.Title }</title>
            if meta.Description != "" {
                <meta name="description" content={ meta.Description }/>
            }
            
            // Basecoat CSS
            <link rel="stylesheet" href="/static/css/basecoat.css"/>
            
            // HTMX
            <script src="/static/js/htmx.min.js" defer></script>
            
            // Alpine.js
            <script src="/static/js/alpine.min.js" defer></script>
        </head>
        <body class="antialiased">
            { children... }
        </body>
    </html>
}
```

### Dashboard Layout

```go
// views/components/templates/dashboard_layout.templ

package templates

type DashboardLayoutProps struct {
    Title string
    User  UserInfo
}

type UserInfo struct {
    Name   string
    Email  string
    Avatar string
}

templ DashboardLayout(props DashboardLayoutProps) {
    @BaseLayout(PageMeta{Title: props.Title}) {
        <div class="layout-dashboard" x-data="{ sidebarOpen: true }">
            // Sidebar
            <aside 
                class="sidebar"
                :data-open="sidebarOpen"
            >
                <div class="sidebar-header">
                    { sidebarHeader... }
                </div>
                <nav class="sidebar-nav">
                    { sidebarNav... }
                </nav>
                <div class="sidebar-footer">
                    { sidebarFooter... }
                </div>
            </aside>
            
            // Main content
            <div class="layout-main">
                <header class="topbar">
                    <button 
                        class="btn"
                        data-variant="ghost"
                        data-size="icon"
                        @click="sidebarOpen = !sidebarOpen"
                        aria-label="Toggle sidebar"
                    >
                        @Icon("menu")
                    </button>
                    
                    <div class="topbar-spacer"></div>
                    
                    { topbarActions... }
                </header>
                
                <main class="main-content">
                    { children... }
                </main>
            </div>
        </div>
    }
}
```

---

## Helper Functions

```go
// views/components/helpers.go

package components

import "strconv"

// boolStr converts bool to "true"/"false" string for data attributes
func boolStr(b bool) string {
    return strconv.FormatBool(b)
}

// boolAttr returns the attribute value only if true
func boolAttr(b bool) string {
    if b {
        return "true"
    }
    return ""
}

// HTMX attribute helpers
type HtmxAttrs struct {
    Get       string
    Post      string
    Put       string
    Delete    string
    Target    string
    Swap      string
    Trigger   string
    Indicator string
    Confirm   string
    PushUrl   bool
}
```

---

## CSS Requirements

For these patterns to work, your `basecoat.css` must include data-attribute selectors:

```css
/* Button variants via data-variant */
.btn[data-variant="primary"] { /* primary styles */ }
.btn[data-variant="secondary"] { /* secondary styles */ }
.btn[data-variant="outline"] { /* outline styles */ }
.btn[data-variant="ghost"] { /* ghost styles */ }

/* Sizes via data-size */
.btn[data-size="sm"] { /* small styles */ }
.btn[data-size="md"] { /* medium styles */ }
.btn[data-size="lg"] { /* large styles */ }
.btn[data-size="icon"] { /* icon button styles */ }

/* States via data-* */
.btn[data-loading="true"] { 
    opacity: 0.5; 
    cursor: wait; 
    pointer-events: none;
}

/* Field error state */
.field[data-error="true"] .input {
    border-color: var(--color-danger);
}

.field[data-error="true"] .field-error {
    display: block;
}
```

---

## Checklist for Each Component

When refactoring a component, verify:

- [ ] No `TwMerge` or `twMerge` usage
- [ ] No `ClassName` or `className` props
- [ ] Uses Basecoat semantic classes
- [ ] Uses data-* attributes for variants/states
- [ ] Has typed Go constants for variants
- [ ] Props struct if >3 parameters
- [ ] Proper accessibility (aria-*, role)
- [ ] HTMX integration if interactive
- [ ] Compiles with `templ generate`
- [ ] Matches Basecoat docs examples

---

*Last updated: Component Refactoring Project*
