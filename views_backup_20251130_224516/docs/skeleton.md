# Skeleton Component

Used to show a placeholder while content is loading.

## Important Note

**There is no dedicated Skeleton component in Basecoat.** Simply use the `animate-pulse` class to create a skeleton loader.

## Basic Usage

```html
<div class="flex items-center gap-4">
  <div class="bg-accent animate-pulse size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent animate-pulse rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent animate-pulse rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard Tailwind utilities

### Supporting Classes
- **`animate-pulse`** - Applies pulsing animation
- **`bg-accent`** - Background color for skeleton elements
- **Size classes**: `h-*`, `w-*`, `size-*` for dimensions
- **Shape classes**: `rounded-*` for different shapes

### Tailwind Utilities Used
- `animate-pulse` - Pulsing animation effect
- `bg-accent` - Light background color
- `bg-muted` - Alternative muted background
- `h-4` - Height (16px)
- `w-[150px]` - Fixed width
- `size-10` - Square dimensions (40x40px)
- `rounded-full` - Circular shape
- `rounded-md` - Rounded rectangle
- `shrink-0` - Prevents shrinking in flex containers

## Component Attributes

### Element Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "animate-pulse" and bg color | Yes |
| `aria-hidden` | boolean | Should be "true" for decorative elements | Recommended |

### No JavaScript Required
This is a pure CSS implementation using Tailwind's animation utilities.

## HTML Structure

```html
<!-- Basic skeleton element -->
<div class="bg-accent animate-pulse [shape-classes] [size-classes]"></div>

<!-- Skeleton group -->
<div class="[layout-classes]">
  <div class="bg-accent animate-pulse [avatar-shape]"></div>
  <div class="[content-layout]">
    <div class="bg-accent animate-pulse [text-line-size]"></div>
    <div class="bg-accent animate-pulse [text-line-size]"></div>
  </div>
</div>
```

## Examples

### Text Skeletons
```html
<!-- Single line -->
<div class="bg-accent animate-pulse h-4 w-48 rounded-md"></div>

<!-- Multiple lines -->
<div class="space-y-2">
  <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
  <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
  <div class="bg-accent animate-pulse h-4 w-3/5 rounded-md"></div>
</div>

<!-- Different text sizes -->
<div class="space-y-3">
  <!-- Heading -->
  <div class="bg-accent animate-pulse h-6 w-2/3 rounded-md"></div>
  <!-- Paragraph -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-11/12 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
  </div>
</div>
```

### Avatar Skeletons
```html
<!-- Circular avatar -->
<div class="bg-accent animate-pulse size-10 rounded-full"></div>

<!-- Square avatar -->
<div class="bg-accent animate-pulse size-10 rounded-lg"></div>

<!-- Different sizes -->
<div class="flex items-center gap-3">
  <div class="bg-accent animate-pulse size-6 rounded-full"></div>
  <div class="bg-accent animate-pulse size-8 rounded-full"></div>
  <div class="bg-accent animate-pulse size-10 rounded-full"></div>
  <div class="bg-accent animate-pulse size-12 rounded-full"></div>
</div>
```

### User Profile Skeleton
```html
<div class="flex items-center gap-4">
  <!-- Avatar -->
  <div class="bg-accent animate-pulse size-12 shrink-0 rounded-full"></div>
  <!-- Content -->
  <div class="flex-1 space-y-2">
    <!-- Name -->
    <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
    <!-- Description -->
    <div class="bg-accent animate-pulse h-3 w-24 rounded-md"></div>
  </div>
</div>

<!-- Extended profile -->
<div class="space-y-4">
  <div class="flex items-center gap-4">
    <div class="bg-accent animate-pulse size-16 shrink-0 rounded-full"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-5 w-40 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-28 rounded-md"></div>
    </div>
  </div>
  <!-- Bio -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-5/6 rounded-md"></div>
  </div>
