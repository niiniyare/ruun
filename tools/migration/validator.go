package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// ValidationRule defines what to check after migration
type ValidationRule struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"` // "syntax", "import", "component", "compilation"
	Pattern     string   `json:"pattern,omitempty"`
	Required    bool     `json:"required"`
	FileTypes   []string `json:"file_types"`
	ErrorMsg    string   `json:"error_msg"`
}

// ValidationIssue represents a validation problem
type ValidationIssue struct {
	ID       string `json:"id"`
	RuleID   string `json:"rule_id"`
	File     string `json:"file"`
	Line     int    `json:"line"`
	Column   int    `json:"column"`
	Code     string `json:"code"`
	Severity string `json:"severity"` // "error", "warning", "info"
	Message  string `json:"message"`
	Fix      string `json:"fix,omitempty"`
}

// ValidationReport contains validation results
type ValidationReport struct {
	Timestamp     time.Time           `json:"timestamp"`
	ProjectPath   string              `json:"project_path"`
	SessionID     string              `json:"session_id,omitempty"`
	FilesChecked  int                 `json:"files_checked"`
	Issues        []ValidationIssue   `json:"issues"`
	Summary       map[string]int      `json:"summary"`
	Rules         []ValidationRule    `json:"rules"`
	PassedRules   []string            `json:"passed_rules"`
	FailedRules   []string            `json:"failed_rules"`
	CompileErrors []CompileError      `json:"compile_errors,omitempty"`
}

// CompileError represents a compilation error
type CompileError struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
	Message string `json:"message"`
}

// MigrationValidator validates migrated code
type MigrationValidator struct {
	rules        []ValidationRule
	fileSet      *token.FileSet
	issues       []ValidationIssue
	filesChecked int
	passedRules  map[string]bool
	failedRules  map[string]bool
}

// NewMigrationValidator creates a new validator
func NewMigrationValidator() *MigrationValidator {
	return &MigrationValidator{
		rules:        getDefaultValidationRules(),
		fileSet:      token.NewFileSet(),
		issues:       make([]ValidationIssue, 0),
		filesChecked: 0,
		passedRules:  make(map[string]bool),
		failedRules:  make(map[string]bool),
	}
}

// getDefaultValidationRules returns predefined validation rules
func getDefaultValidationRules() []ValidationRule {
	return []ValidationRule{
		// Syntax validation
		{
			ID:          "valid-go-syntax",
			Name:        "Valid Go Syntax",
			Description: "All Go files should have valid syntax",
			Type:        "syntax",
			Required:    true,
			FileTypes:   []string{".go"},
			ErrorMsg:    "Go file has syntax errors",
		},
		{
			ID:          "valid-templ-syntax",
			Name:        "Valid Templ Syntax", 
			Description: "All templ files should have valid syntax",
			Type:        "syntax",
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Templ file has syntax errors",
		},
		
		// Import validation
		{
			ID:          "no-old-imports",
			Name:        "No Old Import Paths",
			Description: "Should not import old component paths",
			Type:        "import",
			Pattern:     `import\s+"[^"]*views/components/[^/]+\.templ"`,
			Required:    true,
			FileTypes:   []string{".go", ".templ"},
			ErrorMsg:    "Found import of old component path - should use atoms/molecules/organisms",
		},
		{
			ID:          "atoms-imports-present",
			Name:        "Atoms Import Present",
			Description: "Files using atoms should import atoms package",
			Type:        "import",
			Pattern:     `@atoms\.\w+`,
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Using atoms components but missing import",
		},
		{
			ID:          "molecules-imports-present",
			Name:        "Molecules Import Present",
			Description: "Files using molecules should import molecules package", 
			Type:        "import",
			Pattern:     `@molecules\.\w+`,
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Using molecules components but missing import",
		},
		
		// Component usage validation
		{
			ID:          "button-props-struct",
			Name:        "Button Uses Props Struct",
			Description: "Button components should use props struct",
			Type:        "component",
			Pattern:     `@atoms\.Button\([^{]`,
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Button should use ButtonProps struct",
		},
		{
			ID:          "formfield-props-struct",
			Name:        "FormField Uses Props Struct",
			Description: "FormField components should use props struct",
			Type:        "component",
			Pattern:     `@molecules\.FormField\([^{]`,
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "FormField should use FormFieldProps struct",
		},
		{
			ID:          "input-props-struct",
			Name:        "Input Uses Props Struct",
			Description: "Input components should use props struct",
			Type:        "component",
			Pattern:     `@atoms\.Input\([^{]`,
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Input should use InputProps struct",
		},
		
		// Theme validation
		{
			ID:          "no-theme-get-calls",
			Name:        "No theme.Get() Calls",
			Description: "Should not use dynamic theme.Get() calls",
			Type:        "component",
			Pattern:     `theme\.Get\(`,
			Required:    true,
			FileTypes:   []string{".go", ".templ"},
			ErrorMsg:    "Found theme.Get() call - should use compiled CSS classes",
		},
		{
			ID:          "no-hardcoded-colors",
			Name:        "No Hardcoded Colors",
			Description: "Should not use hardcoded Tailwind color classes",
			Type:        "component",
			Pattern:     `class="[^"]*(?:text-red-|bg-blue-|border-green-)\d+`,
			Required:    false,
			FileTypes:   []string{".templ", ".html"},
			ErrorMsg:    "Found hardcoded color - should use theme tokens",
		},
		
		// HTMX validation
		{
			ID:          "htmx-via-props",
			Name:        "HTMX via Props",
			Description: "HTMX attributes should be passed via component props",
			Type:        "component",
			Pattern:     `hx-(?:post|get|put|delete)="`,
			Required:    false,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "HTMX attribute found in template - should pass via props",
		},
		
		// Alpine.js validation
		{
			ID:          "alpine-via-props",
			Name:        "Alpine.js via Props",
			Description: "Alpine directives should be passed via component props",
			Type:        "component",
			Pattern:     `x-(?:data|show|if|for|on)="`,
			Required:    false,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Alpine directive found in template - should pass via props",
		},
		
		// Compilation validation
		{
			ID:          "templ-generates",
			Name:        "Templ Files Generate",
			Description: "Templ files should generate Go code successfully",
			Type:        "compilation",
			Required:    true,
			FileTypes:   []string{".templ"},
			ErrorMsg:    "Templ file failed to generate Go code",
		},
		{
			ID:          "go-builds",
			Name:        "Go Code Builds",
			Description: "Go code should compile successfully",
			Type:        "compilation",
			Required:    true,
			FileTypes:   []string{".go"},
			ErrorMsg:    "Go code failed to compile",
		},
	}
}

// ValidateProject validates an entire project after migration
func (v *MigrationValidator) ValidateProject(projectPath string, sessionID string) (*ValidationReport, error) {
	fmt.Printf("Starting validation for: %s\n", projectPath)
	
	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories and irrelevant files
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == "node_modules" || name == ".git" ||
			   name == "dist" || name == "build" || name == "migration_backups" {
				return filepath.SkipDir
			}
			return nil
		}
		
		if v.isRelevantFile(path) {
			return v.validateFile(path)
		}
		
		return nil
	})
	
	if err != nil {
		return nil, fmt.Errorf("error walking project: %w", err)
	}
	
	return v.generateReport(projectPath, sessionID), nil
}

// isRelevantFile checks if a file should be validated
func (v *MigrationValidator) isRelevantFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".go" || ext == ".templ" || ext == ".html"
}

// validateFile validates a single file
func (v *MigrationValidator) validateFile(filePath string) error {
	v.filesChecked++
	
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	
	ext := filepath.Ext(filePath)
	
	// Syntax validation
	if err := v.validateSyntax(filePath, string(content), ext); err != nil {
		v.addIssue(ValidationIssue{
			ID:       fmt.Sprintf("syntax-%s", filePath),
			RuleID:   v.getSyntaxRuleID(ext),
			File:     filePath,
			Line:     1,
			Column:   1,
			Code:     err.Error(),
			Severity: "error",
			Message:  fmt.Sprintf("Syntax error: %v", err),
		})
	}
	
	// Pattern-based validation
	return v.validatePatterns(filePath, string(content), ext)
}

