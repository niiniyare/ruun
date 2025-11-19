# Enhanced Form Organism

The Enhanced Form organism provides a comprehensive form system that integrates business logic, schema-driven generation, and advanced validation capabilities. It leverages the refactored FormField molecule and provides seamless integration with the `pkg/schema` system.

## Features

- **Schema Integration**: Automatic form generation from schema definitions
- **Business Logic Integration**: Built-in form state management and data transformation
- **Real-time Validation**: Client and server-side validation with debouncing
- **Auto-save Functionality**: Automatic draft saving with configurable intervals
- **Multi-step Forms**: Wizard-style forms with progress tracking
- **Conditional Logic**: Dynamic field visibility based on other field values
- **HTMX Integration**: Real-time form updates and validation
- **Alpine.js Integration**: Reactive form state and interactions
- **Compiled Theme Classes**: Uses compiled theme classes instead of runtime token resolution
- **Accessibility**: Full ARIA support and keyboard navigation
- **Internationalization**: Built-in i18n support

## Basic Usage

### Simple Form

```go
@EnhancedForm(EnhancedFormProps{
    ID:                "contact-form",
    Title:             "Contact Us",
    Description:       "Send us a message",
    SubmitURL:         "/api/contact",
    ValidationStrategy: "realtime",
    Fields: []molecules.FormFieldProps{
        {
            ID:       "name",
            Name:     "name",
            Label:    "Full Name",
            Type:     "text",
            Required: true,
        },
        {
            ID:       "email",
            Name:     "email",
            Label:    "Email",
            Type:     "email",
            Required: true,
        },
    },
})
```

### Schema-Driven Form

```go
// Using the schema builder
form := CreateCRUDForm(userSchema).
    WithTitle("User Management").
    WithValidationURL("/api/users/validate").
    WithAutoSave(true, 30).
    Build()

@EnhancedForm(form)
```

## Form Properties

### Core Properties

- `ID`: Unique form identifier
- `Title`: Form title displayed at the top
- `Description`: Form description text
- `Schema`: Optional schema definition for automatic form generation
- `Layout`: Form layout (`vertical`, `horizontal`, `inline`, `grid`)
- `Size`: Form size (`sm`, `md`, `lg`)
- `ReadOnly`: Make the entire form read-only

### State Management

- `InitialData`: Pre-populate form with data
- `AutoSave`: Enable automatic draft saving
- `AutoSaveInterval`: Auto-save interval in seconds (default: 30)
- `ShowSaveState`: Show save status indicator

### Validation

- `ValidationStrategy`: Validation timing (`realtime`, `onblur`, `onsubmit`)
- `ValidationURL`: Server-side validation endpoint
- `SubmitURL`: Form submission endpoint

### Layout Options

#### Grid Layout

```go
EnhancedFormProps{
    Layout:   FormLayoutGrid,
    GridCols: 2, // Number of columns
    Fields:   [...],
}
```

#### Multi-step Layout

```go
EnhancedFormProps{
    ShowProgress: true,
    Sections: []EnhancedFormSection{
        {
            Title: "Personal Information",
            Fields: [...],
        },
        {
            Title: "Contact Details", 
            Fields: [...],
        },
    },
}
```

## Form Sections

Sections allow you to group related fields together:

```go
EnhancedFormSection{
    ID:          "personal-info",
    Title:       "Personal Information",
    Description: "Enter your personal details",
    Icon:        "user",
    Collapsible: true,
    Collapsed:   false,
    Fields:      [...],
    Conditional: "formData.userType === 'individual'", // Show/hide condition
}
```

### Section Properties

- `Collapsible`: Allow users to collapse/expand the section
- `Collapsed`: Initial collapsed state
- `Required`: Mark section as required
- `Conditional`: JavaScript expression for conditional display
- `Layout`: Section-specific layout override

## Field Integration

The Enhanced Form automatically integrates with the refactored FormField molecule:

```go
molecules.FormFieldProps{
    // Basic properties
    ID:          "email",
    Name:        "email", 
    Label:       "Email Address",
    Type:        "email",
    Required:    true,
    
    // Validation integration
    ValidationState:    molecules.ValidationStateIdle,
    OnValidate:         "/api/validate/email",
    ValidationDebounce: 300,
    
    // Alpine.js integration
    AlpineModel:  "formData.email",
    AlpineChange: "updateField('email', $event.target.value)",
    AlpineBlur:   "validateField('email')",
    
    // HTMX integration
    HXPost:    "/api/validate/email",
    HXTarget:  "#email-feedback",
    HXTrigger: "change delay:500ms",
}
```

## Actions

Define custom actions for your form:

```go
EnhancedFormAction{
    ID:       "save-draft",
    Type:     "save-draft",
    Text:     "Save Draft",
    Variant:  atoms.ButtonGhost,
    Position: "left", // "left", "center", "right"
    OnClick:  "autoSave()",
    Condition: "isDirty", // Show only when form has changes
}
```

### Action Types

- `submit`: Form submission button
- `reset`: Reset form to initial state
- `cancel`: Cancel form editing
- `save-draft`: Save current state as draft
- `button`: Custom button action

## Schema Integration

### Using Schema Builder

The `SchemaFormBuilder` provides a fluent interface for creating forms from schema definitions:

```go
// Basic schema form
form := NewSchemaFormBuilder(schema).
    WithValidationStrategy("realtime").
    WithAutoSave(true, 30).
    Build()

// CRUD form with common actions
form := CreateCRUDForm(schema).
    WithTitle("User Management").
    WithValidationURL("/api/users/validate").
    Build()

// Multi-step wizard
form := CreateWizardForm(schema).
    WithProgress(true).
    Build()

// Read-only form
form := CreateReadOnlyForm(schema).
    WithTitle("User Profile").
    Build()
```

### Schema Field Mapping

The system automatically maps schema field types to form field types:

| Schema Type | Form Type | Notes |
|-------------|-----------|-------|
| `string` | `text` | Basic text input |
| `email` | `email` | Email validation |
| `password` | `password` | Password input |
| `number` | `number` | Numeric input |
| `text` | `textarea` | Multi-line text |
| `select` | `select` | Dropdown selection |
| `radio` | `radio` | Radio button group |
| `checkbox` | `checkbox` | Checkbox group |
| `date` | `date` | Date picker |
| `file` | `file` | File upload |

## Validation

### Validation Strategies

1. **Real-time**: Validate as user types (with debouncing)
2. **On Blur**: Validate when field loses focus
3. **On Submit**: Validate only when form is submitted

### Client-side Validation

```javascript
// Built-in validation
validateField(fieldName) {
    const field = this.getFieldSchema(fieldName);
    const value = this.formData[fieldName];
    
    // Required validation
    if (field.required && !value) {
        this.errors[fieldName] = 'This field is required';
        return;
    }
    
    // Length validation
    if (field.minLength && value.length < field.minLength) {
        this.errors[fieldName] = `Minimum length is ${field.minLength}`;
        return;
    }
    
    // Pattern validation
    if (field.pattern && !new RegExp(field.pattern).test(value)) {
        this.errors[fieldName] = 'Invalid format';
        return;
    }
    
    // Clear errors if validation passes
    delete this.errors[fieldName];
}
```

### Server-side Validation

Configure server-side validation endpoints:

```go
EnhancedFormProps{
    ValidationURL: "/api/validate",
    ValidationStrategy: "realtime",
}
```

The form will send validation requests to your endpoint:

```json
{
    "field": "email",
    "value": "user@example.com",
    "formData": { ... }
}
```

Expected response format:

```json
{
    "valid": true,
    "error": null
}
// or
{
    "valid": false, 
    "error": "Email already exists"
}
```

## Auto-save

Enable automatic draft saving:

```go
EnhancedFormProps{
    AutoSave:         true,
    AutoSaveInterval: 30, // seconds
    AutoSaveURL:     "/api/forms/auto-save",
    ShowSaveState:   true, // Show save status indicator
}
```

The form will automatically save drafts when:
- Form data changes and is valid
- User stops interacting for the specified interval
- No submission is in progress

## Conditional Logic

Show/hide fields based on other field values:

```go
// Field-level conditional
molecules.FormFieldProps{
    Name:      "salePrice",
    Condition: "formData.onSale === true", // JavaScript expression
}

// Section-level conditional
EnhancedFormSection{
    Name:        "billing",
    Conditional: "formData.needsBilling === true",
}
```

### Conditional Expressions

Use JavaScript expressions to define conditions:

```javascript
// Simple equality
"formData.userType === 'premium'"

// Multiple conditions
"formData.age >= 18 && formData.country === 'US'"

// Array inclusion
"['admin', 'moderator'].includes(formData.role)"

// Field existence
"formData.email && formData.email !== ''"
```

## Event Handling

### Form-level Events

```go
EnhancedFormProps{
    OnSubmit:      "handleFormSubmit(formData)",
    OnValidate:    "handleValidation($event.detail)",
    OnFieldChange: "handleFieldChange($event.detail)",
    OnAutoSave:    "handleAutoSave()",
    OnReset:       "handleFormReset()",
}
```

### Field-level Events

```go
molecules.FormFieldProps{
    AlpineChange: "updateField('email', $event.target.value)",
    AlpineBlur:   "validateField('email')",
    AlpineFocus:  "trackFieldFocus('email')",
}
```

## HTMX Integration

### Form-level HTMX

```go
EnhancedFormProps{
    HXTarget:    "#form-result",
    HXSwap:      "innerHTML",
    HXTrigger:   "submit",
    HXIndicator: "#form-spinner",
}
```

### Field-level HTMX

```go
molecules.FormFieldProps{
    HXPost:    "/api/validate/username",
    HXTarget:  "#username-feedback",
    HXTrigger: "change delay:500ms",
    HXSwap:    "innerHTML",
}
```

## Theme Integration

The Enhanced Form uses compiled theme classes for optimal performance:

```go
EnhancedFormProps{
    ThemeID: "corporate-theme",
    TokenOverrides: map[string]string{
        "form.background": "#ffffff",
        "form.border":     "#e2e8f0", 
    },
    DarkMode: false,
}
```

## Alpine.js State

The form exposes comprehensive state through Alpine.js:

```javascript
{
    // Form data
    formData: {...},
    initialData: {...},
    
    // Validation state
    errors: {},
    touched: {},
    dirty: {},
    valid: true,
    validating: {},
    
    // Form lifecycle
    loading: false,
    submitting: false,
    saving: false,
    
    // Progress tracking
    currentStep: 1,
    totalSteps: 3,
    
    // Computed properties
    isDirty: true,
    hasErrors: false,
    progressPercentage: "33%",
    
    // Methods
    updateField(name, value),
    validateField(name),
    submitForm(),
    resetForm(),
    nextStep(),
    prevStep(),
    autoSave()
}
```

## Examples

See `enhanced_form_examples.templ` for comprehensive examples including:

- Basic contact form
- Multi-step user registration
- Grid layout product form
- Schema-driven forms
- Read-only data display

## Best Practices

1. **Use Schema Integration**: Leverage the schema system for consistent forms
2. **Progressive Enhancement**: Start with basic HTML forms, enhance with HTMX/Alpine
3. **Validation Strategy**: Choose appropriate validation timing for your use case
4. **Auto-save Wisely**: Use auto-save for long forms, avoid for short forms
5. **Conditional Logic**: Keep conditions simple and testable
6. **Error Handling**: Provide clear, actionable error messages
7. **Accessibility**: Always include proper labels and ARIA attributes
8. **Performance**: Use compiled themes and avoid runtime token resolution

## Migration from Legacy Forms

To migrate from the existing form system:

1. Replace `@organisms.Form()` with `@organisms.EnhancedForm()`
2. Update field props to use `molecules.FormFieldProps`
3. Add validation strategy and URLs
4. Configure auto-save if needed
5. Update event handlers to use Alpine.js syntax
6. Test conditional logic and validation

## Troubleshooting

### Common Issues

1. **Validation not working**: Check `ValidationURL` and server response format
2. **Auto-save failing**: Verify `AutoSaveURL` endpoint and response
3. **Conditional fields not showing**: Check JavaScript expression syntax
4. **HTMX not triggering**: Verify target selectors and trigger events
5. **Theme classes not applied**: Ensure theme compilation is complete

### Debug Mode

Enable debug mode to inspect form state:

```go
EnhancedFormProps{
    Debug: true,
}
```

This will render a debug panel showing:
- Current form state
- Validation errors  
- Schema information
- Alpine.js data