# Conditional Logic

## Show/Hide Fields

### Simple Condition

```json
{
  "name": "shipping_address",
  "type": "text",
  "label": "Shipping Address",
  "showIf": "needs_shipping === true"
}
```

Uses Alpine.js on client:
```html
<div x-show="needs_shipping === true">
  <input name="shipping_address" />
</div>
```

### Complex Condition

```json
{
  "showIf": "order_total > 100 && country === 'US'"
}
```

## Permission-Based Visibility

```json
{
  "name": "salary",
  "type": "currency",
  "label": "Salary",
  "requirePermission": "hr.view_salary"
}
```

Server sets runtime visibility:
```go
field.Runtime = &FieldRuntime{
    Visible: user.HasPermission("hr.view_salary"),
}
```

## Conditional Required

```json
{
  "name": "tax_id",
  "type": "text",
  "label": "Tax ID",
  "requiredIf": "is_business === true"
}
```

## Example: Shipping Form

```json
{
  "fields": [
    {
      "name": "needs_shipping",
      "type": "checkbox",
      "label": "Ship to different address"
    },
    {
      "name": "shipping_street",
      "type": "text",
      "label": "Street",
      "showIf": "needs_shipping === true",
      "requiredIf": "needs_shipping === true"
    },
    {
      "name": "shipping_city",
      "type": "text",
      "label": "City",
      "showIf": "needs_shipping === true",
      "requiredIf": "needs_shipping === true"
    }
  ]
}
```

## Advanced Implementation

### Business Rules Engine

The system includes a sophisticated business rules engine for complex conditional logic:

```go
// BusinessRuleEngine manages and executes business rules with thread safety
type BusinessRuleEngine struct {
    rules     map[string]*BusinessRule
    evaluator *condition.Evaluator
    mu        sync.RWMutex              // Thread-safe rule management
}

// BusinessRule represents a business logic rule
type BusinessRule struct {
    ID          string                    `json:"id"`
    Name        string                    `json:"name"`
    Description string                    `json:"description"`
    Type        BusinessRuleType          `json:"type"`
    Priority    int                       `json:"priority"`     // Higher = executed first
    Enabled     bool                      `json:"enabled"`
    Condition   *condition.ConditionGroup `json:"condition"`    // When to apply
    Actions     []BusinessRuleAction      `json:"actions"`      // What to do
    Metadata    map[string]any            `json:"metadata"`
    CreatedAt   time.Time                 `json:"created_at"`
    UpdatedAt   time.Time                 `json:"updated_at"`
}
```

### Rule Types & Actions

```go
// Field-level rules
const (
    RuleTypeFieldVisibility BusinessRuleType = "field_visibility"
    RuleTypeFieldRequired   BusinessRuleType = "field_required"
    RuleTypeFieldDefault    BusinessRuleType = "field_default"
    RuleTypeFieldValidation BusinessRuleType = "field_validation"
    RuleTypeFieldOptions    BusinessRuleType = "field_options"
)

// Action types
const (
    ActionShowField     BusinessRuleActionType = "show_field"
    ActionHideField     BusinessRuleActionType = "hide_field"
    ActionRequireField  BusinessRuleActionType = "require_field"
    ActionOptionalField BusinessRuleActionType = "optional_field"
    ActionSetDefault    BusinessRuleActionType = "set_default"
    ActionCalculate     BusinessRuleActionType = "calculate"
)

// BusinessRuleAction defines what happens when condition is met
type BusinessRuleAction struct {
    Type   BusinessRuleActionType `json:"type"`
    Target string                 `json:"target"`    // Field name
    Value  any                    `json:"value"`     // Value to set
    Params map[string]any         `json:"params"`    // Additional parameters
}
```

### Priority-Based Rule Execution

