# Tabs Component

A set of layered sections of content—known as tab panels—that are displayed one at a time.

## Basic Usage

```html
<div class="tabs w-full" id="demo-tabs">
  <nav role="tablist" aria-orientation="horizontal" class="w-full">
    <button type="button" role="tab" id="demo-tabs-tab-1" aria-controls="demo-tabs-panel-1" aria-selected="true" tabindex="0">Account</button>
    <button type="button" role="tab" id="demo-tabs-tab-2" aria-controls="demo-tabs-panel-2" aria-selected="false" tabindex="0">Password</button>
  </nav>

  <div role="tabpanel" id="demo-tabs-panel-1" aria-labelledby="demo-tabs-tab-1" tabindex="-1" aria-selected="true">
    <div class="card">
      <header>
        <h2>Account</h2>
        <p>Make changes to your account here. Click save when you're done.</p>
      </header>
      <section>
        <form class="form grid gap-6">
          <div class="grid gap-3">
            <label for="account-name">Name</label>
            <input type="text" id="account-name" value="Pedro Duarte" />
          </div>
          <div class="grid gap-3">
            <label for="account-username">Username</label>
            <input type="text" id="account-username" value="@peduarte" />
          </div>
        </form>
      </section>
      <footer>
        <button type="button" class="btn">Save changes</button>
      </footer>
    </div>
  </div>

  <div role="tabpanel" id="demo-tabs-panel-2" aria-labelledby="demo-tabs-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="card">
      <header>
        <h2>Password</h2>
        <p>Change your password here. After saving, you'll be logged out.</p>
      </header>
      <section>
        <form class="form grid gap-6">
          <div class="grid gap-3">
            <label for="password-current">Current password</label>
            <input type="password" id="password-current" />
          </div>
          <div class="grid gap-3">
            <label for="password-new">New password</label>
            <input type="password" id="password-new" />
          </div>
        </form>
      </section>
      <footer>
        <button type="button" class="btn">Save Password</button>
      </footer>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/tabs.min.js" defer></script>
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
- **`tabs`** - Applied to the main container
- **`w-full`** - Full width container (optional but common)

### Supporting Classes
- Card classes for content panels
- Form classes for form content
- Button classes for actions

### Tailwind Utilities Used
- `w-full` - Full width
- `grid gap-*` - Grid layout with spacing
- Various spacing and layout utilities

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "tabs" | Yes |
| `id` | string | Unique identifier | Recommended |

### Tablist (Nav) Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "tablist" | Yes |
| `aria-orientation` | string | "horizontal" or "vertical" | Recommended |

### Tab Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `role` | string | Must be "tab" | Yes |
| `id` | string | Unique identifier | Yes |
| `aria-controls` | string | References panel ID | Yes |
| `aria-selected` | boolean | Current selection state | Yes |
| `tabindex` | number | 0 for focusable, -1 for not | Yes |

### Tab Panel Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "tabpanel" | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `aria-labelledby` | string | References tab button ID | Yes |
| `tabindex` | number | Usually -1 | Yes |
| `aria-selected` | boolean | Selection state | Yes |
| `hidden` | boolean | Hide inactive panels | Conditional |

## HTML Structure

```html
<div class="tabs" id="tabs-id">
  <!-- Tab navigation -->
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" 
            id="tab-1" 
            aria-controls="panel-1" 
            aria-selected="true" 
            tabindex="0">
      Tab 1
    </button>
    <button type="button" role="tab" 
            id="tab-2" 
            aria-controls="panel-2" 
            aria-selected="false" 
            tabindex="0">
      Tab 2
    </button>
  </nav>

  <!-- Tab panels -->
  <div role="tabpanel" 
       id="panel-1" 
       aria-labelledby="tab-1" 
       tabindex="-1" 
       aria-selected="true">
    <!-- Panel 1 content -->
  </div>

  <div role="tabpanel" 
       id="panel-2" 
       aria-labelledby="tab-2" 
       tabindex="-1" 
       aria-selected="false" 
       hidden>
    <!-- Panel 2 content -->
  </div>
