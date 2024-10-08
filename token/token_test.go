package token

import "testing"

func TestLookupWordTokenType(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput TokenType
	}{
		{"let", LET},
		{"fn", FUNC},
		{"xxx", IDENT},
		{"if", IF},
		{"else", ELSE},
		{"true", TRUE},
		{"false", FALSE},
		{"return", RETURN},
	}

	for i, tt := range tests {
		tokenType := LookupWordTokenType(tt.input)
		if tokenType != tt.expectedOutput {
			t.Fatalf("tests[%d] - expected token type %q got %q", i, tt.expectedOutput, tokenType)
		}
	}
}
