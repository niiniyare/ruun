package organisms

import (
	"context"
	"fmt"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
)

// Enhanced Modal Examples
// This file provides comprehensive usage examples for the Enhanced Modal organism
// following the same patterns as other organism examples in the codebase.

// ExampleModalConfigs provides various pre-configured modal examples
type ExampleModalConfigs struct{}

// Basic Modal Examples

// SimpleConfirmationModal demonstrates a basic confirmation modal
func (e *ExampleModalConfigs) SimpleConfirmationModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "simple-confirmation",
		Name:           "confirmationModal",
		Title:          "Delete Item",
		Description:    "Are you sure you want to delete this item? This action cannot be undone.",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeSM,
		Variant:        EnhancedModalDestructive,
		Icon:           "trash",
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: false,
		Header:         true,
		Footer:         true,
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:          "delete",
				Text:        "Delete",
				Type:        "destructive",
				Variant:     atoms.ButtonDestructive,
				AutoClose:   true,
				Position:    "right",
				ConfirmText: "Are you absolutely sure?",
				HXPost:      "/api/items/delete",
				HXTarget:    "#item-list",
				HXSwap:      "outerHTML",
			},
		},
	}
}

// InfoAlertModal demonstrates an information alert modal
func (e *ExampleModalConfigs) InfoAlertModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "info-alert",
		Name:           "infoModal",
		Title:          "Update Available",
		Description:    "A new version of the application is available. Would you like to update now?",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeMD,
		Variant:        EnhancedModalInfo,
		Icon:           "download",
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Header:         true,
		Footer:         true,
		Actions: []EnhancedModalAction{
			{
				ID:        "later",
				Text:      "Update Later",
				Type:      "secondary",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "left",
			},
			{
				ID:          "update",
				Text:        "Update Now",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				IconRight:   "download",
				AutoClose:   true,
				Position:    "right",
				AlpineClick: "startUpdate()",
			},
		},
	}
}

// Drawer Modal Examples

// SidebarSettingsDrawer demonstrates a settings drawer modal
func (e *ExampleModalConfigs) SidebarSettingsDrawer() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "settings-drawer",
		Name:             "settingsDrawer",
		Title:            "Settings",
		Subtitle:         "Customize your experience",
		Type:             EnhancedModalDrawer,
		Size:             EnhancedModalSizeMD,
		Position:         EnhancedModalPositionRight,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   true,
		Scrollable:       true,
		Header:           true,
		Footer:           true,
		AutoSave:         true,
		AutoSaveInterval: 5,
		FormData: map[string]any{
			"theme":         "light",
			"notifications": true,
			"language":      "en",
			"timezone":      "UTC",
		},
		Actions: []EnhancedModalAction{
			{
				ID:          "reset",
				Text:        "Reset to Defaults",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				AlpineClick: "resetSettings()",
			},
			{
				ID:        "save",
				Text:      "Save Settings",
				Type:      "primary",
				Variant:   atoms.ButtonPrimary,
				AutoClose: true,
				Position:  "right",
				HXPost:    "/api/settings/save",
				HXTarget:  "body",
				HXSwap:    "none",
			},
		},
	}
}

// NotificationDrawer demonstrates a notification drawer
func (e *ExampleModalConfigs) NotificationDrawer() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "notifications-drawer",
		Name:           "notificationsDrawer",
		Title:          "Notifications",
		Type:           EnhancedModalDrawer,
		Size:           EnhancedModalSizeLG,
		Position:       EnhancedModalPositionRight,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Scrollable:     true,
		Header:         true,
		Footer:         true,
		HXGet:          "/api/notifications",
		LoadOnOpen:     true,
		ReloadOnOpen:   true,
		Actions: []EnhancedModalAction{
			{
				ID:          "mark-all-read",
				Text:        "Mark All as Read",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				HXPost:      "/api/notifications/mark-all-read",
				HXTarget:    "#notifications-list",
				AlpineClick: "markAllAsRead()",
			},
			{
				ID:        "close",
				Text:      "Close",
				Type:      "primary",
				Variant:   atoms.ButtonPrimary,
				AutoClose: true,
				Position:  "right",
			},
		},
	}
}

// Form Modal Examples

// UserProfileEditModal demonstrates a form modal for editing user profiles
func (e *ExampleModalConfigs) UserProfileEditModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "edit-profile",
		Name:             "profileModal",
		Title:            "Edit Profile",
		Description:      "Update your profile information",
		Type:             EnhancedModalDialog,
		Size:             EnhancedModalSizeLG,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   false,
		Scrollable:       true,
		Header:           true,
		Footer:           true,
		AutoFocus:        true,
		TrapFocus:        true,
		RestoreFocus:     true,
		FormValidation:   true,
		AutoSave:         true,
		AutoSaveInterval: 30,
		ConfirmOnClose:   true,
		FormSchema: &schema.FormSchema{
			ID:      "user-profile",
			Version: "1.0.0",
			Title:   "User Profile",
			Fields: []schema.Field{
				{
					Name:        "firstName",
					Label:       "First Name",
					Type:        schema.FieldString,
					Required:    true,
					Placeholder: "Enter your first name",
				},
				{
					Name:        "lastName",
					Label:       "Last Name",
					Type:        schema.FieldString,
					Required:    true,
					Placeholder: "Enter your last name",
				},
				{
					Name:        "email",
					Label:       "Email Address",
					Type:        schema.FieldEmail,
					Required:    true,
					Placeholder: "Enter your email address",
				},
				{
					Name:        "phone",
					Label:       "Phone Number",
					Type:        schema.FieldTel,
					Placeholder: "Enter your phone number",
				},
				{
					Name:        "bio",
					Label:       "Bio",
					Type:        schema.FieldText,
					Placeholder: "Tell us about yourself",
				},
			},
		},
		FormData: map[string]any{
			"firstName": "John",
			"lastName":  "Doe",
			"email":     "john.doe@example.com",
			"phone":     "+1234567890",
			"bio":       "Software developer with a passion for creating great user experiences.",
		},
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:             "save",
				Text:           "Save Changes",
				Type:           "primary",
				Variant:        atoms.ButtonPrimary,
				AutoClose:      true,
				Position:       "right",
				RequiredFields: []string{"firstName", "lastName", "email"},
				HXPost:         "/api/profile/update",
				HXTarget:       "#profile-section",
				HXSwap:         "outerHTML",
			},
		},
	}
}

// CreateProjectModal demonstrates a complex form modal with validation
func (e *ExampleModalConfigs) CreateProjectModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "create-project",
		Name:             "createProjectModal",
		Title:            "Create New Project",
		Description:      "Set up your new project with initial configuration",
		Type:             EnhancedModalDialog,
		Size:             EnhancedModalSizeXL,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   false,
		Scrollable:       true,
		Header:           true,
		Footer:           true,
		AutoFocus:        true,
		TrapFocus:        true,
		FormValidation:   true,
		AutoSave:         true,
		AutoSaveInterval: 15,
		FormSchema: &schema.FormSchema{
			ID:      "create-project",
			Version: "1.0.0",
			Title:   "Project Information",
			Fields: []schema.Field{
				{
					Name:        "name",
					Label:       "Project Name",
					Type:        schema.FieldString,
					Required:    true,
					Placeholder: "Enter project name",
				},
				{
					Name:        "description",
					Label:       "Description",
					Type:        schema.FieldText,
					Placeholder: "Describe your project",
				},
				{
					Name:     "category",
					Label:    "Category",
					Type:     schema.FieldSelect,
					Required: true,
					Options: []schema.FieldOption{
						{Value: "web", Label: "Web Application"},
						{Value: "mobile", Label: "Mobile App"},
						{Value: "desktop", Label: "Desktop Application"},
						{Value: "api", Label: "API/Backend"},
					},
				},
				{
					Name:  "framework",
					Label: "Framework",
					Type:  schema.FieldSelect,
					Options: []schema.FieldOption{
						{Value: "react", Label: "React"},
						{Value: "vue", Label: "Vue.js"},
						{Value: "angular", Label: "Angular"},
						{Value: "go", Label: "Go"},
						{Value: "node", Label: "Node.js"},
					},
				},
				{
					Name:  "features",
					Label: "Features",
					Type:  schema.FieldCheckbox,
					Options: []schema.FieldOption{
						{Value: "auth", Label: "Authentication"},
						{Value: "db", Label: "Database"},
						{Value: "api", Label: "REST API"},
						{Value: "tests", Label: "Unit Tests"},
						{Value: "docker", Label: "Docker"},
					},
				},
			},
		},
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:             "create",
				Text:           "Create Project",
				Type:           "primary",
				Variant:        atoms.ButtonPrimary,
				AutoClose:      true,
				Position:       "right",
				RequiredFields: []string{"name", "category"},
				HXPost:         "/api/projects/create",
				HXTarget:       "#projects-list",
				HXSwap:         "afterbegin",
				Loading:        false,
			},
		},
	}
}

// Multi-Step Modal Examples

// OnboardingWizardModal demonstrates a multi-step onboarding process
func (e *ExampleModalConfigs) OnboardingWizardModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "onboarding-wizard",
		Name:             "onboardingModal",
		Title:            "Welcome to Our Platform",
		Subtitle:         "Let's get you set up in just a few steps",
		Type:             EnhancedModalDialog,
		Size:             EnhancedModalSizeLG,
		Dismissible:      false, // Prevent dismissing during onboarding
		CloseOnEscape:    false,
		CloseOnOverlay:   false,
		Header:           true,
		Footer:           true,
		ShowProgress:     true,
		ShowStepNav:      true,
		StepNavigation:   true,
		LinearProgress:   true,
		AutoSave:         true,
		AutoSaveInterval: 10,
		Steps: []EnhancedModalStep{
			{
				ID:          "welcome",
				Title:       "Welcome",
				Description: "Welcome to our platform! Let's get started.",
				Icon:        "hand",
				Required:    false,
				Active:      true,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{
							Name:        "fullName",
							Label:       "Full Name",
							Type:        schema.FieldString,
							Required:    true,
							Placeholder: "Enter your full name",
						},
						{
							Name:     "role",
							Label:    "Role",
							Type:     schema.FieldSelect,
							Required: true,
							Options: []schema.FieldOption{
								{Value: "developer", Label: "Developer"},
								{Value: "designer", Label: "Designer"},
								{Value: "manager", Label: "Project Manager"},
								{Value: "other", Label: "Other"},
							},
						},
					},
				},
			},
			{
				ID:          "preferences",
				Title:       "Preferences",
				Description: "Configure your workspace preferences",
				Icon:        "settings",
				Required:    true,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{
							Name:     "theme",
							Label:    "Theme",
							Type:     schema.FieldRadio,
							Required: true,
							Options: []schema.FieldOption{
								{Value: "light", Label: "Light"},
								{Value: "dark", Label: "Dark"},
								{Value: "auto", Label: "Auto"},
							},
						},
						{
							Name:     "language",
							Label:    "Language",
							Type:     schema.FieldSelect,
							Required: true,
							Options: []schema.FieldOption{
								{Value: "en", Label: "English"},
								{Value: "es", Label: "Spanish"},
								{Value: "fr", Label: "French"},
								{Value: "de", Label: "German"},
							},
						},
						{
							Name:  "notifications",
							Label: "Notification Preferences",
							Type:  schema.FieldCheckbox,
							Options: []schema.FieldOption{
								{Value: "email", Label: "Email Notifications"},
								{Value: "push", Label: "Push Notifications"},
								{Value: "sms", Label: "SMS Notifications"},
							},
						},
					},
				},
			},
			{
				ID:          "team",
				Title:       "Team Setup",
				Description: "Invite your team members (optional)",
				Icon:        "users",
				Required:    false,
				Skippable:   true,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{
							Name:        "teamName",
							Label:       "Team Name",
							Type:        schema.FieldString,
							Placeholder: "Enter your team name",
						},
						{
							Name:        "inviteEmails",
							Label:       "Invite Members",
							Type:        schema.FieldText,
							Placeholder: "Enter email addresses separated by commas",
						},
					},
				},
			},
			{
				ID:          "complete",
				Title:       "All Set!",
				Description: "Your account is ready to go",
				Icon:        "check-circle",
				Required:    false,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{
							Name:  "newsletter",
							Label: "Subscribe to Newsletter",
							Type:  schema.FieldCheckbox,
							Options: []schema.FieldOption{
								{Value: "subscribe", Label: "Yes, send me product updates and tips"},
							},
						},
					},
				},
			},
		},
		Actions: []EnhancedModalAction{
			{
				ID:          "complete",
				Text:        "Get Started",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				Position:    "right",
				HXPost:      "/api/onboarding/complete",
				HXTarget:    "body",
				HXSwap:      "none",
				AlpineClick: "completeOnboarding()",
			},
		},
	}
}

// ProductCreationWizard demonstrates a complex multi-step product creation process
func (e *ExampleModalConfigs) ProductCreationWizard() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "product-wizard",
		Name:             "productWizard",
		Title:            "Create New Product",
		Type:             EnhancedModalDialog,
		Size:             EnhancedModalSizeXL,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   false,
		Header:           true,
		Footer:           true,
		ShowProgress:     true,
		ShowStepNav:      true,
		StepNavigation:   true,
		LinearProgress:   false, // Allow jumping between steps
		AutoSave:         true,
		AutoSaveInterval: 20,
		ConfirmOnClose:   true,
		Steps: []EnhancedModalStep{
			{
				ID:          "basic-info",
				Title:       "Basic Information",
				Description: "Product details and description",
				Icon:        "package",
				Required:    true,
				Active:      true,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{Name: "name", Label: "Product Name", Type: schema.FieldString, Required: true},
						{Name: "sku", Label: "SKU", Type: schema.FieldString, Required: true},
						{Name: "description", Label: "Description", Type: schema.FieldText, Required: true},
						{Name: "shortDescription", Label: "Short Description", Type: schema.FieldString},
					},
				},
			},
			{
				ID:          "pricing",
				Title:       "Pricing & Inventory",
				Description: "Set price and inventory details",
				Icon:        "dollar-sign",
				Required:    true,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{Name: "price", Label: "Price", Type: schema.FieldNumber, Required: true},
						{Name: "costPrice", Label: "Cost Price", Type: schema.FieldNumber},
						{Name: "stock", Label: "Stock Quantity", Type: schema.FieldNumber, Required: true},
						{Name: "lowStockThreshold", Label: "Low Stock Alert", Type: schema.FieldNumber},
					},
				},
			},
			{
				ID:          "categories",
				Title:       "Categories & Tags",
				Description: "Organize your product",
				Icon:        "tag",
				Required:    false,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{
							Name:  "category",
							Label: "Primary Category",
							Type:  schema.FieldSelect,
							Options: []schema.FieldOption{
								{Value: "electronics", Label: "Electronics"},
								{Value: "clothing", Label: "Clothing"},
								{Value: "books", Label: "Books"},
								{Value: "home", Label: "Home & Garden"},
							},
						},
						{Name: "tags", Label: "Tags", Type: schema.FieldString, Placeholder: "Enter tags separated by commas"},
						{Name: "brand", Label: "Brand", Type: schema.FieldString},
					},
				},
			},
			{
				ID:          "images",
				Title:       "Product Images",
				Description: "Upload product photos",
				Icon:        "image",
				Required:    false,
				// This step would use the file upload functionality
			},
			{
				ID:          "shipping",
				Title:       "Shipping & Dimensions",
				Description: "Shipping information and physical dimensions",
				Icon:        "truck",
				Required:    false,
				Content: &schema.FormSchema{
					Fields: []schema.Field{
						{Name: "weight", Label: "Weight (kg)", Type: schema.FieldNumber},
						{Name: "length", Label: "Length (cm)", Type: schema.FieldNumber},
						{Name: "width", Label: "Width (cm)", Type: schema.FieldNumber},
						{Name: "height", Label: "Height (cm)", Type: schema.FieldNumber},
						{Name: "freeShipping", Label: "Free Shipping", Type: schema.FieldCheckbox},
					},
				},
			},
		},
		Actions: []EnhancedModalAction{
			{
				ID:          "save-draft",
				Text:        "Save as Draft",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				HXPost:      "/api/products/draft",
				AlpineClick: "saveDraft()",
			},
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:             "publish",
				Text:           "Publish Product",
				Type:           "primary",
				Variant:        atoms.ButtonPrimary,
				AutoClose:      true,
				Position:       "right",
				RequiredFields: []string{"name", "sku", "price", "stock"},
				HXPost:         "/api/products/create",
				HXTarget:       "#products-grid",
				HXSwap:         "afterbegin",
			},
		},
	}
}

// Search and Selection Modal Examples

// UserSearchModal demonstrates a search and selection modal
func (e *ExampleModalConfigs) UserSearchModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "user-search",
		Name:           "userSearchModal",
		Title:          "Select Team Members",
		Description:    "Search and select users to add to your team",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeLG,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Header:         true,
		Footer:         true,
		Searchable:     true,
		SearchURL:      "/api/users/search",
		Selectable:     true,
		MultiSelect:    true,
		DebounceMs:     300,
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:          "add-members",
				Text:        "Add Selected Members",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				AutoClose:   true,
				Position:    "right",
				DisabledIf:  "selectedItems.length === 0",
				AlpineClick: "addSelectedMembers()",
				HXPost:      "/api/team/add-members",
			},
		},
	}
}

// ProductLookupModal demonstrates a product lookup with filters
func (e *ExampleModalConfigs) ProductLookupModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "product-lookup",
		Name:           "productLookupModal",
		Title:          "Select Products",
		Description:    "Search and select products to add to your order",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeXL,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Header:         true,
		Footer:         true,
		Scrollable:     true,
		Searchable:     true,
		SearchURL:      "/api/products/search",
		Selectable:     true,
		MultiSelect:    true,
		DebounceMs:     500,
		HXGet:          "/api/products/search-form",
		LoadOnOpen:     true,
		Actions: []EnhancedModalAction{
			{
				ID:          "clear",
				Text:        "Clear Selection",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				AlpineClick: "clearSelection()",
				VisibleIf:   "selectedItems.length > 0",
			},
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:          "add-products",
				Text:        "Add to Order",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				AutoClose:   true,
				Position:    "right",
				DisabledIf:  "selectedItems.length === 0",
				AlpineClick: "addToOrder()",
				HXPost:      "/api/orders/add-products",
			},
		},
	}
}

// File Upload Modal Examples

// DocumentUploadModal demonstrates a document upload modal
func (e *ExampleModalConfigs) DocumentUploadModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "document-upload",
		Name:           "documentUploadModal",
		Title:          "Upload Documents",
		Description:    "Upload important documents for your application",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeLG,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Header:         true,
		Footer:         true,
		FileUpload:     true,
		UploadURL:      "/api/documents/upload",
		UploadMultiple: true,
		UploadMaxFiles: 5,
		UploadMaxSize:  10, // 10MB
		UploadAccept:   ".pdf,.doc,.docx,.txt",
		ShowProgress:   true,
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:          "upload",
				Text:        "Upload Documents",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				AutoClose:   true,
				Position:    "right",
				DisabledIf:  "files.length === 0 || uploadQueue.length > 0",
				AlpineClick: "startUpload()",
			},
		},
	}
}

// ImageGalleryUpload demonstrates an image upload modal with preview
func (e *ExampleModalConfigs) ImageGalleryUpload() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "image-upload",
		Name:           "imageUploadModal",
		Title:          "Upload Images",
		Description:    "Add images to your gallery",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeXL,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: false,
		Header:         true,
		Footer:         true,
		FileUpload:     true,
		UploadURL:      "/api/gallery/upload",
		UploadMultiple: true,
		UploadMaxFiles: 20,
		UploadMaxSize:  5, // 5MB per image
		UploadAccept:   ".jpg,.jpeg,.png,.gif,.webp",
		ShowProgress:   true,
		AutoSave:       false, // Don't auto-save file uploads
		ConfirmOnClose: true,
		Actions: []EnhancedModalAction{
			{
				ID:          "remove-all",
				Text:        "Remove All",
				Type:        "destructive",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				VisibleIf:   "files.length > 0",
				ConfirmText: "Are you sure you want to remove all images?",
				AlpineClick: "removeAllFiles()",
			},
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:          "upload",
				Text:        "Upload Images",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				AutoClose:   true,
				Position:    "right",
				DisabledIf:  "files.length === 0 || uploadQueue.length > 0",
				AlpineClick: "uploadImages()",
				IconLeft:    "upload",
			},
		},
	}
}

// Fullscreen Modal Examples

// DataVisualizationModal demonstrates a fullscreen data modal
func (e *ExampleModalConfigs) DataVisualizationModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:            "data-visualization",
		Name:          "dataModal",
		Title:         "Sales Analytics Dashboard",
		Subtitle:      "Comprehensive view of your sales data",
		Type:          EnhancedModalFullscreen,
		Dismissible:   true,
		CloseOnEscape: true,
		Header:        true,
		Footer:        true,
		Scrollable:    true,
		HXGet:         "/api/analytics/dashboard",
		LoadOnOpen:    true,
		ReloadOnOpen:  true,
		LoadingText:   "Loading analytics data...",
		Actions: []EnhancedModalAction{
			{
				ID:          "export",
				Text:        "Export Data",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				IconLeft:    "download",
				AlpineClick: "exportData()",
				HXPost:      "/api/analytics/export",
			},
			{
				ID:          "refresh",
				Text:        "Refresh",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				IconLeft:    "refresh-cw",
				AlpineClick: "refreshData()",
				HXGet:       "/api/analytics/dashboard",
				HXTarget:    "#dashboard-content",
			},
			{
				ID:        "close",
				Text:      "Close",
				Type:      "primary",
				Variant:   atoms.ButtonPrimary,
				AutoClose: true,
				Position:  "right",
			},
		},
	}
}

// ReportViewerModal demonstrates a fullscreen document viewer
func (e *ExampleModalConfigs) ReportViewerModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:            "report-viewer",
		Name:          "reportViewer",
		Title:         "Monthly Report",
		Type:          EnhancedModalFullscreen,
		Dismissible:   true,
		CloseOnEscape: true,
		Header:        true,
		Footer:        true,
		Scrollable:    true,
		HXGet:         "/api/reports/monthly/view",
		LoadOnOpen:    true,
		Actions: []EnhancedModalAction{
			{
				ID:          "download-pdf",
				Text:        "Download PDF",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				IconLeft:    "download",
				AlpineClick: "downloadPDF()",
			},
			{
				ID:          "print",
				Text:        "Print",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Position:    "left",
				IconLeft:    "printer",
				AlpineClick: "printReport()",
			},
			{
				ID:        "close",
				Text:      "Close",
				Type:      "primary",
				Variant:   atoms.ButtonPrimary,
				AutoClose: true,
				Position:  "right",
			},
		},
	}
}

// Bottom Sheet Modal Examples (Mobile-Optimized)

// MobileMenuBottomSheet demonstrates a mobile menu bottom sheet
func (e *ExampleModalConfigs) MobileMenuBottomSheet() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "mobile-menu",
		Name:             "mobileMenu",
		Title:            "Menu",
		Type:             EnhancedModalBottomSheet,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   true,
		Header:           true,
		Footer:           false,
		MobileFullscreen: true,
		HXGet:            "/api/menu/mobile",
		LoadOnOpen:       true,
		Actions: []EnhancedModalAction{
			{
				ID:        "close",
				Text:      "Close",
				Type:      "secondary",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
				FullWidth: true,
			},
		},
	}
}

// MobileShareSheet demonstrates a mobile share bottom sheet
func (e *ExampleModalConfigs) MobileShareSheet() EnhancedModalProps {
	return EnhancedModalProps{
		ID:               "share-sheet",
		Name:             "shareSheet",
		Title:            "Share",
		Description:      "Share this content",
		Type:             EnhancedModalBottomSheet,
		Size:             EnhancedModalSizeAuto,
		Dismissible:      true,
		CloseOnEscape:    true,
		CloseOnOverlay:   true,
		Header:           true,
		Footer:           false,
		MobileFullscreen: false,
		Actions: []EnhancedModalAction{
			{
				ID:          "copy-link",
				Text:        "Copy Link",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				IconLeft:    "copy",
				FullWidth:   true,
				Position:    "center",
				AlpineClick: "copyToClipboard()",
			},
			{
				ID:          "share-email",
				Text:        "Share via Email",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				IconLeft:    "mail",
				FullWidth:   true,
				Position:    "center",
				AlpineClick: "shareViaEmail()",
			},
		},
	}
}

// Popover Modal Examples

// UserProfilePopover demonstrates a user profile popover
func (e *ExampleModalConfigs) UserProfilePopover() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "user-profile-popover",
		Name:           "profilePopover",
		Title:          "John Doe",
		Type:           EnhancedModalPopover,
		Size:           EnhancedModalSizeAuto,
		Position:       EnhancedModalPositionTop,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Backdrop:       false,
		Header:         false,
		Footer:         true,
		HXGet:          "/api/users/profile-popover",
		HXTarget:       "this",
		LoadOnOpen:     true,
		Actions: []EnhancedModalAction{
			{
				ID:          "view-profile",
				Text:        "View Profile",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				Size:        atoms.ButtonSizeSM,
				AutoClose:   true,
				Position:    "center",
				AlpineClick: "viewFullProfile()",
			},
			{
				ID:          "send-message",
				Text:        "Send Message",
				Type:        "secondary",
				Variant:     atoms.ButtonOutline,
				Size:        atoms.ButtonSizeSM,
				AutoClose:   true,
				Position:    "center",
				AlpineClick: "openMessageDialog()",
			},
		},
	}
}

// QuickActionsPopover demonstrates a quick actions popover
func (e *ExampleModalConfigs) QuickActionsPopover() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "quick-actions",
		Name:           "quickActions",
		Title:          "Quick Actions",
		Type:           EnhancedModalPopover,
		Size:           EnhancedModalSizeAuto,
		Position:       EnhancedModalPositionBottom,
		Dismissible:    true,
		CloseOnEscape:  true,
		CloseOnOverlay: true,
		Backdrop:       false,
		Header:         true,
		Footer:         false,
		Actions: []EnhancedModalAction{
			{
				ID:          "edit",
				Text:        "Edit",
				Type:        "primary",
				Variant:     atoms.ButtonGhost,
				IconLeft:    "edit",
				AutoClose:   true,
				Position:    "center",
				FullWidth:   true,
				AlpineClick: "editItem()",
			},
			{
				ID:          "duplicate",
				Text:        "Duplicate",
				Type:        "secondary",
				Variant:     atoms.ButtonGhost,
				IconLeft:    "copy",
				AutoClose:   true,
				Position:    "center",
				FullWidth:   true,
				AlpineClick: "duplicateItem()",
			},
			{
				ID:          "delete",
				Text:        "Delete",
				Type:        "destructive",
				Variant:     atoms.ButtonGhost,
				IconLeft:    "trash",
				AutoClose:   true,
				Position:    "center",
				FullWidth:   true,
				ConfirmText: "Are you sure you want to delete this item?",
				AlpineClick: "deleteItem()",
			},
		},
	}
}

// Advanced Modal Examples

// WorkflowApprovalModal demonstrates a workflow approval modal with business logic
func (e *ExampleModalConfigs) WorkflowApprovalModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:                  "workflow-approval",
		Name:                "approvalModal",
		Title:               "Approval Required",
		Description:         "Review and approve this expense request",
		Type:                EnhancedModalDialog,
		Size:                EnhancedModalSizeLG,
		Variant:             EnhancedModalInfo,
		Icon:                "check-circle",
		Dismissible:         true,
		CloseOnEscape:       true,
		CloseOnOverlay:      false,
		Header:              true,
		Footer:              true,
		Scrollable:          true,
		WorkflowID:          "expense-approval",
		ProcessID:           "exp-001",
		TaskID:              "approval-123",
		EntityID:            "expense-456",
		EntityType:          "expense",
		RequiredPermissions: []string{"approve_expenses"},
		SecureMode:          true,
		HXGet:               "/api/expenses/456/approval-details",
		LoadOnOpen:          true,
		FormSchema: &schema.FormSchema{
			Fields: []schema.Field{
				{
					Name:        "comments",
					Label:       "Approval Comments",
					Type:        schema.FieldText,
					Placeholder: "Add any comments about this approval",
				},
				{
					Name:  "notifySubmitter",
					Label: "Notify Submitter",
					Type:  schema.FieldCheckbox,
					Options: []schema.FieldOption{
						{Value: "notify", Label: "Send notification to submitter"},
					},
				},
			},
		},
		Actions: []EnhancedModalAction{
			{
				ID:             "reject",
				Text:           "Reject",
				Type:           "destructive",
				Variant:        atoms.ButtonDestructive,
				Position:       "left",
				ConfirmText:    "Are you sure you want to reject this request?",
				RequiredFields: []string{"comments"},
				HXPost:         "/api/expenses/456/reject",
				HXTarget:       "#expense-details",
				AutoClose:      true,
			},
			{
				ID:        "cancel",
				Text:      "Cancel",
				Type:      "cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
				Position:  "right",
			},
			{
				ID:        "approve",
				Text:      "Approve",
				Type:      "primary",
				Variant:   atoms.ButtonPrimary,
				AutoClose: true,
				Position:  "right",
				HXPost:    "/api/expenses/456/approve",
				HXTarget:  "#expense-details",
				IconLeft:  "check",
			},
		},
	}
}

// DataImportWizard demonstrates a complex data import wizard
func (e *ExampleModalConfigs) DataImportWizard() EnhancedModalProps {
	return EnhancedModalProps{
		ID:             "data-import-wizard",
		Name:           "importWizard",
		Title:          "Import Data",
		Description:    "Import data from CSV or Excel files",
		Type:           EnhancedModalDialog,
		Size:           EnhancedModalSizeXL,
		Dismissible:    true,
		CloseOnEscape:  false, // Prevent accidental close during import
		CloseOnOverlay: false,
		Header:         true,
		Footer:         true,
		ShowProgress:   true,
		ShowStepNav:    true,
		StepNavigation: true,
		LinearProgress: true,
		AutoSave:       false, // Don't auto-save import process
		ConfirmOnClose: true,
		Steps: []EnhancedModalStep{
			{
				ID:          "upload",
				Title:       "Upload File",
				Description: "Select the file you want to import",
				Icon:        "upload",
				Required:    true,
				Active:      true,
				// This would integrate with file upload functionality
			},
			{
				ID:          "mapping",
				Title:       "Field Mapping",
				Description: "Map your file columns to database fields",
				Icon:        "shuffle",
				Required:    true,
				// This would show a dynamic mapping interface
			},
			{
				ID:          "validation",
				Title:       "Validation",
				Description: "Review and fix any data issues",
				Icon:        "check-square",
				Required:    true,
				// This would show validation results
			},
			{
				ID:          "preview",
				Title:       "Preview",
				Description: "Preview the data before import",
				Icon:        "eye",
				Required:    false,
				// This would show a preview of the import
			},
			{
				ID:          "import",
				Title:       "Import",
				Description: "Import the data into your system",
				Icon:        "download",
				Required:    true,
				// This would show import progress
			},
		},
		Actions: []EnhancedModalAction{
			{
				ID:          "cancel",
				Text:        "Cancel",
				Type:        "cancel",
				Variant:     atoms.ButtonOutline,
				AutoClose:   true,
				Position:    "right",
				ConfirmText: "Are you sure you want to cancel the import? All progress will be lost.",
			},
			{
				ID:          "start-import",
				Text:        "Start Import",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				Position:    "right",
				VisibleIf:   "currentStep === totalSteps",
				AlpineClick: "startImport()",
				HXPost:      "/api/data/import/start",
			},
		},
	}
}

// Schema-Driven Modal Examples

// GetSchemaBasedModal demonstrates creating modals from schema configuration
func (e *ExampleModalConfigs) GetSchemaBasedModal(modalType string) (*ModalSchema, error) {
	switch modalType {
	case "confirmation":
		return CreateConfirmationModalSchema(
			"confirm-delete",
			"Delete Item",
			"This action cannot be undone. Are you sure you want to delete this item?",
			true,
		), nil

	case "user-form":
		userSchema := &schema.FormSchema{
			ID:      "user-form",
			Version: "1.0.0",
			Title:   "User Information",
			Fields: []schema.Field{
				{Name: "firstName", Label: "First Name", Type: schema.FieldString, Required: true},
				{Name: "lastName", Label: "Last Name", Type: schema.FieldString, Required: true},
				{Name: "email", Label: "Email", Type: schema.FieldEmail, Required: true},
				{Name: "role", Label: "Role", Type: schema.FieldSelect, Required: true, Options: []schema.FieldOption{
					{Value: "admin", Label: "Administrator"},
					{Value: "user", Label: "User"},
					{Value: "viewer", Label: "Viewer"},
				}},
			},
		}
		return CreateFormModalSchema("user-form", "Edit User", userSchema), nil

	case "search-users":
		return CreateSearchModalSchema(
			"search-users",
			"Search Users",
			"/api/users/search",
			true,
		), nil

	case "upload-documents":
		return CreateUploadModalSchema(
			"upload-documents",
			"Upload Documents",
			"/api/documents/upload",
			5,
			[]string{".pdf", ".doc", ".docx", ".txt"},
		), nil

	default:
		return nil, fmt.Errorf("unknown modal type: %s", modalType)
	}
}

// Modal Usage Examples

// Example Alpine.js data for modal interaction
func (e *ExampleModalConfigs) GetModalAlpineData() string {
	return `{
		// Modal state
		confirmationModalOpen: false,
		profileModalOpen: false,
		searchModalOpen: false,
		
		// Form data
		userForm: {
			firstName: '',
			lastName: '',
			email: '',
			role: ''
		},
		
		// Search state
		searchQuery: '',
		selectedUsers: [],
		
		// Methods
		openConfirmation() {
			this.confirmationModalOpen = true;
		},
		
		closeConfirmation() {
			this.confirmationModalOpen = false;
		},
		
		openProfileEdit() {
			this.profileModalOpen = true;
		},
		
		closeProfileEdit() {
			this.profileModalOpen = false;
		},
		
		openUserSearch() {
			this.searchModalOpen = true;
		},
		
		closeUserSearch() {
			this.searchModalOpen = false;
		},
		
		saveProfile() {
			// Validate and save profile
			fetch('/api/profile/update', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(this.userForm)
			})
			.then(response => response.json())
			.then(data => {
				if (data.success) {
					this.closeProfileEdit();
					// Show success message
				}
			});
		},
		
		deleteItem(id) {
			// Confirm and delete
			if (confirm('Are you sure you want to delete this item?')) {
				fetch('/api/items/' + id, {
					method: 'DELETE'
				})
				.then(response => {
					if (response.ok) {
						// Remove item from UI
						location.reload();
					}
				});
			}
		}
	}`
}

// Example CSS classes for modal styling
func (e *ExampleModalConfigs) GetModalCSS() string {
	return `
/* Enhanced Modal Custom Styles */
.enhanced-modal-overlay {
	backdrop-filter: blur(4px);
}

.enhanced-modal-content {
	box-shadow: 
		0 25px 50px -12px rgba(0, 0, 0, 0.25),
		0 0 0 1px rgba(255, 255, 255, 0.05);
}

.enhanced-modal-header {
	background: linear-gradient(to bottom, 
		rgba(255, 255, 255, 0.1), 
		rgba(255, 255, 255, 0.05));
}

/* Modal animations */
@keyframes modalSlideIn {
	from {
		opacity: 0;
		transform: translateY(16px) scale(0.95);
	}
	to {
		opacity: 1;
		transform: translateY(0) scale(1);
	}
}

@keyframes modalSlideOut {
	from {
		opacity: 1;
		transform: translateY(0) scale(1);
	}
	to {
		opacity: 0;
		transform: translateY(16px) scale(0.95);
	}
}

/* Drawer animations */
@keyframes drawerSlideInRight {
	from {
		transform: translateX(100%);
	}
	to {
		transform: translateX(0);
	}
}

@keyframes drawerSlideOutRight {
	from {
		transform: translateX(0);
	}
	to {
		transform: translateX(100%);
	}
}

/* Bottom sheet animations */
@keyframes bottomSheetSlideIn {
	from {
		transform: translateY(100%);
	}
	to {
		transform: translateY(0);
	}
}

@keyframes bottomSheetSlideOut {
	from {
		transform: translateY(0);
	}
	to {
		transform: translateY(100%);
	}
}

/* Step navigation styles */
.step-nav-item {
	transition: all 0.2s ease;
}

.step-nav-item--active {
	transform: scale(1.05);
}

/* File upload styles */
.file-upload-area {
	transition: all 0.2s ease;
}

.file-upload-area--dragging {
	transform: scale(1.02);
	border-color: var(--primary);
	background-color: var(--primary-50);
}

/* Search result styles */
.search-result-item {
	transition: all 0.15s ease;
}

.search-result-item:hover {
	transform: translateY(-1px);
	box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

/* Mobile optimizations */
@media (max-width: 640px) {
	.enhanced-modal-content {
		margin: 0;
		border-radius: 0;
		min-height: 100vh;
	}
	
	.enhanced-modal-overlay {
		padding: 0;
	}
}
`
}

// Integration examples with HTMX

// ExampleHTMXIntegration demonstrates HTMX integration patterns
func (e *ExampleModalConfigs) ExampleHTMXIntegration() map[string]string {
	return map[string]string{
		"trigger-modal": `
<!-- Button to trigger modal -->
<button 
	hx-get="/api/modals/user-edit/123"
	hx-target="body"
	hx-swap="beforeend"
	class="btn btn-primary"
>
	Edit User
</button>
`,

		"dynamic-modal": `
<!-- Server returns complete modal HTML -->
<div id="user-edit-modal">
	<!-- Modal content loaded via HTMX -->
</div>
`,

		"form-submission": `
<!-- Form inside modal with HTMX -->
<form 
	hx-post="/api/users/123/update"
	hx-target="#user-list"
	hx-swap="outerHTML"
	hx-on::after-request="closeModal()"
>
	<!-- Form fields -->
	<button type="submit">Save Changes</button>
</form>
`,

		"search-integration": `
<!-- Search input with HTMX -->
<input 
	type="search"
	hx-get="/api/search"
	hx-target="#search-results"
	hx-trigger="keyup changed delay:300ms"
	hx-params="query"
	placeholder="Search..."
/>
<div id="search-results">
	<!-- Results loaded here -->
</div>
`,

		"file-upload": `
<!-- File upload with progress -->
<input 
	type="file"
	multiple
	hx-post="/api/upload"
	hx-target="#upload-results"
	hx-swap="afterbegin"
	hx-indicator="#upload-spinner"
/>
<div id="upload-spinner" class="htmx-indicator">
	Uploading...
</div>
<div id="upload-results">
	<!-- Upload results appear here -->
</div>
`,
	}
}

// Performance optimization examples
func (e *ExampleModalConfigs) PerformanceOptimizedModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:                "performance-modal",
		Name:              "perfModal",
		Title:             "Large Dataset",
		Type:              EnhancedModalDialog,
		Size:              EnhancedModalSizeLG,
		LazyLoad:          true, // Lazy load content
		DebounceMs:        500,  // Debounce search
		AnimationDuration: 200,  // Faster animations
		HXGet:             "/api/data/large-dataset",
		LoadOnOpen:        false, // Manual loading
		ReloadOnOpen:      false, // Don't reload if already loaded
		Actions: []EnhancedModalAction{
			{
				ID:          "load-data",
				Text:        "Load Data",
				Type:        "primary",
				Variant:     atoms.ButtonPrimary,
				Position:    "right",
				AlpineClick: "loadData()",
				HXGet:       "/api/data/large-dataset",
				HXTarget:    "#data-container",
				HXIndicator: "#loading-spinner",
			},
		},
	}
}

// Accessibility-focused modal example
func (e *ExampleModalConfigs) AccessibleModal() EnhancedModalProps {
	return EnhancedModalProps{
		ID:              "accessible-modal",
		Name:            "accessibleModal",
		Title:           "Accessible Form",
		Description:     "Form with full accessibility support",
		Type:            EnhancedModalDialog,
		Size:            EnhancedModalSizeMD,
		AutoFocus:       true,
		TrapFocus:       true,
		RestoreFocus:    true,
		AriaLabel:       "Edit user profile form",
		AriaDescription: "accessible-modal-description",
		Role:            "dialog",
		FormValidation:  true,
		Actions: []EnhancedModalAction{
			{
				ID:        "cancel",
				Text:      "Cancel",
				Variant:   atoms.ButtonOutline,
				AutoClose: true,
			},
			{
				ID:             "save",
				Text:           "Save Changes",
				Variant:        atoms.ButtonPrimary,
				Type:           "submit",
				RequiredFields: []string{"name", "email"},
			},
		},
	}
}

// GetUsageExamples returns various usage patterns and examples
func (e *ExampleModalConfigs) GetUsageExamples() map[string]interface{} {
	return map[string]interface{}{
		"basic_confirmation":    e.SimpleConfirmationModal(),
		"info_alert":            e.InfoAlertModal(),
		"settings_drawer":       e.SidebarSettingsDrawer(),
		"notifications_drawer":  e.NotificationDrawer(),
		"profile_edit_form":     e.UserProfileEditModal(),
		"create_project_form":   e.CreateProjectModal(),
		"onboarding_wizard":     e.OnboardingWizardModal(),
		"product_wizard":        e.ProductCreationWizard(),
		"user_search":           e.UserSearchModal(),
		"product_lookup":        e.ProductLookupModal(),
		"document_upload":       e.DocumentUploadModal(),
		"image_gallery":         e.ImageGalleryUpload(),
		"data_visualization":    e.DataVisualizationModal(),
		"report_viewer":         e.ReportViewerModal(),
		"mobile_menu":           e.MobileMenuBottomSheet(),
		"share_sheet":           e.MobileShareSheet(),
		"profile_popover":       e.UserProfilePopover(),
		"quick_actions":         e.QuickActionsPopover(),
		"workflow_approval":     e.WorkflowApprovalModal(),
		"data_import":           e.DataImportWizard(),
		"performance_optimized": e.PerformanceOptimizedModal(),
		"accessibility_focused": e.AccessibleModal(),
		"alpine_data":           e.GetModalAlpineData(),
		"css_styles":            e.GetModalCSS(),
		"htmx_integration":      e.ExampleHTMXIntegration(),
	}
}

// GetModalByUseCase returns a modal configuration for a specific use case
func GetModalByUseCase(useCase string, params map[string]interface{}) (*EnhancedModalProps, error) {
	examples := &ExampleModalConfigs{}

	switch useCase {
	case "delete_confirmation":
		modal := examples.SimpleConfirmationModal()
		if title, ok := params["title"].(string); ok {
			modal.Title = title
		}
		if description, ok := params["description"].(string); ok {
			modal.Description = description
		}
		return &modal, nil

	case "edit_form":
		modal := examples.UserProfileEditModal()
		if title, ok := params["title"].(string); ok {
			modal.Title = title
		}
		if formSchema, ok := params["form_schema"].(*schema.FormSchema); ok {
			modal.FormSchema = formSchema
		}
		return &modal, nil

	case "search_select":
		modal := examples.UserSearchModal()
		if title, ok := params["title"].(string); ok {
			modal.Title = title
		}
		if searchURL, ok := params["search_url"].(string); ok {
			modal.SearchURL = searchURL
		}
		if multiSelect, ok := params["multi_select"].(bool); ok {
			modal.MultiSelect = multiSelect
		}
		return &modal, nil

	case "file_upload":
		modal := examples.DocumentUploadModal()
		if title, ok := params["title"].(string); ok {
			modal.Title = title
		}
		if uploadURL, ok := params["upload_url"].(string); ok {
			modal.UploadURL = uploadURL
		}
		if maxFiles, ok := params["max_files"].(int); ok {
			modal.UploadMaxFiles = maxFiles
		}
		if acceptedTypes, ok := params["accepted_types"].(string); ok {
			modal.UploadAccept = acceptedTypes
		}
		return &modal, nil

	default:
		return nil, fmt.Errorf("unknown use case: %s", useCase)
	}
}

// GetModalTemplateByType returns a basic modal template for a given type
func GetModalTemplateByType(modalType EnhancedModalType) EnhancedModalProps {
	examples := &ExampleModalConfigs{}

	switch modalType {
	case EnhancedModalDialog:
		return examples.SimpleConfirmationModal()
	case EnhancedModalDrawer:
		return examples.SidebarSettingsDrawer()
	case EnhancedModalBottomSheet:
		return examples.MobileMenuBottomSheet()
	case EnhancedModalFullscreen:
		return examples.DataVisualizationModal()
	case EnhancedModalPopover:
		return examples.UserProfilePopover()
	default:
		return examples.SimpleConfirmationModal()
	}
}

// CreateCustomModalExample creates a custom modal example based on requirements
func CreateCustomModalExample(req CustomModalRequest) (*EnhancedModalProps, error) {
	if req.ID == "" {
		return nil, fmt.Errorf("modal ID is required")
	}

	if req.Title == "" {
		return nil, fmt.Errorf("modal title is required")
	}

	props := &EnhancedModalProps{
		ID:             req.ID,
		Name:           req.Name,
		Title:          req.Title,
		Description:    req.Description,
		Type:           req.Type,
		Size:           req.Size,
		Position:       req.Position,
		Variant:        req.Variant,
		Icon:           req.Icon,
		Dismissible:    req.Dismissible,
		CloseOnEscape:  req.CloseOnEscape,
		CloseOnOverlay: req.CloseOnOverlay,
		Header:         req.ShowHeader,
		Footer:         req.ShowFooter,
		Scrollable:     req.Scrollable,
		AutoFocus:      req.AutoFocus,
		TrapFocus:      req.TrapFocus,
		RestoreFocus:   req.RestoreFocus,
		FormSchema:     req.FormSchema,
		FormData:       req.InitialData,
		Actions:        req.Actions,
		HXGet:          req.LoadURL,
		LoadOnOpen:     req.LoadOnOpen,
		Context:        context.Background(),
	}

	// Set defaults based on type
	switch req.Type {
	case EnhancedModalDrawer:
		if props.Position == "" {
			props.Position = EnhancedModalPositionRight
		}
		if props.Size == "" {
			props.Size = EnhancedModalSizeMD
		}

	case EnhancedModalBottomSheet:
		props.MobileFullscreen = true

	case EnhancedModalFullscreen:
		props.Size = EnhancedModalSizeFull

	case EnhancedModalPopover:
		props.Backdrop = false
		if props.Size == "" {
			props.Size = EnhancedModalSizeAuto
		}

	default: // Dialog
		if props.Size == "" {
			props.Size = EnhancedModalSizeMD
		}
		props.Centered = true
	}

	return props, nil
}

// CustomModalRequest defines the requirements for a custom modal
type CustomModalRequest struct {
	ID             string                `json:"id"`
	Name           string                `json:"name,omitempty"`
	Title          string                `json:"title"`
	Description    string                `json:"description,omitempty"`
	Type           EnhancedModalType     `json:"type"`
	Size           EnhancedModalSize     `json:"size,omitempty"`
	Position       EnhancedModalPosition `json:"position,omitempty"`
	Variant        EnhancedModalVariant  `json:"variant,omitempty"`
	Icon           string                `json:"icon,omitempty"`
	Dismissible    bool                  `json:"dismissible"`
	CloseOnEscape  bool                  `json:"closeOnEscape"`
	CloseOnOverlay bool                  `json:"closeOnOverlay"`
	ShowHeader     bool                  `json:"showHeader"`
	ShowFooter     bool                  `json:"showFooter"`
	Scrollable     bool                  `json:"scrollable"`
	AutoFocus      bool                  `json:"autoFocus"`
	TrapFocus      bool                  `json:"trapFocus"`
	RestoreFocus   bool                  `json:"restoreFocus"`
	FormSchema     *schema.FormSchema    `json:"formSchema,omitempty"`
	InitialData    map[string]any        `json:"initialData,omitempty"`
	Actions        []EnhancedModalAction `json:"actions,omitempty"`
	LoadURL        string                `json:"loadUrl,omitempty"`
	LoadOnOpen     bool                  `json:"loadOnOpen"`
}
