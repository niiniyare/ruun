# Runtime Interfaces Analysis

## Overview

This analysis identified the missing runtime-related interfaces in the ruun schema package and provides their complete definitions based on usage patterns found in the codebase.

## Missing Interfaces Identified

The following interfaces were referenced in the runtime implementation but not defined:

1. **RuntimeRenderer** - UI rendering interface
2. **RuntimeValidator** - Real-time validation interface  
3. **RuntimeConditionalEngine** - Conditional logic evaluation interface
4. **RuntimeStateManager** - Form state management interface
5. **RuntimeStats** - Runtime performance statistics

## Interface Definitions Created

### 1. RuntimeRenderer

**Purpose**: Converts schema definitions into UI components for any frontend framework.

**Key Methods**:
- `RenderField()` - Renders individual form fields
- `RenderForm()` - Renders complete forms
- `RenderAction()` - Renders buttons/actions
- `RenderLayout()`, `RenderSection()`, `RenderTab()`, etc. - Layout components
- `RenderErrors()` - Error message display

**Usage**: UI frameworks (React, Templ, Go templates) implement this interface to render schema-driven forms.

### 2. RuntimeValidator  

**Purpose**: Performs real-time validation during user interaction.

**Key Methods**:
- `ValidateField()` - Validates single field values
- `ValidateAllFields()` - Validates entire form
- `ValidateAction()` - Validates if actions can be performed
- `ValidateWorkflow()` - Validates workflow state transitions
- `ValidateBusinessRules()` - Applies business logic validation
- `ValidateFieldAsync()` - Async validation for expensive operations

**Usage**: Integrates with the existing validator system to provide runtime validation.

### 3. RuntimeConditionalEngine

**Purpose**: Evaluates conditional logic to show/hide fields and determine field states.

**Key Methods**:
- `EvaluateFieldVisibility()` - Determines if fields should be visible
- `EvaluateFieldRequired()` - Determines if fields should be required
- `EvaluateFieldEditable()` - Determines if fields should be editable
- `EvaluateActionConditions()` - Evaluates action availability
- `EvaluateAllConditions()` - Batch evaluation for performance

**Usage**: Implements dynamic form behavior based on user input and form state.

### 4. RuntimeStateManager

**Purpose**: Manages form state including values, touched fields, dirty tracking, and errors.

**Key Methods**:
- `GetValue()`, `SetValue()` - Field value management
- `IsTouched()`, `Touch()` - User interaction tracking
- `IsDirty()`, `IsAnyDirty()` - Change tracking
- `GetErrors()`, `SetErrors()` - Validation error management
- `CreateSnapshot()`, `RestoreSnapshot()` - State snapshots for undo/redo

**Usage**: Provides state management abstraction for runtime implementation.

### 5. RuntimeStats

**Purpose**: Provides performance and usage statistics for the runtime.

**Structure**: Contains metrics like field counts, error counts, event statistics, and timing information.

**Usage**: Monitoring and debugging runtime performance.

## Implementation Patterns Found

### 1. Interface Usage in Runtime

The runtime implementation shows these interfaces are used as:
- **Dependency injection** - Interfaces injected via RuntimeBuilder
- **Optional dependencies** - Graceful degradation when interfaces not provided  
- **Builder pattern** - Fluent configuration of runtime with implementations

### 2. Error Handling

The interfaces use consistent error handling patterns:
- Return `error` for operations that can fail
- Return `[]string` for validation errors (multiple errors per field)
- Use context for cancellation and timeouts

### 3. Performance Considerations

Several performance optimization patterns were identified:
- **Batch operations** - `EvaluateAllConditions()` for efficiency
- **Async validation** - Non-blocking validation with callbacks
- **Debouncing support** - Built into validation timing configuration
- **Caching support** - State snapshots for undo/redo functionality

## Integration with Existing Code

### Runtime.go Integration

The `runtime.go` file shows these interfaces are used in:

1. **Builder pattern configuration**:
   ```go
   runtime := NewRuntimeBuilder(schema).
       WithRenderer(renderer).
       WithValidator(validator).
       WithConditionalEngine(engine).
       Build()
   ```

2. **Event handling integration**:
   - Field changes trigger validation via `RuntimeValidator`
   - Conditional logic applied via `RuntimeConditionalEngine` 
   - UI updates via `RuntimeRenderer`

3. **State management**:
   - State tracked via `RuntimeStateManager`
   - Statistics collected via `RuntimeStats`

### Existing Validator Integration

The new `RuntimeValidator` interface integrates with existing validation:
- Extends the current `validator.go` implementation
- Supports both sync and async validation
- Maintains compatibility with business rules engine

## Implementation Recommendations

### 1. For UI Framework Developers

Implement `RuntimeRenderer` interface:
```go
type MyFrameworkRenderer struct {
    templates map[string]Template
    theme     *Theme
}

func (r *MyFrameworkRenderer) RenderField(ctx context.Context, field *Field, ...) (string, error) {
    // Map field type to component template
    // Apply theme tokens
    // Render with framework-specific logic
}
```

### 2. For Validation Extensions

Extend existing validator to implement `RuntimeValidator`:
```go
type ExtendedValidator struct {
    *schema.Validator // Embed existing validator
    db Database      // For async validations
}

func (v *ExtendedValidator) ValidateFieldAsync(ctx context.Context, field *Field, ...) error {
    // Implement expensive validation logic
}
```

### 3. For Conditional Logic

Implement `RuntimeConditionalEngine` with expression evaluator:
```go
type ConditionalEngine struct {
    evaluator *condition.Evaluator
}

func (e *ConditionalEngine) EvaluateFieldVisibility(ctx context.Context, field *Field, data map[string]any) (bool, error) {
    // Evaluate field conditions against current form data
}
```

## Testing Strategy

The interfaces enable comprehensive testing:

1. **Mock implementations** for unit testing
2. **Interface compliance tests** to verify implementations
3. **Integration tests** with real UI frameworks
4. **Performance benchmarks** using the statistics interfaces

## Future Extensibility

The interface design supports future extensions:

- **Plugin architecture** - New renderers/validators as plugins
- **Custom conditional logic** - Domain-specific condition evaluators  
- **Advanced state management** - Time-travel debugging, persistence
- **Performance monitoring** - Detailed runtime analytics

## Files Created

1. **`runtime_interfaces.go`** - Complete interface definitions
2. **`RUNTIME_INTERFACES_ANALYSIS.md`** - This analysis document

These interfaces provide the foundation for a flexible, extensible runtime system that can work with any UI framework while maintaining strong typing and clear separation of concerns.