# AWO ERP Component Mapping: Basecoat vs Custom

This document maps the component needs for AWO ERP against what Basecoat provides, identifying which components can use Basecoat classes directly versus requiring custom implementation.

## ✅ BASECOAT PROVIDED COMPONENTS

These components are **fully available** in Basecoat with complete styling and behavior:

### Form & Input Components
- **`btn`** (+ variants: primary, secondary, outline, ghost, link, destructive, all sizes)
- **`field`** - Form field wrapper
- **`input`** - All input types (text, email, password, number, file, tel, url, search, date/time)
- **`textarea`** - Multiline text input
- **`select`** - Native and custom select with popover
- **`checkbox`** - Standard checkbox input
- **`radio`** - Radio button input
- **`switch`** - Toggle switch (checkbox with role="switch")
- **`label`** - Form labels

### Layout & Container Components  
- **`card`** - Card container with header/section/footer structure
- **`sidebar`** - Collapsible sidebar navigation
- **`table`** - Complete table styling (thead, tbody, tr, th, td)

### Interactive Components
- **`dialog`** - Modal dialogs with backdrop
- **`dropdown-menu`** - Dropdown menus with popover positioning
- **`popover`** - Positioned popover containers
- **`tabs`** - Tab navigation and panels
- **`command`** - Command palette/menu

### Feedback Components
- **`alert`** (+ destructive variant) - Status alerts
- **`badge`** (+ variants: primary, secondary, destructive, outline) - Status badges
- **`toast`** - Toast notifications with container
- **`tooltip`** - Hover tooltips

### Utility Components
- **`accordion`** (via `details` element) - Collapsible content
- **`separator`** - Visual dividers
- **`skeleton`** - Loading placeholders
- **`progress`** (via `input[type="range"]`) - Progress bars
- **`kbd`** - Keyboard key styling
- **`scroll-area`** - Scrollable content areas
- **`collapsible`** - Expandable content
- **`navigation-menu`** - Navigation menus
- **`spinner`** - Loading indicators
- **`sheet`** - Side panels/sheets

---

## ❌ CUSTOM CSS REQUIRED COMPONENTS

These components need **custom implementation** as they're not provided by Basecoat:

### Typography & Content
- **`icon`** - Icon display and sizing
- **`link`** - Link styling beyond button variants
- **`divider`** - Content dividers (beyond separator)
- **`heading`** - Typography hierarchy (h1-h6)
- **`text`** - Text content styling (body, caption, etc.)

### Specialized Components
- **`status-badge`** - Status-specific badges beyond basic variants
- **`pagination`** - Page navigation controls
- **`empty-state`** - Empty state illustrations and messaging
- **`filter-chip`** - Removable filter tags
- **`input-group`** - Input with attached buttons/icons
- **`data-table`** - Enhanced tables with sorting, filtering, actions
- **`stats-card`** - Metric cards with trends and comparisons
- **`wizard`** - Multi-step form navigation
- **`breadcrumbs`** - Hierarchical navigation

### Layout Templates
- **`layouts`** - Page layout templates (master-detail, form, dashboard)

---

## 🔄 IMPLEMENTATION STRATEGY

### Phase 1: Leverage Basecoat (85% Coverage)
Use Basecoat classes directly for:
```templ
// Forms
@Button(ButtonProps{Variant: "primary"})      // Uses .btn-primary
@Input(InputProps{Type: "text"})              // Uses .input
@Select(SelectProps{Options: opts})           // Uses .select

// Layout  
@Card(CardProps{Title: "Invoice"})            // Uses .card
@Dialog(DialogProps{Open: true})              // Uses .dialog
@Tabs(TabsProps{Items: tabs})                 // Uses .tabs

// Feedback
@Alert(AlertProps{Variant: "error"})          // Uses .alert-destructive
@Badge(BadgeProps{Variant: "success"})        // Uses .badge-primary
@Toast(ToastProps{Message: "Saved"})          // Uses .toast
```

### Phase 2: Custom Component Extensions
Build custom components that compose Basecoat foundations:

```templ
// Custom DataTable using Basecoat table + custom controls
templ DataTable(props DataTableProps) {
    <div class="space-y-4">
        <!-- Custom filter chips -->
        @FilterChips(props.Filters)
        
        <!-- Basecoat table -->
        <table class="table">
            @DataTableHeader(props.Columns)
            @DataTableBody(props.Rows)
        </table>
        
        <!-- Custom pagination -->
        @Pagination(props.PaginationConfig)
    </div>
}

// Custom StatsCard using Basecoat card + custom layout
templ StatsCard(props StatsCardProps) {
    <div class="card">
        <div class="flex items-center justify-between">
            @Icon(IconProps{Name: props.Icon})
            @Badge(BadgeProps{
                Label:   props.Trend,
                Variant: props.IsPositive ? "primary" : "destructive",
            })
        </div>
        @Heading(HeadingProps{Level: 3, Text: props.Value})
        @Text(TextProps{Size: "sm", Text: props.Label})
    </div>
}
```

### Phase 3: Theme Integration
Basecoat's CSS custom properties integrate seamlessly with our token system:

```css
/* Our theme tokens map to Basecoat variables */
:root {
  --primary: oklch(0.205 0 0);           /* Maps to Basecoat --primary */
  --primary-foreground: oklch(0.985 0 0); /* Maps to --primary-foreground */
  --border: oklch(0.922 0 0);            /* Maps to Basecoat --border */
  /* ... */
}
```

---

## 📊 COVERAGE ANALYSIS

| Component Category | Basecoat Coverage | Custom Required |
|---|---|---|
| **Form Controls** | 100% | 0% |
| **Basic Layout** | 90% | 10% |
| **Interactive** | 95% | 5% |
| **Feedback** | 85% | 15% |
| **Typography** | 20% | 80% |
| **Data Display** | 70% | 30% |
| **Navigation** | 80% | 20% |
| **Specialized** | 0% | 100% |

**Overall: ~75% Basecoat coverage, 25% custom implementation needed**

---

## 🎯 BENEFITS OF THIS APPROACH

### ✅ Advantages
1. **Rapid Development** - 75% of components work out-of-the-box
2. **Consistent Design** - Basecoat provides cohesive design system
3. **Accessibility** - Basecoat includes ARIA attributes and keyboard navigation
4. **Theme Support** - Dark/light modes built-in
5. **Responsive Design** - Mobile-first responsive utilities
6. **Maintenance** - Framework handles browser compatibility and updates

### ⚠️ Considerations
1. **Bundle Size** - Including full Basecoat (~50KB compressed)
2. **Customization Limits** - Some design constraints from framework
3. **Learning Curve** - Team needs to understand Basecoat patterns
4. **Custom Components** - Still need 25% custom implementation

---

## 🚀 RECOMMENDED NEXT STEPS

1. **Audit Current Components** - Review existing templ components against this mapping
2. **Basecoat Integration** - Add Basecoat CSS to build pipeline  
3. **Component Refactoring** - Update existing atoms/molecules to use Basecoat classes
4. **Custom Component Development** - Build the 25% that requires custom CSS
5. **Theme Synchronization** - Align our token system with Basecoat variables
6. **Documentation** - Update component docs with Basecoat usage patterns

This mapping provides a clear path to leverage Basecoat's robust foundation while maintaining flexibility for AWO ERP's specific requirements.