// validateSyntax validates file syntax
func (v *MigrationValidator) validateSyntax(filePath, content, ext string) error {
	switch ext {
	case ".go":
		fset := token.NewFileSet()
		_, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
		if err != nil {
			v.failedRules["valid-go-syntax"] = true
			return err
		}
		v.passedRules["valid-go-syntax"] = true
		
	case ".templ":
		// Basic templ validation - check for common syntax errors
		if err := v.validateTemplSyntax(content); err != nil {
			v.failedRules["valid-templ-syntax"] = true
			return err
		}
		v.passedRules["valid-templ-syntax"] = true
	}
	
	return nil
}

// validateTemplSyntax performs basic templ syntax validation
func (v *MigrationValidator) validateTemplSyntax(content string) error {
	// Check for balanced braces in templ expressions
	braceCount := 0
	for _, char := range content {
		if char == '{' {
			braceCount++
		} else if char == '}' {
			braceCount--
			if braceCount < 0 {
				return fmt.Errorf("unbalanced closing brace")
			}
		}
	}
	
	if braceCount > 0 {
		return fmt.Errorf("unbalanced opening brace")
	}
	
	// Check for valid templ component declarations
	templRegex := regexp.MustCompile(`templ\s+\w+\s*\([^)]*\)\s*{`)
	if strings.Contains(content, "templ ") && !templRegex.MatchString(content) {
		return fmt.Errorf("invalid templ component declaration")
	}
	
	return nil
}

// validatePatterns validates content against pattern-based rules
func (v *MigrationValidator) validatePatterns(filePath, content, ext string) error {
	lines := strings.Split(content, "\n")
	
	for _, rule := range v.rules {
		if rule.Type == "syntax" || rule.Type == "compilation" {
			continue // Already handled
		}
		
		if !v.ruleAppliesTo(rule, ext) {
			continue
		}
		
		if rule.Pattern == "" {
			continue
		}
		
		regex, err := regexp.Compile(rule.Pattern)
		if err != nil {
			fmt.Printf("Warning: Invalid regex in rule %s: %v\n", rule.ID, err)
			continue
		}
		
		// Check if pattern exists in content
		if regex.MatchString(content) {
			// Find specific lines with issues
			for lineNum, line := range lines {
				if regex.MatchString(line) {
					matches := regex.FindStringSubmatch(line)
					
					issue := ValidationIssue{
						ID:       fmt.Sprintf("%s-%s-%d", rule.ID, filePath, lineNum+1),
						RuleID:   rule.ID,
						File:     filePath,
						Line:     lineNum + 1,
						Column:   strings.Index(line, matches[0]) + 1,
						Code:     matches[0],
						Severity: v.getSeverity(rule),
						Message:  rule.ErrorMsg,
						Fix:      v.getSuggestedFix(rule, matches[0]),
					}
					
					v.addIssue(issue)
					v.failedRules[rule.ID] = true
				}
			}
		} else {
			// Rule passed - pattern not found where it shouldn't be
			if rule.Required {
				v.passedRules[rule.ID] = true
			}
		}
		
		// Special handling for required imports
		if rule.Type == "import" && rule.Required {
			v.validateRequiredImports(filePath, content, rule)
		}
	}
	
	return nil
}

// validateRequiredImports checks if required imports are present
func (v *MigrationValidator) validateRequiredImports(filePath, content string, rule ValidationRule) {
	// Check if atoms/molecules are used without imports
	if rule.ID == "atoms-imports-present" {
		if strings.Contains(content, "@atoms.") && !strings.Contains(content, `"github.com/niiniyare/ruun/views/components/atoms"`) {
			issue := ValidationIssue{
				ID:       fmt.Sprintf("%s-%s", rule.ID, filePath),
				RuleID:   rule.ID,
				File:     filePath,
				Line:     1,
				Column:   1,
				Code:     "@atoms.",
				Severity: "error",
				Message:  "Using atoms components but missing import statement",
				Fix:      `Add: import "github.com/niiniyare/ruun/views/components/atoms"`,
			}
			v.addIssue(issue)
			v.failedRules[rule.ID] = true
		}
	}
	
	if rule.ID == "molecules-imports-present" {
		if strings.Contains(content, "@molecules.") && !strings.Contains(content, `"github.com/niiniyare/ruun/views/components/molecules"`) {
			issue := ValidationIssue{
				ID:       fmt.Sprintf("%s-%s", rule.ID, filePath),
				RuleID:   rule.ID,
				File:     filePath,
				Line:     1,
				Column:   1,
				Code:     "@molecules.",
				Severity: "error",
				Message:  "Using molecules components but missing import statement",
				Fix:      `Add: import "github.com/niiniyare/ruun/views/components/molecules"`,
			}
			v.addIssue(issue)
			v.failedRules[rule.ID] = true
		}
	}
}

// ruleAppliesTo checks if a rule applies to a file type
func (v *MigrationValidator) ruleAppliesTo(rule ValidationRule, ext string) bool {
	for _, fileType := range rule.FileTypes {
		if fileType == ext {
			return true
		}
	}
	return false
}

// getSyntaxRuleID returns the syntax rule ID for a file extension
func (v *MigrationValidator) getSyntaxRuleID(ext string) string {
	switch ext {
	case ".go":
		return "valid-go-syntax"
	case ".templ":
		return "valid-templ-syntax"
	default:
		return "unknown-syntax"
	}
}

// getSeverity determines issue severity based on rule
func (v *MigrationValidator) getSeverity(rule ValidationRule) string {
	if rule.Required {
		return "error"
	}
	return "warning"
}

// getSuggestedFix provides a suggested fix for common issues
func (v *MigrationValidator) getSuggestedFix(rule ValidationRule, code string) string {
	switch rule.ID {
	case "no-old-imports":
		return "Update import to use atoms/molecules/organisms package"
	case "button-props-struct":
		return "Use @atoms.Button(atoms.ButtonProps{...})"
	case "formfield-props-struct":
		return "Use @molecules.FormField(molecules.FormFieldProps{...})"
	case "no-theme-get-calls":
		return "Use compiled CSS classes instead of theme.Get()"
	case "no-hardcoded-colors":
		return "Use theme tokens like text-error, bg-primary"
	case "htmx-via-props":
		return "Pass HTMX attributes via component props"
	case "alpine-via-props":
		return "Pass Alpine directives via component props"
	default:
		return ""
	}
}

// addIssue adds a validation issue
func (v *MigrationValidator) addIssue(issue ValidationIssue) {
	v.issues = append(v.issues, issue)
}

// generateReport creates a validation report
func (v *MigrationValidator) generateReport(projectPath, sessionID string) *ValidationReport {
	summary := make(map[string]int)
	
	// Count issues by severity
	for _, issue := range v.issues {
		summary[issue.Severity]++
		summary["total"]++
	}
	
	// Count passed/failed rules
	var passedRules, failedRules []string
	for ruleID := range v.passedRules {
		passedRules = append(passedRules, ruleID)
	}
	for ruleID := range v.failedRules {
		failedRules = append(failedRules, ruleID)
	}
	
	summary["passed_rules"] = len(passedRules)
	summary["failed_rules"] = len(failedRules)
	
	return &ValidationReport{
		Timestamp:    time.Now(),
		ProjectPath:  projectPath,
		SessionID:    sessionID,
		FilesChecked: v.filesChecked,
		Issues:       v.issues,
		Summary:      summary,
		Rules:        v.rules,
		PassedRules:  passedRules,
		FailedRules:  failedRules,
	}
}

// SaveReport saves the validation report
func (v *MigrationValidator) SaveReport(report *ValidationReport, outputPath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling report: %w", err)
	}
	
	return os.WriteFile(outputPath, data, 0644)
}

