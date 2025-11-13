package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/a-h/templ"
	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/validation"
)

// TemplateRenderer implements schema.RuntimeRenderer using Templ components
// This is the main bridge between the pure schema system and the Views layer
// EXTENDED with UI validation orchestration
type TemplateRenderer struct {
	registry     *ComponentRegistry
	mapper       *FieldMapper
	schema       *schema.Schema
	stateManager *StateManager                      // Form state management
	orchestrator *validation.UIValidationOrchestrator // UI validation orchestration
}

// NewTemplateRenderer creates a new template renderer with dependencies
func NewTemplateRenderer(registry *ComponentRegistry, mapper *FieldMapper, s *schema.Schema, stateManager *StateManager, orchestrator *validation.UIValidationOrchestrator) *TemplateRenderer {
	return &TemplateRenderer{
		registry:     registry,
		mapper:       mapper,
		schema:       s,
		stateManager: stateManager,
		orchestrator: orchestrator,
	}
}

// RenderField implements schema.RuntimeRenderer.RenderField
// Renders a single field with schema context and current state values
// ENHANCED with validation orchestration integration
func (r *TemplateRenderer) RenderField(ctx context.Context, field *schema.Field, value any, errors []string, touched, dirty bool) (string, error) {
	if field == nil {
		return "", fmt.Errorf("field cannot be nil")
	}

	// Use state manager for real validation states if available
	var realValue any = value
	var realErrors []string = errors
	var realTouched bool = touched
	var realDirty bool = dirty
	
	if r.stateManager != nil {
		// Get actual state from state manager
		realValue = r.stateManager.GetValue(field.Name)
		realErrors = r.stateManager.GetFieldErrors(field.Name)
		realTouched = r.stateManager.IsFieldTouched(field.Name)
		realDirty = r.stateManager.IsFieldDirty(field.Name)
		
		// Override with provided values if they are explicitly set
		if value != nil {
			realValue = value
		}
		if len(errors) > 0 {
			realErrors = errors
		}
	}

	// Convert schema field to FormField props with real validation state
	props, err := r.mapper.ConvertField(field, realValue, realErrors, realTouched, realDirty)
	if err != nil {
		return "", fmt.Errorf("failed to convert field %s: %w", field.Name, err)
	}
	
	// Enhance props with validation orchestration data
	r.enhancePropsWithValidation(field, &props)

	// Resolve component renderer for this field type
	renderer, err := r.registry.Resolve(field)
	if err != nil {
		return "", fmt.Errorf("failed to resolve renderer for field %s (type: %s): %w", field.Name, field.Type, err)
	}

	// Render using the resolved component
	html, err := renderer.Render(ctx, props)
	if err != nil {
		return "", fmt.Errorf("failed to render field %s: %w", field.Name, err)
	}

	return html, nil
}

// RenderForm implements schema.RuntimeRenderer.RenderForm
// Renders the complete form using schema and runtime state
func (r *TemplateRenderer) RenderForm(ctx context.Context, schema *schema.Schema, state map[string]any, errors map[string][]string) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("schema cannot be nil")
	}

	var formHTML strings.Builder

	// Build form opening tag with schema attributes
	formAttrs := r.buildFormAttributes(schema)
	formHTML.WriteString(fmt.Sprintf("<form %s>\n", formAttrs))

	// Render form fields based on layout or default order
	if schema.Layout != nil {
		html, err := r.renderFormWithLayout(ctx, schema, state, errors)
		if err != nil {
			return "", fmt.Errorf("failed to render form with layout: %w", err)
		}
		formHTML.WriteString(html)
	} else {
		html, err := r.renderFormFields(ctx, schema.Fields, state, errors)
		if err != nil {
			return "", fmt.Errorf("failed to render form fields: %w", err)
		}
		formHTML.WriteString(html)
	}

	// Render form actions
	if len(schema.Actions) > 0 {
		actionsHTML, err := r.renderFormActions(ctx, schema.Actions)
		if err != nil {
			return "", fmt.Errorf("failed to render form actions: %w", err)
		}
		formHTML.WriteString(actionsHTML)
	}

	formHTML.WriteString("</form>")

	return formHTML.String(), nil
}

// RenderAction implements schema.RuntimeRenderer.RenderAction
// Renders an action button/element
func (r *TemplateRenderer) RenderAction(ctx context.Context, action *schema.Action, enabled bool) (string, error) {
	if action == nil {
		return "", fmt.Errorf("action cannot be nil")
	}

	// Convert action to button props
	buttonProps := r.mapper.ConvertAction(action, enabled)

	// Render action as button
	// In actual implementation, this would use atoms.Button templ component
	return r.renderActionButton(buttonProps), nil
}

