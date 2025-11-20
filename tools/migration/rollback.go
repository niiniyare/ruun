package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// RollbackSession represents a rollback operation
type RollbackSession struct {
	ID              string                 `json:"id"`
	Timestamp       time.Time              `json:"timestamp"`
	ProjectPath     string                 `json:"project_path"`
	MigrationID     string                 `json:"migration_id"`
	BackupPath      string                 `json:"backup_path"`
	RollbackType    string                 `json:"rollback_type"` // "full", "partial", "selective"
	FilesRestored   []string               `json:"files_restored"`
	FilesSkipped    []string               `json:"files_skipped"`
	Summary         map[string]int         `json:"summary"`
	DryRun          bool                   `json:"dry_run"`
	Error           string                 `json:"error,omitempty"`
	CustomRollbacks []CustomRollbackRule   `json:"custom_rollbacks,omitempty"`
}

// CustomRollbackRule defines custom rollback logic
type CustomRollbackRule struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	FilePattern string `json:"file_pattern"`
	Action      string `json:"action"` // "restore", "skip", "transform"
	Condition   string `json:"condition,omitempty"`
}

// FileRestoreResult tracks file restoration results
type FileRestoreResult struct {
	SourceFile      string `json:"source_file"`
	DestinationFile string `json:"destination_file"`
	Success         bool   `json:"success"`
	Error           string `json:"error,omitempty"`
	Skipped         bool   `json:"skipped"`
	SkipReason      string `json:"skip_reason,omitempty"`
}

// MigrationRollback handles rollback operations
type MigrationRollback struct {
	session *RollbackSession
}

// NewMigrationRollback creates a new rollback handler
func NewMigrationRollback(projectPath, migrationID, backupPath string, dryRun bool) *MigrationRollback {
	session := &RollbackSession{
		ID:            fmt.Sprintf("rollback-%d", time.Now().Unix()),
		Timestamp:     time.Now(),
		ProjectPath:   projectPath,
		MigrationID:   migrationID,
		BackupPath:    backupPath,
		RollbackType:  "full",
		FilesRestored: make([]string, 0),
		FilesSkipped:  make([]string, 0),
		Summary:       make(map[string]int),
		DryRun:        dryRun,
	}
	
	return &MigrationRollback{
		session: session,
	}
}

// LoadMigrationSession loads a migration session to understand what was changed
func (r *MigrationRollback) LoadMigrationSession(sessionFile string) error {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return fmt.Errorf("error reading migration session: %w", err)
	}
	
	var migrationSession map[string]interface{}
	if err := json.Unmarshal(data, &migrationSession); err != nil {
		return fmt.Errorf("error parsing migration session: %w", err)
	}
	
	// Extract information we need for rollback
	if id, ok := migrationSession["id"].(string); ok {
		r.session.MigrationID = id
	}
	
	if backupPath, ok := migrationSession["backup_path"].(string); ok {
		r.session.BackupPath = backupPath
	}
	
	return nil
}

// ExecuteFullRollback restores all files from backup
func (r *MigrationRollback) ExecuteFullRollback() error {
	fmt.Printf("Starting full rollback for migration: %s\n", r.session.MigrationID)
	fmt.Printf("Dry run: %v\n", r.session.DryRun)
	
	if r.session.BackupPath == "" {
		return fmt.Errorf("no backup path specified")
	}
	
	if _, err := os.Stat(r.session.BackupPath); os.IsNotExist(err) {
		return fmt.Errorf("backup directory not found: %s", r.session.BackupPath)
	}
	
	// Walk through backup directory and restore files
	err := filepath.Walk(r.session.BackupPath, func(backupFile string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		// Calculate destination path
		relPath, err := filepath.Rel(r.session.BackupPath, backupFile)
		if err != nil {
			return fmt.Errorf("error calculating relative path: %w", err)
		}
		
		destPath := filepath.Join(r.session.ProjectPath, relPath)
		
		return r.restoreFile(backupFile, destPath)
	})
	
	if err != nil {
		r.session.Error = err.Error()
		return fmt.Errorf("rollback failed: %w", err)
	}
	
	r.generateSummary()
	return nil
}

