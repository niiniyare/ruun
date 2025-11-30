# Card Component

Displays a card with header, content, and footer sections for organizing related information.

## Basic Usage

```html
<div class="card">
  <header>
    <h2>Card Title</h2>
    <p>Card Description</p>
  </header>
  <section>
    <p>Card Content</p>
  </section>
  <footer>
    <p>Card Footer</p>
  </footer>
</div>
```

## CSS Classes

### Primary Classes
- **`card`** - Applied to the main container element

### Supporting Classes
- Standard semantic HTML elements (`header`, `section`, `footer`)
- Form classes when containing forms
- Button classes for actions in footer
- Text utilities for content styling

### Tailwind Utilities Used
- `w-full` - Full width cards
- `grid gap-*` - Grid layouts for form content
- `flex items-center` - Flexible footer layouts
- `text-sm` - Small text sizing
- `mt-*` - Margin top spacing
- `pl-*` - Padding left for lists
- `space-x-*` - Horizontal spacing
- `rounded-*` - Border radius utilities

## Component Attributes

### Card Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "card" | Yes |

### No JavaScript Required
This component is purely CSS-based and does not require JavaScript initialization.

## HTML Structure

```html
<div class="card">
  <!-- Header section -->
  <header>
    <h2>Card Title</h2>
    <p>Card Description (optional)</p>
  </header>
  
  <!-- Content section -->
  <section>
    <!-- Card content goes here -->
  </section>
  
  <!-- Footer section (optional) -->
  <footer>
    <!-- Card actions or metadata -->
  </footer>
</div>
```

## Examples

### Basic Information Card
```html
<div class="card">
  <header>
    <h2>Meeting Notes</h2>
    <p>Transcript from the meeting with the client.</p>
  </header>
  <section class="text-sm">
    <p>Client requested dashboard redesign with focus on mobile responsiveness.</p>
    <ol class="mt-4 flex list-decimal flex-col gap-2 pl-6">
      <li>New analytics widgets for daily/weekly metrics</li>
      <li>Simplified navigation menu</li>
      <li>Dark mode support</li>
      <li>Timeline: 6 weeks</li>
      <li>Follow-up meeting scheduled for next Tuesday</li>
    </ol>
  </section>
  <footer class="flex items-center">
    <div class="flex -space-x-2 [&_img]:ring-card [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full">
      <img alt="@hunvreus" src="https://github.com/hunvreus.png">
      <img alt="@shadcn" src="https://github.com/shadcn.png">
      <img alt="@adamwathan" src="https://github.com/adamwathan.png">
    </div>
  </footer>
</div>
```

### Login Form Card
```html
<div class="card w-full">
  <header>
    <h2>Login to your account</h2>
    <p>Enter your details below to login to your account</p>
  </header>
  <section>
    <form class="form grid gap-6">
      <div class="grid gap-2">
        <label for="email">Email</label>
        <input type="email" id="email">
      </div>
      <div class="grid gap-2">
        <div class="flex items-center gap-2">
          <label for="password">Password</label>
          <a href="#" class="ml-auto inline-block text-sm underline-offset-4 hover:underline">Forgot your password?</a>
        </div>
        <input type="password" id="password">
      </div>
    </form>
  </section>
  <footer class="flex flex-col items-center gap-2">
    <button type="button" class="btn w-full">Login</button>
    <button type="button" class="btn-outline w-full">Login with Google</button>
    <p class="mt-4 text-center text-sm">Don't have an account? <a href="#" class="underline-offset-4 hover:underline">Sign up</a></p>
  </footer>
</div>
```

### Simple Content Card
```html
<div class="card">
  <header>
    <h2>Project Status</h2>
  </header>
  <section>
    <p>Current phase: Development</p>
    <div class="mt-4">
      <div class="flex justify-between text-sm mb-1">
        <span>Progress</span>
        <span>75%</span>
      </div>
      <div class="w-full bg-gray-200 rounded-full h-2">
        <div class="bg-blue-600 h-2 rounded-full" style="width: 75%"></div>
      </div>
    </div>
  </section>
  <footer>
    <p class="text-sm text-muted-foreground">Last updated 2 hours ago</p>
  </footer>
</div>
```

### Card with Actions
```html
<div class="card">
  <header>
    <h2>Team Member</h2>
    <p>Software Developer</p>
  </header>
  <section>
    <div class="flex items-center gap-4">
      <img src="/avatar.jpg" alt="John Doe" class="size-12 rounded-full">
      <div>
        <h3 class="font-medium">John Doe</h3>
        <p class="text-sm text-muted-foreground">john.doe@example.com</p>
      </div>
    </div>
  </section>
  <footer class="flex gap-2">
    <button type="button" class="btn">View Profile</button>
    <button type="button" class="btn-outline">Send Message</button>
  </footer>
</div>
```

