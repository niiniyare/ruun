package schema

import (
	"fmt"
	"time"
)

// StorageBackendType defines the type of storage backend
type StorageBackendType string

const (
	StorageBackendMemory StorageBackendType = "memory"
	StorageBackendFile   StorageBackendType = "file"
	StorageBackendRedis  StorageBackendType = "redis"
	StorageBackendS3     StorageBackendType = "s3"
)

// IsValid checks if the storage backend type is valid
func (sbt StorageBackendType) IsValid() bool {
	switch sbt {
	case StorageBackendMemory, StorageBackendFile, StorageBackendRedis, StorageBackendS3:
		return true
	default:
		return false
	}
}

// UnifiedStorageConfig provides a comprehensive storage configuration system
type UnifiedStorageConfig struct {
	// Backend Configuration
	Type    StorageBackendType `json:"type" validate:"required"`
	Enabled bool               `json:"enabled"`

	// Common Storage Settings
	KeyPrefix  string        `json:"keyPrefix,omitempty"`
	TTL        time.Duration `json:"ttl,omitempty"`
	Timeout    time.Duration `json:"timeout,omitempty"`
	RetryCount int           `json:"retryCount,omitempty"`
	RetryDelay time.Duration `json:"retryDelay,omitempty"`

	// Caching Configuration
	Cache *CacheConfig `json:"cache,omitempty"`

	// Backend-specific settings (stored as map for flexibility)
	BackendSettings map[string]any `json:"backendSettings,omitempty"`

	// Registry Features
	Features *StorageFeatures `json:"features,omitempty"`

	// Connection and Performance
	Connection *ConnectionConfig `json:"connection,omitempty"`
}

// CacheConfig defines caching behavior
type CacheConfig struct {
	EnableMemoryCache      bool          `json:"enableMemoryCache"`
	MemoryCacheTTL         time.Duration `json:"memoryCacheTTL,omitempty"`
	MaxMemoryCacheSize     int           `json:"maxMemoryCacheSize,omitempty"`
	EnableDistributedCache bool          `json:"enableDistributedCache"`
	DistributedCacheTTL    time.Duration `json:"distributedCacheTTL,omitempty"`
	CacheCompression       bool          `json:"cacheCompression,omitempty"`
	CacheEvictionPolicy    string        `json:"cacheEvictionPolicy,omitempty"` // "lru", "lfu", "ttl"
}

// StorageFeatures defines what features are enabled
type StorageFeatures struct {
	ValidateOnStore   bool `json:"validateOnStore"`
	ValidateOnLoad    bool `json:"validateOnLoad"`
	EnableVersioning  bool `json:"enableVersioning"`
	MaxVersions       int  `json:"maxVersions,omitempty"`
	EnableEvents      bool `json:"enableEvents"`
	EnableMetrics     bool `json:"enableMetrics"`
	EnableCompression bool `json:"enableCompression"`
	EnableEncryption  bool `json:"enableEncryption"`
	EnableReplication bool `json:"enableReplication"`
	EnableSharding    bool `json:"enableSharding"`
}

// ConnectionConfig defines connection settings
type ConnectionConfig struct {
	MaxConnections      int           `json:"maxConnections,omitempty"`
	ConnectionTimeout   time.Duration `json:"connectionTimeout,omitempty"`
	IdleTimeout         time.Duration `json:"idleTimeout,omitempty"`
	MaxRetries          int           `json:"maxRetries,omitempty"`
	HealthCheckInterval time.Duration `json:"healthCheckInterval,omitempty"`
	EnablePooling       bool          `json:"enablePooling,omitempty"`
}

// Backend-specific configuration helpers

// SetRedisConfig sets Redis-specific configuration
func (usc *UnifiedStorageConfig) SetRedisConfig(client any, prefix string, ttl time.Duration) {
	usc.Type = StorageBackendRedis
	usc.KeyPrefix = prefix
	usc.TTL = ttl
	if usc.BackendSettings == nil {
		usc.BackendSettings = make(map[string]any)
	}
	usc.BackendSettings["client"] = client
}

