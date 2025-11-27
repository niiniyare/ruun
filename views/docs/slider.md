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