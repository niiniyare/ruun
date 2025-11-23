# AWO ERP Component Implementation Guide

## EXECUTIVE SUMMARY

**Component Framework:** Pure templ components with Tailwind CSS utilities  
**Architecture:** Atomic Design (Atoms → Molecules → Organisms → Templates → Pages)  
**Styling:** Tailwind classes generated from theme tokens  
**State:** Server-driven (HTMX) + Client-side (Alpine.js)  
**Philosophy:** Zero business logic, props-first, presentation-only

---

## TOKEN USAGE IN COMPONENTS

### Theme Token Structure

Components reference tokens via Tailwind utility classes. Tokens are defined in theme JSON files:

```json
{
  "name": "Default",
  "tokens": {
    "colors": {
      "primary": "hsl(217, 91%, 60%)",
      "primary-foreground": "hsl(0, 0%, 100%)",
      "background": "hsl(0, 0%, 100%)",
      "foreground": "hsl(0, 0%, 9%)",
      "border": "hsl(0, 0%, 90%)",
      "muted": "hsl(0, 0%, 98%)",
      "muted-foreground": "hsl(0, 0%, 50%)"
    },
    "spacing": {
      "xs": "0.5rem",
      "sm": "0.75rem",
      "md": "1rem",
      "lg": "1.5rem"
    },
    "radius": {
      "sm": "0.25rem",
      "md": "0.5rem",
      "lg": "0.75rem"
    }
  },
  "components": {
    "button": {
      "primary": {
        "background-color": "hsl(217, 91%, 60%)",
        "color": "hsl(0, 0%, 100%)",
        "border-radius": "0.5rem",
        "padding": "0.5rem 1rem"
      }
    }
  }
}
```

### How Components Use Tokens

**Theme tokens compile to Tailwind utilities:**

```css
/* Generated from theme tokens */
.bg-primary { background-color: hsl(217, 91%, 60%); }
.text-primary { color: hsl(217, 91%, 60%); }
.text-primary-foreground { color: hsl(0, 0%, 100%); }
.border-border { border-color: hsl(0, 0%, 90%); }
.rounded-md { border-radius: 0.5rem; }
```

**Components use these classes:**

```templ
templ Button(props ButtonProps) {
    <button class="bg-primary text-primary-foreground rounded-md px-4 py-2">
        {props.Label}
    </button>
}
```

---

## ATOMIC DESIGN STRUCTURE

### Level 1: ATOMS (Foundation Elements)

**Definition:** Smallest UI units that cannot be broken down further.

**Key Atoms:**
- Button, Input, Label, Icon, Badge, Spinner, Avatar
- Checkbox, Radio, Select, Textarea, Toggle
- Link, Image, Heading, Text, Divider

**Atom Pattern:**

