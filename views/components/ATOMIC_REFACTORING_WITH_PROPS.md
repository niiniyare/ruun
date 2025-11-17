# Atomic Design Refactoring Strategy with Props Pattern

## Executive Summary

This document provides a comprehensive refactoring strategy using strongly-typed props at every level of the atomic design hierarchy, ensuring type safety, maintainability, and clear component contracts.

## Props Architecture Overview

```
DESIGN TOKENS → ATOM PROPS → MOLECULE PROPS → ORGANISM PROPS → TEMPLATE PROPS → PAGE PROPS
```

## 1. Atoms Layer - Foundation Props with HTMX/Alpine

### 1.1 Button Atom with Props

```go
// views/components/atoms/button.templ
package atoms

type ButtonProps struct {
    // Core properties
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
    ClassName   string  // Additional classes
    
    // HTMX attributes
    HXGet       string  // hx-get URL
    HXPost      string  // hx-post URL
    HXPut       string  // hx-put URL
    HXDelete    string  // hx-delete URL
    HXPatch     string  // hx-patch URL
    HXTarget    string  // hx-target selector
    HXSwap      string  // hx-swap strategy
    HXTrigger   string  // hx-trigger events
    HXPushURL   string  // hx-push-url
    HXConfirm   string  // hx-confirm message
    HXIndicator string  // hx-indicator selector
    HXHeaders   map[string]string  // hx-headers
    
    // Alpine.js attributes
    XData       string  // x-data Alpine component
    XShow       string  // x-show condition
    XBind       map[string]string  // x-bind attributes
    XOn         map[string]string  // x-on event handlers
    XText       string  // x-text content
    XHTML       string  // x-html content
    XRef        string  // x-ref reference
    
    // Accessibility
    AriaLabel       string
    AriaPressed     bool
    AriaExpanded    bool
    AriaControls    string
    
    // Data attributes
    DataTestID      string
    DataAttributes  map[string]string
}

type ButtonVariant string
const (
    ButtonPrimary     ButtonVariant = "primary"
    ButtonSecondary   ButtonVariant = "secondary"
    ButtonDestructive ButtonVariant = "destructive"
    ButtonOutline     ButtonVariant = "outline"
    ButtonGhost       ButtonVariant = "ghost"
    ButtonLink        ButtonVariant = "link"
)

type ButtonSize string
const (
    ButtonXS ButtonSize = "xs"
    ButtonSM ButtonSize = "sm"
    ButtonMD ButtonSize = "md"
    ButtonLG ButtonSize = "lg"
    ButtonXL ButtonSize = "xl"
)

type ButtonType string
const (
    ButtonTypeButton ButtonType = "button"
    ButtonTypeSubmit ButtonType = "submit"
    ButtonTypeReset  ButtonType = "reset"
)

templ Button(props ButtonProps) {
    // Generate unique ID if not provided
    buttonID := utils.IfElse(props.ID != "", props.ID, utils.RandomID())
    
    // Build all attributes using utils
    baseAttrs := templ.Attributes{
        "id":   buttonID,
        "type": string(utils.IfElse(props.Type != "", props.Type, ButtonTypeButton)),
        "class": buttonClasses(props),
        "disabled": props.Disabled || props.Loading,
    }
    
    // HTMX attributes
    htmxAttrs := templ.Attributes{
        "hx-get":       props.HXGet,
        "hx-post":      props.HXPost,
        "hx-put":       props.HXPut,
        "hx-delete":    props.HXDelete,
        "hx-patch":     props.HXPatch,
        "hx-target":    props.HXTarget,
        "hx-swap":      props.HXSwap,
        "hx-trigger":   props.HXTrigger,
        "hx-push-url":  props.HXPushURL,
        "hx-confirm":   props.HXConfirm,
        "hx-indicator": props.HXIndicator,
    }
    
    // Alpine.js attributes
    alpineAttrs := templ.Attributes{
        "x-data": props.XData,
        "x-show": props.XShow,
        "x-text": props.XText,
        "x-html": props.XHTML,
        "x-ref":  props.XRef,
    }
    
    // Accessibility attributes
    ariaAttrs := templ.Attributes{
        "aria-label":    props.AriaLabel,
        "aria-pressed":  utils.IfElse(props.AriaPressed, "true", ""),
        "aria-expanded": utils.IfElse(props.AriaExpanded, "true", ""),
        "aria-controls": props.AriaControls,
    }
    
    // Data attributes
    dataAttrs := templ.Attributes{
        "data-testid": props.DataTestID,
    }
    
    // Merge all attributes using utils
    allAttrs := utils.MergeAttributes(
        baseAttrs,
        htmxAttrs,
        alpineAttrs,
        ariaAttrs,
        dataAttrs,
        htmxHeaders(props.HXHeaders),
        alpineBind(props.XBind),
        alpineOn(props.XOn),
        props.DataAttributes,
    )
    
    <button { allAttrs... }>
        if props.Loading {
            @Spinner(SpinnerProps{Size: mapButtonSizeToSpinner(props.Size)})
        }
        { children... }
    </button>
}

func buttonClasses(props ButtonProps) string {
    classes := []string{"btn"}
    classes = append(classes, fmt.Sprintf("btn-%s", props.Variant))
    classes = append(classes, fmt.Sprintf("btn-%s", props.Size))
    
    if props.Active {
        classes = append(classes, "active")
    }
    if props.FullWidth {
        classes = append(classes, "w-full")
    }
    if props.ClassName != "" {
        classes = append(classes, props.ClassName)
    }
    
    return strings.Join(classes, " ")
}

// Helper functions for HTMX and Alpine attributes using existing utils
import "views/components/utils"

func htmxHeaders(headers map[string]string) templ.Attributes {
    attrs := make(templ.Attributes)
    if len(headers) > 0 {
        // Convert map to JSON for hx-headers
        headerJSON, _ := json.Marshal(headers)
        attrs["hx-headers"] = string(headerJSON)
    }
    return attrs
}

func alpineBind(bindings map[string]string) templ.Attributes {
    attrs := make(templ.Attributes)
    for key, value := range bindings {
        attrs[fmt.Sprintf("x-bind:%s", key)] = value
    }
    return attrs
}

func alpineOn(handlers map[string]string) templ.Attributes {
    attrs := make(templ.Attributes)
    for event, handler := range handlers {
        attrs[fmt.Sprintf("x-on:%s", event)] = handler
    }
    return attrs
}

// Enhanced class generation using TwMerge
func buttonClasses(props ButtonProps) string {
    baseClasses := []string{"btn"}
    variantClass := fmt.Sprintf("btn-%s", props.Variant)
    sizeClass := fmt.Sprintf("btn-%s", props.Size)
    
    conditionalClasses := []string{
        utils.If(props.Active, "active"),
        utils.If(props.FullWidth, "w-full"),
        utils.If(props.Loading, "loading"),
        utils.If(props.Disabled, "disabled"),
        props.ClassName, // Additional custom classes
    }
    
    allClasses := append(baseClasses, variantClass, sizeClass)
    allClasses = append(allClasses, conditionalClasses...)
    
    // Use TwMerge to resolve conflicts and clean up classes
    return utils.TwMerge(allClasses...)
}

// Enhanced attribute merging
func mergeComponentAttributes(base templ.Attributes, additional ...templ.Attributes) templ.Attributes {
    all := append([]templ.Attributes{base}, additional...)
    return utils.MergeAttributes(all...)
}
```

### 1.2 Input Atom with Props

