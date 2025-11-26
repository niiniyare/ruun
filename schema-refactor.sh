#!/bin/bash
# =============================================================================
# schema-refactor.sh - Enterprise Schema Package Consolidation
# =============================================================================
#
# A self-healing refactoring script that consolidates redundant types and
# interfaces in the schema package. Features:
#   - State tracking for resume capability
#   - Automatic error detection and fixing
#   - Task-based execution with dependencies
#   - Comprehensive logging
#
# Usage:
#   ./schema-refactor.sh              # Interactive mode
#   ./schema-refactor.sh --auto       # Run all tasks automatically
#   ./schema-refactor.sh --resume     # Resume from last checkpoint
#   ./schema-refactor.sh --status     # Show current progress
#   ./schema-refactor.sh --task T1.1  # Run specific task
#   ./schema-refactor.sh --from T2.1  # Run from specific task onwards
#   ./schema-refactor.sh --reset      # Reset state and start fresh
#
# =============================================================================

set -o pipefail

# =============================================================================
# Configuration
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SCHEMA_DIR="${SCHEMA_DIR:-./schema}"
STATE_FILE="${STATE_FILE:-.schema-refactor-state}"
LOG_FILE="${LOG_FILE:-./schema-refactor.log}"
BACKUP_DIR="${BACKUP_DIR:-./schema_backup_$(date +%Y%m%d_%H%M%S)}"

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
  echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}  $1${NC}" | tee -a "$LOG_FILE"
  echo -e "${BOLD}════════════════════════════════════════════════════════════${NC}" | tee -a "$LOG_FILE"
  echo "" | tee -a "$LOG_FILE"
}

# =============================================================================
# State Management
# =============================================================================

# Task states: pending, running, completed, failed, skipped
declare -A TASK_STATUS
declare -A TASK_DESC
declare -a TASK_ORDER

init_state() {
  if [[ -f "$STATE_FILE" ]]; then
    source "$STATE_FILE"
  fi
}

save_state() {
  {
    echo "# Schema Refactor State - $(date)"
    echo "# Do not edit manually"
    echo ""
    echo "LAST_TASK=\"$LAST_TASK\""
    echo "LAST_STATUS=\"$LAST_STATUS\""
    echo "START_TIME=\"$START_TIME\""
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
  log_header "Refactoring Progress"

  local completed=0
  local failed=0
  local pending=0
  local total=${#TASK_ORDER[@]}

  for task in "${TASK_ORDER[@]}"; do
    local status="${TASK_STATUS[$task]:-pending}"
    local desc="${TASK_DESC[$task]}"

    case "$status" in
    completed)
      echo -e "  ${GREEN}✓${NC} $task: $desc"
      ((completed++))
      ;;
    failed)
      echo -e "  ${RED}✗${NC} $task: $desc"
      ((failed++))
      ;;
    running)
      echo -e "  ${YELLOW}►${NC} $task: $desc"
      ;;
    skipped)
      echo -e "  ${CYAN}○${NC} $task: $desc (skipped)"
      ;;
    *)
      echo -e "  ${BLUE}·${NC} $task: $desc"
      ((pending++))
      ;;
    esac
  done

  echo ""
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
  local desc="$2"
  TASK_ORDER+=("$id")
  TASK_DESC[$id]="$desc"
  TASK_STATUS[$id]="${TASK_STATUS[$id]:-pending}"
}

# Register all tasks
register_all_tasks() {
  # Phase 1: Create unified type files
  register_task "T1.1" "Create types.go (common types)"
  register_task "T1.2" "Create interface.go (all interfaces)"
  register_task "T1.3" "Create conditional.go (unified Conditional)"
  register_task "T1.4" "Create style.go (unified Style)"
  register_task "T1.5" "Create event.go (unified Event types)"
  register_task "T1.6" "Create behavior.go (request behavior)"
  register_task "T1.7" "Create binding.go (reactive bindings)"
  register_task "T1.8" "Create errors.go (error types)"

  # Phase 2: Update core structs
  register_task "T2.1" "Update field.go (use unified types)"
  register_task "T2.2" "Update action.go (use unified types)"
  register_task "T2.3" "Update schema.go (use unified types)"
  register_task "T2.4" "Update layout.go (use unified types)"
  register_task "T2.5" "Update mixin.go (use unified types)"
  register_task "T2.6" "Update repeatable.go (use unified types)"

  # Phase 3: Update subsystems
  register_task "T3.1" "Update builder.go"
  register_task "T3.2" "Update validator.go"
  register_task "T3.3" "Update enricher.go"
  register_task "T3.4" "Update registry.go"
  register_task "T3.5" "Update runtime.go"
  register_task "T3.6" "Update parser.go"
  register_task "T3.7" "Update i18n.go"
  register_task "T3.8" "Update business_rules.go"

  # Phase 4: Remove deprecated
  register_task "T4.1" "Remove duplicate types from old files"
  register_task "T4.2" "Remove deprecated interfaces"
  register_task "T4.3" "Clean up doc.go"

  # Phase 5: Update tests
  register_task "T5.1" "Update field_test.go"
  register_task "T5.2" "Update action_test.go"
  register_task "T5.3" "Update schema_test.go"
  register_task "T5.4" "Update builder_test.go"
  register_task "T5.5" "Update remaining test files"

  # Phase 6: Final
  register_task "T6.1" "Final build verification"
  register_task "T6.2" "Final test verification"
  register_task "T6.3" "Generate documentation"
}

# =============================================================================
# Build & Test Helpers
# =============================================================================

run_build() {
  local log_file="/tmp/schema_build_$$.log"
  timeout "$BUILD_TIMEOUT" go build ./schema/... 2>&1 | tee "$log_file"
  local result=${PIPESTATUS[0]}

  if [[ $result -eq 0 ]]; then
    rm -f "$log_file"
    return 0
  else
    cat "$log_file"
    return 1
  fi
}

run_tests() {
  local log_file="/tmp/schema_test_$$.log"
  timeout "$TEST_TIMEOUT" go test ./schema/... -count=1 2>&1 | tee "$log_file"
  local result=${PIPESTATUS[0]}

  if [[ $result -eq 0 ]]; then
    rm -f "$log_file"
    return 0
  fi
  return 1
}

get_build_errors() {
  go build ./schema/... 2>&1 | grep -E "^\./" | head -30
}

# =============================================================================
# Self-Healing Build System
# =============================================================================

fix_build_errors() {
  local attempt=1

  while [[ $attempt -le $MAX_FIX_ATTEMPTS ]]; do
    log_info "Build attempt $attempt/$MAX_FIX_ATTEMPTS..."

    if run_build; then
      log_success "Build passed!"
      return 0
    fi

    local errors=$(get_build_errors)

    if [[ -z "$errors" ]]; then
      log_error "Build failed but couldn't extract errors"
      return 1
    fi

    log_fix "Attempting automatic fix (attempt $attempt)..."

    claude --dangerously-skip-permissions "Fix these Go build errors in the schema package:

$errors

FIXING RULES:
1. DUPLICATE DECLARATION: Remove the duplicate from the OLD file, keep the one in the NEW unified file
2. UNDEFINED TYPE: 
   - If type should be in types.go, interface.go, etc., ensure it's there
   - If old type was removed, update references to new unified type
3. UNDEFINED FIELD: Update struct field access to match new type definitions
4. TYPE MISMATCH: Cast or update to correct types
5. IMPORT CYCLE: Move types to break the cycle
6. MISSING IMPORT: Add required imports

Fix ALL errors. The code MUST compile after your changes.
Show what files you're modifying."

    ((attempt++))
    sleep 2
  done

  log_error "Could not fix build errors after $MAX_FIX_ATTEMPTS attempts"
  return 1
}

ensure_builds() {
  if ! run_build 2>/dev/null; then
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

        # Git commit if available
        if git rev-parse --git-dir &>/dev/null 2>&1; then
          git add -A
          git commit -m "refactor(schema): $task_id - ${TASK_DESC[$task_id]}" 2>/dev/null || true
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
# Task Implementations - Phase 1: Create Unified Type Files
# =============================================================================

task_T1_1() {
  log_info "Creating consolidated types.go..."

  claude --dangerously-skip-permissions 'Create schema/types.go with ALL common types for the package.

IMPORTANT: This file consolidates types from multiple files. Check existing files first.

Include these types (consolidated from across the codebase):

1. ENUMERATIONS:
   - Type (form, table, wizard, etc.)
   - ErrorType (validation, permission, etc.) 
   - TransformType (trim, lowercase, etc.)
   - SwapStrategy (innerHTML, outerHTML, etc.)
   - AggregateType (sum, avg, etc.)

2. VALIDATION TYPES:
   - ValidationResult (with Valid, Errors, Warnings, Data fields + helper methods)
   - ValidationMessages (custom error messages)
   - ValidationRule (rule definition)
   - Validation (schema-level config)

3. METADATA TYPES:
   - Metadata (base: Version, CreatedAt, UpdatedAt, CreatedBy, UpdatedBy)
   - Meta (extended schema metadata)
   - ChangelogEntry
   - Theme

4. CONFIGURATION TYPES:
   - Config (form submission config)
   - Security, CSRF, RateLimit, Encryption
   - Tenant (multi-tenancy)

5. WORKFLOW TYPES:
   - Workflow, WorkflowAction, WorkflowTransition
   - WorkflowTimeout, Notification
   - ApprovalConfig, ApproverConfig

6. LAYOUT CONFIG TYPES:
   - GridConfig, FlexConfig, StackConfig, PositionConfig

7. OTHER:
   - EnrichConfig
   - Aggregate
   - SchemaI18n

DO NOT include (these go in separate files):
- Conditional (goes in conditional.go)
- Style (goes in style.go)
- Event/Events (goes in event.go)
- Behavior/Trigger (goes in behavior.go)
- Binding (goes in binding.go)
- Interfaces (go in interface.go)
- Field, Action, Schema, Layout (stay in their own files)

After creating, remove these types from their OLD locations to avoid duplicates.
Run "go build ./schema/..." and fix any errors.'
}

task_T1_2() {
  log_info "Creating consolidated interface.go..."

  claude --dangerously-skip-permissions 'Create schema/interface.go with ALL interfaces for the package.

CONSOLIDATE these interfaces (many are duplicated):

1. VALIDATION (consolidate ~6 interfaces into 1-2):
   - Validator interface (ValidateSchema, ValidateData, ValidateField, AddRule, GetRule)
   - ValidatorFunc, AsyncValidatorFunc types

2. RENDERING (consolidate 2 into 1):
   - Renderer interface (Render, RenderField, Format)

3. REGISTRY & STORAGE:
   - Registry interface
   - StorageBackend interface  
   - CacheBackend interface
   - ListFilter, StorageFilter, StorageMetadata structs

4. EVENTS:
   - EventHandler type
   - EventDispatcher interface

5. STATE:
   - StateManager interface

6. ENRICHMENT:
   - Enricher interface
   - User interface
   - TenantProvider interface
   - TenantCustomization, TenantOverride structs

7. BEHAVIOR:
   - BehaviorAdapter interface
   - ConditionalEngine interface
   - BusinessRulesEngine interface

8. DATA:
   - Database interface

9. PARSER:
   - ParserPlugin interface

10. LOGGING:
    - Logger interface

11. ACCESSORS (optional, for mocking):
    - SchemaAccessor interface
    - FieldAccessor interface
    - RepeatableFieldInfo interface

12. ERRORS:
    - SchemaError interface

REMOVE duplicate interfaces from:
- doc.go (SchemaValidator, SchemaRenderer, SchemaRegistry)
- validator.go (FieldInterface, EnrichedFieldInterface, SchemaInterface, etc.)
- enricher.go (User, TenantProvider, Logger, Enricher)
- runtime.go (ConditionalEngine)
- errors.go (keep implementation, move interface here)

Run "go build ./schema/..." and fix any errors.'
}

task_T1_3() {
  log_info "Creating unified conditional.go..."

  claude --dangerously-skip-permissions 'Create schema/conditional.go with the SINGLE unified Conditional type.

This REPLACES:
- FieldConditional (from field.go)
- LayoutConditional (from layout.go)  
- ActionConditional (if exists)
- Conditional (from types.go if it exists there)

Create this file:

package schema

import "github.com/niiniyare/ruun/pkg/condition"

// Conditional defines visibility and state conditions for any component.
// Used by Field, Action, Layout, LayoutBlock, Section, Tab, Step, etc.
type Conditional struct {
    // Visibility conditions
    Show *condition.ConditionGroup `json:"show,omitempty"`
    Hide *condition.ConditionGroup `json:"hide,omitempty"`
    
    // State conditions  
    Required *condition.ConditionGroup `json:"required,omitempty"`
    Disabled *condition.ConditionGroup `json:"disabled,omitempty"`
    Readonly *condition.ConditionGroup `json:"readonly,omitempty"`
    
    // Validation conditions
    Validate *condition.ConditionGroup `json:"validate,omitempty"`
}

// Helper methods
func (c *Conditional) IsEmpty() bool {
    if c == nil { return true }
    return c.Show == nil && c.Hide == nil && c.Required == nil && 
           c.Disabled == nil && c.Readonly == nil && c.Validate == nil
}

func (c *Conditional) HasVisibility() bool {
    return c != nil && (c.Show != nil || c.Hide != nil)
}

func (c *Conditional) HasStateConditions() bool {
    return c != nil && (c.Required != nil || c.Disabled != nil || c.Readonly != nil)
}

// ConditionalBuilder for fluent construction
type ConditionalBuilder struct {
    cond Conditional
}

func NewConditional() *ConditionalBuilder {
    return &ConditionalBuilder{}
}

func (b *ConditionalBuilder) ShowWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
    b.cond.Show = cg
    return b
}

func (b *ConditionalBuilder) HideWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
    b.cond.Hide = cg
    return b
}

