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