</div>
```

## Examples

### Basic Text Tabs
```html
<div class="tabs" id="text-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="text-tab-1" aria-controls="text-panel-1" aria-selected="true" tabindex="0">
      Overview
    </button>
    <button type="button" role="tab" id="text-tab-2" aria-controls="text-panel-2" aria-selected="false" tabindex="0">
      Features
    </button>
    <button type="button" role="tab" id="text-tab-3" aria-controls="text-panel-3" aria-selected="false" tabindex="0">
      Reviews
    </button>
  </nav>

  <div role="tabpanel" id="text-panel-1" aria-labelledby="text-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <h3 class="text-lg font-semibold">Product Overview</h3>
      <p class="mt-2 text-muted-foreground">
        This is the overview content for our product. It provides a general introduction
        and highlights the main benefits.
      </p>
    </div>
  </div>

  <div role="tabpanel" id="text-panel-2" aria-labelledby="text-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="text-lg font-semibold">Key Features</h3>
      <ul class="mt-2 space-y-2 text-muted-foreground">
        <li>• Feature 1: Advanced functionality</li>
        <li>• Feature 2: User-friendly interface</li>
        <li>• Feature 3: Cross-platform support</li>
      </ul>
    </div>
  </div>

  <div role="tabpanel" id="text-panel-3" aria-labelledby="text-tab-3" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="text-lg font-semibold">Customer Reviews</h3>
      <p class="mt-2 text-muted-foreground">
        See what our customers are saying about this product.
      </p>
    </div>
  </div>
</div>
```

### Tabs with Icons
```html
<div class="tabs" id="icon-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="icon-tab-1" aria-controls="icon-panel-1" aria-selected="true" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
        <polyline points="9,22 9,12 15,12 15,22" />
      </svg>
      Home
    </button>
    <button type="button" role="tab" id="icon-tab-2" aria-controls="icon-panel-2" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
        <circle cx="12" cy="12" r="3" />
      </svg>
      Settings
    </button>
  </nav>

  <div role="tabpanel" id="icon-panel-1" aria-labelledby="icon-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <p>Home panel content</p>
    </div>
  </div>

  <div role="tabpanel" id="icon-panel-2" aria-labelledby="icon-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <p>Settings panel content</p>
    </div>
  </div>
</div>
```

### Vertical Tabs
```html
<div class="tabs flex gap-6" id="vertical-tabs">
  <nav role="tablist" aria-orientation="vertical" class="flex flex-col w-48">
    <button type="button" role="tab" id="v-tab-1" aria-controls="v-panel-1" aria-selected="true" tabindex="0" class="text-left">
      General
    </button>
    <button type="button" role="tab" id="v-tab-2" aria-controls="v-panel-2" aria-selected="false" tabindex="0" class="text-left">
      Security
    </button>
    <button type="button" role="tab" id="v-tab-3" aria-controls="v-panel-3" aria-selected="false" tabindex="0" class="text-left">
      Notifications
    </button>
    <button type="button" role="tab" id="v-tab-4" aria-controls="v-panel-4" aria-selected="false" tabindex="0" class="text-left">
      Advanced
    </button>
  </nav>

  <div class="flex-1">
    <div role="tabpanel" id="v-panel-1" aria-labelledby="v-tab-1" tabindex="-1" aria-selected="true">
      <h3 class="text-lg font-semibold mb-4">General Settings</h3>
      <p>Configure general application settings here.</p>
    </div>

    <div role="tabpanel" id="v-panel-2" aria-labelledby="v-tab-2" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Security Settings</h3>
      <p>Manage your security preferences and authentication options.</p>
    </div>

    <div role="tabpanel" id="v-panel-3" aria-labelledby="v-tab-3" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Notification Preferences</h3>
      <p>Choose how and when you want to receive notifications.</p>
    </div>

    <div role="tabpanel" id="v-panel-4" aria-labelledby="v-tab-4" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Advanced Options</h3>
      <p>Advanced settings for power users.</p>
    </div>
  </div>
