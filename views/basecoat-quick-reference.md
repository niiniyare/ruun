# Basecoat CSS Quick Reference

## Component Class Patterns

### Buttons
```html
<!-- Single combined classes -->
<button class="btn">Default (Primary)</button>
<button class="btn-secondary">Secondary</button>
<button class="btn-sm-outline">Small Outline</button>
<button class="btn-lg-destructive">Large Destructive</button>
<button class="btn-icon">Icon Only</button>
<button class="btn-sm-icon-ghost">Small Icon Ghost</button>

<!-- States via attributes -->
<button class="btn" disabled>Disabled</button>
<button class="btn" aria-pressed="true">Pressed</button>
```

### Form Elements
```html
<!-- Form field wrapper -->
<div class="field" data-invalid="true">
  <label>Email</label>
  <input type="email" placeholder="Enter email">
  <p role="alert">Error message appears here</p>
</div>

<!-- Switch -->
<input type="checkbox" role="switch" checked>

<!-- Select -->
<select class="select">
  <option>Choose option</option>
</select>
```

### Layout Components
```html
<!-- Card -->
<div class="card">
  <header>
    <h2>Card Title</h2>
    <p>Description</p>
  </header>
  <section>
    Card content goes here
  </section>
  <footer>
    Action buttons
  </footer>
</div>

<!-- Sidebar -->
<div class="sidebar" data-side="left">
  <nav>
    <section>Navigation content</section>
  </nav>
</div>

<!-- Dialog -->
<dialog class="dialog" open>
  <div>
    <header>
      <h2>Dialog Title</h2>
    </header>
    <section>Dialog content</section>
    <footer>Action buttons</footer>
  </div>
</dialog>
```

### Navigation
```html
<!-- Dropdown Menu -->
<div class="dropdown-menu">
  <button>Trigger</button>
  <div data-popover aria-hidden="true">
    <div role="menuitem">Item 1</div>
    <div role="menuitem">Item 2</div>
  </div>
</div>

<!-- Tabs -->
<div class="tabs">
  <div role="tablist">
    <button role="tab" aria-selected="true">Tab 1</button>
    <button role="tab">Tab 2</button>
  </div>
  <div role="tabpanel">Tab content</div>
</div>
```

### Feedback Components
```html
<!-- Alert -->
<div class="alert">
  <svg>...</svg>
  <h3>Alert Title</h3>
  <section>
    <p>Alert message content</p>
  </section>
</div>

<!-- Badge -->
<span class="badge">Default</span>
<span class="badge-secondary">Secondary</span>
<span class="badge-destructive">Error</span>
<span class="badge-outline">Outline</span>

<!-- Toast -->
<div class="toaster" data-align="end">
  <div class="toast">
    <div class="toast-content">
      <svg>...</svg>
      <section>
        <h2>Toast Title</h2>
        <p>Toast message</p>
      </section>
    </div>
  </div>
</div>
```

## Data Attributes Reference

### State Attributes
| Attribute | Values | Usage |
|-----------|--------|-------|
| `aria-hidden` | `true`, `false` | Show/hide popover, sidebar, toast |
| `aria-selected` | `true`, `false` | Tab/option selection |
| `aria-pressed` | `true`, `false` | Button pressed state |
| `aria-invalid` | `true`, `false` | Form validation errors |
| `disabled` | boolean | Disable interactive elements |
| `data-invalid` | `true`, `false` | Field validation state |

### Layout Attributes
| Attribute | Values | Usage |
|-----------|--------|-------|
| `data-side` | `top`, `bottom`, `left`, `right` | Popover/sidebar position |
| `data-align` | `start`, `end`, `center` | Popover/toast alignment |
| `data-orientation` | `vertical`, `horizontal` | Layout direction |
| `data-variant` | `default`, `outline` | Component styling variant |
| `data-size` | `sm`, `default`, `lg` | Component size |

### Component-Specific
| Attribute | Values | Usage |
|-----------|--------|-------|
| `role` | `switch`, `tab`, `tabpanel`, `menuitem`, etc. | Semantic meaning |
| `data-empty` | Custom text | Empty state message |
| `open` | boolean | Dialog/details open state |

## CSS Custom Properties

### Colors (Light/Dark)
```css
--background: oklch(1 0 0) / oklch(0.145 0 0)
--foreground: oklch(0.145 0 0) / oklch(0.985 0 0)
--primary: oklch(0.205 0 0) / oklch(0.922 0 0)
--secondary: oklch(0.97 0 0) / oklch(0.269 0 0)
--muted: oklch(0.97 0 0) / oklch(0.269 0 0)
--border: oklch(0.922 0 0) / oklch(1 0 0 / 10%)
--destructive: oklch(0.577 0.245 27.325) / oklch(0.704 0.191 22.216)
```

### Layout
```css
--radius: 0.625rem
--sidebar-width: 16rem
--sidebar-mobile-width: 18rem
```

## JavaScript Integration

### Component Initialization
```javascript
// Import component
import { DropdownMenu } from '/static/dist/basecoat/dropdown-menu.js';

// Initialize
const dropdown = new DropdownMenu(document.querySelector('.dropdown-menu'));

// Available components:
// - DropdownMenu
// - Sidebar
// - Dialog
// - Command
// - Tabs
// - Select (for searchable)
```

### Event Handling
```javascript
// Dropdown events
dropdown.on('open', () => console.log('Opened'));
dropdown.on('close', () => console.log('Closed'));
dropdown.on('select', (value) => console.log('Selected:', value));

// Manual control
dropdown.open();
dropdown.close();
dropdown.toggle();
```

## Common Patterns

### Form with Validation
```html
<form>
  <div class="field" data-invalid="false">
    <label for="email">Email</label>
    <input type="email" id="email" required>
    <p>Enter your email address</p>
  </div>
  
  <div class="field" data-invalid="true">
    <label for="password">Password</label>
    <input type="password" id="password" required>
    <p role="alert">Password must be at least 8 characters</p>
  </div>
  
  <button type="submit" class="btn">Submit</button>
</form>
```

### Card with Actions
```html
<div class="card">
  <header>
    <h2>Project Name</h2>
    <p>Project description and details</p>
  </header>
  <section>
    <p>Main content area with project information.</p>
  </section>
  <footer>
    <button class="btn-outline">Edit</button>
    <button class="btn">View</button>
  </footer>
</div>
```

### Data Table
```html
<table class="table">
  <caption>User List</caption>
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
      <td><span class="badge-secondary">Active</span></td>
    </tr>
  </tbody>
</table>
```

## Responsive Behavior

### Mobile-First Breakpoints
- Components adapt automatically
- Sidebar becomes overlay on mobile (`max-md:` classes)
- Dialog width adjusts (`sm:max-w-lg`)
- Toast positioning adapts

### Dark Mode
- Automatic via `.dark` class on `<html>`
- All components support both themes
- Uses CSS custom properties for colors

## Debugging Tips

1. **Check data attributes**: Use browser devtools to verify attributes
2. **Inspect CSS variables**: See computed styles for color values
3. **Validate HTML**: Ensure proper semantic structure
4. **Test keyboard navigation**: Tab, Enter, Escape, Arrow keys
5. **Screen reader testing**: Use browser's built-in reader

## Browser Support

- Modern browsers (Chrome 90+, Firefox 88+, Safari 14+)
- CSS Grid and Flexbox required
- JavaScript modules for interactive components
- Graceful degradation without JavaScript