// GetRedisConfig extracts Redis configuration
func (usc *UnifiedStorageConfig) GetRedisConfig() (client any, prefix string, ttl time.Duration) {
	if usc.BackendSettings != nil {
		client = usc.BackendSettings["client"]
	}
	prefix = usc.KeyPrefix
	if prefix == "" {
		prefix = "schema:" // Default Redis prefix
	}
	ttl = usc.TTL
	return
}

// SetS3Config sets S3-specific configuration
func (usc *UnifiedStorageConfig) SetS3Config(client any, bucket, keyPrefix string) {
	usc.Type = StorageBackendS3
	usc.KeyPrefix = keyPrefix
	if usc.BackendSettings == nil {
		usc.BackendSettings = make(map[string]any)
	}
	usc.BackendSettings["client"] = client
	usc.BackendSettings["bucket"] = bucket
}

// GetS3Config extracts S3 configuration
func (usc *UnifiedStorageConfig) GetS3Config() (client any, bucket, keyPrefix string) {
	if usc.BackendSettings != nil {
		client = usc.BackendSettings["client"]
		if b, ok := usc.BackendSettings["bucket"].(string); ok {
			bucket = b
		}
	}
	keyPrefix = usc.KeyPrefix
	if keyPrefix == "" {
		keyPrefix = "schemas/" // Default S3 prefix
	}
	return
}

// SetFileConfig sets file-based storage configuration
func (usc *UnifiedStorageConfig) SetFileConfig(basePath string) {
	usc.Type = StorageBackendFile
	if usc.BackendSettings == nil {
		usc.BackendSettings = make(map[string]any)
	}
	usc.BackendSettings["basePath"] = basePath
}

// GetFileConfig extracts file storage configuration
func (usc *UnifiedStorageConfig) GetFileConfig() (basePath string) {
	if usc.BackendSettings != nil {
		if p, ok := usc.BackendSettings["basePath"].(string); ok {
			basePath = p
		}
	}
	if basePath == "" {
		basePath = "./schemas" // Default file path
	}
	return
}

// Validation methods

// Validate checks if the configuration is valid
func (usc *UnifiedStorageConfig) Validate() error {
	if !usc.Type.IsValid() {
		return fmt.Errorf("invalid storage backend type: %s", usc.Type)
	}

	if usc.RetryCount < 0 {
		return fmt.Errorf("retry count cannot be negative")
	}

	if usc.Features != nil && usc.Features.MaxVersions < 0 {
		return fmt.Errorf("max versions cannot be negative")
	}

	// Backend-specific validation
	switch usc.Type {
	case StorageBackendRedis:
		client, _, _ := usc.GetRedisConfig()
		if client == nil {
			return fmt.Errorf("Redis client is required for Redis backend")
		}
	case StorageBackendS3:
		client, bucket, _ := usc.GetS3Config()
		if client == nil {
			return fmt.Errorf("S3 client is required for S3 backend")
		}
		if bucket == "" {
			return fmt.Errorf("S3 bucket is required for S3 backend")
		}
	case StorageBackendFile:
		basePath := usc.GetFileConfig()
		if basePath == "" {
			return fmt.Errorf("base path is required for file backend")
		}
	}

	return nil
}

// Helper methods

// IsDistributedBackend returns true if this is a distributed storage backend
func (usc *UnifiedStorageConfig) IsDistributedBackend() bool {
	return usc.Type == StorageBackendRedis || usc.Type == StorageBackendS3
}

// SupportsCaching returns true if the backend supports caching
func (usc *UnifiedStorageConfig) SupportsCaching() bool {
	return true // All backends can support caching
}

// SupportsVersioning returns true if versioning is supported and enabled
func (usc *UnifiedStorageConfig) SupportsVersioning() bool {
	return usc.Features != nil && usc.Features.EnableVersioning
}

// GetEffectiveTTL returns the effective TTL with fallbacks
func (usc *UnifiedStorageConfig) GetEffectiveTTL() time.Duration {
	if usc.TTL > 0 {
		return usc.TTL
	}
	// Default TTL based on backend type
	switch usc.Type {
	case StorageBackendMemory:
		return 1 * time.Hour
	case StorageBackendRedis:
		return 24 * time.Hour
	case StorageBackendS3, StorageBackendFile:
		return 0 // No TTL for persistent storage by default
	default:
		return 1 * time.Hour
	}
}

