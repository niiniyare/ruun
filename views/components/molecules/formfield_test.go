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

// TestValidation tests the validation logic
func (s *FormFieldTestSuite) TestValidation() {
	tests := []struct {
		name      string
		props     FormFieldProps
		value     string
		wantError bool
	}{
		{
			name: "required field with value",
			props: FormFieldProps{
				Type:     FormFieldText,
				Required: true,
			},
			value:     "test",
			wantError: false,
		},
		{
			name: "required field without value",
			props: FormFieldProps{
				Type:     FormFieldText,
				Required: true,
			},
			value:     "",
			wantError: true,
		},
		{
			name: "email validation - valid",
			props: FormFieldProps{
				Type: FormFieldEmail,
			},
			value:     "user@example.com",
			wantError: false,
		},
		{
			name: "email validation - invalid",
			props: FormFieldProps{
				Type: FormFieldEmail,
			},
			value:     "invalid-email",
			wantError: true,
		},
		{
			name: "minLength validation - pass",
			props: FormFieldProps{
				Type:      FormFieldText,
				MinLength: 5,
			},
			value:     "12345",
			wantError: false,
		},
		{
			name: "minLength validation - fail",
			props: FormFieldProps{
				Type:      FormFieldText,
				MinLength: 5,
			},
			value:     "123",
			wantError: true,
		},
		{
			name: "maxLength validation - pass",
			props: FormFieldProps{
				Type:      FormFieldText,
				MaxLength: 10,
			},
			value:     "12345",
			wantError: false,
		},
		{
			name: "maxLength validation - fail",
			props: FormFieldProps{
				Type:      FormFieldText,
				MaxLength: 10,
			},
			value:     "12345678901",
			wantError: true,
		},
		{
			name: "pattern validation - pass",
			props: FormFieldProps{
				Type:    FormFieldText,
				Pattern: "^[a-zA-Z]+$",
			},
			value:     "hello",
			wantError: false,
		},
		{
			name: "pattern validation - fail",
			props: FormFieldProps{
				Type:    FormFieldText,
				Pattern: "^[a-zA-Z]+$",
			},
			value:     "hello123",
			wantError: true,
		},
		{
			name: "URL validation - valid",
			props: FormFieldProps{
				Type: FormFieldURL,
			},
			value:     "https://example.com",
			wantError: false,
		},
		{
			name: "URL validation - invalid",
			props: FormFieldProps{
				Type: FormFieldURL,
			},
			value:     "not-a-url",
			wantError: true,
		},
		{
			name: "number validation - valid",
			props: FormFieldProps{
				Type: FormFieldNumber,
			},
			value:     "123.45",
			wantError: false,
		},
		{
			name: "number validation - invalid",
			props: FormFieldProps{
				Type: FormFieldNumber,
			},
			value:     "not-a-number",
			wantError: true,
		},
		{
			name: "optional field empty",
			props: FormFieldProps{
				Type:     FormFieldText,
				Required: false,
			},
			value:     "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := tt.props.Validate(tt.value)
			if tt.wantError {
				require.Error(s.T(), err, "Expected validation error")
			} else {
				require.NoError(s.T(), err, "Expected no validation error")
			}
		})
	}
}

// TestValidationRules tests custom validation rules
func (s *FormFieldTestSuite) TestValidationRules() {
	tests := []struct {
		name      string
		props     FormFieldProps
		value     string
		wantError bool
		errMsg    string
	}{
		{
			name: "custom minLength rule",
			props: FormFieldProps{
				Type: FormFieldText,
				ValidationRules: []ValidationRule{
					{
						Type:    "minLength",
						Value:   5,
						Message: "Too short",
					},
				},
			},
			value:     "test",
			wantError: true,
			errMsg:    "Too short",
		},
		{
			name: "custom maxLength rule",
			props: FormFieldProps{
				Type: FormFieldText,
				ValidationRules: []ValidationRule{
					{
						Type:    "maxLength",
						Value:   10,
						Message: "Too long",
					},
				},
			},
			value:     "this is a very long text",
			wantError: true,
			errMsg:    "Too long",
		},
		{
			name: "custom pattern rule",
			props: FormFieldProps{
				Type: FormFieldText,
				ValidationRules: []ValidationRule{
					{
						Type:    "pattern",
						Value:   "^[0-9]+$",
						Message: "Must be numeric",
					},
				},
			},
			value:     "abc",
			wantError: true,
			errMsg:    "Must be numeric",
		},
		{
			name: "passing custom rule",
			props: FormFieldProps{
				Type: FormFieldText,
				ValidationRules: []ValidationRule{
					{
						Type:  "minLength",
						Value: 3,
					},
				},
			},
			value:     "test",
			wantError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := tt.props.Validate(tt.value)
			if tt.wantError {
				require.Error(s.T(), err)
				if tt.errMsg != "" {
					require.Contains(s.T(), err.Error(), tt.errMsg)
				}
			} else {
				require.NoError(s.T(), err)
			}
		})
	}
}

// TestValidateMultiple tests multi-value validation
func (s *FormFieldTestSuite) TestValidateMultiple() {
	tests := []struct {
		name      string
		props     FormFieldProps
		values    []string
		wantError bool
	}{
		{
			name: "required with values",
			props: FormFieldProps{
				Type:     FormFieldCheckboxGroup,
				Required: true,
			},
			values:    []string{"value1", "value2"},
			wantError: false,
		},
		{
			name: "required without values",
			props: FormFieldProps{
				Type:     FormFieldCheckboxGroup,
				Required: true,
			},
			values:    []string{},
			wantError: true,
		},
		{
			name: "optional without values",
			props: FormFieldProps{
				Type:     FormFieldCheckboxGroup,
				Required: false,
			},
			values:    []string{},
			wantError: false,
		},
		{
			name: "required with nil values",
			props: FormFieldProps{
				Type:     FormFieldCheckboxGroup,
				Required: true,
			},
			values:    nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := tt.props.ValidateMultiple(tt.values)
			if tt.wantError {
				require.Error(s.T(), err)
			} else {
				require.NoError(s.T(), err)
			}
		})
	}
}

// TestNormalizeDefaults tests default value normalization
func (s *FormFieldTestSuite) TestNormalizeDefaults() {
	s.Run("size defaults to MD", func() {
		props := FormFieldProps{}
		props.NormalizeDefaults()
		require.Equal(s.T(), FormFieldSizeMD, props.Size)
	})

	s.Run("variant defaults to default", func() {
		props := FormFieldProps{}
		props.NormalizeDefaults()
		require.Equal(s.T(), FormFieldVariantDefault, props.Variant)
	})

	s.Run("autocomplete minChars defaults to 3", func() {
		props := FormFieldProps{Type: FormFieldAutoComplete}
		props.NormalizeDefaults()
		require.Equal(s.T(), 3, props.MinChars)
	})

	s.Run("autocomplete maxResults defaults to 10", func() {
		props := FormFieldProps{Type: FormFieldAutoComplete}
		props.NormalizeDefaults()
		require.Equal(s.T(), 10, props.MaxResults)
	})

	s.Run("autocomplete debounce defaults to 300", func() {
		props := FormFieldProps{Type: FormFieldAutoComplete}
		props.NormalizeDefaults()
		require.Equal(s.T(), 300, props.Debounce)
	})

	s.Run("textarea rows defaults to 3", func() {
		props := FormFieldProps{Type: FormFieldTextarea}
		props.NormalizeDefaults()
		require.Equal(s.T(), 3, props.Rows)
	})

	s.Run("checkbox group columns defaults to 1", func() {
		props := FormFieldProps{Type: FormFieldCheckboxGroup}
		props.NormalizeDefaults()
		require.Equal(s.T(), 1, props.Columns)
	})

	s.Run("radio group columns defaults to 1", func() {
		props := FormFieldProps{Type: FormFieldRadio}
		props.NormalizeDefaults()
		require.Equal(s.T(), 1, props.Columns)
	})

	s.Run("rounded defaults to md", func() {
		props := FormFieldProps{}
		props.NormalizeDefaults()
		require.Equal(s.T(), "md", props.Rounded)
	})

	s.Run("disabled sets tabindex to -1", func() {
		props := FormFieldProps{Disabled: true}
		props.NormalizeDefaults()
		require.Equal(s.T(), -1, props.TabIndex)
	})

	s.Run("date format defaults for date field", func() {
		props := FormFieldProps{Type: FormFieldDate}
		props.NormalizeDefaults()
		require.Equal(s.T(), "YYYY-MM-DD", props.DateFormat)
	})

	s.Run("aria-describedby set for error", func() {
		props := FormFieldProps{
			ID:        "test-field",
			ErrorText: "Error message",
		}
		props.NormalizeDefaults()
		require.Equal(s.T(), "test-field-error", props.AriaDescribedBy)
	})

	s.Run("aria-describedby set for help text", func() {
		props := FormFieldProps{
			ID:       "test-field",
			HelpText: "Help text",
		}
		props.NormalizeDefaults()
		require.Equal(s.T(), "test-field-help", props.AriaDescribedBy)
	})

	s.Run("custom values not overridden", func() {
		props := FormFieldProps{
			Type:     FormFieldAutoComplete,
			Size:     FormFieldSizeLG,
			MinChars: 5,
			Debounce: 500,
			Rounded:  "lg",
			TabIndex: 2,
		}
		props.NormalizeDefaults()
		require.Equal(s.T(), FormFieldSizeLG, props.Size)
		require.Equal(s.T(), 5, props.MinChars)
		require.Equal(s.T(), 500, props.Debounce)
		require.Equal(s.T(), "lg", props.Rounded)
		require.Equal(s.T(), 2, props.TabIndex)
	})
}

