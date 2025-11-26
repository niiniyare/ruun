#!/bin/bash
# =============================================================================
# component-refactor.sh - Complete UI Component Refactoring System
# =============================================================================
#
# Refactoring script for Awo ERP UI components following:
# - Basecoat UI patterns (semantic CSS classes like shadcn/ui)
# - Atomic Design principles (atoms → molecules → organisms → templates → pages)
# - HTMX + Alpine.js integration
# - Templ best practices (Classes, KV, Attributes, Component)
#
# BASECOAT PATTERN (from basecoatui.com):
#   - Semantic classes: card, card-header, card-title, btn, field, input
#   - Part classes: card-content, field-label, field-input, checkbox-indicator
#   - Variant classes: btn-outline, btn-destructive
#   - Tailwind utilities for layout: grid, flex, gap-4, w-full, items-center
#
# TEMPL PATTERNS:
#   ✅ TwMerge - resolve Tailwind layout utility conflicts
#   ✅ templ.Classes() - conditional class composition
#   ✅ templ.KV() - conditional class inclusion
#   ✅ templ.Attributes - spread hx-*, aria-*, data-* attributes
#   ✅ templ.Component - slots, icons, composition
#   ❌ ClassName string prop - leaky abstraction, breaks encapsulation
#
# DEMO SERVER:
#   After refactoring, run: ./component-refactor.sh --demo
#   Opens kitchen-sink style demo at http://localhost:3000
#
# Component Counts:
#   - 24 Atoms, 20 Molecules, 18 Organisms, 12 Templates, 15 Pages
#
# Usage:
#   ./component-refactor.sh              # Interactive mode
#   ./component-refactor.sh --auto       # Run all tasks automatically
#   ./component-refactor.sh --resume     # Resume from last checkpoint
#   ./component-refactor.sh --status     # Show current progress
#   ./component-refactor.sh --task T1.1  # Run specific task
#   ./component-refactor.sh --phase 3    # Run specific phase
#   ./component-refactor.sh --demo       # Start demo server
#   ./component-refactor.sh --css-only   # Generate custom CSS only
#
# =============================================================================

set -o pipefail

# =============================================================================
# Configuration
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Directory paths
VIEWS_DIR="${VIEWS_DIR:-./views}"
COMPONENTS_DIR="${COMPONENTS_DIR:-$VIEWS_DIR/components}"
ATOMS_DIR="${ATOMS_DIR:-$COMPONENTS_DIR/atoms}"
MOLECULES_DIR="${MOLECULES_DIR:-$COMPONENTS_DIR/molecules}"
ORGANISMS_DIR="${ORGANISMS_DIR:-$COMPONENTS_DIR/organisms}"
TEMPLATES_DIR="${TEMPLATES_DIR:-$COMPONENTS_DIR/templates}"
PAGES_DIR="${PAGES_DIR:-$VIEWS_DIR/pages}"

# Static assets
STATIC_DIR="${STATIC_DIR:-./static}"
CSS_DIR="${CSS_DIR:-$STATIC_DIR/dist/css}"
CUSTOM_CSS_DIR="${CUSTOM_CSS_DIR:-$CSS_DIR/components}"
JS_DIR="${JS_DIR:-$STATIC_DIR/js}"

# Basecoat resources
BASECOAT_CSS="${BASECOAT_CSS:-$CSS_DIR/basecoat.css}"
BASECOAT_JS_DIR="${BASECOAT_JS_DIR:-$JS_DIR/basecoat}"

# Utils package
UTILS_PKG="${UTILS_PKG:-pkg/utils}"

# Demo server
DEMO_DIR="${DEMO_DIR:-$VIEWS_DIR/demo}"
DEMO_SERVER="${DEMO_SERVER:-$DEMO_DIR/server.go}"
DEMO_PAGE="${DEMO_PAGE:-$DEMO_DIR/demo.templ}"
DEMO_PORT="${DEMO_PORT:-3000}"

# State and logging
STATE_FILE="${STATE_FILE:-.component-refactor-state}"
LOG_FILE="${LOG_FILE:-./component-refactor.log}"
REFACTOR_GUIDE="${REFACTOR_GUIDE:-$VIEWS_DIR/refactoring_guide.md}"
BACKUP_DIR="${BACKUP_DIR:-./views_backup_$(date +%Y%m%d_%H%M%S)}"

# Limits
MAX_FIX_ATTEMPTS=5
BUILD_TIMEOUT=120

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
WHITE='\033[1;37m'
BOLD='\033[1m'
DIM='\033[2m'
NC='\033[0m'

# =============================================================================
# Basecoat Component Classes Reference
# =============================================================================

read -r -d '' BASECOAT_CLASSES <<'EOF'
# Cards
card card-header card-title card-description card-content card-footer

# Buttons
btn btn-outline btn-secondary btn-destructive btn-ghost btn-link btn-icon

# Form Fields
field field-label field-input field-description field-error

# Input, Textarea, Select, Checkbox, Radio, Switch
input textarea select checkbox checkbox-indicator radio switch switch-thumb

# Badge, Alert, Avatar
badge badge-outline badge-secondary badge-destructive
alert alert-title alert-description
avatar avatar-image avatar-fallback

# Dialog, Dropdown, Popover, Tabs, Toast
dialog dialog-overlay dialog-content dialog-header dialog-footer
dropdown-menu dropdown-menu-trigger dropdown-menu-content dropdown-menu-item
popover popover-trigger popover-content
tabs tabs-list tabs-trigger tabs-content
toast toast-title toast-description toaster

# Sidebar, Command, Table, Accordion
sidebar sidebar-header sidebar-content sidebar-nav sidebar-nav-item
command command-input command-list command-item
table table-header table-body table-row table-head table-cell
accordion accordion-item accordion-trigger accordion-content

# Separator, Skeleton, Progress, Spinner
separator skeleton progress progress-indicator spinner

# Tooltip, Sheet, Scroll Area
tooltip tooltip-trigger tooltip-content
sheet sheet-content sheet-header sheet-footer
scroll-area scroll-area-viewport scroll-area-scrollbar
EOF

# Components that need custom CSS (not in Basecoat)
CUSTOM_COMPONENTS=(
  "icon" "kbd" "code" "link" "divider" "heading" "text"
  "status-badge" "quantity-selector" "currency-input" "date-range-picker"
  "tag-input" "file-upload" "breadcrumbs" "filter-chip" "pagination"
  "empty-state" "input-group" "media-object" "stat-display" "list-item"
  "data-table" "stats-card" "wizard" "filter-panel" "chart-widget"
  "user-menu" "notification-list" "calendar-widget" "comment-thread"
  "file-explorer" "invoice-card" "product-card"
  "layout-dashboard" "layout-master-detail" "layout-form" "layout-auth"
  "layout-error" "layout-wizard" "layout-kanban" "layout-report" "layout-settings"
)

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
# State Management
# =============================================================================

declare -A TASK_STATUS
declare -A TASK_DESC
declare -A TASK_PHASE
declare -a TASK_ORDER

init_state() {
  if [[ -f "$STATE_FILE" ]]; then
    source "$STATE_FILE"
  fi
}

save_state() {
  {
    echo "# Component Refactor State - $(date)"
    echo "LAST_TASK=\"$LAST_TASK\""
    echo "LAST_STATUS=\"$LAST_STATUS\""
    echo "START_TIME=\"$START_TIME\""
    echo ""
    for task in "${TASK_ORDER[@]}"; do
      echo "TASK_STATUS[$task]=\"${TASK_STATUS[$task]:-pending}\""
    done
  } >"$STATE_FILE"
}

set_task_status() {
  local task="$1"
  local status="$2"
  TASK_STATUS[$task]="$status"
  LAST_TASK="$task"
  LAST_STATUS="$status"
  save_state
}

get_task_status() {
  echo "${TASK_STATUS[$1]:-pending}"
}

show_status() {
  log_header "Component Refactoring Progress"

  local completed=0 failed=0 pending=0
  local total=${#TASK_ORDER[@]}
  local current_phase=""

  for task in "${TASK_ORDER[@]}"; do
    local status="${TASK_STATUS[$task]:-pending}"
    local desc="${TASK_DESC[$task]}"
    local phase="${TASK_PHASE[$task]}"

    if [[ "$phase" != "$current_phase" ]]; then
      current_phase="$phase"
      echo ""
      echo -e "${BOLD}$phase${NC}"
    fi

    case "$status" in
    completed)
      echo -e "  ${GREEN}✓${NC} $task: $desc"
      ((completed++))
      ;;
    failed)
      echo -e "  ${RED}✗${NC} $task: $desc"
      ((failed++))
      ;;
    running) echo -e "  ${YELLOW}►${NC} $task: $desc" ;;
    skipped) echo -e "  ${CYAN}○${NC} $task: $desc (skipped)" ;;
    *)
      echo -e "  ${BLUE}·${NC} $task: $desc"
      ((pending++))
      ;;
    esac
  done

  echo ""
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo "Progress: $completed/$total completed, $failed failed, $pending pending"
  [[ -n "$LAST_TASK" ]] && echo "Last task: $LAST_TASK ($LAST_STATUS)"
  echo ""
  echo "Demo server: ./component-refactor.sh --demo"
}

reset_state() {
  rm -f "$STATE_FILE"
  log_info "State reset. Starting fresh."
}

# =============================================================================
# Task Registry
# =============================================================================

