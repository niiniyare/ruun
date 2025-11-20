package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// CLIConfig holds configuration for the migration CLI
type CLIConfig struct {
	ProjectPath  string
	Verbose      bool
	DryRun       bool
	AutoFix      bool
	Interactive  bool
	BackupDir    string
	OutputDir    string
	ConfigFile   string
	LogLevel     string
}

// MigrationPipeline orchestrates the complete migration process
type MigrationPipeline struct {
	config      *CLIConfig
	sessionID   string
	toolsDir    string
	workingDir  string
	logFile     *os.File
}

// NewMigrationPipeline creates a new migration pipeline
func NewMigrationPipeline(config *CLIConfig) (*MigrationPipeline, error) {
	sessionID := fmt.Sprintf("migration-%d", time.Now().Unix())
	
	// Setup working directory
	workingDir := filepath.Join(config.OutputDir, sessionID)
	if err := os.MkdirAll(workingDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create working directory: %w", err)
	}
	
	// Setup logging
	logPath := filepath.Join(workingDir, "migration.log")
	logFile, err := os.Create(logPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create log file: %w", err)
	}
	
	return &MigrationPipeline{
		config:     config,
		sessionID:  sessionID,
		toolsDir:   ".",
		workingDir: workingDir,
		logFile:    logFile,
	}, nil
}

// Execute runs the complete migration pipeline
func (mp *MigrationPipeline) Execute() error {
	defer mp.logFile.Close()
	
	mp.log("info", "Starting migration pipeline")
	mp.log("info", fmt.Sprintf("Session ID: %s", mp.sessionID))
	mp.log("info", fmt.Sprintf("Project Path: %s", mp.config.ProjectPath))
	
	// Phase 1: Analysis
	fmt.Println("📊 Phase 1: Analyzing codebase...")
	analysisReport, err := mp.runAnalysis()
	if err != nil {
		return fmt.Errorf("analysis phase failed: %w", err)
	}
	
	mp.printPhaseResult("Analysis", analysisReport.Summary["total"], "issues found")
	
	if analysisReport.Summary["total"] == 0 {
		fmt.Println("✅ No migration issues found. Your codebase is already up to date!")
		return nil
	}
	
	// Phase 2: User confirmation
	if mp.config.Interactive && !mp.confirmProceed(analysisReport) {
		fmt.Println("Migration cancelled by user")
		return nil
	}
	
	// Phase 3: Backup
	fmt.Println("💾 Phase 2: Creating backup...")
	backupPath, err := mp.createBackup()
	if err != nil {
		return fmt.Errorf("backup phase failed: %w", err)
	}
	
	mp.printPhaseResult("Backup", 1, fmt.Sprintf("created at %s", backupPath))
	
	// Phase 4: Migration
	fmt.Println("🔄 Phase 3: Executing migration...")
	migrationSession, err := mp.runMigration()
	if err != nil {
		return fmt.Errorf("migration phase failed: %w", err)
	}
	
	mp.printPhaseResult("Migration", migrationSession.Summary["total_changes"], "changes applied")
	
	// Phase 5: Validation
	fmt.Println("✅ Phase 4: Validating results...")
	validationReport, err := mp.runValidation(migrationSession.ID)
	if err != nil {
		return fmt.Errorf("validation phase failed: %w", err)
	}
	
	mp.printPhaseResult("Validation", validationReport.Summary["total"], "issues remaining")
	
	// Phase 6: Testing (if configured)
	if mp.shouldRunTests() {
		fmt.Println("🧪 Phase 5: Running tests...")
		testReport, err := mp.runTests()
		if err != nil {
			mp.log("warning", fmt.Sprintf("Testing phase had issues: %v", err))
			fmt.Printf("⚠️  Testing completed with warnings (see log for details)\n")
		} else {
			mp.printPhaseResult("Testing", testReport.Passed, fmt.Sprintf("of %d tests passed", testReport.TotalTests))
		}
	}
	
	// Phase 7: Summary and recommendations
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("🎉 MIGRATION COMPLETE")
	fmt.Println(strings.Repeat("=", 60))
	
	mp.printFinalSummary(analysisReport, migrationSession, validationReport)
	
	return nil
}

