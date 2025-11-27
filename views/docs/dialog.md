# Dialog Component

A window overlaid on either the primary window or another dialog window, rendering the content underneath inert.

## Basic Usage

```html
<button type="button" onclick="document.getElementById('demo-dialog').showModal()" class="btn-outline">
  Open Dialog
</button>

<dialog id="demo-dialog" class="dialog w-full sm:max-w-[425px] max-h-[612px]" aria-labelledby="demo-dialog-title" aria-describedby="demo-dialog-description" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="demo-dialog-title">Dialog Title</h2>
      <p id="demo-dialog-description">Dialog description here.</p>
    </header>

    <section>
      <!-- Dialog content -->
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Cancel</button>
      <button class="btn" onclick="this.closest('dialog').close()">Confirm</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

## CSS Classes

### Primary Classes
- **`dialog`** - Applied to the native `<dialog>` element

### Supporting Classes
- **`scrollbar`** - For scrollable content sections
- Button classes (`btn`, `btn-outline`) for actions
- Form classes for form content

### Tailwind Utilities Used
- `w-full` - Full width
- `sm:max-w-[425px]` - Responsive max width
- `max-h-[612px]` - Maximum height
- `overflow-y-auto` - Vertical scrolling
- `space-y-4` - Vertical spacing
- `grid gap-3` - Grid layout for forms
- Various layout and spacing utilities

## Component Attributes

### Dialog Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "dialog" class | Yes |
| `id` | string | Unique identifier | Yes |
| `aria-labelledby` | string | References title element ID | Recommended |
| `aria-describedby` | string | References description element ID | Recommended |
| `onclick` | string | Backdrop click handler | Recommended |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `onclick` | string | Should call showModal() | Yes |

## HTML Structure

```html
<!-- Trigger button -->
<button type="button" onclick="document.getElementById('dialog-id').showModal()">
  Trigger Text
</button>

<!-- Dialog element -->
<dialog id="dialog-id" class="dialog" aria-labelledby="title-id" aria-describedby="desc-id" onclick="if (event.target === this) this.close()">
  <div>
    <!-- Header section -->
    <header>
      <h2 id="title-id">Title</h2>
      <p id="desc-id">Description</p> <!-- Optional -->
    </header>
    
    <!-- Content section -->
    <section>
      <!-- Dialog content -->
    </section>
    
    <!-- Footer section -->
    <footer>
      <!-- Action buttons -->
    </footer>
    
    <!-- Close button -->
    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <!-- Close icon -->
    </button>
  </div>
</dialog>
```

## Examples

### Basic Dialog
```html
<button type="button" onclick="document.getElementById('basic-dialog').showModal()" class="btn-outline">
  Open Basic Dialog
</button>

<dialog id="basic-dialog" class="dialog w-full sm:max-w-[425px]" aria-labelledby="basic-dialog-title" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="basic-dialog-title">Confirm Action</h2>
    </header>

    <section>
      <p>Are you sure you want to perform this action? This cannot be undone.</p>
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Cancel</button>
      <button class="btn-destructive" onclick="this.closest('dialog').close()">Delete</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

### Form Dialog
```html
<button type="button" onclick="document.getElementById('form-dialog').showModal()" class="btn-outline">
  Edit Profile
</button>

<dialog id="form-dialog" class="dialog w-full sm:max-w-[425px]" aria-labelledby="form-dialog-title" aria-describedby="form-dialog-description" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="form-dialog-title">Edit Profile</h2>
      <p id="form-dialog-description">Make changes to your profile here. Click save when you're done.</p>
    </header>

    <section>
      <form class="form grid gap-4">
        <div class="grid gap-3">
          <label for="profile-name">Name</label>
          <input type="text" id="profile-name" value="John Doe" autofocus />
        </div>
        <div class="grid gap-3">
          <label for="profile-email">Email</label>
          <input type="email" id="profile-email" value="john@example.com" />
        </div>
        <div class="grid gap-3">
          <label for="profile-bio">Bio</label>
          <textarea id="profile-bio" rows="3" placeholder="Tell us about yourself"></textarea>
        </div>
      </form>
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Cancel</button>
      <button class="btn" onclick="this.closest('dialog').close()">Save Changes</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

### Scrollable Content Dialog
```html
<button type="button" onclick="document.getElementById('scrollable-dialog').showModal()" class="btn-outline">
  Long Content
</button>

<dialog id="scrollable-dialog" class="dialog w-full sm:max-w-[425px] max-h-[612px]" aria-labelledby="scrollable-dialog-title" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="scrollable-dialog-title">Terms and Conditions</h2>
    </header>

    <section class="overflow-y-auto scrollbar">
      <div class="space-y-4 text-sm">
        <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua...</p>
        <p>Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat...</p>
        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur...</p>
        <!-- More content -->
      </div>
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Decline</button>
      <button class="btn" onclick="this.closest('dialog').close()">Accept</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

### Confirmation Dialog
```html
<button type="button" onclick="document.getElementById('confirmation-dialog').showModal()" class="btn-destructive">
  Delete Account
</button>

<dialog id="confirmation-dialog" class="dialog w-full sm:max-w-[400px]" aria-labelledby="confirmation-dialog-title" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="confirmation-dialog-title">Delete Account</h2>
    </header>

    <section>
      <div class="space-y-3">
        <p>Are you absolutely sure you want to delete your account?</p>
        <div class="rounded-lg border border-destructive/20 bg-destructive/10 p-3">
          <p class="text-sm text-destructive">
            <strong>Warning:</strong> This action cannot be undone. This will permanently delete your account and remove all associated data.
          </p>
        </div>
      </div>
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Cancel</button>
      <button class="btn-destructive" onclick="this.closest('dialog').close()">Yes, Delete Account</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

### Dialog without Footer
```html
<button type="button" onclick="document.getElementById('simple-dialog').showModal()" class="btn-outline">
  Show Message
</button>

<dialog id="simple-dialog" class="dialog w-full sm:max-w-[350px]" aria-labelledby="simple-dialog-title" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <h2 id="simple-dialog-title">Success!</h2>
    </header>

    <section>
      <p>Your changes have been saved successfully.</p>
    </section>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

### Alert Dialog
```html
<button type="button" onclick="document.getElementById('alert-dialog').showModal()" class="btn-outline">
  Show Alert
</button>

<dialog id="alert-dialog" class="dialog w-full sm:max-w-[400px]" aria-labelledby="alert-dialog-title" onclick="if (event.target === this) this.close()">
  <div>
    <header>
      <div class="flex items-center gap-3">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-warning">
          <path d="m21 16-4 4-4-4"/>
          <path d="M17 20V9.5a2.5 2.5 0 0 0-5 0V20"/>
          <path d="M12 13.5h5"/>
          <circle cx="5.5" cy="11.5" r="4.5"/>
        </svg>
        <h2 id="alert-dialog-title" class="text-lg font-semibold">Warning</h2>
      </div>
    </header>

    <section>
      <p>Your session is about to expire. Would you like to continue?</p>
    </section>

    <footer>
      <button class="btn-outline" onclick="this.closest('dialog').close()">Log Out</button>
      <button class="btn" onclick="this.closest('dialog').close()">Stay Signed In</button>
    </footer>

    <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M18 6 6 18" />
        <path d="m6 6 12 12" />
      </svg>
    </button>
  </div>
</dialog>
```

## JavaScript Methods

### Opening Dialog
```javascript
// Show modal dialog (blocks interaction with page)
document.getElementById('dialog-id').showModal();

// Show non-modal dialog (allows page interaction)
document.getElementById('dialog-id').show();
```

### Closing Dialog
```javascript
// Close dialog
document.getElementById('dialog-id').close();

// Close with return value
document.getElementById('dialog-id').close('confirmed');
```

### Dialog Events
```javascript
const dialog = document.getElementById('my-dialog');

// Listen for dialog close
dialog.addEventListener('close', (e) => {
  console.log('Dialog closed with return value:', dialog.returnValue);
});

// Listen for dialog cancel (ESC key)
dialog.addEventListener('cancel', (e) => {
  console.log('Dialog cancelled');
  // Optionally prevent closing
  // e.preventDefault();
});
```

### Form Integration
```javascript
// Using form method="dialog" to close
function handleFormSubmit(form) {
  const formData = new FormData(form);
  
  // Process form data
  console.log('Form submitted:', Object.fromEntries(formData));
  
  // Close dialog
  form.closest('dialog').close('submitted');
}
```

```html
<!-- Alternative form approach -->
<form method="dialog">
  <input type="text" name="username" required>
  <button type="submit" value="save">Save</button>
  <button type="submit" value="cancel" formnovalidate>Cancel</button>
</form>
```

## Event Handling

### Click Outside to Close
```javascript
// Add to dialog element
function handleBackdropClick(event) {
  if (event.target === this) {
    this.close();
  }
}

// Inline version (already shown in examples)
onclick="if (event.target === this) this.close()"
```

### Keyboard Handling
```javascript
document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape') {
    // ESC key automatically closes modal dialogs
    // Custom handling if needed
  }
});
```

### Dynamic Dialog Creation
```javascript
function createDialog(title, content, actions = []) {
  const dialog = document.createElement('dialog');
  dialog.className = 'dialog w-full sm:max-w-[425px]';
  dialog.setAttribute('aria-labelledby', 'dynamic-title');
  dialog.onclick = function(e) { if (e.target === this) this.close(); };
  
  const actionsHTML = actions.map(action => 
    `<button class="${action.class}" onclick="this.closest('dialog').close('${action.value}')">${action.label}</button>`
  ).join('');
  
  dialog.innerHTML = `
    <div>
      <header>
        <h2 id="dynamic-title">${title}</h2>
      </header>
      <section>
        ${content}
      </section>
      ${actions.length ? `<footer>${actionsHTML}</footer>` : ''}
      <button type="button" aria-label="Close dialog" onclick="this.closest('dialog').close()">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M18 6 6 18" />
          <path d="m6 6 12 12" />
        </svg>
      </button>
    </div>
  `;
  
  document.body.appendChild(dialog);
  return dialog;
}

// Usage
const confirmDialog = createDialog(
  'Confirm Delete',
  '<p>Are you sure you want to delete this item?</p>',
  [
    { label: 'Cancel', class: 'btn-outline', value: 'cancel' },
    { label: 'Delete', class: 'btn-destructive', value: 'delete' }
  ]
);

confirmDialog.showModal();
confirmDialog.addEventListener('close', (e) => {
  if (confirmDialog.returnValue === 'delete') {
    // Handle deletion
  }
  document.body.removeChild(confirmDialog);
});
```

## Accessibility Features

- **Modal Behavior**: Properly blocks interaction with underlying content
- **Keyboard Navigation**: ESC key closes, focus management
- **Screen Reader Support**: Proper labeling with `aria-labelledby` and `aria-describedby`
- **Focus Management**: Auto-focus first interactive element
- **Role Semantics**: Native dialog element provides proper semantics

### Enhanced Accessibility
```html
<dialog id="accessible-dialog" class="dialog" role="dialog" aria-modal="true" aria-labelledby="accessible-title" aria-describedby="accessible-desc">
  <div>
    <header>
      <h2 id="accessible-title">Dialog Title</h2>
      <p id="accessible-desc">Dialog description for screen readers</p>
    </header>
    
    <section>
      <label for="first-input">First Input</label>
      <input id="first-input" type="text" autofocus>
    </section>
    
    <footer>
      <button type="button" onclick="this.closest('dialog').close()">Cancel</button>
      <button type="button" onclick="this.closest('dialog').close()" class="btn">Save</button>
    </footer>
  </div>
</dialog>
```

## Best Practices

1. **Use Modal for Important Actions**: Use `showModal()` for critical interactions
2. **Provide Clear Actions**: Always include clear cancel/confirm options
3. **Keyboard Support**: Ensure ESC key closes and tab navigation works
4. **Focus Management**: Set initial focus appropriately
5. **Backdrop Clicks**: Allow clicking outside to close for non-critical dialogs
6. **Content Length**: Use scrolling for long content, set max-height
7. **Responsive Design**: Use responsive width classes
8. **Form Integration**: Use proper form handling for data collection

## Common Patterns

### Confirmation Pattern
```javascript
function confirmAction(message, onConfirm) {
  const dialog = createDialog(
    'Confirm Action',
    `<p>${message}</p>`,
    [
      { label: 'Cancel', class: 'btn-outline', value: 'cancel' },
      { label: 'Confirm', class: 'btn', value: 'confirm' }
    ]
  );
  
  dialog.showModal();
  dialog.addEventListener('close', () => {
    if (dialog.returnValue === 'confirm') {
      onConfirm();
    }
    document.body.removeChild(dialog);
  });
}

// Usage
confirmAction('Delete this item?', () => {
  // Handle deletion
});
```

### Form Dialog Pattern
```javascript
function openFormDialog(formHTML, onSubmit) {
  const dialog = createDialog('Form', formHTML);
  
  dialog.showModal();
  
  const form = dialog.querySelector('form');
  if (form) {
    form.addEventListener('submit', (e) => {
      e.preventDefault();
      const data = new FormData(form);
      onSubmit(Object.fromEntries(data));
      dialog.close('submitted');
    });
  }
  
  dialog.addEventListener('close', () => {
    document.body.removeChild(dialog);
  });
}
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% set footer %}
  <button class="btn-outline" onclick="this.closest('dialog').close()">Cancel</button>
  <button class="btn" onclick="this.closest('dialog').close()">Save</button>
{% endset %}

{% call dialog(
  id="edit-dialog",
  title="Edit Item",
  description="Make your changes below",
  trigger="Edit",
  trigger_attrs={"class": "btn-outline"},
  dialog_attrs={"class": "w-full sm:max-w-[425px]"},
  footer=footer
) %}
  <form class="form grid gap-4">
    <div class="grid gap-3">
      <label for="item-name">Name</label>
      <input type="text" id="item-name">
    </div>
  </form>
{% endcall %}
```

## Related Components

- [Button](./button.md) - For dialog triggers and actions
- [Form](./form.md) - For form content within dialogs
- [Alert Dialog](./alert-dialog.md) - For alert-specific dialogs
- [Toast](./toast.md) - For non-blocking notifications