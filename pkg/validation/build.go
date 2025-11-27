package validation

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// BuildIntegration provides validation integration with build processes
type BuildIntegration struct {
	config    BuildIntegrationConfig
	watchers  map[string]*FileWatcher
	pipelines []BuildPipeline
	hooks     []BuildHook
	cache     *BuildCache
	reporter  *BuildReporter
	mutex     sync.RWMutex
}

// BuildIntegrationConfig configures build integration
type BuildIntegrationConfig struct {
	EnableFileWatching    bool          `json:"enableFileWatching"`
	EnablePreCommitHooks  bool          `json:"enablePreCommitHooks"`
	EnableCIPipeline      bool          `json:"enableCIPipeline"`
	EnableCaching         bool          `json:"enableCaching"`
	WatchDirectories      []string      `json:"watchDirectories"`
	WatchFileExtensions   []string      `json:"watchFileExtensions"`
	IgnorePatterns        []string      `json:"ignorePatterns"`
	CacheTTL              time.Duration `json:"cacheTTL"`
	ValidationTimeout     time.Duration `json:"validationTimeout"`
	ConcurrentValidations int           `json:"concurrentValidations"`
	ReportPath            string        `json:"reportPath"`
	ExitOnFailure         bool          `json:"exitOnFailure"`
	Verbose               bool          `json:"verbose"`
}

// BuildPipeline represents a validation pipeline in the build process
type BuildPipeline struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Stages          []BuildStage      `json:"stages"`
	Triggers        []BuildTrigger    `json:"triggers"`
	Environment     map[string]string `json:"environment"`
	Timeout         time.Duration     `json:"timeout"`
	Parallel        bool              `json:"parallel"`
	ContinueOnError bool              `json:"continueOnError"`
}

// BuildStage represents a stage in the build pipeline
type BuildStage struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Type         BuildStageType    `json:"type"`
	Commands     []BuildCommand    `json:"commands"`
	Validators   []string          `json:"validators"`
	Dependencies []string          `json:"dependencies"`
	Parallel     bool              `json:"parallel"`
	Optional     bool              `json:"optional"`
	Timeout      time.Duration     `json:"timeout"`
	Environment  map[string]string `json:"environment"`
}

// BuildStageType represents the type of build stage
type BuildStageType string

const (
	BuildStageTypeValidation   BuildStageType = "validation"
	BuildStageTypeLinting      BuildStageType = "linting"
	BuildStageTypeTesting      BuildStageType = "testing"
	BuildStageTypeBuilding     BuildStageType = "building"
	BuildStageTypeDeployment   BuildStageType = "deployment"
	BuildStageTypeNotification BuildStageType = "notification"
)

// BuildCommand represents a command to execute
type BuildCommand struct {
	Name        string            `json:"name"`
	Command     string            `json:"command"`
	Args        []string          `json:"args"`
	WorkingDir  string            `json:"workingDir"`
	Environment map[string]string `json:"environment"`
	Timeout     time.Duration     `json:"timeout"`
	RetryCount  int               `json:"retryCount"`
	IgnoreError bool              `json:"ignoreError"`
}

// BuildTrigger represents when a pipeline should be triggered
type BuildTrigger struct {
	Type      BuildTriggerType `json:"type"`
	Pattern   string           `json:"pattern"`
	Branch    string           `json:"branch,omitempty"`
	Schedule  string           `json:"schedule,omitempty"`
	Manual    bool             `json:"manual"`
	Condition string           `json:"condition,omitempty"`
}

// BuildTriggerType represents the type of trigger
type BuildTriggerType string

const (
	BuildTriggerTypeFileChange BuildTriggerType = "file_change"
	BuildTriggerTypeGitCommit  BuildTriggerType = "git_commit"
	BuildTriggerTypeGitPush    BuildTriggerType = "git_push"
	BuildTriggerTypeSchedule   BuildTriggerType = "schedule"
	BuildTriggerTypeManual     BuildTriggerType = "manual"
	BuildTriggerTypeAPI        BuildTriggerType = "api"
)

// BuildHook represents hooks that can be executed at various build stages
type BuildHook struct {
	Name        string         `json:"name"`
	Type        BuildHookType  `json:"type"`
	Stage       BuildStageType `json:"stage"`
	Command     string         `json:"command"`
	Script      string         `json:"script,omitempty"`
	Condition   string         `json:"condition,omitempty"`
	FailOnError bool           `json:"failOnError"`
	Async       bool           `json:"async"`
}

// BuildHookType represents the type of hook
type BuildHookType string

const (
	BuildHookTypePreValidation  BuildHookType = "pre_validation"
	BuildHookTypePostValidation BuildHookType = "post_validation"
	BuildHookTypePreCommit      BuildHookType = "pre_commit"
	BuildHookTypePostCommit     BuildHookType = "post_commit"
	BuildHookTypePrePush        BuildHookType = "pre_push"
	BuildHookTypePostPush       BuildHookType = "post_push"
	BuildHookTypePreBuild       BuildHookType = "pre_build"
	BuildHookTypePostBuild      BuildHookType = "post_build"
)

// FileWatcher monitors files for changes
type FileWatcher struct {
	paths      []string
	extensions []string
	ignore     []string
	callback   func(string, FileEvent)
	running    bool
	stopChan   chan struct{}
	mutex      sync.RWMutex
}

// FileEvent represents a file system event
type FileEvent struct {
	Type      FileEventType `json:"type"`
	Path      string        `json:"path"`
	Timestamp time.Time     `json:"timestamp"`
}

// FileEventType represents the type of file event
type FileEventType string

const (
	FileEventTypeCreated  FileEventType = "created"
	FileEventTypeModified FileEventType = "modified"
	FileEventTypeDeleted  FileEventType = "deleted"
)

// BuildCache provides caching for build artifacts and validation results
type BuildCache struct {
	entries map[string]*BuildCacheEntry
	mutex   sync.RWMutex
	ttl     time.Duration
	maxSize int
}

// BuildCacheEntry represents a cached build artifact
type BuildCacheEntry struct {
	Key       string         `json:"key"`
	Data      any            `json:"data"`
	Hash      string         `json:"hash"`
	Timestamp time.Time      `json:"timestamp"`
	TTL       time.Duration  `json:"ttl"`
	Metadata  map[string]any `json:"metadata"`
}

// BuildReporter generates build reports
type BuildReporter struct {
	config  BuildIntegrationConfig
	results []BuildResult
	summary BuildSummary
	mutex   sync.RWMutex
}

// BuildResult represents the result of a build stage or pipeline
type BuildResult struct {
	Pipeline   string            `json:"pipeline"`
	Stage      string            `json:"stage"`
	Command    string            `json:"command,omitempty"`
	Status     BuildStatus       `json:"status"`
	StartTime  time.Time         `json:"startTime"`
	EndTime    time.Time         `json:"endTime"`
	Duration   time.Duration     `json:"duration"`
	ExitCode   int               `json:"exitCode"`
	Output     string            `json:"output"`
	Error      string            `json:"error,omitempty"`
	Validation *ValidationResult `json:"validation,omitempty"`
	Metadata   map[string]any    `json:"metadata"`
	Artifacts  []BuildArtifact   `json:"artifacts,omitempty"`
}

// BuildStatus represents the status of a build
type BuildStatus string

const (
	BuildStatusPending   BuildStatus = "pending"
	BuildStatusRunning   BuildStatus = "running"
	BuildStatusSuccess   BuildStatus = "success"
	BuildStatusFailed    BuildStatus = "failed"
	BuildStatusCancelled BuildStatus = "cancelled"
	BuildStatusSkipped   BuildStatus = "skipped"
)

// BuildArtifact represents a build artifact
type BuildArtifact struct {
	Name     string         `json:"name"`
	Path     string         `json:"path"`
	Type     string         `json:"type"`
	Size     int64          `json:"size"`
	Hash     string         `json:"hash"`
	Created  time.Time      `json:"created"`
	Metadata map[string]any `json:"metadata,omitempty"`
}

// BuildSummary contains overall build summary
type BuildSummary struct {
	TotalPipelines   int               `json:"totalPipelines"`
	SuccessPipelines int               `json:"successPipelines"`
	FailedPipelines  int               `json:"failedPipelines"`
	TotalStages      int               `json:"totalStages"`
	SuccessStages    int               `json:"successStages"`
	FailedStages     int               `json:"failedStages"`
	TotalDuration    time.Duration     `json:"totalDuration"`
	SuccessRate      float64           `json:"successRate"`
	Validations      ValidationSummary `json:"validations"`
	StartTime        time.Time         `json:"startTime"`
	EndTime          time.Time         `json:"endTime"`
	Metadata         map[string]any    `json:"metadata"`
}

// ValidationSummary contains validation-specific summary
type ValidationSummary struct {
	TotalValidations      int            `json:"totalValidations"`
	PassedValidations     int            `json:"passedValidations"`
	FailedValidations     int            `json:"failedValidations"`
	ValidationsByType     map[string]int `json:"validationsByType"`
	CriticalIssues        int            `json:"criticalIssues"`
	Warnings              int            `json:"warnings"`
	AccessibilityScore    float64        `json:"accessibilityScore"`
	PerformanceScore      float64        `json:"performanceScore"`
	ThemeConsistencyScore float64        `json:"themeConsistencyScore"`
}

// NewBuildIntegration creates a new build integration
func NewBuildIntegration() *BuildIntegration {
	config := BuildIntegrationConfig{
		EnableFileWatching:    true,
		EnablePreCommitHooks:  true,
		EnableCIPipeline:      true,
		EnableCaching:         true,
		WatchDirectories:      []string{"./src", "./components", "./schemas", "./themes"},
		WatchFileExtensions:   []string{".go", ".templ", ".json", ".yaml", ".js", ".ts", ".tsx"},
		IgnorePatterns:        []string{"node_modules", ".git", "dist", "build"},
		CacheTTL:              time.Hour * 24,
		ValidationTimeout:     time.Minute * 10,
		ConcurrentValidations: 4,
		ReportPath:            "./validation-reports",
		ExitOnFailure:         true,
		Verbose:               false,
	}

	return &BuildIntegration{
		config:    config,
		watchers:  make(map[string]*FileWatcher),
		pipelines: make([]BuildPipeline, 0),
		hooks:     make([]BuildHook, 0),
		cache:     NewBuildCache(1000, config.CacheTTL),
		reporter:  NewBuildReporter(config),
	}
}

// NewBuildCache creates a new build cache
func NewBuildCache(maxSize int, ttl time.Duration) *BuildCache {
	return &BuildCache{
		entries: make(map[string]*BuildCacheEntry),
		ttl:     ttl,
		maxSize: maxSize,
	}
}

// NewBuildReporter creates a new build reporter
func NewBuildReporter(config BuildIntegrationConfig) *BuildReporter {
	return &BuildReporter{
		config:  config,
		results: make([]BuildResult, 0),
	}
}

// SetupGitHooks sets up Git hooks for validation
func (bi *BuildIntegration) SetupGitHooks() error {
	if !bi.config.EnablePreCommitHooks {
		return nil
	}

	gitDir, err := bi.findGitDirectory()
	if err != nil {
		return fmt.Errorf("failed to find git directory: %w", err)
	}

	hooksDir := filepath.Join(gitDir, "hooks")
	if err := os.MkdirAll(hooksDir, 0755); err != nil {
		return fmt.Errorf("failed to create hooks directory: %w", err)
	}

	// Setup pre-commit hook
	preCommitPath := filepath.Join(hooksDir, "pre-commit")
	preCommitScript := `#!/bin/sh
# Auto-generated validation pre-commit hook

echo "Running validation checks..."
validation-cli validate --pre-commit

if [ $? -ne 0 ]; then
    echo "Validation failed. Commit aborted."
    exit 1
fi

echo "Validation passed. Proceeding with commit."
exit 0
`

	if err := os.WriteFile(preCommitPath, []byte(preCommitScript), 0755); err != nil {
		return fmt.Errorf("failed to write pre-commit hook: %w", err)
	}

	// Setup pre-push hook
	prePushPath := filepath.Join(hooksDir, "pre-push")
	prePushScript := `#!/bin/sh
# Auto-generated validation pre-push hook

echo "Running comprehensive validation checks..."
validation-cli validate --comprehensive

if [ $? -ne 0 ]; then
    echo "Validation failed. Push aborted."
    exit 1
fi

echo "All validations passed. Proceeding with push."
exit 0
`

	if err := os.WriteFile(prePushPath, []byte(prePushScript), 0755); err != nil {
		return fmt.Errorf("failed to write pre-push hook: %w", err)
	}

	return nil
}

// StartFileWatching starts watching files for changes
func (bi *BuildIntegration) StartFileWatching() error {
	if !bi.config.EnableFileWatching {
		return nil
	}

	for _, dir := range bi.config.WatchDirectories {
		watcher := NewFileWatcher([]string{dir}, bi.config.WatchFileExtensions, bi.config.IgnorePatterns)
		watcher.SetCallback(bi.onFileChanged)

		if err := watcher.Start(); err != nil {
			return fmt.Errorf("failed to start file watcher for %s: %w", dir, err)
		}

		bi.watchers[dir] = watcher
	}

	return nil
}

// StopFileWatching stops file watching
func (bi *BuildIntegration) StopFileWatching() {
	for _, watcher := range bi.watchers {
		watcher.Stop()
	}
}

// onFileChanged handles file change events
func (bi *BuildIntegration) onFileChanged(path string, event FileEvent) {
	if bi.config.Verbose {
		fmt.Printf("File changed: %s (%s)\n", path, event.Type)
	}

	// Determine which validations to run based on file type
	validations := bi.determineValidationsForFile(path)

	// Run validations asynchronously
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), bi.config.ValidationTimeout)
		defer cancel()

		for _, validation := range validations {
			if err := bi.runValidation(ctx, validation, path); err != nil {
				fmt.Printf("Validation failed for %s: %v\n", path, err)
			}
		}
	}()
}

// AddPipeline adds a build pipeline
func (bi *BuildIntegration) AddPipeline(pipeline BuildPipeline) {
	bi.mutex.Lock()
	defer bi.mutex.Unlock()
	bi.pipelines = append(bi.pipelines, pipeline)
}

