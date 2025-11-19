package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
	"github.com/niiniyare/ruun/views/components/utils"
)

// NavigationService provides business logic for navigation operations
type NavigationService struct {
	// Configuration
	DefaultCacheTimeout time.Duration
	MaxMenuDepth       int
	
	// Feature flags
	EnableBreadcrumbs  bool
	EnableSearch       bool
	EnableNotifications bool
	
	// Performance settings
	MaxMenuItems       int
	SearchDebounce     time.Duration
	
	// Auth integration
	AuthProvider       AuthProvider
	PermissionProvider PermissionProvider
}

// AuthProvider interface for user authentication
type AuthProvider interface {
	GetCurrentUser(ctx context.Context) (*User, error)
	IsAuthenticated(ctx context.Context) bool
	HasRole(ctx context.Context, role string) bool
	HasPermission(ctx context.Context, permission string) bool
}

// PermissionProvider interface for permission-based navigation
type PermissionProvider interface {
	CanAccessMenuItem(ctx context.Context, user *User, item *NavigationItem) bool
	GetUserPermissions(ctx context.Context, user *User) ([]string, error)
	GetVisibleMenuItems(ctx context.Context, user *User, items []NavigationItem) ([]NavigationItem, error)
}

// User represents the current user
type User struct {
	ID          string   `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	Avatar      string   `json:"avatar,omitempty"`
	DisplayName string   `json:"displayName"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	Preferences map[string]any `json:"preferences,omitempty"`
	LastActive  time.Time `json:"lastActive"`
}

// NavigationContext provides context for navigation rendering
type NavigationContext struct {
	User            *User                  `json:"user,omitempty"`
	CurrentPath     string                 `json:"currentPath"`
	Breadcrumbs     []BreadcrumbItem       `json:"breadcrumbs"`
	ActiveItems     []string               `json:"activeItems"`
	NotificationCounts map[string]int      `json:"notificationCounts,omitempty"`
	Preferences     NavigationPreferences  `json:"preferences"`
	Permissions     []string               `json:"permissions"`
	
	// Request context
	RequestID       string                 `json:"requestId,omitempty"`
	SessionID       string                 `json:"sessionId,omitempty"`
	UserAgent       string                 `json:"userAgent,omitempty"`
	IPAddress       string                 `json:"ipAddress,omitempty"`
}

// NavigationPreferences represents user navigation preferences
type NavigationPreferences struct {
	SidebarCollapsed   bool              `json:"sidebarCollapsed"`
	Theme             string            `json:"theme"` // "light", "dark", "auto"
	Language          string            `json:"language"`
	MenuOrder         []string          `json:"menuOrder,omitempty"`
	HiddenMenuItems   []string          `json:"hiddenMenuItems,omitempty"`
	FavoriteMenuItems []string          `json:"favoriteMenuItems,omitempty"`
	RecentMenuItems   []RecentMenuItem  `json:"recentMenuItems,omitempty"`
	
	// Advanced preferences
	ShowIcons         bool              `json:"showIcons"`
	ShowBadges        bool              `json:"showBadges"`
	ShowDescriptions  bool              `json:"showDescriptions"`
	CompactMode       bool              `json:"compactMode"`
	AnimationEnabled  bool              `json:"animationEnabled"`
}

// RecentMenuItem represents a recently accessed menu item
type RecentMenuItem struct {
	ID          string    `json:"id"`
	Text        string    `json:"text"`
	URL         string    `json:"url"`
	Icon        string    `json:"icon,omitempty"`
	AccessedAt  time.Time `json:"accessedAt"`
}

