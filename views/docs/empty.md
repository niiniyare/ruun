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