// runAnalysis executes the analysis tool
func (mp *MigrationPipeline) runAnalysis() (*AnalysisReport, error) {
	mp.log("info", "Running migration analysis")
	
	outputPath := filepath.Join(mp.workingDir, "analysis-report.json")
	
	cmd := exec.Command("go", "run", "analyzer.go", mp.config.ProjectPath, outputPath)
	cmd.Dir = mp.toolsDir
	
	output, err := cmd.CombinedOutput()
	mp.log("debug", fmt.Sprintf("Analyzer output: %s", string(output)))
	
	if err != nil {
		return nil, fmt.Errorf("analyzer failed: %w", err)
	}
	
	// Load the analysis report
	data, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read analysis report: %w", err)
	}
	
	var report AnalysisReport
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("failed to parse analysis report: %w", err)
	}
	
	return &report, nil
}

// runMigration executes the migration tool
func (mp *MigrationPipeline) runMigration() (*MigrationSessionData, error) {
	mp.log("info", "Running automated migration")
	
	outputPath := filepath.Join(mp.workingDir, "migration-session.json")
	
	args := []string{"run", "migrator.go", mp.config.ProjectPath}
	if mp.config.DryRun {
		args = append(args, "--dry-run")
	}
	args = append(args, fmt.Sprintf("--output=%s", outputPath))
	
	cmd := exec.Command("go", args...)
	cmd.Dir = mp.toolsDir
	
	output, err := cmd.CombinedOutput()
	mp.log("debug", fmt.Sprintf("Migrator output: %s", string(output)))
	
	if err != nil {
		return nil, fmt.Errorf("migrator failed: %w", err)
	}
	
	// Load the migration session
	data, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read migration session: %w", err)
	}
	
	var session MigrationSessionData
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, fmt.Errorf("failed to parse migration session: %w", err)
	}
	
	return &session, nil
}

// runValidation executes the validation tool
func (mp *MigrationPipeline) runValidation(sessionID string) (*ValidationReportData, error) {
	mp.log("info", "Running migration validation")
	
	outputPath := filepath.Join(mp.workingDir, "validation-report.json")
	
	args := []string{"run", "validator.go", mp.config.ProjectPath}
	args = append(args, fmt.Sprintf("--session=%s", sessionID))
	args = append(args, fmt.Sprintf("--output=%s", outputPath))
	
	cmd := exec.Command("go", args...)
	cmd.Dir = mp.toolsDir
	
	output, err := cmd.CombinedOutput()
	mp.log("debug", fmt.Sprintf("Validator output: %s", string(output)))
	
	// Validator may return non-zero exit code if issues found
	// We still want to process the report
	
	// Load the validation report
	data, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read validation report: %w", err)
	}
	
	var report ValidationReportData
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("failed to parse validation report: %w", err)
	}
	
	return &report, nil
}

// runTests executes the test framework
func (mp *MigrationPipeline) runTests() (*TestReportData, error) {
	mp.log("info", "Running migration tests")
	
	// First, generate default test suite if it doesn't exist
	testSuitePath := filepath.Join(mp.workingDir, "test-suite.json")
	if _, err := os.Stat(testSuitePath); os.IsNotExist(err) {
		cmd := exec.Command("go", "run", "test_framework.go", "generate-default")
		cmd.Dir = mp.toolsDir
		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("failed to generate test suite: %w", err)
		}
		
		// Move generated file to working directory
		if err := os.Rename("migration-test-suite.json", testSuitePath); err != nil {
			return nil, fmt.Errorf("failed to move test suite: %w", err)
		}
	}
	
	outputPath := filepath.Join(mp.workingDir, "test-report.json")
	
	cmd := exec.Command("go", "run", "test_framework.go", "run", testSuitePath, 
		fmt.Sprintf("--output=%s", outputPath))
	cmd.Dir = mp.toolsDir
	
	output, err := cmd.CombinedOutput()
	mp.log("debug", fmt.Sprintf("Test framework output: %s", string(output)))
	
	if err != nil {
		// Tests may fail but we still want the report
		mp.log("warning", fmt.Sprintf("Test framework had failures: %v", err))
	}
	
	// Load the test report
	data, err := os.ReadFile(outputPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read test report: %w", err)
	}
	
	var report TestReportData
	if err := json.Unmarshal(data, &report); err != nil {
		return nil, fmt.Errorf("failed to parse test report: %w", err)
	}
	
	return &report, nil
}

