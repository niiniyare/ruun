# Form Component

Building forms with Basecoat components.

## Basic Usage

```html
<form class="form grid gap-6">
  <div class="grid gap-2">
    <label for="username">Username</label>
    <input type="text" id="username" placeholder="hunvreus">
    <p class="text-muted-foreground text-sm">This is your public display name.</p>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

## CSS Classes

### Form Container Classes
- **`form`** - Main form styling that cascades to child elements
- **`grid gap-6`** - Recommended layout classes for spacing

### Child Element Styling
The `form` class automatically applies styling to child elements:
- **`label`** - Automatically styled when inside `.form`
- **`input`** - Automatically styled when inside `.form`  
- **`textarea`** - Automatically styled when inside `.form`
- **`select`** - Automatically styled when inside `.form`

### Field Layout Classes
- **`grid gap-2`** - Standard field layout with spacing
- **`flex flex-col gap-3`** - Alternative column layout for radio groups
- **`flex flex-row items-start justify-between`** - Horizontal layout for switches

### Supporting Classes
- **`text-muted-foreground text-sm`** - Helper text styling
- **`font-normal`** - Reset font weight for radio/checkbox labels
- **`leading-normal`** - Normal line height for switch labels
- **`rounded-lg border p-4 shadow-xs`** - Card-style containers

## Component Attributes

### Form Element
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "form" for automatic styling | Yes |
| `method` | string | HTTP method (GET/POST) | Optional |
| `action` | string | Form submission URL | Optional |

### No JavaScript Required (Basic)
Basic forms work with standard HTML form handling.

## HTML Structure

```html
<!-- Basic form structure -->
<form class="form grid gap-6">
  <!-- Text field -->
  <div class="grid gap-2">
    <label for="field-id">Label</label>
    <input type="text" id="field-id" placeholder="Placeholder">
    <p class="text-muted-foreground text-sm">Helper text</p>
  </div>
  
  <!-- Radio group -->
  <div class="flex flex-col gap-3">
    <label for="radio-group">Group Label</label>
    <fieldset id="radio-group" class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="group" value="1" checked>
        Option 1
      </label>
    </fieldset>
  </div>
  
  <!-- Switch/checkbox section -->
  <section class="grid gap-4">
    <h3 class="text-lg font-medium">Section Title</h3>
    <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
      <div class="flex flex-col gap-0.5">
        <label for="switch" class="leading-normal">Switch Label</label>
        <p class="text-muted-foreground text-sm">Switch description</p>
      </div>
      <input type="checkbox" id="switch" role="switch">
    </div>
  </section>
  
  <!-- Submit -->
  <button type="submit" class="btn">Submit</button>
</form>
```

## Examples

### Complete User Profile Form

```html
<form class="form grid gap-6">
  <div class="grid gap-2">
    <label for="username">Username</label>
    <input type="text" id="username" placeholder="hunvreus">
    <p class="text-muted-foreground text-sm">This is your public display name.</p>
  </div>

  <div class="grid gap-2">
    <label for="email">Email</label>
    <select id="email">
      <option value="m@example.com">m@example.com</option>
      <option value="m@google.com">m@google.com</option>
      <option value="m@support.com">m@support.com</option>
    </select>
    <p class="text-muted-foreground text-sm">You can manage email addresses in your email settings.</p>
  </div>

  <div class="grid gap-2">
    <label for="bio">Bio</label>
    <textarea id="bio" placeholder="I like to..." rows="3"></textarea>
    <p class="text-muted-foreground text-sm">You can @mention other users and organizations.</p>
  </div>

  <div class="grid gap-2">
    <label for="birth-date">Date of birth</label>
    <input type="date" id="birth-date">
    <p class="text-muted-foreground text-sm">Your date of birth is used to calculate your age.</p>
  </div>

  <div class="flex flex-col gap-3">
    <label for="notifications">Notify me about...</label>
    <fieldset id="notifications" class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="notifications" value="all" checked>
        All new messages
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="direct">
        Direct messages and mentions
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="none">
        Nothing
      </label>
    </fieldset>
  </div>

  <section class="grid gap-4">
    <h3 class="text-lg font-medium">Email Notifications</h3>
    <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
      <div class="flex flex-col gap-0.5">
        <label for="marketing" class="leading-normal">Marketing emails</label>
        <p class="text-muted-foreground text-sm">Receive emails about new products, features, and more.</p>
      </div>
      <input type="checkbox" id="marketing" role="switch">
    </div>
    <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
      <div class="flex flex-col gap-0.5 opacity-60">
        <label for="security" class="leading-normal">Security emails</label>
        <p class="text-muted-foreground text-sm">Receive emails about your account security.</p>
      </div>
      <input type="checkbox" id="security" role="switch" disabled>
    </div>
  </section>

  <button type="submit" class="btn">Update profile</button>
