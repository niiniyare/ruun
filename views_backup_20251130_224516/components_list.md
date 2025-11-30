# UI Component Reference: Complete Atomic Design System

**Version**: 2.0  
**Last Updated**: November 2025  
**System**: Awo ERP - Go + Templ + HTMX + Alpine.js + PostgreSQL  
**Token System**: `pkg/schema` Three-Tier Architecture

---

## Table of Contents

1. [System Overview](#system-overview)
2. [Dependency Graph](#dependency-graph)
3. [Token Integration](#token-integration)
4. [Level 1: Atoms (Complete)](#level-1-atoms)
5. [Level 2: Molecules (Complete)](#level-2-molecules)
6. [Level 3: Organisms (Complete)](#level-3-organisms)
7. [Level 4: Templates (Complete)](#level-4-templates)
8. [Level 5: Pages (Complete)](#level-5-pages)
9. [Component Registry](#component-registry)
10. [Usage Patterns](#usage-patterns)
11. [Accessibility Guidelines](#accessibility-guidelines)
12. [Responsive Design Patterns](#responsive-design-patterns)

---

## System Overview

### Architecture Principles

```
Design Tokens (pkg/schema)
    ↓
Primitives → Semantic → Components
    ↓
Atoms → Molecules → Organisms → Templates → Pages
    ↓
HTMX (behavior) + Alpine.js (state)
```

### Key Rules

1. **Unidirectional Dependencies**: Lower levels never import higher levels
2. **Token-Based Styling**: All visual properties use schema token references
3. **Single Responsibility**: Each component has one clear purpose
4. **Composition Over Configuration**: Build complexity through composition
5. **Schema-Driven**: Support runtime schema interpretation
6. **Progressive**: Work without JavaScript, enhanced with it
7. **Accessibility First**: WCAG 2.1 AA compliance minimum
8. **Mobile First**: Responsive from smallest screens up

### Technology Integration

```go
// Component with HTMX + Alpine.js
templ Button(props ButtonProps) {
    <button
        type={ props.Type }
        class={ getButtonClasses(props) }
        disabled?={ props.Disabled }
        // HTMX attributes for server interaction
        if props.HxPost != "" {
            hx-post={ props.HxPost }
            hx-target={ props.HxTarget }
            hx-swap={ props.HxSwap }
        }
        // Alpine.js for client-side state
        if props.Loading {
            x-data="{ loading: true }"
            x-bind:disabled="loading"
        }
    >
        if props.Loading {
            @Spinner(SpinnerProps{Size: SpinnerSM})
        }
        if props.IconLeft != "" {
            @Icon(IconProps{Name: props.IconLeft, Size: IconSM})
        }
        { props.Label }
        if props.IconRight != "" {
            @Icon(IconProps{Name: props.IconRight, Size: IconSM})
        }
    </button>
}
```

---

## Dependency Graph

### Visual Hierarchy 

```
┌──────────────────────────────────────────────────────────────┐
│ LEVEL 5: PAGES                                               │
│ Dependencies: Templates, Organisms, Molecules, Atoms         │
│ State: Application state, data fetching, routing             │
│ Examples: InvoiceDetailPage, DashboardPage, LoginPage        │
│ Reusability: Single-use (1 per route)                        │
└────────────────────┬─────────────────────────────────────────┘
                     │
┌────────────────────┴─────────────────────────────────────────┐
│ LEVEL 4: TEMPLATES                                           │
│ Dependencies: Organisms, Molecules, Atoms                    │
│ State: Layout state (collapsed, active tab, viewport)        │
│ Examples: MasterDetailLayout, DashboardLayout, FormLayout    │
│ Reusability: Medium-High (3-10 usages)                       │
└────────────────────┬─────────────────────────────────────────┘
                     │
┌────────────────────┴─────────────────────────────────────────┐
│ LEVEL 3: ORGANISMS                                           │
│ Dependencies: Molecules, Atoms, (other Organisms carefully)  │
│ State: Business logic state, complex UI state                │
│ Examples: DataTable, NavigationHeader, MultiSectionForm      │
│ Reusability: Medium (10-30 usages)                           │
└────────────────────┬─────────────────────────────────────────┘
                     │
┌────────────────────┴─────────────────────────────────────────┐
│ LEVEL 2: MOLECULES                                           │
│ Dependencies: Atoms only (2-5 atoms combined)                │
│ State: Local UI state (show/hide, input value, hover)        │
│ Examples: FormField, SearchBar, Pagination                   │
│ Reusability: High (30-100 usages)                            │
└────────────────────┬─────────────────────────────────────────┘
                     │
┌────────────────────┴─────────────────────────────────────────┐
│ LEVEL 1: ATOMS                                               │
│ Dependencies: Design Tokens ONLY                             │
│ State: Presentational only, no business logic                │
│ Examples: Button, Input, Icon, Badge, Checkbox               │
│ Reusability: Universal (100+ usages)                         │
└────────────────────┬─────────────────────────────────────────┘
                     │
┌────────────────────┴─────────────────────────────────────────┐
│ LEVEL 0: DESIGN TOKENS (pkg/schema)                          │
│ Primitives → Semantic → Components                           │
│ Examples: colors, spacing, typography, breakpoints           │
└──────────────────────────────────────────────────────────────┘
```

---

## Token Integration

### Token Reference Pattern

```go
// Token resolution with fallbacks and validation
type ComponentTokens struct {
    // Primary tokens
    Background   schema.TokenReference `token:"components.button.{variant}.background" fallback:"semantic.colors.primary"`
    Color        schema.TokenReference `token:"components.button.{variant}.color" fallback:"semantic.colors.text.on-primary"`
    Border       schema.TokenReference `token:"components.button.{variant}.border" fallback:"transparent"`
    
    // State tokens
    HoverBg      schema.TokenReference `token:"components.button.{variant}.hover.background"`
    ActiveBg     schema.TokenReference `token:"components.button.{variant}.active.background"`
    DisabledBg   schema.TokenReference `token:"components.button.{variant}.disabled.background"`
    
    // Spacing tokens
    PaddingX     schema.TokenReference `token:"primitives.spacing.{size}.x" default:"primitives.spacing.4"`
    PaddingY     schema.TokenReference `token:"primitives.spacing.{size}.y" default:"primitives.spacing.2"`
    
    // Typography tokens
    FontSize     schema.TokenReference `token:"primitives.typography.{size}.size"`
    FontWeight   schema.TokenReference `token:"primitives.typography.{size}.weight"`
    LineHeight   schema.TokenReference `token:"primitives.typography.{size}.line-height"`
}

// Runtime resolution with caching
func (b *Button) ResolveTokens(ctx context.Context) (*ResolvedTokens, error) {
    registry := schema.GetTokenRegistryFromContext(ctx)
    
    // Check cache first
    cacheKey := fmt.Sprintf("button:%s:%s", b.Variant, b.Size)
    if cached, ok := registry.GetCached(cacheKey); ok {
        return cached, nil
    }
    
    // Resolve with variable substitution
    vars := map[string]string{
        "variant": string(b.Variant),
        "size":    string(b.Size),
    }
    
    resolved, err := registry.ResolveWithVars(ctx, &ComponentTokens{}, vars)
    if err != nil {
        return nil, err
    }
    
    // Cache for future use
    registry.SetCached(cacheKey, resolved)
    
    return resolved, nil
}
```

---

## Level 1: Atoms

### Complete Atom Catalog (24 Essential Components)

#### 1. Button 

**Type Signature**:
```go
type ButtonVariant string
const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    ButtonTertiary  ButtonVariant = "tertiary"
    ButtonDanger    ButtonVariant = "danger"
    ButtonSuccess   ButtonVariant = "success"
    ButtonWarning   ButtonVariant = "warning"
    ButtonGhost     ButtonVariant = "ghost"
    ButtonLink      ButtonVariant = "link"
)

type ButtonSize string
const (
    ButtonXS ButtonSize = "xs" // 24px height
    ButtonSM ButtonSize = "sm" // 32px height
    ButtonMD ButtonSize = "md" // 40px height
    ButtonLG ButtonSize = "lg" // 48px height
    ButtonXL ButtonSize = "xl" // 56px height
)

type ButtonProps struct {
    Label         string        `json:"label"`
    Variant       ButtonVariant `json:"variant"`
    Size          ButtonSize    `json:"size"`
    Disabled      bool          `json:"disabled"`
    Loading       bool          `json:"loading"`
    FullWidth     bool          `json:"fullWidth"`
    IconLeft      string        `json:"iconLeft,omitempty"`
    IconRight     string        `json:"iconRight,omitempty"`
    IconOnly      bool          `json:"iconOnly"`
    Type          string        `json:"type"` // "button", "submit", "reset"
    
    // HTMX attributes
    HxPost        string        `json:"hxPost,omitempty"`
    HxGet         string        `json:"hxGet,omitempty"`
    HxPut         string        `json:"hxPut,omitempty"`
    HxDelete      string        `json:"hxDelete,omitempty"`
    HxTarget      string        `json:"hxTarget,omitempty"`
    HxSwap        string        `json:"hxSwap,omitempty"`
    HxConfirm     string        `json:"hxConfirm,omitempty"`
    
    // Alpine.js integration
    XData         string        `json:"xData,omitempty"`
    XOn           map[string]string `json:"xOn,omitempty"`
    
    // Accessibility
    AriaLabel     string        `json:"ariaLabel,omitempty"`
    AriaDescribedBy string      `json:"ariaDescribedBy,omitempty"`
}
```

**Token References**:
```
components.button.{variant}.background
components.button.{variant}.color
components.button.{variant}.border
components.button.{variant}.hover.background
components.button.{variant}.hover.border
components.button.{variant}.active.background
components.button.{variant}.focus.ring
components.button.{variant}.disabled.background
components.button.{variant}.disabled.color
primitives.spacing.{size}.padding-x
primitives.spacing.{size}.padding-y
primitives.borders.radius.md
primitives.typography.{size}.size
primitives.typography.{size}.weight
primitives.transitions.base
primitives.shadows.{variant}.shadow
```

**ARIA Attributes**:
- `role="button"` (when not using `<button>`)
- `aria-busy="true"` (when loading)
- `aria-disabled="true"` (when disabled)
- `aria-label` (for icon-only buttons)

---

#### 2. Input 

**Type Signature**:
```go
type InputType string
const (
    InputText     InputType = "text"
    InputEmail    InputType = "email"
    InputPassword InputType = "password"
    InputNumber   InputType = "number"
    InputTel      InputType = "tel"
    InputURL      InputType = "url"
    InputSearch   InputType = "search"
    InputDate     InputType = "date"
    InputTime     InputType = "time"
    InputDateTime InputType = "datetime-local"
    InputFile     InputType = "file"
    InputHidden   InputType = "hidden"
)

type InputSize string
const (
    InputSM InputSize = "sm"
    InputMD InputSize = "md"
    InputLG InputSize = "lg"
)

type InputProps struct {
    Name          string      `json:"name"`
    Type          InputType   `json:"type"`
    Value         string      `json:"value"`
    Placeholder   string      `json:"placeholder,omitempty"`
    Size          InputSize   `json:"size"`
    Disabled      bool        `json:"disabled"`
    ReadOnly      bool        `json:"readonly"`
    Required      bool        `json:"required"`
    Error         bool        `json:"error"`
    Success       bool        `json:"success"`
    AutoFocus     bool        `json:"autofocus"`
    AutoComplete  string      `json:"autocomplete,omitempty"`
    MaxLength     int         `json:"maxlength,omitempty"`
    MinLength     int         `json:"minlength,omitempty"`
    Pattern       string      `json:"pattern,omitempty"`
    
    // Number-specific
    Min           *float64    `json:"min,omitempty"`
    Max           *float64    `json:"max,omitempty"`
    Step          *float64    `json:"step,omitempty"`
    
    // HTMX
    HxPost        string      `json:"hxPost,omitempty"`
    HxTrigger     string      `json:"hxTrigger,omitempty"` // "keyup changed delay:500ms"
    HxTarget      string      `json:"hxTarget,omitempty"`
    HxSwap        string      `json:"hxSwap,omitempty"`
    
    // Alpine.js
    XModel        string      `json:"xModel,omitempty"`
    XOn           map[string]string `json:"xOn,omitempty"`
    
    // Accessibility
    AriaLabel     string      `json:"ariaLabel,omitempty"`
    AriaDescribedBy string    `json:"ariaDescribedBy,omitempty"`
    AriaInvalid   bool        `json:"ariaInvalid"`
}
```

**Token References**:
```
components.input.default.background
components.input.default.border
components.input.default.color
components.input.disabled.background
components.input.disabled.color
components.input.error.border
components.input.error.background
components.input.success.border
components.input.focus.border
components.input.focus.ring
primitives.spacing.{size}.padding-x
primitives.spacing.{size}.padding-y
primitives.borders.radius.md
primitives.typography.{size}.size
primitives.transitions.base
```

---

#### 3. Select 

**Type Signature**:
```go
type SelectOption struct {
    Value    string `json:"value"`
    Label    string `json:"label"`
    Disabled bool   `json:"disabled"`
    Group    string `json:"group,omitempty"`
}

type SelectProps struct {
    Name         string         `json:"name"`
    Value        string         `json:"value"`
    Options      []SelectOption `json:"options"`
    Placeholder  string         `json:"placeholder,omitempty"`
    Size         InputSize      `json:"size"`
    Disabled     bool           `json:"disabled"`
    Required     bool           `json:"required"`
    Multiple     bool           `json:"multiple"`
    Error        bool           `json:"error"`
    
    // HTMX
    HxPost       string         `json:"hxPost,omitempty"`
    HxTrigger    string         `json:"hxTrigger,omitempty"`
    HxTarget     string         `json:"hxTarget,omitempty"`
    
    // Alpine.js for custom dropdown
    CustomDropdown bool         `json:"customDropdown"`
    Searchable     bool         `json:"searchable"`
    
    // Accessibility
    AriaLabel    string         `json:"ariaLabel,omitempty"`
}
```

**Token References**:
```
components.select.background
components.select.border
components.select.color
components.select.focus.border
components.select.focus.ring
components.select.option.hover.background
components.select.icon.color
primitives.spacing.{size}.padding
primitives.borders.radius.md
primitives.transitions.base
```

**Custom Dropdown with Alpine.js**:
```html
<div x-data="{ open: false, selected: '', search: '' }" 
     @click.away="open = false"
     class="relative">
    <button @click="open = !open" 
            type="button"
            class="select-trigger">
        <span x-text="selected || 'Select option'"></span>
        @Icon(IconProps{Name: "chevron-down"})
    </button>
    
    <div x-show="open" 
         x-transition
         class="select-dropdown">
        <input x-model="search" 
               placeholder="Search..."
               class="select-search">
        
        <template x-for="option in filteredOptions" :key="option.value">
            <button @click="selected = option.label; open = false"
                    class="select-option"
                    x-text="option.label"></button>
        </template>
    </div>
</div>
```

---

#### 4. Textarea 

**Type Signature**:
```go
type TextareaProps struct {
    Name         string    `json:"name"`
    Value        string    `json:"value"`
    Placeholder  string    `json:"placeholder,omitempty"`
    Rows         int       `json:"rows"` // Default: 3
    MaxRows      int       `json:"maxRows,omitempty"` // Auto-expand limit
    Disabled     bool      `json:"disabled"`
    ReadOnly     bool      `json:"readonly"`
    Required     bool      `json:"required"`
    Error        bool      `json:"error"`
    MaxLength    int       `json:"maxlength,omitempty"`
    AutoResize   bool      `json:"autoResize"` // Auto-expand on type
    
    // HTMX
    HxPost       string    `json:"hxPost,omitempty"`
    HxTrigger    string    `json:"hxTrigger,omitempty"`
    
    // Alpine.js
    XModel       string    `json:"xModel,omitempty"`
    
    // Accessibility
    AriaLabel    string    `json:"ariaLabel,omitempty"`
}
```

**Auto-resize Implementation**:
```html
<textarea 
    x-data="{ resize() { $el.style.height = 'auto'; $el.style.height = $el.scrollHeight + 'px'; } }"
    x-init="resize()"
    @input="resize()"
    :style="{ minHeight: '80px', maxHeight: '400px' }"
>
</textarea>
```

---

#### 5. Toggle/Switch 

**Type Signature**:
```go
type ToggleSize string
const (
    ToggleSM ToggleSize = "sm" // 36px width
    ToggleMD ToggleSize = "md" // 44px width
    ToggleLG ToggleLG ToggleSize = "lg" // 52px width
)

type ToggleProps struct {
    Name         string     `json:"name"`
    Checked      bool       `json:"checked"`
    Disabled     bool       `json:"disabled"`
    Size         ToggleSize `json:"size"`
    Label        string     `json:"label,omitempty"`
    LabelLeft    bool       `json:"labelLeft"` // Label on left side
    
    // HTMX
    HxPost       string     `json:"hxPost,omitempty"`
    HxTrigger    string     `json:"hxTrigger,omitempty"`
    
    // Alpine.js
    XModel       string     `json:"xModel,omitempty"`
    
    // Accessibility
    AriaLabel    string     `json:"ariaLabel,omitempty"`
}
```

**Token References**:
```
components.toggle.background.off
components.toggle.background.on
components.toggle.handle.background
components.toggle.border
components.toggle.focus.ring
primitives.transitions.base
```

**Alpine.js Implementation**:
```html
<button 
    type="button"
    role="switch"
    x-data="{ on: false }"
    @click="on = !on"
    :aria-checked="on.toString()"
    :class="{ 'bg-primary': on, 'bg-gray-200': !on }"
    class="toggle-switch"
>
    <span 
        :class="{ 'translate-x-5': on, 'translate-x-0': !on }"
        class="toggle-handle"
    ></span>
</button>
```

---

#### 6. Link 

**Type Signature**:
```go
type LinkVariant string
const (
    LinkDefault   LinkVariant = "default"
    LinkPrimary   LinkVariant = "primary"
    LinkSecondary LinkVariant = "secondary"
    LinkMuted     LinkVariant = "muted"
)

type LinkProps struct {
    Href         string      `json:"href"`
    Label        string      `json:"label,omitempty"` // If empty, uses children
    Variant      LinkVariant `json:"variant"`
    External     bool        `json:"external"` // Opens in new tab
    Disabled     bool        `json:"disabled"`
    Underline    bool        `json:"underline"` // Always underlined
    IconLeft     string      `json:"iconLeft,omitempty"`
    IconRight    string      `json:"iconRight,omitempty"`
    
    // Accessibility
    AriaLabel    string      `json:"ariaLabel,omitempty"`
    AriaCurrent  string      `json:"ariaCurrent,omitempty"` // "page", "step", etc.
}
```

**Token References**:
```
components.link.{variant}.color
components.link.{variant}.hover.color
components.link.{variant}.visited.color
components.link.{variant}.active.color
components.link.underline.thickness
primitives.transitions.base
```

---

#### 7. Tooltip 

**Type Signature**:
```go
type TooltipPlacement string
const (
    TooltipTop    TooltipPlacement = "top"
    TooltipBottom TooltipPlacement = "bottom"
    TooltipLeft   TooltipPlacement = "left"
    TooltipRight  TooltipPlacement = "right"
)

type TooltipProps struct {
    Content      string           `json:"content"`
    Placement    TooltipPlacement `json:"placement"`
    Delay        int              `json:"delay"` // ms before showing
    MaxWidth     string           `json:"maxWidth,omitempty"`
}
```

**Token References**:
```
components.tooltip.background
components.tooltip.color
components.tooltip.shadow
primitives.borders.radius.md
primitives.spacing.2.padding
primitives.typography.sm.size
primitives.transitions.fast
```

**Alpine.js Implementation**:
```html
<div x-data="{ show: false, timeout: null }"
     @mouseenter="timeout = setTimeout(() => show = true, 300)"
     @mouseleave="clearTimeout(timeout); show = false"
     class="relative inline-block">
    
    <!-- Trigger -->
    <slot></slot>
    
    <!-- Tooltip -->
    <div x-show="show"
         x-transition
         :class="placement"
         class="tooltip"
         role="tooltip">
        { content }
        <div class="tooltip-arrow"></div>
    </div>
</div>
```

---

#### 8. ProgressBar 

**Type Signature**:
```go
type ProgressVariant string
const (
    ProgressDefault ProgressVariant = "default"
    ProgressSuccess ProgressVariant = "success"
    ProgressWarning ProgressVariant = "warning"
    ProgressDanger  ProgressVariant = "danger"
)

type ProgressSize string
const (
    ProgressSM ProgressSize = "sm" // 4px height
    ProgressMD ProgressSize = "md" // 8px height
    ProgressLG ProgressSize = "lg" // 12px height
)

type ProgressProps struct {
    Value        float64          `json:"value"` // 0-100
    Max          float64          `json:"max"` // Default: 100
    Variant      ProgressVariant  `json:"variant"`
    Size         ProgressSize     `json:"size"`
    Striped      bool             `json:"striped"`
    Animated     bool             `json:"animated"`
    ShowLabel    bool             `json:"showLabel"`
    Label        string           `json:"label,omitempty"` // Custom label
    Indeterminate bool            `json:"indeterminate"` // Loading spinner style
    
    // Accessibility
    AriaLabel    string           `json:"ariaLabel,omitempty"`
}
```

**Token References**:
```
components.progress.background
components.progress.{variant}.fill
primitives.borders.radius.full
primitives.transitions.base
```

---

#### 9. Skeleton 

**Type Signature**:
```go
type SkeletonVariant string
const (
    SkeletonText      SkeletonVariant = "text"
    SkeletonCircle    SkeletonVariant = "circle"
    SkeletonRectangle SkeletonVariant = "rectangle"
)

type SkeletonProps struct {
    Variant      SkeletonVariant `json:"variant"`
    Width        string          `json:"width,omitempty"` // CSS value
    Height       string          `json:"height,omitempty"`
    Animated     bool            `json:"animated"` // Pulse or shimmer
    Count        int             `json:"count"` // Multiple lines for text
}
```

**Token References**:
```
components.skeleton.background
components.skeleton.shimmer
primitives.borders.radius.md
```

**CSS Animation**:
```css
@keyframes skeleton-loading {
    0% { background-position: -200px 0; }
    100% { background-position: calc(200px + 100%) 0; }
}

.skeleton {
    background: linear-gradient(
        90deg,
        var(--skeleton-background) 0px,
        var(--skeleton-shimmer) 40px,
        var(--skeleton-background) 80px
    );
    background-size: 200px 100%;
    animation: skeleton-loading 1.2s ease-in-out infinite;
}
```

---

#### 10. Alert 

**Type Signature**:
```go
type AlertVariant string
const (
    AlertInfo    AlertVariant = "info"
    AlertSuccess AlertVariant = "success"
    AlertWarning AlertVariant = "warning"
    AlertDanger  AlertVariant = "danger"
)

type AlertProps struct {
    Variant      AlertVariant `json:"variant"`
    Title        string       `json:"title,omitempty"`
    Message      string       `json:"message"`
    Dismissible  bool         `json:"dismissible"`
    Icon         string       `json:"icon,omitempty"` // Auto-selected if empty
    
    // HTMX for server-side dismissal
    HxDelete     string       `json:"hxDelete,omitempty"`
    
    // Accessibility
    Role         string       `json:"role"` // "alert", "status"
    AriaLive     string       `json:"ariaLive"` // "polite", "assertive"
}
```

**Token References**:
```
components.alert.{variant}.background
components.alert.{variant}.border
components.alert.{variant}.color
components.alert.{variant}.icon
primitives.spacing.4.padding
primitives.borders.radius.md
```

**Alpine.js Dismissal**:
```html
<div x-data="{ show: true }"
     x-show="show"
     x-transition
     role="alert"
     class="alert alert-{{ variant }}">
    
    @Icon(IconProps{Name: icon})
    
    <div class="alert-content">
        if title != "" {
            <div class="alert-title">{ title }</div>
        }
        <div class="alert-message">{ message }</div>
    </div>
    
    if dismissible {
        <button @click="show = false" 
                class="alert-close"
                aria-label="Dismiss">
            @Icon(IconProps{Name: "x"})
        </button>
    }
</div>
```

---

### Complete Atoms Summary

| # | Atom | Purpose | Token Category | Reusability | ARIA Role |
|---|------|---------|---------------|-------------|-----------|
| 1 | Button | Action trigger | Component | Universal | button |
| 2 | Input | Text entry | Component | Universal | textbox |
| 3 | Label | Field identifier | Semantic | Universal | - |
| 4 | Icon | Visual indicator | Primitive | Universal | img |
| 5 | Badge | Status/category | Component | Universal | status |
| 6 | Spinner | Loading state | Semantic | Universal | status |
| 7 | Divider | Visual separator | Semantic | Universal | separator |
| 8 | Avatar | User representation | Component | Very High | img |
| 9 | Checkbox | Boolean selection | Component | Very High | checkbox |
| 10 | Radio | Single selection | Component | Very High | radio |
| 11 | Select | Dropdown selection | Component | Very High | listbox |
| 12 | Textarea | Multi-line text | Component | Very High | textbox |
| 13 | Toggle | Binary state | Component | Very High | switch |
| 14 | Link | Navigation | Semantic | Universal | link |
| 15 | Image | Media display | Primitive | Universal | img |
| 16 | Tooltip | Help text | Component | Very High | tooltip |
| 17 | ProgressBar | Progress indicator | Component | High | progressbar |
| 18 | Skeleton | Loading placeholder | Semantic | High | - |
| 19 | Alert | Inline message | Component | High | alert |
| 20 | Heading | Text hierarchy | Semantic | Universal | heading |
| 21 | Text | Body copy | Semantic | Universal | - |
| 22 | Code | Code snippet | Semantic | High | code |
| 23 | Pre | Preformatted text | Semantic | High | - |
| 24 | Kbd | Keyboard shortcut | Semantic | Medium | - |

---

## Level 2: Molecules

### Complete Molecule Catalog (20 Essential Components)

#### 11. StatusBadge 

**Type Signature**:
```go
type StatusType string
const (
    StatusActive    StatusType = "active"
    StatusInactive  StatusType = "inactive"
    StatusPending   StatusType = "pending"
    StatusApproved  StatusType = "approved"
    StatusRejected  StatusType = "rejected"
    StatusDraft     StatusType = "draft"
    StatusPublished StatusType = "published"
    StatusArchived  StatusType = "archived"
)

type StatusBadgeProps struct {
    Status       StatusType `json:"status"`
    Label        string     `json:"label,omitempty"` // Override default label
    WithIcon     bool       `json:"withIcon"`
    WithDot      bool       `json:"withDot"`
    Pulsing      bool       `json:"pulsing"` // Animated pulse effect
}
```

**Composition**:
```
StatusBadge
├── Badge (atom)
├── Icon (atom) - if WithIcon
└── Dot (span) - if WithDot
```

**Token References**:
```
components.status.{status}.background
components.status.{status}.color
components.status.{status}.border
components.status.{status}.icon
```

**Auto Icon Mapping**:
```go
var statusIcons = map[StatusType]string{
    StatusActive:    "check-circle",
    StatusInactive:  "x-circle",
    StatusPending:   "clock",
    StatusApproved:  "check",
    StatusRejected:  "x",
    StatusDraft:     "edit",
    StatusPublished: "eye",
    StatusArchived:  "archive",
}
```

---

#### 12. Pagination 

**Type Signature**:
```go
type PaginationProps struct {
    CurrentPage  int    `json:"currentPage"`
    TotalPages   int    `json:"totalPages"`
    TotalItems   int    `json:"totalItems"`
    PageSize     int    `json:"pageSize"`
    PageSizes    []int  `json:"pageSizes,omitempty"` // [10, 25, 50, 100]
    ShowSizeSelector bool `json:"showSizeSelector"`
    ShowFirstLast    bool `json:"showFirstLast"`
    MaxPages     int    `json:"maxPages"` // Max page buttons to show (default: 7)
    
    // HTMX
    HxGet        string `json:"hxGet"` // Base URL
    HxTarget     string `json:"hxTarget"`
    HxSwap       string `json:"hxSwap"`
}
```

**Composition**:
```
Pagination
├── PageInfo ("Showing 1-10 of 100")
├── PageSizeSelector (if enabled)
│   └── Select (atom)
└── PageButtons
    ├── FirstButton (if enabled)
    ├── PrevButton
    ├── PageNumbers (with ellipsis)
    ├── NextButton
    └── LastButton (if enabled)
```

**Token References**:
```
components.pagination.button.background
components.pagination.button.active.background
components.pagination.button.hover.background
primitives.spacing.2.gap
```

**HTMX Implementation**:
```html
<nav class="pagination" aria-label="Pagination">
    <div class="pagination-info">
        Showing {{ (currentPage-1)*pageSize + 1 }}-{{ min(currentPage*pageSize, totalItems) }} of {{ totalItems }}
    </div>
    
    <div class="pagination-controls">
        @Button(ButtonProps{
            Label: "Previous",
            Disabled: currentPage == 1,
            HxGet: fmt.Sprintf("%s?page=%d", baseURL, currentPage-1),
            HxTarget: "#table-container",
            HxSwap: "outerHTML",
        })
        
        <!-- Page numbers -->
        for _, page := range getPageRange(currentPage, totalPages, maxPages) {
            if page == "..." {
                <span class="pagination-ellipsis">...</span>
            } else {
                @Button(ButtonProps{
                    Label: fmt.Sprintf("%d", page),
                    Variant: page == currentPage ? ButtonPrimary : ButtonGhost,
                    HxGet: fmt.Sprintf("%s?page=%d", baseURL, page),
                    HxTarget: "#table-container",
                })
            }
        }
        
        @Button(ButtonProps{
            Label: "Next",
            Disabled: currentPage == totalPages,
            HxGet: fmt.Sprintf("%s?page=%d", baseURL, currentPage+1),
        })
    </div>
</nav>
```

---

#### 13. EmptyState 

**Type Signature**:
```go
type EmptyStateProps struct {
    Icon         string  `json:"icon"`
    Title        string  `json:"title"`
    Description  string  `json:"description"`
    ActionLabel  string  `json:"actionLabel,omitempty"`
    ActionURL    string  `json:"actionUrl,omitempty"`
    ActionClick  string  `json:"actionClick,omitempty"` // HTMX or Alpine.js
    SecondaryActionLabel string `json:"secondaryActionLabel,omitempty"`
    SecondaryActionURL   string `json:"secondaryActionUrl,omitempty"`
    Compact      bool    `json:"compact"` // Smaller version
}
```

**Composition**:
```
EmptyState
├── Icon (atom) - large size
├── Title (heading)
├── Description (text)
└── Actions
    ├── PrimaryAction (button)
    └── SecondaryAction (button/link)
```

**Token References**:
```
semantic.colors.text.muted
semantic.colors.icon.muted
primitives.spacing.8.gap
primitives.typography.lg.size (title)
```

---

#### 14. FilterChip 

**Type Signature**:
```go
type FilterChipProps struct {
    Label        string `json:"label"`
    Value        string `json:"value"`
    Removable    bool   `json:"removable"`
    Active       bool   `json:"active"`
    
    // HTMX for removal
    HxDelete     string `json:"hxDelete,omitempty"`
    HxTarget     string `json:"hxTarget,omitempty"`
}
```

**Composition**:
```
FilterChip
├── Badge (atom) with interactive styling
└── RemoveButton (icon button)
```

**Alpine.js Implementation**:
```html
<div x-data="{ active: {{ active }} }"
     :class="{ 'filter-chip-active': active }"
     class="filter-chip">
    
    <span>{{ label }}</span>
    
    if removable {
        <button @click="active = false; $dispatch('filter-removed', { value: '{{ value }}' })"
                class="filter-chip-remove"
                aria-label="Remove filter">
            @Icon(IconProps{Name: "x", Size: IconXS})
        </button>
    }
</div>
```

---

#### 15. ActionButtons 

**Type Signature**:
```go
type ActionButton struct {
    Label       string        `json:"label"`
    Icon        string        `json:"icon,omitempty"`
    Variant     ButtonVariant `json:"variant"`
    Disabled    bool          `json:"disabled"`
    HxPost      string        `json:"hxPost,omitempty"`
    HxDelete    string        `json:"hxDelete,omitempty"`
    HxConfirm   string        `json:"hxConfirm,omitempty"`
    Permission  string        `json:"permission,omitempty"` // Required permission
}

type ActionButtonsProps struct {
    Actions      []ActionButton `json:"actions"`
    Layout       string         `json:"layout"` // "horizontal", "vertical", "dropdown"
    Align        string         `json:"align"` // "left", "center", "right"
    Spacing      string         `json:"spacing"` // Token reference
}
```

**Composition**:
```
ActionButtons
└── For each action:
    └── Button (atom) with permission check
```

**With Dropdown (for mobile)**:
```html
<div x-data="{ open: false }" class="action-buttons">
    <!-- Desktop: Horizontal buttons -->
    <div class="hidden md:flex md:gap-2">
        for _, action := range actions {
            @Button(ButtonProps{
                Label: action.Label,
                Variant: action.Variant,
                HxPost: action.HxPost,
            })
        }
    </div>
    
    <!-- Mobile: Dropdown menu -->
    <div class="md:hidden relative">
        <button @click="open = !open" class="btn-icon">
            @Icon(IconProps{Name: "more-vertical"})
        </button>
        
        <div x-show="open" 
             @click.away="open = false"
             class="dropdown-menu">
            for _, action := range actions {
                <button class="dropdown-item">
                    @Icon(IconProps{Name: action.Icon})
                    {{ action.Label }}
                </button>
            }
        </div>
    </div>
</div>
```

---

#### 16. CardHeader 

**Type Signature**:
```go
type CardHeaderProps struct {
    Title        string `json:"title"`
    Subtitle     string `json:"subtitle,omitempty"`
    Icon         string `json:"icon,omitempty"`
    Actions      []ActionButton `json:"actions,omitempty"`
    Badge        *BadgeProps `json:"badge,omitempty"`
}
```

**Composition**:
```
CardHeader
├── Left Section
│   ├── Icon (atom)
│   ├── Titles
│   │   ├── Title
│   │   └── Subtitle
│   └── Badge (atom)
└── Right Section
    └── ActionButtons (molecule)
```

---

#### 17. ListItem 

**Type Signature**:
```go
type ListItemProps struct {
    Title        string      `json:"title"`
    Description  string      `json:"description,omitempty"`
    ImageURL     string      `json:"imageUrl,omitempty"`
    Icon         string      `json:"icon,omitempty"`
    Badge        *BadgeProps `json:"badge,omitempty"`
    Metadata     []string    `json:"metadata,omitempty"` // ["2 hours ago", "3 comments"]
    Actions      []ActionButton `json:"actions,omitempty"`
    Clickable    bool        `json:"clickable"`
    Selected     bool        `json:"selected"`
    
    // HTMX
    HxGet        string      `json:"hxGet,omitempty"`
    HxTarget     string      `json:"hxTarget,omitempty"`
}
```

**Composition**:
```
ListItem
├── SelectCheckbox (if selectable)
├── MediaObject (molecule)
│   ├── Avatar/Icon/Image
│   └── Content
│       ├── Title
│       ├── Description
│       ├── Metadata (small text)
│       └── Badge
└── Actions (molecule)
```

---

#### 18. Tabs 

**Type Signature**:
```go
type Tab struct {
    ID       string `json:"id"`
    Label    string `json:"label"`
    Icon     string `json:"icon,omitempty"`
    Badge    string `json:"badge,omitempty"` // Count or status
    Disabled bool   `json:"disabled"`
}

type TabsProps struct {
    Tabs         []Tab  `json:"tabs"`
    ActiveTab    string `json:"activeTab"`
    Variant      string `json:"variant"` // "underline", "pills", "enclosed"
    FullWidth    bool   `json:"fullWidth"`
    
    // HTMX for lazy loading tab content
    HxGet        string `json:"hxGet,omitempty"` // Base URL for tab content
    HxTarget     string `json:"hxTarget,omitempty"`
}
```

**Composition**:
```
Tabs
├── TabList (navigation)
│   └── For each tab:
│       ├── TabButton
│       │   ├── Icon (if exists)
│       │   ├── Label
│       │   └── Badge (if exists)
│       └── ActiveIndicator (underline/background)
└── TabPanel (content slot)
```

**HTMX Lazy Loading**:
```html
<div class="tabs">
    <div role="tablist" class="tab-list">
        for _, tab := range tabs {
            <button role="tab"
                    :aria-selected="{{ tab.ID == activeTab }}"
                    hx-get="{{ baseURL }}/{{ tab.ID }}"
                    hx-target="#tab-content"
                    hx-swap="innerHTML"
                    class="tab-button">
                {{ tab.Label }}
            </button>
        }
    </div>
    
    <div id="tab-content" 
         role="tabpanel" 
         class="tab-panel">
        <!-- Content loaded via HTMX -->
    </div>
</div>
```

---

#### 19. DropdownMenu 

**Type Signature**:
```go
type MenuItem struct {
    ID          string     `json:"id"`
    Label       string     `json:"label"`
    Icon        string     `json:"icon,omitempty"`
    Variant     string     `json:"variant"` // "default", "danger"
    Divider     bool       `json:"divider"` // Show divider after this item
    Disabled    bool       `json:"disabled"`
    Children    []MenuItem `json:"children,omitempty"` // Submenu
    
    // Actions
    HxPost      string     `json:"hxPost,omitempty"`
    HxDelete    string     `json:"hxDelete,omitempty"`
    HxConfirm   string     `json:"hxConfirm,omitempty"`
    Href        string     `json:"href,omitempty"`
}

type DropdownMenuProps struct {
    Items        []MenuItem `json:"items"`
    Trigger      string     `json:"trigger"` // "click", "hover", "context"
    Placement    string     `json:"placement"` // "bottom-start", "bottom-end", "top-start", etc.
    TriggerLabel string     `json:"triggerLabel,omitempty"`
    TriggerIcon  string     `json:"triggerIcon,omitempty"`
}
```

**Composition**:
```
DropdownMenu
├── Trigger (button)
└── Menu (absolute positioned)
    └── For each item:
        ├── MenuItem
        │   ├── Icon
        │   ├── Label
        │   └── SubmenuIndicator (if children)
        ├── Divider (if divider)
        └── Submenu (if children)
```

**Alpine.js Implementation**:
```html
<div x-data="{ open: false, submenu: null }"
     @click.away="open = false"
     @keydown.escape="open = false"
     class="dropdown">
    
    <!-- Trigger -->
    <button @click="open = !open"
            :aria-expanded="open"
            aria-haspopup="true"
            class="dropdown-trigger">
        {{ triggerLabel }}
        @Icon(IconProps{Name: "chevron-down"})
    </button>
    
    <!-- Menu -->
    <div x-show="open"
         x-transition
         role="menu"
         class="dropdown-menu"
         style="display: none;">
        
        for i, item := range items {
            if item.Divider && i > 0 {
                @Divider(DividerProps{})
            }
            
            <button role="menuitem"
                    :class="{ 'dropdown-item-danger': '{{ item.Variant }}' === 'danger' }"
                    class="dropdown-item"
                    @click="open = false"
                    hx-post="{{ item.HxPost }}"
                    hx-confirm="{{ item.HxConfirm }}">
                
                if item.Icon != "" {
                    @Icon(IconProps{Name: item.Icon, Size: IconSM})
                }
                
                <span>{{ item.Label }}</span>
                
                if len(item.Children) > 0 {
                    @Icon(IconProps{Name: "chevron-right", Size: IconXS})
                }
            </button>
        }
    </div>
</div>
```

---

#### 20. InputGroup 

**Type Signature**:
```go
type InputGroupProps struct {
    Label        string      `json:"label,omitempty"`
    Prefix       string      `json:"prefix,omitempty"` // Text or icon
    Suffix       string      `json:"suffix,omitempty"` // Text or icon
    Input        InputProps  `json:"input"`
    HelpText     string      `json:"helpText,omitempty"`
    Error        string      `json:"error,omitempty"`
}
```

**Composition**:
```
InputGroup
├── Label (atom)
├── InputWrapper
│   ├── Prefix (text/icon)
│   ├── Input (atom)
│   └── Suffix (text/icon)
├── HelpText
└── ErrorMessage
```

**Examples**:
```html
<!-- Currency Input -->
<div class="input-group">
    <span class="input-prefix">KES</span>
    <input type="number" class="input">
    <span class="input-suffix">.00</span>
</div>

<!-- Search Input -->
<div class="input-group">
    <span class="input-prefix">
        @Icon(IconProps{Name: "search"})
    </span>
    <input type="search" class="input" placeholder="Search...">
    <button class="input-suffix-button">
        @Icon(IconProps{Name: "filter"})
    </button>
</div>
```

---

### Complete Molecules Summary

| # | Molecule | Atoms Used | Purpose | Reusability |
|---|----------|-----------|---------|-------------|
| 1 | FormField | Label, Input, Icon | Standard form input | Very High |
| 2 | SearchBar | Icon, Input, Button | Search functionality | High |
| 3 | StatDisplay | Icon, Text, Badge | Metric display | High |
| 4 | QuantitySelector | Button, Input, Icon | Numeric adjustment | High |
| 5 | MediaObject | Avatar/Icon, Text | Content with media | Very High |
| 6 | TagInput | Badge, Button, Icon | Tag management | Medium-High |
| 7 | DateRangePicker | Input, Icon, Label | Date range selection | High |
| 8 | CurrencyInput | Input, Select, Icon | Currency amounts | High |
| 9 | FileUploadButton | Button, Icon, Progress | File upload | High |
| 10 | BreadcrumbItem | Icon, Link | Navigation path | Medium |
| 11 | StatusBadge | Badge, Icon | Status indicator | High |
| 12 | Pagination | Button, Select, Text | Page navigation | High |
| 13 | EmptyState | Icon, Text, Button | No data display | High |
| 14 | FilterChip | Badge, Icon | Active filter | Medium |
| 15 | ActionButtons | Button | Grouped actions | Medium |
| 16 | CardHeader | Icon, Text, Badge, Button | Card title section | High |
| 17 | ListItem | Avatar, Text, Badge, Button | List row | Very High |
| 18 | Tabs | Button, Badge, Icon | Tab navigation | High |
| 19 | DropdownMenu | Button, Icon, Divider | Contextual menu | Very High |
| 20 | InputGroup | Input, Icon, Text | input | High |

---

## Level 3: Organisms

### Complete Organism Catalog (18 Essential Components)

I'll add the missing organisms to complete the catalog:

#### 11. UserMenu 

**Type Signature**:
```go
type UserMenuProps struct {
    User         User           `json:"user"`
    MenuItems    []MenuItem     `json:"menuItems"`
    ShowStatus   bool           `json:"showStatus"` // Online/offline indicator
    Notifications int           `json:"notifications"`
}
```

**Composition**:
```
UserMenu
├── Trigger
│   ├── Avatar (atom)
│   ├── UserInfo (name, role)
│   └── NotificationBadge (if count > 0)
└── DropdownMenu (molecule)
    ├── UserProfile Section
    │   ├── Avatar (large)
    │   ├── Name
    │   ├── Email
    │   └── StatusToggle (if enabled)
    ├── Divider
    ├── Navigation Items
    │   ├── Profile
    │   ├── Settings
    │   └── Preferences
    ├── Divider
    └── Logout Button
```

---

#### 12. NotificationList 

**Type Signature**:
```go
type Notification struct {
    ID          string    `json:"id"`
    Type        string    `json:"type"` // "info", "success", "warning", "danger"
    Title       string    `json:"title"`
    Message     string    `json:"message"`
    Icon        string    `json:"icon,omitempty"`
    Read        bool      `json:"read"`
    Timestamp   time.Time `json:"timestamp"`
    ActionLabel string    `json:"actionLabel,omitempty"`
    ActionURL   string    `json:"actionUrl,omitempty"`
}

type NotificationListProps struct {
    Notifications []Notification `json:"notifications"`
    ShowViewAll   bool           `json:"showViewAll"`
    MaxDisplay    int            `json:"maxDisplay"` // Show only first N
    EmptyMessage  string         `json:"emptyMessage"`
}
```

**Composition**:
```
NotificationList
├── Header
│   ├── Title ("Notifications")
│   └── MarkAllReadButton
├── NotificationItems
│   └── For each notification:
│       ├── ListItem (molecule)
│       │   ├── Icon (type-based)
│       │   ├── Content
│       │   │   ├── Title
│       │   │   ├── Message
│       │   │   └── Timestamp
│       │   └── Actions
│       │       ├── ActionButton (if exists)
│       │       └── MarkReadButton
│       └── ReadIndicator (dot)
├── Divider
└── Footer
    └── ViewAllLink
```

**HTMX Real-time Updates**:
```html
<div hx-get="/api/notifications"
     hx-trigger="every 30s"
     hx-swap="innerHTML"
     class="notification-list">
    <!-- Notifications render here -->
</div>
```

---

#### 13. Dropdown (Generic) 

**Type Signature**:
```go
type DropdownProps struct {
    Trigger      templ.Component `json:"-"` // Any component as trigger
    Content      templ.Component `json:"-"` // Any content
    Placement    string          `json:"placement"` // "bottom", "top", "left", "right"
    CloseOnClick bool            `json:"closeOnClick"`
    Width        string          `json:"width"` // Token or CSS value
}
```

**Alpine.js Implementation**:
```html
<div x-data="{ 
        open: false,
        toggle() { this.open = !this.open },
        close() { this.open = false }
     }"
     @keydown.escape.window="close()"
     @click.away="close()"
     class="dropdown">
    
    <!-- Trigger (slot) -->
    <div @click="toggle()">
        { trigger }
    </div>
    
    <!-- Content (slot) -->
    <div x-show="open"
         x-transition
         :class="placement"
         class="dropdown-content"
         style="display: none;">
        { content }
    </div>
</div>
```

---

#### 14. CalendarWidget 

**Type Signature**:
```go
type CalendarEvent struct {
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Start       time.Time `json:"start"`
    End         time.Time `json:"end"`
    AllDay      bool      `json:"allDay"`
    Color       string    `json:"color"` // Token reference
    Description string    `json:"description,omitempty"`
}

type CalendarWidgetProps struct {
    Events       []CalendarEvent `json:"events"`
    CurrentDate  time.Time       `json:"currentDate"`
    View         string          `json:"view"` // "month", "week", "day"
    Selectable   bool            `json:"selectable"`
    Editable     bool            `json:"editable"`
    
    // HTMX for navigation
    HxGet        string          `json:"hxGet"`
}
```

**Composition**:
```
CalendarWidget
├── Header
│   ├── NavigationButtons (prev/next)
│   ├── CurrentPeriod (Month Year)
│   ├── ViewToggle (month/week/day)
│   └── TodayButton
├── Calendar Grid
│   ├── DayHeaders (Sun, Mon, ...)
│   └── For each day:
│       ├── DayCell
│       │   ├── DayNumber
│       │   └── Events
│       │       └── EventBadge (clickable)
│       └── EmptyState (if no events)
└── EventModal (when clicked)
```

---

#### 15. CommentThread 

**Type Signature**:
```go
type Comment struct {
    ID           string     `json:"id"`
    Author       User       `json:"author"`
    Content      string     `json:"content"`
    Timestamp    time.Time  `json:"timestamp"`
    Edited       bool       `json:"edited"`
    Replies      []Comment  `json:"replies,omitempty"`
    Likes        int        `json:"likes"`
    UserLiked    bool       `json:"userLiked"`
}

type CommentThreadProps struct {
    Comments     []Comment  `json:"comments"`
    CurrentUser  User       `json:"currentUser"`
    MaxDepth     int        `json:"maxDepth"` // Max nesting level
    Permissions  []string   `json:"permissions"`
}
```

**Composition**:
```
CommentThread
├── CommentForm (new comment)
│   ├── Avatar
│   ├── Textarea
│   └── SubmitButton
└── CommentList
    └── For each comment (recursive):
        ├── CommentCard
        │   ├── Avatar
        │   ├── Header
        │   │   ├── AuthorName
        │   │   ├── Timestamp
        │   │   └── EditedBadge
        │   ├── Content (markdown)
        │   ├── Actions
        │   │   ├── LikeButton
        │   │   ├── ReplyButton
        │   │   ├── EditButton (if owner)
        │   │   └── DeleteButton (if owner)
        │   └── Replies (nested)
        └── ReplyForm (if replying)
```

**HTMX Optimistic Updates**:
```html
<form hx-post="/comments"
      hx-target="#comments-list"
      hx-swap="afterbegin"
      hx-on::after-request="this.reset()">
    
    <textarea name="content" required></textarea>
    
    <button type="submit">Post Comment</button>
</form>

<div id="comments-list">
    <!-- Comments render here -->
</div>
```

---

#### 16. FileExplorer 

**Type Signature**:
```go
type FileItem struct {
    ID           string    `json:"id"`
    Name         string    `json:"name"`
    Type         string    `json:"type"` // "folder", "file"
    Size         int64     `json:"size"`
    Modified     time.Time `json:"modified"`
    Icon         string    `json:"icon"`
    PreviewURL   string    `json:"previewUrl,omitempty"`
    DownloadURL  string    `json:"downloadUrl,omitempty"`
}

type FileExplorerProps struct {
    Files        []FileItem `json:"files"`
    CurrentPath  string     `json:"currentPath"`
    View         string     `json:"view"` // "list", "grid", "tree"
    Selectable   bool       `json:"selectable"`
    Uploadable   bool       `json:"uploadable"`
    Permissions  []string   `json:"permissions"`
}
```

**Composition**:
```
FileExplorer
├── Toolbar
│   ├── Breadcrumbs (path navigation)
│   ├── SearchBar
│   ├── ViewToggle (list/grid)
│   ├── UploadButton (if uploadable)
│   └── SortSelect
├── FileList/Grid
│   └── For each file:
│       ├── FileCard (grid view) or
│       ├── FileRow (list view)
│       │   ├── Checkbox (if selectable)
│       │   ├── Icon (file type)
│       │   ├── Name
│       │   ├── Size
│       │   ├── Modified
│       │   └── Actions (download, delete, rename)
├── ContextMenu (right-click)
└── UploadZone (drag & drop)
```

---

#### 17. StatsCard 

**Type Signature**:
```go
type StatsCardProps struct {
    Title        string          `json:"title"`
    Value        string          `json:"value"`
    Subtitle     string          `json:"subtitle,omitempty"`
    Icon         string          `json:"icon,omitempty"`
    Trend        TrendDirection  `json:"trend"`
    TrendValue   string          `json:"trendValue,omitempty"`
    TrendLabel   string          `json:"trendLabel,omitempty"` // "vs last month"
    Chart        *MiniChartData  `json:"chart,omitempty"` // Sparkline
    Actions      []ActionButton  `json:"actions,omitempty"`
    Loading      bool            `json:"loading"`
}

type MiniChartData struct {
    Type         string    `json:"type"` // "line", "bar", "area"
    Data         []float64 `json:"data"`
    Labels       []string  `json:"labels,omitempty"`
}
```

**Composition**:
```
StatsCard
├── Header
│   ├── Icon (large, colored background)
│   └── Actions (dropdown menu)
├── Content
│   ├── Title
│   ├── Value (large number)
│   └── Subtitle
├── Trend
│   ├── TrendIcon (up/down arrow)
│   ├── TrendValue (percentage)
│   └── TrendLabel
└── Chart (if exists)
    └── MiniChart (sparkline)
```

---

#### 18. Wizard 

**Type Signature**:
```go
type WizardStep struct {
    ID          string          `json:"id"`
    Title       string          `json:"title"`
    Description string          `json:"description,omitempty"`
    Icon        string          `json:"icon,omitempty"`
    Completed   bool            `json:"completed"`
    Disabled    bool            `json:"disabled"`
}

type WizardProps struct {
    Steps        []WizardStep    `json:"steps"`
    CurrentStep  int             `json:"currentStep"` // 0-indexed
    Linear       bool            `json:"linear"` // Must complete in order
    ShowProgress bool            `json:"showProgress"`
    
    // Navigation
    NextLabel    string          `json:"nextLabel"` // "Next", "Continue"
    BackLabel    string          `json:"backLabel"` // "Back", "Previous"
    FinishLabel  string          `json:"finishLabel"` // "Finish", "Submit"
    
    // HTMX
    HxPost       string          `json:"hxPost,omitempty"` // For final submission
}
```

**Composition**:
```
Wizard
├── StepIndicator
│   ├── ProgressBar (if enabled)
│   └── StepList
│       └── For each step:
│           ├── StepCircle (number/icon)
│           ├── StepTitle
│           └── Connector (line to next)
├── StepContent
│   └── Current step content (slot)
└── Navigation
    ├── BackButton
    ├── CancelButton
    └── NextButton / FinishButton
```

**Alpine.js State Management**:
```html
<div x-data="{
        currentStep: 0,
        steps: {{ jsonEncode(steps) }},
        canGoNext() {
            return this.currentStep < this.steps.length - 1;
        },
        canGoBack() {
            return this.currentStep > 0;
        },
        next() {
            if (this.canGoNext()) {
                this.currentStep++;
            }
        },
        back() {
            if (this.canGoBack()) {
                this.currentStep--;
            }
        },
        goToStep(index) {
            if (!{{ linear }} || this.steps[index].completed) {
                this.currentStep = index;
            }
        }
     }"
     class="wizard">
    
    <!-- Step indicator -->
    <div class="wizard-steps">
        <template x-for="(step, index) in steps" :key="step.id">
            <button @click="goToStep(index)"
                    :class="{ 
                        'active': index === currentStep,
                        'completed': step.completed,
                        'disabled': step.disabled
                    }"
                    class="wizard-step">
                <span x-text="index + 1"></span>
                <span x-text="step.title"></span>
            </button>
        </template>
    </div>
    
    <!-- Step content (slot) -->
    <div class="wizard-content">
        <template x-for="(step, index) in steps" :key="step.id">
            <div x-show="currentStep === index" class="wizard-panel">
                <!-- Content for this step -->
            </div>
        </template>
    </div>
    
    <!-- Navigation -->
    <div class="wizard-nav">
        <button @click="back()" 
                x-show="canGoBack()"
                class="btn-secondary">
            Back
        </button>
        
        <button @click="next()" 
                x-show="canGoNext()"
                class="btn-primary">
            Next
        </button>
        
        <button x-show="currentStep === steps.length - 1"
                class="btn-primary">
            Finish
        </button>
    </div>
</div>
```

---

### Complete Organisms Summary

| # | Organism | Dependencies | Business Logic | Reusability |
|---|----------|-------------|----------------|-------------|
| 1 | DataTable | Checkbox, Button, Badge, Pagination | Sorting, filtering, selection | High |
| 2 | NavigationHeader | SearchBar, Avatar, Badge, Dropdown | User state, notifications | Medium |
| 3 | SidebarNavigation | Button, Icon, Badge | Active state, collapse | Medium |
| 4 | InvoiceCard | StatusBadge, MediaObject, CurrencyDisplay | Permissions, calculations | Medium |
| 5 | ProductCard | Image, Badge, QuantitySelector | Stock status, pricing | Medium |
| 6 | MultiSectionForm | FormField, Button | Section collapse, validation | High |
| 7 | FilterPanel | FormField, FilterChip, DateRangePicker | Filter logic, saved filters | High |
| 8 | ChartWidget | Button, Icon, Spinner | Chart rendering, export | High |
| 9 | Modal | Button, Icon | Open/close state | Very High |
| 10 | Breadcrumbs | BreadcrumbItem | Path tracking | High |
| 11 | UserMenu | Avatar, DropdownMenu | User session, status | Medium |
| 12 | NotificationList | ListItem, Badge | Read/unread state | Medium |
| 13 | Dropdown | Any component | Show/hide state | Very High |
| 14 | CalendarWidget | Button, Badge | Date navigation, events | Medium |
| 15 | CommentThread | Avatar, Textarea, Button | Nested replies, likes | Medium |
| 16 | FileExplorer | Icon, Button, Checkbox | File operations, upload | Medium |
| 17 | StatsCard | Icon, Badge, Chart | Trend calculation | High |
| 18 | Wizard | Button, ProgressBar | Step navigation, validation | Medium |

---

## Level 4: Templates

### Complete Template Catalog (12 Essential Layouts)

#### 9. WizardLayout 

**Type Signature**:
```go
type WizardLayoutProps struct {
    Steps            []WizardStep `json:"steps"`
    CurrentStep      int          `json:"currentStep"`
    ShowSidebar      bool         `json:"showSidebar"` // Step list on side
    ShowProgress     bool         `json:"showProgress"`
    AllowSkip        bool         `json:"allowSkip"`
    StickyNavigation bool         `json:"stickyNavigation"`
}
```

**Composition**:
```
WizardLayout
├── Header
│   ├── Logo
│   └── Title
├── Container
│   ├── Sidebar (if enabled)
│   │   └── StepList (vertical)
│   └── Main
│       ├── ProgressBar (if enabled)
│       ├── StepContent (slot)
│       └── StickyNavigation (if enabled)
│           ├── BackButton
│           ├── SkipButton (if allowed)
│           └── NextButton
└── Footer (optional)
```

---

#### 10. ErrorLayout 

**Type Signature**:
```go
type ErrorType string
const (
    Error404        ErrorType = "404"
    Error403        ErrorType = "403"
    Error500        ErrorType = "500"
    ErrorMaintenance ErrorType = "maintenance"
)

type ErrorLayoutProps struct {
    ErrorType    ErrorType `json:"errorType"`
    Title        string    `json:"title,omitempty"` // Override default
    Message      string    `json:"message,omitempty"` // Override default
    ShowHome     bool      `json:"showHome"`
    ShowBack     bool      `json:"showBack"`
    ShowSupport  bool      `json:"showSupport"`
    CustomAction *ActionButton `json:"customAction,omitempty"`
}
```

**Composition**:
```
ErrorLayout
├── Centered Container
│   ├── ErrorIllustration (large icon/image)
│   ├── ErrorCode (large text)
│   ├── Title
│   ├── Message
│   └── Actions
│       ├── HomeButton
│       ├── BackButton
│       ├── SupportButton
│       └── CustomAction
└── Footer (minimal)
```

**Default Messages**:
```go
var errorDefaults = map[ErrorType]struct{
    Title   string
    Message string
    Icon    string
}{
    Error404: {
        Title:   "Page Not Found",
        Message: "The page you're looking for doesn't exist or has been moved.",
        Icon:    "search-x",
    },
    Error403: {
        Title:   "Access Denied",
        Message: "You don't have permission to access this resource.",
        Icon:    "shield-x",
    },
    Error500: {
        Title:   "Server Error",
        Message: "Something went wrong on our end. We're working to fix it.",
        Icon:    "server-crash",
    },
    ErrorMaintenance: {
        Title:   "Under Maintenance",
        Message: "We're performing scheduled maintenance. Please check back soon.",
        Icon:    "tool",
    },
}
```

---

#### 11. ListDetailLayout 

**Type Signature**:
```go
type ListDetailLayoutProps struct {
    ListWidth      string `json:"listWidth"` // "300px", "33%", "400px"
    Collapsible    bool   `json:"collapsible"`
    InitialState   string `json:"initialState"` // "list", "detail", "both"
    StickyList     bool   `json:"stickyList"`
    ShowBreadcrumbs bool  `json:"showBreadcrumbs"`
}
```

**Composition**:
```
ListDetailLayout
├── Header
│   ├── Breadcrumbs (if enabled)
│   └── Actions
├── Container
│   ├── List (sidebar)
│   │   ├── ListHeader
│   │   │   ├── SearchBar
│   │   │   └── FilterButton
│   │   └── ListContent (scrollable)
│   └── Detail (main)
│       ├── DetailHeader
│       ├── DetailContent
│       └── DetailActions
└── MobileToggle (show list/detail on mobile)
```

**Responsive Behavior**:
```html
<!-- Desktop: Side by side -->
<div class="hidden lg:grid lg:grid-cols-[300px_1fr]">
    <div>List</div>
    <div>Detail</div>
</div>

<!-- Mobile: Toggle between views -->
<div class="lg:hidden">
    <div x-show="view === 'list'">List</div>
    <div x-show="view === 'detail'">
        <button @click="view = 'list'">← Back to list</button>
        Detail
    </div>
</div>
```

---

#### 12. KanbanLayout 

**Type Signature**:
```go
type KanbanColumn struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Color    string `json:"color"` // Token reference
    Limit    int    `json:"limit,omitempty"` // WIP limit
    Collapsed bool  `json:"collapsed"`
}

type KanbanLayoutProps struct {
    Columns      []KanbanColumn `json:"columns"`
    ShowFilters  bool           `json:"showFilters"`
    ShowGroupBy  bool           `json:"showGroupBy"`
    Sortable     bool           `json:"sortable"` // Drag & drop
    
    // HTMX for updates
    HxPost       string         `json:"hxPost"` // Card move endpoint
}
```

**Composition**:
```
KanbanLayout
├── Header
│   ├── Title
│   ├── FilterPanel (if enabled)
│   └── GroupBySelect (if enabled)
├── Board (horizontal scroll)
│   └── For each column:
│       ├── ColumnHeader
│       │   ├── Title
│       │   ├── Count / Limit
│       │   └── CollapseButton
│       ├── ColumnContent (cards)
│       │   └── Card slots (draggable)
│       └── AddCardButton
└── FloatingActions
    └── CreateCardButton
```

**Sortable.js Integration**:
```html
<script>
// Initialize drag & drop
document.querySelectorAll('.kanban-column').forEach(column => {
    new Sortable(column, {
        group: 'shared',
        animation: 150,
        onEnd: function(evt) {
            // Send HTMX request to update card position
            htmx.ajax('POST', '/api/cards/move', {
                values: {
                    cardId: evt.item.dataset.id,
                    fromColumn: evt.from.dataset.columnId,
                    toColumn: evt.to.dataset.columnId,
                    position: evt.newIndex
                }
            });
        }
    });
});
</script>
```

---

### Complete Templates Summary

| # | Template | Organisms Used | Use Case | Responsive Strategy |
|---|----------|---------------|----------|-------------------|
| 1 | MasterDetailLayout | List, Detail | CRUD operations | Stack on mobile |
| 2 | DashboardLayout | Navigation, Widgets | Dashboards | Grid reflow |
| 3 | FormLayout | Form, Breadcrumbs | Create/Edit | Single column |
| 4 | ReportLayout | Filters, Charts, Table | Reports | Filters to drawer |
| 5 | ContentLayout | Article, TOC | Documentation | TOC to drawer |
| 6 | AuthLayout | Form | Login/Register | Center always |
| 7 | SplitViewLayout | Various | Side-by-side | Stack on mobile |
| 8 | SettingsLayout | Tabs, Forms | Settings | Tabs to dropdown |
| 9 | WizardLayout | Wizard, Steps | Multi-step | Linear flow |
| 10 | ErrorLayout | EmptyState | Error pages | Center always |
| 11 | ListDetailLayout | List, Detail | Email/Messages | Toggle views |
| 12 | KanbanLayout | Cards, Columns | Project management | Horizontal scroll |

---

## Level 5: Pages

### Complete Page Catalog (15 Essential Pages)

I'll add the missing pages:

#### 6. InvoiceCreatePage 

**Route**: `/invoices/new`

**Type Signature**:
```go
type InvoiceCreatePageProps struct {
    Customers    []Customer      `json:"customers"`
    Products     []Product       `json:"products"`
    TaxRates     []TaxRate       `json:"taxRates"`
    DefaultCurrency string       `json:"defaultCurrency"`
    NextNumber   string          `json:"nextNumber"` // Auto-generated invoice number
    User         User            `json:"user"`
    Template     *Invoice        `json:"template,omitempty"` // For duplicating
}
```

**Composition**:
```
InvoiceCreatePage
└── FormLayout (wizard variant)
    ├── slot("header")
    │   ├── Breadcrumbs
    │   └── Title ("Create Invoice")
    ├── slot("progress")
    │   └── StepIndicator (Customer → Items → Details → Review)
    ├── slot("form")
    │   └── MultiSectionForm
    │       ├── Step 1: Customer Section
    │       │   ├── CustomerSelect (searchable)
    │       │   ├── BillingAddress
    │       │   └── ShippingAddress
    │       ├── Step 2: Line Items Section
    │       │   ├── ProductSearch
    │       │   ├── LineItemsTable
    │       │   │   └── For each item:
    │       │   │       ├── ProductSelect
    │       │   │       ├── QuantityInput
    │       │   │       ├── PriceInput
    │       │   │       ├── TaxSelect
    │       │   │       └── RemoveButton
    │       │   ├── AddItemButton
    │       │   └── Summary
    │       │       ├── Subtotal
    │       │       ├── Tax
    │       │       └── Total
    │       ├── Step 3: Details Section
    │       │   ├── InvoiceNumber
    │       │   ├── InvoiceDate
    │       │   ├── DueDate
    │       │   ├── Terms (select)
    │       │   ├── Notes (textarea)
    │       │   └── AttachFile
    │       └── Step 4: Review Section
    │           └── InvoicePreview
    └── slot("actions")
        ├── SaveDraftButton
        └── SendInvoiceButton
```

**HTMX Dynamic Line Items**:
```html
<div id="line-items">
    <!-- Existing items -->
    
    <button hx-get="/invoices/line-item-row"
            hx-target="#line-items"
            hx-swap="beforeend">
        + Add Item
    </button>
</div>

<!-- Server returns new row HTML -->
<tr hx-target="this" hx-swap="outerHTML">
    <td>@ProductSelect(...)</td>
    <td>@QuantityInput(...)</td>
    <td>@CurrencyInput(...)</td>
    <td>@TaxSelect(...)</td>
    <td><span class="calculated-total">0.00</span></td>
    <td>
        <button hx-delete="/invoices/line-item/temp-{{id}}"
                hx-confirm="Remove this item?">
            Remove
        </button>
    </td>
</tr>
```

**Alpine.js Calculator**:
```html
<div x-data="{
        items: [],
        taxRate: 0.16,
        get subtotal() {
            return this.items.reduce((sum, item) => 
                sum + (item.quantity * item.price), 0
            );
        },
        get tax() {
            return this.subtotal * this.taxRate;
        },
        get total() {
            return this.subtotal + this.tax;
        }
     }">
    
    <!-- Item rows -->
    
    <!-- Summary -->
    <div class="invoice-summary">
        <div>Subtotal: <span x-text="subtotal.toFixed(2)"></span></div>
        <div>Tax (16%): <span x-text="tax.toFixed(2)"></span></div>
        <div class="total">Total: <span x-text="total.toFixed(2)"></span></div>
    </div>
</div>
```

---

#### 7. CustomerListPage 

**Route**: `/customers`

**Type Signature**:
```go
type CustomerListPageProps struct {
    Customers    []Customer       `json:"customers"`
    Pagination   PaginationState  `json:"pagination"`
    Filters      CustomerFilters  `json:"filters"`
    Sort         SortState        `json:"sort"`
    Stats        CustomerStats    `json:"stats"`
    User         User             `json:"user"`
    Permissions  []string         `json:"permissions"`
}

type CustomerFilters struct {
    Search       string   `json:"search,omitempty"`
    Status       []string `json:"status,omitempty"` // "active", "inactive"
    Type         []string `json:"type,omitempty"` // "individual", "company"
    Tags         []string `json:"tags,omitempty"`
    CreatedFrom  time.Time `json:"createdFrom,omitempty"`
    CreatedTo    time.Time `json:"createdTo,omitempty"`
}

type CustomerStats struct {
    Total        int             `json:"total"`
    Active       int             `json:"active"`
    Inactive     int             `json:"inactive"`
    TotalRevenue decimal.Decimal `json:"totalRevenue"`
}
```

**Composition**:
```
CustomerListPage
└── MasterDetailLayout
    ├── slot("header")
    │   ├── PageTitle ("Customers")
    │   ├── StatsRow
    │   │   ├── StatDisplay (Total)
    │   │   ├── StatDisplay (Active)
    │   │   ├── StatDisplay (Inactive)
    │   │   └── StatDisplay (Revenue)
    │   └── CreateButton (if has permission)
    ├── slot("master")
    │   ├── SearchBar (instant search)
    │   ├── FilterPanel
    │   │   ├── StatusFilter (checkboxes)
    │   │   ├── TypeFilter (radio)
    │   │   ├── TagsFilter (multi-select)
    │   │   └── DateRangeFilter
    │   └── DataTable
    │       └── Columns:
    │           ├── Checkbox (select)
    │           ├── Customer (name + email)
    │           ├── Type (badge)
    │           ├── Status (status badge)
    │           ├── Revenue (currency)
    │           ├── Invoices (count)
    │           └── Actions (dropdown)
    ├── slot("detail")
    │   └── CustomerDetailPreview
    │       ├── Header (avatar, name, status)
    │       ├── Tabs
    │       │   ├── Overview
    │       │   │   ├── ContactInfo
    │       │   │   ├── Billing Address
    │       │   │   └── QuickStats
    │       │   ├── Invoices
    │       │   │   └── InvoiceList (mini)
    │       │   ├── Payments
    │       │   │   └── PaymentHistory
    │       │   └── Activity
    │       │       └── ActivityLog
    │       └── Actions
    │           ├── EditButton
    │           ├── SendEmailButton
    │           └── MoreButton (dropdown)
    └── slot("toolbar")
        └── BulkActions (if selections)
            ├── ExportButton
            ├── EmailButton
            ├── TagButton
            └── DeleteButton
```

---

#### 8. SettingsPage 

**Route**: `/settings`

**Type Signature**:
```go
type SettingsPageProps struct {
    Tabs         []SettingsTab    `json:"tabs"`
    ActiveTab    string           `json:"activeTab"`
    Settings     map[string]any   `json:"settings"`
    User         User             `json:"user"`
    Permissions  []string         `json:"permissions"`
}

type SettingsTab struct {
    ID          string `json:"id"`
    Label       string `json:"label"`
    Icon        string `json:"icon"`
    Description string `json:"description,omitempty"`
}
```

**Composition**:
```
SettingsPage
└── SettingsLayout
    ├── slot("header")
    │   └── PageTitle ("Settings")
    ├── slot("sidebar")
    │   └── SettingsNavigation
    │       ├── General Tab
    │       ├── Profile Tab
    │       ├── Company Tab
    │       ├── Billing Tab
    │       ├── Team Tab
    │       ├── Integrations Tab
    │       ├── Security Tab
    │       └── Notifications Tab
    └── slot("content")
        └── Switch on ActiveTab:
            ├── GeneralSettings
            │   ├── Section: Application
            │   │   ├── AppName
            │   │   ├── AppURL
            │   │   ├── Timezone
            │   │   └── Language
            │   ├── Section: Regional
            │   │   ├── Currency
            │   │   ├── DateFormat
            │   │   └── NumberFormat
            │   └── SaveButton
            ├── ProfileSettings
            │   ├── Avatar Upload
            │   ├── Personal Info
            │   ├── Password Change
            │   └── 2FA Setup
            ├── CompanySettings
            │   ├── Company Info
            │   ├── Logo Upload
            │   ├── Addresses
            │   └── Tax Info
            ├── BillingSettings
            │   ├── Subscription Info
            │   ├── Payment Method
            │   ├── Invoice History
            │   └── Usage Stats
            ├── TeamSettings
            │   ├── Member List
            │   ├── Invite Members
            │   └── Roles & Permissions
            ├── IntegrationsSettings
            │   └── For each integration:
            │       ├── Integration Card
            │       ├── ConnectButton
            │       └── Configuration
            ├── SecuritySettings
            │   ├── 2FA Status
            │   ├── Active Sessions
            │   ├── Login History
            │   └── API Keys
            └── NotificationsSettings
                └── For each category:
                    ├── Category Toggle
                    └── Channels (email, SMS, push)
```

**Auto-save with HTMX**:
```html
<form hx-post="/settings/general"
      hx-trigger="change delay:500ms"
      hx-target="#save-status"
      hx-swap="innerHTML">
    
    <input name="appName" value="...">
    
    <div id="save-status" class="text-sm text-gray-500">
        <!-- Status updates here: "Saving...", "Saved", "Error" -->
    </div>
</form>
```

---

#### 9. LoginPage 

**Route**: `/login`

**Type Signature**:
```go
type LoginPageProps struct {
    RedirectURL  string `json:"redirectUrl,omitempty"`
    Error        string `json:"error,omitempty"`
    Email        string `json:"email,omitempty"` // Pre-filled email
    ShowSSO      bool   `json:"showSso"`
    SSOProviders []SSOProvider `json:"ssoProviders,omitempty"`
    ShowSignup   bool   `json:"showSignup"`
}

type SSOProvider struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Icon    string `json:"icon"`
    AuthURL string `json:"authUrl"`
}
```

**Composition**:
```
LoginPage
└── AuthLayout
    ├── slot("logo")
    │   └── CompanyLogo
    ├── slot("form")
    │   └── LoginForm
    │       ├── Title ("Welcome back")
    │       ├── Error Alert (if error)
    │       ├── SSO Buttons (if enabled)
    │       │   └── For each provider:
    │       │       └── SSOButton (Google, Microsoft, etc.)
    │       ├── Divider ("or")
    │       ├── Email FormField
    │       ├── Password FormField
    │       │   └── ForgotPasswordLink
    │       ├── RememberMe Checkbox
    │       └── LoginButton (loading state)
    │   
    │   if ShowSignup {
    │       ├── Divider
    │       └── SignupPrompt ("Don't have an account? Sign up")
    │   }
    └── slot("footer")
        ├── Terms Link
        └── Privacy Link
```

**HTMX Form Submission**:
```html
<form hx-post="/api/auth/login"
      hx-target="#error-container"
      hx-swap="innerHTML"
      hx-disabled-elt="#login-button"
      class="auth-form">
    
    <div id="error-container">
        <!-- Errors render here -->
    </div>
    
    @FormField(FormFieldProps{
        Name: "email",
        Type: InputEmail,
        Label: "Email",
        Required: true,
        AutoComplete: "username",
    })
    
    @FormField(FormFieldProps{
        Name: "password",
        Type: InputPassword,
        Label: "Password",
        Required: true,
        AutoComplete: "current-password",
    })
    
    @Button(ButtonProps{
        ID: "login-button",
        Label: "Sign In",
        Type: "submit",
        Variant: ButtonPrimary,
        FullWidth: true,
    })
</form>
```

---

#### 10. DashboardHomePage 

**Route**: `/dashboard` or `/`

**Type Signature**:
```go
type DashboardHomePageProps struct {
    Widgets      []WidgetConfig  `json:"widgets"`
    User         User            `json:"user"`
    Permissions  []string        `json:"permissions"`
    DateRange    DateRange       `json:"dateRange"`
    Customizable bool            `json:"customizable"` // User can rearrange
}

type WidgetConfig struct {
    ID       string         `json:"id"`
    Type     string         `json:"type"` // "stat", "chart", "table", "list"
    Title    string         `json:"title"`
    Span     int            `json:"span"` // Grid columns (1-12)
    Height   string         `json:"height"`
    Data     any            `json:"data"`
    Refresh  int            `json:"refresh,omitempty"` // Seconds (auto-refresh)
}
```

**Composition**:
```
DashboardHomePage
└── DashboardLayout
    ├── slot("header")
    │   ├── GreetingMessage ("Good morning, John")
    │   ├── DateRangePicker (global filter)
    │   └── CustomizeButton (if customizable)
    ├── slot("sidebar")
    │   └── QuickActions
    │       ├── CreateInvoiceButton
    │       ├── AddCustomerButton
    │       ├── RecordPaymentButton
    │       └── NewProductButton
    └── slot("main")
        └── WidgetGrid (responsive, draggable if customizable)
            ├── RevenueWidget (stat + chart)
            │   ├── StatDisplay (large)
            │   └── MiniChart (line)
            ├── InvoicesWidget (stat + trend)
            ├── CustomersWidget (stat + trend)
            ├── SalesChartWidget (full chart)
            │   └── ChartWidget (bar/line)
            ├── RecentInvoicesWidget (table)
            │   └── Mini DataTable
            ├── PendingPaymentsWidget (list)
            │   └── Payment list with actions
            ├── TopProductsWidget (chart)
            │   └── ChartWidget (pie/bar)
            └── ActivityFeedWidget (timeline)
                └── Activity items with timestamps
```

**Draggable Grid (Gridstack.js)**:
```html
<div class="grid-stack">
    for _, widget := range widgets {
        <div class="grid-stack-item"
             gs-w="{{ widget.Span }}"
             gs-h="{{ widget.Height }}"
             gs-id="{{ widget.ID }}">
            
            <div class="grid-stack-item-content">
                @RenderWidget(widget)
            </div>
        </div>
    }
</div>

<script>
GridStack.init({
    float: true,
    cellHeight: 70,
    minRow: 1
}).on('change', function(event, items) {
    // Save layout via HTMX
    htmx.ajax('POST', '/dashboard/layout', {
        values: { layout: JSON.stringify(items) }
    });
});
</script>
```

---

#### 11. ProductDetailPage 

**Route**: `/products/:id`

**Type Signature**:
```go
type ProductDetailPageProps struct {
    Product       Product         `json:"product"`
    Variants      []ProductVariant `json:"variants,omitempty"`
    Inventory     []StockLevel    `json:"inventory"`
    PriceHistory  []PricePoint    `json:"priceHistory"`
    SalesData     SalesAnalytics  `json:"salesData"`
    RelatedProducts []Product     `json:"relatedProducts"`
    User          User            `json:"user"`
    Permissions   []string        `json:"permissions"`
}

type ProductVariant struct {
    ID          string          `json:"id"`
    SKU         string          `json:"sku"`
    Name        string          `json:"name"`
    Options     map[string]string `json:"options"` // {"size": "L", "color": "Blue"}
    Price       decimal.Decimal `json:"price"`
    Stock       int             `json:"stock"`
    ImageURL    string          `json:"imageUrl,omitempty"`
}
```

**Composition**:
```
ProductDetailPage
└── DashboardLayout
    ├── slot("header")
    │   ├── Breadcrumbs
    │   └── NavigationHeader
    ├── slot("main")
    │   ├── ProductHeader
    │   │   ├── ImageGallery (main + thumbnails)
    │   │   ├── ProductInfo
    │   │   │   ├── Title
    │   │   │   ├── SKU + StatusBadge
    │   │   │   ├── Category + Tags
    │   │   │   ├── Price (large)
    │   │   │   └── Description
    │   │   └── QuickActions
    │   │       ├── EditButton
    │   │       ├── DuplicateButton
    │   │       └── MoreButton
    │   │
    │   ├── Tabs
    │   │   ├── Overview Tab
    │   │   │   ├── DetailsSection
    │   │   │   │   ├── Specifications
    │   │   │   │   ├── Dimensions
    │   │   │   │   └── Weight
    │   │   │   ├── InventorySection
    │   │   │   │   ├── Stock Levels (by location)
    │   │   │   │   ├── Reorder Point
    │   │   │   │   └── Supplier Info
    │   │   │   └── PricingSection
    │   │   │       ├── Cost Price
    │   │   │       ├── Selling Price
    │   │   │       └── Profit Margin
    │   │   │
    │   │   ├── Variants Tab (if has variants)
    │   │   │   └── VariantsTable
    │   │   │       └── For each variant:
    │   │   │           ├── Options
    │   │   │           ├── SKU
    │   │   │           ├── Price
    │   │   │           ├── Stock
    │   │   │           └── Actions
    │   │   │
    │   │   ├── Sales Tab
    │   │   │   ├── SalesStats
    │   │   │   │   ├── TotalSold
    │   │   │   │   ├── Revenue
    │   │   │   │   └── AvgOrderValue
    │   │   │   ├── SalesChart (over time)
    │   │   │   └── TopCustomers
    │   │   │
    │   │   ├── Inventory Tab
    │   │   │   ├── StockMovements (table)
    │   │   │   ├── AdjustmentForm
    │   │   │   └── TransferForm
    │   │   │
    │   │   └── History Tab
    │   │       └── AuditLog
    │   │
    │   └── RelatedProducts (sidebar/bottom)
    │       └── ProductGrid (mini cards)
    │
    └── slot("actions")
        └── FloatingActionButton (edit on mobile)
```

---

#### 12. UserProfilePage 

**Route**: `/profile` or `/users/:id`

**Type Signature**:
```go
type UserProfilePageProps struct {
    User         User            `json:"user"`
    IsOwnProfile bool            `json:"isOwnProfile"`
    Stats        UserStats       `json:"stats"`
    Activity     []ActivityItem  `json:"activity"`
    Permissions  []string        `json:"permissions"`
}

type UserStats struct {
    InvoicesCreated int       `json:"invoicesCreated"`
    TotalRevenue    decimal.Decimal `json:"totalRevenue"`
    CustomersAdded  int       `json:"customersAdded"`
    LastLogin       time.Time `json:"lastLogin"`
    MemberSince     time.Time `json:"memberSince"`
}
```

**Composition**:
```
UserProfilePage
└── ContentLayout
    ├── slot("header")
    │   └── ProfileHeader
    │       ├── Avatar (large)
    │       ├── UserInfo
    │       │   ├── Name
    │       │   ├── Role Badge
    │       │   ├── Email
    │       │   └── Status (online/offline)
    │       └── Actions (if own profile)
    │           └── EditProfileButton
    │
    ├── slot("sidebar")
    │   ├── StatsCard
    │   │   ├── MemberSince
    │   │   ├── LastLogin
    │   │   └── Role
    │   └── ContactInfo
    │       ├── Email
    │       ├── Phone
    │       └── Location
    │
    └── slot("main")
        └── Tabs
            ├── Overview Tab
            │   ├── AboutSection
            │   ├── StatsGrid
            │   │   ├── InvoicesCreated
            │   │   ├── Revenue
            │   │   └── Customers
            │   └── RecentActivity
            │
            ├── Activity Tab
            │   └── ActivityFeed (full)
            │
            ├── Settings Tab (if own profile)
            │   ├── PersonalInfo Form
            │   ├── ChangePassword Form
            │   ├── NotificationPreferences
            │   └── PrivacySettings
            │
            └── Security Tab (if own profile)
                ├── 2FA Setup
                ├── Active Sessions
                └── Login History
```

---

#### 13. ReportsIndexPage 

**Route**: `/reports`

**Type Signature**:
```go
type ReportCategory struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Icon        string   `json:"icon"`
    Reports     []ReportInfo `json:"reports"`
}

type ReportInfo struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    URL         string `json:"url"`
    Icon        string `json:"icon"`
    Popular     bool   `json:"popular"`
}

type ReportsIndexPageProps struct {
    Categories   []ReportCategory `json:"categories"`
    Recent       []ReportInfo     `json:"recent"`
    Favorites    []ReportInfo     `json:"favorites"`
    User         User             `json:"user"`
}
```

**Composition**:
```
ReportsIndexPage
└── DashboardLayout
    ├── slot("header")
    │   ├── PageTitle ("Reports")
    │   └── SearchBar
    │
    └── slot("main")
        ├── QuickAccess
        │   ├── RecentReports
        │   │   └── For each report:
        │   │       └── ReportCard (mini)
        │   └── FavoriteReports
        │       └── For each report:
        │           └── ReportCard (mini)
        │
        └── Categories
            └── For each category:
                ├── CategoryHeader
                │   ├── Icon
                │   ├── Name
                │   └── Description
                └── ReportGrid
                    └── For each report:
                        └── ReportCard
                            ├── Icon
                            ├── Name
                            ├── Description
                            ├── PopularBadge (if popular)
                            └── Actions
                                ├── ViewButton
                                └── FavoriteButton
```

**Report Categories**:
```go
var reportCategories = []ReportCategory{
    {
        ID:          "sales",
        Name:        "Sales Reports",
        Description: "Revenue, invoices, and sales analytics",
        Icon:        "trending-up",
        Reports: []ReportInfo{
            {ID: "sales-summary", Name: "Sales Summary"},
            {ID: "invoice-aging", Name: "Invoice Aging"},
            {ID: "revenue-by-customer", Name: "Revenue by Customer"},
            {ID: "sales-by-product", Name: "Sales by Product"},
        },
    },
    {
        ID:          "inventory",
        Name:        "Inventory Reports",
        Description: "Stock levels, movements, and valuation",
        Icon:        "package",
        Reports: []ReportInfo{
            {ID: "stock-levels", Name: "Current Stock Levels"},
            {ID: "stock-movements", Name: "Stock Movements"},
            {ID: "inventory-valuation", Name: "Inventory Valuation"},
            {ID: "low-stock", Name: "Low Stock Alert"},
        },
    },
    {
        ID:          "financial",
        Name:        "Financial Reports",
        Description: "P&L, balance sheet, cash flow",
        Icon:        "dollar-sign",
        Reports: []ReportInfo{
            {ID: "profit-loss", Name: "Profit & Loss"},
            {ID: "balance-sheet", Name: "Balance Sheet"},
            {ID: "cash-flow", Name: "Cash Flow Statement"},
            {ID: "trial-balance", Name: "Trial Balance"},
        },
    },
    {
        ID:          "customers",
        Name:        "Customer Reports",
        Description: "Customer activity and insights",
        Icon:        "users",
        Reports: []ReportInfo{
            {ID: "customer-list", Name: "Customer List"},
            {ID: "customer-aging", Name: "Customer Aging"},
            {ID: "top-customers", Name: "Top Customers"},
        },
    },
}
```

---

#### 14. SearchResultsPage 

**Route**: `/search?q=...`

**Type Signature**:
```go
type SearchResult struct {
    Type        string    `json:"type"` // "invoice", "customer", "product"
    ID          string    `json:"id"`
    Title       string    `json:"title"`
    Subtitle    string    `json:"subtitle"`
    Snippet     string    `json:"snippet"` // Highlighted excerpt
    URL         string    `json:"url"`
    Icon        string    `json:"icon"`
    Metadata    []string  `json:"metadata"` // ["2 days ago", "Draft"]
    Relevance   float64   `json:"relevance"`
}

type SearchResultsPageProps struct {
    Query       string         `json:"query"`
    Results     []SearchResult `json:"results"`
    Total       int            `json:"total"`
    Filters     SearchFilters  `json:"filters"`
    Suggestions []string       `json:"suggestions"`
    Time        time.Duration  `json:"time"` // Search time
}

type SearchFilters struct {
    Types       []string   `json:"types,omitempty"`
    DateFrom    time.Time  `json:"dateFrom,omitempty"`
    DateTo      time.Time  `json:"dateTo,omitempty"`
}
```

**Composition**:
```
SearchResultsPage
└── DashboardLayout
    ├── slot("header")
    │   └── SearchBar (with current query)
    │
    ├── slot("sidebar")
    │   └── FilterPanel
    │       ├── TypeFilter
    │       │   ├── Invoices (count)
    │       │   ├── Customers (count)
    │       │   ├── Products (count)
    │       │   └── All (count)
    │       └── DateFilter
    │
    └── slot("main")
        ├── ResultsHeader
        │   ├── ResultCount ("23 results")
        │   ├── SearchTime ("0.14s")
        │   └── SortSelect
        │
        if len(suggestions) > 0 {
            ├── Suggestions
            │   └── "Did you mean: {suggestion}?"
        }
        │
        ├── ResultsList
        │   └── For each result:
        │       └── SearchResultCard
        │           ├── Icon (type-based)
        │           ├── Content
        │           │   ├── Type Badge
        │           │   ├── Title (highlighted)
        │           │   ├── Snippet (highlighted)
        │           │   └── Metadata
        │           └── Actions
        │               └── ViewButton
        │
        └── Pagination
```

**Highlight Implementation**:
```go
// Highlight search terms in results
func HighlightSnippet(text, query string) string {
    words := strings.Split(query, " ")
    for _, word := range words {
        text = strings.ReplaceAll(
            text,
            word,
            fmt.Sprintf("<mark>%s</mark>", word),
        )
    }
    return text
}
```

---

#### 15. NotFoundPage (404) 

**Route**: `*` (catch-all)

**Type Signature**:
```go
type NotFoundPageProps struct {
    Path         string   `json:"path"` // Requested path
    Suggestions  []string `json:"suggestions,omitempty"` // Similar paths
    ShowSearch   bool     `json:"showSearch"`
    ShowHome     bool     `json:"showHome"`
}
```

**Composition**:
```
NotFoundPage
└── ErrorLayout
    ├── slot("content")
    │   ├── ErrorCode ("404")
    │   ├── Title ("Page Not Found")
    │   ├── Message
    │   │   └── "The page '{path}' doesn't exist or has been moved."
    │   │
    │   if len(suggestions) > 0 {
    │       └── Suggestions
    │           ├── Heading ("Did you mean:")
    │           └── SuggestionList
    │               └── For each suggestion:
    │                   └── Link to suggested page
    │   }
    │   │
    │   if showSearch {
    │       └── SearchBar ("Search for what you need")
    │   }
    │   │
    │   └── Actions
    │       ├── HomeButton ("Go to Dashboard")
    │       └── BackButton ("Go Back")
    │
    └── slot("footer")
        └── SupportLink ("Need help? Contact support")
```

---

### Complete Pages Summary

| # | Page | Route | Template | Data Sources | Auth Required |
|---|------|-------|----------|-------------|---------------|
| 1 | InvoiceListPage | /invoices | MasterDetail | PostgreSQL + Filters | Yes |
| 2 | InvoiceDetailPage | /invoices/:id | Dashboard | PostgreSQL + Related | Yes |
| 3 | InvoiceCreatePage | /invoices/new | Form (Wizard) | Lookups + Schema | Yes |
| 4 | ProductCatalogPage | /products | MasterDetail | PostgreSQL + Filters | Yes |
| 5 | ProductDetailPage | /products/:id | Dashboard | PostgreSQL + Analytics | Yes |
| 6 | StationDashboardPage | /dashboard/station/:id | Dashboard | Real-time + Analytics | Yes |
| 7 | SalesReportPage | /reports/sales | Report | Analytics DB | Yes |
| 8 | CustomerListPage | /customers | MasterDetail | PostgreSQL + Filters | Yes |
| 9 | SettingsPage | /settings | Settings | User Preferences | Yes |
| 10 | LoginPage | /login | Auth | Session | No |
| 11 | DashboardHomePage | / | Dashboard | Analytics + Quick Stats | Yes |
| 12 | UserProfilePage | /profile | Content | User Data + Activity | Yes |
| 13 | ReportsIndexPage | /reports | Dashboard | Report Registry | Yes |
| 14 | SearchResultsPage | /search | Dashboard | Search Index | Yes |
| 15 | NotFoundPage | * | Error | - | No |

---

## Accessibility Guidelines

### WCAG 2.1 AA Compliance

#### Color Contrast
```go
// Minimum contrast ratios in token system
var contrastRatios = map[string]float64{
    "text.primary":   4.5,  // Normal text
    "text.large":     3.0,  // Large text (18pt+ or 14pt+ bold)
    "ui.controls":    3.0,  // Buttons, form controls
    "focus.indicator": 3.0,  // Focus outlines
}
```

#### Keyboard Navigation
```html
<!-- All interactive elements must be keyboard accessible -->
<button tabindex="0" 
        @keydown.enter="handleClick()"
        @keydown.space.prevent="handleClick()">
    Action
</button>

<!-- Skip links for screen readers -->
<a href="#main-content" class="skip-link">
    Skip to main content
</a>
```

#### ARIA Landmarks
```html
<header role="banner">
    <nav role="navigation" aria-label="Main navigation">
        <!-- Navigation items -->
    </nav>
</header>

<main role="main" id="main-content">
    <!-- Main content -->
</main>

<aside role="complementary" aria-label="Filters">
    <!-- Filters -->
</aside>

<footer role="contentinfo">
    <!-- Footer -->
</footer>
```

#### Form Accessibility
```html
<!-- Associate labels with inputs -->
<label for="email-input">Email Address</label>
<input id="email-input" 
       type="email"
       aria-required="true"
       aria-invalid="false"
       aria-describedby="email-error email-help">

<div id="email-help" class="help-text">
    We'll never share your email
</div>

<div id="email-error" class="error-text" role="alert">
    <!-- Error message if invalid -->
</div>
```

#### Live Regions
```html
<!-- For dynamic content updates -->
<div role="status" aria-live="polite" aria-atomic="true">
    {{ statusMessage }}
</div>

<div role="alert" aria-live="assertive">
    {{ errorMessage }}
</div>
```

---

## Responsive Design Patterns

### Breakpoint System
```go
// Token definitions
var breakpoints = map[string]string{
    "xs": "0px",      // Mobile portrait
    "sm": "640px",    // Mobile landscape
    "md": "768px",    // Tablet portrait
    "lg": "1024px",   // Tablet landscape / Desktop
    "xl": "1280px",   // Desktop
    "2xl": "1536px",  // Large desktop
}
```

### Mobile-First Approach
```css
/* Base styles (mobile) */
.container {
    padding: 1rem;
}

/* Tablet and up */
@media (min-width: 768px) {
    .container {
        padding: 2rem;
        max-width: 1200px;
    }
}

/* Desktop and up */
@media (min-width: 1024px) {
    .container {
        padding: 3rem;
    }
}
```

### Responsive Patterns

#### 1. Stack to Sidebar
```html
<!-- Mobile: Stacked -->
<div class="flex flex-col md:flex-row">
    <aside class="w-full md:w-64">
        <!-- Sidebar -->
    </aside>
    <main class="w-full md:flex-1">
        <!-- Main content -->
    </main>
</div>
```

#### 2. Grid Reflow
```html
<!-- Responsive grid -->
<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4">
    <!-- Cards -->
</div>
```

#### 3. Hamburger Menu
```html
<nav x-data="{ open: false }">
    <!-- Mobile: Hamburger -->
    <button @click="open = !open" 
            class="md:hidden">
        @Icon(IconProps{Name: "menu"})
    </button>
    
    <!-- Mobile: Drawer -->
    <div x-show="open" 
         class="fixed inset-0 z-50 md:hidden">
        <div @click="open = false" 
             class="absolute inset-0 bg-black/50"></div>
        <nav class="absolute left-0 top-0 h-full w-64 bg-white">
            <!-- Menu items -->
        </nav>
    </div>
    
    <!-- Desktop: Always visible -->
    <nav class="hidden md:flex">
        <!-- Menu items -->
    </nav>
</nav>
```

#### 4. Table to Cards
```html
<!-- Desktop: Table -->
<table class="hidden lg:table">
    <!-- Table rows -->
</table>

<!-- Mobile: Cards -->
<div class="lg:hidden space-y-4">
    <div class="card">
        <!-- Card layout of row data -->
    </div>
</div>
```

#### 5. Adaptive Typography
```go
// Token system for fluid typography
var typography = map[string]map[string]string{
    "h1": {
        "mobile":  "clamp(2rem, 5vw, 3rem)",
        "desktop": "3rem",
    },
    "body": {
        "mobile":  "0.875rem",  // 14px
        "desktop": "1rem",      // 16px
    },
}
```

---

## Component Registry 

### Registration System

```go
// registry/registry.go
package registry

type ComponentRegistry struct {
    components map[string]ComponentFactory
    tokens     *schema.TokenRegistry
    mu         sync.RWMutex
}

type ComponentFactory func(
    ctx context.Context,
    props map[string]any,
    children ...templ.Component,
) (templ.Component, error)

func NewRegistry() *ComponentRegistry {
    r := &ComponentRegistry{
        components: make(map[string]ComponentFactory),
        tokens:     schema.NewTokenRegistry(),
    }
    
    // Register all components
    r.registerAtoms()
    r.registerMolecules()
    r.registerOrganisms()
    
    return r
}

func (r *ComponentRegistry) Register(name string, factory ComponentFactory) {
    r.mu.Lock()
    defer r.mu.Unlock()
    r.components[name] = factory
}

func (r *ComponentRegistry) Get(name string) (ComponentFactory, bool) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    factory, exists := r.components[name]
    return factory, exists
}

func (r *ComponentRegistry) Render(
    ctx context.Context,
    name string,
    props map[string]any,
    children ...templ.Component,
) (templ.Component, error) {
    factory, exists := r.Get(name)
    if !exists {
        return nil, fmt.Errorf("component not found: %s", name)
    }
    
    return factory(ctx, props, children...)
}
```

### Schema-Driven Rendering

```go
// Render a page from schema
func RenderFromSchema(
    ctx context.Context,
    schema *schema.PageSchema,
    data map[string]any,
) (templ.Component, error) {
    registry := GetRegistryFromContext(ctx)
    
    // Render layout
    layout, err := registry.Render(ctx, schema.Layout, schema.LayoutProps)
    if err != nil {
        return nil, err
    }
    
    // Render sections
    sections := make([]templ.Component, 0, len(schema.Sections))
    for _, section := range schema.Sections {
        // Apply conditions
        if !evaluateConditions(ctx, section.Conditions, data) {
            continue
        }
        
        // Render component
        component, err := registry.Render(
            ctx,
            section.Component,
            mergeProps(section.Props, data),
        )
        if err != nil {
            return nil, err
        }
        
        sections = append(sections, component)
    }
    
    // Compose final page
    return composePage(layout, sections), nil
}
```

---

## Usage Patterns 

### Pattern 1: HTMX Infinite Scroll

```html
<div id="product-list">
    <!-- Initial products -->
    
    <div hx-get="/products?page=2"
         hx-trigger="revealed"
         hx-swap="afterend"
         hx-indicator="#loading">
        <!-- Trigger element -->
    </div>
</div>

<div id="loading" class="htmx-indicator">
    @Spinner(SpinnerProps{Size: SpinnerMD})
</div>
```

### Pattern 2: Alpine.js Form Wizard

```html
<div x-data="formWizard()">
    <!-- Steps indicator -->
    <div class="steps">
        <template x-for="(step, index) in steps">
            <div :class="{ 'active': currentStep === index }">
                <span x-text="step.title"></span>
            </div>
        </template>
    </div>
    
    <!-- Current step content -->
    <div x-show="currentStep === 0">
        @CustomerForm()
    </div>
    
    <div x-show="currentStep === 1">
        @LineItemsForm()
    </div>
    
    <!-- Navigation -->
    <div class="wizard-nav">
        <button @click="prev()" x-show="currentStep > 0">
            Previous
        </button>
        
        <button @click="next()" 
                x-show="currentStep < steps.length - 1"
                :disabled="!canProceed()">
            Next
        </button>
        
        <button @click="submit()" 
                x-show="currentStep === steps.length - 1">
            Submit
        </button>
    </div>
</div>

<script>
function formWizard() {
    return {
        currentStep: 0,
        steps: [
            { title: 'Customer', completed: false },
            { title: 'Items', completed: false },
            { title: 'Review', completed: false }
        ],
        canProceed() {
            // Validation logic
            return this.steps[this.currentStep].completed;
        },
        next() {
            if (this.canProceed()) {
                this.currentStep++;
            }
        },
        prev() {
            this.currentStep--;
        },
        submit() {
            // Submit form
            htmx.ajax('POST', '/invoices', {
                source: '#invoice-form'
            });
        }
    }
}
</script>
```

### Pattern 3: Real-time Updates with SSE

```html
<!-- Server-Sent Events for live updates -->
<div hx-ext="sse" 
     sse-connect="/stream/dashboard"
     sse-swap="message">
    
    <!-- Content updated via SSE -->
    <div id="stats">
        @StatsDisplay(stats)
    </div>
</div>
```

```go
// Server-side SSE handler
func StreamDashboard(c *fiber.Ctx) error {
    c.Set("Content-Type", "text/event-stream")
    c.Set("Cache-Control", "no-cache")
    c.Set("Connection", "keep-alive")
    
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            stats := fetchLatestStats()
            html := renderStats(stats)
            
            fmt.Fprintf(c, "event: message\n")
            fmt.Fprintf(c, "data: %s\n\n", html)
            c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
                w.Flush()
            }))
            
        case <-c.Context().Done():
            return nil
        }
    }
}
```

### Pattern 4: Optimistic UI Updates

```html
<form hx-post="/comments"
      hx-target="#comments-list"
      hx-swap="afterbegin"
      hx-on::before-request="addOptimisticComment(event)">
    
    <textarea name="content"></textarea>
    <button type="submit">Post Comment</button>
</form>

<div id="comments-list">
    <!-- Comments -->
</div>

<script>
function addOptimisticComment(event) {
    const form = event.detail.elt;
    const content = form.querySelector('textarea').value;
    
    // Add temporary comment
    const tempComment = createCommentHTML(content, {
        temporary: true,
        id: 'temp-' + Date.now()
    });
    
    document.getElementById('comments-list')
        .insertAdjacentHTML('afterbegin', tempComment);
    
    // Clear form
    form.reset();
}

// HTMX will replace the temp comment with real one from server
</script>
```

---

## Summary

This complete component reference provides:

✅ **24 Atoms**: All foundational UI elements  
✅ **20 Molecules**: Essential 2-5 atom combinations  
✅ **18 Organisms**: Complex, self-contained components  
✅ **12 Templates**: All major page layout patterns  
✅ **15 Pages**: Complete page examples for core features  

✅ **Features**:
- Full HTMX integration patterns
- Alpine.js state management examples
- Comprehensive accessibility guidelines (WCAG 2.1 AA)
- Mobile-first responsive patterns
- Real-time update strategies (SSE, WebSocket)
- Optimistic UI patterns
- Schema-driven rendering system
- Multi-tenant theming support

✅ **Production Ready**:
- Type-safe Go structs
- Token-based styling (no hard-coded values)
- Clear dependency hierarchy
- Reusability metrics
- Performance considerations
- Security best practices

**Implementation Priority**:
1. **Phase 1**: Core Atoms (Buttons, Inputs, Icons) - Week 1-2
2. **Phase 2**: Essential Molecules (FormField, SearchBar, Pagination) - Week 3-4
3. **Phase 3**: Key Organisms (DataTable, Forms, Navigation) - Week 5-7
4. **Phase 4**: Main Templates (Dashboard, Form, MasterDetail) - Week 8-9
5. **Phase 5**: Priority Pages (Login, Dashboard, Invoice List) - Week 10-12

---

**Version**: 2.0  
**Maintained By**: Awo ERP Development Team  
**Last Updated**: November 2025  
**Next Review**: December 2025
