# Field Types Reference


---

## Table of Contents

### Core Documentation
- [Overview](#overview)
- [Field Type Anatomy](#field-type-anatomy)
- [Version Compatibility](#version-compatibility)
- [Quick Start Guide](#quick-start-guide)

### Field Categories
- [Text Input Fields](#text-input-fields)
- [Numeric Fields](#numeric-fields)
- [Date & Time Fields](#date--time-fields)
- [Selection Fields](#selection-fields)
- [File & Media Fields](#file--media-fields)
- [Text Content Fields](#text-content-fields)
- [Geographic & Location Fields](#geographic--location-fields)
- [Biometric & Capture Fields](#biometric--capture-fields)
- [Specialized Input Fields](#specialized-input-fields)
- [Relational Fields](#relational-fields)
- [Layout & Container Fields](#layout--container-fields)
- [Display-Only Fields](#display-only-fields)

### Technical Reference
- [Database Type Mapping](#database-type-mapping)
- [Validation Rules Reference](#validation-rules-reference)
- [Security Guidelines](#security-guidelines)
- [Performance Optimization](#performance-optimization)

---

## Overview

Awo ERP's field type system provides a declarative approach to building forms that automatically generates:

- **UI Components** - Production-ready React components with professional styling
- **Validation** - Client-side and server-side validation with custom rules
- **Database Schema** - Automatic SQL DDL generation with proper indexing
- **API Contracts** - Type-safe REST and GraphQL API definitions
- **Documentation** - Auto-generated API and field documentation

### Design Principles

1. **Type Safety** - Strong typing with runtime validation and TypeScript support
2. **Progressive Enhancement** - Functional without JavaScript, enhanced with it
3. **Accessibility First** - WCAG 2.1 AA compliant by default with screen reader support
4. **Performance** - Optimized for large datasets with virtualization and lazy loading
5. **Security** - Built-in protection against XSS, CSRF, SQL injection, and file upload attacks
6. **Developer Experience** - Intuitive API with comprehensive error messages
7. **Extensibility** - Custom field types and validation rules supported

### Key Features

- **Conditional Logic** - Show/hide fields based on other field values
- **Dynamic Options** - Load dropdown options from APIs with caching
- **Multi-Language** - Built-in i18n support for labels and validation messages
- **Real-Time Validation** - Async validation with debouncing
- **Responsive Design** - Mobile-first responsive layouts
- **Dark Mode** - Automatic dark mode support
- **Audit Trail** - Track all field changes with timestamps
- **Version Control** - Field schema versioning and migration tools

---

## Field Type Anatomy

Every field type follows this standardized structure:

```json
{
  "name": "field_identifier",
  "type": "field_type",
  "label": "Human Readable Label",
  "placeholder": "Hint text shown when empty",
  "description": "Detailed help text explaining the field purpose",
  "required": false,
  "readonly": false,
  "disabled": false,
  "default": null,
  "hidden": false,
  
  "config": {
    // Field-specific configuration options
    "customOption": "value"
  },
  
  "validation": {
    // Validation rules applied to field value
    "minLength": 3,
    "maxLength": 50,
    "pattern": "^[a-zA-Z]+$",
    "custom": "validateFunction",
    "message": "Custom error message"
  },
  
  "showIf": {
    // Conditional visibility rules
    "field": "other_field_name",
    "operator": "equals",
    "value": "expected_value"
  },
  
  "permissions": {
    "view": ["user", "admin", "manager"],
    "edit": ["admin", "manager"],
    "delete": ["admin"]
  },
  
  "metadata": {
    "version": "1.0.0",
    "category": "contact_info",
    "tags": ["pii", "required", "searchable"],
    "helpUrl": "https://docs.example.com/fields/email",
    "searchable": true,
    "sortable": true,
    "filterable": true,
    "exportable": true
  },
  
  "i18n": {
    "en": {
      "label": "Email Address",
      "placeholder": "Enter your email",
      "description": "We'll never share your email"
    },
    "es": {
      "label": "Dirección de Correo",
      "placeholder": "Ingrese su correo",
      "description": "Nunca compartiremos su correo"
    }
  }
}
```

### Core Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `name` | string | ✓ | Unique identifier (snake_case, alphanumeric + underscore) |
| `type` | string | ✓ | Field type identifier from supported types |
| `label` | string | ✓ | Display label shown to users |
| `placeholder` | string | | Hint text shown in empty input |
| `description` | string | | Help text displayed below field |
| `required` | boolean | | Field must have a value to submit form |
| `readonly` | boolean | | Field is visible but not editable |
| `disabled` | boolean | | Field is grayed out and not submitted |
| `hidden` | boolean | | Field is hidden from UI but included in submission |
| `default` | any | | Default value when creating new records |
| `config` | object | | Type-specific configuration options |
| `validation` | object | | Validation rules and custom validators |
| `showIf` | object | | Conditional display rules |
| `permissions` | object | | Role-based access control |
| `metadata` | object | | Additional metadata for indexing, search, etc. |
| `i18n` | object | | Internationalization translations |

### Conditional Logic

Fields can be shown or hidden based on other field values:

```json
{
  "showIf": {
    "field": "employment_type",
    "operator": "equals",
    "value": "contract"
  }
}
```

**Supported Operators:**
- `equals` - Exact match
- `not_equals` - Not equal to
- `contains` - String contains substring
- `not_contains` - String doesn't contain substring
- `in` - Value in array
- `not_in` - Value not in array
- `greater_than` - Numeric comparison
- `less_than` - Numeric comparison
- `greater_than_or_equal` - Numeric comparison
- `less_than_or_equal` - Numeric comparison
- `is_empty` - Field has no value
- `is_not_empty` - Field has a value
- `matches` - Regular expression match

**Complex Conditions:**

```json
{
  "showIf": {
    "operator": "and",
    "conditions": [
      {"field": "country", "operator": "equals", "value": "US"},
      {"field": "age", "operator": "greater_than", "value": 18}
    ]
  }
}
```

---

## Version Compatibility

| Field Type | Introduced | Last Modified | Status | Breaking Changes | Notes |
|------------|-----------|---------------|--------|------------------|-------|
| `text` | v1.0 | - | Stable | None | Core field type |
| `email` | v1.0 | v1.3 | Stable | None | Added multi-email support in v1.3 |
| `password` | v1.0 | v1.4 | Stable | None | Added strength meter in v1.4 |
| `number` | v1.0 | v1.2 | Stable | None | Added precision control in v1.2 |
| `currency` | v1.1 | v1.5 | Stable | None | Added multi-currency in v1.5 |
| `phone` | v1.0 | v1.5 | Stable | None | Added international formatting in v1.5 |
| `url` | v1.0 | - | Stable | None | Core field type |
| `date` | v1.0 | v1.3 | Stable | None | Added range presets in v1.3 |
| `time` | v1.0 | - | Stable | None | Core field type |
| `datetime` | v1.0 | v1.3 | Stable | None | Added timezone support in v1.3 |
| `daterange` | v1.2 | v1.4 | Stable | None | Added relative ranges in v1.4 |
| `select` | v1.0 | v1.4 | Stable | None | Added async loading in v1.4 |
| `radio` | v1.0 | - | Stable | None | Core field type |
| `checkbox` | v1.0 | - | Stable | None | Core field type |
| `checkboxes` | v1.0 | - | Stable | None | Core field type |
| `switch` | v1.1 | - | Stable | None | New in v1.1 |
| `tree-select` | v1.2 | v1.5 | Stable | None | Added multi-select in v1.5 |
| `file` | v1.0 | v1.5 | Stable | v1.5 | Changed upload API format in v1.5 |
| `image` | v1.0 | v1.4 | Stable | None | Added crop & resize in v1.4 |
| `video` | v1.3 | - | Stable | None | New in v1.3 |
| `audio` | v1.3 | - | Stable | None | New in v1.3 |
| `textarea` | v1.0 | - | Stable | None | Core field type |
| `richtext` | v1.1 | v1.5 | Stable | v1.5 | Switched editor engine in v1.5 |
| `code` | v1.2 | v1.4 | Stable | None | Added more languages in v1.4 |
| `json` | v1.2 | - | Stable | None | New in v1.2 |
| `markdown` | v1.3 | - | Stable | None | New in v1.3 |
| `tags` | v1.1 | v1.3 | Stable | None | Added suggestions in v1.3 |
| `rating` | v1.1 | - | Stable | None | New in v1.1 |
| `slider` | v1.1 | - | Stable | None | New in v1.1 |
| `color` | v1.1 | v1.3 | Stable | None | Added alpha channel in v1.3 |
| `icon-picker` | v1.3 | - | Stable | None | New in v1.3 |
| `location` | v1.4 | v1.5 | Stable | None | Enhanced geocoding in v1.5 |
| `map` | v1.4 | - | Beta | None | New in v1.4 |
| `signature` | v1.3 | - | Stable | None | New in v1.3 |
| `qrcode` | v1.4 | - | Beta | None | New in v1.4 |
| `barcode` | v1.4 | - | Beta | None | New in v1.4 |
| `ipaddress` | v1.5 | - | Stable | None | New in v1.5 |
| `mac-address` | v1.5 | - | Stable | None | New in v1.5 |
| `duration` | v1.5 | - | Beta | None | New in v1.5 |
| `cron` | v1.5 | - | Beta | None | New in v1.5 |
| `relation` | v1.0 | v1.5 | Stable | None | Enhanced filtering in v1.5 |
| `autocomplete` | v1.1 | v1.5 | Stable | None | Added caching in v1.5 |
| `repeatable` | v1.2 | v1.5 | Stable | None | Added virtualization in v1.5 |
| `group` | v1.0 | - | Stable | None | Core field type |
| `tabs` | v1.2 | - | Stable | None | New in v1.2 |
| `accordion` | v1.3 | - | Stable | None | New in v1.3 |
| `fieldset` | v1.0 | - | Stable | None | Core field type |
| `static` | v1.0 | - | Stable | None | Core field type |
| `divider` | v1.0 | - | Stable | None | Core field type |
| `spacer` | v1.2 | - | Stable | None | New in v1.2 |
| `formula` | v1.3 | v1.5 | Stable | None | Enhanced expressions in v1.5 |
| `chart` | v1.4 | - | Beta | None | New in v1.4 |

### Deprecated Field Types

| Field Type | Deprecated | Removed | Replacement | Migration Path |
|------------|-----------|---------|-------------|----------------|
| `html` | v1.4 | v2.0 | `richtext` | Use richtext with HTML mode enabled |
| `multiselect` | v1.3 | v1.5 | `select` | Set `multiple: true` in config |
| `upload` | v1.2 | v1.4 | `file` or `image` | Use specific file/image type |
| `wysiwyg` | v1.1 | v1.3 | `richtext` | Direct replacement, same API |

---

## Quick Start Guide

### Basic Form Example

```json
{
  "form": {
    "name": "contact_form",
    "title": "Contact Information",
    "description": "Enter your contact details",
    "fields": [
      {
        "name": "full_name",
        "type": "text",
        "label": "Full Name",
        "required": true,
        "validation": {
          "minLength": 2,
          "maxLength": 100
        }
      },
      {
        "name": "email",
        "type": "email",
        "label": "Email Address",
        "required": true
      },
      {
        "name": "phone",
        "type": "phone",
        "label": "Phone Number",
        "config": {
          "defaultCountry": "US"
        }
      },
      {
        "name": "message",
        "type": "textarea",
        "label": "Message",
        "config": {
          "rows": 4,
          "maxLength": 500,
          "showCount": true
        }
      }
    ]
  }
}
```

---

## Text Input Fields

### `text`

Single-line text input for general text data.

**HTML Rendering:** `<input type="text">`  
**Database Type:** `VARCHAR(255)` or `TEXT`  
**Validation:** Length, pattern, custom functions

```json
{
  "name": "username",
  "type": "text",
  "label": "Username",
  "placeholder": "Enter username",
  "required": true,
  "config": {
    "minLength": 3,
    "maxLength": 50,
    "autocomplete": "username",
    "spellcheck": false,
    "prefix": "@",
    "suffix": null,
    "transform": "lowercase",
    "trim": true,
    "mask": null,
    "debounce": 300,
    "validateOnBlur": true,
    "validateOnChange": false
  },
  "validation": {
    "minLength": 3,
    "maxLength": 50,
    "pattern": "^[a-zA-Z0-9_-]+$",
    "message": "Username must be 3-50 characters, alphanumeric with - and _",
    "async": "/api/validate/username"
  },
  "default": null
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `minLength` | integer | null | Minimum character length |
| `maxLength` | integer | null | Maximum character length |
| `autocomplete` | string | null | HTML autocomplete attribute value |
| `spellcheck` | boolean | true | Enable browser spellcheck |
| `prefix` | string | null | Text or icon displayed before input |
| `suffix` | string | null | Text or icon displayed after input |
| `transform` | string | null | Transform input: `lowercase`, `uppercase`, `capitalize`, `trim` |
| `trim` | boolean | true | Automatically trim whitespace |
| `mask` | string | null | Input mask pattern (e.g., "999-999-9999") |
| `debounce` | integer | 0 | Debounce validation in milliseconds |
| `validateOnBlur` | boolean | true | Validate when field loses focus |
| `validateOnChange` | boolean | false | Validate on every keystroke |

**Input Masking Examples:**

```json
{
  "config": {
    "mask": "999-999-9999",
    "maskPlaceholder": "_",
    "maskChar": "#"
  }
}
```

**Mask Patterns:**
- `9` - Numeric (0-9)
- `a` - Alphabetic (a-z, A-Z)
- `*` - Alphanumeric
- `#` - Numeric or space

**Common Use Cases:**
- Usernames and account identifiers
- Names (first, last, full)
- Titles and headlines
- Short descriptions
- Product codes and SKUs
- Search queries

**Advanced Validation:**

```json
{
  "validation": {
    "minLength": 3,
    "maxLength": 50,
    "pattern": "^[a-zA-Z0-9_-]+$",
    "message": "Invalid username format",
    "async": "/api/validate/username",
    "asyncMethod": "POST",
    "asyncDebounce": 500,
    "custom": "validateUsernameCustom"
  }
}
```

---

### `email`

Email address input with built-in validation and multi-email support.

**HTML Rendering:** `<input type="email">`  
**Database Type:** `VARCHAR(255)` or `TEXT[]` (for multiple emails)  
**Validation:** RFC 5322 email format

```json
{
  "name": "email",
  "type": "email",
  "label": "Email Address",
  "placeholder": "user@example.com",
  "required": true,
  "config": {
    "autocomplete": "email",
    "multiple": false,
    "separator": ",",
    "maxEmails": 5,
    "verifyDomain": false,
    "allowedDomains": [],
    "blockedDomains": ["tempmail.com", "disposable.com"],
    "lowercase": true,
    "showTags": true,
    "tagCloseable": true
  },
  "validation": {
    "format": "rfc5322",
    "dns": false,
    "disposable": false,
    "role": false,
    "message": "Please enter a valid email address"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `multiple` | boolean | false | Allow multiple email addresses |
| `separator` | string | "," | Separator for multiple emails |
| `maxEmails` | integer | null | Maximum number of emails allowed |
| `verifyDomain` | boolean | false | Verify DNS MX records |
| `allowedDomains` | array | [] | Whitelist of allowed domains |
| `blockedDomains` | array | [] | Blacklist of blocked domains |
| `lowercase` | boolean | true | Convert email to lowercase |
| `showTags` | boolean | false | Display emails as tags (multiple mode) |
| `tagCloseable` | boolean | true | Allow removing individual email tags |

**Multiple Email Configuration:**

```json
{
  "config": {
    "multiple": true,
    "maxEmails": 5,
    "separator": ",",
    "showTags": true,
    "tagCloseable": true,
    "allowDuplicates": false
  }
}
```

**Domain Validation:**

```json
{
  "validation": {
    "allowedDomains": ["company.com", "partner.org", "client.net"],
    "message": "Only company, partner, or client email addresses are allowed"
  }
}
```

**Advanced Validation Options:**

```json
{
  "validation": {
    "format": "rfc5322",
    "dns": true,
    "disposable": true,
    "role": true,
    "blockedDomains": ["tempmail.com", "10minutemail.com"],
    "customValidator": "/api/validate/email"
  }
}
```

**Validation Types:**
- `format` - Check email format (RFC 5322 or simple)
- `dns` - Verify domain has valid MX records
- `disposable` - Block disposable/temporary email services
- `role` - Block role-based emails (admin@, support@, etc.)

---

### `password`

Secure password input with strength indicator and requirements.

**HTML Rendering:** `<input type="password">`  
**Database Type:** `VARCHAR(255)` (store hashed with bcrypt/argon2)  
**Security:** NEVER store passwords in plain text

```json
{
  "name": "password",
  "type": "password",
  "label": "Password",
  "placeholder": "Enter secure password",
  "required": true,
  "config": {
    "showStrength": true,
    "showToggle": true,
    "showRequirements": true,
    "autocomplete": "new-password",
    "strengthLevels": {
      "weak": 0,
      "fair": 40,
      "good": 60,
      "strong": 80,
      "excellent": 90
    },
    "requirements": [
      {
        "rule": "minLength",
        "value": 8,
        "message": "At least 8 characters"
      },
      {
        "rule": "uppercase",
        "message": "One uppercase letter (A-Z)"
      },
      {
        "rule": "lowercase",
        "message": "One lowercase letter (a-z)"
      },
      {
        "rule": "number",
        "message": "One number (0-9)"
      },
      {
        "rule": "special",
        "message": "One special character (!@#$%^&*)"
      }
    ],
    "prohibitCommon": true,
    "prohibitUserInfo": true,
    "zxcvbn": true
  },
  "validation": {
    "minLength": 8,
    "maxLength": 128,
    "pattern": "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%*?&])[A-Za-z\\d@$!%*?&]+$",
    "message": "Password must meet all security requirements",
    "commonPasswords": true,
    "breached": false
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `showStrength` | boolean | true | Display password strength meter |
| `showToggle` | boolean | true | Show/hide password toggle button |
| `showRequirements` | boolean | true | Display password requirements checklist |
| `strengthLevels` | object | default | Score thresholds for strength levels |
| `requirements` | array | [] | Array of requirement rules |
| `prohibitCommon` | boolean | true | Block common/leaked passwords |
| `prohibitUserInfo` | boolean | false | Block passwords containing user info |
| `zxcvbn` | boolean | false | Use zxcvbn library for strength |

**Password Confirmation Pattern:**

```json
[
  {
    "name": "password",
    "type": "password",
    "label": "Password",
    "config": {
      "showStrength": true,
      "showRequirements": true
    }
  },
  {
    "name": "password_confirmation",
    "type": "password",
    "label": "Confirm Password",
    "config": {
      "showStrength": false,
      "showRequirements": false
    },
    "validation": {
      "matches": "password",
      "message": "Passwords must match"
    }
  }
]
```

**Security Best Practices:**
- Always hash passwords using bcrypt, argon2, or PBKDF2
- Never log password values in application logs
- Implement rate limiting on authentication endpoints
- Require HTTPS for password transmission
- Support passkeys and multi-factor authentication
- Set reasonable maximum length (128-256 characters)
- Use constant-time comparison for password verification

**Common Password Requirements:**

```json
{
  "requirements": [
    {"rule": "minLength", "value": 12, "message": "At least 12 characters"},
    {"rule": "maxLength", "value": 128, "message": "Maximum 128 characters"},
    {"rule": "uppercase", "message": "At least one uppercase letter"},
    {"rule": "lowercase", "message": "At least one lowercase letter"},
    {"rule": "number", "message": "At least one number"},
    {"rule": "special", "chars": "!@#$%^&*()_+-=[]{}|;:,.<>?", "message": "At least one special character"},
    {"rule": "noSpaces", "message": "No spaces allowed"},
    {"rule": "notCommon", "message": "Not a commonly used password"},
    {"rule": "notUsername", "message": "Cannot contain username"},
    {"rule": "notEmail", "message": "Cannot contain email address"}
  ]
}
```

---

### `url`

URL/website address input with protocol handling and validation.

**HTML Rendering:** `<input type="url">`  
**Database Type:** `VARCHAR(2048)` or `TEXT`  
**Validation:** URL format, protocol, DNS resolution

```json
{
  "name": "website",
  "type": "url",
  "label": "Website URL",
  "placeholder": "https://example.com",
  "config": {
    "protocols": ["http", "https"],
    "requireProtocol": true,
    "defaultProtocol": "https",
    "allowRelative": false,
    "validateDNS": false,
    "preview": true,
    "favicon": true,
    "openInNewTab": true,
    "allowedDomains": [],
    "blockedDomains": [],
    "maxLength": 2048
  },
  "validation": {
    "protocols": ["http", "https"],
    "requireProtocol": true,
    "pattern": "^https?://[\\w\\-]+(\\.[\\w\\-]+)+[/#?]?.*$",
    "dns": false,
    "ssl": false,
    "message": "Please enter a valid URL"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `protocols` | array | ["http","https"] | Allowed URL protocols |
| `requireProtocol` | boolean | true | Protocol must be specified |
| `defaultProtocol` | string | "https" | Auto-add if missing |
| `allowRelative` | boolean | false | Allow relative URLs (e.g., /page) |
| `validateDNS` | boolean | false | Verify DNS resolution |
| `preview` | boolean | false | Show URL preview card |
| `favicon` | boolean | false | Display site favicon |
| `openInNewTab` | boolean | true | Open links in new tab |
| `allowedDomains` | array | [] | Whitelist of domains |
| `blockedDomains` | array | [] | Blacklist of domains |
| `maxLength` | integer | 2048 | Maximum URL length |

**Advanced URL Validation:**

```json
{
  "validation": {
    "protocols": ["https"],
    "requireProtocol": true,
    "requireTLS": true,
    "maxLength": 2048,
    "allowedDomains": ["company.com", "partner.org"],
    "blockedDomains": ["malicious.com", "spam.net"],
    "validateSSL": true,
    "checkBlacklist": true
  }
}
```

**URL Preview Configuration:**

```json
{
  "config": {
    "preview": true,
    "previewDelay": 500,
    "previewData": ["title", "description", "image", "favicon"],
    "previewCache": true,
    "previewCacheTTL": 3600
  }
}
```

---

### `phone`

Phone number input with international formatting and validation.

**HTML Rendering:** `<input type="tel">`  
**Database Type:** `VARCHAR(50)`  
**Validation:** E.164 format, country-specific patterns

```json
{
  "name": "phone",
  "type": "phone",
  "label": "Phone Number",
  "placeholder": "+1 (555) 123-4567",
  "config": {
    "format": "international",
    "defaultCountry": "US",
    "preferredCountries": ["US", "GB", "KE", "NG"],
    "allowExtension": true,
    "autoFormat": true,
    "validateMobile": false,
    "showCountryCode": true,
    "nationalMode": false,
    "separateDialCode": true,
    "countrySearch": true,
    "onlyCountries": [],
    "excludeCountries": []
  },
  "validation": {
    "pattern": "^\\+?[1-9]\\d{1,14}$",
    "format": "e164",
    "validateLength": true,
    "validateType": false,
    "types": ["mobile", "fixed_line"]
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `format` | string | "international" | Format: `international`, `national`, `e164` |
| `defaultCountry` | string | "US" | ISO 3166-1 alpha-2 country code |
| `preferredCountries` | array | [] | Countries shown at top of list |
| `allowExtension` | boolean | false | Allow phone extensions |
| `autoFormat` | boolean | true | Format number as user types |
| `validateMobile` | boolean | false | Validate as mobile number only |
| `showCountryCode` | boolean | true | Show country code selector |
| `nationalMode` | boolean | false | Use national format |
| `separateDialCode` | boolean | false | Show dial code separately |
| `countrySearch` | boolean | true | Enable country search |
| `onlyCountries` | array | [] | Limit to specific countries |
| `excludeCountries` | array | [] | Exclude specific countries |

**Phone Extension Support:**

```json
{
  "config": {
    "allowExtension": true,
    "extensionLabel": "Ext.",
    "extensionPlaceholder": "1234",
    "extensionMaxLength": 6
  }
}
```

**SMS Verification:**

```json
{
  "config": {
    "verify": true,
    "verifyUrl": "/api/phone/verify",
    "verifyMethod": "sms",
    "codeLength": 6,
    "resendDelay": 60,
    "maxAttempts": 3
  }
}
```

**Common Country Codes:**
- US/Canada: +1
- UK: +44
- Kenya: +254
- Nigeria: +234
- South Africa: +27
- Australia: +61
- India: +91

---

### `hidden`

Hidden field for storing values not displayed to users.

**HTML Rendering:** `<input type="hidden">`  
**Database Type:** Any type based on value  
**Security:** Always validate on server-side

```json
{
  "name": "user_id",
  "type": "hidden",
  "default": "${session.user_id}",
  "readonly": true
}
```

**Dynamic Values with Template Variables:**

```json
{
  "name": "csrf_token",
  "type": "hidden",
  "default": "${csrf.token}"
},
{
  "name": "timestamp",
  "type": "hidden",
  "default": "${now.timestamp}"
},
{
  "name": "referrer",
  "type": "hidden",
  "default": "${request.referrer}"
},
{
  "name": "form_version",
  "type": "hidden",
  "default": "2.1.0"
}
```

**Security Warnings:**
- Never trust hidden field values without server-side validation
- Always re-validate and sanitize hidden field data
- Don't store sensitive data in hidden fields (visible in HTML source)
- Use signed tokens or encryption for critical hidden values
- Implement CSRF protection separately from hidden fields
- Hidden fields can be modified by users using browser dev tools

---

### `ipaddress`

IP address input with IPv4/IPv6 validation and CIDR notation support.

**HTML Rendering:** `<input type="text">`  
**Database Type:** `INET` (PostgreSQL) or `VARCHAR(45)`  
**Validation:** IP format, version, range

```json
{
  "name": "ip_address",
  "type": "ipaddress",
  "label": "IP Address",
  "placeholder": "192.168.1.1",
  "config": {
    "version": "both",
    "cidr": false,
    "private": true,
    "autoDetect": false,
    "showVersion": true,
    "formatOnBlur": true
  },
  "validation": {
    "version": "both",
    "allowPrivate": true,
    "allowReserved": false,
    "allowLoopback": false,
    "cidr": false,
    "range": null
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `version` | string | "both" | IP version: `ipv4`, `ipv6`, `both` |
| `cidr` | boolean | false | Allow CIDR notation (e.g., 192.168.1.0/24) |
| `private` | boolean | true | Allow private IP addresses |
| `autoDetect` | boolean | false | Auto-fill with user's IP address |
| `showVersion` | boolean | false | Display IP version indicator |
| `formatOnBlur` | boolean | true | Format IP when field loses focus |

**CIDR Notation Support:**

```json
{
  "config": {
    "cidr": true,
    "showSubnetInfo": true,
    "calculateHosts": true
  }
}
```

**IP Range Validation:**

```json
{
  "validation": {
    "version": "ipv4",
    "range": {
      "start": "192.168.0.0",
      "end": "192.168.255.255"
    },
    "allowPrivate": true,
    "allowReserved": false
  }
}
```

**Private IP Ranges:**
- IPv4: 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
- IPv6: fc00::/7, fe80::/10

---

### `mac-address`

MAC (Media Access Control) address input with formatting.

**HTML Rendering:** `<input type="text">`  
**Database Type:** `MACADDR` (PostgreSQL) or `VARCHAR(17)`  
**Validation:** MAC address format

```json
{
  "name": "mac_address",
  "type": "mac-address",
  "label": "MAC Address",
  "placeholder": "00:1A:2B:3C:4D:5E",
  "config": {
    "separator": ":",
    "uppercase": true,
    "format": "colon",
    "autoFormat": true,
    "showVendor": false
  },
  "validation": {
    "format": "ieee802",
    "allowEUI64": false
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `separator` | string | ":" | Separator character |
| `uppercase` | boolean | true | Convert to uppercase |
| `format` | string | "colon" | Format: `colon`, `hyphen`, `dot`, `none` |
| `autoFormat` | boolean | true | Format as user types |
| `showVendor` | boolean | false | Display vendor name (OUI lookup) |

**Supported Formats:**
- Colon: `00:1A:2B:3C:4D:5E`
- Hyphen: `00-1A-2B-3C-4D-5E`
- Dot: `001A.2B3C.4D5E` (Cisco format)
- None: `001A2B3C4D5E`

---

## Numeric Fields

### `number`

Numeric input with optional step controls and formatting.

**HTML Rendering:** `<input type="number">`  
**Database Type:** `INTEGER`, `BIGINT`, `NUMERIC`, `REAL`, `DOUBLE PRECISION`  
**Validation:** Min, max, step, precision

```json
{
  "name": "quantity",
  "type": "number",
  "label": "Quantity",
  "placeholder": "Enter quantity",
  "required": true,
  "config": {
    "step": 1,
    "precision": 0,
    "showStepper": true,
    "stepperPosition": "right",
    "stepperOrientation": "vertical",
    "thousandsSeparator": true,
    "decimalSeparator": ".",
    "prefix": null,
    "suffix": "items",
    "keyboard": true,
    "allowNegative": false,
    "allowDecimal": false,
    "notation": "standard",
    "compactDisplay": "short"
  },
  "validation": {
    "min": 0,
    "max": 9999,
    "step": 1,
    "integer": true,
    "positive": true,
    "multipleOf": 1
  },
  "default": 1
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `step` | number | 1 | Increment/decrement step value |
| `precision` | integer | 0 | Number of decimal places |
| `showStepper` | boolean | true | Show increment/decrement buttons |
| `stepperPosition` | string | "right" | Position: `left`, `right`, `both` |
| `stepperOrientation` | string | "vertical" | Orientation: `vertical`, `horizontal` |
| `thousandsSeparator` | boolean | false | Format with thousands separator |
| `decimalSeparator` | string | "." | Decimal separator character |
| `prefix` | string | null | Text/symbol before number |
| `suffix` | string | null | Text/symbol after number |
| `keyboard` | boolean | true | Arrow key increment/decrement |
| `allowNegative` | boolean | true | Allow negative numbers |
| `allowDecimal` | boolean | true | Allow decimal numbers |
| `notation` | string | "standard" | Number notation: `standard`, `scientific`, `engineering`, `compact` |
| `compactDisplay` | string | "short" | Compact display: `short`, `long` |

**Decimal Number Configuration:**

```json
{
  "name": "weight",
  "type": "number",
  "label": "Weight (kg)",
  "config": {
    "step": 0.1,
    "precision": 2,
    "min": 0,
    "allowNegative": false,
    "suffix": "kg",
    "decimalSeparator": "."
  },
  "validation": {
    "min": 0,
    "max": 1000,
    "precision": 2
  }
}
```

**Large Number Formatting:**

```json
{
  "name": "population",
  "type": "number",
  "label": "Population",
  "config": {
    "thousandsSeparator": true,
    "notation": "compact",
    "compactDisplay": "short",
    "suffix": null
  }
}
```

**Compact Notation Examples:**
- 1,234 → 1.2K
- 1,234,567 → 1.2M
- 1,234,567,890 → 1.2B

---

### `currency`

Monetary value input with multi-currency support and exchange rates.

**HTML Rendering:** `<input type="text">` with currency formatting  
**Database Type:** `NUMERIC(19,4)` or `MONEY`  
**Validation:** Amount range, currency rules

```json
{
  "name": "amount",
  "type": "currency",
  "label": "Amount",
  "placeholder": "0.00",
  "required": true,
  "config": {
    "currency": "USD",
    "locale": "en-US",
    "decimals": 2,
    "showSymbol": true,
    "symbolPosition": "before",
    "thousandsSeparator": true,
    "allowNegative": false,
    "currencySelect": false,
    "exchangeRate": null,
    "baseCurrency": "USD",
    "showExchangeRate": false,
    "roundingMode": "halfUp",
    "notation": "standard"
  },
  "validation": {
    "min": 0,
    "max": 999999999.99,
    "step": 0.01,
    "precision": 2
  },
  "default": 0
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `currency` | string | "USD" | ISO 4217 currency code |
| `locale` | string | "en-US" | Locale for number formatting |
| `decimals` | integer | 2 | Number of decimal places |
| `showSymbol` | boolean | true | Display currency symbol |
| `symbolPosition` | string | "before" | Symbol position: `before`, `after` |
| `thousandsSeparator` | boolean | true | Use thousands separator |
| `allowNegative` | boolean | false | Allow negative amounts |
| `currencySelect` | boolean | false | Show currency dropdown |
| `exchangeRate` | number | null | Exchange rate to base currency |
| `baseCurrency` | string | "USD" | Base currency for conversions |
| `showExchangeRate` | boolean | false | Display exchange rate info |
| `roundingMode` | string | "halfUp" | Rounding: `up`, `down`, `halfUp`, `halfDown` |
| `notation` | string | "standard" | Display notation |

**Multi-Currency Configuration:**

```json
{
  "name": "price",
  "type": "currency",
  "config": {
    "currencySelect": true,
    "currencies": [
      {
        "code": "USD",
        "symbol": "$",
        "decimals": 2,
        "name": "US Dollar"
      },
      {
        "code": "EUR",
        "symbol": "€",
        "decimals": 2,
        "name": "Euro"
      },
      {
        "code": "GBP",
        "symbol": "£",
        "decimals": 2,
        "name": "British Pound"
      },
      {
        "code": "KES",
        "symbol": "KSh",
        "decimals": 2,
        "name": "Kenyan Shilling"
      }
    ],
    "defaultCurrency": "USD",
    "showExchangeRate": true,
    "exchangeRateSource": "/api/exchange-rates"
  }
}
```

**Major World Currencies:**

```json
{
  "currencies": [
    {"code": "USD", "symbol": "$", "decimals": 2, "name": "US Dollar"},
    {"code": "EUR", "symbol": "€", "decimals": 2, "name": "Euro"},
    {"code": "GBP", "symbol": "£", "decimals": 2, "name": "British Pound"},
    {"code": "JPY", "symbol": "¥", "decimals": 0, "name": "Japanese Yen"},
    {"code": "CHF", "symbol": "Fr", "decimals": 2, "name": "Swiss Franc"},
    {"code": "CAD", "symbol": "C$", "decimals": 2, "name": "Canadian Dollar"},
    {"code": "AUD", "symbol": "A$", "decimals": 2, "name": "Australian Dollar"},
    {"code": "CNY", "symbol": "¥", "decimals": 2, "name": "Chinese Yuan"},
    {"code": "INR", "symbol": "₹", "decimals": 2, "name": "Indian Rupee"},
    {"code": "KES", "symbol": "KSh", "decimals": 2, "name": "Kenyan Shilling"},
    {"code": "NGN", "symbol": "₦", "decimals": 2, "name": "Nigerian Naira"},
    {"code": "ZAR", "symbol": "R", "decimals": 2, "name": "South African Rand"},
    {"code": "BRL", "symbol": "R$", "decimals": 2, "name": "Brazilian Real"}
  ]
}
```

**Exchange Rate Integration:**

```json
{
  "config": {
    "exchangeRate": "${rates.EUR_USD}",
    "exchangeRateSource": "/api/exchange-rates",
    "updateInterval": 3600,
    "showLastUpdated": true,
    "autoConvert": true
  }
}
```

---

### `percentage`

Percentage input with automatic formatting and validation.

**HTML Rendering:** `<input type="number">`  
**Database Type:** `NUMERIC(5,2)`  
**Validation:** 0-100 range (or custom range)

```json
{
  "name": "discount",
  "type": "percentage",
  "label": "Discount Percentage",
  "placeholder": "0",
  "config": {
    "decimals": 2,
    "min": 0,
    "max": 100,
    "suffix": "%",
    "step": 0.1,
    "showStepper": true,
    "allowOver100": false,
    "displayMultiplier": 1
  },
  "validation": {
    "min": 0,
    "max": 100,
    "precision": 2
  },
  "default": 0
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `decimals` | integer | 2 | Decimal places |
| `min` | number | 0 | Minimum percentage |
| `max` | number | 100 | Maximum percentage |
| `suffix` | string | "%" | Suffix symbol |
| `step` | number | 0.1 | Step increment |
| `showStepper` | boolean | true | Show +/- buttons |
| `allowOver100` | boolean | false | Allow values over 100% |
| `displayMultiplier` | number | 1 | Display multiplier (e.g., 0.5 shows as 50%) |

**Percentage as Decimal Storage:**

```json
{
  "config": {
    "displayMultiplier": 100,
    "storeAsDecimal": true
  }
}
```

---

### `duration`

Duration/time period input with multiple unit support.

**HTML Rendering:** Custom component  
**Database Type:** `INTERVAL` (PostgreSQL) or `INTEGER` (seconds)  
**Validation:** Min/max duration

```json
{
  "name": "session_timeout",
  "type": "duration",
  "label": "Session Timeout",
  "config": {
    "units": ["hours", "minutes", "seconds"],
    "defaultUnit": "minutes",
    "maxUnit": "hours",
    "minUnit": "seconds",
    "format": "iso8601",
    "showAll": false,
    "compact": false,
    "separator": " ",
    "allowNegative": false
  },
  "validation": {
    "min": 60,
    "max": 86400,
    "unit": "seconds"
  },
  "default": 1800
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `units` | array | all | Available time units |
| `defaultUnit` | string | "seconds" | Default display unit |
| `maxUnit` | string | null | Largest allowed unit |
| `minUnit` | string | null | Smallest allowed unit |
| `format` | string | "iso8601" | Output format |
| `showAll` | boolean | false | Show all units even if zero |
| `compact` | boolean | false | Use compact display |
| `separator` | string | " " | Separator between units |
| `allowNegative` | boolean | false | Allow negative durations |

**Supported Time Units:**
- `years` (y)
- `months` (mo)
- `weeks` (w)
- `days` (d)
- `hours` (h)
- `minutes` (m)
- `seconds` (s)
- `milliseconds` (ms)

**Duration Format Examples:**

```json
{
  "formats": {
    "iso8601": "PT1H30M45S",
    "human": "1 hour 30 minutes 45 seconds",
    "compact": "1h 30m 45s",
    "colon": "01:30:45",
    "seconds": "5445"
  }
}
```

**Complex Duration Example:**

```json
{
  "name": "project_duration",
  "type": "duration",
  "label": "Project Duration",
  "config": {
    "units": ["months", "weeks", "days"],
    "defaultUnit": "weeks",
    "showAll": true,
    "format": "human",
    "validation": {
      "min": 7,
      "max": 365,
      "unit": "days"
    }
  }
}
```

---

## Date & Time Fields

### `date`

Date picker with calendar interface and range validation.

**HTML Rendering:** `<input type="date">` with calendar picker  
**Database Type:** `DATE`  
**Validation:** Date range, format, business days

```json
{
  "name": "birth_date",
  "type": "date",
  "label": "Date of Birth",
  "placeholder": "MM/DD/YYYY",
  "required": true,
  "config": {
    "format": "YYYY-MM-DD",
    "displayFormat": "MM/DD/YYYY",
    "locale": "en-US",
    "minDate": "1900-01-01",
    "maxDate": "today",
    "firstDayOfWeek": 0,
    "showWeekNumbers": false,
    "disabledDates": [],
    "disabledDaysOfWeek": [],
    "highlightedDates": [],
    "shortcuts": true,
    "clearable": true,
    "inline": false,
    "numberOfMonths": 1,
    "showToday": true,
    "todayButton": true
  },
  "validation": {
    "min": "1900-01-01",
    "max": "${today}",
    "businessDays": false,
    "holidays": [],
    "format": "YYYY-MM-DD"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `format` | string | "YYYY-MM-DD" | Internal storage format |
| `displayFormat` | string | locale | User display format |
| `locale` | string | "en-US" | Locale for calendar |
| `minDate` | string | null | Minimum selectable date |
| `maxDate` | string | null | Maximum selectable date |
| `firstDayOfWeek` | integer | 0 | First day (0=Sunday, 1=Monday) |
| `showWeekNumbers` | boolean | false | Display week numbers |
| `disabledDates` | array | [] | Specific disabled dates |
| `disabledDaysOfWeek` | array | [] | Disabled days (0-6) |
| `highlightedDates` | array | [] | Highlighted dates |
| `shortcuts` | boolean/array | true | Show date shortcuts |
| `clearable` | boolean | true | Show clear button |
| `inline` | boolean | false | Always show calendar |
| `numberOfMonths` | integer | 1 | Number of months to display |
| `showToday` | boolean | true | Highlight today's date |
| `todayButton` | boolean | true | Show "Today" button |

**Date Format Tokens:**
- `YYYY` - 4-digit year (2025)
- `YY` - 2-digit year (25)
- `MMMM` - Full month name (January)
- `MMM` - Short month name (Jan)
- `MM` - 2-digit month (01)
- `M` - Month (1)
- `DD` - 2-digit day (05)
- `D` - Day (5)
- `dddd` - Full day name (Monday)
- `ddd` - Short day name (Mon)

**Date Shortcuts:**

```json
{
  "config": {
    "shortcuts": [
      {"label": "Today", "value": "today"},
      {"label": "Yesterday", "value": "yesterday"},
      {"label": "Tomorrow", "value": "tomorrow"},
      {"label": "Last Week", "value": "today - 7 days"},
      {"label": "Last Month", "value": "today - 1 month"},
      {"label": "Last Year", "value": "today - 1 year"},
      {"label": "Start of Month", "value": "startOfMonth"},
      {"label": "End of Month", "value": "endOfMonth"},
      {"label": "Start of Year", "value": "startOfYear"}
    ]
  }
}
```

**Disabled Dates Configuration:**

```json
{
  "config": {
    "disabledDates": [
      "2025-12-25",
      "2025-01-01",
      "2025-07-04"
    ],
    "disabledDaysOfWeek": [0, 6],
    "disabledDateRanges": [
      {
        "start": "2025-07-01",
        "end": "2025-07-31",
        "reason": "Summer holiday"
      }
    ]
  }
}
```

**Business Days Only:**

```json
{
  "validation": {
    "businessDays": true,
    "holidays": [
      "2025-01-01",
      "2025-07-04",
      "2025-12-25"
    ],
    "customWorkingDays": [1, 2, 3, 4, 5],
    "message": "Please select a business day"
  }
}
```

---

### `time`

Time picker input with hour/minute selection.

**HTML Rendering:** `<input type="time">`  
**Database Type:** `TIME`  
**Validation:** Time range, format

```json
{
  "name": "meeting_time",
  "type": "time",
  "label": "Meeting Time",
  "placeholder": "HH:MM",
  "config": {
    "format": "HH:mm",
    "displayFormat": "hh:mm A",
    "use24Hour": false,
    "minuteStep": 15,
    "hourStep": 1,
    "showSeconds": false,
    "secondStep": 1,
    "clearable": true,
    "hourCycle": 12,
    "showNow": true,
    "disabledHours": [],
    "disabledMinutes": [],
    "hideDisabledOptions": false
  },
  "validation": {
    "min": "08:00",
    "max": "18:00",
    "excludeRanges": []
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `format` | string | "HH:mm" | Internal time format |
| `displayFormat` | string | locale | Display format |
| `use24Hour` | boolean | false | Use 24-hour format |
| `minuteStep` | integer | 1 | Minute increment |
| `hourStep` | integer | 1 | Hour increment |
| `showSeconds` | boolean | false | Show seconds |
| `secondStep` | integer | 1 | Second increment |
| `clearable` | boolean | true | Show clear button |
| `hourCycle` | integer | 12 | 12 or 24 hour cycle |
| `showNow` | boolean | true | Show "Now" button |
| `disabledHours` | array | [] | Disabled hours |
| `disabledMinutes` | array | [] | Disabled minutes |
| `hideDisabledOptions` | boolean | false | Hide disabled options |

**Time Range Validation:**

```json
{
  "validation": {
    "min": "09:00",
    "max": "17:00",
    "excludeRanges": [
      {"start": "12:00", "end": "13:00", "reason": "Lunch break"}
    ],
    "message": "Please select a time during business hours"
  }
}
```

**Disabled Hours Example:**

```json
{
  "config": {
    "disabledHours": [0, 1, 2, 3, 4, 5, 6, 7, 20, 21, 22, 23],
    "disabledMinutes": [1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14],
    "minuteStep": 15
  }
}
```

---

### `datetime`

Combined date and time picker with timezone support.

**HTML Rendering:** `<input type="datetime-local">`  
**Database Type:** `TIMESTAMP` or `TIMESTAMPTZ`  
**Validation:** DateTime range, timezone

```json
{
  "name": "scheduled_at",
  "type": "datetime",
  "label": "Schedule Date & Time",
  "placeholder": "Select date and time",
  "required": true,
  "config": {
    "format": "YYYY-MM-DD HH:mm:ss",
    "displayFormat": "MM/DD/YYYY hh:mm A",
    "timezone": "UTC",
    "showTimezone": true,
    "timezoneSelect": false,
    "minDate": "today",
    "maxDate": "today + 1 year",
    "minuteStep": 15,
    "use24Hour": false,
    "showSeconds": false,
    "convertToUTC": true,
    "preserveLocal": false,
    "defaultTime": "09:00"
  },
  "validation": {
    "min": "${now}",
    "max": "${now + 1 year}",
    "timezone": "UTC"
  },
  "default": "${now}"
}
```

**Timezone Configuration:**

```json
{
  "config": {
    "timezone": "America/New_York",
    "showTimezone": true,
    "timezoneSelect": true,
    "timezones": [
      {"value": "UTC", "label": "UTC"},
      {"value": "America/New_York", "label": "Eastern Time"},
      {"value": "America/Chicago", "label": "Central Time"},
      {"value": "America/Denver", "label": "Mountain Time"},
      {"value": "America/Los_Angeles", "label": "Pacific Time"},
      {"value": "Europe/London", "label": "London (GMT/BST)"},
      {"value": "Asia/Tokyo", "label": "Tokyo (JST)"},
      {"value": "Africa/Nairobi", "label": "Nairobi (EAT)"}
    ],
    "convertToUTC": true,
    "preserveLocal": false
  }
}
```

**Common Timezones:**
- `UTC` - Coordinated Universal Time
- `America/New_York` - Eastern Time (EST/EDT)
- `America/Chicago` - Central Time (CST/CDT)
- `America/Denver` - Mountain Time (MST/MDT)
- `America/Los_Angeles` - Pacific Time (PST/PDT)
- `Europe/London` - British Time (GMT/BST)
- `Europe/Paris` - Central European Time
- `Asia/Tokyo` - Japan Standard Time
- `Asia/Shanghai` - China Standard Time
- `Asia/Dubai` - Gulf Standard Time
- `Africa/Nairobi` - East Africa Time
- `Africa/Lagos` - West Africa Time
- `Australia/Sydney` - Australian Eastern Time

---

### `daterange`

Date range picker for selecting start and end dates.

**HTML Rendering:** Two date inputs or custom range picker  
**Database Type:** Two `DATE` columns or `DATERANGE` (PostgreSQL)  
**Validation:** Range constraints, span limits

```json
{
  "name": "project_duration",
  "type": "daterange",
  "label": "Project Duration",
  "placeholder": "Select date range",
  "required": true,
  "config": {
    "startField": "start_date",
    "endField": "end_date",
    "separator": " to ",
    "minDays": 1,
    "maxDays": 365,
    "minDate": "today",
    "maxDate": "today + 2 years",
    "singleCalendar": false,
    "linkedCalendars": true,
    "autoApply": false,
    "showRanges": true,
    "alwaysShowCalendars": false,
    "opens": "right",
    "drops": "down",
    "ranges": {
      "Today": ["today", "today"],
      "Yesterday": ["yesterday", "yesterday"],
      "This Week": ["startOfWeek", "endOfWeek"],
      "Last Week": ["startOfWeek - 1 week", "endOfWeek - 1 week"],
      "This Month": ["startOfMonth", "endOfMonth"],
      "Last Month": ["startOfMonth - 1 month", "endOfMonth - 1 month"],
      "This Quarter": ["startOfQuarter", "endOfQuarter"],
      "Last Quarter": ["startOfQuarter - 3 months", "endOfQuarter - 3 months"],
      "This Year": ["startOfYear", "endOfYear"],
      "Last Year": ["startOfYear - 1 year", "endOfYear - 1 year"],
      "Last 7 Days": ["today - 7 days", "today"],
      "Last 30 Days": ["today - 30 days", "today"],
      "Last 90 Days": ["today - 90 days", "today"]
    }
  },
  "validation": {
    "required": true,
    "maxSpan": 365,
    "minSpan": 1,
    "startBeforeEnd": true
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `startField` | string | "start" | Start date field name |
| `endField` | string | "end" | End date field name |
| `separator` | string | " to " | Visual separator |
| `minDays` | integer | null | Minimum range in days |
| `maxDays` | integer | null | Maximum range in days |
| `minDate` | string | null | Earliest selectable date |
| `maxDate` | string | null | Latest selectable date |
| `singleCalendar` | boolean | false | Show single calendar |
| `linkedCalendars` | boolean | true | Link calendar navigation |
| `autoApply` | boolean | false | Auto-apply on selection |
| `showRanges` | boolean | true | Show preset ranges |
| `alwaysShowCalendars` | boolean | false | Always show calendars |
| `opens` | string | "right" | Picker position: `left`, `right`, `center` |
| `drops` | string | "down" | Picker direction: `up`, `down` |

**Relative Date Range Expressions:**

```json
{
  "ranges": {
    "Today": ["today", "today"],
    "Yesterday": ["yesterday", "yesterday"],
    "Tomorrow": ["tomorrow", "tomorrow"],
    "This Week": ["startOfWeek", "endOfWeek"],
    "Next Week": ["startOfWeek + 1 week", "endOfWeek + 1 week"],
    "Last 7 Days": ["today - 7 days", "today"],
    "Last 30 Days": ["today - 30 days", "today"],
    "Last 90 Days": ["today - 90 days", "today"],
    "This Month": ["startOfMonth", "endOfMonth"],
    "Last Month": ["startOfMonth - 1 month", "endOfMonth - 1 month"],
    "This Quarter": ["startOfQuarter", "endOfQuarter"],
    "This Year": ["startOfYear", "endOfYear"],
    "All Time": [null, null]
  }
}
```

**Custom Range Validation:**

```json
{
  "validation": {
    "maxSpan": 90,
    "minSpan": 1,
    "startBeforeEnd": true,
    "businessDaysOnly": false,
    "excludeHolidays": false,
    "message": "Date range must be between 1 and 90 days"
  }
}
```

---

### `month`

Month picker for selecting year and month.

**HTML Rendering:** `<input type="month">`  
**Database Type:** `DATE` or `VARCHAR(7)` (YYYY-MM)  
**Validation:** Month range

```json
{
  "name": "report_month",
  "type": "month",
  "label": "Report Month",
  "placeholder": "Select month",
  "config": {
    "format": "YYYY-MM",
    "displayFormat": "MMMM YYYY",
    "locale": "en-US",
    "minMonth": "2020-01",
    "maxMonth": "today",
    "showYear": true,
    "yearFirst": true,
    "clearable": true
  },
  "validation": {
    "min": "2020-01",
    "max": "${currentMonth}"
  }
}
```

---

### `year`

Year picker for selecting a year.

**HTML Rendering:** `<select>` or custom year picker  
**Database Type:** `SMALLINT` or `INTEGER`  
**Validation:** Year range

```json
{
  "name": "graduation_year",
  "type": "year",
  "label": "Graduation Year",
  "placeholder": "Select year",
  "config": {
    "minYear": 1950,
    "maxYear": "${currentYear + 10}",
    "step": 1,
    "reverse": true,
    "format": "YYYY",
    "showCentury": false
  },
  "validation": {
    "min": 1950,
    "max": 2035
  }
}
```

---

### `cron`

Cron expression builder for scheduling tasks.

**HTML Rendering:** Custom cron builder component  
**Database Type:** `VARCHAR(100)`  
**Validation:** Valid cron syntax

```json
{
  "name": "backup_schedule",
  "type": "cron",
  "label": "Backup Schedule",
  "config": {
    "presets": [
      {"label": "Every minute", "value": "* * * * *"},
      {"label": "Every 5 minutes", "value": "*/5 * * * *"},
      {"label": "Every 15 minutes", "value": "*/15 * * * *"},
      {"label": "Every 30 minutes", "value": "*/30 * * * *"},
      {"label": "Every hour", "value": "0 * * * *"},
      {"label": "Every 6 hours", "value": "0 */6 * * *"},
      {"label": "Every 12 hours", "value": "0 */12 * * *"},
      {"label": "Daily at midnight", "value": "0 0 * * *"},
      {"label": "Daily at noon", "value": "0 12 * * *"},
      {"label": "Weekly on Sunday", "value": "0 0 * * 0"},
      {"label": "Weekly on Monday", "value": "0 0 * * 1"},
      {"label": "Monthly on 1st", "value": "0 0 1 * *"},
      {"label": "Yearly on Jan 1st", "value": "0 0 1 1 *"}
    ],
    "showDescription": true,
    "showNextRuns": 5,
    "timezone": "UTC",
    "allowAdvanced": true,
    "mode": "simple"
  },
  "validation": {
    "format": "cron",
    "allowSpecialChars": true
  }
}
```

**Cron Expression Format:**
```
* * * * * *
│ │ │ │ │ │
│ │ │ │ │ └─── Day of week (0-7, Sunday=0 or 7)
│ │ │ │ └───── Month (1-12)
│ │ │ └─────── Day of month (1-31)
│ │ └───────── Hour (0-23)
│ └─────────── Minute (0-59)
└───────────── Second (0-59) [optional]
```

**Cron Special Characters:**
- `*` - Any value
- `,` - List separator (1,2,3)
- `-` - Range (1-5)
- `/` - Step values (*/5)
- `?` - No specific value
- `L` - Last (day of month/week)
- `W` - Weekday nearest to given day
- `#` - Nth occurrence (1#2 = second Monday)

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `presets` | array | [] | Preset cron expressions |
| `showDescription` | boolean | true | Show human-readable description |
| `showNextRuns` | integer | 5 | Show next N execution times |
| `timezone` | string | "UTC" | Timezone for execution |
| `allowAdvanced` | boolean | true | Allow raw cron editing |
| `mode` | string | "simple" | UI mode: `simple`, `advanced` |

---

## Selection Fields

### `select`

Dropdown selection field with single or multiple selection support.

**HTML Rendering:** `<select>` or custom dropdown component  
**Database Type:** `VARCHAR`, `INTEGER`, or `TEXT[]` (multiple)  
**Validation:** Option exists in list

```json
{
  "name": "country",
  "type": "select",
  "label": "Country",
  "placeholder": "Select a country",
  "required": true,
  "config": {
    "searchable": true,
    "clearable": true,
    "multiple": false,
    "maxItems": null,
    "closeOnSelect": true,
    "createable": false,
    "loading": false,
    "disabled": false,
    "virtualScroll": true,
    "groupBy": null,
    "optionHeight": 40,
    "showCheckboxes": false,
    "allowSelectAll": false,
    "noOptionsText": "No options available",
    "loadingText": "Loading...",
    "searchPlaceholder": "Search..."
  },
  "options": [
    {
      "value": "us",
      "label": "United States",
      "icon": "🇺🇸",
      "disabled": false,
      "group": "North America",
      "description": "USA"
    },
    {
      "value": "gb",
      "label": "United Kingdom",
      "icon": "🇬🇧",
      "group": "Europe"
    },
    {
      "value": "ke",
      "label": "Kenya",
      "icon": "🇰🇪",
      "group": "Africa"
    }
  ],
  "validation": {
    "required": true,
    "in": ["us", "gb", "ke"]
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `searchable` | boolean | false | Enable search/filter |
| `clearable` | boolean | false | Show clear button |
| `multiple` | boolean | false | Allow multiple selections |
| `maxItems` | integer | null | Max selections (multiple mode) |
| `closeOnSelect` | boolean | true | Close dropdown on selection |
| `createable` | boolean | false | Allow creating new options |
| `loading` | boolean | false | Show loading state |
| `disabled` | boolean | false | Disable entire field |
| `virtualScroll` | boolean | false | Use virtual scrolling (large lists) |
| `groupBy` | string | null | Group options by field |
| `optionHeight` | integer | 40 | Height of each option (px) |
| `showCheckboxes` | boolean | false | Show checkboxes (multiple mode) |
| `allowSelectAll` | boolean | false | Add "Select All" option |
| `noOptionsText` | string | default | Text when no options |
| `loadingText` | string | default | Text when loading |
| `searchPlaceholder` | string | default | Search box placeholder |

**Option Format:**

```json
{
  "value": "unique_value",
  "label": "Display Text",
  "icon": "icon-name",
  "image": "/path/to/image.png",
  "disabled": false,
  "group": "category",
  "description": "Additional info text",
  "color": "#FF0000",
  "metadata": {
    "custom": "data"
  }
}
```

**Dynamic Options (Async Loading):**

```json
{
  "name": "category_id",
  "type": "select",
  "label": "Category",
  "config": {
    "optionsSource": "/api/categories",
    "searchSource": "/api/categories/search",
    "dependsOn": ["parent_id"],
    "valueField": "id",
    "labelField": "name",
    "iconField": "icon",
    "groupField": "group",
    "cache": true,
    "cacheTimeout": 300000,
    "debounce": 300,
    "minSearchLength": 2
  }
}
```

**Grouped Options:**

```json
{
  "options": [
    {
      "group": "Fruits",
      "options": [
        {"value": "apple", "label": "Apple", "icon": "🍎"},
        {"value": "banana", "label": "Banana", "icon": "🍌"},
        {"value": "orange", "label": "Orange", "icon": "🍊"}
      ]
    },
    {
      "group": "Vegetables",
      "options": [
        {"value": "carrot", "label": "Carrot", "icon": "🥕"},
        {"value": "lettuce", "label": "Lettuce", "icon": "🥬"},
        {"value": "tomato", "label": "Tomato", "icon": "🍅"}
      ]
    }
  ]
}
```

**Multiple Selection Configuration:**

```json
{
  "config": {
    "multiple": true,
    "maxItems": 5,
    "showTags": true,
    "tagCloseable": true,
    "tagColor": "primary",
    "selectAll": true,
    "selectAllText": "Select All",
    "showCheckboxes": true,
    "allowDuplicates": false
  },
  "validation": {
    "minItems": 1,
    "maxItems": 5,
    "message": "Select between 1 and 5 items"
  }
}
```

**Createable Options:**

```json
{
  "config": {
    "createable": true,
    "createText": "Create \"{input}\"",
    "createValidator": "^[a-zA-Z0-9\\s-]+$",
    "onCreate": "/api/options/create",
    "allowDuplicates": false
  }
}
```

---

### `radio`

Radio button group for single selection from visible options.

**HTML Rendering:** Multiple `<input type="radio">`  
**Database Type:** `VARCHAR`, `INTEGER`, `BOOLEAN`  
**Validation:** Required, option exists

```json
{
  "name": "payment_method",
  "type": "radio",
  "label": "Payment Method",
  "required": true,
  "config": {
    "inline": false,
    "showDescription": true,
    "columns": 1,
    "buttonStyle": false,
    "buttonSize": "medium",
    "spacing": "normal",
    "fullWidth": false
  },
  "options": [
    {
      "value": "cash",
      "label": "Cash",
      "description": "Pay with cash on delivery",
      "icon": "💵",
      "disabled": false
    },
    {
      "value": "card",
      "label": "Credit Card",
      "description": "Secure online card payment",
      "icon": "💳"
    },
    {
      "value": "mpesa",
      "label": "M-Pesa",
      "description": "Mobile money payment",
      "icon": "📱"
    },
    {
      "value": "bank",
      "label": "Bank Transfer",
      "description": "Direct bank transfer",
      "icon": "🏦"
    }
  ],
  "default": "card"
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `inline` | boolean | false | Display horizontally |
| `showDescription` | boolean | false | Show option descriptions |
| `columns` | integer | 1 | Number of columns in grid layout |
| `buttonStyle` | boolean | false | Use button-style radios |
| `buttonSize` | string | "medium" | Button size: `small`, `medium`, `large` |
| `spacing` | string | "normal" | Spacing: `compact`, `normal`, `relaxed` |
| `fullWidth` | boolean | false | Buttons fill container width |

**Button Style Configuration:**

```json
{
  "config": {
    "buttonStyle": true,
    "buttonSize": "large",
    "fullWidth": true,
    "columns": 2,
    "showIcon": true,
    "iconPosition": "left"
  }
}
```

**Card Style Radio Buttons:**

```json
{
  "config": {
    "style": "card",
    "cardPadding": "medium",
    "showBorder": true,
    "highlightSelected": true,
    "columns": 3
  }
}
```

---

### `checkbox`

Single checkbox for boolean values or agreement confirmations.

**HTML Rendering:** `<input type="checkbox">`  
**Database Type:** `BOOLEAN`  
**Validation:** Required (must be checked)

```json
{
  "name": "agree_terms",
  "type": "checkbox",
  "label": "I agree to the terms and conditions",
  "required": true,
  "config": {
    "checkedValue": true,
    "uncheckedValue": false,
    "indeterminate": false,
    "labelPosition": "after",
    "size": "medium",
    "color": "primary",
    "showCheck": true
  },
  "default": false,
  "validation": {
    "required": true,
    "message": "You must agree to the terms and conditions"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `checkedValue` | any | true | Value when checked |
| `uncheckedValue` | any | false | Value when unchecked |
| `indeterminate` | boolean | false | Support indeterminate state |
| `labelPosition` | string | "after" | Label position: `before`, `after` |
| `size` | string | "medium" | Size: `small`, `medium`, `large` |
| `color` | string | "primary" | Color theme |
| `showCheck` | boolean | true | Show checkmark icon |

**Custom Values:**

```json
{
  "config": {
    "checkedValue": "yes",
    "uncheckedValue": "no"
  }
}
```

**Indeterminate State:**

```json
{
  "config": {
    "indeterminate": true,
    "checkedValue": true,
    "uncheckedValue": false,
    "indeterminateValue": null
  }
}
```

---

### `checkboxes`

Multiple checkbox group for multi-selection.

**HTML Rendering:** Multiple `<input type="checkbox">`  
**Database Type:** `TEXT[]` or `JSONB`  
**Validation:** Min/max checked items

```json
{
  "name": "permissions",
  "type": "checkboxes",
  "label": "User Permissions",
  "config": {
    "inline": false,
    "columns": 2,
    "selectAll": true,
    "selectAllText": "Select All",
    "showDescription": true,
    "spacing": "normal",
    "sortable": false
  },
  "options": [
    {
      "value": "read",
      "label": "Read",
      "description": "View records and data"
    },
    {
      "value": "write",
      "label": "Write",
      "description": "Create and edit records"
    },
    {
      "value": "delete",
      "label": "Delete",
      "description": "Delete records permanently"
    },
    {
      "value": "admin",
      "label": "Admin",
      "description": "Full administrative access"
    }
  ],
  "validation": {
    "minChecked": 1,
    "maxChecked": 3,
    "message": "Select between 1 and 3 permissions"
  },
  "default": ["read"]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `inline` | boolean | false | Display horizontally |
| `columns` | integer | 1 | Number of columns |
| `selectAll` | boolean | false | Show "Select All" checkbox |
| `selectAllText` | string | "Select All" | Select all label |
| `showDescription` | boolean | false | Show descriptions |
| `spacing` | string | "normal" | Spacing between items |
| `sortable` | boolean | false | Allow drag-to-reorder |

**Grouped Checkboxes:**

```json
{
  "options": [
    {
      "group": "Content",
      "options": [
        {"value": "content.read", "label": "View Content"},
        {"value": "content.write", "label": "Edit Content"},
        {"value": "content.delete", "label": "Delete Content"}
      ]
    },
    {
      "group": "Users",
      "options": [
        {"value": "users.read", "label": "View Users"},
        {"value": "users.write", "label": "Edit Users"},
        {"value": "users.delete", "label": "Delete Users"}
      ]
    }
  ]
}
```

---

### `switch`

Toggle switch for boolean values.

**HTML Rendering:** `<input type="checkbox">` (styled as switch)  
**Database Type:** `BOOLEAN`  
**Validation:** Required

```json
{
  "name": "is_active",
  "type": "switch",
  "label": "Active Status",
  "config": {
    "onLabel": "Active",
    "offLabel": "Inactive",
    "onValue": true,
    "offValue": false,
    "showLabels": true,
    "size": "medium",
    "loading": false,
    "color": "success",
    "labelPosition": "left"
  },
  "default": true
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `onLabel` | string | null | Label when switch is on |
| `offLabel` | string | null | Label when switch is off |
| `onValue` | any | true | Value when on |
| `offValue` | any | false | Value when off |
| `showLabels` | boolean | false | Display on/off labels |
| `size` | string | "medium" | Size: `small`, `medium`, `large` |
| `loading` | boolean | false | Show loading spinner |
| `color` | string | "primary" | Color theme |
| `labelPosition` | string | "right" | Position: `left`, `right` |

**With Confirmation:**

```json
{
  "config": {
    "confirmChange": true,
    "confirmTitle": "Confirm Status Change",
    "confirmMessage": "Are you sure you want to change the status?",
    "confirmOnText": "Activate",
    "confirmOffText": "Deactivate"
  }
}
```

---

### `tree-select`

Hierarchical tree-based selector for nested data.

**HTML Rendering:** Custom tree component  
**Database Type:** `INTEGER` or `VARCHAR` (for selected node ID)  
**Validation:** Node exists in tree

```json
{
  "name": "department_id",
  "type": "tree-select",
  "label": "Department",
  "placeholder": "Select department",
  "config": {
    "checkable": false,
    "multiple": false,
    "showSearch": true,
    "treeDefaultExpandAll": false,
    "treeDefaultExpandedKeys": [],
    "treeCheckable": false,
    "treeCheckStrictly": false,
    "showCheckedStrategy": "all",
    "maxTagCount": 3,
    "treeNodeFilterProp": "label",
    "showLine": true,
    "showIcon": true,
    "virtual": true,
    "loadData": null
  },
  "treeData": [
    {
      "value": "1",
      "label": "Engineering",
      "disabled": false,
      "selectable": true,
      "checkable": true,
      "icon": "🔧",
      "children": [
        {
          "value": "1-1",
          "label": "Backend Team",
          "children": [
            {"value": "1-1-1", "label": "API Development"},
            {"value": "1-1-2", "label": "Database"}
          ]
        },
        {"value": "1-2", "label": "Frontend Team"},
        {"value": "1-3", "label": "DevOps Team"}
      ]
    },
    {
      "value": "2",
      "label": "Sales",
      "children": [
        {"value": "2-1", "label": "Enterprise Sales"},
        {"value": "2-2", "label": "SMB Sales"}
      ]
    },
    {
      "value": "3",
      "label": "Support"
    }
  ]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `checkable` | boolean | false | Show checkboxes |
| `multiple` | boolean | false | Allow multiple selections |
| `showSearch` | boolean | true | Enable tree search |
| `treeDefaultExpandAll` | boolean | false | Expand all nodes by default |
| `treeDefaultExpandedKeys` | array | [] | Array of expanded node keys |
| `treeCheckStrictly` | boolean | false | Independent parent-child checking |
| `showCheckedStrategy` | string | "all" | Strategy: `all`, `parent`, `child` |
| `maxTagCount` | integer | null | Max tags to display (multiple mode) |
| `treeNodeFilterProp` | string | "label" | Property to filter on |
| `showLine` | boolean | true | Show connecting lines |
| `showIcon` | boolean | true | Show node icons |
| `virtual` | boolean | false | Use virtual scrolling |
| `loadData` | function/url | null | Lazy load child nodes |

**Dynamic Tree Data (Lazy Loading):**

```json
{
  "config": {
    "treeDataSource": "/api/departments/tree",
    "lazy": true,
    "loadData": "/api/departments/{id}/children",
    "cache": true,
    "cacheTimeout": 600000
  }
}
```

**Tree Node Format:**

```json
{
  "value": "unique_id",
  "label": "Display Name",
  "icon": "icon-name",
  "disabled": false,
  "selectable": true,
  "checkable": true,
  "isLeaf": false,
  "children": [],
  "metadata": {}
}
```

---

## File & Media Fields

### `file`

Generic file upload with drag-and-drop, chunked upload, and resume support.

**HTML Rendering:** `<input type="file">` with custom UI  
**Database Type:** `VARCHAR` (file path) or `JSONB` (metadata)  
**Validation:** File size, type, count

```json
{
  "name": "document",
  "type": "file",
  "label": "Upload Document",
  "config": {
    "multiple": false,
    "maxSize": 10485760,
    "maxFiles": 10,
    "accept": [".pdf", ".doc", ".docx", ".xls", ".xlsx", ".txt"],
    "acceptMime": ["application/pdf", "application/msword", "application/vnd.ms-excel"],
    "uploadUrl": "/api/upload",
    "downloadUrl": "/api/files/{id}",
    "deleteUrl": "/api/files/{id}",
    "autoUpload": true,
    "chunkSize": 1048576,
    "resumable": true,
    "showProgress": true,
    "showPreview": true,
    "drag": true,
    "directory": false,
    "withCredentials": true,
    "headers": {},
    "data": {}
  },
  "validation": {
    "maxSize": 10485760,
    "accept": ["application/pdf", "application/msword"],
    "maxFiles": 5,
    "virusScan": true,
    "required": true
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `multiple` | boolean | false | Allow multiple files |
| `maxSize` | integer | null | Max file size in bytes |
| `maxFiles` | integer | null | Max number of files |
| `accept` | array | [] | Allowed file extensions |
| `acceptMime` | array | [] | Allowed MIME types |
| `uploadUrl` | string | null | Upload endpoint URL |
| `downloadUrl` | string | null | Download URL template |
| `deleteUrl` | string | null | Delete endpoint URL |
| `autoUpload` | boolean | true | Upload immediately on select |
| `chunkSize` | integer | 1MB | Chunk size for large files |
| `resumable` | boolean | false | Support resumable uploads |
| `showProgress` | boolean | true | Show upload progress bar |
| `showPreview` | boolean | true | Show file preview |
| `drag` | boolean | true | Enable drag-and-drop |
| `directory` | boolean | false | Allow folder upload |
| `withCredentials` | boolean | false | Send cookies with request |
| `headers` | object | {} | Custom HTTP headers |
| `data` | object | {} | Additional form data |

**Chunked Upload Configuration:**

```json
{
  "config": {
    "chunkSize": 1048576,
    "resumable": true,
    "retryChunks": 3,
    "parallelChunks": 3,
    "testChunks": true,
    "uploadMethod": "POST"
  }
}
```

**File Metadata Structure:**

```json
{
  "id": "file_abc123",
  "name": "document.pdf",
  "size": 1024000,
  "type": "application/pdf",
  "extension": "pdf",
  "lastModified": 1699999999999,
  "url": "/api/files/abc123",
  "thumbnailUrl": "/api/files/abc123/thumbnail",
  "downloadUrl": "/api/files/abc123/download",
  "uploadedAt": "2025-11-13T10:00:00Z",
  "uploadedBy": "user_123",
  "hash": "sha256:abc...",
  "virusScanStatus": "clean"
}
```

**Common MIME Types:**

```json
{
  "documents": [
    "application/pdf",
    "application/msword",
    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
    "application/vnd.ms-excel",
    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
    "application/vnd.ms-powerpoint",
    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
    "text/plain",
    "text/csv"
  ],
  "images": [
    "image/jpeg",
    "image/png",
    "image/gif",
    "image/webp",
    "image/svg+xml"
  ],
  "archives": [
    "application/zip",
    "application/x-rar-compressed",
    "application/x-7z-compressed",
    "application/x-tar"
  ]
}
```

**Security Best Practices:**
- Always validate file types on server-side
- Scan uploaded files for viruses and malware
- Store files outside web root directory
- Use random file names to prevent overwriting
- Implement file size limits
- Check file content, not just extension
- Use secure file storage (S3, cloud storage)
- Implement access control for downloads

---

### `image`

Image upload with preview, cropping, and editing capabilities.

**HTML Rendering:** `<input type="file" accept="image/*">`  
**Database Type:** `VARCHAR` (image path) or `JSONB` (metadata)  
**Validation:** Size, dimensions, format

```json
{
  "name": "avatar",
  "type": "image",
  "label": "Profile Picture",
  "config": {
    "multiple": false,
    "maxSize": 5242880,
    "accept": ["image/jpeg", "image/png", "image/webp", "image/gif"],
    "crop": true,
    "aspectRatio": 1,
    "cropShape": "circle",
    "resize": {
      "maxWidth": 800,
      "maxHeight": 800,
      "quality": 0.9,
      "format": "webp"
    },
    "preview": true,
    "showDimensions": true,
    "filters": ["brightness", "contrast", "saturation"],
    "uploadUrl": "/api/upload/image",
    "generateThumbnails": true,
    "thumbnailSizes": [
      {"width": 150, "height": 150, "name": "thumb"},
      {"width": 300, "height": 300, "name": "small"},
      {"width": 600, "height": 600, "name": "medium"}
    ],
    "allowedFormats": ["jpeg", "png", "webp"],
    "convertTo": "webp",
    "removeExif": true
  },
  "validation": {
    "maxSize": 5242880,
    "accept": ["image/jpeg", "image/png", "image/webp"],
    "dimensions": {
      "minWidth": 100,
      "minHeight": 100,
      "maxWidth": 4000,
      "maxHeight": 4000,
      "aspectRatio": 1
    }
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `crop` | boolean | false | Enable image cropping |
| `aspectRatio` | number | null | Crop aspect ratio (width/height) |
| `cropShape` | string | "rect" | Crop shape: `rect`, `circle` |
| `resize` | object | null | Auto-resize settings |
| `preview` | boolean | true | Show image preview |
| `showDimensions` | boolean | true | Display image dimensions |
| `filters` | array | [] | Available image filters |
| `generateThumbnails` | boolean | false | Auto-generate thumbnails |
| `thumbnailSizes` | array | [] | Thumbnail size configurations |
| `allowedFormats` | array | all | Allowed image formats |
| `convertTo` | string | null | Auto-convert format |
| `removeExif` | boolean | false | Strip EXIF metadata |

**Image Editing Configuration:**

```json
{
  "config": {
    "edit": true,
    "editTools": ["crop", "rotate", "flip", "filter", "adjust"],
    "filters": [
      "grayscale",
      "sepia",
      "vintage",
      "warm",
      "cool",
      "bright",
      "contrast"
    ],
    "adjustments": [
      "brightness",
      "contrast",
      "saturation",
      "hue",
      "blur",
      "sharpen"
    ],
    "rotateSteps": 90,
    "allowFreeRotate": false
  }
}
```

**Aspect Ratio Presets:**

```json
{
  "config": {
    "aspectRatioPresets": [
      {"label": "Square", "value": 1},
      {"label": "Portrait", "value": 0.75},
      {"label": "Landscape", "value": 1.33},
      {"label": "16:9", "value": 1.78},
      {"label": "4:3", "value": 1.33},
      {"label": "Free", "value": null}
    ]
  }
}
```

**Responsive Image Generation:**

```json
{
  "config": {
    "responsive": true,
    "breakpoints": [
      {"name": "mobile", "width": 480},
      {"name": "tablet", "width": 768},
      {"name": "desktop", "width": 1200},
      {"name": "full", "width": 1920}
    ],
    "format": "webp",
    "quality": 85
  }
}
```

---

### `video`

Video upload with player preview and thumbnail generation.

**HTML Rendering:** `<input type="file" accept="video/*">`  
**Database Type:** `VARCHAR` (video path) or `JSONB` (metadata)  
**Validation:** Size, duration, format

```json
{
  "name": "promo_video",
  "type": "video",
  "label": "Promotional Video",
  "config": {
    "maxSize": 104857600,
    "accept": ["video/mp4", "video/webm", "video/ogg"],
    "maxDuration": 300,
    "generateThumbnail": true,
    "thumbnailTime": 1,
    "thumbnailCount": 3,
    "showPlayer": true,
    "playerControls": true,
    "autoplay": false,
    "muted": false,
    "loop": false,
    "preload": "metadata",
    "showMetadata": true,
    "extractAudio": false,
    "transcodeFormats": ["mp4", "webm"]
  },
  "validation": {
    "maxSize": 104857600,
    "maxDuration": 300,
    "minDuration": 1,
    "accept": ["video/mp4", "video/webm"],
    "dimensions": {
      "minWidth": 640,
      "minHeight": 480,
      "maxWidth": 1920,
      "maxHeight": 1080
    }
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `maxDuration` | integer | null | Max video duration (seconds) |
| `generateThumbnail` | boolean | true | Auto-generate thumbnail |
| `thumbnailTime` | number | 1 | Time for thumbnail (seconds) |
| `thumbnailCount` | integer | 1 | Number of thumbnails |
| `showPlayer` | boolean | true | Show video player preview |
| `playerControls` | boolean | true | Show player controls |
| `autoplay` | boolean | false | Auto-play video |
| `muted` | boolean | false | Mute audio by default |
| `loop` | boolean | false | Loop video |
| `preload` | string | "metadata" | Preload: `none`, `metadata`, `auto` |
| `showMetadata` | boolean | true | Display video metadata |
| `extractAudio` | boolean | false | Extract audio track |
| `transcodeFormats` | array | [] | Auto-transcode to formats |

---

### `audio`

Audio file upload with player and waveform visualization.

**HTML Rendering:** `<input type="file" accept="audio/*">`  
**Database Type:** `VARCHAR` (audio path) or `JSONB` (metadata)  
**Validation:** Size, duration, format

```json
{
  "name": "voice_note",
  "type": "audio",
  "label": "Voice Note",
  "config": {
    "maxSize": 10485760,
    "accept": ["audio/mp3", "audio/wav", "audio/ogg", "audio/webm"],
    "maxDuration": 300,
    "showPlayer": true,
    "waveform": true,
    "waveformColor": "#4F46E5",
    "record": true,
    "recordFormat": "audio/webm",
    "maxRecordDuration": 60,
    "showDuration": true,
    "showFileSize": true,
    "visualizer": "waveform"
  },
  "validation": {
    "maxSize": 10485760,
    "maxDuration": 300,
    "minDuration": 1,
    "accept": ["audio/mp3", "audio/wav", "audio/webm"]
  }
}
```

**Audio Recording Configuration:**

```json
{
  "config": {
    "record": true,
    "maxRecordDuration": 120,
    "recordFormat": "audio/webm",
    "recordCodec": "opus",
    "sampleRate": 48000,
    "channelCount": 1,
    "audioBitsPerSecond": 128000,
    "showWaveform": true,
    "showTimer": true,
    "countdown": 3,
    "pauseRecording": true
  }
}
```

**Waveform Visualization:**

```json
{
  "config": {
    "waveform": true,
    "waveformColor": "#4F46E5",
    "waveformProgressColor": "#818CF8",
    "waveformHeight": 80,
    "waveformBarWidth": 2,
    "waveformBarGap": 1,
    "normalize": true,
    "interact": true
  }
}
```

---

## Text Content Fields

### `textarea`

Multi-line text input for longer text content.

**HTML Rendering:** `<textarea>`  
**Database Type:** `TEXT`  
**Validation:** Length, pattern

```json
{
  "name": "description",
  "type": "textarea",
  "label": "Description",
  "placeholder": "Enter a detailed description...",
  "config": {
    "rows": 4,
    "minRows": 2,
    "maxRows": 10,
    "maxLength": 1000,
    "showCount": true,
    "autoSize": true,
    "resize": "vertical",
    "spellcheck": true,
    "wordWrap": true,
    "lineNumbers": false,
    "highlightActiveLine": false
  },
  "validation": {
    "minLength": 10,
    "maxLength": 1000,
    "pattern": null,
    "message": "Description must be between 10 and 1000 characters"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `rows` | integer | 4 | Initial row count |
| `minRows` | integer | null | Minimum rows (autoSize) |
| `maxRows` | integer | null | Maximum rows (autoSize) |
| `maxLength` | integer | null | Character limit |
| `showCount` | boolean | false | Show character counter |
| `autoSize` | boolean | false | Auto-grow height |
| `resize` | string | "vertical" | Resize: `none`, `both`, `horizontal`, `vertical` |
| `spellcheck` | boolean | true | Enable spellcheck |
| `wordWrap` | boolean | true | Wrap long lines |
| `lineNumbers` | boolean | false | Show line numbers |
| `highlightActiveLine` | boolean | false | Highlight current line |

---

### `richtext`

WYSIWYG rich text editor with formatting tools.

**HTML Rendering:** Custom rich text editor component  
**Database Type:** `TEXT` or `JSONB`  
**Validation:** Length, allowed tags

```json
{
  "name": "content",
  "type": "richtext",
  "label": "Content",
  "config": {
    "toolbar": [
      "bold", "italic", "underline", "strike",
      "heading", "blockquote", "code", "codeBlock",
      "bulletList", "orderedList", "taskList",
      "link", "image", "video", "table",
      "horizontalRule", "undo", "redo",
      "align", "textColor", "backgroundColor"
    ],
    "toolbarSticky": true,
    "height": 400,
    "minHeight": 200,
    "maxHeight": 800,
    "maxLength": 50000,
    "allowImages": true,
    "imageUploadUrl": "/api/upload/image",
    "allowVideos": true,
    "videoEmbedProviders": ["youtube", "vimeo", "dailymotion"],
    "allowTables": true,
    "allowCodeBlocks": true,
    "codeLanguages": [
      "javascript",
      "python",
      "java",
      "sql",
      "html",
      "css",
      "php",
      "ruby"
    ],
    "placeholder": "Start writing...",
    "spellcheck": true,
    "autosave": true,
    "autosaveInterval": 30000,
    "characterCount": true,
    "wordCount": true,
    "readingTime": true
  },
  "validation": {
    "minLength": 50,
    "maxLength": 50000,
    "allowedTags": ["p", "h1", "h2", "h3", "strong", "em", "a", "img", "ul", "ol", "li"],
    "allowedAttributes": {
      "a": ["href", "title", "target"],
      "img": ["src", "alt", "width", "height"]
    }
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `toolbar` | array | all | Enabled toolbar buttons |
| `toolbarSticky` | boolean | false | Sticky toolbar on scroll |
| `height` | integer | null | Fixed editor height |
| `minHeight` | integer | null | Minimum editor height |
| `maxHeight` | integer | null | Maximum editor height |
| `maxLength` | integer | null | Character limit |
| `allowImages` | boolean | true | Allow image insertion |
| `imageUploadUrl` | string | null | Image upload endpoint |
| `allowVideos` | boolean | true | Allow video embeds |
| `videoEmbedProviders` | array | [] | Allowed video providers |
| `allowTables` | boolean | true | Allow table insertion |
| `allowCodeBlocks` | boolean | true | Allow code blocks |
| `codeLanguages` | array | [] | Syntax highlight languages |
| `autosave` | boolean | false | Auto-save content |
| `autosaveInterval` | integer | 30000 | Auto-save interval (ms) |
| `characterCount` | boolean | false | Show character count |
| `wordCount` | boolean | false | Show word count |
| `readingTime` | boolean | false | Show reading time |

**Toolbar Configuration:**

Available toolbar buttons:
- **Text formatting:** `bold`, `italic`, `underline`, `strike`, `code`
- **Headings:** `heading`, `h1`, `h2`, `h3`, `h4`, `h5`, `h6`
- **Lists:** `bulletList`, `orderedList`, `taskList`
- **Blocks:** `blockquote`, `codeBlock`, `horizontalRule`
- **Media:** `link`, `image`, `video`, `table`
- **Alignment:** `alignLeft`, `alignCenter`, `alignRight`, `alignJustify`
- **Colors:** `textColor`, `backgroundColor`, `highlight`
- **Actions:** `undo`, `redo`, `clear`
- **View:** `fullscreen`, `preview`, `source`

**Custom Toolbar Groups:**

```json
{
  "toolbar": [
    ["bold", "italic", "underline"],
    ["h1", "h2", "h3"],
    ["bulletList", "orderedList"],
    ["link", "image"],
    ["undo", "redo"]
  ]
}
```

---

### `code`

Code editor with syntax highlighting and linting.

**HTML Rendering:** Code editor component (Monaco, CodeMirror)  
**Database Type:** `TEXT`  
**Validation:** Syntax, length

```json
{
  "name": "source_code",
  "type": "code",
  "label": "Source Code",
  "config": {
    "language": "javascript",
    "theme": "vs-dark",
    "height": 400,
    "minLines": 10,
    "maxLines": 1000,
    "lineNumbers": true,
    "lineWrapping": false,
    "tabSize": 2,
    "indentWithTabs": false,
    "autoCloseBrackets": true,
    "matchBrackets": true,
    "lint": true,
    "autocomplete": true,
    "foldGutter": true,
    "highlightActiveLine": true,
    "highlightSelectionMatches": true,
    "readOnly": false,
    "placeholder": "// Enter code here..."
  },
  "validation": {
    "syntax": true,
    "maxLength": 100000,
    "language": "javascript"
  }
}
```

**Supported Languages:**
- `javascript` / `typescript`
- `python`
- `java`
- `csharp`
- `php`
- `ruby`
- `go`
- `rust`
- `sql`
- `html` / `css`
- `json` / `yaml` / `xml`
- `markdown`
- `shell` / `bash`

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `language` | string | "javascript" | Programming language |
| `theme` | string | "vs-dark" | Editor theme |
| `height` | integer | 400 | Editor height (px) |
| `minLines` | integer | null | Minimum lines |
| `maxLines` | integer | null | Maximum lines |
| `lineNumbers` | boolean | true | Show line numbers |
| `lineWrapping` | boolean | false | Wrap long lines |
| `tabSize` | integer | 2 | Tab size in spaces |
| `indentWithTabs` | boolean | false | Use tabs for indentation |
| `autoCloseBrackets` | boolean | true | Auto-close brackets |
| `matchBrackets` | boolean | true | Highlight matching brackets |
| `lint` | boolean | false | Enable linting |
| `autocomplete` | boolean | true | Enable autocomplete |
| `foldGutter` | boolean | true | Show code folding |
| `highlightActiveLine` | boolean | true | Highlight active line |
| `readOnly` | boolean | false | Read-only mode |

---

### `json`

JSON editor with validation and formatting.

**HTML Rendering:** JSON editor component  
**Database Type:** `JSONB` or `TEXT`  
**Validation:** Valid JSON syntax

```json
{
  "name": "config",
  "type": "json",
  "label": "Configuration",
  "config": {
    "height": 400,
    "theme": "dark",
    "mode": "tree",
    "modes": ["tree", "code", "form", "text"],
    "indentation": 2,
    "sortKeys": false,
    "expandAll": false,
    "search": true,
    "history": true,
    "schema": null,
    "schemaRefs": {},
    "templates": []
  },
  "validation": {
    "syntax": true,
    "schema": {
      "type": "object",
      "properties": {
        "name": {"type": "string"},
        "age": {"type": "number", "minimum": 0}
      },
      "required": ["name"]
    }
  }
}
```

**JSON Editor Modes:**
- `tree` - Interactive tree view
- `code` - Code editor with syntax highlighting
- `form` - Auto-generated form from schema
- `text` - Plain text editor
- `view` - Read-only view

---

### `markdown`

Markdown editor with live preview.

**HTML Rendering:** Markdown editor with preview pane  
**Database Type:** `TEXT`  
**Validation:** Length

```json
{
  "name": "readme",
  "type": "markdown",
  "label": "README",
  "config": {
    "height": 500,
    "showPreview": true,
    "previewPosition": "right",
    "syncScroll": true,
    "toolbar": true,
    "spellcheck": true,
    "allowHTML": false,
    "sanitize": true,
    "highlightCode": true,
    "lineNumbers": false,
    "wordWrap": true,
    "shortcuts": true
  },
  "validation": {
    "maxLength": 100000
  }
}
```

**Markdown Toolbar:**
- Bold, Italic, Strikethrough
- Headings (H1-H6)
- Lists (ordered, unordered, tasks)
- Links, Images
- Code blocks, Inline code
- Blockquotes
- Horizontal rules
- Tables

---

## Geographic & Location Fields

### `location`

Location picker with map interface and geocoding.

**HTML Rendering:** Map picker component  
**Database Type:** `POINT` (PostGIS) or `JSON`  
**Validation:** Coordinates, bounds

```json
{
  "name": "address_location",
  "type": "location",
  "label": "Location",
  "config": {
    "mapProvider": "google",
    "apiKey": "YOUR_API_KEY",
    "defaultZoom": 13,
    "defaultCenter": {
      "lat": -1.286389,
      "lng": 36.817223
    },
    "searchBox": true,
    "geocoding": true,
    "reverseGeocoding": true,
    "showMarker": true,
    "draggableMarker": true,
    "showAddress": true,
    "showCoordinates": true,
    "mapHeight": 400,
    "detectLocation": true,
    "bounds": null,
    "types": ["address", "establishment"],
    "country": null
  },
  "validation": {
    "required": true,
    "bounds": {
      "north": 5.0,
      "south": -5.0,
      "east": 42.0,
      "west": 33.0
    }
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `mapProvider` | string | "google" | Map provider: `google`, `mapbox`, `osm` |
| `apiKey` | string | null | API key for map provider |
| `defaultZoom` | integer | 13 | Initial zoom level |
| `defaultCenter` | object | null | Default map center coordinates |
| `searchBox` | boolean | true | Show address search box |
| `geocoding` | boolean | true | Enable address-to-coordinates |
| `reverseGeocoding` | boolean | true | Enable coordinates-to-address |
| `showMarker` | boolean | true | Show location marker |
| `draggableMarker` | boolean | true | Allow dragging marker |
| `showAddress` | boolean | true | Display formatted address |
| `showCoordinates` | boolean | false | Display lat/lng |
| `mapHeight` | integer | 400 | Map height (px) |
| `detectLocation` | boolean | false | Auto-detect user location |
| `bounds` | object | null | Restrict to geographic bounds |
| `types` | array | [] | Place types to search |
| `country` | string | null | Restrict to country (ISO code) |

**Location Data Format:**

```json
{
  "lat": -1.286389,
  "lng": 36.817223,
  "address": "123 Kenyatta Avenue, Nairobi, Kenya",
  "formattedAddress": "123 Kenyatta Ave, Nairobi, Kenya",
  "placeId": "ChIJ...",
  "components": {
    "street_number": "123",
    "route": "Kenyatta Avenue",
    "locality": "Nairobi",
    "administrative_area_level_1": "Nairobi County",
    "country": "Kenya",
    "postal_code": "00100"
  }
}
```

---

### `map`

Interactive map display for multiple locations.

**HTML Rendering:** Interactive map component  
**Database Type:** `JSON` or PostGIS  
**Validation:** Valid coordinates

```json
{
  "name": "store_locations",
  "type": "map",
  "label": "Store Locations",
  "config": {
    "mapProvider": "google",
    "apiKey": "YOUR_API_KEY",
    "height": 600,
    "zoom": 10,
    "center": {"lat": -1.286389, "lng": 36.817223},
    "markers": [],
    "clustering": true,
    "drawingTools": ["marker", "polygon", "circle"],
    "layers": ["traffic", "transit"],
    "styles": null,
    "controls": {
      "zoom": true,
      "mapType": true,
      "streetView": true,
      "fullscreen": true
    }
  }
}
```

---

## Biometric & Capture Fields

### `signature`

Digital signature capture pad.

**HTML Rendering:** Canvas-based signature pad  
**Database Type:** `TEXT` (base64) or `VARCHAR` (file path)  
**Validation:** Required, not empty

```json
{
  "name": "customer_signature",
  "type": "signature",
  "label": "Signature",
  "required": true,
  "config": {
    "width": 500,
    "height": 200,
    "penColor": "#000000",
    "backgroundColor": "#ffffff",
    "penWidth": 2,
    "minPenWidth": 1,
    "maxPenWidth": 3,
    "velocityFilterWeight": 0.7,
    "clear": true,
    "undo": true,
    "dotSize": 1,
    "throttle": 16,
    "minDistance": 5,
    "format": "image/png",
    "quality": 1,
    "trim": true
  },
  "validation": {
    "required": true,
    "notEmpty": true,
    "minPoints": 10
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `width` | integer | 500 | Canvas width (px) |
| `height` | integer | 200 | Canvas height (px) |
| `penColor` | string | "#000" | Pen stroke color |
| `backgroundColor` | string | "#fff" | Canvas background |
| `penWidth` | number | 2 | Base pen width |
| `minPenWidth` | number | 1 | Minimum pen width |
| `maxPenWidth` | number | 3 | Maximum pen width |
| `clear` | boolean | true | Show clear button |
| `undo` | boolean | true | Show undo button |
| `format` | string | "image/png" | Output format |
| `quality` | number | 1 | Image quality (0-1) |
| `trim` | boolean | false | Trim whitespace |

---

### `qrcode`

QR code scanner and generator.

**HTML Rendering:** QR code scanner/display  
**Database Type:** `TEXT`  
**Validation:** Valid QR data

```json
{
  "name": "qr_data",
  "type": "qrcode",
  "label": "QR Code",
  "config": {
    "mode": "both",
    "scanner": {
      "enabled": true,
      "camera": "environment",
      "fps": 10,
      "aspectRatio": 1,
      "constraints": {}
    },
    "generator": {
      "enabled": true,
      "size": 256,
      "margin": 4,
      "errorCorrectionLevel": "M",
      "format": "png",
      "color": {
        "dark": "#000000",
        "light": "#ffffff"
      },
      "logo": null,
      "logoSize": 0.2
    }
  }
}
```

**QR Code Scanner:**
- Front/back camera selection
- Real-time scanning
- Multiple format support (QR, barcode)
- Flashlight control

**QR Code Generator:**
- Customizable size and colors
- Error correction levels (L, M, Q, H)
- Logo/image overlay
- Multiple output formats

---

### `barcode`

Barcode scanner and generator.

**HTML Rendering:** Barcode scanner/display  
**Database Type:** `VARCHAR(255)`  
**Validation:** Valid barcode format

```json
{
  "name": "product_barcode",
  "type": "barcode",
  "label": "Product Barcode",
  "config": {
    "mode": "both",
    "formats": [
      "EAN13",
      "EAN8",
      "UPC",
      "CODE39",
      "CODE128",
      "ITF",
      "CODABAR"
    ],
    "scanner": {
      "enabled": true,
      "camera": "environment",
      "continuous": false,
      "playSound": true
    },
    "generator": {
      "enabled": true,
      "format": "CODE128",
      "width": 2,
      "height": 100,
      "displayValue": true,
      "fontSize": 14,
      "margin": 10,
      "background": "#ffffff",
      "lineColor": "#000000"
    }
  }
}
```

**Supported Barcode Formats:**
- EAN-13, EAN-8
- UPC-A, UPC-E
- CODE39, CODE128
- ITF (Interleaved 2 of 5)
- CODABAR
- QR Code
- Data Matrix
- PDF417

---

## Specialized Input Fields

### `tags`

Tag input for multiple keyword entries.

**HTML Rendering:** Tag input component  
**Database Type:** `TEXT[]` or `JSON`  
**Validation:** Max tags, pattern

```json
{
  "name": "keywords",
  "type": "tags",
  "label": "Keywords",
  "placeholder": "Add keywords...",
  "config": {
    "maxTags": 10,
    "minLength": 2,
    "maxLength": 50,
    "duplicates": false,
    "caseSensitive": false,
    "suggestions": [],
    "suggestionsSource": "/api/tags/suggestions",
    "createOnBlur": true,
    "createOnEnter": true,
    "createOnComma": true,
    "createOnSpace": false,
    "trim": true,
    "lowercase": false,
    "pattern": null,
    "tagColor": "primary",
    "closeable": true,
    "draggable": false
  },
  "validation": {
    "maxTags": 10,
    "minTags": 1,
    "pattern": "^[a-zA-Z0-9-]+$"
  }
}
```

---

### `rating`

Star rating or custom rating input.

**HTML Rendering:** Interactive rating component  
**Database Type:** `NUMERIC(3,2)` or `SMALLINT`  
**Validation:** Min/max rating

```json
{
  "name": "product_rating",
  "type": "rating",
  "label": "Rating",
  "config": {
    "max": 5,
    "precision": 0.5,
    "icon": "star",
    "emptyIcon": "star_outline",
    "size": "large",
    "color": "warning",
    "allowHalf": true,
    "allowClear": false,
    "showValue": true,
    "valuePosition": "after",
    "readonly": false,
    "disabled": false,
    "highlightSelected": true
  },
  "validation": {
    "min": 1,
    "max": 5,
    "required": true
  },
  "default": 0
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `max` | integer | 5 | Maximum rating value |
| `precision` | number | 1 | Rating precision (0.5 for half stars) |
| `icon` | string | "star" | Icon name |
| `emptyIcon` | string | "star_outline" | Empty state icon |
| `size` | string | "medium" | Size: `small`, `medium`, `large` |
| `color` | string | "warning" | Color theme |
| `allowHalf` | boolean | false | Allow half values |
| `allowClear` | boolean | false | Allow clearing rating |
| `showValue` | boolean | false | Display numeric value |
| `valuePosition` | string | "after" | Value position: `before`, `after` |
| `readonly` | boolean | false | Read-only mode |
| `highlightSelected` | boolean | true | Highlight on hover |

**Custom Icons:**

```json
{
  "config": {
    "icon": "❤️",
    "emptyIcon": "🤍",
    "max": 5
  }
}
```

---

### `slider`

Slider input for numeric range selection.

**HTML Rendering:** Range slider component  
**Database Type:** `INTEGER` or `NUMERIC`  
**Validation:** Min/max range

```json
{
  "name": "price_range",
  "type": "slider",
  "label": "Price Range",
  "config": {
    "min": 0,
    "max": 1000,
    "step": 10,
    "range": true,
    "marks": {
      "0": "$0",
      "250": "$250",
      "500": "$500",
      "750": "$750",
      "1000": "$1000"
    },
    "showMarks": true,
    "showValue": true,
    "showInput": true,
    "vertical": false,
    "height": null,
    "tooltipVisible": "auto",
    "tooltipFormat": "${value}",
    "color": "primary"
  },
  "validation": {
    "min": 0,
    "max": 1000
  },
  "default": [100, 500]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `min` | number | 0 | Minimum value |
| `max` | number | 100 | Maximum value |
| `step` | number | 1 | Step increment |
| `range` | boolean | false | Range mode (two handles) |
| `marks` | object | {} | Mark labels at positions |
| `showMarks` | boolean | false | Display marks |
| `showValue` | boolean | true | Display current value |
| `showInput` | boolean | false | Show numeric input |
| `vertical` | boolean | false | Vertical orientation |
| `height` | integer | null | Height for vertical slider |
| `tooltipVisible` | string | "auto" | Tooltip: `always`, `auto`, `never` |
| `tooltipFormat` | string | null | Tooltip format template |
| `color` | string | "primary" | Color theme |

**Range Slider Example:**

```json
{
  "config": {
    "range": true,
    "min": 0,
    "max": 100,
    "step": 5,
    "default": [25, 75],
    "tooltipFormat": "{value}%"
  }
}
```

---

### `color`

Color picker with multiple formats.

**HTML Rendering:** Color picker component  
**Database Type:** `VARCHAR(50)`  
**Validation:** Color format

```json
{
  "name": "brand_color",
  "type": "color",
  "label": "Brand Color",
  "config": {
    "format": "hex",
    "showAlpha": false,
    "showPreview": true,
    "showInput": true,
    "presets": [
      "#FF0000",
      "#00FF00",
      "#0000FF",
      "#FFFF00",
      "#FF00FF",
      "#00FFFF"
    ],
    "presetColors": [
      {"color": "#FF0000", "label": "Red"},
      {"color": "#00FF00", "label": "Green"},
      {"color": "#0000FF", "label": "Blue"}
    ],
    "disableAlpha": false,
    "showEyeDropper": true,
    "showGradient": false
  },
  "validation": {
    "format": "hex",
    "allowAlpha": false
  },
  "default": "#4F46E5"
}
```

**Color Formats:**
- `hex` - #RRGGBB or #RRGGBBAA
- `rgb` - rgb(r, g, b) or rgba(r, g, b, a)
- `hsl` - hsl(h, s%, l%) or hsla(h, s%, l%, a)
- `hsv` - hsv(h, s%, v%)

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `format` | string | "hex" | Output color format |
| `showAlpha` | boolean | false | Show alpha/opacity control |
| `showPreview` | boolean | true | Show color preview |
| `showInput` | boolean | true | Show color input field |
| `presets` | array | [] | Preset color array |
| `presetColors` | array | [] | Preset colors with labels |
| `disableAlpha` | boolean | false | Disable alpha channel |
| `showEyeDropper` | boolean | false | Show eyedropper tool |
| `showGradient` | boolean | false | Show gradient picker |

---

### `icon-picker`

Icon selector from icon libraries.

**HTML Rendering:** Icon picker component  
**Database Type:** `VARCHAR(100)`  
**Validation:** Valid icon name

```json
{
  "name": "menu_icon",
  "type": "icon-picker",
  "label": "Menu Icon",
  "config": {
    "library": "lucide",
    "libraries": ["lucide", "heroicons", "fontawesome"],
    "searchable": true,
    "categories": true,
    "previewSize": "medium",
    "gridSize": 6,
    "showName": true,
    "allowCustom": false,
    "customIconUrl": null,
    "color": "#000000",
    "size": 24
  },
  "validation": {
    "required": true
  }
}
```

**Supported Icon Libraries:**
- Lucide Icons
- Heroicons
- Font Awesome
- Material Design Icons
- Bootstrap Icons
- Feather Icons

---

## Relational Fields

### `relation`

Reference to records in another table.

**HTML Rendering:** Select/autocomplete component  
**Database Type:** Foreign key (INTEGER, UUID)  
**Validation:** Record exists

```json
{
  "name": "user_id",
  "type": "relation",
  "label": "Assigned User",
  "required": true,
  "config": {
    "table": "users",
    "valueField": "id",
    "labelField": "name",
    "searchFields": ["name", "email"],
    "displayFields": ["name", "email", "avatar"],
    "displayTemplate": "{name} ({email})",
    "filters": {
      "status": "active"
    },
    "orderBy": "name",
    "orderDirection": "asc",
    "limit": 100,
    "searchable": true,
    "clearable": true,
    "multiple": false,
    "createNew": false,
    "createNewUrl": "/users/create",
    "cache": true,
    "cacheTimeout": 300000,
    "dependsOn": [],
    "cascadeDelete": false
  },
  "validation": {
    "exists": true,
    "table": "users",
    "field": "id"
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `table` | string | required | Related table name |
| `valueField` | string | "id" | Field to use as value |
| `labelField` | string | "name" | Field to display as label |
| `searchFields` | array | [] | Fields to search |
| `displayFields` | array | [] | Fields to display in dropdown |
| `displayTemplate` | string | null | Custom display template |
| `filters` | object | {} | Additional filters |
| `orderBy` | string | null | Sort field |
| `orderDirection` | string | "asc" | Sort direction |
| `limit` | integer | 100 | Max records to load |
| `searchable` | boolean | true | Enable search |
| `clearable` | boolean | true | Allow clearing selection |
| `multiple` | boolean | false | Allow multiple selections |
| `createNew` | boolean | false | Allow creating new records |
| `createNewUrl` | string | null | URL for creating new record |
| `cache` | boolean | true | Cache results |
| `cacheTimeout` | integer | 300000 | Cache timeout (ms) |
| `dependsOn` | array | [] | Dependent fields |
| `cascadeDelete` | boolean | false | Delete when parent deleted |

**Many-to-Many Relationships:**

```json
{
  "name": "tags",
  "type": "relation",
  "config": {
    "table": "tags",
    "multiple": true,
    "pivotTable": "post_tags",
    "pivotFields": {
      "post_id": "id",
      "tag_id": "tag_id"
    }
  }
}
```

---

### `autocomplete`

Autocomplete input with async data loading.

**HTML Rendering:** Autocomplete input  
**Database Type:** Varies  
**Validation:** Varies

```json
{
  "name": "customer",
  "type": "autocomplete",
  "label": "Customer",
  "placeholder": "Search customers...",
  "config": {
    "source": "/api/customers/search",
    "method": "GET",
    "params": {},
    "valueField": "id",
    "labelField": "name",
    "minLength": 2,
    "debounce": 300,
    "limit": 10,
    "cache": true,
    "cacheTimeout": 300000,
    "showNoResults": true,
    "noResultsText": "No customers found",
    "loadingText": "Searching...",
    "clearable": true,
    "freeSolo": false,
    "autoSelect": true,
    "selectOnFocus": false
  }
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `source` | string/function | required | Data source URL or function |
| `method` | string | "GET" | HTTP method |
| `params` | object | {} | Additional query params |
| `valueField` | string | "id" | Value field name |
| `labelField` | string | "name" | Label field name |
| `minLength` | integer | 2 | Min chars to trigger search |
| `debounce` | integer | 300 | Debounce delay (ms) |
| `limit` | integer | 10 | Max results to show |
| `cache` | boolean | true | Cache results |
| `cacheTimeout` | integer | 300000 | Cache timeout (ms) |
| `showNoResults` | boolean | true | Show no results message |
| `noResultsText` | string | default | No results message |
| `loadingText` | string | default | Loading message |
| `clearable` | boolean | true | Show clear button |
| `freeSolo` | boolean | false | Allow arbitrary input |
| `autoSelect` | boolean | false | Auto-select first match |
| `selectOnFocus` | boolean | false | Select text on focus |

---

### `repeatable`

Repeatable field group for dynamic lists.

**HTML Rendering:** Repeatable field group  
**Database Type:** `JSONB` or separate table  
**Validation:** Min/max items

```json
{
  "name": "phone_numbers",
  "type": "repeatable",
  "label": "Phone Numbers",
  "config": {
    "minItems": 1,
    "maxItems": 5,
    "addText": "Add Phone Number",
    "removeText": "Remove",
    "sortable": true,
    "collapsible": false,
    "defaultExpanded": true,
    "itemLabel": "Phone {index}",
    "confirmRemove": true,
    "duplicateButton": true,
    "virtualize": false
  },
  "fields": [
    {
      "name": "type",
      "type": "select",
      "label": "Type",
      "options": [
        {"value": "mobile", "label": "Mobile"},
        {"value": "home", "label": "Home"},
        {"value": "work", "label": "Work"}
      ],
      "default": "mobile"
    },
    {
      "name": "number",
      "type": "phone",
      "label": "Number",
      "required": true
    },
    {
      "name": "primary",
      "type": "checkbox",
      "label": "Primary"
    }
  ],
  "validation": {
    "minItems": 1,
    "maxItems": 5
  },
  "default": [
    {"type": "mobile", "number": "", "primary": true}
  ]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `minItems` | integer | 0 | Minimum number of items |
| `maxItems` | integer | null | Maximum number of items |
| `addText` | string | "Add Item" | Add button text |
| `removeText` | string | "Remove" | Remove button text |
| `sortable` | boolean | false | Allow drag-to-reorder |
| `collapsible` | boolean | false | Make items collapsible |
| `defaultExpanded` | boolean | true | Default expanded state |
| `itemLabel` | string | "Item {index}" | Item label template |
| `confirmRemove` | boolean | false | Confirm before removing |
| `duplicateButton` | boolean | false | Show duplicate button |
| `virtualize` | boolean | false | Use virtual scrolling |

---

## Layout & Container Fields

### `group`

Visual grouping of fields.

**HTML Rendering:** Fieldset or div container  
**Database Type:** N/A (container only)  
**Validation:** N/A

```json
{
  "name": "personal_info",
  "type": "group",
  "label": "Personal Information",
  "description": "Enter your personal details",
  "config": {
    "collapsible": false,
    "collapsed": false,
    "border": true,
    "padding": "normal",
    "columns": 1,
    "gap": "normal"
  },
  "fields": [
    {
      "name": "first_name",
      "type": "text",
      "label": "First Name",
      "required": true
    },
    {
      "name": "last_name",
      "type": "text",
      "label": "Last Name",
      "required": true
    },
    {
      "name": "date_of_birth",
      "type": "date",
      "label": "Date of Birth"
    }
  ]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `collapsible` | boolean | false | Can be collapsed |
| `collapsed` | boolean | false | Initially collapsed |
| `border` | boolean | true | Show border |
| `padding` | string | "normal" | Padding: `none`, `small`, `normal`, `large` |
| `columns` | integer | 1 | Number of columns |
| `gap` | string | "normal" | Gap between fields |

---

### `tabs`

Tabbed container for organizing fields.

**HTML Rendering:** Tab component  
**Database Type:** N/A (container only)  
**Validation:** N/A

```json
{
  "name": "product_tabs",
  "type": "tabs",
  "config": {
    "position": "top",
    "variant": "line",
    "keepAlive": true,
    "lazy": false
  },
  "tabs": [
    {
      "key": "basic",
      "label": "Basic Info",
      "icon": "info",
      "fields": [
        {
          "name": "name",
          "type": "text",
          "label": "Product Name",
          "required": true
        },
        {
          "name": "description",
          "type": "textarea",
          "label": "Description"
        }
      ]
    },
    {
      "key": "pricing",
      "label": "Pricing",
      "icon": "dollar",
      "fields": [
        {
          "name": "price",
          "type": "currency",
          "label": "Price",
          "required": true
        },
        {
          "name": "cost",
          "type": "currency",
          "label": "Cost"
        }
      ]
    },
    {
      "key": "inventory",
      "label": "Inventory",
      "icon": "package",
      "disabled": false,
      "fields": [
        {
          "name": "sku",
          "type": "text",
          "label": "SKU"
        },
        {
          "name": "stock",
          "type": "number",
          "label": "Stock Quantity"
        }
      ]
    }
  ]
}
```

**Configuration Options:**

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `position` | string | "top" | Tab position: `top`, `left`, `right`, `bottom` |
| `variant` | string | "line" | Variant: `line`, `card`, `pills` |
| `keepAlive` | boolean | true | Keep tab content mounted |
| `lazy` | boolean | false | Lazy load tab content |

---

### `accordion`

Accordion/collapsible panels.

**HTML Rendering:** Accordion component  
**Database Type:** N/A (container only)  
**Validation:** N/A

```json
{
  "name": "faq_sections",
  "type": "accordion",
  "config": {
    "multiple": false,
    "collapsible": true,
    "defaultExpanded": [0],
    "expandIcon": "chevron-down",
    "variant": "default"
  },
  "panels": [
    {
      "key": "general",
      "title": "General Information",
      "icon": "info",
      "fields": [...]
    },
    {
      "key": "address",
      "title": "Address Details",
      "icon": "map-pin",
      "fields": [...]
    }
  ]
}
```

---

### `fieldset`

HTML fieldset for semantic grouping.

**HTML Rendering:** `<fieldset>` element  
**Database Type:** N/A (container only)  
**Validation:** N/A

```json
{
  "name": "contact_fieldset",
  "type": "fieldset",
  "legend": "Contact Information",
  "description": "How can we reach you?",
  "config": {
    "border": true,
    "disabled": false
  },
  "fields": [
    {
      "name": "email",
      "type": "email",
      "label": "Email"
    },
    {
      "name": "phone",
      "type": "phone",
      "label": "Phone"
    }
  ]
}
```

---

## Display-Only Fields

### `static`

Display static text or HTML content.

**HTML Rendering:** Static content  
**Database Type:** N/A  
**Validation:** N/A

```json
{
  "name": "terms_notice",
  "type": "static",
  "content": "By submitting this form, you agree to our <a href='/terms'>Terms of Service</a>.",
  "config": {
    "variant": "info",
    "icon": "info-circle",
    "dismissible": false,
    "markdown": false,
    "html": true
  }
}
```

**Variants:**
- `default` - Plain text
- `info` - Info message
- `success` - Success message
- `warning` - Warning message
- `error` - Error message

---

### `divider`

Visual separator between fields.

**HTML Rendering:** `<hr>` or styled divider  
**Database Type:** N/A  
**Validation:** N/A

```json
{
  "name": "section_divider",
  "type": "divider",
  "label": "Additional Information",
  "config": {
    "variant": "solid",
    "thickness": 1,
    "color": "gray",
    "spacing": "normal",
    "orientation": "horizontal"
  }
}
```

---

### `spacer`

Empty space for layout.

**HTML Rendering:** Empty div with height  
**Database Type:** N/A  
**Validation:** N/A

```json
{
  "name": "spacer_1",
  "type": "spacer",
  "config": {
    "height": 32
  }
}
```

---

### `formula`

Calculated field based on other field values.

**HTML Rendering:** Display component  
**Database Type:** Virtual or computed column  
**Validation:** N/A (computed)

```json
{
  "name": "total_price",
  "type": "formula",
  "label": "Total Price",
  "config": {
    "formula": "quantity * unit_price",
    "format": "currency",
    "currency": "USD",
    "readonly": true,
    "recalculateOn": ["quantity", "unit_price"]
  }
}
```

**Formula Syntax:**

```javascript
// Arithmetic
"price * quantity"
"(subtotal + tax) * discount"

// Functions
"SUM(line_items.total)"
"AVERAGE(ratings)"
"MAX(prices)"
"MIN(dates)"
"COUNT(items)"

// Conditionals
"IF(quantity > 10, price * 0.9, price)"
"IF(status == 'active', 1, 0)"

// String operations
"CONCAT(first_name, ' ', last_name)"
"UPPER(email)"
"LOWER(username)"

// Date operations
"DATEDIFF(end_date, start_date)"
"DATEADD(start_date, 30, 'days')"
"NOW()"
"TODAY()"
```

---

### `chart`

Data visualization chart.

**HTML Rendering:** Chart component  
**Database Type:** N/A (display only)  
**Validation:** N/A

```json
{
  "name": "sales_chart",
  "type": "chart",
  "label": "Sales Overview",
  "config": {
    "chartType": "line",
    "dataSource": "/api/sales/chart-data",
    "height": 300,
    "options": {
      "responsive": true,
      "maintainAspectRatio": false,
      "plugins": {
        "legend": {
          "display": true,
          "position": "top"
        },
        "title": {
          "display": true,
          "text": "Monthly Sales"
        }
      }
    },
    "refreshInterval": 0
  }
}
```

**Chart Types:**
- `line` - Line chart
- `bar` - Bar chart
- `pie` - Pie chart
- `doughnut` - Doughnut chart
- `radar` - Radar chart
- `polarArea` - Polar area chart
- `scatter` - Scatter plot
- `bubble` - Bubble chart

---

## Database Type Mapping

Recommended database types for each field type:

| Field Type | PostgreSQL | MySQL | SQLite | MongoDB |
|------------|-----------|-------|--------|---------|
| `text` | VARCHAR/TEXT | VARCHAR/TEXT | TEXT | String |
| `email` | VARCHAR(255) | VARCHAR(255) | TEXT | String |
| `password` | VARCHAR(255) | VARCHAR(255) | TEXT | String |
| `url` | VARCHAR(2048) | VARCHAR(2048) | TEXT | String |
| `phone` | VARCHAR(50) | VARCHAR(50) | TEXT | String |
| `number` | INTEGER/NUMERIC | INT/DECIMAL | INTEGER/REAL | Number |
| `currency` | NUMERIC(19,4) | DECIMAL(19,4) | REAL | Number |
| `percentage` | NUMERIC(5,2) | DECIMAL(5,2) | REAL | Number |
| `date` | DATE | DATE | TEXT | Date |
| `time` | TIME | TIME | TEXT | String |
| `datetime` | TIMESTAMP | DATETIME | TEXT | Date |
| `daterange` | DATERANGE | N/A | TEXT | Object |
| `select` | VARCHAR/INTEGER | VARCHAR/INT | TEXT/INTEGER | String/Number |
| `checkbox` | BOOLEAN | TINYINT(1) | INTEGER | Boolean |
| `checkboxes` | TEXT[] | JSON | TEXT | Array |
| `file` | VARCHAR | VARCHAR | TEXT | String |
| `image` | VARCHAR | VARCHAR | TEXT | String |
| `textarea` | TEXT | TEXT | TEXT | String |
| `richtext` | TEXT | TEXT | TEXT | String |
| `code` | TEXT | TEXT | TEXT | String |
| `json` | JSONB | JSON | TEXT | Object |
| `tags` | TEXT[] | JSON | TEXT | Array |
| `rating` | NUMERIC(3,2) | DECIMAL(3,2) | REAL | Number |
| `color` | VARCHAR(50) | VARCHAR(50) | TEXT | String |
| `location` | POINT | POINT | TEXT | GeoJSON |
| `signature` | TEXT | TEXT | TEXT | String |
| `ipaddress` | INET | VARCHAR(45) | TEXT | String |
| `mac-address` | MACADDR | VARCHAR(17) | TEXT | String |
| `duration` | INTERVAL | INTEGER | INTEGER | Number |
| `relation` | INTEGER/UUID | INT/VARCHAR | INTEGER/TEXT | ObjectId |

---

## Validation Rules Reference

### Common Validation Rules

```json
{
  "validation": {
    "required": true,
    "message": "Custom error message",
    
    // String validations
    "minLength": 3,
    "maxLength": 100,
    "pattern": "^[a-zA-Z]+$",
    "email": true,
    "url": true,
    "alpha": true,
    "alphanumeric": true,
    "numeric": true,
    
    // Number validations
    "min": 0,
    "max": 100,
    "integer": true,
    "positive": true,
    "negative": false,
    "multipleOf": 5,
    "precision": 2,
    
    // Array validations
    "minItems": 1,
    "maxItems": 10,
    "uniqueItems": true,
    
    // Date validations
    "minDate": "2020-01-01",
    "maxDate": "today",
    "before": "end_date",
    "after": "start_date",
    "businessDays": true,
    
    // File validations
    "maxSize": 5242880,
    "accept": ["image/jpeg", "image/png"],
    "dimensions": {
      "minWidth": 100,
      "maxHeight": 2000
    },
    
    // Custom validations
    "custom": "validateCustomFunction",
    "async": "/api/validate/field",
    "asyncDebounce": 500,
    
    // Comparison validations
    "equals": "other_field",
    "notEquals": "other_field",
    "greaterThan": "other_field",
    "lessThan": "other_field",
    "matches": "password",
    
    // Conditional validations
    "requiredIf": {
      "field": "type",
      "value": "custom"
    },
    "requiredUnless": {
      "field": "status",
      "value": "draft"
    }
  }
}
```

### Custom Validation Functions

```javascript
// Client-side validation
function validateCustomFunction(value, field, formData) {
  if (!value) return true;
  
  // Custom validation logic
  if (value.length < 5) {
    return "Value must be at least 5 characters";
  }
  
  return true; // Valid
}

// Async validation
async function validateAsync(value, field, formData) {
  const response = await fetch(`/api/validate/${field.name}`, {
    method: 'POST',
    body: JSON.stringify({ value })
  });
  
  const result = await response.json();
  return result.valid ? true : result.message;
}
```

---

## Security Guidelines

### Input Sanitization

1. **Never Trust User Input**
   - Always validate and sanitize on server-side
   - Client-side validation is for UX only

2. **XSS Prevention**
   ```json
   {
     "config": {
       "sanitize": true,
       "allowedTags": ["p", "strong", "em"],
       "stripScripts": true
     }
   }
   ```

3. **SQL Injection Prevention**
   - Use parameterized queries
   - Never concatenate SQL with user input
   - Use ORM with proper escaping

4. **File Upload Security**
   - Validate file types by content, not extension
   - Scan for malware
   - Store outside web root
   - Use random filenames
   - Implement size limits
   - Check MIME types

5. **CSRF Protection**
   - Include CSRF tokens in forms
   - Validate tokens on server
   - Use SameSite cookies

### Access Control

```json
{
  "permissions": {
    "view": ["user", "admin"],
    "edit": ["admin"],
    "delete": ["admin"]
  }
}
```

### Data Encryption

- Encrypt sensitive fields at rest
- Use HTTPS for transmission
- Hash passwords with bcrypt/argon2
- Encrypt PII data
- Use environment variables for secrets

---

## Performance Optimization

### Field-Level Optimizations

1. **Virtual Scrolling** for large lists
   ```json
   {
     "config": {
       "virtualScroll": true,
       "itemHeight": 40
     }
   }
   ```

2. **Lazy Loading** for heavy components
   ```json
   {
     "config": {
       "lazy": true,
       "lazyPlaceholder": "Loading..."
     }
   }
   ```

3. **Debouncing** for async operations
   ```json
   {
     "config": {
       "debounce": 300,
       "asyncDebounce": 500
     }
   }
   ```

4. **Caching** for API calls
   ```json
   {
     "config": {
       "cache": true,
       "cacheTimeout": 300000
     }
   }
   ```

[← Back to Schema Structure](04-schema-structure.md) | [Next: Validation →](06-validation.md)
