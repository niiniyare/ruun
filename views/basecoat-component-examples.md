# Basecoat Component Examples

A practical reference guide with copy-paste examples for all major Basecoat components.

## Table of Contents

- [Button Components](#button-components)
- [Form Components](#form-components)
- [Layout Components](#layout-components)
- [Navigation Components](#navigation-components)
- [Feedback Components](#feedback-components)

---

## Button Components

### Basic Button Usage

```html
<!-- Default Button -->
<button data-component="button">Default Button</button>

<!-- With explicit variant -->
<button data-component="button" data-variant="primary">Primary Button</button>
```

### All Button Variants

```html
<!-- Primary (default) -->
<button data-component="button" data-variant="primary">Primary</button>

<!-- Secondary -->
<button data-component="button" data-variant="secondary">Secondary</button>

<!-- Outline -->
<button data-component="button" data-variant="outline">Outline</button>

<!-- Ghost -->
<button data-component="button" data-variant="ghost">Ghost</button>

<!-- Destructive -->
<button data-component="button" data-variant="destructive">Destructive</button>

<!-- Link style -->
<button data-component="button" data-variant="link">Link</button>
```

### Button Sizes

```html
<!-- Small -->
<button data-component="button" data-size="sm">Small Button</button>

<!-- Default -->
<button data-component="button" data-size="default">Default Size</button>

<!-- Large -->
<button data-component="button" data-size="lg">Large Button</button>

<!-- Icon-only button -->
<button data-component="button" data-size="icon" aria-label="Settings">
  <svg><!-- icon --></svg>
</button>
```

### Button States

```html
<!-- Disabled -->
<button data-component="button" disabled>Disabled Button</button>

<!-- Loading state -->
<button data-component="button" data-loading="true" disabled>
  <span data-slot="loader" class="mr-2">
    <svg class="animate-spin h-4 w-4"><!-- spinner icon --></svg>
  </span>
  Loading...
</button>

<!-- With icon -->
<button data-component="button">
  <svg class="mr-2 h-4 w-4"><!-- icon --></svg>
  Button with Icon
</button>
```

### Button as Link

```html
<!-- Button styled link -->
<a href="/page" data-component="button" data-variant="primary">Link Button</a>

<!-- External link -->
<a href="https://example.com" data-component="button" target="_blank" rel="noopener noreferrer">
  External Link
  <svg class="ml-2 h-4 w-4"><!-- external icon --></svg>
</a>
```

---

## Form Components

### Input Field

```html
<!-- Basic input -->
<div data-component="form-field">
  <label data-slot="label" for="email">Email</label>
  <input data-slot="input" type="email" id="email" name="email" placeholder="Enter your email">
</div>

<!-- Required field -->
<div data-component="form-field">
  <label data-slot="label" for="username">
    Username <span aria-label="required">*</span>
  </label>
  <input data-slot="input" type="text" id="username" name="username" required>
</div>

<!-- With help text -->
<div data-component="form-field">
  <label data-slot="label" for="password">Password</label>
  <input data-slot="input" type="password" id="password" name="password">
  <p data-slot="description">Must be at least 8 characters</p>
</div>

<!-- Error state -->
<div data-component="form-field" data-error="true">
  <label data-slot="label" for="email-error">Email</label>
  <input data-slot="input" type="email" id="email-error" aria-invalid="true" aria-describedby="email-error-msg">
  <p data-slot="error" id="email-error-msg">Please enter a valid email address</p>
</div>

<!-- Disabled -->
<div data-component="form-field">
  <label data-slot="label" for="disabled-input">Disabled Field</label>
  <input data-slot="input" type="text" id="disabled-input" disabled value="Cannot edit">
</div>
```

### Textarea

```html
<!-- Basic textarea -->
<div data-component="form-field">
  <label data-slot="label" for="message">Message</label>
  <textarea data-slot="input" id="message" name="message" rows="4" placeholder="Enter your message"></textarea>
</div>

<!-- With character count -->
<div data-component="form-field">
  <label data-slot="label" for="bio">Bio</label>
  <textarea data-slot="input" id="bio" name="bio" maxlength="200"></textarea>
  <p data-slot="description">
    <span data-count="bio">0</span>/200 characters
  </p>
</div>
```

### Select Dropdown

```html
<!-- Basic select -->
<div data-component="form-field">
  <label data-slot="label" for="country">Country</label>
  <select data-slot="input" id="country" name="country">
    <option value="">Select a country</option>
    <option value="us">United States</option>
    <option value="uk">United Kingdom</option>
    <option value="ca">Canada</option>
  </select>
</div>

<!-- Custom select with dropdown component -->
<div data-component="dropdown" data-form-field="true">
  <button data-slot="trigger" type="button" aria-haspopup="listbox" aria-expanded="false">
    <span data-slot="value">Select option</span>
    <svg data-slot="icon"><!-- chevron icon --></svg>
  </button>
  <div data-slot="content" role="listbox" data-align="start">
    <div data-slot="item" role="option" data-value="option1">Option 1</div>
    <div data-slot="item" role="option" data-value="option2">Option 2</div>
    <div data-slot="item" role="option" data-value="option3">Option 3</div>
  </div>
</div>
```

### Checkbox

```html
<!-- Single checkbox -->
<div data-component="checkbox">
  <input data-slot="input" type="checkbox" id="terms" name="terms">
  <label data-slot="label" for="terms">I agree to the terms and conditions</label>
</div>

<!-- Checkbox group -->
<fieldset data-component="form-field">
  <legend data-slot="label">Select your interests</legend>
  <div class="space-y-2">
    <div data-component="checkbox">
      <input data-slot="input" type="checkbox" id="coding" name="interests" value="coding">
      <label data-slot="label" for="coding">Coding</label>
    </div>
    <div data-component="checkbox">
      <input data-slot="input" type="checkbox" id="design" name="interests" value="design">
      <label data-slot="label" for="design">Design</label>
    </div>
    <div data-component="checkbox">
      <input data-slot="input" type="checkbox" id="music" name="interests" value="music">
      <label data-slot="label" for="music">Music</label>
    </div>
  </div>
</fieldset>

<!-- Indeterminate state (requires JS) -->
<div data-component="checkbox">
  <input data-slot="input" type="checkbox" id="select-all" name="select-all">
  <label data-slot="label" for="select-all">Select All</label>
</div>
<script>
  document.getElementById('select-all').indeterminate = true;
</script>
```

### Radio Group

```html
<!-- Radio group -->
<fieldset data-component="form-field">
  <legend data-slot="label">Choose your plan</legend>
  <div class="space-y-2" role="radiogroup">
    <div data-component="radio">
      <input data-slot="input" type="radio" id="free" name="plan" value="free">
      <label data-slot="label" for="free">
        <span>Free</span>
        <span class="text-sm text-muted-foreground">$0/month</span>
      </label>
    </div>
    <div data-component="radio">
      <input data-slot="input" type="radio" id="pro" name="plan" value="pro" checked>
      <label data-slot="label" for="pro">
        <span>Pro</span>
        <span class="text-sm text-muted-foreground">$10/month</span>
      </label>
    </div>
    <div data-component="radio">
      <input data-slot="input" type="radio" id="enterprise" name="plan" value="enterprise">
      <label data-slot="label" for="enterprise">
        <span>Enterprise</span>
        <span class="text-sm text-muted-foreground">Contact us</span>
      </label>
    </div>
  </div>
</fieldset>
```

### Switch/Toggle

```html
<!-- Basic switch -->
<div data-component="switch">
  <input data-slot="input" type="checkbox" id="notifications" role="switch">
  <label data-slot="label" for="notifications">Enable notifications</label>
</div>

<!-- Switch with description -->
<div data-component="form-field">
  <div class="flex items-center justify-between">
    <div class="space-y-0.5">
      <label data-slot="label" for="marketing">Marketing emails</label>
      <p data-slot="description">Receive emails about new products and features</p>
    </div>
    <div data-component="switch">
      <input data-slot="input" type="checkbox" id="marketing" role="switch">
    </div>
  </div>
</div>
```

### Complete Form Example

```html
<form data-component="form" method="POST" action="/submit">
  <!-- Name field -->
  <div data-component="form-field">
    <label data-slot="label" for="name">Full Name</label>
    <input data-slot="input" type="text" id="name" name="name" required>
  </div>

  <!-- Email field -->
  <div data-component="form-field">
    <label data-slot="label" for="form-email">Email</label>
    <input data-slot="input" type="email" id="form-email" name="email" required>
    <p data-slot="description">We'll never share your email</p>
  </div>

  <!-- Message field -->
  <div data-component="form-field">
    <label data-slot="label" for="form-message">Message</label>
    <textarea data-slot="input" id="form-message" name="message" rows="4"></textarea>
  </div>

  <!-- Terms checkbox -->
  <div data-component="checkbox">
    <input data-slot="input" type="checkbox" id="form-terms" name="terms" required>
    <label data-slot="label" for="form-terms">I agree to the terms</label>
  </div>

  <!-- Submit button -->
  <button data-component="button" type="submit">Submit Form</button>
</form>
```

---

## Layout Components

### Card

```html
<!-- Basic card -->
<div data-component="card">
  <div data-slot="header">
    <h3 data-slot="title">Card Title</h3>
    <p data-slot="description">Card description goes here</p>
  </div>
  <div data-slot="content">
    <p>Card content goes here. This is the main body of the card.</p>
  </div>
  <div data-slot="footer">
    <button data-component="button" data-variant="outline">Cancel</button>
    <button data-component="button">Save</button>
  </div>
</div>

<!-- Card with image -->
<div data-component="card">
  <img data-slot="image" src="/image.jpg" alt="Card image">
  <div data-slot="header">
    <h3 data-slot="title">Product Name</h3>
  </div>
  <div data-slot="content">
    <p class="text-2xl font-bold">$99.00</p>
    <p class="text-muted-foreground">In stock</p>
  </div>
  <div data-slot="footer">
    <button data-component="button" class="w-full">Add to Cart</button>
  </div>
</div>

<!-- Clickable card -->
<article data-component="card" data-interactive="true" tabindex="0" role="button">
  <div data-slot="header">
    <h3 data-slot="title">Clickable Card</h3>
  </div>
  <div data-slot="content">
    <p>This entire card is clickable</p>
  </div>
</article>
```

### Dialog/Modal

```html
<!-- Dialog trigger and content -->
<button data-component="button" data-dialog-trigger="example-dialog">
  Open Dialog
</button>

<div data-component="dialog" id="example-dialog" aria-labelledby="dialog-title" aria-modal="true">
  <div data-slot="overlay"></div>
  <div data-slot="content">
    <div data-slot="header">
      <h2 data-slot="title" id="dialog-title">Dialog Title</h2>
      <button data-slot="close" aria-label="Close">
        <svg><!-- close icon --></svg>
      </button>
    </div>
    <div data-slot="body">
      <p>Dialog content goes here.</p>
    </div>
    <div data-slot="footer">
      <button data-component="button" data-variant="outline" data-dialog-close>
        Cancel
      </button>
      <button data-component="button">
        Confirm
      </button>
    </div>
  </div>
</div>

<!-- Alert dialog -->
<div data-component="dialog" data-variant="alert" id="delete-dialog">
  <div data-slot="overlay"></div>
  <div data-slot="content">
    <div data-slot="header">
      <h2 data-slot="title">Are you sure?</h2>
    </div>
    <div data-slot="body">
      <p>This action cannot be undone. This will permanently delete your account.</p>
    </div>
    <div data-slot="footer">
      <button data-component="button" data-variant="outline" data-dialog-close>
        Cancel
      </button>
      <button data-component="button" data-variant="destructive">
        Delete Account
      </button>
    </div>
  </div>
</div>
```

### Dialog JavaScript

```javascript
// Initialize dialog
const dialog = new BasecoatDialog('example-dialog');

// Open programmatically
dialog.open();

// Close programmatically
dialog.close();

// Listen for events
dialog.on('open', () => console.log('Dialog opened'));
dialog.on('close', () => console.log('Dialog closed'));
```

### Sidebar

```html
<!-- Sidebar layout -->
<div data-component="sidebar-layout">
  <!-- Sidebar -->
  <aside data-slot="sidebar" data-state="expanded">
    <div data-slot="header">
      <h2>App Name</h2>
    </div>
    <nav data-slot="content">
      <a href="/dashboard" data-component="sidebar-item" data-active="true">
        <svg data-slot="icon"><!-- icon --></svg>
        <span data-slot="label">Dashboard</span>
      </a>
      <a href="/projects" data-component="sidebar-item">
        <svg data-slot="icon"><!-- icon --></svg>
        <span data-slot="label">Projects</span>
      </a>
      <a href="/tasks" data-component="sidebar-item">
        <svg data-slot="icon"><!-- icon --></svg>
        <span data-slot="label">Tasks</span>
        <span data-slot="badge" data-component="badge" data-variant="default">5</span>
      </a>
      
      <!-- Separator -->
      <div data-slot="separator"></div>
      
      <a href="/settings" data-component="sidebar-item">
        <svg data-slot="icon"><!-- icon --></svg>
        <span data-slot="label">Settings</span>
      </a>
    </nav>
    <div data-slot="footer">
      <button data-component="button" data-variant="ghost" data-sidebar-toggle>
        <svg><!-- collapse icon --></svg>
      </button>
    </div>
  </aside>

  <!-- Main content -->
  <main data-slot="main">
    <!-- Page content -->
  </main>
</div>

<!-- Collapsible sidebar -->
<aside data-component="sidebar" data-collapsible="true" data-state="collapsed">
  <!-- Same content structure -->
</aside>
```

### Accordion

```html
<!-- Single accordion -->
<div data-component="accordion" data-type="single">
  <div data-slot="item">
    <button data-slot="trigger" aria-expanded="false">
      <span>Accordion Item 1</span>
      <svg data-slot="icon"><!-- chevron icon --></svg>
    </button>
    <div data-slot="content">
      <p>Content for the first accordion item.</p>
    </div>
  </div>
  <div data-slot="item">
    <button data-slot="trigger" aria-expanded="false">
      <span>Accordion Item 2</span>
      <svg data-slot="icon"><!-- chevron icon --></svg>
    </button>
    <div data-slot="content">
      <p>Content for the second accordion item.</p>
    </div>
  </div>
</div>

<!-- Multiple selection accordion -->
<div data-component="accordion" data-type="multiple">
  <!-- Same item structure -->
</div>
```

---

## Navigation Components

### Dropdown Menu

```html
<!-- Basic dropdown -->
<div data-component="dropdown">
  <button data-slot="trigger" data-component="button" data-variant="outline">
    Options
    <svg data-slot="icon"><!-- chevron icon --></svg>
  </button>
  <div data-slot="content" data-align="end">
    <a href="#" data-slot="item">
      <svg data-slot="icon"><!-- icon --></svg>
      Profile
    </a>
    <a href="#" data-slot="item">
      <svg data-slot="icon"><!-- icon --></svg>
      Settings
    </a>
    <div data-slot="separator"></div>
    <button data-slot="item" data-variant="destructive">
      <svg data-slot="icon"><!-- icon --></svg>
      Logout
    </button>
  </div>
</div>

<!-- User menu dropdown -->
<div data-component="dropdown">
  <button data-slot="trigger" data-component="avatar">
    <img src="/user.jpg" alt="User">
  </button>
  <div data-slot="content" class="w-56">
    <div data-slot="header">
      <p class="font-semibold">John Doe</p>
      <p class="text-sm text-muted-foreground">john@example.com</p>
    </div>
    <div data-slot="separator"></div>
    <a href="/profile" data-slot="item">Profile</a>
    <a href="/billing" data-slot="item">Billing</a>
    <a href="/settings" data-slot="item">Settings</a>
    <div data-slot="separator"></div>
    <button data-slot="item">Log out</button>
  </div>
</div>
```

### Tabs

```html
<!-- Basic tabs -->
<div data-component="tabs" data-default-value="tab1">
  <div data-slot="list" role="tablist">
    <button data-slot="trigger" data-value="tab1" role="tab" aria-selected="true">
      Tab 1
    </button>
    <button data-slot="trigger" data-value="tab2" role="tab">
      Tab 2
    </button>
    <button data-slot="trigger" data-value="tab3" role="tab">
      Tab 3
    </button>
  </div>
  <div data-slot="content" data-value="tab1" role="tabpanel">
    <p>Content for tab 1</p>
  </div>
  <div data-slot="content" data-value="tab2" role="tabpanel" hidden>
    <p>Content for tab 2</p>
  </div>
  <div data-slot="content" data-value="tab3" role="tabpanel" hidden>
    <p>Content for tab 3</p>
  </div>
</div>

<!-- Tabs with icons -->
<div data-component="tabs">
  <div data-slot="list">
    <button data-slot="trigger" data-value="overview">
      <svg data-slot="icon"><!-- icon --></svg>
      Overview
    </button>
    <button data-slot="trigger" data-value="analytics">
      <svg data-slot="icon"><!-- icon --></svg>
      Analytics
    </button>
  </div>
</div>
```

### Command Palette

```html
<!-- Command palette -->
<div data-component="command" role="combobox" aria-expanded="false">
  <div data-slot="input-wrapper">
    <svg data-slot="icon"><!-- search icon --></svg>
    <input data-slot="input" type="text" placeholder="Type a command or search...">
  </div>
  <div data-slot="list" role="listbox">
    <div data-slot="group">
      <div data-slot="group-heading">Suggestions</div>
      <button data-slot="item" role="option">
        <svg data-slot="icon"><!-- icon --></svg>
        <span>Search Documentation</span>
        <kbd data-slot="shortcut">⌘K</kbd>
      </button>
      <button data-slot="item" role="option">
        <svg data-slot="icon"><!-- icon --></svg>
        <span>Create New Project</span>
      </button>
    </div>
    <div data-slot="separator"></div>
    <div data-slot="group">
      <div data-slot="group-heading">Recent</div>
      <button data-slot="item" role="option">
        <span>Project Alpha</span>
      </button>
    </div>
  </div>
  <div data-slot="empty">No results found</div>
</div>
```

### Command JavaScript

```javascript
// Initialize command palette
const command = new BasecoatCommand('command-palette');

// Open command palette
command.open();

// Set up keyboard shortcut
document.addEventListener('keydown', (e) => {
  if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
    e.preventDefault();
    command.toggle();
  }
});

// Listen for selection
command.on('select', (item) => {
  console.log('Selected:', item);
});
```

### Breadcrumb

```html
<!-- Basic breadcrumb -->
<nav data-component="breadcrumb" aria-label="Breadcrumb">
  <ol data-slot="list">
    <li data-slot="item">
      <a href="/">Home</a>
    </li>
    <li data-slot="separator">/</li>
    <li data-slot="item">
      <a href="/products">Products</a>
    </li>
    <li data-slot="separator">/</li>
    <li data-slot="item">
      <span aria-current="page">Product Details</span>
    </li>
  </ol>
</nav>

<!-- Breadcrumb with dropdown -->
<nav data-component="breadcrumb">
  <ol data-slot="list">
    <li data-slot="item">
      <a href="/">Home</a>
    </li>
    <li data-slot="separator">/</li>
    <li data-slot="item">
      <div data-component="dropdown">
        <button data-slot="trigger">
          ...
        </button>
        <div data-slot="content">
          <a href="/category1" data-slot="item">Category 1</a>
          <a href="/category2" data-slot="item">Category 2</a>
        </div>
      </div>
    </li>
    <li data-slot="separator">/</li>
    <li data-slot="item">
      <span aria-current="page">Current Page</span>
    </li>
  </ol>
</nav>
```

---

## Feedback Components

### Alert

```html
<!-- Default alert -->
<div data-component="alert">
  <svg data-slot="icon"><!-- info icon --></svg>
  <div data-slot="content">
    <h5 data-slot="title">Heads up!</h5>
    <p data-slot="description">You can add components to your app using the CLI.</p>
  </div>
</div>

<!-- Success alert -->
<div data-component="alert" data-variant="success">
  <svg data-slot="icon"><!-- check icon --></svg>
  <div data-slot="content">
    <h5 data-slot="title">Success!</h5>
    <p data-slot="description">Your changes have been saved.</p>
  </div>
</div>

<!-- Error alert -->
<div data-component="alert" data-variant="destructive">
  <svg data-slot="icon"><!-- alert icon --></svg>
  <div data-slot="content">
    <h5 data-slot="title">Error</h5>
    <p data-slot="description">There was a problem with your request.</p>
  </div>
</div>

<!-- Warning alert with action -->
<div data-component="alert" data-variant="warning">
  <svg data-slot="icon"><!-- warning icon --></svg>
  <div data-slot="content">
    <h5 data-slot="title">Warning</h5>
    <p data-slot="description">Your session will expire in 5 minutes.</p>
  </div>
  <button data-slot="action" data-component="button" data-variant="outline" data-size="sm">
    Extend Session
  </button>
</div>

<!-- Dismissible alert -->
<div data-component="alert" data-dismissible="true">
  <div data-slot="content">
    <p data-slot="description">This is a dismissible alert.</p>
  </div>
  <button data-slot="close" aria-label="Close">
    <svg><!-- close icon --></svg>
  </button>
</div>
```

### Toast Notifications

```html
<!-- Toast container (place once in layout) -->
<div data-component="toast-container" data-position="bottom-right">
  <!-- Toasts will be dynamically inserted here -->
</div>

<!-- Toast structure (for reference) -->
<div data-component="toast" role="status" aria-live="polite">
  <div data-slot="content">
    <h6 data-slot="title">Notification Title</h6>
    <p data-slot="description">This is a toast message.</p>
  </div>
  <button data-slot="close" aria-label="Close">
    <svg><!-- close icon --></svg>
  </button>
</div>
```

### Toast JavaScript

```javascript
// Initialize toast system
const toaster = new BasecoatToaster();

// Show basic toast
toaster.show({
  title: 'Success!',
  description: 'Your changes have been saved.',
  variant: 'success'
});

// Toast with action
toaster.show({
  title: 'New Message',
  description: 'You have a new message from John',
  action: {
    label: 'View',
    onClick: () => console.log('View clicked')
  }
});

// Toast with custom duration
toaster.show({
  title: 'Processing',
  description: 'This might take a few moments...',
  duration: 10000 // 10 seconds
});

// Dismiss all toasts
toaster.dismissAll();
```

### Badge

```html
<!-- Default badge -->
<span data-component="badge">Badge</span>

<!-- Badge variants -->
<span data-component="badge" data-variant="default">Default</span>
<span data-component="badge" data-variant="secondary">Secondary</span>
<span data-component="badge" data-variant="destructive">Destructive</span>
<span data-component="badge" data-variant="outline">Outline</span>
<span data-component="badge" data-variant="success">Success</span>
<span data-component="badge" data-variant="warning">Warning</span>

<!-- Badge with icon -->
<span data-component="badge">
  <svg data-slot="icon"><!-- icon --></svg>
  With Icon
</span>

<!-- Clickable badge -->
<button data-component="badge" data-variant="secondary">
  Clickable Badge
</button>

<!-- Badge in context -->
<h3>
  Inbox
  <span data-component="badge" data-variant="destructive" class="ml-2">99+</span>
</h3>
```

### Progress

```html
<!-- Basic progress bar -->
<div data-component="progress" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100">
  <div data-slot="bar" style="width: 60%"></div>
</div>

<!-- Progress with label -->
<div class="space-y-2">
  <div class="flex justify-between text-sm">
    <span>Uploading...</span>
    <span>60%</span>
  </div>
  <div data-component="progress" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100">
    <div data-slot="bar" style="width: 60%"></div>
  </div>
</div>

<!-- Indeterminate progress -->
<div data-component="progress" data-indeterminate="true">
  <div data-slot="bar"></div>
</div>

<!-- Circular progress (spinner) -->
<div data-component="spinner" role="status">
  <svg class="animate-spin h-4 w-4" viewBox="0 0 24 24">
    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
  </svg>
  <span class="sr-only">Loading...</span>
</div>
```

### Skeleton Loading

```html
<!-- Text skeleton -->
<div data-component="skeleton" class="h-4 w-[250px]"></div>
<div data-component="skeleton" class="h-4 w-[200px]"></div>

<!-- Card skeleton -->
<div data-component="card">
  <div data-slot="header">
    <div data-component="skeleton" class="h-6 w-[150px]"></div>
    <div data-component="skeleton" class="h-4 w-[100px] mt-2"></div>
  </div>
  <div data-slot="content" class="space-y-2">
    <div data-component="skeleton" class="h-4 w-full"></div>
    <div data-component="skeleton" class="h-4 w-full"></div>
    <div data-component="skeleton" class="h-4 w-[80%]"></div>
  </div>
</div>

<!-- Avatar skeleton -->
<div data-component="skeleton" class="h-12 w-12 rounded-full"></div>

<!-- Button skeleton -->
<div data-component="skeleton" class="h-10 w-[100px] rounded-md"></div>
```

---

## Data Attributes Reference

### Common Data Attributes

```html
<!-- Component identification -->
data-component="button"              <!-- Identifies the component type -->
data-variant="primary"              <!-- Visual variant -->
data-size="lg"                      <!-- Component size -->
data-state="active"                 <!-- Current state -->

<!-- Behavior modifiers -->
data-disabled="true"                <!-- Disabled state -->
data-loading="true"                 <!-- Loading state -->
data-error="true"                   <!-- Error state -->
data-required="true"                <!-- Required field -->

<!-- Layout and positioning -->
data-align="start|center|end"       <!-- Alignment -->
data-position="top|right|bottom|left" <!-- Position -->
data-side="top|right|bottom|left"   <!-- Side placement -->

<!-- Interactive states -->
data-open="true"                    <!-- Open state (dropdowns, dialogs) -->
data-selected="true"                <!-- Selected state -->
data-active="true"                  <!-- Active state -->
data-expanded="true"                <!-- Expanded state -->

<!-- Component slots -->
data-slot="trigger"                 <!-- Component part identifier -->
data-slot="content"                 <!-- Content area -->
data-slot="header"                  <!-- Header section -->
data-slot="footer"                  <!-- Footer section -->

<!-- Form-specific -->
data-value="option1"                <!-- Form value -->
data-name="fieldName"               <!-- Form field name -->
data-type="text|email|password"     <!-- Input type -->

<!-- Accessibility -->
aria-label="Description"            <!-- Accessible label -->
aria-describedby="id"              <!-- Description reference -->
aria-expanded="true|false"          <!-- Expanded state -->
aria-selected="true|false"          <!-- Selected state -->
role="button|tab|dialog"           <!-- ARIA role -->
```

---

## JavaScript Initialization

### Auto-initialization

```javascript
// Initialize all Basecoat components on page load
document.addEventListener('DOMContentLoaded', () => {
  Basecoat.init();
});

// Initialize components in a specific container
const container = document.getElementById('dynamic-content');
Basecoat.init(container);
```

### Manual Initialization

```javascript
// Initialize specific component types
Basecoat.initComponent('dropdown');
Basecoat.initComponent('dialog');
Basecoat.initComponent('tabs');

// Initialize individual component
const dropdown = new BasecoatDropdown(element);
const dialog = new BasecoatDialog('dialog-id');
const tabs = new BasecoatTabs(tabsElement);
```

### Component API Examples

```javascript
// Dropdown API
const dropdown = new BasecoatDropdown(element);
dropdown.open();
dropdown.close();
dropdown.toggle();
dropdown.on('open', () => console.log('Opened'));
dropdown.on('close', () => console.log('Closed'));

// Dialog API
const dialog = new BasecoatDialog('my-dialog');
dialog.open();
dialog.close();
dialog.on('open', () => console.log('Dialog opened'));
dialog.on('close', () => console.log('Dialog closed'));

// Tabs API
const tabs = new BasecoatTabs(element);
tabs.select('tab2');
tabs.on('change', (value) => console.log('Selected tab:', value));

// Toast API
const toaster = new BasecoatToaster();
const toast = toaster.show({
  title: 'Hello',
  description: 'This is a toast',
  duration: 5000
});
toast.dismiss();
```

### Event Handling

```javascript
// Listen for component events
document.addEventListener('basecoat:dropdown:open', (event) => {
  console.log('Dropdown opened:', event.detail.element);
});

document.addEventListener('basecoat:dialog:close', (event) => {
  console.log('Dialog closed:', event.detail.id);
});

// Custom event handling
element.addEventListener('basecoat:change', (event) => {
  console.log('Component changed:', event.detail);
});
```

---

## Styling and Customization

### CSS Custom Properties

```css
/* Override theme colors */
[data-component="button"][data-variant="primary"] {
  --button-bg: var(--primary);
  --button-text: var(--primary-foreground);
  --button-border: transparent;
}

/* Custom component styling */
[data-component="card"] {
  --card-padding: 1.5rem;
  --card-radius: var(--radius);
  --card-shadow: var(--shadow-sm);
}

/* State-based styling */
[data-component="input"][data-error="true"] {
  --input-border: var(--destructive);
  --input-focus-ring: var(--destructive);
}
```

### Utility Classes

```html
<!-- Combine with Tailwind utilities -->
<button data-component="button" class="w-full sm:w-auto">
  Full width on mobile
</button>

<!-- Custom spacing -->
<div data-component="card" class="mt-6 mb-4">
  <!-- Card content -->
</div>

<!-- Responsive behavior -->
<div data-component="alert" class="hidden sm:flex">
  <!-- Alert content -->
</div>
```

---

## Best Practices

1. **Accessibility First**: Always include proper ARIA attributes and keyboard support
2. **Semantic HTML**: Use appropriate HTML elements (button for actions, a for links)
3. **Progressive Enhancement**: Ensure components work without JavaScript
4. **Consistent Data Attributes**: Follow the established data attribute patterns
5. **Error Handling**: Always provide error states and helpful messages
6. **Loading States**: Show loading indicators for async operations
7. **Mobile First**: Design for mobile and enhance for larger screens
8. **Keyboard Navigation**: Ensure all interactive elements are keyboard accessible

---

## Troubleshooting

### Common Issues

```javascript
// Component not initializing
// Solution: Ensure DOM is loaded
document.addEventListener('DOMContentLoaded', () => {
  Basecoat.init();
});

// Dynamic content not working
// Solution: Re-initialize after adding content
fetch('/api/content').then(response => response.text()).then(html => {
  container.innerHTML = html;
  Basecoat.init(container);
});

// Event listeners not working
// Solution: Use event delegation
document.addEventListener('click', (e) => {
  const button = e.target.closest('[data-component="button"]');
  if (button) {
    // Handle button click
  }
});
```

### Debug Mode

```javascript
// Enable debug mode for verbose logging
Basecoat.debug = true;

// Check component initialization
console.log(Basecoat.components);

// Log all component events
Basecoat.on('*', (event) => {
  console.log('Basecoat event:', event);
});
```