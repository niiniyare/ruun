# Button Component

A versatile, accessible button component with multiple variants, sizes, loading states, and full HTMX and Alpine.js integration. Built with design system tokens for consistent theming.

## Table of Contents

- [Basic Usage](#basic-usage)
- [Variants](#variants)
- [Sizes](#sizes)
- [Icons](#icons)
- [States](#states)
- [Interactive Buttons](#interactive-buttons)
- [API Patterns](#api-patterns)
- [JSON Schema](#json-schema)
- [Examples](#examples)
- [API Reference](#api-reference)

## Basic Usage

### Simple Button

```go
// Basic button with text
@Button(ButtonProps{Text: "Click Me"})

// With children (takes precedence over Text)
@Button(ButtonProps{}) { Click Me }

// Primary button (default)
@Button(ButtonProps{Variant: ButtonPrimary, Text: "Save"})
```

### Using Convenience Components

```go
@PrimaryButton(ButtonProps{Text: "Save"})
@SecondaryButton(ButtonProps{Text: "Cancel"})
@DestructiveButton(ButtonProps{Text: "Delete"})
```

## Variants

Six visual variants following your design system:

```go
// Primary (Default)
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Primary",
})

// Secondary
@Button(ButtonProps{
    Variant: ButtonSecondary,
    Text: "Secondary",
})

// Destructive (Danger)
@Button(ButtonProps{
    Variant: ButtonDestructive,
    Text: "Delete",
})

// Outline
@Button(ButtonProps{
    Variant: ButtonOutline,
    Text: "Outline",
})

// Ghost
@Button(ButtonProps{
    Variant: ButtonGhost,
    Text: "Ghost",
})

// Link
@Button(ButtonProps{
    Variant: ButtonLink,
    Text: "Link Style",
})
```

## Sizes

Five size options:

```go
// Extra Small
@Button(ButtonProps{Size: ButtonSizeXS, Text: "XS"})

// Small
@Button(ButtonProps{Size: ButtonSizeSM, Text: "Small"})

// Medium (default)
@Button(ButtonProps{Size: ButtonSizeMD, Text: "Medium"})

// Large
@Button(ButtonProps{Size: ButtonSizeLG, Text: "Large"})

// Extra Large
@Button(ButtonProps{Size: ButtonSizeXL, Text: "XL"})
```

## Icons

Add icons to buttons for better visual communication:

```go
// Left icon
@Button(ButtonProps{
    IconLeft: "save",
    Text: "Save",
})

// Right icon
@Button(ButtonProps{
    IconRight: "arrow-right",
    Text: "Next",
})

// Both icons
@Button(ButtonProps{
    IconLeft: "download",
    IconRight: "external-link",
    Text: "Export",
})

// Icon only (no text)
@Button(ButtonProps{
    Icon: "trash",
    Variant: ButtonDestructive,
    Size: ButtonSizeSM,
})
```

## States

### Disabled State

```go
@Button(ButtonProps{
    Text: "Submit",
    Disabled: true,
})
```

### Loading State

Shows a spinner and disables interaction:

```go
@Button(ButtonProps{
    Text: "Saving...",
    Loading: true,
})
```

### Button Types

```go
// Default button type
@Button(ButtonProps{
    Text: "Button",
    Type: "button", // default
})

// Submit button (for forms)
@Button(ButtonProps{
    Text: "Submit",
    Type: "submit",
})

// Reset button (for forms)
@Button(ButtonProps{
    Text: "Reset",
    Type: "reset",
})
```

### Form Integration

```go
<form>
    <input type="text" name="email" />
    
    @Button(ButtonProps{
        Text: "Submit Form",
        Type: "submit",
        Name: "action",
        Value: "save",
    })
</form>
```

## Interactive Buttons

### HTMX Integration

```go
// GET request
@Button(ButtonProps{
    Text: "Load More",
    HXGet: "/api/items",
    HXTarget: "#items-list",
    HXSwap: "beforeend",
})

// POST request
@Button(ButtonProps{
    Text: "Delete Item",
    Variant: ButtonDestructive,
    HXPost: "/api/items/123",
    HXTarget: "#item-123",
    HXSwap: "outerHTML",
    HXTrigger: "click",
})

// With loading state
@Button(ButtonProps{
    Text: "Save Changes",
    HXPost: "/api/save",
    HXTarget: "#result",
    IconLeft: "save",
})
```

### Alpine.js Click Handler

```go
@Button(ButtonProps{
    Text: "Toggle Modal",
    AlpineClick: "open = !open",
})

// With Alpine data
<div x-data="{ count: 0 }">
    @Button(ButtonProps{
        AlpineClick: "count++",
    }) {
        <span>Clicked: </span>
        <span x-text="count"></span>
    }
</div>
```

## API Patterns

### 1. Direct Props (Standard)

```go
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size: ButtonSizeLG,
    IconLeft: "save",
    Text: "Save Changes",
    Type: "submit",
}) 
```

### 2. Fluent API (Builder Pattern)

Chain methods for readable, intuitive syntax:

```go
btn := NewButton().
    Primary().
    Large().
    WithIconLeft("save").
    WithText("Save Changes").
    AsSubmit().
    Build()

@Button(btn)
```

**All Builder Methods:**

```go
// Variants
.Primary()
.Secondary()
.Destructive()
.Outline()
.Ghost()
.Link()

// Sizes
.ExtraSmall()
.Small()
.Medium()
.Large()
.ExtraLarge()

// Content
.WithText(text)
.WithIcon(icon)
.WithIconLeft(icon)
.WithIconRight(icon)
.WithClass(class)
.WithID(id)

// Attributes
.AsSubmit()
.AsReset()
.WithName(name)
.WithValue(value)
.Disabled()
.Loading()

// HTMX
.HXPost(url)
.HXGet(url)
.HXTarget(target)
.HXSwap(swap)
.HXTrigger(trigger)

// Alpine.js
.OnClick(handler)

// Build
.Build() // Returns ButtonProps
```

**Complex Example:**

```go
saveButton := NewButton().
    Primary().
    Large().
    WithIconLeft("save").
    WithText("Save Changes").
    AsSubmit().
    HXPost("/api/save").
    HXTarget("#result").
    HXSwap("innerHTML").
    Build()

@Button(saveButton)
```

**Loading Button:**

```go
loadingBtn := NewButton().
    Secondary().
    WithText("Processing...").
    Loading().
    Build()

@Button(loadingBtn)
```

### 3. Functional Options Pattern

Composable and reusable:

```go
@Button(NewButtonWithOptions(
    WithPrimary(),
    WithLarge(),
    WithButtonIconLeft("save"),
    WithButtonText("Save"),
))
```

**Reusable Configurations:**

```go
var (
    SaveButtonOptions = []ButtonOption{
        WithPrimary(),
        WithButtonIconLeft("save"),
        WithButtonText("Save"),
        AsSubmit(),
    }
    
    CancelButtonOptions = []ButtonOption{
        WithSecondary(),
        WithButtonText("Cancel"),
        WithButtonAlpineClick("closeModal()"),
    }
    
    DeleteButtonOptions = []ButtonOption{
        WithDestructive(),
        WithSmall(),
        WithButtonIcon("trash"),
    }
)

// Use predefined options
@Button(NewButtonWithOptions(SaveButtonOptions...))

// Extend with more options
@Button(NewButtonWithOptions(
    append(SaveButtonOptions, AsLoading())...,
))
```

**All Option Functions:**

```go
// Variants
WithVariant(variant), WithPrimary(), WithSecondary(), WithDestructive(),
WithOutline(), WithGhost(), WithLink()

// Sizes
WithSize(size), WithExtraSmall(), WithSmall(), WithMedium(), 
WithLarge(), WithExtraLarge()

// Content
WithButtonText(text), WithButtonIcon(icon), WithButtonIconLeft(icon),
WithButtonIconRight(icon), WithButtonClass(class), WithButtonID(id)

// Attributes
AsSubmit(), AsReset(), WithButtonName(name), WithButtonValue(value),
AsDisabled(), AsLoading()

// HTMX
WithButtonHXPost(url), WithButtonHXGet(url), WithButtonHXTarget(target),
WithButtonHXSwap(swap), WithButtonHXTrigger(trigger)

// Alpine.js
WithButtonAlpineClick(handler)
```

## JSON Schema

Load buttons dynamically from JSON:

### Basic JSON Button

```go
jsonButton := `{
    "text": "Save Changes",
    "variant": "primary",
    "iconLeft": "save"
}`

@ButtonFrom(jsonButton)
```

### JSON Schema Structure

```json
{
    "variant": "primary",           // primary, secondary, destructive, outline, ghost, link
    "size": "md",                   // xs, sm, md, lg, xl
    "class": "custom-class",
    "id": "save-btn",
    "text": "Button Text",
    "icon": "bell",                 // single icon (icon-only button)
    "iconLeft": "save",             // left icon
    "iconRight": "arrow-right",     // right icon
    "type": "button",               // button, submit, reset
    "name": "action",
    "value": "save",
    "disabled": false,
    "loading": false,
    "hxPost": "/api/save",
    "hxGet": "/api/load",
    "hxTarget": "#result",
    "hxSwap": "innerHTML",
    "hxTrigger": "click",
    "alpineClick": "handleClick()"
}
```

### Predefined JSON Examples

```go
// Primary button with icon
@ButtonFrom(ExamplePrimaryButton)

// Loading button
@ButtonFrom(ExampleLoadingButton)

// HTMX button
@ButtonFrom(ExampleHTMXButton)

// Icon-only button
@ButtonFrom(ExampleIconButton)

// Submit button
@ButtonFrom(ExampleSubmitButton)
```

### Creating Props from JSON

```go
props, err := ButtonFromJSON(jsonString)
if err != nil {
    // Handle error
}

@Button(props)
```

## Examples

### Form Actions

```go
<form hx-post="/api/users" hx-target="#result">
    <input type="text" name="username" />
    
    <div class="flex gap-2">
        @Button(ButtonProps{
            Text: "Save",
            Type: "submit",
            Variant: ButtonPrimary,
            IconLeft: "save",
        })
        
        @Button(ButtonProps{
            Text: "Cancel",
            Type: "button",
            Variant: ButtonSecondary,
            AlpineClick: "resetForm()",
        })
    </div>
</form>
```

### Loading States

```go
<div x-data="{ loading: false }">
    @Button(ButtonProps{
        AlpineClick: "loading = true; setTimeout(() => loading = false, 2000)",
    }) {
        <span x-show="!loading">Submit</span>
        <span x-show="loading">Loading...</span>
    }
</div>
```

### Icon Buttons

```go
<div class="flex gap-2">
    @Button(ButtonProps{
        Icon: "edit",
        Variant: ButtonGhost,
        Size: ButtonSizeSM,
    })
    
    @Button(ButtonProps{
        Icon: "trash",
        Variant: ButtonDestructive,
        Size: ButtonSizeSM,
    })
    
    @Button(ButtonProps{
        Icon: "share",
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
    })
</div>
```

### HTMX Infinite Scroll

```go
@Button(ButtonProps{
    Text: "Load More",
    Variant: ButtonOutline,
    HXGet: "/api/items?page=2",
    HXTarget: "#items-list",
    HXSwap: "beforeend",
    IconRight: "arrow-down",
})
```

### Confirmation Dialog

```go
<div x-data="{ showConfirm: false }">
    @Button(ButtonProps{
        Text: "Delete Account",
        Variant: ButtonDestructive,
        AlpineClick: "showConfirm = true",
    })
    
    <div x-show="showConfirm" class="mt-4 p-4 border rounded">
        <p>Are you sure?</p>
        <div class="flex gap-2 mt-2">
            @Button(ButtonProps{
                Text: "Confirm",
                Variant: ButtonDestructive,
                Size: ButtonSizeSM,
                HXPost: "/api/delete-account",
            })
            
            @Button(ButtonProps{
                Text: "Cancel",
                Variant: ButtonGhost,
                Size: ButtonSizeSM,
                AlpineClick: "showConfirm = false",
            })
        </div>
    </div>
</div>
```

### Button Group

```go
<div class="inline-flex rounded-md shadow-sm" role="group">
    @Button(ButtonProps{
        Text: "Left",
        Variant: ButtonOutline,
        Class: "rounded-r-none",
    })
    
    @Button(ButtonProps{
        Text: "Middle",
        Variant: ButtonOutline,
        Class: "rounded-none border-l-0",
    })
    
    @Button(ButtonProps{
        Text: "Right",
        Variant: ButtonOutline,
        Class: "rounded-l-none border-l-0",
    })
</div>
```

### Dropdown Trigger

```go
<div x-data="{ open: false }" class="relative">
    @Button(ButtonProps{
        Variant: ButtonOutline,
        IconRight: "chevron-down",
        AlpineClick: "open = !open",
    }) { Options }
    
    <div x-show="open" class="absolute mt-2 w-48 rounded-md shadow-lg">
        <!-- Dropdown content -->
    </div>
</div>
```

### Using Fluent API in Template

```go
templ SaveForm() {
    <form>
        <input type="text" name="title" />
        
        @Button(NewButton().
            Primary().
            Large().
            WithIconLeft("save").
            WithText("Save Draft").
            AsSubmit().
            HXPost("/api/drafts").
            HXTarget("#result").
            Build(),
        )
    </form>
}
```

### Async Action with Feedback

```go
<div 
    x-data="{ saving: false }"
    @htmx:before-request="saving = true"
    @htmx:after-request="saving = false"
>
    @Button(ButtonProps{
        HXPost: "/api/save",
        HXTarget: "#result",
    }) {
        <span x-show="!saving">Save</span>
        <span x-show="saving">Saving...</span>
    }
</div>
```

### Social Login Buttons

```go
<div class="flex flex-col gap-2">
    @Button(ButtonProps{
        Variant: ButtonOutline,
        IconLeft: "github",
        Text: "Continue with GitHub",
        HXGet: "/auth/github",
    })
    
    @Button(ButtonProps{
        Variant: ButtonOutline,
        IconLeft: "google",
        Text: "Continue with Google",
        HXGet: "/auth/google",
    })
</div>
```

## API Reference

### ButtonProps

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `Variant` | `ButtonVariant` | `ButtonPrimary` | Visual style variant |
| `Size` | `ButtonSize` | `ButtonSizeMD` | Button size |
| `Class` | `string` | `""` | Additional CSS classes |
| `ID` | `string` | `""` | HTML id attribute |
| `Text` | `string` | `""` | Simple text content (alternative to children) |
| `Icon` | `string` | `""` | Single icon (icon-only button) |
| `IconLeft` | `string` | `""` | Left icon |
| `IconRight` | `string` | `""` | Right icon |
| `Type` | `string` | `"button"` | Button type: button, submit, reset |
| `Name` | `string` | `""` | Form field name |
| `Value` | `string` | `""` | Form field value |
| `Disabled` | `bool` | `false` | Disable the button |
| `Loading` | `bool` | `false` | Show loading spinner |
| `HXPost` | `string` | `""` | HTMX POST URL |
| `HXGet` | `string` | `""` | HTMX GET URL |
| `HXTarget` | `string` | `""` | HTMX target selector |
| `HXSwap` | `string` | `""` | HTMX swap strategy |
| `HXTrigger` | `string` | `""` | HTMX trigger event |
| `AlpineClick` | `string` | `""` | Alpine.js click handler |

### ButtonVariant

```go
const (
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonDestructive ButtonVariant = "destructive"
    ButtonOutline     ButtonVariant = "outline"
    ButtonGhost       ButtonVariant = "ghost"
    ButtonLink        ButtonVariant = "link"
)
```

### ButtonSize

```go
const (
    ButtonSizeXS ButtonSize = "xs"
    ButtonSizeSM ButtonSize = "sm"
    ButtonSizeMD ButtonSize = "md"
    ButtonSizeLG ButtonSize = "lg"
    ButtonSizeXL ButtonSize = "xl"
)
```

### Components

#### Button

Main button component:

```go
templ Button(props ButtonProps, children ...templ.Component)
```

#### Convenience Components

```go
templ PrimaryButton(props ButtonProps, children ...templ.Component)
templ SecondaryButton(props ButtonProps, children ...templ.Component)
templ DestructiveButton(props ButtonProps, children ...templ.Component)
```

#### JSON Component

```go
templ ButtonFrom(schemaJSON string)
```

### Helper Functions

```go
// Fluent API
func NewButton() *ButtonBuilder

// Functional Options
func NewButtonWithOptions(opts ...ButtonOption) ButtonProps

// JSON Parsing
func ButtonFromJSON(jsonStr string) (ButtonProps, error)
```

## Best Practices

1. **Use semantic variants** - Match button variant to action importance (primary for main actions, destructive for dangerous actions)
2. **Include meaningful text** - Even with icons, include text for accessibility
3. **Loading states** - Always show loading feedback for async actions
4. **Disable during submission** - Prevent double-submissions by using loading or disabled states
5. **Icon placement** - Use left icons for actions (save, add), right icons for navigation (next, external)
6. **Button types** - Use `type="submit"` for form submissions, `type="button"` for other actions
7. **HTMX integration** - Leverage HTMX for progressive enhancement
8. **Consistent sizing** - Use size variants consistently across your app
9. **Accessible labels** - The component handles loading aria-labels automatically

## Accessibility

The button component is built with accessibility in mind:

- ✅ Proper semantic `<button>` element
- ✅ Keyboard navigation support (Enter/Space)
- ✅ Focus visible styles with ring
- ✅ Loading state aria-label
- ✅ Disabled state prevents interaction
- ✅ Icon-only buttons should include aria-label (add manually if needed)

```go
// For icon-only buttons, add aria-label
@Button(ButtonProps{
    Icon: "trash",
    Class: "aria-label='Delete item'",
})
```

## Styling Customization

The button uses CSS custom properties (design tokens) that can be customized:

```css
:root {
    --primary: ...;
    --primary-foreground: ...;
    --secondary: ...;
    --destructive: ...;
    --radius-md: ...;
    --space-10: ...;
    /* etc */
}
```

For one-off customization:

```go
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Class: "bg-gradient-to-r from-purple-500 to-pink-500",
    Text: "Custom Gradient",
})
```

## Design System Integration

This button component integrates with your design system through CSS custom properties:

- **Colors**: `--primary`, `--secondary`, `--destructive`, `--accent`
- **Spacing**: `--space-{n}` tokens
- **Radius**: `--radius-md`
- **Ring**: `--ring` for focus states

This ensures consistent theming across your entire application.
