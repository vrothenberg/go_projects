package qualifying

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	tests := []struct {
		name     string
		hexStr   string
		expected string
		hasError bool
	}{
		{
			name:     "Crypto Challenge 1",
			hexStr:   "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			expected: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
			hasError: false,
		},
		{
			name:     "Simple Hello",
			hexStr:   "48656c6c6f",
			expected: "SGVsbG8=",
			hasError: false,
		},
		{
			name:     "Simple World",
			hexStr:   "776f726c64",
			expected: "d29ybGQ=",
			hasError: false,
		},
		{
			name:     "Hello World",
			hexStr:   "48656c6c6f20776f726c64",
			expected: "SGVsbG8gd29ybGQ=",
			hasError: false,
		},
		{
			name:     "Invalid Hex Input",
			hexStr:   "gibberish",
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing: %s", tt.name)
			result, err := HexToBase64(tt.hexStr)

			if tt.hasError {
				if err == nil {
					t.Fatalf("Expected error for input %q, but got none", tt.hexStr)
				} else {
					t.Logf("Received expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error for input %q: %v", tt.hexStr, err)
				}
				if result != tt.expected {
					t.Fatalf("HexToBase64(%q) = %q, expected %q", tt.hexStr, result, tt.expected)
				} else {
					t.Logf("Success! Output: %q", result)
				}
			}
		})
	}
}
