# Button Component

Displays a button or a component that looks like a button.

## Basic Usage

```html
<button class="btn">Button</button>
```

## CSS Classes

### Button Variants
- **`btn`** or **`btn-primary`** - Primary buttons (default)
- **`btn-secondary`** - Secondary buttons  
- **`btn-destructive`** - Destructive/danger buttons
- **`btn-outline`** - Outlined buttons
- **`btn-ghost`** - Ghost/transparent buttons
- **`btn-link`** - Link-style buttons
- **`btn-icon`** - Icon-only buttons

### Button Sizes
- **Default** - Standard button size
- **`btn-sm`** - Small buttons
- **`btn-lg`** - Large buttons

### Size + Variant Combinations
You can combine sizes with any variant:
- `btn-lg-destructive` - Large destructive button
- `btn-sm-outline` - Small outline button
- `btn-sm-icon-outline` - Small outline icon button
- `btn-icon-destructive` - Destructive icon button

## Component Attributes

### Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Button styling classes | Yes |
| `type` | string | "button", "submit", "reset" | Recommended |
| `disabled` | boolean | Disables the button | Optional |
| `aria-label` | string | Accessible label for icon buttons | Icon buttons |

### No JavaScript Required (Basic)
Basic buttons work with pure CSS and HTML.

## HTML Structure

```html
<!-- Basic button -->
<button class="btn" type="button">Label</button>

<!-- Button with icon -->
<button class="btn">
  <svg><!-- icon --></svg>
  Label
</button>

<!-- Icon-only button -->
<button class="btn-icon" aria-label="Description">
  <svg><!-- icon --></svg>
</button>
```

## Examples

### Button Variants

```html
<!-- Primary (default) -->
<button class="btn">Primary</button>
<button class="btn-primary">Primary</button>

<!-- Secondary -->
<button class="btn-secondary">Secondary</button>

<!-- Destructive -->
<button class="btn-destructive">Destructive</button>

<!-- Outline -->
<button class="btn-outline">Outline</button>

<!-- Ghost -->
<button class="btn-ghost">Ghost</button>

<!-- Link -->
<button class="btn-link">Link</button>
```

### Button Sizes

```html
<!-- Small buttons -->
<button class="btn-sm">Small</button>
<button class="btn-sm-secondary">Small Secondary</button>
<button class="btn-sm-outline">Small Outline</button>

<!-- Default size -->
<button class="btn">Default</button>

<!-- Large buttons -->
<button class="btn-lg">Large</button>
<button class="btn-lg-destructive">Large Destructive</button>
<button class="btn-lg-outline">Large Outline</button>
```

### Icon Buttons

```html
<!-- Icon-only buttons -->
<button class="btn-icon" aria-label="Next">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="m9 18 6-6-6-6" />
  </svg>
</button>

<button class="btn-icon-outline" aria-label="Settings">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
    <circle cx="12" cy="12" r="3" />
  </svg>
</button>

<button class="btn-icon-destructive" aria-label="Delete">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3 6h18" />
    <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
    <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
  </svg>
</button>

<!-- Small icon buttons -->
<button class="btn-sm-icon" aria-label="Edit">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
    <path d="M18.375 2.625a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.375-9.375Z" />
  </svg>
</button>
```

### Buttons with Icons and Text

```html
<!-- Icon + text buttons -->
<button class="btn">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M14.536 21.686a.5.5 0 0 0 .937-.024l6.5-19a.496.496 0 0 0-.635-.635l-19 6.5a.5.5 0 0 0-.024.937l7.93 3.18a2 2 0 0 1 1.112 1.11z" />
    <path d="m21.854 2.147-10.94 10.939" />
  </svg>
  Send email
</button>

<button class="btn-outline">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
    <polyline points="7,10 12,15 17,10" />
    <line x1="12" x2="12" y1="15" y2="3" />
  </svg>
  Download
</button>

<button class="btn-secondary">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M5 12l5 5l10 -10" />
  </svg>
  Save changes
</button>

<!-- Icon after text -->
<button class="btn">
  Continue
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="m9 18 6-6-6-6" />
  </svg>
</button>
```

### Loading State Buttons

```html
<!-- Button with loading spinner -->
<button class="btn" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Loading...
</button>

<!-- Loading icon only -->
<button class="btn-icon" disabled aria-label="Loading">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>

<!-- Different loading states -->
<button class="btn-outline" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Please wait...
</button>

<button class="btn-sm-secondary" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Processing
</button>
```

### Button States

