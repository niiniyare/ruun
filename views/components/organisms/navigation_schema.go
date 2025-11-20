package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
)

// NavigationSchemaBuilder provides schema-driven Navigation configuration
type NavigationSchemaBuilder struct {
	schema      *schema.Schema
	config      *NavigationConfig
	menuItems   []NavigationMenuItem
	permissions []NavigationPermission
	themes      []NavigationTheme
	authContext *AuthContext
}

// NavigationConfig holds configuration for schema-driven navigation
type NavigationConfig struct {
	// Display options
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Logo        string `json:"logo,omitempty"`
	LogoURL     string `json:"logoURL,omitempty"`

	// Layout configuration
	Layout   NavigationLayout   `json:"layout"`
	Variant  NavigationVariant  `json:"variant"`
	Size     NavigationSize     `json:"size"`
	Position NavigationPosition `json:"position"`

	// Feature flags
	EnableSearch        bool `json:"enableSearch"`
	EnableBreadcrumbs   bool `json:"enableBreadcrumbs"`
	EnableNotifications bool `json:"enableNotifications"`
	EnableUserMenu      bool `json:"enableUserMenu"`
	EnableMobileMenu    bool `json:"enableMobileMenu"`
	EnableCollapsible   bool `json:"enableCollapsible"`
	EnableReordering    bool `json:"enableReordering"`

	// Search configuration
	SearchConfig SearchConfig `json:"searchConfig"`

	// Responsive behavior
	ResponsiveConfig ResponsiveConfig `json:"responsiveConfig"`

	// Theme configuration
	ThemeConfig ThemeConfig `json:"themeConfig"`

	// Animation configuration
	AnimationConfig AnimationConfig `json:"animationConfig"`

	// Security configuration
	SecurityConfig SecurityConfig `json:"securityConfig"`

	// Performance configuration
	PerformanceConfig PerformanceConfig `json:"performanceConfig"`

	// Accessibility configuration
	AccessibilityConfig AccessibilityConfig `json:"accessibilityConfig"`

	// Custom templates
	CustomTemplates map[string]string `json:"customTemplates,omitempty"`
}

// NavigationLayout defines navigation layout types
type NavigationLayout string

const (
	LayoutSidebarMain NavigationLayout = "sidebar-main" // Sidebar + main content
	LayoutTopbarMain  NavigationLayout = "topbar-main"  // Topbar + main content
	LayoutBoth        NavigationLayout = "both"         // Sidebar + topbar + main
	LayoutTabs        NavigationLayout = "tabs"         // Tab navigation only
	LayoutBreadcrumb  NavigationLayout = "breadcrumb"   // Breadcrumb only
	LayoutSteps       NavigationLayout = "steps"        // Step navigation
	LayoutMega        NavigationLayout = "mega"         // Mega menu
)

// NavigationVariant defines visual variants
type NavigationVariant string

const (
	VariantDefault  NavigationVariant = "default"
	VariantMinimal  NavigationVariant = "minimal"
	VariantSidebar  NavigationVariant = "sidebar"
	VariantTopbar   NavigationVariant = "topbar"
	VariantFloating NavigationVariant = "floating"
	VariantCard     NavigationVariant = "card"
)

// NavigationPosition defines navigation position
type NavigationPosition string

const (
	PositionLeft   NavigationPosition = "left"
	PositionRight  NavigationPosition = "right"
	PositionTop    NavigationPosition = "top"
	PositionBottom NavigationPosition = "bottom"
)