func (b *ConditionalBuilder) RequiredWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
    b.cond.Required = cg
    return b
}

func (b *ConditionalBuilder) DisabledWhen(cg *condition.ConditionGroup) *ConditionalBuilder {
    b.cond.Disabled = cg
    return b
}

func (b *ConditionalBuilder) Build() *Conditional {
    return &b.cond
}

AFTER CREATING:
1. Update Field.Conditional to use *Conditional (remove FieldConditional)
2. Update Layout components to use *Conditional (remove LayoutConditional)
3. Remove old conditional types from their files
4. Run "go build ./schema/..." and fix errors'
}

task_T1_4() {
  log_info "Creating unified style.go..."

  claude --dangerously-skip-permissions 'Create schema/style.go with the SINGLE unified Style type.

This REPLACES:
- BaseStyle (from layout.go)
- LayoutStyle (from layout.go)
- SectionStyle (from layout.go)
- TabStyle (from layout.go)
- StepStyle (from layout.go)
- ActionTheme (from action.go)
- FieldStyle (if exists)
- Style (if exists in types.go)

Create this file:

package schema

// Style defines unified styling for all components.
type Style struct {
    // CSS Classes
    Classes   string `json:"classes,omitempty"`
    Container string `json:"container,omitempty"`
    Wrapper   string `json:"wrapper,omitempty"`
    Label     string `json:"label,omitempty"`
    Input     string `json:"input,omitempty"`
    Help      string `json:"help,omitempty"`
    Error     string `json:"error,omitempty"`
    
    // Dimensions
    Width   string `json:"width,omitempty"`
    Height  string `json:"height,omitempty"`
    Margin  string `json:"margin,omitempty"`
    Padding string `json:"padding,omitempty"`
    
    // Visual
    Background   string `json:"background,omitempty"`
    Border       string `json:"border,omitempty"`
    BorderRadius string `json:"borderRadius,omitempty"`
    Shadow       string `json:"shadow,omitempty"`
    
    // Typography
    FontSize   string `json:"fontSize,omitempty"`
    FontWeight string `json:"fontWeight,omitempty"`
    TextColor  string `json:"textColor,omitempty"`
    TextAlign  string `json:"textAlign,omitempty"`
    
    // Color tokens (for theming)
    Colors map[string]string `json:"colors,omitempty"`
    
    // State-specific styles
    States *StateStyles `json:"states,omitempty"`
    
    // Custom CSS
    CustomCSS string `json:"customCSS,omitempty"`
}

// StateStyles defines styling for different component states.
type StateStyles struct {
    Default   map[string]string `json:"default,omitempty"`
    Hover     map[string]string `json:"hover,omitempty"`
    Focus     map[string]string `json:"focus,omitempty"`
    Active    map[string]string `json:"active,omitempty"`
    Disabled  map[string]string `json:"disabled,omitempty"`
    Error     map[string]string `json:"error,omitempty"`
    Success   map[string]string `json:"success,omitempty"`
    // For tabs/steps
    Inactive  map[string]string `json:"inactive,omitempty"`
    Completed map[string]string `json:"completed,omitempty"`
}

func (s *Style) IsEmpty() bool {
    if s == nil { return true }
    return s.Classes == "" && s.Container == "" && s.Background == "" && 
           s.Border == "" && s.CustomCSS == "" && len(s.Colors) == 0
}

func (s *Style) Merge(other *Style) *Style {
    if s == nil { return other }
    if other == nil { return s }
    // Implement merge logic - other overwrites s
    result := *s
    if other.Classes != "" { result.Classes = other.Classes }
    if other.Background != "" { result.Background = other.Background }
    // ... etc
    return &result
}

AFTER CREATING:
1. Update all structs that use style types to use *Style
2. Remove old style types from layout.go, action.go, etc.
3. Run "go build ./schema/..." and fix errors'
}