// PrintSummary prints a human-readable summary
func (v *MigrationValidator) PrintSummary(report *ValidationReport) {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("MIGRATION VALIDATION REPORT\n")
	fmt.Printf(strings.Repeat("=", 60) + "\n")
	
	fmt.Printf("Project Path: %s\n", report.ProjectPath)
	if report.SessionID != "" {
		fmt.Printf("Migration Session: %s\n", report.SessionID)
	}
	fmt.Printf("Files Checked: %d\n", report.FilesChecked)
	fmt.Printf("Validation Time: %s\n", report.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("\n")
	
	// Overall status
	if report.Summary["total"] == 0 {
		fmt.Printf("✅ VALIDATION PASSED - No issues found!\n")
	} else {
		fmt.Printf("❌ VALIDATION FAILED - Found %d issues\n", report.Summary["total"])
	}
	fmt.Printf("\n")
	
	// Issues by severity
	fmt.Printf("Issues by Severity:\n")
	fmt.Printf("  - Errors: %d\n", report.Summary["error"])
	fmt.Printf("  - Warnings: %d\n", report.Summary["warning"])
	fmt.Printf("  - Info: %d\n", report.Summary["info"])
	fmt.Printf("\n")
	
	// Rules summary
	fmt.Printf("Rules Summary:\n")
	fmt.Printf("  - Passed: %d rules\n", report.Summary["passed_rules"])
	fmt.Printf("  - Failed: %d rules\n", report.Summary["failed_rules"])
	fmt.Printf("\n")
	
	// Show failed rules
	if len(report.FailedRules) > 0 {
		fmt.Printf("Failed Rules:\n")
		for _, ruleID := range report.FailedRules {
			for _, rule := range report.Rules {
				if rule.ID == ruleID {
					fmt.Printf("  - %s: %s\n", rule.Name, rule.Description)
					break
				}
			}
		}
		fmt.Printf("\n")
	}
	
	// Show critical issues
	criticalIssues := 0
	for _, issue := range report.Issues {
		if issue.Severity == "error" {
			criticalIssues++
		}
	}
	
	if criticalIssues > 0 {
		fmt.Printf("Critical Issues (first 5):\n")
		count := 0
		for _, issue := range report.Issues {
			if issue.Severity == "error" && count < 5 {
				fmt.Printf("  %s:%d - %s\n", issue.File, issue.Line, issue.Message)
				if issue.Fix != "" {
					fmt.Printf("    Fix: %s\n", issue.Fix)
				}
				count++
			}
		}
	}
	
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	
	if report.Summary["total"] == 0 {
		fmt.Printf("🎉 Migration validation successful! Your code is ready to use.\n")
	} else {
		fmt.Printf("🔧 Please fix the validation issues above and run validation again.\n")
		fmt.Printf("📝 See the detailed report for complete issue list and fixes.\n")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run validator.go <project-path> [--session=session-id] [--output=report.json]")
		fmt.Println("\nOptions:")
		fmt.Println("  --session=ID: Associate with a migration session ID")
		fmt.Println("  --output=FILE: Save validation report to JSON file")
		os.Exit(1)
	}
	
	projectPath := os.Args[1]
	sessionID := ""
	outputFile := "validation-report.json"
	
	// Parse command line arguments
	for _, arg := range os.Args[2:] {
		if strings.HasPrefix(arg, "--session=") {
			sessionID = strings.TrimPrefix(arg, "--session=")
		} else if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
		}
	}
	
	validator := NewMigrationValidator()
	
	report, err := validator.ValidateProject(projectPath, sessionID)
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		os.Exit(1)
	}
	
	// Save report
	if err := validator.SaveReport(report, outputFile); err != nil {
		fmt.Printf("Error saving report: %v\n", err)
		os.Exit(1)
	}
	
	// Print summary
	validator.PrintSummary(report)
	
	fmt.Printf("\nValidation report saved to: %s\n", outputFile)
	
	// Exit with error code if validation failed
	if report.Summary["total"] > 0 {
		os.Exit(1)
	}
}