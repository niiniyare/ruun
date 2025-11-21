package theme

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// ThemeSwitcher provides runtime theme switching with the new Manager architecture
type ThemeSwitcher struct {
	manager       *Manager
	tenantManager *TenantManager
	storage       UserPreferenceStorage
	observers     []ThemeSwitchObserver
	config        *SwitcherConfig
	mu            sync.RWMutex
}

// SwitcherConfig configures theme switching behavior
type SwitcherConfig struct {
	EnablePersistence  bool          `json:"enablePersistence"`
	EnableAutoDetect   bool          `json:"enableAutoDetect"`
	EnableTransitions  bool          `json:"enableTransitions"`
	TransitionDuration time.Duration `json:"transitionDuration"`
	PreloadThemes      bool          `json:"preloadThemes"`
	MaxHistory         int           `json:"maxHistory"`
	StorageKey         string        `json:"storageKey"`
	CookieExpiry       time.Duration `json:"cookieExpiry"`
}

// UserPreferences stores user theme preferences
type UserPreferences struct {
	CurrentTheme string            `json:"currentTheme"`
	DarkMode     bool              `json:"darkMode"`
	AutoDarkMode bool              `json:"autoDarkMode"`
	CustomTokens map[string]string `json:"customTokens,omitempty"`
	History      []string          `json:"history,omitempty"`
	LastSwitched time.Time         `json:"lastSwitched"`
	TenantID     string            `json:"tenantId,omitempty"`
}

// UserPreferenceStorage interface for persisting user preferences
type UserPreferenceStorage interface {
	SavePreferences(ctx context.Context, userID string, prefs *UserPreferences) error
	LoadPreferences(ctx context.Context, userID string) (*UserPreferences, error)
	DeletePreferences(ctx context.Context, userID string) error
}

// ThemeSwitchObserver interface for theme switch notifications
type ThemeSwitchObserver interface {
	OnThemeSwitched(ctx context.Context, event *ThemeSwitchEvent)
	OnDarkModeToggled(ctx context.Context, event *DarkModeEvent)
}

// ThemeSwitchEvent represents a theme switch event
type ThemeSwitchEvent struct {
	UserID        string         `json:"userId"`
	TenantID      string         `json:"tenantId,omitempty"`
	PreviousTheme string         `json:"previousTheme"`
	NewTheme      string         `json:"newTheme"`
	Timestamp     time.Time      `json:"timestamp"`
	Metadata      map[string]any `json:"metadata,omitempty"`
}

// DarkModeEvent represents a dark mode toggle event
type DarkModeEvent struct {
	UserID       string    `json:"userId"`
	TenantID     string    `json:"tenantId,omitempty"`
	DarkMode     bool      `json:"darkMode"`
	PreviousMode bool      `json:"previousMode"`
	Timestamp    time.Time `json:"timestamp"`
}

// ThemeSwitchResult represents the result of a theme switch operation
type ThemeSwitchResult struct {
	Success       bool              `json:"success"`
	ThemeID       string            `json:"themeId"`
	CSS           string            `json:"css"`
	Variables     map[string]string `json:"variables,omitempty"`
	TransitionCSS string            `json:"transitionCss,omitempty"`
	Error         string            `json:"error,omitempty"`
	Metadata      map[string]any    `json:"metadata,omitempty"`
}

// AlpineJSData provides data structure for Alpine.js integration
type AlpineJSData struct {
	CurrentTheme    string            `json:"currentTheme"`
	AvailableThemes []*ThemeInfo      `json:"availableThemes"`
	DarkMode        bool              `json:"darkMode"`
	IsLoading       bool              `json:"isLoading"`
	Error           string            `json:"error,omitempty"`
	Preferences     *UserPreferences  `json:"preferences"`
	Config          map[string]any    `json:"config"`
	Methods         map[string]string `json:"methods"`
}

// ThemeInfo provides basic theme information for UI
type ThemeInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
}

// NewThemeSwitcher creates a new theme switcher with new package architecture
func NewThemeSwitcher(manager *Manager, tenantManager *TenantManager, storage UserPreferenceStorage, config *SwitcherConfig) *ThemeSwitcher {
	if config == nil {
		config = DefaultSwitcherConfig()
	}

	return &ThemeSwitcher{
		manager:       manager,
		tenantManager: tenantManager,
		storage:       storage,
		config:        config,
		observers:     make([]ThemeSwitchObserver, 0),
	}
}