</form>
```

### Contact Form

```html
<form class="form grid gap-6" method="POST" action="/contact">
  <div class="grid gap-2">
    <label for="name">Full Name</label>
    <input type="text" id="name" name="name" required placeholder="John Doe">
  </div>

  <div class="grid gap-2">
    <label for="contact-email">Email Address</label>
    <input type="email" id="contact-email" name="email" required placeholder="john@example.com">
  </div>

  <div class="grid gap-2">
    <label for="subject">Subject</label>
    <select id="subject" name="subject" required>
      <option value="">Select a subject</option>
      <option value="general">General Inquiry</option>
      <option value="support">Technical Support</option>
      <option value="billing">Billing Question</option>
      <option value="feedback">Feedback</option>
    </select>
  </div>

  <div class="grid gap-2">
    <label for="message">Message</label>
    <textarea id="message" name="message" required placeholder="How can we help you?" rows="5"></textarea>
    <p class="text-muted-foreground text-sm">Please be as detailed as possible.</p>
  </div>

  <div class="flex flex-col gap-3">
    <label for="priority">Priority Level</label>
    <fieldset id="priority" class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="priority" value="low" checked>
        Low - General inquiry
      </label>
      <label class="font-normal">
        <input type="radio" name="priority" value="medium">
        Medium - Need assistance
      </label>
      <label class="font-normal">
        <input type="radio" name="priority" value="high">
        High - Urgent issue
      </label>
    </fieldset>
  </div>

  <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
    <div class="flex flex-col gap-0.5">
      <label for="newsletter" class="leading-normal">Subscribe to newsletter</label>
      <p class="text-muted-foreground text-sm">Get updates about new features and products.</p>
    </div>
    <input type="checkbox" id="newsletter" name="newsletter" role="switch">
  </div>

  <div class="flex gap-3">
    <button type="submit" class="btn">Send Message</button>
    <button type="reset" class="btn-outline">Clear Form</button>
  </div>
</form>
```

### Settings Form with Validation

```html
<form class="form grid gap-6" id="settings-form">
  <div class="grid gap-2">
    <label for="current-password">Current Password</label>
    <input type="password" id="current-password" name="current_password" required>
  </div>

  <div class="grid gap-2">
    <label for="new-password">New Password</label>
    <input type="password" id="new-password" name="new_password" required minlength="8">
    <p class="text-muted-foreground text-sm">Must be at least 8 characters long.</p>
  </div>

  <div class="grid gap-2">
    <label for="confirm-password">Confirm New Password</label>
    <input type="password" id="confirm-password" name="confirm_password" required>
  </div>

  <div class="grid gap-2">
    <label for="timezone">Timezone</label>
    <select id="timezone" name="timezone">
      <option value="UTC">UTC</option>
      <option value="America/New_York">Eastern Time</option>
      <option value="America/Chicago">Central Time</option>
      <option value="America/Denver">Mountain Time</option>
      <option value="America/Los_Angeles">Pacific Time</option>
    </select>
  </div>

  <section class="grid gap-4">
    <h3 class="text-lg font-medium">Privacy Settings</h3>
    <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
      <div class="flex flex-col gap-0.5">
        <label for="profile-public" class="leading-normal">Public profile</label>
        <p class="text-muted-foreground text-sm">Allow others to see your profile information.</p>
      </div>
      <input type="checkbox" id="profile-public" name="profile_public" role="switch" checked>
    </div>
    <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
      <div class="flex flex-col gap-0.5">
        <label for="activity-public" class="leading-normal">Public activity</label>
        <p class="text-muted-foreground text-sm">Show your activity to other users.</p>
      </div>
      <input type="checkbox" id="activity-public" name="activity_public" role="switch">
    </div>
  </section>

  <div class="flex flex-col gap-3">
    <label for="delete-data">Data retention</label>
    <fieldset id="delete-data" class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="data_retention" value="keep" checked>
        Keep my data indefinitely
      </label>
      <label class="font-normal">
        <input type="radio" name="data_retention" value="1year">
        Delete after 1 year of inactivity
      </label>
      <label class="font-normal">
        <input type="radio" name="data_retention" value="6months">
        Delete after 6 months of inactivity
      </label>
    </fieldset>
  </div>

  <div class="flex gap-3">
    <button type="submit" class="btn">Save Changes</button>
    <button type="button" class="btn-outline">Cancel</button>
  </div>
