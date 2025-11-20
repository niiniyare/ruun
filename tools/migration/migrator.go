package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// MigrationRule defines how to transform code
type MigrationRule struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Pattern     string            `json:"pattern"`
	Replacement string            `json:"replacement"`
	FileTypes   []string          `json:"file_types"`
	Conditions  map[string]string `json:"conditions,omitempty"`
	PostActions []string          `json:"post_actions,omitempty"`
}

// MigrationResult tracks the result of a migration
type MigrationResult struct {
	File        string `json:"file"`
	RuleID      string `json:"rule_id"`
	Line        int    `json:"line"`
	OldCode     string `json:"old_code"`
	NewCode     string `json:"new_code"`
	Success     bool   `json:"success"`
	Error       string `json:"error,omitempty"`
}

// MigrationSession represents a migration session
type MigrationSession struct {
	ID            string            `json:"id"`
	Timestamp     time.Time         `json:"timestamp"`
	ProjectPath   string            `json:"project_path"`
	Rules         []MigrationRule   `json:"rules"`
	Results       []MigrationResult `json:"results"`
	Summary       map[string]int    `json:"summary"`
	BackupPath    string            `json:"backup_path"`
	DryRun        bool              `json:"dry_run"`
}

// Migrator handles automated code migration
type Migrator struct {
	session *MigrationSession
	rules   []MigrationRule
}

// NewMigrator creates a new migrator with default rules
func NewMigrator(projectPath string, dryRun bool) *Migrator {
	session := &MigrationSession{
		ID:          fmt.Sprintf("migration-%d", time.Now().Unix()),
		Timestamp:   time.Now(),
		ProjectPath: projectPath,
		Rules:       getDefaultMigrationRules(),
		Results:     make([]MigrationResult, 0),
		Summary:     make(map[string]int),
		DryRun:      dryRun,
	}
	
	return &Migrator{
		session: session,
		rules:   session.Rules,
	}
}