register_task() {
  local id="$1" phase="$2" desc="$3"
  TASK_ORDER+=("$id")
  TASK_DESC[$id]="$desc"
  TASK_PHASE[$id]="$phase"
  TASK_STATUS[$id]="${TASK_STATUS[$id]:-pending}"
}

register_all_tasks() {
  # PHASE 0: Demo Setup
  register_task "T0.1" "PHASE 0: Demo Setup" "Create demo server (server.go)"
  register_task "T0.2" "PHASE 0: Demo Setup" "Create demo page template (demo.templ)"
  register_task "T0.3" "PHASE 0: Demo Setup" "Create demo section templates"

  # PHASE 1: Analysis
  register_task "T1.1" "PHASE 1: Analysis" "Parse basecoat.css for component classes"
  register_task "T1.2" "PHASE 1: Analysis" "Scan Basecoat JS for behavior patterns"
  register_task "T1.3" "PHASE 1: Analysis" "Inventory existing templ components"
  register_task "T1.4" "PHASE 1: Analysis" "Map Basecoat vs Custom components"
  register_task "T1.5" "PHASE 1: Analysis" "Detect anti-patterns"
  register_task "T1.6" "PHASE 1: Analysis" "Generate analysis report"

  # PHASE 2: Strategy
  register_task "T2.1" "PHASE 2: Strategy" "Define component API standards"
  register_task "T2.2" "PHASE 2: Strategy" "Create shared types (types.go)"
  register_task "T2.3" "PHASE 2: Strategy" "Update helpers (helpers.go)"
  register_task "T2.4" "PHASE 2: Strategy" "Define refactoring priority"
  register_task "T2.5" "PHASE 2: Strategy" "Generate strategy report"

  # PHASE 3: Custom CSS
  register_task "T3.1" "PHASE 3: Custom CSS" "Create CSS directory structure"
  register_task "T3.2" "PHASE 3: Custom CSS" "Generate icon.css"
  register_task "T3.3" "PHASE 3: Custom CSS" "Generate kbd.css"
  register_task "T3.4" "PHASE 3: Custom CSS" "Generate code.css"
  register_task "T3.5" "PHASE 3: Custom CSS" "Generate link.css"
  register_task "T3.6" "PHASE 3: Custom CSS" "Generate divider.css"
  register_task "T3.7" "PHASE 3: Custom CSS" "Generate status-badge.css"
  register_task "T3.8" "PHASE 3: Custom CSS" "Generate pagination.css"
  register_task "T3.9" "PHASE 3: Custom CSS" "Generate empty-state.css"
  register_task "T3.10" "PHASE 3: Custom CSS" "Generate filter-chip.css"
  register_task "T3.11" "PHASE 3: Custom CSS" "Generate input-group.css"
  register_task "T3.12" "PHASE 3: Custom CSS" "Generate data-table.css"
  register_task "T3.13" "PHASE 3: Custom CSS" "Generate stats-card.css"
  register_task "T3.14" "PHASE 3: Custom CSS" "Generate wizard.css"
  register_task "T3.15" "PHASE 3: Custom CSS" "Generate breadcrumbs.css"
  register_task "T3.16" "PHASE 3: Custom CSS" "Generate layout CSS files"
  register_task "T3.17" "PHASE 3: Custom CSS" "Create components index.css"

  # PHASE 4: Atoms (24) + Demo updates
  register_task "T4.1" "PHASE 4: Atoms" "Refactor button.templ"
  register_task "T4.2" "PHASE 4: Atoms" "Refactor input.templ"
  register_task "T4.3" "PHASE 4: Atoms" "Refactor label.templ"
  register_task "T4.4" "PHASE 4: Atoms" "Refactor icon.templ"
  register_task "T4.5" "PHASE 4: Atoms" "Refactor badge.templ"
  register_task "T4.6" "PHASE 4: Atoms" "Refactor spinner.templ"
  register_task "T4.7" "PHASE 4: Atoms" "Refactor separator.templ"
  register_task "T4.8" "PHASE 4: Atoms" "Refactor avatar.templ"
  register_task "T4.9" "PHASE 4: Atoms" "Refactor checkbox.templ"
  register_task "T4.10" "PHASE 4: Atoms" "Refactor radio.templ"
  register_task "T4.11" "PHASE 4: Atoms" "Refactor select.templ"
  register_task "T4.12" "PHASE 4: Atoms" "Refactor textarea.templ"
  register_task "T4.13" "PHASE 4: Atoms" "Refactor switch.templ"
  register_task "T4.14" "PHASE 4: Atoms" "Create link.templ"
  register_task "T4.15" "PHASE 4: Atoms" "Create image.templ"
  register_task "T4.16" "PHASE 4: Atoms" "Refactor tooltip.templ"
  register_task "T4.17" "PHASE 4: Atoms" "Refactor progress.templ"
  register_task "T4.18" "PHASE 4: Atoms" "Refactor skeleton.templ"
  register_task "T4.19" "PHASE 4: Atoms" "Refactor alert.templ"
  register_task "T4.20" "PHASE 4: Atoms" "Create heading.templ"
  register_task "T4.21" "PHASE 4: Atoms" "Create text.templ"
  register_task "T4.22" "PHASE 4: Atoms" "Create code.templ"
  register_task "T4.23" "PHASE 4: Atoms" "Refactor kbd.templ"
  register_task "T4.24" "PHASE 4: Atoms" "Refactor tags.templ"
  register_task "T4.25" "PHASE 4: Atoms" "Update demo: Atoms section"

  # PHASE 5: Molecules (20) + Demo updates
  register_task "T5.1" "PHASE 5: Molecules" "Refactor form_field.templ"
  register_task "T5.2" "PHASE 5: Molecules" "Refactor searchbox.templ"
  register_task "T5.3" "PHASE 5: Molecules" "Create stat_display.templ"
  register_task "T5.4" "PHASE 5: Molecules" "Create quantity_selector.templ"
  register_task "T5.5" "PHASE 5: Molecules" "Create media_object.templ"
  register_task "T5.6" "PHASE 5: Molecules" "Create tag_input.templ"
  register_task "T5.7" "PHASE 5: Molecules" "Refactor date_picker.templ"
  register_task "T5.8" "PHASE 5: Molecules" "Create currency_input.templ"
  register_task "T5.9" "PHASE 5: Molecules" "Create file_upload.templ"
  register_task "T5.10" "PHASE 5: Molecules" "Create breadcrumb_item.templ"
  register_task "T5.11" "PHASE 5: Molecules" "Create status_badge.templ"
  register_task "T5.12" "PHASE 5: Molecules" "Create pagination.templ"
  register_task "T5.13" "PHASE 5: Molecules" "Create empty_state.templ"
  register_task "T5.14" "PHASE 5: Molecules" "Create filter_chip.templ"
  register_task "T5.15" "PHASE 5: Molecules" "Create action_buttons.templ"
  register_task "T5.16" "PHASE 5: Molecules" "Refactor card.templ"
  register_task "T5.17" "PHASE 5: Molecules" "Create list_item.templ"
  register_task "T5.18" "PHASE 5: Molecules" "Refactor tabs.templ"
  register_task "T5.19" "PHASE 5: Molecules" "Refactor dropdown_menu.templ"
  register_task "T5.20" "PHASE 5: Molecules" "Create input_group.templ"
  register_task "T5.21" "PHASE 5: Molecules" "Update demo: Molecules section"

  # PHASE 6: Organisms (18) + Demo updates
  register_task "T6.1" "PHASE 6: Organisms" "Refactor datatable.templ"
  register_task "T6.2" "PHASE 6: Organisms" "Refactor navigation.templ"
  register_task "T6.3" "PHASE 6: Organisms" "Refactor sidebar.templ"
  register_task "T6.4" "PHASE 6: Organisms" "Create invoice_card.templ"
  register_task "T6.5" "PHASE 6: Organisms" "Create product_card.templ"
  register_task "T6.6" "PHASE 6: Organisms" "Refactor form.templ"
  register_task "T6.7" "PHASE 6: Organisms" "Create filter_panel.templ"
  register_task "T6.8" "PHASE 6: Organisms" "Create chart_widget.templ"
  register_task "T6.9" "PHASE 6: Organisms" "Refactor dialog.templ (modal)"
  register_task "T6.10" "PHASE 6: Organisms" "Create breadcrumbs.templ"
  register_task "T6.11" "PHASE 6: Organisms" "Create user_menu.templ"
  register_task "T6.12" "PHASE 6: Organisms" "Create notification_list.templ"
  register_task "T6.13" "PHASE 6: Organisms" "Refactor popover.templ"
  register_task "T6.14" "PHASE 6: Organisms" "Create calendar_widget.templ"
  register_task "T6.15" "PHASE 6: Organisms" "Create comment_thread.templ"
  register_task "T6.16" "PHASE 6: Organisms" "Create file_explorer.templ"
  register_task "T6.17" "PHASE 6: Organisms" "Create stats_card.templ"
  register_task "T6.18" "PHASE 6: Organisms" "Create wizard.templ"
  register_task "T6.19" "PHASE 6: Organisms" "Update demo: Organisms section"

  # PHASE 7: Templates (12)
  register_task "T7.1" "PHASE 7: Templates" "Refactor base_layout.templ"
  register_task "T7.2" "PHASE 7: Templates" "Create master_detail_layout.templ"
  register_task "T7.3" "PHASE 7: Templates" "Refactor dashboard_layout.templ"
  register_task "T7.4" "PHASE 7: Templates" "Create form_layout.templ"
  register_task "T7.5" "PHASE 7: Templates" "Create report_layout.templ"
  register_task "T7.6" "PHASE 7: Templates" "Create content_layout.templ"
  register_task "T7.7" "PHASE 7: Templates" "Create auth_layout.templ"
  register_task "T7.8" "PHASE 7: Templates" "Create split_view_layout.templ"
  register_task "T7.9" "PHASE 7: Templates" "Create settings_layout.templ"
  register_task "T7.10" "PHASE 7: Templates" "Create wizard_layout.templ"
  register_task "T7.11" "PHASE 7: Templates" "Create error_layout.templ"
  register_task "T7.12" "PHASE 7: Templates" "Create kanban_layout.templ"
  register_task "T7.13" "PHASE 7: Templates" "Update demo: Templates section"

  # PHASE 8: Pages (15)
  register_task "T8.1" "PHASE 8: Pages" "Create login_page.templ"
  register_task "T8.2" "PHASE 8: Pages" "Create dashboard_home_page.templ"
  register_task "T8.3" "PHASE 8: Pages" "Create invoice_list_page.templ"
  register_task "T8.4" "PHASE 8: Pages" "Create invoice_detail_page.templ"
  register_task "T8.5" "PHASE 8: Pages" "Create invoice_create_page.templ"
  register_task "T8.6" "PHASE 8: Pages" "Create customer_list_page.templ"
  register_task "T8.7" "PHASE 8: Pages" "Create product_catalog_page.templ"
  register_task "T8.8" "PHASE 8: Pages" "Create product_detail_page.templ"
  register_task "T8.9" "PHASE 8: Pages" "Create settings_page.templ"
  register_task "T8.10" "PHASE 8: Pages" "Create user_profile_page.templ"
  register_task "T8.11" "PHASE 8: Pages" "Create reports_index_page.templ"
  register_task "T8.12" "PHASE 8: Pages" "Create sales_report_page.templ"
  register_task "T8.13" "PHASE 8: Pages" "Create search_results_page.templ"
  register_task "T8.14" "PHASE 8: Pages" "Create not_found_page.templ"
  register_task "T8.15" "PHASE 8: Pages" "Create station_dashboard_page.templ"

  # PHASE 9: Validation
  register_task "T9.1" "PHASE 9: Validation" "Check templ compilation"
  register_task "T9.2" "PHASE 9: Validation" "Verify Basecoat classes exist"
  register_task "T9.3" "PHASE 9: Validation" "Verify custom CSS complete"
  register_task "T9.4" "PHASE 9: Validation" "Scan for remaining anti-patterns"
  register_task "T9.5" "PHASE 9: Validation" "Test HTMX integration"
  register_task "T9.6" "PHASE 9: Validation" "Test Alpine.js integration"
  register_task "T9.7" "PHASE 9: Validation" "Finalize component gallery demo"
  register_task "T9.8" "PHASE 9: Validation" "Generate final report"
}

