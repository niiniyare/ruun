# Schema Package Refactoring Context

## Current Task
Refactoring types.go to eliminate:
1. Unused types (types with no references)
2. Redundant types (different names, same structure)

## Rules
- Preserve all exported types that are part of public API
- When consolidating, prefer the more descriptive name
- Update all references when renaming
- Maintain backward compatibility where possible

## Files to Focus
- types.go (primary)
- field.go, action.go, layout.go (consumers)
