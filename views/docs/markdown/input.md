# Input Component

Displays a form input field or a component that looks like an input field.

## Basic Usage

```html
<input class="input" type="email" placeholder="Email">
```

## CSS Classes

### Primary Class
- **`input`** - Main input styling class

### Tailwind Utilities Used
- `w-full` - Full width
- `rounded-md` - Medium border radius
- `border` - Border styling
- `bg-input` - Input background color
- `px-3` - Horizontal padding
- `py-2` - Vertical padding
- `text-sm` - Small text size
- `transition-colors` - Color transition effect
- `focus:outline-none` - Remove default focus outline
- `focus:ring-2` - Focus ring size
- `focus:ring-offset-2` - Focus ring offset

### State Classes
- **Invalid state**: `aria-invalid="true"` - Changes border to error color
- **Disabled state**: `disabled` - Applies muted styling and cursor

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "input" class | Yes |
| `type` | string | Input type (text, email, password, etc.) | No |
| `placeholder` | string | Placeholder text | No |
| `disabled` | boolean | Disables the input | No |
| `aria-invalid` | boolean | Marks input as invalid for error states | No |
| `id` | string | Unique identifier for label association | No |

## Examples

### Default Input
```html
<input class="input" type="email" placeholder="Email">
```

### Invalid Input
```html
<input class="input" type="email" placeholder="Email" aria-invalid="true">
```

### Disabled Input
```html
<input class="input" type="email" placeholder="Email" disabled>
```

### With Label
```html
<div class="grid gap-3">
  <label for="input-with-label" class="label">Label</label>
  <input class="input" id="input-with-label" type="email" placeholder="Email">
</div>
```

### With Help Text
```html
<div class="grid gap-3">
  <label for="input-with-text" class="label">Label</label>
  <input class="input" id="input-with-text" type="email" placeholder="Email">
  <p class="text-muted-foreground text-sm">Fill in your email address.</p>
</div>
```

### With Button
```html
<div class="flex items-center space-x-2">
  <input class="input" type="email" placeholder="Email">
  <button type="submit" class="btn">Submit</button>
</div>
```

### Form Context
```html
<form class="form space-y-6 w-full">
  <div class="grid gap-3">
    <label for="input-form" class="label">Username</label>
    <input class="input" id="input-form" type="text" placeholder="hunvreus">
    <p class="text-muted-foreground text-sm">This is your public display name.</p>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

## JavaScript Integration

- **Copy functionality**: Uses navigator.clipboard.writeText() for code copying
- **HTMX Support**: Can be used with HTMX attributes for dynamic behavior
- **Theme support**: Automatically adapts to light/dark themes

## Accessibility Features

- Supports `aria-invalid` for screen readers
- Works with `for` attribute on labels
- Focus management with visible focus rings
- Proper color contrast in all states

## Related Components

- [Label](./label.md) - For accessible labeling
- [Form](./form.md) - For form context styling
- [Button](./button.md) - Often used together