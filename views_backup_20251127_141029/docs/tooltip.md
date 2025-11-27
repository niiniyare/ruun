# Tooltip Component

A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it.

## Basic Usage

```html
<button class="btn-outline" data-tooltip="Add to library">Hover</button>
```

## CSS Classes

### No Custom Classes Required
Tooltips are controlled entirely through data attributes. The Basecoat JavaScript automatically handles styling.

### Supporting Classes
- **Any element** can have a tooltip by adding `data-tooltip` attribute
- **Layout classes** from other components work normally

### Tailwind Utilities Used
Tooltips use internal styling from the Basecoat JavaScript - no custom Tailwind classes needed.

## Component Attributes

### Required Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-tooltip` | string | Text content of the tooltip | Yes |

### Optional Attributes
| Attribute | Type | Values | Description | Required |
|-----------|------|--------|-------------|----------|
| `data-side` | string | "top", "bottom", "left", "right" | Position of tooltip relative to element | No (defaults to top) |
| `data-align` | string | "start", "center", "end" | Alignment of tooltip on the chosen side | No (defaults to center) |

## JavaScript Required

This component requires the Basecoat JavaScript files for tooltip functionality:

```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
```

## HTML Structure

```html
<!-- Any element can have a tooltip -->
<element data-tooltip="Tooltip text" data-side="position" data-align="alignment">
  Element content
</element>
```

## Examples

### Default Tooltip (Top Center)

```html
<button class="btn-outline" data-tooltip="Default tooltip">Default</button>
```

### Position Variations

```html
<!-- Top (default) -->
<button class="btn-outline" data-tooltip="Top tooltip" data-side="top">Top</button>

<!-- Bottom -->
<button class="btn-outline" data-tooltip="Bottom tooltip" data-side="bottom">Bottom</button>

<!-- Left -->
<button class="btn-outline" data-tooltip="Left tooltip" data-side="left">Left</button>

<!-- Right -->
<button class="btn-outline" data-tooltip="Right side tooltip" data-side="right">Right</button>
```

### Alignment Variations

```html
<!-- Center aligned (default) -->
<button class="btn-outline" data-tooltip="Center aligned tooltip">Center</button>

<!-- Start aligned -->
<button class="btn-outline" data-tooltip="Start aligned tooltip" data-align="start">Start</button>

<!-- End aligned -->
<button class="btn-outline" data-tooltip="End aligned tooltip" data-align="end">End</button>
```

### Combined Position and Alignment

```html
<!-- Bottom + Start -->
<button class="btn-outline" data-tooltip="Bottom side and start aligned tooltip" data-side="bottom" data-align="start">Bottom + Start</button>

<!-- Right + End -->
<button class="btn-outline" data-tooltip="Right side and end aligned tooltip" data-side="right" data-align="end">Right + End</button>

<!-- Left + Center -->
<button class="btn-outline" data-tooltip="Left side and center aligned tooltip" data-side="left" data-align="center">Left + Center</button>

<!-- Top + End -->
<button class="btn-outline" data-tooltip="Top side and end aligned tooltip" data-side="top" data-align="end">Top + End</button>
```

### Icon Button Tooltips

```html
<!-- Icon buttons with descriptive tooltips -->
<button class="btn-icon-outline" data-tooltip="Save document" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" />
    <polyline points="17,21 17,13 7,13 7,21" />
    <polyline points="7,3 7,8 15,8" />
  </svg>
</button>

<button class="btn-icon-outline" data-tooltip="Edit item" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
    <path d="M18.375 2.625a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.375-9.375Z" />
  </svg>
</button>

<button class="btn-icon-destructive" data-tooltip="Delete forever" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3 6h18" />
    <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
    <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
  </svg>
</button>
```

### Link Tooltips

```html
<!-- External link with tooltip -->
<a href="https://example.com" target="_blank" class="btn-link" data-tooltip="Opens in new window">
  Visit External Site
</a>

<!-- Navigation link -->
<a href="/dashboard" class="btn-ghost" data-tooltip="View your dashboard">Dashboard</a>

<!-- Download link -->
<a href="/files/document.pdf" download class="btn-outline" data-tooltip="Download PDF (2.3MB)">
  Download Report
</a>
```

