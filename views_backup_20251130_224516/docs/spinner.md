# Spinner Component

An indicator that can be used to show a loading state.

## Important Note

**There is no dedicated Spinner component in Basecoat.** Spinners are pure HTML using the Lucide `loader-circle` icon with the `animate-spin` Tailwind utility.

## Basic Usage

```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard SVG with Tailwind utilities

### Supporting Classes
- **`animate-spin`** - Applies continuous rotation animation
- **`size-*`** - Controls spinner dimensions
- **`text-*`** - Controls spinner color

### Tailwind Utilities Used
- `animate-spin` - Continuous rotation animation
- `size-3` - 12px dimensions (0.75rem)
- `size-4` - 16px dimensions (1rem) 
- `size-6` - 24px dimensions (1.5rem)
- `size-8` - 32px dimensions (2rem)
- `text-*` - Color utilities (text-red-500, text-primary, etc.)

## Component Attributes

### SVG Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `xmlns` | string | SVG namespace | Yes |
| `width` | string | SVG width (usually "24") | Yes |
| `height` | string | SVG height (usually "24") | Yes |
| `viewBox` | string | SVG viewBox (usually "0 0 24 24") | Yes |
| `fill` | string | Should be "none" for outline style | Yes |
| `stroke` | string | Should be "currentColor" | Yes |
| `stroke-width` | string | Stroke thickness (usually "2") | Yes |
| `stroke-linecap` | string | Line cap style (usually "round") | Yes |
| `stroke-linejoin` | string | Line join style (usually "round") | Yes |
| `role` | string | Should be "status" for accessibility | Recommended |
| `aria-label` | string | Should be "Loading" or descriptive | Recommended |
| `class` | string | Must include "animate-spin" | Yes |

### No JavaScript Required
This is a pure CSS/HTML implementation using SVG animation.

## HTML Structure

```html
<svg xmlns="http://www.w3.org/2000/svg" 
     width="24" height="24" 
     viewBox="0 0 24 24" 
     fill="none" 
     stroke="currentColor" 
     stroke-width="2" 
     stroke-linecap="round" 
     stroke-linejoin="round" 
     role="status" 
     aria-label="Loading" 
     class="animate-spin [size-classes] [color-classes]">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## Examples

### Different Sizes
```html
<!-- Extra Small (12px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-3 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Small (16px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Medium (24px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Large (32px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-8 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

### Different Colors
```html
<!-- Red spinner -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-red-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Green spinner -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-green-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Primary color -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-primary">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Muted color -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-muted-foreground">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

### Button with Spinner
```html
<!-- Primary button with spinner -->
<button class="btn-sm" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Loading...
</button>

<!-- Outline button with spinner -->
<button class="btn-sm-outline" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Please wait
</button>

<!-- Secondary button with spinner -->
<button class="btn-sm-secondary" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Processing
</button>
```

### Spinner Only (No Text)
```html
<!-- Icon button with spinner only -->
<button class="btn-icon" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>

<!-- Small icon button -->
<button class="btn-icon-sm" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>
```

### Card Loading State
```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-transparent bg-muted/50 p-4 gap-4">
  <div class="flex shrink-0 items-center justify-center [&_svg]:pointer-events-none [&_svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium line-clamp-1">Processing payment...</h3>
  </div>
  <div class="flex flex-col gap-1 flex-none text-center">
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance">$100.00</p>
  </div>
</article>
```

### Inline Spinner
```html
<!-- Inline with text -->
<p class="flex items-center gap-2">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Updating content...
</p>

<!-- Inline with different alignments -->
<div class="flex items-start gap-2">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin mt-0.5">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  <div>
    <p class="font-medium">Loading data...</p>
    <p class="text-sm text-muted-foreground">This may take a moment</p>
  </div>
</div>
```

### Full Page Loading
```html
<!-- Full screen loading overlay -->
<div class="fixed inset-0 bg-background/80 backdrop-blur-sm z-50 flex items-center justify-center">
  <div class="flex flex-col items-center gap-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-8 animate-spin text-primary">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <p class="text-muted-foreground">Loading application...</p>
  </div>
</div>

<!-- Centered in container -->
<div class="flex items-center justify-center h-64">
  <div class="text-center">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-muted-foreground mx-auto mb-2">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <p class="text-sm text-muted-foreground">Loading content...</p>
  </div>
</div>
```

### Input Loading State
```html
<!-- Input with spinner -->
<div class="relative">
  <input type="text" class="input pr-10" placeholder="Search..." disabled />
  <div class="absolute inset-y-0 right-0 flex items-center pr-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Searching" class="size-4 animate-spin text-muted-foreground">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
</div>

<!-- Select with spinner -->
<div class="relative">
  <select class="select pr-10" disabled>
    <option>Loading options...</option>
  </select>
  <div class="absolute inset-y-0 right-8 flex items-center pr-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin text-muted-foreground">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
</div>
```

### Custom Animation Speed
```html
<!-- Faster animation -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin animation-duration-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Slower animation -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin animation-duration-2000">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Custom CSS for animation timing -->
<style>
  .animate-slow-spin {
    animation: spin 2s linear infinite;
  }
  .animate-fast-spin {
    animation: spin 0.5s linear infinite;
  }
</style>

<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-slow-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## Accessibility Features

- **Role Attribute**: Use `role="status"` to announce loading states
- **ARIA Label**: Provide descriptive `aria-label` for screen readers
- **Live Regions**: Can be used with `aria-live` regions for dynamic updates
- **Focus Management**: Consider focus management when loading states change

### Enhanced Accessibility
```html
<!-- With live region -->
<div aria-live="polite" aria-atomic="true">
  <div class="flex items-center gap-2">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading user data" class="size-4 animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <span>Loading user data...</span>
  </div>
