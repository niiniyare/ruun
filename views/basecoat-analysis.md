# Basecoat CSS Component Analysis

## CSS Custom Properties (Variables)

### Color Variables
| Variable | Light Mode | Dark Mode | Description |
|----------|------------|-----------|-------------|
| `--background` | `oklch(1 0 0)` | `oklch(0.145 0 0)` | Page background |
| `--foreground` | `oklch(0.145 0 0)` | `oklch(0.985 0 0)` | Default text color |
| `--primary` | `oklch(0.205 0 0)` | `oklch(0.922 0 0)` | Primary brand color |
| `--primary-foreground` | `oklch(0.985 0 0)` | `oklch(0.205 0 0)` | Text on primary |
| `--secondary` | `oklch(0.97 0 0)` | `oklch(0.269 0 0)` | Secondary color |
| `--secondary-foreground` | `oklch(0.205 0 0)` | `oklch(0.985 0 0)` | Text on secondary |
| `--muted` | `oklch(0.97 0 0)` | `oklch(0.269 0 0)` | Muted backgrounds |
| `--muted-foreground` | `oklch(0.556 0 0)` | `oklch(0.708 0 0)` | Muted text |
| `--accent` | `oklch(0.97 0 0)` | `oklch(0.371 0 0)` | Accent color |
| `--accent-foreground` | `oklch(0.205 0 0)` | `oklch(0.985 0 0)` | Text on accent |
| `--destructive` | `oklch(0.577 0.245 27.325)` | `oklch(0.704 0.191 22.216)` | Error/danger color |
| `--border` | `oklch(0.922 0 0)` | `oklch(1 0 0 / 10%)` | Default border color |
| `--input` | `oklch(0.922 0 0)` | `oklch(1 0 0 / 15%)` | Input border color |
| `--ring` | `oklch(0.708 0 0)` | `oklch(0.556 0 0)` | Focus ring color |

### Layout Variables
| Variable | Value | Description |
|----------|-------|-------------|
| `--radius` | `0.625rem` | Base border radius |
| `--radius-sm` | `calc(var(--radius) - 4px)` | Small radius |
| `--radius-md` | `calc(var(--radius) - 2px)` | Medium radius |
| `--radius-lg` | `var(--radius)` | Large radius |
| `--radius-xl` | `calc(var(--radius) + 4px)` | Extra large radius |
| `--sidebar-width` | `16rem` | Desktop sidebar width |
| `--sidebar-mobile-width` | `18rem` | Mobile sidebar width |

## Component Classes

### Alert Component
| Class | Description |
|-------|-------------|
| `.alert` | Base alert styles with card background |
| `.alert-destructive` | Destructive/error alert variant |

### Badge Component
| Class | Description |
|-------|-------------|
| `.badge` | Base badge (same as primary) |
| `.badge-primary` | Primary colored badge |
| `.badge-secondary` | Secondary colored badge |
| `.badge-destructive` | Destructive/error badge |
| `.badge-outline` | Outlined badge variant |

### Button Component

#### Base Classes
| Class | Description |
|-------|-------------|
| `.btn` | Default button (same as primary) |
| `.btn-primary` | Primary button |
| `.btn-secondary` | Secondary button |
| `.btn-outline` | Outlined button |
| `.btn-ghost` | Ghost button (no background) |
| `.btn-link` | Link-styled button |
| `.btn-destructive` | Destructive/danger button |

#### Size Variants
| Size | Regular | Icon Only |
|------|---------|-----------|
| Small | `.btn-sm`, `.btn-sm-primary`, etc. | `.btn-sm-icon`, `.btn-sm-icon-primary`, etc. |
| Default | `.btn`, `.btn-primary`, etc. | `.btn-icon`, `.btn-icon-primary`, etc. |
| Large | `.btn-lg`, `.btn-lg-primary`, etc. | `.btn-lg-icon`, `.btn-lg-icon-primary`, etc. |

#### States
- Disabled: `disabled` attribute
- Pressed: `aria-pressed="true"`
- Loading: Can be implemented with custom logic

### Button Group
| Class | Description |
|-------|-------------|
| `.button-group` | Container for grouped buttons |
| `[data-orientation="vertical"]` | Vertical button group |

### Card Component
| Class | Description |
|-------|-------------|
| `.card` | Card container with header, section, footer slots |

### Form Components

#### Checkbox
| Class | Description |
|-------|-------------|
| `.input[type='checkbox']` | Styled checkbox |
| `[role='switch']` | Switch variant of checkbox |

#### Input
| Class | Description |
|-------|-------------|
| `.input[type='text']` | Text input |
| `.input[type='email']` | Email input |
| `.input[type='password']` | Password input |
| `.input[type='number']` | Number input |
| `.input[type='file']` | File input |
| `.input[type='tel']` | Phone input |
| `.input[type='url']` | URL input |
| `.input[type='search']` | Search input |
| `.input[type='date']` | Date input |
| `.input[type='datetime-local']` | DateTime input |
| `.input[type='month']` | Month input |
| `.input[type='week']` | Week input |
| `.input[type='time']` | Time input |

