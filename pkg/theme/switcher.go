package theme

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
)

// ThemeSwitcher provides runtime theme switching capabilities with persistence
type ThemeSwitcher struct {
	runtime       *Runtime
	storage       ThemeStorage
	preferences   *UserPreferences
	observers     []ThemeObserver
	config        *SwitcherConfig
	mu            sync.RWMutex
	lastSwitched  time.Time
}

// SwitcherConfig configures theme switching behavior
type SwitcherConfig struct {
	EnablePersistence   bool          `json:"enablePersistence"`
	EnableAutoDetect    bool          `json:"enableAutoDetect"`
	EnableTransitions   bool          `json:"enableTransitions"`
	TransitionDuration  time.Duration `json:"transitionDuration"`
	PreloadThemes       bool          `json:"preloadThemes"`
	MaxHistory          int           `json:"maxHistory"`
	StorageKey          string        `json:"storageKey"`
	CookieExpiry        time.Duration `json:"cookieExpiry"`
}

// UserPreferences stores user theme preferences
type UserPreferences struct {
	CurrentTheme   string    `json:"currentTheme"`
	DarkMode       bool      `json:"darkMode"`
	AutoDarkMode   bool      `json:"autoDarkMode"`
	CustomTokens   map[string]string `json:"customTokens,omitempty"`
	History        []string  `json:"history,omitempty"`
	LastSwitched   time.Time `json:"lastSwitched"`
	TenantID       string    `json:"tenantId,omitempty"`
}

// ThemeStorage interface for persisting theme preferences
type ThemeStorage interface {
	SavePreferences(userID string, prefs *UserPreferences) error
	LoadPreferences(userID string) (*UserPreferences, error)
	DeletePreferences(userID string) error
}

// ThemeObserver interface for theme change notifications
type ThemeObserver interface {
	OnThemeChanged(event *ThemeChangeEvent)
}

// ThemeChangeEvent represents a theme change event
type ThemeChangeEvent struct {
	PreviousTheme string                 `json:"previousTheme"`
	NewTheme      string                 `json:"newTheme"`
	DarkMode      bool                   `json:"darkMode"`
	UserID        string                 `json:"userId,omitempty"`
	TenantID      string                 `json:"tenantId,omitempty"`
	Timestamp     time.Time              `json:"timestamp"`
	Metadata      map[string]any `json:"metadata,omitempty"`
}

// ThemeSwitchResult represents the result of a theme switch operation
type ThemeSwitchResult struct {
	Success        bool              `json:"success"`
	ThemeID        string            `json:"themeId"`
	CSS            string            `json:"css"`
	Variables      map[string]string `json:"variables"`
	Classes        []string          `json:"classes"`
	Error          string            `json:"error,omitempty"`
	TransitionCSS  string            `json:"transitionCss,omitempty"`
	Metadata       map[string]any `json:"metadata,omitempty"`
}

// AlpineJSData provides data structure for Alpine.js integration
type AlpineJSData struct {
	CurrentTheme    string                 `json:"currentTheme"`
	AvailableThemes []ThemeInfo            `json:"availableThemes"`
	DarkMode        bool                   `json:"darkMode"`
	IsLoading       bool                   `json:"isLoading"`
	Error           string                 `json:"error,omitempty"`
	Preferences     *UserPreferences       `json:"preferences"`
	Config          map[string]any `json:"config"`
	Methods         map[string]string      `json:"methods"`
}

