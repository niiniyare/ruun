package validation

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ValidationTrigger defines when validation should occur
type ValidationTrigger int

const (
	ValidationTriggerImmediate ValidationTrigger = iota // Immediate validation (for required fields)
	ValidationTriggerDebounced                          // Debounced validation (for patterns, async checks)
	ValidationTriggerOnBlur                             // On-blur validation
	ValidationTriggerOnSubmit                           // On-submit validation only
)

// UIValidationOrchestrator coordinates validation timing and UI state
// REUSES existing schema validation infrastructure completely
type UIValidationOrchestrator struct {
	// REUSE existing schema validation (from pkg/schema/validate/)
	schema    *schema.Schema
	validator schema.RuntimeValidator // Existing validator interface

	// NEW: UI-specific timing and state control
	debouncer *FieldDebouncer
	stateMgr  *ValidationStateManager
	eventBus  *ValidationEventBus

	// Configuration
	config *ValidationConfig

	// Thread safety
	mu sync.RWMutex
}

// ValidationConfig holds UI validation configuration
type ValidationConfig struct {
	// Timing configuration
	ImmediateFields    []string      // Fields that validate immediately
	DebounceDelay      time.Duration // Default debounce delay
	OnBlurDelay        time.Duration // Delay for onBlur validation
	AsyncTimeoutDelay  time.Duration // Timeout for async validation

	// Behavior
	ValidateOnChange   bool // Enable change-based validation
	ValidateOnBlur     bool // Enable blur-based validation
	ShowErrorsOnlyOnBlur bool // Only show errors after blur
	ClearErrorsOnChange bool // Clear errors immediately on change

	// Client-side validation
	EnableClientRules  bool // Enable client-side rule extraction
}

// NewUIValidationOrchestrator creates a new UI validation orchestrator
func NewUIValidationOrchestrator(
	schema *schema.Schema,
	validator schema.RuntimeValidator,
	stateMgr *ValidationStateManager,
) *UIValidationOrchestrator {
	config := &ValidationConfig{
		DebounceDelay:        300 * time.Millisecond,
		OnBlurDelay:          100 * time.Millisecond,
		AsyncTimeoutDelay:    5 * time.Second,
		ValidateOnChange:     true,
		ValidateOnBlur:       true,
		ShowErrorsOnlyOnBlur: false,
		ClearErrorsOnChange:  true,
		EnableClientRules:    true,
	}

	orchestrator := &UIValidationOrchestrator{
		schema:    schema,
		validator: validator,
		stateMgr:  stateMgr,
		config:    config,
		debouncer: NewFieldDebouncer(),
		eventBus:  NewValidationEventBus(),
	}

	// Connect event bus to state manager for real-time updates
	orchestrator.connectEventHandlers()

	return orchestrator
}

// WithConfig sets custom validation configuration
func (o *UIValidationOrchestrator) WithConfig(config *ValidationConfig) *UIValidationOrchestrator {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.config = config
	return o
}

// ValidateFieldWithUI orchestrates validation with UI timing control
// This is the main entry point for UI-driven validation
func (o *UIValidationOrchestrator) ValidateFieldWithUI(
	ctx context.Context,
	fieldName string,
	value any,
	trigger ValidationTrigger,
) error {
	o.mu.RLock()
	schema := o.schema
	validator := o.validator
	config := o.config
	o.mu.RUnlock()

	if validator == nil {
		return fmt.Errorf("no validator configured")
	}

	// REUSE existing field validation from schema
	field, exists := schema.GetField(fieldName)
	if !exists {
		return fmt.Errorf("field %s not found", fieldName)
	}

	// Set UI state: validating
	o.stateMgr.SetValidationState(fieldName, ValidationStateValidating)
	o.eventBus.EmitValidationStart(fieldName)

	// Use existing schema validation based on trigger
	switch trigger {
	case ValidationTriggerImmediate:
		return o.validateImmediate(ctx, &field, value, config)
	case ValidationTriggerDebounced:
		return o.validateDebounced(ctx, &field, value, config)
	case ValidationTriggerOnBlur:
		return o.validateOnBlur(ctx, &field, value, config)
	case ValidationTriggerOnSubmit:
		return o.validateOnSubmit(ctx, &field, value, config)
	default:
		return o.validateImmediate(ctx, &field, value, config)
	}
}

// validateImmediate performs immediate validation using existing schema validation
func (o *UIValidationOrchestrator) validateImmediate(
	ctx context.Context,
	field *schema.Field,
	value any,
	config *ValidationConfig,
) error {
	// REUSE: Use existing schema RuntimeValidator interface
	allData := o.stateMgr.GetAllValues()
	errors := o.validator.ValidateField(ctx, field, value, allData)

	fieldName := field.Name

	if len(errors) > 0 {
		// Handle validation errors
		o.handleValidationErrors(fieldName, errors)
		
		// Emit validation error event
		o.eventBus.EmitValidationError(fieldName, errors)
		
		// Update UI state
		o.stateMgr.SetValidationState(fieldName, ValidationStateInvalid)
		
		return fmt.Errorf("validation failed: %v", errors)
	}

	// Validation passed
	o.handleValidationSuccess(fieldName)
	o.eventBus.EmitValidationSuccess(fieldName)
	o.stateMgr.SetValidationState(fieldName, ValidationStateValid)
	
	return nil
}