```go
// Props struct with typed fields
type ButtonProps struct {
    Label    string
    Variant  ButtonVariant  // "primary" | "secondary" | "outline" | "ghost"
    Size     ButtonSize     // "xs" | "sm" | "md" | "lg" | "xl"
    Disabled bool
    Loading  bool
    Type     string         // "button" | "submit" | "reset"
    
    // HTMX
    HxPost   string
    HxTarget string
    HxSwap   string
    
    // Additional
    ClassName string
    AriaLabel string
}

// Variant types for type safety
type ButtonVariant string
const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonOutline   ButtonVariant = "outline"
    ButtonGhost     ButtonVariant = "ghost"
)

type ButtonSize string
const (
    ButtonXS ButtonSize = "xs"
    ButtonSM ButtonSize = "sm"
    ButtonMD ButtonSize = "md"
    ButtonLG ButtonSize = "lg"
    ButtonXL ButtonSize = "xl"
)

// Class composition function
func getButtonClasses(variant ButtonVariant, size ButtonSize, disabled bool) string {
    // Base classes (always applied)
    base := "inline-flex items-center justify-center font-medium transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2"
    
    // Variant-specific classes
    var variantClasses string
    switch variant {
    case ButtonPrimary:
        variantClasses = "bg-primary text-primary-foreground hover:bg-primary/90 focus:ring-primary"
    case ButtonSecondary:
        variantClasses = "bg-secondary text-secondary-foreground hover:bg-secondary/80 focus:ring-secondary"
    case ButtonOutline:
        variantClasses = "border border-border bg-background hover:bg-muted focus:ring-ring"
    case ButtonGhost:
        variantClasses = "hover:bg-muted hover:text-foreground"
    }
    
    // Size-specific classes
    var sizeClasses string
    switch size {
    case ButtonXS:
        sizeClasses = "text-xs px-2 py-1 rounded-sm"
    case ButtonSM:
        sizeClasses = "text-sm px-3 py-1.5 rounded-md"
    case ButtonMD:
        sizeClasses = "text-sm px-4 py-2 rounded-md"
    case ButtonLG:
        sizeClasses = "text-base px-5 py-2.5 rounded-lg"
    case ButtonXL:
        sizeClasses = "text-base px-6 py-3 rounded-lg"
    }
    
    // State classes
    classes := []string{base, variantClasses, sizeClasses}
    if disabled {
        classes = append(classes, "opacity-50 pointer-events-none")
    }
    
    return strings.Join(classes, " ")
}

// Templ component
templ Button(props ButtonProps) {
    <button
        type={props.Type}
        class={getButtonClasses(props.Variant, props.Size, props.Disabled)}
        disabled?={props.Disabled || props.Loading}
        if props.HxPost != "" {
            hx-post={props.HxPost}
            hx-target={props.HxTarget}
            hx-swap={props.HxSwap}
        }
        if props.AriaLabel != "" {
            aria-label={props.AriaLabel}
        }
    >
        if props.Loading {
            @Spinner(SpinnerProps{Size: "sm", ClassName: "mr-2"})
        }
        {props.Label}
    </button>
}
```

**Input Atom Example:**

```go
type InputProps struct {
    Name        string
    Type        string  // "text" | "email" | "password" | "number"
    Value       string
    Placeholder string
    Disabled    bool
    Error       bool
    Required    bool
    ClassName   string
    
    // Accessibility
    ID              string
    AriaLabel       string
    AriaDescribedBy string
    AriaInvalid     bool
}

func getInputClasses(error bool, disabled bool, className string) string {
    base := "w-full rounded-md border bg-input px-3 py-2 text-sm transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2"
    
    var stateClasses string
    if error {
        stateClasses = "border-error focus:ring-error"
    } else if disabled {
        stateClasses = "bg-muted cursor-not-allowed opacity-50"
    } else {
        stateClasses = "border-border focus:ring-ring"
    }
    
    classes := []string{base, stateClasses}
    if className != "" {
        classes = append(classes, className)
    }
    
    return strings.Join(classes, " ")
}

templ Input(props InputProps) {
    <input
        type={props.Type}
        name={props.Name}
        value={props.Value}
        placeholder={props.Placeholder}
        class={getInputClasses(props.Error, props.Disabled, props.ClassName)}
        disabled?={props.Disabled}
        required?={props.Required}
        if props.ID != "" {
            id={props.ID}
        }
        if props.AriaLabel != "" {
            aria-label={props.AriaLabel}
        }
        if props.AriaDescribedBy != "" {
            aria-describedby={props.AriaDescribedBy}
        }
        aria-invalid={strconv.FormatBool(props.AriaInvalid)}
    />
}
```

**Badge Atom Example:**

```go
type BadgeProps struct {
    Label     string
    Variant   BadgeVariant  // "default" | "success" | "warning" | "error"
    Size      BadgeSize     // "sm" | "md" | "lg"
    ClassName string
}

type BadgeVariant string
const (
    BadgeDefault BadgeVariant = "default"
    BadgeSuccess BadgeVariant = "success"
    BadgeWarning BadgeVariant = "warning"
    BadgeError   BadgeVariant = "error"
)

func getBadgeClasses(variant BadgeVariant, size string) string {
    base := "inline-flex items-center rounded-full font-medium"
    
    variantMap := map[BadgeVariant]string{
        BadgeDefault: "bg-muted text-muted-foreground",
        BadgeSuccess: "bg-success text-success-foreground",
        BadgeWarning: "bg-warning text-warning-foreground",
        BadgeError:   "bg-error text-error-foreground",
    }
    
    sizeMap := map[string]string{
        "sm": "text-xs px-2 py-0.5",
        "md": "text-sm px-2.5 py-0.5",
        "lg": "text-base px-3 py-1",
    }
    
    return fmt.Sprintf("%s %s %s", base, variantMap[variant], sizeMap[size])
}

templ Badge(props BadgeProps) {
    <span class={getBadgeClasses(props.Variant, "md")}>
        {props.Label}
    </span>
}
```

