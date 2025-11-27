# Toast Component

A succinct message that is displayed temporarily to provide feedback to users about actions they have taken.

## Basic Usage

### Setup
First, add the toaster container to your HTML (typically at the end of the `<body>`):

```html
<div id="toaster" class="toaster"></div>
```

### Trigger from Frontend
```html
<button
  class="btn-outline"
  onclick="document.dispatchEvent(new CustomEvent('basecoat:toast', {
    detail: {
      config: {
        category: 'success',
        title: 'Success',
        description: 'A success toast called from the front-end.',
        cancel: {
          label: 'Dismiss'
        }
      }
    }
  }))"
>
  Toast from front-end
</button>
```

### Trigger from Backend (HTMX)
```html
<button
  class="btn-outline"
  hx-trigger="click"
  hx-get="/fragments/toast/success"
  hx-target="#toaster"
  hx-swap="beforeend"
>
  Toast from backend (with HTMX)
</button>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/toast.min.js" defer></script>
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
- **`toaster`** - Applied to the container that holds all toasts
- **`toast`** - Applied to individual toast elements

### Supporting Classes
- **`toast-content`** - Applied to the content wrapper inside each toast
- Button classes for actions (`btn`, `btn-outline`)
- Category-specific styling is applied automatically

### Tailwind Utilities Used
- Positioning and layout classes for the toaster
- Spacing utilities for content arrangement
- Animation classes for enter/exit transitions

## Component Attributes

### Toaster Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "toaster" | Yes |
| `id` | string | Unique identifier (commonly "toaster") | Recommended |

### Toast Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "toast" | Yes |
| `data-duration` | number | Duration in milliseconds | No |

## HTML Structure

### Toaster Container
```html
<div id="toaster" class="toaster">
  <!-- Toasts are dynamically inserted here -->
</div>
```

### Individual Toast Structure
```html
<div class="toast" data-duration="5000">
  <div class="toast-content">
    <!-- Optional icon -->
    <svg aria-hidden="true">
      <!-- Category icon -->
    </svg>
    
    <!-- Message section -->
    <section>
      <h2>Toast Title</h2>
      <p>Toast description (optional)</p>
    </section>
    
    <!-- Optional footer with buttons -->
    <footer>
      <button type="button" class="btn" onclick="handleAction()">Action</button>
      <button type="button" class="btn-outline" onclick="handleCancel()">Cancel</button>
    </footer>
  </div>
</div>
```

## Toast Configuration

### JavaScript Config Object
```javascript
const toastConfig = {
  duration: 5000,                    // Duration in milliseconds (optional)
  category: 'success',               // 'success', 'info', 'warning', 'error' (optional)
  title: 'Operation Successful',     // Toast title (required)
  description: 'Your data has been saved.', // Toast description (optional)
  
  // Action button (optional)
  action: {
    label: 'View Details',
    onclick: 'showDetails()'
  },
  
  // Cancel/dismiss button (optional)
  cancel: {
    label: 'Dismiss',               // Defaults to "Dismiss" if not provided
    onclick: 'handleCancel()'       // Optional custom handler
  }
};

// Trigger the toast
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: { config: toastConfig }
}));
```

### Configuration Properties

| Property | Type | Description | Default |
|----------|------|-------------|---------|
| `duration` | number | Duration in milliseconds | 3000ms (3s) or 5000ms (5s) for errors |
| `category` | string | Toast type: 'success', 'info', 'warning', 'error' | undefined |
| `title` | string | Toast title | Required |
| `description` | string | Toast description | undefined |
| `action` | object | Action button configuration | undefined |
| `cancel` | object | Cancel button configuration | undefined |

## Examples

### Basic Success Toast
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    config: {
      category: 'success',
      title: 'Settings Saved',
      description: 'Your preferences have been updated successfully.'
    }
  }
}));
```

### Error Toast with Custom Duration
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    config: {
      category: 'error',
      title: 'Upload Failed',
      description: 'The file could not be uploaded. Please try again.',
      duration: 10000  // 10 seconds
    }
  }
}));
```

### Info Toast with Action
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    config: {
      category: 'info',
      title: 'New Update Available',
      description: 'Version 2.0 is now available.',
      action: {
        label: 'Update Now',
        onclick: 'startUpdate()'
      },
      cancel: {
        label: 'Later'
      }
    }
  }
}));
```

### Warning Toast with Multiple Actions
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    config: {
      category: 'warning',
      title: 'Unsaved Changes',
      description: 'You have unsaved changes that will be lost.',
      action: {
        label: 'Save',
        onclick: 'saveChanges()'
      },
      cancel: {
        label: 'Discard',
        onclick: 'discardChanges()'
      }
    }
  }
}));
```

### Toast with Link Action
```javascript
document.dispatchEvent(new CustomEvent('basecoat:toast', {
  detail: {
    config: {
      category: 'success',
      title: 'Account Created',
      description: 'Welcome! Your account has been created successfully.',
      action: {
        label: 'Get Started',
        href: '/dashboard'  // Use href instead of onclick for links
      }
    }
  }
}));
```

## JavaScript Events

### Component Events
```javascript
const toaster = document.getElementById('toaster');

