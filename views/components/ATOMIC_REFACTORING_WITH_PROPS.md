# Atomic Design Refactoring Strategy with Props Pattern

## Executive Summary

Comprehensive refactoring strategy using strongly-typed props across atomic design hierarchy with HTMX/Alpine.js integration for type safety, maintainability, and clear component contracts.

**Props Flow**: `DESIGN TOKENS → ATOMS → MOLECULES → ORGANISMS → TEMPLATES → PAGES`

---

## 1. Atoms Layer - Foundation Props

### Core Atom Structure

```go
// views/components/atoms/button.templ
type ButtonProps struct {
    // Identity & Core
    ID          string
    Variant     ButtonVariant  // primary, secondary, destructive, outline, ghost, link
    Size        ButtonSize     // xs, sm, md, lg, xl
    Type        ButtonType     // button, submit, reset
    
    // States
    Disabled    bool
    Loading     bool
    Active      bool
    
    // Styling
    FullWidth   bool
    ClassName   string
    
    // HTMX
    HXGet, HXPost, HXPut, HXDelete, HXPatch string
    HXTarget, HXSwap, HXTrigger string
    HXPushURL, HXConfirm, HXIndicator string
    HXHeaders   map[string]string
    
    // Alpine.js
    XData       string
    XShow       string
    XBind       map[string]string
    XOn         map[string]string
    XText, XHTML, XRef string
    
    // Accessibility
    AriaLabel, AriaControls string
    AriaPressed, AriaExpanded bool
    
    // Data
    DataTestID      string
    DataAttributes  map[string]string
}

templ Button(props ButtonProps) {
    <button { buildButtonAttributes(props)... }>
        if props.Loading {
            @Spinner(SpinnerProps{Size: mapButtonSizeToSpinner(props.Size)})
        }
        { children... }
    </button>
}

func buildButtonAttributes(props ButtonProps) templ.Attributes {
    return utils.MergeAttributes(
        templ.Attributes{
            "id":   utils.IfElse(props.ID != "", props.ID, utils.RandomID()),
            "type": string(utils.IfElse(props.Type != "", props.Type, ButtonTypeButton)),
            "class": utils.TwMerge(
                "btn",
                fmt.Sprintf("btn-%s btn-%s", props.Variant, props.Size),
                utils.If(props.Active, "active"),
                utils.If(props.FullWidth, "w-full"),
                utils.If(props.Loading, "loading"),
                props.ClassName,
            ),
            "disabled": props.Disabled || props.Loading,
        },
        buildHTMXAttrs(props),
        buildAlpineAttrs(props),
        buildAriaAttrs(props),
        props.DataAttributes,
    )
}
```

### Other Essential Atoms

```go
// Input with validation
type InputProps struct {
    ID, Name, Type, Value, Placeholder string
    Required, Disabled, ReadOnly, Invalid bool
    Pattern string
    MinLength, MaxLength int
    Size InputSize
    
    // HTMX validation
    HXPost, HXTrigger, HXTarget string
    HXValidate bool
    
    // Alpine binding
    XModel string
    XOn    map[string]string
    
    AriaLabel, AriaDescribedBy string
}

// Label, Icon, Spinner follow similar pattern with focused props
```

---

## 2. Molecules Layer - Composition Props

### Form Field Molecule

```go
type FormFieldProps struct {
    ID, Name, Label string
    Type        atoms.InputType
    Value       string
    InputSize   atoms.InputSize
    Layout      FormFieldLayout  // vertical, horizontal, floating
    
    // Validation
    Required     bool
    Invalid      bool
    ErrorMessage string
    HelpText     string
    
    ClassName   string
}

templ FormField(props FormFieldProps) {
    <div class={ formFieldClasses(props) }>
        @atoms.Label(atoms.LabelProps{For: props.ID, Required: props.Required}) {
            { props.Label }
        }
        
        @atoms.Input(atoms.InputProps{
            ID: props.ID,
            Name: props.Name,
            Type: props.Type,
            Value: props.Value,
            Size: props.InputSize,
            Invalid: props.Invalid,
            AriaDescribedBy: getAriaDescribedBy(props),
        })
        
        if props.ErrorMessage != "" {
            @ErrorMessage(ErrorMessageProps{Message: props.ErrorMessage})
        }
        if props.HelpText != "" {
            @HelpText(HelpTextProps{Text: props.HelpText})
        }
    </div>
}
```

### Search Bar with HTMX Live Search

