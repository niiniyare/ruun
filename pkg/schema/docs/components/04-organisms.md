# Organisms

**Complex components with business logic that combine multiple molecules and atoms**

## Definition

Organisms are sophisticated UI components that contain business logic, manage complex state, and often coordinate between multiple data sources. They represent complete functional units that users interact with to accomplish specific business tasks.

## Key Characteristics

- **Contains business logic** - handles data processing and business rules
- **Complex state management** - manages application state and side effects
- **Multi-molecule composition** - combines 5+ atoms/molecules
- **Data integration** - fetches, processes, and validates data
- **Workflow orchestration** - coordinates multi-step user interactions
- **Context-aware** - adapts behavior based on business context

## Core Organism Categories

### 1. Data Management Organisms

#### Data Table with Advanced Features

Enterprise-grade data table with sorting, filtering, pagination, and bulk actions.

```go
type DataTableProps struct {
    // Data configuration
    Data        []map[string]interface{} // Table data
    Columns     []TableColumn           // Column definitions
    Schema      *schema.Schema          // Schema definition for validation
    
    // Features
    Sortable    bool
    Filterable  bool
    Searchable  bool
    Selectable  bool   // Row selection
    Paginated   bool
    
    // Pagination config
    PageSize    int
    CurrentPage int
    TotalRows   int
    
    // Actions
    BulkActions []BulkAction
    RowActions  []RowAction
    
    // Business context
    UserPermissions []string
    TenantID        string
    WorkflowState   string
    
    // HTMX integration
    HXGet       string // Data fetching endpoint
    HXTarget    string // Update target
    HXTrigger   string // Refresh trigger
    
    // Alpine.js state
    AlpineData  string // Complex state object
}

type TableColumn struct {
    Key         string
    Label       string
    Type        ColumnType // text, number, date, currency, etc.
    Sortable    bool
    Filterable  bool
    Required    bool
    Width       string
    Align       ColumnAlign
    Format      string     // Formatting rules
    Validation  *schema.FieldValidation
    Permissions []string   // View permissions
}

type BulkAction struct {
    ID          string
    Label       string
    Icon        string
    Variant     ButtonVariant
    Confirmation string    // Confirmation message
    Endpoint    string     // HTMX endpoint
    Permission  string     // Required permission
}
```

**Implementation**:

```go
func dataTableClasses(props DataTableProps) string {
    var classes []string
    
    classes = append(classes,
        "w-full",
        "bg-[var(--card)]",
        "border",
        "border-[var(--border)]",
        "rounded-[var(--radius-lg)]",
        "shadow-[var(--shadow-sm)]",
        "overflow-hidden",
    )
    
    return strings.Join(classes, " ")
}

templ DataTable(props DataTableProps) {
    <div 
        class={ dataTableClasses(props) }
        x-data={ fmt.Sprintf(`{
            data: %s,
            filteredData: %s,
            selectedRows: [],
            currentPage: %d,
            pageSize: %d,
            sortColumn: '',
            sortDirection: 'asc',
            filters: {},
            searchTerm: '',
            
            // Computed properties
            get totalPages() {
                return Math.ceil(this.filteredData.length / this.pageSize);
            },
            
            get paginatedData() {
                const start = (this.currentPage - 1) * this.pageSize;
                return this.filteredData.slice(start, start + this.pageSize);
            },
            
            // Business logic methods
            applyFilters() {
                this.filteredData = this.data.filter(row => {
                    // Apply search filter
                    if (this.searchTerm) {
                        const searchMatch = Object.values(row).some(value => 
                            String(value).toLowerCase().includes(this.searchTerm.toLowerCase())
                        );
                        if (!searchMatch) return false;
                    }
                    
                    // Apply column filters
                    for (const [column, value] of Object.entries(this.filters)) {
                        if (value && row[column] !== value) return false;
                    }
                    
                    return true;
                });
                this.currentPage = 1;
            },
            
            sortData(column) {
                if (this.sortColumn === column) {
                    this.sortDirection = this.sortDirection === 'asc' ? 'desc' : 'asc';
                } else {
                    this.sortColumn = column;
                    this.sortDirection = 'asc';
                }
                
                this.filteredData.sort((a, b) => {
                    const aVal = a[column];
                    const bVal = b[column];
                    const modifier = this.sortDirection === 'asc' ? 1 : -1;
                    
                    if (aVal < bVal) return -1 * modifier;
                    if (aVal > bVal) return 1 * modifier;
                    return 0;
                });
            },
            
            toggleRowSelection(rowId) {
                const index = this.selectedRows.indexOf(rowId);
                if (index > -1) {
                    this.selectedRows.splice(index, 1);
                } else {
                    this.selectedRows.push(rowId);
                }
            },
            
            selectAllRows() {
                if (this.selectedRows.length === this.paginatedData.length) {
                    this.selectedRows = [];
                } else {
                    this.selectedRows = this.paginatedData.map(row => row.id);
                }
            }
        }`, 
        toJSON(props.Data), toJSON(props.Data), props.CurrentPage, props.PageSize) }
        x-init="applyFilters()"
    >
        // Table header with search and bulk actions
        <div class="p-[var(--space-4)] border-b border-[var(--border)]">
            <div class="flex items-center justify-between gap-[var(--space-4)]">
                // Search bar
                if props.Searchable {
                    @SearchBar(SearchBarProps{
                        Name:        "table-search",
                        Placeholder: "Search records...",
                        AlpineModel: "searchTerm",
                        Class:       "max-w-xs",
                    })
                }
                
                // Bulk actions
                if len(props.BulkActions) > 0 {
                    <div class="flex items-center gap-[var(--space-2)]" x-show="selectedRows.length > 0">
                        <span class="text-[var(--muted-foreground)] text-[var(--font-size-sm)]" x-text="`${selectedRows.length} selected`"></span>
                        for _, action := range props.BulkActions {
                            @Button(ButtonProps{
                                Variant:     action.Variant,
                                Size:        ButtonSizeSM,
                                Icon:        action.Icon,
                                HXPost:      action.Endpoint,
                                HXTarget:    props.HXTarget,
                                AlpineClick: if action.Confirmation != "" {
                                    fmt.Sprintf("if (confirm('%s')) { /* submit */ }", action.Confirmation)
                                } else { "" },
                            }) {
                                { action.Label }
                            }
                        }
                    </div>
                }
            </div>
        </div>
        
        // Table content
        <div class="overflow-x-auto">
            <table class="w-full">
                // Table header
                <thead class="bg-[var(--muted)]">
                    <tr>
                        if props.Selectable {
                            <th class="w-[var(--space-12)] p-[var(--space-3)]">
                                @Checkbox(CheckboxProps{
                                    AlpineModel: "selectAllRows()",
                                    Class:       "mx-auto",
                                })
                            </th>
                        }
                        
                        for _, column := range props.Columns {
                            <th class={ "p-[var(--space-3)] text-left font-medium text-[var(--muted-foreground)]" +
                                      if column.Sortable { " cursor-pointer hover:text-[var(--foreground)]" } else { "" } }>
                                <div class="flex items-center gap-[var(--space-2)]">
                                    <span>{ column.Label }</span>
                                    if column.Sortable {
                                        <div 
                                            x-on:click={ fmt.Sprintf("sortData('%s')", column.Key) }
                                            class="flex flex-col"
                                        >
                                            @Icon(IconProps{
                                                Name: "chevron-up",
                                                Size: IconSizeXS,
                                                Class: "transition-colors " + 
                                                      "x-bind:class=\"sortColumn === '" + column.Key + "' && sortDirection === 'asc' ? 'text-[var(--primary)]' : 'text-[var(--muted-foreground)]'\"",
                                            })
                                            @Icon(IconProps{
                                                Name: "chevron-down", 
                                                Size: IconSizeXS,
                                                Class: "transition-colors " +
                                                      "x-bind:class=\"sortColumn === '" + column.Key + "' && sortDirection === 'desc' ? 'text-[var(--primary)]' : 'text-[var(--muted-foreground)]'\"",
                                            })
                                        </div>
                                    }
                                </div>
                            </th>
                        }
                        
                        if len(props.RowActions) > 0 {
                            <th class="w-[var(--space-20)] p-[var(--space-3)]">Actions</th>
                        }
                    </tr>
                </thead>
                
                // Table body
                <tbody class="divide-y divide-[var(--border)]">
                    <template x-for="(row, index) in paginatedData" :key="row.id">
                        <tr class="hover:bg-[var(--muted)]/30 transition-colors">
                            if props.Selectable {
                                <td class="p-[var(--space-3)] text-center">
                                    @Checkbox(CheckboxProps{
                                        AlpineModel: "selectedRows.includes(row.id)",
                                        AlpineClick: "toggleRowSelection(row.id)",
                                    })
                                </td>
                            }
                            
                            for _, column := range props.Columns {
                                <td class="p-[var(--space-3)]">
                                    @tableCellContent(column)
                                </td>
                            }
                            
                            if len(props.RowActions) > 0 {
                                <td class="p-[var(--space-3)]">
                                    @tableRowActions(props.RowActions)
                                </td>
                            }
                        </tr>
                    </template>
                </tbody>
            </table>
        </div>
        
        // Pagination
        if props.Paginated {
            <div class="p-[var(--space-4)] border-t border-[var(--border)] bg-[var(--muted)]/20">
                @Pagination(PaginationProps{
                    CurrentPage: props.CurrentPage,
                    TotalPages:  "totalPages",
                    PageSize:    props.PageSize,
                    TotalItems:  "filteredData.length",
                    AlpineModel: "currentPage",
                    HXGet:       props.HXGet,
                    HXTarget:    props.HXTarget,
                })
            </div>
        }
    </div>
}
```

