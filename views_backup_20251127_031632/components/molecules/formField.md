# Enterprise Form Field Component

A comprehensive, production-ready form field component library for Go/Templ applications with enterprise-grade features.

## Features

✅ **30+ Field Types** - Text, email, password, select, multi-select, date pickers, and more  
✅ **Validation** - Built-in server-side validation with custom rules  
✅ **Accessibility** - WCAG 2.1 AA compliant with proper ARIA attributes  
✅ **HTMX Integration** - Full support for HTMX attributes  
✅ **Alpine.js Support** - Reactive UI with Alpine.js directives  
✅ **Responsive Design** - Mobile-first with responsive column layouts  
✅ **Internationalization** - RTL support and locale-aware formatting  
✅ **Advanced Features** - Auto-resize, copy buttons, password toggles, character counts  
✅ **Type Safety** - Full Go type safety with compile-time checks  
✅ **XSS Protection** - Sanitized Alpine.js expressions  

## Installation

```bash
go get github.com/niiniyare/ruun/views/components/molecules
```

## Quick Start

### Basic Text Input

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldText,
    ID:          "username",
    Name:        "username",
    Label:       "Username",
    Placeholder: "Enter your username",
    Required:    true,
    HelpText:    "Choose a unique username",
})
```

### Email with Validation

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldEmail,
    ID:          "email",
    Name:        "email",
    Label:       "Email Address",
    Required:    true,
    MaxLength:   100,
    ErrorText:   emailError, // Set if validation fails
    Autocomplete: "email",
})
```

### Password with Show/Hide Toggle

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:         molecules.FormFieldPassword,
    ID:           "password",
    Name:         "password",
    Label:        "Password",
    Required:     true,
    MinLength:    8,
    ShowPassword: true, // Enables toggle button
    HelpText:     "Must be at least 8 characters",
})
```

## Field Types

### Text Inputs

```go
// Standard text input
FormFieldText

// Email with validation
FormFieldEmail

// Password with optional visibility toggle
FormFieldPassword

// Number with step controls
FormFieldNumber

// URL with validation
FormFieldURL

// Telephone with formatting
FormFieldTel

// Search with clear button
FormFieldSearch
```

### Selection Fields

```go
// Radio button group
@molecules.FormField(molecules.FormFieldProps{
    Type:    molecules.FormFieldRadio,
    Name:    "plan",
    Label:   "Choose Plan",
    Inline:  true,
    Columns: 3,
    Options: []molecules.SelectOption{
        {Value: "basic", Label: "Basic", Description: "$9/month"},
        {Value: "pro", Label: "Pro", Description: "$29/month"},
        {Value: "enterprise", Label: "Enterprise", Description: "Custom"},
    },
})

// Checkbox group
@molecules.FormField(molecules.FormFieldProps{
    Type:    molecules.FormFieldCheckboxGroup,
    Name:    "features",
    Label:   "Select Features",
    Values:  []string{"api", "support"}, // Pre-selected values
    Options: []molecules.SelectOption{
        {Value: "api", Label: "API Access"},
        {Value: "support", Label: "Priority Support"},
        {Value: "analytics", Label: "Advanced Analytics"},
    },
})

// Multi-select dropdown
@molecules.FormField(molecules.FormFieldProps{
    Type:       molecules.FormFieldMultiSelect,
    Name:       "tags",
    Label:      "Tags",
    Searchable: true,
    Clearable:  true,
    Options:    tagOptions,
})

// Autocomplete with remote search
@molecules.FormField(molecules.FormFieldProps{
    Type:         molecules.FormFieldAutoComplete,
    Name:         "city",
    Label:        "City",
    SearchURL:    "/api/cities/search",
    MinChars:     3,
    MaxResults:   10,
    Debounce:     300,
    FreeForm:     true, // Allow non-matching values
})
```

### Date and Time

```go
// Date picker
@molecules.FormField(molecules.FormFieldProps{
    Type:         molecules.FormFieldDate,
    Name:         "birthdate",
    Label:        "Date of Birth",
    ShowCalendar: true,
    MinDate:      "1900-01-01",
    MaxDate:      "2010-12-31",
    DateFormat:   "MM/DD/YYYY",
})

// Time picker
@molecules.FormField(molecules.FormFieldProps{
    Type:     molecules.FormFieldTime,
    Name:     "meeting_time",
    Label:    "Meeting Time",
    Format24: false, // 12-hour format with AM/PM
})

// DateTime picker
@molecules.FormField(molecules.FormFieldProps{
    Type:  molecules.FormFieldDateTime,
    Name:  "appointment",
    Label: "Appointment",
})

// Date range picker
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldDateRange,
    Name:      "date_range",
    Label:     "Select Date Range",
    StartDate: "2024-01-01",
    EndDate:   "2024-12-31",
})
```

### Complex Inputs

```go
// Textarea with auto-resize
@molecules.FormField(molecules.FormFieldProps{
    Type:           molecules.FormFieldTextarea,
    Name:           "description",
    Label:          "Description",
    Rows:           4,
    AutoResize:     true,
    CharacterCount: true,
    MaxLength:      500,
})

// Editable tags
@molecules.FormField(molecules.FormFieldProps{
    Type:         molecules.FormFieldTags,
    Name:         "keywords",
    Label:        "Keywords",
    TagsEditable: true,
    MaxTags:      10,
})

// File upload with drag & drop
@molecules.FormField(molecules.FormFieldProps{
    Type:          molecules.FormFieldFile,
    Name:          "documents",
    Label:         "Upload Documents",
    Accept:        ".pdf,.doc,.docx",
    MaxFileSize:   10485760, // 10MB
    MultipleFiles: true,
    DropZone:      true,
    ShowPreview:   true,
})

// Range slider
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldRange,
    Name:      "budget",
    Label:     "Budget Range",
    Min:       "0",
    Max:       "10000",
    Step:      "100",
    ShowValue: true,
    ShowMinMax: true,
})

// Color picker
@molecules.FormField(molecules.FormFieldProps{
    Type:  molecules.FormFieldColor,
    Name:  "brand_color",
    Label: "Brand Color",
    Value: "#3B82F6",
})
```

## Advanced Features

### Server-Side Validation

```go
// Define validation rules
props := molecules.FormFieldProps{
    Type:  molecules.FormFieldText,
    Name:  "username",
    Label: "Username",
    ValidationRules: []molecules.ValidationRule{
        {
            Type:    "minLength",
            Value:   3,
            Message: "Username must be at least 3 characters",
        },
        {
            Type:    "pattern",
            Value:   "^[a-zA-Z0-9_]+$",
            Message: "Only letters, numbers, and underscores allowed",
        },
    },
}

// Validate input
if err := props.Validate(userInput); err != nil {
    props.ErrorText = err.Error()
}

// Render with error
@molecules.FormField(props)
```

### HTMX Integration

```go
// Live validation
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldEmail,
    Name:      "email",
    Label:     "Email",
    HXPost:    "/validate/email",
    HXTrigger: "blur changed delay:500ms",
    HXTarget:  "#email-validation",
    HXSwap:    "outerHTML",
})

// Dependent fields
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldSelect,
    Name:      "country",
    Label:     "Country",
    HXGet:     "/api/states",
    HXTarget:  "#state-field",
    HXTrigger: "change",
    Options:   countryOptions,
})

// Form submission with confirmation
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldText,
    Name:      "amount",
    HXPost:    "/transfer",
    HXConfirm: "Are you sure you want to transfer this amount?",
})
```

### Alpine.js Reactivity

```go
// Conditional display
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldText,
    Name:      "other_reason",
    Label:     "Please specify",
    Condition: "reason === 'other'", // Show only when condition is true
})

// Two-way binding
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldText,
    Name:        "price",
    Label:       "Price",
    AlpineModel: "price",
})

// Live calculation
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldNumber,
    Name:        "quantity",
    Label:       "Quantity",
    AlpineModel: "quantity",
    AlpineInput: "total = price * quantity",
})
```

### Responsive Layouts

```go
// Responsive column layout for radio/checkbox groups
@molecules.FormField(molecules.FormFieldProps{
    Type:      molecules.FormFieldCheckboxGroup,
    Name:      "interests",
    Label:     "Interests",
    Columns:   1,  // 1 column on mobile
    ColumnsMD: 2,  // 2 columns on tablet
    ColumnsLG: 3,  // 3 columns on desktop
    Options:   interestOptions,
})
```

### Custom Styling

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:         molecules.FormFieldText,
    Name:         "custom",
    Label:        "Custom Styled",
    Class:        "my-custom-wrapper",
    LabelClass:   "font-bold text-lg",
    InputClass:   "border-2 border-blue-500",
    WrapperClass: "bg-gray-50 p-4",
    Rounded:      "lg", // none, sm, md, lg, full
})
```

### Prefix and Suffix

```go
// With currency prefix
@molecules.FormField(molecules.FormFieldProps{
    Type:   molecules.FormFieldNumber,
    Name:   "price",
    Label:  "Price",
    Prefix: "$",
})

// With unit suffix
@molecules.FormField(molecules.FormFieldProps{
    Type:   molecules.FormFieldNumber,
    Name:   "weight",
    Label:  "Weight",
    Suffix: "kg",
})

// With icon
@molecules.FormField(molecules.FormFieldProps{
    Type:       molecules.FormFieldText,
    Name:       "search",
    PrefixIcon: "search",
    SuffixIcon: "filter",
})

// With copy button
@molecules.FormField(molecules.FormFieldProps{
    Type:       molecules.FormFieldText,
    Name:       "api_key",
    Label:      "API Key",
    Value:      apiKey,
    Readonly:   true,
    CopyButton: true,
})
```

### Loading States

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:    molecules.FormFieldText,
    Name:    "address",
    Label:   "Address",
    Loading: isGeocoding, // Show spinner
})
```

### Success and Warning States

```go
// Success state
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldEmail,
    Name:        "email",
    Value:       "user@example.com",
    SuccessText: "Email verified!",
})

// Warning state
@molecules.FormField(molecules.FormFieldProps{
    Type:        molecules.FormFieldPassword,
    Name:        "password",
    WarningText: "Password strength: weak",
})
```

## Accessibility Features

All form fields include:

- Proper ARIA labels and descriptions
- Keyboard navigation support
- Screen reader announcements for errors
- Focus management
- High contrast mode support
- Touch-friendly targets (minimum 44x44px)

```go
@molecules.FormField(molecules.FormFieldProps{
    Type:            molecules.FormFieldText,
    Name:            "custom_aria",
    AriaLabel:       "Custom accessible name",
    AriaDescribedBy: "custom-help-text",
    TabIndex:        1,
    Role:            "textbox",
})
```

## Internationalization

```go
// RTL support for Arabic, Hebrew, etc.
@molecules.FormField(molecules.FormFieldProps{
    Type:   molecules.FormFieldText,
    Name:   "name_ar",
    Label:  "الاسم",
    RTL:    true,
    Locale: "ar-SA",
})

// Date formatting based on locale
@molecules.FormField(molecules.FormFieldProps{
    Type:       molecules.FormFieldDate,
    Name:       "date",
    Locale:     "fr-FR",
    DateFormat: "DD/MM/YYYY",
})
```

## Performance Tips

1. **Use `GenerateID: true`** for auto-generated IDs to avoid manual ID management
2. **Debounce validation** with `DebounceMs` for expensive operations
3. **Lazy load options** for large select lists using HTMX
4. **Pre-select values** efficiently using `GetSelectedValue()` and `GetSelectedValues()`
5. **Batch validation** on form submit rather than field-by-field

## Security

### XSS Prevention

```go
// Alpine expressions are automatically sanitized
props := molecules.FormFieldProps{
    AlpineChange: userInput, // Automatically escaped
}

// Manual sanitization if needed
safeExpr, err := props.SanitizeAlpineExpression(untrustedInput)
```

### CSRF Protection

Add CSRF tokens at the form level:

```templ
<form method="post">
    <input type="hidden" name="csrf_token" value={ csrfToken }>
    @molecules.FormField(fieldProps)
</form>
```

## Testing

```go
func TestFormFieldValidation(t *testing.T) {
    props := molecules.FormFieldProps{
        Type:      molecules.FormFieldEmail,
        Required:  true,
        MaxLength: 100,
    }
    
    // Test required validation
    err := props.Validate("")
    assert.Error(t, err)
    
    // Test email format
    err = props.Validate("invalid-email")
    assert.Error(t, err)
    
    // Test valid email
    err = props.Validate("user@example.com")
    assert.NoError(t, err)
}
```

## Browser Support

- Chrome/Edge 90+
- Firefox 88+
- Safari 14+
- Mobile browsers (iOS Safari 14+, Chrome Mobile)

## Dependencies

- Go 1.21+
- Templ
- TailwindCSS 3.0+
- HTMX 1.9+ (optional)
- Alpine.js 3.0+ (optional)

## License

MIT

## Support

For issues and questions:
- GitHub Issues: https://github.com/niiniyare/ruun/issues
- Documentation: https://docs.ruun.dev
- Discord: https://discord.gg/ruun
