package organisms

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// ============================================================================
// INTERNAL STATE MANAGEMENT
// ============================================================================

// formState holds processed rendering state (internal only)
type formState struct {
	props           FormProps
	id              string
	config          map[string]any
	actionGroups    map[string][]Action
	dependencyGraph map[string][]Dependency
	flags           formFlags

	// Computed flags for template rendering
	hasHeader         bool
	hasProgress       bool
	hasAutoSave       bool
	hasStorageRestore bool
	hasDebug          bool
}

// formFlags consolidates feature detection
type formFlags struct {
	hasValidation   bool
	hasAutoSave     bool
	hasStorage      bool
	hasProgress     bool
	hasDependencies bool
	hasSSE          bool
	hasAdvanced     bool
	hasHTMX         bool
	hasDebug        bool
}

// ============================================================================
// STATE BUILDING - Convert props to internal state
// ============================================================================

func buildFormState(props FormProps) formState {
	// Generate ID
	id := generateFormID(props.ID, props.Title)

	// Detect enabled features
	flags := detectFeatures(props)

	// Build dependency graph and validate
	depGraph := buildDependencyGraph(props)
	if flags.hasDependencies {
		if err := validateDependencyGraph(depGraph); err != nil {
			fmt.Printf("Form dependency validation warning: %v\n", err)
		}
	}

	// Group actions by position
	actionGroups := groupActionsByPosition(props.Actions)

	// Build Alpine.js configuration
	config := buildAlpineConfig(id, props, flags, depGraph)

	return formState{
		props:             props,
		id:                id,
		config:            config,
		actionGroups:      actionGroups,
		dependencyGraph:   depGraph,
		flags:             flags,
		hasHeader:         props.Title != "" || props.Description != "",
		hasProgress:       flags.hasProgress && props.Progress != nil && props.Progress.TotalSteps > 1,
		hasAutoSave:       flags.hasAutoSave && props.AutoSave != nil,
		hasStorageRestore: flags.hasStorage && props.Storage != nil && props.Storage.RestoreOnLoad,
		hasDebug:          flags.hasDebug && props.Debug != nil && props.Debug.Enabled,
	}
}

func generateFormID(providedID, title string) string {
	if providedID != "" {
		return sanitizeID(providedID)
	}
	if title != "" {
		return sanitizeID(title)
	}
	// Generate from timestamp
	hash := sha256.Sum256([]byte(fmt.Sprintf("form-%d", time.Now().UnixNano())))
	return fmt.Sprintf("form-%x", hash[:8])
}

func sanitizeID(id string) string {
	id = strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_':
			return r
		default:
			return '-'
		}
	}, id)
	return strings.Trim(id, "-")
}

func detectFeatures(props FormProps) formFlags {
	return formFlags{
		hasValidation:   props.Validation != nil && props.Validation.URL != "",
		hasAutoSave:     props.AutoSave != nil && props.AutoSave.Enabled,
		hasStorage:      props.Storage != nil && props.Storage.Strategy != StorageNone,
		hasProgress:     props.Progress != nil && props.Progress.TotalSteps > 1,
		hasDependencies: props.Dependencies != nil && len(props.Dependencies.Graph) > 0,
		hasSSE:          (props.Validation != nil && props.Validation.SSE != nil) || (props.Advanced != nil && props.Advanced.EnableSSE),
		hasAdvanced:     props.Advanced != nil,
		hasHTMX:         props.HTMX != nil,
		hasDebug:        props.Debug != nil && props.Debug.Enabled,
	}
}

func buildDependencyGraph(props FormProps) map[string][]Dependency {
	graph := make(map[string][]Dependency)

	// Add dependencies from config
	if props.Dependencies != nil && props.Dependencies.Graph != nil {
		for field, deps := range props.Dependencies.Graph {
			graph[field] = deps
		}
	}

	// Add field-level dependencies
	allFields := collectAllFields(props)
	for _, field := range allFields {
		if len(field.Dependencies) > 0 {
			graph[field.Name] = append(graph[field.Name], field.Dependencies...)
		}
	}

	return graph
}

func collectAllFields(props FormProps) []Field {
	var fields []Field
	fields = append(fields, props.Fields...)

	for _, section := range props.Sections {
		fields = append(fields, section.Fields...)
	}

	return fields
}

func validateDependencyGraph(graph map[string][]Dependency) error {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var hasCycle func(string) bool
	hasCycle = func(field string) bool {
		visited[field] = true
		recStack[field] = true

		for _, dep := range graph[field] {
			target := dep.TargetField
			if !visited[target] {
				if hasCycle(target) {
					return true
				}
			} else if recStack[target] {
				return true
			}
		}

		recStack[field] = false
		return false
	}

	for field := range graph {
		if !visited[field] && hasCycle(field) {
			return fmt.Errorf("circular dependency detected involving field: %s", field)
		}
	}
	return nil
}

func groupActionsByPosition(actions []Action) map[string][]Action {
	groups := make(map[string][]Action)

	for _, action := range actions {
		position := action.Position
		if position == "" {
			position = "right" // Default position
		}
		groups[position] = append(groups[position], action)
	}

	return groups
}

func buildAlpineConfig(id string, props FormProps, flags formFlags, depGraph map[string][]Dependency) map[string]any {
	config := map[string]any{
		"formId":      id,
		"initialData": props.Fields, // This would be processed to extract initial values

		// Feature flags
		"hasValidation":   flags.hasValidation,
		"hasAutoSave":     flags.hasAutoSave,
		"hasStorage":      flags.hasStorage,
		"hasProgress":     flags.hasProgress,
		"hasDependencies": flags.hasDependencies,
		"hasSSE":          flags.hasSSE,
		"hasAdvanced":     flags.hasAdvanced,
		"hasDebug":        flags.hasDebug,
	}

	// Add validation config
	if flags.hasValidation {
		config["validationURL"] = props.Validation.URL
		config["validationStrategy"] = defaultString(string(props.Validation.Strategy), string(ValidationOnBlur))
		config["validationDebounce"] = defaultInt(props.Validation.Debounce, 300)
	}

	// Add auto-save config
	if flags.hasAutoSave {
		config["autoSave"] = true
		config["autoSaveURL"] = props.AutoSave.URL
		config["autoSaveStrategy"] = defaultString(string(props.AutoSave.Strategy), string(AutoSaveDebounced))
		config["autoSaveInterval"] = defaultInt(props.AutoSave.Interval, 30)
	}

	// Add storage config
	if flags.hasStorage {
		config["storageStrategy"] = string(props.Storage.Strategy)
		config["storageKey"] = defaultString(props.Storage.Key, fmt.Sprintf("form_%s", id))
		config["storageTTL"] = defaultInt(props.Storage.TTL, 86400)
		config["restoreFromStorage"] = props.Storage.RestoreOnLoad
	}

	// Add progress config
	if flags.hasProgress {
		config["currentStep"] = defaultInt(props.Progress.CurrentStep, 1)
		config["totalSteps"] = props.Progress.TotalSteps
	}

	// Add dependency config
	if flags.hasDependencies {
		config["dependencyGraph"] = depGraph
	}

	// Add submission config
	if props.SubmitURL != "" {
		config["submitURL"] = props.SubmitURL
	}

	// Add debug config
	if flags.hasDebug {
		config["debug"] = true
	}

	// Add advanced config
	if flags.hasAdvanced && props.Advanced != nil {
		config["hasAdvanced"] = map[string]bool{
			"enableSSE":          props.Advanced.EnableSSE,
			"enableCrossTabSync": props.Advanced.EnableCrossTabSync,
			"enableOfflineMode":  props.Advanced.EnableOfflineMode,
		}
	}

	return config
}

// ============================================================================
// ATTRIBUTE BUILDERS
// ============================================================================

func buildFormAttributes(state formState) templ.Attributes {
	props := state.props
	attrs := templ.Attributes{
		"id":         state.id,
		"class":      buildFormClasses(state),
		"method":     "POST",
		"novalidate": "true",
	}

	// Alpine.js store initialization
	attrs["x-data"] = fmt.Sprintf("formStore_%s()", state.id)

	// Restoration
	if state.hasStorageRestore {
		attrs["x-init"] = "init()"
	}

	// Submission
	if props.SubmitURL != "" {
		attrs["action"] = props.SubmitURL
	}

	// Event handlers
	if props.OnSubmit != "" {
		attrs["x-on:submit.prevent"] = props.OnSubmit
	} else {
		attrs["x-on:submit.prevent"] = "submitForm($event)"
	}

	// HTMX
	if state.flags.hasHTMX && props.HTMX != nil {
		if props.HTMX.Target != "" {
			attrs["hx-target"] = props.HTMX.Target
		}
		if props.HTMX.Swap != "" {
			attrs["hx-swap"] = props.HTMX.Swap
		}
		if props.HTMX.Trigger != "" {
			attrs["hx-trigger"] = props.HTMX.Trigger
		}
		if props.HTMX.Indicator != "" {
			attrs["hx-indicator"] = props.HTMX.Indicator
		}
	}

	// Theme overrides
	if props.Theme != nil && len(props.Theme.TokenOverrides) > 0 {
		vars := make([]string, 0, len(props.Theme.TokenOverrides))
		for k, v := range props.Theme.TokenOverrides {
			vars = append(vars, fmt.Sprintf("--%s: %s", k, v))
		}
		attrs["style"] = strings.Join(vars, "; ")
	}

	return attrs
}

func buildSectionAttributes(section Section) templ.Attributes {
	attrs := templ.Attributes{
		"class": buildSectionClasses(section),
	}

	if section.Conditional != "" {
		attrs["x-show"] = section.Conditional
		attrs["x-transition"] = "true"
	}

	if section.Collapsible {
		attrs["x-data"] = fmt.Sprintf("{ collapsed: %t }", section.Collapsed)
	}

	return attrs
}

func buildFieldContainerAttributes(field Field, state formState) templ.Attributes {
	attrs := templ.Attributes{
		"class": "form-field-container",
	}

	if field.Conditional != "" {
		attrs["x-show"] = field.Conditional
		attrs["x-transition"] = "true"
	}

	if field.Name != "" {
		attrs["aria-describedby"] = fmt.Sprintf("%s-%s-feedback", state.id, field.Name)
	}

	if len(field.Dependencies) > 0 {
		attrs["x-effect"] = fmt.Sprintf("handleDependencies('%s')", field.Name)
	}

	return attrs
}

// ============================================================================
// CLASS BUILDERS
// ============================================================================

func buildFormClasses(state formState) string {
	props := state.props
	classes := []string{
		"form",
		fmt.Sprintf("form--%s", defaultString(string(props.Layout), string(FormLayoutVertical))),
		fmt.Sprintf("form--%s", defaultString(string(props.Size), string(FormSizeMD))),
	}

	// State classes
	if props.ReadOnly {
		classes = append(classes, "form--readonly")
	}
	if state.hasProgress {
		classes = append(classes, "form--with-progress")
	}
	if state.hasAutoSave {
		classes = append(classes, "form--autosave")
	}
	if state.hasStorageRestore {
		classes = append(classes, "form--with-storage")
	}
	if state.flags.hasSSE {
		classes = append(classes, "form--with-sse")
	}
	if state.flags.hasDependencies {
		classes = append(classes, "form--with-dependencies")
	}
	if props.Theme != nil && props.Theme.DarkMode {
		classes = append(classes, "form--dark")
	}
	if state.hasDebug {
		classes = append(classes, "form--debug")
	}

	// Custom class
	if props.ClassName != "" {
		classes = append(classes, props.ClassName)
	}

	return strings.Join(classes, " ")
}

func buildSectionClasses(section Section) string {
	classes := []string{"form-section"}

	if section.Collapsible {
		classes = append(classes, "form-section--collapsible")
	}
	if section.Collapsed {
		classes = append(classes, "form-section--collapsed")
	}
	if section.Required {
		classes = append(classes, "form-section--required")
	}
	if section.Layout != "" {
		classes = append(classes, fmt.Sprintf("form-section--%s", section.Layout))
	}

	return strings.Join(classes, " ")
}

func buildFieldsContainerClasses(props FormProps) string {
	classes := []string{"form-fields"}

	if props.Layout == FormLayoutGrid && len(props.Fields) > 0 {
		// Could be enhanced to specify grid columns
		classes = append(classes, "form-fields--grid")
	}

	return strings.Join(classes, " ")
}

func buildActionsClasses(position string) string {
	classes := []string{"form-actions"}

	if position != "" {
		classes = append(classes, fmt.Sprintf("form-actions--%s", position))
	}

	return strings.Join(classes, " ")
}

// ============================================================================
// FIELD ENHANCEMENT
// ============================================================================

func enhanceFieldProps(field Field, state formState) molecules.FormFieldProps {
	enhanced := field.FormFieldProps

	// Skip if no name
	if enhanced.Name == "" {
		return enhanced
	}

	// Alpine.js model binding
	// This would need to be added to FormFieldProps or handled differently

	// Add validation attributes based on strategy
	if state.flags.hasValidation && state.props.Validation != nil {
		// This would be implemented based on how FormFieldProps supports Alpine.js
		// For now, we'll assume it supports x-model and x-on attributes
	}

	// Read-only state
	if state.props.ReadOnly {
		enhanced.Readonly = true
	}

	return enhanced
}

// ============================================================================
// BUTTON PROPS MERGING
// ============================================================================

func mergeButtonProps(custom, defaults atoms.ButtonProps) atoms.ButtonProps {
	// Start with defaults
	merged := defaults

	// Override with non-zero custom values
	if custom.ID != "" {
		merged.ID = custom.ID
	}
	if custom.Type != "" {
		merged.Type = custom.Type
	}
	if custom.Text != "" {
		merged.Text = custom.Text
	}
	if custom.Variant != "" {
		merged.Variant = custom.Variant
	}
	if custom.Size != "" {
		merged.Size = custom.Size
	}
	if custom.OnClick != "" {
		merged.OnClick = custom.OnClick
	}
	if custom.ClassName != "" {
		if defaults.ClassName != "" {
			merged.ClassName = defaults.ClassName + " " + custom.ClassName
		} else {
			merged.ClassName = custom.ClassName
		}
	}
	if custom.Disabled {
		merged.Disabled = custom.Disabled
	}

	return merged
}

// ============================================================================
// UTILITY FUNCTIONS
// ============================================================================

func defaultString(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func defaultInt(value, defaultValue int) int {
	if value == 0 {
		return defaultValue
	}
	return value
}

func mustMarshalJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// HasElements checks if slice has elements
func HasElements[T any](slice []T) bool {
	return len(slice) > 0
}

// Coalesce returns first non-empty string
func Coalesce(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

// Default returns value if non-zero, otherwise defaultValue
func Default[T comparable](value, defaultValue T) T {
	var zero T
	if value == zero {
		return defaultValue
	}
	return value
}
