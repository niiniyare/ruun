package organisms

import (
	"time"

	"github.com/niiniyare/ruun/views/components/atoms"
)

// ExampleSidebarNavigation demonstrates a complete sidebar navigation with auth
func ExampleSidebarNavigation() NavigationProps {
	return NavigationProps{
		ID:          "main-sidebar",
		Type:        NavigationSidebar,
		Size:        NavigationSizeMD,
		Title:       "Dashboard",
		Description: "Main application navigation",
		Logo:        "/assets/logo.svg",
		LogoURL:     "/",
		Collapsible: true,
		Collapsed:   false,
		Variant:     "default",

		// Search configuration
		Search: NavigationSearch{
			Enabled:     true,
			Placeholder: "Search menu items...",
			MinLength:   2,
			Debounce:    300,
		},

		// Mobile menu
		MobileMenu: NavigationMobileMenu{
			Enabled:     true,
			Position:    "left",
			Overlay:     true,
			CloseButton: true,
			Swipeable:   true,
		},

		// Navigation items with nested structure
		Items: []NavigationItem{
			{
				ID:     "dashboard",
				Text:   "Dashboard",
				Icon:   "home",
				URL:    "/dashboard",
				Active: true,
			},
			{
				ID:      "divider-main",
				Divider: true,
			},
			{
				ID:           "analytics",
				Text:         "Analytics",
				Description:  "View reports and insights",
				Icon:         "bar-chart-3",
				URL:          "/analytics",
				Badge:        "New",
				BadgeVariant: atoms.BadgePrimary,
			},
			{
				ID:          "users",
				Text:        "User Management",
				Icon:        "users",
				Collapsible: true,
				Items: []NavigationItem{
					{
						ID:           "users-list",
						Text:         "All Users",
						URL:          "/users",
						Icon:         "list",
						Badge:        "1,234",
						BadgeVariant: atoms.BadgeSecondary,
					},
					{
						ID:   "users-roles",
						Text: "Roles & Permissions",
						URL:  "/users/roles",
						Icon: "shield",
					},
					{
						ID:   "users-activity",
						Text: "Activity Log",
						URL:  "/users/activity",
						Icon: "activity",
					},
				},
			},
			{
				ID:          "content",
				Text:        "Content Management",
				Icon:        "file-text",
				Collapsible: true,
				Collapsed:   true,
				Items: []NavigationItem{
					{
						ID:           "content-posts",
						Text:         "Posts",
						URL:          "/content/posts",
						Icon:         "edit-3",
						Badge:        "42",
						BadgeVariant: atoms.BadgeDestructive,
					},
					{
						ID:   "content-pages",
						Text: "Pages",
						URL:  "/content/pages",
						Icon: "file",
					},
					{
						ID:   "content-media",
						Text: "Media Library",
						URL:  "/content/media",
						Icon: "image",
					},
				},
			},
			{
				ID:      "divider-admin",
				Divider: true,
			},
			{
				ID:          "settings",
				Text:        "Settings",
				Description: "Application configuration",
				Icon:        "settings",
				URL:         "/settings",
			},
		},

		// Auth context
		AuthContext: &AuthContext{
			User: &User{
				ID:          "user-123",
				Username:    "admin",
				Email:       "admin@example.com",
				DisplayName: "John Doe",
				Avatar:      "/assets/avatars/admin.jpg",
				Roles:       []string{"admin", "moderator"},
				Permissions: []string{"users:read", "users:write", "content:read", "content:write"},
			},
			Authenticated: true,
		},

		// User interface
		UserName:   "John Doe",
		UserAvatar: "/assets/avatars/admin.jpg",
		UserOnline: true,
		UserRole:   "Administrator",

		// User menu
		UserMenu: []NavigationItem{
			{
				ID:   "profile",
				Text: "Profile",
				URL:  "/profile",
				Icon: "user",
			},
			{
				ID:   "preferences",
				Text: "Preferences",
				URL:  "/preferences",
				Icon: "settings",
			},
			{
				ID:      "divider-user",
				Divider: true,
			},
			{
				ID:   "help",
				Text: "Help & Support",
				URL:  "/help",
				Icon: "help-circle",
			},
			{
				ID:   "logout",
				Text: "Sign Out",
				URL:  "/auth/logout",
				Icon: "log-out",
			},
		},

		// Performance
		LazyLoad:  true,
		CacheData: true,

		// Theme
		DarkMode: false,

		// Accessibility
		AriaLabel: "Main navigation",
		AriaLabels: map[string]string{
			"search":     "Search navigation",
			"userMenu":   "User menu",
			"mobileMenu": "Mobile menu",
		},
	}
}

