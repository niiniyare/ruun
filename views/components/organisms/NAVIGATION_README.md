# Enhanced Navigation Organism

The Navigation organism provides a comprehensive, enterprise-grade navigation system with authentication integration, role-based access control, and advanced features like search, notifications, and responsive design.

## Architecture Overview

The Enhanced Navigation organism follows the same architectural principles as the Form and DataTable organisms:

- **Main Template** (`navigation.templ`): Core navigation rendering with multiple layout types
- **Business Logic** (`navigation_logic.go`): Auth integration, permission handling, and navigation state management  
- **Schema Configuration** (`navigation_schema.go`): Schema-driven navigation structure and configuration
- **Enhanced Components** (`navigation_enhanced_components.templ`): Specialized components for advanced features
- **Examples** (`navigation_examples.go`): Comprehensive usage examples and patterns

## Key Features

### 🔐 Authentication & Authorization
- Role-based menu item visibility
- Permission-based access control
- User context integration
- Session management
- CSRF protection

### 🎨 Multiple Navigation Types
- **Sidebar Navigation**: Collapsible sidebar with nested menus
- **Topbar Navigation**: Horizontal navigation bar
- **Breadcrumb Navigation**: Path-based navigation
- **Tab Navigation**: Content tabs with variants
- **Step Navigation**: Progress-based navigation

### 🔍 Advanced Search
- Real-time search with debouncing
- Search highlighting
- Keyboard shortcuts (/ to open)
- Search result ranking
- Searchable navigation items

### 📱 Mobile Responsive
- Touch-optimized interface
- Swipeable mobile menu
- Adaptive layouts
- Mobile-specific interactions
- Responsive breakpoints

### 🔔 Notification System
- Real-time notifications
- Unread count badges
- Toast notifications
- Notification types (info, warning, error, success)
- Notification history

### ⚡ Performance Features
- Lazy loading support
- Data caching
- Virtual scrolling (for large menus)
- Debounced search
- Image optimization

### ♿ Accessibility
- ARIA labels and roles
- Keyboard navigation
- Screen reader support
- Focus management
- High contrast support
- Reduced motion preferences

### 🎯 Theme Integration
- Compiled theme classes
- Dark/light mode support
- Theme switching
- Custom color schemes
- Runtime theme overrides

## Quick Start

### Basic Sidebar Navigation

```go
func BasicSidebarExample() NavigationProps {
    return NavigationProps{
        Type:        NavigationSidebar,
        Title:       "My App",
        Logo:        "/logo.svg",
        Collapsible: true,
        
        Items: []NavigationItem{
            {
                ID:   "dashboard",
                Text: "Dashboard", 
                Icon: "home",
                URL:  "/dashboard",
                Active: true,
            },
            {
                ID:   "users",
                Text: "Users",
                Icon: "users", 
                URL:  "/users",
            },
        },
    }
}
```

### With Authentication

```go
func AuthenticatedNavigation() NavigationProps {
    return NavigationProps{
        Type:  NavigationSidebar,
        Title: "Admin Panel",
        
        // Auth context
        AuthContext: &AuthContext{
            User: &User{
                ID:          "admin-123",
                DisplayName: "John Admin",
                Email:       "admin@example.com",
                Roles:       []string{"admin"},
                Permissions: []string{"users:read", "users:write"},
            },
            Authenticated: true,
        },
        
        // User interface
        UserName:   "John Admin",
        UserAvatar: "/avatars/admin.jpg",
        UserMenu: []NavigationItem{
            {ID: "profile", Text: "Profile", URL: "/profile", Icon: "user"},
            {ID: "logout", Text: "Logout", URL: "/logout", Icon: "log-out"},
        },
        
        Items: []NavigationItem{
            {
                ID:          "users", 
                Text:        "User Management",
                Icon:        "users",
                Permissions: []string{"users:read"}, // Requires permission
                Items: []NavigationItem{
                    {
                        ID:          "users-list",
                        Text:        "All Users", 
                        URL:         "/users",
                        Permissions: []string{"users:read"},
                    },
                    {
                        ID:          "users-create",
                        Text:        "Add User",
                        URL:         "/users/create", 
                        Permissions: []string{"users:write"},
                    },
                },
            },
        },
    }
}
```

### Schema-Driven Navigation