// NavigationMenuItem represents a navigation menu item with schema support
type NavigationMenuItem struct {
	// Core properties
	ID          string       `json:"id"`
	Type        MenuItemType `json:"type"`
	Text        string       `json:"text"`
	Description string       `json:"description,omitempty"`
	URL         string       `json:"url,omitempty"`
	Icon        string       `json:"icon,omitempty"`
	Image       string       `json:"image,omitempty"`

	// Visual properties
	Badge BadgeConfig `json:"badge,omitempty"`
	Color string      `json:"color,omitempty"`
	Order int         `json:"order"`
	Group string      `json:"group,omitempty"`

	// Behavior properties
	Active   bool `json:"active"`
	Disabled bool `json:"disabled"`
	Hidden   bool `json:"hidden"`
	External bool `json:"external"`
	NewTab   bool `json:"newTab"`
	Divider  bool `json:"divider"`

	// Hierarchy
	Items       []NavigationMenuItem `json:"items,omitempty"`
	ParentID    string               `json:"parentId,omitempty"`
	Collapsible bool                 `json:"collapsible"`
	Collapsed   bool                 `json:"collapsed"`
	Depth       int                  `json:"depth"`

	// Permissions and access control
	Permissions []string `json:"permissions,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	Condition   string   `json:"condition,omitempty"`

	// HTMX integration
	HTMXConfig HTMXConfig `json:"htmxConfig,omitempty"`

	// Alpine.js integration
	AlpineConfig AlpineConfig `json:"alpineConfig,omitempty"`

	// Metadata
	Metadata map[string]any `json:"metadata,omitempty"`
	Tags     []string       `json:"tags,omitempty"`
	Category string         `json:"category,omitempty"`

	// Analytics
	TrackingID string          `json:"trackingId,omitempty"`
	Analytics  AnalyticsConfig `json:"analytics,omitempty"`
}

// MenuItemType defines different types of menu items
type MenuItemType string

const (
	MenuItemLink     MenuItemType = "link"
	MenuItemButton   MenuItemType = "button"
	MenuItemDivider  MenuItemType = "divider"
	MenuItemHeading  MenuItemType = "heading"
	MenuItemDropdown MenuItemType = "dropdown"
	MenuItemMega     MenuItemType = "mega"
	MenuItemSearch   MenuItemType = "search"
	MenuItemUser     MenuItemType = "user"
	MenuItemTheme    MenuItemType = "theme"
	MenuItemLanguage MenuItemType = "language"
)

// BadgeConfig configures item badges
type BadgeConfig struct {
	Text     string             `json:"text,omitempty"`
	Count    int                `json:"count,omitempty"`
	Variant  atoms.BadgeVariant `json:"variant"`
	Color    string             `json:"color,omitempty"`
	Icon     string             `json:"icon,omitempty"`
	Position string             `json:"position"` // "top-right", "top-left", etc.
	Animated bool               `json:"animated"`
	Pulse    bool               `json:"pulse"`
}

// HTMXConfig configures HTMX behavior
type HTMXConfig struct {
	Get     string `json:"get,omitempty"`
	Post    string `json:"post,omitempty"`
	Put     string `json:"put,omitempty"`
	Delete  string `json:"delete,omitempty"`
	Target  string `json:"target,omitempty"`
	Swap    string `json:"swap,omitempty"`
	Trigger string `json:"trigger,omitempty"`
	Confirm string `json:"confirm,omitempty"`
}

// AlpineConfig configures Alpine.js behavior
type AlpineConfig struct {
	Click string `json:"click,omitempty"`
	Show  string `json:"show,omitempty"`
	If    string `json:"if,omitempty"`
	Model string `json:"model,omitempty"`
	Data  string `json:"data,omitempty"`
	Init  string `json:"init,omitempty"`
}

// AnalyticsConfig configures analytics tracking
type AnalyticsConfig struct {
	Enabled    bool              `json:"enabled"`
	Events     []AnalyticsEvent  `json:"events,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	Provider   string            `json:"provider,omitempty"` // "google", "mixpanel", etc.
}

// AnalyticsEvent defines trackable events
type AnalyticsEvent struct {
	Name       string            `json:"name"`
	Action     string            `json:"action"` // "click", "view", "hover"
	Properties map[string]string `json:"properties,omitempty"`
}

