# Schema-Component Integration

**Comprehensive guide for integrating the schema system with atomic design components**

## Overview

This document details how the backend schema system seamlessly integrates with frontend atomic design components, enabling JSON-driven UI generation, dynamic behavior, and business logic coordination.

## Integration Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                    SCHEMA-COMPONENT FLOW                       │
├─────────────────────────────────────────────────────────────────┤
│  Backend Schema → Runtime Processing → Component Rendering     │
│                                                                 │
│  1. Schema Definition (Go)                                      │
│     └── JSON schema with fields, validation, business rules    │
│                                                                 │
│  2. Runtime Processing                                          │
│     ├── Schema validation and enrichment                       │
│     ├── Business rule evaluation                               │
│     ├── Conditional logic processing                           │
│     └── Theme and token application                            │
│                                                                 │
│  3. Component Rendering (Templ)                                │
│     ├── Schema-driven component selection                      │
│     ├── Props mapping from schema fields                       │
│     ├── Dynamic validation and state management                │
│     └── HTMX/Alpine.js integration                             │
└─────────────────────────────────────────────────────────────────┘
```

## Core Integration Patterns

### 1. Schema-Driven Component Selection

Components are automatically selected based on schema field definitions:

```go
// Schema field definition drives component selection
type FieldRenderer struct {
    registry map[schema.FieldType]ComponentRenderer
}

type ComponentRenderer interface {
    Render(field schema.Field, value interface{}, ctx RenderContext) templ.Component
    SupportedTypes() []schema.FieldType
    RequiredTokens() []string
}

// Field type to component mapping
func NewFieldRenderer() *FieldRenderer {
    return &FieldRenderer{
        registry: map[schema.FieldType]ComponentRenderer{
            schema.FieldTypeText:     &TextInputRenderer{},
            schema.FieldTypeEmail:    &EmailInputRenderer{},
            schema.FieldTypePassword: &PasswordInputRenderer{},
            schema.FieldTypeNumber:   &NumberInputRenderer{},
            schema.FieldTypeDate:     &DateInputRenderer{},
            schema.FieldTypeSelect:   &SelectRenderer{},
            schema.FieldTypeMultiSelect: &MultiSelectRenderer{},
            schema.FieldTypeTextarea: &TextareaRenderer{},
            schema.FieldTypeCheckbox: &CheckboxRenderer{},
            schema.FieldTypeRadio:    &RadioRenderer{},
            schema.FieldTypeFile:     &FileInputRenderer{},
            schema.FieldTypeCustom:   &CustomFieldRenderer{},
        },
    }
}

// Automatic component rendering based on schema
func (r *FieldRenderer) RenderField(field schema.Field, value interface{}, ctx RenderContext) templ.Component {
    renderer, exists := r.registry[field.Type]
    if !exists {
        return r.renderFallbackField(field, value, ctx)
    }
    
    return renderer.Render(field, value, ctx)
}
```

### 2. Schema Field to Component Props Mapping

Schema fields are automatically mapped to component props:

```go
// Text input renderer with schema integration
type TextInputRenderer struct{}

func (r *TextInputRenderer) Render(field schema.Field, value interface{}, ctx RenderContext) templ.Component {
    // Map schema field to component props
    props := FormFieldProps{
        // Basic field properties
        Name:        field.Name,
        Label:       field.Label,
        Type:        InputText,
        Value:       toString(value),
        Placeholder: field.Placeholder,
        Required:    field.Required,
        Disabled:    field.Disabled || ctx.ReadOnly,
        
        // Schema validation integration
        Error:       ctx.HasError(field.Name),
        ErrorText:   ctx.GetError(field.Name),
        HelpText:    field.HelpText,
        
        // Size and styling from schema
        Size:        mapSchemaSize(field.Size),
        Class:       field.CSSClass,
        
        // Business logic integration
        HXPost:      field.ValidationEndpoint,
        HXTrigger:   field.ValidationTrigger,
        HXTarget:    ctx.ValidationTarget,
        
        // Alpine.js state management
        AlpineModel: ctx.GetAlpineModel(field.Name),
        
        // Conditional display
        AlpineShow:  r.buildConditionalLogic(field.ConditionalLogic, ctx),
        
        // Accessibility from schema
        AriaLabel:       field.AriaLabel,
        AriaDescribedBy: field.AriaDescribedBy,
    }
    
    return FormField(props)
}