### 2. Form Organisms

#### Multi-Step Wizard Form

Complex form with validation, conditional logic, and workflow management.

```go
type WizardFormProps struct {
    // Wizard configuration
    Steps       []WizardStep
    CurrentStep int
    Schema      *schema.Schema
    
    // Data management
    FormData    map[string]interface{}
    Validation  map[string][]string // Validation errors
    
    // Business logic
    BusinessRules map[string]*schema.BusinessRule
    Workflow      *schema.Workflow
    
    // Behavior
    AllowSkip     bool
    SaveProgress  bool
    ShowProgress  bool
    
    // Integration
    HXPost     string
    HXTarget   string
    OnComplete string // JavaScript callback
}

type WizardStep struct {
    ID          string
    Title       string
    Description string
    Icon        string
    Fields      []schema.Field
    Validation  *schema.StepValidation
    Conditional *schema.ConditionalLogic
    Required    bool
}
```

**Implementation**:

```go
templ WizardForm(props WizardFormProps) {
    <div 
        class="max-w-4xl mx-auto bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] shadow-[var(--shadow-lg)]"
        x-data={ fmt.Sprintf(`{
            currentStep: %d,
            totalSteps: %d,
            formData: %s,
            errors: %s,
            isSubmitting: false,
            
            get currentStepData() {
                return this.steps[this.currentStep];
            },
            
            get canProceed() {
                return this.validateCurrentStep();
            },
            
            get canSkip() {
                return %t && !this.currentStepData.required;
            },
            
            validateCurrentStep() {
                const step = this.steps[this.currentStep];
                // Validation logic based on schema
                return true; // Simplified
            },
            
            nextStep() {
                if (this.canProceed && this.currentStep < this.totalSteps - 1) {
                    this.currentStep++;
                    if (%t) this.saveProgress();
                }
            },
            
            prevStep() {
                if (this.currentStep > 0) {
                    this.currentStep--;
                }
            },
            
            skipStep() {
                if (this.canSkip) {
                    this.nextStep();
                }
            },
            
            goToStep(stepIndex) {
                if (stepIndex >= 0 && stepIndex < this.totalSteps) {
                    this.currentStep = stepIndex;
                }
            },
            
            submitForm() {
                this.isSubmitting = true;
                // Form submission logic
            },
            
            saveProgress() {
                // Auto-save logic for draft state
            }
        }`, props.CurrentStep, len(props.Steps), toJSON(props.FormData), 
            toJSON(props.Validation), props.AllowSkip, props.SaveProgress) }
    >
        // Progress indicator
        if props.ShowProgress {
            <div class="p-[var(--space-6)] border-b border-[var(--border)]">
                @WizardProgressBar(WizardProgressBarProps{
                    Steps:       props.Steps,
                    CurrentStep: props.CurrentStep,
                    AlpineModel: "currentStep",
                })
            </div>
        }
        
        // Step content
        <div class="p-[var(--space-6)]">
            // Step header
            <div class="mb-[var(--space-8)]">
                <template x-for="(step, index) in steps" :key="step.id">
                    <div x-show="currentStep === index" class="transition-all duration-200">
                        <div class="flex items-center gap-[var(--space-4)] mb-[var(--space-4)]">
                            <div class="w-[var(--space-12)] h-[var(--space-12)] bg-[var(--primary)] rounded-full flex items-center justify-center">
                                @Icon(IconProps{
                                    Name:  "x-bind:step.icon",
                                    Size:  IconSizeLG,
                                    Class: "text-[var(--primary-foreground)]",
                                })
                            </div>
                            <div>
                                <h2 class="text-xl font-semibold text-[var(--foreground)]" x-text="step.title"></h2>
                                <p class="text-[var(--muted-foreground)]" x-text="step.description"></p>
                            </div>
                        </div>
                    </div>
                </template>
            </div>
            
            // Dynamic form fields based on current step
            <form 
                hx-post={ props.HXPost }
                hx-target={ props.HXTarget }
                hx-trigger="submit"
                x-on:submit.prevent="submitForm"
            >
                for i, step := range props.Steps {
                    <div x-show={ fmt.Sprintf("currentStep === %d", i) } class="space-y-[var(--space-6)]">
                        for _, field := range step.Fields {
                            @schemaFieldRenderer(field, props.Schema, props.FormData)
                        }
                    </div>
                }
                
                // Form navigation
                <div class="flex items-center justify-between pt-[var(--space-8)] border-t border-[var(--border)] mt-[var(--space-8)]">
                    // Previous button
                    <div>
                        @Button(ButtonProps{
                            Variant:     ButtonOutline,
                            Size:        ButtonSizeLG,
                            IconLeft:    "chevron-left",
                            Type:        "button",
                            AlpineClick: "prevStep()",
                            Class:       "x-bind:disabled=\"currentStep === 0\"",
                        }) {
                            Previous
                        }
                    </div>
                    
                    // Action buttons
                    <div class="flex items-center gap-[var(--space-4)]">
                        // Skip button
                        if props.AllowSkip {
                            @Button(ButtonProps{
                                Variant:     ButtonGhost,
                                Size:        ButtonSizeLG,
                                Type:        "button",
                                AlpineClick: "skipStep()",
                                Class:       "x-show=\"canSkip\"",
                            }) {
                                Skip
                            }
                        }
                        
                        // Next/Submit button
                        @Button(ButtonProps{
                            Variant:     ButtonPrimary,
                            Size:        ButtonSizeLG,
                            IconRight:   "chevron-right",
                            Type:        "button",
                            Loading:     true,
                            AlpineClick: "currentStep === totalSteps - 1 ? submitForm() : nextStep()",
                            Class:       "x-bind:disabled=\"!canProceed\" x-bind:class=\"isSubmitting ? 'pointer-events-none opacity-50' : ''\"",
                        }) {
                            <span x-text="currentStep === totalSteps - 1 ? 'Complete' : 'Next'"></span>
                        }
                    </div>
                </div>
            </form>
        </div>
    </div>
}
```

