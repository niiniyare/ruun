# Textarea Component

Displays a form textarea or a component that looks like a textarea.

## Basic Usage

```html
<textarea class="textarea" placeholder="Type your message here"></textarea>
```

## CSS Classes

### Primary Class
- **`textarea`** - Main textarea styling class

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
- `resize-none` - Disable manual resizing (optional)
- `min-h-[80px]` - Minimum height

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "textarea" class | Yes |
| `placeholder` | string | Placeholder text | No |
| `disabled` | boolean | Disables the textarea | No |
| `aria-invalid` | boolean | Marks textarea as invalid for error states | No |
| `id` | string | Unique identifier for label association | No |
| `name` | string | Form field name | No |
| `rows` | number | Visible text lines | No |
| `cols` | number | Visible character width | No |
| `maxlength` | number | Maximum character count | No |
| `required` | boolean | Makes field required | No |

## Examples

### Basic Textarea
```html
<textarea class="textarea" placeholder="Type your message here"></textarea>
```

### Invalid Textarea
```html
<textarea class="textarea" placeholder="Type your message here" aria-invalid="true"></textarea>
```

### Disabled Textarea
```html
<textarea class="textarea" placeholder="Type your message here" disabled></textarea>
```

### With Label
```html
<div class="grid gap-3">
  <label for="textarea-with-label" class="label">Message</label>
  <textarea id="textarea-with-label" class="textarea" placeholder="Type your message here"></textarea>
</div>
```

### With Help Text
```html
<div class="grid gap-3">
  <label for="textarea-with-help" class="label">Comment</label>
  <textarea id="textarea-with-help" class="textarea" placeholder="Type your message here"></textarea>
  <p class="text-muted-foreground text-sm">Type your message and press Ctrl+Enter to send.</p>
</div>
```

### With Character Counter
```html
<div class="grid gap-3">
  <label for="textarea-counter" class="label">Bio</label>
  <textarea id="textarea-counter" class="textarea" placeholder="Tell us about yourself" maxlength="160"></textarea>
  <p class="text-muted-foreground text-sm">0/160 characters</p>
</div>
```

### Form Context
```html
<form class="form space-y-6">
  <div class="grid gap-3">
    <label for="textarea-form" class="label">Bio</label>
    <textarea id="textarea-form" placeholder="Tell us a bit about yourself"></textarea>
    <p class="text-muted-foreground text-sm">You can @mention other users and organizations.</p>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

### Auto-resizing Textarea
```html
<textarea 
  class="textarea resize-none overflow-hidden" 
  placeholder="This textarea grows as you type..."
  oninput="this.style.height = 'auto'; this.style.height = this.scrollHeight + 'px'"></textarea>
```

### Fixed Height Textarea
```html
<textarea class="textarea h-32 resize-none" placeholder="Fixed height textarea"></textarea>
```

## Size Variants

### Small
```html
<textarea class="textarea text-sm min-h-[60px]" placeholder="Small textarea"></textarea>
```

### Default
```html
<textarea class="textarea" placeholder="Default textarea"></textarea>
```

### Large
```html
<textarea class="textarea text-base min-h-[120px]" placeholder="Large textarea"></textarea>
```

## States

### Default State
Regular appearance with placeholder text

### Focus State
Visible focus ring and enhanced border

### Invalid State
Error styling when `aria-invalid="true"` is set

### Disabled State
Reduced opacity and no interaction

### Filled State
Contains user input text

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, textareas inherit styling automatically:

```html
<form class="form space-y-4">
  <!-- No need for explicit textarea class -->
  <div class="grid gap-3">
    <label for="auto-textarea">Description</label>
    <textarea id="auto-textarea" placeholder="Automatically styled"></textarea>
  </div>
</form>
```

### Validation Integration
```html
<div class="grid gap-3">
  <label for="validated-textarea" class="label">Required Field</label>
  <textarea 
    id="validated-textarea" 
    class="textarea" 
    required 
    aria-describedby="textarea-error">
  </textarea>
  <p id="textarea-error" class="text-error text-sm hidden">This field is required</p>
</div>
```

## JavaScript Integration

### Auto-resize Functionality
```javascript
function autoResize(textarea) {
  textarea.style.height = 'auto';
  textarea.style.height = textarea.scrollHeight + 'px';
}

// Apply to all auto-resize textareas
document.querySelectorAll('.textarea[data-auto-resize]').forEach(textarea => {
  textarea.addEventListener('input', () => autoResize(textarea));
  autoResize(textarea); // Initial sizing
});
```

### Character Counter
```javascript
function updateCounter(textarea) {
  const counter = textarea.nextElementSibling;
  if (counter && counter.dataset.counter) {
    const maxLength = textarea.maxLength;
    const currentLength = textarea.value.length;
    counter.textContent = `${currentLength}/${maxLength} characters`;
    
    // Visual feedback when approaching limit
    if (currentLength > maxLength * 0.9) {
      counter.classList.add('text-warning');
    } else {
      counter.classList.remove('text-warning');
    }
  }
}
```

### HTMX Integration
```html
<textarea 
  class="textarea" 
  hx-post="/api/save-draft" 
  hx-trigger="keyup changed delay:500ms" 
  hx-target="#save-status"
  placeholder="Auto-saves as you type">
</textarea>
<div id="save-status"></div>
```

## Accessibility Features

- **Keyboard Navigation**: Full keyboard support including Tab navigation
- **Screen Reader Support**: Proper labeling with `aria-describedby`
- **Error States**: `aria-invalid` for validation feedback
- **Focus Management**: Clear focus indicators
- **Label Association**: Proper `for` attribute usage

## Best Practices

1. Always associate with a label using `for` attribute
2. Provide helpful placeholder text
3. Use appropriate sizing for content type
4. Consider auto-resize for dynamic content
5. Implement character counting for limited fields
6. Provide clear validation feedback
7. Test with keyboard navigation
8. Ensure sufficient contrast in all states

## Styling Customization

### Border Variants
```html
<!-- Default border -->
<textarea class="textarea"></textarea>

<!-- No border -->
<textarea class="textarea border-0"></textarea>

<!-- Thick border -->
<textarea class="textarea border-2"></textarea>
```

### Background Variants
```html
<!-- Default background -->
<textarea class="textarea"></textarea>

<!-- Transparent background -->
<textarea class="textarea bg-transparent"></textarea>

<!-- Custom background -->
<textarea class="textarea bg-muted"></textarea>
```

### Resize Control
```html
<!-- No resize -->
<textarea class="textarea resize-none"></textarea>

<!-- Vertical resize only -->
<textarea class="textarea resize-y"></textarea>

<!-- Horizontal resize only -->
<textarea class="textarea resize-x"></textarea>

<!-- Both directions (default) -->
<textarea class="textarea resize"></textarea>
```

## Related Components

- [Input](./input.md) - For single-line text input
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling
- [Button](./button.md) - Often used with textarea in forms