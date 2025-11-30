# Molecules

**Composite components that combine 2-5 atoms to serve a single, focused purpose**

## Definition

Molecules are simple groups of atoms that work together as a unit. They represent the smallest complete UI patterns and introduce basic interactivity while maintaining a single, clear responsibility.

## Key Characteristics

- **Combines 2-5 atoms** into functional groups
- **Single, clear purpose** - can be described in one sentence
- **Reusable patterns** - used across multiple features
- **Simple state management** - local UI state only
- **Self-contained interactions** - no external dependencies
- **Medium complexity** - more than atoms, less than organisms

## Core Molecule Patterns

### 1. Form Field Group

Combines label, input, help text, and error message into a standardized form field pattern.

```go
type FormFieldProps struct {
    // Field identity
    Name        string
    Label       string
    
    // Input configuration
    Type        InputType
    Value       string
    Placeholder string
    
    // Field state
    Required    bool
    Disabled    bool
    Readonly    bool
    Error       bool
    
    // Support content
    HelpText    string
    ErrorText   string
    
    // Styling
    Size        InputSize
    Class       string
    
    // HTMX integration
    HXPost      string
    HXTrigger   string
    HXTarget    string
    
    // Alpine.js integration
    AlpineModel string
}
```

**Implementation**:

```go
func formFieldClasses(props FormFieldProps) string {
    var classes []string
    
    // Base field container
    classes = append(classes, "space-y-[var(--space-2)]")
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ FormField(props FormFieldProps) {
    <div class={ formFieldClasses(props) }>
        // Label with required indicator
        @Label(LabelProps{
            For:      props.Name,
            Required: props.Required,
        }) {
            { props.Label }
        }
        
        // Input with error state
        @Input(InputProps{
            Type:        props.Type,
            Name:        props.Name,
            ID:          props.Name,
            Value:       props.Value,
            Placeholder: props.Placeholder,
            Required:    props.Required,
            Disabled:    props.Disabled,
            Readonly:    props.Readonly,
            Error:       props.Error,
            Size:        props.Size,
            HXPost:      props.HXPost,
            HXTrigger:   props.HXTrigger,
            HXTarget:    props.HXTarget,
            AlpineModel: props.AlpineModel,
        })
        
        // Help text (when no error)
        if props.HelpText != "" && !props.Error {
            <p class="text-[var(--muted-foreground)] text-xs">
                { props.HelpText }
            </p>
        }
        
        // Error message
        if props.Error && props.ErrorText != "" {
            <p class="text-[var(--destructive)] text-xs">
                { props.ErrorText }
            </p>
        }
    </div>
}
```

### 2. Search Bar

Combines input field with search icon and optional clear button.

```go
type SearchBarProps struct {
    // Search configuration
    Name        string
    Value       string
    Placeholder string
    
    // Behavior
    AutoFocus   bool
    ShowClear   bool // Usually derived from Value != ""
    Disabled    bool
    
    // Styling
    Size        InputSize
    Class       string
    
    // HTMX integration
    HXPost      string
    HXTrigger   string
    HXTarget    string
    
    // Alpine.js integration
    AlpineModel string
}
```

**Implementation**:

```go
func searchBarClasses(props SearchBarProps) string {
    var classes []string
    
    classes = append(classes, "relative", "flex", "items-center")
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ SearchBar(props SearchBarProps) {
    <div class={ searchBarClasses(props) }>
        // Search icon
        <div class="absolute left-[var(--space-3)] flex items-center pointer-events-none">
            @Icon(IconProps{
                Name:  "search",
                Size:  IconSizeSM,
                Class: "text-[var(--muted-foreground)]",
            })
        </div>
        
        // Search input with left padding for icon
        @Input(InputProps{
            Type:        InputSearch,
            Name:        props.Name,
            Value:       props.Value,
            Placeholder: props.Placeholder,
            AutoFocus:   props.AutoFocus,
            Disabled:    props.Disabled,
            Size:        props.Size,
            Class:       "pl-[var(--space-10)]" + 
                        if props.ShowClear { " pr-[var(--space-10)]" } else { "" },
            HXPost:      props.HXPost,
            HXTrigger:   props.HXTrigger,
            HXTarget:    props.HXTarget,
            AlpineModel: props.AlpineModel,
        })
        
        // Clear button (when there's a value)
        if props.ShowClear {
            <div class="absolute right-[var(--space-3)] flex items-center">
                @Button(ButtonProps{
                    Variant:     ButtonGhost,
                    Size:        ButtonSizeSM,
                    Icon:        "x",
                    Type:        "button",
                    AlpineClick: "value = ''; $refs.searchInput.focus()",
                    Class:       "h-auto p-1",
                })
            </div>
        }
    </div>
}
```