// AddHook adds a build hook
func (bi *BuildIntegration) AddHook(hook BuildHook) {
	bi.mutex.Lock()
	defer bi.mutex.Unlock()
	bi.hooks = append(bi.hooks, hook)
}

// RunPipeline executes a build pipeline
func (bi *BuildIntegration) RunPipeline(ctx context.Context, pipelineName string) (*BuildSummary, error) {
	var pipeline *BuildPipeline
	for i := range bi.pipelines {
		if bi.pipelines[i].Name == pipelineName {
			pipeline = &bi.pipelines[i]
			break
		}
	}

	if pipeline == nil {
		return nil, fmt.Errorf("pipeline %s not found", pipelineName)
	}

	start := time.Now()
	summary := &BuildSummary{
		TotalPipelines: 1,
		StartTime:      start,
		Metadata:       make(map[string]any),
		Validations: ValidationSummary{
			ValidationsByType: make(map[string]int),
		},
	}

	// Run pre-build hooks
	if err := bi.runHooks(ctx, BuildHookTypePreBuild, pipeline); err != nil {
		return summary, fmt.Errorf("pre-build hooks failed: %w", err)
	}

	// Execute pipeline stages
	results := make([]BuildResult, 0)

	if pipeline.Parallel {
		results = bi.runStagesParallel(ctx, pipeline)
	} else {
		results = bi.runStagesSequential(ctx, pipeline)
	}

	// Run post-build hooks
	if err := bi.runHooks(ctx, BuildHookTypePostBuild, pipeline); err != nil {
		fmt.Printf("Warning: post-build hooks failed: %v\n", err)
	}

	// Update summary
	summary.EndTime = time.Now()
	summary.TotalDuration = summary.EndTime.Sub(summary.StartTime)
	bi.updateSummaryFromResults(summary, results)

	// Generate report
	bi.reporter.AddResults(results)
	if err := bi.reporter.GenerateReport(); err != nil {
		fmt.Printf("Warning: failed to generate report: %v\n", err)
	}

	return summary, nil
}

// runStagesSequential runs pipeline stages sequentially
func (bi *BuildIntegration) runStagesSequential(ctx context.Context, pipeline *BuildPipeline) []BuildResult {
	results := make([]BuildResult, 0)

	for _, stage := range pipeline.Stages {
		stageResults := bi.runStage(ctx, pipeline, &stage)
		results = append(results, stageResults...)

		// Check if stage failed and should stop pipeline
		stageFailed := false
		for _, result := range stageResults {
			if result.Status == BuildStatusFailed {
				stageFailed = true
				break
			}
		}

		if stageFailed && !pipeline.ContinueOnError {
			break
		}
	}

	return results
}

