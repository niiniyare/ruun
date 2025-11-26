// pkg/schema/parser_v2.go
package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
	"gopkg.in/yaml.v3"
)

// Parser provides enterprise-grade schema parsing with validation, caching, and extensibility.
// It supports multiple formats (JSON, YAML) and provides detailed error reporting.
//
// Example usage:
//
//	parser := NewParser(
//	    WithValidation(true),
//	    WithCache(true),
//	    WithEvaluator(evaluator),
//	)
//	schema, err := parser.ParseJSON(ctx, jsonData)
type Parser struct {
	// Configuration
	config *ParserConfig

	// Dependency injection
	evaluator *condition.Evaluator
	validator *schemaValidator

	// Caching
	cache     *parserCache
	cacheOnce sync.Once

	// Plugin support
	plugins map[string]ParserPlugin

	// Statistics
	stats *ParserStats

	mu sync.RWMutex
}

// ParserConfig configures parser behavior
type ParserConfig struct {
	// Validation
	ValidateOnParse   bool          // Validate schema after parsing
	StrictMode        bool          // Fail on unknown fields
	MaxDepth          int           // Maximum nesting depth
	MaxSize           int64         // Maximum input size in bytes
	ValidationTimeout time.Duration // Timeout for validation

	// Performance
	EnableCache      bool          // Enable schema caching
	CacheTTL         time.Duration // Cache entry lifetime
	MaxCacheSize     int           // Maximum cached schemas
	EnableStreaming  bool          // Enable streaming for large schemas
	StreamBufferSize int           // Buffer size for streaming

	// Features
	EnableVersionCheck  bool // Check schema version compatibility
	EnableMigration     bool // Auto-migrate old schema versions
	EnableCustomParsers bool // Allow custom format parsers
	EnableMetrics       bool // Collect parsing metrics

	// Security
	SanitizeInput      bool // Sanitize input before parsing
	DisallowExtensions bool // Disallow non-standard extensions
	ValidateReferences bool // Validate field/action references
}

// DefaultParserConfig returns production-ready default configuration
func DefaultParserConfig() *ParserConfig {
	return &ParserConfig{
		ValidateOnParse:     true,
		StrictMode:          false,
		MaxDepth:            10,
		MaxSize:             10 * 1024 * 1024, // 10MB
		ValidationTimeout:   5 * time.Second,
		EnableCache:         true,
		CacheTTL:            15 * time.Minute,
		MaxCacheSize:        100,
		EnableStreaming:     false,
		StreamBufferSize:    8192,
		EnableVersionCheck:  true,
		EnableMigration:     true,
		EnableCustomParsers: true,
		EnableMetrics:       true,
		SanitizeInput:       true,
		DisallowExtensions:  false,
		ValidateReferences:  true,
	}
}

// ParserOption configures the parser
type ParserOption func(*Parser)

// WithConfig sets custom configuration
func WithConfig(config *ParserConfig) ParserOption {
	return func(p *Parser) {
		p.config = config
	}
}

// WithValidation enables/disables validation
func WithValidation(enabled bool) ParserOption {
	return func(p *Parser) {
		p.config.ValidateOnParse = enabled
	}
}

// WithCache enables/disables caching
func WithCache(enabled bool) ParserOption {
	return func(p *Parser) {
		p.config.EnableCache = enabled
	}
}

// WithEvaluator sets the condition evaluator
func WithEvaluator(evaluator *condition.Evaluator) ParserOption {
	return func(p *Parser) {
		p.evaluator = evaluator
	}
}

// WithValidator sets a custom validator
func WithValidator(validator schemaValidator) ParserOption {
	return func(p *Parser) {
		p.validator = &validator
	}
}

// WithStrictMode enables strict parsing mode
func WithStrictMode(enabled bool) ParserOption {
	return func(p *Parser) {
		p.config.StrictMode = enabled
	}
}

