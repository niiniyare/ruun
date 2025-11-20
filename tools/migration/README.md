# Migration Tools

Comprehensive migration system for transitioning from old component patterns to the new atomic design architecture.

## Overview

This migration toolkit provides safe, automated, and incremental migration capabilities with full rollback support. The system includes analysis, automated transformation, validation, testing, and rollback tools.

## Quick Start

### 1. Run Complete Migration Pipeline

```bash
cd tools/migration
go run migrate.go /path/to/your/project
```

This will:
- Analyze your codebase
- Create backups
- Apply automated migrations
- Validate results
- Run tests
- Provide detailed reporting

### 2. Dry Run (Preview Only)

```bash
go run migrate.go /path/to/project --dry-run
```

Preview all changes without modifying files.

### 3. Individual Tools

For more control, run tools individually:

```bash
# 1. Analyze codebase
go run analyzer.go /path/to/project

# 2. Apply migrations
go run migrator.go /path/to/project

# 3. Validate results
go run validator.go /path/to/project

# 4. Run tests
go run test_framework.go run migration-test-suite.json

# 5. Rollback if needed
go run rollback.go list /path/to/project
go run rollback.go full /path/to/project /backup/path
```

## Tool Details

### 1. Migration Analyzer (`analyzer.go`)

Scans codebase to identify patterns that need migration.

**Features:**
- Detects old import statements
- Identifies outdated component usage
- Finds hardcoded styles and colors
- Checks theme.Get() calls
- Reports HTMX/Alpine patterns

**Usage:**
```bash
go run analyzer.go <project-path> [output-file]

# Example
go run analyzer.go ../../ analysis-report.json
```

**Output:** JSON report with detailed findings and migration recommendations.

### 2. Automated Migrator (`migrator.go`)

Applies automated code transformations.

**Features:**
- Updates import statements
- Converts component calls to props structs
- Transforms CSS classes to theme tokens
- Migrates HTMX/Alpine attributes
- Creates automatic backups

**Usage:**
```bash
go run migrator.go <project-path> [options]

Options:
  --dry-run           Preview changes without applying
  --rules=FILE        Use custom migration rules
  --output=FILE       Save session details
```

**Transformations:**

| Old Pattern | New Pattern |
|-------------|-------------|
| `@Button("Save", "primary")` | `@atoms.Button(atoms.ButtonProps{Text: "Save", Variant: atoms.ButtonPrimary})` |
| `@FormField("email", "email", "Enter email")` | `@molecules.FormField(molecules.FormFieldProps{Name: "email", Type: molecules.FormFieldEmail, Placeholder: "Enter email"})` |
| `class="text-red-500"` | `class="text-error"` |
| `hx-post="/api/save"` | `HXPost: "/api/save"` |
| `theme.Get("button.primary")` | `"button-primary"` |

### 3. Migration Validator (`validator.go`)

Validates that migrations were applied correctly.

**Features:**
- Syntax validation (Go/Templ)
- Import statement verification
- Component usage validation
- Theme compliance checking
- Compilation testing

**Usage:**
```bash
go run validator.go <project-path> [options]

Options:
  --session=ID        Link to migration session
  --output=FILE       Save validation report
```

**Validation Rules:**
- ✅ All Go/Templ files have valid syntax
- ✅ No old import paths remain
- ✅ Components use proper props structs
- ✅ No hardcoded colors or theme.Get() calls
- ✅ HTMX/Alpine attributes passed via props

### 4. Rollback Tool (`rollback.go`)

Provides safe rollback capabilities.

**Features:**
- Full project rollback
- Selective file rollback
- Partial rollback with custom rules
- Backup management
- Dry run support

**Usage:**
```bash
# List available backups
go run rollback.go list <project-path>

# Full rollback
go run rollback.go full <project-path> <backup-path>

# Selective rollback (specific files/patterns)
go run rollback.go selective <project-path> <backup-path> "*.templ" "views/*"

# Partial rollback (with custom rules)
go run rollback.go partial <project-path> <backup-path>

Options:
  --dry-run           Preview rollback without applying
  --migration=ID      Link to migration session
  --output=FILE       Save rollback session
```

### 5. Test Framework (`test_framework.go`)