// ExecuteSelectiveRollback restores only specific files or patterns
func (r *MigrationRollback) ExecuteSelectiveRollback(filePatterns []string) error {
	fmt.Printf("Starting selective rollback for patterns: %v\n", filePatterns)
	
	if r.session.BackupPath == "" {
		return fmt.Errorf("no backup path specified")
	}
	
	// Walk through backup and restore matching files
	err := filepath.Walk(r.session.BackupPath, func(backupFile string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		
		relPath, err := filepath.Rel(r.session.BackupPath, backupFile)
		if err != nil {
			return err
		}
		
		// Check if file matches any pattern
		if r.matchesPatterns(relPath, filePatterns) {
			destPath := filepath.Join(r.session.ProjectPath, relPath)
			return r.restoreFile(backupFile, destPath)
		}
		
		return nil
	})
	
	if err != nil {
		r.session.Error = err.Error()
		return fmt.Errorf("selective rollback failed: %w", err)
	}
	
	r.session.RollbackType = "selective"
	r.generateSummary()
	return nil
}

// ExecutePartialRollback restores files but applies transformation rules
func (r *MigrationRollback) ExecutePartialRollback(rules []CustomRollbackRule) error {
	fmt.Printf("Starting partial rollback with %d custom rules\n", len(rules))
	
	r.session.CustomRollbacks = rules
	
	err := filepath.Walk(r.session.BackupPath, func(backupFile string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		
		relPath, err := filepath.Rel(r.session.BackupPath, backupFile)
		if err != nil {
			return err
		}
		
		destPath := filepath.Join(r.session.ProjectPath, relPath)
		
		// Apply custom rules
		action := r.getActionForFile(relPath, rules)
		
		switch action {
		case "restore":
			return r.restoreFile(backupFile, destPath)
		case "skip":
			r.skipFile(relPath, "custom rule")
			return nil
		case "transform":
			return r.transformAndRestoreFile(backupFile, destPath)
		default:
			return r.restoreFile(backupFile, destPath)
		}
	})
	
	if err != nil {
		r.session.Error = err.Error()
		return fmt.Errorf("partial rollback failed: %w", err)
	}
	
	r.session.RollbackType = "partial"
	r.generateSummary()
	return nil
}

// restoreFile restores a single file from backup
func (r *MigrationRollback) restoreFile(sourcePath, destPath string) error {
	fmt.Printf("Restoring: %s -> %s\n", sourcePath, destPath)
	
	if r.session.DryRun {
		r.session.FilesRestored = append(r.session.FilesRestored, destPath)
		return nil
	}
	
	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("error creating directory %s: %w", destDir, err)
	}
	
	// Copy file content
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}
	
	if err := os.WriteFile(destPath, content, 0644); err != nil {
		return fmt.Errorf("error writing destination file: %w", err)
	}
	
	r.session.FilesRestored = append(r.session.FilesRestored, destPath)
	return nil
}

// skipFile records that a file was skipped
func (r *MigrationRollback) skipFile(filePath, reason string) {
	fmt.Printf("Skipping: %s (%s)\n", filePath, reason)
	r.session.FilesSkipped = append(r.session.FilesSkipped, filePath)
}

// transformAndRestoreFile applies transformations before restoring
func (r *MigrationRollback) transformAndRestoreFile(sourcePath, destPath string) error {
	fmt.Printf("Transforming and restoring: %s -> %s\n", sourcePath, destPath)
	
	content, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("error reading source file: %w", err)
	}
	
	// Apply transformations (this is a simplified version)
	transformedContent := r.applyRollbackTransformations(string(content), destPath)
	
	if r.session.DryRun {
		r.session.FilesRestored = append(r.session.FilesRestored, destPath)
		return nil
	}
	
	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("error creating directory %s: %w", destDir, err)
	}
	
	if err := os.WriteFile(destPath, []byte(transformedContent), 0644); err != nil {
		return fmt.Errorf("error writing transformed file: %w", err)
	}
	
	r.session.FilesRestored = append(r.session.FilesRestored, destPath)
	return nil
}

// applyRollbackTransformations applies custom transformations during rollback
func (r *MigrationRollback) applyRollbackTransformations(content, filePath string) string {
	// Example transformations - keep some new patterns while rolling back others
	
	// Keep new import structure but restore old component usage
	if strings.HasSuffix(filePath, ".templ") {
		// Keep atoms/molecules imports
		if strings.Contains(content, `"github.com/niiniyare/ruun/views/components/atoms"`) {
			// But restore simpler component calls if they existed
			content = strings.ReplaceAll(content, `@atoms.Button(atoms.ButtonProps{Text: "`, `@Button("`)
			content = strings.ReplaceAll(content, `"})`, `")`)
		}
	}
	
	return content
}