// Clone creates a copy of the configuration
func (usc *UnifiedStorageConfig) Clone() *UnifiedStorageConfig {
	if usc == nil {
		return nil
	}

	clone := *usc

	// Deep copy maps
	if usc.BackendSettings != nil {
		clone.BackendSettings = make(map[string]any)
		for k, v := range usc.BackendSettings {
			clone.BackendSettings[k] = v
		}
	}

	// Deep copy nested structs
	if usc.Cache != nil {
		cacheCopy := *usc.Cache
		clone.Cache = &cacheCopy
	}

	if usc.Features != nil {
		featuresCopy := *usc.Features
		clone.Features = &featuresCopy
	}

	if usc.Connection != nil {
		connectionCopy := *usc.Connection
		clone.Connection = &connectionCopy
	}

	return &clone
}

// Builder for UnifiedStorageConfig

// StorageConfigBuilder provides a fluent API for building storage configurations
type StorageConfigBuilder struct {
	config *UnifiedStorageConfig
	*BuilderMixin
}

// NewStorageConfig creates a new storage configuration builder
func NewStorageConfig(backendType StorageBackendType) *StorageConfigBuilder {
	return &StorageConfigBuilder{
		config: &UnifiedStorageConfig{
			Type:    backendType,
			Enabled: true,
		},
		BuilderMixin: NewBuilderMixin(),
	}
}

// WithPrefix sets the key prefix
func (b *StorageConfigBuilder) WithPrefix(prefix string) *StorageConfigBuilder {
	b.config.KeyPrefix = prefix
	return b
}

// WithTTL sets the TTL
func (b *StorageConfigBuilder) WithTTL(ttl time.Duration) *StorageConfigBuilder {
	if ttl < 0 {
		b.Context.AddError(fmt.Errorf("TTL cannot be negative"))
		return b
	}
	b.config.TTL = ttl
	return b
}

// WithRetries sets retry configuration
func (b *StorageConfigBuilder) WithRetries(count int, delay time.Duration) *StorageConfigBuilder {
	if count < 0 {
		b.Context.AddError(fmt.Errorf("retry count cannot be negative"))
		return b
	}
	b.config.RetryCount = count
	b.config.RetryDelay = delay
	return b
}

// WithCaching enables caching with the specified configuration
func (b *StorageConfigBuilder) WithCaching(memoryTTL, distributedTTL time.Duration) *StorageConfigBuilder {
	if b.config.Cache == nil {
		b.config.Cache = &CacheConfig{}
	}
	b.config.Cache.EnableMemoryCache = true
	b.config.Cache.MemoryCacheTTL = memoryTTL
	b.config.Cache.EnableDistributedCache = true
	b.config.Cache.DistributedCacheTTL = distributedTTL
	return b
}

// WithFeatures sets storage features
func (b *StorageConfigBuilder) WithFeatures(features *StorageFeatures) *StorageConfigBuilder {
	if features != nil && features.MaxVersions < 0 {
		b.Context.AddError(fmt.Errorf("max versions cannot be negative"))
		return b
	}
	b.config.Features = features
	return b
}

// EnableVersioning enables versioning with max versions
func (b *StorageConfigBuilder) EnableVersioning(maxVersions int) *StorageConfigBuilder {
	if maxVersions < 0 {
		b.Context.AddError(fmt.Errorf("max versions cannot be negative"))
		return b
	}
	if b.config.Features == nil {
		b.config.Features = &StorageFeatures{}
	}
	b.config.Features.EnableVersioning = true
	b.config.Features.MaxVersions = maxVersions
	return b
}

// EnableMetrics enables metrics collection
func (b *StorageConfigBuilder) EnableMetrics() *StorageConfigBuilder {
	if b.config.Features == nil {
		b.config.Features = &StorageFeatures{}
	}
	b.config.Features.EnableMetrics = true
	return b
}

// Build implements BaseBuilder interface
func (b *StorageConfigBuilder) Build() (*UnifiedStorageConfig, error) {
	if err := b.CheckBuild("StorageConfig"); err != nil {
		return nil, err
	}

	// Validate the configuration
	if err := b.config.Validate(); err != nil {
		return nil, fmt.Errorf("storage config validation failed: %w", err)
	}

	return b.config.Clone(), nil
}