// Listen for toaster initialization
toaster.addEventListener('basecoat:initialized', (e) => {
  console.log('Toaster initialized');
});

// Listen for toast events
document.addEventListener('basecoat:toast', (e) => {
  console.log('Toast triggered:', e.detail.config);
});
```

### Manual Toast Creation
```javascript
function createToast(config) {
  document.dispatchEvent(new CustomEvent('basecoat:toast', {
    detail: { config }
  }));
}

// Usage
createToast({
  category: 'success',
  title: 'Task Complete',
  description: 'Your task has been completed successfully.'
});
```

### Toast Helper Functions
```javascript
// Helper functions for common toast types
const Toast = {
  success(title, description = '') {
    createToast({ category: 'success', title, description });
  },
  
  error(title, description = '') {
    createToast({ 
      category: 'error', 
      title, 
      description,
      duration: 5000 
    });
  },
  
  info(title, description = '') {
    createToast({ category: 'info', title, description });
  },
  
  warning(title, description = '') {
    createToast({ category: 'warning', title, description });
  },
  
  confirm(title, description, onConfirm, onCancel) {
    createToast({
      category: 'warning',
      title,
      description,
      action: {
        label: 'Confirm',
        onclick: onConfirm
      },
      cancel: {
        label: 'Cancel',
        onclick: onCancel
      }
    });
  }
};

// Usage
Toast.success('Data Saved', 'Your changes have been saved successfully.');
Toast.error('Connection Failed', 'Unable to connect to the server.');
Toast.confirm(
  'Delete Item', 
  'Are you sure you want to delete this item?',
  'deleteItem()',
  'cancelDelete()'
);
```

## HTMX Integration

### Backend-Generated Toasts
Your server can return toast HTML that gets inserted into the toaster:

```html
<!-- Server response (e.g., from /fragments/toast/success) -->
<div class="toast" data-duration="4000">
  <div class="toast-content">
    <svg aria-hidden="true">
      <!-- Success icon -->
      <path d="M20 6 9 17l-5-5" />
    </svg>
    <section>
      <h2>Data Saved</h2>
      <p>Your changes have been saved successfully.</p>
    </section>
    <footer>
      <button type="button" class="btn-outline">Dismiss</button>
    </footer>
  </div>
</div>
```

### HTMX Toast Triggers
```html
<!-- Form submission with toast feedback -->
<form hx-post="/api/save" hx-target="#toaster" hx-swap="beforeend">
  <input type="text" name="data" required>
  <button type="submit" class="btn">Save</button>
</form>

<!-- Button with confirmation toast -->
<button 
  hx-delete="/api/item/123"
  hx-target="#toaster"
  hx-swap="beforeend"
  hx-confirm="Are you sure you want to delete this item?"
  class="btn-destructive"
>
  Delete
</button>

<!-- Auto-refresh with toast notification -->
<div 
  hx-get="/api/status" 
  hx-target="#status"
  hx-trigger="every 30s"
  hx-on::after-request="showStatusToast()"
>
  Status: <span id="status">Checking...</span>
</div>
```

## Accessibility Features

- **Screen Reader Support**: Toasts are announced when they appear
- **Keyboard Navigation**: Action buttons are keyboard accessible
- **Focus Management**: Focus is maintained appropriately
- **Role Semantics**: Proper ARIA roles for alerting users

### Enhanced Accessibility
```javascript
// Create accessible toast with proper ARIA attributes
function createAccessibleToast(config) {
  const toastConfig = {
    ...config,
    // Add ARIA attributes for better accessibility
    'aria-live': config.category === 'error' ? 'assertive' : 'polite',
    'aria-atomic': 'true',
    role: 'alert'
  };
  
  document.dispatchEvent(new CustomEvent('basecoat:toast', {
    detail: { config: toastConfig }
  }));
}
```

## Styling and Theming

### Custom Toast Duration
```html
<!-- Toast with custom 10-second duration -->
<div class="toast" data-duration="10000">
  <!-- Toast content -->
</div>
```

### Category-Based Styling
The toast component automatically applies category-specific styling:

```css
/* Automatic category styles applied by the component */
.toast[data-category="success"] {
  /* Success styling */
}

.toast[data-category="error"] {
  /* Error styling */
}

.toast[data-category="warning"] {
  /* Warning styling */
}