// NavigationPermission defines permission-based access
type NavigationPermission struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`    // "view", "edit", "delete", etc.
	Resources   []string `json:"resources"`  // menu items, sections, etc.
	Roles       []string `json:"roles"`      // required roles
	Conditions  []string `json:"conditions"` // additional conditions
}

// SearchConfig configures navigation search
type SearchConfig struct {
	Enabled       bool     `json:"enabled"`
	Placeholder   string   `json:"placeholder"`
	MinLength     int      `json:"minLength"`
	MaxResults    int      `json:"maxResults"`
	Debounce      int      `json:"debounce"` // milliseconds
	SearchFields  []string `json:"searchFields"`
	HighlightTerm bool     `json:"highlightTerm"`
	Shortcuts     bool     `json:"shortcuts"` // enable keyboard shortcuts
}

// ResponsiveConfig configures responsive behavior
type ResponsiveConfig struct {
	Breakpoints    map[string]int `json:"breakpoints"`
	MobileFirst    bool           `json:"mobileFirst"`
	CollapseMobile bool           `json:"collapseMobile"`
	HiddenMobile   []string       `json:"hiddenMobile"` // items to hide on mobile
	OnlyMobile     []string       `json:"onlyMobile"`   // items only on mobile
	TouchOptimized bool           `json:"touchOptimized"`
}

// ThemeConfig configures theming
type ThemeConfig struct {
	DefaultTheme    string            `json:"defaultTheme"`
	AvailableThemes []NavigationTheme `json:"availableThemes"`
	AutoDetect      bool              `json:"autoDetect"`
	Persistent      bool              `json:"persistent"`
	SystemTheme     bool              `json:"systemTheme"`
}

// NavigationTheme defines a navigation theme
type NavigationTheme struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`
	Colors      ThemeColors     `json:"colors"`
	Fonts       ThemeFonts      `json:"fonts,omitempty"`
	Spacing     ThemeSpacing    `json:"spacing,omitempty"`
	Borders     ThemeBorders    `json:"borders,omitempty"`
	Shadows     ThemeShadows    `json:"shadows,omitempty"`
	Animations  ThemeAnimations `json:"animations,omitempty"`
}

// ThemeColors defines theme colors
type ThemeColors struct {
	Primary    string `json:"primary"`
	Secondary  string `json:"secondary"`
	Background string `json:"background"`
	Surface    string `json:"surface"`
	Text       string `json:"text"`
	TextMuted  string `json:"textMuted"`
	Border     string `json:"border"`
	Hover      string `json:"hover"`
	Active     string `json:"active"`
	Selected   string `json:"selected"`
}

// ThemeFonts defines theme typography
type ThemeFonts struct {
	Family     string `json:"family"`
	Size       string `json:"size"`
	Weight     string `json:"weight"`
	LineHeight string `json:"lineHeight"`
}

// ThemeSpacing defines theme spacing
type ThemeSpacing struct {
	XS string `json:"xs"`
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
	XL string `json:"xl"`
}

// ThemeBorders defines theme borders
type ThemeBorders struct {
	Width  string `json:"width"`
	Style  string `json:"style"`
	Radius string `json:"radius"`
	Color  string `json:"color"`
}

// ThemeShadows defines theme shadows
type ThemeShadows struct {
	SM string `json:"sm"`
	MD string `json:"md"`
	LG string `json:"lg"`
	XL string `json:"xl"`
}

// ThemeAnimations defines theme animations
type ThemeAnimations struct {
	Duration string `json:"duration"`
	Easing   string `json:"easing"`
	Delay    string `json:"delay"`
}

// AnimationConfig configures animations
type AnimationConfig struct {
	Enabled       bool   `json:"enabled"`
	Duration      string `json:"duration"`
	Easing        string `json:"easing"`
	ReducedMotion bool   `json:"reducedMotion"`
	Prefers       string `json:"prefers"` // "motion", "no-motion"
}

// SecurityConfig configures security features
type SecurityConfig struct {
	CSRFProtection      bool     `json:"csrfProtection"`
	ClickjackProtection bool     `json:"clickjackProtection"`
	ContentSecurity     bool     `json:"contentSecurity"`
	AllowedOrigins      []string `json:"allowedOrigins,omitempty"`
	TrustedDomains      []string `json:"trustedDomains,omitempty"`
}

// PerformanceConfig configures performance features
type PerformanceConfig struct {
	LazyLoad          bool `json:"lazyLoad"`
	Prefetch          bool `json:"prefetch"`
	Cache             bool `json:"cache"`
	CacheDuration     int  `json:"cacheDuration"` // seconds
	Debounce          bool `json:"debounce"`
	VirtualScrolling  bool `json:"virtualScrolling"`
	ImageOptimization bool `json:"imageOptimization"`
}

