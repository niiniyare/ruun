# Progressive Enhancement Guide for Organisms

## Overview

This guide documents the progressive enhancement pattern implemented for the Navigation and DataTable organisms, based on the proven architecture from the Form organism. This pattern allows components to scale from simple functionality to enterprise-grade features while maintaining performance and usability.

## Architecture Philosophy

### 1. Progressive Enhancement Principle
- **Core Functionality First**: Start with essential features that work without JavaScript
- **Feature Flags**: Use nil pointer checks to enable/disable advanced features
- **Zero Configuration**: Components work with minimal props, add complexity as needed
- **Performance Optimized**: Only load JavaScript and features that are actually configured

### 2. Component Structure

```go
type ComponentProps struct {
    // Core properties (always required)
    ID    string
    Items []Item    // Essential data

    // Basic Enhancement (optional)
    Title       string
    Description string
    ClassName   string

    // Progressive Enhancement (nil = disabled)
    Feature1 *FeatureConfig `json:"feature1"`
    Feature2 *FeatureConfig `json:"feature2"`
    Advanced *AdvancedConfig `json:"advanced"`
}
```

### 3. Feature Detection Pattern

```go
// Feature detection functions
func hasFeature1(props ComponentProps) bool {
    return props.Feature1 != nil && props.Feature1.Enabled
}

func hasFeature2(props ComponentProps) bool {
    return props.Feature2 != nil && props.Feature2.Enabled
}
```

## Navigation Organism Refactoring

### Before (Monolithic)
```go
type NavigationProps struct {
    // 50+ properties mixed together
    Type        NavigationType
    Collapsible bool
    Search      NavigationSearch  // Always included
    UserAvatar  string
    MobileMenu  NavigationMobileMenu
    // ... many more
}
```

### After (Progressive Enhancement)
```go
type NavigationProps struct {
    // Core properties (always required)
    ID    string             `json:"id"`
    Type  NavigationType     `json:"type"`
    Items []NavigationItem   `json:"items"`
    
    // Basic Enhancement (optional)
    Title       string         `json:"title"`
    Logo        string         `json:"logo"`
    
    // Progressive Enhancement (nil = disabled)
    Layout     *LayoutConfig     `json:"layout"`     // Collapsible, responsive
    Search     *SearchConfig     `json:"search"`     // Search functionality
    User       *UserConfig       `json:"user"`       // Authentication/profile
    Mobile     *MobileConfig     `json:"mobile"`     // Mobile features
    Theme      *ThemeConfig      `json:"theme"`      // Theme switching
    Analytics  *AnalyticsConfig  `json:"analytics"`  // Usage tracking
}
```

### Navigation Usage Examples

#### 1. Simple Navigation
```go
@organisms.Navigation(organisms.NavigationProps{
    ID:   "main-nav",
    Type: organisms.NavigationSidebar,
    Items: []organisms.NavigationItem{
        {Text: "Dashboard", URL: "/dashboard", Icon: "home"},
        {Text: "Settings", URL: "/settings", Icon: "settings"},
    },
})
```

#### 2. Enhanced Navigation
```go
@organisms.Navigation(organisms.NavigationProps{
    ID:    "main-nav",
    Type:  organisms.NavigationSidebar,
    Title: "My App",
    Logo:  "/logo.svg",
    Items: navItems,
    
    // Progressive Enhancement: Add search
    Search: &organisms.SearchConfig{
        Enabled:     true,
        Placeholder: "Search navigation...",
        Shortcut:    "/",
    },
    
    // Progressive Enhancement: Add user profile
    User: &organisms.UserConfig{
        Enabled: true,
        Name:    "John Doe",
        Avatar:  "/avatar.jpg",
        Menu:    userMenuItems,
    },
})
```