</div>
```

### Card Skeleton
```html
<div class="card w-full">
  <header>
    <!-- Title -->
    <div class="bg-accent animate-pulse rounded-md h-4 w-2/3"></div>
    <!-- Subtitle -->
    <div class="bg-accent animate-pulse rounded-md h-4 w-1/2"></div>
  </header>
  <section>
    <!-- Main content area -->
    <div class="bg-accent animate-pulse rounded-md aspect-square w-full"></div>
  </section>
</div>

<!-- Card with actions -->
<div class="card">
  <header>
    <div class="flex justify-between items-start">
      <div class="space-y-2">
        <div class="bg-accent animate-pulse h-5 w-48 rounded-md"></div>
        <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      </div>
      <div class="bg-accent animate-pulse h-8 w-20 rounded-md"></div>
    </div>
  </header>
  <section>
    <div class="space-y-3">
      <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-3/5 rounded-md"></div>
    </div>
  </section>
  <footer>
    <div class="flex justify-between items-center">
      <div class="bg-accent animate-pulse h-8 w-16 rounded-md"></div>
      <div class="bg-accent animate-pulse h-8 w-24 rounded-md"></div>
    </div>
  </footer>
</div>
```

### Table Skeleton
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-18 rounded-md"></div></th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></td>
      </tr>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-36 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-14 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-22 rounded-md"></div></td>
      </tr>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-30 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></td>
      </tr>
    </tbody>
  </table>
</div>

<!-- Simplified table skeleton -->
<div class="space-y-3">
  <!-- Header -->
  <div class="flex gap-4">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
  </div>
  <!-- Rows -->
  <div class="space-y-2">
    <div class="flex gap-4">
      <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
    </div>
    <div class="flex gap-4">
      <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    </div>
  </div>
</div>
```

### Image Skeleton
```html
<!-- Basic image placeholder -->
<div class="bg-accent animate-pulse w-full h-48 rounded-lg"></div>

<!-- Different aspect ratios -->
<div class="space-y-4">
  <!-- Square -->
  <div class="bg-accent animate-pulse aspect-square w-full rounded-lg"></div>
  <!-- 16:9 Video -->
  <div class="bg-accent animate-pulse aspect-video w-full rounded-lg"></div>
  <!-- 4:3 Photo -->
  <div class="bg-accent animate-pulse aspect-[4/3] w-full rounded-lg"></div>
</div>

<!-- Image gallery skeleton -->
<div class="grid grid-cols-2 md:grid-cols-3 gap-4">
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
</div>
```

### Form Skeleton
```html
<div class="space-y-4">
  <!-- Form field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  
  <!-- Another field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  
  <!-- Textarea field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div>
    <div class="bg-accent animate-pulse h-24 w-full rounded-md"></div>
  </div>
  
  <!-- Submit button -->
  <div class="bg-accent animate-pulse h-10 w-24 rounded-md"></div>
</div>

<!-- Inline form -->
<div class="flex gap-3 items-end">
  <div class="flex-1 space-y-2">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  <div class="bg-accent animate-pulse h-10 w-20 rounded-md"></div>
</div>
```

### Navigation Skeleton
```html
<!-- Header navigation -->
<div class="flex items-center justify-between p-4 border-b">
  <!-- Logo -->
  <div class="bg-accent animate-pulse h-8 w-32 rounded-md"></div>
  
  <!-- Navigation links -->
  <div class="hidden md:flex items-center gap-6">
    <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-18 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
  </div>
  
  <!-- User menu -->
  <div class="bg-accent animate-pulse size-8 rounded-full"></div>
</div>

<!-- Sidebar navigation -->
<div class="w-64 space-y-2 p-4">
  <!-- Logo -->
  <div class="bg-accent animate-pulse h-6 w-28 rounded-md mb-6"></div>
  
  <!-- Nav items -->
  <div class="space-y-1">
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
  </div>
  
  <!-- Section -->
  <div class="pt-4">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md mb-2"></div>
    <div class="space-y-1">
      <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
      <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    </div>
  </div>
</div>
```