# =============================================================================
# Build & Validation
# =============================================================================

run_templ_generate() {
  if ! command -v templ &>/dev/null; then
    log_warn "templ not found - skipping"
    return 0
  fi
  timeout "$BUILD_TIMEOUT" templ generate 2>&1
}

run_go_build() {
  timeout "$BUILD_TIMEOUT" go build ./... 2>&1
}

detect_anti_patterns() {
  local dir="${1:-$COMPONENTS_DIR}"
  local issues=""

  # ClassName prop detection (the actual anti-pattern)
  local classname=$(grep -rln 'ClassName\s*string' "$dir" --include="*.templ" --include="*.go" 2>/dev/null || true)
  [[ -n "$classname" ]] && issues+="ClassName props found in:\n$classname\n\n"

  echo -e "$issues"
}

fix_build_errors() {
  local attempt=1

  while [[ $attempt -le $MAX_FIX_ATTEMPTS ]]; do
    log_info "Build attempt $attempt/$MAX_FIX_ATTEMPTS..."

    if run_templ_generate && run_go_build; then
      log_success "Build passed!"
      return 0
    fi

    local errors=$(templ generate 2>&1 | grep -E "(error|Error)" | head -20)
    [[ -z "$errors" ]] && errors=$(go build ./... 2>&1 | grep -E "^\./" | head -20)

    if [[ -z "$errors" ]]; then
      log_error "Build failed but couldn't extract errors"
      return 1
    fi

    log_fix "Attempting fix (attempt $attempt)..."

    claude --dangerously-skip-permissions "Fix these build errors:

$errors

Rules:
1. Fix syntax errors
2. Fix undefined references  
3. Fix type mismatches
4. The code MUST compile."

    ((attempt++))
    sleep 2
  done

  log_error "Could not fix after $MAX_FIX_ATTEMPTS attempts"
  return 1
}

ensure_builds() {
  if ! run_templ_generate 2>/dev/null; then
    log_warn "Build failed, attempting fixes..."
    fix_build_errors || return 1
  fi
  return 0
}

# =============================================================================
# Demo Server
# =============================================================================

start_demo_server() {
  log_header "Starting Demo Server"

  if [[ ! -f "$DEMO_SERVER" ]]; then
    log_error "Demo server not found: $DEMO_SERVER"
    log_info "Run: $0 --task T0.1 to create it"
    return 1
  fi

  # Generate templ files first
  log_info "Generating templ files..."
  run_templ_generate

  # Start server
  log_info "Starting demo server on http://localhost:$DEMO_PORT"
  log_info "Press Ctrl+C to stop"
  echo ""

  cd "$DEMO_DIR" && go run server.go
}

# =============================================================================
# Task Execution
# =============================================================================

execute_task() {
  local task_id="$1"
  local task_func="task_${task_id//./_}"

  log_task "Executing $task_id: ${TASK_DESC[$task_id]}"
  set_task_status "$task_id" "running"

  if type "$task_func" &>/dev/null; then
    if $task_func; then
      if ensure_builds; then
        set_task_status "$task_id" "completed"
        log_success "Task $task_id completed"

        if git rev-parse --git-dir &>/dev/null 2>&1; then
          git add -A
          git commit -m "refactor(ui): $task_id - ${TASK_DESC[$task_id]}" 2>/dev/null || true
        fi
        return 0
      fi
    fi
  else
    log_error "Task function $task_func not found"
  fi

  set_task_status "$task_id" "failed"
  log_error "Task $task_id failed"
  return 1
}

run_all_tasks() {
  for task in "${TASK_ORDER[@]}"; do
    local status=$(get_task_status "$task")
    if [[ "$status" != "completed" ]]; then
      if ! execute_task "$task"; then
        log_error "Failed at task: $task"
        echo "To resume: $0 --resume"
        return 1
      fi
    else
      log_info "Skipping completed: $task"
    fi
  done

  log_header "ALL TASKS COMPLETED"
  show_status

  echo ""
  log_success "View components: $0 --demo"
}

run_from_task() {
  local start_task="$1"
  local started=false

  for task in "${TASK_ORDER[@]}"; do
    [[ "$task" == "$start_task" ]] && started=true
    if $started; then
      local status=$(get_task_status "$task")
      if [[ "$status" != "completed" ]]; then
        execute_task "$task" || return 1
      fi
    fi
  done
}

run_phase() {
  local phase_num="$1"
  local phase_prefix="T${phase_num}."

  log_header "Running Phase $phase_num"

  for task in "${TASK_ORDER[@]}"; do
    if [[ "$task" == $phase_prefix* ]]; then
      local status=$(get_task_status "$task")
      [[ "$status" != "completed" ]] && execute_task "$task"
    fi
  done

  log_success "Phase $phase_num complete!"
  log_info "View components: $0 --demo"
}

resume_tasks() {
  local resume_from=""
  for task in "${TASK_ORDER[@]}"; do
    local status=$(get_task_status "$task")
    if [[ "$status" == "failed" || "$status" == "running" || "$status" == "pending" ]]; then
      resume_from="$task"
      break
    fi
  done

  if [[ -z "$resume_from" ]]; then
    log_success "All tasks already completed!"
    return 0
  fi

  log_info "Resuming from: $resume_from"
  run_from_task "$resume_from"
}

# =============================================================================
# PHASE 0: Demo Setup Tasks
# =============================================================================

task_T0_1() {
  log_info "Creating demo server..."

  mkdir -p "$DEMO_DIR"

  cat >"$DEMO_SERVER" <<'GOEOF'
package main

import (
	"context"
	"log"
	"os"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Awo ERP Component Demo",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// Serve static assets (CSS, JS)
	app.Static("/static", "../static")

	// Main demo page (kitchen sink)
	app.Get("/", func(c *fiber.Ctx) error {
		return render(c, DemoPage())
	})

	// Section routes for quick navigation
	app.Get("/atoms", func(c *fiber.Ctx) error {
		return render(c, AtomsSection())
	})

	app.Get("/molecules", func(c *fiber.Ctx) error {
		return render(c, MoleculesSection())
	})

	app.Get("/organisms", func(c *fiber.Ctx) error {
		return render(c, OrganismsSection())
	})

	app.Get("/templates", func(c *fiber.Ctx) error {
		return render(c, TemplatesSection())
	})

	// Get port from environment or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Demo server starting on http://localhost:%s", port)
	log.Printf("📦 Kitchen Sink: http://localhost:%s/", port)
	log.Printf("⚛️  Atoms: http://localhost:%s/atoms", port)
	log.Printf("🧬 Molecules: http://localhost:%s/molecules", port)
	log.Printf("🦠 Organisms: http://localhost:%s/organisms", port)
	log.Fatal(app.Listen(":" + port))
}

// render helper for templ components
func render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(context.Background(), c.Response().BodyWriter())
}
GOEOF

  log_success "Created: $DEMO_SERVER"
}

