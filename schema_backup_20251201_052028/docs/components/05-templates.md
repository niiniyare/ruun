# Templates

**Page layout structures that define content arrangement patterns without specific data**

## Definition

Templates are wireframes that define the structural skeleton of pages. They establish layout patterns, content areas, and component relationships without containing actual content. Templates focus on arrangement, spacing, and responsive behavior.

## Key Characteristics

- **Layout-focused** - defines structural patterns and content areas
- **Content-agnostic** - works with different data sets and content types
- **Responsive design** - adapts to different screen sizes and devices
- **Reusable structures** - same template serves multiple similar pages
- **Accessibility foundation** - establishes proper semantic structure
- **Performance optimized** - efficient layout patterns for fast rendering

## Core Template Categories

### 1. Dashboard Layouts

#### Executive Dashboard Template

Layout for data-heavy dashboards with metrics, charts, and activity feeds.

```go
type DashboardTemplateProps struct {
    // Layout configuration
    Layout      DashboardLayout // grid, sidebar, tabs
    Columns     int             // Number of grid columns
    Responsive  bool            // Responsive grid behavior
    
    // Content areas
    HeaderArea    bool // Top navigation/breadcrumbs
    SidebarArea   bool // Side navigation
    MetricsArea   bool // Key metrics display
    ChartsArea    bool // Data visualization
    ActivityArea  bool // Recent activity feed
    AlertsArea    bool // Notifications/alerts
    
    // Responsive behavior
    SidebarCollapsible bool
    MobileLayout       string // stacked, tabs, accordion
    
    // Schema integration
    Schema      *schema.Schema
    LayoutRules map[string]*schema.LayoutRule
    
    // Styling
    Spacing     SpacingLevel // compact, normal, spacious
    Theme       string       // light, dark, auto
    Class       string
}

type DashboardLayout string

const (
    DashboardLayoutGrid    DashboardLayout = "grid"
    DashboardLayoutSidebar DashboardLayout = "sidebar"
    DashboardLayoutTabs    DashboardLayout = "tabs"
)

type SpacingLevel string

const (
    SpacingCompact  SpacingLevel = "compact"
    SpacingNormal   SpacingLevel = "normal"
    SpacingSpacious SpacingLevel = "spacious"
)
```

**Implementation**:

```go
func dashboardTemplateClasses(props DashboardTemplateProps) string {
    var classes []string
    
    // Base layout classes
    classes = append(classes,
        "min-h-screen",
        "bg-[var(--background)]",
        "text-[var(--foreground)]",
    )
    
    // Layout-specific classes
    switch props.Layout {
    case DashboardLayoutGrid:
        classes = append(classes, "grid", fmt.Sprintf("grid-cols-%d", props.Columns), "gap-[var(--space-6)]")
    case DashboardLayoutSidebar:
        classes = append(classes, "flex", "min-h-screen")
    case DashboardLayoutTabs:
        classes = append(classes, "flex", "flex-col")
    }
    
    // Spacing adjustments
    switch props.Spacing {
    case SpacingCompact:
        classes = append(classes, "space-compact")
    case SpacingSpacious:
        classes = append(classes, "space-spacious")
    default:
        classes = append(classes, "space-normal")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ DashboardTemplate(props DashboardTemplateProps) {
    <div 
        class={ dashboardTemplateClasses(props) }
        x-data={ fmt.Sprintf(`{
            sidebarOpen: true,
            currentTab: 'overview',
            isMobile: window.innerWidth < 768,
            
            init() {
                // Handle responsive behavior
                this.$watch('isMobile', value => {
                    if (value && this.sidebarOpen) {
                        this.sidebarOpen = false;
                    }
                });
                
                // Listen for window resize
                window.addEventListener('resize', () => {
                    this.isMobile = window.innerWidth < 768;
                });
            },
            
            toggleSidebar() {
                this.sidebarOpen = !this.sidebarOpen;
            }
        }`) }
    >
        // Header area (always present)
        if props.HeaderArea {
            <header class="bg-[var(--card)] border-b border-[var(--border)] px-[var(--space-6)] py-[var(--space-4)]">
                <div class="flex items-center justify-between">
                    // Left: Logo and navigation
                    <div class="flex items-center gap-[var(--space-6)]">
                        if props.SidebarArea && props.SidebarCollapsible {
                            @Button(ButtonProps{
                                Variant:     ButtonGhost,
                                Size:        ButtonSizeMD,
                                Icon:        "menu",
                                AlpineClick: "toggleSidebar()",
                                Class:       "lg:hidden",
                            })
                        }
                        
                        // Breadcrumbs slot
                        <div id="breadcrumbs-slot">
                            <!-- Breadcrumbs will be injected here -->
                        </div>
                    </div>
                    
                    // Right: User menu and notifications
                    <div class="flex items-center gap-[var(--space-4)]">
                        <div id="header-actions-slot">
                            <!-- Header actions will be injected here -->
                        </div>
                        
                        <div id="notifications-slot">
                            <!-- Notifications will be injected here -->
                        </div>
                        
                        <div id="user-menu-slot">
                            <!-- User menu will be injected here -->
                        </div>
                    </div>
                </div>
            </header>
        }
        
        // Main content area
        <div class={ "flex-1" + if props.Layout == DashboardLayoutSidebar { " flex" } else { "" } }>
            // Sidebar area
            if props.SidebarArea {
                <aside 
                    class="bg-[var(--card)] border-r border-[var(--border)] transition-all duration-200"
                    x-bind:class="sidebarOpen ? 'w-64' : 'w-0 lg:w-16'"
                    x-show="sidebarOpen || !isMobile"
                    x-transition:enter="transition ease-out duration-200"
                    x-transition:enter-start="transform -translate-x-full"
                    x-transition:enter-end="transform translate-x-0"
                >
                    <div id="sidebar-slot" class="h-full">
                        <!-- Sidebar navigation will be injected here -->
                    </div>
                </aside>
            }
            
            // Main content
            <main class="flex-1 overflow-auto">
                // Metrics area
                if props.MetricsArea {
                    <section class="p-[var(--space-6)] border-b border-[var(--border)]">
                        <div id="metrics-slot" class={ getMetricsGridClasses(props.Columns) }>
                            <!-- Metrics cards will be injected here -->
                        </div>
                    </section>
                }
                
                // Content grid for charts and activities
                <div class="p-[var(--space-6)] grid gap-[var(--space-6)]" 
                     style={ getContentGridStyle(props.ChartsArea, props.ActivityArea) }>
                    
                    // Charts area
                    if props.ChartsArea {
                        <section class="bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] p-[var(--space-6)]">
                            <div id="charts-slot">
                                <!-- Charts and visualizations will be injected here -->
                            </div>
                        </section>
                    }
                    
                    // Activity/Alerts area
                    <div class="space-y-[var(--space-6)]">
                        if props.ActivityArea {
                            <section class="bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] p-[var(--space-6)]">
                                <div id="activity-slot">
                                    <!-- Activity feed will be injected here -->
                                </div>
                            </section>
                        }
                        
                        if props.AlertsArea {
                            <section class="bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] p-[var(--space-6)]">
                                <div id="alerts-slot">
                                    <!-- Alerts and notifications will be injected here -->
                                </div>
                            </section>
                        }
                    </div>
                </div>
            </main>
        </div>
    </div>
}

// Helper functions for responsive grid
func getMetricsGridClasses(columns int) string {
    return fmt.Sprintf("grid grid-cols-1 md:grid-cols-2 lg:grid-cols-%d gap-[var(--space-4)]", 
                      min(columns, 4))
}

func getContentGridStyle(hasCharts, hasActivity bool) string {
    if hasCharts && hasActivity {
        return "grid-template-columns: 2fr 1fr;"
    } else if hasCharts || hasActivity {
        return "grid-template-columns: 1fr;"
    }
    return ""
}
```

### 2. Form Layouts

#### Multi-Column Form Template

Responsive form layout with multiple columns and section groupings.

```go
type FormTemplateProps struct {
    // Layout configuration
    Columns      int  // Number of form columns
    Responsive   bool // Responsive column behavior
    Sectioned    bool // Group fields into sections
    
    // Form areas
    HeaderArea   bool // Form title and description
    ToolbarArea  bool // Form actions and tools
    ContentArea  bool // Main form fields
    SidebarArea  bool // Additional info/help
    FooterArea   bool // Submit/cancel actions
    
    // Responsive behavior
    MobileStack  bool // Stack on mobile
    TabletColumns int // Columns on tablet
    
    // Schema integration
    Schema       *schema.Schema
    FormLayout   *schema.FormLayout
    
    // Styling
    Spacing      SpacingLevel
    Bordered     bool
    Elevated     bool
    Class        string
}
```