// validateDebounced performs debounced validation for better UX
func (o *UIValidationOrchestrator) validateDebounced(
	ctx context.Context,
	field *schema.Field,
	value any,
	config *ValidationConfig,
) error {
	fieldName := field.Name

	// Cancel any pending validation for this field
	o.debouncer.CancelValidation(fieldName)

	// Schedule debounced validation
	validationFunc := func() {
		if err := o.validateImmediate(ctx, field, value, config); err != nil {
			// Error already handled in validateImmediate
			return
		}
	}

	o.debouncer.ScheduleValidation(fieldName, config.DebounceDelay, validationFunc)
	
	// Clear errors immediately if configured
	if config.ClearErrorsOnChange {
		o.stateMgr.ClearFieldErrors(fieldName)
		o.stateMgr.SetValidationState(fieldName, ValidationStateIdle)
	}

	return nil
}

// validateOnBlur performs validation on field blur event
func (o *UIValidationOrchestrator) validateOnBlur(
	ctx context.Context,
	field *schema.Field,
	value any,
	config *ValidationConfig,
) error {
	// Short delay for onBlur validation to improve UX
	time.Sleep(config.OnBlurDelay)
	
	return o.validateImmediate(ctx, field, value, config)
}

// validateOnSubmit performs validation only on form submission
func (o *UIValidationOrchestrator) validateOnSubmit(
	ctx context.Context,
	field *schema.Field,
	value any,
	config *ValidationConfig,
) error {
	// Only validate during form submission - no immediate feedback
	o.stateMgr.SetValidationState(field.Name, ValidationStateIdle)
	return nil
}

// ValidateAllFieldsWithUI validates all fields using UI orchestration
func (o *UIValidationOrchestrator) ValidateAllFieldsWithUI(ctx context.Context) map[string][]string {
	o.mu.RLock()
	schema := o.schema
	validator := o.validator
	o.mu.RUnlock()

	if validator == nil {
		return make(map[string][]string)
	}

	// REUSE: Use existing schema RuntimeValidator.ValidateAllFields
	allData := o.stateMgr.GetAllValues()
	allErrors := validator.ValidateAllFields(ctx, schema, allData)

	// Update UI state for all fields
	for _, field := range schema.Fields {
		fieldName := field.Name
		
		if errors, hasErrors := allErrors[fieldName]; hasErrors && len(errors) > 0 {
			o.handleValidationErrors(fieldName, errors)
			o.stateMgr.SetValidationState(fieldName, ValidationStateInvalid)
			o.eventBus.EmitValidationError(fieldName, errors)
		} else {
			o.handleValidationSuccess(fieldName)
			o.stateMgr.SetValidationState(fieldName, ValidationStateValid)
			o.eventBus.EmitValidationSuccess(fieldName)
		}
	}

	return allErrors
}

// GetValidationStateForField returns the current validation state for a field
func (o *UIValidationOrchestrator) GetValidationStateForField(fieldName string) ValidationState {
	return o.stateMgr.GetValidationState(fieldName)
}

// IsValidationInProgress checks if any field is currently being validated
func (o *UIValidationOrchestrator) IsValidationInProgress() bool {
	return o.stateMgr.HasFieldsInState(ValidationStateValidating)
}

// GetPendingValidations returns fields with pending validations
func (o *UIValidationOrchestrator) GetPendingValidations() []string {
	return o.debouncer.GetPendingFields()
}

// connectEventHandlers connects the event bus to the state manager
func (o *UIValidationOrchestrator) connectEventHandlers() {
	// Connect validation events to state updates
	o.eventBus.AddListener("*", func(event ValidationEvent) {
		// Update validation timestamps
		o.stateMgr.SetValidationTime(event.FieldName, event.Timestamp)
		
		// Additional event-driven state updates can be added here
	})
}

// handleValidationErrors processes validation errors from schema validation
func (o *UIValidationOrchestrator) handleValidationErrors(fieldName string, errors []string) {
	// Update state manager with errors
	o.stateMgr.SetFieldErrors(fieldName, errors)
	
	// Mark field as touched since we validated it
	o.stateMgr.SetFieldTouched(fieldName, true)
}

// handleValidationSuccess processes successful validation
func (o *UIValidationOrchestrator) handleValidationSuccess(fieldName string) {
	// Clear any existing errors
	o.stateMgr.ClearFieldErrors(fieldName)
	
	// Mark field as touched
	o.stateMgr.SetFieldTouched(fieldName, true)
}