```go
// views/components/atoms/input.templ
package atoms

type InputProps struct {
    // Core properties
    ID          string
    Name        string
    Type        InputType
    Value       string
    Placeholder string
    
    // Validation
    Required    bool
    Pattern     string
    MinLength   int
    MaxLength   int
    Min         string  // for number/date inputs
    Max         string  // for number/date inputs
    
    // States
    Disabled    bool
    ReadOnly    bool
    Invalid     bool
    
    // Styling
    Size        InputSize
    ClassName   string
    
    // HTMX attributes for inputs
    HXPost          string  // hx-post for form submission
    HXTrigger       string  // hx-trigger (e.g., "keyup changed delay:300ms")
    HXTarget        string  // hx-target for validation results
    HXSwap          string  // hx-swap strategy
    HXValidate      bool    // hx-validate for instant validation
    HXIndicator     string  // hx-indicator selector
    
    // Alpine.js attributes
    XData       string  // x-data Alpine component
    XModel      string  // x-model for two-way binding
    XShow       string  // x-show condition
    XBind       map[string]string  // x-bind attributes
    XOn         map[string]string  // x-on event handlers
    XRef        string  // x-ref reference
    
    // Accessibility
    AriaLabel           string
    AriaDescribedBy     string
    AriaInvalid         bool
    
    // Behavior
    AutoComplete        string
    AutoFocus          bool
    
    // Data attributes
    DataTestID         string
    DataAttributes     map[string]string
}

type InputType string
const (
    InputText     InputType = "text"
    InputEmail    InputType = "email"
    InputPassword InputType = "password"
    InputNumber   InputType = "number"
    InputTel      InputType = "tel"
    InputURL      InputType = "url"
    InputSearch   InputType = "search"
    InputDate     InputType = "date"
    InputTime     InputType = "time"
)

type InputSize string
const (
    InputSM InputSize = "sm"
    InputMD InputSize = "md"
    InputLG InputSize = "lg"
)

templ Input(props InputProps) {
    <input
        id={ props.ID }
        name={ props.Name }
        type={ string(props.Type) }
        value={ props.Value }
        placeholder={ props.Placeholder }
        class={ inputClasses(props) }
        required={ props.Required }
        pattern={ props.Pattern }
        minlength={ fmt.Sprintf("%d", props.MinLength) }
        maxlength={ fmt.Sprintf("%d", props.MaxLength) }
        min={ props.Min }
        max={ props.Max }
        disabled={ props.Disabled }
        readonly={ props.ReadOnly }
        
        // HTMX attributes
        hx-post={ props.HXPost }
        hx-trigger={ props.HXTrigger }
        hx-target={ props.HXTarget }
        hx-swap={ props.HXSwap }
        hx-validate={ fmt.Sprintf("%t", props.HXValidate) }
        hx-indicator={ props.HXIndicator }
        
        // Alpine.js attributes
        x-data={ props.XData }
        x-model={ props.XModel }
        x-show={ props.XShow }
        x-ref={ props.XRef }
        { alpineBind(props.XBind)... }
        { alpineOn(props.XOn)... }
        
        // Accessibility
        aria-label={ props.AriaLabel }
        aria-describedby={ props.AriaDescribedBy }
        aria-invalid={ fmt.Sprintf("%t", props.AriaInvalid || props.Invalid) }
        
        // Behavior
        autocomplete={ props.AutoComplete }
        autofocus={ props.AutoFocus }
        
        // Data attributes
        data-testid={ props.DataTestID }
        { dataAttributes(props.DataAttributes)... }
    />
}
```

### 1.3 Other Essential Atoms with Props

```go
// views/components/atoms/label.templ
type LabelProps struct {
    For         string
    Required    bool
    Size        LabelSize
    Weight      LabelWeight
    ClassName   string
    DataTestID  string
}

type LabelSize string
const (
    LabelXS LabelSize = "xs"
    LabelSM LabelSize = "sm"
    LabelMD LabelSize = "md"
    LabelLG LabelSize = "lg"
)

templ Label(props LabelProps) {
    <label
        for={ props.For }
        class={ labelClasses(props) }
        data-testid={ props.DataTestID }
    >
        { children... }
        if props.Required {
            <span class="text-destructive ml-1">*</span>
        }
    </label>
}

// views/components/atoms/icon.templ
type IconProps struct {
    Name        string
    Size        IconSize
    Color       string
    ClassName   string
    AriaHidden  bool
    DataTestID  string
}

type IconSize string
const (
    IconXS IconSize = "xs"  // 12px
    IconSM IconSize = "sm"  // 16px
    IconMD IconSize = "md"  // 20px
    IconLG IconSize = "lg"  // 24px
    IconXL IconSize = "xl"  // 32px
)

templ Icon(props IconProps) {
    <i
        class={ iconClasses(props) }
        aria-hidden={ fmt.Sprintf("%t", props.AriaHidden) }
        data-testid={ props.DataTestID }
    />
}

// views/components/atoms/spinner.templ
type SpinnerProps struct {
    Size        SpinnerSize
    Color       string
    ClassName   string
    AriaLabel   string
}

type SpinnerSize string
const (
    SpinnerXS SpinnerSize = "xs"
    SpinnerSM SpinnerSize = "sm"
    SpinnerMD SpinnerSize = "md"
    SpinnerLG SpinnerSize = "lg"
)

templ Spinner(props SpinnerProps) {
    <div
        class={ spinnerClasses(props) }
        role="status"
        aria-label={ getAriaLabel(props.AriaLabel, "Loading...") }
    >
        <svg class="animate-spin" viewBox="0 0 24 24">
            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4" />
            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
        </svg>
    </div>
}
```

## 2. Molecules Layer - Composition Props

### 2.1 Form Field Molecule with Props

```go
// views/components/molecules/form_field.templ
package molecules

import "views/components/atoms"

type FormFieldProps struct {
    // Field identity
    ID          string
    Name        string
    
    // Label configuration
    Label       string
    LabelSize   atoms.LabelSize
    Required    bool
    
    // Input configuration
    Type        atoms.InputType
    Value       string
    Placeholder string
    InputSize   atoms.InputSize
    
    // Validation
    Pattern     string
    MinLength   int
    MaxLength   int
    
    // States
    Disabled    bool
    ReadOnly    bool
    Invalid     bool
    
    // Messages
    ErrorMessage    string
    HelpText        string
    
    // Layout
    Layout      FormFieldLayout  // vertical, horizontal, floating
    
    // Styling
    ClassName   string
    
    // Data attributes
    DataTestID  string
}

type FormFieldLayout string
const (
    FieldLayoutVertical   FormFieldLayout = "vertical"
    FieldLayoutHorizontal FormFieldLayout = "horizontal"
    FieldLayoutFloating   FormFieldLayout = "floating"
)

templ FormField(props FormFieldProps) {
    <div class={ formFieldClasses(props) } data-testid={ props.DataTestID }>
        if props.Layout != FieldLayoutFloating {
            @atoms.Label(atoms.LabelProps{
                For:      props.ID,
                Required: props.Required,
                Size:     props.LabelSize,
            }) {
                { props.Label }
            }
        }
        
        <div class="form-field-input-wrapper">
            @atoms.Input(atoms.InputProps{
                ID:              props.ID,
                Name:            props.Name,
                Type:            props.Type,
                Value:           props.Value,
                Placeholder:     props.Placeholder,
                Size:            props.InputSize,
                Required:        props.Required,
                Pattern:         props.Pattern,
                MinLength:       props.MinLength,
                MaxLength:       props.MaxLength,
                Disabled:        props.Disabled,
                ReadOnly:        props.ReadOnly,
                Invalid:         props.Invalid,
                AriaDescribedBy: getAriaDescribedBy(props),
                AriaInvalid:     props.Invalid || props.ErrorMessage != "",
            })
            
            if props.Layout == FieldLayoutFloating {
                @atoms.Label(atoms.LabelProps{
                    For:      props.ID,
                    Required: props.Required,
                    Size:     props.LabelSize,
                    ClassName: "floating-label",
                }) {
                    { props.Label }
                }
            }
        </div>
        
        if props.ErrorMessage != "" {
            @ErrorMessage(ErrorMessageProps{
                ID:      props.ID + "-error",
                Message: props.ErrorMessage,
            })
        }
        
        if props.HelpText != "" {
            @HelpText(HelpTextProps{
                ID:      props.ID + "-help",
                Text:    props.HelpText,
            })
        }
    </div>
}
```

### 2.2 Button Group Molecule with Props

