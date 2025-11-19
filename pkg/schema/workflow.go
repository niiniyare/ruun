package schema

import (
	"time"

	"github.com/niiniyare/ruun/pkg/condition"
)

// Workflow defines workflow and approval configuration
type Workflow struct {
	Enabled       bool                 `json:"enabled"`                                 // Enable workflow
	ID            string               `json:"id,omitempty" validate:"uuid"`            // Workflow ID
	Stage         string               `json:"stage,omitempty"`                         // Current stage
	Status        string               `json:"status,omitempty"`                        // Current status
	Actions       []WorkflowAction     `json:"actions,omitempty" validate:"dive"`       // Available actions
	Transitions   []WorkflowTransition `json:"transitions,omitempty" validate:"dive"`   // Stage transitions
	Notifications []Notification       `json:"notifications,omitempty" validate:"dive"` // Notifications
	Approvals     *ApprovalConfig      `json:"approvals,omitempty"`                     // Approval configuration
	History       bool                 `json:"history,omitempty"`                       // Track workflow history
	Timeout       *WorkflowTimeout     `json:"timeout,omitempty"`                       // Stage timeout
	Audit         bool                 `json:"audit,omitempty"`                         // Enable audit logging
	Versioning    bool                 `json:"versioning,omitempty"`                    // Enable versioning
}

// WorkflowAction defines an action that can be taken in the workflow
type WorkflowAction struct {
	ID          string                    `json:"id" validate:"required"`
	Label       string                    `json:"label" validate:"required"`
	Type        string                    `json:"type" validate:"oneof=approve reject submit return delegate escalate" example:"approve"`
	ToStage     string                    `json:"toStage,omitempty"`     // Next stage
	ToStatus    string                    `json:"toStatus,omitempty"`    // Next status
	Permissions []string                  `json:"permissions,omitempty"` // Required permissions
	RequireNote bool                      `json:"requireNote,omitempty"` // Require comment
	Icon        string                    `json:"icon,omitempty" validate:"icon_name"`
	Variant     string                    `json:"variant,omitempty" validate:"oneof=primary secondary outline destructive"`
	Condition   *condition.ConditionGroup `json:"condition,omitempty"` // Action visibility conditions
}

// WorkflowTransition defines a stage transition
type WorkflowTransition struct {
	From       string                    `json:"from" validate:"required"`   // From stage
	To         string                    `json:"to" validate:"required"`     // To stage
	Action     string                    `json:"action" validate:"required"` // Triggering action
	Conditions *condition.ConditionGroup `json:"conditions,omitempty"`       // Transition conditions
	Validator  string                    `json:"validator,omitempty"`        // Custom validator
	Auto       bool                      `json:"auto,omitempty"`             // Automatic transition
	Delay      time.Duration             `json:"delay,omitempty"`            // Transition delay
}

// WorkflowTimeout defines timeout for workflow stages
type WorkflowTimeout struct {
	Duration time.Duration `json:"duration" validate:"min=1"` // Timeout duration
	Action   string        `json:"action" validate:"oneof=escalate auto-approve auto-reject notify"`
	NotifyAt []int         `json:"notifyAt,omitempty"` // Notify at % of timeout (e.g., [50, 75, 90])
}