### List Skeleton
```html
<!-- Simple list -->
<div class="space-y-3">
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-48 rounded-md"></div>
  </div>
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-40 rounded-md"></div>
  </div>
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-52 rounded-md"></div>
  </div>
</div>

<!-- Detailed list items -->
<div class="space-y-4">
  <div class="flex items-start gap-3">
    <div class="bg-accent animate-pulse size-10 rounded-lg shrink-0"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-4 w-3/4 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-1/2 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-2/3 rounded-md"></div>
    </div>
    <div class="bg-accent animate-pulse h-6 w-16 rounded-md"></div>
  </div>
  
  <div class="flex items-start gap-3">
    <div class="bg-accent animate-pulse size-10 rounded-lg shrink-0"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-4 w-2/3 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-3/5 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-4/5 rounded-md"></div>
    </div>
    <div class="bg-accent animate-pulse h-6 w-20 rounded-md"></div>
  </div>
</div>
```

### Alternative Styling
```html
<!-- Using muted background -->
<div class="flex items-center gap-4">
  <div class="bg-muted animate-pulse size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-muted animate-pulse rounded-md h-4 w-[150px]"></div>
    <div class="bg-muted animate-pulse rounded-md h-4 w-[100px]"></div>
  </div>
</div>

<!-- Custom gradient skeleton -->
<style>
.skeleton-gradient {
  background: linear-gradient(-90deg, #e0e0e0 25%, #f0f0f0 50%, #e0e0e0 75%);
  background-size: 400% 100%;
  animation: loading 1.4s ease-in-out infinite;
}

@keyframes loading {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}
</style>

<div class="flex items-center gap-4">
  <div class="skeleton-gradient size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="skeleton-gradient rounded-md h-4 w-[150px]"></div>
    <div class="skeleton-gradient rounded-md h-4 w-[100px]"></div>
  </div>
</div>

<!-- Without animation (static) -->
<div class="flex items-center gap-4">
  <div class="bg-accent size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## Accessibility Features

- **Hidden from Screen Readers**: Use `aria-hidden="true"` on skeleton elements
- **Loading State**: Provide context about loading state
- **Alternative Text**: Consider providing loading feedback
- **Reduced Motion**: Respect user motion preferences

### Enhanced Accessibility
```html
<!-- With proper ARIA attributes -->
<div aria-live="polite" aria-busy="true">
  <span class="sr-only">Loading content...</span>
  <div class="flex items-center gap-4" aria-hidden="true">
    <div class="bg-accent animate-pulse size-10 shrink-0 rounded-full"></div>
    <div class="grid gap-2">
      <div class="bg-accent animate-pulse rounded-md h-4 w-[150px]"></div>
      <div class="bg-accent animate-pulse rounded-md h-4 w-[100px]"></div>
    </div>
  </div>
</div>

<!-- Reduced motion support -->
<div class="flex items-center gap-4">
  <div class="bg-accent animate-pulse motion-reduce:animate-none size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent animate-pulse motion-reduce:animate-none rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent animate-pulse motion-reduce:animate-none rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## JavaScript Integration

### React Skeleton Component
```jsx
import React from 'react';

function Skeleton({ className = '', width, height, circle = false, ...props }) {
  const baseClasses = "bg-accent animate-pulse";
  const shapeClasses = circle ? "rounded-full" : "rounded-md";
  const sizeClasses = width || height ? "" : "h-4 w-20";
  
  const style = {};
  if (width) style.width = width;
  if (height) style.height = height;
  
  return (
    <div 
      className={`${baseClasses} ${shapeClasses} ${sizeClasses} ${className}`}
      style={style}
      aria-hidden="true"
      {...props}
    />
  );
}

// Usage
<div className="flex items-center gap-4">
  <Skeleton circle width="40px" height="40px" />
  <div className="space-y-2">
    <Skeleton width="150px" />
    <Skeleton width="100px" />
  </div>
</div>

// Text skeleton component
function TextSkeleton({ lines = 1, className = "" }) {
  return (
    <div className={`space-y-2 ${className}`}>
      {Array.from({ length: lines }).map((_, i) => (
        <Skeleton 
          key={i} 
          width={i === lines - 1 ? "80%" : "100%"} 
        />
      ))}
    </div>
  );
}
```

