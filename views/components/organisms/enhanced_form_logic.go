package organisms

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	
	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/components/utils"
)

// Class generation functions for enhanced form

func getEnhancedFormClasses(props EnhancedFormProps) string {
	return utils.TwMerge(
		"enhanced-form",
		"space-y-6",
		getFormLayoutClasses(props.Layout, props.GridCols),
		utils.If(props.ReadOnly, "form-readonly"),
		utils.If(props.Debug, "form-debug"),
		props.Class,
	)
}

func getFormLayoutClasses(layout FormLayout, gridCols int) string {
	switch layout {
	case FormLayoutHorizontal:
		return "space-y-4"
	case FormLayoutInline:
		return "flex flex-wrap items-end gap-4 space-y-0"
	case FormLayoutGrid:
		cols := utils.IfElse(gridCols > 0, gridCols, 2)
		return fmt.Sprintf("grid gap-6 space-y-0 %s", getGridClasses(cols))
	default:
		return "space-y-6"
	}
}

func getGridClasses(cols int) string {
	switch cols {
	case 1:
		return "grid-cols-1"
	case 2:
		return "grid-cols-1 md:grid-cols-2"
	case 3:
		return "grid-cols-1 md:grid-cols-2 lg:grid-cols-3"
	case 4:
		return "grid-cols-1 md:grid-cols-2 lg:grid-cols-4"
	default:
		return "grid-cols-1 md:grid-cols-2"
	}
}

func getFormContentClasses(props EnhancedFormProps) string {
	return utils.TwMerge(
		"form-content",
		utils.If(props.Layout == FormLayoutGrid, "space-y-0"),
		utils.If(props.Layout != FormLayoutGrid, "space-y-6"),
	)
}

func getDirectFieldsClasses(props EnhancedFormProps) string {
	switch props.Layout {
	case FormLayoutGrid:
		return fmt.Sprintf("grid gap-4 %s", getGridClasses(utils.IfElse(props.GridCols > 0, props.GridCols, 2)))
	case FormLayoutInline:
		return "flex flex-wrap gap-4"
	case FormLayoutHorizontal:
		return "grid grid-cols-1 md:grid-cols-2 gap-4"
	default:
		return "space-y-4"
	}
}

func getSectionClasses(section EnhancedFormSection, formProps EnhancedFormProps) string {
	return utils.TwMerge(
		"form-section",
		"p-6 bg-card rounded-lg border",
		utils.If(section.Collapsible, "collapsible-section"),
		utils.If(section.Required, "required-section"),
		section.Class,
	)
}

func getSectionFieldsClasses(section EnhancedFormSection, formProps EnhancedFormProps) string {
	if section.Layout == "grid" || (section.Layout == "" && formProps.Layout == FormLayoutGrid) {
		cols := utils.IfElse(section.Columns > 0, section.Columns, 2)
		return fmt.Sprintf("grid gap-4 %s", getGridClasses(cols))
	}
	
	switch section.Layout {
	case "horizontal":
		return "grid grid-cols-1 md:grid-cols-2 gap-4"
	case "inline":
		return "flex flex-wrap gap-4"
	default:
		return "space-y-4"
	}
}

func getFormActionsClasses(props EnhancedFormProps) string {
	return utils.TwMerge(
		"form-actions",
		"flex items-center justify-end",
		"pt-6 mt-8 border-t border-border",
		utils.If(props.Layout == FormLayoutInline, "border-t-0 pt-0 mt-0"),
	)
}

// Helper functions for form behavior

func getFormMethod(method string) string {
	if method != "" {
		return strings.ToUpper(method)
	}
	return "POST"
}

func getCSRFFieldName(fieldName string) string {
	if fieldName != "" {
		return fieldName
	}
	return "_csrf_token"
}

// Alpine.js data generation for enhanced form state management
func getEnhancedFormAlpineData(props EnhancedFormProps) string {
	initialData := props.InitialData
	if initialData == nil {
		initialData = make(map[string]any)
	}
	
	initialDataJSON, _ := json.Marshal(initialData)
	
	sections := len(props.Sections)
	if props.Schema != nil && len(props.Schema.Children) > 0 {
		sections = len(props.Schema.Children)
	}
	
	autoSaveInterval := utils.IfElse(props.AutoSaveInterval > 0, props.AutoSaveInterval, 30)
	
	return fmt.Sprintf(`{
		// Core form state
		formData: %s,
		initialData: %s,
		errors: {},
		touched: {},
		dirty: {},
		
		// Validation state
		validating: {},
		validateCount: {},
		valid: true,
		
		// Form lifecycle
		loading: false,
		submitting: false,
		submitCount: 0,
		
		// Auto-save
		autoSaving: false,
		autoSaveCount: 0,
		autoSaveEnabled: %s,
		autoSaveInterval: %d,
		autoSaveTimer: null,
		lastSaved: null,
		saving: false,
		
		// Progress tracking
		currentStep: 1,
		totalSteps: %d,
		showProgress: %s,
		showSaveState: %s,
		
		// Schema integration
		schemaVersion: '%s',
		dynamicSchema: %s,
		conditionalFields: [],
		
		// Computed properties
		get progressText() {
			return this.totalSteps > 1 ? 'Step ' + this.currentStep + ' of ' + this.totalSteps : '';
		},
		
		get progressPercentage() {
			return this.totalSteps > 1 ? Math.round((this.currentStep / this.totalSteps) * 100) + '%% complete' : '';
		},
		
		get progressStyle() {
			return this.totalSteps > 1 ? 'width: ' + ((this.currentStep / this.totalSteps) * 100) + '%%' : '';
		},
		
		get isDirty() {
			return Object.keys(this.dirty).some(key => this.dirty[key]);
		},
		
		get hasErrors() {
			return Object.keys(this.errors).length > 0;
		},
		
		// Form initialization
		initEnhancedForm() {
			this.setupAutoSave();
			this.setupValidation();
			this.setupConditionalLogic();
			this.loadInitialData();
			%s // Custom initialization
		},
		
		// Data loading
		loadInitialData() {
			%s
		},
		
		// Field management
		updateField(fieldName, value) {
			const oldValue = this.formData[fieldName];
			this.formData[fieldName] = value;
			this.touched[fieldName] = true;
			this.dirty[fieldName] = value !== this.initialData[fieldName];
			
			// Trigger validation if strategy is realtime
			if ('%s' === 'realtime') {
				this.validateField(fieldName);
			}
			
			// Trigger conditional logic evaluation
			this.evaluateConditionals();
			
			// Reset auto-save timer
			if (this.autoSaveEnabled) {
				this.resetAutoSaveTimer();
			}
			
			// Custom field change handler
			%s
		},
		
		// Validation
		validateField(fieldName) {
			if (this.validating[fieldName]) return;
			
			this.validating[fieldName] = true;
			this.validateCount[fieldName] = (this.validateCount[fieldName] || 0) + 1;
			
			// Client-side validation first
			const error = this.validateFieldClientSide(fieldName);
			if (error) {
				this.errors[fieldName] = error;
				this.validating[fieldName] = false;
				this.updateValidState();
				return;
			}
			
			// Server-side validation if URL provided
			if ('%s' && this.formData[fieldName] !== undefined) {
				fetch('%s', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'X-Requested-With': 'XMLHttpRequest'
					},
					body: JSON.stringify({
						field: fieldName,
						value: this.formData[fieldName],
						formData: this.formData
					})
				})
				.then(response => response.json())
				.then(data => {
					if (data.error) {
						this.errors[fieldName] = data.error;
					} else {
						delete this.errors[fieldName];
					}
					this.validating[fieldName] = false;
					this.updateValidState();
				})
				.catch(() => {
					this.validating[fieldName] = false;
				});
			} else {
				delete this.errors[fieldName];
				this.validating[fieldName] = false;
				this.updateValidState();
			}
		},
		
		validateFieldClientSide(fieldName) {
			const value = this.formData[fieldName];
			// Add client-side validation logic here
			// This would integrate with the FormField validation rules
			return null;
		},
		
		validateAllFields() {
			Object.keys(this.formData).forEach(fieldName => {
				this.validateField(fieldName);
			});
		},
		
		updateValidState() {
			this.valid = Object.keys(this.errors).length === 0;
		},
		
		// Form submission
		submitForm() {
			this.submitting = true;
			this.submitCount++;
			
			// Final validation
			this.validateAllFields();
			
			if (!this.valid) {
				this.submitting = false;
				return false;
			}
			
			// Custom submit handler
			%s
			
			return true;
		},
		
		// Auto-save functionality
		setupAutoSave() {
			if (!this.autoSaveEnabled) return;
			
			this.autoSaveTimer = setInterval(() => {
				if (this.isDirty && this.valid && !this.submitting) {
					this.autoSave();
				}
			}, this.autoSaveInterval * 1000);
		},
		
		autoSave() {
			if (this.autoSaving || !this.isDirty) return;
			
			this.autoSaving = true;
			this.saving = true;
			
			const url = '%s' || ('%s' + '/auto-save');
			
			fetch(url, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-Requested-With': 'XMLHttpRequest'
				},
				body: JSON.stringify(this.formData)
			})
			.then(response => response.json())
			.then(data => {
				if (data.success) {
					this.lastSaved = new Date().toLocaleTimeString();
					this.autoSaveCount++;
					// Reset dirty state
					Object.keys(this.dirty).forEach(key => {
						this.dirty[key] = false;
					});
				}
			})
			.finally(() => {
				this.autoSaving = false;
				this.saving = false;
			});
			
			// Custom auto-save handler
			%s
		},
		
		resetAutoSaveTimer() {
			if (this.autoSaveTimer) {
				clearInterval(this.autoSaveTimer);
				this.setupAutoSave();
			}
		},
		
		// Step navigation for multi-step forms
		nextStep() {
			if (this.currentStep < this.totalSteps) {
				this.currentStep++;
			}
		},
		
		prevStep() {
			if (this.currentStep > 1) {
				this.currentStep--;
			}
		},
		
		goToStep(step) {
			if (step >= 1 && step <= this.totalSteps) {
				this.currentStep = step;
			}
		},
		
		// Conditional logic
		setupConditionalLogic() {
			this.$watch('formData', (newData, oldData) => {
				this.evaluateConditionals();
			}, { deep: true });
		},
		
		evaluateConditionals() {
			// Evaluate conditional field visibility
			// This would integrate with the schema conditional logic
		},
		
		// Utility methods
		resetForm() {
			this.formData = { ...this.initialData };
			this.errors = {};
			this.touched = {};
			this.dirty = {};
			this.valid = true;
			this.currentStep = 1;
		},
		
		cancelForm() {
			if (this.isDirty) {
				if (confirm('You have unsaved changes. Are you sure you want to cancel?')) {
					this.resetForm();
				}
			} else {
				this.resetForm();
			}
		},
		
		// Setup validation strategy
		setupValidation() {
			const strategy = '%s';
			if (strategy === 'onblur') {
				// Setup blur validation
			} else if (strategy === 'realtime') {
				// Already handled in updateField
			}
			// onsubmit is handled in submitForm
		}
	}`,
		string(initialDataJSON),
		string(initialDataJSON),
		strconv.FormatBool(props.AutoSave),
		autoSaveInterval,
		sections,
		strconv.FormatBool(props.ShowProgress),
		strconv.FormatBool(props.ShowSaveState),
		getSchemaVersion(props),
		strconv.FormatBool(props.Schema != nil),
		utils.IfElse(props.OnFieldChange != "", props.OnFieldChange, "// No custom initialization"),
		utils.IfElse(props.LoadDataURL != "", fmt.Sprintf("fetch('%s').then(r => r.json()).then(data => { this.formData = {...this.formData, ...data}; })", props.LoadDataURL), "// No data loading URL"),
		props.ValidationStrategy,
		utils.IfElse(props.OnFieldChange != "", props.OnFieldChange, "// No custom field change handler"),
		props.ValidationURL,
		props.ValidationURL,
		utils.IfElse(props.OnSubmit != "", props.OnSubmit, "// No custom submit handler"),
		props.AutoSaveURL,
		props.SubmitURL,
		utils.IfElse(props.OnAutoSave != "", props.OnAutoSave, "// No custom auto-save handler"),
		props.ValidationStrategy,
	)
}

func getSchemaVersion(props EnhancedFormProps) string {
	if props.Schema != nil && props.Schema.Version != "" {
		return props.Schema.Version
	}
	return "1.0.0"
}

// Section management functions

func getSectionAlpineData(section EnhancedFormSection) string {
	return fmt.Sprintf(`{
		collapsed: %s,
		toggleSection() {
			this.collapsed = !this.collapsed;
		}
	}`, strconv.FormatBool(section.Collapsed))
}

// Schema integration functions

func convertSchemaFieldToFormField(field schema.Field, props EnhancedFormProps) molecules.FormFieldProps {
	formField := molecules.FormFieldProps{
		ID:          field.Name,
		Name:        field.Name,
		Label:       field.GetLocalizedLabel(utils.IfElse(props.Locale != "", props.Locale, "en")),
		Type:        mapSchemaFieldTypeToFormFieldType(field.Type),
		Value:       getFieldValueFromData(field.Name, props.InitialData),
		Placeholder: field.GetLocalizedPlaceholder(utils.IfElse(props.Locale != "", props.Locale, "en")),
		HelpText:    field.Description,
		Required:    field.Required,
		Disabled:    props.ReadOnly || field.Disabled,
		Readonly:    field.Readonly,
	}
	
	// Add validation integration
	if field.Validation != nil {
		formField.ValidationState = molecules.ValidationStateIdle
		formField.OnValidate = utils.IfElse(props.ValidationURL != "", props.ValidationURL, "")
		formField.ValidationDebounce = 300
	}
	
	// Add HTMX integration for real-time validation
	if props.ValidationStrategy == "realtime" && props.ValidationURL != "" {
		formField.HXPost = props.ValidationURL
		formField.HXTarget = fmt.Sprintf("#field-%s-feedback", field.Name)
		formField.HXTrigger = "change, keyup delay:300ms"
	}
	
	// Add Alpine.js integration
	formField.AlpineModel = fmt.Sprintf("formData.%s", field.Name)
	formField.AlpineChange = fmt.Sprintf("updateField('%s', $event.target.value)", field.Name)
	
	if props.ValidationStrategy == "onblur" {
		formField.AlpineBlur = fmt.Sprintf("validateField('%s')", field.Name)
	}
	
	// Handle field options for select/radio/checkbox types
	if len(field.Options) > 0 {
		formField.Options = convertSchemaOptionsToFormOptions(field.Options)
	}
	
	return enhanceFormFieldProps(formField, props)
}

func mapSchemaFieldTypeToFormFieldType(fieldType schema.FieldType) string {
	switch fieldType {
	case schema.FieldString:
		return "text"
	case schema.FieldEmail:
		return "email"
	case schema.FieldPassword:
		return "password"
	case schema.FieldNumber:
		return "number"
	case schema.FieldText:
		return "textarea"
	case schema.FieldSelect:
		return "select"
	case schema.FieldRadio:
		return "radio"
	case schema.FieldCheckbox:
		return "checkbox"
	case schema.FieldDate:
		return "date"
	case schema.FieldTime:
		return "time"
	case schema.FieldDatetime:
		return "datetime"
	case schema.FieldFile:
		return "file"
	case schema.FieldUrl:
		return "url"
	case schema.FieldTel:
		return "tel"
	default:
		return "text"
	}
}

func getFieldValueFromData(fieldName string, data map[string]any) string {
	if data == nil {
		return ""
	}
	
	if value, exists := data[fieldName]; exists && value != nil {
		return fmt.Sprintf("%v", value)
	}
	
	return ""
}

func convertSchemaOptionsToFormOptions(options []schema.FieldOption) []molecules.SelectOption {
	formOptions := make([]molecules.SelectOption, len(options))
	for i, option := range options {
		formOptions[i] = molecules.SelectOption{
			Value:       option.Value,
			Label:       option.Label,
			Description: option.Description,
			Disabled:    option.Disabled,
			Icon:        option.Icon,
			Group:       option.Group,
		}
	}
	return formOptions
}

// Form field enhancement functions

func enhanceFormFieldProps(field molecules.FormFieldProps, formProps EnhancedFormProps) molecules.FormFieldProps {
	// Apply theme integration
	if formProps.ThemeID != "" {
		field.ThemeID = formProps.ThemeID
	}
	
	if formProps.TokenOverrides != nil {
		field.TokenOverrides = formProps.TokenOverrides
	}
	
	field.DarkMode = formProps.DarkMode
	
	// Apply form-level settings
	if formProps.ReadOnly {
		field.Disabled = true
	}
	
	// Ensure validation integration
	if field.ValidationState == "" {
		field.ValidationState = molecules.ValidationStateIdle
	}
	
	// Add form context to Alpine model
	if field.AlpineModel == "" && field.Name != "" {
		field.AlpineModel = fmt.Sprintf("formData.%s", field.Name)
	}
	
	return field
}

// Action management functions

func getActionsByPosition(actions []EnhancedFormAction, position string) []EnhancedFormAction {
	var filtered []EnhancedFormAction
	defaultPosition := utils.IfElse(position == "right", "right", position)
	
	for _, action := range actions {
		actionPosition := utils.IfElse(action.Position != "", action.Position, "right")
		if actionPosition == defaultPosition {
			filtered = append(filtered, action)
		}
	}
	
	return filtered
}

func getActionType(actionType string) string {
	switch actionType {
	case "submit", "reset", "button":
		return actionType
	case "cancel", "save-draft":
		return "button"
	default:
		return "button"
	}
}

func getActionVariant(variant atoms.ButtonVariant) atoms.ButtonVariant {
	if variant != "" {
		return variant
	}
	return atoms.ButtonPrimary
}

func getActionSize(size atoms.ButtonSize) atoms.ButtonSize {
	if size != "" {
		return size
	}
	return atoms.ButtonSizeMD
}

func getActionClickHandler(action EnhancedFormAction, formProps EnhancedFormProps) string {
	if action.OnClick != "" {
		return action.OnClick
	}
	
	switch action.Type {
	case "submit":
		return "submitForm()"
	case "reset":
		return "resetForm()"
	case "cancel":
		return "cancelForm()"
	case "save-draft":
		return "autoSave()"
	default:
		return ""
	}
}

func getActionTarget(actionTarget, formTarget string) string {
	if actionTarget != "" {
		return actionTarget
	}
	return formTarget
}

func getActionSwap(actionSwap, formSwap string) string {
	if actionSwap != "" {
		return actionSwap
	}
	if formSwap != "" {
		return formSwap
	}
	return "innerHTML"
}

func getSubmitButtonText(props EnhancedFormProps) string {
	if props.AutoSave {
		return "Submit"
	}
	return "Save"
}