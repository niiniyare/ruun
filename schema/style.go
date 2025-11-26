package schema

// Style defines unified styling for all components.
type Style struct {
	// CSS Classes
	Classes   string `json:"classes,omitempty"`
	Container string `json:"container,omitempty"`
	Wrapper   string `json:"wrapper,omitempty"`
	Label     string `json:"label,omitempty"`
	Input     string `json:"input,omitempty"`
	Help      string `json:"help,omitempty"`
	Error     string `json:"error,omitempty"`

	// Dimensions
	Width   string `json:"width,omitempty"`
	Height  string `json:"height,omitempty"`
	Margin  string `json:"margin,omitempty"`
	Padding string `json:"padding,omitempty"`

	// Visual
	Background   string `json:"background,omitempty"`
	Border       string `json:"border,omitempty"`
	BorderRadius string `json:"borderRadius,omitempty"`
	Shadow       string `json:"shadow,omitempty"`

	// Typography
	FontSize   string `json:"fontSize,omitempty"`
	FontWeight string `json:"fontWeight,omitempty"`
	TextColor  string `json:"textColor,omitempty"`
	TextAlign  string `json:"textAlign,omitempty"`

	// Color tokens (for theming)
	Colors map[string]string `json:"colors,omitempty"`

	// State-specific styles
	States *StateStyles `json:"states,omitempty"`

	// Custom CSS
	CustomCSS string `json:"customCSS,omitempty"`
}

// StateStyles defines styling for different component states.
type StateStyles struct {
	Default   map[string]string `json:"default,omitempty"`
	Hover     map[string]string `json:"hover,omitempty"`
	Focus     map[string]string `json:"focus,omitempty"`
	Active    map[string]string `json:"active,omitempty"`
	Disabled  map[string]string `json:"disabled,omitempty"`
	Error     map[string]string `json:"error,omitempty"`
	Success   map[string]string `json:"success,omitempty"`
	// For tabs/steps
	Inactive  map[string]string `json:"inactive,omitempty"`
	Completed map[string]string `json:"completed,omitempty"`
}

// IsEmpty returns true if the style has no values set.
func (s *Style) IsEmpty() bool {
	if s == nil {
		return true
	}
	return s.Classes == "" && s.Container == "" && s.Wrapper == "" &&
		s.Label == "" && s.Input == "" && s.Help == "" && s.Error == "" &&
		s.Width == "" && s.Height == "" && s.Margin == "" && s.Padding == "" &&
		s.Background == "" && s.Border == "" && s.BorderRadius == "" && s.Shadow == "" &&
		s.FontSize == "" && s.FontWeight == "" && s.TextColor == "" && s.TextAlign == "" &&
		len(s.Colors) == 0 && s.States == nil && s.CustomCSS == ""
}

// Merge combines two styles, with 'other' values taking precedence.
func (s *Style) Merge(other *Style) *Style {
	if s == nil {
		return other
	}
	if other == nil {
		return s
	}

	result := *s

	// CSS Classes
	if other.Classes != "" {
		result.Classes = other.Classes
	}
	if other.Container != "" {
		result.Container = other.Container
	}
	if other.Wrapper != "" {
		result.Wrapper = other.Wrapper
	}
	if other.Label != "" {
		result.Label = other.Label
	}
	if other.Input != "" {
		result.Input = other.Input
	}
	if other.Help != "" {
		result.Help = other.Help
	}
	if other.Error != "" {
		result.Error = other.Error
	}

	// Dimensions
	if other.Width != "" {
		result.Width = other.Width
	}
	if other.Height != "" {
		result.Height = other.Height
	}
	if other.Margin != "" {
		result.Margin = other.Margin
	}
	if other.Padding != "" {
		result.Padding = other.Padding
	}

	// Visual
	if other.Background != "" {
		result.Background = other.Background
	}
	if other.Border != "" {
		result.Border = other.Border
	}
	if other.BorderRadius != "" {
		result.BorderRadius = other.BorderRadius
	}
	if other.Shadow != "" {
		result.Shadow = other.Shadow
	}

	// Typography
	if other.FontSize != "" {
		result.FontSize = other.FontSize
	}
	if other.FontWeight != "" {
		result.FontWeight = other.FontWeight
	}
	if other.TextColor != "" {
		result.TextColor = other.TextColor
	}
	if other.TextAlign != "" {
		result.TextAlign = other.TextAlign
	}

	// Color tokens
	if len(other.Colors) > 0 {
		if result.Colors == nil {
			result.Colors = make(map[string]string)
		}
		for k, v := range other.Colors {
			result.Colors[k] = v
		}
	}

	// State styles
	if other.States != nil {
		result.States = other.States
	}

	// Custom CSS
	if other.CustomCSS != "" {
		result.CustomCSS = other.CustomCSS
	}

	return &result
}