// OnFieldChange handles field value changes with appropriate validation trigger
func (o *UIValidationOrchestrator) OnFieldChange(
	ctx context.Context,
	fieldName string,
	newValue any,
) error {
	o.mu.RLock()
	config := o.config
	o.mu.RUnlock()

	if !config.ValidateOnChange {
		return nil // Validation on change disabled
	}

	// Determine validation trigger based on field and configuration
	trigger := o.determineValidationTrigger(fieldName)
	
	return o.ValidateFieldWithUI(ctx, fieldName, newValue, trigger)
}

// OnFieldBlur handles field blur events
func (o *UIValidationOrchestrator) OnFieldBlur(
	ctx context.Context,
	fieldName string,
	value any,
) error {
	o.mu.RLock()
	config := o.config
	o.mu.RUnlock()

	if !config.ValidateOnBlur {
		return nil // Validation on blur disabled
	}

	return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerOnBlur)
}

// determineValidationTrigger determines the appropriate validation trigger for a field
func (o *UIValidationOrchestrator) determineValidationTrigger(fieldName string) ValidationTrigger {
	// Check if field should validate immediately
	for _, immediatField := range o.config.ImmediateFields {
		if immediatField == fieldName {
			return ValidationTriggerImmediate
		}
	}

	// Get field from schema to check validation complexity
	field, exists := o.schema.GetField(fieldName)
	if !exists {
		return ValidationTriggerDebounced // Safe default
	}

	// Use immediate validation for simple checks
	if field.Validation != nil {
		// Required fields - immediate feedback
		if field.Required {
			return ValidationTriggerImmediate
		}
		
		// Complex validation (async, business rules) - debounced
		if field.Validation.Unique || field.Validation.Custom != "" {
			return ValidationTriggerDebounced
		}
	}

	// Default to debounced for better UX
	return ValidationTriggerDebounced
}

// Cleanup cancels all pending validations
func (o *UIValidationOrchestrator) Cleanup() {
	o.debouncer.CancelAllValidations()
}

// ═══════════════════════════════════════════════════════════════════════════
// HELPER METHODS FOR INTERACTION HANDLERS
// ═══════════════════════════════════════════════════════════════════════════

// buildHTMXValidationAttributes builds HTMX-specific attributes for validation responses
func (o *UIValidationOrchestrator) buildHTMXValidationAttributes(fieldName string, state ValidationState) map[string]string {
	attributes := make(map[string]string)
	
	// Set validation state attributes
	attributes["hx-trigger"] = o.getHTMXTriggerForField(fieldName)
	attributes["data-validation-state"] = string(state)
	
	// Add debounce timing for async validation
	if o.HasAsyncValidation(fieldName) {
		debounceTime := o.GetDebounceTimeForField(fieldName)
		attributes["hx-trigger"] = fmt.Sprintf("keyup changed delay:%dms", debounceTime)
	}
	
	// Add validation endpoint
	attributes["hx-post"] = fmt.Sprintf("/api/validate/field/%s", fieldName)
	attributes["hx-target"] = fmt.Sprintf("#%s-validation", fieldName)
	attributes["hx-swap"] = "innerHTML"
	
	return attributes
}

// getHTMXTriggerForField determines optimal HTMX trigger for field
func (o *UIValidationOrchestrator) getHTMXTriggerForField(fieldName string) string {
	trigger := o.determineValidationTrigger(fieldName)
	
	switch trigger {
	case ValidationTriggerImmediate:
		return "keyup changed"
	case ValidationTriggerDebounced:
		debounceTime := o.GetDebounceTimeForField(fieldName)
		return fmt.Sprintf("keyup changed delay:%dms", debounceTime)
	case ValidationTriggerOnBlur:
		return "blur"
	default:
		return "keyup changed delay:300ms"
	}
}

// countTotalErrors counts total validation errors across all fields
func (o *UIValidationOrchestrator) countTotalErrors(fieldErrors map[string][]string) int {
	totalErrors := 0
	for _, errors := range fieldErrors {
		totalErrors += len(errors)
	}
	return totalErrors
}

// getTouchedFields returns list of fields that have been touched
func (o *UIValidationOrchestrator) getTouchedFields() []string {
	if o.stateMgr == nil {
		return []string{}
	}
	
	var touchedFields []string
	for _, field := range o.schema.Fields {
		if o.stateMgr.IsFieldTouched(field.Name) {
			touchedFields = append(touchedFields, field.Name)
		}
	}
	return touchedFields
}

// getDirtyFields returns list of fields that have been modified
func (o *UIValidationOrchestrator) getDirtyFields() []string {
	if o.stateMgr == nil {
		return []string{}
	}
	
	var dirtyFields []string
	for _, field := range o.schema.Fields {
		if o.isFieldDirty(field.Name) {
			dirtyFields = append(dirtyFields, field.Name)
		}
	}
	return dirtyFields
}