// AccessibilityConfig configures accessibility features
type AccessibilityConfig struct {
	ARIA            bool              `json:"aria"`
	KeyboardNav     bool              `json:"keyboardNav"`
	FocusManagement bool              `json:"focusManagement"`
	ScreenReader    bool              `json:"screenReader"`
	HighContrast    bool              `json:"highContrast"`
	ReducedMotion   bool              `json:"reducedMotion"`
	CustomLabels    map[string]string `json:"customLabels,omitempty"`
	SkipLinks       bool              `json:"skipLinks"`
}

// AuthContext provides authentication context for navigation
type AuthContext struct {
	User          *User    `json:"user,omitempty"`
	Roles         []string `json:"roles"`
	Permissions   []string `json:"permissions"`
	Authenticated bool     `json:"authenticated"`
	SessionID     string   `json:"sessionId,omitempty"`
	CSRF          string   `json:"csrf,omitempty"`
}

// NewNavigationSchemaBuilder creates a new schema-driven navigation builder
func NewNavigationSchemaBuilder(navSchema *schema.Schema) *NavigationSchemaBuilder {
	return &NavigationSchemaBuilder{
		schema:      navSchema,
		config:      getDefaultNavigationConfig(),
		menuItems:   []NavigationMenuItem{},
		permissions: []NavigationPermission{},
		themes:      getDefaultNavigationThemes(),
	}
}

// getDefaultNavigationConfig returns default configuration
func getDefaultNavigationConfig() *NavigationConfig {
	return &NavigationConfig{
		Layout:              LayoutSidebarMain,
		Variant:             VariantDefault,
		Size:                NavigationSizeMD,
		Position:            PositionLeft,
		EnableSearch:        true,
		EnableBreadcrumbs:   true,
		EnableNotifications: true,
		EnableUserMenu:      true,
		EnableMobileMenu:    true,
		EnableCollapsible:   true,
		EnableReordering:    false,

		SearchConfig: SearchConfig{
			Enabled:       true,
			Placeholder:   "Search navigation...",
			MinLength:     2,
			MaxResults:    10,
			Debounce:      300,
			SearchFields:  []string{"text", "description"},
			HighlightTerm: true,
			Shortcuts:     true,
		},

		ResponsiveConfig: ResponsiveConfig{
			Breakpoints: map[string]int{
				"sm": 640,
				"md": 768,
				"lg": 1024,
				"xl": 1280,
			},
			MobileFirst:    true,
			CollapseMobile: true,
			TouchOptimized: true,
		},

		ThemeConfig: ThemeConfig{
			DefaultTheme: "light",
			AutoDetect:   true,
			Persistent:   true,
			SystemTheme:  true,
		},

		AnimationConfig: AnimationConfig{
			Enabled:       true,
			Duration:      "200ms",
			Easing:        "ease-in-out",
			ReducedMotion: true,
			Prefers:       "motion",
		},

		SecurityConfig: SecurityConfig{
			CSRFProtection:      true,
			ClickjackProtection: true,
			ContentSecurity:     true,
		},

		PerformanceConfig: PerformanceConfig{
			LazyLoad:          true,
			Prefetch:          true,
			Cache:             true,
			CacheDuration:     300,
			Debounce:          true,
			VirtualScrolling:  false,
			ImageOptimization: true,
		},

		AccessibilityConfig: AccessibilityConfig{
			ARIA:            true,
			KeyboardNav:     true,
			FocusManagement: true,
			ScreenReader:    true,
			HighContrast:    true,
			ReducedMotion:   true,
			SkipLinks:       true,
			CustomLabels: map[string]string{
				"navigation":  "Main navigation",
				"search":      "Search navigation",
				"userMenu":    "User menu",
				"mobileMenu":  "Mobile menu",
				"breadcrumbs": "Breadcrumb navigation",
			},
		},
	}
}

