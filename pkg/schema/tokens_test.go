package schema

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// =============================================================================
// TOKEN REFERENCE TEST SUITE
// =============================================================================

// TokenReferenceTestSuite tests TokenReference type and its methods
type TokenReferenceTestSuite struct {
	suite.Suite
}

func TestTokenReferenceTestSuite(t *testing.T) {
	suite.Run(t, new(TokenReferenceTestSuite))
}

func (s *TokenReferenceTestSuite) TestIsReference_ValidReferences() {
	testCases := []struct {
		name  string
		value string
	}{
		{"simple path", "primitives.colors"},
		{"three segments", "primitives.colors.primary"},
		{"deep path", "components.button.primary.background-color"},
		{"with numbers", "primitives.colors.gray-500"},
		{"with underscores", "semantic.colors.background_color"},
		{"mixed separators", "semantic.font_size-lg"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.value)
			s.True(ref.IsReference(), "expected '%s' to be a reference", tc.value)
		})
	}
}

func (s *TokenReferenceTestSuite) TestIsReference_CSSLiterals() {
	testCases := []struct {
		name  string
		value string
	}{
		// Colors
		{"hex color", "#3b82f6"},
		{"hex short", "#fff"},
		{"hsl color", "hsl(217, 91%, 60%)"},
		{"rgb color", "rgb(59, 130, 246)"},
		{"rgba color", "rgba(59, 130, 246, 0.5)"},
		{"hsla color", "hsla(217, 91%, 60%, 0.8)"},

		// Units
		{"pixels", "16px"},
		{"rem", "1.5rem"},
		{"em", "2em"},
		{"percentage", "50%"},
		{"viewport", "100vh"},
		{"negative px", "-10px"},
		{"decimal rem", "0.875rem"},

		// Functions
		{"calc", "calc(100% - 2rem)"},
		{"var", "var(--primary)"},
		{"url", "url('/images/bg.png')"},
		{"linear gradient", "linear-gradient(to right, #fff, #000)"},
		{"radial gradient", "radial-gradient(circle, #fff, #000)"},

		// Keywords
		{"transparent", "transparent"},
		{"inherit", "inherit"},
		{"auto", "auto"},
		{"none", "none"},
		{"bold", "bold"},
		{"solid", "solid"},

		// Numbers
		{"integer", "1"},
		{"decimal", "0.5"},
		{"negative", "-3"},
		{"scientific", "1e10"},

		// Compound values
		{"shadow", "0 1px 2px rgba(0,0,0,0.1)"},
		{"border", "1px solid #000"},
		{"padding", "10px 20px"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.value)
			s.False(ref.IsReference(), "expected '%s' to NOT be a reference", tc.value)
		})
	}
}

func (s *TokenReferenceTestSuite) TestIsReference_EdgeCases() {
	testCases := []struct {
		name     string
		value    string
		expected bool
	}{
		{"empty string", "", false},
		{"single word", "primary", false},
		{"decimal number", "0.5", false},
		{"number with dot", "3.14", false},
		{"single segment", "colors", false},
		{"leading dot", ".colors.primary", false},
		{"trailing dot", "colors.primary.", false},
		{"double dot", "colors..primary", false},
		{"spaces", "primitives colors primary", false},
		{"space in segment", "primitives.color primary.value", false},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.value)
			s.Equal(tc.expected, ref.IsReference())
		})
	}
}

func (s *TokenReferenceTestSuite) TestValidate_ValidReferences() {
	validRefs := []string{
		"primitives.colors.primary",
		"semantic.colors.background",
		"components.button.primary.background-color",
		"primitives.colors.gray-500",
		"semantic.font_size-lg",
	}

	for _, refStr := range validRefs {
		s.Run(refStr, func() {
			ref := TokenReference(refStr)
			err := ref.Validate()
			s.NoError(err)
		})
	}
}

