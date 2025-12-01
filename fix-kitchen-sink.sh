#!/bin/bash
# =============================================================================
# fix-kitchen-sink.sh - Analyze and Fix Kitchen Sink Demo
# =============================================================================
#
# This script:
# 1. Analyzes ruun-demo/ for build errors
# 2. Analyzes views/components/ to see what ACTUALLY exists
# 3. Reads views/docs/*.md to understand actual component APIs
# 4. Creates missing components OR fixes demo to match reality
# 5. Loops until server successfully runs
#
# Usage:
#   ./fix-kitchen-sink.sh              # Fix errors
#   ./fix-kitchen-sink.sh --serve      # Fix and run server
#   ./fix-kitchen-sink.sh --create     # Create missing layout primitives
#
# =============================================================================

set -o pipefail

# =============================================================================
# Configuration
# =============================================================================

RUUN_PKG="github.com/niiniyare/ruun"
DEMO_DIR="${DEMO_DIR:-./ruun-demo}"
COMPONENTS_DIR="${COMPONENTS_DIR:-./views/components}"
DOCS_DIR="${DOCS_DIR:-./views/docs}"

# Subdirectories
ATOMS_DIR="$COMPONENTS_DIR/atoms"
MOLECULES_DIR="$COMPONENTS_DIR/molecules"
ORGANISMS_DIR="$COMPONENTS_DIR/organisms"
TEMPLATES_DIR="$COMPONENTS_DIR/templates"

MAX_FIX_ATTEMPTS=10
LOG_FILE="./fix-kitchen-sink.log"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
BOLD='\033[1m'
NC='\033[0m'

# =============================================================================
# Logging
# =============================================================================

log() { echo -e "$1" | tee -a "$LOG_FILE"; }
log_info() { log "${BLUE}[INFO]${NC} $1"; }
log_success() { log "${GREEN}[SUCCESS]${NC} $1"; }
log_warn() { log "${YELLOW}[WARN]${NC} $1"; }
log_error() { log "${RED}[ERROR]${NC} $1"; }
log_task() { log "${MAGENTA}[TASK]${NC} $1"; }
log_fix() { log "${CYAN}[FIX]${NC} $1"; }

log_header() {
  echo "" | tee -a "$LOG_FILE"
  echo -e "${BOLD}════════════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}  $1${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}════════════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo "" | tee -a "$LOG_FILE"
}

# =============================================================================
# Helper: Process templ files in a directory
# =============================================================================

process_templ_files() {
  local dir="$1"
  local output_file="$2"

  if [[ ! -d "$dir" ]]; then
    echo "Directory not found: $dir" >>"$output_file"
    return
  fi

  # Use find instead of glob to avoid syntax issues
  while IFS= read -r -d '' f; do
    local name
    name=$(basename "$f" .templ)
    echo "- $name.templ" >>"$output_file"

    # Extract exported templ functions
    if grep -qE "^templ [A-Z]" "$f" 2>/dev/null; then
      grep -E "^templ [A-Z]" "$f" | sed 's/templ /  - /' | sed 's/(.*$/()/' >>"$output_file"
    fi
  done < <(find "$dir" -maxdepth 1 -name "*.templ" -print0 2>/dev/null)
}

# =============================================================================
# Analysis Functions
# =============================================================================

# Discover all existing .templ files and extract component names
discover_components() {
  log_task "Discovering existing components..."

  local inventory_file="$DEMO_DIR/COMPONENT_INVENTORY.md"

  mkdir -p "$DEMO_DIR"

  cat >"$inventory_file" <<'HEADER'
# Ruun Package - Component Inventory

This is an auto-generated inventory of what ACTUALLY exists in the ruun package.

HEADER

  # Atoms
  echo "## Atoms (views/components/atoms/)" >>"$inventory_file"
  echo "" >>"$inventory_file"
  process_templ_files "$ATOMS_DIR" "$inventory_file"
  echo "" >>"$inventory_file"

  # Molecules
  echo "## Molecules (views/components/molecules/)" >>"$inventory_file"
  echo "" >>"$inventory_file"
  process_templ_files "$MOLECULES_DIR" "$inventory_file"
  echo "" >>"$inventory_file"

  # Organisms
  echo "## Organisms (views/components/organisms/)" >>"$inventory_file"
  echo "" >>"$inventory_file"
  process_templ_files "$ORGANISMS_DIR" "$inventory_file"
  echo "" >>"$inventory_file"

  # Templates
  echo "## Templates (views/components/templates/)" >>"$inventory_file"
  echo "" >>"$inventory_file"
  process_templ_files "$TEMPLATES_DIR" "$inventory_file"
  echo "" >>"$inventory_file"

  # Types
  echo "## Types (views/components/types.go)" >>"$inventory_file"
  echo "" >>"$inventory_file"
  if [[ -f "$COMPONENTS_DIR/types.go" ]]; then
    echo '```go' >>"$inventory_file"
    cat "$COMPONENTS_DIR/types.go" >>"$inventory_file"
    echo '```' >>"$inventory_file"
  fi

  log_success "Created $inventory_file"
}