```go
type SearchBarProps struct {
    ID, Name, Value, Placeholder string
    SearchEndpoint string  // HTMX endpoint
    MinChars, DebounceMs int
    
    // HTMX config
    HXTarget, HXSwap, HXIndicator string
    
    // Alpine state
    XData string
    
    ShowButton bool
    ButtonText string
    Size ComponentSize
}

templ SearchBar(props SearchBarProps) {
    <form x-data={ getSearchAlpineData(props) }>
        @atoms.Input(atoms.InputProps{
            ID: props.ID,
            Type: atoms.InputSearch,
            HXPost: props.SearchEndpoint,
            HXTrigger: fmt.Sprintf("keyup changed delay:%dms", props.DebounceMs),
            HXTarget: props.HXTarget,
            XModel: "query",
        })
        
        if props.ShowButton {
            @atoms.Button(atoms.ButtonProps{
                Type: atoms.ButtonTypeSubmit,
                XShow: fmt.Sprintf("query.length >= %d", props.MinChars),
            }) { { props.ButtonText } }
        }
    </form>
}

func getSearchAlpineData(props SearchBarProps) string {
    return fmt.Sprintf(`{query: '%s', loading: false, minChars: %d}`, 
        props.Value, props.MinChars)
}
```

---

## 3. Organisms Layer - Complex Props

### Form Organism with Business Logic

```go
type FormProps struct {
    ID, Name string
    Schema   schema.FormSchema
    Values   map[string]any
    Errors   map[string][]string
    
    // Context
    User     *auth.User
    Tenant   *auth.Tenant
    
    // HTMX
    HXPost, HXTarget, HXSwap string
    HXValidate, HXBoost bool
    
    // Alpine state
    XData string
    
    // Features
    AutoSave bool
    AutoSaveEndpoint string
    SaveDelayMs int
    ValidateOnBlur bool
    ShowUnsavedWarning bool
    
    Layout FormLayout  // stacked, grid, tabs, wizard
}

templ Form(props FormProps) {
    <form
        { buildFormAttributes(props)... }
        x-data={ getFormAlpineData(props) }
    >
        if props.ShowUnsavedWarning {
            <div x-show="hasUnsavedChanges" class="unsaved-warning">
                You have unsaved changes
                @atoms.Button(atoms.ButtonProps{
                    XOn: map[string]string{"click": "saveForm()"},
                }) { Save Now }
            </div>
        }
        
        { renderFormLayout(props) }
        @FormActions(FormActionsProps{...})
    </form>
}

func getFormAlpineData(props FormProps) string {
    return fmt.Sprintf(`{
        loading: false,
        errors: %s,
        originalValues: %s,
        currentValues: %s,
        hasUnsavedChanges: false,
        
        init() {
            this.setupChangeTracking();
            %s
        },
        
        setupChangeTracking() {
            this.$watch('currentValues', () => {
                this.hasUnsavedChanges = JSON.stringify(this.originalValues) !== 
                    JSON.stringify(this.currentValues);
            }, { deep: true });
        },
        
        saveForm() {
            htmx.ajax('POST', '%s', {source: this.$el, swap: 'none'});
        }
    }`,
        toJSON(props.Errors),
        toJSON(props.Values),
        toJSON(props.Values),
        utils.If(props.AutoSave, "this.setupAutoSave();"),
        props.AutoSaveEndpoint,
    )
}
```

### Data Table with Sorting/Filtering

```go
type DataTableProps struct {
    ID       string
    Columns  []TableColumn
    Data     []map[string]any
    
    // Features
    Sortable, Filterable, Selectable bool
    Paginated bool
    PageSize, CurrentPage, TotalItems int
    
    BulkActions []BulkAction
    RowActions  []RowAction
    
    Loading bool
    EmptyTitle, EmptyMessage string
}

templ DataTable(props DataTableProps) {
    <div class="data-table-wrapper">
        if props.Loading {
            @TableSkeleton(TableSkeletonProps{...})
        } else if len(props.Data) == 0 {
            @EmptyState(EmptyStateProps{...})
        } else {
            <table>
                @TableHeader(TableHeaderProps{...})
                <tbody>
                    for _, row := range props.Data {
                        @TableRow(TableRowProps{...})
                    }
                </tbody>
            </table>
        }
        
        if props.Paginated {
            @Pagination(PaginationProps{...})
        }
    </div>
}
```

---

## 4. Templates Layer - Layout Props

```go
type DashboardTemplateProps struct {
    User     *auth.User
    Tenant   *auth.Tenant
    
    Navigation  navigation.Config
    Breadcrumbs []navigation.Breadcrumb
    
    SidebarCollapsed bool
    SidebarPosition  SidebarPosition
    HeaderFixed      bool
    
    ShowSearch, ShowNotifications bool
    Logo, AppName, Theme string
    Title, Description string
}

templ DashboardTemplate(props DashboardTemplateProps) {
    <!DOCTYPE html>
    <html data-theme={ props.Theme }>
        <head>
            @MetaTags(MetaTagsProps{...})
            @ThemeStyles(props.Tenant.ID)
        </head>
        <body>
            <div class={ dashboardLayoutClasses(props) }>
                @organisms.Sidebar(organisms.SidebarProps{...})
                <div class="dashboard-main">
                    @organisms.Header(organisms.HeaderProps{...})
                    @molecules.Breadcrumbs(molecules.BreadcrumbsProps{...})
                    <main>{ children... }</main>
                    @organisms.Footer(organisms.FooterProps{...})
                </div>
            </div>
            @AppScripts()
        </body>
    </html>
}
```

---