// Complex field renderer for business objects
type CustomerSelectRenderer struct{}

func (r *CustomerSelectRenderer) Render(field schema.Field, value interface{}, ctx RenderContext) templ.Component {
    // Extract business context from schema
    businessConfig := field.BusinessConfig.(*schema.CustomerSelectConfig)
    
    props := CustomerSelectorProps{
        Name:         field.Name,
        Label:        field.Label,
        Selected:     r.mapSelectedCustomer(value, businessConfig),
        Required:     field.Required,
        
        // Business logic from schema
        TenantFilter: businessConfig.TenantFilter,
        StatusFilter: businessConfig.AllowedStatuses,
        SearchFields: businessConfig.SearchFields,
        
        // Dynamic data loading
        SearchEndpoint: businessConfig.SearchEndpoint,
        DetailsEndpoint: businessConfig.DetailsEndpoint,
        
        // HTMX integration
        HXGet:       businessConfig.SearchEndpoint,
        HXTarget:    "#customer-options",
        HXTrigger:   "keyup changed delay:300ms",
        
        // Business validation
        ValidationRules: field.BusinessRules,
        
        // Permissions
        CreatePermission: businessConfig.CreatePermission,
        EditPermission:   businessConfig.EditPermission,
    }
    
    return CustomerSelector(props)
}
```

### 3. Dynamic Form Generation

Complete forms are generated from schema definitions:

```go
// Schema-driven form generator
type SchemaFormGenerator struct {
    fieldRenderer *FieldRenderer
    layoutEngine  *LayoutEngine
    validator     *SchemaValidator
}

func (g *SchemaFormGenerator) GenerateForm(schema *schema.Schema, data map[string]interface{}, ctx RenderContext) templ.Component {
    return SchemaForm(SchemaFormProps{
        Schema:     schema,
        Data:       data,
        Context:    ctx,
        Generator:  g,
    })
}

templ SchemaForm(props SchemaFormProps) {
    <form 
        class={ getFormClasses(props.Schema.Layout) }
        hx-post={ props.Schema.SubmitEndpoint }
        hx-target={ props.Context.FormTarget }
        x-data={ buildFormAlpineData(props.Schema, props.Data) }
        x-on:submit.prevent="submitForm()"
    >
        // Generate form layout based on schema
        if props.Schema.Layout != nil {
            @renderFormLayout(props.Schema.Layout, props)
        } else {
            @renderDefaultLayout(props)
        }
        
        // Form actions from schema
        @renderFormActions(props.Schema.Actions, props.Context)
    </form>
}

// Layout-aware field rendering
templ renderFormLayout(layout *schema.FormLayout, props SchemaFormProps) {
    switch layout.Type {
    case schema.LayoutTypeSingleColumn:
        <div class="space-y-[var(--space-6)]">
            for _, field := range props.Schema.Fields {
                @renderSchemaField(field, props)
            }
        </div>
    
    case schema.LayoutTypeMultiColumn:
        <div class={ getGridClasses(layout.Columns) }>
            for _, field := range props.Schema.Fields {
                <div class={ getFieldColumnClasses(field, layout) }>
                    @renderSchemaField(field, props)
                </div>
            }
        </div>
    
    case schema.LayoutTypeSections:
        for _, section := range layout.Sections {
            @renderFormSection(section, props)
        }
    
    case schema.LayoutTypeWizard:
        @renderWizardLayout(layout.WizardConfig, props)
    }
}