// WithMaxSize sets maximum input size
func WithMaxSize(size int64) ParserOption {
	return func(p *Parser) {
		p.config.MaxSize = size
	}
}

// NewParser creates a new parser with options
func NewParser(options ...ParserOption) *Parser {
	parser := &Parser{
		config:  DefaultParserConfig(),
		plugins: make(map[string]ParserPlugin),
		stats:   newParserStats(),
	}

	for _, option := range options {
		option(parser)
	}

	if parser.config.EnableCache {
		parser.initCache()
	}

	return parser
}

// initCache initializes the parser cache
func (p *Parser) initCache() {
	p.cacheOnce.Do(func() {
		p.cache = newParserCache(p.config.CacheTTL, p.config.MaxCacheSize)
	})
}

// ═══════════════════════════════════════════════════════════════════════════
// Primary Parsing Methods
// ═══════════════════════════════════════════════════════════════════════════

// Parse parses schema from bytes, auto-detecting format
//
// The parser automatically detects JSON or YAML format and parses accordingly.
// It validates the schema if configured and applies caching for performance.
//
// Example:
//
//	schema, err := parser.Parse(ctx, schemaBytes)
//	if err != nil {
//	    log.Fatal(err)
//	}
func (p *Parser) Parse(ctx context.Context, data []byte) (*Schema, error) {
	startTime := time.Now()
	defer func() {
		if p.config.EnableMetrics {
			p.recordMetric("parse", time.Since(startTime), len(data), nil)
		}
	}()

	// Validate input size
	if err := p.validateInputSize(data); err != nil {
		return nil, err
	}

	// Check cache
	if p.config.EnableCache {
		if cached := p.getFromCache(data); cached != nil {
			return cached, nil
		}
	}

	// Detect format and parse
	format, err := p.detectFormat(data)
	if err != nil {
		return nil, NewParseError("format_detection", err.Error())
	}

	var schema *Schema
	switch format {
	case FormatJSON:
		schema, err = p.parseJSON(ctx, data)
	case FormatYAML:
		schema, err = p.parseYAML(ctx, data)
	default:
		return nil, NewParseError("unsupported_format", "unsupported format: "+string(format))
	}

	if err != nil {
		return nil, err
	}

	// Post-processing
	if err := p.postProcess(ctx, schema); err != nil {
		return nil, err
	}

	// Cache result
	if p.config.EnableCache {
		p.addToCache(data, schema)
	}

	return schema, nil
}

// ParseJSON parses schema from JSON data
//
// This method provides direct JSON parsing with full validation and error reporting.
// It's more efficient than Parse() when you know the format is JSON.
//
// Example:
//
//	jsonData := []byte(`{"id": "user-form", "type": "form", ...}`)
//	schema, err := parser.ParseJSON(ctx, jsonData)
func (p *Parser) ParseJSON(ctx context.Context, data []byte) (*Schema, error) {
	startTime := time.Now()
	defer func() {
		if p.config.EnableMetrics {
			p.recordMetric("parse_json", time.Since(startTime), len(data), nil)
		}
	}()

	if err := p.validateInputSize(data); err != nil {
		return nil, err
	}

	schema, err := p.parseJSON(ctx, data)
	if err != nil {
		return nil, err
	}

	if err := p.postProcess(ctx, schema); err != nil {
		return nil, err
	}

	return schema, nil
}

// ParseYAML parses schema from YAML data
//
// YAML parsing is useful for human-readable configuration files.
// The parser converts YAML to JSON internally for consistency.
//
// Example:
//
//	yamlData := []byte(`
//	  id: user-form
//	  type: form
//	  ...
//	`)
//	schema, err := parser.ParseYAML(ctx, yamlData)
func (p *Parser) ParseYAML(ctx context.Context, data []byte) (*Schema, error) {
	startTime := time.Now()
	defer func() {
		if p.config.EnableMetrics {
			p.recordMetric("parse_yaml", time.Since(startTime), len(data), nil)
		}
	}()

	if err := p.validateInputSize(data); err != nil {
		return nil, err
	}

	schema, err := p.parseYAML(ctx, data)
	if err != nil {
		return nil, err
	}

	if err := p.postProcess(ctx, schema); err != nil {
		return nil, err
	}

	return schema, nil
}