// createBackup creates a backup of the project
func (mp *MigrationPipeline) createBackup() (string, error) {
	if mp.config.DryRun {
		mp.log("info", "Skipping backup creation (dry run mode)")
		return "dry-run-no-backup", nil
	}
	
	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(mp.config.BackupDir, fmt.Sprintf("backup_%s", timestamp))
	
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return "", fmt.Errorf("failed to create backup directory: %w", err)
	}
	
	// Copy project files (excluding common ignore patterns)
	err := filepath.Walk(mp.config.ProjectPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip directories to ignore
		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" || 
			   name == "dist" || name == "build" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		
		// Only backup relevant files
		ext := filepath.Ext(path)
		if ext == ".go" || ext == ".templ" || ext == ".html" || ext == ".js" || 
		   ext == ".ts" || ext == ".css" || ext == ".json" {
			relPath, err := filepath.Rel(mp.config.ProjectPath, path)
			if err != nil {
				return err
			}
			
			destPath := filepath.Join(backupPath, relPath)
			destDir := filepath.Dir(destPath)
			
			if err := os.MkdirAll(destDir, 0755); err != nil {
				return err
			}
			
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			
			return os.WriteFile(destPath, content, info.Mode())
		}
		
		return nil
	})
	
	if err != nil {
		return "", fmt.Errorf("backup creation failed: %w", err)
	}
	
	mp.log("info", fmt.Sprintf("Backup created at: %s", backupPath))
	return backupPath, nil
}

// confirmProceed asks user for confirmation to proceed
func (mp *MigrationPipeline) confirmProceed(report *AnalysisReport) bool {
	fmt.Printf("\nMigration Analysis Summary:\n")
	fmt.Printf("  - Total Issues: %d\n", report.Summary["total"])
	fmt.Printf("  - Auto-fixable: %d\n", report.Summary["auto_fixable"])
	fmt.Printf("  - Errors: %d\n", report.Summary["error"])
	fmt.Printf("  - Warnings: %d\n", report.Summary["warning"])
	fmt.Printf("\nProceed with migration? (y/N): ")
	
	var response string
	fmt.Scanln(&response)
	
	return strings.ToLower(response) == "y" || strings.ToLower(response) == "yes"
}

// shouldRunTests determines if tests should be run
func (mp *MigrationPipeline) shouldRunTests() bool {
	// Skip tests in dry run mode
	if mp.config.DryRun {
		return false
	}
	
	// Run tests if not disabled
	return true
}

// printPhaseResult prints a formatted phase result
func (mp *MigrationPipeline) printPhaseResult(phase string, count interface{}, description string) {
	fmt.Printf("  ✅ %s: %v %s\n", phase, count, description)
}

// printFinalSummary prints the final migration summary
func (mp *MigrationPipeline) printFinalSummary(analysis *AnalysisReport, migration *MigrationSessionData, validation *ValidationReportData) {
	fmt.Printf("Session ID: %s\n", mp.sessionID)
	fmt.Printf("Working Directory: %s\n", mp.workingDir)
	fmt.Printf("\n")
	
	fmt.Printf("Migration Summary:\n")
	fmt.Printf("  - Issues Found: %d\n", analysis.Summary["total"])
	fmt.Printf("  - Changes Applied: %d\n", migration.Summary["total_changes"])
	fmt.Printf("  - Files Modified: %d\n", migration.Summary["files_modified"])
	fmt.Printf("  - Remaining Issues: %d\n", validation.Summary["total"])
	
	if validation.Summary["total"] > 0 {
		fmt.Printf("\n⚠️  %d validation issues remain. Please review and fix manually.\n", validation.Summary["total"])
		fmt.Printf("   See validation report: %s/validation-report.json\n", mp.workingDir)
	}
	
	fmt.Printf("\nNext Steps:\n")
	if validation.Summary["total"] == 0 {
		fmt.Printf("  ✅ Migration completed successfully!\n")
		fmt.Printf("  📝 Test your application thoroughly\n")
		fmt.Printf("  📚 Update documentation as needed\n")
	} else {
		fmt.Printf("  🔧 Fix remaining validation issues\n")
		fmt.Printf("  🔄 Run validator again: go run validator.go %s\n", mp.config.ProjectPath)
		fmt.Printf("  📝 Test your application after fixes\n")
	}
	
	if !mp.config.DryRun {
		fmt.Printf("  🗂️  Backup available for rollback if needed\n")
		fmt.Printf("  📋 Review migration logs: %s/migration.log\n", mp.workingDir)
	}
}

// log writes a log message
func (mp *MigrationPipeline) log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s: %s\n", timestamp, strings.ToUpper(level), message)
	
	mp.logFile.WriteString(logLine)
	
	if mp.config.Verbose || level == "error" {
		fmt.Print(logLine)
	}
}

// Data structures for reports (simplified versions)
type AnalysisReport struct {
	Summary map[string]int `json:"summary"`
}