```go
func SchemaBasedNavigation() (*NavigationProps, error) {
    // Load navigation schema
    navSchema := &schema.Schema{
        ID:    "main-navigation",
        Title: "Application Navigation",
        Type:  schema.TypeNavigation,
    }
    
    // Configure navigation
    config := &NavigationConfig{
        EnableSearch:        true,
        EnableBreadcrumbs:   true,
        EnableNotifications: true,
        SearchConfig: SearchConfig{
            Placeholder: "Search menu...",
            MinLength:   2,
            Debounce:    300,
        },
    }
    
    // Build with auth context
    authCtx := &AuthContext{
        User: getCurrentUser(),
        Authenticated: true,
    }
    
    builder := NewNavigationSchemaBuilder(navSchema).
        WithConfig(config).
        WithAuthContext(authCtx)
    
    return builder.Build(context.Background())
}
```

## Navigation Types

### Sidebar Navigation

```go
NavigationProps{
    Type:        NavigationSidebar,
    Size:        NavigationSizeMD,
    Collapsible: true,
    Variant:     "default", // "default", "minimal", "card"
    Width:       "256px",   // Custom width
    Position:    "left",    // "left" or "right"
}
```

**Features:**
- Collapsible with smooth animations
- Nested menu support (unlimited depth)
- User section at bottom
- Search integration
- Recent items tracking

### Topbar Navigation  

```go
NavigationProps{
    Type: NavigationTopbar,
    Size: NavigationSizeLG, // Controls height
    
    // Search in topbar
    Search: NavigationSearch{
        Enabled:     true,
        Placeholder: "Search...",
    },
    
    // User menu dropdown
    UserMenu: []NavigationItem{...},
}
```

**Features:**
- Responsive design (mobile hamburger menu)
- Dropdown menus
- Search integration
- Notification bell
- User avatar/menu

### Breadcrumb Navigation

```go
NavigationProps{
    Type: NavigationBreadcrumb,
    
    Breadcrumbs: []BreadcrumbItem{
        {Text: "Home", URL: "/", Clickable: true},
        {Text: "Products", URL: "/products", Clickable: true}, 
        {Text: "Product A", Active: true},
    },
    
    Separator: "/", // Custom separator
}
```

### Tab Navigation

```go
NavigationProps{
    Type:       NavigationTabs,
    TabVariant: "underline", // "default", "pills", "underline"
    
    Items: []NavigationItem{
        {ID: "overview", Text: "Overview", Active: true},
        {ID: "settings", Text: "Settings"},
    },
}
```

### Step Navigation

```go
NavigationProps{
    Type:        NavigationSteps,
    CurrentStep: 2,
    TotalSteps:  4,
    
    Items: []NavigationItem{
        {Text: "Setup", Description: "Configure account"},
        {Text: "Import", Description: "Import data"},
        {Text: "Review", Description: "Review settings"},
        {Text: "Complete", Description: "Finish setup"},
    },
}
```

## Advanced Features

### Search Configuration

```go
Search: NavigationSearch{
    Enabled:     true,
    Placeholder: "Search navigation...",
    MinLength:   2,        // Minimum characters to search
    Debounce:    300,      // Debounce delay in ms
    Results:     []NavigationSearchResult{...},
    Loading:     false,    // Show loading state
    Visible:     false,    // Search overlay visible
}
```

**Search Features:**
- Real-time search with highlighting
- Keyboard shortcuts (/ to open, Esc to close)
- Search across text and descriptions  
- Ranking and relevance scoring
- Recent search suggestions

### Mobile Menu Configuration

```go
MobileMenu: NavigationMobileMenu{
    Enabled:     true,
    Position:    "left",      // "left", "right", "top", "bottom"
    Overlay:     true,        // Show backdrop overlay
    CloseButton: true,        // Show close button
    Swipeable:   true,        // Enable swipe gestures
}
```

### Notification System

```go
Notifications: []Notification{
    {
        ID:      "notif-1",
        Type:    "info",          // "info", "warning", "error", "success"
        Title:   "New Message",
        Message: "You have a new message",
        Icon:    "mail",          // Custom icon
        URL:     "/messages/123", // Click destination
        Read:    false,           // Read status
        CreatedAt: time.Now(),
    },
}
```

### Permission-Based Items

```go
NavigationItem{
    ID:          "admin-panel",
    Text:        "Admin Panel", 
    Permissions: []string{"admin:access"},    // Required permissions
    Roles:       []string{"admin", "super"},  // Required roles
    Condition:   "user.verified === true",    // JavaScript condition
}
```

## HTMX Integration

### Dynamic Loading

