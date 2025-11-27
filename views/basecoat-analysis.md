# Basecoat CSS Component Analysis

## Overview
Analysis of `./static/css/basecoat.css` following Basecoat UI patterns:
- **Base classes**: Core component classes (e.g., `.card`, `.btn`, `.field`)
- **Part classes**: Component sub-elements (e.g., `.card-header`, `.field-label`)
- **Variant classes**: Component variations (e.g., `.btn-outline`, `.badge-secondary`)

---

## Base Component Classes

| Base Class | Description | Line References |
|------------|-------------|----------------|
| `.alert` | Alert/notification component | 153-178 |
| `.badge` | Badge/label component | 183-202 |
| `.btn` | Button component | 207-380 |
| `.button-group` | Button grouping container | 385-423 |
| `.card` | Card container component | 428-448 |
| `.command` | Command palette component | 486-591 |
| `.command-dialog` | Command dialog modal | 487-543 |
| `.dialog` | Modal dialog component | 595-652 |
| `.dropdown-menu` | Dropdown menu component | 657-691 |
| `.field` | Form field component | 703-735 |
| `.fieldset` | Form fieldset component | 696-702 |
| `.input` | Input component (various types) | 739-769, 867-873, 878-922, 1116-1123 |
| `.kbd` | Keyboard key component | 774-776 |
| `.label` | Form label component | 781-789 |
| `.popover` | Popover/tooltip component | 793-861 |
| `.select` | Select dropdown component | 926-989 |
| `.sidebar` | Sidebar navigation component | 994-1111 |
| `.table` | Table component | 1128-1151 |
| `.tabs` | Tab navigation component | 1156-1173 |
| `.textarea` | Textarea input component | 1178-1181 |
| `.toaster` | Toast notification container | 1186-1233 |
| `.toast` | Individual toast notification | 1199-1232 |

---

## Part Classes (Component-Part Pattern)

| Part Class | Parent Component | Description | Line References |
|------------|------------------|-------------|----------------|
| `.toast-content` | `.toast` | Toast message content | 1202-1223 |

**Note**: Most components in this Basecoat implementation use semantic HTML elements (header, section, footer) as parts rather than explicit part classes.

---

## Variant Classes

### Alert Variants
| Variant Class | Description | Line References |
|---------------|-------------|----------------|
| `.alert-destructive` | Destructive/error alert variant | 154, 172-178 |

### Badge Variants
| Variant Class | Description | Line References |
|---------------|-------------|----------------|
| `.badge-primary` | Primary badge variant | 184, 191-192 |
| `.badge-secondary` | Secondary badge variant | 185, 194-195 |
| `.badge-destructive` | Destructive/error badge variant | 186, 197-198 |
| `.badge-outline` | Outline badge variant | 187, 200-202 |

### Button Variants
| Variant Class | Description | Line References |
|---------------|-------------|----------------|
| `.btn-primary` | Primary button variant | 208, 252, 306, 308, 310-321 |
| `.btn-secondary` | Secondary button variant | 209, 253, 322-333 |
| `.btn-outline` | Outline button variant | 210, 254, 334-345 |
| `.btn-ghost` | Ghost button variant | 211, 255, 346-356 |
| `.btn-link` | Link-style button variant | 212, 256, 357-368 |
| `.btn-destructive` | Destructive button variant | 213, 257, 369-380 |

### Button Size Variants
| Variant Class | Description | Line References |
|---------------|-------------|----------------|
| `.btn-sm` | Small button variant | 214, 269-277 |
| `.btn-sm-primary` | Small primary button | 215, 270, 307-309, 313-321 |
| `.btn-sm-secondary` | Small secondary button | 216, 271, 323-333 |
| `.btn-sm-outline` | Small outline button | 217, 272, 335-345 |
| `.btn-sm-ghost` | Small ghost button | 218, 273, 347-356 |
| `.btn-sm-link` | Small link button | 219, 274, 358-368 |
| `.btn-sm-destructive` | Small destructive button | 220, 275, 370-380 |
| `.btn-lg` | Large button variant | 221, 287-295 |
| `.btn-lg-primary` | Large primary button | 222, 288, 309-321 |
| `.btn-lg-secondary` | Large secondary button | 223, 289, 324-333 |
| `.btn-lg-outline` | Large outline button | 224, 290, 336-345 |
| `.btn-lg-ghost` | Large ghost button | 225, 291, 348-356 |
| `.btn-lg-link` | Large link button | 226, 292, 359-368 |
| `.btn-lg-destructive` | Large destructive button | 227, 293, 371-380 |

