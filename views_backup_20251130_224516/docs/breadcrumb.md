# Breadcrumb Component

Displays the path to the current resource using a hierarchy of links.

## Important Note

**There is no dedicated Breadcrumb component in Basecoat.** Instead, use the pattern shown below with semantic HTML and Tailwind utilities.

## Basic Usage

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Components</a>
  </li>
</ol>
```

## CSS Classes

### Primary Classes
No dedicated breadcrumb classes - uses standard semantic HTML with utility classes.

### Tailwind Utilities Used
- `text-muted-foreground` - Muted text color for overall breadcrumb
- `flex flex-wrap` - Flexible, wrapping layout
- `items-center` - Vertical centering
- `gap-1.5` - Small spacing between elements
- `text-sm` - Small text size
- `break-words` - Word breaking for long links
- `sm:gap-2.5` - Responsive spacing
- `inline-flex` - Inline flex layout for links
- `hover:text-foreground` - Hover color change
- `transition-colors` - Smooth color transitions
- `size-3.5` - Icon sizing
- `text-foreground` - Full color for current page
- `font-normal` - Normal font weight

## HTML Structure

### Basic Breadcrumb
```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Breadcrumb item -->
  <li class="inline-flex items-center gap-1.5">
    <a href="/path">Link Text</a>
  </li>
  
  <!-- Separator -->
  <li>
    <!-- Chevron right icon -->
  </li>
  
  <!-- Current page (no link) -->
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Current Page</span>
  </li>
</ol>
```

### Separator Icon
Use a chevron-right icon between breadcrumb items:

```html
<li>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
    <path d="m9 18 6-6-6-6" />
  </svg>
</li>
```

## Examples

### Simple Breadcrumb
```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="/" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="/components" class="hover:text-foreground transition-colors">Components</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Breadcrumb</span>
  </li>
</ol>
```

### Breadcrumb with Dropdown Menu
For collapsed or overflow items, use the dropdown menu pattern:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <div id="breadcrumb-menu" class="dropdown-menu">
      <button type="button" id="breadcrumb-menu-trigger" aria-haspopup="menu" aria-controls="breadcrumb-menu-menu" aria-expanded="false" class="flex size-9 items-center justify-center h-4 w-4 hover:text-foreground cursor-pointer">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </button>
      <div id="breadcrumb-menu-popover" data-popover aria-hidden="true">
        <div role="menu" id="breadcrumb-menu-menu" aria-labelledby="breadcrumb-menu-trigger">
          <nav role="menu">
            <button type="button" role="menuitem">Documentation</button>
            <button type="button" role="menuitem">Themes</button>
            <button type="button" role="menuitem">GitHub</button>
          </nav>
        </div>
      </div>
    </div>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Components</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Breadcrumb</span>
  </li>
</ol>
```

### Breadcrumb with Icons
Add icons to breadcrumb items for enhanced visual hierarchy:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
      <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
      <polyline points="9,22 9,12 15,12 15,22" />
    </svg>
    <a href="/" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
      <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z" />
      <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z" />
    </svg>
    <a href="/docs" class="hover:text-foreground transition-colors">Documentation</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Current Page</span>
  </li>
</ol>
```

## Navigation Patterns

### Mobile Responsive
The breadcrumb automatically wraps on smaller screens:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Items will wrap naturally on mobile -->
</ol>
```

### Condensed for Mobile
For mobile, consider showing only the parent and current page:

```html
<!-- Desktop: Full breadcrumb -->
<ol class="hidden sm:flex text-muted-foreground flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Full breadcrumb trail -->
</ol>

<!-- Mobile: Condensed breadcrumb -->
<ol class="sm:hidden text-muted-foreground flex items-center gap-1.5 text-sm">
  <li class="inline-flex items-center gap-1.5">
    <a href="/parent" class="hover:text-foreground transition-colors">← Back</a>
  </li>
</ol>
```

## Accessibility Features

- **Semantic HTML**: Uses `<ol>` for proper list semantics
- **Navigation Landmark**: Consider wrapping in `<nav>` with `aria-label`
- **Current Page**: Non-interactive element for current location
- **Keyboard Navigation**: All links are keyboard accessible
- **Screen Reader Support**: Proper list and link semantics

