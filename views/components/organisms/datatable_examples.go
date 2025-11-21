package organisms

//
// import (
// 	"fmt"
// 	"time"
//
// 	"github.com/niiniyare/ruun/pkg/schema"
// 	"github.com/niiniyare/ruun/views/components/atoms"
// 	"github.com/niiniyare/ruun/views/components/molecules"
// )
//
// // DataTableExamples provides comprehensive usage examples for the DataTable organism
// type DataTableExamples struct{}
//
// // Example 1: Basic DataTable with minimal configuration
// func (e *DataTableExamples) BasicExample() *DataTableProps {
// 	return &DataTableProps{
// 		ID:          "basic-table",
// 		Title:       "Users",
// 		Description: "Manage system users",
// 		Columns: []DataTableColumn{
// 			TextColumn("name", "Name"),
// 			TextColumn("email", "Email"),
// 			DateColumn("created_at", "Created", "2006-01-02"),
// 		},
// 		Rows: []DataTableRow{
// 			{
// 				ID: "1",
// 				Data: map[string]any{
// 					"name":       "John Doe",
// 					"email":      "john@example.com",
// 					"created_at": time.Now().AddDate(0, -1, 0),
// 				},
// 			},
// 			{
// 				ID: "2",
// 				Data: map[string]any{
// 					"name":       "Jane Smith",
// 					"email":      "jane@example.com",
// 					"created_at": time.Now().AddDate(0, -2, 0),
// 				},
// 			},
// 		},
// 		Selectable: true,
// 		Sortable:   true,
// 		Search: DataTableSearch{
// 			Enabled:     true,
// 			Placeholder: "Search users...",
// 		},
// 		Pagination: DataTablePagination{
// 			Enabled:     true,
// 			CurrentPage: 1,
// 			PageSize:    10,
// 			TotalItems:  2,
// 			TotalPages:  1,
// 		},
// 	}
// }
//
// // Example 2: Advanced DataTable with all features
// func (e *DataTableExamples) AdvancedExample() *DataTableProps {
// 	return &DataTableProps{
// 		ID:          "advanced-table",
// 		Title:       "Product Inventory",
// 		Description: "Comprehensive product management with advanced features",
// 		Variant:     DataTableStriped,
// 		Size:        DataTableSizeLG,
// 		Density:     DataTableDensityCompact,
// 		Selectable:  true,
// 		MultiSelect: true,
// 		Sortable:    true,
// 		Resizable:   true,
// 		Filterable:  true,
// 		Expandable:  true,
//
// 		Columns: []DataTableColumn{
// 			{
// 				Key:        "product_id",
// 				Title:      "ID",
// 				Type:       ColumnTypeNumber,
// 				Width:      "80px",
// 				Sortable:   true,
// 				Filterable: true,
// 				Align:      "center",
// 			},
// 			{
// 				Key:          "name",
// 				Title:        "Product Name",
// 				Type:         ColumnTypeText,
// 				Sortable:     true,
// 				Searchable:   true,
// 				Filterable:   true,
// 				Clickable:    true,
// 				ClickHandler: "viewProduct",
// 			},
// 			{
// 				Key:      "category",
// 				Title:    "Category",
// 				Type:     ColumnTypeBadge,
// 				Sortable: true,
// 				BadgeMap: map[string]atoms.BadgeVariant{
// 					"Electronics": atoms.BadgePrimary,
// 					"Clothing":    atoms.BadgeSecondary,
// 					"Books":       atoms.BadgeSuccess,
// 					"Home":        atoms.BadgeWarning,
// 				},
// 			},
// 			{
// 				Key:          "price",
// 				Title:        "Price",
// 				Type:         ColumnTypeCurrency,
// 				Sortable:     true,
// 				Filterable:   true,
// 				Align:        "right",
// 				CurrencyCode: "USD",
// 				Precision:    2,
// 			},
// 			{
// 				Key:      "stock",
// 				Title:    "Stock",
// 				Type:     ColumnTypeProgress,
// 				Sortable: true,
// 				Align:    "center",
// 			},
// 			{
// 				Key:        "status",
// 				Title:      "Status",
// 				Type:       ColumnTypeBadge,
// 				Sortable:   true,
// 				Filterable: true,
// 				BadgeMap: map[string]atoms.BadgeVariant{
// 					"active":       atoms.BadgeSuccess,
// 					"inactive":     atoms.BadgeSecondary,
// 					"out_of_stock": atoms.BadgeDestructive,
// 				},
// 			},
// 			{
// 				Key:        "image",
// 				Title:      "Image",
// 				Type:       ColumnTypeImage,
// 				Width:      "60px",
// 				Sortable:   false,
// 				Filterable: false,
// 			},
// 			ActionsColumn([]molecules.MenuItemProps{
// 				{Text: "Edit", Icon: "edit", Action: "edit"},
// 				{Text: "Duplicate", Icon: "copy", Action: "duplicate"},
// 				{Text: "Delete", Icon: "trash", Action: "delete", Destructive: true},
// 			}),
// 		},
//
// 		Rows: []DataTableRow{
// 			{
// 				ID: "prod-001",
// 				Data: map[string]any{
// 					"product_id": 1,
// 					"name":       "Wireless Headphones",
// 					"category":   "Electronics",
// 					"price":      299.99,
// 					"stock":      75,
// 					"status":     "active",
// 					"image":      "/images/headphones.jpg",
// 				},
// 				Actions: []molecules.MenuItemProps{
// 					{Text: "View Details", Icon: "eye", Action: "view"},
// 				},
// 			},
// 			{
// 				ID: "prod-002",
// 				Data: map[string]any{
// 					"product_id": 2,
// 					"name":       "Cotton T-Shirt",
// 					"category":   "Clothing",
// 					"price":      29.99,
// 					"stock":      45,
// 					"status":     "active",
// 					"image":      "/images/tshirt.jpg",
// 				},
// 			},
// 			{
// 				ID: "prod-003",
// 				Data: map[string]any{
// 					"product_id": 3,
// 					"name":       "Programming Book",
// 					"category":   "Books",
// 					"price":      49.99,
// 					"stock":      0,
// 					"status":     "out_of_stock",
// 					"image":      "/images/book.jpg",
// 				},
// 				Variant: "warning", // Highlight out of stock items
// 			},
// 		},
//
// 		Search: DataTableSearch{
// 			Enabled:       true,
// 			Placeholder:   "Search products...",
// 			Columns:       []string{"name", "category"},
// 			CaseSensitive: false,
// 			Highlight:     true,
// 			Advanced:      true,
// 		},
//
// 		Filters: []DataTableFilter{
// 			{
// 				Key:    "category",
// 				Label:  "Category",
// 				Type:   "select",
// 				Active: false,
// 				Options: []molecules.SelectOption{
// 					{Value: "Electronics", Label: "Electronics"},
// 					{Value: "Clothing", Label: "Clothing"},
// 					{Value: "Books", Label: "Books"},
// 					{Value: "Home", Label: "Home & Garden"},
// 				},
// 			},
// 			{
// 				Key:    "status",
// 				Label:  "Status",
// 				Type:   "select",
// 				Active: false,
// 				Options: []molecules.SelectOption{
// 					{Value: "active", Label: "Active"},
// 					{Value: "inactive", Label: "Inactive"},
// 					{Value: "out_of_stock", Label: "Out of Stock"},
// 				},
// 			},
// 			{
// 				Key:    "price",
// 				Label:  "Price Range",
// 				Type:   "numberrange",
// 				Active: false,
// 			},
// 		},
//
// 		QuickFilters: []DataTableFilter{
// 			{Key: "status", Operator: FilterEquals, Value: "active", Label: "Active", Active: false},
// 			{Key: "status", Operator: FilterEquals, Value: "out_of_stock", Label: "Out of Stock", Active: false},
// 			{Key: "stock", Operator: FilterGreaterThan, Value: 50, Label: "High Stock", Active: false},
// 		},
//
// 		Pagination: DataTablePagination{
// 			Enabled:         true,
// 			CurrentPage:     1,
// 			PageSize:        25,
// 			TotalPages:      10,
// 			TotalItems:      250,
// 			PageSizeOptions: []int{10, 25, 50, 100},
// 			ShowTotal:       true,
// 			ShowPageSize:    true,
// 			ShowQuickJump:   true,
// 			ServerSide:      true,
// 		},
//
// 		Actions: []DataTableAction{
// 			{
// 				ID:      "add_product",
// 				Text:    "Add Product",
// 				Icon:    "plus",
// 				Variant: atoms.ButtonPrimary,
// 				HXGet:   "/products/new",
// 			},
// 			{
// 				ID:      "import",
// 				Text:    "Import",
// 				Icon:    "upload",
// 				Variant: atoms.ButtonOutline,
// 				Type:    "dropdown",
// 				Items: []DataTableAction{
// 					{ID: "import_csv", Text: "Import CSV", Icon: "file-text"},
// 					{ID: "import_excel", Text: "Import Excel", Icon: "file-spreadsheet"},
// 				},
// 			},
// 			{
// 				ID:      "refresh",
// 				Text:    "Refresh",
// 				Icon:    "refresh-cw",
// 				Variant: atoms.ButtonGhost,
// 			},
// 		},
//
// 		BulkActions: []DataTableBulkAction{
// 			{
// 				ID:      "bulk_edit",
// 				Text:    "Edit Selected",
// 				Icon:    "edit",
// 				Variant: atoms.ButtonSecondary,
// 			},
// 			{
// 				ID:      "bulk_activate",
// 				Text:    "Activate Selected",
// 				Icon:    "check",
// 				Variant: atoms.ButtonSuccess,
// 			},
// 			{
// 				ID:          "bulk_delete",
// 				Text:        "Delete Selected",
// 				Icon:        "trash-2",
// 				Variant:     atoms.ButtonDestructive,
// 				Destructive: true,
// 				Confirm:     true,
// 				Message:     "Are you sure you want to delete the selected products?",
// 			},
// 		},
//
// 		Export: DataTableExport{
// 			Enabled:    true,
// 			Formats:    []ExportFormat{ExportCSV, ExportExcel, ExportPDF},
// 			Filename:   "products_export",
// 			AllData:    true,
// 			ServerSide: false,
// 		},
//
// 		Aggregation: struct {
// 			Enabled   bool              `json:\"enabled\"`
// 			Functions map[string]string `json:\"functions\"`
// 			Position  string            `json:\"position\"`
// 		}{
// 			Enabled: true,
// 			Functions: map[string]string{
// 				"price": "avg",
// 				"stock": "sum",
// 			},
// 			Position: "bottom",
// 		},
//
// 		// HTMX Configuration
// 		HXGet:    "/api/products",
// 		HXTarget: "#product-table",
// 		HXSwap:   "innerHTML",
//
// 		// Event Handlers
// 		OnRowClick:   "handleProductClick",
// 		OnRowSelect:  "handleProductSelect",
// 		OnSort:       "handleSort",
// 		OnFilter:     "handleFilter",
// 		OnSearch:     "handleSearch",
// 		OnPageChange: "handlePageChange",
// 		OnExport:     "handleExport",
// 	}
// }
//
// // Example 3: Schema-driven DataTable
// func (e *DataTableExamples) SchemaExample() (*DataTableProps, error) {
// 	// Create a schema for users
// 	userSchema := schema.NewSchema("user-table", schema.TypeForm, "User Management")
// 	userSchema.Description = "Manage system users with role-based access"
//
// 	// Add fields to the schema
// 	userSchema.AddField(schema.Field{
// 		Name:       "id",
// 		Type:       schema.FieldNumber,
// 		Label:      "ID",
// 		Hidden:     false,
// 		Sortable:   true,
// 		Searchable: false,
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:       "full_name",
// 		Type:       schema.FieldText,
// 		Label:      "Full Name",
// 		Required:   true,
// 		Sortable:   true,
// 		Searchable: true,
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:       "email",
// 		Type:       schema.FieldEmail,
// 		Label:      "Email Address",
// 		Required:   true,
// 		Sortable:   true,
// 		Searchable: true,
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:     "role",
// 		Type:     schema.FieldSelect,
// 		Label:    "Role",
// 		Sortable: true,
// 		Options: []schema.FieldOption{
// 			{Value: "admin", Label: "Administrator"},
// 			{Value: "editor", Label: "Editor"},
// 			{Value: "viewer", Label: "Viewer"},
// 		},
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:     "status",
// 		Type:     schema.FieldSelect,
// 		Label:    "Status",
// 		Sortable: true,
// 		Options: []schema.FieldOption{
// 			{Value: "active", Label: "Active"},
// 			{Value: "inactive", Label: "Inactive"},
// 			{Value: "suspended", Label: "Suspended"},
// 		},
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:     "created_at",
// 		Type:     schema.FieldDatetime,
// 		Label:    "Created At",
// 		Sortable: true,
// 	})
//
// 	userSchema.AddField(schema.Field{
// 		Name:     "last_login",
// 		Type:     schema.FieldDatetime,
// 		Label:    "Last Login",
// 		Sortable: true,
// 	})
//
// 	// Add actions to the schema
// 	userSchema.AddAction(schema.Action{
// 		ID:   "add_user",
// 		Text: "Add User",
// 		Icon: "plus",
// 		Type: schema.ActionSubmit,
// 	})
//
// 	userSchema.AddAction(schema.Action{
// 		ID:   "edit_user",
// 		Text: "Edit",
// 		Icon: "edit",
// 		Type: schema.ActionButton,
// 	})
//
// 	userSchema.AddAction(schema.Action{
// 		ID:             "delete_user",
// 		Text:           "Delete",
// 		Icon:           "trash",
// 		Type:           schema.ActionButton,
// 		Confirm:        true,
// 		ConfirmMessage: "Are you sure you want to delete this user?",
// 	})
//
// 	// Build DataTable from schema
// 	builder := NewDataTableSchemaBuilder(userSchema)
//
// 	// Configure custom column mappings
// 	builder.WithFieldMapping("role", ColumnMapping{
// 		ColumnType: ColumnTypeBadge,
// 		BadgeConfig: &BadgeConfig{
// 			VariantMap: map[string]string{
// 				"admin":  "primary",
// 				"editor": "secondary",
// 				"viewer": "default",
// 			},
// 		},
// 	})
//
// 	builder.WithFieldMapping("status", ColumnMapping{
// 		ColumnType: ColumnTypeBadge,
// 		BadgeConfig: &BadgeConfig{
// 			VariantMap: map[string]string{
// 				"active":    "success",
// 				"inactive":  "secondary",
// 				"suspended": "destructive",
// 			},
// 		},
// 	})
//
// 	// Configure table settings
// 	config := &DataTableConfig{
// 		EnableSelection:   true,
// 		EnableMultiSelect: true,
// 		EnableSorting:     true,
// 		EnableFiltering:   true,
// 		EnableSearch:      true,
// 		EnableExport:      true,
// 		EnablePagination:  true,
// 		DefaultPageSize:   25,
// 		SearchPlaceholder: "Search users by name or email...",
// 		SearchFields:      []string{"full_name", "email"},
// 		ExportFormats:     []ExportFormat{ExportCSV, ExportExcel},
// 		Variant:           DataTableDefault,
// 		Size:              DataTableSizeMD,
// 		Density:           DataTableDensityComfortable,
// 	}
//
// 	builder.WithConfig(config)
//
// 	props, err := builder.Build(nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	// Add sample data
// 	props.Rows = []DataTableRow{
// 		{
// 			ID: "user-1",
// 			Data: map[string]any{
// 				"id":         1,
// 				"full_name":  "John Admin",
// 				"email":      "admin@example.com",
// 				"role":       "admin",
// 				"status":     "active",
// 				"created_at": time.Now().AddDate(-1, 0, 0),
// 				"last_login": time.Now().Add(-2 * time.Hour),
// 			},
// 		},
// 		{
// 			ID: "user-2",
// 			Data: map[string]any{
// 				"id":         2,
// 				"full_name":  "Jane Editor",
// 				"email":      "editor@example.com",
// 				"role":       "editor",
// 				"status":     "active",
// 				"created_at": time.Now().AddDate(0, -6, 0),
// 				"last_login": time.Now().Add(-1 * time.Hour),
// 			},
// 		},
// 		{
// 			ID: "user-3",
// 			Data: map[string]any{
// 				"id":         3,
// 				"full_name":  "Bob Viewer",
// 				"email":      "viewer@example.com",
// 				"role":       "viewer",
// 				"status":     "suspended",
// 				"created_at": time.Now().AddDate(0, -3, 0),
// 				"last_login": time.Now().AddDate(0, 0, -7),
// 			},
// 		},
// 	}
//
// 	// Update pagination
// 	props.Pagination.TotalItems = len(props.Rows)
// 	props.Pagination.TotalPages = 1
//
// 	return props, nil
// }
//
// // Example 4: DataTable Builder Pattern
// func (e *DataTableExamples) BuilderExample() *DataTableProps {
// 	return NewDataTableBuilder("orders-table", "Order Management").
// 		WithDescription("Track and manage customer orders").
// 		WithVariant(DataTableHover).
// 		WithSize(DataTableSizeLG).
// 		AddColumn(DataTableColumn{
// 			Key:      "order_id",
// 			Title:    "Order #",
// 			Type:     ColumnTypeText,
// 			Width:    "120px",
// 			Sortable: true,
// 			Fixed:    true,
// 		}).
// 		AddColumn(DataTableColumn{
// 			Key:        "customer",
// 			Title:      "Customer",
// 			Type:       ColumnTypeText,
// 			Sortable:   true,
// 			Searchable: true,
// 			Clickable:  true,
// 		}).
// 		AddColumn(DataTableColumn{
// 			Key:          "total",
// 			Title:        "Total",
// 			Type:         ColumnTypeCurrency,
// 			Sortable:     true,
// 			Align:        "right",
// 			CurrencyCode: "USD",
// 			Precision:    2,
// 		}).
// 		AddColumn(DataTableColumn{
// 			Key:      "status",
// 			Title:    "Status",
// 			Type:     ColumnTypeBadge,
// 			Sortable: true,
// 			BadgeMap: map[string]atoms.BadgeVariant{
// 				"pending":   atoms.BadgeWarning,
// 				"confirmed": atoms.BadgePrimary,
// 				"shipped":   atoms.BadgeSecondary,
// 				"delivered": atoms.BadgeSuccess,
// 				"cancelled": atoms.BadgeDestructive,
// 			},
// 		}).
// 		AddColumn(DateColumn("created_at", "Order Date", "2006-01-02")).
// 		WithRows([]DataTableRow{
// 			{
// 				ID: "order-001",
// 				Data: map[string]any{
// 					"order_id":   "ORD-001",
// 					"customer":   "Alice Johnson",
// 					"total":      299.99,
// 					"status":     "delivered",
// 					"created_at": time.Now().AddDate(0, 0, -3),
// 				},
// 			},
// 			{
// 				ID: "order-002",
// 				Data: map[string]any{
// 					"order_id":   "ORD-002",
// 					"customer":   "Bob Wilson",
// 					"total":      149.50,
// 					"status":     "shipped",
// 					"created_at": time.Now().AddDate(0, 0, -1),
// 				},
// 			},
// 			{
// 				ID: "order-003",
// 				Data: map[string]any{
// 					"order_id":   "ORD-003",
// 					"customer":   "Carol Brown",
// 					"total":      89.99,
// 					"status":     "pending",
// 					"created_at": time.Now(),
// 				},
// 			},
// 		}).
// 		WithPagination(DataTablePagination{
// 			Enabled:         true,
// 			CurrentPage:     1,
// 			PageSize:        20,
// 			PageSizeOptions: []int{10, 20, 50},
// 			ShowTotal:       true,
// 		}).
// 		WithSearch(DataTableSearch{
// 			Enabled:     true,
// 			Placeholder: "Search orders...",
// 			Columns:     []string{"order_id", "customer"},
// 		}).
// 		WithActions([]DataTableAction{
// 			{
// 				ID:      "new_order",
// 				Text:    "New Order",
// 				Icon:    "plus",
// 				Variant: atoms.ButtonPrimary,
// 			},
// 			{
// 				ID:      "export_orders",
// 				Text:    "Export",
// 				Icon:    "download",
// 				Variant: atoms.ButtonOutline,
// 			},
// 		}).
// 		WithBulkActions([]DataTableBulkAction{
// 			{
// 				ID:      "bulk_ship",
// 				Text:    "Ship Selected",
// 				Icon:    "truck",
// 				Variant: atoms.ButtonPrimary,
// 			},
// 			{
// 				ID:          "bulk_cancel",
// 				Text:        "Cancel Selected",
// 				Icon:        "x",
// 				Variant:     atoms.ButtonDestructive,
// 				Destructive: true,
// 				Confirm:     true,
// 				Message:     "Cancel selected orders?",
// 			},
// 		}).
// 		WithRowActions([]DataTableAction{
// 			{ID: "view", Text: "View", Icon: "eye"},
// 			{ID: "edit", Text: "Edit", Icon: "edit"},
// 			{ID: "duplicate", Text: "Duplicate", Icon: "copy"},
// 		}).
// 		Build()
// }
//
// // Example 5: Responsive DataTable for mobile
// func (e *DataTableExamples) ResponsiveExample() *DataTableProps {
// 	return &DataTableProps{
// 		ID:          "mobile-table",
// 		Title:       "Mobile Contacts",
// 		Description: "Optimized for mobile devices",
// 		Variant:     DataTableCompact,
// 		Size:        DataTableSizeSM,
// 		Density:     DataTableDensityCondensed,
// 		Responsive:  true,
// 		Stackable:   true,
//
// 		Columns: []DataTableColumn{
// 			{
// 				Key:      "avatar",
// 				Title:    "",
// 				Type:     ColumnTypeAvatar,
// 				Width:    "40px",
// 				Sortable: false,
// 				Hidden:   false,
// 			},
// 			{
// 				Key:        "name",
// 				Title:      "Name",
// 				Type:       ColumnTypeText,
// 				Sortable:   true,
// 				Searchable: true,
// 			},
// 			{
// 				Key:    "phone",
// 				Title:  "Phone",
// 				Type:   ColumnTypeText,
// 				Hidden: true, // Hide on mobile, show in expanded view
// 			},
// 			{
// 				Key:   "status",
// 				Title: "Status",
// 				Type:  ColumnTypeBadge,
// 				Width: "80px",
// 				BadgeMap: map[string]atoms.BadgeVariant{
// 					"online":  atoms.BadgeSuccess,
// 					"offline": atoms.BadgeSecondary,
// 					"busy":    atoms.BadgeWarning,
// 				},
// 			},
// 		},
//
// 		Rows: []DataTableRow{
// 			{
// 				ID: "contact-1",
// 				Data: map[string]any{
// 					"avatar": map[string]any{
// 						"src":  "/avatars/john.jpg",
// 						"name": "John Doe",
// 					},
// 					"name":   "John Doe",
// 					"phone":  "+1 (555) 123-4567",
// 					"status": "online",
// 				},
// 			},
// 			{
// 				ID: "contact-2",
// 				Data: map[string]any{
// 					"avatar": map[string]any{
// 						"src":  "/avatars/jane.jpg",
// 						"name": "Jane Smith",
// 					},
// 					"name":   "Jane Smith",
// 					"phone":  "+1 (555) 987-6543",
// 					"status": "busy",
// 				},
// 			},
// 		},
//
// 		HideColumns: []string{"phone"}, // Hidden on mobile
//
// 		Search: DataTableSearch{
// 			Enabled:     true,
// 			Placeholder: "Search contacts...",
// 			MinLength:   1,
// 		},
//
// 		Pagination: DataTablePagination{
// 			Enabled:  true,
// 			PageSize: 10,
// 			Compact:  true,
// 		},
// 	}
// }
//
// // Example 6: High-performance virtualized table
// func (e *DataTableExamples) VirtualizedExample() *DataTableProps {
// 	// Generate large dataset
// 	rows := make([]DataTableRow, 10000)
// 	for i := 0; i < 10000; i++ {
// 		rows[i] = DataTableRow{
// 			ID: fmt.Sprintf("row-%d", i),
// 			Data: map[string]any{
// 				"id":       i + 1,
// 				"name":     fmt.Sprintf("Item %d", i+1),
// 				"category": []string{"A", "B", "C", "D"}[i%4],
// 				"value":    float64(i * 10),
// 				"active":   i%3 == 0,
// 			},
// 		}
// 	}
//
// 	return &DataTableProps{
// 		ID:          "virtualized-table",
// 		Title:       "Large Dataset (10,000 rows)",
// 		Description: "Demonstrates virtual scrolling for performance",
// 		Virtualized: true,
//
// 		Columns: []DataTableColumn{
// 			NumberColumn("id", "ID", 0),
// 			TextColumn("name", "Name"),
// 			BadgeColumn("category", "Category", map[string]atoms.BadgeVariant{
// 				"A": atoms.BadgePrimary,
// 				"B": atoms.BadgeSecondary,
// 				"C": atoms.BadgeSuccess,
// 				"D": atoms.BadgeWarning,
// 			}),
// 			NumberColumn("value", "Value", 2),
// 			{
// 				Key:   "active",
// 				Title: "Active",
// 				Type:  ColumnTypeCheckbox,
// 				Align: "center",
// 			},
// 		},
//
// 		Rows: rows,
//
// 		Search: DataTableSearch{
// 			Enabled:    true,
// 			ServerSide: true, // Server-side search for large datasets
// 		},
//
// 		Pagination: DataTablePagination{
// 			Enabled:    true,
// 			ServerSide: true, // Server-side pagination
// 			PageSize:   100,
// 		},
//
// 		VirtualScrollOptions: map[string]any{
// 			"rowHeight":    40,
// 			"overscan":     10,
// 			"enableResize": true,
// 		},
//
// 		LazyLoad:  true,
// 		CacheData: true,
// 	}
// }
//
// // Example 7: DataTable with custom cell renderers
// func (e *DataTableExamples) CustomRenderingExample() *DataTableProps {
// 	return &DataTableProps{
// 		ID:          "custom-table",
// 		Title:       "Custom Cell Rendering",
// 		Description: "Demonstrates custom column rendering",
//
// 		Columns: []DataTableColumn{
// 			TextColumn("product", "Product"),
// 			{
// 				Key:   "rating",
// 				Title: "Rating",
// 				Type:  ColumnTypeRating,
// 				Align: "center",
// 				Width: "120px",
// 			},
// 			{
// 				Key:      "tags",
// 				Title:    "Tags",
// 				Type:     ColumnTypeTags,
// 				Sortable: false,
// 			},
// 			{
// 				Key:   "progress",
// 				Title: "Completion",
// 				Type:  ColumnTypeProgress,
// 				Width: "150px",
// 			},
// 			{
// 				Key:       "link",
// 				Title:     "Documentation",
// 				Type:      ColumnTypeLink,
// 				Clickable: true,
// 			},
// 		},
//
// 		Rows: []DataTableRow{
// 			{
// 				ID: "item-1",
// 				Data: map[string]any{
// 					"product":  "Wireless Headphones",
// 					"rating":   4.5,
// 					"tags":     []any{"Electronics", "Audio", "Wireless"},
// 					"progress": 85.0,
// 					"link":     "View Docs",
// 					"link_url": "https://docs.example.com/headphones",
// 				},
// 			},
// 			{
// 				ID: "item-2",
// 				Data: map[string]any{
// 					"product":  "Smart Watch",
// 					"rating":   4.2,
// 					"tags":     []any{"Electronics", "Wearable", "Fitness"},
// 					"progress": 92.0,
// 					"link":     "User Guide",
// 					"link_url": "https://docs.example.com/watch",
// 				},
// 			},
// 		},
// 	}
// }
//
// // Example 8: DataTable with grouping and aggregation
// func (e *DataTableExamples) GroupingExample() *DataTableProps {
// 	return &DataTableProps{
// 		ID:          "grouped-table",
// 		Title:       "Sales by Region",
// 		Description: "Revenue data grouped by region",
//
// 		Columns: []DataTableColumn{
// 			TextColumn("region", "Region"),
// 			TextColumn("salesperson", "Salesperson"),
// 			NumberColumn("revenue", "Revenue", 2),
// 			NumberColumn("deals", "Deals Closed", 0),
// 			DateColumn("last_sale", "Last Sale", "2006-01-02"),
// 		},
//
// 		Rows: []DataTableRow{
// 			{
// 				ID: "sale-1",
// 				Data: map[string]any{
// 					"region":      "North",
// 					"salesperson": "Alice Johnson",
// 					"revenue":     125000.00,
// 					"deals":       15,
// 					"last_sale":   time.Now().AddDate(0, 0, -2),
// 				},
// 			},
// 			{
// 				ID: "sale-2",
// 				Data: map[string]any{
// 					"region":      "North",
// 					"salesperson": "Bob Wilson",
// 					"revenue":     98000.00,
// 					"deals":       12,
// 					"last_sale":   time.Now().AddDate(0, 0, -1),
// 				},
// 			},
// 			{
// 				ID: "sale-3",
// 				Data: map[string]any{
// 					"region":      "South",
// 					"salesperson": "Carol Davis",
// 					"revenue":     156000.00,
// 					"deals":       18,
// 					"last_sale":   time.Now(),
// 				},
// 			},
// 		},
//
// 		Grouping: struct {
// 			Enabled     bool     `json:"enabled"`
// 			Columns     []string `json:"columns"`
// 			Expanded    []string `json:"expanded"`
// 			Collapsible bool     `json:"collapsible"`
// 		}{
// 			Enabled:     true,
// 			Columns:     []string{"region"},
// 			Collapsible: true,
// 		},
//
// 		Aggregation: struct {
// 			Enabled   bool              `json:"enabled"`
// 			Functions map[string]string `json:"functions"`
// 			Position  string            `json:"position"`
// 		}{
// 			Enabled: true,
// 			Functions: map[string]string{
// 				"revenue": "sum",
// 				"deals":   "sum",
// 			},
// 			Position: "both", // Show aggregates at top and bottom
// 		},
// 	}
// }
//
// // GetAllExamples returns all available examples
// func (e *DataTableExamples) GetAllExamples() map[string]*DataTableProps {
// 	examples := make(map[string]*DataTableProps)
//
// 	examples["basic"] = e.BasicExample()
// 	examples["advanced"] = e.AdvancedExample()
// 	examples["builder"] = e.BuilderExample()
// 	examples["responsive"] = e.ResponsiveExample()
// 	examples["virtualized"] = e.VirtualizedExample()
// 	examples["custom"] = e.CustomRenderingExample()
// 	examples["grouping"] = e.GroupingExample()
//
// 	// Schema example requires error handling
// 	if schemaExample, err := e.SchemaExample(); err == nil {
// 		examples["schema"] = schemaExample
// 	}
//
// 	return examples
// }
//
// // GetExampleByName returns a specific example by name
// func (e *DataTableExamples) GetExampleByName(name string) (*DataTableProps, bool) {
// 	examples := e.GetAllExamples()
// 	example, exists := examples[name]
// 	return example, exists
// }
//
// // Example usage in Go handler:
// /*
// func ProductTableHandler(w http.ResponseWriter, r *http.Request) {
// 	examples := &DataTableExamples{}
// 	props := examples.AdvancedExample()
//
// 	// Apply query parameters for filtering, sorting, pagination
// 	if searchQuery := r.URL.Query().Get("search"); searchQuery != "" {
// 		props.Search.Query = searchQuery
// 	}
//
// 	// Render the table
// 	component := organisms.DataTable(*props)
// 	component.Render(r.Context(), w)
// }
//
// // Example usage in templ template:
// templ ProductPage() {
// 	@Layout("Products") {
// 		<div class="container mx-auto py-8">
// 			@organisms.DataTable(organisms.DataTableExamples{}.AdvancedExample())
// 		</div>
// 	}
// }
// */