```go
NavigationItem{
    ID:       "dynamic-content",
    Text:     "Dynamic Section",
    HXGet:    "/api/navigation/section",  // Load content
    HXTarget: "#nav-content",             // Target container
    HXSwap:   "innerHTML",                // Swap strategy
    HXTrigger: "click",                   // Trigger event
}
```

### Form Actions

```go
NavigationItem{
    ID:     "logout",
    Text:   "Sign Out",
    HXPost: "/auth/logout",       // POST request
    HXSwap: "none",               // No content swap
    HXTrigger: "click",
}
```

## Alpine.js Integration

### Custom Click Handlers

```go
NavigationItem{
    ID:          "theme-toggle",
    Text:        "Toggle Theme",
    AlpineClick: "toggleTheme()",     // Custom Alpine method
}
```

### Conditional Display

```go
NavigationItem{
    ID:   "premium-features",
    Text: "Premium Features", 
    AlpineConfig: AlpineConfig{
        Show: "user.isPremium",       // Show condition
    },
}
```

### State Management

```javascript
// Available Alpine.js state and methods
{
    // Navigation state
    sidebarCollapsed: false,
    mobileMenuOpen: false, 
    searchOpen: false,
    searchQuery: '',
    searchResults: [],
    
    // User preferences  
    preferences: {
        theme: 'auto',
        recentMenuItems: [],
        favoriteMenuItems: [],
    },
    
    // Methods
    toggleSidebar(),
    toggleMobileMenu(),
    toggleSearch(),
    navigateToItem(item),
    performSearch(),
    savePreferences(),
}
```

## Styling & Theming

### Theme Integration

```go
NavigationProps{
    ThemeID:  "dark-theme",              // Theme identifier
    DarkMode: true,                      // Dark mode override
    TokenOverrides: map[string]string{   // Custom token overrides
        "nav-bg":     "#1a1a1a",
        "nav-border": "#333333",
    },
}
```

### CSS Classes

The navigation system uses semantic CSS classes for easy customization:

```css
/* Navigation container */
.navigation { }
.navigation-sidebar { }
.navigation-topbar { }

/* Navigation variants */
.navigation-sidebar--minimal { }
.navigation-sidebar--card { }
.navigation-tabs--pills { }

/* Navigation items */
.nav-menu-item { }
.nav-menu-item--active { }
.nav-menu-item--disabled { }
.nav-menu-item--depth-1 { }

/* States */
.navigation--searchable { }
.navigation--mobile { }
.navigation--dark { }
```

## Business Logic Integration

### Navigation Service

```go
// Create navigation service with auth integration
service := NewNavigationService()
service.AuthProvider = myAuthProvider
service.PermissionProvider = myPermissionProvider

// Build navigation context
navCtx, err := service.BuildNavigationContext(ctx, "/current/path")

// Process items with permissions
processedItems, err := service.ProcessNavigationItems(ctx, items, navCtx)

// Search functionality
results, err := service.SearchMenuItems(ctx, "query", items, navCtx)
```

### Auth Provider Interface

```go
type AuthProvider interface {
    GetCurrentUser(ctx context.Context) (*User, error)
    IsAuthenticated(ctx context.Context) bool
    HasRole(ctx context.Context, role string) bool
    HasPermission(ctx context.Context, permission string) bool
}

type PermissionProvider interface {
    CanAccessMenuItem(ctx context.Context, user *User, item *NavigationItem) bool
    GetUserPermissions(ctx context.Context, user *User) ([]string, error)
    GetVisibleMenuItems(ctx context.Context, user *User, items []NavigationItem) ([]NavigationItem, error)
}
```

## Performance Optimization

### Lazy Loading

```go
NavigationProps{
    LazyLoad: true,               // Enable lazy loading
    HXGet:    "/api/nav/lazy",    // Lazy load endpoint
    HXTarget: "#nav-content",     // Target container
}
```

### Caching

```go
NavigationProps{
    CacheData: true,              // Enable data caching
}

// Service configuration
service := &NavigationService{
    EnableCaching:   true,
    CacheTimeout:    5 * time.Minute,
    MaxMenuItems:    1000,
    SearchDebounce:  300 * time.Millisecond,
}
```

### Virtual Scrolling

For large navigation trees:

```go
NavigationProps{
    VirtualScrolling: true,       // Enable for large lists
    MaxVisibleItems:  50,         // Limit visible items
}
```

## Accessibility Features

### ARIA Support

