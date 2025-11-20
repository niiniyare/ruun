# Migration Guide: Old Component System → New Atomic Design Architecture

This comprehensive guide will help you migrate from the old component system to the new atomic design architecture safely and incrementally.

## Table of Contents

1. [Migration Overview](#migration-overview)
2. [Pre-Migration Checklist](#pre-migration-checklist)
3. [Migration Tools](#migration-tools)
4. [Step-by-Step Migration Process](#step-by-step-migration-process)
5. [Component Migration Patterns](#component-migration-patterns)
6. [Validation and Testing](#validation-and-testing)
7. [Rollback Procedures](#rollback-procedures)
8. [Troubleshooting](#troubleshooting)
9. [Best Practices](#best-practices)

## Migration Overview

The migration process transforms your codebase from the old component system to a new atomic design architecture that provides:

- **Pure Presentation Components**: Components focused solely on UI rendering
- **Props-based Configuration**: Structured props interfaces instead of parameter lists
- **Compiled Theme Classes**: CSS classes generated from JSON theme definitions
- **Better Type Safety**: Strongly typed props and component interfaces
- **Enhanced Composition**: Atomic components that compose into molecules and organisms

### Key Changes

| Old System | New System |
|------------|------------|
| `@Button("Save", "primary")` | `@atoms.Button(atoms.ButtonProps{Text: "Save", Variant: atoms.ButtonPrimary})` |
| `@FormField("name", "text", "Enter name")` | `@molecules.FormField(molecules.FormFieldProps{Name: "name", Type: "text", Placeholder: "Enter name"})` |
| Direct HTMX attributes | HTMX via props: `HXPost: "/api/save"` |
| `theme.Get("button.primary")` | Compiled CSS class: `"button-primary"` |
| Hardcoded colors: `"text-red-500"` | Theme tokens: `"text-error"` |

## Pre-Migration Checklist

Before starting the migration, ensure you have:

- [ ] **Backup Strategy**: Full project backup or version control
- [ ] **Testing Environment**: Dedicated environment for testing changes
- [ ] **Time Allocation**: Plan for 2-4 hours per medium-sized project
- [ ] **Team Coordination**: Inform team members about migration timeline
- [ ] **Documentation Access**: This guide and component documentation
- [ ] **Tool Preparation**: Migration tools built and tested

### System Requirements

- Go 1.21 or later
- Templ CLI installed
- Project uses the ruun framework
- Access to terminal/command line

## Migration Tools

The migration process uses four specialized tools:

### 1. Migration Analyzer (`analyzer.go`)
Scans your codebase to identify migration requirements.

```bash
go run analyzer.go /path/to/project migration-report.json
```

### 2. Automated Migrator (`migrator.go`)
Applies automated transformations to your code.

```bash
# Dry run (preview changes)
go run migrator.go /path/to/project --dry-run

# Apply migrations
go run migrator.go /path/to/project
```

### 3. Migration Validator (`validator.go`)
Validates that migrations were applied correctly.

```bash
go run validator.go /path/to/project
```

### 4. Rollback Tool (`rollback.go`)
Safely rolls back migrations if needed.

```bash
# List available backups
go run rollback.go list /path/to/project

# Full rollback
go run rollback.go full /path/to/project /backup/path
```

## Step-by-Step Migration Process

### Phase 1: Analysis and Planning (15-30 minutes)

#### Step 1.1: Run Migration Analysis

```bash
cd tools/migration
go run analyzer.go ../../ analysis-report.json
```

**Expected Output:**
```
MIGRATION ANALYSIS SUMMARY
==========================================================
Project Path: /path/to/project
Files Scanned: 45
Total Issues Found: 23
Auto-fixable Issues: 18

Issues by Severity:
  - Errors: 5
  - Warnings: 15
  - Info: 3

Issues by Type:
  - Import Issues: 8
  - Component Issues: 12
  - Class Issues: 2
  - Theme Issues: 1
```

#### Step 1.2: Review Analysis Report

Open `analysis-report.json` to understand:
- Which components need migration
- Estimated effort per component
- Manual vs. automated migration requirements

#### Step 1.3: Plan Migration Order

Prioritize migrations in this order:
1. **Atoms first**: Button, Input, Icon components
2. **Molecules next**: FormField, SearchBox components  
3. **Organisms last**: Form, Modal, Navigation components
4. **Pages finally**: Update page-level imports

### Phase 2: Backup and Preparation (5-10 minutes)

#### Step 2.1: Create Project Backup

```bash
# Option 1: Git commit (recommended)
git add .
git commit -m "Pre-migration backup"

# Option 2: Manual backup
cp -r /path/to/project /path/to/project-backup
```

#### Step 2.2: Prepare Migration Environment

```bash
# Build migration tools
cd tools/migration
go build -o analyzer analyzer.go
go build -o migrator migrator.go
go build -o validator validator.go
go build -o rollback rollback.go
```

### Phase 3: Automated Migration (30-60 minutes)

#### Step 3.1: Dry Run Migration

```bash
./migrator ../../ --dry-run --output=dry-run-session.json
```

**Review dry run results:**
- Check which files will be modified
- Verify transformation patterns look correct
- Note any unexpected changes

#### Step 3.2: Execute Migration

```bash
./migrator ../../ --output=migration-session.json
```

**Monitor output for:**
- Files being processed
- Applied transformation rules
- Any errors or warnings

#### Step 3.3: Validate Migration

```bash
./validator ../../ --session=migration-session.json --output=validation-report.json
```

**Check validation results:**
- Syntax errors in migrated files
- Missing imports
- Component usage issues

### Phase 4: Manual Migration (15-45 minutes)

#### Step 4.1: Address Validation Issues

Fix any issues found during validation:

```bash
# Common fixes needed:
# 1. Add missing imports
import "github.com/niiniyare/ruun/views/components/atoms"

# 2. Fix component prop structures
@atoms.Button(atoms.ButtonProps{
    Text: "Save",
    Variant: atoms.ButtonPrimary,
    HXPost: "/api/save",
})

# 3. Update complex component usage
@molecules.FormField(molecules.FormFieldProps{
    ID: "email",
    Name: "email", 
    Type: molecules.FormFieldEmail,
    Label: "Email Address",
    Required: true,
    Validation: validation.Rules{
        validation.Required(),
        validation.Email(),
    },
})
```

#### Step 4.2: Update Complex Components

Some components require manual migration:

**Enhanced Forms:**
```go
// Old
@Form(formData, "/api/submit")

// New  
@organisms.EnhancedForm(organisms.FormProps{
    Schema: &schema.Form{
        Fields: []schema.Field{
            {Name: "email", Type: "email", Required: true},
            {Name: "password", Type: "password", Required: true},
        },
    },
    Action: "/api/submit",
    Method: "POST",
})
```

**Navigation Components:**
```go
// Old
@Navigation(menuItems, user)

// New
@organisms.EnhancedNavigation(organisms.NavigationProps{
    Schema: &schema.Navigation{
        Brand: "MyApp",
        Items: navItems,
    },
    User: user,
    Theme: currentTheme,
})
```

#### Step 4.3: Update Theme Usage

Replace theme.Get() calls with compiled classes:

```go
// Old
class={ theme.Get("button.primary.background") }

// New  
class="button-primary"

// Old
style={ fmt.Sprintf("color: %s", theme.Get("text.error")) }

// New
class="text-error"
```

### Phase 5: Testing and Validation (30-60 minutes)

#### Step 5.1: Run Comprehensive Validation

```bash
./validator ../../ --output=final-validation.json
```

#### Step 5.2: Compile and Test

```bash
# Generate templ files
templ generate

# Build application
go build

# Run tests
go test ./...
```

#### Step 5.3: Visual Testing

1. Start your application in development mode
2. Navigate through all pages/components
3. Test interactive features (forms, buttons, modals)
4. Verify styling and themes work correctly
5. Test responsive design

#### Step 5.4: Performance Testing

1. Check bundle sizes haven't increased significantly
2. Verify page load times
3. Test component rendering performance

### Phase 6: Documentation and Cleanup (15-30 minutes)

#### Step 6.1: Update Documentation

- Update component usage in README files
- Update code examples in documentation
- Update API documentation if applicable

#### Step 6.2: Clean Up Migration Files

```bash
# Archive migration reports
mkdir migration-archive-$(date +%Y%m%d)
mv *.json migration-archive-*/

# Keep tools for future use
# Consider committing migration tools to version control
```

#### Step 6.3: Team Communication

- Document migration completion
- Share new component patterns with team
- Schedule team training if needed

## Component Migration Patterns

### Button Components

**Before:**
```go
@Button("Save")
@Button("Save", "primary")  
@Button("Save", "primary", "lg")
@Button("Save", "primary", "lg", true) // disabled
```

**After:**
```go
@atoms.Button(atoms.ButtonProps{Text: "Save"})
@atoms.Button(atoms.ButtonProps{Text: "Save", Variant: atoms.ButtonPrimary})
@atoms.Button(atoms.ButtonProps{Text: "Save", Variant: atoms.ButtonPrimary, Size: atoms.ButtonSizeLG})
@atoms.Button(atoms.ButtonProps{Text: "Save", Variant: atoms.ButtonPrimary, Size: atoms.ButtonSizeLG, Disabled: true})
```

### Form Field Components

**Before:**
```go
@FormField("email", "email", "Enter your email")
@FormField("email", "email", "Enter your email", true) // required
```

**After:**
```go
@molecules.FormField(molecules.FormFieldProps{
    Name: "email",
    Type: molecules.FormFieldEmail,
    Placeholder: "Enter your email",
})

@molecules.FormField(molecules.FormFieldProps{
    Name: "email", 
    Type: molecules.FormFieldEmail,
    Placeholder: "Enter your email",
    Required: true,
})
```

### Input Components

**Before:**
```go
@Input("search", "Search...")
@Input("search", "Search...", "text-sm")
```

**After:**
```go
@atoms.Input(atoms.InputProps{
    Name: "search",
    Placeholder: "Search...",
})

@atoms.Input(atoms.InputProps{
    Name: "search",
    Placeholder: "Search...", 
    Size: atoms.InputSizeSM,
})
```

### Complex HTMX Integration

**Before:**
```html
<button hx-post="/api/save" hx-target="#result">Save</button>
```

**After:**
```go
@atoms.Button(atoms.ButtonProps{
    Text: "Save",
    HXPost: "/api/save",
    HXTarget: "#result",
})
```

### Alpine.js Integration

**Before:**
```html
<div x-data="{ open: false }" x-show="open">
    <button x-on:click="open = true">Open</button>
</div>
```

**After:**
```go
@molecules.Modal(molecules.ModalProps{
    AlpineData: "{ open: false }",
    AlpineShow: "open",
    Trigger: atoms.ButtonProps{
        Text: "Open",
        AlpineClick: "open = true",
    },
})
```

## Validation and Testing

### Automated Validation Checks

The validator checks for:

1. **Syntax Errors**: Ensures all Go and Templ files have valid syntax
2. **Import Issues**: Verifies correct import paths for atoms/molecules
3. **Component Usage**: Ensures components use proper props structures
4. **Theme Compliance**: Checks for hardcoded styles and theme.Get() calls
5. **HTMX/Alpine Integration**: Validates proper attribute passing

### Manual Testing Checklist

- [ ] All pages load without errors
- [ ] Forms submit correctly
- [ ] Buttons trigger expected actions
- [ ] Modal dialogs open/close properly
- [ ] Navigation works as expected
- [ ] Responsive design maintains layout
- [ ] Dark/light theme switching works
- [ ] Interactive components respond correctly
- [ ] Error states display properly
- [ ] Loading states work as expected

### Performance Validation

- [ ] JavaScript bundle size hasn't increased significantly
- [ ] CSS bundle is properly optimized
- [ ] Page load times are acceptable
- [ ] Component rendering is smooth
- [ ] No memory leaks in interactive components

## Rollback Procedures

### When to Consider Rollback

Consider rolling back if you encounter:
- Critical functionality is broken
- Performance has significantly degraded
- Major visual issues across the application
- Security vulnerabilities introduced
- Timeline constraints prevent fixing issues

### Rollback Options

#### 1. Full Rollback
Restores all files from backup:

```bash
./rollback full /path/to/project /backup/path
```

#### 2. Selective Rollback
Rollback specific files or patterns:

```bash
./rollback selective /path/to/project /backup/path "views/components/atoms/*"
```

#### 3. Partial Rollback
Keep some changes while rolling back others:

```bash
./rollback partial /path/to/project /backup/path
```

### Post-Rollback Steps

1. **Verify Functionality**: Test that critical features work
2. **Update Documentation**: Note what was rolled back and why
3. **Plan Re-migration**: Address issues that caused rollback
4. **Team Communication**: Inform team of rollback and next steps

## Troubleshooting

### Common Issues and Solutions

#### Issue: Templ Generation Fails
**Symptoms:** `templ generate` produces errors
**Solution:**
```bash
# Check for syntax errors
go run validator.go /path/to/project

# Fix any syntax issues reported
# Regenerate templ files
templ generate
```

#### Issue: Import Errors
**Symptoms:** `package X is not in GOROOT or GOPATH`
**Solution:**
```bash
# Update go.mod
go mod tidy

# Verify import paths are correct
# Should be: "github.com/niiniyare/ruun/views/components/atoms"
```

#### Issue: Component Props Errors
**Symptoms:** Type errors when using components
**Solution:**
```go
// Wrong
@atoms.Button("Save")

// Correct
@atoms.Button(atoms.ButtonProps{Text: "Save"})
```

#### Issue: Theme Classes Not Working
**Symptoms:** Styling doesn't apply correctly
**Solution:**
1. Ensure theme compilation is up to date
2. Check CSS classes are properly generated
3. Verify theme tokens are correctly mapped

#### Issue: HTMX/Alpine Not Working
**Symptoms:** Interactive features broken
**Solution:**
```go
// Wrong - attributes in template
<button hx-post="/api/save">Save</button>

// Correct - via props
@atoms.Button(atoms.ButtonProps{
    Text: "Save",
    HXPost: "/api/save",
})
```

### Getting Help

1. **Check Validation Report**: Look for specific error messages
2. **Review Migration Session**: Check what transformations were applied
3. **Compare with Examples**: Use this guide's examples as reference
4. **Test in Isolation**: Create minimal test cases for problematic components
5. **Use Rollback**: If stuck, rollback and retry with manual approach

## Best Practices

### Migration Strategy

1. **Start Small**: Begin with simple atoms before complex organisms
2. **Test Frequently**: Validate after each component type migration
3. **Use Version Control**: Commit after each successful phase
4. **Document Changes**: Keep notes on manual modifications needed
5. **Plan Rollback**: Always have a rollback strategy ready

### Code Organization

1. **Consistent Props**: Use structured props for all components
2. **Theme Tokens**: Always use theme tokens instead of hardcoded values
3. **Component Composition**: Favor composition over complex single components
4. **Type Safety**: Leverage Go's type system for component interfaces

### Performance Optimization

1. **Compiled Classes**: Use compiled CSS classes for better performance
2. **Minimal Props**: Only pass necessary props to components
3. **Lazy Loading**: Consider lazy loading for complex organisms
4. **Bundle Optimization**: Monitor and optimize bundle sizes

### Maintenance

1. **Regular Validation**: Run validator periodically to catch issues early
2. **Update Dependencies**: Keep migration tools and framework updated
3. **Team Training**: Ensure team understands new patterns
4. **Documentation**: Keep component documentation current

## Migration Completion Checklist

- [ ] All automated migrations applied successfully
- [ ] Validation passes without critical errors
- [ ] Manual migrations completed for complex components
- [ ] All tests pass
- [ ] Visual testing completed
- [ ] Performance validation passed
- [ ] Documentation updated
- [ ] Team informed of changes
- [ ] Migration artifacts archived
- [ ] Rollback plan documented

## Next Steps

After successful migration:

1. **Explore Enhanced Features**: Take advantage of new component capabilities
2. **Optimize Performance**: Fine-tune theme and component usage
3. **Extend Components**: Create new atoms/molecules as needed
4. **Share Knowledge**: Document lessons learned for future migrations
5. **Plan Enhancements**: Consider additional improvements enabled by new architecture

---

**Need Help?** 
- Check the troubleshooting section above
- Review component documentation
- Test with minimal examples
- Use the rollback tool if needed