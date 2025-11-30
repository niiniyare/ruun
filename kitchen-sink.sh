#!/bin/bash
# =============================================================================
# kitchen-sink.sh - Pure Go Component Library Demo Generator
# =============================================================================
#
# Generates a kitchen-sink demo for github.com/niiniyare/ruun UI library
#
# PURPOSE:
#   Demonstrate how to USE the ruun package as an external dependency.
#   This is documentation-as-code showing real-world usage patterns.
#
# RULES:
#   ✅ ONLY component function calls: atoms.Button(), molecules.Card(), etc.
#   ✅ Layout primitives as components: atoms.Flex(), atoms.Grid(), atoms.Stack()
#   ✅ All spacing, alignment via component props - NOT Tailwind classes
#   ❌ NO raw HTML tags (<div>, <span>, etc.)
#   ❌ NO Tailwind utility classes in demo code
#
# The ruun package exports:
#   - atoms.*      (Button, Input, Badge, Text, Heading, Flex, Grid, Stack, Box, etc.)
#   - molecules.*  (Card, FormField, Tabs, Pagination, etc.)
#   - organisms.*  (DataTable, Dialog, Form, Sidebar, etc.)
#   - templates.*  (PageLayout, DashboardLayout, AuthLayout, etc.)
#   - pages.*      (LoginPage, DashboardPage, etc.) [optional]
#
# Usage:
#   ./kitchen-sink.sh              # Generate demo
#   ./kitchen-sink.sh --serve      # Generate and run server
#   ./kitchen-sink.sh --init       # Initialize demo project structure
#
# =============================================================================

set -o pipefail

# =============================================================================
# Configuration
# =============================================================================

# The ruun package
RUUN_PKG="github.com/niiniyare/ruun"

# Demo output location (separate from ruun package)
DEMO_DIR="${DEMO_DIR:-./ruun-demo}"
DEMO_SECTIONS_DIR="$DEMO_DIR/sections"

# Server
DEMO_PORT="${DEMO_PORT:-3000}"

# Build
MAX_FIX_ATTEMPTS=5
LOG_FILE="./ruun-demo.log"

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
# Build System
# =============================================================================

run_templ_generate() {
  cd "$DEMO_DIR" && timeout 120 templ generate 2>&1
}

run_go_build() {
  cd "$DEMO_DIR" && timeout 120 go build ./... 2>&1
}

fix_build_errors() {
  local attempt=1
  local build_log="tmp/log/ruun-demo-build-$$.log"

  while [[ $attempt -le $MAX_FIX_ATTEMPTS ]]; do
    log_info "Build attempt $attempt/$MAX_FIX_ATTEMPTS..."

    if run_templ_generate >"$build_log" 2>&1 && run_go_build >>"$build_log" 2>&1; then
      log_success "Build passed!"
      rm -f "$build_log"
      return 0
    fi

    local errors=$(cat "$build_log" | grep -E "(error|Error|undefined|duplicate)" | head -30)

    if [[ -z "$errors" ]]; then
      log_error "Build failed - errors:"
      cat "$build_log"
      rm -f "$build_log"
      return 1
    fi

    log_fix "Attempting fix (attempt $attempt)..."

    claude --dangerously-skip-permissions "Fix these build errors in the ruun demo:

$errors

CRITICAL RULES FOR RUUN PACKAGE DEMO:
1. Import paths: $RUUN_PKG/atoms, $RUUN_PKG/molecules, etc.
2. Use ONLY exported component functions - NO raw HTML
3. Layout must use layout components: atoms.Flex(), atoms.Grid(), atoms.Stack(), atoms.Box()
4. NO Tailwind classes - all styling via component props
5. Fix undefined types by checking what ruun package actually exports

Example correct usage:
  @atoms.Flex(atoms.FlexProps{Direction: atoms.Row, Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
      @atoms.Button(atoms.ButtonProps{Text: \"Save\"})
      @atoms.Button(atoms.ButtonProps{Text: \"Cancel\", Variant: atoms.ButtonOutline})
  }

NOT:
  <div class=\"flex gap-4 items-center\">  // WRONG - raw HTML

Fix all errors."

    ((attempt++))
    sleep 2
  done

  log_error "Could not fix after $MAX_FIX_ATTEMPTS attempts"
  rm -f "$build_log"
  return 1
}

ensure_builds() {
  if ! run_templ_generate 2>/dev/null || ! run_go_build 2>/dev/null; then
    log_warn "Build failed, attempting fixes..."
    fix_build_errors || return 1
  fi
  return 0
}

# =============================================================================
# Initialize Demo Project
# =============================================================================

init_demo_project() {
  log_header "Initializing Ruun Demo Project"

  mkdir -p "$DEMO_DIR"
  mkdir -p "$DEMO_SECTIONS_DIR"

  # Create go.mod
  cat >"$DEMO_DIR/go.mod" <<GOMOD
module ruun-demo

go 1.22

require (
    github.com/a-h/templ v0.2.778
    github.com/gofiber/fiber/v2 v2.52.5
    $RUUN_PKG v0.0.0
)
GOMOD

  log_success "Created $DEMO_DIR/go.mod"

  # Create README
  cat >"$DEMO_DIR/README.md" <<'README'
# Ruun UI Library - Kitchen Sink Demo

This demo showcases how to use the `github.com/niiniyare/ruun` UI component library.

## Running the Demo
```bash
go run main.go
```

Then open http://localhost:3000

## Package Structure
```go
import (
    "github.com/niiniyare/ruun/atoms"      // Basic components
    "github.com/niiniyare/ruun/molecules"  // Combined components
    "github.com/niiniyare/ruun/organisms"  // Complex components
    "github.com/niiniyare/ruun/templates"  // Page layouts
)
```

## Usage Pattern

Everything is a component function call. No raw HTML.
```go
// Layout with Flex
@atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
    @atoms.Button(atoms.ButtonProps{Text: "Save"})
    @atoms.Button(atoms.ButtonProps{Text: "Cancel", Variant: atoms.ButtonOutline})
}

// Grid layout
@atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap6}) {
    @molecules.Card(molecules.CardProps{Title: "Card 1"}) { ... }
    @molecules.Card(molecules.CardProps{Title: "Card 2"}) { ... }
    @molecules.Card(molecules.CardProps{Title: "Card 3"}) { ... }
}
```
README

  log_success "Created README.md"
}

# =============================================================================
# Generate Layout Primitives Documentation
# =============================================================================

generate_layout_primitives_spec() {
  log_task "Generating layout primitives specification..."

  # This documents what the ruun package MUST export for the demo to work
  cat >"$DEMO_DIR/LAYOUT_PRIMITIVES.md" <<'SPEC'
# Ruun Layout Primitives

The ruun package must export these layout primitives in `atoms/` for pure component-based layouts:

## atoms.Box
Container with padding, margin, background.
```go
type BoxProps struct {
    Padding   Spacing    // P0, P1, P2, P4, P6, P8
    Margin    Spacing
    BgColor   Color      // BgBackground, BgMuted, BgCard
    Border    bool
    Rounded   Rounded    // RoundedNone, RoundedSm, RoundedMd, RoundedLg, RoundedFull
    Shadow    Shadow     // ShadowNone, ShadowSm, ShadowMd, ShadowLg
    Attrs     templ.Attributes
}
```

## atoms.Flex
Flexbox container.
```go
type FlexProps struct {
    Direction   FlexDirection  // Row, Col, RowReverse, ColReverse
    Wrap        FlexWrap       // NoWrap, Wrap, WrapReverse
    Justify     Justify        // JustifyStart, JustifyCenter, JustifyEnd, JustifyBetween, JustifyAround
    Align       Align          // AlignStart, AlignCenter, AlignEnd, AlignStretch, AlignBaseline
    Gap         Spacing        // Gap0, Gap1, Gap2, Gap4, Gap6, Gap8
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Grid
CSS Grid container.
```go
type GridProps struct {
    Cols        int            // 1-12
    Rows        int
    Gap         Spacing
    ColGap      Spacing
    RowGap      Spacing
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Stack
Vertical or horizontal stack (simplified Flex).
```go
type StackProps struct {
    Direction   StackDirection  // Vertical, Horizontal
    Gap         Spacing
    Align       Align
    Padding     Spacing
    Attrs       templ.Attributes
}
```

## atoms.Center
Centers content both horizontally and vertically.
```go
type CenterProps struct {
    Padding     Spacing
    MaxWidth    MaxWidth  // MaxWidthSm, MaxWidthMd, MaxWidthLg, MaxWidthXl, MaxWidth2xl
    Attrs       templ.Attributes
}
```

## atoms.Spacer
Flexible space for pushing items apart.
```go
type SpacerProps struct {
    Size    Spacing  // Fixed size, or leave empty for flex-grow
}
```

## atoms.Divider
Horizontal or vertical divider line.
```go
type DividerProps struct {
    Orientation  Orientation  // Horizontal, Vertical
    Label        string       // Optional label in middle
}
```

## Spacing Constants
```go
type Spacing int

const (
    Gap0 Spacing = iota  // 0
    Gap1                  // 0.25rem
    Gap2                  // 0.5rem
    Gap3                  // 0.75rem
    Gap4                  // 1rem
    Gap5                  // 1.25rem
    Gap6                  // 1.5rem
    Gap8                  // 2rem
    Gap10                 // 2.5rem
    Gap12                 // 3rem
)

// Aliases
const (
    P0 = Gap0; P1 = Gap1; P2 = Gap2; P4 = Gap4; P6 = Gap6; P8 = Gap8
    M0 = Gap0; M1 = Gap1; M2 = Gap2; M4 = Gap4; M6 = Gap6; M8 = Gap8
)
```

## Usage in Demo
```templ
// Instead of: <div class="flex gap-4 items-center">
@atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
    @atoms.Button(atoms.ButtonProps{Text: "Button 1"})
    @atoms.Button(atoms.ButtonProps{Text: "Button 2"})
}

// Instead of: <div class="grid grid-cols-3 gap-6">
@atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap6}) {
    @molecules.Card(...) { ... }
    @molecules.Card(...) { ... }
    @molecules.Card(...) { ... }
}

// Instead of: <div class="space-y-4">
@atoms.Stack(atoms.StackProps{Direction: atoms.Vertical, Gap: atoms.Gap4}) {
    @atoms.Text(...)
    @atoms.Text(...)
}
```
SPEC

  log_success "Created LAYOUT_PRIMITIVES.md"
}

# =============================================================================
# Generate Main Server
# =============================================================================

generate_main_server() {
  log_task "Generating main server..."

  claude --dangerously-skip-permissions "Create the demo server for ruun package.

Create: $DEMO_DIR/main.go

This is a standalone demo app that imports and uses the ruun package.

\`\`\`go
package main

import (
    \"context\"
    \"log\"
    \"os\"

    \"github.com/a-h/templ\"
    \"github.com/gofiber/fiber/v2\"
    \"github.com/gofiber/fiber/v2/middleware/logger\"
    \"github.com/gofiber/fiber/v2/middleware/recover\"

    \"ruun-demo/sections\"
)

func main() {
    app := fiber.New(fiber.Config{
        AppName: \"Ruun UI - Kitchen Sink\",
    })

    app.Use(logger.New())
    app.Use(recover.New())

    // Routes
    app.Get(\"/\", render(sections.HomePage()))
    app.Get(\"/atoms\", render(sections.AtomsPage()))
    app.Get(\"/molecules\", render(sections.MoleculesPage()))
    app.Get(\"/organisms\", render(sections.OrganismsPage()))
    app.Get(\"/templates\", render(sections.TemplatesPage()))
    app.Get(\"/examples\", render(sections.ExamplesPage()))

    port := os.Getenv(\"PORT\")
    if port == \"\" {
        port = \"$DEMO_PORT\"
    }

    log.Printf(\"🎨 Ruun Kitchen Sink: http://localhost:%s\", port)
    log.Fatal(app.Listen(\":\" + port))
}

func render(c templ.Component) fiber.Handler {
    return func(ctx *fiber.Ctx) error {
        ctx.Set(\"Content-Type\", \"text/html\")
        return c.Render(context.Background(), ctx.Response().BodyWriter())
    }
}
\`\`\`

Create exactly this file."

  ensure_builds
}

# =============================================================================
# Generate Demo Layout
# =============================================================================

generate_demo_layout() {
  log_task "Generating demo layout..."

  claude --dangerously-skip-permissions "Create the demo layout component using ONLY ruun package functions.

Create: $DEMO_SECTIONS_DIR/layout.templ

CRITICAL RULES:
1. Import from \"$RUUN_PKG/atoms\", \"$RUUN_PKG/molecules\", \"$RUUN_PKG/templates\"
2. Use ONLY component function calls - ZERO raw HTML
3. Use atoms.Flex(), atoms.Grid(), atoms.Stack(), atoms.Box() for ALL layout
4. NO Tailwind classes, NO class strings anywhere
5. All spacing via props: Gap: atoms.Gap4, Padding: atoms.P6, etc.

The ruun package provides:
- atoms.Flex(FlexProps{Direction, Gap, Align, Justify, ...})
- atoms.Stack(StackProps{Direction, Gap, ...})
- atoms.Grid(GridProps{Cols, Gap, ...})
- atoms.Box(BoxProps{Padding, Border, Rounded, ...})
- atoms.Text(TextProps{Text, Size, Weight, Muted, ...})
- atoms.Heading(HeadingProps{Level, Text, ...})
- atoms.Link(LinkProps{Href, Text, Variant, ...})
- atoms.Spacer(SpacerProps{})
- atoms.Divider(DividerProps{})
- templates.BaseHTML(BaseHTMLProps{Title, ...}) for HTML document wrapper

Example structure:
\`\`\`templ
package sections

import (
    \"$RUUN_PKG/atoms\"
    \"$RUUN_PKG/molecules\"
    \"$RUUN_PKG/templates\"
)

// DemoLayout wraps all demo pages
templ DemoLayout(title string, activeSection string) {
    @templates.BaseHTML(templates.BaseHTMLProps{
        Title: title + \" - Ruun Kitchen Sink\",
        // BaseHTML handles <html>, <head>, CSS/JS links, <body>
    }) {
        @atoms.Flex(atoms.FlexProps{Direction: atoms.Row}) {
            // Sidebar
            @DemoSidebar(activeSection)
            
            // Main content
            @atoms.Box(atoms.BoxProps{Padding: atoms.P8, Flex: atoms.Flex1}) {
                { children... }
            }
        }
    }
}

// DemoSidebar - navigation sidebar
templ DemoSidebar(activeSection string) {
    @atoms.Box(atoms.BoxProps{
        Width: atoms.W64,
        Padding: atoms.P4,
        Border: true,
        BorderSide: atoms.BorderRight,
        BgColor: atoms.BgMuted,
        MinHeight: atoms.MinHScreen,
    }) {
        // Logo/Title
        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap2, Padding: atoms.P4}) {
            @atoms.Heading(atoms.HeadingProps{Level: 4, Text: \"Ruun UI\"})
            @atoms.Text(atoms.TextProps{Text: \"Kitchen Sink\", Size: atoms.TextSm, Muted: true})
        }
        
        @atoms.Divider(atoms.DividerProps{})
        
        // Navigation
        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap1, Padding: atoms.P2}) {
            @NavItem(\"/\", \"Overview\", \"overview\", activeSection)
            @NavItem(\"/atoms\", \"Atoms\", \"atoms\", activeSection)
            @NavItem(\"/molecules\", \"Molecules\", \"molecules\", activeSection)
            @NavItem(\"/organisms\", \"Organisms\", \"organisms\", activeSection)
            @NavItem(\"/templates\", \"Templates\", \"templates\", activeSection)
            @NavItem(\"/examples\", \"Examples\", \"examples\", activeSection)
        }
    }
}

// NavItem - single navigation link
templ NavItem(href, label, section, activeSection string) {
    @atoms.Link(atoms.LinkProps{
        Href: href,
        Text: label,
        Variant: getNavVariant(section, activeSection),
        Block: true,
        Padding: atoms.P2,
        Rounded: atoms.RoundedMd,
    })
}

func getNavVariant(section, active string) atoms.LinkVariant {
    if section == active {
        return atoms.LinkActive
    }
    return atoms.LinkNav
}

// Section - demo section wrapper
templ Section(title string) {
    @atoms.Box(atoms.BoxProps{Padding: atoms.P6, Border: true, Rounded: atoms.RoundedLg, Margin: atoms.M4}) {
        @atoms.Heading(atoms.HeadingProps{Level: 3, Text: title})
        @atoms.Box(atoms.BoxProps{Padding: atoms.P4}) {
            { children... }
        }
    }
}
\`\`\`

REMEMBER: 
- ZERO HTML tags
- ZERO class=\"...\" attributes  
- ZERO Tailwind utilities
- Everything is a component function call with typed props"

  ensure_builds
}

# =============================================================================
# Generate Home Page
# =============================================================================

generate_home_page() {
  log_task "Generating home page..."

  claude --dangerously-skip-permissions "Create the home/overview page for ruun kitchen sink.

Create: $DEMO_SECTIONS_DIR/home.templ

CRITICAL: Use ONLY ruun component functions. NO HTML. NO Tailwind classes.

\`\`\`templ
package sections

import (
    \"$RUUN_PKG/atoms\"
    \"$RUUN_PKG/molecules\"
    \"$RUUN_PKG/organisms\"
)

templ HomePage() {
    @DemoLayout(\"Overview\", \"overview\") {
        // Hero
        @atoms.Center(atoms.CenterProps{Padding: atoms.P12}) {
            @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
                @atoms.Heading(atoms.HeadingProps{Level: 1, Text: \"Ruun UI Components\"})
                @atoms.Text(atoms.TextProps{
                    Text: \"A pure Go/Templ component library. No HTML, no CSS classes - just functions.\",
                    Size: atoms.TextLg,
                    Muted: true,
                })
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap2}) {
                    @atoms.Badge(atoms.BadgeProps{Text: \"Go\"})
                    @atoms.Badge(atoms.BadgeProps{Text: \"Templ\", Variant: atoms.BadgeSecondary})
                    @atoms.Badge(atoms.BadgeProps{Text: \"HTMX\", Variant: atoms.BadgeOutline})
                    @atoms.Badge(atoms.BadgeProps{Text: \"Alpine.js\", Variant: atoms.BadgeOutline})
                }
            }
        }
        
        // Stats Grid
        @atoms.Grid(atoms.GridProps{Cols: 4, Gap: atoms.Gap6, Padding: atoms.P6}) {
            @organisms.StatsCard(organisms.StatsCardProps{
                Icon: \"⚛️\",
                Value: \"24\",
                Label: \"Atoms\",
            })
            @organisms.StatsCard(organisms.StatsCardProps{
                Icon: \"🧬\",
                Value: \"20\",
                Label: \"Molecules\",
            })
            @organisms.StatsCard(organisms.StatsCardProps{
                Icon: \"🦠\",
                Value: \"18\",
                Label: \"Organisms\",
            })
            @organisms.StatsCard(organisms.StatsCardProps{
                Icon: \"📐\",
                Value: \"12\",
                Label: \"Templates\",
            })
        }
        
        // Quick Preview
        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap8, Padding: atoms.P6}) {
            @atoms.Heading(atoms.HeadingProps{Level: 2, Text: \"Quick Preview\"})
            
            // Buttons preview
            @Section(\"Buttons\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Wrap: atoms.Wrap}) {
                    @atoms.Button(atoms.ButtonProps{Text: \"Primary\"})
                    @atoms.Button(atoms.ButtonProps{Text: \"Secondary\", Variant: atoms.ButtonSecondary})
                    @atoms.Button(atoms.ButtonProps{Text: \"Outline\", Variant: atoms.ButtonOutline})
                    @atoms.Button(atoms.ButtonProps{Text: \"Destructive\", Variant: atoms.ButtonDestructive})
                    @atoms.Button(atoms.ButtonProps{Text: \"Ghost\", Variant: atoms.ButtonGhost})
                }
            }
            
            // Cards preview
            @Section(\"Cards\") {
                @atoms.Grid(atoms.GridProps{Cols: 2, Gap: atoms.Gap6}) {
                    @molecules.Card(molecules.CardProps{
                        Title: \"Card Title\",
                        Description: \"Card description here\",
                    }) {
                        @atoms.Text(atoms.TextProps{Text: \"Card content\"})
                    }
                    @molecules.Card(molecules.CardProps{Title: \"Another Card\"}) {
                        @atoms.Text(atoms.TextProps{Text: \"More content\"})
                    }
                }
            }
        }
        
        // Usage Example
        @Section(\"How to Use\") {
            @molecules.Card(molecules.CardProps{Title: \"Installation\"}) {
                @atoms.Code(atoms.CodeProps{
                    Block: true,
                    Lang: \"bash\",
                    Text: \"go get github.com/niiniyare/ruun\",
                })
            }
            
            @atoms.Spacer(atoms.SpacerProps{Size: atoms.Gap4})
            
            @molecules.Card(molecules.CardProps{Title: \"Import\"}) {
                @atoms.Code(atoms.CodeProps{
                    Block: true,
                    Lang: \"go\",
                    Text: \"import (\\n    \\\"github.com/niiniyare/ruun/atoms\\\"\\n    \\\"github.com/niiniyare/ruun/molecules\\\"\\n    \\\"github.com/niiniyare/ruun/organisms\\\"\\n)\",
                })
            }
            
            @atoms.Spacer(atoms.SpacerProps{Size: atoms.Gap4})
            
            @molecules.Card(molecules.CardProps{Title: \"Usage\"}) {
                @atoms.Code(atoms.CodeProps{
                    Block: true,
                    Lang: \"go\",
                    Text: \"@atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {\\n    @atoms.Button(atoms.ButtonProps{Text: \\\"Save\\\"})\\n    @atoms.Button(atoms.ButtonProps{\\n        Text: \\\"Cancel\\\",\\n        Variant: atoms.ButtonOutline,\\n    })\\n}\",
                })
            }
        }
    }
}
\`\`\`

NO HTML. NO class attributes. Only component functions."

  ensure_builds
}

# =============================================================================
# Generate Atoms Page
# =============================================================================

generate_atoms_page() {
  log_task "Generating atoms page..."

  claude --dangerously-skip-permissions "Create comprehensive atoms showcase page.

Create: $DEMO_SECTIONS_DIR/atoms.templ

CRITICAL RULES:
1. Import from \"$RUUN_PKG/atoms\" and \"$RUUN_PKG/molecules\"
2. ZERO HTML tags - use atoms.Box, atoms.Flex, atoms.Grid, atoms.Stack for ALL layout
3. ZERO Tailwind/CSS classes - use component props for ALL styling
4. Show EVERY atom with ALL variants, sizes, states

Components to showcase (use atoms.* functions):

LAYOUT PRIMITIVES:
- atoms.Box(BoxProps{Padding, Border, Rounded, Shadow, BgColor, ...})
- atoms.Flex(FlexProps{Direction, Gap, Align, Justify, Wrap, ...})
- atoms.Grid(GridProps{Cols, Gap, ...})
- atoms.Stack(StackProps{Direction, Gap, ...})
- atoms.Center(CenterProps{MaxWidth, ...})
- atoms.Spacer(SpacerProps{Size})

BUTTONS:
- atoms.Button - variants: Default, Secondary, Outline, Destructive, Ghost, Link
- atoms.Button - sizes: Sm, Default, Lg
- atoms.Button - states: Disabled, Loading
- atoms.Button - with icons: IconLeft, IconRight, IconOnly

FORM INPUTS:
- atoms.Input - types: Text, Email, Password, Number, Search, Tel, URL
- atoms.Input - states: Default, Disabled, Readonly, Error
- atoms.Textarea - with rows, resize options
- atoms.Select - with options
- atoms.Checkbox - checked, unchecked, disabled, with description
- atoms.Radio - in groups
- atoms.Switch - on, off, disabled

DISPLAY:
- atoms.Text - sizes: Xs, Sm, Base, Lg, Xl; weights: Normal, Medium, Semibold, Bold; muted
- atoms.Heading - levels: 1, 2, 3, 4, 5, 6
- atoms.Badge - variants: Default, Secondary, Outline, Destructive
- atoms.Avatar - sizes, with image, with fallback initials
- atoms.Icon - sizes, spin animation
- atoms.Code - inline and block
- atoms.Kbd - keyboard shortcuts

FEEDBACK:
- atoms.Alert - variants: Default, Success, Warning, Destructive
- atoms.Progress - values: 0-100
- atoms.Spinner - sizes
- atoms.Skeleton - shapes: text, circle, rectangle

STRUCTURE:
- atoms.Separator - horizontal, vertical
- atoms.Divider - with optional label
- atoms.Link - variants: Default, Muted, Nav, External

Example structure:
\`\`\`templ
package sections

import (
    \"$RUUN_PKG/atoms\"
    \"$RUUN_PKG/molecules\"
)

templ AtomsPage() {
    @DemoLayout(\"Atoms\", \"atoms\") {
        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap8}) {
            @atoms.Heading(atoms.HeadingProps{Level: 1, Text: \"Atoms\"})
            @atoms.Text(atoms.TextProps{
                Text: \"Basic building blocks. Each atom is a single-purpose component.\",
                Muted: true,
            })
            
            // ════════════════════════════════════════════════════════════
            // LAYOUT PRIMITIVES
            // ════════════════════════════════════════════════════════════
            @Section(\"Layout: Flex\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap6}) {
                    @atoms.Text(atoms.TextProps{Text: \"Row with gap\", Weight: atoms.FontMedium})
                    @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
                        @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                            @atoms.Text(atoms.TextProps{Text: \"Item 1\"})
                        }
                        @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                            @atoms.Text(atoms.TextProps{Text: \"Item 2\"})
                        }
                        @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                            @atoms.Text(atoms.TextProps{Text: \"Item 3\"})
                        }
                    }
                    
                    @atoms.Text(atoms.TextProps{Text: \"Justify between\", Weight: atoms.FontMedium})
                    @atoms.Flex(atoms.FlexProps{Justify: atoms.JustifyBetween}) {
                        @atoms.Button(atoms.ButtonProps{Text: \"Left\"})
                        @atoms.Button(atoms.ButtonProps{Text: \"Right\"})
                    }
                    
                    @atoms.Text(atoms.TextProps{Text: \"Align center\", Weight: atoms.FontMedium})
                    @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
                        @atoms.Avatar(atoms.AvatarProps{Initials: \"JD\", Size: atoms.SizeLg})
                        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap1}) {
                            @atoms.Text(atoms.TextProps{Text: \"John Doe\", Weight: atoms.FontMedium})
                            @atoms.Text(atoms.TextProps{Text: \"john@example.com\", Size: atoms.TextSm, Muted: true})
                        }
                    }
                }
            }
            
            @Section(\"Layout: Grid\") {
                @atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap4}) {
                    @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                        @atoms.Text(atoms.TextProps{Text: \"Col 1\"})
                    }
                    @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                        @atoms.Text(atoms.TextProps{Text: \"Col 2\"})
                    }
                    @atoms.Box(atoms.BoxProps{Padding: atoms.P4, BgColor: atoms.BgMuted, Rounded: atoms.RoundedMd}) {
                        @atoms.Text(atoms.TextProps{Text: \"Col 3\"})
                    }
                }
            }
            
            @Section(\"Layout: Stack\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Text(atoms.TextProps{Text: \"Item 1\"})
                    @atoms.Text(atoms.TextProps{Text: \"Item 2\"})
                    @atoms.Text(atoms.TextProps{Text: \"Item 3\"})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // BUTTONS
            // ════════════════════════════════════════════════════════════
            @Section(\"Button Variants\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Wrap: atoms.Wrap}) {
                    @atoms.Button(atoms.ButtonProps{Text: \"Default\"})
                    @atoms.Button(atoms.ButtonProps{Text: \"Secondary\", Variant: atoms.ButtonSecondary})
                    @atoms.Button(atoms.ButtonProps{Text: \"Outline\", Variant: atoms.ButtonOutline})
                    @atoms.Button(atoms.ButtonProps{Text: \"Destructive\", Variant: atoms.ButtonDestructive})
                    @atoms.Button(atoms.ButtonProps{Text: \"Ghost\", Variant: atoms.ButtonGhost})
                    @atoms.Button(atoms.ButtonProps{Text: \"Link\", Variant: atoms.ButtonLink})
                }
            }
            
            @Section(\"Button Sizes\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
                    @atoms.Button(atoms.ButtonProps{Text: \"Small\", Size: atoms.SizeSm})
                    @atoms.Button(atoms.ButtonProps{Text: \"Default\", Size: atoms.SizeDefault})
                    @atoms.Button(atoms.ButtonProps{Text: \"Large\", Size: atoms.SizeLg})
                }
            }
            
            @Section(\"Button States\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
                    @atoms.Button(atoms.ButtonProps{Text: \"Disabled\", Disabled: true})
                    @atoms.Button(atoms.ButtonProps{Text: \"Loading\", Loading: true})
                }
            }
            
            @Section(\"Button with Icons\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
                    @atoms.Button(atoms.ButtonProps{
                        Text: \"Settings\",
                        Icon: atoms.IconCog,
                        IconPosition: atoms.IconLeft,
                    })
                    @atoms.Button(atoms.ButtonProps{
                        Text: \"Next\",
                        Icon: atoms.IconArrowRight,
                        IconPosition: atoms.IconRight,
                    })
                    @atoms.Button(atoms.ButtonProps{
                        Icon: atoms.IconPlus,
                        Variant: atoms.ButtonIcon,
                        AriaLabel: \"Add item\",
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // INPUTS
            // ════════════════════════════════════════════════════════════
            @Section(\"Input Types\") {
                @atoms.Grid(atoms.GridProps{Cols: 2, Gap: atoms.Gap4}) {
                    @atoms.Input(atoms.InputProps{
                        Type: atoms.InputText,
                        Placeholder: \"Text input\",
                    })
                    @atoms.Input(atoms.InputProps{
                        Type: atoms.InputEmail,
                        Placeholder: \"Email\",
                    })
                    @atoms.Input(atoms.InputProps{
                        Type: atoms.InputPassword,
                        Placeholder: \"Password\",
                    })
                    @atoms.Input(atoms.InputProps{
                        Type: atoms.InputNumber,
                        Placeholder: \"Number\",
                    })
                }
            }
            
            @Section(\"Input States\") {
                @atoms.Grid(atoms.GridProps{Cols: 2, Gap: atoms.Gap4}) {
                    @atoms.Input(atoms.InputProps{Placeholder: \"Default\"})
                    @atoms.Input(atoms.InputProps{Placeholder: \"Disabled\", Disabled: true})
                    @atoms.Input(atoms.InputProps{Value: \"Readonly\", Readonly: true})
                    @atoms.Input(atoms.InputProps{Placeholder: \"Error\", Error: true})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // CHECKBOX & RADIO & SWITCH
            // ════════════════════════════════════════════════════════════
            @Section(\"Checkbox\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Checkbox(atoms.CheckboxProps{ID: \"c1\", Label: \"Unchecked\"})
                    @atoms.Checkbox(atoms.CheckboxProps{ID: \"c2\", Label: \"Checked\", Checked: true})
                    @atoms.Checkbox(atoms.CheckboxProps{ID: \"c3\", Label: \"Disabled\", Disabled: true})
                    @atoms.Checkbox(atoms.CheckboxProps{
                        ID: \"c4\",
                        Label: \"With description\",
                        Description: \"This is helpful text\",
                    })
                }
            }
            
            @Section(\"Radio\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Radio(atoms.RadioProps{Name: \"r\", ID: \"r1\", Label: \"Option A\", Value: \"a\"})
                    @atoms.Radio(atoms.RadioProps{Name: \"r\", ID: \"r2\", Label: \"Option B\", Value: \"b\", Checked: true})
                    @atoms.Radio(atoms.RadioProps{Name: \"r\", ID: \"r3\", Label: \"Option C\", Value: \"c\", Disabled: true})
                }
            }
            
            @Section(\"Switch\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Switch(atoms.SwitchProps{ID: \"s1\", Label: \"Off\"})
                    @atoms.Switch(atoms.SwitchProps{ID: \"s2\", Label: \"On\", Checked: true})
                    @atoms.Switch(atoms.SwitchProps{ID: \"s3\", Label: \"Disabled\", Disabled: true})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // SELECT & TEXTAREA
            // ════════════════════════════════════════════════════════════
            @Section(\"Select\") {
                @atoms.Box(atoms.BoxProps{MaxWidth: atoms.MaxWidthMd}) {
                    @atoms.Select(atoms.SelectProps{
                        Placeholder: \"Select an option\",
                        Options: []atoms.Option{
                            {Value: \"1\", Label: \"Option 1\"},
                            {Value: \"2\", Label: \"Option 2\"},
                            {Value: \"3\", Label: \"Option 3\"},
                        },
                    })
                }
            }
            
            @Section(\"Textarea\") {
                @atoms.Box(atoms.BoxProps{MaxWidth: atoms.MaxWidthMd}) {
                    @atoms.Textarea(atoms.TextareaProps{
                        Placeholder: \"Enter your message...\",
                        Rows: 4,
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // BADGES
            // ════════════════════════════════════════════════════════════
            @Section(\"Badge Variants\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
                    @atoms.Badge(atoms.BadgeProps{Text: \"Default\"})
                    @atoms.Badge(atoms.BadgeProps{Text: \"Secondary\", Variant: atoms.BadgeSecondary})
                    @atoms.Badge(atoms.BadgeProps{Text: \"Outline\", Variant: atoms.BadgeOutline})
                    @atoms.Badge(atoms.BadgeProps{Text: \"Destructive\", Variant: atoms.BadgeDestructive})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // AVATAR
            // ════════════════════════════════════════════════════════════
            @Section(\"Avatar\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
                    @atoms.Avatar(atoms.AvatarProps{Initials: \"SM\", Size: atoms.SizeSm})
                    @atoms.Avatar(atoms.AvatarProps{Initials: \"MD\", Size: atoms.SizeDefault})
                    @atoms.Avatar(atoms.AvatarProps{Initials: \"LG\", Size: atoms.SizeLg})
                    @atoms.Avatar(atoms.AvatarProps{
                        Src: \"https://github.com/shadcn.png\",
                        Alt: \"User\",
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // ALERTS
            // ════════════════════════════════════════════════════════════
            @Section(\"Alert\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Alert(atoms.AlertProps{
                        Title: \"Default\",
                        Description: \"This is a default alert.\",
                    })
                    @atoms.Alert(atoms.AlertProps{
                        Title: \"Success\",
                        Description: \"Operation completed.\",
                        Variant: atoms.AlertSuccess,
                    })
                    @atoms.Alert(atoms.AlertProps{
                        Title: \"Warning\",
                        Description: \"Please review.\",
                        Variant: atoms.AlertWarning,
                    })
                    @atoms.Alert(atoms.AlertProps{
                        Title: \"Error\",
                        Description: \"Something went wrong.\",
                        Variant: atoms.AlertDestructive,
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // PROGRESS & SPINNER
            // ════════════════════════════════════════════════════════════
            @Section(\"Progress\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Progress(atoms.ProgressProps{Value: 25})
                    @atoms.Progress(atoms.ProgressProps{Value: 50})
                    @atoms.Progress(atoms.ProgressProps{Value: 75})
                    @atoms.Progress(atoms.ProgressProps{Value: 100})
                }
            }
            
            @Section(\"Spinner\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap8, Align: atoms.AlignCenter}) {
                    @atoms.Spinner(atoms.SpinnerProps{Size: atoms.SizeSm})
                    @atoms.Spinner(atoms.SpinnerProps{Size: atoms.SizeDefault})
                    @atoms.Spinner(atoms.SpinnerProps{Size: atoms.SizeLg})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // SKELETON
            // ════════════════════════════════════════════════════════════
            @Section(\"Skeleton\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Align: atoms.AlignCenter}) {
                    @atoms.Skeleton(atoms.SkeletonProps{Shape: atoms.SkeletonCircle, Size: atoms.SizeLg})
                    @atoms.Stack(atoms.StackProps{Gap: atoms.Gap2}) {
                        @atoms.Skeleton(atoms.SkeletonProps{Width: atoms.W64, Height: atoms.H4})
                        @atoms.Skeleton(atoms.SkeletonProps{Width: atoms.W48, Height: atoms.H4})
                    }
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // TYPOGRAPHY
            // ════════════════════════════════════════════════════════════
            @Section(\"Headings\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap2}) {
                    @atoms.Heading(atoms.HeadingProps{Level: 1, Text: \"Heading 1\"})
                    @atoms.Heading(atoms.HeadingProps{Level: 2, Text: \"Heading 2\"})
                    @atoms.Heading(atoms.HeadingProps{Level: 3, Text: \"Heading 3\"})
                    @atoms.Heading(atoms.HeadingProps{Level: 4, Text: \"Heading 4\"})
                }
            }
            
            @Section(\"Text\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap2}) {
                    @atoms.Text(atoms.TextProps{Text: \"Regular text\"})
                    @atoms.Text(atoms.TextProps{Text: \"Muted text\", Muted: true})
                    @atoms.Text(atoms.TextProps{Text: \"Small text\", Size: atoms.TextSm})
                    @atoms.Text(atoms.TextProps{Text: \"Large text\", Size: atoms.TextLg})
                    @atoms.Text(atoms.TextProps{Text: \"Bold text\", Weight: atoms.FontBold})
                }
            }
            
            @Section(\"Code & Kbd\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap2, Align: atoms.AlignCenter}) {
                        @atoms.Text(atoms.TextProps{Text: \"Inline code:\"})
                        @atoms.Code(atoms.CodeProps{Text: \"const x = 1\"})
                    }
                    @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap2, Align: atoms.AlignCenter}) {
                        @atoms.Text(atoms.TextProps{Text: \"Keyboard shortcut:\"})
                        @atoms.Kbd(atoms.KbdProps{Keys: []string{\"⌘\", \"K\"}})
                    }
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // SEPARATOR & DIVIDER
            // ════════════════════════════════════════════════════════════
            @Section(\"Separator\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Text(atoms.TextProps{Text: \"Above\"})
                    @atoms.Separator(atoms.SeparatorProps{})
                    @atoms.Text(atoms.TextProps{Text: \"Below\"})
                }
            }
            
            @Section(\"Divider with Label\") {
                @atoms.Divider(atoms.DividerProps{Label: \"OR\"})
            }
            
            // ════════════════════════════════════════════════════════════
            // LINK
            // ════════════════════════════════════════════════════════════
            @Section(\"Link\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap2}) {
                    @atoms.Link(atoms.LinkProps{Href: \"#\", Text: \"Default link\"})
                    @atoms.Link(atoms.LinkProps{Href: \"#\", Text: \"Muted link\", Variant: atoms.LinkMuted})
                    @atoms.Link(atoms.LinkProps{Href: \"https://example.com\", Text: \"External\", External: true})
                }
            }
        }
    }
}
\`\`\`

CRITICAL: ZERO HTML. ZERO class attributes. Every layout uses atoms.Flex/Grid/Stack/Box."

  ensure_builds
}

# =============================================================================
# Generate Molecules Page
# =============================================================================

generate_molecules_page() {
  log_task "Generating molecules page..."

  claude --dangerously-skip-permissions "Create comprehensive molecules showcase page.

Create: $DEMO_SECTIONS_DIR/molecules.templ

CRITICAL: Use ONLY component functions. NO HTML. NO Tailwind.

Molecules to showcase (adapt to actual ruun package exports):
- molecules.Card - with title, description, footer
- molecules.FormField - label + input + error + description
- molecules.Tabs - tabbed content
- molecules.Pagination - page navigation
- molecules.StatusBadge - status indicators
- molecules.EmptyState - no data state
- molecules.SearchBox - search input
- molecules.InputGroup - input with prefix/suffix
- molecules.FilterChip - filter pills
- molecules.StatDisplay - stat with trend
- molecules.DropdownMenu - dropdown with items
- molecules.Tooltip - hover tooltip

Example:
\`\`\`templ
package sections

import (
    \"$RUUN_PKG/atoms\"
    \"$RUUN_PKG/molecules\"
)

templ MoleculesPage() {
    @DemoLayout(\"Molecules\", \"molecules\") {
        @atoms.Stack(atoms.StackProps{Gap: atoms.Gap8}) {
            @atoms.Heading(atoms.HeadingProps{Level: 1, Text: \"Molecules\"})
            @atoms.Text(atoms.TextProps{
                Text: \"Combinations of atoms working together as a unit.\",
                Muted: true,
            })
            
            // ════════════════════════════════════════════════════════════
            // CARD
            // ════════════════════════════════════════════════════════════
            @Section(\"Card\") {
                @atoms.Grid(atoms.GridProps{Cols: 2, Gap: atoms.Gap6}) {
                    @molecules.Card(molecules.CardProps{
                        Title: \"Card Title\",
                        Description: \"Card description goes here\",
                    }) {
                        @atoms.Text(atoms.TextProps{Text: \"Card content area\"})
                    }
                    
                    @molecules.Card(molecules.CardProps{
                        Title: \"Card with Footer\",
                    }) {
                        @atoms.Text(atoms.TextProps{Text: \"Content\"})
                        @molecules.CardFooter() {
                            @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap2, Justify: atoms.JustifyEnd}) {
                                @atoms.Button(atoms.ButtonProps{Text: \"Cancel\", Variant: atoms.ButtonOutline})
                                @atoms.Button(atoms.ButtonProps{Text: \"Save\"})
                            }
                        }
                    }
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // FORM FIELD
            // ════════════════════════════════════════════════════════════
            @Section(\"Form Field\") {
                @atoms.Grid(atoms.GridProps{Cols: 2, Gap: atoms.Gap6}) {
                    @molecules.FormField(molecules.FormFieldProps{
                        Label: \"Email\",
                        Name: \"email\",
                        Type: atoms.InputEmail,
                        Placeholder: \"name@example.com\",
                    })
                    @molecules.FormField(molecules.FormFieldProps{
                        Label: \"Password\",
                        Name: \"password\",
                        Type: atoms.InputPassword,
                    })
                    @molecules.FormField(molecules.FormFieldProps{
                        Label: \"With Description\",
                        Name: \"desc\",
                        Description: \"This is helpful text\",
                    })
                    @molecules.FormField(molecules.FormFieldProps{
                        Label: \"With Error\",
                        Name: \"err\",
                        Error: \"This field is required\",
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // TABS
            // ════════════════════════════════════════════════════════════
            @Section(\"Tabs\") {
                @molecules.Tabs(molecules.TabsProps{
                    DefaultTab: \"account\",
                    Tabs: []molecules.Tab{
                        {ID: \"account\", Label: \"Account\"},
                        {ID: \"password\", Label: \"Password\"},
                        {ID: \"settings\", Label: \"Settings\"},
                    },
                }) {
                    @molecules.TabPanel(\"account\") {
                        @atoms.Text(atoms.TextProps{Text: \"Account settings\"})
                    }
                    @molecules.TabPanel(\"password\") {
                        @atoms.Text(atoms.TextProps{Text: \"Password settings\"})
                    }
                    @molecules.TabPanel(\"settings\") {
                        @atoms.Text(atoms.TextProps{Text: \"General settings\"})
                    }
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // PAGINATION
            // ════════════════════════════════════════════════════════════
            @Section(\"Pagination\") {
                @molecules.Pagination(molecules.PaginationProps{
                    CurrentPage: 3,
                    TotalPages: 10,
                    TotalItems: 100,
                    ItemsPerPage: 10,
                })
            }
            
            // ════════════════════════════════════════════════════════════
            // STATUS BADGE
            // ════════════════════════════════════════════════════════════
            @Section(\"Status Badge\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4, Wrap: atoms.Wrap}) {
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusActive})
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusInactive})
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusPending})
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusApproved})
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusRejected})
                    @molecules.StatusBadge(molecules.StatusBadgeProps{Status: molecules.StatusDraft})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // EMPTY STATE
            // ════════════════════════════════════════════════════════════
            @Section(\"Empty State\") {
                @molecules.EmptyState(molecules.EmptyStateProps{
                    Icon: \"📭\",
                    Title: \"No results found\",
                    Description: \"Try adjusting your search or filters.\",
                    Action: atoms.ButtonProps{Text: \"Clear Filters\"},
                })
            }
            
            // ════════════════════════════════════════════════════════════
            // SEARCH BOX
            // ════════════════════════════════════════════════════════════
            @Section(\"Search Box\") {
                @atoms.Box(atoms.BoxProps{MaxWidth: atoms.MaxWidthMd}) {
                    @molecules.SearchBox(molecules.SearchBoxProps{
                        Placeholder: \"Search...\",
                        Name: \"search\",
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // INPUT GROUP
            // ════════════════════════════════════════════════════════════
            @Section(\"Input Group\") {
                @atoms.Stack(atoms.StackProps{Gap: atoms.Gap4}) {
                    @atoms.Box(atoms.BoxProps{MaxWidth: atoms.MaxWidthMd}) {
                        @molecules.InputGroup(molecules.InputGroupProps{
                            Prefix: \"https://\",
                            Placeholder: \"example.com\",
                        })
                    }
                    @atoms.Box(atoms.BoxProps{MaxWidth: atoms.MaxWidthMd}) {
                        @molecules.InputGroup(molecules.InputGroupProps{
                            Prefix: \"$\",
                            Suffix: \".00\",
                            Placeholder: \"0\",
                            Type: atoms.InputNumber,
                        })
                    }
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // FILTER CHIP
            // ════════════════════════════════════════════════════════════
            @Section(\"Filter Chip\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap2, Wrap: atoms.Wrap}) {
                    @molecules.FilterChip(molecules.FilterChipProps{Text: \"All\", Active: true})
                    @molecules.FilterChip(molecules.FilterChipProps{Text: \"Active\"})
                    @molecules.FilterChip(molecules.FilterChipProps{Text: \"Pending\"})
                    @molecules.FilterChip(molecules.FilterChipProps{Text: \"Status: Active\", Removable: true})
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // STAT DISPLAY
            // ════════════════════════════════════════════════════════════
            @Section(\"Stat Display\") {
                @atoms.Grid(atoms.GridProps{Cols: 3, Gap: atoms.Gap6}) {
                    @molecules.StatDisplay(molecules.StatDisplayProps{
                        Label: \"Revenue\",
                        Value: \"$45,231\",
                        Trend: \"+12.5%\",
                        TrendUp: true,
                    })
                    @molecules.StatDisplay(molecules.StatDisplayProps{
                        Label: \"Orders\",
                        Value: \"2,345\",
                        Trend: \"+8.2%\",
                        TrendUp: true,
                    })
                    @molecules.StatDisplay(molecules.StatDisplayProps{
                        Label: \"Returns\",
                        Value: \"23\",
                        Trend: \"-2.4%\",
                        TrendUp: false,
                    })
                }
            }
            
            // ════════════════════════════════════════════════════════════
            // TOOLTIP
            // ════════════════════════════════════════════════════════════
            @Section(\"Tooltip\") {
                @atoms.Flex(atoms.FlexProps{Gap: atoms.Gap4}) {
                    @molecules.Tooltip(molecules.TooltipProps{Content: \"Top tooltip\", Side: molecules.TooltipTop}) {
                        @atoms.Button(atoms.ButtonProps{Text: \"Hover (Top)\", Variant: atoms.ButtonOutline})
                    }
                    @molecules.Tooltip(molecules.TooltipProps{Content: \"Right tooltip\", Side: molecules.TooltipRight}) {
                        @atoms.Button(atoms.ButtonProps{Text: \"Hover (Right)\", Variant: atoms.ButtonOutline})
                    }
                }
            }
        }
    }
}
\`\`\`

CRITICAL: ZERO HTML tags. ZERO class strings. Pure component functions."

  ensure_builds
}

# =============================================================================
# Generate Organisms Page
# =============================================================================

generate_organisms_page() {
  log_task "Generating organisms page..."

  claude --dangerously-skip-permissions "Create comprehensive organisms showcase page.

Create: $DEMO_SECTIONS_DIR/organisms.templ

CRITICAL: Use ONLY component functions. NO HTML. NO Tailwind.

Organisms to showcase:
- organisms.DataTable - full data table with columns, sorting, pagination
- organisms.Dialog - modal dialog
- organisms.Form - complete form organism
- organisms.StatsCard - dashboard stat card
- organisms.Breadcrumbs - navigation breadcrumbs
- organisms.Wizard - multi-step wizard
- organisms.Sidebar - navigation sidebar
- organisms.FilterPanel - filter controls panel

Use atoms.Flex/Grid/Stack/Box for ALL layout. NO HTML divs.

Include Alpine.js bindings via Attrs prop where needed for interactivity."

  ensure_builds
}

# =============================================================================
# Generate Templates Page
# =============================================================================

generate_templates_page() {
  log_task "Generating templates page..."

  claude --dangerously-skip-permissions "Create templates showcase page.

Create: $DEMO_SECTIONS_DIR/templates.templ

Show available page layouts:
- templates.PageLayout - base page layout
- templates.DashboardLayout - sidebar + main
- templates.AuthLayout - centered card layout
- templates.FormLayout - centered form
- templates.ErrorLayout - error pages

Use ONLY component functions for the showcase grid/cards."

  ensure_builds
}

# =============================================================================
# Generate Examples Page
# =============================================================================

generate_examples_page() {
  log_task "Generating examples page..."

  claude --dangerously-skip-permissions "Create real-world examples page.

Create: $DEMO_SECTIONS_DIR/examples.templ

Show complete, realistic examples:

1. Invoice List - DataTable with filters, search, pagination, actions
2. Login Form - AuthLayout with FormField components
3. Dashboard - StatsCards + DataTable
4. Settings Form - Tabs + FormFields
5. Product Grid - Grid of Cards

CRITICAL: ZERO HTML. Use only:
- atoms.* for primitives and layout (Flex, Grid, Stack, Box)
- molecules.* for combined components
- organisms.* for complex components
- templates.* for page layouts

Show code snippets using atoms.Code with Block: true."

  ensure_builds
}

# =============================================================================
# Run All
# =============================================================================

run_all() {
  log_header "Generating Ruun Kitchen Sink Demo"

  init_demo_project
  generate_layout_primitives_spec
  generate_main_server
  generate_demo_layout
  generate_home_page
  generate_atoms_page
  generate_molecules_page
  generate_organisms_page
  generate_templates_page
  generate_examples_page

  # Final build check
  log_info "Final build verification..."
  cd "$DEMO_DIR"
  if run_templ_generate && run_go_build; then
    log_success "Build successful!"
  else
    fix_build_errors
  fi

  log_header "Generation Complete!"
  log_success "Demo generated at: $DEMO_DIR"
  log_info "Start server: cd $DEMO_DIR && go run main.go"
  log_info "Or: $0 --serve"
}

start_server() {
  log_header "Starting Ruun Demo Server"

  cd "$DEMO_DIR"

  log_info "Generating templ..."
  templ generate

  log_info "Starting server at http://localhost:$DEMO_PORT"
  go run main.go
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
  log_success "Prerequisites OK"
}

show_help() {
  cat <<'EOF'
Ruun Kitchen Sink Generator

Generates a pure Go/Templ demo for the github.com/niiniyare/ruun UI library.

Usage:
  ./kitchen-sink.sh              # Generate demo
  ./kitchen-sink.sh --serve      # Generate and run server
  ./kitchen-sink.sh --init       # Initialize project only

Key Principles:
  ✅ ONLY component function calls
  ✅ Layout via atoms.Flex/Grid/Stack/Box
  ✅ All styling via typed props
  ❌ NO raw HTML tags
  ❌ NO Tailwind/CSS class strings

Package Structure:
  github.com/niiniyare/ruun/
  ├── atoms/       # Button, Input, Flex, Grid, Stack, Box, Text, etc.
  ├── molecules/   # Card, FormField, Tabs, Pagination, etc.
  ├── organisms/   # DataTable, Dialog, Form, Sidebar, etc.
  └── templates/   # PageLayout, DashboardLayout, etc.

Output:
  ruun-demo/
  ├── main.go
  ├── go.mod
  ├── README.md
  ├── LAYOUT_PRIMITIVES.md
  └── sections/
      ├── layout.templ
      ├── home.templ
      ├── atoms.templ
      ├── molecules.templ
      ├── organisms.templ
      ├── templates.templ
      └── examples.templ
EOF
}

main() {
  case "${1:-}" in
  --help | -h)
    show_help
    exit 0
    ;;
  --serve)
    check_prerequisites
    run_all
    start_server
    ;;
  --init)
    check_prerequisites
    init_demo_project
    generate_layout_primitives_spec
    log_success "Project initialized at $DEMO_DIR"
    ;;
  *)
    check_prerequisites
    run_all
    ;;
  esac
}

main "$@"
