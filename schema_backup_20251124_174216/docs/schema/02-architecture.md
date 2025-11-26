# Architecture Overview

## Three-Layer Architecture

```
┌─────────────────────────────────────┐
│  Layer 1: Schema Definition         │
│  JSON Schema → Go Struct            │
└────────────┬────────────────────────┘
             ↓
┌─────────────────────────────────────┐
│  Layer 2: Processing                │
│  Parse → Validate → Enrich → Cache  │
└────────────┬────────────────────────┘
             ↓
┌─────────────────────────────────────┐
│  Layer 3: Rendering                 │
│  Schema → templ → HTML              │
└─────────────────────────────────────┘
```

## System Components

```
┌────────────────────────────────────────────────┐
│              Browser (Frontend)                 │
│  ┌──────────┐  ┌──────────┐  ┌──────────────┐ │
│  │  HTML5   │  │  HTMX    │  │  Alpine.js   │ │
│  └──────────┘  └──────────┘  └──────────────┘ │
└────────────────────┬───────────────────────────┘
                     │ HTTPS (HTML over wire)
┌────────────────────▼───────────────────────────┐
│            Go Server (Fiber v2)                 │
│                                                 │
│  ┌──────────────────────────────────────────┐  │
│  │  Fiber Handlers + Middleware             │  │
│  │  ├─ Auth                                 │  │
│  │  ├─ Tenant Isolation                     │  │
│  │  └─ Rate Limiting                        │  │
│  └────────────┬─────────────────────────────┘  │
│               │                                 │
│  ┌────────────▼─────────────────────────────┐  │
│  │  Schema Processing                       │  │
│  │  ├─ Registry (storage + cache)          │  │
│  │  ├─ Parser (JSON → Go)                   │  │
│  │  ├─ Validator (rules)                    │  │
│  │  └─ Enricher (permissions)               │  │
│  └────────────┬─────────────────────────────┘  │
│               │                                 │
│  ┌────────────▼─────────────────────────────┐  │
│  │  Rendering (templ)                       │  │
│  │  ├─ Form components                      │  │
│  │  ├─ Field components                     │  │
│  │  └─ Layout components                    │  │
│  └────────────┬─────────────────────────────┘  │
│               │                                 │
│               ▼ HTML Output                     │
└─────────────────────────────────────────────────┘
                     │
┌────────────────────▼───────────────────────────┐
│              Storage Layer                      │
│  ┌──────────┐  ┌──────────┐  ┌──────────────┐ │
│  │PostgreSQL│  │  Redis   │  │ S3/Filesystem│ │
│  └──────────┘  └──────────┘  └──────────────┘ │
└─────────────────────────────────────────────────┘
```

## Component Responsibilities

### Frontend

**HTML5**
- Native form validation (instant feedback)
- Semantic markup
- Accessibility (ARIA)

**HTMX**
- Form submission without page reload
- Partial page updates
- HTTP requests triggered by events

**Alpine.js**
- Client-side state
- Show/hide logic
- Reactive data binding

### Backend

**Fiber Handler**
```go
func HandleForm(c *fiber.Ctx) error {
    // 1. Get user from context
    user := GetUserFromContext(c)
    
    // 2. Load schema
    schema, err := registry.Get(c.Context(), "user-form")
    if err != nil {
        return c.Status(404).SendString("Schema not found")
    }
    
    // 3. Enrich with permissions
    enriched, err := enricher.Enrich(c.Context(), schema, user)
    
    // 4. Render
    return views.FormPage(enriched, nil).Render(c.Context(), c.Response().BodyWriter())
}
```

**Registry**
- Loads schemas from storage (PostgreSQL, Redis, files, S3)
- Caches parsed schemas
- Validates schema structure
- Version management

**Parser**
- Converts JSON to Go structs
- Validates JSON structure
- Handles errors
- Sets defaults

**Validator**
- HTML5 validation rules
- Server-side business rules
- Cross-field validation
- Integration with condition package

**Enricher**
- Adds runtime permissions
- Populates default values
- Applies tenant customization
- Injects user context

