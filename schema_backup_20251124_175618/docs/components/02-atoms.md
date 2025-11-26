# Atoms

**Basic building blocks - indivisible UI elements with single responsibilities**

## Definition

Atoms are the foundational UI elements that cannot be broken down further without losing their meaning or function. They serve as the building blocks for all other components in the system.

## Key Characteristics

- **Zero business logic** - purely presentational
- **Single responsibility** - does one thing well
- **Highly reusable** - used 50+ times across the application
- **Design token driven** - all styling comes from schema tokens
- **Stateless** - no internal state management
- **Self-contained** - works in complete isolation

## Core Atom Categories

### 1. Interactive Elements

#### Button Component

```go
// Button variants for different visual styles
type ButtonVariant string

const (
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonDestructive ButtonVariant = "destructive"
    ButtonOutline     ButtonVariant = "outline"
    ButtonGhost       ButtonVariant = "ghost"
    ButtonLink        ButtonVariant = "link"
)

// Button sizes using schema spacing tokens
type ButtonSize string

const (
    ButtonSizeXS ButtonSize = "xs"
    ButtonSizeSM ButtonSize = "sm"
    ButtonSizeMD ButtonSize = "md"
    ButtonSizeLG ButtonSize = "lg"
    ButtonSizeXL ButtonSize = "xl"
)

// Button properties - purely presentational
type ButtonProps struct {
    Variant  ButtonVariant
    Size     ButtonSize
    Type     string // "button", "submit", "reset"
    Disabled bool
    Loading  bool
    Icon     string
    IconLeft string
    IconRight string
    Class    string // Additional CSS classes
    ID       string
    Name     string
    Value    string
    // HTMX integration
    HXPost   string
    HXGet    string
    HXTarget string
    HXSwap   string
    // Alpine.js integration
    AlpineClick string
}
```

**Implementation with Schema Tokens**:

```go
// buttonClasses generates CSS classes using design tokens
func buttonClasses(props ButtonProps) string {
    var classes []string
    
    // Base classes using schema tokens
    classes = append(classes,
        "inline-flex",
        "items-center",
        "justify-center",
        "whitespace-nowrap",
        "rounded-[var(--radius-md)]",
        "text-[var(--font-size-sm)]",
        "font-medium",
        "ring-offset-[var(--background)]",
        "transition-colors",
        "focus-visible:outline-none",
        "focus-visible:ring-2",
        "focus-visible:ring-[var(--ring)]",
        "focus-visible:ring-offset-2",
        "disabled:pointer-events-none",
        "disabled:opacity-50",
    )
    
    // Variant-specific styling using semantic tokens
    switch props.Variant {
    case ButtonPrimary:
        classes = append(classes,
            "bg-[var(--primary)]",
            "text-[var(--primary-foreground)]",
            "hover:bg-[var(--primary)]/90",
        )
    case ButtonSecondary:
        classes = append(classes,
            "bg-[var(--secondary)]",
            "text-[var(--secondary-foreground)]",
            "hover:bg-[var(--secondary)]/80",
        )
    case ButtonDestructive:
        classes = append(classes,
            "bg-[var(--destructive)]",
            "text-[var(--destructive-foreground)]",
            "hover:bg-[var(--destructive)]/90",
        )
    case ButtonOutline:
        classes = append(classes,
            "border",
            "border-[var(--input)]",
            "bg-[var(--background)]",
            "hover:bg-[var(--accent)]",
            "hover:text-[var(--accent-foreground)]",
        )
    case ButtonGhost:
        classes = append(classes,
            "hover:bg-[var(--accent)]",
            "hover:text-[var(--accent-foreground)]",
        )
    case ButtonLink:
        classes = append(classes,
            "text-[var(--primary)]",
            "underline-offset-4",
            "hover:underline",
        )
    }
    
    // Size classes using spacing tokens
    switch props.Size {
    case ButtonSizeXS:
        classes = append(classes, "h-[var(--space-6)]", "px-[var(--space-2)]", "text-xs")
    case ButtonSizeSM:
        classes = append(classes, "h-[var(--space-8)]", "px-[var(--space-3)]", "text-xs")
    case ButtonSizeMD:
        classes = append(classes, "h-[var(--space-10)]", "px-[var(--space-4)]", "py-[var(--space-2)]")
    case ButtonSizeLG:
        classes = append(classes, "h-[var(--space-11)]", "px-[var(--space-8)]")
    case ButtonSizeXL:
        classes = append(classes, "h-[var(--space-14)]", "px-[var(--space-8)]", "text-lg")
    }
    
    if props.Loading {
        classes = append(classes, "pointer-events-none")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}
```

