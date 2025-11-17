# Component Development Workflow

**Complete guide for developing, testing, and maintaining atomic design components in the schema-driven system**

## Overview

This guide outlines the end-to-end workflow for developing components within our atomic design system, from initial design through deployment and maintenance. It covers tooling, testing strategies, integration patterns, and best practices.

## Development Environment Setup

### Prerequisites

```bash
# Required tools
go version  # Go 1.21+
templ version  # Templ v0.2.543+
npm --version  # Node.js 18+
air --version  # Live reload for Go development

# Project dependencies
go mod download
npm install

# Schema validation tools
go install ./cmd/schema-validator
go install ./cmd/component-analyzer
```

### Project Structure

```
pkg/schema/docs/components/
├── 01-design-tokens.md     # Token system documentation
├── 02-atoms.md            # Basic building blocks
├── 03-molecules.md        # Composite components
├── 04-organisms.md        # Complex components
├── 05-templates.md        # Page layout structures
├── 06-pages.md           # Complete implementations
├── 07-integration.md     # Schema integration patterns
└── 08-development.md     # This workflow guide

views/components/
├── atoms/
│   ├── button/
│   │   ├── button.templ          # Component implementation
│   │   ├── button_test.go        # Unit tests
│   │   ├── button.stories.go     # Storybook stories
│   │   └── button.schema.json    # Schema definition
│   └── ...
├── molecules/
├── organisms/
├── templates/
└── pages/

tools/
├── component-generator/     # Scaffolding tools
├── schema-validator/       # Schema validation
├── visual-regression/      # Visual testing
└── performance-analyzer/   # Performance tools
```

## Component Development Lifecycle

### 1. Planning and Design

#### Schema Definition First

Every component starts with a schema definition that describes its structure, validation rules, and business context.

```json
{
  "id": "customer-form",
  "type": "organism",
  "title": "Customer Form",
  "description": "Customer creation and editing form with validation",
  "category": "forms",
  "atomicLevel": "organism",
  "dependencies": [
    "form-field",
    "customer-selector",
    "address-input"
  ],
  "schema": {
    "fields": [
      {
        "name": "firstName",
        "type": "text",
        "required": true,
        "validation": {
          "minLength": 2,
          "maxLength": 50,
          "pattern": "^[a-zA-Z\\s]+$"
        }
      },
      {
        "name": "email",
        "type": "email",
        "required": true,
        "validation": {
          "format": "email",
          "unique": true
        }
      }
    ],
    "businessRules": [
      {
        "name": "validateCustomerUniqueness",
        "condition": "email.changed",
        "action": "validateUnique"
      }
    ]
  },
  "props": {
    "customer": {
      "type": "Customer",
      "required": false,
      "description": "Existing customer data for editing"
    },
    "mode": {
      "type": "FormMode",
      "required": true,
      "enum": ["create", "edit", "view"],
      "default": "create"
    },
    "onSubmit": {
      "type": "string",
      "description": "HTMX endpoint for form submission"
    }
  },
  "tokens": {
    "spacing": ["space-4", "space-6"],
    "colors": ["primary", "destructive", "muted"],
    "typography": ["font-size-sm", "font-size-base"]
  }
}
```

#### Design Token Planning

Map component styling needs to design tokens:

```go
// Token usage planning for new component
type ComponentTokenUsage struct {
    Colors []string `json:"colors"`
    Spacing []string `json:"spacing"`
    Typography []string `json:"typography"`
    BorderRadius []string `json:"borderRadius"`
    Shadows []string `json:"shadows"`
}

// Example: Planning tokens for a new card component
cardTokens := ComponentTokenUsage{
    Colors: []string{
        "card",           // Background
        "card-foreground", // Text color
        "border",         // Border color
        "muted",          // Secondary background
    },
    Spacing: []string{
        "space-4",        // Internal padding
        "space-6",        // Large padding
        "space-2",        // Small gaps
    },
    BorderRadius: []string{
        "radius-lg",      // Card border radius
    },
    Shadows: []string{
        "shadow-sm",      // Subtle shadow
        "shadow-md",      // Elevated state
    },
}
```

