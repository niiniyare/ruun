# Label Component

Renders an accessible label associated with controls.

## Basic Usage

```html
<label class="label" for="email">Your email address</label>
```

## CSS Classes

### Primary Class
- **`label`** - Main label styling class

### Tailwind Utilities Used
- `text-sm` - Small text size
- `font-medium` - Medium font weight
- `leading-none` - No line height
- `peer-disabled:cursor-not-allowed` - Disabled cursor when peer is disabled
- `peer-disabled:opacity-70` - Reduced opacity when peer is disabled

### Layout Classes
- `gap-3` - For spacing between label and control when used as flex container

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "label" class | Yes |
| `for` | string | ID of associated form control | Recommended |

## Examples

### Basic Label
```html
<label class="label" for="email">Your email address</label>
```

### Label with Input
```html
<div class="grid gap-3">
  <label class="label" for="email">Email</label>
  <input class="input" id="email" type="email" placeholder="Email">
</div>
```

### Inline Label with Control
```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

### Label with Disabled Input
```html
<div class="grid gap-3">
  <label class="label" for="email">Email</label>
  <input class="input" id="email" type="email" placeholder="Email" disabled>
</div>
```

## State Behavior

### Disabled State
When the associated input is disabled, the label automatically:
- Shows a "not-allowed" cursor
- Reduces opacity to 70%
- Maintains accessibility

This is handled through CSS peer selectors:
```css
.peer-disabled:cursor-not-allowed
.peer-disabled:opacity-70
```

### Form Context
Labels automatically inherit proper styling when used within a `.form` container, eliminating the need to manually add the `label` class.

## Layout Patterns

### Vertical Layout (Recommended)
```html
<div class="grid gap-3">
  <label class="label" for="input-field">Field Label</label>
  <input class="input" id="input-field" type="text">
</div>
```

### Horizontal Layout  
```html
<label class="label gap-3 flex items-center">
  <input type="checkbox" class="input">
  <span>Checkbox Label</span>
</label>
```

## JavaScript Integration

- **Theme support**: Automatically adapts to light/dark themes
- **Form validation**: Works with validation libraries
- **HTMX support**: Compatible with dynamic form updates

## Accessibility Features

- **Semantic association**: Uses `for` attribute to associate with controls
- **Screen reader support**: Properly announced by assistive technology  
- **Focus management**: Clicking label focuses associated control
- **Disabled state**: Visual and interaction changes when control is disabled
- **Form context**: Maintains accessibility within form groups

## Best Practices

1. Always use `for` attribute with unique ID of associated control
2. Keep label text clear and concise  
3. Place labels above inputs in vertical layouts
4. Use inline labels only for checkboxes and radio buttons
5. Ensure sufficient color contrast
6. Test with screen readers

## Form Integration

Labels work seamlessly with the form component system:

```html
<form class="form space-y-6">
  <div class="grid gap-3">
    <label for="username">Username</label> <!-- Inherits label styling -->
    <input id="username" type="text" placeholder="Enter username">
  </div>
</form>
```

## Related Components

- [Input](./input.md) - Primary form control
- [Form](./form.md) - Form container context
- [Checkbox](./checkbox.md) - Checkbox inputs
- [Radio Group](./radio-group.md) - Radio button groups
- [Textarea](./textarea.md) - Multi-line text inputs