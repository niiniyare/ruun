package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// DataTableSchemaBuilder helps build DataTable components from schema definitions
// Note: Schema package implementation pending
type DataTableSchemaBuilder struct {
	// schema        *schema.Schema
	config        *DataTableConfig
	fieldMappings map[string]ColumnMapping
	filters       []SchemaFilter
	actions       []SchemaAction
}

// DataTableConfig contains configuration for schema-driven tables
type DataTableConfig struct {
	// Basic display
	ShowTitle       bool `json:"showTitle"`
	ShowDescription bool `json:"showDescription"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`

	// Data operations
	EnableSelection   bool `json:"enableSelection"`
	EnableMultiSelect bool `json:"enableMultiSelect"`
	EnableSorting     bool `json:"enableSorting"`
	EnableFiltering   bool `json:"enableFiltering"`
	EnableSearch      bool `json:"enableSearch"`
	EnableExport      bool `json:"enableExport"`

	// Pagination
	EnablePagination  bool  `json:"enablePagination"`
	DefaultPageSize   int   `json:"defaultPageSize"`
	PageSizeOptions   []int `json:"pageSizeOptions"`

	// Search configuration
	SearchPlaceholder string   `json:"searchPlaceholder,omitempty"`
	SearchFields      []string `json:"searchFields,omitempty"`

	// Export configuration
	ExportFormats  []ExportFormat `json:"exportFormats,omitempty"`
	ExportFilename string         `json:"exportFilename,omitempty"`

	// Display settings
	Variant DataTableVariant `json:"variant"`
	Size    DataTableSize    `json:"size"`
	Density DataTableDensity `json:"density"`

	// Column configuration
	ColumnDefaults  ColumnDefaults            `json:"columnDefaults"`
	ColumnOverrides map[string]ColumnOverride `json:"columnOverrides,omitempty"`

	// Actions
	EnableDefaultActions bool   `json:"enableDefaultActions"`
	ActionPosition       string `json:"actionPosition"` // "left", "right"

	// UI features
	EnableColumnResize  bool `json:"enableColumnResize"`
	EnableColumnReorder bool `json:"enableColumnReorder"`
	VirtualScrolling    bool `json:"virtualScrolling"`
	LazyLoading         bool `json:"lazyLoading"`
	CacheData           bool `json:"cacheData"`

	// Accessibility
	AriaLabels map[string]string `json:"ariaLabels,omitempty"`

	// Custom templates
	CustomTemplates map[string]string `json:"customTemplates,omitempty"`

	// Conditional display
	ShowConditions []ShowCondition `json:"showConditions,omitempty"`
}

// ColumnDefaults provides default settings for all columns
type ColumnDefaults struct {
	Visible    bool   `json:"visible"`
	Sortable   bool   `json:"sortable"`
	Filterable bool   `json:"filterable"`
	Searchable bool   `json:"searchable"`
	Resizable  bool   `json:"resizable"`
	Align      string `json:"align"`
	Width      string `json:"width,omitempty"`
	MinWidth   string `json:"minWidth,omitempty"`
	MaxWidth   string `json:"maxWidth,omitempty"`
}

// ColumnMapping defines how a schema field maps to a table column
type ColumnMapping struct {
	// Basic mapping
	FieldName   string     `json:"fieldName"`
	ColumnType  ColumnType `json:"columnType"`
	DisplayName string     `json:"displayName,omitempty"`

	// Rendering options
	Renderer string `json:"renderer,omitempty"` // custom, badge, link, etc.
	Template string `json:"template,omitempty"` // custom template
	Format   string `json:"format,omitempty"`   // date/number format

	// Column properties
	Visible    *bool `json:"visible,omitempty"`
	Sortable   *bool `json:"sortable,omitempty"`
	Filterable *bool `json:"filterable,omitempty"`
	Searchable *bool `json:"searchable,omitempty"`
	Resizable  *bool `json:"resizable,omitempty"`
	Clickable  *bool `json:"clickable,omitempty"`

	// Layout
	Align         string `json:"align,omitempty"`
	VerticalAlign string `json:"verticalAlign,omitempty"`
	Width         string `json:"width,omitempty"`
	MinWidth      string `json:"minWidth,omitempty"`
	MaxWidth      string `json:"maxWidth,omitempty"`
	Fixed         bool   `json:"fixed"`

	// Special rendering
	BadgeConfig    *BadgeConfig    `json:"badgeConfig,omitempty"`
	LinkConfig     *LinkConfig     `json:"linkConfig,omitempty"`
	ImageConfig    *ImageConfig    `json:"imageConfig,omitempty"`
	ProgressConfig *ProgressConfig `json:"progressConfig,omitempty"`

	// Filtering
	FilterType    string         `json:"filterType,omitempty"` // text, select, date, number
	FilterOptions []FilterOption `json:"filterOptions,omitempty"`

	// Actions
	ColumnActions []ActionConfig `json:"columnActions,omitempty"`

	// Conditional display
	ShowCondition string `json:"showCondition,omitempty"`
	HideCondition string `json:"hideCondition,omitempty"`
}

// ColumnOverride provides specific overrides for individual columns
type ColumnOverride struct {
	Mapping ColumnMapping `json:"mapping"`
	Order   int           `json:"order"`           // Display order
	Group   string        `json:"group,omitempty"` // Column group
}

// BadgeConfig configures badge rendering
type BadgeConfig struct {
	VariantMap     map[string]string `json:"variantMap"`              // value -> variant
	ColorMap       map[string]string `json:"colorMap,omitempty"`      // value -> color
	IconMap        map[string]string `json:"iconMap,omitempty"`       // value -> icon
	DefaultVariant string            `json:"defaultVariant,omitempty"`
}

