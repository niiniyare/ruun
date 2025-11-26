# Pages

**Complete page implementations that combine templates with real data and business context**

## Definition

Pages are the highest level in atomic design - complete instances of templates filled with real content, data, and business logic. They represent actual user-facing screens with specific purposes and functionality within the application.

## Key Characteristics

- **Complete implementations** - fully functional pages with real data
- **Business context** - tied to specific user workflows and tasks  
- **Data integration** - connected to backend data sources and APIs
- **State management** - handles full application state for the page
- **User interactions** - complete user experience with all interactions
- **Performance optimized** - implements loading, caching, and optimization strategies

## Core Page Categories

### 1. Dashboard Pages

#### Executive Dashboard Page

Complete executive dashboard with real-time data, metrics, and business insights.

```go
type ExecutiveDashboardPageProps struct {
    // Data context
    UserID      string
    TenantID    string
    DateRange   DateRange
    Permissions []string
    
    // Real data
    Metrics     []DashboardMetric
    Charts      []ChartConfig
    Activities  []ActivityItem
    Alerts      []AlertItem
    
    // Business context
    Department  string
    Role        string
    Region      string
    
    // Schema context
    Schema      *schema.Schema
    Workflow    *schema.Workflow
    
    // Page state
    RefreshInterval time.Duration
    LastUpdate     time.Time
    LoadingState   LoadingState
    ErrorState     *ErrorState
}

type DashboardMetric struct {
    ID          string
    Title       string
    Value       any
    PrevValue   any
    Trend       TrendDirection
    TrendValue  string
    Format      MetricFormat
    Unit        string
    Target      any
    Permission  string
    DrillDown   *DrillDownConfig
}

type ChartConfig struct {
    ID       string
    Type     ChartType
    Title    string
    Data     any
    Options  map[string]any
    Height   int
    Width    int
    Position GridPosition
}

type ActivityItem struct {
    ID        string
    Type      ActivityType
    Title     string
    Message   string
    UserName  string
    UserAvatar string
    Timestamp time.Time
    Icon      string
    Color     string
    Link      string
}
```

**Implementation**:

```go
templ ExecutiveDashboardPage(props ExecutiveDashboardPageProps) {
    @DashboardTemplate(DashboardTemplateProps{
        Layout:      DashboardLayoutSidebar,
        HeaderArea:  true,
        SidebarArea: true,
        MetricsArea: true,
        ChartsArea:  true,
        ActivityArea: true,
        AlertsArea:  true,
        Responsive:  true,
    }) {
        // Inject real content into template slots
        
        // Header slot
        @headerContent(props)
        
        // Sidebar slot  
        @sidebarNavigation(props)
        
        // Metrics slot
        @metricsSection(props.Metrics, props.Permissions)
        
        // Charts slot
        @chartsSection(props.Charts, props.DateRange)
        
        // Activity slot
        @activityFeed(props.Activities)
        
        // Alerts slot
        @alertsSection(props.Alerts)
    }
}

templ headerContent(props ExecutiveDashboardPageProps) {
    <div slot="breadcrumbs-slot">
        @Breadcrumb(BreadcrumbProps{
            Items: []BreadcrumbItem{
                {Label: "Home", Href: "/"},
                {Label: "Dashboard", Href: "/dashboard"},
                {Label: "Executive Overview"},
            },
        })
    </div>
    
    <div slot="header-actions-slot">
        @DateRangePicker(DateRangePickerProps{
            Value:     props.DateRange,
            HXPost:    "/dashboard/date-range",
            HXTarget:  "#dashboard-content",
        })
        
        @Button(ButtonProps{
            Variant:  ButtonOutline,
            Size:     ButtonSizeMD,
            Icon:     "download",
            HXPost:   "/dashboard/export",
            HXTarget: "#export-modal",
        }) {
            Export
        }
    </div>
    
    <div slot="notifications-slot">
        @NotificationBell(NotificationBellProps{
            Count:   len(props.Alerts),
            HXGet:   "/notifications",
            HXTarget: "#notifications-modal",
        })
    </div>
    
    <div slot="user-menu-slot">
        @UserMenu(UserMenuProps{
            UserID:      props.UserID,
            Role:        props.Role,
            Department:  props.Department,
            ShowAvatar:  true,
        })
    </div>
}

templ metricsSection(metrics []DashboardMetric, permissions []string) {
    <div slot="metrics-slot">
        for _, metric := range metrics {
            if hasPermission(permissions, metric.Permission) {
                @MetricCard(MetricCardProps{
                    Title:      metric.Title,
                    Value:      formatValue(metric.Value, metric.Format),
                    PrevValue:  formatValue(metric.PrevValue, metric.Format),
                    Trend:      metric.Trend,
                    TrendValue: metric.TrendValue,
                    Unit:       metric.Unit,
                    Target:     formatValue(metric.Target, metric.Format),
                    DrillDown:  metric.DrillDown,
                    HXGet:      fmt.Sprintf("/metrics/%s/details", metric.ID),
                    HXTarget:   "#metric-detail-modal",
                })
            }
        }
    </div>
}

templ chartsSection(charts []ChartConfig, dateRange DateRange) {
    <div slot="charts-slot">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-[var(--space-6)]">
            for _, chart := range charts {
                @ChartCard(ChartCardProps{
                    ID:        chart.ID,
                    Type:      chart.Type,
                    Title:     chart.Title,
                    Data:      chart.Data,
                    Options:   chart.Options,
                    Height:    chart.Height,
                    DateRange: dateRange,
                    HXGet:     fmt.Sprintf("/charts/%s/data", chart.ID),
                    HXTrigger: "load, every 30s",
                })
            }
        </div>
    </div>
}

templ activityFeed(activities []ActivityItem) {
    <div slot="activity-slot">
        <div class="space-y-[var(--space-4)]">
            <h3 class="text-lg font-semibold text-[var(--foreground)]">Recent Activity</h3>
            
            <div class="space-y-[var(--space-3)]">
                for _, activity := range activities {
                    @ActivityItem(ActivityItemProps{
                        Type:      activity.Type,
                        Title:     activity.Title,
                        Message:   activity.Message,
                        UserName:  activity.UserName,
                        UserAvatar: activity.UserAvatar,
                        Timestamp: activity.Timestamp,
                        Icon:      activity.Icon,
                        Color:     activity.Color,
                        Link:      activity.Link,
                    })
                }
            </div>
            
            @Button(ButtonProps{
                Variant:  ButtonLink,
                Size:     ButtonSizeSM,
                HXGet:    "/activity/all",
                HXTarget: "#activity-modal",
            }) {
                View all activity
            }
        </div>
    </div>
}
```

### 2. Form Pages

#### Invoice Creation Page

Complete invoice creation form with complex business logic and validation.

```go
type InvoiceCreationPageProps struct {
    // Business context
    CompanyID    string
    UserID       string
    CustomerID   string // Optional for edit
    InvoiceID    string // For editing existing
    
    // Form data
    Invoice      *Invoice
    Customer     *Customer
    LineItems    []LineItem
    TaxSettings  TaxSettings
    
    // Supporting data
    Customers    []CustomerOption
    Products     []ProductOption
    TaxRates     []TaxRate
    
    // Schema context
    Schema       *schema.Schema
    Workflow     *schema.Workflow
    Validation   *schema.Validation
    
    // Page state
    Mode         FormMode // create, edit, view
    SaveState    SaveState
    Errors       map[string][]string
    Warnings     []string
}

type Invoice struct {
    ID           string
    Number       string
    Date         time.Time
    DueDate      time.Time
    CustomerID   string
    Status       InvoiceStatus
    Subtotal     decimal.Decimal
    TaxAmount    decimal.Decimal
    Total        decimal.Decimal
    Notes        string
    Terms        string
    Currency     string
}

type LineItem struct {
    ID           string
    ProductID    string
    Description  string
    Quantity     decimal.Decimal
    UnitPrice    decimal.Decimal
    TaxRate      decimal.Decimal
    Total        decimal.Decimal
}
```

**Implementation**:

```go
templ InvoiceCreationPage(props InvoiceCreationPageProps) {
    @FormTemplate(FormTemplateProps{
        Columns:     2,
        Responsive:  true,
        HeaderArea:  true,
        ToolbarArea: true,
        ContentArea: true,
        SidebarArea: true,
        FooterArea:  true,
        Bordered:    true,
        Elevated:    true,
    }) {
        // Inject invoice-specific content
        
        // Header slot
        @invoiceFormHeader(props)
        
        // Toolbar slot
        @invoiceFormToolbar(props)
        
        // Content slot
        @invoiceFormContent(props)
        
        // Sidebar slot
        @invoiceFormSidebar(props)
        
        // Footer slot
        @invoiceFormFooter(props)
    }
}

templ invoiceFormContent(props InvoiceCreationPageProps) {
    <div slot="form-content-slot" 
         x-data={ fmt.Sprintf(`{
             invoice: %s,
             lineItems: %s,
             selectedCustomer: %s,
             taxSettings: %s,
             
             // Computed properties
             get subtotal() {
                 return this.lineItems.reduce((sum, item) => 
                     sum + (item.quantity * item.unitPrice), 0);
             },
             
             get taxAmount() {
                 return this.lineItems.reduce((sum, item) => 
                     sum + (item.quantity * item.unitPrice * item.taxRate / 100), 0);
             },
             
             get total() {
                 return this.subtotal + this.taxAmount;
             },
             
             // Business logic methods
             addLineItem() {
                 this.lineItems.push({
                     id: generateId(),
                     productId: '',
                     description: '',
                     quantity: 1,
                     unitPrice: 0,
                     taxRate: this.taxSettings.defaultRate,
                     total: 0
                 });
             },
             
             removeLineItem(index) {
                 if (this.lineItems.length > 1) {
                     this.lineItems.splice(index, 1);
                 }
             },
             
             updateLineItem(index, field, value) {
                 this.lineItems[index][field] = value;
                 this.calculateLineTotal(index);
             },
             
             calculateLineTotal(index) {
                 const item = this.lineItems[index];
                 item.total = item.quantity * item.unitPrice;
             },
             
             validateInvoice() {
                 // Invoice validation logic
                 return true;
             },
             
             async saveInvoice(isDraft = false) {
                 if (!isDraft && !this.validateInvoice()) return;
                 
                 const payload = {
                     invoice: this.invoice,
                     lineItems: this.lineItems,
                     isDraft
                 };
                 
                 // Save via HTMX
             }
         }`, 
         toJSON(props.Invoice), toJSON(props.LineItems), 
         toJSON(props.Customer), toJSON(props.TaxSettings)) }>
        
        // Customer selection
        <div class="col-span-2">
            @CustomerSelector(CustomerSelectorProps{
                Selected:     props.Customer,
                Options:      props.Customers,
                Required:     true,
                AlpineModel:  "selectedCustomer",
                HXPost:       "/customers/search",
                HXTarget:     "#customer-options",
                HXTrigger:    "keyup changed delay:300ms",
            })
        </div>
        
        // Invoice details
        <div class="space-y-[var(--space-4)]">
            @FormField(FormFieldProps{
                Name:        "invoice_number",
                Label:       "Invoice Number",
                Type:        InputText,
                Value:       props.Invoice.Number,
                Required:    true,
                AlpineModel: "invoice.number",
            })
            
            @FormField(FormFieldProps{
                Name:        "invoice_date",
                Label:       "Invoice Date",
                Type:        InputDate,
                Value:       formatDate(props.Invoice.Date),
                Required:    true,
                AlpineModel: "invoice.date",
            })
            
            @FormField(FormFieldProps{
                Name:        "due_date",
                Label:       "Due Date", 
                Type:        InputDate,
                Value:       formatDate(props.Invoice.DueDate),
                Required:    true,
                AlpineModel: "invoice.dueDate",
            })
        </div>
        
        // Line items section
        <div class="col-span-2">
            <div class="space-y-[var(--space-4)]">
                <div class="flex items-center justify-between">
                    <h3 class="text-lg font-semibold">Line Items</h3>
                    @Button(ButtonProps{
                        Variant:     ButtonOutline,
                        Size:        ButtonSizeMD,
                        Icon:        "plus",
                        AlpineClick: "addLineItem()",
                    }) {
                        Add Item
                    }
                </div>
                
                // Line items table
                @InvoiceLineItemsTable(InvoiceLineItemsTableProps{
                    LineItems:    props.LineItems,
                    Products:     props.Products,
                    TaxRates:     props.TaxRates,
                    AlpineModel:  "lineItems",
                    OnUpdate:     "updateLineItem",
                    OnRemove:     "removeLineItem",
                })
            </div>
        </div>
        
        // Notes and terms
        <div class="col-span-2 space-y-[var(--space-4)]">
            @FormField(FormFieldProps{
                Name:        "notes",
                Label:       "Notes",
                Type:        InputTextarea,
                Value:       props.Invoice.Notes,
                Placeholder: "Additional notes for the customer",
                AlpineModel: "invoice.notes",
            })
            
            @FormField(FormFieldProps{
                Name:        "terms",
                Label:       "Payment Terms",
                Type:        InputTextarea,
                Value:       props.Invoice.Terms,
                Placeholder: "Payment terms and conditions",
                AlpineModel: "invoice.terms",
            })
        </div>
    </div>
}

templ invoiceFormSidebar(props InvoiceCreationPageProps) {
    <div slot="form-sidebar-slot">
        // Invoice totals
        @InvoiceTotalsSummary(InvoiceTotalsProps{
            Subtotal:    "subtotal",
            TaxAmount:   "taxAmount", 
            Total:       "total",
            Currency:    props.Invoice.Currency,
            AlpineData:  true,
        })
        
        // Validation summary
        if len(props.Errors) > 0 {
            @ValidationSummary(ValidationSummaryProps{
                Errors: props.Errors,
                Title:  "Please correct the following errors:",
            })
        }
        
        // Warnings
        if len(props.Warnings) > 0 {
            @WarningsSummary(WarningsSummaryProps{
                Warnings: props.Warnings,
                Title:    "Please review:",
            })
        }
        
        // Quick actions
        @QuickActions(QuickActionsProps{
            Actions: []QuickAction{
                {
                    Label:   "Preview Invoice",
                    Icon:    "eye",
                    HXGet:   fmt.Sprintf("/invoices/%s/preview", props.InvoiceID),
                    HXTarget: "#preview-modal",
                },
                {
                    Label:   "Duplicate Invoice",
                    Icon:    "copy",
                    HXPost:  fmt.Sprintf("/invoices/%s/duplicate", props.InvoiceID),
                    HXTarget: "#main-content",
                },
            },
        })
    </div>
}

templ invoiceFormFooter(props InvoiceCreationPageProps) {
    <div slot="form-footer-slot">
        // Cancel button
        <div>
            @Button(ButtonProps{
                Variant: ButtonOutline,
                Size:    ButtonSizeLG,
                Type:    "button",
                HXGet:   "/invoices",
                HXTarget: "#main-content",
            }) {
                Cancel
            }
        </div>
        
        // Save actions
        <div class="flex items-center gap-[var(--space-4)]">
            // Save as draft
            @Button(ButtonProps{
                Variant:     ButtonSecondary,
                Size:        ButtonSizeLG,
                Icon:        "save",
                AlpineClick: "saveInvoice(true)",
                Loading:     true,
                Class:       "x-bind:class=\"saveState === 'saving' ? 'pointer-events-none opacity-50' : ''\"",
            }) {
                Save Draft
            }
            
            // Save and send
            @Button(ButtonProps{
                Variant:     ButtonPrimary,
                Size:        ButtonSizeLG,
                Icon:        "send",
                AlpineClick: "saveInvoice(false)",
                Loading:     true,
                Class:       "x-bind:disabled=\"!validateInvoice()\" x-bind:class=\"saveState === 'saving' ? 'pointer-events-none opacity-50' : ''\"",
            }) {
                { if props.Mode == FormModeEdit { "Update" } else { "Create" } } & Send
            }
        </div>
    </div>
}
```

