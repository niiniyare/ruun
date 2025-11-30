package schema

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
)

// ValidationRegistry manages custom validators and provides extensible validation
type ValidationRegistry struct {
	validators      map[string]ValidatorFunc
	asyncValidators map[string]*AsyncValidator
	mu              sync.RWMutex
}

// AsyncValidator wraps async validation with caching and debouncing
type AsyncValidator struct {
	Name     string
	Validate AsyncValidatorFunc
	Debounce time.Duration
	Cache    bool
	CacheTTL time.Duration
	cache    map[string]*CacheEntry
	mu       sync.RWMutex
	cleanup  context.CancelFunc
}

// CrossFieldValidator validates multiple fields together using condition engine
type CrossFieldValidator struct {
	Name      string
	Fields    []string
	Condition *condition.ConditionGroup
	Validate  func(ctx context.Context, values map[string]any) error
}

// CrossFieldValidationRegistry manages cross-field validators
type CrossFieldValidationRegistry struct {
	validators map[string]*CrossFieldValidator
	evaluator  *condition.Evaluator
	mu         sync.RWMutex
}

// ValidatorChain allows chaining multiple validators
type ValidatorChain struct {
	validators []ValidatorFunc
}

// NewValidationRegistry creates a new validation registry with built-in validators
func NewValidationRegistry() *ValidationRegistry {
	registry := &ValidationRegistry{
		validators:      make(map[string]ValidatorFunc),
		asyncValidators: make(map[string]*AsyncValidator),
	}
	// Register built-in validators
	registry.registerBuiltInValidators()
	return registry
}

// NewCrossFieldValidationRegistry creates a new cross-field validation registry
func NewCrossFieldValidationRegistry() *CrossFieldValidationRegistry {
	// Initialize condition evaluator with reasonable defaults
	config := &condition.Config{
		MaxLevel: 5,
		Types: map[string]condition.TypeConfig{
			"text": {
				Operators: []condition.OperatorType{
					condition.OpEqual, condition.OpNotEqual, condition.OpContains,
					condition.OpStartsWith, condition.OpEndsWith, condition.OpIsEmpty,
					condition.OpIsNotEmpty, condition.OpMatchRegexp,
				},
				ValueTypes: []condition.ValueType{condition.ValueTypeValue, condition.ValueTypeField},
			},
			"number": {
				Operators: []condition.OperatorType{
					condition.OpEqual, condition.OpNotEqual, condition.OpLess,
					condition.OpLessOrEqual, condition.OpGreater, condition.OpGreaterOrEqual,
					condition.OpBetween, condition.OpNotBetween,
				},
				ValueTypes: []condition.ValueType{condition.ValueTypeValue, condition.ValueTypeField},
			},
			"date": {
				Operators: []condition.OperatorType{
					condition.OpEqual, condition.OpNotEqual, condition.OpLess,
					condition.OpLessOrEqual, condition.OpGreater, condition.OpGreaterOrEqual,
					condition.OpBetween, condition.OpNotBetween,
				},
				ValueTypes: []condition.ValueType{condition.ValueTypeValue, condition.ValueTypeField},
			},
		},
	}
	evaluator := condition.NewEvaluator(config, condition.DefaultEvalOptions())
	registry := &CrossFieldValidationRegistry{
		validators: make(map[string]*CrossFieldValidator),
		evaluator:  evaluator,
	}
	// Register built-in cross-field validators
	registry.registerBuiltInValidators()
	return registry
}

// NewValidatorChain creates a new validator chain
func NewValidatorChain() *ValidatorChain {
	return &ValidatorChain{
		validators: []ValidatorFunc{},
	}
}