</div>
```

### Tabs with Badges
```html
<div class="tabs" id="badge-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="badge-tab-1" aria-controls="badge-panel-1" aria-selected="true" tabindex="0">
      Messages
      <span class="badge ml-1.5">12</span>
    </button>
    <button type="button" role="tab" id="badge-tab-2" aria-controls="badge-panel-2" aria-selected="false" tabindex="0">
      Notifications
      <span class="badge ml-1.5">3</span>
    </button>
    <button type="button" role="tab" id="badge-tab-3" aria-controls="badge-panel-3" aria-selected="false" tabindex="0">
      Updates
    </button>
  </nav>

  <div role="tabpanel" id="badge-panel-1" aria-labelledby="badge-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <h3 class="font-semibold">Unread Messages (12)</h3>
      <!-- Message list -->
    </div>
  </div>

  <div role="tabpanel" id="badge-panel-2" aria-labelledby="badge-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="font-semibold">Recent Notifications (3)</h3>
      <!-- Notification list -->
    </div>
  </div>

  <div role="tabpanel" id="badge-panel-3" aria-labelledby="badge-tab-3" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="font-semibold">System Updates</h3>
      <!-- Update list -->
    </div>
  </div>
</div>
```

### Disabled Tabs
```html
<div class="tabs" id="disabled-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="d-tab-1" aria-controls="d-panel-1" aria-selected="true" tabindex="0">
      Available
    </button>
    <button type="button" role="tab" id="d-tab-2" aria-controls="d-panel-2" aria-selected="false" tabindex="0">
      Premium
    </button>
    <button type="button" role="tab" id="d-tab-3" aria-controls="d-panel-3" aria-selected="false" tabindex="-1" aria-disabled="true" class="opacity-50 cursor-not-allowed">
      Coming Soon
    </button>
  </nav>

  <div role="tabpanel" id="d-panel-1" aria-labelledby="d-tab-1" tabindex="-1" aria-selected="true">
    <p class="py-4">This feature is available to all users.</p>
  </div>

  <div role="tabpanel" id="d-panel-2" aria-labelledby="d-tab-2" tabindex="-1" aria-selected="false" hidden>
    <p class="py-4">This feature is available to premium users.</p>
  </div>

  <div role="tabpanel" id="d-panel-3" aria-labelledby="d-tab-3" tabindex="-1" aria-selected="false" hidden>
    <p class="py-4">This feature is coming soon.</p>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const tabs = document.getElementById('my-tabs');

// Listen for initialization
tabs.addEventListener('basecoat:initialized', (e) => {
  console.log('Tabs initialized');
});

// Listen for tab changes
tabs.addEventListener('basecoat:tabs:change', (e) => {
  console.log('Tab changed to:', e.detail.tabId);
});
```

### Manual Tab Control
```javascript
// Switch to specific tab
function switchToTab(tabId) {
  const tab = document.getElementById(tabId);
  if (tab) {
    tab.click();
  }
}

// Get active tab
function getActiveTab() {
  return document.querySelector('[role="tab"][aria-selected="true"]');
}