// Section-based form rendering
templ renderFormSection(section *schema.FormSection, props SchemaFormProps) {
    <fieldset class="space-y-[var(--space-4)]">
        if section.Title != "" {
            <legend class="text-lg font-semibold text-[var(--foreground)] mb-[var(--space-4)]">
                { section.Title }
            </legend>
        }
        
        if section.Description != "" {
            <p class="text-[var(--muted-foreground)] text-[var(--font-size-sm)] mb-[var(--space-6)]">
                { section.Description }
            </p>
        }
        
        <div class={ getSectionLayoutClasses(section.Layout) }>
            for _, fieldName := range section.Fields {
                if field := props.Schema.GetField(fieldName); field != nil {
                    @renderSchemaField(*field, props)
                }
            }
        </div>
    </fieldset>
}
```

### 4. Business Rules Integration

Schema business rules are integrated into component behavior:

```go
// Business rule engine integration
type BusinessRuleEngine struct {
    rules map[string]*schema.BusinessRule
    ctx   *RenderContext
}

func (e *BusinessRuleEngine) EvaluateRules(fieldName string, value interface{}) (*RuleEvaluation, error) {
    evaluation := &RuleEvaluation{
        FieldName: fieldName,
        Value:     value,
        Results:   make(map[string]*RuleResult),
    }
    
    for ruleName, rule := range e.rules {
        if rule.AppliesTo(fieldName) {
            result, err := e.evaluateRule(rule, value)
            if err != nil {
                return nil, err
            }
            evaluation.Results[ruleName] = result
        }
    }
    
    return evaluation, nil
}

// Component with business rules integration
templ BusinessRuleField(field schema.Field, value interface{}, ctx RenderContext) {
    <div 
        x-data={ fmt.Sprintf(`{
            value: %s,
            rules: %s,
            ruleResults: {},
            
            async validateBusinessRules() {
                if (!this.value) return;
                
                const response = await fetch('/api/validate/business-rules', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify({
                        fieldName: '%s',
                        value: this.value,
                        context: %s
                    })
                });
                
                const results = await response.json();
                this.ruleResults = results;
                
                // Update UI based on rule results
                this.updateFieldState(results);
            },
            
            updateFieldState(results) {
                // Update visibility
                if (results.visibility) {
                    this.$el.style.display = results.visibility.visible ? 'block' : 'none';
                }
                
                // Update validation
                if (results.validation) {
                    this.errors = results.validation.errors || [];
                }
                
                // Update options (for selects)
                if (results.options) {
                    this.availableOptions = results.options.allowed || [];
                }
            }
        }`, toJSON(value), toJSON(field.BusinessRules), field.Name, toJSON(ctx)) }
        x-on:change="validateBusinessRules()"
    >
        // Render field with business rule integration
        @renderFieldWithRules(field, value, ctx)
    </div>
}

// Dynamic field visibility based on business rules
templ ConditionalField(field schema.Field, value interface{}, ctx RenderContext) {
    <div 
        x-show={ buildConditionalExpression(field.ConditionalLogic) }
        x-transition:enter="transition ease-out duration-200"
        x-transition:enter-start="opacity-0 transform scale-95"
        x-transition:enter-end="opacity-100 transform scale-100"
    >
        @renderSchemaField(field, value, ctx)
    </div>
}

