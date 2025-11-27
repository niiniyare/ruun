# Select Component

Displays a list of options for the user to pick from—triggered by a button.

## Basic Usage

```html
<select class="select w-[180px]">
  <option>Apple</option>
  <option>Banana</option>
  <option>Orange</option>
</select>
```

## CSS Classes

### Primary Classes
- **`select`** - Applied to select element or custom select wrapper
- **`btn-outline`** - Button styling for custom select trigger
- **`justify-between`** - Space between text and chevron icon
- **`font-normal`** - Normal font weight

### Tailwind Utilities Used
- `w-[180px]` - Fixed width (customizable)
- `truncate` - Text overflow handling
- `text-muted-foreground` - Muted text color
- `opacity-50` - Icon opacity
- `shrink-0` - Prevent icon shrinking

## Component Types

### Native Select
Uses the browser's native select element with custom styling.

### Custom Select
Enhanced select with JavaScript functionality and custom appearance.

## Examples

### Native Select
```html
<select class="select w-[180px]">
  <optgroup label="Fruits">
    <option>Apple</option>
    <option>Banana</option>
    <option>Blueberry</option>
    <option>Grapes</option>
    <option>Orange</option>
  </optgroup>
  <optgroup label="Vegetables">
    <option>Aubergine</option>
    <option>Broccoli</option>
    <option>Carrot</option>
    <option selected>Courgette</option>
  </optgroup>
</select>
```

### Custom Select (JavaScript-Enhanced)
```html
<div id="select-942316" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[180px]" 
          id="select-942316-trigger" 
          aria-haspopup="listbox" 
          aria-expanded="false" 
          aria-controls="select-942316-listbox">
    <span class="truncate">Apple</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" 
         fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" 
         class="lucide lucide-chevrons-up-down text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="select-942316-popover" data-popover aria-hidden="true">
    <!-- Options list -->
    <div role="listbox" id="select-942316-listbox" aria-labelledby="select-942316-trigger">
      <div role="option" class="option">Apple</div>
      <div role="option" class="option">Banana</div>
      <div role="option" class="option">Orange</div>
    </div>
  </div>
</div>
```

### Disabled Select
```html
<div id="select-disabled" class="select">
  <button type="button" class="btn-outline justify-between font-normal" 
          id="select-disabled-trigger" 
          aria-haspopup="listbox" 
          aria-expanded="false" 
          aria-controls="select-disabled-listbox" 
          disabled="disabled">
    <span class="truncate">Select an option</span>
    <svg class="lucide lucide-chevron-down text-muted-foreground opacity-50 shrink-0">
      <path d="m6 9 6 6 6-6" />
    </svg>
  </button>
</div>
```

### With Label
```html
<div class="grid gap-3">
  <label for="select-with-label" class="label">Choose a fruit</label>
  <select id="select-with-label" class="select w-[180px]">
    <option value="">Select a fruit</option>
    <option value="apple">Apple</option>
    <option value="banana">Banana</option>
    <option value="orange">Orange</option>
  </select>
</div>
```

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "select" class | Yes |
| `id` | string | Unique identifier | Recommended |
| `name` | string | Form field name | No |
| `disabled` | boolean | Disables the select | No |
| `multiple` | boolean | Allow multiple selections | No |
| `size` | number | Visible options count | No |

## Custom Select Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `aria-haspopup` | string | "listbox" for select behavior | Yes |
| `aria-expanded` | string | "true"/"false" for open state | Yes |
| `aria-controls` | string | ID of listbox element | Yes |
| `data-popover` | string | Marks popover container | Yes |
| `role` | string | "listbox" and "option" for options | Yes |

## JavaScript Integration

### Required Scripts
The custom select requires JavaScript for functionality:

```html
<script src="/assets/js/select.js" defer></script>
```

### Popover System
Uses the `data-popover` attribute for dropdown positioning and behavior:

```javascript
// Automatic initialization when page loads
document.addEventListener('DOMContentLoaded', () => {
  if (window.basecoat) window.basecoat.initAll();
});
```

### Event Handling
Custom selects dispatch standard `change` events:

```javascript
document.getElementById('my-select-trigger').addEventListener('change', (e) => {
  console.log('Selected value:', e.detail.value);
});
```

## Accessibility Features

- **Keyboard Navigation**: Arrow keys, Enter, Escape, and type-ahead
- **Screen Reader Support**: Proper ARIA attributes and roles
- **Focus Management**: Maintains focus correctly when opening/closing
- **Label Association**: Works with `<label>` elements
- **Option Groups**: Support for `<optgroup>` in native selects

## States

### Default State
Closed select with placeholder or selected value

### Open State
Dropdown is visible with options list

### Focused State
Visible focus ring around trigger button

### Disabled State
Reduced opacity and no interaction

### Selected State
Highlighted option and updated trigger text

## Form Integration

### Native Form Support
```html
<form class="form space-y-4">
  <div class="grid gap-3">
    <label for="fruit-select">Favorite Fruit</label>
    <select id="fruit-select" name="fruit" class="select">
      <option value="">Choose a fruit</option>
      <option value="apple">Apple</option>
      <option value="banana">Banana</option>
    </select>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

### Custom Select in Forms
Custom selects include hidden input fields for form submission:

```html
<form class="form space-y-4">
  <div id="custom-select" class="select">
    <input type="hidden" name="fruit" value="apple">
    <button type="button" class="btn-outline justify-between font-normal">
      <span class="truncate">Apple</span>
      <!-- Chevron icon -->
    </button>
    <!-- Popover with options -->
  </div>
</form>
```

## Best Practices

1. Use native selects for simple lists and better mobile experience
2. Use custom selects when you need enhanced styling or functionality
3. Always provide a default "empty" option for non-required fields
4. Group related options with `<optgroup>` for native selects
5. Ensure adequate contrast for all states
6. Test thoroughly with keyboard navigation
7. Consider mobile-first design (native selects work better on mobile)

## Styling Customization

### Width Control
```html
<select class="select w-32">...</select>   <!-- 128px width -->
<select class="select w-48">...</select>   <!-- 192px width -->
<select class="select w-full">...</select> <!-- Full width -->
```

### Size Variants
```html
<!-- Small -->
<select class="select text-sm h-8">...</select>

<!-- Default -->
<select class="select">...</select>

<!-- Large -->
<select class="select text-lg h-12">...</select>
```

## Related Components

- [Input](./input.md) - For text input fields
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling
- [Popover](./popover.md) - For custom dropdown positioning
- [Button](./button.md) - For trigger styling