---

### Level 2: MOLECULES (Simple Compositions)

**Definition:** 2-5 atoms combined for a single purpose.

**Key Molecules:**
- FormField (Label + Input + HelpText + ErrorMessage)
- SearchBar (Icon + Input + ClearButton)
- StatCard (Icon + Label + Value + Trend)

**Molecule Pattern (FormField):**

```go
type FormFieldProps struct {
    Label       string
    Name        string
    Type        string
    Value       string
    Placeholder string
    Required    bool
    Disabled    bool
    HelpText    string
    ErrorMsg    string
    ShowError   bool
    ClassName   string
}

templ FormField(props FormFieldProps) {
    <div class={templ.Classes("space-y-2", props.ClassName)}>
        // Label atom
        @Label(LabelProps{
            For:      props.Name,
            Text:     props.Label,
            Required: props.Required,
        })
        
        // Input atom
        @Input(InputProps{
            ID:              props.Name,
            Name:            props.Name,
            Type:            props.Type,
            Value:           props.Value,
            Placeholder:     props.Placeholder,
            Disabled:        props.Disabled,
            Required:        props.Required,
            Error:           props.ShowError,
            AriaDescribedBy: props.Name + "-help " + props.Name + "-error",
            AriaInvalid:     props.ShowError,
        })
        
        // Help text (conditional)
        if props.HelpText != "" && !props.ShowError {
            <p id={props.Name + "-help"} class="text-sm text-muted-foreground">
                {props.HelpText}
            </p>
        }
        
        // Error message (conditional)
        if props.ShowError && props.ErrorMsg != "" {
            <p id={props.Name + "-error"} class="text-sm text-error" role="alert">
                {props.ErrorMsg}
            </p>
        }
    </div>
}
```

**SearchBar Molecule:**

```go
type SearchBarProps struct {
    Placeholder string
    Value       string
    OnClear     string
    HxGet       string
    HxTarget    string
    HxTrigger   string
    ClassName   string
}

templ SearchBar(props SearchBarProps) {
    <div class={templ.Classes("relative", props.ClassName)}>
        // Icon atom (search icon)
        <div class="absolute left-3 top-1/2 -translate-y-1/2">
            @Icon(IconProps{Name: "search", Size: "sm", ClassName: "text-muted-foreground"})
        </div>
        
        // Input atom with padding for icon
        <input
            type="search"
            value={props.Value}
            placeholder={props.Placeholder}
            class="w-full rounded-md border border-border bg-background pl-10 pr-10 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-ring"
            if props.HxGet != "" {
                hx-get={props.HxGet}
                hx-target={props.HxTarget}
                hx-trigger={props.HxTrigger}
            }
        />
        
        // Clear button (conditional)
        if props.Value != "" {
            <button
                type="button"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground"
                onclick={props.OnClear}
            >
                @Icon(IconProps{Name: "x", Size: "sm"})
            </button>
        }
    </div>
}
```

**StatCard Molecule:**

```go
type StatCardProps struct {
    Icon      string
    Label     string
    Value     string
    Trend     string  // "+12%" or "-5%"
    Positive  bool
    ClassName string
}

templ StatCard(props StatCardProps) {
    <div class={templ.Classes("rounded-lg border border-border bg-background p-6", props.ClassName)}>
        <div class="flex items-center justify-between">
            // Icon atom
            <div class="rounded-full bg-primary/10 p-3">
                @Icon(IconProps{Name: props.Icon, Size: "md", ClassName: "text-primary"})
            </div>
            
            // Trend badge
            if props.Trend != "" {
                @Badge(BadgeProps{
                    Label:   props.Trend,
                    Variant: templ.KV(BadgeSuccess, props.Positive).(BadgeVariant),
                })
            }
        </div>
        
        <div class="mt-4 space-y-1">
            <p class="text-sm text-muted-foreground">{props.Label}</p>
            <p class="text-2xl font-bold text-foreground">{props.Value}</p>
        </div>
    </div>
}
```

---

### Level 3: ORGANISMS (Complex Components)

