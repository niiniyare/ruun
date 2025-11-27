# Templ Utilities Guide

## Overview
The `pkg/utils` package provides essential helper functions for working with templ components, focusing on class management, conditional rendering, and attribute handling.

## Utility Functions

### 1. TwMerge - Tailwind Class Merging

**Purpose:** Intelligently merges Tailwind classes and resolves conflicts.

**Usage:**
```go
import "project/pkg/utils"

// Basic merging
classes := utils.TwMerge("bg-red-500 text-white", "bg-blue-500 hover:bg-green-500")
// Result: "text-white bg-blue-500 hover:bg-green-500"

// In component class functions
func getButtonClasses(variant string, size string, className string) string {
    base := "inline-flex items-center justify-center font-medium"
    variantClasses := map[string]string{
        "primary": "bg-primary text-primary-foreground",
        "secondary": "bg-secondary text-secondary-foreground",
    }
    sizeClasses := map[string]string{
        "sm": "px-3 py-1.5 text-sm",
        "lg": "px-6 py-3 text-base",
    }
    
    return utils.TwMerge(base, variantClasses[variant], sizeClasses[size], className)
}
```

**Templ Integration:**
```templ
templ Button(props ButtonProps) {
    <button class={utils.TwMerge(
        "btn-base",
        utils.If(props.Primary, "btn-primary"),
        utils.If(props.Disabled, "btn-disabled"),
        props.ClassName,
    )}>
        {props.Label}
    </button>
}
```

### 2. If - Conditional Values

**Purpose:** Returns a value only if condition is true, otherwise returns zero value.

**Usage:**
```go
// Basic conditional classes
classes := utils.If(isActive, "bg-primary text-white")

// Conditional attributes
attrs := templ.Attributes{
    "disabled": utils.If(isDisabled, "disabled"),
    "aria-pressed": utils.If(isPressed, "true"),
}
```

**Templ Examples:**
```templ
templ Alert(variant string, dismissible bool) {
    <div class={utils.TwMerge(
        "alert-base",
        utils.If(variant == "error", "alert-error"),
        utils.If(variant == "success", "alert-success"),
        utils.If(dismissible, "pr-12"),
    )}>
        // Content
    </div>
}

templ Input(props InputProps) {
    <input
        type={props.Type}
        class={utils.TwMerge(
            "input-base",
            utils.If(props.Error, "input-error"),
            utils.If(props.Disabled, "input-disabled"),
        )}
        disabled?={props.Disabled}
        required?={utils.If(props.Required, true)}
    />
}
```

### 3. IfElse - Binary Conditional Values

**Purpose:** Returns one of two values based on condition.

**Usage:**
```go
// Toggle between two states
iconName := utils.IfElse(isExpanded, "chevron-up", "chevron-down")
buttonText := utils.IfElse(isLoading, "Loading...", "Submit")
```

**Templ Examples:**
```templ
templ Toggle(isOn bool) {
    <button class={utils.TwMerge(
        "toggle-base",
        utils.IfElse(isOn, "toggle-on", "toggle-off"),
    )}>
        @Icon(utils.IfElse(isOn, "check", "x"))
    </button>
}

templ Badge(count int) {
    <span class={utils.IfElse(
        count > 0,
        "badge-visible",
        "badge-hidden",
    )}>
        {utils.IfElse(count > 99, "99+", fmt.Sprint(count))}
    </span>
}
```

### 4. MergeAttributes - Attribute Combining

**Purpose:** Combines multiple `templ.Attributes` into a single map.

**Usage:**
```go
baseAttrs := templ.Attributes{"class": "base-class"}
conditionalAttrs := templ.Attributes{}

if props.Disabled {
    conditionalAttrs["disabled"] = "disabled"
    conditionalAttrs["aria-disabled"] = "true"
}

merged := utils.MergeAttributes(baseAttrs, conditionalAttrs, props.Attributes)
```