type MigrationSessionData struct {
	ID      string         `json:"id"`
	Summary map[string]int `json:"summary"`
}

type ValidationReportData struct {
	Summary map[string]int `json:"summary"`
}

type TestReportData struct {
	TotalTests int `json:"total_tests"`
	Passed     int `json:"passed"`
	Failed     int `json:"failed"`
}

// parseFlags parses command line flags
func parseFlags() (*CLIConfig, error) {
	config := &CLIConfig{
		ProjectPath: ".",
		Verbose:     false,
		DryRun:      false,
		AutoFix:     true,
		Interactive: true,
		BackupDir:   "./migration_backups",
		OutputDir:   "./migration_sessions",
		LogLevel:    "info",
	}
	
	args := os.Args[1:]
	i := 0
	
	for i < len(args) {
		arg := args[i]
		
		switch {
		case arg == "--help" || arg == "-h":
			printHelp()
			os.Exit(0)
		case arg == "--verbose" || arg == "-v":
			config.Verbose = true
		case arg == "--dry-run":
			config.DryRun = true
		case arg == "--no-interactive":
			config.Interactive = false
		case strings.HasPrefix(arg, "--project="):
			config.ProjectPath = strings.TrimPrefix(arg, "--project=")
		case strings.HasPrefix(arg, "--backup-dir="):
			config.BackupDir = strings.TrimPrefix(arg, "--backup-dir=")
		case strings.HasPrefix(arg, "--output-dir="):
			config.OutputDir = strings.TrimPrefix(arg, "--output-dir=")
		default:
			if !strings.HasPrefix(arg, "-") && config.ProjectPath == "." {
				config.ProjectPath = arg
			}
		}
		i++
	}
	
	// Validate project path exists
	if _, err := os.Stat(config.ProjectPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("project path does not exist: %s", config.ProjectPath)
	}
	
	// Ensure output directories exist
	if err := os.MkdirAll(config.BackupDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create backup directory: %w", err)
	}
	
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}
	
	return config, nil
}

// printHelp prints usage information
func printHelp() {
	fmt.Println("Migration Tool - Comprehensive Component Migration Pipeline")
	fmt.Println("")
	fmt.Println("USAGE:")
	fmt.Println("  go run migrate.go [project-path] [options]")
	fmt.Println("")
	fmt.Println("ARGUMENTS:")
	fmt.Println("  project-path    Path to the project to migrate (default: current directory)")
	fmt.Println("")
	fmt.Println("OPTIONS:")
	fmt.Println("  --dry-run              Preview migration without making changes")
	fmt.Println("  --verbose, -v          Show verbose output")
	fmt.Println("  --no-interactive       Skip user prompts")
	fmt.Println("  --project=PATH         Specify project path explicitly")
	fmt.Println("  --backup-dir=DIR       Backup directory (default: ./migration_backups)")
	fmt.Println("  --output-dir=DIR       Output directory (default: ./migration_sessions)")
	fmt.Println("  --help, -h             Show this help message")
	fmt.Println("")
	fmt.Println("EXAMPLES:")
	fmt.Println("  # Migrate current directory")
	fmt.Println("  go run migrate.go")
	fmt.Println("")
	fmt.Println("  # Dry run on specific project")
	fmt.Println("  go run migrate.go /path/to/project --dry-run")
	fmt.Println("")
	fmt.Println("  # Non-interactive migration with verbose output")
	fmt.Println("  go run migrate.go --no-interactive --verbose")
	fmt.Println("")
	fmt.Println("MIGRATION PHASES:")
	fmt.Println("  1. Analysis    - Scan codebase for migration requirements")
	fmt.Println("  2. Backup      - Create backup of current code")
	fmt.Println("  3. Migration   - Apply automated transformations")
	fmt.Println("  4. Validation  - Verify migration correctness")
	fmt.Println("  5. Testing     - Run migration test suite")
	fmt.Println("")
	fmt.Println("For more information, see MIGRATION_GUIDE.md")
}

func main() {
	config, err := parseFlags()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Use --help for usage information")
		os.Exit(1)
	}
	
	pipeline, err := NewMigrationPipeline(config)
	if err != nil {
		fmt.Printf("Error creating migration pipeline: %v\n", err)
		os.Exit(1)
	}
	
	if err := pipeline.Execute(); err != nil {
		fmt.Printf("Migration failed: %v\n", err)
		fmt.Printf("Check logs at: %s/migration.log\n", pipeline.workingDir)
		os.Exit(1)
	}
}