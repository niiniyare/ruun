# DataTable Organism

The DataTable organism is a comprehensive, enterprise-grade table component that provides advanced features for displaying, sorting, filtering, and managing tabular data. It follows Atomic Design principles and integrates seamlessly with the schema system for dynamic configuration.

## Table of Contents

1. [Overview](#overview)
2. [Architecture](#architecture)
3. [Features](#features)
4. [Basic Usage](#basic-usage)
5. [Advanced Usage](#advanced-usage)
6. [Schema Integration](#schema-integration)
7. [Configuration Options](#configuration-options)
8. [Styling & Theming](#styling--theming)
9. [Performance Considerations](#performance-considerations)
10. [Accessibility](#accessibility)
11. [API Reference](#api-reference)
12. [Examples](#examples)

## Overview

The DataTable organism provides:

- **Enterprise Features**: Advanced filtering, sorting, pagination, row selection, bulk actions
- **Schema Integration**: Automatic table generation from schema definitions
- **Performance**: Virtual scrolling for large datasets, server-side operations
- **Responsive Design**: Mobile-optimized layouts with stackable columns
- **HTMX Integration**: Real-time updates and server-side operations
- **Alpine.js Reactivity**: Client-side state management and interactivity
- **Accessibility**: Full ARIA compliance and keyboard navigation
- **Export Capabilities**: CSV, Excel, PDF export functionality

## Architecture

The DataTable follows a modular architecture:

```
organisms/
├── datatable.templ          # Main template component
├── datatable_logic.go       # Business logic & data processing
├── datatable_schema.go      # Schema-driven configuration
├── datatable_examples.go    # Usage examples
└── DATATABLE_README.md      # This documentation
```

### Key Components

1. **DataTable Template** (`datatable.templ`): Main UI component with all visual elements
2. **Business Logic** (`datatable_logic.go`): Data processing, filtering, sorting, pagination
3. **Schema Builder** (`datatable_schema.go`): Dynamic configuration from schema definitions
4. **Examples** (`datatable_examples.go`): Comprehensive usage examples

## Features

### Core Features

- ✅ **Column Management**: Sortable, resizable, reorderable columns
- ✅ **Row Selection**: Single and multi-row selection with bulk actions
- ✅ **Search & Filtering**: Global search, column-specific filters, advanced filtering
- ✅ **Pagination**: Client-side and server-side pagination
- ✅ **Sorting**: Single and multi-column sorting
- ✅ **Export**: CSV, Excel, PDF export with customizable options
- ✅ **Row Actions**: Contextual actions with dropdown menus
- ✅ **Responsive Design**: Mobile-friendly layouts

### Advanced Features

- ✅ **Virtual Scrolling**: Handle large datasets (10,000+ rows)
- ✅ **Row Expansion**: Expandable rows for additional details
- ✅ **Column Grouping**: Group related columns
- ✅ **Data Aggregation**: Sum, average, count, min, max calculations
- ✅ **Real-time Updates**: HTMX integration for live data
- ✅ **Custom Cell Renderers**: Badge, progress, image, link renderers
- ✅ **Conditional Formatting**: Dynamic styling based on data
- ✅ **State Persistence**: Remember user preferences

## Basic Usage

### Simple Table

```go
package main

import "github.com/niiniyare/ruun/views/components/organisms"

func CreateBasicTable() *organisms.DataTableProps {
    return &organisms.DataTableProps{
        ID:    "users-table",
        Title: "System Users",
        
        Columns: []organisms.DataTableColumn{
            organisms.TextColumn("name", "Name"),
            organisms.TextColumn("email", "Email"),
            organisms.DateColumn("created_at", "Created", "2006-01-02"),
        },
        
        Rows: []organisms.DataTableRow{
            {
                ID: "1",
                Data: map[string]any{
                    "name":       "John Doe",
                    "email":      "john@example.com",
                    "created_at": time.Now(),
                },
            },
        },
        
        Selectable: true,
        Sortable:   true,
        Search: organisms.DataTableSearch{
            Enabled: true,
        },
        Pagination: organisms.DataTablePagination{
            Enabled:  true,
            PageSize: 25,
        },
    }
}
```

### Template Usage

```templ
package main

templ UserManagementPage() {
    <div class="container mx-auto py-8">
        @organisms.DataTable(*CreateBasicTable())
    </div>
}
```

## Advanced Usage

### Using the Builder Pattern

```go
func CreateAdvancedTable() *organisms.DataTableProps {
    return organisms.NewDataTableBuilder("products", "Product Inventory").
        WithDescription("Manage product catalog").
        WithVariant(organisms.DataTableStriped).
        WithSize(organisms.DataTableSizeLG).
        AddColumn(organisms.DataTableColumn{
            Key:      "name",
            Title:    "Product Name",
            Type:     organisms.ColumnTypeText,
            Sortable: true,
            Clickable: true,
        }).
        AddColumn(organisms.DataTableColumn{
            Key:      "category",
            Title:    "Category",
            Type:     organisms.ColumnTypeBadge,
            BadgeMap: map[string]atoms.BadgeVariant{
                "Electronics": atoms.BadgePrimary,
                "Clothing":    atoms.BadgeSecondary,
            },
        }).
        AddColumn(organisms.DataTableColumn{
            Key:          "price",
            Title:        "Price",
            Type:         organisms.ColumnTypeCurrency,
            CurrencyCode: "USD",
            Precision:    2,
            Align:        "right",
        }).
        WithSearch(organisms.DataTableSearch{
            Enabled:     true,
            Placeholder: "Search products...",
            Columns:     []string{"name", "category"},
        }).
        WithActions([]organisms.DataTableAction{
            {
                ID:      "add",
                Text:    "Add Product",
                Icon:    "plus",
                Variant: atoms.ButtonPrimary,
                HXGet:   "/products/new",
            },
        }).
        WithBulkActions([]organisms.DataTableBulkAction{
            {
                ID:          "delete",
                Text:        "Delete Selected",
                Icon:        "trash",
                Variant:     atoms.ButtonDestructive,
                Destructive: true,
                Confirm:     true,
            },
        }).
        Build()
}
```

## Schema Integration

### Automatic Table Generation

The DataTable can automatically generate its configuration from a schema definition:

```go
func CreateSchemaTable() (*organisms.DataTableProps, error) {
    // Define schema
    userSchema := schema.NewSchema("users", schema.TypeForm, "Users")
    
    userSchema.AddField(schema.Field{
        Name:       "full_name",
        Type:       schema.FieldText,
        Label:      "Full Name",
        Sortable:   true,
        Searchable: true,
    })
    
    userSchema.AddField(schema.Field{
        Name:     "role",
        Type:     schema.FieldSelect,
        Label:    "Role",
        Options: []schema.FieldOption{
            {Value: "admin", Label: "Administrator"},
            {Value: "user", Label: "User"},
        },
    })
    
    // Build table from schema
    builder := organisms.NewDataTableSchemaBuilder(userSchema)
    
    // Customize specific columns
    builder.WithFieldMapping("role", organisms.ColumnMapping{
        ColumnType: organisms.ColumnTypeBadge,
        BadgeConfig: &organisms.BadgeConfig{
            VariantMap: map[string]string{
                "admin": "primary",
                "user":  "secondary",
            },
        },
    })
    
    return builder.Build(context.Background())
}
```

### Schema Configuration

```go
func ConfigureSchemaTable() *organisms.DataTableSchemaBuilder {
    builder := organisms.NewDataTableSchemaBuilder(schema)
    
    config := &organisms.DataTableConfig{
        EnableSelection:   true,
        EnableMultiSelect: true,
        EnableSorting:     true,
        EnableFiltering:   true,
        EnableSearch:      true,
        DefaultPageSize:   50,
        SearchPlaceholder: "Search users...",
        ExportFormats:     []organisms.ExportFormat{
            organisms.ExportCSV,
            organisms.ExportExcel,
        },
        Variant: organisms.DataTableDefault,
        Size:    organisms.DataTableSizeMD,
    }
    
    return builder.WithConfig(config)
}
```

## Configuration Options

### Column Types

The DataTable supports various column types with specific rendering:

```go
// Text column
organisms.TextColumn("name", "Name")

// Number column with precision
organisms.NumberColumn("price", "Price", 2)

// Date column with format
organisms.DateColumn("created_at", "Created", "2006-01-02")

// Badge column with variant mapping
organisms.BadgeColumn("status", "Status", map[string]atoms.BadgeVariant{
    "active":   atoms.BadgeSuccess,
    "inactive": atoms.BadgeSecondary,
})

// Custom column
organisms.DataTableColumn{
    Key:     "progress",
    Title:   "Completion",
    Type:    organisms.ColumnTypeProgress,
    Width:   "150px",
    Align:   "center",
}
```

### Search Configuration

```go
Search: organisms.DataTableSearch{
    Enabled:       true,
    Placeholder:   "Search...",
    Columns:       []string{"name", "email"}, // Specific columns
    CaseSensitive: false,
    MinLength:     2,
    Delay:         300, // Debounce delay in ms
    Highlight:     true, // Highlight matches
    ServerSide:    false, // Client or server search
    Advanced:      true, // Enable advanced search
}
```

### Pagination Configuration

```go
Pagination: organisms.DataTablePagination{
    Enabled:         true,
    CurrentPage:     1,
    PageSize:        25,
    TotalPages:      10,
    TotalItems:      250,
    PageSizeOptions: []int{10, 25, 50, 100},
    ShowTotal:       true,
    ShowPageSize:    true,
    ShowQuickJump:   true,
    ServerSide:      false,
    Compact:         false,
}
```

### Export Configuration

```go
Export: organisms.DataTableExport{
    Enabled:    true,
    Formats:    []organisms.ExportFormat{
        organisms.ExportCSV,
        organisms.ExportExcel,
        organisms.ExportPDF,
    },
    Filename:   "data_export",
    AllData:    true, // Export all data or just visible
    ServerSide: false,
}
```

## Styling & Theming

### Variants

```go
// Available variants
DataTableDefault   // Clean, minimal design
DataTableBordered  // With borders
DataTableStriped   // Alternating row colors
DataTableHover     // Hover effects
DataTableCompact   // Reduced padding
```

### Sizes

```go
// Available sizes
DataTableSizeSM    // Compact for mobile
DataTableSizeMD    // Standard size
DataTableSizeLG    // Large for desktop
```

### Density

```go
// Row density options
DataTableDensityComfortable // Standard spacing
DataTableDensityCompact     // Reduced spacing
DataTableDensityCondensed   // Minimal spacing
```

### CSS Classes

The component uses a consistent class naming convention:

```css
.datatable                    /* Main container */
.datatable-header            /* Title and controls */
.datatable-search            /* Search box */
.datatable-filters           /* Filter controls */
.datatable-table            /* Table element */
.datatable-th                /* Table headers */
.datatable-td                /* Table cells */
.datatable-pagination       /* Pagination controls */
.datatable-bulk-actions      /* Bulk action bar */
```

## Performance Considerations

### Large Datasets

For tables with large datasets (1000+ rows), consider:

1. **Server-side Operations**: Enable server-side pagination, sorting, and filtering
2. **Virtual Scrolling**: Use virtualization for smooth performance
3. **Data Chunking**: Load data in chunks
4. **Debounced Search**: Use search delays to reduce API calls

```go
// Performance configuration
VirtualScrolling: true,
LazyLoading:     true,
CacheData:       true,

// Virtual scroll options
VirtualScrollOptions: map[string]any{
    "rowHeight":    40,
    "overscan":     10,
    "enableResize": true,
},

// Server-side operations
Search: organisms.DataTableSearch{
    ServerSide: true,
    Delay:      500, // Longer delay for server calls
},

Pagination: organisms.DataTablePagination{
    ServerSide: true,
    PageSize:   100, // Larger page size
},
```

### Memory Optimization

- Use `LazyLoading` for data that's not immediately visible
- Enable `CacheData` for frequently accessed data
- Implement proper cleanup in Alpine.js components

## Accessibility

The DataTable is fully accessible with:

### ARIA Labels

```go
AriaLabels: map[string]string{
    "table":     "Data table",
    "search":    "Search table data",
    "filter":    "Filter table data",
    "sort":      "Sort column",
    "select":    "Select row",
    "selectAll": "Select all rows",
}
```

### Keyboard Navigation

- **Tab**: Navigate through interactive elements
- **Space**: Select rows, activate buttons
- **Enter**: Activate buttons, links
- **Arrow Keys**: Navigate table cells
- **Escape**: Close dropdowns, cancel actions

### Screen Reader Support

- Proper table structure with `<thead>`, `<tbody>`, `<tfoot>`
- Column headers with scope attributes
- Row and cell ARIA labels
- Live regions for dynamic content updates

## API Reference

### DataTableProps

Main configuration object for the DataTable component.

```go
type DataTableProps struct {
    // Basic configuration
    ID          string
    Title       string
    Description string
    Variant     DataTableVariant
    Size        DataTableSize
    Density     DataTableDensity
    
    // Data
    Columns     []DataTableColumn
    Rows        []DataTableRow
    
    // Features
    Selectable  bool
    MultiSelect bool
    Sortable    bool
    Filterable  bool
    Resizable   bool
    Expandable  bool
    
    // Search and filtering
    Search      DataTableSearch
    Filters     []DataTableFilter
    
    // Pagination
    Pagination  DataTablePagination
    
    // Actions
    Actions     []DataTableAction
    BulkActions []DataTableBulkAction
    RowActions  []DataTableAction
    
    // Export
    Export      DataTableExport
    
    // Performance
    Virtualized bool
    LazyLoad    bool
    CacheData   bool
    
    // HTMX integration
    HXGet       string
    HXPost      string
    HXTarget    string
    HXSwap      string
    
    // Event handlers
    OnRowClick  string
    OnSort      string
    OnFilter    string
    OnSearch    string
}
```

### DataTableColumn

Column configuration with rendering options.

```go
type DataTableColumn struct {
    Key           string
    Title         string
    Type          ColumnType
    Width         string
    Sortable      bool
    Searchable    bool
    Filterable    bool
    Visible       bool
    Align         string
    Format        string
    BadgeMap      map[string]atoms.BadgeVariant
    ActionItems   []molecules.MenuItemProps
}
```

### DataTableRow

Row data with metadata and actions.

```go
type DataTableRow struct {
    ID       string
    Data     map[string]any
    Selected bool
    Expanded bool
    Disabled bool
    Class    string
    Actions  []molecules.MenuItemProps
    Meta     map[string]any
}
```

## Examples

### Basic Table

```go
examples := &organisms.DataTableExamples{}
basicTable := examples.BasicExample()
```

### Advanced Features

```go
advancedTable := examples.AdvancedExample()
```

### Schema-Driven

```go
schemaTable, err := examples.SchemaExample()
if err != nil {
    // Handle error
}
```

### Builder Pattern

```go
builderTable := examples.BuilderExample()
```

### Responsive Design

```go
mobileTable := examples.ResponsiveExample()
```

### Virtual Scrolling

```go
largeDataTable := examples.VirtualizedExample()
```

### Custom Rendering

```go
customTable := examples.CustomRenderingExample()
```

### Grouping & Aggregation

```go
groupedTable := examples.GroupingExample()
```

## Integration with HTMX

The DataTable integrates seamlessly with HTMX for server-side operations:

```html
<!-- Auto-refresh every 30 seconds -->
<div hx-get="/api/data" hx-trigger="every 30s" hx-target="#data-table">
    @organisms.DataTable(props)
</div>

<!-- Search with debouncing -->
<input 
    hx-get="/api/search" 
    hx-trigger="keyup changed delay:500ms" 
    hx-target="#table-body"
    hx-include="[name='filters']"
>
```

### Server-Side Endpoints

Expected query parameters for server-side operations:

```
GET /api/data?page=1&size=25&sort=name&direction=asc&search=query&filter[status]=active
```

Expected response format:

```json
{
    "rows": [...],
    "total": 1000,
    "page": 1,
    "totalPages": 40
}
```

## Best Practices

1. **Performance**: Use server-side operations for large datasets
2. **UX**: Provide clear loading states and error messages
3. **Accessibility**: Test with screen readers and keyboard navigation
4. **Responsive**: Design mobile-first with progressive enhancement
5. **State Management**: Persist user preferences (sort, filters, page size)
6. **Error Handling**: Gracefully handle network failures and invalid data
7. **Security**: Validate and sanitize all user inputs
8. **Testing**: Include unit tests for business logic and integration tests for UI

## Migration Guide

When upgrading from the basic Table component:

1. Update import statements
2. Replace `TableProps` with `DataTableProps`
3. Update column definitions to use new column types
4. Add new features gradually (search, pagination, etc.)
5. Test thoroughly with your existing data

## Contributing

To contribute to the DataTable component:

1. Follow the existing code patterns
2. Add comprehensive tests
3. Update documentation
4. Include usage examples
5. Ensure accessibility compliance

The DataTable organism provides a solid foundation for enterprise data management with room for customization and extension based on specific requirements.