### Button Icon Variants
| Variant Class | Description | Line References |
|---------------|-------------|----------------|
| `.btn-icon` | Icon-only button | 228, 260-268, 311-321 |
| `.btn-icon-primary` | Icon-only primary button | 229, 261, 312-321 |
| `.btn-icon-secondary` | Icon-only secondary button | 230, 262, 325-333 |
| `.btn-icon-outline` | Icon-only outline button | 231, 263, 337-345 |
| `.btn-icon-ghost` | Icon-only ghost button | 232, 264, 349-356 |
| `.btn-icon-link` | Icon-only link button | 233, 265, 360-368 |
| `.btn-icon-destructive` | Icon-only destructive button | 234, 266, 372-380 |
| `.btn-sm-icon` | Small icon-only button | 235, 278-286 |
| `.btn-sm-icon-primary` | Small icon-only primary button | 236, 279, 313-321 |
| `.btn-sm-icon-secondary` | Small icon-only secondary button | 237, 280, 326-333 |
| `.btn-sm-icon-outline` | Small icon-only outline button | 238, 281, 338-345 |
| `.btn-sm-icon-ghost` | Small icon-only ghost button | 239, 282, 350-356 |
| `.btn-sm-icon-link` | Small icon-only link button | 240, 283, 361-368 |
| `.btn-sm-icon-destructive` | Small icon-only destructive button | 241, 284, 373-380 |
| `.btn-lg-icon` | Large icon-only button | 242, 296-304 |
| `.btn-lg-icon-primary` | Large icon-only primary button | 243, 297, 315-321 |
| `.btn-lg-icon-secondary` | Large icon-only secondary button | 244, 298, 327-333 |
| `.btn-lg-icon-outline` | Large icon-only outline button | 245, 299, 339-345 |
| `.btn-lg-icon-ghost` | Large icon-only ghost button | 246, 300, 351-356 |
| `.btn-lg-icon-link` | Large icon-only link button | 247, 301, 362-368 |
| `.btn-lg-icon-destructive` | Large icon-only destructive button | 248, 302, 374-380 |

---

## CSS Custom Properties (Design Tokens)

### Color Tokens
| Property | Description | Light Value | Dark Value |
|----------|-------------|-------------|------------|
| `--background` | Main background color | `oklch(1 0 0)` | `oklch(0.145 0 0)` |
| `--foreground` | Main text color | `oklch(0.145 0 0)` | `oklch(0.985 0 0)` |
| `--primary` | Primary brand color | `oklch(0.205 0 0)` | `oklch(0.922 0 0)` |
| `--primary-foreground` | Primary text color | `oklch(0.985 0 0)` | `oklch(0.205 0 0)` |
| `--secondary` | Secondary color | `oklch(0.97 0 0)` | `oklch(0.269 0 0)` |
| `--secondary-foreground` | Secondary text color | `oklch(0.205 0 0)` | `oklch(0.985 0 0)` |
| `--muted` | Muted background | `oklch(0.97 0 0)` | `oklch(0.269 0 0)` |
| `--muted-foreground` | Muted text color | `oklch(0.556 0 0)` | `oklch(0.708 0 0)` |
| `--accent` | Accent color | `oklch(0.97 0 0)` | `oklch(0.371 0 0)` |
| `--accent-foreground` | Accent text color | `oklch(0.205 0 0)` | `oklch(0.985 0 0)` |
| `--destructive` | Error/destructive color | `oklch(0.577 0.245 27.325)` | `oklch(0.704 0.191 22.216)` |
| `--border` | Border color | `oklch(0.922 0 0)` | `oklch(1 0 0 / 10%)` |
| `--input` | Input background | `oklch(0.922 0 0)` | `oklch(1 0 0 / 15%)` |
| `--ring` | Focus ring color | `oklch(0.708 0 0)` | `oklch(0.556 0 0)` |
| `--card` | Card background | `oklch(1 0 0)` | `oklch(0.205 0 0)` |
| `--card-foreground` | Card text color | `oklch(0.145 0 0)` | `oklch(0.985 0 0)` |
| `--popover` | Popover background | `oklch(1 0 0)` | `oklch(0.269 0 0)` |
| `--popover-foreground` | Popover text color | `oklch(0.145 0 0)` | `oklch(0.985 0 0)` |