**Templ Component**:

```go
// Button renders an accessible, schema-token styled button
templ Button(props ButtonProps, children ...templ.Component) {
    <button
        type={ props.Type }
        if props.ID != "" {
            id={ props.ID }
        }
        if props.Name != "" {
            name={ props.Name }
        }
        if props.Value != "" {
            value={ props.Value }
        }
        if props.Disabled {
            disabled
        }
        class={ buttonClasses(props) }
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
        if props.Loading {
            aria-label="Loading..."
        }
    >
        // Loading state
        if props.Loading {
            @LoadingIcon()
        }
        
        // Left icon
        if !props.Loading && props.IconLeft != "" {
            @Icon(IconProps{Name: props.IconLeft, Size: IconSizeSM, Class: "mr-2"})
        }
        
        // Single icon (for icon-only buttons)
        if !props.Loading && props.Icon != "" && len(children) == 0 {
            @Icon(IconProps{Name: props.Icon, Size: IconSizeMD})
        }
        
        // Button content
        for _, child := range children {
            @child
        }
        
        // Right icon
        if !props.Loading && props.IconRight != "" {
            @Icon(IconProps{Name: props.IconRight, Size: IconSizeSM, Class: "ml-2"})
        }
    </button>
}
```

### 2. Form Controls

#### Input Component

```go
type InputType string

const (
    InputText     InputType = "text"
    InputEmail    InputType = "email"
    InputPassword InputType = "password"
    InputNumber   InputType = "number"
    InputDate     InputType = "date"
    InputSearch   InputType = "search"
    InputTel      InputType = "tel"
    InputURL      InputType = "url"
)

type InputSize string

const (
    InputSizeSM InputSize = "sm"
    InputSizeMD InputSize = "md"
    InputSizeLG InputSize = "lg"
)

type InputProps struct {
    Type         InputType
    Name         string
    Value        string
    Placeholder  string
    Disabled     bool
    Readonly     bool
    Required     bool
    AutoFocus    bool
    AutoComplete string
    MinLength    int
    MaxLength    int
    Size         InputSize
    Error        bool // Visual error state only
    Class        string
    ID           string
    // HTMX integration
    HXPost    string
    HXTrigger string
    HXTarget  string
    // Alpine.js integration
    AlpineModel string
}
```

**Implementation**:

```go
func inputClasses(props InputProps) string {
    var classes []string
    
    // Base input styling using design tokens
    classes = append(classes,
        "flex",
        "w-full",
        "rounded-[var(--radius-md)]",
        "border",
        "bg-[var(--input)]",
        "px-[var(--space-3)]",
        "text-[var(--font-size-sm)]",
        "ring-offset-[var(--background)]",
        "file:border-0",
        "file:bg-transparent",
        "file:text-[var(--font-size-sm)]",
        "file:font-medium",
        "placeholder:text-[var(--muted-foreground)]",
        "focus-visible:outline-none",
        "focus-visible:ring-2",
        "focus-visible:ring-[var(--ring)]",
        "focus-visible:ring-offset-2",
        "disabled:cursor-not-allowed",
        "disabled:opacity-50",
    )
    
    // Error state styling
    if props.Error {
        classes = append(classes,
            "border-[var(--destructive)]",
            "focus-visible:ring-[var(--destructive)]",
        )
    } else {
        classes = append(classes, "border-[var(--input)]")
    }
    
    // Size-specific styling
    switch props.Size {
    case InputSizeSM:
        classes = append(classes, "h-[var(--space-8)]", "px-[var(--space-2)]", "text-xs")
    case InputSizeLG:
        classes = append(classes, "h-[var(--space-12)]", "px-[var(--space-4)]", "text-[var(--font-size-base)]")
    default: // MD
        classes = append(classes, "h-[var(--space-10)]", "px-[var(--space-3)]")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Input(props InputProps) {
    <input
        type={ string(props.Type) }
        if props.ID != "" {
            id={ props.ID }
        }
        if props.Name != "" {
            name={ props.Name }
        }
        if props.Value != "" {
            value={ props.Value }
        }
        if props.Placeholder != "" {
            placeholder={ props.Placeholder }
        }
        if props.Disabled {
            disabled
        }
        if props.Readonly {
            readonly
        }
        if props.Required {
            required
        }
        if props.AutoFocus {
            autofocus
        }
        if props.AutoComplete != "" {
            autocomplete={ props.AutoComplete }
        }
        if props.MinLength > 0 {
            minlength={ strconv.Itoa(props.MinLength) }
        }
        if props.MaxLength > 0 {
            maxlength={ strconv.Itoa(props.MaxLength) }
        }
        class={ inputClasses(props) }
        if props.HXPost != "" {
            hx-post={ props.HXPost }
        }
        if props.HXTrigger != "" {
            hx-trigger={ props.HXTrigger }
        }
        if props.HXTarget != "" {
            hx-target={ props.HXTarget }
        }
        if props.AlpineModel != "" {
            x-model={ props.AlpineModel }
        }
    />
}
```

