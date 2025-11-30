# Avatar Component

An image element with a fallback for representing the user.

## Important Note

**There is no dedicated Avatar component in Basecoat.** Avatars are simply `<img>` elements styled with Tailwind utility classes.

## Basic Usage

```html
<img class="size-8 shrink-0 object-cover rounded-full" alt="@hunvreus" src="https://github.com/hunvreus.png" />
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard Tailwind utilities

### Supporting Classes
- **`size-*`** - Sets width and height (size-6, size-8, size-10, size-12, etc.)
- **`shrink-0`** - Prevents image from shrinking in flex containers
- **`object-cover`** - Ensures proper image scaling within container
- **`rounded-*`** - Border radius (rounded-full, rounded-lg, rounded-md)

### Tailwind Utilities Used
- `size-8` - Sets width and height to 2rem (32px)
- `shrink-0` - Prevents flex shrinking
- `object-cover` - Cover object fit for proper image scaling
- `rounded-full` - Circular avatar shape
- `rounded-lg` - Rounded rectangle shape
- `ring-*` - Ring/border effects for grouped avatars
- `grayscale` - Desaturated color effect
- `-space-x-*` - Negative horizontal spacing for overlapped avatars

## Component Attributes

### Image Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `src` | string | Image source URL | Yes |
| `alt` | string | Alternative text for accessibility | Yes |
| `class` | string | Tailwind utility classes | Yes |

### No JavaScript Required
This is a pure CSS/HTML implementation using standard image elements.

## HTML Structure

```html
<img class="[size-classes] [shape-classes] [object-classes]" alt="[description]" src="[image-url]" />
```

## Examples

### Different Sizes
```html
<!-- Extra Small (24px) -->
<img class="size-6 shrink-0 object-cover rounded-full" alt="Small avatar" src="https://github.com/hunvreus.png" />

<!-- Small (32px) -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="Medium avatar" src="https://github.com/hunvreus.png" />

<!-- Medium (40px) -->
<img class="size-10 shrink-0 object-cover rounded-full" alt="Medium avatar" src="https://github.com/hunvreus.png" />

<!-- Large (48px) -->
<img class="size-12 shrink-0 object-cover rounded-full" alt="Large avatar" src="https://github.com/hunvreus.png" />

<!-- Extra Large (64px) -->
<img class="size-16 shrink-0 object-cover rounded-full" alt="Extra large avatar" src="https://github.com/hunvreus.png" />
```

### Different Shapes
```html
<!-- Circular (most common) -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="Circular avatar" src="https://github.com/hunvreus.png" />

<!-- Rounded rectangle -->
<img class="size-8 shrink-0 object-cover rounded-lg" alt="Rounded avatar" src="https://github.com/shadcn.png" />

<!-- Square -->
<img class="size-8 shrink-0 object-cover rounded-md" alt="Square avatar" src="https://github.com/hunvreus.png" />

<!-- No rounding -->
<img class="size-8 shrink-0 object-cover" alt="Square avatar" src="https://github.com/hunvreus.png" />
```

### Avatar Group (Overlapping)
```html
<div class="flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full">
  <img alt="@hunvreus" src="https://github.com/hunvreus.png" />
  <img alt="@shadcn" src="https://github.com/shadcn.png" />
  <img alt="@adamwathan" src="https://github.com/adamwathan.png" />
</div>
```

### Avatar with Ring/Border
```html
<!-- Simple ring -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-primary" alt="Avatar with ring" src="https://github.com/hunvreus.png" />

<!-- Ring with offset -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-primary ring-offset-2 ring-offset-background" alt="Avatar with offset ring" src="https://github.com/hunvreus.png" />

<!-- Status indicator ring -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-success" alt="Online user" src="https://github.com/hunvreus.png" />
```

### Avatar with Fallback Text
```html
<!-- Using a div for text fallback when no image -->
<div class="size-8 shrink-0 rounded-full bg-primary text-primary-foreground flex items-center justify-center text-sm font-medium">
  JD
</div>

<!-- Using initials for fallback -->
<div class="size-10 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-sm font-semibold">
  AB
</div>
```

### Avatar with Status Indicator
```html
<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Online status -->
  <div class="absolute bottom-0 right-0 size-3 bg-success rounded-full ring-2 ring-background"></div>