### Form Element Tooltips

```html
<!-- Input field with help tooltip -->
<div class="field">
  <label for="username">Username</label>
  <input 
    type="text" 
    id="username" 
    placeholder="Enter username"
    data-tooltip="Must be 3-20 characters, letters and numbers only"
    data-side="right"
  >
</div>

<!-- Checkbox with explanation -->
<div class="field">
  <label class="checkbox">
    <input 
      type="checkbox" 
      data-tooltip="We'll only send important updates about your account"
      data-side="right"
    >
    <span>Subscribe to notifications</span>
  </label>
</div>

<!-- Select with context -->
<div class="field">
  <label for="timezone">Timezone</label>
  <select 
    id="timezone" 
    class="select"
    data-tooltip="Used for scheduling and displaying dates"
    data-side="bottom"
  >
    <option>UTC</option>
    <option>EST</option>
    <option>PST</option>
  </select>
</div>
```

### Image Tooltips

```html
<!-- Avatar with user info -->
<img 
  src="/avatars/user.jpg" 
  alt="John Doe" 
  class="size-10 rounded-full"
  data-tooltip="John Doe - Product Manager"
  data-side="bottom"
>

<!-- Status indicator -->
<div class="flex items-center gap-2">
  <span class="size-3 rounded-full bg-green-500" data-tooltip="Online" data-side="top"></span>
  <span>User Status</span>
</div>

<!-- Chart data point -->
<div class="relative">
  <div 
    class="size-4 rounded-full bg-blue-500 cursor-pointer"
    data-tooltip="Sales: $12,450 (Mar 15)"
    data-side="top"
  ></div>
</div>
```

### Truncated Text Tooltips

```html
<!-- Long text with truncation -->
<div class="max-w-xs">
  <p 
    class="truncate text-sm"
    data-tooltip="This is a very long piece of text that would normally be truncated but the full content is shown in the tooltip"
    data-side="bottom"
  >
    This is a very long piece of text that would normally be truncated...
  </p>
</div>

<!-- Card title -->
<div class="card">
  <h3 
    class="text-lg font-semibold truncate"
    data-tooltip="Advanced Data Analytics and Reporting Dashboard Configuration Settings"
    data-side="bottom"
  >
    Advanced Data Analytics and...
  </h3>
</div>
```

### Interactive Element Tooltips

```html
<!-- Toggle switch -->
<label class="switch">
  <input type="checkbox" data-tooltip="Enable dark mode" data-side="right">
  <span>Dark Mode</span>
</label>

<!-- Slider -->
<div class="field">
  <label for="volume">Volume</label>
  <input 
    type="range" 
    id="volume" 
    min="0" 
    max="100" 
    value="75"
    data-tooltip="Current volume: 75%"
    data-side="top"
  >
</div>

<!-- Progress bar -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Upload Progress</span>
    <span>65%</span>
  </div>
  <div 
    class="progress"
    data-tooltip="Uploading... 65% complete"
    data-side="bottom"
  >
    <div class="progress-bar" style="width: 65%"></div>
  </div>
</div>
```

### Disabled Element Tooltips

```html
<!-- Disabled button with explanation -->
<button 
  class="btn" 
  disabled
  data-tooltip="Save is disabled until all required fields are completed"
  data-side="bottom"
>
  Save Changes
</button>

<!-- Disabled input -->
<input 
  type="text" 
  placeholder="Read only" 
  readonly
  data-tooltip="This field is automatically calculated"
  data-side="right"
>
```

### Badge and Status Tooltips

```html
<!-- Status badge -->
<span 
  class="badge-success"
  data-tooltip="Payment processed successfully on March 15, 2024"
  data-side="bottom"
>
  Paid
</span>

<!-- Priority badge -->
<span 
  class="badge-destructive"
  data-tooltip="Requires immediate attention"
  data-side="top"
>
  High Priority
</span>

<!-- Version badge -->
<span 
  class="badge-outline"
  data-tooltip="Released on March 10, 2024"
  data-side="bottom"
>
  v2.1.0
</span>
```

### Complex Content Tooltips

