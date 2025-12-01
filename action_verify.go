package main

import (
	"fmt"
	"github.com/niiniyare/ruun/schema"
)

func main() {
	// Test Action struct with unified types
	action := schema.NewAction("test", schema.ActionButton, "Test Action")
	
	// Test unified Config
	config := map[string]any{
		"url": "https://example.com",
		"method": "POST",
	}
	
	// Test unified Style
	style := &schema.Style{
		Classes: "btn btn-primary",
		Colors: map[string]string{
			"background": "#007bff",
			"text": "#ffffff",
		},
	}
	
	// Test unified Behavior  
	behavior := &schema.Behavior{
		URL: "https://api.example.com/action",
		Method: "POST",
	}
	
	// Test unified Binding
	binding := &schema.Binding{
		Model: "user.name",
		Event: "input",
	}
	
	// Test unified Conditional
	conditional := &schema.Conditional{}
	
	// Build action with unified types
	builtAction := action.
		WithConfig(config).
		WithStyle(style).
		WithBehavior(behavior).
		WithBinding(binding).
		WithConditional(conditional).
		WithPermissions([]string{"user", "admin"}).
		Build()
	
	// Test helper methods
	fmt.Printf("Action ID: %s\n", builtAction.ID)
	fmt.Printf("Config URL: %s\n", builtAction.GetConfigString("url"))
	fmt.Printf("Has Permission 'admin': %t\n", builtAction.HasPermission("admin"))
	fmt.Printf("Has Permission 'guest': %t\n", builtAction.HasPermission("guest"))
	
	fmt.Println("✅ Action struct with unified types working correctly!")
}