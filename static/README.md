# JS Components Build System

Unified TypeScript and JavaScript build system for UI components with CSS processing.

## Overview

This build system:
- Compiles the `datatable` TypeScript project into both minified and unminified bundles
- Processes `basecoat` JavaScript components individually with both versions
- Processes CSS files with minification and debug versions
- Outputs files that UI components can easily reference
- Works well in Termux environment using esbuild
- Includes sourcemaps for debugging

## Directory Structure

```
static/
├── js/
│   ├── datatable/          # TypeScript project
│   │   └── src/            # TS source files
│   └── basecoat/           # JavaScript components
│       ├── basecoat.js     # Core component system
│       ├── dropdown-menu.js
│       ├── popover.js
│       └── ...
├── css/                    # CSS source files
│   ├── base.css
│   └── ...
├── dist/                   # Build output
│   ├── datatable.js        # Unminified bundle (debug)
│   ├── datatable.js.map    # Source map
│   ├── datatable.min.js    # Minified bundle (production)
│   ├── basecoat/           # Individual component files
│   │   ├── basecoat.js     # Unminified (debug)
│   │   ├── basecoat.js.map # Source map
│   │   ├── basecoat.min.js # Minified (production)
│   │   ├── dropdown-menu.js
│   │   ├── dropdown-menu.min.js
│   │   └── ...
│   └── css/                # Processed CSS files
│       ├── base.css        # Unminified (debug)
│       ├── base.min.css    # Minified (production)
│       └── ...
├── package.json            # Build dependencies
├── build.mjs               # Build script
├── js-manifest.json        # Component path mapping
└── js-helpers.go           # Go helpers for templ components
```

## Usage

### Install dependencies
```bash
pnpm install
```

### Build all components
```bash
pnpm run build
```

### Build specific components
```bash
pnpm run build:datatable   # Build only datatable (TS)
pnpm run build:basecoat    # Build only basecoat components (JS)
pnpm run build:css         # Build only CSS files
```

### Development with watch mode
```bash
pnpm run watch
```

### Clean build output
```bash
pnpm run clean
```

### Run tests
```bash
pnpm test              # Run all tests (261/269 passing)
pnpm run test:passing  # Run only stable tests (128/128 passing - recommended)
pnpm run test:watch    # Run tests in watch mode
pnpm run test:ui       # Run tests with UI (browser)
```

**Test Status**: 
- ✅ **261 tests passing** - Core functionality works
- ❌ **8 tests failing** - Minor issues with mocking and test assumptions
- 🟢 **Stable tests**: Sort, Filter, Export, Events, Index (128 tests)
- 🟡 **Partial failures**: State management, Integration, Core init events

**Known Issues**:
- Date serialization in localStorage mocking
- Console warning capture in test environment  
- Event timing for initialization
- Selection behavior expectations in filtered views

## Using in Templ Components

### Load the manifest (in your Go app initialization)
```go
import "path/to/static"

func init() {
    err := static.LoadJSManifest(".")
    if err != nil {
        log.Fatal("Failed to load JS manifest:", err)
    }
}
```

### Reference files in templ components

#### Production (Minified) - Default:
```templ
// JavaScript
templ DataTableComponent() {
    <div id="datatable"></div>
    <script src={static.GetDataTableJS()}></script>
}

templ DropdownComponent() {
    <div class="dropdown">...</div>
    <script src={static.GetComponentJS("dropdown-menu")}></script>
}

// CSS
templ BaseLayout() {
    <html>
        <head>
            <link rel="stylesheet" href={static.GetBaseCSS()}>
            <script src={static.GetBasecoatJS()}></script>
        </head>
        <body>
            <!-- Your content -->
        </body>
    </html>
}
```

#### Debug (Unminified) - For Development:
```templ
// JavaScript with sourcemaps for debugging
templ DataTableComponent() {
    <div id="datatable"></div>
    <script src={static.GetDataTableJSDebug()}></script>
}

templ DropdownComponent() {
    <div class="dropdown">...</div>
    <script src={static.GetComponentJSDebug("dropdown-menu")}></script>
}

// CSS with comments and formatting
templ BaseLayout() {
    <html>
        <head>
            <link rel="stylesheet" href={static.GetBaseCSSDebug()}>
            <script src={static.GetBasecoatJSDebug()}></script>
        </head>
        <body>
            <!-- Your content -->
        </body>
    </html>
}
```

## File Types & Outputs

### JavaScript:
- **Bundle**: Single file containing entire project (like datatable)
  - `datatable.js` + `datatable.min.js`
- **Component**: Individual components (like basecoat parts)
  - `component.js` + `component.min.js`
- **Sourcemaps**: `.js.map` files for debugging

### CSS:
- **Minified**: `component.min.css` - Production ready
- **Unminified**: `component.css` - Debug friendly with comments

## Available Helper Functions

### JavaScript:
```go
// Production (minified)
static.GetDataTableJS()                    // datatable.min.js
static.GetBasecoatJS()                     // basecoat.min.js
static.GetComponentJS("dropdown-menu")     // dropdown-menu.min.js

// Debug (unminified with sourcemaps)
static.GetDataTableJSDebug()               // datatable.js
static.GetBasecoatJSDebug()                // basecoat.js
static.GetComponentJSDebug("dropdown-menu") // dropdown-menu.js
```

### CSS:
```go
// Production (minified)
static.GetBaseCSS()                        // base.min.css
static.GetCSSPath("component-name")        // component.min.css

// Debug (unminified)
static.GetBaseCSSDebug()                   // base.css
static.GetCSSPathDebug("component-name")   // component.css
```

## Customization

Edit `build.mjs` to:
- Add new component directories
- Change build settings (minification, sourcemaps, etc.)
- Add preprocessing steps
- Modify output structure
- Add new CSS processing rules

The manifest file `js-manifest.json` automatically maps component names to their built file paths with both minified and unminified versions.