package organisms

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/niiniyare/ruun/pkg/schema"
	"github.com/niiniyare/ruun/views/components/atoms"
	"github.com/niiniyare/ruun/views/components/molecules"
)

// DataTableService provides business logic for data table operations
type DataTableService struct {
	// Configuration
	DefaultPageSize   int
	MaxPageSize      int
	DefaultSortOrder SortDirection
	
	// Feature flags
	EnableVirtualization bool
	EnableCaching        bool
	CacheTimeout        time.Duration
	
	// Performance settings
	MaxRows             int
	ChunkSize           int
	SearchDebounce      time.Duration
}

// NewDataTableService creates a new service with default configuration
func NewDataTableService() *DataTableService {
	return &DataTableService{
		DefaultPageSize:     25,
		MaxPageSize:        1000,
		DefaultSortOrder:   SortAsc,
		EnableVirtualization: true,
		EnableCaching:      true,
		CacheTimeout:       5 * time.Minute,
		MaxRows:           10000,
		ChunkSize:         100,
		SearchDebounce:    300 * time.Millisecond,
	}
}

// DataTableQuery represents a query for data table operations
type DataTableQuery struct {
	// Pagination
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	
	// Sorting
	SortBy    string             `json:"sortBy"`
	SortOrder SortDirection      `json:"sortOrder"`
	MultiSort []SortCriteria     `json:"multiSort,omitempty"`
	
	// Search
	Search       string           `json:"search"`
	SearchFields []string         `json:"searchFields,omitempty"`
	
	// Filtering
	Filters      []FilterCriteria `json:"filters"`
	QuickFilters []string         `json:"quickFilters,omitempty"`
	
	// Selection
	SelectedRows []string         `json:"selectedRows,omitempty"`
	
	// Display options
	VisibleColumns []string       `json:"visibleColumns,omitempty"`
	ColumnWidths   map[string]int `json:"columnWidths,omitempty"`
	
	// Advanced options
	GroupBy      []string         `json:"groupBy,omitempty"`
	Aggregates   []AggregateCriteria `json:"aggregates,omitempty"`
	
	// Meta
	RequestID    string          `json:"requestId,omitempty"`
	UserID       string          `json:"userId,omitempty"`
	SessionID    string          `json:"sessionId,omitempty"`
}

// SortCriteria represents sorting criteria for multi-column sorting
type SortCriteria struct {
	Column    string        `json:"column"`
	Direction SortDirection `json:"direction"`
	Priority  int           `json:"priority"`
}

// FilterCriteria represents filtering criteria
type FilterCriteria struct {
	Column   string           `json:"column"`
	Operator FilterOperator   `json:"operator"`
	Value    any              `json:"value"`
	Values   []any            `json:"values,omitempty"`
	Type     string           `json:"type,omitempty"`
}

// AggregateCriteria represents aggregation criteria
type AggregateCriteria struct {
	Column   string `json:"column"`
	Function string `json:"function"` // sum, avg, count, min, max
	Label    string `json:"label,omitempty"`
}

// DataTableResponse represents the response from data table operations
type DataTableResponse struct {
	// Data
	Rows         []DataTableRow    `json:"rows"`
	TotalRows    int               `json:"totalRows"`
	FilteredRows int               `json:"filteredRows"`
	
	// Pagination
	Page         int               `json:"page"`
	PageSize     int               `json:"pageSize"`
	TotalPages   int               `json:"totalPages"`
	HasNext      bool              `json:"hasNext"`
	HasPrev      bool              `json:"hasPrev"`
	
	// Aggregations
	Aggregates   map[string]any    `json:"aggregates,omitempty"`
	
	// Performance
	QueryTime    time.Duration     `json:"queryTime"`
	CacheHit     bool              `json:"cacheHit"`
	
	// Meta
	RequestID    string            `json:"requestId"`
	Timestamp    time.Time         `json:"timestamp"`
	Warnings     []string          `json:"warnings,omitempty"`
}