**templ Renderer**
- Type-safe template rendering
- Maps field types to components
- Generates HTML
- Injects HTMX/Alpine.js attributes

## Data Flow

### Display Form (GET)

```
1. Browser → GET /forms/user-registration
             ↓
2. Fiber Handler
   ├─ Auth middleware (verify user)
   ├─ Tenant middleware (isolate tenant)
   └─ Handler function
             ↓
3. Registry.Get("user-registration")
   ├─ Check Redis cache
   ├─ If miss: load from PostgreSQL
   ├─ Parse JSON to Schema struct
   └─ Return *Schema
             ↓
4. Enricher.Enrich(schema, user)
   ├─ Get user permissions
   ├─ Apply tenant customization
   ├─ Set field.Runtime.Visible
   └─ Return enriched Schema
             ↓
5. Renderer (templ)
   ├─ Loop through fields
   ├─ Render components based on type
   ├─ Add HTMX attributes
   └─ Generate HTML
             ↓
6. Response → HTML (50KB, ~200ms)
             ↓
7. Browser displays form immediately
```

### Submit Form (POST)

```
1. Browser → Fill form → Click submit
             ↓
2. HTML5 validation (instant client-side)
             ↓
3. HTMX → POST /api/users
             ↓
4. Fiber Handler
   ├─ Auth middleware
   ├─ Parse request body
   └─ Handler function
             ↓
5. Registry.Get("user-registration")
   (get schema for validation)
             ↓
6. Validator.ValidateData(schema, data)
   ├─ Check required fields
   ├─ Check field types
   ├─ Check constraints (min, max, pattern)
   ├─ Check uniqueness (database query)
   └─ Run business rules (condition package)
             ↓
7. If valid:
   ├─ Save to PostgreSQL
   ├─ Return success HTML fragment
   └─ HTMX swaps into #result div
   
   If invalid:
   ├─ Return error HTML fragment
   └─ HTMX displays errors
```

## Package Structure

```
github.com/niiniyare/ruun/
│
├── pkg/
│   ├── condition/          # Business rules engine
│   └── schema/             # Schema system
│       ├── schema.go       # Core Schema struct
│       ├── field.go        # Field struct
│       ├── types.go        # Enums (SchemaType, FieldType)
│       ├── validation.go   # Validation structures
│       ├── runtime.go      # Runtime structs
│       │
│       ├── parse/          # Parser
│       │   └── parser.go
│       │
│       ├── validate/       # Validator
│       │   └── validator.go
│       │
│       ├── enrich/         # Enricher
│       │   └── enricher.go
│       │
│       └── registry/       # Registry
│           ├── registry.go     # Interface
│           ├── postgres.go     # PostgreSQL impl
│           ├── redis.go        # Redis impl
│           ├── filesystem.go   # File impl
│           └── s3.go          # S3/Minio impl
│
├── views/                  # templ components
│   ├── form.templ
│   └── fields/
│       ├── text.templ
│       ├── select.templ
│       └── ...
│
└── internal/
    └── handlers/           # Fiber handlers
        ├── form.go
        └── submit.go
```

## Storage Interface

The registry uses an interface to support multiple storage backends:

```go
type Storage interface {
    // Get retrieves a schema by ID
    Get(ctx context.Context, id string) ([]byte, error)
    
    // Set stores a schema
    Set(ctx context.Context, id string, data []byte) error
    
    // Delete removes a schema
    Delete(ctx context.Context, id string) error
    
    // List returns all schema IDs
    List(ctx context.Context) ([]string, error)
    
    // Exists checks if schema exists
    Exists(ctx context.Context, id string) (bool, error)
}
```

**Implementations:**
- `PostgresStorage` - ACID compliant, primary storage
- `RedisStorage` - Fast cache, optional TTL
- `FilesystemStorage` - Simple, good for development
- `S3Storage` - Cloud object storage (S3, Minio)

See [09-storage-interface.md](09-storage-interface.md) for details.

## Middleware Stack