### 3. Navigation Organisms

#### Sidebar Navigation with Permissions

Dynamic navigation with role-based access control and multi-level menus.

```go
type SidebarNavigationProps struct {
    // Navigation structure
    MenuItems   []MenuItem
    CurrentPath string
    
    // Business context
    UserRole        string
    UserPermissions []string
    TenantID        string
    
    // Configuration
    Collapsible bool
    DefaultOpen bool
    ShowIcons   bool
    ShowBadges  bool
    
    // State management
    AlpineModel string // For collapse state
}

type MenuItem struct {
    ID          string
    Label       string
    Icon        string
    Path        string
    Children    []MenuItem
    Badge       *MenuBadge
    Permission  string        // Required permission
    Roles       []string      // Allowed roles
    Conditional *schema.ConditionalLogic
    HXGet       string        // HTMX navigation
    External    bool          // External link
}

type MenuBadge struct {
    Text    string
    Variant BadgeVariant
    Count   int
    Source  string // Data source for dynamic count
}
```

**Implementation**:

```go
templ SidebarNavigation(props SidebarNavigationProps) {
    <nav 
        class="flex flex-col bg-[var(--card)] border-r border-[var(--border)] h-full"
        x-data={ fmt.Sprintf(`{
            isOpen: %t,
            currentPath: '%s',
            userPermissions: %s,
            
            hasPermission(permission) {
                return !permission || this.userPermissions.includes(permission);
            },
            
            isActive(path) {
                return this.currentPath === path || this.currentPath.startsWith(path + '/');
            },
            
            toggle() {
                this.isOpen = !this.isOpen;
            }
        }`, props.DefaultOpen, props.CurrentPath, toJSON(props.UserPermissions)) }
        x-bind:class="isOpen ? 'w-64' : 'w-16'"
        class="transition-all duration-200"
    >
        // Navigation header
        <div class="p-[var(--space-4)] border-b border-[var(--border)]">
            <div class="flex items-center justify-between">
                <div class="flex items-center gap-[var(--space-3)]" x-show="isOpen" x-transition>
                    @Logo(LogoProps{Size: "md"})
                    <span class="font-semibold text-[var(--foreground)]">Dashboard</span>
                </div>
                
                if props.Collapsible {
                    @Button(ButtonProps{
                        Variant:     ButtonGhost,
                        Size:        ButtonSizeSM,
                        Icon:        "menu",
                        Type:        "button",
                        AlpineClick: "toggle()",
                        Class:       "ml-auto",
                    })
                }
            </div>
        </div>
        
        // Navigation menu
        <div class="flex-1 p-[var(--space-2)] space-y-[var(--space-1)]">
            for _, item := range props.MenuItems {
                @navigationItem(item, 0, props)
            }
        </div>
        
        // User section
        <div class="p-[var(--space-4)] border-t border-[var(--border)]">
            @UserMenu(UserMenuProps{
                Collapsed: "!isOpen",
                ShowAvatar: true,
                ShowName: true,
            })
        </div>
    </nav>
}

templ navigationItem(item MenuItem, depth int, props SidebarNavigationProps) {
    // Check permissions
    if item.Permission != "" {
        <div x-show={ fmt.Sprintf("hasPermission('%s')", item.Permission) }>
            @renderNavigationItem(item, depth, props)
        </div>
    } else {
        @renderNavigationItem(item, depth, props)
    }
}

templ renderNavigationItem(item MenuItem, depth int, props SidebarNavigationProps) {
    if len(item.Children) > 0 {
        // Parent menu item with children
        <div x-data="{ open: false }" class="space-y-[var(--space-1)]">
            <button
                type="button"
                x-on:click="open = !open"
                class={ "w-full flex items-center gap-[var(--space-3)] px-[var(--space-3)] py-[var(--space-2)] rounded-[var(--radius-md)] text-[var(--muted-foreground)] hover:bg-[var(--accent)] hover:text-[var(--accent-foreground)] transition-colors" +
                       fmt.Sprintf(" pl-[var(--space-%d)]", 3 + (depth * 4)) }
            >
                if props.ShowIcons && item.Icon != "" {
                    @Icon(IconProps{
                        Name: item.Icon,
                        Size: IconSizeMD,
                        Class: "flex-shrink-0",
                    })
                }
                
                <span class="flex-1 text-left" x-show="$parent.isOpen" x-transition>{ item.Label }</span>
                
                <div x-show="$parent.isOpen" x-transition>
                    @Icon(IconProps{
                        Name: "chevron-right",
                        Size: IconSizeSM,
                        Class: "transform transition-transform x-bind:class=\"open ? 'rotate-90' : ''\"",
                    })
                </div>
            </button>
            
            // Submenu
            <div x-show="open && $parent.isOpen" x-transition class="space-y-[var(--space-1)]">
                for _, child := range item.Children {
                    @navigationItem(child, depth+1, props)
                }
            </div>
        </div>
    } else {
        // Leaf menu item
        if item.External {
            <a
                href={ item.Path }
                target="_blank"
                rel="noopener noreferrer"
                class={ "flex items-center gap-[var(--space-3)] px-[var(--space-3)] py-[var(--space-2)] rounded-[var(--radius-md)] transition-colors" +
                       " x-bind:class=\"isActive('" + item.Path + "') ? 'bg-[var(--primary)] text-[var(--primary-foreground)]' : 'text-[var(--muted-foreground)] hover:bg-[var(--accent)] hover:text-[var(--accent-foreground)]'\"" +
                       fmt.Sprintf(" pl-[var(--space-%d)]", 3 + (depth * 4)) }
            >
                @menuItemContent(item, props)
            </a>
        } else if item.HXGet != "" {
            <button
                type="button"
                hx-get={ item.HXGet }
                hx-target="#main-content"
                hx-push-url="true"
                class={ "w-full flex items-center gap-[var(--space-3)] px-[var(--space-3)] py-[var(--space-2)] rounded-[var(--radius-md)] transition-colors" +
                       " x-bind:class=\"isActive('" + item.Path + "') ? 'bg-[var(--primary)] text-[var(--primary-foreground)]' : 'text-[var(--muted-foreground)] hover:bg-[var(--accent)] hover:text-[var(--accent-foreground)]'\"" +
                       fmt.Sprintf(" pl-[var(--space-%d)]", 3 + (depth * 4)) }
            >
                @menuItemContent(item, props)
            </button>
        } else {
            <a
                href={ item.Path }
                class={ "flex items-center gap-[var(--space-3)] px-[var(--space-3)] py-[var(--space-2)] rounded-[var(--radius-md)] transition-colors" +
                       " x-bind:class=\"isActive('" + item.Path + "') ? 'bg-[var(--primary)] text-[var(--primary-foreground)]' : 'text-[var(--muted-foreground)] hover:bg-[var(--accent)] hover:text-[var(--accent-foreground)]'\"" +
                       fmt.Sprintf(" pl-[var(--space-%d)]", 3 + (depth * 4)) }
            >
                @menuItemContent(item, props)
            </a>
        }
    }
}

templ menuItemContent(item MenuItem, props SidebarNavigationProps) {
    if props.ShowIcons && item.Icon != "" {
        @Icon(IconProps{
            Name: item.Icon,
            Size: IconSizeMD,
            Class: "flex-shrink-0",
        })
    }
    
    <span class="flex-1" x-show="$parent.isOpen" x-transition>{ item.Label }</span>
    
    if props.ShowBadges && item.Badge != nil {
        <div x-show="$parent.isOpen" x-transition>
            @Badge(BadgeProps{
                Variant: item.Badge.Variant,
                Size:    BadgeSizeSM,
            }) {
                if item.Badge.Count > 0 {
                    { strconv.Itoa(item.Badge.Count) }
                } else {
                    { item.Badge.Text }
                }
            }
        </div>
    }
}
```

