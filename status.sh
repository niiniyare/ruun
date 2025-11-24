#!/bin/bash
echo "🎯 Refactoring Status Dashboard"
echo "================================"

total=$(find views/components/atoms -name "*.templ" | wc -l)
refactored=$(grep -l "class=\"btn\|class=\"input\|class=\"badge\"" views/components/atoms/*.templ 2>/dev/null | wc -l)
percent=$((refactored * 100 / total))

echo "Atoms: $refactored/$total ($percent%)"
echo ""
echo "📋 Next to refactor:"
grep "^- \[ \]" refactoring_progress.md | head -5