## 5. Props Patterns & Utils Integration

### Builder Pattern

```go
type ButtonBuilder struct {
    props ButtonProps
}

func NewButton() *ButtonBuilder {
    return &ButtonBuilder{props: DefaultButtonProps()}
}

func (b *ButtonBuilder) Variant(v ButtonVariant) *ButtonBuilder {
    b.props.Variant = v
    return b
}

func (b *ButtonBuilder) WithHTMX(endpoint, target string) *ButtonBuilder {
    b.props.HXPost = endpoint
    b.props.HXTarget = target
    return b
}

func (b *ButtonBuilder) WithClasses(classes ...string) *ButtonBuilder {
    b.props.ClassName = utils.TwMerge(classes...)
    return b
}

func (b *ButtonBuilder) Build() ButtonProps {
    return b.props
}

// Usage
@atoms.Button(NewButton().
    Variant(atoms.ButtonPrimary).
    WithHTMX("/api/save", "#result").
    WithClasses("ml-4").
    Build())
```

### Utils Integration Benefits

```go
// Smart class merging with conflict resolution
func buttonClasses(props ButtonProps) string {
    return utils.TwMerge(
        "btn",
        fmt.Sprintf("btn-%s btn-%s", props.Variant, props.Size),
        utils.If(props.Active, "active"),
        utils.If(props.FullWidth, "w-full"),
        props.ClassName,  // Conflicts resolved automatically
    )
}

// Clean conditional attributes
func buildInputAttrs(props InputProps) templ.Attributes {
    return utils.MergeAttributes(
        templ.Attributes{
            "type": string(props.Type),
            "id":   utils.IfElse(props.ID != "", props.ID, utils.RandomID()),
            "class": utils.TwMerge("input", props.ClassName),
        },
        templ.Attributes{
            "required": utils.If(props.Required, true),
            "disabled": utils.If(props.Disabled, true),
        },
    )
}

// Cache-busted scripts
<script src={ fmt.Sprintf("/static/js/app.js?v=%s", utils.ScriptVersion) }></script>
```

### Props Inheritance

```go
// Base props for reuse
type BaseComponentProps struct {
    ID, ClassName, DataTestID string
    DataAttributes map[string]string
}

type BaseFormElementProps struct {
    BaseComponentProps
    Name, Value string
    Disabled, Required bool
}

// Specific component embeds base
type SelectProps struct {
    BaseFormElementProps
    Options  []SelectOption
    Multiple bool
}
```

---

## 6. HTMX + Alpine.js Integration Examples

### Real-time Search with Filters

```go
@molecules.SearchBar(molecules.SearchBarProps{
    SearchEndpoint: "/api/products/search",
    HXTarget: "#search-results",
    DebounceMs: 300,
})

<div 
    id="search-results"
    x-data="{ filters: {} }"
    x-init="$watch('filters', () => htmx.trigger('#search input', 'filter-change'))"
>
    // HTMX-loaded results
</div>
```

### Multi-step Wizard

```go
<div x-data="{
    currentStep: 1,
    totalSteps: 4,
    
    nextStep() {
        this.currentStep++;
        htmx.ajax('GET', '/api/step/' + this.currentStep, {
            target: '#step-content'
        });
    }
}">
    @ProgressIndicator()
    <div id="step-content"></div>
    
    @atoms.Button(atoms.ButtonProps{
        XOn: map[string]string{"click": "nextStep()"},
        XShow: "currentStep < totalSteps",
    }) { Next }
</div>
```

### Real-time Notifications

```go
<div x-data="{
    notifications: [],
    unreadCount: 0,
    
    init() {
        const es = new EventSource('/api/notifications/stream');
        es.onmessage = (e) => {
            this.notifications.unshift(JSON.parse(e.data));
            this.unreadCount++;
        };
    }
}">
    @atoms.Button(atoms.ButtonProps{
        XOn: map[string]string{"click": "isOpen = !isOpen"},
    }) {
        @atoms.Icon(atoms.IconProps{Name: "bell"})
        <span x-show="unreadCount > 0" x-text="unreadCount"></span>
    }
    
    <div x-show="isOpen" id="notification-list">
        // HTMX-loaded list
    </div>
</div>
```

---

## Key Benefits

- **Type Safety**: Strongly-typed props with HTMX/Alpine attributes
- **Rich Interactions**: Server-client communication via HTMX
- **Reactive UI**: Client state with Alpine.js
- **Progressive Enhancement**: Works without JS
- **Performance**: Minimal JS footprint, server-side rendering
- **Maintainability**: Clear server/client separation
- **Developer Experience**: Self-documenting typed props
- **Utils Integration**: TwMerge conflicts, conditional logic, smart defaults

## Architecture Flow

```
Design Tokens → Atoms (Button, Input) → 
Molecules (FormField, SearchBar) → 
Organisms (Form, DataTable) → 
Templates (Dashboard, FormPage) → 
Pages (UserProfile, ProductList)
```

Each level uses strongly-typed props with HTMX for server interactions and Alpine.js for client reactivity.