## Organism Design Principles

### 1. Business Logic Integration

Organisms should contain and coordinate business logic:

**✅ Good Example**:
```go
// Order management organism with business rules
templ OrderManagement(props OrderManagementProps) {
    <div x-data="{
        orders: [],
        filters: {},
        
        // Business logic methods
        async updateOrderStatus(orderId, status) {
            const result = await this.validateStatusChange(orderId, status);
            if (result.valid) {
                // Apply business rules and update
                await this.applyOrderWorkflow(orderId, status);
            }
        },
        
        validateStatusChange(orderId, newStatus) {
            // Business validation logic
            const order = this.orders.find(o => o.id === orderId);
            return this.orderWorkflow.canTransition(order.status, newStatus);
        }
    }">
        // Complex order management interface
    </div>
}
```

### 2. Data Orchestration

Organisms coordinate between multiple data sources:

**✅ Good Pattern**:
```go
// Dashboard organism that orchestrates multiple data sources
templ ExecutiveDashboard(props DashboardProps) {
    <div x-data="{
        // Multiple data streams
        metrics: {},
        alerts: [],
        activities: [],
        reports: [],
        
        // Data loading coordination
        async loadDashboardData() {
            const [metrics, alerts, activities, reports] = await Promise.all([
                this.fetchMetrics(),
                this.fetchAlerts(),
                this.fetchRecentActivities(),
                this.fetchReports()
            ]);
            
            // Coordinate data updates
            this.updateDashboard({metrics, alerts, activities, reports});
        }
    }">
        // Complex dashboard layout
    </div>
}
```