```go
app := fiber.New()

// Global middleware
app.Use(logger.New())
app.Use(recover.New())

// Schema routes
schemas := app.Group("/schemas")
schemas.Use(authMiddleware)      // Verify JWT
schemas.Use(tenantMiddleware)    // Extract tenant_id
schemas.Use(rateLimitMiddleware) // Rate limiting

// Handlers
schemas.Get("/:id/new", handleFormNew)
schemas.Post("/:id", handleFormSubmit)
schemas.Get("/:id/:recordId", handleFormEdit)
schemas.Put("/:id/:recordId", handleFormUpdate)
```

## Request Context

Middleware adds data to Fiber context:

```go
type RequestContext struct {
    UserID      string
    TenantID    string
    Permissions []string
    Roles       []string
}

// In middleware
func tenantMiddleware(c *fiber.Ctx) error {
    tenantID := extractTenantID(c)
    c.Locals("tenant_id", tenantID)
    return c.Next()
}

// In handler
func handleForm(c *fiber.Ctx) error {
    tenantID := c.Locals("tenant_id").(string)
    // Use tenantID for isolation
}
```

## Caching Strategy

```
Request → Registry.Get(id)
          ├─ Check memory cache (Go map)
          │  └─ If found: return (1ms)
          │
          ├─ Check Redis cache
          │  └─ If found: parse + return (5ms)
          │
          └─ Load from PostgreSQL
             ├─ Parse JSON
             ├─ Cache in Redis (TTL: 1 hour)
             ├─ Cache in memory (TTL: 5 min)
             └─ Return (50ms)
```

## Security Layers

1. **Transport**: HTTPS only
2. **Authentication**: JWT tokens (middleware)
3. **Authorization**: Permission checks (enricher)
4. **Tenant Isolation**: Row-level security (PostgreSQL)
5. **Validation**: Client (HTML5) + Server (Go)
6. **Rate Limiting**: Per user/IP (middleware)
7. **CSRF**: Token validation (optional)

## Advanced Implementation

### Multi-Layered Security Architecture

The schema system implements enterprise-grade security with multiple layers:

```go
// Security configuration in schema
type Security struct {
    // CSRF protection
    CSRF *CSRFConfig `json:"csrf,omitempty"`
    
    // Rate limiting
    RateLimit *RateLimitConfig `json:"rateLimit,omitempty"`
    
    // Field-level encryption
    Encryption *EncryptionConfig `json:"encryption,omitempty"`
    
    // Audit logging
    Audit *AuditConfig `json:"audit,omitempty"`
    
    // Access control
    AccessControl *AccessControlConfig `json:"accessControl,omitempty"`
    
    // Data masking
    DataMasking *DataMaskingConfig `json:"dataMasking,omitempty"`
}

// CSRF protection configuration
type CSRFConfig struct {
    Enabled      bool   `json:"enabled"`
    TokenField   string `json:"tokenField,omitempty"`   // Default: "_csrf_token"
    HeaderName   string `json:"headerName,omitempty"`   // Default: "X-CSRF-Token"
    CookieName   string `json:"cookieName,omitempty"`   // Default: "csrf_token"
    CookieSecure bool   `json:"cookieSecure,omitempty"` // HTTPS only
    CookieSameSite string `json:"cookieSameSite,omitempty"` // Strict, Lax, None
}

// Rate limiting configuration
type RateLimitConfig struct {
    Enabled     bool   `json:"enabled"`
    RequestsPerMinute int `json:"requestsPerMinute"`  // Default: 60
    BurstSize   int    `json:"burstSize"`           // Default: 10
    KeyFunc     string `json:"keyFunc"`             // "ip", "user", "tenant"
    SkipPaths   []string `json:"skipPaths,omitempty"`
}
```

### Field-Level Encryption