**Templ Examples:**
```templ
templ Button(props ButtonProps) {
    <button
        {utils.MergeAttributes(
            templ.Attributes{"type": props.Type},
            utils.If(props.Disabled, templ.Attributes{
                "disabled": "disabled",
                "aria-disabled": "true",
            }),
            props.ExtraAttrs,
        )...}
        class={getButtonClasses(props)}
    >
        {props.Label}
    </button>
}
```

### 5. RandomID - Unique Identifiers

**Purpose:** Generates cryptographically secure random IDs for component instances.

**Usage:**
```go
// Generate unique ID for form elements
fieldID := utils.RandomID() // "id-1a2b3c4d"
```

**Templ Examples:**
```templ
templ FormField(props FormFieldProps) {
    { id := utils.IfElse(props.ID != "", props.ID, utils.RandomID()) }
    <div>
        <label for={id}>{props.Label}</label>
        <input id={id} name={props.Name} />
        <div id={id + "-help"}>{props.HelpText}</div>
    </div>
}

templ Accordion(items []AccordionItem) {
    for _, item := range items {
        { panelID := utils.RandomID() }
        <div>
            <button aria-controls={panelID}>{item.Title}</button>
            <div id={panelID}>{item.Content}</div>
        </div>
    }
}
```

### 6. GetStringValue - Safe String Conversion

**Purpose:** Safely converts any value to string representation.

**Usage:**
```go
// Handle dynamic values safely
displayValue := utils.GetStringValue(props.Value) // Works with any type
```

**Templ Examples:**
```templ
templ DataCell(value any) {
    <td>{utils.GetStringValue(value)}</td>
}
```

## Advanced Patterns

### Complex Class Composition
```go
func getCardClasses(props CardProps) string {
    return utils.TwMerge(
        "card-base rounded-lg border p-4",
        utils.If(props.Hoverable, "hover:shadow-md transition-shadow"),
        utils.IfElse(props.Variant == "primary", "border-primary", "border-border"),
        utils.If(props.Selected, "ring-2 ring-primary"),
        props.ClassName,
    )
}
```

### Dynamic Attribute Building
```go
func buildFormAttrs(props FormProps) templ.Attributes {
    attrs := templ.Attributes{}
    
    if props.HxPost != "" {
        attrs["hx-post"] = props.HxPost
        attrs["hx-target"] = props.HxTarget
        attrs["hx-swap"] = utils.IfElse(props.HxSwap != "", props.HxSwap, "innerHTML")
    }
    
    if props.NoValidate {
        attrs["novalidate"] = "novalidate"
    }
    
    return utils.MergeAttributes(attrs, props.ExtraAttrs)
}
```

### Conditional Component Rendering
```templ
templ Card(props CardProps) {
    <div class={getCardClasses(props)}>
        if props.Header != nil {
            <header class="card-header">
                {props.Header...}
            </header>
        }
        
        <div class={utils.TwMerge(
            "card-content",
            utils.If(props.Padded, "p-4"),
        )}>
            {props.Children...}
        </div>
        
        if len(props.Actions) > 0 {
            <footer class="card-footer">
                for _, action := range props.Actions {
                    @Button(ButtonProps{
                        Label: action.Label,
                        Variant: utils.IfElse(action.Primary, ButtonPrimary, ButtonSecondary),
                    })
                }
            </footer>
        }
    </div>
}
```

## Best Practices

1. **Always use TwMerge** when combining multiple class sources
2. **Prefer If/IfElse** over ternary operators for readability
3. **Use RandomID** for accessibility relationships (aria-describedby, for/id)
4. **MergeAttributes** when building complex interactive components
5. **Keep utility usage** in helper functions when logic becomes complex

## Performance Notes

- `TwMerge` has minimal overhead and caches results
- `If/IfElse` are zero-allocation for simple types
- `RandomID` uses crypto/rand for security but generates short IDs for performance
- `MergeAttributes` creates new maps, use sparingly in hot paths