// BreadcrumbItem represents a breadcrumb navigation item
type BreadcrumbItem struct {
	ID          string `json:"id"`
	Text        string `json:"text"`
	URL         string `json:"url,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Active      bool   `json:"active"`
	Clickable   bool   `json:"clickable"`
}

// NavigationState represents the current state of navigation
type NavigationState struct {
	// Core state
	ActivePath      string            `json:"activePath"`
	CollapsedItems  map[string]bool   `json:"collapsedItems"`
	SearchQuery     string            `json:"searchQuery"`
	SearchResults   []NavigationItem  `json:"searchResults"`
	
	// Mobile state
	MobileMenuOpen  bool              `json:"mobileMenuOpen"`
	UserMenuOpen    bool              `json:"userMenuOpen"`
	SearchOpen      bool              `json:"searchOpen"`
	
	// Sidebar state
	SidebarWidth    int               `json:"sidebarWidth"`
	SidebarCollapsed bool             `json:"sidebarCollapsed"`
	SidebarPinned   bool              `json:"sidebarPinned"`
	
	// Loading states
	Loading         bool              `json:"loading"`
	SearchLoading   bool              `json:"searchLoading"`
	
	// Notification state
	Notifications   []Notification    `json:"notifications"`
	UnreadCount     int               `json:"unreadCount"`
	
	// Cache state
	LastUpdated     time.Time         `json:"lastUpdated"`
	CacheValid      bool              `json:"cacheValid"`
}

// Notification represents a navigation notification
type Notification struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"` // "info", "warning", "error", "success"
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Icon      string    `json:"icon,omitempty"`
	URL       string    `json:"url,omitempty"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewNavigationService creates a new navigation service
func NewNavigationService() *NavigationService {
	return &NavigationService{
		DefaultCacheTimeout: 5 * time.Minute,
		MaxMenuDepth:       5,
		EnableBreadcrumbs:  true,
		EnableSearch:       true,
		EnableNotifications: true,
		MaxMenuItems:       1000,
		SearchDebounce:     300 * time.Millisecond,
	}
}

// BuildNavigationContext creates navigation context for the current request
func (s *NavigationService) BuildNavigationContext(ctx context.Context, currentPath string) (*NavigationContext, error) {
	navCtx := &NavigationContext{
		CurrentPath:        currentPath,
		ActiveItems:        []string{},
		NotificationCounts: make(map[string]int),
		Permissions:        []string{},
		Preferences: NavigationPreferences{
			SidebarCollapsed:  false,
			Theme:            "auto",
			Language:         "en",
			ShowIcons:        true,
			ShowBadges:       true,
			ShowDescriptions: true,
			CompactMode:      false,
			AnimationEnabled: true,
		},
	}
	
	// Get current user if auth provider available
	if s.AuthProvider != nil {
		if user, err := s.AuthProvider.GetCurrentUser(ctx); err == nil && user != nil {
			navCtx.User = user
			
			// Get user permissions
			if s.PermissionProvider != nil {
				if permissions, err := s.PermissionProvider.GetUserPermissions(ctx, user); err == nil {
					navCtx.Permissions = permissions
				}
			}
			
			// Load user preferences
			if prefs, ok := user.Preferences["navigation"]; ok {
				if prefsData, ok := prefs.(map[string]any); ok {
					s.loadUserPreferences(&navCtx.Preferences, prefsData)
				}
			}
		}
	}
	
	// Build breadcrumbs
	breadcrumbs, err := s.BuildBreadcrumbs(ctx, currentPath)
	if err != nil {
		return nil, fmt.Errorf("failed to build breadcrumbs: %w", err)
	}
	navCtx.Breadcrumbs = breadcrumbs
	
	// Determine active menu items
	navCtx.ActiveItems = s.getActiveMenuItems(currentPath)
	
	return navCtx, nil
}

// ProcessNavigationItems processes navigation items with auth and permissions
func (s *NavigationService) ProcessNavigationItems(ctx context.Context, items []NavigationItem, navCtx *NavigationContext) ([]NavigationItem, error) {
	// Apply permission filtering
	visibleItems, err := s.filterByPermissions(ctx, items, navCtx)
	if err != nil {
		return nil, fmt.Errorf("failed to filter by permissions: %w", err)
	}
	
	// Apply user preferences
	processedItems := s.applyUserPreferences(visibleItems, navCtx.Preferences)
	
	// Set active states
	processedItems = s.setActiveStates(processedItems, navCtx.ActiveItems)
	
	// Add notification counts
	processedItems = s.addNotificationCounts(processedItems, navCtx.NotificationCounts)
	
	// Sort items by order and preferences
	processedItems = s.sortMenuItems(processedItems, navCtx.Preferences.MenuOrder)
	
	return processedItems, nil
}

// BuildBreadcrumbs creates breadcrumb navigation
func (s *NavigationService) BuildBreadcrumbs(ctx context.Context, currentPath string) ([]BreadcrumbItem, error) {
	if !s.EnableBreadcrumbs {
		return []BreadcrumbItem{}, nil
	}
	
	var breadcrumbs []BreadcrumbItem
	
	// Split path into segments
	segments := strings.Split(strings.Trim(currentPath, "/"), "/")
	currentURL := ""
	
	// Always add home
	breadcrumbs = append(breadcrumbs, BreadcrumbItem{
		ID:        "home",
		Text:      "Home",
		URL:       "/",
		Icon:      "home",
		Clickable: true,
	})
	
	// Add path segments
	for i, segment := range segments {
		if segment == "" {
			continue
		}
		
		currentURL += "/" + segment
		isLast := i == len(segments)-1
		
		breadcrumb := BreadcrumbItem{
			ID:        fmt.Sprintf("breadcrumb-%d", i),
			Text:      s.humanizePathSegment(segment),
			URL:       currentURL,
			Clickable: !isLast,
			Active:    isLast,
		}
		
		breadcrumbs = append(breadcrumbs, breadcrumb)
	}
	
	return breadcrumbs, nil
}

// SearchMenuItems searches through navigation items
func (s *NavigationService) SearchMenuItems(ctx context.Context, query string, items []NavigationItem, navCtx *NavigationContext) ([]NavigationItem, error) {
	if !s.EnableSearch || len(query) < 2 {
		return []NavigationItem{}, nil
	}
	
	query = strings.ToLower(strings.TrimSpace(query))
	var results []NavigationItem
	
	// Recursively search through all items
	s.searchRecursive(items, query, &results, navCtx)
	
	// Sort results by relevance
	sort.Slice(results, func(i, j int) bool {
		scoreI := s.calculateSearchScore(results[i], query)
		scoreJ := s.calculateSearchScore(results[j], query)
		return scoreI > scoreJ
	})
	
	// Limit results
	maxResults := 10
	if len(results) > maxResults {
		results = results[:maxResults]
	}
	
	return results, nil
}

// UpdateNotificationCounts updates notification counts for menu items
func (s *NavigationService) UpdateNotificationCounts(ctx context.Context, counts map[string]int, navCtx *NavigationContext) error {
	navCtx.NotificationCounts = counts
	return nil
}

// SaveUserPreferences saves user navigation preferences
func (s *NavigationService) SaveUserPreferences(ctx context.Context, preferences NavigationPreferences) error {
	if navCtx, ok := ctx.Value("navigationContext").(*NavigationContext); ok && navCtx.User != nil {
		// Save preferences to user storage
		// This would typically save to database or user service
		navCtx.Preferences = preferences
		
		// Update recent menu items
		s.updateRecentMenuItems(navCtx, preferences.RecentMenuItems)
	}
	
	return nil
}

// GetNavigationState gets the current navigation state
func (s *NavigationService) GetNavigationState(ctx context.Context) (*NavigationState, error) {
	state := &NavigationState{
		CollapsedItems:   make(map[string]bool),
		SearchResults:    []NavigationItem{},
		MobileMenuOpen:   false,
		UserMenuOpen:     false,
		SearchOpen:       false,
		SidebarWidth:     256,
		SidebarCollapsed: false,
		SidebarPinned:    true,
		Loading:          false,
		SearchLoading:    false,
		Notifications:    []Notification{},
		UnreadCount:      0,
		LastUpdated:      time.Now(),
		CacheValid:       true,
	}
	
	return state, nil
}

// Helper methods

func (s *NavigationService) loadUserPreferences(prefs *NavigationPreferences, data map[string]any) {
	if val, ok := data["sidebarCollapsed"].(bool); ok {
		prefs.SidebarCollapsed = val
	}
	if val, ok := data["theme"].(string); ok {
		prefs.Theme = val
	}
	if val, ok := data["language"].(string); ok {
		prefs.Language = val
	}
	if val, ok := data["showIcons"].(bool); ok {
		prefs.ShowIcons = val
	}
	if val, ok := data["showBadges"].(bool); ok {
		prefs.ShowBadges = val
	}
	if val, ok := data["compactMode"].(bool); ok {
		prefs.CompactMode = val
	}
}

func (s *NavigationService) filterByPermissions(ctx context.Context, items []NavigationItem, navCtx *NavigationContext) ([]NavigationItem, error) {
	if s.PermissionProvider == nil || navCtx.User == nil {
		return items, nil
	}
	
	return s.PermissionProvider.GetVisibleMenuItems(ctx, navCtx.User, items)
}

func (s *NavigationService) applyUserPreferences(items []NavigationItem, prefs NavigationPreferences) []NavigationItem {
	var result []NavigationItem
	
	for _, item := range items {
		// Skip hidden items
		if s.isItemHidden(item.ID, prefs.HiddenMenuItems) {
			continue
		}
		
		// Apply preference overrides
		processedItem := item
		if !prefs.ShowIcons {
			processedItem.Icon = ""
		}
		if !prefs.ShowBadges {
			processedItem.Badge = ""
		}
		if !prefs.ShowDescriptions {
			processedItem.Description = ""
		}
		
		// Process nested items
		if len(processedItem.Items) > 0 {
			processedItem.Items = s.applyUserPreferences(processedItem.Items, prefs)
		}
		
		result = append(result, processedItem)
	}
	
	return result
}

func (s *NavigationService) setActiveStates(items []NavigationItem, activeItems []string) []NavigationItem {
	var result []NavigationItem
	
	for _, item := range items {
		processedItem := item
		processedItem.Active = s.isItemActive(item.ID, activeItems)
		
		// Process nested items
		if len(processedItem.Items) > 0 {
			processedItem.Items = s.setActiveStates(processedItem.Items, activeItems)
		}
		
		result = append(result, processedItem)
	}
	
	return result
}

func (s *NavigationService) addNotificationCounts(items []NavigationItem, counts map[string]int) []NavigationItem {
	var result []NavigationItem
	
	for _, item := range items {
		processedItem := item
		
		if count, exists := counts[item.ID]; exists && count > 0 {
			processedItem.Badge = fmt.Sprintf("%d", count)
			processedItem.BadgeVariant = atoms.BadgeDestructive
		}
		
		// Process nested items
		if len(processedItem.Items) > 0 {
			processedItem.Items = s.addNotificationCounts(processedItem.Items, counts)
		}
		
		result = append(result, processedItem)
	}
	
	return result
}

func (s *NavigationService) sortMenuItems(items []NavigationItem, order []string) []NavigationItem {
	if len(order) == 0 {
		return items
	}
	
	// Create order map for O(1) lookup
	orderMap := make(map[string]int)
	for i, id := range order {
		orderMap[id] = i
	}
	
	// Sort items
	sort.Slice(items, func(i, j int) bool {
		orderI, existsI := orderMap[items[i].ID]
		orderJ, existsJ := orderMap[items[j].ID]
		
		if existsI && existsJ {
			return orderI < orderJ
		}
		if existsI {
			return true
		}
		if existsJ {
			return false
		}
		
		// Fall back to alphabetical
		return items[i].Text < items[j].Text
	})
	
	// Sort nested items
	for i := range items {
		if len(items[i].Items) > 0 {
			items[i].Items = s.sortMenuItems(items[i].Items, order)
		}
	}
	
	return items
}

func (s *NavigationService) getActiveMenuItems(currentPath string) []string {
	var activeItems []string
	
	// Simple path-based matching
	if currentPath == "/" {
		activeItems = append(activeItems, "home")
	} else {
		// Extract route segments for matching
		segments := strings.Split(strings.Trim(currentPath, "/"), "/")
		for i, segment := range segments {
			if segment != "" {
				// Add both the segment and progressive path as active
				activeItems = append(activeItems, segment)
				if i < len(segments)-1 {
					activeItems = append(activeItems, strings.Join(segments[:i+1], "_"))
				}
			}
		}
	}
	
	return activeItems
}

func (s *NavigationService) humanizePathSegment(segment string) string {
	// Convert path segment to human-readable text
	humanized := strings.ReplaceAll(segment, "-", " ")
	humanized = strings.ReplaceAll(humanized, "_", " ")
	humanized = strings.Title(humanized)
	return humanized
}

func (s *NavigationService) searchRecursive(items []NavigationItem, query string, results *[]NavigationItem, navCtx *NavigationContext) {
	for _, item := range items {
		// Check if current user can access this item
		if navCtx.User != nil && s.PermissionProvider != nil {
			if canAccess, err := s.PermissionProvider.CanAccessMenuItem(context.Background(), navCtx.User, &item); err != nil || !canAccess {
				continue
			}
		}
		
		// Check if item matches search
		if s.itemMatchesSearch(item, query) {
			*results = append(*results, item)
		}
		
		// Search nested items
		if len(item.Items) > 0 {
			s.searchRecursive(item.Items, query, results, navCtx)
		}
	}
}

func (s *NavigationService) itemMatchesSearch(item NavigationItem, query string) bool {
	text := strings.ToLower(item.Text)
	description := strings.ToLower(item.Description)
	
	return strings.Contains(text, query) || strings.Contains(description, query)
}

func (s *NavigationService) calculateSearchScore(item NavigationItem, query string) int {
	score := 0
	text := strings.ToLower(item.Text)
	description := strings.ToLower(item.Description)
	
	// Exact match in title gets highest score
	if text == query {
		score += 100
	} else if strings.HasPrefix(text, query) {
		score += 50
	} else if strings.Contains(text, query) {
		score += 20
	}
	
	// Description matches get lower score
	if strings.Contains(description, query) {
		score += 10
	}
	
	return score
}

func (s *NavigationService) isItemHidden(itemID string, hiddenItems []string) bool {
	for _, hiddenID := range hiddenItems {
		if hiddenID == itemID {
			return true
		}
	}
	return false
}

func (s *NavigationService) isItemActive(itemID string, activeItems []string) bool {
	for _, activeID := range activeItems {
		if activeID == itemID {
			return true
		}
	}
	return false
}

func (s *NavigationService) updateRecentMenuItems(navCtx *NavigationContext, recentItems []RecentMenuItem) {
	// Update recent items list (limit to last 10)
	maxRecent := 10
	if len(recentItems) > maxRecent {
		navCtx.Preferences.RecentMenuItems = recentItems[:maxRecent]
	} else {
		navCtx.Preferences.RecentMenuItems = recentItems
	}
}

// Alpine.js data generation for navigation state management
func GetNavigationAlpineData(props NavigationProps, navCtx *NavigationContext) string {
	state := NavigationState{
		ActivePath:       navCtx.CurrentPath,
		CollapsedItems:   make(map[string]bool),
		SearchQuery:      "",
		SearchResults:    []NavigationItem{},
		MobileMenuOpen:   false,
		UserMenuOpen:     false,
		SearchOpen:       false,
		SidebarWidth:     256,
		SidebarCollapsed: navCtx.Preferences.SidebarCollapsed,
		SidebarPinned:    true,
		Loading:          false,
		SearchLoading:    false,
		Notifications:    []Notification{},
		UnreadCount:      0,
		LastUpdated:      time.Now(),
		CacheValid:       true,
	}
	
	// Initialize collapsed items based on navigation structure
	for _, item := range props.Items {
		if item.Collapsible {
			state.CollapsedItems[item.ID] = item.Collapsed
		}
	}
	
	stateJSON, _ := json.Marshal(state)
	prefsJSON, _ := json.Marshal(navCtx.Preferences)
	
	return fmt.Sprintf(`{
		// Core navigation state
		...(%s),
		
		// User preferences
		preferences: %s,
		
		// Computed properties
		get isMobile() {
			return window.innerWidth < 768;
		},
		
		get sidebarClasses() {
			let classes = 'transition-all duration-300 ease-in-out';
			if (this.sidebarCollapsed) {
				classes += ' w-16';
			} else {
				classes += ' w-64';
			}
			return classes;
		},
		
		get searchResultsVisible() {
			return this.searchQuery.length >= 2 && this.searchResults.length > 0;
		},
		
		// Navigation methods
		toggleMobileMenu() {
			this.mobileMenuOpen = !this.mobileMenuOpen;
			
			// Close other menus
			this.userMenuOpen = false;
			this.searchOpen = false;
		},
		
		toggleUserMenu() {
			this.userMenuOpen = !this.userMenuOpen;
			
			// Close other menus
			this.mobileMenuOpen = false;
			this.searchOpen = false;
		},
		
		toggleSearch() {
			this.searchOpen = !this.searchOpen;
			
			if (this.searchOpen) {
				// Focus search input
				this.$nextTick(() => {
					const input = document.querySelector('[data-search-input]');
					if (input) input.focus();
				});
			} else {
				this.searchQuery = '';
				this.searchResults = [];
			}
			
			// Close other menus
			this.mobileMenuOpen = false;
			this.userMenuOpen = false;
		},
		
		toggleSidebar() {
			this.sidebarCollapsed = !this.sidebarCollapsed;
			this.preferences.sidebarCollapsed = this.sidebarCollapsed;
			this.savePreferences();
		},
		
		toggleMenuItem(itemId) {
			this.collapsedItems[itemId] = !this.collapsedItems[itemId];
		},
		
		// Search functionality
		async performSearch() {
			if (this.searchQuery.length < 2) {
				this.searchResults = [];
				return;
			}
			
			this.searchLoading = true;
			
			try {
				const response = await fetch('/api/navigation/search', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'X-Requested-With': 'XMLHttpRequest'
					},
					body: JSON.stringify({
						query: this.searchQuery
					})
				});
				
				if (response.ok) {
					const data = await response.json();
					this.searchResults = data.results || [];
				}
			} catch (error) {
				console.error('Search failed:', error);
			} finally {
				this.searchLoading = false;
			}
		},
		
		// Navigation actions
		navigateToItem(item) {
			// Track recent items
			this.addToRecentItems(item);
			
			// Close mobile menu if open
			if (this.isMobile) {
				this.mobileMenuOpen = false;
			}
			
			// Close search if open
			this.searchOpen = false;
			this.searchQuery = '';
			this.searchResults = [];
			
			// Update active path
			this.activePath = item.url || '';
		},
		
		addToRecentItems(item) {
			const recentItem = {
				id: item.id,
				text: item.text,
				url: item.url,
				icon: item.icon,
				accessedAt: new Date().toISOString()
			};
			
			// Remove if already exists
			this.preferences.recentMenuItems = this.preferences.recentMenuItems.filter(
				recent => recent.id !== item.id
			);
			
			// Add to beginning
			this.preferences.recentMenuItems.unshift(recentItem);
			
			// Keep only last 10
			this.preferences.recentMenuItems = this.preferences.recentMenuItems.slice(0, 10);
			
			this.savePreferences();
		},
		
		// Notification methods
		markNotificationAsRead(notificationId) {
			const notification = this.notifications.find(n => n.id === notificationId);
			if (notification && !notification.read) {
				notification.read = true;
				this.unreadCount = Math.max(0, this.unreadCount - 1);
				
				// Sync with server
				this.syncNotificationState(notificationId, true);
			}
		},
		
		async syncNotificationState(notificationId, read) {
			try {
				await fetch('/api/notifications/' + notificationId, {
					method: 'PATCH',
					headers: {
						'Content-Type': 'application/json',
						'X-Requested-With': 'XMLHttpRequest'
					},
					body: JSON.stringify({ read })
				});
			} catch (error) {
				console.error('Failed to sync notification state:', error);
			}
		},
		
		// Preference management
		savePreferences() {
			// Debounce preference saving
			clearTimeout(this.preferencesTimeout);
			this.preferencesTimeout = setTimeout(() => {
				fetch('/api/navigation/preferences', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json',
						'X-Requested-With': 'XMLHttpRequest'
					},
					body: JSON.stringify(this.preferences)
				}).catch(error => {
					console.error('Failed to save preferences:', error);
				});
			}, 1000);
		},
		
		// Utility methods
		isItemActive(item) {
			if (!item.url) return false;
			
			// Exact match
			if (item.url === this.activePath) {
				return true;
			}
			
			// Parent path match for nested routes
			if (this.activePath.startsWith(item.url + '/')) {
				return true;
			}
			
			// Check nested items
			if (item.items && item.items.length > 0) {
				return item.items.some(subItem => this.isItemActive(subItem));
			}
			
			return false;
		},
		
		hasVisibleChildren(item) {
			return item.items && item.items.length > 0 && 
				   item.items.some(child => this.isItemVisible(child));
		},
		
		isItemVisible(item) {
			// Check permissions, hidden status, etc.
			if (this.preferences.hiddenMenuItems.includes(item.id)) {
				return false;
			}
			
			// Additional visibility checks would go here
			return true;
		},
		
		// Responsive behavior
		handleResize() {
			if (this.isMobile) {
				this.sidebarCollapsed = true;
				this.mobileMenuOpen = false;
			}
		},
		
		// Keyboard navigation
		handleKeydown(event) {
			switch (event.key) {
				case 'Escape':
					this.mobileMenuOpen = false;
					this.userMenuOpen = false;
					this.searchOpen = false;
					break;
					
				case '/':
					if (!this.searchOpen && !event.target.matches('input, textarea')) {
						event.preventDefault();
						this.toggleSearch();
					}
					break;
					
				case 'b':
					if (event.ctrlKey || event.metaKey) {
						event.preventDefault();
						this.toggleSidebar();
					}
					break;
			}
		},
		
		// Initialization
		init() {
			// Set up search debouncing
			this.$watch('searchQuery', () => {
				clearTimeout(this.searchTimeout);
				this.searchTimeout = setTimeout(() => {
					this.performSearch();
				}, 300);
			});
			
			// Set up keyboard navigation
			document.addEventListener('keydown', this.handleKeydown.bind(this));
			
			// Set up responsive behavior
			window.addEventListener('resize', this.handleResize.bind(this));
			this.handleResize();
			
			// Load initial notifications
			this.loadNotifications();
		},
		
		async loadNotifications() {
			try {
				const response = await fetch('/api/notifications', {
					headers: {
						'X-Requested-With': 'XMLHttpRequest'
					}
				});
				
				if (response.ok) {
					const data = await response.json();
					this.notifications = data.notifications || [];
					this.unreadCount = data.unreadCount || 0;
				}
			} catch (error) {
				console.error('Failed to load notifications:', error);
			}
		}
	}`,
		string(stateJSON),
		string(prefsJSON),
	)
}

// Class generation functions for navigation components

func GetNavigationClasses(props NavigationProps, navCtx *NavigationContext) string {
	classes := []string{"navigation-container"}
	
	// Base type classes
	switch props.Type {
	case NavigationSidebar:
		classes = append(classes, "nav-sidebar")
		if navCtx.Preferences.SidebarCollapsed {
			classes = append(classes, "nav-sidebar--collapsed")
		}
		if navCtx.Preferences.CompactMode {
			classes = append(classes, "nav-sidebar--compact")
		}
		
	case NavigationTopbar:
		classes = append(classes, "nav-topbar")
		
	case NavigationBreadcrumb:
		classes = append(classes, "nav-breadcrumb")
		
	case NavigationTabs:
		classes = append(classes, "nav-tabs")
		
	case NavigationSteps:
		classes = append(classes, "nav-steps")
	}
	
	// Size classes
	switch props.Size {
	case NavigationSizeSM:
		classes = append(classes, "nav-size--sm")
	case NavigationSizeLG:
		classes = append(classes, "nav-size--lg")
	default:
		classes = append(classes, "nav-size--md")
	}
	
	// Theme classes
	if navCtx.Preferences.Theme != "" {
		classes = append(classes, "nav-theme--"+navCtx.Preferences.Theme)
	}
	
	// Animation classes
	if navCtx.Preferences.AnimationEnabled {
		classes = append(classes, "nav-animated")
	}
	
	// Custom classes
	if props.Class != "" {
		classes = append(classes, props.Class)
	}
	
	return strings.Join(classes, " ")
}

func GetMenuItemClasses(item NavigationItem, depth int, navCtx *NavigationContext) string {
	classes := []string{"nav-menu-item"}
	
	// State classes
	if item.Active {
		classes = append(classes, "nav-menu-item--active")
	}
	if item.Disabled {
		classes = append(classes, "nav-menu-item--disabled")
	}
	
	// Depth classes
	if depth > 0 {
		classes = append(classes, fmt.Sprintf("nav-menu-item--depth-%d", depth))
	}
	
	// Type classes
	if len(item.Items) > 0 {
		classes = append(classes, "nav-menu-item--parent")
		if item.Collapsed {
			classes = append(classes, "nav-menu-item--collapsed")
		} else {
			classes = append(classes, "nav-menu-item--expanded")
		}
	} else {
		classes = append(classes, "nav-menu-item--leaf")
	}
	
	// User preference classes
	if navCtx.Preferences.CompactMode {
		classes = append(classes, "nav-menu-item--compact")
	}
	
	// Custom classes
	if item.Class != "" {
		classes = append(classes, item.Class)
	}
	
	return strings.Join(classes, " ")
}