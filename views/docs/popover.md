# Popover Component

Displays rich content in a portal, triggered by a button.

## Basic Usage

```html
<div id="demo-popover" class="popover">
  <button id="demo-popover-trigger" type="button" aria-expanded="false" aria-controls="demo-popover-popover" class="btn-outline">Open popover</button>
  <div id="demo-popover-popover" data-popover aria-hidden="true" class="w-80">
    <div class="grid gap-4">
      <header class="grid gap-1.5">
        <h4 class="leading-none font-medium">Dimensions</h4>
        <p class="text-muted-foreground text-sm">Set the dimensions for the layer.</p>
      </header>
      <form class="form grid gap-2">
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-width">Width</label>
          <input type="text" id="demo-popover-width" value="100%" class="col-span-2 h-8" autofocus />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-max-width">Max. width</label>
          <input type="text" id="demo-popover-max-width" value="300px" class="col-span-2 h-8" />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-height">Height</label>
          <input type="text" id="demo-popover-height" value="25px" class="col-span-2 h-8" />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-max-height">Max. height</label>
          <input type="text" id="demo-popover-max-height" value="none" class="col-span-2 h-8" />
        </div>
      </form>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/popover.min.js" defer></script>
```

### Initialize Component
```javascript
// Automatic initialization when page loads
document.addEventListener('DOMContentLoaded', () => {
  if (window.basecoat) window.basecoat.initAll();
});
```

## CSS Classes

### Primary Classes
- **`popover`** - Applied to the main container
- **`data-popover`** - Applied to the popover content element

### Supporting Classes
- Button classes for triggers (`btn-outline`, `btn`, etc.)
- Width/height utility classes (`w-80`, `h-8`, etc.)
- Grid and spacing classes for content layout

### Tailwind Utilities Used
- `grid gap-4` - Grid layout with gap spacing
- `grid-cols-3` - Three-column grid layout
- `items-center` - Vertical alignment
- `col-span-2` - Span two grid columns
- `w-80` - Fixed width for popover
- `text-muted-foreground` - Muted text color
- `text-sm` - Small text size
- `leading-none` - No line height
- `font-medium` - Medium font weight

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "popover" | Yes |
| `id` | string | Unique identifier | Yes |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `id` | string | Referenced by aria-labelledby | Yes |
| `aria-expanded` | boolean | Tracks open/closed state | Yes |
| `aria-controls` | string | References popover content ID | Yes |
| `popovertarget` | string | Alternative native HTML attribute | Optional |

### Popover Content Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-popover` | boolean | Marks as popover content | Yes |
| `aria-hidden` | boolean | Controls visibility for screen readers | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `data-side` | string | Position: "top", "right", "bottom", "left" | No |
| `data-align` | string | Alignment: "start", "center", "end" | No |

## HTML Structure

```html
<div class="popover" id="popover-id">
  <!-- Trigger button -->
  <button type="button" 
          id="trigger-id" 
          aria-expanded="false" 
          aria-controls="popover-content-id" 
          class="btn-outline">
    Trigger Text
  </button>
  
  <!-- Popover content -->
  <div id="popover-content-id" 
       data-popover 
       aria-hidden="true" 
       class="w-80">
    <!-- Popover content -->
  </div>
</div>
```

## Examples

### Basic Information Popover
```html
<div id="info-popover" class="popover">
  <button id="info-popover-trigger" type="button" aria-expanded="false" aria-controls="info-popover-content" class="btn-outline">
    Show Info
  </button>
  <div id="info-popover-content" data-popover aria-hidden="true" class="w-64">
    <div class="grid gap-3">
      <h4 class="font-medium">Additional Information</h4>
      <p class="text-sm text-muted-foreground">
        This is some helpful information that appears in a popover.
      </p>
    </div>
  </div>
</div>
```

### Form Popover
```html
<div id="form-popover" class="popover">
  <button id="form-popover-trigger" type="button" aria-expanded="false" aria-controls="form-popover-content" class="btn-outline">
    Edit Settings
  </button>
  <div id="form-popover-content" data-popover aria-hidden="true" class="w-80">
    <div class="grid gap-4">
      <header class="grid gap-1.5">
        <h4 class="leading-none font-medium">Settings</h4>
        <p class="text-muted-foreground text-sm">Configure your preferences.</p>
      </header>
      <form class="form grid gap-3">
        <div class="grid gap-2">
          <label for="setting-name">Name</label>
          <input type="text" id="setting-name" placeholder="Enter name">
        </div>
        <div class="grid gap-2">
          <label for="setting-email">Email</label>
          <input type="email" id="setting-email" placeholder="Enter email">
        </div>
        <div class="flex gap-2">
          <button type="submit" class="btn">Save</button>
          <button type="button" class="btn-outline">Cancel</button>
        </div>
      </form>
    </div>
  </div>
</div>
```