// DefaultSwitcherConfig returns default switcher configuration
func DefaultSwitcherConfig() *SwitcherConfig {
	return &SwitcherConfig{
		EnablePersistence:  true,
		EnableAutoDetect:   true,
		EnableTransitions:  true,
		TransitionDuration: 300 * time.Millisecond,
		PreloadThemes:      true,
		MaxHistory:         10,
		StorageKey:         "theme_preferences",
		CookieExpiry:       30 * 24 * time.Hour, // 30 days
	}
}

// DefaultUserPreferences returns default user preferences
func DefaultUserPreferences() *UserPreferences {
	return &UserPreferences{
		CurrentTheme: "default",
		DarkMode:     false,
		AutoDarkMode: true,
		History:      make([]string, 0),
		LastSwitched: time.Now(),
	}
}

// SwitchTheme switches to a new theme using the new Manager architecture
func (ts *ThemeSwitcher) SwitchTheme(ctx context.Context, themeID, userID, tenantID string) (*ThemeSwitchResult, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Load user preferences
	prefs, err := ts.loadPreferences(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load preferences: %w", err)
	}

	previousTheme := prefs.CurrentTheme

	// Get theme using new Manager
	var compiled *CompiledTheme
	if tenantID != "" && ts.tenantManager != nil {
		// Multi-tenant mode
		compiled, err = ts.tenantManager.GetTenantTheme(ctx, tenantID, nil)
	} else {
		// Single-tenant mode
		compiled, err = ts.manager.GetTheme(ctx, themeID, nil)
	}

	if err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Update preferences
	prefs.CurrentTheme = themeID
	prefs.TenantID = tenantID
	prefs.LastSwitched = time.Now()
	ts.addToHistory(prefs, themeID)

	// Save preferences
	if ts.config.EnablePersistence && ts.storage != nil {
		if err := ts.storage.SavePreferences(ctx, userID, prefs); err != nil {
			// Log but don't fail
		}
	}

	// Create event
	event := &ThemeSwitchEvent{
		UserID:        userID,
		TenantID:      tenantID,
		PreviousTheme: previousTheme,
		NewTheme:      themeID,
		Timestamp:     time.Now(),
	}

	// Notify observers
	ts.notifyObservers(ctx, event)

	// Generate transition CSS if enabled
	var transitionCSS string
	if ts.config.EnableTransitions {
		transitionCSS = ts.generateTransitionCSS()
	}

	return &ThemeSwitchResult{
		Success:       true,
		ThemeID:       themeID,
		CSS:           compiled.CSS,
		Variables:     ts.extractVariables(compiled),
		TransitionCSS: transitionCSS,
		Metadata: map[string]any{
			"previousTheme": previousTheme,
			"timestamp":     event.Timestamp,
		},
	}, nil
}

// ToggleDarkMode toggles dark mode using conditional theming
func (ts *ThemeSwitcher) ToggleDarkMode(ctx context.Context, userID, tenantID string) (*ThemeSwitchResult, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Load preferences
	prefs, err := ts.loadPreferences(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to load preferences: %w", err)
	}

	previousMode := prefs.DarkMode
	newMode := !previousMode

	// Update preferences
	prefs.DarkMode = newMode
	prefs.LastSwitched = time.Now()

	// Save preferences
	if ts.config.EnablePersistence && ts.storage != nil {
		if err := ts.storage.SavePreferences(ctx, userID, prefs); err != nil {
			// Log but don't fail
		}
	}

	// Get theme with dark mode via conditional evaluation
	evalData := map[string]any{
		"user": map[string]any{
			"id": userID,
			"preferences": map[string]any{
				"darkMode": newMode,
			},
		},
	}

	var compiled *CompiledTheme
	if tenantID != "" && ts.tenantManager != nil {
		compiled, err = ts.tenantManager.GetTenantTheme(ctx, tenantID, evalData)
	} else {
		compiled, err = ts.manager.GetTheme(ctx, prefs.CurrentTheme, evalData)
	}

	if err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Create event
	event := &DarkModeEvent{
		UserID:       userID,
		TenantID:     tenantID,
		DarkMode:     newMode,
		PreviousMode: previousMode,
		Timestamp:    time.Now(),
	}

	// Notify observers
	ts.notifyDarkModeObservers(ctx, event)

	return &ThemeSwitchResult{
		Success:   true,
		ThemeID:   prefs.CurrentTheme,
		CSS:       compiled.CSS,
		Variables: ts.extractVariables(compiled),
		Metadata: map[string]any{
			"darkModeChanged": true,
			"previousMode":    previousMode,
			"newMode":         newMode,
		},
	}, nil
}