```go
// EncryptionConfig defines field encryption settings
type EncryptionConfig struct {
    Enabled    bool             `json:"enabled"`
    Provider   string           `json:"provider"`    // "aes256", "rsa", "vault"
    KeyID      string           `json:"keyId"`       // Key identifier
    Fields     []EncryptedField `json:"fields"`      // Which fields to encrypt
    AtRest     bool             `json:"atRest"`      // Encrypt in database
    InTransit  bool             `json:"inTransit"`   // Encrypt in API
}

type EncryptedField struct {
    Name       string `json:"name"`                    // Field name
    Algorithm  string `json:"algorithm"`               // "aes256-gcm", "rsa-oaep"
    Searchable bool   `json:"searchable,omitempty"`    // Enable searchable encryption
    KeyRotation bool  `json:"keyRotation,omitempty"`   // Auto key rotation
}

// Field encryption implementation
func (e *Encryptor) EncryptField(ctx context.Context, field *Field, value any) (any, error) {
    if !e.shouldEncrypt(field) {
        return value, nil
    }
    
    stringValue := fmt.Sprintf("%v", value)
    
    switch field.Encryption.Algorithm {
    case "aes256-gcm":
        return e.encryptAES256GCM(stringValue, field.Encryption.KeyID)
    case "rsa-oaep":
        return e.encryptRSAOAEP(stringValue, field.Encryption.KeyID)
    default:
        return nil, fmt.Errorf("unsupported encryption algorithm: %s", field.Encryption.Algorithm)
    }
}

func (e *Encryptor) encryptAES256GCM(plaintext, keyID string) (string, error) {
    key, err := e.keyManager.GetKey(keyID)
    if err != nil {
        return "", fmt.Errorf("failed to get encryption key: %w", err)
    }
    
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    
    nonce := make([]byte, aesGCM.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    
    ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}
```

### Comprehensive Audit Logging

```go
// AuditConfig defines audit logging settings
type AuditConfig struct {
    Enabled     bool     `json:"enabled"`
    Events      []string `json:"events"`          // "create", "read", "update", "delete"
    Fields      []string `json:"fields"`          // Which fields to audit
    Retention   int      `json:"retention"`       // Days to retain logs
    Destination string   `json:"destination"`     // "database", "file", "syslog", "elasticsearch"
    Format      string   `json:"format"`          // "json", "structured", "text"
    PII         bool     `json:"pii"`            // Include PII in logs
}

// AuditLogger handles security event logging
type AuditLogger struct {
    config     *AuditConfig
    writer     AuditWriter
    processor  *AuditProcessor
    encryptor  *AuditEncryptor
    mu         sync.RWMutex
}

type AuditEvent struct {
    ID          string                 `json:"id"`
    Timestamp   time.Time              `json:"timestamp"`
    EventType   string                 `json:"event_type"`    // "schema_access", "field_update", etc.
    UserID      string                 `json:"user_id"`
    TenantID    string                 `json:"tenant_id"`
    SchemaID    string                 `json:"schema_id"`
    Resource    string                 `json:"resource"`      // Field name, schema ID, etc.
    Action      string                 `json:"action"`        // "create", "read", "update", "delete"
    Before      map[string]any         `json:"before,omitempty"`
    After       map[string]any         `json:"after,omitempty"`
    IPAddress   string                 `json:"ip_address"`
    UserAgent   string                 `json:"user_agent"`
    SessionID   string                 `json:"session_id"`
    Success     bool                   `json:"success"`
    Error       string                 `json:"error,omitempty"`
    Metadata    map[string]any         `json:"metadata,omitempty"`
    Checksum    string                 `json:"checksum"`      // Tamper detection
}

func (al *AuditLogger) LogSchemaAccess(ctx context.Context, user *User, schemaID string, action string) error {
    event := &AuditEvent{
        ID:        generateEventID(),
        Timestamp: time.Now().UTC(),
        EventType: "schema_access",
        UserID:    user.ID,
        TenantID:  user.TenantID,
        SchemaID:  schemaID,
        Action:    action,
        IPAddress: extractIPFromContext(ctx),
        UserAgent: extractUserAgentFromContext(ctx),
        SessionID: extractSessionIDFromContext(ctx),
        Success:   true,
    }
    
    // Add checksum for tamper detection
    event.Checksum = al.generateChecksum(event)
    
    return al.writeEvent(event)
}

func (al *AuditLogger) LogFieldChange(
    ctx context.Context,
    user *User,
    schemaID, fieldName string,
    oldValue, newValue any,
) error {
    event := &AuditEvent{
        ID:        generateEventID(),
        Timestamp: time.Now().UTC(),
        EventType: "field_change",
        UserID:    user.ID,
        TenantID:  user.TenantID,
        SchemaID:  schemaID,
        Resource:  fieldName,
        Action:    "update",
        Before:    map[string]any{fieldName: oldValue},
        After:     map[string]any{fieldName: newValue},
        IPAddress: extractIPFromContext(ctx),
        UserAgent: extractUserAgentFromContext(ctx),
        Success:   true,
    }
    
    // Encrypt sensitive data if configured
    if al.config.PII && al.encryptor != nil {
        event.Before = al.encryptor.EncryptSensitiveData(event.Before)
        event.After = al.encryptor.EncryptSensitiveData(event.After)
    }
    
    event.Checksum = al.generateChecksum(event)
    return al.writeEvent(event)
}
```

