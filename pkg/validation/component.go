package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// ComponentValidator validates UI components
type ComponentValidator struct {
	schemas map[string]*ComponentSchema
	engine  *ValidationEngine
}

// ComponentSchema defines the validation schema for a component
type ComponentSchema struct {
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Description string                    `json:"description"`
	Props       map[string]*PropSchema    `json:"props"`
	Children    *ChildrenSchema           `json:"children,omitempty"`
	Composition *CompositionSchema        `json:"composition,omitempty"`
	Variants    map[string]*VariantSchema `json:"variants,omitempty"`
	States      map[string]*StateSchema   `json:"states,omitempty"`
}

// PropSchema defines validation for a single prop
type PropSchema struct {
	Name         string          `json:"name"`
	Type         PropType        `json:"type"`
	Required     bool            `json:"required"`
	Default      any             `json:"default,omitempty"`
	Description  string          `json:"description"`
	Deprecated   bool            `json:"deprecated"`
	Validation   *PropValidation `json:"validation,omitempty"`
	Examples     []any           `json:"examples,omitempty"`
	Dependencies []string        `json:"dependencies,omitempty"` // Props that must be present if this prop is used
	Conflicts    []string        `json:"conflicts,omitempty"`    // Props that cannot be used with this prop
}

// PropType represents the type of a prop
type PropType struct {
	Kind        PropKind             `json:"kind"`
	ElementType *PropType            `json:"elementType,omitempty"` // For arrays
	Properties  map[string]*PropType `json:"properties,omitempty"`  // For objects
	Union       []*PropType          `json:"union,omitempty"`       // For union types
	Enum        []any                `json:"enum,omitempty"`        // For enum types
}

// PropKind represents the kind of prop type
type PropKind string

const (
	PropKindString    PropKind = "string"
	PropKindNumber    PropKind = "number"
	PropKindBoolean   PropKind = "boolean"
	PropKindArray     PropKind = "array"
	PropKindObject    PropKind = "object"
	PropKindFunction  PropKind = "function"
	PropKindComponent PropKind = "component"
	PropKindUnion     PropKind = "union"
	PropKindEnum      PropKind = "enum"
	PropKindAny       PropKind = "any"
)

// PropValidation defines additional validation rules for props
type PropValidation struct {
	MinLength *int             `json:"minLength,omitempty"`
	MaxLength *int             `json:"maxLength,omitempty"`
	Min       *float64         `json:"min,omitempty"`
	Max       *float64         `json:"max,omitempty"`
	Pattern   string           `json:"pattern,omitempty"`
	Custom    []string         `json:"custom,omitempty"` // Custom validator function names
	Rules     []ValidationRule `json:"rules,omitempty"`
}

// ChildrenSchema defines validation for component children
type ChildrenSchema struct {
	Required  bool     `json:"required"`
	Types     []string `json:"types,omitempty"`    // Allowed child component types
	MaxCount  *int     `json:"maxCount,omitempty"` // Maximum number of children
	MinCount  *int     `json:"minCount,omitempty"` // Minimum number of children
	AllowText bool     `json:"allowText"`          // Whether text content is allowed
	AllowHTML bool     `json:"allowHTML"`          // Whether HTML content is allowed
}

// CompositionSchema defines how components can be composed
type CompositionSchema struct {
	Slots       map[string]*SlotSchema `json:"slots,omitempty"`
	Render      *RenderSchema          `json:"render,omitempty"`
	Polymorphic bool                   `json:"polymorphic"` // Can render as different elements
	AsChild     bool                   `json:"asChild"`     // Can accept a child component to render as
}

// SlotSchema defines a slot for composition
type SlotSchema struct {
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Description string   `json:"description"`
	Types       []string `json:"types,omitempty"` // Allowed component types for this slot
}

// RenderSchema defines render prop validation
type RenderSchema struct {
	Required   bool                   `json:"required"`
	Parameters map[string]*PropSchema `json:"parameters"` // Parameters passed to render function
}

// VariantSchema defines component variants
type VariantSchema struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Props       map[string]*PropSchema `json:"props"`     // Additional props for this variant
	Required    []string               `json:"required"`  // Required props for this variant
	Forbidden   []string               `json:"forbidden"` // Forbidden props for this variant
}

// StateSchema defines component state validation
type StateSchema struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Props       map[string]*PropSchema `json:"props"`      // Props affected by this state
	Conditions  []string               `json:"conditions"` // Conditions that trigger this state
}

