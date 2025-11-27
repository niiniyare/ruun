package main

import (
	"strings"
)

// ButtonVariant defines the visual style of a button
type ButtonVariant string

const (
	ButtonPrimary     ButtonVariant = "primary"
	ButtonSecondary   ButtonVariant = "secondary"
	ButtonOutline     ButtonVariant = "outline"
	ButtonGhost       ButtonVariant = "ghost"
	ButtonLink        ButtonVariant = "link"
	ButtonDestructive ButtonVariant = "destructive"
)

// ButtonSize defines the size of a button
type ButtonSize string

const (
	ButtonSizeSm ButtonSize = "sm"
	ButtonSizeMd ButtonSize = "md"
	ButtonSizeLg ButtonSize = "lg"
)

// ButtonType defines whether the button is text, icon-only, or has both
type ButtonType string

const (
	ButtonTypeDefault ButtonType = "default" // text with optional icon
	ButtonTypeIcon    ButtonType = "icon"    // icon only
)

// ButtonProps defines all properties for a button component
type ButtonProps struct {
	Label     string
	Icon      string
	Variant   ButtonVariant
	Size      ButtonSize
	Type      ButtonType
	Disabled  bool
	Loading   bool
	ClassName string
	ID        string
	HTMLType  string
	OnClick   string
	HxPost    string
	HxGet     string
	HxTarget  string
	HxSwap    string
	AriaLabel string
}

// BadgeVariant defines the visual style of a badge
type BadgeVariant string

const (
	BadgePrimary     BadgeVariant = "primary"
	BadgeSecondary   BadgeVariant = "secondary"
	BadgeDestructive BadgeVariant = "destructive"
	BadgeOutline     BadgeVariant = "outline"
)

// BadgeProps defines all properties for a badge component
type BadgeProps struct {
	Label   string
	Icon    string
	Variant BadgeVariant
}

// AlertVariant defines the visual style and semantic meaning of an alert
type AlertVariant string

const (
	AlertDefault     AlertVariant = "default"
	AlertDestructive AlertVariant = "destructive"
	AlertSuccess     AlertVariant = "success"
	AlertWarning     AlertVariant = "warning"
	AlertInfo        AlertVariant = "info"
)

// AlertProps defines all properties for an alert component
type AlertProps struct {
	Title   string
	Message string
	Variant AlertVariant
}

// InputType defines the HTML input type
type InputType string

const (
	InputTypeText     InputType = "text"
	InputTypeEmail    InputType = "email"
	InputTypePassword InputType = "password"
	InputTypeNumber   InputType = "number"
	InputTypeSearch   InputType = "search"
	InputTypeDate     InputType = "date"
	InputTypeFile     InputType = "file"
	InputTypeTel      InputType = "tel"
	InputTypeUrl      InputType = "url"
)

// InputProps defines all properties for an input component
type InputProps struct {
	Name        string
	Type        InputType
	Value       string
	Placeholder string
	Required    bool
	Disabled    bool
}

// LabelProps defines all properties for a label component
type LabelProps struct {
	Text     string
	For      string
	Required bool
	Optional bool
}

// Helper functions for generating CSS classes
func GetButtonClasses(variant ButtonVariant, size ButtonSize, buttonType ButtonType, disabled bool, loading bool) string {
	var classes []string

	// Base class with variant and type
	baseClass := "btn"
	if size != ButtonSizeMd {
		baseClass += "-" + string(size)
	}
	if buttonType == ButtonTypeIcon {
		baseClass += "-icon"
	}
	if variant != ButtonPrimary {
		baseClass += "-" + string(variant)
	}

	classes = append(classes, baseClass)

	// State classes
	if disabled {
		classes = append(classes, "disabled")
	}
	if loading {
		classes = append(classes, "loading")
	}

	return strings.Join(classes, " ")
}

func GetBadgeClasses(variant BadgeVariant) string {
	baseClass := "badge"
	if variant != BadgePrimary {
		return baseClass + "-" + string(variant)
	}
	return baseClass
}

func GetAlertClasses(variant AlertVariant) string {
	baseClass := "alert"
	if variant != AlertDefault {
		return baseClass + "-" + string(variant)
	}
	return baseClass
}

func GetInputClasses(inputType InputType, hasError bool, disabled bool, readOnly bool) string {
	return "input"
}

func GetLabelClasses() string {
	return "label"
}

// Default props functions
func GetDefaultButtonProps() ButtonProps {
	return ButtonProps{
		Variant:  ButtonPrimary,
		Size:     ButtonSizeMd,
		Type:     ButtonTypeDefault,
		HTMLType: "button",
		Disabled: false,
		Loading:  false,
	}
}

func GetDefaultBadgeProps() BadgeProps {
	return BadgeProps{
		Variant: BadgePrimary,
	}
}

func GetDefaultAlertProps() AlertProps {
	return AlertProps{
		Variant: AlertDefault,
	}
}

func GetDefaultInputProps() InputProps {
	return InputProps{
		Type:     InputTypeText,
		Required: false,
		Disabled: false,
	}
}

func GetDefaultLabelProps() LabelProps {
	return LabelProps{
		Required: false,
		Optional: false,
	}
}
