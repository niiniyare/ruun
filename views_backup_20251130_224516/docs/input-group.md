# Input Group Pattern

Display additional information or actions to an input or textarea.

**Note:** There is no dedicated Input Group component in Basecoat. This pattern is achieved using positioning utilities and relative containers.

## Basic Usage

```html
<div class="relative">
  <input type="text" class="input pl-9 pr-20" placeholder="Search...">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">12 results</div>
</div>
```

## CSS Classes

### Container Classes
- **`relative`** - Required for positioning child elements

### Input Classes
- **`input`** - Base input styling
- **`textarea`** - Base textarea styling
- **Padding adjustments**: `pl-9`, `pr-20`, etc. to make room for positioned elements

### Positioned Element Classes
- **`absolute`** - Absolute positioning
- **`left-3`**, `right-3` - Horizontal positioning
- **`top-1/2 -translate-y-1/2`** - Vertical centering
- **`pointer-events-none`** - Disable mouse interaction (for decorative elements)

### Icon and Text Styling
- **`text-muted-foreground`** - Muted text color
- **`[&>svg]:size-4`** - Icon sizing
- **`text-sm`** - Small text size

## Component Attributes

### No specific component attributes - uses standard HTML input/textarea attributes
| Element | Attributes | Description |
|---------|------------|-------------|
| Container | `class="relative"` | Enables absolute positioning |
| Input | Standard input attributes | `type`, `placeholder`, `class`, etc. |
| Positioned elements | `class` with positioning utilities | Absolute positioning and styling |

### No JavaScript Required (Basic)
Basic input groups work with pure CSS positioning.

## HTML Structure

```html
<!-- Basic pattern -->
<div class="relative">
  <input type="text" class="input [padding-adjustments]" placeholder="...">
  <div class="absolute [positioning] [styling]">
    <!-- Icon, text, or interactive element -->
  </div>
</div>
```

## Examples

### Search Input with Icon and Results

```html
<div class="relative">
  <input type="text" class="input pl-9 pr-20" placeholder="Search...">
  <!-- Left icon -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <!-- Right text -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">12 results</div>
</div>
```

### URL Input with Protocol Prefix

```html
<div class="relative">
  <input type="text" class="input pl-15 pr-9" placeholder="example.com">
  <!-- Left prefix text -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">https://</div>
  <!-- Right help icon -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Enter your domain name">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="M12 16v-4" />
      <path d="M12 8h.01" />
    </svg>
  </div>
</div>
```

### Username Input with Validation

```html
<div class="relative">
  <input type="text" class="input pr-9" placeholder="@shadcn">
  <!-- Validation indicator -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-primary text-primary-foreground flex size-4 items-center justify-center rounded-full [&>svg]:size-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 6 9 17l-5-5" />
    </svg>
  </div>
</div>
```

### Enhanced Textarea with Controls

```html
<div class="relative">
  <textarea class="textarea pr-10 min-h-27 pb-12" placeholder="Ask, Search or Chat..."></textarea>
  <!-- Bottom control bar -->
  <footer role="group" class="absolute bottom-0 px-3 pb-3 pt-1.5 flex items-center w-full gap-2">
    <!-- Add button -->
    <button type="button" class="btn-icon-outline rounded-full size-6 text-muted-foreground hover:text-accent-foreground">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M5 12h14" />
        <path d="M12 5v14" />
      </svg>
    </button>
    
    <!-- Dropdown menu -->
    <div class="dropdown-menu">
      <button type="button" class="btn-sm-ghost text-muted-foreground hover:text-accent-foreground h-6 p-2">
        Auto
      </button>
      <div data-popover aria-hidden="true" data-side="top" class="min-w-32">
        <div role="menu">
          <div role="menuitem">Auto</div>
          <div role="menuitem">Agent</div>
          <div role="menuitem">Manual</div>
        </div>
      </div>
    </div>
    
    <!-- Usage indicator -->
    <div class="text-muted-foreground text-sm ml-auto">52% used</div>
    
    <!-- Separator -->
    <hr class="w-0 h-4 border-r border-border shrink-0 m-0">
    
    <!-- Submit button -->
    <button type="button" class="btn-icon rounded-full size-6 bg-muted-foreground hover:bg-foreground" disabled>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m5 12 7-7 7 7" />
        <path d="M12 19V5" />
      </svg>
    </button>
  </footer>
</div>
```

### Email Input with Domain

```html
<div class="relative">
  <input type="email" class="input pr-24" placeholder="username">
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">@company.com</div>
</div>
```

### Currency Input

```html
<div class="relative">
  <input type="number" class="input pl-8 pr-12" placeholder="0.00">
  <!-- Currency symbol -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">$</div>
  <!-- Unit -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">USD</div>
</div>
```

### Password Input with Strength

```html
<div class="relative">
  <input type="password" class="input pr-20" placeholder="Password">
  <!-- Strength indicator -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-1">
    <div class="text-xs text-muted-foreground">Strong</div>
    <div class="flex gap-1">
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
    </div>
  </div>
</div>
```

### Phone Number Input

```html
<div class="relative">
  <input type="tel" class="input pl-16 pr-9" placeholder="123-456-7890">
  <!-- Country code -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">+1</div>
  <!-- Format help -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Format: XXX-XXX-XXXX">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="M12 16v-4" />
      <path d="M12 8h.01" />
    </svg>
  </div>
</div>
```

### File Size Input

```html
<div class="relative">
  <input type="number" class="input pr-12" placeholder="100" min="1">
  <!-- Unit selector -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2">
    <select class="text-sm text-muted-foreground bg-transparent border-none focus:outline-none">
      <option>KB</option>
      <option>MB</option>
      <option>GB</option>
    </select>
  </div>
</div>
```

### Search with Clear Button

```html
<div class="relative">
  <input type="search" class="input pl-9 pr-9" placeholder="Search..." id="search-input">
  <!-- Search icon -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <!-- Clear button (interactive) -->
  <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground [&>svg]:size-4" onclick="document.getElementById('search-input').value = ''">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="m15 9-6 6" />
      <path d="m9 9 6 6" />
    </svg>
  </button>
</div>
```

### Loading State Input

```html
<div class="relative">
  <input type="text" class="input pr-9" placeholder="Processing..." disabled>
  <!-- Loading spinner -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground">
    <svg class="animate-spin size-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="m12 2 0 4c-4.418 0-8 3.582-8 8 0 1.1.224 2.148.63 3.1L2.369 18.9C1.502 17.065 1 15.087 1 13c0-6.075 4.925-11 11-11Z"></path>
    </svg>
  </div>
</div>
```

### Multi-element Input Group

```html
<div class="flex">
  <!-- Input with left addon -->
  <div class="relative flex-1">
    <input type="text" class="input rounded-r-none pr-9" placeholder="repository-name">
    <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Repository name">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10" />
        <path d="M12 16v-4" />
        <path d="M12 8h.01" />
      </svg>
    </div>
  </div>
  <!-- Button addon -->
  <button type="button" class="btn-outline rounded-l-none border-l-0">Clone</button>
</div>
```

## Accessibility Features

- **Proper Labeling**: Use labels or `aria-label` for inputs
- **Interactive Elements**: Ensure clickable elements are accessible
- **Focus Management**: Positioned elements shouldn't interfere with focus
- **Screen Reader Support**: Use `aria-describedby` for helpful text

### Enhanced Accessibility

```html
<!-- Input with accessible help text -->
<div class="field">
  <label for="username-input">Username</label>
  <div class="relative">
    <input 
      type="text" 
      id="username-input"
      class="input pl-8"
      placeholder="username"
      aria-describedby="username-help"
    >
    <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">@</div>
  </div>
  <p id="username-help" class="text-sm text-muted-foreground mt-1">Choose a unique username for your account</p>
</div>

<!-- Search with live region -->
<div class="relative">
  <input 
    type="search" 
    class="input pl-9 pr-20" 
    placeholder="Search..."
    aria-describedby="search-results"
  >
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg><!-- search icon --></svg>
  </div>
  <div 
    id="search-results"
    class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm"
    aria-live="polite"
  >
    12 results
  </div>
</div>
```

## JavaScript Integration

### Dynamic Content Updates