### 2. Component Scaffolding

#### Automated Generation

Use the component generator to create initial structure:

```bash
# Generate new atom component
go run ./tools/component-generator atom button \
    --props "variant,size,disabled,loading" \
    --tokens "primary,secondary,spacing" \
    --htmx-enabled \
    --alpine-enabled

# Generate molecule component
go run ./tools/component-generator molecule form-field \
    --dependencies "button,input,label" \
    --schema-driven \
    --validation-enabled

# Generate organism component  
go run ./tools/component-generator organism data-table \
    --dependencies "button,input,pagination" \
    --business-logic \
    --state-management
```

#### Generated Structure

```
views/components/atoms/button/
├── button.templ              # Main component
├── button.go                 # Go types and helpers
├── button_test.go            # Unit tests
├── button.stories.go         # Development stories
├── button.schema.json        # Schema definition
└── README.md                # Component documentation
```

**Generated button.templ**:

```go
// Auto-generated component structure
package button

import "github.com/a-h/templ"

type ButtonVariant string
type ButtonSize string

const (
    ButtonPrimary   ButtonVariant = "primary"
    ButtonSecondary ButtonVariant = "secondary"
    // ... other variants
)

type ButtonProps struct {
    Variant  ButtonVariant
    Size     ButtonSize
    Disabled bool
    Loading  bool
    Class    string
    
    // HTMX integration
    HXPost   string
    HXTarget string
    
    // Alpine.js integration
    AlpineClick string
    
    // Accessibility
    AriaLabel string
    AriaDescribedBy string
}

func buttonClasses(props ButtonProps) string {
    var classes []string
    
    // Base classes using design tokens
    classes = append(classes,
        "inline-flex",
        "items-center",
        "justify-center",
        "rounded-[var(--radius-md)]",
        "text-[var(--font-size-sm)]",
        "font-medium",
        "transition-colors",
        "focus-visible:outline-none",
        "focus-visible:ring-2",
        "focus-visible:ring-[var(--ring)]",
        "disabled:pointer-events-none",
        "disabled:opacity-50",
    )
    
    // Variant-specific styling
    switch props.Variant {
    case ButtonPrimary:
        classes = append(classes,
            "bg-[var(--primary)]",
            "text-[var(--primary-foreground)]",
            "hover:bg-[var(--primary)]/90",
        )
    case ButtonSecondary:
        classes = append(classes,
            "bg-[var(--secondary)]",
            "text-[var(--secondary-foreground)]",
            "hover:bg-[var(--secondary)]/80",
        )
    }
    
    // Size-specific styling
    switch props.Size {
    case ButtonSizeSM:
        classes = append(classes, "h-[var(--space-8)]", "px-[var(--space-3)]")
    case ButtonSizeLG:
        classes = append(classes, "h-[var(--space-11)]", "px-[var(--space-8)]")
    default: // MD
        classes = append(classes, "h-[var(--space-10)]", "px-[var(--space-4)]")
    }
    
    if props.Class != "" {
        classes = append(classes, props.Class)
    }
    
    return strings.Join(classes, " ")
}

templ Button(props ButtonProps, children ...templ.Component) {
    <button
        type="button"
        class={ buttonClasses(props) }
        if props.Disabled {
            disabled
        }
        if props.HXPost != "" {
            hx-post={ props.HXPost }
        }
        if props.HXTarget != "" {
            hx-target={ props.HXTarget }
        }
        if props.AlpineClick != "" {
            x-on:click={ props.AlpineClick }
        }
        if props.AriaLabel != "" {
            aria-label={ props.AriaLabel }
        }
        if props.AriaDescribedBy != "" {
            aria-describedby={ props.AriaDescribedBy }
        }
    >
        if props.Loading {
            @LoadingSpinner(IconProps{Size: IconSizeSM})
        }
        
        for _, child := range children {
            @child
        }
    </button>
}
```

### 3. Development and Testing

#### Component Development

Follow the atomic design principles during development:

```go
// 1. Start with atoms - ensure they're complete and well-tested
func TestButtonAtom_AllVariants(t *testing.T) {
    variants := []ButtonVariant{
        ButtonPrimary, ButtonSecondary, ButtonDestructive,
        ButtonOutline, ButtonGhost, ButtonLink,
    }
    
    for _, variant := range variants {
        t.Run(string(variant), func(t *testing.T) {
            props := ButtonProps{Variant: variant}
            html := renderToString(Button(props))
            
            assert.Contains(t, html, "<button")
            assert.Contains(t, html, getExpectedVariantClass(variant))
        })
    }
}

// 2. Build molecules using tested atoms
templ FormField(props FormFieldProps) {
    <div class="space-y-[var(--space-2)]">
        @Label(LabelProps{
            For:      props.Name,
            Required: props.Required,
        }) {
            { props.Label }
        }
        
        @Input(InputProps{
            Name:        props.Name,
            Type:        props.Type,
            Value:       props.Value,
            Placeholder: props.Placeholder,
            Required:    props.Required,
            Error:       props.Error,
        })
        
        if props.Error && props.ErrorText != "" {
            <p class="text-[var(--destructive)] text-xs">
                { props.ErrorText }
            </p>
        }
    </div>
}

// 3. Create organisms with business logic
templ CustomerForm(props CustomerFormProps) {
    <form 
        hx-post={ props.SubmitEndpoint }
        hx-target="#form-result"
        x-data={ fmt.Sprintf(`{
            customer: %s,
            errors: {},
            isSubmitting: false,
            
            validateForm() {
                this.errors = {};
                
                if (!this.customer.firstName) {
                    this.errors.firstName = ['First name is required'];
                }
                
                if (!this.customer.email) {
                    this.errors.email = ['Email is required'];
                } else if (!this.isValidEmail(this.customer.email)) {
                    this.errors.email = ['Please enter a valid email'];
                }
                
                return Object.keys(this.errors).length === 0;
            },
            
            submitForm() {
                if (this.validateForm()) {
                    this.isSubmitting = true;
                    // Form will be submitted via HTMX
                }
            }
        }`, toJSON(props.Customer)) }
        x-on:submit.prevent="submitForm()"
    >
        @FormField(FormFieldProps{
            Name:        "firstName",
            Label:       "First Name",
            Type:        InputText,
            Value:       props.Customer.FirstName,
            Required:    true,
            Error:       "errors.firstName",
            AlpineModel: "customer.firstName",
        })
        
        @FormField(FormFieldProps{
            Name:        "email",
            Label:       "Email Address",
            Type:        InputEmail,
            Value:       props.Customer.Email,
            Required:    true,
            Error:       "errors.email",
            AlpineModel: "customer.email",
        })
        
        <div class="flex justify-end gap-[var(--space-4)]">
            @Button(ButtonProps{
                Variant: ButtonOutline,
                Type:    "button",
            }) {
                Cancel
            }
            
            @Button(ButtonProps{
                Variant:     ButtonPrimary,
                Type:        "submit",
                Loading:     true,
                AlpineClick: "submitForm()",
                Class:       "x-bind:disabled='isSubmitting'",
            }) {
                { if props.Mode == "edit" { "Update" } else { "Create" } } Customer
            }
        </div>
    </form>
}
```

#### Live Development

Set up live reload for rapid development:

```bash
# Terminal 1: Start Go server with live reload
air

# Terminal 2: Start Templ watcher
templ generate --watch

# Terminal 3: Start frontend asset watcher
npm run dev

# Terminal 4: Run component development server
go run ./tools/component-dev-server --port 8080
```

#### Component Stories

Create development stories for visual testing:

```go
// button.stories.go
package button

import "github.com/a-h/templ"

type Story struct {
    Name        string
    Description string
    Component   templ.Component
    Code        string
}

func ButtonStories() []Story {
    return []Story{
        {
            Name:        "Primary Button",
            Description: "Primary button for main actions",
            Component: Button(ButtonProps{
                Variant: ButtonPrimary,
                Size:    ButtonSizeMD,
            }) {
                templ.Raw("Click me")
            },
            Code: `@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size:    ButtonSizeMD,
}) {
    Click me
}`,
        },
        {
            Name:        "Loading Button",
            Description: "Button in loading state",
            Component: Button(ButtonProps{
                Variant: ButtonPrimary,
                Size:    ButtonSizeMD,
                Loading: true,
            }) {
                templ.Raw("Loading...")
            },
        },
        {
            Name:        "Disabled Button", 
            Description: "Disabled button state",
            Component: Button(ButtonProps{
                Variant:  ButtonPrimary,
                Size:     ButtonSizeMD,
                Disabled: true,
            }) {
                templ.Raw("Disabled")
            },
        },
        // Interactive examples
        {
            Name:        "HTMX Integration",
            Description: "Button with HTMX behavior",
            Component: Button(ButtonProps{
                Variant:  ButtonPrimary,
                HXPost:   "/api/test",
                HXTarget: "#result",
            }) {
                templ.Raw("Submit")
            },
        },
        {
            Name:        "Alpine.js Integration",
            Description: "Button with Alpine.js interaction",
            Component: Button(ButtonProps{
                Variant:     ButtonSecondary,
                AlpineClick: "alert('Hello from Alpine!')",
            }) {
                templ.Raw("Alert")
            },
        },
    }
}
```

### 4. Testing Strategy

#### Unit Testing

Test component logic and rendering:

```go
func TestButton_Rendering(t *testing.T) {
    tests := []struct {
        name     string
        props    ButtonProps
        children []templ.Component
        want     []string
        notWant  []string
    }{
        {
            name: "primary button",
            props: ButtonProps{
                Variant: ButtonPrimary,
                Size:    ButtonSizeMD,
            },
            children: []templ.Component{templ.Raw("Click me")},
            want: []string{
                `<button`,
                `type="button"`,
                `bg-[var(--primary)]`,
                `text-[var(--primary-foreground)]`,
                `h-[var(--space-10)]`,
                `Click me`,
            },
            notWant: []string{
                `disabled`,
                `hx-post`,
                `loading`,
            },
        },
        {
            name: "disabled button",
            props: ButtonProps{
                Variant:  ButtonPrimary,
                Disabled: true,
            },
            want: []string{
                `disabled`,
                `disabled:pointer-events-none`,
                `disabled:opacity-50`,
            },
        },
        {
            name: "HTMX enabled button",
            props: ButtonProps{
                Variant:  ButtonPrimary,
                HXPost:   "/api/submit",
                HXTarget: "#result",
            },
            want: []string{
                `hx-post="/api/submit"`,
                `hx-target="#result"`,
            },
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            html := renderToString(Button(tt.props, tt.children...))
            
            for _, want := range tt.want {
                assert.Contains(t, html, want)
            }
            
            for _, notWant := range tt.notWant {
                assert.NotContains(t, html, notWant)
            }
        })
    }
}
```

#### Integration Testing

Test component integration with HTMX and Alpine.js:

```go
func TestFormField_Integration(t *testing.T) {
    props := FormFieldProps{
        Name:        "email",
        Label:       "Email Address",
        Type:        InputEmail,
        Required:    true,
        Error:       true,
        ErrorText:   "Please enter a valid email",
        AlpineModel: "user.email",
    }
    
    html := renderToString(FormField(props))
    
    // Test Alpine.js integration
    assert.Contains(t, html, `x-model="user.email"`)
    
    // Test error state
    assert.Contains(t, html, "Please enter a valid email")
    assert.Contains(t, html, "text-[var(--destructive)]")
    
    // Test accessibility
    assert.Contains(t, html, `for="email"`)
    assert.Contains(t, html, `id="email"`)
    assert.Contains(t, html, `required`)
}
```

#### Visual Regression Testing

Automate visual testing with screenshots:

```bash
# Run visual regression tests
go run ./tools/visual-regression \
    --components "button,input,form-field" \
    --browsers "chrome,firefox,safari" \
    --viewports "mobile,tablet,desktop"
```

#### Performance Testing

Test component performance and bundle size:

```go
func BenchmarkButton_Rendering(b *testing.B) {
    props := ButtonProps{
        Variant: ButtonPrimary,
        Size:    ButtonSizeMD,
    }
    children := []templ.Component{templ.Raw("Click me")}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = renderToString(Button(props, children...))
    }
}

func TestButton_BundleSize(t *testing.T) {
    // Test that component CSS doesn't exceed size limits
    css := generateComponentCSS("button")
    assert.Less(t, len(css), 5000, "Button CSS should be under 5KB")
}
```

### 5. Schema Integration

#### Schema-Driven Components

Components that integrate with the schema system:

```go
// Schema-driven form field that adapts based on field definition
templ SchemaFormField(field schema.Field, value interface{}) {
    switch field.Type {
    case schema.FieldTypeText:
        @FormField(FormFieldProps{
            Name:        field.Name,
            Label:       field.Label,
            Type:        InputText,
            Value:       toString(value),
            Required:    field.Required,
            Placeholder: field.Placeholder,
            HelpText:    field.HelpText,
        })
    
    case schema.FieldTypeEmail:
        @FormField(FormFieldProps{
            Name:        field.Name,
            Label:       field.Label,
            Type:        InputEmail,
            Value:       toString(value),
            Required:    field.Required,
            Validation:  getEmailValidation(field),
        })
    
    case schema.FieldTypeSelect:
        @SelectField(SelectFieldProps{
            Name:     field.Name,
            Label:    field.Label,
            Options:  field.Options,
            Selected: toString(value),
            Required: field.Required,
        })
    
    case schema.FieldTypeCustom:
        // Render custom field based on field.Component
        @renderCustomComponent(field, value)
    }
}

// Component that renders entire forms from schema
templ SchemaForm(schema *schema.Schema, data map[string]interface{}) {
    <form 
        hx-post={ schema.SubmitEndpoint }
        hx-target="#form-result"
        x-data={ fmt.Sprintf(`{
            data: %s,
            errors: {},
            
            validateField(fieldName, rules) {
                const value = this.data[fieldName];
                const errors = [];
                
                // Apply schema validation rules
                for (const rule of rules) {
                    if (!validateRule(value, rule)) {
                        errors.push(rule.message);
                    }
                }
                
                this.errors[fieldName] = errors;
                return errors.length === 0;
            }
        }`, toJSON(data)) }
    >
        // Render fields based on schema layout
        if schema.Layout != nil {
            @renderSchemaLayout(schema.Layout, schema.Fields)
        } else {
            // Default single-column layout
            for _, field := range schema.Fields {
                @SchemaFormField(field, data[field.Name])
            }
        }
        
        // Form actions from schema
        if len(schema.Actions) > 0 {
            <div class="flex justify-end gap-[var(--space-4)] pt-[var(--space-6)]">
                for _, action := range schema.Actions {
                    @schemaActionButton(action)
                }
            </div>
        }
    </form>
}
```

#### Runtime Schema Updates

Handle dynamic schema changes:

```go
// Component that updates based on runtime schema changes
templ DynamicForm(props DynamicFormProps) {
    <div 
        x-data={ fmt.Sprintf(`{
            schema: %s,
            data: %s,
            
            async updateSchema(trigger) {
                // Fetch updated schema based on trigger
                const response = await fetch('/api/schema/dynamic', {
                    method: 'POST',
                    body: JSON.stringify({
                        trigger: trigger,
                        currentData: this.data
                    })
                });
                
                const newSchema = await response.json();
                this.schema = newSchema;
                
                // Re-render form with updated schema
                this.$dispatch('schema-updated', newSchema);
            }
        }`, toJSON(props.Schema), toJSON(props.Data)) }
        x-on:field-changed="updateSchema($event.detail)"
    >
        // Dynamically rendered form
        <template x-for="field in schema.fields" :key="field.name">
            <div x-html="renderSchemaField(field, data[field.name])"></div>
        </template>
    </div>
}
```

### 6. Documentation and Maintenance

#### Component Documentation

Generate comprehensive documentation:

```bash
# Generate component documentation
go run ./tools/component-docs-generator \
    --input "./views/components" \
    --output "./docs/components" \
    --format "markdown"
```

**Generated documentation structure**:

```markdown
# Button Component

## Overview
The Button component provides consistent interactive elements across the application.

## Props
| Prop | Type | Required | Default | Description |
|------|------|----------|---------|-------------|
| variant | ButtonVariant | No | primary | Visual style variant |
| size | ButtonSize | No | md | Button size |
| disabled | bool | No | false | Disable button interaction |

## Examples
### Basic Usage
```go
@Button(ButtonProps{
    Variant: ButtonPrimary,
    Size:    ButtonSizeMD,
}) {
    Click me
}
```

### With HTMX
```go
@Button(ButtonProps{
    Variant:  ButtonPrimary,
    HXPost:   "/api/submit",
    HXTarget: "#result",
}) {
    Submit
}
```

## Design Tokens
- Colors: `primary`, `primary-foreground`, `secondary`, `destructive`
- Spacing: `space-4`, `space-8`, `space-10`
- Border Radius: `radius-md`

## Accessibility
- Supports ARIA labels and descriptions
- Keyboard navigation compatible
- Screen reader friendly
- Focus visible indicators

## Testing
Run component tests: `go test ./views/components/atoms/button`
```

#### Changelog and Versioning

Track component changes:

```markdown
# Button Component Changelog

## v2.1.0 - 2024-01-15
### Added
- Loading state with spinner
- Alpine.js click handler support
- New `ghost` variant

### Changed
- Improved focus visible styles
- Updated hover animations

### Fixed
- Disabled state accessibility
- Mobile touch target size

## v2.0.0 - 2024-01-01
### Breaking Changes
- Renamed `type` prop to `variant`
- Removed deprecated `outline` variant
- Changed default size from `sm` to `md`
```

#### Automated Maintenance

Set up automated maintenance tasks:

```bash
# Update component dependencies
go run ./tools/component-updater --check-dependencies

# Validate schema compliance
go run ./tools/schema-validator --components

# Check for unused design tokens
go run ./tools/token-analyzer --unused

# Performance audit
go run ./tools/performance-audit --components
```

## Best Practices

### ✅ DO

- **Start with schema definitions** before writing component code
- **Use design tokens consistently** throughout all styling
- **Write comprehensive tests** for all component variants and states
- **Document component APIs** with clear examples and use cases
- **Follow atomic design principles** strictly for component composition
- **Implement accessibility** from the beginning, not as an afterthought
- **Test with real data** and realistic usage scenarios
- **Monitor performance** and bundle size impacts
- **Version components** and maintain changelogs
- **Use automated tools** for code generation and validation

### ❌ DON'T

- **Skip schema definitions** - they guide the entire development process
- **Hardcode styles** - always use design tokens for consistency
- **Create overly complex components** - follow single responsibility principle
- **Ignore accessibility requirements** - build inclusive components
- **Skip testing edge cases** - components should handle all scenarios gracefully
- **Break atomic design hierarchy** - molecules shouldn't contain organisms
- **Forget responsive behavior** - design mobile-first
- **Ignore performance implications** - components should be efficient
- **Create components without clear use cases** - build for actual needs
- **Skip documentation** - undocumented components become technical debt

## Tools and Resources

### Command Line Tools

```bash
# Component scaffolding
component-generator [type] [name] [options]

# Schema validation
schema-validator --components [path]

# Visual regression testing
visual-regression --components [list]

# Performance analysis
performance-audit --components [list]

# Documentation generation
component-docs-generator --input [path] --output [path]
```

### IDE Integration

- **VS Code Extension**: Templ syntax highlighting and IntelliSense
- **Go Language Server**: Full Go support with type checking
- **Tailwind CSS IntelliSense**: Design token autocomplete
- **HTMX Support**: HTMX attribute validation and snippets

### Useful Resources

- **[Atomic Design Methodology](https://atomicdesign.bradfrost.com/)** - Original atomic design principles
- **[Templ Documentation](https://templ.guide/)** - Templ templating language guide
- **[HTMX Documentation](https://htmx.org/docs/)** - HTMX integration patterns
- **[Alpine.js Documentation](https://alpinejs.dev/)** - Alpine.js reactive patterns

---

This development workflow ensures consistent, maintainable, and high-quality components that integrate seamlessly with the schema-driven architecture while following atomic design principles.