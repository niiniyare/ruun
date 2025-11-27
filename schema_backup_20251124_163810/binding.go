package schema

// Binding defines reactive state bindings.
// Replaces Alpine, ActionAlpine, FieldAlpine
type Binding struct {
	Model      string            `json:"model,omitempty"`
	Text       string            `json:"text,omitempty"`
	HTML       string            `json:"html,omitempty"`
	Show       string            `json:"show,omitempty"`
	If         string            `json:"if,omitempty"`
	Attrs      map[string]string `json:"attrs,omitempty"`
	Class      string            `json:"class,omitempty"`
	Style      string            `json:"style,omitempty"`
	On         map[string]string `json:"on,omitempty"`
	For        string            `json:"for,omitempty"`
	ForKey     string            `json:"forKey,omitempty"`
	Init       string            `json:"init,omitempty"`
	Effect     string            `json:"effect,omitempty"`
	Destroy    string            `json:"destroy,omitempty"`
	Ref        string            `json:"ref,omitempty"`
	Transition string            `json:"transition,omitempty"`
	Data       string            `json:"data,omitempty"`
}

func (b *Binding) HasModel() bool {
	return b != nil && b.Model != ""
}

func (b *Binding) HasVisibility() bool {
	return b != nil && (b.Show != "" || b.If != "")
}