// ProcessQuery processes a data table query and returns results
func (s *DataTableService) ProcessQuery(ctx context.Context, data []DataTableRow, query DataTableQuery) (*DataTableResponse, error) {
	startTime := time.Now()
	
	// Validate query
	if err := s.validateQuery(query); err != nil {
		return nil, fmt.Errorf("invalid query: %w", err)
	}
	
	// Apply search
	filteredData, err := s.applySearch(data, query.Search, query.SearchFields)
	if err != nil {
		return nil, fmt.Errorf("search failed: %w", err)
	}
	
	// Apply filters
	filteredData, err = s.applyFilters(filteredData, query.Filters)
	if err != nil {
		return nil, fmt.Errorf("filtering failed: %w", err)
	}
	
	// Apply sorting
	if query.SortBy != "" {
		filteredData = s.applySorting(filteredData, query.SortBy, query.SortOrder)
	} else if len(query.MultiSort) > 0 {
		filteredData = s.applyMultiSorting(filteredData, query.MultiSort)
	}
	
	// Calculate aggregates
	aggregates := s.calculateAggregates(filteredData, query.Aggregates)
	
	// Apply pagination
	paginatedData, totalPages := s.applyPagination(filteredData, query.Page, query.PageSize)
	
	queryTime := time.Since(startTime)
	
	response := &DataTableResponse{
		Rows:         paginatedData,
		TotalRows:    len(data),
		FilteredRows: len(filteredData),
		Page:         query.Page,
		PageSize:     query.PageSize,
		TotalPages:   totalPages,
		HasNext:      query.Page < totalPages,
		HasPrev:      query.Page > 1,
		Aggregates:   aggregates,
		QueryTime:    queryTime,
		RequestID:    query.RequestID,
		Timestamp:    time.Now(),
	}
	
	return response, nil
}

// validateQuery validates the query parameters
func (s *DataTableService) validateQuery(query DataTableQuery) error {
	if query.Page < 1 {
		return fmt.Errorf("page must be >= 1")
	}
	
	if query.PageSize < 1 {
		return fmt.Errorf("pageSize must be >= 1")
	}
	
	if query.PageSize > s.MaxPageSize {
		return fmt.Errorf("pageSize must be <= %d", s.MaxPageSize)
	}
	
	if query.SortOrder != "" && query.SortOrder != SortAsc && query.SortOrder != SortDesc {
		return fmt.Errorf("invalid sort order: %s", query.SortOrder)
	}
	
	// Validate filter operators
	for _, filter := range query.Filters {
		if !s.isValidFilterOperator(filter.Operator) {
			return fmt.Errorf("invalid filter operator: %s", filter.Operator)
		}
	}
	
	return nil
}

// isValidFilterOperator checks if a filter operator is valid
func (s *DataTableService) isValidFilterOperator(op FilterOperator) bool {
	validOperators := []FilterOperator{
		FilterEquals, FilterNotEquals, FilterContains, FilterNotContains,
		FilterStartsWith, FilterEndsWith, FilterGreaterThan, FilterGreaterEqual,
		FilterLessThan, FilterLessEqual, FilterBetween, FilterIsNull,
		FilterIsNotNull, FilterIn, FilterNotIn,
	}
	
	for _, valid := range validOperators {
		if op == valid {
			return true
		}
	}
	return false
}

// applySearch applies search filtering to the data
func (s *DataTableService) applySearch(data []DataTableRow, searchQuery string, searchFields []string) ([]DataTableRow, error) {
	if searchQuery == "" {
		return data, nil
	}
	
	searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))
	if len(searchQuery) < 2 { // Minimum search length
		return data, nil
	}
	
	var results []DataTableRow
	
	for _, row := range data {
		if s.rowMatchesSearch(row, searchQuery, searchFields) {
			results = append(results, row)
		}
	}
	
	return results, nil
}

// rowMatchesSearch checks if a row matches the search query
func (s *DataTableService) rowMatchesSearch(row DataTableRow, searchQuery string, searchFields []string) bool {
	// If no specific fields specified, search all fields
	if len(searchFields) == 0 {
		for _, value := range row.Data {
			if s.valueMatchesSearch(value, searchQuery) {
				return true
			}
		}
		return false
	}
	
	// Search specific fields
	for _, field := range searchFields {
		if value, exists := row.Data[field]; exists {
			if s.valueMatchesSearch(value, searchQuery) {
				return true
			}
		}
	}
	
	return false
}