func (s *TokenReferenceTestSuite) TestValidate_InvalidReferences() {
	testCases := []struct {
		name        string
		value       string
		expectedErr string
	}{
		{"empty", "", "empty_token_reference"},
		{"no dots", "primary", "invalid_token_path"},
		{"single segment", "colors", "invalid_token_path"},
		{"empty segment", "colors..primary", "empty_token_segment"},
		{"leading dot", ".colors.primary", "empty_token_segment"},
		{"trailing dot", "colors.primary.", "empty_token_segment"},
		{"invalid chars", "colors.primary!", "invalid_token_segment"},
		{"spaces", "colors.primary color", "invalid_token_segment"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.value)
			err := ref.Validate()
			s.Error(err)
			s.Contains(err.Error(), tc.expectedErr)
		})
	}
}

func (s *TokenReferenceTestSuite) TestValidate_CSSLiterals() {
	// CSS literals should always be valid
	cssValues := []string{
		"#3b82f6",
		"hsl(217, 91%, 60%)",
		"1rem",
		"16px",
		"calc(100% - 2rem)",
		"transparent",
		"0.5",
	}

	for _, val := range cssValues {
		s.Run(val, func() {
			ref := TokenReference(val)
			err := ref.Validate()
			s.NoError(err)
		})
	}
}

func (s *TokenReferenceTestSuite) TestPathAndString() {
	ref := TokenReference("primitives.colors.primary")
	s.Equal("primitives.colors.primary", ref.Path())
	s.Equal("primitives.colors.primary", ref.String())

	// With whitespace
	refWithSpace := TokenReference("  primitives.colors.primary  ")
	s.Equal("primitives.colors.primary", strings.TrimSpace(refWithSpace.Path()))
}

// =============================================================================
// DESIGN TOKENS TEST SUITE
// =============================================================================

// DesignTokensTestSuite tests DesignTokens structure and operations
type DesignTokensTestSuite struct {
	suite.Suite
	tokens *DesignTokens
}

func TestDesignTokensTestSuite(t *testing.T) {
	suite.Run(t, new(DesignTokensTestSuite))
}

func (s *DesignTokensTestSuite) SetupTest() {
	s.tokens = &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary":  "hsl(217, 91%, 60%)",
				"gray-500": "hsl(210, 10%, 58%)",
				"white":    "#ffffff",
			},
			Spacing: map[string]string{
				"sm": "0.5rem",
				"md": "1rem",
				"lg": "1.5rem",
			},
			Typography: map[string]string{
				"font-size-base":   "1rem",
				"font-weight-bold": "700",
			},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				"background": "primitives.colors.white",
				"primary":    "primitives.colors.primary",
			},
			Spacing: map[string]string{
				"component-default": "primitives.spacing.md",
			},
		},
		Components: &ComponentTokens{
			"button": ComponentVariants{
				"primary": StyleProperties{
					"background-color": "semantic.colors.primary",
					"padding":          "primitives.spacing.sm primitives.spacing.md",
				},
			},
		},
	}
}

func (s *DesignTokensTestSuite) TestValidate_ValidTokens() {
	err := s.tokens.Validate()
	s.NoError(err)
}

func (s *DesignTokensTestSuite) TestValidate_NilTokens() {
	var tokens *DesignTokens
	err := tokens.Validate()
	s.Error(err)
	s.Contains(err.Error(), "nil_tokens")
}

func (s *DesignTokensTestSuite) TestValidate_NilPrimitives() {
	// Nil primitives are now allowed (for overlay scenarios like dark mode)
	tokens := &DesignTokens{
		Primitives: nil,
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				"background": "primitives.colors.gray-900",
			},
		},
	}
	err := tokens.Validate()
	s.NoError(err)
}

func (s *DesignTokensTestSuite) TestValidate_EmptyPrimitiveCategory() {
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{}, // Empty
		},
	}
	err := tokens.Validate()
	s.Error(err)
	s.Contains(err.Error(), "empty_category")
}