### 3. Display Elements

#### Icon Component

```go
type IconSize string

const (
    IconSizeXS IconSize = "xs" // 12px
    IconSizeSM IconSize = "sm" // 16px
    IconSizeMD IconSize = "md" // 20px
    IconSizeLG IconSize = "lg" // 24px
    IconSizeXL IconSize = "xl" // 32px
)

type IconProps struct {
    Name  string   // Icon identifier
    Size  IconSize
    Color string   // CSS class or token reference
    Class string   // Additional classes
}
```

**Implementation**:

```go
func iconClasses(props IconProps) string {
    var classes []string
    
    // Size classes using design tokens
    switch props.Size {
    case IconSizeXS:
        classes = append(classes, "w-[var(--space-3)]", "h-[var(--space-3)]")
    case IconSizeSM:
        classes = append(classes, "w-[var(--space-4)]", "h-[var(--space-4)]")
    case IconSizeMD:
        classes = append(classes, "w-[var(--space-5)]", "h-[var(--space-5)]")
    case IconSizeLG:
        classes = append(classes, "w-[var(--space-6)]", "h-[var(--space-6)]")
    case IconSizeXL:
        classes = append(classes, "w-[var(--space-8)]", "h-[var(--space-8)]")
    }
    
    // Color classes
    if props.Color != "" {
        classes = append(classes, props.Color)
    } else {
        classes = append(classes, "text-current")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Icon(props IconProps) {
    // SVG icons from icon system
    switch props.Name {
    case "user":
        <svg class={ iconClasses(props) } fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
        </svg>
    case "search":
        <svg class={ iconClasses(props) } fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
        </svg>
    case "spinner":
        <svg class={ iconClasses(props) + " animate-spin" } fill="none" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
        </svg>
    default:
        // Fallback or dynamic icon loading
        <svg class={ iconClasses(props) } fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
        </svg>
    }
}
```

#### Badge Component

```go
type BadgeVariant string

const (
    BadgeDefault     BadgeVariant = "default"
    BadgePrimary     BadgeVariant = "primary"
    BadgeSecondary   BadgeVariant = "secondary"
    BadgeDestructive BadgeVariant = "destructive"
    BadgeOutline     BadgeVariant = "outline"
    BadgeSuccess     BadgeVariant = "success"
    BadgeWarning     BadgeVariant = "warning"
)

type BadgeSize string

const (
    BadgeSizeSM BadgeSize = "sm"
    BadgeSizeMD BadgeSize = "md"
    BadgeSizeLG BadgeSize = "lg"
)

type BadgeProps struct {
    Variant BadgeVariant
    Size    BadgeSize
    Class   string
}
```

**Implementation**:

