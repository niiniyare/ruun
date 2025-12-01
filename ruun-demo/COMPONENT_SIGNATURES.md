# Component Signatures

Extracted from actual .templ files. Use these EXACT signatures in the demo.

## Atoms

### badge
```go
type BadgeProps struct {
	// Content
	Text string          `json:"text"`
	Icon templ.Component `json:"icon,omitempty"`
	
	// Basecoat variants (using shared types)
	Variant components.BadgeVariant `json:"variant,omitempty"`
	Size    components.Size         `json:"size,omitempty"`
	
	// Interactive features
	Clickable bool   `json:"clickable,omitempty"`
	Href      string `json:"href,omitempty"`
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for Tailwind utilities)
	Base components.BaseProps `json:"base,omitempty"`
}

// getBadgeClass returns the correct single Basecoat class
// Uses exact patterns from documentation - no dynamic class building
func getBadgeClass(variant components.BadgeVariant, size components.Size) string {
	// Static switch-based approach following documentation exactly
	switch size {
	case components.SizeSm:
		switch variant {
		case components.BadgeSecondary:
			return "badge-sm-secondary"
		case components.BadgeDestructive:
			return "badge-sm-destructive"
templ Badge(props BadgeProps) {
```

### layout
```go
type FlexProps struct {
    Direction FlexDirection
    Wrap      FlexWrap
    Justify   Justify
    Align     Align
    Gap       Spacing
    Padding   Spacing
    Class     string
    Attrs     templ.Attributes
}

templ Flex(props FlexProps) {
    <div
        class={ templ.Classes(
            "flex",
            string(props.Direction),
            string(props.Wrap),
            string(props.Justify),
            string(props.Align),
            string(props.Gap),
            string(props.Padding),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { children... }
    </div>
}

// =============================================================================
// Grid Component
--
type GridProps struct {
    Cols    int
    Rows    int
    Gap     Spacing
    Padding Spacing
    Class   string
    Attrs   templ.Attributes
}

func gridColsClass(cols int) string {
    switch cols {
    case 1:
        return "grid-cols-1"
    case 2:
        return "grid-cols-2"
    case 3:
        return "grid-cols-3"
    case 4:
templ Flex(props FlexProps) {
templ Grid(props GridProps) {
templ Stack(props StackProps) {
templ Box(props BoxProps) {
templ Center(props CenterProps) {
templ Spacer(props SpacerProps) {
templ Divider(props DividerProps) {
templ Separator(props SeparatorProps) {
```

### input
```go
type InputProps struct {
    // Core HTML attributes
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Type        components.InputType `json:"type"`
    Value       string    `json:"value"`
    Placeholder string    `json:"placeholder"`
    
    // Constraints and validation
    Required    bool   `json:"required"`
    Disabled    bool   `json:"disabled"`
    Readonly    bool   `json:"readonly"`
    AutoFocus   bool   `json:"autoFocus"`
    AutoComplete string `json:"autoComplete"`
    
    // Input-specific constraints
    MinLength   int    `json:"minLength"`
    MaxLength   int    `json:"maxLength"`
    Min         string `json:"min"`
    Max         string `json:"max"`
    Step        string `json:"step"`
    Pattern     string `json:"pattern"`
    Accept      string `json:"accept"` // For file inputs
    Multiple    bool   `json:"multiple"` // For file inputs
    
    // State flags for aria-invalid attribute
    Invalid     bool   `json:"invalid"`
    
    // ARIA accessibility attributes
    AriaLabel       string `json:"ariaLabel"`
    AriaDescribedBy string `json:"ariaDescribedBy"`
templ Input(props InputProps) {
```

### icon
```go
type IconProps struct {
	// Core attributes
	Name string `json:"name"`
	
	// Size using shared types (defaults to inherit from context)
	Size components.Size `json:"size,omitempty"`
	
	// Shared component props
	Base components.BaseProps `json:"base,omitempty"`
}

// getIconSize returns width and height based on Size
func getIconSize(size components.Size) (string, string) {
	switch size {
	case components.SizeXs:
		return "12", "12"
	case components.SizeSm:
		return "16", "16"
	case components.SizeLg:
		return "32", "32"
	case components.SizeXl:
		return "40", "40"
	case components.SizeMd, components.SizeDefault, "":
		return "24", "24"
	default:
		return "24", "24"
	}
}

// buildIconAttributes builds templ.Attributes for the SVG element
func buildIconAttributes(props IconProps) templ.Attributes {
templ Icon(props IconProps) {
```

### text
```go
type TextProps struct {
    Text   string
    Size   TextSize
    Weight FontWeight
    Muted  bool
    Class  string
    Attrs  templ.Attributes
}

templ Text(props TextProps) {
    <p
        class={ templ.Classes(
            string(props.Size),
            string(props.Weight),
            templ.KV("text-muted-foreground", props.Muted),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { props.Text }
    </p>
}

// Span variant for inline text
type SpanProps struct {
    Text   string
    Size   TextSize
    Weight FontWeight
    Muted  bool
    Class  string
    Attrs  templ.Attributes
}

templ Span(props SpanProps) {
    <span
        class={ templ.Classes(
            string(props.Size),
            string(props.Weight),
            templ.KV("text-muted-foreground", props.Muted),
            props.Class,
        ) }
        { props.Attrs... }
    >
        { props.Text }
    </span>
}
templ Text(props TextProps) {
templ Span(props SpanProps) {
```

### code
```go
type CodeProps struct {
    Text  string
    Block bool
    Lang  string
    Class string
    Attrs templ.Attributes
}

templ Code(props CodeProps) {
    if props.Block {
        <pre class={ templ.Classes("bg-muted rounded-md p-4 overflow-x-auto", props.Class) } { props.Attrs... }>
            <code class="text-sm font-mono">
                { props.Text }
            </code>
        </pre>
    } else {
        <code
            class={ templ.Classes(
                "relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm",
                props.Class,
            ) }
            { props.Attrs... }
        >
            { props.Text }
        </code>
    }
}
templ Code(props CodeProps) {
```

### button
```go
type ButtonProps struct {
	// Content
	Text string          `json:"text"`
	Icon templ.Component `json:"icon,omitempty"`
	
	// Basecoat variants (using shared types)
	Variant components.ButtonVariant `json:"variant,omitempty"`
	Size    components.Size          `json:"size,omitempty"`
	
	// Icon positioning (for Tailwind spacing utilities)
	IconPosition string `json:"iconPosition,omitempty"` // "start" | "end" (default: "start")
	
	// HTML attributes  
	Type  string `json:"type,omitempty"` // "button", "submit", "reset"
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	
	// Form attributes
	Form       string `json:"form,omitempty"`
	FormAction string `json:"formAction,omitempty"`
	FormMethod string `json:"formMethod,omitempty"`
	FormTarget string `json:"formTarget,omitempty"`
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for Tailwind utilities)
	Base components.BaseProps `json:"base,omitempty"`
}

// getButtonClass returns the correct single Basecoat class
templ Button(props ButtonProps) {
```