```html
<!-- Normal state -->
<button class="btn">Normal</button>

<!-- Disabled state -->
<button class="btn" disabled>Disabled</button>

<!-- Active/pressed state -->
<button class="btn pressed" aria-pressed="true">Active</button>

<!-- Focus state (handled by CSS automatically) -->
<button class="btn">Focusable</button>

<!-- Different variants disabled -->
<button class="btn-outline" disabled>Disabled Outline</button>
<button class="btn-destructive" disabled>Disabled Destructive</button>
<button class="btn-ghost" disabled>Disabled Ghost</button>
```

### Form Integration

```html
<!-- Form buttons -->
<form>
  <div class="space-y-4">
    <!-- Form fields here -->
    
    <!-- Form actions -->
    <div class="flex justify-end gap-2">
      <button type="button" class="btn-outline">Cancel</button>
      <button type="submit" class="btn">Save</button>
    </div>
  </div>
</form>

<!-- Different submit scenarios -->
<form>
  <!-- Basic submit -->
  <button type="submit" class="btn">Submit</button>
  
  <!-- Submit with loading -->
  <button type="submit" class="btn" id="submit-btn">
    <span class="hidden loading-spinner">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
        <path d="M21 12a9 9 0 1 1-6.219-8.56" />
      </svg>
    </span>
    <span class="button-text">Submit</span>
  </button>
  
  <!-- Reset button -->
  <button type="reset" class="btn-outline">Reset</button>
</form>
```

### Button as Links

```html
<!-- Styled as button, behaves as link -->
<a href="/dashboard" class="btn">Go to Dashboard</a>
<a href="/profile" class="btn-outline">View Profile</a>
<a href="/settings" class="btn-ghost">Settings</a>

<!-- External links -->
<a href="https://example.com" target="_blank" rel="noopener noreferrer" class="btn">
  Visit Site
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M7 7h10v10" />
    <path d="M7 17 17 7" />
  </svg>
</a>
```

### Button Groups

```html
<!-- Simple button group -->
<div class="flex gap-2">
  <button class="btn-outline">Previous</button>
  <button class="btn-outline">Next</button>
</div>

<!-- Segmented control style -->
<div class="inline-flex border border-border rounded-md">
  <button class="btn-ghost rounded-r-none border-r">Day</button>
  <button class="btn-ghost rounded-none border-r">Week</button>
  <button class="btn rounded-l-none">Month</button>
</div>

<!-- Action group -->
<div class="flex items-center gap-2">
  <button class="btn">Save</button>
  <button class="btn-outline">Save & Continue</button>
  <div class="border-l border-border pl-2 ml-2">
    <button class="btn-icon-ghost" aria-label="More options">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="12" cy="5" r="1" />
        <circle cx="12" cy="19" r="1" />
      </svg>
    </button>
  </div>
</div>
```

### Responsive Buttons

```html
<!-- Responsive button sizes -->
<button class="btn-sm md:btn">
  Responsive Size
</button>

<!-- Hide text on small screens -->
<button class="btn">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
    <polyline points="7,10 12,15 17,10" />
    <line x1="12" x2="12" y1="15" y2="3" />
  </svg>
  <span class="hidden sm:inline ml-2">Download</span>
</button>

<!-- Stack buttons on mobile -->
<div class="flex flex-col sm:flex-row gap-2">
  <button class="btn-outline">Cancel</button>
  <button class="btn">Confirm</button>
</div>
```

## Accessibility Features

- **Keyboard Navigation**: All buttons support Tab/Enter/Space
- **Screen Reader Support**: Use `aria-label` for icon-only buttons
- **Focus Indicators**: Automatic focus rings and states
- **Disabled State**: Proper `disabled` attribute handling

### Enhanced Accessibility

```html
<!-- Icon button with proper labeling -->
<button class="btn-icon-outline" aria-label="Close dialog">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M18 6 6 18" />
    <path d="m6 6 12 12" />
  </svg>
</button>

<!-- Toggle button with pressed state -->
<button class="btn-outline" aria-pressed="false" onclick="togglePressed(this)">
  Toggle Feature
</button>

<!-- Button with description -->
<button class="btn" aria-describedby="save-help">
  Save Draft
</button>
<p id="save-help" class="sr-only">
  Saves your work without publishing
</p>

<!-- Loading button with live region -->
<button class="btn" onclick="submitForm(this)" aria-describedby="status">
  Submit
</button>
<div id="status" aria-live="polite" class="sr-only"></div>
```

## JavaScript Integration

### Button State Management