// RenderLayout implements schema.RuntimeRenderer.RenderLayout
// Renders layout container with fields
func (r *TemplateRenderer) RenderLayout(ctx context.Context, layout *schema.Layout, fields []*schema.Field, state map[string]any) (string, error) {
	if layout == nil {
		return "", fmt.Errorf("layout cannot be nil")
	}

	switch layout.Type {
	case schema.LayoutGrid:
		return r.renderGridLayout(ctx, layout, fields, state)
	case schema.LayoutTabs:
		return r.renderTabsLayout(ctx, layout, fields, state)
	case schema.LayoutSections:
		return r.renderSectionsLayout(ctx, layout, fields, state)
	default:
		// Fallback to default field rendering
		return r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
	}
}

// RenderSection implements schema.RuntimeRenderer.RenderSection
// Renders a section container with fields
func (r *TemplateRenderer) RenderSection(ctx context.Context, section *schema.Section, fields []*schema.Field, state map[string]any) (string, error) {
	if section == nil {
		return "", fmt.Errorf("section cannot be nil")
	}

	var sectionHTML strings.Builder

	// Section container
	sectionHTML.WriteString(fmt.Sprintf(`<div class="form-section" data-section-id="%s">`, section.ID))

	// Section title if provided
	if section.Title != "" {
		sectionHTML.WriteString(fmt.Sprintf(`<h3 class="section-title">%s</h3>`, section.Title))
	}

	// Section description if provided
	if section.Description != "" {
		sectionHTML.WriteString(fmt.Sprintf(`<p class="section-description">%s</p>`, section.Description))
	}

	// Render fields in this section
	fieldsHTML, err := r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
	if err != nil {
		return "", fmt.Errorf("failed to render section fields: %w", err)
	}
	sectionHTML.WriteString(fieldsHTML)

	sectionHTML.WriteString("</div>")
	return sectionHTML.String(), nil
}

// RenderTab implements schema.RuntimeRenderer.RenderTab
// Renders a tab container with fields
func (r *TemplateRenderer) RenderTab(ctx context.Context, tab *schema.Tab, fields []*schema.Field, state map[string]any) (string, error) {
	if tab == nil {
		return "", fmt.Errorf("tab cannot be nil")
	}

	var tabHTML strings.Builder

	// Tab content container
	tabHTML.WriteString(fmt.Sprintf(`<div class="tab-content" data-tab-id="%s">`, tab.ID))

	// Render fields in this tab
	fieldsHTML, err := r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
	if err != nil {
		return "", fmt.Errorf("failed to render tab fields: %w", err)
	}
	tabHTML.WriteString(fieldsHTML)

	tabHTML.WriteString("</div>")
	return tabHTML.String(), nil
}

// RenderStep implements schema.RuntimeRenderer.RenderStep
// Renders a step container with fields (for multi-step forms)
func (r *TemplateRenderer) RenderStep(ctx context.Context, step *schema.Step, fields []*schema.Field, state map[string]any) (string, error) {
	if step == nil {
		return "", fmt.Errorf("step cannot be nil")
	}

	var stepHTML strings.Builder

	// Step container
	stepHTML.WriteString(fmt.Sprintf(`<div class="form-step" data-step-id="%s">`, step.ID))

	// Step title if provided
	if step.Title != "" {
		stepHTML.WriteString(fmt.Sprintf(`<h3 class="step-title">%s</h3>`, step.Title))
	}

	// Render fields in this step
	fieldsHTML, err := r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
	if err != nil {
		return "", fmt.Errorf("failed to render step fields: %w", err)
	}
	stepHTML.WriteString(fieldsHTML)

	stepHTML.WriteString("</div>")
	return stepHTML.String(), nil
}

// Helper methods for internal rendering

// renderFormFields renders a list of fields with current state
// ENHANCED with state manager integration
func (r *TemplateRenderer) renderFormFields(ctx context.Context, fields []schema.Field, state map[string]any, formErrors map[string][]string) (string, error) {
	var fieldsHTML strings.Builder

	for _, field := range fields {
		// Get field state (use state manager if available, fallback to provided state)
		var value any
		var errors []string
		var touched, dirty bool
		
		if r.stateManager != nil {
			// Use state manager for accurate field state
			value = r.stateManager.GetValue(field.Name)
			errors = r.stateManager.GetFieldErrors(field.Name)
			touched = r.stateManager.IsFieldTouched(field.Name)
			dirty = r.stateManager.IsFieldDirty(field.Name)
		} else {
			// Fallback to provided state
			value = getFieldValue(state, field.Name)
			errors = getFieldErrors(formErrors, field.Name)
			touched = false
			dirty = false
		}

		// Render individual field
		fieldHTML, err := r.RenderField(ctx, &field, value, errors, touched, dirty)
		if err != nil {
			return "", fmt.Errorf("failed to render field %s: %w", field.Name, err)
		}

		// Wrap field in container with validation state attributes
		containerAttrs := r.buildFieldContainerAttributes(&field)
		fieldsHTML.WriteString(fmt.Sprintf(`<div class="field-container" data-field="%s"%s>`, field.Name, containerAttrs))
		fieldsHTML.WriteString(fieldHTML)
		fieldsHTML.WriteString("</div>\n")
	}

	return fieldsHTML.String(), nil
}