// Business rule expressions for Alpine.js
func buildConditionalExpression(logic *schema.ConditionalLogic) string {
    if logic == nil {
        return "true"
    }
    
    var conditions []string
    
    for _, condition := range logic.Conditions {
        expr := fmt.Sprintf("formData['%s'] %s '%s'", 
                           condition.Field, 
                           condition.Operator, 
                           condition.Value)
        conditions = append(conditions, expr)
    }
    
    operator := " && "
    if logic.Operator == "OR" {
        operator = " || "
    }
    
    return strings.Join(conditions, operator)
}
```

### 5. Real-Time Validation Integration

Schema validation rules are applied in real-time:

```go
// Real-time validation with schema rules
templ ValidatedFormField(field schema.Field, value interface{}, ctx RenderContext) {
    <div 
        x-data={ fmt.Sprintf(`{
            value: %s,
            errors: [],
            isValidating: false,
            validationRules: %s,
            
            async validateField() {
                if (this.isValidating) return;
                this.isValidating = true;
                this.errors = [];
                
                try {
                    // Client-side validation first
                    const clientErrors = this.validateClientRules();
                    if (clientErrors.length > 0) {
                        this.errors = clientErrors;
                        return;
                    }
                    
                    // Server-side validation
                    const response = await fetch('/api/validate/field', {
                        method: 'POST',
                        headers: {'Content-Type': 'application/json'},
                        body: JSON.stringify({
                            fieldName: '%s',
                            value: this.value,
                            rules: this.validationRules,
                            context: %s
                        })
                    });
                    
                    const result = await response.json();
                    this.errors = result.errors || [];
                    
                    // Update dependent fields if validation affects them
                    if (result.affects) {
                        this.$dispatch('field-validated', {
                            field: '%s',
                            value: this.value,
                            affects: result.affects
                        });
                    }
                    
                } catch (error) {
                    this.errors = ['Validation error occurred'];
                } finally {
                    this.isValidating = false;
                }
            },
            
            validateClientRules() {
                const errors = [];
                
                for (const rule of this.validationRules) {
                    if (rule.type === 'required' && !this.value) {
                        errors.push(rule.message || 'This field is required');
                    }
                    
                    if (rule.type === 'minLength' && this.value && this.value.length < rule.value) {
                        errors.push(rule.message || 'Minimum length is ' + rule.value);
                    }
                    
                    if (rule.type === 'maxLength' && this.value && this.value.length > rule.value) {
                        errors.push(rule.message || 'Maximum length is ' + rule.value);
                    }
                    
                    if (rule.type === 'pattern' && this.value && !new RegExp(rule.pattern).test(this.value)) {
                        errors.push(rule.message || 'Invalid format');
                    }
                    
                    if (rule.type === 'email' && this.value && !this.isValidEmail(this.value)) {
                        errors.push(rule.message || 'Please enter a valid email address');
                    }
                }
                
                return errors;
            },
            
            isValidEmail(email) {
                return /^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$/.test(email);
            }
        }`, toJSON(value), toJSON(field.Validation.Rules), field.Name, toJSON(ctx), field.Name) }
        x-on:blur="validateField()"
        x-on:change="validateField()"
    >
        @FormField(FormFieldProps{
            Name:        field.Name,
            Label:       field.Label,
            Type:        getInputType(field.Type),
            Value:       toString(value),
            Required:    field.Required,
            Error:       "errors.length > 0",
            ErrorText:   "errors[0]",
            AlpineModel: "value",
            Class:       "x-bind:class=\"isValidating ? 'opacity-50' : ''\"",
        })
        
        // Validation indicator
        <div x-show="isValidating" class="mt-1">
            @LoadingSpinner(IconProps{Size: IconSizeXS})
            <span class="text-xs text-[var(--muted-foreground)] ml-2">Validating...</span>
        </div>
    </div>
}
```

### 6. Multi-Tenant Schema Integration

Components adapt to tenant-specific schemas:

```go
// Tenant-aware schema rendering
type TenantSchemaRenderer struct {
    baseRenderer    *SchemaFormGenerator
    tenantManager   *TenantManager
    customizations  map[string]*TenantCustomization
}

func (r *TenantSchemaRenderer) RenderTenantForm(tenantID string, schema *schema.Schema, data map[string]interface{}) templ.Component {
    // Apply tenant customizations
    tenantSchema := r.applyTenantCustomizations(tenantID, schema)
    
    // Get tenant-specific theme
    theme := r.tenantManager.GetTenantTheme(tenantID)
    
    // Build tenant context
    ctx := RenderContext{
        TenantID:    tenantID,
        Theme:       theme,
        Permissions: r.tenantManager.GetTenantPermissions(tenantID),
        Locale:      r.tenantManager.GetTenantLocale(tenantID),
    }
    
    return TenantForm(TenantFormProps{
        TenantID:     tenantID,
        Schema:       tenantSchema,
        Data:         data,
        Context:     ctx,
        Theme:       theme,
        Customizations: r.customizations[tenantID],
    })
}