task_T0_2() {
  log_info "Creating demo page template..."

  cat >"$DEMO_PAGE" <<'TEMPLEOF'
package main

// DemoPage renders the full kitchen-sink style component gallery
templ DemoPage() {
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<title>Awo ERP - Component Gallery</title>
		<link rel="stylesheet" href="/static/css/basecoat.css"/>
		<link rel="stylesheet" href="/static/css/components/index.css"/>
		<script src="/static/js/htmx.min.js" defer></script>
		<script src="/static/js/alpine.min.js" defer></script>
	</head>
	<body class="min-h-screen bg-background">
		<!-- Navigation -->
		<nav class="sticky top-0 z-50 border-b bg-background/95 backdrop-blur">
			<div class="container mx-auto px-4">
				<div class="flex h-16 items-center justify-between">
					<div class="flex items-center gap-6">
						<a href="/" class="text-xl font-bold">Awo ERP</a>
						<span class="text-muted-foreground">Component Gallery</span>
					</div>
					<div class="flex items-center gap-4">
						<a href="#atoms" class="text-sm hover:text-primary">Atoms</a>
						<a href="#molecules" class="text-sm hover:text-primary">Molecules</a>
						<a href="#organisms" class="text-sm hover:text-primary">Organisms</a>
						<a href="#templates" class="text-sm hover:text-primary">Templates</a>
					</div>
				</div>
			</div>
		</nav>

		<!-- Hero -->
		<header class="border-b bg-muted/30 py-12">
			<div class="container mx-auto px-4 text-center">
				<h1 class="text-4xl font-bold tracking-tight">Component Gallery</h1>
				<p class="mt-4 text-lg text-muted-foreground">
					Kitchen sink demo of all UI components following Basecoat + Atomic Design patterns
				</p>
				<div class="mt-6 flex justify-center gap-4">
					<span class="badge">Basecoat UI</span>
					<span class="badge badge-secondary">HTMX</span>
					<span class="badge badge-outline">Alpine.js</span>
					<span class="badge badge-outline">Templ</span>
				</div>
			</div>
		</header>

		<!-- Main Content -->
		<main class="container mx-auto px-4 py-8">
			@AtomsSection()
			@MoleculesSection()
			@OrganismsSection()
			@TemplatesSection()
		</main>

		<!-- Footer -->
		<footer class="border-t py-8 text-center text-sm text-muted-foreground">
			<p>Awo ERP Component Library • Built with Basecoat UI + Templ</p>
		</footer>
	</body>
	</html>
}

// AtomsSection displays all atomic components
templ AtomsSection() {
	<section id="atoms" class="py-12">
		<h2 class="text-2xl font-bold mb-8 pb-4 border-b">Atoms</h2>
		<p class="text-muted-foreground mb-8">
			Basic building blocks: buttons, inputs, badges, icons, etc.
		</p>
		
		<div class="grid gap-12">
			<!-- Buttons -->
			@DemoCard("Buttons", "btn, btn-outline, btn-destructive, btn-ghost, btn-link") {
				<div class="flex flex-wrap gap-4">
					<button class="btn">Primary</button>
					<button class="btn btn-secondary">Secondary</button>
					<button class="btn btn-outline">Outline</button>
					<button class="btn btn-destructive">Destructive</button>
					<button class="btn btn-ghost">Ghost</button>
					<button class="btn btn-link">Link</button>
				</div>
				<div class="flex flex-wrap gap-4 mt-4">
					<button class="btn btn-sm">Small</button>
					<button class="btn">Default</button>
					<button class="btn btn-lg">Large</button>
					<button class="btn" disabled>Disabled</button>
				</div>
			}

			<!-- Inputs -->
			@DemoCard("Inputs", "input, field, field-label, field-input, field-description") {
				<div class="grid grid-cols-2 gap-4 max-w-2xl">
					<div class="field">
						<label class="field-label">Email</label>
						<input class="field-input" type="email" placeholder="name@example.com"/>
					</div>
					<div class="field">
						<label class="field-label">Password</label>
						<input class="field-input" type="password" placeholder="••••••••"/>
					</div>
					<div class="field col-span-2">
						<label class="field-label">With Description</label>
						<input class="field-input" type="text" placeholder="Enter value"/>
						<p class="field-description">This is a helpful description.</p>
					</div>
					<div class="field col-span-2">
						<label class="field-label">With Error</label>
						<input class="field-input border-destructive" type="text" value="Invalid value"/>
						<p class="text-sm text-destructive">This field has an error.</p>
					</div>
				</div>
			}

			<!-- Badges -->
			@DemoCard("Badges", "badge, badge-secondary, badge-outline, badge-destructive") {
				<div class="flex flex-wrap gap-4">
					<span class="badge">Default</span>
					<span class="badge badge-secondary">Secondary</span>
					<span class="badge badge-outline">Outline</span>
					<span class="badge badge-destructive">Destructive</span>
				</div>
			}

			<!-- Checkbox & Radio -->
			@DemoCard("Checkbox & Radio", "checkbox, radio, switch") {
				<div class="flex flex-wrap gap-8">
					<div class="flex items-center space-x-2">
						<input type="checkbox" id="c1" class="checkbox"/>
						<label for="c1" class="checkbox-label">Accept terms</label>
					</div>
					<div class="flex items-center space-x-2">
						<input type="checkbox" id="c2" class="checkbox" checked/>
						<label for="c2" class="checkbox-label">Checked</label>
					</div>
					<div class="flex items-center space-x-2">
						<input type="radio" name="demo" id="r1" class="radio"/>
						<label for="r1" class="radio-label">Option A</label>
					</div>
					<div class="flex items-center space-x-2">
						<input type="radio" name="demo" id="r2" class="radio" checked/>
						<label for="r2" class="radio-label">Option B</label>
					</div>
				</div>
			}

			<!-- Alerts -->
			@DemoCard("Alerts", "alert, alert-title, alert-description") {
				<div class="grid gap-4 max-w-xl">
					<div class="alert">
						<h5 class="alert-title">Default Alert</h5>
						<p class="alert-description">This is a default alert message.</p>
					</div>
					<div class="alert alert-destructive">
						<h5 class="alert-title">Error</h5>
						<p class="alert-description">Something went wrong. Please try again.</p>
					</div>
				</div>
			}

			<!-- Progress & Spinner -->
			@DemoCard("Progress & Spinner", "progress, progress-indicator, spinner") {
				<div class="grid gap-4 max-w-xl">
					<div class="progress">
						<div class="progress-indicator" style="width: 60%"></div>
					</div>
					<div class="flex items-center gap-4">
						<span class="spinner"></span>
						<span class="text-sm text-muted-foreground">Loading...</span>
					</div>
				</div>
			}

			<!-- Skeleton -->
			@DemoCard("Skeleton", "skeleton") {
				<div class="flex items-center space-x-4">
					<div class="skeleton h-12 w-12 rounded-full"></div>
					<div class="space-y-2">
						<div class="skeleton h-4 w-[250px]"></div>
						<div class="skeleton h-4 w-[200px]"></div>
					</div>
				</div>
			}

			<!-- Avatar -->
			@DemoCard("Avatar", "avatar, avatar-image, avatar-fallback") {
				<div class="flex items-center gap-4">
					<div class="avatar">
						<span class="avatar-fallback">JD</span>
					</div>
					<div class="avatar">
						<span class="avatar-fallback">AB</span>
					</div>
					<div class="avatar">
						<span class="avatar-fallback">CD</span>
					</div>
				</div>
			}

			<!-- Separator -->
			@DemoCard("Separator", "separator") {
				<div class="max-w-xl">
					<p class="text-sm">Content above</p>
					<div class="separator my-4"></div>
					<p class="text-sm">Content below</p>
				</div>
			}
		</div>
	</section>
}

// MoleculesSection displays molecule components
templ MoleculesSection() {
	<section id="molecules" class="py-12">
		<h2 class="text-2xl font-bold mb-8 pb-4 border-b">Molecules</h2>
		<p class="text-muted-foreground mb-8">
			Combined atoms: cards, tabs, dropdowns, form fields, etc.
		</p>
		
		<div class="grid gap-12">
			<!-- Card -->
			@DemoCard("Card", "card, card-header, card-title, card-description, card-content, card-footer") {
				<div class="grid grid-cols-2 gap-4 max-w-3xl">
					<div class="card">
						<div class="card-header">
							<h3 class="card-title">Card Title</h3>
							<p class="card-description">Card description goes here.</p>
						</div>
						<div class="card-content">
							<p class="text-sm">This is the card content area.</p>
						</div>
						<div class="card-footer flex justify-end gap-2">
							<button class="btn btn-outline btn-sm">Cancel</button>
							<button class="btn btn-sm">Save</button>
						</div>
					</div>
					<div class="card">
						<div class="card-header">
							<h3 class="card-title">Simple Card</h3>
						</div>
						<div class="card-content">
							<p class="text-sm text-muted-foreground">A simpler card without footer.</p>
						</div>
					</div>
				</div>
			}

			<!-- Tabs -->
			@DemoCard("Tabs", "tabs, tabs-list, tabs-trigger, tabs-content") {
				<div class="tabs max-w-xl" x-data="{ tab: 'account' }">
					<div class="tabs-list">
						<button class="tabs-trigger" :class="tab === 'account' && 'tabs-trigger-active'" @click="tab = 'account'">Account</button>
						<button class="tabs-trigger" :class="tab === 'password' && 'tabs-trigger-active'" @click="tab = 'password'">Password</button>
						<button class="tabs-trigger" :class="tab === 'settings' && 'tabs-trigger-active'" @click="tab = 'settings'">Settings</button>
					</div>
					<div class="tabs-content p-4" x-show="tab === 'account'">
						<p class="text-sm text-muted-foreground">Account settings content here.</p>
					</div>
					<div class="tabs-content p-4" x-show="tab === 'password'">
						<p class="text-sm text-muted-foreground">Password settings content here.</p>
					</div>
					<div class="tabs-content p-4" x-show="tab === 'settings'">
						<p class="text-sm text-muted-foreground">General settings content here.</p>
					</div>
				</div>
			}

			<!-- Status Badge (Custom) -->
			@DemoCard("Status Badge", "status-badge, status-badge-{status}") {
				<div class="flex flex-wrap gap-4">
					<span class="status-badge status-badge-active">
						<span class="status-badge-dot"></span>
						Active
					</span>
					<span class="status-badge status-badge-inactive">Inactive</span>
					<span class="status-badge status-badge-pending">Pending</span>
					<span class="status-badge status-badge-approved">Approved</span>
					<span class="status-badge status-badge-rejected">Rejected</span>
					<span class="status-badge status-badge-draft">Draft</span>
				</div>
			}

			<!-- Pagination (Custom) -->
			@DemoCard("Pagination", "pagination, pagination-button, pagination-ellipsis") {
				<div class="pagination">
					<div class="pagination-info">Showing 1-10 of 100</div>
					<div class="pagination-controls">
						<button class="pagination-button">&laquo;</button>
						<button class="pagination-button">&lsaquo;</button>
						<button class="pagination-button pagination-button-active">1</button>
						<button class="pagination-button">2</button>
						<button class="pagination-button">3</button>
						<span class="pagination-ellipsis">...</span>
						<button class="pagination-button">10</button>
						<button class="pagination-button">&rsaquo;</button>
						<button class="pagination-button">&raquo;</button>
					</div>
				</div>
			}

			<!-- Empty State (Custom) -->
			@DemoCard("Empty State", "empty-state, empty-state-icon, empty-state-title, empty-state-description") {
				<div class="empty-state max-w-md mx-auto">
					<div class="empty-state-icon">📭</div>
					<h3 class="empty-state-title">No results found</h3>
					<p class="empty-state-description">Try adjusting your search or filters to find what you're looking for.</p>
					<div class="empty-state-actions">
						<button class="btn">Clear Filters</button>
					</div>
				</div>
			}
		</div>
	</section>
}

// OrganismsSection displays organism components  
templ OrganismsSection() {
	<section id="organisms" class="py-12">
		<h2 class="text-2xl font-bold mb-8 pb-4 border-b">Organisms</h2>
		<p class="text-muted-foreground mb-8">
			Complex components: data tables, forms, dialogs, sidebars, etc.
		</p>
		
		<div class="grid gap-12">
			<!-- Data Table (Custom) -->
			@DemoCard("Data Table", "data-table, data-table-header, data-table-row, data-table-cell") {
				<div class="data-table-wrapper">
					<table class="data-table">
						<thead class="data-table-header">
							<tr>
								<th class="data-table-head">Invoice</th>
								<th class="data-table-head">Status</th>
								<th class="data-table-head">Customer</th>
								<th class="data-table-head text-right">Amount</th>
							</tr>
						</thead>
						<tbody>
							<tr class="data-table-row">
								<td class="data-table-cell font-medium">INV-001</td>
								<td class="data-table-cell"><span class="status-badge status-badge-approved">Paid</span></td>
								<td class="data-table-cell">John Doe</td>
								<td class="data-table-cell text-right">$250.00</td>
							</tr>
							<tr class="data-table-row">
								<td class="data-table-cell font-medium">INV-002</td>
								<td class="data-table-cell"><span class="status-badge status-badge-pending">Pending</span></td>
								<td class="data-table-cell">Jane Smith</td>
								<td class="data-table-cell text-right">$150.00</td>
							</tr>
							<tr class="data-table-row">
								<td class="data-table-cell font-medium">INV-003</td>
								<td class="data-table-cell"><span class="status-badge status-badge-rejected">Overdue</span></td>
								<td class="data-table-cell">Bob Wilson</td>
								<td class="data-table-cell text-right">$350.00</td>
							</tr>
						</tbody>
					</table>
				</div>
			}

			<!-- Stats Card (Custom) -->
			@DemoCard("Stats Card", "stats-card, stats-card-value, stats-card-label, stats-card-trend") {
				<div class="grid grid-cols-3 gap-4 max-w-3xl">
					<div class="stats-card">
						<div class="stats-card-icon">💰</div>
						<div class="stats-card-value">$45,231</div>
						<div class="stats-card-label">Total Revenue</div>
						<div class="stats-card-trend stats-card-trend-up">+12.5%</div>
					</div>
					<div class="stats-card">
						<div class="stats-card-icon">📦</div>
						<div class="stats-card-value">2,345</div>
						<div class="stats-card-label">Orders</div>
						<div class="stats-card-trend stats-card-trend-up">+8.2%</div>
					</div>
					<div class="stats-card">
						<div class="stats-card-icon">👥</div>
						<div class="stats-card-value">1,234</div>
						<div class="stats-card-label">Customers</div>
						<div class="stats-card-trend stats-card-trend-down">-2.4%</div>
					</div>
				</div>
			}

			<!-- Dialog/Modal -->
			@DemoCard("Dialog", "dialog, dialog-content, dialog-header, dialog-title, dialog-footer") {
				<div x-data="{ open: false }">
					<button class="btn" @click="open = true">Open Dialog</button>
					<div class="dialog" x-show="open" x-cloak>
						<div class="dialog-overlay" @click="open = false"></div>
						<div class="dialog-content">
							<div class="dialog-header">
								<h3 class="dialog-title">Dialog Title</h3>
								<p class="dialog-description">This is a dialog description.</p>
							</div>
							<div class="p-6">
								<p class="text-sm text-muted-foreground">Dialog content goes here.</p>
							</div>
							<div class="dialog-footer">
								<button class="btn btn-outline" @click="open = false">Cancel</button>
								<button class="btn" @click="open = false">Confirm</button>
							</div>
						</div>
					</div>
				</div>
			}

			<!-- Breadcrumbs (Custom) -->
			@DemoCard("Breadcrumbs", "breadcrumbs, breadcrumbs-item, breadcrumbs-separator") {
				<nav class="breadcrumbs">
					<ol class="breadcrumbs-list">
						<li class="breadcrumbs-item"><a href="#">Home</a></li>
						<span class="breadcrumbs-separator">/</span>
						<li class="breadcrumbs-item"><a href="#">Products</a></li>
						<span class="breadcrumbs-separator">/</span>
						<li class="breadcrumbs-item breadcrumbs-current">Details</li>
					</ol>
				</nav>
			}
		</div>
	</section>
}

// TemplatesSection displays template/layout components
templ TemplatesSection() {
	<section id="templates" class="py-12">
		<h2 class="text-2xl font-bold mb-8 pb-4 border-b">Templates</h2>
		<p class="text-muted-foreground mb-8">
			Page layouts: dashboard, form, auth, error, etc.
		</p>
		
		<div class="grid gap-12">
			@DemoCard("Layout Types", "layout-dashboard, layout-form, layout-auth, layout-error") {
				<div class="grid grid-cols-2 md:grid-cols-4 gap-4">
					<div class="card p-4 text-center">
						<div class="text-2xl mb-2">📊</div>
						<div class="font-medium">Dashboard</div>
						<div class="text-xs text-muted-foreground">Sidebar + Main</div>
					</div>
					<div class="card p-4 text-center">
						<div class="text-2xl mb-2">📝</div>
						<div class="font-medium">Form</div>
						<div class="text-xs text-muted-foreground">Centered Form</div>
					</div>
					<div class="card p-4 text-center">
						<div class="text-2xl mb-2">🔐</div>
						<div class="font-medium">Auth</div>
						<div class="text-xs text-muted-foreground">Login/Register</div>
					</div>
					<div class="card p-4 text-center">
						<div class="text-2xl mb-2">⚠️</div>
						<div class="font-medium">Error</div>
						<div class="text-xs text-muted-foreground">404/500 Pages</div>
					</div>
				</div>
			}
		</div>
	</section>
}

// DemoCard wraps a demo section with title and classes info
templ DemoCard(title, classes string) {
	<div class="card">
		<div class="card-header">
			<h3 class="card-title">{ title }</h3>
			<p class="card-description font-mono text-xs">{ classes }</p>
		</div>
		<div class="card-content">
			{ children... }
		</div>
	</div>
}
TEMPLEOF

  log_success "Created: $DEMO_PAGE"
}