### Vue Skeleton Component
```vue
<template>
  <div 
    :class="skeletonClasses"
    :style="skeletonStyle"
    aria-hidden="true"
    v-bind="$attrs"
  />
</template>

<script>
export default {
  props: {
    width: String,
    height: String,
    circle: Boolean,
    className: String
  },
  computed: {
    skeletonClasses() {
      const base = "bg-accent animate-pulse";
      const shape = this.circle ? "rounded-full" : "rounded-md";
      const size = this.width || this.height ? "" : "h-4 w-20";
      return `${base} ${shape} ${size} ${this.className || ""}`;
    },
    skeletonStyle() {
      const style = {};
      if (this.width) style.width = this.width;
      if (this.height) style.height = this.height;
      return style;
    }
  }
};
</script>
```

### Dynamic Skeleton Loading
```javascript
// Show/hide skeleton based on loading state
function toggleSkeleton(containerId, isLoading) {
  const container = document.getElementById(containerId);
  const skeleton = container.querySelector('.skeleton-wrapper');
  const content = container.querySelector('.content-wrapper');
  
  if (isLoading) {
    skeleton.style.display = 'block';
    content.style.display = 'none';
    container.setAttribute('aria-busy', 'true');
  } else {
    skeleton.style.display = 'none';
    content.style.display = 'block';
    container.setAttribute('aria-busy', 'false');
  }
}

// Usage
toggleSkeleton('user-profile', true);  // Show skeleton
setTimeout(() => {
  toggleSkeleton('user-profile', false); // Show content
}, 2000);
```

## Best Practices

1. **Match Content Structure**: Skeleton should mirror the layout of actual content
2. **Consistent Sizing**: Use similar dimensions to prevent layout shifts
3. **Appropriate Animation**: Use pulse animation for better UX
4. **Reduced Motion**: Respect user motion preferences
5. **Semantic HTML**: Use appropriate ARIA attributes
6. **Performance**: Avoid complex nested skeletons
7. **Visual Hierarchy**: Maintain visual hierarchy with skeleton elements
8. **Loading Context**: Provide context about what's loading

## Common Patterns

### Progressive Loading
```html
<!-- Initially show skeleton -->
<div id="content-container" aria-busy="true">
  <div class="skeleton-wrapper">
    <div class="flex items-center gap-4" aria-hidden="true">
      <div class="bg-accent animate-pulse size-12 shrink-0 rounded-full"></div>
      <div class="flex-1 space-y-2">
        <div class="bg-accent animate-pulse h-5 w-2/3 rounded-md"></div>
        <div class="bg-accent animate-pulse h-4 w-1/2 rounded-md"></div>
      </div>
    </div>
  </div>
  
  <!-- Content loads here -->
  <div class="content-wrapper" style="display: none;">
    <!-- Actual content -->
  </div>
</div>
```

### Staggered Loading
```css
/* Staggered animation delays */
.skeleton-1 { animation-delay: 0s; }
.skeleton-2 { animation-delay: 0.1s; }
.skeleton-3 { animation-delay: 0.2s; }
.skeleton-4 { animation-delay: 0.3s; }
```

```html
<div class="space-y-3">
  <div class="bg-accent animate-pulse skeleton-1 h-4 w-full rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-2 h-4 w-5/6 rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-3 h-4 w-4/5 rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-4 h-4 w-3/4 rounded-md"></div>
</div>
```

## Related Components

- [Spinner](./spinner.md) - For indeterminate loading states  
- [Progress](./progress.md) - For determinate loading states
- [Card](./card.md) - For skeleton card layouts
- [Avatar](./avatar.md) - For skeleton avatar shapes