// MustBuild builds and panics on error
func (b *StorageConfigBuilder) MustBuild() *UnifiedStorageConfig {
	result, err := b.Build()
	if err != nil {
		panic(fmt.Sprintf("StorageConfigBuilder.MustBuild() failed: %v", err))
	}
	return result
}

// Implement remaining BaseBuilder interface methods
func (b *StorageConfigBuilder) HasErrors() bool {
	return b.Context.HasErrors()
}

func (b *StorageConfigBuilder) GetErrors() []error {
	return b.Context.GetErrors()
}

func (b *StorageConfigBuilder) AddError(err error) BaseBuilder[UnifiedStorageConfig] {
	b.Context.AddError(err)
	return b
}

func (b *StorageConfigBuilder) ClearErrors() BaseBuilder[UnifiedStorageConfig] {
	b.Context.ClearErrors()
	return b
}

// Convenience constructors

// NewRedisStorageConfig creates Redis storage configuration
func NewRedisStorageConfig(client any) *StorageConfigBuilder {
	builder := NewStorageConfig(StorageBackendRedis)
	builder.config.SetRedisConfig(client, "schema:", 24*time.Hour)
	return builder
}

// NewS3StorageConfig creates S3 storage configuration
func NewS3StorageConfig(client any, bucket string) *StorageConfigBuilder {
	builder := NewStorageConfig(StorageBackendS3)
	builder.config.SetS3Config(client, bucket, "schemas/")
	return builder
}

// NewFileStorageConfig creates file storage configuration
func NewFileStorageConfig(basePath string) *StorageConfigBuilder {
	builder := NewStorageConfig(StorageBackendFile)
	builder.config.SetFileConfig(basePath)
	return builder
}

// NewMemoryStorageConfig creates memory storage configuration
func NewMemoryStorageConfig() *StorageConfigBuilder {
	return NewStorageConfig(StorageBackendMemory).
		WithTTL(1*time.Hour).
		WithCaching(5*time.Minute, 15*time.Minute)
}

// Legacy RegistryConfig compatibility methods
func (usc *UnifiedStorageConfig) GetEnableStorage() bool {
	return usc.Enabled
}

func (usc *UnifiedStorageConfig) SetEnableStorage(enabled bool) {
	usc.Enabled = enabled
}

func (usc *UnifiedStorageConfig) GetStorageType() string {
	return string(usc.Type)
}

func (usc *UnifiedStorageConfig) SetStorageType(storageType string) {
	usc.Type = StorageBackendType(storageType)
}

func (usc *UnifiedStorageConfig) GetEnableMemoryCache() bool {
	return usc.Cache != nil && usc.Cache.EnableMemoryCache
}

func (usc *UnifiedStorageConfig) GetMemoryCacheTTL() time.Duration {
	if usc.Cache != nil {
		return usc.Cache.MemoryCacheTTL
	}
	return 0
}

func (usc *UnifiedStorageConfig) GetMaxMemoryCacheSize() int {
	if usc.Cache != nil {
		return usc.Cache.MaxMemoryCacheSize
	}
	return 0
}

func (usc *UnifiedStorageConfig) GetEnableDistributedCache() bool {
	return usc.Cache != nil && usc.Cache.EnableDistributedCache
}

func (usc *UnifiedStorageConfig) GetDistributedCacheTTL() time.Duration {
	if usc.Cache != nil {
		return usc.Cache.DistributedCacheTTL
	}
	return 0
}

func (usc *UnifiedStorageConfig) GetValidateOnStore() bool {
	return usc.Features != nil && usc.Features.ValidateOnStore
}

func (usc *UnifiedStorageConfig) GetValidateOnLoad() bool {
	return usc.Features != nil && usc.Features.ValidateOnLoad
}

func (usc *UnifiedStorageConfig) GetEnableVersioning() bool {
	return usc.Features != nil && usc.Features.EnableVersioning
}

func (usc *UnifiedStorageConfig) GetMaxVersions() int {
	if usc.Features != nil {
		return usc.Features.MaxVersions
	}
	return 0
}

