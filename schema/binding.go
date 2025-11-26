package schema

// Binding defines reactive state bindings.
// Can be rendered to Alpine.js, Vue, or other reactive frameworks.
type Binding struct {
    // Data binding
    Data   string `json:"data,omitempty"`   // x-data equivalent
    Model  string `json:"model,omitempty"`  // Two-way binding
    Text   string `json:"text,omitempty"`   // Text content binding
    HTML   string `json:"html,omitempty"`   // HTML content binding
    
    // Visibility
    Show string `json:"show,omitempty"`
    If   string `json:"if,omitempty"`
    
    // Attributes
    Attrs map[string]string `json:"attrs,omitempty"`
    Class string            `json:"class,omitempty"`
    Style string            `json:"style,omitempty"`
    
    // Events
    On map[string]string `json:"on,omitempty"`
    
    // Iteration
    For    string `json:"for,omitempty"`
    ForKey string `json:"forKey,omitempty"`
    
    // Lifecycle
    Init    string `json:"init,omitempty"`
    Effect  string `json:"effect,omitempty"`
    Destroy string `json:"destroy,omitempty"`
    
    // Reference
    Ref string `json:"ref,omitempty"`
    
    // Transitions
    Transition string `json:"transition,omitempty"`
}

func (b *Binding) IsEmpty() bool {
    return b == nil || (b.Data == "" && b.Model == "" && b.Show == "" && len(b.On) == 0)
}

func (b *Binding) HasModel() bool {
    return b != nil && b.Model != ""
}

func (b *Binding) HasVisibility() bool {
    return b != nil && (b.Show != "" || b.If != "")
}

func (b *Binding) HasEvents() bool {
    return b != nil && len(b.On) > 0
}

// BindingBuilder for fluent construction
type BindingBuilder struct {
    binding Binding
}

func NewBinding() *BindingBuilder {
    return &BindingBuilder{binding: Binding{On: make(map[string]string)}}
}

func (b *BindingBuilder) Data(expr string) *BindingBuilder {
    b.binding.Data = expr
    return b
}

func (b *BindingBuilder) Model(field string) *BindingBuilder {
    b.binding.Model = field
    return b
}

func (b *BindingBuilder) Show(condition string) *BindingBuilder {
    b.binding.Show = condition
    return b
}

func (b *BindingBuilder) On(event, handler string) *BindingBuilder {
    b.binding.On[event] = handler
    return b
}

func (b *BindingBuilder) Build() *Binding {
    return &b.binding
}