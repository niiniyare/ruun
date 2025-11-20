package theme

import (
	"fmt"
	"regexp"
	"strings"
)

type TokenReference string

func (tr TokenReference) String() string {
	return string(tr)
}

func (tr TokenReference) IsReference() bool {
	value := strings.TrimSpace(string(tr))
	
	if value == "" || strings.Contains(value, " ") {
		return false
	}
	
	if !strings.Contains(value, ".") {
		return false
	}
	
	if tr.isCSSLiteral(value) {
		return false
	}
	
	segments := strings.Split(value, ".")
	if len(segments) < 2 {
		return false
	}
	
	for _, segment := range segments {
		if segment == "" {
			return false
		}
	}
	
	return true
}

func (tr TokenReference) isCSSLiteral(value string) bool {
	if strings.HasPrefix(value, "#") {
		return true
	}
	
	cssFunctions := []string{
		"calc(", "var(", "url(", "rgb(", "rgba(", "hsl(", "hsla(",
		"linear-gradient(", "radial-gradient(", "conic-gradient(",
	}
	for _, fn := range cssFunctions {
		if strings.HasPrefix(value, fn) {
			return true
		}
	}
	
	if matched, _ := regexp.MatchString(`^-?\d+\.\d+[a-zA-Z%]*$`, value); matched {
		return false
	}
	
	return false
}

func (tr TokenReference) Validate() error {
	value := strings.TrimSpace(string(tr))
	
	if value == "" {
		return NewError(ErrCodeValidation, "token reference cannot be empty")
	}
	
	if !tr.IsReference() {
		return nil
	}
	
	path := value
	
	if !strings.Contains(path, ".") {
		return NewErrorf(ErrCodeValidation, "token path must contain at least one dot: %s", path)
	}
	
	parts := strings.Split(path, ".")
	if len(parts) < 2 {
		return NewErrorf(ErrCodeValidation, "token path must have at least 2 segments: %s", path)
	}
	
	for i, part := range parts {
		if part == "" {
			return NewErrorf(ErrCodeValidation, "empty segment at position %d in path: %s", i, path)
		}
		if !isValidTokenSegment(part) {
			return NewErrorf(ErrCodeValidation, "invalid segment '%s' at position %d in path: %s", part, i, path)
		}
	}
	
	return nil
}

func isValidTokenSegment(s string) bool {
	for _, char := range s {
		isValid := (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_'
		if !isValid {
			return false
		}
	}
	return true
}

type Tokens struct {
	Primitives *PrimitiveTokens `json:"primitives" yaml:"primitives"`
	Semantic   *SemanticTokens  `json:"semantic" yaml:"semantic"`
	Components *ComponentTokens `json:"components" yaml:"components"`
}

func (t *Tokens) Validate() error {
	if t == nil {
		return NewError(ErrCodeValidation, "tokens cannot be nil")
	}
	
	if t.Primitives != nil {
		if err := t.Primitives.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid primitives", err)
		}
	}
	
	if t.Semantic != nil {
		if err := t.Semantic.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid semantic tokens", err)
		}
	}
	
	if t.Components != nil {
		if err := t.Components.Validate(); err != nil {
			return WrapError(ErrCodeValidation, "invalid component tokens", err)
		}
	}
	
	return nil
}

func (t *Tokens) Clone() *Tokens {
	if t == nil {
		return nil
	}
	
	cloned := &Tokens{}
	
	if t.Primitives != nil {
		cloned.Primitives = t.Primitives.Clone()
	}
	if t.Semantic != nil {
		cloned.Semantic = t.Semantic.Clone()
	}
	if t.Components != nil {
		cloned.Components = t.Components.Clone()
	}
	
	return cloned
}

type PrimitiveTokens struct {
	Colors      map[string]string `json:"colors" yaml:"colors"`
	Spacing     map[string]string `json:"spacing" yaml:"spacing"`
	Radius      map[string]string `json:"radius" yaml:"radius"`
	Typography  map[string]string `json:"typography" yaml:"typography"`
	Borders     map[string]string `json:"borders" yaml:"borders"`
	Shadows     map[string]string `json:"shadows" yaml:"shadows"`
	Effects     map[string]string `json:"effects" yaml:"effects"`
	Animation   map[string]string `json:"animation" yaml:"animation"`
	ZIndex      map[string]string `json:"zindex" yaml:"zindex"`
	Breakpoints map[string]string `json:"breakpoints" yaml:"breakpoints"`
}

