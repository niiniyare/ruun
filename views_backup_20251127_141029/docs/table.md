# Table Component

A responsive table component for displaying tabular data with proper styling and structure.

## Basic Usage

```html
<div class="overflow-x-auto">
  <table class="table">
    <caption>A list of your recent invoices.</caption>
    <thead>
      <tr>
        <th>Invoice</th>
        <th>Status</th>
        <th>Method</th>
        <th>Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">INV001</td>
        <td>Paid</td>
        <td>Credit Card</td>
        <td class="text-right">$250.00</td>
      </tr>
      <tr>
        <td class="font-medium">INV002</td>
        <td>Pending</td>
        <td>PayPal</td>
        <td class="text-right">$150.00</td>
      </tr>
    </tbody>
    <tfoot>
      <tr>
        <td colspan="3">Total</td>
        <td class="text-right">$400.00</td>
      </tr>
    </tfoot>
  </table>
</div>
```

## CSS Classes

### Primary Classes
- **`table`** - Applied to the `<table>` element

### Supporting Classes
- **`overflow-x-auto`** - Container for responsive horizontal scrolling
- Text alignment classes (`text-left`, `text-center`, `text-right`)
- Font weight classes (`font-medium`, `font-semibold`)

### Tailwind Utilities Used
- `font-medium` - Medium font weight for headers/important data
- `text-right` - Right-align numerical data
- `text-center` - Center-align data
- `overflow-x-auto` - Horizontal scroll for responsive tables

## Component Attributes

### Table Attributes
| Attribute | Type | Description | Required |
|-----------|------|-------------|----------|
| `class` | string | Must include "table" | Yes |

### No JavaScript Required
This component is purely CSS-based and does not require JavaScript initialization.

## HTML Structure

```html
<div class="overflow-x-auto">
  <table class="table">
    <caption>Table description (optional)</caption>
    <thead>
      <tr>
        <th>Header 1</th>
        <th>Header 2</th>
        <th>Header 3</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>Data 1</td>
        <td>Data 2</td>
        <td>Data 3</td>
      </tr>
    </tbody>
    <tfoot>
      <tr>
        <td colspan="2">Footer</td>
        <td>Total</td>
      </tr>
    </tfoot>
  </table>
</div>
```

## Examples