```go
func badgeClasses(props BadgeProps) string {
    var classes []string
    
    // Base styling using design tokens
    classes = append(classes,
        "inline-flex",
        "items-center",
        "rounded-[var(--radius-md)]",
        "border",
        "font-semibold",
        "transition-colors",
        "focus:outline-none",
        "focus:ring-2",
        "focus:ring-[var(--ring)]",
        "focus:ring-offset-2",
    )
    
    // Variant styling
    switch props.Variant {
    case BadgePrimary:
        classes = append(classes,
            "border-transparent",
            "bg-[var(--primary)]",
            "text-[var(--primary-foreground)]",
        )
    case BadgeSecondary:
        classes = append(classes,
            "border-transparent",
            "bg-[var(--secondary)]",
            "text-[var(--secondary-foreground)]",
        )
    case BadgeDestructive:
        classes = append(classes,
            "border-transparent",
            "bg-[var(--destructive)]",
            "text-[var(--destructive-foreground)]",
        )
    case BadgeSuccess:
        classes = append(classes,
            "border-transparent",
            "bg-[var(--success)]",
            "text-[var(--success-foreground)]",
        )
    case BadgeWarning:
        classes = append(classes,
            "border-transparent",
            "bg-[var(--warning)]",
            "text-[var(--warning-foreground)]",
        )
    case BadgeOutline:
        classes = append(classes,
            "border-[var(--border)]",
            "text-[var(--foreground)]",
        )
    default: // Default
        classes = append(classes,
            "border-transparent",
            "bg-[var(--muted)]",
            "text-[var(--muted-foreground)]",
        )
    }
    
    // Size styling
    switch props.Size {
    case BadgeSizeSM:
        classes = append(classes, "px-[var(--space-1-5)]", "py-0", "text-xs")
    case BadgeSizeLG:
        classes = append(classes, "px-[var(--space-3)]", "py-[var(--space-1)]", "text-[var(--font-size-base)]")
    default: // MD
        classes = append(classes, "px-[var(--space-2)]", "py-0", "text-[var(--font-size-sm)]")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Badge(props BadgeProps, children ...templ.Component) {
    <div class={ badgeClasses(props) }>
        for _, child := range children {
            @child
        }
    </div>
}
```

### 4. Typography Elements

#### Label Component

```go
type LabelProps struct {
    For      string // HTML for attribute
    Required bool   // Show required indicator
    Class    string
}

func labelClasses(props LabelProps) string {
    var classes []string
    
    classes = append(classes,
        "text-[var(--font-size-sm)]",
        "font-medium",
        "leading-none",
        "peer-disabled:cursor-not-allowed",
        "peer-disabled:opacity-70",
    )
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Label(props LabelProps, children ...templ.Component) {
    <label
        if props.For != "" {
            for={ props.For }
        }
        class={ labelClasses(props) }
    >
        for _, child := range children {
            @child
        }
        if props.Required {
            <span class="text-[var(--destructive)] ml-1">*</span>
        }
    </label>
}
```

### 5. Layout Helpers

#### Divider Component

```go
type DividerOrientation string

const (
    DividerHorizontal DividerOrientation = "horizontal"
    DividerVertical   DividerOrientation = "vertical"
)

type DividerProps struct {
    Orientation DividerOrientation
    Class       string
}

func dividerClasses(props DividerProps) string {
    var classes []string
    
    classes = append(classes, "border-[var(--border)]")
    
    switch props.Orientation {
    case DividerVertical:
        classes = append(classes, "border-l", "h-full")
    default: // Horizontal
        classes = append(classes, "border-b", "w-full")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Divider(props DividerProps) {
    <hr class={ dividerClasses(props) } />
}
```

## Atom Design Principles

### 1. Token-First Approach

**✅ Correct - Using Design Tokens**:
```go
"bg-[var(--primary)]"           // Semantic token
"px-[var(--space-3)]"           // Spacing token  
"text-[var(--font-size-sm)]"    // Typography token
"rounded-[var(--radius-md)]"    // Border radius token
```

**❌ Incorrect - Hardcoded Values**:
```go
"bg-blue-500"    // Hardcoded color
"px-3"           // Hardcoded spacing
"text-sm"        // Hardcoded text size
"rounded-md"     // Hardcoded radius
```

### 2. Stateless Design

Atoms should have no internal state:

**✅ Correct**:
```go
type ButtonProps struct {
    Disabled bool // State passed as prop
    Loading  bool // Visual state from parent
}
```

**❌ Incorrect**:
```go
type Button struct {
    isClicked    bool // Internal state
    clickCount   int  // Business logic
    lastClicked  time.Time // Complex state
}
```

### 3. Single Responsibility

Each atom has one clear purpose:

**✅ Correct**:
```go
templ Button() {} // Renders clickable button
templ Input() {}  // Renders text input
templ Icon() {}   // Renders SVG icon
```

**❌ Incorrect**:
```go
templ SuperWidget() {
    // Renders button + input + validation + formatting
    // Too many responsibilities!
}
```