**Definition:** Complex, self-contained sections with multiple molecules/atoms.

**Key Organisms:**
- DataTable (with sorting, filtering, pagination)
- Navigation (header, menu, user dropdown)
- Card (header, content, footer with actions)

**Card Organism:**

```go
type CardProps struct {
    Title       string
    Description string
    Badge       *BadgeProps
    Actions     []ActionButton
    Children    templ.Component
    ClassName   string
}

type ActionButton struct {
    Label   string
    Variant ButtonVariant
    OnClick string
    HxPost  string
}

templ Card(props CardProps) {
    <div class={templ.Classes("rounded-lg border border-border bg-background", props.ClassName)}>
        // Header (if title or actions)
        if props.Title != "" || len(props.Actions) > 0 {
            <div class="flex items-center justify-between border-b border-border px-6 py-4">
                <div class="space-y-1">
                    <div class="flex items-center gap-2">
                        <h3 class="text-lg font-semibold text-foreground">{props.Title}</h3>
                        if props.Badge != nil {
                            @Badge(*props.Badge)
                        }
                    </div>
                    if props.Description != "" {
                        <p class="text-sm text-muted-foreground">{props.Description}</p>
                    }
                </div>
                
                // Actions
                if len(props.Actions) > 0 {
                    <div class="flex items-center gap-2">
                        for _, action := range props.Actions {
                            @Button(ButtonProps{
                                Label:   action.Label,
                                Variant: action.Variant,
                                Size:    ButtonSM,
                                HxPost:  action.HxPost,
                            })
                        }
                    </div>
                }
            </div>
        }
        
        // Content
        <div class="px-6 py-4">
            {props.Children...}
        </div>
    </div>
}
```

**DataTable Organism:**

```go
type DataTableProps struct {
    Columns      []Column
    Rows         []map[string]any
    Sortable     bool
    CurrentSort  *SortConfig
    Pagination   *PaginationConfig
    HxGet        string
    HxTarget     string
    EmptyMessage string
}

type Column struct {
    Key       string
    Label     string
    Sortable  bool
    Width     string
    Align     string  // "left" | "center" | "right"
    Format    func(any) string
}

type SortConfig struct {
    Column    string
    Direction string  // "asc" | "desc"
}

type PaginationConfig struct {
    Page       int
    PageSize   int
    Total      int
    HxGet      string
}

templ DataTable(props DataTableProps) {
    <div class="rounded-lg border border-border overflow-hidden">
        // Table
        <table class="w-full">
            <thead class="bg-muted">
                <tr>
                    for _, col := range props.Columns {
                        <th 
                            class={getHeaderClasses(col.Align, col.Sortable)}
                            if col.Width != "" {
                                style={"width: " + col.Width}
                            }
                        >
                            if col.Sortable && props.Sortable {
                                <button
                                    class="flex items-center gap-2 font-medium text-sm"
                                    hx-get={props.HxGet + "?sort=" + col.Key}
                                    hx-target={props.HxTarget}
                                >
                                    {col.Label}
                                    @SortIcon(props.CurrentSort, col.Key)
                                </button>
                            } else {
                                <span class="font-medium text-sm">{col.Label}</span>
                            }
                        </th>
                    }
                </tr>
            </thead>
            <tbody class="divide-y divide-border">
                if len(props.Rows) == 0 {
                    <tr>
                        <td colspan={fmt.Sprintf("%d", len(props.Columns))} class="px-6 py-8 text-center">
                            <p class="text-sm text-muted-foreground">{props.EmptyMessage}</p>
                        </td>
                    </tr>
                } else {
                    for _, row := range props.Rows {
                        <tr class="hover:bg-muted/50">
                            for _, col := range props.Columns {
                                <td class={getCellClasses(col.Align)}>
                                    if col.Format != nil {
                                        {col.Format(row[col.Key])}
                                    } else {
                                        {fmt.Sprint(row[col.Key])}
                                    }
                                </td>
                            }
                        </tr>
                    }
                }
            </tbody>
        </table>
        
        // Pagination
        if props.Pagination != nil {
            @Pagination(*props.Pagination)
        }
    </div>
}

func getHeaderClasses(align string, sortable bool) string {
    base := "px-6 py-3 text-left"
    
    alignMap := map[string]string{
        "left":   "text-left",
        "center": "text-center",
        "right":  "text-right",
    }
    
    classes := []string{base, alignMap[align]}
    if sortable {
        classes = append(classes, "cursor-pointer hover:bg-muted/80")
    }
    
    return strings.Join(classes, " ")
}

func getCellClasses(align string) string {
    alignMap := map[string]string{
        "left":   "text-left",
        "center": "text-center",
        "right":  "text-right",
    }
    
    return fmt.Sprintf("px-6 py-4 text-sm %s", alignMap[align])
}
```

