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