// runStagesParallel runs pipeline stages in parallel
func (bi *BuildIntegration) runStagesParallel(ctx context.Context, pipeline *BuildPipeline) []BuildResult {
	resultChan := make(chan []BuildResult, len(pipeline.Stages))
	var wg sync.WaitGroup

	for _, stage := range pipeline.Stages {
		wg.Add(1)
		go func(s BuildStage) {
			defer wg.Done()
			stageResults := bi.runStage(ctx, pipeline, &s)
			resultChan <- stageResults
		}(stage)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	results := make([]BuildResult, 0)
	for stageResults := range resultChan {
		results = append(results, stageResults...)
	}

	return results
}

// runStage executes a build stage
func (bi *BuildIntegration) runStage(ctx context.Context, pipeline *BuildPipeline, stage *BuildStage) []BuildResult {
	results := make([]BuildResult, 0)

	// Run pre-validation hooks
	bi.runHooks(ctx, BuildHookTypePreValidation, pipeline)

	if stage.Parallel {
		// Run commands in parallel
		resultChan := make(chan BuildResult, len(stage.Commands))
		var wg sync.WaitGroup

		for _, command := range stage.Commands {
			wg.Add(1)
			go func(cmd BuildCommand) {
				defer wg.Done()
				result := bi.runCommand(ctx, pipeline, stage, &cmd)
				resultChan <- result
			}(command)
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		for result := range resultChan {
			results = append(results, result)
		}
	} else {
		// Run commands sequentially
		for _, command := range stage.Commands {
			result := bi.runCommand(ctx, pipeline, stage, &command)
			results = append(results, result)

			// Stop if command failed and stage is not optional
			if result.Status == BuildStatusFailed && !stage.Optional {
				break
			}
		}
	}

	// Run post-validation hooks
	bi.runHooks(ctx, BuildHookTypePostValidation, pipeline)

	return results
}

// runCommand executes a build command
func (bi *BuildIntegration) runCommand(ctx context.Context, pipeline *BuildPipeline, stage *BuildStage, command *BuildCommand) BuildResult {
	start := time.Now()

	result := BuildResult{
		Pipeline:  pipeline.Name,
		Stage:     stage.Name,
		Command:   command.Name,
		Status:    BuildStatusRunning,
		StartTime: start,
		Metadata:  make(map[string]any),
	}

	// Check cache first
	if bi.config.EnableCaching {
		cacheKey := bi.generateCacheKey(pipeline.Name, stage.Name, command)
		if cached := bi.cache.Get(cacheKey); cached != nil {
			result.Status = BuildStatusSuccess
			result.Output = "Retrieved from cache"
			result.EndTime = time.Now()
			result.Duration = time.Since(start)
			return result
		}
	}

	// Execute command based on stage type
	switch stage.Type {
	case BuildStageTypeValidation:
		result = bi.runValidationCommand(ctx, command, result)
	case BuildStageTypeLinting:
		result = bi.runLintingCommand(ctx, command, result)
	case BuildStageTypeTesting:
		result = bi.runTestingCommand(ctx, command, result)
	default:
		result = bi.runGenericCommand(ctx, command, result)
	}

	result.EndTime = time.Now()
	result.Duration = time.Since(start)

	// Cache successful results
	if bi.config.EnableCaching && result.Status == BuildStatusSuccess {
		cacheKey := bi.generateCacheKey(pipeline.Name, stage.Name, command)
		bi.cache.Set(cacheKey, result)
	}

	return result
}

// runValidationCommand runs a validation command
func (bi *BuildIntegration) runValidationCommand(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	// Determine validation type from command
	validationType := bi.getValidationTypeFromCommand(command)

	switch validationType {
	case "component":
		return bi.runComponentValidation(ctx, command, result)
	case "schema":
		return bi.runSchemaValidation(ctx, command, result)
	case "accessibility":
		return bi.runAccessibilityValidation(ctx, command, result)
	case "performance":
		return bi.runPerformanceValidation(ctx, command, result)
	case "theme":
		return bi.runThemeValidation(ctx, command, result)
	default:
		return bi.runGenericValidation(ctx, command, result)
	}
}

// runComponentValidation runs component validation
func (bi *BuildIntegration) runComponentValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	validator := NewComponentValidator()

	// Load component files from working directory
	components, err := bi.loadComponentsFromDirectory(command.WorkingDir)
	if err != nil {
		result.Status = BuildStatusFailed
		result.Error = fmt.Sprintf("Failed to load components: %v", err)
		return result
	}

	// Validate each component
	validationResults := make([]*ValidationResult, 0)
	for _, component := range components {
		componentResult := validator.ValidateComponent(component)
		validationResults = append(validationResults, componentResult)
	}

	// Combine results
	combinedResult := bi.combineValidationResults(validationResults)
	result.Validation = combinedResult

	if combinedResult.Valid {
		result.Status = BuildStatusSuccess
		result.Output = fmt.Sprintf("Validated %d components successfully", len(components))
	} else {
		result.Status = BuildStatusFailed
		result.Error = fmt.Sprintf("Component validation failed with %d errors", len(combinedResult.Errors))
	}

	return result
}

// runSchemaValidation runs schema validation
func (bi *BuildIntegration) runSchemaValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	validator := NewSchemaValidator()

	// Load schema files
	schemas, err := bi.loadSchemasFromDirectory(command.WorkingDir)
	if err != nil {
		result.Status = BuildStatusFailed
		result.Error = fmt.Sprintf("Failed to load schemas: %v", err)
		return result
	}

	// Validate each schema
	allValid := true
	errorCount := 0

	for _, schema := range schemas {
		schemaResult := validator.ValidateSchema(schema)
		if !schemaResult.Valid {
			allValid = false
			errorCount += len(schemaResult.Errors)
		}
	}

	if allValid {
		result.Status = BuildStatusSuccess
		result.Output = fmt.Sprintf("Validated %d schemas successfully", len(schemas))
	} else {
		result.Status = BuildStatusFailed
		result.Error = fmt.Sprintf("Schema validation failed with %d errors", errorCount)
	}

	return result
}

// File and directory operations

func (bi *BuildIntegration) loadComponentsFromDirectory(dir string) ([]*ComponentInstance, error) {
	components := make([]*ComponentInstance, 0)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if bi.shouldIgnoreFile(path) {
			return nil
		}

		if bi.isComponentFile(path) {
			component, err := bi.loadComponentFromFile(path)
			if err != nil {
				return fmt.Errorf("failed to load component from %s: %w", path, err)
			}
			if component != nil {
				components = append(components, component)
			}
		}

		return nil
	})

	return components, err
}