---

### Level 4: TEMPLATES (Page Layouts)

**Definition:** Page structure with slot-based composition.

**MasterDetailLayout Template:**

```go
type MasterDetailLayoutProps struct {
    MasterWidth string  // "300px" | "25%" | "33%"
    Master      templ.Component
    Detail      templ.Component
    Header      templ.Component
}

templ MasterDetailLayout(props MasterDetailLayoutProps) {
    <div class="flex flex-col h-screen">
        // Header
        if props.Header != nil {
            <header class="border-b border-border bg-background">
                {props.Header...}
            </header>
        }
        
        // Main content
        <div class="flex flex-1 overflow-hidden">
            // Master (sidebar)
            <aside 
                class="border-r border-border bg-background overflow-y-auto"
                style={"width: " + props.MasterWidth}
            >
                {props.Master...}
            </aside>
            
            // Detail (main)
            <main class="flex-1 overflow-y-auto p-6">
                {props.Detail...}
            </main>
        </div>
    </div>
}
```

**FormLayout Template:**

```go
type FormLayoutProps struct {
    Title       string
    Description string
    Breadcrumbs []Breadcrumb
    Sections    []FormSection
    Actions     []ActionButton
}

type FormSection struct {
    Title       string
    Description string
    Fields      templ.Component
}

templ FormLayout(props FormLayoutProps) {
    <div class="mx-auto max-w-4xl space-y-6 p-6">
        // Breadcrumbs
        if len(props.Breadcrumbs) > 0 {
            @Breadcrumbs(BreadcrumbsProps{Items: props.Breadcrumbs})
        }
        
        // Header
        <div class="space-y-2">
            <h1 class="text-3xl font-bold text-foreground">{props.Title}</h1>
            if props.Description != "" {
                <p class="text-muted-foreground">{props.Description}</p>
            }
        </div>
        
        // Form sections
        <form class="space-y-8">
            for _, section := range props.Sections {
                <div class="rounded-lg border border-border bg-background p-6">
                    <div class="space-y-4">
                        <div>
                            <h2 class="text-lg font-semibold text-foreground">{section.Title}</h2>
                            if section.Description != "" {
                                <p class="text-sm text-muted-foreground mt-1">{section.Description}</p>
                            }
                        </div>
                        
                        <div class="space-y-4">
                            {section.Fields...}
                        </div>
                    </div>
                </div>
            }
            
            // Actions
            <div class="flex items-center justify-end gap-3 border-t border-border pt-6">
                for _, action := range props.Actions {
                    @Button(ButtonProps{
                        Label:   action.Label,
                        Variant: action.Variant,
                        HxPost:  action.HxPost,
                    })
                }
            </div>
        </form>
    </div>
}
```

---

### Level 5: PAGES (Route-Specific)

**Definition:** Complete pages with data and business logic integration.

**InvoiceListPage:**