### Basic Data Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Role</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td>Admin</td>
        <td>Active</td>
      </tr>
      <tr>
        <td class="font-medium">Jane Smith</td>
        <td>jane@example.com</td>
        <td>Editor</td>
        <td>Active</td>
      </tr>
      <tr>
        <td class="font-medium">Mike Johnson</td>
        <td>mike@example.com</td>
        <td>Viewer</td>
        <td>Inactive</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Actions
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Product</th>
        <th>SKU</th>
        <th>Price</th>
        <th>Stock</th>
        <th>Actions</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">Wireless Headphones</td>
        <td>WH-001</td>
        <td class="text-right">$99.99</td>
        <td class="text-center">25</td>
        <td>
          <div class="flex gap-2">
            <button type="button" class="btn-ghost text-sm">Edit</button>
            <button type="button" class="btn-ghost text-sm text-destructive">Delete</button>
          </div>
        </td>
      </tr>
      <tr>
        <td class="font-medium">Bluetooth Speaker</td>
        <td>BS-002</td>
        <td class="text-right">$149.99</td>
        <td class="text-center">12</td>
        <td>
          <div class="flex gap-2">
            <button type="button" class="btn-ghost text-sm">Edit</button>
            <button type="button" class="btn-ghost text-sm text-destructive">Delete</button>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Status Badges
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Order ID</th>
        <th>Customer</th>
        <th>Date</th>
        <th>Status</th>
        <th>Amount</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">#ORD-001</td>
        <td>Alice Johnson</td>
        <td>2024-01-15</td>
        <td>
          <span class="badge bg-success text-success-foreground">Completed</span>
        </td>
        <td class="text-right">$299.99</td>
      </tr>
      <tr>
        <td class="font-medium">#ORD-002</td>
        <td>Bob Wilson</td>
        <td>2024-01-14</td>
        <td>
          <span class="badge bg-warning text-warning-foreground">Pending</span>
        </td>
        <td class="text-right">$149.50</td>
      </tr>
      <tr>
        <td class="font-medium">#ORD-003</td>
        <td>Carol Davis</td>
        <td>2024-01-13</td>
        <td>
          <span class="badge bg-error text-error-foreground">Cancelled</span>
        </td>
        <td class="text-right">$89.99</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Avatars
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>User</th>
        <th>Department</th>
        <th>Role</th>
        <th>Last Active</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>
          <div class="flex items-center gap-3">
            <img src="/avatar1.jpg" alt="John Doe" class="size-8 rounded-full">
            <div>
              <div class="font-medium">John Doe</div>
              <div class="text-sm text-muted-foreground">john@example.com</div>
            </div>
          </div>
        </td>
        <td>Engineering</td>
        <td>Senior Developer</td>
        <td class="text-muted-foreground">2 minutes ago</td>
      </tr>
      <tr>
        <td>
          <div class="flex items-center gap-3">
            <img src="/avatar2.jpg" alt="Jane Smith" class="size-8 rounded-full">
            <div>
              <div class="font-medium">Jane Smith</div>
              <div class="text-sm text-muted-foreground">jane@example.com</div>
            </div>
          </div>
        </td>
        <td>Design</td>
        <td>UX Designer</td>
        <td class="text-muted-foreground">1 hour ago</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Sortable Table Headers
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Name
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Date
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
        <th>
          <button type="button" class="flex items-center gap-2 font-medium hover:text-foreground">
            Amount
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="m7 15 5 5 5-5" />
              <path d="m7 9 5-5 5 5" />
            </svg>
          </button>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">Transaction 1</td>
        <td>2024-01-15</td>
        <td class="text-right">$250.00</td>
      </tr>
      <tr>
        <td class="font-medium">Transaction 2</td>
        <td>2024-01-14</td>
        <td class="text-right">$175.50</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Table with Selection
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>
          <input type="checkbox" class="checkbox" aria-label="Select all">
        </th>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td>Active</td>
      </tr>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">Jane Smith</td>
        <td>jane@example.com</td>
        <td>Active</td>
      </tr>
      <tr>
        <td>
          <input type="checkbox" class="checkbox" aria-label="Select row">
        </td>
        <td class="font-medium">Mike Johnson</td>
        <td>mike@example.com</td>
        <td>Inactive</td>
      </tr>
    </tbody>
  </table>
</div>
```

### Empty State Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Role</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="4" class="text-center py-8">
          <div class="flex flex-col items-center gap-2">
            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round" class="text-muted-foreground">
              <rect width="7" height="7" x="3" y="3" rx="1" />
              <rect width="7" height="7" x="14" y="3" rx="1" />
              <rect width="7" height="7" x="14" y="14" rx="1" />
              <rect width="7" height="7" x="3" y="14" rx="1" />
            </svg>
            <div class="text-muted-foreground">No data found</div>
            <p class="text-sm text-muted-foreground">Get started by adding your first entry.</p>
          </div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Compact Table
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th class="py-2">Product</th>
        <th class="py-2">Price</th>
        <th class="py-2">Stock</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="py-2 font-medium">Item 1</td>
        <td class="py-2 text-right">$99.99</td>
        <td class="py-2 text-center">25</td>
      </tr>
      <tr>
        <td class="py-2 font-medium">Item 2</td>
        <td class="py-2 text-right">$149.99</td>
        <td class="py-2 text-center">12</td>
      </tr>
    </tbody>
  </table>
</div>
```

## Responsive Design

### Horizontal Scroll
```html
<div class="overflow-x-auto">
  <table class="table min-w-full">
    <!-- Table content -->
  </table>
</div>
```