# Collect all documentation
collect_docs() {
  log_task "Collecting component documentation..."

  local docs_file="$DEMO_DIR/COMPONENT_DOCS.md"

  mkdir -p "$DEMO_DIR"

  cat >"$docs_file" <<'HEADER'
# Ruun Package - Component Documentation

Auto-collected from views/docs/*.md

HEADER

  if [[ -d "$DOCS_DIR" ]]; then
    while IFS= read -r -d '' f; do
      local name
      name=$(basename "$f" .md)
      echo "---" >>"$docs_file"
      echo "## $name" >>"$docs_file"
      echo "" >>"$docs_file"
      cat "$f" >>"$docs_file"
      echo "" >>"$docs_file"
    done < <(find "$DOCS_DIR" -maxdepth 1 -name "*.md" -print0 2>/dev/null)
  fi

  log_success "Created $docs_file"
}

# Extract actual component signatures from .templ files
extract_signatures() {
  log_task "Extracting component signatures..."

  local sigs_file="$DEMO_DIR/COMPONENT_SIGNATURES.md"

  mkdir -p "$DEMO_DIR"

  cat >"$sigs_file" <<'HEADER'
# Component Signatures

Extracted from actual .templ files. Use these EXACT signatures in the demo.

HEADER

  echo "## Atoms" >>"$sigs_file"
  echo "" >>"$sigs_file"

  if [[ -d "$ATOMS_DIR" ]]; then
    while IFS= read -r -d '' f; do
      local name
      name=$(basename "$f" .templ)
      echo "### $name" >>"$sigs_file"
      echo '```go' >>"$sigs_file"
      # Extract type definitions (first 50 lines of struct definitions)
      grep -E "^type [A-Z].*struct" "$f" -A 30 2>/dev/null | head -50 >>"$sigs_file"
      # Extract templ function signatures
      grep -E "^templ [A-Z]" "$f" 2>/dev/null >>"$sigs_file"
      echo '```' >>"$sigs_file"
      echo "" >>"$sigs_file"
    done < <(find "$ATOMS_DIR" -maxdepth 1 -name "*.templ" -print0 2>/dev/null)
  fi

  echo "## Molecules" >>"$sigs_file"
  echo "" >>"$sigs_file"

  if [[ -d "$MOLECULES_DIR" ]]; then
    while IFS= read -r -d '' f; do
      local name
      name=$(basename "$f" .templ)
      echo "### $name" >>"$sigs_file"
      echo '```go' >>"$sigs_file"
      grep -E "^type [A-Z].*struct" "$f" -A 30 2>/dev/null | head -50 >>"$sigs_file"
      grep -E "^templ [A-Z]" "$f" 2>/dev/null >>"$sigs_file"
      echo '```' >>"$sigs_file"
      echo "" >>"$sigs_file"
    done < <(find "$MOLECULES_DIR" -maxdepth 1 -name "*.templ" -print0 2>/dev/null)
  fi

  echo "## Organisms" >>"$sigs_file"
  echo "" >>"$sigs_file"

  if [[ -d "$ORGANISMS_DIR" ]]; then
    while IFS= read -r -d '' f; do
      local name
      name=$(basename "$f" .templ)
      echo "### $name" >>"$sigs_file"
      echo '```go' >>"$sigs_file"
      grep -E "^type [A-Z].*struct" "$f" -A 30 2>/dev/null | head -50 >>"$sigs_file"
      grep -E "^templ [A-Z]" "$f" 2>/dev/null >>"$sigs_file"
      echo '```' >>"$sigs_file"
      echo "" >>"$sigs_file"
    done < <(find "$ORGANISMS_DIR" -maxdepth 1 -name "*.templ" -print0 2>/dev/null)
  fi

  echo "## Templates" >>"$sigs_file"
  echo "" >>"$sigs_file"

  if [[ -d "$TEMPLATES_DIR" ]]; then
    while IFS= read -r -d '' f; do
      local name
      name=$(basename "$f" .templ)
      echo "### $name" >>"$sigs_file"
      echo '```go' >>"$sigs_file"
      grep -E "^type [A-Z].*struct" "$f" -A 30 2>/dev/null | head -50 >>"$sigs_file"
      grep -E "^templ [A-Z]" "$f" 2>/dev/null >>"$sigs_file"
      echo '```' >>"$sigs_file"
      echo "" >>"$sigs_file"
    done < <(find "$TEMPLATES_DIR" -maxdepth 1 -name "*.templ" -print0 2>/dev/null)
  fi

  log_success "Created $sigs_file"
}

# =============================================================================
# Create Missing Layout Primitives
# =============================================================================

create_layout_primitives() {
  log_header "Creating Missing Layout Primitives"

  log_task "Creating layout.templ with Flex, Grid, Stack, Box, Center, Spacer, Divider..."

  cat >"$ATOMS_DIR/layout.templ" <<'LAYOUT_TEMPL'
package atoms

import "github.com/a-h/templ"

// =============================================================================
// Spacing and Size Types
// =============================================================================

type Spacing string

const (
    Gap0  Spacing = "gap-0"
    Gap1  Spacing = "gap-1"
    Gap2  Spacing = "gap-2"
    Gap3  Spacing = "gap-3"
    Gap4  Spacing = "gap-4"
    Gap5  Spacing = "gap-5"
    Gap6  Spacing = "gap-6"
    Gap8  Spacing = "gap-8"
    Gap10 Spacing = "gap-10"
    Gap12 Spacing = "gap-12"
)

// Padding constants
const (
    P0  Spacing = "p-0"
    P1  Spacing = "p-1"
    P2  Spacing = "p-2"
    P3  Spacing = "p-3"
    P4  Spacing = "p-4"
    P5  Spacing = "p-5"
    P6  Spacing = "p-6"
    P8  Spacing = "p-8"
    P10 Spacing = "p-10"
    P12 Spacing = "p-12"
)

// Margin constants
const (
    M0  Spacing = "m-0"
    M1  Spacing = "m-1"
    M2  Spacing = "m-2"
    M3  Spacing = "m-3"
    M4  Spacing = "m-4"
    M5  Spacing = "m-5"
    M6  Spacing = "m-6"
    M8  Spacing = "m-8"
)

type FlexDirection string

const (
    Row        FlexDirection = "flex-row"
    Col        FlexDirection = "flex-col"
    RowReverse FlexDirection = "flex-row-reverse"
    ColReverse FlexDirection = "flex-col-reverse"
)

type FlexWrap string

const (
    NoWrap      FlexWrap = "flex-nowrap"
    Wrap        FlexWrap = "flex-wrap"
    WrapReverse FlexWrap = "flex-wrap-reverse"
)

type Justify string

const (
    JustifyStart   Justify = "justify-start"
    JustifyCenter  Justify = "justify-center"
    JustifyEnd     Justify = "justify-end"
    JustifyBetween Justify = "justify-between"
    JustifyAround  Justify = "justify-around"
    JustifyEvenly  Justify = "justify-evenly"
)

type Align string

const (
    AlignStart    Align = "items-start"
    AlignCenter   Align = "items-center"
    AlignEnd      Align = "items-end"
    AlignStretch  Align = "items-stretch"
    AlignBaseline Align = "items-baseline"
)

type Rounded string

const (
    RoundedNone Rounded = "rounded-none"
    RoundedSm   Rounded = "rounded-sm"
    RoundedMd   Rounded = "rounded-md"
    RoundedLg   Rounded = "rounded-lg"
    RoundedXl   Rounded = "rounded-xl"
    RoundedFull Rounded = "rounded-full"
)

type Shadow string

const (
    ShadowNone Shadow = "shadow-none"
    ShadowSm   Shadow = "shadow-sm"
    ShadowMd   Shadow = "shadow-md"
    ShadowLg   Shadow = "shadow-lg"
    ShadowXl   Shadow = "shadow-xl"
)

type BgColor string

const (
    BgBackground  BgColor = "bg-background"
    BgMuted       BgColor = "bg-muted"
    BgCard        BgColor = "bg-card"
    BgPrimary     BgColor = "bg-primary"
    BgSecondary   BgColor = "bg-secondary"
    BgAccent      BgColor = "bg-accent"
    BgDestructive BgColor = "bg-destructive"
)

type MaxWidth string

const (
    MaxWidthXs   MaxWidth = "max-w-xs"
    MaxWidthSm   MaxWidth = "max-w-sm"
    MaxWidthMd   MaxWidth = "max-w-md"
    MaxWidthLg   MaxWidth = "max-w-lg"
    MaxWidthXl   MaxWidth = "max-w-xl"
    MaxWidth2xl  MaxWidth = "max-w-2xl"
    MaxWidthFull MaxWidth = "max-w-full"
)

type Width string

const (
    WAuto Width = "w-auto"
    WFull Width = "w-full"
    W48   Width = "w-48"
    W64   Width = "w-64"
    W96   Width = "w-96"
)

type Height string

const (
    HAuto Height = "h-auto"
    HFull Height = "h-full"
    H4    Height = "h-4"
    H8    Height = "h-8"
    H12   Height = "h-12"
)

// =============================================================================
// Flex Component
// =============================================================================

type FlexProps struct {
    Direction FlexDirection
    Wrap      FlexWrap
    Justify   Justify
    Align     Align
    Gap       Spacing
    Padding   Spacing
    Class     string
    Attrs     templ.Attributes
}

templ Flex(props FlexProps) {
    <div
        class={ templ.Classes(
            "flex",
            string(props.Direction),
            string(props.Wrap),
            string(props.Justify),
            string(props.Align),
            string(props.Gap),
            string(props.Padding),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Grid Component
// =============================================================================

type GridProps struct {
    Cols    int
    Rows    int
    Gap     Spacing
    Padding Spacing
    Class   string
    Attrs   templ.Attributes
}

func gridColsClass(cols int) string {
    switch cols {
    case 1:
        return "grid-cols-1"
    case 2:
        return "grid-cols-2"
    case 3:
        return "grid-cols-3"
    case 4:
        return "grid-cols-4"
    case 5:
        return "grid-cols-5"
    case 6:
        return "grid-cols-6"
    case 12:
        return "grid-cols-12"
    default:
        return "grid-cols-1"
    }
}

templ Grid(props GridProps) {
    <div
        class={ templ.Classes(
            "grid",
            gridColsClass(props.Cols),
            string(props.Gap),
            string(props.Padding),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Stack Component (simplified Flex - vertical by default)
// =============================================================================

type StackDirection string

const (
    Vertical   StackDirection = "flex-col"
    Horizontal StackDirection = "flex-row"
)

type StackProps struct {
    Direction StackDirection
    Gap       Spacing
    Align     Align
    Padding   Spacing
    Class     string
    Attrs     templ.Attributes
}

templ Stack(props StackProps) {
    <div
        class={ templ.Classes(
            "flex",
            "flex-col",
            templ.KV(string(props.Direction), props.Direction != "" && props.Direction != Vertical),
            string(props.Gap),
            string(props.Align),
            string(props.Padding),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Box Component (generic container)
// =============================================================================

type BoxProps struct {
    Padding  Spacing
    Margin   Spacing
    BgColor  BgColor
    Border   bool
    Rounded  Rounded
    Shadow   Shadow
    MaxWidth MaxWidth
    Width    Width
    Height   Height
    Class    string
    Attrs    templ.Attributes
}

templ Box(props BoxProps) {
    <div
        class={ templ.Classes(
            string(props.Padding),
            string(props.Margin),
            string(props.BgColor),
            templ.KV("border", props.Border),
            string(props.Rounded),
            string(props.Shadow),
            string(props.MaxWidth),
            string(props.Width),
            string(props.Height),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Center Component
// =============================================================================

type CenterProps struct {
    Padding  Spacing
    MaxWidth MaxWidth
    Class    string
    Attrs    templ.Attributes
}

templ Center(props CenterProps) {
    <div
        class={ templ.Classes(
            "flex",
            "items-center",
            "justify-center",
            string(props.Padding),
            string(props.MaxWidth),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Spacer Component
// =============================================================================

type SpacerProps struct {
    Size  Spacing
    Grow  bool
    Class string
}

templ Spacer(props SpacerProps) {
    if props.Grow {
        <div class={ templ.Classes("flex-grow", props.Class) }></div>
    } else {
        <div class={ templ.Classes(string(props.Size), props.Class) }></div>
    }
}

// =============================================================================
// Divider/Separator Component
// =============================================================================

type Orientation string

const (
    OrientationHorizontal Orientation = "horizontal"
    OrientationVertical   Orientation = "vertical"
)

type DividerProps struct {
    Orientation Orientation
    Label       string
    Class       string
}

templ Divider(props DividerProps) {
    if props.Label != "" {
        <div class={ templ.Classes("flex items-center gap-4", props.Class) }>
            <div class="flex-1 border-t border-border"></div>
            <span class="text-sm text-muted-foreground">{ props.Label }</span>
            <div class="flex-1 border-t border-border"></div>
        </div>
    } else if props.Orientation == OrientationVertical {
        <div class={ templ.Classes("w-px bg-border self-stretch", props.Class) }></div>
    } else {
        <div class={ templ.Classes("h-px bg-border w-full", props.Class) }></div>
    }
}

// Separator is an alias for Divider
type SeparatorProps struct {
    Orientation Orientation
    Class       string
}

templ Separator(props SeparatorProps) {
    @Divider(DividerProps{Orientation: props.Orientation, Class: props.Class})
}
LAYOUT_TEMPL

  log_success "Created $ATOMS_DIR/layout.templ"
}

# Create Text component if missing
create_text_component() {
  if [[ -f "$ATOMS_DIR/text.templ" ]]; then
    log_info "text.templ already exists, skipping"
    return
  fi

  log_task "Creating text.templ..."

  cat >"$ATOMS_DIR/text.templ" <<'TEXT_TEMPL'
package atoms

import "github.com/a-h/templ"

type TextSize string

const (
    TextXs   TextSize = "text-xs"
    TextSm   TextSize = "text-sm"
    TextBase TextSize = "text-base"
    TextLg   TextSize = "text-lg"
    TextXl   TextSize = "text-xl"
    Text2xl  TextSize = "text-2xl"
)

type FontWeight string

const (
    FontNormal   FontWeight = "font-normal"
    FontMedium   FontWeight = "font-medium"
    FontSemibold FontWeight = "font-semibold"
    FontBold     FontWeight = "font-bold"
)

type TextProps struct {
    Text   string
    Size   TextSize
    Weight FontWeight
    Muted  bool
    Class  string
    Attrs  templ.Attributes
}

templ Text(props TextProps) {
    <p
        class={ templ.Classes(
            string(props.Size),
            string(props.Weight),
            templ.KV("text-muted-foreground", props.Muted),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { props.Text }
    </p>
}

// Span variant for inline text
type SpanProps struct {
    Text   string
    Size   TextSize
    Weight FontWeight
    Muted  bool
    Class  string
    Attrs  templ.Attributes
}

templ Span(props SpanProps) {
    <span
        class={ templ.Classes(
            string(props.Size),
            string(props.Weight),
            templ.KV("text-muted-foreground", props.Muted),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { props.Text }
    </span>
}
TEXT_TEMPL

  log_success "Created $ATOMS_DIR/text.templ"
}

# Create Code component if missing
create_code_component() {
  if [[ -f "$ATOMS_DIR/code.templ" ]]; then
    log_info "code.templ already exists, skipping"
    return
  fi

  log_task "Creating code.templ..."

  cat >"$ATOMS_DIR/code.templ" <<'CODE_TEMPL'
package atoms

import "github.com/a-h/templ"

type CodeProps struct {
    Text  string
    Block bool
    Lang  string
    Class string
    Attrs templ.Attributes
}

templ Code(props CodeProps) {
    if props.Block {
        <pre class={ templ.Classes("bg-muted rounded-md p-4 overflow-x-auto", props.Class) } { props.Attrs... }>
            <code class="text-sm font-mono">
                { props.Text }
            </code>
        </pre>
    } else {
        <code
            class={ templ.Classes(
                "relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm",
                props.Class,
            ) }
            { props.Attrs... }
        >
            { props.Text }
        </code>
    }
}
CODE_TEMPL

  log_success "Created $ATOMS_DIR/code.templ"
}

# =============================================================================
# Build and Fix Loop
# =============================================================================

get_build_errors() {
  local build_log="$1"

  cd "$DEMO_DIR" 2>/dev/null || return 1

  # Run templ generate
  templ generate >"$build_log" 2>&1
  local templ_exit=$?

  # Run go build
  go build ./... >>"$build_log" 2>&1
  local go_exit=$?

  if [[ $templ_exit -eq 0 && $go_exit -eq 0 ]]; then
    return 0
  fi

  return 1
}

fix_with_claude() {
  local errors="$1"
  local inventory="$2"
  local signatures="$3"

  claude --dangerously-skip-permissions "Fix build errors in the ruun-demo kitchen sink.

## BUILD ERRORS:
$errors

## WHAT ACTUALLY EXISTS IN RUUN PACKAGE:
$inventory

## ACTUAL COMPONENT SIGNATURES:
$signatures

## INSTRUCTIONS:

1. Read the errors carefully
2. Look at what components ACTUALLY exist in the ruun package
3. Fix the demo files to use ONLY components that exist
4. If a component doesn't exist, either:
   a) Remove it from the demo
   b) Use an alternative that exists
   c) Use raw HTML with Tailwind (as last resort for layout only)

CRITICAL RULES:
- Import paths must be: $RUUN_PKG/views/components/atoms, etc.
- Use ONLY exported functions that actually exist
- Match the EXACT function signatures from the signatures above
- Layout primitives (Flex, Grid, Stack, Box) are now in atoms/layout.templ

Fix ALL errors. Make the demo compile and run."
}

run_fix_loop() {
  log_header "Running Fix Loop"

  local attempt=1
  local build_log="./tmp/log/ruun-fix-build-$$.log"

  # First, collect information about what exists
  discover_components
  extract_signatures
  collect_docs

  local inventory
  local signatures
  inventory=$(cat "$DEMO_DIR/COMPONENT_INVENTORY.md" 2>/dev/null)
  signatures=$(cat "$DEMO_DIR/COMPONENT_SIGNATURES.md" 2>/dev/null)

  while [[ $attempt -le $MAX_FIX_ATTEMPTS ]]; do
    log_info "Fix attempt $attempt/$MAX_FIX_ATTEMPTS..."

    if get_build_errors "$build_log"; then
      log_success "Build successful!"
      rm -f "$build_log"
      return 0
    fi

    local errors
    errors=$(grep -E "(error|Error|undefined|cannot|invalid)" "$build_log" | head -50)

    if [[ -z "$errors" ]]; then
      log_warn "No specific errors found, showing full log:"
      cat "$build_log"
    fi

    log_fix "Errors found:"
    echo "$errors" | head -20

    log_fix "Asking Claude to fix..."
    fix_with_claude "$errors" "$inventory" "$signatures"

    ((attempt++))
    sleep 2
  done

  log_error "Could not fix after $MAX_FIX_ATTEMPTS attempts"
  log_error "Last errors:"
  cat "$build_log"
  rm -f "$build_log"
  return 1
}

# =============================================================================
# Comprehensive Analysis and Fix
# =============================================================================

analyze_and_fix() {
  log_header "Comprehensive Analysis and Fix"

  # Step 1: Discover what exists
  discover_components
  extract_signatures
  collect_docs

  # Step 2: Check if layout primitives exist, create if not
  if ! grep -q "templ Flex" "$ATOMS_DIR"/*.templ 2>/dev/null; then
    log_warn "Layout primitives (Flex, Grid, Stack, Box) not found"
    log_info "Creating layout primitives..."
    create_layout_primitives
  else
    log_info "Layout primitives already exist"
  fi

  # Step 3: Check for Text component
  create_text_component

  # Step 4: Check for Code component
  create_code_component

  # Step 5: Re-discover after creating components
  discover_components
  extract_signatures

  # Step 6: Now ask Claude to fix the demo based on what actually exists
  log_task "Fixing demo to match actual components..."

  local inventory
  local signatures
  inventory=$(cat "$DEMO_DIR/COMPONENT_INVENTORY.md" 2>/dev/null)
  signatures=$(cat "$DEMO_DIR/COMPONENT_SIGNATURES.md" 2>/dev/null)

  claude --dangerously-skip-permissions "Analyze and fix the ruun-demo kitchen sink to work with the actual ruun package.

## DEMO LOCATION: $DEMO_DIR

## RUUN PACKAGE LOCATION: $COMPONENTS_DIR

## WHAT ACTUALLY EXISTS:
$inventory

## ACTUAL COMPONENT SIGNATURES:
$signatures

## TASK:

1. Read ALL .templ files in $DEMO_DIR/sections/
2. Check what components they try to use
3. Compare against what ACTUALLY exists in the ruun package
4. Fix ALL files to use only components that exist

## IMPORT PATHS:
The demo should import from:
- \"$RUUN_PKG/views/components/atoms\"
- \"$RUUN_PKG/views/components/molecules\"
- \"$RUUN_PKG/views/components/organisms\"
- \"$RUUN_PKG/views/components/templates\"

## FIXING STRATEGY:

1. If a component exists but with different props, fix the props
2. If a component doesn't exist at all:
   - Use an alternative component that exists
   - Or use raw HTML with Tailwind classes for layout
   - Or remove that section from the demo
3. Make sure all imports are correct

## LAYOUT PRIMITIVES NOW AVAILABLE:
These are now in atoms/layout.templ:
- Flex(FlexProps)
- Grid(GridProps)
- Stack(StackProps)
- Box(BoxProps)
- Center(CenterProps)
- Spacer(SpacerProps)
- Divider(DividerProps)
- Separator(SeparatorProps)

## EXAMPLE FIXES:

If demo uses:
  @atoms.Button(atoms.ButtonProps{Text: \"Save\", Variant: atoms.ButtonPrimary})
  
But Button has different props, read the actual signature and fix accordingly.

Go through EVERY file in $DEMO_DIR/sections/ and fix ALL errors.
After fixing, the demo should compile with: cd $DEMO_DIR && templ generate && go build ./..."

  # Step 7: Run the fix loop
  run_fix_loop
}

# =============================================================================
# Start Server
# =============================================================================

start_server() {
  log_header "Starting Demo Server"

  cd "$DEMO_DIR" || {
    log_error "Demo directory not found"
    exit 1
  }

  log_info "Generating templ files..."
  templ generate || {
    log_error "templ generate failed"
    exit 1
  }

  log_info "Building..."
  go build -o server . || {
    log_error "Build failed"
    exit 1
  }

  log_success "Starting server at http://localhost:3000"
  ./server
}

# =============================================================================
# Main
# =============================================================================

check_prerequisites() {
  log_info "Checking prerequisites..."

  command -v claude &>/dev/null || {
    log_error "Claude CLI required"
    exit 1
  }
  command -v go &>/dev/null || {
    log_error "Go required"
    exit 1
  }
  command -v templ &>/dev/null || {
    log_error "templ CLI required"
    exit 1
  }

  [[ -d "$COMPONENTS_DIR" ]] || {
    log_error "Components directory not found: $COMPONENTS_DIR"
    exit 1
  }

  log_success "Prerequisites OK"
}

show_help() {
  cat <<'EOF'
Fix Kitchen Sink Demo

Analyzes the ruun-demo and views/components to fix all errors.

Usage:
  ./fix-kitchen-sink.sh              # Analyze and fix errors
  ./fix-kitchen-sink.sh --serve      # Fix and run server
  ./fix-kitchen-sink.sh --create     # Create missing layout primitives only
  ./fix-kitchen-sink.sh --analyze    # Just analyze, don't fix
  ./fix-kitchen-sink.sh --loop       # Run fix loop only (assumes analysis done)

What it does:
1. Discovers what components actually exist in views/components/
2. Extracts actual function signatures from .templ files
3. Collects documentation from views/docs/
4. Creates missing layout primitives (Flex, Grid, Stack, Box) if needed
5. Fixes demo files to use only existing components
6. Loops until build succeeds

Output files (in ruun-demo/):
  - COMPONENT_INVENTORY.md  - List of all existing components
  - COMPONENT_SIGNATURES.md - Actual function signatures
  - COMPONENT_DOCS.md       - Collected documentation
EOF
}

main() {
  # Initialize log
  echo "=== Fix Kitchen Sink $(date) ===" >"$LOG_FILE"

  case "${1:-}" in
  --help | -h)
    show_help
    exit 0
    ;;
  --serve)
    check_prerequisites
    analyze_and_fix
    start_server
    ;;
  --create)
    check_prerequisites
    create_layout_primitives
    create_text_component
    create_code_component
    log_success "Created missing components"
    ;;
  --analyze)
    check_prerequisites
    discover_components
    extract_signatures
    collect_docs
    log_success "Analysis complete. Check $DEMO_DIR/*.md files"
    ;;
  --loop)
    check_prerequisites
    run_fix_loop
    ;;
  *)
    check_prerequisites
    analyze_and_fix
    ;;
  esac
}

main "$@"
