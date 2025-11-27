# Dropdown Menu Component

A menu that appears when a trigger element is activated, containing a list of choices that users can select from.

## Basic Usage

```html
<div id="demo-dropdown-menu" class="dropdown-menu">
  <button type="button" id="demo-dropdown-menu-trigger" aria-haspopup="menu" aria-controls="demo-dropdown-menu-menu" aria-expanded="false" class="btn-outline">
    Open
  </button>
  <div id="demo-dropdown-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="demo-dropdown-menu-menu" aria-labelledby="demo-dropdown-menu-trigger">
      <div role="menuitem">Profile</div>
      <div role="menuitem">Settings</div>
      <hr role="separator" />
      <div role="menuitem">Logout</div>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/dropdown-menu.min.js" defer></script>
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
- **`dropdown-menu`** - Applied to the wrapper container
- **`min-w-56`** - Minimum width for the popover (commonly used)

### Supporting Classes
- Button classes for triggers (`btn-outline`, `btn`, etc.)
- Popover positioning classes (`data-side`, `data-align`)
- Text styling classes for shortcuts and descriptions

### Tailwind Utilities Used
- `ml-auto` - Push keyboard shortcuts to the right
- `text-xs tracking-widest` - Styling for keyboard shortcuts
- `text-muted-foreground` - Muted text color
- `invisible` - Hide unchecked checkboxes
- `group-aria-checked:visible` - Show when checked

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "dropdown-menu" | Yes |
| `id` | string | Unique identifier | Yes |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `id` | string | Referenced by aria-labelledby | Yes |
| `aria-haspopup` | string | Must be "menu" | Yes |
| `aria-controls` | string | References menu ID | Yes |
| `aria-expanded` | boolean | Tracks open/closed state | Yes |

### Popover Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-popover` | boolean | Marks as popover element | Yes |
| `aria-hidden` | boolean | Controls visibility for screen readers | Yes |
| `data-side` | string | Position: "top", "right", "bottom", "left" | No |
| `data-align` | string | Alignment: "start", "center", "end" | No |

### Menu Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "menu" | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `aria-labelledby` | string | References trigger button ID | Yes |

### Menu Item Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | "menuitem", "menuitemcheckbox", or "menuitemradio" | Yes |
| `aria-checked` | boolean | For checkbox/radio items | For checkboxes/radios |
| `aria-disabled` | boolean | Marks item as disabled | No |

## HTML Structure

```html
<div class="dropdown-menu" id="dropdown-id">
  <!-- Trigger button -->
  <button type="button" 
          id="trigger-id" 
          aria-haspopup="menu" 
          aria-controls="menu-id" 
          aria-expanded="false" 
          class="btn-outline">
    Trigger Text
  </button>
  
  <!-- Popover container -->
  <div id="popover-id" 
       data-popover 
       aria-hidden="true" 
       class="min-w-56">
    
    <!-- Menu container -->
    <div role="menu" 
         id="menu-id" 
         aria-labelledby="trigger-id">
      
      <!-- Menu items -->
      <div role="menuitem">Menu Item</div>
      
      <!-- Separator -->
      <hr role="separator" />
      
      <!-- Grouped items -->
      <div role="group" aria-labelledby="group-heading">
        <div role="heading" id="group-heading">Group Title</div>
        <div role="menuitem">Grouped Item</div>
      </div>
      
    </div>
  </div>
</div>
```

## Examples

### Basic Menu
```html
<div id="basic-menu" class="dropdown-menu">
  <button type="button" id="basic-menu-trigger" aria-haspopup="menu" aria-controls="basic-menu-menu" aria-expanded="false" class="btn-outline">
    Actions
  </button>
  <div id="basic-menu-popover" data-popover aria-hidden="true" class="min-w-40">
    <div role="menu" id="basic-menu-menu" aria-labelledby="basic-menu-trigger">
      <div role="menuitem">Edit</div>
      <div role="menuitem">Copy</div>
      <div role="menuitem">Move</div>
      <hr role="separator" />
      <div role="menuitem" aria-disabled="true">Delete</div>
    </div>
  </div>
</div>
```

