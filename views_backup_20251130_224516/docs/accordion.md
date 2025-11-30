# Accordion Component

A vertically stacked set of interactive headings that each reveal a section of content.

## Basic Usage

```html
<section class="accordion">
  <details class="group border-b last:border-b-0">
    <summary class="w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md">
      <h2 class="flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline">
        Is it accessible?
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Yes. It adheres to the WAI-ARIA design pattern.</p>
    </section>
  </details>
  
  <details class="group border-b last:border-b-0">
    <summary class="w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md">
      <h2 class="flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline">
        Is it styled?
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Yes. It comes with default styles that matches the other components' aesthetic.</p>
    </section>
  </details>
</section>
```

## Required JavaScript (Optional)

For single-item expansion behavior (only one item open at a time):

```javascript
(() => {
  const accordions = document.querySelectorAll(".accordion");
  accordions.forEach((accordion) => {
    accordion.addEventListener("click", (event) => {
      const summary = event.target.closest("summary");
      if (!summary) return;
      const details = summary.closest("details");
      if (!details) return;
      accordion.querySelectorAll("details").forEach((detailsEl) => {
        if (detailsEl !== details) {
          detailsEl.removeAttribute("open");
        }
      });
    });
  });
})();
```

## CSS Classes

### Primary Classes
- **`accordion`** - Applied to the container section element
- **`group`** - Applied to each details element for state-based styling

### Supporting Classes
- **`border-b`** - Bottom border for each item
- **`last:border-b-0`** - Remove border from last item
- **`w-full`** - Full width summary elements
- **`focus-visible:*`** - Focus state styling
- **`transition-all`** - Smooth transitions

### Tailwind Utilities Used
- `flex` - Flexible layout for headers
- `items-start` - Align items to start
- `justify-between` - Space between title and icon
- `gap-4` - Gap spacing
- `py-4` - Vertical padding
- `text-left` - Left-aligned text
- `text-sm` - Small text size
- `font-medium` - Medium font weight
- `hover:underline` - Underline on hover
- `group-open:rotate-180` - Rotate icon when open
- `transition-transform` - Smooth icon rotation
- `duration-200` - Animation duration

## Component Attributes

### Section Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Should include "accordion" | Recommended |

### Details Attributes  
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Should include "group" for styling | Recommended |
| `open` | boolean | Whether item is expanded | No |

### Summary Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Styling classes | Recommended |

## HTML Structure

```html
<section class="accordion">
  <details class="group">
    <summary>
      <h2>
        <!-- Accordion title -->
        <svg><!-- Chevron icon --></svg>
      </h2>
    </summary>
    <section>
      <!-- Accordion content -->
    </section>
  </details>
  <!-- Additional accordion items -->
</section>
```

## Native HTML Details Animation

Basecoat includes default animations for `<details>` elements. The accordion component builds on this with additional styling and optional JavaScript for single-item expansion.

## Examples

### Basic Accordion
```html
<section class="accordion">
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Question 1
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Answer to question 1...</p>
    </section>
  </details>
  
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Question 2
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Answer to question 2...</p>
    </section>
  </details>
</section>
```

### Accordion with Default Open Item
```html
<section class="accordion">
  <details class="group border-b" open>
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Getting Started
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Welcome to our getting started guide...</p>
      <ul class="mt-3 space-y-2 text-sm">
        <li>• Step 1: Install the package</li>
        <li>• Step 2: Configure settings</li>
        <li>• Step 3: Start building</li>
      </ul>
    </section>
  </details>
  
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Advanced Features
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm">Learn about advanced features...</p>
    </section>
  </details>
</section>
```

### Accordion with Rich Content
```html
<section class="accordion">
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Product Features
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <div class="grid gap-4">
        <div class="flex gap-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary mt-0.5">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          <div>
            <h3 class="text-sm font-medium">Real-time Collaboration</h3>
            <p class="text-sm text-muted-foreground mt-1">Work together with your team in real-time</p>
          </div>
        </div>
        
        <div class="flex gap-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary mt-0.5">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          <div>
            <h3 class="text-sm font-medium">Advanced Security</h3>
            <p class="text-sm text-muted-foreground mt-1">Enterprise-grade security features</p>
          </div>
        </div>
      </div>
    </section>
  </details>
</section>
```

### FAQ Accordion
```html
<section class="accordion max-w-2xl mx-auto">
  <h1 class="text-2xl font-bold mb-6">Frequently Asked Questions</h1>
  
  <details class="group border-b">
    <summary class="w-full outline-none cursor-pointer">
      <h2 class="flex justify-between py-4 text-sm font-medium hover:text-primary">
        What payment methods do you accept?
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm text-muted-foreground">
        We accept all major credit cards (Visa, Mastercard, American Express), 
        PayPal, and bank transfers for enterprise customers.
      </p>
    </section>
  </details>
  
  <details class="group border-b">
    <summary class="w-full outline-none cursor-pointer">
      <h2 class="flex justify-between py-4 text-sm font-medium hover:text-primary">
        Can I cancel my subscription anytime?
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm text-muted-foreground">
        Yes, you can cancel your subscription at any time. 
        Your service will continue until the end of your current billing period.
      </p>
    </section>
  </details>
  
  <details class="group border-b">
    <summary class="w-full outline-none cursor-pointer">
      <h2 class="flex justify-between py-4 text-sm font-medium hover:text-primary">
        Do you offer refunds?
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm text-muted-foreground">
        We offer a 30-day money-back guarantee for all new customers. 
        If you're not satisfied, contact our support team for a full refund.
      </p>
    </section>
  </details>
</section>
```

### Nested Accordion
```html
<section class="accordion">
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        API Documentation
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <p class="text-sm mb-4">Explore our API documentation by category:</p>
      
      <!-- Nested accordion -->
      <div class="ml-4 space-y-2">
        <details class="group">
          <summary class="cursor-pointer">
            <span class="text-sm font-medium hover:underline">Authentication</span>
          </summary>
          <div class="pl-4 pt-2">
            <p class="text-sm text-muted-foreground">Learn about API authentication methods...</p>
          </div>
        </details>
        
        <details class="group">
          <summary class="cursor-pointer">
            <span class="text-sm font-medium hover:underline">Endpoints</span>
          </summary>
          <div class="pl-4 pt-2">
            <p class="text-sm text-muted-foreground">Browse available API endpoints...</p>
          </div>
        </details>
      </div>
    </section>
  </details>
</section>
```

## Accessibility Features

- **Semantic HTML**: Uses native `<details>` and `<summary>` elements
- **Keyboard Navigation**: Space/Enter keys toggle items
- **Screen Reader Support**: Native announcement of expanded/collapsed state
- **Focus Management**: Proper focus indicators
- **ARIA Compliance**: Follows WAI-ARIA patterns

### Enhanced Accessibility
```html
<section class="accordion" role="region" aria-labelledby="accordion-title">
  <h2 id="accordion-title" class="sr-only">Frequently Asked Questions</h2>
  
  <details class="group border-b">
    <summary class="w-full outline-none" aria-label="Toggle answer for: What is your return policy?">
      <h3 class="flex justify-between py-4 text-sm font-medium">
        What is your return policy?
        <svg aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h3>
    </summary>
    <section class="pb-4" role="region" aria-labelledby="return-policy">
      <p id="return-policy" class="text-sm">
        We offer a 30-day return policy for all items...
      </p>
    </section>
  </details>
</section>
```

## JavaScript Behavior Options

### Allow Multiple Open Items (Default)
```javascript
// No JavaScript needed - native HTML behavior
```

### Single Item Expansion
```javascript
// Only one item can be open at a time
(() => {
  const accordions = document.querySelectorAll(".accordion");
  accordions.forEach((accordion) => {
    accordion.addEventListener("click", (event) => {
      const summary = event.target.closest("summary");
      if (!summary) return;
      const details = summary.closest("details");
      if (!details) return;
      accordion.querySelectorAll("details").forEach((detailsEl) => {
        if (detailsEl !== details) {
          detailsEl.removeAttribute("open");
        }
      });
    });
  });
})();
```

### Programmatic Control
```javascript
// Open specific item
document.querySelector("#accordion-item-1").setAttribute("open", true);

// Close all items
document.querySelectorAll(".accordion details").forEach(item => {
  item.removeAttribute("open");
});

// Toggle item
const item = document.querySelector("#accordion-item-1");
item.hasAttribute("open") ? item.removeAttribute("open") : item.setAttribute("open", true);
```

## Styling Customization

### Custom Icon
```html
<summary class="w-full outline-none">
  <h2 class="flex justify-between py-4 text-sm font-medium">
    Title
    <!-- Plus/Minus icon -->
    <span class="text-muted-foreground">
      <span class="group-open:hidden">+</span>
      <span class="hidden group-open:inline">−</span>
    </span>
  </h2>
</summary>
```

### Bordered Style
```html
<section class="accordion">
  <details class="group border rounded-lg mb-2 overflow-hidden">
    <summary class="w-full outline-none bg-muted px-4">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Title
      </h2>
    </summary>
    <section class="px-4 pb-4 pt-2">
      <p class="text-sm">Content...</p>
    </section>
  </details>
</section>
```

## Best Practices

1. **Clear Headings**: Use descriptive, concise headings
2. **Logical Order**: Arrange items in a logical sequence
3. **Icon Feedback**: Provide visual indicator of state
4. **Smooth Animations**: Use CSS transitions for state changes
5. **Touch Targets**: Ensure adequate touch target size on mobile
6. **Progressive Enhancement**: Works without JavaScript
7. **Semantic Structure**: Use proper heading hierarchy
8. **Content Length**: Keep expanded content concise

## Common Patterns

### Settings Accordion
```html
<section class="accordion">
  <details class="group border-b">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        <span class="flex items-center gap-3">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
          Account Settings
        </span>
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4 pl-9">
      <div class="space-y-3">
        <button type="button" class="text-sm hover:underline">Change Password</button>
        <button type="button" class="text-sm hover:underline">Update Email</button>
        <button type="button" class="text-sm hover:underline">Two-Factor Authentication</button>
      </div>
    </section>
  </details>
</section>
```

### Filter Accordion
```html
<section class="accordion">
  <details class="group border-b" open>
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Price Range
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4">
      <div class="space-y-2">
        <label class="flex items-center gap-2 text-sm">
          <input type="checkbox" class="checkbox">
          Under $25
        </label>
        <label class="flex items-center gap-2 text-sm">
          <input type="checkbox" class="checkbox">
          $25 - $50
        </label>
        <label class="flex items-center gap-2 text-sm">
          <input type="checkbox" class="checkbox">
          $50 - $100
        </label>
      </div>
    </section>
  </details>
</section>
```

## Integration Examples

### React Integration
```jsx
import React, { useState } from 'react';

function Accordion({ items, allowMultiple = false }) {
  const [openItems, setOpenItems] = useState(new Set());

  const toggleItem = (index) => {
    const newOpenItems = new Set(openItems);
    
    if (allowMultiple) {
      if (newOpenItems.has(index)) {
        newOpenItems.delete(index);
      } else {
        newOpenItems.add(index);
      }
    } else {
      newOpenItems.clear();
      if (!openItems.has(index)) {
        newOpenItems.add(index);
      }
    }
    
    setOpenItems(newOpenItems);
  };

  return (
    <section className="accordion">
      {items.map((item, index) => (
        <details 
          key={index} 
          className="group border-b last:border-b-0"
          open={openItems.has(index)}
        >
          <summary 
            className="w-full outline-none cursor-pointer"
            onClick={(e) => {
              e.preventDefault();
              toggleItem(index);
            }}
          >
            <h2 className="flex justify-between py-4 text-sm font-medium">
              {item.title}
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="transition-transform group-open:rotate-180">
                <path d="m6 9 6 6 6-6" />
              </svg>
            </h2>
          </summary>
          <section className="pb-4">
            {item.content}
          </section>
        </details>
      ))}
    </section>
  );
}
```

### Vue Integration
```vue
<template>
  <section class="accordion">
    <details 
      v-for="(item, index) in items" 
      :key="index"
      class="group border-b last:border-b-0"
      :open="openItems.includes(index)"
      @toggle="toggleItem(index, $event)"
    >
      <summary class="w-full outline-none cursor-pointer">
        <h2 class="flex justify-between py-4 text-sm font-medium">
          {{ item.title }}
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
            <path d="m6 9 6 6 6-6" />
          </svg>
        </h2>
      </summary>
      <section class="pb-4">
        {{ item.content }}
      </section>
    </details>
  </section>
</template>

<script>
export default {
  props: {
    items: Array,
    allowMultiple: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      openItems: []
    };
  },
  methods: {
    toggleItem(index, event) {
      if (!this.allowMultiple) {
        this.openItems = event.target.open ? [index] : [];
      }
    }
  }
};
</script>
```

### HTMX Integration
```html
<section class="accordion">
  <details class="group border-b" hx-get="/api/content/1" hx-trigger="toggle once">
    <summary class="w-full outline-none">
      <h2 class="flex justify-between py-4 text-sm font-medium">
        Load Content on Expand
        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="transition-transform group-open:rotate-180">
          <path d="m6 9 6 6 6-6" />
        </svg>
      </h2>
    </summary>
    <section class="pb-4" hx-target="this">
      <div class="skeleton">Loading...</div>
    </section>
  </details>
</section>
```

## Related Components

- [Card](./card.md) - For grouped content without expand/collapse
- [Tabs](./tabs.md) - For switching between content views
- [Details](./details.md) - Native HTML details element
- [Collapse](./collapse.md) - For simple show/hide functionality