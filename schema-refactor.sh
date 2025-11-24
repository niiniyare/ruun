#!/bin/bash
set -e

echo "=== Phase 1.1: Create types_v2.go ==="
# Run Claude CLI command for Step 1.1
git add -A && git commit -m "feat(schema): add unified Behavior, Binding, Style types"

echo "=== Phase 1.2: Create events_unified.go ==="
# Run Claude CLI command for Step 1.2
git add -A && git commit -m "feat(schema): add unified Events type"

echo "=== Phase 1.3: Create layout_unified.go ==="
# Run Claude CLI command for Step 1.3
git add -A && git commit -m "feat(schema): add unified LayoutBlock type"

echo "=== Phase 1.4: Create interfaces.go ==="
# Run Claude CLI command for Step 1.4
git add -A && git commit -m "feat(schema): add public interfaces"

echo "=== Phase 2: Update core structs ==="
# Run Claude CLI commands for Steps 2.1-2.4
git add -A && git commit -m "refactor(schema): update Schema, Field, Action, Layout to use unified types"

echo "=== Phase 3: Fix compilation ==="
# Run Claude CLI command for Phase 3
git add -A && git commit -m "fix(schema): resolve compilation errors after type unification"

echo "=== Phase 4: Update builders ==="
# Run Claude CLI command for Phase 4
git add -A && git commit -m "refactor(schema): update builders for unified types"

echo "=== Phase 5: Delete old types ==="
# Run Claude CLI command for Phase 5
git add -A && git commit -m "refactor(schema): remove deprecated HTMX/Alpine/Style types"

echo "=== Phase 6: Final cleanup ==="
# Run Claude CLI command for Phase 6
git add -A && git commit -m "chore(schema): cleanup and consolidate files"

echo "✅ Refactoring complete!"
go test ./schema/...