### tags
```go
type TagProps struct {
	// Core attributes
	ID   string `json:"id"`
	Text string `json:"text"`
	Value string `json:"value"`
	
	// Basecoat badge variants only
	Variant string `json:"variant,omitempty"` // "primary", "secondary", "destructive", "outline"
	
	// Icon configuration
	Icon      templ.Component `json:"icon,omitempty"`
	IconLeft  templ.Component `json:"iconLeft,omitempty"`
	IconRight templ.Component `json:"iconRight,omitempty"`
	
	// Interactive features
	Selected  bool `json:"selected,omitempty"`
	Removable bool `json:"removable,omitempty"`
	Clickable bool `json:"clickable,omitempty"`
	Disabled  bool `json:"disabled,omitempty"`
	
	// Event handlers (pre-resolved externally)
	OnClick  string `json:"onClick"`
	OnRemove string `json:"onRemove"`
	OnHover  string `json:"onHover"`
	
	// ARIA accessibility attributes
	AriaLabel       string `json:"ariaLabel"`
	AriaDescribedBy string `json:"ariaDescribedBy"`
	AriaPressed     string `json:"ariaPressed"`
	AriaSelected    string `json:"ariaSelected"`
	
templ Tag(props TagProps) {
```

### alert
```go
type AlertProps struct {
	// Core content
	Title       string          `json:"title"`
	Description string          `json:"description,omitempty"`
	Icon        templ.Component `json:"icon,omitempty"`
	
	// Basecoat variants (using shared types)
	Variant components.AlertVariant `json:"variant,omitempty"`
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes accessibility, HTMX, etc.)
	Base components.BaseProps `json:"base,omitempty"`
}

// getAlertClass returns the correct Basecoat class following documentation patterns
func getAlertClass(variant components.AlertVariant) string {
	switch variant {
	case components.AlertDestructive, components.AlertError:
		return "alert-destructive"
	default:
		return "alert"
	}
}

// buildAlertAttributes creates all HTML attributes for the alert element
func buildAlertAttributes(props AlertProps) templ.Attributes {
	attrs := templ.Attributes{
		"class": templ.Classes(
			getAlertClass(props.Variant),
templ Alert(props AlertProps) {
```

### switch
```go
type SwitchProps struct {
	// Core attributes
	ID      string `json:"id"`
	Name    string `json:"name"`
	Value   string `json:"value,omitempty"`
	Checked bool   `json:"checked,omitempty"`
	
	// Content
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	
	// Styling
	Variant SwitchVariant `json:"variant"`
	Size    SwitchSize    `json:"size"`
	
	// Form attributes
	Required bool `json:"required,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	
	// ARIA accessibility
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
	AriaRequired    string `json:"ariaRequired,omitempty"`
	
	// Additional HTML attributes
	TabIndex   int                `json:"tabIndex,omitempty"`
	DataAttrs  map[string]string  `json:"dataAttrs,omitempty"`
	Attributes templ.Attributes   `json:"attributes,omitempty"`
}

