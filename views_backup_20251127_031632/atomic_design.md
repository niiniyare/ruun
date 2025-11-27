# Atomic Design: Building Scalable UI Systems

**A Practical Guide for Go & Templ Developers**

---

## Table of Contents

1. [Core Concepts](#core-concepts)
2. [The Five Levels Explained](#the-five-levels-explained)
3. [Level 1: Atoms](#level-1-atoms)
4. [Level 2: Molecules](#level-2-molecules)
5. [Level 3: Organisms](#level-3-organisms)
6. [Level 4: Templates](#level-4-templates)
7. [Level 5: Pages](#level-5-pages)
8. [Composition Principles](#composition-principles)
9. [Design Tokens Foundation](#design-tokens-foundation)
10. [Schema-Driven Implementation](#schema-driven-implementation)
11. [Testing Strategy](#testing-strategy)
12. [Real-World ERP Examples](#real-world-erp-examples)
13. [Migration Guide](#migration-guide)

---

## Core Concepts

### What is Atomic Design?

Atomic Design is a methodology created by Brad Frost that breaks down user interfaces into hierarchical building blocks. Think of it like constructing a building:

- **Atoms** are the raw materials (bricks, wood, nails)
- **Molecules** are simple assemblies (door frame, window unit)
- **Organisms** are complete sections (wall with windows, roof structure)
- **Templates** are blueprints (floor plan, structural layout)
- **Pages** are the actual building (123 Main Street, fully furnished)

### Why This Matters for Go/Templ Development

**Traditional Approach Problems:**
```
❌ Copy-paste templ components across pages
❌ Inconsistent styling and behavior
❌ Hard to maintain and update
❌ Testing requires full page context
❌ No clear component boundaries
```

**Atomic Design Benefits:**
```
✅ Single source of truth for each component
✅ Predictable composition patterns
✅ Easy to test in isolation
✅ Clear responsibility boundaries
✅ Schema-driven UI becomes natural
```

### The Hierarchy at a Glance

| Level | Complexity | Example | State | Reusability |
|-------|-----------|---------|-------|-------------|
| **Atom** | Minimal | Button, Input | None/Presentational | Universal |
| **Molecule** | Simple | Search bar, Form field | Local UI | High |
| **Organism** | Complex | Data table, Nav bar | Business logic | Medium |
| **Template** | Structural | Dashboard layout | Layout state | Medium |
| **Page** | Complete | Invoice #123 | App state | Single-use |

### Decision Framework

When creating or classifying a component:

```
┌─────────────────────────────────────┐
│ Can it be broken down further?      │
│ No → ATOM                            │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│ Does it combine 2-5 simple parts?   │
│ Yes → MOLECULE                       │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│ Does it contain business logic?     │
│ Yes → ORGANISM                       │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│ Is it defining page structure?      │
│ Yes → TEMPLATE                       │
└─────────────────────────────────────┘
                ↓
┌─────────────────────────────────────┐
│ Is it a specific instance with data?│
│ Yes → PAGE                           │
└─────────────────────────────────────┘
```

---

## The Five Levels Explained

### Dependency Direction

**Critical Rule:** Dependencies flow in ONE direction only.

```
Pages ──────────► Templates
                      ▲
Templates ───────► Organisms
                      ▲
Organisms ───────► Molecules
                      ▲
Molecules ───────► Atoms
                      ▲
Atoms ───────────► Design Tokens
```

**What This Means:**
- Atoms never import Molecules
- Molecules never import Organisms
- Changes to Atoms may affect everything above
- Changes to Pages never affect Atoms

This creates **predictable cascading** and **isolated testing**.

### Composition Over Configuration

Build complexity through composition, not through giant configuration objects.

**Poor Approach:**
```go
// ❌ Massive configuration object
type MegaComponent struct {
    ShowHeader bool
    ShowFooter bool
    HeaderColor string
    HeaderSize string
    BodyPadding string
    // ... 50 more fields
}
```

**Better Approach:**
```go
// ✅ Composable pieces
type Card struct {
    Header HeaderAtom
    Body   BodyAtom
    Footer FooterAtom
}
```

---

## Level 1: Atoms

### Definition

The smallest, indivisible UI elements that cannot be broken down further without losing their meaning or function.

**Key Characteristics:**
- Zero business logic
- Pure presentation
- Highly reusable (used 50+ times)
- No knowledge of application context
- Always use design tokens, never hard-coded values
- Should work in complete isolation

### When Something is NOT an Atom

If it contains:
- Multiple distinct parts that could be separated
- Business logic or validation rules
- API calls or data fetching
- Complex state management
- Knowledge of where it's being used

### Core Atom Categories

**1. Form Controls**
```go
// Button atom types
type ButtonVariant string
const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonDanger    ButtonVariant = "danger"
    ButtonGhost     ButtonVariant = "ghost"
)

type ButtonSize string
const (
    ButtonSmall  ButtonSize = "sm"
    ButtonMedium ButtonSize = "md"
    ButtonLarge  ButtonSize = "lg"
)

type ButtonProps struct {
    Label    string
    Variant  ButtonVariant
    Size     ButtonSize
    Disabled bool
    Loading  bool
    Type     string // "button", "submit", "reset"
    // No onClick handler - atoms don't handle events
}

// Input atom types
type InputType string
const (
    InputText     InputType = "text"
    InputEmail    InputType = "email"
    InputPassword InputType = "password"
    InputNumber   InputType = "number"
    InputDate     InputType = "date"
)

type InputProps struct {
    Type        InputType
    Name        string
    Value       string
    Placeholder string
    Disabled    bool
    ReadOnly    bool
    Required    bool
    Error       bool // Visual state only, not the error message
    // No onChange handler - atoms don't handle events
}

// Select dropdown
type SelectProps struct {
    Name     string
    Value    string
    Disabled bool
    Error    bool
    // Options passed via children/slots
}

// Checkbox
type CheckboxProps struct {
    Name     string
    Checked  bool
    Disabled bool
    Value    string
}
```

**2. Typography**
```go
type HeadingLevel string
const (
    H1 HeadingLevel = "h1"
    H2 HeadingLevel = "h2"
    H3 HeadingLevel = "h3"
    H4 HeadingLevel = "h4"
    H5 HeadingLevel = "h5"
    H6 HeadingLevel = "h6"
)

type HeadingProps struct {
    Level HeadingLevel
    Text  string
    Class string // For token-based styling
}

type TextProps struct {
    Content string
    Variant string // "body", "caption", "label", "code"
    Weight  string // "normal", "medium", "semibold", "bold"
}

type LabelProps struct {
    For      string // HTML for attribute
    Text     string
    Required bool
}
```

**3. Visual Elements**
```go
type IconProps struct {
    Name  string // Icon identifier from icon set
    Size  string // "sm", "md", "lg" - maps to tokens
    Color string // Token reference, not hex
}

type BadgeProps struct {
    Text    string
    Variant string // "success", "warning", "error", "info", "neutral"
    Size    string
}

type SpinnerProps struct {
    Size  string
    Color string // Token reference
}

type DividerProps struct {
    Orientation string // "horizontal", "vertical"
    Spacing     string // Token reference
}

type AvatarProps struct {
    ImageURL string
    Alt      string
    Size     string
    Fallback string // Initials or icon name
}
```

**4. Media**
```go
type ImageProps struct {
    Src         string
    Alt         string
    Width       int
    Height      int
    Loading     string // "lazy", "eager"
    ObjectFit   string // "cover", "contain"
    Placeholder bool   // Show loading state
}
```

### Templ Signatures

```go
// atoms/button.templ
templ Button(props ButtonProps) {
    // Renders button using only design tokens for styling
    // No event handlers defined here
}

// atoms/input.templ
templ Input(props InputProps) {
    // Renders input with token-based styling
    // Error state affects only visual appearance
}

// atoms/icon.templ
templ Icon(props IconProps) {
    // Renders SVG icon from icon system
}

// atoms/badge.templ
templ Badge(props BadgeProps) {
    // Renders badge with variant-based token styling
}
```

### Design Token Integration

**How Atoms Use Tokens:**

```go
// tokens/tokens.go
type DesignTokens struct {
    Colors  ColorTokens
    Spacing SpacingTokens
    Typography TypographyTokens
    Borders BorderTokens
    Shadows ShadowTokens
}

type ColorTokens struct {
    // Primitive
    Blue500 string // "#3B82F6"
    Red600  string // "#DC2626"
    
    // Semantic
    Primary string // "blue500"
    Error   string // "red600"
    
    // Component
    ButtonPrimaryBg     string // "primary"
    ButtonPrimaryText   string // "white"
    ButtonPrimaryHover  string // "blue600"
}

type SpacingTokens struct {
    Unit string // "4px"
    XS   string // "4px"
    SM   string // "8px"
    MD   string // "16px"
    LG   string // "24px"
    XL   string // "32px"
}
```

**In Templ:**
```html
<!-- atoms/button.templ -->
<button 
  class={ buttonClass(props.Variant, props.Size) }
  disabled?={ props.Disabled }
  type={ props.Type }
>
  @if props.Loading {
    @Icon(IconProps{Name: "spinner", Size: "sm"})
  }
  { props.Label }
</button>

<!-- The buttonClass function returns token-based class names -->
<!-- Never inline styles or hard-coded colors -->
```

### State Management in Atoms

Atoms should be **stateless** or have minimal **presentational state only**.

**Acceptable State:**
- Hover (CSS pseudo-class)
- Focus (CSS pseudo-class)
- Active (CSS pseudo-class)
- Disabled (passed as prop)

**Unacceptable State:**
- Form values
- Validation errors
- API call status
- User permissions

### Common Atom Mistakes

❌ **Adding business logic:**
```go
// WRONG - Button knows about permissions
type ButtonProps struct {
    RequiresPermission string
    UserRole string
}
```

✅ **Keep it presentational:**
```go
// CORRECT - Visibility determined by parent
type ButtonProps struct {
    Disabled bool // Parent decides based on permissions
}
```

❌ **Context-specific atoms:**
```go
// WRONG - Too specific
templ SaveInvoiceButton()

// CORRECT - Reusable
templ Button(props ButtonProps)
// Parent passes Label: "Save Invoice"
```

---

## Level 2: Molecules

### Definition

Simple groups of 2-5 atoms that work together to serve a single, focused purpose. Molecules introduce basic interactivity and simple state management.

**Key Characteristics:**
- Combines 2-5 atoms (occasionally more)
- Single, clear responsibility
- Minimal business logic
- Reusable across multiple contexts
- Can maintain simple local UI state
- Self-contained interaction patterns

### Molecule vs. Organism Boundary

**It's a Molecule if:**
- Purpose fits in one sentence
- Used across many different features
- Logic is simple and predictable
- No complex state machines
- No API calls

**It's an Organism if:**
- Purpose requires multiple sentences
- Feature-specific or domain-specific
- Complex business rules
- Multi-step workflows
- Connects to data sources

### Core Molecule Patterns

**1. Form Field Group**

Combines label, input, help text, and error message into a standardized form field.

```go
type FormFieldProps struct {
    Label       string
    Name        string
    Type        InputType
    Value       string
    Placeholder string
    Required    bool
    Disabled    bool
    HelpText    string
    ErrorMsg    string  // Empty string = no error
    ShowError   bool    // Controls error visibility
}

// molecules/form-field.templ
templ FormField(props FormFieldProps) {
    // Renders:
    // - Label atom (with required indicator if needed)
    // - Input atom
    // - Help text (if provided)
    // - Error message (if ShowError && ErrorMsg != "")
}
```

**Why it's a Molecule:**
- Always used together as a unit
- Simple validation display logic
- No business rules, just presentation of validation state
- Reused across all forms

**2. Search Bar**

Combines input, search icon, and clear button.

```go
type SearchBarProps struct {
    Placeholder string
    Value       string
    Name        string
    ShowClear   bool    // Typically: Value != ""
    Disabled    bool
    AutoFocus   bool
}

// molecules/search-bar.templ
templ SearchBar(props SearchBarProps) {
    // Renders:
    // - Icon atom (magnifying glass)
    // - Input atom
    // - Icon button atom (X) if ShowClear
}
```

**Why it's a Molecule:**
- Single purpose: capture search query
- Simple conditional rendering (show/hide clear button)
- Used across many features
- No search execution logic

**3. Quantity Selector**

Number input with increment/decrement buttons.

```go
type QuantitySelectorProps struct {
    Value    int
    Name     string
    Min      int
    Max      int
    Step     int
    Disabled bool
}

// molecules/quantity-selector.templ
templ QuantitySelector(props QuantitySelectorProps) {
    // Renders:
    // - Button atom (minus icon)
    // - Input atom (number type)
    // - Button atom (plus icon)
    // 
    // Buttons disabled when at min/max
}
```

**Why it's a Molecule:**
- Common pattern (shopping carts, inventory adjustments)
- Simple constraint logic (min/max)
- No business rules about what the quantity means

**4. Stat Display**

Shows a metric with label, value, and optional trend.

```go
type TrendDirection string
const (
    TrendUp   TrendDirection = "up"
    TrendDown TrendDirection = "down"
    TrendFlat TrendDirection = "flat"
)

type StatDisplayProps struct {
    Label     string
    Value     string      // Pre-formatted
    Icon      string      // Optional icon name
    Trend     TrendDirection
    TrendValue string     // e.g., "+12%"
    Variant   string      // "default", "success", "warning", "danger"
}

// molecules/stat-display.templ
templ StatDisplay(props StatDisplayProps) {
    // Renders:
    // - Icon atom (if provided)
    // - Label text atom
    // - Value text atom (larger, prominent)
    // - Trend indicator (icon + value) if trend != flat
}
```

**Why it's a Molecule:**
- Pure presentation of a metric
- No calculation logic (receives formatted value)
- Reused in dashboards, reports, cards

**5. Media Object**

Image/icon with title and description.

```go
type MediaObjectProps struct {
    ImageURL    string
    ImageAlt    string
    Icon        string      // Alternative to image
    Title       string
    Description string
    ImageSize   string      // "sm", "md", "lg"
}

// molecules/media-object.templ
templ MediaObject(props MediaObjectProps) {
    // Renders:
    // - Image or Icon atom
    // - Title (heading atom)
    // - Description (text atom)
    // 
    // Common in lists, cards, notifications
}
```

**6. Tag Input**

Shows selected tags with remove functionality.

```go
type Tag struct {
    ID    string
    Label string
}

type TagInputProps struct {
    Tags []Tag
    Name string
}

// molecules/tag-input.templ
templ TagInput(props TagInputProps) {
    // Renders:
    // - Badge atoms for each tag
    // - Each badge has remove icon button
    // - Hidden inputs for form submission
}
```

### Local State in Molecules

Molecules can maintain simple UI state:

**Acceptable State:**
- Input field current value (before submission)
- Dropdown open/closed
- Tooltip visible/hidden
- File upload progress percentage

**Logic Should Be:**
- Deterministic
- Synchronous
- Presentational
- Reversible without side effects

### Event Handling

Molecules can have simple event handlers:

```go
// In the parent component that uses the molecule:
<form hx-post="/search" hx-target="#results">
  @SearchBar(SearchBarProps{
    Name: "q",
    Placeholder: "Search products...",
  })
</form>

// The SearchBar molecule doesn't know about HTMX
// It just renders a form field that works with the form
```

**Molecules don't:**
- Make API calls
- Update global state
- Navigate to other pages
- Transform business data

**Molecules do:**
- Render atoms in a pattern
- Show/hide elements based on props
- Provide form fields with names for submission
- Display validation states passed from parent

### Templ Signatures

```go
// molecules/form-field.templ
templ FormField(props FormFieldProps) {}

// molecules/search-bar.templ
templ SearchBar(props SearchBarProps) {}

// molecules/quantity-selector.templ
templ QuantitySelector(props QuantitySelectorProps) {}

// molecules/stat-display.templ
templ StatDisplay(props StatDisplayProps) {}

// molecules/media-object.templ
templ MediaObject(props MediaObjectProps) {}

// molecules/tag-input.templ
templ TagInput(props TagInputProps) {}

// molecules/breadcrumb-item.templ
templ BreadcrumbItem(props BreadcrumbItemProps) {}

// molecules/avatar-with-name.templ
templ AvatarWithName(props AvatarWithNameProps) {}
```

### When to Create a Molecule

✅ **Create when:**
- You find yourself using the same 2-5 atoms together 3+ times
- The grouping has a clear, single purpose
- The pattern is recognizable to users (search bar, form field)
- You want to enforce consistency

❌ **Don't create when:**
- The combination is one-off or unique
- The atoms need to be independently configurable
- The logic is complex (use organism)
- You're just grouping for layout convenience (use template)

### Common Molecule Examples for ERP

```go
// Currency Display with Icon
type CurrencyDisplayProps struct {
    Amount   decimal.Decimal
    Currency string // "KES", "USD", "EUR"
    ShowIcon bool
    Size     string
}

// Date Range Picker
type DateRangeProps struct {
    StartDate time.Time
    EndDate   time.Time
    Name      string
    MinDate   time.Time
    MaxDate   time.Time
}

// Status Badge with Timestamp
type StatusBadgeProps struct {
    Status    string // "pending", "approved", "rejected"
    Timestamp time.Time
    ShowTime  bool
}

// File Upload Button with Preview
type FileUploadProps struct {
    Name         string
    Accept       string // "image/*", ".pdf"
    MaxSize      int64
    CurrentFile  string // Filename if already uploaded
    PreviewURL   string
}
```

---

## Level 3: Organisms

### Definition

Complex, self-contained UI sections that combine molecules, atoms, and sometimes other organisms. Organisms implement business logic, connect to data, and form distinct, recognizable sections of your interface.

**Key Characteristics:**
- High complexity and specificity
- Contains business logic and rules
- May fetch or transform data
- Forms a complete, recognizable interface section
- Medium reusability (used within a feature domain)
- Can have multiple responsibilities within a bounded context

### Organism vs. Template Boundary

**It's an Organism if:**
- It implements functionality
- It contains business logic
- It manages complex state
- It can be extracted and used in different page layouts
- It's recognizable as a distinct UI pattern (navigation bar, data table)

**It's a Template if:**
- It defines structure without implementing functionality
- It arranges organisms spatially
- It handles responsive layout logic
- It's essentially a wireframe with slots

### Core Organism Categories

**1. Navigation Systems**

```go
// Header Navigation
type NavItem struct {
    Label    string
    URL      string
    Icon     string
    Active   bool
    Badge    string // Notification count, etc.
    Children []NavItem
}

type HeaderNavigationProps struct {
    Logo        LogoConfig
    NavItems    []NavItem
    SearchBar   bool
    UserMenu    UserMenuConfig
    Notifications []Notification
    CurrentUser User
}

// organisms/header-navigation.templ
templ HeaderNavigation(props HeaderNavigationProps) {
    // Composes:
    // - Logo (link atom)
    // - Navigation items (molecules for each)
    // - Search bar molecule
    // - Notification bell with dropdown
    // - User menu organism
    //
    // Business Logic:
    // - Highlight active nav item
    // - Show notification count
    // - Handle mobile menu toggle
}

// Sidebar Navigation
type SidebarNavigationProps struct {
    MenuSections []MenuSection
    Collapsed    bool
    CurrentPath  string
}

type MenuSection struct {
    Title string
    Items []NavItem
    Icon  string
}

// organisms/sidebar-navigation.templ
templ SidebarNavigation(props SidebarNavigationProps) {
    // Business Logic:
    // - Expand/collapse sections
    // - Highlight active item based on CurrentPath
    // - Show/hide labels when collapsed
    // - Implement permission-based visibility
}
```

**2. Data Display**

```go
// Data Table - Most complex organism in most ERP systems
type ColumnDefinition struct {
    Key        string
    Label      string
    Sortable   bool
    Filterable bool
    Width      string
    Align      string // "left", "center", "right"
    Format     string // "currency", "date", "number", "text"
    Render     string // Custom templ component name
}

type SortConfig struct {
    Column    string
    Direction string // "asc", "desc"
}

type FilterConfig struct {
    Column   string
    Operator string // "equals", "contains", "gt", "lt"
    Value    any
}

type PaginationConfig struct {
    Page      int
    PageSize  int
    Total     int
    PageSizes []int // [10, 25, 50, 100]
}

type BulkAction struct {
    ID    string
    Label string
    Icon  string
    Variant string
}

type DataTableProps struct {
    Columns       []ColumnDefinition
    Rows          []map[string]any // Or specific type
    Sortable      bool
    Filterable    bool
    Selectable    bool
    Sort          *SortConfig
    Filters       []FilterConfig
    Pagination    *PaginationConfig
    BulkActions   []BulkAction
    SelectedRows  []string // Row IDs
    Loading       bool
    EmptyMessage  string
}

// organisms/data-table.templ
templ DataTable(props DataTableProps) {
    // Composes:
    // - Table header with sortable columns
    // - Filter row (if filterable)
    // - Checkbox column (if selectable)
    // - Data rows with formatted cells
    // - Pagination controls
    // - Bulk action toolbar (if selections exist)
    // - Loading overlay
    // - Empty state
    //
    // Business Logic:
    // - Sort state management
    // - Filter state management
    // - Selection state management
    // - Cell formatting based on column type
    // - Row expansion (if supported)
}
```

**3. Forms and Input Collections**

```go
// Multi-Section Form
type FormSection struct {
    ID          string
    Title       string
    Description string
    Fields      []FormField
    Collapsible bool
    Expanded    bool
}

type FormField struct {
    Name        string
    Label       string
    Type        string
    Required    bool
    Validation  ValidationRule
    Options     []SelectOption // For select/radio
    Conditional *ConditionalRule // Show if condition met
}

type ValidationRule struct {
    Type    string // "required", "min", "max", "pattern", "custom"
    Value   any
    Message string
}

type ConditionalRule struct {
    Field    string
    Operator string
    Value    any
}

type MultiSectionFormProps struct {
    Sections      []FormSection
    Values        map[string]any
    Errors        map[string]string
    Disabled      bool
    ShowRequiredIndicator bool
}

// organisms/multi-section-form.templ
templ MultiSectionForm(props MultiSectionFormProps) {
    // Business Logic:
    // - Render fields based on type
    // - Show/hide fields based on conditionals
    // - Display validation errors
    // - Section expand/collapse
    // - Required field indicators
}

// Filter Panel
type FilterGroup struct {
    ID     string
    Label  string
    Type   string // "select", "range", "date-range", "checkbox-list"
    Options []FilterOption
    Value  any
}

type FilterOption struct {
    Value string
    Label string
    Count int // Number of results with this value
}

type FilterPanelProps struct {
    Groups        []FilterGroup
    ActiveFilters map[string]any
    SavedFilters  []SavedFilter
    ShowSave      bool
}

type SavedFilter struct {
    ID      string
    Name    string
    Filters map[string]any
}

// organisms/filter-panel.templ
templ FilterPanel(props FilterPanelProps) {
    // Business Logic:
    // - Apply filters via HTMX
    // - Clear all filters
    // - Save current filter set
    // - Load saved filter
    // - Show active filter count
}
```

**4. Content Display Cards**

```go
// Product Card
type ProductCardProps struct {
    Product      Product
    ShowActions  bool
    Compact      bool
    UserRole     string // For permission-based actions
}

type Product struct {
    ID          string
    Name        string
    SKU         string
    Description string
    ImageURL    string
    Price       decimal.Decimal
    Currency    string
    Stock       int
    Category    string
    Tags        []string
}

// organisms/product-card.templ
templ ProductCard(props ProductCardProps) {
    // Composes:
    // - Image atom
    // - Product name (heading)
    // - SKU badge
    // - Price display molecule
    // - Stock indicator
    // - Tags
    // - Action buttons (if ShowActions)
    //
    // Business Logic:
    // - Show/hide actions based on permissions
    // - Stock status color coding
    // - Price formatting by currency
    // - Out of stock handling
}

// Invoice Card
type InvoiceCardProps struct {
    Invoice      Invoice
    ShowDetails  bool
    Permissions  []string
}

type Invoice struct {
    ID          string
    Number      string
    Customer    Customer
    Date        time.Time
    DueDate     time.Time
    Total       decimal.Decimal
    Currency    string
    Status      string // "draft", "sent", "paid", "overdue"
    PaymentStatus string
}

// organisms/invoice-card.templ
templ InvoiceCard(props InvoiceCardProps) {
    // Business Logic:
    // - Status badge coloring
    // - Overdue calculation and highlighting
    // - Permission-based action visibility
    // - Payment status display
}
```

**5. Interactive Widgets**

```go
// Chart Widget
type ChartType string
const (
    ChartLine   ChartType = "line"
    ChartBar    ChartType = "bar"
    ChartPie    ChartType = "pie"
    ChartArea   ChartType = "area"
)

type ChartDataset struct {
    Label string
    Data  []float64
    Color string
}

type ChartWidgetProps struct {
    Title       string
    Type        ChartType
    Labels      []string
    Datasets    []ChartDataset
    ShowLegend  bool
    ShowExport  bool
    Height      string
    Loading     bool
}

// organisms/chart-widget.templ
templ ChartWidget(props ChartWidgetProps) {
    // Renders chart using JS library
    // Includes export functionality
    // Loading and empty states
}

// Calendar Widget
type CalendarEvent struct {
    ID        string
    Title     string
    Start     time.Time
    End       time.Time
    Type      string
    Color     string
}

type CalendarWidgetProps struct {
    Month    time.Time
    Events   []CalendarEvent
    OnSelect string // HTMX endpoint
}

// organisms/calendar-widget.templ
templ CalendarWidget(props CalendarWidgetProps) {
    // Business Logic:
    // - Day grid rendering
    // - Event placement
    // - Month navigation
    // - Event selection
}
```

**6. Modal Dialogs**

```go
type ModalSize string
const (
    ModalSmall  ModalSize = "sm"
    ModalMedium ModalSize = "md"
    ModalLarge  ModalSize = "lg"
    ModalFull   ModalSize = "full"
)

type ModalProps struct {
    ID          string
    Title       string
    Size        ModalSize
    Closable    bool
    CloseOnEscape bool
    CloseOnBackdrop bool
}

// organisms/modal.templ
templ Modal(props ModalProps) {
    // Wrapper organism that provides modal structure
    // Content passed as children
    // Manages open/close state via HTMX or Alpine.js
}

// Confirmation Dialog
type ConfirmDialogProps struct {
    Title       string
    Message     string
    ConfirmText string
    CancelText  string
    Variant     string // "danger", "warning", "info"
    Action      string // HTMX endpoint or action
}

// organisms/confirm-dialog.templ
templ ConfirmDialog(props ConfirmDialogProps) {
    // Uses Modal organism
    // Adds specific confirmation UI
}
```

### State Management in Organisms

Organisms handle complex state:

**Business State:**
- Form values and validation
- Selected items
- Sort/filter configurations
- Expanded/collapsed sections
- Current tab/step in multi-step process

**UI State:**
- Loading indicators
- Error messages
- Success confirmations
- Intermediate calculation results

**State Patterns:**

```go
// Option 1: Server-driven state via HTMX
// The organism renders based on props from server
// User interactions trigger server requests
// Server returns updated HTML

// Option 2: Client-driven state via Alpine.js
// The organism uses x-data for local state
// Useful for purely UI interactions

// Option 3: Hybrid
// Complex state managed server-side
// UI niceties (dropdowns, tooltips) client-side
```

### Templ Signatures

```go
// organisms/header-navigation.templ
templ HeaderNavigation(props HeaderNavigationProps) {}

// organisms/sidebar-navigation.templ
templ SidebarNavigation(props SidebarNavigationProps) {}

// organisms/data-table.templ
templ DataTable(props DataTableProps) {}

// organisms/multi-section-form.templ
templ MultiSectionForm(props MultiSectionFormProps) {}

// organisms/filter-panel.templ
templ FilterPanel(props FilterPanelProps) {}

// organisms/product-card.templ
templ ProductCard(props ProductCardProps) {}

// organisms/invoice-card.templ
templ InvoiceCard(props InvoiceCardProps) {}

// organisms/chart-widget.templ
templ ChartWidget(props ChartWidgetProps) {}

// organisms/calendar-widget.templ
templ CalendarWidget(props CalendarWidgetProps) {}

// organisms/modal.templ
templ Modal(props ModalProps) {}

// organisms/user-menu.templ
templ UserMenu(props UserMenuProps) {}

// organisms/notification-list.templ
templ NotificationList(props NotificationListProps) {}
```

### When to Create an Organism

✅ **Create when:**
- Component has clear business purpose
- Multiple molecules need coordination
- Complex state management required
- Feature-specific logic involved
- Component forms recognizable UI pattern

❌ **Don't create when:**
- Simple composition of atoms/molecules suffices
- No business logic required
- Component is too page-specific (might be template)
- Just doing layout (definitely template)

---

## Level 4: Templates

### Definition

Page-level structures that define the spatial arrangement of organisms without specific content. Templates are essentially wireframes that establish information architecture and responsive behavior.

**Key Characteristics:**
- Define structure and layout
- No real content, only placeholders or props
- Reusable across similar page types
- Handle responsive layout logic
- Focus on information architecture
- Provide slots or regions for organisms

### Template vs. Page Boundary

**It's a Template if:**
- Works with any data of the right type
- Defines HOW things are arranged
- Reusable across multiple routes
- Contains placeholder/example content

**It's a Page if:**
- Tied to specific route/URL
- Fetches specific data
- Contains real business data
- One instance per route

### Core Template Patterns

**1. Master-Detail Layout**

Common in list-to-detail flows (invoices, customers, products).

```go
type MasterDetailTemplateProps struct {
    MasterWidth    string // "320px", "25%", "33%"
    Collapsible    bool
    InitialState   string // "expanded", "collapsed"
    StickyHeader   bool
}

// templates/master-detail.templ
templ MasterDetailTemplate(props MasterDetailTemplateProps) {
    // Structure:
    // ┌─────────────┬──────────────────────┐
    // │   Master    │      Detail          │
    // │   List      │      Panel           │
    // │             │                      │
    // │  Sidebar    │   Main Content       │
    // │  (Organism) │   (Organism)         │
    // │             │                      │
    // │             │                      │
    // └─────────────┴──────────────────────┘
    //
    // Responsive:
    // Mobile: Stacked, master hides when detail shown
    // Tablet+: Side-by-side
    //
    // Slots:
    // - Master: List organism
    // - Detail: Detail organism
    // - Toolbar: Optional action toolbar
}
```

**Usage Example:**
```go
// Page uses template with actual organisms
@MasterDetailTemplate(MasterDetailTemplateProps{
    MasterWidth: "320px",
    Collapsible: true,
}) {
    @slot("master") {
        @InvoiceList(InvoiceListProps{...})
    }
    @slot("detail") {
        @InvoiceDetail(InvoiceDetailProps{...})
    }
    @slot("toolbar") {
        @InvoiceActions(InvoiceActionsProps{...})
    }
}
```

**2. Dashboard Layout**

Grid-based layout for widgets and cards.

```go
type DashboardTemplateProps struct {
    Columns      int    // Desktop columns: 12, 16, 24
    Gap          string // Token reference
    SidebarWidth string
    HasSidebar   bool
}

// templates/dashboard.templ
templ DashboardTemplate(props DashboardTemplateProps) {
    // Structure:
    // ┌─────────┬────────────────────────────┐
    // │ Sidebar │  Widget Grid               │
    // │ (opt.)  │  ┌──────┬──────┬──────┐    │
    // │         │  │  W1  │  W2  │  W3  │    │
    // │         │  ├──────┴──────┼──────┤    │
    // │         │  │     W4      │  W5  │    │
    // │         │  └─────────────┴──────┘    │
    // └─────────┴────────────────────────────┘
    //
    // Responsive:
    // Mobile: Single column, sidebar becomes drawer
    // Tablet: 2 columns
    // Desktop: N columns based on props
    //
    // Slots:
    // - Sidebar: Navigation or filters
    // - Widgets: Array of positioned organisms
}
```

**3. Form Layout**

Structured layout for create/edit forms.

```go
type FormLayoutVariant string
const (
    FormSingle   FormLayoutVariant = "single"   // One column
    FormSplit    FormLayoutVariant = "split"    // Two columns
    FormWizard   FormLayoutVariant = "wizard"   // Multi-step
    FormSections FormLayoutVariant = "sections" // Collapsible sections
)

type FormTemplateProps struct {
    Variant       FormLayoutVariant
    StickyActions bool // Actions bar stays at bottom
    ShowProgress  bool // For wizard variant
    BreadcrumbNav bool
}

// templates/form-layout.templ
templ FormTemplate(props FormTemplateProps) {
    // Structure:
    // ┌────────────────────────────────┐
    // │  Breadcrumbs (optional)        │
    // ├────────────────────────────────┤
    // │  Page Header                   │
    // │  Title, Description, etc.      │
    // ├────────────────────────────────┤
    // │  Form Content (organism)       │
    // │                                │
    // │  (Layout varies by variant)    │
    // │                                │
    // ├────────────────────────────────┤
    // │  Sticky Actions Bar            │
    // │  [Cancel] [Save Draft] [Submit]│
    // └────────────────────────────────┘
    //
    // Responsive:
    // Split variant becomes single column on mobile
    //
    // Slots:
    // - Header: Page title organism
    // - Content: Form organism
    // - Actions: Action buttons
    // - Progress: For wizard (optional)
}
```

**4. Report Layout**

For data analysis and reporting pages.

```go
type ReportTemplateProps struct {
    FiltersPosition string // "sidebar", "top", "drawer"
    FiltersWidth    string
    ChartsFirst     bool   // Charts before table or after
}

// templates/report-layout.templ
templ ReportTemplate(props ReportTemplateProps) {
    // Structure:
    // ┌────────┬─────────────────────────┐
    // │Filters │  Summary Stats          │
    // │Panel   │  (Stats organisms)      │
    // │        ├─────────────────────────┤
    // │        │  Chart                  │
    // │        │  (Chart organism)       │
    // │        ├─────────────────────────┤
    // │        │  Data Table             │
    // │        │  (Table organism)       │
    // │        ├─────────────────────────┤
    // │        │  Export Toolbar         │
    // └────────┴─────────────────────────┘
    //
    // Slots:
    // - Filters: Filter panel organism
    // - Stats: Array of stat organisms
    // - Charts: Array of chart organisms
    // - Table: Data table organism
    // - Toolbar: Export/action toolbar
}
```

**5. Content Layout**

For documentation, articles, help pages.

```go
type ContentTemplateProps struct {
    HasSidebar    bool
    SidebarSticky bool
    TableOfContents bool
    MaxWidth      string // Content max-width
}

// templates/content-layout.templ
templ ContentTemplate(props ContentTemplateProps) {
    // Structure:
    // ┌──────────────────┬───────────────┐
    // │   Main Content   │   Sidebar     │
    // │                  │   - TOC       │
    // │   Article        │   - Related   │
    // │   Body           │   - Links     │
    // │                  │               │
    // └──────────────────┴───────────────┘
    //
    // Responsive:
    // Mobile: Sidebar moves to bottom
    // Tablet+: Side-by-side
}
```

### Layout Primitives

Templates use layout primitives for consistent spacing:

```go
type ContainerProps struct {
    MaxWidth string // "sm", "md", "lg", "xl", "full"
    Padding  bool
}

type StackProps struct {
    Direction string // "vertical", "horizontal"
    Gap       string // Token reference
    Wrap      bool
}

type GridProps struct {
    Columns      int
    Gap          string
    Responsive   bool
}

// These are reusable layout components used by templates
```

### Responsive Behavior

Templates define responsive rules:

```go
type ResponsiveConfig struct {
    Mobile  LayoutConfig // < 768px
    Tablet  LayoutConfig // 768px - 1024px
    Desktop LayoutConfig // > 1024px
}

type LayoutConfig struct {
    Columns     int
    SidebarPos  string // "top", "left", "right", "hidden"
    Stack       bool
}
```

### Templ Signatures

```go
// templates/master-detail.templ
templ MasterDetailTemplate(props MasterDetailTemplateProps) {}

// templates/dashboard.templ
templ DashboardTemplate(props DashboardTemplateProps) {}

// templates/form-layout.templ
templ FormTemplate(props FormTemplateProps) {}

// templates/report-layout.templ
templ ReportTemplate(props ReportTemplateProps) {}

// templates/content-layout.templ
templ ContentTemplate(props ContentTemplateProps) {}

// templates/auth-layout.templ (centered forms)
templ AuthTemplate(props AuthTemplateProps) {}

// templates/split-view.templ (50/50 split)
templ SplitViewTemplate(props SplitViewTemplateProps) {}

// templates/settings-layout.templ (tabs + content)
templ SettingsTemplate(props SettingsTemplateProps) {}
```

### Template Usage Pattern

```go
// In a page handler
func InvoiceDetailPage(invoice Invoice) templ.Component {
    return MasterDetailTemplate(MasterDetailTemplateProps{
        MasterWidth: "320px",
    }) {
        @InvoiceList(...)  // Master
        @InvoiceDetail(invoice) // Detail
    }
}
```

---

## Level 5: Pages

### Definition

Specific instances of templates populated with real data, connected to routing, and tied to specific URLs. Pages represent what users actually interact with.

**Key Characteristics:**
- Tied to specific route/URL
- Fetches or receives real data
- Implements page-specific business logic
- Coordinates multiple organisms
- Handles authentication/authorization
- Manages page-level state

### Page Responsibilities

**1. Data Management**

```go
type InvoiceDetailPageData struct {
    Invoice      Invoice
    LineItems    []LineItem
    Payments     []Payment
    Customer     Customer
    User         User
    Permissions  []string
}

// pages/invoice-detail.go
func InvoiceDetailPage(w http.ResponseWriter, r *http.Request) {
    invoiceID := chi.URLParam(r, "id")
    
    // Fetch data
    data, err := fetchInvoiceData(r.Context(), invoiceID)
    if err != nil {
        // Handle error
        return
    }
    
    // Render page
    component := invoiceDetailPageTemplate(data)
    component.Render(r.Context(), w)
}
```

**2. Routing Context**

```go
// URL: /invoices/123?tab=payments&sort=date

type InvoiceDetailPageProps struct {
    InvoiceID string        // From URL path
    ActiveTab string        // From query param
    Sort      string        // From query param
    User      User          // From session
    Flash     FlashMessage  // From session
}
```

**3. State Coordination**

```go
// Page manages state between organisms

type DashboardPageState struct {
    DateRange    DateRange    // Affects all widgets
    Filters      FilterSet    // Affects table and charts
    SelectedPump string       // Affects detail panel
}

// When user changes date range:
// 1. Update URL query params
// 2. Fetch new data for all widgets
// 3. Re-render affected organisms
```

**4. Authorization**

```go
func InvoiceEditPage(w http.ResponseWriter, r *http.Request) {
    user := GetUserFromContext(r.Context())
    invoice, _ := fetchInvoice(...)
    
    // Page-level permission check
    if !user.Can("invoice.edit", invoice) {
        http.Error(w, "Forbidden", 403)
        return
    }
    
    // Pass permissions to organisms
    props := InvoiceEditPageProps{
        Invoice: invoice,
        CanEdit: true,
        CanDelete: user.Can("invoice.delete", invoice),
        CanApprove: user.Can("invoice.approve", invoice),
    }
    
    // Organisms use these flags to show/hide actions
    component := invoiceEditPageTemplate(props)
    component.Render(r.Context(), w)
}
```

### Page Types and Examples

**1. List Pages**

```go
type InvoiceListPageProps struct {
    Invoices    []Invoice
    Pagination  PaginationState
    Filters     ActiveFilters
    Sort        SortState
    User        User
}

// pages/invoice-list.templ
templ InvoiceListPage(props InvoiceListPageProps) {
    @MasterDetailTemplate(...) {
        @slot("filters") {
            @InvoiceFilters(props.Filters)
        }
        @slot("list") {
            @InvoiceTable(InvoiceTableProps{
                Invoices: props.Invoices,
                Sort: props.Sort,
            })
        }
        @slot("pagination") {
            @Pagination(props.Pagination)
        }
    }
}
```

**2. Detail Pages**

```go
type ProductDetailPageProps struct {
    Product     Product
    Variants    []ProductVariant
    Suppliers   []Supplier
    Inventory   InventoryStatus
    ActiveTab   string
    User        User
    Permissions []string
}

// pages/product-detail.templ
templ ProductDetailPage(props ProductDetailPageProps) {
    @DashboardTemplate(...) {
        @slot("header") {
            @ProductHeader(props.Product, props.Permissions)
        }
        @slot("tabs") {
            @TabNavigation([]Tab{
                {ID: "details", Label: "Details"},
                {ID: "variants", Label: "Variants"},
                {ID: "inventory", Label: "Inventory"},
                {ID: "suppliers", Label: "Suppliers"},
            }, props.ActiveTab)
        }
        @slot("content") {
            @switch props.ActiveTab {
                @case "details"
                    @ProductDetailsTab(props.Product)
                @case "variants"
                    @ProductVariantsTab(props.Variants)
                @case "inventory"
                    @ProductInventoryTab(props.Inventory)
                @case "suppliers"
                    @ProductSuppliersTab(props.Suppliers)
            }
        }
    }
}
```

**3. Form Pages (Create/Edit)**

```go
type InvoiceCreatePageProps struct {
    Draft       *Invoice // If resuming draft
    Customers   []Customer
    Products    []Product
    Errors      ValidationErrors
    Step        int // For wizard
    User        User
}

// pages/invoice-create.templ
templ InvoiceCreatePage(props InvoiceCreatePageProps) {
    @FormTemplate(FormTemplateProps{
        Variant: FormWizard,
        ShowProgress: true,
    }) {
        @slot("progress") {
            @WizardProgress(InvoiceWizardSteps, props.Step)
        }
        @slot("content") {
            @switch props.Step {
                @case 1
                    @InvoiceCustomerStep(props.Customers, props.Draft)
                @case 2
                    @InvoiceLineItemsStep(props.Products, props.Draft)
                @case 3
                    @InvoiceReviewStep(props.Draft)
            }
        }
        @slot("actions") {
            @WizardActions(props.Step, len(InvoiceWizardSteps))
        }
    }
}
```

**4. Dashboard Pages**

```go
type StationDashboardPageProps struct {
    Station     Station
    Pumps       []PumpStatus
    Tanks       []TankLevel
    Sales       SalesData
    Alerts      []Alert
    DateRange   DateRange
    User        User
    RefreshRate int // Seconds
}

// pages/station-dashboard.templ
templ StationDashboardPage(props StationDashboardPageProps) {
    @DashboardTemplate(...) {
        @slot("header") {
            @DashboardHeader(props.Station, props.DateRange)
        }
        @slot("widgets") {
            <div class="grid grid-cols-12 gap-4">
                <div class="col-span-8">
                    @PumpStatusGrid(props.Pumps)
                </div>
                <div class="col-span-4">
                    @TankLevelPanel(props.Tanks)
                </div>
                <div class="col-span-6">
                    @SalesChart(props.Sales)
                </div>
                <div class="col-span-6">
                    @AlertList(props.Alerts)
                </div>
            </div>
        }
    }
}
```

**5. Report Pages**

```go
type SalesReportPageProps struct {
    Report      SalesReport
    Filters     ReportFilters
    Chart       ChartData
    Table       []ReportRow
    Summary     ReportSummary
    ExportURL   string
}

// pages/sales-report.templ
templ SalesReportPage(props SalesReportPageProps) {
    @ReportTemplate(...) {
        @slot("filters") {
            @ReportFilters(props.Filters)
        }
        @slot("summary") {
            @ReportSummary(props.Summary)
        }
        @slot("chart") {
            @SalesChart(props.Chart)
        }
        @slot("table") {
            @ReportTable(props.Table)
        }
        @slot("export") {
            @ExportToolbar(props.ExportURL)
        }
    }
}
```

### Page Data Flow Patterns

**1. Server-Side Rendering (Templ + HTMX)**

```go
// Initial page load
func InvoiceListHandler(w http.ResponseWriter, r *http.Request) {
    // Parse query params
    filters := parseFilters(r.URL.Query())
    page := parsePage(r.URL.Query())
    
    // Fetch data
    invoices, total := fetchInvoices(filters, page)
    
    // Render full page
    props := InvoiceListPageProps{
        Invoices: invoices,
        Pagination: PaginationState{Page: page, Total: total},
        Filters: filters,
    }
    
    InvoiceListPage(props).Render(r.Context(), w)
}

// Partial updates via HTMX
func InvoiceTablePartial(w http.ResponseWriter, r *http.Request) {
    filters := parseFilters(r.URL.Query())
    page := parsePage(r.URL.Query())
    
    invoices, total := fetchInvoices(filters, page)
    
    // Render just the table organism
    InvoiceTable(InvoiceTableProps{
        Invoices: invoices,
    }).Render(r.Context(), w)
}
```

**2. Real-Time Updates (WebSocket/SSE)**

```go
type DashboardPageProps struct {
    // ... other props
    WebSocketURL string // Client connects for updates
}

// Server pushes updates
// Client receives and swaps specific organisms via HTMX
```

### Page Handler Signature Pattern

```go
// Standard page handler
type PageHandler func(w http.ResponseWriter, r *http.Request)

// pages/handlers.go
func InvoiceListPageHandler(db *sql.DB) PageHandler {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Extract context (user, tenant)
        user := middleware.GetUser(r.Context())
        
        // 2. Parse URL params/query
        filters := parseInvoiceFilters(r.URL.Query())
        
        // 3. Fetch data
        invoices, err := fetchInvoices(r.Context(), db, user.TenantID, filters)
        if err != nil {
            // Handle error
            return
        }
        
        // 4. Build props
        props := InvoiceListPageProps{
            Invoices: invoices,
            Filters: filters,
            User: user,
        }
        
        // 5. Render
        InvoiceListPage(props).Render(r.Context(), w)
    }
}
```

### Templ Signatures

```go
// pages/invoice-list.templ
templ InvoiceListPage(props InvoiceListPageProps) {}

// pages/invoice-detail.templ
templ InvoiceDetailPage(props InvoiceDetailPageProps) {}

// pages/invoice-create.templ
templ InvoiceCreatePage(props InvoiceCreatePageProps) {}

// pages/invoice-edit.templ
templ InvoiceEditPage(props InvoiceEditPageProps) {}

// pages/dashboard.templ
templ DashboardPage(props DashboardPageProps) {}

// pages/station-dashboard.templ
templ StationDashboardPage(props StationDashboardPageProps) {}

// pages/product-detail.templ
templ ProductDetailPage(props ProductDetailPageProps) {}

// pages/sales-report.templ
templ SalesReportPage(props SalesReportPageProps) {}

// pages/customer-profile.templ
templ CustomerProfilePage(props CustomerProfilePageProps) {}
```

---

## Composition Principles

### Unidirectional Dependency Flow

```
Pages
  ├─ Can import Templates
  ├─ Can import Organisms
  ├─ Can import Molecules
  └─ Can import Atoms

Templates
  ├─ Can import Organisms
  ├─ Can import Molecules
  └─ Can import Atoms

Organisms
  ├─ Can import Molecules
  ├─ Can import Atoms
  └─ Can import other Organisms (carefully)

Molecules
  ├─ Can import Atoms
  └─ Can import other Molecules (rarely)

Atoms
  └─ Can ONLY import Design Tokens
```

### Slot-Based Composition

Templates and some organisms use named slots:

```go
// In template
templ DashboardTemplate(props DashboardTemplateProps) {
    <div class="dashboard-layout">
        <header>
            { children... } // Named slot: "header"
        </header>
        <aside>
            { children... } // Named slot: "sidebar"
        </aside>
        <main>
            { children... } // Named slot: "main"
        </main>
    </div>
}

// Usage (conceptual - Templ syntax may vary)
@DashboardTemplate(props) {
    @slot("header") {
        @DashboardHeader(...)
    }
    @slot("sidebar") {
        @DashboardNav(...)
    }
    @slot("main") {
        @DashboardWidgets(...)
    }
}
```

### Props-Down, Events-Up Pattern

```go
// Data flows DOWN through props
Page {
    props.Invoices → Template {
        props.Invoices → Organism {
            props.Invoice → Molecule {
                props.Invoice.Number → Atom
            }
        }
    }
}

// Events flow UP through HTMX/Alpine.js
Atom (button clicked)
    ↓ hx-post="/api/invoice/123/approve"
    ↓ Server handler
    ↓ Re-render organism or page
    ↓ New props flow down
```

### Component Communication Patterns

**1. Via Props (Preferred)**
```go
// Parent passes data to child
@InvoiceTable(InvoiceTableProps{
    Invoices: invoices,
    OnSort: "/invoices?sort=%s",
})
```

**2. Via HTMX (Server Coordination)**
```html
<!-- Child triggers server action -->
<button hx-post="/invoices/123/delete" 
        hx-target="#invoice-list"
        hx-swap="outerHTML">
    Delete
</button>
<!-- Server re-renders parent organism with updated data -->
```

**3. Via Alpine.js (Client Coordination)**
```html
<!-- Shared Alpine.js store -->
<div x-data="{ selectedId: null }">
    <!-- List sets selectedId -->
    <div @click="selectedId = '123'">...</div>
    
    <!-- Detail reads selectedId -->
    <div x-show="selectedId === '123'">...</div>
</div>
```

### Avoiding Anti-Patterns

❌ **Circular Dependencies**
```go
// WRONG
// organisms/invoice-list.templ imports invoice-detail
// organisms/invoice-detail.templ imports invoice-list
```

✅ **Lift to Common Parent**
```go
// CORRECT
// Page imports both and coordinates
@InvoiceListPage() {
    @InvoiceList(...)
    @InvoiceDetail(...)
}
```

❌ **Props Drilling Through Many Levels**
```go
// WRONG - passing userId through 5 levels
Page(userId) → Template(userId) → Organism(userId) → Molecule(userId) → Atom(userId)
```

✅ **Use Context for Cross-Cutting Concerns**
```go
// CORRECT - use context or session
// Only pass what each level actually uses
```

---

## Design Tokens Foundation

### What Are Design Tokens?

Design tokens are the foundational layer below atoms—named variables that define your design system's visual properties.

```go
// tokens/tokens.go
package tokens

type Tokens struct {
    Colors     ColorTokens
    Spacing    SpacingTokens
    Typography TypographyTokens
    Borders    BorderTokens
    Shadows    ShadowTokens
    Breakpoints BreakpointTokens
}

type ColorTokens struct {
    // Primitive - raw colors
    Blue500  string // "#3B82F6"
    Red600   string // "#DC2626"
    Gray100  string // "#F3F4F6"
    
    // Semantic - meaning-based
    Primary   string // "blue500"
    Error     string // "red600"
    TextBase  string // "gray900"
    
    // Component - usage-specific
    ButtonPrimaryBg     string // "primary"
    ButtonPrimaryText   string // "white"
    ButtonPrimaryHover  string // "blue600"
    InputBorder         string // "gray300"
    InputBorderFocus    string // "primary"
}

type SpacingTokens struct {
    Unit string // "4px"
    XS   string // "4px"   (1 unit)
    SM   string // "8px"   (2 units)
    MD   string // "16px"  (4 units)
    LG   string // "24px"  (6 units)
    XL   string // "32px"  (8 units)
    XXL  string // "48px"  (12 units)
}

type TypographyTokens struct {
    // Font families
    FontSans  string // "'Inter', sans-serif"
    FontMono  string // "'JetBrains Mono', monospace"
    
    // Font sizes
    TextXS  string // "12px"
    TextSM  string // "14px"
    TextBase string // "16px"
    TextLG  string // "18px"
    TextXL  string // "20px"
    Text2XL string // "24px"
    
    // Font weights
    FontNormal    string // "400"
    FontMedium    string // "500"
    FontSemibold  string // "600"
    FontBold      string // "700"
    
    // Line heights
    LeadingTight  string // "1.25"
    LeadingNormal string // "1.5"
    LeadingRelaxed string // "1.75"
}
```

### Token Hierarchy

```
Primitive Tokens (raw values)
    ↓
Semantic Tokens (meaning)
    ↓
Component Tokens (usage)
    ↓
CSS Variables
    ↓
Atoms use via classes
```

### Generating CSS from Tokens

```go
// tokens/generator.go
func GenerateCSS(tokens Tokens) string {
    var css strings.Builder
    
    css.WriteString(":root {\n")
    
    // Colors
    css.WriteString(fmt.Sprintf("  --color-primary: %s;\n", tokens.Colors.Primary))
    css.WriteString(fmt.Sprintf("  --color-error: %s;\n", tokens.Colors.Error))
    
    // Spacing
    css.WriteString(fmt.Sprintf("  --spacing-xs: %s;\n", tokens.Spacing.XS))
    css.WriteString(fmt.Sprintf("  --spacing-sm: %s;\n", tokens.Spacing.SM))
    
    // Typography
    css.WriteString(fmt.Sprintf("  --font-sans: %s;\n", tokens.Typography.FontSans))
    css.WriteString(fmt.Sprintf("  --text-base: %s;\n", tokens.Typography.TextBase))
    
    css.WriteString("}\n")
    
    return css.String()
}
```

### Multi-Tenant Theming

```go
// Each tenant has custom token values
type TenantTheme struct {
    TenantID string
    Tokens   Tokens
}

// Generate tenant-specific CSS
func GenerateTenantCSS(theme TenantTheme) string {
    var css strings.Builder
    
    css.WriteString(fmt.Sprintf("[data-tenant='%s'] {\n", theme.TenantID))
    css.WriteString(fmt.Sprintf("  --color-primary: %s;\n", theme.Tokens.Colors.Primary))
    // ... more tokens
    css.WriteString("}\n")
    
    return css.String()
}

// In HTML
<html data-tenant="shell-maanzoni">
```

### Atoms Must Use Tokens

```go
// ❌ WRONG - Hard-coded values in atom
templ Button(props ButtonProps) {
    <button style="background: #3B82F6; padding: 8px 16px;">
        { props.Label }
    </button>
}

// ✅ CORRECT - Token-based classes
templ Button(props ButtonProps) {
    <button class={ buttonClasses(props.Variant, props.Size) }>
        { props.Label }
    </button>
}

// buttonClasses returns token-based class names
func buttonClasses(variant ButtonVariant, size ButtonSize) string {
    classes := []string{"btn"} // Base class with token styles
    
    switch variant {
    case ButtonPrimary:
        classes = append(classes, "btn-primary") // Uses --color-primary
    case ButtonSecondary:
        classes = append(classes, "btn-secondary")
    }
    
    switch size {
    case ButtonSmall:
        classes = append(classes, "btn-sm") // Uses --spacing-xs, --text-sm
    case ButtonLarge:
        classes = append(classes, "btn-lg")
    }
    
    return strings.Join(classes, " ")
}
```

---

## Schema-Driven Implementation

### Why Schema-Driven UI?

In ERP systems, you often need:
- Non-developers to customize forms
- Per-tenant UI customization
- Rapid form creation
- Consistent validation rules
- Version control for UI definitions

Schema-driven UI defines your interface through JSON/YAML configuration.

### Schema Structure

```go
// schema/types.go
type PageSchema struct {
    ID          string           `json:"id"`
    Title       string           `json:"title"`
    Description string           `json:"description"`
    Template    string           `json:"template"`
    Sections    []SectionSchema  `json:"sections"`
    Actions     []ActionSchema   `json:"actions"`
    Permissions []string         `json:"permissions"`
}

type SectionSchema struct {
    ID          string        `json:"id"`
    Title       string        `json:"title"`
    Organism    string        `json:"organism"`
    Props       map[string]any `json:"props"`
    Conditional *Condition    `json:"conditional,omitempty"`
}

type ActionSchema struct {
    ID       string   `json:"id"`
    Label    string   `json:"label"`
    Variant  string   `json:"variant"`
    Action   string   `json:"action"` // URL or handler name
    Requires []string `json:"requires"` // Conditions
}

type Condition struct {
    Field    string `json:"field"`
    Operator string `json:"operator"`
    Value    any    `json:"value"`
}

// Form-specific schema
type FormSchema struct {
    ID          string         `json:"id"`
    Fields      []FieldSchema  `json:"fields"`
    Validation  ValidationSchema `json:"validation"`
    Layout      LayoutConfig   `json:"layout"`
}

type FieldSchema struct {
    Name        string        `json:"name"`
    Label       string        `json:"label"`
    Type        string        `json:"type"`
    Required    bool          `json:"required"`
    Placeholder string        `json:"placeholder,omitempty"`
    Options     []FieldOption `json:"options,omitempty"` // For select
    Validation  []ValidationRule `json:"validation,omitempty"`
    Conditional *Condition    `json:"conditional,omitempty"`
}

type ValidationRule struct {
    Type    string `json:"type"`    // "required", "min", "max", "pattern"
    Value   any    `json:"value"`
    Message string `json:"message"`
}
```

### Example Schema

```json
{
  "id": "invoice-create",
  "title": "Create Invoice",
  "template": "form-layout",
  "sections": [
    {
      "id": "customer",
      "title": "Customer Information",
      "organism": "CustomerSelector",
      "props": {
        "required": true,
        "searchable": true
      }
    },
    {
      "id": "lineItems",
      "title": "Line Items",
      "organism": "LineItemsTable",
      "props": {
        "minRows": 1,
        "showTax": true
      }
    }
  ],
  "actions": [
    {
      "id": "save-draft",
      "label": "Save Draft",
      "variant": "secondary",
      "action": "/api/invoices/draft"
    },
    {
      "id": "submit",
      "label": "Create Invoice",
      "variant": "primary",
      "action": "/api/invoices",
      "requires": ["validation:passed"]
    }
  ]
}
```

### Schema Runtime System

```go
// runtime/renderer.go
type SchemaRenderer struct {
    Registry *ComponentRegistry
    Validator *SchemaValidator
}

type ComponentRegistry struct {
    organisms map[string]OrganismFactory
    templates map[string]TemplateFactory
}

type OrganismFactory func(props map[string]any) templ.Component

func (sr *SchemaRenderer) RenderPage(schema PageSchema, data map[string]any) templ.Component {
    // 1. Validate schema
    if err := sr.Validator.Validate(schema); err != nil {
        return errorPage(err)
    }
    
    // 2. Get template
    template := sr.Registry.GetTemplate(schema.Template)
    
    // 3. Build organisms from sections
    organisms := make([]templ.Component, 0, len(schema.Sections))
    for _, section := range schema.Sections {
        // Check conditional
        if section.Conditional != nil {
            if !evaluateCondition(section.Conditional, data) {
                continue
            }
        }
        
        // Get organism factory
        factory := sr.Registry.GetOrganism(section.Organism)
        
        // Create organism instance
        organism := factory(section.Props)
        organisms = append(organisms, organism)
    }
    
    // 4. Render template with organisms
    return template(TemplateProps{
        Title: schema.Title,
        Sections: organisms,
        Actions: schema.Actions,
    })
}

// Registering components
func (r *ComponentRegistry) RegisterOrganism(name string, factory OrganismFactory) {
    r.organisms[name] = factory
}

// Example registration
registry.RegisterOrganism("CustomerSelector", func(props map[string]any) templ.Component {
    return CustomerSelector(CustomerSelectorProps{
        Required: props["required"].(bool),
        Searchable: props["searchable"].(bool),
    })
})
```

### Condition Evaluation

```go
// runtime/conditions.go
func evaluateCondition(cond *Condition, data map[string]any) bool {
    fieldValue := getNestedField(data, cond.Field)
    
    switch cond.Operator {
    case "equals":
        return fieldValue == cond.Value
    case "not_equals":
        return fieldValue != cond.Value
    case "greater_than":
        return compareNumbers(fieldValue, cond.Value, ">")
    case "less_than":
        return compareNumbers(fieldValue, cond.Value, "<")
    case "contains":
        return containsString(fieldValue, cond.Value)
    case "is_empty":
        return isEmpty(fieldValue)
    default:
        return false
    }
}
```

### Benefits

✅ **Configuration Over Code**
- Change forms without recompiling
- A/B test UI layouts
- Per-tenant customization

✅ **Consistency Enforcement**
- All forms use same components
- Validation rules centralized
- Guaranteed accessibility

✅ **Rapid Development**
- New pages in minutes
- Component reuse maximized
- Less Go code to write

✅ **Version Control**
- Schema changes tracked in Git
- Easy rollback
- Review UI changes in PRs

---

## Testing Strategy

### Testing by Level

**Atoms: Unit Tests**

```go
// atoms/button_test.go
func TestButton_Variants(t *testing.T) {
    tests := []struct {
        variant ButtonVariant
        want    string // Expected class
    }{
        {ButtonPrimary, "btn-primary"},
        {ButtonSecondary, "btn-secondary"},
        {ButtonDanger, "btn-danger"},
    }
    
    for _, tt := range tests {
        t.Run(string(tt.variant), func(t *testing.T) {
            classes := buttonClasses(tt.variant, ButtonMedium)
            assert.Contains(t, classes, tt.want)
        })
    }
}

func TestButton_DisabledState(t *testing.T) {
    // Test that disabled prop affects rendering
}

func TestButton_LoadingState(t *testing.T) {
    // Test that loading state shows spinner
}
```

**Molecules: Integration Tests**

```go
// molecules/form-field_test.go
func TestFormField_ShowsError(t *testing.T) {
    props := FormFieldProps{
        Label: "Email",
        Name: "email",
        ErrorMsg: "Invalid email",
        ShowError: true,
    }
    
    // Render and check error message appears
    html := renderToString(FormField(props))
    assert.Contains(t, html, "Invalid email")
}

func TestFormField_HidesErrorWhenShowErrorFalse(t *testing.T) {
    props := FormFieldProps{
        Label: "Email",
        ErrorMsg: "Invalid email",
        ShowError: false, // Should not show
    }
    
    html := renderToString(FormField(props))
    assert.NotContains(t, html, "Invalid email")
}
```

**Organisms: Behavior Tests**

```go
// organisms/data-table_test.go
func TestDataTable_Sorting(t *testing.T) {
    props := DataTableProps{
        Columns: []ColumnDefinition{
            {Key: "name", Label: "Name", Sortable: true},
        },
        Rows: []map[string]any{
            {"name": "Zebra"},
            {"name": "Apple"},
        },
        Sort: &SortConfig{Column: "name", Direction: "asc"},
    }
    
    html := renderToString(DataTable(props))
    
    // Verify rows are sorted
    // Verify sort indicators shown
}

func TestDataTable_EmptyState(t *testing.T) {
    props := DataTableProps{
        Rows: []map[string]any{},
        EmptyMessage: "No records found",
    }
    
    html := renderToString(DataTable(props))
    assert.Contains(t, html, "No records found")
}
```

**Templates: Visual Tests**

```go
// Use Storybook or visual regression tools
// Test responsive behavior at different viewports
// Test with placeholder content
```

**Pages: E2E Tests**

```go
// Use Playwright or similar
func TestInvoiceCreationFlow(t *testing.T) {
    // 1. Navigate to create invoice page
    // 2. Fill in customer
    // 3. Add line items
    // 4. Submit form
    // 5. Verify invoice created
    // 6. Verify redirect to detail page
}
```

### Test Coverage Goals

- **Atoms:** 100% (they're simple and foundational)
- **Molecules:** 90%+ (interaction logic critical)
- **Organisms:** 80%+ (complex business logic)
- **Templates:** Visual regression tests
- **Pages:** Critical paths only (E2E expensive)

---

## Real-World ERP Examples

### Example 1: Fuel Station Pump Dashboard

**Atoms Used:**
- Status indicator (colored circle)
- Number display
- Icon (fuel pump, warning)
- Button (stop, resume)

**Molecules Used:**
- Stat display (volume sold, amount)
- Status badge (active, idle, error)

**Organisms Used:**
- Pump card (pump# + status + current transaction + volume + controls)
- Pump grid (12 pump cards in grid)
- Tank gauge panel (4 tanks with levels)
- Alert list (recent alerts with severity)

**Template:**
- Dashboard layout (header + pump grid + sidebar with tanks/alerts)

**Page:**
- Station dashboard page (connects to WebSocket, renders specific station data)

### Example 2: Invoice Management

**Atoms:**
- Input, Select, Button
- Date picker
- Currency display
- Status badge

**Molecules:**
- Form field (label + input + error)
- Currency input (amount + currency selector)
- Customer search (search bar + results dropdown)

**Organisms:**
- Invoice header (number, date, customer, status)
- Line items table (products, quantities, prices, totals)
- Payment history (list of payments)
- Invoice actions (print, email, approve, void - permission-based)

**Template:**
- Form layout (for create/edit)
- Master-detail layout (for list → detail)

**Pages:**
- Invoice list page (filterable, sortable table)
- Invoice detail page (specific invoice data)
- Invoice create page (multi-step wizard)

### Example 3: Product Catalog

**Atomic Breakdown:**

```go
// Atoms
Button, Input, Select, Image, Badge, Icon

// Molecules  
PricingDisplay (amount + currency + tax indicator)
StockBadge (quantity + unit + status color)
ProductImage (image + zoom + lightbox trigger)
CategoryBreadcrumb (home > category > subcategory)

// Organisms
ProductCard (image + name + SKU + price + stock + actions)
ProductGrid (filterable, sortable grid of cards)
ProductFilters (category, price range, stock status, attributes)
ProductDetailTabs (overview, variants, inventory, suppliers, history)

// Template
MasterDetailTemplate (filters + grid + detail panel)

// Pages
ProductCatalogPage (all products with filters)
ProductDetailPage (single product with tabs)
ProductCreatePage (multi-step form)
```

---

## Migration Guide

### Phase 1: Foundation (Week 1-2)

**Goals:**
- Establish design tokens
- Create 10-15 core atoms
- Set up file structure

**Tasks:**
1. Define token system
2. Generate CSS from tokens
3. Create button, input, select, icon, badge atoms
4. Document in README
5. Get team alignment

### Phase 2: Molecules (Week 3-4)

**Goals:**
- Extract common patterns
- Create 8-10 molecules

**Tasks:**
1. Identify repeated atom combinations
2. Create form field, search bar, stat display
3. Replace usage in 2-3 pages as proof of concept
4. Gather team feedback

### Phase 3: Organisms (Week 5-8)

**Goals:**
- Build feature-specific organisms
- Standardize complex components

**Tasks:**
1. Start with most-used organisms (data table, navigation)
2. Build 5-8 core organisms
3. Use in new features
4. Plan migration of existing features

### Phase 4: Templates (Week 9-10)

**Goals:**
- Standardize page layouts

**Tasks:**
1. Identify layout patterns
2. Create 4-5 core templates
3. Apply to new pages
4. Document template usage

### Phase 5: Adoption (Ongoing)

**Rules:**
- New features MUST use atomic components
- Refactor old features during major updates
- Deprecate old components gradually
- Celebrate wins

### Success Metrics

✅ Component reuse rate >70%
✅ New features ship faster
✅ Design consistency across pages
✅ Fewer UI bugs
✅ Easier onboarding

---

## Conclusion

### Key Principles

**1. Think in Systems, Not Pages**
- Components are building blocks
- Consistency through reuse
- Changes cascade predictably

**2. Respect the Hierarchy**
- Dependencies flow one direction
- Each level has clear responsibilities
- No circular dependencies

**3. Start Simple**
- Build what you need
- Extract patterns when they emerge
- Don't over-engineer early

**4. Use Design Tokens**
- Never hard-code colors/spacing
- Tokens enable theming
- Single source of truth

**5. Schema-Driven When Appropriate**
- Great for forms and data entry
- Enables non-developer customization
- Not for everything

### When Atomic Design Succeeds

✅ Enterprise applications (ERP, CRM)
✅ Design systems
✅ Multi-tenant applications
✅ Long-term maintenance
✅ Large teams

### When to Be Pragmatic

⚠️ Don't enforce strictness over productivity
⚠️ Skip levels when it makes sense
⚠️ Duplication is okay initially
⚠️ Refactor when patterns emerge

### Final Advice

**For Students:**
- Practice breaking down UIs you see daily
- Understand the "why" behind each level
- Build a personal component library

**For Developers:**
- Start with design tokens and atoms
- Let the system grow organically
- Resist premature abstraction
- Document as you go

**For Teams:**
- Get designer-developer alignment
- Establish naming conventions
- Review components together
- Celebrate reuse wins

---

**Version:** 2.0  
**Last Updated:** November 2025  
**Authors:** For Awo ERP Development  
**Stack:** Go + Templ + HTMX + PostgreSQL