**Implementation**:

```go
templ FormTemplate(props FormTemplateProps) {
    <div 
        class={ getFormTemplateClasses(props) }
        x-data={ fmt.Sprintf(`{
            currentSection: 0,
            formData: {},
            validationErrors: {},
            isSubmitting: false,
            
            get formSections() {
                return document.querySelectorAll('[data-form-section]');
            },
            
            validateSection(sectionIndex) {
                // Section validation logic
                return true;
            },
            
            nextSection() {
                if (this.currentSection < this.formSections.length - 1) {
                    this.currentSection++;
                }
            },
            
            prevSection() {
                if (this.currentSection > 0) {
                    this.currentSection--;
                }
            }
        }`) }
    >
        // Form header
        if props.HeaderArea {
            <header class="mb-[var(--space-8)]">
                <div id="form-header-slot">
                    <!-- Form title, description, and progress will be injected here -->
                </div>
            </header>
        }
        
        // Form toolbar
        if props.ToolbarArea {
            <div class="mb-[var(--space-6)] p-[var(--space-4)] bg-[var(--muted)] rounded-[var(--radius-lg)]">
                <div id="form-toolbar-slot" class="flex items-center justify-between">
                    <!-- Form tools and actions will be injected here -->
                </div>
            </div>
        }
        
        // Main form layout
        <div class={ getMainFormClasses(props) }>
            // Form content
            <div class="flex-1">
                <form class={ getFormGridClasses(props) }>
                    <div id="form-content-slot" class="contents">
                        <!-- Form fields will be injected here -->
                    </div>
                </form>
            </div>
            
            // Form sidebar
            if props.SidebarArea {
                <aside class="w-80 space-y-[var(--space-6)]">
                    <div id="form-sidebar-slot">
                        <!-- Help, validation summary, etc. will be injected here -->
                    </div>
                </aside>
            }
        </div>
        
        // Form footer
        if props.FooterArea {
            <footer class={ "mt-[var(--space-8)] pt-[var(--space-6)]" +
                           if props.Bordered { " border-t border-[var(--border)]" } else { "" } }>
                <div id="form-footer-slot" class="flex items-center justify-between">
                    <!-- Form actions will be injected here -->
                </div>
            </footer>
        }
    </div>
}

func getFormTemplateClasses(props FormTemplateProps) string {
    var classes []string
    
    classes = append(classes, "max-w-5xl", "mx-auto", "p-[var(--space-6)]")
    
    if props.Elevated {
        classes = append(classes, "bg-[var(--card)]", "border", "border-[var(--border)]", 
                        "rounded-[var(--radius-lg)]", "shadow-[var(--shadow-lg)]")
    }
    
    return strings.Join(classes, " ")
}

func getMainFormClasses(props FormTemplateProps) string {
    var classes []string
    
    if props.SidebarArea {
        classes = append(classes, "flex", "gap-[var(--space-8)]")
    }
    
    return strings.Join(classes, " ")
}

func getFormGridClasses(props FormTemplateProps) string {
    var classes []string
    
    classes = append(classes, "grid", "gap-[var(--space-6)]")
    
    if props.Responsive {
        switch props.Columns {
        case 1:
            classes = append(classes, "grid-cols-1")
        case 2:
            classes = append(classes, "grid-cols-1", "lg:grid-cols-2")
        case 3:
            classes = append(classes, "grid-cols-1", "md:grid-cols-2", "lg:grid-cols-3")
        default:
            classes = append(classes, "grid-cols-1", "md:grid-cols-2", "lg:grid-cols-3", "xl:grid-cols-4")
        }
    } else {
        classes = append(classes, fmt.Sprintf("grid-cols-%d", props.Columns))
    }
    
    return strings.Join(classes, " ")
}
```

### 3. List and Detail Layouts

#### Master-Detail Template

Layout for list/detail views with optional split-screen or modal detail.

