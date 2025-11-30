# Link Component

Displays a link element for navigation and external references.

## Basic Usage

```html
<a class="link" href="/dashboard">Dashboard</a>
```

## CSS Classes

### Link Variants
- **`link`** - Default link styling with primary color
- **`link-muted`** - Muted/secondary link variant
- **`link-external`** - External link with visual indicator
- **`link-light`** - Light colored link for dark backgrounds

### Link States
- **`link-disabled`** - Disabled state styling
- **`link-no-underline`** - Remove underline on hover
- **`link-underline`** - Always show underline
- **`link-block`** - Block-level link

### Variant Combinations
You can combine external styling with any variant:
- `link link-external` - Default external link
- `link link-muted link-external` - Muted external link
- `link link-light link-external` - Light external link

## Component Attributes

### Link Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Link styling classes | Yes |
| `href` | string | Link destination URL | Yes |
| `target` | string | "_blank", "_self", "_parent", "_top" | Optional |
| `rel` | string | Relationship attributes | External links |
| `download` | string | Download filename | Optional |

### Accessibility
| Attribute | Description | When to Use |
|-----------|-------------|-------------|
| `aria-label` | Accessible label | When link text is unclear |
| `aria-describedby` | Additional description | For complex links |
| `role` | ARIA role override | Special cases only |

## HTML Structure

```html
<!-- Basic link -->
<a class="link" href="/page">Link Text</a>

<!-- External link -->
<a class="link link-external" href="https://example.com" target="_blank" rel="noopener noreferrer">
  External Site
</a>

<!-- Link with icon -->
<a class="link" href="/dashboard">
  <svg><!-- icon --></svg>
  Dashboard
</a>
```

## Examples

### Link Variants

```html
<!-- Default link -->
<a class="link" href="/dashboard">Dashboard</a>

<!-- Muted link -->
<a class="link-muted" href="/help">Help & Support</a>

<!-- Light link (for dark backgrounds) -->
<a class="link-light" href="/profile">User Profile</a>

<!-- External link -->
<a class="link link-external" href="https://example.com" target="_blank" rel="noopener noreferrer">
  Visit External Site
</a>

<!-- Muted external link -->
<a class="link link-muted link-external" href="https://docs.example.com" target="_blank" rel="noopener noreferrer">
  Documentation
</a>
```

### Link States

```html
<!-- Normal state -->
<a class="link" href="/page">Normal Link</a>

<!-- Disabled state -->
<a class="link-disabled" href="/page" tabindex="-1" aria-disabled="true">Disabled Link</a>

<!-- No underline on hover -->
<a class="link link-no-underline" href="/page">No Hover Underline</a>

<!-- Always underlined -->
<a class="link link-underline" href="/page">Always Underlined</a>

<!-- Block-level link -->
<a class="link link-block" href="/page">Block Level Link</a>
```

### Links with Icons

```html
<!-- Icon before text -->
<a class="link" href="/dashboard">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
    <circle cx="8.5" cy="8.5" r="1.5"/>
    <polyline points="21,15 16,10 5,21"/>
  </svg>
  Dashboard
</a>

<!-- Icon after text -->
<a class="link" href="/settings">
  Settings
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"/>
    <circle cx="12" cy="12" r="3"/>
  </svg>
</a>

<!-- External link with automatic icon -->
<a class="link link-external" href="https://github.com" target="_blank" rel="noopener noreferrer">
  GitHub Repository
</a>
```

### Navigation Links

```html
<!-- Breadcrumb navigation -->
<nav aria-label="Breadcrumb">
  <ol class="flex items-center space-x-2">
    <li><a class="link-muted" href="/">Home</a></li>
    <li class="text-muted-foreground">/</li>
    <li><a class="link-muted" href="/products">Products</a></li>
    <li class="text-muted-foreground">/</li>
    <li><span class="text-foreground">Product Name</span></li>
  </ol>
</nav>

<!-- Primary navigation -->
<nav class="flex space-x-6">
  <a class="link" href="/dashboard">Dashboard</a>
  <a class="link" href="/projects">Projects</a>
  <a class="link" href="/team">Team</a>
  <a class="link-muted" href="/settings">Settings</a>
</nav>

<!-- Footer navigation -->
<footer class="bg-muted p-6">
  <div class="grid grid-cols-4 gap-4">
    <div>
      <h3 class="font-semibold mb-2">Product</h3>
      <ul class="space-y-1">
        <li><a class="link-light" href="/features">Features</a></li>
        <li><a class="link-light" href="/pricing">Pricing</a></li>
        <li><a class="link-light" href="/docs">Documentation</a></li>
      </ul>
    </div>
  </div>
</footer>
```

### Download Links