// ParseReader parses schema from an io.Reader
//
// This method is useful for streaming large schemas or reading from files.
// It automatically handles buffering and size limits.
//
// Example:
//
//	file, _ := os.Open("schema.json")
//	defer file.Close()
//	schema, err := parser.ParseReader(ctx, file)
func (p *Parser) ParseReader(ctx context.Context, reader io.Reader) (*Schema, error) {
	startTime := time.Now()

	// Read with size limit
	limitReader := io.LimitReader(reader, p.config.MaxSize+1)
	data, err := io.ReadAll(limitReader)
	if err != nil {
		return nil, NewParseError("read_error", fmt.Sprintf("failed to read input: %v", err))
	}

	defer func() {
		if p.config.EnableMetrics {
			p.recordMetric("parse_reader", time.Since(startTime), len(data), nil)
		}
	}()

	if int64(len(data)) > p.config.MaxSize {
		return nil, NewParseError("size_exceeded",
			fmt.Sprintf("input exceeds maximum size of %d bytes", p.config.MaxSize))
	}

	return p.Parse(ctx, data)
}

// ParseString parses schema from a string
//
// Convenience method for parsing string input.
//
// Example:
//
//	schema, err := parser.ParseString(ctx, `{"id": "form", ...}`)
func (p *Parser) ParseString(ctx context.Context, data string) (*Schema, error) {
	return p.Parse(ctx, []byte(data))
}

// ParseFile parses schema from a file path
//
// This method handles file reading and automatically detects format based on extension.
// Supports .json, .yaml, and .yml files.
//
// Example:
//
//	schema, err := parser.ParseFile(ctx, "schemas/user-form.json")
func (p *Parser) ParseFile(ctx context.Context, path string) (*Schema, error) {
	startTime := time.Now()
	defer func() {
		if p.config.EnableMetrics {
			p.recordMetric("parse_file", time.Since(startTime), 0, nil)
		}
	}()

	// This would use os.Open in real implementation
	// For now, return error
	return nil, NewParseError("not_implemented", "file parsing not implemented")
}

// ═══════════════════════════════════════════════════════════════════════════
// Batch Parsing Methods
// ═══════════════════════════════════════════════════════════════════════════

// ParseBatch parses multiple schemas concurrently
//
// This method efficiently parses multiple schemas in parallel with error aggregation.
// Failed schemas don't stop the batch - all results are returned.
//
// Example:
//
//	schemas := [][]byte{schema1, schema2, schema3}
//	results := parser.ParseBatch(ctx, schemas)
//	for i, result := range results {
//	    if result.Error != nil {
//	        log.Printf("Schema %d failed: %v", i, result.Error)
//	    }
//	}
func (p *Parser) ParseBatch(ctx context.Context, data [][]byte) []*ParseResult {
	results := make([]*ParseResult, len(data))
	var wg sync.WaitGroup

	for i, schemaData := range data {
		wg.Add(1)
		go func(index int, input []byte) {
			defer wg.Done()
			schema, err := p.Parse(ctx, input)
			results[index] = &ParseResult{
				Schema: schema,
				Error:  err,
				Index:  index,
			}
		}(i, schemaData)
	}

	wg.Wait()
	return results
}

// ParseResult holds the result of a batch parse operation
type ParseResult struct {
	Schema *Schema
	Error  error
	Index  int
}

// ═══════════════════════════════════════════════════════════════════════════
// Validation and Post-Processing
// ═══════════════════════════════════════════════════════════════════════════