### Chart Colors
| Property | Description | Light Value | Dark Value |
|----------|-------------|-------------|------------|
| `--chart-1` | Chart color 1 | `oklch(0.646 0.222 41.116)` | `oklch(0.488 0.243 264.376)` |
| `--chart-2` | Chart color 2 | `oklch(0.6 0.118 184.704)` | `oklch(0.696 0.17 162.48)` |
| `--chart-3` | Chart color 3 | `oklch(0.398 0.07 227.392)` | `oklch(0.769 0.188 70.08)` |
| `--chart-4` | Chart color 4 | `oklch(0.828 0.189 84.429)` | `oklch(0.627 0.265 303.9)` |
| `--chart-5` | Chart color 5 | `oklch(0.769 0.188 70.08)` | `oklch(0.645 0.246 16.439)` |

### Sidebar Tokens
| Property | Description | Light Value | Dark Value |
|----------|-------------|-------------|------------|
| `--sidebar` | Sidebar background | `oklch(0.985 0 0)` | `oklch(0.205 0 0)` |
| `--sidebar-foreground` | Sidebar text color | `oklch(0.145 0 0)` | `oklch(0.985 0 0)` |
| `--sidebar-primary` | Sidebar primary color | `oklch(0.205 0 0)` | `oklch(0.488 0.243 264.376)` |
| `--sidebar-primary-foreground` | Sidebar primary text | `oklch(0.985 0 0)` | `oklch(0.985 0 0)` |
| `--sidebar-accent` | Sidebar accent color | `oklch(0.97 0 0)` | `oklch(0.269 0 0)` |
| `--sidebar-accent-foreground` | Sidebar accent text | `oklch(0.205 0 0)` | `oklch(0.985 0 0)` |
| `--sidebar-border` | Sidebar border color | `oklch(0.922 0 0)` | `oklch(1 0 0 / 10%)` |
| `--sidebar-ring` | Sidebar focus ring | `oklch(0.708 0 0)` | `oklch(0.439 0 0)` |

### Sizing Tokens
| Property | Description | Value |
|----------|-------------|-------|
| `--radius` | Base border radius | `0.625rem` |
| `--sidebar-width` | Desktop sidebar width | `16rem` |
| `--sidebar-mobile-width` | Mobile sidebar width | `18rem` |

### Scrollbar Tokens
| Property | Description | Light Value | Dark Value |
|----------|-------------|-------------|------------|
| `--scrollbar-track` | Scrollbar track color | `transparent` | `transparent` |
| `--scrollbar-thumb` | Scrollbar thumb color | `rgba(0, 0, 0, 0.3)` | `rgba(255, 255, 255, 0.3)` |
| `--scrollbar-width` | Scrollbar width | `6px` | `6px` |
| `--scrollbar-radius` | Scrollbar border radius | `6px` | `6px` |

### Icon Tokens
| Property | Description | Light Value | Dark Value |
|----------|-------------|-------------|------------|
| `--chevron-down-icon` | Chevron down SVG icon | SVG with `oklch(0.556 0 0)` | SVG with `oklch(0.708 0 0)` |
| `--chevron-down-icon-50` | Chevron down 50% opacity | SVG with `oklch(0.556 0 0 / 0.5)` | SVG with `oklch(0.708 0 0 / 0.5)` |
| `--check-icon` | Check mark SVG icon | SVG with `oklch(0.556 0 0)` | SVG with `oklch(0.708 0 0)` |

---