```go
type InvoiceListPageProps struct {
    Invoices   []Invoice
    Stats      InvoiceStats
    Filters    ActiveFilters
    Pagination PaginationConfig
    CanCreate  bool
}

templ InvoiceListPage(props InvoiceListPageProps) {
    @MasterDetailLayout(MasterDetailLayoutProps{
        MasterWidth: "400px",
        Header: invoiceListHeader(props),
        Master: invoiceListMaster(props),
        Detail: invoiceListDetail(props),
    })
}

templ invoiceListHeader(props InvoiceListPageProps) {
    <div class="flex items-center justify-between px-6 py-4">
        <div class="flex items-center gap-4">
            <h1 class="text-xl font-semibold">Invoices</h1>
            @StatCard(StatCardProps{
                Icon:  "file-text",
                Label: "Total",
                Value: fmt.Sprintf("%d", props.Stats.Total),
            })
        </div>
        
        if props.CanCreate {
            @Button(ButtonProps{
                Label:   "New Invoice",
                Variant: ButtonPrimary,
                HxPost:  "/invoices/new",
            })
        }
    </div>
}

templ invoiceListMaster(props InvoiceListPageProps) {
    <div class="p-4 space-y-4">
        // Search
        @SearchBar(SearchBarProps{
            Placeholder: "Search invoices...",
            HxGet:       "/invoices/search",
            HxTarget:    "#invoice-table",
            HxTrigger:   "keyup changed delay:300ms",
        })
        
        // Filters
        @FilterPanel(FilterPanelProps{
            Active:   props.Filters,
            HxGet:    "/invoices/filter",
            HxTarget: "#invoice-table",
        })
        
        // Table
        <div id="invoice-table">
            @DataTable(DataTableProps{
                Columns: getInvoiceColumns(),
                Rows:    invoicesToRows(props.Invoices),
                Pagination: &props.Pagination,
            })
        </div>
    </div>
}
```

---

## TEMPL FUNCTION PATTERNS

### 1. Component Composition

**Children Slots:**

```templ
templ Container(title string, children templ.Component) {
    <div class="container">
        <h2>{title}</h2>
        <div class="content">
            {children...}
        </div>
    </div>
}

// Usage
@Container("Dashboard",
    @StatCard(StatCardProps{...})
)
```

**Multiple Slots:**

```templ
type DialogProps struct {
    Header  templ.Component
    Content templ.Component
    Footer  templ.Component
}

templ Dialog(props DialogProps) {
    <div class="dialog">
        <div class="header">{props.Header...}</div>
        <div class="content">{props.Content...}</div>
        <div class="footer">{props.Footer...}</div>
    </div>
}
```

### 2. Conditional Classes

**Using templ.Classes:**

```templ
templ Alert(variant string, message string) {
    <div class={templ.Classes(
        "rounded-lg p-4 border",
        templ.KV("bg-success/10 border-success text-success", variant == "success"),
        templ.KV("bg-error/10 border-error text-error", variant == "error"),
        templ.KV("bg-warning/10 border-warning text-warning", variant == "warning"),
    )}>
        {message}
    </div>
}
```

### 3. Dynamic Lists

**Rendering Collections:**

```templ
templ TagList(tags []string) {
    <div class="flex flex-wrap gap-2">
        for _, tag := range tags {
            @Badge(BadgeProps{Label: tag, Variant: BadgeDefault})
        }
    </div>
}
```

### 4. Helper Functions

**For Complex Class Logic:**

```go
// In same file as component
func getAlertClasses(variant string, dismissible bool) string {
    base := "rounded-lg p-4 border"
    
    variantMap := map[string]string{
        "success": "bg-success/10 border-success text-success",
        "error":   "bg-error/10 border-error text-error",
        "warning": "bg-warning/10 border-warning text-warning",
        "info":    "bg-primary/10 border-primary text-primary",
    }
    
    classes := []string{base, variantMap[variant]}
    if dismissible {
        classes = append(classes, "pr-12")
    }
    
    return strings.Join(classes, " ")
}

templ Alert(props AlertProps) {
    <div class={getAlertClasses(props.Variant, props.Dismissible)}>
        {props.Message}
    </div>
}
```

---

## COMPONENT FILE STRUCTURE

```
views/
├── components/
│   ├── atoms/
│   │   ├── button.templ
│   │   ├── button_types.go       # Props, variants, helper functions
│   │   ├── input.templ
│   │   ├── input_types.go
│   │   ├── badge.templ
│   │   ├── badge_types.go
│   │   └── ...
│   ├── molecules/
│   │   ├── form_field.templ
│   │   ├── form_field_types.go
│   │   ├── search_bar.templ
│   │   ├── search_bar_types.go
│   │   └── ...
│   ├── organisms/
│   │   ├── data_table.templ
│   │   ├── data_table_types.go
│   │   ├── card.templ
│   │   ├── card_types.go
│   │   └── ...
│   ├── templates/
│   │   ├── master_detail_layout.templ
│   │   ├── master_detail_layout_types.go
│   │   └── ...
│   └── pages/
│       ├── invoice_list.templ
│       ├── invoice_list_types.go
│       └── ...
└── styles/
    └── themes/
        ├── tenant_a.json
        ├── tenant_b.json
        └── default.json
```