#### 3. Enterprise Navigation
```go
@organisms.Navigation(organisms.NavigationProps{
    ID:    "enterprise-nav",
    Type:  organisms.NavigationSidebar,
    Title: "Enterprise Dashboard",
    Items: complexNavItems,
    
    // All progressive features enabled
    Layout: &organisms.LayoutConfig{
        Collapsible: true,
        Responsive:  true,
        Position:    "fixed",
    },
    
    Search: &organisms.SearchConfig{
        Enabled:   true,
        URL:       "/api/search/navigation",
        Highlight: true,
    },
    
    User: &organisms.UserConfig{
        Enabled:    true,
        ShowStatus: true,
        // ... user config
    },
    
    Mobile: &organisms.MobileConfig{
        Enabled:   true,
        Overlay:   true,
        Swipeable: true,
    },
    
    Theme: &organisms.ThemeConfig{
        AutoDetect: true,
        DarkMode:   false,
    },
    
    Analytics: &organisms.AnalyticsConfig{
        Enabled:     true,
        TrackClicks: true,
        TrackSearch: true,
    },
})
```

## DataTable Organism Refactoring

### Before (Monolithic)
```go
type DataTableProps struct {
    // 40+ properties mixed together
    Selectable   bool
    Sortable     bool
    Filterable   bool
    Search       DataTableSearch
    Pagination   DataTablePagination
    Export       DataTableExport
    Virtualized  bool
    // ... many more
}
```

### After (Progressive Enhancement)
```go
type DataTableProps struct {
    // Core properties (always required)
    ID      string              `json:"id"`
    Columns []DataTableColumn   `json:"columns"`
    Rows    []DataTableRow      `json:"rows"`
    
    // Basic Enhancement (optional)
    Title       string           `json:"title"`
    Description string           `json:"description"`
    
    // Progressive Enhancement (nil = disabled)
    Selection   *SelectionConfig   `json:"selection"`   // Row selection
    Sorting     *SortingConfig     `json:"sorting"`     // Column sorting
    Filtering   *FilteringConfig   `json:"filtering"`   // Search & filters
    Pagination  *PaginationConfig  `json:"pagination"`  // Pagination
    Actions     *ActionsConfig     `json:"actions"`     // Table/row actions
    Export      *ExportConfig      `json:"export"`      // Export functionality
    Performance *PerformanceConfig `json:"performance"` // Virtualization
    Analytics   *AnalyticsConfig   `json:"analytics"`   // Usage tracking
}
```

### DataTable Usage Examples

#### 1. Simple Table
```go
@organisms.DataTable(organisms.DataTableProps{
    ID: "users-table",
    Columns: []organisms.DataTableColumn{
        {Key: "name", Title: "Name", Type: organisms.ColumnTypeText},
        {Key: "email", Title: "Email", Type: organisms.ColumnTypeText},
    },
    Rows: userRows,
})
```

#### 2. Enhanced Table
```go
@organisms.DataTable(organisms.DataTableProps{
    ID:          "products-table",
    Title:       "Product Inventory",
    Description: "Manage your product catalog",
    Columns:     productColumns,
    Rows:        productRows,
    
    // Progressive Enhancement: Add sorting
    Sorting: &organisms.SortingConfig{
        Enabled:     true,
        MultiColumn: true,
        DefaultSort: "name:asc",
    },
    
    // Progressive Enhancement: Add pagination
    Pagination: &organisms.PaginationConfig{
        Enabled:  true,
        PageSize: 25,
        ShowTotal: true,
    },
})
```

#### 3. Enterprise Table
```go
@organisms.DataTable(organisms.DataTableProps{
    ID:      "enterprise-data",
    Columns: complexColumns,
    Rows:    enterpriseRows,
    
    // All progressive features enabled
    Selection: &organisms.SelectionConfig{
        Enabled:     true,
        MultiSelect: true,
        Checkbox:    true,
    },
    
    Sorting: &organisms.SortingConfig{
        Enabled:    true,
        ServerSide: true,
        URL:        "/api/sort",
    },
    
    Filtering: &organisms.FilteringConfig{
        Enabled: true,
        Search: &organisms.SearchConfig{
            Enabled: true,
            URL:     "/api/search",
        },
        ColumnFilters: true,
        Advanced:      true,
    },
    
    Pagination: &organisms.PaginationConfig{
        Enabled:    true,
        ServerSide: true,
        URL:        "/api/paginate",
    },
    
    Actions: &organisms.ActionsConfig{
        Enabled: true,
        BulkActions: []organisms.DataTableBulkAction{
            {Text: "Delete Selected", Icon: "trash", Destructive: true},
        },
    },
    
    Export: &organisms.ExportConfig{
        Enabled: true,
        Formats: []organisms.ExportFormat{
            organisms.ExportCSV,
            organisms.ExportExcel,
            organisms.ExportPDF,
        },
    },
    
    Performance: &organisms.PerformanceConfig{
        Virtualized: true,
        LazyLoad:    true,
        CacheData:   true,
    },
    
    Analytics: &organisms.AnalyticsConfig{
        Enabled:        true,
        TrackSort:      true,
        TrackFilter:    true,
        TrackSelection: true,
    },
})
```

