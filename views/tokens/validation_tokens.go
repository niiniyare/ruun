package tokens

import "github.com/niiniyare/ruun/pkg/schema"

// ValidationTokens provides comprehensive validation token definitions
// Three-tier architecture: Primitives → Semantic → Components
var ValidationTokens = &schema.DesignTokens{
	Primitives: &schema.PrimitiveTokens{
		Colors: &schema.ColorPrimitives{
			// Base color palette for validation states
			Gray: &schema.GrayScale{
				Scale50:  "#f9fafb",
				Scale100: "#f3f4f6", 
				Scale200: "#e5e7eb",
				Scale300: "#d1d5db",
				Scale400: "#9ca3af",
				Scale500: "#6b7280",
				Scale600: "#4b5563",
				Scale700: "#374151",
				Scale800: "#1f2937",
				Scale900: "#111827",
			},
			Blue: &schema.ColorScale{
				Scale50:  "#eff6ff",
				Scale100: "#dbeafe",
				Scale200: "#bfdbfe", 
				Scale300: "#93c5fd",
				Scale400: "#60a5fa",
				Scale500: "#3b82f6",
				Scale600: "#2563eb",
				Scale700: "#1d4ed8",
				Scale800: "#1e40af",
				Scale900: "#1e3a8a",
			},
			Green: &schema.ColorScale{
				Scale50:  "#f0fdf4",
				Scale100: "#dcfce7",
				Scale200: "#bbf7d0",
				Scale300: "#86efac",
				Scale400: "#4ade80",
				Scale500: "#22c55e",
				Scale600: "#16a34a",
				Scale700: "#15803d",
				Scale800: "#166534",
				Scale900: "#14532d",
			},
			Red: &schema.ColorScale{
				Scale50:  "#fef2f2",
				Scale100: "#fee2e2",
				Scale200: "#fecaca",
				Scale300: "#f87171",
				Scale400: "#f56565",
				Scale500: "#dc2626",
				Scale600: "#b91c1c",
				Scale700: "#991b1b",
				Scale800: "#7f1d1d",
				Scale900: "#450a0a",
			},
			Yellow: &schema.ColorScale{
				Scale50:  "#fefce8",
				Scale100: "#fef3c7",
				Scale200: "#fde68a",
				Scale300: "#fcd34d",
				Scale400: "#fbbf24",
				Scale500: "#f59e0b",
				Scale600: "#d97706",
				Scale700: "#b45309",
				Scale800: "#92400e",
				Scale900: "#78350f",
			},
			// Pure colors for contrast
			White: "#ffffff",
			Black: "#000000",
		},
		
		Spacing: &schema.SpacingScale{
			None: "0",
			XS:   "0.25rem",
			SM:   "0.5rem", 
			MD:   "0.75rem",
			LG:   "1rem",
			XL:   "1.5rem",
			XXL:  "2rem",
			XXXL: "3rem",
			Huge: "4rem",
		},
		
		Typography: &schema.TypographyPrimitives{
			FontSizes: &schema.FontSizeScale{
				XS:   "0.75rem",
				SM:   "0.875rem", 
				Base: "1rem",
				LG:   "1.125rem",
				XL:   "1.25rem",
				XXL:  "1.5rem",
				XXXL: "1.875rem",
				Huge: "2.25rem",
			},
			FontWeights: &schema.FontWeightScale{
				Thin:      "100",
				Light:     "300",
				Normal:    "400", 
				Medium:    "500",
				Semibold:  "600",
				Bold:      "700",
				Extrabold: "800",
			},
		},
		
		Borders: &schema.BorderPrimitives{
			Width: &schema.BorderWidthScale{
				None:   "0",
				Thin:   "1px",
				Medium: "2px", 
				Thick:  "4px",
			},
			Radius: &schema.BorderRadiusScale{
				None: "0",
				SM:   "0.125rem",
				MD:   "0.25rem",
				LG:   "0.5rem",
				XL:   "1rem",
				Full: "9999px",
			},
		},
		
		Shadows: &schema.ShadowScale{
			None:  "none",
			SM:    "0 1px 2px 0 rgb(0 0 0 / 0.05)",
			MD:    "0 1px 3px 0 rgb(0 0 0 / 0.1), 0 1px 2px -1px rgb(0 0 0 / 0.1)",
			LG:    "0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1)",
			XL:    "0 10px 15px -3px rgb(0 0 0 / 0.1), 0 4px 6px -4px rgb(0 0 0 / 0.1)",
			XXL:   "0 20px 25px -5px rgb(0 0 0 / 0.1), 0 8px 10px -6px rgb(0 0 0 / 0.1)",
			Inner: "inset 0 2px 4px 0 rgb(0 0 0 / 0.06)",
		},
		
		Animations: &schema.AnimationPrimitives{
			Duration: &schema.AnimationDurationScale{
				Fast:   "150ms",
				Normal: "300ms",
				Slow:   "500ms",
			},
			Easing: &schema.AnimationEasingScale{
				Linear:    "linear",
				EaseIn:    "cubic-bezier(0.4, 0.0, 1, 1)",
				EaseOut:   "cubic-bezier(0, 0, 0.2, 1)", 
				EaseInOut: "cubic-bezier(0.4, 0, 0.2, 1)",
			},
		},
	},
	
	Semantic: &schema.SemanticTokens{
		Colors: &schema.SemanticColors{
			Background: &schema.BackgroundColors{
				Default:  schema.TokenReference("primitives.colors.white"),
				Subtle:   schema.TokenReference("primitives.colors.gray.50"),
				Emphasis: schema.TokenReference("primitives.colors.gray.900"),
				Overlay:  schema.TokenReference("primitives.colors.gray.800"),
			},
			Text: &schema.TextColors{
				Default:   schema.TokenReference("primitives.colors.gray.900"),
				Subtle:    schema.TokenReference("primitives.colors.gray.600"),
				Disabled:  schema.TokenReference("primitives.colors.gray.400"),
				Inverted:  schema.TokenReference("primitives.colors.white"),
				Link:      schema.TokenReference("primitives.colors.blue.600"),
				LinkHover: schema.TokenReference("primitives.colors.blue.700"),
			},
			Border: &schema.BorderColors{
				Default: schema.TokenReference("primitives.colors.gray.200"),
				Focus:   schema.TokenReference("primitives.colors.blue.600"),
				Strong:  schema.TokenReference("primitives.colors.gray.300"),
				Subtle:  schema.TokenReference("primitives.colors.gray.100"),
			},
			Interactive: &schema.InteractiveColors{
				Primary: &schema.InteractiveColorSet{
					Default: schema.TokenReference("primitives.colors.blue.600"),
					Hover:   schema.TokenReference("primitives.colors.blue.700"),
					Active:  schema.TokenReference("primitives.colors.blue.800"),
					Focus:   schema.TokenReference("primitives.colors.blue.600"),
				},
			},
			Feedback: &schema.FeedbackColors{
				Success: &schema.FeedbackColorSet{
					Default: schema.TokenReference("primitives.colors.green.600"),
					Subtle:  schema.TokenReference("primitives.colors.green.100"),
					Strong:  schema.TokenReference("primitives.colors.green.800"),
				},
				Error: &schema.FeedbackColorSet{
					Default: schema.TokenReference("primitives.colors.red.600"),
					Subtle:  schema.TokenReference("primitives.colors.red.100"), 
					Strong:  schema.TokenReference("primitives.colors.red.800"),
				},
				Warning: &schema.FeedbackColorSet{
					Default: schema.TokenReference("primitives.colors.yellow.600"),
					Subtle:  schema.TokenReference("primitives.colors.yellow.100"),
					Strong:  schema.TokenReference("primitives.colors.yellow.800"),
				},
				Info: &schema.FeedbackColorSet{
					Default: schema.TokenReference("primitives.colors.blue.600"),
					Subtle:  schema.TokenReference("primitives.colors.blue.100"),
					Strong:  schema.TokenReference("primitives.colors.blue.800"),
				},
			},
		},
		
		Typography: &schema.SemanticTypography{
			Body: &schema.BodyTokens{
				Large: &schema.TypographyToken{
					FontSize:   schema.TokenReference("primitives.typography.fontSizes.lg"),
					FontWeight: schema.TokenReference("primitives.typography.fontWeights.normal"),
				},
				Default: &schema.TypographyToken{
					FontSize:   schema.TokenReference("primitives.typography.fontSizes.base"),
					FontWeight: schema.TokenReference("primitives.typography.fontWeights.normal"),
				},
				Small: &schema.TypographyToken{
					FontSize:   schema.TokenReference("primitives.typography.fontSizes.sm"),
					FontWeight: schema.TokenReference("primitives.typography.fontWeights.normal"),
				},
			},
		},
	},
}