### 3. List Pages

#### Customer Management Page

Complete customer management interface with search, filtering, and bulk operations.

```go
type CustomerManagementPageProps struct {
    // Data context
    Customers    []Customer
    TotalCount   int
    
    // Filter and search state
    SearchTerm   string
    Filters      CustomerFilters
    SortBy       string
    SortOrder    string
    
    // Pagination
    Page         int
    PageSize     int
    
    // Business context
    UserID       string
    Permissions  []string
    
    // Schema context
    Schema       *schema.Schema
    ListSchema   *schema.Schema
    
    // Page state
    SelectedCustomers []string
    BulkActions      []BulkAction
    ViewMode         string // table, grid, cards
}

type Customer struct {
    ID           string
    Name         string
    Email        string
    Phone        string
    Company      string
    Status       CustomerStatus
    TotalOrders  int
    TotalRevenue decimal.Decimal
    LastOrder    *time.Time
    CreatedAt    time.Time
}

type CustomerFilters struct {
    Status      []CustomerStatus
    Company     []string
    DateRange   *DateRange
    RevenueRange *RevenueRange
}
```

**Implementation**:

```go
templ CustomerManagementPage(props CustomerManagementPageProps) {
    @MasterDetailTemplate(MasterDetailTemplateProps{
        Layout:       MasterDetailSplit,
        SplitRatio:   "2:1",
        HeaderArea:   true,
        FilterArea:   true,
        ListArea:     true,
        DetailArea:   true,
        ActionsArea:  true,
        Responsive:   true,
    }) {
        // Header slot
        @customerPageHeader(props)
        
        // Filters slot
        @customerFilters(props)
        
        // List slot
        @customerList(props)
        
        // Detail slot (injected via HTMX when customer is selected)
        @customerDetailPlaceholder()
    }
}

templ customerPageHeader(props CustomerManagementPageProps) {
    <div slot="page-header-slot">
        <div class="flex items-center justify-between">
            <div>
                <h1 class="text-2xl font-bold text-[var(--foreground)]">Customers</h1>
                <p class="text-[var(--muted-foreground)]">
                    { strconv.Itoa(props.TotalCount) } total customers
                </p>
            </div>
            
            <div class="flex items-center gap-[var(--space-4)]">
                // View mode toggle
                @ViewModeToggle(ViewModeToggleProps{
                    Options: []ViewModeOption{
                        {Value: "table", Icon: "table", Label: "Table"},
                        {Value: "grid", Icon: "grid", Label: "Grid"},
                        {Value: "cards", Icon: "card", Label: "Cards"},
                    },
                    Selected:    props.ViewMode,
                    HXPost:      "/customers/view-mode",
                    HXTarget:    "#customer-list",
                })
                
                // Actions
                if hasPermission(props.Permissions, "customers.create") {
                    @Button(ButtonProps{
                        Variant:  ButtonPrimary,
                        Size:     ButtonSizeMD,
                        Icon:     "plus",
                        HXGet:    "/customers/new",
                        HXTarget: "#main-content",
                    }) {
                        New Customer
                    }
                }
                
                @Button(ButtonProps{
                    Variant:  ButtonOutline,
                    Size:     ButtonSizeMD,
                    Icon:     "download",
                    HXPost:   "/customers/export",
                    HXTrigger: "click",
                }) {
                    Export
                }
            </div>
        </div>
    </div>
}

templ customerFilters(props CustomerManagementPageProps) {
    <div slot="filters-slot" 
         x-data={ fmt.Sprintf(`{
             searchTerm: '%s',
             filters: %s,
             viewMode: '%s',
             
             updateFilters() {
                 // Auto-submit filters via HTMX
                 this.$el.querySelector('form').dispatchEvent(new Event('submit'));
             },
             
             clearFilters() {
                 this.searchTerm = '';
                 this.filters = {};
                 this.updateFilters();
             }
         }`, props.SearchTerm, toJSON(props.Filters), props.ViewMode) }>
        
        <form 
            hx-post="/customers/filter"
            hx-target="#customer-list"
            hx-trigger="submit"
            x-on:submit.prevent="$el.submit()"
            class="grid grid-cols-1 md:grid-cols-4 gap-[var(--space-4)]"
        >
            // Search
            @SearchBar(SearchBarProps{
                Name:        "search",
                Value:       props.SearchTerm,
                Placeholder: "Search customers...",
                AlpineModel: "searchTerm",
                HXPost:      "/customers/search",
                HXTarget:    "#customer-list",
                HXTrigger:   "keyup changed delay:300ms",
            })
            
            // Status filter
            @MultiSelect(MultiSelectProps{
                Name:        "status",
                Label:       "Status",
                Options:     getCustomerStatusOptions(),
                Selected:    statusSliceToString(props.Filters.Status),
                AlpineModel: "filters.status",
                Placeholder: "All statuses",
            })
            
            // Date range filter
            @DateRangePicker(DateRangePickerProps{
                Name:        "date_range",
                Label:       "Created Date",
                Value:       props.Filters.DateRange,
                AlpineModel: "filters.dateRange",
                Placeholder: "All time",
            })
            
            // Clear filters
            <div class="flex items-end">
                @Button(ButtonProps{
                    Variant:     ButtonGhost,
                    Size:        ButtonSizeMD,
                    Icon:        "x",
                    Type:        "button",
                    AlpineClick: "clearFilters()",
                }) {
                    Clear
                }
            </div>
        </form>
    </div>
}

templ customerList(props CustomerManagementPageProps) {
    <div slot="list-slot">
        <div 
            id="customer-list"
            x-data={ fmt.Sprintf(`{
                selectedCustomers: %s,
                selectAll: false,
                
                toggleCustomer(customerId) {
                    const index = this.selectedCustomers.indexOf(customerId);
                    if (index > -1) {
                        this.selectedCustomers.splice(index, 1);
                    } else {
                        this.selectedCustomers.push(customerId);
                    }
                    this.updateSelectAll();
                },
                
                toggleSelectAll() {
                    if (this.selectAll) {
                        this.selectedCustomers = [];
                    } else {
                        this.selectedCustomers = %s;
                    }
                    this.selectAll = !this.selectAll;
                },
                
                updateSelectAll() {
                    this.selectAll = this.selectedCustomers.length === %d;
                },
                
                selectCustomer(customerId) {
                    // Load customer details
                    htmx.ajax('GET', '/customers/' + customerId + '/details', {
                        target: '#customer-detail'
                    });
                }
            }`, toJSON(props.SelectedCustomers), 
            toJSON(getCustomerIds(props.Customers)), len(props.Customers)) }
        >
            // Bulk actions bar
            if len(props.BulkActions) > 0 {
                <div 
                    class="mb-[var(--space-4)] p-[var(--space-4)] bg-[var(--accent)] rounded-[var(--radius-md)]"
                    x-show="selectedCustomers.length > 0"
                    x-transition
                >
                    <div class="flex items-center justify-between">
                        <span class="text-[var(--accent-foreground)]" x-text="`${selectedCustomers.length} customers selected`"></span>
                        
                        <div class="flex items-center gap-[var(--space-2)]">
                            for _, action := range props.BulkActions {
                                @Button(ButtonProps{
                                    Variant:     action.Variant,
                                    Size:        ButtonSizeSM,
                                    Icon:        action.Icon,
                                    HXPost:      action.Endpoint,
                                    HXTarget:    "#customer-list",
                                    AlpineClick: if action.Confirmation != "" {
                                        fmt.Sprintf("if (confirm('%s')) { /* continue */ }", action.Confirmation)
                                    } else { "" },
                                }) {
                                    { action.Label }
                                }
                            }
                        </div>
                    </div>
                </div>
            }
            
            // Customer table/grid based on view mode
            if props.ViewMode == "table" {
                @CustomerTable(CustomerTableProps{
                    Customers:         props.Customers,
                    SelectedCustomers: props.SelectedCustomers,
                    SortBy:           props.SortBy,
                    SortOrder:        props.SortOrder,
                    Permissions:      props.Permissions,
                    OnSelect:         "selectCustomer",
                    OnToggleSelect:   "toggleCustomer",
                    OnToggleAll:      "toggleSelectAll",
                })
            } else {
                @CustomerGrid(CustomerGridProps{
                    Customers:         props.Customers,
                    SelectedCustomers: props.SelectedCustomers,
                    ViewMode:         props.ViewMode,
                    Permissions:      props.Permissions,
                    OnSelect:         "selectCustomer",
                    OnToggleSelect:   "toggleCustomer",
                })
            }
            
            // Pagination
            @Pagination(PaginationProps{
                CurrentPage:  props.Page,
                TotalPages:   getTotalPages(props.TotalCount, props.PageSize),
                PageSize:     props.PageSize,
                TotalItems:   props.TotalCount,
                HXGet:        "/customers",
                HXTarget:     "#customer-list",
                ShowSizeSelector: true,
            })
        </div>
    </div>
}
```

