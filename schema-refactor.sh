#!/bin/bash
# schema-refactor.sh

set -e # Exit on error

echo "=== Phase 1.1: Merge GroupStyle into SectionStyle ==="
claude "Merge GroupStyle into SectionStyle in schema package. 
Replace all GroupStyle usages with SectionStyle, delete GroupStyle definition.
Verify with go build."

git add -A && git commit -m "refactor(schema): merge GroupStyle into SectionStyle"

echo "=== Phase 1.2: Create BaseStyle ==="
claude "Create BaseStyle struct with Colors, CustomCSS, BorderRadius fields.
Refactor LayoutStyle, SectionStyle, TabStyle, StepStyle, ActionTheme to embed it.
Run go build to verify."

git add -A && git commit -m "refactor(schema): extract BaseStyle for style composition"

echo "=== Phase 2: Layout Component Base ==="
claude "Create LayoutComponent base struct with ID, Title, Description, Fields, Order, Conditional.
Refactor Section, Group, Tab, Step to embed it. Run tests."

git add -A && git commit -m "refactor(schema): extract LayoutComponent base struct"

echo "=== Phase 3: Inline RepeatableOperations ==="
claude "Move RepeatableOperations methods to RepeatableField, delete RepeatableOperations struct.
Run tests to verify."

git add -A && git commit -m "refactor(schema): inline RepeatableOperations into RepeatableField"

echo "=== Phase 4: BaseMetadata ==="
claude "Create BaseMetadata with Version, CreatedAt, UpdatedAt.
Refactor Meta, StorageMetadata, TranslationMetadata to embed it. Run tests."

git add -A && git commit -m "refactor(schema): extract BaseMetadata for metadata composition"

echo "=== Final Verification ==="
go test ./schema/...
echo "✅ All refactoring complete!"
