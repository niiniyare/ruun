# Kbd Component

Used to display textual user input from keyboard, including shortcuts and key combinations.

## Basic Usage

```html
<kbd class="kbd">K</kbd>
```

## CSS Classes

### Base Classes
- **`kbd`** - Base keyboard key styling with muted background and proper typography

### Size Variants
- **`kbd-sm`** - Smaller keyboard key (text-xs, reduced padding)
- **`kbd-lg`** - Larger keyboard key (increased padding)

### Style Variants
- **`kbd-outline`** - Outline style with border
- **`kbd-solid`** - Solid background style
- **`kbd-accent`** - Accent colored style

### State Classes
- **`kbd-disabled`** - Disabled state with reduced opacity
- **`kbd-active`** - Active/pressed state

### Tailwind Utilities Used
- **`inline-flex items-center justify-center`** - Layout and alignment
- **`min-h-5 h-5 px-1.5 py-0.5`** - Size and spacing
- **`text-xs font-mono font-medium`** - Typography
- **`bg-muted text-muted-foreground`** - Colors
- **`border border-border`** - Border styling
- **`rounded shadow-xs`** - Shape and shadow

## Component Attributes

### Standard HTML Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "kbd" base class | Yes |

### Data Attributes (Optional)
| Attribute | Values | Description | Required |
|-----------|--------|-------------|----------|
| `data-key` | "cmd", "shift", "option", "ctrl", "enter", "esc", "tab", "space", "arrow-up", "arrow-down", "arrow-left", "arrow-right" | Shows corresponding key symbol | No |

## No JavaScript Required
The kbd component is purely presentational and requires no JavaScript.

## HTML Structure

```html
<!-- Basic key -->
<kbd class="kbd">Key</kbd>

<!-- Key combination -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">Ctrl</kbd>
  <span class="text-muted-foreground">+</span>
  <kbd class="kbd">B</kbd>
</span>

<!-- With data attributes -->
<kbd class="kbd" data-key="cmd"></kbd>
```

## Examples

### Single Keys

```html
<!-- Letter keys -->
<kbd class="kbd">K</kbd>
<kbd class="kbd">A</kbd>
<kbd class="kbd">Z</kbd>

<!-- Number keys -->
<kbd class="kbd">1</kbd>
<kbd class="kbd">2</kbd>
<kbd class="kbd">0</kbd>

<!-- Special keys -->
<kbd class="kbd">Enter</kbd>
<kbd class="kbd">Esc</kbd>
<kbd class="kbd">Tab</kbd>
<kbd class="kbd">Space</kbd>
```

### Symbol Keys

```html
<!-- Direct symbols -->
<kbd class="kbd">⌘</kbd>
<kbd class="kbd">⇧</kbd>
<kbd class="kbd">⌥</kbd>
<kbd class="kbd">⌃</kbd>
<kbd class="kbd">⏎</kbd>

<!-- Using data attributes -->
<kbd class="kbd" data-key="cmd"></kbd>
<kbd class="kbd" data-key="shift"></kbd>
<kbd class="kbd" data-key="option"></kbd>
<kbd class="kbd" data-key="ctrl"></kbd>
<kbd class="kbd" data-key="enter"></kbd>
```

### Key Combinations

```html
<!-- Text combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">Ctrl</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">B</kbd>
</span>

<!-- Symbol combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <kbd class="kbd">K</kbd>
</span>

<!-- Complex combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</span>
```

### Size Variants

```html
<!-- Small -->
<kbd class="kbd kbd-sm">⌘</kbd>
<kbd class="kbd kbd-sm">K</kbd>

<!-- Default -->
<kbd class="kbd">Enter</kbd>

<!-- Large -->
<kbd class="kbd kbd-lg">Space</kbd>
```

### Style Variants

```html
<!-- Default (muted) -->
<kbd class="kbd">Ctrl</kbd>

<!-- Outline -->
<kbd class="kbd kbd-outline">Alt</kbd>

<!-- Solid -->
<kbd class="kbd kbd-solid">Shift</kbd>

<!-- Accent -->
<kbd class="kbd kbd-accent">⌘</kbd>
```

### Arrow Keys

```html
<div class="inline-flex items-center gap-1">
  <kbd class="kbd" data-key="arrow-up"></kbd>
  <kbd class="kbd" data-key="arrow-down"></kbd>
  <kbd class="kbd" data-key="arrow-left"></kbd>
  <kbd class="kbd" data-key="arrow-right"></kbd>
</div>

<!-- Or with direct symbols -->
<div class="inline-flex items-center gap-1">
  <kbd class="kbd">↑</kbd>
  <kbd class="kbd">↓</kbd>
  <kbd class="kbd">←</kbd>
  <kbd class="kbd">→</kbd>
</div>
```

### In Buttons

```html
<!-- Button with keyboard shortcut -->
<button class="btn-sm-outline inline-flex items-center gap-2">
  Save
  <kbd class="kbd kbd-sm">⌘S</kbd>
</button>

<button class="btn-ghost inline-flex items-center gap-2">
  Search
  <kbd class="kbd kbd-sm">⌘K</kbd>
</button>

<!-- Accept/Cancel actions -->
<div class="flex gap-2">
  <button class="btn-sm inline-flex items-center gap-2">
    Accept
    <kbd class="kbd kbd-sm">⏎</kbd>
  </button>
  <button class="btn-sm-outline inline-flex items-center gap-2">
    Cancel
    <kbd class="kbd kbd-sm">Esc</kbd>
  </button>
</div>
```

### In Tooltips

```html
<button 
  class="btn-icon-outline" 
  data-tooltip="Bold (⌘B)"
  data-side="bottom"
>
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>
    <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>
  </svg>
</button>

<!-- Or in tooltip content -->
<div class="tooltip-content">
  <div class="font-medium">Bold</div>
  <div class="text-xs text-muted-foreground mt-1">
    Press <kbd class="kbd kbd-sm">⌘</kbd> <kbd class="kbd kbd-sm">B</kbd>
  </div>
</div>
```

### Help Documentation

```html
<div class="space-y-4">
  <h3 class="font-semibold">Keyboard Shortcuts</h3>
  <dl class="space-y-2">
    <div class="flex items-center justify-between">
      <dt>Open command palette</dt>
      <dd><kbd class="kbd">⌘K</kbd></dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Save document</dt>
      <dd><kbd class="kbd">⌘S</kbd></dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Bold text</dt>
      <dd>
        <span class="inline-flex items-center gap-1">
          <kbd class="kbd">⌘</kbd>
          <kbd class="kbd">B</kbd>
        </span>
      </dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Find and replace</dt>
      <dd>
        <span class="inline-flex items-center gap-1">
          <kbd class="kbd">⌘</kbd>
          <span class="text-muted-foreground text-xs">+</span>
          <kbd class="kbd">⇧</kbd>
          <span class="text-muted-foreground text-xs">+</span>
          <kbd class="kbd">F</kbd>
        </span>
      </dd>
    </div>
  </dl>
</div>
```

### Command Menu Items

```html
<div class="command-menu">
  <div class="command-item flex items-center justify-between p-2 rounded hover:bg-muted">
    <div class="flex items-center gap-3">
      <svg class="w-4 h-4"><!-- icon --></svg>
      <span>Open file</span>
    </div>
    <kbd class="kbd kbd-sm">⌘O</kbd>
  </div>
  
  <div class="command-item flex items-center justify-between p-2 rounded hover:bg-muted">
    <div class="flex items-center gap-3">
      <svg class="w-4 h-4"><!-- icon --></svg>
      <span>Close tab</span>
    </div>
    <kbd class="kbd kbd-sm">⌘W</kbd>
  </div>
</div>
```

### State Demonstrations

```html
<!-- Disabled state -->
<kbd class="kbd kbd-disabled">Unavailable</kbd>

<!-- Active/pressed state -->
<kbd class="kbd kbd-active">Pressed</kbd>

<!-- In different contexts -->
<div class="space-y-2">
  <p>Press <kbd class="kbd">Space</kbd> to continue</p>
  <p class="text-sm text-muted-foreground">
    Use <kbd class="kbd kbd-sm">Tab</kbd> and <kbd class="kbd kbd-sm">⇧Tab</kbd> to navigate
  </p>
</div>
```