func (usc *UnifiedStorageConfig) GetEnableEvents() bool {
	return usc.Features != nil && usc.Features.EnableEvents
}

func (usc *UnifiedStorageConfig) GetEnableMetrics() bool {
	return usc.Features != nil && usc.Features.EnableMetrics
}

// Legacy field access for RegistryConfig compatibility
// These fields are exposed as properties to maintain API compatibility

// EnableStorage field getter/setter
func (usc *UnifiedStorageConfig) EnableStorage() bool {
	return usc.GetEnableStorage()
}

// StorageType field getter
func (usc *UnifiedStorageConfig) StorageType() string {
	return usc.GetStorageType()
}

// And so on for other fields...
// Note: In practice, we'd need actual struct fields for backward compatibility
// Let me add a proper compatibility layer

// LegacyRegistryConfig provides field-level compatibility for old RegistryConfig usage
type LegacyRegistryConfig struct {
	EnableStorage          bool
	StorageType            string
	EnableMemoryCache      bool
	MemoryCacheTTL         time.Duration
	MaxMemoryCacheSize     int
	EnableDistributedCache bool
	DistributedCacheTTL    time.Duration
	ValidateOnStore        bool
	ValidateOnLoad         bool
	EnableVersioning       bool
	MaxVersions            int
	EnableEvents           bool
	EnableMetrics          bool
}

// ToUnifiedConfig converts legacy config to unified config
func (lrc *LegacyRegistryConfig) ToUnifiedConfig() *UnifiedStorageConfig {
	config := &UnifiedStorageConfig{
		Type:    StorageBackendType(lrc.StorageType),
		Enabled: lrc.EnableStorage,
	}

	if lrc.EnableMemoryCache || lrc.EnableDistributedCache {
		config.Cache = &CacheConfig{
			EnableMemoryCache:      lrc.EnableMemoryCache,
			MemoryCacheTTL:         lrc.MemoryCacheTTL,
			MaxMemoryCacheSize:     lrc.MaxMemoryCacheSize,
			EnableDistributedCache: lrc.EnableDistributedCache,
			DistributedCacheTTL:    lrc.DistributedCacheTTL,
		}
	}

	if lrc.ValidateOnStore || lrc.ValidateOnLoad || lrc.EnableVersioning ||
		lrc.EnableEvents || lrc.EnableMetrics {
		config.Features = &StorageFeatures{
			ValidateOnStore:  lrc.ValidateOnStore,
			ValidateOnLoad:   lrc.ValidateOnLoad,
			EnableVersioning: lrc.EnableVersioning,
			MaxVersions:      lrc.MaxVersions,
			EnableEvents:     lrc.EnableEvents,
			EnableMetrics:    lrc.EnableMetrics,
		}
	}

	return config
}

// FromUnifiedConfig creates legacy config from unified config
func (usc *UnifiedStorageConfig) ToLegacyConfig() *LegacyRegistryConfig {
	legacy := &LegacyRegistryConfig{
		EnableStorage: usc.Enabled,
		StorageType:   string(usc.Type),
	}

	if usc.Cache != nil {
		legacy.EnableMemoryCache = usc.Cache.EnableMemoryCache
		legacy.MemoryCacheTTL = usc.Cache.MemoryCacheTTL
		legacy.MaxMemoryCacheSize = usc.Cache.MaxMemoryCacheSize
		legacy.EnableDistributedCache = usc.Cache.EnableDistributedCache
		legacy.DistributedCacheTTL = usc.Cache.DistributedCacheTTL
	}

	if usc.Features != nil {
		legacy.ValidateOnStore = usc.Features.ValidateOnStore
		legacy.ValidateOnLoad = usc.Features.ValidateOnLoad
		legacy.EnableVersioning = usc.Features.EnableVersioning
		legacy.MaxVersions = usc.Features.MaxVersions
		legacy.EnableEvents = usc.Features.EnableEvents
		legacy.EnableMetrics = usc.Features.EnableMetrics
	}

	return legacy
}

// Backward compatibility type aliases
type RedisConfig = UnifiedStorageConfig
type S3Config = UnifiedStorageConfig
type RegistryConfig = LegacyRegistryConfig // Use legacy struct for now