### Minimal Card (No Footer)
```html
<div class="card">
  <header>
    <h2>Quick Note</h2>
  </header>
  <section>
    <p>Remember to review the quarterly reports before the meeting tomorrow at 2 PM.</p>
  </section>
</div>
```

### Card with Badge
```html
<div class="card">
  <header>
    <div class="flex items-center justify-between">
      <h2>Feature Request</h2>
      <span class="badge">New</span>
    </div>
    <p>User authentication improvements</p>
  </header>
  <section>
    <p>Add support for two-factor authentication and social login options to improve security and user experience.</p>
    <div class="mt-4 flex gap-2">
      <span class="text-xs bg-blue-100 text-blue-800 px-2 py-1 rounded">Authentication</span>
      <span class="text-xs bg-green-100 text-green-800 px-2 py-1 rounded">Security</span>
    </div>
  </section>
  <footer class="flex justify-between items-center">
    <span class="text-sm text-muted-foreground">Priority: High</span>
    <button type="button" class="btn">Review</button>
  </footer>
</div>
```

### Statistics Card
```html
<div class="card">
  <header>
    <h2>Monthly Revenue</h2>
  </header>
  <section>
    <div class="text-3xl font-bold">$12,450</div>
    <div class="flex items-center gap-1 mt-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-green-600">
        <path d="m7 11 2-2-2-2" />
        <path d="M11 13h4" />
      </svg>
      <span class="text-sm text-green-600">+12.5% from last month</span>
    </div>
  </section>
  <footer>
    <button type="button" class="btn-ghost text-sm">View Details</button>
  </footer>
</div>
```

### List Card
```html
<div class="card">
  <header>
    <h2>Recent Activity</h2>
  </header>
  <section>
    <div class="space-y-3">
      <div class="flex items-center gap-3">
        <div class="size-2 bg-green-600 rounded-full"></div>
        <span class="text-sm">User registered</span>
        <span class="text-xs text-muted-foreground ml-auto">2 min ago</span>
      </div>
      <div class="flex items-center gap-3">
        <div class="size-2 bg-blue-600 rounded-full"></div>
        <span class="text-sm">New order placed</span>
        <span class="text-xs text-muted-foreground ml-auto">5 min ago</span>
      </div>
      <div class="flex items-center gap-3">
        <div class="size-2 bg-orange-600 rounded-full"></div>
        <span class="text-sm">Payment processed</span>
        <span class="text-xs text-muted-foreground ml-auto">10 min ago</span>
      </div>
    </div>
  </section>
</div>
```

## Layout Patterns

### Grid Layout
```html
<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
  <div class="card">
    <header><h2>Card 1</h2></header>
    <section><p>Content...</p></section>
  </div>
  
  <div class="card">
    <header><h2>Card 2</h2></header>
    <section><p>Content...</p></section>
  </div>
  
  <div class="card">
    <header><h2>Card 3</h2></header>
    <section><p>Content...</p></section>
  </div>
</div>
```

### Flex Layout
```html
<div class="flex flex-col lg:flex-row gap-6">
  <div class="card flex-1">
    <header><h2>Main Content</h2></header>
    <section><p>Primary information...</p></section>
  </div>
  
  <div class="card w-full lg:w-80">
    <header><h2>Sidebar</h2></header>
    <section><p>Additional info...</p></section>
  </div>
</div>
```

### Masonry Layout
```html
<div class="columns-1 md:columns-2 lg:columns-3 gap-6 space-y-6">
  <div class="card break-inside-avoid">
    <header><h2>Short Card</h2></header>
    <section><p>Brief content.</p></section>
  </div>
  
  <div class="card break-inside-avoid">
    <header><h2>Tall Card</h2></header>
    <section>
      <p>Much longer content that takes up more vertical space...</p>
      <p>Additional paragraphs...</p>
    </section>
  </div>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper `header`, `section`, and `footer` elements
- **Heading Structure**: Maintains logical heading hierarchy
- **Focus Management**: Interactive elements are keyboard accessible
- **Screen Reader Support**: Clear content structure for assistive technology

### Enhanced Accessibility
```html
<div class="card" role="article" aria-labelledby="card-title">
  <header>
    <h2 id="card-title">Accessible Card</h2>
    <p>Card with enhanced accessibility features</p>
  </header>
  <section>
    <p>Content that is properly structured for screen readers.</p>
  </section>
  <footer>
    <button type="button" class="btn" aria-describedby="card-title">
      Take Action
    </button>
  </footer>
</div>
```

## Styling Customization

### Card Variants
```css
/* Elevated card */
.card-elevated {
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
}

/* Bordered card */
.card-bordered {
  border: 2px solid var(--border);
}

/* Compact card */
.card-compact {
  padding: 1rem;
}