.toast[data-category="info"] {
  /* Info styling */
}
```

## Best Practices

1. **Appropriate Duration**: Use longer durations for error messages (5s) and shorter for success (3s)
2. **Clear Messaging**: Keep titles concise and descriptions helpful
3. **Meaningful Actions**: Provide relevant actions when needed
4. **Category Usage**: Use appropriate categories to convey the right tone
5. **Avoid Spam**: Don't overwhelm users with too many toasts
6. **Important Messages**: Use error category for critical information
7. **Undo Actions**: Provide undo functionality for destructive actions
8. **Accessibility**: Ensure toast content is screen reader friendly

## Common Patterns

### Undo Pattern
```javascript
function deleteWithUndo(itemId) {
  // Perform the deletion
  deleteItem(itemId);
  
  // Show toast with undo option
  document.dispatchEvent(new CustomEvent('basecoat:toast', {
    detail: {
      config: {
        category: 'info',
        title: 'Item Deleted',
        description: 'The item has been moved to trash.',
        action: {
          label: 'Undo',
          onclick: `restoreItem(${itemId})`
        },
        duration: 8000  // Longer duration for undo actions
      }
    }
  }));
}
```

### Progress Notification Pattern
```javascript
function showProgress() {
  let progress = 0;
  const interval = setInterval(() => {
    progress += 20;
    
    if (progress <= 100) {
      document.dispatchEvent(new CustomEvent('basecoat:toast', {
        detail: {
          config: {
            category: 'info',
            title: 'Processing...',
            description: `Progress: ${progress}%`,
            duration: 1000
          }
        }
      }));
    } else {
      clearInterval(interval);
      document.dispatchEvent(new CustomEvent('basecoat:toast', {
        detail: {
          config: {
            category: 'success',
            title: 'Complete',
            description: 'Processing finished successfully.'
          }
        }
      }));
    }
  }, 1000);
}
```

### Form Validation Pattern
```javascript
function validateAndSave(formData) {
  const errors = validateForm(formData);
  
  if (errors.length > 0) {
    // Show validation errors
    document.dispatchEvent(new CustomEvent('basecoat:toast', {
      detail: {
        config: {
          category: 'error',
          title: 'Validation Failed',
          description: `Please fix ${errors.length} error(s) and try again.`,
          duration: 5000
        }
      }
    }));
  } else {
    // Save and show success
    saveForm(formData).then(() => {
      document.dispatchEvent(new CustomEvent('basecoat:toast', {
        detail: {
          config: {
            category: 'success',
            title: 'Saved Successfully',
            description: 'Your changes have been saved.'
          }
        }
      }));
    });
  }
}
```

## Integration Examples

### React Integration
```jsx
import React, { useEffect } from 'react';

function ToastProvider({ children }) {
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  const showToast = (config) => {
    document.dispatchEvent(new CustomEvent('basecoat:toast', {
      detail: { config }
    }));
  };

  return (
    <>
      {children}
      <div id="toaster" className="toaster" />
    </>
  );
}

// Hook for using toasts
function useToast() {
  return {
    success: (title, description) => showToast({ category: 'success', title, description }),
    error: (title, description) => showToast({ category: 'error', title, description }),
    info: (title, description) => showToast({ category: 'info', title, description }),
    warning: (title, description) => showToast({ category: 'warning', title, description })
  };
}
```

### Vue Integration
```vue
<template>
  <div id="app">
    <slot />
    <div id="toaster" class="toaster" />
  </div>
</template>

<script>
export default {
  mounted() {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  },
  
  methods: {
    $toast(config) {
      document.dispatchEvent(new CustomEvent('basecoat:toast', {
        detail: { config }
      }));
    }
  }
}
</script>
```

## Jinja/Nunjucks Macros

### Toaster Setup
```jinja2
{% from "toast.njk" import toaster %}
{{ toaster(
  toasts=[
    {
      type: "success",
      title: "Success",
      description: "A success toast called from the front-end.",
      action: { label: "Dismiss", click: "close()" }
    },
    {
      type: "info",
      title: "Info", 
      description: "An info toast called from the front-end.",
      action: { label: "Dismiss", click: "close()" }
    }
  ]
) }}
```

### Individual Toast
```jinja2
{% from "toast.njk" import toast %}
{{ toast(
  title="Event has been created",
  description="Sunday, December 03, 2023 at 9:00 AM",
  cancel={ label: "Undo" }
) }}
```

### Dynamic Toast Generation
```jinja2
{% for message in flash_messages %}
  {{ toast(
    category=message.category,
    title=message.title,
    description=message.text,
    cancel={ label: "Dismiss" }
  ) }}
{% endfor %}
```

## Related Components

- [Button](./button.md) - For toast actions and triggers
- [Dialog](./dialog.md) - For more complex user interactions
- [Alert](./alert.md) - For persistent status messages
- [Notification](./notification.md) - For system-level notifications