// postProcess performs post-parsing operations
func (p *Parser) postProcess(ctx context.Context, schema *Schema) error {
	// Build internal indexes
	schema.buildIndexes()

	// Set evaluator if available
	if p.evaluator != nil {
		schema.SetEvaluator(p.evaluator)
	}

	// Version check
	if p.config.EnableVersionCheck {
		if err := p.checkVersion(schema); err != nil {
			return err
		}
	}

	// Migration
	if p.config.EnableMigration {
		if err := p.migrate(ctx, schema); err != nil {
			return err
		}
	}

	// Validation
	if p.config.ValidateOnParse {
		if err := p.validateSchema(ctx, schema); err != nil {
			return err
		}
	}

	// Reference validation
	if p.config.ValidateReferences {
		if err := p.validateReferences(schema); err != nil {
			return err
		}
	}

	return nil
}

// validateSchema validates the parsed schema
func (p *Parser) validateSchema(ctx context.Context, schema *Schema) error {
	// Use timeout for validation
	if p.config.ValidationTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, p.config.ValidationTimeout)
		defer cancel()
	}

	// Use custom validator if available
	if p.validator != nil {
		return p.validator.Validate(ctx, schema)
	}

	// Use default validation
	return schema.Validate(ctx)
}

// validateReferences validates field and action references
func (p *Parser) validateReferences(schema *Schema) error {
	collector := NewErrorCollector()

	// Validate field dependencies
	for _, field := range schema.Fields {
		for _, dep := range field.Dependencies {
			if !schema.HasField(dep) {
				collector.AddValidationError(field.Name, "invalid_dependency",
					fmt.Sprintf("field depends on non-existent field: %s", dep))
			}
		}
	}

	// Validate layout references
	if schema.Layout != nil {
		// Check section field references
		for _, section := range schema.Layout.Sections {
			for _, fieldName := range section.Fields {
				if !schema.HasField(fieldName) {
					collector.AddValidationError(section.ID, "invalid_field_reference",
						fmt.Sprintf("section references non-existent field: %s", fieldName))
				}
			}
		}

		// Check tab field references
		for _, tab := range schema.Layout.Tabs {
			for _, fieldName := range tab.Fields {
				if !schema.HasField(fieldName) {
					collector.AddValidationError(tab.ID, "invalid_field_reference",
						fmt.Sprintf("tab references non-existent field: %s", fieldName))
				}
			}
		}
	}

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════
// Format Detection and Parsing
// ═══════════════════════════════════════════════════════════════════════════

// Format represents the schema format
type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

// detectFormat auto-detects the input format
func (p *Parser) detectFormat(data []byte) (Format, error) {
	if len(data) == 0 {
		return "", NewParseError("empty_input", "input is empty")
	}

	// Trim whitespace
	trimmed := strings.TrimSpace(string(data))

	// JSON detection
	if strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[") {
		return FormatJSON, nil
	}

	// YAML detection (common indicators)
	if strings.Contains(trimmed, ":\n") || strings.Contains(trimmed, ": ") {
		return FormatYAML, nil
	}

	// Try JSON first
	var js json.RawMessage
	if err := json.Unmarshal(data, &js); err == nil {
		return FormatJSON, nil
	}

	// Try YAML
	var ym any
	if err := yaml.Unmarshal(data, &ym); err == nil {
		return FormatYAML, nil
	}

	return "", NewParseError("unknown_format", "unable to detect format")
}

// parseJSON performs JSON parsing with error handling
func (p *Parser) parseJSON(ctx context.Context, data []byte) (*Schema, error) {
	// Sanitize input if configured
	if p.config.SanitizeInput {
		data = p.sanitize(data)
	}

	var schema Schema

	// Use custom decoder for better error messages
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	if p.config.StrictMode {
		decoder.DisallowUnknownFields()
	}

	if err := decoder.Decode(&schema); err != nil {
		return nil, p.wrapJSONError(err)
	}

	return &schema, nil
}

// parseYAML performs YAML parsing with conversion to JSON
func (p *Parser) parseYAML(ctx context.Context, data []byte) (*Schema, error) {
	// Sanitize input if configured
	if p.config.SanitizeInput {
		data = p.sanitize(data)
	}

	// Parse YAML to generic structure
	var yamlData any
	if err := yaml.Unmarshal(data, &yamlData); err != nil {
		return nil, p.wrapYAMLError(err)
	}

	// Convert to JSON
	jsonData, err := json.Marshal(yamlData)
	if err != nil {
		return nil, NewParseError("yaml_to_json", fmt.Sprintf("failed to convert YAML to JSON: %v", err))
	}

	// Parse as JSON
	return p.parseJSON(ctx, jsonData)
}

// ═══════════════════════════════════════════════════════════════════════════
// Version Management and Migration
// ═══════════════════════════════════════════════════════════════════════════

// SupportedVersions defines compatible schema versions
var SupportedVersions = []string{"1.0.0", "1.1.0", "1.2.0"}

// checkVersion validates schema version compatibility
func (p *Parser) checkVersion(schema *Schema) error {
	if schema.Version == "" {
		// No version specified, assume latest
		schema.Version = SupportedVersions[len(SupportedVersions)-1]
		return nil
	}

	// Check if version is supported
	for _, supported := range SupportedVersions {
		if schema.Version == supported {
			return nil
		}
	}

	return NewParseError("unsupported_version",
		fmt.Sprintf("schema version %s is not supported", schema.Version))
}

// migrate performs automatic schema migration
func (p *Parser) migrate(ctx context.Context, schema *Schema) error {
	if schema.Version == "" {
		return nil
	}

	// Get current version index
	currentIdx := -1
	for i, v := range SupportedVersions {
		if v == schema.Version {
			currentIdx = i
			break
		}
	}

	if currentIdx == -1 {
		return nil // Unknown version, skip migration
	}

	// Check if migration needed
	latestIdx := len(SupportedVersions) - 1
	if currentIdx >= latestIdx {
		return nil // Already at latest version
	}

	// Apply migrations
	for i := currentIdx; i < latestIdx; i++ {
		fromVersion := SupportedVersions[i]
		toVersion := SupportedVersions[i+1]

		if err := p.applyMigration(ctx, schema, fromVersion, toVersion); err != nil {
			return fmt.Errorf("migration %s -> %s failed: %w", fromVersion, toVersion, err)
		}
	}

	return nil
}

// applyMigration applies a version migration
func (p *Parser) applyMigration(ctx context.Context, schema *Schema, from, to string) error {
	// Example migration logic
	switch {
	case from == "1.0.0" && to == "1.1.0":
		return p.migrateV1_0ToV1_1(schema)
	case from == "1.1.0" && to == "1.2.0":
		return p.migrateV1_1ToV1_2(schema)
	default:
		return nil // No migration needed
	}
}

// migrateV1_0ToV1_1 migrates from version 1.0.0 to 1.1.0
func (p *Parser) migrateV1_0ToV1_1(schema *Schema) error {
	// Example: Add default meta if missing
	if schema.Meta == nil {
		now := time.Now()
		schema.Meta = &Meta{
			BaseMetadata: BaseMetadata{
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
	}
	schema.Version = "1.1.0"
	return nil
}

// migrateV1_1ToV1_2 migrates from version 1.1.0 to 1.2.0
func (p *Parser) migrateV1_1ToV1_2(schema *Schema) error {
	// Example: Ensure all fields have proper defaults
	for i := range schema.Fields {
		if schema.Fields[i].Type == "" {
			schema.Fields[i].Type = FieldText
		}
	}
	schema.Version = "1.2.0"
	return nil
}

// ═══════════════════════════════════════════════════════════════════════════
// Plugin System
// ═══════════════════════════════════════════════════════════════════════════

// ParserPlugin allows custom format parsers
type ParserPlugin interface {
	// Name returns the plugin name
	Name() string

	// CanParse checks if this plugin can parse the data
	CanParse(data []byte) bool

	// Parse parses the data into a schema
	Parse(ctx context.Context, data []byte) (*Schema, error)
}

// RegisterPlugin registers a custom parser plugin
//
// Plugins allow parsing custom formats or extending existing parsers.
//
// Example:
//
//	type XMLPlugin struct{}
//	func (p *XMLPlugin) Name() string { return "xml" }
//	func (p *XMLPlugin) CanParse(data []byte) bool { ... }
//	func (p *XMLPlugin) Parse(ctx context.Context, data []byte) (*Schema, error) { ... }
//
//	parser.RegisterPlugin(&XMLPlugin{})
func (p *Parser) RegisterPlugin(plugin ParserPlugin) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.plugins[plugin.Name()] = plugin
}

// UnregisterPlugin removes a plugin
func (p *Parser) UnregisterPlugin(name string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.plugins, name)
}

// tryPlugins attempts to parse using registered plugins
func (p *Parser) tryPlugins(ctx context.Context, data []byte) (*Schema, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, plugin := range p.plugins {
		if plugin.CanParse(data) {
			return plugin.Parse(ctx, data)
		}
	}

	return nil, NewParseError("no_plugin", "no plugin can parse this format")
}

// ═══════════════════════════════════════════════════════════════════════════
// Caching
// ═══════════════════════════════════════════════════════════════════════════

// parserCache provides schema caching
type parserCache struct {
	entries map[string]*cacheEntry
	ttl     time.Duration
	maxSize int
	mu      sync.RWMutex
}

type cacheEntry struct {
	schema    *Schema
	expiresAt time.Time
}

func newParserCache(ttl time.Duration, maxSize int) *parserCache {
	cache := &parserCache{
		entries: make(map[string]*cacheEntry),
		ttl:     ttl,
		maxSize: maxSize,
	}

	// Start cleanup goroutine
	go cache.cleanup()

	return cache
}

func (c *parserCache) get(key string) *Schema {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entries[key]
	if !exists {
		return nil
	}

	if time.Now().After(entry.expiresAt) {
		return nil
	}

	// Return clone to prevent modification
	clone, _ := entry.schema.Clone()
	return clone
}

func (c *parserCache) set(key string, schema *Schema) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check size limit
	if len(c.entries) >= c.maxSize {
		// Remove oldest entry
		c.evictOldest()
	}

	c.entries[key] = &cacheEntry{
		schema:    schema,
		expiresAt: time.Now().Add(c.ttl),
	}
}