// Disable/enable tab
function setTabEnabled(tabId, enabled) {
  const tab = document.getElementById(tabId);
  if (tab) {
    tab.setAttribute('aria-disabled', !enabled);
    tab.setAttribute('tabindex', enabled ? '0' : '-1');
    tab.classList.toggle('opacity-50', !enabled);
    tab.classList.toggle('cursor-not-allowed', !enabled);
  }
}
```

## Keyboard Navigation

The tabs component supports full keyboard navigation:

- **Tab**: Move focus between tabs
- **Arrow Keys**: Navigate between tabs (when focused on tablist)
- **Home**: Go to first tab
- **End**: Go to last tab
- **Space/Enter**: Activate focused tab

## Accessibility Features

- **ARIA Roles**: Proper tablist, tab, and tabpanel roles
- **Screen Reader Support**: Clear announcements of tab states
- **Keyboard Navigation**: Full keyboard support
- **Focus Management**: Proper focus handling
- **State Communication**: Selected and disabled states

### Enhanced Accessibility
```html
<div class="tabs" id="accessible-tabs">
  <h2 id="tabs-heading" class="text-lg font-semibold mb-4">User Preferences</h2>
  <nav role="tablist" aria-orientation="horizontal" aria-labelledby="tabs-heading">
    <button type="button" 
            role="tab" 
            id="a-tab-1" 
            aria-controls="a-panel-1" 
            aria-selected="true" 
            tabindex="0"
            aria-describedby="a-tab-1-desc">
      Profile
    </button>
    <span id="a-tab-1-desc" class="sr-only">Edit your profile information</span>
    
    <button type="button" 
            role="tab" 
            id="a-tab-2" 
            aria-controls="a-panel-2" 
            aria-selected="false" 
            tabindex="0"
            aria-describedby="a-tab-2-desc">
      Privacy
    </button>
    <span id="a-tab-2-desc" class="sr-only">Manage your privacy settings</span>
  </nav>

  <div role="tabpanel" 
       id="a-panel-1" 
       aria-labelledby="a-tab-1" 
       tabindex="-1" 
       aria-selected="true">
    <!-- Profile content -->
  </div>

  <div role="tabpanel" 
       id="a-panel-2" 
       aria-labelledby="a-tab-2" 
       tabindex="-1" 
       aria-selected="false" 
       hidden>
    <!-- Privacy content -->
  </div>
</div>
```

## Styling Customization

### Custom Tab Styles
```css
/* Pill-style tabs */
.tabs-pills [role="tab"] {
  border-radius: 9999px;
  padding: 0.5rem 1rem;
}

/* Underline tabs */
.tabs-underline [role="tab"] {
  border-bottom: 2px solid transparent;
}

.tabs-underline [role="tab"][aria-selected="true"] {
  border-bottom-color: var(--primary);
}

/* Boxed tabs */
.tabs-boxed [role="tab"] {
  border: 1px solid var(--border);
  margin-right: -1px;
}

.tabs-boxed [role="tab"]:first-child {
  border-radius: 0.5rem 0 0 0.5rem;
}

.tabs-boxed [role="tab"]:last-child {
  border-radius: 0 0.5rem 0.5rem 0;
}
```

### Responsive Tabs
```html
<div class="tabs" id="responsive-tabs">
  <nav role="tablist" aria-orientation="horizontal" class="flex overflow-x-auto scrollbar-none">
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-1" aria-controls="r-panel-1" aria-selected="true" tabindex="0">
      Dashboard
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-2" aria-controls="r-panel-2" aria-selected="false" tabindex="0">
      Analytics
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-3" aria-controls="r-panel-3" aria-selected="false" tabindex="0">
      Reports
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-4" aria-controls="r-panel-4" aria-selected="false" tabindex="0">
      Settings
    </button>
  </nav>
  <!-- Tab panels -->
</div>
```

## Best Practices

1. **Clear Labels**: Use descriptive tab labels
2. **Logical Order**: Arrange tabs in a logical sequence
3. **Visual Feedback**: Clearly indicate active tab
4. **Keyboard Support**: Ensure full keyboard navigation
5. **Touch Targets**: Adequate size for mobile
6. **Content Loading**: Consider lazy loading for heavy content
7. **State Persistence**: Consider saving active tab state
8. **Responsive Design**: Handle overflow gracefully on mobile

## Common Patterns

### Settings Tabs
```html
<div class="tabs" id="settings-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="s-tab-1" aria-controls="s-panel-1" aria-selected="true" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
        <circle cx="12" cy="7" r="4" />
      </svg>
      Profile
    </button>
    <button type="button" role="tab" id="s-tab-2" aria-controls="s-panel-2" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <path d="M17 11h1a3 3 0 0 1 0 6h-1" />
        <path d="M9 12v6" />
        <path d="M13 12v6" />
        <path d="M14 7.5c-1 0-1.44.5-3 .5s-2-.5-3-.5" />
        <path d="M2 5c0-1.1.9-2 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5z" />
      </svg>
      Billing
    </button>
    <button type="button" role="tab" id="s-tab-3" aria-controls="s-panel-3" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
      Advanced
    </button>
  </nav>
  <!-- Tab panels with forms -->
