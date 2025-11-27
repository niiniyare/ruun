# Sidebar Component

A composable, themeable and customizable sidebar component.

## Basic Usage

```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-content-1">
        <h3 id="group-label-content-1">Getting started</h3>
        <ul>
          <li>
            <a href="#">
              <svg><!-- Icon --></svg>
              <span>Playground</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>

<main>
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar'))">Toggle sidebar</button>
  <h1>Content</h1>
</main>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/sidebar.min.js" defer></script>
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
- **`sidebar`** - Applied to the main sidebar container (`<aside>`)
- **`scrollbar`** - Applied to scrollable sections

### Supporting Classes
- Standard navigation and list styling
- Button classes for interactive elements
- Responsive utilities for mobile behavior

## Component Attributes

### Sidebar Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "sidebar" class | Yes |
| `aria-hidden` | boolean | Controls default visibility state | Recommended |
| `data-side` | string | Position: "left" or "right" (default: "left") | No |
| `id` | string | Unique identifier for multiple sidebars | No |

### Navigation Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `aria-label` | string | Describes the navigation purpose | Recommended |

### Group Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "group" | Yes |
| `aria-labelledby` | string | References heading ID | Recommended |

### Link Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-keep-mobile-sidebar-open` | boolean | Prevents mobile auto-close | No |

## HTML Structure

```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <!-- Optional header -->
    <header>
      <!-- Header content -->
    </header>
    
    <!-- Main content section -->
    <section class="scrollbar">
      <!-- Navigation group -->
      <div role="group" aria-labelledby="group-label-1">
        <h3 id="group-label-1">Group Title</h3>
        
        <ul>
          <!-- Regular link -->
          <li>
            <a href="/page">
              <svg><!-- Icon --></svg>
              <span>Page Title</span>
            </a>
          </li>
          
          <!-- Collapsible submenu -->
          <li>
            <details id="submenu-1">
              <summary aria-controls="submenu-1-content">
                <svg><!-- Icon --></svg>
                Submenu Title
              </summary>
              <ul id="submenu-1-content">
                <li>
                  <a href="/sub-page">
                    <span>Sub Page</span>
                  </a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
    
    <!-- Optional footer -->
    <footer>
      <!-- Footer content -->
    </footer>
  </nav>
</aside>

<main>
  <!-- Page content -->
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar'))">
    Toggle sidebar
  </button>
</main>
```

## Examples

### Basic Sidebar
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-main">
        <h3 id="group-label-main">Main Navigation</h3>
        
        <ul>
          <li>
            <a href="/">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
                <polyline points="9,22 9,12 15,12 15,22" />
              </svg>
              <span>Home</span>
            </a>
          </li>
          
          <li>
            <a href="/dashboard">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect width="7" height="9" x="3" y="3" rx="1" />
                <rect width="7" height="5" x="14" y="3" rx="1" />
                <rect width="7" height="9" x="14" y="12" rx="1" />
                <rect width="7" height="5" x="3" y="16" rx="1" />
              </svg>
              <span>Dashboard</span>
            </a>
          </li>
          
          <li>
            <a href="/projects">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z" />
                <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z" />
              </svg>
              <span>Projects</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Sidebar with Collapsible Sections
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-content">
        <h3 id="group-label-content">Getting started</h3>

        <ul>
          <li>
            <a href="/playground">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m7 11 2-2-2-2" />
                <path d="M11 13h4" />
                <rect width="18" height="18" x="3" y="3" rx="2" ry="2" />
              </svg>
              <span>Playground</span>
            </a>
          </li>

          <li>
            <a href="/models">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 8V4H8" />
                <rect width="16" height="12" x="4" y="8" rx="2" />
                <path d="M2 14h2" />
                <path d="M20 14h2" />
                <path d="M15 13v2" />
                <path d="M9 13v2" />
              </svg>
              <span>Models</span>
            </a>
          </li>

          <li>
            <details id="submenu-settings">
              <summary aria-controls="submenu-settings-content">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
                Settings
              </summary>
              <ul id="submenu-settings-content">
                <li>
                  <a href="/settings/general">
                    <span>General</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/team">
                    <span>Team</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/billing">
                    <span>Billing</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/limits">
                    <span>Limits</span>
                  </a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Sidebar with Header and Footer
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <!-- Header -->
    <header>
      <a href="/" class="btn-ghost p-2 h-12 w-full justify-start">
        <div class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256" class="h-4 w-4">
            <rect width="256" height="256" fill="none"></rect>
            <line x1="208" y1="128" x2="128" y2="208" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line>
            <line x1="192" y1="40" x2="40" y2="192" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line>
          </svg>
        </div>
        <div class="grid flex-1 text-left text-sm leading-tight">
          <span class="truncate font-medium">Your App</span>
          <span class="truncate text-xs">v1.0.0</span>
        </div>
      </a>
    </header>
    
    <!-- Main content -->
    <section class="scrollbar">
      <!-- Navigation groups -->
    </section>
    
    <!-- Footer -->
    <footer>
      <div class="p-4">
        <button type="button" class="btn-ghost w-full justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
            <polyline points="16,17 21,12 16,7" />
            <line x1="21" y1="12" x2="9" y2="12" />
          </svg>
          <span>Sign out</span>
        </button>
      </div>
    </footer>
  </nav>
