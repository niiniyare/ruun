// pkg/schema/layout_v2.go
package schema

import (
	"context"
	"fmt"
	"sort"

	"github.com/niiniyare/ruun/pkg/condition"
)

// Layout defines the visual structure of form elements
type Layout struct {
	Type LayoutType `json:"type"`

	// Grid/Flex Configuration
	Columns   int    `json:"columns,omitempty"`
	Gap       string `json:"gap,omitempty"`
	Direction string `json:"direction,omitempty"`
	Wrap      bool   `json:"wrap,omitempty"`

	// Layout Components
	Sections []Section `json:"sections,omitempty"`
	Groups   []Group   `json:"groups,omitempty"`
	Tabs     []Tab     `json:"tabs,omitempty"`
	Steps    []Step    `json:"steps,omitempty"`

	// Responsive Behavior
	Breakpoints *Breakpoints `json:"breakpoints,omitempty"`

	// Styling
	Style *LayoutStyle `json:"style,omitempty"`

	// Internal
	evaluator *condition.Evaluator `json:"-"`
}

// LayoutType defines layout algorithms
type LayoutType string

const (
	LayoutGrid     LayoutType = "grid"
	LayoutFlex     LayoutType = "flex"
	LayoutTabs     LayoutType = "tabs"
	LayoutSteps    LayoutType = "steps"
	LayoutSections LayoutType = "sections"
	LayoutGroups   LayoutType = "groups"
)

// Section represents a logical grouping of fields
type Section struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`

	// Content
	Fields []string `json:"fields"`

	// Behavior
	Collapsible bool `json:"collapsible,omitempty"`
	Collapsed   bool `json:"collapsed,omitempty"`

	// Layout
	Columns int `json:"columns,omitempty"`
	Order   int `json:"order,omitempty"`

	// Conditional
	Conditional *LayoutConditional `json:"conditional,omitempty"`

	// Styling
	Style *SectionStyle `json:"style,omitempty"`
}

// Group represents a visual grouping (like a fieldset)
type Group struct {
	ID          string `json:"id"`
	Label       string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`

	// Content
	Fields []string `json:"fields"`

	// Behavior
	Border  bool `json:"border,omitempty"`
	Columns int  `json:"columns,omitempty"`
	Order   int  `json:"order,omitempty"`

	// Conditional
	Conditional *LayoutConditional `json:"conditional,omitempty"`

	// Styling
	Style *GroupStyle `json:"style,omitempty"`
}

// Tab represents a tab in tabbed interface
type Tab struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Icon        string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
	Badge       string `json:"badge,omitempty"`

	// Content
	Fields []string `json:"fields"`

	// State
	Disabled bool `json:"disabled,omitempty"`
	Order    int  `json:"order,omitempty"`

	// Conditional
	Conditional *LayoutConditional `json:"conditional,omitempty"`

	// Styling
	Style *TabStyle `json:"style,omitempty"`
}