### Advanced Access Control

```go
// AccessControlConfig defines fine-grained permissions
type AccessControlConfig struct {
    Enabled     bool                    `json:"enabled"`
    Mode        string                  `json:"mode"`         // "rbac", "abac", "hybrid"
    Policies    []AccessPolicy          `json:"policies"`
    DefaultDeny bool                    `json:"defaultDeny"`  // Fail closed
    Cache       *AccessControlCache     `json:"cache,omitempty"`
}

type AccessPolicy struct {
    ID          string            `json:"id"`
    Name        string            `json:"name"`
    Description string            `json:"description"`
    Effect      string            `json:"effect"`      // "allow", "deny"
    Subjects    []PolicySubject   `json:"subjects"`    // Who
    Resources   []PolicyResource  `json:"resources"`   // What
    Actions     []string          `json:"actions"`     // Which operations
    Conditions  []PolicyCondition `json:"conditions"`  // When
    Priority    int               `json:"priority"`    // Order of evaluation
}

type PolicySubject struct {
    Type  string `json:"type"`   // "user", "role", "group", "tenant"
    ID    string `json:"id"`     // Subject identifier
    Match string `json:"match"`  // "exact", "pattern", "contains"
}

type PolicyResource struct {
    Type   string `json:"type"`    // "schema", "field", "action"
    ID     string `json:"id"`      // Resource identifier  
    Match  string `json:"match"`   // "exact", "pattern", "contains"
    Tenant string `json:"tenant"`  // Tenant-specific resource
}

type PolicyCondition struct {
    Type     string `json:"type"`      // "time", "ip", "attribute"
    Operator string `json:"operator"`  // "eq", "ne", "in", "between"
    Value    any    `json:"value"`     // Condition value
}

// AccessController enforces access policies
type AccessController struct {
    policies    []AccessPolicy
    evaluator   *PolicyEvaluator
    cache       *AccessCache
    auditor     *AuditLogger
    mu          sync.RWMutex
}

func (ac *AccessController) CheckAccess(
    ctx context.Context,
    subject *PolicySubject,
    resource *PolicyResource,
    action string,
) (bool, error) {
    // Check cache first
    cacheKey := ac.buildCacheKey(subject, resource, action)
    if decision, found := ac.cache.Get(cacheKey); found {
        return decision, nil
    }
    
    // Evaluate policies
    decision := ac.evaluatePolicies(ctx, subject, resource, action)
    
    // Cache result
    ac.cache.Set(cacheKey, decision, 5*time.Minute)
    
    // Audit access decision
    ac.auditor.LogAccessDecision(ctx, subject, resource, action, decision)
    
    return decision, nil
}

func (ac *AccessController) evaluatePolicies(
    ctx context.Context,
    subject *PolicySubject,
    resource *PolicyResource,
    action string,
) bool {
    // Sort policies by priority (higher first)
    sort.Slice(ac.policies, func(i, j int) bool {
        return ac.policies[i].Priority > ac.policies[j].Priority
    })
    
    for _, policy := range ac.policies {
        if ac.policyMatches(ctx, policy, subject, resource, action) {
            if policy.Effect == "deny" {
                return false // Explicit deny
            }
            if policy.Effect == "allow" {
                return true // Explicit allow
            }
        }
    }
    
    // Default behavior based on configuration
    return !ac.config.DefaultDeny
}
```

