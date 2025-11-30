# Button Component Documentation
 > *This under Review See #3020*
A versatile, accessible button component with multiple variants, sizes, loading states, and full HTMX and Alpine.js integration. Built with compiled design system tokens for consistent theming.

## Table of Contents

- [Quick Start](#quick-start)
- [Basic Usage](#basic-usage)
- [Variants](#variants)
- [Sizes](#sizes)
- [Icons](#icons)
- [States](#states)
- [Interactive Buttons](#interactive-buttons)
- [Builder Pattern API](#builder-pattern-api)
- [Reusable Primitives](#reusable-primitives)
- [Advanced Examples](#advanced-examples)
- [API Reference](#api-reference)
- [Design System Integration](#design-system-integration)
- [Accessibility](#accessibility)
- [Best Practices](#best-practices)
- [Migration Guide](#migration-guide)
- [Troubleshooting](#troubleshooting)

---

## Quick Start

```go
// Simplest button
@Button(ButtonProps{Text: "Click Me"})

// Primary button with icon
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Save",
    IconLeft: "save",
})

// Using convenience component
@PrimaryButton(ButtonProps{Text: "Save"})

// Using builder pattern
@Button(NewButton(
    WithVariant(ButtonPrimary),
    WithText("Save"),
    WithIconLeft("save"),
))
```

---

## Basic Usage

### Simple Button

```go
// Basic button with text
@Button(ButtonProps{Text: "Click Me"})

// With children (takes precedence over Text)
@Button(ButtonProps{}) {
    Click Me
}

// Primary button (default variant)
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Save",
})
```

### Using Convenience Components

Pre-configured components for common button types:

```go
// Primary button
@PrimaryButton(ButtonProps{Text: "Save"})

// Secondary button
@SecondaryButton(ButtonProps{Text: "Cancel"})

// Outline button
@OutlineButton(ButtonProps{Text: "Learn More"})

// Destructive button (for dangerous actions)
@DestructiveButton(ButtonProps{Text: "Delete"})

// Ghost button (minimal style)
@GhostButton(ButtonProps{Text: "Skip"})

// Link button (styled as link)
@LinkButton(ButtonProps{Text: "Read More"})
```

> **NOTE:** Convenience components are backward compatible wrappers that set the variant automatically while passing through all other props.

---

## Variants

Six visual variants following your compiled design system:

```go
// Primary (Default) - Main call-to-action
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Primary Action",
})

// Secondary - Supporting actions
@Button(ButtonProps{
    Variant: ButtonSecondary,
    Text: "Secondary Action",
})

// Destructive - Dangerous actions (delete, remove)
@Button(ButtonProps{
    Variant: ButtonDestructive,
    Text: "Delete",
})

// Outline - Alternative style with border
@Button(ButtonProps{
    Variant: ButtonOutline,
    Text: "Outline",
})

// Ghost - Minimal style, no background
@Button(ButtonProps{
    Variant: ButtonGhost,
    Text: "Ghost",
})

// Link - Looks like a link but acts like a button
@Button(ButtonProps{
    Variant: ButtonLink,
    Text: "Link Style",
})
```

### Variant Usage Guidelines

| Variant | Use Case | Example |
|---------|----------|---------|
| **Primary** | Main action, call-to-action | Submit, Save, Continue |
| **Secondary** | Supporting actions | Cancel, Back, Close |
| **Destructive** | Dangerous/irreversible actions | Delete, Remove, Disconnect |
| **Outline** | Alternative primary action | Learn More, View Details |
| **Ghost** | Tertiary actions, minimal UI | Skip, Dismiss, Hide |
| **Link** | Navigation, inline actions | Read More, See All, Edit |

**CSS Classes Used:**
- `.btn` - Base button styles
- `.btn-primary`, `.btn-secondary`, `.btn-destructive`, `.btn-outline`, `.btn-ghost`, `.btn-link` - Variant styles from compiled CSS

---

## Sizes

Five size options with consistent spacing:

```go
// Extra Small
@Button(ButtonProps{
    Size: ButtonSizeXS,
    Text: "Extra Small",
})

// Small
@Button(ButtonProps{
    Size: ButtonSizeSM,
    Text: "Small",
})

// Medium (Default)
@Button(ButtonProps{
    Size: ButtonSizeMD,
    Text: "Medium",
})

// Large
@Button(ButtonProps{
    Size: ButtonSizeLG,
    Text: "Large",
})

// Extra Large
@Button(ButtonProps{
    Size: ButtonSizeXL,
    Text: "Extra Large",
})
```

**CSS Classes Used:**
- `.btn-xs`, `.btn-sm`, `.btn-md`, `.btn-lg`, `.btn-xl` - Size variants from compiled CSS

**Size Mapping:**
| Size | Height | Padding | Font Size | Icon Size | Use Case |
|------|--------|---------|-----------|-----------|----------|
| XS   | 1.75rem | 0.5rem | 0.75rem | xs | Dense UI, compact forms |
| SM   | 2rem   | 0.75rem | 0.875rem | sm | Secondary actions, toolbars |
| MD   | 2.5rem | 1rem   | 0.875rem | md | Default, most common use |
| LG   | 3rem   | 1.25rem | 1rem    | lg | Primary CTAs, hero sections |
| XL   | 3.5rem | 1.5rem | 1.125rem | xl | Marketing pages, landing pages |

> **NOTE:** Icon sizes automatically match button sizes via `mapButtonSizeToIconSize()` helper.

---

## Icons

Add icons to buttons for better visual communication:

### Left Icon

Icon appears before the text:

```go
@Button(ButtonProps{
    IconLeft: "save",
    Text: "Save Changes",
})
```

**CSS Class:** `.btn-icon-left` (provides consistent spacing)

### Right Icon

Icon appears after the text:

```go
@Button(ButtonProps{
    IconRight: "arrow-right",
    Text: "Next Step",
})
```

**CSS Class:** `.btn-icon-right` (provides consistent spacing)

### Both Icons

```go
@Button(ButtonProps{
    IconLeft: "download",
    IconRight: "external-link",
    Text: "Export Data",
})
```

### Icon-Only Button

For action buttons without text:

```go
@Button(ButtonProps{
    Icon: "trash",
    Variant: ButtonDestructive,
    Size: ButtonSizeSM,
    AriaLabel: "Delete item",
})
```

**CSS Class:** `.btn-icon` (centers the icon)

### Icon Best Practices

**✅ DO:**
- Use universally recognized icons (save, delete, search)
- Match icon semantics to action (trash for delete, checkmark for confirm)
- Always provide aria-label for icon-only buttons
- Use left icons for actions (save, download)
- Use right icons for navigation (next, external link)

**❌ DON'T:**
- Use obscure or ambiguous icons without text
- Mix icon styles (outline + solid)
- Use icons purely for decoration
- Forget accessibility labels

### Common Icon Patterns

```go
// Save/Submit actions - left icon
@Button(ButtonProps{
    IconLeft: "save",
    Text: "Save",
})

// Navigation - right icon
@Button(ButtonProps{
    IconRight: "arrow-right",
    Text: "Continue",
})

// External links - right icon
@Button(ButtonProps{
    IconRight: "external-link",
    Text: "Open Documentation",
})

// Downloads - left icon
@Button(ButtonProps{
    IconLeft: "download",
    Text: "Download Report",
})

// Refresh/Reload - left icon
@Button(ButtonProps{
    IconLeft: "refresh-cw",
    Text: "Refresh",
})
```

---

## States

### Disabled State

Prevents interaction and applies disabled styling:

```go
@Button(ButtonProps{
    Text: "Cannot Submit",
    Disabled: true,
})
```

**CSS Class:** `.btn-disabled`

**Behavior:**
- Sets `disabled` HTML attribute
- Pointer events disabled
- Visual opacity reduction
- Prevents form submission
- Blocks HTMX/Alpine handlers

**Use Cases:**
- Form validation (submit disabled until valid)
- Insufficient permissions
- Prerequisites not met
- Processing states

### Loading State

Shows a spinner and disables interaction:

```go
@Button(ButtonProps{
    Text: "Saving...",
    Loading: true,
})
```

**CSS Class:** `.btn-loading`

**Features:**
- Displays animated spinner (from `LoadingSpinner` component)
- Hides all icons and content
- Sets `aria-label="Loading..."` for accessibility
- Button remains disabled but visually indicates progress
- Prevents duplicate submissions

**Best Practices:**
```go
// ✅ Good: Change text during loading
@Button(ButtonProps{
    Text: if isLoading { "Saving..." } else { "Save" },
    Loading: isLoading,
})

// ✅ Good: With HTMX loading states
@Button(ButtonProps{
    Text: "Submit",
    HXPost: "/api/save",
    ClassName: "htmx-indicator",
})

// ❌ Bad: Same text when loading
@Button(ButtonProps{
    Text: "Save",
    Loading: true,
})
```

### Button Types

Standard HTML button types:

```go
// Default button (no form action)
@Button(ButtonProps{
    Text: "Button",
    Type: "button", // default
})

// Submit button (submits parent form)
@Button(ButtonProps{
    Text: "Submit Form",
    Type: "submit",
})

// Reset button (resets parent form)
@Button(ButtonProps{
    Text: "Reset",
    Type: "reset",
})
```

> **NOTE:** Type defaults to `"button"` to prevent accidental form submissions.

### Focus State

Automatically handled via CSS:

```css
.btn:focus-visible {
    outline: 2px solid hsl(var(--ring));
    outline-offset: 2px;
}
```

**Keyboard Interaction:**
- **Tab** - Navigate to button
- **Enter/Space** - Activate button
- **Shift+Tab** - Navigate backward

---

## Interactive Buttons

### HTMX Integration

Full support for HTMX attributes:

#### GET Request

```go
@Button(ButtonProps{
    Text: "Load More",
    HXGet: "/api/items",
    HXTarget: "#items-list",
    HXSwap: "beforeend",
})
```

#### POST Request

```go
@Button(ButtonProps{
    Text: "Delete Item",
    Variant: ButtonDestructive,
    HXPost: "/api/items/123",
    HXTarget: "#item-123",
    HXSwap: "outerHTML",
    HXConfirm: "Are you sure you want to delete this item?",
})
```

#### Custom Triggers

```go
@Button(ButtonProps{
    Text: "Auto Save",
    HXPost: "/api/save",
    HXTrigger: "click throttle:1s",
    IconLeft: "save",
})
```

#### Advanced HTMX Patterns

```go
// Polling button
@Button(ButtonProps{
    Text: "Start Monitoring",
    HXGet: "/api/status",
    HXTrigger: "click, every 5s",
    HXTarget: "#status",
})

// Optimistic UI
@Button(ButtonProps{
    Text: "Like",
    HXPost: "/api/like",
    HXSwap: "outerHTML swap:0s",
    IconLeft: "heart",
})

// Push URL
@Button(ButtonProps{
    Text: "Next Page",
    HXGet: "/page/2",
    HXPushURL: "true",
    HXTarget: "#content",
})
```

**Supported HTMX Attributes:**
- `hx-post` - POST request
- `hx-get` - GET request
- `hx-target` - Where to swap response
- `hx-swap` - How to swap (innerHTML, outerHTML, beforeend, afterbegin, beforebegin, afterend, delete, none)
- `hx-trigger` - Custom trigger conditions
- `hx-confirm` - Confirmation dialog before request
- `hx-push-url` - Update browser URL
- `hx-select` - Select portion of response

### Alpine.js Integration

```go
// Simple click handler
@Button(ButtonProps{
    Text: "Toggle Modal",
    AlpineClick: "open = !open",
})

// With Alpine data context
<div x-data="{ count: 0 }">
    @Button(ButtonProps{
        AlpineClick: "count++",
    }) {
        <span>Clicked: </span>
        <span x-text="count"></span>
    }
</div>

// Method calls
@Button(ButtonProps{
    Text: "Submit",
    AlpineClick: "submitForm()",
})

// Multiple statements
@Button(ButtonProps{
    Text: "Process",
    AlpineClick: "loading = true; await processData(); loading = false",
})
```

### Combined HTMX + Alpine

```go
<div x-data="{ saving: false }">
    @Button(ButtonProps{
        Text: "Save",
        HXPost: "/api/save",
        HXTarget: "#result",
        AlpineClick: "saving = true",
        ClassName: "@htmx:after-request='saving = false'",
    })
    
    <span x-show="saving" class="text-sm text-muted-foreground">
        Saving in progress...
    </span>
</div>
```

### HTMX Loading Indicators

```go
// Method 1: Using htmx-indicator class
<div x-data="{ }">
    @Button(ButtonProps{
        Text: "Submit",
        HXPost: "/api/save",
        HXIndicator: "#spinner",
    })
    
    <div id="spinner" class="htmx-indicator">
        @LoadingSpinner(LoadingSpinnerProps{})
    </div>
</div>

// Method 2: Using Alpine with HTMX events
<div 
    x-data="{ loading: false }"
    @htmx:before-request="loading = true"
    @htmx:after-request="loading = false"
>
    @Button(ButtonProps{
        Text: "Submit",
        HXPost: "/api/save",
        ClassName: "x-bind:disabled='loading'",
    })
</div>
```

---

## Builder Pattern API

Fluent API for constructing buttons with chainable methods:

### Basic Usage

```go
props := NewButton(
    WithVariant(ButtonPrimary),
    WithText("Save Changes"),
    WithIconLeft("save"),
    WithSize(ButtonSizeLG),
)

@Button(props)
```

### All Builder Functions

#### Variant Options

```go
WithVariant(ButtonVariant)    // Set any variant
WithPrimary()                 // Shortcut for primary variant
WithSecondary()               // Shortcut for secondary variant
WithDestructive()             // Shortcut for destructive variant
WithOutline()                 // Shortcut for outline variant
WithGhost()                   // Shortcut for ghost variant
WithLink()                    // Shortcut for link variant
```

#### Size Options

```go
WithSize(ButtonSize)          // Set any size
WithSizeXS()                  // Extra small size
WithSizeSM()                  // Small size
WithSizeMD()                  // Medium size (default)
WithSizeLG()                  // Large size
WithSizeXL()                  // Extra large size
```

#### Content Options

```go
WithText(string)              // Button text
WithIcon(string)              // Single icon (icon-only)
WithIconLeft(string)          // Left icon
WithIconRight(string)         // Right icon
WithClass(string)             // Custom CSS classes
WithID(string)                // HTML id attribute
WithAriaLabel(string)         // Accessibility label
```

#### Attribute Options

```go
WithType(string)              // Button type (button, submit, reset)
WithName(string)              // Form field name
WithValue(string)             // Form field value
AsSubmit()                    // Shortcut for type="submit"
AsReset()                     // Shortcut for type="reset"
AsDisabled()                  // Mark as disabled
AsLoading()                   // Show loading state
```

#### HTMX Options

```go
WithHXPost(string)            // hx-post URL
WithHXGet(string)             // hx-get URL
WithHXTarget(string)          // hx-target selector
WithHXSwap(string)            // hx-swap strategy
WithHXTrigger(string)         // hx-trigger conditions
WithHXConfirm(string)         // hx-confirm dialog
WithHXPushURL(string)         // hx-push-url
WithHXSelect(string)          // hx-select
WithHXIndicator(string)       // hx-indicator
```

#### Alpine.js Options

```go
WithAlpineClick(string)       // x-on:click handler
WithAlpineBind(string)        // x-bind attributes
```

### Complex Example

```go
saveButton := NewButton(
    WithPrimary(),
    WithSizeLG(),
    WithText("Save Draft"),
    WithIconLeft("save"),
    WithIconRight("cloud"),
    AsSubmit(),
    WithHXPost("/api/drafts"),
    WithHXTarget("#result"),
    WithHXSwap("innerHTML"),
    WithHXConfirm("Save this draft?"),
)

@Button(saveButton)
```

### Preset Configurations

Create reusable button configurations:

```go
// Define presets
var (
    SaveButtonBase = []ButtonOption{
        WithPrimary(),
        WithIconLeft("save"),
        WithText("Save"),
        AsSubmit(),
    }
    
    CancelButtonBase = []ButtonOption{
        WithSecondary(),
        WithText("Cancel"),
    }
    
    DeleteButtonBase = []ButtonOption{
        WithDestructive(),
        WithIconLeft("trash"),
        WithHXConfirm("Are you sure?"),
    }
)

// Use presets
@Button(NewButton(SaveButtonBase...))

// Extend presets
@Button(NewButton(
    append(SaveButtonBase, WithSizeLG())...,
))

// Combine presets
customButton := NewButton(
    append(append(SaveButtonBase, WithHXPost("/api/save")), 
        WithHXTarget("#result"))...,
)
```

### Builder Pattern Benefits

**✅ Advantages:**
- Type-safe configuration
- Autocomplete support
- Reusable presets
- Composable options
- Clear intent

**Example Comparisons:**

```go
// Traditional struct initialization
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size: ButtonSizeLG,
    Text: "Save",
    IconLeft: "save",
    Type: "submit",
    HXPost: "/api/save",
    HXTarget: "#result",
})

// Builder pattern - more readable
@Button(NewButton(
    WithPrimary(),
    WithSizeLG(),
    WithText("Save"),
    WithIconLeft("save"),
    AsSubmit(),
    WithHXPost("/api/save"),
    WithHXTarget("#result"),
))
```

---

## Reusable Primitives

### LoadingSpinner Component

The loading spinner is extracted as a reusable primitive that can be used anywhere:

#### Usage in Buttons (Automatic)

```go
@Button(ButtonProps{
    Loading: true,
    Text: "Processing...",
})
```

#### Standalone Usage

```go
// Default spinner
@LoadingSpinner(LoadingSpinnerProps{})

// Custom sized spinner
@LoadingSpinner(LoadingSpinnerProps{
    Size: "lg",
})

// Spinner with custom styling
@LoadingSpinner(LoadingSpinnerProps{
    Size: "md",
    ClassName: "text-blue-500",
})

// Spinner with custom aria-label
@LoadingSpinner(LoadingSpinnerProps{
    Size: "sm",
    AriaLabel: "Processing your request",
})
```

#### LoadingSpinnerProps

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `Size` | `string` | `"md"` | Size: xs, sm, md, lg, xl |
| `ClassName` | `string` | `""` | Additional CSS classes |
| `AriaLabel` | `string` | `"Loading..."` | Accessibility label |

**CSS Classes:**
- `.spinner` - Base spinner styles with animation
- `.spinner-xs`, `.spinner-sm`, `.spinner-md`, `.spinner-lg`, `.spinner-xl` - Size variants

**Use Cases:**
- Button loading states (automatic)
- Page loading overlays
- Form submission indicators
- Data fetching states
- Inline content loading
- Skeleton screens
- Progress indicators

#### Advanced Spinner Patterns

```go
// Full-page loader
<div class="fixed inset-0 flex items-center justify-center bg-black/50 z-50">
    @LoadingSpinner(LoadingSpinnerProps{
        Size: "xl",
        ClassName: "text-white",
    })
</div>

// Inline loading
<div class="flex items-center gap-2">
    @LoadingSpinner(LoadingSpinnerProps{Size: "sm"})
    <span>Loading data...</span>
</div>

// Table cell loading
<td class="text-center py-4">
    @LoadingSpinner(LoadingSpinnerProps{Size: "sm"})
</td>

// Card loading state
<div class="card">
    <div class="flex justify-center py-8">
        @LoadingSpinner(LoadingSpinnerProps{
            Size: "lg",
            AriaLabel: "Loading content",
        })
    </div>
</div>
```

> **NOTE:** This is a primitive component following the refactoring strategy. It can be composed into other components beyond buttons.

---

## Advanced Examples

### Form Submission

```go
<form hx-post="/api/users" hx-target="#result">
    <div class="space-y-4">
        <input 
            type="text" 
            name="username" 
            placeholder="Username"
            class="input"
            required
        />
        <input 
            type="email" 
            name="email" 
            placeholder="Email"
            class="input"
            required
        />
    </div>
    
    <div class="flex gap-2 mt-6">
        @Button(ButtonProps{
            Text: "Create User",
            Type: "submit",
            Variant: ButtonPrimary,
            IconLeft: "user-plus",
            Size: ButtonSizeLG,
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

### Loading State with Alpine

```go
<div 
    x-data="{ loading: false }"
    @htmx:before-request="loading = true"
    @htmx:after-request="loading = false"
    @htmx:response-error="loading = false"
>
    @Button(ButtonProps{
        HXPost: "/api/save",
        HXTarget: "#result",
    }) {
        <span x-show="!loading">Submit</span>
        <span x-show="loading" class="flex items-center gap-2">
            @LoadingSpinner(LoadingSpinnerProps{Size: "sm"})
            <span>Saving...</span>
        </span>
    }
</div>
```

### Icon Button Grid

```go
<div class="flex gap-2">
    @Button(ButtonProps{
        Icon: "edit",
        Variant: ButtonGhost,
        Size: ButtonSizeSM,
        AriaLabel: "Edit",
        HXGet: "/edit/123",
    })
    
    @Button(ButtonProps{
        Icon: "trash",
        Variant: ButtonDestructive,
        Size: ButtonSizeSM,
        AriaLabel: "Delete",
        HXPost: "/delete/123",
        HXConfirm: "Delete this item?",
    })
    
    @Button(ButtonProps{
        Icon: "share",
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
        AriaLabel: "Share",
        AlpineClick: "shareModal = true",
    })
    
    @Button(ButtonProps{
        Icon: "download",
        Variant: ButtonPrimary,
        Size: ButtonSizeSM,
        AriaLabel: "Download",
        HXGet: "/download/123",
    })
</div>
```

### HTMX Infinite Scroll

```go
<div id="items-list" class="space-y-4">
    <!-- Items loaded dynamically -->
</div>

<div class="mt-6 text-center">
    @Button(ButtonProps{
        Text: "Load More",
        Variant: ButtonOutline,
        Size: ButtonSizeLG,
        HXGet: "/api/items?page=2",
        HXTarget: "#items-list",
        HXSwap: "beforeend",
        IconRight: "arrow-down",
        ClassName: "w-full max-w-xs",
    })
</div>
```

### Confirmation Dialog with Alpine

```go
<div x-data="{ showConfirm: false }">
    @Button(ButtonProps{
        Text: "Delete Account",
        Variant: ButtonDestructive,
        IconLeft: "alert-triangle",
        AlpineClick: "showConfirm = true",
    })
    
    <div 
        x-show="showConfirm" 
        class="mt-4 p-4 border-2 border-destructive rounded-lg bg-destructive/10"
        x-transition:enter="transition ease-out duration-200"
        x-transition:enter-start="opacity-0 scale-95"
        x-transition:enter-end="opacity-100 scale-100"
    >
        <div class="flex items-start gap-3">
            <div class="flex-shrink-0 text-destructive">
                <!-- Alert Icon -->
            </div>
            <div class="flex-1">
                <h3 class="font-semibold text-lg">Are you absolutely sure?</h3>
                <p class="text-sm text-muted-foreground mt-1">
                    This action cannot be undone. This will permanently delete your account
                    and remove all your data from our servers.
                </p>
                
                <div class="flex gap-2 mt-4">
                    @Button(ButtonProps{
                        Text: "Yes, Delete My Account",
                        Variant: ButtonDestructive,
                        Size: ButtonSizeSM,
                        HXPost: "/api/delete-account",
                        HXTarget: "body",
                        HXConfirm: "Type DELETE to confirm",
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
    </div>
</div>
```

### Button Group

```go
<div class="inline-flex rounded-md shadow-sm" role="group" aria-label="Text alignment">
    @Button(ButtonProps{
        Icon: "align-left",
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
        ClassName: "rounded-r-none",
        AriaLabel: "Align left",
        AlpineClick: "align = 'left'",
    })
    
    @Button(ButtonProps{
        Icon: "align-center",
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
        ClassName: "rounded-none border-l-0",
        AriaLabel: "Align center",
        AlpineClick: "align = 'center'",
    })
    
    @Button(ButtonProps{
        Icon: "align-right",
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
        ClassName: "rounded-l-none border-l-0",
        AriaLabel: "Align right",
        AlpineClick: "align = 'right'",
    })
</div>
```

### Dropdown Trigger

```go
<div x-data="{ open: false }" class="relative">
    @Button(ButtonProps{
        Variant: ButtonOutline,
        Text: "Options",
        IconRight: "chevron-down",
        AlpineClick: "open = !open",
        ClassName: "x-bind:aria-expanded='open'",
    })
    
    <div 
        x-show="open" 
        @click.away="open = false"
        x-transition:enter="transition ease-out duration-100"
        x-transition:enter-start="opacity-0 scale-95"
        x-transition:enter-end="opacity-100 scale-100"
        class="absolute right-0 mt-2 w-56 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 z-50"
    >
        <div class="py-1" role="menu" aria-orientation="vertical">
            <a href="#" class="block px-4 py-2 text-sm hover:bg-gray-100" role="menuitem">
                Edit
            </a>
            <a href="#" class="block px-4 py-2 text-sm hover:bg-gray-100" role="menuitem">
                Duplicate
            </a>
            <a href="#" class="block px-4 py-2 text-sm text-destructive hover:bg-gray-100" role="menuitem">
                Delete
            </a>
        </div>
    </div>
</div>
```

### Social Login Buttons

```go
<div class="flex flex-col gap-3">
    @Button(ButtonProps{
        Variant: ButtonOutline,
        Size: ButtonSizeLG,
        IconLeft: "github",
        Text: "Continue with GitHub",
        HXGet: "/auth/github",
        ClassName: "w-full justify-center",
    })
    
    @Button(ButtonProps{
        Variant: ButtonOutline,
        Size: ButtonSizeLG,
        IconLeft: "google",
        Text: "Continue with Google",
        HXGet: "/auth/google",
        ClassName: "w-full justify-center",
    })
    
    @Button(ButtonProps{
        Variant: ButtonOutline,
        Size: ButtonSizeLG,
        IconLeft: "twitter",
        Text: "Continue with Twitter",
        HXGet: "/auth/twitter",
        ClassName: "w-full justify-center",
    })
</div>

    <div class="relative my-6">
    <div class="absolute inset-0 flex items-center">
        <div class="w-full border-t border-gray-300"></div>
    </div>
    <div class="relative flex justify-center text-sm">
        <span class="px-2 bg-white text-gray-500">Or continue with</span>
    </div>
</div>
```

### Multi-Step Form Navigation

```go
<div x-data="{ step: 1, maxSteps: 3 }">
    <!-- Progress indicator -->
    <div class="mb-8">
        <div class="flex justify-between mb-2">
            <span class="text-sm font-medium">Step <span x-text="step"></span> of 3</span>
            <span class="text-sm text-muted-foreground" x-text="Math.round((step / maxSteps) * 100) + '%'"></span>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-2">
            <div 
                class="bg-primary h-2 rounded-full transition-all duration-300"
                x-bind:style="'width: ' + (step / maxSteps * 100) + '%'"
            ></div>
        </div>
    </div>
    
    <!-- Step content -->
    <div x-show="step === 1" class="space-y-4">
        <h2 class="text-xl font-semibold">Personal Information</h2>
        <!-- Form fields -->
    </div>
    
    <div x-show="step === 2" class="space-y-4">
        <h2 class="text-xl font-semibold">Account Details</h2>
        <!-- Form fields -->
    </div>
    
    <div x-show="step === 3" class="space-y-4">
        <h2 class="text-xl font-semibold">Review & Submit</h2>
        <!-- Summary -->
    </div>
    
    <!-- Navigation -->
    <div class="flex justify-between mt-8">
        @Button(ButtonProps{
            Text: "Previous",
            Variant: ButtonSecondary,
            IconLeft: "arrow-left",
            AlpineClick: "step--",
            ClassName: "x-show='step > 1'",
        })
        
        <div class="ml-auto flex gap-2">
            @Button(ButtonProps{
                Text: "Save Draft",
                Variant: ButtonGhost,
                IconLeft: "save",
                HXPost: "/api/draft",
                ClassName: "x-show='step < maxSteps'",
            })
            
            @Button(ButtonProps{
                Variant: ButtonPrimary,
                IconRight: "arrow-right",
                AlpineClick: "step++",
                ClassName: "x-show='step < maxSteps'",
            }) {
                <span x-text="step < maxSteps ? 'Next' : ''"></span>
            }
            
            @Button(ButtonProps{
                Text: "Submit",
                Variant: ButtonPrimary,
                Type: "submit",
                IconLeft: "check",
                ClassName: "x-show='step === maxSteps'",
            })
        </div>
    </div>
</div>
```

### Async Action with Error Handling

```go
<div 
    x-data="{ 
        loading: false, 
        error: null,
        async handleSubmit() {
            this.loading = true;
            this.error = null;
            try {
                const response = await fetch('/api/save', {
                    method: 'POST',
                    body: JSON.stringify(this.formData)
                });
                if (!response.ok) throw new Error('Save failed');
                this.loading = false;
            } catch (err) {
                this.error = err.message;
                this.loading = false;
            }
        }
    }"
>
    @Button(ButtonProps{
        AlpineClick: "handleSubmit()",
        ClassName: "x-bind:disabled='loading'",
    }) {
        <span x-show="!loading">Save Changes</span>
        <span x-show="loading" class="flex items-center gap-2">
            @LoadingSpinner(LoadingSpinnerProps{Size: "sm"})
            <span>Saving...</span>
        </span>
    }
    
    <div x-show="error" class="mt-2 text-sm text-destructive" x-text="error"></div>
</div>
```

### File Upload Button

```go
<div x-data="{ fileName: null }">
    <input 
        type="file" 
        id="file-upload" 
        class="hidden"
        @change="fileName = $event.target.files[0]?.name"
    />
    
    <label for="file-upload">
        @Button(ButtonProps{
            Variant: ButtonOutline,
            IconLeft: "upload",
            ClassName: "cursor-pointer",
        }) {
            <span x-show="!fileName">Choose File</span>
            <span x-show="fileName" x-text="fileName"></span>
        }
    </label>
</div>
```

### Toast Notification Trigger

```go
<div x-data="{ showToast: false }">
    @Button(ButtonProps{
        Text: "Show Notification",
        Variant: ButtonPrimary,
        IconLeft: "bell",
        AlpineClick: "showToast = true; setTimeout(() => showToast = false, 3000)",
    })
    
    <div 
        x-show="showToast"
        x-transition:enter="transition ease-out duration-300"
        x-transition:enter-start="opacity-0 translate-y-4"
        x-transition:enter-end="opacity-100 translate-y-0"
        class="fixed bottom-4 right-4 bg-white rounded-lg shadow-lg p-4 max-w-sm z-50"
    >
        <div class="flex items-start gap-3">
            <div class="text-green-500">✓</div>
            <div>
                <h4 class="font-semibold">Success</h4>
                <p class="text-sm text-muted-foreground">Your changes have been saved.</p>
            </div>
        </div>
    </div>
</div>
```

### Copy to Clipboard

```go
<div x-data="{ copied: false }">
    @Button(ButtonProps{
        Variant: ButtonOutline,
        Size: ButtonSizeSM,
        AlpineClick: "navigator.clipboard.writeText('https://example.com'); copied = true; setTimeout(() => copied = false, 2000)",
    }) {
        <span x-show="!copied" class="flex items-center gap-2">
            <span>Copy Link</span>
        </span>
        <span x-show="copied" class="flex items-center gap-2 text-green-600">
            <span>✓ Copied!</span>
        </span>
    }
</div>
```

---

## API Reference

### ButtonProps

Complete property reference:

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `Variant` | `ButtonVariant` | `ButtonPrimary` | Visual style variant |
| `Size` | `ButtonSize` | `ButtonSizeMD` | Button dimensions |
| `ClassName` | `string` | `""` | Additional CSS classes (merged with TwMerge) |
| `ID` | `string` | `""` | HTML id attribute |
| `Text` | `string` | `""` | Simple text content (ignored if children provided) |
| `Icon` | `string` | `""` | Single icon for icon-only buttons |
| `IconLeft` | `string` | `""` | Icon before content |
| `IconRight` | `string` | `""` | Icon after content |
| `Type` | `string` | `"button"` | HTML button type (button, submit, reset) |
| `Name` | `string` | `""` | Form field name |
| `Value` | `string` | `""` | Form field value |
| `Disabled` | `bool` | `false` | Disable interaction |
| `Loading` | `bool` | `false` | Show loading spinner |
| `AriaLabel` | `string` | `""` | Accessibility label (required for icon-only) |
| `HXPost` | `string` | `""` | HTMX POST URL |
| `HXGet` | `string` | `""` | HTMX GET URL |
| `HXTarget` | `string` | `""` | HTMX target selector |
| `HXSwap` | `string` | `""` | HTMX swap strategy |
| `HXTrigger` | `string` | `""` | HTMX trigger event |
| `HXConfirm` | `string` | `""` | HTMX confirmation dialog |
| `HXPushURL` | `string` | `""` | HTMX push URL to history |
| `HXSelect` | `string` | `""` | HTMX select portion of response |
| `HXIndicator` | `string` | `""` | HTMX loading indicator selector |
| `AlpineClick` | `string` | `""` | Alpine.js click handler |

### ButtonVariant Constants

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

**Mapped CSS Classes:**
- `ButtonPrimary` → `.btn-primary`
- `ButtonSecondary` → `.btn-secondary`
- `ButtonDestructive` → `.btn-destructive`
- `ButtonOutline` → `.btn-outline`
- `ButtonGhost` → `.btn-ghost`
- `ButtonLink` → `.btn-link`

### ButtonSize Constants

```go
const (
    ButtonSizeXS ButtonSize = "xs"
    ButtonSizeSM ButtonSize = "sm"
    ButtonSizeMD ButtonSize = "md"
    ButtonSizeLG ButtonSize = "lg"
    ButtonSizeXL ButtonSize = "xl"
)
```

**Mapped CSS Classes:**
- `ButtonSizeXS` → `.btn-xs`
- `ButtonSizeSM` → `.btn-sm`
- `ButtonSizeMD` → `.btn-md`
- `ButtonSizeLG` → `.btn-lg`
- `ButtonSizeXL` → `.btn-xl`

### Component Signatures

#### Main Button Component

```go
templ Button(props ButtonProps, children ...templ.Component)
```

Renders a button with full customization support.

**Parameters:**
- `props` - ButtonProps struct with all configuration
- `children` - Optional child components (takes precedence over `Text` prop)

**Returns:** templ.Component

#### Convenience Components

```go
templ PrimaryButton(props ButtonProps, children ...templ.Component)
templ SecondaryButton(props ButtonProps, children ...templ.Component)
templ OutlineButton(props ButtonProps, children ...templ.Component)
templ DestructiveButton(props ButtonProps, children ...templ.Component)
templ GhostButton(props ButtonProps, children ...templ.Component)
templ LinkButton(props ButtonProps, children ...templ.Component)
```

Pre-configured buttons with variant defaults. All props except `Variant` can be customized.

#### Loading Spinner Primitive

```go
templ LoadingSpinner(props LoadingSpinnerProps)
```

Reusable animated loading indicator.

**LoadingSpinnerProps:**
```go
type LoadingSpinnerProps struct {
    Size      string // xs, sm, md, lg, xl
    ClassName string // Additional CSS classes
    AriaLabel string // Accessibility label
}
```

### Helper Functions

#### Builder Pattern

```go
func NewButton(opts ...ButtonOption) ButtonProps
```

Creates ButtonProps using functional options pattern.

**Example:**
```go
props := NewButton(
    WithPrimary(),
    WithText("Save"),
    WithIconLeft("save"),
)
```

#### Button Options (Complete List)

```go
// Variant shortcuts
func WithVariant(variant ButtonVariant) ButtonOption
func WithPrimary() ButtonOption
func WithSecondary() ButtonOption
func WithDestructive() ButtonOption
func WithOutline() ButtonOption
func WithGhost() ButtonOption
func WithLink() ButtonOption

// Size shortcuts
func WithSize(size ButtonSize) ButtonOption
func WithSizeXS() ButtonOption
func WithSizeSM() ButtonOption
func WithSizeMD() ButtonOption
func WithSizeLG() ButtonOption
func WithSizeXL() ButtonOption

// Content
func WithText(text string) ButtonOption
func WithIcon(icon string) ButtonOption
func WithIconLeft(icon string) ButtonOption
func WithIconRight(icon string) ButtonOption
func WithClass(class string) ButtonOption
func WithID(id string) ButtonOption
func WithAriaLabel(label string) ButtonOption

// Attributes
func WithType(buttonType string) ButtonOption
func WithName(name string) ButtonOption
func WithValue(value string) ButtonOption
func AsSubmit() ButtonOption
func AsReset() ButtonOption
func AsDisabled() ButtonOption
func AsLoading() ButtonOption

// HTMX (Complete)
func WithHXPost(url string) ButtonOption
func WithHXGet(url string) ButtonOption
func WithHXTarget(target string) ButtonOption
func WithHXSwap(swap string) ButtonOption
func WithHXTrigger(trigger string) ButtonOption
func WithHXConfirm(message string) ButtonOption
func WithHXPushURL(url string) ButtonOption
func WithHXSelect(selector string) ButtonOption
func WithHXIndicator(selector string) ButtonOption

// Alpine.js
func WithAlpineClick(handler string) ButtonOption
```

#### Internal Helpers

```go
func mapButtonSizeToIconSize(size ButtonSize) string
```

Maps button size to corresponding icon size.

**Mapping:**
- `ButtonSizeXS` → `"xs"`
- `ButtonSizeSM` → `"sm"`
- `ButtonSizeMD` → `"md"`
- `ButtonSizeLG` → `"lg"`
- `ButtonSizeXL` → `"xl"`

```go
func buttonClasses(props ButtonProps) string
```

Generates the complete CSS class string for a button using TwMerge.

**Returns:** Space-separated string of CSS classes with conflicts resolved.

---

## Design System Integration

### Compiled CSS Classes

The button component uses compiled CSS classes from your design system:

```css
/* Base button styles */
.btn {
    @apply inline-flex items-center justify-center rounded-md;
    @apply font-medium transition-colors;
    @apply focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-offset-2;
    @apply disabled:pointer-events-none disabled:opacity-50;
    @apply select-none;
}

/* Variant classes */
.btn-primary {
    background-color: hsl(var(--primary));
    color: hsl(var(--primary-foreground));
}

.btn-primary:hover {
    background-color: hsl(var(--primary) / 0.9);
}

.btn-secondary {
    background-color: hsl(var(--secondary));
    color: hsl(var(--secondary-foreground));
}

.btn-secondary:hover {
    background-color: hsl(var(--secondary) / 0.8);
}

.btn-destructive {
    background-color: hsl(var(--destructive));
    color: hsl(var(--destructive-foreground));
}

.btn-destructive:hover {
    background-color: hsl(var(--destructive) / 0.9);
}

.btn-outline {
    border: 1px solid hsl(var(--input));
    background-color: transparent;
}

.btn-outline:hover {
    background-color: hsl(var(--accent));
    color: hsl(var(--accent-foreground));
}

.btn-ghost {
    background-color: transparent;
}

.btn-ghost:hover {
    background-color: hsl(var(--accent));
    color: hsl(var(--accent-foreground));
}

.btn-link {
    color: hsl(var(--primary));
    text-decoration-line: underline;
    background-color: transparent;
}

.btn-link:hover {
    text-decoration-line: none;
}

/* Size classes */
.btn-xs {
    height: 1.75rem;
    padding-left: 0.5rem;
    padding-right: 0.5rem;
    font-size: 0.75rem;
}

.btn-sm {
    height: 2rem;
    padding-left: 0.75rem;
    padding-right: 0.75rem;
    font-size: 0.875rem;
}

.btn-md {
    height: 2.5rem;
    padding-left: 1rem;
    padding-right: 1rem;
    font-size: 0.875rem;
}

.btn-lg {
    height: 3rem;
    padding-left: 1.25rem;
    padding-right: 1.25rem;
    font-size: 1rem;
}

.btn-xl {
    height: 3.5rem;
    padding-left: 1.5rem;
    padding-right: 1.5rem;
    font-size: 1.125rem;
}

/* Icon positioning */
.btn-icon-left {
    margin-right: 0.5rem;
}

.btn-icon-right {
    margin-left: 0.5rem;
}

.btn-icon {
    padding: 0;
    aspect-ratio: 1;
}

/* State classes */
.btn-disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.btn-loading {
    position: relative;
    color: transparent;
}

.btn-loading > * {
    visibility: hidden;
}

/* Loading spinner */
.spinner {
    animation: spin 1s linear infinite;
    border: 2px solid currentColor;
    border-right-color: transparent;
    border-radius: 50%;
}

.spinner-xs { width: 0.75rem; height: 0.75rem; }
.spinner-sm { width: 1rem; height: 1rem; }
.spinner-md { width: 1.25rem; height: 1.25rem; }
.spinner-lg { width: 1.5rem; height: 1.5rem; }
.spinner-xl { width: 2rem; height: 2rem; }

@keyframes spin {
    from { transform: rotate(0deg); }
    to { transform: rotate(360deg); }
}
```

### CSS Custom Properties (Design Tokens)

The button uses CSS variables that can be customized:

```css
:root {
    /* Colors - HSL format for easy manipulation */
    --primary: 222.2 47.4% 11.2%;
    --primary-foreground: 210 40% 98%;
    --secondary: 210 40% 96.1%;
    --secondary-foreground: 222.2 47.4% 11.2%;
    --destructive: 0 84.2% 60.2%;
    --destructive-foreground: 210 40% 98%;
    --accent: 210 40% 96.1%;
    --accent-foreground: 222.2 47.4% 11.2%;
    --input: 214.3 31.8% 91.4%;
    
    /* Spacing tokens */
    --spacing-xs: 0.25rem;   /* 4px */
    --spacing-sm: 0.5rem;    /* 8px */
    --spacing-md: 1rem;      /* 16px */
    --spacing-lg: 1.5rem;    /* 24px */
    --spacing-xl: 2rem;      /* 32px */
    
    /* Border radius */
    --radius-sm: 0.125rem;   /* 2px */
    --radius-md: 0.375rem;   /* 6px */
    --radius-lg: 0.5rem;     /* 8px */
    --radius-xl: 0.75rem;    /* 12px */
    
    /* Focus ring */
    --ring: 222.2 47.4% 11.2%;
    --ring-offset: 2px;
    
    /* Transition timing */
    --transition-fast: 150ms;
    --transition-base: 200ms;
    --transition-slow: 300ms;
}

/* Dark mode overrides */
.dark {
    --primary: 210 40% 98%;
    --primary-foreground: 222.2 47.4% 11.2%;
    --secondary: 217.2 32.6% 17.5%;
    --secondary-foreground: 210 40% 98%;
    --destructive: 0 62.8% 30.6%;
    --destructive-foreground: 210 40% 98%;
    --accent: 217.2 32.6% 17.5%;
    --accent-foreground: 210 40% 98%;
}

/* Tenant-specific themes */
[data-tenant="enterprise"] {
    --primary: 217 91% 60%;
    --spacing-md: 1.25rem;
    --radius-md: 0.25rem;
}

[data-tenant="minimal"] {
    --primary: 0 0% 0%;
    --radius-md: 0;
}
```

### Theme Customization

Override design tokens at different levels:

```css
/* Global theme */
:root {
    --primary: 220 70% 50%;
}

/* Page-specific */
.marketing-page {
    --primary: 280 60% 55%;
    --spacing-md: 1.5rem;
}

/* Component-specific */
.hero-section .btn-primary {
    --primary: 340 75% 55%;
}
```

### TwMerge Integration

The component uses `TwMerge` to intelligently resolve class conflicts:

```go
// If both base and custom classes have padding
buttonClasses(ButtonProps{
    Variant: ButtonPrimary,  // includes .btn-md with padding
    ClassName: "px-8",       // custom padding
})
// Result: .btn .btn-primary .btn-md px-8
// TwMerge ensures px-8 takes precedence over btn-md padding
```

**TwMerge Benefits:**
- ✅ No class conflicts
- ✅ Last class wins for same property
- ✅ Maintains specificity
- ✅ Optimized output

**Examples:**

```go
// Padding override
@Button(ButtonProps{
    Size: ButtonSizeMD,      // btn-md has padding: 1rem
    ClassName: "px-12",      // Override horizontal padding
})
// Result: px-12 wins

// Color override
@Button(ButtonProps{
    Variant: ButtonPrimary,  // Primary color
    ClassName: "bg-purple-500 hover:bg-purple-600",
})
// Result: Custom colors win

// Complete override
@Button(ButtonProps{
    ClassName: "bg-gradient-to-r from-purple-500 to-pink-500 text-white px-8 py-4",
})
// Result: Fully custom styling
```

---

## Accessibility

The button component is built with accessibility in mind following WCAG 2.1 Level AA guidelines:

### ✅ Implemented Features

#### Semantic HTML
Uses proper `<button>` element with correct ARIA attributes:

```html
<button type="button" role="button">Click Me</button>
```

#### Keyboard Navigation
Full keyboard support:
- **Tab** - Navigate to button
- **Shift + Tab** - Navigate backward
- **Enter** - Activate button
- **Space** - Activate button

#### Focus Management
Visible focus indicators:

```css
.btn:focus-visible {
    outline: 2px solid hsl(var(--ring));
    outline-offset: 2px;
    ring: 2px;
}
```

#### Loading States
Proper aria-label for screen readers:

```go
@Button(ButtonProps{
    Loading: true,
    Text: "Saving...",
})
// Renders: aria-label="Loading..."
```

#### Disabled States
Uses proper `disabled` attribute:

```html
<button disabled aria-disabled="true">
    Cannot Submit
</button>
```

### 📝 Required Implementations

#### Icon-Only Buttons

**Always** add descriptive aria-labels:

```go
// ✅ Good: Has aria-label
@Button(ButtonProps{
    Icon: "trash",
    AriaLabel: "Delete item",
})

// ❌ Bad: Missing aria-label
@Button(ButtonProps{
    Icon: "trash",
})
```

#### Dynamic Content

For buttons with changing content:

```go
<div x-data="{ count: 0 }">
    @Button(ButtonProps{
        AlpineClick: "count++",
        AriaLabel: "Increment counter",
    }) {
        <span aria-live="polite">Count: <span x-text="count"></span></span>
    }
</div>
```

#### Confirmation Dialogs

Use `hx-confirm` or custom modals:

```go
// With HTMX confirmation
@Button(ButtonProps{
    Text: "Delete",
    Variant: ButtonDestructive,
    HXPost: "/delete/123",
    HXConfirm: "Are you sure you want to delete this item? This action cannot be undone.",
    AriaLabel: "Delete item - requires confirmation",
})
```

### WCAG Compliance Checklist

#### ✅ Color Contrast (Level AA)

All button variants meet minimum contrast ratios:

| Variant | Background | Foreground | Ratio | Status |
|---------|------------|------------|-------|--------|
| Primary | `--primary` | `--primary-foreground` | 4.5:1 | ✅ Pass |
| Secondary | `--secondary` | `--secondary-foreground` | 4.5:1 | ✅ Pass |
| Destructive | `--destructive` | `--destructive-foreground` | 4.5:1 | ✅ Pass |
| Outline | `--background` | `--foreground` | 7:1 | ✅ Pass AAA |
| Ghost | `--accent` | `--accent-foreground` | 4.5:1 | ✅ Pass |

#### ✅ Keyboard Accessible

- All buttons are keyboard navigable
- Focus indicators are clearly visible
- No keyboard traps
- Logical tab order

#### ✅ Screen Reader Support

- Proper button role
- Descriptive labels
- State announcements (loading, disabled)
- Live regions for dynamic content

#### ⚠️ Touch Target Size

Minimum touch target is 44x44px (WCAG Level AAA):

```go
// ✅ Good: Large enough for touch
@Button(ButtonProps{
    Size: ButtonSizeMD,  // 40px height - close to target
    ClassName: "min-h-[44px]",  // Ensure minimum
})

// ⚠️ Warning: Too small for touch
@Button(ButtonProps{
    Size: ButtonSizeXS,  // 28px height - below target
})
```

**Recommendation:** Use `ButtonSizeMD` or larger for touch interfaces.

### Screen Reader Testing

Test your buttons with screen readers:

**NVDA (Windows):**
- "Save button" - ✅ Good
- "Button" - ❌ Bad (no label)
- "Loading... button" - ✅ Good (loading state)

**VoiceOver (macOS):**
- "Save, button" - ✅ Good
- "Button, unlabeled" - ❌ Bad
- "Save, button, dimmed" - ✅ Good (disabled)

**JAWS (Windows):**
- "Save button" - ✅ Good
- "Button" - ❌ Bad

### Best Practices Summary

1. **Always provide labels** - Text or aria-label
2. **Use semantic HTML** - `<button>` not `<div>`
3. **Maintain focus order** - Logical tab sequence
4. **Test with keyboard only** - No mouse required
5. **Check color contrast** - 4.5:1 minimum
6. **Announce state changes** - Loading, disabled
7. **Size touch targets** - 44x44px minimum
8. **Test with screen readers** - NVDA, VoiceOver, JAWS

---

## Best Practices

### 1. Semantic Variant Usage

Match button variant to action importance:

```go
// ✅ Good: Primary for main action
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Save Changes",
    Type: "submit",
})

// ✅ Good: Destructive for dangerous actions
@Button(ButtonProps{
    Variant: ButtonDestructive,
    Text: "Delete Account",
    HXConfirm: "This cannot be undone",
})

// ✅ Good: Secondary for cancel/back
@Button(ButtonProps{
    Variant: ButtonSecondary,
    Text: "Cancel",
})

// ❌ Bad: Primary for destructive action
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Delete Everything",
})

// ❌ Bad: Multiple primary buttons
<div>
    @PrimaryButton(ButtonProps{Text: "Save"})
    @PrimaryButton(ButtonProps{Text: "Submit"})  // Confusing!
</div>
```

### 2. Icon Usage Guidelines

**✅ DO:**

```go
// Use universally understood icons
@Button(ButtonProps{
    IconLeft: "save",
    Text: "Save",
})

// Provide aria-label for icon-only
@Button(ButtonProps{
    Icon: "settings",
    AriaLabel: "Open settings",
})

// Use directional icons for navigation
@Button(ButtonProps{
    Text: "Next",
    IconRight: "arrow-right",
})
```

**❌ DON'T:**

```go
// Don't use obscure icons without text
@Button(ButtonProps{
    Icon: "triangle-hexagon-square",  // What does this mean?
})

// Don't forget aria-labels
@Button(ButtonProps{
    Icon: "trash",  // Screen readers won't know the purpose
})

// Don't use icons that contradict text
@Button(ButtonProps{
    IconLeft: "trash",
    Text: "Save",  // Confusing!
})
```

### 3. Loading States

**✅ DO:**

```go
// Update text during loading
<div x-data="{ saving: false }">
    @Button(ButtonProps{
        Text: saving ? "Saving..." : "Save",
        Loading: saving,
        AlpineClick: "saving = true",
    })
</div>

// Disable during async operations
@Button(ButtonProps{
    Text: "Submit",
    HXPost: "/api/save",
    HXIndicator: "#spinner",
})

// Show progress indication
@Button(ButtonProps{
    Loading: true,
    Text: "Processing (2/5)",
})
```

**❌ DON'T:**

```go
// Don't allow duplicate submissions
@Button(ButtonProps{
    Text: "Pay $99.99",
    HXPost: "/api/payment",
    // Missing: Loading state or disabled logic
})

// Don't hide loading state
@Button(ButtonProps{
    Text: "Save",
    Loading: true,
    ClassName: "opacity-0",  // User can't see it's loading!
})
```

### 4. Size Hierarchy

Use sizes to establish visual hierarchy:

```go
// ✅ Good: Size matches importance
<div class="space-y-4">
    <h1>Welcome to Our Platform</h1>
    
    @Button(ButtonProps{
        Variant: ButtonPrimary,
        Size: ButtonSizeLG,
        Text: "Get Started",
    })
    
    @Button(ButtonProps{
        Variant: ButtonSecondary,
        Size: ButtonSizeMD,
        Text: "Learn More",
    })
</div>

// ❌ Bad: Inconsistent sizing
<div>
    @Button(ButtonProps{
        Size: ButtonSizeXL,
        Text: "Minor Action",  // Too large!
    })
</div>
```

**Size Usage Guidelines:**

| Size | Use Case | Example |
|------|----------|---------|
| **XL** | Hero CTAs, landing pages | "Start Free Trial" |
| **LG** | Primary page actions | "Save Changes", "Submit Form" |
| **MD** | Default actions, forms | "Add Item", "Search" |
| **SM** | Secondary actions, toolbars | "Edit", "Cancel" |
| **XS** | Dense UIs, table actions | Icon-only buttons |

### 5. Form Buttons

**✅ DO:**

```go
// Use proper types
<form hx-post="/api/save">
    <input type="text" name="title" />
    
    <div class="flex gap-2 mt-4">
        @Button(ButtonProps{
            Type: "submit",
            Variant: ButtonPrimary,
            Text: "Save",
        })
        
        @Button(ButtonProps{
            Type: "button",  // Prevents form submission
            Variant: ButtonSecondary,
            Text: "Cancel",
        })
    </div>
</form>

// Disable until valid
<form x-data="{ isValid: false }">
    <input 
        type="email" 
        @input="isValid = $el.checkValidity()"
    />
    
    @Button(ButtonProps{
        Type: "submit",
        ClassName: "x-bind:disabled='!isValid'",
    })
</form>
```

**❌ DON'T:**

```go
// Don't forget button types
<form>
    @Button(ButtonProps{
        Text: "Cancel",
        // Missing Type: "button" - will submit form!
    })
</form>

// Don't use div as button
<form>
    <div onclick="submitForm()">Submit</div>  // Not accessible!
</form>
```

### 6. HTMX Patterns

**✅ DO:**

```go
// Use appropriate HTTP methods
@Button(ButtonProps{
    HXPost: "/api/create",  // Creating
    Text: "Create",
})

@Button(ButtonProps{
    HXGet: "/api/items",  // Reading
    Text: "Load More",
})

@Button(ButtonProps{
    HXDelete: "/api/items/123",  // Deleting (if supported)
    Text: "Delete",
})

// Provide user feedback
@Button(ButtonProps{
    HXPost: "/api/save",
    HXTarget: "#result",
    HXSwap: "innerHTML",
    ClassName: "htmx-indicator",
})

// Use confirmation for destructive actions
@Button(ButtonProps{
    HXPost: "/api/delete-account",
    HXConfirm: "This will permanently delete your account. Are you absolutely sure?",
    Variant: ButtonDestructive,
})
```

**❌ DON'T:**

```go
// Don't use GET for mutations
@Button(ButtonProps{
    HXGet: "/api/delete/123",  // Should be POST/DELETE!
})

// Don't forget targets
@Button(ButtonProps{
    HXPost: "/api/save",
    // Missing HXTarget - where does response go?
})

// Don't ignore errors
@Button(ButtonProps{
    HXPost: "/api/risky-operation",
    // No error handling!
})
```

### 7. Alpine.js Integration

**✅ DO:**

```go
// Keep logic simple
@Button(ButtonProps{
    AlpineClick: "open = true",
})

// Use methods for complex logic
<div x-data="{ 
    async handleSubmit() {
        this.loading = true;
        try {
            await api.save(this.data);
            this.success = true;
        } catch (err) {
            this.error = err.message;
        } finally {
            this.loading = false;
        }
    }
}">
    @Button(ButtonProps{
        AlpineClick: "handleSubmit()",
    })
</div>

// Provide visual feedback
@Button(ButtonProps{
    AlpineClick: "count++",
}) {
    <span>Clicked: </span>
    <span x-text="count"></span>
}
```

**❌ DON'T:**

```go
// Don't put too much logic inline
@Button(ButtonProps{
    AlpineClick: "if (validate()) { loading = true; fetch('/api').then(r => r.json()).then(d => { result = d; loading = false; }); }",  // Too complex!
})

// Don't forget error handling
@Button(ButtonProps{
    AlpineClick: "await dangerousOperation()",  // What if it fails?
})
```

### 8. Accessibility First

**✅ DO:**

```go
// Always label icon-only buttons
@Button(ButtonProps{
    Icon: "close",
    AriaLabel: "Close dialog",
})

// Use semantic HTML
@Button(ButtonProps{
    Type: "button",
    Text: "Click Me",
})

// Announce dynamic changes
<div x-data="{ count: 0 }">
    @Button(ButtonProps{
        AlpineClick: "count++",
    }) {
        <span aria-live="polite">Count: <span x-text="count"></span></span>
    }
</div>

// Provide context
@Button(ButtonProps{
    Text: "Delete",
    Variant: ButtonDestructive,
    AriaLabel: "Delete invoice #12345",
})
```

**❌ DON'T:**

```go
// Don't use divs as buttons
<div onclick="doSomething()">Click</div>

// Don't forget labels
@Button(ButtonProps{
    Icon: "mysterious-icon",
    // No text, no aria-label - what does it do?
})

// Don't hide focus
@Button(ButtonProps{
    ClassName: "focus:outline-none",  // Bad for keyboard users!
})
```

### 9. Performance Optimization

**✅ DO:**

```go
// Throttle expensive operations
@Button(ButtonProps{
    HXPost: "/api/search",
    HXTrigger: "click throttle:1s",
})

// Debounce user input
@Button(ButtonProps{
    HXGet: "/api/autocomplete",
    HXTrigger: "keyup changed delay:300ms",
})

// Use proper swap strategies
@Button(ButtonProps{
    HXGet: "/api/items",
    HXSwap: "beforeend",  // Append instead of replace
})

// Optimize re-renders
<div x-data="{ items: [] }">
    <template x-for="item in items">
        @Button(ButtonProps{
            ClassName: "x-bind:key='item.id'",
        })
    </template>
</div>
```

**❌ DON'T:**

```go
// Don't spam the server
@Button(ButtonProps{
    HXPost: "/api/expensive-operation",
    // No throttling on expensive operation!
})

// Don't replace entire pages unnecessarily
@Button(ButtonProps{
    HXGet: "/data",
    HXTarget: "body",
    HXSwap: "outerHTML",  // Replaces everything!
})
```

### 10. Consistent Styling

**✅ DO:**

```go
// Use design system variants
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size: ButtonSizeMD,
})

// Extend with utility classes
@Button(ButtonProps{
    Variant: ButtonPrimary,
    ClassName: "w-full",  // Add responsive width
})

// Create consistent button groups
<div class="flex gap-2">
    @Button(ButtonProps{
        Variant: ButtonPrimary,
        Size: ButtonSizeMD,
    })
    @Button(ButtonProps{
        Variant: ButtonSecondary,
        Size: ButtonSizeMD,  // Consistent size
    })
</div>
```

**❌ DON'T:**

```go
// Don't fight the design system
@Button(ButtonProps{
    Variant: ButtonPrimary,
    ClassName: "bg-red-500",  // Overriding primary color
})

// Don't create one-off custom styles
@Button(ButtonProps{
    ClassName: "bg-[#ff00ff] text-[14.5px] px-[13px]",  // Random values
})

// Don't mix sizes randomly
<div>
    @Button(ButtonProps{Size: ButtonSizeXL})
    @Button(ButtonProps{Size: ButtonSizeXS})  // Inconsistent!
</div>
```

---

## Migration Guide

### From Legacy Button Components

If you're migrating from an older button implementation:

#### Before (Legacy)

```go
// Old API
@OldButton("primary", "large", "Save", "save-icon")

// Old inline styles
<button class="custom-btn custom-primary" onclick="save()">
    <i class="icon-save"></i>
    Save
</button>
```

#### After (New API)

```go
// New API - Props struct
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size: ButtonSizeLG,
    Text: "Save",
    IconLeft: "save",
})

// New API - Builder pattern
@Button(NewButton(
    WithPrimary(),
    WithSizeLG(),
    WithText("Save"),
    WithIconLeft("save"),
))

// New API - HTMX
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Text: "Save",
    IconLeft: "save",
    HXPost: "/api/save",
})
```

### Breaking Changes

#### 1. Prop Names Changed

```go
// ❌ Old
Size: "large"

// ✅ New
Size: ButtonSizeLG
```

#### 2. Icon Props Split

```go
// ❌ Old
Icon: "save"  // Could be left or right

// ✅ New
IconLeft: "save"   // Explicit positioning
IconRight: "arrow"
Icon: "close"      // Only for icon-only buttons
```

#### 3. Click Handlers

```go
// ❌ Old
OnClick: "handleClick()"

// ✅ New - Use HTMX or Alpine
HXPost: "/api/endpoint"
// OR
AlpineClick: "handleClick()"
```

#### 4. Default Variant

```go
// ❌ Old - No default
@Button(ButtonProps{Text: "Click"})  // Would be unstyled

// ✅ New - Primary is default
@Button(ButtonProps{Text: "Click"})  // Primary variant
```

### Migration Checklist

- [ ] Replace all old button imports
- [ ] Update prop names to new constants
- [ ] Split Icon props into IconLeft/IconRight
- [ ] Convert onClick to HXPost/HXGet or AlpineClick
- [ ] Add AriaLabel to icon-only buttons
- [ ] Update size constants (large → ButtonSizeLG)
- [ ] Test keyboard navigation
- [ ] Verify loading states work
- [ ] Check form submission buttons have Type prop
- [ ] Validate HTMX targets are set

---

## Troubleshooting

### Common Issues

#### 1. Button Not Clickable

**Symptoms:**
- Button appears but doesn't respond to clicks
- No hover effects

**Causes & Solutions:**

```go
// ❌ Problem: Disabled without visual indicator
@Button(ButtonProps{
    Disabled: true,
    ClassName: "opacity-100",  // Hiding disabled state
})

// ✅ Solution: Remove conflicting styles
@Button(ButtonProps{
    Disabled: true,
})

// ❌ Problem: Z-index issues
<div class="relative z-50">
    <div class="absolute inset-0"></div>  // Covering button
    @Button(ButtonProps{})
</div>

// ✅ Solution: Adjust z-index
<div class="relative">
    <div class="absolute inset-0 z-0"></div>
    @Button(ButtonProps{
        ClassName: "relative z-10",
    })
</div>
```

#### 2. HTMX Request Not Firing

**Symptoms:**
- Button click does nothing
- No network requests in DevTools

**Causes & Solutions:**

```go
// ❌ Problem: Missing HTMX script
// Solution: Add to <head>
<script src="https://unpkg.com/htmx.org@1.9.10"></script>

// ❌ Problem: Wrong HTTP method
@Button(ButtonProps{
    HXGet: "/api/save",  // Should be POST
})

// ✅ Solution: Use correct method
@Button(ButtonProps{
    HXPost: "/api/save",
})

// ❌ Problem: Invalid target selector
@Button(ButtonProps{
    HXPost: "/api/save",
    HXTarget: "result",  // Missing #
})

// ✅ Solution: Use valid selector
@Button(ButtonProps{
    HXPost: "/api/save",
    HXTarget: "#result",
})
```

#### 3. Icons Not Showing

**Symptoms:**
- Button renders but icon is missing
- Empty space where icon should be

**Causes & Solutions:**

```go
// ❌ Problem: Icon component not imported
// Solution: Ensure Icon component is available

// ❌ Problem: Wrong icon name
@Button(ButtonProps{
    IconLeft: "save-icon",  // Should be "save"
})

// ✅ Solution: Use correct icon name
@Button(ButtonProps{
    IconLeft: "save",
})

// ❌ Problem: Icon library not loaded
// Solution: Import icon library in layout
<link rel="stylesheet" href="icons.css">
```

#### 4. Loading Spinner Not Visible

**Symptoms:**
- Loading prop is true but no spinner appears
- Button text still visible during loading

**Causes & Solutions:**

```go
// ❌ Problem: Missing LoadingSpinner component
// Solution: Ensure component is defined

// ❌ Problem: Conflicting styles
@Button(ButtonProps{
    Loading: true,
    ClassName: "text-transparent",  // Hides spinner too
})

// ✅ Solution: Remove conflicting styles
@Button(ButtonProps{
    Loading: true,
})

// Check spinner CSS is loaded
.spinner {
    animation: spin 1s linear infinite;
}
```

#### 5. Alpine.js Not Working

**Symptoms:**
- AlpineClick does nothing
- x-data not recognized

**Causes & Solutions:**

```go
// ❌ Problem: Alpine.js not loaded
// Solution: Add to <head>
<script defer src="https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"></script>

// ❌ Problem: Syntax error
@Button(ButtonProps{
    AlpineClick: "count+",  // Missing +
})

// ✅ Solution: Fix syntax
@Button(ButtonProps{
    AlpineClick: "count++",
})

// ❌ Problem: Variable not in scope
<div>
    @Button(ButtonProps{
        AlpineClick: "count++",  // count doesn't exist
    })
</div>

// ✅ Solution: Add x-data
<div x-data="{ count: 0 }">
    @Button(ButtonProps{
        AlpineClick: "count++",
    })
</div>
```

#### 6. Form Not Submitting

**Symptoms:**
- Submit button doesn't submit form
- Page doesn't reload/update

**Causes & Solutions:**

```go
// ❌ Problem: Wrong button type
@Button(ButtonProps{
    Text: "Submit",
    // Missing Type: "submit"
})

// ✅ Solution: Set type
@Button(ButtonProps{
    Text: "Submit",
    Type: "submit",
})

// ❌ Problem: Button outside form
<div>
    <form id="myform">
        <input type="text" />
    </form>
    
    @Button(ButtonProps{
        Type: "submit",  // Not in form!
    })
</div>

// ✅ Solution: Move inside or use form attribute
<form id="myform">
    <input type="text" />
    
    @Button(ButtonProps{
        Type: "submit",
    })
</form>
```

#### 7. Styles Not Applied

**Symptoms:**
- Button appears unstyled
- Custom classes ignored

**Causes & Solutions:**

```go
// ❌ Problem: CSS not loaded
// Solution: Import button CSS
@import "button.css";

// ❌ Problem: TwMerge conflict
@Button(ButtonProps{
    Variant: ButtonPrimary,
    ClassName: "btn-primary",  // Duplicate
})

// ✅ Solution: Let variant handle it
@Button(ButtonProps{
    Variant: ButtonPrimary,
})

// ❌ Problem: Specificity issue
.my-custom-class {
    background: blue;  // Not specific enough
}

// ✅ Solution: Increase specificity
.btn.my-custom-class {
    background: blue;
}
```

#### 8. Accessibility Issues

**Symptoms:**
- Screen reader announces incorrectly
- Keyboard navigation doesn't work
- No focus indicator

**Causes & Solutions:**

```go
// ❌ Problem: Missing aria-label
@Button(ButtonProps{
    Icon: "trash",
})

// ✅ Solution: Add aria-label
@Button(ButtonProps{
    Icon: "trash",
    AriaLabel: "Delete item",
})

// ❌ Problem: Focus hidden
@Button(ButtonProps{
    ClassName: "focus:outline-none focus:ring-0",
})

// ✅ Solution: Keep focus visible
@Button(ButtonProps{
    // Default focus styles are accessible
})

// ❌ Problem: Using div instead of button
<div onclick="submit()">Submit</div>

// ✅ Solution: Use semantic button
@Button(ButtonProps{
    Text: "Submit",
    AlpineClick: "submit()",
})
```

### Debug Checklist

When troubleshooting button issues:

1. **Check Browser Console**
   - Look for JavaScript errors
   - Verify HTMX/Alpine loaded

2. **Inspect Element**
   - Verify HTML structure
   - Check applied CSS classes
   - Look for conflicting styles

3. **Network Tab**
   - Confirm HTMX requests firing
   - Check request/response
   - Verify endpoints

4. **Accessibility Tree**
   - Check button role
   - Verify aria-labels
   - Test keyboard navigation

5. **Component Props**
   - Print props to console
   - Verify correct types
   - Check for typos

### Getting Help

If you're still stuck:

1. **Check Documentation** - Review relevant sections above
2. **Search Issues** - Look for similar problems in the issue tracker
3. **Minimal Reproduction** - Create a minimal example that reproduces the issue
4. **Ask for Help** - Provide context, code, and error messages

---

## Summary

The Button component is a production-ready, accessible, and feature-rich component that supports:

✅ **Six variants** (Primary, Secondary, Destructive, Outline, Ghost, Link)  
✅ **Five sizes** (XS, SM, MD, LG, XL)  
✅ **Multiple icon positions** (Left, Right, Icon-only)  
✅ **State management** (Disabled, Loading)  
✅ **HTMX integration** (GET, POST, and all attributes)  
✅ **Alpine.js support** (Click handlers and data binding)  
✅ **Builder pattern API** (Fluent, chainable configuration)  
✅ **Reusable primitives** (LoadingSpinner component)  
✅ **Full accessibility** (WCAG 2.1 Level AA compliant)  
✅ **Design system integration** (CSS custom properties and compiled classes)  
✅ **TypeScript-like safety** (Via Go constants and builder pattern)

### Quick Reference

```go
// Basic button
@Button(ButtonProps{Text: "Click Me"})

// Styled button
@PrimaryButton(ButtonProps{
    Text: "Save",
    IconLeft: "save",
    Size: ButtonSizeLG,
})

// Interactive button
@Button(ButtonProps{
    Text: "Submit",
    Type: "submit",
    HXPost: "/api/save",
    HXTarget: "#result",
    Loading: isLoading,
})

// Builder pattern
@Button(NewButton(
    WithPrimary(),
    WithSizeLG(),
    WithText("Save"),
    WithIconLeft("save"),
    AsSubmit(),
    WithHXPost("/api/save"),
))
```

### Next Steps

1. **Implement the component** - Copy the implementation code
2. **Add to your design system** - Include required CSS
3. **Test accessibility** - Verify with keyboard and screen readers
4. **Create presets** - Build reusable button configurations
5. **Document usage** - Share patterns with your team

---

**Version:** 1.0.0  
**Last Updated:** 2024  
**Maintainer:** Your Team  
**License:** MIT

# *~~OLD DOC STARTING HERE~~*
# *Button Component*

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