### Menu with Icons and Shortcuts
```html
<div id="icon-menu" class="dropdown-menu">
  <button type="button" id="icon-menu-trigger" aria-haspopup="menu" aria-controls="icon-menu-menu" aria-expanded="false" class="btn-outline">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="1" />
      <circle cx="19" cy="12" r="1" />
      <circle cx="5" cy="12" r="1" />
    </svg>
  </button>
  <div id="icon-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="icon-menu-menu" aria-labelledby="icon-menu-trigger">
      <div role="group" aria-labelledby="account-options">
        <div role="heading" id="account-options">My Account</div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
          Profile
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘P</span>
        </div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="20" height="14" x="2" y="5" rx="2" />
            <line x1="2" x2="22" y1="10" y2="10" />
          </svg>
          Billing
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘B</span>
        </div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
            <circle cx="12" cy="12" r="3" />
          </svg>
          Settings
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘S</span>
        </div>
      </div>
      <hr role="separator" />
      <div role="menuitem">GitHub</div>
      <div role="menuitem">Support</div>
      <hr role="separator" />
      <div role="menuitem">
        Logout
        <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘Q</span>
      </div>
    </div>
  </div>
</div>
```

### Menu with Checkboxes
```html
<div id="checkbox-menu" class="dropdown-menu">
  <button type="button" id="checkbox-menu-trigger" aria-haspopup="menu" aria-controls="checkbox-menu-menu" aria-expanded="false" class="btn-outline">
    View Options
  </button>
  <div id="checkbox-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="checkbox-menu-menu" aria-labelledby="checkbox-menu-trigger">
      <div role="group" aria-labelledby="appearance-options">
        <div role="heading" id="appearance-options">Appearance</div>
        <div role="menuitemcheckbox" aria-checked="true" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Status Bar
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘S</span>
        </div>
        <div role="menuitemcheckbox" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Activity Bar
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘B</span>
        </div>
        <div role="menuitemcheckbox" aria-checked="false" class="group" aria-disabled="true">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Panel (Disabled)
        </div>
      </div>
    </div>
  </div>
</div>
```

### Menu with Radio Buttons
```html
<div id="radio-menu" class="dropdown-menu">
  <button type="button" id="radio-menu-trigger" aria-haspopup="menu" aria-controls="radio-menu-menu" aria-expanded="false" class="btn-outline">
    Theme: Light
  </button>
  <div id="radio-menu-popover" data-popover aria-hidden="true" class="min-w-48">
    <div role="menu" id="radio-menu-menu" aria-labelledby="radio-menu-trigger">
      <div role="group" aria-labelledby="theme-options">
        <div role="heading" id="theme-options">Theme</div>
        <div role="menuitemradio" aria-checked="true" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Light
        </div>
        <div role="menuitemradio" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Dark
        </div>
        <div role="menuitemradio" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Auto
        </div>
      </div>
    </div>
  </div>
</div>
```

### Positioned Menu
```html
<div id="positioned-menu" class="dropdown-menu">
  <button type="button" id="positioned-menu-trigger" aria-haspopup="menu" aria-controls="positioned-menu-menu" aria-expanded="false" class="btn-outline">
    Top Menu
  </button>
  <div id="positioned-menu-popover" data-popover aria-hidden="true" class="min-w-48" data-side="top" data-align="end">
    <div role="menu" id="positioned-menu-menu" aria-labelledby="positioned-menu-trigger">
      <div role="menuitem">Item 1</div>
      <div role="menuitem">Item 2</div>
      <div role="menuitem">Item 3</div>
    </div>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const dropdownMenu = document.getElementById('my-dropdown');

// Listen for component initialization
dropdownMenu.addEventListener('basecoat:initialized', (e) => {
  console.log('Dropdown menu initialized');
});

// Listen for popover open/close
document.addEventListener('basecoat:popover', (e) => {
  console.log('A popover was opened, others will close');
});
```