// NewThemeSwitcher creates a new theme switcher
func NewThemeSwitcher(runtime *Runtime, storage ThemeStorage, config *SwitcherConfig) *ThemeSwitcher {
	if config == nil {
		config = DefaultSwitcherConfig()
	}

	return &ThemeSwitcher{
		runtime:     runtime,
		storage:     storage,
		config:      config,
		observers:   make([]ThemeObserver, 0),
		preferences: DefaultUserPreferences(),
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

// SwitchTheme switches to a new theme
func (ts *ThemeSwitcher) SwitchTheme(themeID string, userID string) (*ThemeSwitchResult, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	previousTheme := ts.preferences.CurrentTheme

	// Load user preferences
	if err := ts.loadUserPreferences(userID); err != nil {
		return nil, fmt.Errorf("failed to load user preferences: %w", err)
	}

	// Switch theme in runtime
	if err := ts.runtime.SetTheme(themeID); err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Get compiled theme data
	css, err := ts.runtime.GetThemeCSS()
	if err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Update preferences
	ts.preferences.CurrentTheme = themeID
	ts.preferences.LastSwitched = time.Now()
	ts.addToHistory(themeID)

	// Save preferences if persistence is enabled
	if ts.config.EnablePersistence && ts.storage != nil {
		if err := ts.storage.SavePreferences(userID, ts.preferences); err != nil {
			// Don't fail the operation, just log the warning
		}
	}

	// Create theme change event
	event := &ThemeChangeEvent{
		PreviousTheme: previousTheme,
		NewTheme:      themeID,
		DarkMode:      ts.preferences.DarkMode,
		UserID:        userID,
		TenantID:      ts.preferences.TenantID,
		Timestamp:     time.Now(),
	}

	// Notify observers
	ts.notifyObservers(event)

	// Generate transition CSS if enabled
	var transitionCSS string
	if ts.config.EnableTransitions {
		transitionCSS = ts.generateTransitionCSS()
	}

	result := &ThemeSwitchResult{
		Success:       true,
		ThemeID:       themeID,
		CSS:           css,
		Variables:     ts.extractVariables(css),
		Classes:       ts.extractClasses(css),
		TransitionCSS: transitionCSS,
		Metadata: map[string]any{
			"previousTheme": previousTheme,
			"timestamp":     event.Timestamp,
		},
	}

	ts.lastSwitched = time.Now()
	return result, nil
}

// ToggleDarkMode toggles dark mode
func (ts *ThemeSwitcher) ToggleDarkMode(userID string) (*ThemeSwitchResult, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	// Load user preferences
	if err := ts.loadUserPreferences(userID); err != nil {
		return nil, fmt.Errorf("failed to load user preferences: %w", err)
	}

	previousMode := ts.preferences.DarkMode
	newMode := !previousMode

	// Set dark mode in runtime
	if err := ts.runtime.SetDarkMode(newMode); err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Update preferences
	ts.preferences.DarkMode = newMode
	ts.preferences.LastSwitched = time.Now()

	// Save preferences
	if ts.config.EnablePersistence && ts.storage != nil {
		if err := ts.storage.SavePreferences(userID, ts.preferences); err != nil {
			// Don't fail the operation
		}
	}

	// Get updated CSS
	css, err := ts.runtime.GetThemeCSS()
	if err != nil {
		return &ThemeSwitchResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Create event
	event := &ThemeChangeEvent{
		PreviousTheme: ts.preferences.CurrentTheme,
		NewTheme:      ts.preferences.CurrentTheme,
		DarkMode:      newMode,
		UserID:        userID,
		Timestamp:     time.Now(),
		Metadata: map[string]any{
			"darkModeToggle": true,
			"previousMode":   previousMode,
		},
	}

	// Notify observers
	ts.notifyObservers(event)

	return &ThemeSwitchResult{
		Success:   true,
		ThemeID:   ts.preferences.CurrentTheme,
		CSS:       css,
		Variables: ts.extractVariables(css),
		Classes:   ts.extractClasses(css),
		Metadata: map[string]any{
			"darkModeChanged": true,
			"previousMode":    previousMode,
			"newMode":         newMode,
		},
	}, nil
}

// LoadUserPreferences loads preferences for a user
func (ts *ThemeSwitcher) LoadUserPreferences(userID string) (*UserPreferences, error) {
	if !ts.config.EnablePersistence || ts.storage == nil {
		return ts.preferences, nil
	}

	prefs, err := ts.storage.LoadPreferences(userID)
	if err != nil {
		return DefaultUserPreferences(), nil // Return defaults on error
	}

	ts.mu.Lock()
	ts.preferences = prefs
	ts.mu.Unlock()

	// Apply loaded preferences to runtime
	if err := ts.runtime.SetTheme(prefs.CurrentTheme); err != nil {
		// Fall back to default theme
		ts.runtime.SetTheme("default")
		prefs.CurrentTheme = "default"
	}

	if err := ts.runtime.SetDarkMode(prefs.DarkMode); err != nil {
		// Ignore dark mode errors
	}

	return prefs, nil
}

// GetAlpineJSData returns data structure for Alpine.js integration
func (ts *ThemeSwitcher) GetAlpineJSData() *AlpineJSData {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	availableThemes := ts.runtime.GetAvailableThemes()

	return &AlpineJSData{
		CurrentTheme:    ts.preferences.CurrentTheme,
		AvailableThemes: availableThemes,
		DarkMode:        ts.preferences.DarkMode,
		IsLoading:       false,
		Preferences:     ts.preferences,
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
	}
}

// GenerateAlpineJSComponent generates Alpine.js component for theme switching
func (ts *ThemeSwitcher) GenerateAlpineJSComponent() string {
	data := ts.GetAlpineJSData()
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
                        userId: this.getCurrentUserId()
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
                        userId: this.getCurrentUserId()
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
`, dataJSON, ts.config.StorageKey, ts.config.StorageKey)
}

// AddObserver adds a theme change observer
func (ts *ThemeSwitcher) AddObserver(observer ThemeObserver) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.observers = append(ts.observers, observer)
}

// RemoveObserver removes a theme change observer
func (ts *ThemeSwitcher) RemoveObserver(observer ThemeObserver) {
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

func (ts *ThemeSwitcher) loadUserPreferences(userID string) error {
	if !ts.config.EnablePersistence || ts.storage == nil {
		return nil
	}

	prefs, err := ts.storage.LoadPreferences(userID)
	if err != nil {
		return err
	}

	ts.preferences = prefs
	return nil
}

func (ts *ThemeSwitcher) addToHistory(themeID string) {
	// Remove if already exists
	for i, theme := range ts.preferences.History {
		if theme == themeID {
			ts.preferences.History = append(ts.preferences.History[:i], ts.preferences.History[i+1:]...)
			break
		}
	}

	// Add to front
	ts.preferences.History = append([]string{themeID}, ts.preferences.History...)

	// Limit history size
	if len(ts.preferences.History) > ts.config.MaxHistory {
		ts.preferences.History = ts.preferences.History[:ts.config.MaxHistory]
	}
}

func (ts *ThemeSwitcher) notifyObservers(event *ThemeChangeEvent) {
	for _, observer := range ts.observers {
		go observer.OnThemeChanged(event)
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

func (ts *ThemeSwitcher) extractVariables(css string) map[string]string {
	// Simple variable extraction
	variables := make(map[string]string)
	// This would be implemented properly with CSS parsing
	return variables
}

func (ts *ThemeSwitcher) extractClasses(css string) []string {
	// Simple class extraction
	var classes []string
	// This would be implemented properly with CSS parsing
	return classes
}

// SimpleThemeStorage provides a simple in-memory storage implementation
type SimpleThemeStorage struct {
	preferences map[string]*UserPreferences
	mu          sync.RWMutex
}

// NewSimpleThemeStorage creates a new simple theme storage
func NewSimpleThemeStorage() *SimpleThemeStorage {
	return &SimpleThemeStorage{
		preferences: make(map[string]*UserPreferences),
	}
}

// SavePreferences saves user preferences
func (s *SimpleThemeStorage) SavePreferences(userID string, prefs *UserPreferences) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.preferences[userID] = prefs
	return nil
}

// LoadPreferences loads user preferences
func (s *SimpleThemeStorage) LoadPreferences(userID string) (*UserPreferences, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if prefs, exists := s.preferences[userID]; exists {
		return prefs, nil
	}
	
	return DefaultUserPreferences(), nil
}

// DeletePreferences deletes user preferences
func (s *SimpleThemeStorage) DeletePreferences(userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.preferences, userID)
	return nil
}