</div>

<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Away status -->
  <div class="absolute bottom-0 right-0 size-3 bg-warning rounded-full ring-2 ring-background"></div>
</div>

<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Offline status -->
  <div class="absolute bottom-0 right-0 size-3 bg-muted rounded-full ring-2 ring-background"></div>
</div>
```

### Avatar Sizes Reference
```html
<!-- Size reference -->
<div class="flex items-center gap-4">
  <!-- xs: 16px -->
  <img class="size-4 shrink-0 object-cover rounded-full" alt="Extra small" src="https://github.com/hunvreus.png" />
  
  <!-- sm: 20px -->
  <img class="size-5 shrink-0 object-cover rounded-full" alt="Small" src="https://github.com/hunvreus.png" />
  
  <!-- md: 24px -->
  <img class="size-6 shrink-0 object-cover rounded-full" alt="Medium" src="https://github.com/hunvreus.png" />
  
  <!-- lg: 32px -->
  <img class="size-8 shrink-0 object-cover rounded-full" alt="Large" src="https://github.com/hunvreus.png" />
  
  <!-- xl: 40px -->
  <img class="size-10 shrink-0 object-cover rounded-full" alt="Extra large" src="https://github.com/hunvreus.png" />
  
  <!-- 2xl: 48px -->
  <img class="size-12 shrink-0 object-cover rounded-full" alt="2X large" src="https://github.com/hunvreus.png" />
</div>
```

### Avatar with Custom Dimensions
```html
<!-- Custom width/height -->
<img class="w-20 h-20 shrink-0 object-cover rounded-full" alt="Custom size avatar" src="https://github.com/hunvreus.png" />

<!-- Aspect ratio preservation -->
<img class="w-16 aspect-square shrink-0 object-cover rounded-lg" alt="Square aspect ratio" src="https://github.com/hunvreus.png" />
```

### Avatar Group with Count Indicator
```html
<div class="flex items-center">
  <!-- Avatar group -->
  <div class="flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full">
    <img alt="User 1" src="https://github.com/hunvreus.png" />
    <img alt="User 2" src="https://github.com/shadcn.png" />
    <img alt="User 3" src="https://github.com/adamwathan.png" />
  </div>
  
  <!-- Count indicator -->
  <div class="ml-2 size-8 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-xs font-medium">
    +5
  </div>
</div>
```

### Clickable Avatar
```html
<!-- Link avatar -->
<a href="/profile" class="block">
  <img class="size-8 shrink-0 object-cover rounded-full hover:ring-2 hover:ring-primary transition-all" alt="Profile" src="https://github.com/hunvreus.png" />
</a>

<!-- Button avatar -->
<button type="button" class="block">
  <img class="size-8 shrink-0 object-cover rounded-full hover:ring-2 hover:ring-primary transition-all" alt="User menu" src="https://github.com/hunvreus.png" />
</button>
```

## Accessibility Features

- **Alt Text**: Always provide descriptive alternative text
- **Semantic HTML**: Uses standard `<img>` elements
- **Focus States**: Support for focus indicators when clickable
- **Screen Reader Support**: Proper alternative text announcements

### Enhanced Accessibility
```html
<!-- Descriptive alt text -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="John Smith, Software Engineer" src="/avatars/john-smith.jpg" />

<!-- Role for decorative avatars -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="" role="presentation" src="/avatars/decorative.jpg" />

<!-- With ARIA label for interactive avatars -->
<button type="button" aria-label="Open user menu for John Smith">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="" src="/avatars/john-smith.jpg" />
</button>
```

## Best Practices

1. **Consistent Sizing**: Use consistent avatar sizes within the same context
2. **Meaningful Alt Text**: Provide descriptive alternative text
3. **Loading States**: Consider placeholder or loading states for slow networks
4. **Fallback Content**: Provide fallbacks for missing images
5. **Performance**: Optimize images for file size and format
6. **Responsive**: Consider different sizes for different screen sizes
7. **Accessibility**: Ensure proper contrast for text fallbacks
8. **Error Handling**: Handle broken image URLs gracefully

## Common Patterns

### User Profile Card
```html
<div class="flex items-center gap-3 p-4 border rounded-lg">
  <img class="size-12 shrink-0 object-cover rounded-full" alt="John Doe" src="https://github.com/hunvreus.png" />
  <div>
    <h3 class="font-medium">John Doe</h3>
    <p class="text-sm text-muted-foreground">Software Engineer</p>
  </div>
