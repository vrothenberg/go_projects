package qualifying

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// HexToBase64 converts a hexadecimal string to a base64-encoded string
func HexToBase64(hexStr string) (string, error) {
	// Decode hex string
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", errors.New("invalid hex input")
	}
	
	// Encode to base64
	return base64.StdEncoding.EncodeToString(bytes), nil
}
