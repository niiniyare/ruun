// package models  provides a comprehensive JSON-driven UI schema system for enterprise applications.
/*
````md
# Overview
This package enables backend developers to define complete user interfaces through JSON
configuration files without writing frontend code. It supports forms, multi-step wizards,
tabbed interfaces, and complex enterprise workflows.
# Core Concepts
## Schema
The Schema is the root object that defines a complete form or UI component. It includes:
  - Fields: Input elements and form controls
  - Actions: Buttons and interactive elements
  - Layout: Visual structure (grid, tabs, sections, steps)
  - Security: CSRF, rate limiting, encryption
  - Workflow: Approval processes and state transitions
## Fields
Fields represent individual form inputs. The system supports 30+ field types including:
  - Basic: text, email, password, number
  - Selection: select, multi-select, radio, checkbox
  - Rich content: textarea, rich text editor, code editor
  - Files: file upload, image upload, signature
  - Specialized: currency, tags, location, date ranges
## Validation
Comprehensive validation at multiple levels:
  - Field-level: required, min/max, patterns, custom validators
  - Cross-field: compare fields, complex business rules
  - Async: server-side validation during input
## Multi-Tenancy
Built-in tenant isolation with three modes:
  - Strict: Complete data isolation per tenant
  - Shared: Data shared across tenants
  - Hybrid: Mix of shared and isolated data
# Usage
## Basic Form
	schema := nodels.NewBuilder("user-form", nodels.TypeForm, "Create User").
		WithConfig(nodels.NewSimpleConfig("/api/v1/users", "POST")).
		AddTextField("name", "Full Name", true).
		AddEmailField("email", "Email", true).
		AddSubmitButton("Create").
		MustBuild()
## Form with Validation
	minLen := 3
	maxLen := 50
	field := nodels.NewField("username", nodels.FieldText).
		WithLabel("Username").
		Required().
		WithValidation(&nodels.FieldValidation{
			MinLength: &minLen,
			MaxLength: &maxLen,
			Pattern:   "^[a-zA-Z0-9_]+$",
		}).
		Build()
## Multi-Tenant Form
	schema := nodels.NewBuilder("invoice", nodels.TypeForm, "Invoice").
		WithTenant("tenant_id", "strict").
		WithCSRF().
		AddTextField("invoice_number", "Invoice #", true).
		MustBuild()
## Multi-Step Wizard
	schema := nodels.NewBuilder("onboarding", nodels.TypeWorkflow, "Onboarding").
		WithLayout(&nodels.Layout{
			Type: nodels.LayoutSteps,
			Steps: []nodels.Step{
				{ID: "step1", Title: "Personal", Fields: []string{"name", "email"}},
				{ID: "step2", Title: "Company", Fields: []string{"company", "role"}},
			},
		}).
		AddTextField("name", "Name", true).
		AddEmailField("email", "Email", true).
		MustBuild()
# Framework Integration
## HTMX
Enable HTMX for form submission and dynamic updates:
	nodels.HTMX = &nodels.HTMX{
		Enabled: true,
		Post:    "/api/v1/users",
		Target:  "#main-content",
		Swap:    "innerHTML",
	}
## Reactive Binding
Add reactive behavior with universal binding:
	nodels.Binding = &nodels.Binding{
		Data: `{count: 0, total: 0}`,
		Show: "count > 0",
	}
# Enterprise Features
## Security
	nodels.Security = &nodels.Security{
		CSRF: &nodels.CSRF{Enabled: true},
		RateLimit: &nodels.RateLimit{
			Enabled:     true,
			MaxRequests: 100,
		},
		Encryption: &nodels.Encryption{
			Enabled: true,
			Fields:  []string{"ssn", "credit_card"},
		},
	}
## Workflows
	nodels.Workflow = &nodels.Workflow{
		Enabled: true,
		Actions: []nodels.WorkflowAction{
			{ID: "approve", Label: "Approve", Type: "approve"},
			{ID: "reject", Label: "Reject", Type: "reject"},
		},
		Approvals: &nodels.ApprovalConfig{
			Required:     true,
			MinApprovals: 2,
		},
	}
# Extensibility
Implement interfaces to extend functionality:
	// See interface.go for complete interface definitions including:
	// - Validator: ValidateSchema, ValidateData, ValidateField
	// - Renderer: Render, RenderField, Format
	// - Registry: Register, Get, List, Delete, Exists
# Error Handling
The package provides typed errors for better error handling:
	schema, err := builder.Build()
	if err != nil {
		var validationErr nodels.ValidationError
		if errors.As(err, &validationErr) {
			log.Printf("Field %s: %s", validationErr.Field, validationErr.Message)
		}
	}
````
*/
package schema