```go
// Rules are applied in priority order (higher priority first)
func (bre *BusinessRuleEngine) ApplyRules(
    ctx context.Context, 
    schema *Schema, 
    data map[string]any,
) (*Schema, error) {
    bre.mu.RLock()
    defer bre.mu.RUnlock()
    
    // Create deep copy to modify
    modifiedSchema := bre.deepCopySchema(schema)
    
    // Get rules sorted by priority
    applicableRules := bre.getApplicableRules()
    collector := NewErrorCollector()
    
    // Apply each rule in order
    for _, rule := range applicableRules {
        if !rule.Enabled {
            continue
        }
        
        // Evaluate condition
        shouldApply, err := bre.evaluateRuleCondition(ctx, rule, data)
        if err != nil {
            collector.AddValidationError(rule.ID, "condition_failed", err.Error())
            continue
        }
        
        if shouldApply {
            if err := bre.applyRuleActions(ctx, rule, modifiedSchema, data); err != nil {
                collector.AddValidationError(rule.ID, "action_failed", err.Error())
            }
        }
    }
    
    if collector.HasErrors() {
        return schema, collector.Errors()
    }
    
    return modifiedSchema, nil
}

// Sort rules by priority (higher first)
func (bre *BusinessRuleEngine) getApplicableRules() []*BusinessRule {
    rules := make([]*BusinessRule, 0, len(bre.rules))
    for _, rule := range bre.rules {
        rules = append(rules, rule)
    }
    
    sort.Slice(rules, func(i, j int) bool {
        return rules[i].Priority > rules[j].Priority
    })
    
    return rules
}
```

### Complex Condition Evaluation

```go
// Condition evaluator with support for multiple operators
func NewBusinessRuleEngine() *BusinessRuleEngine {
    config := &condition.Config{
        MaxLevel: 10,
        Types: map[string]condition.TypeConfig{
            "text": {
                Operators: []condition.OperatorType{
                    condition.OpEqual, condition.OpNotEqual, condition.OpContains,
                    condition.OpStartsWith, condition.OpEndsWith, condition.OpIsEmpty,
                    condition.OpIsNotEmpty, condition.OpMatchRegexp,
                },
                ValueTypes: []condition.ValueType{
                    condition.ValueTypeValue, condition.ValueTypeField,
                },
            },
            "number": {
                Operators: []condition.OperatorType{
                    condition.OpEqual, condition.OpNotEqual, condition.OpLess,
                    condition.OpLessOrEqual, condition.OpGreater, 
                    condition.OpGreaterOrEqual, condition.OpBetween, 
                    condition.OpNotBetween,
                },
            },
            // ... more type configurations
        },
    }
    
    evaluator := condition.NewEvaluator(config, condition.DefaultEvalOptions())
    return &BusinessRuleEngine{
        rules:     make(map[string]*BusinessRule),
        evaluator: evaluator,
    }
}
```

### Calculated Fields

```go
// Rules can perform calculations between fields
func (bre *BusinessRuleEngine) calculateFieldValue(
    ctx context.Context,
    schema *Schema,
    fieldName string,
    params map[string]any,
    data map[string]any,
) error {
    formula, exists := params["formula"].(string)
    if !exists {
        return NewValidationError("missing_formula", "calculation requires formula parameter")
    }
    
    // Find and configure the field
    for i := range schema.Fields {
        if schema.Fields[i].Name == fieldName {
            // Mark as readonly calculated field
            schema.Fields[i].Readonly = true
            
            // Store formula for frontend
            if schema.Fields[i].Config == nil {
                schema.Fields[i].Config = make(map[string]any)
            }
            schema.Fields[i].Config["formula"] = formula
            schema.Fields[i].Config["calculated"] = true
            
            // Evaluate formula
            var result any
            switch formula {
            case "quantity * unit_price":
                qty, qtyOk := data["quantity"]
                price, priceOk := data["unit_price"]
                if !qtyOk || !priceOk {
                    return NewValidationError("formula_failed", "missing variables")
                }
                
                // Handle type conversion
                qtyFloat := toFloat64(qty)
                priceFloat := toFloat64(price)
                result = qtyFloat * priceFloat
                
            default:
                return NewValidationError("unsupported_formula", 
                    fmt.Sprintf("formula not supported: %s", formula))
            }
            
            // Set calculated value
            schema.Fields[i].Default = result
            return nil
        }
    }
    
    return NewValidationError("field_not_found", 
        fmt.Sprintf("field %s not found", fieldName))
}
```