// requiresServerValidation checks if field requires server-side validation
func (o *UIValidationOrchestrator) requiresServerValidation(field *schema.Field) bool {
	if field.Validation == nil {
		return false
	}
	
	// Async validation rules require server validation
	if field.Validation.Unique || field.Validation.Custom != "" {
		return true
	}
	
	// Complex business rules require server validation
	if field.Validation.Custom != "" {
		return true
	}
	
	return false
}

// determineValidationStrategy determines the optimal validation strategy
func (o *UIValidationOrchestrator) determineValidationStrategy(field *schema.Field, context *RealTimeValidationContext) ValidationStrategy {
	// Immediate validation for critical fields
	if field.Required && context.IsFirstInput {
		return ValidationStrategyImmediate
	}
	
	// Client-first validation for simple rules
	if !o.requiresServerValidation(field) && !context.IsUserTyping {
		return ValidationStrategyClientFirst
	}
	
	// On-blur validation for complex fields while typing
	if context.IsUserTyping && o.HasAsyncValidation(field.Name) {
		return ValidationStrategyOnBlur
	}
	
	// Default to debounced validation
	return ValidationStrategyDebounced
}

// triggerClientSideValidation triggers client-side validation for a field
func (o *UIValidationOrchestrator) triggerClientSideValidation(ctx context.Context, fieldName string, value any, context *RealTimeValidationContext) error {
	// This would integrate with client-side validation libraries
	// For now, just proceed with server validation
	return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerDebounced)
}

// ═══════════════════════════════════════════════════════════════════════════
// HELPER METHODS FOR STATE MANAGEMENT DELEGATION
// ═══════════════════════════════════════════════════════════════════════════

// getFieldValue gets field value from state manager
func (o *UIValidationOrchestrator) getFieldValue(fieldName string) any {
	if o.stateMgr != nil && o.stateMgr.stateManager != nil {
		return o.stateMgr.stateManager.GetValue(fieldName)
	}
	return nil
}

// setFieldValue sets field value in state manager
func (o *UIValidationOrchestrator) setFieldValue(fieldName string, value any) {
	if o.stateMgr != nil && o.stateMgr.stateManager != nil {
		o.stateMgr.stateManager.SetValue(fieldName, value)
	}
}

// isFieldDirty checks if field is dirty via state manager
func (o *UIValidationOrchestrator) isFieldDirty(fieldName string) bool {
	if o.stateMgr != nil && o.stateMgr.stateManager != nil {
		return o.stateMgr.stateManager.IsFieldDirty(fieldName)
	}
	return false
}