Comprehensive testing system for migration validation.

**Features:**
- Unit tests for component migrations
- Integration tests for syntax validation
- Visual regression testing capabilities
- Performance testing
- Custom test case support

**Usage:**
```bash
# Generate default test suite
go run test_framework.go generate-default

# Run test suite
go run test_framework.go run migration-test-suite.json

Options:
  --output=FILE       Save test report
  --verbose           Show detailed output
```

**Test Types:**
- **Unit Tests**: Individual component migrations
- **Integration Tests**: End-to-end migration validation  
- **Syntax Tests**: Go/Templ compilation verification
- **Performance Tests**: Bundle size and runtime validation

### 6. Migration CLI (`migrate.go`)

Orchestrates the complete migration pipeline.

**Features:**
- Automated pipeline execution
- Interactive confirmation prompts
- Comprehensive logging
- Phase-by-phase execution
- Detailed reporting

**Usage:**
```bash
go run migrate.go [project-path] [options]

Options:
  --dry-run              Preview migration without changes
  --verbose, -v          Show verbose output  
  --no-interactive       Skip user prompts
  --project=PATH         Specify project path
  --backup-dir=DIR       Custom backup directory
  --output-dir=DIR       Custom output directory
```

**Pipeline Phases:**
1. **Analysis** - Scan codebase for migration requirements
2. **Backup** - Create backup of current code
3. **Migration** - Apply automated transformations
4. **Validation** - Verify migration correctness
5. **Testing** - Run migration test suite
6. **Summary** - Provide results and recommendations

## Migration Patterns

### Component Migrations

**Button Components:**
```go
// Before
@Button("Save")
@Button("Save", "primary", "lg")

// After  
@atoms.Button(atoms.ButtonProps{Text: "Save"})
@atoms.Button(atoms.ButtonProps{
    Text: "Save", 
    Variant: atoms.ButtonPrimary,
    Size: atoms.ButtonSizeLG,
})
```

**Form Fields:**
```go
// Before
@FormField("email", "email", "Enter your email")

// After
@molecules.FormField(molecules.FormFieldProps{
    Name: "email",
    Type: molecules.FormFieldEmail,
    Placeholder: "Enter your email",
})
```

### Import Migrations

```go
// Before
import "views/components/button.templ"
import "views/components/formfield.templ"

// After
import "github.com/niiniyare/ruun/views/components/atoms"
import "github.com/niiniyare/ruun/views/components/molecules"
```

### Theme Migrations

```go
// Before
class={ theme.Get("button.primary.background") }
class="text-red-500 bg-blue-600"

// After
class="button-primary"
class="text-error bg-primary"
```

### HTMX/Alpine Migrations

```html
<!-- Before -->
<button hx-post="/api/save" x-on:click="handleClick()">Save</button>

<!-- After -->
@atoms.Button(atoms.ButtonProps{
    Text: "Save",
    HXPost: "/api/save",
    AlpineClick: "handleClick()",
})
```

## Configuration

### Custom Migration Rules

Create `custom-rules.json`:

```json
{
  "rules": [
    {
      "id": "custom-component-migration",
      "name": "Custom Component Migration",
      "description": "Migrate custom component patterns",
      "pattern": "@MyComponent\\(([^)]+)\\)",
      "replacement": "@atoms.MyComponent(atoms.MyComponentProps{$1})",
      "file_types": [".templ"]
    }
  ]
}
```

Use with migrator:
```bash
go run migrator.go /path/to/project --rules=custom-rules.json
```

### Custom Test Cases

Create custom test cases in `custom-tests.json`:

```json
{
  "id": "custom-test-suite",
  "name": "Custom Migration Tests",
  "tests": [
    {
      "id": "my-component-test",
      "name": "My Component Migration",
      "type": "unit",
      "pre_state": {
        "files": [
          {
            "path": "test.templ",
            "content": "@MyComponent(\"test\")",
            "type": "input"
          }
        ]
      },
      "commands": [
        {
          "name": "migrate",
          "command": "go",
          "args": ["run", "migrator.go", "."],
          "timeout": "30s",
          "expected_exit": 0
        }
      ],
      "assertions": [
        {
          "type": "file_content",
          "target": "test.templ", 
          "expected": "MyComponentProps",
          "description": "Should use props struct"
        }
      ]
    }
  ]
}
```

