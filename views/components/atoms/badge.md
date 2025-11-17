# Badge Component

A versatile badge component for displaying labels, tags, status indicators, and notification counts. Supports both inline badges and positioned overlay badges with full HTMX and Alpine.js integration.

## Table of Contents

- [Basic Usage](#basic-usage)
- [Variants](#variants)
- [Sizes](#sizes)
- [Modes](#modes)
- [Icons](#icons)
- [Overlay Badges](#overlay-badges)
- [Interactive Badges](#interactive-badges)
- [API Patterns](#api-patterns)
- [JSON Schema](#json-schema)
- [Examples](#examples)
- [API Reference](#api-reference)

## Basic Usage

### Simple Badge

```go
// Basic badge with text
@Badge(BadgeProps{}) { New }

// Success badge
@Badge(BadgeProps{Variant: BadgeSuccess}) { Active }

// With custom class
@Badge(BadgeProps{Class: "ml-2"}) { Beta }
```

### Using Convenience Components

```go
@DefaultBadge(BadgeProps{}) { Default }
@SuccessBadge(BadgeProps{}) { Success }
@WarningBadge(BadgeProps{}) { Warning }
@DestructiveBadge(BadgeProps{}) { Error }
```

## Variants

Six visual variants following your design system:

```go
// Default (Primary)
@Badge(BadgeProps{Variant: BadgeDefault}) { Default }

// Secondary
@Badge(BadgeProps{Variant: BadgeSecondary}) { Secondary }

// Success (Green)
@Badge(BadgeProps{Variant: BadgeSuccess}) { Success }

// Warning (Yellow)
@Badge(BadgeProps{Variant: BadgeWarning}) { Warning }

// Destructive (Red)
@Badge(BadgeProps{Variant: BadgeDestructive}) { Error }

// Outline
@Badge(BadgeProps{Variant: BadgeOutline}) { Outline }
```

## Sizes

Three size options:

```go
// Small
@Badge(BadgeProps{Size: BadgeSizeSM}) { Small }

// Medium (default)
@Badge(BadgeProps{Size: BadgeSizeMD}) { Medium }

// Large
@Badge(BadgeProps{Size: BadgeSizeLG}) { Large }
```

## Modes

Special display modes for overlay badges:

### Text Mode (Default)

Standard badge with text content:

```go
@Badge(BadgeProps{
    Mode: BadgeModeText,
    Text: 5,
    Variant: BadgeDestructive,
})
```

### Dot Mode

Small status indicator dot:

```go
@Badge(BadgeProps{
    Mode: BadgeModeDot,
    Variant: BadgeSuccess,
})
```

### Ribbon Mode

Wider badge for labels:

```go
@Badge(BadgeProps{
    Mode: BadgeModeRibbon,
    Text: "Premium",
    Variant: BadgeWarning,
})
```

## Icons

Add icons to badges:

```go
// Left icon
@Badge(BadgeProps{IconLeft: "check"}) { Verified }

// Right icon
@Badge(BadgeProps{IconRight: "arrow-right"}) { Next }

// Both icons
@Badge(BadgeProps{IconLeft: "star", IconRight: "x"}) { Featured }

// Icon only (no text)
@Badge(BadgeProps{Icon: "bell"})
```

## Overlay Badges

Position badges on top of other elements:

### Basic Overlay

```go
@BadgeWrapper(
    buttonComponent("Messages"),
    BadgeProps{
        Text: 5,
        Mode: BadgeModeText,
        Position: BadgePositionTopRight,
        Variant: BadgeDestructive,
    },
)
```

### Positions

Four corner positions available:

```go
// Top Right (default)
BadgePositionTopRight

// Top Left
BadgePositionTopLeft

// Bottom Right
BadgePositionBottomRight

// Bottom Left
BadgePositionBottomLeft
```

### With Custom Offset

```go
@BadgeWrapper(
    avatarComponent(),
    BadgeProps{
        Text: "NEW",
        Position: BadgePositionTopRight,
        Offset: []interface{}{10, -5}, // [x, y] in pixels
        Variant: BadgeSuccess,
    },
)
```

### Status Dot Example

```go
@BadgeWrapper(
    avatarComponent(),
    BadgeProps{
        Mode: BadgeModeDot,
        Position: BadgePositionBottomRight,
        Variant: BadgeSuccess, // Online status
    },
)
```

### Overflow Count

Display counts with overflow handling:

```go
@BadgeWrapper(
    iconComponent("bell"),
    BadgeProps{
        Text: 150,
        OverflowCount: 99, // Shows "99+"
        Position: BadgePositionTopRight,
        Variant: BadgeDestructive,
    },
)
```

## Interactive Badges

### Removable Badges

```go
@Badge(BadgeProps{
    Variant: BadgeSecondary,
    Removable: true,
    OnRemove: "$el.remove()", // Alpine.js handler
}) { Removable Tag }
```

### With Animation

```go
@Badge(BadgeProps{
    Text: 3,
    Animation: true, // Pulse animation
    Variant: BadgeDestructive,
})
```

### Conditional Visibility

```go
@Badge(BadgeProps{
    Text: 5,
    VisibleOn: "unreadCount > 0", // Alpine.js expression
    Variant: BadgeWarning,
})
```

### HTMX Integration

```go
@Badge(BadgeProps{
    Variant: BadgeWarning,
    HXPost: "/api/dismiss",
    HXTarget: "#notifications",
    HXSwap: "outerHTML",
    HXTrigger: "click",
}) { Dismiss }
```

### Alpine.js Click Handler

```go
@Badge(BadgeProps{
    Variant: BadgeDefault,
    AlpineClick: "handleBadgeClick($event)",
}) { Click Me }
```

## API Patterns

### 1. Direct Props (Standard)

```go
@Badge(BadgeProps{
    Variant: BadgeSuccess,
    Size: BadgeSizeSM,
    IconLeft: "check",
}) { Verified }
```

### 2. Fluent API (Builder Pattern)

Chain methods for readable, intuitive syntax:

```go
badge := NewBadge().
    Success().
    Small().
    WithIconLeft("check").
    Build()

@Badge(badge) { Verified }
```

**All Builder Methods:**

```go
// Variants
.Default()
.Secondary()
.Success()
.Warning()
.Destructive()
.Outline()

// Sizes
.Small()
.Medium()
.Large()

// Content
.WithText(text)
.WithIcon(icon)
.WithIconLeft(icon)
.WithIconRight(icon)
.WithClass(class)
.WithID(id)
.WithStyle(style)

// Modes
.AsText()
.AsDot()
.AsRibbon()

// Positions
.TopRight()
.TopLeft()
.BottomRight()
.BottomLeft()
.WithOffset(x, y)
.WithOverflowCount(count)

// Interactivity
.Removable(onRemove)
.Animated()
.VisibleWhen(condition)

// HTMX
.HXPost(url)
.HXGet(url)
.HXTarget(target)
.HXSwap(swap)
.HXTrigger(trigger)

// Alpine.js
.OnClick(handler)

// Build
.Build() // Returns BadgeProps
```

**Complex Example:**

```go
notificationBadge := NewBadge().
    Destructive().
    AsText().
    TopRight().
    WithText(150).
    WithOverflowCount(99).
    Animated().
    VisibleWhen("notifications > 0").
    Build()

@BadgeWrapper(bellIcon, notificationBadge)
```

### 3. Functional Options Pattern

Composable and reusable:

```go
@Badge(NewBadgeWithOptions(
    WithSuccess(),
    WithSmall(),
    WithIconLeft("check"),
)) { Verified }
```

**Reusable Configurations:**

```go
var (
    SuccessTagOptions = []BadgeOption{
        WithSuccess(),
        WithIconLeft("check"),
        WithSmall(),
    }
    
    NotificationBadgeOptions = []BadgeOption{
        WithDestructive(),
        AsTextMode(),
        AtTopRight(),
        WithOverflowCount(99),
        WithAnimation(),
    }
)

// Use predefined options
@Badge(NewBadgeWithOptions(SuccessTagOptions...)) { Active }

// Extend with more options
@Badge(NewBadgeWithOptions(
    append(SuccessTagOptions, AsRemovable("remove()"))...,
)) { Tag }
```

**All Option Functions:**

```go
// Variants
WithVariant(variant), WithDefault(), WithSecondary(), WithSuccess(), 
WithWarning(), WithDestructive(), WithOutline()

// Sizes
WithSize(size), WithSmall(), WithMedium(), WithLarge()

// Content
WithText(text), WithIcon(icon), WithIconLeft(icon), WithIconRight(icon),
WithClass(class), WithID(id), WithStyle(style)

// Modes
WithMode(mode), AsTextMode(), AsDotMode(), AsRibbonMode()

// Positions
WithPosition(position), AtTopRight(), AtTopLeft(), AtBottomRight(), 
AtBottomLeft(), WithOffset(x, y), WithOverflowCount(count)

// Interactivity
AsRemovable(onRemove), WithAnimation(), VisibleWhen(condition)

// HTMX
WithHXPost(url), WithHXGet(url), WithHXTarget(target), 
WithHXSwap(swap), WithHXTrigger(trigger)

// Alpine.js
WithAlpineClick(handler)
```

## JSON Schema

Load badges dynamically from JSON:

### Basic JSON Badge

```go
jsonBadge := `{
    "variant": "success",
    "size": "sm",
    "text": "New",
    "iconLeft": "star"
}`

@BadgeFrom(jsonBadge)
```

### JSON Schema Structure

```json
{
    "variant": "success",        // default, secondary, success, warning, destructive, outline
    "size": "md",                // sm, md, lg
    "class": "ml-2",
    "text": "Badge Text",        // string or number
    "icon": "bell",              // single icon (icon-only badge)
    "iconLeft": "check",         // left icon
    "iconRight": "arrow",        // right icon
    "mode": "text",              // text, dot, ribbon
    "position": "top-right",     // top-right, top-left, bottom-right, bottom-left
    "offset": [10, -5],          // [x, y] position offset
    "overflowCount": 99,         // max count before showing "+"
    "visibleOn": "count > 0",    // Alpine.js visibility condition
    "animation": true,           // enable pulse animation
    "style": "opacity: 0.8",     // custom inline styles
    "removable": true            // show remove button
}
```

### Predefined JSON Examples

```go
// Notification badge
@BadgeFrom(ExampleNotificationBadge)

// Overflow count badge
@BadgeFrom(ExampleOverflowBadge)

// Status dot
@BadgeFrom(ExampleStatusDot)

// Dynamic visibility
@BadgeFrom(ExampleDynamicBadge)

// Custom offset
@BadgeFrom(ExampleCustomOffsetBadge)
```

### Creating Props from JSON

```go
props, err := BadgeFromJSON(jsonString)
if err != nil {
    // Handle error
}

@Badge(props) { Content }
```

## Examples

### Status Tags

```go
<div class="flex gap-2">
    @Badge(BadgeProps{Variant: BadgeSuccess, IconLeft: "check"}) { Active }
    @Badge(BadgeProps{Variant: BadgeWarning, IconLeft: "clock"}) { Pending }
    @Badge(BadgeProps{Variant: BadgeDestructive, IconLeft: "x"}) { Inactive }
</div>
```

### Notification Count

```go
@BadgeWrapper(
    Button(ButtonProps{}) { <i class="icon-bell"></i> },
    BadgeProps{
        Text: 5,
        Mode: BadgeModeText,
        Position: BadgePositionTopRight,
        Variant: BadgeDestructive,
        Animation: true,
    },
)
```

### User Status Indicator

```go
@BadgeWrapper(
    Avatar(AvatarProps{Src: user.Avatar}),
    BadgeProps{
        Mode: BadgeModeDot,
        Position: BadgePositionBottomRight,
        Variant: if user.Online { BadgeSuccess } else { BadgeSecondary },
    },
)
```

### Removable Filter Tags

```go
<div class="flex flex-wrap gap-2" x-data="{ tags: ['Design', 'Development', 'Marketing'] }">
    <template x-for="tag in tags" :key="tag">
        @Badge(BadgeProps{
            Variant: BadgeSecondary,
            Size: BadgeSizeSM,
            Removable: true,
            OnRemove: "tags = tags.filter(t => t !== tag)",
        }) {
            <span x-text="tag"></span>
        }
    </template>
</div>
```

### Interactive HTMX Badge

```go
@Badge(BadgeProps{
    Variant: BadgeWarning,
    IconLeft: "trash",
    HXPost: "/api/clear-notifications",
    HXTarget: "#notification-area",
    HXSwap: "innerHTML",
    HXTrigger: "click",
}) { Clear All }
```

### Count with Overflow

```go
badge := NewBadge().
    Destructive().
    AsText().
    TopRight().
    WithText(notificationCount).
    WithOverflowCount(99).
    Animated().
    Build()

@BadgeWrapper(mailIcon, badge)
// If count is 150, displays "99+"
```

### Conditional Badge

```go
<div x-data="{ unread: 5 }">
    @BadgeWrapper(
        Button(ButtonProps{}) { Messages },
        BadgeProps{
            Text: nil, // Text set via Alpine
            Mode: BadgeModeText,
            Position: BadgePositionTopRight,
            Variant: BadgeDestructive,
            VisibleOn: "unread > 0",
        },
    )
    <span x-show="unread > 0" class="absolute top-0 right-0" x-text="unread"></span>
</div>
```

### Using Fluent API in Template

```go
templ NotificationBell(count int) {
    @BadgeWrapper(
        Icon(IconProps{Name: "bell"}),
        NewBadge().
            Destructive().
            AsText().
            TopRight().
            WithText(count).
            WithOverflowCount(99).
            Animated().
            VisibleWhen("count > 0").
            Build(),
    )
}
```

### Product Labels

```go
<div class="relative">
    <img src="product.jpg" />
    @Badge(NewBadgeWithOptions(
        WithDestructive(),
        AsRibbonMode(),
        WithText("Sale"),
        AtTopLeft(),
        WithOffset(-5, -5),
    ))
    @Badge(NewBadgeWithOptions(
        WithSuccess(),
        WithText("New"),
        AtTopRight(),
    ))
</div>
```

## API Reference

### BadgeProps

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `Variant` | `BadgeVariant` | `BadgeDefault` | Visual style variant |
| `Size` | `BadgeSize` | `BadgeSizeMD` | Badge size |
| `Class` | `string` | `""` | Additional CSS classes |
| `ID` | `string` | `""` | HTML id attribute |
| `Text` | `interface{}` | `nil` | Text content (string or number) |
| `Icon` | `string` | `""` | Single icon (icon-only) |
| `IconLeft` | `string` | `""` | Left icon |
| `IconRight` | `string` | `""` | Right icon |
| `Mode` | `BadgeMode` | `BadgeModeText` | Display mode |
| `Position` | `BadgePosition` | `BadgePositionTopRight` | Corner position (overlay) |
| `Offset` | `[]interface{}` | `nil` | Position offset [x, y] |
| `OverflowCount` | `int` | `0` | Max count before "+" |
| `Removable` | `bool` | `false` | Show remove button |
| `OnRemove` | `string` | `""` | Alpine.js remove handler |
| `Animation` | `bool` | `false` | Enable pulse animation |
| `VisibleOn` | `string` | `""` | Alpine.js visibility condition |
| `HXPost` | `string` | `""` | HTMX POST URL |
| `HXGet` | `string` | `""` | HTMX GET URL |
| `HXTarget` | `string` | `""` | HTMX target selector |
| `HXSwap` | `string` | `""` | HTMX swap strategy |
| `HXTrigger` | `string` | `""` | HTMX trigger event |
| `AlpineClick` | `string` | `""` | Alpine.js click handler |
| `Style` | `string` | `""` | Custom inline styles |

### BadgeVariant

```go
const (
    BadgeDefault     BadgeVariant = "default"
    BadgeSecondary   BadgeVariant = "secondary"
    BadgeSuccess     BadgeVariant = "success"
    BadgeWarning     BadgeVariant = "warning"
    BadgeDestructive BadgeVariant = "destructive"
    BadgeOutline     BadgeVariant = "outline"
)
```

### BadgeSize

```go
const (
    BadgeSizeSM BadgeSize = "sm"
    BadgeSizeMD BadgeSize = "md"
    BadgeSizeLG BadgeSize = "lg"
)
```

### BadgeMode

```go
const (
    BadgeModeText   BadgeMode = "text"
    BadgeModeDot    BadgeMode = "dot"
    BadgeModeRibbon BadgeMode = "ribbon"
)
```

### BadgePosition

```go
const (
    BadgePositionTopRight    BadgePosition = "top-right"
    BadgePositionTopLeft     BadgePosition = "top-left"
    BadgePositionBottomRight BadgePosition = "bottom-right"
    BadgePositionBottomLeft  BadgePosition = "bottom-left"
)
```

### Components

#### Badge

Main badge component:

```go
templ Badge(props BadgeProps, children ...templ.Component)
```

#### BadgeWrapper

Wraps content with an overlay badge:

```go
templ BadgeWrapper(content templ.Component, badge BadgeProps)
```

#### Convenience Components

```go
templ DefaultBadge(props BadgeProps, children ...templ.Component)
templ SuccessBadge(props BadgeProps, children ...templ.Component)
templ WarningBadge(props BadgeProps, children ...templ.Component)
templ DestructiveBadge(props BadgeProps, children ...templ.Component)
```

#### JSON Components

```go
templ BadgeFrom(schemaJSON string)
```

### Helper Functions

```go
// Fluent API
func NewBadge() *BadgeBuilder

// Functional Options
func NewBadgeWithOptions(opts ...BadgeOption) BadgeProps

// JSON Parsing
func BadgeFromJSON(jsonStr string) (BadgeProps, error)
```

## Best Practices

1. **Use semantic variants** - Match badge variant to meaning (success for positive, destructive for errors)
2. **Keep text concise** - Badges work best with 1-3 words
3. **Use overflow counts** - For notifications, set a reasonable overflow limit (e.g., 99)
4. **Prefer dot mode for status** - Use dots for simple online/offline indicators
5. **Choose the right pattern** - Use fluent API for complex badges, direct props for simple ones
6. **Leverage JSON** - Use JSON schema for badges loaded from backend/database
7. **Accessible removables** - Always include proper `aria-label` on remove buttons (handled automatically)
8. **Test animations** - Pulse animations draw attention - use sparingly

## Styling Customization

The badge uses design system tokens that can be customized in your Tailwind configuration:

```css
/* Customize in your tailwind.config.js */
{
  theme: {
    extend: {
      colors: {
        primary: { ... },
        secondary: { ... },
        success: { ... },
        warning: { ... },
        destructive: { ... },
      }
    }
  }
}
```

For one-off customization:

```go
@Badge(BadgeProps{
    Variant: BadgeDefault,
    Class: "bg-purple-500 text-white hover:bg-purple-600",
}) { Custom }
```