### 4. Framework Integration

#### HTMX Integration

```go
// Atoms provide HTMX attribute support
templ Button(props ButtonProps) {
    <button
        if props.HXPost != "" {
            hx-post={ props.HXPost }
            hx-target={ props.HXTarget }
            hx-swap={ props.HXSwap }
        }
    >
        { children... }
    </button>
}
```

#### Alpine.js Integration

```go
// Atoms support Alpine.js directives  
templ Input(props InputProps) {
    <input
        if props.AlpineModel != "" {
            x-model={ props.AlpineModel }
        }
        x-on:blur="touched = true"
        x-bind:disabled="disabled"
    />
}
```

## Common Atom Patterns

### Convenience Components

```go
// Specific variants of base atoms
templ PrimaryButton(props ButtonProps, children ...templ.Component) {
    @Button(ButtonProps{
        Variant: ButtonPrimary,
        Size:    props.Size,
        Type:    props.Type,
        // ... copy other props
    }, children...)
}

templ EmailInput(props InputProps) {
    @Input(InputProps{
        Type:         InputEmail,
        AutoComplete: "email",
        Name:         props.Name,
        // ... copy other props
    })
}
```

### Loading States

```go
templ LoadingSpinner(props IconProps) {
    @Icon(IconProps{
        Name:  "spinner",
        Size:  props.Size,
        Class: props.Class + " animate-spin",
    })
}

templ LoadingButton(props ButtonProps, children ...templ.Component) {
    @Button(ButtonProps{
        Loading:  true,
        Disabled: true,
        Variant:  props.Variant,
        Size:     props.Size,
    }, children...)
}
```

## Testing Atoms

### Unit Tests

```go
func TestButton_VariantClasses(t *testing.T) {
    tests := []struct {
        variant  ButtonVariant
        expected string
    }{
        {ButtonPrimary, "bg-[var(--primary)]"},
        {ButtonSecondary, "bg-[var(--secondary)]"},
        {ButtonDestructive, "bg-[var(--destructive)]"},
    }
    
    for _, tt := range tests {
        t.Run(string(tt.variant), func(t *testing.T) {
            classes := buttonClasses(ButtonProps{Variant: tt.variant})
            assert.Contains(t, classes, tt.expected)
        })
    }
}

func TestInput_ErrorState(t *testing.T) {
    props := InputProps{Error: true}
    classes := inputClasses(props)
    
    assert.Contains(t, classes, "border-[var(--destructive)]")
    assert.Contains(t, classes, "focus-visible:ring-[var(--destructive)]")
}
```

### Visual Regression Tests

```go
// Storybook stories for visual testing
func ButtonStories() {
    // Test all variants and sizes
    for _, variant := range []ButtonVariant{ButtonPrimary, ButtonSecondary, ButtonDestructive} {
        for _, size := range []ButtonSize{ButtonSizeSM, ButtonSizeMD, ButtonSizeLG} {
            // Generate story
        }
    }
}
```

## Migration Guidelines

### From Hardcoded to Token-Based

1. **Identify hardcoded styles** in existing components
2. **Map to appropriate tokens** from design system
3. **Update component classes** to use token references
4. **Test visual consistency** across themes
5. **Document token usage** in component docs

### Component Extraction

1. **Find repeated patterns** - same elements used 3+ times
2. **Extract as atom** - create reusable component
3. **Replace usage** - update existing code to use new atom
4. **Document component** - add to design system docs
5. **Add tests** - ensure reliability

## Best Practices

### ✅ DO

- Use design tokens for all styling
- Keep components stateless and pure
- Follow single responsibility principle
- Support accessibility attributes
- Integrate with HTMX and Alpine.js
- Write comprehensive tests
- Document component APIs

### ❌ DON'T

- Add business logic to atoms
- Use hardcoded colors or spacing
- Create atoms that are too specific
- Implement complex state management
- Ignore accessibility requirements
- Skip testing for foundational components

## Related Documentation

- **[Design Tokens](./01-design-tokens.md)** - Token system and usage
- **[Molecules](./03-molecules.md)** - Next level in atomic design
- **[Schema Integration](./07-integration.md)** - Schema-driven component patterns

---

Atoms form the foundation of our design system. By keeping them simple, reusable, and token-driven, they enable consistent and maintainable UIs that scale across the entire application.