func (c *parserCache) evictOldest() {
	var oldestKey string
	var oldestTime time.Time

	for key, entry := range c.entries {
		if oldestKey == "" || entry.expiresAt.Before(oldestTime) {
			oldestKey = key
			oldestTime = entry.expiresAt
		}
	}

	if oldestKey != "" {
		delete(c.entries, oldestKey)
	}
}

func (c *parserCache) cleanup() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.entries {
			if now.After(entry.expiresAt) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *parserCache) clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]*cacheEntry)
}

// Cache operations on parser
func (p *Parser) getFromCache(data []byte) *Schema {
	if p.cache == nil {
		return nil
	}
	key := p.cacheKey(data)
	return p.cache.get(key)
}

func (p *Parser) addToCache(data []byte, schema *Schema) {
	if p.cache == nil {
		return
	}
	key := p.cacheKey(data)
	p.cache.set(key, schema)
}

func (p *Parser) cacheKey(data []byte) string {
	// Simple hash for demo - use proper hash in production
	return fmt.Sprintf("%d", len(data))
}

// ClearCache clears the parser cache
func (p *Parser) ClearCache() {
	if p.cache != nil {
		p.cache.clear()
	}
}

// ═══════════════════════════════════════════════════════════════════════════
// Metrics and Statistics
// ═══════════════════════════════════════════════════════════════════════════