```html
<!-- Shortcut reference -->
<button 
  class="btn-outline"
  data-tooltip="Save (Ctrl+S)"
  data-side="bottom"
>
  Save
</button>

<!-- Feature description -->
<button 
  class="btn-ghost"
  data-tooltip="AI-powered content suggestions based on your writing style"
  data-side="right"
>
  Smart Compose
</button>

<!-- Technical info -->
<div 
  class="text-xs text-muted-foreground font-mono"
  data-tooltip="Last updated: 2024-03-15 14:23:15 UTC"
  data-side="top"
>
  #abc123
</div>
```

### Grid and Table Tooltips

```html
<!-- Table cell with overflow -->
<table class="table">
  <tbody>
    <tr>
      <td 
        class="max-w-xs truncate"
        data-tooltip="john.doe@example-company-with-long-domain.com"
        data-side="bottom"
      >
        john.doe@example-company...
      </td>
    </tr>
  </tbody>
</table>

<!-- Grid item -->
<div class="grid grid-cols-4 gap-4">
  <div 
    class="p-4 rounded border"
    data-tooltip="Click to view detailed analytics"
    data-side="top"
  >
    <div class="text-2xl font-bold">1.2k</div>
    <div class="text-sm text-muted-foreground">Views</div>
  </div>
</div>
```

## Accessibility Features

- **Keyboard Support**: Tooltips appear on focus for keyboard users
- **Mouse Support**: Tooltips appear on hover for mouse users
- **Screen Reader Friendly**: Uses appropriate ARIA attributes
- **Focus Management**: Doesn't interfere with keyboard navigation
- **Touch Friendly**: Works on touch devices with tap

### Enhanced Accessibility

```html
<!-- Using aria-label as fallback -->
<button 
  class="btn-icon-outline" 
  aria-label="Save document"
  data-tooltip="Save document"
  data-side="bottom"
>
  <svg><!-- save icon --></svg>
</button>

<!-- Descriptive tooltip for complex actions -->
<button 
  class="btn-destructive"
  aria-describedby="delete-tooltip"
  data-tooltip="Permanently delete this item. This action cannot be undone."
  data-side="top"
>
  Delete
</button>

<!-- Form field with help -->
<div class="field">
  <label for="password">Password</label>
  <input 
    type="password" 
    id="password"
    aria-describedby="password-help"
    data-tooltip="Must be at least 8 characters with uppercase, lowercase, number, and special character"
    data-side="right"
  >
</div>
```

## JavaScript Integration

### Dynamic Tooltip Content

```javascript
// Update tooltip content dynamically
function updateTooltip(element, newText) {
  element.setAttribute('data-tooltip', newText);
}

// Example: Update button tooltip based on state
const saveButton = document.getElementById('save-btn');
updateTooltip(saveButton, 'Save changes (Ctrl+S)');

// For form validation feedback
const emailInput = document.getElementById('email');
emailInput.addEventListener('blur', (e) => {
  const isValid = e.target.checkValidity();
  const tooltip = isValid 
    ? 'Valid email address' 
    : 'Please enter a valid email address';
  updateTooltip(e.target, tooltip);
});
```

### Conditional Tooltips

```javascript
// Show tooltip only when needed
function setConditionalTooltip(element, condition, tooltip) {
  if (condition) {
    element.setAttribute('data-tooltip', tooltip);
  } else {
    element.removeAttribute('data-tooltip');
  }
}

// Example: Truncated text
const textElement = document.querySelector('.truncated-text');
const isOverflowing = textElement.scrollWidth > textElement.clientWidth;
setConditionalTooltip(textElement, isOverflowing, textElement.textContent);

// Example: Disabled state explanation
const button = document.querySelector('#submit-btn');
const isDisabled = button.disabled;
setConditionalTooltip(button, isDisabled, 'Complete all required fields to enable');
```

### React Integration