---

## CRITICAL PATTERNS

### ✅ DO: Props-First Architecture

```go
// External logic computes everything
props := ButtonProps{
    Label:    "Save Invoice",
    Variant:  ButtonPrimary,
    Disabled: !invoice.IsValid(),  // Business logic HERE
    HxPost:   fmt.Sprintf("/api/invoices/%s", invoice.ID),  // URL built HERE
}

@Button(props)
```

### ✅ DO: Type-Safe Variants

```go
type ButtonVariant string
const (
    ButtonPrimary ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
)

// Prevents typos, enables autocomplete
@Button(ButtonProps{Variant: ButtonPrimary})
```

### ✅ DO: Class Composition Functions

```go
func getCardClasses(variant string, hoverable bool) string {
    base := "rounded-lg border bg-background"
    
    variantMap := map[string]string{
        "default":   "border-border",
        "highlighted": "border-primary bg-primary/5",
    }
    
    classes := []string{base, variantMap[variant]}
    if hoverable {
        classes = append(classes, "hover:shadow-md transition-shadow")
    }
    
    return strings.Join(classes, " ")
}
```

### ❌ DON'T: Business Logic in Components

```templ
<!-- WRONG -->
templ InvoiceButton(invoice Invoice, user User) {
    if user.HasPermission("invoice.edit") && invoice.Status == "draft" {
        <button>Edit</button>
    }
}

<!-- CORRECT -->
templ InvoiceButton(canEdit bool) {
    if canEdit {
        <button>Edit</button>
    }
}
```

### ❌ DON'T: Dynamic URL Construction

```templ
<!-- WRONG -->
<form hx-post={fmt.Sprintf("/api/items/%s", itemID)}>

<!-- CORRECT -->
<form hx-post={props.HxPost}>  // URL passed as prop
```

---

## INTEGRATION EXAMPLES

### HTMX Integration

```templ
templ InvoiceForm(props InvoiceFormProps) {
    <form
        hx-post="/api/invoices"
        hx-target="#invoice-list"
        hx-swap="beforeend"
        hx-disabled-elt="#submit-btn"
    >
        @FormField(FormFieldProps{
            Label: "Customer",
            Name:  "customer_id",
            Type:  "text",
        })
        
        @Button(ButtonProps{
            ID:      "submit-btn",
            Label:   "Create Invoice",
            Variant: ButtonPrimary,
            Type:    "submit",
        })
    </form>
}
```

### Alpine.js Integration

```templ
templ Dropdown(props DropdownProps) {
    <div x-data="{ open: false }" @click.away="open = false">
        <button 
            @click="open = !open"
            class="inline-flex items-center gap-2 rounded-md px-4 py-2 border border-border"
        >
            {props.Label}
            @Icon(IconProps{Name: "chevron-down", Size: "sm"})
        </button>
        
        <div 
            x-show="open"
            x-transition
            class="absolute mt-2 w-48 rounded-md border border-border bg-background shadow-lg"
        >
            for _, item := range props.Items {
                <button class="block w-full px-4 py-2 text-left hover:bg-muted">
                    {item.Label}
                </button>
            }
        </div>
    </div>
}
```

---

## QUICK REFERENCE

### When to Use Each Level

| Level | Use When | Example |
|-------|----------|---------|
| **Atom** | Cannot be broken down | Button, Input, Badge |
| **Molecule** | 2-5 atoms, single purpose | FormField, SearchBar |
| **Organism** | Complex, multiple molecules | DataTable, Card, Navigation |
| **Template** | Page structure | MasterDetailLayout, FormLayout |
| **Page** | Specific route | InvoiceListPage, DashboardPage |

### Component Checklist

- [ ] Props struct defined with types
- [ ] Variant enums for type safety (if applicable)
- [ ] Class composition helper function
- [ ] templ component using helpers
- [ ] Conditional rendering with if/for
- [ ] HTMX attributes as props (if interactive)
- [ ] Alpine.js directives (if client state needed)
- [ ] Accessibility attributes (aria-*, role)
- [ ] No business logic in component
- [ ] No dynamic URL construction

---

**Remember:** Components are pure presentation. All logic, data fetching, permissions, and URL construction happen BEFORE props are passed in.
