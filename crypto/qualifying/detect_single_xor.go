package qualifying

import (
	"bufio"
	"os"
)

// DetectSingleCharacterXOR reads a file and finds the line encrypted with single-character XOR
func DetectSingleCharacterXOR(filename string) (bestDecryption string, bestKey byte, bestLine string, bestScore float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", 0, "", 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bestScore = 0

	// Scan through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		decrypted, key, score, err := SingleByteXOR(line)
		if err != nil {
			continue // Skip invalid hex lines
		}

		// Keep track of the best decryption
		if score > bestScore {
			bestScore = score
			bestDecryption = decrypted
			bestKey = key
			bestLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		return "", 0, "", 0, err
	}

	return bestDecryption, bestKey, bestLine, bestScore, nil
}