task_T0_3() {
  log_info "Creating demo section templates..."

  # Create go.mod for demo if needed
  if [[ ! -f "$DEMO_DIR/go.mod" ]]; then
    cat >"$DEMO_DIR/go.mod" <<'MODEOF'
module demo

go 1.22

require (
	github.com/a-h/templ v0.2.778
	github.com/gofiber/fiber/v2 v2.52.5
)
MODEOF
    log_info "Created: $DEMO_DIR/go.mod"

    # Run go mod tidy
    cd "$DEMO_DIR" && go mod tidy 2>/dev/null || true
  fi

  log_success "Demo setup complete!"
  log_info "Start demo: $0 --demo"
}

# =============================================================================
# PHASE 1: Analysis Tasks
# =============================================================================

task_T1_1() {
  log_info "Parsing basecoat.css for component classes..."

  claude --dangerously-skip-permissions "Analyze Basecoat CSS and extract all component classes.

Read: $BASECOAT_CSS

Create: $VIEWS_DIR/basecoat-analysis.md

BASECOAT PATTERN (from basecoatui.com):
- Base class: .card, .btn, .field, .input, .checkbox
- Part classes: .card-header, .card-title, .field-label, .field-input
- Variant classes: .btn-outline, .btn-destructive, .badge-secondary

Extract and document:
1. All base component classes
2. All part classes (component-part naming)
3. All variant classes
4. CSS custom properties used

Format as markdown tables for reference."
}