// renderFormWithLayout renders form using layout specifications
func (r *TemplateRenderer) renderFormWithLayout(ctx context.Context, s *schema.Schema, state map[string]any, errors map[string][]string) (string, error) {
	// Convert []schema.Field to []*schema.Field for layout methods
	fieldPtrs := make([]*schema.Field, len(s.Fields))
	for i := range s.Fields {
		fieldPtrs[i] = &s.Fields[i]
	}

	return r.RenderLayout(ctx, s.Layout, fieldPtrs, state)
}

// renderFormActions renders form action buttons
func (r *TemplateRenderer) renderFormActions(ctx context.Context, actions []schema.Action) (string, error) {
	var actionsHTML strings.Builder

	actionsHTML.WriteString(`<div class="form-actions">`)

	for _, action := range actions {
		actionHTML, err := r.RenderAction(ctx, &action, true) // Assume enabled for now
		if err != nil {
			return "", fmt.Errorf("failed to render action %s: %w", action.ID, err)
		}
		actionsHTML.WriteString(actionHTML)
	}

	actionsHTML.WriteString("</div>")
	return actionsHTML.String(), nil
}

// buildFormAttributes builds HTML attributes for form tag
func (r *TemplateRenderer) buildFormAttributes(s *schema.Schema) string {
	var attrs []string

	// Basic attributes
	if s.ID != "" {
		attrs = append(attrs, fmt.Sprintf(`id="%s"`, s.ID))
	}

	attrs = append(attrs, `class="schema-form"`)
	attrs = append(attrs, fmt.Sprintf(`data-schema-type="%s"`, s.Type))

	// Configuration-based attributes
	if s.Config != nil {
		if s.Config.Method != "" {
			attrs = append(attrs, fmt.Sprintf(`method="%s"`, s.Config.Method))
		}
		if s.Config.Action != "" {
			attrs = append(attrs, fmt.Sprintf(`action="%s"`, s.Config.Action))
		}
		// Add HTMX target if specified
		if s.Config.Target != "" {
			attrs = append(attrs, fmt.Sprintf(`hx-target="%s"`, s.Config.Target))
		}
	}

	// Add schema-level HTMX attributes if present
	if s.HTMX != nil {
		// Add various HTMX attributes that might be configured at schema level
		// This would need to be expanded based on the actual HTMX struct definition
	}

	return strings.Join(attrs, " ")
}

// Layout rendering methods

// renderGridLayout renders fields in a CSS grid layout
func (r *TemplateRenderer) renderGridLayout(ctx context.Context, layout *schema.Layout, fields []*schema.Field, state map[string]any) (string, error) {
	var gridHTML strings.Builder

	// Grid container with columns
	cols := 1
	if layout.Columns > 0 {
		cols = layout.Columns
	}

	gridHTML.WriteString(fmt.Sprintf(`<div class="grid grid-cols-%d gap-4">`, cols))

	// Convert []*schema.Field back to []schema.Field for renderFormFields
	fieldValues := make([]schema.Field, len(fields))
	for i, field := range fields {
		fieldValues[i] = *field
	}

	fieldsHTML, err := r.renderFormFields(ctx, fieldValues, state, nil)
	if err != nil {
		return "", err
	}
	gridHTML.WriteString(fieldsHTML)

	gridHTML.WriteString("</div>")
	return gridHTML.String(), nil
}

// renderTabsLayout renders fields in a tabbed interface
func (r *TemplateRenderer) renderTabsLayout(ctx context.Context, layout *schema.Layout, fields []*schema.Field, state map[string]any) (string, error) {
	// For now, render as simple sections
	// In a full implementation, this would create actual tab interface
	return r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
}

// renderSectionsLayout renders fields grouped by sections
func (r *TemplateRenderer) renderSectionsLayout(ctx context.Context, layout *schema.Layout, fields []*schema.Field, state map[string]any) (string, error) {
	// For now, render as simple sections
	return r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
}

