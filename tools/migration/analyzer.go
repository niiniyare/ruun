package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// MigrationPattern represents an old pattern that needs migration
type MigrationPattern struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Pattern     string   `json:"pattern"`
	Type        string   `json:"type"`     // "import", "component", "prop", "class", "theme"
	Severity    string   `json:"severity"` // "error", "warning", "info"
	AutoFix     bool     `json:"auto_fix"`
	NewPattern  string   `json:"new_pattern,omitempty"`
	Examples    []string `json:"examples,omitempty"`
}

// MigrationIssue represents a found issue that needs migration
type MigrationIssue struct {
	ID          string `json:"id"`
	PatternID   string `json:"pattern_id"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Column      int    `json:"column"`
	OldCode     string `json:"old_code"`
	NewCode     string `json:"new_code,omitempty"`
	Context     string `json:"context"`
	Severity    string `json:"severity"`
	AutoFixable bool   `json:"auto_fixable"`
	Message     string `json:"message"`
}

// MigrationReport contains analysis results
type MigrationReport struct {
	Timestamp    string             `json:"timestamp"`
	ProjectPath  string             `json:"project_path"`
	FilesScanned int                `json:"files_scanned"`
	Issues       []MigrationIssue   `json:"issues"`
	Summary      map[string]int     `json:"summary"`
	Patterns     []MigrationPattern `json:"patterns"`
}

// MigrationAnalyzer scans codebase for migration issues
type MigrationAnalyzer struct {
	patterns     []MigrationPattern
	fileSet      *token.FileSet
	issues       []MigrationIssue
	filesScanned int
}

// NewMigrationAnalyzer creates a new analyzer with predefined patterns
func NewMigrationAnalyzer() *MigrationAnalyzer {
	return &MigrationAnalyzer{
		patterns:     getDefaultPatterns(),
		fileSet:      token.NewFileSet(),
		issues:       make([]MigrationIssue, 0),
		filesScanned: 0,
	}
}

// getDefaultPatterns returns predefined migration patterns
func getDefaultPatterns() []MigrationPattern {
	return []MigrationPattern{
		// Import patterns
		{
			ID:          "old-button-import",
			Name:        "Old Button Import",
			Description: "Import statement uses old button component path",
			Pattern:     `import.*".*views/components.*button\.templ"`,
			Type:        "import",
			Severity:    "error",
			AutoFix:     true,
			NewPattern:  `"github.com/niiniyare/ruun/views/components/atoms"`,
			Examples:    []string{`import "views/components/button.templ"`},
		},
		{
			ID:          "old-formfield-import",
			Name:        "Old FormField Import",
			Description: "Import statement uses old formfield component path",
			Pattern:     `import.*".*views/components.*formfield\.templ"`,
			Type:        "import",
			Severity:    "error",
			AutoFix:     true,
			NewPattern:  `"github.com/niiniyare/ruun/views/components/molecules"`,
			Examples:    []string{`import "views/components/formfield.templ"`},
		},
		// Component usage patterns
		{
			ID:          "old-button-usage",
			Name:        "Old Button Usage",
			Description: "Button component used without props struct",
			Pattern:     `@Button\([^{].*\)`,
			Type:        "component",
			Severity:    "warning",
			AutoFix:     true,
			NewPattern:  `@atoms.Button(atoms.ButtonProps{Text: "..."}`,
			Examples:    []string{`@Button("Save", "primary", "md")`},
		},
		{
			ID:          "old-formfield-usage",
			Name:        "Old FormField Usage",
			Description: "FormField component used without props struct",
			Pattern:     `@FormField\([^{].*\)`,
			Type:        "component",
			Severity:    "warning",
			AutoFix:     true,
			NewPattern:  `@molecules.FormField(molecules.FormFieldProps{...})`,
			Examples:    []string{`@FormField("name", "text", "Enter name")`},
		},
		// CSS class patterns
		{
			ID:          "hardcoded-colors",
			Name:        "Hardcoded Colors",
			Description: "Hardcoded color classes instead of theme tokens",
			Pattern:     `class="[^"]*(?:text-red-|bg-blue-|border-green-)`,
			Type:        "class",
			Severity:    "warning",
			AutoFix:     false,
			NewPattern:  "Use theme tokens: text-error, bg-primary, border-success",
			Examples:    []string{`class="text-red-500"`, `class="bg-blue-600"`},
		},
		{
			ID:          "inline-styles",
			Name:        "Inline Styles",
			Description: "Inline styles should be converted to theme tokens",
			Pattern:     `style="[^"]*(?:color:|background:|border:)`,
			Type:        "class",
			Severity:    "info",
			AutoFix:     false,
			NewPattern:  "Use CSS classes with theme tokens",
			Examples:    []string{`style="color: #ff0000"`},
		},
		// HTMX patterns
		{
			ID:          "old-htmx-pattern",
			Name:        "Old HTMX Pattern",
			Description: "HTMX attributes applied directly instead of via props",
			Pattern:     `hx-(?:post|get|put|delete)="[^"]*"`,
			Type:        "component",
			Severity:    "warning",
			AutoFix:     true,
			NewPattern:  "Pass HTMX attributes via component props",
			Examples:    []string{`hx-post="/api/save"`},
		},
		// Theme patterns
		{
			ID:          "old-theme-usage",
			Name:        "Old Theme Usage",
			Description: "Direct theme.Get() calls instead of compiled classes",
			Pattern:     `theme\.Get\(`,
			Type:        "theme",
			Severity:    "error",
			AutoFix:     true,
			NewPattern:  "Use compiled theme classes from JSON",
			Examples:    []string{`theme.Get("button.primary.background")`},
		},
		// Alpine.js patterns
		{
			ID:          "old-alpine-pattern",
			Name:        "Old Alpine Pattern",
			Description: "Alpine.js directives applied directly instead of via props",
			Pattern:     `x-(?:data|show|if|for|on)="[^"]*"`,
			Type:        "component",
			Severity:    "warning",
			AutoFix:     true,
			NewPattern:  "Pass Alpine directives via component props",
			Examples:    []string{`x-show="open"`},
		},
		// Validation patterns
		{
			ID:          "old-validation-pattern",
			Name:        "Old Validation Pattern",
			Description: "Manual validation display instead of ValidationState",
			Pattern:     `if.*error.*{`,
			Type:        "component",
			Severity:    "info",
			AutoFix:     false,
			NewPattern:  "Use ValidationState enum and validation messages",
			Examples:    []string{`if field.Error != "" {`},
		},
	}
}