func (bi *BuildIntegration) loadSchemasFromDirectory(dir string) ([]any, error) {
	schemas := make([]any, 0)

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if bi.shouldIgnoreFile(path) {
			return nil
		}

		if bi.isSchemaFile(path) {
			schema, err := bi.loadSchemaFromFile(path)
			if err != nil {
				return fmt.Errorf("failed to load schema from %s: %w", path, err)
			}
			if schema != nil {
				schemas = append(schemas, schema)
			}
		}

		return nil
	})

	return schemas, err
}

func (bi *BuildIntegration) loadComponentFromFile(path string) (*ComponentInstance, error) {
	// Simplified component loading - would need proper parsing based on file format
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Try to parse as JSON first
	var component ComponentInstance
	if err := json.Unmarshal(data, &component); err == nil {
		component.Location = &SourceLocation{
			File: path,
			Line: 1,
		}
		return &component, nil
	}

	// For other file types (.templ, .go), would need specialized parsing
	return nil, nil
}

func (bi *BuildIntegration) loadSchemaFromFile(path string) (any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var schema any
	if err := json.Unmarshal(data, &schema); err != nil {
		return nil, err
	}

	return schema, nil
}

// Helper methods

func (bi *BuildIntegration) shouldIgnoreFile(path string) bool {
	for _, pattern := range bi.config.IgnorePatterns {
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

func (bi *BuildIntegration) isComponentFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".templ" || ext == ".tsx" || ext == ".jsx" ||
		(ext == ".json" && strings.Contains(path, "component"))
}

func (bi *BuildIntegration) isSchemaFile(path string) bool {
	ext := filepath.Ext(path)
	return (ext == ".json" || ext == ".yaml" || ext == ".yml") &&
		strings.Contains(strings.ToLower(path), "schema")
}

func (bi *BuildIntegration) determineValidationsForFile(path string) []string {
	validations := make([]string, 0)

	if bi.isComponentFile(path) {
		validations = append(validations, "component", "accessibility")
	}

	if bi.isSchemaFile(path) {
		validations = append(validations, "schema")
	}

	if strings.Contains(path, "theme") {
		validations = append(validations, "theme")
	}

	return validations
}

func (bi *BuildIntegration) runValidation(ctx context.Context, validationType, path string) error {
	// Simplified validation execution
	switch validationType {
	case "component":
		return bi.validateComponentFile(ctx, path)
	case "schema":
		return bi.validateSchemaFile(ctx, path)
	case "accessibility":
		return bi.validateAccessibilityFile(ctx, path)
	case "theme":
		return bi.validateThemeFile(ctx, path)
	default:
		return fmt.Errorf("unknown validation type: %s", validationType)
	}
}

func (bi *BuildIntegration) validateComponentFile(ctx context.Context, path string) error {
	component, err := bi.loadComponentFromFile(path)
	if err != nil {
		return err
	}

	if component == nil {
		return nil // Not a component file
	}

	validator := NewComponentValidator()
	result := validator.ValidateComponent(component)

	if !result.Valid {
		return fmt.Errorf("component validation failed: %d errors", len(result.Errors))
	}

	return nil
}

func (bi *BuildIntegration) validateSchemaFile(ctx context.Context, path string) error {
	schema, err := bi.loadSchemaFromFile(path)
	if err != nil {
		return err
	}

	validator := NewSchemaValidator()
	result := validator.ValidateSchema(schema)

	if !result.Valid {
		return fmt.Errorf("schema validation failed: %d errors", len(result.Errors))
	}

	return nil
}

func (bi *BuildIntegration) validateAccessibilityFile(ctx context.Context, path string) error {
	// Accessibility validation for file
	return nil
}

func (bi *BuildIntegration) validateThemeFile(ctx context.Context, path string) error {
	// Theme validation for file
	return nil
}

// Additional command runners and utility methods...

func (bi *BuildIntegration) runLintingCommand(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	return bi.runGenericCommand(ctx, command, result)
}

func (bi *BuildIntegration) runTestingCommand(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	return bi.runGenericCommand(ctx, command, result)
}

func (bi *BuildIntegration) runGenericCommand(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	cmd := exec.CommandContext(ctx, command.Command, command.Args...)

	if command.WorkingDir != "" {
		cmd.Dir = command.WorkingDir
	}

	// Set environment variables
	cmd.Env = os.Environ()
	for key, value := range command.Environment {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", key, value))
	}

	output, err := cmd.CombinedOutput()
	result.Output = string(output)

	if err != nil {
		result.Status = BuildStatusFailed
		result.Error = err.Error()
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
		}
	} else {
		result.Status = BuildStatusSuccess
		result.ExitCode = 0
	}

	return result
}

func (bi *BuildIntegration) runGenericValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	// Run generic validation command
	return bi.runGenericCommand(ctx, command, result)
}

func (bi *BuildIntegration) runAccessibilityValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	// Accessibility validation logic
	return result
}

