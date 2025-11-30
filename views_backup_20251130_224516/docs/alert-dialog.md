# Alert Dialog Component

A modal dialog that interrupts the user with important content and expects a response.

## Basic Usage

```html
<button type="button" onclick="document.getElementById('alert-dialog').showModal()" class="btn-outline">Open alert dialog</button>

<dialog id="alert-dialog" class="dialog" aria-labelledby="alert-dialog-title" aria-describedby="alert-dialog-description">
  <div>
    <header>
      <h2 id="alert-dialog-title">Are you absolutely sure?</h2>
      <p id="alert-dialog-description">This action cannot be undone. This will permanently delete your account and remove your data from our servers.</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('alert-dialog').close()">Cancel</button>
      <button class="btn-primary" onclick="document.getElementById('alert-dialog').close()">Continue</button>
    </footer>
  </div>
</dialog>
```

## CSS Classes

### Primary Classes
- **`dialog`** - Core dialog styling and behavior (same as Dialog component)

### Supporting Classes
- **Layout**: Standard HTML5 dialog structure
- **Typography**: `h2` for titles, `p` for descriptions
- **Buttons**: Standard button classes for actions

### Tailwind Utilities Used
Same as Dialog component - relies on `.dialog` class for styling.

## Component Attributes

### Dialog Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for dialog | Yes |
| `class` | string | Must include "dialog" | Yes |
| `aria-labelledby` | string | References title element ID | Recommended |
| `aria-describedby` | string | References description element ID | Recommended |

### Title Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Must match aria-labelledby value | If using aria-labelledby |

### Description Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Must match aria-describedby value | If using aria-describedby |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `onclick` | string | JavaScript to open dialog | Yes |
| `type` | string | Should be "button" | Yes |

## JavaScript Required

Uses native HTML5 Dialog API:
- `showModal()` to open the dialog
- `close()` to close the dialog

No additional JavaScript frameworks required.

## HTML Structure

```html
<!-- Trigger button (optional) -->
<button type="button" onclick="document.getElementById('dialog-id').showModal()">
  Open Alert Dialog
</button>

<!-- Dialog -->
<dialog id="dialog-id" class="dialog" aria-labelledby="title-id" aria-describedby="description-id">
  <div>
    <!-- Header (required) -->
    <header>
      <h2 id="title-id">Dialog Title</h2>
      <p id="description-id">Dialog description</p>
    </header>
    
    <!-- Content (optional) -->
    <section>
      Additional content here
    </section>
    
    <!-- Footer with actions (recommended) -->
    <footer>
      <button class="btn-outline" onclick="document.getElementById('dialog-id').close()">Cancel</button>
      <button class="btn-destructive" onclick="handleAction()">Confirm</button>
    </footer>
  </div>
</dialog>
```

## Key Differences from Dialog

**Alert Dialog is identical to Dialog except:**
1. **No close button** in the header
2. **No backdrop click to close** - user must explicitly choose an action
3. **Requires explicit user action** to dismiss

## Examples

### Basic Confirmation

```html
<button type="button" onclick="document.getElementById('confirm-delete').showModal()" class="btn-destructive">Delete Item</button>

<dialog id="confirm-delete" class="dialog" aria-labelledby="confirm-delete-title" aria-describedby="confirm-delete-description">
  <div>
    <header>
      <h2 id="confirm-delete-title">Delete Item</h2>
      <p id="confirm-delete-description">Are you sure you want to delete this item? This action cannot be undone.</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('confirm-delete').close()">Cancel</button>
      <button class="btn-destructive" onclick="deleteItem(); document.getElementById('confirm-delete').close()">Delete</button>
    </footer>
  </div>
</dialog>
```

### Account Deletion Warning

```html
<button type="button" onclick="document.getElementById('delete-account').showModal()" class="btn-destructive">Delete Account</button>

<dialog id="delete-account" class="dialog" aria-labelledby="delete-account-title" aria-describedby="delete-account-description">
  <div>
    <header>
      <h2 id="delete-account-title">Are you absolutely sure?</h2>
      <p id="delete-account-description">This action cannot be undone. This will permanently delete your account and remove your data from our servers.</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('delete-account').close()">Cancel</button>
      <button class="btn-destructive" onclick="deleteAccount(); document.getElementById('delete-account').close()">Yes, delete my account</button>
    </footer>
  </div>
</dialog>
```