### 3. Quantity Selector

Number input with increment/decrement buttons for quantity selection.

```go
type QuantitySelectorProps struct {
    // Value configuration
    Name     string
    Value    int
    Min      int
    Max      int
    Step     int
    
    // Behavior
    Disabled bool
    
    // Styling
    Size     InputSize
    Class    string
    
    // HTMX integration
    HXPost   string
    HXTrigger string
    HXTarget string
    
    // Alpine.js integration
    AlpineModel string
}
```

**Implementation**:

```go
func quantitySelectorClasses(props QuantitySelectorProps) string {
    var classes []string
    
    classes = append(classes, "flex", "items-center", "border", "border-[var(--input)]", "rounded-[var(--radius-md)]")
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ QuantitySelector(props QuantitySelectorProps) {
    <div 
        class={ quantitySelectorClasses(props) }
        x-data={ fmt.Sprintf("{ value: %d, min: %d, max: %d, step: %d }", 
                props.Value, props.Min, props.Max, props.Step) }
    >
        // Decrement button
        @Button(ButtonProps{
            Variant:     ButtonGhost,
            Size:        ButtonSizeSM,
            Icon:        "minus",
            Type:        "button",
            Disabled:    props.Disabled,
            AlpineClick: "value = Math.max(min, value - step)",
            Class:       "h-full rounded-r-none border-r",
        })
        
        // Number input
        @Input(InputProps{
            Type:     InputNumber,
            Name:     props.Name,
            Value:    strconv.Itoa(props.Value),
            Disabled: props.Disabled,
            Size:     props.Size,
            Class:    "text-center border-none rounded-none focus:ring-0 focus:ring-offset-0",
            AlpineModel: if props.AlpineModel != "" { props.AlpineModel } else { "value" },
            HXPost:   props.HXPost,
            HXTrigger: props.HXTrigger,
            HXTarget: props.HXTarget,
        })
        
        // Increment button
        @Button(ButtonProps{
            Variant:     ButtonGhost,
            Size:        ButtonSizeSM,
            Icon:        "plus",
            Type:        "button",
            Disabled:    props.Disabled,
            AlpineClick: "value = Math.min(max, value + step)",
            Class:       "h-full rounded-l-none border-l",
        })
    </div>
}
```

### 4. Stat Display

Shows a metric with label, value, and optional trend indicator.

```go
type TrendDirection string

const (
    TrendUp   TrendDirection = "up"
    TrendDown TrendDirection = "down"
    TrendFlat TrendDirection = "flat"
)

type StatDisplayProps struct {
    // Content
    Label       string
    Value       string // Pre-formatted value
    SubValue    string // Optional secondary value
    
    // Visual elements
    Icon        string // Optional icon
    Trend       TrendDirection
    TrendValue  string // e.g., "+12%", "-5.2%"
    
    // Styling
    Variant     string // "default", "success", "warning", "danger"
    Size        string // "sm", "md", "lg"
    Class       string
}
```

**Implementation**:

```go
func statDisplayClasses(props StatDisplayProps) string {
    var classes []string
    
    classes = append(classes,
        "flex",
        "flex-col",
        "space-y-[var(--space-1)]",
        "p-[var(--space-4)]",
        "bg-[var(--card)]",
        "border",
        "border-[var(--border)]",
        "rounded-[var(--radius-lg)]",
    )
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

func getTrendIcon(direction TrendDirection) string {
    switch direction {
    case TrendUp:
        return "trending-up"
    case TrendDown:
        return "trending-down"
    default:
        return "minus"
    }
}

func getTrendColor(direction TrendDirection) string {
    switch direction {
    case TrendUp:
        return "text-[var(--success)]"
    case TrendDown:
        return "text-[var(--destructive)]"
    default:
        return "text-[var(--muted-foreground)]"
    }
}

templ StatDisplay(props StatDisplayProps) {
    <div class={ statDisplayClasses(props) }>
        // Header with label and optional icon
        <div class="flex items-center justify-between">
            <div class="flex items-center space-x-[var(--space-2)]">
                if props.Icon != "" {
                    @Icon(IconProps{
                        Name:  props.Icon,
                        Size:  IconSizeSM,
                        Class: "text-[var(--muted-foreground)]",
                    })
                }
                <span class="text-[var(--muted-foreground)] text-[var(--font-size-sm)] font-medium">
                    { props.Label }
                </span>
            </div>
            
            // Trend indicator
            if props.Trend != TrendFlat && props.TrendValue != "" {
                <div class={ "flex items-center space-x-1 " + getTrendColor(props.Trend) }>
                    @Icon(IconProps{
                        Name: getTrendIcon(props.Trend),
                        Size: IconSizeXS,
                    })
                    <span class="text-xs font-medium">
                        { props.TrendValue }
                    </span>
                </div>
            }
        </div>
        
        // Main value
        <div class="flex items-baseline space-x-[var(--space-2)]">
            <span class="text-2xl font-bold text-[var(--foreground)]">
                { props.Value }
            </span>
            if props.SubValue != "" {
                <span class="text-[var(--muted-foreground)] text-[var(--font-size-sm)]">
                    { props.SubValue }
                </span>
            }
        </div>
    </div>
}
```

### 5. Media Object

Combines image/icon with title and description for content display.

```go
type MediaObjectProps struct {
    // Media content
    ImageURL    string
    ImageAlt    string
    Icon        string // Alternative to image
    
    // Text content
    Title       string
    Description string
    Meta        string // Optional metadata (date, author, etc.)
    
    // Configuration
    ImageSize   string // "sm", "md", "lg"
    Alignment   string // "start", "center"
    Clickable   bool
    
    // Links
    HRef        string
    HXGet       string
    HXTarget    string
    
    // Styling
    Class       string
}
```

**Implementation**:

```go
func mediaObjectClasses(props MediaObjectProps) string {
    var classes []string
    
    classes = append(classes, "flex", "space-x-[var(--space-4)]")
    
    if props.Alignment == "center" {
        classes = append(classes, "items-center")
    } else {
        classes = append(classes, "items-start")
    }
    
    if props.Clickable {
        classes = append(classes, 
            "cursor-pointer", 
            "hover:bg-[var(--muted)]/30", 
            "rounded-[var(--radius-md)]",
            "p-[var(--space-3)]",
            "-m-[var(--space-3)]",
        )
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

func getImageSize(size string) (string, string) {
    switch size {
    case "sm":
        return "w-[var(--space-8)]", "h-[var(--space-8)]"
    case "lg":
        return "w-[var(--space-16)]", "h-[var(--space-16)]"
    default: // md
        return "w-[var(--space-12)]", "h-[var(--space-12)]"
    }
}

templ MediaObject(props MediaObjectProps) {
    if props.Clickable && props.HRef != "" {
        <a href={ props.HRef } class={ mediaObjectClasses(props) }>
            @mediaObjectContent(props)
        </a>
    } else if props.Clickable && props.HXGet != "" {
        <div 
            class={ mediaObjectClasses(props) }
            hx-get={ props.HXGet }
            hx-target={ props.HXTarget }
        >
            @mediaObjectContent(props)
        </div>
    } else {
        <div class={ mediaObjectClasses(props) }>
            @mediaObjectContent(props)
        </div>
    }
}

templ mediaObjectContent(props MediaObjectProps) {
    // Media element
    <div class="flex-shrink-0">
        if props.ImageURL != "" {
            <img 
                src={ props.ImageURL }
                alt={ props.ImageAlt }
                class={ fmt.Sprintf("%s %s rounded-[var(--radius-md)] object-cover", getImageSize(props.ImageSize)) }
                loading="lazy"
            />
        } else if props.Icon != "" {
            <div class={ fmt.Sprintf("flex items-center justify-center %s %s bg-[var(--muted)] rounded-[var(--radius-md)]", getImageSize(props.ImageSize)) }>
                @Icon(IconProps{
                    Name: props.Icon,
                    Size: IconSizeLG,
                    Class: "text-[var(--muted-foreground)]",
                })
            </div>
        }
    </div>
    
    // Content
    <div class="flex-1 min-w-0">
        <h4 class="text-[var(--font-size-base)] font-semibold text-[var(--foreground)] truncate">
            { props.Title }
        </h4>
        if props.Description != "" {
            <p class="text-[var(--font-size-sm)] text-[var(--muted-foreground)] line-clamp-2">
                { props.Description }
            </p>
        }
        if props.Meta != "" {
            <p class="text-xs text-[var(--muted-foreground)] mt-[var(--space-1)]">
                { props.Meta }
            </p>
        }
    </div>
}
```

### 6. Tag Input

Displays selected tags with remove functionality.

```go
type TagItem struct {
    ID    string
    Label string
    Color string // Optional color variant
}

type TagInputProps struct {
    // Data
    Tags []TagItem
    Name string // Form field name
    
    // Configuration
    Removable  bool
    MaxDisplay int // Show "N more" after this many tags
    
    // Styling
    Size  string // "sm", "md", "lg"  
    Class string
    
    // HTMX integration
    HXPost   string
    HXTarget string
}
```

**Implementation**:

```go
func tagInputClasses(props TagInputProps) string {
    var classes []string
    
    classes = append(classes, "flex", "flex-wrap", "gap-[var(--space-2)]")
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ TagInput(props TagInputProps) {
    <div class={ tagInputClasses(props) }>
        // Display tags (up to MaxDisplay)
        for i, tag := range props.Tags {
            if props.MaxDisplay == 0 || i < props.MaxDisplay {
                @tagItem(tag, props)
            }
        }
        
        // "N more" indicator
        if props.MaxDisplay > 0 && len(props.Tags) > props.MaxDisplay {
            <span class="text-[var(--muted-foreground)] text-[var(--font-size-sm)] self-center">
                +{ strconv.Itoa(len(props.Tags) - props.MaxDisplay) } more
            </span>
        }
        
        // Hidden inputs for form submission
        for _, tag := range props.Tags {
            <input type="hidden" name={ props.Name } value={ tag.ID } />
        }
    </div>
}

templ tagItem(tag TagItem, props TagInputProps) {
    <div class="flex items-center space-x-1 bg-[var(--muted)] text-[var(--muted-foreground)] px-[var(--space-2)] py-1 rounded-[var(--radius-sm)] text-[var(--font-size-sm)]">
        <span>{ tag.Label }</span>
        if props.Removable {
            @Button(ButtonProps{
                Variant:     ButtonGhost,
                Size:        ButtonSizeXS,
                Icon:        "x",
                Type:        "button",
                AlpineClick: fmt.Sprintf("removeTag('%s')", tag.ID),
                Class:       "h-auto p-0 ml-1",
                HXPost:      if props.HXPost != "" { fmt.Sprintf("%s?remove=%s", props.HXPost, tag.ID) } else { "" },
                HXTarget:    props.HXTarget,
            })
        }
    </div>
}
```

## Molecule Design Principles

### 1. Single Purpose Rule

Each molecule should have one clear, easily describable purpose:

**✅ Good Examples**:
- FormField: "Renders a labeled input with help text and error display"
- SearchBar: "Provides text input with search icon and clear button"
- StatDisplay: "Shows a metric with label, value, and trend indicator"

**❌ Bad Examples**:
- UserProfileEditor: "Handles user profile editing, avatar upload, and preferences"
- CompleteFormBuilder: "Builds any type of form with validation and submission"

### 2. Minimal State Management

Molecules can have simple, local UI state:

**✅ Acceptable State**:
```go
// Dropdown open/closed
x-data="{ open: false }"

// Input current value
x-data="{ value: @js(props.Value) }"

// Show/hide toggle
x-data="{ visible: true }"
```

**❌ Unacceptable State**:
```go
// Complex business state
x-data="{ user: {}, orders: [], currentStep: 1 }"

// API call management
x-data="{ loading: false, error: null, retryCount: 0 }"
```

### 3. Composition Over Configuration

Build molecules through composition rather than complex configuration:

**✅ Good Approach**:
```go
// Separate molecules for different purposes
templ EmailFormField(props FormFieldProps)
templ PasswordFormField(props FormFieldProps)
templ SearchInput(props SearchBarProps)
```

**❌ Bad Approach**:
```go
// Over-configured single component
type MegaInputProps struct {
    Type string
    ShowSearch bool
    ShowPassword bool
    ShowValidation bool
    ShowCounter bool
    ShowIcon bool
    // ... 20 more configuration flags
}
```

### 4. Framework Integration

#### HTMX Integration

```go
// Molecules provide HTMX endpoints for interactions
templ SearchBar(props SearchBarProps) {
    <input 
        hx-post="/api/search"
        hx-trigger="keyup changed delay:300ms"
        hx-target="#search-results"
        hx-swap="innerHTML"
    />
}
```

#### Alpine.js State Management

```go
// Molecules use Alpine.js for local interactions
templ QuantitySelector(props QuantitySelectorProps) {
    <div x-data="{ 
        value: @js(props.Value),
        increment() { this.value = Math.min(@js(props.Max), this.value + @js(props.Step)) },
        decrement() { this.value = Math.max(@js(props.Min), this.value - @js(props.Step)) }
    }">
        // Implementation
    </div>
}
```

## Common Molecule Patterns

### 1. Input Enhancement Patterns

```go
// Password input with visibility toggle
templ PasswordField(props FormFieldProps) {
    <div x-data="{ showPassword: false }">
        @FormField(FormFieldProps{
            Type: if showPassword { InputText } else { InputPassword },
            // ... other props
        })
        @Button(ButtonProps{
            Icon:        if showPassword { "eye-off" } else { "eye" },
            AlpineClick: "showPassword = !showPassword",
        })
    </div>
}

// Currency input with currency selector
templ CurrencyField(props CurrencyFieldProps) {
    <div class="flex">
        @Select(SelectProps{
            Name:    props.CurrencyName,
            Value:   props.Currency,
            Options: props.CurrencyOptions,
            Class:   "rounded-r-none",
        })
        @Input(InputProps{
            Type:  InputNumber,
            Name:  props.AmountName,
            Value: props.Amount,
            Class: "rounded-l-none border-l-0",
        })
    </div>
}
```

### 2. Display Enhancement Patterns

```go
// Avatar with fallback initials
templ AvatarWithFallback(props AvatarProps) {
    <div class="relative">
        if props.ImageURL != "" {
            <img 
                src={ props.ImageURL }
                alt={ props.Alt }
                class="rounded-full"
                x-show="imageLoaded"
                x-on:load="imageLoaded = true"
                x-on:error="imageLoaded = false"
            />
        }
        <div 
            class="flex items-center justify-center bg-[var(--muted)] rounded-full"
            x-show="!imageLoaded"
        >
            <span class="text-[var(--muted-foreground)] font-medium">
                { getInitials(props.Name) }
            </span>
        </div>
    </div>
}

// Badge with count
templ NotificationBadge(props NotificationBadgeProps) {
    <div class="relative">
        @Icon(IconProps{Name: "bell"})
        if props.Count > 0 {
            @Badge(BadgeProps{
                Variant: BadgeDestructive,
                Size:    BadgeSizeSM,
                Class:   "absolute -top-2 -right-2 min-w-[1.5rem] h-6 flex items-center justify-center",
            }) {
                { if props.Count > 99 { "99+" } else { strconv.Itoa(props.Count) } }
            }
        }
    </div>
}
```

### 3. Navigation Patterns