### Data Masking & Privacy Protection

```go
// DataMaskingConfig defines data privacy settings
type DataMaskingConfig struct {
    Enabled   bool              `json:"enabled"`
    Rules     []MaskingRule     `json:"rules"`
    DefaultMask string          `json:"defaultMask"`   // "****"
    PreservePlaintext bool      `json:"preservePlaintext"` // Keep original for authorized users
}

type MaskingRule struct {
    FieldPattern string   `json:"fieldPattern"`  // Regex pattern for field names
    MaskType     string   `json:"maskType"`      // "partial", "full", "hash", "custom"
    MaskChar     string   `json:"maskChar"`      // Character to use for masking
    PreserveStart int     `json:"preserveStart"` // Characters to preserve at start
    PreserveEnd   int     `json:"preserveEnd"`   // Characters to preserve at end
    Roles        []string `json:"roles"`         // Roles that can see unmasked data
    Custom       string   `json:"custom"`        // Custom masking function
}

// DataMasker handles sensitive data masking
type DataMasker struct {
    config *DataMaskingConfig
    hasher *DataHasher
}

func (dm *DataMasker) MaskField(ctx context.Context, field *Field, value any, user *User) any {
    if !dm.config.Enabled {
        return value
    }
    
    stringValue := fmt.Sprintf("%v", value)
    
    // Find applicable masking rule
    rule := dm.findMaskingRule(field.Name)
    if rule == nil {
        return value // No masking rule
    }
    
    // Check if user has permission to see unmasked data
    if dm.userCanSeeUnmasked(user, rule) {
        return value
    }
    
    // Apply masking
    switch rule.MaskType {
    case "partial":
        return dm.maskPartial(stringValue, rule)
    case "full":
        return strings.Repeat(rule.MaskChar, len(stringValue))
    case "hash":
        return dm.hasher.Hash(stringValue)
    case "custom":
        return dm.applyCustomMask(stringValue, rule.Custom)
    default:
        return dm.config.DefaultMask
    }
}

func (dm *DataMasker) maskPartial(value string, rule *MaskingRule) string {
    if len(value) <= rule.PreserveStart+rule.PreserveEnd {
        return value // Too short to mask
    }
    
    start := value[:rule.PreserveStart]
    end := value[len(value)-rule.PreserveEnd:]
    middle := strings.Repeat(rule.MaskChar, len(value)-rule.PreserveStart-rule.PreserveEnd)
    
    return start + middle + end
}

// Example: Email masking "john.doe@example.com" → "jo****@example.com"
func (dm *DataMasker) maskEmail(email string) string {
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return dm.config.DefaultMask
    }
    
    username := parts[0]
    domain := parts[1]
    
    if len(username) <= 2 {
        return username + "@" + domain
    }
    
    maskedUsername := username[:2] + strings.Repeat("*", len(username)-2)
    return maskedUsername + "@" + domain
}
```

### Secure Session Management

