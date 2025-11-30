# Basecoat UI Component Documentation

Complete documentation for all Basecoat UI components with examples, usage patterns, and integration guides.

## Overview

This documentation covers **34+ components** from the Basecoat UI library, providing comprehensive guides for implementation, customization, and best practices. Each component includes HTML examples, CSS classes, accessibility features, and JavaScript integration patterns.

## 📚 Component Categories

### Form Components
- **[Button](./button.md)** - Interactive buttons with multiple variants and states
- **[Button Group](./button-group.md)** - Groups of related buttons with consistent styling
- **[Input](./input.md)** - Text input fields with validation and styling options
- **[Input Group](./input-group.md)** - Input fields with icons, labels, and enhancements
- **[Textarea](./textarea.md)** - Multi-line text input areas
- **[Select](./select.md)** - Dropdown selection menus with custom styling
- **[Combobox](./combobox.md)** - Searchable select components
- **[Checkbox](./checkbox.md)** - Checkbox inputs with custom styling
- **[Radio Group](./radio-group.md)** - Radio button selections
- **[Switch](./switch.md)** - Toggle switch controls
- **[Slider](./slider.md)** - Range input controls with custom styling
- **[Field](./field.md)** - Complete form field wrapper with labels and validation
- **[Form](./form.md)** - Form container with automatic child element styling
- **[Label](./label.md)** - Form labels with proper accessibility
- **[Kbd](./kbd.md)** - Keyboard shortcut and key combination display

### Navigation Components
- **[Breadcrumb](./breadcrumb.md)** - Hierarchical navigation paths
- **[Pagination](./pagination.md)** - Page navigation controls
- **[Sidebar](./sidebar.md)** - Collapsible sidebar navigation

### Interactive Components
- **[Dialog](./dialog.md)** - Modal dialogs and overlays
- **[Alert Dialog](./alert-dialog.md)** - Confirmation and alert modals
- **[Dropdown Menu](./dropdown-menu.md)** - Context menus and dropdowns
- **[Popover](./popover.md)** - Floating content containers
- **[Tooltip](./tooltip.md)** - Hover and focus information displays
- **[Command](./command.md)** - Command palette and search interfaces
- **[Toast](./toast.md)** - Temporary notification messages
- **[Theme Switcher](./theme-switcher.md)** - Dark/light mode toggle controls

### Layout Components
- **[Card](./card.md)** - Content containers with headers and actions
- **[Accordion](./accordion.md)** - Collapsible content sections
- **[Tabs](./tabs.md)** - Tabbed content interfaces
- **[Table](./table.md)** - Data tables with sorting and styling

### Feedback Components
- **[Alert](./alert.md)** - Information and status messages
- **[Empty](./empty.md)** - Empty state displays with actions and guidance

### Utility Components
- **[Avatar](./avatar.md)** - User profile images and placeholders
- **[Badge](./badge.md)** - Status indicators and labels
- **[Spinner](./spinner.md)** - Loading indicators
- **[Progress](./progress.md)** - Progress bars and indicators
- **[Skeleton](./skeleton.md)** - Loading placeholders

### Data Components
- **[Chart](./chart.md)** - Data visualizations with Chart.js integration
- **[Carousel](./carousel.md)** - Image and content carousels

### Pattern Components
- **[Item](./item.md)** - Versatile content display patterns

## 🚀 Getting Started

### Basic Usage

Each component follows a consistent pattern:

```html
<element class="base-class variant-class size-class">
  Content
</element>
```

### Example: Button Component

```html
<!-- Basic button -->
<button class="btn">Click me</button>

<!-- Button variants -->
<button class="btn btn-outline">Outline</button>
<button class="btn btn-ghost">Ghost</button>

<!-- Button sizes -->
<button class="btn btn-sm">Small</button>
<button class="btn btn-lg">Large</button>
```

### CSS Integration

Include Basecoat CSS in your project:

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/basecoat-css@latest/dist/basecoat.css">
```

### JavaScript Integration

For interactive components, include Basecoat JavaScript:

```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@latest/dist/js/basecoat.min.js" defer></script>
```

## 🎨 Design Principles

### Atomic Design
Components are organized following atomic design principles:
- **Atoms**: Basic building blocks (Button, Input, Badge)
- **Molecules**: Simple combinations (Form Field, Search Bar)
- **Organisms**: Complex components (Data Table, Navigation)
- **Templates**: Page layouts and structures
- **Pages**: Complete interfaces

### Theme Integration
All components use CSS custom properties for theming:

```css
:root {
  --primary: 217 91% 60%;
  --primary-foreground: 0 0% 100%;
  --background: 0 0% 100%;
  --foreground: 0 0% 9%;
  --muted: 0 0% 96%;
  --muted-foreground: 0 0% 45%;
  --border: 0 0% 90%;
  --ring: 217 91% 60%;
}
```

### Accessibility First
Every component includes:
- Proper semantic HTML
- ARIA attributes and roles
- Keyboard navigation support
- Screen reader compatibility
- High contrast mode support

## 💻 Framework Integration

### React Example

```jsx
import React from 'react';

function MyButton({ variant = 'default', size = 'md', children, ...props }) {
  const classes = [
    'btn',
    variant !== 'default' && `btn-${variant}`,
    size !== 'md' && `btn-${size}`
  ].filter(Boolean).join(' ');
  
  return (
    <button className={classes} {...props}>
      {children}
    </button>
  );
}
```

### Vue Example

```vue
<template>
  <button :class="buttonClasses" v-bind="$attrs">
    <slot />
  </button>
</template>