</aside>
```

### Right-Side Sidebar
```html
<aside class="sidebar" data-side="right" aria-hidden="true">
  <nav aria-label="Right sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-tools">
        <h3 id="group-label-tools">Tools</h3>
        
        <ul>
          <li>
            <a href="/tools/inspector">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m21 21-6-6m6 6v-4.8m0 4.8h-4.8" />
                <path d="M3 16.2V21m0 0h4.8M3 21l6-6" />
                <path d="M21 7.8V3m0 0h-4.8M21 3l-6 6" />
                <path d="M3 7.8V3m0 0h4.8M3 3l6 6" />
              </svg>
              <span>Inspector</span>
            </a>
          </li>
          
          <li>
            <a href="/tools/debugger">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect width="20" height="16" x="2" y="4" rx="2" />
                <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
              </svg>
              <span>Debugger</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Multiple Sidebars
```html
<!-- Main navigation sidebar -->
<aside id="main-navigation" class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Main navigation">
    <!-- Main navigation content -->
  </nav>
</aside>

<!-- Tools sidebar -->
<aside id="tools-sidebar" class="sidebar" data-side="right" aria-hidden="true">
  <nav aria-label="Tools navigation">
    <!-- Tools content -->
  </nav>
</aside>

<main>
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: 'main-navigation', action: 'toggle' } }))">
    Toggle Main Nav
  </button>
  
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: 'tools-sidebar', action: 'toggle' } }))">
    Toggle Tools
  </button>
</main>
```

## JavaScript Events

### Toggle Sidebar
```javascript
// Toggle default sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar'));

// Toggle specific sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { id: 'main-navigation' }
}));
```

### Open/Close Sidebar
```javascript
// Open sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { action: 'open' }
}));

// Close sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { action: 'close' }
}));

// Open specific sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { id: 'main-navigation', action: 'open' }
}));
```

### Listen for Initialization
```javascript
// Listen for component initialization
document.querySelectorAll('.sidebar').forEach(sidebar => {
  sidebar.addEventListener('basecoat:initialized', (e) => {
    console.log('Sidebar initialized:', e.target.id);
  });
});
```

### Custom Event Handlers
```javascript
// Create custom toggle button
function toggleSidebar(sidebarId = null, action = 'toggle') {
  const detail = { action };
  if (sidebarId) detail.id = sidebarId;
  
  document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail }));
}

// Usage
toggleSidebar('main-nav', 'open');
toggleSidebar(null, 'close'); // Close default sidebar
```

## Mobile Behavior

### Auto-Close on Mobile
By default, clicking links in the sidebar will close it on mobile devices:

```html
<li>
  <a href="/page">Page Link</a> <!-- Closes sidebar on mobile -->
</li>
```

### Prevent Auto-Close
Use `data-keep-mobile-sidebar-open` to prevent auto-closing:

```html
<li>
  <a href="/page" data-keep-mobile-sidebar-open>
    Page Link <!-- Keeps sidebar open on mobile -->
  </a>
</li>

<li>
  <button type="button" data-keep-mobile-sidebar-open>
    Button <!-- Keeps sidebar open on mobile -->
  </button>
</li>
```

## Accessibility Features

- **Semantic HTML**: Uses `<aside>` and `<nav>` landmarks
- **ARIA Labels**: Descriptive labels for navigation
- **Keyboard Navigation**: Full keyboard support
- **Screen Reader Support**: Proper grouping and labeling
- **Focus Management**: Logical tab order and focus indicators
- **State Announcement**: Hidden/visible states properly announced

