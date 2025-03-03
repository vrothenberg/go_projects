package qualifying

import (
	"encoding/hex"
	"unicode"
	// "fmt"
)

// Character frequency map based on English letter frequency
var letterFrequencies = map[rune]float64{
	'e': 12.70, 't': 9.06, 'a': 8.17, 'o': 7.51, 'i': 6.97, 'n': 6.75,
	's': 6.33, 'h': 6.09, 'r': 5.99, 'd': 4.25, 'l': 4.03, 'c': 2.78,
	'u': 2.76, 'm': 2.41, 'w': 2.36, 'f': 2.23, 'g': 2.02, 'y': 1.97,
	'p': 1.93, 'b': 1.29, 'v': 0.98, 'k': 0.77, 'j': 0.15, 'x': 0.15,
	'q': 0.10, 'z': 0.07, ' ': 15.00, // Spaces are very common in English
}

// scoreText computes how likely a given text is to be English based on character frequency
func scoreText(text string) float64 {
	score := 0.0
	for _, char := range text {
		if freq, exists := letterFrequencies[unicode.ToLower(char)]; exists {
			score += freq
		}
	}
	return score
}

// SingleByteXOR attempts to decrypt a hex-encoded string encrypted with a single-byte XOR cipher
func SingleByteXOR(hexStr string) (bestDecryption string, bestKey byte, bestScore float64, err error) {
	// Decode the hex string into bytes
	cipherBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", 0, 0, err
	}

	bestScore = 0
	bestDecryption = ""

	// Try all possible single-byte keys (0x00 - 0xFF)
	for key := 0; key < 256; key++ {
		decoded := make([]byte, len(cipherBytes))

		// XOR each byte with the candidate key
		for i := range cipherBytes {
			decoded[i] = cipherBytes[i] ^ byte(key)
		}

		// Convert to string
		plaintext := string(decoded)
		score := scoreText(plaintext)

		// Keep track of the best-scoring plaintext
		if score > bestScore {
			bestScore = score
			bestDecryption = plaintext
			bestKey = byte(key)
			// fmt.Printf("Best decryption: %s\n", bestDecryption)

		}
	}

	return bestDecryption, bestKey, bestScore, nil
}
