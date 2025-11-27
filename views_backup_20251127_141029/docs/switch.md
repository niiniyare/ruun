# Switch Component

A control that allows the user to toggle between on and off states, similar to a physical switch.

## Basic Usage

```html
<label class="label">
  <input type="checkbox" name="switch" role="switch" class="input">
  Airplane Mode
</label>
```

## CSS Classes

### Primary Classes
- **`input`** - Applied to checkbox input element with `role="switch"`
- **`label`** - Applied to label wrapper
- **`form`** - Parent container for automatic styling

### Tailwind Utilities Used
- `gap-2` - Spacing between switch and content
- `flex flex-row` - Horizontal layout
- `items-start` - Vertical alignment
- `justify-between` - Space between content and switch
- `rounded-lg` - Container border radius
- `border` - Container border
- `p-4` - Container padding
- `shadow-xs` - Subtle shadow
- `opacity-60` - Disabled state opacity
- `leading-normal` - Normal line height
- `text-muted-foreground` - Muted text color
- `text-sm` - Small text size

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Must be "checkbox" | Yes |
| `role` | string | Must be "switch" | Yes |
| `class` | string | Must include "input" class | Yes |
| `id` | string | Unique identifier for label association | Recommended |
| `name` | string | Form field name | No |
| `checked` | boolean | Initial on/off state | No |
| `disabled` | boolean | Disables the switch | No |

## Examples

### Basic Switch
```html
<label class="label">
  <input type="checkbox" name="switch" role="switch" class="input">
  Airplane Mode
</label>
```

### Switch with Description
```html
<form class="form grid gap-4">
  <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
    <div class="flex flex-col gap-0.5">
      <label for="email-notifications" class="leading-normal">Email notifications</label>
      <p class="text-muted-foreground text-sm">Receive emails about new products, features, and more.</p>
    </div>
    <input type="checkbox" id="email-notifications" role="switch">
  </div>
</form>
```

### Disabled Switch
```html
<div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
  <div class="flex flex-col gap-0.5 opacity-60">
    <label for="disabled-switch" class="leading-normal">Marketing emails</label>
    <p class="text-muted-foreground text-sm">This feature is currently unavailable.</p>
  </div>
  <input type="checkbox" id="disabled-switch" role="switch" disabled>
</div>
```

### Group of Switches
```html
<form class="form grid gap-4">
  <h3 class="text-lg font-medium">Notification Preferences</h3>
  
  <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
    <div class="flex flex-col gap-0.5">
      <label for="push-notifications" class="leading-normal">Push notifications</label>
      <p class="text-muted-foreground text-sm">Receive push notifications on your device.</p>
    </div>
    <input type="checkbox" id="push-notifications" role="switch" checked>
  </div>
  
  <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
    <div class="flex flex-col gap-0.5">
      <label for="email-notifications" class="leading-normal">Email notifications</label>
      <p class="text-muted-foreground text-sm">Receive emails about updates and news.</p>
    </div>
    <input type="checkbox" id="email-notifications" role="switch">
  </div>
  
  <div class="gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs">
    <div class="flex flex-col gap-0.5">
      <label for="sms-notifications" class="leading-normal">SMS notifications</label>
      <p class="text-muted-foreground text-sm">Receive text messages for urgent updates.</p>
    </div>
    <input type="checkbox" id="sms-notifications" role="switch">
  </div>
</form>
```

### Simple Switch List
```html
<div class="grid gap-4">
  <label class="flex items-center justify-between">
    <span>Dark mode</span>
    <input type="checkbox" role="switch" class="input">
  </label>
  
  <label class="flex items-center justify-between">
    <span>Auto-save</span>
    <input type="checkbox" role="switch" class="input" checked>
  </label>
  
  <label class="flex items-center justify-between">
    <span>Show tooltips</span>
    <input type="checkbox" role="switch" class="input">
  </label>
</div>
```

### Form Integration
```html
<form class="form space-y-6">
  <div class="grid gap-4">
    <h3 class="text-lg font-medium">Account Settings</h3>
    
    <div class="space-y-4">
      <div class="flex items-center justify-between">
        <div>
          <label for="public-profile" class="label">Public profile</label>
          <p class="text-muted-foreground text-sm">Make your profile visible to other users</p>
        </div>
        <input type="checkbox" id="public-profile" name="public_profile" role="switch">
      </div>
      
      <div class="flex items-center justify-between">
        <div>
          <label for="activity-status" class="label">Show activity status</label>
          <p class="text-muted-foreground text-sm">Let others see when you're online</p>
        </div>
        <input type="checkbox" id="activity-status" name="activity_status" role="switch" checked>
      </div>
    </div>
  </div>
  
  <button type="submit" class="btn">Save settings</button>
</form>
```