// getDefaultMigrationRules returns predefined migration rules
func getDefaultMigrationRules() []MigrationRule {
	return []MigrationRule{
		// Import migrations
		{
			ID:          "import-button-atoms",
			Name:        "Migrate Button Import to Atoms",
			Description: "Update import paths to use new atoms package",
			Pattern:     `import\s+"[^"]*views/components[^"]*button\.templ"`,
			Replacement: `import "github.com/niiniyare/ruun/views/components/atoms"`,
			FileTypes:   []string{".go", ".templ"},
		},
		{
			ID:          "import-formfield-molecules",
			Name:        "Migrate FormField Import to Molecules",
			Description: "Update import paths to use new molecules package",
			Pattern:     `import\s+"[^"]*views/components[^"]*formfield\.templ"`,
			Replacement: `import "github.com/niiniyare/ruun/views/components/molecules"`,
			FileTypes:   []string{".go", ".templ"},
		},
		{
			ID:          "import-input-atoms",
			Name:        "Migrate Input Import to Atoms",
			Description: "Update import paths to use new atoms package",
			Pattern:     `import\s+"[^"]*views/components[^"]*input\.templ"`,
			Replacement: `import "github.com/niiniyare/ruun/views/components/atoms"`,
			FileTypes:   []string{".go", ".templ"},
		},
		
		// Component usage migrations
		{
			ID:          "button-simple-string-props",
			Name:        "Migrate Simple Button Calls",
			Description: "Convert simple button calls to props struct",
			Pattern:     `@Button\("([^"]+)"\)`,
			Replacement: `@atoms.Button(atoms.ButtonProps{Text: "$1"})`,
			FileTypes:   []string{".templ"},
		},
		{
			ID:          "button-text-variant-props",
			Name:        "Migrate Button with Variant",
			Description: "Convert button calls with text and variant to props struct",
			Pattern:     `@Button\("([^"]+)",\s*"([^"]+)"\)`,
			Replacement: `@atoms.Button(atoms.ButtonProps{Text: "$1", Variant: atoms.Button$2})`,
			FileTypes:   []string{".templ"},
			PostActions: []string{"capitalize-variant"},
		},
		{
			ID:          "formfield-basic-props",
			Name:        "Migrate Basic FormField Calls",
			Description: "Convert basic formfield calls to props struct",
			Pattern:     `@FormField\("([^"]+)",\s*"([^"]+)",\s*"([^"]*)"\)`,
			Replacement: `@molecules.FormField(molecules.FormFieldProps{Name: "$1", Type: "$2", Placeholder: "$3"})`,
			FileTypes:   []string{".templ"},
		},
		{
			ID:          "input-basic-props",
			Name:        "Migrate Basic Input Calls",
			Description: "Convert basic input calls to props struct",
			Pattern:     `@Input\("([^"]+)",\s*"([^"]+)"\)`,
			Replacement: `@atoms.Input(atoms.InputProps{Name: "$1", Placeholder: "$2"})`,
			FileTypes:   []string{".templ"},
		},
		
		// CSS class migrations
		{
			ID:          "hardcoded-red-colors",
			Name:        "Replace Hardcoded Red Colors",
			Description: "Replace hardcoded red colors with theme tokens",
			Pattern:     `(text|bg|border)-red-\d+`,
			Replacement: `$1-error`,
			FileTypes:   []string{".templ", ".html", ".go"},
		},
		{
			ID:          "hardcoded-blue-colors",
			Name:        "Replace Hardcoded Blue Colors", 
			Description: "Replace hardcoded blue colors with theme tokens",
			Pattern:     `(text|bg|border)-blue-\d+`,
			Replacement: `$1-primary`,
			FileTypes:   []string{".templ", ".html", ".go"},
		},
		{
			ID:          "hardcoded-green-colors",
			Name:        "Replace Hardcoded Green Colors",
			Description: "Replace hardcoded green colors with theme tokens",
			Pattern:     `(text|bg|border)-green-\d+`,
			Replacement: `$1-success`,
			FileTypes:   []string{".templ", ".html", ".go"},
		},
		
		// HTMX attribute migrations
		{
			ID:          "htmx-to-props",
			Name:        "Migrate HTMX Attributes to Props",
			Description: "Move HTMX attributes from HTML to component props",
			Pattern:     `hx-(post|get|put|delete)="([^"]+)"`,
			Replacement: `HX$1: "$2"`,
			FileTypes:   []string{".templ"},
			PostActions: []string{"capitalize-htmx-method"},
		},
		
		// Alpine.js migrations
		{
			ID:          "alpine-show-to-props",
			Name:        "Migrate Alpine x-show to Props",
			Description: "Move Alpine x-show to component props",
			Pattern:     `x-show="([^"]+)"`,
			Replacement: `AlpineShow: "$1"`,
			FileTypes:   []string{".templ"},
		},
		{
			ID:          "alpine-click-to-props",
			Name:        "Migrate Alpine x-on:click to Props",
			Description: "Move Alpine x-on:click to component props", 
			Pattern:     `x-on:click="([^"]+)"`,
			Replacement: `AlpineClick: "$1"`,
			FileTypes:   []string{".templ"},
		},
		
		// Theme migrations
		{
			ID:          "theme-get-to-classes",
			Name:        "Replace theme.Get() with Compiled Classes",
			Description: "Replace dynamic theme calls with compiled CSS classes",
			Pattern:     `theme\.Get\("([^"]+)"\)`,
			Replacement: `"$1"`, // Simplified - would need more complex logic
			FileTypes:   []string{".go", ".templ"},
		},
	}
}

// Run executes the migration
func (m *Migrator) Run() error {
	fmt.Printf("Starting migration for: %s\n", m.session.ProjectPath)
	fmt.Printf("Dry run: %v\n", m.session.DryRun)
	
	// Create backup if not dry run
	if !m.session.DryRun {
		backupPath, err := m.createBackup()
		if err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}
		m.session.BackupPath = backupPath
		fmt.Printf("Backup created at: %s\n", backupPath)
	}
	
	// Walk through all files
	err := filepath.Walk(m.session.ProjectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories and irrelevant files
		if info.IsDir() || !m.isRelevantFile(path) {
			return nil
		}
		
		return m.migrateFile(path)
	})
	
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	
	m.generateSummary()
	return nil
}

// createBackup creates a backup of the project
func (m *Migrator) createBackup() (string, error) {
	timestamp := time.Now().Format("20060102_150405")
	backupDir := filepath.Join(m.session.ProjectPath, "migration_backups")
	backupPath := filepath.Join(backupDir, fmt.Sprintf("backup_%s", timestamp))
	
	// Create backup directory
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return "", err
	}
	
	// Copy relevant files
	return backupPath, filepath.Walk(m.session.ProjectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		
		// Skip backup directory itself
		if strings.Contains(path, "migration_backups") {
			return nil
		}
		
		if m.isRelevantFile(path) {
			relPath, err := filepath.Rel(m.session.ProjectPath, path)
			if err != nil {
				return err
			}
			
			destPath := filepath.Join(backupPath, relPath)
			destDir := filepath.Dir(destPath)
			
			if err := os.MkdirAll(destDir, 0755); err != nil {
				return err
			}
			
			return m.copyFile(path, destPath)
		}
		
		return nil
	})
}