// getDefaultNavigationThemes returns default themes
func getDefaultNavigationThemes() []NavigationTheme {
	return []NavigationTheme{
		{
			ID:          "light",
			Name:        "Light Theme",
			Description: "Clean light theme for better readability",
			Colors: ThemeColors{
				Primary:    "#3b82f6",
				Secondary:  "#64748b",
				Background: "#ffffff",
				Surface:    "#f8fafc",
				Text:       "#1e293b",
				TextMuted:  "#64748b",
				Border:     "#e2e8f0",
				Hover:      "#f1f5f9",
				Active:     "#e2e8f0",
				Selected:   "#dbeafe",
			},
		},
		{
			ID:          "dark",
			Name:        "Dark Theme",
			Description: "Dark theme for reduced eye strain",
			Colors: ThemeColors{
				Primary:    "#60a5fa",
				Secondary:  "#94a3b8",
				Background: "#0f172a",
				Surface:    "#1e293b",
				Text:       "#f1f5f9",
				TextMuted:  "#94a3b8",
				Border:     "#334155",
				Hover:      "#334155",
				Active:     "#475569",
				Selected:   "#1e40af",
			},
		},
	}
}

// WithConfig applies custom configuration
func (b *NavigationSchemaBuilder) WithConfig(config *NavigationConfig) *NavigationSchemaBuilder {
	if config != nil {
		b.config = config
	}
	return b
}

// WithAuthContext sets the authentication context
func (b *NavigationSchemaBuilder) WithAuthContext(authCtx *AuthContext) *NavigationSchemaBuilder {
	b.authContext = authCtx
	return b
}

// WithMenuItems adds menu items
func (b *NavigationSchemaBuilder) WithMenuItems(items []NavigationMenuItem) *NavigationSchemaBuilder {
	b.menuItems = items
	return b
}

// AddMenuItem adds a single menu item
func (b *NavigationSchemaBuilder) AddMenuItem(item NavigationMenuItem) *NavigationSchemaBuilder {
	b.menuItems = append(b.menuItems, item)
	return b
}

// WithPermissions adds navigation permissions
func (b *NavigationSchemaBuilder) WithPermissions(permissions []NavigationPermission) *NavigationSchemaBuilder {
	b.permissions = permissions
	return b
}

// WithThemes sets custom themes
func (b *NavigationSchemaBuilder) WithThemes(themes []NavigationTheme) *NavigationSchemaBuilder {
	b.themes = themes
	return b
}

// Build generates the NavigationProps from schema and configuration
func (b *NavigationSchemaBuilder) Build(ctx context.Context) (*NavigationProps, error) {
	if b.schema == nil {
		return nil, fmt.Errorf("schema is required")
	}

	props := &NavigationProps{
		ID:          b.schema.ID,
		Type:        b.mapLayoutToType(b.config.Layout),
		Size:        b.config.Size,
		Title:       getConfigValue(b.config.Title, b.schema.Title),
		Description: getConfigValue(b.config.Description, b.schema.Description),
		Logo:        b.config.Logo,
		LogoURL:     b.config.LogoURL,
		Collapsible: b.config.EnableCollapsible,
		Class:       b.buildNavigationClasses(),
	}

	// Build menu items from schema and configuration
	items, err := b.buildNavigationItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to build navigation items: %w", err)
	}
	props.Items = items

	// Configure user menu if enabled
	if b.config.EnableUserMenu && b.authContext != nil && b.authContext.User != nil {
		userMenu, err := b.buildUserMenu(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to build user menu: %w", err)
		}
		props.UserMenu = userMenu
		props.UserName = b.authContext.User.DisplayName
		props.UserAvatar = b.authContext.User.Avatar
	}

	// Add Alpine.js integration
	props.AlpineData = b.buildAlpineData()

	// Add HTMX configuration
	if b.config.PerformanceConfig.LazyLoad {
		props.HXGet = "/api/navigation/lazy"
		props.HXTarget = "#navigation-content"
		props.HXSwap = "innerHTML"
	}

	return props, nil
}

// buildNavigationItems creates navigation items from schema
func (b *NavigationSchemaBuilder) buildNavigationItems(ctx context.Context) ([]NavigationItem, error) {
	var items []NavigationItem

	// Convert schema menu items to navigation items
	for _, schemaItem := range b.menuItems {
		// Check permissions
		if !b.hasPermission(ctx, schemaItem) {
			continue
		}

		item, err := b.convertSchemaItemToNavItem(schemaItem)
		if err != nil {
			return nil, fmt.Errorf("failed to convert schema item %s: %w", schemaItem.ID, err)
		}

		items = append(items, item)
	}

	// Auto-generate items from schema if no explicit items
	if len(items) == 0 && len(b.schema.Fields) > 0 {
		autoItems, err := b.generateItemsFromSchema(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to generate items from schema: %w", err)
		}
		items = autoItems
	}

	// Sort items by order
	sort.Slice(items, func(i, j int) bool {
		return b.getItemOrder(items[i].ID) < b.getItemOrder(items[j].ID)
	})

	return items, nil
}