// ComponentInstance represents an instance of a component being validated
type ComponentInstance struct {
	Type     string              `json:"type"`
	Props    map[string]any      `json:"props"`
	Children []ComponentInstance `json:"children,omitempty"`
	Slots    map[string]any      `json:"slots,omitempty"`
	Location *SourceLocation     `json:"location,omitempty"`
	Metadata map[string]any      `json:"metadata,omitempty"`
}

// NewComponentValidator creates a new component validator
func NewComponentValidator() *ComponentValidator {
	return &ComponentValidator{
		schemas: make(map[string]*ComponentSchema),
	}
}

// RegisterComponentSchema registers a component schema
func (cv *ComponentValidator) RegisterComponentSchema(schema *ComponentSchema) {
	cv.schemas[schema.Name] = schema
}

// ValidateComponent validates a component instance
func (cv *ComponentValidator) ValidateComponent(instance *ComponentInstance) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Level:     ValidationLevelError,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Find component schema
	schema, exists := cv.schemas[instance.Type]
	if !exists {
		result.Errors = append(result.Errors, NewValidationError(
			"component.unknown",
			fmt.Sprintf("Unknown component type: %s", instance.Type),
			"type",
			ValidationLevelError,
		))
		result.Valid = false
		return result
	}

	// Validate props
	propResult := cv.validateProps(instance.Props, schema.Props, instance.Type)
	result = cv.mergeResults(result, propResult)

	// Validate children
	if schema.Children != nil {
		childrenResult := cv.validateChildren(instance.Children, schema.Children, instance.Type)
		result = cv.mergeResults(result, childrenResult)
	}

	// Validate composition
	if schema.Composition != nil {
		compResult := cv.validateComposition(instance, schema.Composition)
		result = cv.mergeResults(result, compResult)
	}

	// Validate variants
	if len(schema.Variants) > 0 {
		variantResult := cv.validateVariants(instance, schema.Variants)
		result = cv.mergeResults(result, variantResult)
	}

	return result
}