// copyFile copies a file from src to dst
func (m *Migrator) copyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	
	return os.WriteFile(dst, input, 0644)
}

// isRelevantFile checks if a file should be processed
func (m *Migrator) isRelevantFile(path string) bool {
	ext := filepath.Ext(path)
	relevantExts := []string{".go", ".templ", ".html", ".js", ".ts", ".css"}
	
	for _, relevantExt := range relevantExts {
		if ext == relevantExt {
			return true
		}
	}
	
	return false
}

// migrateFile applies migration rules to a single file
func (m *Migrator) migrateFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	
	originalContent := string(content)
	modifiedContent := originalContent
	fileModified := false
	
	ext := filepath.Ext(filePath)
	
	// Apply each rule that matches the file type
	for _, rule := range m.rules {
		if !m.ruleAppliesTo(rule, ext) {
			continue
		}
		
		regex, err := regexp.Compile(rule.Pattern)
		if err != nil {
			fmt.Printf("Warning: Invalid regex in rule %s: %v\n", rule.ID, err)
			continue
		}
		
		if regex.MatchString(modifiedContent) {
			newContent := regex.ReplaceAllString(modifiedContent, rule.Replacement)
			
			if newContent != modifiedContent {
				// Apply post-actions
				newContent = m.applyPostActions(newContent, rule.PostActions)
				
				// Record the change
				result := MigrationResult{
					File:    filePath,
					RuleID:  rule.ID,
					OldCode: m.extractMatch(modifiedContent, regex),
					NewCode: m.extractMatch(newContent, regex),
					Success: true,
				}
				
				m.session.Results = append(m.session.Results, result)
				modifiedContent = newContent
				fileModified = true
				
				fmt.Printf("Applied rule '%s' to %s\n", rule.Name, filePath)
			}
		}
	}
	
	// Write the modified content if changes were made and not dry run
	if fileModified && !m.session.DryRun {
		if err := os.WriteFile(filePath, []byte(modifiedContent), 0644); err != nil {
			return fmt.Errorf("error writing file %s: %w", filePath, err)
		}
	}
	
	return nil
}

// ruleAppliesTo checks if a rule applies to a file type
func (m *Migrator) ruleAppliesTo(rule MigrationRule, ext string) bool {
	for _, fileType := range rule.FileTypes {
		if fileType == ext {
			return true
		}
	}
	return false
}

// extractMatch extracts the first match from content for logging
func (m *Migrator) extractMatch(content string, regex *regexp.Regexp) string {
	matches := regex.FindStringSubmatch(content)
	if len(matches) > 0 {
		return matches[0]
	}
	return ""
}

// applyPostActions applies post-processing actions to content
func (m *Migrator) applyPostActions(content string, actions []string) string {
	for _, action := range actions {
		switch action {
		case "capitalize-variant":
			content = m.capitalizeVariants(content)
		case "capitalize-htmx-method":
			content = m.capitalizeHTMXMethods(content)
		}
	}
	return content
}

// capitalizeVariants capitalizes button variant names
func (m *Migrator) capitalizeVariants(content string) string {
	variants := map[string]string{
		"primary":     "Primary",
		"secondary":   "Secondary", 
		"destructive": "Destructive",
		"outline":     "Outline",
		"ghost":       "Ghost",
		"link":        "Link",
	}
	
	for old, new := range variants {
		content = strings.ReplaceAll(content, "Button"+old, "Button"+new)
	}
	
	return content
}

// capitalizeHTMXMethods capitalizes HTMX method names
func (m *Migrator) capitalizeHTMXMethods(content string) string {
	methods := map[string]string{
		"HXpost":   "HXPost",
		"HXget":    "HXGet", 
		"HXput":    "HXPut",
		"HXdelete": "HXDelete",
	}
	
	for old, new := range methods {
		content = strings.ReplaceAll(content, old, new)
	}
	
	return content
}

// generateSummary generates migration summary
func (m *Migrator) generateSummary() {
	summary := make(map[string]int)
	
	// Count results by rule
	ruleCount := make(map[string]int)
	for _, result := range m.session.Results {
		ruleCount[result.RuleID]++
		if result.Success {
			summary["successful"]++
		} else {
			summary["failed"]++
		}
	}
	
	summary["total_changes"] = len(m.session.Results)
	summary["rules_applied"] = len(ruleCount)
	
	// Count files modified
	fileSet := make(map[string]bool)
	for _, result := range m.session.Results {
		fileSet[result.File] = true
	}
	summary["files_modified"] = len(fileSet)
	
	m.session.Summary = summary
}