// LinkConfig configures link rendering
type LinkConfig struct {
	URLField string `json:"urlField"`              // Field containing URL
	Target   string `json:"target"`                // _blank, _self, etc.
	ShowIcon bool   `json:"showIcon"`              // Show external link icon
	Template string `json:"template,omitempty"`    // URL template
}

// ImageConfig configures image rendering
type ImageConfig struct {
	AltField    string `json:"altField,omitempty"`    // Field containing alt text
	Size        string `json:"size"`                  // sm, md, lg
	Shape       string `json:"shape"`                 // square, circle, rounded
	Lazy        bool   `json:"lazy"`                  // Lazy loading
	Placeholder string `json:"placeholder,omitempty"` // Placeholder image
}

// ProgressConfig configures progress bar rendering
type ProgressConfig struct {
	Min             float64          `json:"min"`                               // Minimum value
	Max             float64          `json:"max"`                               // Maximum value
	ShowValue       bool             `json:"showValue"`                         // Show percentage
	ColorThresholds []ColorThreshold `json:"colorThresholds,omitempty"`
}

// ColorThreshold defines color changes at specific values
type ColorThreshold struct {
	Threshold float64 `json:"threshold"`
	Color     string  `json:"color"`
	Variant   string  `json:"variant,omitempty"`
}

// FilterOption represents an option in a filter dropdown
type FilterOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Icon  string `json:"icon,omitempty"`
	Color string `json:"color,omitempty"`
}

// ActionConfig configures column/row actions
type ActionConfig struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	Icon       string `json:"icon,omitempty"`
	Variant    string `json:"variant,omitempty"`
	Action     string `json:"action"` // htmx, alpine, javascript
	URL        string `json:"url,omitempty"`
	Confirm    bool   `json:"confirm"`
	Message    string `json:"message,omitempty"`
	Condition  string `json:"condition,omitempty"`
	Permission string `json:"permission,omitempty"`
}

// SchemaFilter represents a filter derived from schema
// Note: Schema package implementation pending
type SchemaFilter struct {
	// Field        schema.Field     `json:"field"`
	Type         FilterType       `json:"type"`
	Operators    []FilterOperator `json:"operators"`
	Options      []FilterOption   `json:"options,omitempty"`
	DefaultValue any              `json:"defaultValue,omitempty"`
	QuickFilter  bool             `json:"quickFilter"`
}

// FilterType represents different filter UI types
type FilterType string

const (
	FilterTypeText        FilterType = "text"
	FilterTypeSelect      FilterType = "select"
	FilterTypeMultiSelect FilterType = "multiselect"
	FilterTypeDate        FilterType = "date"
	FilterTypeDateRange   FilterType = "daterange"
	FilterTypeNumber      FilterType = "number"
	FilterTypeNumberRange FilterType = "numberrange"
	FilterTypeBoolean     FilterType = "boolean"
)

// SchemaAction represents an action derived from schema
// Note: Schema package implementation pending
type SchemaAction struct {
	ID          string     `json:"id"`
	Label       string     `json:"label"`
	Icon        string     `json:"icon,omitempty"`
	Type        ActionType `json:"type"` // table, bulk, row
	// Action      schema.Action `json:"action"`
	Condition   string `json:"condition,omitempty"`
	Permission  string `json:"permission,omitempty"`
	Confirm     bool   `json:"confirm"`
	Destructive bool   `json:"destructive"`
}

// ActionType represents different action scopes
type ActionType string

const (
	ActionTypeTable ActionType = "table"
	ActionTypeBulk  ActionType = "bulk"
	ActionTypeRow   ActionType = "row"
)

// ShowCondition represents a condition for showing/hiding elements
type ShowCondition struct {
	Condition string `json:"condition"` // JavaScript expression
	Target    string `json:"target"`    // element to show/hide
	Action    string `json:"action"`    // show, hide, enable, disable
}

// NewDataTableSchemaBuilder creates a new schema-driven table builder
// Note: Schema package implementation pending
func NewDataTableSchemaBuilder() *DataTableSchemaBuilder {
	return &DataTableSchemaBuilder{
		// schema:        tableSchema,
		config:        getDefaultDataTableConfig(),
		fieldMappings: make(map[string]ColumnMapping),
		filters:       []SchemaFilter{},
		actions:       []SchemaAction{},
	}
}

// getDefaultDataTableConfig returns default configuration
func getDefaultDataTableConfig() *DataTableConfig {
	return &DataTableConfig{
		ShowTitle:         true,
		ShowDescription:   true,
		EnableSelection:   true,
		EnableMultiSelect: true,
		EnableSorting:     true,
		EnableFiltering:   true,
		EnableSearch:      true,
		EnableExport:      true,
		EnablePagination:  true,
		DefaultPageSize:   25,
		PageSizeOptions:   []int{10, 25, 50, 100},
		SearchPlaceholder: "Search...",
		ExportFormats:     []ExportFormat{ExportCSV, ExportExcel, ExportPDF},
		Variant:           DataTableDefault,
		Size:              DataTableSizeMD,
		Density:           DataTableDensityComfortable,
		ColumnDefaults: ColumnDefaults{
			Visible:    true,
			Sortable:   true,
			Filterable: true,
			Searchable: true,
			Resizable:  true,
			Align:      "left",
		},
		EnableDefaultActions: true,
		ActionPosition:       "right",
		EnableColumnResize:   true,
		EnableColumnReorder:  true,
		VirtualScrolling:     false,
		LazyLoading:          false,
		CacheData:            true,
		AriaLabels: map[string]string{
			"table":     "Data table",
			"search":    "Search table",
			"filter":    "Filter table",
			"sort":      "Sort column",
			"select":    "Select row",
			"selectAll": "Select all rows",
		},
	}
}