## Computed Radius Tokens (@theme)
| Property | Description | Value |
|----------|-------------|-------|
| `--radius-sm` | Small radius | `calc(var(--radius) - 4px)` |
| `--radius-md` | Medium radius | `calc(var(--radius) - 2px)` |
| `--radius-lg` | Large radius | `var(--radius)` |
| `--radius-xl` | Extra large radius | `calc(var(--radius) + 4px)` |

---

## Key Implementation Notes

1. **Extensive Button System**: Basecoat provides comprehensive button variants covering size (sm/md/lg), style (primary/secondary/outline/ghost/link/destructive), and icon-only versions.

2. **Token-Driven Design**: All colors use CSS custom properties that adapt to light/dark themes automatically.

3. **Semantic Structure**: Components use semantic HTML elements (header, section, footer) as parts rather than explicit CSS part classes.

4. **Focus-First Accessibility**: All interactive components include comprehensive focus-visible states with ring indicators.

5. **Responsive Design**: Components include responsive behavior, especially the sidebar system with mobile/desktop breakpoints.

6. **Icon Integration**: Built-in support for SVG icons through CSS mask properties and background images.

7. **Form System**: Comprehensive form field system with validation states, labels, and various input types.

8. **Modern CSS Features**: Uses cutting-edge CSS features like `@starting-style`, `field-sizing`, and `calc-size()`.

This analysis provides a complete reference for implementing Basecoat UI components with proper token usage and variant patterns.

---

## JavaScript Behavior Analysis

### Core Architecture

#### 1. Basecoat Registration System (`basecoat.js`)

**Initialization Pattern:**
- Central registry for all components
- Auto-initialization on DOM ready
- MutationObserver for dynamic content
- Component lifecycle management

**Key APIs:**
```javascript
window.basecoat = {
  register: (name, selector, initFunction) => {},  // Register component
  init: (componentName) => {},                      // Reinit specific component
  initAll: () => {},                               // Reinit all components
  start: () => {},                                 // Start MutationObserver
  stop: () => {}                                   // Stop MutationObserver
}
```

**Registration Pattern:**
```javascript
window.basecoat.register('component-name', '.selector:not([data-component-initialized])', initFunction);
```

#### 2. Event System

**Global Events:**
- `basecoat:initialized` - Dispatched when component initializes
- `basecoat:popover` - Close other popovers when one opens
- `basecoat:sidebar` - Control sidebar state
- `basecoat:toast` - Create dynamic toasts
- `basecoat:locationchange` - Custom navigation event

---

### Component Patterns

#### Dropdown Menu (`dropdown-menu.js`)

**Required HTML Structure:**
```html
<div class="dropdown-menu">
  <button aria-expanded="false">Trigger</button>
  <div data-popover aria-hidden="true">
    <div role="menu">
      <button role="menuitem">Item 1</button>
      <button role="menuitem">Item 2</button>
    </div>
  </div>
</div>
```

**Behavior:**
- Click trigger to toggle menu
- Keyboard navigation (Arrow keys, Home/End, Enter, Escape)
- Mouse hover highlights items
- Auto-close on outside click
- Dispatches `basecoat:popover` to close other dropdowns

#### Sidebar (`sidebar.js`)

**Required HTML Structure:**
```html
<aside id="sidebar-id" class="sidebar" 
       data-initial-open="true" 
       data-initial-mobile-open="false"
       data-breakpoint="768">
  <nav>
    <a href="/page1">Link 1</a>
    <a href="/page2" data-ignore-current>Link 2</a>
    <a href="/page3" data-keep-mobile-sidebar-open>Link 3</a>
  </nav>
</aside>
```

**Behavior:**
- Responsive open/close based on breakpoint
- Auto-highlights current page link (`aria-current="page"`)
- Mobile: closes on link click (unless `data-keep-mobile-sidebar-open`)
- Patches history API for client-side navigation detection
- Controlled via `basecoat:sidebar` events

**Control Events:**
```javascript
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { id: 'sidebar-id', action: 'open' | 'close' | 'toggle' }
}));
```

#### Tabs (`tabs.js`)