// SaveSession saves the migration session to a JSON file
func (m *Migrator) SaveSession(outputPath string) error {
	data, err := json.MarshalIndent(m.session, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling session: %w", err)
	}
	
	return os.WriteFile(outputPath, data, 0644)
}

// PrintSummary prints a human-readable summary
func (m *Migrator) PrintSummary() {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("MIGRATION SUMMARY\n")
	fmt.Printf(strings.Repeat("=", 60) + "\n")
	
	fmt.Printf("Session ID: %s\n", m.session.ID)
	fmt.Printf("Project Path: %s\n", m.session.ProjectPath)
	fmt.Printf("Dry Run: %v\n", m.session.DryRun)
	if m.session.BackupPath != "" {
		fmt.Printf("Backup Location: %s\n", m.session.BackupPath)
	}
	fmt.Printf("\n")
	
	fmt.Printf("Migration Results:\n")
	fmt.Printf("  - Total Changes: %d\n", m.session.Summary["total_changes"])
	fmt.Printf("  - Files Modified: %d\n", m.session.Summary["files_modified"])
	fmt.Printf("  - Rules Applied: %d\n", m.session.Summary["rules_applied"])
	fmt.Printf("  - Successful: %d\n", m.session.Summary["successful"])
	fmt.Printf("  - Failed: %d\n", m.session.Summary["failed"])
	fmt.Printf("\n")
	
	// Show changes by rule
	ruleCount := make(map[string]int)
	for _, result := range m.session.Results {
		ruleCount[result.RuleID]++
	}
	
	fmt.Printf("Changes by Rule:\n")
	for ruleID, count := range ruleCount {
		for _, rule := range m.rules {
			if rule.ID == ruleID {
				fmt.Printf("  - %s: %d changes\n", rule.Name, count)
				break
			}
		}
	}
	
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	
	if m.session.DryRun {
		fmt.Printf("💡 This was a dry run. Re-run without --dry-run to apply changes.\n")
	} else if m.session.Summary["total_changes"] > 0 {
		fmt.Printf("✅ Migration completed successfully!\n")
		fmt.Printf("🔄 Run the validation tool to verify the changes.\n")
		if m.session.BackupPath != "" {
			fmt.Printf("🗂️  Original files backed up to: %s\n", m.session.BackupPath)
		}
	} else {
		fmt.Printf("ℹ️  No changes were needed - your code is already up to date!\n")
	}
}

// LoadRulesFromFile loads migration rules from a JSON file
func (m *Migrator) LoadRulesFromFile(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading rules file: %w", err)
	}
	
	var rules []MigrationRule
	if err := json.Unmarshal(data, &rules); err != nil {
		return fmt.Errorf("error parsing rules file: %w", err)
	}
	
	m.rules = rules
	m.session.Rules = rules
	
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run migrator.go <project-path> [--dry-run] [--rules=rules.json] [--output=session.json]")
		fmt.Println("\nOptions:")
		fmt.Println("  --dry-run: Preview changes without modifying files")
		fmt.Println("  --rules=FILE: Load custom migration rules from JSON file")
		fmt.Println("  --output=FILE: Save migration session to JSON file")
		os.Exit(1)
	}
	
	projectPath := os.Args[1]
	dryRun := false
	rulesFile := ""
	outputFile := "migration-session.json"
	
	// Parse command line arguments
	for _, arg := range os.Args[2:] {
		if arg == "--dry-run" {
			dryRun = true
		} else if strings.HasPrefix(arg, "--rules=") {
			rulesFile = strings.TrimPrefix(arg, "--rules=")
		} else if strings.HasPrefix(arg, "--output=") {
			outputFile = strings.TrimPrefix(arg, "--output=")
		}
	}
	
	migrator := NewMigrator(projectPath, dryRun)
	
	// Load custom rules if specified
	if rulesFile != "" {
		if err := migrator.LoadRulesFromFile(rulesFile); err != nil {
			fmt.Printf("Error loading rules file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loaded custom rules from: %s\n", rulesFile)
	}
	
	// Run migration
	if err := migrator.Run(); err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		os.Exit(1)
	}
	
	// Save session
	if err := migrator.SaveSession(outputFile); err != nil {
		fmt.Printf("Error saving session: %v\n", err)
		os.Exit(1)
	}
	
	// Print summary
	migrator.PrintSummary()
	
	fmt.Printf("\nMigration session saved to: %s\n", outputFile)
}