task_T1_2() {
  log_info "Scanning Basecoat JS files..."

  claude --dangerously-skip-permissions "Analyze Basecoat JavaScript for behavior patterns.

Directory: $BASECOAT_JS_DIR

Document:
1. Initialization patterns
2. Event handling
3. Alpine.js integration
4. Required HTML structure

Append to: $VIEWS_DIR/basecoat-analysis.md"
}

task_T1_3() {
  log_info "Inventorying existing components..."

  claude --dangerously-skip-permissions "Create inventory of all existing templ components.

Scan:
- $ATOMS_DIR/*.templ
- $MOLECULES_DIR/*.templ  
- $ORGANISMS_DIR/*.templ
- $TEMPLATES_DIR/*.templ

For each file document:
1. File path
2. Exported components
3. Props/parameters
4. Current class usage
5. templ features used (Classes, KV, Attributes, Component)

Create: $VIEWS_DIR/component-inventory.md"
}

task_T1_4() {
  log_info "Mapping Basecoat vs Custom components..."

  claude --dangerously-skip-permissions "Create mapping of Basecoat-provided vs custom components.

BASECOAT PROVIDES:
card, btn, field, input, textarea, select, checkbox, radio, switch,
badge, alert, avatar, dialog, dropdown-menu, popover, tabs, toast,
sidebar, command, table, accordion, separator, skeleton, progress,
tooltip, sheet, scroll-area, collapsible, navigation-menu, spinner

NEED CUSTOM CSS:
icon, kbd, code, link, divider, heading, text,
status-badge, pagination, empty-state, filter-chip, input-group,
data-table, stats-card, wizard, breadcrumbs, layouts

Create: $VIEWS_DIR/component-mapping.md"
}

task_T1_5() {
  log_info "Detecting anti-patterns..."

  local issues=$(detect_anti_patterns "$VIEWS_DIR")

  claude --dangerously-skip-permissions "Perform anti-pattern detection.

Already found:
$issues

Search for:
1. ClassName string props (ANTI-PATTERN - leaky abstraction)

Note: These are ALLOWED and encouraged:
- utils.TwMerge for resolving layout utility conflicts
- templ.Classes() for conditional class composition
- templ.KV() for conditional class inclusion
- templ.Attributes for spreading hx-*, aria-*, data-*
- templ.Component for slots, icons, composition

Create: $VIEWS_DIR/anti-patterns-report.md"
}

task_T1_6() {
  log_info "Generating analysis report..."

  claude --dangerously-skip-permissions "Generate Phase 1 Analysis Report.

Compile into: $VIEWS_DIR/analysis-report.md

Include:
1. Basecoat CSS Coverage
2. Component Inventory Summary
3. Basecoat vs Custom Mapping
4. Anti-Pattern Summary
5. Effort Estimate

Update $REFACTOR_GUIDE"
}

# =============================================================================
# PHASE 2: Strategy Tasks
# =============================================================================

task_T2_1() {
  log_info "Defining component API standards..."

  claude --dangerously-skip-permissions "Define component API standards.

Create: $VIEWS_DIR/component-api-standards.md

TEMPL PATTERNS TO USE:

1. templ.Classes() + templ.KV() for conditional classes:
\`\`\`go
class={ templ.Classes(
    \"btn\",
    templ.KV(\"btn-outline\", props.Variant == ButtonOutline),
    templ.KV(\"btn-destructive\", props.Variant == ButtonDestructive),
) }
\`\`\`

2. templ.Attributes for HTMX/ARIA spreading:
\`\`\`go
type ButtonProps struct {
    Attrs templ.Attributes
}
// Usage: { props.Attrs... }
\`\`\`

3. templ.Component for slots:
\`\`\`go
type CardProps struct {
    Header templ.Component
    Footer templ.Component
}
\`\`\`

4. utils.TwMerge for layout conflicts:
\`\`\`go
class={ utils.TwMerge(\"flex gap-4\", props.Layout) }
\`\`\`

KEY RULES:
- ✅ templ.Classes(), templ.KV(), templ.Attributes, templ.Component
- ✅ utils.TwMerge for Tailwind layout conflicts
- ✅ Basecoat semantic classes
- ❌ ClassName string props"
}

task_T2_2() {
  log_info "Creating shared types..."

  claude --dangerously-skip-permissions "Create shared type definitions.

Create: $COMPONENTS_DIR/types.go

Include:
- Size (SizeDefault, SizeSm, SizeLg)
- ButtonVariant (ButtonDefault, ButtonOutline, ButtonSecondary, ButtonDestructive, ButtonGhost, ButtonLink)
- BadgeVariant, AlertVariant
- Status (StatusActive, StatusInactive, StatusPending, StatusApproved, StatusRejected, StatusDraft)
- InputType
- Position
- TrendDirection
- WithIcon struct { Icon templ.Component; IconPos string }
- WithSlots struct { Header, Footer templ.Component }"
}

task_T2_3() {
  log_info "Updating helper functions..."

  claude --dangerously-skip-permissions "Review and update helper functions.

File: $UTILS_PKG/util.go

The existing utils are GOOD:
- TwMerge - resolve Tailwind conflicts ✅
- If/IfElse - conditional values ✅
- MergeAttributes - combine templ.Attributes ✅
- RandomID - generate unique IDs ✅

Create: $VIEWS_DIR/utils-guide.md documenting how to use these with templ."
}

task_T2_4() {
  log_info "Defining refactoring priority..."

  claude --dangerously-skip-permissions "Define refactoring priority order.

Create: $VIEWS_DIR/refactor-order.md

Week-by-week plan from atoms to pages."
}

task_T2_5() {
  log_info "Generating strategy report..."

  claude --dangerously-skip-permissions "Generate Phase 2 Strategy Report.

Create: $VIEWS_DIR/strategy-report.md

Update $REFACTOR_GUIDE"
}

# =============================================================================
# PHASE 3: Custom CSS Tasks
# =============================================================================

task_T3_1() {
  log_info "Creating CSS directory structure..."

  mkdir -p "$CUSTOM_CSS_DIR"

  claude --dangerously-skip-permissions "Create CSS variables file.

Create: $CUSTOM_CSS_DIR/_variables.css

Include variables for:
- Icon sizes
- Status colors (active, inactive, pending, approved, rejected, draft)
- Data table colors
- Stats card trend colors
- Wizard step sizes
- Layout dimensions"
}

generate_css() {
  local component="$1"
  local content="$2"

  claude --dangerously-skip-permissions "Create CSS for $component component.

File: $CUSTOM_CSS_DIR/$component.css

Follow Basecoat pattern (semantic classes):
- Base class: .$component
- Part classes: .$component-{part}
- Variant classes: .$component-{variant}

Component spec:
$content

Create production-ready CSS."
}

task_T3_2() { generate_css "icon" "Icon with sizes xs/sm/md/lg/xl, spin animation. Classes: icon, icon-xs, icon-sm, icon-lg, icon-xl, icon-spin"; }
task_T3_3() { generate_css "kbd" "Keyboard shortcut display. Classes: kbd, kbd-group, kbd-separator"; }
task_T3_4() { generate_css "code" "Code display inline and block. Classes: code, code-block, code-block-header"; }
task_T3_5() { generate_css "link" "Link styling. Classes: link, link-muted, link-external"; }
task_T3_6() { generate_css "divider" "Divider with optional label. Classes: divider, divider-vertical, divider-label"; }
task_T3_7() { generate_css "status-badge" "Status indicator. Classes: status-badge, status-badge-{active|inactive|pending|approved|rejected|draft}, status-badge-dot"; }
task_T3_8() { generate_css "pagination" "Pagination controls. Classes: pagination, pagination-info, pagination-controls, pagination-button, pagination-button-active, pagination-ellipsis"; }
task_T3_9() { generate_css "empty-state" "Empty state display. Classes: empty-state, empty-state-compact, empty-state-icon, empty-state-title, empty-state-description, empty-state-actions"; }
task_T3_10() { generate_css "filter-chip" "Filter chip. Classes: filter-chip, filter-chip-active, filter-chip-remove"; }
task_T3_11() { generate_css "input-group" "Input with addons. Classes: input-group, input-group-addon"; }
task_T3_12() { generate_css "data-table" "Data table. Classes: data-table, data-table-wrapper, data-table-header, data-table-head, data-table-row, data-table-cell, data-table-sortable, data-table-loading, data-table-empty"; }
task_T3_13() { generate_css "stats-card" "Stats card. Classes: stats-card, stats-card-icon, stats-card-value, stats-card-label, stats-card-trend, stats-card-trend-up, stats-card-trend-down"; }
task_T3_14() { generate_css "wizard" "Wizard/stepper. Classes: wizard, wizard-steps, wizard-step, wizard-step-completed, wizard-step-active, wizard-connector, wizard-content, wizard-nav"; }
task_T3_15() { generate_css "breadcrumbs" "Breadcrumbs. Classes: breadcrumbs, breadcrumbs-list, breadcrumbs-item, breadcrumbs-separator, breadcrumbs-current"; }

task_T3_16() {
  log_info "Generating layout CSS..."

  claude --dangerously-skip-permissions "Create CSS for all layout components in $CUSTOM_CSS_DIR/:

1. layout-dashboard.css - Sidebar + main grid
2. layout-master-detail.css - Split view
3. layout-form.css - Centered form
4. layout-auth.css - Auth card centered
5. layout-error.css - Error display
6. layout-wizard.css - Wizard layout
7. layout-kanban.css - Kanban board

Include responsive breakpoints."
}

task_T3_17() {
  log_info "Creating components index.css..."

  claude --dangerously-skip-permissions "Create index file importing all custom CSS.

Create: $CUSTOM_CSS_DIR/index.css

Import all component CSS files."
}

# =============================================================================
# PHASE 4: Atoms
# =============================================================================

refactor_atom() {
  local component="$1"
  local file="$ATOMS_DIR/${component}.templ"
  local create="${2:-false}"

  [[ ! -f "$file" && "$create" != "true" ]] && {
    log_warn "File not found: $file"
    return 0
  }

  claude --dangerously-skip-permissions "Refactor/Create atom: $file

Use these templ patterns:
- templ.Classes() + templ.KV() for conditional classes
- templ.Attributes (Attrs field) for hx-*, aria-*, data-*
- templ.Component for icon injection
- Basecoat classes: btn, btn-outline, field-input, badge
- Tailwind for layout: flex, mr-2, gap-4
- NO ClassName string prop

Create complete component for: $component"
}

task_T4_1() { refactor_atom "button"; }
task_T4_2() { refactor_atom "input"; }
task_T4_3() { refactor_atom "label"; }
task_T4_4() { refactor_atom "icon"; }
task_T4_5() { refactor_atom "badge"; }
task_T4_6() { refactor_atom "spinner"; }
task_T4_7() { refactor_atom "separator"; }
task_T4_8() { refactor_atom "avatar"; }
task_T4_9() { refactor_atom "checkbox"; }
task_T4_10() { refactor_atom "radio"; }
task_T4_11() { refactor_atom "select"; }
task_T4_12() { refactor_atom "textarea"; }
task_T4_13() { refactor_atom "switch"; }
task_T4_14() { refactor_atom "link" "true"; }
task_T4_15() { refactor_atom "image" "true"; }
task_T4_16() { refactor_atom "tooltip"; }
task_T4_17() { refactor_atom "progress"; }
task_T4_18() { refactor_atom "skeleton"; }
task_T4_19() { refactor_atom "alert"; }
task_T4_20() { refactor_atom "heading" "true"; }
task_T4_21() { refactor_atom "text" "true"; }
task_T4_22() { refactor_atom "code" "true"; }
task_T4_23() { refactor_atom "kbd"; }
task_T4_24() { refactor_atom "tags"; }

task_T4_25() {
  log_info "Updating demo: Atoms section..."

  claude --dangerously-skip-permissions "Update the AtomsSection in $DEMO_PAGE to showcase all newly created atom components.

Import the atom components from $ATOMS_DIR and use them in the demo.
Show all variants and states for each component.
Update the DemoCard blocks to use the actual templ components instead of raw HTML."
}

# =============================================================================
# PHASE 5: Molecules
# =============================================================================

refactor_molecule() {
  local component="$1"
  local file="$MOLECULES_DIR/${component}.templ"
  local create="${2:-false}"

  [[ ! -f "$file" && "$create" != "true" ]] && create="true"

  claude --dangerously-skip-permissions "Refactor/Create molecule: $file

Use templ patterns:
- templ.Classes() + templ.KV()
- templ.Attributes
- templ.Component for slots
- Compose atoms from $ATOMS_DIR

Create molecule: $component"
}

task_T5_1() { refactor_molecule "form_field"; }
task_T5_2() { refactor_molecule "searchbox"; }
task_T5_3() { refactor_molecule "stat_display" "true"; }
task_T5_4() { refactor_molecule "quantity_selector" "true"; }
task_T5_5() { refactor_molecule "media_object" "true"; }
task_T5_6() { refactor_molecule "tag_input" "true"; }
task_T5_7() { refactor_molecule "date_picker"; }
task_T5_8() { refactor_molecule "currency_input" "true"; }
task_T5_9() { refactor_molecule "file_upload" "true"; }
task_T5_10() { refactor_molecule "breadcrumb_item" "true"; }
task_T5_11() { refactor_molecule "status_badge" "true"; }
task_T5_12() { refactor_molecule "pagination" "true"; }
task_T5_13() { refactor_molecule "empty_state" "true"; }
task_T5_14() { refactor_molecule "filter_chip" "true"; }
task_T5_15() { refactor_molecule "action_buttons" "true"; }
task_T5_16() { refactor_molecule "card"; }
task_T5_17() { refactor_molecule "list_item" "true"; }
task_T5_18() { refactor_molecule "tabs"; }
task_T5_19() { refactor_molecule "dropdown_menu"; }
task_T5_20() { refactor_molecule "input_group" "true"; }

task_T5_21() {
  log_info "Updating demo: Molecules section..."

  claude --dangerously-skip-permissions "Update the MoleculesSection in $DEMO_PAGE to showcase all molecule components.

Import molecules from $MOLECULES_DIR and use them in the demo.
Show cards, tabs, status badges, pagination, empty states, etc."
}

# =============================================================================
# PHASE 6: Organisms
# =============================================================================

refactor_organism() {
  local component="$1"
  local file="$ORGANISMS_DIR/${component}.templ"
  local create="${2:-false}"

  [[ ! -f "$file" && "$create" != "true" ]] && create="true"

  claude --dangerously-skip-permissions "Refactor/Create organism: $file

Compose molecules and atoms.
Use templ.Component for flexible slots.
Use HTMX for server interactions.
Use Alpine.js for client state.

Create organism: $component"
}

task_T6_1() { refactor_organism "datatable"; }
task_T6_2() { refactor_organism "navigation"; }
task_T6_3() { refactor_organism "sidebar"; }
task_T6_4() { refactor_organism "invoice_card" "true"; }
task_T6_5() { refactor_organism "product_card" "true"; }
task_T6_6() { refactor_organism "form"; }
task_T6_7() { refactor_organism "filter_panel" "true"; }
task_T6_8() { refactor_organism "chart_widget" "true"; }
task_T6_9() { refactor_organism "dialog"; }
task_T6_10() { refactor_organism "breadcrumbs" "true"; }
task_T6_11() { refactor_organism "user_menu" "true"; }
task_T6_12() { refactor_organism "notification_list" "true"; }
task_T6_13() { refactor_organism "popover"; }
task_T6_14() { refactor_organism "calendar_widget" "true"; }
task_T6_15() { refactor_organism "comment_thread" "true"; }
task_T6_16() { refactor_organism "file_explorer" "true"; }
task_T6_17() { refactor_organism "stats_card" "true"; }
task_T6_18() { refactor_organism "wizard" "true"; }

task_T6_19() {
  log_info "Updating demo: Organisms section..."

  claude --dangerously-skip-permissions "Update the OrganismsSection in $DEMO_PAGE to showcase all organism components.

Import organisms from $ORGANISMS_DIR.
Show data tables, stats cards, dialogs, breadcrumbs, etc."
}

# =============================================================================
# PHASE 7: Templates
# =============================================================================

refactor_template() {
  local component="$1"
  local file="$TEMPLATES_DIR/${component}.templ"
  local create="${2:-false}"

  [[ ! -f "$file" && "$create" != "true" ]] && create="true"

  claude --dangerously-skip-permissions "Refactor/Create template: $file

Use templ.Component slots for Sidebar, Topbar, etc.
Include CSS and JS links.

Create template: $component"
}

task_T7_1() { refactor_template "base_layout"; }
task_T7_2() { refactor_template "master_detail_layout" "true"; }
task_T7_3() { refactor_template "dashboard_layout"; }
task_T7_4() { refactor_template "form_layout" "true"; }
task_T7_5() { refactor_template "report_layout" "true"; }
task_T7_6() { refactor_template "content_layout" "true"; }
task_T7_7() { refactor_template "auth_layout" "true"; }
task_T7_8() { refactor_template "split_view_layout" "true"; }
task_T7_9() { refactor_template "settings_layout" "true"; }
task_T7_10() { refactor_template "wizard_layout" "true"; }
task_T7_11() { refactor_template "error_layout" "true"; }
task_T7_12() { refactor_template "kanban_layout" "true"; }

task_T7_13() {
  log_info "Updating demo: Templates section..."

  claude --dangerously-skip-permissions "Update the TemplatesSection in $DEMO_PAGE to showcase layout templates.

Show previews/thumbnails of each layout type.
Link to full-page demos if available."
}

# =============================================================================
# PHASE 8: Pages
# =============================================================================

create_page() {
  local component="$1"
  local file="$PAGES_DIR/${component}.templ"

  claude --dangerously-skip-permissions "Create page: $file

Compose templates with content.
Use HTMX for data loading.
Use Alpine.js for UI state.

Create complete page: $component"
}

task_T8_1() { create_page "login_page"; }
task_T8_2() { create_page "dashboard_home_page"; }
task_T8_3() { create_page "invoice_list_page"; }
task_T8_4() { create_page "invoice_detail_page"; }
task_T8_5() { create_page "invoice_create_page"; }
task_T8_6() { create_page "customer_list_page"; }
task_T8_7() { create_page "product_catalog_page"; }
task_T8_8() { create_page "product_detail_page"; }
task_T8_9() { create_page "settings_page"; }
task_T8_10() { create_page "user_profile_page"; }
task_T8_11() { create_page "reports_index_page"; }
task_T8_12() { create_page "sales_report_page"; }
task_T8_13() { create_page "search_results_page"; }
task_T8_14() { create_page "not_found_page"; }
task_T8_15() { create_page "station_dashboard_page"; }

# =============================================================================
# PHASE 9: Validation
# =============================================================================

task_T9_1() {
  log_info "Checking templ compilation..."
  run_templ_generate || fix_build_errors
}

task_T9_2() {
  log_info "Verifying Basecoat classes..."
  claude --dangerously-skip-permissions "Verify all Basecoat classes used exist in $BASECOAT_CSS.
Create: $VIEWS_DIR/class-validation-report.md"
}

task_T9_3() {
  log_info "Verifying custom CSS..."
  claude --dangerously-skip-permissions "Verify $CUSTOM_CSS_DIR has all needed files.
Create: $VIEWS_DIR/css-validation-report.md"
}

task_T9_4() {
  log_info "Scanning for anti-patterns..."
  local issues=$(detect_anti_patterns "$VIEWS_DIR")
  if [[ -n "$issues" ]]; then
    log_warn "Anti-patterns found:"
    echo "$issues"
    claude --dangerously-skip-permissions "Remove anti-patterns:
$issues

Remember: TwMerge, templ.Classes, templ.KV, templ.Attributes, templ.Component are ALLOWED.
Only ClassName string props are anti-patterns."
  fi
}

task_T9_5() {
  log_info "Testing HTMX..."
  claude --dangerously-skip-permissions "Verify HTMX patterns in all components.
Create: $VIEWS_DIR/htmx-validation.md"
}

task_T9_6() {
  log_info "Testing Alpine.js..."
  claude --dangerously-skip-permissions "Verify Alpine.js patterns in all components.
Create: $VIEWS_DIR/alpine-validation.md"
}

task_T9_7() {
  log_info "Finalizing component gallery demo..."

  claude --dangerously-skip-permissions "Finalize the demo page at $DEMO_PAGE.

1. Ensure all components are imported and displayed
2. Add interactive examples using Alpine.js
3. Add HTMX examples with mock endpoints
4. Polish styling and navigation
5. Add a component count summary"
}

task_T9_8() {
  log_info "Generating final report..."
  claude --dangerously-skip-permissions "Generate final report.
Create: $VIEWS_DIR/final-report.md

Include:
- Component summary (count by type)
- Templ patterns used
- Migration guide
- Demo server instructions"
}

# =============================================================================
# Prerequisites & Setup
# =============================================================================

check_prerequisites() {
  log_info "Checking prerequisites..."
  local missing=""
  command -v claude &>/dev/null || missing+="- Claude CLI\n"
  command -v go &>/dev/null || missing+="- Go\n"
  [[ -d "$VIEWS_DIR" ]] || missing+="- Views dir: $VIEWS_DIR\n"

  if [[ -n "$missing" ]]; then
    log_error "Missing:\n$missing"
    exit 1
  fi
  log_success "Prerequisites OK"
}

create_backup() {
  if [[ ! -d "$BACKUP_DIR" ]]; then
    log_info "Creating backup: $BACKUP_DIR"
    cp -r "$VIEWS_DIR" "$BACKUP_DIR"
  fi
}

init_refactor_guide() {
  [[ -f "$REFACTOR_GUIDE" ]] && return
  cat >"$REFACTOR_GUIDE" <<'EOF'
# UI Component Refactoring Guide

## Patterns

### Allowed (✅)
- `utils.TwMerge` - resolve Tailwind layout conflicts
- `templ.Classes()` - conditional class composition
- `templ.KV()` - conditional class inclusion
- `templ.Attributes` - spread hx-*, aria-*, data-*
- `templ.Component` - slots, icons, composition
- Basecoat classes: btn, card, field, input, badge
- Tailwind for layout: grid, flex, gap-4

### Not Allowed (❌)
- `ClassName string` prop - leaky abstraction

## Demo Server
```bash
./component-refactor.sh --demo
```
Opens http://localhost:3000 with kitchen-sink component gallery.

## Status
| Phase | Tasks | Status |
|-------|-------|--------|
| 0 Demo Setup | 3 | 🔵 |
| 1 Analysis | 6 | 🔵 |
| 2 Strategy | 5 | 🔵 |
| 3 CSS | 17 | 🔵 |
| 4 Atoms | 25 | 🔵 |
| 5 Molecules | 21 | 🔵 |
| 6 Organisms | 19 | 🔵 |
| 7 Templates | 13 | 🔵 |
| 8 Pages | 15 | 🔵 |
| 9 Validation | 8 | 🔵 |

**Total: 132 tasks**
EOF
}

# =============================================================================
# Main
# =============================================================================

show_help() {
  cat <<'EOF'
UI Component Refactoring Script

Usage:
  ./component-refactor.sh [OPTIONS]

Options:
  --auto        Run all tasks
  --resume      Resume from checkpoint
  --status      Show progress
  --task ID     Run specific task
  --from ID     Run from task onwards
  --phase N     Run specific phase (0-9)
  --demo        Start demo server (kitchen sink)
  --css-only    CSS generation only (Phase 3)
  --reset       Reset state
  --help        Show help

Demo Server:
  After running tasks, start the demo server to view components:
  
    ./component-refactor.sh --demo
  
  Opens http://localhost:3000 with:
  - Kitchen sink component gallery
  - Atoms, Molecules, Organisms sections
  - Interactive examples with Alpine.js

Templ Patterns Used:
  ✅ templ.Classes() - conditional class composition
  ✅ templ.KV()      - conditional class inclusion
  ✅ templ.Attributes - spread hx-*, aria-*, data-*
  ✅ templ.Component  - slots, icons, composition
  ✅ utils.TwMerge   - resolve Tailwind conflicts
  ❌ ClassName prop  - leaky abstraction

Phases:
  0  Demo Setup (3 tasks)
  1  Analysis (6 tasks)
  2  Strategy (5 tasks)
  3  Custom CSS (17 tasks)
  4  Atoms (25 tasks)
  5  Molecules (21 tasks)
  6  Organisms (19 tasks)
  7  Templates (13 tasks)
  8  Pages (15 tasks)
  9  Validation (8 tasks)
EOF
}

main() {
  register_all_tasks
  init_state
  init_refactor_guide

  case "${1:-}" in
  --help | -h)
    show_help
    exit 0
    ;;
  --status)
    show_status
    exit 0
    ;;
  --reset)
    reset_state
    exit 0
    ;;
  --demo)
    start_demo_server
    exit 0
    ;;
  --auto)
    log_header "Full Refactoring (132 Tasks)"
    check_prerequisites
    create_backup
    START_TIME=$(date +%s)
    run_all_tasks
    ;;
  --resume)
    log_header "Resuming"
    check_prerequisites
    resume_tasks
    ;;
  --task)
    [[ -z "${2:-}" ]] && {
      log_error "Task ID required"
      exit 1
    }
    check_prerequisites
    execute_task "$2"
    ;;
  --from)
    [[ -z "${2:-}" ]] && {
      log_error "Task ID required"
      exit 1
    }
    check_prerequisites
    create_backup
    run_from_task "$2"
    ;;
  --phase)
    [[ -z "${2:-}" ]] && {
      log_error "Phase required"
      exit 1
    }
    check_prerequisites
    run_phase "$2"
    ;;
  --css-only)
    check_prerequisites
    run_phase 3
    ;;
  "")
    log_header "Interactive Mode"
    check_prerequisites
    echo "1) Run all  2) Resume  3) Status  4) Phase  5) Task  6) Demo  q) Quit"
    read -p "Choice: " c
    case "$c" in
    1)
      create_backup
      run_all_tasks
      ;;
    2) resume_tasks ;;
    3) show_status ;;
    4)
      read -p "Phase (0-9): " p
      run_phase "$p"
      ;;
    5)
      show_status
      read -p "Task: " t
      execute_task "$t"
      ;;
    6) start_demo_server ;;
    q) exit 0 ;;
    esac
    ;;
  *)
    log_error "Unknown: $1"
    show_help
    exit 1
    ;;
  esac
}

main "$@"