// LoadUserPreferences loads preferences for a user
func (ts *ThemeSwitcher) LoadUserPreferences(ctx context.Context, userID string) (*UserPreferences, error) {
	return ts.loadPreferences(ctx, userID)
}

// GetAlpineJSData returns data structure for Alpine.js integration
func (ts *ThemeSwitcher) GetAlpineJSData(ctx context.Context) (*AlpineJSData, error) {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	// Get available themes
	var themes []*Theme
	var err error

	tenantID := getTenantID(ctx)
	if tenantID != "" && ts.tenantManager != nil {
		themes, err = ts.tenantManager.ListTenantThemes(ctx, tenantID)
		if err != nil {
			return nil, err
		}
	} else {
		themes, err = ts.manager.ListThemes(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Convert to ThemeInfo
	themeInfos := make([]*ThemeInfo, len(themes))
	for i, theme := range themes {
		themeInfos[i] = &ThemeInfo{
			ID:          theme.ID,
			Name:        theme.Name,
			Description: theme.Description,
			Version:     theme.Version,
		}
	}

	return &AlpineJSData{
		CurrentTheme:    "default", // Would load from user prefs
		AvailableThemes: themeInfos,
		DarkMode:        false, // Would load from user prefs
		IsLoading:       false,
		Preferences:     DefaultUserPreferences(),
		Config: map[string]any{
			"enableTransitions":  ts.config.EnableTransitions,
			"transitionDuration": ts.config.TransitionDuration.Milliseconds(),
			"enableAutoDetect":   ts.config.EnableAutoDetect,
		},
		Methods: map[string]string{
			"switchTheme":    "switchTheme(themeId)",
			"toggleDarkMode": "toggleDarkMode()",
			"resetTheme":     "resetTheme()",
			"exportTheme":    "exportTheme()",
		},
	}, nil
}

// GenerateAlpineJSComponent generates Alpine.js component for theme switching
func (ts *ThemeSwitcher) GenerateAlpineJSComponent(ctx context.Context, tenantID string) (string, error) {
	ctx = WithTenant(ctx, tenantID)
	data, err := ts.GetAlpineJSData(ctx)
	if err != nil {
		return "", err
	}

	dataJSON, _ := json.Marshal(data)

	return fmt.Sprintf(`
window.themeData = %s;

function createThemeSwitcher() {
    return {
        ...window.themeData,
        
        init() {
            this.detectSystemTheme();
            this.loadSavedPreferences();
            this.applyTheme();
            
            // Listen for system dark mode changes
            if (this.config.enableAutoDetect) {
                window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
                    if (this.preferences.autoDarkMode) {
                        this.setDarkMode(e.matches);
                    }
                });
            }
        },
        
        async switchTheme(themeId) {
            this.isLoading = true;
            this.error = '';
            
            try {
                const response = await fetch('/api/themes/switch', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        themeId: themeId,
                        userId: this.getCurrentUserId(),
                        tenantId: this.getCurrentTenantId()
                    })
                });
                
                if (!response.ok) {
                    throw new Error('Failed to switch theme');
                }
                
                const result = await response.json();
                
                if (result.success) {
                    this.currentTheme = result.themeId;
                    this.applyThemeCSS(result.css, result.transitionCss);
                    this.savePreferences();
                    this.notifyThemeChanged(result);
                } else {
                    throw new Error(result.error || 'Unknown error');
                }
            } catch (error) {
                this.error = error.message;
                console.error('Theme switch failed:', error);
            } finally {
                this.isLoading = false;
            }
        },
        
        async toggleDarkMode() {
            this.isLoading = true;
            this.error = '';
            
            try {
                const response = await fetch('/api/themes/toggle-dark', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        userId: this.getCurrentUserId(),
                        tenantId: this.getCurrentTenantId()
                    })
                });
                
                if (!response.ok) {
                    throw new Error('Failed to toggle dark mode');
                }
                
                const result = await response.json();
                
                if (result.success) {
                    this.darkMode = !this.darkMode;
                    this.applyThemeCSS(result.css, result.transitionCss);
                    this.savePreferences();
                    this.notifyThemeChanged(result);
                } else {
                    throw new Error(result.error || 'Unknown error');
                }
            } catch (error) {
                this.error = error.message;
                console.error('Dark mode toggle failed:', error);
            } finally {
                this.isLoading = false;
            }
        },
        
        resetTheme() {
            this.switchTheme('default');
            this.setDarkMode(false);
            this.preferences.autoDarkMode = true;
            this.savePreferences();
        },
        
        exportTheme() {
            const themeData = {
                currentTheme: this.currentTheme,
                darkMode: this.darkMode,
                preferences: this.preferences
            };
            
            const dataStr = JSON.stringify(themeData, null, 2);
            const dataUri = 'data:application/json;charset=utf-8,'+ encodeURIComponent(dataStr);
            
            const exportFileDefaultName = 'theme-preferences.json';
            
            const linkElement = document.createElement('a');
            linkElement.setAttribute('href', dataUri);
            linkElement.setAttribute('download', exportFileDefaultName);
            linkElement.click();
        },
        
        detectSystemTheme() {
            if (this.config.enableAutoDetect && this.preferences.autoDarkMode) {
                const systemDarkMode = window.matchMedia('(prefers-color-scheme: dark)').matches;
                this.setDarkMode(systemDarkMode);
            }
        },
        
        setDarkMode(enabled) {
            this.darkMode = enabled;
            document.documentElement.classList.toggle('dark', enabled);
        },
        
        applyTheme() {
            document.documentElement.setAttribute('data-theme', this.currentTheme);
            this.setDarkMode(this.darkMode);
        },
        
        applyThemeCSS(css, transitionCss) {
            // Apply transition CSS if provided
            if (transitionCss && this.config.enableTransitions) {
                const transitionStyle = document.createElement('style');
                transitionStyle.textContent = transitionCss;
                document.head.appendChild(transitionStyle);
                
                // Remove transition after duration
                setTimeout(() => {
                    document.head.removeChild(transitionStyle);
                }, this.config.transitionDuration);
            }
            
            // Apply new theme CSS
            let themeStyle = document.getElementById('theme-styles');
            if (!themeStyle) {
                themeStyle = document.createElement('style');
                themeStyle.id = 'theme-styles';
                document.head.appendChild(themeStyle);
            }
            themeStyle.textContent = css;
            
            this.applyTheme();
        },
        
        loadSavedPreferences() {
            try {
                const saved = localStorage.getItem('%s');
                if (saved) {
                    const prefs = JSON.parse(saved);
                    this.preferences = { ...this.preferences, ...prefs };
                    this.currentTheme = prefs.currentTheme || this.currentTheme;
                    this.darkMode = prefs.darkMode || this.darkMode;
                }
            } catch (error) {
                console.warn('Failed to load saved theme preferences:', error);
            }
        },
        
        savePreferences() {
            try {
                const prefs = {
                    currentTheme: this.currentTheme,
                    darkMode: this.darkMode,
                    autoDarkMode: this.preferences.autoDarkMode,
                    lastSwitched: new Date().toISOString()
                };
                localStorage.setItem('%s', JSON.stringify(prefs));
            } catch (error) {
                console.warn('Failed to save theme preferences:', error);
            }
        },
        
        getCurrentUserId() {
            // This would be implemented based on your auth system
            return window.currentUser?.id || 'anonymous';
        },
        
        getCurrentTenantId() {
            // This would be implemented based on your multi-tenant setup
            return document.body.dataset.tenantId || '';
        },
        
        notifyThemeChanged(result) {
            // Dispatch custom event for other components to listen
            const event = new CustomEvent('themeChanged', {
                detail: {
                    theme: this.currentTheme,
                    darkMode: this.darkMode,
                    result: result
                }
            });
            window.dispatchEvent(event);
        },
        
        // Utility methods for templates
        isCurrentTheme(themeId) {
            return this.currentTheme === themeId;
        },
        
        getThemeDisplayName(themeId) {
            const theme = this.availableThemes.find(t => t.id === themeId);
            return theme ? theme.name : themeId;
        },
        
        getThemeDescription(themeId) {
            const theme = this.availableThemes.find(t => t.id === themeId);
            return theme ? theme.description : '';
        }
    }
}

// Initialize Alpine.js component
document.addEventListener('alpine:init', () => {
    Alpine.data('themeSwitcher', createThemeSwitcher);
});
`, dataJSON, ts.config.StorageKey, ts.config.StorageKey), nil
}

// AddObserver adds a theme switch observer
func (ts *ThemeSwitcher) AddObserver(observer ThemeSwitchObserver) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.observers = append(ts.observers, observer)
}

// RemoveObserver removes a theme switch observer
func (ts *ThemeSwitcher) RemoveObserver(observer ThemeSwitchObserver) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i, obs := range ts.observers {
		if obs == observer {
			ts.observers = append(ts.observers[:i], ts.observers[i+1:]...)
			break
		}
	}
}