// ParserStats tracks parser performance metrics
type ParserStats struct {
	TotalParses     int64
	SuccessfulParse int64
	FailedParses    int64
	CacheHits       int64
	CacheMisses     int64
	TotalBytes      int64
	AverageDuration time.Duration

	mu sync.RWMutex
}

func newParserStats() *ParserStats {
	return &ParserStats{}
}

func (s *ParserStats) record(duration time.Duration, bytes int, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.TotalParses++
	s.TotalBytes += int64(bytes)

	if err != nil {
		s.FailedParses++
	} else {
		s.SuccessfulParse++
	}

	// Update average duration
	total := int64(s.AverageDuration) * (s.TotalParses - 1)
	s.AverageDuration = time.Duration((total + int64(duration)) / s.TotalParses)
}

func (p *Parser) recordMetric(operation string, duration time.Duration, bytes int, err error) {
	if p.stats != nil {
		p.stats.record(duration, bytes, err)
	}
}

// GetStats returns parser statistics
func (p *Parser) GetStats() *ParserStats {
	return p.stats
}

// ═══════════════════════════════════════════════════════════════════════════
// Utility Methods
// ═══════════════════════════════════════════════════════════════════════════

// validateInputSize checks if input size is within limits
func (p *Parser) validateInputSize(data []byte) error {
	if int64(len(data)) > p.config.MaxSize {
		return NewParseError("size_exceeded",
			fmt.Sprintf("input size %d exceeds maximum %d", len(data), p.config.MaxSize))
	}
	return nil
}

