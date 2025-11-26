#!/bin/bash
# =============================================================================
# component-refactor.sh - Complete UI Component Refactoring System
# =============================================================================
#
# A comprehensive refactoring script for Awo ERP UI components following:
# - Basecoat UI patterns (shadcn without React)
# - Atomic Design principles (atoms → molecules → organisms → templates → pages)
# - HTMX + Alpine.js integration
# - Three-tier token architecture from pkg/schema
#
# Features:
#   - State tracking for resume capability
#   - Automatic error detection and fixing
#   - Custom CSS generation for non-Basecoat components
#   - Component inventory and analysis
#   - Visual gallery generation
#
# Component Counts:
#   - 24 Atoms
#   - 20 Molecules
#   - 18 Organisms
#   - 12 Templates
#   - 15 Pages
#   - Custom CSS generation
#
# Constraints Enforced:
#   - NO utils.TwMerge anywhere
#   - NO ClassName props of any kind
#   - NO manual Tailwind utilities that duplicate Basecoat
#   - Data attributes preferred over utility classes
#   - Trust basecoat.css - read it first
#
# Usage:
#   ./component-refactor.sh              # Interactive mode
#   ./component-refactor.sh --auto       # Run all tasks automatically
#   ./component-refactor.sh --resume     # Resume from last checkpoint
#   ./component-refactor.sh --status     # Show current progress
#   ./component-refactor.sh --task T1.1  # Run specific task
#   ./component-refactor.sh --from T2.1  # Run from specific task onwards
#   ./component-refactor.sh --reset      # Reset state and start fresh
#   ./component-refactor.sh --phase 3    # Run specific phase
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
CSS_DIR="${CSS_DIR:-$STATIC_DIR/css}"
CUSTOM_CSS_DIR="${CUSTOM_CSS_DIR:-$CSS_DIR/components}"
JS_DIR="${JS_DIR:-$STATIC_DIR/js}"

# Basecoat resources
BASECOAT_CSS="${BASECOAT_CSS:-$CSS_DIR/basecoat.css}"
BASECOAT_JS_DIR="${BASECOAT_JS_DIR:-$JS_DIR/basecoat}"
BASECOAT_JINJA_DIR="${BASECOAT_JINJA_DIR:-./basecoat/src/jinja}"

# State and logging
STATE_FILE="${STATE_FILE:-.component-refactor-state}"
LOG_FILE="${LOG_FILE:-./component-refactor.log}"
REFACTOR_GUIDE="${REFACTOR_GUIDE:-$VIEWS_DIR/refactoring_guide.md}"
BACKUP_DIR="${BACKUP_DIR:-./views_backup_$(date +%Y%m%d_%H%M%S)}"

# Limits
MAX_FIX_ATTEMPTS=5
BUILD_TIMEOUT=120
TEST_TIMEOUT=300

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
# Component Registry - What Basecoat Provides vs Custom
# =============================================================================

# Components that Basecoat provides (from jinja templates and CSS)
BASECOAT_COMPONENTS=(
  "button"
  "input"
  "label"
  "checkbox"
  "radio"
  "switch"
  "select"
  "textarea"
  "badge"
  "alert"
  "avatar"
  "card"
  "dialog"
  "dropdown-menu"
  "popover"
  "command"
  "sidebar"
  "tabs"
  "toast"
  "skeleton"
  "progress"
  "spinner"
  "separator"
  "scroll-area"
  "sheet"
  "tooltip"
  "accordion"
  "collapsible"
  "context-menu"
  "menubar"
  "navigation-menu"
  "hover-card"
  "aspect-ratio"
  "form"
  "table"
)

# Components that need custom CSS (not in Basecoat)
CUSTOM_CSS_COMPONENTS=(
  "icon"
  "kbd"
  "code"
  "pre"
  "heading"
  "text"
  "link"
  "image"
  "divider"
  "status-badge"
  "quantity-selector"
  "currency-input"
  "date-range-picker"
  "tag-input"
  "file-upload"
  "breadcrumbs"
  "filter-chip"
  "action-buttons"
  "card-header"
  "list-item"
  "media-object"
  "search-bar"
  "stat-display"
  "pagination"
  "empty-state"
  "input-group"
  "data-table"
  "navigation-header"
  "sidebar-navigation"
  "invoice-card"
  "product-card"
  "multi-section-form"
  "filter-panel"
  "chart-widget"
  "modal"
  "user-menu"
  "notification-list"
  "calendar-widget"
  "comment-thread"
  "file-explorer"
  "stats-card"
  "wizard"
  "master-detail-layout"
  "dashboard-layout"
  "form-layout"
  "report-layout"
  "content-layout"
  "auth-layout"
  "split-view-layout"
  "settings-layout"
  "wizard-layout"
  "error-layout"
  "list-detail-layout"
  "kanban-layout"
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
log_phase() { log "${WHITE}[PHASE]${NC} $1"; }

log_header() {
  echo "" | tee -a "$LOG_FILE"
  echo -e "${BOLD}════════════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}  $1${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}════════════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo "" | tee -a "$LOG_FILE"
}

log_subheader() {
  echo "" | tee -a "$LOG_FILE"
  echo -e "${DIM}────────────────────────────────────────────────────────────${NC}" | tee -a "$LOG_FILE"
  echo -e "${CYAN}  $1${NC}" | tee -a "$LOG_FILE"
  echo -e "${DIM}────────────────────────────────────────────────────────────${NC}" | tee -a "$LOG_FILE"
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
    echo "# Do not edit manually"
    echo ""
    echo "LAST_TASK=\"$LAST_TASK\""
    echo "LAST_STATUS=\"$LAST_STATUS\""
    echo "START_TIME=\"$START_TIME\""
    echo "CURRENT_PHASE=\"$CURRENT_PHASE\""
    echo ""
    echo "# Task statuses"
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

  local completed=0
  local failed=0
  local pending=0
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

  if [[ -n "$LAST_TASK" ]]; then
    echo "Last task: $LAST_TASK ($LAST_STATUS)"
  fi
}

reset_state() {
  rm -f "$STATE_FILE"
  log_info "State reset. Starting fresh."
}

# =============================================================================
# Task Registry
# =============================================================================

register_task() {
  local id="$1"
  local phase="$2"
  local desc="$3"
  TASK_ORDER+=("$id")
  TASK_DESC[$id]="$desc"
  TASK_PHASE[$id]="$phase"
  TASK_STATUS[$id]="${TASK_STATUS[$id]:-pending}"
}

register_all_tasks() {
  # =========================================================================
  # PHASE 1: Analysis (T1.x)
  # =========================================================================
  register_task "T1.1" "PHASE 1: Analysis" "Load and parse basecoat.css"
  register_task "T1.2" "PHASE 1: Analysis" "Scan Basecoat jinja templates"
  register_task "T1.3" "PHASE 1: Analysis" "Scan Basecoat JS files"
  register_task "T1.4" "PHASE 1: Analysis" "Inventory all existing components"
  register_task "T1.5" "PHASE 1: Analysis" "Map Basecoat vs Custom components"
  register_task "T1.6" "PHASE 1: Analysis" "Detect anti-patterns"
  register_task "T1.7" "PHASE 1: Analysis" "Generate analysis report"

  # =========================================================================
  # PHASE 2: Strategy & Standards (T2.x)
  # =========================================================================
  register_task "T2.1" "PHASE 2: Strategy" "Define component API standards"
  register_task "T2.2" "PHASE 2: Strategy" "Create shared types (types.go)"
  register_task "T2.3" "PHASE 2: Strategy" "Create helper functions (helpers.go)"
  register_task "T2.4" "PHASE 2: Strategy" "Define refactoring priority order"
  register_task "T2.5" "PHASE 2: Strategy" "Generate strategy report"

  # =========================================================================
  # PHASE 3: Custom CSS Generation (T3.x)
  # =========================================================================
  register_task "T3.1" "PHASE 3: Custom CSS" "Create CSS directory structure"
  register_task "T3.2" "PHASE 3: Custom CSS" "Generate icon.css"
  register_task "T3.3" "PHASE 3: Custom CSS" "Generate kbd.css"
  register_task "T3.4" "PHASE 3: Custom CSS" "Generate code.css"
  register_task "T3.5" "PHASE 3: Custom CSS" "Generate link.css"
  register_task "T3.6" "PHASE 3: Custom CSS" "Generate divider.css"
  register_task "T3.7" "PHASE 3: Custom CSS" "Generate status-badge.css"
  register_task "T3.8" "PHASE 3: Custom CSS" "Generate quantity-selector.css"
  register_task "T3.9" "PHASE 3: Custom CSS" "Generate currency-input.css"
  register_task "T3.10" "PHASE 3: Custom CSS" "Generate date-range-picker.css"
  register_task "T3.11" "PHASE 3: Custom CSS" "Generate tag-input.css"
  register_task "T3.12" "PHASE 3: Custom CSS" "Generate file-upload.css"
  register_task "T3.13" "PHASE 3: Custom CSS" "Generate breadcrumbs.css"
  register_task "T3.14" "PHASE 3: Custom CSS" "Generate filter-chip.css"
  register_task "T3.15" "PHASE 3: Custom CSS" "Generate pagination.css"
  register_task "T3.16" "PHASE 3: Custom CSS" "Generate empty-state.css"
  register_task "T3.17" "PHASE 3: Custom CSS" "Generate input-group.css"
  register_task "T3.18" "PHASE 3: Custom CSS" "Generate data-table.css"
  register_task "T3.19" "PHASE 3: Custom CSS" "Generate stats-card.css"
  register_task "T3.20" "PHASE 3: Custom CSS" "Generate wizard.css"
  register_task "T3.21" "PHASE 3: Custom CSS" "Generate layout components CSS"
  register_task "T3.22" "PHASE 3: Custom CSS" "Create components index.css"

  # =========================================================================
  # PHASE 4: Atoms (T4.x) - 24 components
  # =========================================================================
  register_task "T4.1" "PHASE 4: Atoms" "Refactor button.templ"
  register_task "T4.2" "PHASE 4: Atoms" "Refactor input.templ"
  register_task "T4.3" "PHASE 4: Atoms" "Refactor label.templ"
  register_task "T4.4" "PHASE 4: Atoms" "Refactor icon.templ"
  register_task "T4.5" "PHASE 4: Atoms" "Refactor badge.templ"
  register_task "T4.6" "PHASE 4: Atoms" "Refactor spinner.templ"
  register_task "T4.7" "PHASE 4: Atoms" "Refactor divider.templ (separator)"
  register_task "T4.8" "PHASE 4: Atoms" "Refactor avatar.templ"
  register_task "T4.9" "PHASE 4: Atoms" "Refactor checkbox.templ"
  register_task "T4.10" "PHASE 4: Atoms" "Refactor radio.templ"
  register_task "T4.11" "PHASE 4: Atoms" "Refactor select.templ"
  register_task "T4.12" "PHASE 4: Atoms" "Refactor textarea.templ"
  register_task "T4.13" "PHASE 4: Atoms" "Refactor switch.templ"
  register_task "T4.14" "PHASE 4: Atoms" "Refactor link.templ"
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

  # =========================================================================
  # PHASE 5: Molecules (T5.x) - 20 components
  # =========================================================================
  register_task "T5.1" "PHASE 5: Molecules" "Refactor form_field.templ"
  register_task "T5.2" "PHASE 5: Molecules" "Refactor searchbox.templ (SearchBar)"
  register_task "T5.3" "PHASE 5: Molecules" "Create stat_display.templ"
  register_task "T5.4" "PHASE 5: Molecules" "Create quantity_selector.templ"
  register_task "T5.5" "PHASE 5: Molecules" "Create media_object.templ"
  register_task "T5.6" "PHASE 5: Molecules" "Create tag_input.templ"
  register_task "T5.7" "PHASE 5: Molecules" "Refactor date_picker.templ (DateRangePicker)"
  register_task "T5.8" "PHASE 5: Molecules" "Create currency_input.templ"
  register_task "T5.9" "PHASE 5: Molecules" "Create file_upload.templ"
  register_task "T5.10" "PHASE 5: Molecules" "Create breadcrumb_item.templ"
  register_task "T5.11" "PHASE 5: Molecules" "Create status_badge.templ"
  register_task "T5.12" "PHASE 5: Molecules" "Create pagination.templ"
  register_task "T5.13" "PHASE 5: Molecules" "Create empty_state.templ"
  register_task "T5.14" "PHASE 5: Molecules" "Create filter_chip.templ"
  register_task "T5.15" "PHASE 5: Molecules" "Create action_buttons.templ"
  register_task "T5.16" "PHASE 5: Molecules" "Create card_header.templ"
  register_task "T5.17" "PHASE 5: Molecules" "Create list_item.templ"
  register_task "T5.18" "PHASE 5: Molecules" "Refactor tabs.templ (molecule part)"
  register_task "T5.19" "PHASE 5: Molecules" "Refactor dropdown_menu.templ"
  register_task "T5.20" "PHASE 5: Molecules" "Create input_group.templ"

  # =========================================================================
  # PHASE 6: Organisms (T6.x) - 18 components
  # =========================================================================
  register_task "T6.1" "PHASE 6: Organisms" "Refactor datatable.templ"
  register_task "T6.2" "PHASE 6: Organisms" "Refactor navigation.templ (NavigationHeader)"
  register_task "T6.3" "PHASE 6: Organisms" "Create sidebar_navigation.templ"
  register_task "T6.4" "PHASE 6: Organisms" "Create invoice_card.templ"
  register_task "T6.5" "PHASE 6: Organisms" "Create product_card.templ"
  register_task "T6.6" "PHASE 6: Organisms" "Refactor form.templ (MultiSectionForm)"
  register_task "T6.7" "PHASE 6: Organisms" "Create filter_panel.templ"
  register_task "T6.8" "PHASE 6: Organisms" "Create chart_widget.templ"
  register_task "T6.9" "PHASE 6: Organisms" "Create modal.templ"
  register_task "T6.10" "PHASE 6: Organisms" "Create breadcrumbs.templ"
  register_task "T6.11" "PHASE 6: Organisms" "Create user_menu.templ"
  register_task "T6.12" "PHASE 6: Organisms" "Create notification_list.templ"
  register_task "T6.13" "PHASE 6: Organisms" "Create dropdown.templ (generic)"
  register_task "T6.14" "PHASE 6: Organisms" "Create calendar_widget.templ"
  register_task "T6.15" "PHASE 6: Organisms" "Create comment_thread.templ"
  register_task "T6.16" "PHASE 6: Organisms" "Create file_explorer.templ"
  register_task "T6.17" "PHASE 6: Organisms" "Create stats_card.templ"
  register_task "T6.18" "PHASE 6: Organisms" "Create wizard.templ"

  # =========================================================================
  # PHASE 7: Templates (T7.x) - 12 layouts
  # =========================================================================
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

  # =========================================================================
  # PHASE 8: Pages (T8.x) - 15 pages
  # =========================================================================
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

  # =========================================================================
  # PHASE 9: Validation (T9.x)
  # =========================================================================
  register_task "T9.1" "PHASE 9: Validation" "Check templ compilation"
  register_task "T9.2" "PHASE 9: Validation" "Verify all Basecoat classes exist"
  register_task "T9.3" "PHASE 9: Validation" "Verify custom CSS is complete"
  register_task "T9.4" "PHASE 9: Validation" "Scan for remaining anti-patterns"
  register_task "T9.5" "PHASE 9: Validation" "Test HTMX integration"
  register_task "T9.6" "PHASE 9: Validation" "Test Alpine.js integration"
  register_task "T9.7" "PHASE 9: Validation" "Create visual component gallery"
  register_task "T9.8" "PHASE 9: Validation" "Generate final report"
}

# =============================================================================
# Build & Validation Helpers
# =============================================================================

run_templ_generate() {
  local log_file="tmp/log/templ_generate_$$.log"

  if ! command -v templ &>/dev/null; then
    log_warn "templ command not found - skipping generation"
    return 0
  fi

  timeout "$BUILD_TIMEOUT" templ generate 2>&1 | tee "$log_file"
  local result=${PIPESTATUS[0]}

  if [[ $result -eq 0 ]]; then
    rm -f "$log_file"
    return 0
  else
    cat "$log_file"
    return 1
  fi
}

run_go_build() {
  local log_file="tmp/log/go_build_$$.log"
  timeout "$BUILD_TIMEOUT" go build ./... 2>&1 | tee "$log_file"
  local result=${PIPESTATUS[0]}

  if [[ $result -eq 0 ]]; then
    rm -f "$log_file"
    return 0
  fi
  return 1
}

# Check if component is in Basecoat
is_basecoat_component() {
  local component="$1"
  for bc in "${BASECOAT_COMPONENTS[@]}"; do
    if [[ "$bc" == "$component" ]]; then
      return 0
    fi
  done
  return 1
}

# Anti-pattern detection
detect_anti_patterns() {
  local dir="${1:-$COMPONENTS_DIR}"
  local issues=""

  local twmerge=$(grep -rl "TwMerge\|twMerge\|tw-merge\|tw\.Merge" "$dir" --include="*.templ" 2>/dev/null || true)
  if [[ -n "$twmerge" ]]; then
    issues+="TwMerge found in:\n$twmerge\n\n"
  fi

  local classname=$(grep -rl "ClassName\|className\|class_name" "$dir" --include="*.templ" 2>/dev/null || true)
  if [[ -n "$classname" ]]; then
    issues+="ClassName props found in:\n$classname\n\n"
  fi

  echo -e "$issues"
}

# =============================================================================
# Self-Healing Build System
# =============================================================================

fix_build_errors() {
  local attempt=1

  while [[ $attempt -le $MAX_FIX_ATTEMPTS ]]; do
    log_info "Build attempt $attempt/$MAX_FIX_ATTEMPTS..."

    if run_templ_generate && run_go_build; then
      log_success "Build passed!"
      return 0
    fi

    local errors=$(templ generate 2>&1 | grep -E "(error|Error)" | head -20)

    if [[ -z "$errors" ]]; then
      errors=$(go build ./... 2>&1 | grep -E "^\./" | head -20)
    fi

    if [[ -z "$errors" ]]; then
      log_error "Build failed but couldn't extract errors"
      return 1
    fi

    log_fix "Attempting automatic fix (attempt $attempt)..."

    claude --dangerously-skip-permissions "Fix these templ/Go build errors:

$errors

FIXING RULES:
1. SYNTAX ERROR: Fix templ syntax (proper @ expressions, correct HTML)
2. UNDEFINED: Import missing packages or fix references
3. TYPE ERROR: Ensure proper Go types in templ expressions
4. DUPLICATE: Remove duplicate definitions

Fix ALL errors. The code MUST compile after your changes."

    ((attempt++))
    sleep 2
  done

  log_error "Could not fix build errors after $MAX_FIX_ATTEMPTS attempts"
  return 1
}

ensure_builds() {
  if ! run_templ_generate 2>/dev/null; then
    log_warn "Build failed, attempting fixes..."
    if ! fix_build_errors; then
      return 1
    fi
  fi
  return 0
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

run_from_task() {
  local start_task="$1"
  local started=false

  for task in "${TASK_ORDER[@]}"; do
    if [[ "$task" == "$start_task" ]]; then
      started=true
    fi

    if $started; then
      local status=$(get_task_status "$task")
      if [[ "$status" != "completed" ]]; then
        if ! execute_task "$task"; then
          log_error "Stopping at failed task: $task"
          return 1
        fi
      else
        log_info "Skipping completed task: $task"
      fi
    fi
  done

  return 0
}

run_all_tasks() {
  for task in "${TASK_ORDER[@]}"; do
    local status=$(get_task_status "$task")
    if [[ "$status" != "completed" ]]; then
      if ! execute_task "$task"; then
        log_error "Failed at task: $task"
        echo ""
        echo "To resume, run: $0 --resume"
        echo "To retry this task: $0 --task $task"
        return 1
      fi
    else
      log_info "Skipping completed: $task"
    fi
  done

  log_header "ALL TASKS COMPLETED"
  show_status
  return 0
}

run_phase() {
  local phase_num="$1"
  local phase_prefix="T${phase_num}."

  log_header "Running Phase $phase_num"

  for task in "${TASK_ORDER[@]}"; do
    if [[ "$task" == $phase_prefix* ]]; then
      local status=$(get_task_status "$task")
      if [[ "$status" != "completed" ]]; then
        if ! execute_task "$task"; then
          return 1
        fi
      fi
    fi
  done

  log_success "Phase $phase_num complete!"
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

  log_info "Resuming from task: $resume_from"
  run_from_task "$resume_from"
}

# =============================================================================
# PHASE 1: Analysis Tasks
# =============================================================================

task_T1_1() {
  log_info "Loading and parsing basecoat.css..."

  claude --dangerously-skip-permissions "Analyze the Basecoat CSS file and extract all component classes.

Read: $BASECOAT_CSS

Create a structured analysis at: $VIEWS_DIR/basecoat-analysis.md

Include:
1. All semantic component classes (.btn, .card, .input, etc.)
2. All variant classes (.btn-primary, .btn-outline, etc.)
3. All size modifiers (.btn-sm, .btn-lg, etc.)
4. All state classes (.is-loading, .is-disabled, etc.)
5. Data-attribute patterns used for variants
6. CSS custom properties (variables) available

Format as markdown with tables for easy reference.

Update $REFACTOR_GUIDE with the summary."
}

task_T1_2() {
  log_info "Scanning Basecoat jinja templates..."

  claude --dangerously-skip-permissions "Scan Basecoat jinja templates for component patterns.

Directory: $BASECOAT_JINJA_DIR

Expected files:
- command.html.jinja
- dialog.html.jinja
- dropdown-menu.html.jinja
- popover.html.jinja
- select.html.jinja
- sidebar.html.jinja
- tabs.html.jinja
- toast.html.jinja

For each template, document:
1. HTML structure
2. Class naming conventions
3. Data attributes used
4. Accessibility attributes (aria-*, role)
5. JavaScript integration points

Append to: $VIEWS_DIR/basecoat-analysis.md"
}

task_T1_3() {
  log_info "Scanning Basecoat JS files..."

  claude --dangerously-skip-permissions "Analyze Basecoat JavaScript files for component behavior.

Directory: $BASECOAT_JS_DIR

Files to analyze:
- basecoat.js (main/initialization)
- command.js
- dropdown-menu.js
- popover.js
- select.js
- sidebar.js
- tabs.js
- toast.js

Document:
1. Initialization patterns
2. Event handling
3. Data attribute triggers
4. Alpine.js integration patterns
5. State management approaches

Append to: $VIEWS_DIR/basecoat-analysis.md"
}

task_T1_4() {
  log_info "Inventorying all existing components..."

  claude --dangerously-skip-permissions "Create comprehensive inventory of ALL existing templ components.

Scan directories:
- $ATOMS_DIR/*.templ
- $MOLECULES_DIR/*.templ
- $ORGANISMS_DIR/*.templ
- $TEMPLATES_DIR/*.templ
- $PAGES_DIR/*.templ

For EACH file found, document:
1. File path
2. Package name
3. Exported functions/components
4. Props/parameters
5. Current class usage
6. Anti-patterns present (TwMerge, ClassName, etc.)
7. HTMX attributes used
8. Alpine.js patterns used

Create: $VIEWS_DIR/component-inventory.md

Format as structured markdown with tables."
}

task_T1_5() {
  log_info "Mapping Basecoat vs Custom components..."

  claude --dangerously-skip-permissions "Create a mapping of which components Basecoat provides vs need custom CSS.

Based on the analysis, create: $VIEWS_DIR/component-mapping.md

BASECOAT PROVIDED (use existing classes):
- button, input, label, checkbox, radio, switch, select, textarea
- badge, alert, avatar, card, dialog, dropdown-menu
- popover, command, sidebar, tabs, toast
- skeleton, progress, spinner, separator, scroll-area
- sheet, tooltip, accordion, collapsible, table, form

CUSTOM CSS NEEDED (create in $CUSTOM_CSS_DIR):
- icon, kbd, code, pre, heading, text, link, image, divider
- status-badge, quantity-selector, currency-input, date-range-picker
- tag-input, file-upload, breadcrumbs, filter-chip
- pagination, empty-state, input-group, data-table
- stats-card, wizard, calendar-widget, etc.

Format as two tables with component name, Basecoat equivalent (if any), and custom CSS file needed."
}

task_T1_6() {
  log_info "Detecting anti-patterns..."

  local issues=$(detect_anti_patterns "$VIEWS_DIR")

  claude --dangerously-skip-permissions "Perform comprehensive anti-pattern detection.

Already found:
$issues

Perform deep scan for:
1. TwMerge usage: grep -rn 'TwMerge\|twMerge\|tw\.Merge' $VIEWS_DIR --include='*.templ'
2. ClassName props: grep -rn 'ClassName\|className' $VIEWS_DIR --include='*.templ'
3. Excessive inline Tailwind (>10 classes per element)
4. Hardcoded colors (not using CSS variables)
5. Missing accessibility attributes
6. Inconsistent naming patterns

Create: $VIEWS_DIR/anti-patterns-report.md

Include file, line number, code snippet, and recommended fix for each issue."
}

task_T1_7() {
  log_info "Generating analysis report..."

  claude --dangerously-skip-permissions "Generate comprehensive Phase 1 Analysis Report.

Compile all analysis into: $VIEWS_DIR/analysis-report.md

Include:
1. Executive Summary
2. Basecoat CSS Coverage (from T1.1)
3. Jinja Template Patterns (from T1.2)
4. JavaScript Integration Patterns (from T1.3)
5. Existing Component Inventory (from T1.4)
6. Basecoat vs Custom Mapping (from T1.5)
7. Anti-Pattern Summary (from T1.6)
8. Estimated Effort by Phase

Update $REFACTOR_GUIDE with analysis summary."
}

# =============================================================================
# PHASE 2: Strategy Tasks
# =============================================================================

task_T2_1() {
  log_info "Defining component API standards..."

  claude --dangerously-skip-permissions "Define standardized component API patterns.

Create: $VIEWS_DIR/component-api-standards.md

Include patterns for:

1. TYPE-SAFE VARIANTS:
   type ButtonVariant string
   const (
       ButtonPrimary ButtonVariant = \"primary\"
       ButtonSecondary ButtonVariant = \"secondary\"
   )

2. PROPS STRUCT for >3 parameters:
   type ButtonProps struct {
       Variant  ButtonVariant
       Size     Size
       Disabled bool
   }

3. DATA ATTRIBUTES for state:
   data-variant={ string(props.Variant) }
   data-loading={ strconv.FormatBool(props.Loading) }

4. SEMANTIC CLASSES only:
   class=\"btn\" not class=\"px-4 py-2...\"

5. HTMX INTEGRATION:
   hx-post, hx-target, hx-swap, hx-trigger patterns

6. ALPINE.JS INTEGRATION:
   x-data, x-show, x-on patterns

Document patterns for atoms, molecules, organisms, templates, and pages."
}

task_T2_2() {
  log_info "Creating shared types..."

  claude --dangerously-skip-permissions "Create shared type definitions for all components.

Create: $COMPONENTS_DIR/types.go

Include:

package components

// Common sizes
type Size string
const (
    SizeXS Size = \"xs\"
    SizeSM Size = \"sm\"
    SizeMD Size = \"md\"
    SizeLG Size = \"lg\"
    SizeXL Size = \"xl\"
)

// Intent/semantic colors
type Intent string
const (
    IntentDefault Intent = \"default\"
    IntentPrimary Intent = \"primary\"
    IntentSuccess Intent = \"success\"
    IntentWarning Intent = \"warning\"
    IntentDanger  Intent = \"danger\"
    IntentInfo    Intent = \"info\"
)

// Position for dropdowns, tooltips
type Position string
const (
    PositionTop    Position = \"top\"
    PositionBottom Position = \"bottom\"
    PositionLeft   Position = \"left\"
    PositionRight  Position = \"right\"
)

// Trend direction for stats
type TrendDirection string
const (
    TrendUp   TrendDirection = \"up\"
    TrendDown TrendDirection = \"down\"
    TrendFlat TrendDirection = \"flat\"
)

// Common status types
type Status string
const (
    StatusActive   Status = \"active\"
    StatusInactive Status = \"inactive\"
    StatusPending  Status = \"pending\"
    StatusApproved Status = \"approved\"
    StatusRejected Status = \"rejected\"
    StatusDraft    Status = \"draft\"
)

// HTMX attributes struct
type HtmxAttrs struct {
    Get       string
    Post      string
    Put       string
    Delete    string
    Target    string
    Swap      string
    Trigger   string
    Indicator string
    Confirm   string
    PushUrl   bool
}

Add any other common types needed across components."
}

task_T2_3() {
  log_info "Creating helper functions..."

  claude --dangerously-skip-permissions "Create helper functions for templ components.

Create: $COMPONENTS_DIR/helpers.go

Include:

package components

import \"strconv\"

// boolStr converts bool to string for data attributes
func BoolStr(b bool) string {
    return strconv.FormatBool(b)
}

// BoolAttr returns string only if true
func BoolAttr(b bool) string {
    if b { return \"true\" }
    return \"\"
}

// IfElse returns a or b based on condition
func IfElse(cond bool, a, b string) string {
    if cond { return a }
    return b
}

// JoinClasses joins non-empty class strings
func JoinClasses(classes ...string) string {
    var result []string
    for _, c := range classes {
        if c != \"\" {
            result = append(result, c)
        }
    }
    return strings.Join(result, \" \")
}

// Ptr returns pointer to value
func Ptr[T any](v T) *T {
    return &v
}

// FormatCurrency formats decimal as currency
func FormatCurrency(amount decimal.Decimal, currency string) string {
    // Implementation
}

// FormatDate formats time for display
func FormatDate(t time.Time, format string) string {
    // Implementation
}

// Truncate truncates string with ellipsis
func Truncate(s string, maxLen int) string {
    if len(s) <= maxLen { return s }
    return s[:maxLen-3] + \"...\"
}

Add any other helper functions needed."
}

task_T2_4() {
  log_info "Defining refactoring priority order..."

  claude --dangerously-skip-permissions "Define the optimal refactoring order.

Create: $VIEWS_DIR/refactor-order.md

Order based on:
1. Dependencies (atoms before molecules)
2. Usage frequency (most used first)
3. Anti-pattern severity (most problematic first)
4. Business criticality

Format:

## Priority 1 (Foundation - Week 1)
| Component | Level | Depends On | Used By |
|-----------|-------|------------|---------|
| button | Atom | - | ~50 components |
| input | Atom | - | ~30 components |
| label | Atom | - | ~25 components |

## Priority 2 (Forms - Week 2)
...

Include dependency graph and timeline estimate."
}

task_T2_5() {
  log_info "Generating strategy report..."

  claude --dangerously-skip-permissions "Generate comprehensive Phase 2 Strategy Report.

Create: $VIEWS_DIR/strategy-report.md

Include:
1. API Standards Summary
2. Shared Types Overview
3. Helper Functions Reference
4. Refactoring Order & Timeline
5. Risk Assessment
6. Testing Strategy

Update $REFACTOR_GUIDE with strategy summary."
}

# =============================================================================
# PHASE 3: Custom CSS Generation Tasks
# =============================================================================

task_T3_1() {
  log_info "Creating CSS directory structure..."

  mkdir -p "$CUSTOM_CSS_DIR"

  claude --dangerously-skip-permissions "Create the CSS directory structure and base files.

Create directories:
- $CUSTOM_CSS_DIR/atoms/
- $CUSTOM_CSS_DIR/molecules/
- $CUSTOM_CSS_DIR/organisms/
- $CUSTOM_CSS_DIR/layouts/

Create base variable file: $CUSTOM_CSS_DIR/_variables.css

/* Custom Component Variables - extends Basecoat */
:root {
    /* Component-specific tokens */
    --icon-size-xs: 0.75rem;
    --icon-size-sm: 1rem;
    --icon-size-md: 1.25rem;
    --icon-size-lg: 1.5rem;
    --icon-size-xl: 2rem;
    
    /* Status colors */
    --status-active: var(--color-success);
    --status-inactive: var(--color-muted);
    --status-pending: var(--color-warning);
    --status-danger: var(--color-danger);
    
    /* Data table */
    --table-header-bg: var(--color-muted-50);
    --table-row-hover: var(--color-muted-100);
    --table-border: var(--color-border);
    
    /* Pagination */
    --pagination-gap: 0.25rem;
    --pagination-button-size: 2rem;
    
    /* Add more as needed */
}

Create: $CUSTOM_CSS_DIR/_base.css with shared utilities."
}

# CSS generation helper function
generate_custom_css() {
  local component="$1"
  local category="$2"
  local css_content="$3"

  local css_file="$CUSTOM_CSS_DIR/$category/$component.css"

  claude --dangerously-skip-permissions "Create custom CSS for $component component.

File: $css_file

Requirements:
1. Use CSS custom properties from Basecoat where possible
2. Support data-* attributes for variants/states
3. Include responsive styles
4. Include dark mode support (@media (prefers-color-scheme: dark))
5. Follow BEM-like naming (.component, .component-part, .component--modifier)

Component specifications:
$css_content

Example structure:
/* $component.css */

.${component} {
    /* Base styles */
}

.${component}[data-variant=\"primary\"] {
    /* Variant styles */
}

.${component}[data-size=\"sm\"] {
    /* Size modifier */
}

.${component}[data-disabled=\"true\"] {
    /* Disabled state */
}

@media (prefers-color-scheme: dark) {
    .${component} {
        /* Dark mode overrides */
    }
}

Create complete, production-ready CSS."
}

task_T3_2() {
  log_info "Generating icon.css..."

  generate_custom_css "icon" "atoms" "
Icon component:
- Sizes: xs (12px), sm (16px), md (20px), lg (24px), xl (32px)
- Should support inline and block display
- Color should inherit from parent by default
- Support for spin animation (loading)
- Support for custom colors via data-color attribute
"
}

task_T3_3() {
  log_info "Generating kbd.css..."

  generate_custom_css "kbd" "atoms" "
Keyboard shortcut component:
- Styled as a key cap
- Light background, subtle border, slight shadow
- Support for size variants
- Can be combined (Ctrl+C shows two kbd elements)
"
}

task_T3_4() {
  log_info "Generating code.css..."

  generate_custom_css "code" "atoms" "
Code display component:
- Inline code (.code)
- Code block (.code-block)
- Syntax highlighting support
- Copy button positioning
- Line numbers support
- Language label
"
}

task_T3_5() {
  log_info "Generating link.css..."

  generate_custom_css "link" "atoms" "
Link component:
- Variants: default, primary, secondary, muted
- Underline on hover (optional)
- External link indicator
- Focus states for accessibility
- Visited state
"
}

task_T3_6() {
  log_info "Generating divider.css..."

  generate_custom_css "divider" "atoms" "
Divider/separator component:
- Horizontal and vertical variants
- Optional label in middle
- Thickness variants
- Color variants
"
}

task_T3_7() {
  log_info "Generating status-badge.css..."

  generate_custom_css "status-badge" "molecules" "
Status badge component:
- Status types: active, inactive, pending, approved, rejected, draft, published, archived
- Optional dot indicator (pulsing for active)
- Optional icon
- Size variants
"
}

task_T3_8() {
  log_info "Generating quantity-selector.css..."

  generate_custom_css "quantity-selector" "molecules" "
Quantity selector component:
- Minus button, input, plus button
- Size variants
- Disabled state
- Min/max reached states
"
}

task_T3_9() {
  log_info "Generating currency-input.css..."

  generate_custom_css "currency-input" "molecules" "
Currency input component:
- Currency symbol prefix/suffix
- Number formatting
- Currency selector dropdown
- Size variants
"
}

task_T3_10() {
  log_info "Generating date-range-picker.css..."

  generate_custom_css "date-range-picker" "molecules" "
Date range picker component:
- Two date inputs
- Separator
- Preset ranges dropdown
- Calendar popup styling
- Mobile-friendly layout
"
}

task_T3_11() {
  log_info "Generating tag-input.css..."

  generate_custom_css "tag-input" "molecules" "
Tag input component:
- Tags displayed as badges
- Input field for new tags
- Remove button on tags
- Autocomplete dropdown
- Max tags indicator
"
}

task_T3_12() {
  log_info "Generating file-upload.css..."

  generate_custom_css "file-upload" "molecules" "
File upload component:
- Drag and drop zone
- File list display
- Progress bar for uploads
- Error states
- Preview thumbnails
"
}

task_T3_13() {
  log_info "Generating breadcrumbs.css..."

  generate_custom_css "breadcrumbs" "molecules" "
Breadcrumbs component:
- Breadcrumb container
- Breadcrumb items with separators
- Current item styling
- Truncation for long paths
- Mobile collapse
"
}

task_T3_14() {
  log_info "Generating filter-chip.css..."

  generate_custom_css "filter-chip" "molecules" "
Filter chip component:
- Removable chip
- Active/inactive states
- Icon support
- Count badge
"
}

task_T3_15() {
  log_info "Generating pagination.css..."

  generate_custom_css "pagination" "molecules" "
Pagination component:
- Page buttons
- Active page highlight
- Prev/Next buttons
- First/Last buttons
- Page size selector
- Page info text
- Ellipsis for many pages
"
}

task_T3_16() {
  log_info "Generating empty-state.css..."

  generate_custom_css "empty-state" "molecules" "
Empty state component:
- Centered container
- Large icon
- Title and description
- Action buttons
- Compact variant
"
}

task_T3_17() {
  log_info "Generating input-group.css..."

  generate_custom_css "input-group" "molecules" "
Input group component:
- Prefix/suffix addons
- Icon addons
- Button addons
- Connected inputs
"
}

task_T3_18() {
  log_info "Generating data-table.css..."

  generate_custom_css "data-table" "organisms" "
Data table component:
- Table wrapper with horizontal scroll
- Header row (sticky)
- Sortable column headers
- Selection checkboxes
- Row hover
- Striped variant
- Compact variant
- Loading overlay
- Empty state
- Responsive card view
"
}

task_T3_19() {
  log_info "Generating stats-card.css..."

  generate_custom_css "stats-card" "organisms" "
Stats card component:
- Icon with colored background
- Large value display
- Title and subtitle
- Trend indicator
- Mini chart area
- Actions dropdown
"
}

task_T3_20() {
  log_info "Generating wizard.css..."

  generate_custom_css "wizard" "organisms" "
Wizard/stepper component:
- Step indicator (horizontal/vertical)
- Step circles with numbers/icons
- Connectors between steps
- Completed/active/pending states
- Content area
- Navigation buttons
- Progress bar
"
}

task_T3_21() {
  log_info "Generating layout CSS..."

  claude --dangerously-skip-permissions "Create CSS for all layout components.

Create files in $CUSTOM_CSS_DIR/layouts/:

1. master-detail-layout.css
   - Split view with resizable panels
   - Mobile stack behavior

2. dashboard-layout.css
   - Sidebar + main content
   - Collapsible sidebar
   - Top navigation

3. form-layout.css
   - Centered form container
   - Section dividers
   - Sticky footer

4. report-layout.css
   - Filter sidebar
   - Chart area
   - Data table area

5. auth-layout.css
   - Centered card
   - Background pattern/image

6. wizard-layout.css
   - Step indicator area
   - Content area
   - Navigation area

7. error-layout.css
   - Centered error display
   - Illustration area

8. kanban-layout.css
   - Horizontal scrolling board
   - Column styling
   - Card drag indicators

Each should support:
- Responsive breakpoints
- Dark mode
- CSS custom properties"
}

task_T3_22() {
  log_info "Creating components index.css..."

  claude --dangerously-skip-permissions "Create the main index file that imports all custom CSS.

Create: $CUSTOM_CSS_DIR/index.css

/* Custom Component Styles - Awo ERP */
/* Extends Basecoat UI with custom components */

/* Variables */
@import '_variables.css';
@import '_base.css';

/* Atoms */
@import 'atoms/icon.css';
@import 'atoms/kbd.css';
@import 'atoms/code.css';
@import 'atoms/link.css';
@import 'atoms/divider.css';

/* Molecules */
@import 'molecules/status-badge.css';
@import 'molecules/quantity-selector.css';
@import 'molecules/currency-input.css';
@import 'molecules/date-range-picker.css';
@import 'molecules/tag-input.css';
@import 'molecules/file-upload.css';
@import 'molecules/breadcrumbs.css';
@import 'molecules/filter-chip.css';
@import 'molecules/pagination.css';
@import 'molecules/empty-state.css';
@import 'molecules/input-group.css';

/* Organisms */
@import 'organisms/data-table.css';
@import 'organisms/stats-card.css';
@import 'organisms/wizard.css';

/* Layouts */
@import 'layouts/master-detail-layout.css';
@import 'layouts/dashboard-layout.css';
@import 'layouts/form-layout.css';
@import 'layouts/report-layout.css';
@import 'layouts/auth-layout.css';
@import 'layouts/wizard-layout.css';
@import 'layouts/error-layout.css';
@import 'layouts/kanban-layout.css';

Also update base.css to import this file after basecoat.css."
}

# =============================================================================
# PHASE 4: Atoms Tasks
# =============================================================================

refactor_atom() {
  local component="$1"
  local file="$ATOMS_DIR/${component}.templ"
  local create_new="${2:-false}"

  if [[ ! -f "$file" && "$create_new" != "true" ]]; then
    log_warn "File not found: $file - skipping"
    return 0
  fi

  local basecoat_note=""
  if is_basecoat_component "$component"; then
    basecoat_note="USES BASECOAT: This component has Basecoat CSS. Use .${component} class."
  else
    basecoat_note="CUSTOM CSS: Use classes from $CUSTOM_CSS_DIR/atoms/${component}.css"
  fi

  claude --dangerously-skip-permissions "Refactor/Create atom: $file

$basecoat_note

READ FIRST:
1. $BASECOAT_CSS - Find relevant classes
2. $VIEWS_DIR/component-api-standards.md - Follow patterns
3. $COMPONENTS_DIR/types.go - Use shared types
4. $COMPONENTS_DIR/helpers.go - Use helper functions

RULES:
1. ❌ NO TwMerge
2. ❌ NO ClassName props
3. ✅ Use Basecoat semantic classes OR custom CSS classes
4. ✅ Use data-* attributes for variants/states
5. ✅ Add typed Go constants for variants
6. ✅ Include HTMX support where appropriate
7. ✅ Include accessibility attributes

Example pattern:
\`\`\`templ
package atoms

type ${component^}Variant string

const (
    ${component^}Default ${component^}Variant = \"default\"
    ${component^}Primary ${component^}Variant = \"primary\"
)

type ${component^}Props struct {
    Variant  ${component^}Variant
    Size     Size
    Disabled bool
}

templ ${component^}(props ${component^}Props) {
    <element
        class=\"${component}\"
        data-variant={ string(props.Variant) }
        data-size={ string(props.Size) }
        disabled?={ props.Disabled }
    >
        { children... }
    </element>
}
\`\`\`

Create complete, production-ready component.
Update $REFACTOR_GUIDE log."
}

task_T4_1() { refactor_atom "button"; }
task_T4_2() { refactor_atom "input"; }
task_T4_3() { refactor_atom "label"; }
task_T4_4() { refactor_atom "icon"; }
task_T4_5() { refactor_atom "badge"; }
task_T4_6() { refactor_atom "spinner"; }
task_T4_7() { refactor_atom "divider"; }
task_T4_8() { refactor_atom "avatar"; }
task_T4_9() { refactor_atom "checkbox"; }
task_T4_10() { refactor_atom "radio"; }
task_T4_11() { refactor_atom "select"; }
task_T4_12() { refactor_atom "textarea"; }
task_T4_13() { refactor_atom "switch"; }
task_T4_14() { refactor_atom "link" "true"; }
task_T4_15() { refactor_atom "image" "true"; }
task_T4_16() { refactor_atom "tooltip" "true"; }
task_T4_17() { refactor_atom "progress"; }
task_T4_18() { refactor_atom "skeleton"; }
task_T4_19() { refactor_atom "alert"; }
task_T4_20() { refactor_atom "heading" "true"; }
task_T4_21() { refactor_atom "text" "true"; }
task_T4_22() { refactor_atom "code" "true"; }
task_T4_23() { refactor_atom "kbd"; }
task_T4_24() { refactor_atom "tags"; }

# =============================================================================
# PHASE 5: Molecules Tasks
# =============================================================================

refactor_molecule() {
  local component="$1"
  local file="$MOLECULES_DIR/${component}.templ"
  local create_new="${2:-false}"

  if [[ ! -f "$file" && "$create_new" != "true" ]]; then
    log_warn "File not found: $file - creating new"
    create_new="true"
  fi

  claude --dangerously-skip-permissions "Refactor/Create molecule: $file

READ FIRST:
1. Refactored atoms in $ATOMS_DIR/
2. $VIEWS_DIR/component-api-standards.md
3. $COMPONENTS_DIR/types.go and helpers.go
4. Custom CSS: $CUSTOM_CSS_DIR/molecules/

MOLECULE RULES:
1. Molecules COMPOSE atoms - import and use them
2. ❌ NO TwMerge, ClassName
3. ✅ Use Basecoat OR custom CSS classes
4. ✅ Use data-* attributes for state
5. ✅ Include HTMX patterns for server interaction
6. ✅ Include Alpine.js for client state

Example:
\`\`\`templ
package molecules

import \"yourapp/views/components/atoms\"

type FormFieldProps struct {
    Name     string
    Label    string
    Error    string
    Required bool
}

templ FormField(props FormFieldProps) {
    <div class=\"field\" data-error={ components.BoolStr(props.Error != \"\") }>
        @atoms.Label(atoms.LabelProps{For: props.Name, Required: props.Required}) {
            { props.Label }
        }
        { children... }
        if props.Error != \"\" {
            <p class=\"field-error\" role=\"alert\">{ props.Error }</p>
        }
    </div>
}
\`\`\`

Create complete molecule using refactored atoms."
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
task_T5_16() { refactor_molecule "card_header" "true"; }
task_T5_17() { refactor_molecule "list_item" "true"; }
task_T5_18() { refactor_molecule "tabs"; }
task_T5_19() { refactor_molecule "dropdown_menu"; }
task_T5_20() { refactor_molecule "input_group" "true"; }

# =============================================================================
# PHASE 6: Organisms Tasks
# =============================================================================

refactor_organism() {
  local component="$1"
  local file="$ORGANISMS_DIR/${component}.templ"
  local create_new="${2:-false}"

  if [[ ! -f "$file" && "$create_new" != "true" ]]; then
    log_warn "File not found: $file - creating new"
    create_new="true"
  fi

  claude --dangerously-skip-permissions "Refactor/Create organism: $file

READ FIRST:
1. Refactored atoms in $ATOMS_DIR/
2. Refactored molecules in $MOLECULES_DIR/
3. $VIEWS_DIR/component-api-standards.md
4. Custom CSS: $CUSTOM_CSS_DIR/organisms/

ORGANISM RULES:
1. Organisms compose molecules + atoms
2. May have significant internal state (Alpine.js)
3. Heavy HTMX integration for server interactions
4. ❌ NO TwMerge, ClassName
5. ✅ Consider breaking into sub-components if >200 lines
6. ✅ Full accessibility support

For complex organisms:
- Use proper Go types for configuration
- Implement slots pattern for customization
- Add comprehensive ARIA attributes
- Support keyboard navigation

Create complete, production-ready organism."
}

task_T6_1() { refactor_organism "datatable"; }
task_T6_2() { refactor_organism "navigation"; }
task_T6_3() { refactor_organism "sidebar_navigation" "true"; }
task_T6_4() { refactor_organism "invoice_card" "true"; }
task_T6_5() { refactor_organism "product_card" "true"; }
task_T6_6() { refactor_organism "form"; }
task_T6_7() { refactor_organism "filter_panel" "true"; }
task_T6_8() { refactor_organism "chart_widget" "true"; }
task_T6_9() { refactor_organism "modal" "true"; }
task_T6_10() { refactor_organism "breadcrumbs" "true"; }
task_T6_11() { refactor_organism "user_menu" "true"; }
task_T6_12() { refactor_organism "notification_list" "true"; }
task_T6_13() { refactor_organism "dropdown" "true"; }
task_T6_14() { refactor_organism "calendar_widget" "true"; }
task_T6_15() { refactor_organism "comment_thread" "true"; }
task_T6_16() { refactor_organism "file_explorer" "true"; }
task_T6_17() { refactor_organism "stats_card" "true"; }
task_T6_18() { refactor_organism "wizard" "true"; }

# =============================================================================
# PHASE 7: Templates Tasks
# =============================================================================

refactor_template() {
  local component="$1"
  local file="$TEMPLATES_DIR/${component}.templ"
  local create_new="${2:-false}"

  if [[ ! -f "$file" && "$create_new" != "true" ]]; then
    log_warn "File not found: $file - creating new"
    create_new="true"
  fi

  claude --dangerously-skip-permissions "Refactor/Create template: $file

READ FIRST:
1. All refactored organisms, molecules, atoms
2. Custom CSS: $CUSTOM_CSS_DIR/layouts/
3. $VIEWS_DIR/component-api-standards.md

TEMPLATE RULES:
1. Templates define page structure/layout
2. Use CSS Grid/Flexbox for layout
3. Implement proper slots for content areas
4. Handle responsive behavior
5. ❌ NO TwMerge, ClassName

Include:
- HTML5 semantic structure
- Meta tags support
- Script/style injection points
- Responsive breakpoints

Example structure:
\`\`\`templ
package templates

type DashboardLayoutProps struct {
    Title string
    User  UserInfo
}

templ DashboardLayout(props DashboardLayoutProps) {
    @BaseLayout(PageMeta{Title: props.Title}) {
        <div class=\"layout-dashboard\">
            <aside class=\"sidebar\">
                { sidebar... }
            </aside>
            <main class=\"main-content\">
                { children... }
            </main>
        </div>
    }
}
\`\`\`

Create complete, responsive template."
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

# =============================================================================
# PHASE 8: Pages Tasks
# =============================================================================

create_page() {
  local component="$1"
  local file="$PAGES_DIR/${component}.templ"

  claude --dangerously-skip-permissions "Create page: $file

READ FIRST:
1. All templates in $TEMPLATES_DIR/
2. All organisms, molecules, atoms
3. Component list reference for page composition

PAGE RULES:
1. Pages compose templates with content
2. Handle data fetching/props
3. Wire up HTMX endpoints
4. Manage page-level state with Alpine.js

Based on the component list, implement this page with:
- Proper template usage
- All required organisms/molecules
- HTMX data loading patterns
- Proper error handling
- Loading states

Create complete, functional page."
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
# PHASE 9: Validation Tasks
# =============================================================================

task_T9_1() {
  log_info "Checking templ compilation..."

  if ! run_templ_generate; then
    log_error "templ generation failed"

    claude --dangerously-skip-permissions "Fix all templ compilation errors.

Run: templ generate 2>&1

Fix ALL errors. Common issues:
1. Import path errors
2. Undefined types
3. Syntax errors
4. Missing closing tags"

    if ! run_templ_generate; then
      return 1
    fi
  fi

  log_success "templ compilation passed!"
}

task_T9_2() {
  log_info "Verifying Basecoat classes..."

  claude --dangerously-skip-permissions "Verify all Basecoat classes used actually exist.

For each .templ file:
1. Extract class=\"...\" values
2. Check against $BASECOAT_CSS
3. Report any missing classes

Create: $VIEWS_DIR/class-validation-report.md"
}

task_T9_3() {
  log_info "Verifying custom CSS..."

  claude --dangerously-skip-permissions "Verify all custom CSS files are complete.

Check $CUSTOM_CSS_DIR for:
1. All expected files exist
2. All classes used in components are defined
3. No syntax errors
4. Dark mode variants present
5. Responsive breakpoints included

Create: $VIEWS_DIR/css-validation-report.md"
}

task_T9_4() {
  log_info "Scanning for remaining anti-patterns..."

  local issues=$(detect_anti_patterns "$VIEWS_DIR")

  if [[ -n "$issues" ]]; then
    log_warn "Anti-patterns found:"
    echo "$issues"

    claude --dangerously-skip-permissions "Remove ALL remaining anti-patterns.

Found:
$issues

Fix every instance."
  fi

  issues=$(detect_anti_patterns "$VIEWS_DIR")
  if [[ -z "$issues" ]]; then
    log_success "No anti-patterns remaining!"
  else
    log_error "Some anti-patterns still present"
    echo "$issues"
  fi
}

task_T9_5() {
  log_info "Testing HTMX integration..."

  claude --dangerously-skip-permissions "Verify HTMX integration patterns.

Check all components for:
1. Proper hx-* attribute usage
2. Target elements exist
3. Swap strategies appropriate
4. Loading indicators defined
5. Error handling present

Create: $VIEWS_DIR/htmx-validation-report.md"
}

task_T9_6() {
  log_info "Testing Alpine.js integration..."

  claude --dangerously-skip-permissions "Verify Alpine.js integration patterns.

Check all components for:
1. x-data initialization
2. x-show/x-if usage
3. Event handling (@click, @keydown)
4. Reactivity patterns
5. No conflicts with HTMX

Create: $VIEWS_DIR/alpine-validation-report.md"
}

task_T9_7() {
  log_info "Creating visual component gallery..."

  claude --dangerously-skip-permissions "Create comprehensive component gallery page.

File: $PAGES_DIR/component_gallery.templ

Include ALL components organized by level:
1. Atoms (24 components)
2. Molecules (20 components)
3. Organisms (18 components)
4. Template previews (12 layouts)

Each component should show:
- All variants
- All sizes
- All states
- Code examples

Make it navigable with sidebar and anchor links."
}

task_T9_8() {
  log_info "Generating final report..."

  claude --dangerously-skip-permissions "Generate comprehensive final report.

Create: $VIEWS_DIR/final-report.md

Include:
1. Executive Summary
2. Components Refactored (by level)
3. Custom CSS Created
4. Anti-Patterns Removed
5. HTMX Integration Status
6. Alpine.js Integration Status
7. Accessibility Compliance
8. Breaking Changes
9. Migration Guide
10. Remaining Work

Update $REFACTOR_GUIDE to mark complete."
}

# =============================================================================
# Prerequisites & Setup
# =============================================================================

check_prerequisites() {
  log_info "Checking prerequisites..."

  local missing=""

  if ! command -v claude &>/dev/null; then
    missing+="- Claude CLI not found\n"
  fi

  if ! command -v go &>/dev/null; then
    missing+="- Go not found\n"
  fi

  if [[ ! -d "$VIEWS_DIR" ]]; then
    missing+="- Views directory not found: $VIEWS_DIR\n"
  fi

  if [[ ! -f "$BASECOAT_CSS" ]]; then
    log_warn "Basecoat CSS not found at $BASECOAT_CSS"
  fi

  if [[ -n "$missing" ]]; then
    log_error "Missing prerequisites:\n$missing"
    exit 1
  fi

  log_success "Prerequisites OK"
}

create_backup() {
  if [[ ! -d "$BACKUP_DIR" ]]; then
    log_info "Creating backup: $BACKUP_DIR"
    cp -r "$VIEWS_DIR" "$BACKUP_DIR"
    log_success "Backup created"
  fi
}

init_refactor_guide() {
  if [[ ! -f "$REFACTOR_GUIDE" ]]; then
    cat >"$REFACTOR_GUIDE" <<'GUIDE_EOF'
# UI Component Refactoring Guide - Full System

## Overview
Complete refactoring of Awo ERP UI components following:
- Basecoat UI patterns
- Atomic Design (24 atoms, 20 molecules, 18 organisms, 12 templates, 15 pages)
- HTMX + Alpine.js integration
- Custom CSS for non-Basecoat components

## Status Dashboard

| Phase | Description | Tasks | Status |
|-------|-------------|-------|--------|
| 1 | Analysis | 7 | 🔵 Pending |
| 2 | Strategy | 5 | 🔵 Pending |
| 3 | Custom CSS | 22 | 🔵 Pending |
| 4 | Atoms | 24 | 🔵 Pending |
| 5 | Molecules | 20 | 🔵 Pending |
| 6 | Organisms | 18 | 🔵 Pending |
| 7 | Templates | 12 | 🔵 Pending |
| 8 | Pages | 15 | 🔵 Pending |
| 9 | Validation | 8 | 🔵 Pending |

**Total: 131 tasks**

## Constraints
- ❌ NO utils.TwMerge
- ❌ NO ClassName props
- ✅ Basecoat semantic classes
- ✅ Custom CSS for non-Basecoat
- ✅ Data attributes for state

## Refactoring Log
<!-- Entries added after each task -->

---
*Last updated: Never*
GUIDE_EOF
  fi
}

# =============================================================================
# Main Entry Point
# =============================================================================

show_help() {
  cat <<'EOF'
UI Component Refactoring Script - Full System

Usage:
  ./component-refactor.sh [OPTIONS]

Options:
  --auto        Run all 131 tasks automatically
  --resume      Resume from last checkpoint
  --status      Show progress (131 tasks across 9 phases)
  --task ID     Run specific task (e.g., --task T4.1)
  --from ID     Run from specific task onwards
  --phase N     Run specific phase (1-9)
  --css-only    Run only CSS generation phase (Phase 3)
  --reset       Reset state and start fresh
  --help        Show this help

Phases:
  1  Analysis (7 tasks)     - Inventory, mapping, anti-patterns
  2  Strategy (5 tasks)     - Standards, types, helpers
  3  Custom CSS (22 tasks)  - CSS for non-Basecoat components
  4  Atoms (24 tasks)       - 24 foundational components
  5  Molecules (20 tasks)   - 20 composed components
  6  Organisms (18 tasks)   - 18 complex components
  7  Templates (12 tasks)   - 12 layout patterns
  8  Pages (15 tasks)       - 15 full pages
  9  Validation (8 tasks)   - Compilation, testing, gallery

Total: 131 tasks

Examples:
  ./component-refactor.sh --auto          # Run everything
  ./component-refactor.sh --phase 3       # CSS only
  ./component-refactor.sh --from T4.1     # Start at atoms

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
    register_all_tasks
    init_state
    exit 0
    ;;
  --auto)
    log_header "Full UI Component Refactoring (131 Tasks)"
    check_prerequisites
    create_backup
    START_TIME=$(date +%s)
    run_all_tasks
    exit $?
    ;;
  --resume)
    log_header "Resuming Refactoring"
    check_prerequisites
    resume_tasks
    exit $?
    ;;
  --task)
    if [[ -z "${2:-}" ]]; then
      log_error "Task ID required"
      exit 1
    fi
    check_prerequisites
    execute_task "$2"
    exit $?
    ;;
  --from)
    if [[ -z "${2:-}" ]]; then
      log_error "Task ID required"
      exit 1
    fi
    log_header "Running from $2"
    check_prerequisites
    create_backup
    run_from_task "$2"
    exit $?
    ;;
  --phase)
    if [[ -z "${2:-}" ]]; then
      log_error "Phase number required (1-9)"
      exit 1
    fi
    check_prerequisites
    run_phase "$2"
    exit $?
    ;;
  --css-only)
    log_header "CSS Generation Only"
    check_prerequisites
    run_phase 3
    exit $?
    ;;
  "")
    log_header "UI Component Refactoring (Interactive)"
    check_prerequisites

    echo ""
    echo "Select mode:"
    echo "  1) Run ALL 131 tasks"
    echo "  2) Resume from checkpoint"
    echo "  3) Show status"
    echo "  4) Run specific phase"
    echo "  5) Run specific task"
    echo "  6) Generate CSS only (Phase 3)"
    echo "  7) Reset and start fresh"
    echo "  q) Quit"
    echo ""
    read -p "Choice: " choice

    case "$choice" in
    1)
      create_backup
      START_TIME=$(date +%s)
      run_all_tasks
      ;;
    2) resume_tasks ;;
    3) show_status ;;
    4)
      read -p "Phase (1-9): " phase
      run_phase "$phase"
      ;;
    5)
      show_status
      read -p "Task ID: " tid
      execute_task "$tid"
      ;;
    6) run_phase 3 ;;
    7)
      reset_state
      echo "Reset complete"
      ;;
    q | Q) exit 0 ;;
    *)
      log_error "Invalid choice"
      exit 1
      ;;
    esac
    ;;
  *)
    log_error "Unknown option: $1"
    show_help
    exit 1
    ;;
  esac
}

cleanup() {
  rm -f tmp/log/templ_generate_$$.log tmp/log/go_build_$$.log 2>/dev/null
}
trap cleanup EXIT

main "$@"
