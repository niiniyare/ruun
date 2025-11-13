package views

import (
	"context"
	"fmt"
	"strings"

	"github.com/a-h/templ"
	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// TemplateRenderer implements schema.RuntimeRenderer using Templ components
// This is the main bridge between the pure schema system and the Views layer
// EXTENDED with UI validation orchestration
type TemplateRenderer struct {
	registry     *ComponentRegistry
	mapper       *FieldMapper
	schema       *schema.Schema
	stateManager *StateManager                // Form state management
	orchestrator *UIValidationOrchestrator    // UI validation orchestration
}

// NewTemplateRenderer creates a new template renderer with dependencies
func NewTemplateRenderer(registry *ComponentRegistry, mapper *FieldMapper, s *schema.Schema, stateManager *StateManager, orchestrator *UIValidationOrchestrator) *TemplateRenderer {
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
func (r *TemplateRenderer) enhancePropsWithValidation(field *schema.Field, props *molecules.FormFieldProps) {
	if r.orchestrator == nil {
		return
	}
	
	fieldName := field.Name
	
	// Add validation state CSS classes
	validationState := r.orchestrator.GetValidationStateForField(fieldName)
	validationClass := r.getValidationStateClass(validationState)
	if validationClass != "" {
		if props.Class != "" {
			props.Class += " " + validationClass
		} else {
			props.Class = validationClass
		}
	}
	
	// Add client validation rules if enabled
	if r.mapper.HasClientValidationRules(field) {
		clientRules := r.mapper.ExtractClientValidationRules(field)
		
		// Add client validation attributes to props
		if props.Attributes == nil {
			props.Attributes = make(map[string]string)
		}
		
		// Add validation trigger attributes
		if clientRules.Required {
			props.Attributes["data-validate-required"] = "true"
		}
		if clientRules.MinLength != nil {
			props.Attributes["data-validate-min-length"] = fmt.Sprintf("%d", *clientRules.MinLength)
		}
		if clientRules.MaxLength != nil {
			props.Attributes["data-validate-max-length"] = fmt.Sprintf("%d", *clientRules.MaxLength)
		}
		if clientRules.Pattern != "" {
			props.Attributes["data-validate-pattern"] = clientRules.Pattern
		}
		if clientRules.Format != "" {
			props.Attributes["data-validate-format"] = clientRules.Format
		}
		
		// Add debounce timing
		props.Attributes["data-validate-debounce"] = "300"
		props.Attributes["data-field-name"] = fieldName
	}
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
func (r *TemplateRenderer) getValidationStateClass(state ValidationState) string {
	switch state {
	case ValidationStateValidating:
		return "field-validating"
	case ValidationStateValid:
		return "field-valid"
	case ValidationStateInvalid:
		return "field-invalid"
	case ValidationStateWarning:
		return "field-warning"
	default:
		return ""
	}
}