```go
type MasterDetailTemplateProps struct {
    // Layout configuration
    Layout       MasterDetailLayout // split, modal, page
    SplitRatio   string            // e.g., "1:2", "1:1"
    Responsive   bool              // Stack on mobile
    
    // Areas
    HeaderArea   bool // Page header
    FilterArea   bool // Filters and search
    ListArea     bool // Master list
    DetailArea   bool // Detail content
    ActionsArea  bool // Detail actions
    
    // Behavior
    SelectionMode string // single, multiple, none
    DetailMode    string // overlay, inline, page
    
    // Schema integration
    Schema        *schema.Schema
    ListSchema    *schema.Schema
    DetailSchema  *schema.Schema
    
    // Styling
    ListMinWidth  string
    DetailMinWidth string
    Class         string
}

type MasterDetailLayout string

const (
    MasterDetailSplit MasterDetailLayout = "split"
    MasterDetailModal MasterDetailLayout = "modal"
    MasterDetailPage  MasterDetailLayout = "page"
)
```

**Implementation**:

```go
templ MasterDetailTemplate(props MasterDetailTemplateProps) {
    <div 
        class={ getMasterDetailClasses(props) }
        x-data={ fmt.Sprintf(`{
            selectedItem: null,
            detailVisible: false,
            isMobile: window.innerWidth < 768,
            
            selectItem(item) {
                this.selectedItem = item;
                if ('%s' === 'modal' || this.isMobile) {
                    this.detailVisible = true;
                }
            },
            
            closeDetail() {
                this.detailVisible = false;
                if ('%s' === 'page') {
                    this.selectedItem = null;
                }
            },
            
            init() {
                window.addEventListener('resize', () => {
                    this.isMobile = window.innerWidth < 768;
                    if (this.isMobile && this.detailVisible && '%s' === 'split') {
                        this.detailVisible = true; // Convert to modal on mobile
                    }
                });
            }
        }`, props.Layout, props.Layout, props.Layout) }
    >
        // Page header
        if props.HeaderArea {
            <header class="mb-[var(--space-6)]">
                <div id="page-header-slot">
                    <!-- Page title, breadcrumbs will be injected here -->
                </div>
            </header>
        }
        
        // Filters area
        if props.FilterArea {
            <section class="mb-[var(--space-6)]">
                <div id="filters-slot" class="bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] p-[var(--space-4)]">
                    <!-- Search and filters will be injected here -->
                </div>
            </section>
        }
        
        // Main content layout
        <div class={ getMainContentClasses(props) }>
            // Master list
            if props.ListArea {
                <section class={ getListSectionClasses(props) }>
                    <div id="list-slot" class="h-full">
                        <!-- Master list will be injected here -->
                    </div>
                </section>
            }
            
            // Detail content (for split layout)
            if props.DetailArea && props.Layout == MasterDetailSplit {
                <section 
                    class={ getDetailSectionClasses(props) }
                    x-show="selectedItem || !isMobile"
                    x-transition:enter="transition ease-out duration-200"
                    x-transition:enter-start="transform translate-x-full opacity-0"
                    x-transition:enter-end="transform translate-x-0 opacity-100"
                >
                    <div class="bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] h-full">
                        // Detail header
                        <header class="p-[var(--space-6)] border-b border-[var(--border)]">
                            <div class="flex items-center justify-between">
                                <div id="detail-header-slot">
                                    <!-- Detail title will be injected here -->
                                </div>
                                
                                if props.ActionsArea {
                                    <div id="detail-actions-slot">
                                        <!-- Detail actions will be injected here -->
                                    </div>
                                }
                            </div>
                        </header>
                        
                        // Detail content
                        <div class="p-[var(--space-6)]">
                            <div id="detail-content-slot">
                                <!-- Detail content will be injected here -->
                            </div>
                        </div>
                    </div>
                </section>
            }
        </div>
        
        // Modal detail (for modal layout)
        if props.DetailArea && props.Layout == MasterDetailModal {
            <div 
                x-show="detailVisible"
                x-transition:enter="transition ease-out duration-300"
                x-transition:enter-start="opacity-0"
                x-transition:enter-end="opacity-100"
                class="fixed inset-0 z-50 overflow-y-auto"
            >
                // Modal backdrop
                <div class="fixed inset-0 bg-black bg-opacity-50" x-on:click="closeDetail()"></div>
                
                // Modal content
                <div class="relative min-h-screen flex items-center justify-center p-4">
                    <div 
                        class="relative bg-[var(--card)] border border-[var(--border)] rounded-[var(--radius-lg)] shadow-[var(--shadow-xl)] max-w-4xl w-full max-h-[90vh] overflow-hidden"
                        x-on:click.away="closeDetail()"
                    >
                        // Modal header
                        <header class="p-[var(--space-6)] border-b border-[var(--border)]">
                            <div class="flex items-center justify-between">
                                <div id="modal-detail-header-slot">
                                    <!-- Modal detail title will be injected here -->
                                </div>
                                
                                <div class="flex items-center gap-[var(--space-2)]">
                                    if props.ActionsArea {
                                        <div id="modal-detail-actions-slot">
                                            <!-- Modal detail actions will be injected here -->
                                        </div>
                                    }
                                    
                                    @Button(ButtonProps{
                                        Variant:     ButtonGhost,
                                        Size:        ButtonSizeSM,
                                        Icon:        "x",
                                        AlpineClick: "closeDetail()",
                                    })
                                </div>
                            </div>
                        </header>
                        
                        // Modal content
                        <div class="p-[var(--space-6)] overflow-y-auto">
                            <div id="modal-detail-content-slot">
                                <!-- Modal detail content will be injected here -->
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        }
    </div>
}
```

