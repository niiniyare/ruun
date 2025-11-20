package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/pkg/utils"
)

// ModalService provides business logic for modal operations
type ModalService struct {
	// Configuration
	DefaultCacheTimeout time.Duration
	MaxModalStack       int

	// Feature flags
	EnableModalStack bool
	EnableAutoSave   bool
	EnableVersioning bool
	EnableFileUpload bool
	EnableSearch     bool

	// Performance settings
	MaxFileSize       int64 // bytes
	UploadConcurrency int
	SearchDebounce    time.Duration
	AutoSaveInterval  time.Duration

	// Integration providers
	FileProvider       FileUploadProvider
	SearchProvider     SearchProvider
	ValidationProvider ValidationProvider
	WorkflowProvider   WorkflowProvider
	PermissionProvider ModalPermissionProvider

	// Storage providers
	DataProvider  ModalDataProvider
	StateProvider ModalStateProvider
	CacheProvider CacheProvider
}

// ModalDataProvider interface for data persistence
type ModalDataProvider interface {
	Save(ctx context.Context, modalID string, data map[string]any) error
	Load(ctx context.Context, modalID string) (map[string]any, error)
	Delete(ctx context.Context, modalID string) error
	AutoSave(ctx context.Context, modalID string, data map[string]any) error
	GetVersions(ctx context.Context, modalID string) ([]DataVersion, error)
	SaveVersion(ctx context.Context, modalID string, data map[string]any, comment string) (*DataVersion, error)
}

// ModalStateProvider interface for modal state management
type ModalStateProvider interface {
	GetState(ctx context.Context, sessionID string) (*ModalStackState, error)
	UpdateState(ctx context.Context, sessionID string, state *ModalStackState) error
	PushModal(ctx context.Context, sessionID string, modalID string, config *EnhancedModalProps) error
	PopModal(ctx context.Context, sessionID string) (*EnhancedModalProps, error)
	ClearStack(ctx context.Context, sessionID string) error
}

// FileUploadProvider interface for file upload operations
type FileUploadProvider interface {
	Upload(ctx context.Context, file *UploadedFile) (*FileMetadata, error)
	Delete(ctx context.Context, fileID string) error
	GetMetadata(ctx context.Context, fileID string) (*FileMetadata, error)
	ValidateFile(ctx context.Context, file *UploadedFile, constraints *FileConstraints) error
}

// SearchProvider interface for modal search functionality
type SearchProvider interface {
	Search(ctx context.Context, query string, filters map[string]any) ([]SearchResult, error)
	GetSuggestions(ctx context.Context, query string, limit int) ([]SearchSuggestion, error)
	SaveSearch(ctx context.Context, userID string, query string) error
	GetRecentSearches(ctx context.Context, userID string, limit int) ([]string, error)
}

// ValidationProvider interface for form validation
type ValidationProvider interface {
	ValidateField(ctx context.Context, field string, value any, rules []ValidationRule) (*ValidationResult, error)
	ValidateForm(ctx context.Context, data map[string]any, schema *schema.FormSchema) (*ValidationResult, error)
	ValidateStep(ctx context.Context, stepData map[string]any, step *EnhancedModalStep) (*ValidationResult, error)
}

// WorkflowProvider interface for business workflow integration
type WorkflowProvider interface {
	StartWorkflow(ctx context.Context, workflowID string, input map[string]any) (*WorkflowInstance, error)
	ContinueWorkflow(ctx context.Context, instanceID string, action string, data map[string]any) (*WorkflowInstance, error)
	GetWorkflowState(ctx context.Context, instanceID string) (*WorkflowState, error)
	CompleteWorkflow(ctx context.Context, instanceID string, result map[string]any) error
}

// ModalPermissionProvider interface for permission-based modal access
type ModalPermissionProvider interface {
	CanOpenModal(ctx context.Context, user *User, modalType string, entityID string) bool
	CanPerformAction(ctx context.Context, user *User, modalID string, action string) bool
	GetVisibleActions(ctx context.Context, user *User, actions []EnhancedModalAction) ([]EnhancedModalAction, error)
	ValidateSecureAccess(ctx context.Context, user *User, modalProps *EnhancedModalProps) error
}

// CacheProvider interface for caching modal data
type CacheProvider interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

// Data structures for modal business logic

// ModalStackState represents the current modal stack state
type ModalStackState struct {
	SessionID    string           `json:"sessionId"`
	Stack        []ModalStackItem `json:"stack"`
	MaxDepth     int              `json:"maxDepth"`
	CurrentDepth int              `json:"currentDepth"`
	LastUpdated  time.Time        `json:"lastUpdated"`
}

// ModalStackItem represents a modal in the stack
type ModalStackItem struct {
	ID       string              `json:"id"`
	Type     EnhancedModalType   `json:"type"`
	Props    *EnhancedModalProps `json:"props"`
	Data     map[string]any      `json:"data,omitempty"`
	State    ModalState          `json:"state"`
	OpenedAt time.Time           `json:"openedAt"`
	ParentID string              `json:"parentId,omitempty"`
	Context  map[string]any      `json:"context,omitempty"`
}

// ModalState represents the current state of a modal
type ModalState struct {
	// Basic state
	Open    bool `json:"open"`
	Loading bool `json:"loading"`
	Saving  bool `json:"saving"`
	Dirty   bool `json:"dirty"`
	Valid   bool `json:"valid"`

	// Multi-step state
	CurrentStep    int   `json:"currentStep"`
	TotalSteps     int   `json:"totalSteps"`
	CompletedSteps []int `json:"completedSteps"`
	SkippedSteps   []int `json:"skippedSteps"`

	// Form state
	FormData        map[string]any    `json:"formData"`
	Errors          map[string]string `json:"errors"`
	Touched         map[string]bool   `json:"touched"`
	ValidationState string            `json:"validationState"`

	// File upload state
	UploadQueue    []UploadTask   `json:"uploadQueue,omitempty"`
	UploadProgress map[string]int `json:"uploadProgress,omitempty"`
	UploadedFiles  []FileMetadata `json:"uploadedFiles,omitempty"`

	// Search state
	SearchQuery   string         `json:"searchQuery"`
	SearchResults []SearchResult `json:"searchResults"`
	SearchLoading bool           `json:"searchLoading"`
	SelectedItems []string       `json:"selectedItems"`

	// Auto-save state
	LastSaved       *time.Time `json:"lastSaved,omitempty"`
	AutoSaveEnabled bool       `json:"autoSaveEnabled"`
	SaveCount       int        `json:"saveCount"`

	// Workflow state
	WorkflowID    string `json:"workflowId,omitempty"`
	InstanceID    string `json:"instanceId,omitempty"`
	WorkflowState string `json:"workflowState,omitempty"`

	// Version control state
	CurrentVersion string `json:"currentVersion,omitempty"`
	HasVersions    bool   `json:"hasVersions"`
	CanUndo        bool   `json:"canUndo"`
	CanRedo        bool   `json:"canRedo"`

	// Custom state
	CustomState map[string]any `json:"customState,omitempty"`
}

// Supporting data structures

// DataVersion represents a versioned data snapshot
type DataVersion struct {
	ID        string         `json:"id"`
	Version   string         `json:"version"`
	Data      map[string]any `json:"data"`
	Comment   string         `json:"comment,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	CreatedBy string         `json:"createdBy,omitempty"`
}

// UploadedFile represents a file being uploaded
type UploadedFile struct {
	Name     string            `json:"name"`
	Size     int64             `json:"size"`
	Type     string            `json:"type"`
	Content  []byte            `json:"-"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// FileMetadata represents uploaded file metadata
type FileMetadata struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Size         int64             `json:"size"`
	Type         string            `json:"type"`
	URL          string            `json:"url"`
	ThumbnailURL string            `json:"thumbnailUrl,omitempty"`
	UploadedAt   time.Time         `json:"uploadedAt"`
	UploadedBy   string            `json:"uploadedBy,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// FileConstraints represents file upload constraints
type FileConstraints struct {
	MaxSize      int64    `json:"maxSize"`
	AllowedTypes []string `json:"allowedTypes"`
	Required     bool     `json:"required"`
}

// UploadTask represents a file upload task
type UploadTask struct {
	ID       string        `json:"id"`
	File     *UploadedFile `json:"file"`
	Progress int           `json:"progress"`
	Status   string        `json:"status"` // "pending", "uploading", "completed", "failed"
	Error    string        `json:"error,omitempty"`
}

// SearchResult represents a search result item
type SearchResult struct {
	ID          string         `json:"id"`
	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Icon        string         `json:"icon,omitempty"`
	Score       float64        `json:"score"`
	Metadata    map[string]any `json:"metadata,omitempty"`
}

// SearchSuggestion represents a search suggestion
type SearchSuggestion struct {
	Query string  `json:"query"`
	Type  string  `json:"type"`
	Count int     `json:"count,omitempty"`
	Score float64 `json:"score"`
}

// ValidationRule represents a validation rule
type ValidationRule struct {
	Type       string         `json:"type"`
	Value      any            `json:"value,omitempty"`
	Message    string         `json:"message,omitempty"`
	Conditions map[string]any `json:"conditions,omitempty"`
}

// ValidationResult represents validation result
type ValidationResult struct {
	Valid    bool              `json:"valid"`
	Errors   map[string]string `json:"errors,omitempty"`
	Warnings map[string]string `json:"warnings,omitempty"`
}

// WorkflowInstance represents a workflow instance
type WorkflowInstance struct {
	ID          string         `json:"id"`
	WorkflowID  string         `json:"workflowId"`
	State       string         `json:"state"`
	CurrentStep string         `json:"currentStep"`
	Data        map[string]any `json:"data"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}

// WorkflowState represents workflow execution state
type WorkflowState struct {
	Current   string         `json:"current"`
	Previous  []string       `json:"previous"`
	Available []string       `json:"available"`
	Data      map[string]any `json:"data"`
	Metadata  map[string]any `json:"metadata"`
}

// User represents the current user (referenced from navigation logic)
type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	DisplayName string   `json:"displayName"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

// NewModalService creates a new modal service
func NewModalService() *ModalService {
	return &ModalService{
		DefaultCacheTimeout: 5 * time.Minute,
		MaxModalStack:       5,
		EnableModalStack:    true,
		EnableAutoSave:      true,
		EnableVersioning:    false,
		EnableFileUpload:    true,
		EnableSearch:        true,
		MaxFileSize:         10 * 1024 * 1024, // 10MB
		UploadConcurrency:   3,
		SearchDebounce:      300 * time.Millisecond,
		AutoSaveInterval:    30 * time.Second,
	}
}

// Core modal service methods

// OpenModal opens a modal and manages the modal stack
func (s *ModalService) OpenModal(ctx context.Context, sessionID string, props *EnhancedModalProps) (*ModalState, error) {
	// Check permissions
	if s.PermissionProvider != nil {
		user := getUserFromContext(ctx)
		if user != nil && props.SecureMode {
			if err := s.PermissionProvider.ValidateSecureAccess(ctx, user, props); err != nil {
				return nil, fmt.Errorf("permission denied: %w", err)
			}
		}
	}

	// Initialize modal state
	state := &ModalState{
		Open:            true,
		Loading:         false,
		Saving:          false,
		Dirty:           false,
		Valid:           true,
		CurrentStep:     1,
		TotalSteps:      len(props.Steps),
		FormData:        make(map[string]any),
		Errors:          make(map[string]string),
		Touched:         make(map[string]bool),
		AutoSaveEnabled: props.AutoSave,
		CustomState:     make(map[string]any),
	}

	// Initialize form data if provided
	if props.FormData != nil {
		state.FormData = props.FormData
	}

	// Load initial data if URL provided
	if props.LoadDataURL != "" && props.LoadOnOpen {
		if err := s.loadModalData(ctx, props.LoadDataURL, state); err != nil {
			return nil, fmt.Errorf("failed to load initial data: %w", err)
		}
	}

	// Start workflow if specified
	if props.WorkflowID != "" {
		if err := s.startModalWorkflow(ctx, props, state); err != nil {
			return nil, fmt.Errorf("failed to start workflow: %w", err)
		}
	}

	// Add to modal stack if enabled
	if s.EnableModalStack && s.StateProvider != nil {
		if err := s.StateProvider.PushModal(ctx, sessionID, props.ID, props); err != nil {
			return nil, fmt.Errorf("failed to add to modal stack: %w", err)
		}
	}

	// Setup auto-save if enabled
	if props.AutoSave && s.EnableAutoSave {
		s.setupAutoSave(ctx, props.ID, state)
	}

	return state, nil
}

// CloseModal closes a modal and cleans up resources
func (s *ModalService) CloseModal(ctx context.Context, sessionID string, modalID string, forced bool) error {
	// Get current state
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return fmt.Errorf("failed to get modal state: %w", err)
	}

	// Check if can close (dirty state, validation, etc.)
	if !forced && state.Dirty {
		// This would typically trigger a confirmation dialog
		return fmt.Errorf("modal has unsaved changes")
	}

	// Stop auto-save
	if state.AutoSaveEnabled {
		s.stopAutoSave(ctx, modalID)
	}

	// Clean up file uploads
	if len(state.UploadQueue) > 0 {
		s.cancelPendingUploads(ctx, modalID)
	}

	// Complete workflow if active
	if state.InstanceID != "" {
		s.completeModalWorkflow(ctx, state)
	}

	// Remove from modal stack
	if s.EnableModalStack && s.StateProvider != nil {
		_, err := s.StateProvider.PopModal(ctx, sessionID)
		if err != nil {
			return fmt.Errorf("failed to remove from modal stack: %w", err)
		}
	}

	// Clean up cached data
	if s.CacheProvider != nil {
		s.CacheProvider.Delete(ctx, fmt.Sprintf("modal_state_%s", modalID))
	}

	return nil
}

// UpdateModalData updates modal data and triggers validation/auto-save
func (s *ModalService) UpdateModalData(ctx context.Context, modalID string, field string, value any) (*ValidationResult, error) {
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modal state: %w", err)
	}

	// Update form data
	oldValue := state.FormData[field]
	state.FormData[field] = value
	state.Touched[field] = true
	state.Dirty = true

	// Clear previous error for this field
	delete(state.Errors, field)

	// Validate field if validation provider is available
	var result *ValidationResult
	if s.ValidationProvider != nil {
		result, err = s.ValidationProvider.ValidateField(ctx, field, value, nil) // Rules would come from schema
		if err != nil {
			return nil, fmt.Errorf("validation failed: %w", err)
		}

		if !result.Valid {
			for field, error := range result.Errors {
				state.Errors[field] = error
			}
		}

		// Update overall validation state
		state.Valid = len(state.Errors) == 0
	}

	// Trigger auto-save if enabled and data is valid
	if state.AutoSaveEnabled && state.Valid {
		go s.triggerAutoSave(ctx, modalID, state)
	}

	// Save state
	if err := s.saveModalState(ctx, modalID, state); err != nil {
		return nil, fmt.Errorf("failed to save modal state: %w", err)
	}

	return result, nil
}

// SubmitModal submits modal form data and processes the result
func (s *ModalService) SubmitModal(ctx context.Context, modalID string, submitURL string) (*ModalSubmissionResult, error) {
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modal state: %w", err)
	}

	// Validate all data before submission
	if s.ValidationProvider != nil {
		// Get form schema from modal props (would need to be passed in or stored)
		result, err := s.ValidationProvider.ValidateForm(ctx, state.FormData, nil)
		if err != nil {
			return nil, fmt.Errorf("validation failed: %w", err)
		}

		if !result.Valid {
			state.Errors = result.Errors
			state.Valid = false
			s.saveModalState(ctx, modalID, state)
			return &ModalSubmissionResult{
				Success: false,
				Errors:  result.Errors,
			}, nil
		}
	}

	// Mark as saving
	state.Saving = true
	s.saveModalState(ctx, modalID, state)

	// Save data using data provider
	if s.DataProvider != nil {
		if err := s.DataProvider.Save(ctx, modalID, state.FormData); err != nil {
			state.Saving = false
			s.saveModalState(ctx, modalID, state)
			return nil, fmt.Errorf("failed to save data: %w", err)
		}
	}

	// Mark as clean and not saving
	state.Dirty = false
	state.Saving = false
	state.SaveCount++
	now := time.Now()
	state.LastSaved = &now

	s.saveModalState(ctx, modalID, state)

	return &ModalSubmissionResult{
		Success: true,
		Data:    state.FormData,
	}, nil
}

// ModalSubmissionResult represents the result of modal submission
type ModalSubmissionResult struct {
	Success     bool              `json:"success"`
	Data        map[string]any    `json:"data,omitempty"`
	Errors      map[string]string `json:"errors,omitempty"`
	Message     string            `json:"message,omitempty"`
	RedirectURL string            `json:"redirectUrl,omitempty"`
}

// Step navigation methods

// NextStep moves to the next step in multi-step modal
func (s *ModalService) NextStep(ctx context.Context, modalID string) error {
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return fmt.Errorf("failed to get modal state: %w", err)
	}

	if state.CurrentStep >= state.TotalSteps {
		return fmt.Errorf("already at last step")
	}

	// Validate current step before proceeding
	if s.ValidationProvider != nil {
		// Validation logic for current step
		// This would validate only the fields relevant to the current step
	}

	state.CurrentStep++
	state.CompletedSteps = append(state.CompletedSteps, state.CurrentStep-1)

	return s.saveModalState(ctx, modalID, state)
}

// PrevStep moves to the previous step in multi-step modal
func (s *ModalService) PrevStep(ctx context.Context, modalID string) error {
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return fmt.Errorf("failed to get modal state: %w", err)
	}

	if state.CurrentStep <= 1 {
		return fmt.Errorf("already at first step")
	}

	state.CurrentStep--

	return s.saveModalState(ctx, modalID, state)
}

// GoToStep jumps to a specific step
func (s *ModalService) GoToStep(ctx context.Context, modalID string, step int) error {
	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return fmt.Errorf("failed to get modal state: %w", err)
	}

	if step < 1 || step > state.TotalSteps {
		return fmt.Errorf("invalid step number: %d", step)
	}

	state.CurrentStep = step

	return s.saveModalState(ctx, modalID, state)
}

// File upload methods

// UploadFile handles file upload for modals
func (s *ModalService) UploadFile(ctx context.Context, modalID string, file *UploadedFile) (*FileMetadata, error) {
	if !s.EnableFileUpload || s.FileProvider == nil {
		return nil, fmt.Errorf("file upload not enabled")
	}

	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modal state: %w", err)
	}

	// Create upload task
	task := &UploadTask{
		ID:       generateUploadTaskID(),
		File:     file,
		Progress: 0,
		Status:   "pending",
	}

	state.UploadQueue = append(state.UploadQueue, *task)
	s.saveModalState(ctx, modalID, state)

	// Perform upload
	task.Status = "uploading"
	metadata, err := s.FileProvider.Upload(ctx, file)
	if err != nil {
		task.Status = "failed"
		task.Error = err.Error()
		s.saveModalState(ctx, modalID, state)
		return nil, fmt.Errorf("upload failed: %w", err)
	}

	// Update state
	task.Status = "completed"
	task.Progress = 100
	state.UploadedFiles = append(state.UploadedFiles, *metadata)

	// Remove from queue
	for i, queueTask := range state.UploadQueue {
		if queueTask.ID == task.ID {
			state.UploadQueue = append(state.UploadQueue[:i], state.UploadQueue[i+1:]...)
			break
		}
	}

	state.Dirty = true
	s.saveModalState(ctx, modalID, state)

	return metadata, nil
}

// Search methods

// Search performs search within modal context
func (s *ModalService) Search(ctx context.Context, modalID string, query string, filters map[string]any) ([]SearchResult, error) {
	if !s.EnableSearch || s.SearchProvider == nil {
		return nil, fmt.Errorf("search not enabled")
	}

	state, err := s.getModalState(ctx, modalID)
	if err != nil {
		return nil, fmt.Errorf("failed to get modal state: %w", err)
	}

	// Update search state
	state.SearchQuery = query
	state.SearchLoading = true
	s.saveModalState(ctx, modalID, state)

	// Perform search
	results, err := s.SearchProvider.Search(ctx, query, filters)
	if err != nil {
		state.SearchLoading = false
		s.saveModalState(ctx, modalID, state)
		return nil, fmt.Errorf("search failed: %w", err)
	}

	// Update state with results
	state.SearchResults = results
	state.SearchLoading = false
	s.saveModalState(ctx, modalID, state)

	return results, nil
}

// Helper methods

func (s *ModalService) getModalState(ctx context.Context, modalID string) (*ModalState, error) {
	// Try cache first
	if s.CacheProvider != nil {
		data, err := s.CacheProvider.Get(ctx, fmt.Sprintf("modal_state_%s", modalID))
		if err == nil {
			var state ModalState
			if err := json.Unmarshal(data, &state); err == nil {
				return &state, nil
			}
		}
	}

	// Fallback to creating new state
	return &ModalState{
		FormData:    make(map[string]any),
		Errors:      make(map[string]string),
		Touched:     make(map[string]bool),
		CustomState: make(map[string]any),
	}, nil
}

func (s *ModalService) saveModalState(ctx context.Context, modalID string, state *ModalState) error {
	if s.CacheProvider != nil {
		data, err := json.Marshal(state)
		if err != nil {
			return fmt.Errorf("failed to marshal state: %w", err)
		}

		if err := s.CacheProvider.Set(ctx, fmt.Sprintf("modal_state_%s", modalID), data, s.DefaultCacheTimeout); err != nil {
			return fmt.Errorf("failed to cache state: %w", err)
		}
	}

	return nil
}

func (s *ModalService) loadModalData(ctx context.Context, url string, state *ModalState) error {
	// Implementation would make HTTP request to load data
	// This is a placeholder for the actual implementation
	return nil
}

func (s *ModalService) startModalWorkflow(ctx context.Context, props *EnhancedModalProps, state *ModalState) error {
	if s.WorkflowProvider == nil {
		return nil
	}

	instance, err := s.WorkflowProvider.StartWorkflow(ctx, props.WorkflowID, state.FormData)
	if err != nil {
		return err
	}

	state.WorkflowID = props.WorkflowID
	state.InstanceID = instance.ID
	state.WorkflowState = instance.State

	return nil
}

func (s *ModalService) completeModalWorkflow(ctx context.Context, state *ModalState) error {
	if s.WorkflowProvider == nil || state.InstanceID == "" {
		return nil
	}

	return s.WorkflowProvider.CompleteWorkflow(ctx, state.InstanceID, state.FormData)
}

func (s *ModalService) setupAutoSave(ctx context.Context, modalID string, state *ModalState) {
	// Implementation would setup auto-save timer
	// This is a placeholder for the actual implementation
}

func (s *ModalService) stopAutoSave(ctx context.Context, modalID string) {
	// Implementation would stop auto-save timer
}

func (s *ModalService) triggerAutoSave(ctx context.Context, modalID string, state *ModalState) {
	// Implementation would perform auto-save
}

func (s *ModalService) cancelPendingUploads(ctx context.Context, modalID string) {
	// Implementation would cancel pending uploads
}

func getUserFromContext(ctx context.Context) *User {
	if user, ok := ctx.Value("user").(*User); ok {
		return user
	}
	return nil
}

func generateUploadTaskID() string {
	return fmt.Sprintf("upload_%d", time.Now().UnixNano())
}

// Class generation functions for enhanced modal components

func getEnhancedModalOverlayClasses(props EnhancedModalProps) string {
	classes := []string{}

	// Base classes
	classes = append(classes, "enhanced-modal-overlay", "fixed", "inset-0", "z-50")

	// Type-specific positioning
	switch props.Type {
	case EnhancedModalDialog:
		classes = append(classes, "flex", "items-center", "justify-center", "p-4")
		if props.Centered {
			classes = append(classes, "items-center")
		} else {
			classes = append(classes, "items-start", "pt-16")
		}

	case EnhancedModalDrawer:
		classes = append(classes, "flex")
		switch props.Position {
		case EnhancedModalPositionRight:
			classes = append(classes, "justify-end")
		case EnhancedModalPositionLeft:
			classes = append(classes, "justify-start")
		case EnhancedModalPositionTop:
			classes = append(classes, "items-start")
		case EnhancedModalPositionBottom:
			classes = append(classes, "items-end")
		}

	case EnhancedModalBottomSheet:
		classes = append(classes, "flex", "items-end")

	case EnhancedModalFullscreen:
		// No flex needed for fullscreen

	case EnhancedModalPopover:
		classes = append(classes, "flex")
		// Position would be set dynamically based on trigger element
	}

	// Backdrop
	if props.Backdrop {
		if props.BlurBackdrop {
			classes = append(classes, "bg-black/80", "backdrop-blur-sm")
		} else {
			classes = append(classes, "bg-black/80")
		}
	}

	// Mobile responsiveness
	if props.MobileFullscreen {
		classes = append(classes, "sm:flex", "sm:items-center", "sm:justify-center", "sm:p-4")
	}

	// Custom classes
	if props.OverlayClass != "" {
		classes = append(classes, props.OverlayClass)
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getEnhancedModalContentClasses(props EnhancedModalProps) string {
	classes := []string{}

	// Base classes
	classes = append(classes, "enhanced-modal-content", "relative", "bg-background", "text-foreground", "shadow-lg", "focus:outline-none")

	// Type-specific styles
	switch props.Type {
	case EnhancedModalDialog, EnhancedModalPopover:
		classes = append(classes, "rounded-lg", "border")

	case EnhancedModalDrawer:
		switch props.Position {
		case EnhancedModalPositionRight:
			classes = append(classes, "h-full", "border-l", "rounded-l-lg")
		case EnhancedModalPositionLeft:
			classes = append(classes, "h-full", "border-r", "rounded-r-lg")
		case EnhancedModalPositionTop:
			classes = append(classes, "w-full", "border-b", "rounded-b-lg")
		case EnhancedModalPositionBottom:
			classes = append(classes, "w-full", "border-t", "rounded-t-lg")
		default:
			classes = append(classes, "h-full", "border-l", "rounded-l-lg")
		}

	case EnhancedModalBottomSheet:
		classes = append(classes, "w-full", "border-t", "rounded-t-lg")
		if props.MobileFullscreen {
			classes = append(classes, "h-full", "sm:h-auto", "sm:max-h-[90vh]", "sm:rounded-lg", "sm:border")
		}

	case EnhancedModalFullscreen:
		classes = append(classes, "w-screen", "h-screen")
	}

	// Size classes
	if props.Type != EnhancedModalFullscreen && props.Type != EnhancedModalBottomSheet {
		switch props.Size {
		case EnhancedModalSizeXS:
			classes = append(classes, "w-full", "max-w-xs")
		case EnhancedModalSizeSM:
			classes = append(classes, "w-full", "max-w-sm")
		case EnhancedModalSizeMD:
			classes = append(classes, "w-full", "max-w-md")
		case EnhancedModalSizeLG:
			classes = append(classes, "w-full", "max-w-lg")
		case EnhancedModalSizeXL:
			classes = append(classes, "w-full", "max-w-xl")
		case EnhancedModalSize2XL:
			classes = append(classes, "w-full", "max-w-2xl")
		case EnhancedModalSize3XL:
			classes = append(classes, "w-full", "max-w-3xl")
		case EnhancedModalSizeAuto:
			classes = append(classes, "w-auto")
		case EnhancedModalSizeFull:
			if props.Type == EnhancedModalDrawer {
				switch props.Position {
				case EnhancedModalPositionRight, EnhancedModalPositionLeft:
					classes = append(classes, "w-full", "sm:w-96", "md:w-[32rem]")
				case EnhancedModalPositionTop, EnhancedModalPositionBottom:
					classes = append(classes, "h-64", "sm:h-80")
				}
			} else {
				classes = append(classes, "w-full", "max-w-6xl")
			}
		default:
			classes = append(classes, "w-full", "max-w-md")
		}
	}

	// Drawer specific width for different positions
	if props.Type == EnhancedModalDrawer && props.Size != EnhancedModalSizeFull {
		switch props.Position {
		case EnhancedModalPositionRight, EnhancedModalPositionLeft:
			switch props.Size {
			case EnhancedModalSizeXS:
				classes = append(classes, "w-64")
			case EnhancedModalSizeSM:
				classes = append(classes, "w-80")
			case EnhancedModalSizeMD:
				classes = append(classes, "w-96")
			case EnhancedModalSizeLG:
				classes = append(classes, "w-[28rem]")
			case EnhancedModalSizeXL:
				classes = append(classes, "w-[32rem]")
			default:
				classes = append(classes, "w-96")
			}
		case EnhancedModalPositionTop, EnhancedModalPositionBottom:
			switch props.Size {
			case EnhancedModalSizeXS:
				classes = append(classes, "h-32")
			case EnhancedModalSizeSM:
				classes = append(classes, "h-48")
			case EnhancedModalSizeMD:
				classes = append(classes, "h-64")
			case EnhancedModalSizeLG:
				classes = append(classes, "h-80")
			case EnhancedModalSizeXL:
				classes = append(classes, "h-96")
			default:
				classes = append(classes, "h-64")
			}
		}
	}

	// Scrollable content
	if props.Scrollable {
		classes = append(classes, "max-h-[90vh]", "flex", "flex-col")
	}

	// Theme variant classes
	switch props.Variant {
	case EnhancedModalDestructive:
		classes = append(classes, "border-destructive/20")
	case EnhancedModalSuccess:
		classes = append(classes, "border-success/20")
	case EnhancedModalWarning:
		classes = append(classes, "border-warning/20")
	case EnhancedModalInfo:
		classes = append(classes, "border-info/20")
	}

	// Custom content classes
	if props.ContentClass != "" {
		classes = append(classes, props.ContentClass)
	}

	// Base custom class
	if props.Class != "" {
		classes = append(classes, props.Class)
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getEnhancedModalHeaderClasses(props EnhancedModalProps) string {
	classes := []string{"enhanced-modal-header", "flex", "items-start", "justify-between", "p-6"}

	// Variant-specific styling
	switch props.Variant {
	case EnhancedModalDestructive:
		classes = append(classes, "border-b", "border-destructive/20")
	case EnhancedModalSuccess:
		classes = append(classes, "border-b", "border-success/20")
	case EnhancedModalWarning:
		classes = append(classes, "border-b", "border-warning/20")
	case EnhancedModalInfo:
		classes = append(classes, "border-b", "border-info/20")
	default:
		classes = append(classes, "border-b")
	}

	// Fullscreen header styling
	if props.Type == EnhancedModalFullscreen {
		classes = append(classes, "bg-background/95", "backdrop-blur", "border-b", "sticky", "top-0", "z-10")
	}

	// Custom header classes
	if props.HeaderClass != "" {
		classes = append(classes, props.HeaderClass)
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getEnhancedModalBodyClasses(props EnhancedModalProps) string {
	classes := []string{"enhanced-modal-body", "p-6"}

	// Scrollable content
	if props.Scrollable {
		classes = append(classes, "overflow-y-auto", "flex-1")
	}

	// Type-specific body styling
	if props.Type == EnhancedModalFullscreen {
		classes = append(classes, "container", "mx-auto", "max-w-6xl")
	}

	// Custom body classes
	if props.BodyClass != "" {
		classes = append(classes, props.BodyClass)
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getEnhancedModalFooterClasses(props EnhancedModalProps) string {
	classes := []string{"enhanced-modal-footer", "flex", "items-center", "justify-between", "gap-2", "p-6", "border-t"}

	// Fullscreen footer styling
	if props.Type == EnhancedModalFullscreen {
		classes = append(classes, "bg-background/95", "backdrop-blur", "sticky", "bottom-0", "z-10")
	}

	// Mobile optimization
	classes = append(classes, "flex-col", "sm:flex-row")

	// Custom footer classes
	if props.FooterClass != "" {
		classes = append(classes, props.FooterClass)
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

// Step navigation classes

func getStepNavItemClasses(step EnhancedModalStep, index int, props EnhancedModalProps) string {
	classes := []string{"step-nav-item", "flex", "items-center", "text-sm"}

	if step.Active {
		classes = append(classes, "text-primary")
	} else if step.Completed {
		classes = append(classes, "text-success")
	} else {
		classes = append(classes, "text-muted-foreground")
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getStepNavButtonClasses(step EnhancedModalStep, index int, props EnhancedModalProps) string {
	classes := []string{"step-nav-button"}

	if props.StepNavigation && !props.LinearProgress {
		classes = append(classes, "cursor-pointer", "transition-colors", "hover:text-primary")
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

func getStepIconContainerClasses(step EnhancedModalStep, index int, props EnhancedModalProps) string {
	classes := []string{"flex", "items-center", "justify-center", "w-8", "h-8", "rounded-full", "border-2"}

	if step.Active {
		classes = append(classes, "bg-primary", "border-primary", "text-primary-foreground")
	} else if step.Completed {
		classes = append(classes, "bg-success", "border-success", "text-success-foreground")
	} else {
		classes = append(classes, "border-muted-foreground", "text-muted-foreground")
	}

	return utils.TwMerge(strings.Join(classes, " "))
}

// Action helper functions

func getActionsByPosition(actions []EnhancedModalAction, position string) []EnhancedModalAction {
	var filtered []EnhancedModalAction
	for _, action := range actions {
		actionPosition := utils.IfElse(action.Position != "", action.Position, "right")
		if actionPosition == position {
			filtered = append(filtered, action)
		}
	}

	// Sort by priority
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].Priority < filtered[j].Priority
	})

	return filtered
}

func getEnhancedModalActionVariant(action EnhancedModalAction) atoms.ButtonVariant {
	if action.Variant != "" {
		return action.Variant
	}

	// Default variants based on action type
	switch action.Type {
	case "primary", "submit", "save":
		return atoms.ButtonPrimary
	case "destructive", "delete":
		return atoms.ButtonDestructive
	case "secondary", "cancel":
		return atoms.ButtonSecondary
	default:
		return atoms.ButtonOutline
	}
}

func getEnhancedModalActionSize(action EnhancedModalAction) atoms.ButtonSize {
	if action.Size != "" {
		return action.Size
	}
	return atoms.ButtonSizeMD
}

func getEnhancedModalActionType(action EnhancedModalAction) string {
	switch action.Type {
	case "submit":
		return "submit"
	case "reset":
		return "reset"
	default:
		return "button"
	}
}

func getEnhancedModalActionClasses(action EnhancedModalAction) string {
	classes := []string{}

	if action.FullWidth {
		classes = append(classes, "w-full", "sm:w-auto")
	}

	return strings.Join(classes, " ")
}

func getEnhancedModalActionTarget(action EnhancedModalAction, modalProps EnhancedModalProps) string {
	if action.HXTarget != "" {
		return action.HXTarget
	}
	return modalProps.HXTarget
}

func getEnhancedModalActionSwap(action EnhancedModalAction, modalProps EnhancedModalProps) string {
	if action.HXSwap != "" {
		return action.HXSwap
	}
	if modalProps.HXSwap != "" {
		return modalProps.HXSwap
	}
	return "innerHTML"
}

func getEnhancedModalActionClick(action EnhancedModalAction, modalProps EnhancedModalProps) string {
	clickHandler := action.AlpineClick

	// Add confirmation if specified
	if action.ConfirmText != "" {
		confirmHandler := fmt.Sprintf("if(confirm('%s'))", action.ConfirmText)
		if clickHandler != "" {
			clickHandler = fmt.Sprintf("%s && (%s)", confirmHandler, clickHandler)
		} else {
			clickHandler = confirmHandler
		}
	}

	// Add auto-close if specified
	if action.AutoClose {
		if clickHandler != "" {
			clickHandler = fmt.Sprintf("(%s) && closeModal()", clickHandler)
		} else {
			clickHandler = "closeModal()"
		}
	}

	// Add default actions based on type
	if clickHandler == "" {
		switch action.Type {
		case "submit":
			clickHandler = "submitForm()"
		case "cancel":
			clickHandler = "closeModal()"
		case "save":
			clickHandler = "saveForm()"
		}
	}

	return clickHandler
}

// Transition helper functions

func getModalEnterTransition(props EnhancedModalProps) string {
	duration := utils.IfElse(props.AnimationDuration > 0, props.AnimationDuration, 300)
	return fmt.Sprintf("ease-out duration-%d", duration)
}

func getModalEnterStartTransition(props EnhancedModalProps) string {
	switch props.Type {
	case EnhancedModalDrawer:
		switch props.Position {
		case EnhancedModalPositionRight:
			return "opacity-0 translate-x-full"
		case EnhancedModalPositionLeft:
			return "opacity-0 -translate-x-full"
		case EnhancedModalPositionTop:
			return "opacity-0 -translate-y-full"
		case EnhancedModalPositionBottom:
			return "opacity-0 translate-y-full"
		}
	case EnhancedModalBottomSheet:
		return "opacity-0 translate-y-full"
	case EnhancedModalPopover:
		return "opacity-0 scale-95"
	default:
		return "opacity-0"
	}
	return "opacity-0"
}

func getModalEnterEndTransition(props EnhancedModalProps) string {
	return "opacity-100 translate-x-0 translate-y-0 scale-100"
}

func getModalLeaveTransition(props EnhancedModalProps) string {
	duration := utils.IfElse(props.AnimationDuration > 0, props.AnimationDuration/2, 200)
	return fmt.Sprintf("ease-in duration-%d", duration)
}

func getModalLeaveStartTransition(props EnhancedModalProps) string {
	return "opacity-100 translate-x-0 translate-y-0 scale-100"
}

func getModalLeaveEndTransition(props EnhancedModalProps) string {
	return getModalEnterStartTransition(props) // Same as enter start
}

func getModalShowExpression(props EnhancedModalProps) string {
	if props.Name != "" {
		return props.Name + "Open"
	}
	return "open"
}

// Header icon classes based on variant
func getHeaderIconClasses(variant EnhancedModalVariant) string {
	switch variant {
	case EnhancedModalDestructive:
		return "text-destructive"
	case EnhancedModalSuccess:
		return "text-success"
	case EnhancedModalWarning:
		return "text-warning"
	case EnhancedModalInfo:
		return "text-info"
	default:
		return "text-muted-foreground"
	}
}

// Header title classes based on variant
func getHeaderTitleClasses(variant EnhancedModalVariant) string {
	classes := []string{"text-lg", "font-semibold"}

	switch variant {
	case EnhancedModalDestructive:
		classes = append(classes, "text-destructive")
	case EnhancedModalSuccess:
		classes = append(classes, "text-success")
	case EnhancedModalWarning:
		classes = append(classes, "text-warning")
	case EnhancedModalInfo:
		classes = append(classes, "text-info")
	default:
		classes = append(classes, "text-foreground")
	}

	return strings.Join(classes, " ")
}

// Alpine.js data generation for enhanced modal state management
func getEnhancedModalAlpineData(props EnhancedModalProps) string {
	initialData := props.FormData
	if initialData == nil {
		initialData = make(map[string]any)
	}

	initialDataJSON, _ := json.Marshal(initialData)

	totalSteps := len(props.Steps)
	if totalSteps == 0 {
		totalSteps = 1
	}

	autoSaveInterval := utils.IfElse(props.AutoSaveInterval > 0, props.AutoSaveInterval, 30)

	return fmt.Sprintf(`{
		// Core modal state
		open: %s,
		loading: %s,
		saving: %s,
		dirty: false,
		valid: true,
		
		// Modal type and configuration
		modalType: '%s',
		modalSize: '%s',
		modalPosition: '%s',
		modalVariant: '%s',
		
		// Multi-step state
		currentStep: %d,
		totalSteps: %d,
		completedSteps: [],
		skippedSteps: [],
		showStepNavigation: %s,
		linearProgress: %s,
		
		// Form data and validation
		formData: %s,
		initialData: %s,
		errors: {},
		touched: {},
		validationState: 'idle',
		
		// File upload state
		files: [],
		uploadQueue: [],
		uploadProgress: {},
		isDragOver: false,
		maxFiles: %d,
		maxFileSize: %d,
		acceptedTypes: '%s',
		
		// Search state
		searchQuery: '',
		searchResults: [],
		searchLoading: false,
		selectedItems: [],
		multiSelect: %s,
		
		// Auto-save state
		autoSaveEnabled: %s,
		autoSaveInterval: %d,
		lastSaved: null,
		autoSaveTimer: null,
		saveCount: 0,
		
		// Modal stack state
		stackDepth: 0,
		maxStackDepth: %d,
		modalStack: %s,
		
		// Workflow state
		workflowId: '%s',
		processId: '%s',
		instanceId: '%s',
		workflowState: 'idle',
		
		// UI state
		headerVisible: %s,
		footerVisible: %s,
		progressVisible: %s,
		
		// Mobile and responsive state
		isMobile: false,
		isTablet: false,
		mobileFullscreen: %s,
		
		// Accessibility state
		focusTrap: %s,
		restoreFocus: %s,
		
		// Computed properties
		get progressPercentage() {
			return this.totalSteps > 1 ? Math.round((this.currentStep / this.totalSteps) * 100) : 100;
		},
		
		get progressStyle() {
			return 'width: ' + this.progressPercentage + '%%';
		},
		
		get canGoNext() {
			return this.currentStep < this.totalSteps;
		},
		
		get canGoPrev() {
			return this.currentStep > 1;
		},
		
		get hasErrors() {
			return Object.keys(this.errors).length > 0;
		},
		
		get isLastStep() {
			return this.currentStep === this.totalSteps;
		},
		
		get isFirstStep() {
			return this.currentStep === 1;
		},
		
		get uploadedFilesCount() {
			return this.files.filter(f => f.status === 'completed').length;
		},
		
		get totalFileSize() {
			return this.files.reduce((sum, file) => sum + (file.size || 0), 0);
		},
		
		get hasSelectedItems() {
			return this.selectedItems.length > 0;
		},
		
		// Core modal methods
		initEnhancedModal() {
			this.setupResponsive();
			this.setupAutoSave();
			this.setupKeyboardNavigation();
			this.setupFileUpload();
			this.setupValidation();
			this.loadInitialData();
			%s // Custom initialization
		},
		
		openModal() {
			this.open = true;
			document.body.style.overflow = 'hidden';
			this.restoreFocusElement = document.activeElement;
			
			this.$nextTick(() => {
				if (this.autoFocus) {
					const firstInput = this.$el.querySelector('input, select, textarea');
					if (firstInput) firstInput.focus();
				}
			});
			
			%s // Custom onOpen handler
		},
		
		closeModal() {
			if (this.dirty && !confirm('You have unsaved changes. Are you sure you want to close?')) {
				return false;
			}
			
			%s // Custom onBeforeClose handler
			
			this.open = false;
			document.body.style.overflow = '';
			
			if (this.restoreFocus && this.restoreFocusElement) {
				this.restoreFocusElement.focus();
			}
			
			this.cleanupModal();
			%s // Custom onClose handler
			
			return true;
		},
		
		closeModalOnOverlay(event) {
			if (event.target === event.currentTarget) {
				this.closeModal();
			}
		},
		
		cleanupModal() {
			this.stopAutoSave();
			this.cancelPendingUploads();
			this.clearErrors();
		},
		
		// Step navigation
		nextStep() {
			if (!this.canGoNext) return;
			
			if (this.linearProgress && !this.validateCurrentStep()) {
				return;
			}
			
			this.currentStep++;
			if (!this.completedSteps.includes(this.currentStep - 1)) {
				this.completedSteps.push(this.currentStep - 1);
			}
			
			%s // Custom onStepChange handler
		},
		
		prevStep() {
			if (!this.canGoPrev) return;
			this.currentStep--;
			%s // Custom onStepChange handler
		},
		
		goToStep(step) {
			if (step >= 1 && step <= this.totalSteps) {
				if (!this.linearProgress || this.completedSteps.includes(step - 1)) {
					this.currentStep = step;
					%s // Custom onStepChange handler
				}
			}
		},
		
		validateCurrentStep() {
			// Implementation would validate current step fields
			return this.valid;
		},
		
		// Form data management
		updateField(field, value) {
			const oldValue = this.formData[field];
			this.formData[field] = value;
			this.touched[field] = true;
			this.dirty = value !== (this.initialData[field] || '');
			
			// Clear field error
			delete this.errors[field];
			
			// Trigger validation
			this.validateField(field, value);
			
			// Reset auto-save timer
			if (this.autoSaveEnabled) {
				this.resetAutoSaveTimer();
			}
			
			%s // Custom onDataChange handler
		},
		
		validateField(field, value) {
			// Implementation would validate individual field
			%s // Custom onValidation handler
		},
		
		validateForm() {
			this.valid = Object.keys(this.errors).length === 0;
			return this.valid;
		},
		
		clearErrors() {
			this.errors = {};
			this.valid = true;
		},
		
		resetForm() {
			this.formData = { ...this.initialData };
			this.errors = {};
			this.touched = {};
			this.dirty = false;
			this.valid = true;
			this.currentStep = 1;
			this.completedSteps = [];
		},
		
		// Form submission
		submitForm() {
			this.saving = true;
			
			if (!this.validateForm()) {
				this.saving = false;
				return;
			}
			
			// Implementation would submit form data
			// This is handled by HTMX or custom submit handler
			%s // Custom submit handler
		},
		
		saveForm() {
			this.saving = true;
			// Implementation would save form data without submission
			this.performAutoSave();
		},
		
		// Auto-save functionality
		setupAutoSave() {
			if (!this.autoSaveEnabled) return;
			
			this.autoSaveTimer = setInterval(() => {
				if (this.dirty && this.valid && !this.saving) {
					this.performAutoSave();
				}
			}, this.autoSaveInterval * 1000);
		},
		
		performAutoSave() {
			if (!this.dirty || this.saving) return;
			
			// Implementation would save form data
			fetch('/api/modal/auto-save', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-Requested-With': 'XMLHttpRequest'
				},
				body: JSON.stringify({
					modalId: '%s',
					data: this.formData
				})
			})
			.then(response => response.json())
			.then(data => {
				if (data.success) {
					this.lastSaved = new Date();
					this.saveCount++;
					this.dirty = false;
				}
			})
			.catch(error => {
				console.error('Auto-save failed:', error);
			});
			
			%s // Custom auto-save handler
		},
		
		resetAutoSaveTimer() {
			if (this.autoSaveTimer) {
				clearInterval(this.autoSaveTimer);
				this.setupAutoSave();
			}
		},
		
		stopAutoSave() {
			if (this.autoSaveTimer) {
				clearInterval(this.autoSaveTimer);
				this.autoSaveTimer = null;
			}
		},
		
		// File upload methods
		setupFileUpload() {
			this.files = [];
			this.uploadProgress = {};
		},
		
		handleDrop(event) {
			this.isDragOver = false;
			const files = Array.from(event.dataTransfer.files);
			this.addFiles(files);
		},
		
		handleFileSelect(event) {
			const files = Array.from(event.target.files);
			this.addFiles(files);
			event.target.value = ''; // Reset input
		},
		
		addFiles(files) {
			for (const file of files) {
				if (this.validateFile(file)) {
					this.files.push({
						id: this.generateFileId(),
						name: file.name,
						size: file.size,
						type: file.type,
						file: file,
						status: 'pending',
						progress: 0
					});
					
					this.uploadFile(this.files[this.files.length - 1]);
				}
			}
			
			this.dirty = true;
		},
		
		validateFile(file) {
			if (this.maxFiles > 0 && this.files.length >= this.maxFiles) {
				alert('Maximum number of files exceeded');
				return false;
			}
			
			if (this.maxFileSize > 0 && file.size > this.maxFileSize * 1024 * 1024) {
				alert('File size exceeds maximum limit');
				return false;
			}
			
			if (this.acceptedTypes && !this.acceptedTypes.split(',').some(type => 
				file.type.includes(type.trim()) || file.name.includes(type.trim())
			)) {
				alert('File type not supported');
				return false;
			}
			
			return true;
		},
		
		uploadFile(fileObj) {
			const formData = new FormData();
			formData.append('file', fileObj.file);
			formData.append('modalId', '%s');
			
			fileObj.status = 'uploading';
			
			fetch('/api/modal/upload', {
				method: 'POST',
				body: formData
			})
			.then(response => response.json())
			.then(data => {
				if (data.success) {
					fileObj.status = 'completed';
					fileObj.progress = 100;
					fileObj.url = data.url;
					fileObj.id = data.id;
				} else {
					fileObj.status = 'failed';
					fileObj.error = data.error || 'Upload failed';
				}
			})
			.catch(error => {
				fileObj.status = 'failed';
				fileObj.error = error.message;
			});
			
			%s // Custom upload progress handler
		},
		
		removeFile(index) {
			const file = this.files[index];
			if (file.status === 'uploading') {
				// Cancel upload if possible
			}
			
			this.files.splice(index, 1);
			this.dirty = true;
		},
		
		cancelPendingUploads() {
			this.files = this.files.filter(f => f.status === 'completed');
		},
		
		formatFileSize(bytes) {
			if (bytes === 0) return '0 Bytes';
			const k = 1024;
			const sizes = ['Bytes', 'KB', 'MB', 'GB'];
			const i = Math.floor(Math.log(bytes) / Math.log(k));
			return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
		},
		
		generateFileId() {
			return 'file_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9);
		},
		
		// Search functionality
		debounceSearch() {
			clearTimeout(this.searchTimeout);
			this.searchTimeout = setTimeout(() => {
				this.performSearch();
			}, %d);
		},
		
		performSearch() {
			if (this.searchQuery.length < 2) {
				this.searchResults = [];
				return;
			}
			
			this.searchLoading = true;
			
			fetch('/api/modal/search', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					'X-Requested-With': 'XMLHttpRequest'
				},
				body: JSON.stringify({
					query: this.searchQuery,
					modalId: '%s'
				})
			})
			.then(response => response.json())
			.then(data => {
				this.searchResults = data.results || [];
			})
			.catch(error => {
				console.error('Search failed:', error);
				this.searchResults = [];
			})
			.finally(() => {
				this.searchLoading = false;
			});
			
			%s // Custom search handler
		},
		
		toggleSelection(item) {
			const index = this.selectedItems.indexOf(item.id);
			if (index > -1) {
				this.selectedItems.splice(index, 1);
			} else {
				if (this.multiSelect) {
					this.selectedItems.push(item.id);
				} else {
					this.selectedItems = [item.id];
				}
			}
			
			this.dirty = true;
			%s // Custom selection handler
		},
		
		isSelected(item) {
			return this.selectedItems.includes(item.id);
		},
		
		clearSelection() {
			this.selectedItems = [];
			this.dirty = true;
		},
		
		// Data loading
		loadInitialData() {
			if ('%s' && %s) {
				this.loading = true;
				
				fetch('%s')
				.then(response => response.json())
				.then(data => {
					this.formData = { ...this.formData, ...data };
					this.initialData = { ...this.initialData, ...data };
				})
				.catch(error => {
					console.error('Failed to load initial data:', error);
				})
				.finally(() => {
					this.loading = false;
				});
			}
		},
		
		// Responsive behavior
		setupResponsive() {
			this.updateResponsiveState();
			window.addEventListener('resize', () => {
				this.updateResponsiveState();
			});
		},
		
		updateResponsiveState() {
			this.isMobile = window.innerWidth < 768;
			this.isTablet = window.innerWidth >= 768 && window.innerWidth < 1024;
		},
		
		// Keyboard navigation
		setupKeyboardNavigation() {
			document.addEventListener('keydown', (event) => {
				if (!this.open) return;
				
				switch (event.key) {
					case 'Escape':
						if ('%s' === 'true') {
							this.closeModal();
						}
						break;
					case 'Enter':
						if (event.ctrlKey || event.metaKey) {
							if ('%s' === 'true') {
								this.submitForm();
							}
						}
						break;
					case 'Tab':
						// Focus trap logic would go here
						break;
				}
			});
		},
		
		// Utility methods
		generateId() {
			return 'modal_' + Date.now() + '_' + Math.random().toString(36).substr(2, 9);
		},
		
		log(message, data = null) {
			if ('%s' === 'true') {
				console.log('[Enhanced Modal]', message, data);
			}
		}
	}`,
		strconv.FormatBool(props.Open),
		strconv.FormatBool(props.Loading),
		strconv.FormatBool(props.Saving),
		props.Type,
		props.Size,
		props.Position,
		props.Variant,
		props.CurrentStep,
		totalSteps,
		strconv.FormatBool(props.ShowStepNav),
		strconv.FormatBool(props.LinearProgress),
		string(initialDataJSON),
		string(initialDataJSON),
		props.UploadMaxFiles,
		props.UploadMaxSize,
		props.UploadAccept,
		strconv.FormatBool(props.MultiSelect),
		strconv.FormatBool(props.AutoSave),
		autoSaveInterval,
		props.MaxStackDepth,
		strconv.FormatBool(props.ModalStack),
		props.WorkflowID,
		props.ProcessID,
		props.EntityID,
		strconv.FormatBool(props.Header),
		strconv.FormatBool(props.Footer),
		strconv.FormatBool(props.ShowProgress),
		strconv.FormatBool(props.MobileFullscreen),
		strconv.FormatBool(props.TrapFocus),
		strconv.FormatBool(props.RestoreFocus),
		utils.IfElse(props.AlpineInit != "", props.AlpineInit, "// No custom initialization"),
		utils.IfElse(props.OnOpen != "", props.OnOpen, "// No onOpen handler"),
		utils.IfElse(props.OnBeforeClose != "", props.OnBeforeClose, "// No onBeforeClose handler"),
		utils.IfElse(props.OnClose != "", props.OnClose, "// No onClose handler"),
		utils.IfElse(props.OnStepChange != "", props.OnStepChange, "// No onStepChange handler"),
		utils.IfElse(props.OnStepChange != "", props.OnStepChange, "// No onStepChange handler"),
		utils.IfElse(props.OnStepChange != "", props.OnStepChange, "// No onStepChange handler"),
		utils.IfElse(props.OnDataChange != "", props.OnDataChange, "// No onDataChange handler"),
		utils.IfElse(props.OnValidation != "", props.OnValidation, "// No onValidation handler"),
		utils.IfElse(props.OnBeforeClose != "", props.OnBeforeClose, "// No custom submit handler"),
		props.ID,
		utils.IfElse(props.OnUploadProgress != "", props.OnUploadProgress, "// No upload progress handler"),
		props.ID,
		utils.IfElse(props.OnUploadProgress != "", props.OnUploadProgress, "// No upload progress handler"),
		props.DebounceMs,
		props.ID,
		utils.IfElse(props.OnSearch != "", props.OnSearch, "// No search handler"),
		utils.IfElse(props.OnSelection != "", props.OnSelection, "// No selection handler"),
		props.LoadDataURL,
		strconv.FormatBool(props.LoadOnOpen),
		props.LoadDataURL,
		strconv.FormatBool(props.CloseOnEscape),
		strconv.FormatBool(props.SubmitOnEnter),
		"false", // Debug mode - could be a prop
	)
}
