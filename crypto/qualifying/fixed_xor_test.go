package qualifying

import (
	"testing"
)

func TestFixedXOR(t *testing.T) {
	tests := []struct {
		name     string
		hex1     string
		hex2     string
		expected string
		hasError bool
	}{
		{
			name:     "Crypto Challenge 2",
			hex1:     "1c0111001f010100061a024b53535009181c",
			hex2:     "686974207468652062756c6c277320657965",
			expected: "746865206b696420646f6e277420706c6179",
			hasError: false,
		},
		{
			name:     "Different Length Inputs",
			hex1:     "1c0111001f01",
			hex2:     "68697420746865",
			expected: "",
			hasError: true,
		},
		{
			name:     "Invalid Hex Input",
			hex1:     "xyz123",
			hex2:     "abcdef",
			expected: "",
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Testing: %s", tt.name)
			result, err := FixedXOR(tt.hex1, tt.hex2)

			if tt.hasError {
				if err == nil {
					t.Fatalf("Expected error for input (%q, %q), but got none", tt.hex1, tt.hex2)
				} else {
					t.Logf("Received expected error: %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error for input (%q, %q): %v", tt.hex1, tt.hex2, err)
				}
				if result != tt.expected {
					t.Fatalf("FixedXOR(%q, %q) = %q, expected %q", tt.hex1, tt.hex2, result, tt.expected)
				} else {
					t.Logf("Success! Output: %q", result)
				}
			}
		})
	}
}
