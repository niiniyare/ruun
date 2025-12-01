# Ruun Package - Component Documentation

Auto-collected from views/docs/*.md

---
## input

# Input Component

Displays a form input field or a component that looks like an input field.

## Basic Usage

```html
<input class="input" type="email" placeholder="Email">
```

## CSS Classes

### Primary Class
- **`input`** - Main input styling class

### Tailwind Utilities Used
- `w-full` - Full width
- `rounded-md` - Medium border radius
- `border` - Border styling
- `bg-input` - Input background color
- `px-3` - Horizontal padding
- `py-2` - Vertical padding
- `text-sm` - Small text size
- `transition-colors` - Color transition effect
- `focus:outline-none` - Remove default focus outline
- `focus:ring-2` - Focus ring size
- `focus:ring-offset-2` - Focus ring offset

### State Classes
- **Invalid state**: `aria-invalid="true"` - Changes border to error color
- **Disabled state**: `disabled` - Applies muted styling and cursor

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "input" class | Yes |
| `type` | string | Input type (text, email, password, etc.) | No |
| `placeholder` | string | Placeholder text | No |
| `disabled` | boolean | Disables the input | No |
| `aria-invalid` | boolean | Marks input as invalid for error states | No |
| `id` | string | Unique identifier for label association | No |

## Examples

### Default Input
```html
<input class="input" type="email" placeholder="Email">
```

### Invalid Input
```html
<input class="input" type="email" placeholder="Email" aria-invalid="true">
```

### Disabled Input
```html
<input class="input" type="email" placeholder="Email" disabled>
```

### With Label
```html
<div class="grid gap-3">
  <label for="input-with-label" class="label">Label</label>
  <input class="input" id="input-with-label" type="email" placeholder="Email">
</div>
```

### With Help Text
```html
<div class="grid gap-3">
  <label for="input-with-text" class="label">Label</label>
  <input class="input" id="input-with-text" type="email" placeholder="Email">
  <p class="text-muted-foreground text-sm">Fill in your email address.</p>
</div>
```

### With Button
```html
<div class="flex items-center space-x-2">
  <input class="input" type="email" placeholder="Email">
  <button type="submit" class="btn">Submit</button>
</div>
```

### Form Context
```html
<form class="form space-y-6 w-full">
  <div class="grid gap-3">
    <label for="input-form" class="label">Username</label>
    <input class="input" id="input-form" type="text" placeholder="hunvreus">
    <p class="text-muted-foreground text-sm">This is your public display name.</p>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

## JavaScript Integration

- **Copy functionality**: Uses navigator.clipboard.writeText() for code copying
- **HTMX Support**: Can be used with HTMX attributes for dynamic behavior
- **Theme support**: Automatically adapts to light/dark themes

## Accessibility Features

- Supports `aria-invalid` for screen readers
- Works with `for` attribute on labels
- Focus management with visible focus rings
- Proper color contrast in all states

## Related Components

- [Label](./label.md) - For accessible labeling
- [Form](./form.md) - For form context styling
- [Button](./button.md) - Often used together
---
## badge

# Badge Component

Displays a badge or a component that looks like a badge.

## Basic Usage

```html
<span class="badge">Badge</span>
```

## CSS Classes

### Variant Classes
- **`badge`** - Primary badge (default styling)
- **`badge-primary`** - Primary variant
- **`badge-secondary`** - Secondary variant  
- **`badge-destructive`** - Destructive/error variant
- **`badge-outline`** - Outline variant

### Tailwind Utilities Used
- `inline-flex` - Inline flex display
- `items-center` - Center items vertically
- `rounded-full` - Full border radius for circular badges
- `px-2.5` - Horizontal padding
- `py-0.5` - Vertical padding
- `text-xs` - Extra small text
- `font-medium` - Medium font weight
- `border` - Border styling (outline variant)

### Modifier Classes
- `rounded-full` - Makes badge circular
- `h-5 min-w-5` - Fixed height and minimum width for counters
- `px-1` - Reduced padding for small badges
- `font-mono tabular-nums` - Monospace font for numbers

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include badge variant class | Yes |

## Examples

### Basic Variants
```html
<span class="badge">Badge</span>
<span class="badge-secondary">Secondary</span>
<span class="badge-destructive">Destructive</span>
<span class="badge-outline">Outline</span>
```

### With Custom Colors
```html
<span class="badge-secondary bg-blue-500 text-white dark:bg-blue-600">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z" />
    <path d="m9 12 2 2 4-4" />
  </svg>
  Verified
</span>
```

### Counter Badges
```html
<span class="badge rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">8</span>
<span class="badge-destructive rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">99</span>
<span class="badge-outline rounded-full h-5 min-w-5 px-1 font-mono tabular-nums">20+</span>
```

### With Icons
```html
<span class="badge-destructive">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <circle cx="12" cy="12" r="10" />
    <line x1="12" x2="12" y1="8" y2="12" />
    <line x1="12" x2="12.01" y1="16" y2="16" />
  </svg>
  With icon
</span>
```

### As Links
```html
<a href="#" class="badge-outline">
  Link
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M5 12h14" />
    <path d="m12 5 7 7-7 7" />
  </svg>
</a>
```

### Layout Examples
```html
<div class="flex flex-col items-center gap-2">
  <div class="flex w-full flex-wrap gap-2">
    <span class="badge">Badge</span>
    <span class="badge-secondary">Secondary</span>
    <span class="badge-destructive">Destructive</span>
    <span class="badge-outline">Outline</span>
  </div>
</div>
```

## Color Variants

### Primary
- Background: Primary theme color
- Text: Primary foreground color
- Use: Default/primary actions

### Secondary  
- Background: Secondary theme color
- Text: Secondary foreground color
- Use: Secondary information

### Destructive
- Background: Error/danger color
- Text: Error foreground color  
- Use: Errors, warnings, deletions

### Outline
- Background: Transparent
- Border: Border color
- Text: Foreground color
- Use: Subtle emphasis

## JavaScript Integration

- **Copy functionality**: Code examples include clipboard copying
- **Theme support**: Adapts to light/dark themes automatically
- **Interactive states**: Can be used as clickable elements (links, buttons)

## Accessibility Features

- Semantic HTML structure
- Proper contrast ratios
- Screen reader friendly
- Focus states when interactive

## Best Practices

1. Use meaningful text or numbers
2. Consider accessibility for color-blind users
3. Use appropriate variant for context
4. Keep text concise for small badges
5. Use `font-mono tabular-nums` for numeric badges

## Related Components

- [Button](./button.md) - For interactive actions
- [Avatar](./avatar.md) - Often used with status badges
- [Alert](./alert.md) - For larger status messages
---
## label

# Label Component

Renders an accessible label associated with controls.

## Basic Usage

```html
<label class="label" for="email">Your email address</label>
```

## CSS Classes

### Primary Class
- **`label`** - Main label styling class

### Tailwind Utilities Used
- `text-sm` - Small text size
- `font-medium` - Medium font weight
- `leading-none` - No line height
- `peer-disabled:cursor-not-allowed` - Disabled cursor when peer is disabled
- `peer-disabled:opacity-70` - Reduced opacity when peer is disabled

### Layout Classes
- `gap-3` - For spacing between label and control when used as flex container

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "label" class | Yes |
| `for` | string | ID of associated form control | Recommended |

## Examples

### Basic Label
```html
<label class="label" for="email">Your email address</label>
```

### Label with Input
```html
<div class="grid gap-3">
  <label class="label" for="email">Email</label>
  <input class="input" id="email" type="email" placeholder="Email">
</div>
```

### Inline Label with Control
```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

### Label with Disabled Input
```html
<div class="grid gap-3">
  <label class="label" for="email">Email</label>
  <input class="input" id="email" type="email" placeholder="Email" disabled>
</div>
```

## State Behavior

### Disabled State
When the associated input is disabled, the label automatically:
- Shows a "not-allowed" cursor
- Reduces opacity to 70%
- Maintains accessibility

This is handled through CSS peer selectors:
```css
.peer-disabled:cursor-not-allowed
.peer-disabled:opacity-70
```

### Form Context
Labels automatically inherit proper styling when used within a `.form` container, eliminating the need to manually add the `label` class.

## Layout Patterns

### Vertical Layout (Recommended)
```html
<div class="grid gap-3">
  <label class="label" for="input-field">Field Label</label>
  <input class="input" id="input-field" type="text">
</div>
```

### Horizontal Layout  
```html
<label class="label gap-3 flex items-center">
  <input type="checkbox" class="input">
  <span>Checkbox Label</span>
</label>
```

## JavaScript Integration

- **Theme support**: Automatically adapts to light/dark themes
- **Form validation**: Works with validation libraries
- **HTMX support**: Compatible with dynamic form updates

## Accessibility Features

- **Semantic association**: Uses `for` attribute to associate with controls
- **Screen reader support**: Properly announced by assistive technology  
- **Focus management**: Clicking label focuses associated control
- **Disabled state**: Visual and interaction changes when control is disabled
- **Form context**: Maintains accessibility within form groups

## Best Practices

1. Always use `for` attribute with unique ID of associated control
2. Keep label text clear and concise  
3. Place labels above inputs in vertical layouts
4. Use inline labels only for checkboxes and radio buttons
5. Ensure sufficient color contrast
6. Test with screen readers

## Form Integration

Labels work seamlessly with the form component system:

```html
<form class="form space-y-6">
  <div class="grid gap-3">
    <label for="username">Username</label> <!-- Inherits label styling -->
    <input id="username" type="text" placeholder="Enter username">
  </div>
</form>
```

## Related Components

- [Input](./input.md) - Primary form control
- [Form](./form.md) - Form container context
- [Checkbox](./checkbox.md) - Checkbox inputs
- [Radio Group](./radio-group.md) - Radio button groups
- [Textarea](./textarea.md) - Multi-line text inputs
---
## checkbox

# Checkbox Component

A control that allows the user to toggle between checked and not checked.

## Basic Usage

```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

## CSS Classes

### Primary Classes
- **`input`** - Applied to the checkbox input element
- **`label`** - Applied to the label wrapper
- **`form`** - Parent container that provides automatic styling

### Tailwind Utilities Used
- `gap-3` - Spacing between checkbox and text
- `flex` - Flexbox layout for labels
- `items-start` - Vertical alignment
- `text-muted-foreground` - Muted text color
- `text-sm` - Small text size
- `leading-snug` - Tight line height

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "input" class | Yes |
| `type` | string | Must be "checkbox" | Yes |
| `id` | string | Unique identifier for label association | Recommended |
| `name` | string | Form field name for multiple checkboxes | No |
| `value` | string | Value when checked | No |
| `checked` | boolean | Pre-checked state | No |
| `disabled` | boolean | Disables the checkbox | No |

## Examples

### Basic Checkbox
```html
<label class="label gap-3">
  <input type="checkbox" class="input">
  Accept terms and conditions
</label>
```

### Disabled Checkbox
```html
<label class="label gap-3">
  <input type="checkbox" class="input" disabled>
  Accept terms and conditions
</label>
```

### With Descriptive Text
```html
<div class="flex items-start gap-3">
  <input type="checkbox" id="checkbox-with-text" class="input">
  <div class="grid gap-2">
    <label for="checkbox-with-text" class="label">Accept terms and conditions</label>
    <p class="text-muted-foreground text-sm">By clicking this checkbox, you agree to the terms and conditions.</p>
  </div>
</div>
```

### Form Context (Auto-styling)
```html
<form class="form flex flex-row items-start gap-3 rounded-md border p-4 shadow-xs">
  <input type="checkbox" id="checkbox-form-1">
  <div class="flex flex-col gap-1">
    <label for="checkbox-form-1" class="leading-snug">Use different settings for my mobile devices</label>
    <p class="text-muted-foreground text-sm leading-snug">You can manage your mobile notifications in the mobile settings page.</p>
  </div>
</form>
```

### Multiple Checkboxes in Fieldset
```html
<form class="form flex flex-col gap-4">
  <fieldset id="demo-form-checkboxes" class="flex flex-col gap-2">
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="1" checked>
      Recents
    </label>
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="2" checked>
      Home
    </label>
    <label class="font-normal leading-tight">
      <input type="checkbox" name="demo-form-checkboxes" value="3">
      Applications
    </label>
  </fieldset>
</form>
```

## States

### Default State
Regular unchecked appearance with focus styles

### Checked State
Shows checkmark when selected

### Disabled State
Reduced opacity and no interaction

### Focus State
Visible focus ring for keyboard navigation

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, checkboxes inherit styling automatically:

```html
<form class="form">
  <!-- No need for explicit classes on checkbox -->
  <input type="checkbox" id="auto-styled">
  <label for="auto-styled">Auto-styled checkbox</label>
</form>
```

### Multiple Selection
Use the same `name` attribute for related checkboxes:

```html
<fieldset>
  <legend>Select your preferences:</legend>
  <label><input type="checkbox" name="preferences" value="email"> Email notifications</label>
  <label><input type="checkbox" name="preferences" value="sms"> SMS notifications</label>
  <label><input type="checkbox" name="preferences" value="push"> Push notifications</label>
</fieldset>
```

## JavaScript Integration

- **HTMX Support**: Compatible with HTMX attributes for dynamic behavior
- **Form Validation**: Works with native and custom validation
- **Theme Support**: Automatically adapts to light/dark themes

## Accessibility Features

- **Keyboard Navigation**: Full keyboard support with Tab and Space
- **Screen Reader**: Proper labeling and state announcement
- **Focus Management**: Visible focus indicators
- **Form Association**: Proper label-input relationships

## Best Practices

1. Always associate with a label using `for` attribute or wrapping
2. Use `fieldset` and `legend` for related checkbox groups
3. Provide descriptive text for complex choices
4. Consider the reading order for screen readers
5. Use consistent naming conventions for grouped checkboxes

## Related Components

- [Radio Group](./radio-group.md) - For mutually exclusive options
- [Switch](./switch.md) - For on/off toggles
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling
---
## select

# Select Component

Displays a list of options for the user to pick from—triggered by a button.

## Basic Usage

```html
<select class="select w-[180px]">
  <option>Apple</option>
  <option>Banana</option>
  <option>Orange</option>
</select>
```

## CSS Classes

### Primary Classes
- **`select`** - Applied to select element or custom select wrapper
- **`btn-outline`** - Button styling for custom select trigger
- **`justify-between`** - Space between text and chevron icon
- **`font-normal`** - Normal font weight

### Tailwind Utilities Used
- `w-[180px]` - Fixed width (customizable)
- `truncate` - Text overflow handling
- `text-muted-foreground` - Muted text color
- `opacity-50` - Icon opacity
- `shrink-0` - Prevent icon shrinking

## Component Types

### Native Select
Uses the browser's native select element with custom styling.

### Custom Select
Enhanced select with JavaScript functionality and custom appearance.

## Examples

### Native Select
```html
<select class="select w-[180px]">
  <optgroup label="Fruits">
    <option>Apple</option>
    <option>Banana</option>
    <option>Blueberry</option>
    <option>Grapes</option>
    <option>Orange</option>
  </optgroup>
  <optgroup label="Vegetables">
    <option>Aubergine</option>
    <option>Broccoli</option>
    <option>Carrot</option>
    <option selected>Courgette</option>
  </optgroup>
</select>
```

### Custom Select (JavaScript-Enhanced)
```html
<div id="select-942316" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[180px]" 
          id="select-942316-trigger" 
          aria-haspopup="listbox" 
          aria-expanded="false" 
          aria-controls="select-942316-listbox">
    <span class="truncate">Apple</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" 
         fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" 
         class="lucide lucide-chevrons-up-down text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="select-942316-popover" data-popover aria-hidden="true">
    <!-- Options list -->
    <div role="listbox" id="select-942316-listbox" aria-labelledby="select-942316-trigger">
      <div role="option" class="option">Apple</div>
      <div role="option" class="option">Banana</div>
      <div role="option" class="option">Orange</div>
    </div>
  </div>
</div>
```

### Disabled Select
```html
<div id="select-disabled" class="select">
  <button type="button" class="btn-outline justify-between font-normal" 
          id="select-disabled-trigger" 
          aria-haspopup="listbox" 
          aria-expanded="false" 
          aria-controls="select-disabled-listbox" 
          disabled="disabled">
    <span class="truncate">Select an option</span>
    <svg class="lucide lucide-chevron-down text-muted-foreground opacity-50 shrink-0">
      <path d="m6 9 6 6 6-6" />
    </svg>
  </button>
</div>
```

### With Label
```html
<div class="grid gap-3">
  <label for="select-with-label" class="label">Choose a fruit</label>
  <select id="select-with-label" class="select w-[180px]">
    <option value="">Select a fruit</option>
    <option value="apple">Apple</option>
    <option value="banana">Banana</option>
    <option value="orange">Orange</option>
  </select>
</div>
```

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "select" class | Yes |
| `id` | string | Unique identifier | Recommended |
| `name` | string | Form field name | No |
| `disabled` | boolean | Disables the select | No |
| `multiple` | boolean | Allow multiple selections | No |
| `size` | number | Visible options count | No |

## Custom Select Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `aria-haspopup` | string | "listbox" for select behavior | Yes |
| `aria-expanded` | string | "true"/"false" for open state | Yes |
| `aria-controls` | string | ID of listbox element | Yes |
| `data-popover` | string | Marks popover container | Yes |
| `role` | string | "listbox" and "option" for options | Yes |

## JavaScript Integration

### Required Scripts
The custom select requires JavaScript for functionality:

```html
<script src="/assets/js/select.js" defer></script>
```

### Popover System
Uses the `data-popover` attribute for dropdown positioning and behavior:

```javascript
// Automatic initialization when page loads
document.addEventListener('DOMContentLoaded', () => {
  if (window.basecoat) window.basecoat.initAll();
});
```

### Event Handling
Custom selects dispatch standard `change` events:

```javascript
document.getElementById('my-select-trigger').addEventListener('change', (e) => {
  console.log('Selected value:', e.detail.value);
});
```

## Accessibility Features

- **Keyboard Navigation**: Arrow keys, Enter, Escape, and type-ahead
- **Screen Reader Support**: Proper ARIA attributes and roles
- **Focus Management**: Maintains focus correctly when opening/closing
- **Label Association**: Works with `<label>` elements
- **Option Groups**: Support for `<optgroup>` in native selects

## States

### Default State
Closed select with placeholder or selected value

### Open State
Dropdown is visible with options list

### Focused State
Visible focus ring around trigger button

### Disabled State
Reduced opacity and no interaction

### Selected State
Highlighted option and updated trigger text

## Form Integration

### Native Form Support
```html
<form class="form space-y-4">
  <div class="grid gap-3">
    <label for="fruit-select">Favorite Fruit</label>
    <select id="fruit-select" name="fruit" class="select">
      <option value="">Choose a fruit</option>
      <option value="apple">Apple</option>
      <option value="banana">Banana</option>
    </select>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

### Custom Select in Forms
Custom selects include hidden input fields for form submission:

```html
<form class="form space-y-4">
  <div id="custom-select" class="select">
    <input type="hidden" name="fruit" value="apple">
    <button type="button" class="btn-outline justify-between font-normal">
      <span class="truncate">Apple</span>
      <!-- Chevron icon -->
    </button>
    <!-- Popover with options -->
  </div>
</form>
```

## Best Practices

1. Use native selects for simple lists and better mobile experience
2. Use custom selects when you need enhanced styling or functionality
3. Always provide a default "empty" option for non-required fields
4. Group related options with `<optgroup>` for native selects
5. Ensure adequate contrast for all states
6. Test thoroughly with keyboard navigation
7. Consider mobile-first design (native selects work better on mobile)

## Styling Customization

### Width Control
```html
<select class="select w-32">...</select>   <!-- 128px width -->
<select class="select w-48">...</select>   <!-- 192px width -->
<select class="select w-full">...</select> <!-- Full width -->
```

### Size Variants
```html
<!-- Small -->
<select class="select text-sm h-8">...</select>

<!-- Default -->
<select class="select">...</select>

<!-- Large -->
<select class="select text-lg h-12">...</select>
```

## Related Components

- [Input](./input.md) - For text input fields
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling
- [Popover](./popover.md) - For custom dropdown positioning
- [Button](./button.md) - For trigger styling
---
## textarea

# Textarea Component

Displays a form textarea or a component that looks like a textarea.

## Basic Usage

```html
<textarea class="textarea" placeholder="Type your message here"></textarea>
```

## CSS Classes

### Primary Class
- **`textarea`** - Main textarea styling class

### Tailwind Utilities Used
- `w-full` - Full width
- `rounded-md` - Medium border radius
- `border` - Border styling
- `bg-input` - Input background color
- `px-3` - Horizontal padding
- `py-2` - Vertical padding
- `text-sm` - Small text size
- `transition-colors` - Color transition effect
- `focus:outline-none` - Remove default focus outline
- `focus:ring-2` - Focus ring size
- `focus:ring-offset-2` - Focus ring offset
- `resize-none` - Disable manual resizing (optional)
- `min-h-[80px]` - Minimum height

## Component Attributes

| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "textarea" class | Yes |
| `placeholder` | string | Placeholder text | No |
| `disabled` | boolean | Disables the textarea | No |
| `aria-invalid` | boolean | Marks textarea as invalid for error states | No |
| `id` | string | Unique identifier for label association | No |
| `name` | string | Form field name | No |
| `rows` | number | Visible text lines | No |
| `cols` | number | Visible character width | No |
| `maxlength` | number | Maximum character count | No |
| `required` | boolean | Makes field required | No |

## Examples

### Basic Textarea
```html
<textarea class="textarea" placeholder="Type your message here"></textarea>
```

### Invalid Textarea
```html
<textarea class="textarea" placeholder="Type your message here" aria-invalid="true"></textarea>
```

### Disabled Textarea
```html
<textarea class="textarea" placeholder="Type your message here" disabled></textarea>
```

### With Label
```html
<div class="grid gap-3">
  <label for="textarea-with-label" class="label">Message</label>
  <textarea id="textarea-with-label" class="textarea" placeholder="Type your message here"></textarea>
</div>
```

### With Help Text
```html
<div class="grid gap-3">
  <label for="textarea-with-help" class="label">Comment</label>
  <textarea id="textarea-with-help" class="textarea" placeholder="Type your message here"></textarea>
  <p class="text-muted-foreground text-sm">Type your message and press Ctrl+Enter to send.</p>
</div>
```

### With Character Counter
```html
<div class="grid gap-3">
  <label for="textarea-counter" class="label">Bio</label>
  <textarea id="textarea-counter" class="textarea" placeholder="Tell us about yourself" maxlength="160"></textarea>
  <p class="text-muted-foreground text-sm">0/160 characters</p>
</div>
```

### Form Context
```html
<form class="form space-y-6">
  <div class="grid gap-3">
    <label for="textarea-form" class="label">Bio</label>
    <textarea id="textarea-form" placeholder="Tell us a bit about yourself"></textarea>
    <p class="text-muted-foreground text-sm">You can @mention other users and organizations.</p>
  </div>
  <button type="submit" class="btn">Submit</button>
</form>
```

### Auto-resizing Textarea
```html
<textarea 
  class="textarea resize-none overflow-hidden" 
  placeholder="This textarea grows as you type..."
  oninput="this.style.height = 'auto'; this.style.height = this.scrollHeight + 'px'"></textarea>
```

### Fixed Height Textarea
```html
<textarea class="textarea h-32 resize-none" placeholder="Fixed height textarea"></textarea>
```

## Size Variants

### Small
```html
<textarea class="textarea text-sm min-h-[60px]" placeholder="Small textarea"></textarea>
```

### Default
```html
<textarea class="textarea" placeholder="Default textarea"></textarea>
```

### Large
```html
<textarea class="textarea text-base min-h-[120px]" placeholder="Large textarea"></textarea>
```

## States

### Default State
Regular appearance with placeholder text

### Focus State
Visible focus ring and enhanced border

### Invalid State
Error styling when `aria-invalid="true"` is set

### Disabled State
Reduced opacity and no interaction

### Filled State
Contains user input text

## Form Integration

### Automatic Styling
When using `class="form"` on a parent container, textareas inherit styling automatically:

```html
<form class="form space-y-4">
  <!-- No need for explicit textarea class -->
  <div class="grid gap-3">
    <label for="auto-textarea">Description</label>
    <textarea id="auto-textarea" placeholder="Automatically styled"></textarea>
  </div>
</form>
```

### Validation Integration
```html
<div class="grid gap-3">
  <label for="validated-textarea" class="label">Required Field</label>
  <textarea 
    id="validated-textarea" 
    class="textarea" 
    required 
    aria-describedby="textarea-error">
  </textarea>
  <p id="textarea-error" class="text-error text-sm hidden">This field is required</p>
</div>
```

## JavaScript Integration

### Auto-resize Functionality
```javascript
function autoResize(textarea) {
  textarea.style.height = 'auto';
  textarea.style.height = textarea.scrollHeight + 'px';
}

// Apply to all auto-resize textareas
document.querySelectorAll('.textarea[data-auto-resize]').forEach(textarea => {
  textarea.addEventListener('input', () => autoResize(textarea));
  autoResize(textarea); // Initial sizing
});
```

### Character Counter
```javascript
function updateCounter(textarea) {
  const counter = textarea.nextElementSibling;
  if (counter && counter.dataset.counter) {
    const maxLength = textarea.maxLength;
    const currentLength = textarea.value.length;
    counter.textContent = `${currentLength}/${maxLength} characters`;
    
    // Visual feedback when approaching limit
    if (currentLength > maxLength * 0.9) {
      counter.classList.add('text-warning');
    } else {
      counter.classList.remove('text-warning');
    }
  }
}
```

### HTMX Integration
```html
<textarea 
  class="textarea" 
  hx-post="/api/save-draft" 
  hx-trigger="keyup changed delay:500ms" 
  hx-target="#save-status"
  placeholder="Auto-saves as you type">
</textarea>
<div id="save-status"></div>
```

## Accessibility Features

- **Keyboard Navigation**: Full keyboard support including Tab navigation
- **Screen Reader Support**: Proper labeling with `aria-describedby`
- **Error States**: `aria-invalid` for validation feedback
- **Focus Management**: Clear focus indicators
- **Label Association**: Proper `for` attribute usage

## Best Practices

1. Always associate with a label using `for` attribute
2. Provide helpful placeholder text
3. Use appropriate sizing for content type
4. Consider auto-resize for dynamic content
5. Implement character counting for limited fields
6. Provide clear validation feedback
7. Test with keyboard navigation
8. Ensure sufficient contrast in all states

## Styling Customization

### Border Variants
```html
<!-- Default border -->
<textarea class="textarea"></textarea>

<!-- No border -->
<textarea class="textarea border-0"></textarea>

<!-- Thick border -->
<textarea class="textarea border-2"></textarea>
```

### Background Variants
```html
<!-- Default background -->
<textarea class="textarea"></textarea>

<!-- Transparent background -->
<textarea class="textarea bg-transparent"></textarea>

<!-- Custom background -->
<textarea class="textarea bg-muted"></textarea>
```

### Resize Control
```html
<!-- No resize -->
<textarea class="textarea resize-none"></textarea>

<!-- Vertical resize only -->
<textarea class="textarea resize-y"></textarea>

<!-- Horizontal resize only -->
<textarea class="textarea resize-x"></textarea>

<!-- Both directions (default) -->
<textarea class="textarea resize"></textarea>
```

## Related Components

- [Input](./input.md) - For single-line text input
- [Label](./label.md) - For proper labeling
- [Form](./form.md) - For form context styling
- [Button](./button.md) - Often used with textarea in forms
---
## radio-group

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
---
## switch

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
---
## field

# Field Component

Combine labels, controls, and help text to compose accessible form fields.

## Basic Usage

```html
<div role="group" class="field">
  <label for="username">Username</label>
  <input id="username" type="text" placeholder="evilrabbit" aria-describedby="username-desc">
  <p id="username-desc">Choose a unique username for your account.</p>
</div>
```

## CSS Classes

### Primary Classes
- **`field`** - Applied to individual field container with `role="group"`
- **`fieldset`** - Applied to fieldset element for grouping related fields

### Tailwind Utilities Used
- `gap-3` - Spacing within field containers
- `grid gap-6` - Layout spacing between fields
- `grid gap-3` - Tighter spacing for checkbox/radio groups
- `flex flex-col` - Vertical layout for field groups
- `w-full max-w-md` - Width constraints
- `font-normal` - Normal font weight for labels
- `grid-cols-3 gap-4` - Grid layout for multiple fields
- `data-orientation="horizontal"` - Horizontal field layout
- `self-start` - Vertical alignment for horizontal fields

## Component Structure

### Field Container
A field consists of a container with `role="group" class="field"` containing:
- Label element
- Form control (input, select, textarea, etc.)
- Optional helper text
- Optional error message

### Fieldset Container  
A fieldset groups related fields with `<fieldset class="fieldset">` containing:
- Legend element for the group title
- Optional description paragraph
- Multiple field containers or form controls

## Component Attributes

### Field Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "group" | Yes |
| `class` | string | Must include "field" class | Yes |
| `data-orientation` | string | "horizontal" for side-by-side layout | No |

### Form Control Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier | Recommended |
| `aria-describedby` | string | References helper text and error elements | Recommended |
| `aria-invalid` | boolean | "true" when field has validation errors | No |

### Helper Text and Error Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Referenced by aria-describedby | Yes |
| `role` | string | "alert" for error messages | For errors |

## Examples

### Basic Field
```html
<div role="group" class="field">
  <label for="username">Username</label>
  <input id="username" type="text" placeholder="evilrabbit" aria-describedby="username-desc">
  <p id="username-desc">Choose a unique username for your account.</p>
</div>
```

### Field with Error
```html
<div role="group" class="field">
  <label for="username">Username</label>
  <input id="username" type="text" placeholder="@evilrabbit" aria-invalid="true" aria-describedby="username-error">
  <p id="username-error" role="alert">Username is already taken</p>
</div>
```

### Multiple Fields
```html
<div class="grid gap-6">
  <div role="group" class="field">
    <label for="username">Username</label>
    <input id="username" type="text" placeholder="evilrabbit" aria-describedby="username-desc">
    <p id="username-desc">Choose a unique username for your account.</p>
  </div>

  <div role="group" class="field">
    <label for="password">Password</label>
    <p id="password-desc">Must be at least 8 characters long.</p>
    <input id="password" type="password" placeholder="••••••••" aria-describedby="password-desc">
  </div>
</div>
```

### Textarea Field
```html
<div role="group" class="field">
  <label for="feedback">Feedback</label>
  <textarea id="feedback" rows="4" placeholder="Your feedback helps us improve..." aria-describedby="feedback-desc"></textarea>
  <p id="feedback-desc">Share your thoughts about our service.</p>
</div>
```

### Select Field
```html
<div role="group" class="field">
  <label for="department">Department</label>
  <select id="department" class="select w-full" aria-describedby="department-desc">
    <option value="">Choose department</option>
    <option value="engineering">Engineering</option>
    <option value="design">Design</option>
    <option value="marketing">Marketing</option>
    <option value="sales">Sales</option>
  </select>
  <p id="department-desc">Select your department or area of work.</p>
</div>
```

### Range/Slider Field
```html
<div role="group" class="field">
  <label for="price-range">Price</label>
  <p id="price-range-desc">Set your budget: $<span id="price-range-value">150</span></p>
  <input id="price-range" type="range" min="0" max="500" value="150">
</div>
```

### Horizontal Field Layout
```html
<div class="field" data-orientation="horizontal">
  <input id="sync-desktop-documents" type="checkbox" checked class="input self-start">
  <section>
    <label for="sync-desktop-documents">Sync Desktop & Documents folders</label>
    <p>Your Desktop & Documents folders are being synced with iCloud Drive. You can access them from other devices.</p>
  </section>
</div>
```

## Fieldset Examples

### Basic Fieldset
```html
<fieldset class="fieldset">
  <legend>Profile</legend>
  <p>This information will be displayed on your profile</p>
  
  <div role="group" class="field">
    <label for="full-name">Full name</label>
    <input id="full-name" type="text" placeholder="Evil Rabbit" aria-describedby="name-desc">
    <p id="name-desc">Your first and last name</p>
  </div>
  
  <div role="group" class="field">
    <label for="display-name">Username</label>
    <input id="display-name" type="text" placeholder="@evilrabbit" aria-invalid="true" aria-describedby="username-error">
    <p id="username-error" role="alert">Username is already taken</p>
  </div>
</fieldset>
```

### Fieldset with Payment Form
```html
<form class="w-full max-w-md space-y-6">
  <fieldset class="fieldset">
    <legend>Payment Method</legend>
    <p>All transactions are secure and encrypted</p>
    
    <div role="group" class="field">
      <label for="card-name">Name on Card</label>
      <input id="card-name" type="text" placeholder="Evil Rabbit" required>
    </div>
    
    <div role="group" class="field">
      <label for="card-number">Card Number</label>
      <input id="card-number" type="text" placeholder="1234 5678 9012 3456" aria-describedby="card-number-desc" required>
      <p id="card-number-desc">Enter your 16-digit card number</p>
    </div>
    
    <div class="grid grid-cols-3 gap-4">
      <div role="group" class="field">
        <label for="exp-month">Month</label>
        <select id="exp-month" class="select w-full">
          <option value="">MM</option>
          <option value="01">01</option>
          <option value="02">02</option>
        </select>
      </div>
      <div role="group" class="field">
        <label for="exp-year">Year</label>
        <select id="exp-year" class="select w-full">
          <option value="">YYYY</option>
          <option value="2024">2024</option>
          <option value="2025">2025</option>
        </select>
      </div>
      <div role="group" class="field">
        <label for="cvv">CVV</label>
        <input id="cvv" type="text" placeholder="123" required>
      </div>
    </div>
  </fieldset>
</form>
```

### Checkbox Fieldset
```html
<fieldset class="fieldset">
  <legend>Show these items on the desktop</legend>
  <p>Select the items you want to show on the desktop.</p>
  
  <div class="flex flex-col gap-3">
    <div class="field">
      <label class="gap-3">
        <input type="checkbox">
        Hard disks
      </label>
    </div>
    <div class="field">
      <label class="gap-3">
        <input type="checkbox">
        External disks
      </label>
    </div>
    <div class="field">
      <label class="gap-3">
        <input type="checkbox">
        CDs, DVDs and iPods
      </label>
    </div>
  </div>
</fieldset>
```

### Radio Group Fieldset
```html
<fieldset class="fieldset">
  <legend>Subscription plan</legend>
  <p>Yearly and lifetime plans offer significant savings.</p>
  
  <div role="radiogroup" class="grid gap-3">
    <div class="field">
      <label class="gap-3">
        <input type="radio" name="subscription-plan" checked>
        Monthly ($9.99/month)
      </label>
    </div>
    <div class="field">
      <label class="gap-3">
        <input type="radio" name="subscription-plan">
        Yearly ($99.99/year)
      </label>
    </div>
    <div class="field">
      <label class="gap-3">
        <input type="radio" name="subscription-plan">
        Lifetime ($299.99)
      </label>
    </div>
  </div>
</fieldset>
```

## HTML Structure Guidelines

### Field Structure
```html
<div role="group" class="field">
  <label><!-- Label text --></label>
  <input> <!-- or select, textarea, etc. -->
  <p><!-- Optional helper text --></p>
  <p role="alert"><!-- Optional error message --></p>
</div>
```

### Fieldset Structure
```html
<fieldset class="fieldset">
  <legend><!-- Group title --></legend>
  <p><!-- Optional description --></p>
  <!-- Multiple .field containers or form controls -->
</fieldset>
```

### Horizontal Field Structure
```html
<div class="field" data-orientation="horizontal">
  <input class="input self-start">
  <section>
    <label><!-- Label text --></label>
    <p><!-- Description text --></p>
  </section>
</div>
```

## Layout Patterns

### Vertical Fields (Default)
```html
<div class="grid gap-6">
  <!-- Multiple field containers -->
</div>
```

### Horizontal Form Layout
```html
<div class="grid grid-cols-2 gap-4">
  <!-- Side-by-side fields -->
</div>
```

### Mixed Layout Grid
```html
<div class="grid grid-cols-3 gap-4">
  <!-- Three-column layout for related fields -->
</div>
```

## States

### Default State
Normal field appearance with proper spacing and typography

### Focus State  
Enhanced focus indicators for form controls

### Error State
Red styling when `aria-invalid="true"` is present

### Disabled State
Reduced opacity and interaction for disabled controls

## Form Integration

### Automatic Styling
Fields inherit proper styling when used within forms or fieldsets:

```html
<form class="space-y-6">
  <!-- Fields automatically styled -->
</form>
```

### Validation Integration
```html
<div role="group" class="field">
  <label for="email">Email</label>
  <input id="email" type="email" aria-invalid="true" aria-describedby="email-error" required>
  <p id="email-error" role="alert">Please enter a valid email address</p>
</div>
```

## Accessibility Features

- **Semantic HTML**: Proper use of fieldset, legend, and labels
- **ARIA Support**: Proper roles and relationships
- **Screen Reader**: Clear announcement of field purpose and errors
- **Keyboard Navigation**: Full keyboard accessibility
- **Error Handling**: Proper error association and announcement
- **Focus Management**: Clear focus indicators and logical tab order

## Best Practices

1. Always associate labels with form controls using `for` and `id`
2. Use `aria-describedby` to connect helper text and errors
3. Use `fieldset` and `legend` for grouping related fields
4. Include `role="alert"` for error messages
5. Use `aria-invalid="true"` when fields have validation errors
6. Provide helpful placeholder text and descriptions
7. Test with screen readers and keyboard navigation
8. Consider horizontal layout for simple checkbox/radio fields with descriptions

## Layout Orientation

### Vertical Layout (Default)
Labels appear above form controls with helper text below:

```html
<div role="group" class="field">
  <label>Label</label>
  <input>
  <p>Helper text</p>
</div>
```

### Horizontal Layout
Form control appears beside label and description:

```html
<div class="field" data-orientation="horizontal">
  <input class="input self-start">
  <section>
    <label>Label</label>
    <p>Description</p>
  </section>
</div>
```

## JavaScript Integration

### Form Validation
```javascript
// Check field validity
function validateField(fieldId) {
  const input = document.getElementById(fieldId);
  const isValid = input.checkValidity();
  
  input.setAttribute('aria-invalid', !isValid);
  
  const errorElement = document.getElementById(`${fieldId}-error`);
  if (errorElement) {
    errorElement.style.display = isValid ? 'none' : 'block';
  }
  
  return isValid;
}
```

### HTMX Integration
```html
<div role="group" class="field">
  <label for="username">Username</label>
  <input id="username" 
         type="text" 
         hx-post="/validate-username" 
         hx-trigger="blur"
         hx-target="#username-feedback"
         aria-describedby="username-feedback">
  <div id="username-feedback"></div>
</div>
```

### Dynamic Field Updates
```javascript
// Update helper text dynamically
function updateHelperText(fieldId, message) {
  const helperId = `${fieldId}-help`;
  const helperElement = document.getElementById(helperId);
  
  if (helperElement) {
    helperElement.textContent = message;
  }
}
```

## Related Components

- [Input](./input.md) - Text input fields
- [Label](./label.md) - Proper labeling
- [Checkbox](./checkbox.md) - Multiple selections
- [Radio Group](./radio-group.md) - Mutually exclusive options
- [Select](./select.md) - Dropdown selections
- [Textarea](./textarea.md) - Multi-line text input
- [Switch](./switch.md) - Binary toggles
- [Button](./button.md) - Form actions
---
## breadcrumb

# Breadcrumb Component

Displays the path to the current resource using a hierarchy of links.

## Important Note

**There is no dedicated Breadcrumb component in Basecoat.** Instead, use the pattern shown below with semantic HTML and Tailwind utilities.

## Basic Usage

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Components</a>
  </li>
</ol>
```

## CSS Classes

### Primary Classes
No dedicated breadcrumb classes - uses standard semantic HTML with utility classes.

### Tailwind Utilities Used
- `text-muted-foreground` - Muted text color for overall breadcrumb
- `flex flex-wrap` - Flexible, wrapping layout
- `items-center` - Vertical centering
- `gap-1.5` - Small spacing between elements
- `text-sm` - Small text size
- `break-words` - Word breaking for long links
- `sm:gap-2.5` - Responsive spacing
- `inline-flex` - Inline flex layout for links
- `hover:text-foreground` - Hover color change
- `transition-colors` - Smooth color transitions
- `size-3.5` - Icon sizing
- `text-foreground` - Full color for current page
- `font-normal` - Normal font weight

## HTML Structure

### Basic Breadcrumb
```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Breadcrumb item -->
  <li class="inline-flex items-center gap-1.5">
    <a href="/path">Link Text</a>
  </li>
  
  <!-- Separator -->
  <li>
    <!-- Chevron right icon -->
  </li>
  
  <!-- Current page (no link) -->
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Current Page</span>
  </li>
</ol>
```

### Separator Icon
Use a chevron-right icon between breadcrumb items:

```html
<li>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
    <path d="m9 18 6-6-6-6" />
  </svg>
</li>
```

## Examples

### Simple Breadcrumb
```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="/" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="/components" class="hover:text-foreground transition-colors">Components</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Breadcrumb</span>
  </li>
</ol>
```

### Breadcrumb with Dropdown Menu
For collapsed or overflow items, use the dropdown menu pattern:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <div id="breadcrumb-menu" class="dropdown-menu">
      <button type="button" id="breadcrumb-menu-trigger" aria-haspopup="menu" aria-controls="breadcrumb-menu-menu" aria-expanded="false" class="flex size-9 items-center justify-center h-4 w-4 hover:text-foreground cursor-pointer">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </button>
      <div id="breadcrumb-menu-popover" data-popover aria-hidden="true">
        <div role="menu" id="breadcrumb-menu-menu" aria-labelledby="breadcrumb-menu-trigger">
          <nav role="menu">
            <button type="button" role="menuitem">Documentation</button>
            <button type="button" role="menuitem">Themes</button>
            <button type="button" role="menuitem">GitHub</button>
          </nav>
        </div>
      </div>
    </div>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <a href="#" class="hover:text-foreground transition-colors">Components</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Breadcrumb</span>
  </li>
</ol>
```

### Breadcrumb with Icons
Add icons to breadcrumb items for enhanced visual hierarchy:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <li class="inline-flex items-center gap-1.5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
      <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
      <polyline points="9,22 9,12 15,12 15,22" />
    </svg>
    <a href="/" class="hover:text-foreground transition-colors">Home</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
      <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z" />
      <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z" />
    </svg>
    <a href="/docs" class="hover:text-foreground transition-colors">Documentation</a>
  </li>
  <li>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5"><path d="m9 18 6-6-6-6" /></svg>
  </li>
  <li class="inline-flex items-center gap-1.5">
    <span class="text-foreground font-normal">Current Page</span>
  </li>
</ol>
```

## Navigation Patterns

### Mobile Responsive
The breadcrumb automatically wraps on smaller screens:

```html
<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Items will wrap naturally on mobile -->
</ol>
```

### Condensed for Mobile
For mobile, consider showing only the parent and current page:

```html
<!-- Desktop: Full breadcrumb -->
<ol class="hidden sm:flex text-muted-foreground flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
  <!-- Full breadcrumb trail -->
</ol>

<!-- Mobile: Condensed breadcrumb -->
<ol class="sm:hidden text-muted-foreground flex items-center gap-1.5 text-sm">
  <li class="inline-flex items-center gap-1.5">
    <a href="/parent" class="hover:text-foreground transition-colors">← Back</a>
  </li>
</ol>
```

## Accessibility Features

- **Semantic HTML**: Uses `<ol>` for proper list semantics
- **Navigation Landmark**: Consider wrapping in `<nav>` with `aria-label`
- **Current Page**: Non-interactive element for current location
- **Keyboard Navigation**: All links are keyboard accessible
- **Screen Reader Support**: Proper list and link semantics

### Enhanced Accessibility
```html
<nav aria-label="Breadcrumb">
  <ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
    <li class="inline-flex items-center gap-1.5">
      <a href="/" class="hover:text-foreground transition-colors">Home</a>
    </li>
    <li>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5" aria-hidden="true">
        <path d="m9 18 6-6-6-6" />
      </svg>
    </li>
    <li class="inline-flex items-center gap-1.5">
      <span class="text-foreground font-normal" aria-current="page">Current Page</span>
    </li>
  </ol>
</nav>
```

## JavaScript Integration

### Dynamic Breadcrumbs
```javascript
function generateBreadcrumb(path) {
  const breadcrumbContainer = document.getElementById('breadcrumb');
  const pathSegments = path.split('/').filter(segment => segment);
  
  let breadcrumbHTML = '<ol class="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">';
  
  // Home link
  breadcrumbHTML += `
    <li class="inline-flex items-center gap-1.5">
      <a href="/" class="hover:text-foreground transition-colors">Home</a>
    </li>
  `;
  
  pathSegments.forEach((segment, index) => {
    const isLast = index === pathSegments.length - 1;
    const href = '/' + pathSegments.slice(0, index + 1).join('/');
    const title = segment.charAt(0).toUpperCase() + segment.slice(1);
    
    // Separator
    breadcrumbHTML += `
      <li>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-3.5">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </li>
    `;
    
    // Breadcrumb item
    if (isLast) {
      breadcrumbHTML += `
        <li class="inline-flex items-center gap-1.5">
          <span class="text-foreground font-normal" aria-current="page">${title}</span>
        </li>
      `;
    } else {
      breadcrumbHTML += `
        <li class="inline-flex items-center gap-1.5">
          <a href="${href}" class="hover:text-foreground transition-colors">${title}</a>
        </li>
      `;
    }
  });
  
  breadcrumbHTML += '</ol>';
  breadcrumbContainer.innerHTML = breadcrumbHTML;
}

// Usage
generateBreadcrumb(window.location.pathname);
```

### HTMX Integration
```html
<nav aria-label="Breadcrumb" hx-get="/breadcrumb" hx-trigger="path-change">
  <!-- Breadcrumb content loaded dynamically -->
</nav>
```

## Best Practices

1. **Start with Home**: Always begin breadcrumbs with the home/root page
2. **Current Page**: Don't link the current page, use a `<span>` instead
3. **Meaningful Labels**: Use clear, descriptive text for each level
4. **Reasonable Length**: Limit to 5-7 levels to avoid overcrowding
5. **Mobile Consideration**: Use responsive design or mobile-specific patterns
6. **Accessibility**: Include proper ARIA labels and semantic HTML
7. **Visual Hierarchy**: Use consistent styling and separators
8. **Dropdown for Overflow**: Use dropdown menus for collapsed intermediate levels

## Integration with Routing

### React Router
```jsx
import { useLocation } from 'react-router-dom';

function Breadcrumb() {
  const location = useLocation();
  const pathnames = location.pathname.split('/').filter(x => x);
  
  return (
    <nav aria-label="Breadcrumb">
      <ol className="text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5">
        <li className="inline-flex items-center gap-1.5">
          <a href="/" className="hover:text-foreground transition-colors">Home</a>
        </li>
        {pathnames.map((pathname, index) => {
          const routeTo = `/${pathnames.slice(0, index + 1).join('/')}`;
          const isLast = index === pathnames.length - 1;
          
          return (
            <React.Fragment key={routeTo}>
              <li>
                <svg className="size-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor">
                  <path d="m9 18 6-6-6-6" />
                </svg>
              </li>
              <li className="inline-flex items-center gap-1.5">
                {isLast ? (
                  <span className="text-foreground font-normal" aria-current="page">
                    {pathname}
                  </span>
                ) : (
                  <a href={routeTo} className="hover:text-foreground transition-colors">
                    {pathname}
                  </a>
                )}
              </li>
            </React.Fragment>
          );
        })}
      </ol>
    </nav>
  );
}
```

## Related Components

- [Dropdown Menu](./dropdown-menu.md) - For collapsed breadcrumb items
- [Button](./button.md) - For interactive elements
- [Link](./link.md) - For navigation links
- [Navigation](./navigation.md) - For main site navigation
---
## dialog

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
---
## dropdown-menu

# Dropdown Menu Component

A menu that appears when a trigger element is activated, containing a list of choices that users can select from.

## Basic Usage

```html
<div id="demo-dropdown-menu" class="dropdown-menu">
  <button type="button" id="demo-dropdown-menu-trigger" aria-haspopup="menu" aria-controls="demo-dropdown-menu-menu" aria-expanded="false" class="btn-outline">
    Open
  </button>
  <div id="demo-dropdown-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="demo-dropdown-menu-menu" aria-labelledby="demo-dropdown-menu-trigger">
      <div role="menuitem">Profile</div>
      <div role="menuitem">Settings</div>
      <hr role="separator" />
      <div role="menuitem">Logout</div>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/dropdown-menu.min.js" defer></script>
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
- **`dropdown-menu`** - Applied to the wrapper container
- **`min-w-56`** - Minimum width for the popover (commonly used)

### Supporting Classes
- Button classes for triggers (`btn-outline`, `btn`, etc.)
- Popover positioning classes (`data-side`, `data-align`)
- Text styling classes for shortcuts and descriptions

### Tailwind Utilities Used
- `ml-auto` - Push keyboard shortcuts to the right
- `text-xs tracking-widest` - Styling for keyboard shortcuts
- `text-muted-foreground` - Muted text color
- `invisible` - Hide unchecked checkboxes
- `group-aria-checked:visible` - Show when checked

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "dropdown-menu" | Yes |
| `id` | string | Unique identifier | Yes |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `id` | string | Referenced by aria-labelledby | Yes |
| `aria-haspopup` | string | Must be "menu" | Yes |
| `aria-controls` | string | References menu ID | Yes |
| `aria-expanded` | boolean | Tracks open/closed state | Yes |

### Popover Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-popover` | boolean | Marks as popover element | Yes |
| `aria-hidden` | boolean | Controls visibility for screen readers | Yes |
| `data-side` | string | Position: "top", "right", "bottom", "left" | No |
| `data-align` | string | Alignment: "start", "center", "end" | No |

### Menu Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "menu" | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `aria-labelledby` | string | References trigger button ID | Yes |

### Menu Item Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | "menuitem", "menuitemcheckbox", or "menuitemradio" | Yes |
| `aria-checked` | boolean | For checkbox/radio items | For checkboxes/radios |
| `aria-disabled` | boolean | Marks item as disabled | No |

## HTML Structure

```html
<div class="dropdown-menu" id="dropdown-id">
  <!-- Trigger button -->
  <button type="button" 
          id="trigger-id" 
          aria-haspopup="menu" 
          aria-controls="menu-id" 
          aria-expanded="false" 
          class="btn-outline">
    Trigger Text
  </button>
  
  <!-- Popover container -->
  <div id="popover-id" 
       data-popover 
       aria-hidden="true" 
       class="min-w-56">
    
    <!-- Menu container -->
    <div role="menu" 
         id="menu-id" 
         aria-labelledby="trigger-id">
      
      <!-- Menu items -->
      <div role="menuitem">Menu Item</div>
      
      <!-- Separator -->
      <hr role="separator" />
      
      <!-- Grouped items -->
      <div role="group" aria-labelledby="group-heading">
        <div role="heading" id="group-heading">Group Title</div>
        <div role="menuitem">Grouped Item</div>
      </div>
      
    </div>
  </div>
</div>
```

## Examples

### Basic Menu
```html
<div id="basic-menu" class="dropdown-menu">
  <button type="button" id="basic-menu-trigger" aria-haspopup="menu" aria-controls="basic-menu-menu" aria-expanded="false" class="btn-outline">
    Actions
  </button>
  <div id="basic-menu-popover" data-popover aria-hidden="true" class="min-w-40">
    <div role="menu" id="basic-menu-menu" aria-labelledby="basic-menu-trigger">
      <div role="menuitem">Edit</div>
      <div role="menuitem">Copy</div>
      <div role="menuitem">Move</div>
      <hr role="separator" />
      <div role="menuitem" aria-disabled="true">Delete</div>
    </div>
  </div>
</div>
```

### Menu with Icons and Shortcuts
```html
<div id="icon-menu" class="dropdown-menu">
  <button type="button" id="icon-menu-trigger" aria-haspopup="menu" aria-controls="icon-menu-menu" aria-expanded="false" class="btn-outline">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="1" />
      <circle cx="19" cy="12" r="1" />
      <circle cx="5" cy="12" r="1" />
    </svg>
  </button>
  <div id="icon-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="icon-menu-menu" aria-labelledby="icon-menu-trigger">
      <div role="group" aria-labelledby="account-options">
        <div role="heading" id="account-options">My Account</div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
          Profile
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘P</span>
        </div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="20" height="14" x="2" y="5" rx="2" />
            <line x1="2" x2="22" y1="10" y2="10" />
          </svg>
          Billing
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘B</span>
        </div>
        <div role="menuitem">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
            <circle cx="12" cy="12" r="3" />
          </svg>
          Settings
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘S</span>
        </div>
      </div>
      <hr role="separator" />
      <div role="menuitem">GitHub</div>
      <div role="menuitem">Support</div>
      <hr role="separator" />
      <div role="menuitem">
        Logout
        <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘Q</span>
      </div>
    </div>
  </div>
</div>
```

### Menu with Checkboxes
```html
<div id="checkbox-menu" class="dropdown-menu">
  <button type="button" id="checkbox-menu-trigger" aria-haspopup="menu" aria-controls="checkbox-menu-menu" aria-expanded="false" class="btn-outline">
    View Options
  </button>
  <div id="checkbox-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="checkbox-menu-menu" aria-labelledby="checkbox-menu-trigger">
      <div role="group" aria-labelledby="appearance-options">
        <div role="heading" id="appearance-options">Appearance</div>
        <div role="menuitemcheckbox" aria-checked="true" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Status Bar
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘S</span>
        </div>
        <div role="menuitemcheckbox" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Activity Bar
          <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘B</span>
        </div>
        <div role="menuitemcheckbox" aria-checked="false" class="group" aria-disabled="true">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <path d="M20 6 9 17l-5-5" />
          </svg>
          Panel (Disabled)
        </div>
      </div>
    </div>
  </div>
</div>
```

### Menu with Radio Buttons
```html
<div id="radio-menu" class="dropdown-menu">
  <button type="button" id="radio-menu-trigger" aria-haspopup="menu" aria-controls="radio-menu-menu" aria-expanded="false" class="btn-outline">
    Theme: Light
  </button>
  <div id="radio-menu-popover" data-popover aria-hidden="true" class="min-w-48">
    <div role="menu" id="radio-menu-menu" aria-labelledby="radio-menu-trigger">
      <div role="group" aria-labelledby="theme-options">
        <div role="heading" id="theme-options">Theme</div>
        <div role="menuitemradio" aria-checked="true" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Light
        </div>
        <div role="menuitemradio" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Dark
        </div>
        <div role="menuitemradio" aria-checked="false" class="group">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="invisible group-aria-checked:visible" aria-hidden="true">
            <circle cx="12" cy="12" r="3" />
          </svg>
          Auto
        </div>
      </div>
    </div>
  </div>
</div>
```

### Positioned Menu
```html
<div id="positioned-menu" class="dropdown-menu">
  <button type="button" id="positioned-menu-trigger" aria-haspopup="menu" aria-controls="positioned-menu-menu" aria-expanded="false" class="btn-outline">
    Top Menu
  </button>
  <div id="positioned-menu-popover" data-popover aria-hidden="true" class="min-w-48" data-side="top" data-align="end">
    <div role="menu" id="positioned-menu-menu" aria-labelledby="positioned-menu-trigger">
      <div role="menuitem">Item 1</div>
      <div role="menuitem">Item 2</div>
      <div role="menuitem">Item 3</div>
    </div>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const dropdownMenu = document.getElementById('my-dropdown');

// Listen for component initialization
dropdownMenu.addEventListener('basecoat:initialized', (e) => {
  console.log('Dropdown menu initialized');
});

// Listen for popover open/close
document.addEventListener('basecoat:popover', (e) => {
  console.log('A popover was opened, others will close');
});
```

### Menu Item Click Handling
```javascript
// Handle menu item clicks
document.querySelectorAll('[role="menuitem"]').forEach(item => {
  item.addEventListener('click', (e) => {
    const text = e.target.textContent.trim();
    console.log('Menu item clicked:', text);
    
    // Close the dropdown menu
    const dropdown = e.target.closest('.dropdown-menu');
    const trigger = dropdown.querySelector('[aria-haspopup="menu"]');
    trigger.setAttribute('aria-expanded', 'false');
    
    // Hide popover
    const popover = dropdown.querySelector('[data-popover]');
    popover.setAttribute('aria-hidden', 'true');
  });
});
```

### Checkbox/Radio State Management
```javascript
// Handle checkbox menu items
document.querySelectorAll('[role="menuitemcheckbox"]').forEach(item => {
  item.addEventListener('click', (e) => {
    const currentState = item.getAttribute('aria-checked') === 'true';
    item.setAttribute('aria-checked', (!currentState).toString());
    
    console.log(`${item.textContent.trim()}: ${!currentState ? 'checked' : 'unchecked'}`);
  });
});

// Handle radio button menu items
document.querySelectorAll('[role="menuitemradio"]').forEach(item => {
  item.addEventListener('click', (e) => {
    // Uncheck all radio items in the same group
    const group = item.closest('[role="group"]');
    if (group) {
      group.querySelectorAll('[role="menuitemradio"]').forEach(radio => {
        radio.setAttribute('aria-checked', 'false');
      });
    }
    
    // Check the clicked item
    item.setAttribute('aria-checked', 'true');
    
    console.log(`Selected: ${item.textContent.trim()}`);
  });
});
```

## Accessibility Features

- **Keyboard Navigation**: Arrow keys, Enter, Escape
- **Focus Management**: Proper focus handling when opening/closing
- **Screen Reader Support**: Full ARIA menu roles and properties
- **State Communication**: Checkbox/radio states properly announced
- **Group Structure**: Menu item groups with headings

### Enhanced Accessibility
```html
<div id="accessible-menu" class="dropdown-menu">
  <button type="button" 
          id="accessible-menu-trigger" 
          aria-haspopup="menu" 
          aria-controls="accessible-menu-menu" 
          aria-expanded="false" 
          aria-label="Open user account menu"
          class="btn-outline">
    Account
  </button>
  <div id="accessible-menu-popover" data-popover aria-hidden="true" class="min-w-56">
    <div role="menu" id="accessible-menu-menu" aria-labelledby="accessible-menu-trigger">
      <div role="group" aria-labelledby="account-heading">
        <div role="heading" id="account-heading" aria-level="3">Account Settings</div>
        <div role="menuitem" tabindex="-1">
          Profile Settings
        </div>
        <div role="menuitem" tabindex="-1" aria-describedby="billing-desc">
          Billing
          <div id="billing-desc" class="sr-only">Manage your subscription and payments</div>
        </div>
      </div>
      <hr role="separator" aria-hidden="true" />
      <div role="menuitem" tabindex="-1">
        Sign Out
      </div>
    </div>
  </div>
</div>
```

## Positioning

### Popover Positioning
Use `data-side` and `data-align` attributes on the popover element:

```html
<!-- Position above trigger, aligned to end -->
<div data-popover data-side="top" data-align="end">
  <!-- Menu content -->
</div>

<!-- Position to the right, center aligned -->
<div data-popover data-side="right" data-align="center">
  <!-- Menu content -->
</div>

<!-- Position below (default), aligned to start -->
<div data-popover data-side="bottom" data-align="start">
  <!-- Menu content -->
</div>
```

### Side Options
- `top` - Above the trigger
- `right` - To the right of the trigger  
- `bottom` - Below the trigger (default)
- `left` - To the left of the trigger

### Alignment Options
- `start` - Align to the start edge
- `center` - Center align
- `end` - Align to the end edge

## Best Practices

1. **Use Semantic Roles**: Always use proper ARIA menu roles
2. **Group Related Items**: Use groups with headings for organization
3. **Keyboard Shortcuts**: Include keyboard shortcuts where applicable
4. **Clear Icons**: Use consistent, meaningful icons
5. **Logical Order**: Organize items logically
6. **Disable Appropriately**: Use `aria-disabled` for unavailable actions
7. **State Communication**: Properly manage checkbox/radio states
8. **Focus Management**: Ensure proper keyboard navigation

## Common Patterns

### Context Menu Pattern
```javascript
// Right-click context menu
document.addEventListener('contextmenu', (e) => {
  e.preventDefault();
  
  const contextMenu = createContextMenu([
    { label: 'Copy', shortcut: '⌘C', action: () => copy() },
    { label: 'Paste', shortcut: '⌘V', action: () => paste() },
    { separator: true },
    { label: 'Delete', shortcut: 'Del', action: () => delete() }
  ]);
  
  showMenuAtPosition(contextMenu, e.clientX, e.clientY);
});
```

### User Account Menu Pattern
```html
<div id="user-menu" class="dropdown-menu">
  <button type="button" id="user-menu-trigger" aria-haspopup="menu" aria-controls="user-menu-menu" aria-expanded="false" class="btn-ghost">
    <img src="/avatar.jpg" alt="User Avatar" class="size-8 rounded-full">
    John Doe
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m6 9 6 6 6-6" />
    </svg>
  </button>
  <div id="user-menu-popover" data-popover aria-hidden="true" class="min-w-56" data-align="end">
    <div role="menu" id="user-menu-menu" aria-labelledby="user-menu-trigger">
      <div role="group">
        <div class="px-3 py-2 text-sm">
          <div class="font-medium">John Doe</div>
          <div class="text-muted-foreground">john@example.com</div>
        </div>
      </div>
      <hr role="separator" />
      <div role="menuitem">Dashboard</div>
      <div role="menuitem">Settings</div>
      <div role="menuitem">Billing</div>
      <hr role="separator" />
      <div role="menuitem">Sign Out</div>
    </div>
  </div>
</div>
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function DropdownMenu({ trigger, items, onSelect }) {
  const [isOpen, setIsOpen] = useState(false);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="dropdown-menu">
      <button 
        type="button"
        aria-haspopup="menu"
        aria-expanded={isOpen}
        className="btn-outline"
      >
        {trigger}
      </button>
      <div data-popover aria-hidden={!isOpen} className="min-w-56">
        <div role="menu">
          {items.map((item, index) => (
            item.separator ? (
              <hr key={index} role="separator" />
            ) : (
              <div 
                key={index}
                role="menuitem" 
                onClick={() => onSelect(item)}
              >
                {item.label}
                {item.shortcut && (
                  <span className="text-muted-foreground ml-auto text-xs tracking-widest">
                    {item.shortcut}
                  </span>
                )}
              </div>
            )
          ))}
        </div>
      </div>
    </div>
  );
}
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% call dropdown_menu(
  id="user-menu",
  trigger="Account",
  trigger_attrs={"class": "btn-outline"},
  popover_attrs={"class": "min-w-56", "data-align": "end"}
) %}
  <div role="group" aria-labelledby="account-options">
    <div role="heading" id="account-options">My Account</div>
    <div role="menuitem">
      Profile
      <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘P</span>
    </div>
    <div role="menuitem">
      Settings
      <span class="text-muted-foreground ml-auto text-xs tracking-widest">⌘S</span>
    </div>
  </div>
  <hr role="separator">
  <div role="menuitem">
    Logout
    <span class="text-muted-foreground ml-auto text-xs tracking-widest">⇧⌘Q</span>
  </div>
{% endcall %}
```

### Advanced Configuration
```jinja2
{% set menu_items = [
  {"type": "group", "heading": "Account", "items": [
    {"label": "Profile", "shortcut": "⇧⌘P"},
    {"label": "Settings", "shortcut": "⌘S"}
  ]},
  {"type": "separator"},
  {"label": "Logout", "shortcut": "⇧⌘Q"}
] %}

{{ dropdown_menu(
  id="dynamic-menu",
  trigger="Options",
  items=menu_items,
  trigger_attrs={"class": "btn-ghost"},
  popover_attrs={"class": "min-w-48"}
) }}
```

## Related Components

- [Button](./button.md) - For menu triggers
- [Popover](./popover.md) - Underlying positioning system
- [Dialog](./dialog.md) - For modal interactions
- [Select](./select.md) - For form-based selection
---
## pagination

# Pagination Component

Pagination with page navigation, next and previous links.

## Important Note

**There is no dedicated pagination component in Basecoat.** Instead, use the pattern shown below with Basecoat button classes and semantic HTML.

## Basic Usage

```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="#" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="#" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <a href="#" class="btn-icon-outline">2</a>
    </li>
    <li>
      <a href="#" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="#" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

## CSS Classes

### Primary Classes
Uses existing Basecoat button classes:
- **`btn-ghost`** - For Previous/Next links
- **`btn-icon-ghost`** - For inactive page numbers
- **`btn-icon-outline`** - For current/active page

### Tailwind Utilities Used
- `mx-auto` - Center the pagination horizontally
- `flex w-full justify-center` - Flexible centering layout
- `flex flex-row items-center` - Horizontal list layout
- `gap-1` - Small spacing between pagination items
- `size-9` - Size for ellipsis container
- `size-4 shrink-0` - Icon sizing

## HTML Structure

```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Previous link -->
    <li>
      <a href="/prev-page" class="btn-ghost">Previous</a>
    </li>
    
    <!-- Page numbers -->
    <li>
      <a href="/page-1" class="btn-icon-ghost">1</a>
    </li>
    
    <!-- Current page (highlighted) -->
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    
    <!-- More page numbers -->
    <li>
      <a href="/page-3" class="btn-icon-ghost">3</a>
    </li>
    
    <!-- Ellipsis for truncated pages -->
    <li>
      <div class="size-9 flex items-center justify-center">
        <!-- Ellipsis icon -->
      </div>
    </li>
    
    <!-- Next link -->
    <li>
      <a href="/next-page" class="btn-ghost">Next</a>
    </li>
  </ul>
</nav>
```

## Examples

### Simple Pagination
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Pagination with Ellipsis
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/4" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <div class="size-9 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </div>
    </li>
    <li>
      <a href="/page/4" class="btn-icon-ghost">4</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">5</span>
    </li>
    <li>
      <a href="/page/6" class="btn-icon-ghost">6</a>
    </li>
    <li>
      <div class="size-9 flex items-center justify-center">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
          <circle cx="12" cy="12" r="1" />
          <circle cx="19" cy="12" r="1" />
          <circle cx="5" cy="12" r="1" />
        </svg>
      </div>
    </li>
    <li>
      <a href="/page/20" class="btn-icon-ghost">20</a>
    </li>
    <li>
      <a href="/page/6" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Disabled States
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Disabled Previous (first page) -->
    <li>
      <span class="btn-ghost opacity-50 cursor-not-allowed">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </span>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">1</span>
    </li>
    <li>
      <a href="/page/2" class="btn-icon-ghost">2</a>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/2" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Mobile-Friendly Pagination
```html
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    <!-- Mobile: Show only Previous/Next on small screens -->
    <li class="sm:hidden">
      <a href="/page/4" class="btn-ghost">Previous</a>
    </li>
    <li class="sm:hidden">
      <span class="btn-outline px-4">Page 5 of 20</span>
    </li>
    <li class="sm:hidden">
      <a href="/page/6" class="btn-ghost">Next</a>
    </li>
    
    <!-- Desktop: Full pagination -->
    <li class="hidden sm:inline">
      <a href="/page/4" class="btn-ghost">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
        Previous
      </a>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li class="hidden sm:inline">
      <div class="size-9 flex items-center justify-center">…</div>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/4" class="btn-icon-ghost">4</a>
    </li>
    <li class="hidden sm:inline">
      <span class="btn-icon-outline" aria-current="page">5</span>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/6" class="btn-icon-ghost">6</a>
    </li>
    <li class="hidden sm:inline">
      <div class="size-9 flex items-center justify-center">…</div>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/20" class="btn-icon-ghost">20</a>
    </li>
    <li class="hidden sm:inline">
      <a href="/page/6" class="btn-ghost">
        Next
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

### Compact Pagination
```html
<nav role="navigation" aria-label="pagination" class="flex justify-center">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-icon-ghost" title="Previous page">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m15 18-6-6 6-6" />
        </svg>
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost" title="Next page">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="m9 18 6-6-6-6" />
        </svg>
      </a>
    </li>
  </ul>
</nav>
```

## Icon SVGs

### Previous Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="m15 18-6-6 6-6" />
</svg>
```

### Next Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
  <path d="m9 18 6-6-6-6" />
</svg>
```

### Ellipsis Icon
```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="size-4 shrink-0">
  <circle cx="12" cy="12" r="1" />
  <circle cx="19" cy="12" r="1" />
  <circle cx="5" cy="12" r="1" />
</svg>
```

## Accessibility Features

- **Semantic HTML**: Uses `<nav>` and `<ul>` for proper navigation structure
- **Navigation Landmark**: `role="navigation"` identifies the pagination
- **ARIA Labels**: `aria-label="pagination"` describes the navigation purpose
- **Current Page**: `aria-current="page"` marks the active page
- **Keyboard Navigation**: All links are keyboard accessible
- **Screen Reader Support**: Proper list and link semantics

### Enhanced Accessibility
```html
<nav role="navigation" aria-label="Pagination Navigation">
  <ul class="flex flex-row items-center gap-1">
    <li>
      <a href="/page/1" class="btn-ghost" aria-label="Go to previous page">
        <svg aria-hidden="true"><!-- Previous icon --></svg>
        Previous
      </a>
    </li>
    <li>
      <a href="/page/1" class="btn-icon-ghost" aria-label="Go to page 1">1</a>
    </li>
    <li>
      <span class="btn-icon-outline" aria-current="page" aria-label="Current page, page 2">2</span>
    </li>
    <li>
      <a href="/page/3" class="btn-icon-ghost" aria-label="Go to page 3">3</a>
    </li>
    <li>
      <a href="/page/3" class="btn-ghost" aria-label="Go to next page">
        Next
        <svg aria-hidden="true"><!-- Next icon --></svg>
      </a>
    </li>
  </ul>
</nav>
```

## JavaScript Integration

### Dynamic Pagination
```javascript
function generatePagination(currentPage, totalPages, baseUrl) {
  const paginationContainer = document.getElementById('pagination');
  const maxVisiblePages = 5;
  
  let html = '<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">';
  html += '<ul class="flex flex-row items-center gap-1">';
  
  // Previous button
  if (currentPage > 1) {
    html += `
      <li>
        <a href="${baseUrl}?page=${currentPage - 1}" class="btn-ghost">
          <svg><!-- Previous icon --></svg>
          Previous
        </a>
      </li>
    `;
  } else {
    html += `
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">
          <svg><!-- Previous icon --></svg>
          Previous
        </span>
      </li>
    `;
  }
  
  // Page numbers
  const startPage = Math.max(1, currentPage - Math.floor(maxVisiblePages / 2));
  const endPage = Math.min(totalPages, startPage + maxVisiblePages - 1);
  
  if (startPage > 1) {
    html += `<li><a href="${baseUrl}?page=1" class="btn-icon-ghost">1</a></li>`;
    if (startPage > 2) {
      html += '<li><div class="size-9 flex items-center justify-center">…</div></li>';
    }
  }
  
  for (let page = startPage; page <= endPage; page++) {
    if (page === currentPage) {
      html += `<li><span class="btn-icon-outline" aria-current="page">${page}</span></li>`;
    } else {
      html += `<li><a href="${baseUrl}?page=${page}" class="btn-icon-ghost">${page}</a></li>`;
    }
  }
  
  if (endPage < totalPages) {
    if (endPage < totalPages - 1) {
      html += '<li><div class="size-9 flex items-center justify-center">…</div></li>';
    }
    html += `<li><a href="${baseUrl}?page=${totalPages}" class="btn-icon-ghost">${totalPages}</a></li>`;
  }
  
  // Next button
  if (currentPage < totalPages) {
    html += `
      <li>
        <a href="${baseUrl}?page=${currentPage + 1}" class="btn-ghost">
          Next
          <svg><!-- Next icon --></svg>
        </a>
      </li>
    `;
  } else {
    html += `
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">
          Next
          <svg><!-- Next icon --></svg>
        </span>
      </li>
    `;
  }
  
  html += '</ul></nav>';
  paginationContainer.innerHTML = html;
}

// Usage
generatePagination(5, 20, '/products');
```

### HTMX Integration
```html
<nav id="pagination" hx-get="/api/pagination" hx-trigger="page-change">
  <!-- Pagination content loaded dynamically -->
</nav>

<script>
// Custom event to trigger pagination updates
function changePage(page) {
  document.dispatchEvent(new CustomEvent('page-change', { detail: { page } }));
}
</script>
```

### Keyboard Navigation Enhancement
```javascript
document.addEventListener('keydown', function(e) {
  const currentPage = parseInt(document.querySelector('[aria-current="page"]').textContent);
  const totalPages = document.querySelectorAll('.btn-icon-ghost, .btn-icon-outline').length;
  
  if (e.key === 'ArrowLeft' && currentPage > 1) {
    window.location.href = `?page=${currentPage - 1}`;
  } else if (e.key === 'ArrowRight' && currentPage < totalPages) {
    window.location.href = `?page=${currentPage + 1}`;
  }
});
```

## Best Practices

1. **Provide Context**: Always include total page count and current position
2. **Limit Visible Pages**: Show 5-7 page numbers to avoid overcrowding
3. **Use Ellipsis**: Indicate skipped pages with ellipsis
4. **Disable When Appropriate**: Disable Previous on first page, Next on last page
5. **Mobile Optimization**: Simplify pagination for mobile devices
6. **Clear Labels**: Use descriptive aria-labels for screen readers
7. **Consistent Styling**: Follow button design patterns
8. **Fast Navigation**: Include first/last page links for large datasets

## URL Structure

### Query Parameters
```
/products?page=5
/search?q=term&page=3
/articles?category=tech&page=2&limit=10
```

### RESTful Paths
```
/products/page/5
/search/term/page/3
/articles/tech/page/2
```

## Integration Examples

### PHP/Laravel
```php
// Controller
$products = Product::paginate(10);

// Blade template
<nav role="navigation" aria-label="pagination" class="mx-auto flex w-full justify-center">
  <ul class="flex flex-row items-center gap-1">
    @if ($products->onFirstPage())
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">Previous</span>
      </li>
    @else
      <li>
        <a href="{{ $products->previousPageUrl() }}" class="btn-ghost">Previous</a>
      </li>
    @endif
    
    @foreach ($elements as $element)
      @if (is_array($element))
        @foreach ($element as $page => $url)
          @if ($page == $products->currentPage())
            <li>
              <span class="btn-icon-outline" aria-current="page">{{ $page }}</span>
            </li>
          @else
            <li>
              <a href="{{ $url }}" class="btn-icon-ghost">{{ $page }}</a>
            </li>
          @endif
        @endforeach
      @endif
    @endforeach
    
    @if ($products->hasMorePages())
      <li>
        <a href="{{ $products->nextPageUrl() }}" class="btn-ghost">Next</a>
      </li>
    @else
      <li>
        <span class="btn-ghost opacity-50 cursor-not-allowed">Next</span>
      </li>
    @endif
  </ul>
</nav>
```

## Related Components

- [Button](./button.md) - For pagination links and controls
- [Breadcrumb](./breadcrumb.md) - For hierarchical navigation
- [Table](./table.md) - Often used with pagination
- [Navigation](./navigation.md) - For primary site navigation
---
## popover

# Popover Component

Displays rich content in a portal, triggered by a button.

## Basic Usage

```html
<div id="demo-popover" class="popover">
  <button id="demo-popover-trigger" type="button" aria-expanded="false" aria-controls="demo-popover-popover" class="btn-outline">Open popover</button>
  <div id="demo-popover-popover" data-popover aria-hidden="true" class="w-80">
    <div class="grid gap-4">
      <header class="grid gap-1.5">
        <h4 class="leading-none font-medium">Dimensions</h4>
        <p class="text-muted-foreground text-sm">Set the dimensions for the layer.</p>
      </header>
      <form class="form grid gap-2">
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-width">Width</label>
          <input type="text" id="demo-popover-width" value="100%" class="col-span-2 h-8" autofocus />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-max-width">Max. width</label>
          <input type="text" id="demo-popover-max-width" value="300px" class="col-span-2 h-8" />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-height">Height</label>
          <input type="text" id="demo-popover-height" value="25px" class="col-span-2 h-8" />
        </div>
        <div class="grid grid-cols-3 items-center gap-4">
          <label for="demo-popover-max-height">Max. height</label>
          <input type="text" id="demo-popover-max-height" value="none" class="col-span-2 h-8" />
        </div>
      </form>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/popover.min.js" defer></script>
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
- **`popover`** - Applied to the main container
- **`data-popover`** - Applied to the popover content element

### Supporting Classes
- Button classes for triggers (`btn-outline`, `btn`, etc.)
- Width/height utility classes (`w-80`, `h-8`, etc.)
- Grid and spacing classes for content layout

### Tailwind Utilities Used
- `grid gap-4` - Grid layout with gap spacing
- `grid-cols-3` - Three-column grid layout
- `items-center` - Vertical alignment
- `col-span-2` - Span two grid columns
- `w-80` - Fixed width for popover
- `text-muted-foreground` - Muted text color
- `text-sm` - Small text size
- `leading-none` - No line height
- `font-medium` - Medium font weight

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "popover" | Yes |
| `id` | string | Unique identifier | Yes |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `id` | string | Referenced by aria-labelledby | Yes |
| `aria-expanded` | boolean | Tracks open/closed state | Yes |
| `aria-controls` | string | References popover content ID | Yes |
| `popovertarget` | string | Alternative native HTML attribute | Optional |

### Popover Content Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-popover` | boolean | Marks as popover content | Yes |
| `aria-hidden` | boolean | Controls visibility for screen readers | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `data-side` | string | Position: "top", "right", "bottom", "left" | No |
| `data-align` | string | Alignment: "start", "center", "end" | No |

## HTML Structure

```html
<div class="popover" id="popover-id">
  <!-- Trigger button -->
  <button type="button" 
          id="trigger-id" 
          aria-expanded="false" 
          aria-controls="popover-content-id" 
          class="btn-outline">
    Trigger Text
  </button>
  
  <!-- Popover content -->
  <div id="popover-content-id" 
       data-popover 
       aria-hidden="true" 
       class="w-80">
    <!-- Popover content -->
  </div>
</div>
```

## Examples

### Basic Information Popover
```html
<div id="info-popover" class="popover">
  <button id="info-popover-trigger" type="button" aria-expanded="false" aria-controls="info-popover-content" class="btn-outline">
    Show Info
  </button>
  <div id="info-popover-content" data-popover aria-hidden="true" class="w-64">
    <div class="grid gap-3">
      <h4 class="font-medium">Additional Information</h4>
      <p class="text-sm text-muted-foreground">
        This is some helpful information that appears in a popover.
      </p>
    </div>
  </div>
</div>
```

### Form Popover
```html
<div id="form-popover" class="popover">
  <button id="form-popover-trigger" type="button" aria-expanded="false" aria-controls="form-popover-content" class="btn-outline">
    Edit Settings
  </button>
  <div id="form-popover-content" data-popover aria-hidden="true" class="w-80">
    <div class="grid gap-4">
      <header class="grid gap-1.5">
        <h4 class="leading-none font-medium">Settings</h4>
        <p class="text-muted-foreground text-sm">Configure your preferences.</p>
      </header>
      <form class="form grid gap-3">
        <div class="grid gap-2">
          <label for="setting-name">Name</label>
          <input type="text" id="setting-name" placeholder="Enter name">
        </div>
        <div class="grid gap-2">
          <label for="setting-email">Email</label>
          <input type="email" id="setting-email" placeholder="Enter email">
        </div>
        <div class="flex gap-2">
          <button type="submit" class="btn">Save</button>
          <button type="button" class="btn-outline">Cancel</button>
        </div>
      </form>
    </div>
  </div>
</div>
```

### Positioned Popover
```html
<div id="positioned-popover" class="popover">
  <button id="positioned-popover-trigger" type="button" aria-expanded="false" aria-controls="positioned-popover-content" class="btn-outline">
    Top Popover
  </button>
  <div id="positioned-popover-content" data-popover aria-hidden="true" class="w-56" data-side="top" data-align="end">
    <div class="grid gap-3">
      <h4 class="font-medium">Positioned Above</h4>
      <p class="text-sm text-muted-foreground">
        This popover appears above the trigger button, aligned to the end.
      </p>
    </div>
  </div>
</div>
```

### Rich Content Popover
```html
<div id="rich-popover" class="popover">
  <button id="rich-popover-trigger" type="button" aria-expanded="false" aria-controls="rich-popover-content" class="btn-outline">
    Show Details
  </button>
  <div id="rich-popover-content" data-popover aria-hidden="true" class="w-72">
    <div class="grid gap-4">
      <div class="flex items-center gap-3">
        <div class="size-12 rounded-full bg-primary/10 flex items-center justify-center">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
            <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
            <circle cx="12" cy="7" r="4" />
          </svg>
        </div>
        <div class="grid gap-1">
          <h4 class="font-medium">John Doe</h4>
          <p class="text-sm text-muted-foreground">Software Developer</p>
        </div>
      </div>
      <div class="grid gap-2">
        <div class="flex items-center gap-2 text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="20" height="16" x="2" y="4" rx="2" />
            <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
          </svg>
          john.doe@example.com
        </div>
        <div class="flex items-center gap-2 text-sm">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z" />
          </svg>
          +1 (555) 123-4567
        </div>
      </div>
      <div class="flex gap-2">
        <button type="button" class="btn">Contact</button>
        <button type="button" class="btn-outline">View Profile</button>
      </div>
    </div>
  </div>
</div>
```

### List Popover
```html
<div id="list-popover" class="popover">
  <button id="list-popover-trigger" type="button" aria-expanded="false" aria-controls="list-popover-content" class="btn-outline">
    Options
  </button>
  <div id="list-popover-content" data-popover aria-hidden="true" class="w-48">
    <div class="grid gap-1">
      <h4 class="font-medium px-3 py-2 border-b">Quick Actions</h4>
      <div class="grid gap-1 p-1">
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M4 13.5V4a1 1 0 0 1 1-1h4.5" />
            <path d="M14 4h5a1 1 0 0 1 1 1v5" />
            <path d="M4 20v-5.5a1 1 0 0 1 1-1H10" />
            <path d="M20 10v10a1 1 0 0 1-1 1h-5" />
          </svg>
          Expand
        </button>
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
            <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
          </svg>
          Copy
        </button>
        <button type="button" class="btn-ghost justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M3 6h18" />
            <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
            <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
          </svg>
          Delete
        </button>
      </div>
    </div>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const popover = document.getElementById('my-popover');

// Listen for component initialization
popover.addEventListener('basecoat:initialized', (e) => {
  console.log('Popover initialized');
});

// Listen for popover open (dispatched on document)
document.addEventListener('basecoat:popover', (e) => {
  console.log('A popover was opened, others will close');
});
```

### Manual Control
```javascript
// Get popover elements
const trigger = document.getElementById('my-popover-trigger');
const content = document.getElementById('my-popover-content');

// Open popover
function openPopover() {
  trigger.setAttribute('aria-expanded', 'true');
  content.setAttribute('aria-hidden', 'false');
  content.focus();
}

// Close popover
function closePopover() {
  trigger.setAttribute('aria-expanded', 'false');
  content.setAttribute('aria-hidden', 'true');
  trigger.focus();
}

// Toggle popover
function togglePopover() {
  const isOpen = trigger.getAttribute('aria-expanded') === 'true';
  if (isOpen) {
    closePopover();
  } else {
    openPopover();
  }
}
```

### Click Outside to Close
```javascript
document.addEventListener('click', (e) => {
  const popover = e.target.closest('.popover');
  if (!popover) {
    // Click outside all popovers - close any open ones
    document.querySelectorAll('.popover').forEach(p => {
      const trigger = p.querySelector('[aria-expanded]');
      const content = p.querySelector('[data-popover]');
      if (trigger && content) {
        trigger.setAttribute('aria-expanded', 'false');
        content.setAttribute('aria-hidden', 'true');
      }
    });
  }
});
```

## Positioning

### Popover Positioning
Use `data-side` and `data-align` attributes on the popover content element:

```html
<!-- Position above trigger, aligned to start -->
<div data-popover data-side="top" data-align="start">
  <!-- Popover content -->
</div>

<!-- Position to the right, center aligned -->
<div data-popover data-side="right" data-align="center">
  <!-- Popover content -->
</div>

<!-- Position below (default), aligned to end -->
<div data-popover data-side="bottom" data-align="end">
  <!-- Popover content -->
</div>
```

### Side Options
- `top` - Above the trigger
- `right` - To the right of the trigger  
- `bottom` - Below the trigger (default)
- `left` - To the left of the trigger

### Alignment Options
- `start` - Align to the start edge
- `center` - Center align
- `end` - Align to the end edge

## Accessibility Features

- **Keyboard Navigation**: ESC key closes, focus management
- **Screen Reader Support**: Proper ARIA attributes and state management
- **Focus Management**: Automatic focus handling when opening/closing
- **State Communication**: Expanded/collapsed states properly announced

### Enhanced Accessibility
```html
<div id="accessible-popover" class="popover">
  <button id="accessible-popover-trigger" 
          type="button" 
          aria-expanded="false" 
          aria-controls="accessible-popover-content"
          aria-describedby="accessible-popover-desc"
          class="btn-outline">
    Settings
  </button>
  <div id="accessible-popover-content" 
       data-popover 
       aria-hidden="true" 
       aria-labelledby="accessible-popover-trigger"
       role="dialog"
       class="w-80">
    <div class="grid gap-4">
      <h4 id="accessible-popover-title">Configuration Settings</h4>
      <p id="accessible-popover-desc" class="text-sm text-muted-foreground">
        Adjust your application preferences below.
      </p>
      <!-- Content -->
    </div>
  </div>
</div>
```

## Best Practices

1. **Clear Triggers**: Use descriptive button text that indicates what the popover contains
2. **Appropriate Sizing**: Size popovers appropriately for their content
3. **Keyboard Support**: Ensure ESC key closes and tab navigation works
4. **Focus Management**: Set appropriate focus when opening
5. **Click Outside**: Allow clicking outside to close for non-critical content
6. **Content Length**: Keep content concise or use scrolling for long content
7. **Positioning**: Choose appropriate positioning based on context
8. **Form Integration**: Use proper form handling for interactive content

## Common Patterns

### Tooltip-Style Popover
```javascript
// Simple text content popover
function createTooltipPopover(trigger, content) {
  const popover = document.createElement('div');
  popover.className = 'popover';
  popover.innerHTML = `
    <div data-popover aria-hidden="true" class="max-w-xs">
      <p class="text-sm">${content}</p>
    </div>
  `;
  
  trigger.parentNode.insertBefore(popover, trigger.nextSibling);
  
  trigger.addEventListener('mouseenter', () => {
    popover.querySelector('[data-popover]').setAttribute('aria-hidden', 'false');
  });
  
  trigger.addEventListener('mouseleave', () => {
    popover.querySelector('[data-popover]').setAttribute('aria-hidden', 'true');
  });
}
```

### Settings Panel Pattern
```javascript
// Reusable settings popover
function createSettingsPopover(triggerSelector, settings) {
  const trigger = document.querySelector(triggerSelector);
  const settingsHTML = settings.map(setting => `
    <div class="grid gap-2">
      <label for="${setting.id}">${setting.label}</label>
      <input type="${setting.type}" id="${setting.id}" value="${setting.value || ''}">
    </div>
  `).join('');
  
  // Create and attach popover
  // Handle form submission
}
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function Popover({ trigger, children, side = 'bottom', align = 'start' }) {
  const [isOpen, setIsOpen] = useState(false);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="popover">
      <button 
        type="button"
        aria-expanded={isOpen}
        className="btn-outline"
        onClick={() => setIsOpen(!isOpen)}
      >
        {trigger}
      </button>
      <div 
        data-popover 
        aria-hidden={!isOpen}
        data-side={side}
        data-align={align}
        className="w-80"
      >
        {children}
      </div>
    </div>
  );
}
```

### HTMX Integration
```html
<div class="popover">
  <button 
    type="button" 
    aria-expanded="false" 
    aria-controls="dynamic-popover-content"
    hx-get="/api/popover-content"
    hx-target="#dynamic-popover-content"
    hx-trigger="click"
    class="btn-outline"
  >
    Load Content
  </button>
  <div id="dynamic-popover-content" data-popover aria-hidden="true" class="w-80">
    <!-- Content loaded dynamically -->
  </div>
</div>
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% call popover(
  id="demo-popover",
  trigger="Open popover",
  trigger_attrs={"class": "btn-outline"},
  popover_attrs={"class": "w-80"}
) %}
  <div class="grid gap-4">
    <header class="grid gap-1.5">
      <h4 class="leading-none font-medium">Dimensions</h4>
      <p class="text-muted-foreground text-sm">Set the dimensions for the layer.</p>
    </header>
    <form class="form grid gap-2">
      <div class="grid grid-cols-3 items-center gap-4">
        <label for="demo-popover-width">Width</label>
        <input type="text" id="demo-popover-width" value="100%" class="col-span-2 h-8"/>
      </div>
    </form>
  </div>
{% endcall %}
```

### Advanced Configuration
```jinja2
{% set popover_content %}
  <div class="grid gap-3">
    <h4 class="font-medium">{{ title }}</h4>
    <p class="text-sm text-muted-foreground">{{ description }}</p>
    {% for item in items %}
      <div class="flex items-center gap-2">
        {{ item.label }}: {{ item.value }}
      </div>
    {% endfor %}
  </div>
{% endset %}

{{ popover(
  id="config-popover",
  trigger="Configuration",
  content=popover_content,
  trigger_attrs={"class": "btn-ghost"},
  popover_attrs={"class": "w-64", "data-side": "right"}
) }}
```

## Related Components

- [Button](./button.md) - For popover triggers
- [Dialog](./dialog.md) - For modal interactions
- [Dropdown Menu](./dropdown-menu.md) - For menu-based interactions
- [Tooltip](./tooltip.md) - For simple text overlays
---
## sidebar

# Sidebar Component

A composable, themeable and customizable sidebar component.

## Basic Usage

```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-content-1">
        <h3 id="group-label-content-1">Getting started</h3>
        <ul>
          <li>
            <a href="#">
              <svg><!-- Icon --></svg>
              <span>Playground</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>

<main>
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar'))">Toggle sidebar</button>
  <h1>Content</h1>
</main>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/sidebar.min.js" defer></script>
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
- **`sidebar`** - Applied to the main sidebar container (`<aside>`)
- **`scrollbar`** - Applied to scrollable sections

### Supporting Classes
- Standard navigation and list styling
- Button classes for interactive elements
- Responsive utilities for mobile behavior

## Component Attributes

### Sidebar Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "sidebar" class | Yes |
| `aria-hidden` | boolean | Controls default visibility state | Recommended |
| `data-side` | string | Position: "left" or "right" (default: "left") | No |
| `id` | string | Unique identifier for multiple sidebars | No |

### Navigation Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `aria-label` | string | Describes the navigation purpose | Recommended |

### Group Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "group" | Yes |
| `aria-labelledby` | string | References heading ID | Recommended |

### Link Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-keep-mobile-sidebar-open` | boolean | Prevents mobile auto-close | No |

## HTML Structure

```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <!-- Optional header -->
    <header>
      <!-- Header content -->
    </header>
    
    <!-- Main content section -->
    <section class="scrollbar">
      <!-- Navigation group -->
      <div role="group" aria-labelledby="group-label-1">
        <h3 id="group-label-1">Group Title</h3>
        
        <ul>
          <!-- Regular link -->
          <li>
            <a href="/page">
              <svg><!-- Icon --></svg>
              <span>Page Title</span>
            </a>
          </li>
          
          <!-- Collapsible submenu -->
          <li>
            <details id="submenu-1">
              <summary aria-controls="submenu-1-content">
                <svg><!-- Icon --></svg>
                Submenu Title
              </summary>
              <ul id="submenu-1-content">
                <li>
                  <a href="/sub-page">
                    <span>Sub Page</span>
                  </a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
    
    <!-- Optional footer -->
    <footer>
      <!-- Footer content -->
    </footer>
  </nav>
</aside>

<main>
  <!-- Page content -->
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar'))">
    Toggle sidebar
  </button>
</main>
```

## Examples

### Basic Sidebar
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-main">
        <h3 id="group-label-main">Main Navigation</h3>
        
        <ul>
          <li>
            <a href="/">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
                <polyline points="9,22 9,12 15,12 15,22" />
              </svg>
              <span>Home</span>
            </a>
          </li>
          
          <li>
            <a href="/dashboard">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect width="7" height="9" x="3" y="3" rx="1" />
                <rect width="7" height="5" x="14" y="3" rx="1" />
                <rect width="7" height="9" x="14" y="12" rx="1" />
                <rect width="7" height="5" x="3" y="16" rx="1" />
              </svg>
              <span>Dashboard</span>
            </a>
          </li>
          
          <li>
            <a href="/projects">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z" />
                <path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z" />
              </svg>
              <span>Projects</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Sidebar with Collapsible Sections
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-content">
        <h3 id="group-label-content">Getting started</h3>

        <ul>
          <li>
            <a href="/playground">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m7 11 2-2-2-2" />
                <path d="M11 13h4" />
                <rect width="18" height="18" x="3" y="3" rx="2" ry="2" />
              </svg>
              <span>Playground</span>
            </a>
          </li>

          <li>
            <a href="/models">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M12 8V4H8" />
                <rect width="16" height="12" x="4" y="8" rx="2" />
                <path d="M2 14h2" />
                <path d="M20 14h2" />
                <path d="M15 13v2" />
                <path d="M9 13v2" />
              </svg>
              <span>Models</span>
            </a>
          </li>

          <li>
            <details id="submenu-settings">
              <summary aria-controls="submenu-settings-content">
                <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
                  <circle cx="12" cy="12" r="3" />
                </svg>
                Settings
              </summary>
              <ul id="submenu-settings-content">
                <li>
                  <a href="/settings/general">
                    <span>General</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/team">
                    <span>Team</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/billing">
                    <span>Billing</span>
                  </a>
                </li>
                <li>
                  <a href="/settings/limits">
                    <span>Limits</span>
                  </a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Sidebar with Header and Footer
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <!-- Header -->
    <header>
      <a href="/" class="btn-ghost p-2 h-12 w-full justify-start">
        <div class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256" class="h-4 w-4">
            <rect width="256" height="256" fill="none"></rect>
            <line x1="208" y1="128" x2="128" y2="208" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line>
            <line x1="192" y1="40" x2="40" y2="192" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line>
          </svg>
        </div>
        <div class="grid flex-1 text-left text-sm leading-tight">
          <span class="truncate font-medium">Your App</span>
          <span class="truncate text-xs">v1.0.0</span>
        </div>
      </a>
    </header>
    
    <!-- Main content -->
    <section class="scrollbar">
      <!-- Navigation groups -->
    </section>
    
    <!-- Footer -->
    <footer>
      <div class="p-4">
        <button type="button" class="btn-ghost w-full justify-start">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
            <polyline points="16,17 21,12 16,7" />
            <line x1="21" y1="12" x2="9" y2="12" />
          </svg>
          <span>Sign out</span>
        </button>
      </div>
    </footer>
  </nav>
</aside>
```

### Right-Side Sidebar
```html
<aside class="sidebar" data-side="right" aria-hidden="true">
  <nav aria-label="Right sidebar navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="group-label-tools">
        <h3 id="group-label-tools">Tools</h3>
        
        <ul>
          <li>
            <a href="/tools/inspector">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="m21 21-6-6m6 6v-4.8m0 4.8h-4.8" />
                <path d="M3 16.2V21m0 0h4.8M3 21l6-6" />
                <path d="M21 7.8V3m0 0h-4.8M21 3l-6 6" />
                <path d="M3 7.8V3m0 0h4.8M3 3l6 6" />
              </svg>
              <span>Inspector</span>
            </a>
          </li>
          
          <li>
            <a href="/tools/debugger">
              <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect width="20" height="16" x="2" y="4" rx="2" />
                <path d="m22 7-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7" />
              </svg>
              <span>Debugger</span>
            </a>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

### Multiple Sidebars
```html
<!-- Main navigation sidebar -->
<aside id="main-navigation" class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Main navigation">
    <!-- Main navigation content -->
  </nav>
</aside>

<!-- Tools sidebar -->
<aside id="tools-sidebar" class="sidebar" data-side="right" aria-hidden="true">
  <nav aria-label="Tools navigation">
    <!-- Tools content -->
  </nav>
</aside>

<main>
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: 'main-navigation', action: 'toggle' } }))">
    Toggle Main Nav
  </button>
  
  <button type="button" onclick="document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail: { id: 'tools-sidebar', action: 'toggle' } }))">
    Toggle Tools
  </button>
</main>
```

## JavaScript Events

### Toggle Sidebar
```javascript
// Toggle default sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar'));

// Toggle specific sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { id: 'main-navigation' }
}));
```

### Open/Close Sidebar
```javascript
// Open sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { action: 'open' }
}));

// Close sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { action: 'close' }
}));

// Open specific sidebar
document.dispatchEvent(new CustomEvent('basecoat:sidebar', {
  detail: { id: 'main-navigation', action: 'open' }
}));
```

### Listen for Initialization
```javascript
// Listen for component initialization
document.querySelectorAll('.sidebar').forEach(sidebar => {
  sidebar.addEventListener('basecoat:initialized', (e) => {
    console.log('Sidebar initialized:', e.target.id);
  });
});
```

### Custom Event Handlers
```javascript
// Create custom toggle button
function toggleSidebar(sidebarId = null, action = 'toggle') {
  const detail = { action };
  if (sidebarId) detail.id = sidebarId;
  
  document.dispatchEvent(new CustomEvent('basecoat:sidebar', { detail }));
}

// Usage
toggleSidebar('main-nav', 'open');
toggleSidebar(null, 'close'); // Close default sidebar
```

## Mobile Behavior

### Auto-Close on Mobile
By default, clicking links in the sidebar will close it on mobile devices:

```html
<li>
  <a href="/page">Page Link</a> <!-- Closes sidebar on mobile -->
</li>
```

### Prevent Auto-Close
Use `data-keep-mobile-sidebar-open` to prevent auto-closing:

```html
<li>
  <a href="/page" data-keep-mobile-sidebar-open>
    Page Link <!-- Keeps sidebar open on mobile -->
  </a>
</li>

<li>
  <button type="button" data-keep-mobile-sidebar-open>
    Button <!-- Keeps sidebar open on mobile -->
  </button>
</li>
```

## Accessibility Features

- **Semantic HTML**: Uses `<aside>` and `<nav>` landmarks
- **ARIA Labels**: Descriptive labels for navigation
- **Keyboard Navigation**: Full keyboard support
- **Screen Reader Support**: Proper grouping and labeling
- **Focus Management**: Logical tab order and focus indicators
- **State Announcement**: Hidden/visible states properly announced

### Enhanced Accessibility
```html
<aside class="sidebar" data-side="left" aria-hidden="false" aria-label="Main navigation">
  <nav aria-label="Primary site navigation">
    <section class="scrollbar">
      <div role="group" aria-labelledby="main-nav-heading">
        <h3 id="main-nav-heading">Main Navigation</h3>
        
        <ul role="list">
          <li role="listitem">
            <a href="/" aria-current="page">
              <svg aria-hidden="true"><!-- Icon --></svg>
              <span>Home</span>
            </a>
          </li>
          
          <li role="listitem">
            <details id="settings-menu">
              <summary aria-controls="settings-submenu" aria-expanded="false">
                <svg aria-hidden="true"><!-- Icon --></svg>
                Settings
              </summary>
              <ul id="settings-submenu" role="list">
                <li role="listitem">
                  <a href="/settings/general">General</a>
                </li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </section>
  </nav>
</aside>
```

## Styling Customization

### Custom Sidebar Positioning
```css
/* Left sidebar (default) */
.sidebar[data-side="left"] {
  left: 0;
}

/* Right sidebar */
.sidebar[data-side="right"] {
  right: 0;
}
```

### Custom Width
```css
.sidebar {
  width: 280px; /* Custom width */
}

@media (max-width: 768px) {
  .sidebar {
    width: 100%; /* Full width on mobile */
  }
}
```

### Custom Scrollbar
```css
.sidebar .scrollbar {
  scrollbar-width: thin;
  scrollbar-color: var(--muted) var(--background);
}

.sidebar .scrollbar::-webkit-scrollbar {
  width: 6px;
}

.sidebar .scrollbar::-webkit-scrollbar-track {
  background: var(--background);
}

.sidebar .scrollbar::-webkit-scrollbar-thumb {
  background: var(--muted);
  border-radius: 3px;
}
```

## Best Practices

1. **Clear Grouping**: Use semantic grouping with proper headings
2. **Consistent Icons**: Use consistent icon styles throughout
3. **Logical Order**: Organize navigation items logically
4. **Mobile-First**: Consider mobile experience and auto-close behavior
5. **Accessibility**: Provide proper ARIA labels and landmarks
6. **Keyboard Support**: Ensure all interactive elements are keyboard accessible
7. **Visual Hierarchy**: Use proper heading levels and visual distinction
8. **State Management**: Handle open/closed states appropriately

## Integration Examples

### React Integration
```jsx
import React, { useEffect } from 'react';

function Sidebar({ isOpen, onToggle }) {
  useEffect(() => {
    // Initialize Basecoat sidebar
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <aside className="sidebar" data-side="left" aria-hidden={!isOpen}>
      <nav aria-label="Sidebar navigation">
        <section className="scrollbar">
          {/* Navigation content */}
        </section>
      </nav>
    </aside>
  );
}
```

### Vue Integration
```vue
<template>
  <aside class="sidebar" data-side="left" :aria-hidden="!isOpen">
    <nav aria-label="Sidebar navigation">
      <section class="scrollbar">
        <!-- Navigation content -->
      </section>
    </nav>
  </aside>
</template>

<script>
export default {
  props: {
    isOpen: Boolean
  },
  mounted() {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }
}
</script>
```

### HTMX Integration
```html
<aside class="sidebar" data-side="left" aria-hidden="false">
  <nav aria-label="Sidebar navigation">
    <section class="scrollbar" hx-get="/api/navigation" hx-trigger="load">
      <!-- Navigation loaded dynamically -->
    </section>
  </nav>
</aside>
```

## Jinja/Nunjucks Macros

### Basic Usage
```jinja2
{% set menu = [
  { type: "group", label: "Getting started", items: [
    { label: "Playground", url: "#" },
    { label: "Models", url: "#" },
    { label: "Settings", type: "submenu", items: [
      { label: "General", url: "#" },
      { label: "Team", url: "#" },
      { label: "Billing", url: "#" },
      { label: "Limits", url: "#" }
    ] }
  ]}
] %}

{{ sidebar(
  label="Sidebar navigation",
  menu=menu
) }}
```

### Advanced Configuration
```jinja2
{{ sidebar(
  label="Main navigation",
  side="left",
  hidden=false,
  id="main-nav",
  menu=navigation_items,
  header=header_content,
  footer=footer_content
) }}
```

## Related Components

- [Button](./button.md) - For interactive elements within sidebar
- [Navigation](./navigation.md) - For primary site navigation
- [Breadcrumb](./breadcrumb.md) - For hierarchical navigation
- [Dropdown Menu](./dropdown-menu.md) - For nested menu items
---
## toast

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
---
## card

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
---
## accordion

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
---
## tabs

# Tabs Component

A set of layered sections of content—known as tab panels—that are displayed one at a time.

## Basic Usage

```html
<div class="tabs w-full" id="demo-tabs">
  <nav role="tablist" aria-orientation="horizontal" class="w-full">
    <button type="button" role="tab" id="demo-tabs-tab-1" aria-controls="demo-tabs-panel-1" aria-selected="true" tabindex="0">Account</button>
    <button type="button" role="tab" id="demo-tabs-tab-2" aria-controls="demo-tabs-panel-2" aria-selected="false" tabindex="0">Password</button>
  </nav>

  <div role="tabpanel" id="demo-tabs-panel-1" aria-labelledby="demo-tabs-tab-1" tabindex="-1" aria-selected="true">
    <div class="card">
      <header>
        <h2>Account</h2>
        <p>Make changes to your account here. Click save when you're done.</p>
      </header>
      <section>
        <form class="form grid gap-6">
          <div class="grid gap-3">
            <label for="account-name">Name</label>
            <input type="text" id="account-name" value="Pedro Duarte" />
          </div>
          <div class="grid gap-3">
            <label for="account-username">Username</label>
            <input type="text" id="account-username" value="@peduarte" />
          </div>
        </form>
      </section>
      <footer>
        <button type="button" class="btn">Save changes</button>
      </footer>
    </div>
  </div>

  <div role="tabpanel" id="demo-tabs-panel-2" aria-labelledby="demo-tabs-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="card">
      <header>
        <h2>Password</h2>
        <p>Change your password here. After saving, you'll be logged out.</p>
      </header>
      <section>
        <form class="form grid gap-6">
          <div class="grid gap-3">
            <label for="password-current">Current password</label>
            <input type="password" id="password-current" />
          </div>
          <div class="grid gap-3">
            <label for="password-new">New password</label>
            <input type="password" id="password-new" />
          </div>
        </form>
      </section>
      <footer>
        <button type="button" class="btn">Save Password</button>
      </footer>
    </div>
  </div>
</div>
```

## Required JavaScript

### CDN Installation
```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/tabs.min.js" defer></script>
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
- **`tabs`** - Applied to the main container
- **`w-full`** - Full width container (optional but common)

### Supporting Classes
- Card classes for content panels
- Form classes for form content
- Button classes for actions

### Tailwind Utilities Used
- `w-full` - Full width
- `grid gap-*` - Grid layout with spacing
- Various spacing and layout utilities

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "tabs" | Yes |
| `id` | string | Unique identifier | Recommended |

### Tablist (Nav) Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "tablist" | Yes |
| `aria-orientation` | string | "horizontal" or "vertical" | Recommended |

### Tab Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "button" | Yes |
| `role` | string | Must be "tab" | Yes |
| `id` | string | Unique identifier | Yes |
| `aria-controls` | string | References panel ID | Yes |
| `aria-selected` | boolean | Current selection state | Yes |
| `tabindex` | number | 0 for focusable, -1 for not | Yes |

### Tab Panel Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Must be "tabpanel" | Yes |
| `id` | string | Referenced by aria-controls | Yes |
| `aria-labelledby` | string | References tab button ID | Yes |
| `tabindex` | number | Usually -1 | Yes |
| `aria-selected` | boolean | Selection state | Yes |
| `hidden` | boolean | Hide inactive panels | Conditional |

## HTML Structure

```html
<div class="tabs" id="tabs-id">
  <!-- Tab navigation -->
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" 
            id="tab-1" 
            aria-controls="panel-1" 
            aria-selected="true" 
            tabindex="0">
      Tab 1
    </button>
    <button type="button" role="tab" 
            id="tab-2" 
            aria-controls="panel-2" 
            aria-selected="false" 
            tabindex="0">
      Tab 2
    </button>
  </nav>

  <!-- Tab panels -->
  <div role="tabpanel" 
       id="panel-1" 
       aria-labelledby="tab-1" 
       tabindex="-1" 
       aria-selected="true">
    <!-- Panel 1 content -->
  </div>

  <div role="tabpanel" 
       id="panel-2" 
       aria-labelledby="tab-2" 
       tabindex="-1" 
       aria-selected="false" 
       hidden>
    <!-- Panel 2 content -->
  </div>
</div>
```

## Examples

### Basic Text Tabs
```html
<div class="tabs" id="text-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="text-tab-1" aria-controls="text-panel-1" aria-selected="true" tabindex="0">
      Overview
    </button>
    <button type="button" role="tab" id="text-tab-2" aria-controls="text-panel-2" aria-selected="false" tabindex="0">
      Features
    </button>
    <button type="button" role="tab" id="text-tab-3" aria-controls="text-panel-3" aria-selected="false" tabindex="0">
      Reviews
    </button>
  </nav>

  <div role="tabpanel" id="text-panel-1" aria-labelledby="text-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <h3 class="text-lg font-semibold">Product Overview</h3>
      <p class="mt-2 text-muted-foreground">
        This is the overview content for our product. It provides a general introduction
        and highlights the main benefits.
      </p>
    </div>
  </div>

  <div role="tabpanel" id="text-panel-2" aria-labelledby="text-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="text-lg font-semibold">Key Features</h3>
      <ul class="mt-2 space-y-2 text-muted-foreground">
        <li>• Feature 1: Advanced functionality</li>
        <li>• Feature 2: User-friendly interface</li>
        <li>• Feature 3: Cross-platform support</li>
      </ul>
    </div>
  </div>

  <div role="tabpanel" id="text-panel-3" aria-labelledby="text-tab-3" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="text-lg font-semibold">Customer Reviews</h3>
      <p class="mt-2 text-muted-foreground">
        See what our customers are saying about this product.
      </p>
    </div>
  </div>
</div>
```

### Tabs with Icons
```html
<div class="tabs" id="icon-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="icon-tab-1" aria-controls="icon-panel-1" aria-selected="true" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
        <polyline points="9,22 9,12 15,12 15,22" />
      </svg>
      Home
    </button>
    <button type="button" role="tab" id="icon-tab-2" aria-controls="icon-panel-2" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
        <circle cx="12" cy="12" r="3" />
      </svg>
      Settings
    </button>
  </nav>

  <div role="tabpanel" id="icon-panel-1" aria-labelledby="icon-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <p>Home panel content</p>
    </div>
  </div>

  <div role="tabpanel" id="icon-panel-2" aria-labelledby="icon-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <p>Settings panel content</p>
    </div>
  </div>
</div>
```

### Vertical Tabs
```html
<div class="tabs flex gap-6" id="vertical-tabs">
  <nav role="tablist" aria-orientation="vertical" class="flex flex-col w-48">
    <button type="button" role="tab" id="v-tab-1" aria-controls="v-panel-1" aria-selected="true" tabindex="0" class="text-left">
      General
    </button>
    <button type="button" role="tab" id="v-tab-2" aria-controls="v-panel-2" aria-selected="false" tabindex="0" class="text-left">
      Security
    </button>
    <button type="button" role="tab" id="v-tab-3" aria-controls="v-panel-3" aria-selected="false" tabindex="0" class="text-left">
      Notifications
    </button>
    <button type="button" role="tab" id="v-tab-4" aria-controls="v-panel-4" aria-selected="false" tabindex="0" class="text-left">
      Advanced
    </button>
  </nav>

  <div class="flex-1">
    <div role="tabpanel" id="v-panel-1" aria-labelledby="v-tab-1" tabindex="-1" aria-selected="true">
      <h3 class="text-lg font-semibold mb-4">General Settings</h3>
      <p>Configure general application settings here.</p>
    </div>

    <div role="tabpanel" id="v-panel-2" aria-labelledby="v-tab-2" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Security Settings</h3>
      <p>Manage your security preferences and authentication options.</p>
    </div>

    <div role="tabpanel" id="v-panel-3" aria-labelledby="v-tab-3" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Notification Preferences</h3>
      <p>Choose how and when you want to receive notifications.</p>
    </div>

    <div role="tabpanel" id="v-panel-4" aria-labelledby="v-tab-4" tabindex="-1" aria-selected="false" hidden>
      <h3 class="text-lg font-semibold mb-4">Advanced Options</h3>
      <p>Advanced settings for power users.</p>
    </div>
  </div>
</div>
```

### Tabs with Badges
```html
<div class="tabs" id="badge-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="badge-tab-1" aria-controls="badge-panel-1" aria-selected="true" tabindex="0">
      Messages
      <span class="badge ml-1.5">12</span>
    </button>
    <button type="button" role="tab" id="badge-tab-2" aria-controls="badge-panel-2" aria-selected="false" tabindex="0">
      Notifications
      <span class="badge ml-1.5">3</span>
    </button>
    <button type="button" role="tab" id="badge-tab-3" aria-controls="badge-panel-3" aria-selected="false" tabindex="0">
      Updates
    </button>
  </nav>

  <div role="tabpanel" id="badge-panel-1" aria-labelledby="badge-tab-1" tabindex="-1" aria-selected="true">
    <div class="py-4">
      <h3 class="font-semibold">Unread Messages (12)</h3>
      <!-- Message list -->
    </div>
  </div>

  <div role="tabpanel" id="badge-panel-2" aria-labelledby="badge-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="font-semibold">Recent Notifications (3)</h3>
      <!-- Notification list -->
    </div>
  </div>

  <div role="tabpanel" id="badge-panel-3" aria-labelledby="badge-tab-3" tabindex="-1" aria-selected="false" hidden>
    <div class="py-4">
      <h3 class="font-semibold">System Updates</h3>
      <!-- Update list -->
    </div>
  </div>
</div>
```

### Disabled Tabs
```html
<div class="tabs" id="disabled-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="d-tab-1" aria-controls="d-panel-1" aria-selected="true" tabindex="0">
      Available
    </button>
    <button type="button" role="tab" id="d-tab-2" aria-controls="d-panel-2" aria-selected="false" tabindex="0">
      Premium
    </button>
    <button type="button" role="tab" id="d-tab-3" aria-controls="d-panel-3" aria-selected="false" tabindex="-1" aria-disabled="true" class="opacity-50 cursor-not-allowed">
      Coming Soon
    </button>
  </nav>

  <div role="tabpanel" id="d-panel-1" aria-labelledby="d-tab-1" tabindex="-1" aria-selected="true">
    <p class="py-4">This feature is available to all users.</p>
  </div>

  <div role="tabpanel" id="d-panel-2" aria-labelledby="d-tab-2" tabindex="-1" aria-selected="false" hidden>
    <p class="py-4">This feature is available to premium users.</p>
  </div>

  <div role="tabpanel" id="d-panel-3" aria-labelledby="d-tab-3" tabindex="-1" aria-selected="false" hidden>
    <p class="py-4">This feature is coming soon.</p>
  </div>
</div>
```

## JavaScript Events

### Component Events
```javascript
const tabs = document.getElementById('my-tabs');

// Listen for initialization
tabs.addEventListener('basecoat:initialized', (e) => {
  console.log('Tabs initialized');
});

// Listen for tab changes
tabs.addEventListener('basecoat:tabs:change', (e) => {
  console.log('Tab changed to:', e.detail.tabId);
});
```

### Manual Tab Control
```javascript
// Switch to specific tab
function switchToTab(tabId) {
  const tab = document.getElementById(tabId);
  if (tab) {
    tab.click();
  }
}

// Get active tab
function getActiveTab() {
  return document.querySelector('[role="tab"][aria-selected="true"]');
}

// Disable/enable tab
function setTabEnabled(tabId, enabled) {
  const tab = document.getElementById(tabId);
  if (tab) {
    tab.setAttribute('aria-disabled', !enabled);
    tab.setAttribute('tabindex', enabled ? '0' : '-1');
    tab.classList.toggle('opacity-50', !enabled);
    tab.classList.toggle('cursor-not-allowed', !enabled);
  }
}
```

## Keyboard Navigation

The tabs component supports full keyboard navigation:

- **Tab**: Move focus between tabs
- **Arrow Keys**: Navigate between tabs (when focused on tablist)
- **Home**: Go to first tab
- **End**: Go to last tab
- **Space/Enter**: Activate focused tab

## Accessibility Features

- **ARIA Roles**: Proper tablist, tab, and tabpanel roles
- **Screen Reader Support**: Clear announcements of tab states
- **Keyboard Navigation**: Full keyboard support
- **Focus Management**: Proper focus handling
- **State Communication**: Selected and disabled states

### Enhanced Accessibility
```html
<div class="tabs" id="accessible-tabs">
  <h2 id="tabs-heading" class="text-lg font-semibold mb-4">User Preferences</h2>
  <nav role="tablist" aria-orientation="horizontal" aria-labelledby="tabs-heading">
    <button type="button" 
            role="tab" 
            id="a-tab-1" 
            aria-controls="a-panel-1" 
            aria-selected="true" 
            tabindex="0"
            aria-describedby="a-tab-1-desc">
      Profile
    </button>
    <span id="a-tab-1-desc" class="sr-only">Edit your profile information</span>
    
    <button type="button" 
            role="tab" 
            id="a-tab-2" 
            aria-controls="a-panel-2" 
            aria-selected="false" 
            tabindex="0"
            aria-describedby="a-tab-2-desc">
      Privacy
    </button>
    <span id="a-tab-2-desc" class="sr-only">Manage your privacy settings</span>
  </nav>

  <div role="tabpanel" 
       id="a-panel-1" 
       aria-labelledby="a-tab-1" 
       tabindex="-1" 
       aria-selected="true">
    <!-- Profile content -->
  </div>

  <div role="tabpanel" 
       id="a-panel-2" 
       aria-labelledby="a-tab-2" 
       tabindex="-1" 
       aria-selected="false" 
       hidden>
    <!-- Privacy content -->
  </div>
</div>
```

## Styling Customization

### Custom Tab Styles
```css
/* Pill-style tabs */
.tabs-pills [role="tab"] {
  border-radius: 9999px;
  padding: 0.5rem 1rem;
}

/* Underline tabs */
.tabs-underline [role="tab"] {
  border-bottom: 2px solid transparent;
}

.tabs-underline [role="tab"][aria-selected="true"] {
  border-bottom-color: var(--primary);
}

/* Boxed tabs */
.tabs-boxed [role="tab"] {
  border: 1px solid var(--border);
  margin-right: -1px;
}

.tabs-boxed [role="tab"]:first-child {
  border-radius: 0.5rem 0 0 0.5rem;
}

.tabs-boxed [role="tab"]:last-child {
  border-radius: 0 0.5rem 0.5rem 0;
}
```

### Responsive Tabs
```html
<div class="tabs" id="responsive-tabs">
  <nav role="tablist" aria-orientation="horizontal" class="flex overflow-x-auto scrollbar-none">
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-1" aria-controls="r-panel-1" aria-selected="true" tabindex="0">
      Dashboard
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-2" aria-controls="r-panel-2" aria-selected="false" tabindex="0">
      Analytics
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-3" aria-controls="r-panel-3" aria-selected="false" tabindex="0">
      Reports
    </button>
    <button type="button" role="tab" class="whitespace-nowrap flex-shrink-0" id="r-tab-4" aria-controls="r-panel-4" aria-selected="false" tabindex="0">
      Settings
    </button>
  </nav>
  <!-- Tab panels -->
</div>
```

## Best Practices

1. **Clear Labels**: Use descriptive tab labels
2. **Logical Order**: Arrange tabs in a logical sequence
3. **Visual Feedback**: Clearly indicate active tab
4. **Keyboard Support**: Ensure full keyboard navigation
5. **Touch Targets**: Adequate size for mobile
6. **Content Loading**: Consider lazy loading for heavy content
7. **State Persistence**: Consider saving active tab state
8. **Responsive Design**: Handle overflow gracefully on mobile

## Common Patterns

### Settings Tabs
```html
<div class="tabs" id="settings-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="s-tab-1" aria-controls="s-panel-1" aria-selected="true" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
        <circle cx="12" cy="7" r="4" />
      </svg>
      Profile
    </button>
    <button type="button" role="tab" id="s-tab-2" aria-controls="s-panel-2" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <path d="M17 11h1a3 3 0 0 1 0 6h-1" />
        <path d="M9 12v6" />
        <path d="M13 12v6" />
        <path d="M14 7.5c-1 0-1.44.5-3 .5s-2-.5-3-.5" />
        <path d="M2 5c0-1.1.9-2 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5z" />
      </svg>
      Billing
    </button>
    <button type="button" role="tab" id="s-tab-3" aria-controls="s-panel-3" aria-selected="false" tabindex="0">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
      Advanced
    </button>
  </nav>
  <!-- Tab panels with forms -->
</div>
```

### Product Details Tabs
```html
<div class="tabs" id="product-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" role="tab" id="p-tab-1" aria-controls="p-panel-1" aria-selected="true" tabindex="0">
      Description
    </button>
    <button type="button" role="tab" id="p-tab-2" aria-controls="p-panel-2" aria-selected="false" tabindex="0">
      Specifications
    </button>
    <button type="button" role="tab" id="p-tab-3" aria-controls="p-panel-3" aria-selected="false" tabindex="0">
      Reviews (24)
    </button>
    <button type="button" role="tab" id="p-tab-4" aria-controls="p-panel-4" aria-selected="false" tabindex="0">
      Q&A (8)
    </button>
  </nav>
  <!-- Product information panels -->
</div>
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function Tabs({ items, defaultTab = 0 }) {
  const [activeTab, setActiveTab] = useState(defaultTab);
  
  useEffect(() => {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }, []);

  return (
    <div className="tabs">
      <nav role="tablist" aria-orientation="horizontal">
        {items.map((item, index) => (
          <button
            key={index}
            type="button"
            role="tab"
            id={`tab-${index}`}
            aria-controls={`panel-${index}`}
            aria-selected={activeTab === index}
            tabindex={activeTab === index ? 0 : -1}
            onClick={() => setActiveTab(index)}
          >
            {item.label}
          </button>
        ))}
      </nav>

      {items.map((item, index) => (
        <div
          key={index}
          role="tabpanel"
          id={`panel-${index}`}
          aria-labelledby={`tab-${index}`}
          tabindex="-1"
          aria-selected={activeTab === index}
          hidden={activeTab !== index}
        >
          {item.content}
        </div>
      ))}
    </div>
  );
}
```

### Vue Integration
```vue
<template>
  <div class="tabs">
    <nav role="tablist" aria-orientation="horizontal">
      <button
        v-for="(item, index) in items"
        :key="index"
        type="button"
        role="tab"
        :id="`tab-${index}`"
        :aria-controls="`panel-${index}`"
        :aria-selected="activeTab === index"
        :tabindex="activeTab === index ? 0 : -1"
        @click="activeTab = index"
      >
        {{ item.label }}
      </button>
    </nav>

    <div
      v-for="(item, index) in items"
      :key="index"
      role="tabpanel"
      :id="`panel-${index}`"
      :aria-labelledby="`tab-${index}`"
      tabindex="-1"
      :aria-selected="activeTab === index"
      :hidden="activeTab !== index"
    >
      <slot :name="`panel-${index}`">
        {{ item.content }}
      </slot>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    items: Array,
    defaultTab: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      activeTab: this.defaultTab
    };
  },
  mounted() {
    if (window.basecoat) {
      window.basecoat.initAll();
    }
  }
};
</script>
```

### HTMX Integration
```html
<div class="tabs" id="htmx-tabs">
  <nav role="tablist" aria-orientation="horizontal">
    <button type="button" 
            role="tab" 
            id="htmx-tab-1" 
            aria-controls="htmx-panel-1" 
            aria-selected="true" 
            tabindex="0"
            hx-get="/api/content/overview"
            hx-target="#htmx-panel-1"
            hx-trigger="click once">
      Overview
    </button>
    <button type="button" 
            role="tab" 
            id="htmx-tab-2" 
            aria-controls="htmx-panel-2" 
            aria-selected="false" 
            tabindex="0"
            hx-get="/api/content/details"
            hx-target="#htmx-panel-2"
            hx-trigger="click once">
      Details
    </button>
  </nav>

  <div role="tabpanel" id="htmx-panel-1" aria-labelledby="htmx-tab-1" tabindex="-1" aria-selected="true">
    <div class="skeleton">Loading overview...</div>
  </div>

  <div role="tabpanel" id="htmx-panel-2" aria-labelledby="htmx-tab-2" tabindex="-1" aria-selected="false" hidden>
    <div class="skeleton">Loading details...</div>
  </div>
</div>
```

## Related Components

- [Accordion](./accordion.md) - For expandable content sections
- [Card](./card.md) - For content containers within tabs
- [Navigation](./navigation.md) - For page-level navigation
- [Button Group](./button-group.md) - For alternative switching patterns
---
## avatar

# Avatar Component

An image element with a fallback for representing the user.

## Important Note

**There is no dedicated Avatar component in Basecoat.** Avatars are simply `<img>` elements styled with Tailwind utility classes.

## Basic Usage

```html
<img class="size-8 shrink-0 object-cover rounded-full" alt="@hunvreus" src="https://github.com/hunvreus.png" />
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard Tailwind utilities

### Supporting Classes
- **`size-*`** - Sets width and height (size-6, size-8, size-10, size-12, etc.)
- **`shrink-0`** - Prevents image from shrinking in flex containers
- **`object-cover`** - Ensures proper image scaling within container
- **`rounded-*`** - Border radius (rounded-full, rounded-lg, rounded-md)

### Tailwind Utilities Used
- `size-8` - Sets width and height to 2rem (32px)
- `shrink-0` - Prevents flex shrinking
- `object-cover` - Cover object fit for proper image scaling
- `rounded-full` - Circular avatar shape
- `rounded-lg` - Rounded rectangle shape
- `ring-*` - Ring/border effects for grouped avatars
- `grayscale` - Desaturated color effect
- `-space-x-*` - Negative horizontal spacing for overlapped avatars

## Component Attributes

### Image Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `src` | string | Image source URL | Yes |
| `alt` | string | Alternative text for accessibility | Yes |
| `class` | string | Tailwind utility classes | Yes |

### No JavaScript Required
This is a pure CSS/HTML implementation using standard image elements.

## HTML Structure

```html
<img class="[size-classes] [shape-classes] [object-classes]" alt="[description]" src="[image-url]" />
```

## Examples

### Different Sizes
```html
<!-- Extra Small (24px) -->
<img class="size-6 shrink-0 object-cover rounded-full" alt="Small avatar" src="https://github.com/hunvreus.png" />

<!-- Small (32px) -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="Medium avatar" src="https://github.com/hunvreus.png" />

<!-- Medium (40px) -->
<img class="size-10 shrink-0 object-cover rounded-full" alt="Medium avatar" src="https://github.com/hunvreus.png" />

<!-- Large (48px) -->
<img class="size-12 shrink-0 object-cover rounded-full" alt="Large avatar" src="https://github.com/hunvreus.png" />

<!-- Extra Large (64px) -->
<img class="size-16 shrink-0 object-cover rounded-full" alt="Extra large avatar" src="https://github.com/hunvreus.png" />
```

### Different Shapes
```html
<!-- Circular (most common) -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="Circular avatar" src="https://github.com/hunvreus.png" />

<!-- Rounded rectangle -->
<img class="size-8 shrink-0 object-cover rounded-lg" alt="Rounded avatar" src="https://github.com/shadcn.png" />

<!-- Square -->
<img class="size-8 shrink-0 object-cover rounded-md" alt="Square avatar" src="https://github.com/hunvreus.png" />

<!-- No rounding -->
<img class="size-8 shrink-0 object-cover" alt="Square avatar" src="https://github.com/hunvreus.png" />
```

### Avatar Group (Overlapping)
```html
<div class="flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full">
  <img alt="@hunvreus" src="https://github.com/hunvreus.png" />
  <img alt="@shadcn" src="https://github.com/shadcn.png" />
  <img alt="@adamwathan" src="https://github.com/adamwathan.png" />
</div>
```

### Avatar with Ring/Border
```html
<!-- Simple ring -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-primary" alt="Avatar with ring" src="https://github.com/hunvreus.png" />

<!-- Ring with offset -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-primary ring-offset-2 ring-offset-background" alt="Avatar with offset ring" src="https://github.com/hunvreus.png" />

<!-- Status indicator ring -->
<img class="size-8 shrink-0 object-cover rounded-full ring-2 ring-success" alt="Online user" src="https://github.com/hunvreus.png" />
```

### Avatar with Fallback Text
```html
<!-- Using a div for text fallback when no image -->
<div class="size-8 shrink-0 rounded-full bg-primary text-primary-foreground flex items-center justify-center text-sm font-medium">
  JD
</div>

<!-- Using initials for fallback -->
<div class="size-10 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-sm font-semibold">
  AB
</div>
```

### Avatar with Status Indicator
```html
<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Online status -->
  <div class="absolute bottom-0 right-0 size-3 bg-success rounded-full ring-2 ring-background"></div>
</div>

<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Away status -->
  <div class="absolute bottom-0 right-0 size-3 bg-warning rounded-full ring-2 ring-background"></div>
</div>

<div class="relative">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="User avatar" src="https://github.com/hunvreus.png" />
  <!-- Offline status -->
  <div class="absolute bottom-0 right-0 size-3 bg-muted rounded-full ring-2 ring-background"></div>
</div>
```

### Avatar Sizes Reference
```html
<!-- Size reference -->
<div class="flex items-center gap-4">
  <!-- xs: 16px -->
  <img class="size-4 shrink-0 object-cover rounded-full" alt="Extra small" src="https://github.com/hunvreus.png" />
  
  <!-- sm: 20px -->
  <img class="size-5 shrink-0 object-cover rounded-full" alt="Small" src="https://github.com/hunvreus.png" />
  
  <!-- md: 24px -->
  <img class="size-6 shrink-0 object-cover rounded-full" alt="Medium" src="https://github.com/hunvreus.png" />
  
  <!-- lg: 32px -->
  <img class="size-8 shrink-0 object-cover rounded-full" alt="Large" src="https://github.com/hunvreus.png" />
  
  <!-- xl: 40px -->
  <img class="size-10 shrink-0 object-cover rounded-full" alt="Extra large" src="https://github.com/hunvreus.png" />
  
  <!-- 2xl: 48px -->
  <img class="size-12 shrink-0 object-cover rounded-full" alt="2X large" src="https://github.com/hunvreus.png" />
</div>
```

### Avatar with Custom Dimensions
```html
<!-- Custom width/height -->
<img class="w-20 h-20 shrink-0 object-cover rounded-full" alt="Custom size avatar" src="https://github.com/hunvreus.png" />

<!-- Aspect ratio preservation -->
<img class="w-16 aspect-square shrink-0 object-cover rounded-lg" alt="Square aspect ratio" src="https://github.com/hunvreus.png" />
```

### Avatar Group with Count Indicator
```html
<div class="flex items-center">
  <!-- Avatar group -->
  <div class="flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full">
    <img alt="User 1" src="https://github.com/hunvreus.png" />
    <img alt="User 2" src="https://github.com/shadcn.png" />
    <img alt="User 3" src="https://github.com/adamwathan.png" />
  </div>
  
  <!-- Count indicator -->
  <div class="ml-2 size-8 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-xs font-medium">
    +5
  </div>
</div>
```

### Clickable Avatar
```html
<!-- Link avatar -->
<a href="/profile" class="block">
  <img class="size-8 shrink-0 object-cover rounded-full hover:ring-2 hover:ring-primary transition-all" alt="Profile" src="https://github.com/hunvreus.png" />
</a>

<!-- Button avatar -->
<button type="button" class="block">
  <img class="size-8 shrink-0 object-cover rounded-full hover:ring-2 hover:ring-primary transition-all" alt="User menu" src="https://github.com/hunvreus.png" />
</button>
```

## Accessibility Features

- **Alt Text**: Always provide descriptive alternative text
- **Semantic HTML**: Uses standard `<img>` elements
- **Focus States**: Support for focus indicators when clickable
- **Screen Reader Support**: Proper alternative text announcements

### Enhanced Accessibility
```html
<!-- Descriptive alt text -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="John Smith, Software Engineer" src="/avatars/john-smith.jpg" />

<!-- Role for decorative avatars -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="" role="presentation" src="/avatars/decorative.jpg" />

<!-- With ARIA label for interactive avatars -->
<button type="button" aria-label="Open user menu for John Smith">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="" src="/avatars/john-smith.jpg" />
</button>
```

## Best Practices

1. **Consistent Sizing**: Use consistent avatar sizes within the same context
2. **Meaningful Alt Text**: Provide descriptive alternative text
3. **Loading States**: Consider placeholder or loading states for slow networks
4. **Fallback Content**: Provide fallbacks for missing images
5. **Performance**: Optimize images for file size and format
6. **Responsive**: Consider different sizes for different screen sizes
7. **Accessibility**: Ensure proper contrast for text fallbacks
8. **Error Handling**: Handle broken image URLs gracefully

## Common Patterns

### User Profile Card
```html
<div class="flex items-center gap-3 p-4 border rounded-lg">
  <img class="size-12 shrink-0 object-cover rounded-full" alt="John Doe" src="https://github.com/hunvreus.png" />
  <div>
    <h3 class="font-medium">John Doe</h3>
    <p class="text-sm text-muted-foreground">Software Engineer</p>
  </div>
</div>
```

### Comment Thread
```html
<div class="flex gap-3">
  <img class="size-8 shrink-0 object-cover rounded-full" alt="Sarah Wilson" src="https://github.com/shadcn.png" />
  <div class="flex-1">
    <div class="flex items-center gap-2">
      <span class="font-medium text-sm">Sarah Wilson</span>
      <span class="text-xs text-muted-foreground">2 hours ago</span>
    </div>
    <p class="text-sm mt-1">This looks great! Thanks for the update.</p>
  </div>
</div>
```

### Team List
```html
<div class="space-y-3">
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-3">
      <img class="size-10 shrink-0 object-cover rounded-full" alt="Alice Johnson" src="https://github.com/hunvreus.png" />
      <div>
        <h4 class="font-medium">Alice Johnson</h4>
        <p class="text-sm text-muted-foreground">Team Lead</p>
      </div>
    </div>
    <div class="relative">
      <div class="size-3 bg-success rounded-full"></div>
    </div>
  </div>
</div>
```

## Error Handling

### Image Fallback with JavaScript
```html
<img 
  class="size-8 shrink-0 object-cover rounded-full" 
  alt="User avatar" 
  src="https://example.com/avatar.jpg"
  onerror="this.style.display='none'; this.nextElementSibling.style.display='flex'"
/>
<div 
  class="size-8 shrink-0 rounded-full bg-muted text-muted-foreground flex items-center justify-center text-sm font-medium" 
  style="display: none;"
>
  UA
</div>
```

### CSS-Only Fallback
```html
<div class="size-8 shrink-0 rounded-full overflow-hidden bg-muted flex items-center justify-center">
  <img 
    class="w-full h-full object-cover" 
    alt="User avatar" 
    src="https://example.com/avatar.jpg"
    style="display: block;"
    onerror="this.style.display='none'"
  />
  <span class="text-sm font-medium text-muted-foreground">UA</span>
</div>
```

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Avatar({ src, alt, size = 8, shape = 'rounded-full', className = '', ...props }) {
  const sizeClass = `size-${size}`;
  
  return (
    <img
      className={`${sizeClass} shrink-0 object-cover ${shape} ${className}`}
      alt={alt}
      src={src}
      {...props}
    />
  );
}

// Usage
<Avatar src="https://github.com/hunvreus.png" alt="User" size={10} />
```

### Vue Integration
```vue
<template>
  <img 
    :class="avatarClasses"
    :alt="alt"
    :src="src"
    v-bind="$attrs"
  />
</template>

<script>
export default {
  props: {
    src: String,
    alt: String,
    size: {
      type: Number,
      default: 8
    },
    shape: {
      type: String,
      default: 'rounded-full'
    }
  },
  computed: {
    avatarClasses() {
      return `size-${this.size} shrink-0 object-cover ${this.shape}`;
    }
  }
};
</script>
```

### Avatar with Loading State
```html
<!-- Loading skeleton -->
<div class="size-8 shrink-0 rounded-full bg-muted animate-pulse"></div>

<!-- Loaded avatar -->
<img class="size-8 shrink-0 object-cover rounded-full" alt="User" src="https://github.com/hunvreus.png" />
```

## Size Guide

| Class | Size | Pixels | Use Case |
|-------|------|--------|----------|
| `size-4` | 1rem | 16px | Tiny avatars, icons |
| `size-5` | 1.25rem | 20px | Small inline avatars |
| `size-6` | 1.5rem | 24px | Default small size |
| `size-8` | 2rem | 32px | Standard avatar size |
| `size-10` | 2.5rem | 40px | Medium avatars |
| `size-12` | 3rem | 48px | Large avatars |
| `size-16` | 4rem | 64px | Profile headers |
| `size-20` | 5rem | 80px | Hero avatars |

## Related Components

- [Button](./button.md) - For clickable avatar functionality
- [Badge](./badge.md) - For status indicators on avatars
- [Card](./card.md) - For profile cards containing avatars
- [Dropdown Menu](./dropdown-menu.md) - For avatar-triggered menus
---
## progress

# Progress Component

Displays an indicator showing the completion progress of a task, typically displayed as a progress bar.

## Important Note

**There is no dedicated Progress component in Basecoat.** Progress bars are pure HTML composition using Tailwind utility classes.

## Basic Usage

```html
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 66%"></div>
</div>
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard Tailwind utilities

### Supporting Classes
- **Track Container**: `bg-primary/20`, `relative`, `h-*`, `w-full`, `overflow-hidden`, `rounded-*`
- **Progress Indicator**: `bg-primary`, `h-full`, `w-full`, `flex-1`, `transition-all`

### Tailwind Utilities Used
- `bg-primary/20` - Semi-transparent background for track
- `bg-primary` - Solid background for progress indicator
- `relative` - Positioning context
- `h-2` - Height of progress bar (8px)
- `w-full` - Full width
- `overflow-hidden` - Clips overflowing content
- `rounded-full` - Fully rounded ends
- `h-full` - Full height of indicator
- `flex-1` - Flexible growth
- `transition-all` - Smooth animations

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Track styling classes | Yes |

### Indicator Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Progress styling classes | Yes |
| `style` | string | Inline width percentage | Yes |
| `role` | string | "progressbar" for accessibility | Recommended |
| `aria-valuenow` | number | Current value | Recommended |
| `aria-valuemin` | number | Minimum value (usually 0) | Recommended |
| `aria-valuemax` | number | Maximum value (usually 100) | Recommended |

### No JavaScript Required (Basic)
Basic progress bars work with pure CSS and inline styles.

## HTML Structure

```html
<!-- Track container -->
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <!-- Progress indicator -->
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: [percentage]%"></div>
</div>
```

## Examples

### Different Sizes
```html
<!-- Small (4px) -->
<div class="bg-primary/20 relative h-1 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 45%"></div>
</div>

<!-- Default (8px) -->
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 66%"></div>
</div>

<!-- Medium (12px) -->
<div class="bg-primary/20 relative h-3 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 80%"></div>
</div>

<!-- Large (16px) -->
<div class="bg-primary/20 relative h-4 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 90%"></div>
</div>
```

### Different Colors
```html
<!-- Success progress -->
<div class="bg-success/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-success h-full w-full flex-1 transition-all" style="width: 75%"></div>
</div>

<!-- Warning progress -->
<div class="bg-warning/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-warning h-full w-full flex-1 transition-all" style="width: 50%"></div>
</div>

<!-- Error progress -->
<div class="bg-error/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-error h-full w-full flex-1 transition-all" style="width: 30%"></div>
</div>

<!-- Custom colors -->
<div class="bg-blue-200 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-blue-600 h-full w-full flex-1 transition-all" style="width: 65%"></div>
</div>

<div class="bg-purple-200 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-purple-600 h-full w-full flex-1 transition-all" style="width: 85%"></div>
</div>
```

### Different Shapes
```html
<!-- Fully rounded (default) -->
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 60%"></div>
</div>

<!-- Rounded corners -->
<div class="bg-primary/20 relative h-3 w-full overflow-hidden rounded-lg">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 60%"></div>
</div>

<!-- Square corners -->
<div class="bg-primary/20 relative h-3 w-full overflow-hidden">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 60%"></div>
</div>

<!-- Custom rounded -->
<div class="bg-primary/20 relative h-3 w-full overflow-hidden rounded-md">
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 60%"></div>
</div>
```

### With Labels
```html
<!-- Progress with percentage -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Progress</span>
    <span>75%</span>
  </div>
  <div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 75%"></div>
  </div>
</div>

<!-- Progress with status text -->
<div class="space-y-2">
  <div class="flex justify-between items-center">
    <div>
      <h4 class="text-sm font-medium">Uploading files...</h4>
      <p class="text-xs text-muted-foreground">3 of 4 files completed</p>
    </div>
    <span class="text-sm text-muted-foreground">75%</span>
  </div>
  <div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 75%"></div>
  </div>
</div>

<!-- Progress with time remaining -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Processing</span>
    <span>2 minutes remaining</span>
  </div>
  <div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 45%"></div>
  </div>
</div>
```

### Segmented Progress
```html
<!-- Multi-step progress -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Step 2 of 4</span>
    <span>50%</span>
  </div>
  <div class="bg-muted relative h-2 w-full overflow-hidden rounded-full">
    <!-- Completed segments -->
    <div class="absolute left-0 top-0 h-full bg-primary" style="width: 50%"></div>
    <!-- Segment dividers -->
    <div class="absolute left-1/4 top-0 h-full w-px bg-background"></div>
    <div class="absolute left-1/2 top-0 h-full w-px bg-background"></div>
    <div class="absolute left-3/4 top-0 h-full w-px bg-background"></div>
  </div>
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>Info</span>
    <span>Review</span>
    <span>Payment</span>
    <span>Complete</span>
  </div>
</div>

<!-- Multiple progress bars -->
<div class="space-y-3">
  <div>
    <div class="flex justify-between text-sm mb-1">
      <span>HTML</span>
      <span>90%</span>
    </div>
    <div class="bg-orange-200 relative h-2 w-full overflow-hidden rounded-full">
      <div class="bg-orange-600 h-full w-full flex-1 transition-all" style="width: 90%"></div>
    </div>
  </div>
  
  <div>
    <div class="flex justify-between text-sm mb-1">
      <span>CSS</span>
      <span>75%</span>
    </div>
    <div class="bg-blue-200 relative h-2 w-full overflow-hidden rounded-full">
      <div class="bg-blue-600 h-full w-full flex-1 transition-all" style="width: 75%"></div>
    </div>
  </div>
  
  <div>
    <div class="flex justify-between text-sm mb-1">
      <span>JavaScript</span>
      <span>45%</span>
    </div>
    <div class="bg-yellow-200 relative h-2 w-full overflow-hidden rounded-full">
      <div class="bg-yellow-600 h-full w-full flex-1 transition-all" style="width: 45%"></div>
    </div>
  </div>
</div>
```

### Animated Progress
```html
<!-- Indeterminate progress (loading animation) -->
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full absolute left-0 animate-pulse" style="width: 30%"></div>
</div>

<!-- Sliding animation -->
<style>
@keyframes slide {
  0% { transform: translateX(-100%); }
  100% { transform: translateX(400%); }
}
.progress-slide {
  animation: slide 2s ease-in-out infinite;
}
</style>

<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div class="bg-primary h-full w-1/4 progress-slide"></div>
</div>

<!-- Growing animation -->
<div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
  <div id="growing-progress" class="bg-primary h-full w-full flex-1 transition-all duration-1000" style="width: 0%"></div>
</div>

<script>
// Animate to 80% over 3 seconds
let width = 0;
const target = 80;
const interval = setInterval(() => {
  width += 2;
  document.getElementById('growing-progress').style.width = width + '%';
  if (width >= target) {
    clearInterval(interval);
  }
}, 75);
</script>
```

### Progress with Icon
```html
<!-- Progress with success icon -->
<div class="space-y-2">
  <div class="flex justify-between items-center">
    <div class="flex items-center gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-success">
        <path d="M20 6 9 17l-5-5" />
      </svg>
      <span class="text-sm">Upload Complete</span>
    </div>
    <span class="text-sm text-success">100%</span>
  </div>
  <div class="bg-success/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-success h-full w-full flex-1 transition-all" style="width: 100%"></div>
  </div>
</div>

<!-- Progress with loading spinner -->
<div class="space-y-2">
  <div class="flex justify-between items-center">
    <div class="flex items-center gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
        <path d="M21 12a9 9 0 1 1-6.219-8.56" />
      </svg>
      <span class="text-sm">Processing...</span>
    </div>
    <span class="text-sm">67%</span>
  </div>
  <div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 67%"></div>
  </div>
</div>
```

### Stacked Progress Bars
```html
<!-- Multiple overlapping progress indicators -->
<div class="relative h-3 w-full bg-muted rounded-full overflow-hidden">
  <!-- Background layer -->
  <div class="absolute inset-0 bg-red-200"></div>
  <!-- First layer -->
  <div class="absolute left-0 top-0 h-full bg-red-500" style="width: 30%"></div>
  <!-- Second layer -->  
  <div class="absolute left-0 top-0 h-full bg-yellow-500" style="width: 60%"></div>
  <!-- Third layer -->
  <div class="absolute left-0 top-0 h-full bg-green-500" style="width: 80%"></div>
</div>

<div class="flex justify-between text-xs text-muted-foreground mt-1">
  <span>Critical: 30%</span>
  <span>Warning: 60%</span>
  <span>Normal: 80%</span>
</div>
```

### Circular Progress
```html
<!-- SVG circular progress -->
<div class="relative inline-flex items-center justify-center">
  <svg class="size-16" viewBox="0 0 100 100">
    <!-- Background circle -->
    <circle cx="50" cy="50" r="40" stroke="currentColor" stroke-width="8" fill="none" class="text-muted" />
    <!-- Progress circle -->
    <circle 
      cx="50" cy="50" r="40" 
      stroke="currentColor" 
      stroke-width="8" 
      fill="none" 
      class="text-primary"
      stroke-dasharray="251.2"
      stroke-dashoffset="75.36"
      stroke-linecap="round"
      style="transform: rotate(-90deg); transform-origin: 50% 50%;"
    />
  </svg>
  <span class="absolute text-sm font-medium">70%</span>
</div>

<!-- Simplified circular with CSS -->
<style>
.circular-progress {
  background: conic-gradient(from 0deg, hsl(var(--primary)) 70%, hsl(var(--muted)) 70%);
}
</style>

<div class="circular-progress relative size-16 rounded-full flex items-center justify-center">
  <div class="size-12 bg-background rounded-full flex items-center justify-center">
    <span class="text-sm font-medium">70%</span>
  </div>
</div>
```

## Accessibility Features

- **ARIA Roles**: Use `role="progressbar"` for screen readers
- **Value Attributes**: Provide current, min, and max values
- **Labels**: Include descriptive text or `aria-label`
- **Live Updates**: Use `aria-live` for dynamic updates

### Enhanced Accessibility
```html
<!-- Fully accessible progress bar -->
<div>
  <label id="progress-label" class="text-sm font-medium">File Upload Progress</label>
  <div 
    class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full mt-2"
    role="progressbar"
    aria-labelledby="progress-label"
    aria-valuenow="75"
    aria-valuemin="0"
    aria-valuemax="100"
    aria-live="polite"
  >
    <div 
      class="bg-primary h-full w-full flex-1 transition-all" 
      style="width: 75%"
    ></div>
  </div>
  <div class="flex justify-between text-xs text-muted-foreground mt-1">
    <span>3 of 4 files uploaded</span>
    <span>75% complete</span>
  </div>
</div>

<!-- Progress with description -->
<div 
  role="progressbar"
  aria-label="Installation progress"
  aria-describedby="progress-desc"
  aria-valuenow="45"
  aria-valuemin="0"
  aria-valuemax="100"
  class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"
>
  <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 45%"></div>
</div>
<p id="progress-desc" class="text-sm text-muted-foreground mt-1">
  Installing dependencies... This may take several minutes.
</p>
```

## JavaScript Integration

### Dynamic Progress Updates
```javascript
// Update progress bar
function updateProgress(elementId, percentage) {
  const progressBar = document.getElementById(elementId);
  const progressFill = progressBar.querySelector('[style*="width"]');
  
  // Update visual
  progressFill.style.width = percentage + '%';
  
  // Update accessibility
  progressBar.setAttribute('aria-valuenow', percentage);
}

// Example usage
updateProgress('my-progress', 75);

// Animated progress update
function animateProgress(elementId, targetPercentage, duration = 1000) {
  const progressBar = document.getElementById(elementId);
  const progressFill = progressBar.querySelector('[style*="width"]');
  const currentWidth = parseInt(progressFill.style.width) || 0;
  
  const startTime = Date.now();
  const difference = targetPercentage - currentWidth;
  
  function animate() {
    const elapsed = Date.now() - startTime;
    const progress = Math.min(elapsed / duration, 1);
    const currentValue = currentWidth + (difference * progress);
    
    progressFill.style.width = currentValue + '%';
    progressBar.setAttribute('aria-valuenow', Math.round(currentValue));
    
    if (progress < 1) {
      requestAnimationFrame(animate);
    }
  }
  
  requestAnimationFrame(animate);
}
```

### File Upload Progress
```html
<div id="upload-progress" class="space-y-2" style="display: none;">
  <div class="flex justify-between text-sm">
    <span>Uploading...</span>
    <span id="upload-percentage">0%</span>
  </div>
  <div 
    class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"
    role="progressbar"
    aria-label="File upload progress"
    aria-valuenow="0"
    aria-valuemin="0"
    aria-valuemax="100"
  >
    <div id="upload-fill" class="bg-primary h-full w-full flex-1 transition-all" style="width: 0%"></div>
  </div>
</div>

<script>
// Simulate file upload progress
function simulateUpload() {
  const progressContainer = document.getElementById('upload-progress');
  const progressFill = document.getElementById('upload-fill');
  const progressBar = progressContainer.querySelector('[role="progressbar"]');
  const progressText = document.getElementById('upload-percentage');
  
  progressContainer.style.display = 'block';
  let progress = 0;
  
  const interval = setInterval(() => {
    progress += Math.random() * 10;
    if (progress >= 100) {
      progress = 100;
      clearInterval(interval);
    }
    
    const roundedProgress = Math.round(progress);
    progressFill.style.width = roundedProgress + '%';
    progressBar.setAttribute('aria-valuenow', roundedProgress);
    progressText.textContent = roundedProgress + '%';
    
    if (progress === 100) {
      setTimeout(() => {
        progressContainer.style.display = 'none';
      }, 1000);
    }
  }, 200);
}
</script>
```

## Best Practices

1. **Meaningful Progress**: Only show progress for operations that take time
2. **Accurate Values**: Ensure progress accurately reflects completion
3. **Clear Labels**: Provide context about what's progressing  
4. **Accessibility**: Include ARIA attributes and live updates
5. **Visual Feedback**: Use appropriate colors and animations
6. **Responsive Design**: Ensure progress bars work on all screen sizes
7. **Error Handling**: Handle failed operations gracefully
8. **Performance**: Avoid excessive DOM updates

## Common Patterns

### Form Completion
```html
<div class="space-y-4">
  <div class="flex justify-between items-center">
    <h3 class="text-lg font-semibold">Complete Your Profile</h3>
    <span class="text-sm text-muted-foreground">3 of 5 steps</span>
  </div>
  
  <div class="bg-muted relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 60%"></div>
  </div>
  
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>Personal Info</span>
    <span>Work Experience</span>  
    <span>Skills</span>
    <span>Portfolio</span>
    <span>Review</span>
  </div>
</div>
```

### Download Progress
```html
<div class="border rounded-lg p-4">
  <div class="flex items-center gap-3 mb-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
      <polyline points="7,10 12,15 17,10" />
      <line x1="12" x2="12" y1="15" y2="3" />
    </svg>
    <div class="flex-1">
      <h4 class="font-medium">document.pdf</h4>
      <p class="text-sm text-muted-foreground">2.4 MB of 5.1 MB</p>
    </div>
  </div>
  
  <div class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full">
    <div class="bg-primary h-full w-full flex-1 transition-all" style="width: 47%"></div>
  </div>
  
  <div class="flex justify-between text-xs text-muted-foreground mt-2">
    <span>47% complete</span>
    <span>2 minutes remaining</span>
  </div>
</div>
```

## Integration Examples

### React Integration
```jsx
import React, { useState, useEffect } from 'react';

function ProgressBar({ value = 0, max = 100, label, className = '' }) {
  const percentage = Math.round((value / max) * 100);
  
  return (
    <div className={className}>
      {label && (
        <div className="flex justify-between text-sm mb-2">
          <span>{label}</span>
          <span>{percentage}%</span>
        </div>
      )}
      <div 
        className="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"
        role="progressbar"
        aria-label={label}
        aria-valuenow={percentage}
        aria-valuemin="0"
        aria-valuemax="100"
      >
        <div 
          className="bg-primary h-full w-full flex-1 transition-all"
          style={{ width: `${percentage}%` }}
        />
      </div>
    </div>
  );
}

// Usage
function App() {
  const [progress, setProgress] = useState(0);
  
  useEffect(() => {
    const timer = setInterval(() => {
      setProgress(prev => {
        if (prev >= 100) {
          clearInterval(timer);
          return 100;
        }
        return prev + 10;
      });
    }, 500);
    
    return () => clearInterval(timer);
  }, []);
  
  return (
    <ProgressBar 
      value={progress} 
      label="Loading..." 
      className="w-80" 
    />
  );
}
```

### Vue Integration
```vue
<template>
  <div :class="className">
    <div v-if="label" class="flex justify-between text-sm mb-2">
      <span>{{ label }}</span>
      <span>{{ percentage }}%</span>
    </div>
    <div 
      class="bg-primary/20 relative h-2 w-full overflow-hidden rounded-full"
      role="progressbar"
      :aria-label="label"
      :aria-valuenow="percentage"
      aria-valuemin="0"
      aria-valuemax="100"
    >
      <div 
        class="bg-primary h-full w-full flex-1 transition-all"
        :style="{ width: percentage + '%' }"
      />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: Number,
      default: 0
    },
    max: {
      type: Number,
      default: 100
    },
    label: String,
    className: String
  },
  computed: {
    percentage() {
      return Math.round((this.value / this.max) * 100);
    }
  }
};
</script>
```

## Related Components

- [Spinner](./spinner.md) - For indeterminate loading states
- [Button](./button.md) - For triggering progressive actions
- [Card](./card.md) - For containing progress indicators
- [Toast](./toast.md) - For progress notifications
---
## skeleton

# Skeleton Component

Used to show a placeholder while content is loading.

## Important Note

**There is no dedicated Skeleton component in Basecoat.** Simply use the `animate-pulse` class to create a skeleton loader.

## Basic Usage

```html
<div class="flex items-center gap-4">
  <div class="bg-accent animate-pulse size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent animate-pulse rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent animate-pulse rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard Tailwind utilities

### Supporting Classes
- **`animate-pulse`** - Applies pulsing animation
- **`bg-accent`** - Background color for skeleton elements
- **Size classes**: `h-*`, `w-*`, `size-*` for dimensions
- **Shape classes**: `rounded-*` for different shapes

### Tailwind Utilities Used
- `animate-pulse` - Pulsing animation effect
- `bg-accent` - Light background color
- `bg-muted` - Alternative muted background
- `h-4` - Height (16px)
- `w-[150px]` - Fixed width
- `size-10` - Square dimensions (40x40px)
- `rounded-full` - Circular shape
- `rounded-md` - Rounded rectangle
- `shrink-0` - Prevents shrinking in flex containers

## Component Attributes

### Element Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "animate-pulse" and bg color | Yes |
| `aria-hidden` | boolean | Should be "true" for decorative elements | Recommended |

### No JavaScript Required
This is a pure CSS implementation using Tailwind's animation utilities.

## HTML Structure

```html
<!-- Basic skeleton element -->
<div class="bg-accent animate-pulse [shape-classes] [size-classes]"></div>

<!-- Skeleton group -->
<div class="[layout-classes]">
  <div class="bg-accent animate-pulse [avatar-shape]"></div>
  <div class="[content-layout]">
    <div class="bg-accent animate-pulse [text-line-size]"></div>
    <div class="bg-accent animate-pulse [text-line-size]"></div>
  </div>
</div>
```

## Examples

### Text Skeletons
```html
<!-- Single line -->
<div class="bg-accent animate-pulse h-4 w-48 rounded-md"></div>

<!-- Multiple lines -->
<div class="space-y-2">
  <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
  <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
  <div class="bg-accent animate-pulse h-4 w-3/5 rounded-md"></div>
</div>

<!-- Different text sizes -->
<div class="space-y-3">
  <!-- Heading -->
  <div class="bg-accent animate-pulse h-6 w-2/3 rounded-md"></div>
  <!-- Paragraph -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-11/12 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
  </div>
</div>
```

### Avatar Skeletons
```html
<!-- Circular avatar -->
<div class="bg-accent animate-pulse size-10 rounded-full"></div>

<!-- Square avatar -->
<div class="bg-accent animate-pulse size-10 rounded-lg"></div>

<!-- Different sizes -->
<div class="flex items-center gap-3">
  <div class="bg-accent animate-pulse size-6 rounded-full"></div>
  <div class="bg-accent animate-pulse size-8 rounded-full"></div>
  <div class="bg-accent animate-pulse size-10 rounded-full"></div>
  <div class="bg-accent animate-pulse size-12 rounded-full"></div>
</div>
```

### User Profile Skeleton
```html
<div class="flex items-center gap-4">
  <!-- Avatar -->
  <div class="bg-accent animate-pulse size-12 shrink-0 rounded-full"></div>
  <!-- Content -->
  <div class="flex-1 space-y-2">
    <!-- Name -->
    <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
    <!-- Description -->
    <div class="bg-accent animate-pulse h-3 w-24 rounded-md"></div>
  </div>
</div>

<!-- Extended profile -->
<div class="space-y-4">
  <div class="flex items-center gap-4">
    <div class="bg-accent animate-pulse size-16 shrink-0 rounded-full"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-5 w-40 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-28 rounded-md"></div>
    </div>
  </div>
  <!-- Bio -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-5/6 rounded-md"></div>
  </div>
</div>
```

### Card Skeleton
```html
<div class="card w-full">
  <header>
    <!-- Title -->
    <div class="bg-accent animate-pulse rounded-md h-4 w-2/3"></div>
    <!-- Subtitle -->
    <div class="bg-accent animate-pulse rounded-md h-4 w-1/2"></div>
  </header>
  <section>
    <!-- Main content area -->
    <div class="bg-accent animate-pulse rounded-md aspect-square w-full"></div>
  </section>
</div>

<!-- Card with actions -->
<div class="card">
  <header>
    <div class="flex justify-between items-start">
      <div class="space-y-2">
        <div class="bg-accent animate-pulse h-5 w-48 rounded-md"></div>
        <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      </div>
      <div class="bg-accent animate-pulse h-8 w-20 rounded-md"></div>
    </div>
  </header>
  <section>
    <div class="space-y-3">
      <div class="bg-accent animate-pulse h-4 w-full rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-4/5 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-3/5 rounded-md"></div>
    </div>
  </section>
  <footer>
    <div class="flex justify-between items-center">
      <div class="bg-accent animate-pulse h-8 w-16 rounded-md"></div>
      <div class="bg-accent animate-pulse h-8 w-24 rounded-md"></div>
    </div>
  </footer>
</div>
```

### Table Skeleton
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div></th>
        <th><div class="bg-accent animate-pulse h-4 w-18 rounded-md"></div></th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></td>
      </tr>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-36 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-14 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-22 rounded-md"></div></td>
      </tr>
      <tr>
        <td><div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-30 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div></td>
        <td><div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div></td>
      </tr>
    </tbody>
  </table>
</div>

<!-- Simplified table skeleton -->
<div class="space-y-3">
  <!-- Header -->
  <div class="flex gap-4">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
  </div>
  <!-- Rows -->
  <div class="space-y-2">
    <div class="flex gap-4">
      <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-32 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
    </div>
    <div class="flex gap-4">
      <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-12 rounded-md"></div>
      <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    </div>
  </div>
</div>
```

### Image Skeleton
```html
<!-- Basic image placeholder -->
<div class="bg-accent animate-pulse w-full h-48 rounded-lg"></div>

<!-- Different aspect ratios -->
<div class="space-y-4">
  <!-- Square -->
  <div class="bg-accent animate-pulse aspect-square w-full rounded-lg"></div>
  <!-- 16:9 Video -->
  <div class="bg-accent animate-pulse aspect-video w-full rounded-lg"></div>
  <!-- 4:3 Photo -->
  <div class="bg-accent animate-pulse aspect-[4/3] w-full rounded-lg"></div>
</div>

<!-- Image gallery skeleton -->
<div class="grid grid-cols-2 md:grid-cols-3 gap-4">
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
  <div class="bg-accent animate-pulse aspect-square rounded-lg"></div>
</div>
```

### Form Skeleton
```html
<div class="space-y-4">
  <!-- Form field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  
  <!-- Another field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  
  <!-- Textarea field -->
  <div class="space-y-2">
    <div class="bg-accent animate-pulse h-4 w-28 rounded-md"></div>
    <div class="bg-accent animate-pulse h-24 w-full rounded-md"></div>
  </div>
  
  <!-- Submit button -->
  <div class="bg-accent animate-pulse h-10 w-24 rounded-md"></div>
</div>

<!-- Inline form -->
<div class="flex gap-3 items-end">
  <div class="flex-1 space-y-2">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-10 w-full rounded-md"></div>
  </div>
  <div class="bg-accent animate-pulse h-10 w-20 rounded-md"></div>
</div>
```

### Navigation Skeleton
```html
<!-- Header navigation -->
<div class="flex items-center justify-between p-4 border-b">
  <!-- Logo -->
  <div class="bg-accent animate-pulse h-8 w-32 rounded-md"></div>
  
  <!-- Navigation links -->
  <div class="hidden md:flex items-center gap-6">
    <div class="bg-accent animate-pulse h-4 w-16 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-18 rounded-md"></div>
    <div class="bg-accent animate-pulse h-4 w-24 rounded-md"></div>
  </div>
  
  <!-- User menu -->
  <div class="bg-accent animate-pulse size-8 rounded-full"></div>
</div>

<!-- Sidebar navigation -->
<div class="w-64 space-y-2 p-4">
  <!-- Logo -->
  <div class="bg-accent animate-pulse h-6 w-28 rounded-md mb-6"></div>
  
  <!-- Nav items -->
  <div class="space-y-1">
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
  </div>
  
  <!-- Section -->
  <div class="pt-4">
    <div class="bg-accent animate-pulse h-4 w-20 rounded-md mb-2"></div>
    <div class="space-y-1">
      <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
      <div class="bg-accent animate-pulse h-8 w-full rounded-md"></div>
    </div>
  </div>
</div>
```

### List Skeleton
```html
<!-- Simple list -->
<div class="space-y-3">
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-48 rounded-md"></div>
  </div>
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-40 rounded-md"></div>
  </div>
  <div class="flex items-center gap-3">
    <div class="bg-accent animate-pulse size-6 rounded-full shrink-0"></div>
    <div class="bg-accent animate-pulse h-4 w-52 rounded-md"></div>
  </div>
</div>

<!-- Detailed list items -->
<div class="space-y-4">
  <div class="flex items-start gap-3">
    <div class="bg-accent animate-pulse size-10 rounded-lg shrink-0"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-4 w-3/4 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-1/2 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-2/3 rounded-md"></div>
    </div>
    <div class="bg-accent animate-pulse h-6 w-16 rounded-md"></div>
  </div>
  
  <div class="flex items-start gap-3">
    <div class="bg-accent animate-pulse size-10 rounded-lg shrink-0"></div>
    <div class="flex-1 space-y-2">
      <div class="bg-accent animate-pulse h-4 w-2/3 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-3/5 rounded-md"></div>
      <div class="bg-accent animate-pulse h-3 w-4/5 rounded-md"></div>
    </div>
    <div class="bg-accent animate-pulse h-6 w-20 rounded-md"></div>
  </div>
</div>
```

### Alternative Styling
```html
<!-- Using muted background -->
<div class="flex items-center gap-4">
  <div class="bg-muted animate-pulse size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-muted animate-pulse rounded-md h-4 w-[150px]"></div>
    <div class="bg-muted animate-pulse rounded-md h-4 w-[100px]"></div>
  </div>
</div>

<!-- Custom gradient skeleton -->
<style>
.skeleton-gradient {
  background: linear-gradient(-90deg, #e0e0e0 25%, #f0f0f0 50%, #e0e0e0 75%);
  background-size: 400% 100%;
  animation: loading 1.4s ease-in-out infinite;
}

@keyframes loading {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}
</style>

<div class="flex items-center gap-4">
  <div class="skeleton-gradient size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="skeleton-gradient rounded-md h-4 w-[150px]"></div>
    <div class="skeleton-gradient rounded-md h-4 w-[100px]"></div>
  </div>
</div>

<!-- Without animation (static) -->
<div class="flex items-center gap-4">
  <div class="bg-accent size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## Accessibility Features

- **Hidden from Screen Readers**: Use `aria-hidden="true"` on skeleton elements
- **Loading State**: Provide context about loading state
- **Alternative Text**: Consider providing loading feedback
- **Reduced Motion**: Respect user motion preferences

### Enhanced Accessibility
```html
<!-- With proper ARIA attributes -->
<div aria-live="polite" aria-busy="true">
  <span class="sr-only">Loading content...</span>
  <div class="flex items-center gap-4" aria-hidden="true">
    <div class="bg-accent animate-pulse size-10 shrink-0 rounded-full"></div>
    <div class="grid gap-2">
      <div class="bg-accent animate-pulse rounded-md h-4 w-[150px]"></div>
      <div class="bg-accent animate-pulse rounded-md h-4 w-[100px]"></div>
    </div>
  </div>
</div>

<!-- Reduced motion support -->
<div class="flex items-center gap-4">
  <div class="bg-accent animate-pulse motion-reduce:animate-none size-10 shrink-0 rounded-full"></div>
  <div class="grid gap-2">
    <div class="bg-accent animate-pulse motion-reduce:animate-none rounded-md h-4 w-[150px]"></div>
    <div class="bg-accent animate-pulse motion-reduce:animate-none rounded-md h-4 w-[100px]"></div>
  </div>
</div>
```

## JavaScript Integration

### React Skeleton Component
```jsx
import React from 'react';

function Skeleton({ className = '', width, height, circle = false, ...props }) {
  const baseClasses = "bg-accent animate-pulse";
  const shapeClasses = circle ? "rounded-full" : "rounded-md";
  const sizeClasses = width || height ? "" : "h-4 w-20";
  
  const style = {};
  if (width) style.width = width;
  if (height) style.height = height;
  
  return (
    <div 
      className={`${baseClasses} ${shapeClasses} ${sizeClasses} ${className}`}
      style={style}
      aria-hidden="true"
      {...props}
    />
  );
}

// Usage
<div className="flex items-center gap-4">
  <Skeleton circle width="40px" height="40px" />
  <div className="space-y-2">
    <Skeleton width="150px" />
    <Skeleton width="100px" />
  </div>
</div>

// Text skeleton component
function TextSkeleton({ lines = 1, className = "" }) {
  return (
    <div className={`space-y-2 ${className}`}>
      {Array.from({ length: lines }).map((_, i) => (
        <Skeleton 
          key={i} 
          width={i === lines - 1 ? "80%" : "100%"} 
        />
      ))}
    </div>
  );
}
```

### Vue Skeleton Component
```vue
<template>
  <div 
    :class="skeletonClasses"
    :style="skeletonStyle"
    aria-hidden="true"
    v-bind="$attrs"
  />
</template>

<script>
export default {
  props: {
    width: String,
    height: String,
    circle: Boolean,
    className: String
  },
  computed: {
    skeletonClasses() {
      const base = "bg-accent animate-pulse";
      const shape = this.circle ? "rounded-full" : "rounded-md";
      const size = this.width || this.height ? "" : "h-4 w-20";
      return `${base} ${shape} ${size} ${this.className || ""}`;
    },
    skeletonStyle() {
      const style = {};
      if (this.width) style.width = this.width;
      if (this.height) style.height = this.height;
      return style;
    }
  }
};
</script>
```

### Dynamic Skeleton Loading
```javascript
// Show/hide skeleton based on loading state
function toggleSkeleton(containerId, isLoading) {
  const container = document.getElementById(containerId);
  const skeleton = container.querySelector('.skeleton-wrapper');
  const content = container.querySelector('.content-wrapper');
  
  if (isLoading) {
    skeleton.style.display = 'block';
    content.style.display = 'none';
    container.setAttribute('aria-busy', 'true');
  } else {
    skeleton.style.display = 'none';
    content.style.display = 'block';
    container.setAttribute('aria-busy', 'false');
  }
}

// Usage
toggleSkeleton('user-profile', true);  // Show skeleton
setTimeout(() => {
  toggleSkeleton('user-profile', false); // Show content
}, 2000);
```

## Best Practices

1. **Match Content Structure**: Skeleton should mirror the layout of actual content
2. **Consistent Sizing**: Use similar dimensions to prevent layout shifts
3. **Appropriate Animation**: Use pulse animation for better UX
4. **Reduced Motion**: Respect user motion preferences
5. **Semantic HTML**: Use appropriate ARIA attributes
6. **Performance**: Avoid complex nested skeletons
7. **Visual Hierarchy**: Maintain visual hierarchy with skeleton elements
8. **Loading Context**: Provide context about what's loading

## Common Patterns

### Progressive Loading
```html
<!-- Initially show skeleton -->
<div id="content-container" aria-busy="true">
  <div class="skeleton-wrapper">
    <div class="flex items-center gap-4" aria-hidden="true">
      <div class="bg-accent animate-pulse size-12 shrink-0 rounded-full"></div>
      <div class="flex-1 space-y-2">
        <div class="bg-accent animate-pulse h-5 w-2/3 rounded-md"></div>
        <div class="bg-accent animate-pulse h-4 w-1/2 rounded-md"></div>
      </div>
    </div>
  </div>
  
  <!-- Content loads here -->
  <div class="content-wrapper" style="display: none;">
    <!-- Actual content -->
  </div>
</div>
```

### Staggered Loading
```css
/* Staggered animation delays */
.skeleton-1 { animation-delay: 0s; }
.skeleton-2 { animation-delay: 0.1s; }
.skeleton-3 { animation-delay: 0.2s; }
.skeleton-4 { animation-delay: 0.3s; }
```

```html
<div class="space-y-3">
  <div class="bg-accent animate-pulse skeleton-1 h-4 w-full rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-2 h-4 w-5/6 rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-3 h-4 w-4/5 rounded-md"></div>
  <div class="bg-accent animate-pulse skeleton-4 h-4 w-3/4 rounded-md"></div>
</div>
```

## Related Components

- [Spinner](./spinner.md) - For indeterminate loading states  
- [Progress](./progress.md) - For determinate loading states
- [Card](./card.md) - For skeleton card layouts
- [Avatar](./avatar.md) - For skeleton avatar shapes
---
## spinner

# Spinner Component

An indicator that can be used to show a loading state.

## Important Note

**There is no dedicated Spinner component in Basecoat.** Spinners are pure HTML using the Lucide `loader-circle` icon with the `animate-spin` Tailwind utility.

## Basic Usage

```html
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## CSS Classes

### Primary Classes
- **No dedicated classes** - Uses standard SVG with Tailwind utilities

### Supporting Classes
- **`animate-spin`** - Applies continuous rotation animation
- **`size-*`** - Controls spinner dimensions
- **`text-*`** - Controls spinner color

### Tailwind Utilities Used
- `animate-spin` - Continuous rotation animation
- `size-3` - 12px dimensions (0.75rem)
- `size-4` - 16px dimensions (1rem) 
- `size-6` - 24px dimensions (1.5rem)
- `size-8` - 32px dimensions (2rem)
- `text-*` - Color utilities (text-red-500, text-primary, etc.)

## Component Attributes

### SVG Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `xmlns` | string | SVG namespace | Yes |
| `width` | string | SVG width (usually "24") | Yes |
| `height` | string | SVG height (usually "24") | Yes |
| `viewBox` | string | SVG viewBox (usually "0 0 24 24") | Yes |
| `fill` | string | Should be "none" for outline style | Yes |
| `stroke` | string | Should be "currentColor" | Yes |
| `stroke-width` | string | Stroke thickness (usually "2") | Yes |
| `stroke-linecap` | string | Line cap style (usually "round") | Yes |
| `stroke-linejoin` | string | Line join style (usually "round") | Yes |
| `role` | string | Should be "status" for accessibility | Recommended |
| `aria-label` | string | Should be "Loading" or descriptive | Recommended |
| `class` | string | Must include "animate-spin" | Yes |

### No JavaScript Required
This is a pure CSS/HTML implementation using SVG animation.

## HTML Structure

```html
<svg xmlns="http://www.w3.org/2000/svg" 
     width="24" height="24" 
     viewBox="0 0 24 24" 
     fill="none" 
     stroke="currentColor" 
     stroke-width="2" 
     stroke-linecap="round" 
     stroke-linejoin="round" 
     role="status" 
     aria-label="Loading" 
     class="animate-spin [size-classes] [color-classes]">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## Examples

### Different Sizes
```html
<!-- Extra Small (12px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-3 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Small (16px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Medium (24px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Large (32px) -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-8 animate-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

### Different Colors
```html
<!-- Red spinner -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-red-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Green spinner -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-green-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Primary color -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-primary">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Muted color -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-muted-foreground">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

### Button with Spinner
```html
<!-- Primary button with spinner -->
<button class="btn-sm" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Loading...
</button>

<!-- Outline button with spinner -->
<button class="btn-sm-outline" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Please wait
</button>

<!-- Secondary button with spinner -->
<button class="btn-sm-secondary" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Processing
</button>
```

### Spinner Only (No Text)
```html
<!-- Icon button with spinner only -->
<button class="btn-icon" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>

<!-- Small icon button -->
<button class="btn-icon-sm" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>
```

### Card Loading State
```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-transparent bg-muted/50 p-4 gap-4">
  <div class="flex shrink-0 items-center justify-center [&_svg]:pointer-events-none [&_svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium line-clamp-1">Processing payment...</h3>
  </div>
  <div class="flex flex-col gap-1 flex-none text-center">
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance">$100.00</p>
  </div>
</article>
```

### Inline Spinner
```html
<!-- Inline with text -->
<p class="flex items-center gap-2">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Updating content...
</p>

<!-- Inline with different alignments -->
<div class="flex items-start gap-2">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin mt-0.5">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  <div>
    <p class="font-medium">Loading data...</p>
    <p class="text-sm text-muted-foreground">This may take a moment</p>
  </div>
</div>
```

### Full Page Loading
```html
<!-- Full screen loading overlay -->
<div class="fixed inset-0 bg-background/80 backdrop-blur-sm z-50 flex items-center justify-center">
  <div class="flex flex-col items-center gap-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-8 animate-spin text-primary">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <p class="text-muted-foreground">Loading application...</p>
  </div>
</div>

<!-- Centered in container -->
<div class="flex items-center justify-center h-64">
  <div class="text-center">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin text-muted-foreground mx-auto mb-2">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <p class="text-sm text-muted-foreground">Loading content...</p>
  </div>
</div>
```

### Input Loading State
```html
<!-- Input with spinner -->
<div class="relative">
  <input type="text" class="input pr-10" placeholder="Search..." disabled />
  <div class="absolute inset-y-0 right-0 flex items-center pr-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Searching" class="size-4 animate-spin text-muted-foreground">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
</div>

<!-- Select with spinner -->
<div class="relative">
  <select class="select pr-10" disabled>
    <option>Loading options...</option>
  </select>
  <div class="absolute inset-y-0 right-8 flex items-center pr-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin text-muted-foreground">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  </div>
</div>
```

### Custom Animation Speed
```html
<!-- Faster animation -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin animation-duration-500">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Slower animation -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin animation-duration-2000">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- Custom CSS for animation timing -->
<style>
  .animate-slow-spin {
    animation: spin 2s linear infinite;
  }
  .animate-fast-spin {
    animation: spin 0.5s linear infinite;
  }
</style>

<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-slow-spin">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>
```

## Accessibility Features

- **Role Attribute**: Use `role="status"` to announce loading states
- **ARIA Label**: Provide descriptive `aria-label` for screen readers
- **Live Regions**: Can be used with `aria-live` regions for dynamic updates
- **Focus Management**: Consider focus management when loading states change

### Enhanced Accessibility
```html
<!-- With live region -->
<div aria-live="polite" aria-atomic="true">
  <div class="flex items-center gap-2">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading user data" class="size-4 animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <span>Loading user data...</span>
  </div>
</div>

<!-- With hidden text for screen readers -->
<button class="btn" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-hidden="true" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  <span class="sr-only">Loading, please wait</span>
  <span aria-hidden="true">Loading...</span>
</button>
```

## Performance Considerations

### Conditional Rendering
```html
<!-- Show spinner only when loading -->
<div>
  <!-- Loading state -->
  <div class="flex items-center gap-2" style="display: none;" id="loading-state">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-4 animate-spin">
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    Loading...
  </div>
  
  <!-- Content state -->
  <div id="content-state">
    Content loaded successfully!
  </div>
</div>

<script>
function showLoading() {
  document.getElementById('loading-state').style.display = 'flex';
  document.getElementById('content-state').style.display = 'none';
}

function hideLoading() {
  document.getElementById('loading-state').style.display = 'none';
  document.getElementById('content-state').style.display = 'block';
}
</script>
```

### Reduced Motion Support
```html
<!-- Respects user's motion preferences -->
<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="size-6 animate-spin motion-reduce:animate-none">
  <path d="M21 12a9 9 0 1 1-6.219-8.56" />
</svg>

<!-- CSS alternative -->
<style>
  @media (prefers-reduced-motion: reduce) {
    .spinner-respectful {
      animation: none;
    }
  }
</style>
```

## Best Practices

1. **Descriptive Labels**: Use meaningful `aria-label` attributes
2. **Appropriate Sizing**: Match spinner size to context
3. **Color Contrast**: Ensure sufficient contrast for visibility
4. **Loading States**: Clear communication of what's loading
5. **Timeout Handling**: Provide fallbacks for long loading times
6. **Reduced Motion**: Respect user motion preferences
7. **Performance**: Only show spinners when necessary
8. **User Feedback**: Combine with progress indicators when possible

## Size Reference

| Class | Size | Pixels | Use Case |
|-------|------|--------|----------|
| `size-3` | 0.75rem | 12px | Tiny indicators, inline text |
| `size-4` | 1rem | 16px | Small buttons, input fields |
| `size-5` | 1.25rem | 20px | Medium inline elements |
| `size-6` | 1.5rem | 24px | Standard buttons, cards |
| `size-8` | 2rem | 32px | Large buttons, prominent loading |
| `size-10` | 2.5rem | 40px | Modal dialogs, major sections |
| `size-12` | 3rem | 48px | Full-page loading |

## Common Patterns

### Form Submission
```html
<form onsubmit="showLoadingState()">
  <div class="space-y-4">
    <!-- Form fields -->
    <input type="email" class="input" placeholder="Email" required />
    
    <!-- Submit button with loading state -->
    <button type="submit" class="btn" id="submit-btn">
      <span class="flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading" class="animate-spin hidden" id="submit-spinner">
          <path d="M21 12a9 9 0 1 1-6.219-8.56" />
        </svg>
        <span id="submit-text">Submit</span>
      </span>
    </button>
  </div>
</form>

<script>
function showLoadingState() {
  const btn = document.getElementById('submit-btn');
  const spinner = document.getElementById('submit-spinner');
  const text = document.getElementById('submit-text');
  
  btn.disabled = true;
  spinner.classList.remove('hidden');
  text.textContent = 'Submitting...';
}
</script>
```

### Data Table Loading
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody id="table-body">
      <!-- Loading state -->
      <tr>
        <td colspan="3" class="text-center py-8">
          <div class="flex flex-col items-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="status" aria-label="Loading data" class="size-6 animate-spin text-muted-foreground">
              <path d="M21 12a9 9 0 1 1-6.219-8.56" />
            </svg>
            <p class="text-muted-foreground">Loading data...</p>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Spinner({ size = 'size-4', color = 'text-current', className = '', ...props }) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
      role="status"
      aria-label="Loading"
      className={`animate-spin ${size} ${color} ${className}`}
      {...props}
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
  );
}

// Usage
<Spinner size="size-6" color="text-primary" />
<button disabled>
  <Spinner size="size-4" />
  Loading...
</button>
```

### Vue Integration
```vue
<template>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    role="status"
    aria-label="Loading"
    :class="spinnerClasses"
    v-bind="$attrs"
  >
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</template>

<script>
export default {
  props: {
    size: {
      type: String,
      default: 'size-4'
    },
    color: {
      type: String,
      default: 'text-current'
    }
  },
  computed: {
    spinnerClasses() {
      return `animate-spin ${this.size} ${this.color}`;
    }
  }
};
</script>
```

### Web Components
```javascript
class SpinnerElement extends HTMLElement {
  constructor() {
    super();
    this.attachShadow({ mode: 'open' });
  }
  
  connectedCallback() {
    const size = this.getAttribute('size') || 'size-4';
    const color = this.getAttribute('color') || 'text-current';
    
    this.shadowRoot.innerHTML = `
      <style>
        @import url('path/to/tailwind.css');
        .spinner { animation: spin 1s linear infinite; }
        @keyframes spin {
          from { transform: rotate(0deg); }
          to { transform: rotate(360deg); }
        }
      </style>
      <svg xmlns="http://www.w3.org/2000/svg" 
           width="24" height="24" 
           viewBox="0 0 24 24" 
           fill="none" 
           stroke="currentColor" 
           stroke-width="2" 
           stroke-linecap="round" 
           stroke-linejoin="round" 
           role="status" 
           aria-label="Loading"
           class="spinner ${size} ${color}">
        <path d="M21 12a9 9 0 1 1-6.219-8.56" />
      </svg>
    `;
  }
}

customElements.define('loading-spinner', SpinnerElement);
```

## Related Components

- [Button](./button.md) - For loading button states
- [Card](./card.md) - For card loading states
- [Table](./table.md) - For table loading states
- [Progress](./progress.md) - For determinate progress indicators
---
## table

# Table Component

A responsive table component for displaying tabular data with proper styling and structure.

## Basic Usage

```html
<div class="overflow-x-auto">
  <table class="table">
    <caption>A list of your recent invoices.</caption>
    <thead>
      <tr>
        <th>Invoice</th>
        <th>Status</th>
        <th>Method</th>
        <th>Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">INV001</td>
        <td>Paid</td>
        <td>Credit Card</td>
        <td class="text-right">$250.00</td>
      </tr>
      <tr>
        <td class="font-medium">INV002</td>
        <td>Pending</td>
        <td>PayPal</td>
        <td class="text-right">$150.00</td>
      </tr>
    </tbody>
    <tfoot>
      <tr>
        <td colspan="3">Total</td>
        <td class="text-right">$400.00</td>
      </tr>
    </tfoot>
  </table>
</div>
```

## CSS Classes

### Primary Classes
- **`table`** - Applied to the `<table>` element

### Supporting Classes
- **`overflow-x-auto`** - Container for responsive horizontal scrolling
- Text alignment classes (`text-left`, `text-center`, `text-right`)
- Font weight classes (`font-medium`, `font-semibold`)

### Tailwind Utilities Used
- `font-medium` - Medium font weight for headers/important data
- `text-right` - Right-align numerical data
- `text-center` - Center-align data
- `overflow-x-auto` - Horizontal scroll for responsive tables

## Component Attributes

### Table Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "table" | Yes |

### No JavaScript Required
This component is purely CSS-based and does not require JavaScript initialization.

## HTML Structure

```html
<div class="overflow-x-auto">
  <table class="table">
    <caption>Table description (optional)</caption>
    <thead>
      <tr>
        <th>Header 1</th>
        <th>Header 2</th>
        <th>Header 3</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>Data 1</td>
        <td>Data 2</td>
        <td>Data 3</td>
      </tr>
    </tbody>
    <tfoot>
      <tr>
        <td colspan="2">Footer</td>
        <td>Total</td>
      </tr>
    </tfoot>
  </table>
</div>
```

## Examples

### Basic Data Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Role</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td>Admin</td>
        <td>Active</td>
      </tr>
      <tr>
        <td class="font-medium">Jane Smith</td>
        <td>jane@example.com</td>
        <td>Editor</td>
        <td>Active</td>
      </tr>
      <tr>
        <td class="font-medium">Mike Johnson</td>
        <td>mike@example.com</td>
        <td>Viewer</td>
        <td>Inactive</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Actions
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Product</th>
        <th>SKU</th>
        <th>Price</th>
        <th>Stock</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">Wireless Headphones</td>
        <td>WH-001</td>
        <td class="text-right">$99.99</td>
        <td class="text-center">25</td>
        <td>
          <div class="flex gap-2">
            <button type="button" class="btn-ghost text-sm">Edit</button>
            <button type="button" class="btn-ghost text-sm text-destructive">Delete</button>
          </div>
        </td>
      </tr>
      <tr>
        <td class="font-medium">Bluetooth Speaker</td>
        <td>BS-002</td>
        <td class="text-right">$149.99</td>
        <td class="text-center">12</td>
        <td>
          <div class="flex gap-2">
            <button type="button" class="btn-ghost text-sm">Edit</button>
            <button type="button" class="btn-ghost text-sm text-destructive">Delete</button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Status Badges
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Order ID</th>
        <th>Customer</th>
        <th>Date</th>
        <th>Status</th>
        <th>Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">#ORD-001</td>
        <td>Alice Johnson</td>
        <td>2024-01-15</td>
        <td>
          <span class="badge bg-success text-success-foreground">Completed</span>
        </td>
        <td class="text-right">$299.99</td>
      </tr>
      <tr>
        <td class="font-medium">#ORD-002</td>
        <td>Bob Wilson</td>
        <td>2024-01-14</td>
        <td>
          <span class="badge bg-warning text-warning-foreground">Pending</span>
        </td>
        <td class="text-right">$149.50</td>
      </tr>
      <tr>
        <td class="font-medium">#ORD-003</td>
        <td>Carol Davis</td>
        <td>2024-01-13</td>
        <td>
          <span class="badge bg-error text-error-foreground">Cancelled</span>
        </td>
        <td class="text-right">$89.99</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Avatars
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>User</th>
        <th>Department</th>
        <th>Role</th>
        <th>Last Active</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>
          <div class="flex items-center gap-3">
            <img src="/avatar1.jpg" alt="John Doe" class="size-8 rounded-full">
            <div>
              <div class="font-medium">John Doe</div>
              <div class="text-sm text-muted-foreground">john@example.com</div>
            </div>
          </div>
        </td>
        <td>Engineering</td>
        <td>Senior Developer</td>
        <td class="text-muted-foreground">2 minutes ago</td>
      </tr>
      <tr>
        <td>
          <div class="flex items-center gap-3">
            <img src="/avatar2.jpg" alt="Jane Smith" class="size-8 rounded-full">
            <div>
              <div class="font-medium">Jane Smith</div>
              <div class="text-sm text-muted-foreground">jane@example.com</div>
            </div>
          </div>
        </td>
        <td>Design</td>
        <td>UX Designer</td>
        <td class="text-muted-foreground">1 hour ago</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Sortable Table Headers
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Name
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Date
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Amount
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">Transaction 1</td>
        <td>2024-01-15</td>
        <td class="text-right">$250.00</td>
      </tr>
      <tr>
        <td class="font-medium">Transaction 2</td>
        <td>2024-01-14</td>
        <td class="text-right">$175.50</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Selection
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>
          <input type="checkbox" class="checkbox" aria-label="Select all">
        </th>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td>Active</td>
      </tr>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">Jane Smith</td>
        <td>jane@example.com</td>
        <td>Active</td>
      </tr>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">Mike Johnson</td>
        <td>mike@example.com</td>
        <td>Inactive</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Empty State Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Role</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="4" class="text-center py-8">
          <div class="flex flex-col items-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
              <rect width="7" height="7" x="3" y="3" rx="1" />
              <rect width="7" height="7" x="14" y="3" rx="1" />
              <rect width="7" height="7" x="14" y="14" rx="1" />
              <rect width="7" height="7" x="3" y="14" rx="1" />
            </svg>
            <div class="text-muted-foreground">No data found</div>
            <p class="text-sm text-muted-foreground">Get started by adding your first entry.</p>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Compact Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th class="py-2">Product</th>
        <th class="py-2">Price</th>
        <th class="py-2">Stock</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="py-2 font-medium">Item 1</td>
        <td class="py-2 text-right">$99.99</td>
        <td class="py-2 text-center">25</td>
      </tr>
      <tr>
        <td class="py-2 font-medium">Item 2</td>
        <td class="py-2 text-right">$149.99</td>
        <td class="py-2 text-center">12</td>
      </tr>
    </tbody>
  </table>
</div>
```

## Responsive Design

### Horizontal Scroll
```html
<div class="overflow-x-auto">
  <table class="table min-w-full">
    <!-- Table content -->
  </table>
</div>
```

### Stack on Mobile
```html
<div class="block md:hidden">
  <!-- Mobile card view -->
  <div class="space-y-4">
    <div class="border rounded-lg p-4">
      <div class="font-medium">John Doe</div>
      <div class="text-sm text-muted-foreground">john@example.com</div>
      <div class="mt-2 flex justify-between">
        <span>Role: Admin</span>
        <span class="badge">Active</span>
      </div>
    </div>
  </div>
</div>

<div class="hidden md:block">
  <!-- Desktop table view -->
  <div class="overflow-x-auto">
    <table class="table">
      <!-- Full table -->
    </table>
  </div>
</div>
```

### Responsive Columns
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th class="hidden md:table-cell">Phone</th>
        <th class="hidden lg:table-cell">Department</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td class="hidden md:table-cell">(555) 123-4567</td>
        <td class="hidden lg:table-cell">Engineering</td>
        <td>Active</td>
      </tr>
    </tbody>
  </table>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper table markup
- **Screen Reader Support**: Table headers properly associated with data
- **Keyboard Navigation**: Focusable interactive elements
- **ARIA Labels**: Appropriate labels for complex tables

### Enhanced Accessibility
```html
<div class="overflow-x-auto">
  <table class="table" role="table" aria-label="User management data">
    <caption class="sr-only">
      List of users with their roles and status information
    </caption>
    <thead>
      <tr role="row">
        <th role="columnheader" scope="col" aria-sort="none">
          Name
        </th>
        <th role="columnheader" scope="col" aria-sort="none">
          Role
        </th>
        <th role="columnheader" scope="col" aria-sort="ascending">
          Status
        </th>
      </tr>
    </thead>
    <tbody>
      <tr role="row">
        <td role="gridcell">John Doe</td>
        <td role="gridcell">Admin</td>
        <td role="gridcell">
          <span class="sr-only">Status:</span>
          Active
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

## Table Pagination

```html
<div class="space-y-4">
  <!-- Table -->
  <div class="overflow-x-auto">
    <table class="table">
      <!-- Table content -->
    </table>
  </div>
  
  <!-- Pagination -->
  <nav role="navigation" aria-label="Table pagination" class="flex items-center justify-between">
    <div class="text-sm text-muted-foreground">
      Showing 1 to 10 of 97 results
    </div>
    <div class="flex items-center gap-2">
      <button type="button" class="btn-outline" disabled>
        Previous
      </button>
      <button type="button" class="btn">1</button>
      <button type="button" class="btn-outline">2</button>
      <button type="button" class="btn-outline">3</button>
      <span class="px-2">...</span>
      <button type="button" class="btn-outline">10</button>
      <button type="button" class="btn-outline">
        Next
      </button>
    </div>
  </nav>
</div>
```

## JavaScript Functionality

### Sortable Table
```javascript
function makeSortable(table) {
  const headers = table.querySelectorAll('th[data-sortable]');
  
  headers.forEach(header => {
    header.style.cursor = 'pointer';
    header.addEventListener('click', () => {
      const column = header.dataset.sortable;
      const order = header.dataset.order === 'asc' ? 'desc' : 'asc';
      
      // Reset other headers
      headers.forEach(h => h.dataset.order = '');
      header.dataset.order = order;
      
      sortTable(table, column, order);
    });
  });
}

function sortTable(table, column, order) {
  const tbody = table.querySelector('tbody');
  const rows = Array.from(tbody.querySelectorAll('tr'));
  
  rows.sort((a, b) => {
    const aValue = a.querySelector(`[data-column="${column}"]`).textContent;
    const bValue = b.querySelector(`[data-column="${column}"]`).textContent;
    
    if (order === 'asc') {
      return aValue.localeCompare(bValue, undefined, { numeric: true });
    } else {
      return bValue.localeCompare(aValue, undefined, { numeric: true });
    }
  });
  
  rows.forEach(row => tbody.appendChild(row));
}
```

### Row Selection
```javascript
function addRowSelection(table) {
  const selectAll = table.querySelector('thead input[type="checkbox"]');
  const rowCheckboxes = table.querySelectorAll('tbody input[type="checkbox"]');
  
  // Select all functionality
  selectAll.addEventListener('change', () => {
    rowCheckboxes.forEach(checkbox => {
      checkbox.checked = selectAll.checked;
      toggleRowSelection(checkbox.closest('tr'), checkbox.checked);
    });
  });
  
  // Individual row selection
  rowCheckboxes.forEach(checkbox => {
    checkbox.addEventListener('change', () => {
      toggleRowSelection(checkbox.closest('tr'), checkbox.checked);
      updateSelectAllState();
    });
  });
  
  function toggleRowSelection(row, selected) {
    row.classList.toggle('bg-muted/50', selected);
  }
  
  function updateSelectAllState() {
    const checkedCount = [...rowCheckboxes].filter(cb => cb.checked).length;
    selectAll.checked = checkedCount === rowCheckboxes.length;
    selectAll.indeterminate = checkedCount > 0 && checkedCount < rowCheckboxes.length;
  }
}
```

## Best Practices

1. **Clear Headers**: Use descriptive column headers
2. **Consistent Alignment**: Align numerical data right, text left
3. **Responsive Design**: Provide horizontal scroll or alternative layouts
4. **Loading States**: Show skeleton loaders for dynamic data
5. **Empty States**: Provide helpful messages when no data
6. **Row Actions**: Keep action buttons consistent and accessible
7. **Sorting**: Provide clear visual feedback for sortable columns
8. **Pagination**: Break up large datasets appropriately

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Table({ columns, data, onSort, onSelect }) {
  return (
    <div className="overflow-x-auto">
      <table className="table">
        <thead>
          <tr>
            {columns.map((column) => (
              <th key={column.key}>
                {column.sortable ? (
                  <button 
                    onClick={() => onSort(column.key)}
                    className="flex items-center gap-2 font-medium hover:text-foreground"
                  >
                    {column.label}
                    <SortIcon />
                  </button>
                ) : (
                  column.label
                )}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row) => (
            <tr key={row.id}>
              {columns.map((column) => (
                <td key={column.key} className={column.className}>
                  {column.render ? column.render(row[column.key], row) : row[column.key]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
```

### HTMX Integration
```html
<div class="overflow-x-auto">
  <table class="table" 
         hx-get="/api/table-data" 
         hx-trigger="load"
         hx-target="tbody">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="3" class="text-center py-4">
          <div class="skeleton w-full h-4"></div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Vue Integration
```vue
<template>
  <div class="overflow-x-auto">
    <table class="table">
      <thead>
        <tr>
          <th v-for="column in columns" :key="column.key">
            <button 
              v-if="column.sortable"
              @click="sort(column.key)"
              class="flex items-center gap-2 font-medium hover:text-foreground"
            >
              {{ column.label }}
              <SortIcon />
            </button>
            <span v-else>{{ column.label }}</span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in data" :key="row.id">
          <td 
            v-for="column in columns" 
            :key="column.key"
            :class="column.className"
          >
            <slot :name="`cell-${column.key}`" :row="row" :value="row[column.key]">
              {{ row[column.key] }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
```

## Related Components

- [Card](./card.md) - For alternative data display
- [Badge](./badge.md) - For status indicators in tables
- [Button](./button.md) - For table actions
- [Checkbox](./checkbox.md) - For row selection
---
## alert-dialog

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
---
## alert

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
---
## button-group

# Button Group Component

A container that groups related buttons together with consistent styling.

## Basic Usage

```html
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Archive</button>
  <button type="button" class="btn-outline">Report</button>
</div>
```

## CSS Classes

### Primary Classes
- **`button-group`** - Core button group class for connected buttons

### Supporting Classes
- **Layout**: `flex`, `w-fit`, `items-stretch`, `gap-2`
- **Grouping**: `role="group"` for accessibility

### Tailwind Utilities Used
- `flex` - Flexbox layout
- `w-fit` - Width fits content
- `items-stretch` - Equal height buttons
- `gap-2` - Spacing between groups
- `button-group` - Connected button styling

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include layout classes | Yes |
| `role` | string | "group" for accessibility | Recommended |

### Button Group Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | "button-group" class | Yes |
| `role` | string | "group" for semantic grouping | Recommended |

### No JavaScript Required (Basic)
Basic button groups work with pure CSS styling.

## HTML Structure

```html
<!-- Container for multiple groups -->
<div class="flex w-fit items-stretch gap-2">
  <!-- Individual button groups -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Button 1</button>
    <button type="button" class="btn-outline">Button 2</button>
  </div>
</div>
```

## Examples

### Basic Button Groups

```html
<!-- Simple two-button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Save</button>
  <button type="button" class="btn-outline">Cancel</button>
</div>

<!-- Multiple button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Bold</button>
  <button type="button" class="btn-outline">Italic</button>
  <button type="button" class="btn-outline">Underline</button>
</div>

<!-- Mixed button types -->
<div role="group" class="button-group">
  <button type="button" class="btn">Primary</button>
  <button type="button" class="btn-outline">Secondary</button>
</div>
```

### Multiple Button Groups

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- Back button (standalone) -->
  <button type="button" class="btn-icon-outline" aria-label="Go Back">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m12 19-7-7 7-7" />
      <path d="M19 12H5" />
    </svg>
  </button>

  <!-- First button group -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Archive</button>
    <button type="button" class="btn-outline">Report</button>
  </div>

  <!-- Second button group with dropdown -->
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Snooze</button>
    <button type="button" class="btn-icon-outline" aria-label="More options">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
    </button>
  </div>
</div>
```

### Different Button Variants

```html
<!-- Primary button group -->
<div role="group" class="button-group">
  <button type="button" class="btn">Save</button>
  <button type="button" class="btn">Continue</button>
  <button type="button" class="btn">Finish</button>
</div>

<!-- Outline button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Left</button>
  <button type="button" class="btn-outline">Center</button>
  <button type="button" class="btn-outline">Right</button>
</div>

<!-- Secondary button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-secondary">Option A</button>
  <button type="button" class="btn-secondary">Option B</button>
  <button type="button" class="btn-secondary">Option C</button>
</div>

<!-- Ghost button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-ghost">Draft</button>
  <button type="button" class="btn-ghost">Published</button>
  <button type="button" class="btn-ghost">Archived</button>
</div>
```

### Segmented Control Style

```html
<!-- Selection group (radio-like behavior) -->
<div role="group" class="button-group" aria-label="View options">
  <button type="button" class="btn-outline" aria-pressed="false">Day</button>
  <button type="button" class="btn-outline" aria-pressed="true">Week</button>
  <button type="button" class="btn-outline" aria-pressed="false">Month</button>
</div>

<!-- Toggle group -->
<div role="group" class="button-group" aria-label="Text formatting">
  <button type="button" class="btn-outline" aria-pressed="true" aria-label="Bold">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
      <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
    </svg>
  </button>
  <button type="button" class="btn-outline" aria-pressed="false" aria-label="Italic">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="19" x2="10" y1="4" y2="4" />
      <line x1="14" x2="5" y1="20" y2="20" />
      <line x1="15" x2="9" y1="4" y2="20" />
    </svg>
  </button>
  <button type="button" class="btn-outline" aria-pressed="false" aria-label="Underline">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4v6a6 6 0 0 0 12 0V4" />
      <line x1="4" x2="20" y1="20" y2="20" />
    </svg>
  </button>
</div>
```

### Icon Button Groups

```html
<!-- Icon-only button group -->
<div role="group" class="button-group" aria-label="Text alignment">
  <button type="button" class="btn-icon-outline" aria-label="Align left">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="21" x2="3" y1="6" y2="6" />
      <line x1="15" x2="3" y1="12" y2="12" />
      <line x1="17" x2="3" y1="18" y2="18" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Align center">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="18" x2="6" y1="6" y2="6" />
      <line x1="21" x2="3" y1="12" y2="12" />
      <line x1="18" x2="6" y1="18" y2="18" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Align right">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="21" x2="3" y1="6" y2="6" />
      <line x1="21" x2="9" y1="12" y2="12" />
      <line x1="21" x2="7" y1="18" y2="18" />
    </svg>
  </button>
</div>

<!-- Media controls -->
<div role="group" class="button-group" aria-label="Media controls">
  <button type="button" class="btn-icon-outline" aria-label="Previous">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="19,20 9,12 19,4" />
      <line x1="5" x2="5" y1="19" y2="5" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Play">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="5,3 19,12 5,21" />
    </svg>
  </button>
  <button type="button" class="btn-icon-outline" aria-label="Next">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="5,4 15,12 5,20" />
      <line x1="19" x2="19" y1="5" y2="19" />
    </svg>
  </button>
</div>
```

### Small Button Groups

```html
<!-- Small button group -->
<div role="group" class="button-group">
  <button type="button" class="btn-sm-outline">Small</button>
  <button type="button" class="btn-sm-outline">Buttons</button>
  <button type="button" class="btn-sm-outline">Group</button>
</div>

<!-- Small icon group -->
<div role="group" class="button-group">
  <button type="button" class="btn-sm-icon-outline" aria-label="Cut">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="6" cy="6" r="3" />
      <circle cx="6" cy="18" r="3" />
      <line x1="20" x2="8.12" y1="4" y2="15.88" />
      <line x1="14.47" x2="20" y1="14.48" y2="20" />
      <line x1="8.12" x2="14.47" y1="8.12" y2="14.47" />
    </svg>
  </button>
  <button type="button" class="btn-sm-icon-outline" aria-label="Copy">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect width="14" height="14" x="8" y="8" rx="2" ry="2" />
      <path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2" />
    </svg>
  </button>
  <button type="button" class="btn-sm-icon-outline" aria-label="Paste">
    <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <rect width="8" height="4" x="8" y="2" rx="1" ry="1" />
      <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2" />
    </svg>
  </button>
</div>
```

### Button Group with Dropdown

```html
<div role="group" class="button-group">
  <button type="button" class="btn-outline">Snooze</button>
  
  <!-- Dropdown trigger -->
  <div class="dropdown-menu">
    <button type="button" class="btn-icon-outline" aria-haspopup="menu" aria-expanded="false">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
    </button>
    <!-- Dropdown menu content -->
    <div data-popover aria-hidden="true">
      <div role="menu">
        <div role="menuitem">Archive</div>
        <div role="menuitem">Delete</div>
      </div>
    </div>
  </div>
</div>

<!-- Split button pattern -->
<div role="group" class="button-group">
  <button type="button" class="btn">Save</button>
  <div class="dropdown-menu">
    <button type="button" class="btn-icon" aria-haspopup="menu" aria-label="Save options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m6 9 6 6 6-6" />
      </svg>
    </button>
  </div>
</div>
```

### Pagination Button Group

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- Previous button -->
  <button type="button" class="btn-icon-outline" aria-label="Previous page">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m15 18-6-6 6-6" />
    </svg>
  </button>
  
  <!-- Page numbers -->
  <div role="group" class="button-group" aria-label="Page navigation">
    <button type="button" class="btn-outline">1</button>
    <button type="button" class="btn" aria-current="page">2</button>
    <button type="button" class="btn-outline">3</button>
    <button type="button" class="btn-outline">4</button>
    <button type="button" class="btn-outline">5</button>
  </div>
  
  <!-- Next button -->
  <button type="button" class="btn-icon-outline" aria-label="Next page">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m9 18 6-6-6-6" />
    </svg>
  </button>
</div>
```

### Form Button Groups

```html
<!-- Form actions -->
<div class="flex w-fit items-stretch gap-2">
  <button type="button" class="btn-outline">Cancel</button>
  
  <div role="group" class="button-group">
    <button type="submit" class="btn">Save Draft</button>
    <button type="submit" class="btn">Publish</button>
  </div>
</div>

<!-- Modal actions -->
<div role="group" class="button-group">
  <button type="button" class="btn-outline" onclick="closeModal()">Cancel</button>
  <button type="button" class="btn-destructive" onclick="confirmDelete()">Delete</button>
</div>

<!-- Wizard navigation -->
<div class="flex justify-between">
  <button type="button" class="btn-outline">Back</button>
  
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Save Draft</button>
    <button type="submit" class="btn">Continue</button>
  </div>
</div>
```

## Accessibility Features

- **Role Attributes**: Use `role="group"` for semantic grouping
- **Keyboard Navigation**: Standard Tab/Shift+Tab navigation between groups
- **Arrow Key Navigation**: Can implement custom arrow key handling within groups
- **Screen Reader Support**: Groups are announced as units
- **Labels**: Use `aria-label` to describe group purpose

### Enhanced Accessibility

```html
<!-- Labeled button group -->
<div role="group" class="button-group" aria-label="Text formatting options">
  <button type="button" class="btn-outline" aria-pressed="false">Bold</button>
  <button type="button" class="btn-outline" aria-pressed="true">Italic</button>
  <button type="button" class="btn-outline" aria-pressed="false">Underline</button>
</div>

<!-- Group with description -->
<fieldset class="border-none p-0 m-0">
  <legend class="sr-only">Alignment options</legend>
  <div role="group" class="button-group">
    <button type="button" class="btn-outline" aria-pressed="true">Left</button>
    <button type="button" class="btn-outline" aria-pressed="false">Center</button>
    <button type="button" class="btn-outline" aria-pressed="false">Right</button>
  </div>
</fieldset>

<!-- Radio-like group -->
<div role="radiogroup" class="button-group" aria-label="View mode">
  <button type="button" role="radio" class="btn" aria-checked="false">Grid</button>
  <button type="button" role="radio" class="btn-outline" aria-checked="true">List</button>
  <button type="button" role="radio" class="btn-outline" aria-checked="false">Card</button>
</div>
```

## JavaScript Integration

### Toggle Selection

```javascript
// Single selection (radio-like)
function selectButton(button, group) {
  // Remove selection from all buttons in group
  group.querySelectorAll('button').forEach(btn => {
    btn.classList.remove('btn');
    btn.classList.add('btn-outline');
    btn.setAttribute('aria-pressed', 'false');
  });
  
  // Select clicked button
  button.classList.remove('btn-outline');
  button.classList.add('btn');
  button.setAttribute('aria-pressed', 'true');
}

// Multi-selection (checkbox-like)
function toggleButton(button) {
  const isPressed = button.getAttribute('aria-pressed') === 'true';
  button.setAttribute('aria-pressed', (!isPressed).toString());
  
  if (isPressed) {
    button.classList.remove('btn');
    button.classList.add('btn-outline');
  } else {
    button.classList.remove('btn-outline');
    button.classList.add('btn');
  }
}

// Initialize button group
document.querySelectorAll('.button-group').forEach(group => {
  group.addEventListener('click', (e) => {
    if (e.target.tagName === 'BUTTON') {
      const isRadio = group.hasAttribute('data-radio');
      
      if (isRadio) {
        selectButton(e.target, group);
      } else {
        toggleButton(e.target);
      }
    }
  });
});
```

### Keyboard Navigation

```javascript
// Arrow key navigation within groups
function initButtonGroupKeyboard() {
  document.querySelectorAll('.button-group').forEach(group => {
    const buttons = Array.from(group.querySelectorAll('button'));
    
    buttons.forEach((button, index) => {
      button.addEventListener('keydown', (e) => {
        let targetIndex;
        
        switch (e.key) {
          case 'ArrowLeft':
            e.preventDefault();
            targetIndex = index > 0 ? index - 1 : buttons.length - 1;
            buttons[targetIndex].focus();
            break;
            
          case 'ArrowRight':
            e.preventDefault();
            targetIndex = index < buttons.length - 1 ? index + 1 : 0;
            buttons[targetIndex].focus();
            break;
            
          case 'Home':
            e.preventDefault();
            buttons[0].focus();
            break;
            
          case 'End':
            e.preventDefault();
            buttons[buttons.length - 1].focus();
            break;
        }
      });
    });
  });
}

// Initialize on page load
document.addEventListener('DOMContentLoaded', initButtonGroupKeyboard);
```

### React Integration

```jsx
import React, { useState } from 'react';

function ButtonGroup({ children, variant = 'single', onChange, className = '' }) {
  const [selected, setSelected] = useState(variant === 'single' ? 0 : []);
  
  const handleClick = (index) => {
    if (variant === 'single') {
      setSelected(index);
      onChange?.(index);
    } else {
      const newSelected = selected.includes(index)
        ? selected.filter(i => i !== index)
        : [...selected, index];
      setSelected(newSelected);
      onChange?.(newSelected);
    }
  };
  
  return (
    <div role="group" className={`button-group ${className}`}>
      {React.Children.map(children, (child, index) => 
        React.cloneElement(child, {
          'aria-pressed': variant === 'single' 
            ? selected === index 
            : selected.includes(index),
          onClick: () => handleClick(index),
          className: variant === 'single'
            ? (selected === index ? 'btn' : 'btn-outline')
            : (selected.includes(index) ? 'btn' : 'btn-outline')
        })
      )}
    </div>
  );
}

// Usage
function App() {
  return (
    <ButtonGroup variant="single" onChange={(index) => console.log('Selected:', index)}>
      <button type="button">Option 1</button>
      <button type="button">Option 2</button>
      <button type="button">Option 3</button>
    </ButtonGroup>
  );
}
```

### Vue Integration

```vue
<template>
  <div role="group" class="button-group">
    <button
      v-for="(option, index) in options"
      :key="option.id"
      type="button"
      :class="getButtonClass(index)"
      :aria-pressed="isSelected(index)"
      @click="handleClick(index)"
    >
      {{ option.label }}
    </button>
  </div>
</template>

<script>
export default {
  props: {
    options: Array,
    modelValue: [Number, Array],
    multiple: Boolean
  },
  emits: ['update:modelValue'],
  methods: {
    isSelected(index) {
      return this.multiple 
        ? this.modelValue.includes(index)
        : this.modelValue === index;
    },
    getButtonClass(index) {
      return this.isSelected(index) ? 'btn' : 'btn-outline';
    },
    handleClick(index) {
      if (this.multiple) {
        const selected = [...this.modelValue];
        const selectedIndex = selected.indexOf(index);
        
        if (selectedIndex > -1) {
          selected.splice(selectedIndex, 1);
        } else {
          selected.push(index);
        }
        
        this.$emit('update:modelValue', selected);
      } else {
        this.$emit('update:modelValue', index);
      }
    }
  }
};
</script>
```

## Best Practices

1. **Logical Grouping**: Group related actions together
2. **Consistent Variants**: Use same button variant within a group
3. **Accessibility**: Include proper ARIA attributes
4. **Visual Separation**: Use gaps between different button groups
5. **Icon Guidelines**: Use consistent icon sizes within groups
6. **Responsive Design**: Consider button group behavior on mobile
7. **Clear Labels**: Provide descriptive labels for screen readers
8. **State Management**: Handle selection states appropriately

## Common Patterns

### Toolbar Groups

```html
<div class="flex w-fit items-stretch gap-2">
  <!-- File operations -->
  <div role="group" class="button-group" aria-label="File operations">
    <button type="button" class="btn-sm-icon-outline" aria-label="New file">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
        <polyline points="14,2 14,8 20,8" />
      </svg>
    </button>
    <button type="button" class="btn-sm-icon-outline" aria-label="Save">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" />
        <polyline points="17,21 17,13 7,13 7,21" />
        <polyline points="7,3 7,8 15,8" />
      </svg>
    </button>
  </div>
  
  <!-- Edit operations -->
  <div role="group" class="button-group" aria-label="Edit operations">
    <button type="button" class="btn-sm-icon-outline" aria-label="Undo">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M3 7v6h6" />
        <path d="M21 17a9 9 0 0 0-9-9 9 9 0 0 0-6 2.3L3 13" />
      </svg>
    </button>
    <button type="button" class="btn-sm-icon-outline" aria-label="Redo">
      <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M21 7v6h-6" />
        <path d="M3 17a9 9 0 0 1 9-9 9 9 0 0 1 6 2.3l3 2.7" />
      </svg>
    </button>
  </div>
</div>
```

### Status Filter Groups

```html
<div class="space-y-2">
  <label class="text-sm font-medium">Filter by status</label>
  <div role="group" class="button-group" aria-label="Status filter">
    <button type="button" class="btn" aria-pressed="true">All</button>
    <button type="button" class="btn-outline" aria-pressed="false">Active</button>
    <button type="button" class="btn-outline" aria-pressed="false">Pending</button>
    <button type="button" class="btn-outline" aria-pressed="false">Inactive</button>
  </div>
</div>
```

### Modal Actions

```html
<div class="flex justify-end gap-2 pt-6 border-t">
  <button type="button" class="btn-outline">Cancel</button>
  
  <div role="group" class="button-group">
    <button type="button" class="btn-outline">Save Draft</button>
    <button type="button" class="btn">Publish Now</button>
  </div>
</div>
```

## Related Components

- [Button](./button.md) - Individual button styling
- [Dropdown Menu](./dropdown-menu.md) - For button group dropdowns
- [Toggle](./toggle.md) - For toggle-style button groups
- [Toolbar](./toolbar.md) - For complex tool arrangements
---
## button

# Button Component

Displays a button or a component that looks like a button.

## Basic Usage

```html
<button class="btn">Button</button>
```

## CSS Classes

### Button Variants
- **`btn`** or **`btn-primary`** - Primary buttons (default)
- **`btn-secondary`** - Secondary buttons  
- **`btn-destructive`** - Destructive/danger buttons
- **`btn-outline`** - Outlined buttons
- **`btn-ghost`** - Ghost/transparent buttons
- **`btn-link`** - Link-style buttons
- **`btn-icon`** - Icon-only buttons

### Button Sizes
- **Default** - Standard button size
- **`btn-sm`** - Small buttons
- **`btn-lg`** - Large buttons

### Size + Variant Combinations
You can combine sizes with any variant:
- `btn-lg-destructive` - Large destructive button
- `btn-sm-outline` - Small outline button
- `btn-sm-icon-outline` - Small outline icon button
- `btn-icon-destructive` - Destructive icon button

## Component Attributes

### Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Button styling classes | Yes |
| `type` | string | "button", "submit", "reset" | Recommended |
| `disabled` | boolean | Disables the button | Optional |
| `aria-label` | string | Accessible label for icon buttons | Icon buttons |

### No JavaScript Required (Basic)
Basic buttons work with pure CSS and HTML.

## HTML Structure

```html
<!-- Basic button -->
<button class="btn" type="button">Label</button>

<!-- Button with icon -->
<button class="btn">
  <svg><!-- icon --></svg>
  Label
</button>

<!-- Icon-only button -->
<button class="btn-icon" aria-label="Description">
  <svg><!-- icon --></svg>
</button>
```

## Examples

### Button Variants

```html
<!-- Primary (default) -->
<button class="btn">Primary</button>
<button class="btn-primary">Primary</button>

<!-- Secondary -->
<button class="btn-secondary">Secondary</button>

<!-- Destructive -->
<button class="btn-destructive">Destructive</button>

<!-- Outline -->
<button class="btn-outline">Outline</button>

<!-- Ghost -->
<button class="btn-ghost">Ghost</button>

<!-- Link -->
<button class="btn-link">Link</button>
```

### Button Sizes

```html
<!-- Small buttons -->
<button class="btn-sm">Small</button>
<button class="btn-sm-secondary">Small Secondary</button>
<button class="btn-sm-outline">Small Outline</button>

<!-- Default size -->
<button class="btn">Default</button>

<!-- Large buttons -->
<button class="btn-lg">Large</button>
<button class="btn-lg-destructive">Large Destructive</button>
<button class="btn-lg-outline">Large Outline</button>
```

### Icon Buttons

```html
<!-- Icon-only buttons -->
<button class="btn-icon" aria-label="Next">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="m9 18 6-6-6-6" />
  </svg>
</button>

<button class="btn-icon-outline" aria-label="Settings">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
    <circle cx="12" cy="12" r="3" />
  </svg>
</button>

<button class="btn-icon-destructive" aria-label="Delete">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3 6h18" />
    <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
    <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
  </svg>
</button>

<!-- Small icon buttons -->
<button class="btn-sm-icon" aria-label="Edit">
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
    <path d="M18.375 2.625a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.375-9.375Z" />
  </svg>
</button>
```

### Buttons with Icons and Text

```html
<!-- Icon + text buttons -->
<button class="btn">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M14.536 21.686a.5.5 0 0 0 .937-.024l6.5-19a.496.496 0 0 0-.635-.635l-19 6.5a.5.5 0 0 0-.024.937l7.93 3.18a2 2 0 0 1 1.112 1.11z" />
    <path d="m21.854 2.147-10.94 10.939" />
  </svg>
  Send email
</button>

<button class="btn-outline">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
    <polyline points="7,10 12,15 17,10" />
    <line x1="12" x2="12" y1="15" y2="3" />
  </svg>
  Download
</button>

<button class="btn-secondary">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M5 12l5 5l10 -10" />
  </svg>
  Save changes
</button>

<!-- Icon after text -->
<button class="btn">
  Continue
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="m9 18 6-6-6-6" />
  </svg>
</button>
```

### Loading State Buttons

```html
<!-- Button with loading spinner -->
<button class="btn" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Loading...
</button>

<!-- Loading icon only -->
<button class="btn-icon" disabled aria-label="Loading">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
</button>

<!-- Different loading states -->
<button class="btn-outline" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Please wait...
</button>

<button class="btn-sm-secondary" disabled>
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
    <path d="M21 12a9 9 0 1 1-6.219-8.56" />
  </svg>
  Processing
</button>
```

### Button States

```html
<!-- Normal state -->
<button class="btn">Normal</button>

<!-- Disabled state -->
<button class="btn" disabled>Disabled</button>

<!-- Active/pressed state -->
<button class="btn pressed" aria-pressed="true">Active</button>

<!-- Focus state (handled by CSS automatically) -->
<button class="btn">Focusable</button>

<!-- Different variants disabled -->
<button class="btn-outline" disabled>Disabled Outline</button>
<button class="btn-destructive" disabled>Disabled Destructive</button>
<button class="btn-ghost" disabled>Disabled Ghost</button>
```

### Form Integration

```html
<!-- Form buttons -->
<form>
  <div class="space-y-4">
    <!-- Form fields here -->
    
    <!-- Form actions -->
    <div class="flex justify-end gap-2">
      <button type="button" class="btn-outline">Cancel</button>
      <button type="submit" class="btn">Save</button>
    </div>
  </div>
</form>

<!-- Different submit scenarios -->
<form>
  <!-- Basic submit -->
  <button type="submit" class="btn">Submit</button>
  
  <!-- Submit with loading -->
  <button type="submit" class="btn" id="submit-btn">
    <span class="hidden loading-spinner">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="animate-spin">
        <path d="M21 12a9 9 0 1 1-6.219-8.56" />
      </svg>
    </span>
    <span class="button-text">Submit</span>
  </button>
  
  <!-- Reset button -->
  <button type="reset" class="btn-outline">Reset</button>
</form>
```

### Button as Links

```html
<!-- Styled as button, behaves as link -->
<a href="/dashboard" class="btn">Go to Dashboard</a>
<a href="/profile" class="btn-outline">View Profile</a>
<a href="/settings" class="btn-ghost">Settings</a>

<!-- External links -->
<a href="https://example.com" target="_blank" rel="noopener noreferrer" class="btn">
  Visit Site
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M7 7h10v10" />
    <path d="M7 17 17 7" />
  </svg>
</a>
```

### Button Groups

```html
<!-- Simple button group -->
<div class="flex gap-2">
  <button class="btn-outline">Previous</button>
  <button class="btn-outline">Next</button>
</div>

<!-- Segmented control style -->
<div class="inline-flex border border-border rounded-md">
  <button class="btn-ghost rounded-r-none border-r">Day</button>
  <button class="btn-ghost rounded-none border-r">Week</button>
  <button class="btn rounded-l-none">Month</button>
</div>

<!-- Action group -->
<div class="flex items-center gap-2">
  <button class="btn">Save</button>
  <button class="btn-outline">Save & Continue</button>
  <div class="border-l border-border pl-2 ml-2">
    <button class="btn-icon-ghost" aria-label="More options">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="12" cy="5" r="1" />
        <circle cx="12" cy="19" r="1" />
      </svg>
    </button>
  </div>
</div>
```

### Responsive Buttons

```html
<!-- Responsive button sizes -->
<button class="btn-sm md:btn">
  Responsive Size
</button>

<!-- Hide text on small screens -->
<button class="btn">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
    <polyline points="7,10 12,15 17,10" />
    <line x1="12" x2="12" y1="15" y2="3" />
  </svg>
  <span class="hidden sm:inline ml-2">Download</span>
</button>

<!-- Stack buttons on mobile -->
<div class="flex flex-col sm:flex-row gap-2">
  <button class="btn-outline">Cancel</button>
  <button class="btn">Confirm</button>
</div>
```

## Accessibility Features

- **Keyboard Navigation**: All buttons support Tab/Enter/Space
- **Screen Reader Support**: Use `aria-label` for icon-only buttons
- **Focus Indicators**: Automatic focus rings and states
- **Disabled State**: Proper `disabled` attribute handling

### Enhanced Accessibility

```html
<!-- Icon button with proper labeling -->
<button class="btn-icon-outline" aria-label="Close dialog">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M18 6 6 18" />
    <path d="m6 6 12 12" />
  </svg>
</button>

<!-- Toggle button with pressed state -->
<button class="btn-outline" aria-pressed="false" onclick="togglePressed(this)">
  Toggle Feature
</button>

<!-- Button with description -->
<button class="btn" aria-describedby="save-help">
  Save Draft
</button>
<p id="save-help" class="sr-only">
  Saves your work without publishing
</p>

<!-- Loading button with live region -->
<button class="btn" onclick="submitForm(this)" aria-describedby="status">
  Submit
</button>
<div id="status" aria-live="polite" class="sr-only"></div>
```

## JavaScript Integration

### Button State Management

```javascript
// Toggle loading state
function setButtonLoading(button, isLoading) {
  const spinner = button.querySelector('.loading-spinner');
  const text = button.querySelector('.button-text');
  
  if (isLoading) {
    button.disabled = true;
    spinner?.classList.remove('hidden');
    text && (text.textContent = 'Loading...');
  } else {
    button.disabled = false;
    spinner?.classList.add('hidden');
    text && (text.textContent = 'Submit');
  }
}

// Toggle pressed state
function togglePressed(button) {
  const isPressed = button.getAttribute('aria-pressed') === 'true';
  button.setAttribute('aria-pressed', (!isPressed).toString());
  button.classList.toggle('pressed', !isPressed);
}

// Form submission with loading
function submitForm(button) {
  setButtonLoading(button, true);
  
  // Simulate async operation
  setTimeout(() => {
    setButtonLoading(button, false);
  }, 2000);
}
```

### React Integration

```jsx
import React, { useState } from 'react';

function Button({ 
  variant = 'btn',
  size = '',
  children,
  loading = false,
  disabled = false,
  className = '',
  ...props 
}) {
  const buttonClasses = [
    size ? `${size}-${variant}` : variant,
    className
  ].filter(Boolean).join(' ');
  
  return (
    <button 
      className={buttonClasses}
      disabled={disabled || loading}
      {...props}
    >
      {loading && (
        <svg className="animate-spin mr-2" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
          <path d="M21 12a9 9 0 1 1-6.219-8.56" />
        </svg>
      )}
      {children}
    </button>
  );
}

// Usage
function App() {
  const [loading, setLoading] = useState(false);
  
  const handleSubmit = async () => {
    setLoading(true);
    await new Promise(resolve => setTimeout(resolve, 2000));
    setLoading(false);
  };
  
  return (
    <div className="space-x-2">
      <Button variant="btn">Primary</Button>
      <Button variant="btn-outline">Outline</Button>
      <Button variant="btn" loading={loading} onClick={handleSubmit}>
        Submit
      </Button>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <button 
    :class="buttonClasses"
    :disabled="disabled || loading"
    v-bind="$attrs"
  >
    <svg 
      v-if="loading"
      class="animate-spin mr-2" 
      width="20" height="20" 
      viewBox="0 0 24 24" 
      fill="none" 
      stroke="currentColor" 
      stroke-width="2"
    >
      <path d="M21 12a9 9 0 1 1-6.219-8.56" />
    </svg>
    <slot />
  </button>
</template>

<script>
export default {
  props: {
    variant: {
      type: String,
      default: 'btn'
    },
    size: String,
    loading: Boolean,
    disabled: Boolean
  },
  computed: {
    buttonClasses() {
      const classes = [];
      if (this.size) {
        classes.push(`${this.size}-${this.variant}`);
      } else {
        classes.push(this.variant);
      }
      return classes.join(' ');
    }
  }
};
</script>
```

## Best Practices

1. **Use Semantic HTML**: Use `<button>` for actions, `<a>` for navigation
2. **Button Types**: Specify `type` attribute for form buttons
3. **Accessibility**: Provide `aria-label` for icon-only buttons
4. **Loading States**: Disable buttons during async operations
5. **Visual Hierarchy**: Use appropriate variants for action priority
6. **Consistent Sizing**: Maintain consistent button sizes within groups
7. **Icon Guidelines**: Use 24px icons for default buttons, 16px for small
8. **Focus Management**: Ensure proper focus indicators and flow

## Common Patterns

### Confirmation Dialogs

```html
<div class="space-y-4">
  <p>Are you sure you want to delete this item?</p>
  <div class="flex justify-end gap-2">
    <button type="button" class="btn-outline" onclick="closeDialog()">
      Cancel
    </button>
    <button type="button" class="btn-destructive" onclick="confirmDelete()">
      Delete
    </button>
  </div>
</div>
```

### Toolbar Actions

```html
<div class="flex items-center gap-1 p-2 border-b">
  <button class="btn-sm-icon" aria-label="Bold">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
      <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z" />
    </svg>
  </button>
  <button class="btn-sm-icon" aria-label="Italic">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <line x1="19" x2="10" y1="4" y2="4" />
      <line x1="14" x2="5" y1="20" y2="20" />
      <line x1="15" x2="9" y1="4" y2="20" />
    </svg>
  </button>
  <div class="border-l border-border mx-1 h-6"></div>
  <button class="btn-sm-outline">Save</button>
</div>
```

### Call-to-Action Sections

```html
<div class="text-center space-y-4 py-12">
  <h2 class="text-2xl font-bold">Ready to get started?</h2>
  <p class="text-muted-foreground">Join thousands of satisfied customers today.</p>
  <div class="flex justify-center gap-3">
    <button class="btn-lg">Start Free Trial</button>
    <button class="btn-lg-outline">Learn More</button>
  </div>
</div>
```

## Related Components

- [Button Group](./button-group.md) - For grouping related buttons
- [Spinner](./spinner.md) - For button loading states
- [Input](./input.md) - For form integration
- [Dialog](./dialog.md) - For confirmation buttons
---
## combobox

# Combobox Component

Autocomplete input and command palette with a list of suggestions.

## Basic Usage

```html
<div id="combobox-demo" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[200px]" id="combobox-demo-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="combobox-demo-listbox">
    <span class="truncate"></span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="combobox-demo-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Search..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="combobox-demo-listbox" aria-labelledby="combobox-demo-trigger" />
    </header>
    <div role="listbox" id="combobox-demo-listbox" aria-orientation="vertical" aria-labelledby="combobox-demo-trigger" data-empty="No results found.">
      <div role="option" data-value="option1">Option 1</div>
      <div role="option" data-value="option2">Option 2</div>
      <div role="option" data-value="option3">Option 3</div>
    </div>
  </div>
  <input type="hidden" name="combobox-value" value="" />
</div>
```

## Important Note

**Combobox uses the same markup and JavaScript as the [Select component](./select.md)**, with the key difference being the search input at the top of the listbox.

## CSS Classes

### Primary Classes
- **`select`** - Core combobox container class

### Button Classes
- **`btn-outline`** - Trigger button styling
- **`justify-between`** - Space between text and icon
- **`font-normal`** - Normal font weight
- **`w-[200px]`** - Fixed width (customizable)

### Supporting Classes
- **`truncate`** - Text truncation
- **`text-muted-foreground`** - Muted text color
- **`opacity-50`** - Reduced opacity
- **`shrink-0`** - Prevent shrinking

## JavaScript Required

This component requires the Basecoat JavaScript files:
- `basecoat.js` - Core functionality
- `select.js` - Combobox/Select logic

```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/select.min.js" defer></script>
```

## Component Attributes

### Container Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for combobox | Yes |
| `class` | string | Must include "select" | Yes |

### Trigger Button Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Button identifier | Yes |
| `type` | string | Should be "button" | Yes |
| `class` | string | Button styling classes | Yes |
| `aria-haspopup` | string | Should be "listbox" | Yes |
| `aria-expanded` | boolean | Expansion state | Yes |
| `aria-controls` | string | References listbox ID | Yes |

### Search Input Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Should be "text" | Yes |
| `placeholder` | string | Search placeholder text | Recommended |
| `autocomplete` | string | Should be "off" | Yes |
| `autocorrect` | string | Should be "off" | Yes |
| `spellcheck` | boolean | Should be false | Yes |
| `role` | string | Should be "combobox" | Yes |
| `aria-autocomplete` | string | Should be "list" | Yes |
| `aria-expanded` | boolean | Expansion state | Yes |
| `aria-controls` | string | References listbox ID | Yes |
| `aria-labelledby` | string | References trigger ID | Yes |

### Listbox Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `role` | string | Should be "listbox" | Yes |
| `id` | string | Listbox identifier | Yes |
| `aria-orientation` | string | Should be "vertical" | Yes |
| `aria-labelledby` | string | References trigger ID | Yes |
| `data-empty` | string | Empty state message | Recommended |

## HTML Structure

```html
<div id="[unique-id]" class="select">
  <!-- Trigger button -->
  <button type="button" class="btn-outline justify-between font-normal w-[width]" 
          id="[unique-id]-trigger" 
          aria-haspopup="listbox" 
          aria-expanded="false" 
          aria-controls="[unique-id]-listbox">
    <span class="truncate">[Selected value]</span>
    <svg><!-- chevron icon --></svg>
  </button>
  
  <!-- Popover container -->
  <div id="[unique-id]-popover" data-popover aria-hidden="true">
    <!-- Search header -->
    <header>
      <svg><!-- search icon --></svg>
      <input type="text" 
             placeholder="[Search placeholder]" 
             autocomplete="off" 
             autocorrect="off" 
             spellcheck="false"
             aria-autocomplete="list" 
             role="combobox" 
             aria-expanded="false" 
             aria-controls="[unique-id]-listbox" 
             aria-labelledby="[unique-id]-trigger" />
    </header>
    
    <!-- Options listbox -->
    <div role="listbox" 
         id="[unique-id]-listbox" 
         aria-orientation="vertical" 
         aria-labelledby="[unique-id]-trigger" 
         data-empty="[Empty message]">
      <div role="option" data-value="[value]">[Option text]</div>
      <!-- More options -->
    </div>
  </div>
  
  <!-- Hidden form input -->
  <input type="hidden" name="[name]" value="" />
</div>
```

## Examples

### Framework Selector

```html
<div id="framework-select" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[200px]" id="framework-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="framework-select-listbox">
    <span class="truncate"></span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="framework-select-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Search framework..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="framework-select-listbox" aria-labelledby="framework-select-trigger" />
    </header>
    <div role="listbox" id="framework-select-listbox" aria-orientation="vertical" aria-labelledby="framework-select-trigger" data-empty="No framework found.">
      <div role="option" data-value="Next.js">Next.js</div>
      <div role="option" data-value="SvelteKit">SvelteKit</div>
      <div role="option" data-value="Nuxt.js">Nuxt.js</div>
      <div role="option" data-value="Remix">Remix</div>
      <div role="option" data-value="Astro">Astro</div>
    </div>
  </div>
  <input type="hidden" name="framework" value="" />
</div>
```

### Language Selector

```html
<div id="language-select" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[250px]" id="language-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="language-select-listbox">
    <span class="truncate">Select language...</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="language-select-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Search languages..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="language-select-listbox" aria-labelledby="language-select-trigger" />
    </header>
    <div role="listbox" id="language-select-listbox" aria-orientation="vertical" aria-labelledby="language-select-trigger" data-empty="No language found.">
      <div role="option" data-value="en">English</div>
      <div role="option" data-value="es">Español</div>
      <div role="option" data-value="fr">Français</div>
      <div role="option" data-value="de">Deutsch</div>
      <div role="option" data-value="it">Italiano</div>
      <div role="option" data-value="pt">Português</div>
      <div role="option" data-value="ru">Русский</div>
      <div role="option" data-value="ja">日本語</div>
      <div role="option" data-value="ko">한국어</div>
      <div role="option" data-value="zh">中文</div>
    </div>
  </div>
  <input type="hidden" name="language" value="" />
</div>
```

### User Picker

```html
<div id="user-select" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[300px]" id="user-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="user-select-listbox">
    <span class="truncate">Assign to...</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="user-select-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Search users..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="user-select-listbox" aria-labelledby="user-select-trigger" />
    </header>
    <div role="listbox" id="user-select-listbox" aria-orientation="vertical" aria-labelledby="user-select-trigger" data-empty="No users found.">
      <div role="option" data-value="john" class="flex items-center gap-2">
        <img class="size-6 rounded-full" src="https://github.com/johnsmith.png" alt="John Smith" />
        <div>
          <div class="font-medium">John Smith</div>
          <div class="text-xs text-muted-foreground">john@company.com</div>
        </div>
      </div>
      <div role="option" data-value="jane" class="flex items-center gap-2">
        <img class="size-6 rounded-full" src="https://github.com/janedoe.png" alt="Jane Doe" />
        <div>
          <div class="font-medium">Jane Doe</div>
          <div class="text-xs text-muted-foreground">jane@company.com</div>
        </div>
      </div>
      <div role="option" data-value="alex" class="flex items-center gap-2">
        <img class="size-6 rounded-full" src="https://github.com/alexjohnson.png" alt="Alex Johnson" />
        <div>
          <div class="font-medium">Alex Johnson</div>
          <div class="text-xs text-muted-foreground">alex@company.com</div>
        </div>
      </div>
    </div>
  </div>
  <input type="hidden" name="assignee" value="" />
</div>
```

### Tag Selector

```html
<div id="tag-select" class="select">
  <button type="button" class="btn-outline justify-between font-normal w-[200px]" id="tag-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="tag-select-listbox">
    <span class="truncate">Add tags...</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="tag-select-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z" />
      </svg>
      <input type="text" placeholder="Search tags..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="tag-select-listbox" aria-labelledby="tag-select-trigger" />
    </header>
    <div role="listbox" id="tag-select-listbox" aria-orientation="vertical" aria-labelledby="tag-select-trigger" data-empty="No tags found.">
      <div role="option" data-value="bug" class="flex items-center gap-2">
        <span class="w-3 h-3 bg-red-500 rounded-full"></span>
        bug
      </div>
      <div role="option" data-value="feature" class="flex items-center gap-2">
        <span class="w-3 h-3 bg-green-500 rounded-full"></span>
        feature
      </div>
      <div role="option" data-value="enhancement" class="flex items-center gap-2">
        <span class="w-3 h-3 bg-blue-500 rounded-full"></span>
        enhancement
      </div>
      <div role="option" data-value="documentation" class="flex items-center gap-2">
        <span class="w-3 h-3 bg-yellow-500 rounded-full"></span>
        documentation
      </div>
    </div>
  </div>
  <input type="hidden" name="tags" value="" />
</div>
```

### Small Combobox

```html
<div id="small-select" class="select">
  <button type="button" class="btn-sm-outline justify-between font-normal w-[150px]" id="small-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="small-select-listbox">
    <span class="truncate text-sm">Select...</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="small-select-popover" data-popover aria-hidden="true">
    <header class="p-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Search..." class="text-sm" autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="small-select-listbox" aria-labelledby="small-select-trigger" />
    </header>
    <div role="listbox" id="small-select-listbox" aria-orientation="vertical" aria-labelledby="small-select-trigger" data-empty="Nothing found.">
      <div role="option" data-value="xs" class="text-sm">XS</div>
      <div role="option" data-value="sm" class="text-sm">Small</div>
      <div role="option" data-value="md" class="text-sm">Medium</div>
      <div role="option" data-value="lg" class="text-sm">Large</div>
      <div role="option" data-value="xl" class="text-sm">XL</div>
    </div>
  </div>
  <input type="hidden" name="size" value="" />
</div>
```

### Command Palette Style

```html
<div id="command-select" class="select">
  <button type="button" class="btn-outline justify-start font-normal w-full max-w-md text-muted-foreground" id="command-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="command-select-listbox">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="mr-2">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
    <span class="truncate">Search commands...</span>
    <div class="ml-auto flex items-center gap-1">
      <kbd class="inline-flex items-center rounded border bg-muted px-1.5 py-0.5 text-xs font-mono text-muted-foreground">⌘K</kbd>
    </div>
  </button>
  <div id="command-select-popover" data-popover aria-hidden="true">
    <header>
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" placeholder="Type a command or search..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="command-select-listbox" aria-labelledby="command-select-trigger" />
    </header>
    <div role="listbox" id="command-select-listbox" aria-orientation="vertical" aria-labelledby="command-select-trigger" data-empty="No commands found.">
      <div class="px-3 py-2 text-xs font-medium text-muted-foreground">Suggestions</div>
      <div role="option" data-value="new-file" class="flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
          <polyline points="14,2 14,8 20,8" />
        </svg>
        <div>Create New File</div>
        <div class="ml-auto text-xs text-muted-foreground">⌘N</div>
      </div>
      <div role="option" data-value="open-file" class="flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" />
          <polyline points="14,2 14,8 20,8" />
          <line x1="16" x2="8" y1="13" y2="13" />
          <line x1="16" x2="8" y1="17" y2="17" />
          <line x1="10" x2="8" y1="9" y2="9" />
        </svg>
        <div>Open File</div>
        <div class="ml-auto text-xs text-muted-foreground">⌘O</div>
      </div>
      <div role="option" data-value="save" class="flex items-center gap-2">
        <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" />
          <polyline points="17,21 17,13 7,13 7,21" />
          <polyline points="7,3 7,8 15,8" />
        </svg>
        <div>Save</div>
        <div class="ml-auto text-xs text-muted-foreground">⌘S</div>
      </div>
    </div>
  </div>
  <input type="hidden" name="command" value="" />
</div>
```

### Form Integration

```html
<form class="space-y-4">
  <div class="space-y-2">
    <label for="project-select-trigger" class="label">Project</label>
    <div id="project-select" class="select">
      <button type="button" class="btn-outline justify-between font-normal w-full" id="project-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="project-select-listbox">
        <span class="truncate">Choose project...</span>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
          <path d="m7 15 5 5 5-5" />
          <path d="m7 9 5-5 5 5" />
        </svg>
      </button>
      <div id="project-select-popover" data-popover aria-hidden="true">
        <header>
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="11" cy="11" r="8" />
            <path d="m21 21-4.3-4.3" />
          </svg>
          <input type="text" placeholder="Search projects..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="project-select-listbox" aria-labelledby="project-select-trigger" />
        </header>
        <div role="listbox" id="project-select-listbox" aria-orientation="vertical" aria-labelledby="project-select-trigger" data-empty="No projects found.">
          <div role="option" data-value="website">Website Redesign</div>
          <div role="option" data-value="mobile-app">Mobile App</div>
          <div role="option" data-value="api">API Development</div>
        </div>
      </div>
      <input type="hidden" name="project" value="" required />
    </div>
  </div>
  
  <button type="submit" class="btn">Create Task</button>
</form>
```

## Accessibility Features

- **ARIA Combobox**: Proper `role="combobox"` implementation
- **Keyboard Navigation**: Arrow keys, Enter, Escape support
- **Screen Reader Support**: Comprehensive ARIA attributes
- **Live Updates**: Search results announced to screen readers
- **Focus Management**: Proper focus handling throughout interaction

### Enhanced Accessibility

```html
<div id="accessible-select" class="select">
  <label for="accessible-select-trigger" class="label sr-only">Choose option</label>
  <button type="button" class="btn-outline justify-between font-normal w-[200px]" id="accessible-select-trigger" aria-haspopup="listbox" aria-expanded="false" aria-controls="accessible-select-listbox" aria-describedby="accessible-select-help">
    <span class="truncate">Select option...</span>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="accessible-select-help" class="sr-only">Use arrow keys to navigate options</div>
  <div id="accessible-select-popover" data-popover aria-hidden="true">
    <header>
      <label for="accessible-select-search" class="sr-only">Search options</label>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8" />
        <path d="m21 21-4.3-4.3" />
      </svg>
      <input type="text" id="accessible-select-search" placeholder="Search..." autocomplete="off" autocorrect="off" spellcheck="false" aria-autocomplete="list" role="combobox" aria-expanded="false" aria-controls="accessible-select-listbox" aria-labelledby="accessible-select-trigger" />
    </header>
    <div role="listbox" id="accessible-select-listbox" aria-orientation="vertical" aria-labelledby="accessible-select-trigger" data-empty="No options found." aria-live="polite">
      <div role="option" data-value="option1" aria-selected="false">Option 1</div>
      <div role="option" data-value="option2" aria-selected="false">Option 2</div>
      <div role="option" data-value="option3" aria-selected="false">Option 3</div>
    </div>
  </div>
  <input type="hidden" name="value" value="" />
</div>
```

## JavaScript Integration

### Basic Implementation

```javascript
// Initialize combobox
document.addEventListener('DOMContentLoaded', function() {
  // The select.js handles all combobox functionality automatically
  // No additional initialization needed if using basecoat scripts
});

// Access selected value
function getComboboxValue(comboboxId) {
  const hiddenInput = document.querySelector(`#${comboboxId} input[type="hidden"]`);
  return hiddenInput ? hiddenInput.value : null;
}

// Set combobox value programmatically
function setComboboxValue(comboboxId, value) {
  const combobox = document.getElementById(comboboxId);
  const hiddenInput = combobox.querySelector('input[type="hidden"]');
  const trigger = combobox.querySelector('button');
  const option = combobox.querySelector(`[data-value="${value}"]`);
  
  if (hiddenInput && option) {
    hiddenInput.value = value;
    trigger.querySelector('span').textContent = option.textContent;
  }
}
```

### Custom Search Function

```javascript
// Override default search behavior
function initCustomCombobox(comboboxId, searchFunction) {
  const combobox = document.getElementById(comboboxId);
  const searchInput = combobox.querySelector('input[type="text"]');
  const listbox = combobox.querySelector('[role="listbox"]');
  
  searchInput.addEventListener('input', function(e) {
    const query = e.target.value.toLowerCase();
    const options = listbox.querySelectorAll('[role="option"]');
    
    options.forEach(option => {
      const matches = searchFunction(option.textContent, query);
      option.style.display = matches ? 'block' : 'none';
    });
    
    // Update empty state
    const visibleOptions = listbox.querySelectorAll('[role="option"]:not([style*="display: none"])');
    const emptyMessage = listbox.getAttribute('data-empty');
    
    if (visibleOptions.length === 0 && emptyMessage) {
      listbox.innerHTML = `<div class="px-3 py-2 text-sm text-muted-foreground">${emptyMessage}</div>`;
    }
  });
}

// Fuzzy search implementation
function fuzzySearch(text, query) {
  const textLower = text.toLowerCase();
  const queryLower = query.toLowerCase();
  
  let queryIndex = 0;
  for (let textIndex = 0; textIndex < textLower.length && queryIndex < queryLower.length; textIndex++) {
    if (textLower[textIndex] === queryLower[queryIndex]) {
      queryIndex++;
    }
  }
  
  return queryIndex === queryLower.length;
}

// Usage
initCustomCombobox('my-combobox', fuzzySearch);
```

### React Integration

```jsx
import React, { useState, useRef, useEffect } from 'react';

function Combobox({ options, placeholder, onSelect, value, className = '' }) {
  const [isOpen, setIsOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [selectedIndex, setSelectedIndex] = useState(-1);
  const searchInputRef = useRef(null);
  
  const filteredOptions = options.filter(option =>
    option.label.toLowerCase().includes(searchQuery.toLowerCase())
  );
  
  const selectedOption = options.find(opt => opt.value === value);
  
  const handleKeyDown = (e) => {
    switch (e.key) {
      case 'ArrowDown':
        e.preventDefault();
        setSelectedIndex(prev => 
          prev < filteredOptions.length - 1 ? prev + 1 : prev
        );
        break;
      case 'ArrowUp':
        e.preventDefault();
        setSelectedIndex(prev => prev > 0 ? prev - 1 : -1);
        break;
      case 'Enter':
        e.preventDefault();
        if (selectedIndex >= 0) {
          onSelect(filteredOptions[selectedIndex]);
          setIsOpen(false);
        }
        break;
      case 'Escape':
        setIsOpen(false);
        break;
    }
  };
  
  useEffect(() => {
    if (isOpen && searchInputRef.current) {
      searchInputRef.current.focus();
    }
  }, [isOpen]);
  
  return (
    <div className={`select ${className}`}>
      <button
        type="button"
        className="btn-outline justify-between font-normal w-full"
        onClick={() => setIsOpen(!isOpen)}
        aria-haspopup="listbox"
        aria-expanded={isOpen}
      >
        <span className="truncate">
          {selectedOption ? selectedOption.label : placeholder}
        </span>
        <svg className="text-muted-foreground opacity-50 shrink-0" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2">
          <path d="m7 15 5 5 5-5" />
          <path d="m7 9 5-5 5 5" />
        </svg>
      </button>
      
      {isOpen && (
        <div className="absolute z-10 mt-1 w-full bg-background border rounded-md shadow-lg">
          <div className="flex items-center gap-2 p-3 border-b">
            <svg className="w-4 h-4 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <circle cx="11" cy="11" r="8" />
              <path d="m21 21-4.3-4.3" />
            </svg>
            <input
              ref={searchInputRef}
              type="text"
              className="flex-1 bg-transparent outline-none"
              placeholder="Search..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              onKeyDown={handleKeyDown}
            />
          </div>
          
          <div className="max-h-60 overflow-y-auto">
            {filteredOptions.length === 0 ? (
              <div className="px-3 py-2 text-sm text-muted-foreground">
                No options found
              </div>
            ) : (
              filteredOptions.map((option, index) => (
                <div
                  key={option.value}
                  className={`px-3 py-2 cursor-pointer hover:bg-muted ${
                    index === selectedIndex ? 'bg-muted' : ''
                  }`}
                  onClick={() => {
                    onSelect(option);
                    setIsOpen(false);
                  }}
                >
                  {option.label}
                </div>
              ))
            )}
          </div>
        </div>
      )}
    </div>
  );
}

// Usage
function App() {
  const [selectedFramework, setSelectedFramework] = useState('');
  
  const frameworks = [
    { value: 'next', label: 'Next.js' },
    { value: 'svelte', label: 'SvelteKit' },
    { value: 'nuxt', label: 'Nuxt.js' },
    { value: 'remix', label: 'Remix' },
    { value: 'astro', label: 'Astro' }
  ];
  
  return (
    <Combobox
      options={frameworks}
      placeholder="Select framework..."
      value={selectedFramework}
      onSelect={(option) => setSelectedFramework(option.value)}
    />
  );
}
```

### Vue Integration

```vue
<template>
  <div class="select" :class="className">
    <button
      type="button"
      class="btn-outline justify-between font-normal w-full"
      @click="toggleOpen"
      :aria-expanded="isOpen"
      aria-haspopup="listbox"
    >
      <span class="truncate">
        {{ selectedOption?.label || placeholder }}
      </span>
      <svg class="text-muted-foreground opacity-50 shrink-0" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
        <path d="m7 15 5 5 5-5" />
        <path d="m7 9 5-5 5 5" />
      </svg>
    </button>
    
    <div v-if="isOpen" class="absolute z-10 mt-1 w-full bg-background border rounded-md shadow-lg">
      <div class="flex items-center gap-2 p-3 border-b">
        <svg class="w-4 h-4 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <circle cx="11" cy="11" r="8" />
          <path d="m21 21-4.3-4.3" />
        </svg>
        <input
          ref="searchInput"
          type="text"
          class="flex-1 bg-transparent outline-none"
          placeholder="Search..."
          v-model="searchQuery"
          @keydown="handleKeyDown"
        />
      </div>
      
      <div class="max-h-60 overflow-y-auto">
        <div
          v-if="filteredOptions.length === 0"
          class="px-3 py-2 text-sm text-muted-foreground"
        >
          No options found
        </div>
        <div
          v-else
          v-for="(option, index) in filteredOptions"
          :key="option.value"
          class="px-3 py-2 cursor-pointer hover:bg-muted"
          :class="{ 'bg-muted': index === selectedIndex }"
          @click="selectOption(option)"
        >
          {{ option.label }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    options: Array,
    placeholder: String,
    modelValue: String,
    className: String
  },
  emits: ['update:modelValue'],
  data() {
    return {
      isOpen: false,
      searchQuery: '',
      selectedIndex: -1
    };
  },
  computed: {
    filteredOptions() {
      return this.options.filter(option =>
        option.label.toLowerCase().includes(this.searchQuery.toLowerCase())
      );
    },
    selectedOption() {
      return this.options.find(opt => opt.value === this.modelValue);
    }
  },
  methods: {
    toggleOpen() {
      this.isOpen = !this.isOpen;
      if (this.isOpen) {
        this.$nextTick(() => {
          this.$refs.searchInput?.focus();
        });
      }
    },
    selectOption(option) {
      this.$emit('update:modelValue', option.value);
      this.isOpen = false;
      this.searchQuery = '';
    },
    handleKeyDown(e) {
      switch (e.key) {
        case 'ArrowDown':
          e.preventDefault();
          this.selectedIndex = this.selectedIndex < this.filteredOptions.length - 1 
            ? this.selectedIndex + 1 
            : this.selectedIndex;
          break;
        case 'ArrowUp':
          e.preventDefault();
          this.selectedIndex = this.selectedIndex > 0 ? this.selectedIndex - 1 : -1;
          break;
        case 'Enter':
          e.preventDefault();
          if (this.selectedIndex >= 0) {
            this.selectOption(this.filteredOptions[this.selectedIndex]);
          }
          break;
        case 'Escape':
          this.isOpen = false;
          break;
      }
    }
  }
};
</script>
```

## Best Practices

1. **Searchable Content**: Only use for lists with many options (>10)
2. **Clear Placeholders**: Use descriptive search placeholder text
3. **Empty States**: Provide helpful empty state messages
4. **Performance**: Debounce search for large datasets
5. **Keyboard Support**: Ensure full keyboard accessibility
6. **Mobile UX**: Consider touch-friendly interactions
7. **Loading States**: Show loading for async data
8. **Error Handling**: Handle search failures gracefully

## Common Patterns

### Async Data Loading

```javascript
async function searchUsers(query) {
  const response = await fetch(`/api/users?search=${encodeURIComponent(query)}`);
  const users = await response.json();
  return users;
}

function initAsyncCombobox(comboboxId) {
  const searchInput = document.querySelector(`#${comboboxId} input[type="text"]`);
  const listbox = document.querySelector(`#${comboboxId} [role="listbox"]`);
  
  let debounceTimer;
  
  searchInput.addEventListener('input', function(e) {
    clearTimeout(debounceTimer);
    
    debounceTimer = setTimeout(async () => {
      const query = e.target.value;
      
      if (query.length < 2) {
        listbox.innerHTML = '';
        return;
      }
      
      // Show loading
      listbox.innerHTML = '<div class="px-3 py-2 text-sm">Searching...</div>';
      
      try {
        const results = await searchUsers(query);
        
        if (results.length === 0) {
          listbox.innerHTML = '<div class="px-3 py-2 text-sm text-muted-foreground">No users found</div>';
        } else {
          listbox.innerHTML = results.map(user => 
            `<div role="option" data-value="${user.id}">${user.name}</div>`
          ).join('');
        }
      } catch (error) {
        listbox.innerHTML = '<div class="px-3 py-2 text-sm text-destructive">Error loading results</div>';
      }
    }, 300);
  });
}
```

### Multi-select Combobox

```html
<!-- Multi-select combobox with tags -->
<div id="multi-select" class="select">
  <button type="button" class="btn-outline justify-between font-normal min-h-[40px] w-full" id="multi-select-trigger">
    <div class="flex flex-wrap gap-1 flex-1">
      <!-- Selected tags will be inserted here -->
    </div>
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground opacity-50 shrink-0 ml-2">
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  </button>
  <div id="multi-select-popover" data-popover aria-hidden="true">
    <header>
      <input type="text" placeholder="Search options..." />
    </header>
    <div role="listbox">
      <!-- Options with checkboxes -->
    </div>
  </div>
  <input type="hidden" name="selected-values" value="" />
</div>
```

## Related Components

- [Select](./select.md) - Basic dropdown selection
- [Input](./input.md) - Text input field
- [Command](./command.md) - Command palette interface
- [Popover](./popover.md) - Floating container
---
## tooltip

# Tooltip Component

A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it.

## Basic Usage

```html
<button class="btn-outline" data-tooltip="Add to library">Hover</button>
```

## CSS Classes

### No Custom Classes Required
Tooltips are controlled entirely through data attributes. The Basecoat JavaScript automatically handles styling.

### Supporting Classes
- **Any element** can have a tooltip by adding `data-tooltip` attribute
- **Layout classes** from other components work normally

### Tailwind Utilities Used
Tooltips use internal styling from the Basecoat JavaScript - no custom Tailwind classes needed.

## Component Attributes

### Required Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `data-tooltip` | string | Text content of the tooltip | Yes |

### Optional Attributes
| Attribute | Type | Values | Description | Required |
|-----------|------|--------|-------------|----------|
| `data-side` | string | "top", "bottom", "left", "right" | Position of tooltip relative to element | No (defaults to top) |
| `data-align` | string | "start", "center", "end" | Alignment of tooltip on the chosen side | No (defaults to center) |

## JavaScript Required

This component requires the Basecoat JavaScript files for tooltip functionality:

```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@0.3.6/dist/js/basecoat.min.js" defer></script>
```

## HTML Structure

```html
<!-- Any element can have a tooltip -->
<element data-tooltip="Tooltip text" data-side="position" data-align="alignment">
  Element content
</element>
```

## Examples

### Default Tooltip (Top Center)

```html
<button class="btn-outline" data-tooltip="Default tooltip">Default</button>
```

### Position Variations

```html
<!-- Top (default) -->
<button class="btn-outline" data-tooltip="Top tooltip" data-side="top">Top</button>

<!-- Bottom -->
<button class="btn-outline" data-tooltip="Bottom tooltip" data-side="bottom">Bottom</button>

<!-- Left -->
<button class="btn-outline" data-tooltip="Left tooltip" data-side="left">Left</button>

<!-- Right -->
<button class="btn-outline" data-tooltip="Right side tooltip" data-side="right">Right</button>
```

### Alignment Variations

```html
<!-- Center aligned (default) -->
<button class="btn-outline" data-tooltip="Center aligned tooltip">Center</button>

<!-- Start aligned -->
<button class="btn-outline" data-tooltip="Start aligned tooltip" data-align="start">Start</button>

<!-- End aligned -->
<button class="btn-outline" data-tooltip="End aligned tooltip" data-align="end">End</button>
```

### Combined Position and Alignment

```html
<!-- Bottom + Start -->
<button class="btn-outline" data-tooltip="Bottom side and start aligned tooltip" data-side="bottom" data-align="start">Bottom + Start</button>

<!-- Right + End -->
<button class="btn-outline" data-tooltip="Right side and end aligned tooltip" data-side="right" data-align="end">Right + End</button>

<!-- Left + Center -->
<button class="btn-outline" data-tooltip="Left side and center aligned tooltip" data-side="left" data-align="center">Left + Center</button>

<!-- Top + End -->
<button class="btn-outline" data-tooltip="Top side and end aligned tooltip" data-side="top" data-align="end">Top + End</button>
```

### Icon Button Tooltips

```html
<!-- Icon buttons with descriptive tooltips -->
<button class="btn-icon-outline" data-tooltip="Save document" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z" />
    <polyline points="17,21 17,13 7,13 7,21" />
    <polyline points="7,3 7,8 15,8" />
  </svg>
</button>

<button class="btn-icon-outline" data-tooltip="Edit item" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M12 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7" />
    <path d="M18.375 2.625a2.121 2.121 0 1 1 3 3L12 15l-4 1 1-4 9.375-9.375Z" />
  </svg>
</button>

<button class="btn-icon-destructive" data-tooltip="Delete forever" data-side="bottom">
  <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M3 6h18" />
    <path d="M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6" />
    <path d="M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2" />
  </svg>
</button>
```

### Link Tooltips

```html
<!-- External link with tooltip -->
<a href="https://example.com" target="_blank" class="btn-link" data-tooltip="Opens in new window">
  Visit External Site
</a>

<!-- Navigation link -->
<a href="/dashboard" class="btn-ghost" data-tooltip="View your dashboard">Dashboard</a>

<!-- Download link -->
<a href="/files/document.pdf" download class="btn-outline" data-tooltip="Download PDF (2.3MB)">
  Download Report
</a>
```

### Form Element Tooltips

```html
<!-- Input field with help tooltip -->
<div class="field">
  <label for="username">Username</label>
  <input 
    type="text" 
    id="username" 
    placeholder="Enter username"
    data-tooltip="Must be 3-20 characters, letters and numbers only"
    data-side="right"
  >
</div>

<!-- Checkbox with explanation -->
<div class="field">
  <label class="checkbox">
    <input 
      type="checkbox" 
      data-tooltip="We'll only send important updates about your account"
      data-side="right"
    >
    <span>Subscribe to notifications</span>
  </label>
</div>

<!-- Select with context -->
<div class="field">
  <label for="timezone">Timezone</label>
  <select 
    id="timezone" 
    class="select"
    data-tooltip="Used for scheduling and displaying dates"
    data-side="bottom"
  >
    <option>UTC</option>
    <option>EST</option>
    <option>PST</option>
  </select>
</div>
```

### Image Tooltips

```html
<!-- Avatar with user info -->
<img 
  src="/avatars/user.jpg" 
  alt="John Doe" 
  class="size-10 rounded-full"
  data-tooltip="John Doe - Product Manager"
  data-side="bottom"
>

<!-- Status indicator -->
<div class="flex items-center gap-2">
  <span class="size-3 rounded-full bg-green-500" data-tooltip="Online" data-side="top"></span>
  <span>User Status</span>
</div>

<!-- Chart data point -->
<div class="relative">
  <div 
    class="size-4 rounded-full bg-blue-500 cursor-pointer"
    data-tooltip="Sales: $12,450 (Mar 15)"
    data-side="top"
  ></div>
</div>
```

### Truncated Text Tooltips

```html
<!-- Long text with truncation -->
<div class="max-w-xs">
  <p 
    class="truncate text-sm"
    data-tooltip="This is a very long piece of text that would normally be truncated but the full content is shown in the tooltip"
    data-side="bottom"
  >
    This is a very long piece of text that would normally be truncated...
  </p>
</div>

<!-- Card title -->
<div class="card">
  <h3 
    class="text-lg font-semibold truncate"
    data-tooltip="Advanced Data Analytics and Reporting Dashboard Configuration Settings"
    data-side="bottom"
  >
    Advanced Data Analytics and...
  </h3>
</div>
```

### Interactive Element Tooltips

```html
<!-- Toggle switch -->
<label class="switch">
  <input type="checkbox" data-tooltip="Enable dark mode" data-side="right">
  <span>Dark Mode</span>
</label>

<!-- Slider -->
<div class="field">
  <label for="volume">Volume</label>
  <input 
    type="range" 
    id="volume" 
    min="0" 
    max="100" 
    value="75"
    data-tooltip="Current volume: 75%"
    data-side="top"
  >
</div>

<!-- Progress bar -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Upload Progress</span>
    <span>65%</span>
  </div>
  <div 
    class="progress"
    data-tooltip="Uploading... 65% complete"
    data-side="bottom"
  >
    <div class="progress-bar" style="width: 65%"></div>
  </div>
</div>
```

### Disabled Element Tooltips

```html
<!-- Disabled button with explanation -->
<button 
  class="btn" 
  disabled
  data-tooltip="Save is disabled until all required fields are completed"
  data-side="bottom"
>
  Save Changes
</button>

<!-- Disabled input -->
<input 
  type="text" 
  placeholder="Read only" 
  readonly
  data-tooltip="This field is automatically calculated"
  data-side="right"
>
```

### Badge and Status Tooltips

```html
<!-- Status badge -->
<span 
  class="badge-success"
  data-tooltip="Payment processed successfully on March 15, 2024"
  data-side="bottom"
>
  Paid
</span>

<!-- Priority badge -->
<span 
  class="badge-destructive"
  data-tooltip="Requires immediate attention"
  data-side="top"
>
  High Priority
</span>

<!-- Version badge -->
<span 
  class="badge-outline"
  data-tooltip="Released on March 10, 2024"
  data-side="bottom"
>
  v2.1.0
</span>
```

### Complex Content Tooltips

```html
<!-- Shortcut reference -->
<button 
  class="btn-outline"
  data-tooltip="Save (Ctrl+S)"
  data-side="bottom"
>
  Save
</button>

<!-- Feature description -->
<button 
  class="btn-ghost"
  data-tooltip="AI-powered content suggestions based on your writing style"
  data-side="right"
>
  Smart Compose
</button>

<!-- Technical info -->
<div 
  class="text-xs text-muted-foreground font-mono"
  data-tooltip="Last updated: 2024-03-15 14:23:15 UTC"
  data-side="top"
>
  #abc123
</div>
```

### Grid and Table Tooltips

```html
<!-- Table cell with overflow -->
<table class="table">
  <tbody>
    <tr>
      <td 
        class="max-w-xs truncate"
        data-tooltip="john.doe@example-company-with-long-domain.com"
        data-side="bottom"
      >
        john.doe@example-company...
      </td>
    </tr>
  </tbody>
</table>

<!-- Grid item -->
<div class="grid grid-cols-4 gap-4">
  <div 
    class="p-4 rounded border"
    data-tooltip="Click to view detailed analytics"
    data-side="top"
  >
    <div class="text-2xl font-bold">1.2k</div>
    <div class="text-sm text-muted-foreground">Views</div>
  </div>
</div>
```

## Accessibility Features

- **Keyboard Support**: Tooltips appear on focus for keyboard users
- **Mouse Support**: Tooltips appear on hover for mouse users
- **Screen Reader Friendly**: Uses appropriate ARIA attributes
- **Focus Management**: Doesn't interfere with keyboard navigation
- **Touch Friendly**: Works on touch devices with tap

### Enhanced Accessibility

```html
<!-- Using aria-label as fallback -->
<button 
  class="btn-icon-outline" 
  aria-label="Save document"
  data-tooltip="Save document"
  data-side="bottom"
>
  <svg><!-- save icon --></svg>
</button>

<!-- Descriptive tooltip for complex actions -->
<button 
  class="btn-destructive"
  aria-describedby="delete-tooltip"
  data-tooltip="Permanently delete this item. This action cannot be undone."
  data-side="top"
>
  Delete
</button>

<!-- Form field with help -->
<div class="field">
  <label for="password">Password</label>
  <input 
    type="password" 
    id="password"
    aria-describedby="password-help"
    data-tooltip="Must be at least 8 characters with uppercase, lowercase, number, and special character"
    data-side="right"
  >
</div>
```

## JavaScript Integration

### Dynamic Tooltip Content

```javascript
// Update tooltip content dynamically
function updateTooltip(element, newText) {
  element.setAttribute('data-tooltip', newText);
}

// Example: Update button tooltip based on state
const saveButton = document.getElementById('save-btn');
updateTooltip(saveButton, 'Save changes (Ctrl+S)');

// For form validation feedback
const emailInput = document.getElementById('email');
emailInput.addEventListener('blur', (e) => {
  const isValid = e.target.checkValidity();
  const tooltip = isValid 
    ? 'Valid email address' 
    : 'Please enter a valid email address';
  updateTooltip(e.target, tooltip);
});
```

### Conditional Tooltips

```javascript
// Show tooltip only when needed
function setConditionalTooltip(element, condition, tooltip) {
  if (condition) {
    element.setAttribute('data-tooltip', tooltip);
  } else {
    element.removeAttribute('data-tooltip');
  }
}

// Example: Truncated text
const textElement = document.querySelector('.truncated-text');
const isOverflowing = textElement.scrollWidth > textElement.clientWidth;
setConditionalTooltip(textElement, isOverflowing, textElement.textContent);

// Example: Disabled state explanation
const button = document.querySelector('#submit-btn');
const isDisabled = button.disabled;
setConditionalTooltip(button, isDisabled, 'Complete all required fields to enable');
```

### React Integration

```jsx
import React, { useState } from 'react';

function TooltipButton({ children, tooltip, side = 'top', align = 'center', ...props }) {
  return (
    <button
      data-tooltip={tooltip}
      data-side={side}
      data-align={align}
      {...props}
    >
      {children}
    </button>
  );
}

// Usage examples
function App() {
  const [saveStatus, setSaveStatus] = useState('idle');
  
  const getTooltipText = () => {
    switch (saveStatus) {
      case 'saving': return 'Saving changes...';
      case 'saved': return 'Changes saved successfully';
      case 'error': return 'Error saving changes';
      default: return 'Save changes (Ctrl+S)';
    }
  };
  
  return (
    <div>
      <TooltipButton 
        tooltip={getTooltipText()}
        side="bottom"
        onClick={handleSave}
        disabled={saveStatus === 'saving'}
      >
        Save
      </TooltipButton>
      
      <TooltipButton 
        tooltip="Delete this item permanently"
        side="top"
        className="btn-destructive"
      >
        Delete
      </TooltipButton>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div>
    <button 
      :data-tooltip="tooltipText"
      :data-side="side"
      :data-align="align"
      v-bind="$attrs"
    >
      <slot />
    </button>
  </div>
</template>

<script>
export default {
  props: {
    tooltip: {
      type: String,
      required: true
    },
    side: {
      type: String,
      default: 'top',
      validator: value => ['top', 'bottom', 'left', 'right'].includes(value)
    },
    align: {
      type: String,
      default: 'center',
      validator: value => ['start', 'center', 'end'].includes(value)
    }
  },
  computed: {
    tooltipText() {
      return this.tooltip;
    }
  }
};
</script>
```

### Tooltip Management Service

```javascript
// Centralized tooltip management
class TooltipManager {
  static updateTooltip(selector, text, options = {}) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    element.setAttribute('data-tooltip', text);
    if (options.side) element.setAttribute('data-side', options.side);
    if (options.align) element.setAttribute('data-align', options.align);
  }
  
  static removeTooltip(selector) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    element.removeAttribute('data-tooltip');
    element.removeAttribute('data-side');
    element.removeAttribute('data-align');
  }
  
  static setTemporaryTooltip(selector, text, duration = 3000) {
    const element = document.querySelector(selector);
    if (!element) return;
    
    const originalTooltip = element.getAttribute('data-tooltip');
    element.setAttribute('data-tooltip', text);
    
    setTimeout(() => {
      if (originalTooltip) {
        element.setAttribute('data-tooltip', originalTooltip);
      } else {
        element.removeAttribute('data-tooltip');
      }
    }, duration);
  }
}

// Usage
TooltipManager.updateTooltip('#save-btn', 'Saved successfully!', { side: 'bottom' });
TooltipManager.setTemporaryTooltip('#copy-btn', 'Copied to clipboard!');
```

## Best Practices

1. **Concise Content**: Keep tooltip text short and informative
2. **Essential Information**: Only include tooltips for helpful context
3. **Consistent Positioning**: Use consistent sides for similar elements
4. **Keyboard Accessible**: Ensure tooltips work with keyboard navigation
5. **Touch Friendly**: Consider touch device interactions
6. **No Critical Info**: Don't put essential information only in tooltips
7. **Clear Language**: Use simple, clear language
8. **Performance**: Avoid too many tooltips on complex pages

## Common Patterns

### Icon Button Explanations

```html
<div class="flex gap-2">
  <button class="btn-icon-outline" data-tooltip="Bold (Ctrl+B)" data-side="bottom">
    <svg><!-- bold icon --></svg>
  </button>
  <button class="btn-icon-outline" data-tooltip="Italic (Ctrl+I)" data-side="bottom">
    <svg><!-- italic icon --></svg>
  </button>
  <button class="btn-icon-outline" data-tooltip="Underline (Ctrl+U)" data-side="bottom">
    <svg><!-- underline icon --></svg>
  </button>
</div>
```

### Status Indicators

```html
<div class="flex items-center gap-2">
  <div 
    class="size-2 rounded-full bg-green-500"
    data-tooltip="Service operational"
    data-side="top"
  ></div>
  <span>API Status</span>
</div>
```

### Form Field Help

```html
<div class="field">
  <label for="api-key">API Key</label>
  <input 
    type="password" 
    id="api-key"
    data-tooltip="Found in your account settings under 'API Access'"
    data-side="right"
  >
</div>
```

## Related Components

- [Popover](./popover.md) - For more complex tooltip content
- [Button](./button.md) - Common element that uses tooltips
- [Dialog](./dialog.md) - For detailed help content
- [Badge](./badge.md) - Often used with explanatory tooltips
---
## input-group

# Input Group Pattern

Display additional information or actions to an input or textarea.

**Note:** There is no dedicated Input Group component in Basecoat. This pattern is achieved using positioning utilities and relative containers.

## Basic Usage

```html
<div class="relative">
  <input type="text" class="input pl-9 pr-20" placeholder="Search...">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">12 results</div>
</div>
```

## CSS Classes

### Container Classes
- **`relative`** - Required for positioning child elements

### Input Classes
- **`input`** - Base input styling
- **`textarea`** - Base textarea styling
- **Padding adjustments**: `pl-9`, `pr-20`, etc. to make room for positioned elements

### Positioned Element Classes
- **`absolute`** - Absolute positioning
- **`left-3`**, `right-3` - Horizontal positioning
- **`top-1/2 -translate-y-1/2`** - Vertical centering
- **`pointer-events-none`** - Disable mouse interaction (for decorative elements)

### Icon and Text Styling
- **`text-muted-foreground`** - Muted text color
- **`[&>svg]:size-4`** - Icon sizing
- **`text-sm`** - Small text size

## Component Attributes

### No specific component attributes - uses standard HTML input/textarea attributes
| Element | Attributes | Description |
|---------|------------|-------------|
| Container | `class="relative"` | Enables absolute positioning |
| Input | Standard input attributes | `type`, `placeholder`, `class`, etc. |
| Positioned elements | `class` with positioning utilities | Absolute positioning and styling |

### No JavaScript Required (Basic)
Basic input groups work with pure CSS positioning.

## HTML Structure

```html
<!-- Basic pattern -->
<div class="relative">
  <input type="text" class="input [padding-adjustments]" placeholder="...">
  <div class="absolute [positioning] [styling]">
    <!-- Icon, text, or interactive element -->
  </div>
</div>
```

## Examples

### Search Input with Icon and Results

```html
<div class="relative">
  <input type="text" class="input pl-9 pr-20" placeholder="Search...">
  <!-- Left icon -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <!-- Right text -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">12 results</div>
</div>
```

### URL Input with Protocol Prefix

```html
<div class="relative">
  <input type="text" class="input pl-15 pr-9" placeholder="example.com">
  <!-- Left prefix text -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">https://</div>
  <!-- Right help icon -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Enter your domain name">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="M12 16v-4" />
      <path d="M12 8h.01" />
    </svg>
  </div>
</div>
```

### Username Input with Validation

```html
<div class="relative">
  <input type="text" class="input pr-9" placeholder="@shadcn">
  <!-- Validation indicator -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-primary text-primary-foreground flex size-4 items-center justify-center rounded-full [&>svg]:size-3">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 6 9 17l-5-5" />
    </svg>
  </div>
</div>
```

### Enhanced Textarea with Controls

```html
<div class="relative">
  <textarea class="textarea pr-10 min-h-27 pb-12" placeholder="Ask, Search or Chat..."></textarea>
  <!-- Bottom control bar -->
  <footer role="group" class="absolute bottom-0 px-3 pb-3 pt-1.5 flex items-center w-full gap-2">
    <!-- Add button -->
    <button type="button" class="btn-icon-outline rounded-full size-6 text-muted-foreground hover:text-accent-foreground">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M5 12h14" />
        <path d="M12 5v14" />
      </svg>
    </button>
    
    <!-- Dropdown menu -->
    <div class="dropdown-menu">
      <button type="button" class="btn-sm-ghost text-muted-foreground hover:text-accent-foreground h-6 p-2">
        Auto
      </button>
      <div data-popover aria-hidden="true" data-side="top" class="min-w-32">
        <div role="menu">
          <div role="menuitem">Auto</div>
          <div role="menuitem">Agent</div>
          <div role="menuitem">Manual</div>
        </div>
      </div>
    </div>
    
    <!-- Usage indicator -->
    <div class="text-muted-foreground text-sm ml-auto">52% used</div>
    
    <!-- Separator -->
    <hr class="w-0 h-4 border-r border-border shrink-0 m-0">
    
    <!-- Submit button -->
    <button type="button" class="btn-icon rounded-full size-6 bg-muted-foreground hover:bg-foreground" disabled>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m5 12 7-7 7 7" />
        <path d="M12 19V5" />
      </svg>
    </button>
  </footer>
</div>
```

### Email Input with Domain

```html
<div class="relative">
  <input type="email" class="input pr-24" placeholder="username">
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">@company.com</div>
</div>
```

### Currency Input

```html
<div class="relative">
  <input type="number" class="input pl-8 pr-12" placeholder="0.00">
  <!-- Currency symbol -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">$</div>
  <!-- Unit -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">USD</div>
</div>
```

### Password Input with Strength

```html
<div class="relative">
  <input type="password" class="input pr-20" placeholder="Password">
  <!-- Strength indicator -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-1">
    <div class="text-xs text-muted-foreground">Strong</div>
    <div class="flex gap-1">
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
      <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
    </div>
  </div>
</div>
```

### Phone Number Input

```html
<div class="relative">
  <input type="tel" class="input pl-16 pr-9" placeholder="123-456-7890">
  <!-- Country code -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">+1</div>
  <!-- Format help -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Format: XXX-XXX-XXXX">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="M12 16v-4" />
      <path d="M12 8h.01" />
    </svg>
  </div>
</div>
```

### File Size Input

```html
<div class="relative">
  <input type="number" class="input pr-12" placeholder="100" min="1">
  <!-- Unit selector -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2">
    <select class="text-sm text-muted-foreground bg-transparent border-none focus:outline-none">
      <option>KB</option>
      <option>MB</option>
      <option>GB</option>
    </select>
  </div>
</div>
```

### Search with Clear Button

```html
<div class="relative">
  <input type="search" class="input pl-9 pr-9" placeholder="Search..." id="search-input">
  <!-- Search icon -->
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  </div>
  <!-- Clear button (interactive) -->
  <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground [&>svg]:size-4" onclick="document.getElementById('search-input').value = ''">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <circle cx="12" cy="12" r="10" />
      <path d="m15 9-6 6" />
      <path d="m9 9 6 6" />
    </svg>
  </button>
</div>
```

### Loading State Input

```html
<div class="relative">
  <input type="text" class="input pr-9" placeholder="Processing..." disabled>
  <!-- Loading spinner -->
  <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground">
    <svg class="animate-spin size-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="m12 2 0 4c-4.418 0-8 3.582-8 8 0 1.1.224 2.148.63 3.1L2.369 18.9C1.502 17.065 1 15.087 1 13c0-6.075 4.925-11 11-11Z"></path>
    </svg>
  </div>
</div>
```

### Multi-element Input Group

```html
<div class="flex">
  <!-- Input with left addon -->
  <div class="relative flex-1">
    <input type="text" class="input rounded-r-none pr-9" placeholder="repository-name">
    <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4" data-tooltip="Repository name">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="10" />
        <path d="M12 16v-4" />
        <path d="M12 8h.01" />
      </svg>
    </div>
  </div>
  <!-- Button addon -->
  <button type="button" class="btn-outline rounded-l-none border-l-0">Clone</button>
</div>
```

## Accessibility Features

- **Proper Labeling**: Use labels or `aria-label` for inputs
- **Interactive Elements**: Ensure clickable elements are accessible
- **Focus Management**: Positioned elements shouldn't interfere with focus
- **Screen Reader Support**: Use `aria-describedby` for helpful text

### Enhanced Accessibility

```html
<!-- Input with accessible help text -->
<div class="field">
  <label for="username-input">Username</label>
  <div class="relative">
    <input 
      type="text" 
      id="username-input"
      class="input pl-8"
      placeholder="username"
      aria-describedby="username-help"
    >
    <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm">@</div>
  </div>
  <p id="username-help" class="text-sm text-muted-foreground mt-1">Choose a unique username for your account</p>
</div>

<!-- Search with live region -->
<div class="relative">
  <input 
    type="search" 
    class="input pl-9 pr-20" 
    placeholder="Search..."
    aria-describedby="search-results"
  >
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg><!-- search icon --></svg>
  </div>
  <div 
    id="search-results"
    class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground text-sm"
    aria-live="polite"
  >
    12 results
  </div>
</div>
```

## JavaScript Integration

### Dynamic Content Updates

```javascript
// Update result count
function updateSearchResults(count) {
  const resultsElement = document.querySelector('#search-results');
  if (resultsElement) {
    resultsElement.textContent = `${count} results`;
  }
}

// Toggle validation state
function updateValidationState(input, isValid) {
  const container = input.parentElement;
  const indicator = container.querySelector('.validation-indicator');
  
  if (indicator) {
    if (isValid) {
      indicator.className = 'absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-green-500 text-white flex size-4 items-center justify-center rounded-full [&>svg]:size-3';
      indicator.innerHTML = '<svg><!-- checkmark --></svg>';
    } else {
      indicator.className = 'absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none bg-red-500 text-white flex size-4 items-center justify-center rounded-full [&>svg]:size-3';
      indicator.innerHTML = '<svg><!-- x mark --></svg>';
    }
  }
}

// Clear input
function clearInput(inputId) {
  const input = document.getElementById(inputId);
  if (input) {
    input.value = '';
    input.focus();
  }
}
```

### React Integration

```jsx
import React, { useState } from 'react';

function InputGroup({ 
  children, 
  leftElement, 
  rightElement, 
  className = '',
  ...props 
}) {
  return (
    <div className={`relative ${className}`}>
      {React.cloneElement(children, {
        className: `input ${leftElement ? 'pl-9' : ''} ${rightElement ? 'pr-9' : ''} ${children.props.className || ''}`,
        ...props
      })}
      {leftElement && (
        <div className="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
          {leftElement}
        </div>
      )}
      {rightElement && (
        <div className="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4">
          {rightElement}
        </div>
      )}
    </div>
  );
}

// Usage examples
function App() {
  const [searchValue, setSearchValue] = useState('');
  const [results, setResults] = useState(0);
  
  return (
    <div className="space-y-4">
      <InputGroup
        leftElement={<SearchIcon />}
        rightElement={<span className="text-sm pointer-events-none">{results} results</span>}
      >
        <input 
          type="search" 
          placeholder="Search..."
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
        />
      </InputGroup>
      
      <InputGroup
        leftElement={<span className="text-sm">https://</span>}
        rightElement={
          <button onClick={() => console.log('Help clicked')}>
            <InfoIcon />
          </button>
        }
      >
        <input type="url" placeholder="example.com" />
      </InputGroup>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <div class="relative">
    <input 
      v-bind="$attrs"
      :class="inputClasses"
      @input="$emit('input', $event.target.value)"
    />
    
    <div v-if="leftElement" class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
      <slot name="left" />
    </div>
    
    <div v-if="rightElement" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground [&>svg]:size-4">
      <slot name="right" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    leftElement: Boolean,
    rightElement: Boolean
  },
  computed: {
    inputClasses() {
      let classes = 'input';
      if (this.leftElement) classes += ' pl-9';
      if (this.rightElement) classes += ' pr-9';
      return classes;
    }
  }
};
</script>
```

## Best Practices

1. **Padding Adjustments**: Always adjust input padding to accommodate positioned elements
2. **Pointer Events**: Use `pointer-events-none` for decorative elements
3. **Interactive Elements**: Make buttons and clickable icons accessible
4. **Consistent Sizing**: Use consistent icon sizes and positioning
5. **Visual Hierarchy**: Use muted colors for secondary elements
6. **Mobile Friendly**: Ensure touch targets are large enough
7. **Loading States**: Show loading indicators for async operations

## Common Patterns

### Form Field Enhancement

```html
<div class="field">
  <label for="enhanced-input">Enhanced Input</label>
  <div class="relative">
    <input type="text" id="enhanced-input" class="input pl-9 pr-9" placeholder="Enter value">
    <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
      <svg><!-- icon --></svg>
    </div>
    <button type="button" class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground hover:text-foreground">
      <svg><!-- action icon --></svg>
    </button>
  </div>
  <p class="text-sm text-muted-foreground mt-1">Helper text for this field</p>
</div>
```

### Search Interface

```html
<div class="relative">
  <input type="search" class="input pl-9 pr-24" placeholder="Search anything...">
  <div class="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-muted-foreground [&>svg]:size-4">
    <svg><!-- search icon --></svg>
  </div>
  <div class="absolute right-3 top-1/2 -translate-y-1/2 flex items-center gap-2">
    <span class="text-xs text-muted-foreground">⌘K</span>
  </div>
</div>
```

### Status Input

```html
<div class="relative">
  <input type="text" class="input pr-9" readonly value="Connected">
  <div class="absolute right-3 top-1/2 -translate-y-1/2 pointer-events-none">
    <div class="size-2 rounded-full bg-green-500"></div>
  </div>
</div>
```

## Related Components

- [Input](./input.md) - Base input component
- [Textarea](./textarea.md) - Base textarea component
- [Button](./button.md) - For interactive elements
- [Field](./field.md) - For complete form field structure
---
## item

# Item Pattern

A versatile component that you can use to display any content.

**Note:** There is no dedicated Item component in Basecoat. Items are pure HTML composition using Tailwind utility classes.

## Basic Usage

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Basic Item</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">A simple item with title and description.</p>
  </div>
  <button class="btn-sm-outline">Action</button>
</article>
```

## CSS Classes

### Core Item Classes
- **`group/item`** - Groups hover and focus states
- **`flex items-center`** - Horizontal layout with center alignment
- **`border rounded-md`** - Basic styling
- **`text-sm`** - Text size
- **`transition-colors`** - Smooth color transitions

### Layout Classes
- **`p-4 gap-4`** - Default padding and spacing
- **`py-3 px-4 gap-2.5`** - Compact size
- **`p-6 gap-6`** - Large size

### Interaction Classes
- **`[a]:hover:bg-accent/50`** - Hover state for links
- **`outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]`** - Focus states

### Content Classes
- **`flex flex-1 flex-col gap-1`** - Content area layout
- **`text-muted-foreground line-clamp-2`** - Description styling
- **`flex shrink-0 items-center justify-center`** - Icon/avatar container

## Component Attributes

### Container Attributes
| Attribute | Element | Description | Required |
|-----------|---------|-------------|----------|
| `class` | Any container | Must include base item classes | Yes |
| `href` | `<a>` | For clickable items | Optional |
| `role` | Various | Semantic role (article, button, etc.) | Recommended |

### No JavaScript Required (Basic)
Basic items work with pure CSS styling and HTML structure.

## HTML Structure

```html
<!-- Basic item pattern -->
<article class="group/item flex items-center border rounded-md p-4 gap-4 [base-classes]">
  <!-- Optional: Icon/Avatar -->
  <div class="flex shrink-0 items-center justify-center [icon-classes]">
    <!-- Icon or avatar content -->
  </div>
  
  <!-- Main content -->
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="text-sm leading-snug font-medium">Title</h3>
    <p class="text-muted-foreground text-sm">Description</p>
  </div>
  
  <!-- Optional: Actions -->
  <button class="btn-sm-outline">Action</button>
</article>
```

## Examples

### Basic Item with Action

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Task Completed</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">Your data backup has finished successfully.</p>
  </div>
  <button class="btn-sm-outline">View Details</button>
</article>
```

### Clickable Link Item

```html
<a href="/notifications/123" class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border py-3 px-4 gap-2.5">
  <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M3.85 8.62a4 4 0 0 1 4.78-4.77 4 4 0 0 1 6.74 0 4 4 0 0 1 4.78 4.78 4 4 0 0 1 0 6.74 4 4 0 0 1-4.77 4.78 4 4 0 0 1-6.75 0 4 4 0 0 1-4.78-4.77 4 4 0 0 1 0-6.76Z" />
      <path d="m9 12 2 2 4-4" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Your profile has been verified.</h3>
  </div>
  <div class="flex items-center gap-2 [&_svg]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="m9 18 6-6-6-6" />
    </svg>
  </div>
</a>
```

### Item Variants

```html
<!-- Default variant -->
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-transparent p-4 gap-4">
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Default Variant</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">Standard styling with subtle background and borders.</p>
  </div>
  <button class="btn-sm-outline">Open</button>
</article>

<!-- Outline variant -->
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Outline Variant</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">Outlined style with clear borders and transparent background.</p>
  </div>
  <button class="btn-sm-outline">Open</button>
</article>

<!-- Muted variant -->
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-transparent bg-muted/50 p-4 gap-4">
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Muted Variant</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">Subdued appearance with muted colors for secondary content.</p>
  </div>
  <button class="btn-sm-outline">Open</button>
</article>
```

### Item with Icon

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- Icon container -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none size-8 border rounded-sm bg-muted [&_svg:not([class*='size-'])]:size-4">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 13c0 5-3.5 7.5-7.66 8.95a1 1 0 0 1-.67-.01C7.5 20.5 4 18 4 13V6a1 1 0 0 1 1-1c2 0 4.5-1.2 6.24-2.72a1.17 1.17 0 0 1 1.52 0C14.51 3.81 17 5 19 5a1 1 0 0 1 1 1z" />
      <path d="M12 8v4" />
      <path d="M12 16h.01" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">Security Alert</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">New login detected from unknown device.</p>
  </div>
  <button class="btn-sm-outline">Review</button>
</article>
```

### Item with Avatar

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- Avatar -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none">
    <img src="/avatars/user1.jpg" alt="John Doe" class="size-8 rounded-full">
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">John Doe mentioned you</h3>
    <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance [&>a:hover]:text-primary [&>a]:underline [&>a]:underline-offset-4">@sarah can you review this pull request?</p>
  </div>
  <div class="text-xs text-muted-foreground">2m ago</div>
</article>
```

### Notification Item

```html
<div class="flex flex-col gap-2">
  <!-- Unread notification -->
  <article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border bg-blue-50 dark:bg-blue-950/30 p-4 gap-4">
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-blue-500 text-white size-8 rounded-full [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9" />
        <path d="m13.73 21a2 2 0 0 1-3.46 0" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col gap-1">
      <div class="flex items-center gap-2">
        <h3 class="text-sm leading-snug font-medium">New message received</h3>
        <div class="size-2 rounded-full bg-blue-500"></div>
      </div>
      <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance">You have a new message from the support team.</p>
    </div>
    <div class="flex flex-col items-end gap-1">
      <div class="text-xs text-muted-foreground">5m ago</div>
      <button class="btn-sm-outline">View</button>
    </div>
  </article>

  <!-- Read notification -->
  <article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-muted size-8 rounded-full [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M20 6 9 17l-5-5" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col gap-1">
      <h3 class="text-sm leading-snug font-medium">Task completed</h3>
      <p class="text-muted-foreground line-clamp-2 text-sm leading-normal font-normal text-balance">Your backup has finished successfully.</p>
    </div>
    <div class="text-xs text-muted-foreground">1h ago</div>
  </article>
</div>
```

### File Item

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- File icon -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none size-10 border rounded bg-muted [&_svg]:size-5">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
      <polyline points="14,2 14,8 20,8" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">project-proposal.pdf</h3>
    <p class="text-muted-foreground text-sm">2.4 MB • Modified 2 hours ago</p>
  </div>
  <div class="flex items-center gap-2">
    <button class="btn-sm-icon-ghost" data-tooltip="Download">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
        <polyline points="7,10 12,15 17,10" />
        <line x1="12" x2="12" y1="15" y2="3" />
      </svg>
    </button>
    <button class="btn-sm-icon-ghost" data-tooltip="More options">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="12" cy="12" r="1" />
        <circle cx="19" cy="12" r="1" />
        <circle cx="5" cy="12" r="1" />
      </svg>
    </button>
  </div>
</article>
```

### Contact Item

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- Avatar -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none">
    <img src="/avatars/sarah.jpg" alt="Sarah Wilson" class="size-10 rounded-full">
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">
      Sarah Wilson
      <span class="badge-outline">Team Lead</span>
    </h3>
    <p class="text-muted-foreground text-sm">sarah.wilson@company.com</p>
    <div class="flex items-center gap-2 mt-1">
      <div class="size-2 rounded-full bg-green-500"></div>
      <span class="text-xs text-muted-foreground">Online</span>
    </div>
  </div>
  <div class="flex flex-col gap-2">
    <button class="btn-sm-outline">Message</button>
    <button class="btn-sm-ghost">Call</button>
  </div>
</article>
```

### Product Item

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- Product image -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none">
    <img src="/products/laptop.jpg" alt="MacBook Pro" class="size-16 rounded border object-cover">
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="flex w-fit items-center gap-2 text-sm leading-snug font-medium">MacBook Pro 16-inch</h3>
    <p class="text-muted-foreground text-sm">Apple M2 Pro chip • 16GB RAM • 512GB SSD</p>
    <div class="flex items-center gap-2 mt-1">
      <span class="text-lg font-semibold">$2,399</span>
      <span class="badge-outline">In Stock</span>
    </div>
  </div>
  <div class="flex flex-col gap-2">
    <button class="btn-sm">Add to Cart</button>
    <button class="btn-sm-ghost">Save</button>
  </div>
</article>
```

### Activity Item

```html
<article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-3 gap-3">
  <!-- Timeline indicator -->
  <div class="flex shrink-0 items-center justify-center gap-2 self-start [&_svg]:pointer-events-none size-8 border-2 border-green-500 bg-green-50 dark:bg-green-950 rounded-full [&_svg]:size-4 text-green-600">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M20 6 9 17l-5-5" />
    </svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="text-sm leading-snug font-medium">Deployment successful</h3>
    <p class="text-muted-foreground text-sm">Version 2.1.0 deployed to production</p>
    <div class="text-xs text-muted-foreground">2:30 PM • Deploy #142</div>
  </div>
  <button class="btn-sm-ghost">View Logs</button>
</article>
```

### Compact Size Items

```html
<div class="flex flex-col gap-2">
  <!-- Small padding for compact lists -->
  <article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border py-2 px-3 gap-3">
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
        <polyline points="14,2 14,8 20,8" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col">
      <h3 class="text-sm leading-snug font-medium">Quick item</h3>
    </div>
    <div class="text-xs text-muted-foreground">2MB</div>
  </article>

  <article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border py-2 px-3 gap-3">
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <rect width="18" height="18" x="3" y="4" rx="2" ry="2" />
        <line x1="16" x2="16" y1="2" y2="6" />
        <line x1="8" x2="8" y1="2" y2="6" />
        <line x1="3" x2="21" y1="10" y2="10" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col">
      <h3 class="text-sm leading-snug font-medium">Meeting notes</h3>
    </div>
    <div class="text-xs text-muted-foreground">1.2MB</div>
  </article>
</div>
```

### Selection Items

```html
<div class="flex flex-col gap-2">
  <!-- Selectable items with checkboxes -->
  <label class="group/item flex items-center border text-sm rounded-md transition-colors hover:bg-accent/50 cursor-pointer outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-3 gap-3">
    <input type="checkbox" class="checkbox">
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z" />
        <polyline points="14,2 14,8 20,8" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col gap-1">
      <h3 class="text-sm leading-snug font-medium">Document.pdf</h3>
      <p class="text-muted-foreground text-sm">2.4 MB • Modified today</p>
    </div>
    <div class="text-xs text-muted-foreground">Select</div>
  </label>

  <label class="group/item flex items-center border text-sm rounded-md transition-colors hover:bg-accent/50 cursor-pointer outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-3 gap-3">
    <input type="checkbox" class="checkbox" checked>
    <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <rect x="3" y="3" width="18" height="18" rx="2" />
        <path d="M9 12l2 2 4-4" />
      </svg>
    </div>
    <div class="flex flex-1 flex-col gap-1">
      <h3 class="text-sm leading-snug font-medium">Spreadsheet.xlsx</h3>
      <p class="text-muted-foreground text-sm">1.8 MB • Modified yesterday</p>
    </div>
    <div class="text-xs text-primary">Selected</div>
  </label>
</div>
```

## Accessibility Features

- **Semantic HTML**: Use appropriate elements (`article`, `a`, `button`)
- **Keyboard Navigation**: Focus states and proper tab order
- **Screen Reader Support**: Proper heading hierarchy and content structure
- **Interactive States**: Clear hover and focus indicators
- **ARIA Attributes**: When needed for complex interactions

### Enhanced Accessibility

```html
<!-- Proper semantic structure -->
<article 
  class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4"
  aria-labelledby="item-title-123"
  aria-describedby="item-desc-123"
>
  <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-5" aria-hidden="true">
    <svg><!-- decorative icon --></svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 id="item-title-123" class="text-sm leading-snug font-medium">Accessible Item</h3>
    <p id="item-desc-123" class="text-muted-foreground text-sm">This item has proper accessibility attributes.</p>
  </div>
  <button class="btn-sm-outline" aria-label="View details for Accessible Item">
    View
  </button>
</article>

<!-- Clickable item with proper link semantics -->
<a 
  href="/item/123" 
  class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4"
  aria-describedby="item-meta-123"
>
  <div class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-5" aria-hidden="true">
    <svg><!-- decorative icon --></svg>
  </div>
  <div class="flex flex-1 flex-col gap-1">
    <h3 class="text-sm leading-snug font-medium">Clickable Item Title</h3>
    <p id="item-meta-123" class="text-muted-foreground text-sm">Additional context about this clickable item</p>
  </div>
  <div class="flex items-center gap-2 [&_svg]:size-4" aria-hidden="true">
    <svg><!-- arrow icon --></svg>
  </div>
</a>
```

## JavaScript Integration

### Selection Management

```javascript
// Handle item selection
function initializeItemSelection() {
  const items = document.querySelectorAll('.group\\/item input[type="checkbox"]');
  const selectAllBtn = document.getElementById('select-all');
  
  // Update select all state
  function updateSelectAllState() {
    const totalItems = items.length;
    const selectedItems = Array.from(items).filter(item => item.checked).length;
    
    if (selectAllBtn) {
      selectAllBtn.checked = selectedItems === totalItems;
      selectAllBtn.indeterminate = selectedItems > 0 && selectedItems < totalItems;
    }
  }
  
  // Handle individual item selection
  items.forEach(item => {
    item.addEventListener('change', updateSelectAllState);
  });
  
  // Handle select all
  selectAllBtn?.addEventListener('change', (e) => {
    items.forEach(item => {
      item.checked = e.target.checked;
    });
  });
}

// Get selected items
function getSelectedItems() {
  const selected = [];
  document.querySelectorAll('.group\\/item input[type="checkbox"]:checked').forEach(checkbox => {
    const item = checkbox.closest('.group\\/item');
    const title = item.querySelector('h3')?.textContent;
    const id = item.dataset.id;
    selected.push({ id, title, element: item });
  });
  return selected;
}
```

### React Integration

```jsx
import React, { useState } from 'react';

function Item({ 
  title, 
  description, 
  icon, 
  actions, 
  href, 
  variant = 'default',
  size = 'default',
  className = '',
  ...props 
}) {
  const baseClasses = "group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]";
  
  const variantClasses = {
    default: "border-transparent",
    outline: "border-border",
    muted: "border-transparent bg-muted/50"
  };
  
  const sizeClasses = {
    compact: "py-2 px-3 gap-3",
    default: "p-4 gap-4",
    large: "p-6 gap-6"
  };
  
  const classes = `${baseClasses} ${variantClasses[variant]} ${sizeClasses[size]} ${className}`;
  
  const content = (
    <>
      {icon && (
        <div className="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-5">
          {icon}
        </div>
      )}
      <div className="flex flex-1 flex-col gap-1">
        <h3 className="text-sm leading-snug font-medium">{title}</h3>
        {description && (
          <p className="text-muted-foreground text-sm line-clamp-2">{description}</p>
        )}
      </div>
      {actions && (
        <div className="flex items-center gap-2">
          {actions}
        </div>
      )}
    </>
  );
  
  if (href) {
    return (
      <a href={href} className={classes} {...props}>
        {content}
        <div className="flex items-center gap-2 [&_svg]:size-4">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="m9 18 6-6-6-6" />
          </svg>
        </div>
      </a>
    );
  }
  
  return (
    <article className={classes} {...props}>
      {content}
    </article>
  );
}

// Usage examples
function ItemList() {
  const [selectedItems, setSelectedItems] = useState([]);
  
  const handleItemSelect = (id, selected) => {
    if (selected) {
      setSelectedItems([...selectedItems, id]);
    } else {
      setSelectedItems(selectedItems.filter(item => item !== id));
    }
  };
  
  return (
    <div className="flex flex-col gap-4">
      <Item
        title="Task Completed"
        description="Your data backup has finished successfully."
        icon={<CheckIcon />}
        actions={<button className="btn-sm-outline">View Details</button>}
      />
      
      <Item
        title="View Profile"
        description="Click to view your profile settings"
        href="/profile"
        variant="outline"
      />
      
      <Item
        title="Archive Item"
        description="This item is archived and read-only"
        variant="muted"
        size="compact"
      />
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <component
    :is="href ? 'a' : 'article'"
    :href="href"
    :class="itemClasses"
    v-bind="$attrs"
  >
    <div v-if="icon" class="flex shrink-0 items-center justify-center gap-2 [&_svg]:pointer-events-none bg-transparent [&_svg]:size-5">
      <slot name="icon">{{ icon }}</slot>
    </div>
    
    <div class="flex flex-1 flex-col gap-1">
      <h3 class="text-sm leading-snug font-medium">{{ title }}</h3>
      <p v-if="description" class="text-muted-foreground text-sm line-clamp-2">{{ description }}</p>
    </div>
    
    <div v-if="$slots.actions" class="flex items-center gap-2">
      <slot name="actions" />
    </div>
    
    <div v-if="href" class="flex items-center gap-2 [&_svg]:size-4">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m9 18 6-6-6-6" />
      </svg>
    </div>
  </component>
</template>

<script>
export default {
  props: {
    title: { type: String, required: true },
    description: String,
    icon: String,
    href: String,
    variant: { type: String, default: 'default' },
    size: { type: String, default: 'default' }
  },
  computed: {
    itemClasses() {
      const baseClasses = "group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]";
      
      const variantClasses = {
        default: "border-transparent",
        outline: "border-border",
        muted: "border-transparent bg-muted/50"
      };
      
      const sizeClasses = {
        compact: "py-2 px-3 gap-3",
        default: "p-4 gap-4",
        large: "p-6 gap-6"
      };
      
      return `${baseClasses} ${variantClasses[this.variant]} ${sizeClasses[this.size]}`;
    }
  }
};
</script>
```

## Best Practices

1. **Semantic HTML**: Use appropriate semantic elements (`article`, `a`, `section`)
2. **Consistent Layout**: Follow the flex layout pattern for alignment
3. **Visual Hierarchy**: Use proper heading levels and text sizing
4. **Interactive States**: Ensure clear hover and focus indicators
5. **Content Structure**: Keep titles concise and descriptions informative
6. **Icon Usage**: Use consistent icon sizing and placement
7. **Spacing**: Maintain consistent padding and gaps across items
8. **Accessibility**: Include proper labels and ARIA attributes

## Common Patterns

### List Items

```html
<div class="flex flex-col gap-2">
  <!-- Multiple items in a list -->
  <article class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-3 gap-3">
    <!-- Item content -->
  </article>
  <!-- More items... -->
</div>
```

### Card-style Items

```html
<div class="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
  <!-- Items in a grid layout -->
  <article class="group/item flex flex-col border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
    <!-- Vertical card layout -->
  </article>
</div>
```

### Interactive Lists

```html
<!-- For clickable items that navigate -->
<a href="/item/1" class="group/item flex items-center border text-sm rounded-md transition-colors [a]:hover:bg-accent/50 [a]:transition-colors duration-100 flex-wrap outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] border-border p-4 gap-4">
  <!-- Link content -->
</a>
```

## Related Components

- [Card](./card.md) - For more complex content containers
- [Button](./button.md) - For item actions
- [Avatar](./avatar.md) - For user-related items
- [Badge](./badge.md) - For status indicators
- [Checkbox](./checkbox.md) - For selectable items
---
## carousel

# Carousel Component

A carousel with motion and swipe built with CSS scroll snapping.

## Basic Usage

```html
<div class="slider">
  <!-- Navigation links -->
  <a href="#slide-1">1</a>
  <a href="#slide-2">2</a>
  <a href="#slide-3">3</a>
  
  <!-- Slides container -->
  <div class="slides">
    <div id="slide-1">Slide 1 Content</div>
    <div id="slide-2">Slide 2 Content</div>
    <div id="slide-3">Slide 3 Content</div>
  </div>
</div>
```

## CSS Classes

### Container Classes
- **`slider`** - Main carousel container with fixed width and overflow hidden
- **`slides`** - Scrollable container with flex layout and scroll snap

### Custom CSS Required
The carousel component requires custom CSS for scroll snapping behavior:

```css
.slider {
  width: 300px;
  text-align: center;
  overflow: hidden;
}

.slides {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  scroll-behavior: smooth;
  -webkit-overflow-scrolling: touch;
  scrollbar-width: none; /* Firefox */
}

.slides::-webkit-scrollbar {
  display: none; /* Webkit browsers */
}

.slides > div {
  scroll-snap-align: start;
  flex-shrink: 0;
  width: 300px;
  height: 300px;
  margin-right: 50px;
  border-radius: 10px;
  background: #eee;
  transform-origin: center center;
  transform: scale(1);
  transition: transform 0.5s;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 100px;
}
```

### Tailwind Utility Alternative
Using Tailwind utilities for responsive carousel:

```css
/* Alternative Tailwind-based carousel */
.carousel-container {
  @apply w-full max-w-lg mx-auto overflow-hidden;
}

.carousel-slides {
  @apply flex overflow-x-auto snap-x snap-mandatory scroll-smooth;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.carousel-slides::-webkit-scrollbar {
  display: none;
}

.carousel-slide {
  @apply snap-start flex-shrink-0 w-full h-64 bg-muted rounded-lg mr-4 flex items-center justify-center text-2xl font-bold;
}
```

## Component Attributes

### Carousel Container
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "slider" for carousel styling | Yes |

### Slides Container  
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "slides" for scroll behavior | Yes |

### Individual Slides
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for navigation links | Yes |

### Navigation Links
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `href` | string | Points to slide ID (e.g., "#slide-1") | Yes |

## No JavaScript Required (Basic)
The carousel uses CSS scroll snapping and hash navigation, requiring no JavaScript for basic functionality.

## HTML Structure

```html
<!-- Basic carousel structure -->
<div class="slider">
  <!-- Navigation (optional) -->
  <nav class="carousel-nav">
    <a href="#slide-1">•</a>
    <a href="#slide-2">•</a>
    <a href="#slide-3">•</a>
  </nav>
  
  <!-- Slides container -->
  <div class="slides">
    <div id="slide-1" class="slide">
      <!-- Slide content -->
    </div>
    <div id="slide-2" class="slide">
      <!-- Slide content -->
    </div>
    <div id="slide-3" class="slide">
      <!-- Slide content -->
    </div>
  </div>
</div>
```

## Examples

### Image Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="image-1" class="slide">
      <img src="/images/photo1.jpg" alt="Beautiful landscape" class="w-full h-full object-cover rounded-lg">
    </div>
    <div id="image-2" class="slide">
      <img src="/images/photo2.jpg" alt="City skyline" class="w-full h-full object-cover rounded-lg">
    </div>
    <div id="image-3" class="slide">
      <img src="/images/photo3.jpg" alt="Ocean view" class="w-full h-full object-cover rounded-lg">
    </div>
  </div>
  
  <!-- Dot navigation -->
  <nav class="flex justify-center gap-2 mt-4">
    <a href="#image-1" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
    <a href="#image-2" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
    <a href="#image-3" class="w-3 h-3 rounded-full bg-muted hover:bg-primary transition-colors"></a>
  </nav>
</div>
```

### Product Showcase Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="product-1" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/laptop.jpg" alt="Premium Laptop" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Premium Laptop</h3>
        <p class="text-muted-foreground mb-4">High-performance laptop for professionals</p>
        <span class="text-2xl font-bold text-primary">$1,299</span>
      </div>
    </div>
    
    <div id="product-2" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/phone.jpg" alt="Smartphone" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Smartphone</h3>
        <p class="text-muted-foreground mb-4">Latest flagship with advanced camera</p>
        <span class="text-2xl font-bold text-primary">$899</span>
      </div>
    </div>
    
    <div id="product-3" class="slide bg-background border rounded-lg p-6">
      <div class="text-center">
        <img src="/products/tablet.jpg" alt="Tablet" class="w-32 h-32 mx-auto mb-4 object-cover rounded">
        <h3 class="text-lg font-semibold mb-2">Tablet</h3>
        <p class="text-muted-foreground mb-4">Versatile tablet for work and entertainment</p>
        <span class="text-2xl font-bold text-primary">$599</span>
      </div>
    </div>
  </div>
</div>
```

### Testimonials Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="testimonial-1" class="slide bg-muted rounded-lg p-8">
      <div class="text-center max-w-md mx-auto">
        <svg class="w-8 h-8 mx-auto mb-4 text-primary" fill="currentColor" viewBox="0 0 24 24">
          <path d="M14.17 18.45c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36zm-8 0c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36z"/>
        </svg>
        <blockquote class="text-lg italic mb-4">
          "This product has completely transformed our workflow. Highly recommended!"
        </blockquote>
        <div>
          <img src="/avatars/user1.jpg" alt="Sarah Johnson" class="w-12 h-12 rounded-full mx-auto mb-2">
          <cite class="font-semibold not-italic">Sarah Johnson</cite>
          <div class="text-sm text-muted-foreground">Product Manager</div>
        </div>
      </div>
    </div>
    
    <div id="testimonial-2" class="slide bg-muted rounded-lg p-8">
      <div class="text-center max-w-md mx-auto">
        <svg class="w-8 h-8 mx-auto mb-4 text-primary" fill="currentColor" viewBox="0 0 24 24">
          <path d="M14.17 18.45c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36zm-8 0c.4.72 1.47 1.23 2.58 1.23 1.61 0 2.91-1.3 2.91-2.91 0-1.61-1.3-2.91-2.91-2.91-.72 0-1.38.26-1.89.71-.51-.85-.84-1.83-.84-2.88 0-2.21 1.79-4 4-4v-1.5c-3.04 0-5.5 2.46-5.5 5.5 0 1.8.87 3.4 2.21 4.4l-.56 1.36z"/>
        </svg>
        <blockquote class="text-lg italic mb-4">
          "Amazing user experience and excellent customer support. Five stars!"
        </blockquote>
        <div>
          <img src="/avatars/user2.jpg" alt="Michael Chen" class="w-12 h-12 rounded-full mx-auto mb-2">
          <cite class="font-semibold not-italic">Michael Chen</cite>
          <div class="text-sm text-muted-foreground">Software Developer</div>
        </div>
      </div>
    </div>
  </div>
  
  <!-- Arrow navigation -->
  <div class="flex justify-between items-center mt-4">
    <a href="#testimonial-1" class="btn-icon-outline">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m15 18-6-6 6-6"/>
      </svg>
    </a>
    <div class="flex gap-2">
      <a href="#testimonial-1" class="w-2 h-2 rounded-full bg-primary"></a>
      <a href="#testimonial-2" class="w-2 h-2 rounded-full bg-muted"></a>
    </div>
    <a href="#testimonial-2" class="btn-icon-outline">
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <path d="m9 18 6-6-6-6"/>
      </svg>
    </a>
  </div>
</div>
```

### Card Carousel

```html
<div class="slider">
  <div class="slides">
    <div id="feature-1" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Easy to Use</h3>
        </div>
        <p class="text-muted-foreground">
          Get started in minutes with our intuitive interface and comprehensive documentation.
        </p>
      </div>
    </div>
    
    <div id="feature-2" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M13 2L3 14h9l-1 8 10-12h-9l1-8z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Lightning Fast</h3>
        </div>
        <p class="text-muted-foreground">
          Optimized performance ensures your application runs smoothly at any scale.
        </p>
      </div>
    </div>
    
    <div id="feature-3" class="slide">
      <div class="card p-6 h-full">
        <div class="flex items-center gap-3 mb-4">
          <div class="w-10 h-10 bg-primary/10 rounded-lg flex items-center justify-center">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-primary">
              <path d="M9 12l2 2 4-4"/>
              <path d="M21 12c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M3 12c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M12 21c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
              <path d="M12 3c.552 0 1-.448 1-1s-.448-1-1-1-1 .448-1 1 .448 1 1 1z"/>
            </svg>
          </div>
          <h3 class="text-lg font-semibold">Secure</h3>
        </div>
        <p class="text-muted-foreground">
          Built-in security features protect your data with enterprise-grade encryption.
        </p>
      </div>
    </div>
  </div>
</div>
```

### Responsive Multi-Item Carousel

```html
<style>
.multi-carousel {
  @apply w-full max-w-6xl mx-auto overflow-hidden;
}

.multi-slides {
  @apply flex overflow-x-auto snap-x snap-mandatory scroll-smooth gap-4 pb-4;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.multi-slides::-webkit-scrollbar {
  display: none;
}

.multi-slide {
  @apply snap-start flex-shrink-0 w-64 h-48 bg-muted rounded-lg p-4;
}

@media (min-width: 640px) {
  .multi-slide {
    @apply w-72;
  }
}

@media (min-width: 1024px) {
  .multi-slide {
    @apply w-80;
  }
}
</style>

<div class="multi-carousel">
  <div class="multi-slides">
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 1</h3>
      <p class="text-sm text-muted-foreground mb-4">Lorem ipsum dolor sit amet consectetur adipiscing elit...</p>
      <a href="/blog/post-1" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 2</h3>
      <p class="text-sm text-muted-foreground mb-4">Sed do eiusmod tempor incididunt ut labore et dolore...</p>
      <a href="/blog/post-2" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 3</h3>
      <p class="text-sm text-muted-foreground mb-4">Ut enim ad minim veniam quis nostrud exercitation...</p>
      <a href="/blog/post-3" class="text-primary text-sm hover:underline">Read more</a>
    </div>
    
    <div class="multi-slide">
      <h3 class="font-semibold mb-2">Blog Post 4</h3>
      <p class="text-sm text-muted-foreground mb-4">Duis aute irure dolor in reprehenderit in voluptate...</p>
      <a href="/blog/post-4" class="text-primary text-sm hover:underline">Read more</a>
    </div>
  </div>
</div>
```

## Accessibility Features

- **Keyboard Navigation**: Use arrow keys or tab to navigate
- **Focus Management**: Proper focus states for navigation links
- **Screen Reader Support**: Use descriptive alt text and labels
- **Reduced Motion**: Respect user preferences for motion

### Enhanced Accessibility

```html
<div class="slider" role="region" aria-label="Image carousel" aria-live="polite">
  <div class="slides" role="list">
    <div id="slide-1" class="slide" role="listitem" aria-label="Slide 1 of 3">
      <img src="/images/photo1.jpg" alt="Beautiful mountain landscape with snow-capped peaks">
    </div>
    <div id="slide-2" class="slide" role="listitem" aria-label="Slide 2 of 3">
      <img src="/images/photo2.jpg" alt="Bustling city skyline at sunset">
    </div>
    <div id="slide-3" class="slide" role="listitem" aria-label="Slide 3 of 3">
      <img src="/images/photo3.jpg" alt="Peaceful ocean view with clear blue water">
    </div>
  </div>
  
  <nav aria-label="Carousel navigation" class="flex justify-center gap-2 mt-4">
    <a href="#slide-1" aria-label="Go to slide 1" class="carousel-nav-dot"></a>
    <a href="#slide-2" aria-label="Go to slide 2" class="carousel-nav-dot"></a>
    <a href="#slide-3" aria-label="Go to slide 3" class="carousel-nav-dot"></a>
  </nav>
</div>
```

## JavaScript Enhancement

### Auto-Play Carousel

```javascript
class AutoCarousel {
  constructor(container, options = {}) {
    this.container = container;
    this.slides = container.querySelectorAll('.slides > div');
    this.currentSlide = 0;
    this.interval = options.interval || 5000;
    this.autoPlayId = null;
    
    this.init();
  }
  
  init() {
    this.startAutoPlay();
    this.addEventListeners();
  }
  
  startAutoPlay() {
    this.autoPlayId = setInterval(() => {
      this.nextSlide();
    }, this.interval);
  }
  
  stopAutoPlay() {
    if (this.autoPlayId) {
      clearInterval(this.autoPlayId);
      this.autoPlayId = null;
    }
  }
  
  nextSlide() {
    this.currentSlide = (this.currentSlide + 1) % this.slides.length;
    this.goToSlide(this.currentSlide);
  }
  
  goToSlide(index) {
    const slide = this.slides[index];
    if (slide) {
      slide.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
      this.currentSlide = index;
    }
  }
  
  addEventListeners() {
    // Pause on hover
    this.container.addEventListener('mouseenter', () => this.stopAutoPlay());
    this.container.addEventListener('mouseleave', () => this.startAutoPlay());
    
    // Pause on focus
    this.container.addEventListener('focusin', () => this.stopAutoPlay());
    this.container.addEventListener('focusout', () => this.startAutoPlay());
  }
}

// Usage
const carousel = new AutoCarousel(document.querySelector('.slider'), {
  interval: 3000
});
```

### Touch/Swipe Support

```javascript
class SwipeCarousel {
  constructor(container) {
    this.container = container;
    this.slides = container.querySelector('.slides');
    this.startX = 0;
    this.currentX = 0;
    this.isDragging = false;
    
    this.init();
  }
  
  init() {
    this.slides.addEventListener('touchstart', (e) => this.handleStart(e));
    this.slides.addEventListener('touchmove', (e) => this.handleMove(e));
    this.slides.addEventListener('touchend', () => this.handleEnd());
    
    // Mouse events for desktop
    this.slides.addEventListener('mousedown', (e) => this.handleStart(e));
    this.slides.addEventListener('mousemove', (e) => this.handleMove(e));
    this.slides.addEventListener('mouseup', () => this.handleEnd());
    this.slides.addEventListener('mouseleave', () => this.handleEnd());
  }
  
  handleStart(e) {
    this.isDragging = true;
    this.startX = e.type === 'touchstart' ? e.touches[0].clientX : e.clientX;
    this.slides.style.scrollBehavior = 'auto';
  }
  
  handleMove(e) {
    if (!this.isDragging) return;
    
    e.preventDefault();
    this.currentX = e.type === 'touchmove' ? e.touches[0].clientX : e.clientX;
    const diffX = this.startX - this.currentX;
    this.slides.scrollLeft += diffX;
    this.startX = this.currentX;
  }
  
  handleEnd() {
    this.isDragging = false;
    this.slides.style.scrollBehavior = 'smooth';
  }
}

// Usage
new SwipeCarousel(document.querySelector('.slider'));
```

### React Carousel Component

```jsx
import React, { useState, useEffect, useRef } from 'react';

function Carousel({ children, autoPlay = false, interval = 5000 }) {
  const [currentSlide, setCurrentSlide] = useState(0);
  const slidesRef = useRef(null);
  const autoPlayRef = useRef(null);
  
  const totalSlides = React.Children.count(children);
  
  const goToSlide = (index) => {
    setCurrentSlide(index);
    const slide = slidesRef.current?.children[index];
    if (slide) {
      slide.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
    }
  };
  
  const nextSlide = () => {
    goToSlide((currentSlide + 1) % totalSlides);
  };
  
  const prevSlide = () => {
    goToSlide(currentSlide === 0 ? totalSlides - 1 : currentSlide - 1);
  };
  
  useEffect(() => {
    if (autoPlay) {
      autoPlayRef.current = setInterval(nextSlide, interval);
      return () => clearInterval(autoPlayRef.current);
    }
  }, [autoPlay, interval, currentSlide]);
  
  const pauseAutoPlay = () => {
    if (autoPlayRef.current) {
      clearInterval(autoPlayRef.current);
    }
  };
  
  const resumeAutoPlay = () => {
    if (autoPlay) {
      autoPlayRef.current = setInterval(nextSlide, interval);
    }
  };
  
  return (
    <div 
      className="slider"
      onMouseEnter={pauseAutoPlay}
      onMouseLeave={resumeAutoPlay}
    >
      <div 
        ref={slidesRef}
        className="slides"
      >
        {React.Children.map(children, (child, index) => (
          <div key={index} id={`slide-${index}`} className="slide">
            {child}
          </div>
        ))}
      </div>
      
      {/* Navigation dots */}
      <div className="flex justify-center gap-2 mt-4">
        {Array.from({ length: totalSlides }, (_, index) => (
          <button
            key={index}
            onClick={() => goToSlide(index)}
            className={`w-3 h-3 rounded-full transition-colors ${
              index === currentSlide ? 'bg-primary' : 'bg-muted hover:bg-primary/50'
            }`}
            aria-label={`Go to slide ${index + 1}`}
          />
        ))}
      </div>
      
      {/* Arrow navigation */}
      <div className="flex justify-between items-center mt-4">
        <button
          onClick={prevSlide}
          className="btn-icon-outline"
          aria-label="Previous slide"
        >
          ←
        </button>
        <button
          onClick={nextSlide}
          className="btn-icon-outline"
          aria-label="Next slide"
        >
          →
        </button>
      </div>
    </div>
  );
}

// Usage
function App() {
  return (
    <Carousel autoPlay interval={4000}>
      <div>Slide 1 Content</div>
      <div>Slide 2 Content</div>
      <div>Slide 3 Content</div>
    </Carousel>
  );
}
```

## Best Practices

1. **Performance**: Use CSS scroll-snap for smooth native scrolling
2. **Accessibility**: Provide keyboard navigation and screen reader support
3. **Responsive Design**: Adapt slide sizes for different screen sizes
4. **Auto-play**: Include pause/play controls for auto-playing carousels
5. **Touch Support**: Enable swipe gestures on touch devices
6. **Loading States**: Show placeholders for content that's loading
7. **Navigation**: Provide multiple ways to navigate (dots, arrows, keyboard)

## Common Patterns

### Hero Carousel

```html
<div class="slider w-full h-96">
  <div class="slides">
    <div id="hero-1" class="slide bg-gradient-to-r from-blue-500 to-purple-600 text-white">
      <div class="flex items-center justify-center h-full">
        <div class="text-center">
          <h1 class="text-4xl font-bold mb-4">Welcome to Our Platform</h1>
          <p class="text-xl mb-6">Discover amazing features and capabilities</p>
          <button class="btn bg-white text-blue-600 hover:bg-gray-100">Get Started</button>
        </div>
      </div>
    </div>
  </div>
</div>
```

### Media Carousel with Thumbnails

```html
<div class="slider">
  <!-- Main carousel -->
  <div class="slides mb-4">
    <div id="main-1" class="slide">
      <img src="/media/image1.jpg" alt="Product image 1" class="w-full h-full object-cover">
    </div>
    <div id="main-2" class="slide">
      <img src="/media/image2.jpg" alt="Product image 2" class="w-full h-full object-cover">
    </div>
  </div>
  
  <!-- Thumbnail navigation -->
  <div class="flex gap-2 justify-center">
    <a href="#main-1" class="w-16 h-16 rounded border-2 border-transparent hover:border-primary">
      <img src="/media/thumb1.jpg" alt="Thumbnail 1" class="w-full h-full object-cover rounded">
    </a>
    <a href="#main-2" class="w-16 h-16 rounded border-2 border-transparent hover:border-primary">
      <img src="/media/thumb2.jpg" alt="Thumbnail 2" class="w-full h-full object-cover rounded">
    </a>
  </div>
</div>
```

## Related Components

- [Card](./card.md) - For carousel slide content
- [Button](./button.md) - For navigation controls
- [Image](./avatar.md) - For image carousels
- [Badge](./badge.md) - For slide indicators
---
## chart

# Chart Component

Interactive charts and data visualizations built with Chart.js and integrated with Basecoat theming.

## Basic Usage

```html
<div class="chart-container">
  <canvas id="myChart" width="400" height="200"></canvas>
</div>

<script>
  const ctx = document.getElementById('myChart').getContext('2d');
  new Chart(ctx, {
    type: 'line',
    data: {
      labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May'],
      datasets: [{
        label: 'Sales',
        data: [12, 19, 3, 5, 2],
        borderColor: 'hsl(var(--chart-1))',
        backgroundColor: 'hsla(var(--chart-1), 0.1)',
        tension: 0.3
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false
    }
  });
</script>
```

## CSS Classes

### Container Classes
- **`chart-container`** - Wrapper for chart canvas with proper sizing
- **`card`** - Often wrapped in card component for consistent styling

### Chart Variables
Basecoat provides CSS custom properties for consistent chart colors:
- **`--chart-1`** - Primary chart color (blue)
- **`--chart-2`** - Secondary chart color (emerald)  
- **`--chart-3`** - Tertiary chart color (yellow)
- **`--chart-4`** - Quaternary chart color (red)
- **`--chart-5`** - Quinary chart color (purple)

### Theme Integration
Charts automatically adapt to light/dark mode using CSS variables:
- **`--background`** - Chart background color
- **`--border`** - Grid line and border color
- **`--foreground`** - Text and label color
- **`--muted-foreground`** - Secondary text color

## Component Attributes

### Canvas Element
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `id` | string | Unique identifier for Chart.js initialization | Yes |
| `width` | number | Canvas width in pixels | Optional |
| `height` | number | Canvas height in pixels | Optional |

## JavaScript Required

This component requires Chart.js library:

```html
<script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
```

## HTML Structure

```html
<!-- Basic chart structure -->
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Chart Title</h3>
    <p class="text-sm text-muted-foreground">Chart description</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="chart-id"></canvas>
  </div>
</div>
```

## Examples

### Area Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Sales Overview</h3>
    <p class="text-sm text-muted-foreground">Monthly sales data</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="areaChart"></canvas>
  </div>
</div>

<script>
const areaCtx = document.getElementById('areaChart').getContext('2d');
new Chart(areaCtx, {
  type: 'line',
  data: {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun'],
    datasets: [{
      label: 'Sales',
      data: [30, 40, 45, 50, 49, 60],
      borderColor: 'hsl(var(--chart-1))',
      backgroundColor: createGradient(areaCtx, 'var(--chart-1)'),
      fill: true,
      tension: 0.4
    }]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});

function createGradient(ctx, color) {
  const gradient = ctx.createLinearGradient(0, 0, 0, 200);
  gradient.addColorStop(0, `hsla(${color}, 0.3)`);
  gradient.addColorStop(1, `hsla(${color}, 0)`);
  return gradient;
}
</script>
```

### Bar Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Revenue by Product</h3>
    <p class="text-sm text-muted-foreground">Quarterly comparison</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="barChart"></canvas>
  </div>
</div>

<script>
const barCtx = document.getElementById('barChart').getContext('2d');
new Chart(barCtx, {
  type: 'bar',
  data: {
    labels: ['Q1', 'Q2', 'Q3', 'Q4'],
    datasets: [
      {
        label: 'Product A',
        data: [65, 59, 80, 81],
        backgroundColor: 'hsl(var(--chart-1))',
        borderColor: 'hsl(var(--chart-1))',
        borderWidth: 1
      },
      {
        label: 'Product B',
        data: [28, 48, 40, 19],
        backgroundColor: 'hsl(var(--chart-2))',
        borderColor: 'hsl(var(--chart-2))',
        borderWidth: 1
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});
</script>
```

### Doughnut Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Market Share</h3>
    <p class="text-sm text-muted-foreground">By company segment</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="doughnutChart"></canvas>
  </div>
</div>

<script>
const doughnutCtx = document.getElementById('doughnutChart').getContext('2d');
new Chart(doughnutCtx, {
  type: 'doughnut',
  data: {
    labels: ['Desktop', 'Mobile', 'Tablet'],
    datasets: [{
      data: [300, 50, 100],
      backgroundColor: [
        'hsl(var(--chart-1))',
        'hsl(var(--chart-2))',
        'hsl(var(--chart-3))'
      ],
      borderColor: [
        'hsl(var(--chart-1))',
        'hsl(var(--chart-2))',
        'hsl(var(--chart-3))'
      ],
      borderWidth: 2
    }]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    cutout: '60%'
  }
});
</script>
```

### Line Chart with Multiple Datasets

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Website Analytics</h3>
    <p class="text-sm text-muted-foreground">Visitors and page views</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="lineChart"></canvas>
  </div>
</div>

<script>
const lineCtx = document.getElementById('lineChart').getContext('2d');
new Chart(lineCtx, {
  type: 'line',
  data: {
    labels: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
    datasets: [
      {
        label: 'Visitors',
        data: [1200, 1900, 3000, 5000, 2000, 3000, 4500],
        borderColor: 'hsl(var(--chart-1))',
        backgroundColor: 'hsla(var(--chart-1), 0.1)',
        tension: 0.3
      },
      {
        label: 'Page Views',
        data: [2400, 3800, 6000, 10000, 4000, 6000, 9000],
        borderColor: 'hsl(var(--chart-2))',
        backgroundColor: 'hsla(var(--chart-2), 0.1)',
        tension: 0.3
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    interaction: {
      intersect: false,
      mode: 'index'
    },
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      y: {
        beginAtZero: true
      }
    }
  }
});
</script>
```

### Stacked Bar Chart

```html
<div class="card p-6">
  <header class="mb-6">
    <h3 class="text-lg font-semibold">Revenue Breakdown</h3>
    <p class="text-sm text-muted-foreground">By region and quarter</p>
  </header>
  
  <div class="chart-container h-64">
    <canvas id="stackedChart"></canvas>
  </div>
</div>

<script>
const stackedCtx = document.getElementById('stackedChart').getContext('2d');
new Chart(stackedCtx, {
  type: 'bar',
  data: {
    labels: ['Q1', 'Q2', 'Q3', 'Q4'],
    datasets: [
      {
        label: 'North America',
        data: [120, 150, 180, 200],
        backgroundColor: 'hsl(var(--chart-1))',
      },
      {
        label: 'Europe',
        data: [80, 90, 100, 110],
        backgroundColor: 'hsl(var(--chart-2))',
      },
      {
        label: 'Asia Pacific',
        data: [60, 70, 85, 95],
        backgroundColor: 'hsl(var(--chart-3))',
      }
    ]
  },
  options: {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'bottom'
      }
    },
    scales: {
      x: {
        stacked: true
      },
      y: {
        stacked: true,
        beginAtZero: true
      }
    }
  }
});
</script>
```

## Custom Tooltip Implementation

```javascript
// Custom tooltip for better styling integration
const customTooltip = (context) => {
  // Get or create tooltip element
  let tooltipEl = document.getElementById('chartjs-tooltip');
  
  if (!tooltipEl) {
    tooltipEl = document.createElement('div');
    tooltipEl.id = 'chartjs-tooltip';
    tooltipEl.className = 'absolute bg-background border border-border rounded-lg shadow-lg p-3 text-sm pointer-events-none z-50 opacity-0 transition-opacity';
    document.body.appendChild(tooltipEl);
  }

  // Hide if no tooltip
  const tooltipModel = context.tooltip;
  if (tooltipModel.opacity === 0) {
    tooltipEl.style.opacity = 0;
    return;
  }

  // Set content
  if (tooltipModel.body) {
    const titleLines = tooltipModel.title || [];
    const bodyLines = tooltipModel.body.map(item => item.lines);

    let innerHtml = '<div class="font-medium mb-1">';
    titleLines.forEach(title => {
      innerHtml += title;
    });
    innerHtml += '</div>';

    bodyLines.forEach((body, i) => {
      const colors = tooltipModel.labelColors[i];
      innerHtml += `
        <div class="flex items-center gap-2">
          <div class="w-3 h-3 rounded-full" style="background-color: ${colors.backgroundColor}"></div>
          <span>${body}</span>
        </div>
      `;
    });

    tooltipEl.innerHTML = innerHtml;
  }

  // Position
  const position = context.chart.canvas.getBoundingClientRect();
  tooltipEl.style.opacity = 1;
  tooltipEl.style.left = position.left + window.pageXOffset + tooltipModel.caretX + 'px';
  tooltipEl.style.top = position.top + window.pageYOffset + tooltipModel.caretY + 'px';
};

// Usage in chart configuration
const chartOptions = {
  plugins: {
    tooltip: {
      enabled: false,
      external: customTooltip
    }
  }
};
```

## Helper Functions

```javascript
// Color utilities for charts
const chartHelpers = {
  // Convert chart color variables to usable format
  getChartColor(variable) {
    return `hsl(${getComputedStyle(document.documentElement).getPropertyValue(variable)})`;
  },

  // Add alpha channel to chart colors
  setAlpha(color, alpha) {
    const hslMatch = color.match(/hsl\((.*?)\)/);
    if (hslMatch) {
      return `hsla(${hslMatch[1]}, ${alpha})`;
    }
    return color;
  },

  // Create gradient for area charts
  createGradient(ctx, colorVar, height = 200) {
    const gradient = ctx.createLinearGradient(0, 0, 0, height);
    const color = this.getChartColor(colorVar);
    
    gradient.addColorStop(0, this.setAlpha(color, 0.3));
    gradient.addColorStop(1, this.setAlpha(color, 0));
    
    return gradient;
  },

  // Format dates for chart labels
  formatDateLabel(dateString) {
    const date = new Date(dateString);
    return date.toLocaleDateString('en-US', { 
      month: 'short', 
      day: 'numeric' 
    });
  },

  // Generate chart data from API response
  transformData(apiData, labelKey, valueKey) {
    return {
      labels: apiData.map(item => this.formatDateLabel(item[labelKey])),
      datasets: [{
        data: apiData.map(item => item[valueKey]),
        borderColor: this.getChartColor('--chart-1'),
        backgroundColor: this.createGradient(null, '--chart-1')
      }]
    };
  }
};
```

## Responsive Chart Component

```javascript
class ResponsiveChart {
  constructor(canvasId, config) {
    this.canvas = document.getElementById(canvasId);
    this.ctx = this.canvas.getContext('2d');
    this.config = config;
    this.chart = null;
    
    this.init();
    this.setupResize();
  }
  
  init() {
    this.chart = new Chart(this.ctx, {
      ...this.config,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        ...this.config.options
      }
    });
  }
  
  setupResize() {
    window.addEventListener('resize', () => {
      if (this.chart) {
        this.chart.resize();
      }
    });
  }
  
  updateData(newData) {
    this.chart.data = newData;
    this.chart.update();
  }
  
  destroy() {
    if (this.chart) {
      this.chart.destroy();
    }
  }
}

// Usage
const salesChart = new ResponsiveChart('salesChart', {
  type: 'line',
  data: {
    labels: ['Jan', 'Feb', 'Mar'],
    datasets: [{
      label: 'Sales',
      data: [100, 200, 150]
    }]
  }
});
```

## React Chart Component

```jsx
import React, { useEffect, useRef } from 'react';
import Chart from 'chart.js/auto';

function ChartComponent({ 
  type = 'line', 
  data, 
  options = {}, 
  className = '',
  height = 256 
}) {
  const canvasRef = useRef(null);
  const chartRef = useRef(null);
  
  useEffect(() => {
    if (canvasRef.current) {
      // Destroy existing chart
      if (chartRef.current) {
        chartRef.current.destroy();
      }
      
      // Create new chart
      chartRef.current = new Chart(canvasRef.current, {
        type,
        data,
        options: {
          responsive: true,
          maintainAspectRatio: false,
          plugins: {
            tooltip: {
              backgroundColor: 'hsl(var(--background))',
              titleColor: 'hsl(var(--foreground))',
              bodyColor: 'hsl(var(--foreground))',
              borderColor: 'hsl(var(--border))',
              borderWidth: 1
            }
          },
          ...options
        }
      });
    }
    
    return () => {
      if (chartRef.current) {
        chartRef.current.destroy();
      }
    };
  }, [type, data, options]);
  
  return (
    <div className={`chart-container ${className}`} style={{ height }}>
      <canvas ref={canvasRef} />
    </div>
  );
}

// Usage
function Dashboard() {
  const chartData = {
    labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May'],
    datasets: [{
      label: 'Revenue',
      data: [12, 19, 3, 5, 2],
      borderColor: 'hsl(var(--chart-1))',
      backgroundColor: 'hsla(var(--chart-1), 0.1)',
    }]
  };
  
  return (
    <div className="card p-6">
      <header className="mb-6">
        <h3 className="text-lg font-semibold">Sales Dashboard</h3>
        <p className="text-sm text-muted-foreground">Monthly revenue tracking</p>
      </header>
      
      <ChartComponent
        type="line"
        data={chartData}
        height={300}
        options={{
          scales: {
            y: {
              beginAtZero: true
            }
          }
        }}
      />
    </div>
  );
}
```

## Accessibility Features

- **Keyboard Navigation**: Charts support keyboard interaction when focused
- **Screen Reader Support**: Include descriptive labels and summaries
- **High Contrast**: Colors are accessible in both light and dark modes
- **Alternative Text**: Provide text alternatives for chart data

### Enhanced Accessibility

```html
<div class="card p-6">
  <header className="mb-6">
    <h3 className="text-lg font-semibold" id="sales-chart-title">
      Sales Performance
    </h3>
    <p className="text-sm text-muted-foreground">
      Monthly sales data from January to June 2024
    </p>
  </header>
  
  <div className="chart-container h-64">
    <canvas 
      id="salesChart"
      role="img"
      aria-labelledby="sales-chart-title"
      aria-describedby="sales-chart-summary"
    ></canvas>
  </div>
  
  <div id="sales-chart-summary" className="sr-only">
    Chart showing sales performance over 6 months. 
    January: $30k, February: $40k, March: $45k, 
    April: $50k, May: $49k, June: $60k. 
    Overall trend is positive with steady growth.
  </div>
  
  <!-- Data table alternative -->
  <details className="mt-4">
    <summary className="text-sm text-muted-foreground cursor-pointer">
      View data table
    </summary>
    <table className="table mt-2">
      <thead>
        <tr>
          <th>Month</th>
          <th>Sales</th>
        </tr>
      </thead>
      <tbody>
        <tr><td>January</td><td>$30,000</td></tr>
        <tr><td>February</td><td>$40,000</td></tr>
        <tr><td>March</td><td>$45,000</td></tr>
        <tr><td>April</td><td>$50,000</td></tr>
        <tr><td>May</td><td>$49,000</td></tr>
        <tr><td>June</td><td>$60,000</td></tr>
      </tbody>
    </table>
  </details>
</div>
```

## Best Practices

1. **Performance**: Use Chart.js animations sparingly on mobile devices
2. **Accessibility**: Always provide data tables as alternatives
3. **Responsive Design**: Test charts on different screen sizes
4. **Color Choice**: Use the provided chart color variables
5. **Loading States**: Show skeleton loaders while data loads
6. **Error Handling**: Display fallback content when charts fail to load
7. **Data Updates**: Use Chart.js update methods for smooth transitions

## Common Patterns

### Dashboard Grid

```html
<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
  <!-- KPI Cards -->
  <div className="card p-6">
    <div className="flex items-center justify-between">
      <div>
        <p className="text-sm text-muted-foreground">Total Revenue</p>
        <p className="text-2xl font-bold">$45,231</p>
      </div>
      <div className="chart-container h-16 w-16">
        <canvas id="revenueSparkline"></canvas>
      </div>
    </div>
  </div>
  
  <!-- Main Chart -->
  <div className="card p-6 md:col-span-2">
    <div className="chart-container h-64">
      <canvas id="mainChart"></canvas>
    </div>
  </div>
</div>
```

### Chart with Controls

```html
<div className="card p-6">
  <header className="flex items-center justify-between mb-6">
    <div>
      <h3 className="text-lg font-semibold">Analytics</h3>
      <p className="text-sm text-muted-foreground">User engagement metrics</p>
    </div>
    <div className="flex gap-2">
      <select className="select" id="timeRange">
        <option value="7d">Last 7 days</option>
        <option value="30d">Last 30 days</option>
        <option value="90d">Last 90 days</option>
      </select>
    </div>
  </header>
  
  <div className="chart-container h-64">
    <canvas id="analyticsChart"></canvas>
  </div>
</div>
```

## Related Components

- [Card](./card.md) - For chart containers
- [Select](./select.md) - For chart controls
- [Badge](./badge.md) - For chart legends
- [Table](./table.md) - For data table alternatives
---
## empty

# Empty State Component

Display empty states with icons, titles, descriptions, and actions to guide users when content is unavailable.

## Basic Usage

```html
<div class="empty-state">
  <div class="flex items-center justify-center size-12 mx-auto mb-4 bg-muted rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
      <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/>
      <polyline points="14,2 14,8 20,8"/>
    </svg>
  </div>
  <h3 class="text-lg font-semibold mb-2">No Documents</h3>
  <p class="text-muted-foreground mb-6">You haven't created any documents yet. Get started by creating your first document.</p>
  <button class="btn">Create Document</button>
</div>
```

## CSS Classes

### Container Classes
- **`empty-state`** - Base styling for empty state container
- **`flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12`** - Complete layout classes

### Content Classes
- **`text-lg font-semibold`** - Title styling
- **`text-muted-foreground`** - Description text color
- **`size-12 bg-muted rounded-lg`** - Icon container styling

### Layout Classes
- **`flex items-center justify-center`** - Icon centering
- **`mx-auto mb-4`** - Icon positioning
- **`mb-2`**, `mb-6` - Spacing between elements

### Tailwind Utilities
- **`text-center`** - Center align text
- **`text-balance`** - Better text wrapping
- **`gap-6`** - Consistent spacing between sections
- **`md:p-12`** - Responsive padding

## Component Attributes

### No Specific Attributes Required
Empty states are pure HTML compositions using Tailwind utilities.

### Icon Container
| Element | Classes | Description |
|---------|---------|-------------|
| Icon container | `flex items-center justify-center size-12 bg-muted rounded-lg` | Standard icon background |
| Icon | `text-muted-foreground` | Muted icon color |

## No JavaScript Required
Empty states are static content displays requiring no JavaScript.

## HTML Structure

```html
<!-- Basic empty state structure -->
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <!-- Icon/Image -->
  <div class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
    <!-- SVG icon -->
  </div>
  
  <!-- Content -->
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">Title</h3>
    <p class="text-muted-foreground">Description text explaining the empty state</p>
  </div>
  
  <!-- Actions -->
  <div class="flex flex-col sm:flex-row gap-3">
    <button class="btn">Primary Action</button>
    <button class="btn-outline">Secondary Action</button>
  </div>
  
  <!-- Optional help link -->
  <a href="/help" class="text-sm text-muted-foreground hover:text-foreground">
    Need help? Contact support
  </a>
</div>
```

## Examples

### No Projects State

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
      <path d="M2 20h20"/>
      <path d="M4 20V10a2 2 0 0 1 2-2h2.586a1 1 0 0 0 .707-.293l1.414-1.414A1 1 0 0 1 11.414 6H18a2 2 0 0 1 2 2v12"/>
      <path d="M7 14h3m-3 3h5"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">No Projects Yet</h3>
    <p class="text-muted-foreground">
      You haven't created any projects yet. Get started by creating your first project.
    </p>
  </div>
  
  <div class="flex flex-col sm:flex-row gap-3">
    <button class="btn">Create Project</button>
    <button class="btn-outline">Import Project</button>
  </div>
  
  <a href="/docs" class="text-sm text-muted-foreground hover:text-foreground inline-flex items-center gap-1">
    Learn More
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6"/>
      <polyline points="15,3 21,3 21,9"/>
      <line x1="10" x2="21" y1="14" y2="3"/>
    </svg>
  </a>
</div>
```

### Cloud Storage Empty

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg border p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-16 mx-auto bg-muted rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
      <path d="M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">Cloud Storage Empty</h3>
    <p class="text-muted-foreground">
      Upload files to your cloud storage to access them anywhere.
    </p>
  </div>
  
  <button class="btn-sm-outline">Upload Files</button>
</div>
```

### User Offline State

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <img 
    src="/avatars/user.jpg" 
    alt="John Doe" 
    class="size-16 rounded-full mx-auto grayscale"
  >
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">User Offline</h3>
    <p class="text-muted-foreground">
      This user is currently offline. You can leave a message to notify them or try again later.
    </p>
  </div>
  
  <button class="btn">Leave Message</button>
</div>
```

### No Team Members

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="flex -space-x-2 mx-auto">
    <img src="/avatars/user1.jpg" alt="User 1" class="size-12 rounded-full border-2 border-background grayscale">
    <img src="/avatars/user2.jpg" alt="User 2" class="size-12 rounded-full border-2 border-background grayscale">
    <div class="size-12 rounded-full border-2 border-background bg-muted flex items-center justify-center">
      <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
        <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/>
        <circle cx="9" cy="7" r="4"/>
        <path d="M22 21v-2a4 4 0 0 0-3-3.87"/>
        <path d="M16 3.13a4 4 0 0 1 0 7.75"/>
      </svg>
    </div>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">No Team Members</h3>
    <p class="text-muted-foreground">
      Invite your team to collaborate on this project.
    </p>
  </div>
  
  <button class="btn inline-flex items-center gap-2">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <path d="M5 12h14"/>
      <path d="M12 5v14"/>
    </svg>
    Invite Members
  </button>
</div>
```

### 404 Page Not Found

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="space-y-2">
    <h3 class="text-2xl font-bold">404 - Not Found</h3>
    <p class="text-muted-foreground">
      The page you're looking for doesn't exist. Try searching for what you need below.
    </p>
  </div>
  
  <div class="w-full max-w-sm relative">
    <div class="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
        <circle cx="11" cy="11" r="8"/>
        <path d="m21 21-4.3-4.3"/>
      </svg>
    </div>
    <input 
      type="search" 
      placeholder="Search..."
      class="input w-full pl-10 pr-20"
    >
    <div class="absolute right-3 top-1/2 -translate-y-1/2 text-muted-foreground text-sm">
      ⌘K
    </div>
  </div>
  
  <a href="/support" class="text-sm text-muted-foreground hover:text-foreground">
    Need help? Contact support
  </a>
</div>
```

### No Search Results

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
      <circle cx="11" cy="11" r="8"/>
      <path d="m21 21-4.3-4.3"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">No Results Found</h3>
    <p class="text-muted-foreground">
      We couldn't find anything matching "<span class="font-medium">your search</span>". 
      Try adjusting your search terms.
    </p>
  </div>
  
  <div class="flex flex-col sm:flex-row gap-3">
    <button class="btn-outline">Clear Search</button>
    <button class="btn-ghost">Browse All Items</button>
  </div>
</div>
```

### No Notifications

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
      <path d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9"/>
      <path d="M10.3 21a1.94 1.94 0 0 0 3.4 0"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">You're All Caught Up!</h3>
    <p class="text-muted-foreground">
      No new notifications. We'll notify you when something important happens.
    </p>
  </div>
  
  <button class="btn-outline">Notification Settings</button>
</div>
```

### Connection Error

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg border-dashed border-2 p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-12 mx-auto bg-destructive/10 rounded-lg">
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-destructive">
      <path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09z"/>
      <path d="M12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"/>
      <path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0"/>
      <path d="M12 15v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">Connection Failed</h3>
    <p class="text-muted-foreground">
      Unable to connect to the server. Please check your internet connection and try again.
    </p>
  </div>
  
  <div class="flex flex-col sm:flex-row gap-3">
    <button class="btn">Try Again</button>
    <button class="btn-ghost">Work Offline</button>
  </div>
</div>
```

### Loading State Empty

```html
<div class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12">
  <div class="flex items-center justify-center size-12 mx-auto">
    <svg class="animate-spin h-6 w-6 text-muted-foreground" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
      <path class="opacity-75" fill="currentColor" d="m12 2 0 4c-4.418 0-8 3.582-8 8 0 1.1.224 2.148.63 3.1L2.369 18.9C1.502 17.065 1 15.087 1 13c0-6.075 4.925-11 11-11Z"></path>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 class="text-lg font-semibold">Loading...</h3>
    <p class="text-muted-foreground">
      Please wait while we fetch your data.
    </p>
  </div>
</div>
```

## Accessibility Features

- **Semantic HTML**: Use proper headings and structure
- **Alt Text**: Provide descriptive alt text for images
- **Focus Management**: Ensure interactive elements are keyboard accessible
- **Screen Reader Support**: Use proper labels and descriptions

### Enhanced Accessibility

```html
<div 
  class="flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12"
  role="region"
  aria-labelledby="empty-title"
  aria-describedby="empty-description"
>
  <div class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
    <svg 
      xmlns="http://www.w3.org/2000/svg" 
      width="24" 
      height="24" 
      viewBox="0 0 24 24" 
      fill="none" 
      stroke="currentColor" 
      stroke-width="2" 
      stroke-linecap="round" 
      stroke-linejoin="round" 
      class="text-muted-foreground"
      aria-hidden="true"
    >
      <path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/>
      <polyline points="14,2 14,8 20,8"/>
    </svg>
  </div>
  
  <div class="space-y-2">
    <h3 id="empty-title" class="text-lg font-semibold">No Documents</h3>
    <p id="empty-description" class="text-muted-foreground">
      You haven't created any documents yet. Get started by creating your first document.
    </p>
  </div>
  
  <button class="btn" aria-describedby="empty-description">
    Create Document
  </button>
</div>
```

## React Integration

```jsx
import React from 'react';

function EmptyState({
  icon,
  title,
  description,
  actions,
  className = '',
  variant = 'default' // 'default' | 'bordered' | 'error'
}) {
  const containerClasses = [
    'flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12',
    variant === 'bordered' && 'border',
    variant === 'error' && 'border-dashed border-2',
    className
  ].filter(Boolean).join(' ');

  return (
    <div className={containerClasses}>
      {icon && (
        <div className="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
          {icon}
        </div>
      )}
      
      <div className="space-y-2">
        <h3 className="text-lg font-semibold">{title}</h3>
        {description && (
          <p className="text-muted-foreground">{description}</p>
        )}
      </div>
      
      {actions && (
        <div className="flex flex-col sm:flex-row gap-3">
          {actions.map((action, index) => action)}
        </div>
      )}
    </div>
  );
}

// Usage
function ProjectsPage() {
  const projectIcon = (
    <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" className="text-muted-foreground">
      <path d="M2 20h20"/>
      <path d="M4 20V10a2 2 0 0 1 2-2h2.586a1 1 0 0 0 .707-.293l1.414-1.414A1 1 0 0 1 11.414 6H18a2 2 0 0 1 2 2v12"/>
    </svg>
  );

  return (
    <EmptyState
      icon={projectIcon}
      title="No Projects Yet"
      description="You haven't created any projects yet. Get started by creating your first project."
      actions={[
        <button key="create" className="btn">Create Project</button>,
        <button key="import" className="btn-outline">Import Project</button>
      ]}
    />
  );
}
```

### Vue Integration

```vue
<template>
  <div :class="containerClasses">
    <div v-if="icon" class="flex items-center justify-center size-12 mx-auto bg-muted rounded-lg">
      <component :is="icon" class="text-muted-foreground" />
    </div>
    
    <div class="space-y-2">
      <h3 class="text-lg font-semibold">{{ title }}</h3>
      <p v-if="description" class="text-muted-foreground">{{ description }}</p>
    </div>
    
    <div v-if="$slots.actions" class="flex flex-col sm:flex-row gap-3">
      <slot name="actions" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    icon: Object,
    title: String,
    description: String,
    variant: {
      type: String,
      default: 'default',
      validator: value => ['default', 'bordered', 'error'].includes(value)
    }
  },
  computed: {
    containerClasses() {
      const base = 'flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg p-6 text-center text-balance md:p-12';
      const variants = {
        bordered: 'border',
        error: 'border-dashed border-2'
      };
      
      return [base, variants[this.variant]].filter(Boolean).join(' ');
    }
  }
};
</script>
```

## Best Practices

1. **Clear Messaging**: Use descriptive titles and helpful descriptions
2. **Actionable**: Provide relevant actions to help users proceed
3. **Visual Hierarchy**: Use consistent icon sizes and spacing
4. **Context Aware**: Tailor empty states to specific scenarios
5. **Progressive Disclosure**: Don't overwhelm with too many actions
6. **Accessible**: Include proper semantic markup and ARIA labels
7. **Consistent**: Use similar patterns across your application

## Common Patterns

### Data Table Empty State

```html
<tr>
  <td colspan="4" class="p-8">
    <div class="flex flex-col items-center gap-4 text-center">
      <div class="flex items-center justify-center size-10 bg-muted rounded-lg">
        <svg class="w-5 h-5 text-muted-foreground"><!-- table icon --></svg>
      </div>
      <div>
        <h4 class="font-medium">No entries found</h4>
        <p class="text-sm text-muted-foreground mt-1">Start by adding your first entry</p>
      </div>
      <button class="btn-sm">Add Entry</button>
    </div>
  </td>
</tr>
```

### Sidebar Empty State

```html
<div class="p-6 text-center">
  <div class="flex items-center justify-center size-8 mx-auto mb-3 bg-muted rounded">
    <svg class="w-4 h-4 text-muted-foreground"><!-- folder icon --></svg>
  </div>
  <p class="text-sm font-medium mb-1">No folders</p>
  <p class="text-xs text-muted-foreground mb-3">Create folders to organize your files</p>
  <button class="btn-sm-outline w-full">New Folder</button>
</div>
```

### Modal Empty State

```html
<div class="dialog-content text-center py-8">
  <div class="flex items-center justify-center size-16 mx-auto mb-4 bg-muted rounded-full">
    <svg class="w-8 h-8 text-muted-foreground"><!-- icon --></svg>
  </div>
  <h3 class="text-xl font-semibold mb-2">All Set!</h3>
  <p class="text-muted-foreground mb-6">You've completed all available tasks.</p>
  <button class="btn">Close</button>
</div>
```

## Related Components

- [Button](./button.md) - For action buttons
- [Card](./card.md) - For container styling
- [Spinner](./spinner.md) - For loading states
- [Avatar](./avatar.md) - For user-related empty states
---
## form

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
---
## kbd

# Kbd Component

Used to display textual user input from keyboard, including shortcuts and key combinations.

## Basic Usage

```html
<kbd class="kbd">K</kbd>
```

## CSS Classes

### Base Classes
- **`kbd`** - Base keyboard key styling with muted background and proper typography

### Size Variants
- **`kbd-sm`** - Smaller keyboard key (text-xs, reduced padding)
- **`kbd-lg`** - Larger keyboard key (increased padding)

### Style Variants
- **`kbd-outline`** - Outline style with border
- **`kbd-solid`** - Solid background style
- **`kbd-accent`** - Accent colored style

### State Classes
- **`kbd-disabled`** - Disabled state with reduced opacity
- **`kbd-active`** - Active/pressed state

### Tailwind Utilities Used
- **`inline-flex items-center justify-center`** - Layout and alignment
- **`min-h-5 h-5 px-1.5 py-0.5`** - Size and spacing
- **`text-xs font-mono font-medium`** - Typography
- **`bg-muted text-muted-foreground`** - Colors
- **`border border-border`** - Border styling
- **`rounded shadow-xs`** - Shape and shadow

## Component Attributes

### Standard HTML Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "kbd" base class | Yes |

### Data Attributes (Optional)
| Attribute | Values | Description | Required |
|-----------|--------|-------------|----------|
| `data-key` | "cmd", "shift", "option", "ctrl", "enter", "esc", "tab", "space", "arrow-up", "arrow-down", "arrow-left", "arrow-right" | Shows corresponding key symbol | No |

## No JavaScript Required
The kbd component is purely presentational and requires no JavaScript.

## HTML Structure

```html
<!-- Basic key -->
<kbd class="kbd">Key</kbd>

<!-- Key combination -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">Ctrl</kbd>
  <span class="text-muted-foreground">+</span>
  <kbd class="kbd">B</kbd>
</span>

<!-- With data attributes -->
<kbd class="kbd" data-key="cmd"></kbd>
```

## Examples

### Single Keys

```html
<!-- Letter keys -->
<kbd class="kbd">K</kbd>
<kbd class="kbd">A</kbd>
<kbd class="kbd">Z</kbd>

<!-- Number keys -->
<kbd class="kbd">1</kbd>
<kbd class="kbd">2</kbd>
<kbd class="kbd">0</kbd>

<!-- Special keys -->
<kbd class="kbd">Enter</kbd>
<kbd class="kbd">Esc</kbd>
<kbd class="kbd">Tab</kbd>
<kbd class="kbd">Space</kbd>
```

### Symbol Keys

```html
<!-- Direct symbols -->
<kbd class="kbd">⌘</kbd>
<kbd class="kbd">⇧</kbd>
<kbd class="kbd">⌥</kbd>
<kbd class="kbd">⌃</kbd>
<kbd class="kbd">⏎</kbd>

<!-- Using data attributes -->
<kbd class="kbd" data-key="cmd"></kbd>
<kbd class="kbd" data-key="shift"></kbd>
<kbd class="kbd" data-key="option"></kbd>
<kbd class="kbd" data-key="ctrl"></kbd>
<kbd class="kbd" data-key="enter"></kbd>
```

### Key Combinations

```html
<!-- Text combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">Ctrl</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">B</kbd>
</span>

<!-- Symbol combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <kbd class="kbd">K</kbd>
</span>

<!-- Complex combinations -->
<span class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</span>
```

### Size Variants

```html
<!-- Small -->
<kbd class="kbd kbd-sm">⌘</kbd>
<kbd class="kbd kbd-sm">K</kbd>

<!-- Default -->
<kbd class="kbd">Enter</kbd>

<!-- Large -->
<kbd class="kbd kbd-lg">Space</kbd>
```

### Style Variants

```html
<!-- Default (muted) -->
<kbd class="kbd">Ctrl</kbd>

<!-- Outline -->
<kbd class="kbd kbd-outline">Alt</kbd>

<!-- Solid -->
<kbd class="kbd kbd-solid">Shift</kbd>

<!-- Accent -->
<kbd class="kbd kbd-accent">⌘</kbd>
```

### Arrow Keys

```html
<div class="inline-flex items-center gap-1">
  <kbd class="kbd" data-key="arrow-up"></kbd>
  <kbd class="kbd" data-key="arrow-down"></kbd>
  <kbd class="kbd" data-key="arrow-left"></kbd>
  <kbd class="kbd" data-key="arrow-right"></kbd>
</div>

<!-- Or with direct symbols -->
<div class="inline-flex items-center gap-1">
  <kbd class="kbd">↑</kbd>
  <kbd class="kbd">↓</kbd>
  <kbd class="kbd">←</kbd>
  <kbd class="kbd">→</kbd>
</div>
```

### In Buttons

```html
<!-- Button with keyboard shortcut -->
<button class="btn-sm-outline inline-flex items-center gap-2">
  Save
  <kbd class="kbd kbd-sm">⌘S</kbd>
</button>

<button class="btn-ghost inline-flex items-center gap-2">
  Search
  <kbd class="kbd kbd-sm">⌘K</kbd>
</button>

<!-- Accept/Cancel actions -->
<div class="flex gap-2">
  <button class="btn-sm inline-flex items-center gap-2">
    Accept
    <kbd class="kbd kbd-sm">⏎</kbd>
  </button>
  <button class="btn-sm-outline inline-flex items-center gap-2">
    Cancel
    <kbd class="kbd kbd-sm">Esc</kbd>
  </button>
</div>
```

### In Tooltips

```html
<button 
  class="btn-icon-outline" 
  data-tooltip="Bold (⌘B)"
  data-side="bottom"
>
  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
    <path d="M6 4h8a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>
    <path d="M6 12h9a4 4 0 0 1 4 4 4 4 0 0 1-4 4H6z"/>
  </svg>
</button>

<!-- Or in tooltip content -->
<div class="tooltip-content">
  <div class="font-medium">Bold</div>
  <div class="text-xs text-muted-foreground mt-1">
    Press <kbd class="kbd kbd-sm">⌘</kbd> <kbd class="kbd kbd-sm">B</kbd>
  </div>
</div>
```

### Help Documentation

```html
<div class="space-y-4">
  <h3 class="font-semibold">Keyboard Shortcuts</h3>
  <dl class="space-y-2">
    <div class="flex items-center justify-between">
      <dt>Open command palette</dt>
      <dd><kbd class="kbd">⌘K</kbd></dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Save document</dt>
      <dd><kbd class="kbd">⌘S</kbd></dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Bold text</dt>
      <dd>
        <span class="inline-flex items-center gap-1">
          <kbd class="kbd">⌘</kbd>
          <kbd class="kbd">B</kbd>
        </span>
      </dd>
    </div>
    <div class="flex items-center justify-between">
      <dt>Find and replace</dt>
      <dd>
        <span class="inline-flex items-center gap-1">
          <kbd class="kbd">⌘</kbd>
          <span class="text-muted-foreground text-xs">+</span>
          <kbd class="kbd">⇧</kbd>
          <span class="text-muted-foreground text-xs">+</span>
          <kbd class="kbd">F</kbd>
        </span>
      </dd>
    </div>
  </dl>
</div>
```

### Command Menu Items

```html
<div class="command-menu">
  <div class="command-item flex items-center justify-between p-2 rounded hover:bg-muted">
    <div class="flex items-center gap-3">
      <svg class="w-4 h-4"><!-- icon --></svg>
      <span>Open file</span>
    </div>
    <kbd class="kbd kbd-sm">⌘O</kbd>
  </div>
  
  <div class="command-item flex items-center justify-between p-2 rounded hover:bg-muted">
    <div class="flex items-center gap-3">
      <svg class="w-4 h-4"><!-- icon --></svg>
      <span>Close tab</span>
    </div>
    <kbd class="kbd kbd-sm">⌘W</kbd>
  </div>
</div>
```

### State Demonstrations

```html
<!-- Disabled state -->
<kbd class="kbd kbd-disabled">Unavailable</kbd>

<!-- Active/pressed state -->
<kbd class="kbd kbd-active">Pressed</kbd>

<!-- In different contexts -->
<div class="space-y-2">
  <p>Press <kbd class="kbd">Space</kbd> to continue</p>
  <p class="text-sm text-muted-foreground">
    Use <kbd class="kbd kbd-sm">Tab</kbd> and <kbd class="kbd kbd-sm">⇧Tab</kbd> to navigate
  </p>
</div>
```

### Platform-Specific Keys

```html
<!-- Mac shortcuts -->
<div class="mac-shortcuts space-y-2">
  <div class="flex items-center justify-between">
    <span>Copy</span>
    <kbd class="kbd">⌘C</kbd>
  </div>
  <div class="flex items-center justify-between">
    <span>Paste</span>
    <kbd class="kbd">⌘V</kbd>
  </div>
</div>

<!-- Windows/Linux shortcuts -->
<div class="pc-shortcuts space-y-2">
  <div class="flex items-center justify-between">
    <span>Copy</span>
    <kbd class="kbd">Ctrl+C</kbd>
  </div>
  <div class="flex items-center justify-between">
    <span>Paste</span>
    <kbd class="kbd">Ctrl+V</kbd>
  </div>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper `<kbd>` element for keyboard input
- **Screen Reader Support**: Screen readers announce kbd content appropriately
- **High Contrast**: Sufficient contrast in both light and dark modes
- **Font Choice**: Monospace font for consistent character width

### Enhanced Accessibility

```html
<!-- With descriptive text -->
<p>
  To save your work, press 
  <kbd class="kbd" title="Command key">⌘</kbd>
  <kbd class="kbd" title="Letter S key">S</kbd>
</p>

<!-- With aria-label for complex combinations -->
<span 
  class="inline-flex items-center gap-1"
  aria-label="Command plus Shift plus P"
>
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</span>

<!-- In help content with proper structure -->
<section aria-labelledby="shortcuts-heading">
  <h3 id="shortcuts-heading">Available Shortcuts</h3>
  <dl>
    <dt>Save document</dt>
    <dd>
      <kbd class="kbd" aria-label="Command S">⌘S</kbd>
    </dd>
  </dl>
</section>
```

## JavaScript Integration

### Dynamic Key Display

```javascript
// Show platform-appropriate shortcuts
function getShortcutDisplay(action) {
  const isMac = navigator.platform.toUpperCase().indexOf('MAC') >= 0;
  
  const shortcuts = {
    save: isMac ? '⌘S' : 'Ctrl+S',
    copy: isMac ? '⌘C' : 'Ctrl+C',
    paste: isMac ? '⌘V' : 'Ctrl+V',
    find: isMac ? '⌘F' : 'Ctrl+F'
  };
  
  return shortcuts[action] || '';
}

// Usage
document.addEventListener('DOMContentLoaded', () => {
  document.querySelectorAll('[data-shortcut]').forEach(el => {
    const action = el.dataset.shortcut;
    const kbd = el.querySelector('.kbd');
    if (kbd) {
      kbd.textContent = getShortcutDisplay(action);
    }
  });
});
```

### Keyboard Event Handling

```javascript
// Listen for shortcuts and highlight corresponding kbd elements
document.addEventListener('keydown', (e) => {
  const keys = [];
  
  if (e.metaKey || e.ctrlKey) keys.push(navigator.platform.includes('Mac') ? '⌘' : 'Ctrl');
  if (e.shiftKey) keys.push('⇧');
  if (e.altKey) keys.push(navigator.platform.includes('Mac') ? '⌥' : 'Alt');
  
  keys.push(e.key.toUpperCase());
  
  const combination = keys.join('+');
  
  // Find matching kbd elements and highlight temporarily
  document.querySelectorAll('.kbd').forEach(kbd => {
    if (kbd.textContent === combination) {
      kbd.classList.add('kbd-active');
      setTimeout(() => kbd.classList.remove('kbd-active'), 150);
    }
  });
});
```

### React Integration

```jsx
import React from 'react';

function Kbd({ 
  children, 
  dataKey, 
  size = 'default', 
  variant = 'default',
  disabled = false,
  className = '' 
}) {
  const keySymbols = {
    cmd: '⌘',
    shift: '⇧',
    option: '⌥',
    ctrl: '⌃',
    enter: '⏎',
    esc: '⎋',
    tab: '⇥',
    space: '␣',
    'arrow-up': '↑',
    'arrow-down': '↓',
    'arrow-left': '←',
    'arrow-right': '→'
  };

  const sizeClasses = {
    sm: 'kbd-sm',
    default: '',
    lg: 'kbd-lg'
  };

  const variantClasses = {
    default: '',
    outline: 'kbd-outline',
    solid: 'kbd-solid',
    accent: 'kbd-accent'
  };

  const classes = [
    'kbd',
    sizeClasses[size],
    variantClasses[variant],
    disabled && 'kbd-disabled',
    className
  ].filter(Boolean).join(' ');

  const content = dataKey ? keySymbols[dataKey] || dataKey : children;

  return (
    <kbd className={classes}>
      {content}
    </kbd>
  );
}

// Usage
function ShortcutDisplay() {
  return (
    <div className="space-y-4">
      <div className="flex items-center gap-2">
        <span>Save:</span>
        <Kbd dataKey="cmd" />
        <Kbd>S</Kbd>
      </div>
      
      <button className="btn inline-flex items-center gap-2">
        Search
        <Kbd size="sm" dataKey="cmd" />
        <Kbd size="sm">K</Kbd>
      </button>
    </div>
  );
}
```

### Vue Integration

```vue
<template>
  <kbd :class="kbdClasses">
    {{ displayText }}
  </kbd>
</template>

<script>
export default {
  props: {
    dataKey: String,
    size: {
      type: String,
      default: 'default',
      validator: value => ['sm', 'default', 'lg'].includes(value)
    },
    variant: {
      type: String,
      default: 'default',
      validator: value => ['default', 'outline', 'solid', 'accent'].includes(value)
    },
    disabled: Boolean
  },
  computed: {
    keySymbols() {
      return {
        cmd: '⌘',
        shift: '⇧',
        option: '⌥',
        ctrl: '⌃',
        enter: '⏎',
        esc: '⎋',
        tab: '⇥',
        space: '␣'
      };
    },
    kbdClasses() {
      return [
        'kbd',
        this.size !== 'default' && `kbd-${this.size}`,
        this.variant !== 'default' && `kbd-${this.variant}`,
        this.disabled && 'kbd-disabled'
      ].filter(Boolean).join(' ');
    },
    displayText() {
      if (this.dataKey) {
        return this.keySymbols[this.dataKey] || this.dataKey;
      }
      return this.$slots.default?.[0]?.children || '';
    }
  }
};
</script>
```

## Best Practices

1. **Consistent Symbols**: Use standard symbols (⌘, ⇧, ⌥, ⌃) for modifier keys
2. **Platform Awareness**: Show appropriate shortcuts for user's platform
3. **Logical Grouping**: Group related keys with proper spacing
4. **Context Appropriate**: Use smaller sizes in compact interfaces
5. **Accessibility**: Include descriptive text or aria-labels
6. **Visual Hierarchy**: Don't overuse accent variants
7. **Readability**: Ensure sufficient contrast in all themes

## Common Patterns

### Shortcut List

```html
<div class="space-y-3">
  <div class="flex items-center justify-between py-2">
    <span>New file</span>
    <kbd class="kbd">⌘N</kbd>
  </div>
  <div class="flex items-center justify-between py-2">
    <span>Open file</span>
    <kbd class="kbd">⌘O</kbd>
  </div>
  <div class="flex items-center justify-between py-2">
    <span>Save file</span>
    <kbd class="kbd">⌘S</kbd>
  </div>
</div>
```

### Inline Instructions

```html
<p class="text-sm text-muted-foreground">
  Press <kbd class="kbd kbd-sm">Enter</kbd> to submit or 
  <kbd class="kbd kbd-sm">Esc</kbd> to cancel.
</p>
```

### Complex Shortcuts

```html
<div class="inline-flex items-center gap-1">
  <kbd class="kbd">⌘</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">⇧</kbd>
  <span class="text-muted-foreground text-xs">+</span>
  <kbd class="kbd">P</kbd>
</div>
```

## Related Components

- [Button](./button.md) - Often contains kbd elements
- [Tooltip](./tooltip.md) - For showing shortcuts on hover
- [Command](./command.md) - For command palette interfaces
- [Badge](./badge.md) - Similar styling patterns
---
## slider

# Slider Component

An input component that allows users to select a value from a range by dragging a handle along a track.

## Basic Usage

```html
<input type="range" class="slider" min="0" max="100" value="50">
```

## CSS Classes

### Base Classes
- **`slider`** - Base range input styling with custom track and thumb

### Size Variants
- **`slider-sm`** - Smaller slider for compact interfaces
- **`slider-lg`** - Larger slider for better visibility

### Color Variants
- **`slider-secondary`** - Secondary color theme
- **`slider-success`** - Success/green color theme
- **`slider-warning`** - Warning/yellow color theme
- **`slider-error`** - Error/red color theme

### Container Classes
- **`slider-container`** - Wrapper for slider with labels
- **`slider-track`** - Custom track styling
- **`slider-thumb`** - Custom thumb styling

### State Classes
- **`slider:disabled`** - Disabled state styling
- **`slider:focus`** - Focus state with ring
- **`slider[data-error]`** - Error state styling

## Component Attributes

### Range Input Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `type` | string | Must be "range" | Yes |
| `class` | string | Must include "slider" | Yes |
| `min` | number | Minimum value | Optional |
| `max` | number | Maximum value | Optional |
| `value` | number | Current value | Optional |
| `step` | number | Step increment | Optional |

### Optional Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `disabled` | boolean | Disable interaction | No |
| `data-error` | boolean | Error state indicator | No |
| `aria-label` | string | Accessibility label | Recommended |
| `aria-describedby` | string | Reference to description | Optional |

## No JavaScript Required (Basic)
Basic slider functionality works with native HTML5 range input behavior.

## HTML Structure

```html
<!-- Basic slider -->
<input type="range" class="slider" min="0" max="100" value="50">

<!-- With container and labels -->
<div class="slider-container">
  <div class="flex items-center justify-between mb-2">
    <label for="volume">Volume</label>
    <span class="text-sm text-muted-foreground">50%</span>
  </div>
  <input type="range" id="volume" class="slider" min="0" max="100" value="50">
  <div class="flex justify-between text-xs text-muted-foreground mt-1">
    <span>0</span>
    <span>100</span>
  </div>
</div>
```

## Examples

### Basic Slider

```html
<div class="space-y-4">
  <div>
    <label for="basic-slider" class="block text-sm font-medium mb-2">Basic Slider</label>
    <input type="range" id="basic-slider" class="slider w-full" min="0" max="100" value="30">
  </div>
</div>
```

### Slider with Value Display

```html
<div class="space-y-2">
  <div class="flex items-center justify-between">
    <label for="volume-slider" class="text-sm font-medium">Volume</label>
    <span id="volume-value" class="text-sm text-muted-foreground">75%</span>
  </div>
  <input 
    type="range" 
    id="volume-slider" 
    class="slider w-full" 
    min="0" 
    max="100" 
    value="75"
    oninput="document.getElementById('volume-value').textContent = this.value + '%'"
  >
</div>
```

### Size Variants

```html
<div class="space-y-6">
  <!-- Small -->
  <div>
    <label class="block text-sm font-medium mb-2">Small Slider</label>
    <input type="range" class="slider slider-sm w-full" min="0" max="100" value="25">
  </div>
  
  <!-- Default -->
  <div>
    <label class="block text-sm font-medium mb-2">Default Slider</label>
    <input type="range" class="slider w-full" min="0" max="100" value="50">
  </div>
  
  <!-- Large -->
  <div>
    <label class="block text-sm font-medium mb-2">Large Slider</label>
    <input type="range" class="slider slider-lg w-full" min="0" max="100" value="75">
  </div>
</div>
```

### Color Variants

```html
<div class="space-y-6">
  <!-- Default -->
  <div>
    <label class="block text-sm font-medium mb-2">Default (Primary)</label>
    <input type="range" class="slider w-full" min="0" max="100" value="60">
  </div>
  
  <!-- Secondary -->
  <div>
    <label class="block text-sm font-medium mb-2">Secondary</label>
    <input type="range" class="slider slider-secondary w-full" min="0" max="100" value="40">
  </div>
  
  <!-- Success -->
  <div>
    <label class="block text-sm font-medium mb-2">Success</label>
    <input type="range" class="slider slider-success w-full" min="0" max="100" value="80">
  </div>
  
  <!-- Warning -->
  <div>
    <label class="block text-sm font-medium mb-2">Warning</label>
    <input type="range" class="slider slider-warning w-full" min="0" max="100" value="30">
  </div>
  
  <!-- Error -->
  <div>
    <label class="block text-sm font-medium mb-2">Error</label>
    <input type="range" class="slider slider-error w-full" min="0" max="100" value="20">
  </div>
</div>
```

### Slider with Min/Max Labels

```html
<div class="space-y-2">
  <label for="range-slider" class="block text-sm font-medium">Price Range</label>
  <input type="range" id="range-slider" class="slider w-full" min="0" max="1000" value="250" step="10">
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>$0</span>
    <span>$1,000</span>
  </div>
</div>
```

### Disabled Slider

```html
<div class="space-y-2">
  <label for="disabled-slider" class="block text-sm font-medium text-muted-foreground">
    Disabled Slider
  </label>
  <input 
    type="range" 
    id="disabled-slider" 
    class="slider w-full" 
    min="0" 
    max="100" 
    value="40" 
    disabled
  >
  <p class="text-xs text-muted-foreground">This setting is currently unavailable</p>
</div>
```

### Slider with Steps and Ticks

```html
<div class="space-y-2">
  <label for="stepped-slider" class="block text-sm font-medium">Rating</label>
  <input 
    type="range" 
    id="stepped-slider" 
    class="slider w-full" 
    min="1" 
    max="5" 
    value="3" 
    step="1"
  >
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>1</span>
    <span>2</span>
    <span>3</span>
    <span>4</span>
    <span>5</span>
  </div>
</div>
```

### Brightness Control

```html
<div class="space-y-2">
  <div class="flex items-center justify-between">
    <div class="flex items-center gap-2">
      <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
        <circle cx="12" cy="12" r="4"/>
        <path d="M12 2v2"/>
        <path d="M12 20v2"/>
        <path d="m4.93 4.93 1.41 1.41"/>
        <path d="m17.66 17.66 1.41 1.41"/>
        <path d="M2 12h2"/>
        <path d="M20 12h2"/>
        <path d="m6.34 17.66-1.41 1.41"/>
        <path d="m19.07 4.93-1.41 1.41"/>
      </svg>
      <label for="brightness" class="text-sm font-medium">Brightness</label>
    </div>
    <span id="brightness-value" class="text-sm text-muted-foreground">70%</span>
  </div>
  <input 
    type="range" 
    id="brightness" 
    class="slider w-full" 
    min="0" 
    max="100" 
    value="70"
    oninput="document.getElementById('brightness-value').textContent = this.value + '%'"
  >
</div>
```

### Temperature Control

```html
<div class="space-y-2">
  <div class="flex items-center justify-between">
    <label for="temperature" class="text-sm font-medium">Temperature</label>
    <span id="temp-value" class="text-sm text-muted-foreground">22°C</span>
  </div>
  <input 
    type="range" 
    id="temperature" 
    class="slider slider-warning w-full" 
    min="16" 
    max="30" 
    value="22"
    oninput="document.getElementById('temp-value').textContent = this.value + '°C'"
  >
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>❄️ 16°C</span>
    <span>🔥 30°C</span>
  </div>
</div>
```

### Progress Indicator

```html
<div class="space-y-2">
  <div class="flex items-center justify-between">
    <label for="progress" class="text-sm font-medium">Upload Progress</label>
    <span id="progress-value" class="text-sm text-muted-foreground">65%</span>
  </div>
  <input 
    type="range" 
    id="progress" 
    class="slider slider-success w-full" 
    min="0" 
    max="100" 
    value="65"
    disabled
  >
  <p class="text-xs text-muted-foreground">Uploading file... Please wait.</p>
</div>
```

### Media Player Volume

```html
<div class="flex items-center gap-3">
  <button class="btn-icon-ghost">
    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
      <polygon points="11 5 6 9 2 9 2 15 6 15 11 19 11 5"/>
      <path d="M15.54 8.46a5 5 0 0 1 0 7.07"/>
    </svg>
  </button>
  <input 
    type="range" 
    class="slider w-24" 
    min="0" 
    max="100" 
    value="75"
    aria-label="Volume control"
  >
</div>
```

### Form Field Integration

```html
<div class="field">
  <label for="opacity" class="block text-sm font-medium mb-2">
    Opacity
    <span class="text-muted-foreground font-normal">(Optional)</span>
  </label>
  <div class="space-y-2">
    <input 
      type="range" 
      id="opacity" 
      class="slider w-full" 
      min="0" 
      max="100" 
      value="90"
      aria-describedby="opacity-help"
    >
    <div class="flex justify-between text-xs text-muted-foreground">
      <span>Transparent</span>
      <span>Opaque</span>
    </div>
  </div>
  <p id="opacity-help" class="text-sm text-muted-foreground mt-1">
    Adjust the transparency of the element from 0% (invisible) to 100% (fully visible).
  </p>
</div>
```

### Error State

```html
<div class="field">
  <label for="invalid-slider" class="block text-sm font-medium mb-2">
    Value Range
  </label>
  <input 
    type="range" 
    id="invalid-slider" 
    class="slider slider-error w-full" 
    min="10" 
    max="90" 
    value="5"
    data-error="true"
    aria-invalid="true"
    aria-describedby="slider-error"
  >
  <p id="slider-error" class="text-sm text-destructive mt-1" role="alert">
    Value must be between 10 and 90.
  </p>
</div>
```

## Accessibility Features

- **Keyboard Support**: Arrow keys for fine adjustment, Page Up/Down for larger steps
- **Screen Reader Support**: Proper labeling and value announcements
- **Focus Management**: Clear focus indicators
- **ARIA Attributes**: Support for aria-label, aria-describedby, aria-invalid

### Enhanced Accessibility

```html
<div class="field">
  <label for="accessible-slider" id="slider-label" class="block text-sm font-medium mb-2">
    Audio Volume
  </label>
  <input 
    type="range" 
    id="accessible-slider"
    class="slider w-full"
    min="0" 
    max="100" 
    value="50"
    step="5"
    role="slider"
    aria-labelledby="slider-label"
    aria-describedby="slider-instructions slider-value"
    aria-valuemin="0"
    aria-valuemax="100"
    aria-valuenow="50"
    aria-valuetext="50 percent"
  >
  <div class="flex justify-between items-center text-xs text-muted-foreground mt-1">
    <span>Mute</span>
    <span id="slider-value" aria-live="polite">50%</span>
    <span>Max</span>
  </div>
  <p id="slider-instructions" class="sr-only">
    Use arrow keys to adjust volume. Left and right arrows make small adjustments, page up and page down make larger adjustments.
  </p>
</div>
```

## JavaScript Integration

### Basic Value Updates

```javascript
// Update display value
function updateSliderValue(slider, displayElement, suffix = '') {
  displayElement.textContent = slider.value + suffix;
  slider.addEventListener('input', () => {
    displayElement.textContent = slider.value + suffix;
  });
}

// Usage
const volumeSlider = document.getElementById('volume');
const volumeDisplay = document.getElementById('volume-value');
updateSliderValue(volumeSlider, volumeDisplay, '%');
```

### Advanced Slider Controls

```javascript
class SliderControl {
  constructor(element, options = {}) {
    this.slider = element;
    this.options = {
      displayValue: true,
      suffix: '',
      prefix: '',
      formatter: null,
      onChange: null,
      ...options
    };
    
    this.init();
  }
  
  init() {
    if (this.options.displayValue) {
      this.createValueDisplay();
    }
    
    this.slider.addEventListener('input', (e) => {
      this.updateValue();
      if (this.options.onChange) {
        this.options.onChange(e.target.value, e);
      }
    });
    
    // Initial value
    this.updateValue();
  }
  
  createValueDisplay() {
    this.valueDisplay = document.createElement('span');
    this.valueDisplay.className = 'text-sm text-muted-foreground';
    
    // Insert after label or before slider
    const label = this.slider.previousElementSibling;
    if (label && label.tagName === 'LABEL') {
      const wrapper = document.createElement('div');
      wrapper.className = 'flex items-center justify-between mb-2';
      label.parentNode.insertBefore(wrapper, label);
      wrapper.appendChild(label);
      wrapper.appendChild(this.valueDisplay);
    }
  }
  
  updateValue() {
    const value = parseInt(this.slider.value);
    let displayValue = value;
    
    if (this.options.formatter) {
      displayValue = this.options.formatter(value);
    } else {
      displayValue = this.options.prefix + value + this.options.suffix;
    }
    
    if (this.valueDisplay) {
      this.valueDisplay.textContent = displayValue;
    }
    
    // Update ARIA
    this.slider.setAttribute('aria-valuenow', value);
    this.slider.setAttribute('aria-valuetext', displayValue);
  }
  
  setValue(value) {
    this.slider.value = value;
    this.updateValue();
  }
  
  getValue() {
    return parseInt(this.slider.value);
  }
}

// Usage examples
new SliderControl(document.getElementById('volume'), {
  suffix: '%',
  onChange: (value) => console.log('Volume changed to:', value)
});

new SliderControl(document.getElementById('temperature'), {
  suffix: '°C',
  formatter: (value) => `${value}°C (${Math.round(value * 9/5 + 32)}°F)`
});
```

### React Integration

```jsx
import React, { useState, useCallback } from 'react';

function Slider({
  min = 0,
  max = 100,
  value = 50,
  step = 1,
  disabled = false,
  size = 'default',
  variant = 'default',
  label,
  showValue = false,
  suffix = '',
  prefix = '',
  formatter,
  onChange,
  className = ''
}) {
  const [currentValue, setCurrentValue] = useState(value);
  
  const handleChange = useCallback((e) => {
    const newValue = parseInt(e.target.value);
    setCurrentValue(newValue);
    if (onChange) {
      onChange(newValue);
    }
  }, [onChange]);
  
  const sizeClasses = {
    sm: 'slider-sm',
    default: '',
    lg: 'slider-lg'
  };
  
  const variantClasses = {
    default: '',
    secondary: 'slider-secondary',
    success: 'slider-success',
    warning: 'slider-warning',
    error: 'slider-error'
  };
  
  const sliderClasses = [
    'slider',
    'w-full',
    sizeClasses[size],
    variantClasses[variant],
    className
  ].filter(Boolean).join(' ');
  
  const formatValue = (val) => {
    if (formatter) return formatter(val);
    return prefix + val + suffix;
  };
  
  return (
    <div className="space-y-2">
      {(label || showValue) && (
        <div className="flex items-center justify-between">
          {label && <label className="text-sm font-medium">{label}</label>}
          {showValue && (
            <span className="text-sm text-muted-foreground">
              {formatValue(currentValue)}
            </span>
          )}
        </div>
      )}
      
      <input
        type="range"
        className={sliderClasses}
        min={min}
        max={max}
        step={step}
        value={currentValue}
        disabled={disabled}
        onChange={handleChange}
        aria-valuemin={min}
        aria-valuemax={max}
        aria-valuenow={currentValue}
        aria-valuetext={formatValue(currentValue)}
      />
    </div>
  );
}

// Usage
function VolumeControl() {
  const [volume, setVolume] = useState(75);
  
  return (
    <Slider
      label="Volume"
      value={volume}
      onChange={setVolume}
      showValue
      suffix="%"
      variant="default"
    />
  );
}
```

### Vue Integration

```vue
<template>
  <div class="space-y-2">
    <div v-if="label || showValue" class="flex items-center justify-between">
      <label v-if="label" class="text-sm font-medium">{{ label }}</label>
      <span v-if="showValue" class="text-sm text-muted-foreground">
        {{ formatValue(currentValue) }}
      </span>
    </div>
    
    <input
      v-model="currentValue"
      type="range"
      :class="sliderClasses"
      :min="min"
      :max="max"
      :step="step"
      :disabled="disabled"
      :aria-valuemin="min"
      :aria-valuemax="max"
      :aria-valuenow="currentValue"
      :aria-valuetext="formatValue(currentValue)"
      @input="handleChange"
    />
  </div>
</template>

<script>
export default {
  props: {
    min: { type: Number, default: 0 },
    max: { type: Number, default: 100 },
    value: { type: Number, default: 50 },
    step: { type: Number, default: 1 },
    disabled: { type: Boolean, default: false },
    size: { type: String, default: 'default' },
    variant: { type: String, default: 'default' },
    label: String,
    showValue: { type: Boolean, default: false },
    suffix: { type: String, default: '' },
    prefix: { type: String, default: '' },
    formatter: Function
  },
  emits: ['update:value', 'change'],
  data() {
    return {
      currentValue: this.value
    };
  },
  computed: {
    sliderClasses() {
      const sizeClasses = {
        sm: 'slider-sm',
        default: '',
        lg: 'slider-lg'
      };
      
      const variantClasses = {
        default: '',
        secondary: 'slider-secondary',
        success: 'slider-success',
        warning: 'slider-warning',
        error: 'slider-error'
      };
      
      return [
        'slider',
        'w-full',
        sizeClasses[this.size],
        variantClasses[this.variant]
      ].filter(Boolean).join(' ');
    }
  },
  methods: {
    formatValue(value) {
      if (this.formatter) return this.formatter(value);
      return this.prefix + value + this.suffix;
    },
    handleChange(e) {
      const newValue = parseInt(e.target.value);
      this.currentValue = newValue;
      this.$emit('update:value', newValue);
      this.$emit('change', newValue);
    }
  },
  watch: {
    value(newValue) {
      this.currentValue = newValue;
    }
  }
};
</script>
```

## Best Practices

1. **Clear Labels**: Always provide descriptive labels for sliders
2. **Value Display**: Show current value for important controls
3. **Reasonable Ranges**: Use appropriate min/max values
4. **Logical Steps**: Choose step values that make sense
5. **Visual Feedback**: Use colors to indicate state or importance
6. **Accessibility**: Include proper ARIA attributes
7. **Responsive Design**: Ensure sliders work on touch devices

## Common Patterns

### Settings Panel

```html
<div class="space-y-6">
  <div class="space-y-2">
    <div class="flex items-center justify-between">
      <label class="text-sm font-medium">Master Volume</label>
      <span class="text-sm text-muted-foreground">80%</span>
    </div>
    <input type="range" class="slider w-full" min="0" max="100" value="80">
  </div>
  
  <div class="space-y-2">
    <div class="flex items-center justify-between">
      <label class="text-sm font-medium">Brightness</label>
      <span class="text-sm text-muted-foreground">60%</span>
    </div>
    <input type="range" class="slider slider-warning w-full" min="0" max="100" value="60">
  </div>
</div>
```

### Range Selection

```html
<div class="space-y-4">
  <label class="text-sm font-medium">Price Range</label>
  <div class="space-y-2">
    <input type="range" class="slider w-full" min="0" max="1000" value="100" placeholder="Min price">
    <input type="range" class="slider w-full" min="0" max="1000" value="800" placeholder="Max price">
  </div>
  <div class="flex justify-between text-xs text-muted-foreground">
    <span>$0</span>
    <span>$1,000</span>
  </div>
</div>
```

## Related Components

- [Input](./input.md) - For text-based input alternatives
- [Button](./button.md) - For increment/decrement controls
- [Progress](./progress.md) - For progress indication
- [Field](./field.md) - For form field integration
---
## README

# Basecoat UI Component Documentation

Complete documentation for all Basecoat UI components with examples, usage patterns, and integration guides.

## Overview

This documentation covers **34+ components** from the Basecoat UI library, providing comprehensive guides for implementation, customization, and best practices. Each component includes HTML examples, CSS classes, accessibility features, and JavaScript integration patterns.

## 📚 Component Categories

### Form Components
- **[Button](./button.md)** - Interactive buttons with multiple variants and states
- **[Button Group](./button-group.md)** - Groups of related buttons with consistent styling
- **[Input](./input.md)** - Text input fields with validation and styling options
- **[Input Group](./input-group.md)** - Input fields with icons, labels, and enhancements
- **[Textarea](./textarea.md)** - Multi-line text input areas
- **[Select](./select.md)** - Dropdown selection menus with custom styling
- **[Combobox](./combobox.md)** - Searchable select components
- **[Checkbox](./checkbox.md)** - Checkbox inputs with custom styling
- **[Radio Group](./radio-group.md)** - Radio button selections
- **[Switch](./switch.md)** - Toggle switch controls
- **[Slider](./slider.md)** - Range input controls with custom styling
- **[Field](./field.md)** - Complete form field wrapper with labels and validation
- **[Form](./form.md)** - Form container with automatic child element styling
- **[Label](./label.md)** - Form labels with proper accessibility
- **[Kbd](./kbd.md)** - Keyboard shortcut and key combination display

### Navigation Components
- **[Breadcrumb](./breadcrumb.md)** - Hierarchical navigation paths
- **[Pagination](./pagination.md)** - Page navigation controls
- **[Sidebar](./sidebar.md)** - Collapsible sidebar navigation

### Interactive Components
- **[Dialog](./dialog.md)** - Modal dialogs and overlays
- **[Alert Dialog](./alert-dialog.md)** - Confirmation and alert modals
- **[Dropdown Menu](./dropdown-menu.md)** - Context menus and dropdowns
- **[Popover](./popover.md)** - Floating content containers
- **[Tooltip](./tooltip.md)** - Hover and focus information displays
- **[Command](./command.md)** - Command palette and search interfaces
- **[Toast](./toast.md)** - Temporary notification messages
- **[Theme Switcher](./theme-switcher.md)** - Dark/light mode toggle controls

### Layout Components
- **[Card](./card.md)** - Content containers with headers and actions
- **[Accordion](./accordion.md)** - Collapsible content sections
- **[Tabs](./tabs.md)** - Tabbed content interfaces
- **[Table](./table.md)** - Data tables with sorting and styling

### Feedback Components
- **[Alert](./alert.md)** - Information and status messages
- **[Empty](./empty.md)** - Empty state displays with actions and guidance

### Utility Components
- **[Avatar](./avatar.md)** - User profile images and placeholders
- **[Badge](./badge.md)** - Status indicators and labels
- **[Spinner](./spinner.md)** - Loading indicators
- **[Progress](./progress.md)** - Progress bars and indicators
- **[Skeleton](./skeleton.md)** - Loading placeholders

### Data Components
- **[Chart](./chart.md)** - Data visualizations with Chart.js integration
- **[Carousel](./carousel.md)** - Image and content carousels

### Pattern Components
- **[Item](./item.md)** - Versatile content display patterns

## 🚀 Getting Started

### Basic Usage

Each component follows a consistent pattern:

```html
<element class="base-class variant-class size-class">
  Content
</element>
```

### Example: Button Component

```html
<!-- Basic button -->
<button class="btn">Click me</button>

<!-- Button variants -->
<button class="btn btn-outline">Outline</button>
<button class="btn btn-ghost">Ghost</button>

<!-- Button sizes -->
<button class="btn btn-sm">Small</button>
<button class="btn btn-lg">Large</button>
```

### CSS Integration

Include Basecoat CSS in your project:

```html
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/basecoat-css@latest/dist/basecoat.css">
```

### JavaScript Integration

For interactive components, include Basecoat JavaScript:

```html
<script src="https://cdn.jsdelivr.net/npm/basecoat-css@latest/dist/js/basecoat.min.js" defer></script>
```

## 🎨 Design Principles

### Atomic Design
Components are organized following atomic design principles:
- **Atoms**: Basic building blocks (Button, Input, Badge)
- **Molecules**: Simple combinations (Form Field, Search Bar)
- **Organisms**: Complex components (Data Table, Navigation)
- **Templates**: Page layouts and structures
- **Pages**: Complete interfaces

### Theme Integration
All components use CSS custom properties for theming:

```css
:root {
  --primary: 217 91% 60%;
  --primary-foreground: 0 0% 100%;
  --background: 0 0% 100%;
  --foreground: 0 0% 9%;
  --muted: 0 0% 96%;
  --muted-foreground: 0 0% 45%;
  --border: 0 0% 90%;
  --ring: 217 91% 60%;
}
```

### Accessibility First
Every component includes:
- Proper semantic HTML
- ARIA attributes and roles
- Keyboard navigation support
- Screen reader compatibility
- High contrast mode support

## 💻 Framework Integration

### React Example

```jsx
import React from 'react';

function MyButton({ variant = 'default', size = 'md', children, ...props }) {
  const classes = [
    'btn',
    variant !== 'default' && `btn-${variant}`,
    size !== 'md' && `btn-${size}`
  ].filter(Boolean).join(' ');
  
  return (
    <button className={classes} {...props}>
      {children}
    </button>
  );
}
```

### Vue Example

```vue
<template>
  <button :class="buttonClasses" v-bind="$attrs">
    <slot />
  </button>
</template>

<script>
export default {
  props: {
    variant: { type: String, default: 'default' },
    size: { type: String, default: 'md' }
  },
  computed: {
    buttonClasses() {
      return [
        'btn',
        this.variant !== 'default' && `btn-${this.variant}`,
        this.size !== 'md' && `btn-${this.size}`
      ].filter(Boolean).join(' ');
    }
  }
};
</script>
```

## 🎯 Common Patterns

### Form Layout

```html
<form class="form grid gap-6">
  <div class="grid gap-2">
    <label for="username">Username</label>
    <input type="text" id="username" placeholder="Enter username">
  </div>
  
  <div class="grid gap-2">
    <label for="email">Email</label>
    <input type="email" id="email" placeholder="Enter email">
  </div>
  
  <button type="submit" class="btn">Submit</button>
</form>
```

### Card with Actions

```html
<div class="card p-6">
  <header class="mb-4">
    <h3 class="text-lg font-semibold">Card Title</h3>
    <p class="text-muted-foreground">Card description</p>
  </header>
  
  <div class="content">
    <!-- Card content -->
  </div>
  
  <footer class="flex gap-2 mt-4">
    <button class="btn">Primary Action</button>
    <button class="btn-outline">Secondary</button>
  </footer>
</div>
```

### Data Display

```html
<div class="space-y-4">
  <!-- Stats -->
  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <div class="card p-4 text-center">
      <div class="text-2xl font-bold">1,234</div>
      <div class="text-sm text-muted-foreground">Total Users</div>
    </div>
  </div>
  
  <!-- Table -->
  <div class="card">
    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>John Doe</td>
          <td>john@example.com</td>
          <td><span class="badge badge-success">Active</span></td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
```

## 🔧 Customization

### CSS Custom Properties

Override theme variables for customization:

```css
:root {
  --primary: 142 76% 36%;     /* Custom green */
  --secondary: 221 83% 53%;   /* Custom blue */
  --radius: 0.75rem;          /* More rounded corners */
}
```

### Component Variants

Create custom component variants:

```css
.btn-custom {
  @apply bg-gradient-to-r from-purple-500 to-pink-500 text-white;
}

.btn-custom:hover {
  @apply from-purple-600 to-pink-600;
}
```

### Responsive Breakpoints

Use Tailwind's responsive prefixes:

```html
<button class="btn btn-sm md:btn-md lg:btn-lg">
  Responsive Button
</button>
```

## 📱 Mobile Considerations

### Touch-Friendly Sizes

```html
<!-- Minimum 44px touch target -->
<button class="btn min-h-11 px-4">Mobile Button</button>
```

### Responsive Typography

```html
<h1 class="text-xl md:text-2xl lg:text-3xl">
  Responsive Heading
</h1>
```

### Mobile Navigation

```html
<nav class="sidebar md:sidebar-desktop">
  <!-- Navigation content -->
</nav>
```

## ♿ Accessibility Guidelines

### Semantic HTML
Always use appropriate HTML elements:

```html
<!-- Good -->
<button class="btn">Submit</button>
<nav class="breadcrumb">...</nav>

<!-- Avoid -->
<div class="btn" onclick="submit()">Submit</div>
```

### ARIA Labels

```html
<button class="btn-icon" aria-label="Close dialog">
  <svg><!-- close icon --></svg>
</button>
```

### Focus Management

```html
<div class="dialog" role="dialog" aria-labelledby="dialog-title">
  <h2 id="dialog-title">Dialog Title</h2>
  <!-- Dialog content -->
</div>
```

### Color Contrast

Ensure sufficient contrast ratios:
- Normal text: 4.5:1 minimum
- Large text: 3:1 minimum
- Interactive elements: 3:1 minimum

## 🚀 Performance Tips

### CSS Optimization

```css
/* Use CSS containment for performance */
.card {
  contain: layout style paint;
}

/* Optimize animations */
.btn {
  will-change: transform;
  transition: transform 0.2s ease;
}
```

### JavaScript Best Practices

```javascript
// Use event delegation
document.addEventListener('click', (e) => {
  if (e.target.matches('.btn')) {
    // Handle button click
  }
});

// Debounce search inputs
const debounce = (fn, delay) => {
  let timeoutId;
  return (...args) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn.apply(null, args), delay);
  };
};
```

## 📋 Best Practices Checklist

### Component Usage
- [ ] Use semantic HTML elements
- [ ] Include proper ARIA attributes
- [ ] Test keyboard navigation
- [ ] Verify screen reader compatibility
- [ ] Check color contrast ratios
- [ ] Test on multiple devices/browsers
- [ ] Validate responsive behavior

### Development
- [ ] Follow consistent naming conventions
- [ ] Use appropriate component variants
- [ ] Include loading and error states
- [ ] Handle edge cases gracefully
- [ ] Optimize for performance
- [ ] Document component usage
- [ ] Test with real data

### Design
- [ ] Maintain visual hierarchy
- [ ] Use consistent spacing
- [ ] Follow brand guidelines
- [ ] Consider user workflow
- [ ] Provide clear feedback
- [ ] Support multiple themes
- [ ] Plan for internationalization

## 🔗 Resources

- **[Basecoat GitHub](https://github.com/hunvreus/basecoat)** - Source code and issues
- **[Basecoat Website](https://basecoatui.com)** - Official documentation
- **[Tailwind CSS](https://tailwindcss.com)** - Utility-first CSS framework
- **[Web Accessibility Guidelines](https://www.w3.org/WAI/WCAG21/quickref/)** - WCAG 2.1 quick reference

## 📄 License

This documentation is provided under the same license as the Basecoat UI library. Please refer to the main project repository for license details.

---

**Last Updated**: November 2024
**Version**: Compatible with Basecoat v0.3.6+
---
## link

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
