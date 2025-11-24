package static

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type JSComponent struct {
	Minified   string `json:"minified"`
	Unminified string `json:"unminified"`
	Version    string `json:"version"`
	Type       string `json:"type"`
}

type CSSComponent struct {
	Minified   string `json:"minified"`
	Unminified string `json:"unminified"`
	Version    string `json:"version"`
}

type JSManifest struct {
	Components map[string]JSComponent  `json:"components"`
	CSS        map[string]CSSComponent `json:"css"`
}

var jsManifest *JSManifest

// LoadJSManifest loads the JavaScript manifest file
func LoadJSManifest(rootPath string) error {
	manifestPath := filepath.Join(rootPath, "static", "js-manifest.json")
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read JS manifest: %w", err)
	}

	var manifest JSManifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return fmt.Errorf("failed to parse JS manifest: %w", err)
	}

	jsManifest = &manifest
	return nil
}

// GetJSPath returns the minified JS path for a component
func GetJSPath(componentName string) string {
	if jsManifest == nil {
		// Fallback if manifest not loaded
		return fmt.Sprintf("/static/dist/%s.min.js", componentName)
	}

	component, exists := jsManifest.Components[componentName]
	if !exists {
		// Fallback if component not found
		return fmt.Sprintf("/static/dist/%s.min.js", componentName)
	}

	return component.Minified
}

// GetJSPathDebug returns the unminified JS path for debugging
func GetJSPathDebug(componentName string) string {
	if jsManifest == nil {
		// Fallback if manifest not loaded
		return fmt.Sprintf("/static/dist/%s.js", componentName)
	}

	component, exists := jsManifest.Components[componentName]
	if !exists {
		// Fallback if component not found
		return fmt.Sprintf("/static/dist/%s.js", componentName)
	}

	return component.Unminified
}

// GetCSSPath returns the minified CSS path for a component
func GetCSSPath(componentName string) string {
	if jsManifest == nil {
		// Fallback if manifest not loaded
		return fmt.Sprintf("/static/dist/css/%s.min.css", componentName)
	}

	component, exists := jsManifest.CSS[componentName]
	if !exists {
		// Fallback if component not found
		return fmt.Sprintf("/static/dist/css/%s.min.css", componentName)
	}

	return component.Minified
}

// GetCSSPathDebug returns the unminified CSS path for debugging
func GetCSSPathDebug(componentName string) string {
	if jsManifest == nil {
		// Fallback if manifest not loaded
		return fmt.Sprintf("/static/dist/css/%s.css", componentName)
	}

	component, exists := jsManifest.CSS[componentName]
	if !exists {
		// Fallback if component not found
		return fmt.Sprintf("/static/dist/css/%s.css", componentName)
	}

	return component.Unminified
}

// GetJSComponent returns the full component info
func GetJSComponent(componentName string) (JSComponent, bool) {
	if jsManifest == nil {
		return JSComponent{}, false
	}

	component, exists := jsManifest.Components[componentName]
	return component, exists
}

// GetDataTableJS returns the path to the datatable bundle (minified)
func GetDataTableJS() string {
	return GetJSPath("datatable")
}

// GetDataTableJSDebug returns the path to the datatable bundle (unminified)
func GetDataTableJSDebug() string {
	return GetJSPathDebug("datatable")
}

// GetBasecoatJS returns the path to the basecoat core component (minified)
func GetBasecoatJS() string {
	return GetJSPath("basecoat")
}

// GetBasecoatJSDebug returns the path to the basecoat core component (unminified)
func GetBasecoatJSDebug() string {
	return GetJSPathDebug("basecoat")
}

// GetComponentJS returns the path for a specific basecoat component (minified)
func GetComponentJS(componentName string) string {
	return GetJSPath(componentName)
}

// GetComponentJSDebug returns the path for a specific basecoat component (unminified)
func GetComponentJSDebug(componentName string) string {
	return GetJSPathDebug(componentName)
}

// GetBaseCSS returns the path to base CSS (minified)
func GetBaseCSS() string {
	return GetCSSPath("base")
}

// GetBaseCSSDebug returns the path to base CSS (unminified)
func GetBaseCSSDebug() string {
	return GetCSSPathDebug("base")
}