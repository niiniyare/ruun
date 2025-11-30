# Radio Group Component

A set of checkable buttons—known as radio buttons—where no more than one of the buttons can be checked at a time.

## Basic Usage

```html
<fieldset class="grid gap-3">
  <label class="label"><input type="radio" name="size" value="default" class="input">Default</label>
  <label class="label"><input type="radio" name="size" value="comfortable" class="input" checked>Comfortable</label>
  <label class="label"><input type="radio" name="size" value="compact" class="input">Compact</label>
</fieldset>
```

## CSS Classes

### Primary Classes
- **`input`** - Applied to radio input elements
- **`label`** - Applied to label elements
- **`fieldset`** - Container for grouping radio buttons
- **`form`** - Parent container for automatic styling

### Tailwind Utilities Used
- `grid gap-3` - Vertical layout with spacing
- `flex flex-col` - Alternative vertical layout
- `font-normal` - Normal font weight for labels
- `leading-tight` - Tight line height
- `text-muted-foreground` - Muted text color for descriptions
- `text-sm` - Small text size for descriptions

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Must be "radio" | Yes |
| `name` | string | Group identifier (same for all options) | Yes |
| `value` | string | Value when selected | Yes |
| `class` | string | Must include "input" class | Yes |
| `id` | string | Unique identifier for label association | Recommended |
| `checked` | boolean | Pre-selected option | No |
| `disabled` | boolean | Disables specific option | No |
| `required` | boolean | Makes group selection required | No |

## Examples

### Basic Radio Group
```html
<fieldset class="grid gap-3">
  <legend class="label">Choose your display size</legend>
  <label class="label"><input type="radio" name="display" value="default" class="input">Default</label>
  <label class="label"><input type="radio" name="display" value="comfortable" class="input" checked>Comfortable</label>
  <label class="label"><input type="radio" name="display" value="compact" class="input">Compact</label>
</fieldset>
```

### With Descriptions
```html
<fieldset class="grid gap-4">
  <legend class="label">Notification preferences</legend>
  
  <div class="flex items-start gap-3">
    <input type="radio" id="all-notifications" name="notifications" value="all" class="input">
    <div class="grid gap-1">
      <label for="all-notifications" class="label">All notifications</label>
      <p class="text-muted-foreground text-sm">Receive all email notifications and updates</p>
    </div>
  </div>
  
  <div class="flex items-start gap-3">
    <input type="radio" id="important-only" name="notifications" value="important" class="input" checked>
    <div class="grid gap-1">
      <label for="important-only" class="label">Important only</label>
      <p class="text-muted-foreground text-sm">Only receive critical updates and messages</p>
    </div>
  </div>
  
  <div class="flex items-start gap-3">
    <input type="radio" id="no-notifications" name="notifications" value="none" class="input">
    <div class="grid gap-1">
      <label for="no-notifications" class="label">None</label>
      <p class="text-muted-foreground text-sm">Turn off all email notifications</p>
    </div>
  </div>
</fieldset>
```

### Form Context
```html
<form class="form space-y-6 w-full">
  <div class="flex flex-col gap-3">
    <label>Notify me about...</label>
    <fieldset class="grid gap-3">
      <label class="font-normal">
        <input type="radio" name="notifications" value="all">
        All new messages
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="mentions">
        Direct messages and mentions
      </label>
      <label class="font-normal">
        <input type="radio" name="notifications" value="none">
        Nothing
      </label>
    </fieldset>
  </div>
  <button type="submit" class="btn">Save preferences</button>
</form>
```

### Disabled Options
```html
<fieldset class="grid gap-3">
  <legend class="label">Account type</legend>
  <label class="label"><input type="radio" name="account" value="free" class="input" checked>Free</label>
  <label class="label"><input type="radio" name="account" value="premium" class="input">Premium</label>
  <label class="label opacity-50"><input type="radio" name="account" value="enterprise" class="input" disabled>Enterprise (Coming soon)</label>
</fieldset>
```

### Inline Radio Group
```html
<fieldset class="flex gap-6">
  <legend class="label sr-only">Gender</legend>
  <label class="label flex items-center gap-2">
    <input type="radio" name="gender" value="male" class="input">
    Male
  </label>
  <label class="label flex items-center gap-2">
    <input type="radio" name="gender" value="female" class="input">
    Female
  </label>
  <label class="label flex items-center gap-2">
    <input type="radio" name="gender" value="other" class="input">
    Other
  </label>
</fieldset>
```