// valueMatchesSearch checks if a value matches the search query
func (s *DataTableService) valueMatchesSearch(value any, searchQuery string) bool {
	if value == nil {
		return false
	}
	
	valueStr := strings.ToLower(fmt.Sprintf("%v", value))
	return strings.Contains(valueStr, searchQuery)
}

// applyFilters applies filters to the data
func (s *DataTableService) applyFilters(data []DataTableRow, filters []FilterCriteria) ([]DataTableRow, error) {
	if len(filters) == 0 {
		return data, nil
	}
	
	var results []DataTableRow
	
	for _, row := range data {
		matches := true
		
		for _, filter := range filters {
			if !s.rowMatchesFilter(row, filter) {
				matches = false
				break
			}
		}
		
		if matches {
			results = append(results, row)
		}
	}
	
	return results, nil
}

// rowMatchesFilter checks if a row matches a filter
func (s *DataTableService) rowMatchesFilter(row DataTableRow, filter FilterCriteria) bool {
	value, exists := row.Data[filter.Column]
	if !exists {
		return filter.Operator == FilterIsNull
	}
	
	switch filter.Operator {
	case FilterEquals:
		return s.compareValues(value, filter.Value) == 0
		
	case FilterNotEquals:
		return s.compareValues(value, filter.Value) != 0
		
	case FilterContains:
		return s.stringContains(value, filter.Value)
		
	case FilterNotContains:
		return !s.stringContains(value, filter.Value)
		
	case FilterStartsWith:
		return s.stringStartsWith(value, filter.Value)
		
	case FilterEndsWith:
		return s.stringEndsWith(value, filter.Value)
		
	case FilterGreaterThan:
		return s.compareValues(value, filter.Value) > 0
		
	case FilterGreaterEqual:
		return s.compareValues(value, filter.Value) >= 0
		
	case FilterLessThan:
		return s.compareValues(value, filter.Value) < 0
		
	case FilterLessEqual:
		return s.compareValues(value, filter.Value) <= 0
		
	case FilterBetween:
		if len(filter.Values) != 2 {
			return false
		}
		return s.compareValues(value, filter.Values[0]) >= 0 &&
			   s.compareValues(value, filter.Values[1]) <= 0
			   
	case FilterIsNull:
		return value == nil
		
	case FilterIsNotNull:
		return value != nil
		
	case FilterIn:
		return s.valueInSlice(value, filter.Values)
		
	case FilterNotIn:
		return !s.valueInSlice(value, filter.Values)
		
	default:
		return false
	}
}

// compareValues compares two values and returns -1, 0, or 1
func (s *DataTableService) compareValues(a, b any) int {
	if a == nil && b == nil {
		return 0
	}
	if a == nil {
		return -1
	}
	if b == nil {
		return 1
	}
	
	// Try to compare as strings first
	aStr := fmt.Sprintf("%v", a)
	bStr := fmt.Sprintf("%v", b)
	
	// Try to parse as numbers
	if aFloat, errA := strconv.ParseFloat(aStr, 64); errA == nil {
		if bFloat, errB := strconv.ParseFloat(bStr, 64); errB == nil {
			if aFloat < bFloat {
				return -1
			} else if aFloat > bFloat {
				return 1
			}
			return 0
		}
	}
	
	// Try to parse as time
	if aTime, errA := time.Parse(time.RFC3339, aStr); errA == nil {
		if bTime, errB := time.Parse(time.RFC3339, bStr); errB == nil {
			if aTime.Before(bTime) {
				return -1
			} else if aTime.After(bTime) {
				return 1
			}
			return 0
		}
	}
	
	// Fall back to string comparison
	return strings.Compare(aStr, bStr)
}

// stringContains checks if a value contains a substring
func (s *DataTableService) stringContains(value, substring any) bool {
	valueStr := strings.ToLower(fmt.Sprintf("%v", value))
	substringStr := strings.ToLower(fmt.Sprintf("%v", substring))
	return strings.Contains(valueStr, substringStr)
}