### 3. State Management

Organisms manage complex application state:

```go
// Shopping cart organism with complex state
templ ShoppingCart(props CartProps) {
    <div x-data="{
        items: [],
        totals: { subtotal: 0, tax: 0, shipping: 0, total: 0 },
        discounts: [],
        
        // Complex state calculations
        recalculateTotals() {
            this.totals.subtotal = this.items.reduce((sum, item) => 
                sum + (item.price * item.quantity), 0);
            
            this.totals.tax = this.calculateTax();
            this.totals.shipping = this.calculateShipping();
            
            // Apply discounts
            const discountAmount = this.calculateDiscounts();
            this.totals.total = this.totals.subtotal + this.totals.tax + 
                               this.totals.shipping - discountAmount;
        },
        
        addItem(product, quantity = 1) {
            // Business logic for adding items
            const existingItem = this.items.find(item => item.id === product.id);
            if (existingItem) {
                existingItem.quantity += quantity;
            } else {
                this.items.push({...product, quantity});
            }
            this.recalculateTotals();
        }
    }">
        // Shopping cart interface
    </div>
}
```

## Testing Organisms

### Integration Testing

```go
func TestDataTable_SortingAndFiltering(t *testing.T) {
    props := DataTableProps{
        Data: []map[string]interface{}{
            {"id": 1, "name": "John", "role": "Admin"},
            {"id": 2, "name": "Jane", "role": "User"},
            {"id": 3, "name": "Bob", "role": "Admin"},
        },
        Columns: []TableColumn{
            {Key: "name", Label: "Name", Sortable: true},
            {Key: "role", Label: "Role", Filterable: true},
        },
        Sortable:   true,
        Filterable: true,
    }
    
    // Test sorting functionality
    html := renderToString(DataTable(props))
    
    // Verify sorting controls are present
    assert.Contains(t, html, "x-on:click=\"sortData('name')\"")
    assert.Contains(t, html, "chevron-up")
    
    // Test Alpine.js state management
    assert.Contains(t, html, "sortColumn:")
    assert.Contains(t, html, "sortDirection:")
}

func TestWizardForm_StepValidation(t *testing.T) {
    props := WizardFormProps{
        Steps: []WizardStep{
            {
                ID:       "personal",
                Title:    "Personal Information",
                Required: true,
                Fields: []schema.Field{
                    {Name: "firstName", Required: true},
                    {Name: "lastName", Required: true},
                },
            },
            {
                ID:    "preferences",
                Title: "Preferences",
            },
        },
        CurrentStep: 0,
    }
    
    html := renderToString(WizardForm(props))
    
    // Verify step validation logic
    assert.Contains(t, html, "validateCurrentStep()")
    assert.Contains(t, html, "canProceed")
    
    // Test required step handling
    assert.Contains(t, html, "currentStepData.required")
}
```