// matchesPatterns checks if a file path matches any of the given patterns
func (r *MigrationRollback) matchesPatterns(filePath string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, filePath); matched {
			return true
		}
		
		// Also check if pattern is a substring of the path
		if strings.Contains(filePath, pattern) {
			return true
		}
	}
	return false
}

// getActionForFile determines what action to take for a file based on rules
func (r *MigrationRollback) getActionForFile(filePath string, rules []CustomRollbackRule) string {
	for _, rule := range rules {
		if r.matchesPatterns(filePath, []string{rule.FilePattern}) {
			return rule.Action
		}
	}
	return "restore" // default action
}

// generateSummary generates rollback summary
func (r *MigrationRollback) generateSummary() {
	r.session.Summary["files_restored"] = len(r.session.FilesRestored)
	r.session.Summary["files_skipped"] = len(r.session.FilesSkipped)
	r.session.Summary["total_processed"] = len(r.session.FilesRestored) + len(r.session.FilesSkipped)
	
	if r.session.Error != "" {
		r.session.Summary["errors"] = 1
	} else {
		r.session.Summary["errors"] = 0
	}
}

// SaveSession saves the rollback session
func (r *MigrationRollback) SaveSession(outputPath string) error {
	data, err := json.MarshalIndent(r.session, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling session: %w", err)
	}
	
	return os.WriteFile(outputPath, data, 0644)
}

// PrintSummary prints a human-readable summary
func (r *MigrationRollback) PrintSummary() {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("ROLLBACK SUMMARY\n")
	fmt.Printf(strings.Repeat("=", 60) + "\n")
	
	fmt.Printf("Rollback ID: %s\n", r.session.ID)
	fmt.Printf("Migration ID: %s\n", r.session.MigrationID)
	fmt.Printf("Project Path: %s\n", r.session.ProjectPath)
	fmt.Printf("Backup Path: %s\n", r.session.BackupPath)
	fmt.Printf("Rollback Type: %s\n", r.session.RollbackType)
	fmt.Printf("Dry Run: %v\n", r.session.DryRun)
	fmt.Printf("\n")
	
	fmt.Printf("Rollback Results:\n")
	fmt.Printf("  - Files Restored: %d\n", r.session.Summary["files_restored"])
	fmt.Printf("  - Files Skipped: %d\n", r.session.Summary["files_skipped"])
	fmt.Printf("  - Total Processed: %d\n", r.session.Summary["total_processed"])
	fmt.Printf("  - Errors: %d\n", r.session.Summary["errors"])
	
	if r.session.Error != "" {
		fmt.Printf("  - Error Details: %s\n", r.session.Error)
	}
	
	fmt.Printf("\n")
	
	// Show restored files (first 10)
	if len(r.session.FilesRestored) > 0 {
		fmt.Printf("Restored Files (showing first 10):\n")
		for i, file := range r.session.FilesRestored {
			if i >= 10 {
				fmt.Printf("  ... and %d more\n", len(r.session.FilesRestored)-10)
				break
			}
			fmt.Printf("  - %s\n", file)
		}
	}
	
	// Show skipped files
	if len(r.session.FilesSkipped) > 0 {
		fmt.Printf("\nSkipped Files:\n")
		for _, file := range r.session.FilesSkipped {
			fmt.Printf("  - %s\n", file)
		}
	}
	
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	
	if r.session.Summary["errors"] > 0 {
		fmt.Printf("❌ Rollback completed with errors. Check the details above.\n")
	} else if r.session.DryRun {
		fmt.Printf("💡 This was a dry run. Re-run without --dry-run to apply rollback.\n")
	} else {
		fmt.Printf("✅ Rollback completed successfully!\n")
		fmt.Printf("🔧 Consider running tests to verify the rollback.\n")
	}
}