// TestGetSelectedValue tests getting selected value from options
func (s *FormFieldTestSuite) TestGetSelectedValue() {
	s.Run("selected option present", func() {
		props := FormFieldProps{
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1"},
				{Value: "opt2", Label: "Option 2", Selected: true},
				{Value: "opt3", Label: "Option 3"},
			},
		}
		result := props.GetSelectedValue()
		require.Equal(s.T(), "opt2", result)
	})

	s.Run("no selected option, uses value", func() {
		props := FormFieldProps{
			Value: "fallback",
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1"},
				{Value: "opt2", Label: "Option 2"},
			},
		}
		result := props.GetSelectedValue()
		require.Equal(s.T(), "fallback", result)
	})

	s.Run("no selected option or value", func() {
		props := FormFieldProps{
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1"},
			},
		}
		result := props.GetSelectedValue()
		require.Empty(s.T(), result)
	})

	s.Run("multiple selected options returns first", func() {
		props := FormFieldProps{
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1", Selected: true},
				{Value: "opt2", Label: "Option 2", Selected: true},
			},
		}
		result := props.GetSelectedValue()
		require.Equal(s.T(), "opt1", result)
	})

	s.Run("empty options slice", func() {
		props := FormFieldProps{
			Value:   "default",
			Options: []SelectOption{},
		}
		result := props.GetSelectedValue()
		require.Equal(s.T(), "default", result)
	})
}

// TestGetSelectedValues tests getting multiple selected values
func (s *FormFieldTestSuite) TestGetSelectedValues() {
	s.Run("values present", func() {
		props := FormFieldProps{
			Values: []string{"val1", "val2"},
		}
		result := props.GetSelectedValues()
		require.Equal(s.T(), []string{"val1", "val2"}, result)
	})

	s.Run("selected options", func() {
		props := FormFieldProps{
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1", Selected: true},
				{Value: "opt2", Label: "Option 2"},
				{Value: "opt3", Label: "Option 3", Selected: true},
			},
		}
		result := props.GetSelectedValues()
		require.Equal(s.T(), []string{"opt1", "opt3"}, result)
	})

	s.Run("values takes precedence over options", func() {
		props := FormFieldProps{
			Values: []string{"val1"},
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1", Selected: true},
			},
		}
		result := props.GetSelectedValues()
		require.Equal(s.T(), []string{"val1"}, result)
	})

	s.Run("no values or selected options", func() {
		props := FormFieldProps{
			Options: []SelectOption{
				{Value: "opt1", Label: "Option 1"},
			},
		}
		result := props.GetSelectedValues()
		require.Empty(s.T(), result)
	})

	s.Run("empty values slice", func() {
		props := FormFieldProps{
			Values: []string{},
		}
		result := props.GetSelectedValues()
		require.NotNil(s.T(), result)
		require.Len(s.T(), result, 0)
	})
}

// TestSanitizeAlpineExpression tests XSS prevention
func (s *FormFieldTestSuite) TestSanitizeAlpineExpression() {
	tests := []struct {
		name      string
		expr      string
		wantError bool
	}{
		{
			name:      "safe expression",
			expr:      "count++",
			wantError: false,
		},
		{
			name:      "safe method call",
			expr:      "updateTotal()",
			wantError: false,
		},
		{
			name:      "script tag lowercase",
			expr:      "<script>alert('xss')</script>",
			wantError: true,
		},
		{
			name:      "script tag uppercase",
			expr:      "<SCRIPT>alert('xss')</SCRIPT>",
			wantError: true,
		},
		{
			name:      "javascript protocol",
			expr:      "javascript:alert('xss')",
			wantError: true,
		},
		{
			name:      "onerror attribute",
			expr:      "onerror=alert('xss')",
			wantError: true,
		},
		{
			name:      "onload attribute",
			expr:      "onload=alert('xss')",
			wantError: true,
		},
		{
			name:      "empty expression",
			expr:      "",
			wantError: false,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			props := FormFieldProps{}
			result, err := props.SanitizeAlpineExpression(tt.expr)
			if tt.wantError {
				require.Error(s.T(), err)
				require.Empty(s.T(), result)
			} else {
				require.NoError(s.T(), err)
				if tt.expr != "" {
					require.NotEmpty(s.T(), result)
				}
			}
		})
	}
}

// TestConversionFunctions tests option conversion helpers
func (s *FormFieldTestSuite) TestConvertToRadioOptions() {
	s.Run("converts options correctly", func() {
		input := []SelectOption{
			{Value: "1", Label: "One", Description: "First option"},
			{Value: "2", Label: "Two", Disabled: true},
			{Value: "3", Label: "Three", Icon: "star"},
		}

		result := convertToRadioOptions(input)

		require.Len(s.T(), result, 3)
		require.Equal(s.T(), "1", result[0].Value)
		require.Equal(s.T(), "One", result[0].Label)
		require.Equal(s.T(), "First option", result[0].Description)
		require.True(s.T(), result[1].Disabled)
		require.Equal(s.T(), "star", result[2].Icon)
	})

	s.Run("handles nil input", func() {
		result := convertToRadioOptions(nil)
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})

	s.Run("handles empty slice", func() {
		result := convertToRadioOptions([]SelectOption{})
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

func (s *FormFieldTestSuite) TestConvertToCheckboxOptions() {
	s.Run("converts options with selected state", func() {
		input := []SelectOption{
			{Value: "1", Label: "One", Selected: true},
			{Value: "2", Label: "Two"},
			{Value: "3", Label: "Three", Disabled: true, Selected: true},
		}

		result := convertToCheckboxOptions(input)

		require.Len(s.T(), result, 3)
		require.True(s.T(), result[0].Checked)
		require.False(s.T(), result[1].Checked)
		require.True(s.T(), result[2].Checked)
		require.True(s.T(), result[2].Disabled)
	})

	s.Run("handles nil input", func() {
		result := convertToCheckboxOptions(nil)
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

func (s *FormFieldTestSuite) TestConvertToSelectOptions() {
	s.Run("converts with all properties", func() {
		input := []SelectOption{
			{Value: "1", Label: "One", Group: "Numbers", Selected: true},
			{Value: "a", Label: "A", Group: "Letters", Icon: "letter"},
		}

		result := convertToSelectOptions(input)

		require.Len(s.T(), result, 2)
		require.Equal(s.T(), "Numbers", result[0].Group)
		require.True(s.T(), result[0].Selected)
		require.Equal(s.T(), "letter", result[1].Icon)
	})

	s.Run("handles nil input", func() {
		result := convertToSelectOptions(nil)
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

func (s *FormFieldTestSuite) TestConvertToAutoCompleteOptions() {
	s.Run("converts with meta data", func() {
		input := []SelectOption{
			{Value: "1", Label: "One", Description: "Desc", Meta: "metadata"},
			{Value: "2", Label: "Two", Icon: "icon2"},
		}

		result := convertToAutoCompleteOptions(input)

		require.Len(s.T(), result, 2)
		require.Equal(s.T(), "metadata", result[0].Meta)
		require.Equal(s.T(), "Desc", result[0].Description)
		require.Equal(s.T(), "icon2", result[1].Icon)
	})

	s.Run("handles nil input", func() {
		result := convertToAutoCompleteOptions(nil)
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

func (s *FormFieldTestSuite) TestConvertSelectOptionsToTagProps() {
	s.Run("converts to tag props", func() {
		input := []SelectOption{
			{Value: "tag1", Label: "Tag 1", Selected: true},
			{Value: "tag2", Label: "Tag 2", Disabled: true},
		}

		result := convertSelectOptionsToTagProps(input)

		require.Len(s.T(), result, 2)
		require.Equal(s.T(), "tag1", result[0].Value)
		require.Equal(s.T(), "Tag 1", result[0].Text)
		require.True(s.T(), result[0].Selected)
		require.True(s.T(), result[1].Disabled)
	})

	s.Run("handles nil input", func() {
		result := convertSelectOptionsToTagProps(nil)
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

// TestEmailValidation tests email validation helper
func (s *FormFieldTestSuite) TestEmailValidation() {
	validEmails := []string{
		"user@example.com",
		"user.name@example.com",
		"user+tag@example.co.uk",
		"test_user@sub.domain.com",
		"123@example.com",
	}

	for _, email := range validEmails {
		s.Run("valid: "+email, func() {
			require.True(s.T(), isValidEmail(email))
		})
	}

	invalidEmails := []string{
		"invalid",
		"@example.com",
		"user@",
		"user @example.com",
		"user@example",
		"",
		"user@@example.com",
	}

	for _, email := range invalidEmails {
		s.Run("invalid: "+email, func() {
			require.False(s.T(), isValidEmail(email))
		})
	}
}

// TestURLValidation tests URL validation helper
func (s *FormFieldTestSuite) TestURLValidation() {
	validURLs := []string{
		"https://example.com",
		"http://example.com",
		"https://example.com/path",
		"https://sub.example.com",
		"https://example.com:8080",
		"https://example.com/path?query=value",
	}

	for _, url := range validURLs {
		s.Run("valid: "+url, func() {
			require.True(s.T(), isValidURL(url))
		})
	}

	invalidURLs := []string{
		"ftp://example.com",
		"example.com",
		"not a url",
		"",
		"javascript:alert('xss')",
		"//example.com",
	}

	for _, url := range invalidURLs {
		s.Run("invalid: "+url, func() {
			require.False(s.T(), isValidURL(url))
		})
	}
}

// TestGroupOptionsByName tests option grouping
func (s *FormFieldTestSuite) TestGroupOptionsByName() {
	s.Run("groups options correctly", func() {
		options := []SelectOption{
			{Value: "1", Label: "One", Group: "Numbers"},
			{Value: "2", Label: "Two", Group: "Numbers"},
			{Value: "a", Label: "A", Group: "Letters"},
			{Value: "3", Label: "Three"},
		}

		result := groupOptionsByName(options)

		require.Len(s.T(), result, 3)
		require.Len(s.T(), result["Numbers"], 2)
		require.Len(s.T(), result["Letters"], 1)
		require.Len(s.T(), result[""], 1)
	})

	s.Run("handles empty options", func() {
		result := groupOptionsByName([]SelectOption{})
		require.NotNil(s.T(), result)
		require.Empty(s.T(), result)
	})
}

// TestValidationEdgeCases tests edge cases in validation
func (s *FormFieldTestSuite) TestValidationEdgeCases() {
	s.Run("validate with whitespace-only value", func() {
		props := FormFieldProps{
			Type:     FormFieldText,
			Required: true,
		}
		err := props.Validate("   ")
		require.Error(s.T(), err)
	})

	s.Run("validate number with leading zeros", func() {
		props := FormFieldProps{
			Type: FormFieldNumber,
		}
		err := props.Validate("0123")
		require.NoError(s.T(), err)
	})

	s.Run("validate with invalid pattern regex", func() {
		props := FormFieldProps{
			Type:    FormFieldText,
			Pattern: "[invalid(",
		}
		err := props.Validate("test")
		require.Error(s.T(), err)
	})

	s.Run("validate email with special characters", func() {
		props := FormFieldProps{
			Type: FormFieldEmail,
		}
		err := props.Validate("user+tag@example.com")
		require.NoError(s.T(), err)
	})

	s.Run("validate URL with port", func() {
		props := FormFieldProps{
			Type: FormFieldURL,
		}
		err := props.Validate("https://example.com:8080")
		require.NoError(s.T(), err)
	})
}

// Benchmark tests
func BenchmarkValidate(b *testing.B) {
	props := FormFieldProps{
		Type:      FormFieldEmail,
		Required:  true,
		MinLength: 5,
		MaxLength: 100,
	}

	value := "user@example.com"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = props.Validate(value)
	}
}

func BenchmarkNormalizeDefaults(b *testing.B) {
	for i := 0; i < b.N; i++ {
		props := FormFieldProps{
			Type: FormFieldAutoComplete,
		}
		props.NormalizeDefaults()
	}
}

func BenchmarkConvertOptions(b *testing.B) {
	options := []SelectOption{
		{Value: "1", Label: "One"},
		{Value: "2", Label: "Two"},
		{Value: "3", Label: "Three"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = convertToRadioOptions(options)
	}
}

func BenchmarkValidateComplex(b *testing.B) {
	props := FormFieldProps{
		Type:      FormFieldText,
		Required:  true,
		MinLength: 5,
		MaxLength: 50,
		Pattern:   "^[a-zA-Z0-9_]+$",
		ValidationRules: []ValidationRule{
			{Type: "minLength", Value: 5},
			{Type: "pattern", Value: "^[a-zA-Z0-9_]+$"},
		},
	}

	value := "valid_username_123"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = props.Validate(value)
	}
}

func BenchmarkGetSelectedValues(b *testing.B) {
	props := FormFieldProps{
		Options: []SelectOption{
			{Value: "1", Selected: true},
			{Value: "2", Selected: false},
			{Value: "3", Selected: true},
			{Value: "4", Selected: false},
			{Value: "5", Selected: true},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = props.GetSelectedValues()
	}
}