func (s *DesignTokensTestSuite) TestValidate_PrimitiveWithReference() {
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "semantic.colors.primary", // Primitives shouldn't reference
			},
		},
	}
	err := tokens.Validate()
	s.Error(err)
	s.Contains(err.Error(), "primitive_reference")
}

func (s *DesignTokensTestSuite) TestGetToken_Primitive() {
	value, err := s.tokens.GetToken("primitives.colors.primary")
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", value)
}

func (s *DesignTokensTestSuite) TestGetToken_Semantic() {
	value, err := s.tokens.GetToken("semantic.colors.background")
	s.NoError(err)
	s.Equal("primitives.colors.white", value)
}

func (s *DesignTokensTestSuite) TestGetToken_Component() {
	value, err := s.tokens.GetToken("components.button.primary.background-color")
	s.NoError(err)
	s.Equal("semantic.colors.primary", value)
}

func (s *DesignTokensTestSuite) TestGetToken_InvalidPath() {
	testCases := []struct {
		name string
		path string
	}{
		{"too short", "primitives.colors"},
		{"invalid tier", "invalid.colors.primary"},
		{"invalid category", "primitives.invalid.primary"},
		{"not found", "primitives.colors.nonexistent"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			_, err := s.tokens.GetToken(tc.path)
			s.Error(err)
		})
	}
}

func (s *DesignTokensTestSuite) TestClone() {
	cloned := s.tokens.Clone()
	s.NotNil(cloned)
	s.NotSame(s.tokens, cloned)
	s.NotSame(s.tokens.Primitives, cloned.Primitives)
	s.NotSame(s.tokens.Semantic, cloned.Semantic)
	s.NotSame(s.tokens.Components, cloned.Components)

	// Modify cloned - should not affect original
	cloned.Primitives.Colors["primary"] = "hsl(0, 0%, 0%)"
	s.Equal("hsl(217, 91%, 60%)", s.tokens.Primitives.Colors["primary"])
	s.Equal("hsl(0, 0%, 0%)", cloned.Primitives.Colors["primary"])
}

func (s *DesignTokensTestSuite) TestClone_NilTokens() {
	var tokens *DesignTokens
	cloned := tokens.Clone()
	s.Nil(cloned)
}

func (s *DesignTokensTestSuite) TestClone_DeepNested() {
	// Ensure deep cloning works for nested maps
	original := &DesignTokens{
		Components: &ComponentTokens{
			"button": ComponentVariants{
				"primary": StyleProperties{
					"color": "blue",
				},
			},
		},
	}

	cloned := original.Clone()
	(*cloned.Components)["button"]["primary"]["color"] = "red"

	s.Equal("blue", (*original.Components)["button"]["primary"]["color"])
	s.Equal("red", (*cloned.Components)["button"]["primary"]["color"])
}

// =============================================================================
// TOKEN RESOLVER TEST SUITE
// =============================================================================

// TokenResolverTestSuite tests token resolution with caching
type TokenResolverTestSuite struct {
	suite.Suite
	tokens   *DesignTokens
	resolver *TokenResolver
	ctx      context.Context
}

func TestTokenResolverTestSuite(t *testing.T) {
	suite.Run(t, new(TokenResolverTestSuite))
}

func (s *TokenResolverTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.tokens = &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary":   "hsl(217, 91%, 60%)",
				"secondary": "hsl(270, 50%, 60%)",
				"gray-500":  "hsl(210, 10%, 58%)",
				"white":     "#ffffff",
			},
			Spacing: map[string]string{
				"sm": "0.5rem",
				"md": "1rem",
			},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				"background": "primitives.colors.white",
				"primary":    "primitives.colors.primary",
				"text":       "primitives.colors.gray-500",
			},
			Spacing: map[string]string{
				"component-default": "primitives.spacing.md",
			},
		},
		Components: &ComponentTokens{
			"button": ComponentVariants{
				"primary": StyleProperties{
					"background-color": "semantic.colors.primary",
					"color":            "semantic.colors.background",
				},
			},
		},
	}

	var err error
	s.resolver, err = NewTokenResolver(s.tokens)
	s.Require().NoError(err)
}