### Enhanced Accessibility
```html
<nav aria-label="Breadcrumb">
  <ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
    <li class="inline-flex items-center gap-1.5">
      <a href="/" class="hover:text-foreground transition-colors">Home</a>
    </li>
    <li>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5" aria-hidden="true">
        <path d="m9 18 6-6-6-6" />
      </svg>
    </li>
    <li class="inline-flex items-center gap-1.5">
      <span class="text-foreground font-normal" aria-current="page">Current Page</span>
    </li>
  </ol>
</nav>
```

## JavaScript Integration

### Dynamic Breadcrumbs
```javascript
function generateBreadcrumb(path) {
  const breadcrumbContainer = document.getElementById('breadcrumb');
  const pathSegments = path.split('/').filter(segment => segment);
  
  let breadcrumbHTML = '<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">';
  
  // Home link
  breadcrumbHTML += `
    <li class="inline-flex items-center gap-1.5">
      <a href="/" class="hover:text-foreground transition-colors">Home</a>
    </li>
  `;
  
  pathSegments.forEach((segment, index) => {
    const isLast = index === pathSegments.length - 1;
    const href = '/' + pathSegments.slice(0, index + 1).join('/');
    const title = segment.charAt(0).toUpperCase() + segment.slice(1);
    
    // Separator
    breadcrumbHTML += `
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </li>
    `;
    
    // Breadcrumb item
    if (isLast) {
      breadcrumbHTML += `
        <li class="inline-flex items-center gap-1.5">
          <span class="text-foreground font-normal" aria-current="page">${title}</span>
        </li>
      `;
    } else {
      breadcrumbHTML += `
        <li class="inline-flex items-center gap-1.5">
          <a href="${href}" class="hover:text-foreground transition-colors">${title}</a>
        </li>
      `;
    }
  });
  
  breadcrumbHTML += '</ol>';
  breadcrumbContainer.innerHTML = breadcrumbHTML;
}

// Usage
generateBreadcrumb(window.location.pathname);
```

### HTMX Integration
```html
<nav aria-label="Breadcrumb" hx-get="/breadcrumb" hx-trigger="path-change">
  <!-- Breadcrumb content loaded dynamically -->
</nav>
```

## Best Practices

1. **Start with Home**: Always begin breadcrumbs with the home/root page
2. **Current Page**: Don't link the current page, use a `<span>` instead
3. **Meaningful Labels**: Use clear, descriptive text for each level
4. **Reasonable Length**: Limit to 5-7 levels to avoid overcrowding
5. **Mobile Consideration**: Use responsive design or mobile-specific patterns
6. **Accessibility**: Include proper ARIA labels and semantic HTML
7. **Visual Hierarchy**: Use consistent styling and separators
8. **Dropdown for Overflow**: Use dropdown menus for collapsed intermediate levels

## Integration with Routing

### React Router
```jsx
import { useLocation } from 'react-router-dom';

function Breadcrumb() {
  const location = useLocation();
  const pathnames = location.pathname.split('/').filter(x => x);
  
  return (
    <nav aria-label="Breadcrumb">
      <ol className="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
        <li className="inline-flex items-center gap-1.5">
          <a href="/" className="hover:text-foreground transition-colors">Home</a>
        </li>
        {pathnames.map((pathname, index) => {
          const routeTo = `/${pathnames.slice(0, index + 1).join('/')}`;
          const isLast = index === pathnames.length - 1;
          
          return (
            <React.Fragment key={routeTo}>
              <li>
                <svg className="size-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                  <path d="m9 18 6-6-6-6" />
                </svg>
              </li>
              <li className="inline-flex items-center gap-1.5">
                {isLast ? (
                  <span className="text-foreground font-normal" aria-current="page">
                    {pathname}
                  </span>
                ) : (
                  <a href={routeTo} className="hover:text-foreground transition-colors">
                    {pathname}
                  </a>
                )}
              </li>
            </React.Fragment>
          );
        })}
      </ol>
    </nav>
  );
}
```

## Related Components

- [Dropdown Menu](./dropdown-menu.md) - For collapsed breadcrumb items
- [Button](./button.md) - For interactive elements
- [Link](./link.md) - For navigation links
- [Navigation](./navigation.md) - For main site navigation