package theme

import (
	"encoding/json"
	"fmt"
	"maps"
	"time"

	"gopkg.in/yaml.v3"
)

type Theme struct {
	ID          string `json:"id" yaml:"id"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Version     string `json:"version,omitempty" yaml:"version,omitempty"`
	Author      string `json:"author,omitempty" yaml:"author,omitempty"`

	Tokens        *Tokens              `json:"tokens" yaml:"tokens"`
	DarkMode      *DarkModeConfig      `json:"darkMode,omitempty" yaml:"darkMode,omitempty"`
	Conditions    []*Condition         `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Accessibility *AccessibilityConfig `json:"accessibility,omitempty" yaml:"accessibility,omitempty"`
	CustomCSS     string               `json:"customCSS,omitempty" yaml:"customCSS,omitempty"`
	CustomJS      string               `json:"customJS,omitempty" yaml:"customJS,omitempty"`
	Metadata      *ThemeMetadata       `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	createdAt time.Time
	updatedAt time.Time
}

type DarkModeConfig struct {
	Enabled    bool    `json:"enabled" yaml:"enabled"`
	Default    bool    `json:"default,omitempty" yaml:"default,omitempty"`
	Strategy   string  `json:"strategy,omitempty" yaml:"strategy,omitempty"`
	DarkTokens *Tokens `json:"darkTokens,omitempty" yaml:"darkTokens,omitempty"`
}

type Condition struct {
	ID          string            `json:"id" yaml:"id"`
	Expression  string            `json:"expression" yaml:"expression"`
	Priority    int               `json:"priority" yaml:"priority"`
	Overrides   map[string]string `json:"overrides" yaml:"overrides"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
}

type AccessibilityConfig struct {
	HighContrast      bool    `json:"highContrast,omitempty" yaml:"highContrast,omitempty"`
	MinContrastRatio  float64 `json:"minContrastRatio,omitempty" yaml:"minContrastRatio,omitempty"`
	FocusIndicator    bool    `json:"focusIndicator,omitempty" yaml:"focusIndicator,omitempty"`
	FocusOutlineColor string  `json:"focusOutlineColor,omitempty" yaml:"focusOutlineColor,omitempty"`
	FocusOutlineWidth string  `json:"focusOutlineWidth,omitempty" yaml:"focusOutlineWidth,omitempty"`
	KeyboardNav       bool    `json:"keyboardNav,omitempty" yaml:"keyboardNav,omitempty"`
	ReducedMotion     bool    `json:"reducedMotion,omitempty" yaml:"reducedMotion,omitempty"`
	ScreenReader      bool    `json:"screenReader,omitempty" yaml:"screenReader,omitempty"`
	AriaLive          string  `json:"ariaLive,omitempty" yaml:"ariaLive,omitempty"`
}

type ThemeMetadata struct {
	Tags       []string       `json:"tags,omitempty" yaml:"tags,omitempty"`
	License    string         `json:"license,omitempty" yaml:"license,omitempty"`
	Repository string         `json:"repository,omitempty" yaml:"repository,omitempty"`
	Homepage   string         `json:"homepage,omitempty" yaml:"homepage,omitempty"`
	Preview    string         `json:"preview,omitempty" yaml:"preview,omitempty"`
	CustomData map[string]any `json:"customData,omitempty" yaml:"customData,omitempty"`
	CreatedAt  time.Time      `json:"createdAt" yaml:"createdAt,omitempty"`
	UpdatedAt  time.Time      `json:"updatedAt" yaml:"updatedAt,omitempty"`
}

func (t *Theme) Validate() error {
	if t.ID == "" {
		return NewError(ErrCodeValidation, "theme ID cannot be empty")
	}
	if t.Name == "" {
		return NewError(ErrCodeValidation, "theme name cannot be empty")
	}
	if t.Tokens == nil {
		return NewError(ErrCodeValidation, "theme tokens cannot be nil")
	}

	if err := t.Tokens.Validate(); err != nil {
		return WrapError(ErrCodeValidation, "invalid tokens", err)
	}

	if t.DarkMode != nil {
		if err := t.DarkMode.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid dark mode", err)
		}
	}

	for i, cond := range t.Conditions {
		if err := cond.Validate(); err != nil {
			return WrapError(ErrCodeValidation,
				fmt.Sprintf("invalid condition %d (%s)", i, cond.ID), err)
		}
	}

	if t.Accessibility != nil {
		if err := t.Accessibility.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid accessibility", err)
		}
	}

	const maxCustomCodeSize = 1 << 20
	if len(t.CustomCSS) > maxCustomCodeSize {
		return NewErrorf(ErrCodeValidation,
			"custom CSS exceeds maximum size of %d bytes", maxCustomCodeSize)
	}
	if len(t.CustomJS) > maxCustomCodeSize {
		return NewErrorf(ErrCodeValidation,
			"custom JS exceeds maximum size of %d bytes", maxCustomCodeSize)
	}

	return nil
}

func (d *DarkModeConfig) Validate() error {
	if d == nil {
		return nil
	}

	if d.Strategy != "" {
		validStrategies := map[string]bool{
			"class": true,
			"media": true,
			"auto":  true,
		}
		if !validStrategies[d.Strategy] {
			return NewErrorf(ErrCodeValidation,
				"invalid dark mode strategy: %s (must be class, media, or auto)", d.Strategy)
		}
	}

	if d.DarkTokens != nil {
		if err := d.DarkTokens.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid dark tokens", err)
		}
	}

	return nil
}

func (c *Condition) Validate() error {
	if c.ID == "" {
		return NewError(ErrCodeValidation, "condition ID cannot be empty")
	}
	if c.Expression == "" {
		return NewError(ErrCodeValidation, "expression cannot be empty")
	}
	if len(c.Overrides) == 0 {
		return NewError(ErrCodeValidation, "overrides cannot be empty")
	}

	for path, value := range c.Overrides {
		if path == "" {
			return NewError(ErrCodeValidation, "override path cannot be empty")
		}
		ref := TokenReference(value)
		if err := ref.Validate(); err != nil {
			return WrapError(ErrCodeValidation,
				fmt.Sprintf("invalid override value for '%s'", path), err)
		}
	}

	return nil
}

func (a *AccessibilityConfig) Validate() error {
	if a == nil {
		return nil
	}

	if a.MinContrastRatio < 0 || a.MinContrastRatio > 21 {
		return NewErrorf(ErrCodeValidation,
			"contrast ratio must be between 0 and 21, got: %f", a.MinContrastRatio)
	}

	if a.AriaLive != "" {
		validValues := map[string]bool{
			"off":       true,
			"polite":    true,
			"assertive": true,
		}
		if !validValues[a.AriaLive] {
			return NewErrorf(ErrCodeValidation,
				"invalid ARIA live value: %s (must be off, polite, or assertive)", a.AriaLive)
		}
	}

	return nil
}

func (t *Theme) ToYAML() ([]byte, error) {
	return yaml.Marshal(t)
}

func ThemeFromYAML(data []byte) (*Theme, error) {
	var theme Theme
	if err := yaml.Unmarshal(data, &theme); err != nil {
		return nil, WrapError(ErrCodeValidation, "failed to parse theme YAML", err)
	}
	if err := theme.Validate(); err != nil {
		return nil, WrapError(ErrCodeValidation, "invalid theme", err)
	}
	return &theme, nil
}

func (t *Theme) Clone() *Theme {
	if t == nil {
		return nil
	}

	cloned := &Theme{
		ID:          t.ID,
		Name:        t.Name,
		Description: t.Description,
		Version:     t.Version,
		Author:      t.Author,
		CustomCSS:   t.CustomCSS,
		CustomJS:    t.CustomJS,
		createdAt:   t.createdAt,
		updatedAt:   t.updatedAt,
	}

	if t.Tokens != nil {
		cloned.Tokens = t.Tokens.Clone()
	}
	if t.DarkMode != nil {
		cloned.DarkMode = t.DarkMode.Clone()
	}
	if t.Accessibility != nil {
		cloned.Accessibility = t.Accessibility.Clone()
	}
	if t.Metadata != nil {
		cloned.Metadata = t.Metadata.Clone()
	}
	if len(t.Conditions) > 0 {
		cloned.Conditions = make([]*Condition, len(t.Conditions))
		for i, cond := range t.Conditions {
			cloned.Conditions[i] = cond.Clone()
		}
	}

	return cloned
}

func (d *DarkModeConfig) Clone() *DarkModeConfig {
	if d == nil {
		return nil
	}
	cloned := &DarkModeConfig{
		Enabled:  d.Enabled,
		Default:  d.Default,
		Strategy: d.Strategy,
	}
	if d.DarkTokens != nil {
		cloned.DarkTokens = d.DarkTokens.Clone()
	}
	return cloned
}

func (c *Condition) Clone() *Condition {
	if c == nil {
		return nil
	}
	return &Condition{
		ID:          c.ID,
		Expression:  c.Expression,
		Priority:    c.Priority,
		Overrides:   cloneStringMap(c.Overrides),
		Description: c.Description,
	}
}

func (a *AccessibilityConfig) Clone() *AccessibilityConfig {
	if a == nil {
		return nil
	}
	cloned := *a
	return &cloned
}

func (m *ThemeMetadata) Clone() *ThemeMetadata {
	if m == nil {
		return nil
	}
	cloned := &ThemeMetadata{
		License:    m.License,
		Repository: m.Repository,
		Homepage:   m.Homepage,
		Preview:    m.Preview,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	if len(m.Tags) > 0 {
		cloned.Tags = make([]string, len(m.Tags))
		copy(cloned.Tags, m.Tags)
	}
	if m.CustomData != nil {
		cloned.CustomData = make(map[string]any, len(m.CustomData))
		maps.Copy(cloned.CustomData, m.CustomData)
	}
	return cloned
}

func (t *Theme) ToJSON() ([]byte, error) {
	return json.MarshalIndent(t, "", "  ")
}

func ThemeFromJSON(data []byte) (*Theme, error) {
	var theme Theme
	if err := json.Unmarshal(data, &theme); err != nil {
		return nil, WrapError(ErrCodeValidation, "failed to parse theme JSON", err)
	}
	if err := theme.Validate(); err != nil {
		return nil, WrapError(ErrCodeValidation, "invalid theme", err)
	}
	return &theme, nil
}

func (t *Theme) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Theme) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Theme) SetCreatedAt(tm time.Time) {
	t.createdAt = tm
}

func (t *Theme) SetUpdatedAt(tm time.Time) {
	t.updatedAt = tm
}