## Page Design Principles

### 1. Real Data Integration

Pages work with actual business data and handle all data states:

```go
// Page handles all data states
templ CustomerManagementPage(props CustomerManagementPageProps) {
    <div x-data="{
        loading: false,
        error: null,
        
        async loadCustomers() {
            this.loading = true;
            try {
                // Load real data
                const response = await fetch('/api/customers');
                const data = await response.json();
                // Update UI with real data
            } catch (err) {
                this.error = err.message;
            } finally {
                this.loading = false;
            }
        }
    }">
        // Handle loading states
        <div x-show="loading">@LoadingSpinner()</div>
        <div x-show="error">@ErrorMessage()</div>
        <div x-show="!loading && !error">
            // Main content
        </div>
    </div>
}
```

### 2. Complete User Experience

Pages provide complete, end-to-end user experiences:

```go
// Complete invoice page with all user flows
templ InvoiceDetailPage(props InvoiceDetailPageProps) {
    // Page handles all possible user actions
    <div x-data="{
        // All possible states
        mode: 'view', // view, edit, print, email
        
        // All possible actions
        editInvoice() { this.mode = 'edit'; },
        saveInvoice() { /* save logic */ },
        printInvoice() { /* print logic */ },
        emailInvoice() { /* email logic */ },
        duplicateInvoice() { /* duplicate logic */ },
        deleteInvoice() { /* delete logic */ }
    }">
        // Complete interface for all actions
    </div>
}
```