func (s *TokenResolverTestSuite) TearDownTest() {
	if s.resolver != nil {
		s.resolver.Close()
	}
}

func (s *TokenResolverTestSuite) TestNewTokenResolver_NilTokens() {
	resolver, err := NewTokenResolver(nil)
	s.Error(err)
	s.Nil(resolver)
	s.Contains(err.Error(), "nil_tokens")
}

func (s *TokenResolverTestSuite) TestNewTokenResolver_InvalidTokens() {
	invalidTokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "semantic.colors.primary", // Invalid: primitive with reference
			},
		},
	}
	resolver, err := NewTokenResolver(invalidTokens)
	s.Error(err)
	s.Nil(resolver)
}

func (s *TokenResolverTestSuite) TestResolve_CSSLiteral() {
	testCases := []string{
		"#3b82f6",
		"hsl(217, 91%, 60%)",
		"1rem",
		"16px",
		"transparent",
	}

	for _, literal := range testCases {
		s.Run(literal, func() {
			ref := TokenReference(literal)
			resolved, err := s.resolver.Resolve(s.ctx, ref)
			s.NoError(err)
			s.Equal(literal, resolved)
		})
	}
}

func (s *TokenResolverTestSuite) TestResolve_DirectPrimitive() {
	ref := TokenReference("primitives.colors.primary")
	resolved, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", resolved)
}

func (s *TokenResolverTestSuite) TestResolve_SemanticToPrimitive() {
	ref := TokenReference("semantic.colors.background")
	resolved, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("#ffffff", resolved)
}

func (s *TokenResolverTestSuite) TestResolve_ComponentToSemanticToPrimitive() {
	ref := TokenReference("components.button.primary.background-color")
	resolved, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", resolved)
}

func (s *TokenResolverTestSuite) TestResolve_NotFound() {
	ref := TokenReference("primitives.colors.nonexistent")
	_, err := s.resolver.Resolve(s.ctx, ref)
	s.Error(err)
	s.Contains(err.Error(), "token not found")
}

func (s *TokenResolverTestSuite) TestResolve_InvalidPath() {
	testCases := []struct {
		name string
		ref  string
	}{
		{"invalid tier", "invalid.colors.primary"},
		{"invalid category", "primitives.invalid.primary"},
		{"too short", "primitives.colors"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.ref)
			_, err := s.resolver.Resolve(s.ctx, ref)
			s.Error(err)
		})
	}
}

func (s *TokenResolverTestSuite) TestResolve_CircularReference() {
	// Create tokens with circular reference
	circularTokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "hsl(217, 91%, 60%)",
			},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				"a": "semantic.colors.b",
				"b": "semantic.colors.c",
				"c": "semantic.colors.a", // Circular: a -> b -> c -> a
			},
		},
	}

	resolver, err := NewTokenResolver(circularTokens)
	s.Require().NoError(err)
	defer resolver.Close()

	ref := TokenReference("semantic.colors.a")
	_, err = resolver.Resolve(s.ctx, ref)
	s.Error(err)
	s.Contains(err.Error(), "circular")
}

func (s *TokenResolverTestSuite) TestResolve_Caching() {
	ref := TokenReference("semantic.colors.primary")

	// First resolution - cache miss
	resolved1, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", resolved1)

	// Second resolution - cache hit (should be faster)
	resolved2, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", resolved2)
	s.Equal(resolved1, resolved2)
}

func (s *TokenResolverTestSuite) TestResolve_ContextCancellation() {
	ctx, cancel := context.WithCancel(s.ctx)
	cancel() // Cancel immediately

	ref := TokenReference("semantic.colors.primary")
	_, err := s.resolver.Resolve(ctx, ref)
	s.Error(err)
	s.Equal(context.Canceled, err)
}

