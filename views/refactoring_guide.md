# Basecoat CSS Refactoring Guide

## Context & Mission

You are refactoring Go templ components from a utility-class approach to use Basecoat CSS component classes. The codebase has:

- **CSS Framework**: Basecoat (Tailwind plugin with semantic component classes)
- **Template Engine**: Go templ
- **Location**: `views/components/atoms/`, `views/components/molecules/`, `views/components/organisms/`
- **CSS Files**: `static/css/basecoat.css` (component definitions), `static/css/themes/default.css` (theme variables)
- **JS Components**: `static/dist/basecoat/*.js` (interactive behaviors)

**Your Goal**: Transform components to use Basecoat's semantic classes while maintaining functionality and improving code quality.

---

## Phase 1: Discovery & Analysis (30 minutes)

### Task 1.1: Inventory Current Components ✓
**Action**: Scan `views/components/` directories and create a list of all existing components.

**Your Decisions**:
- Which components exist?
- What props do they accept?
- What classes are they currently using?
- Which ones use dynamic class building (TwMerge, string concatenation)?

**Outcome**: A markdown table with columns: `Component Name | Type (Atom/Molecule/Organism) | Current Classes | Has Dynamic Logic | Priority (High/Medium/Low)`

### Task 1.2: Map Basecoat CSS Classes ✓
**Action**: Read `static/css/basecoat.css` completely.

**Your Decisions**:
- What component classes are available? (Hint: Look for `@layer components`)
- What variants exist for each component? (e.g., `.btn-primary`, `.btn-secondary`)
- What size modifiers exist? (e.g., `.btn-sm`, `.btn-lg`)
- Are there icon variants? (e.g., `.btn-icon`)

**CRITICAL DISCOVERY**: Basecoat uses **SINGLE COMBINED CLASSES**, not multiple classes:
- ✅ CORRECT: `class="btn-sm-primary"` (one combined class)
- ❌ WRONG: `class="btn btn-sm btn-primary"` (multiple classes)

**Outcome**: A reference document mapping:
```
Component → Available Classes → Variants → Sizes → Special States
```

### Task 1.3: Identify Gaps ✓
**Action**: Compare your component inventory with available Basecoat classes.

**Your Decisions**:
- Which components have direct Basecoat equivalents?
- Which components need custom CSS extensions?
- Which Basecoat components are unused but might be useful?

**Outcome**: Three lists:
1. **Direct Mappings**: Components that can use Basecoat classes as-is
2. **Need Extensions**: Components requiring custom CSS in `static/css/extensions.css`
3. **Opportunities**: Basecoat components not yet used

---

## Phase 2: Refactoring Strategy (15 minutes)

### Task 2.1: Prioritization ✓
**Your Decisions**:
- Which components are used most frequently?
- Which are simplest to refactor (atoms vs organisms)?
- Which have the most complex dynamic logic?
- What's the optimal refactoring order?

**Outcome**: An ordered list of components to refactor with justification for ordering.

### Task 2.2: Pattern Recognition ✓
**Action**: Study the Basecoat class naming patterns.

**Your Intelligence Challenge**:
- What's the pattern for variants? (class vs data-attribute)
- How are sizes handled?
- How are states represented? (disabled, loading, error, etc.)
- What's the relationship between form components and `.field` wrapper?

**Basecoat Patterns Discovered**:
- **Base + Variant**: `.btn` (primary), `.btn-secondary`, `.btn-outline`, `.btn-ghost`, `.btn-link`, `.btn-destructive`
- **Size + Variant**: `.btn-sm`, `.btn-sm-secondary`, `.btn-lg-destructive`, etc.
- **Icon Variants**: `.btn-icon`, `.btn-sm-icon-outline`, `.btn-lg-icon-primary`, etc.
- **State Handling**: HTML attributes (`disabled`, `aria-pressed="true"`)
- **Features to Keep**: All variants, sizes, disabled state, aria attributes
- **Features to Remove**: `loading` prop (use standard loading patterns), `fullWidth` (use CSS grid/flexbox)

**Outcome**: A "Basecoat Patterns Cheatsheet" documenting:
- Variant pattern: Combined single classes
- Size pattern: Part of combined class name
- State pattern: HTML attributes, not CSS classes
- Composition pattern: Single semantic class per element

---

## Phase 3: Component Refactoring (4-6 hours)

### Task 3.1: Refactor Atoms ✓

**General Refactoring Rules**:
1. **Remove all dynamic class building** (TwMerge, fmt.Sprintf, string concatenation)
2. **Use semantic Basecoat classes** from basecoat.css
3. **Simplify props** - remove className props, keep only semantic props
4. **Follow HTML semantics** - correct element types, ARIA attributes
5. **Preserve functionality** - onclick, form behaviors, validation

**Your Decisions for Each Atom**:

#### Button
- Study `.btn`, `.btn-primary`, `.btn-secondary`, `.btn-outline`, `.btn-ghost`, `.btn-link`, `.btn-destructive`
- Study size variants: `.btn-sm`, `.btn-lg`
- Study icon variants: `.btn-icon`, `.btn-sm-icon`, `.btn-lg-icon`
- Decision: What props structure is cleanest?
- Decision: How to handle icons inside buttons?

#### Input
- Study form input classes in basecoat.css
- Notice: Inputs inside `.field` get special treatment
- Decision: Should Input be standalone or always wrapped in Field?
- Decision: How to represent error/success states?

#### Badge
- Study `.badge`, `.badge-primary`, `.badge-secondary`, `.badge-destructive`, `.badge-outline`
- Decision: Simple string prop for variant or enum?

#### Alert
- Study `.alert`, `.alert-destructive`
- Notice: Grid layout with SVG support
- Decision: How to structure title vs description?

**For Each Atom Component**:

**Step 1**: Read the Basecoat CSS for that component type
**Step 2**: Identify the base class and all variants
**Step 3**: Design the Props struct (minimal, semantic)
**Step 4**: Write the templ component using only Basecoat classes
**Step 5**: Test in isolation

**Required Outcomes**:
- ✓ No dynamic class building
- ✓ Props are simple and semantic
- ✓ Uses correct Basecoat classes
- ✓ Preserves all functionality
- ✓ HTML is semantic and accessible
- ✓ Component renders correctly

### Task 3.2: Refactor Molecules ✓

**Additional Challenges**:
- Molecules compose atoms
- They may need custom layout CSS
- They coordinate multiple interactive elements

**Your Decisions for Each Molecule**:

#### Card
- Study `.card` structure: header, section, footer
- Notice: Header uses grid for title/description/action
- Decision: How to pass header actions?
- Decision: How to handle optional sections?

#### FormField (or Field)
- Study `.field` and `.fieldset`
- Notice: Complex structure with label, input, hint, error
- Notice: data-orientation for horizontal layouts
- Decision: How to compose with Input/Select/Checkbox atoms?
- Decision: How to pass error messages?

#### SearchBar, Dropdown, etc.
- Study relevant Basecoat components
- Decision: Which need custom CSS extensions?
- Decision: How to integrate Basecoat JS components?

**Required Outcomes**:
- ✓ Atoms compose cleanly
- ✓ Layout uses Basecoat patterns or minimal custom CSS
- ✓ Interactive behaviors preserved
- ✓ Form validation works

### Task 3.3: Refactor Organisms ✓

**Complex Challenges**:
- Organisms are complete UI sections
- They may use Basecoat JS for interactions
- They need to integrate with HTMX

**Your Decisions for Each Organism**:

#### Form
- Study `.form` class behavior
- Decision: How to integrate progressive enhancement?
- Decision: How to handle HTMX attributes?

#### Table
- Study `.table` classes
- Notice: Caption, thead, tbody, tfoot styling
- Decision: How to make sortable/filterable?
- Decision: How to handle empty states?

#### Sidebar (if you have one)
- Study `.sidebar` classes in basecoat.css
- Notice: Complex responsive behavior with animations
- Notice: Requires `static/dist/basecoat/sidebar.js`
- Decision: How to initialize JS component?
- Decision: How to handle mobile toggle?

**Required Outcomes**:
- ✓ Complete functional UI sections
- ✓ HTMX integration preserved
- ✓ JS components initialized correctly
- ✓ Responsive behavior works
- ✓ Progressive enhancement maintained

---

## Phase 4: Extensions & Custom Components (2-3 hours)

### Task 4.1: Identify Missing Components ✓

**Your Analysis**:
- Which components do you need that Basecoat doesn't provide?
- Common examples: DataTable, EmptyState, LoadingSkeleton, Stats Cards

**Outcome**: List of custom components needed.

### Task 4.2: Create Custom Extensions ✓

**Location**: `static/css/extensions.css`

**Your Decisions**:
- How to name custom classes? (follow Basecoat conventions)
- How to use Basecoat CSS variables? (--primary, --muted, etc.)
- How to make them responsive and accessible?

**Extension Pattern**:
```css
/* static/css/extensions.css */

@layer components {
  .your-custom-component {
    @apply /* Tailwind utilities */;
    /* Use Basecoat CSS variables */
    color: oklch(var(--foreground));
    background: oklch(var(--background));
    border-color: oklch(var(--border));
    border-radius: var(--radius);
  }
  
  .your-custom-component-variant {
    /* variant styles */
  }
}
```

**Required Outcomes**:
- ✓ Custom components follow Basecoat conventions
- ✓ Use OKLCH color variables
- ✓ Use --radius for border-radius
- ✓ Responsive and accessible
- ✓ Well-documented in code comments

### Task 4.3: Interactive Components with JS ✓

**Components Needing JS**: Dialog, Dropdown Menu, Select (with search), Sidebar, Tabs, Toast, Command

**Your Decisions**:
- Which components need the JS from `static/dist/basecoat/`?
- How to initialize them? (study the Jinja templates in `basecoat/src/jinja/`)
- How to integrate with templ?

**Pattern for JS Components**:
```go
// In your templ component
templ Dropdown(props DropdownProps) {
    <div class="dropdown-menu">
        <button>Trigger</button>
        <div data-popover>
            // Menu items
        </div>
    </div>
    
    // Initialize JS
    <script type="module">
        import { DropdownMenu } from '/static/dist/basecoat/dropdown-menu.js';
        const dropdown = new DropdownMenu(document.querySelector('.dropdown-menu'));
    </script>
}
```

**Required Outcomes**:
- ✓ JS components initialize correctly
- ✓ Interactive behaviors work (open/close, select, etc.)
- ✓ Keyboard navigation works
- ✓ ARIA attributes are correct
- ✓ Works with HTMX (if applicable)

---

## Phase 5: Testing & Validation (2-3 hours)

### Task 5.1: Visual Regression Testing ✓

**Your Testing Checklist**:
```
□ Each component renders correctly in isolation
□ Components compose correctly in molecules/organisms
□ All variants display correctly (primary, secondary, outline, etc.)
□ All sizes display correctly (sm, default, lg)
□ All states work (hover, focus, disabled, loading, error)
□ Dark mode works (if applicable)
□ Responsive breakpoints work
□ Icons display correctly
```

### Task 5.2: Functional Testing ✓

**Your Testing Checklist**:
```
□ Forms submit correctly
□ Validation displays errors
□ Buttons trigger actions
□ Dropdowns open/close
□ Dialogs open/close
□ Selects work (including searchable)
□ Tabs switch content
□ Toasts appear/disappear
□ Sidebar toggles (mobile)
□ HTMX interactions work
□ Progressive enhancement works (JS disabled)
```

### Task 5.3: Accessibility Testing ✓

**Your Testing Checklist**:
```
□ Keyboard navigation works (Tab, Enter, Escape, Arrow keys)
□ Screen reader announcements (test with browser's reader)
□ Focus indicators visible
□ ARIA attributes present and correct
□ Color contrast meets WCAG AA (use browser devtools)
□ Forms have proper labels
□ Buttons have accessible names
□ Interactive elements have proper roles
```

### Task 5.4: Performance Validation ✓

**Your Measurements**:
```
□ CSS bundle size reduced? (compare before/after)
□ No unused CSS classes
□ JS bundles load correctly
□ No console errors
□ Page load time acceptable
□ No layout shifts (CLS)
```

---

## Phase 6: Cleanup & Documentation (1-2 hours)

### Task 6.1: Remove Old Code ✓

**Your Cleanup**:
```
□ Delete TwMerge utility (if exists)
□ Remove old utility CSS files
□ Remove unused Go files (enums, class builders)
□ Remove className props from all components
□ Clean up imports
□ Remove commented-out code
```

### Task 6.2: Update Documentation ✓

**Required Documentation**:

1. **Component Library** (`docs/components.md`):
   - List all components with examples
   - Show all variants and sizes
   - Include copy-paste examples

2. **Styling Guide** (`docs/styling.md`):
   - How to use Basecoat classes
   - How to create custom extensions
   - CSS variables reference
   - Theme customization guide

3. **Migration Notes** (`docs/migration.md`):
   - What changed
   - Breaking changes
   - Before/after examples
   - Update guide for other developers

### Task 6.3: Create Component Templates ✓

**Location**: `templates/components/`

**Your Deliverable**: Create reusable component templates for common patterns:

```
templates/components/
├── atoms/
│   ├── button.example.templ
│   ├── input.example.templ
│   └── badge.example.templ
├── molecules/
│   ├── card.example.templ
│   └── form-field.example.templ
└── organisms/
    ├── form.example.templ
    └── table.example.templ
```

---

## Success Criteria Checklist

### Code Quality ✓
```
□ No dynamic class building (TwMerge, string concat)
□ All components use semantic Basecoat classes
□ Props are minimal and semantic
□ No hardcoded colors (use CSS variables)
□ No inline styles (except data attributes)
□ Code is DRY and maintainable
□ No console errors or warnings
```

### Functionality ✓
```
□ All features work as before
□ Forms validate correctly
□ Interactive components work (dropdowns, dialogs, etc.)
□ HTMX integration preserved
□ Progressive enhancement works
□ No broken layouts
□ Mobile responsive
```

### Accessibility ✓
```
□ WCAG AA contrast ratios
□ Keyboard navigation complete
□ Screen reader compatible
□ Proper ARIA attributes
□ Focus indicators visible
□ Semantic HTML
```