// buildSwitchAttributes creates all HTML attributes for the switch input
templ Switch(props SwitchProps) {
```

### kbd
```go
type KbdProps struct {
	// Core content
	Text string          `json:"text"`
	Icon templ.Component `json:"icon,omitempty"`
	
	// HTML attributes
	ID string `json:"id,omitempty"`
	
	// ARIA accessibility
	AriaLabel string `json:"ariaLabel,omitempty"`
}

// Kbd renders a Basecoat kbd atom for keyboard shortcuts
templ Kbd(props KbdProps) {
	<kbd 
		class="kbd"
		if props.ID != "" {
			id={props.ID}
		}
		if props.AriaLabel != "" {
			aria-label={props.AriaLabel}
		}
	>
		if props.Icon != nil {
			@props.Icon
		}
		if props.Text != "" {
			{props.Text}
		}
	</kbd>
}
templ Kbd(props KbdProps) {
```

### slider
```go
type SliderProps struct {
	// Core attributes
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
	
	// Range constraints
	Min  string `json:"min,omitempty"`  // Default: "0"
	Max  string `json:"max,omitempty"`  // Default: "100"
	Step string `json:"step,omitempty"` // Default: "1"
	
	// Form attributes
	Required bool `json:"required,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	
	// Event handlers (pre-resolved externally)
	OnChange string `json:"onChange,omitempty"`
	OnInput  string `json:"onInput,omitempty"`
	OnFocus  string `json:"onFocus,omitempty"`
	OnBlur   string `json:"onBlur,omitempty"`
	
	// ARIA accessibility
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
	AriaValueMin    string `json:"ariaValueMin,omitempty"`
	AriaValueMax    string `json:"ariaValueMax,omitempty"`
	AriaValueNow    string `json:"ariaValueNow,omitempty"`
	AriaValueText   string `json:"ariaValueText,omitempty"`
	
	// Additional HTML attributes
	TabIndex   int               `json:"tabIndex,omitempty"`
templ Slider(props SliderProps) {
```

### radio
```go
type RadioProps struct {
	// Core HTML attributes
	ID      string // The ID of the radio input
	Name    string // The name of the radio input (groups radios together)
	Value   string // The value of the radio input
	Label   string // The label text for the radio
	Checked bool   // Whether the radio is checked

	// Constraints and validation
	Required  bool // Whether the radio is required
	Disabled  bool // Whether the radio is disabled
	Readonly  bool // Whether the radio is readonly
	AutoFocus bool // Whether the radio should autofocus

	// Layout and presentation
	Description string // Additional description text
	Horizontal  bool   // Whether to render label and radio horizontally

	// Event handlers
	OnChange string // JavaScript to execute on change
	OnBlur   string // JavaScript to execute on blur
	OnFocus  string // JavaScript to execute on focus
	OnClick  string // JavaScript to execute on click

	// ARIA accessibility attributes
	AriaLabel       string // Aria label for the radio
	AriaDescribedBy string // ID of element describing the radio
	AriaInvalid     any    // "true", "false", or bool for aria-invalid
	AriaRequired    string // Aria required attribute
	AriaChecked     string // Aria checked attribute

templ Radio(props RadioProps) {
```

### checkbox
```go
type CheckboxProps struct {
    // Core HTML attributes
    Name    string `json:"name"`
    Value   string `json:"value"`
    Checked bool   `json:"checked"`
    
    // Label content - optional for labeled variant
    Label string `json:"label"`
    
    // Description text for complex layouts (optional)
    Description string `json:"description,omitempty"`
    
    // Form state
    Required  bool `json:"required"`
    Disabled  bool `json:"disabled"`
    Readonly  bool `json:"readonly"`
    AutoFocus bool `json:"autofocus"`
    
    // Layout options
    Horizontal bool `json:"horizontal,omitempty"` // Use horizontal layout with description
    
    // Event handlers (pre-resolved externally)
    OnChange string `json:"onChange,omitempty"`
    OnBlur   string `json:"onBlur,omitempty"`
    OnFocus  string `json:"onFocus,omitempty"`
    OnClick  string `json:"onClick,omitempty"`
    
    // Accessibility (using shared AccessibilityProps pattern)
    AriaLabel       string `json:"ariaLabel,omitempty"`
    AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
    AriaInvalid     string `json:"ariaInvalid,omitempty"`
templ Checkbox(props CheckboxProps) {
```

### select
```go
type SelectOption struct {
    Value    string `json:"value"`
    Label    string `json:"label"`
    Selected bool   `json:"selected"`
    Disabled bool   `json:"disabled"`
    Group    string `json:"group"` // For optgroup
}

// SelectProps defines all properties for the Select atom
type SelectProps struct {
	// Core HTML attributes
	ID          string         // The ID of the select element
	Name        string         // The name of the select element
	Value       string         // The current value of the select
	Options     []SelectOption // The list of options
	Placeholder string         // Placeholder text for empty state

	// Constraints and validation
	Required  bool // Whether the select is required
	Disabled  bool // Whether the select is disabled
	Readonly  bool // Whether the select is readonly
	AutoFocus bool // Whether the select should autofocus
	Multiple  bool // Whether multiple selections are allowed
	Size      int  // HTML size attribute (number of visible options)

	// Event handlers
	OnChange string // JavaScript to execute on change
	OnBlur   string // JavaScript to execute on blur
	OnFocus  string // JavaScript to execute on focus
	OnClick  string // JavaScript to execute on click

	// ARIA accessibility attributes
	AriaLabel       string // Aria label for the select
	AriaDescribedBy string // ID of element describing the select
	AriaInvalid     any    // "true", "false", or bool for aria-invalid
	AriaRequired    string // Aria required attribute
	AriaExpanded    string // Aria expanded attribute for custom selects

	// Base component properties
	components.BaseProps
templ Select(props SelectProps) {
```

### date_picker
```go
type DatePickerProps struct {
	// Core attributes
	ID          string `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`        // ISO date format: YYYY-MM-DD
	Placeholder string `json:"placeholder,omitempty"`   // For browsers that don't support date inputs
	
	// Date constraints
	Min string `json:"min,omitempty"`  // ISO date format: YYYY-MM-DD  
	Max string `json:"max,omitempty"`  // ISO date format: YYYY-MM-DD
	
	// Form attributes
	Required bool `json:"required,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Readonly bool `json:"readonly,omitempty"`
	
	// Event handlers (pre-resolved externally)
	OnChange string `json:"onChange,omitempty"`
	OnFocus  string `json:"onFocus,omitempty"`
	OnBlur   string `json:"onBlur,omitempty"`
	
	// ARIA accessibility attributes
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
	AriaRequired    string `json:"ariaRequired,omitempty"`
	AriaInvalid     string `json:"ariaInvalid,omitempty"`
	
	// Additional HTML attributes
	TabIndex   int               `json:"tabIndex,omitempty"`
	DataAttrs  map[string]string `json:"dataAttrs,omitempty"`
	Attributes templ.Attributes  `json:"attributes,omitempty"`
templ DatePicker(props DatePickerProps) {
```

### label
```go
type LabelProps struct {
	// Content
	Text string `json:"text"`
	
	// Core HTML attributes
	For string `json:"for"` // Associates with form control via id
	
	// State indicators
	Required     bool   `json:"required"`     // Shows required indicator
	Optional     bool   `json:"optional"`     // Shows optional indicator  
	RequiredText string `json:"requiredText"` // Custom required text (default "*")
	OptionalText string `json:"optionalText"` // Custom optional text (default "(optional)")
	
	// HTML attributes using shared types
	Base components.BaseProps `json:"base,omitempty"`
	
	// Additional attributes using templ.Attributes for extensibility  
	Attrs templ.Attributes `json:"attrs,omitempty"`
}

// getBaseLabelClasses returns the correct Basecoat label class + Tailwind utilities
// Following documentation pattern from label.md
func getBaseLabelClasses(className string) string {
	// Base Basecoat class as documented
	base := "label"
	
	// Combine with Tailwind utilities from props
	if className != "" {
		return base + " " + className
	}
	
templ Label(props LabelProps) {
```

### textarea
```go
type TextareaProps struct {
    // Core HTML attributes
    ID          string `json:"id"`
    Name        string `json:"name"`
    Value       string `json:"value"`
    Placeholder string `json:"placeholder"`
    
    // Constraints and validation
    Required     bool   `json:"required"`
    Disabled     bool   `json:"disabled"`
    Readonly     bool   `json:"readonly"`
    AutoFocus    bool   `json:"autoFocus"`
    AutoComplete string `json:"autoComplete"`
    
    // Textarea-specific constraints
    MinLength int    `json:"minLength"`
    MaxLength int    `json:"maxLength"`
    Rows      int    `json:"rows"`
    Cols      int    `json:"cols"`
    Wrap      string `json:"wrap"` // "soft" | "hard" | "off"
    
    // State flags for aria-invalid attribute
    Invalid bool `json:"invalid"`
    
    // ARIA accessibility attributes
    AriaLabel       string `json:"ariaLabel"`
    AriaDescribedBy string `json:"ariaDescribedBy"`
    AriaInvalid     string `json:"ariaInvalid"`
    AriaRequired    string `json:"ariaRequired"`
    
    // Base component properties for HTMX, data attributes, etc.
templ Textarea(props TextareaProps) {
```

### autocomplete
```go
type AutocompleteOption struct {
	Value       string          `json:"value"`
	Label       string          `json:"label"`
	Description string          `json:"description,omitempty"`
	Icon        templ.Component `json:"icon,omitempty"`
	Disabled    bool            `json:"disabled,omitempty"`
}

// AutocompleteProps defines all properties for the Autocomplete component using Basecoat select
type AutocompleteProps struct {
	// Core attributes
	ID          string `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	
	// Options and configuration
	Options       []AutocompleteOption `json:"options"`
	SearchURL     string               `json:"searchURL,omitempty"`     // For dynamic loading
	MinChars      int                  `json:"minChars,omitempty"`      // Min chars before search
	EmptyMessage  string               `json:"emptyMessage,omitempty"`  // No results message
	
	// Form attributes
	Required bool `json:"required,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	
	// Event handlers (pre-resolved externally)
	OnChange string `json:"onChange,omitempty"`
	OnSelect string `json:"onSelect,omitempty"`
	OnFocus  string `json:"onFocus,omitempty"`
	OnBlur   string `json:"onBlur,omitempty"`
	
	// ARIA accessibility attributes
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
	AriaRequired    string `json:"ariaRequired,omitempty"`
	AriaInvalid     string `json:"ariaInvalid,omitempty"`
	
	// Additional HTML attributes
	TabIndex   int               `json:"tabIndex,omitempty"`
templ Autocomplete(props AutocompleteProps) {
```

### spinner
```go
type SpinnerProps struct {
	// Visual properties using shared Size types
	Size components.Size `json:"size,omitempty"`
	
	// Shared component props
	Base components.BaseProps `json:"base,omitempty"`
}

// getSpinnerSize returns width and height for spinner based on Size
func getSpinnerSize(size components.Size) (string, string) {
	switch size {
	case components.SizeXs:
		return "12", "12"
	case components.SizeSm:
		return "16", "16"
	case components.SizeLg:
		return "32", "32"
	case components.SizeXl:
		return "40", "40"
	case components.SizeMd, components.SizeDefault, "":
		return "24", "24"
	default:
		return "24", "24"
	}
}

// buildSpinnerAttributes builds templ.Attributes for the SVG spinner element
func buildSpinnerAttributes(props SpinnerProps) templ.Attributes {
	width, height := getSpinnerSize(props.Size)
	
	attrs := templ.Attributes{
templ Spinner(props SpinnerProps) {
```

### avatar
```go
type AvatarProps struct {
	// Image properties
	Src      string `json:"src"`
	Alt      string `json:"alt"`
	Initials string `json:"initials,omitempty"` // Fallback text when no image
	
	// Visual properties (using shared types)
	Size  components.Size `json:"size,omitempty"`  // Avatar size variant
	Shape AvatarShape     `json:"shape,omitempty"` // Avatar shape variant
	
	// Status indicator
	ShowStatus  bool              `json:"showStatus,omitempty"`  // Show status indicator
	StatusColor AvatarStatusColor `json:"statusColor,omitempty"` // Status color variant
	
	// Ring/border
	HasRing    bool            `json:"hasRing,omitempty"`    // Add ring border
	RingColor  AvatarRingColor `json:"ringColor,omitempty"`  // Ring color variant
	RingOffset bool            `json:"ringOffset,omitempty"` // Add ring offset
	
	// Interactive
	Clickable bool `json:"clickable,omitempty"` // Add hover effects
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ID, data attributes)
	Base components.BaseProps `json:"base,omitempty"`
}

// getAvatarSizeClass returns Tailwind size classes based on shared Size type
func getAvatarSizeClass(size components.Size) string {
templ Avatar(props AvatarProps) {
```

### progress
```go
type ProgressProps struct {
	// Core properties
	ID    string `json:"id,omitempty"`
	Value int    `json:"value"`     // Current progress value (0-100)
	Max   int    `json:"max,omitempty"`       // Maximum value (default: 100)
	
	// Visual properties
	Size    ProgressSize    `json:"size,omitempty"`    // "sm", "md", "lg", "xl"
	Variant ProgressVariant `json:"variant,omitempty"` // "default", "success", "warning", "error"
	
	// Accessibility
	AriaLabel       string `json:"ariaLabel,omitempty"`
	AriaLabelledBy  string `json:"ariaLabelledBy,omitempty"`
	AriaDescribedBy string `json:"ariaDescribedBy,omitempty"`
	
	// Additional attributes (HTMX, Alpine.js, etc.)
	Attrs templ.Attributes `json:"attrs,omitempty"`
}

// ProgressSize defines type-safe size variants
type ProgressSize string

const (
	ProgressSM ProgressSize = "sm"
	ProgressMD ProgressSize = "md"
	ProgressLG ProgressSize = "lg"
	ProgressXL ProgressSize = "xl"
)

// ProgressVariant defines type-safe color variants
type ProgressVariant string
templ Progress(props ProgressProps) {
```

### skeleton
```go
type SkeletonProps struct {
	// Core properties
	ID string `json:"id,omitempty"`
	
	// Shape presets for common patterns
	Shape string `json:"shape,omitempty"` // "text", "heading", "avatar", "button", "card", "image"
	
	// Size variants for predefined shapes
	Size string `json:"size,omitempty"` // "xs", "sm", "md", "lg", "xl"
	
	// Custom dimensions (overrides shape/size)
	Width  string `json:"width,omitempty"`
	Height string `json:"height,omitempty"`
	
	// Style options
	Circle bool   `json:"circle,omitempty"` // Makes avatar/image circular
	Muted  bool   `json:"muted,omitempty"`  // Uses bg-muted instead of bg-accent
	Animate bool  `json:"animate,omitempty"` // Controls animate-pulse (default: true)
	
	// Additional attributes for HTMX, Alpine.js, etc
	Attrs templ.Attributes `json:"attrs,omitempty"`
}

// getSkeletonClasses builds Tailwind classes based on shape and size
func getSkeletonClasses(props SkeletonProps) string {
	classes := []string{}
	
	// Base animation and background
	if props.Animate != false { // default true
		classes = append(classes, "animate-pulse")
		// Support reduced motion
templ Skeleton(props SkeletonProps) {
```

### link
```go
type LinkProps struct {
	// Content
	Text string          `json:"text"`
	Icon templ.Component `json:"icon,omitempty"`
	Href string          `json:"href"`
	
	// Basecoat variants (using shared types)
	Variant LinkVariant `json:"variant,omitempty"`
	
	// Icon positioning (for Tailwind spacing utilities)
	IconPosition string `json:"iconPosition,omitempty"` // "start" | "end" (default: "start")
	
	// Link behavior
	Target    string `json:"target,omitempty"`    // "_blank", "_self", "_parent", "_top"
	Rel       string `json:"rel,omitempty"`       // "noopener", "noreferrer", etc.
	Download  string `json:"download,omitempty"`  // Download attribute
	External  bool   `json:"external,omitempty"`  // Add external link styling
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for Tailwind utilities)
	Base components.BaseProps `json:"base,omitempty"`
}

// LinkVariant represents link-specific style variations
type LinkVariant string

const (
	LinkDefault LinkVariant = "default"
	LinkMuted   LinkVariant = "muted"
templ Link(props LinkProps) {
```

### image
```go
type ImageProps struct {
	// Core attributes
	Src string `json:"src"`
	Alt string `json:"alt"`
	
	// Basecoat variants (using shared types)
	Variant ImageVariant    `json:"variant,omitempty"`
	Size    components.Size `json:"size,omitempty"`
	
	// Image-specific attributes
	Width     string `json:"width,omitempty"`
	Height    string `json:"height,omitempty"`
	Sizes     string `json:"sizes,omitempty"`     // Responsive sizes
	SrcSet    string `json:"srcSet,omitempty"`    // Responsive sources
	Loading   string `json:"loading,omitempty"`   // "lazy" | "eager" 
	Decoding  string `json:"decoding,omitempty"`  // "async" | "sync" | "auto"
	FetchPrio string `json:"fetchPrio,omitempty"` // "high" | "low" | "auto"
	
	// Fallback handling
	FallbackSrc string `json:"fallbackSrc,omitempty"` // Fallback image URL
	Placeholder string `json:"placeholder,omitempty"`  // Placeholder content/base64
	
	// Layout properties  
	ObjectFit    string `json:"objectFit,omitempty"`    // CSS object-fit value
	ObjectPos    string `json:"objectPos,omitempty"`    // CSS object-position value
	AspectRatio  string `json:"aspectRatio,omitempty"`  // CSS aspect-ratio value
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for Tailwind utilities)
templ Image(props ImageProps) {
```

### heading
```go
type HeadingProps struct {
	// Content
	Level int    `json:"level"`           // 1-6 for h1-h6
	Text  string `json:"text"`            // Heading text content
	
	// Visual styling
	Variant HeadingVariant `json:"variant,omitempty"` // Color variant
	Weight  HeadingWeight  `json:"weight,omitempty"`  // Font weight
	Align   HeadingAlign   `json:"align,omitempty"`   // Text alignment
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for additional styling)
	Base components.BaseProps `json:"base,omitempty"`
}

// getHeadingBaseClasses returns the base classes for heading styling
// Since Basecoat doesn't provide heading classes, we use Tailwind utilities
func getHeadingBaseClasses(level int) string {
	switch level {
	case 1:
		return "text-4xl leading-tight" // h1: 36px, tight line height
	case 2:
		return "text-3xl leading-tight" // h2: 30px
	case 3:
		return "text-2xl leading-snug"  // h3: 24px
	case 4:
		return "text-xl leading-snug"   // h4: 20px
	case 5:
		return "text-lg leading-normal" // h5: 18px
templ Heading(props HeadingProps) {
```

## Molecules

### errorhandling
```go
type ErrorDetail struct {
	Code        string                 `json:"code"`
	Message     string                 `json:"message"`
	Field       string                 `json:"field"`
	Type        ErrorType              `json:"type"`
	Severity    ErrorSeverity          `json:"severity"`
	Timestamp   string                 `json:"timestamp"`
	Context     map[string]any `json:"context"`
	Suggestions []string               `json:"suggestions"`
	HelpURL     string                 `json:"helpUrl"`
	Retryable   bool                   `json:"retryable"`
}

// ErrorBoundaryProps defines properties for error boundaries
type ErrorBoundaryProps struct {
	Errors       []ErrorDetail `json:"errors"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	ShowDetails  bool          `json:"showDetails"`
	ShowStack    bool          `json:"showStack"`
	Recoverable  bool          `json:"recoverable"`
	AutoRetry    bool          `json:"autoRetry"`
	RetryCount   int           `json:"retryCount"`
	MaxRetries   int           `json:"maxRetries"`
	RetryDelay   int           `json:"retryDelay"` // milliseconds
	Dismissible  bool          `json:"dismissible"`
	Persistent   bool          `json:"persistent"` // Don't auto-dismiss
	Class        string        `json:"class"`
	ID           string        `json:"id"`
	// Actions
	Actions      []ErrorAction `json:"actions"`
	// Events
	OnRetry      string        `json:"onRetry"`
	OnDismiss    string        `json:"onDismiss"`
	OnReport     string        `json:"onReport"`
	// HTMX
	HXRetry      string        `json:"hxRetry"`
	HXReport     string        `json:"hxReport"`
}

// ErrorAction represents an action button for error handling
type ErrorAction struct {
	Text        string             `json:"text"`
	Icon        string             `json:"icon"`
	Variant     string              `json:"variant"`
	Action      string             `json:"action"` // "retry", "dismiss", "report", "custom"
	URL         string             `json:"url"`
	OnClick     string             `json:"onClick"`
}

templ ErrorBoundary(props ErrorBoundaryProps) {
templ Notification(props NotificationProps) {
templ NotificationContainer(position string) {
templ GlobalErrorHandler() {
```

### card
```go
type CardProps struct {
	// Core content
	Title       string `json:"title"`
	Description string `json:"description"`
	
	// Slot-based composition
	Header   templ.Component `json:"header"`   // Custom header content
	Content  templ.Component `json:"content"`  // Main content area
	Footer   templ.Component `json:"footer"`   // Footer content
	
	// Action slot for header
	HeaderAction templ.Component `json:"headerAction"`
	
	// Accessibility
	ID        string `json:"id"`
	AriaLabel string `json:"ariaLabel"`
	
	// Semantic variants (replaces ClassName)
	Border   bool `json:"border"`   // Add border styling
	Compact  bool `json:"compact"`  // Compact spacing
	Elevated bool `json:"elevated"` // Elevated appearance
}

// Card renders a Basecoat card component with header, content, and footer
templ Card(props CardProps) {
	<div 
		class="card"
		if props.Border {
			data-variant="border"
		}
		if props.Compact {
templ Card(props CardProps) {
templ CardWithBorder(props CardProps) {
templ SimpleCard(title string, content templ.Component) {
templ ActionCard(title, description string, content templ.Component, actions []atoms.ButtonProps) {
templ ImageCard(props CardProps, imageSrc, imageAlt string) {
templ StatsCard(title, value, description string, icon templ.Component, trend string, trendPositive bool) {
```

### menuitem
```go
type MenuItemProps struct {
	// Basic properties
	Type        MenuItemType `json:"type"`
	Size        MenuItemSize `json:"size,omitempty"`
	Text        string       `json:"text,omitempty"`
	Description string       `json:"description,omitempty"`
	Value       string       `json:"value,omitempty"`
	URL         string       `json:"url,omitempty"`
	Target      string       `json:"target,omitempty"` // For links: _blank, _self, etc.
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`

	// Visual elements
	Icon         string                `json:"icon,omitempty"`
	IconLeft     string                `json:"iconLeft,omitempty"`
	IconRight    string                `json:"iconRight,omitempty"`
	Badge        string                `json:"badge,omitempty"`
	BadgeVariant string               `json:"badgeVariant,omitempty"`

	// States
	Active      bool `json:"active,omitempty"`
	Selected    bool `json:"selected,omitempty"`
	Checked     bool `json:"checked,omitempty"`
	Disabled    bool `json:"disabled,omitempty"`
	Destructive bool `json:"destructive,omitempty"`

	// Submenu items
	SubItems []MenuItemProps `json:"subItems,omitempty"`

	// HTMX attributes
	HXPost    string `json:"hxPost,omitempty"`
templ MenuItem(props MenuItemProps) {
templ MenuDivider() {
templ MenuLink(props MenuItemProps) {
templ MenuButton(props MenuItemProps) {
templ MenuCheckbox(props MenuItemProps) {
templ MenuRadio(props MenuItemProps) {
templ MenuSubmenu(props MenuItemProps) {
```

### searchbox
```go
type SearchSuggestion struct {
	Value       string `json:"value"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Category    string `json:"category,omitempty"`
	URL         string `json:"url,omitempty"` // For navigation suggestions
}

// SearchBoxProps defines the properties for the SearchBox component
type SearchBoxProps struct {
	// Basic properties
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Value       string        `json:"value,omitempty"`
	Placeholder string        `json:"placeholder,omitempty"`
	Size        SearchBoxSize `json:"size,omitempty"`
	Disabled    bool          `json:"disabled,omitempty"`
	
	// Search behavior
	Suggestions   []SearchSuggestion `json:"suggestions,omitempty"`
	Loading       bool               `json:"loading,omitempty"`
	MinChars      int                `json:"minChars,omitempty"`      // Minimum characters before showing suggestions
	Debounce      int                `json:"debounce,omitempty"`      // Debounce delay in milliseconds
	ClearOnSelect bool               `json:"clearOnSelect,omitempty"` // Clear input after selection
	
	// HTMX attributes
	HXPost    string `json:"hxPost,omitempty"`    // URL for search API
	HXGet     string `json:"hxGet,omitempty"`     // URL for search API
	HXTarget  string `json:"hxTarget,omitempty"`  // Target for search results
	HXSwap    string `json:"hxSwap,omitempty"`    // Swap strategy
	HXTrigger string `json:"hxTrigger,omitempty"` // Custom trigger
	
	// Alpine.js attributes
	AlpineData   string `json:"alpineData,omitempty"`   // Custom x-data
	AlpineModel  string `json:"alpineModel,omitempty"`  // x-model for input
	AlpineSearch string `json:"alpineSearch,omitempty"` // Custom search method
	AlpineSelect string `json:"alpineSelect,omitempty"` // Method called when suggestion is selected
	AlpineClear  string `json:"alpineClear,omitempty"`  // Method called when clear button is clicked
	
	// Events
templ SearchBox(props SearchBoxProps) {
```

### validation
```go
type ValidationMessage struct {
	Type    ValidationMessageType `json:"type"`
	Message string               `json:"message"`
	Field   string               `json:"field,omitempty"`
	Code    string               `json:"code,omitempty"`
}

// ValidationProps defines properties for validation components using Basecoat patterns
type ValidationProps struct {
	// Core properties
	FieldName    string              `json:"fieldName"`
	Messages     []ValidationMessage `json:"messages"`
	
	// Display options
	ShowSuccess  bool `json:"showSuccess,omitempty"`
	ShowWarnings bool `json:"showWarnings,omitempty"`
	
	// HTML attributes
	ID string `json:"id,omitempty"`
	
	// HTMX integration
	HXPost    string `json:"hxPost,omitempty"`
	HXGet     string `json:"hxGet,omitempty"`
	HXTarget  string `json:"hxTarget,omitempty"`
	HXSwap    string `json:"hxSwap,omitempty"`
	HXTrigger string `json:"hxTrigger,omitempty"`
	
	// Alpine.js integration  
	AlpineData string `json:"alpineData,omitempty"`
	
	// Event handlers
	OnValidate string `json:"onValidate,omitempty"`
	OnError    string `json:"onError,omitempty"`
	OnSuccess  string `json:"onSuccess,omitempty"`
}

// getValidationIcon maps message types to icons
func getValidationIcon(messageType ValidationMessageType) atoms.IconProps {
	switch messageType {
--
type FieldValidationProps struct {
	FieldName   string              `json:"fieldName"`
	Messages    []ValidationMessage `json:"messages"`
	ShowIcon    bool                `json:"showIcon,omitempty"`
	Inline      bool                `json:"inline,omitempty"`
	ID          string              `json:"id,omitempty"`
	AlpineData  string              `json:"alpineData,omitempty"`
}

// FieldValidation renders field-level validation in a more compact form
templ ValidationMessages(props ValidationProps) {
templ FieldValidation(props FieldValidationProps) {
templ ValidationSummary(props ValidationSummaryProps) {
```

### dropdown_menu
```go
type DropdownMenuItem struct {
	// Core properties
	Text        string       `json:"text"`
	Icon        string       `json:"icon"`
	Type        DropdownMenuItemType `json:"type"`
	
	// Actions
	OnClick  string `json:"onClick"`
	HxPost   string `json:"hxPost"`
	HxGet    string `json:"hxGet"`
	HxTarget string `json:"hxTarget"`
	HxSwap   string `json:"hxSwap"`
	
	// States
	Disabled bool `json:"disabled"`
	Checked  bool `json:"checked"`  // For checkbox/radio items
	
	// Keyboard shortcut
	Shortcut string `json:"shortcut"`
	
	// ARIA
	AriaLabel string `json:"ariaLabel"`
}

// DropdownMenuGroup represents a group of dropdown menu items
type DropdownMenuGroup struct {
	Heading string             `json:"heading"`
	Items   []DropdownMenuItem `json:"items"`
}

// DropdownMenuProps defines properties for the DropdownMenu molecule
type DropdownMenuProps struct {
	// Core properties
	ID      string `json:"id"`
	Trigger string `json:"trigger"`
	
	// Trigger customization
	TriggerComponent templ.Component `json:"triggerComponent"` // Custom trigger
	TriggerButton    atoms.ButtonProps `json:"triggerButton"` // Button props for trigger
	
	// Content
	Items  []DropdownMenuItem  `json:"items"`
	Groups []DropdownMenuGroup `json:"groups"`
	
	// Popover positioning
	Side  string `json:"side"`  // "top", "bottom", "left", "right"
	Align string `json:"align"` // "start", "center", "end"
	
	// Styling - semantic width options
	WidthSize  string `json:"widthSize"` // "sm", "md", "lg", "xl" or custom
templ DropdownMenu(props DropdownMenuProps) {
templ SimpleDropdownMenu(id, trigger string, items []DropdownMenuItem) {
templ UserDropdownMenu(id, userName string) {
templ ActionsDropdownMenu(id string, actions []DropdownMenuItem) {
```

### form_field
```go
type FormFieldOption struct {
    Value    string
    Label    string
    Selected bool
    Disabled bool
}

// FormFieldProps defines all properties for FormField molecule
type FormFieldProps struct {
    // Core properties
    ID          string
    Name        string
    Type        FormFieldType
    Label       string
    Value       string
    Placeholder string
    HelpText    string
    
    // Options for select/radio/checkbox
    Options     []FormFieldOption
    
    // State (pre-computed externally)
    Required    bool
    Disabled    bool
    Readonly    bool
    HasError    bool
    Horizontal  bool  // Use horizontal layout
    
    // Validation (pre-computed externally) 
    Errors      []string
    
    // Event handlers (pre-resolved HTML attributes)
    OnChange    string
    OnBlur      string
    OnFocus     string
    OnInput     string
}

// buildEventHandlerAttrs creates a templ.Attributes map with event handlers
templ FormField(props FormFieldProps) {
```

### button_group
```go
type ButtonGroupProps struct {
	// Content - array of buttons to render
	Buttons []ButtonConfig `json:"buttons"`
	
	// Layout behavior
	Orientation components.LayoutOrientation `json:"orientation,omitempty"` // "horizontal" (default) | "vertical"
	Size        components.Size              `json:"size,omitempty"`        // affects all buttons in group
	
	// Selection behavior (for segmented control style)
	SelectionMode string `json:"selectionMode,omitempty"` // "" | "single" | "multiple"
	Selected      []int  `json:"selected,omitempty"`      // indices of selected buttons
	
	// Accessibility
	AriaLabel       string `json:"ariaLabel,omitempty"`       // describes the group purpose
	AriaLabelledBy  string `json:"ariaLabelledBy,omitempty"`  // references label element
	Role            string `json:"role,omitempty"`            // "group" (default) | "radiogroup" | "toolbar"
	
	// Additional attributes using templ.Attributes for extensibility
	Attrs templ.Attributes `json:"attrs,omitempty"`
	
	// Shared component props (includes ClassName for additional Tailwind utilities)
	Base components.BaseProps `json:"base,omitempty"`
}

// ButtonConfig represents individual button configuration within the group
type ButtonConfig struct {
	Text         string                   `json:"text"`
	Icon         templ.Component          `json:"icon,omitempty"`
	IconPosition string                   `json:"iconPosition,omitempty"` // "start" | "end"
	Variant      components.ButtonVariant `json:"variant,omitempty"`      // if empty, inherits from group logic
	Disabled     bool                     `json:"disabled,omitempty"`
	
	// Button-specific attributes
	Type  string `json:"type,omitempty"`  // "button", "submit", "reset"
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	
	// HTMX for individual buttons
	HTMX components.HTMXConfig `json:"htmx,omitempty"`
	
	// Accessibility for individual buttons
	AriaLabel    string `json:"ariaLabel,omitempty"`    // required for icon-only buttons
	AriaPressed  string `json:"ariaPressed,omitempty"`  // for toggle buttons
	AriaCurrent  string `json:"ariaCurrent,omitempty"`  // for pagination/navigation
	AriaExpanded string `json:"ariaExpanded,omitempty"` // for dropdown triggers
	AriaHaspopup string `json:"ariaHaspopup,omitempty"` // for dropdown triggers
	
	// Event handlers  
	OnClick templ.ComponentScript `json:"onClick,omitempty"`
	
templ ButtonGroup(props ButtonGroupProps) {
```

## Organisms

### navigation
```go
type NavigationItem struct {
	// Core properties
	ID          string `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Icon        string `json:"icon"`
	Badge       string `json:"badge"`
	
	// States
	Active   bool `json:"active"`
	Disabled bool `json:"disabled"`
	
	// Permissions
	Permissions []string `json:"permissions"`
	Roles       []string `json:"roles"`
	Condition   string   `json:"condition"`
	
	// Event handlers
	OnClick    string `json:"onClick"`
	HXGet      string `json:"hxGet"`
	HXPost     string `json:"hxPost"`
	HXTarget   string `json:"hxTarget"`
	HXSwap     string `json:"hxSwap"`
	HXTrigger  string `json:"hxTrigger"`
	AlpineClick string `json:"alpineClick"`
	
	// Nested items
	Items []NavigationItem `json:"items"`
}

--
type NavigationSearch struct {
	Enabled     bool   `json:"enabled"`
	Placeholder string `json:"placeholder"`
	MinLength   int    `json:"minLength"`
	Debounce    int    `json:"debounce"`
	Query       string `json:"query"`
	Loading     bool   `json:"loading"`
	Visible     bool   `json:"visible"`
}

// NavigationMobileMenu defines mobile menu configuration
type NavigationMobileMenu struct {
	Enabled     bool   `json:"enabled"`
	Position    string `json:"position"`
	Overlay     bool   `json:"overlay"`
	CloseButton bool   `json:"closeButton"`
	Swipeable   bool   `json:"swipeable"`
}
templ Navigation(props NavigationProps) {
```

### form
```go
type FormProps struct {
	// CORE - Required for all forms
	ID     string  `json:"id"`
	Fields []Field `json:"fields"`

	// BASIC - Optional enhancements
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Layout      FormLayout `json:"layout,omitempty"`        // Default: vertical
	Size        FormSize   `json:"size,omitempty"`          // Default: md
	ReadOnly    bool       `json:"readOnly,omitempty"`

	// SECTIONING - Group related fields
	Sections []Section `json:"sections,omitempty"`

	// SUBMISSION - Basic form handling
	SubmitURL string `json:"submitURL,omitempty"`
	OnSubmit  string `json:"onSubmit,omitempty"`

	// ACTIONS - Custom buttons (uses atoms.Button)
	Actions []Action `json:"actions,omitempty"`

	// PROGRESSIVE ENHANCEMENT - Nil = disabled
	Advanced    *AdvancedConfig    `json:"advanced,omitempty"`
	Validation  *ValidationConfig  `json:"validation,omitempty"`
	AutoSave    *AutoSaveConfig    `json:"autoSave,omitempty"`
	Storage     *FormStorageConfig     `json:"storage,omitempty"`
	Progress    *FormProgressConfig    `json:"progress,omitempty"`
	Dependencies *DependencyConfig `json:"dependencies,omitempty"`
	Debug       *DebugConfig       `json:"debug,omitempty"`

--
type Field struct {
	molecules.FormFieldProps        // Reuse existing molecule

	// PROGRESSIVE ENHANCEMENT
	Conditional  string       `json:"conditional,omitempty"`  // Alpine.js visibility
	Dependencies []Dependency `json:"dependencies,omitempty"` // Field relationships

	// FIELD-LEVEL OVERRIDES
	AutoSave   *FieldAutoSave   `json:"autoSave,omitempty"`
	Validation *FieldValidation `json:"validation,omitempty"`
	Storage    *FieldStorage    `json:"storage,omitempty"`
}

// Section - Reuse existing pattern
type Section struct {
	ID          string     `json:"id,omitempty"`
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
templ Form(props FormProps) {
```

### datatable
```go
type DataTableColumn struct {
	// Core properties
	Key           string `json:"key"`
	Title         string `json:"title"`
	Type          ColumnType `json:"type"`
	Width         string `json:"width"`
	MinWidth      string `json:"minWidth"`
	MaxWidth      string `json:"maxWidth"`
	
	// Features
	Sortable      bool `json:"sortable"`
	Searchable    bool `json:"searchable"`
	Filterable    bool `json:"filterable"`
	Resizable     bool `json:"resizable"`
	Visible       bool `json:"visible"`
	Clickable     bool `json:"clickable"`
	
	// Appearance
	Align         string `json:"align"`
	Format        string `json:"format"`
	
	// Type-specific configurations
	BadgeMap      map[string]string `json:"badgeMap"`
	CurrencyCode  string `json:"currencyCode"`
	Precision     int    `json:"precision"`
	DateFormat    string `json:"dateFormat"`
	
	// Actions for this column
	ActionItems   []molecules.MenuItemProps `json:"actionItems"`
	OnClick       string `json:"onClick"`
}
--
type DataTableRow struct {
	// Core data
	ID       string         `json:"id"`
	Data     map[string]any `json:"data"`
	
	// State
	Selected bool `json:"selected"`
	Expanded bool `json:"expanded"`
	Disabled bool `json:"disabled"`
	
	// Appearance
	Class    string `json:"class"`
	
	// Row-specific actions
	Actions  []molecules.MenuItemProps `json:"actions"`
	
	// Additional metadata
	Meta     map[string]any `json:"meta"`
templ DataTable(props DataTableProps) {
```

### tabs
```go
type TabItem struct {
	// Core properties
	ID      string `json:"id"`
	Label   string `json:"label"`
	Icon    string `json:"icon"`
	Content templ.Component `json:"content"`
	
	// States
	Active   bool `json:"active"`
	Disabled bool `json:"disabled"`
	
	// Badge/indicator
	Badge string `json:"badge"`
	Count int    `json:"count"`
	
	// ARIA
	AriaLabel string `json:"ariaLabel"`
}

// TabsProps defines properties for the Tabs organism
type TabsProps struct {
	// Core properties
	ID    string    `json:"id"`
	Items []TabItem `json:"items"`
	
	// Active tab (by ID or index)
	ActiveTab     string `json:"activeTab"`     // Tab ID
	ActiveIndex   int    `json:"activeIndex"`   // Tab index (0-based)
	
	// Layout and styling
	Orientation string `json:"orientation"` // "horizontal" | "vertical"
	FullWidth   bool   `json:"fullWidth"`   // Full width tab list
	
	// Event handlers
	OnTabChange string `json:"onTabChange"` // JavaScript function name
	
	// HTMX integration
	HxGet    string `json:"hxGet"`    // Endpoint for dynamic content loading
	HxTarget string `json:"hxTarget"` // Target for HTMX updates
	HxSwap   string `json:"hxSwap"`   // How to swap content
	
	// ARIA
	AriaLabel string `json:"ariaLabel"`
}

// Tabs renders a Basecoat tabs component
templ Tabs(props TabsProps) {
	<div 
		class="tabs"
		id={ props.ID }
templ Tabs(props TabsProps) {
templ SimpleTabs(id string, items []TabItem) {
templ FullWidthTabs(id string, items []TabItem) {
templ VerticalTabs(id string, items []TabItem) {
templ DynamicTabs(id string, items []TabItem, endpoint string) {
templ CardTabs(id string, items []TabItem, title, description string) {
templ NavStyleTabs(id string, items []TabItem) {
```

## Templates

### pagelayout
```go
type BreadcrumbItem struct {
	Text string
	URL  string
	Icon string
}

// PageMeta represents page metadata
type PageMeta struct {
	Title       string
	Description string
	Keywords    []string
	Author      string
	Favicon     string
	Canonical   string
}

// PageHeader represents page header content
type PageHeader struct {
	Title       string
	Description string
	Icon        string
	Badge       string
	BadgeVariant atoms.BadgeVariant
	Actions     []organisms.TableAction
	Breadcrumbs []BreadcrumbItem
	Tabs        []organisms.NavigationItem
	Class       string
}

// PageSidebar represents sidebar configuration
type PageSidebar struct {
	Title       string
	Logo        string
	LogoURL     string
	Width       string
	Navigation  []organisms.NavigationItem
	Collapsible bool
	Collapsed   bool
	Footer      bool
	Class       string
}

// PageTopbar represents topbar configuration
type PageTopbar struct {
	Title        string
	Logo         string
	LogoURL      string
	Navigation   []organisms.NavigationItem
	UserMenu     []organisms.NavigationItem
	UserName     string
templ PageLayout(props PageLayoutProps, children ...templ.Component) {
templ SidebarLayout(props PageLayoutProps, children ...templ.Component) {
templ TopbarLayout(props PageLayoutProps, children ...templ.Component) {
templ CombinedLayout(props PageLayoutProps, children ...templ.Component) {
templ FullscreenLayout(props PageLayoutProps, children ...templ.Component) {
```

### base_layout
```go
type PageMeta struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Author      string   `json:"author"`
	Canonical   string   `json:"canonical"`
	Favicon     string   `json:"favicon"`
	CSRFToken   string   `json:"csrfToken"`
	OGImage     string   `json:"ogImage"`
	OGType      string   `json:"ogType"`
}

// ThemeConfig defines theme configuration
type ThemeConfig struct {
	ID         string `json:"id"`
	DarkMode   bool   `json:"darkMode"`
	SystemSync bool   `json:"systemSync"` // Sync with system preferences
}

// HTMXConfig defines HTMX configuration
type HTMXConfig struct {
	Boost   bool   `json:"boost"`
	Target  string `json:"target"`
	Swap    string `json:"swap"`
	PushURL bool   `json:"pushURL"`
}

// BaseLayoutProps defines properties for the base layout template
type BaseLayoutProps struct {
	// Meta information
	Meta PageMeta `json:"meta"`

	// Theme configuration
	Theme ThemeConfig `json:"theme"`

	// Core properties
	ID         string            `json:"id"`
	Class      string            `json:"class"`
	Attributes templ.Attributes  `json:"attributes"`

	// HTMX configuration
	HTMX HTMXConfig `json:"htmx"`

	// Alpine.js configuration
	AlpineData string `json:"alpineData"`

	// Resources
	CustomCSS []string `json:"customCSS"`
	CustomJS  []string `json:"customJS"`

templ BaseLayout(props BaseLayoutProps, children ...templ.Component) {
templ BaseHead(props BaseLayoutProps) {
templ BaseScripts(props BaseLayoutProps) {
templ PageLoader(customLoader templ.Component) {
templ DefaultLoader() {
```

### layout_adapters
```go
templ PageLayoutAdapter(props PageLayoutProps, children ...templ.Component) {
templ DashboardLayoutAdapter(props DashboardLayoutProps, children ...templ.Component) {
```

### layout_examples
```go
templ ExampleSidebarLayout() {
templ ExampleAnalyticsDashboard() {
templ ExampleTopbarLayout() {
templ ExampleFullscreenLayout() {
```

### dashboard_layout
```go
type DashboardLayoutProps struct {
	// Extends enhanced page layout
	EnhancedPageLayoutProps
	
	// Dashboard configuration
	Title           string                       `json:"title"`
	Description     string                       `json:"description"`
	Section         DashboardSection             `json:"section"`
	
	// Dashboard components
	Metrics         []DashboardMetric            `json:"metrics"`
	Charts          []DashboardChart             `json:"charts"`
	Tables          []DashboardTable             `json:"tables"`
	Widgets         []DashboardWidget            `json:"widgets"`
	Stats           []DashboardStats             `json:"stats"`
	QuickActions    []DashboardQuickAction       `json:"quickActions"`
	
	// Layout configuration
	GridConfig      DashboardGridConfig          `json:"gridConfig"`
	
	// Features
	EnableRefresh   bool                         `json:"enableRefresh"`
	RefreshInterval int                          `json:"refreshInterval"` // seconds
	EnableExport    bool                         `json:"enableExport"`
	EnableFilters   bool                         `json:"enableFilters"`
	ShowStats       bool                         `json:"showStats"`
	ShowActions     bool                         `json:"showActions"`
	
	// Data configuration
	DataSource      string                       `json:"dataSource"`
	RealTimeData    bool                         `json:"realTimeData"`
--
type DashboardMetric struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Value       string                 `json:"value"`
	Unit        string                 `json:"unit"`
	Change      float64                `json:"change"`
	ChangeType  string                 `json:"changeType"` // "increase", "decrease", "neutral"
	Icon        string                 `json:"icon"`
	Color       string                 `json:"color"`
	Sparkline   []float64              `json:"sparkline"`
	Loading     bool                   `json:"loading"`
	HXGet       string                 `json:"hxGet"`
	HXTrigger   string                 `json:"hxTrigger"`
}

// DashboardChart represents a chart widget
type DashboardChart struct {
	ID          string                 `json:"id"`
templ DashboardLayout(props DashboardLayoutProps, children ...templ.Component) {
templ AnalyticsDashboard(props DashboardLayoutProps, children ...templ.Component) {
templ OverviewDashboard(props DashboardLayoutProps, children ...templ.Component) {
```