templ TenantForm(props TenantFormProps) {
    // Inject tenant-specific tokens
    <style>
        :root {
            for token, value := range props.Theme.Tokens {
                --{ token }: { value };
            }
        }
    </style>
    
    // Tenant-specific form wrapper
    <div class={ getTenantFormClasses(props.TenantID, props.Customizations) }>
        // Tenant branding
        if props.Customizations.ShowBranding {
            <div class="mb-[var(--space-6)]">
                <img src={ props.Theme.LogoURL } alt={ props.Theme.CompanyName } class="h-8" />
            </div>
        }
        
        // Standard schema form with tenant context
        @SchemaForm(SchemaFormProps{
            Schema:  props.Schema,
            Data:    props.Data,
            Context: props.Context,
        })
    </div>
}
```

### 7. Workflow Integration

Schema workflows are integrated with component state management:

```go
// Workflow-aware form rendering
templ WorkflowForm(schema *schema.Schema, workflow *schema.Workflow, currentState string, data map[string]interface{}) {
    <div 
        x-data={ fmt.Sprintf(`{
            currentState: '%s',
            workflow: %s,
            formData: %s,
            
            get allowedActions() {
                const state = this.workflow.states[this.currentState];
                return state ? state.allowedActions : [];
            },
            
            get requiredFields() {
                const state = this.workflow.states[this.currentState];
                return state ? state.requiredFields : [];
            },
            
            get readOnlyFields() {
                const state = this.workflow.states[this.currentState];
                return state ? state.readOnlyFields : [];
            },
            
            async transitionTo(newState, action) {
                // Validate transition
                if (!this.canTransition(newState)) {
                    alert('Cannot transition to ' + newState);
                    return;
                }
                
                // Perform workflow transition
                const response = await fetch('/api/workflow/transition', {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                    body: JSON.stringify({
                        fromState: this.currentState,
                        toState: newState,
                        action: action,
                        data: this.formData
                    })
                });
                
                const result = await response.json();
                if (result.success) {
                    this.currentState = newState;
                    // Refresh form with new state
                    htmx.ajax('GET', '/api/form/refresh', {
                        target: '#workflow-form',
                        values: {state: newState}
                    });
                }
            },
            
            canTransition(newState) {
                const currentStateConfig = this.workflow.states[this.currentState];
                return currentStateConfig && 
                       currentStateConfig.transitions.includes(newState);
            }
        }`, currentState, toJSON(workflow), toJSON(data)) }>
        
        // Workflow progress indicator
        @WorkflowProgress(WorkflowProgressProps{
            Workflow:     workflow,
            CurrentState: currentState,
        })
        
        // Dynamic form based on workflow state
        <form id="workflow-form" class="space-y-[var(--space-6)]">
            for _, field := range schema.Fields {
                <div 
                    x-show={ fmt.Sprintf("!readOnlyFields.includes('%s')", field.Name) }
                    x-bind:class={ fmt.Sprintf("requiredFields.includes('%s') ? 'required' : ''", field.Name) }
                >
                    @WorkflowField(WorkflowFieldProps{
                        Field:        field,
                        Value:        data[field.Name],
                        WorkflowState: currentState,
                        Required:     fmt.Sprintf("requiredFields.includes('%s')", field.Name),
                        ReadOnly:     fmt.Sprintf("readOnlyFields.includes('%s')", field.Name),
                    })
                </div>
            }
        </form>
        
        // Workflow actions
        <div class="flex justify-end gap-[var(--space-4)] pt-[var(--space-6)]">
            <template x-for="action in allowedActions" :key="action.id">
                @WorkflowActionButton(WorkflowActionButtonProps{
                    AlpineAction: "action",
                    AlpineClick:  "transitionTo(action.targetState, action)",
                })
            </template>
        </div>
    </div>
}
```

## Performance Optimizations

### Schema Caching

Implement efficient schema caching for performance:

```go
// Cached schema renderer
type CachedSchemaRenderer struct {
    cache       *lru.Cache
    ttl         time.Duration
    renderer    *SchemaFormGenerator
}

