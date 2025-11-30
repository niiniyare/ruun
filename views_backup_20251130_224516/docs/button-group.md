# Button Group Component

A container that groups related buttons together with consistent styling.

## Basic Usage

```html
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Archive</button>
  <button type="button" class="btn-outline">Report</button>
</div>
```

## CSS Classes

### Primary Classes
- **`button-group`** - Core button group class for connected buttons

### Supporting Classes
- **Layout**: `flex`, `w-fit`, `items-stretch`, `gap-2`
- **Grouping**: `role="group"` for accessibility

### Tailwind Utilities Used
- `flex` - Flexbox layout
- `w-fit` - Width fits content
- `items-stretch` - Equal height buttons
- `gap-2` - Spacing between groups
- `button-group` - Connected button styling

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include layout classes | Yes |
| `role` | string | "group" for accessibility | Recommended |

### Button Group Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | "button-group" class | Yes |
| `role` | string | "group" for semantic grouping | Recommended |

### No JavaScript Required (Basic)
Basic button groups work with pure CSS styling.

## HTML Structure

```html
<!-- Container for multiple groups -->
<div class="flex w-fit items-stretch gap-2">
  <!-- Individual button groups -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Button 1</button>
    <button type="button" class="btn-outline">Button 2</button>
  </div>
</div>
```

## Examples

### Basic Button Groups

```html
<!-- Simple two-button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Save</button>
  <button type="button" class="btn-outline">Cancel</button>
</div>

<!-- Multiple button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Bold</button>
  <button type="button" class="btn-outline">Italic</button>
  <button type="button" class="btn-outline">Underline</button>
</div>

<!-- Mixed button types -->
<div role="group" class="button-group">
  <button type="button" class="btn">Primary</button>
  <button type="button" class="btn-outline">Secondary</button>
</div>
```

### Multiple Button Groups

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- Back button (standalone) -->
  <button type="button" class="btn-icon-outline" aria-label="Go Back">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m12 19-7-7 7-7" />
      <path d="M19 12H5" />
    </svg>
  </button>

  <!-- First button group -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Archive</button>
    <button type="button" class="btn-outline">Report</button>
  </div>

  <!-- Second button group with dropdown -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Snooze</button>
    <button type="button" class="btn-icon-outline" aria-label="More options">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
    </button>
  </div>
</div>
```

### Different Button Variants

```html
<!-- Primary button group -->
<div role="group" class="button-group">
  <button type="button" class="btn">Save</button>
  <button type="button" class="btn">Continue</button>
  <button type="button" class="btn">Finish</button>
</div>

<!-- Outline button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Left</button>
  <button type="button" class="btn-outline">Center</button>
  <button type="button" class="btn-outline">Right</button>
</div>

<!-- Secondary button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-secondary">Option A</button>
  <button type="button" class="btn-secondary">Option B</button>
  <button type="button" class="btn-secondary">Option C</button>
</div>

<!-- Ghost button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-ghost">Draft</button>
  <button type="button" class="btn-ghost">Published</button>
  <button type="button" class="btn-ghost">Archived</button>
</div>
```

### Segmented Control Style

```html
<!-- Selection group (radio-like behavior) -->
<div role="group" class="button-group" aria-label="View options">
  <button type="button" class="btn-outline" aria-pressed="false">Day</button>
  <button type="button" class="btn-outline" aria-pressed="true">Week</button>
  <button type="button" class="btn-outline" aria-pressed="false">Month</button>
</div>

<!-- Toggle group -->
<div role="group" class="button-group" aria-label="Text formatting">
  <button type="button" class="btn-outline" aria-pressed="true" aria-label="Bold">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
      <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
    </svg>
  </button>
  <button type="button" class="btn-outline" aria-pressed="false" aria-label="Italic">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="19" x2="10" y1="4" y2="4" />
      <line x1="14" x2="5" y1="20" y2="20" />
      <line x1="15" x2="9" y1="4" y2="20" />
    </svg>
  </button>
  <button type="button" class="btn-outline" aria-pressed="false" aria-label="Underline">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4v6a6 6 0 0 0 12 0V4" />
      <line x1="4" x2="20" y1="20" y2="20" />
    </svg>
  </button>