### Positioned Popover
```html
<div id="positioned-popover" class="popover">
  <button id="positioned-popover-trigger" type="button" aria-expanded="false" aria-controls="positioned-popover-content" class="btn-outline">
    Top Popover
  </button>
  <div id="positioned-popover-content" data-popover aria-hidden="true" class="w-56" data-side="top" data-align="end">
    <div class="grid gap-3">
      <h4 class="font-medium">Positioned Above</h4>
      <p class="text-sm text-muted-foreground">
        This popover appears above the trigger button, aligned to the end.
      </p>
    </div>
  </div>
</div>
```

### Rich Content Popover
```html
<div id="rich-popover" class="popover">
  <button id="rich-popover-trigger" type="button" aria-expanded="false" aria-controls="rich-popover-content" class="btn-outline">
    Show Details
  </button>
  <div id="rich-popover-content" data-popover aria-hidden="true" class="w-72">
    <div class="grid gap-4">
      <div class="flex items-center gap-3">
        <div class="size-12 rounded-full bg-primary/10 flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
        </div>
        <div class="grid gap-1">
          <h4 class="font-medium">John Doe</h4>
          <p class="text-sm text-muted-foreground">Software Developer</p>
        </div>
      </div>
      <div class="grid gap-2">
        <div class="flex items-center gap-2 text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="20" height="16" x="2" y="4" rx="2" />
            <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
          </svg>
          john.doe@example.com
        </div>
        <div class="flex items-center gap-2 text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z" />
          </svg>
          +1 (555) 123-4567
        </div>
      </div>
      <div class="flex gap-2">
        <button type="button" class="btn">Contact</button>
        <button type="button" class="btn-outline">View Profile</button>
      </div>
    </div>
  </div>
</div>
```

### List Popover
```html
<div id="list-popover" class="popover">
  <button id="list-popover-trigger" type="button" aria-expanded="false" aria-controls="list-popover-content" class="btn-outline">
    Options
  </button>
  <div id="list-popover-content" data-popover aria-hidden="true" class="w-48">
    <div class="grid gap-1">
      <h4 class="font-medium px-3 py-2 border-b">Quick Actions</h4>
      <div class="grid gap-1 p-1">
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M4 13.5V4a1 1 0 0 1 1-1h4.5" />
            <path d="M14 4h5a1 1 0 0 1 1 1v5" />
            <path d="M4 20v-5.5a1 1 0 0 1 1-1H10" />
            <path d="M20 10v10a1 1 0 0 1-1 1h-5" />
          </svg>
          Expand
        </button>
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
            <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
          </svg>
          Copy
        </button>
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 6h18" />
            <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
            <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
          </svg>
          Delete
        </button>
      </div>
    </div>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const popover = document.getElementById('my-popover');

// Listen for component initialization
popover.addEventListener('basecoat:initialized', (e) => {
  console.log('Popover initialized');
});

// Listen for popover open (dispatched on document)
document.addEventListener('basecoat:popover', (e) => {
  console.log('A popover was opened, others will close');
});
```

### Manual Control
```javascript
// Get popover elements
const trigger = document.getElementById('my-popover-trigger');
const content = document.getElementById('my-popover-content');

// Open popover
function openPopover() {
  trigger.setAttribute('aria-expanded', 'true');
  content.setAttribute('aria-hidden', 'false');
  content.focus();
}

// Close popover
function closePopover() {
  trigger.setAttribute('aria-expanded', 'false');
  content.setAttribute('aria-hidden', 'true');
  trigger.focus();
}

// Toggle popover
function togglePopover() {
  const isOpen = trigger.getAttribute('aria-expanded') === 'true';
  if (isOpen) {
    closePopover();
  } else {
    openPopover();
  }
}
```

### Click Outside to Close
```javascript
document.addEventListener('click', (e) => {
  const popover = e.target.closest('.popover');
  if (!popover) {
    // Click outside all popovers - close any open ones
    document.querySelectorAll('.popover').forEach(p => {
      const trigger = p.querySelector('[aria-expanded]');
      const content = p.querySelector('[data-popover]');
      if (trigger && content) {
        trigger.setAttribute('aria-expanded', 'false');
        content.setAttribute('aria-hidden', 'true');
      }
    });
  }
});
```

## Positioning

### Popover Positioning
Use `data-side` and `data-align` attributes on the popover content element:

```html
<!-- Position above trigger, aligned to start -->
<div data-popover data-side="top" data-align="start">
  <!-- Popover content -->
</div>

<!-- Position to the right, center aligned -->
<div data-popover data-side="right" data-align="center">
  <!-- Popover content -->
</div>

<!-- Position below (default), aligned to end -->
<div data-popover data-side="bottom" data-align="end">
  <!-- Popover content -->
</div>
```

### Side Options
- `top` - Above the trigger
- `right` - To the right of the trigger  
- `bottom` - Below the trigger (default)
- `left` - To the left of the trigger

### Alignment Options
- `start` - Align to the start edge
- `center` - Center align
- `end` - Align to the end edge

## Accessibility Features

- **Keyboard Navigation**: ESC key closes, focus management
- **Screen Reader Support**: Proper ARIA attributes and state management
- **Focus Management**: Automatic focus handling when opening/closing
- **State Communication**: Expanded/collapsed states properly announced