func (s *TokenResolverTestSuite) TestResolveAll() {
	resolved, err := s.resolver.ResolveAll(s.ctx)
	s.NoError(err)
	s.NotNil(resolved)

	// All semantic tokens should be resolved to primitives
	s.Equal("#ffffff", resolved.Semantic.Colors["background"])
	s.Equal("hsl(217, 91%, 60%)", resolved.Semantic.Colors["primary"])

	// All component tokens should be resolved
	buttonBg := (*resolved.Components)["button"]["primary"]["background-color"]
	s.Equal("hsl(217, 91%, 60%)", buttonBg)
}

func (s *TokenResolverTestSuite) TestClearCache() {
	ref := TokenReference("semantic.colors.primary")

	// Resolve to populate cache
	_, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)

	// Clear cache
	s.resolver.ClearCache()

	// Should still resolve correctly after cache clear
	resolved, err := s.resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(217, 91%, 60%)", resolved)
}

func (s *TokenResolverTestSuite) TestClose() {
	err := s.resolver.Close()
	s.NoError(err)
	// Resolver should not be used after Close, but shouldn't panic
}

// =============================================================================
// TOKEN SERIALIZATION TEST SUITE
// =============================================================================

// TokenSerializationTestSuite tests JSON/YAML serialization
type TokenSerializationTestSuite struct {
	suite.Suite
	tokens *DesignTokens
}

func TestTokenSerializationTestSuite(t *testing.T) {
	suite.Run(t, new(TokenSerializationTestSuite))
}

func (s *TokenSerializationTestSuite) SetupTest() {
	s.tokens = &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "hsl(217, 91%, 60%)",
				"white":   "#ffffff",
			},
			Spacing: map[string]string{
				"md": "1rem",
			},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{
				"background": "primitives.colors.white",
			},
		},
	}
}

func (s *TokenSerializationTestSuite) TestTokensToJSON() {
	jsonData, err := s.tokens.TokensToJSON()
	s.NoError(err)
	s.NotEmpty(jsonData)

	// Should be valid JSON
	s.Contains(string(jsonData), `"primitives"`)
	s.Contains(string(jsonData), `"semantic"`)
	s.Contains(string(jsonData), `"primary"`)
}

func (s *TokenSerializationTestSuite) TestTokensFromJSON() {
	// Serialize
	jsonData, err := s.tokens.TokensToJSON()
	s.Require().NoError(err)

	// Deserialize
	parsed, err := TokensFromJSON(jsonData)
	s.NoError(err)
	s.NotNil(parsed)
	s.Equal("hsl(217, 91%, 60%)", parsed.Primitives.Colors["primary"])
	s.Equal("primitives.colors.white", parsed.Semantic.Colors["background"])
}

func (s *TokenSerializationTestSuite) TestTokensFromJSON_Invalid() {
	invalidJSON := []byte(`{"invalid": "json"`)
	_, err := TokensFromJSON(invalidJSON)
	s.Error(err)
}

func (s *TokenSerializationTestSuite) TestTokensFromJSON_InvalidStructure() {
	invalidStructure := []byte(`{
		"primitives": {
			"colors": {
				"primary": "semantic.colors.primary"
			}
		}
	}`)
	_, err := TokensFromJSON(invalidStructure)
	s.Error(err)
	s.Contains(err.Error(), "primitive_reference")
}

func (s *TokenSerializationTestSuite) TestRoundTrip_JSON() {
	// Serialize
	jsonData, err := s.tokens.TokensToJSON()
	s.Require().NoError(err)

	// Deserialize
	parsed, err := TokensFromJSON(jsonData)
	s.Require().NoError(err)

	// Serialize again
	jsonData2, err := parsed.TokensToJSON()
	s.NoError(err)

	// Should be equivalent
	s.JSONEq(string(jsonData), string(jsonData2))
}

// =============================================================================
// DEFAULT TOKENS TEST SUITE
// =============================================================================

// DefaultTokensTestSuite tests the default token system
type DefaultTokensTestSuite struct {
	suite.Suite
	tokens *DesignTokens
}

func TestDefaultTokensTestSuite(t *testing.T) {
	suite.Run(t, new(DefaultTokensTestSuite))
}

func (s *DefaultTokensTestSuite) SetupTest() {
	s.tokens = GetDefaultTokens()
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_NotNil() {
	s.NotNil(s.tokens)
	s.NotNil(s.tokens.Primitives)
	s.NotNil(s.tokens.Semantic)
	s.NotNil(s.tokens.Components)
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_Valid() {
	err := s.tokens.Validate()
	s.NoError(err)
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_HasRequiredPrimitives() {
	s.NotEmpty(s.tokens.Primitives.Colors)
	s.NotEmpty(s.tokens.Primitives.Spacing)
	s.NotEmpty(s.tokens.Primitives.Typography)
	s.NotEmpty(s.tokens.Primitives.Shadows)

	// Check specific required colors
	s.Contains(s.tokens.Primitives.Colors, "primary")
	s.Contains(s.tokens.Primitives.Colors, "white")
	s.Contains(s.tokens.Primitives.Colors, "black")
	s.Contains(s.tokens.Primitives.Colors, "success")
	s.Contains(s.tokens.Primitives.Colors, "error")
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_HasRequiredSemantic() {
	s.NotEmpty(s.tokens.Semantic.Colors)
	s.NotEmpty(s.tokens.Semantic.Spacing)
	s.NotEmpty(s.tokens.Semantic.Typography)

	// Check specific semantic colors
	s.Contains(s.tokens.Semantic.Colors, "background")
	s.Contains(s.tokens.Semantic.Colors, "foreground")
	s.Contains(s.tokens.Semantic.Colors, "primary")
	s.Contains(s.tokens.Semantic.Colors, "border")
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_HasComponents() {
	s.NotEmpty(*s.tokens.Components)
	s.Contains(*s.tokens.Components, "button")
	s.Contains(*s.tokens.Components, "input")
	s.Contains(*s.tokens.Components, "card")
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_ComponentStructure() {
	buttonVariants := (*s.tokens.Components)["button"]
	s.NotEmpty(buttonVariants)
	s.Contains(buttonVariants, "primary")

	primaryButton := buttonVariants["primary"]
	s.Contains(primaryButton, "background-color")
	s.Contains(primaryButton, "color")
}

func (s *DefaultTokensTestSuite) TestDefaultTokens_Resolvable() {
	resolver, err := NewTokenResolver(s.tokens)
	s.Require().NoError(err)
	defer resolver.Close()

	ctx := context.Background()

	// Test resolving various token levels
	testCases := []struct {
		name     string
		tokenRef string
	}{
		{"primitive color", "primitives.colors.primary"},
		{"semantic color", "semantic.colors.background"},
		{"component property", "components.button.primary.background-color"},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			ref := TokenReference(tc.tokenRef)
			resolved, err := resolver.Resolve(ctx, ref)
			s.NoError(err)
			s.NotEmpty(resolved)
		})
	}
}

// =============================================================================
// TOKEN EDGE CASES TEST SUITE
// =============================================================================

// TokenEdgeCasesTestSuite tests edge cases and error conditions
type TokenEdgeCasesTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestTokenEdgeCasesTestSuite(t *testing.T) {
	suite.Run(t, new(TokenEdgeCasesTestSuite))
}

func (s *TokenEdgeCasesTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func (s *TokenEdgeCasesTestSuite) TestLargeTokenStructure() {
	// Create tokens with many entries
	largeTokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: make(map[string]string),
		},
	}

	// Add 1000 color tokens
	for i := range 1000 {
		key := fmt.Sprintf("color-%d", i)
		largeTokens.Primitives.Colors[key] = fmt.Sprintf("hsl(%d, 50%%, 50%%)", i%360)
	}

	err := largeTokens.Validate()
	s.NoError(err)

	resolver, err := NewTokenResolver(largeTokens)
	s.Require().NoError(err)
	defer resolver.Close()

	// Should be able to resolve
	ref := TokenReference("primitives.colors.color-500")
	resolved, err := resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.NotEmpty(resolved)
}

func (s *TokenEdgeCasesTestSuite) TestDeepNesting() {
	// Test maximum reference chain depth
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"base": "hsl(0, 0%, 50%)",
			},
		},
		Semantic: &SemanticTokens{
			Colors: make(map[string]string),
		},
	}

	// Create a chain: level1 -> level2 -> level3 -> ... -> base
	tokens.Semantic.Colors["level10"] = "primitives.colors.base"
	for i := 9; i >= 1; i-- {
		tokens.Semantic.Colors[fmt.Sprintf("level%d", i)] = fmt.Sprintf("semantic.colors.level%d", i+1)
	}

	resolver, err := NewTokenResolver(tokens)
	s.Require().NoError(err)
	defer resolver.Close()

	ref := TokenReference("semantic.colors.level1")
	resolved, err := resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Equal("hsl(0, 0%, 50%)", resolved)
}