### Performance ✓
```
□ CSS bundle optimized
□ No unused classes
□ Fast page loads
□ No layout shifts
□ JS loads asynchronously
```

### Documentation ✓
```
□ Component library complete
□ Styling guide written
□ Migration notes documented
□ Examples provided
□ Code comments clear
```

---

## Intelligent Decision-Making Framework

As you work through this refactoring, use this framework:

### 1. When You Encounter a Component:

**Ask**:
- Does Basecoat have this? → Check basecoat.css
- What's the semantic HTML? → Use correct element
- What are the variants? → Study the CSS classes
- What props are essential? → Keep only semantic props
- Does it need JS? → Check basecoat/src/jinja/ examples

### 2. When You're Unsure:

**Research**:
- Read the Basecoat CSS definition
- Study the Jinja template examples
- Check similar components for patterns
- Test in browser devtools

**Don't**:
- Guess class names
- Copy utility classes
- Create custom classes prematurely
- Skip accessibility

### 3. When You Find a Gap:

**Evaluate**:
- Can I compose existing components?
- Do I need a custom extension?
- Is there a Basecoat component I'm missing?
- What's the simplest solution?

**Extend Carefully**:
- Follow Basecoat naming conventions
- Use Basecoat CSS variables
- Keep it simple and reusable
- Document your extension

---

## Example Refactoring Walkthrough

Let me show you **one example** so you understand the thinking process. Then you do the rest.

### Original Button (Utility Approach)

```go
type ButtonVariant string
const (
    ButtonVariantPrimary   ButtonVariant = "primary"
    ButtonVariantSecondary ButtonVariant = "secondary"
)

type ButtonProps struct {
    Text      string
    Variant   ButtonVariant
    Size      string
    ClassName string
    OnClick   string
}

templ Button(props ButtonProps) {
    <button 
        class={ utils.TwMerge(
            "btn",
            fmt.Sprintf("btn-%s", string(props.Variant)),
            fmt.Sprintf("btn-%s", props.Size),
            props.ClassName,
        ) }
        onclick={ props.OnClick }
    >
        { props.Text }
    </button>
}
```

### Refactored Button (Basecoat Approach)

**Step 1**: Study basecoat.css
```css
/* ACTUAL basecoat.css pattern - SINGLE COMBINED CLASSES: */
.btn, .btn-primary { /* primary (default) */ }
.btn-secondary, .btn-sm-secondary, .btn-lg-secondary { /* secondary in all sizes */ }
.btn-sm, .btn-sm-primary { /* small primary */ }
.btn-lg, .btn-lg-primary { /* large primary */ }
.btn-icon, .btn-sm-icon, .btn-lg-icon { /* icon-only variants */ }
.btn-sm-icon-destructive { /* small icon destructive */ }
/* Each variant+size combination has its own class! */
```

**Step 2**: Design props
- Remove ClassName (not needed)
- Variant: simple string  
- Size: simple string (optional)
- Type: for form buttons

**Step 3**: Refactor with CORRECT class generation
```go
type ButtonProps struct {
    Text     string
    Variant  string // "primary", "secondary", "outline", "ghost", "destructive"
    Size     string // "sm", "", "lg" (empty = default)
    Icon     templ.Component
    Type     string // "button", "submit", "reset"
    Disabled bool
    OnClick  string
}

// Generate the SINGLE combined Basecoat class
func getButtonClass(variant, size string, iconOnly bool) string {
    // Start with base
    class := "btn"
    
    // Add size prefix if specified
    if size != "" {
        class = "btn-" + size
    }
    
    // Add icon modifier if icon-only
    if iconOnly {
        class = class + "-icon"
    }
    
    // Add variant suffix (except for primary which is default)
    if variant != "" && variant != "primary" {
        class = class + "-" + variant
    }
    
    return class
}

templ Button(props ButtonProps) {
    <button
        class={ getButtonClass(props.Variant, props.Size, props.Icon != nil && props.Text == "") }
        if props.Type != "" {
            type={ props.Type }
        } else {
            type="button"
        }
        if props.Disabled {
            disabled
        }
        if props.OnClick != "" {
            onclick={ props.OnClick }
        }
    >
        if props.Icon != nil {
            @props.Icon
        }
        if props.Text != "" {
            { props.Text }
        }
    </button>
}
```

**Result**:
- ✓ Uses Basecoat classes
- ✓ No dynamic string building
- ✓ Simple, semantic props
- ✓ Handles icons
- ✓ Accessible (type, disabled)

---

## Your Turn

Now you have:
- ✓ The strategy
- ✓ The patterns
- ✓ The success criteria
- ✓ An example

**Start with Task 1.1** and work through systematically. Use your intelligence to make the right decisions. When in doubt, **read the basecoat.css** - it has all the answers.

Good luck! 🚀