</form>
```

### Multi-Step Wizard Form

```html
<form class="form grid gap-6" id="wizard-form">
  <!-- Step indicator -->
  <div class="flex justify-between items-center mb-4">
    <div class="flex space-x-4">
      <div class="flex items-center">
        <div class="rounded-full h-8 w-8 flex items-center justify-center border-2 border-primary bg-primary text-primary-foreground">1</div>
        <span class="ml-2 text-sm">Personal Info</span>
      </div>
      <div class="flex items-center">
        <div class="rounded-full h-8 w-8 flex items-center justify-center border-2 border-muted">2</div>
        <span class="ml-2 text-sm text-muted-foreground">Account</span>
      </div>
      <div class="flex items-center">
        <div class="rounded-full h-8 w-8 flex items-center justify-center border-2 border-muted">3</div>
        <span class="ml-2 text-sm text-muted-foreground">Preferences</span>
      </div>
    </div>
  </div>

  <!-- Step 1: Personal Info -->
  <div id="step-1" class="step">
    <div class="grid gap-4">
      <div class="grid gap-2">
        <label for="first-name">First Name</label>
        <input type="text" id="first-name" name="first_name" required>
      </div>

      <div class="grid gap-2">
        <label for="last-name">Last Name</label>
        <input type="text" id="last-name" name="last_name" required>
      </div>

      <div class="grid gap-2">
        <label for="wizard-email">Email</label>
        <input type="email" id="wizard-email" name="email" required>
      </div>

      <div class="grid gap-2">
        <label for="phone">Phone Number</label>
        <input type="tel" id="phone" name="phone" placeholder="+1 (555) 123-4567">
      </div>
    </div>
  </div>

  <!-- Navigation -->
  <div class="flex justify-between">
    <button type="button" class="btn-outline" disabled>Previous</button>
    <button type="button" class="btn">Next</button>
  </div>
</form>
```

## Accessibility Features

- **Semantic HTML**: Uses proper form elements and structure
- **Label Association**: All inputs properly linked to labels
- **Required Fields**: Uses `required` attribute for validation
- **Fieldsets**: Groups related options with `fieldset` and `legend`
- **Helper Text**: Descriptive text linked with `aria-describedby`
- **Switch Role**: Checkboxes used as switches have `role="switch"`

### Enhanced Accessibility

```html
<form class="form grid gap-6" aria-labelledby="form-title">
  <h2 id="form-title">Account Settings</h2>
  
  <div class="grid gap-2">
    <label for="accessible-username">Username</label>
    <input 
      type="text" 
      id="accessible-username" 
      name="username"
      required
      aria-describedby="username-help username-error"
      aria-invalid="false"
    >
    <p id="username-help" class="text-muted-foreground text-sm">
      Choose a unique username between 3-20 characters.
    </p>
    <p id="username-error" class="text-red-500 text-sm hidden" role="alert">
      Username must be at least 3 characters long.
    </p>
  </div>

  <fieldset class="flex flex-col gap-3">
    <legend>Notification Preferences</legend>
    <div class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="notifications" value="all" checked>
        All notifications
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="important">
        Important only
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="none">
        None
      </label>
    </div>
  </fieldset>
  
  <button type="submit" class="btn" aria-describedby="submit-help">
    Save Changes
  </button>
  <p id="submit-help" class="text-muted-foreground text-sm">
    Changes will take effect immediately.
  </p>
</form>
```

## JavaScript Integration

### Form Validation

```javascript
// Basic form validation
const form = document.getElementById('settings-form');

form.addEventListener('submit', (e) => {
  e.preventDefault();
  
  const formData = new FormData(form);
  const newPassword = formData.get('new_password');
  const confirmPassword = formData.get('confirm_password');
  
  // Validation
  if (newPassword !== confirmPassword) {
    alert('Passwords do not match');
    return;
  }
  
  // Submit form
  fetch(form.action, {
    method: 'POST',
    body: formData
  }).then(response => {
    if (response.ok) {
      // Show success message
    }
  });
});

// Real-time validation
const usernameInput = document.getElementById('username');
const usernameError = document.getElementById('username-error');

