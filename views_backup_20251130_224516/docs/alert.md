# Alert Component

Displays a callout for user attention.

## Basic Usage

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  <h2>Success! Your changes have been saved</h2>
  <section>This is an alert with icon, title and description.</section>
</div>
```

## CSS Classes

### Primary Classes
- **`alert`** - Default alert styling for informational/success messages
- **`alert-destructive`** - Destructive/error alert styling

### Supporting Classes
- **Typography**: `h2` for titles, `section` for descriptions
- **Icons**: Standard SVG icons (24x24px recommended)
- **Lists**: `list-inside`, `list-disc`, `text-sm` for content formatting

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "alert" or "alert-destructive" | Yes |

### Icon Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| Standard SVG attributes | various | Width, height, viewBox, stroke properties | Optional |

### No JavaScript Required
Basic alerts work with pure CSS and HTML.

## HTML Structure

```html
<!-- Full alert structure -->
<div class="alert">
  <svg><!-- Icon (optional) --></svg>
  <h2>Title</h2>
  <section>Description (optional)</section>
</div>

<!-- Alert without icon -->
<div class="alert">
  <h2>Title</h2>
  <section>Description</section>
</div>

<!-- Alert without description -->
<div class="alert">
  <svg><!-- Icon --></svg>
  <h2>Title</h2>
</div>
```

## Examples

### Default Alert

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  <h2>Success! Your changes have been saved</h2>
  <section>This is an alert with icon, title and description.</section>
</div>
```

### Destructive Alert

```html
<div class="alert-destructive">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <line x1="12" x2="12" y1="8" y2="12" />
    <line x1="12" x2="12.01" y1="16" y2="16" />
  </svg>
  <h2>Something went wrong!</h2>
  <section>Your session has expired. Please log in again.</section>
</div>
```

### Alert with Complex Description

```html
<div class="alert-destructive">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  <h2>Unable to process your payment.</h2>
  <section>
    <p>Please verify your billing information and try again.</p>
    <ul class="list-inside list-disc text-sm">
      <li>Check your card details</li>
      <li>Ensure sufficient funds</li>
      <li>Verify billing address</li>
    </ul>
  </section>
</div>
```

### Alert Without Icon

```html
<div class="alert">
  <h2>Success! Your changes have been saved</h2>
  <section>This is an alert with title and description but no icon.</section>
</div>
```

### Alert Without Description

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z" />
    <path d="M12 8v4" />
    <path d="M12 16h.01" />
  </svg>
  <h2>This alert only has an icon and title</h2>
</div>
```

### Alert Title Only

```html
<div class="alert">
  <h2>Simple alert with just a title</h2>
</div>
```

### Alert with Long Title

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z" />
    <path d="M12 8v4" />
    <path d="M12 16h.01" />
  </svg>
  <h2>This is a very long alert title that demonstrates how the component handles extended text content and potentially wraps across multiple lines</h2>
</div>
```

### Common Alert Icons

```html
<!-- Success/Check -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <circle cx="12" cy="12" r="10" />
  <path d="m9 12 2 2 4-4" />
</svg>

<!-- Error/Warning -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <circle cx="12" cy="12" r="10" />
  <line x1="12" x2="12" y1="8" y2="12" />
  <line x1="12" x2="12.01" y1="16" y2="16" />
</svg>

<!-- Info -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <circle cx="12" cy="12" r="10" />
  <path d="M12 16v-4" />
  <path d="M12 8h.01" />
</svg>

<!-- Warning/Alert Triangle -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="m21.73 18-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z" />
  <path d="M12 9v4" />
  <path d="M12 17h.01" />
</svg>

<!-- Shield (Security) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z" />
  <path d="M12 8v4" />
  <path d="M12 16h.01" />
</svg>
```

### Multiple Alerts Stack

```html
<div class="grid w-full max-w-xl items-start gap-4">
  <!-- Success Alert -->
  <div class="alert">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="m9 12 2 2 4-4" />
    </svg>
    <h2>Success! Your changes have been saved</h2>
    <section>This is an alert with icon, title and description.</section>
  </div>

  <!-- Warning Alert -->
  <div class="alert">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M18 8a2 2 0 0 0 0-4 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0-4 0 2 2 0 0 0 0 4" />
      <path d="M10 22 9 8" />
      <path d="m14 22 1-14" />
      <path d="M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z" />
    </svg>
    <h2>This Alert has a title and an icon. No description.</h2>
    <section>This is an alert with icon, title and description.</section>
  </div>

  <!-- Error Alert -->
  <div class="alert-destructive">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="m9 12 2 2 4-4" />
    </svg>
    <h2>Unable to process your payment.</h2>
    <section>
      <p>Please verify your billing information and try again.</p>
      <ul class="list-inside list-disc text-sm">
        <li>Check your card details</li>
        <li>Ensure sufficient funds</li>
        <li>Verify billing address</li>
      </ul>
    </section>
  </div>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper heading tags (`h2`) for titles
- **Color Contrast**: High contrast colors for readability
- **Icon Support**: Icons are decorative and don't require alt text
- **Screen Reader Support**: Proper heading hierarchy and content structure

### Enhanced Accessibility

```html
<!-- Alert with ARIA role -->
<div class="alert" role="alert" aria-live="polite">
  <h2>System Notification</h2>
  <section>Your changes have been automatically saved.</section>
</div>

<!-- Alert with specific ARIA attributes -->
<div class="alert-destructive" role="alert" aria-live="assertive" aria-atomic="true">
  <h2>Critical Error</h2>
  <section>Immediate action required to prevent data loss.</section>
</div>

<!-- Alert with additional description -->
<div class="alert" aria-labelledby="alert-title" aria-describedby="alert-description">
  <h2 id="alert-title">Upload Complete</h2>
  <section id="alert-description">Your file has been successfully uploaded and is now available in your dashboard.</section>
</div>
```

## JavaScript Integration

### Dynamic Alert Creation

```javascript
// Create alert programmatically
function createAlert(type, title, description, icon) {
  const alert = document.createElement('div');
  alert.className = type === 'error' ? 'alert-destructive' : 'alert';
  
  let html = '';
  
  // Add icon if provided
  if (icon) {
    html += icon;
  }
  
  // Add title
  html += `<h2>${title}</h2>`;
  
  // Add description if provided
  if (description) {
    html += `<section>${description}</section>`;
  }
  
  alert.innerHTML = html;
  return alert;
}

// Usage examples
const successAlert = createAlert(
  'success', 
  'Changes Saved', 
  'Your profile has been updated successfully.',
  '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10" /><path d="m9 12 2 2 4-4" /></svg>'
);

const errorAlert = createAlert(
  'error',
  'Upload Failed',
  'Please check your internet connection and try again.',
  '<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10" /><line x1="12" x2="12" y1="8" y2="12" /><line x1="12" x2="12.01" y1="16" y2="16" /></svg>'
);

// Add to page
document.body.appendChild(successAlert);
```

### Alert with Dismiss Button

```html
<div class="alert" id="dismissible-alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  <h2>Success! Your changes have been saved</h2>
  <section>You can dismiss this alert by clicking the X button.</section>
  <button onclick="this.parentElement.remove()" class="btn-icon-ghost ml-auto" aria-label="Dismiss alert">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  </button>
</div>
```

### Temporary Alert

```javascript
// Show temporary alert that auto-dismisses
function showTemporaryAlert(type, title, description, duration = 5000) {
  const alertContainer = document.getElementById('alert-container') || document.body;
  
  const alert = createAlert(type, title, description);
  alertContainer.appendChild(alert);
  
  // Auto-dismiss after specified duration
  setTimeout(() => {
    alert.style.transition = 'opacity 0.3s ease-out';
    alert.style.opacity = '0';
    
    setTimeout(() => {
      alert.remove();
    }, 300);
  }, duration);
}

// Usage
showTemporaryAlert('success', 'Saved!', 'Your changes have been saved.', 3000);
```

### React Integration

```jsx
import React, { useState } from 'react';

function Alert({ type = 'info', title, children, icon, onDismiss, className = '' }) {
  const [isVisible, setIsVisible] = useState(true);
  
  const handleDismiss = () => {
    setIsVisible(false);
    onDismiss?.();
  };
  
  if (!isVisible) return null;
  
  const alertClasses = type === 'error' ? 'alert-destructive' : 'alert';
  
  return (
    <div className={`${alertClasses} ${className}`} role="alert">
      {icon && <div dangerouslySetInnerHTML={{ __html: icon }} />}
      <h2>{title}</h2>
      {children && <section>{children}</section>}
      {onDismiss && (
        <button onClick={handleDismiss} className="btn-icon-ghost ml-auto" aria-label="Dismiss alert">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M18 6 6 18" />
            <path d="m6 6 12 12" />
          </svg>
        </button>
      )}
    </div>
  );
}

// Usage
function App() {
  return (
    <div className="space-y-4">
      <Alert type="success" title="Success!" onDismiss={() => console.log('Alert dismissed')}>
        Your changes have been saved successfully.
      </Alert>
      
      <Alert 
        type="error" 
        title="Error occurred" 
        icon='<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10" /><line x1="12" x2="12" y1="8" y2="12" /><line x1="12" x2="12.01" y1="16" y2="16" /></svg>'
      >
        Please check your internet connection and try again.
      </Alert>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div v-if="visible" :class="alertClasses" role="alert">
    <div v-if="icon" v-html="icon"></div>
    <h2>{{ title }}</h2>
    <section v-if="$slots.default"><slot /></section>
    <button 
      v-if="dismissible" 
      @click="dismiss"
      class="btn-icon-ghost ml-auto"
      aria-label="Dismiss alert"
    >
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</template>

<script>
export default {
  props: {
    type: {
      type: String,
      default: 'info',
      validator: value => ['info', 'success', 'error'].includes(value)
    },
    title: {
      type: String,
      required: true
    },
    icon: String,
    dismissible: Boolean
  },
  data() {
    return {
      visible: true
    };
  },
  computed: {
    alertClasses() {
      return this.type === 'error' ? 'alert-destructive' : 'alert';
    }
  },
  methods: {
    dismiss() {
      this.visible = false;
      this.$emit('dismiss');
    }
  }
};
</script>
```

## Best Practices

1. **Use Appropriate Types**: Use `alert` for general info/success, `alert-destructive` for errors
2. **Clear Titles**: Keep titles concise and descriptive
3. **Meaningful Icons**: Use icons that match the alert type and message
4. **Structured Content**: Use lists for multiple action items in descriptions
5. **Color Independence**: Don't rely solely on color to convey meaning
6. **Consistent Sizing**: Use 24x24px icons for visual consistency
7. **ARIA Support**: Include proper ARIA attributes for dynamic alerts

## Common Patterns

### Form Validation Alerts

```html
<div class="alert-destructive">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <line x1="12" x2="12" y1="8" y2="12" />
    <line x1="12" x2="12.01" y1="16" y2="16" />
  </svg>
  <h2>Form submission failed</h2>
  <section>
    <p>Please correct the following errors:</p>
    <ul class="list-inside list-disc text-sm mt-2">
      <li>Email address is required</li>
      <li>Password must be at least 8 characters</li>
      <li>Please accept the terms of service</li>
    </ul>
  </section>
</div>
```

### Success Confirmation

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  <h2>Account created successfully</h2>
  <section>Welcome! Please check your email to verify your account.</section>
</div>
```

### System Maintenance

```html
<div class="alert">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <path d="M12 16v-4" />
    <path d="M12 8h.01" />
  </svg>
  <h2>Scheduled maintenance</h2>
  <section>Our service will be temporarily unavailable on Sunday, December 15th from 2:00 AM to 4:00 AM EST for scheduled maintenance.</section>
</div>
```

## Related Components

- [Alert Dialog](./alert-dialog.md) - For interactive confirmation alerts
- [Toast](./toast.md) - For temporary notifications
- [Button](./button.md) - For adding action buttons to alerts
- [Badge](./badge.md) - For status indicators within alerts