// stringStartsWith checks if a value starts with a prefix
func (s *DataTableService) stringStartsWith(value, prefix any) bool {
	valueStr := strings.ToLower(fmt.Sprintf("%v", value))
	prefixStr := strings.ToLower(fmt.Sprintf("%v", prefix))
	return strings.HasPrefix(valueStr, prefixStr)
}

// stringEndsWith checks if a value ends with a suffix
func (s *DataTableService) stringEndsWith(value, suffix any) bool {
	valueStr := strings.ToLower(fmt.Sprintf("%v", value))
	suffixStr := strings.ToLower(fmt.Sprintf("%v", suffix))
	return strings.HasSuffix(valueStr, suffixStr)
}

// valueInSlice checks if a value is in a slice of values
func (s *DataTableService) valueInSlice(value any, values []any) bool {
	for _, v := range values {
		if s.compareValues(value, v) == 0 {
			return true
		}
	}
	return false
}

// applySorting applies single-column sorting to the data
func (s *DataTableService) applySorting(data []DataTableRow, sortBy string, sortOrder SortDirection) []DataTableRow {
	if sortBy == "" {
		return data
	}
	
	// Make a copy to avoid modifying the original slice
	sorted := make([]DataTableRow, len(data))
	copy(sorted, data)
	
	sort.Slice(sorted, func(i, j int) bool {
		valueI := sorted[i].Data[sortBy]
		valueJ := sorted[j].Data[sortBy]
		
		comparison := s.compareValues(valueI, valueJ)
		
		if sortOrder == SortDesc {
			return comparison > 0
		}
		return comparison < 0
	})
	
	return sorted
}

// applyMultiSorting applies multi-column sorting to the data
func (s *DataTableService) applyMultiSorting(data []DataTableRow, criteria []SortCriteria) []DataTableRow {
	if len(criteria) == 0 {
		return data
	}
	
	// Sort criteria by priority
	sort.Slice(criteria, func(i, j int) bool {
		return criteria[i].Priority < criteria[j].Priority
	})
	
	// Make a copy to avoid modifying the original slice
	sorted := make([]DataTableRow, len(data))
	copy(sorted, data)
	
	sort.Slice(sorted, func(i, j int) bool {
		for _, criterion := range criteria {
			valueI := sorted[i].Data[criterion.Column]
			valueJ := sorted[j].Data[criterion.Column]
			
			comparison := s.compareValues(valueI, valueJ)
			
			if comparison != 0 {
				if criterion.Direction == SortDesc {
					return comparison > 0
				}
				return comparison < 0
			}
		}
		return false // All criteria are equal
	})
	
	return sorted
}

// calculateAggregates calculates aggregates for the data
func (s *DataTableService) calculateAggregates(data []DataTableRow, aggregates []AggregateCriteria) map[string]any {
	results := make(map[string]any)
	
	for _, agg := range aggregates {
		key := agg.Column + "_" + agg.Function
		if agg.Label != "" {
			key = agg.Label
		}
		
		switch agg.Function {
		case "count":
			results[key] = len(data)
			
		case "sum":
			sum := 0.0
			count := 0
			for _, row := range data {
				if value, exists := row.Data[agg.Column]; exists {
					if floatVal, err := s.toFloat(value); err == nil {
						sum += floatVal
						count++
					}
				}
			}
			results[key] = sum
			
		case "avg":
			sum := 0.0
			count := 0
			for _, row := range data {
				if value, exists := row.Data[agg.Column]; exists {
					if floatVal, err := s.toFloat(value); err == nil {
						sum += floatVal
						count++
					}
				}
			}
			if count > 0 {
				results[key] = sum / float64(count)
			} else {
				results[key] = 0
			}
			
		case "min":
			var min *float64
			for _, row := range data {
				if value, exists := row.Data[agg.Column]; exists {
					if floatVal, err := s.toFloat(value); err == nil {
						if min == nil || floatVal < *min {
							min = &floatVal
						}
					}
				}
			}
			if min != nil {
				results[key] = *min
			}
			
		case "max":
			var max *float64
			for _, row := range data {
				if value, exists := row.Data[agg.Column]; exists {
					if floatVal, err := s.toFloat(value); err == nil {
						if max == nil || floatVal > *max {
							max = &floatVal
						}
					}
				}
			}
			if max != nil {
				results[key] = *max
			}
		}
	}
	
	return results
}