### Card-style Radio Group
```html
<fieldset class="grid gap-4">
  <legend class="label">Choose your plan</legend>
  
  <label class="flex items-center gap-4 rounded-lg border p-4 hover:bg-muted cursor-pointer has-[:checked]:ring-2 has-[:checked]:ring-primary">
    <input type="radio" name="plan" value="basic" class="input">
    <div class="flex-1">
      <h3 class="font-medium">Basic Plan</h3>
      <p class="text-muted-foreground text-sm">Perfect for getting started</p>
      <p class="font-semibold">$9/month</p>
    </div>
  </label>
  
  <label class="flex items-center gap-4 rounded-lg border p-4 hover:bg-muted cursor-pointer has-[:checked]:ring-2 has-[:checked]:ring-primary">
    <input type="radio" name="plan" value="pro" class="input">
    <div class="flex-1">
      <h3 class="font-medium">Pro Plan</h3>
      <p class="text-muted-foreground text-sm">For growing businesses</p>
      <p class="font-semibold">$29/month</p>
    </div>
  </label>
</fieldset>
```

## States

### Default State
Unselected radio button with normal appearance

### Selected State
Shows filled circle when selected (only one per group)

### Hover State
Slight visual feedback on hover

### Focus State
Visible focus ring for keyboard navigation

### Disabled State
Reduced opacity and no interaction

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, radio groups inherit styling:

```html
<form class="form space-y-6">
  <fieldset>
    <legend>Payment method</legend>
    <!-- Radio buttons inherit form styling -->
    <label><input type="radio" name="payment" value="card"> Credit Card</label>
    <label><input type="radio" name="payment" value="paypal"> PayPal</label>
  </fieldset>
</form>
```

### Validation
```html
<fieldset class="grid gap-3" required>
  <legend class="label">Required selection</legend>
  <label class="label">
    <input type="radio" name="required-group" value="option1" class="input" required>
    Option 1
  </label>
  <label class="label">
    <input type="radio" name="required-group" value="option2" class="input">
    Option 2
  </label>
</fieldset>
```

## JavaScript Integration

### Getting Selected Value
```javascript
function getSelectedValue(groupName) {
  const selected = document.querySelector(`input[name="${groupName}"]:checked`);
  return selected ? selected.value : null;
}

// Example usage
const selectedPlan = getSelectedValue('plan');
console.log('Selected plan:', selectedPlan);
```

### Change Event Handling
```javascript
document.querySelectorAll('input[type="radio"]').forEach(radio => {
  radio.addEventListener('change', (e) => {
    if (e.target.checked) {
      console.log(`Selected ${e.target.name}: ${e.target.value}`);
    }
  });
});
```

### HTMX Integration
```html
<fieldset class="grid gap-3">
  <label class="label">
    <input type="radio" name="theme" value="light" class="input" 
           hx-post="/api/update-theme" 
           hx-trigger="change">
    Light theme
  </label>
  <label class="label">
    <input type="radio" name="theme" value="dark" class="input"
           hx-post="/api/update-theme" 
           hx-trigger="change">
    Dark theme
  </label>
</fieldset>
```

## Accessibility Features

- **Keyboard Navigation**: Arrow keys navigate within group, Tab moves between groups
- **Screen Reader Support**: Proper grouping with `fieldset` and `legend`
- **Focus Management**: Visual focus indicators and logical tab order
- **Label Association**: Proper label-input relationships
- **Group Semantics**: `fieldset` and `legend` provide group context

## Best Practices

1. Always use `fieldset` and `legend` for grouping
2. Ensure all radio buttons in a group have the same `name` attribute
3. Provide unique `value` attributes for each option
4. Use descriptive labels and help text
5. Consider the logical order of options
6. Mark required groups appropriately
7. Test with keyboard navigation
8. Ensure sufficient visual distinction between states

## Layout Patterns

### Vertical Stack (Default)
```html
<fieldset class="grid gap-3">
  <!-- Radio buttons stacked vertically -->
</fieldset>
```

### Horizontal Layout
```html
<fieldset class="flex gap-4">
  <!-- Radio buttons in a row -->
</fieldset>
```

### Grid Layout
```html
<fieldset class="grid grid-cols-2 gap-3">
  <!-- Radio buttons in a grid -->
</fieldset>
```

### Card Layout
```html
<fieldset class="grid gap-4">
  <!-- Each radio button in a card -->
</fieldset>
```

## Related Components

- [Checkbox](./checkbox.md) - For multiple selections
- [Switch](./switch.md) - For binary toggles
- [Select](./select.md) - For dropdown selections
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling