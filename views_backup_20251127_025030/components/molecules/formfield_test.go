package molecules

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// FormFieldTestSuite is the test suite for form field components
type FormFieldTestSuite struct {
	suite.Suite
}

// TestFormFieldSuite runs the test suite
func TestFormFieldSuite(t *testing.T) {
	suite.Run(t, new(FormFieldTestSuite))
}

// TestFormFieldProps tests basic FormFieldProps functionality
func (s *FormFieldTestSuite) TestFormFieldProps() {
	s.Run("create basic FormFieldProps", func() {
		props := FormFieldProps{
			ID:       "test-field",
			Name:     "testName", 
			Type:     "text",
			Value:    "test value",
			Required: true,
		}
		
		require.Equal(s.T(), "test-field", props.ID)
		require.Equal(s.T(), "testName", props.Name)
		require.Equal(s.T(), "text", props.Type)
		require.Equal(s.T(), "test value", props.Value)
		require.True(s.T(), props.Required)
	})
	
	s.Run("create FormFieldProps with options", func() {
		options := []SelectOption{
			{Value: "opt1", Label: "Option 1"},
			{Value: "opt2", Label: "Option 2", Selected: true},
		}
		
		props := FormFieldProps{
			Type:    "select",
			Options: options,
		}
		
		require.Len(s.T(), props.Options, 2)
		require.Equal(s.T(), "opt1", props.Options[0].Value)
		require.True(s.T(), props.Options[1].Selected)
	})
}

// TestSelectOption tests SelectOption functionality  
func (s *FormFieldTestSuite) TestSelectOption() {
	s.Run("create basic SelectOption", func() {
		option := SelectOption{
			Value:    "test-value",
			Label:    "Test Label",
			Disabled: false,
			Selected: true,
		}
		
		require.Equal(s.T(), "test-value", option.Value)
		require.Equal(s.T(), "Test Label", option.Label)
		require.False(s.T(), option.Disabled)
		require.True(s.T(), option.Selected)
	})
}

// TestFormFieldClasses tests CSS class generation
func (s *FormFieldTestSuite) TestFormFieldClasses() {
	s.Run("generates form field classes", func() {
		props := FormFieldProps{
			Size:     "lg",
			Variant:  "outline", 
			Required: true,
			Disabled: false,
		}
		
		classes := getFormFieldClasses(props)
		require.Contains(s.T(), classes, "form-field")
		require.Contains(s.T(), classes, "form-field-lg")
		require.Contains(s.T(), classes, "form-field-outline")
		require.Contains(s.T(), classes, "form-field-required")
	})
}

// TestLabelClasses tests label CSS class generation
func (s *FormFieldTestSuite) TestLabelClasses() {
	s.Run("generates label classes", func() {
		props := FormFieldProps{
			Size:     "md",
			Required: true,
			Disabled: false,
		}
		
		classes := getLabelClasses(props)
		require.Contains(s.T(), classes, "form-field-label")
		require.Contains(s.T(), classes, "form-field-label-md")
		require.Contains(s.T(), classes, "form-field-label-required")
	})
}

// TestValidationStateClasses tests validation state CSS classes
func (s *FormFieldTestSuite) TestValidationStateClasses() {
	s.Run("generates validation state classes", func() {
		testCases := []struct {
			state    string
			expected string
		}{
			{string(ValidationStateIdle), "form-field-state-idle"},
			{string(ValidationStateValidating), "form-field-state-validating"},
			{string(ValidationStateValid), "form-field-state-valid"},
			{string(ValidationStateInvalid), "form-field-state-invalid"},
			{string(ValidationStateWarning), "form-field-state-warning"},
		}
		
		for _, tc := range testCases {
			result := getValidationStateClasses(tc.state)
			require.Equal(s.T(), tc.expected, result)
		}
	})
}

// Placeholder tests for methods not yet implemented
func (s *FormFieldTestSuite) TestPlaceholderMethods() {
	s.Run("validation methods placeholder", func() {
		// Validation methods (Validate, ValidateMultiple) not implemented
		// Tests will be added when validation is implemented
		s.T().Skip("Validation methods not implemented")
	})
	
	s.Run("normalization methods placeholder", func() {
		// NormalizeDefaults method not implemented
		// Tests will be added when normalization is implemented  
		s.T().Skip("NormalizeDefaults method not implemented")
	})
	
	s.Run("selection methods placeholder", func() {
		// GetSelectedValue, GetSelectedValues methods not implemented
		// Tests will be added when selection methods are implemented
		s.T().Skip("Selection methods not implemented")
	})
}