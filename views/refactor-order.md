# AWO ERP Component Refactoring Priority Order

## Overview
This document outlines the week-by-week refactoring plan to migrate the existing codebase to the atomic design component architecture. The approach follows a bottom-up strategy, starting with atoms and progressing to pages.

## Week 1: Foundation Atoms
**Focus:** Core UI building blocks that all other components depend on

### Priority 1 - Typography & Display
- [ ] Text (heading variants, body text)
- [ ] Label (form labels, required indicators)
- [ ] Icon (icon wrapper with size variants)
- [ ] Divider (horizontal/vertical separators)

### Priority 2 - Form Controls
- [ ] Button (primary, secondary, outline, ghost variants)
- [ ] Input (text, email, password, number types)
- [ ] Textarea (multiline text input)
- [ ] Select (dropdown selection)
- [ ] Checkbox (boolean input)
- [ ] Radio (single selection from group)

### Priority 3 - Feedback Elements
- [ ] Badge (status indicators, counts)
- [ ] Spinner (loading states)
- [ ] Tooltip (contextual help)
- [ ] Alert (inline messages)

## Week 2: Essential Molecules
**Focus:** Common combinations used throughout the application

### Priority 1 - Form Components
- [ ] FormField (label + input + error/help text)
- [ ] FormGroup (grouped form fields)
- [ ] SearchBar (icon + input + clear button)
- [ ] DatePicker (calendar selection)

### Priority 2 - Data Display
- [ ] StatCard (icon + label + value + trend)
- [ ] ListItem (avatar + content + actions)
- [ ] EmptyState (icon + message + action)
- [ ] KeyValue (label-value pairs)

### Priority 3 - Navigation
- [ ] MenuItem (icon + label + indicator)
- [ ] TabItem (label + count + state)
- [ ] Breadcrumb (navigation path)
- [ ] Pagination (page navigation)

## Week 3: Complex Organisms
**Focus:** Self-contained sections with business value

### Priority 1 - Data Management
- [ ] DataTable (sortable, filterable, paginated)
- [ ] FilterPanel (multi-criteria filtering)
- [ ] Card (header + content + actions)
- [ ] Modal (dialog container)

### Priority 2 - Navigation Components
- [ ] Navigation (app header with user menu)
- [ ] Sidebar (collapsible menu structure)
- [ ] TabPanel (tabbed content sections)
- [ ] Dropdown (menu with actions)

### Priority 3 - Forms & Wizards
- [ ] FormSection (grouped form fields)
- [ ] StepIndicator (wizard progress)
- [ ] FileUpload (drag & drop upload)
- [ ] SearchResults (filtered list display)

## Week 4: Layout Templates
**Focus:** Reusable page structures

### Priority 1 - Core Layouts
- [ ] AppLayout (header + sidebar + content)
- [ ] MasterDetailLayout (list + detail view)
- [ ] FormLayout (form-centric pages)
- [ ] DashboardLayout (widget-based layout)

### Priority 2 - Specialized Layouts
- [ ] ReportLayout (print-optimized)
- [ ] WizardLayout (multi-step process)
- [ ] SettingsLayout (tabbed configuration)
- [ ] AuthLayout (login/register pages)

## Week 5: Business Module Pages
**Focus:** Complete page implementations

### Priority 1 - Core Business Pages
- [ ] InvoiceListPage (search, filter, actions)
- [ ] InvoiceDetailPage (view/edit invoice)
- [ ] CustomerListPage (customer management)
- [ ] ProductCatalogPage (inventory browser)

### Priority 2 - Transaction Pages
- [ ] OrderFormPage (create/edit orders)
- [ ] PaymentPage (process payments)
- [ ] QuotationPage (generate quotes)
- [ ] PurchaseOrderPage (supplier orders)

### Priority 3 - Reporting Pages
- [ ] DashboardPage (business metrics)
- [ ] ReportsPage (financial reports)
- [ ] AnalyticsPage (data visualization)
- [ ] AuditLogPage (activity tracking)

## Week 6: Polish & Integration
**Focus:** Refinement and system integration

### Priority 1 - Theme Integration
- [ ] Theme token validation
- [ ] Dark mode support
- [ ] Responsive breakpoints
- [ ] Print styles

### Priority 2 - Interactivity
- [ ] HTMX attribute standardization
- [ ] Alpine.js state management
- [ ] Loading states
- [ ] Error boundaries

### Priority 3 - Accessibility
- [ ] ARIA attributes
- [ ] Keyboard navigation
- [ ] Screen reader support
- [ ] Focus management

## Migration Strategy

### For Each Component:
1. **Analyze** existing implementation
2. **Extract** presentation logic to props
3. **Create** type definitions
4. **Implement** templ component
5. **Test** all variants
6. **Document** usage examples
7. **Replace** old implementations

### Testing Approach:
- Unit tests for helper functions
- Visual regression tests
- Accessibility audits
- Performance benchmarks

### Success Metrics:
- Zero business logic in components
- 100% props-based configuration
- Consistent token usage
- Full type safety
- Improved render performance

## Risk Mitigation

### Potential Challenges:
1. **State Management**: Ensure proper separation of server/client state
2. **Data Flow**: Maintain clean props passing without prop drilling
3. **Performance**: Monitor bundle size and render performance
4. **Backwards Compatibility**: Gradual migration with fallbacks

### Mitigation Strategies:
- Feature flags for gradual rollout
- Parallel implementation (old/new)
- Comprehensive testing suite
- Performance monitoring
- Rollback procedures

## Notes

- Each week builds on previous work
- Components should be immediately usable
- Documentation updated alongside implementation
- Regular team reviews at week boundaries
- Flexibility to adjust based on discoveries