### Platform-Specific Keys

```html
<!-- Mac shortcuts -->
<div class="mac-shortcuts space-y-2">
  <div class="flex items-center justify-between">
    <span>Copy</span>
    <kbd class="kbd">⌘C</kbd>
  </div>
  <div class="flex items-center justify-between">
    <span>Paste</span>
    <kbd class="kbd">⌘V</kbd>
  </div>
</div>

<!-- Windows/Linux shortcuts -->
<div class="pc-shortcuts space-y-2">
  <div class="flex items-center justify-between">
    <span>Copy</span>
    <kbd class="kbd">Ctrl+C</kbd>
  </div>
  <div class="flex items-center justify-between">
    <span>Paste</span>
    <kbd class="kbd">Ctrl+V</kbd>
  </div>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper `<kbd>` element for keyboard input
- **Screen Reader Support**: Screen readers announce kbd content appropriately
- **High Contrast**: Sufficient contrast in both light and dark modes
- **Font Choice**: Monospace font for consistent character width

### Enhanced Accessibility

```html
<!-- With descriptive text -->
<p>
  To save your work, press 
  <kbd class="kbd" title="Command key">⌘</kbd>
  <kbd class="kbd" title="Letter S key">S</kbd>
</p>

<!-- With aria-label for complex combinations -->
<span 
  class="inline-flex items-center gap-1"
  aria-label="Command plus Shift plus P"
>
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</span>

<!-- In help content with proper structure -->
<section aria-labelledby="shortcuts-heading">
  <h3 id="shortcuts-heading">Available Shortcuts</h3>
  <dl>
    <dt>Save document</dt>
    <dd>
      <kbd class="kbd" aria-label="Command S">⌘S</kbd>
    </dd>
  </dl>
</section>
```

## JavaScript Integration

### Dynamic Key Display

```javascript
// Show platform-appropriate shortcuts
function getShortcutDisplay(action) {
  const isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
  
  const shortcuts = {
    save: isMac ? '⌘S' : 'Ctrl+S',
    copy: isMac ? '⌘C' : 'Ctrl+C',
    paste: isMac ? '⌘V' : 'Ctrl+V',
    find: isMac ? '⌘F' : 'Ctrl+F'
  };
  
  return shortcuts[action] || '';
}

// Usage
document.addEventListener('DOMContentLoaded', () => {
  document.querySelectorAll('[data-shortcut]').forEach(el => {
    const action = el.dataset.shortcut;
    const kbd = el.querySelector('.kbd');
    if (kbd) {
      kbd.textContent = getShortcutDisplay(action);
    }
  });
});
```

### Keyboard Event Handling

```javascript
// Listen for shortcuts and highlight corresponding kbd elements
document.addEventListener('keydown', (e) => {
  const keys = [];
  
  if (e.metaKey || e.ctrlKey) keys.push(navigator.platform.includes('Mac') ? '⌘' : 'Ctrl');
  if (e.shiftKey) keys.push('⇧');
  if (e.altKey) keys.push(navigator.platform.includes('Mac') ? '⌥' : 'Alt');
  
  keys.push(e.key.toUpperCase());
  
  const combination = keys.join('+');
  
  // Find matching kbd elements and highlight temporarily
  document.querySelectorAll('.kbd').forEach(kbd => {
    if (kbd.textContent === combination) {
      kbd.classList.add('kbd-active');
      setTimeout(() => kbd.classList.remove('kbd-active'), 150);
    }
  });
});
```

### React Integration

```jsx
import React from 'react';