### Stack on Mobile
```html
<div class="block md:hidden">
  <!-- Mobile card view -->
  <div class="space-y-4">
    <div class="border rounded-lg p-4">
      <div class="font-medium">John Doe</div>
      <div class="text-sm text-muted-foreground">john@example.com</div>
      <div class="mt-2 flex justify-between">
        <span>Role: Admin</span>
        <span class="badge">Active</span>
      </div>
    </div>
  </div>
</div>

<div class="hidden md:block">
  <!-- Desktop table view -->
  <div class="overflow-x-auto">
    <table class="table">
      <!-- Full table -->
    </table>
  </div>
</div>
```

### Responsive Columns
```html
<div class="overflow-x-auto">
  <table class="table">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th class="hidden md:table-cell">Phone</th>
        <th class="hidden lg:table-cell">Department</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td class="font-medium">John Doe</td>
        <td>john@example.com</td>
        <td class="hidden md:table-cell">(555) 123-4567</td>
        <td class="hidden lg:table-cell">Engineering</td>
        <td>Active</td>
      </tr>
    </tbody>
  </table>
</div>
```

## Accessibility Features

- **Semantic HTML**: Uses proper table markup
- **Screen Reader Support**: Table headers properly associated with data
- **Keyboard Navigation**: Focusable interactive elements
- **ARIA Labels**: Appropriate labels for complex tables

### Enhanced Accessibility
```html
<div class="overflow-x-auto">
  <table class="table" role="table" aria-label="User management data">
    <caption class="sr-only">
      List of users with their roles and status information
    </caption>
    <thead>
      <tr role="row">
        <th role="columnheader" scope="col" aria-sort="none">
          Name
        </th>
        <th role="columnheader" scope="col" aria-sort="none">
          Role
        </th>
        <th role="columnheader" scope="col" aria-sort="ascending">
          Status
        </th>
      </tr>
    </thead>
    <tbody>
      <tr role="row">
        <td role="gridcell">John Doe</td>
        <td role="gridcell">Admin</td>
        <td role="gridcell">
          <span class="sr-only">Status:</span>
          Active
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

## Table Pagination

```html
<div class="space-y-4">
  <!-- Table -->
  <div class="overflow-x-auto">
    <table class="table">
      <!-- Table content -->
    </table>
  </div>
  
  <!-- Pagination -->
  <nav role="navigation" aria-label="Table pagination" class="flex items-center justify-between">
    <div class="text-sm text-muted-foreground">
      Showing 1 to 10 of 97 results
    </div>
    <div class="flex items-center gap-2">
      <button type="button" class="btn-outline" disabled>
        Previous
      </button>
      <button type="button" class="btn">1</button>
      <button type="button" class="btn-outline">2</button>
      <button type="button" class="btn-outline">3</button>
      <span class="px-2">...</span>
      <button type="button" class="btn-outline">10</button>
      <button type="button" class="btn-outline">
        Next
      </button>
    </div>
  </nav>
</div>
```

## JavaScript Functionality

### Sortable Table
```javascript
function makeSortable(table) {
  const headers = table.querySelectorAll('th[data-sortable]');
  
  headers.forEach(header => {
    header.style.cursor = 'pointer';
    header.addEventListener('click', () => {
      const column = header.dataset.sortable;
      const order = header.dataset.order === 'asc' ? 'desc' : 'asc';
      
      // Reset other headers
      headers.forEach(h => h.dataset.order = '');
      header.dataset.order = order;
      
      sortTable(table, column, order);
    });
  });
}