```javascript
// Update result count
function updateSearchResults(count) {
  const resultsElement = document.querySelector('#search-results');
  if (resultsElement) {
    resultsElement.textContent = `${count} results`;
  }
}

// Toggle validation state
function updateValidationState(input, isValid) {
  const container = input.parentElement;
  const indicator = container.querySelector('.validation-indicator');
  
  if (indicator) {
    if (isValid) {
      indicator.className = 'absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-green-500 text-white flex size-4 items-center justify-center rounded-full [&>svg]:size-3';
      indicator.innerHTML = '<svg><!-- checkmark --></svg>';
    } else {
      indicator.className = 'absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-red-500 text-white flex size-4 items-center justify-center rounded-full [&>svg]:size-3';
      indicator.innerHTML = '<svg><!-- x mark --></svg>';
    }
  }
}

// Clear input
function clearInput(inputId) {
  const input = document.getElementById(inputId);
  if (input) {
    input.value = '';
    input.focus();
  }
}
```

### React Integration

```jsx
import React, { useState } from 'react';

function InputGroup({ 
  children, 
  leftElement, 
  rightElement, 
  className = '',
  ...props 
}) {
  return (
    <div className={`relative ${className}`}>
      {React.cloneElement(children, {
        className: `input ${leftElement ? 'pl-9' : ''} ${rightElement ? 'pr-9' : ''} ${children.props.className || ''}`,
        ...props
      })}
      {leftElement && (
        <div className="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
          {leftElement}
        </div>
      )}
      {rightElement && (
        <div className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4">
          {rightElement}
        </div>
      )}
    </div>
  );
}

// Usage examples
function App() {
  const [searchValue, setSearchValue] = useState('');
  const [results, setResults] = useState(0);
  
  return (
    <div className="space-y-4">
      <InputGroup
        leftElement={<SearchIcon />}
        rightElement={<span className="text-sm pointer-events-none">{results} results</span>}
      >
        <input 
          type="search" 
          placeholder="Search..."
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
        />
      </InputGroup>
      
      <InputGroup
        leftElement={<span className="text-sm">https://</span>}
        rightElement={
          <button onClick={() => console.log('Help clicked')}>
            <InfoIcon />
          </button>
        }
      >
        <input type="url" placeholder="example.com" />
      </InputGroup>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div class="relative">
    <input 
      v-bind="$attrs"
      :class="inputClasses"
      @input="$emit('input', $event.target.value)"
    />
    
    <div v-if="leftElement" class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
      <slot name="left" />
    </div>
    
    <div v-if="rightElement" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4">
      <slot name="right" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    leftElement: Boolean,
    rightElement: Boolean
  },
  computed: {
    inputClasses() {
      let classes = 'input';
      if (this.leftElement) classes += ' pl-9';
      if (this.rightElement) classes += ' pr-9';
      return classes;
    }
  }
};
</script>
```

## Best Practices

1. **Padding Adjustments**: Always adjust input padding to accommodate positioned elements
2. **Pointer Events**: Use `pointer-events-none` for decorative elements
3. **Interactive Elements**: Make buttons and clickable icons accessible
4. **Consistent Sizing**: Use consistent icon sizes and positioning
5. **Visual Hierarchy**: Use muted colors for secondary elements
6. **Mobile Friendly**: Ensure touch targets are large enough
7. **Loading States**: Show loading indicators for async operations

## Common Patterns

### Form Field Enhancement

```html
<div class="field">
  <label for="enhanced-input">Enhanced Input</label>
  <div class="relative">
    <input type="text" id="enhanced-input" class="input pl-9 pr-9" placeholder="Enter value">
    <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
      <svg><!-- icon --></svg>
    </div>
    <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground">
      <svg><!-- action icon --></svg>
    </button>
  </div>
  <p class="text-sm text-muted-foreground mt-1">Helper text for this field</p>
</div>
```

### Search Interface

```html
<div class="relative">
  <input type="search" class="input pl-9 pr-24" placeholder="Search anything...">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg><!-- search icon --></svg>
  </div>
  <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-2">
    <span class="text-xs text-muted-foreground">⌘K</span>
  </div>
</div>
```

### Status Input

```html
<div class="relative">
  <input type="text" class="input pr-9" readonly value="Connected">
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none">
    <div class="size-2 rounded-full bg-green-500"></div>
  </div>
</div>
```

## Related Components

- [Input](./input.md) - Base input component
- [Textarea](./textarea.md) - Base textarea component
- [Button](./button.md) - For interactive elements
- [Field](./field.md) - For complete form field structure