// Private helper methods

func (ts *ThemeSwitcher) loadPreferences(ctx context.Context, userID string) (*UserPreferences, error) {
	if !ts.config.EnablePersistence || ts.storage == nil {
		return DefaultUserPreferences(), nil
	}

	prefs, err := ts.storage.LoadPreferences(ctx, userID)
	if err != nil {
		return DefaultUserPreferences(), nil // Return defaults on error
	}

	return prefs, nil
}

func (ts *ThemeSwitcher) addToHistory(prefs *UserPreferences, themeID string) {
	// Remove if already exists
	for i, theme := range prefs.History {
		if theme == themeID {
			prefs.History = append(prefs.History[:i], prefs.History[i+1:]...)
			break
		}
	}

	// Add to front
	prefs.History = append([]string{themeID}, prefs.History...)

	// Limit history size
	if len(prefs.History) > ts.config.MaxHistory {
		prefs.History = prefs.History[:ts.config.MaxHistory]
	}
}

func (ts *ThemeSwitcher) notifyObservers(ctx context.Context, event *ThemeSwitchEvent) {
	for _, observer := range ts.observers {
		go observer.OnThemeSwitched(ctx, event)
	}
}

func (ts *ThemeSwitcher) notifyDarkModeObservers(ctx context.Context, event *DarkModeEvent) {
	for _, observer := range ts.observers {
		go observer.OnDarkModeToggled(ctx, event)
	}
}

