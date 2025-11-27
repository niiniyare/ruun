# Checkbox Component

A control that allows the user to toggle between checked and not checked.

## Basic Usage

```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

## CSS Classes

### Primary Classes
- **`input`** - Applied to the checkbox input element
- **`label`** - Applied to the label wrapper
- **`form`** - Parent container that provides automatic styling

### Tailwind Utilities Used
- `gap-3` - Spacing between checkbox and text
- `flex` - Flexbox layout for labels
- `items-start` - Vertical alignment
- `text-muted-foreground` - Muted text color
- `text-sm` - Small text size
- `leading-snug` - Tight line height

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "input" class | Yes |
| `type` | string | Must be "checkbox" | Yes |
| `id` | string | Unique identifier for label association | Recommended |
| `name` | string | Form field name for multiple checkboxes | No |
| `value` | string | Value when checked | No |
| `checked` | boolean | Pre-checked state | No |
| `disabled` | boolean | Disables the checkbox | No |

## Examples

### Basic Checkbox
```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

### Disabled Checkbox
```html
<label class="label gap-3">
  <input type="checkbox" class="input" disabled>
  Accept terms and conditions
</label>
```

### With Descriptive Text
```html
<div class="flex items-start gap-3">
  <input type="checkbox" id="checkbox-with-text" class="input">
  <div class="grid gap-2">
    <label for="checkbox-with-text" class="label">Accept terms and conditions</label>
    <p class="text-muted-foreground text-sm">By clicking this checkbox, you agree to the terms and conditions.</p>
  </div>
</div>
```

### Form Context (Auto-styling)
```html
<form class="form flex flex-row items-start gap-3 rounded-md border p-4 shadow-xs">
  <input type="checkbox" id="checkbox-form-1">
  <div class="flex flex-col gap-1">
    <label for="checkbox-form-1" class="leading-snug">Use different settings for my mobile devices</label>
    <p class="text-muted-foreground text-sm leading-snug">You can manage your mobile notifications in the mobile settings page.</p>
  </div>
</form>
```

### Multiple Checkboxes in Fieldset
```html
<form class="form flex flex-col gap-4">
  <fieldset id="demo-form-checkboxes" class="flex flex-col gap-2">
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="1" checked>
      Recents
    </label>
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="2" checked>
      Home
    </label>
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="3">
      Applications
    </label>
  </fieldset>
</form>
```

## States

### Default State
Regular unchecked appearance with focus styles

### Checked State
Shows checkmark when selected

### Disabled State
Reduced opacity and no interaction

### Focus State
Visible focus ring for keyboard navigation

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, checkboxes inherit styling automatically:

```html
<form class="form">
  <!-- No need for explicit classes on checkbox -->
  <input type="checkbox" id="auto-styled">
  <label for="auto-styled">Auto-styled checkbox</label>
</form>
```

### Multiple Selection
Use the same `name` attribute for related checkboxes:

```html
<fieldset>
  <legend>Select your preferences:</legend>
  <label><input type="checkbox" name="preferences" value="email"> Email notifications</label>
  <label><input type="checkbox" name="preferences" value="sms"> SMS notifications</label>
  <label><input type="checkbox" name="preferences" value="push"> Push notifications</label>
</fieldset>
```

## JavaScript Integration

- **HTMX Support**: Compatible with HTMX attributes for dynamic behavior
- **Form Validation**: Works with native and custom validation
- **Theme Support**: Automatically adapts to light/dark themes

## Accessibility Features

- **Keyboard Navigation**: Full keyboard support with Tab and Space
- **Screen Reader**: Proper labeling and state announcement
- **Focus Management**: Visible focus indicators
- **Form Association**: Proper label-input relationships

## Best Practices

1. Always associate with a label using `for` attribute or wrapping
2. Use `fieldset` and `legend` for related checkbox groups
3. Provide descriptive text for complex choices
4. Consider the reading order for screen readers
5. Use consistent naming conventions for grouped checkboxes

## Related Components

- [Radio Group](./radio-group.md) - For mutually exclusive options
- [Switch](./switch.md) - For on/off toggles
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling