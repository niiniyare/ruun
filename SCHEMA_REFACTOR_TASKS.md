# Schema Package Refactoring - Task Tracking

## Overview

This document tracks the consolidation of redundant types and interfaces in the schema package.

## Current State Analysis

### Duplicate Types Identified

| Type | Locations | Action |
|------|-----------|--------|
| `Event` | types.go:129, interface.go:99 | Consolidate to event.go |
| `BehaviorMetadata` | types.go:140, interface.go:110 | Consolidate to event.go |
| `ValidationResult` | interface.go:120, validator.go:68 (commented) | Keep in types.go |
| `SwapStrategy` | types.go, behavior.go | Keep in behavior.go only |
| `Events` | types.go:150 | Move to event.go |

### Redundant Type Families

| Old Types | New Unified Type | Location |
|-----------|------------------|----------|
| `FieldConditional`, `LayoutConditional`, `ActionConditional`, `Conditional` | `Conditional` | conditional.go |
| `BaseStyle`, `LayoutStyle`, `SectionStyle`, `TabStyle`, `StepStyle`, `ActionTheme`, `FieldStyle` | `Style` | style.go |
| `HTMX`, `ActionHTMX`, `FieldHTMX` | `Behavior` | behavior.go |
| `Alpine`, `ActionAlpine`, `FieldAlpine` | `Binding` | binding.go |
| `ActionConfig` | `map[string]any` | (inline) |
| `ActionPermissions` | `[]string` | (inline) |

### Interface Consolidation

| Old Interfaces | New Interface | Location |
|----------------|---------------|----------|
| `SchemaValidator` (doc.go), `Validator` (interface.go) | `Validator` | interface.go |
| `SchemaRenderer` (doc.go), `Renderer` (interface.go) | `Renderer` | interface.go |
| `SchemaRegistry` (doc.go), `Registry` | `Registry` | interface.go |
| `FieldInterface`, `EnrichedFieldInterface`, `SchemaInterface`, `EnrichedSchemaInterface` | Use concrete types or `FieldAccessor`, `SchemaAccessor` | interface.go |

---

## Task Checklist

### Phase 1: Create Unified Type Files

- [ ] **T1.1** Create `types.go` - Common types (enums, configs, workflow, validation)
- [ ] **T1.2** Create `interface.go` - All package interfaces
- [ ] **T1.3** Create `conditional.go` - Unified `Conditional` type
- [ ] **T1.4** Create `style.go` - Unified `Style` type
- [ ] **T1.5** Create `event.go` - `Event`, `Events`, `BehaviorMetadata`
- [ ] **T1.6** Create `behavior.go` - `Behavior`, `Trigger`, `SwapStrategy`
- [ ] **T1.7** Create `binding.go` - `Binding` type
- [ ] **T1.8** Update `errors.go` - Consolidate error types

### Phase 2: Update Core Structs

- [ ] **T2.1** Update `field.go` - Use `*Conditional`, `*Style`, `*Behavior`, `*Binding`
- [ ] **T2.2** Update `action.go` - Use unified types, remove old definitions
- [ ] **T2.3** Update `schema.go` - Use unified types
- [ ] **T2.4** Update `layout.go` - Use unified types for all blocks
- [ ] **T2.5** Update `mixin.go` - Verify compatibility
- [ ] **T2.6** Update `repeatable.go` - Verify compatibility

### Phase 3: Update Subsystems

- [ ] **T3.1** Update `builder.go` - New builder methods
- [ ] **T3.2** Update `validator.go` - Remove old interfaces
- [ ] **T3.3** Update `enricher.go` - Remove moved interfaces
- [ ] **T3.4** Update `registry.go` - Remove moved interfaces
- [ ] **T3.5** Update `runtime.go` - Remove moved interfaces
- [ ] **T3.6** Update `parser.go` - Remove moved interfaces
- [ ] **T3.7** Update `i18n.go` - Verify no conflicts
- [ ] **T3.8** Update `business_rules.go` - Verify no conflicts

### Phase 4: Remove Deprecated

- [ ] **T4.1** Remove duplicate types from old locations
- [ ] **T4.2** Remove deprecated interfaces from doc.go, validator.go
- [ ] **T4.3** Clean up doc.go

### Phase 5: Update Tests

- [ ] **T5.1** Update `field_test.go`
- [ ] **T5.2** Update `action_test.go`
- [ ] **T5.3** Update `schema_test.go`
- [ ] **T5.4** Update `builder_test.go`
- [ ] **T5.5** Update all remaining test files

### Phase 6: Final Verification

- [ ] **T6.1** Final build verification
- [ ] **T6.2** Final test verification
- [ ] **T6.3** Generate documentation

---

## File Structure After Refactoring

```
schema/
├── types.go           # Common types (enums, configs, workflow, metadata)
├── interface.go       # ALL interfaces
├── conditional.go     # Unified Conditional type
├── style.go           # Unified Style type  
├── event.go           # Event, Events, BehaviorMetadata
├── behavior.go        # Behavior, Trigger, SwapStrategy
├── binding.go         # Binding type
├── errors.go          # Error types and implementations
│
├── schema.go          # Schema struct (uses unified types)
├── field.go           # Field struct + FieldType, FieldValidation, etc.
├── action.go          # Action struct + ActionType, ActionConfirm
├── layout.go          # Layout, Section, Group, Tab, Step, LayoutBlock
│
├── builder.go         # SchemaBuilder
├── validator.go       # DataValidator implementation
├── enricher.go        # DefaultEnricher implementation
├── registry.go        # Registry implementation
├── runtime.go         # Runtime implementation
├── parser.go          # Parser implementation
│
├── mixin.go           # Mixin, MixinRegistry
├── repeatable.go      # RepeatableField
├── business_rules.go  # BusinessRule, BusinessRuleEngine
├── i18n.go            # I18nManager, translations
│
├── doc.go             # Package documentation only
│
└── *_test.go          # Test files
```

---

## Type Mapping Reference

### Field Struct Changes

```go
// Before
type Field struct {
    Conditional *FieldConditional  // REMOVE
    HTMX        *FieldHTMX         // REMOVE
    Alpine      *FieldAlpine       // REMOVE
    Style       *FieldStyle        // REMOVE
}

// After
type Field struct {
    Conditional *Conditional       // from conditional.go
    Behavior    *Behavior          // from behavior.go
    Binding     *Binding           // from binding.go
    Style       *Style             // from style.go
    Events      *Events            // from event.go
}
```

### Action Struct Changes

```go
// Before
type Action struct {
    HTMX        *ActionHTMX
    Theme       *ActionTheme
    Config      *ActionConfig
    Permissions *ActionPermissions
}

// After
type Action struct {
    Behavior    *Behavior          // replaces HTMX
    Binding     *Binding           // new
    Style       *Style             // replaces Theme
    Conditional *Conditional       // new
    Config      map[string]any     // simplified
    Permissions []string           // simplified
    Events      *Events            // new
}
```

### Layout Component Changes

```go
// Before (each had different style types)
type Section struct { Style *SectionStyle }
type Tab struct { Style *TabStyle }
type Step struct { Style *StepStyle }

// After (all use unified Style)
type Section struct { Style *Style; Conditional *Conditional }
type Tab struct { Style *Style; Conditional *Conditional }
type Step struct { Style *Style; Conditional *Conditional }
```

---

## Recovery Procedures

### If Build Fails

1. Check the error message
2. Run `./schema-refactor.sh --status` to see progress
3. Fix the specific error manually or run `./schema-refactor.sh --task TX.X`
4. Resume with `./schema-refactor.sh --resume`

### If Tests Fail

1. Run `go test ./schema/... -v` to see detailed output
2. Check if old type names are still used in tests
3. Update test files to use new types
4. Re-run tests

### To Start Over

```bash
./schema-refactor.sh --reset
rm -rf ./schema
cp -r ./schema_backup_* ./schema
./schema-refactor.sh --auto
```

---

## Estimated Reduction

| Category | Before | After | Reduction |
|----------|--------|-------|-----------|
| Conditional types | 4 | 1 | 75% |
| Style types | 7 | 1 | 86% |
| Behavior/HTMX types | 3 | 1 | 67% |
| Binding/Alpine types | 3 | 1 | 67% |
| Validator interfaces | 6 | 2 | 67% |
| Renderer interfaces | 2 | 1 | 50% |
| Duplicate structs | 4 | 0 | 100% |
| **Total types/interfaces** | ~30 | ~12 | **60%** |