// Register adds a custom validator
func (r *ValidationRegistry) Register(name string, validator ValidatorFunc) error {
	if name == "" {
		return NewValidationError("validator_name", "validator name is required")
	}
	if validator == nil {
		return NewValidationError("validator_func", "validator function is required")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.validators[name] = validator
	return nil
}

// RegisterAsync adds an async validator with cache cleanup
func (r *ValidationRegistry) RegisterAsync(validator *AsyncValidator) error {
	if validator == nil {
		return NewValidationError("async_validator", "async validator is required")
	}
	if validator.Name == "" {
		return NewValidationError("async_validator_name", "async validator name is required")
	}
	if validator.Validate == nil {
		return NewValidationError("async_validator_func", "async validator function is required")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if validator.cache == nil {
		validator.cache = make(map[string]*CacheEntry)
	}
	// Start cache cleanup goroutine if caching is enabled
	if validator.Cache && validator.CacheTTL > 0 {
		ctx, cancel := context.WithCancel(context.Background())
		validator.cleanup = cancel
		validator.startCacheCleanup(ctx)
	}
	r.asyncValidators[validator.Name] = validator
	return nil
}

// Get retrieves a validator by name
func (r *ValidationRegistry) Get(name string) (ValidatorFunc, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	validator, exists := r.validators[name]
	return validator, exists
}

// Validate executes a validator by name
func (r *ValidationRegistry) Validate(ctx context.Context, name string, value any, params map[string]any) error {
	r.mu.RLock()
	validator, exists := r.validators[name]
	r.mu.RUnlock()
	if !exists {
		return NewValidationError("validator_not_found", fmt.Sprintf("validator %s not found", name))
	}
	return validator(ctx, value, params)
}

// ValidateAsync executes an async validator with caching and debouncing
func (r *ValidationRegistry) ValidateAsync(ctx context.Context, name string, value any, params map[string]any) error {
	r.mu.RLock()
	validator, exists := r.asyncValidators[name]
	r.mu.RUnlock()
	if !exists {
		return NewValidationError("async_validator_not_found", fmt.Sprintf("async validator %s not found", name))
	}
	// Check cache if enabled
	if validator.Cache {
		cacheKey := fmt.Sprintf("%s:%v", name, value)
		validator.mu.RLock()
		entry, exists := validator.cache[cacheKey]
		validator.mu.RUnlock()
		if exists && time.Now().Before(entry.ExpiresAt) {
			return entry.Result
		}
	}
	// Apply debouncing if configured
	if validator.Debounce > 0 {
		select {
		case <-time.After(validator.Debounce):
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	// Perform validation
	result := validator.Validate(ctx, value, params)
	// Cache result if enabled
	if validator.Cache {
		cacheKey := fmt.Sprintf("%s:%v", name, value)
		validator.mu.Lock()
		validator.cache[cacheKey] = &CacheEntry{
			Result:    result,
			ExpiresAt: time.Now().Add(validator.CacheTTL),
		}
		validator.mu.Unlock()
	}
	return result
}

// UnregisterAsync removes an async validator and stops its cleanup goroutine
func (r *ValidationRegistry) UnregisterAsync(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	validator, exists := r.asyncValidators[name]
	if !exists {
		return NewValidationError("async_validator_not_found", fmt.Sprintf("async validator %s not found", name))
	}
	// Stop cleanup goroutine
	validator.Stop()
	delete(r.asyncValidators, name)
	return nil
}

// Register adds a cross-field validator
func (r *CrossFieldValidationRegistry) Register(validator *CrossFieldValidator) error {
	if validator == nil {
		return NewValidationError("cross_validator", "cross-field validator is required")
	}
	if validator.Name == "" {
		return NewValidationError("cross_validator_name", "cross-field validator name is required")
	}
	if len(validator.Fields) == 0 {
		return NewValidationError("cross_validator_fields", "cross-field validator must specify fields")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.validators[validator.Name] = validator
	return nil
}

// Validate executes a cross-field validator
func (r *CrossFieldValidationRegistry) Validate(ctx context.Context, name string, values map[string]any) error {
	r.mu.RLock()
	validator, exists := r.validators[name]
	r.mu.RUnlock()
	if !exists {
		return NewValidationError("cross_validator_not_found", fmt.Sprintf("cross-field validator %s not found", name))
	}
	// Check if all required fields are present
	fieldValues := make(map[string]any)
	for _, fieldName := range validator.Fields {
		if value, exists := values[fieldName]; exists {
			fieldValues[fieldName] = value
		}
	}
	// If condition is specified, evaluate it first
	if validator.Condition != nil {
		evalCtx := condition.NewEvalContext(fieldValues, condition.DefaultEvalOptions())
		shouldValidate, err := r.evaluator.Evaluate(ctx, validator.Condition, evalCtx)
		if err != nil {
			return WrapError(err, "condition_evaluation_failed", "failed to evaluate condition")
		}
		// Only validate if condition is true
		if !shouldValidate {
			return nil
		}
	}
	return validator.Validate(ctx, fieldValues)
}

// Add adds a validator to the chain
func (vc *ValidatorChain) Add(validator ValidatorFunc) *ValidatorChain {
	vc.validators = append(vc.validators, validator)
	return vc
}

// Validate executes all validators in the chain
func (vc *ValidatorChain) Validate(ctx context.Context, value any, params map[string]any) error {
	for _, validator := range vc.validators {
		if err := validator(ctx, value, params); err != nil {
			return err
		}
	}
	return nil
}

// Stop stops the cache cleanup goroutine
func (av *AsyncValidator) Stop() {
	if av.cleanup != nil {
		av.cleanup()
	}
}

// startCacheCleanup starts a goroutine to clean up expired cache entries
func (av *AsyncValidator) startCacheCleanup(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				av.cleanupExpiredCache()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// cleanupExpiredCache removes expired entries from the cache
func (av *AsyncValidator) cleanupExpiredCache() {
	av.mu.Lock()
	defer av.mu.Unlock()
	now := time.Now()
	for key, entry := range av.cache {
		if now.After(entry.ExpiresAt) {
			delete(av.cache, key)
		}
	}
}

// WrapValidationError wraps an error with additional context
func WrapValidationError(err error, code, message string) error {
	return fmt.Errorf("%s (%s): %w", message, code, err)
}