// AnalyzeProject scans the entire project for migration issues
func (ma *MigrationAnalyzer) AnalyzeProject(projectPath string) (*MigrationReport, error) {
	fmt.Printf("Starting migration analysis for: %s\n", projectPath)

	err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip vendor, node_modules, .git directories
		if info.IsDir() {
			name := info.Name()
			if name == "vendor" || name == "node_modules" || name == ".git" ||
				name == "dist" || name == "build" {
				return filepath.SkipDir
			}
			return nil
		}

		// Only analyze relevant files
		if ma.isRelevantFile(path) {
			return ma.analyzeFile(path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking project: %w", err)
	}

	return ma.generateReport(projectPath), nil
}

// isRelevantFile checks if a file should be analyzed
func (ma *MigrationAnalyzer) isRelevantFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".templ" || ext == ".go" || ext == ".html" ||
		ext == ".js" || ext == ".ts" || ext == ".css"
}

// analyzeFile scans a single file for migration issues
func (ma *MigrationAnalyzer) analyzeFile(filePath string) error {
	ma.filesScanned++

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	defer file.Close()

	// Read file content
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}

	// Analyze based on file type
	switch filepath.Ext(filePath) {
	case ".go", ".templ":
		return ma.analyzeGoFile(filePath, string(content))
	case ".html", ".js", ".ts", ".css":
		return ma.analyzeTextFile(filePath, string(content))
	}

	return nil
}

