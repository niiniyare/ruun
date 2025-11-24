// pkg/schema/i18n_v2.go
package schema

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// I18nManager manages translations for the application
type I18nManager struct {
	config       *I18nConfig
	translations map[string]*Translation // locale -> translation
	currentLocale string                  // current active locale
	mu           sync.RWMutex

	// Public fields for backward compatibility
	Enabled          bool     `json:"enabled"`
	DefaultLocale    string   `json:"defaultLocale"`
	SupportedLocales []string `json:"supportedLocales"`
	
	// Direct message access for testing
	messages map[string]map[string]string
}

// I18nConfig configures the i18n system
type I18nConfig struct {
	DefaultLocale    string   // Default language (e.g., "en")
	FallbackLocale   string   // Fallback language (e.g., "en")
	SupportedLocales []string // Supported languages
	TranslationsPath string   // Path to translation files
	FilePattern      string   // File pattern (e.g., "{locale}.json")
	EnableCache      bool     // Enable translation caching
	EnablePlurals    bool     // Enable pluralization
}

// Translation holds translations for a single locale
type Translation struct {
	Locale       string                 `json:"locale"`
	Translations map[string]interface{} `json:"translations"`
	metadata     *TranslationMetadata
	mu           sync.RWMutex
}

// TranslationMetadata holds translation file metadata
type TranslationMetadata struct {
	BaseMetadata                     // Embedded base metadata
	Author      string `json:"author"`      // Translation author
	LastUpdated string `json:"lastUpdated"` // Last updated timestamp (string format for compatibility)
	Description string `json:"description"` // Translation description
}

// TranslateParams holds parameters for translation
type TranslateParams struct {
	Key          string         // Translation key (e.g., "common.welcome")
	Params       map[string]any // Parameters for interpolation
	Count        *int           // Count for pluralization
	DefaultValue string         // Default value if translation not found
}

// NewI18nManager creates a new i18n manager
func NewI18nManager(config *I18nConfig) *I18nManager {
	if config == nil {
		config = DefaultI18nConfig()
	}

	manager := &I18nManager{
		config:           config,
		translations:     make(map[string]*Translation),
		currentLocale:    config.DefaultLocale,
		Enabled:          true,
		DefaultLocale:    config.DefaultLocale,
		SupportedLocales: config.SupportedLocales,
		messages:         make(map[string]map[string]string),
	}

	return manager
}

// DefaultI18nConfig returns default configuration
func DefaultI18nConfig() *I18nConfig {
	return &I18nConfig{
		DefaultLocale:    "en",
		FallbackLocale:   "en",
		SupportedLocales: []string{"en"},
		TranslationsPath: "./locales",
		FilePattern:      "{locale}.json",
		EnableCache:      true,
		EnablePlurals:    true,
	}
}

// LoadTranslations loads all translation files
func (m *I18nManager) LoadTranslations() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, locale := range m.config.SupportedLocales {
		if err := m.loadTranslationFile(locale); err != nil {
			return fmt.Errorf("load locale %s: %w", locale, err)
		}
	}

	return nil
}

// LoadTranslation loads a specific translation file
func (m *I18nManager) LoadTranslation(locale string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.loadTranslationFile(locale)
}

// loadTranslationFile loads a translation file (must be called with lock held)
func (m *I18nManager) loadTranslationFile(locale string) error {
	filename := m.getTranslationFilePath(locale)

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read file: %w", err)
	}

	var translation Translation
	if err := json.Unmarshal(data, &translation); err != nil {
		return fmt.Errorf("unmarshal json: %w", err)
	}

	translation.Locale = locale
	m.translations[locale] = &translation

	return nil
}

// LoadFromFS loads translations from an embedded filesystem
func (m *I18nManager) LoadFromFS(fsys fs.FS, pattern string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, locale := range m.config.SupportedLocales {
		filename := strings.ReplaceAll(pattern, "{locale}", locale)

		data, err := fs.ReadFile(fsys, filename)
		if err != nil {
			return fmt.Errorf("read file %s: %w", filename, err)
		}

		var translation Translation
		if err := json.Unmarshal(data, &translation); err != nil {
			return fmt.Errorf("unmarshal json: %w", err)
		}

		translation.Locale = locale
		m.translations[locale] = &translation
	}

	return nil
}

// LoadFromMap loads translations from a map
func (m *I18nManager) LoadFromMap(locale string, translations map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.translations[locale] = &Translation{
		Locale:       locale,
		Translations: translations,
	}

	return nil
}

// T translates a key with parameters
func (m *I18nManager) T(locale, key string, params map[string]any) string {
	return m.Translate(&TranslateParams{
		Key:    key,
		Params: params,
	}, locale)
}

// Translate translates with full options
func (m *I18nManager) Translate(params *TranslateParams, locale string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// Get translation for locale
	translation := m.getTranslation(locale, params.Key)
	if translation == "" {
		// Try fallback locale
		if locale != m.config.FallbackLocale {
			translation = m.getTranslation(m.config.FallbackLocale, params.Key)
		}

		// Use default value if still not found
		if translation == "" {
			if params.DefaultValue != "" {
				return params.DefaultValue
			}
			return params.Key // Return key as last resort
		}
	}

	// Handle pluralization
	if m.config.EnablePlurals && params.Count != nil {
		translation = m.selectPluralForm(translation, *params.Count)
	}

	// Interpolate parameters
	if params.Params != nil {
		translation = m.interpolate(translation, params.Params)
	}

	return translation
}

// TCount translates with pluralization
func (m *I18nManager) TCount(locale, key string, count int, params map[string]any) string {
	if params == nil {
		params = make(map[string]any)
	}
	params["count"] = count

	return m.Translate(&TranslateParams{
		Key:    key,
		Params: params,
		Count:  &count,
	}, locale)
}

// getTranslation gets a translation by key
func (m *I18nManager) getTranslation(locale, key string) string {
	trans, ok := m.translations[locale]
	if !ok {
		return ""
	}

	return trans.Get(key)
}

// getTranslationFilePath returns the path to a translation file
func (m *I18nManager) getTranslationFilePath(locale string) string {
	filename := strings.ReplaceAll(m.config.FilePattern, "{locale}", locale)
	return filepath.Join(m.config.TranslationsPath, filename)
}

// selectPluralForm selects the appropriate plural form
func (m *I18nManager) selectPluralForm(translation string, count int) string {
	// Check if translation is a plural object
	if !strings.Contains(translation, "|") {
		return translation
	}

	// Split plural forms: "no items|one item|{count} items"
	forms := strings.Split(translation, "|")

	switch {
	case count == 0 && len(forms) > 0:
		return strings.TrimSpace(forms[0])
	case count == 1 && len(forms) > 1:
		return strings.TrimSpace(forms[1])
	default:
		if len(forms) > 2 {
			return strings.TrimSpace(forms[2])
		}
		if len(forms) > 1 {
			return strings.TrimSpace(forms[1])
		}
		return strings.TrimSpace(forms[0])
	}
}

// interpolate replaces parameters in translation string
func (m *I18nManager) interpolate(translation string, params map[string]any) string {
	result := translation

	for key, value := range params {
		placeholder := fmt.Sprintf("{%s}", key)
		replacement := fmt.Sprintf("%v", value)
		result = strings.ReplaceAll(result, placeholder, replacement)
	}

	return result
}

// HasLocale checks if a locale is supported
func (m *I18nManager) HasLocale(locale string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.translations[locale]
	return ok
}

// GetSupportedLocales returns list of supported locales
func (m *I18nManager) GetSupportedLocales() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return append([]string{}, m.config.SupportedLocales...)
}

// SetLocale sets the current active locale
func (m *I18nManager) SetLocale(locale string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	m.currentLocale = locale
}

// LoadDefaultTranslations loads default translations for supported locales
func (m *I18nManager) LoadDefaultTranslations() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Load translations for each supported locale
	for _, locale := range m.config.SupportedLocales {
		if _, exists := m.translations[locale]; !exists {
			// Create basic default translations
			translations := map[string]interface{}{
				"common": map[string]interface{}{
					"save":   "Save",
					"cancel": "Cancel",
					"ok":     "OK",
				},
				"validation": map[string]interface{}{
					"required":       "{{.field}} is required",
					"min_length":     "{{.field}} must be at least {{.min}} characters",
					"invalid_email":  "Please enter a valid email address",
					"invalid_range":  "{{.field}} must be between {{.min}} and {{.max}}",
				},
				"actions": map[string]interface{}{
					"save":   "Save",
					"cancel": "Cancel",
					"submit": "Submit",
					"delete": "Delete",
				},
				"messages": map[string]interface{}{
					"save_success": "Changes saved successfully",
					"save_error":   "Failed to save changes",
					"loading":      "Loading...",
				},
			}
			
			// Add Spanish translations
			if locale == "es" {
				translations = map[string]interface{}{
					"common": map[string]interface{}{
						"save":   "Guardar",
						"cancel": "Cancelar",
						"ok":     "OK",
					},
					"validation": map[string]interface{}{
						"required":       "{{.field}} es requerido",
						"min_length":     "{{.field}} debe tener al menos {{.min}} caracteres",
						"invalid_email":  "Por favor ingresa una dirección de correo válida",
						"invalid_range":  "{{.field}} debe estar entre {{.min}} y {{.max}}",
					},
					"actions": map[string]interface{}{
						"save":   "Guardar",
						"cancel": "Cancelar",
						"submit": "Enviar",
						"delete": "Eliminar",
					},
					"messages": map[string]interface{}{
						"save_success": "Cambios guardados exitosamente",
						"save_error":   "Error al guardar cambios",
						"loading":      "Cargando...",
					},
				}
			}
			
			m.translations[locale] = &Translation{
				Locale:       locale,
				Translations: translations,
			}
		}
	}
	return nil
}

// GetValidationMessage returns a validation message for a specific error type
func (m *I18nManager) GetValidationMessage(errorType string, params map[string]any) string {
	key := fmt.Sprintf("validation.%s", errorType)
	locale := m.GetLocale()
	return m.T(locale, key, params)
}

// interpolateMessage interpolates parameters into a message template
func (m *I18nManager) interpolateMessage(message string, params map[string]any) string {
	return m.interpolate(message, params)
}

// GetMessage gets a translation message for the current locale
func (m *I18nManager) GetMessage(key string, params map[string]any) string {
	locale := m.GetLocale()
	return m.T(locale, key, params)
}

// GetPluralMessage gets a pluralized message
func (m *I18nManager) GetPluralMessage(key string, count int, locale string) string {
	params := map[string]any{"count": count}
	return m.TCount(locale, key, count, params)
}

// DetectLocale detects best locale from preferences
func (m *I18nManager) DetectLocale(preferences []string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, pref := range preferences {
		// Try exact match
		if m.hasLocaleUnsafe(pref) {
			return pref
		}

		// Try language code only (e.g., "en" from "en-US")
		if len(pref) > 2 {
			langCode := pref[:2]
			if m.hasLocaleUnsafe(langCode) {
				return langCode
			}
		}

		// Try matching by language prefix
		for _, locale := range m.config.SupportedLocales {
			if strings.HasPrefix(locale, pref[:2]) {
				return locale
			}
		}
	}

	return m.config.DefaultLocale
}

// hasLocaleUnsafe checks locale without locking
func (m *I18nManager) hasLocaleUnsafe(locale string) bool {
	_, ok := m.translations[locale]
	return ok
}

// ReloadTranslations reloads all translations
func (m *I18nManager) ReloadTranslations() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.translations = make(map[string]*Translation)

	for _, locale := range m.config.SupportedLocales {
		if err := m.loadTranslationFile(locale); err != nil {
			return fmt.Errorf("reload locale %s: %w", locale, err)
		}
	}

	return nil
}

// AddLocale adds a new supported locale
func (m *I18nManager) AddLocale(locale string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Check if already supported
	for _, l := range m.config.SupportedLocales {
		if l == locale {
			return nil
		}
	}

	m.config.SupportedLocales = append(m.config.SupportedLocales, locale)
	return m.loadTranslationFile(locale)
}

// Translation methods

// Get gets a translation by key (supports dot notation)
func (t *Translation) Get(key string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.getValueByPath(t.Translations, key)
}

// getValueByPath gets value from nested map using dot notation
func (t *Translation) getValueByPath(data map[string]interface{}, path string) string {
	parts := strings.Split(path, ".")

	var current interface{} = data
	for _, part := range parts {
		switch v := current.(type) {
		case map[string]interface{}:
			current = v[part]
		default:
			return ""
		}

		if current == nil {
			return ""
		}
	}

	// Convert to string
	switch v := current.(type) {
	case string:
		return v
	case map[string]interface{}:
		// If it's still a map, try to get default value
		if def, ok := v["_default"].(string); ok {
			return def
		}
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Set sets a translation value
func (t *Translation) Set(key, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.setValueByPath(t.Translations, key, value)
}

// setValueByPath sets value in nested map using dot notation
func (t *Translation) setValueByPath(data map[string]interface{}, path, value string) {
	parts := strings.Split(path, ".")

	current := data
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]

		if _, ok := current[part]; !ok {
			current[part] = make(map[string]interface{})
		}

		if next, ok := current[part].(map[string]interface{}); ok {
			current = next
		} else {
			return // Path is invalid
		}
	}

	current[parts[len(parts)-1]] = value
}

// ToJSON converts translation to JSON
func (t *Translation) ToJSON() ([]byte, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return json.MarshalIndent(t, "", "  ")
}

// I18nBuilder provides fluent API for building i18n manager
type I18nBuilder struct {
	config *I18nConfig
}

// NewI18n starts building i18n manager
func NewI18n() *I18nBuilder {
	return &I18nBuilder{
		config: DefaultI18nConfig(),
	}
}

func (b *I18nBuilder) WithDefaultLocale(locale string) *I18nBuilder {
	b.config.DefaultLocale = locale
	return b
}

func (b *I18nBuilder) WithFallbackLocale(locale string) *I18nBuilder {
	b.config.FallbackLocale = locale
	return b
}

func (b *I18nBuilder) WithSupportedLocales(locales ...string) *I18nBuilder {
	b.config.SupportedLocales = locales
	return b
}

func (b *I18nBuilder) WithTranslationsPath(path string) *I18nBuilder {
	b.config.TranslationsPath = path
	return b
}

func (b *I18nBuilder) WithFilePattern(pattern string) *I18nBuilder {
	b.config.FilePattern = pattern
	return b
}

func (b *I18nBuilder) WithCache(enable bool) *I18nBuilder {
	b.config.EnableCache = enable
	return b
}

func (b *I18nBuilder) WithPlurals(enable bool) *I18nBuilder {
	b.config.EnablePlurals = enable
	return b
}

func (b *I18nBuilder) Build() *I18nManager {
	return NewI18nManager(b.config)
}

func (b *I18nBuilder) BuildAndLoad() (*I18nManager, error) {
	manager := NewI18nManager(b.config)
	if err := manager.LoadTranslations(); err != nil {
		return nil, err
	}
	return manager, nil
}

// Localizer provides a convenient interface for a specific locale
type Localizer struct {
	manager *I18nManager
	locale  string
}

// NewLocalizer creates a localizer for a specific locale
func (m *I18nManager) NewLocalizer(locale string) *Localizer {
	return &Localizer{
		manager: m,
		locale:  locale,
	}
}

// T translates a key
func (l *Localizer) T(key string, params ...map[string]any) string {
	var p map[string]any
	if len(params) > 0 {
		p = params[0]
	}
	return l.manager.T(l.locale, key, p)
}

// TCount translates with count
func (l *Localizer) TCount(key string, count int, params ...map[string]any) string {
	var p map[string]any
	if len(params) > 0 {
		p = params[0]
	}
	return l.manager.TCount(l.locale, key, count, p)
}

// TDefault translates with default value
func (l *Localizer) TDefault(key, defaultValue string, params ...map[string]any) string {
	var p map[string]any
	if len(params) > 0 {
		p = params[0]
	}
	return l.manager.Translate(&TranslateParams{
		Key:          key,
		Params:       p,
		DefaultValue: defaultValue,
	}, l.locale)
}

// GetLocale returns current locale
func (l *Localizer) GetLocale() string {
	return l.locale
}

// TranslationValidator validates translation files
type TranslationValidator struct {
	requiredKeys []string
}

// NewTranslationValidator creates a validator
func NewTranslationValidator(requiredKeys ...string) *TranslationValidator {
	return &TranslationValidator{
		requiredKeys: requiredKeys,
	}
}

// Validate validates a translation
func (v *TranslationValidator) Validate(translation *Translation) error {
	missing := []string{}

	for _, key := range v.requiredKeys {
		if translation.Get(key) == "" {
			missing = append(missing, key)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required keys: %v", missing)
	}

	return nil
}

// ValidateAll validates all translations
func (v *TranslationValidator) ValidateAll(manager *I18nManager) error {
	errors := []string{}

	for _, locale := range manager.GetSupportedLocales() {
		if trans, ok := manager.translations[locale]; ok {
			if err := v.Validate(trans); err != nil {
				errors = append(errors, fmt.Sprintf("%s: %v", locale, err))
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("validation errors: %v", errors)
	}

	return nil
}

// TranslationExporter exports translations
type TranslationExporter struct{}

// ExportToJSON exports translation to JSON file
func (e *TranslationExporter) ExportToJSON(translation *Translation, filepath string) error {
	data, err := translation.ToJSON()
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(filepath, data, 0o644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

// FormatMessage formats a message with parameters (advanced formatting)
type MessageFormatter struct {
	paramPattern *regexp.Regexp
}

// NewMessageFormatter creates a message formatter
func NewMessageFormatter() *MessageFormatter {
	return &MessageFormatter{
		paramPattern: regexp.MustCompile(`\{([^}]+)\}`),
	}
}

// Format formats a message with parameters
func (f *MessageFormatter) Format(message string, params map[string]any) string {
	return f.paramPattern.ReplaceAllStringFunc(message, func(match string) string {
		key := match[1 : len(match)-1] // Remove { }

		// Check for formatting options (e.g., {count:number})
		parts := strings.Split(key, ":")
		paramKey := parts[0]

		value, ok := params[paramKey]
		if !ok {
			return match
		}

		// Apply formatting if specified
		if len(parts) > 1 {
			format := parts[1]
			return f.applyFormat(value, format)
		}

		return fmt.Sprintf("%v", value)
	})
}

// applyFormat applies formatting to a value
func (f *MessageFormatter) applyFormat(value any, format string) string {
	switch format {
	case "number":
		if num, ok := toNumber(value); ok {
			return formatNumber(num)
		}
	case "currency":
		if num, ok := toNumber(value); ok {
			return formatCurrency(num)
		}
	case "percent":
		if num, ok := toNumber(value); ok {
			return formatPercent(num)
		}
	}

	return fmt.Sprintf("%v", value)
}

// Helper functions

func toNumber(value any) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	default:
		return 0, false
	}
}

func formatNumber(num float64) string {
	return fmt.Sprintf("%.2f", num)
}

func formatCurrency(num float64) string {
	return fmt.Sprintf("$%.2f", num)
}

func formatPercent(num float64) string {
	return fmt.Sprintf("%.1f%%", num*100)
}

// I18nLoader loads translation files from filesystem
type I18nLoader struct {
	basePath string
	cache    map[string]map[string]string
	mu       sync.RWMutex
}

// NewI18nLoader creates a new translation file loader
func NewI18nLoader(basePath string) *I18nLoader {
	return &I18nLoader{
		basePath: basePath,
		cache:    make(map[string]map[string]string),
	}
}

// LoadLocale loads translations for a specific locale
func (l *I18nLoader) LoadLocale(locale string) (map[string]string, error) {
	l.mu.RLock()
	if cached, ok := l.cache[locale]; ok {
		l.mu.RUnlock()
		return cached, nil
	}
	l.mu.RUnlock()

	// Load from file
	filename := filepath.Join(l.basePath, locale+".json")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var translations map[string]interface{}
	if err := json.Unmarshal(data, &translations); err != nil {
		return nil, fmt.Errorf("unmarshal json: %w", err)
	}

	// Flatten the translations
	flattened := flattenJSON(translations)

	// Cache the result
	l.mu.Lock()
	l.cache[locale] = flattened
	l.mu.Unlock()

	return flattened, nil
}

// flattenJSON flattens a nested JSON structure
func flattenJSON(data map[string]interface{}) map[string]string {
	result := make(map[string]string)
	flattenJSONHelper(data, "", result)
	return result
}

// flattenJSONHelper recursively flattens JSON
func flattenJSONHelper(data map[string]interface{}, prefix string, result map[string]string) {
	for key, value := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		switch v := value.(type) {
		case string:
			result[fullKey] = v
		case map[string]interface{}:
			flattenJSONHelper(v, fullKey, result)
		default:
			result[fullKey] = fmt.Sprintf("%v", v)
		}
	}
}