// ListBackups lists available migration backups
func ListBackups(projectPath string) error {
	backupDir := filepath.Join(projectPath, "migration_backups")
	
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		fmt.Printf("No backup directory found at: %s\n", backupDir)
		return nil
	}
	
	entries, err := os.ReadDir(backupDir)
	if err != nil {
		return fmt.Errorf("error reading backup directory: %w", err)
	}
	
	fmt.Printf("Available Backups in %s:\n", backupDir)
	fmt.Printf(strings.Repeat("-", 50) + "\n")
	
	if len(entries) == 0 {
		fmt.Printf("No backups found.\n")
		return nil
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			
			backupPath := filepath.Join(backupDir, entry.Name())
			
			// Count files in backup
			fileCount := 0
			filepath.Walk(backupPath, func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					fileCount++
				}
				return nil
			})
			
			fmt.Printf("📁 %s\n", entry.Name())
			fmt.Printf("   Created: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
			fmt.Printf("   Files: %d\n", fileCount)
			fmt.Printf("   Path: %s\n", backupPath)
			fmt.Printf("\n")
		}
	}
	
	return nil
}

// getDefaultCustomRules returns default custom rollback rules
func getDefaultCustomRules() []CustomRollbackRule {
	return []CustomRollbackRule{
		{
			ID:          "keep-new-imports",
			Name:        "Keep New Import Structure",
			Description: "Keep new atoms/molecules imports but restore old usage patterns",
			FilePattern: "*.templ",
			Action:      "transform",
		},
		{
			ID:          "skip-theme-files",
			Name:        "Skip Theme Files",
			Description: "Don't rollback theme-related changes",
			FilePattern: "*theme*",
			Action:      "skip",
		},
		{
			ID:          "restore-components",
			Name:        "Restore Component Files",
			Description: "Fully restore component files",
			FilePattern: "views/components/*",
			Action:      "restore",
		},
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run rollback.go <command> [options]")
		fmt.Println("\nCommands:")
		fmt.Println("  list <project-path>                    - List available backups")
		fmt.Println("  full <project-path> <backup-path>      - Full rollback from backup")
		fmt.Println("  selective <project-path> <backup-path> <pattern>... - Selective rollback")
		fmt.Println("  partial <project-path> <backup-path>   - Partial rollback with rules")
		fmt.Println("\nOptions:")
		fmt.Println("  --dry-run           Preview rollback without changes")
		fmt.Println("  --migration=ID      Migration session ID")
		fmt.Println("  --output=FILE       Save rollback session to file")
		os.Exit(1)
	}
	
	command := os.Args[1]
	
	switch command {
	case "list":
		if len(os.Args) < 3 {
			fmt.Println("Usage: rollback.go list <project-path>")
			os.Exit(1)
		}
		projectPath := os.Args[2]
		if err := ListBackups(projectPath); err != nil {
			fmt.Printf("Error listing backups: %v\n", err)
			os.Exit(1)
		}
		
	case "full", "selective", "partial":
		if len(os.Args) < 4 {
			fmt.Printf("Usage: rollback.go %s <project-path> <backup-path> [options]\n", command)
			os.Exit(1)
		}
		
		projectPath := os.Args[2]
		backupPath := os.Args[3]
		
		// Parse options
		dryRun := false
		migrationID := ""
		outputFile := "rollback-session.json"
		var patterns []string
		
		for _, arg := range os.Args[4:] {
			if arg == "--dry-run" {
				dryRun = true
			} else if strings.HasPrefix(arg, "--migration=") {
				migrationID = strings.TrimPrefix(arg, "--migration=")
			} else if strings.HasPrefix(arg, "--output=") {
				outputFile = strings.TrimPrefix(arg, "--output=")
			} else if command == "selective" {
				patterns = append(patterns, arg)
			}
		}
		
		rollback := NewMigrationRollback(projectPath, migrationID, backupPath, dryRun)
		
		var err error
		switch command {
		case "full":
			err = rollback.ExecuteFullRollback()
		case "selective":
			if len(patterns) == 0 {
				fmt.Println("Error: selective rollback requires file patterns")
				os.Exit(1)
			}
			err = rollback.ExecuteSelectiveRollback(patterns)
		case "partial":
			rules := getDefaultCustomRules()
			err = rollback.ExecutePartialRollback(rules)
		}
		
		if err != nil {
			fmt.Printf("Rollback failed: %v\n", err)
			os.Exit(1)
		}
		
		// Save session
		if err := rollback.SaveSession(outputFile); err != nil {
			fmt.Printf("Error saving session: %v\n", err)
		}
		
		// Print summary
		rollback.PrintSummary()
		
		fmt.Printf("\nRollback session saved to: %s\n", outputFile)
		
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}