func (pt *PrimitiveTokens) Validate() error {
	if pt == nil {
		return NewError(ErrCodeValidation, "primitive tokens cannot be nil")
	}
	
	categories := map[string]map[string]string{
		"colors":      pt.Colors,
		"spacing":     pt.Spacing,
		"radius":      pt.Radius,
		"typography":  pt.Typography,
		"borders":     pt.Borders,
		"shadows":     pt.Shadows,
		"effects":     pt.Effects,
		"animation":   pt.Animation,
		"zindex":      pt.ZIndex,
		"breakpoints": pt.Breakpoints,
	}
	
	for category, tokens := range categories {
		if tokens == nil {
			continue
		}
		
		for key, value := range tokens {
			ref := TokenReference(value)
			if err := ref.Validate(); err != nil {
				return WrapError(ErrCodeValidation,
					fmt.Sprintf("invalid primitive token %s.%s", category, key), err)
			}
			
			if ref.IsReference() {
				return NewErrorf(ErrCodeValidation,
					"primitive token %s.%s contains reference '%s' (must be CSS literal)",
					category, key, value)
			}
		}
	}
	
	return nil
}

func (pt *PrimitiveTokens) Clone() *PrimitiveTokens {
	if pt == nil {
		return nil
	}
	
	return &PrimitiveTokens{
		Colors:      cloneStringMap(pt.Colors),
		Spacing:     cloneStringMap(pt.Spacing),
		Radius:      cloneStringMap(pt.Radius),
		Typography:  cloneStringMap(pt.Typography),
		Borders:     cloneStringMap(pt.Borders),
		Shadows:     cloneStringMap(pt.Shadows),
		Effects:     cloneStringMap(pt.Effects),
		Animation:   cloneStringMap(pt.Animation),
		ZIndex:      cloneStringMap(pt.ZIndex),
		Breakpoints: cloneStringMap(pt.Breakpoints),
	}
}

type SemanticTokens struct {
	Colors      map[string]string `json:"colors" yaml:"colors"`
	Spacing     map[string]string `json:"spacing" yaml:"spacing"`
	Typography  map[string]string `json:"typography" yaml:"typography"`
	Interactive map[string]string `json:"interactive" yaml:"interactive"`
}

func (st *SemanticTokens) Validate() error {
	if st == nil {
		return nil
	}
	
	categories := map[string]map[string]string{
		"colors":      st.Colors,
		"spacing":     st.Spacing,
		"typography":  st.Typography,
		"interactive": st.Interactive,
	}
	
	for category, tokens := range categories {
		if tokens == nil {
			continue
		}
		
		for key, value := range tokens {
			ref := TokenReference(value)
			if err := ref.Validate(); err != nil {
				return WrapError(ErrCodeValidation,
					fmt.Sprintf("invalid semantic token %s.%s", category, key), err)
			}
		}
	}
	
	return nil
}

func (st *SemanticTokens) Clone() *SemanticTokens {
	if st == nil {
		return nil
	}
	
	return &SemanticTokens{
		Colors:      cloneStringMap(st.Colors),
		Spacing:     cloneStringMap(st.Spacing),
		Typography:  cloneStringMap(st.Typography),
		Interactive: cloneStringMap(st.Interactive),
	}
}

type ComponentTokens map[string]ComponentVariants
type ComponentVariants map[string]StyleProperties
type StyleProperties map[string]string

func (ct *ComponentTokens) Validate() error {
	if ct == nil {
		return nil
	}
	
	for componentName, variants := range *ct {
		if componentName == "" {
			return NewError(ErrCodeValidation, "component name cannot be empty")
		}
		
		if err := variants.Validate(componentName); err != nil {
			return WrapError(ErrCodeValidation,
				fmt.Sprintf("invalid component '%s'", componentName), err)
		}
	}
	
	return nil
}

func (cv ComponentVariants) Validate(component string) error {
	if cv == nil {
		return nil
	}
	
	for variant, props := range cv {
		if variant == "" {
			return NewErrorf(ErrCodeValidation,
				"variant name cannot be empty in component '%s'", component)
		}
		
		if len(props) == 0 {
			return NewErrorf(ErrCodeValidation,
				"variant '%s' in component '%s' has no properties", variant, component)
		}
		
		for property, value := range props {
			if property == "" {
				return NewErrorf(ErrCodeValidation,
					"property name cannot be empty in component '%s', variant '%s'",
					component, variant)
			}
			
			if strings.TrimSpace(value) == "" {
				return NewErrorf(ErrCodeValidation,
					"empty value in component '%s', variant '%s', property '%s'",
					component, variant, property)
			}
		}
	}
	
	return nil
}

func (ct *ComponentTokens) Clone() *ComponentTokens {
	if ct == nil {
		return nil
	}
	
	cloned := make(ComponentTokens, len(*ct))
	for component, variants := range *ct {
		cloned[component] = variants.Clone()
	}
	
	return &cloned
}

func (cv ComponentVariants) Clone() ComponentVariants {
	if cv == nil {
		return nil
	}
	
	cloned := make(ComponentVariants, len(cv))
	for variant, props := range cv {
		cloned[variant] = cloneStringMap(props)
	}
	
	return cloned
}

func cloneStringMap(m map[string]string) map[string]string {
	if m == nil {
		return nil
	}
	
	cloned := make(map[string]string, len(m))
	for k, v := range m {
		cloned[k] = v
	}
	
	return cloned
}