// toFloat converts a value to float64
func (s *DataTableService) toFloat(value any) (float64, error) {
	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return strconv.ParseFloat(fmt.Sprintf("%v", v), 64)
	}
}

// applyPagination applies pagination to the data
func (s *DataTableService) applyPagination(data []DataTableRow, page, pageSize int) ([]DataTableRow, int) {
	totalRows := len(data)
	totalPages := (totalRows + pageSize - 1) / pageSize
	
	if totalPages == 0 {
		return []DataTableRow{}, 0
	}
	
	if page > totalPages {
		page = totalPages
	}
	
	start := (page - 1) * pageSize
	end := start + pageSize
	
	if end > totalRows {
		end = totalRows
	}
	
	return data[start:end], totalPages
}

// BuildDataTableFromSchema creates a DataTable configuration from a schema
func BuildDataTableFromSchema(ctx context.Context, tableSchema *schema.Schema) (*DataTableProps, error) {
	if tableSchema == nil {
		return nil, fmt.Errorf("schema cannot be nil")
	}
	
	props := &DataTableProps{
		ID:          tableSchema.ID,
		Title:       tableSchema.Title,
		Description: tableSchema.Description,
		Columns:     []DataTableColumn{},
		Rows:        []DataTableRow{},
		
		// Default configuration
		Selectable:  true,
		MultiSelect: true,
		Sortable:    true,
		Filterable:  true,
		
		// Search configuration
		Search: DataTableSearch{
			Enabled:     true,
			Placeholder: "Search " + tableSchema.Title + "...",
			MinLength:   2,
			Delay:       300,
		},
		
		// Pagination configuration
		Pagination: DataTablePagination{
			Enabled:         true,
			CurrentPage:     1,
			PageSize:        25,
			PageSizeOptions: []int{10, 25, 50, 100},
			ShowTotal:       true,
			ShowPageSize:    true,
		},
		
		// Export configuration
		Export: DataTableExport{
			Enabled: true,
			Formats: []ExportFormat{ExportCSV, ExportExcel, ExportPDF},
		},
	}
	
	// Build columns from schema fields
	for _, field := range tableSchema.Fields {
		column := buildColumnFromField(field)
		props.Columns = append(props.Columns, column)
	}
	
	// Set up default actions
	props.Actions = buildDefaultActions()
	props.BulkActions = buildDefaultBulkActions()
	props.RowActions = buildDefaultRowActions()
	
	return props, nil
}

// buildColumnFromField converts a schema field to a data table column
func buildColumnFromField(field schema.Field) DataTableColumn {
	column := DataTableColumn{
		Key:          field.Name,
		Title:        getFieldDisplayName(field),
		Visible:      !field.Hidden,
		Sortable:     field.Sortable,
		Searchable:   field.Searchable,
		Filterable:   true,
		SchemaField:  &field,
	}
	
	// Map field types to column types
	switch field.Type {
	case schema.FieldText, schema.FieldTextarea:
		column.Type = ColumnTypeText
	case schema.FieldNumber:
		column.Type = ColumnTypeNumber
	case schema.FieldEmail:
		column.Type = ColumnTypeLink
	case schema.FieldDate:
		column.Type = ColumnTypeDate
	case schema.FieldDatetime:
		column.Type = ColumnTypeDateTime
	case schema.FieldSelect, schema.FieldRadio:
		column.Type = ColumnTypeBadge
		column.BadgeMap = buildBadgeMapFromOptions(field.Options)
	case schema.FieldCheckbox:
		column.Type = ColumnTypeCheckbox
	case schema.FieldFile:
		if isImageField(field) {
			column.Type = ColumnTypeImage
		} else {
			column.Type = ColumnTypeLink
		}
	default:
		column.Type = ColumnTypeText
	}
	
	// Set column alignment based on type
	switch column.Type {
	case ColumnTypeNumber, ColumnTypeCurrency, ColumnTypePercent:
		column.Align = "right"
	case ColumnTypeDate, ColumnTypeDateTime:
		column.Align = "center"
	default:
		column.Align = "left"
	}
	
	return column
}

// getFieldDisplayName returns the display name for a field
func getFieldDisplayName(field schema.Field) string {
	if field.Label != "" {
		return field.Label
	}
	return strings.Title(strings.ReplaceAll(field.Name, "_", " "))
}

// buildBadgeMapFromOptions creates a badge variant map from field options
func buildBadgeMapFromOptions(options []schema.FieldOption) map[string]atoms.BadgeVariant {
	badgeMap := make(map[string]atoms.BadgeVariant)
	
	for i, option := range options {
		var variant atoms.BadgeVariant
		
		// Assign variants in a cycle
		switch i % 4 {
		case 0:
			variant = atoms.BadgePrimary
		case 1:
			variant = atoms.BadgeSecondary
		case 2:
			variant = atoms.BadgeSuccess
		case 3:
			variant = atoms.BadgeWarning
		}
		
		badgeMap[option.Value] = variant
	}
	
	return badgeMap
}

// isImageField checks if a file field is for images
func isImageField(field schema.Field) bool {
	// Check if the field has image-specific validation or metadata
	if field.Validation != nil {
		// This is a simplified check - in practice you'd check MIME types, extensions, etc.
		return strings.Contains(strings.ToLower(field.Name), "image") ||
			   strings.Contains(strings.ToLower(field.Name), "photo") ||
			   strings.Contains(strings.ToLower(field.Name), "avatar")
	}
	return false
}

// buildDefaultActions creates default table actions
func buildDefaultActions() []DataTableAction {
	return []DataTableAction{
		{
			ID:      "add",
			Text:    "Add New",
			Icon:    "plus",
			Variant: atoms.ButtonPrimary,
			Type:    "button",
		},
		{
			ID:      "refresh",
			Text:    "Refresh",
			Icon:    "refresh-cw",
			Variant: atoms.ButtonOutline,
			Type:    "button",
		},
	}
}

// buildDefaultBulkActions creates default bulk actions
func buildDefaultBulkActions() []DataTableBulkAction {
	return []DataTableBulkAction{
		{
			ID:          "delete",
			Text:        "Delete Selected",
			Icon:        "trash-2",
			Variant:     atoms.ButtonDestructive,
			Destructive: true,
			Confirm:     true,
			Message:     "Are you sure you want to delete the selected items?",
		},
		{
			ID:      "export",
			Text:    "Export Selected",
			Icon:    "download",
			Variant: atoms.ButtonOutline,
		},
	}
}

// buildDefaultRowActions creates default row actions
func buildDefaultRowActions() []DataTableAction {
	return []DataTableAction{
		{
			ID:   "view",
			Text: "View",
			Icon: "eye",
		},
		{
			ID:   "edit",
			Text: "Edit",
			Icon: "edit",
		},
		{
			ID:   "delete",
			Text: "Delete",
			Icon: "trash-2",
		},
	}
}

// DataTableBuilder provides a fluent interface for building data tables
type DataTableBuilder struct {
	props *DataTableProps
}

// NewDataTableBuilder creates a new data table builder
func NewDataTableBuilder(id, title string) *DataTableBuilder {
	return &DataTableBuilder{
		props: &DataTableProps{
			ID:          id,
			Title:       title,
			Variant:     DataTableDefault,
			Size:        DataTableSizeMD,
			Density:     DataTableDensityComfortable,
			Columns:     []DataTableColumn{},
			Rows:        []DataTableRow{},
			Selectable:  true,
			MultiSelect: true,
			Sortable:    true,
			Filterable:  true,
			Search: DataTableSearch{
				Enabled: true,
			},
			Pagination: DataTablePagination{
				Enabled:         true,
				CurrentPage:     1,
				PageSize:        25,
				PageSizeOptions: []int{10, 25, 50, 100},
				ShowTotal:       true,
				ShowPageSize:    true,
			},
		},
	}
}

// WithDescription sets the table description
func (b *DataTableBuilder) WithDescription(description string) *DataTableBuilder {
	b.props.Description = description
	return b
}

// WithVariant sets the table variant
func (b *DataTableBuilder) WithVariant(variant DataTableVariant) *DataTableBuilder {
	b.props.Variant = variant
	return b
}

// WithSize sets the table size
func (b *DataTableBuilder) WithSize(size DataTableSize) *DataTableBuilder {
	b.props.Size = size
	return b
}

// WithColumns sets the table columns
func (b *DataTableBuilder) WithColumns(columns []DataTableColumn) *DataTableBuilder {
	b.props.Columns = columns
	return b
}

// AddColumn adds a column to the table
func (b *DataTableBuilder) AddColumn(column DataTableColumn) *DataTableBuilder {
	b.props.Columns = append(b.props.Columns, column)
	return b
}

// WithRows sets the table rows
func (b *DataTableBuilder) WithRows(rows []DataTableRow) *DataTableBuilder {
	b.props.Rows = rows
	return b
}

// WithPagination configures pagination
func (b *DataTableBuilder) WithPagination(pagination DataTablePagination) *DataTableBuilder {
	b.props.Pagination = pagination
	return b
}

// WithSearch configures search
func (b *DataTableBuilder) WithSearch(search DataTableSearch) *DataTableBuilder {
	b.props.Search = search
	return b
}

// WithExport configures export
func (b *DataTableBuilder) WithExport(export DataTableExport) *DataTableBuilder {
	b.props.Export = export
	return b
}

// WithActions sets table actions
func (b *DataTableBuilder) WithActions(actions []DataTableAction) *DataTableBuilder {
	b.props.Actions = actions
	return b
}

// WithBulkActions sets bulk actions
func (b *DataTableBuilder) WithBulkActions(actions []DataTableBulkAction) *DataTableBuilder {
	b.props.BulkActions = actions
	return b
}

// WithRowActions sets row actions
func (b *DataTableBuilder) WithRowActions(actions []DataTableAction) *DataTableBuilder {
	b.props.RowActions = actions
	return b
}

// Build returns the built DataTableProps
func (b *DataTableBuilder) Build() *DataTableProps {
	return b.props
}

// Convenience functions for common column types

// TextColumn creates a text column
func TextColumn(key, title string) DataTableColumn {
	return DataTableColumn{
		Key:        key,
		Title:      title,
		Type:       ColumnTypeText,
		Visible:    true,
		Sortable:   true,
		Searchable: true,
		Filterable: true,
		Align:      "left",
	}
}

// NumberColumn creates a number column
func NumberColumn(key, title string, precision int) DataTableColumn {
	return DataTableColumn{
		Key:        key,
		Title:      title,
		Type:       ColumnTypeNumber,
		Precision:  precision,
		Visible:    true,
		Sortable:   true,
		Filterable: true,
		Align:      "right",
	}
}

// DateColumn creates a date column
func DateColumn(key, title string, format string) DataTableColumn {
	return DataTableColumn{
		Key:        key,
		Title:      title,
		Type:       ColumnTypeDate,
		DateFormat: format,
		Visible:    true,
		Sortable:   true,
		Filterable: true,
		Align:      "center",
	}
}

// BadgeColumn creates a badge column
func BadgeColumn(key, title string, badgeMap map[string]atoms.BadgeVariant) DataTableColumn {
	return DataTableColumn{
		Key:        key,
		Title:      title,
		Type:       ColumnTypeBadge,
		BadgeMap:   badgeMap,
		Visible:    true,
		Sortable:   true,
		Filterable: true,
		Align:      "center",
	}
}

// ActionsColumn creates an actions column
func ActionsColumn(actions []molecules.MenuItemProps) DataTableColumn {
	return DataTableColumn{
		Key:         "actions",
		Title:       "Actions",
		Type:        ColumnTypeActions,
		ActionItems: actions,
		Visible:     true,
		Sortable:    false,
		Filterable:  false,
		Align:       "right",
		Width:       "120px",
	}
}