### Enhanced Accessibility
```html
<aside class="sidebar" data-side="left" aria-hidden="false" aria-label="Main navigation">
  <nav aria-label="Primary site navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="main-nav-heading">
        <h3 id="main-nav-heading">Main Navigation</h3>
        
        <ul role="list">
          <li role="listitem">
            <a href="/" aria-current="page">
              <svg aria-hidden="true"><!-- Icon --></svg>
              <span>Home</span>
            </a>
          </li>
          
          <li role="listitem">
            <details id="settings-menu">
              <summary aria-controls="settings-submenu" aria-expanded="false">
                <svg aria-hidden="true"><!-- Icon --></svg>
                Settings
              </summary>
              <ul id="settings-submenu" role="list">
                <li role="listitem">
                  <a href="/settings/general">General</a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

## Styling Customization

### Custom Sidebar Positioning
```css
/* Left sidebar (default) */
.sidebar[data-side="left"] {
  left: 0;
}

/* Right sidebar */
.sidebar[data-side="right"] {
  right: 0;
}
```

### Custom Width
```css
.sidebar {
  width: 280px; /* Custom width */
}

@media (max-width: 768px) {
  .sidebar {
    width: 100%; /* Full width on mobile */
  }
}
```

### Custom Scrollbar
```css
.sidebar .scrollbar {
  scrollbar-width: thin;
  scrollbar-color: var(--muted) var(--background);
}

.sidebar .scrollbar::-webkit-scrollbar {
  width: 6px;
}

.sidebar .scrollbar::-webkit-scrollbar-track {
  background: var(--background);
}

.sidebar .scrollbar::-webkit-scrollbar-thumb {
  background: var(--muted);
  border-radius: 3px;
}
```

## Best Practices

1. **Clear Grouping**: Use semantic grouping with proper headings
2. **Consistent Icons**: Use consistent icon styles throughout
3. **Logical Order**: Organize navigation items logically
4. **Mobile-First**: Consider mobile experience and auto-close behavior
5. **Accessibility**: Provide proper ARIA labels and landmarks
6. **Keyboard Support**: Ensure all interactive elements are keyboard accessible
7. **Visual Hierarchy**: Use proper heading levels and visual distinction
8. **State Management**: Handle open/closed states appropriately

## Integration Examples

### React Integration
```jsx
import React, { useEffect } from 'react';

function Sidebar({ isOpen, onToggle }) {
  useEffect(() => {
    // Initialize Basecoat sidebar
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <aside className="sidebar" data-side="left" aria-hidden={!isOpen}>
      <nav aria-label="Sidebar navigation">
        <section className="scrollbar">
          {/* Navigation content */}
        </section>
      </nav>
    </aside>
  );
}
```

### Vue Integration
```vue
<template>
  <aside class="sidebar" data-side="left" :aria-hidden="!isOpen">
    <nav aria-label="Sidebar navigation">
      <section class="scrollbar">
        <!-- Navigation content -->
      </section>
    </nav>
  </aside>
</template>

<script>
export default {
  props: {
    isOpen: Boolean
  },
  mounted() {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }
}
</script>
```

### HTMX Integration
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar" hx-get="/api/navigation" hx-trigger="load">
      <!-- Navigation loaded dynamically -->
    </section>
  </nav>
</aside>
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% set menu = [
  { type: "group", label: "Getting started", items: [
    { label: "Playground", url: "#" },
    { label: "Models", url: "#" },
    { label: "Settings", type: "submenu", items: [
      { label: "General", url: "#" },
      { label: "Team", url: "#" },
      { label: "Billing", url: "#" },
      { label: "Limits", url: "#" }
    ] }
  ]}
] %}

{{ sidebar(
  label="Sidebar navigation",
  menu=menu
) }}
```

### Advanced Configuration
```jinja2
{{ sidebar(
  label="Main navigation",
  side="left",
  hidden=false,
  id="main-nav",
  menu=navigation_items,
  header=header_content,
  footer=footer_content
) }}
```

## Related Components

- [Button](./button.md) - For interactive elements within sidebar
- [Navigation](./navigation.md) - For primary site navigation
- [Breadcrumb](./breadcrumb.md) - For hierarchical navigation
- [Dropdown Menu](./dropdown-menu.md) - For nested menu items