## Implementation Benefits

### 1. Performance Optimizations
- **Conditional JavaScript**: Only load features that are enabled
- **Feature Detection**: Automatic based on nil pointer checks
- **Zero Overhead**: Disabled features add no computational cost
- **Progressive Loading**: Features can be lazy-loaded

```go
// Alpine.js state is built dynamically based on enabled features
func buildNavigationAlpineData(props NavigationProps) string {
    data := "{ sidebarCollapsed: false"
    
    if hasSearch(props) {
        data += ", searchQuery: '', searchOpen: false"
    }
    
    if hasTheme(props) {
        data += ", darkMode: " + strconv.FormatBool(props.Theme.DarkMode)
    }
    
    data += " }"
    return data
}
```

### 2. Developer Experience
- **Type Safety**: All configurations are strongly typed
- **IntelliSense**: IDE auto-completion for all options
- **Zero Config**: Components work with minimal setup
- **Incremental Adoption**: Add features as needed

### 3. Maintainability
- **Clear Separation**: Features are isolated in separate configs
- **Single Responsibility**: Each config handles one concern
- **Easy Testing**: Features can be tested independently
- **Documentation**: Self-documenting through types

## Template Structure Pattern

### 1. Feature Detection in Templates
```go
// Navigation example
templ navigationSidebar(props NavigationProps) {
    <div class="navigation-sidebar-container">
        <!-- Core Content Always Rendered -->
        <div class="navigation-menu">
            for _, item := range props.Items {
                @navigationMenuItem(item, 0)
            }
        </div>
        
        <!-- Progressive Enhancement: Search -->
        if hasSearch(props) {
            @navigationSearch(props.Search)
        }
        
        <!-- Progressive Enhancement: User Section -->
        if hasUser(props) {
            @navigationUserSection(props)
        }
    </div>
}
```

### 2. Alpine.js State Management
```javascript
// Dynamically generated based on enabled features
{
    // Core state (always present)
    loading: false,
    
    // Progressive Enhancement: Search (only if enabled)
    searchQuery: '',
    searchOpen: false,
    
    // Progressive Enhancement: Selection (only if enabled)
    selectedRows: [],
    selectAll: false,
    
    // Methods are also conditionally added
    toggleSearch() { /* only if search enabled */ },
    toggleSelection(id) { /* only if selection enabled */ }
}
```

### 3. CSS Classes
```go
func getNavigationClasses(props NavigationProps) string {
    return utils.TwMerge(
        "navigation",
        fmt.Sprintf("navigation-%s", string(props.Type)),
        utils.If(props.Search != nil && props.Search.Enabled, "navigation-searchable"),
        utils.If(props.Mobile != nil && props.Mobile.Enabled, "navigation-mobile"),
        utils.If(props.Theme != nil && props.Theme.DarkMode, "navigation-dark"),
        props.ClassName,
    )
}
```

## Configuration Patterns

### 1. Basic Configuration
```go
// Minimal setup - just the essentials
BasicConfig{
    Enabled: true,
}
```

### 2. Standard Configuration
```go
// Common use cases with sensible defaults
StandardConfig{
    Enabled:     true,
    Placeholder: "Search...",
    MinLength:   2,
    Debounce:    300,
}
```