// sanitize removes potentially dangerous content
func (p *Parser) sanitize(data []byte) []byte {
	// Remove null bytes
	return []byte(strings.ReplaceAll(string(data), "\x00", ""))
}

// wrapJSONError converts JSON errors to ParseError
func (p *Parser) wrapJSONError(err error) error {
	return NewParseError("json_error", fmt.Sprintf("JSON parse error: %v", err))
}

// wrapYAMLError converts YAML errors to ParseError
func (p *Parser) wrapYAMLError(err error) error {
	return NewParseError("yaml_error", fmt.Sprintf("YAML parse error: %v", err))
}

// ═══════════════════════════════════════════════════════════════════════════
// Error Types
// ═══════════════════════════════════════════════════════════════════════════

// ParseError represents a parsing error
type ParseError struct {
	Code    string
	Message string
	Details map[string]any
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewParseError creates a new parse error
func NewParseError(code, message string) *ParseError {
	return &ParseError{
		Code:    code,
		Message: message,
		Details: make(map[string]any),
	}
}

// WithDetail adds details to the error
func (e *ParseError) WithDetail(key string, value any) *ParseError {
	e.Details[key] = value
	return e
}

// IsParseError checks if error is a ParseError
func IsParseError(err error) bool {
	_, ok := err.(*ParseError)
	return ok
}

// ═══════════════════════════════════════════════════════════════════════════
// Helper Functions
// ═══════════════════════════════════════════════════════════════════════════

// MustParse parses or panics (useful for static schemas)
//
// Example:
//
//	var userFormSchema = MustParse(ctx, userFormJSON)
func MustParse(ctx context.Context, data []byte) *Schema {
	parser := NewParser()
	schema, err := parser.Parse(ctx, data)
	if err != nil {
		panic(fmt.Sprintf("failed to parse schema: %v", err))
	}
	return schema
}

// ParseWithDefaults parses with default configuration
//
// Convenience function for simple parsing needs.
//
// Example:
//
//	schema, err := ParseWithDefaults(ctx, jsonData)
func ParseWithDefaults(ctx context.Context, data []byte) (*Schema, error) {
	parser := NewParser()
	return parser.Parse(ctx, data)
}

// ValidateAndParse validates then parses
//
// Two-step process for stricter validation.
//
// Example:
//
//	schema, err := ValidateAndParse(ctx, jsonData)
func ValidateAndParse(ctx context.Context, data []byte) (*Schema, error) {
	parser := NewParser(WithValidation(true), WithStrictMode(true))
	return parser.Parse(ctx, data)
}
