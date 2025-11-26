package schema

import (
	"fmt"
	"sync"
)

// BaseBuilder provides common builder functionality with generic type support
type BaseBuilder[T any] interface {
	Build() (*T, error)
	HasErrors() bool
	GetErrors() []error
	AddError(error) BaseBuilder[T]
	ClearErrors() BaseBuilder[T]
}

// BuilderContext provides shared builder state and error handling
type BuilderContext struct {
	errors []error
	mu     sync.RWMutex
}

// NewBuilderContext creates a new builder context
func NewBuilderContext() *BuilderContext {
	return &BuilderContext{
		errors: []error{},
	}
}

// AddError adds an error to the builder context
func (bc *BuilderContext) AddError(err error) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.errors = append(bc.errors, err)
}

// AddErrorf adds a formatted error to the builder context
func (bc *BuilderContext) AddErrorf(format string, args ...any) {
	bc.AddError(fmt.Errorf(format, args...))
}

// HasErrors returns true if there are any errors
func (bc *BuilderContext) HasErrors() bool {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return len(bc.errors) > 0
}

// GetErrors returns a copy of all errors
func (bc *BuilderContext) GetErrors() []error {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	if len(bc.errors) == 0 {
		return nil
	}
	errorsCopy := make([]error, len(bc.errors))
	copy(errorsCopy, bc.errors)
	return errorsCopy
}

// ClearErrors removes all errors from the context
func (bc *BuilderContext) ClearErrors() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.errors = []error{}
}

// ErrorCount returns the number of errors
func (bc *BuilderContext) ErrorCount() int {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	return len(bc.errors)
}

// FirstError returns the first error, or nil if no errors
func (bc *BuilderContext) FirstError() error {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	if len(bc.errors) == 0 {
		return nil
	}
	return bc.errors[0]
}

// CombinedError returns all errors combined as a single error
func (bc *BuilderContext) CombinedError() error {
	errors := bc.GetErrors()
	if len(errors) == 0 {
		return nil
	}
	if len(errors) == 1 {
		return errors[0]
	}
	
	// Combine multiple errors
	message := fmt.Sprintf("builder has %d errors:", len(errors))
	for i, err := range errors {
		message += fmt.Sprintf("\n  %d. %v", i+1, err)
	}
	return fmt.Errorf(message)
}

// ValidateField validates a field name is non-empty
func (bc *BuilderContext) ValidateField(fieldName, value, context string) {
	if value == "" {
		bc.AddErrorf("%s %s cannot be empty", context, fieldName)
	}
}

// ValidateRequired validates a required value is present
func (bc *BuilderContext) ValidateRequired(value any, fieldName, context string) {
	if value == nil {
		bc.AddErrorf("%s %s is required", context, fieldName)
		return
	}
	
	// Check for empty strings
	if str, ok := value.(string); ok && str == "" {
		bc.AddErrorf("%s %s cannot be empty", context, fieldName)
	}
}

// ValidateRange validates a numeric value is within range
func (bc *BuilderContext) ValidateRange(value, min, max int, fieldName, context string) {
	if value < min || value > max {
		bc.AddErrorf("%s %s must be between %d and %d (got %d)", 
			context, fieldName, min, max, value)
	}
}

// ValidateOneOf validates a value is one of the allowed options
func (bc *BuilderContext) ValidateOneOf(value string, allowed []string, fieldName, context string) {
	for _, option := range allowed {
		if value == option {
			return
		}
	}
	bc.AddErrorf("%s %s must be one of %v (got %q)", 
		context, fieldName, allowed, value)
}

// BuilderValidator provides common validation methods for builders
type BuilderValidator struct {
	context *BuilderContext
}

// NewBuilderValidator creates a new validator with the given context
func NewBuilderValidator(context *BuilderContext) *BuilderValidator {
	return &BuilderValidator{context: context}
}

// ValidateID validates an ID field is valid
func (bv *BuilderValidator) ValidateID(id, builderType string) {
	if id == "" {
		bv.context.AddErrorf("%s ID cannot be empty", builderType)
		return
	}
	
	// Basic ID validation (alphanumeric, hyphens, underscores)
	for _, char := range id {
		if !((char >= 'a' && char <= 'z') || 
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_' || char == '.') {
			bv.context.AddErrorf("%s ID contains invalid character: %c", builderType, char)
			return
		}
	}
}

// ValidateEnum validates an enum value
func (bv *BuilderValidator) ValidateEnum(value string, validValues []string, fieldName, builderType string) {
	if value == "" {
		return // Empty is often acceptable for optional enums
	}
	
	for _, valid := range validValues {
		if value == valid {
			return
		}
	}
	
	bv.context.AddErrorf("%s %s must be one of %v (got %q)", 
		builderType, fieldName, validValues, value)
}

// BuilderMixin provides reusable builder functionality
type BuilderMixin struct {
	Context   *BuilderContext
	Validator *BuilderValidator
}

// NewBuilderMixin creates a new builder mixin
func NewBuilderMixin() *BuilderMixin {
	ctx := NewBuilderContext()
	return &BuilderMixin{
		Context:   ctx,
		Validator: NewBuilderValidator(ctx),
	}
}

// CheckBuild performs final validation before building
func (bm *BuilderMixin) CheckBuild(builderType string) error {
	if bm.Context.HasErrors() {
		return fmt.Errorf("cannot build %s with errors: %w", 
			builderType, bm.Context.CombinedError())
	}
	return nil
}

// SafeBuild safely builds an object, returning an error if there are validation issues
func SafeBuild[T any](builder BaseBuilder[T], builderType string) (*T, error) {
	if builder.HasErrors() {
		return nil, fmt.Errorf("cannot build %s: %s", builderType, 
			combineErrors(builder.GetErrors()))
	}
	return builder.Build()
}

// combineErrors combines multiple errors into a single error message
func combineErrors(errors []error) string {
	if len(errors) == 0 {
		return "no errors"
	}
	if len(errors) == 1 {
		return errors[0].Error()
	}
	
	message := fmt.Sprintf("%d errors:", len(errors))
	for i, err := range errors {
		message += fmt.Sprintf(" (%d) %v", i+1, err)
	}
	return message
}

// BuilderStats provides statistics about builder usage (optional)
type BuilderStats struct {
	TotalBuilds    int64 `json:"totalBuilds"`
	SuccessfulBuilds int64 `json:"successfulBuilds"`
	FailedBuilds   int64 `json:"failedBuilds"`
	mu             sync.RWMutex
}

// RecordBuild records a build attempt
func (bs *BuilderStats) RecordBuild(success bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	bs.TotalBuilds++
	if success {
		bs.SuccessfulBuilds++
	} else {
		bs.FailedBuilds++
	}
}

// GetStats returns a copy of the current stats
func (bs *BuilderStats) GetStats() BuilderStats {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	return BuilderStats{
		TotalBuilds:      bs.TotalBuilds,
		SuccessfulBuilds: bs.SuccessfulBuilds,
		FailedBuilds:     bs.FailedBuilds,
	}
}