### 3. Advanced Configuration
```go
// Enterprise features with full customization
AdvancedConfig{
    Enabled:       true,
    URL:           "/api/search",
    Highlight:     true,
    ServerSide:    true,
    CacheResults:  true,
    Analytics:     true,
}
```

## Best Practices

### 1. Start Simple, Enhance Progressively
```go
// Start with this
@Navigation(NavigationProps{
    ID:   "nav",
    Type: NavigationSidebar,
    Items: navItems,
})

// Then add features as needed
@Navigation(NavigationProps{
    ID:    "nav",
    Type:  NavigationSidebar,
    Items: navItems,
    Search: &SearchConfig{Enabled: true}, // Add search
})
```

### 2. Use Feature Detection
```go
// In templates
if hasSearch(props) {
    @searchComponent(props.Search)
}

// In functions
func buildClasses(props Props) string {
    return utils.TwMerge(
        "base-class",
        utils.If(hasFeature(props), "feature-class"),
    )
}
```

### 3. Type Safety
```go
// Use typed enums, not strings
Search: &SearchConfig{
    Strategy: SearchStrategyServerSide, // ✅ Type-safe
    // Strategy: "server-side",         // ❌ Error-prone
}
```

### 4. Configuration Validation
```go
// Validate configurations at creation
func NewSearchConfig(enabled bool, url string) *SearchConfig {
    if enabled && url == "" {
        log.Warn("Search enabled but no URL provided, falling back to client-side")
    }
    
    return &SearchConfig{
        Enabled: enabled,
        URL:     url,
        MinLength: utils.IfElse(enabled, 2, 0),
    }
}
```

## Migration Guide

### From Monolithic to Progressive Enhancement

#### 1. Identify Feature Groups
```go
// Old monolithic structure
type OldProps struct {
    Searchable    bool
    SearchURL     string
    SearchDelay   int
    Sortable      bool
    SortURL       string
    MultiSort     bool
    // ... 50+ more properties
}

// New progressive structure
type NewProps struct {
    // Core
    ID   string
    Data []Item
    
    // Progressive Enhancement
    Search  *SearchConfig
    Sorting *SortingConfig
}
```

#### 2. Create Feature Configs
```go
// Group related properties together
type SearchConfig struct {
    Enabled     bool   `json:"enabled"`
    URL         string `json:"url"`
    Placeholder string `json:"placeholder"`
    MinLength   int    `json:"minLength"`
    Debounce    int    `json:"debounce"`
}
```

#### 3. Update Templates
```go
// Old approach
if props.Searchable {
    // render search
}

// New approach
if hasSearch(props) {
    @searchComponent(props.Search)
}
```

#### 4. Migration Helper
```go
// Provide migration helper for backward compatibility
func MigrateFromOldProps(old OldProps) NewProps {
    props := NewProps{
        ID:   old.ID,
        Data: old.Data,
    }
    
    if old.Searchable {
        props.Search = &SearchConfig{
            Enabled: true,
            URL:     old.SearchURL,
            Debounce: old.SearchDelay,
        }
    }
    
    return props
}
```

## Analytics Integration

### 1. Feature Usage Tracking
```javascript
// Automatically track feature adoption
if (hasSearch(props)) {
    gtag('event', 'feature_enabled', {
        feature: 'search',
        component: 'navigation'
    });
}
```

### 2. Performance Monitoring
```javascript
// Track feature impact on performance
const startTime = performance.now();
// Feature initialization
const endTime = performance.now();

gtag('event', 'feature_performance', {
    feature: 'search',
    initialization_time: endTime - startTime
});
```

## Conclusion

The progressive enhancement pattern provides a scalable, maintainable, and performant approach to building complex UI organisms. By starting with core functionality and layering on advanced features through configuration, we achieve:

1. **Better Performance** - Only load what's needed
2. **Improved DX** - Clear, typed configurations
3. **Easier Maintenance** - Isolated, testable features
4. **Future-Proof** - Easy to add new capabilities

This pattern has been successfully applied to Navigation and DataTable organisms, and can be extended to other complex components in the system.