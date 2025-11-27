# Badge Component

Displays a badge or a component that looks like a badge.

## Basic Usage

```html
<span class="badge">Badge</span>
```

## CSS Classes

### Variant Classes
- **`badge`** - Primary badge (default styling)
- **`badge-primary`** - Primary variant
- **`badge-secondary`** - Secondary variant  
- **`badge-destructive`** - Destructive/error variant
- **`badge-outline`** - Outline variant

### Tailwind Utilities Used
- `inline-flex` - Inline flex display
- `items-center` - Center items vertically
- `rounded-full` - Full border radius for circular badges
- `px-2.5` - Horizontal padding
- `py-0.5` - Vertical padding
- `text-xs` - Extra small text
- `font-medium` - Medium font weight
- `border` - Border styling (outline variant)

### Modifier Classes
- `rounded-full` - Makes badge circular
- `h-5 min-w-5` - Fixed height and minimum width for counters
- `px-1` - Reduced padding for small badges
- `font-mono tabular-nums` - Monospace font for numbers

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include badge variant class | Yes |

## Examples

### Basic Variants
```html
<span class="badge">Badge</span>
<span class="badge-secondary">Secondary</span>
<span class="badge-destructive">Destructive</span>
<span class="badge-outline">Outline</span>
```

### With Custom Colors
```html
<span class="badge-secondary bg-blue-500 text-white dark:bg-blue-600">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  Verified
</span>
```

### Counter Badges
```html
<span class="badge rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">8</span>
<span class="badge-destructive rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">99</span>
<span class="badge-outline rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">20+</span>
```

### With Icons
```html
<span class="badge-destructive">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <line x1="12" x2="12" y1="8" y2="12" />
    <line x1="12" x2="12.01" y1="16" y2="16" />
  </svg>
  With icon
</span>
```

### As Links
```html
<a href="#" class="badge-outline">
  Link
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M5 12h14" />
    <path d="m12 5 7 7-7 7" />
  </svg>
</a>
```

### Layout Examples
```html
<div class="flex flex-col items-center gap-2">
  <div class="flex w-full flex-wrap gap-2">
    <span class="badge">Badge</span>
    <span class="badge-secondary">Secondary</span>
    <span class="badge-destructive">Destructive</span>
    <span class="badge-outline">Outline</span>
  </div>
</div>
```

## Color Variants

### Primary
- Background: Primary theme color
- Text: Primary foreground color
- Use: Default/primary actions

### Secondary  
- Background: Secondary theme color
- Text: Secondary foreground color
- Use: Secondary information

### Destructive
- Background: Error/danger color
- Text: Error foreground color  
- Use: Errors, warnings, deletions

### Outline
- Background: Transparent
- Border: Border color
- Text: Foreground color
- Use: Subtle emphasis

## JavaScript Integration

- **Copy functionality**: Code examples include clipboard copying
- **Theme support**: Adapts to light/dark themes automatically
- **Interactive states**: Can be used as clickable elements (links, buttons)

## Accessibility Features

- Semantic HTML structure
- Proper contrast ratios
- Screen reader friendly
- Focus states when interactive

## Best Practices

1. Use meaningful text or numbers
2. Consider accessibility for color-blind users
3. Use appropriate variant for context
4. Keep text concise for small badges
5. Use `font-mono tabular-nums` for numeric badges

## Related Components

- [Button](./button.md) - For interactive actions
- [Avatar](./avatar.md) - Often used with status badges
- [Alert](./alert.md) - For larger status messages