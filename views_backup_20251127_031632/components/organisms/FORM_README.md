# Form Organism

The Form organism provides progressive enhancement from simple contact forms to enterprise-grade form systems with advanced validation, auto-save, persistence, and dependency management. Built with Alpine.js and HTMX integration for progressive enhancement.

## Progressive Enhancement Philosophy

The form scales from simple to enterprise complexity using feature flags:

- **Simple**: Just ID + Fields
- **Basic**: Add validation and layout
- **Advanced**: Enable auto-save, storage, dependencies
- **Enterprise**: Full feature set with SSE, cross-tab sync, offline mode

## Features

- **Progressive Enhancement**: Scales from simple to enterprise complexity
- **Component Consistency**: Uses atoms/molecules for styling inheritance
- **Type-Safe Configuration**: Compile-time safety with Go enums
- **Zero-Config Binding**: Automatic Alpine.js integration
- **Server-Side Validation**: With optional SSE support for long-running validations
- **Field Dependencies**: Cascading fields with multiple dependency types
- **Auto-save**: Multiple strategies (debounced, immediate, interval, manual)
- **Persistence**: localStorage, sessionStorage, or IndexedDB
- **Cross-Tab Sync**: Real-time synchronization across browser tabs
- **Offline Support**: Network-aware with offline detection
- **Multi-step Forms**: Wizard-style forms with progress tracking
- **HTMX Integration**: Progressive enhancement for real-time updates
- **Full Accessibility**: ARIA support and keyboard navigation
- **Debug Mode**: Comprehensive debugging panel

## Basic Usage

### Simple Contact Form
```go
// Minimal form - just core fields
@organisms.Form(organisms.FormProps{
    ID: "contact",
    Title: "Contact Us",
    Fields: []organisms.Field{
        {FormFieldProps: molecules.FormFieldProps{Name: "name", Type: "text", Label: "Name"}},
        {FormFieldProps: molecules.FormFieldProps{Name: "email", Type: "email", Label: "Email"}},
        {FormFieldProps: molecules.FormFieldProps{Name: "message", Type: "textarea", Label: "Message"}},
    },
    SubmitURL: "/contact",
})
```

### Form with Basic Enhancement
```go
// Add validation and better layout
@organisms.Form(organisms.FormProps{
    ID: "user-profile",
    Title: "User Profile",
    Description: "Update your information",
    Layout: organisms.FormLayoutGrid,
    Size: organisms.FormSizeMD,
    
    Fields: []organisms.Field{
        {FormFieldProps: molecules.FormFieldProps{Name: "firstName", Type: "text", Label: "First Name"}},
        {FormFieldProps: molecules.FormFieldProps{Name: "lastName", Type: "text", Label: "Last Name"}},
    },
    
    // Enable validation
    Validation: &organisms.ValidationConfig{
        Strategy: organisms.ValidationOnBlur,
        URL:      "/validate",
    },
    
    SubmitURL: "/profile/update",
})
```

### Enterprise Form with All Features
```go
// Full enterprise feature set
@organisms.Form(organisms.FormProps{
    ID: "enterprise-form",
    Title: "Advanced User Profile",
    Description: "Complete profile with enterprise features",
    Layout: organisms.FormLayoutGrid,
    Size: organisms.FormSizeLG,
    
    // Core fields with dependencies
    Fields: []organisms.Field{
        {
            FormFieldProps: molecules.FormFieldProps{Name: "country", Type: "select", Label: "Country"},
            Dependencies: []organisms.Dependency{
                {TargetField: "state", Type: organisms.DependencyOptions, OptionsURL: "/states/{value}"},
            },
        },
        {
            FormFieldProps: molecules.FormFieldProps{Name: "state", Type: "select", Label: "State"},
            Conditional: "country !== ''",
        },
    },
    
    // Sections for organization
    Sections: []organisms.Section{
        {
            Title: "Personal Information",
            Icon: "user",
            Collapsible: true,
            Fields: []organisms.Field{
                {FormFieldProps: molecules.FormFieldProps{Name: "bio", Type: "textarea", Label: "Biography"}},
            },
        },
    },
    
    // Enterprise features (progressive enhancement)
    Advanced: &organisms.AdvancedConfig{
        EnableSSE:          true,
        EnableCrossTabSync: true,
        EnableOfflineMode:  true,
    },
    
    Validation: &organisms.ValidationConfig{
        Strategy: organisms.ValidationRealtime,
        URL:      "/validate",
        Debounce: 500,
    },
    
    AutoSave: &organisms.AutoSaveConfig{
        Enabled:   true,
        Strategy:  organisms.AutoSaveDebounced,
        Interval:  30,
        URL:       "/autosave",
        ShowState: true,
    },
    
    Storage: &organisms.StorageConfig{
        Strategy:       organisms.StorageLocal,
        RestoreOnLoad:  true,
        TTL:           86400,
    },
    
    Progress: &organisms.ProgressConfig{
        ShowProgress: true,
        CurrentStep:  1,
        TotalSteps:   3,
    },
    
    Dependencies: &organisms.DependencyConfig{
        DetectCycles: true,
    },
    
    Debug: &organisms.DebugConfig{
        Enabled:      true,
        ShowPanel:    true,
        LogToConsole: true,
    },
    
    // Custom actions using atoms.Button
    Actions: []organisms.Action{
        {
            Type:     "button",
            Label:    "Save Draft",
            Position: "left",
            ButtonProps: atoms.ButtonProps{
                Variant: atoms.ButtonVariantSecondary,
            },
        },
    },
    
    SubmitURL: "/profile/update",
})
```

## Architecture Benefits

### 1. Progressive Enhancement
- Start with simple form, add complexity as needed
- Feature flags automatically enable/disable functionality
- No overhead for unused features

### 2. Component Consistency  
- Uses `@atoms.Button()` instead of `<button>`
- Uses `@atoms.Icon()` instead of `<svg>`
- Uses semantic HTML with consistent CSS classes
- Styling inheritance through design tokens

### 3. Feature Detection
The form automatically enables features based on configuration:

```go
// Automatically detected feature flags
hasValidation   := props.Validation != nil && props.Validation.URL != ""
hasAutoSave     := props.AutoSave != nil && props.AutoSave.Enabled  
hasStorage      := props.Storage != nil && props.Storage.Strategy != StorageNone
hasProgress     := props.Progress != nil && props.Progress.TotalSteps > 1
hasDependencies := props.Dependencies != nil && len(props.Dependencies.Graph) > 0
hasSSE          := (props.Validation != nil && props.Validation.SSE != nil) || 
                   (props.Advanced != nil && props.Advanced.EnableSSE)
hasAdvanced     := props.Advanced != nil
hasDebug        := props.Debug != nil && props.Debug.Enabled
```

## Form Properties

### Core Properties (Always Available)
- `ID` (string): Unique form identifier
- `Fields` ([]Field): Form fields (required)

### Basic Enhancement  
- `Title` (string): Form title
- `Description` (string): Form description
- `Layout` (FormLayout): `FormLayoutVertical` (default), `FormLayoutHorizontal`, `FormLayoutGrid`
- `Size` (FormSize): `FormSizeSM`, `FormSizeMD` (default), `FormSizeLG`
- `ClassName` (string): Custom CSS classes
- `ReadOnly` (bool): Make entire form read-only
- `SubmitURL` (string): Form submission endpoint
- `OnSubmit` (string): Alpine.js submit handler

### Progressive Enhancement (Nil = Disabled)

#### Validation
```go
Validation: &organisms.ValidationConfig{
    Strategy: organisms.ValidationOnBlur,    // When to validate
    URL:      "/validate",                   // Server endpoint
    Debounce: 300,                          // ms delay
    SSE:      &organisms.SSEConfig{...},    // Long-running validation
}
```

#### Auto-Save
```go
AutoSave: &organisms.AutoSaveConfig{
    Enabled:   true,
    Strategy:  organisms.AutoSaveDebounced,  // How to save
    Interval:  30,                          // seconds  
    URL:       "/autosave",                 // Server endpoint
    ShowState: true,                        // Show save indicator
}
```

#### Storage/Persistence
```go
Storage: &organisms.StorageConfig{
    Strategy:       organisms.StorageLocal,     // Where to store
    Key:           "form-key",                  // Storage key
    TTL:           86400,                       // Time to live (seconds)
    SyncAcrossTabs: true,                      // Cross-tab sync
    RestoreOnLoad:  true,                      // Restore on mount
}
```

#### Advanced Features
```go
Advanced: &organisms.AdvancedConfig{
    EnableSSE:          true,    // Server-Sent Events
    EnableCrossTabSync: true,    // Real-time tab sync
    EnableOfflineMode:  true,    // Offline support
    EnableTelemetry:    false,   // Usage analytics
}
```

#### Multi-step Progress
```go
Progress: &organisms.ProgressConfig{
    ShowProgress: true,
    CurrentStep:  1,
    TotalSteps:   3,
}
```

#### Field Dependencies
```go
Dependencies: &organisms.DependencyConfig{
    Graph: map[string][]organisms.Dependency{
        "country": {{TargetField: "state", Type: organisms.DependencyOptions}},
    },
    DetectCycles: true,
}
```

#### Debug Tools
```go
Debug: &organisms.DebugConfig{
    Enabled:      true,     // Enable debug features
    ShowPanel:    true,     // Show debug panel
    LogToConsole: true,     // Console logging
}
```

## Field Enhancement

### Basic Field
```go
organisms.Field{
    FormFieldProps: molecules.FormFieldProps{
        Name:     "email",
        Label:    "Email Address", 
        Type:     "email",
        Required: true,
    },
}
```

### Field with Progressive Features
```go
organisms.Field{
    FormFieldProps: molecules.FormFieldProps{
        Name:  "country",
        Label: "Country",
        Type:  "select",
    },
    
    // Progressive enhancement
    Conditional:  "userType === 'international'",  // Visibility
    Dependencies: []organisms.Dependency{           // Field relationships
        {
            TargetField: "state",
            Type:        organisms.DependencyOptions,
            OptionsURL:  "/states/{value}",
            ClearOnChange: true,
        },
    },
    AutoSave:   &organisms.FieldAutoSave{...},     // Field-level auto-save
    Validation: &organisms.FieldValidation{...},   // Field-level validation
    Storage:    &organisms.FieldStorage{...},      // Storage exclusion
}
```

## Dependency Types

### 1. Options (Cascading Dropdowns)
```go
organisms.Dependency{
    TargetField:   "state",
    Type:          organisms.DependencyOptions,
    OptionsURL:    "/api/states?country={value}",
    ClearOnChange: true,
    Debounce:      300,
}
```

### 2. Calculated Values
```go
organisms.Dependency{
    TargetField:     "total",
    Type:            organisms.DependencyValue, 
    ValueExpression: "quantity * price",
}
```

### 3. Conditional Visibility
```go
organisms.Dependency{
    TargetField: "phoneNumber",
    Type:        organisms.DependencyVisibility,
    Condition:   "contactMethod === 'phone'",
}
```

### 4. Conditional Required
```go
organisms.Dependency{
    TargetField: "cardNumber", 
    Type:        organisms.DependencyRequired,
    Condition:   "paymentMethod === 'card'",
}
```

### 5. Dynamic Validation
```go
organisms.Dependency{
    TargetField:     "confirmPassword",
    Type:            organisms.DependencyValidation,
    ValidationRules: `{"match": "password"}`,
}
```

## Actions System

Actions use `atoms.Button` for consistency:

```go
organisms.Action{
    Type:     "submit",           // "submit" | "button" | "reset"
    Label:    "Save Profile",     // Button text
    Position: "right",            // "left" | "center" | "right"
    
    // Atoms integration
    ButtonProps: atoms.ButtonProps{
        Variant: atoms.ButtonVariantPrimary,
        Size:    atoms.ButtonSizeMD,
    },
    
    // Behavior
    OnClick:     "customHandler()",
    Conditional: "isDirty",        // Alpine.js visibility
    
    // HTMX support
    HTMX: &organisms.ActionHTMX{
        Post:    "/api/save",
        Target:  "#result",
        Confirm: "Are you sure?",
    },
}
```

## Validation Strategies

```go
type ValidationStrategy string

const (
    ValidationRealtime ValidationStrategy = "realtime" // Keystroke + debounce
    ValidationOnBlur   ValidationStrategy = "onblur"   // Field blur (default)  
    ValidationOnChange ValidationStrategy = "onchange" // Value change + debounce
    ValidationOnSubmit ValidationStrategy = "onsubmit" // Submit only
)
```

### Server-Side Validation

**Request format:**
```json
{
    "field": "email",
    "value": "user@example.com", 
    "formData": {...}
}
```

**Response format:**
```json
{
    "valid": true,
    "message": ""
}
```

## Auto-Save Strategies

```go
type AutoSaveStrategy string

const (
    AutoSaveDebounced AutoSaveStrategy = "debounced" // Wait for inactivity
    AutoSaveImmediate AutoSaveStrategy = "immediate" // Save immediately  
    AutoSaveInterval  AutoSaveStrategy = "interval"  // Fixed intervals
    AutoSaveManual    AutoSaveStrategy = "manual"    // Manual trigger only
)
```

## Storage Strategies

```go
type StorageStrategy string

const (
    StorageNone      StorageStrategy = "none"      // No persistence
    StorageLocal     StorageStrategy = "local"     // localStorage
    StorageSession   StorageStrategy = "session"   // sessionStorage
    StorageIndexedDB StorageStrategy = "indexeddb" // IndexedDB
)
```

## Multi-Step Forms

```go
@organisms.Form(organisms.FormProps{
    // Enable progress tracking
    Progress: &organisms.ProgressConfig{
        ShowProgress: true,
        CurrentStep:  1,
        TotalSteps:   3,
    },
    
    // Organize fields into sections
    Sections: []organisms.Section{
        {Title: "Personal Info", Fields: [...]},
        {Title: "Contact Details", Fields: [...]}, 
        {Title: "Preferences", Fields: [...]},
    },
    
    // Custom navigation actions
    Actions: []organisms.Action{
        {
            Label:       "Back",
            OnClick:     "prevStep()",
            Position:    "left",
            Conditional: "currentStep > 1",
        },
        {
            Label:       "Next", 
            OnClick:     "nextStep()",
            Position:    "right",
            Conditional: "currentStep < totalSteps",
        },
        {
            Type:        "submit",
            Label:       "Complete",
            Position:    "right", 
            Conditional: "currentStep === totalSteps",
        },
    },
})
```

## Alpine.js State Management

The form exposes comprehensive reactive state:

```javascript
// Minimal form state
{
    formData: {...},              // Form values
    errors: {},                   // Validation errors
    touched: {},                  // Touched fields
    dirty: {},                    // Changed fields
    loading: false,               // Loading state
    submitting: false,            // Submit state
    
    // Computed
    isDirty: boolean,             // Has changes
    hasErrors: boolean,           // Has validation errors
    canSubmit: boolean,           // Can submit form
}

// Progressive enhancement adds:
{
    // Auto-save (if enabled)
    saving: false,
    lastSaved: Date,
    
    // Storage (if enabled)  
    restoredFromStorage: false,
    
    // Progress (if enabled)
    currentStep: 1,
    totalSteps: 3,
    progressPercentage: 33,
    
    // Dependencies (if enabled)
    dependencyCache: {},
    
    // SSE (if enabled)
    sseConnections: {},
    sseStatus: {},
    
    // Debug (if enabled)
    showDebug: false,
    debugTab: 'state',
}
```

## Debug Mode

Enable comprehensive debugging:

```go
Debug: &organisms.DebugConfig{
    Enabled:      true,    // Enable debug features
    ShowPanel:    true,    // Show debug panel
    LogToConsole: true,    // Console logging
}
```

Toggle panel with **Ctrl+Shift+D** or call `toggleDebug()`.

Debug panel shows:
- **State**: Form data, errors, touched fields
- **Dependencies**: Dependency graph and cached options  
- **Storage**: Persistence info and saved data
- **SSE**: Active connections and status

## Performance Optimizations

1. **Progressive Loading**: Only load JS for enabled features
2. **Feature Detection**: Automatic based on configuration
3. **Component Reuse**: Consistent atoms/molecules usage
4. **Debouncing**: Built-in for validation and dependencies
5. **Caching**: Dependency options cached client-side
6. **Minimal State**: Only track what's needed

## Best Practices

1. **Start Simple**: Begin with minimal props, add features as needed
2. **Type Safety**: Use typed enums (`ValidationOnBlur` not `"onblur"`)
3. **Component Consistency**: Leverage atoms/molecules for styling
4. **Security**: Set `ExcludeFromStorage` for sensitive fields
5. **Performance**: Use appropriate debounce values (300ms default)
6. **Validation**: Client-side for UX, server-side for security
7. **Storage**: Choose strategy based on data sensitivity and size
8. **Dependencies**: Keep expressions simple and testable

## Complete Example: Enterprise User Profile

```go
@organisms.Form(organisms.FormProps{
    ID:          "user-profile",
    Title:       "User Profile",
    Description: "Complete your profile information",
    Layout:      organisms.FormLayoutGrid,
    
    // Core fields
    Fields: []organisms.Field{
        {
            FormFieldProps: molecules.FormFieldProps{
                Name: "firstName", Type: "text", Label: "First Name", Required: true,
            },
        },
        {
            FormFieldProps: molecules.FormFieldProps{
                Name: "lastName", Type: "text", Label: "Last Name", Required: true,
            },
        },
        {
            FormFieldProps: molecules.FormFieldProps{
                Name: "country", Type: "select", Label: "Country",
            },
            Dependencies: []organisms.Dependency{
                {TargetField: "state", Type: organisms.DependencyOptions, OptionsURL: "/states/{value}"},
            },
        },
        {
            FormFieldProps: molecules.FormFieldProps{
                Name: "state", Type: "select", Label: "State",
            },
            Conditional: "country !== ''",
        },
    },
    
    // Optional sections
    Sections: []organisms.Section{
        {
            Title: "Additional Information",
            Icon: "info", 
            Collapsible: true,
            Fields: []organisms.Field{
                {FormFieldProps: molecules.FormFieldProps{Name: "bio", Type: "textarea", Label: "Bio"}},
                {FormFieldProps: molecules.FormFieldProps{Name: "website", Type: "url", Label: "Website"}},
            },
        },
    },
    
    // Enterprise features (progressive enhancement)
    Validation: &organisms.ValidationConfig{
        Strategy: organisms.ValidationRealtime,
        URL:      "/validate", 
        Debounce: 500,
    },
    
    AutoSave: &organisms.AutoSaveConfig{
        Enabled:   true,
        Strategy:  organisms.AutoSaveDebounced,
        Interval:  30,
        URL:       "/autosave",
        ShowState: true,
    },
    
    Storage: &organisms.StorageConfig{
        Strategy:       organisms.StorageLocal,
        RestoreOnLoad:  true,
        TTL:           86400,
        SyncAcrossTabs: true,
    },
    
    Advanced: &organisms.AdvancedConfig{
        EnableSSE:          true,
        EnableCrossTabSync: true, 
        EnableOfflineMode:  true,
    },
    
    Dependencies: &organisms.DependencyConfig{
        DetectCycles: true,
    },
    
    Debug: &organisms.DebugConfig{
        Enabled:      true,
        ShowPanel:    true,
        LogToConsole: true,
    },
    
    // Custom actions
    Actions: []organisms.Action{
        {
            Type:     "button",
            Label:    "Save Draft", 
            Position: "left",
            OnClick:  "autoSave()",
            ButtonProps: atoms.ButtonProps{
                Variant: atoms.ButtonVariantSecondary,
            },
        },
        {
            Type:     "submit",
            Label:    "Save Profile",
            Position: "right",
            ButtonProps: atoms.ButtonProps{
                Variant: atoms.ButtonVariantPrimary,
            },
        },
    },
    
    SubmitURL: "/profile/update",
})
```

This unified form scales from a simple 3-field contact form to a complex enterprise form with all advanced features, while maintaining component consistency and performance through progressive enhancement.