func (r *CachedSchemaRenderer) RenderForm(schemaID string, data map[string]interface{}) (templ.Component, error) {
    // Check cache first
    cacheKey := fmt.Sprintf("schema:%s:hash:%s", schemaID, hashData(data))
    if cached, ok := r.cache.Get(cacheKey); ok {
        return cached.(templ.Component), nil
    }
    
    // Load and render schema
    schema, err := r.loadSchema(schemaID)
    if err != nil {
        return nil, err
    }
    
    component := r.renderer.GenerateForm(schema, data, RenderContext{})
    
    // Cache result
    r.cache.Set(cacheKey, component, r.ttl)
    
    return component, nil
}
```

### Lazy Loading

Implement lazy loading for complex components:

```go
// Lazy-loaded schema sections
templ LazySchemaSection(sectionID string, schema *schema.Schema) {
    <div 
        x-data="{ loaded: false }"
        x-intersect="if (!loaded) { loaded = true; loadSection() }"
    >
        <div x-show="!loaded">
            @SectionSkeleton()
        </div>
        
        <div 
            x-show="loaded"
            hx-get={ fmt.Sprintf("/api/schema/section/%s", sectionID) }
            hx-trigger="load"
            x-transition
        >
            <!-- Section content loaded via HTMX -->
        </div>
    </div>
}
```

## Testing Schema Integration

```go
func TestSchemaFormGeneration(t *testing.T) {
    schema := &schema.Schema{
        ID:    "test-form",
        Title: "Test Form",
        Fields: []schema.Field{
            {
                Name:     "firstName",
                Type:     schema.FieldTypeText,
                Label:    "First Name",
                Required: true,
            },
            {
                Name:     "email",
                Type:     schema.FieldTypeEmail,
                Label:    "Email",
                Required: true,
            },
        },
    }
    
    data := map[string]interface{}{
        "firstName": "John",
        "email":     "john@example.com",
    }
    
    generator := NewSchemaFormGenerator()
    component := generator.GenerateForm(schema, data, RenderContext{})
    html := renderToString(component)
    
    // Test schema-driven rendering
    assert.Contains(t, html, "First Name")
    assert.Contains(t, html, "Email")
    assert.Contains(t, html, "required")
    assert.Contains(t, html, "john@example.com")
    
    // Test Alpine.js integration
    assert.Contains(t, html, "x-data")
    assert.Contains(t, html, "formData")
    
    // Test HTMX integration
    assert.Contains(t, html, "hx-post")
}
```

## Best Practices

### ✅ DO

- **Start with schema definitions** for all form and component integrations
- **Use type-safe mapping** between schema fields and component props  
- **Implement comprehensive validation** at both client and server levels
- **Cache schema processing results** for performance
- **Test schema-component integration** thoroughly
- **Handle error states gracefully** in schema-driven components
- **Document schema-component mappings** clearly

### ❌ DON'T

- **Skip schema validation** when processing schema-driven components
- **Ignore performance implications** of dynamic schema processing
- **Create tight coupling** between specific schemas and components
- **Forget accessibility** in schema-driven components
- **Skip error handling** for invalid schema definitions

## Related Documentation

- **[Schema Package Usage](../usage.md)** - Core schema system usage
- **[Business Rules](../schema/business_rules.md)** - Business logic integration
- **[Component Development](./08-development.md)** - Component development workflow
- **[Design Tokens](./01-design-tokens.md)** - Token system integration

---

The schema-component integration enables powerful, dynamic, and maintainable user interfaces driven entirely by backend definitions while maintaining the flexibility and composability of atomic design components.