```html
<!-- File download -->
<a class="link" href="/files/report.pdf" download="monthly-report.pdf">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
    <polyline points="7,10 12,15 17,10"/>
    <line x1="12" x2="12" y1="15" y2="3"/>
  </svg>
  Download Report (PDF)
</a>

<!-- Image download -->
<a class="link" href="/images/chart.png" download="sales-chart.png">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
    <circle cx="8.5" cy="8.5" r="1.5"/>
    <polyline points="21,15 16,10 5,21"/>
  </svg>
  Download Chart
</a>

<!-- Document download -->
<a class="link-muted" href="/docs/manual.docx" download>
  User Manual (DOCX)
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
    <polyline points="7,10 12,15 17,10"/>
    <line x1="12" x2="12" y1="15" y2="3"/>
  </svg>
</a>
```

### Social and External Links

```html
<!-- Social media links -->
<div class="flex space-x-4">
  <a class="link link-external" href="https://twitter.com/company" target="_blank" rel="noopener noreferrer">
    Twitter
  </a>
  <a class="link link-external" href="https://github.com/company" target="_blank" rel="noopener noreferrer">
    GitHub
  </a>
  <a class="link link-external" href="https://linkedin.com/company/company" target="_blank" rel="noopener noreferrer">
    LinkedIn
  </a>
</div>

<!-- Reference links -->
<p>
  Learn more about our 
  <a class="link link-external" href="https://docs.example.com/api" target="_blank" rel="noopener noreferrer">
    API Documentation
  </a>
  or check out our 
  <a class="link link-external" href="https://github.com/company/examples" target="_blank" rel="noopener noreferrer">
    code examples
  </a>
  on GitHub.
</p>

<!-- Help and support links -->
<div class="text-sm text-muted-foreground space-y-1">
  <div>
    Need help? Contact our 
    <a class="link" href="/support">support team</a>
    or visit our 
    <a class="link link-external" href="https://help.example.com" target="_blank" rel="noopener noreferrer">
      help center
    </a>.
  </div>
  <div>
    Found a bug? 
    <a class="link link-external" href="https://github.com/company/issues" target="_blank" rel="noopener noreferrer">
      Report it on GitHub
    </a>.
  </div>
</div>
```

### Responsive Link Patterns

```html
<!-- Mobile-friendly navigation -->
<nav class="block sm:flex sm:space-x-6 space-y-2 sm:space-y-0">
  <a class="link block sm:inline" href="/dashboard">Dashboard</a>
  <a class="link block sm:inline" href="/projects">Projects</a>
  <a class="link block sm:inline" href="/team">Team</a>
  <a class="link-muted block sm:inline" href="/settings">Settings</a>
</nav>

<!-- Truncated links -->
<div class="max-w-sm">
  <a class="link block truncate" href="/very-long-url-that-might-overflow">
    This is a very long link text that will be truncated
  </a>
</div>

<!-- Icon-only on mobile -->
<div class="flex space-x-4">
  <a class="link" href="/dashboard">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
    </svg>
    <span class="hidden sm:inline ml-1">Dashboard</span>
  </a>
  <a class="link" href="/profile">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2"/>
      <circle cx="12" cy="7" r="4"/>
    </svg>
    <span class="hidden sm:inline ml-1">Profile</span>
  </a>
</div>
```

## Accessibility Features

- **Keyboard Navigation**: All links support Tab navigation
- **Screen Reader Support**: Proper semantic HTML structure
- **Focus Indicators**: Automatic focus rings and states
- **External Link Indication**: Visual and semantic indicators for external links
- **Disabled State**: Proper `aria-disabled` attribute handling

### Enhanced Accessibility

```html
<!-- Link with additional description -->
<a class="link" href="/advanced-settings" aria-describedby="settings-help">
  Advanced Settings
</a>
<p id="settings-help" class="sr-only">
  Configure advanced system preferences and security options
</p>

<!-- Link with custom label -->
<a class="link" href="/profile/edit" aria-label="Edit your user profile">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
    <path d="M18.375 2.625a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.375-9.375Z"/>
  </svg>
  Edit
</a>

<!-- Skip link -->
<a class="link sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 bg-background p-2 rounded-md" href="#main-content">
  Skip to main content
</a>

<!-- External link with clear indication -->
<a class="link link-external" href="https://api.example.com/docs" target="_blank" rel="noopener noreferrer" aria-describedby="external-link-warning">
  API Documentation
</a>
<span id="external-link-warning" class="sr-only">
  Opens in a new tab
</span>
```

## JavaScript Integration

### Link State Management

