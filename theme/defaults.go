package theme

// GetDefaultTokens returns a comprehensive, production-ready default token set.
func GetDefaultTokens() *Tokens {
	return &Tokens{
		Primitives: getDefaultPrimitives(),
		Semantic:   getDefaultSemantic(),
		Components: getDefaultComponents(),
	}
}

// getDefaultPrimitives returns default primitive tokens.
func getDefaultPrimitives() *PrimitiveTokens {
	return &PrimitiveTokens{
		Colors: map[string]string{
			// Grayscale
			"white":    "#ffffff",
			"black":    "#000000",
			"gray-50":  "#f9fafb",
			"gray-100": "#f3f4f6",
			"gray-200": "#e5e7eb",
			"gray-300": "#d1d5db",
			"gray-400": "#9ca3af",
			"gray-500": "#6b7280",
			"gray-600": "#4b5563",
			"gray-700": "#374151",
			"gray-800": "#1f2937",
			"gray-900": "#111827",

			// Primary (Blue)
			"blue-50":  "#eff6ff",
			"blue-100": "#dbeafe",
			"blue-200": "#bfdbfe",
			"blue-300": "#93c5fd",
			"blue-400": "#60a5fa",
			"blue-500": "#3b82f6",
			"blue-600": "#2563eb",
			"blue-700": "#1d4ed8",
			"blue-800": "#1e40af",
			"blue-900": "#1e3a8a",

			// Success (Green)
			"green-50":  "#f0fdf4",
			"green-100": "#dcfce7",
			"green-200": "#bbf7d0",
			"green-300": "#86efac",
			"green-400": "#4ade80",
			"green-500": "#22c55e",
			"green-600": "#16a34a",
			"green-700": "#15803d",
			"green-800": "#166534",
			"green-900": "#14532d",

			// Warning (Amber)
			"amber-50":  "#fffbeb",
			"amber-100": "#fef3c7",
			"amber-200": "#fde68a",
			"amber-300": "#fcd34d",
			"amber-400": "#fbbf24",
			"amber-500": "#f59e0b",
			"amber-600": "#d97706",
			"amber-700": "#b45309",
			"amber-800": "#92400e",
			"amber-900": "#78350f",

			// Danger (Red)
			"red-50":  "#fef2f2",
			"red-100": "#fee2e2",
			"red-200": "#fecaca",
			"red-300": "#fca5a5",
			"red-400": "#f87171",
			"red-500": "#ef4444",
			"red-600": "#dc2626",
			"red-700": "#b91c1c",
			"red-800": "#991b1b",
			"red-900": "#7f1d1d",

			// Info (Cyan)
			"cyan-50":  "#ecfeff",
			"cyan-100": "#cffafe",
			"cyan-200": "#a5f3fc",
			"cyan-300": "#67e8f9",
			"cyan-400": "#22d3ee",
			"cyan-500": "#06b6d4",
			"cyan-600": "#0891b2",
			"cyan-700": "#0e7490",
			"cyan-800": "#155e75",
			"cyan-900": "#164e63",
		},

		Spacing: map[string]string{
			"0":    "0",
			"0.5":  "0.125rem",
			"1":    "0.25rem",
			"1.5":  "0.375rem",
			"2":    "0.5rem",
			"2.5":  "0.625rem",
			"3":    "0.75rem",
			"3.5":  "0.875rem",
			"4":    "1rem",
			"5":    "1.25rem",
			"6":    "1.5rem",
			"7":    "1.75rem",
			"8":    "2rem",
			"9":    "2.25rem",
			"10":   "2.5rem",
			"11":   "2.75rem",
			"12":   "3rem",
			"14":   "3.5rem",
			"16":   "4rem",
			"20":   "5rem",
			"24":   "6rem",
			"28":   "7rem",
			"32":   "8rem",
			"36":   "9rem",
			"40":   "10rem",
			"44":   "11rem",
			"48":   "12rem",
			"52":   "13rem",
			"56":   "14rem",
			"60":   "15rem",
			"64":   "16rem",
			"72":   "18rem",
			"80":   "20rem",
			"96":   "24rem",
		},

		Radius: map[string]string{
			"none":   "0",
			"sm":     "0.125rem",
			"base":   "0.25rem",
			"md":     "0.375rem",
			"lg":     "0.5rem",
			"xl":     "0.75rem",
			"2xl":    "1rem",
			"3xl":    "1.5rem",
			"full":   "9999px",
		},

		Typography: map[string]string{
			"font-sans":       "ui-sans-serif, system-ui, sans-serif",
			"font-serif":      "ui-serif, Georgia, Cambria, serif",
			"font-mono":       "ui-monospace, SFMono-Regular, Menlo, monospace",
			
			"text-xs":         "0.75rem",
			"text-sm":         "0.875rem",
			"text-base":       "1rem",
			"text-lg":         "1.125rem",
			"text-xl":         "1.25rem",
			"text-2xl":        "1.5rem",
			"text-3xl":        "1.875rem",
			"text-4xl":        "2.25rem",
			"text-5xl":        "3rem",
			"text-6xl":        "3.75rem",
			"text-7xl":        "4.5rem",
			"text-8xl":        "6rem",
			"text-9xl":        "8rem",
			
			"leading-none":    "1",
			"leading-tight":   "1.25",
			"leading-snug":    "1.375",
			"leading-normal":  "1.5",
			"leading-relaxed": "1.625",
			"leading-loose":   "2",
			
			"tracking-tighter": "-0.05em",
			"tracking-tight":   "-0.025em",
			"tracking-normal":  "0",
			"tracking-wide":    "0.025em",
			"tracking-wider":   "0.05em",
			"tracking-widest":  "0.1em",
			
			"font-thin":        "100",
			"font-extralight":  "200",
			"font-light":       "300",
			"font-normal":      "400",
			"font-medium":      "500",
			"font-semibold":    "600",
			"font-bold":        "700",
			"font-extrabold":   "800",
			"font-black":       "900",
		},

		Borders: map[string]string{
			"width-0":      "0",
			"width-1":      "1px",
			"width-2":      "2px",
			"width-4":      "4px",
			"width-8":      "8px",
			
			"style-solid":  "solid",
			"style-dashed": "dashed",
			"style-dotted": "dotted",
			"style-double": "double",
			"style-none":   "none",
		},

		Shadows: map[string]string{
			"none":  "none",
			"xs":    "0 1px 2px 0 rgba(0, 0, 0, 0.05)",
			"sm":    "0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.1)",
			"base":  "0 1px 3px 0 rgba(0, 0, 0, 0.1), 0 1px 2px -1px rgba(0, 0, 0, 0.1)",
			"md":    "0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -2px rgba(0, 0, 0, 0.1)",
			"lg":    "0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -4px rgba(0, 0, 0, 0.1)",
			"xl":    "0 20px 25px -5px rgba(0, 0, 0, 0.1), 0 8px 10px -6px rgba(0, 0, 0, 0.1)",
			"2xl":   "0 25px 50px -12px rgba(0, 0, 0, 0.25)",
			"inner": "inset 0 2px 4px 0 rgba(0, 0, 0, 0.05)",
		},

		Effects: map[string]string{
			"opacity-0":    "0",
			"opacity-5":    "0.05",
			"opacity-10":   "0.1",
			"opacity-20":   "0.2",
			"opacity-25":   "0.25",
			"opacity-30":   "0.3",
			"opacity-40":   "0.4",
			"opacity-50":   "0.5",
			"opacity-60":   "0.6",
			"opacity-70":   "0.7",
			"opacity-75":   "0.75",
			"opacity-80":   "0.8",
			"opacity-90":   "0.9",
			"opacity-95":   "0.95",
			"opacity-100":  "1",
			
			"blur-none":    "0",
			"blur-sm":      "4px",
			"blur-base":    "8px",
			"blur-md":      "12px",
			"blur-lg":      "16px",
			"blur-xl":      "24px",
			"blur-2xl":     "40px",
			"blur-3xl":     "64px",
		},

		Animation: map[string]string{
			"duration-75":    "75ms",
			"duration-100":   "100ms",
			"duration-150":   "150ms",
			"duration-200":   "200ms",
			"duration-300":   "300ms",
			"duration-500":   "500ms",
			"duration-700":   "700ms",
			"duration-1000":  "1000ms",
			
			"ease-linear":    "linear",
			"ease-in":        "cubic-bezier(0.4, 0, 1, 1)",
			"ease-out":       "cubic-bezier(0, 0, 0.2, 1)",
			"ease-in-out":    "cubic-bezier(0.4, 0, 0.2, 1)",
		},

		ZIndex: map[string]string{
			"0":       "0",
			"10":      "10",
			"20":      "20",
			"30":      "30",
			"40":      "40",
			"50":      "50",
			"auto":    "auto",
			"dropdown": "1000",
			"sticky":   "1020",
			"fixed":    "1030",
			"modal":    "1040",
			"popover":  "1050",
			"tooltip":  "1060",
		},

		Breakpoints: map[string]string{
			"xs":   "475px",
			"sm":   "640px",
			"md":   "768px",
			"lg":   "1024px",
			"xl":   "1280px",
			"2xl":  "1536px",
		},
	}
}

// getDefaultSemantic returns default semantic tokens.
func getDefaultSemantic() *SemanticTokens {
	return &SemanticTokens{
		Colors: map[string]string{
			"background":          "primitives.colors.white",
			"background-subtle":   "primitives.colors.gray-50",
			"background-muted":    "primitives.colors.gray-100",
			
			"foreground":          "primitives.colors.gray-900",
			"foreground-subtle":   "primitives.colors.gray-600",
			"foreground-muted":    "primitives.colors.gray-500",
			
			"border":              "primitives.colors.gray-200",
			"border-strong":       "primitives.colors.gray-300",
			
			"primary":             "primitives.colors.blue-600",
			"primary-hover":       "primitives.colors.blue-700",
			"primary-active":      "primitives.colors.blue-800",
			"primary-subtle":      "primitives.colors.blue-50",
			
			"success":             "primitives.colors.green-600",
			"success-hover":       "primitives.colors.green-700",
			"success-subtle":      "primitives.colors.green-50",
			
			"warning":             "primitives.colors.amber-600",
			"warning-hover":       "primitives.colors.amber-700",
			"warning-subtle":      "primitives.colors.amber-50",
			
			"danger":              "primitives.colors.red-600",
			"danger-hover":        "primitives.colors.red-700",
			"danger-subtle":       "primitives.colors.red-50",
			
			"info":                "primitives.colors.cyan-600",
			"info-hover":          "primitives.colors.cyan-700",
			"info-subtle":         "primitives.colors.cyan-50",
		},

		Spacing: map[string]string{
			"page-padding":        "primitives.spacing.6",
			"section-gap":         "primitives.spacing.12",
			"component-gap":       "primitives.spacing.4",
			"content-gap":         "primitives.spacing.3",
			"inline-gap":          "primitives.spacing.2",
		},

		Typography: map[string]string{
			"heading-1":           "primitives.typography.text-4xl",
			"heading-2":           "primitives.typography.text-3xl",
			"heading-3":           "primitives.typography.text-2xl",
			"heading-4":           "primitives.typography.text-xl",
			"heading-5":           "primitives.typography.text-lg",
			"heading-6":           "primitives.typography.text-base",
			"body":                "primitives.typography.text-base",
			"body-sm":             "primitives.typography.text-sm",
			"caption":             "primitives.typography.text-xs",
			"label":               "primitives.typography.text-sm",
		},

		Interactive: map[string]string{
			"focus-ring":          "primitives.colors.blue-500",
			"focus-ring-offset":   "primitives.spacing.0.5",
			"disabled-opacity":    "primitives.effects.opacity-50",
			"hover-opacity":       "primitives.effects.opacity-90",
		},
	}
}

// getDefaultComponents returns default component tokens.
func getDefaultComponents() *ComponentTokens {
	return &ComponentTokens{
		"button": {
			"primary": {
				"background":       "semantic.colors.primary",
				"color":            "primitives.colors.white",
				"border":           "primitives.borders.width-0",
				"border-radius":    "primitives.radius.md",
				"padding":          "primitives.spacing.2 primitives.spacing.4",
				"font-size":        "semantic.typography.body-sm",
				"font-weight":      "primitives.typography.font-medium",
				"transition":       "primitives.animation.duration-150",
			},
			"secondary": {
				"background":       "primitives.colors.white",
				"color":            "semantic.colors.foreground",
				"border":           "primitives.borders.width-1 primitives.borders.style-solid semantic.colors.border",
				"border-radius":    "primitives.radius.md",
				"padding":          "primitives.spacing.2 primitives.spacing.4",
				"font-size":        "semantic.typography.body-sm",
				"font-weight":      "primitives.typography.font-medium",
			},
			"ghost": {
				"background":       "transparent",
				"color":            "semantic.colors.foreground",
				"border":           "primitives.borders.width-0",
				"border-radius":    "primitives.radius.md",
				"padding":          "primitives.spacing.2 primitives.spacing.4",
				"font-size":        "semantic.typography.body-sm",
				"font-weight":      "primitives.typography.font-medium",
			},
			"danger": {
				"background":       "semantic.colors.danger",
				"color":            "primitives.colors.white",
				"border":           "primitives.borders.width-0",
				"border-radius":    "primitives.radius.md",
				"padding":          "primitives.spacing.2 primitives.spacing.4",
				"font-size":        "semantic.typography.body-sm",
				"font-weight":      "primitives.typography.font-medium",
			},
		},
		
		"input": {
			"default": {
				"background":       "primitives.colors.white",
				"color":            "semantic.colors.foreground",
				"border":           "primitives.borders.width-1 primitives.borders.style-solid semantic.colors.border",
				"border-radius":    "primitives.radius.md",
				"padding":          "primitives.spacing.2 primitives.spacing.3",
				"font-size":        "semantic.typography.body",
			},
			"error": {
				"border":           "primitives.borders.width-1 primitives.borders.style-solid semantic.colors.danger",
			},
		},
		
		"card": {
			"default": {
				"background":       "primitives.colors.white",
				"border":           "primitives.borders.width-1 primitives.borders.style-solid semantic.colors.border",
				"border-radius":    "primitives.radius.lg",
				"padding":          "primitives.spacing.6",
				"shadow":           "primitives.shadows.sm",
			},
		},
		
		"badge": {
			"default": {
				"background":       "semantic.colors.background-muted",
				"color":            "semantic.colors.foreground",
				"border-radius":    "primitives.radius.full",
				"padding":          "primitives.spacing.1 primitives.spacing.2.5",
				"font-size":        "semantic.typography.caption",
				"font-weight":      "primitives.typography.font-medium",
			},
			"primary": {
				"background":       "semantic.colors.primary-subtle",
				"color":            "semantic.colors.primary",
			},
			"success": {
				"background":       "semantic.colors.success-subtle",
				"color":            "semantic.colors.success",
			},
			"warning": {
				"background":       "semantic.colors.warning-subtle",
				"color":            "semantic.colors.warning",
			},
			"danger": {
				"background":       "semantic.colors.danger-subtle",
				"color":            "semantic.colors.danger",
			},
		},
	}
}

// GetDefaultTheme returns a complete default theme ready for production use.
func GetDefaultTheme() *Theme {
	return &Theme{
		ID:          "default",
		Name:        "Default Theme",
		Description: "Production-ready default theme with comprehensive token coverage",
		Version:     "1.0.0",
		Author:      "Awo ERP",
		Tokens:      GetDefaultTokens(),
		DarkMode: &DarkModeConfig{
			Enabled:  true,
			Default:  false,
			Strategy: "class",
			DarkTokens: &Tokens{
				Primitives: &PrimitiveTokens{
					Colors: map[string]string{
						"white":     "#ffffff",
						"black":     "#000000",
						"gray-50":   "#18181b",
						"gray-100":  "#27272a",
						"gray-200":  "#3f3f46",
						"gray-300":  "#52525b",
						"gray-400":  "#71717a",
						"gray-500":  "#a1a1aa",
						"gray-600":  "#d4d4d8",
						"gray-700":  "#e4e4e7",
						"gray-800":  "#f4f4f5",
						"gray-900":  "#fafafa",
					},
				},
				Semantic: &SemanticTokens{
					Colors: map[string]string{
						"background":        "primitives.colors.gray-50",
						"background-subtle": "primitives.colors.gray-100",
						"background-muted":  "primitives.colors.gray-200",
						"foreground":        "primitives.colors.gray-900",
						"foreground-subtle": "primitives.colors.gray-600",
						"border":            "primitives.colors.gray-300",
					},
				},
			},
		},
		Accessibility: &AccessibilityConfig{
			HighContrast:      false,
			MinContrastRatio:  4.5,
			FocusIndicator:    true,
			FocusOutlineColor: "#3b82f6",
			FocusOutlineWidth: "2px",
			KeyboardNav:       true,
			ReducedMotion:     false,
			ScreenReader:      true,
			AriaLive:          "polite",
		},
		Metadata: &ThemeMetadata{
			Tags:    []string{"default", "production", "accessible"},
			License: "MIT",
		},
	}
}