```jsx
import React, { useState } from 'react';

function TooltipButton({ children, tooltip, side = 'top', align = 'center', ...props }) {
  return (
    <button
      data-tooltip={tooltip}
      data-side={side}
      data-align={align}
      {...props}
    >
      {children}
    </button>
  );
}

// Usage examples
function App() {
  const [saveStatus, setSaveStatus] = useState('idle');
  
  const getTooltipText = () => {
    switch (saveStatus) {
      case 'saving': return 'Saving changes...';
      case 'saved': return 'Changes saved successfully';
      case 'error': return 'Error saving changes';
      default: return 'Save changes (Ctrl+S)';
    }
  };
  
  return (
    <div>
      <TooltipButton 
        tooltip={getTooltipText()}
        side="bottom"
        onClick={handleSave}
        disabled={saveStatus === 'saving'}
      >
        Save
      </TooltipButton>
      
      <TooltipButton 
        tooltip="Delete this item permanently"
        side="top"
        className="btn-destructive"
      >
        Delete
      </TooltipButton>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div>
    <button 
      :data-tooltip="tooltipText"
      :data-side="side"
      :data-align="align"
      v-bind="$attrs"
    >
      <slot />
    </button>
  </div>
</template>

<script>
export default {
  props: {
    tooltip: {
      type: String,
      required: true
    },
    side: {
      type: String,
      default: 'top',
      validator: value => ['top', 'bottom', 'left', 'right'].includes(value)
    },
    align: {
      type: String,
      default: 'center',
      validator: value => ['start', 'center', 'end'].includes(value)
    }
  },
  computed: {
    tooltipText() {
      return this.tooltip;
    }
  }
};
</script>
```

### Tooltip Management Service

```javascript
// Centralized tooltip management
class TooltipManager {
  static updateTooltip(selector, text, options = {}) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    element.setAttribute('data-tooltip', text);
    if (options.side) element.setAttribute('data-side', options.side);
    if (options.align) element.setAttribute('data-align', options.align);
  }
  
  static removeTooltip(selector) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    element.removeAttribute('data-tooltip');
    element.removeAttribute('data-side');
    element.removeAttribute('data-align');
  }
  
  static setTemporaryTooltip(selector, text, duration = 3000) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    const originalTooltip = element.getAttribute('data-tooltip');
    element.setAttribute('data-tooltip', text);
    
    setTimeout(() => {
      if (originalTooltip) {
        element.setAttribute('data-tooltip', originalTooltip);
      } else {
        element.removeAttribute('data-tooltip');
      }
    }, duration);
  }
}

// Usage
TooltipManager.updateTooltip('#save-btn', 'Saved successfully!', { side: 'bottom' });
TooltipManager.setTemporaryTooltip('#copy-btn', 'Copied to clipboard!');
```

## Best Practices

1. **Concise Content**: Keep tooltip text short and informative
2. **Essential Information**: Only include tooltips for helpful context
3. **Consistent Positioning**: Use consistent sides for similar elements
4. **Keyboard Accessible**: Ensure tooltips work with keyboard navigation
5. **Touch Friendly**: Consider touch device interactions
6. **No Critical Info**: Don't put essential information only in tooltips
7. **Clear Language**: Use simple, clear language
8. **Performance**: Avoid too many tooltips on complex pages

## Common Patterns

### Icon Button Explanations

```html
<div class="flex gap-2">
  <button class="btn-icon-outline" data-tooltip="Bold (Ctrl+B)" data-side="bottom">
    <svg><!-- bold icon --></svg>
  </button>
  <button class="btn-icon-outline" data-tooltip="Italic (Ctrl+I)" data-side="bottom">
    <svg><!-- italic icon --></svg>
  </button>
  <button class="btn-icon-outline" data-tooltip="Underline (Ctrl+U)" data-side="bottom">
    <svg><!-- underline icon --></svg>
  </button>
</div>
```

### Status Indicators

```html
<div class="flex items-center gap-2">
  <div 
    class="size-2 rounded-full bg-green-500"
    data-tooltip="Service operational"
    data-side="top"
  ></div>
  <span>API Status</span>
</div>
```

### Form Field Help

```html
<div class="field">
  <label for="api-key">API Key</label>
  <input 
    type="password" 
    id="api-key"
    data-tooltip="Found in your account settings under 'API Access'"
    data-side="right"
  >
</div>
```

## Related Components

- [Popover](./popover.md) - For more complex tooltip content
- [Button](./button.md) - Common element that uses tooltips
- [Dialog](./dialog.md) - For detailed help content
- [Badge](./badge.md) - Often used with explanatory tooltips