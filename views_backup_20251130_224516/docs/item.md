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