usernameInput.addEventListener('input', (e) => {
  const value = e.target.value;
  const isValid = value.length >= 3;
  
  e.target.setAttribute('aria-invalid', !isValid);
  usernameError.classList.toggle('hidden', isValid);
});
```

### React Integration

```jsx
import React, { useState } from 'react';

function ContactForm() {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    message: '',
    newsletter: false
  });

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Form submitted:', formData);
  };

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  return (
    <form className="form grid gap-6" onSubmit={handleSubmit}>
      <div className="grid gap-2">
        <label htmlFor="name">Name</label>
        <input
          type="text"
          id="name"
          name="name"
          value={formData.name}
          onChange={handleChange}
          required
        />
      </div>

      <div className="grid gap-2">
        <label htmlFor="email">Email</label>
        <input
          type="email"
          id="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          required
        />
      </div>

      <div className="grid gap-2">
        <label htmlFor="message">Message</label>
        <textarea
          id="message"
          name="message"
          value={formData.message}
          onChange={handleChange}
          rows={4}
          required
        />
      </div>

      <div className="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
        <div className="flex flex-col gap-0.5">
          <label htmlFor="newsletter" className="leading-normal">
            Subscribe to newsletter
          </label>
          <p className="text-muted-foreground text-sm">
            Get updates about new features.
          </p>
        </div>
        <input
          type="checkbox"
          id="newsletter"
          name="newsletter"
          role="switch"
          checked={formData.newsletter}
          onChange={handleChange}
        />
      </div>

      <button type="submit" className="btn">
        Send Message
      </button>
    </form>
  );
}
```

### Vue Integration

```vue
<template>
  <form class="form grid gap-6" @submit.prevent="handleSubmit">
    <div class="grid gap-2">
      <label for="vue-name">Name</label>
      <input
        type="text"
        id="vue-name"
        v-model="form.name"
        required
      />
    </div>

    <div class="grid gap-2">
      <label for="vue-email">Email</label>
      <input
        type="email"
        id="vue-email"
        v-model="form.email"
        required
      />
    </div>

    <div class="flex flex-col gap-3">
      <label>Contact Method</label>
      <fieldset class="grid gap-3">
        <label class="font-normal">
          <input
            type="radio"
            v-model="form.contactMethod"
            value="email"
          />
          Email
        </label>
        <label class="font-normal">
          <input
            type="radio"
            v-model="form.contactMethod"
            value="phone"
          />
          Phone
        </label>
      </fieldset>
    </div>

    <button type="submit" class="btn">Submit</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: '',
        email: '',
        contactMethod: 'email'
      }
    };
  },
  methods: {
    handleSubmit() {
      console.log('Form submitted:', this.form);
    }
  }
};
</script>
```

## Best Practices

1. **Use Semantic HTML**: Always use proper form elements and attributes
2. **Label Everything**: Every input should have an associated label
3. **Group Related Fields**: Use fieldsets for radio/checkbox groups
4. **Provide Feedback**: Include helper text and validation messages
5. **Accessible Validation**: Use aria attributes for error states
6. **Progressive Enhancement**: Ensure forms work without JavaScript
7. **Clear Actions**: Make submit buttons descriptive
8. **Consistent Spacing**: Use the recommended grid gap classes

## Common Patterns

### Field with Validation State

```html
<div class="grid gap-2">
  <label for="validated-field">Email Address</label>
  <input 
    type="email" 
    id="validated-field" 
    class="border-red-500" 
    aria-invalid="true"
    aria-describedby="email-error"
  >
  <p id="email-error" class="text-red-500 text-sm" role="alert">
    Please enter a valid email address.
  </p>
</div>
```

### Optional Field Indicator

```html
<div class="grid gap-2">
  <label for="optional-field">
    Phone Number
    <span class="text-muted-foreground font-normal">(optional)</span>
  </label>
  <input type="tel" id="optional-field">
</div>
```

### Inline Form

```html
<form class="form flex gap-4 items-end">
  <div class="grid gap-2 flex-1">
    <label for="inline-email">Email</label>
    <input type="email" id="inline-email" placeholder="Enter your email">
  </div>
  <button type="submit" class="btn">Subscribe</button>
</form>
```

## Related Components

- [Field](./field.md) - Individual form field wrapper
- [Input](./input.md) - Text input component
- [Textarea](./textarea.md) - Multi-line text input
- [Select](./select.md) - Dropdown selection
- [Button](./button.md) - Form submission buttons
- [Checkbox](./checkbox.md) - Checkbox inputs
- [Radio Group](./radio-group.md) - Radio button groups
- [Label](./label.md) - Form labels