// convertSchemaItemToNavItem converts a schema menu item to navigation item
func (b *NavigationSchemaBuilder) convertSchemaItemToNavItem(schemaItem NavigationMenuItem) (NavigationItem, error) {
	item := NavigationItem{
		ID:          schemaItem.ID,
		Text:        schemaItem.Text,
		Description: schemaItem.Description,
		URL:         schemaItem.URL,
		Icon:        schemaItem.Icon,
		Active:      schemaItem.Active,
		Disabled:    schemaItem.Disabled,
		External:    schemaItem.External,
		Divider:     schemaItem.Divider,
		Collapsible: schemaItem.Collapsible,
		Collapsed:   schemaItem.Collapsed,
		Class:       "",
	}

	// Configure badge
	if schemaItem.Badge.Text != "" || schemaItem.Badge.Count > 0 {
		if schemaItem.Badge.Count > 0 {
			item.Badge = fmt.Sprintf("%d", schemaItem.Badge.Count)
		} else {
			item.Badge = schemaItem.Badge.Text
		}
		item.BadgeVariant = schemaItem.Badge.Variant
	}

	// Configure HTMX
	if schemaItem.HTMXConfig.Get != "" {
		item.HXGet = schemaItem.HTMXConfig.Get
	}
	if schemaItem.HTMXConfig.Post != "" {
		item.HXPost = schemaItem.HTMXConfig.Post
	}
	if schemaItem.HTMXConfig.Target != "" {
		item.HXTarget = schemaItem.HTMXConfig.Target
	}
	if schemaItem.HTMXConfig.Swap != "" {
		item.HXSwap = schemaItem.HTMXConfig.Swap
	}

	// Configure Alpine.js
	if schemaItem.AlpineConfig.Click != "" {
		item.AlpineClick = schemaItem.AlpineConfig.Click
	}

	// Process nested items
	if len(schemaItem.Items) > 0 {
		for _, subSchemaItem := range schemaItem.Items {
			if b.hasPermission(context.Background(), subSchemaItem) {
				subItem, err := b.convertSchemaItemToNavItem(subSchemaItem)
				if err != nil {
					return item, fmt.Errorf("failed to convert sub-item %s: %w", subSchemaItem.ID, err)
				}
				item.Items = append(item.Items, subItem)
			}
		}
	}

	return item, nil
}

// generateItemsFromSchema auto-generates navigation items from schema fields
func (b *NavigationSchemaBuilder) generateItemsFromSchema(ctx context.Context) ([]NavigationItem, error) {
	var items []NavigationItem

	// Group schema fields by categories or use predefined structure
	if len(b.schema.Children) > 0 {
		// Use schema children as navigation structure
		for _, child := range b.schema.Children {
			item := NavigationItem{
				ID:          child.ID,
				Text:        getSchemaDisplayName(child),
				Description: child.Description,
				URL:         fmt.Sprintf("/%s", strings.ToLower(child.ID)),
				Icon:        b.getSchemaIcon(child),
			}

			// Add nested items from child fields
			for _, field := range child.Fields {
				if b.shouldIncludeFieldInNav(field) {
					subItem := NavigationItem{
						ID:   fmt.Sprintf("%s_%s", child.ID, field.Name),
						Text: getFieldDisplayName(field),
						URL:  fmt.Sprintf("/%s/%s", strings.ToLower(child.ID), field.Name),
						Icon: b.getFieldIcon(field),
					}
					item.Items = append(item.Items, subItem)
				}
			}

			items = append(items, item)
		}
	} else {
		// Create navigation from top-level fields
		for _, field := range b.schema.Fields {
			if b.shouldIncludeFieldInNav(field) {
				item := NavigationItem{
					ID:   field.Name,
					Text: getFieldDisplayName(field),
					URL:  fmt.Sprintf("/%s", field.Name),
					Icon: b.getFieldIcon(field),
				}
				items = append(items, item)
			}
		}
	}

	return items, nil
}

// buildUserMenu creates user menu items
func (b *NavigationSchemaBuilder) buildUserMenu(ctx context.Context) ([]NavigationItem, error) {
	var userMenuItems []NavigationItem

	// Profile item
	userMenuItems = append(userMenuItems, NavigationItem{
		ID:   "profile",
		Text: "Profile",
		URL:  "/profile",
		Icon: "user",
	})

	// Settings item
	userMenuItems = append(userMenuItems, NavigationItem{
		ID:   "settings",
		Text: "Settings",
		URL:  "/settings",
		Icon: "settings",
	})

	// Theme toggle if enabled
	if len(b.config.ThemeConfig.AvailableThemes) > 1 {
		userMenuItems = append(userMenuItems, NavigationItem{
			ID:          "theme-toggle",
			Text:        "Toggle Theme",
			Icon:        "sun-moon",
			AlpineClick: "toggleTheme()",
		})
	}

	// Logout item
	userMenuItems = append(userMenuItems, NavigationItem{
		ID:      "logout",
		Text:    "Logout",
		URL:     "/auth/logout",
		Icon:    "log-out",
		Divider: true,
	})

	return userMenuItems, nil
}

// Helper methods

func (b *NavigationSchemaBuilder) mapLayoutToType(layout NavigationLayout) NavigationType {
	switch layout {
	case LayoutSidebarMain:
		return NavigationSidebar
	case LayoutTopbarMain:
		return NavigationTopbar
	case LayoutBreadcrumb:
		return NavigationBreadcrumb
	case LayoutTabs:
		return NavigationTabs
	case LayoutSteps:
		return NavigationSteps
	default:
		return NavigationSidebar
	}
}

func (b *NavigationSchemaBuilder) buildNavigationClasses() string {
	var classes []string

	// Base classes
	classes = append(classes, "schema-navigation")

	// Variant classes
	if b.config.Variant != "" {
		classes = append(classes, "nav-variant--"+string(b.config.Variant))
	}

	// Layout classes
	if b.config.Layout != "" {
		classes = append(classes, "nav-layout--"+string(b.config.Layout))
	}

	// Feature classes
	if b.config.EnableSearch {
		classes = append(classes, "nav-searchable")
	}
	if b.config.EnableCollapsible {
		classes = append(classes, "nav-collapsible")
	}
	if b.config.AnimationConfig.Enabled {
		classes = append(classes, "nav-animated")
	}

	return strings.Join(classes, " ")
}

func (b *NavigationSchemaBuilder) buildAlpineData() string {
	config := map[string]any{
		"searchEnabled":    b.config.EnableSearch,
		"searchMinLength":  b.config.SearchConfig.MinLength,
		"searchDebounce":   b.config.SearchConfig.Debounce,
		"animationEnabled": b.config.AnimationConfig.Enabled,
		"responsive":       b.config.ResponsiveConfig,
		"themes":           b.themes,
		"currentTheme":     b.config.ThemeConfig.DefaultTheme,
	}

	configJSON, _ := json.Marshal(config)
	return fmt.Sprintf("navigationSchema(%s)", string(configJSON))
}