// renderActionButton renders an action using the Button Templ component
func (r *TemplateRenderer) renderActionButton(props ButtonProps) string {
	// Convert to atoms.ButtonProps
	buttonProps := r.convertToButtonProps(props)

	// Use templ to render the Button component with text content
	var buf strings.Builder
	err := atoms.Button(buttonProps, templ.Raw(props.Text)).Render(context.Background(), &buf)
	if err != nil {
		// Fallback to basic button HTML in case of rendering failure
		return fmt.Sprintf(`<button type="%s" class="btn btn-%s"%s>%s</button>`,
			props.Type, props.Variant, r.getDisabledAttr(props.Enabled), props.Text)
	}
	return buf.String()
}

// Helper functions

// getFieldValue safely gets field value from state map
func getFieldValue(state map[string]any, fieldName string) any {
	if state == nil {
		return nil
	}
	return state[fieldName]
}

// getFieldErrors safely gets field errors from errors map
func getFieldErrors(errors map[string][]string, fieldName string) []string {
	if errors == nil {
		return nil
	}
	return errors[fieldName]
}

// RenderErrors implements schema.RuntimeRenderer.RenderErrors
// Renders validation errors using the ValidationMessages Templ component
func (r *TemplateRenderer) RenderErrors(ctx context.Context, errors map[string][]string) (string, error) {
	if len(errors) == 0 {
		return "", nil
	}

	// Convert errors map to ValidationMessages format
	var messages []molecules.ValidationMessage
	for fieldName, fieldErrors := range errors {
		for _, err := range fieldErrors {
			messages = append(messages, molecules.ValidationMessage{
				Type:    molecules.ValidationError,
				Message: err,
				Field:   fieldName,
				Code:    "validation_error",
			})
		}
	}

	// Create validation props
	validationProps := molecules.ValidationProps{
		Messages:     messages,
		ShowSuccess:  false,
		ShowWarnings: false,
		Inline:       false,
		Position:     "bottom",
		Class:        "form-errors",
	}

	// Use templ to render the ValidationMessages component
	var buf strings.Builder
	err := molecules.ValidationMessages(validationProps).Render(ctx, &buf)
	if err != nil {
		// Fallback to basic HTML in case of rendering failure
		return r.renderErrorsFallback(errors), nil
	}
	return buf.String(), nil
}

// RenderGroup implements schema.RuntimeRenderer.RenderGroup
// Renders a group container with fields
func (r *TemplateRenderer) RenderGroup(ctx context.Context, group *schema.Group, fields []*schema.Field, state map[string]any) (string, error) {
	if group == nil {
		return "", fmt.Errorf("group cannot be nil")
	}

	var groupHTML strings.Builder

	// Group container with border if specified
	borderClass := ""
	if group.Border {
		borderClass = " form-group-bordered"
	}
	groupHTML.WriteString(fmt.Sprintf(`<div class="form-group%s" data-group-id="%s">`, borderClass, group.ID))

	// Group label if provided (note: Group has Label, not Title)
	if group.Label != "" {
		groupHTML.WriteString(fmt.Sprintf(`<h4 class="group-label">%s</h4>`, group.Label))
	}

	// Group description if provided
	if group.Description != "" {
		groupHTML.WriteString(fmt.Sprintf(`<p class="group-description">%s</p>`, group.Description))
	}

	// Render fields in this group
	fieldsHTML, err := r.renderFormFields(ctx, fieldValuesFromPointers(fields), state, nil)
	if err != nil {
		return "", fmt.Errorf("failed to render group fields: %w", err)
	}
	groupHTML.WriteString(fieldsHTML)

	groupHTML.WriteString("</div>")
	return groupHTML.String(), nil
}

// fieldValuesFromPointers converts []*schema.Field to []schema.Field
func fieldValuesFromPointers(fields []*schema.Field) []schema.Field {
	result := make([]schema.Field, len(fields))
	for i, field := range fields {
		result[i] = *field
	}
	return result
}

// convertToButtonProps converts internal ButtonProps to atoms.ButtonProps
func (r *TemplateRenderer) convertToButtonProps(props ButtonProps) atoms.ButtonProps {
	// Map variant strings to atoms.ButtonVariant
	var variant atoms.ButtonVariant
	switch props.Variant {
	case "primary":
		variant = atoms.ButtonPrimary
	case "secondary":
		variant = atoms.ButtonSecondary
	case "destructive":
		variant = atoms.ButtonDestructive
	case "outline":
		variant = atoms.ButtonOutline
	case "ghost":
		variant = atoms.ButtonGhost
	case "link":
		variant = atoms.ButtonLink
	default:
		variant = atoms.ButtonPrimary
	}

	return atoms.ButtonProps{
		Variant:  variant,
		Size:     atoms.ButtonSizeMD, // Default size
		Type:     props.Type,
		Disabled: !props.Enabled,
		Class:    props.Class,
		ID:       props.ID,
	}
}