### Business Logic Testing

```go
func TestSidebarNavigation_PermissionFiltering(t *testing.T) {
    props := SidebarNavigationProps{
        MenuItems: []MenuItem{
            {ID: "dashboard", Label: "Dashboard", Path: "/dashboard"},
            {ID: "admin", Label: "Admin", Path: "/admin", Permission: "admin.access"},
            {ID: "reports", Label: "Reports", Path: "/reports", Permission: "reports.view"},
        },
        UserPermissions: []string{"reports.view"},
    }
    
    html := renderToString(SidebarNavigation(props))
    
    // Verify permission-based visibility
    assert.Contains(t, html, "hasPermission('admin.access')")
    assert.Contains(t, html, "hasPermission('reports.view')")
    
    // Test permission checking logic
    assert.Contains(t, html, "userPermissions.includes(permission)")
}
```

## Best Practices

### ✅ DO

- Implement complex business logic in organisms
- Coordinate multiple data sources efficiently
- Manage application state appropriately
- Use proper error handling and loading states
- Integrate with backend workflows and permissions
- Provide comprehensive user feedback
- Test business logic thoroughly

### ❌ DON'T

- Create organisms without clear business purpose
- Duplicate business logic across multiple organisms
- Ignore performance implications of complex state
- Skip error handling and edge cases
- Create overly coupled organisms
- Ignore accessibility in complex interfaces

## Related Documentation

- **[Molecules](./03-molecules.md)** - Building blocks for organisms
- **[Templates](./05-templates.md)** - Page layout structures  
- **[Schema Integration](./07-integration.md)** - Schema-component patterns
- **[Business Rules](../schema/business_rules.md)** - Backend business logic system

---

Organisms represent the most sophisticated components in our atomic design system, handling complex business logic and coordinating multiple elements to create complete functional units that users interact with to accomplish business tasks.