<script>
export default {
  props: {
    variant: { type: String, default: 'default' },
    size: { type: String, default: 'md' }
  },
  computed: {
    buttonClasses() {
      return [
        'btn',
        this.variant !== 'default' && `btn-${this.variant}`,
        this.size !== 'md' && `btn-${this.size}`
      ].filter(Boolean).join(' ');
    }
  }
};
</script>
```

## 🎯 Common Patterns

### Form Layout

```html
<form class="form grid gap-6">
  <div class="grid gap-2">
    <label for="username">Username</label>
    <input type="text" id="username" placeholder="Enter username">
  </div>
  
  <div class="grid gap-2">
    <label for="email">Email</label>
    <input type="email" id="email" placeholder="Enter email">
  </div>
  
  <button type="submit" class="btn">Submit</button>
</form>
```

### Card with Actions

```html
<div class="card p-6">
  <header class="mb-4">
    <h3 class="text-lg font-semibold">Card Title</h3>
    <p class="text-muted-foreground">Card description</p>
  </header>
  
  <div class="content">
    <!-- Card content -->
  </div>
  
  <footer class="flex gap-2 mt-4">
    <button class="btn">Primary Action</button>
    <button class="btn-outline">Secondary</button>
  </footer>
</div>
```

### Data Display

```html
<div class="space-y-4">
  <!-- Stats -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <div class="card p-4 text-center">
      <div class="text-2xl font-bold">1,234</div>
      <div class="text-sm text-muted-foreground">Total Users</div>
    </div>
  </div>
  
  <!-- Table -->
  <div class="card">
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>John Doe</td>
          <td>john@example.com</td>
          <td><span class="badge badge-success">Active</span></td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
```

## 🔧 Customization

### CSS Custom Properties

Override theme variables for customization:

```css
:root {
  --primary: 142 76% 36%;     /* Custom green */
  --secondary: 221 83% 53%;   /* Custom blue */
  --radius: 0.75rem;          /* More rounded corners */
}
```

### Component Variants

Create custom component variants:

```css
.btn-custom {
  @apply bg-gradient-to-r from-purple-500 to-pink-500 text-white;
}

.btn-custom:hover {
  @apply from-purple-600 to-pink-600;
}
```

### Responsive Breakpoints

Use Tailwind's responsive prefixes:

```html
<button class="btn btn-sm md:btn-md lg:btn-lg">
  Responsive Button
</button>
```

## 📱 Mobile Considerations

### Touch-Friendly Sizes

```html
<!-- Minimum 44px touch target -->
<button class="btn min-h-11 px-4">Mobile Button</button>
```

### Responsive Typography

```html
<h1 class="text-xl md:text-2xl lg:text-3xl">
  Responsive Heading
</h1>
```

### Mobile Navigation

```html
<nav class="sidebar md:sidebar-desktop">
  <!-- Navigation content -->
</nav>
```

## ♿ Accessibility Guidelines

### Semantic HTML
Always use appropriate HTML elements:

```html
<!-- Good -->
<button class="btn">Submit</button>
<nav class="breadcrumb">...</nav>

<!-- Avoid -->
<div class="btn" onclick="submit()">Submit</div>
```

### ARIA Labels

```html
<button class="btn-icon" aria-label="Close dialog">
  <svg><!-- close icon --></svg>
</button>
```

### Focus Management

```html
<div class="dialog" role="dialog" aria-labelledby="dialog-title">
  <h2 id="dialog-title">Dialog Title</h2>
  <!-- Dialog content -->
</div>
```

### Color Contrast

Ensure sufficient contrast ratios:
- Normal text: 4.5:1 minimum
- Large text: 3:1 minimum
- Interactive elements: 3:1 minimum

## 🚀 Performance Tips

### CSS Optimization

```css
/* Use CSS containment for performance */
.card {
  contain: layout style paint;
}

/* Optimize animations */
.btn {
  will-change: transform;
  transition: transform 0.2s ease;
}
```

### JavaScript Best Practices

```javascript
// Use event delegation
document.addEventListener('click', (e) => {
  if (e.target.matches('.btn')) {
    // Handle button click
  }
});

// Debounce search inputs
const debounce = (fn, delay) => {
  let timeoutId;
  return (...args) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn.apply(null, args), delay);
  };
};
```

## 📋 Best Practices Checklist

### Component Usage
- [ ] Use semantic HTML elements
- [ ] Include proper ARIA attributes
- [ ] Test keyboard navigation
- [ ] Verify screen reader compatibility
- [ ] Check color contrast ratios
- [ ] Test on multiple devices/browsers
- [ ] Validate responsive behavior

### Development
- [ ] Follow consistent naming conventions
- [ ] Use appropriate component variants
- [ ] Include loading and error states
- [ ] Handle edge cases gracefully
- [ ] Optimize for performance
- [ ] Document component usage
- [ ] Test with real data

### Design
- [ ] Maintain visual hierarchy
- [ ] Use consistent spacing
- [ ] Follow brand guidelines
- [ ] Consider user workflow
- [ ] Provide clear feedback
- [ ] Support multiple themes
- [ ] Plan for internationalization

## 🔗 Resources

- **[Basecoat GitHub](https://github.com/hunvreus/basecoat)** - Source code and issues
- **[Basecoat Website](https://basecoatui.com)** - Official documentation
- **[Tailwind CSS](https://tailwindcss.com)** - Utility-first CSS framework
- **[Web Accessibility Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)** - WCAG 2.1 quick reference

## 📄 License

This documentation is provided under the same license as the Basecoat UI library. Please refer to the main project repository for license details.

---

**Last Updated**: November 2024
**Version**: Compatible with Basecoat v0.3.6+