### Menu Item Click Handling
```javascript
// Handle menu item clicks
document.querySelectorAll('[role="menuitem"]').forEach(item => {
  item.addEventListener('click', (e) => {
    const text = e.target.textContent.trim();
    console.log('Menu item clicked:', text);
    
    // Close the dropdown menu
    const dropdown = e.target.closest('.dropdown-menu');
    const trigger = dropdown.querySelector('[aria-haspopup="menu"]');
    trigger.setAttribute('aria-expanded', 'false');
    
    // Hide popover
    const popover = dropdown.querySelector('[data-popover]');
    popover.setAttribute('aria-hidden', 'true');
  });
});
```

### Checkbox/Radio State Management
```javascript
// Handle checkbox menu items
document.querySelectorAll('[role="menuitemcheckbox"]').forEach(item => {
  item.addEventListener('click', (e) => {
    const currentState = item.getAttribute('aria-checked') === 'true';
    item.setAttribute('aria-checked', (!currentState).toString());
    
    console.log(`${item.textContent.trim()}: ${!currentState ? 'checked' : 'unchecked'}`);
  });
});

// Handle radio button menu items
document.querySelectorAll('[role="menuitemradio"]').forEach(item => {
  item.addEventListener('click', (e) => {
    // Uncheck all radio items in the same group
    const group = item.closest('[role="group"]');
    if (group) {
      group.querySelectorAll('[role="menuitemradio"]').forEach(radio => {
        radio.setAttribute('aria-checked', 'false');
      });
    }
    
    // Check the clicked item
    item.setAttribute('aria-checked', 'true');
    
    console.log(`Selected: ${item.textContent.trim()}`);
  });
});
```

## Accessibility Features

- **Keyboard Navigation**: Arrow keys, Enter, Escape
- **Focus Management**: Proper focus handling when opening/closing
- **Screen Reader Support**: Full ARIA menu roles and properties
- **State Communication**: Checkbox/radio states properly announced
- **Group Structure**: Menu item groups with headings

### Enhanced Accessibility
```html
<div id="accessible-menu" class="dropdown-menu">
  <button type="button" 
          id="accessible-menu-trigger" 
          aria-haspopup="menu" 
          aria-controls="accessible-menu-menu" 
          aria-expanded="false" 
          aria-label="Open user account menu"
          class="btn-outline">
    Account
  </button>
  <div id="accessible-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="accessible-menu-menu" aria-labelledby="accessible-menu-trigger">
      <div role="group" aria-labelledby="account-heading">
        <div role="heading" id="account-heading" aria-level="3">Account Settings</div>
        <div role="menuitem" tabindex="-1">
          Profile Settings
        </div>
        <div role="menuitem" tabindex="-1" aria-describedby="billing-desc">
          Billing
          <div id="billing-desc" class="sr-only">Manage your subscription and payments</div>
        </div>
      </div>
      <hr role="separator" aria-hidden="true" />
      <div role="menuitem" tabindex="-1">
        Sign Out
      </div>
    </div>
  </div>
</div>
```

## Positioning

### Popover Positioning
Use `data-side` and `data-align` attributes on the popover element:

```html
<!-- Position above trigger, aligned to end -->
<div data-popover data-side="top" data-align="end">
  <!-- Menu content -->
</div>

<!-- Position to the right, center aligned -->
<div data-popover data-side="right" data-align="center">
  <!-- Menu content -->
</div>

<!-- Position below (default), aligned to start -->
<div data-popover data-side="bottom" data-align="start">
  <!-- Menu content -->
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

## Best Practices

1. **Use Semantic Roles**: Always use proper ARIA menu roles
2. **Group Related Items**: Use groups with headings for organization
3. **Keyboard Shortcuts**: Include keyboard shortcuts where applicable
4. **Clear Icons**: Use consistent, meaningful icons
5. **Logical Order**: Organize items logically
6. **Disable Appropriately**: Use `aria-disabled` for unavailable actions
7. **State Communication**: Properly manage checkbox/radio states
8. **Focus Management**: Ensure proper keyboard navigation

## Common Patterns

### Context Menu Pattern
```javascript
// Right-click context menu
document.addEventListener('contextmenu', (e) => {
  e.preventDefault();
  
  const contextMenu = createContextMenu([
    { label: 'Copy', shortcut: '⌘C', action: () => copy() },
    { label: 'Paste', shortcut: '⌘V', action: () => paste() },
    { separator: true },
    { label: 'Delete', shortcut: 'Del', action: () => delete() }
  ]);
  
  showMenuAtPosition(contextMenu, e.clientX, e.clientY);
});
```

### User Account Menu Pattern
```html
<div id="user-menu" class="dropdown-menu">
  <button type="button" id="user-menu-trigger" aria-haspopup="menu" aria-controls="user-menu-menu" aria-expanded="false" class="btn-ghost">
    <img src="/avatar.jpg" alt="User Avatar" class="size-8 rounded-full">
    John Doe
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m6 9 6 6 6-6" />
    </svg>
  </button>
  <div id="user-menu-popover" data-popover aria-hidden="true" class="min-w-56" data-align="end">
    <div role="menu" id="user-menu-menu" aria-labelledby="user-menu-trigger">
      <div role="group">
        <div class="px-3 py-2 text-sm">
          <div class="font-medium">John Doe</div>
          <div class="text-muted-foreground">john@example.com</div>
        </div>
      </div>
      <hr role="separator" />
      <div role="menuitem">Dashboard</div>
      <div role="menuitem">Settings</div>
      <div role="menuitem">Billing</div>
      <hr role="separator" />
      <div role="menuitem">Sign Out</div>
    </div>
  </div>
</div>
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function DropdownMenu({ trigger, items, onSelect }) {
  const [isOpen, setIsOpen] = useState(false);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="dropdown-menu">
      <button 
        type="button"
        aria-haspopup="menu"
        aria-expanded={isOpen}
        className="btn-outline"
      >
        {trigger}
      </button>
      <div data-popover aria-hidden={!isOpen} className="min-w-56">
        <div role="menu">
          {items.map((item, index) => (
            item.separator ? (
              <hr key={index} role="separator" />
            ) : (
              <div 
                key={index}
                role="menuitem" 
                onClick={() => onSelect(item)}
              >
                {item.label}
                {item.shortcut && (
                  <span className="text-muted-foreground ml-auto text-xs tracking-widest">
                    {item.shortcut}
                  </span>
                )}
              </div>
            )
          ))}
        </div>
      </div>
    </div>
  );
}
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% call dropdown_menu(
  id="user-menu",
  trigger="Account",
  trigger_attrs={"class": "btn-outline"},
  popover_attrs={"class": "min-w-56", "data-align": "end"}
) %}
  <div role="group" aria-labelledby="account-options">
    <div role="heading" id="account-options">My Account</div>
    <div role="menuitem">
      Profile
      <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘P</span>
    </div>
    <div role="menuitem">
      Settings
      <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘S</span>
    </div>
  </div>
  <hr role="separator">
  <div role="menuitem">
    Logout
    <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘Q</span>
  </div>
{% endcall %}
```

### Advanced Configuration
```jinja2
{% set menu_items = [
  {"type": "group", "heading": "Account", "items": [
    {"label": "Profile", "shortcut": "⇧⌘P"},
    {"label": "Settings", "shortcut": "⌘S"}
  ]},
  {"type": "separator"},
  {"label": "Logout", "shortcut": "⇧⌘Q"}
] %}

{{ dropdown_menu(
  id="dynamic-menu",
  trigger="Options",
  items=menu_items,
  trigger_attrs={"class": "btn-ghost"},
  popover_attrs={"class": "min-w-48"}
) }}
```

## Related Components

- [Button](./button.md) - For menu triggers
- [Popover](./popover.md) - Underlying positioning system
- [Dialog](./dialog.md) - For modal interactions
- [Select](./select.md) - For form-based selection