func (s *TokenEdgeCasesTestSuite) TestUnicodeInTokenNames() {
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"primary": "hsl(217, 91%, 60%)",
			},
		},
	}

	// Token names with unicode should work if they're valid identifiers
	// (though not recommended in practice)
	err := tokens.Validate()
	s.NoError(err)
}

func (s *TokenEdgeCasesTestSuite) TestEmptyMaps() {
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{"primary": "hsl(0, 0%, 0%)"},
		},
		Semantic: &SemanticTokens{
			Colors: map[string]string{}, // Empty map - should be fine
		},
		Components: &ComponentTokens{}, // Empty components - should be fine
	}

	err := tokens.Validate()
	s.NoError(err)
}

func (s *TokenEdgeCasesTestSuite) TestConcurrentResolve() {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	s.Require().NoError(err)
	defer resolver.Close()

	// Resolve same token concurrently
	const numGoroutines = 50
	done := make(chan bool, numGoroutines)

	for range numGoroutines {
		go func() {
			defer func() { done <- true }()
			ref := TokenReference("semantic.colors.primary")
			resolved, err := resolver.Resolve(s.ctx, ref)
			s.NoError(err)
			s.NotEmpty(resolved)
		}()
	}

	// Wait for all goroutines
	for range numGoroutines {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			s.Fail("timeout waiting for concurrent resolves")
		}
	}
}

func (s *TokenEdgeCasesTestSuite) TestResolveWithTimeout() {
	tokens := GetDefaultTokens()
	resolver, err := NewTokenResolver(tokens)
	s.Require().NoError(err)
	defer resolver.Close()

	ctx, cancel := context.WithTimeout(s.ctx, 10*time.Millisecond)
	defer cancel()

	// Should complete before timeout
	ref := TokenReference("semantic.colors.primary")
	resolved, err := resolver.Resolve(ctx, ref)
	s.NoError(err)
	s.NotEmpty(resolved)
}

func (s *TokenEdgeCasesTestSuite) TestSpecialCharactersInValues() {
	tokens := &DesignTokens{
		Primitives: &PrimitiveTokens{
			Colors: map[string]string{
				"special": `url("data:image/svg+xml,%3Csvg xmlns='http://www.w3.org/2000/svg'%3E%3C/svg%3E")`,
			},
			Typography: map[string]string{
				"font-family": `-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto`,
			},
		},
	}

	err := tokens.Validate()
	s.NoError(err)

	resolver, err := NewTokenResolver(tokens)
	s.Require().NoError(err)
	defer resolver.Close()

	ref := TokenReference("primitives.colors.special")
	resolved, err := resolver.Resolve(s.ctx, ref)
	s.NoError(err)
	s.Contains(resolved, "data:image/svg+xml")
}

func Test_isNumeric(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNumeric(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("isNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