// analyzeGoFile analyzes Go/Templ files using AST when possible
func (ma *MigrationAnalyzer) analyzeGoFile(filePath string, content string) error {
	// Try AST parsing for .go files
	if filepath.Ext(filePath) == ".go" {
		if err := ma.analyzeGoAST(filePath, content); err != nil {
			// Fall back to text analysis if AST parsing fails
			return ma.analyzeTextFile(filePath, content)
		}
		return nil
	}

	// Use text analysis for .templ files
	return ma.analyzeTextFile(filePath, content)
}

// analyzeGoAST analyzes Go files using AST
func (ma *MigrationAnalyzer) analyzeGoAST(filePath string, content string) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, content, parser.ParseComments)
	if err != nil {
		return err
	}

	// Traverse AST to find import statements
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			if x.Path != nil {
				importPath := x.Path.Value
				ma.checkPatterns(filePath, fset.Position(x.Pos()).Line,
					fset.Position(x.Pos()).Column, importPath, "import")
			}
		case *ast.CallExpr:
			// Check for function calls that might be component usage
			if ident, ok := x.Fun.(*ast.Ident); ok {
				callExpr := ident.Name + "("
				ma.checkPatterns(filePath, fset.Position(x.Pos()).Line,
					fset.Position(x.Pos()).Column, callExpr, "component")
			}
		}
		return true
	})

	return nil
}

// analyzeTextFile analyzes files using regex patterns
func (ma *MigrationAnalyzer) analyzeTextFile(filePath string, content string) error {
	lines := strings.Split(content, "\n")

	for lineNum, line := range lines {
		for _, pattern := range ma.patterns {
			regex, err := regexp.Compile(pattern.Pattern)
			if err != nil {
				log.Printf("Invalid regex pattern %s: %v", pattern.ID, err)
				continue
			}

			if matches := regex.FindAllStringSubmatch(line, -1); len(matches) > 0 {
				for _, match := range matches {
					issue := MigrationIssue{
						ID:          fmt.Sprintf("%s-%d-%d", pattern.ID, lineNum+1, 0),
						PatternID:   pattern.ID,
						File:        filePath,
						Line:        lineNum + 1,
						Column:      strings.Index(line, match[0]) + 1,
						OldCode:     match[0],
						NewCode:     pattern.NewPattern,
						Context:     strings.TrimSpace(line),
						Severity:    pattern.Severity,
						AutoFixable: pattern.AutoFix,
						Message:     pattern.Description,
					}
					ma.issues = append(ma.issues, issue)
				}
			}
		}
	}

	return nil
}

// checkPatterns checks a code fragment against all patterns
func (ma *MigrationAnalyzer) checkPatterns(filePath string, line, column int, code, codeType string) {
	for _, pattern := range ma.patterns {
		if pattern.Type != codeType && pattern.Type != "all" {
			continue
		}

		regex, err := regexp.Compile(pattern.Pattern)
		if err != nil {
			continue
		}

		if regex.MatchString(code) {
			issue := MigrationIssue{
				ID:          fmt.Sprintf("%s-%d-%d", pattern.ID, line, column),
				PatternID:   pattern.ID,
				File:        filePath,
				Line:        line,
				Column:      column,
				OldCode:     code,
				NewCode:     pattern.NewPattern,
				Context:     code,
				Severity:    pattern.Severity,
				AutoFixable: pattern.AutoFix,
				Message:     pattern.Description,
			}
			ma.issues = append(ma.issues, issue)
		}
	}
}

// generateReport creates a comprehensive migration report
func (ma *MigrationAnalyzer) generateReport(projectPath string) *MigrationReport {
	summary := make(map[string]int)

	// Count issues by severity
	for _, issue := range ma.issues {
		summary[issue.Severity]++
		summary["total"]++
	}

	// Count issues by type
	for _, pattern := range ma.patterns {
		count := 0
		for _, issue := range ma.issues {
			if issue.PatternID == pattern.ID {
				count++
			}
		}
		if count > 0 {
			summary[pattern.Type] = summary[pattern.Type] + count
		}
	}

	// Count auto-fixable issues
	autoFixable := 0
	for _, issue := range ma.issues {
		if issue.AutoFixable {
			autoFixable++
		}
	}
	summary["auto_fixable"] = autoFixable

	return &MigrationReport{
		Timestamp:    fmt.Sprintf("%d", os.Getpid()), // Simple timestamp
		ProjectPath:  projectPath,
		FilesScanned: ma.filesScanned,
		Issues:       ma.issues,
		Summary:      summary,
		Patterns:     ma.patterns,
	}
}