// getDisabledAttr returns disabled attribute string for fallback HTML
func (r *TemplateRenderer) getDisabledAttr(enabled bool) string {
	if !enabled {
		return " disabled"
	}
	return ""
}

// renderErrorsFallback provides fallback error rendering when Templ fails
func (r *TemplateRenderer) renderErrorsFallback(errors map[string][]string) string {
	var errorsHTML strings.Builder
	errorsHTML.WriteString(`<div class="form-errors">`)

	for fieldName, fieldErrors := range errors {
		if len(fieldErrors) > 0 {
			errorsHTML.WriteString(fmt.Sprintf(`<div class="field-errors" data-field="%s">`, fieldName))
			errorsHTML.WriteString(fmt.Sprintf(`<span class="field-name">%s:</span>`, fieldName))
			errorsHTML.WriteString(`<ul class="error-list">`)

			for _, err := range fieldErrors {
				errorsHTML.WriteString(fmt.Sprintf(`<li class="error-item">%s</li>`, err))
			}

			errorsHTML.WriteString(`</ul></div>`)
		}
	}

	errorsHTML.WriteString(`</div>`)
	return errorsHTML.String()
}

// ButtonProps represents button properties for action rendering
type ButtonProps struct {
	Type    string
	Text    string
	Variant string
	Enabled bool
	ID      string
	Class   string
}

// ═══════════════════════════════════════════════════════════════════════════
// UI Validation Integration Methods (EXTENSION)
// ═══════════════════════════════════════════════════════════════════════════

// enhancePropsWithValidation enhances FormField props with validation orchestration data
// ENHANCED with comprehensive theme-aware validation token integration
func (r *TemplateRenderer) enhancePropsWithValidation(field *schema.Field, props *molecules.FormFieldProps) {
	if r.orchestrator == nil {
		return
	}
	
	fieldName := field.Name
	
	// Get current validation state from orchestrator
	validationState := r.orchestrator.GetValidationStateForField(fieldName)
	
	// Set validation state and loading indicator
	props.ValidationState = validationState
	props.ValidationLoading = r.orchestrator.IsFieldValidating(fieldName)
	
	// Determine dark mode state from context
	darkMode := r.determineContextDarkMode(field, props)
	props.DarkMode = darkMode
	
	// Resolve and apply comprehensive validation tokens using mapper's token resolution
	r.resolveFieldValidationTokens(field, props, validationState, darkMode)
	
	// Set theme ID from context for token resolution
	if themeID := r.getThemeIDFromContext(field); themeID != "" {
		props.ThemeID = themeID
	}
	
	// Add validation state CSS classes (legacy support)
	validationClass := r.getValidationStateClass(validationState)
	if validationClass != "" {
		if props.Class != "" {
			props.Class += " " + validationClass
		} else {
			props.Class = validationClass
		}
	}
	
	// Set current validation errors from orchestrator
	if errors := r.orchestrator.GetFieldErrors(fieldName); len(errors) > 0 {
		props.Errors = errors
	}
	
	// Add client validation rules if enabled
	if r.mapper.HasClientValidationRules(field) {
		clientRules := r.mapper.ExtractClientValidationRules(field)
		props.ClientRules = &clientRules
		
		// Configure validation endpoint for async validation
		if r.orchestrator.HasAsyncValidation(fieldName) {
			props.OnValidate = r.buildValidationEndpoint(fieldName)
			props.ValidationDebounce = r.orchestrator.GetDebounceTimeForField(fieldName)
		}
		
		// Add client validation attributes as CSS classes
		validationClasses := r.buildValidationClasses(clientRules)
		
		// Add validation classes to existing class string
		if len(validationClasses) > 0 {
			validationClassStr := strings.Join(validationClasses, " ")
			if props.Class != "" {
				props.Class += " " + validationClassStr
			} else {
				props.Class = validationClassStr
			}
		}
	}
}

// resolveFieldValidationTokens resolves and applies comprehensive validation tokens using mapper's token resolution
func (r *TemplateRenderer) resolveFieldValidationTokens(
	field *schema.Field, 
	props *molecules.FormFieldProps, 
	validationState validation.ValidationState,
	darkMode bool,
) {
	if r.mapper == nil {
		return
	}
	
	// Use mapper's comprehensive token resolution system
	ctx := context.Background()
	resolvedTokens := r.mapper.ResolveFieldTokens(ctx, *props, darkMode)
	
	// Set resolved tokens on props
	if props.Tokens == nil {
		props.Tokens = make(map[string]string)
	}
	
	// Apply all resolved tokens
	for tokenKey, tokenValue := range resolvedTokens {
		props.Tokens[tokenKey] = tokenValue
	}
	
	// Ensure validation state-specific tokens are applied
	r.ensureValidationStateTokens(props, validationState, darkMode)
}