### 3. Performance Optimization

Pages implement performance optimizations for real-world usage:

```go
// Page with performance optimizations
templ DashboardPage(props DashboardPageProps) {
    <div x-data="{
        // Lazy loading
        chartsLoaded: false,
        
        // Refresh strategies  
        refreshInterval: 30000,
        
        // Memory management
        cleanup() {
            // Clean up resources
        },
        
        init() {
            // Intersection observer for lazy loading
            const observer = new IntersectionObserver((entries) => {
                if (entries[0].isIntersecting && !this.chartsLoaded) {
                    this.loadCharts();
                }
            });
            
            // Auto refresh
            setInterval(() => this.refreshData(), this.refreshInterval);
        }
    }">
        // Optimized page content
    </div>
}
```

## Testing Pages

```go
func TestExecutiveDashboardPage_DataIntegration(t *testing.T) {
    props := ExecutiveDashboardPageProps{
        UserID:   "user123",
        TenantID: "tenant456",
        Metrics: []DashboardMetric{
            {
                ID:    "revenue",
                Title: "Total Revenue",
                Value: 125000.00,
                Trend: TrendUp,
            },
        },
        Permissions: []string{"dashboard.view", "metrics.revenue"},
    }
    
    html := renderToString(ExecutiveDashboardPage(props))
    
    // Test data integration
    assert.Contains(t, html, "125000.00")
    assert.Contains(t, html, "Total Revenue")
    
    // Test permission handling
    assert.Contains(t, html, "hasPermission")
}

func TestInvoiceCreationPage_BusinessLogic(t *testing.T) {
    props := InvoiceCreationPageProps{
        Mode: FormModeCreate,
        Invoice: &Invoice{
            Currency: "USD",
        },
        LineItems: []LineItem{
            {
                Quantity:  2,
                UnitPrice: decimal.NewFromFloat(100.00),
                TaxRate:   decimal.NewFromFloat(10.0),
            },
        },
    }
    
    html := renderToString(InvoiceCreationPage(props))
    
    // Test business logic integration
    assert.Contains(t, html, "calculateLineTotal")
    assert.Contains(t, html, "validateInvoice")
    assert.Contains(t, html, "get total()")
}
```

## Best Practices

### ✅ DO

- Connect to real data sources
- Handle all loading and error states
- Implement complete user workflows
- Optimize for performance and usability
- Test with realistic data scenarios
- Implement proper error handling
- Provide comprehensive user feedback

### ❌ DON'T

- Use mock or placeholder data in production pages
- Ignore error and edge cases
- Create incomplete user experiences
- Skip performance optimization
- Forget accessibility requirements
- Ignore mobile responsiveness

## Related Documentation

- **[Templates](./05-templates.md)** - Structural patterns for pages
- **[Integration](./07-integration.md)** - Schema-component integration
- **[Schema Workflows](../schema/workflow.md)** - Backend workflow system
- **[Performance Guide](../performance.md)** - Page optimization strategies

---

Pages represent the complete user-facing implementations that combine all atomic design elements with real data, business logic, and comprehensive user experiences.