func (bi *BuildIntegration) runPerformanceValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	// Performance validation logic
	return result
}

func (bi *BuildIntegration) runThemeValidation(ctx context.Context, command *BuildCommand, result BuildResult) BuildResult {
	// Theme validation logic
	return result
}

func (bi *BuildIntegration) getValidationTypeFromCommand(command *BuildCommand) string {
	if strings.Contains(command.Name, "component") {
		return "component"
	}
	if strings.Contains(command.Name, "schema") {
		return "schema"
	}
	if strings.Contains(command.Name, "a11y") || strings.Contains(command.Name, "accessibility") {
		return "accessibility"
	}
	if strings.Contains(command.Name, "performance") || strings.Contains(command.Name, "perf") {
		return "performance"
	}
	if strings.Contains(command.Name, "theme") {
		return "theme"
	}
	return "generic"
}

func (bi *BuildIntegration) combineValidationResults(results []*ValidationResult) *ValidationResult {
	combined := &ValidationResult{
		Valid:     true,
		Timestamp: time.Now(),
		Metadata:  make(map[string]any),
	}

	for _, result := range results {
		if !result.Valid {
			combined.Valid = false
		}
		combined.Errors = append(combined.Errors, result.Errors...)
		combined.Warnings = append(combined.Warnings, result.Warnings...)
	}

	return combined
}

func (bi *BuildIntegration) runHooks(ctx context.Context, hookType BuildHookType, pipeline *BuildPipeline) error {
	for _, hook := range bi.hooks {
		if hook.Type == hookType {
			if err := bi.runHook(ctx, &hook); err != nil && hook.FailOnError {
				return err
			}
		}
	}
	return nil
}

func (bi *BuildIntegration) runHook(ctx context.Context, hook *BuildHook) error {
	if hook.Script != "" {
		return bi.runHookScript(ctx, hook)
	}
	return bi.runHookCommand(ctx, hook)
}

func (bi *BuildIntegration) runHookScript(ctx context.Context, hook *BuildHook) error {
	// Execute hook script
	return nil
}

func (bi *BuildIntegration) runHookCommand(ctx context.Context, hook *BuildHook) error {
	// Execute hook command
	cmd := exec.CommandContext(ctx, "sh", "-c", hook.Command)
	return cmd.Run()
}

func (bi *BuildIntegration) generateCacheKey(pipeline, stage string, command *BuildCommand) string {
	return fmt.Sprintf("%s:%s:%s", pipeline, stage, command.Name)
}

func (bi *BuildIntegration) updateSummaryFromResults(summary *BuildSummary, results []BuildResult) {
	for _, result := range results {
		summary.TotalStages++

		switch result.Status {
		case BuildStatusSuccess:
			summary.SuccessStages++
		case BuildStatusFailed:
			summary.FailedStages++
		}

		if result.Validation != nil {
			summary.Validations.TotalValidations++
			if result.Validation.Valid {
				summary.Validations.PassedValidations++
			} else {
				summary.Validations.FailedValidations++
				summary.Validations.CriticalIssues += len(result.Validation.Errors)
				summary.Validations.Warnings += len(result.Validation.Warnings)
			}
		}
	}

	if summary.TotalStages > 0 {
		summary.SuccessRate = float64(summary.SuccessStages) / float64(summary.TotalStages)
	}

	if summary.TotalStages > 0 {
		summary.SuccessPipelines = 1
	}
}

func (bi *BuildIntegration) findGitDirectory() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		gitDir := filepath.Join(dir, ".git")
		if info, err := os.Stat(gitDir); err == nil && info.IsDir() {
			return gitDir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("not a git repository")
}

// FileWatcher implementation

func NewFileWatcher(paths, extensions, ignore []string) *FileWatcher {
	return &FileWatcher{
		paths:      paths,
		extensions: extensions,
		ignore:     ignore,
		stopChan:   make(chan struct{}),
	}
}

func (fw *FileWatcher) SetCallback(callback func(string, FileEvent)) {
	fw.callback = callback
}

func (fw *FileWatcher) Start() error {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	if fw.running {
		return fmt.Errorf("file watcher already running")
	}

	fw.running = true

	go fw.watch()
	return nil
}

func (fw *FileWatcher) Stop() {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	if !fw.running {
		return
	}

	fw.running = false
	close(fw.stopChan)
}

func (fw *FileWatcher) watch() {
	// Simplified file watching implementation
	// In production, use a proper file watching library like fsnotify
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-fw.stopChan:
			return
		case <-ticker.C:
			fw.checkFiles()
		}
	}
}