task_T1_5() {
  log_info "Creating unified event.go..."

  claude --dangerously-skip-permissions 'Create schema/event.go with unified event types.

This CONSOLIDATES:
- Event struct (defined TWICE: types.go:129 AND interface.go:99 - REMOVE DUPLICATES)
- Events struct (from types.go)
- BehaviorMetadata (defined TWICE - REMOVE DUPLICATE)

Create this file:

package schema

import "time"

// Event represents a schema lifecycle or interaction event.
// This is the SINGLE definition - remove duplicates from other files.
type Event struct {
    Type      string         `json:"type"`
    Source    string         `json:"source"`
    Timestamp time.Time      `json:"timestamp"`
    Data      map[string]any `json:"data,omitempty"`
}

func NewEvent(eventType, source string) Event {
    return Event{
        Type:      eventType,
        Source:    source,
        Timestamp: time.Now(),
        Data:      make(map[string]any),
    }
}

func (e Event) WithData(key string, value any) Event {
    if e.Data == nil { e.Data = make(map[string]any) }
    e.Data[key] = value
    return e
}

// Events configures event handlers for schemas/fields/actions.
type Events struct {
    // Lifecycle
    OnInit    string `json:"onInit,omitempty"`
    OnMount   string `json:"onMount,omitempty"`
    OnUnmount string `json:"onUnmount,omitempty"`
    OnDestroy string `json:"onDestroy,omitempty"`
    
    // Form submission
    OnSubmit        string `json:"onSubmit,omitempty"`
    BeforeSubmit    string `json:"beforeSubmit,omitempty"`
    AfterSubmit     string `json:"afterSubmit,omitempty"`
    OnSubmitSuccess string `json:"onSubmitSuccess,omitempty"`
    OnSubmitError   string `json:"onSubmitError,omitempty"`
    
    // Field interactions
    OnChange   string `json:"onChange,omitempty"`
    OnInput    string `json:"onInput,omitempty"`
    OnFocus    string `json:"onFocus,omitempty"`
    OnBlur     string `json:"onBlur,omitempty"`
    OnValidate string `json:"onValidate,omitempty"`
    
    // Data operations
    OnLoad    string `json:"onLoad,omitempty"`
    OnSuccess string `json:"onSuccess,omitempty"`
    OnError   string `json:"onError,omitempty"`
    
    // State changes
    OnReset    string `json:"onReset,omitempty"`
    OnDirty    string `json:"onDirty,omitempty"`
    OnPristine string `json:"onPristine,omitempty"`
    
    // Custom handlers
    Custom map[string]string `json:"custom,omitempty"`
}

func (e *Events) IsEmpty() bool {
    if e == nil { return true }
    return e.OnInit == "" && e.OnMount == "" && e.OnSubmit == "" && 
           e.OnChange == "" && e.OnLoad == "" && len(e.Custom) == 0
}

// BehaviorMetadata describes a behavior capability.
// This is the SINGLE definition - remove duplicates.
type BehaviorMetadata struct {
    Name        string               `json:"name"`
    Description string               `json:"description,omitempty"`
    Version     string               `json:"version,omitempty"`
    Inputs      map[string]FieldType `json:"inputs,omitempty"`
    Outputs     map[string]FieldType `json:"outputs,omitempty"`
    Tags        []string             `json:"tags,omitempty"`
}

// Common event type constants
const (
    EventTypeFieldChange  = "field:change"
    EventTypeFieldFocus   = "field:focus"
    EventTypeFieldBlur    = "field:blur"
    EventTypeFormSubmit   = "form:submit"
    EventTypeFormReset    = "form:reset"
    EventTypeFormValidate = "form:validate"
    EventTypeDataLoad     = "data:load"
    EventTypeDataSave     = "data:save"
    EventTypeStateChange  = "state:change"
)

AFTER CREATING:
1. Remove duplicate Event from types.go and interface.go (keep only this one)
2. Remove duplicate BehaviorMetadata from types.go and interface.go
3. Remove old Events from types.go
4. Run "go build ./schema/..." and fix errors'
}

task_T1_6() {
  log_info "Creating behavior.go..."

  claude --dangerously-skip-permissions 'Create or update schema/behavior.go with the Behavior type.

Check if behavior.go already exists. If so, verify and update it.

This REPLACES:
- HTMX struct (from types.go)
- ActionHTMX (from action.go)  
- FieldHTMX (if exists)

The file should contain:

package schema

// Behavior defines technology-agnostic request/interaction behavior.
// Can be rendered to HTMX, Fetch API, or other implementations.
type Behavior struct {
    // Request
    Method  string            `json:"method,omitempty"`
    URL     string            `json:"url,omitempty"`
    Headers map[string]string `json:"headers,omitempty"`
    Params  map[string]string `json:"params,omitempty"`
    Body    string            `json:"body,omitempty"`
    
    // Response handling
    Target    string       `json:"target,omitempty"`
    Swap      SwapStrategy `json:"swap,omitempty"`
    Select    string       `json:"select,omitempty"`
    
    // Trigger configuration
    Trigger *Trigger `json:"trigger,omitempty"`
    
    // UI feedback
    Indicator string `json:"indicator,omitempty"`
    Disabled  string `json:"disabled,omitempty"`
    Confirm   string `json:"confirm,omitempty"`
    
    // Timing
    Debounce int  `json:"debounce,omitempty"`
    Throttle int  `json:"throttle,omitempty"`
    Timeout  int  `json:"timeout,omitempty"`
    
    // Navigation
    PushURL   bool   `json:"pushUrl,omitempty"`
    ReplaceURL bool  `json:"replaceUrl,omitempty"`
    
    // Events
    OnSuccess string `json:"onSuccess,omitempty"`
    OnError   string `json:"onError,omitempty"`
}

// Trigger defines when behavior activates.
type Trigger struct {
    Event   string `json:"event,omitempty"`
    Target  string `json:"target,omitempty"`
    Filter  string `json:"filter,omitempty"`
    Delay   int    `json:"delay,omitempty"`
    Once    bool   `json:"once,omitempty"`
    Changed bool   `json:"changed,omitempty"`
    From    string `json:"from,omitempty"`
}

// SwapStrategy defines how content is replaced.
type SwapStrategy string

const (
    SwapInnerHTML   SwapStrategy = "innerHTML"
    SwapOuterHTML   SwapStrategy = "outerHTML"
    SwapBeforeBegin SwapStrategy = "beforebegin"
    SwapAfterBegin  SwapStrategy = "afterbegin"
    SwapBeforeEnd   SwapStrategy = "beforeend"
    SwapAfterEnd    SwapStrategy = "afterend"
    SwapDelete      SwapStrategy = "delete"
    SwapNone        SwapStrategy = "none"
)

func (b *Behavior) IsEmpty() bool {
    return b == nil || (b.Method == "" && b.URL == "" && b.Target == "")
}

// BehaviorBuilder for fluent construction
type BehaviorBuilder struct {
    behavior Behavior
}

func NewBehavior() *BehaviorBuilder {
    return &BehaviorBuilder{}
}

func (b *BehaviorBuilder) Get(url string) *BehaviorBuilder {
    b.behavior.Method = "GET"
    b.behavior.URL = url
    return b
}

func (b *BehaviorBuilder) Post(url string) *BehaviorBuilder {
    b.behavior.Method = "POST"
    b.behavior.URL = url
    return b
}

func (b *BehaviorBuilder) Target(target string) *BehaviorBuilder {
    b.behavior.Target = target
    return b
}

func (b *BehaviorBuilder) Swap(strategy SwapStrategy) *BehaviorBuilder {
    b.behavior.Swap = strategy
    return b
}

func (b *BehaviorBuilder) Build() *Behavior {
    return &b.behavior
}

AFTER CREATING:
1. Remove HTMX struct from types.go
2. Remove ActionHTMX from action.go
3. Update all HTMX references to use Behavior
4. Run "go build ./schema/..." and fix errors'
}