```go
// views/components/molecules/button_group.templ
package molecules

type ButtonGroupProps struct {
    // Layout
    Orientation ButtonGroupOrientation  // horizontal, vertical
    Alignment   ButtonGroupAlignment    // start, center, end, between, around
    
    // Behavior
    Attached    bool  // Buttons visually connected
    
    // Styling
    Size        atoms.ButtonSize  // Applies to all buttons
    FullWidth   bool
    ClassName   string
    
    // Data attributes
    DataTestID  string
}

type ButtonGroupOrientation string
const (
    GroupHorizontal ButtonGroupOrientation = "horizontal"
    GroupVertical   ButtonGroupOrientation = "vertical"
)

type ButtonGroupAlignment string
const (
    GroupStart   ButtonGroupAlignment = "start"
    GroupCenter  ButtonGroupAlignment = "center"
    GroupEnd     ButtonGroupAlignment = "end"
    GroupBetween ButtonGroupAlignment = "between"
    GroupAround  ButtonGroupAlignment = "around"
)

templ ButtonGroup(props ButtonGroupProps) {
    <div 
        class={ buttonGroupClasses(props) }
        role="group"
        data-testid={ props.DataTestID }
    >
        { children... }
    </div>
}

// Convenience component for common use case
templ ActionButtons(props ActionButtonsProps) {
    @ButtonGroup(ButtonGroupProps{
        Orientation: GroupHorizontal,
        Alignment:   GroupEnd,
        ClassName:   props.ClassName,
    }) {
        if props.ShowCancel {
            @atoms.Button(atoms.ButtonProps{
                Variant: atoms.ButtonOutline,
                Size:    props.Size,
            }) {
                { props.CancelText }
            }
        }
        
        @atoms.Button(atoms.ButtonProps{
            Variant:  atoms.ButtonPrimary,
            Size:     props.Size,
            Type:     atoms.ButtonTypeSubmit,
            Loading:  props.SubmitLoading,
            Disabled: props.SubmitDisabled,
        }) {
            { props.SubmitText }
        }
    }
}
```

### 2.3 Search Bar Molecule with Props

```go
// views/components/molecules/search_bar.templ
package molecules

type SearchBarProps struct {
    // Core functionality
    ID              string
    Name            string
    Value           string
    Placeholder     string
    
    // Search configuration
    SearchOnType    bool
    DebounceMs      int
    MinChars        int
    
    // Button configuration
    ShowButton      bool
    ButtonText      string
    ButtonVariant   atoms.ButtonVariant
    
    // States
    Loading         bool
    Disabled        bool
    
    // Styling
    Size            ComponentSize
    ClassName       string
    
    // Events
    OnSearch        string  // JavaScript function
    Action          string  // Form action
    Method          string  // Form method
    
    // Data attributes
    DataTestID      string
}

type SearchBarProps struct {
    // Core functionality
    ID              string
    Name            string
    Value           string
    Placeholder     string
    
    // Search configuration
    SearchEndpoint  string  // HTMX search endpoint
    MinChars        int
    DebounceMs      int
    
    // Button configuration
    ShowButton      bool
    ButtonText      string
    ButtonVariant   atoms.ButtonVariant
    
    // States
    Loading         bool
    Disabled        bool
    
    // HTMX configuration
    HXTarget        string  // Where to put search results
    HXSwap          string  // How to swap results
    HXIndicator     string  // Loading indicator
    
    // Alpine.js integration
    XData           string  // Alpine data for search state
    
    // Styling
    Size            ComponentSize
    ClassName       string
    
    // Form configuration (for non-HTMX fallback)
    Action          string
    Method          string
    
    // Data attributes
    DataTestID      string
}

templ SearchBar(props SearchBarProps) {
    <form 
        class={ searchBarClasses(props) }
        action={ props.Action }
        method={ getMethod(props.Method, "GET") }
        data-testid={ props.DataTestID }
        x-data={ getSearchAlpineData(props) }
    >
        <div class="search-input-wrapper">
            @atoms.Icon(atoms.IconProps{
                Name: "search",
                Size: mapSizeToIcon(props.Size),
                ClassName: "search-icon",
            })
            
            @atoms.Input(atoms.InputProps{
                ID:          props.ID,
                Name:        props.Name,
                Type:        atoms.InputSearch,
                Value:       props.Value,
                Placeholder: props.Placeholder,
                Size:        mapSizeToInput(props.Size),
                Disabled:    props.Disabled,
                
                // HTMX live search
                HXPost:      props.SearchEndpoint,
                HXTrigger:   fmt.Sprintf("keyup changed delay:%dms", props.DebounceMs),
                HXTarget:    props.HXTarget,
                HXSwap:      props.HXSwap,
                HXIndicator: props.HXIndicator,
                
                // Alpine integration for client-side state
                XModel:      "query",
                XOn: map[string]string{
                    "input": "handleInput($event)",
                },
            })
            
            if props.Loading {
                @atoms.Spinner(atoms.SpinnerProps{
                    Size: mapSizeToSpinner(props.Size),
                    ClassName: "search-spinner",
                    XShow: "loading",
                })
            }
        </div>
        
        if props.ShowButton {
            @atoms.Button(atoms.ButtonProps{
                Type:     atoms.ButtonTypeSubmit,
                Variant:  props.ButtonVariant,
                Size:     mapSizeToButton(props.Size),
                Disabled: props.Disabled,
                XShow:    "query.length >= " + fmt.Sprintf("%d", props.MinChars),
            }) {
                { getButtonText(props.ButtonText, "Search") }
            }
        }
        
        // Clear button for Alpine state management
        @atoms.Button(atoms.ButtonProps{
            Type:     atoms.ButtonTypeButton,
            Variant:  atoms.ButtonGhost,
            Size:     atoms.ButtonSM,
            XShow:    "query.length > 0",
            XOn:      map[string]string{"click": "clearSearch()"},
            ClassName: "search-clear",
        }) {
            @atoms.Icon(atoms.IconProps{Name: "x", Size: "sm"})
        }
    </form>
}

// Generate Alpine.js data for search functionality
func getSearchAlpineData(props SearchBarProps) string {
    return fmt.Sprintf(`{
        query: '%s',
        loading: false,
        results: [],
        minChars: %d,
        handleInput(event) {
            if (event.target.value.length >= this.minChars) {
                this.loading = true;
            }
        },
        clearSearch() {
            this.query = '';
            this.results = [];
            this.loading = false;
        }
    }`, props.Value, props.MinChars)
}
```

## 3. Organisms Layer - Complex Props with Business Logic

### 3.1 Form Organism with Props

```go
// views/components/organisms/form.templ
package organisms

import (
    "views/components/molecules"
    "github.com/ruun/pkg/schema"
    "github.com/ruun/pkg/auth"
)

type FormProps struct {
    // Identity
    ID          string
    Name        string
    
    // Schema and context
    Schema      schema.FormSchema
    Values      map[string]interface{}
    Errors      map[string][]string
    
    // User context for permissions
    User        *auth.User
    Tenant      *auth.Tenant
    
    // Form configuration
    Action      string
    Method      string
    EncType     string
    
    // HTMX form handling
    HXPost          string  // Form submission endpoint
    HXTarget        string  // Where to put response
    HXSwap          string  // How to swap response
    HXPushURL       string  // URL to push to history
    HXValidate      bool    // Enable HTMX validation
    HXConfirm       string  // Confirmation message
    HXBoost         bool    // Enable HTMX boost
    
    // Alpine.js state management
    XData           string  // Alpine data for form state
    XValidate       bool    // Client-side validation
    
    // Auto-save with HTMX
    AutoSave        bool
    AutoSaveEndpoint string  // HTMX endpoint for auto-save
    SaveDelayMs     int
    
    // Real-time validation
    ValidateOnBlur  bool
    ValidationEndpoint string  // HTMX validation endpoint
    
    // Layout
    Layout      FormLayout
    Columns     int
    
    // States
    Loading     bool
    Disabled    bool
    ReadOnly    bool
    
    // Styling
    ClassName   string
    
    // Features
    ShowProgress    bool
    ShowSummary     bool
    ShowUnsavedWarning bool  // Alpine-powered unsaved changes warning
    
    // Data attributes
    DataTestID  string
}

type FormLayout string
const (
    FormLayoutStacked FormLayout = "stacked"
    FormLayoutGrid    FormLayout = "grid"
    FormLayoutTabs    FormLayout = "tabs"
    FormLayoutWizard  FormLayout = "wizard"
)

templ Form(props FormProps) {
    // Extract visible sections based on permissions
    visibleSections := getVisibleSections(props.Schema.Sections, props.User, props.Tenant)
    
    if len(visibleSections) == 0 {
        @EmptyState(EmptyStateProps{
            Title: "No Access",
            Description: "You don't have permission to view this form.",
            Icon: "lock",
        })
        return
    }
    
    <form
        id={ props.ID }
        name={ props.Name }
        class={ formClasses(props) }
        action={ props.Action }
        method={ props.Method }
        enctype={ props.EncType }
        novalidate={ props.ValidateOnBlur }
        
        // HTMX attributes for form handling
        hx-post={ props.HXPost }
        hx-target={ props.HXTarget }
        hx-swap={ props.HXSwap }
        hx-push-url={ props.HXPushURL }
        hx-validate={ fmt.Sprintf("%t", props.HXValidate) }
        hx-confirm={ props.HXConfirm }
        hx-boost={ fmt.Sprintf("%t", props.HXBoost) }
        
        // Alpine.js for rich client-side interactions
        x-data={ getFormAlpineData(props) }
        x-on:beforeunload.window={ getBeforeUnloadHandler(props.ShowUnsavedWarning) }
        
        data-testid={ props.DataTestID }
        { formDataAttributes(props)... }
    >
        if props.ShowProgress && props.Layout == FormLayoutWizard {
            @FormProgress(FormProgressProps{
                Steps: len(visibleSections),
                Current: getCurrentStep(props.Values),
                XBind: map[string]string{
                    "class": "getProgressClasses(currentStep, totalSteps)",
                },
            })
        }
        
        if props.ShowSummary && hasErrors(props.Errors) {
            @FormSummary(FormSummaryProps{
                Errors: props.Errors,
                XShow: "showSummary && Object.keys(errors).length > 0",
            })
        }
        
        // Unsaved changes warning
        if props.ShowUnsavedWarning {
            <div 
                x-show="hasUnsavedChanges" 
                x-transition
                class="unsaved-warning bg-warning-50 border border-warning-200 p-4 rounded-md mb-4"
            >
                <div class="flex">
                    @atoms.Icon(atoms.IconProps{Name: "alert-triangle", Size: "sm", ClassName: "text-warning-600"})
                    <span class="text-warning-800 ml-2">You have unsaved changes</span>
                    @atoms.Button(atoms.ButtonProps{
                        Variant: atoms.ButtonLink,
                        Size: atoms.ButtonSM,
                        XOn: map[string]string{"click": "saveForm()"},
                        ClassName: "ml-auto",
                    }) {
                        Save Now
                    }
                </div>
            </div>
        }
        
        switch props.Layout {
        case FormLayoutTabs:
            @TabbedForm(visibleSections, props)
        case FormLayoutWizard:
            @WizardForm(visibleSections, props)
        case FormLayoutGrid:
            @GridForm(visibleSections, props)
        default:
            @StackedForm(visibleSections, props)
        }
        
        @FormActions(FormActionsProps{
            Schema:   props.Schema,
            User:     props.User,
            Disabled: props.Disabled || props.Loading,
            Loading:  props.Loading,
            XBind: map[string]string{
                "disabled": "loading || disabled",
            },
        })
    </form>
}

// Generate Alpine.js data for comprehensive form state management
func getFormAlpineData(props FormProps) string {
    autoSaveConfig := "null"
    if props.AutoSave {
        autoSaveConfig = fmt.Sprintf(`{
            enabled: true,
            endpoint: '%s',
            delay: %d,
            timeout: null
        }`, props.AutoSaveEndpoint, props.SaveDelayMs)
    }
    
    return fmt.Sprintf(`{
        // Form state
        loading: false,
        disabled: %t,
        errors: %s,
        originalValues: %s,
        currentValues: %s,
        hasUnsavedChanges: false,
        showSummary: %t,
        
        // Auto-save configuration
        autoSave: %s,
        
        // Form lifecycle methods
        init() {
            this.originalValues = { ...this.currentValues };
            if (this.autoSave && this.autoSave.enabled) {
                this.setupAutoSave();
            }
            this.setupChangeTracking();
        },
        
        // Auto-save functionality
        setupAutoSave() {
            this.$watch('currentValues', () => {
                if (this.hasUnsavedChanges) {
                    clearTimeout(this.autoSave.timeout);
                    this.autoSave.timeout = setTimeout(() => {
                        this.saveForm();
                    }, this.autoSave.delay);
                }
            });
        },
        
        // Change tracking
        setupChangeTracking() {
            this.$watch('currentValues', () => {
                this.hasUnsavedChanges = JSON.stringify(this.originalValues) !== JSON.stringify(this.currentValues);
            }, { deep: true });
        },
        
        // Save form via HTMX or fallback
        saveForm() {
            if (this.autoSave && this.autoSave.enabled) {
                htmx.ajax('POST', this.autoSave.endpoint, {
                    source: this.$el,
                    swap: 'none'
                });
            }
        },
        
        // Reset form state
        resetForm() {
            this.currentValues = { ...this.originalValues };
            this.hasUnsavedChanges = false;
            this.errors = {};
        },
        
        // Field validation
        validateField(fieldName) {
            if ('%s' !== '') {
                const field = this.$el.querySelector('[name="' + fieldName + '"]');
                htmx.ajax('POST', '%s', {
                    source: field,
                    target: '#' + fieldName + '-validation'
                });
            }
        },
        
        // Before unload warning
        beforeUnload(event) {
            if (this.hasUnsavedChanges && %t) {
                event.preventDefault();
                return 'You have unsaved changes. Are you sure you want to leave?';
            }
        }
    }`,
        props.Disabled,
        toJSONString(props.Errors),
        toJSONString(props.Values),
        toJSONString(props.Values),
        props.ShowSummary,
        autoSaveConfig,
        props.ValidationEndpoint,
        props.ValidationEndpoint,
        props.ShowUnsavedWarning,
    )
}

func getBeforeUnloadHandler(enabled bool) string {
    if enabled {
        return "beforeUnload($event)"
    }
    return ""
}

// Helper to render form sections
templ FormSection(section schema.Section, props FormProps) {
    visibleFields := getVisibleFields(section.Fields, props.User, props.Tenant)
    
    if len(visibleFields) == 0 {
        return
    }
    
    <fieldset class="form-section" disabled={ props.Disabled || props.Loading }>
        if section.Title != "" {
            <legend class="form-section-title">{ section.Title }</legend>
        }
        
        if section.Description != "" {
            <p class="form-section-description">{ section.Description }</p>
        }
        
        <div class={ sectionLayoutClasses(section, props.Layout, props.Columns) }>
            for _, field := range visibleFields {
                @renderField(field, props)
            }
        </div>
    </fieldset>
}

// Render individual field based on permissions and state
templ renderField(field schema.Field, props FormProps) {
    // Determine field state based on permissions
    fieldState := getFieldState(field, props.User, props.Tenant)
    
    switch fieldState {
    case FieldHidden:
        return
    case FieldReadOnly:
        @molecules.ReadOnlyField(molecules.ReadOnlyFieldProps{
            Label: field.Label,
            Value: formatFieldValue(field, props.Values[field.Name]),
        })
    case FieldDisabled:
        @molecules.FormField(toFormFieldProps(field, props, true))
    default:
        @molecules.FormField(toFormFieldProps(field, props, false))
    }
}

// Business logic functions
func getVisibleSections(sections []schema.Section, user *auth.User, tenant *auth.Tenant) []schema.Section {
    visible := []schema.Section{}
    for _, section := range sections {
        if section.CheckPermissions(user) && section.CheckTenant(tenant) {
            visible = append(visible, section)
        }
    }
    return visible
}

func getFieldState(field schema.Field, user *auth.User, tenant *auth.Tenant) FieldState {
    if !field.CheckPermissions(user) {
        return FieldHidden
    }
    if !field.CanEdit(user) || field.ReadOnly {
        return FieldReadOnly  
    }
    if field.DisabledForTenant(tenant) {
        return FieldDisabled
    }
    return FieldEditable
}
```

### 3.2 Data Table Organism with Props

```go
// views/components/organisms/data_table.templ
package organisms

type DataTableProps struct {
    // Identity
    ID          string
    
    // Data configuration
    Columns     []TableColumn
    Data        []map[string]interface{}
    
    // Features
    Sortable    bool
    Filterable  bool
    Selectable  bool
    Expandable  bool
    
    // Pagination
    Paginated   bool
    PageSize    int
    CurrentPage int
    TotalItems  int
    
    // Selection
    SelectedRows    []string
    OnSelectionChange string
    
    // Actions
    BulkActions []BulkAction
    RowActions  []RowAction
    
    // States
    Loading     bool
    
    // Styling
    Striped     bool
    Bordered    bool
    Compact     bool
    ClassName   string
    
    // Empty state
    EmptyTitle      string
    EmptyMessage    string
    
    // Data attributes
    DataTestID  string
}

type TableColumn struct {
    ID          string
    Header      string
    Field       string
    
    // Sorting
    Sortable    bool
    SortOrder   SortOrder
    
    // Formatting
    Format      ColumnFormat
    Formatter   func(interface{}) string
    
    // Alignment
    Align       ColumnAlign
    
    // Width
    Width       string
    MinWidth    string
    
    // Visibility
    Hidden      bool
    Responsive  ResponsiveVisibility
}

type ColumnFormat string
const (
    FormatText     ColumnFormat = "text"
    FormatNumber   ColumnFormat = "number"
    FormatCurrency ColumnFormat = "currency"
    FormatDate     ColumnFormat = "date"
    FormatBoolean  ColumnFormat = "boolean"
    FormatCustom   ColumnFormat = "custom"
)

templ DataTable(props DataTableProps) {
    <div class={ dataTableWrapperClasses(props) } data-testid={ props.DataTestID }>
        if props.Filterable {
            @TableFilters(TableFiltersProps{
                Columns: getFilterableColumns(props.Columns),
                OnFilter: "handleTableFilter",
            })
        }
        
        if len(props.BulkActions) > 0 && hasSelectedRows(props.SelectedRows) {
            @BulkActionsBar(BulkActionsBarProps{
                Actions: props.BulkActions,
                SelectedCount: len(props.SelectedRows),
            })
        }
        
        <div class="table-container">
            if props.Loading {
                @TableSkeleton(TableSkeletonProps{
                    Columns: len(props.Columns),
                    Rows: props.PageSize,
                })
            } else if len(props.Data) == 0 {
                @EmptyState(EmptyStateProps{
                    Title: props.EmptyTitle,
                    Description: props.EmptyMessage,
                    Icon: "inbox",
                })
            } else {
                <table class={ tableClasses(props) }>
                    @TableHeader(TableHeaderProps{
                        Columns: props.Columns,
                        Sortable: props.Sortable,
                        Selectable: props.Selectable,
                        AllSelected: areAllSelected(props.Data, props.SelectedRows),
                    })
                    
                    <tbody>
                        for _, row := range props.Data {
                            @TableRow(TableRowProps{
                                Data: row,
                                Columns: props.Columns,
                                Selectable: props.Selectable,
                                Selected: isRowSelected(row["id"], props.SelectedRows),
                                Expandable: props.Expandable,
                                Actions: props.RowActions,
                            })
                        }
                    </tbody>
                </table>
            }
        </div>
        
        if props.Paginated && props.TotalItems > props.PageSize {
            @molecules.Pagination(molecules.PaginationProps{
                CurrentPage: props.CurrentPage,
                TotalPages: calculateTotalPages(props.TotalItems, props.PageSize),
                PageSize: props.PageSize,
                TotalItems: props.TotalItems,
                OnPageChange: "handlePageChange",
            })
        }
    </div>
}
```

## 4. Templates Layer - Layout Props

### 4.1 Dashboard Template with Props

```go
// views/components/templates/dashboard.templ
package templates

import (
    "views/components/organisms"
    "github.com/ruun/pkg/auth"
    "github.com/ruun/pkg/navigation"
)

type DashboardTemplateProps struct {
    // User context
    User        *auth.User
    Tenant      *auth.Tenant
    
    // Navigation
    Navigation  navigation.Config
    Breadcrumbs []navigation.Breadcrumb
    
    // Layout configuration
    SidebarCollapsed    bool
    SidebarPosition     SidebarPosition  // left, right
    HeaderFixed         bool
    
    // Features
    ShowSearch          bool
    ShowNotifications   bool
    ShowUserMenu       bool
    
    // Branding
    Logo               string
    LogoUrl            string
    AppName            string
    
    // Theme
    Theme              string
    
    // Meta
    Title              string
    Description        string
    
    // Data attributes
    DataTestID         string
}

type SidebarPosition string
const (
    SidebarLeft  SidebarPosition = "left"
    SidebarRight SidebarPosition = "right"
)

templ DashboardTemplate(props DashboardTemplateProps) {
    <!DOCTYPE html>
    <html lang="en" data-theme={ props.Theme }>
        <head>
            @MetaTags(MetaTagsProps{
                Title: props.Title,
                Description: props.Description,
                AppName: props.AppName,
            })
            @ThemeStyles(props.Tenant.ID)
            @AppStyles()
        </head>
        
        <body class="dashboard-body" data-testid={ props.DataTestID }>
            <div class={ dashboardLayoutClasses(props) }>
                @organisms.Sidebar(organisms.SidebarProps{
                    Navigation: props.Navigation,
                    User: props.User,
                    Collapsed: props.SidebarCollapsed,
                    Position: string(props.SidebarPosition),
                    Logo: props.Logo,
                    LogoUrl: props.LogoUrl,
                })
                
                <div class="dashboard-main">
                    @organisms.Header(organisms.HeaderProps{
                        User: props.User,
                        Fixed: props.HeaderFixed,
                        ShowSearch: props.ShowSearch,
                        ShowNotifications: props.ShowNotifications,
                        ShowUserMenu: props.ShowUserMenu,
                        AppName: props.AppName,
                    })
                    
                    if len(props.Breadcrumbs) > 0 {
                        @molecules.Breadcrumbs(molecules.BreadcrumbsProps{
                            Items: props.Breadcrumbs,
                        })
                    }
                    
                    <main class="dashboard-content">
                        { children... }
                    </main>
                    
                    @organisms.Footer(organisms.FooterProps{
                        AppName: props.AppName,
                        Year: getCurrentYear(),
                    })
                </div>
            </div>
            
            @AppScripts()
        </body>
    </html>
}
```

### 4.2 Form Page Template with Props

```go
// views/components/templates/form_page.templ
package templates

type FormPageTemplateProps struct {
    // Page configuration
    Title           string
    Description     string
    Icon           string
    
    // Layout
    Width          FormPageWidth  // narrow, medium, wide, full
    Centered       bool
    
    // Actions
    BackUrl        string
    BackLabel      string
    
    // Features
    ShowProgress   bool
    ShowHelp       bool
    HelpContent    string
    
    // States
    Loading        bool
    
    // Styling
    ClassName      string
    
    // Data attributes
    DataTestID     string
}

type FormPageWidth string
const (
    FormPageNarrow FormPageWidth = "narrow"   // max-w-md
    FormPageMedium FormPageWidth = "medium"   // max-w-2xl
    FormPageWide   FormPageWidth = "wide"     // max-w-4xl
    FormPageFull   FormPageWidth = "full"     // max-w-full
)

templ FormPageTemplate(props FormPageTemplateProps) {
    <div class={ formPageClasses(props) } data-testid={ props.DataTestID }>
        @molecules.PageHeader(molecules.PageHeaderProps{
            Title: props.Title,
            Description: props.Description,
            Icon: props.Icon,
            BackUrl: props.BackUrl,
            BackLabel: props.BackLabel,
        })
        
        <div class={ formPageContentClasses(props) }>
            if props.ShowHelp && props.HelpContent != "" {
                @molecules.HelpCard(molecules.HelpCardProps{
                    Content: props.HelpContent,
                })
            }
            
            <div class="form-page-main">
                if props.Loading {
                    @organisms.FormSkeleton()
                } else {
                    { children... }
                }
            </div>
        </div>
    </div>
}
```

## 5. Pages Layer - Complete Page Props

### 5.1 User Profile Page with Props

```go
// views/components/pages/user_profile.templ
package pages

type UserProfilePageProps struct {
    // Data
    User            *auth.User
    Profile         *models.UserProfile
    
    // Permissions
    CanEdit         bool
    CanDelete       bool
    CanChangePassword bool
    
    // Features
    ShowActivity    bool
    ShowStats       bool
    
    // API endpoints
    UpdateEndpoint  string
    DeleteEndpoint  string
    
    // Navigation
    ReturnUrl       string
}

templ UserProfilePage(props UserProfilePageProps) {
    @templates.DashboardTemplate(templates.DashboardTemplateProps{
        User: props.User,
        Title: fmt.Sprintf("%s's Profile", props.Profile.Name),
        Navigation: getUserNavigation(props.User),
    }) {
        @templates.FormPageTemplate(templates.FormPageTemplateProps{
            Title: "User Profile",
            Description: "Manage your account settings and preferences",
            Icon: "user",
            Width: templates.FormPageMedium,
            BackUrl: props.ReturnUrl,
        }) {
            @organisms.TabPanel(organisms.TabPanelProps{
                Tabs: []organisms.Tab{
                    {ID: "profile", Label: "Profile", Icon: "user"},
                    {ID: "security", Label: "Security", Icon: "lock"},
                    {ID: "preferences", Label: "Preferences", Icon: "settings"},
                    if props.ShowActivity {
                        {ID: "activity", Label: "Activity", Icon: "activity"}
                    },
                },
                DefaultTab: "profile",
            }) {
                <div id="profile" class="tab-content">
                    @organisms.Form(organisms.FormProps{
                        ID: "profile-form",
                        Schema: getProfileSchema(),
                        Values: profileToValues(props.Profile),
                        User: props.User,
                        Action: props.UpdateEndpoint,
                        Method: "POST",
                        ReadOnly: !props.CanEdit,
                    })
                </div>
                
                <div id="security" class="tab-content">
                    if props.CanChangePassword {
                        @organisms.PasswordChangeForm(organisms.PasswordChangeFormProps{
                            User: props.User,
                            Endpoint: "/api/users/change-password",
                        })
                    }
                </div>
                
                <div id="preferences" class="tab-content">
                    @organisms.PreferencesForm(organisms.PreferencesFormProps{
                        Preferences: props.Profile.Preferences,
                        Endpoint: "/api/users/preferences",
                    })
                </div>
                
                if props.ShowActivity {
                    <div id="activity" class="tab-content">
                        @organisms.ActivityLog(organisms.ActivityLogProps{
                            UserID: props.User.ID,
                            Limit: 50,
                        })
                    </div>
                }
            }
        }
    }
}
```

## 6. Enhanced Props Pattern with Utils Integration

### 6.1 Props Design Guidelines with Utils

```go
import "views/components/utils"

// 1. Use specific types for variants
type ButtonVariant string  // Not just string

// 2. Group related props
type ButtonStateProps struct {
    Disabled bool
    Loading  bool
    Active   bool
}

// 3. Provide defaults through functions using utils
func DefaultButtonProps() ButtonProps {
    return ButtonProps{
        ID:      utils.RandomID(), // Auto-generate ID
        Variant: ButtonPrimary,
        Size:    ButtonMD,
        Type:    ButtonTypeButton,
    }
}

// 4. Use builder pattern for complex props with utils
type ButtonBuilder struct {
    props ButtonProps
}

func NewButton() *ButtonBuilder {
    return &ButtonBuilder{
        props: DefaultButtonProps(),
    }
}

func (b *ButtonBuilder) Variant(v ButtonVariant) *ButtonBuilder {
    b.props.Variant = v
    return b
}

func (b *ButtonBuilder) WithHTMX(endpoint string, target string) *ButtonBuilder {
    b.props.HXPost = endpoint
    b.props.HXTarget = target
    return b
}

func (b *ButtonBuilder) WithClasses(classes ...string) *ButtonBuilder {
    b.props.ClassName = utils.TwMerge(classes...)
    return b
}

func (b *ButtonBuilder) Conditional(condition bool, fn func(*ButtonBuilder)) *ButtonBuilder {
    if condition {
        fn(b)
    }
    return b
}

func (b *ButtonBuilder) Build() ButtonProps {
    return b.props
}

// 5. Enhanced validation with utils
func (p ButtonProps) Validate() error {
    if p.AriaLabel == "" && len(p.children) == 0 {
        return errors.New("button must have either aria-label or children")
    }
    
    // Validate HTMX configuration
    htmxMethods := []string{p.HXGet, p.HXPost, p.HXPut, p.HXDelete, p.HXPatch}
    hasHTMX := false
    for _, method := range htmxMethods {
        if method != "" {
            hasHTMX = true
            break
        }
    }
    
    if hasHTMX && p.HXTarget == "" {
        return errors.New("HTMX method specified but no target provided")
    }
    
    return nil
}

// 6. Smart class generation with conflict resolution
func smartButtonClasses(props ButtonProps, additionalClasses ...string) string {
    baseClasses := []string{"btn"}
    
    // Core variant and size
    variantClass := fmt.Sprintf("btn-%s", props.Variant)
    sizeClass := fmt.Sprintf("btn-%s", props.Size)
    
    // State-based classes using utils.If for clean conditional logic
    stateClasses := []string{
        utils.If(props.Active, "btn-active"),
        utils.If(props.Loading, "btn-loading"),
        utils.If(props.Disabled, "btn-disabled"),
        utils.If(props.FullWidth, "w-full"),
    }
    
    // Combine all classes
    allClasses := append(baseClasses, variantClass, sizeClass)
    allClasses = append(allClasses, stateClasses...)
    allClasses = append(allClasses, props.ClassName)
    allClasses = append(allClasses, additionalClasses...)
    
    // Use TwMerge to resolve conflicts and remove duplicates
    return utils.TwMerge(allClasses...)
}

// 7. Enhanced attribute building
func buildButtonAttributes(props ButtonProps) templ.Attributes {
    // Base attributes with smart defaults
    base := templ.Attributes{
        "id":   utils.IfElse(props.ID != "", props.ID, utils.RandomID()),
        "type": string(utils.IfElse(props.Type != "", props.Type, ButtonTypeButton)),
        "class": smartButtonClasses(props),
    }
    
    // Conditional attributes using utils.If
    conditionalAttrs := templ.Attributes{
        "disabled":      utils.If(props.Disabled || props.Loading, true),
        "aria-label":    utils.If(props.AriaLabel != "", props.AriaLabel),
        "aria-pressed":  utils.IfElse(props.AriaPressed, "true", "false"),
        "data-testid":   utils.If(props.DataTestID != "", props.DataTestID),
    }
    
    // HTMX attributes
    htmxAttrs := buildHTMXAttributes(props)
    
    // Alpine attributes  
    alpineAttrs := buildAlpineAttributes(props)
    
    // Merge all using utils.MergeAttributes
    return utils.MergeAttributes(
        base,
        conditionalAttrs,
        htmxAttrs,
        alpineAttrs,
        props.DataAttributes,
    )
}
```

### 6.2 Props Inheritance Pattern

```go
// Base props that other components can embed
type BaseComponentProps struct {
    ID          string
    ClassName   string
    DataTestID  string
    DataAttributes map[string]string
}

type BaseFormElementProps struct {
    BaseComponentProps
    Name        string
    Value       string
    Disabled    bool
    Required    bool
    AriaLabel   string
}

// Specific component embeds base
type SelectProps struct {
    BaseFormElementProps
    Options     []SelectOption
    Multiple    bool
    Size        int
}
```

### 6.3 Props Documentation

```go
// ButtonProps defines all available properties for the Button component
type ButtonProps struct {
    // ID is the unique identifier for the button element
    ID          string
    
    // Variant determines the visual style of the button
    // Options: primary, secondary, destructive, outline, ghost, link
    Variant     ButtonVariant
    
    // Size controls the button dimensions and text size
    // Options: xs, sm, md, lg, xl
    Size        ButtonSize
    
    // Type is the HTML button type attribute
    // Default: "button"
    Type        ButtonType
    
    // Disabled prevents user interaction
    // Also applied when Loading is true
    Disabled    bool
    
    // Loading shows a spinner and prevents interaction
    Loading     bool
    
    // ... additional documentation for each prop
}
```

## 7. HTMX + Alpine.js Integration Examples

### 7.1 Real-time Search with Filtering

```go
// Example: Product search with real-time filtering
templ ProductSearch() {
    @molecules.SearchBar(molecules.SearchBarProps{
        ID: "product-search",
        Placeholder: "Search products...",
        SearchEndpoint: "/api/products/search",
        HXTarget: "#search-results",
        HXSwap: "innerHTML",
        DebounceMs: 300,
        MinChars: 2,
    })
    
    <div 
        id="search-results"
        x-data="{ filters: { category: '', price: '' } }"
        x-init="$watch('filters', () => htmx.trigger('#product-search input', 'search-filter-change'))"
    >
        // Results will be populated by HTMX
    </div>
}
```

### 7.2 Dynamic Form with Conditional Fields

```go
// Example: User registration form with conditional fields
templ RegistrationForm() {
    @organisms.Form(organisms.FormProps{
        ID: "registration-form",
        HXPost: "/api/users/register",
        HXTarget: "#form-container",
        HXSwap: "outerHTML",
        XData: `{
            userType: 'individual',
            showCompanyFields: false,
            init() {
                this.$watch('userType', (value) => {
                    this.showCompanyFields = value === 'company';
                    // Trigger field visibility change via HTMX
                    htmx.ajax('GET', '/api/forms/registration/fields?type=' + value, {
                        target: '#conditional-fields',
                        swap: 'innerHTML'
                    });
                });
            }
        }`,
        ShowUnsavedWarning: true,
        AutoSave: true,
        AutoSaveEndpoint: "/api/users/draft",
    })
}
```

### 7.3 Interactive Data Table with Sorting and Filtering

```go
// Example: User management table
templ UserTable(users []User) {
    @organisms.DataTable(organisms.DataTableProps{
        ID: "user-table",
        Data: usersToTableData(users),
        Filterable: true,
        Sortable: true,
        Selectable: true,
        
        // HTMX for server-side operations
        BulkActions: []organisms.BulkAction{
            {
                ID: "delete",
                Label: "Delete Selected",
                HXDelete: "/api/users/bulk",
                HXConfirm: "Are you sure you want to delete selected users?",
                HXTarget: "#user-table",
            },
        },
        
        // Alpine.js for client-side interactions
        XData: `{
            selectedUsers: [],
            sortBy: 'name',
            sortOrder: 'asc',
            filters: {},
            
            toggleSort(column) {
                if (this.sortBy === column) {
                    this.sortOrder = this.sortOrder === 'asc' ? 'desc' : 'asc';
                } else {
                    this.sortBy = column;
                    this.sortOrder = 'asc';
                }
                this.applySort();
            },
            
            applySort() {
                htmx.ajax('GET', '/api/users?' + new URLSearchParams({
                    sort: this.sortBy,
                    order: this.sortOrder,
                    ...this.filters
                }), {
                    target: '#user-table tbody',
                    swap: 'innerHTML'
                });
            }
        }`,
    })
}
```

### 7.4 Multi-step Wizard with Progress

```go
// Example: Onboarding wizard
templ OnboardingWizard() {
    @templates.FormPageTemplate(templates.FormPageTemplateProps{
        Title: "Welcome to Ruun",
        Width: templates.FormPageMedium,
    }) {
        <div 
            x-data="{
                currentStep: 1,
                totalSteps: 4,
                canProceed: false,
                formData: {},
                
                nextStep() {
                    if (this.currentStep < this.totalSteps && this.canProceed) {
                        this.currentStep++;
                        htmx.ajax('GET', '/api/onboarding/step/' + this.currentStep, {
                            target: '#step-content',
                            swap: 'innerHTML'
                        });
                    }
                },
                
                prevStep() {
                    if (this.currentStep > 1) {
                        this.currentStep--;
                        htmx.ajax('GET', '/api/onboarding/step/' + this.currentStep, {
                            target: '#step-content',
                            swap: 'innerHTML'
                        });
                    }
                },
                
                submitStep() {
                    const form = this.$el.querySelector('#step-form');
                    htmx.ajax('POST', '/api/onboarding/step/' + this.currentStep, {
                        source: form,
                        target: '#step-content'
                    });
                }
            }"
            class="wizard-container"
        >
            // Progress indicator
            <div class="mb-8">
                <div class="flex items-center justify-between">
                    <template x-for="step in totalSteps" :key="step">
                        <div 
                            :class="{
                                'step-complete': step < currentStep,
                                'step-active': step === currentStep,
                                'step-pending': step > currentStep
                            }"
                            class="step-indicator"
                        >
                            <span x-text="step"></span>
                        </div>
                    </template>
                </div>
            </div>
            
            // Dynamic step content loaded via HTMX
            <div id="step-content">
                // Initial step content
            </div>
            
            // Navigation
            <div class="flex justify-between mt-8">
                @atoms.Button(atoms.ButtonProps{
                    Variant: atoms.ButtonOutline,
                    XShow: "currentStep > 1",
                    XOn: map[string]string{"click": "prevStep()"},
                }) {
                    Previous
                }
                
                @atoms.Button(atoms.ButtonProps{
                    Variant: atoms.ButtonPrimary,
                    XShow: "currentStep < totalSteps",
                    XOn: map[string]string{"click": "nextStep()"},
                    XBind: map[string]string{"disabled": "!canProceed"},
                }) {
                    Next
                }
                
                @atoms.Button(atoms.ButtonProps{
                    Variant: atoms.ButtonPrimary,
                    XShow: "currentStep === totalSteps",
                    XOn: map[string]string{"click": "submitStep()"},
                }) {
                    Complete Setup
                }
            </div>
        </div>
    }
}
```

### 7.5 Real-time Notifications

```go
// Example: Notification system with Alpine + HTMX
templ NotificationCenter() {
    <div 
        x-data="{
            notifications: [],
            unreadCount: 0,
            isOpen: false,
            
            init() {
                // Set up SSE for real-time notifications
                this.setupSSE();
                // Fetch initial notifications via HTMX
                htmx.ajax('GET', '/api/notifications', {
                    target: '#notification-list',
                    swap: 'innerHTML'
                });
            },
            
            setupSSE() {
                const eventSource = new EventSource('/api/notifications/stream');
                eventSource.onmessage = (event) => {
                    const notification = JSON.parse(event.data);
                    this.notifications.unshift(notification);
                    this.unreadCount++;
                    this.showToast(notification);
                };
            },
            
            markAsRead(notificationId) {
                htmx.ajax('POST', '/api/notifications/' + notificationId + '/read', {
                    swap: 'none'
                });
                this.unreadCount = Math.max(0, this.unreadCount - 1);
            },
            
            showToast(notification) {
                // Show temporary toast notification
                this.$dispatch('toast', notification);
            }
        }"
        class="notification-center"
    >
        // Notification bell with badge
        @atoms.Button(atoms.ButtonProps{
            Variant: atoms.ButtonGhost,
            XOn: map[string]string{"click": "isOpen = !isOpen"},
            ClassName: "relative",
        }) {
            @atoms.Icon(atoms.IconProps{Name: "bell", Size: "md"})
            <span 
                x-show="unreadCount > 0"
                x-text="unreadCount"
                class="notification-badge"
            />
        }
        
        // Dropdown panel
        <div 
            x-show="isOpen"
            x-transition
            @click.outside="isOpen = false"
            class="notification-dropdown"
        >
            <div class="notification-header">
                <h3>Notifications</h3>
                @atoms.Button(atoms.ButtonProps{
                    Variant: atoms.ButtonLink,
                    Size: atoms.ButtonSM,
                    HXPost: "/api/notifications/mark-all-read",
                    HXSwap: "none",
                    XOn: map[string]string{"click": "unreadCount = 0"},
                }) {
                    Mark all as read
                }
            </div>
            
            <div id="notification-list">
                // Notifications loaded via HTMX
            </div>
        </div>
    </div>
}
```

## 8. Event Handling Patterns

### 8.1 HTMX Events

```go
// Props for HTMX event handling
type HTMXEventProps struct {
    BeforeRequest   string  // hx-on:htmx:beforeRequest
    AfterRequest    string  // hx-on:htmx:afterRequest  
    BeforeSwap      string  // hx-on:htmx:beforeSwap
    AfterSwap       string  // hx-on:htmx:afterSwap
    ResponseError   string  // hx-on:htmx:responseError
    SendError       string  // hx-on:htmx:sendError
}

// Example usage with utils integration
templ APIButton(action string, target string, loadingText string, htmxEvents HTMXEventProps) {
    @atoms.Button(NewButton().
        Variant(atoms.ButtonPrimary).
        WithHTMX(action, target).
        WithClasses("api-button").
        Conditional(loadingText != "", func(b *ButtonBuilder) {
            b.props.XBind = map[string]string{
                "disabled": "loading",
            }
            b.props.XText = utils.IfElse(loadingText != "", 
                fmt.Sprintf("loading ? '%s' : $el.textContent", loadingText),
                "",
            )
        }).
        Build(),
        mergeHTMXEvents(htmxEvents),
    )
}

func mergeHTMXEvents(events HTMXEventProps) templ.Attributes {
    return utils.MergeAttributes(
        utils.If(events.BeforeRequest != "", templ.Attributes{
            "hx-on:htmx:beforeRequest": events.BeforeRequest,
        }),
        utils.If(events.AfterRequest != "", templ.Attributes{
            "hx-on:htmx:afterRequest": events.AfterRequest,
        }),
        utils.If(events.ResponseError != "", templ.Attributes{
            "hx-on:htmx:responseError": events.ResponseError,
        }),
    )
}
```

### 8.2 Alpine.js Event Patterns

```go
// Common Alpine event handlers as props
type AlpineEventProps struct {
    OnInit          string
    OnShow          string  
    OnHide          string
    OnResize        string
    OnKeydown       string
    OnFocus         string
    OnBlur          string
    OnSubmit        string
    OnChange        string
}

// Usage in form fields
templ SmartInput(props InputProps, events AlpineEventProps) {
    @atoms.Input(atoms.InputProps{
        ...props,
        XOn: map[string]string{
            "init":    events.OnInit,
            "focus":   events.OnFocus,
            "blur":    events.OnBlur,
            "change":  events.OnChange,
            "keydown": events.OnKeydown,
        },
    })
}
```

### 8.3 Custom Event Communication

```go
// Component that dispatches custom events
templ FileUploader(props FileUploaderProps) {
    <div 
        x-data="fileUploader()"
        x-on:file-uploaded.window="handleFileUploaded($event)"
        x-on:upload-error.window="handleUploadError($event)"
    >
        @atoms.Input(atoms.InputProps{
            Type: atoms.InputFile,
            XOn: map[string]string{
                "change": "uploadFiles($event)",
            },
        })
        
        <div x-show="uploading" class="upload-progress">
            <div x-bind:style="'width: ' + uploadProgress + '%'"></div>
        </div>
    </div>
}

script fileUploader() {
    return {
        uploading: false,
        uploadProgress: 0,
        
        uploadFiles(event) {
            const files = event.target.files;
            this.uploading = true;
            
            // Upload logic with progress
            // Dispatch events for other components
            this.$dispatch('upload-started', { files: files.length });
        },
        
        handleFileUploaded(event) {
            // Handle successful upload
            this.uploading = false;
            this.$dispatch('toast', { 
                type: 'success', 
                message: 'File uploaded successfully' 
            });
        }
    }
}
```

## Conclusion

This enhanced props-based architecture with HTMX and Alpine.js integration provides:

- **Type Safety**: Strongly typed props at every level including HTMX/Alpine attributes
- **Rich Interactions**: Seamless server-client communication via HTMX
- **Reactive UI**: Client-side reactivity with Alpine.js state management  
- **Event Handling**: Comprehensive event system for component communication
- **Progressive Enhancement**: Works without JavaScript, enhanced with it
- **Performance**: Minimal JavaScript footprint with server-side rendering
- **Maintainability**: Clear separation between server logic (HTMX) and client state (Alpine)
- **Developer Experience**: Strongly-typed props make components predictable and self-documenting

The architecture respects atomic design principles while providing modern web app functionality through the powerful combination of HTMX for server interactions and Alpine.js for client-side reactivity.

## 9. Utils Integration Benefits

The existing `views/components/utils/util.go` provides significant value to our atomic design system:

### 9.1 TwMerge for Class Conflict Resolution

```go
// Before utils integration - potential conflicts
func buttonClasses(props ButtonProps) string {
    return strings.Join([]string{
        "btn", 
        "btn-" + props.Variant, 
        props.ClassName,  // Could conflict with base classes
    }, " ")
}

// After utils integration - conflicts resolved
func buttonClasses(props ButtonProps) string {
    return utils.TwMerge(
        "btn",
        "btn-" + props.Variant,
        utils.If(props.Active, "bg-blue-600"), 
        utils.If(props.ClassName != "", props.ClassName), // "bg-red-500" would override bg-blue-600
    )
    // Result: "btn btn-primary bg-red-500" (bg-blue-600 removed due to conflict)
}
```

### 9.2 Conditional Logic Simplification

```go
// Before utils - verbose conditional attributes
func buildInputAttrs(props InputProps) templ.Attributes {
    attrs := templ.Attributes{
        "type": string(props.Type),
        "name": props.Name,
    }
    
    if props.Required {
        attrs["required"] = true
    }
    
    if props.Disabled {
        attrs["disabled"] = true
    }
    
    if props.AriaLabel != "" {
        attrs["aria-label"] = props.AriaLabel
    }
    
    return attrs
}

// After utils - clean and readable
func buildInputAttrs(props InputProps) templ.Attributes {
    return utils.MergeAttributes(
        templ.Attributes{
            "type": string(props.Type),
            "name": props.Name,
            "id":   utils.IfElse(props.ID != "", props.ID, utils.RandomID()),
        },
        templ.Attributes{
            "required":    utils.If(props.Required, true),
            "disabled":    utils.If(props.Disabled, true),
            "aria-label":  utils.If(props.AriaLabel != "", props.AriaLabel),
            "class":       utils.TwMerge("input", props.ClassName),
        },
    )
}
```

### 9.3 Component Factory with Utils

```go
// Smart component factory using utils
type ComponentFactory struct {
    defaults map[string]interface{}
}

func NewComponentFactory() *ComponentFactory {
    return &ComponentFactory{
        defaults: map[string]interface{}{
            "buttonVariant": atoms.ButtonPrimary,
            "inputSize":     atoms.InputMD,
        },
    }
}

func (cf *ComponentFactory) CreateButton(overrides ButtonProps) templ.Component {
    // Merge with defaults using utils
    props := utils.IfElse(overrides.Variant != "", overrides, ButtonProps{
        ID:      utils.RandomID(),
        Variant: cf.defaults["buttonVariant"].(atoms.ButtonVariant),
        Size:    atoms.ButtonMD,
        Type:    atoms.ButtonTypeButton,
    })
    
    // Apply smart class merging
    props.ClassName = utils.TwMerge(
        "component-factory-btn",
        props.ClassName,
    )
    
    return atoms.Button(props)
}

// Usage examples
func ExampleButtons() {
    factory := NewComponentFactory()
    
    // Simple button with auto-generated ID
    @factory.CreateButton(ButtonProps{}) { "Click me" }
    
    // Override specific properties
    @factory.CreateButton(ButtonProps{
        Variant: atoms.ButtonDestructive,
        ClassName: "ml-4", // Will be merged with factory defaults
    }) { "Delete" }
}
```

### 9.4 Cache-Busted Script Loading

```go
// Leveraging ScriptVersion for cache busting
templ DashboardLayout(props LayoutProps) {
    <html>
        <head>
            <script src={ fmt.Sprintf("/static/js/htmx.min.js?v=%s", utils.ScriptVersion) }></script>
            <script src={ fmt.Sprintf("/static/js/alpine.min.js?v=%s", utils.ScriptVersion) }></script>
            <script src={ fmt.Sprintf("/static/js/app.js?v=%s", utils.ScriptVersion) }></script>
        </head>
        <body>
            { children... }
        </body>
    </html>
}
```

### 9.5 Smart Defaults and Fallbacks

```go
// Component with intelligent defaults using utils
templ SmartModal(props ModalProps) {
    modalID := utils.IfElse(props.ID != "", props.ID, utils.RandomID())
    
    <div 
        id={ modalID }
        class={ utils.TwMerge(
            "modal",
            utils.If(props.Large, "modal-lg"),
            utils.If(props.Centered, "modal-centered"),
            props.ClassName,
        )}
        x-data={ fmt.Sprintf(`{
            open: %t,
            closeOnEscape: %t,
            closeOnBackdrop: %t,
            autoFocus: %t
        }`, 
            props.Open,
            utils.IfElse(props.CloseOnEscape, true, false),  // Default to true
            utils.IfElse(props.CloseOnBackdrop, true, false), // Default to true  
            utils.IfElse(props.AutoFocus, true, false),      // Default to true
        )}
        x-show="open"
        x-transition
    >
        { children... }
    </div>
}
```

### 9.6 Reusable Utility Patterns

```go
// Create reusable utility functions leveraging existing utils
func CreateConditionalClasses(conditions map[string]bool, baseClasses ...string) string {
    var classes []string
    classes = append(classes, baseClasses...)
    
    for class, condition := range conditions {
        classes = append(classes, utils.If(condition, class))
    }
    
    return utils.TwMerge(classes...)
}

// Usage in components
func cardClasses(props CardProps) string {
    return CreateConditionalClasses(
        map[string]bool{
            "card-elevated":  props.Elevated,
            "card-bordered":  props.Bordered,
            "card-compact":   props.Compact,
            "card-loading":   props.Loading,
        },
        "card",
        fmt.Sprintf("card-%s", props.Variant),
        props.ClassName,
    )
}
```

The utils integration provides:
- **Cleaner Code**: Less verbose conditional logic
- **Better Performance**: TwMerge optimizes CSS class conflicts
- **Consistent IDs**: RandomID ensures unique component IDs
- **Smart Defaults**: IfElse provides fallback values
- **Cache Busting**: ScriptVersion handles asset versioning
- **Maintainability**: Centralized utility functions reduce duplication