// setFieldDirty marks field as dirty via state manager
func (o *UIValidationOrchestrator) setFieldDirty(fieldName string, dirty bool) {
	if o.stateMgr != nil && o.stateMgr.stateManager != nil {
		o.stateMgr.stateManager.SetFieldDirty(fieldName, dirty)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// ADDITIONAL ORCHESTRATOR METHODS (for renderer integration)
// ═══════════════════════════════════════════════════════════════════════════

// IsFieldValidating checks if a specific field is currently being validated
func (o *UIValidationOrchestrator) IsFieldValidating(fieldName string) bool {
	return o.stateMgr.GetValidationState(fieldName) == ValidationStateValidating
}

// GetFieldErrors returns current validation errors for a field
func (o *UIValidationOrchestrator) GetFieldErrors(fieldName string) []string {
	return o.stateMgr.GetFieldErrors(fieldName)
}

// GetContextForValidation extracts validation context including theme information
func (o *UIValidationOrchestrator) GetContextForValidation(field *schema.Field) map[string]any {
	validationContext := make(map[string]any)
	
	if field == nil {
		return validationContext
	}
	
	// Include validation state
	validationState := o.GetValidationStateForField(field.Name)
	validationContext["validationState"] = validationState
	validationContext["isValidating"] = o.IsFieldValidating(field.Name)
	validationContext["hasAsyncValidation"] = o.HasAsyncValidation(field.Name)
	
	// Include field errors
	errors := o.GetFieldErrors(field.Name)
	validationContext["errors"] = errors
	validationContext["hasErrors"] = len(errors) > 0
	
	// Include timing information
	if o.stateMgr != nil {
		validationTime := o.stateMgr.GetValidationTime(field.Name)
		validationContext["lastValidated"] = validationTime
	}
	
	// Include debounce time for this field
	validationContext["debounceTime"] = o.GetDebounceTimeForField(field.Name)
	
	return validationContext
}

// HasAsyncValidation checks if a field has async validation rules
func (o *UIValidationOrchestrator) HasAsyncValidation(fieldName string) bool {
	field, exists := o.schema.GetField(fieldName)
	if !exists {
		return false
	}
	
	// Check for validation rules that require async processing
	if field.Validation != nil {
		// Uniqueness checks require database lookups
		if field.Validation.Unique {
			return true
		}
		
		// Custom validation rules might require external calls
		if field.Validation.Custom != "" {
			return true
		}
	}
	
	return false
}

// GetDebounceTimeForField returns the debounce time configured for a field
func (o *UIValidationOrchestrator) GetDebounceTimeForField(fieldName string) int {
	// Check if field has specific debounce configuration
	field, exists := o.schema.GetField(fieldName)
	if exists && field.Config != nil {
		if debounce, ok := field.Config["debounce"].(int); ok {
			return debounce
		}
	}
	
	// Return default debounce time in milliseconds
	return int(o.config.DebounceDelay / time.Millisecond)
}

// ═══════════════════════════════════════════════════════════════════════════
// THEME-AWARE VALIDATION INTEGRATION (NEW)
// ═══════════════════════════════════════════════════════════════════════════

// ThemeAwareOrchestrator enhances the orchestrator with theme-aware validation
type ThemeAwareOrchestrator struct {
	*UIValidationOrchestrator
	
	// Theme-aware extensions
	themeManager       ThemeManager
	validationRenderer ValidationRenderer
	eventEmitter       ValidationEventEmitter
	
	// Theme integration config
	themeConfig *ThemeValidationConfig
}

// ThemeValidationConfig holds theme-specific validation configuration
type ThemeValidationConfig struct {
	// Theme coordination
	AutoDetectDarkMode     bool          // Auto-detect dark mode from system
	ThemeTransitionDelay   time.Duration // Delay for theme transitions
	ValidateThemeContrast  bool          // Validate color contrast for accessibility
	
	// Validation UI theming
	EnableValidationTokens bool // Enable token-based validation styling
	ThemeIDField          string // Field name to use for theme context
	
	// Real-time theme updates
	EnableThemeEvents      bool // Enable theme change events
	ThemeEventDebounce     time.Duration // Debounce theme events
}

// ThemeManager interface for theme coordination
type ThemeManager interface {
	GetActiveTheme(context string) string
	IsDarkModeEnabled(context string) bool
	GetThemeTokens(themeID string, validation ValidationState) map[string]string
	SubscribeToThemeChanges(callback func(themeID string, darkMode bool))
}

// ValidationRenderer interface for theme-aware validation rendering
type ValidationRenderer interface {
	RenderValidationStatus(fieldName string, state ValidationState, tokens map[string]string) string
	RenderValidationMessages(fieldName string, errors []string, tokens map[string]string) string
	UpdateFieldTokens(fieldName string, tokens map[string]string)
}

// ValidationEventEmitter interface for theme-aware validation events
type ValidationEventEmitter interface {
	EmitThemeValidationEvent(fieldName string, state ValidationState, themeContext map[string]any)
	EmitValidationTokenUpdate(fieldName string, tokens map[string]string)
}

// NewThemeAwareOrchestrator creates a new theme-aware validation orchestrator
func NewThemeAwareOrchestrator(
	baseOrchestrator *UIValidationOrchestrator,
	themeManager ThemeManager,
	validationRenderer ValidationRenderer,
	eventEmitter ValidationEventEmitter,
) *ThemeAwareOrchestrator {
	
	config := &ThemeValidationConfig{
		AutoDetectDarkMode:     true,
		ThemeTransitionDelay:   150 * time.Millisecond,
		ValidateThemeContrast:  true,
		EnableValidationTokens: true,
		ThemeIDField:          "theme",
		EnableThemeEvents:      true,
		ThemeEventDebounce:     100 * time.Millisecond,
	}
	
	themeOrchestrator := &ThemeAwareOrchestrator{
		UIValidationOrchestrator: baseOrchestrator,
		themeManager:             themeManager,
		validationRenderer:       validationRenderer,
		eventEmitter:             eventEmitter,
		themeConfig:              config,
	}
	
	// Subscribe to theme changes
	if config.EnableThemeEvents {
		themeOrchestrator.subscribeToThemeEvents()
	}
	
	return themeOrchestrator
}

// ValidateFieldWithTheme performs theme-aware field validation
func (o *ThemeAwareOrchestrator) ValidateFieldWithTheme(
	ctx context.Context,
	fieldName string,
	value any,
	trigger ValidationTrigger,
	themeContext map[string]any,
) error {
	
	// Perform standard validation
	err := o.ValidateFieldWithUI(ctx, fieldName, value, trigger)
	
	// Apply theme-aware enhancements
	if o.themeConfig.EnableValidationTokens {
		o.applyValidationTheme(fieldName, themeContext)
	}
	
	// Emit theme-aware validation event
	if o.themeConfig.EnableThemeEvents {
		validationState := o.GetValidationStateForField(fieldName)
		o.eventEmitter.EmitThemeValidationEvent(fieldName, validationState, themeContext)
	}
	
	return err
}

// applyValidationTheme applies theme-aware styling to validation UI
func (o *ThemeAwareOrchestrator) applyValidationTheme(fieldName string, themeContext map[string]any) {
	// Determine active theme
	activeTheme := o.getActiveThemeForField(fieldName, themeContext)
	
	// Check if dark mode is enabled
	isDarkMode := o.isDarkModeForContext(themeContext)
	
	// Get validation state
	validationState := o.GetValidationStateForField(fieldName)
	
	// Resolve validation tokens for current theme and state
	tokens := o.themeManager.GetThemeTokens(activeTheme, validationState)
	
	// Apply dark mode overrides if needed
	if isDarkMode {
		darkTokens := o.getDarkModeTokenOverrides(activeTheme, validationState)
		for key, value := range darkTokens {
			tokens[key] = value
		}
	}
	
	// Update field tokens through renderer
	o.validationRenderer.UpdateFieldTokens(fieldName, tokens)
	
	// Emit token update event
	o.eventEmitter.EmitValidationTokenUpdate(fieldName, tokens)
}

// getActiveThemeForField determines the active theme for a field
func (o *ThemeAwareOrchestrator) getActiveThemeForField(fieldName string, themeContext map[string]any) string {
	// Check for field-specific theme in context
	if themeID, exists := themeContext["themeId"]; exists {
		if theme, ok := themeID.(string); ok {
			return theme
		}
	}
	
	// Check for theme ID field in form data
	if o.themeConfig.ThemeIDField != "" {
		if themeValue := o.getFieldValue(o.themeConfig.ThemeIDField); themeValue != nil {
			if theme, ok := themeValue.(string); ok {
				return theme
			}
		}
	}
	
	// Use theme manager to get active theme
	return o.themeManager.GetActiveTheme(fieldName)
}

// isDarkModeForContext determines if dark mode is enabled for the current context
func (o *ThemeAwareOrchestrator) isDarkModeForContext(themeContext map[string]any) bool {
	// Check for explicit dark mode setting in context
	if darkMode, exists := themeContext["darkMode"]; exists {
		if dark, ok := darkMode.(bool); ok {
			return dark
		}
	}
	
	// Auto-detect dark mode if enabled
	if o.themeConfig.AutoDetectDarkMode {
		return o.themeManager.IsDarkModeEnabled("")
	}
	
	return false
}

// getDarkModeTokenOverrides gets dark mode token overrides for validation
func (o *ThemeAwareOrchestrator) getDarkModeTokenOverrides(themeID string, state ValidationState) map[string]string {
	// This would integrate with the theme manager to get dark mode overrides
	// For now, return basic dark mode overrides
	return map[string]string{
		"background": "#111827",
		"text":       "#f9fafb",
		"border":     "#374151",
	}
}

// subscribeToThemeEvents subscribes to theme change events
func (o *ThemeAwareOrchestrator) subscribeToThemeEvents() {
	o.themeManager.SubscribeToThemeChanges(func(themeID string, darkMode bool) {
		// Apply theme changes to all fields with debouncing
		o.handleThemeChange(themeID, darkMode)
	})
}

// handleThemeChange handles theme change events
func (o *ThemeAwareOrchestrator) handleThemeChange(themeID string, darkMode bool) {
	// Debounce theme changes to avoid excessive updates
	time.Sleep(o.themeConfig.ThemeEventDebounce)
	
	// Update all field tokens with new theme
	for _, field := range o.schema.Fields {
		themeContext := map[string]any{
			"themeId":  themeID,
			"darkMode": darkMode,
		}
		o.applyValidationTheme(field.Name, themeContext)
	}
}

// OnThemeChange handles explicit theme changes
func (o *ThemeAwareOrchestrator) OnThemeChange(
	ctx context.Context,
	themeID string,
	darkMode bool,
) error {
	// Update theme for all validated fields
	for _, field := range o.schema.Fields {
		fieldName := field.Name
		
		// Only update fields that have been validated
		if o.stateMgr.IsFieldTouched(fieldName) {
			themeContext := map[string]any{
				"themeId":  themeID,
				"darkMode": darkMode,
			}
			o.applyValidationTheme(fieldName, themeContext)
		}
	}
	
	return nil
}

// WithThemeConfig sets theme validation configuration
func (o *ThemeAwareOrchestrator) WithThemeConfig(config *ThemeValidationConfig) *ThemeAwareOrchestrator {
	o.themeConfig = config
	return o
}

// ═══════════════════════════════════════════════════════════════════════════
// HTMX & ALPINE.JS VALIDATION INTERACTION HANDLERS (NEW)
// ═══════════════════════════════════════════════════════════════════════════

// HandleHTMXFieldValidation handles HTMX-triggered field validation requests
// This is called from HTMX endpoints to validate individual fields
func (o *UIValidationOrchestrator) HandleHTMXFieldValidation(
	ctx context.Context,
	fieldName string,
	value any,
	triggerType string,
) (*ValidationResponse, error) {
	// Determine validation trigger from HTMX trigger type
	var trigger ValidationTrigger
	switch triggerType {
	case "input", "keyup":
		trigger = ValidationTriggerDebounced
	case "blur":
		trigger = ValidationTriggerOnBlur
	case "change":
		trigger = ValidationTriggerImmediate
	case "submit":
		trigger = ValidationTriggerOnSubmit
	default:
		trigger = ValidationTriggerDebounced
	}
	
	// Validate field using UI orchestration
	err := o.ValidateFieldWithUI(ctx, fieldName, value, trigger)
	
	// Build validation response for HTMX
	response := &ValidationResponse{
		FieldName:       fieldName,
		ValidationState: o.GetValidationStateForField(fieldName),
		IsValid:         err == nil,
		Errors:          o.GetFieldErrors(fieldName),
		IsValidating:    o.IsFieldValidating(fieldName),
		Trigger:         triggerType,
		Timestamp:       time.Now(),
	}
	
	// Add HTMX-specific response headers and attributes
	response.HTMXAttributes = o.buildHTMXValidationAttributes(fieldName, response.ValidationState)
	
	return response, err
}

// HandleAlpineFieldValidation handles Alpine.js-triggered validation events
// This is called from Alpine.js event handlers for real-time validation
func (o *UIValidationOrchestrator) HandleAlpineFieldValidation(
	ctx context.Context,
	fieldName string,
	value any,
	eventType string,
	alpineData map[string]any,
) (*AlpineValidationResponse, error) {
	// Map Alpine.js event types to validation triggers
	var trigger ValidationTrigger
	switch eventType {
	case "@input", "@keyup":
		trigger = ValidationTriggerDebounced
	case "@blur", "@focusout":
		trigger = ValidationTriggerOnBlur
	case "@change":
		trigger = ValidationTriggerImmediate
	default:
		trigger = ValidationTriggerDebounced
	}
	
	// Update field value in state manager if available
	if o.stateMgr != nil {
		o.setFieldValue(fieldName, value)
	}
	
	// Validate field
	err := o.ValidateFieldWithUI(ctx, fieldName, value, trigger)
	
	// Build Alpine.js validation response
	response := &AlpineValidationResponse{
		FieldName:       fieldName,
		ValidationState: string(o.GetValidationStateForField(fieldName)),
		IsValid:         err == nil,
		Errors:          o.GetFieldErrors(fieldName),
		IsValidating:    o.IsFieldValidating(fieldName),
		EventType:       eventType,
		Timestamp:       time.Now().Unix(),
		
		// Alpine.js specific data
		AlpineData: map[string]any{
			"validationState": string(o.GetValidationStateForField(fieldName)),
			"hasErrors":       len(o.GetFieldErrors(fieldName)) > 0,
			"isValidating":    o.IsFieldValidating(fieldName),
			"debounceTime":    o.GetDebounceTimeForField(fieldName),
			"hasAsync":        o.HasAsyncValidation(fieldName),
		},
	}
	
	// Merge with provided Alpine data
	for key, value := range alpineData {
		response.AlpineData[key] = value
	}
	
	return response, err
}

// HandleFormSubmission validates all fields on form submission
// This coordinates validation for the entire form before submission
func (o *UIValidationOrchestrator) HandleFormSubmission(
	ctx context.Context,
	formData map[string]any,
	submissionContext map[string]any,
) (*FormValidationResponse, error) {
	// Update all field values in state manager
	if o.stateMgr != nil {
		for fieldName, value := range formData {
			o.setFieldValue(fieldName, value)
		}
	}
	
	// Validate all fields
	allErrors := o.ValidateAllFieldsWithUI(ctx)
	
	// Build form validation response
	isValid := len(allErrors) == 0
	response := &FormValidationResponse{
		IsValid:      isValid,
		FieldErrors:  allErrors,
		ErrorCount:   o.countTotalErrors(allErrors),
		IsValidating: o.IsValidationInProgress(),
		Timestamp:    time.Now(),
		
		// Form-level validation metadata
		ValidationMeta: map[string]any{
			"submissionContext": submissionContext,
			"pendingFields":     o.GetPendingValidations(),
			"touchedFields":     o.getTouchedFields(),
			"dirtyFields":       o.getDirtyFields(),
		},
	}
	
	if !isValid {
		return response, fmt.Errorf("form validation failed: %d field(s) have errors", response.ErrorCount)
	}
	
	return response, nil
}

// HandleClientSideValidation processes client-side validation results
// This integrates client-side validation with server-side validation orchestration
func (o *UIValidationOrchestrator) HandleClientSideValidation(
	ctx context.Context,
	fieldName string,
	clientResult *ClientValidationResult,
) error {
	// Update field state with client validation results
	if clientResult.IsValid {
		// Client validation passed - proceed with server validation if needed
		field, exists := o.schema.GetField(fieldName)
		if !exists {
			return fmt.Errorf("field %s not found", fieldName)
		}
		
		// Check if server validation is required
		if o.requiresServerValidation(&field) {
			return o.ValidateFieldWithUI(ctx, fieldName, clientResult.Value, ValidationTriggerDebounced)
		} else {
			// Client validation sufficient - mark as valid
			o.stateMgr.SetValidationState(fieldName, ValidationStateValid)
			o.stateMgr.ClearFieldErrors(fieldName)
			o.eventBus.EmitValidationSuccess(fieldName)
		}
	} else {
		// Client validation failed - show errors immediately
		o.stateMgr.SetFieldErrors(fieldName, clientResult.Errors)
		o.stateMgr.SetValidationState(fieldName, ValidationStateInvalid)
		o.eventBus.EmitValidationError(fieldName, clientResult.Errors)
	}
	
	return nil
}

// HandleRealTimeValidation manages real-time validation with intelligent timing
// This provides smooth UX by coordinating validation timing and feedback
func (o *UIValidationOrchestrator) HandleRealTimeValidation(
	ctx context.Context,
	fieldName string,
	value any,
	validationContext *RealTimeValidationContext,
) error {
	// Get field validation configuration
	field, exists := o.schema.GetField(fieldName)
	if !exists {
		return fmt.Errorf("field %s not found", fieldName)
	}
	
	// Determine optimal validation strategy based on context
	strategy := o.determineValidationStrategy(&field, validationContext)
	
	// Apply validation strategy
	switch strategy {
	case ValidationStrategyImmediate:
		return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerImmediate)
		
	case ValidationStrategyDebounced:
		return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerDebounced)
		
	case ValidationStrategyOnBlur:
		// Set value but don't validate until blur
		if o.stateMgr != nil {
			o.setFieldValue(fieldName, value)
			o.setFieldDirty(fieldName, true)
		}
		return nil
		
	case ValidationStrategyClientFirst:
		// Trigger client-side validation first
		return o.triggerClientSideValidation(ctx, fieldName, value, validationContext)
		
	default:
		return o.ValidateFieldWithUI(ctx, fieldName, value, ValidationTriggerDebounced)
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// VALIDATION RESPONSE TYPES
// ═══════════════════════════════════════════════════════════════════════════

// ValidationResponse represents HTMX validation response
type ValidationResponse struct {
	FieldName       string            `json:"fieldName"`
	ValidationState ValidationState   `json:"validationState"`
	IsValid         bool              `json:"isValid"`
	Errors          []string          `json:"errors"`
	IsValidating    bool              `json:"isValidating"`
	Trigger         string            `json:"trigger"`
	Timestamp       time.Time         `json:"timestamp"`
	HTMXAttributes  map[string]string `json:"htmxAttributes"`
}

// AlpineValidationResponse represents Alpine.js validation response
type AlpineValidationResponse struct {
	FieldName       string         `json:"fieldName"`
	ValidationState string         `json:"validationState"`
	IsValid         bool           `json:"isValid"`
	Errors          []string       `json:"errors"`
	IsValidating    bool           `json:"isValidating"`
	EventType       string         `json:"eventType"`
	Timestamp       int64          `json:"timestamp"`
	AlpineData      map[string]any `json:"alpineData"`
}

// FormValidationResponse represents form-level validation response
type FormValidationResponse struct {
	IsValid        bool              `json:"isValid"`
	FieldErrors    map[string][]string `json:"fieldErrors"`
	ErrorCount     int               `json:"errorCount"`
	IsValidating   bool              `json:"isValidating"`
	Timestamp      time.Time         `json:"timestamp"`
	ValidationMeta map[string]any    `json:"validationMeta"`
}

// ClientValidationResult represents client-side validation result
type ClientValidationResult struct {
	FieldName string   `json:"fieldName"`
	Value     any      `json:"value"`
	IsValid   bool     `json:"isValid"`
	Errors    []string `json:"errors"`
	Rules     []string `json:"rules"`
}

// RealTimeValidationContext provides context for real-time validation decisions
type RealTimeValidationContext struct {
	IsFirstInput    bool              `json:"isFirstInput"`
	TimeSinceChange time.Duration     `json:"timeSinceChange"`
	IsUserTyping    bool              `json:"isUserTyping"`
	HasFocus        bool              `json:"hasFocus"`
	PreviousValue   any               `json:"previousValue"`
	ClientData      map[string]any    `json:"clientData"`
}

// ValidationStrategy defines validation timing strategy
type ValidationStrategy int

const (
	ValidationStrategyImmediate   ValidationStrategy = iota
	ValidationStrategyDebounced
	ValidationStrategyOnBlur
	ValidationStrategyClientFirst
)