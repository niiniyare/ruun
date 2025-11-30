package schema

// UnifiedStyle defines unified styling for all components.
// Replaces LayoutStyle, SectionStyle, TabStyle, StepStyle, ActionTheme, FieldStyle
type UnifiedStyle struct {
	// CSS Classes
	Classes   string `json:"classes,omitempty"`
	Container string `json:"container,omitempty"`
	Label     string `json:"label,omitempty"`
	Input     string `json:"input,omitempty"`
	Help      string `json:"help,omitempty"`
	Error     string `json:"error,omitempty"`

	// Layout
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

	// Color mappings
	Colors map[string]string `json:"colors,omitempty"`

	// State styles
	States *UnifiedStateStyles `json:"states,omitempty"`

	// Custom CSS
	Custom string `json:"custom,omitempty"`
}

type UnifiedStateStyles struct {
	Default   map[string]string `json:"default,omitempty"`
	Hover     map[string]string `json:"hover,omitempty"`
	Focus     map[string]string `json:"focus,omitempty"`
	Active    map[string]string `json:"active,omitempty"`
	Disabled  map[string]string `json:"disabled,omitempty"`
	Completed map[string]string `json:"completed,omitempty"`
}

func (s *UnifiedStyle) IsEmpty() bool {
	if s == nil {
		return true
	}
	return s.Classes == "" && s.Container == "" && s.Background == "" &&
		s.Border == "" && s.Custom == "" && len(s.Colors) == 0
}