// SaveReport saves the migration report to a JSON file
func (ma *MigrationAnalyzer) SaveReport(report *MigrationReport, outputPath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling report: %w", err)
	}

	return os.WriteFile(outputPath, data, 0o644)
}

// PrintSummary prints a human-readable summary of the migration report
func (ma *MigrationAnalyzer) PrintSummary(report *MigrationReport) {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("MIGRATION ANALYSIS SUMMARY\n")
	fmt.Printf(strings.Repeat("=", 60) + "\n")

	fmt.Printf("Project Path: %s\n", report.ProjectPath)
	fmt.Printf("Files Scanned: %d\n", report.FilesScanned)
	fmt.Printf("Total Issues Found: %d\n", report.Summary["total"])
	fmt.Printf("Auto-fixable Issues: %d\n", report.Summary["auto_fixable"])
	fmt.Printf("\n")

	// Issues by severity
	fmt.Printf("Issues by Severity:\n")
	fmt.Printf("  - Errors: %d\n", report.Summary["error"])
	fmt.Printf("  - Warnings: %d\n", report.Summary["warning"])
	fmt.Printf("  - Info: %d\n", report.Summary["info"])
	fmt.Printf("\n")

	// Issues by type
	fmt.Printf("Issues by Type:\n")
	fmt.Printf("  - Import Issues: %d\n", report.Summary["import"])
	fmt.Printf("  - Component Issues: %d\n", report.Summary["component"])
	fmt.Printf("  - Class Issues: %d\n", report.Summary["class"])
	fmt.Printf("  - Theme Issues: %d\n", report.Summary["theme"])
	fmt.Printf("\n")

	// Top issues by frequency
	issueCount := make(map[string]int)
	for _, issue := range report.Issues {
		issueCount[issue.PatternID]++
	}

	type issueFreq struct {
		PatternID string
		Count     int
		Pattern   MigrationPattern
	}

	var frequencies []issueFreq
	for patternID, count := range issueCount {
		for _, pattern := range report.Patterns {
			if pattern.ID == patternID {
				frequencies = append(frequencies, issueFreq{
					PatternID: patternID,
					Count:     count,
					Pattern:   pattern,
				})
				break
			}
		}
	}

	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Count > frequencies[j].Count
	})

	fmt.Printf("Most Common Issues:\n")
	for i, freq := range frequencies {
		if i >= 5 { // Show top 5
			break
		}
		fmt.Printf("  %d. %s (%d occurrences)\n", i+1, freq.Pattern.Name, freq.Count)
		fmt.Printf("      %s\n", freq.Pattern.Description)
	}

	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")

	if report.Summary["total"] > 0 {
		fmt.Printf("💡 Run the migration tool with --auto-fix to automatically fix %d issues\n", report.Summary["auto_fixable"])
		fmt.Printf("📝 See the detailed report for manual migration guidance\n")
	} else {
		fmt.Printf("✅ No migration issues found! Your codebase is ready for the new architecture.\n")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run analyzer.go <project-path> [output-file]")
		os.Exit(1)
	}

	projectPath := os.Args[1]
	outputPath := "migration-report.json"
	if len(os.Args) > 2 {
		outputPath = os.Args[2]
	}

	analyzer := NewMigrationAnalyzer()

	report, err := analyzer.AnalyzeProject(projectPath)
	if err != nil {
		log.Fatalf("Error analyzing project: %v", err)
	}

	// Save detailed report
	if err := analyzer.SaveReport(report, outputPath); err != nil {
		log.Fatalf("Error saving report: %v", err)
	}

	// Print summary
	analyzer.PrintSummary(report)

	fmt.Printf("\nDetailed report saved to: %s\n", outputPath)
}