function sortTable(table, column, order) {
  const tbody = table.querySelector('tbody');
  const rows = Array.from(tbody.querySelectorAll('tr'));
  
  rows.sort((a, b) => {
    const aValue = a.querySelector(`[data-column="${column}"]`).textContent;
    const bValue = b.querySelector(`[data-column="${column}"]`).textContent;
    
    if (order === 'asc') {
      return aValue.localeCompare(bValue, undefined, { numeric: true });
    } else {
      return bValue.localeCompare(aValue, undefined, { numeric: true });
    }
  });
  
  rows.forEach(row => tbody.appendChild(row));
}
```

### Row Selection
```javascript
function addRowSelection(table) {
  const selectAll = table.querySelector('thead input[type="checkbox"]');
  const rowCheckboxes = table.querySelectorAll('tbody input[type="checkbox"]');
  
  // Select all functionality
  selectAll.addEventListener('change', () => {
    rowCheckboxes.forEach(checkbox => {
      checkbox.checked = selectAll.checked;
      toggleRowSelection(checkbox.closest('tr'), checkbox.checked);
    });
  });
  
  // Individual row selection
  rowCheckboxes.forEach(checkbox => {
    checkbox.addEventListener('change', () => {
      toggleRowSelection(checkbox.closest('tr'), checkbox.checked);
      updateSelectAllState();
    });
  });
  
  function toggleRowSelection(row, selected) {
    row.classList.toggle('bg-muted/50', selected);
  }
  
  function updateSelectAllState() {
    const checkedCount = [...rowCheckboxes].filter(cb => cb.checked).length;
    selectAll.checked = checkedCount === rowCheckboxes.length;
    selectAll.indeterminate = checkedCount > 0 && checkedCount < rowCheckboxes.length;
  }
}
```

## Best Practices

1. **Clear Headers**: Use descriptive column headers
2. **Consistent Alignment**: Align numerical data right, text left
3. **Responsive Design**: Provide horizontal scroll or alternative layouts
4. **Loading States**: Show skeleton loaders for dynamic data
5. **Empty States**: Provide helpful messages when no data
6. **Row Actions**: Keep action buttons consistent and accessible
7. **Sorting**: Provide clear visual feedback for sortable columns
8. **Pagination**: Break up large datasets appropriately

## Integration Examples

### React Integration
```jsx
import React from 'react';

function Table({ columns, data, onSort, onSelect }) {
  return (
    <div className="overflow-x-auto">
      <table className="table">
        <thead>
          <tr>
            {columns.map((column) => (
              <th key={column.key}>
                {column.sortable ? (
                  <button 
                    onClick={() => onSort(column.key)}
                    className="flex items-center gap-2 font-medium hover:text-foreground"
                  >
                    {column.label}
                    <SortIcon />
                  </button>
                ) : (
                  column.label
                )}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row) => (
            <tr key={row.id}>
              {columns.map((column) => (
                <td key={column.key} className={column.className}>
                  {column.render ? column.render(row[column.key], row) : row[column.key]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
```

### HTMX Integration
```html
<div class="overflow-x-auto">
  <table class="table" 
         hx-get="/api/table-data" 
         hx-trigger="load"
         hx-target="tbody">
    <thead>
      <tr>
        <th>Name</th>
        <th>Email</th>
        <th>Status</th>
      </tr>
    </thead>
    <tbody>
      <tr>
        <td colspan="3" class="text-center py-4">
          <div class="skeleton w-full h-4"></div>
        </td>
      </tr>
    </tbody>
  </table>
</div>
```

### Vue Integration
```vue
<template>
  <div class="overflow-x-auto">
    <table class="table">
      <thead>
        <tr>
          <th v-for="column in columns" :key="column.key">
            <button 
              v-if="column.sortable"
              @click="sort(column.key)"
              class="flex items-center gap-2 font-medium hover:text-foreground"
            >
              {{ column.label }}
              <SortIcon />
            </button>
            <span v-else>{{ column.label }}</span>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="row in data" :key="row.id">
          <td 
            v-for="column in columns" 
            :key="column.key"
            :class="column.className"
          >
            <slot :name="`cell-${column.key}`" :row="row" :value="row[column.key]">
              {{ row[column.key] }}
            </slot>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
```

## Related Components

- [Card](./card.md) - For alternative data display
- [Badge](./badge.md) - For status indicators in tables
- [Button](./button.md) - For table actions
- [Checkbox](./checkbox.md) - For row selection