## Template Design Principles

### 1. Content Agnostic

Templates should work with any appropriate content:

**✅ Good Example**:
```go
// Generic dashboard template that works with any metrics
templ DashboardTemplate(props DashboardTemplateProps) {
    <div class="dashboard-grid">
        <section id="metrics-slot" class="metrics-area">
            <!-- Any metrics content can be inserted here -->
        </section>
        <section id="charts-slot" class="charts-area">  
            <!-- Any charts can be inserted here -->
        </section>
    </div>
}
```

### 2. Responsive by Default

Templates should handle responsive behavior automatically:

```go
// Responsive grid that adapts to screen size
func getResponsiveGridClasses(columns int) string {
    switch columns {
    case 1:
        return "grid grid-cols-1"
    case 2: 
        return "grid grid-cols-1 lg:grid-cols-2"
    case 3:
        return "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3"
    default:
        return "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
    }
}
```

### 3. Slot-Based Architecture

Templates use slots for flexible content injection:

```go
// Template with multiple content slots
templ PageTemplate(props PageTemplateProps) {
    <div class="page-layout">
        <header id="header-slot"></header>
        <nav id="navigation-slot"></nav>
        <main id="content-slot"></main>
        <aside id="sidebar-slot"></aside>
        <footer id="footer-slot"></footer>
    </div>
}
```

## Schema Integration

Templates can be driven by schema layout definitions:

```go
// Schema-driven template rendering
func renderSchemaTemplate(schema *schema.Schema) templ.Component {
    layout := schema.Layout
    
    props := DashboardTemplateProps{
        Layout:      DashboardLayout(layout.Type),
        Columns:     layout.Columns,
        HeaderArea:  layout.Areas.Header,
        SidebarArea: layout.Areas.Sidebar,
        MetricsArea: layout.Areas.Metrics,
    }
    
    return DashboardTemplate(props)
}
```

## Testing Templates

```go
func TestDashboardTemplate_ResponsiveLayout(t *testing.T) {
    props := DashboardTemplateProps{
        Layout:      DashboardLayoutGrid,
        Columns:     3,
        Responsive:  true,
        HeaderArea:  true,
        MetricsArea: true,
    }
    
    html := renderToString(DashboardTemplate(props))
    
    // Test responsive classes
    assert.Contains(t, html, "grid-cols-1")
    assert.Contains(t, html, "lg:grid-cols-3")
    
    // Test content slots
    assert.Contains(t, html, "id=\"metrics-slot\"")
    assert.Contains(t, html, "id=\"header-slot\"")
}
```

## Best Practices

### ✅ DO

- Create reusable layout patterns
- Design for responsive behavior
- Use semantic HTML structure
- Implement accessible navigation
- Provide flexible content slots
- Test across different screen sizes

### ❌ DON'T

- Include specific content in templates  
- Hard-code layout dimensions
- Ignore mobile experience
- Create templates that are too specific
- Skip accessibility considerations

## Related Documentation

- **[Organisms](./04-organisms.md)** - Complex components for templates
- **[Pages](./06-pages.md)** - Complete page implementations  
- **[Schema Layout](../schema/layout.md)** - Backend layout system
- **[Responsive Design](../styles/responsive.md)** - Responsive design system

---

Templates provide the structural foundation for pages, defining layout patterns and content arrangement without being tied to specific data or business logic.