```go
NavigationProps{
    AriaLabel: "Main navigation",
    AriaLabels: map[string]string{
        "search":     "Search navigation",
        "userMenu":   "User account menu", 
        "mobileMenu": "Mobile navigation menu",
    },
}
```

### Keyboard Navigation

- **Tab/Shift+Tab**: Navigate between items
- **Enter/Space**: Activate items
- **Arrow Keys**: Navigate hierarchically
- **/** : Open search
- **Esc**: Close overlays
- **Ctrl+B**: Toggle sidebar

### Screen Reader Support

```go
NavigationItem{
    Text:        "Dashboard",
    Description: "Main application dashboard", // Screen reader description
}
```

## Error Handling

```go
// Service error handling
navCtx, err := service.BuildNavigationContext(ctx, path)
if err != nil {
    // Handle navigation context error
    log.Error("Failed to build navigation context", err)
    return fallbackNavigation()
}

// Permission error handling  
items, err := service.ProcessNavigationItems(ctx, items, navCtx)
if err != nil {
    // Handle permission processing error
    log.Warn("Permission processing failed", err) 
    // Continue with unfiltered items
}
```

## Security Considerations

### CSRF Protection

```go
NavigationProps{
    AuthContext: &AuthContext{
        CSRF: "csrf-token-here",      // CSRF token
    },
}
```

### Permission Validation

```go
// Server-side permission validation
func (p *MyPermissionProvider) CanAccessMenuItem(ctx context.Context, user *User, item *NavigationItem) bool {
    // Validate permissions server-side
    for _, required := range item.Permissions {
        if !p.userHasPermission(user, required) {
            return false
        }
    }
    return true
}
```

## Testing

### Unit Testing Navigation Logic

```go
func TestNavigationService_ProcessItems(t *testing.T) {
    service := NewNavigationService()
    
    // Mock auth providers
    service.AuthProvider = &MockAuthProvider{
        User: &User{Roles: []string{"user"}},
    }
    
    items := []NavigationItem{
        {ID: "public", Text: "Public"},
        {ID: "admin", Text: "Admin", Roles: []string{"admin"}},
    }
    
    navCtx := &NavigationContext{User: &User{Roles: []string{"user"}}}
    
    processed, err := service.ProcessNavigationItems(context.Background(), items, navCtx)
    assert.NoError(t, err)
    assert.Len(t, processed, 1) // Only public item visible
    assert.Equal(t, "public", processed[0].ID)
}
```

### Integration Testing

```go
func TestNavigationRendering(t *testing.T) {
    props := ExampleSidebarNavigation()
    
    // Render navigation
    html := renderComponent(Navigation(props))
    
    // Verify output
    assert.Contains(t, html, "navigation-sidebar")
    assert.Contains(t, html, "Dashboard") 
    assert.Contains(t, html, "John Doe")
}
```

## Migration Guide

### From Basic Navigation

If upgrading from a basic navigation implementation:

1. **Update imports**: Add the new navigation logic and schema files
2. **Enhance props**: Add auth context and enhanced configuration
3. **Add permissions**: Define role-based access for menu items
4. **Configure search**: Enable search functionality if needed
5. **Update styling**: Apply new CSS classes for enhanced features

### Breaking Changes

- `NavigationProps` structure has been extended with new fields
- Auth integration requires implementing `AuthProvider` and `PermissionProvider`
- CSS classes have been updated for better semantic naming
- Alpine.js data structure has changed for enhanced features

## Examples

See `navigation_examples.go` for comprehensive examples including:

- **ExampleSidebarNavigation()**: Complete sidebar with auth
- **ExampleTopbarNavigation()**: Responsive topbar
- **ExampleBreadcrumbNavigation()**: Breadcrumb implementation
- **ExampleTabsNavigation()**: Tab navigation
- **ExampleStepsNavigation()**: Step-based navigation
- **ExampleMinimalSidebar()**: Minimal design
- **ExampleCardSidebar()**: Card-style sidebar
- **ExamplePublicNavigation()**: Public site navigation  
- **ExampleMobileResponsiveNavigation()**: Mobile-optimized

Each example demonstrates different aspects of the navigation system and can be used as starting points for your implementation.

## Contributing

When contributing to the Navigation organism:

1. Follow the established architectural patterns
2. Ensure accessibility compliance
3. Add comprehensive tests for new features  
4. Update documentation for API changes
5. Consider performance implications
6. Test with different auth scenarios

## License

Part of the RUUN component system. See project license for details.