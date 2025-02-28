package qualifying

import (
	"encoding/hex"
	"errors"
)

// FixedXOR takes two equal-length hex-encoded strings and returns their XOR combination as a hex string.
func FixedXOR(hex1, hex2 string) (string, error) {
	// Decode hex strings to byte slices
	bytes1, err1 := hex.DecodeString(hex1)
	bytes2, err2 := hex.DecodeString(hex2)

	if err1 != nil || err2 != nil {
		return "", errors.New("invalid hex input")
	}

	// Ensure both inputs have the same length
	if len(bytes1) != len(bytes2) {
		return "", errors.New("input strings must have the same length")
	}

	// XOR operation
	xorResult := make([]byte, len(bytes1))
	for i := range bytes1 {
		xorResult[i] = bytes1[i] ^ bytes2[i]
	}

	// Convert the result to a hex string
	return hex.EncodeToString(xorResult), nil
}