</div>
```

### Icon Button Groups

```html
<!-- Icon-only button group -->
<div role="group" class="button-group" aria-label="Text alignment">
  <button type="button" class="btn-icon-outline" aria-label="Align left">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="21" x2="3" y1="6" y2="6" />
      <line x1="15" x2="3" y1="12" y2="12" />
      <line x1="17" x2="3" y1="18" y2="18" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Align center">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="18" x2="6" y1="6" y2="6" />
      <line x1="21" x2="3" y1="12" y2="12" />
      <line x1="18" x2="6" y1="18" y2="18" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Align right">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="21" x2="3" y1="6" y2="6" />
      <line x1="21" x2="9" y1="12" y2="12" />
      <line x1="21" x2="7" y1="18" y2="18" />
    </svg>
  </button>
</div>

<!-- Media controls -->
<div role="group" class="button-group" aria-label="Media controls">
  <button type="button" class="btn-icon-outline" aria-label="Previous">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="19,20 9,12 19,4" />
      <line x1="5" x2="5" y1="19" y2="5" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Play">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="5,3 19,12 5,21" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Next">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="5,4 15,12 5,20" />
      <line x1="19" x2="19" y1="5" y2="19" />
    </svg>
  </button>
</div>
```

### Small Button Groups

```html
<!-- Small button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-sm-outline">Small</button>
  <button type="button" class="btn-sm-outline">Buttons</button>
  <button type="button" class="btn-sm-outline">Group</button>
</div>

<!-- Small icon group -->
<div role="group" class="button-group">
  <button type="button" class="btn-sm-icon-outline" aria-label="Cut">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="6" cy="6" r="3" />
      <circle cx="6" cy="18" r="3" />
      <line x1="20" x2="8.12" y1="4" y2="15.88" />
      <line x1="14.47" x2="20" y1="14.48" y2="20" />
      <line x1="8.12" x2="14.47" y1="8.12" y2="14.47" />
    </svg>
  </button>
  <button type="button" class="btn-sm-icon-outline" aria-label="Copy">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
      <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
    </svg>
  </button>
  <button type="button" class="btn-sm-icon-outline" aria-label="Paste">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect width="8" height="4" x="8" y="2" rx="1" ry="1" />
      <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2" />
    </svg>
  </button>
</div>
```

### Button Group with Dropdown

```html
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Snooze</button>
  
  <!-- Dropdown trigger -->
  <div class="dropdown-menu">
    <button type="button" class="btn-icon-outline" aria-haspopup="menu" aria-expanded="false">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
    </button>
    <!-- Dropdown menu content -->
    <div data-popover aria-hidden="true">
      <div role="menu">
        <div role="menuitem">Archive</div>
        <div role="menuitem">Delete</div>
      </div>
    </div>
  </div>
</div>

<!-- Split button pattern -->
<div role="group" class="button-group">
  <button type="button" class="btn">Save</button>
  <div class="dropdown-menu">
    <button type="button" class="btn-icon" aria-haspopup="menu" aria-label="Save options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m6 9 6 6 6-6" />
      </svg>
    </button>
  </div>
</div>
```

### Pagination Button Group

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- Previous button -->
  <button type="button" class="btn-icon-outline" aria-label="Previous page">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m15 18-6-6 6-6" />
    </svg>
  </button>
  
  <!-- Page numbers -->
  <div role="group" class="button-group" aria-label="Page navigation">
    <button type="button" class="btn-outline">1</button>
    <button type="button" class="btn" aria-current="page">2</button>
    <button type="button" class="btn-outline">3</button>
    <button type="button" class="btn-outline">4</button>
    <button type="button" class="btn-outline">5</button>
  </div>
  
  <!-- Next button -->
  <button type="button" class="btn-icon-outline" aria-label="Next page">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m9 18 6-6-6-6" />
    </svg>
  </button>
</div>
```

### Form Button Groups

```html
<!-- Form actions -->
<div class="flex w-fit items-stretch gap-2">
  <button type="button" class="btn-outline">Cancel</button>
  
  <div role="group" class="button-group">
    <button type="submit" class="btn">Save Draft</button>
    <button type="submit" class="btn">Publish</button>
  </div>
</div>

<!-- Modal actions -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline" onclick="closeModal()">Cancel</button>
  <button type="button" class="btn-destructive" onclick="confirmDelete()">Delete</button>
</div>

<!-- Wizard navigation -->
<div class="flex justify-between">
  <button type="button" class="btn-outline">Back</button>
  
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Save Draft</button>
    <button type="submit" class="btn">Continue</button>
  </div>
</div>
```

## Accessibility Features