## Best Practices

### Before Migration

1. **Create Git Commit**: Ensure clean working directory
2. **Run Analysis**: Understand scope of changes
3. **Test Current State**: Verify everything works
4. **Plan Downtime**: Migration may take 30-60 minutes

### During Migration

1. **Start with Dry Run**: Preview all changes first
2. **Run in Stages**: Use individual tools for complex projects
3. **Validate Frequently**: Check results after each phase
4. **Monitor Logs**: Watch for warnings and errors

### After Migration

1. **Run Full Validation**: Ensure no issues remain
2. **Test Thoroughly**: Check all functionality
3. **Update Documentation**: Reflect new patterns
4. **Train Team**: Share new component usage

## Troubleshooting

### Common Issues

**Import Errors:**
```bash
# Fix missing imports
go mod tidy

# Verify import paths
grep -r "github.com/niiniyare/ruun/views/components" .
```

**Syntax Errors:**
```bash
# Check templ syntax
templ generate

# Validate Go syntax
go build ./...
```

**Component Props Errors:**
```go
// Wrong
@atoms.Button("Save")

// Correct  
@atoms.Button(atoms.ButtonProps{Text: "Save"})
```

**Theme Classes Not Working:**
```bash
# Rebuild theme
go run cmd/build-themes/main.go

# Verify CSS classes
grep -r "button-primary" assets/css/
```

### Getting Help

1. **Check Validation Report**: Look for specific errors
2. **Review Migration Logs**: Check detailed execution logs
3. **Use Rollback**: Restore backup if needed
4. **Run Individual Tools**: Isolate specific issues
5. **Check Examples**: Review this documentation

### Emergency Rollback

If migration causes critical issues:

```bash
# List available backups
go run rollback.go list /path/to/project

# Full rollback to last backup
go run rollback.go full /path/to/project /backup/path

# Verify rollback
go run validator.go /path/to/project
```

## File Organization

```
tools/migration/
├── README.md                 # This file
├── MIGRATION_GUIDE.md        # Detailed migration guide
├── migrate.go                # Main CLI orchestrator
├── analyzer.go               # Analysis tool
├── migrator.go               # Migration tool
├── validator.go              # Validation tool
├── rollback.go               # Rollback tool
├── test_framework.go         # Testing framework
├── migration-test-suite.json # Default test suite
└── examples/
    ├── custom-rules.json     # Example custom rules
    └── custom-tests.json     # Example custom tests
```

## Integration with CI/CD

### GitHub Actions Example

```yaml
name: Migration Validation
on: [push, pull_request]

jobs:
  validate-migration:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v3
        with:
          go-version: '1.21'
      
      - name: Run Migration Analysis
        run: |
          cd tools/migration
          go run analyzer.go ../../
      
      - name: Validate Migration State  
        run: |
          cd tools/migration
          go run validator.go ../../
      
      - name: Run Migration Tests
        run: |
          cd tools/migration
          go run test_framework.go run migration-test-suite.json
```

### Pre-commit Hooks

```bash
#!/bin/sh
# .git/hooks/pre-commit

# Run migration validation before commits
cd tools/migration
go run validator.go ../../

if [ $? -ne 0 ]; then
    echo "Migration validation failed. Please fix issues before committing."
    exit 1
fi
```

## Performance

### Metrics

- **Analysis**: ~30 seconds for 100 files
- **Migration**: ~60 seconds for 100 files  
- **Validation**: ~15 seconds for 100 files
- **Testing**: ~45 seconds for full suite

### Optimization

- Run on dedicated environment for large projects
- Use `--dry-run` for quick previews
- Selective migrations for incremental updates
- Parallel processing planned for future versions

## Changelog

### v1.0.0
- Initial release
- Complete migration pipeline
- All core tools implemented
- Comprehensive test suite
- Full documentation

### Future Plans
- Interactive migration wizard
- Visual diff support
- Parallel processing
- Plugin system for custom migrations
- Integration with popular editors

---

**Need Help?** Check the [Migration Guide](MIGRATION_GUIDE.md) for detailed step-by-step instructions.