#!/bin/bash
# Basecoat Compliance Linting Script
# Prevents regression to dynamic class patterns

set -e

echo "🔍 Running Basecoat compliance checks..."

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Counter for violations
violations=0

# Check for utils.TwMerge usage
echo -e "${BLUE}Checking for utils.TwMerge violations...${NC}"
if grep -r "utils\.TwMerge" views/ --include="*.go" --include="*.templ" 2>/dev/null; then
    echo -e "${RED}❌ VIOLATION: Found utils.TwMerge usage${NC}"
    echo -e "${YELLOW}   Replace with static class mapping or strings.Join()${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ No utils.TwMerge violations found${NC}"
fi

# Check for ClassName props
echo -e "${BLUE}Checking for ClassName prop violations...${NC}"
if grep -r "ClassName.*string" views/ --include="*.go" --include="*.templ" 2>/dev/null; then
    echo -e "${RED}❌ VIOLATION: Found ClassName prop definitions${NC}"
    echo -e "${YELLOW}   Replace with semantic boolean props${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ No ClassName prop violations found${NC}"
fi

# Check for dynamic class building patterns
echo -e "${BLUE}Checking for dynamic class building patterns...${NC}"
dynamic_patterns=(
    'fmt\.Sprintf.*class'
    'strings\.Join.*class'
    '\+.*class.*\+'
    'class.*fmt\.Sprintf'
    'class.*\+.*\+'
)

for pattern in "${dynamic_patterns[@]}"; do
    if grep -r -E "$pattern" views/ --include="*.go" --include="*.templ" 2>/dev/null; then
        echo -e "${RED}❌ VIOLATION: Found dynamic class building pattern: $pattern${NC}"
        echo -e "${YELLOW}   Use static switch statements instead${NC}"
        violations=$((violations + 1))
    fi
done

if [ $violations -eq 0 ]; then
    echo -e "${GREEN}✅ No dynamic class building patterns found${NC}"
fi

# Check for string concatenation in templates
echo -e "${BLUE}Checking for template string concatenation...${NC}"
if grep -r 'class={.*".*".*+' views/ --include="*.templ" 2>/dev/null; then
    echo -e "${RED}❌ VIOLATION: Found string concatenation in class attributes${NC}"
    echo -e "${YELLOW}   Use data attributes instead: data-variant={ variant }${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ No template string concatenation found${NC}"
fi

# Check for deprecated class utilities
echo -e "${BLUE}Checking for deprecated class utilities...${NC}"
deprecated_utils=(
    'clsx'
    'classnames'
    'tw-merge'
    'twMerge'
    'cn\('
    'classNames\('
)

for util in "${deprecated_utils[@]}"; do
    if grep -r -E "$util" views/ --include="*.go" --include="*.templ" 2>/dev/null; then
        echo -e "${RED}❌ VIOLATION: Found deprecated utility: $util${NC}"
        echo -e "${YELLOW}   Use static Basecoat classes instead${NC}"
        violations=$((violations + 1))
    fi
done

if [ $violations -eq 0 ]; then
    echo -e "${GREEN}✅ No deprecated utilities found${NC}"
fi

# Check for missing type-safe enums
echo -e "${BLUE}Checking for type-safe enum usage...${NC}"
if grep -r 'Variant.*:.*"' views/ --include="*.go" --include="*.templ" 2>/dev/null; then
    echo -e "${RED}❌ VIOLATION: Found string variant instead of enum${NC}"
    echo -e "${YELLOW}   Use atoms.ButtonVariantPrimary instead of \"primary\"${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ All variants use type-safe enums${NC}"
fi

# Check for proper data attribute usage
echo -e "${BLUE}Checking for proper data attribute formatting...${NC}"
if grep -r 'data-.*{.*[^t]"' views/ --include="*.templ" 2>/dev/null; then
    echo -e "${RED}❌ VIOLATION: Found improperly formatted data attributes${NC}"
    echo -e "${YELLOW}   Use fmt.Sprintf(\"%t\", boolValue) for boolean attributes${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ All data attributes properly formatted${NC}"
fi

# Check for Basecoat class usage
echo -e "${BLUE}Checking for Basecoat-compliant class patterns...${NC}"
basecoat_classes=(
    'btn'
    'card'
    'field'
    'table'
    'tabs'
    'sidebar'
    'badge'
    'progress'
    'avatar'
)

# Verify at least some Basecoat classes are being used
found_basecoat=0
for class in "${basecoat_classes[@]}"; do
    if grep -r "class=\"$class" views/ --include="*.templ" >/dev/null 2>&1; then
        found_basecoat=$((found_basecoat + 1))
    fi
done

if [ $found_basecoat -eq 0 ]; then
    echo -e "${RED}❌ WARNING: No Basecoat classes found${NC}"
    echo -e "${YELLOW}   Verify Basecoat CSS is being used correctly${NC}"
    violations=$((violations + 1))
else
    echo -e "${GREEN}✅ Found $found_basecoat Basecoat class patterns${NC}"
fi

# Generate compliance report
echo ""
echo "==================================="
echo "📊 BASECOAT COMPLIANCE REPORT"
echo "==================================="

if [ $violations -eq 0 ]; then
    echo -e "${GREEN}🏆 PERFECT COMPLIANCE!${NC}"
    echo -e "${GREEN}   All components follow Basecoat patterns${NC}"
    echo -e "${GREEN}   Zero violations found${NC}"
    exit 0
else
    echo -e "${RED}❌ VIOLATIONS FOUND: $violations${NC}"
    echo ""
    echo "🔧 REMEDIATION STEPS:"
    echo "1. Replace utils.TwMerge with static switch statements"
    echo "2. Remove ClassName props, add semantic boolean props"
    echo "3. Use data attributes instead of dynamic classes"
    echo "4. Convert string variants to type-safe enums"
    echo "5. Follow examples in docs/component-examples.md"
    echo ""
    echo "📖 See migration guide: docs/migration-guide.md"
    exit 1
fi