- **Role Attributes**: Use `role="group"` for semantic grouping
- **Keyboard Navigation**: Standard Tab/Shift+Tab navigation between groups
- **Arrow Key Navigation**: Can implement custom arrow key handling within groups
- **Screen Reader Support**: Groups are announced as units
- **Labels**: Use `aria-label` to describe group purpose

### Enhanced Accessibility

```html
<!-- Labeled button group -->
<div role="group" class="button-group" aria-label="Text formatting options">
  <button type="button" class="btn-outline" aria-pressed="false">Bold</button>
  <button type="button" class="btn-outline" aria-pressed="true">Italic</button>
  <button type="button" class="btn-outline" aria-pressed="false">Underline</button>
</div>

<!-- Group with description -->
<fieldset class="border-none p-0 m-0">
  <legend class="sr-only">Alignment options</legend>
  <div role="group" class="button-group">
    <button type="button" class="btn-outline" aria-pressed="true">Left</button>
    <button type="button" class="btn-outline" aria-pressed="false">Center</button>
    <button type="button" class="btn-outline" aria-pressed="false">Right</button>
  </div>
</fieldset>

<!-- Radio-like group -->
<div role="radiogroup" class="button-group" aria-label="View mode">
  <button type="button" role="radio" class="btn" aria-checked="false">Grid</button>
  <button type="button" role="radio" class="btn-outline" aria-checked="true">List</button>
  <button type="button" role="radio" class="btn-outline" aria-checked="false">Card</button>
</div>
```

## JavaScript Integration

### Toggle Selection

```javascript
// Single selection (radio-like)
function selectButton(button, group) {
  // Remove selection from all buttons in group
  group.querySelectorAll('button').forEach(btn => {
    btn.classList.remove('btn');
    btn.classList.add('btn-outline');
    btn.setAttribute('aria-pressed', 'false');
  });
  
  // Select clicked button
  button.classList.remove('btn-outline');
  button.classList.add('btn');
  button.setAttribute('aria-pressed', 'true');
}

// Multi-selection (checkbox-like)
function toggleButton(button) {
  const isPressed = button.getAttribute('aria-pressed') === 'true';
  button.setAttribute('aria-pressed', (!isPressed).toString());
  
  if (isPressed) {
    button.classList.remove('btn');
    button.classList.add('btn-outline');
  } else {
    button.classList.remove('btn-outline');
    button.classList.add('btn');
  }
}

// Initialize button group
document.querySelectorAll('.button-group').forEach(group => {
  group.addEventListener('click', (e) => {
    if (e.target.tagName === 'BUTTON') {
      const isRadio = group.hasAttribute('data-radio');
      
      if (isRadio) {
        selectButton(e.target, group);
      } else {
        toggleButton(e.target);
      }
    }
  });
});
```

### Keyboard Navigation

```javascript
// Arrow key navigation within groups
function initButtonGroupKeyboard() {
  document.querySelectorAll('.button-group').forEach(group => {
    const buttons = Array.from(group.querySelectorAll('button'));
    
    buttons.forEach((button, index) => {
      button.addEventListener('keydown', (e) => {
        let targetIndex;
        
        switch (e.key) {
          case 'ArrowLeft':
            e.preventDefault();
            targetIndex = index > 0 ? index - 1 : buttons.length - 1;
            buttons[targetIndex].focus();
            break;
            
          case 'ArrowRight':
            e.preventDefault();
            targetIndex = index < buttons.length - 1 ? index + 1 : 0;
            buttons[targetIndex].focus();
            break;
            
          case 'Home':
            e.preventDefault();
            buttons[0].focus();
            break;
            
          case 'End':
            e.preventDefault();
            buttons[buttons.length - 1].focus();
            break;
        }
      });
    });
  });
}

// Initialize on page load
document.addEventListener('DOMContentLoaded', initButtonGroupKeyboard);
```

### React Integration