```go
// Breadcrumb with separators
templ Breadcrumb(props BreadcrumbProps) {
    <nav class="flex items-center space-x-[var(--space-1)]">
        for i, item := range props.Items {
            if i > 0 {
                @Icon(IconProps{
                    Name:  "chevron-right",
                    Size:  IconSizeXS,
                    Class: "text-[var(--muted-foreground)]",
                })
            }
            
            if item.Href != "" {
                <a 
                    href={ item.Href }
                    class="text-[var(--muted-foreground)] hover:text-[var(--foreground)] text-[var(--font-size-sm)]"
                >
                    { item.Label }
                </a>
            } else {
                <span class="text-[var(--foreground)] text-[var(--font-size-sm)] font-medium">
                    { item.Label }
                </span>
            }
        }
    </nav>
}

// Tab navigation
templ TabNavigation(props TabNavigationProps) {
    <nav class="flex space-x-[var(--space-8)] border-b border-[var(--border)]">
        for _, tab := range props.Tabs {
            <a
                href={ tab.Href }
                class={ "pb-[var(--space-4)] text-[var(--font-size-sm)] font-medium border-b-2 transition-colors" +
                       if tab.Active { 
                           " border-[var(--primary)] text-[var(--primary)]" 
                       } else { 
                           " border-transparent text-[var(--muted-foreground)] hover:text-[var(--foreground)]" 
                       } }
                hx-get={ tab.HXGet }
                hx-target={ tab.HXTarget }
            >
                { tab.Label }
            </a>
        }
    </nav>
}
```

## Testing Molecules

### Interaction Testing

```go
func TestQuantitySelector_Increment(t *testing.T) {
    props := QuantitySelectorProps{
        Value: 1,
        Min:   0,
        Max:   10,
        Step:  1,
    }
    
    // Test increment button functionality
    html := renderToString(QuantitySelector(props))
    
    // Verify increment button exists
    assert.Contains(t, html, "x-on:click=\"value = Math.min(10, value + 1)\"")
    
    // Verify input reflects current value
    assert.Contains(t, html, "value=\"1\"")
}

func TestSearchBar_ClearButton(t *testing.T) {
    props := SearchBarProps{
        Value:     "test search",
        ShowClear: true,
    }
    
    html := renderToString(SearchBar(props))
    
    // Verify clear button is present when there's a value
    assert.Contains(t, html, "x-on:click=\"value = ''; $refs.searchInput.focus()\"")
    
    // Test without value
    props.Value = ""
    props.ShowClear = false
    html = renderToString(SearchBar(props))
    
    // Verify clear button is not present
    assert.NotContains(t, html, "x-on:click=\"value = ''")
}
```

### Integration Testing

```go
func TestFormField_ErrorStateIntegration(t *testing.T) {
    props := FormFieldProps{
        Name:      "email",
        Label:     "Email Address",
        Type:      InputEmail,
        Error:     true,
        ErrorText: "Please enter a valid email",
    }
    
    html := renderToString(FormField(props))
    
    // Verify label is rendered
    assert.Contains(t, html, "Email Address")
    
    // Verify input has error styling
    assert.Contains(t, html, "border-[var(--destructive)]")
    
    // Verify error message is displayed
    assert.Contains(t, html, "Please enter a valid email")
    
    // Verify help text is hidden when error is shown
    props.HelpText = "Enter your email address"
    html = renderToString(FormField(props))
    assert.NotContains(t, html, "Enter your email address")
}
```

## Best Practices

### ✅ DO

- Keep molecules focused on a single purpose
- Use local state for simple UI interactions
- Compose atoms rather than recreate functionality
- Support both HTMX and Alpine.js integration
- Provide clear, typed props interfaces
- Test interaction patterns thoroughly
- Make molecules reusable across features

### ❌ DON'T

- Add complex business logic to molecules
- Make molecules too feature-specific
- Create molecules with too many configuration options
- Implement data fetching in molecules
- Break the single responsibility principle
- Create circular dependencies between molecules

## Migration Guidelines

### Identifying Molecule Candidates

1. **Find repeated atom patterns** - same 2-5 atoms used together 3+ times
2. **Look for interaction patterns** - atoms that need to work together
3. **Extract form patterns** - label + input + help + error combinations
4. **Identify display patterns** - icon + text combinations

### Extraction Process

1. **Create molecule component** with proper props interface
2. **Replace existing usage** with the new molecule
3. **Add tests** for functionality and interactions
4. **Document component** with examples and props reference
5. **Review for reusability** across different contexts

## Related Documentation

- **[Atoms](./02-atoms.md)** - Building blocks used in molecules
- **[Organisms](./04-organisms.md)** - Next level in atomic design
- **[Design Tokens](./01-design-tokens.md)** - Token system for consistent styling
- **[Integration](./07-integration.md)** - Schema integration patterns

---

Molecules represent the first level of meaningful component composition, providing reusable patterns that form the foundation for more complex organisms and complete user interfaces.