**Required HTML Structure:**
```html
<div class="tabs">
  <div role="tablist">
    <button role="tab" aria-controls="panel1" aria-selected="true">Tab 1</button>
    <button role="tab" aria-controls="panel2" aria-selected="false">Tab 2</button>
  </div>
  <div id="panel1" role="tabpanel">Content 1</div>
  <div id="panel2" role="tabpanel" hidden>Content 2</div>
</div>
```

**Behavior:**
- Click to select tab
- Keyboard navigation (Arrow keys, Home/End)
- Automatic panel show/hide
- ARIA attributes management

#### Select (`select.js`)

**Required HTML Structure:**
```html
<div class="select" data-value="option1">
  <button class="trigger">
    <span class="placeholder">Choose...</span>
    <span class="value">Option 1</span>
  </button>
  <input type="hidden" name="field-name" value="option1">
  <div data-popover>
    <div class="filter">
      <input type="text" placeholder="Filter...">
    </div>
    <ul role="listbox">
      <li role="option" data-value="option1">Option 1</li>
      <li role="option" data-value="option2">Option 2</li>
    </ul>
  </div>
</div>
```

**Behavior:**
- Custom select with hidden input
- Filterable options
- Keyboard navigation
- Dispatches `change` event on value change
- Programmatic control: `selectComponent.selectByValue(value)`

#### Popover (`popover.js`)

**Required HTML Structure:**
```html
<div class="popover">
  <button>Trigger</button>
  <div data-popover aria-hidden="true">
    <div autofocus>Popover content</div>
  </div>
</div>
```

**Behavior:**
- Simple click toggle
- Auto-focus elements with `autofocus` attribute
- Escape to close
- Outside click to close
- Closes other popovers via event system

#### Toast (`toast.js`)

**Required HTML Structure:**
```html
<div class="toaster" id="toaster"></div>
```

**Dynamic Toast Creation:**
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    category: 'success' | 'error' | 'info' | 'warning',
    message: 'Toast message',
    duration: 3000,  // Optional, ms
    action: {        // Optional
      label: 'Undo',
      href: '/undo'  // or onclick function
    }
  }
}));
```

**Behavior:**
- Auto-dismiss after duration
- Pause on hover
- Manual dismiss button
- ARIA live regions for accessibility
- Category-specific icons

#### Command Palette (`command.js`)

**Required HTML Structure:**
```html
<div class="command">
  <input type="text" placeholder="Search commands..." role="combobox">
  <ul role="listbox">
    <li role="option" data-keywords="save file">
      <a href="/save">Save File</a>
    </li>
  </ul>
</div>
```

**Behavior:**
- Real-time filtering by text/keywords
- Keyboard navigation
- Auto-close parent dialog on selection
- Highlights active option

---

### Alpine.js Integration

While Basecoat provides vanilla JavaScript components, they're designed to work alongside Alpine.js:

1. **Non-conflicting**: Components use data attributes that don't interfere with Alpine
2. **Event-driven**: Can trigger/listen to Basecoat events from Alpine
3. **Progressive Enhancement**: Alpine can enhance Basecoat components

**Example Integration:**
```html
<div x-data="{ customBehavior: false }">
  <div class="dropdown-menu" @basecoat:initialized="customBehavior = true">
    <!-- Basecoat handles core dropdown behavior -->
    <!-- Alpine adds custom enhancements -->
  </div>
</div>
```

---

### Required HTML Patterns Summary

1. **Initialization Flag**: Components check for `data-[component]-initialized` to prevent double init
2. **ARIA Attributes**: All components require proper ARIA roles and properties
3. **Nested Structure**: Most components have specific parent-child relationships
4. **Data Attributes**: Configuration via `data-*` attributes
5. **Event Targets**: Specific elements for event handling (triggers, menus, panels)

---

### Best Practices

1. **Let Basecoat Initialize**: Don't manually initialize components unless needed
2. **Use Events for Control**: Communicate with components via custom events
3. **Maintain Structure**: Keep required HTML structure intact
4. **Accessibility First**: Don't remove ARIA attributes
5. **Progressive Enhancement**: Layer additional behavior with Alpine.js

This comprehensive analysis provides the complete reference for implementing and working with Basecoat JavaScript components.