#### Radio
| Class | Description |
|-------|-------------|
| `.input[type='radio']` | Styled radio button |

#### Range
| Class | Description |
|-------|-------------|
| `.input[type='range']` | Styled range slider |

#### Select
| Class | Description |
|-------|-------------|
| `select.select` | Native select dropdown |
| `.select` | Custom select component |

#### Textarea
| Class | Description |
|-------|-------------|
| `.textarea` | Styled textarea |

#### Field & Label
| Class | Description |
|-------|-------------|
| `.field` | Form field container |
| `.fieldset` | Fieldset container |
| `.label` | Form label |

### Navigation Components

#### Command (Command Palette)
| Class | Description |
|-------|-------------|
| `.command-dialog` | Command palette dialog |
| `.command` | Command container |

#### Dropdown Menu
| Class | Description |
|-------|-------------|
| `.dropdown-menu` | Dropdown menu container |

#### Popover
| Class | Description |
|-------|-------------|
| `.popover` | Popover container |
| `[data-popover]` | Popover content |

#### Sidebar
| Class | Description |
|-------|-------------|
| `.sidebar` | Sidebar container |
| `[data-side="left"]` | Left sidebar (default) |
| `[data-side="right"]` | Right sidebar |

### Layout Components

#### Dialog/Modal
| Class | Description |
|-------|-------------|
| `.dialog` | Modal dialog container |

#### Table
| Class | Description |
|-------|-------------|
| `.table` | Styled table |

#### Tabs
| Class | Description |
|-------|-------------|
| `.tabs` | Tabs container |

### Utility Components

#### Collapsible
| Class | Description |
|-------|-------------|
| `details` | Native details element with animation |

#### Kbd
| Class | Description |
|-------|-------------|
| `.kbd` | Keyboard key display |

#### Toast
| Class | Description |
|-------|-------------|
| `.toaster` | Toast container |
| `.toast` | Individual toast |
| `.toast-content` | Toast content wrapper |

#### Tooltip
| Class | Description |
|-------|-------------|
| `[data-tooltip]` | Element with tooltip |

## Data Attribute Patterns

### Positioning Attributes
| Attribute | Values | Used By |
|-----------|--------|---------|
| `data-side` | `top`, `bottom`, `left`, `right` | Popover, Tooltip, Sidebar |
| `data-align` | `start`, `end`, `center` | Popover, Tooltip, Toaster |

### State Attributes
| Attribute | Values | Used By |
|-----------|--------|---------|
| `aria-hidden` | `true`, `false` | Popover, Sidebar, Toast |
| `aria-selected` | `true`, `false` | Select options, Tabs |
| `aria-pressed` | `true`, `false` | Buttons |
| `aria-disabled` | `true`, `false` | Menu items, Options |
| `aria-invalid` | `true`, `false` | Form inputs |

### Component-Specific Attributes
| Attribute | Values | Used By |
|-----------|--------|---------|
| `data-orientation` | `vertical`, `horizontal` | Button Group |
| `data-variant` | `default`, `outline` | Sidebar items |
| `data-size` | `sm`, `default`, `lg` | Sidebar items |
| `data-empty` | Custom text | Command/Select empty state |

## State Classes

### Loading States
- `.is-loading` - Not explicitly defined, can be added
- Loading states handled via composition (e.g., adding spinner)

### Disabled States
- `disabled` attribute on form elements
- `disabled:opacity-50` utility applied

### Focus States
- `focus-visible:ring-[3px]` - Focus ring on interactive elements
- `focus-visible:border-ring` - Focus border color

### Hover States
- `hover:bg-accent` - Hover background for ghost buttons
- `hover:bg-primary/90` - Hover effect on primary buttons

### Invalid States
- `aria-invalid:border-destructive` - Error border
- `aria-invalid:ring-destructive/20` - Error focus ring

## Responsive Patterns

### Breakpoints Used
- `sm:` - Small screens and up
- `md:` - Medium screens and up
- `max-md:` - Mobile-specific styles

### Mobile-Specific Behaviors
- Sidebar becomes overlay on mobile
- Dialog/modal width adjustments
- Toast positioning adjustments

## Animation & Transitions

### Defined Animations
| Animation | Description |
|-----------|-------------|
| `toast-up` | Toast slide up animation |
| Details open/close | Smooth height transition |
| Popover show/hide | Scale and opacity transition |

### Transition Properties
- `transition-all` - Used for most interactive elements
- `transition-[color,box-shadow]` - Used for buttons
- `transition-[margin,opacity]` - Used for toasts
- Duration typically 200-300ms

## Icon Integration
- SVG icons expected to have `.size-4` class by default
- Icon sizing handled via `[&_svg]:size-4` selectors
- Data URLs for chevron and check icons included

## Accessibility Features
- ARIA attributes fully supported
- Focus management with visible focus rings
- Screen reader support via roles
- Keyboard navigation patterns