/* Large card */
.card-large {
  padding: 2rem;
}
```

### Custom Card Styling
```html
<div class="card bg-gradient-to-r from-blue-50 to-indigo-50 border-blue-200">
  <header>
    <h2 class="text-blue-900">Special Card</h2>
    <p class="text-blue-700">With custom styling</p>
  </header>
  <section class="text-blue-800">
    <p>Custom styled content...</p>
  </section>
</div>
```

## Best Practices

1. **Clear Hierarchy**: Use proper heading levels and semantic structure
2. **Consistent Spacing**: Apply consistent padding and margins
3. **Logical Order**: Place most important content first
4. **Action Placement**: Put primary actions in footer
5. **Content Density**: Don't overcrowd cards with too much information
6. **Responsive Design**: Ensure cards work well on all screen sizes
7. **Visual Balance**: Balance text and visual elements
8. **Loading States**: Provide skeleton states for dynamic content

## Common Patterns

### Dashboard Card
```html
<div class="card">
  <header class="flex items-center justify-between">
    <h2>Active Users</h2>
    <button type="button" class="btn-ghost text-sm">View All</button>
  </header>
  <section>
    <div class="text-3xl font-bold text-primary">1,234</div>
    <p class="text-sm text-muted-foreground mt-2">
      +5.2% increase from last week
    </p>
  </section>
</div>
```

### Product Card
```html
<div class="card">
  <section>
    <img src="/product.jpg" alt="Product" class="w-full h-48 object-cover rounded-t-lg">
  </section>
  <header class="pt-4">
    <h2>Product Name</h2>
    <p class="text-lg font-semibold text-primary">$99.99</p>
  </header>
  <section>
    <p class="text-sm text-muted-foreground">
      Brief product description highlighting key features.
    </p>
  </section>
  <footer>
    <button type="button" class="btn w-full">Add to Cart</button>
  </footer>
</div>
```

### Notification Card
```html
<div class="card border-l-4 border-l-blue-500 bg-blue-50">
  <header class="flex items-start gap-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-blue-600 mt-0.5">
      <circle cx="12" cy="12" r="10" />
      <path d="M12 16v-4" />
      <path d="M12 8h.01" />
    </svg>
    <div class="flex-1">
      <h2 class="text-blue-900">System Update</h2>
      <p class="text-blue-800 text-sm">New features are now available</p>
    </div>
  </header>
  <footer class="ml-8">
    <button type="button" class="btn-outline text-sm">Learn More</button>
  </footer>
</div>
```

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Card({ title, description, children, footer }) {
  return (
    <div className="card">
      {(title || description) && (
        <header>
          {title && <h2>{title}</h2>}
          {description && <p>{description}</p>}
        </header>
      )}
      
      <section>
        {children}
      </section>
      
      {footer && (
        <footer>
          {footer}
        </footer>
      )}
    </div>
  );
}

// Usage
<Card 
  title="User Profile" 
  description="Manage your account settings"
  footer={<button className="btn">Edit Profile</button>}
>
  <p>User information content...</p>
</Card>
```

### Vue Integration
```vue
<template>
  <div class="card">
    <header v-if="title || description">
      <h2 v-if="title">{{ title }}</h2>
      <p v-if="description">{{ description }}</p>
    </header>
    
    <section>
      <slot />
    </section>
    
    <footer v-if="$slots.footer">
      <slot name="footer" />
    </footer>
  </div>
</template>

<script>
export default {
  props: {
    title: String,
    description: String
  }
}
</script>
```

### HTMX Integration
```html
<div class="card" hx-get="/api/card-content" hx-trigger="load">
  <header>
    <h2>Dynamic Content</h2>
  </header>
  <section>
    <div class="skeleton">Loading...</div>
  </section>
</div>
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% macro card(title, description, footer_content) %}
<div class="card">
  {% if title or description %}
    <header>
      {% if title %}<h2>{{ title }}</h2>{% endif %}
      {% if description %}<p>{{ description }}</p>{% endif %}
    </header>
  {% endif %}
  
  <section>
    {{ caller() }}
  </section>
  
  {% if footer_content %}
    <footer>
      {{ footer_content | safe }}
    </footer>
  {% endif %}
</div>
{% endmacro %}

{% call card("User Settings", "Manage your account preferences") %}
  <form class="form grid gap-4">
    <div class="grid gap-2">
      <label for="username">Username</label>
      <input type="text" id="username" value="{{ user.username }}">
    </div>
  </form>
{% endcall %}
```

### Advanced Configuration
```jinja2
{% set footer %}
  <div class="flex justify-between">
    <button type="button" class="btn-outline">Cancel</button>
    <button type="button" class="btn">Save Changes</button>
  </div>
{% endset %}

{{ card(
  title="Edit Profile",
  description="Update your profile information",
  footer_content=footer
) }}
```

## Related Components

- [Button](./button.md) - For card actions and interactions
- [Badge](./badge.md) - For status indicators in cards
- [Avatar](./avatar.md) - For user representations in cards
- [Form](./form.md) - For form content within cards