func (fw *FileWatcher) checkFiles() {
	for _, path := range fw.paths {
		filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if fw.shouldIgnoreFile(filePath) {
				return nil
			}

			if fw.isWatchedFile(filePath) {
				// Check file modification time
				info, err := d.Info()
				if err != nil {
					return nil
				}

				// Simplified change detection
				if time.Since(info.ModTime()) < time.Second*2 {
					event := FileEvent{
						Type:      FileEventTypeModified,
						Path:      filePath,
						Timestamp: time.Now(),
					}

					if fw.callback != nil {
						fw.callback(filePath, event)
					}
				}
			}

			return nil
		})
	}
}

func (fw *FileWatcher) shouldIgnoreFile(path string) bool {
	for _, pattern := range fw.ignore {
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

func (fw *FileWatcher) isWatchedFile(path string) bool {
	if len(fw.extensions) == 0 {
		return true
	}

	ext := filepath.Ext(path)
	for _, watchExt := range fw.extensions {
		if ext == watchExt {
			return true
		}
	}
	return false
}

// BuildCache implementation

func (bc *BuildCache) Get(key string) *BuildCacheEntry {
	bc.mutex.RLock()
	defer bc.mutex.RUnlock()

	entry, exists := bc.entries[key]
	if !exists {
		return nil
	}

	// Check if entry has expired
	if time.Since(entry.Timestamp) > entry.TTL {
		delete(bc.entries, key)
		return nil
	}

	return entry
}

func (bc *BuildCache) Set(key string, data any) {
	bc.mutex.Lock()
	defer bc.mutex.Unlock()

	// Remove oldest entries if cache is full
	if len(bc.entries) >= bc.maxSize {
		bc.evictOldest()
	}

	bc.entries[key] = &BuildCacheEntry{
		Key:       key,
		Data:      data,
		Timestamp: time.Now(),
		TTL:       bc.ttl,
		Metadata:  make(map[string]any),
	}
}

func (bc *BuildCache) evictOldest() {
	oldestKey := ""
	oldestTime := time.Now()

	for key, entry := range bc.entries {
		if entry.Timestamp.Before(oldestTime) {
			oldestTime = entry.Timestamp
			oldestKey = key
		}
	}

	if oldestKey != "" {
		delete(bc.entries, oldestKey)
	}
}

// BuildReporter implementation

func (br *BuildReporter) AddResults(results []BuildResult) {
	br.mutex.Lock()
	defer br.mutex.Unlock()
	br.results = append(br.results, results...)
}

func (br *BuildReporter) GenerateReport() error {
	br.mutex.RLock()
	defer br.mutex.RUnlock()

	// Create report directory
	if err := os.MkdirAll(br.config.ReportPath, 0755); err != nil {
		return fmt.Errorf("failed to create report directory: %w", err)
	}

	// Generate JSON report
	reportData := map[string]any{
		"summary":   br.summary,
		"results":   br.results,
		"timestamp": time.Now(),
	}

	jsonData, err := json.MarshalIndent(reportData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report data: %w", err)
	}

	jsonPath := filepath.Join(br.config.ReportPath, "build-report.json")
	if err := os.WriteFile(jsonPath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON report: %w", err)
	}

	// Generate HTML report
	htmlPath := filepath.Join(br.config.ReportPath, "build-report.html")
	if err := br.generateHTMLReport(htmlPath); err != nil {
		return fmt.Errorf("failed to write HTML report: %w", err)
	}

	return nil
}

func (br *BuildReporter) generateHTMLReport(path string) error {
	// Generate HTML report
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Build Validation Report</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .summary { background: #f5f5f5; padding: 15px; border-radius: 5px; margin-bottom: 20px; }
        .success { color: green; }
        .failed { color: red; }
        .stage { margin: 10px 0; padding: 10px; border: 1px solid #ddd; border-radius: 3px; }
    </style>
</head>
<body>
    <h1>Build Validation Report</h1>
    <div class="summary">
        <h2>Summary</h2>
        <p>Total Stages: %d</p>
        <p class="success">Successful: %d</p>
        <p class="failed">Failed: %d</p>
        <p>Success Rate: %.2f%%</p>
    </div>
    <div class="results">
        <h2>Results</h2>
        %s
    </div>
</body>
</html>`

	resultHTML := ""
	for _, result := range br.results {
		statusClass := "success"
		if result.Status == BuildStatusFailed {
			statusClass = "failed"
		}

		resultHTML += fmt.Sprintf(`
        <div class="stage">
            <h3 class="%s">%s - %s</h3>
            <p>Duration: %v</p>
            <p>Status: %s</p>
            %s
        </div>`,
			statusClass,
			result.Pipeline,
			result.Stage,
			result.Duration,
			result.Status,
			result.Output,
		)
	}

	htmlContent := fmt.Sprintf(html,
		br.summary.TotalStages,
		br.summary.SuccessStages,
		br.summary.FailedStages,
		br.summary.SuccessRate*100,
		resultHTML,
	)

	return os.WriteFile(path, []byte(htmlContent), 0644)
}