</div>

<!-- With hidden text for screen readers -->
<button class="btn" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-hidden="true" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  <span class="sr-only">Loading, please wait</span>
  <span aria-hidden="true">Loading...</span>
</button>
```

## Performance Considerations

### Conditional Rendering
```html
<!-- Show spinner only when loading -->
<div>
  <!-- Loading state -->
  <div class="flex items-center gap-2" style="display: none;" id="loading-state">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    Loading...
  </div>
  
  <!-- Content state -->
  <div id="content-state">
    Content loaded successfully!
  </div>
</div>

<script>
function showLoading() {
  document.getElementById('loading-state').style.display = 'flex';
  document.getElementById('content-state').style.display = 'none';
}

function hideLoading() {
  document.getElementById('loading-state').style.display = 'none';
  document.getElementById('content-state').style.display = 'block';
}
</script>
```

### Reduced Motion Support
```html
<!-- Respects user's motion preferences -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin motion-reduce:animate-none">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- CSS alternative -->
<style>
  @media (prefers-reduced-motion: reduce) {
    .spinner-respectful {
      animation: none;
    }
  }
</style>
```

## Best Practices

1. **Descriptive Labels**: Use meaningful `aria-label` attributes
2. **Appropriate Sizing**: Match spinner size to context
3. **Color Contrast**: Ensure sufficient contrast for visibility
4. **Loading States**: Clear communication of what's loading
5. **Timeout Handling**: Provide fallbacks for long loading times
6. **Reduced Motion**: Respect user motion preferences
7. **Performance**: Only show spinners when necessary
8. **User Feedback**: Combine with progress indicators when possible

## Size Reference

| Class | Size | Pixels | Use Case |
|-------|------|--------|----------|
| `size-3` | 0.75rem | 12px | Tiny indicators, inline text |
| `size-4` | 1rem | 16px | Small buttons, input fields |
| `size-5` | 1.25rem | 20px | Medium inline elements |
| `size-6` | 1.5rem | 24px | Standard buttons, cards |
| `size-8` | 2rem | 32px | Large buttons, prominent loading |
| `size-10` | 2.5rem | 40px | Modal dialogs, major sections |
| `size-12` | 3rem | 48px | Full-page loading |

## Common Patterns

### Form Submission
```html
<form onsubmit="showLoadingState()">
  <div class="space-y-4">
    <!-- Form fields -->
    <input type="email" class="input" placeholder="Email" required />
    
    <!-- Submit button with loading state -->
    <button type="submit" class="btn" id="submit-btn">
      <span class="flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin hidden" id="submit-spinner">
          <path d="M21 12a9 9 0 1 1-6.219-8.56" />
        </svg>
        <span id="submit-text">Submit</span>
      </span>
    </button>
  </div>
</form>

<script>
function showLoadingState() {
  const btn = document.getElementById('submit-btn');
  const spinner = document.getElementById('submit-spinner');
  const text = document.getElementById('submit-text');
  
  btn.disabled = true;
  spinner.classList.remove('hidden');
  text.textContent = 'Submitting...';
}
</script>
```

### Data Table Loading
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody id="table-body">
      <!-- Loading state -->
      <tr>
        <td colspan="3" class="text-center py-8">
          <div class="flex flex-col items-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading data" class="size-6 animate-spin text-muted-foreground">
              <path d="M21 12a9 9 0 1 1-6.219-8.56" />
            </svg>
            <p class="text-muted-foreground">Loading data...</p>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Spinner({ size = 'size-4', color = 'text-current', className = '', ...props }) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      role="status"
      aria-label="Loading"
      className={`animate-spin ${size} ${color} ${className}`}
      {...props}
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  );
}

// Usage
<Spinner size="size-6" color="text-primary" />
<button disabled>
  <Spinner size="size-4" />
  Loading...
</button>
```

### Vue Integration
```vue
<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    role="status"
    aria-label="Loading"
    :class="spinnerClasses"
    v-bind="$attrs"
  >
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</template>

<script>
export default {
  props: {
    size: {
      type: String,
      default: 'size-4'
    },
    color: {
      type: String,
      default: 'text-current'
    }
  },
  computed: {
    spinnerClasses() {
      return `animate-spin ${this.size} ${this.color}`;
    }
  }
};
</script>
```

### Web Components
```javascript
class SpinnerElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
  }
  
  connectedCallback() {
    const size = this.getAttribute('size') || 'size-4';
    const color = this.getAttribute('color') || 'text-current';
    
    this.shadowRoot.innerHTML = `
      <style>
        @import url('path/to/tailwind.css');
        .spinner { animation: spin 1s linear infinite; }
        @keyframes spin {
          from { transform: rotate(0deg); }
          to { transform: rotate(360deg); }
        }
      </style>
      <svg xmlns="http://www.w3.org/2000/svg" 
           width="24" height="24" 
           viewBox="0 0 24 24" 
           fill="none" 
           stroke="currentColor" 
           stroke-width="2" 
           stroke-linecap="round" 
           stroke-linejoin="round" 
           role="status" 
           aria-label="Loading"
           class="spinner ${size} ${color}">
        <path d="M21 12a9 9 0 1 1-6.219-8.56" />
      </svg>
    `;
  }
}

customElements.define('loading-spinner', SpinnerElement);
```

## Related Components

- [Button](./button.md) - For loading button states
- [Card](./card.md) - For card loading states
- [Table](./table.md) - For table loading states
- [Progress](./progress.md) - For determinate progress indicators