## States

### Off State (Unchecked)
Switch appears in the "off" position

### On State (Checked)
Switch appears in the "on" position with active styling

### Focus State
Visible focus ring for keyboard navigation

### Disabled State
Reduced opacity and no interaction

### Hover State
Subtle visual feedback on hover

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, switches inherit proper styling:

```html
<form class="form">
  <!-- Switch inherits form styling -->
  <label>
    <input type="checkbox" role="switch" name="notifications">
    Enable notifications
  </label>
</form>
```

### Value Handling
Switches submit as checkbox values:
- **Checked**: Submits the value (or "on" if no value specified)
- **Unchecked**: Does not submit any value

```html
<input type="checkbox" role="switch" name="feature" value="enabled">
<!-- When checked, submits: feature=enabled -->
<!-- When unchecked, submits nothing -->
```

## JavaScript Integration

### Getting Switch State
```javascript
function getSwitchState(switchId) {
  const switchElement = document.getElementById(switchId);
  return switchElement.checked;
}

// Example usage
const isDarkModeEnabled = getSwitchState('dark-mode-switch');
console.log('Dark mode:', isDarkModeEnabled);
```

### Toggle Switch Programmatically
```javascript
function toggleSwitch(switchId) {
  const switchElement = document.getElementById(switchId);
  switchElement.checked = !switchElement.checked;
  
  // Trigger change event
  switchElement.dispatchEvent(new Event('change', { bubbles: true }));
}
```

### Change Event Handling
```javascript
document.querySelectorAll('input[role="switch"]').forEach(switchElement => {
  switchElement.addEventListener('change', (e) => {
    const isEnabled = e.target.checked;
    const switchName = e.target.name || e.target.id;
    
    console.log(`${switchName} is now ${isEnabled ? 'ON' : 'OFF'}`);
    
    // Save state to localStorage
    if (switchName) {
      localStorage.setItem(switchName, isEnabled.toString());
    }
  });
});
```

### HTMX Integration
```html
<input type="checkbox" 
       role="switch" 
       name="notifications"
       hx-post="/api/toggle-notifications" 
       hx-trigger="change"
       hx-target="#notification-status">
```

### Theme Switching Example
```javascript
// Initialize dark mode switch
const darkModeSwitch = document.getElementById('dark-mode');
const isDarkMode = localStorage.getItem('darkMode') === 'true';
darkModeSwitch.checked = isDarkMode;

// Apply initial theme
document.documentElement.classList.toggle('dark', isDarkMode);

// Handle theme changes
darkModeSwitch.addEventListener('change', (e) => {
  const isDark = e.target.checked;
  document.documentElement.classList.toggle('dark', isDark);
  localStorage.setItem('darkMode', isDark.toString());
});
```

## Accessibility Features

- **Keyboard Support**: Space key toggles, Tab for navigation
- **Screen Reader**: Announced as "switch" control type
- **State Announcement**: "on" or "off" state clearly communicated
- **Label Association**: Proper label-input relationships
- **Focus Management**: Visible focus indicators

## Best Practices

1. Use switches for binary settings that take effect immediately
2. Use checkboxes for options that require form submission
3. Always provide clear labels explaining what the switch controls
4. Include descriptions for complex settings
5. Group related switches logically
6. Consider the default state carefully
7. Test with keyboard navigation and screen readers
8. Provide visual feedback for state changes

## Switch vs. Checkbox vs. Radio

### Use Switch When:
- Toggling a system setting on/off
- Changes take effect immediately
- Binary choice (on/off, enabled/disabled)
- User expects immediate feedback

### Use Checkbox When:
- Multiple selections allowed
- Changes require form submission
- Part of a larger form

### Use Radio When:
- Mutually exclusive options
- More than two choices
- One option must be selected

## Styling Variations

### Compact Switches
```html
<label class="flex items-center gap-2 text-sm">
  <input type="checkbox" role="switch" class="input scale-75">
  Compact switch
</label>
```

### Large Switches
```html
<label class="flex items-center gap-3 text-lg">
  <input type="checkbox" role="switch" class="input scale-125">
  Large switch
</label>
```

### Card-style Layout
```html
<div class="rounded-lg border p-4 space-y-2">
  <div class="flex items-center justify-between">
    <h4 class="font-medium">Feature Name</h4>
    <input type="checkbox" role="switch" class="input">
  </div>
  <p class="text-muted-foreground text-sm">Detailed description of what this feature does.</p>
</div>
```

## Related Components

- [Checkbox](./checkbox.md) - For multiple selections
- [Radio Group](./radio-group.md) - For mutually exclusive options
- [Button](./button.md) - For action triggers
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling