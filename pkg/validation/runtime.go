package validation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

// RuntimeValidator provides runtime validation for user inputs and API calls
type RuntimeValidator struct {
	middleware   []ValidationMiddleware
	interceptors []ValidationInterceptor
	config       RuntimeValidationConfig
	metrics      *RuntimeMetrics
	cache        *ValidationCache
	rateLimiter  *RateLimiter
}

// RuntimeValidationConfig configures runtime validation
type RuntimeValidationConfig struct {
	EnableInputValidation bool          `json:"enableInputValidation"`
	EnableAPIValidation   bool          `json:"enableAPIValidation"`
	EnableSanitization    bool          `json:"enableSanitization"`
	EnableRateLimit       bool          `json:"enableRateLimit"`
	EnableCaching         bool          `json:"enableCaching"`
	ValidateAsync         bool          `json:"validateAsync"`
	CacheTTL              time.Duration `json:"cacheTTL"`
	MaxCacheSize          int           `json:"maxCacheSize"`
	RateLimitRequests     int           `json:"rateLimitRequests"`
	RateLimitWindow       time.Duration `json:"rateLimitWindow"`
	ValidationTimeout     time.Duration `json:"validationTimeout"`
	StrictMode            bool          `json:"strictMode"`
	AllowedOrigins        []string      `json:"allowedOrigins"`
	CSRFProtection        bool          `json:"csrfProtection"`
}

// ValidationMiddleware interface for HTTP middleware
type ValidationMiddleware interface {
	ValidateRequest(ctx context.Context, req *http.Request) *ValidationResult
	ValidateResponse(ctx context.Context, resp *http.Response) *ValidationResult
	GetPriority() int
}

// ValidationInterceptor interface for intercepting validation calls
type ValidationInterceptor interface {
	BeforeValidation(ctx context.Context, data any) (any, error)
	AfterValidation(ctx context.Context, data any, result *ValidationResult) *ValidationResult
	OnValidationError(ctx context.Context, data any, err error) error
}

// RuntimeMetrics contains runtime validation metrics
type RuntimeMetrics struct {
	TotalValidations      int64            `json:"totalValidations"`
	SuccessfulValidations int64            `json:"successfulValidations"`
	FailedValidations     int64            `json:"failedValidations"`
	AverageLatency        time.Duration    `json:"averageLatency"`
	ValidationsByType     map[string]int64 `json:"validationsByType"`
	ErrorsByType          map[string]int64 `json:"errorsByType"`
	CacheHitRate          float64          `json:"cacheHitRate"`
	RateLimitHits         int64            `json:"rateLimitHits"`
	LastUpdated           time.Time        `json:"lastUpdated"`
	mutex                 sync.RWMutex
}

// ValidationCache provides caching for validation results
type ValidationCache struct {
	cache   map[string]*CacheEntry
	mutex   sync.RWMutex
	maxSize int
	ttl     time.Duration
	hits    int64
	misses  int64
}

// CacheEntry represents a cached validation result
type CacheEntry struct {
	Result    *ValidationResult `json:"result"`
	Timestamp time.Time         `json:"timestamp"`
	TTL       time.Duration     `json:"ttl"`
}

// RateLimiter provides rate limiting for validation requests
type RateLimiter struct {
	requests map[string]*RequestCounter
	mutex    sync.RWMutex
	limit    int
	window   time.Duration
}

// RequestCounter tracks requests per client
type RequestCounter struct {
	Count        int       `json:"count"`
	Window       time.Time `json:"window"`
	Blocked      bool      `json:"blocked"`
	BlockedUntil time.Time `json:"blockedUntil"`
}

// RuntimeValidationRequest represents a validation request
type RuntimeValidationRequest struct {
	ID        string         `json:"id"`
	Type      string         `json:"type"` // input, api, form, etc.
	Data      any            `json:"data"`
	Schema    any            `json:"schema,omitempty"`
	Context   map[string]any `json:"context"`
	ClientID  string         `json:"clientId,omitempty"`
	Async     bool           `json:"async"`
	Timeout   time.Duration  `json:"timeout,omitempty"`
	Priority  int            `json:"priority"`
	Timestamp time.Time      `json:"timestamp"`
}

// RuntimeValidationResponse represents a validation response
type RuntimeValidationResponse struct {
	ID        string            `json:"id"`
	Valid     bool              `json:"valid"`
	Result    *ValidationResult `json:"result"`
	Sanitized any               `json:"sanitized,omitempty"`
	Cached    bool              `json:"cached"`
	Duration  time.Duration     `json:"duration"`
	Timestamp time.Time         `json:"timestamp"`
}

// InputValidationMiddleware validates user inputs
type InputValidationMiddleware struct {
	sanitizer *InputSanitizer
	validator *InputValidator
	config    InputValidationConfig
}

// APIValidationMiddleware validates API requests and responses
type APIValidationMiddleware struct {
	schemas   map[string]any
	validator *APIValidator
	config    APIValidationConfig
}

// CSRFValidationMiddleware provides CSRF protection
type CSRFValidationMiddleware struct {
	tokenGenerator *CSRFTokenGenerator
	config         CSRFValidationConfig
}

// InputValidationConfig configures input validation
type InputValidationConfig struct {
	SanitizeHTML     bool     `json:"sanitizeHTML"`
	SanitizeSQL      bool     `json:"sanitizeSQL"`
	SanitizeXSS      bool     `json:"sanitizeXSS"`
	MaxLength        int      `json:"maxLength"`
	AllowedTags      []string `json:"allowedTags"`
	BlockedPatterns  []string `json:"blockedPatterns"`
	RequiredFields   []string `json:"requiredFields"`
	ValidateTypes    bool     `json:"validateTypes"`
	StrictValidation bool     `json:"strictValidation"`
}

// APIValidationConfig configures API validation
type APIValidationConfig struct {
	ValidateRequest     bool     `json:"validateRequest"`
	ValidateResponse    bool     `json:"validateResponse"`
	ValidateHeaders     bool     `json:"validateHeaders"`
	RequiredHeaders     []string `json:"requiredHeaders"`
	MaxRequestSize      int64    `json:"maxRequestSize"`
	AllowedContentTypes []string `json:"allowedContentTypes"`
	ValidateAuth        bool     `json:"validateAuth"`
	LogRequests         bool     `json:"logRequests"`
}

// CSRFValidationConfig configures CSRF protection
type CSRFValidationConfig struct {
	Enabled        bool          `json:"enabled"`
	TokenLength    int           `json:"tokenLength"`
	TokenTTL       time.Duration `json:"tokenTTL"`
	HeaderName     string        `json:"headerName"`
	FormFieldName  string        `json:"formFieldName"`
	CookieName     string        `json:"cookieName"`
	SecureCookie   bool          `json:"secureCookie"`
	SameSite       string        `json:"sameSite"`
	ValidateOrigin bool          `json:"validateOrigin"`
}

// NewRuntimeValidator creates a new runtime validator
func NewRuntimeValidator() *RuntimeValidator {
	config := RuntimeValidationConfig{
		EnableInputValidation: true,
		EnableAPIValidation:   true,
		EnableSanitization:    true,
		EnableRateLimit:       true,
		EnableCaching:         true,
		ValidateAsync:         false,
		CacheTTL:              time.Minute * 5,
		MaxCacheSize:          1000,
		RateLimitRequests:     100,
		RateLimitWindow:       time.Minute,
		ValidationTimeout:     time.Second * 5,
		StrictMode:            false,
		CSRFProtection:        true,
	}

	return &RuntimeValidator{
		middleware:   make([]ValidationMiddleware, 0),
		interceptors: make([]ValidationInterceptor, 0),
		config:       config,
		metrics:      NewRuntimeMetrics(),
		cache:        NewValidationCache(config.MaxCacheSize, config.CacheTTL),
		rateLimiter:  NewRateLimiter(config.RateLimitRequests, config.RateLimitWindow),
	}
}

// NewRuntimeMetrics creates new runtime metrics
func NewRuntimeMetrics() *RuntimeMetrics {
	return &RuntimeMetrics{
		ValidationsByType: make(map[string]int64),
		ErrorsByType:      make(map[string]int64),
		LastUpdated:       time.Now(),
	}
}

// NewValidationCache creates a new validation cache
func NewValidationCache(maxSize int, ttl time.Duration) *ValidationCache {
	return &ValidationCache{
		cache:   make(map[string]*CacheEntry),
		maxSize: maxSize,
		ttl:     ttl,
	}
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]*RequestCounter),
		limit:    limit,
		window:   window,
	}
}

// ValidateInput validates user input data
func (rv *RuntimeValidator) ValidateInput(ctx context.Context, data any, schema any) *RuntimeValidationResponse {
	request := &RuntimeValidationRequest{
		ID:        generateRequestID(),
		Type:      "input",
		Data:      data,
		Schema:    schema,
		Context:   make(map[string]any),
		Async:     rv.config.ValidateAsync,
		Timeout:   rv.config.ValidationTimeout,
		Priority:  1,
		Timestamp: time.Now(),
	}

	return rv.processValidationRequest(ctx, request)
}

// ValidateAPI validates API requests and responses
func (rv *RuntimeValidator) ValidateAPI(ctx context.Context, req *http.Request) *RuntimeValidationResponse {
	request := &RuntimeValidationRequest{
		ID:        generateRequestID(),
		Type:      "api",
		Data:      req,
		Context:   map[string]any{"method": req.Method, "url": req.URL.String()},
		ClientID:  rv.getClientID(req),
		Async:     false, // API validation should be synchronous
		Timeout:   rv.config.ValidationTimeout,
		Priority:  2,
		Timestamp: time.Now(),
	}

	return rv.processValidationRequest(ctx, request)
}

// ValidateForm validates form submissions
func (rv *RuntimeValidator) ValidateForm(ctx context.Context, formData map[string]any, formSchema any) *RuntimeValidationResponse {
	request := &RuntimeValidationRequest{
		ID:        generateRequestID(),
		Type:      "form",
		Data:      formData,
		Schema:    formSchema,
		Context:   make(map[string]any),
		Async:     rv.config.ValidateAsync,
		Timeout:   rv.config.ValidationTimeout,
		Priority:  1,
		Timestamp: time.Now(),
	}

	return rv.processValidationRequest(ctx, request)
}

// processValidationRequest processes a validation request
func (rv *RuntimeValidator) processValidationRequest(ctx context.Context, request *RuntimeValidationRequest) *RuntimeValidationResponse {
	start := time.Now()

	response := &RuntimeValidationResponse{
		ID:        request.ID,
		Timestamp: start,
	}

	// Check rate limiting
	if rv.config.EnableRateLimit {
		if !rv.rateLimiter.Allow(request.ClientID) {
			response.Valid = false
			response.Result = &ValidationResult{
				Valid: false,
				Level: ValidationLevelError,
				Errors: []ValidationError{
					NewValidationError("rate_limit_exceeded", "Too many validation requests", "", ValidationLevelError),
				},
				Timestamp: time.Now(),
				Metadata:  map[string]any{"rateLimited": true},
			}
			rv.metrics.RecordRateLimitHit()
			return response
		}
	}

	// Check cache
	if rv.config.EnableCaching {
		cacheKey := rv.generateCacheKey(request)
		if cached := rv.cache.Get(cacheKey); cached != nil {
			response.Valid = cached.Valid
			response.Result = cached
			response.Cached = true
			response.Duration = time.Since(start)
			rv.metrics.RecordValidation(request.Type, true, response.Duration)
			return response
		}
	}

	// Run interceptors - before validation
	processedData := request.Data
	for _, interceptor := range rv.interceptors {
		var err error
		processedData, err = interceptor.BeforeValidation(ctx, processedData)
		if err != nil {
			response.Valid = false
			response.Result = &ValidationResult{
				Valid: false,
				Errors: []ValidationError{
					NewValidationError("interceptor_error", err.Error(), "", ValidationLevelError),
				},
				Timestamp: time.Now(),
			}
			response.Duration = time.Since(start)
			return response
		}
	}

	// Perform validation based on type
	var result *ValidationResult
	var sanitized any

	switch request.Type {
	case "input":
		result, sanitized = rv.validateInputData(ctx, processedData, request.Schema)
	case "api":
		result = rv.validateAPIRequest(ctx, processedData.(*http.Request))
	case "form":
		result, sanitized = rv.validateFormData(ctx, processedData.(map[string]any), request.Schema)
	default:
		result = &ValidationResult{
			Valid: false,
			Errors: []ValidationError{
				NewValidationError("unknown_validation_type", fmt.Sprintf("Unknown validation type: %s", request.Type), "", ValidationLevelError),
			},
			Timestamp: time.Now(),
		}
	}

	// Run interceptors - after validation
	for _, interceptor := range rv.interceptors {
		result = interceptor.AfterValidation(ctx, processedData, result)
	}

	// Cache the result if enabled
	if rv.config.EnableCaching && result.Valid {
		cacheKey := rv.generateCacheKey(request)
		rv.cache.Set(cacheKey, result)
	}

	response.Valid = result.Valid
	response.Result = result
	response.Sanitized = sanitized
	response.Duration = time.Since(start)

	// Record metrics
	rv.metrics.RecordValidation(request.Type, result.Valid, response.Duration)

	return response
}

// validateInputData validates input data with sanitization
func (rv *RuntimeValidator) validateInputData(ctx context.Context, data any, schema any) (*ValidationResult, any) {
	result := &ValidationResult{
		Valid:     true,
		Level:     ValidationLevelError,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Sanitize input if enabled
	var sanitized any = data
	if rv.config.EnableSanitization {
		sanitized = rv.sanitizeInput(data)
	}

	// Validate against schema if provided
	if schema != nil {
		engine := NewValidationEngine(ValidationOptions{
			StrictMode: rv.config.StrictMode,
		})

		validationCtx := NewValidationContext(ctx, ValidationLevelError, "runtime_input")
		schemaResult := engine.ValidateSchema(validationCtx, schema)

		if !schemaResult.IsValid() {
			result.Valid = false
			result.Errors = append(result.Errors, schemaResult.Errors...)
		}
	}

	// Additional input validation rules
	if dataMap, ok := sanitized.(map[string]any); ok {
		for field, value := range dataMap {
			if err := rv.validateFieldInput(field, value); err != nil {
				result.Valid = false
				result.Errors = append(result.Errors, *err)
			}
		}
	}

	return result, sanitized
}

// validateAPIRequest validates API requests
func (rv *RuntimeValidator) validateAPIRequest(ctx context.Context, req *http.Request) *ValidationResult {
	result := &ValidationResult{
		Valid:     true,
		Level:     ValidationLevelError,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Validate request size
	if req.ContentLength > 0 && req.ContentLength > 10*1024*1024 { // 10MB limit
		result.Valid = false
		result.Errors = append(result.Errors, NewValidationError(
			"request_too_large",
			"Request size exceeds maximum allowed size",
			"content-length",
			ValidationLevelError,
		))
	}

	// Validate content type
	contentType := req.Header.Get("Content-Type")
	if contentType != "" && !rv.isAllowedContentType(contentType) {
		result.Valid = false
		result.Errors = append(result.Errors, NewValidationError(
			"invalid_content_type",
			fmt.Sprintf("Content type %s is not allowed", contentType),
			"content-type",
			ValidationLevelError,
		))
	}

	// Validate required headers
	requiredHeaders := []string{"User-Agent"}
	for _, header := range requiredHeaders {
		if req.Header.Get(header) == "" {
			result.Valid = false
			result.Errors = append(result.Errors, NewValidationError(
				"missing_required_header",
				fmt.Sprintf("Required header %s is missing", header),
				header,
				ValidationLevelError,
			))
		}
	}

	// Validate origin if CSRF protection is enabled
	if rv.config.CSRFProtection {
		if !rv.validateOrigin(req) {
			result.Valid = false
			result.Errors = append(result.Errors, NewValidationError(
				"invalid_origin",
				"Request origin is not allowed",
				"origin",
				ValidationLevelError,
			))
		}
	}

	return result
}

// validateFormData validates form submission data
func (rv *RuntimeValidator) validateFormData(ctx context.Context, data map[string]any, schema any) (*ValidationResult, any) {
	result := &ValidationResult{
		Valid:     true,
		Level:     ValidationLevelError,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	// Sanitize form data
	sanitized := rv.sanitizeFormData(data)

	// Validate required fields
	requiredFields := []string{"csrf_token"} // Example required field
	for _, field := range requiredFields {
		if _, exists := sanitized[field]; !exists {
			result.Valid = false
			result.Errors = append(result.Errors, NewValidationError(
				"missing_required_field",
				fmt.Sprintf("Required field %s is missing", field),
				field,
				ValidationLevelError,
			))
		}
	}

	// Validate field formats and constraints
	for field, value := range sanitized {
		if err := rv.validateFormField(field, value); err != nil {
			result.Valid = false
			result.Errors = append(result.Errors, *err)
		}
	}

	return result, sanitized
}

// Input sanitization methods

func (rv *RuntimeValidator) sanitizeInput(data any) any {
	if dataMap, ok := data.(map[string]any); ok {
		sanitized := make(map[string]any)
		for key, value := range dataMap {
			sanitized[key] = rv.sanitizeValue(value)
		}
		return sanitized
	} else if dataSlice, ok := data.([]any); ok {
		sanitized := make([]any, len(dataSlice))
		for i, value := range dataSlice {
			sanitized[i] = rv.sanitizeValue(value)
		}
		return sanitized
	}

	return rv.sanitizeValue(data)
}

func (rv *RuntimeValidator) sanitizeValue(value any) any {
	if str, ok := value.(string); ok {
		// Basic XSS prevention
		str = rv.removeHTMLTags(str)
		str = rv.escapeHTMLEntities(str)
		str = rv.removeSQLPatterns(str)
		return str
	}

	return value
}

func (rv *RuntimeValidator) removeHTMLTags(input string) string {
	// Simple HTML tag removal - in production, use a proper HTML sanitizer
	return input // Simplified for example
}

func (rv *RuntimeValidator) escapeHTMLEntities(input string) string {
	// Escape HTML entities
	replacements := map[string]string{
		"<":  "&lt;",
		">":  "&gt;",
		"&":  "&amp;",
		"'":  "&#39;",
		"\"": "&quot;",
	}

	result := input
	for old, new := range replacements {
		result = strings.ReplaceAll(result, old, new)
	}

	return result
}

func (rv *RuntimeValidator) removeSQLPatterns(input string) string {
	// Remove potential SQL injection patterns
	sqlPatterns := []string{
		"'", "--", "/*", "*/", "xp_", "sp_",
		"UNION", "SELECT", "INSERT", "UPDATE", "DELETE",
		"DROP", "CREATE", "ALTER", "EXEC",
	}

	result := input
	for _, pattern := range sqlPatterns {
		result = strings.ReplaceAll(result, pattern, "")
		result = strings.ReplaceAll(result, strings.ToLower(pattern), "")
	}

	return result
}

func (rv *RuntimeValidator) sanitizeFormData(data map[string]any) map[string]any {
	sanitized := make(map[string]any)

	for key, value := range data {
		// Skip sensitive fields from logging
		if rv.isSensitiveField(key) {
			sanitized[key] = "[REDACTED]"
			continue
		}

		sanitized[key] = rv.sanitizeValue(value)
	}

	return sanitized
}

func (rv *RuntimeValidator) isSensitiveField(field string) bool {
	sensitiveFields := []string{"password", "token", "secret", "key", "auth"}
	fieldLower := strings.ToLower(field)

	for _, sensitive := range sensitiveFields {
		if strings.Contains(fieldLower, sensitive) {
			return true
		}
	}

	return false
}

// Field validation methods

func (rv *RuntimeValidator) validateFieldInput(field string, value any) *ValidationError {
	// Implement field-specific validation logic
	if field == "email" {
		if str, ok := value.(string); ok {
			if !rv.isValidEmail(str) {
				return &ValidationError{
					Code:    "invalid_email",
					Message: "Invalid email format",
					Field:   field,
					Level:   ValidationLevelError,
				}
			}
		}
	}

	if field == "phone" {
		if str, ok := value.(string); ok {
			if !rv.isValidPhone(str) {
				return &ValidationError{
					Code:    "invalid_phone",
					Message: "Invalid phone number format",
					Field:   field,
					Level:   ValidationLevelError,
				}
			}
		}
	}

	return nil
}

func (rv *RuntimeValidator) validateFormField(field string, value any) *ValidationError {
	// Additional form-specific validation
	if field == "csrf_token" {
		if str, ok := value.(string); ok {
			if !rv.isValidCSRFToken(str) {
				return &ValidationError{
					Code:    "invalid_csrf_token",
					Message: "Invalid CSRF token",
					Field:   field,
					Level:   ValidationLevelError,
				}
			}
		}
	}

	return nil
}

// Helper validation methods

func (rv *RuntimeValidator) isValidEmail(email string) bool {
	// Basic email validation - use a proper email validation library in production
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

func (rv *RuntimeValidator) isValidPhone(phone string) bool {
	// Basic phone validation
	return len(phone) >= 10 && len(phone) <= 15
}

func (rv *RuntimeValidator) isValidCSRFToken(token string) bool {
	// CSRF token validation logic
	return len(token) >= 32
}

func (rv *RuntimeValidator) isAllowedContentType(contentType string) bool {
	allowedTypes := []string{
		"application/json",
		"application/x-www-form-urlencoded",
		"multipart/form-data",
		"text/plain",
	}

	for _, allowed := range allowedTypes {
		if strings.Contains(contentType, allowed) {
			return true
		}
	}

	return false
}

func (rv *RuntimeValidator) validateOrigin(req *http.Request) bool {
	origin := req.Header.Get("Origin")
	referer := req.Header.Get("Referer")

	if len(rv.config.AllowedOrigins) == 0 {
		return true // No origin restriction
	}

	for _, allowed := range rv.config.AllowedOrigins {
		if origin == allowed || strings.HasPrefix(referer, allowed) {
			return true
		}
	}

	return false
}

// Cache methods

func (vc *ValidationCache) Get(key string) *ValidationResult {
	vc.mutex.RLock()
	defer vc.mutex.RUnlock()

	entry, exists := vc.cache[key]
	if !exists {
		vc.misses++
		return nil
	}

	// Check if entry has expired
	if time.Since(entry.Timestamp) > entry.TTL {
		delete(vc.cache, key)
		vc.misses++
		return nil
	}

	vc.hits++
	return entry.Result
}

func (vc *ValidationCache) Set(key string, result *ValidationResult) {
	vc.mutex.Lock()
	defer vc.mutex.Unlock()

	// Remove oldest entries if cache is full
	if len(vc.cache) >= vc.maxSize {
		vc.evictOldest()
	}

	vc.cache[key] = &CacheEntry{
		Result:    result,
		Timestamp: time.Now(),
		TTL:       vc.ttl,
	}
}

func (vc *ValidationCache) evictOldest() {
	oldestKey := ""
	oldestTime := time.Now()

	for key, entry := range vc.cache {
		if entry.Timestamp.Before(oldestTime) {
			oldestTime = entry.Timestamp
			oldestKey = key
		}
	}

	if oldestKey != "" {
		delete(vc.cache, oldestKey)
	}
}

func (vc *ValidationCache) GetHitRate() float64 {
	total := vc.hits + vc.misses
	if total == 0 {
		return 0
	}
	return float64(vc.hits) / float64(total)
}

// Rate limiter methods

func (rl *RateLimiter) Allow(clientID string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	now := time.Now()
	counter, exists := rl.requests[clientID]

	if !exists {
		rl.requests[clientID] = &RequestCounter{
			Count:  1,
			Window: now,
		}
		return true
	}

	// Check if client is currently blocked
	if counter.Blocked && now.Before(counter.BlockedUntil) {
		return false
	}

	// Reset window if expired
	if now.Sub(counter.Window) > rl.window {
		counter.Count = 1
		counter.Window = now
		counter.Blocked = false
		return true
	}

	// Check if limit is exceeded
	if counter.Count >= rl.limit {
		counter.Blocked = true
		counter.BlockedUntil = now.Add(rl.window)
		return false
	}

	counter.Count++
	return true
}

// Metrics methods

func (rm *RuntimeMetrics) RecordValidation(validationType string, success bool, duration time.Duration) {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	rm.TotalValidations++
	rm.ValidationsByType[validationType]++

	if success {
		rm.SuccessfulValidations++
	} else {
		rm.FailedValidations++
		rm.ErrorsByType[validationType]++
	}

	// Update average latency (simple moving average)
	if rm.TotalValidations == 1 {
		rm.AverageLatency = duration
	} else {
		rm.AverageLatency = time.Duration(
			(int64(rm.AverageLatency)*(rm.TotalValidations-1) + int64(duration)) / rm.TotalValidations,
		)
	}

	rm.LastUpdated = time.Now()
}

func (rm *RuntimeMetrics) RecordRateLimitHit() {
	rm.mutex.Lock()
	defer rm.mutex.Unlock()

	rm.RateLimitHits++
}

func (rm *RuntimeMetrics) GetStats() map[string]any {
	rm.mutex.RLock()
	defer rm.mutex.RUnlock()

	successRate := 0.0
	if rm.TotalValidations > 0 {
		successRate = float64(rm.SuccessfulValidations) / float64(rm.TotalValidations)
	}

	return map[string]any{
		"totalValidations":      rm.TotalValidations,
		"successfulValidations": rm.SuccessfulValidations,
		"failedValidations":     rm.FailedValidations,
		"successRate":           successRate,
		"averageLatency":        rm.AverageLatency.Milliseconds(),
		"validationsByType":     rm.ValidationsByType,
		"errorsByType":          rm.ErrorsByType,
		"rateLimitHits":         rm.RateLimitHits,
		"lastUpdated":           rm.LastUpdated,
	}
}

// Utility methods

func (rv *RuntimeValidator) generateCacheKey(request *RuntimeValidationRequest) string {
	data, _ := json.Marshal(map[string]any{
		"type":   request.Type,
		"data":   request.Data,
		"schema": request.Schema,
	})
	return fmt.Sprintf("validation_%x", data)
}

func (rv *RuntimeValidator) getClientID(req *http.Request) string {
	// Try to get client ID from various sources
	if clientID := req.Header.Get("X-Client-ID"); clientID != "" {
		return clientID
	}

	if userAgent := req.Header.Get("User-Agent"); userAgent != "" {
		return fmt.Sprintf("ua_%x", userAgent)
	}

	return req.RemoteAddr
}

func generateRequestID() string {
	return fmt.Sprintf("req_%d", time.Now().UnixNano())
}

// Middleware registration

func (rv *RuntimeValidator) AddMiddleware(middleware ValidationMiddleware) {
	rv.middleware = append(rv.middleware, middleware)
}

func (rv *RuntimeValidator) AddInterceptor(interceptor ValidationInterceptor) {
	rv.interceptors = append(rv.interceptors, interceptor)
}

// HTTP middleware wrapper

func (rv *RuntimeValidator) HTTPMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Validate the request
			response := rv.ValidateAPI(r.Context(), r)

			if !response.Valid {
				// Return validation error
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]any{
					"error":  "Validation failed",
					"result": response.Result,
				})
				return
			}

			// Continue to next handler
			next.ServeHTTP(w, r)
		})
	}
}

// Specialized validators and sanitizers would be implemented here...

type InputSanitizer struct{}
type InputValidator struct{}
type APIValidator struct{}
type CSRFTokenGenerator struct{}

// Additional implementation details would go here...
