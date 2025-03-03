package qualifying

import (
	"encoding/hex"
)

// RepeatingKeyXOR encrypts plaintext using a repeating key XOR
func RepeatingKeyXOR(plaintext, key string) string {
	keyBytes := []byte(key)
	plaintextBytes := []byte(plaintext)
	ciphertext := make([]byte, len(plaintextBytes))

	// XOR each byte of plaintext with the repeating key
	for i := range plaintextBytes {
		ciphertext[i] = plaintextBytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	// Convert result to hex string
	return hex.EncodeToString(ciphertext)
}