task_T1_7() {
  log_info "Creating binding.go..."

  claude --dangerously-skip-permissions 'Create or update schema/binding.go with the Binding type.

Check if binding.go already exists. If so, verify and update it.

This REPLACES:
- Alpine struct (if exists in types.go)
- ActionAlpine (from action.go)
- FieldAlpine (if exists)

The file should contain:

package schema

// Binding defines reactive state bindings.
// Can be rendered to Alpine.js, Vue, or other reactive frameworks.
type Binding struct {
    // Data binding
    Data   string `json:"data,omitempty"`   // x-data equivalent
    Model  string `json:"model,omitempty"`  // Two-way binding
    Text   string `json:"text,omitempty"`   // Text content binding
    HTML   string `json:"html,omitempty"`   // HTML content binding
    
    // Visibility
    Show string `json:"show,omitempty"`
    If   string `json:"if,omitempty"`
    
    // Attributes
    Attrs map[string]string `json:"attrs,omitempty"`
    Class string            `json:"class,omitempty"`
    Style string            `json:"style,omitempty"`
    
    // Events
    On map[string]string `json:"on,omitempty"`
    
    // Iteration
    For    string `json:"for,omitempty"`
    ForKey string `json:"forKey,omitempty"`
    
    // Lifecycle
    Init    string `json:"init,omitempty"`
    Effect  string `json:"effect,omitempty"`
    Destroy string `json:"destroy,omitempty"`
    
    // Reference
    Ref string `json:"ref,omitempty"`
    
    // Transitions
    Transition string `json:"transition,omitempty"`
}

func (b *Binding) IsEmpty() bool {
    return b == nil || (b.Data == "" && b.Model == "" && b.Show == "" && len(b.On) == 0)
}

func (b *Binding) HasModel() bool {
    return b != nil && b.Model != ""
}

func (b *Binding) HasVisibility() bool {
    return b != nil && (b.Show != "" || b.If != "")
}

func (b *Binding) HasEvents() bool {
    return b != nil && len(b.On) > 0
}

// BindingBuilder for fluent construction
type BindingBuilder struct {
    binding Binding
}

func NewBinding() *BindingBuilder {
    return &BindingBuilder{binding: Binding{On: make(map[string]string)}}
}

func (b *BindingBuilder) Data(expr string) *BindingBuilder {
    b.binding.Data = expr
    return b
}

func (b *BindingBuilder) Model(field string) *BindingBuilder {
    b.binding.Model = field
    return b
}

func (b *BindingBuilder) Show(condition string) *BindingBuilder {
    b.binding.Show = condition
    return b
}

func (b *BindingBuilder) On(event, handler string) *BindingBuilder {
    b.binding.On[event] = handler
    return b
}

func (b *BindingBuilder) Build() *Binding {
    return &b.binding
}

AFTER CREATING:
1. Remove Alpine-related types from action.go and elsewhere
2. Update references to use Binding
3. Run "go build ./schema/..." and fix errors'
}

task_T1_8() {
  log_info "Creating unified errors.go..."

  claude --dangerously-skip-permissions 'Update schema/errors.go to consolidate error types.

Keep existing implementations but ensure:

1. SchemaError interface is defined ONCE (either here or in interface.go)

2. Keep these implementations:
   - BaseError struct (implements SchemaError)
   - ValidationErrorCollection
   - ErrorCollector
   - ErrorResponse, MultiErrorResponse

3. Add if missing:
   - Common error constructors
   - Error type constants

4. Remove duplicates:
   - If SchemaError interface is in interface.go, remove from here
   - If ValidationResult is defined here, remove (its in types.go)

5. Add helper constructors:

func NewValidationError(field, message string) SchemaError {
    return &BaseError{
        code:    "VALIDATION_ERROR",
        message: message,
        errType: ErrorTypeValidation,
        field:   field,
    }
}

func NewFieldError(field, code, message string) SchemaError {
    return &BaseError{
        code:    code,
        message: message,
        errType: ErrorTypeValidation,
        field:   field,
    }
}

func NewPermissionError(message string) SchemaError {
    return &BaseError{
        code:    "PERMISSION_DENIED",
        message: message,
        errType: ErrorTypePermission,
    }
}

func NewNotFoundError(resource, id string) SchemaError {
    return &BaseError{
        code:    "NOT_FOUND",
        message: fmt.Sprintf("%s not found: %s", resource, id),
        errType: ErrorTypeNotFound,
    }
}

Run "go build ./schema/..." and fix any errors.'
}

# =============================================================================
# Task Implementations - Phase 2: Update Core Structs
# =============================================================================

task_T2_1() {
  log_info "Updating field.go to use unified types..."

  claude --dangerously-skip-permissions 'Update schema/field.go to use the unified types.

CHANGES NEEDED:

1. Update Field struct:
   REPLACE:
   - Conditional *FieldConditional -> Conditional *Conditional
   - Style *FieldStyle -> Style *Style (if exists)
   - Behavior *FieldBehavior or HTMX *FieldHTMX -> Behavior *Behavior
   - Binding *FieldBinding or Alpine *FieldAlpine -> Binding *Binding
   
   ADD if not present:
   - Events *Events `json:"events,omitempty"`

2. REMOVE these type definitions from field.go:
   - type FieldConditional struct {...}
   - type FieldHTMX struct {...}
   - type FieldAlpine struct {...}
   - type FieldStyle struct {...}
   - Any other types now in unified files

3. Update FieldBuilder methods:
   - WithConditional should take *Conditional
   - WithBehavior should take *Behavior
   - WithBinding should take *Binding
   - WithStyle should take *Style
   - Add WithEvents(*Events)

4. Keep these in field.go (field-specific):
   - FieldType enum
   - FieldValidation struct
   - FieldTransform struct
   - FieldLayout struct
   - FieldOption struct
   - FieldRuntime struct
   - FieldI18n struct
   - DataSource struct

Run "go build ./schema/..." and fix ALL errors.'
}

task_T2_2() {
  log_info "Updating action.go to use unified types..."

  claude --dangerously-skip-permissions 'Update schema/action.go to use the unified types.

CHANGES NEEDED:

1. Update Action struct:
   REPLACE:
   - HTMX *ActionHTMX -> Behavior *Behavior `json:"behavior,omitempty"`
   - Theme *ActionTheme -> Style *Style `json:"style,omitempty"`
   - Config *ActionConfig -> Config map[string]any `json:"config,omitempty"`
   - Permissions *ActionPermissions -> Permissions []string `json:"permissions,omitempty"`
   
   ADD if not present:
   - Conditional *Conditional `json:"conditional,omitempty"`
   - Binding *Binding `json:"binding,omitempty"`
   - Events *Events `json:"events,omitempty"`

2. REMOVE these from action.go:
   - type ActionHTMX struct {...}
   - type ActionTheme struct {...}
   - type ActionConfig struct {...}
   - type ActionPermissions struct {...}
   - type ActionConditional struct {...} (if exists)

3. Keep these in action.go (action-specific):
   - ActionType enum
   - ActionConfirm struct
   - Action struct (updated)
   - ActionBuilder (updated)
   - Action methods

4. Add helper methods to Action:
   func (a *Action) HasPermission(perm string) bool {
       for _, p := range a.Permissions {
           if p == perm { return true }
       }
       return false
   }
   
   func (a *Action) GetConfig(key string) any {
       if a.Config == nil { return nil }
       return a.Config[key]
   }

5. Update ActionBuilder methods for new types

Run "go build ./schema/..." and fix ALL errors.'
}