```javascript
// Toggle loading state
function setButtonLoading(button, isLoading) {
  const spinner = button.querySelector('.loading-spinner');
  const text = button.querySelector('.button-text');
  
  if (isLoading) {
    button.disabled = true;
    spinner?.classList.remove('hidden');
    text && (text.textContent = 'Loading...');
  } else {
    button.disabled = false;
    spinner?.classList.add('hidden');
    text && (text.textContent = 'Submit');
  }
}

// Toggle pressed state
function togglePressed(button) {
  const isPressed = button.getAttribute('aria-pressed') === 'true';
  button.setAttribute('aria-pressed', (!isPressed).toString());
  button.classList.toggle('pressed', !isPressed);
}

// Form submission with loading
function submitForm(button) {
  setButtonLoading(button, true);
  
  // Simulate async operation
  setTimeout(() => {
    setButtonLoading(button, false);
  }, 2000);
}
```

### React Integration

```jsx
import React, { useState } from 'react';

function Button({ 
  variant = 'btn',
  size = '',
  children,
  loading = false,
  disabled = false,
  className = '',
  ...props 
}) {
  const buttonClasses = [
    size ? `${size}-${variant}` : variant,
    className
  ].filter(Boolean).join(' ');
  
  return (
    <button 
      className={buttonClasses}
      disabled={disabled || loading}
      {...props}
    >
      {loading && (
        <svg className="animate-spin mr-2" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
          <path d="M21 12a9 9 0 1 1-6.219-8.56" />
        </svg>
      )}
      {children}
    </button>
  );
}

// Usage
function App() {
  const [loading, setLoading] = useState(false);
  
  const handleSubmit = async () => {
    setLoading(true);
    await new Promise(resolve => setTimeout(resolve, 2000));
    setLoading(false);
  };
  
  return (
    <div className="space-x-2">
      <Button variant="btn">Primary</Button>
      <Button variant="btn-outline">Outline</Button>
      <Button variant="btn" loading={loading} onClick={handleSubmit}>
        Submit
      </Button>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <button 
    :class="buttonClasses"
    :disabled="disabled || loading"
    v-bind="$attrs"
  >
    <svg 
      v-if="loading"
      class="animate-spin mr-2" 
      width="20" height="20" 
      viewBox="0 0 24 24" 
      fill="none" 
      stroke="currentColor" 
      stroke-width="2"
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <slot />
  </button>
</template>

<script>
export default {
  props: {
    variant: {
      type: String,
      default: 'btn'
    },
    size: String,
    loading: Boolean,
    disabled: Boolean
  },
  computed: {
    buttonClasses() {
      const classes = [];
      if (this.size) {
        classes.push(`${this.size}-${this.variant}`);
      } else {
        classes.push(this.variant);
      }
      return classes.join(' ');
    }
  }
};
</script>
```

## Best Practices

1. **Use Semantic HTML**: Use `<button>` for actions, `<a>` for navigation
2. **Button Types**: Specify `type` attribute for form buttons
3. **Accessibility**: Provide `aria-label` for icon-only buttons
4. **Loading States**: Disable buttons during async operations
5. **Visual Hierarchy**: Use appropriate variants for action priority
6. **Consistent Sizing**: Maintain consistent button sizes within groups
7. **Icon Guidelines**: Use 24px icons for default buttons, 16px for small
8. **Focus Management**: Ensure proper focus indicators and flow

## Common Patterns

### Confirmation Dialogs

```html
<div class="space-y-4">
  <p>Are you sure you want to delete this item?</p>
  <div class="flex justify-end gap-2">
    <button type="button" class="btn-outline" onclick="closeDialog()">
      Cancel
    </button>
    <button type="button" class="btn-destructive" onclick="confirmDelete()">
      Delete
    </button>
  </div>
</div>
```

### Toolbar Actions

```html
<div class="flex items-center gap-1 p-2 border-b">
  <button class="btn-sm-icon" aria-label="Bold">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
      <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
    </svg>
  </button>
  <button class="btn-sm-icon" aria-label="Italic">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="19" x2="10" y1="4" y2="4" />
      <line x1="14" x2="5" y1="20" y2="20" />
      <line x1="15" x2="9" y1="4" y2="20" />
    </svg>
  </button>
  <div class="border-l border-border mx-1 h-6"></div>
  <button class="btn-sm-outline">Save</button>
</div>
```

### Call-to-Action Sections

```html
<div class="text-center space-y-4 py-12">
  <h2 class="text-2xl font-bold">Ready to get started?</h2>
  <p class="text-muted-foreground">Join thousands of satisfied customers today.</p>
  <div class="flex justify-center gap-3">
    <button class="btn-lg">Start Free Trial</button>
    <button class="btn-lg-outline">Learn More</button>
  </div>
</div>
```

## Related Components

- [Button Group](./button-group.md) - For grouping related buttons
- [Spinner](./spinner.md) - For button loading states
- [Input](./input.md) - For form integration
- [Dialog](./dialog.md) - For confirmation buttons