```go
// SessionSecurity handles secure session management
type SessionSecurity struct {
    store        SessionStore
    encryptor    *SessionEncryptor
    validator    *SessionValidator
    config       *SessionConfig
}

type SessionConfig struct {
    Timeout      time.Duration `json:"timeout"`       // Session timeout
    IdleTimeout  time.Duration `json:"idleTimeout"`   // Idle timeout
    MaxSessions  int           `json:"maxSessions"`   // Max concurrent sessions per user
    RotateToken  bool          `json:"rotateToken"`   // Rotate tokens on activity
    IPValidation bool          `json:"ipValidation"`  // Validate IP address
    SecureCookie bool          `json:"secureCookie"`  // HTTPS only cookies
    SameSite     string        `json:"sameSite"`      // Cookie SameSite policy
}

func (ss *SessionSecurity) CreateSession(ctx context.Context, user *User) (*Session, error) {
    // Check concurrent session limit
    if err := ss.checkSessionLimit(user.ID); err != nil {
        return nil, err
    }
    
    session := &Session{
        ID:        generateSecureSessionID(),
        UserID:    user.ID,
        TenantID:  user.TenantID,
        CreatedAt: time.Now().UTC(),
        UpdatedAt: time.Now().UTC(),
        ExpiresAt: time.Now().Add(ss.config.Timeout),
        IPAddress: extractIPFromContext(ctx),
        UserAgent: extractUserAgentFromContext(ctx),
        Data:      make(map[string]any),
    }
    
    // Encrypt session data
    if err := ss.encryptor.EncryptSession(session); err != nil {
        return nil, fmt.Errorf("failed to encrypt session: %w", err)
    }
    
    // Store session
    if err := ss.store.Create(session); err != nil {
        return nil, fmt.Errorf("failed to store session: %w", err)
    }
    
    return session, nil
}

func (ss *SessionSecurity) ValidateSession(ctx context.Context, sessionID string) (*Session, error) {
    session, err := ss.store.Get(sessionID)
    if err != nil {
        return nil, err
    }
    
    // Check expiration
    if time.Now().After(session.ExpiresAt) {
        ss.store.Delete(sessionID)
        return nil, fmt.Errorf("session expired")
    }
    
    // Validate IP address if configured
    if ss.config.IPValidation {
        currentIP := extractIPFromContext(ctx)
        if session.IPAddress != currentIP {
            return nil, fmt.Errorf("IP address mismatch")
        }
    }
    
    // Check idle timeout
    if time.Since(session.UpdatedAt) > ss.config.IdleTimeout {
        ss.store.Delete(sessionID)
        return nil, fmt.Errorf("session idle timeout")
    }
    
    // Update last activity
    session.UpdatedAt = time.Now().UTC()
    
    // Rotate token if configured
    if ss.config.RotateToken {
        session.Token = generateNewToken()
    }
    
    ss.store.Update(session)
    return session, nil
}
```

This advanced security implementation provides enterprise-grade protection including field-level encryption, comprehensive audit logging, fine-grained access control, data masking, and secure session management.

## Key Design Decisions

### 1. Why Fiber?
- **Fast**: One of the fastest Go frameworks
- **Express-like**: Familiar API for many developers
- **Middleware**: Rich ecosystem
- **Low memory**: Efficient for high-traffic

### 2. Why templ?
- **Type-safe**: Compile-time checks
- **Fast**: Generates Go code
- **Simple**: Just Go + HTML
- **No runtime**: Zero overhead

### 3. Why Storage Interface?
- **Flexibility**: Choose backend per environment
- **Testing**: Easy to mock
- **Migration**: Switch storage without code changes
- **Scalability**: Start with files, move to PostgreSQL + Redis

### 4. Why Server-Side Rendering?
- **Performance**: Instant first paint
- **SEO**: Content in HTML
- **Simplicity**: No complex build tools
- **Resilience**: Works without JavaScript

## Performance Characteristics

**With caching:**
- Schema load: 1-5ms (memory/Redis)
- Enrichment: 2-10ms (permission checks)
- Rendering: 5-20ms (templ)
- **Total: 10-35ms**

**Without caching:**
- Schema load: 50-200ms (PostgreSQL + parse)
- Enrichment: 2-10ms
- Rendering: 5-20ms
- **Total: 60-230ms**

**Recommendation:** Use Redis for production.

## Next Steps

- [03-quick-start.md](03-quick-start.md) - Build a form
- [04-schema-structure.md](04-schema-structure.md) - Schema reference
- [09-storage-interface.md](09-storage-interface.md) - Storage backends
- [10-registry.md](10-registry.md) - Registry implementation

---

[← Back to Introduction](01-introduction.md) | [Next: Quick Start →](03-quick-start.md)