</div>
```

### Product Details Tabs
```html
<div class="tabs" id="product-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="p-tab-1" aria-controls="p-panel-1" aria-selected="true" tabindex="0">
      Description
    </button>
    <button type="button" role="tab" id="p-tab-2" aria-controls="p-panel-2" aria-selected="false" tabindex="0">
      Specifications
    </button>
    <button type="button" role="tab" id="p-tab-3" aria-controls="p-panel-3" aria-selected="false" tabindex="0">
      Reviews (24)
    </button>
    <button type="button" role="tab" id="p-tab-4" aria-controls="p-panel-4" aria-selected="false" tabindex="0">
      Q&A (8)
    </button>
  </nav>
  <!-- Product information panels -->
</div>
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function Tabs({ items, defaultTab = 0 }) {
  const [activeTab, setActiveTab] = useState(defaultTab);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="tabs">
      <nav role="tablist" aria-orientation="horizontal">
        {items.map((item, index) => (
          <button
            key={index}
            type="button"
            role="tab"
            id={`tab-${index}`}
            aria-controls={`panel-${index}`}
            aria-selected={activeTab === index}
            tabindex={activeTab === index ? 0 : -1}
            onClick={() => setActiveTab(index)}
          >
            {item.label}
          </button>
        ))}
      </nav>

      {items.map((item, index) => (
        <div
          key={index}
          role="tabpanel"
          id={`panel-${index}`}
          aria-labelledby={`tab-${index}`}
          tabindex="-1"
          aria-selected={activeTab === index}
          hidden={activeTab !== index}
        >
          {item.content}
        </div>
      ))}
    </div>
  );
}
```

### Vue Integration
```vue
<template>
  <div class="tabs">
    <nav role="tablist" aria-orientation="horizontal">
      <button
        v-for="(item, index) in items"
        :key="index"
        type="button"
        role="tab"
        :id="`tab-${index}`"
        :aria-controls="`panel-${index}`"
        :aria-selected="activeTab === index"
        :tabindex="activeTab === index ? 0 : -1"
        @click="activeTab = index"
      >
        {{ item.label }}
      </button>
    </nav>

    <div
      v-for="(item, index) in items"
      :key="index"
      role="tabpanel"
      :id="`panel-${index}`"
      :aria-labelledby="`tab-${index}`"
      tabindex="-1"
      :aria-selected="activeTab === index"
      :hidden="activeTab !== index"
    >
      <slot :name="`panel-${index}`">
        {{ item.content }}
      </slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    items: Array,
    defaultTab: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      activeTab: this.defaultTab
    };
  },
  mounted() {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }
};
</script>
```

### HTMX Integration
```html
<div class="tabs" id="htmx-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" 
            role="tab" 
            id="htmx-tab-1" 
            aria-controls="htmx-panel-1" 
            aria-selected="true" 
            tabindex="0"
            hx-get="/api/content/overview"
            hx-target="#htmx-panel-1"
            hx-trigger="click once">
      Overview
    </button>
    <button type="button" 
            role="tab" 
            id="htmx-tab-2" 
            aria-controls="htmx-panel-2" 
            aria-selected="false" 
            tabindex="0"
            hx-get="/api/content/details"
            hx-target="#htmx-panel-2"
            hx-trigger="click once">
      Details
    </button>
  </nav>

  <div role="tabpanel" id="htmx-panel-1" aria-labelledby="htmx-tab-1" tabindex="-1" aria-selected="true">
    <div class="skeleton">Loading overview...</div>
  </div>

  <div role="tabpanel" id="htmx-panel-2" aria-labelledby="htmx-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="skeleton">Loading details...</div>
  </div>
</div>
```

## Related Components

- [Accordion](./accordion.md) - For expandable content sections
- [Card](./card.md) - For content containers within tabs
- [Navigation](./navigation.md) - For page-level navigation
- [Button Group](./button-group.md) - For alternative switching patterns