```javascript
// Toggle link disabled state
function setLinkDisabled(link, disabled) {
  if (disabled) {
    link.classList.add('link-disabled');
    link.setAttribute('aria-disabled', 'true');
    link.setAttribute('tabindex', '-1');
  } else {
    link.classList.remove('link-disabled');
    link.removeAttribute('aria-disabled');
    link.removeAttribute('tabindex');
  }
}

// Track external link clicks
function trackExternalLink(link) {
  // Analytics tracking
  if (link.hostname !== window.location.hostname) {
    analytics.track('external_link_click', {
      url: link.href,
      text: link.textContent.trim()
    });
  }
}

// Add external link indicators dynamically
function markExternalLinks() {
  const links = document.querySelectorAll('a[href^="http"]');
  
  links.forEach(link => {
    if (link.hostname !== window.location.hostname) {
      link.classList.add('link-external');
      link.setAttribute('target', '_blank');
      link.setAttribute('rel', 'noopener noreferrer');
    }
  });
}
```

### React Integration

```jsx
import React from 'react';

function Link({ 
  href,
  variant = 'default',
  external = false,
  target,
  rel,
  children,
  className = '',
  ...props 
}) {
  const isExternal = external || (href && href.startsWith('http') && !href.includes(window.location.hostname));
  
  const linkClasses = [
    'link',
    variant !== 'default' && `link-${variant}`,
    isExternal && 'link-external',
    className
  ].filter(Boolean).join(' ');
  
  const linkProps = {
    href,
    className: linkClasses,
    target: target || (isExternal ? '_blank' : undefined),
    rel: rel || (isExternal ? 'noopener noreferrer' : undefined),
    ...props
  };
  
  return <a {...linkProps}>{children}</a>;
}

// Usage
function App() {
  return (
    <div className="space-y-4">
      <Link href="/dashboard">Dashboard</Link>
      <Link href="/help" variant="muted">Help</Link>
      <Link href="https://example.com" external>External Site</Link>
      <Link href="/download.pdf" download="file.pdf">
        Download File
      </Link>
    </div>
  );
}
```

## Best Practices

1. **Use Semantic HTML**: Always use `<a>` elements for links, not buttons
2. **External Links**: Always include `target="_blank"` and `rel="noopener noreferrer"`
3. **Accessibility**: Provide clear, descriptive link text
4. **Icon Usage**: Include meaningful text alongside icons
5. **Visual Hierarchy**: Use appropriate variants for link importance
6. **Focus Management**: Ensure proper focus indicators
7. **Download Attributes**: Use `download` attribute for file downloads
8. **URL Structure**: Use meaningful, readable URLs

## Common Patterns

### Card Links

```html
<!-- Card with clickable area -->
<article class="border border-border rounded-lg p-6 hover:shadow-md transition-shadow">
  <h3 class="text-lg font-semibold">
    <a class="link link-no-underline" href="/articles/1">
      How to Build Better UIs
    </a>
  </h3>
  <p class="text-muted-foreground mt-2">
    A comprehensive guide to creating user interfaces that users love.
  </p>
  <div class="mt-4 flex justify-between items-center">
    <span class="text-sm text-muted-foreground">March 15, 2024</span>
    <a class="link link-muted" href="/articles/1">Read more →</a>
  </div>
</article>
```

### Action Lists

```html
<div class="space-y-3">
  <div class="flex items-center justify-between p-3 border border-border rounded-lg">
    <div>
      <h4 class="font-medium">Project Settings</h4>
      <p class="text-sm text-muted-foreground">Configure project preferences</p>
    </div>
    <a class="link" href="/project/settings">Configure</a>
  </div>
  
  <div class="flex items-center justify-between p-3 border border-border rounded-lg">
    <div>
      <h4 class="font-medium">Team Management</h4>
      <p class="text-sm text-muted-foreground">Manage team members and roles</p>
    </div>
    <a class="link" href="/team">Manage</a>
  </div>
</div>
```

### Link Collections

```html
<!-- Related links -->
<aside class="bg-muted rounded-lg p-4">
  <h4 class="font-medium mb-3">Related Resources</h4>
  <ul class="space-y-2">
    <li>
      <a class="link" href="/guides/getting-started">Getting Started Guide</a>
    </li>
    <li>
      <a class="link link-external" href="https://api.example.com" target="_blank" rel="noopener noreferrer">
        API Reference
      </a>
    </li>
    <li>
      <a class="link" href="/tutorials">Video Tutorials</a>
    </li>
    <li>
      <a class="link link-external" href="https://community.example.com" target="_blank" rel="noopener noreferrer">
        Community Forum
      </a>
    </li>
  </ul>
</aside>
```

## Related Components

- [Button](./button.md) - For action elements styled as buttons
- [Navigation](./navigation.md) - For navigation menus and breadcrumbs
- [Badge](./badge.md) - For link status indicators
- [Icon](./icon.md) - For link icons and indicators