### Data Loss Warning

```html
<button type="button" onclick="document.getElementById('unsaved-changes').showModal()" class="btn-outline">Leave Page</button>

<dialog id="unsaved-changes" class="dialog" aria-labelledby="unsaved-changes-title" aria-describedby="unsaved-changes-description">
  <div>
    <header>
      <h2 id="unsaved-changes-title">Unsaved Changes</h2>
      <p id="unsaved-changes-description">You have unsaved changes that will be lost if you leave this page. Are you sure you want to continue?</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('unsaved-changes').close()">Stay on Page</button>
      <button class="btn-destructive" onclick="leavePage(); document.getElementById('unsaved-changes').close()">Leave Without Saving</button>
    </footer>
  </div>
</dialog>
```

### Permission Request

```html
<button type="button" onclick="document.getElementById('permission-request').showModal()" class="btn">Enable Notifications</button>

<dialog id="permission-request" class="dialog" aria-labelledby="permission-request-title" aria-describedby="permission-request-description">
  <div>
    <header>
      <h2 id="permission-request-title">Enable Notifications</h2>
      <p id="permission-request-description">We'd like to show you notifications for the latest news and updates. You can change this setting anytime in your browser preferences.</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('permission-request').close()">Not Now</button>
      <button class="btn" onclick="requestNotificationPermission(); document.getElementById('permission-request').close()">Allow</button>
    </footer>
  </div>
</dialog>
```

### Form Submission Confirmation

```html
<button type="button" onclick="document.getElementById('submit-form').showModal()" class="btn">Submit Application</button>

<dialog id="submit-form" class="dialog" aria-labelledby="submit-form-title" aria-describedby="submit-form-description">
  <div>
    <header>
      <h2 id="submit-form-title">Submit Application</h2>
      <p id="submit-form-description">Once submitted, you will not be able to edit your application. Please review your information before proceeding.</p>
    </header>

    <section>
      <div class="alert">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10" />
          <path d="M12 16v-4" />
          <path d="M12 8h.01" />
        </svg>
        <h2>Important</h2>
        <section>Make sure all information is correct before submitting.</section>
      </div>
    </section>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('submit-form').close()">Review Again</button>
      <button class="btn" onclick="submitApplication(); document.getElementById('submit-form').close()">Submit Application</button>
    </footer>
  </div>
</dialog>
```

### Multiple Actions

```html
<button type="button" onclick="document.getElementById('save-options').showModal()" class="btn-outline">Save Document</button>

<dialog id="save-options" class="dialog" aria-labelledby="save-options-title" aria-describedby="save-options-description">
  <div>
    <header>
      <h2 id="save-options-title">Save Document</h2>
      <p id="save-options-description">How would you like to save your document?</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('save-options').close()">Cancel</button>
      <button class="btn-secondary" onclick="saveDraft(); document.getElementById('save-options').close()">Save as Draft</button>
      <button class="btn" onclick="saveAndPublish(); document.getElementById('save-options').close()">Save & Publish</button>
    </footer>
  </div>
</dialog>
```

### Error Acknowledgment

```html
<dialog id="error-dialog" class="dialog" aria-labelledby="error-dialog-title" aria-describedby="error-dialog-description">
  <div>
    <header>
      <h2 id="error-dialog-title">Upload Failed</h2>
      <p id="error-dialog-description">There was an error uploading your file. Please check your internet connection and try again.</p>
    </header>

    <section>
      <div class="alert-destructive">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10" />
          <line x1="12" x2="12" y1="8" y2="12" />
          <line x1="12" x2="12.01" y1="16" y2="16" />
        </svg>
        <h2>Error Details</h2>
        <section>
          <p>File size exceeds the 10MB limit.</p>
          <ul class="list-inside list-disc text-sm mt-2">
            <li>Maximum file size: 10MB</li>
            <li>Your file size: 15.2MB</li>
            <li>Supported formats: PDF, DOC, DOCX</li>
          </ul>
        </section>
      </div>
    </section>

    <footer>
      <button class="btn" onclick="document.getElementById('error-dialog').close()">OK</button>
    </footer>
  </div>
</dialog>
```