// Step represents a step in multi-step wizard
type Step struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`

	// Content
	Fields []string `json:"fields"`

	// Behavior
	Order      int  `json:"order"`
	Skippable  bool `json:"skippable,omitempty"`
	Validation bool `json:"validation,omitempty"`

	// Conditional
	Conditional *LayoutConditional `json:"conditional,omitempty"`

	// Styling
	Style *StepStyle `json:"style,omitempty"`
}

// LayoutConditional defines conditional visibility for layout components
type LayoutConditional struct {
	Show *condition.ConditionGroup `json:"show,omitempty"`
	Hide *condition.ConditionGroup `json:"hide,omitempty"`
}

// Breakpoints defines responsive behavior
type Breakpoints struct {
	Mobile  *BreakpointConfig            `json:"mobile,omitempty"`
	Tablet  *BreakpointConfig            `json:"tablet,omitempty"`
	Desktop *BreakpointConfig            `json:"desktop,omitempty"`
	Custom  map[string]*BreakpointConfig `json:"custom,omitempty"`
}

// BreakpointConfig defines layout at specific screen size
type BreakpointConfig struct {
	Columns    int      `json:"columns,omitempty"`
	Gap        string   `json:"gap,omitempty"`
	Direction  string   `json:"direction,omitempty"`
	HideFields []string `json:"hideFields,omitempty"`
	ShowFields []string `json:"showFields,omitempty"`
}

// Style Types
type LayoutStyle struct {
	Classes      string            `json:"classes,omitempty"`
	Colors       map[string]string `json:"colors,omitempty"`
	Spacing      map[string]string `json:"spacing,omitempty"`
	BorderRadius string            `json:"borderRadius,omitempty"`
	CustomCSS    string            `json:"customCSS,omitempty"`
}

type SectionStyle struct {
	Background   string            `json:"background,omitempty"`
	Border       string            `json:"border,omitempty"`
	Padding      string            `json:"padding,omitempty"`
	BorderRadius string            `json:"borderRadius,omitempty"`
	Colors       map[string]string `json:"colors,omitempty"`
	CustomCSS    string            `json:"customCSS,omitempty"`
}

type GroupStyle struct {
	Background   string            `json:"background,omitempty"`
	Border       string            `json:"border,omitempty"`
	Padding      string            `json:"padding,omitempty"`
	BorderRadius string            `json:"borderRadius,omitempty"`
	Colors       map[string]string `json:"colors,omitempty"`
	CustomCSS    string            `json:"customCSS,omitempty"`
}

type TabStyle struct {
	ActiveBackground   string            `json:"activeBackground,omitempty"`
	InactiveBackground string            `json:"inactiveBackground,omitempty"`
	ActiveColor        string            `json:"activeColor,omitempty"`
	InactiveColor      string            `json:"inactiveColor,omitempty"`
	BorderColor        string            `json:"borderColor,omitempty"`
	Colors             map[string]string `json:"colors,omitempty"`
	CustomCSS          string            `json:"customCSS,omitempty"`
}

type StepStyle struct {
	ActiveColor    string            `json:"activeColor,omitempty"`
	CompletedColor string            `json:"completedColor,omitempty"`
	InactiveColor  string            `json:"inactiveColor,omitempty"`
	ConnectorColor string            `json:"connectorColor,omitempty"`
	Colors         map[string]string `json:"colors,omitempty"`
	CustomCSS      string            `json:"customCSS,omitempty"`
}

// Core Methods

// SetEvaluator sets the condition evaluator
func (l *Layout) SetEvaluator(evaluator *condition.Evaluator) {
	l.evaluator = evaluator
}

// GetEvaluator returns the condition evaluator
func (l *Layout) GetEvaluator() *condition.Evaluator {
	return l.evaluator
}

// Validate validates layout configuration
func (l *Layout) Validate(schema *Schema) error {
	validator := newLayoutValidator()
	return validator.Validate(l, schema)
}

// Configuration Methods

// GetColumns returns appropriate column count
func (l *Layout) GetColumns() int {
	if l.Columns > 0 {
		return l.Columns
	}
	switch l.Type {
	case LayoutGrid:
		return 2
	case LayoutFlex:
		return 1
	default:
		return 1
	}
}

// GetGap returns spacing value with fallback
func (l *Layout) GetGap() string {
	if l.Gap != "" {
		return l.Gap
	}
	return "1rem"
}

// GetDirection returns direction with fallback
func (l *Layout) GetDirection() string {
	if l.Direction != "" {
		return l.Direction
	}
	return "column"
}

// Type Checking Methods

// HasTabs checks if layout uses tabs
func (l *Layout) HasTabs() bool {
	return l.Type == LayoutTabs && len(l.Tabs) > 0
}

// HasSteps checks if layout is multi-step wizard
func (l *Layout) HasSteps() bool {
	return l.Type == LayoutSteps && len(l.Steps) > 0
}

// HasSections checks if layout has sections
func (l *Layout) HasSections() bool {
	return (l.Type == LayoutSections || len(l.Sections) > 0)
}

// HasGroups checks if layout has groups
func (l *Layout) HasGroups() bool {
	return (l.Type == LayoutGroups || len(l.Groups) > 0)
}

// Field Retrieval Methods

// GetFieldsForSection returns fields for a section
func (l *Layout) GetFieldsForSection(sectionID string) []string {
	for _, section := range l.Sections {
		if section.ID == sectionID {
			return section.Fields
		}
	}
	return []string{}
}

// GetFieldsForTab returns fields for a tab
func (l *Layout) GetFieldsForTab(tabID string) []string {
	for _, tab := range l.Tabs {
		if tab.ID == tabID {
			return tab.Fields
		}
	}
	return []string{}
}

// GetFieldsForStep returns fields for a step
func (l *Layout) GetFieldsForStep(stepID string) []string {
	for _, step := range l.Steps {
		if step.ID == stepID {
			return step.Fields
		}
	}
	return []string{}
}

// GetFieldsForGroup returns fields for a group
func (l *Layout) GetFieldsForGroup(groupID string) []string {
	for _, group := range l.Groups {
		if group.ID == groupID {
			return group.Fields
		}
	}
	return []string{}
}

// GetAllFields returns all fields referenced in layout
func (l *Layout) GetAllFields() []string {
	fieldSet := make(map[string]bool)

	for _, section := range l.Sections {
		for _, field := range section.Fields {
			fieldSet[field] = true
		}
	}

	for _, tab := range l.Tabs {
		for _, field := range tab.Fields {
			fieldSet[field] = true
		}
	}

	for _, step := range l.Steps {
		for _, field := range step.Fields {
			fieldSet[field] = true
		}
	}

	for _, group := range l.Groups {
		for _, field := range group.Fields {
			fieldSet[field] = true
		}
	}

	fields := make([]string, 0, len(fieldSet))
	for field := range fieldSet {
		fields = append(fields, field)
	}
	return fields
}

// Component Retrieval Methods

// GetSection returns a section by ID
func (l *Layout) GetSection(id string) (*Section, bool) {
	for i := range l.Sections {
		if l.Sections[i].ID == id {
			return &l.Sections[i], true
		}
	}
	return nil, false
}

// GetTab returns a tab by ID
func (l *Layout) GetTab(id string) (*Tab, bool) {
	for i := range l.Tabs {
		if l.Tabs[i].ID == id {
			return &l.Tabs[i], true
		}
	}
	return nil, false
}

// GetStep returns a step by ID
func (l *Layout) GetStep(id string) (*Step, bool) {
	for i := range l.Steps {
		if l.Steps[i].ID == id {
			return &l.Steps[i], true
		}
	}
	return nil, false
}

// GetGroup returns a group by ID
func (l *Layout) GetGroup(id string) (*Group, bool) {
	for i := range l.Groups {
		if l.Groups[i].ID == id {
			return &l.Groups[i], true
		}
	}
	return nil, false
}

// Ordering Methods

// GetOrderedSections returns sections sorted by order
func (l *Layout) GetOrderedSections() []Section {
	sections := make([]Section, len(l.Sections))
	copy(sections, l.Sections)
	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Order < sections[j].Order
	})
	return sections
}

// GetOrderedGroups returns groups sorted by order
func (l *Layout) GetOrderedGroups() []Group {
	groups := make([]Group, len(l.Groups))
	copy(groups, l.Groups)
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Order < groups[j].Order
	})
	return groups
}

// GetOrderedTabs returns tabs sorted by order
func (l *Layout) GetOrderedTabs() []Tab {
	tabs := make([]Tab, len(l.Tabs))
	copy(tabs, l.Tabs)
	sort.Slice(tabs, func(i, j int) bool {
		return tabs[i].Order < tabs[j].Order
	})
	return tabs
}

// GetOrderedSteps returns steps sorted by order
func (l *Layout) GetOrderedSteps() []Step {
	steps := make([]Step, len(l.Steps))
	copy(steps, l.Steps)
	sort.Slice(steps, func(i, j int) bool {
		return steps[i].Order < steps[j].Order
	})
	return steps
}

// Conditional Visibility Methods

// GetVisibleSections returns visible sections
func (l *Layout) GetVisibleSections(ctx context.Context, data map[string]any) []Section {
	visible := make([]Section, 0)
	for _, section := range l.Sections {
		if isVisible, _ := l.isSectionVisible(ctx, &section, data); isVisible {
			visible = append(visible, section)
		}
	}
	return visible
}

// GetVisibleTabs returns visible tabs
func (l *Layout) GetVisibleTabs(ctx context.Context, data map[string]any) []Tab {
	visible := make([]Tab, 0)
	for _, tab := range l.Tabs {
		if isVisible, _ := l.isTabVisible(ctx, &tab, data); isVisible && !tab.Disabled {
			visible = append(visible, tab)
		}
	}
	return visible
}

// GetVisibleSteps returns visible steps
func (l *Layout) GetVisibleSteps(ctx context.Context, data map[string]any) []Step {
	visible := make([]Step, 0)
	for _, step := range l.Steps {
		if isVisible, _ := l.isStepVisible(ctx, &step, data); isVisible {
			visible = append(visible, step)
		}
	}
	return visible
}

// GetVisibleGroups returns visible groups
func (l *Layout) GetVisibleGroups(ctx context.Context, data map[string]any) []Group {
	visible := make([]Group, 0)
	for _, group := range l.Groups {
		if isVisible, _ := l.isGroupVisible(ctx, &group, data); isVisible {
			visible = append(visible, group)
		}
	}
	return visible
}

// Internal visibility checkers

func (l *Layout) isSectionVisible(ctx context.Context, section *Section, data map[string]any) (bool, error) {
	if section.Conditional == nil {
		return true, nil
	}
	return l.evaluateConditional(ctx, section.Conditional, data)
}

func (l *Layout) isTabVisible(ctx context.Context, tab *Tab, data map[string]any) (bool, error) {
	if tab.Conditional == nil {
		return true, nil
	}
	return l.evaluateConditional(ctx, tab.Conditional, data)
}

func (l *Layout) isStepVisible(ctx context.Context, step *Step, data map[string]any) (bool, error) {
	if step.Conditional == nil {
		return true, nil
	}
	return l.evaluateConditional(ctx, step.Conditional, data)
}

func (l *Layout) isGroupVisible(ctx context.Context, group *Group, data map[string]any) (bool, error) {
	if group.Conditional == nil {
		return true, nil
	}
	return l.evaluateConditional(ctx, group.Conditional, data)
}

func (l *Layout) evaluateConditional(ctx context.Context, conditional *LayoutConditional, data map[string]any) (bool, error) {
	if l.evaluator == nil {
		return true, nil
	}

	// Check Hide first
	if conditional.Hide != nil {
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		hide, err := l.evaluator.Evaluate(ctx, conditional.Hide, evalCtx)
		if err != nil {
			return false, err
		}
		if hide {
			return false, nil
		}
	}

	// Check Show
	if conditional.Show != nil {
		evalCtx := condition.NewEvalContext(data, condition.DefaultEvalOptions())
		return l.evaluator.Evaluate(ctx, conditional.Show, evalCtx)
	}

	return true, nil
}

// Responsive Methods

// GetBreakpointColumns returns columns for breakpoint
func (l *Layout) GetBreakpointColumns(breakpoint string) int {
	if l.Breakpoints == nil {
		return l.GetColumns()
	}

	var config *BreakpointConfig
	switch breakpoint {
	case "mobile":
		config = l.Breakpoints.Mobile
	case "tablet":
		config = l.Breakpoints.Tablet
	case "desktop":
		config = l.Breakpoints.Desktop
	default:
		if l.Breakpoints.Custom != nil {
			config = l.Breakpoints.Custom[breakpoint]
		}
	}

	if config != nil && config.Columns > 0 {
		return config.Columns
	}
	return l.GetColumns()
}

// GetBreakpointGap returns gap for breakpoint
func (l *Layout) GetBreakpointGap(breakpoint string) string {
	if l.Breakpoints == nil {
		return l.GetGap()
	}

	var config *BreakpointConfig
	switch breakpoint {
	case "mobile":
		config = l.Breakpoints.Mobile
	case "tablet":
		config = l.Breakpoints.Tablet
	case "desktop":
		config = l.Breakpoints.Desktop
	default:
		if l.Breakpoints.Custom != nil {
			config = l.Breakpoints.Custom[breakpoint]
		}
	}

	if config != nil && config.Gap != "" {
		return config.Gap
	}
	return l.GetGap()
}

// Clone creates a copy of the layout
func (l *Layout) Clone() *Layout {
	clone := *l
	clone.evaluator = l.evaluator
	return &clone
}

// layoutValidator validates layout configuration
type layoutValidator struct{}

func newLayoutValidator() *layoutValidator {
	return &layoutValidator{}
}

func (v *layoutValidator) Validate(layout *Layout, schema *Schema) error {
	collector := NewErrorCollector()

	// Validate sections
	v.validateSections(layout, schema, collector)

	// Validate tabs
	v.validateTabs(layout, schema, collector)

	// Validate steps
	v.validateSteps(layout, schema, collector)

	// Validate groups
	v.validateGroups(layout, schema, collector)

	// Check for duplicate IDs
	v.checkDuplicateIDs(layout, collector)

	if collector.HasErrors() {
		return collector.Errors()
	}
	return nil
}

func (v *layoutValidator) validateSections(layout *Layout, schema *Schema, collector *ErrorCollector) {
	for i, section := range layout.Sections {
		if section.ID == "" {
			collector.AddValidationError(
				fmt.Sprintf("Sections[%d]", i),
				"missing_id",
				"section ID is required",
			)
		}

		for _, fieldName := range section.Fields {
			if !schema.HasField(fieldName) {
				collector.AddValidationError(
					"Section."+section.ID,
					"invalid_field",
					fmt.Sprintf("references non-existent field: %s", fieldName),
				)
			}
		}
	}
}

func (v *layoutValidator) validateTabs(layout *Layout, schema *Schema, collector *ErrorCollector) {
	for i, tab := range layout.Tabs {
		if tab.ID == "" {
			collector.AddValidationError(
				fmt.Sprintf("Tabs[%d]", i),
				"missing_id",
				"tab ID is required",
			)
		}

		if tab.Label == "" {
			collector.AddValidationError(
				"Tab."+tab.ID,
				"missing_label",
				"tab label is required",
			)
		}

		for _, fieldName := range tab.Fields {
			if !schema.HasField(fieldName) {
				collector.AddValidationError(
					"Tab."+tab.ID,
					"invalid_field",
					fmt.Sprintf("references non-existent field: %s", fieldName),
				)
			}
		}
	}
}

func (v *layoutValidator) validateSteps(layout *Layout, schema *Schema, collector *ErrorCollector) {
	for i, step := range layout.Steps {
		if step.ID == "" {
			collector.AddValidationError(
				fmt.Sprintf("Steps[%d]", i),
				"missing_id",
				"step ID is required",
			)
		}

		if step.Title == "" {
			collector.AddValidationError(
				"Step."+step.ID,
				"missing_title",
				"step title is required",
			)
		}

		for _, fieldName := range step.Fields {
			if !schema.HasField(fieldName) {
				collector.AddValidationError(
					"Step."+step.ID,
					"invalid_field",
					fmt.Sprintf("references non-existent field: %s", fieldName),
				)
			}
		}
	}
}

func (v *layoutValidator) validateGroups(layout *Layout, schema *Schema, collector *ErrorCollector) {
	for i, group := range layout.Groups {
		if group.ID == "" {
			collector.AddValidationError(
				fmt.Sprintf("Groups[%d]", i),
				"missing_id",
				"group ID is required",
			)
		}

		for _, fieldName := range group.Fields {
			if !schema.HasField(fieldName) {
				collector.AddValidationError(
					"Group."+group.ID,
					"invalid_field",
					fmt.Sprintf("references non-existent field: %s", fieldName),
				)
			}
		}
	}
}

func (v *layoutValidator) checkDuplicateIDs(layout *Layout, collector *ErrorCollector) {
	ids := make(map[string]bool)

	for _, section := range layout.Sections {
		if ids[section.ID] {
			collector.AddValidationError(
				"Section."+section.ID,
				"duplicate_id",
				"duplicate section ID: "+section.ID,
			)
		}
		ids[section.ID] = true
	}

	for _, tab := range layout.Tabs {
		if ids[tab.ID] {
			collector.AddValidationError(
				"Tab."+tab.ID,
				"duplicate_id",
				"duplicate tab ID: "+tab.ID,
			)
		}
		ids[tab.ID] = true
	}

	for _, step := range layout.Steps {
		if ids[step.ID] {
			collector.AddValidationError(
				"Step."+step.ID,
				"duplicate_id",
				"duplicate step ID: "+step.ID,
			)
		}
		ids[step.ID] = true
	}

	for _, group := range layout.Groups {
		if ids[group.ID] {
			collector.AddValidationError(
				"Group."+group.ID,
				"duplicate_id",
				"duplicate group ID: "+group.ID,
			)
		}
		ids[group.ID] = true
	}
}

// LayoutBuilder provides fluent API for building layouts
type LayoutBuilder struct {
	layout    Layout
	evaluator *condition.Evaluator
}

// NewLayout starts building a layout
func NewLayout(layoutType LayoutType) *LayoutBuilder {
	return &LayoutBuilder{
		layout: Layout{
			Type: layoutType,
		},
	}
}

func (b *LayoutBuilder) WithColumns(columns int) *LayoutBuilder {
	b.layout.Columns = columns
	return b
}

func (b *LayoutBuilder) WithGap(gap string) *LayoutBuilder {
	b.layout.Gap = gap
	return b
}

func (b *LayoutBuilder) WithDirection(direction string) *LayoutBuilder {
	b.layout.Direction = direction
	return b
}

func (b *LayoutBuilder) WithWrap(wrap bool) *LayoutBuilder {
	b.layout.Wrap = wrap
	return b
}

func (b *LayoutBuilder) AddSection(section Section) *LayoutBuilder {
	b.layout.Sections = append(b.layout.Sections, section)
	return b
}

func (b *LayoutBuilder) AddSections(sections ...Section) *LayoutBuilder {
	b.layout.Sections = append(b.layout.Sections, sections...)
	return b
}

func (b *LayoutBuilder) AddGroup(group Group) *LayoutBuilder {
	b.layout.Groups = append(b.layout.Groups, group)
	return b
}

func (b *LayoutBuilder) AddGroups(groups ...Group) *LayoutBuilder {
	b.layout.Groups = append(b.layout.Groups, groups...)
	return b
}

func (b *LayoutBuilder) AddTab(tab Tab) *LayoutBuilder {
	b.layout.Tabs = append(b.layout.Tabs, tab)
	return b
}

func (b *LayoutBuilder) AddTabs(tabs ...Tab) *LayoutBuilder {
	b.layout.Tabs = append(b.layout.Tabs, tabs...)
	return b
}

func (b *LayoutBuilder) AddStep(step Step) *LayoutBuilder {
	b.layout.Steps = append(b.layout.Steps, step)
	return b
}

func (b *LayoutBuilder) AddSteps(steps ...Step) *LayoutBuilder {
	b.layout.Steps = append(b.layout.Steps, steps...)
	return b
}

func (b *LayoutBuilder) WithBreakpoints(breakpoints *Breakpoints) *LayoutBuilder {
	b.layout.Breakpoints = breakpoints
	return b
}

func (b *LayoutBuilder) WithStyle(style *LayoutStyle) *LayoutBuilder {
	b.layout.Style = style
	return b
}

func (b *LayoutBuilder) WithEvaluator(evaluator *condition.Evaluator) *LayoutBuilder {
	b.evaluator = evaluator
	return b
}

func (b *LayoutBuilder) Build() Layout {
	if b.evaluator != nil {
		b.layout.SetEvaluator(b.evaluator)
	}
	return b.layout
}

// Component Builders

// SectionBuilder builds sections
type SectionBuilder struct {
	section Section
}

// NewSection starts building a section
func NewSection(id, title string) *SectionBuilder {
	return &SectionBuilder{
		section: Section{
			ID:    id,
			Title: title,
		},
	}
}

func (b *SectionBuilder) WithDescription(desc string) *SectionBuilder {
	b.section.Description = desc
	return b
}

func (b *SectionBuilder) WithIcon(icon string) *SectionBuilder {
	b.section.Icon = icon
	return b
}

func (b *SectionBuilder) WithFields(fields ...string) *SectionBuilder {
	b.section.Fields = append(b.section.Fields, fields...)
	return b
}

func (b *SectionBuilder) Collapsible(collapsed bool) *SectionBuilder {
	b.section.Collapsible = true
	b.section.Collapsed = collapsed
	return b
}

func (b *SectionBuilder) WithColumns(columns int) *SectionBuilder {
	b.section.Columns = columns
	return b
}

func (b *SectionBuilder) WithOrder(order int) *SectionBuilder {
	b.section.Order = order
	return b
}

func (b *SectionBuilder) WithConditional(conditional *LayoutConditional) *SectionBuilder {
	b.section.Conditional = conditional
	return b
}

func (b *SectionBuilder) WithStyle(style *SectionStyle) *SectionBuilder {
	b.section.Style = style
	return b
}

func (b *SectionBuilder) Build() Section {
	return b.section
}

// GroupBuilder builds groups
type GroupBuilder struct {
	group Group
}

// NewGroup starts building a group
func NewGroup(id, label string) *GroupBuilder {
	return &GroupBuilder{
		group: Group{
			ID:    id,
			Label: label,
		},
	}
}

func (b *GroupBuilder) WithDescription(desc string) *GroupBuilder {
	b.group.Description = desc
	return b
}

func (b *GroupBuilder) WithFields(fields ...string) *GroupBuilder {
	b.group.Fields = append(b.group.Fields, fields...)
	return b
}

func (b *GroupBuilder) WithBorder(border bool) *GroupBuilder {
	b.group.Border = border
	return b
}

func (b *GroupBuilder) WithColumns(columns int) *GroupBuilder {
	b.group.Columns = columns
	return b
}

func (b *GroupBuilder) WithOrder(order int) *GroupBuilder {
	b.group.Order = order
	return b
}

func (b *GroupBuilder) WithConditional(conditional *LayoutConditional) *GroupBuilder {
	b.group.Conditional = conditional
	return b
}

func (b *GroupBuilder) WithStyle(style *GroupStyle) *GroupBuilder {
	b.group.Style = style
	return b
}

func (b *GroupBuilder) Build() Group {
	return b.group
}

// TabBuilder builds tabs
type TabBuilder struct {
	tab Tab
}

// NewTab starts building a tab
func NewTab(id, label string) *TabBuilder {
	return &TabBuilder{
		tab: Tab{
			ID:    id,
			Label: label,
		},
	}
}

func (b *TabBuilder) WithIcon(icon string) *TabBuilder {
	b.tab.Icon = icon
	return b
}

func (b *TabBuilder) WithDescription(desc string) *TabBuilder {
	b.tab.Description = desc
	return b
}

func (b *TabBuilder) WithBadge(badge string) *TabBuilder {
	b.tab.Badge = badge
	return b
}

func (b *TabBuilder) WithFields(fields ...string) *TabBuilder {
	b.tab.Fields = append(b.tab.Fields, fields...)
	return b
}

func (b *TabBuilder) Disabled() *TabBuilder {
	b.tab.Disabled = true
	return b
}

func (b *TabBuilder) WithOrder(order int) *TabBuilder {
	b.tab.Order = order
	return b
}

func (b *TabBuilder) WithConditional(conditional *LayoutConditional) *TabBuilder {
	b.tab.Conditional = conditional
	return b
}

func (b *TabBuilder) WithStyle(style *TabStyle) *TabBuilder {
	b.tab.Style = style
	return b
}

func (b *TabBuilder) Build() Tab {
	return b.tab
}

// StepBuilder builds steps
type StepBuilder struct {
	step Step
}

// NewStep starts building a step
func NewStep(id, title string, order int) *StepBuilder {
	return &StepBuilder{
		step: Step{
			ID:    id,
			Title: title,
			Order: order,
		},
	}
}

func (b *StepBuilder) WithDescription(desc string) *StepBuilder {
	b.step.Description = desc
	return b
}

func (b *StepBuilder) WithIcon(icon string) *StepBuilder {
	b.step.Icon = icon
	return b
}

func (b *StepBuilder) WithFields(fields ...string) *StepBuilder {
	b.step.Fields = append(b.step.Fields, fields...)
	return b
}

func (b *StepBuilder) Skippable() *StepBuilder {
	b.step.Skippable = true
	return b
}

func (b *StepBuilder) WithValidation(validate bool) *StepBuilder {
	b.step.Validation = validate
	return b
}

func (b *StepBuilder) WithConditional(conditional *LayoutConditional) *StepBuilder {
	b.step.Conditional = conditional
	return b
}

func (b *StepBuilder) WithStyle(style *StepStyle) *StepBuilder {
	b.step.Style = style
	return b
}

func (b *StepBuilder) Build() Step {
	return b.step
}

// Convenience Constructors

// NewGridLayout creates a grid layout
func NewGridLayout(columns int) *LayoutBuilder {
	return NewLayout(LayoutGrid).WithColumns(columns).WithGap("1rem")
}

// NewFlexLayout creates a flex layout
func NewFlexLayout(direction string) *LayoutBuilder {
	return NewLayout(LayoutFlex).WithDirection(direction).WithGap("1rem")
}

// NewTabLayout creates a tabbed layout
func NewTabLayout() *LayoutBuilder {
	return NewLayout(LayoutTabs)
}

// NewStepLayout creates a multi-step wizard
func NewStepLayout() *LayoutBuilder {
	return NewLayout(LayoutSteps)
}

// NewSectionLayout creates a sectioned layout
func NewSectionLayout() *LayoutBuilder {
	return NewLayout(LayoutSections)
}

// NewGroupLayout creates a grouped layout
func NewGroupLayout() *LayoutBuilder {
	return NewLayout(LayoutGroups)
}

// Helper Functions

// CreateBreakpoints creates responsive breakpoints
func CreateBreakpoints(mobile, tablet, desktop int) *Breakpoints {
	return &Breakpoints{
		Mobile: &BreakpointConfig{
			Columns: mobile,
		},
		Tablet: &BreakpointConfig{
			Columns: tablet,
		},
		Desktop: &BreakpointConfig{
			Columns: desktop,
		},
	}
}