function Kbd({ 
  children, 
  dataKey, 
  size = 'default', 
  variant = 'default',
  disabled = false,
  className = '' 
}) {
  const keySymbols = {
    cmd: '⌘',
    shift: '⇧',
    option: '⌥',
    ctrl: '⌃',
    enter: '⏎',
    esc: '⎋',
    tab: '⇥',
    space: '␣',
    'arrow-up': '↑',
    'arrow-down': '↓',
    'arrow-left': '←',
    'arrow-right': '→'
  };

  const sizeClasses = {
    sm: 'kbd-sm',
    default: '',
    lg: 'kbd-lg'
  };

  const variantClasses = {
    default: '',
    outline: 'kbd-outline',
    solid: 'kbd-solid',
    accent: 'kbd-accent'
  };

  const classes = [
    'kbd',
    sizeClasses[size],
    variantClasses[variant],
    disabled && 'kbd-disabled',
    className
  ].filter(Boolean).join(' ');

  const content = dataKey ? keySymbols[dataKey] || dataKey : children;

  return (
    <kbd className={classes}>
      {content}
    </kbd>
  );
}

// Usage
function ShortcutDisplay() {
  return (
    <div className="space-y-4">
      <div className="flex items-center gap-2">
        <span>Save:</span>
        <Kbd dataKey="cmd" />
        <Kbd>S</Kbd>
      </div>
      
      <button className="btn inline-flex items-center gap-2">
        Search
        <Kbd size="sm" dataKey="cmd" />
        <Kbd size="sm">K</Kbd>
      </button>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <kbd :class="kbdClasses">
    {{ displayText }}
  </kbd>
</template>

<script>
export default {
  props: {
    dataKey: String,
    size: {
      type: String,
      default: 'default',
      validator: value => ['sm', 'default', 'lg'].includes(value)
    },
    variant: {
      type: String,
      default: 'default',
      validator: value => ['default', 'outline', 'solid', 'accent'].includes(value)
    },
    disabled: Boolean
  },
  computed: {
    keySymbols() {
      return {
        cmd: '⌘',
        shift: '⇧',
        option: '⌥',
        ctrl: '⌃',
        enter: '⏎',
        esc: '⎋',
        tab: '⇥',
        space: '␣'
      };
    },
    kbdClasses() {
      return [
        'kbd',
        this.size !== 'default' && `kbd-${this.size}`,
        this.variant !== 'default' && `kbd-${this.variant}`,
        this.disabled && 'kbd-disabled'
      ].filter(Boolean).join(' ');
    },
    displayText() {
      if (this.dataKey) {
        return this.keySymbols[this.dataKey] || this.dataKey;
      }
      return this.$slots.default?.[0]?.children || '';
    }
  }
};
</script>
```

## Best Practices

1. **Consistent Symbols**: Use standard symbols (⌘, ⇧, ⌥, ⌃) for modifier keys
2. **Platform Awareness**: Show appropriate shortcuts for user's platform
3. **Logical Grouping**: Group related keys with proper spacing
4. **Context Appropriate**: Use smaller sizes in compact interfaces
5. **Accessibility**: Include descriptive text or aria-labels
6. **Visual Hierarchy**: Don't overuse accent variants
7. **Readability**: Ensure sufficient contrast in all themes

## Common Patterns

### Shortcut List

```html
<div class="space-y-3">
  <div class="flex items-center justify-between py-2">
    <span>New file</span>
    <kbd class="kbd">⌘N</kbd>
  </div>
  <div class="flex items-center justify-between py-2">
    <span>Open file</span>
    <kbd class="kbd">⌘O</kbd>
  </div>
  <div class="flex items-center justify-between py-2">
    <span>Save file</span>
    <kbd class="kbd">⌘S</kbd>
  </div>
</div>
```

### Inline Instructions

```html
<p class="text-sm text-muted-foreground">
  Press <kbd class="kbd kbd-sm">Enter</kbd> to submit or 
  <kbd class="kbd kbd-sm">Esc</kbd> to cancel.
</p>
```

### Complex Shortcuts

```html
<div class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</div>
```

## Related Components

- [Button](./button.md) - Often contains kbd elements
- [Tooltip](./tooltip.md) - For showing shortcuts on hover
- [Command](./command.md) - For command palette interfaces
- [Badge](./badge.md) - Similar styling patterns