func (b *NavigationSchemaBuilder) hasPermission(ctx context.Context, item NavigationMenuItem) bool {
	if b.authContext == nil || !b.authContext.Authenticated {
		// Allow public items when not authenticated
		return len(item.Permissions) == 0 && len(item.Roles) == 0
	}

	// Check roles
	if len(item.Roles) > 0 {
		hasRole := false
		for _, requiredRole := range item.Roles {
			for _, userRole := range b.authContext.Roles {
				if userRole == requiredRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}
		if !hasRole {
			return false
		}
	}

	// Check permissions
	if len(item.Permissions) > 0 {
		hasPermission := false
		for _, requiredPerm := range item.Permissions {
			for _, userPerm := range b.authContext.Permissions {
				if userPerm == requiredPerm {
					hasPermission = true
					break
				}
			}
			if hasPermission {
				break
			}
		}
		if !hasPermission {
			return false
		}
	}

	return true
}

func (b *NavigationSchemaBuilder) getItemOrder(itemID string) int {
	for i, item := range b.menuItems {
		if item.ID == itemID {
			if item.Order != 0 {
				return item.Order
			}
			return i
		}
	}
	return 999999 // Default high value for unordered items
}

func (b *NavigationSchemaBuilder) shouldIncludeFieldInNav(field schema.Field) bool {
	// Skip hidden fields
	if field.Hidden {
		return false
	}

	// Skip certain field types
	switch field.Type {
	case schema.FieldPassword, schema.FieldHidden:
		return false
	default:
		return true
	}
}

func (b *NavigationSchemaBuilder) getSchemaIcon(schema *schema.Schema) string {
	// Map schema types/names to icons
	schemaType := strings.ToLower(schema.Type.String())
	schemaID := strings.ToLower(schema.ID)

	iconMap := map[string]string{
		"user":      "user",
		"users":     "users",
		"product":   "package",
		"products":  "package",
		"order":     "shopping-cart",
		"orders":    "shopping-cart",
		"article":   "file-text",
		"articles":  "file-text",
		"settings":  "settings",
		"admin":     "shield",
		"dashboard": "home",
		"reports":   "bar-chart",
		"analytics": "trending-up",
	}

	for key, icon := range iconMap {
		if strings.Contains(schemaID, key) || strings.Contains(schemaType, key) {
			return icon
		}
	}

	return "circle" // Default icon
}

func (b *NavigationSchemaBuilder) getFieldIcon(field schema.Field) string {
	// Map field types to icons
	switch field.Type {
	case schema.FieldEmail:
		return "mail"
	case schema.FieldDate, schema.FieldDatetime:
		return "calendar"
	case schema.FieldFile:
		return "file"
	case schema.FieldNumber:
		return "hash"
	case schema.FieldCheckbox:
		return "check-square"
	case schema.FieldSelect, schema.FieldRadio:
		return "list"
	default:
		return "edit-3"
	}
}

func getSchemaDisplayName(schema *schema.Schema) string {
	if schema.Title != "" {
		return schema.Title
	}
	return strings.Title(strings.ReplaceAll(schema.ID, "_", " "))
}

func getFieldDisplayName(field schema.Field) string {
	if field.Label != "" {
		return field.Label
	}
	return strings.Title(strings.ReplaceAll(field.Name, "_", " "))
}

func getConfigValue(configValue, schemaValue string) string {
	if configValue != "" {
		return configValue
	}
	return schemaValue
}

// JSON serialization helpers

// ToJSON serializes the configuration to JSON
func (config *NavigationConfig) ToJSON() ([]byte, error) {
	return json.MarshalIndent(config, "", "  ")
}

// FromJSON deserializes configuration from JSON
func (config *NavigationConfig) FromJSON(data []byte) error {
	return json.Unmarshal(data, config)
}

// ToJSON serializes a menu item to JSON
func (item *NavigationMenuItem) ToJSON() ([]byte, error) {
	return json.MarshalIndent(item, "", "  ")
}

// FromJSON deserializes a menu item from JSON
func (item *NavigationMenuItem) FromJSON(data []byte) error {
	return json.Unmarshal(data, item)
}

// Validation helpers

// Validate validates the navigation configuration
func (config *NavigationConfig) Validate() error {
	if config.SearchConfig.MinLength < 1 {
		return fmt.Errorf("searchConfig.minLength must be at least 1")
	}

	if config.SearchConfig.MaxResults < 1 {
		return fmt.Errorf("searchConfig.maxResults must be at least 1")
	}

	if config.SearchConfig.Debounce < 0 {
		return fmt.Errorf("searchConfig.debounce must be non-negative")
	}

	if config.PerformanceConfig.CacheDuration < 0 {
		return fmt.Errorf("performanceConfig.cacheDuration must be non-negative")
	}

	return nil
}

// Validate validates a navigation menu item
func (item *NavigationMenuItem) Validate() error {
	if item.ID == "" {
		return fmt.Errorf("menu item ID is required")
	}

	if item.Text == "" {
		return fmt.Errorf("menu item text is required")
	}

	if item.Type == MenuItemLink && item.URL == "" {
		return fmt.Errorf("link menu items must have a URL")
	}

	// Validate nested items
	for i, subItem := range item.Items {
		if err := subItem.Validate(); err != nil {
			return fmt.Errorf("nested item %d: %w", i, err)
		}
	}

	return nil
}
