package theme

import (
	"bytes"
	"testing"
)

func TestCompiler_generateFontProperties(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		config *CompilerConfig
		// Named input parameters for target function.
		buf    *bytes.Buffer
		tokens *Tokens
		theme  *Theme
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCompiler(tt.config)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			c.generateFontProperties(tt.buf, tt.tokens, tt.theme)
		})
	}
}