// buildFieldContainerAttributes builds HTML attributes for field containers with validation state
func (r *TemplateRenderer) buildFieldContainerAttributes(field *schema.Field) string {
	if r.orchestrator == nil {
		return ""
	}
	
	fieldName := field.Name
	validationState := r.orchestrator.GetValidationStateForField(fieldName)
	
	var attrs []string
	
	// Add validation state attribute
	attrs = append(attrs, fmt.Sprintf(`data-validation-state="%s"`, validationState.String()))
	
	// Add pending validation indicator
	if r.orchestrator.IsValidationInProgress() {
		pendingFields := r.orchestrator.GetPendingValidations()
		for _, pending := range pendingFields {
			if pending == fieldName {
				attrs = append(attrs, `data-validation-pending="true"`)
				break
			}
		}
	}
	
	if len(attrs) > 0 {
		return " " + strings.Join(attrs, " ")
	}
	return ""
}

// getValidationStateClass returns CSS class for validation state
func (r *TemplateRenderer) getValidationStateClass(state validation.ValidationState) string {
	switch state {
	case validation.ValidationStateValidating:
		return "field-validating"
	case validation.ValidationStateValid:
		return "field-valid"
	case validation.ValidationStateInvalid:
		return "field-invalid"
	case validation.ValidationStateWarning:
		return "field-warning"
	default:
		return ""
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// THEME-AWARE VALIDATION TOKEN RESOLUTION METHODS (ENHANCED)
// ═══════════════════════════════════════════════════════════════════════════

// determineContextDarkMode determines dark mode state from multiple sources
func (r *TemplateRenderer) determineContextDarkMode(field *schema.Field, props *molecules.FormFieldProps) bool {
	// Check if already set on props (highest priority)
	if props.DarkMode {
		return true
	}
	
	// Check field-specific dark mode setting
	if field.Config != nil {
		if darkMode, exists := field.Config["darkMode"]; exists {
			if dark, ok := darkMode.(bool); ok {
				return dark
			}
		}
	}
	
	// Schema Config doesn't contain darkMode - would need to be added if required
	// if r.schema != nil && r.schema.Config != nil {
	//   // Add darkMode field to Config struct if needed
	// }
	
	// Check state manager for user preference
	if r.stateManager != nil {
		if darkModeValue := r.stateManager.GetValue("darkMode"); darkModeValue != nil {
			if dark, ok := darkModeValue.(bool); ok {
				return dark
			}
		}
	}
	
	return false
}

// getThemeIDFromContext extracts theme ID from various context sources
func (r *TemplateRenderer) getThemeIDFromContext(field *schema.Field) string {
	// Check field-specific theme ID
	if field.Config != nil {
		if themeID, exists := field.Config["themeId"]; exists {
			if theme, ok := themeID.(string); ok {
				return theme
			}
		}
	}
	
	// Schema Config doesn't contain themeId - would need to be added if required
	// if r.schema != nil && r.schema.Config != nil {
	//   // Add themeId field to Config struct if needed
	// }
	
	// Check state manager for user theme selection
	if r.stateManager != nil {
		if themeValue := r.stateManager.GetValue("theme"); themeValue != nil {
			if theme, ok := themeValue.(string); ok {
				return theme
			}
		}
	}
	
	return "default"
}

// buildValidationClasses builds CSS classes for client-side validation
func (r *TemplateRenderer) buildValidationClasses(clientRules molecules.ClientValidationRules) []string {
	var validationClasses []string
	
	if clientRules.Required {
		validationClasses = append(validationClasses, "validate-required")
	}
	
	if clientRules.MinLength != nil {
		validationClasses = append(validationClasses, fmt.Sprintf("validate-min-length-%d", *clientRules.MinLength))
	}
	
	if clientRules.MaxLength != nil {
		validationClasses = append(validationClasses, fmt.Sprintf("validate-max-length-%d", *clientRules.MaxLength))
	}
	
	if clientRules.Pattern != "" {
		validationClasses = append(validationClasses, "validate-pattern")
	}
	
	if clientRules.Format != "" {
		validationClasses = append(validationClasses, fmt.Sprintf("validate-format-%s", clientRules.Format))
	}
	
	if clientRules.Min != nil {
		validationClasses = append(validationClasses, fmt.Sprintf("validate-min-%.2f", *clientRules.Min))
	}
	
	if clientRules.Max != nil {
		validationClasses = append(validationClasses, fmt.Sprintf("validate-max-%.2f", *clientRules.Max))
	}
	
	return validationClasses
}

// ensureValidationStateTokens ensures validation state-specific tokens are properly set
func (r *TemplateRenderer) ensureValidationStateTokens(
	props *molecules.FormFieldProps, 
	validationState validation.ValidationState,
	darkMode bool,
) {
	// Ensure validation state tokens exist with fallbacks
	validationStateTokens := r.getValidationStateTokens(validationState, darkMode)
	
	for tokenKey, tokenValue := range validationStateTokens {
		// Only set if not already resolved by mapper
		if _, exists := props.Tokens[tokenKey]; !exists {
			props.Tokens[tokenKey] = tokenValue
		}
	}
}

// getValidationStateTokens returns fallback validation tokens for a specific state
func (r *TemplateRenderer) getValidationStateTokens(state validation.ValidationState, darkMode bool) map[string]string {
	// Base colors (light mode defaults)
	baseTokens := map[string]string{
		molecules.TokenFieldBackground: "#ffffff",
		molecules.TokenFieldBorder:     "#e5e7eb", 
		molecules.TokenFieldText:       "#111827",
		molecules.TokenFieldRadius:     "0.375rem",
		molecules.TokenFieldSpacing:    "0.75rem",
		molecules.TokenFieldShadow:     "0 1px 2px 0 rgb(0 0 0 / 0.05)",
	}
	
	// Apply dark mode base tokens if enabled
	if darkMode {
		baseTokens[molecules.TokenFieldBackground] = "#111827"
		baseTokens[molecules.TokenFieldBorder] = "#374151"
		baseTokens[molecules.TokenFieldText] = "#f9fafb"
		baseTokens[molecules.TokenFieldShadow] = "0 1px 2px 0 rgb(0 0 0 / 0.25)"
	}
	
	// Apply validation state-specific tokens
	switch state {
	case validation.ValidationStateValidating:
		borderColor := "#f59e0b"
		if darkMode {
			borderColor = "#fbbf24"
		}
		baseTokens[molecules.TokenFieldBorder] = borderColor
		baseTokens[molecules.TokenValidationLoading+".color"] = borderColor
		baseTokens[molecules.TokenValidationLoading+".message"] = "Validating..."
		baseTokens[molecules.TokenValidationLoading+".icon"] = "⟳"
		
	case validation.ValidationStateValid:
		borderColor := "#16a34a"
		if darkMode {
			borderColor = "#22c55e"
		}
		baseTokens[molecules.TokenFieldBorder] = borderColor
		baseTokens[molecules.TokenValidationSuccess+".color"] = borderColor
		baseTokens[molecules.TokenValidationSuccess+".message"] = "Valid"
		baseTokens[molecules.TokenValidationSuccess+".icon"] = "✓"
		
	case validation.ValidationStateInvalid:
		borderColor := "#dc2626"
		if darkMode {
			borderColor = "#ef4444"
		}
		baseTokens[molecules.TokenFieldBorder] = borderColor
		baseTokens[molecules.TokenValidationError+".color"] = borderColor
		baseTokens[molecules.TokenValidationError+".message"] = "Invalid"
		baseTokens[molecules.TokenValidationError+".icon"] = "✗"
		
	case validation.ValidationStateWarning:
		borderColor := "#d97706"
		if darkMode {
			borderColor = "#f59e0b"
		}
		baseTokens[molecules.TokenFieldBorder] = borderColor
		baseTokens[molecules.TokenValidationWarning+".color"] = borderColor
		baseTokens[molecules.TokenValidationWarning+".message"] = "Warning"
		baseTokens[molecules.TokenValidationWarning+".icon"] = "⚠"
	}
	
	return baseTokens
}

// validationStateToString converts validation state enum to string for token lookup
func (r *TemplateRenderer) validationStateToString(state validation.ValidationState) string {
	switch state {
	case validation.ValidationStateValidating:
		return "validating"
	case validation.ValidationStateValid:
		return "valid"
	case validation.ValidationStateInvalid:
		return "invalid"
	case validation.ValidationStateWarning:
		return "warning"
	default:
		return "idle"
	}
}

// renderFormFieldWithTheme renders a form field with comprehensive theme-aware token resolution
// This is the enhanced rendering method that integrates with the theme coordination layer
func (r *TemplateRenderer) renderFormFieldWithTheme(
	ctx context.Context,
	field *schema.Field, 
	value any, 
	errors []string, 
	touched, dirty bool,
	themeContext map[string]any,
) (string, error) {
	// Convert field to FormField props with enhanced token resolution
	props, err := r.mapper.ConvertField(field, value, errors, touched, dirty)
	if err != nil {
		return "", fmt.Errorf("failed to convert field %s: %w", field.Name, err)
	}
	
	// Apply theme context to props
	if themeID, exists := themeContext["themeId"]; exists {
		if theme, ok := themeID.(string); ok {
			props.ThemeID = theme
		}
	}
	
	if darkMode, exists := themeContext["darkMode"]; exists {
		if dark, ok := darkMode.(bool); ok {
			props.DarkMode = dark
		}
	}
	
	// Enhance props with validation orchestration data and theme-aware tokens
	r.enhancePropsWithValidation(field, &props)
	
	// Resolve component renderer for this field type
	renderer, err := r.registry.Resolve(field)
	if err != nil {
		return "", fmt.Errorf("failed to resolve renderer for field %s (type: %s): %w", field.Name, field.Type, err)
	}
	
	// Render using the resolved component with theme-aware tokens
	html, err := renderer.Render(ctx, props)
	if err != nil {
		return "", fmt.Errorf("failed to render field %s: %w", field.Name, err)
	}
	
	return html, nil
}

// SetThemeManager sets a theme manager for enhanced theme coordination
// This integrates with the theme coordination layer for multi-tenant support
func (r *TemplateRenderer) SetThemeManager(themeManager interface{}) {
	// This would integrate with the theme coordination manager
	// For now, this is a placeholder for future theme manager integration
}

// GetThemeContext extracts theme context from various sources for field rendering
func (r *TemplateRenderer) GetThemeContext(field *schema.Field) map[string]any {
	themeContext := make(map[string]any)
	
	// Extract theme ID
	themeID := r.getThemeIDFromContext(field)
	themeContext["themeId"] = themeID
	
	// Extract dark mode setting
	darkMode := r.determineContextDarkMode(field, &molecules.FormFieldProps{})
	themeContext["darkMode"] = darkMode
	
	// Extract tenant context if available  
	// if r.schema != nil && r.schema.Config != nil {
	//   // Schema Config doesn't contain tenant - would need to be added if required
	// }
	
	return themeContext
}

// ═══════════════════════════════════════════════════════════════════════════
// ENHANCED RENDERING METHODS WITH THEME COORDINATION
// ═══════════════════════════════════════════════════════════════════════════

// RenderFieldWithThemeContext renders a field with explicit theme context
// This is an enhanced version of RenderField that accepts theme context
func (r *TemplateRenderer) RenderFieldWithThemeContext(
	ctx context.Context,
	field *schema.Field,
	value any,
	errors []string,
	touched, dirty bool,
	themeContext map[string]any,
) (string, error) {
	return r.renderFormFieldWithTheme(ctx, field, value, errors, touched, dirty, themeContext)
}

// GetContextForValidation extracts validation context including theme information
func (r *TemplateRenderer) GetContextForValidation(field *schema.Field) map[string]any {
	validationContext := make(map[string]any)
	
	// Include theme context
	themeContext := r.GetThemeContext(field)
	for key, value := range themeContext {
		validationContext[key] = value
	}
	
	// Include validation state
	if r.orchestrator != nil {
		validationState := r.orchestrator.GetValidationStateForField(field.Name)
		validationContext["validationState"] = validationState
		validationContext["isValidating"] = r.orchestrator.IsFieldValidating(field.Name)
		validationContext["hasAsyncValidation"] = r.orchestrator.HasAsyncValidation(field.Name)
	}
	
	return validationContext
}

// buildValidationEndpoint builds HTMX validation endpoint URL for a field
func (r *TemplateRenderer) buildValidationEndpoint(fieldName string) string {
	// Get base validation endpoint from schema config
	baseEndpoint := "/api/validate"
	
	// if r.schema != nil && r.schema.Config != nil {
	//   // Schema Config doesn't contain validationEndpoint - would need to be added if required
	// }
	
	// Add field-specific path
	return fmt.Sprintf("%s/field/%s", baseEndpoint, fieldName)
}

// Helper methods for token resolution


// getValidationTokenSystem returns the validation token system interface
func (r *TemplateRenderer) getValidationTokenSystem() (interface{ 
	GetValidationTokensForState(string) map[string]string
	GetDarkModeValidationTokens(string) map[string]string
}, bool) {
	// This would integrate with the tokens/validation_tokens.go package
	// For now, return false to use fallback values
	return nil, false
}

// resolveTokenToActualValue attempts to resolve token reference to actual CSS value
func (r *TemplateRenderer) resolveTokenToActualValue(tokenRef string) string {
	// This would integrate with the full token resolution system
	// For now, return the token reference itself
	return tokenRef
}
