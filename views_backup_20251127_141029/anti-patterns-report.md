# Component Anti-Pattern Detection Report

## Executive Summary

Found **4 violations** of the ClassName anti-pattern across the codebase. This breaks component encapsulation and violates the established component architecture guidelines.

---

## ClassName String Anti-Pattern Violations

### 🔴 CRITICAL VIOLATIONS

#### 1. BaseProps.ClassName
**File:** `./views/components/types.go:109`
```go
type BaseProps struct {
    ID        string             // HTML id attribute
    ClassName string             // ← ANTI-PATTERN
    Style     string             // Inline styles
    // ...
}
```
**Severity:** HIGH - Shared struct embedded in multiple components

#### 2. Icon.ClassName  
**File:** `./views/components/types.go:369`
```go
type Icon struct {
    Name      string // Icon name/identifier
    Size      Size   // Icon size
    ClassName string // ← ANTI-PATTERN
}
```
**Severity:** HIGH - Used in shared type definitions

#### 3. ButtonProps.ClassName (Demo)
**File:** `./views/demo/atoms.go:45`
```go
type ButtonProps struct {
    // ... other fields
    ClassName string  // ← ANTI-PATTERN
    // ...
}
```
**Severity:** MEDIUM - Only affects demo code

### 🟡 COMPILATION ERRORS FROM USAGE

#### 4. Navigation Template
**File:** `./views/components/organisms/navigation.templ`
- Multiple instances attempting to pass `ClassName` to Icon components
- **ERROR:** IconProps doesn't have ClassName field

#### 5. Form Template  
**File:** `./views/components/organisms/form.templ`
- Multiple instances attempting to pass `ClassName` to Icon/Button components
- **ERROR:** Target components missing ClassName fields

---

## Architecture Compliance

### ✅ PROPER PATTERNS FOUND

These components follow the correct architecture:

1. **Card Component** - Uses semantic props like `Border`, `Compact`, `Elevated`
2. **Button (main)** - Type-safe variants with computed classes
3. **Modal** - Proper slot-based composition

### ❌ VIOLATIONS BREAKDOWN

| Component | File | Line | Severity | Status |
|-----------|------|------|----------|---------|
| BaseProps | types.go | 109 | HIGH | Active |
| Icon | types.go | 369 | HIGH | Active |
| ButtonProps (demo) | demo/atoms.go | 45 | MEDIUM | Active |
| Navigation usage | navigation.templ | Multiple | CRITICAL | Compilation Error |
| Form usage | form.templ | Multiple | CRITICAL | Compilation Error |

---

## Recommended Actions

### IMMEDIATE (Fix Compilation)

1. **Remove ClassName from templates:**
   ```diff
   - ClassName: "text-muted-foreground"
   + // Use semantic props or remove entirely
   ```

### HIGH PRIORITY

1. **Remove BaseProps.ClassName:**
   ```go
   type BaseProps struct {
       ID        string             // HTML id attribute
   -   ClassName string             // REMOVE
       Style     string             // Inline styles
       // ...
   }
   ```

2. **Remove Icon.ClassName:**
   ```go
   type Icon struct {
       Name      string // Icon name/identifier  
       Size      Size   // Icon size
   -   ClassName string // REMOVE
   }
   ```

### MEDIUM PRIORITY

1. **Update demo ButtonProps** to match main component architecture

### ARCHITECTURAL GUIDANCE

Replace `ClassName string` with semantic boolean props:

```go
// ❌ ANTI-PATTERN
type ComponentProps struct {
    ClassName string
}

// ✅ CORRECT PATTERN  
type ComponentProps struct {
    Compact   bool
    Bordered  bool
    Elevated  bool
    Variant   ComponentVariant
}
```

Use class composition functions:
```go
func getComponentClasses(compact bool, bordered bool) string {
    base := "component-base"
    classes := []string{base}
    
    if compact {
        classes = append(classes, "py-2")
    }
    if bordered {
        classes = append(classes, "border border-border")
    }
    
    return strings.Join(classes, " ")
}
```

---

## Impact Assessment

- **Compilation Errors:** 2 templates currently broken
- **Architecture Violations:** 3 critical props structs  
- **Maintenance Risk:** HIGH - Leaky abstraction allows arbitrary style injection
- **Consistency:** Components inconsistent between proper/anti-pattern approaches

**Status:** Codebase in mixed state - some components follow proper patterns, others contain violations.