</div>
```

### Comment Thread
```html
<div class="flex gap-3">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="Sarah Wilson" src="https://github.com/shadcn.png" />
  <div class="flex-1">
    <div class="flex items-center gap-2">
      <span class="font-medium text-sm">Sarah Wilson</span>
      <span class="text-xs text-muted-foreground">2 hours ago</span>
    </div>
    <p class="text-sm mt-1">This looks great! Thanks for the update.</p>
  </div>
</div>
```

### Team List
```html
<div class="space-y-3">
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-3">
      <img class="size-10 shrink-0 object-cover rounded-full" alt="Alice Johnson" src="https://github.com/hunvreus.png" />
      <div>
        <h4 class="font-medium">Alice Johnson</h4>
        <p class="text-sm text-muted-foreground">Team Lead</p>
      </div>
    </div>
    <div class="relative">
      <div class="size-3 bg-success rounded-full"></div>
    </div>
  </div>
</div>
```

## Error Handling

### Image Fallback with JavaScript
```html
<img 
  class="size-8 shrink-0 object-cover rounded-full" 
  alt="User avatar" 
  src="https://example.com/avatar.jpg"
  onerror="this.style.display='none'; this.nextElementSibling.style.display='flex'"
/>
<div 
  class="size-8 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-sm font-medium" 
  style="display: none;"
>
  UA
</div>
```

### CSS-Only Fallback
```html
<div class="size-8 shrink-0 rounded-full overflow-hidden bg-muted flex items-center justify-center">
  <img 
    class="w-full h-full object-cover" 
    alt="User avatar" 
    src="https://example.com/avatar.jpg"
    style="display: block;"
    onerror="this.style.display='none'"
  />
  <span class="text-sm font-medium text-muted-foreground">UA</span>
</div>
```

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Avatar({ src, alt, size = 8, shape = 'rounded-full', className = '', ...props }) {
  const sizeClass = `size-${size}`;
  
  return (
    <img
      className={`${sizeClass} shrink-0 object-cover ${shape} ${className}`}
      alt={alt}
      src={src}
      {...props}
    />
  );
}

// Usage
<Avatar src="https://github.com/hunvreus.png" alt="User" size={10} />
```

### Vue Integration
```vue
<template>
  <img 
    :class="avatarClasses"
    :alt="alt"
    :src="src"
    v-bind="$attrs"
  />
</template>

<script>
export default {
  props: {
    src: String,
    alt: String,
    size: {
      type: Number,
      default: 8
    },
    shape: {
      type: String,
      default: 'rounded-full'
    }
  },
  computed: {
    avatarClasses() {
      return `size-${this.size} shrink-0 object-cover ${this.shape}`;
    }
  }
};
</script>
```

### Avatar with Loading State
```html
<!-- Loading skeleton -->
<div class="size-8 shrink-0 rounded-full bg-muted animate-pulse"></div>

<!-- Loaded avatar -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="User" src="https://github.com/hunvreus.png" />
```

## Size Guide

| Class | Size | Pixels | Use Case |
|-------|------|--------|----------|
| `size-4` | 1rem | 16px | Tiny avatars, icons |
| `size-5` | 1.25rem | 20px | Small inline avatars |
| `size-6` | 1.5rem | 24px | Default small size |
| `size-8` | 2rem | 32px | Standard avatar size |
| `size-10` | 2.5rem | 40px | Medium avatars |
| `size-12` | 3rem | 48px | Large avatars |
| `size-16` | 4rem | 64px | Profile headers |
| `size-20` | 5rem | 80px | Hero avatars |

## Related Components

- [Button](./button.md) - For clickable avatar functionality
- [Badge](./badge.md) - For status indicators on avatars
- [Card](./card.md) - For profile cards containing avatars
- [Dropdown Menu](./dropdown-menu.md) - For avatar-triggered menus