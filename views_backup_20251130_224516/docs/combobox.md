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