// validateProps validates component props
func (cv *ComponentValidator) validateProps(props map[string]any, schemas map[string]*PropSchema, componentType string) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Check for required props
	for propName, propSchema := range schemas {
		if propSchema.Required {
			if _, exists := props[propName]; !exists {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.required",
					fmt.Sprintf("Required prop '%s' is missing", propName),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	// Validate each provided prop
	for propName, propValue := range props {
		propSchema, exists := schemas[propName]
		if !exists {
			result.Warnings = append(result.Warnings, NewValidationWarning(
				"prop.unknown",
				fmt.Sprintf("Unknown prop '%s' for component %s", propName, componentType),
				propName,
			))
			continue
		}

		// Check if prop is deprecated
		if propSchema.Deprecated {
			result.Warnings = append(result.Warnings, NewValidationWarning(
				"prop.deprecated",
				fmt.Sprintf("Prop '%s' is deprecated", propName),
				propName,
			))
		}

		// Validate prop type
		typeResult := cv.validatePropType(propValue, propSchema.Type, propName)
		if !typeResult.Valid {
			result.Errors = append(result.Errors, typeResult.Errors...)
			result.Valid = false
		}

		// Validate prop constraints
		if propSchema.Validation != nil {
			validationResult := cv.validatePropConstraints(propValue, propSchema.Validation, propName)
			if !validationResult.Valid {
				result.Errors = append(result.Errors, validationResult.Errors...)
				result.Valid = false
			}
		}

		// Check dependencies
		for _, dep := range propSchema.Dependencies {
			if _, exists := props[dep]; !exists {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.dependency",
					fmt.Sprintf("Prop '%s' requires prop '%s' to be present", propName, dep),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}

		// Check conflicts
		for _, conflict := range propSchema.Conflicts {
			if _, exists := props[conflict]; exists {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.conflict",
					fmt.Sprintf("Prop '%s' cannot be used with prop '%s'", propName, conflict),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	return result
}

// validatePropType validates the type of a prop value
func (cv *ComponentValidator) validatePropType(value any, propType PropType, propName string) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	if value == nil {
		return result
	}

	switch propType.Kind {
	case PropKindString:
		if _, ok := value.(string); !ok {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.string",
				fmt.Sprintf("Prop '%s' must be a string", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindNumber:
		switch value.(type) {
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			// Valid number types
		case string:
			// Try to parse string as number
			if _, err := strconv.ParseFloat(value.(string), 64); err != nil {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.type.number",
					fmt.Sprintf("Prop '%s' must be a number", propName),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		default:
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.number",
				fmt.Sprintf("Prop '%s' must be a number", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindBoolean:
		if _, ok := value.(bool); !ok {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.boolean",
				fmt.Sprintf("Prop '%s' must be a boolean", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindArray:
		if reflect.TypeOf(value).Kind() != reflect.Slice {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.array",
				fmt.Sprintf("Prop '%s' must be an array", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		} else if propType.ElementType != nil {
			// Validate array elements
			slice := reflect.ValueOf(value)
			for i := 0; i < slice.Len(); i++ {
				element := slice.Index(i).Interface()
				elementResult := cv.validatePropType(element, *propType.ElementType, fmt.Sprintf("%s[%d]", propName, i))
				if !elementResult.Valid {
					result.Errors = append(result.Errors, elementResult.Errors...)
					result.Valid = false
				}
			}
		}

	case PropKindObject:
		if reflect.TypeOf(value).Kind() != reflect.Map {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.object",
				fmt.Sprintf("Prop '%s' must be an object", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		} else if len(propType.Properties) > 0 {
			// Validate object properties
			objMap, ok := value.(map[string]any)
			if ok {
				for propKey, propVal := range objMap {
					if propSchema, exists := propType.Properties[propKey]; exists {
						propResult := cv.validatePropType(propVal, *propSchema, fmt.Sprintf("%s.%s", propName, propKey))
						if !propResult.Valid {
							result.Errors = append(result.Errors, propResult.Errors...)
							result.Valid = false
						}
					}
				}
			}
		}

	case PropKindFunction:
		valType := reflect.TypeOf(value)
		if valType.Kind() != reflect.Func {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.function",
				fmt.Sprintf("Prop '%s' must be a function", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindUnion:
		// Validate against any of the union types
		valid := false
		for _, unionType := range propType.Union {
			unionResult := cv.validatePropType(value, *unionType, propName)
			if unionResult.Valid {
				valid = true
				break
			}
		}
		if !valid {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.union",
				fmt.Sprintf("Prop '%s' does not match any of the allowed types", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindEnum:
		// Check if value is in enum
		valid := false
		for _, enumValue := range propType.Enum {
			if value == enumValue {
				valid = true
				break
			}
		}
		if !valid {
			result.Errors = append(result.Errors, NewValidationError(
				"prop.type.enum",
				fmt.Sprintf("Prop '%s' must be one of the allowed values", propName),
				propName,
				ValidationLevelError,
			))
			result.Valid = false
		}

	case PropKindAny:
		// Any type is allowed
		break

	default:
		result.Warnings = append(result.Warnings, NewValidationWarning(
			"prop.type.unknown",
			fmt.Sprintf("Unknown prop type for '%s'", propName),
			propName,
		))
	}

	return result
}

// validatePropConstraints validates prop constraints
func (cv *ComponentValidator) validatePropConstraints(value any, validation *PropValidation, propName string) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// String length validation
	if validation.MinLength != nil || validation.MaxLength != nil {
		if str, ok := value.(string); ok {
			if validation.MinLength != nil && len(str) < *validation.MinLength {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.constraint.minLength",
					fmt.Sprintf("Prop '%s' must be at least %d characters long", propName, *validation.MinLength),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
			if validation.MaxLength != nil && len(str) > *validation.MaxLength {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.constraint.maxLength",
					fmt.Sprintf("Prop '%s' must be at most %d characters long", propName, *validation.MaxLength),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	// Numeric range validation
	if validation.Min != nil || validation.Max != nil {
		var num float64
		var ok bool

		switch v := value.(type) {
		case int:
			num, ok = float64(v), true
		case int64:
			num, ok = float64(v), true
		case float64:
			num, ok = v, true
		case float32:
			num, ok = float64(v), true
		case string:
			if parsed, err := strconv.ParseFloat(v, 64); err == nil {
				num, ok = parsed, true
			}
		}

		if ok {
			if validation.Min != nil && num < *validation.Min {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.constraint.min",
					fmt.Sprintf("Prop '%s' must be at least %g", propName, *validation.Min),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
			if validation.Max != nil && num > *validation.Max {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.constraint.max",
					fmt.Sprintf("Prop '%s' must be at most %g", propName, *validation.Max),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	// Pattern validation
	if validation.Pattern != "" {
		if str, ok := value.(string); ok {
			// Note: In a real implementation, you'd use regexp.MatchString
			// For this example, we'll do a simple contains check
			if !strings.Contains(str, validation.Pattern) {
				result.Errors = append(result.Errors, NewValidationError(
					"prop.constraint.pattern",
					fmt.Sprintf("Prop '%s' does not match the required pattern", propName),
					propName,
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	return result
}

// validateChildren validates component children
func (cv *ComponentValidator) validateChildren(children []ComponentInstance, schema *ChildrenSchema, componentType string) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Check required children
	if schema.Required && len(children) == 0 {
		result.Errors = append(result.Errors, NewValidationError(
			"children.required",
			fmt.Sprintf("Component %s requires children", componentType),
			"children",
			ValidationLevelError,
		))
		result.Valid = false
	}

	// Check count constraints
	if schema.MinCount != nil && len(children) < *schema.MinCount {
		result.Errors = append(result.Errors, NewValidationError(
			"children.minCount",
			fmt.Sprintf("Component %s requires at least %d children", componentType, *schema.MinCount),
			"children",
			ValidationLevelError,
		))
		result.Valid = false
	}

	if schema.MaxCount != nil && len(children) > *schema.MaxCount {
		result.Errors = append(result.Errors, NewValidationError(
			"children.maxCount",
			fmt.Sprintf("Component %s allows at most %d children", componentType, *schema.MaxCount),
			"children",
			ValidationLevelError,
		))
		result.Valid = false
	}

	// Validate child types
	if len(schema.Types) > 0 {
		for _, child := range children {
			if !contains(schema.Types, child.Type) {
				result.Errors = append(result.Errors, NewValidationError(
					"children.type",
					fmt.Sprintf("Child type '%s' is not allowed in component %s", child.Type, componentType),
					"children",
					ValidationLevelError,
				))
				result.Valid = false
			}
		}
	}

	return result
}

// validateComposition validates component composition rules
func (cv *ComponentValidator) validateComposition(instance *ComponentInstance, schema *CompositionSchema) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Validate slots
	if len(schema.Slots) > 0 {
		for slotName, slotSchema := range schema.Slots {
			slotContent, exists := instance.Slots[slotName]

			if slotSchema.Required && !exists {
				result.Errors = append(result.Errors, NewValidationError(
					"slot.required",
					fmt.Sprintf("Required slot '%s' is missing", slotName),
					slotName,
					ValidationLevelError,
				))
				result.Valid = false
			}

			if exists && slotContent != nil {
				// Validate slot content type if specified
				if len(slotSchema.Types) > 0 {
					// This would need more complex implementation to validate slot content types
				}
			}
		}
	}

	return result
}

// validateVariants validates component variants
func (cv *ComponentValidator) validateVariants(instance *ComponentInstance, variants map[string]*VariantSchema) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Determine which variant is being used (simplified logic)
	for variantName, variantSchema := range variants {
		isVariant := true

		// Check if all required props for this variant are present
		for _, requiredProp := range variantSchema.Required {
			if _, exists := instance.Props[requiredProp]; !exists {
				isVariant = false
				break
			}
		}

		// Check if any forbidden props are present
		for _, forbiddenProp := range variantSchema.Forbidden {
			if _, exists := instance.Props[forbiddenProp]; exists {
				isVariant = false
				break
			}
		}

		if isVariant {
			// Validate variant-specific props
			variantResult := cv.validateProps(instance.Props, variantSchema.Props, instance.Type)
			if !variantResult.Valid {
				result.Errors = append(result.Errors, variantResult.Errors...)
				result.Valid = false
			}
			break
		}
	}

	return result
}

// mergeResults merges two validation results
func (cv *ComponentValidator) mergeResults(base, addition *ValidationResult) *ValidationResult {
	if !addition.Valid {
		base.Valid = false
	}

	base.Errors = append(base.Errors, addition.Errors...)
	base.Warnings = append(base.Warnings, addition.Warnings...)

	// Merge metadata
	for k, v := range addition.Metadata {
		base.Metadata[k] = v
	}

	return base
}

// Built-in validators

// RequiredPropsValidator validates required props
type RequiredPropsValidator struct{}

func (v *RequiredPropsValidator) Validate(ctx *ValidationContext, value any) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	if instance, ok := value.(*ComponentInstance); ok {
		cv := NewComponentValidator()
		return cv.ValidateComponent(instance)
	}

	return result
}

func (v *RequiredPropsValidator) GetRule() ValidationRule {
	return ValidationRule{
		ID:          "component.props.required",
		Name:        "Required Props",
		Description: "Validates that required props are provided",
		Category:    "component",
		Level:       ValidationLevelError,
		Enabled:     true,
	}
}

// PropTypesValidator validates prop types
type PropTypesValidator struct{}

func (v *PropTypesValidator) Validate(ctx *ValidationContext, value any) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Implementation would be similar to RequiredPropsValidator
	// but focus on type validation

	return result
}

func (v *PropTypesValidator) GetRule() ValidationRule {
	return ValidationRule{
		ID:          "component.props.types",
		Name:        "Prop Types",
		Description: "Validates that props have correct types",
		Category:    "component",
		Level:       ValidationLevelError,
		Enabled:     true,
	}
}
