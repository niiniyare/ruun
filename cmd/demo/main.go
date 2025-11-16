package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/components/organisms"
	"github.com/niiniyare/ruun/views/validation"
)

// DemoServer serves the schema-driven UI demo
type DemoServer struct {
	renderer     *views.TemplateRenderer
	stateManager *views.StateManager
	orchestrator *validation.UIValidationOrchestrator
}

// DemoFormData represents the form submission data
type DemoFormData struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
	Country    string `json:"country"`
	Newsletter bool   `json:"newsletter"`
	Bio        string `json:"bio"`
}

// NewDemoServer creates a new demo server instance
func NewDemoServer() *DemoServer {
	// Create demo schema
	demoSchema := createDemoSchema()

	// Initialize state manager
	stateManager := views.NewStateManager(demoSchema, make(map[string]any))

	// Initialize validation orchestrator
	validationStateMgr := validation.NewValidationStateManager(stateManager)
	orchestrator := validation.NewUIValidationOrchestrator(
		demoSchema,
		nil, // Using built-in validation for demo
		validationStateMgr,
	)

	// Initialize field mapper with theme support
	mapper := views.NewFieldMapper(demoSchema, "en")

	// Initialize component registry
	registry := views.NewComponentRegistry()

	// Register default renderers for demo
	registry.Register(schema.FieldText, &views.DefaultFieldRenderer{})
	registry.Register(schema.FieldEmail, &views.DefaultFieldRenderer{})
	registry.Register(schema.FieldNumber, &views.DefaultFieldRenderer{})
	registry.Register(schema.FieldSelect, &views.DefaultFieldRenderer{})
	registry.Register(schema.FieldCheckbox, &views.DefaultFieldRenderer{})
	registry.Register(schema.FieldTextarea, &views.DefaultFieldRenderer{})

	// Initialize template renderer
	renderer := views.NewTemplateRenderer(
		registry,
		mapper,
		demoSchema,
		stateManager,
		orchestrator,
	)

	return &DemoServer{
		renderer:     renderer,
		stateManager: stateManager,
		orchestrator: orchestrator,
	}
}

func createDemoSchema() *schema.Schema {
	return &schema.Schema{
		ID:          "demo-form",
		Title:       "User Registration Demo",
		Type:        schema.TypeForm,
		Version:     "1.0.0",
		Description: "Real-time schema-driven UI components demo",
		Fields: []schema.Field{
			{
				Name:        "username",
				Type:        "text",
				Label:       "Username",
				Description: "Choose a unique username (3-20 characters)",
				Required:    true,
				Validation: &schema.FieldValidation{
					MinLength: ptr(3),
					MaxLength: ptr(20),
					Pattern:   "^[a-zA-Z0-9_-]+$",
				},
			},
			{
				Name:        "email",
				Type:        "email",
				Label:       "Email Address",
				Description: "We'll never share your email",
				Required:    true,
				Validation: &schema.FieldValidation{
					Format: "email",
				},
			},
			{
				Name:        "age",
				Type:        "number",
				Label:       "Age",
				Description: "You must be 18 or older",
				Required:    true,
				Validation: &schema.FieldValidation{
					Min: ptrFloat(18.0),
					Max: ptrFloat(120.0),
				},
			},
			{
				Name:        "country",
				Type:        "select",
				Label:       "Country",
				Description: "Select your country",
				Required:    false,
				Options: []schema.Option{
					{Value: "US", Label: "United States"},
					{Value: "CA", Label: "Canada"},
					{Value: "UK", Label: "United Kingdom"},
					{Value: "DE", Label: "Germany"},
					{Value: "FR", Label: "France"},
					{Value: "AU", Label: "Australia"},
					{Value: "JP", Label: "Japan"},
					{Value: "Other", Label: "Other"},
				},
			},
			{
				Name:        "newsletter",
				Type:        "checkbox",
				Label:       "Subscribe to Newsletter",
				Description: "Get the latest updates and news",
				Required:    false,
			},
			{
				Name:        "bio",
				Type:        "textarea",
				Label:       "Bio",
				Description: "Tell us about yourself (optional)",
				Required:    false,
				Validation: &schema.FieldValidation{
					MaxLength: ptr(500),
				},
			},
		},
	}
}

func (s *DemoServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Ruun Schema-Driven UI Demo</title>
    <script src="https://unpkg.com/htmx.org@1.9.8"></script>
    <script src="https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js" defer></script>
    <style>` + getEmbeddedCSS() + `</style>
</head>
<body>
    <div class="demo-container">
        <!-- Header -->
        <header class="demo-header">
            <div>
                <h1 class="demo-title">Ruun Schema-Driven UI Demo</h1>
                <p class="demo-subtitle">Real Templ Components • Token-Based Theming • Live Validation</p>
            </div>
            <div class="demo-controls">
                <button class="btn btn-secondary" onclick="resetForm()">Reset</button>
                <button class="btn btn-primary" 
                        hx-post="/validate-all" 
                        hx-target="#form-container"
                        hx-swap="innerHTML">
                    Validate All
                </button>
                <button class="theme-toggle" onclick="toggleTheme()" title="Toggle theme">
                    <span id="theme-icon">🌙</span>
                </button>
            </div>
        </header>

        <!-- Main Content -->
        <main class="demo-main">
            <!-- Live Form Panel -->
            <div class="demo-panel">
                <div class="panel-header">
                    <h2 class="panel-title">Generated Templ Components</h2>
                    <span class="panel-badge">Live</span>
                </div>
                
                <div id="form-container" class="form-container">
                    <!-- Form will be loaded here via HTMX -->
                </div>
                
                <!-- Load initial form -->
                <script>
                    htmx.onLoad(function() {
                        htmx.ajax('GET', '/form', {target:'#form-container'});
                    });
                </script>
            </div>

            <!-- Schema & State Panel -->
            <div class="demo-panel">
                <div class="panel-header">
                    <h2 class="panel-title">Schema & State</h2>
                    <span class="panel-badge">Real-time</span>
                </div>
                
                <!-- Schema Display -->
                <div class="state-viewer">
                    <h3>JSON Schema Definition</h3>
                    <pre id="schema-display"></pre>
                </div>

                <!-- Form State -->
                <div class="state-viewer" 
                     hx-get="/state" 
                     hx-trigger="every 1s"
                     hx-swap="innerHTML">
                    <h3>Current Form State</h3>
                    <pre>Loading...</pre>
                </div>

                <!-- Validation State -->
                <div class="state-viewer"
                     hx-get="/validation-state"
                     hx-trigger="every 1s"
                     hx-swap="innerHTML">
                    <h3>Validation Status</h3>
                    <pre>Loading...</pre>
                </div>

                <!-- Token Resolution -->
                <div class="state-viewer"
                     hx-get="/tokens"
                     hx-trigger="every 2s"
                     hx-swap="innerHTML">
                    <h3>Resolved Design Tokens</h3>
                    <pre>Loading...</pre>
                </div>
            </div>
        </main>
    </div>

    <script>` + getEmbeddedJS() + `</script>
</body>
</html>`

	w.Write([]byte(html))
}

func (s *DemoServer) handleForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Get current theme from context
	theme := r.Header.Get("X-Theme")
	if theme == "" {
		theme = "light"
	}

	// Convert schema.Schema to SchemaFormDefinition
	schemaFormDef := organisms.SchemaFormDefinition{
		Name:        s.renderer.GetSchema().ID,
		Title:       s.renderer.GetSchema().Title,
		Description: s.renderer.GetSchema().Description,
		Method:      "POST",
		Action:      "/submit",
		Fields:      convertSchemaFields(s.renderer.GetSchema().Fields),
	}

	// Create schema form component
	schemaFormProps := organisms.SchemaFormProps{
		Schema: schemaFormDef,
		Data:   s.stateManager.GetAllValues(),
		Errors: convertErrorsToStringMap(s.stateManager.GetAllErrors()),
		Locale: "en",
		ID:     "demo-form",
		Class:  "demo-schema-form",
	}

	// Render the form using real Templ components
	formHTML, err := s.renderer.RenderSchemaForm(ctx, schemaFormProps)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error rendering form: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(formHTML))
}

func (s *DemoServer) handleValidateField(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	fieldName := r.FormValue("field")
	fieldValue := r.FormValue("value")

	if fieldName == "" {
		http.Error(w, "Field name required", http.StatusBadRequest)
		return
	}

	// Update form state
	s.stateManager.SetValue(fieldName, fieldValue)
	s.stateManager.SetFieldTouched(fieldName, true)
	s.stateManager.SetFieldDirty(fieldName, true)

	// Perform validation using orchestrator
	_ = s.orchestrator.ValidateFieldWithUI(ctx, fieldName, fieldValue, validation.ValidationTriggerDebounced)

	// Get field definition
	var field *schema.Field
	for _, f := range s.renderer.GetSchema().Fields {
		if f.Name == fieldName {
			field = &f
			break
		}
	}

	if field == nil {
		http.Error(w, "Field not found", http.StatusNotFound)
		return
	}

	// Get validation state
	validationState := s.orchestrator.GetValidationStateForField(fieldName)
	errors := s.stateManager.GetFieldErrors(fieldName)
	touched := s.stateManager.IsFieldTouched(fieldName)
	dirty := s.stateManager.IsFieldDirty(fieldName)

	// Re-render the field with updated validation state
	fieldHTML, renderErr := s.renderer.RenderField(ctx, field, fieldValue, errors, touched, dirty)
	if renderErr != nil {
		http.Error(w, fmt.Sprintf("Error rendering field: %v", renderErr), http.StatusInternalServerError)
		return
	}

	// Return updated field HTML with HTMX-friendly response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Add validation state as HTMX trigger for client-side updates
	w.Header().Set("HX-Trigger", fmt.Sprintf(`{"fieldValidated": {"field": "%s", "state": "%s", "hasErrors": %t}}`,
		fieldName, validationState.String(), len(errors) > 0))

	w.Write([]byte(fieldHTML))
}

func (s *DemoServer) handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Validate all fields
	allValid := true
	validationResults := make(map[string]interface{})

	for _, field := range s.renderer.GetSchema().Fields {
		value := r.FormValue(field.Name)

		// Handle checkbox fields
		if field.Type == "checkbox" {
			value = r.FormValue(field.Name)
			if value == "" {
				value = "false"
			}
		}

		// Update state
		s.stateManager.SetValue(field.Name, value)
		s.stateManager.SetFieldTouched(field.Name, true)

		// Validate
		err := s.orchestrator.ValidateFieldWithUI(ctx, field.Name, value, validation.ValidationTriggerOnSubmit)
		fieldValid := err == nil && len(s.stateManager.GetFieldErrors(field.Name)) == 0

		validationResults[field.Name] = map[string]interface{}{
			"valid":  fieldValid,
			"value":  value,
			"errors": s.stateManager.GetFieldErrors(field.Name),
			"state":  s.orchestrator.GetValidationStateForField(field.Name).String(),
		}

		if !fieldValid {
			allValid = false
		}
	}

	// Prepare response
	response := map[string]interface{}{
		"success":    allValid,
		"message":    getMessage(allValid),
		"validation": validationResults,
		"timestamp":  fmt.Sprintf("%d", time.Now().UnixMilli()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *DemoServer) handleState(w http.ResponseWriter, r *http.Request) {
	formState := s.stateManager.GetAllValues()

	stateHTML := fmt.Sprintf(`
		<h3>Current Form State</h3>
		<pre>%s</pre>
	`, formatJSON(formState))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(stateHTML))
}

func (s *DemoServer) handleValidationState(w http.ResponseWriter, r *http.Request) {
	validationState := make(map[string]interface{})

	for _, field := range s.renderer.GetSchema().Fields {
		validationState[field.Name] = map[string]interface{}{
			"state":      s.orchestrator.GetValidationStateForField(field.Name).String(),
			"errors":     s.stateManager.GetFieldErrors(field.Name),
			"touched":    s.stateManager.IsFieldTouched(field.Name),
			"dirty":      s.stateManager.IsFieldDirty(field.Name),
			"validating": s.orchestrator.IsFieldValidating(field.Name),
		}
	}

	stateHTML := fmt.Sprintf(`
		<h3>Validation Status</h3>
		<pre>%s</pre>
	`, formatJSON(validationState))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(stateHTML))
}

func (s *DemoServer) handleTokens(w http.ResponseWriter, r *http.Request) {
	// Sample token resolution for demo
	sampleProps := molecules.FormFieldProps{
		Name:            "sample",
		Type:            "text",
		ValidationState: validation.ValidationStateValid,
		ThemeID:         r.Header.Get("X-Theme"),
	}

	darkMode := r.Header.Get("X-Theme") == "dark"
	tokens := s.renderer.GetFieldMapper().ResolveFieldTokens(r.Context(), sampleProps, darkMode)

	tokenInfo := map[string]interface{}{
		"theme":           r.Header.Get("X-Theme"),
		"darkMode":        darkMode,
		"resolvedTokens":  tokens,
		"tokenCount":      len(tokens),
		"validationState": sampleProps.ValidationState.String(),
	}

	tokensHTML := fmt.Sprintf(`
		<h3>Resolved Design Tokens</h3>
		<pre>%s</pre>
	`, formatJSON(tokenInfo))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tokensHTML))
}

// Helper functions

func ptr(i int) *int              { return &i }
func ptrFloat(f float64) *float64 { return &f }

func formatJSON(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error formatting: %v", err)
	}
	return string(b)
}

func getMessage(success bool) string {
	if success {
		return "✅ Form validation successful! All fields are valid."
	}
	return "❌ Form validation failed. Please check the errors above."
}

// convertSchemaFields converts schema.Field to organisms.SchemaField
func convertSchemaFields(fields []schema.Field) []organisms.SchemaField {
	result := make([]organisms.SchemaField, len(fields))
	for i, field := range fields {
		result[i] = organisms.SchemaField{
			Name:        field.Name,
			Type:        string(field.Type),
			Label:       field.Label,
			Description: field.Description,
			Required:    field.Required,
			Options:     convertSchemaOptions(field.Options),
		}

		// Convert validation rules
		if field.Validation != nil {
			if field.Validation.MinLength != nil {
				result[i].MinLength = *field.Validation.MinLength
			}
			if field.Validation.MaxLength != nil {
				result[i].MaxLength = *field.Validation.MaxLength
			}
			result[i].Pattern = field.Validation.Pattern
			if field.Validation.Min != nil {
				result[i].Min = *field.Validation.Min
			}
			if field.Validation.Max != nil {
				result[i].Max = *field.Validation.Max
			}
		}
	}
	return result
}

// convertSchemaOptions converts schema.Option to organisms.SchemaFieldOption
func convertSchemaOptions(options []schema.Option) []organisms.SchemaFieldOption {
	result := make([]organisms.SchemaFieldOption, len(options))
	for i, option := range options {
		result[i] = organisms.SchemaFieldOption{
			Value: option.Value,
			Label: option.Label,
		}
	}
	return result
}

// convertErrorsToStringMap converts map[string][]string to map[string]string
func convertErrorsToStringMap(errors map[string][]string) map[string]string {
	result := make(map[string]string)
	for key, errorList := range errors {
		if len(errorList) > 0 {
			result[key] = errorList[0] // Take first error
		}
	}
	return result
}

func main() {
	server := NewDemoServer()

	// Routes
	http.HandleFunc("/", server.handleIndex)
	http.HandleFunc("/form", server.handleForm)
	http.HandleFunc("/validate-field", server.handleValidateField)
	http.HandleFunc("/submit", server.handleSubmit)
	http.HandleFunc("/state", server.handleState)
	http.HandleFunc("/validation-state", server.handleValidationState)
	http.HandleFunc("/tokens", server.handleTokens)

	port := ":8080"
	fmt.Printf("🚀 Ruun Schema-Driven UI Demo starting on http://localhost%s\n", port)
	fmt.Printf("📋 Features:\n")
	fmt.Printf("  • Real Templ component rendering\n")
	fmt.Printf("  • Token-based theming system\n")
	fmt.Printf("  • Live validation orchestration\n")
	fmt.Printf("  • HTMX real-time updates\n")
	fmt.Printf("  • Alpine.js client interactions\n")

	log.Fatal(http.ListenAndServe(port, nil))
}