func (ts *ThemeSwitcher) generateTransitionCSS() string {
	duration := ts.config.TransitionDuration.Milliseconds()
	return fmt.Sprintf(`
* {
    transition: color %dms ease, background-color %dms ease, border-color %dms ease, box-shadow %dms ease;
}
`, duration, duration, duration, duration)
}

func (ts *ThemeSwitcher) extractVariables(compiled *CompiledTheme) map[string]string {
	// Extract CSS variables from compiled theme
	variables := make(map[string]string)

	// The new package already has resolved tokens
	// You could iterate through compiled.Theme.Tokens if needed

	return variables
}

// SimpleUserPreferenceStorage provides a simple in-memory storage implementation
type SimpleUserPreferenceStorage struct {
	preferences map[string]*UserPreferences
	mu          sync.RWMutex
}

// NewSimpleUserPreferenceStorage creates a new simple storage
func NewSimpleUserPreferenceStorage() *SimpleUserPreferenceStorage {
	return &SimpleUserPreferenceStorage{
		preferences: make(map[string]*UserPreferences),
	}
}

// SavePreferences saves user preferences
func (s *SimpleUserPreferenceStorage) SavePreferences(ctx context.Context, userID string, prefs *UserPreferences) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.preferences[userID] = prefs
	return nil
}

// LoadPreferences loads user preferences
func (s *SimpleUserPreferenceStorage) LoadPreferences(ctx context.Context, userID string) (*UserPreferences, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if prefs, exists := s.preferences[userID]; exists {
		return prefs, nil
	}

	return DefaultUserPreferences(), nil
}

// DeletePreferences deletes user preferences
func (s *SimpleUserPreferenceStorage) DeletePreferences(ctx context.Context, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.preferences, userID)
	return nil
}