```jsx
import React, { useState } from 'react';

function ButtonGroup({ children, variant = 'single', onChange, className = '' }) {
  const [selected, setSelected] = useState(variant === 'single' ? 0 : []);
  
  const handleClick = (index) => {
    if (variant === 'single') {
      setSelected(index);
      onChange?.(index);
    } else {
      const newSelected = selected.includes(index)
        ? selected.filter(i => i !== index)
        : [...selected, index];
      setSelected(newSelected);
      onChange?.(newSelected);
    }
  };
  
  return (
    <div role="group" className={`button-group ${className}`}>
      {React.Children.map(children, (child, index) => 
        React.cloneElement(child, {
          'aria-pressed': variant === 'single' 
            ? selected === index 
            : selected.includes(index),
          onClick: () => handleClick(index),
          className: variant === 'single'
            ? (selected === index ? 'btn' : 'btn-outline')
            : (selected.includes(index) ? 'btn' : 'btn-outline')
        })
      )}
    </div>
  );
}

// Usage
function App() {
  return (
    <ButtonGroup variant="single" onChange={(index) => console.log('Selected:', index)}>
      <button type="button">Option 1</button>
      <button type="button">Option 2</button>
      <button type="button">Option 3</button>
    </ButtonGroup>
  );
}
```

### Vue Integration

```vue
<template>
  <div role="group" class="button-group">
    <button
      v-for="(option, index) in options"
      :key="option.id"
      type="button"
      :class="getButtonClass(index)"
      :aria-pressed="isSelected(index)"
      @click="handleClick(index)"
    >
      {{ option.label }}
    </button>
  </div>
</template>

<script>
export default {
  props: {
    options: Array,
    modelValue: [Number, Array],
    multiple: Boolean
  },
  emits: ['update:modelValue'],
  methods: {
    isSelected(index) {
      return this.multiple 
        ? this.modelValue.includes(index)
        : this.modelValue === index;
    },
    getButtonClass(index) {
      return this.isSelected(index) ? 'btn' : 'btn-outline';
    },
    handleClick(index) {
      if (this.multiple) {
        const selected = [...this.modelValue];
        const selectedIndex = selected.indexOf(index);
        
        if (selectedIndex > -1) {
          selected.splice(selectedIndex, 1);
        } else {
          selected.push(index);
        }
        
        this.$emit('update:modelValue', selected);
      } else {
        this.$emit('update:modelValue', index);
      }
    }
  }
};
</script>
```

## Best Practices

1. **Logical Grouping**: Group related actions together
2. **Consistent Variants**: Use same button variant within a group
3. **Accessibility**: Include proper ARIA attributes
4. **Visual Separation**: Use gaps between different button groups
5. **Icon Guidelines**: Use consistent icon sizes within groups
6. **Responsive Design**: Consider button group behavior on mobile
7. **Clear Labels**: Provide descriptive labels for screen readers
8. **State Management**: Handle selection states appropriately

## Common Patterns

### Toolbar Groups

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- File operations -->
  <div role="group" class="button-group" aria-label="File operations">
    <button type="button" class="btn-sm-icon-outline" aria-label="New file">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
        <polyline points="14,2 14,8 20,8" />
      </svg>
    </button>
    <button type="button" class="btn-sm-icon-outline" aria-label="Save">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" />
        <polyline points="17,21 17,13 7,13 7,21" />
        <polyline points="7,3 7,8 15,8" />
      </svg>
    </button>
  </div>
  
  <!-- Edit operations -->
  <div role="group" class="button-group" aria-label="Edit operations">
    <button type="button" class="btn-sm-icon-outline" aria-label="Undo">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M3 7v6h6" />
        <path d="M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13" />
      </svg>
    </button>
    <button type="button" class="btn-sm-icon-outline" aria-label="Redo">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M21 7v6h-6" />
        <path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7" />
      </svg>
    </button>
  </div>
</div>
```

### Status Filter Groups

```html
<div class="space-y-2">
  <label class="text-sm font-medium">Filter by status</label>
  <div role="group" class="button-group" aria-label="Status filter">
    <button type="button" class="btn" aria-pressed="true">All</button>
    <button type="button" class="btn-outline" aria-pressed="false">Active</button>
    <button type="button" class="btn-outline" aria-pressed="false">Pending</button>
    <button type="button" class="btn-outline" aria-pressed="false">Inactive</button>
  </div>
</div>
```

### Modal Actions

```html
<div class="flex justify-end gap-2 pt-6 border-t">
  <button type="button" class="btn-outline">Cancel</button>
  
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Save Draft</button>
    <button type="button" class="btn">Publish Now</button>
  </div>
</div>
```

## Related Components

- [Button](./button.md) - Individual button styling
- [Dropdown Menu](./dropdown-menu.md) - For button group dropdowns
- [Toggle](./toggle.md) - For toggle-style button groups
- [Toolbar](./toolbar.md) - For complex tool arrangements