### Complex Confirmation with Form

```html
<button type="button" onclick="document.getElementById('transfer-funds').showModal()" class="btn-destructive">Transfer Funds</button>

<dialog id="transfer-funds" class="dialog" aria-labelledby="transfer-funds-title" aria-describedby="transfer-funds-description">
  <div>
    <header>
      <h2 id="transfer-funds-title">Confirm Transfer</h2>
      <p id="transfer-funds-description">Please verify the transfer details below:</p>
    </header>

    <section>
      <div class="space-y-4">
        <div class="grid grid-cols-2 gap-4 text-sm">
          <div>
            <dt class="font-medium">Amount:</dt>
            <dd class="text-lg font-bold">$2,500.00</dd>
          </div>
          <div>
            <dt class="font-medium">To Account:</dt>
            <dd>****-****-****-1234</dd>
          </div>
          <div>
            <dt class="font-medium">Transfer Date:</dt>
            <dd>Today, March 15, 2024</dd>
          </div>
          <div>
            <dt class="font-medium">Processing Fee:</dt>
            <dd>$2.50</dd>
          </div>
        </div>
        
        <div class="field">
          <label for="confirmation-code">Enter your authentication code</label>
          <input type="password" id="confirmation-code" placeholder="6-digit code" maxlength="6" required>
        </div>
      </div>
    </section>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('transfer-funds').close()">Cancel</button>
      <button class="btn-destructive" onclick="processTransfer(); document.getElementById('transfer-funds').close()">Confirm Transfer</button>
    </footer>
  </div>
</dialog>
```

## Accessibility Features

- **ARIA Labels**: Use `aria-labelledby` and `aria-describedby` for screen readers
- **Modal Behavior**: Traps focus within the dialog
- **Keyboard Support**: ESC key closes dialog (native HTML5 behavior)
- **Focus Management**: Automatically focuses first interactive element
- **Screen Reader Support**: Announces as modal dialog

### Enhanced Accessibility

```html
<dialog id="accessible-alert" class="dialog" role="alertdialog" aria-labelledby="accessible-alert-title" aria-describedby="accessible-alert-description" aria-modal="true">
  <div>
    <header>
      <h2 id="accessible-alert-title">Critical Action Required</h2>
      <p id="accessible-alert-description">This action requires immediate attention and cannot be undone.</p>
    </header>

    <footer>
      <button class="btn-outline" onclick="document.getElementById('accessible-alert').close()" aria-describedby="cancel-description">
        Cancel
      </button>
      <button class="btn-destructive" onclick="performAction(); document.getElementById('accessible-alert').close()" aria-describedby="confirm-description">
        Confirm Action
      </button>
    </footer>
  </div>
</dialog>

<!-- Hidden descriptions for screen readers -->
<p id="cancel-description" class="sr-only">Closes the dialog without performing any action</p>
<p id="confirm-description" class="sr-only">Performs the destructive action and closes the dialog</p>
```

## JavaScript Integration

### Basic Dialog Management

```javascript
// Open alert dialog
function openAlertDialog(dialogId) {
  const dialog = document.getElementById(dialogId);
  if (dialog) {
    dialog.showModal();
  }
}

// Close alert dialog
function closeAlertDialog(dialogId) {
  const dialog = document.getElementById(dialogId);
  if (dialog) {
    dialog.close();
  }
}

// Close with return value
function closeAlertDialogWithValue(dialogId, returnValue) {
  const dialog = document.getElementById(dialogId);
  if (dialog) {
    dialog.close(returnValue);
  }
}
```

### Promise-based Alert Dialogs

```javascript
// Create promise-based alert dialog
function showAlertDialog(title, message, actions = ['Cancel', 'OK']) {
  return new Promise((resolve) => {
    const dialogId = 'alert-dialog-' + Date.now();
    
    // Create dialog HTML
    const dialogHTML = `
      <dialog id="${dialogId}" class="dialog" aria-labelledby="${dialogId}-title" aria-describedby="${dialogId}-description">
        <div>
          <header>
            <h2 id="${dialogId}-title">${title}</h2>
            <p id="${dialogId}-description">${message}</p>
          </header>
          <footer>
            ${actions.map((action, index) => 
              `<button class="${index === 0 ? 'btn-outline' : 'btn'}" 
                      onclick="resolveDialog('${dialogId}', ${index})">${action}</button>`
            ).join('')}
          </footer>
        </div>
      </dialog>
    `;
    
    // Add to page
    document.body.insertAdjacentHTML('beforeend', dialogHTML);
    
    // Store resolver
    window.dialogResolvers = window.dialogResolvers || {};
    window.dialogResolvers[dialogId] = resolve;
    
    // Open dialog
    document.getElementById(dialogId).showModal();
  });
}

// Resolve dialog promise
function resolveDialog(dialogId, actionIndex) {
  const dialog = document.getElementById(dialogId);
  const resolver = window.dialogResolvers[dialogId];
  
  if (dialog && resolver) {
    dialog.close();
    resolver(actionIndex);
    
    // Cleanup
    setTimeout(() => {
      dialog.remove();
      delete window.dialogResolvers[dialogId];
    }, 100);
  }
}

// Usage examples
showAlertDialog('Delete File', 'Are you sure you want to delete this file?', ['Cancel', 'Delete'])
  .then((actionIndex) => {
    if (actionIndex === 1) {
      console.log('User confirmed deletion');
      deleteFile();
    }
  });

showAlertDialog('Save Changes', 'You have unsaved changes. What would you like to do?', ['Cancel', 'Don\'t Save', 'Save'])
  .then((actionIndex) => {
    switch (actionIndex) {
      case 1:
        discardChanges();
        break;
      case 2:
        saveChanges();
        break;
    }
  });
```

### React Integration

```jsx
import React, { useState, useEffect } from 'react';

function AlertDialog({ isOpen, title, description, actions = [], onAction, onClose }) {
  const dialogRef = useRef(null);
  
  useEffect(() => {
    const dialog = dialogRef.current;
    if (!dialog) return;
    
    if (isOpen) {
      dialog.showModal();
    } else {
      dialog.close();
    }
  }, [isOpen]);
  
  const handleAction = (actionIndex) => {
    onAction?.(actionIndex);
    onClose?.();
  };
  
  const handleKeyDown = (e) => {
    if (e.key === 'Escape') {
      e.preventDefault(); // Prevent default ESC behavior
      onClose?.();
    }
  };
  
  return (
    <dialog 
      ref={dialogRef}
      className="dialog" 
      aria-labelledby="alert-dialog-title" 
      aria-describedby="alert-dialog-description"
      onKeyDown={handleKeyDown}
    >
      <div>
        <header>
          <h2 id="alert-dialog-title">{title}</h2>
          {description && (
            <p id="alert-dialog-description">{description}</p>
          )}
        </header>
        
        <footer>
          {actions.map((action, index) => (
            <button 
              key={index}
              className={index === 0 ? 'btn-outline' : 'btn'}
              onClick={() => handleAction(index)}
            >
              {action}
            </button>
          ))}
        </footer>
      </div>
    </dialog>
  );
}

// Usage
function App() {
  const [alertDialog, setAlertDialog] = useState({ isOpen: false });
  
  const showDeleteConfirmation = () => {
    setAlertDialog({
      isOpen: true,
      title: 'Delete Item',
      description: 'Are you sure you want to delete this item? This action cannot be undone.',
      actions: ['Cancel', 'Delete']
    });
  };
  
  const handleAlertAction = (actionIndex) => {
    if (actionIndex === 1) {
      // User clicked "Delete"
      deleteItem();
    }
    setAlertDialog({ isOpen: false });
  };
  
  return (
    <div>
      <button onClick={showDeleteConfirmation} className="btn-destructive">
        Delete Item
      </button>
      
      <AlertDialog
        {...alertDialog}
        onAction={handleAlertAction}
        onClose={() => setAlertDialog({ isOpen: false })}
      />
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div>
    <button @click="showDeleteDialog" class="btn-destructive">Delete Item</button>
    
    <dialog 
      ref="alertDialog"
      class="dialog" 
      aria-labelledby="alert-dialog-title" 
      aria-describedby="alert-dialog-description"
      @keydown.esc="closeDialog"
    >
      <div>
        <header>
          <h2 id="alert-dialog-title">{{ title }}</h2>
          <p v-if="description" id="alert-dialog-description">{{ description }}</p>
        </header>
        
        <footer>
          <button 
            v-for="(action, index) in actions" 
            :key="index"
            :class="index === 0 ? 'btn-outline' : 'btn'"
            @click="handleAction(index)"
          >
            {{ action }}
          </button>
        </footer>
      </div>
    </dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      title: '',
      description: '',
      actions: [],
      resolver: null
    };
  },
  methods: {
    showAlertDialog(title, description, actions = ['Cancel', 'OK']) {
      return new Promise((resolve) => {
        this.title = title;
        this.description = description;
        this.actions = actions;
        this.resolver = resolve;
        this.$refs.alertDialog.showModal();
      });
    },
    
    showDeleteDialog() {
      this.showAlertDialog(
        'Delete Item', 
        'Are you sure you want to delete this item? This action cannot be undone.',
        ['Cancel', 'Delete']
      ).then((actionIndex) => {
        if (actionIndex === 1) {
          this.deleteItem();
        }
      });
    },
    
    handleAction(actionIndex) {
      this.closeDialog();
      if (this.resolver) {
        this.resolver(actionIndex);
      }
    },
    
    closeDialog() {
      this.$refs.alertDialog.close();
      if (this.resolver) {
        this.resolver(-1); // Indicates dialog was closed without action
      }
    },
    
    deleteItem() {
      console.log('Item deleted');
    }
  }
};
</script>
```

## Best Practices

1. **Critical Actions Only**: Use for destructive or irreversible actions
2. **Clear Messaging**: Make consequences crystal clear
3. **Action Hierarchy**: Primary action on the right, cancel on the left
4. **No Backdrop Close**: Don't allow closing by clicking backdrop
5. **Descriptive Buttons**: Use specific action words like "Delete" not "Yes"
6. **Appropriate Icons**: Include warning icons for destructive actions
7. **Keyboard Support**: Ensure proper keyboard navigation
8. **Mobile-Friendly**: Actions stack vertically on small screens

## Common Patterns

### Delete Confirmation

```html
<dialog id="delete-confirm" class="dialog" aria-labelledby="delete-confirm-title">
  <div>
    <header>
      <h2 id="delete-confirm-title">Delete Item</h2>
      <p>This item will be permanently deleted and cannot be recovered.</p>
    </header>
    <footer>
      <button class="btn-outline" onclick="document.getElementById('delete-confirm').close()">Cancel</button>
      <button class="btn-destructive" onclick="confirmDelete()">Delete Forever</button>
    </footer>
  </div>
</dialog>
```

### Unsaved Changes Warning

```html
<dialog id="unsaved-warning" class="dialog" aria-labelledby="unsaved-warning-title">
  <div>
    <header>
      <h2 id="unsaved-warning-title">Unsaved Changes</h2>
      <p>You have unsaved changes that will be lost. Are you sure you want to leave?</p>
    </header>
    <footer>
      <button class="btn-outline" onclick="document.getElementById('unsaved-warning').close()">Stay</button>
      <button class="btn-destructive" onclick="discardAndLeave()">Leave Without Saving</button>
    </footer>
  </div>
</dialog>
```

### Permission Request

```html
<dialog id="permission-dialog" class="dialog" aria-labelledby="permission-dialog-title">
  <div>
    <header>
      <h2 id="permission-dialog-title">Location Access</h2>
      <p>This app would like to use your location to provide personalized recommendations.</p>
    </header>
    <footer>
      <button class="btn-outline" onclick="document.getElementById('permission-dialog').close()">Not Now</button>
      <button class="btn" onclick="grantLocationAccess()">Allow</button>
    </footer>
  </div>
</dialog>
```

## Related Components

- [Dialog](./dialog.md) - For general purpose dialogs
- [Alert](./alert.md) - For inline notifications
- [Toast](./toast.md) - For temporary feedback messages
- [Button](./button.md) - For dialog actions