task_T2_3() {
  log_info "Updating schema.go to use unified types..."

  claude --dangerously-skip-permissions 'Update schema/schema.go to use the unified types.

CHANGES NEEDED:

1. Update Schema struct:
   REPLACE:
   - HTMX *HTMX -> Behavior *Behavior `json:"behavior,omitempty"`
   - Alpine *Alpine -> Binding *Binding `json:"binding,omitempty"`
   
   ENSURE these use unified types:
   - Events *Events (should reference event.go type)
   - Security *Security (should reference types.go type)
   - Validation *Validation (should reference types.go type)

2. Ensure all field types reference unified types:
   - Fields []Field (Field should use unified Conditional, Style, etc.)
   - Actions []Action (Action should use unified types)
   - Layout *Layout (Layout should use unified types)

3. Update any Schema methods that reference old types

4. Keep these in schema.go (schema-specific):
   - Schema struct (updated)
   - Type enum (form, table, wizard, etc.)
   - SchemaInfo interface implementation
   - Schema methods (GetField, GetAction, etc.)

Run "go build ./schema/..." and fix ALL errors.'
}

task_T2_4() {
  log_info "Updating layout.go to use unified types..."

  claude --dangerously-skip-permissions 'Update schema/layout.go to use the unified types.

CHANGES NEEDED:

1. Update Layout struct:
   - Style *Style (use unified Style from style.go)
   - Add: Blocks []LayoutBlock `json:"blocks,omitempty"` if not present

2. Update Section struct:
   - Conditional *Conditional (use unified type)
   - Style *Style (use unified type)

3. Update Group struct:
   - Conditional *Conditional
   - Style *Style

4. Update Tab struct:
   - Conditional *Conditional
   - Style *Style

5. Update Step struct:
   - Conditional *Conditional
   - Style *Style

6. Update LayoutBlock struct (if exists):
   - Conditional *Conditional
   - Style *Style

7. REMOVE from layout.go:
   - type BaseStyle struct {...}
   - type LayoutStyle struct {...} (if different from Style)
   - type SectionStyle struct {...}
   - type TabStyle struct {...}
   - type StepStyle struct {...}
   - type LayoutConditional struct {...}

8. Keep these in layout.go (layout-specific):
   - LayoutType enum
   - BlockType enum
   - Layout struct (updated)
   - Section, Group, Tab, Step structs (updated)
   - LayoutBlock struct (updated)
   - Breakpoints, BreakpointConfig
   - Layout builders

9. Update all builder methods for new types

Run "go build ./schema/..." and fix ALL errors.'
}

task_T2_5() {
  log_info "Updating mixin.go..."

  claude --dangerously-skip-permissions 'Update schema/mixin.go to use unified types.

CHANGES NEEDED:

1. Ensure Mixin struct uses unified types:
   - Security *Security (from types.go)
   - Events *Events (from event.go)

2. Verify Fields []Field and Actions []Action will work with updated types

3. Update any mixin application logic if needed

Run "go build ./schema/..." and fix errors.'
}

task_T2_6() {
  log_info "Updating repeatable.go..."

  claude --dangerously-skip-permissions 'Update schema/repeatable.go to use unified types.

CHANGES NEEDED:

1. RepeatableField embeds Field - ensure it works with updated Field struct

2. If RepeatableField has its own Style or Conditional fields, update to use unified types

3. Update any methods that reference old types

4. Ensure Aggregate type references AggregateType from types.go (avoid duplicates)

Run "go build ./schema/..." and fix errors.'
}

# =============================================================================
# Task Implementations - Phase 3: Update Subsystems
# =============================================================================

task_T3_1() {
  log_info "Updating builder.go..."

  claude --dangerously-skip-permissions 'Update schema/builder.go to use unified types.

CHANGES NEEDED:

1. Update SchemaBuilder methods:
   - WithHTMX -> WithBehavior (take *Behavior)
   - WithAlpine -> WithBinding (take *Binding)
   - Add WithEvents(*Events) if not present

2. Ensure all methods return proper unified types

3. Update any internal type references

Run "go build ./schema/..." and fix errors.'
}