### Rule Builder API

```go
// Fluent interface for building complex rules
type BusinessRuleBuilder struct {
    rule *BusinessRule
}

// Create visibility rule
visibilityRule, err := NewBusinessRule("shipping-rule", "Shipping Address Rule", RuleTypeFieldVisibility).
    WithDescription("Show shipping fields when shipping is required").
    WithPriority(100).
    WithCondition(conditionGroup).
    WithAction(ActionShowField, "shipping_street", nil).
    WithAction(ActionShowField, "shipping_city", nil).
    WithAction(ActionRequireField, "shipping_street", nil).
    Build()

// Create calculation rule  
calcRule, err := NewBusinessRule("total-calc", "Total Calculation", RuleTypeDataCalculation).
    WithActionAndParams(ActionCalculate, "total", nil, map[string]any{
        "formula": "quantity * unit_price",
    }).
    Build()

// Add to engine
engine := NewBusinessRuleEngine()
engine.AddRule(visibilityRule)
engine.AddRule(calcRule)
```

### Helper Functions for Common Rules

```go
// Quick creation of common rule types
func CreateFieldVisibilityRule(
    id, fieldName string,
    condition *condition.ConditionGroup,
    visible bool,
) (*BusinessRule, error) {
    actionType := ActionShowField
    if !visible {
        actionType = ActionHideField
    }
    
    return NewBusinessRule(id, fmt.Sprintf("%s visibility rule", fieldName), RuleTypeFieldVisibility).
        WithCondition(condition).
        WithAction(actionType, fieldName, nil).
        Build()
}

func CreateCalculationRule(
    id, fieldName, formula string,
    condition *condition.ConditionGroup,
) (*BusinessRule, error) {
    params := map[string]any{"formula": formula}
    
    return NewBusinessRule(id, fmt.Sprintf("%s calculation rule", fieldName), RuleTypeDataCalculation).
        WithCondition(condition).
        WithActionAndParams(ActionCalculate, fieldName, nil, params).
        Build()
}
```

### Rule Testing & Debugging

```go
// Test rules against multiple scenarios
func (bre *BusinessRuleEngine) TestRule(
    ctx context.Context,
    rule *BusinessRule,
    schema *Schema,
    testCases []map[string]any,
) ([]bool, error) {
    results := make([]bool, len(testCases))
    for i, data := range testCases {
        shouldApply, err := bre.evaluateRuleCondition(ctx, rule, data)
        if err != nil {
            return nil, err
        }
        results[i] = shouldApply
    }
    return results, nil
}

// Get detailed explanation of rule behavior
func (bre *BusinessRuleEngine) ExplainRule(
    ctx context.Context,
    rule *BusinessRule,
    data map[string]any,
) (string, error) {
    explanation := fmt.Sprintf("Rule: %s\n", rule.Name)
    explanation += fmt.Sprintf("Type: %s\n", rule.Type)
    explanation += fmt.Sprintf("Priority: %d\n", rule.Priority)
    
    shouldApply, err := bre.evaluateRuleCondition(ctx, rule, data)
    if err != nil {
        return "", err
    }
    
    explanation += fmt.Sprintf("Condition Met: %v\n", shouldApply)
    explanation += fmt.Sprintf("Actions: %d\n", len(rule.Actions))
    
    for i, action := range rule.Actions {
        explanation += fmt.Sprintf("  Action %d: %s on %s\n", 
            i+1, action.Type, action.Target)
    }
    
    return explanation, nil
}
```

---

[← Back](07-layout.md) | [Next: Storage →](09-storage-interface.md)