### Enhanced Accessibility
```html
<div id="accessible-popover" class="popover">
  <button id="accessible-popover-trigger" 
          type="button" 
          aria-expanded="false" 
          aria-controls="accessible-popover-content"
          aria-describedby="accessible-popover-desc"
          class="btn-outline">
    Settings
  </button>
  <div id="accessible-popover-content" 
       data-popover 
       aria-hidden="true" 
       aria-labelledby="accessible-popover-trigger"
       role="dialog"
       class="w-80">
    <div class="grid gap-4">
      <h4 id="accessible-popover-title">Configuration Settings</h4>
      <p id="accessible-popover-desc" class="text-sm text-muted-foreground">
        Adjust your application preferences below.
      </p>
      <!-- Content -->
    </div>
  </div>
</div>
```

## Best Practices

1. **Clear Triggers**: Use descriptive button text that indicates what the popover contains
2. **Appropriate Sizing**: Size popovers appropriately for their content
3. **Keyboard Support**: Ensure ESC key closes and tab navigation works
4. **Focus Management**: Set appropriate focus when opening
5. **Click Outside**: Allow clicking outside to close for non-critical content
6. **Content Length**: Keep content concise or use scrolling for long content
7. **Positioning**: Choose appropriate positioning based on context
8. **Form Integration**: Use proper form handling for interactive content

## Common Patterns

### Tooltip-Style Popover
```javascript
// Simple text content popover
function createTooltipPopover(trigger, content) {
  const popover = document.createElement('div');
  popover.className = 'popover';
  popover.innerHTML = `
    <div data-popover aria-hidden="true" class="max-w-xs">
      <p class="text-sm">${content}</p>
    </div>
  `;
  
  trigger.parentNode.insertBefore(popover, trigger.nextSibling);
  
  trigger.addEventListener('mouseenter', () => {
    popover.querySelector('[data-popover]').setAttribute('aria-hidden', 'false');
  });
  
  trigger.addEventListener('mouseleave', () => {
    popover.querySelector('[data-popover]').setAttribute('aria-hidden', 'true');
  });
}
```

### Settings Panel Pattern
```javascript
// Reusable settings popover
function createSettingsPopover(triggerSelector, settings) {
  const trigger = document.querySelector(triggerSelector);
  const settingsHTML = settings.map(setting => `
    <div class="grid gap-2">
      <label for="${setting.id}">${setting.label}</label>
      <input type="${setting.type}" id="${setting.id}" value="${setting.value || ''}">
    </div>
  `).join('');
  
  // Create and attach popover
  // Handle form submission
}
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function Popover({ trigger, children, side = 'bottom', align = 'start' }) {
  const [isOpen, setIsOpen] = useState(false);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="popover">
      <button 
        type="button"
        aria-expanded={isOpen}
        className="btn-outline"
        onClick={() => setIsOpen(!isOpen)}
      >
        {trigger}
      </button>
      <div 
        data-popover 
        aria-hidden={!isOpen}
        data-side={side}
        data-align={align}
        className="w-80"
      >
        {children}
      </div>
    </div>
  );
}
```

### HTMX Integration
```html
<div class="popover">
  <button 
    type="button" 
    aria-expanded="false" 
    aria-controls="dynamic-popover-content"
    hx-get="/api/popover-content"
    hx-target="#dynamic-popover-content"
    hx-trigger="click"
    class="btn-outline"
  >
    Load Content
  </button>
  <div id="dynamic-popover-content" data-popover aria-hidden="true" class="w-80">
    <!-- Content loaded dynamically -->
  </div>
</div>
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% call popover(
  id="demo-popover",
  trigger="Open popover",
  trigger_attrs={"class": "btn-outline"},
  popover_attrs={"class": "w-80"}
) %}
  <div class="grid gap-4">
    <header class="grid gap-1.5">
      <h4 class="leading-none font-medium">Dimensions</h4>
      <p class="text-muted-foreground text-sm">Set the dimensions for the layer.</p>
    </header>
    <form class="form grid gap-2">
      <div class="grid grid-cols-3 items-center gap-4">
        <label for="demo-popover-width">Width</label>
        <input type="text" id="demo-popover-width" value="100%" class="col-span-2 h-8"/>
      </div>
    </form>
  </div>
{% endcall %}
```

### Advanced Configuration
```jinja2
{% set popover_content %}
  <div class="grid gap-3">
    <h4 class="font-medium">{{ title }}</h4>
    <p class="text-sm text-muted-foreground">{{ description }}</p>
    {% for item in items %}
      <div class="flex items-center gap-2">
        {{ item.label }}: {{ item.value }}
      </div>
    {% endfor %}
  </div>
{% endset %}

{{ popover(
  id="config-popover",
  trigger="Configuration",
  content=popover_content,
  trigger_attrs={"class": "btn-ghost"},
  popover_attrs={"class": "w-64", "data-side": "right"}
) }}
```

## Related Components

- [Button](./button.md) - For popover triggers
- [Dialog](./dialog.md) - For modal interactions
- [Dropdown Menu](./dropdown-menu.md) - For menu-based interactions
- [Tooltip](./tooltip.md) - For simple text overlays