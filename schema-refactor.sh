#!/bin/bash
# schema-refactor-v2.sh
# Self-healing refactoring script that fixes errors automatically

set -o pipefail # Catch pipe failures

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Configuration
SCHEMA_DIR="./schema"
BACKUP_DIR="./schema_backup_$(date +%Y%m%d_%H%M%S)"
MAX_FIX_ATTEMPTS=5
BUILD_LOG="/tmp/schema_build_$$.log"
TEST_LOG="/tmp/schema_test_$$.log"

# Logging
log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }
log_fix() { echo -e "${CYAN}[FIX]${NC} $1"; }

# ============================================================================
# SELF-HEALING BUILD SYSTEM
# ============================================================================

# Run build and capture errors
run_build() {
  go build ./schema/... 2>&1 | tee "$BUILD_LOG"
  return ${PIPESTATUS[0]}
}

# Run tests and capture errors
run_tests() {
  go test ./schema/... 2>&1 | tee "$TEST_LOG"
  return ${PIPESTATUS[0]}
}

# Extract error messages from build log
get_build_errors() {
  cat "$BUILD_LOG" | grep -E "^\./" | head -20
}

# Fix build errors automatically
fix_build_errors() {
  local attempt=1

  while [ $attempt -le $MAX_FIX_ATTEMPTS ]; do
    log_info "Build attempt $attempt of $MAX_FIX_ATTEMPTS..."

    if run_build; then
      log_success "Build passed!"
      return 0
    fi

    local errors=$(get_build_errors)

    if [ -z "$errors" ]; then
      log_error "Build failed but couldn't extract errors"
      return 1
    fi

    log_fix "Attempting to fix build errors (attempt $attempt)..."

    claude "The Go build failed with these errors:

$errors

Please fix ALL these errors. Common issues and solutions:

1. DUPLICATE DECLARATION: 
   - If a type/func is defined in multiple files, remove the duplicate from one file
   - Check both the new unified files AND old files for duplicates
   
2. UNDEFINED TYPE:
   - If using Conditional, Behavior, Binding, Style, Events - ensure the unified type files exist
   - If old types like FieldHTMX, ActionAlpine are referenced, update them to use new unified types
   - Add missing imports if needed

3. UNDEFINED FIELD:
   - Update struct field access to match new unified types
   - Old: field.HTMX.Target -> New: field.Behavior.Target
   - Old: field.Alpine.XModel -> New: field.Binding.Model

4. TYPE MISMATCH:
   - Update type references: *FieldHTMX -> *Behavior, *ActionAlpine -> *Binding, etc.
   - Update style types: *LayoutStyle -> *Style, *SectionStyle -> *Style

5. IMPORT CYCLE:
   - Move types to appropriate files to break cycles
   - Use interfaces instead of concrete types where needed

Fix each error systematically. After fixing, the code must compile.
Do NOT just comment out code - actually fix the root cause.
Make all necessary changes across all affected files."

    ((attempt++))
    sleep 1
  done

  log_error "Failed to fix build errors after $MAX_FIX_ATTEMPTS attempts"
  return 1
}

# Fix test errors automatically
fix_test_errors() {
  local attempt=1

  while [ $attempt -le $MAX_FIX_ATTEMPTS ]; do
    log_info "Test attempt $attempt of $MAX_FIX_ATTEMPTS..."

    if run_tests; then
      log_success "Tests passed!"
      return 0
    fi

    local errors=$(cat "$TEST_LOG" | grep -E "(FAIL|panic|undefined|cannot)" | head -20)

    if [ -z "$errors" ]; then
      log_warn "Some tests failed but no critical errors found"
      return 0 # Continue anyway
    fi

    log_fix "Attempting to fix test errors (attempt $attempt)..."

    claude "The Go tests failed with these errors:

$errors

Please fix ALL test errors:

1. UNDEFINED TYPE in tests:
   - Update test files to use new unified types
   - FieldHTMX -> Behavior
   - FieldAlpine -> Binding  
   - ActionHTMX -> Behavior
   - ActionAlpine -> Binding
   - *Style variants -> Style

2. STRUCT FIELD errors:
   - Update field names to match new unified structs
   - Update test assertions accordingly

3. NIL POINTER:
   - Initialize required struct fields in tests
   - Add nil checks where needed

Fix all test files in schema/*_test.go to use the new unified types."

    ((attempt++))
    sleep 1
  done

  log_warn "Some tests may still be failing - manual review needed"
  return 0
}

# Main check function with auto-fix
check_and_fix_build() {
  if ! fix_build_errors; then
    log_error "Could not fix build errors automatically"
    echo ""
    echo "Manual intervention needed. Check the errors above."
    echo "After fixing manually, re-run the script."
    return 1
  fi
  return 0
}

check_and_fix_tests() {
  fix_test_errors
  return $?
}

# ============================================================================
# SMART PHASE EXECUTION
# ============================================================================

# Execute a phase with error recovery
execute_phase() {
  local phase_name="$1"
  local phase_desc="$2"
  local phase_prompt="$3"

  log_info "=== $phase_desc ==="

  # Run the Claude command
  claude "$phase_prompt"

  # Check and fix build
  if ! check_and_fix_build; then
    return 1
  fi

  # Commit if successful
  if git rev-parse --git-dir >/dev/null 2>&1; then
    git add -A
    git commit -m "refactor(schema): $phase_desc" 2>/dev/null || true
  fi

  log_success "Completed: $phase_desc"
  return 0
}

# ============================================================================
# PREREQUISITES
# ============================================================================

check_prerequisites() {
  log_info "Checking prerequisites..."

  if ! command -v claude &>/dev/null; then
    log_error "Claude CLI not found. Please install it first."
    exit 1
  fi

  if ! command -v go &>/dev/null; then
    log_error "Go not found. Please install it first."
    exit 1
  fi

  if [ ! -d "$SCHEMA_DIR" ]; then
    log_error "Schema directory not found: $SCHEMA_DIR"
    exit 1
  fi

  log_success "Prerequisites check passed"
}

backup_schema() {
  log_info "Creating backup at $BACKUP_DIR..."
  cp -r "$SCHEMA_DIR" "$BACKUP_DIR"
  log_success "Backup created"
}

# ============================================================================
# PHASE DEFINITIONS (with self-healing prompts)
# ============================================================================

phase_1_1() {
  execute_phase "1.1" "Create behavior.go" '
Create file schema/behavior.go with the Behavior type for technology-agnostic request handling.

IMPORTANT: 
- First check if behavior.go already exists. If it does, verify its content is correct or update it.
- If Behavior type exists elsewhere, remove the duplicate.
- Ensure no duplicate type declarations.

The file should contain:

package schema

// Behavior defines technology-agnostic request/interaction behavior.
// Replaces HTMX, ActionHTMX, FieldHTMX
type Behavior struct {
    Method    string            `json:"method,omitempty"`
    URL       string            `json:"url,omitempty"`
    Target    string            `json:"target,omitempty"`
    Swap      SwapStrategy      `json:"swap,omitempty"`
    Trigger   *Trigger          `json:"trigger,omitempty"`
    Headers   map[string]string `json:"headers,omitempty"`
    Params    map[string]string `json:"params,omitempty"`
    Indicator string            `json:"indicator,omitempty"`
    Disabled  string            `json:"disabled,omitempty"`
    Debounce  int               `json:"debounce,omitempty"`
    Throttle  int               `json:"throttle,omitempty"`
    Timeout   int               `json:"timeout,omitempty"`
    Confirm   string            `json:"confirm,omitempty"`
    PushURL   bool              `json:"pushUrl,omitempty"`
    Select    string            `json:"select,omitempty"`
}

type Trigger struct {
    Event   string `json:"event,omitempty"`
    Target  string `json:"target,omitempty"`
    Filter  string `json:"filter,omitempty"`
    Delay   int    `json:"delay,omitempty"`
    Once    bool   `json:"once,omitempty"`
    Changed bool   `json:"changed,omitempty"`
    From    string `json:"from,omitempty"`
}

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

After creating/updating, run "go build ./schema/..." and fix any errors.
'
}

phase_1_2() {
  execute_phase "1.2" "Create binding.go" '
Create file schema/binding.go with the Binding type for reactive state management.

IMPORTANT:
- First check if binding.go or Binding type already exists. Update if needed, avoid duplicates.
- Remove any duplicate Binding declarations from other files.

The file should contain:

package schema

// Binding defines reactive state bindings.
// Replaces Alpine, ActionAlpine, FieldAlpine
type Binding struct {
    Model      string            `json:"model,omitempty"`
    Text       string            `json:"text,omitempty"`
    HTML       string            `json:"html,omitempty"`
    Show       string            `json:"show,omitempty"`
    If         string            `json:"if,omitempty"`
    Attrs      map[string]string `json:"attrs,omitempty"`
    Class      string            `json:"class,omitempty"`
    Style      string            `json:"style,omitempty"`
    On         map[string]string `json:"on,omitempty"`
    For        string            `json:"for,omitempty"`
    ForKey     string            `json:"forKey,omitempty"`
    Init       string            `json:"init,omitempty"`
    Effect     string            `json:"effect,omitempty"`
    Destroy    string            `json:"destroy,omitempty"`
    Ref        string            `json:"ref,omitempty"`
    Transition string            `json:"transition,omitempty"`
    Data       string            `json:"data,omitempty"`
}

func (b *Binding) HasModel() bool {
    return b != nil && b.Model != ""
}

func (b *Binding) HasVisibility() bool {
    return b != nil && (b.Show != "" || b.If != "")
}

After creating/updating, run "go build ./schema/..." and fix any errors.
'
}

phase_1_3() {
  execute_phase "1.3" "Create unified Conditional type" '
Create or update the unified Conditional type in schema package.

IMPORTANT:
- Check if Conditional already exists in any file (conditional.go, field.go, action.go, layout.go, types.go)
- If FieldConditional, ActionConditional, LayoutConditional exist, we will replace them with unified Conditional
- For now, just create the unified Conditional type. We will update references later.
- Put it in a new file schema/conditional_unified.go to avoid conflicts initially

Create schema/conditional_unified.go:

package schema

import "github.com/niiniyare/ruun/pkg/condition" 

// Conditional defines unified conditions for any component.
// This will replace FieldConditional, ActionConditional, LayoutConditional
type Conditional struct {
    Show     *condition.ConditionGroup `json:"show,omitempty"`
    Hide     *condition.ConditionGroup `json:"hide,omitempty"`
    Enable   *condition.ConditionGroup `json:"enable,omitempty"`
    Disable  *condition.ConditionGroup `json:"disable,omitempty"`
    Required *condition.ConditionGroup `json:"required,omitempty"`
    Readonly *condition.ConditionGroup `json:"readonly,omitempty"`
    Validate *condition.ConditionGroup `json:"validate,omitempty"`
}

func (c *Conditional) IsEmpty() bool {
    if c == nil {
        return true
    }
    return c.Show == nil && c.Hide == nil && c.Enable == nil && 
        c.Disable == nil && c.Required == nil && c.Readonly == nil
}

CRITICAL: Find the correct import path for the condition package by looking at existing imports in field.go or action.go, then use that same path.

After creating, run "go build ./schema/..." and fix any errors including import path issues.
'
}

phase_1_4() {
  execute_phase "1.4" "Create unified Style type" '
Create the unified Style type in schema package.

IMPORTANT:
- Check existing style types: BaseStyle, LayoutStyle, SectionStyle, TabStyle, StepStyle, ActionTheme, FieldStyle
- Create a new unified Style type that covers all use cases
- Put it in schema/style_unified.go to avoid initial conflicts

Create schema/style_unified.go:

package schema

// Style defines unified styling for all components.
// Replaces LayoutStyle, SectionStyle, TabStyle, StepStyle, ActionTheme, FieldStyle
type Style struct {
    // CSS Classes
    Classes   string `json:"classes,omitempty"`
    Container string `json:"container,omitempty"`
    Label     string `json:"label,omitempty"`
    Input     string `json:"input,omitempty"`
    Help      string `json:"help,omitempty"`
    Error     string `json:"error,omitempty"`

    // Layout
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

    // Color mappings
    Colors map[string]string `json:"colors,omitempty"`

    // State styles
    States *StateStyles `json:"states,omitempty"`

    // Custom CSS
    Custom string `json:"custom,omitempty"`
}

type StateStyles struct {
    Default   map[string]string `json:"default,omitempty"`
    Hover     map[string]string `json:"hover,omitempty"`
    Focus     map[string]string `json:"focus,omitempty"`
    Active    map[string]string `json:"active,omitempty"`
    Disabled  map[string]string `json:"disabled,omitempty"`
    Completed map[string]string `json:"completed,omitempty"`
}

func (s *Style) IsEmpty() bool {
    if s == nil {
        return true
    }
    return s.Classes == "" && s.Container == "" && s.Background == "" && 
        s.Border == "" && s.Custom == "" && len(s.Colors) == 0
}

After creating, run "go build ./schema/..." and fix any errors.
'
}

phase_1_5() {
  execute_phase "1.5" "Create unified Events type" '
Create the unified Events type in schema package.

IMPORTANT:
- Check if Events type already exists in types.go
- If it does, we need to either update it or create EventsUnified and migrate later
- Put it in schema/events_unified.go to avoid initial conflicts

Create schema/events_unified.go:

package schema

// EventsUnified defines event handlers for all components.
// This will eventually replace the Events type in types.go
type EventsUnified struct {
    // Lifecycle
    OnInit    string `json:"onInit,omitempty"`
    OnMount   string `json:"onMount,omitempty"`
    OnDestroy string `json:"onDestroy,omitempty"`

    // Form
    OnSubmit        string `json:"onSubmit,omitempty"`
    BeforeSubmit    string `json:"beforeSubmit,omitempty"`
    AfterSubmit     string `json:"afterSubmit,omitempty"`
    OnSubmitSuccess string `json:"onSubmitSuccess,omitempty"`
    OnSubmitError   string `json:"onSubmitError,omitempty"`
    OnReset         string `json:"onReset,omitempty"`

    // Field
    OnChange   string `json:"onChange,omitempty"`
    OnInput    string `json:"onInput,omitempty"`
    OnFocus    string `json:"onFocus,omitempty"`
    OnBlur     string `json:"onBlur,omitempty"`
    OnValidate string `json:"onValidate,omitempty"`

    // Data
    OnLoad    string `json:"onLoad,omitempty"`
    OnSuccess string `json:"onSuccess,omitempty"`
    OnError   string `json:"onError,omitempty"`

    // State
    OnDirty    string `json:"onDirty,omitempty"`
    OnPristine string `json:"onPristine,omitempty"`

    // Custom
    Custom map[string]string `json:"custom,omitempty"`
}

func (e *EventsUnified) IsEmpty() bool {
    if e == nil {
        return true
    }
    return e.OnInit == "" && e.OnMount == "" && e.OnSubmit == "" && 
        e.OnChange == "" && e.OnLoad == "" && len(e.Custom) == 0
}

After creating, run "go build ./schema/..." and fix any errors.
'
}

phase_1_6() {
  execute_phase "1.6" "Create interfaces.go" '
Create schema/interfaces.go with public interfaces for the package.

IMPORTANT:
- Check if interfaces.go already exists
- Check if any of these interfaces are already defined elsewhere
- Avoid duplicate interface declarations

Create schema/interfaces.go:

package schema

import "context"

// Renderer renders schema components
type Renderer interface {
    RenderSchema(ctx context.Context, s *Schema) ([]byte, error)
    RenderField(ctx context.Context, f *Field) ([]byte, error)
    RenderAction(ctx context.Context, a *Action) ([]byte, error)
    RenderLayout(ctx context.Context, l *Layout) ([]byte, error)
}

// StateManager manages form state
type StateManager interface {
    Get(field string) any
    Set(field string, value any) error
    GetAll() map[string]any
    SetAll(values map[string]any) error
    Reset() error
    IsDirty(field string) bool
    Touch(field string)
    Errors(field string) []string
    SetErrors(field string, errs []string)
}

// EventDispatcher handles events
type EventDispatcher interface {
    On(event string, handler EventHandler)
    Off(event string)
    Emit(ctx context.Context, event string, data any) error
}

// EventHandler processes events
type EventHandler func(ctx context.Context, data any) error

// BehaviorAdapter converts Behavior to framework-specific output
type BehaviorAdapter interface {
    Adapt(b *Behavior) map[string]string
    Framework() string
}

// BindingAdapter converts Binding to framework-specific output  
type BindingAdapter interface {
    Adapt(b *Binding) map[string]string
    Framework() string
}

After creating, run "go build ./schema/..." and fix any errors.
If there are duplicate interface declarations, remove them from other files.
'
}

phase_1_7() {
  execute_phase "1.7" "Create LayoutBlock type" '
Create schema/layout_block.go with the unified LayoutBlock type.

IMPORTANT:
- This replaces Section, Group, Tab, Step with a single type
- Check for duplicate BlockType declarations
- Use the unified Conditional type (may need to reference it correctly)

Create schema/layout_block.go:

package schema

// LayoutBlock represents a unified layout component
type LayoutBlock struct {
    ID          string       `json:"id"`
    Type        BlockType    `json:"type"`
    Title       string       `json:"title,omitempty"`
    Description string       `json:"description,omitempty"`
    Icon        string       `json:"icon,omitempty"`
    Badge       string       `json:"badge,omitempty"`
    Fields      []string     `json:"fields"`
    Blocks      []LayoutBlock `json:"blocks,omitempty"`
    
    Order       int  `json:"order,omitempty"`
    Disabled    bool `json:"disabled,omitempty"`
    Hidden      bool `json:"hidden,omitempty"`
    Collapsible bool `json:"collapsible,omitempty"`
    Collapsed   bool `json:"collapsed,omitempty"`
    Skippable   bool `json:"skippable,omitempty"`
    Validation  bool `json:"validation,omitempty"`
    Columns     int  `json:"columns,omitempty"`
    Border      bool `json:"border,omitempty"`
    
    Conditional *Conditional `json:"conditional,omitempty"`
    Style       *Style       `json:"style,omitempty"`
}

type BlockType string

const (
    BlockSection BlockType = "section"
    BlockGroup   BlockType = "group"
    BlockTab     BlockType = "tab"
    BlockStep    BlockType = "step"
    BlockCard    BlockType = "card"
    BlockPanel   BlockType = "panel"
)

func (b BlockType) String() string {
    return string(b)
}

func (b BlockType) IsValid() bool {
    switch b {
    case BlockSection, BlockGroup, BlockTab, BlockStep, BlockCard, BlockPanel:
        return true
    default:
        return false
    }
}

// Builder
func NewLayoutBlock(id string, blockType BlockType) *LayoutBlockBuilder {
    return &LayoutBlockBuilder{block: LayoutBlock{ID: id, Type: blockType}}
}

type LayoutBlockBuilder struct {
    block LayoutBlock
}

func (b *LayoutBlockBuilder) Title(t string) *LayoutBlockBuilder {
    b.block.Title = t
    return b
}

func (b *LayoutBlockBuilder) Fields(f ...string) *LayoutBlockBuilder {
    b.block.Fields = f
    return b
}

func (b *LayoutBlockBuilder) Build() LayoutBlock {
    return b.block
}

After creating, run "go build ./schema/..." 

If Conditional or Style types cause errors because they conflict with existing types:
- The unified types should be in conditional_unified.go and style_unified.go
- Check if we need to rename or adjust references

Fix any errors that occur.
'
}

phase_2_update_field() {
  execute_phase "2.1" "Update Field struct" '
Update the Field struct in schema/field.go to use unified types.

STEP BY STEP:

1. First, read the current Field struct in field.go

2. Replace these fields:
   - HTMX *FieldHTMX -> Behavior *Behavior `json:"behavior,omitempty"`
   - Alpine *FieldAlpine -> Binding *Binding `json:"binding,omitempty"`
   - Style *FieldStyle -> Style *Style `json:"style,omitempty"`
   - Conditional *FieldConditional -> Conditional *Conditional `json:"conditional,omitempty"`

3. Add if not present:
   - Events *Events `json:"events,omitempty"` (or *EventsUnified if that is what we created)

4. REMOVE these old type definitions from field.go (if they exist):
   - type FieldHTMX struct {...}
   - type FieldAlpine struct {...}
   - type FieldStyle struct {...}
   - type FieldConditional struct {...}

5. If Style conflicts with existing type, use the one from style_unified.go
   If Conditional conflicts, use the one from conditional_unified.go

After changes, run "go build ./schema/..."

Fix ALL errors:
- If "undefined: Behavior" -> ensure behavior.go exists and is correct
- If "undefined: Binding" -> ensure binding.go exists and is correct
- If "undefined: Style" and we have StyleUnified, update references
- If duplicate declaration, remove the duplicate from the appropriate file
- Update any methods that reference old field types
'
}

phase_2_update_action() {
  execute_phase "2.2" "Update Action struct" '
Update the Action struct in schema/action.go to use unified types.

STEP BY STEP:

1. Read the current Action struct in action.go

2. Replace these fields:
   - HTMX *ActionHTMX -> Behavior *Behavior `json:"behavior,omitempty"`
   - Alpine *ActionAlpine -> Binding *Binding `json:"binding,omitempty"`
   - Theme *ActionTheme -> Style *Style `json:"style,omitempty"`
   - Conditional *ActionConditional -> Conditional *Conditional `json:"conditional,omitempty"`
   - Config *ActionConfig -> Config map[string]any `json:"config,omitempty"`
   - Permissions *ActionPermissions -> Permissions []string `json:"permissions,omitempty"`

3. Add if not present:
   - Events *Events `json:"events,omitempty"`

4. REMOVE these type definitions from action.go:
   - type ActionHTMX struct {...}
   - type ActionAlpine struct {...}
   - type ActionTheme struct {...}
   - type ActionConditional struct {...}
   - type ActionConfig struct {...}
   - type ActionPermissions struct {...}

5. Add helper methods:
   func (a *Action) HasPermission(perm string) bool {
       for _, p := range a.Permissions {
           if p == perm { return true }
       }
       return false
   }

   func (a *Action) GetConfigString(key string) string {
       if a.Config == nil { return "" }
       if v, ok := a.Config[key].(string); ok { return v }
       return ""
   }

6. Update ActionBuilder methods to use new types

After changes, run "go build ./schema/..." and fix ALL errors.
'
}

phase_2_update_schema() {
  execute_phase "2.3" "Update Schema struct" '
Update the Schema struct in schema/schema.go to use unified types.

STEP BY STEP:

1. Read the current Schema struct in schema.go

2. Replace these fields:
   - HTMX *HTMX -> Behavior *Behavior `json:"behavior,omitempty"`
   - Alpine *Alpine -> Binding *Binding `json:"binding,omitempty"`

3. Keep Events field but ensure it uses consistent type

4. DO NOT remove HTMX and Alpine types from types.go yet (we will do that in cleanup phase)

After changes, run "go build ./schema/..." and fix ALL errors.

If there are conflicts between old and new types, the Schema should use the NEW unified types (Behavior, Binding).
'
}

phase_2_update_layout() {
  execute_phase "2.4" "Update Layout struct" '
Update Layout and related types in schema/layout.go to use unified types.

STEP BY STEP:

1. Add Blocks field to Layout struct:
   Blocks []LayoutBlock `json:"blocks,omitempty"`

2. Update Style field in Layout:
   Style *Style `json:"style,omitempty"` (use unified Style type)

3. Update Section struct:
   - Conditional *Conditional (unified type)
   - Style *Style (unified type)

4. Update Group struct:
   - Conditional *Conditional
   - Style *Style  

5. Update Tab struct:
   - Conditional *Conditional
   - Style *Style

6. Update Step struct:
   - Conditional *Conditional
   - Style *Style

7. REMOVE old style types from layout.go (if they exist there):
   - type BaseStyle struct {...}
   - type LayoutStyle struct {...}
   - type SectionStyle struct {...}
   - type TabStyle struct {...}
   - type StepStyle struct {...}
   - type LayoutConditional struct {...}

8. Add conversion methods:
   func (s *Section) ToBlock() LayoutBlock {...}
   func (g *Group) ToBlock() LayoutBlock {...}
   func (t *Tab) ToBlock() LayoutBlock {...}
   func (st *Step) ToBlock() LayoutBlock {...}

After changes, run "go build ./schema/..." and fix ALL errors.
'
}

phase_3_cleanup() {
  execute_phase "3" "Remove deprecated types" '
Remove all deprecated types that have been replaced by unified types.

This is a CLEANUP phase. Remove these types from their respective files:

FROM types.go - REMOVE:
- type HTMX struct {...} (replaced by Behavior)
- type Alpine struct {...} (replaced by Binding)
- type FormSchema struct {...} (unused wrapper)

FROM layout.go - REMOVE (if still present):
- type BaseStyle struct {...}
- type LayoutStyle struct {...}
- type SectionStyle struct {...}
- type TabStyle struct {...}
- type StepStyle struct {...}
- type LayoutConditional struct {...}

FROM field.go - REMOVE (if still present):
- type FieldHTMX struct {...}
- type FieldAlpine struct {...}
- type FieldStyle struct {...}
- type FieldConditional struct {...}

FROM action.go - REMOVE (if still present):
- type ActionHTMX struct {...}
- type ActionAlpine struct {...}
- type ActionTheme struct {...}
- type ActionConditional struct {...}
- type ActionPermissions struct {...}
- type ActionConfig struct {...}

ALSO:
- Remove any methods that ONLY work with deleted types
- Update any remaining references to use unified types
- If removing a type causes errors, update the references first

After each removal, run "go build ./schema/..." and fix errors before continuing.

IMPORTANT: Do not remove types that are still being used. First update all usages, then remove.
'
}

phase_3_rename_unified() {
  execute_phase "3.5" "Rename unified types to final names" '
Now that old types are removed, rename the unified type files to their final names.

1. If we have conditional_unified.go and the old Conditional types are gone:
   - Rename the type from ConditionalUnified to Conditional (if needed)
   - Or keep as Conditional if that is what we named it

2. If we have style_unified.go and old Style types are gone:
   - The type should be called Style
   - Rename file from style_unified.go to style.go (optional, for cleanliness)

3. If we have events_unified.go:
   - Decide: keep old Events in types.go or use EventsUnified
   - If keeping both, ensure they dont conflict
   - If replacing, rename EventsUnified to Events and remove old one

4. Ensure all references throughout the codebase use the correct type names

After changes, run "go build ./schema/..." and fix ALL errors.
'
}

phase_4_update_builders() {
  execute_phase "4" "Update builders" '
Update all builder types to use unified types.

1. FieldBuilder (in field.go):
   - REMOVE methods: WithHTMX, WithAlpine (if they exist)
   - ADD/UPDATE methods:
     func (b *FieldBuilder) WithBehavior(behavior *Behavior) *FieldBuilder {
         b.field.Behavior = behavior
         return b
     }
     func (b *FieldBuilder) WithBinding(binding *Binding) *FieldBuilder {
         b.field.Binding = binding
         return b
     }
     func (b *FieldBuilder) WithStyle(style *Style) *FieldBuilder {
         b.field.Style = style
         return b
     }
     func (b *FieldBuilder) WithConditional(c *Conditional) *FieldBuilder {
         b.field.Conditional = c
         return b
     }
     func (b *FieldBuilder) WithEvents(e *Events) *FieldBuilder {
         b.field.Events = e
         return b
     }

2. ActionBuilder (in action.go):
   - REMOVE methods: WithHTMX, WithAlpine, WithTheme
   - ADD/UPDATE similar methods for Behavior, Binding, Style, Conditional, Events
   - ADD:
     func (b *ActionBuilder) WithPermissions(perms ...string) *ActionBuilder {
         b.action.Permissions = append(b.action.Permissions, perms...)
         return b
     }

3. SchemaBuilder (in builder.go):
   - REMOVE methods: WithHTMX, WithAlpine
   - ADD methods for Behavior, Binding

4. LayoutBuilder (in layout.go):
   - ADD: func (b *LayoutBuilder) AddBlock(block LayoutBlock) *LayoutBuilder
   - UPDATE WithStyle to use *Style

After changes, run "go build ./schema/..." and fix ALL errors.
'
}

phase_5_update_tests() {
  execute_phase "5" "Update tests" '
Update all test files to use unified types.

Search and replace in ALL *_test.go files in schema/:

TYPE REPLACEMENTS:
- FieldHTMX -> Behavior
- FieldAlpine -> Binding
- ActionHTMX -> Behavior
- ActionAlpine -> Binding
- ActionTheme -> Style
- ActionConfig -> use map[string]any for Config field
- ActionConditional -> Conditional
- ActionPermissions -> use []string for Permissions field
- FieldConditional -> Conditional
- LayoutConditional -> Conditional
- LayoutStyle -> Style
- SectionStyle -> Style
- TabStyle -> Style
- StepStyle -> Style
- FieldStyle -> Style

FIELD NAME CHANGES (in struct literals):
- .HTMX -> .Behavior
- .Alpine -> .Binding
- .Theme -> .Style

Also update any test helper functions and mock types.

After changes, run "go test ./schema/..." 

Fix test compilation errors first, then fix failing test assertions.
'
}

phase_6_final() {
  log_info "=== Phase 6: Final verification ==="

  # Search for remaining old type references
  log_info "Searching for remaining old type references..."

  OLD_TYPES="FieldHTMX|FieldAlpine|ActionHTMX|ActionAlpine|ActionTheme|ActionConfig|ActionConditional|ActionPermissions|LayoutStyle|SectionStyle|TabStyle|StepStyle|LayoutConditional|FormSchema"

  remaining=$(grep -rlE "$OLD_TYPES" ./schema/*.go 2>/dev/null | grep -v "_test.go" || true)

  if [ -n "$remaining" ]; then
    log_warn "Found remaining references in: $remaining"

    claude "There are still references to old types in these files:
$remaining

Please find and update ALL remaining references:
- FieldHTMX -> Behavior
- FieldAlpine -> Binding
- ActionHTMX -> Behavior
- ActionAlpine -> Binding
- ActionTheme -> Style
- ActionConfig -> map[string]any
- LayoutStyle/SectionStyle/TabStyle/StepStyle -> Style
- *Conditional variants -> Conditional
- FormSchema -> remove or replace with *Schema

Run go build and fix all errors."

    check_and_fix_build
  fi

  # Final build
  log_info "Final build verification..."
  check_and_fix_build

  # Final tests
  log_info "Final test verification..."
  check_and_fix_tests

  # Commit
  if git rev-parse --git-dir >/dev/null 2>&1; then
    git add -A
    git commit -m "refactor(schema): final cleanup and verification" 2>/dev/null || true
  fi

  log_success "=== REFACTORING COMPLETE ==="
  echo ""
  echo "Summary:"
  echo "  - Created unified types: Behavior, Binding, Style, Conditional, Events, LayoutBlock"
  echo "  - Created interfaces: Renderer, StateManager, EventDispatcher, BehaviorAdapter"
  echo "  - Updated: Schema, Field, Action, Layout to use unified types"
  echo "  - Removed: ~15 redundant type definitions"
  echo "  - Updated: All builders and tests"
  echo ""
  echo "Backup: $BACKUP_DIR"
}

# ============================================================================
# MAIN
# ============================================================================

main() {
  echo "=============================================="
  echo "  Schema Refactoring Script (Self-Healing)"
  echo "=============================================="
  echo ""

  check_prerequisites
  backup_schema

  echo ""
  echo "Select mode:"
  echo "  1) Run ALL phases (recommended)"
  echo "  2) Interactive (confirm each phase)"
  echo "  3) Run specific phase"
  echo "  4) Resume from specific phase"
  echo ""
  read -p "Choice [1-4]: " choice

  case $choice in
  1)
    log_info "Running all phases..."
    phase_1_1
    phase_1_2
    phase_1_3
    phase_1_4
    phase_1_5
    phase_1_6
    phase_1_7
    phase_2_update_field
    phase_2_update_action
    phase_2_update_schema
    phase_2_update_layout
    phase_3_cleanup
    phase_3_rename_unified
    phase_4_update_builders
    phase_5_update_tests
    phase_6_final
    ;;
  2)
    phases=(
      "phase_1_1:Create behavior.go"
      "phase_1_2:Create binding.go"
      "phase_1_3:Create conditional_unified.go"
      "phase_1_4:Create style_unified.go"
      "phase_1_5:Create events_unified.go"
      "phase_1_6:Create interfaces.go"
      "phase_1_7:Create layout_block.go"
      "phase_2_update_field:Update Field struct"
      "phase_2_update_action:Update Action struct"
      "phase_2_update_schema:Update Schema struct"
      "phase_2_update_layout:Update Layout struct"
      "phase_3_cleanup:Remove deprecated types"
      "phase_3_rename_unified:Rename unified types"
      "phase_4_update_builders:Update builders"
      "phase_5_update_tests:Update tests"
      "phase_6_final:Final verification"
    )

    for phase_info in "${phases[@]}"; do
      IFS=':' read -r func desc <<<"$phase_info"
      echo ""
      read -p "Run '$desc'? [Y/n/q]: " confirm
      case $confirm in
      n | N) log_warn "Skipped: $desc" ;;
      q | Q)
        log_info "Quitting"
        exit 0
        ;;
      *) $func ;;
      esac
    done
    ;;
  3 | 4)
    echo ""
    echo "Phases:"
    echo "  1.1-1.7) Create unified types"
    echo "  2.1-2.4) Update core structs"
    echo "  3) Remove deprecated types"
    echo "  3.5) Rename unified types"
    echo "  4) Update builders"
    echo "  5) Update tests"
    echo "  6) Final verification"
    echo ""
    read -p "Start from phase: " phase

    case $phase in
    1.1) phase_1_1 ;&
    1.2) phase_1_2 ;&
    1.3) phase_1_3 ;&
    1.4) phase_1_4 ;&
    1.5) phase_1_5 ;&
    1.6) phase_1_6 ;&
    1.7) phase_1_7 ;&
    2.1) phase_2_update_field ;&
    2.2) phase_2_update_action ;&
    2.3) phase_2_update_schema ;&
    2.4) phase_2_update_layout ;&
    3) phase_3_cleanup ;&
    3.5) phase_3_rename_unified ;&
    4) phase_4_update_builders ;&
    5) phase_5_update_tests ;&
    6) phase_6_final ;;
    *) log_error "Invalid phase" ;;
    esac
    ;;
  *)
    log_error "Invalid choice"
    exit 1
    ;;
  esac
}

# Cleanup on exit
cleanup() {
  rm -f "$BUILD_LOG" "$TEST_LOG" 2>/dev/null
}
trap cleanup EXIT

main "$@"