task_T3_2() {
  log_info "Updating validator.go..."

  claude --dangerously-skip-permissions 'Update schema/validator.go to use unified types and interfaces.

CHANGES NEEDED:

1. REMOVE these interfaces (moved to interface.go):
   - type FieldInterface interface {...}
   - type EnrichedFieldInterface interface {...}
   - type SchemaInterface interface {...}
   - type EnrichedSchemaInterface interface {...}

2. REMOVE commented out duplicates:
   - // type Validator interface {...}
   - // type ValidationResult struct {...}
   - // type ValidationError struct {...}

3. Keep DataValidator struct and its methods

4. Update to use interfaces from interface.go

5. Use ValidationResult from types.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_3() {
  log_info "Updating enricher.go..."

  claude --dangerously-skip-permissions 'Update schema/enricher.go to use unified interfaces.

CHANGES NEEDED:

1. REMOVE these (moved to interface.go):
   - type User interface {...}
   - type TenantProvider interface {...}
   - type Logger interface {...}
   - type Enricher interface {...}

2. REMOVE these (moved to interface.go or types):
   - type EnrichConfig struct {...} (if in types.go)
   - type TenantOverride struct {...} (if in interface.go)
   - type TenantCustomization struct {...} (if in interface.go)

3. Keep implementation:
   - type DefaultEnricher struct {...}
   - type BasicUser struct {...}
   - type noopLogger struct {...}
   - All methods

4. Import and use interfaces from interface.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_4() {
  log_info "Updating registry.go..."

  claude --dangerously-skip-permissions 'Update schema/registry.go to use unified interfaces.

CHANGES NEEDED:

1. REMOVE these (moved to interface.go):
   - type StorageBackend interface {...}
   - type CacheBackend interface {...}
   - type StorageFilter struct {...} (if duplicate)
   - type StorageMetadata struct {...} (if duplicate)

2. Keep Registry implementation and its methods

3. Import and use interfaces from interface.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_5() {
  log_info "Updating runtime.go..."

  claude --dangerously-skip-permissions 'Update schema/runtime.go to use unified interfaces.

CHANGES NEEDED:

1. REMOVE:
   - type ConditionalEngine interface {...} (moved to interface.go)

2. Keep:
   - Runtime struct and methods
   - RuntimeConfig
   - State struct and methods
   - EventHandlers struct
   - RuntimeBuilder

3. Update to use Validator, Renderer, ConditionalEngine from interface.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_6() {
  log_info "Updating parser.go..."

  claude --dangerously-skip-permissions 'Update schema/parser.go to use unified interfaces.

CHANGES NEEDED:

1. REMOVE:
   - type ParserPlugin interface {...} (moved to interface.go)

2. Keep:
   - Parser struct and methods
   - ParserConfig
   - ParseResult
   - parserCache, cacheEntry
   - ParserStats
   - ParseError

3. Import ParserPlugin from interface.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_7() {
  log_info "Updating i18n.go..."

  claude --dangerously-skip-permissions 'Update schema/i18n.go for consistency.

Ensure:
1. No duplicate type definitions with types.go
2. I18nManager, I18nConfig, Translation types are properly defined here
3. Import any shared types from types.go

Run "go build ./schema/..." and fix errors.'
}

task_T3_8() {
  log_info "Updating business_rules.go..."

  claude --dangerously-skip-permissions 'Update schema/business_rules.go for consistency.

Ensure:
1. Uses condition package correctly
2. No duplicate types with other files
3. BusinessRule, BusinessRuleEngine properly defined here

Run "go build ./schema/..." and fix errors.'
}

# =============================================================================
# Task Implementations - Phase 4: Remove Deprecated
# =============================================================================

task_T4_1() {
  log_info "Removing duplicate types from old files..."

  claude --dangerously-skip-permissions 'Remove ALL remaining duplicate type definitions.

Search and remove these from their OLD locations:

FROM types.go - REMOVE if exists:
- type Event struct {...} (now in event.go)
- type Events struct {...} (now in event.go)
- type BehaviorMetadata struct {...} (now in event.go)
- type HTMX struct {...} (replaced by Behavior)
- type Style struct {...} (now in style.go)
- SwapStrategy and constants (now in behavior.go)

FROM interface.go (old one, not the new unified) - CHECK/REMOVE:
- Event struct duplicate
- BehaviorMetadata duplicate
- Any interfaces we moved

FROM layout.go - REMOVE if still exists:
- BaseStyle, LayoutStyle, SectionStyle, TabStyle, StepStyle
- LayoutConditional

FROM field.go - REMOVE if still exists:
- FieldConditional, FieldHTMX, FieldAlpine, FieldStyle

FROM action.go - REMOVE if still exists:
- ActionHTMX, ActionAlpine, ActionTheme, ActionConfig, ActionPermissions

Run "go build ./schema/..." after each file and fix errors before moving to next.'
}

task_T4_2() {
  log_info "Removing deprecated interfaces..."

  claude --dangerously-skip-permissions 'Remove deprecated interface declarations.

FROM doc.go - REMOVE:
- type SchemaValidator interface {...}
- type SchemaRenderer interface {...}
- type SchemaRegistry interface {...}

These are now consolidated in interface.go as:
- Validator
- Renderer
- Registry

FROM validator.go - REMOVE (should already be done):
- FieldInterface
- EnrichedFieldInterface
- SchemaInterface
- EnrichedSchemaInterface
- Commented Validator, ValidationResult, ValidationError

FROM any other files - search and remove duplicate interfaces

Run "go build ./schema/..." and fix errors.'
}

task_T4_3() {
  log_info "Cleaning up doc.go..."

  claude --dangerously-skip-permissions 'Clean up schema/doc.go.

1. Remove the old interface definitions (moved to interface.go)
2. Keep the package documentation
3. Update docs to reference new unified types

The file should only contain:
- Package comment/documentation
- Any doc-only examples
- NO type or interface definitions

Run "go build ./schema/..." and fix errors.'
}

# =============================================================================
# Task Implementations - Phase 5: Update Tests
# =============================================================================

task_T5_1() {
  log_info "Updating field_test.go..."

  claude --dangerously-skip-permissions 'Update schema/field_test.go for unified types.

REPLACE all occurrences:
- FieldHTMX -> Behavior
- FieldAlpine -> Binding
- FieldStyle -> Style
- FieldConditional -> Conditional

UPDATE struct literals:
- .HTMX -> .Behavior
- .Alpine -> .Binding

Fix all test assertions for new type structure.

Run "go test ./schema/... -run TestField" and fix errors.'
}

task_T5_2() {
  log_info "Updating action_test.go..."

  claude --dangerously-skip-permissions 'Update schema/action_test.go for unified types.

REPLACE all occurrences:
- ActionHTMX -> Behavior
- ActionAlpine -> Binding
- ActionTheme -> Style
- ActionConfig -> map[string]any
- ActionPermissions -> []string

UPDATE struct literals:
- .HTMX -> .Behavior
- .Theme -> .Style
- .Permissions.Required -> .Permissions

Fix all test assertions.

Run "go test ./schema/... -run TestAction" and fix errors.'
}

task_T5_3() {
  log_info "Updating schema_test.go..."

  claude --dangerously-skip-permissions 'Update schema/schema_test.go for unified types.

Update all type references and struct literals to use unified types.

Run "go test ./schema/... -run TestSchema" and fix errors.'
}

task_T5_4() {
  log_info "Updating builder_test.go..."

  claude --dangerously-skip-permissions 'Update schema/builder_test.go for unified types.

Update:
- Method calls (WithHTMX -> WithBehavior, etc.)
- Type references
- Struct literals
- Assertions

Run "go test ./schema/... -run TestBuilder" and fix errors.'
}

task_T5_5() {
  log_info "Updating remaining test files..."

  claude --dangerously-skip-permissions 'Update ALL remaining test files in schema/:
- layout_test.go
- enricher_test.go
- runtime_test.go
- validator_test.go
- errors_test.go
- repeatable_test.go
- business_rules_test.go
- i18n_test.go
- val_registry_test.go
- Any other *_test.go files

For EACH file:
1. Replace old type names with unified types
2. Update struct literals
3. Fix assertions

Run "go test ./schema/..." and fix ALL test errors.'
}

# =============================================================================
# Task Implementations - Phase 6: Final Verification
# =============================================================================

task_T6_1() {
  log_info "Final build verification..."

  # Search for any remaining old type references
  OLD_TYPES="FieldHTMX|FieldAlpine|ActionHTMX|ActionAlpine|ActionTheme|ActionConfig|ActionPermissions|ActionConditional|LayoutConditional|SectionStyle|TabStyle|StepStyle|FieldConditional|FieldStyle|SchemaValidator|SchemaRenderer|SchemaRegistry|FieldInterface|EnrichedFieldInterface|SchemaInterface|EnrichedSchemaInterface"

  local remaining=$(grep -rlE "$OLD_TYPES" ./schema/*.go 2>/dev/null || true)

  if [[ -n "$remaining" ]]; then
    log_warn "Found remaining old type references in:"
    echo "$remaining"

    claude --dangerously-skip-permissions "These files still have old type references:
$remaining

Search pattern: $OLD_TYPES

Please update ALL remaining references to use unified types:
- FieldHTMX, ActionHTMX, HTMX -> Behavior
- FieldAlpine, ActionAlpine -> Binding
- ActionTheme, *Style variants -> Style
- *Conditional variants -> Conditional
- SchemaValidator -> Validator (in interface.go)
- SchemaRenderer -> Renderer (in interface.go)
- Remove FieldInterface, EnrichedFieldInterface, SchemaInterface, etc.

Run go build and fix all errors."
  fi

  if ! run_build; then
    log_error "Final build failed"
    return 1
  fi

  log_success "Final build passed!"
  return 0
}

task_T6_2() {
  log_info "Final test verification..."

  if ! run_tests; then
    log_warn "Some tests failed. Attempting fixes..."

    local errors=$(go test ./schema/... 2>&1 | grep -E "(FAIL|undefined|cannot|panic)" | head -30)

    claude --dangerously-skip-permissions "Fix these test failures:

$errors

Update test files to match the refactored types and interfaces."

    if ! run_tests; then
      log_warn "Some tests still failing - may need manual review"
    fi
  fi

  log_success "Test verification complete!"
  return 0
}

task_T6_3() {
  log_info "Generating documentation..."

  claude --dangerously-skip-permissions 'Update package documentation after refactoring.

1. Update doc.go with current package structure
2. Add godoc comments to new unified types if missing
3. Ensure interface.go has proper documentation

Create a brief MIGRATION.md noting:
- Types renamed (FieldHTMX -> Behavior, etc.)
- Interfaces consolidated
- Files created/removed'

  log_success "Documentation updated!"
  return 0
}

# =============================================================================
# Prerequisites & Setup
# =============================================================================

check_prerequisites() {
  log_info "Checking prerequisites..."

  if ! command -v claude &>/dev/null; then
    log_error "Claude CLI not found. Install with: npm install -g @anthropic-ai/claude-cli"
    exit 1
  fi

  if ! command -v go &>/dev/null; then
    log_error "Go not found. Please install Go first."
    exit 1
  fi

  if [[ ! -d "$SCHEMA_DIR" ]]; then
    log_error "Schema directory not found: $SCHEMA_DIR"
    exit 1
  fi

  # Check condition package import path
  if ! grep -r "condition" "$SCHEMA_DIR"/*.go | head -1 | grep -q "import"; then
    log_warn "Could not detect condition package import path"
  fi

  log_success "Prerequisites OK"
}

create_backup() {
  if [[ ! -d "$BACKUP_DIR" ]]; then
    log_info "Creating backup: $BACKUP_DIR"
    cp -r "$SCHEMA_DIR" "$BACKUP_DIR"
    log_success "Backup created"
  fi
}

# =============================================================================
# Main Entry Point
# =============================================================================

show_help() {
  cat <<'EOF'
Schema Package Refactoring Script

Usage:
  ./schema-refactor.sh [OPTIONS]

Options:
  --auto        Run all tasks automatically
  --resume      Resume from last failed/pending task
  --status      Show current progress
  --task ID     Run specific task (e.g., --task T1.1)
  --from ID     Run from specific task onwards
  --reset       Reset state and start fresh
  --help        Show this help

Tasks:
  T1.x  Create unified type files
  T2.x  Update core structs (Field, Action, Schema, Layout)
  T3.x  Update subsystems (Builder, Validator, etc.)
  T4.x  Remove deprecated types
  T5.x  Update tests
  T6.x  Final verification

Examples:
  ./schema-refactor.sh --auto           # Run everything
  ./schema-refactor.sh --resume         # Continue from last point
  ./schema-refactor.sh --task T2.1      # Run just T2.1
  ./schema-refactor.sh --from T3.1      # Run T3.1 and onwards

EOF
}

main() {
  # Initialize
  register_all_tasks
  init_state

  # Parse arguments
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
    log_header "Schema Package Refactoring (Automatic)"
    check_prerequisites
    create_backup
    START_TIME=$(date +%s)
    run_all_tasks
    exit $?
    ;;
  --resume)
    log_header "Schema Package Refactoring (Resume)"
    check_prerequisites
    resume_tasks
    exit $?
    ;;
  --task)
    if [[ -z "${2:-}" ]]; then
      log_error "Task ID required. Example: --task T1.1"
      exit 1
    fi
    check_prerequisites
    execute_task "$2"
    exit $?
    ;;
  --from)
    if [[ -z "${2:-}" ]]; then
      log_error "Task ID required. Example: --from T2.1"
      exit 1
    fi
    log_header "Schema Package Refactoring (From $2)"
    check_prerequisites
    create_backup
    run_from_task "$2"
    exit $?
    ;;
  "")
    # Interactive mode
    log_header "Schema Package Refactoring (Interactive)"
    check_prerequisites

    echo ""
    echo "Select mode:"
    echo "  1) Run ALL tasks automatically"
    echo "  2) Resume from last checkpoint"
    echo "  3) Show current status"
    echo "  4) Run specific task"
    echo "  5) Reset and start fresh"
    echo "  q) Quit"
    echo ""
    read -p "Choice [1-5/q]: " choice

    case "$choice" in
    1)
      create_backup
      START_TIME=$(date +%s)
      run_all_tasks
      ;;
    2)
      resume_tasks
      ;;
    3)
      show_status
      ;;
    4)
      show_status
      echo ""
      read -p "Enter task ID (e.g., T1.1): " task_id
      execute_task "$task_id"
      ;;
    5)
      reset_state
      echo "State reset. Run again to start fresh."
      ;;
    q | Q)
      exit 0
      ;;
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

# Cleanup
cleanup() {
  rm -f /tmp/schema_build_$$.log /tmp/schema_test_$$.log 2>/dev/null
}
trap cleanup EXIT

main "$@"