// ExampleTopbarNavigation demonstrates a topbar navigation
func ExampleTopbarNavigation() NavigationProps {
	return NavigationProps{
		ID:      "main-topbar",
		Type:    NavigationTopbar,
		Size:    NavigationSizeMD,
		Title:   "My Application",
		Logo:    "/assets/logo.svg",
		LogoURL: "/",

		// Search configuration
		Search: NavigationSearch{
			Enabled:     true,
			Placeholder: "Search anything...",
			MinLength:   1,
			Debounce:    500,
		},

		// Mobile menu
		MobileMenu: NavigationMobileMenu{
			Enabled:     true,
			Position:    "top",
			Overlay:     true,
			CloseButton: true,
		},

		// Topbar navigation items
		Items: []NavigationItem{
			{
				ID:   "dashboard",
				Text: "Dashboard",
				URL:  "/dashboard",
				Icon: "home",
			},
			{
				ID:           "projects",
				Text:         "Projects",
				URL:          "/projects",
				Icon:         "folder",
				Badge:        "3",
				BadgeVariant: atoms.BadgePrimary,
			},
			{
				ID:   "teams",
				Text: "Teams",
				URL:  "/teams",
				Icon: "users",
			},
			{
				ID:          "resources",
				Text:        "Resources",
				Icon:        "book",
				Collapsible: true,
				Items: []NavigationItem{
					{
						ID:   "docs",
						Text: "Documentation",
						URL:  "/docs",
						Icon: "book-open",
					},
					{
						ID:   "api",
						Text: "API Reference",
						URL:  "/api-docs",
						Icon: "code",
					},
					{
						ID:   "tutorials",
						Text: "Tutorials",
						URL:  "/tutorials",
						Icon: "play-circle",
					},
				},
			},
		},

		// Auth context
		AuthContext: &AuthContext{
			User: &User{
				ID:          "user-456",
				Username:    "developer",
				Email:       "dev@example.com",
				DisplayName: "Jane Smith",
				Avatar:      "/assets/avatars/dev.jpg",
				Roles:       []string{"developer"},
				Permissions: []string{"projects:read", "projects:write"},
			},
			Authenticated: true,
		},

		// User interface
		UserName:   "Jane Smith",
		UserAvatar: "/assets/avatars/dev.jpg",
		UserOnline: true,
		UserRole:   "Developer",

		// User menu
		UserMenu: []NavigationItem{
			{
				ID:   "account",
				Text: "Account Settings",
				URL:  "/account",
				Icon: "settings",
			},
			{
				ID:   "billing",
				Text: "Billing",
				URL:  "/billing",
				Icon: "credit-card",
			},
			{
				ID:      "divider-user",
				Divider: true,
			},
			{
				ID:   "logout",
				Text: "Sign Out",
				URL:  "/auth/logout",
				Icon: "log-out",
			},
		},

		// Notifications
		Notifications: []Notification{
			{
				ID:        "notif-1",
				Type:      "info",
				Title:     "Welcome to the platform",
				Message:   "Your account has been successfully verified.",
				Read:      false,
				CreatedAt: time.Now().Add(-5 * time.Minute),
			},
			{
				ID:        "notif-2",
				Type:      "warning",
				Title:     "Storage limit warning",
				Message:   "You're approaching your storage limit.",
				Read:      true,
				CreatedAt: time.Now().Add(-2 * time.Hour),
			},
		},

		// Accessibility
		AriaLabel: "Main navigation",
	}
}

// ExampleBreadcrumbNavigation demonstrates breadcrumb navigation
func ExampleBreadcrumbNavigation() NavigationProps {
	return NavigationProps{
		ID:   "breadcrumb-nav",
		Type: NavigationBreadcrumb,
		Size: NavigationSizeMD,

		// Breadcrumb items
		Breadcrumbs: []BreadcrumbItem{
			{
				ID:        "home",
				Text:      "Home",
				URL:       "/",
				Icon:      "home",
				Clickable: true,
			},
			{
				ID:        "projects",
				Text:      "Projects",
				URL:       "/projects",
				Clickable: true,
			},
			{
				ID:        "project-detail",
				Text:      "Website Redesign",
				URL:       "/projects/123",
				Clickable: true,
			},
			{
				ID:     "current",
				Text:   "Settings",
				Active: true,
			},
		},

		Separator: "/",

		// Accessibility
		AriaLabel: "Breadcrumb navigation",
	}
}

// ExampleTabsNavigation demonstrates tab navigation
func ExampleTabsNavigation() NavigationProps {
	return NavigationProps{
		ID:         "content-tabs",
		Type:       NavigationTabs,
		Size:       NavigationSizeMD,
		TabVariant: "underline",

		// Tab items
		Items: []NavigationItem{
			{
				ID:     "overview",
				Text:   "Overview",
				URL:    "/project/overview",
				Icon:   "eye",
				Active: true,
			},
			{
				ID:           "files",
				Text:         "Files",
				URL:          "/project/files",
				Icon:         "folder",
				Badge:        "24",
				BadgeVariant: atoms.BadgeSecondary,
			},
			{
				ID:   "settings",
				Text: "Settings",
				URL:  "/project/settings",
				Icon: "settings",
			},
			{
				ID:           "analytics",
				Text:         "Analytics",
				URL:          "/project/analytics",
				Icon:         "trending-up",
				Badge:        "New",
				BadgeVariant: atoms.BadgeSuccess,
			},
		},

		// Accessibility
		AriaLabel: "Content navigation tabs",
	}
}

// ExampleStepsNavigation demonstrates step navigation
func ExampleStepsNavigation() NavigationProps {
	return NavigationProps{
		ID:          "setup-steps",
		Type:        NavigationSteps,
		Size:        NavigationSizeMD,
		CurrentStep: 2,
		TotalSteps:  4,

		// Step items
		Items: []NavigationItem{
			{
				ID:          "step-1",
				Text:        "Account Setup",
				Description: "Create your account and verify email",
				Icon:        "user",
			},
			{
				ID:          "step-2",
				Text:        "Project Configuration",
				Description: "Configure your first project",
				Icon:        "settings",
			},
			{
				ID:          "step-3",
				Text:        "Team Invitation",
				Description: "Invite team members to collaborate",
				Icon:        "users",
			},
			{
				ID:          "step-4",
				Text:        "Launch",
				Description: "Deploy and go live",
				Icon:        "rocket",
			},
		},

		// Accessibility
		AriaLabel: "Setup progress steps",
	}
}

// ExampleMinimalSidebar demonstrates a minimal sidebar design
func ExampleMinimalSidebar() NavigationProps {
	return NavigationProps{
		ID:          "minimal-sidebar",
		Type:        NavigationSidebar,
		Size:        NavigationSizeSM,
		Variant:     "minimal",
		Collapsible: true,
		Collapsed:   false,

		// Minimal navigation items
		Items: []NavigationItem{
			{
				ID:     "dashboard",
				Icon:   "home",
				URL:    "/dashboard",
				Active: true,
			},
			{
				ID:   "analytics",
				Icon: "bar-chart",
				URL:  "/analytics",
			},
			{
				ID:           "users",
				Icon:         "users",
				URL:          "/users",
				Badge:        "12",
				BadgeVariant: atoms.BadgeDestructive,
			},
			{
				ID:   "settings",
				Icon: "settings",
				URL:  "/settings",
			},
		},

		// Auth context with minimal user
		AuthContext: &AuthContext{
			User: &User{
				ID:       "user-789",
				Username: "user",
				Avatar:   "/assets/avatars/user.jpg",
			},
			Authenticated: true,
		},

		UserAvatar: "/assets/avatars/user.jpg",

		// Accessibility
		AriaLabel: "Minimal navigation",
	}
}

// ExampleCardSidebar demonstrates a card-style sidebar
func ExampleCardSidebar() NavigationProps {
	return NavigationProps{
		ID:          "card-sidebar",
		Type:        NavigationSidebar,
		Size:        NavigationSizeMD,
		Title:       "Project Alpha",
		Description: "Development workspace",
		Variant:     "card",
		Logo:        "/assets/project-logo.svg",

		// Card navigation items with groups
		Items: []NavigationItem{
			{
				ID:          "overview",
				Text:        "Project Overview",
				Description: "Dashboard and summary",
				Icon:        "layout-dashboard",
				URL:         "/project/overview",
				Active:      true,
			},
			{
				ID:      "divider-dev",
				Divider: true,
			},
			{
				ID:           "code",
				Text:         "Code Repository",
				Description:  "Browse and manage code",
				Icon:         "code",
				URL:          "/project/code",
				Badge:        "3 PRs",
				BadgeVariant: atoms.BadgeWarning,
			},
			{
				ID:           "issues",
				Text:         "Issues & Bugs",
				Description:  "Track and resolve issues",
				Icon:         "bug",
				URL:          "/project/issues",
				Badge:        "7",
				BadgeVariant: atoms.BadgeDestructive,
			},
			{
				ID:          "deployments",
				Text:        "Deployments",
				Description: "Manage releases",
				Icon:        "rocket",
				URL:         "/project/deployments",
			},
			{
				ID:      "divider-tools",
				Divider: true,
			},
			{
				ID:          "docs",
				Text:        "Documentation",
				Description: "Project wiki and guides",
				Icon:        "book-open",
				URL:         "/project/docs",
			},
		},

		// Team user context
		AuthContext: &AuthContext{
			User: &User{
				ID:          "user-proj",
				Username:    "teamlead",
				DisplayName: "Sarah Wilson",
				Avatar:      "/assets/avatars/teamlead.jpg",
				Roles:       []string{"team-lead", "developer"},
			},
			Authenticated: true,
		},

		UserName:   "Sarah Wilson",
		UserAvatar: "/assets/avatars/teamlead.jpg",
		UserRole:   "Team Lead",
		UserOnline: true,

		// Project-specific user menu
		UserMenu: []NavigationItem{
			{
				ID:   "project-settings",
				Text: "Project Settings",
				URL:  "/project/settings",
				Icon: "settings",
			},
			{
				ID:   "team-management",
				Text: "Team Management",
				URL:  "/project/team",
				Icon: "users",
			},
			{
				ID:      "divider-project",
				Divider: true,
			},
			{
				ID:          "switch-project",
				Text:        "Switch Project",
				AlpineClick: "showProjectSwitcher()",
				Icon:        "shuffle",
			},
		},

		// Performance
		LazyLoad: true,

		// Accessibility
		AriaLabel: "Project navigation",
	}
}

// ExamplePublicNavigation demonstrates navigation for non-authenticated users
func ExamplePublicNavigation() NavigationProps {
	return NavigationProps{
		ID:      "public-nav",
		Type:    NavigationTopbar,
		Size:    NavigationSizeMD,
		Title:   "Company Website",
		Logo:    "/assets/company-logo.svg",
		LogoURL: "/",

		// Public navigation items
		Items: []NavigationItem{
			{
				ID:     "home",
				Text:   "Home",
				URL:    "/",
				Active: true,
			},
			{
				ID:   "about",
				Text: "About Us",
				URL:  "/about",
			},
			{
				ID:          "products",
				Text:        "Products",
				Collapsible: true,
				Items: []NavigationItem{
					{
						ID:          "product-a",
						Text:        "Product A",
						URL:         "/products/a",
						Description: "Our flagship product",
					},
					{
						ID:          "product-b",
						Text:        "Product B",
						URL:         "/products/b",
						Description: "Enterprise solution",
					},
					{
						ID:   "pricing",
						Text: "Pricing",
						URL:  "/pricing",
					},
				},
			},
			{
				ID:   "blog",
				Text: "Blog",
				URL:  "/blog",
			},
			{
				ID:   "contact",
				Text: "Contact",
				URL:  "/contact",
			},
		},

		// No auth context (public site)
		AuthContext: nil,

		// Call-to-action items in user menu area
		UserMenu: []NavigationItem{
			{
				ID:   "login",
				Text: "Sign In",
				URL:  "/auth/login",
			},
			{
				ID:           "signup",
				Text:         "Get Started",
				URL:          "/auth/signup",
				Badge:        "Free",
				BadgeVariant: atoms.BadgeSuccess,
			},
		},

		// Accessibility
		AriaLabel: "Main website navigation",
	}
}

// ExampleMobileResponsiveNavigation demonstrates responsive navigation patterns
func ExampleMobileResponsiveNavigation() NavigationProps {
	return NavigationProps{
		ID:    "responsive-nav",
		Type:  NavigationTopbar,
		Size:  NavigationSizeMD,
		Title: "Mobile App",
		Logo:  "/assets/mobile-logo.svg",

		// Mobile-optimized navigation
		MobileMenu: NavigationMobileMenu{
			Enabled:     true,
			Position:    "left",
			Overlay:     true,
			CloseButton: true,
			Swipeable:   true,
		},

		// Touch-optimized items
		Items: []NavigationItem{
			{
				ID:     "feed",
				Text:   "Feed",
				Icon:   "rss",
				URL:    "/feed",
				Active: true,
			},
			{
				ID:           "discover",
				Text:         "Discover",
				Icon:         "compass",
				URL:          "/discover",
				Badge:        "3",
				BadgeVariant: atoms.BadgePrimary,
			},
			{
				ID:           "messages",
				Text:         "Messages",
				Icon:         "message-circle",
				URL:          "/messages",
				Badge:        "12",
				BadgeVariant: atoms.BadgeDestructive,
			},
			{
				ID:   "profile",
				Text: "Profile",
				Icon: "user",
				URL:  "/profile",
			},
		},

		// Mobile user context
		AuthContext: &AuthContext{
			User: &User{
				ID:          "mobile-user",
				Username:    "mobileuser",
				DisplayName: "Alex Thompson",
				Avatar:      "/assets/avatars/mobile-user.jpg",
			},
			Authenticated: true,
		},

		UserName:   "Alex",
		UserAvatar: "/assets/avatars/mobile-user.jpg",
		UserOnline: true,

		// Mobile-specific user menu
		UserMenu: []NavigationItem{
			{
				ID:   "edit-profile",
				Text: "Edit Profile",
				URL:  "/profile/edit",
				Icon: "edit",
			},
			{
				ID:           "notifications",
				Text:         "Notifications",
				URL:          "/notifications",
				Icon:         "bell",
				Badge:        "5",
				BadgeVariant: atoms.BadgeDestructive,
			},
			{
				ID:   "settings",
				Text: "Settings",
				URL:  "/settings",
				Icon: "settings",
			},
			{
				ID:      "divider-mobile",
				Divider: true,
			},
			{
				ID:   "help",
				Text: "Help & Support",
				URL:  "/help",
				Icon: "help-circle",
			},
			{
				ID:   "logout",
				Text: "Sign Out",
				URL:  "/auth/logout",
				Icon: "log-out",
			},
		},

		// Search for mobile
		Search: NavigationSearch{
			Enabled:     true,
			Placeholder: "Search...",
			MinLength:   1,
			Debounce:    300,
		},

		// Mobile notifications
		Notifications: []Notification{
			{
				ID:        "mobile-notif-1",
				Type:      "info",
				Title:     "New message",
				Message:   "You have a new message from Sarah",
				Read:      false,
				Icon:      "message-circle",
				URL:       "/messages/sarah",
				CreatedAt: time.Now().Add(-2 * time.Minute),
			},
		},

		// Touch-optimized performance
		LazyLoad: true,

		// Accessibility
		AriaLabel: "Mobile navigation",
		AriaLabels: map[string]string{
			"mobileMenu": "Mobile navigation menu",
			"search":     "Search content",
		},
	}
}

// Helper function to create test auth context
func CreateTestAuthContext(userRole string) *AuthContext {
	var roles []string
	var permissions []string

	switch userRole {
	case "admin":
		roles = []string{"admin", "moderator", "user"}
		permissions = []string{"users:read", "users:write", "content:read", "content:write", "admin:read", "admin:write"}
	case "moderator":
		roles = []string{"moderator", "user"}
		permissions = []string{"users:read", "content:read", "content:write"}
	case "user":
		roles = []string{"user"}
		permissions = []string{"content:read"}
	default:
		return nil
	}

	return &AuthContext{
		User: &User{
			ID:          "test-" + userRole,
			Username:    userRole,
			Email:       userRole + "@example.com",
			DisplayName: "Test " + userRole,
			Avatar:      "/assets/avatars/" + userRole + ".jpg",
			Roles:       roles,
			Permissions: permissions,
		},
		Authenticated: true,
	}
}
