package qualifying

import (
	"encoding/base64"
	"errors"
	// "fmt"
	"math/bits"
	"os"
	"sort"
	"strings"
)

// HammingDistance calculates the number of differing bits between two byte slices.
// For example, the distance between "this is a test" and "wokka wokka!!!" should be 37.
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("slices must be the same length")
	}
	distance := 0
	for i := range a {
		distance += bits.OnesCount8(a[i] ^ b[i])
	}
	return distance, nil
}

// candidateKeySize holds a candidate key size and its normalized edit (Hamming) distance.
type candidateKeySize struct {
	keySize      int
	normDistance float64
}

// ByNormDistance implements sort.Interface for []candidateKeySize based on normDistance.
type ByNormDistance []candidateKeySize

func (a ByNormDistance) Len() int           { return len(a) }
func (a ByNormDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNormDistance) Less(i, j int) bool { return a[i].normDistance < a[j].normDistance }

// SingleByteXORBytes tries all 256 possible single-byte XOR keys on a byte slice.
// It returns the key that produces the best score (using scoreText), along with that score
// and the resulting plaintext bytes.
func SingleByteXORBytes(input []byte) (bestKey byte, bestScore float64, bestPlaintext []byte) {
	bestScore = -1.0
	for key := 0; key < 256; key++ {
		candidate := make([]byte, len(input))
		for i, b := range input {
			candidate[i] = b ^ byte(key)
		}
		score := scoreText(string(candidate))
		if score > bestScore {
			bestScore = score
			bestKey = byte(key)
			bestPlaintext = candidate
		}
	}
	return bestKey, bestScore, bestPlaintext
}

// BreakRepeatingKeyXOR reads a file containing a base64-encoded ciphertext that was encrypted
// with repeating-key XOR. It strips newlines from the base64 text, decodes the ciphertext,
// and then returns the discovered key and the decrypted plaintext.
func BreakRepeatingKeyXOR(filename string) (string, string, error) {
	// Read the file content.
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", "", err
	}

	// Remove newlines (and any extra whitespace) to form a contiguous base64 string.
	base64Str := strings.ReplaceAll(string(data), "\n", "")
	base64Str = strings.TrimSpace(base64Str)

	// Decode the base64 string to get the raw ciphertext.
	ciphertext, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", "", err
	}

	var candidates []candidateKeySize

	// Use up to 4 blocks for averaging.
	numBlocksToUse := 4

	// Try key sizes from 2 to 40.
	for keySize := 2; keySize <= 40; keySize++ {
		// Ensure we have enough data.
		if len(ciphertext) < keySize*numBlocksToUse {
			continue
		}

		totalDistance := 0.0
		count := 0
		// Compare every pair among the first numBlocksToUse blocks.
		for i := 0; i < numBlocksToUse-1; i++ {
			for j := i + 1; j < numBlocksToUse; j++ {
				block1 := ciphertext[i*keySize : (i+1)*keySize]
				block2 := ciphertext[j*keySize : (j+1)*keySize]
				hd, err := HammingDistance(block1, block2)
				if err != nil {
					return "", "", err
				}
				totalDistance += float64(hd)
				count++
			}
		}
		avgDistance := totalDistance / float64(count)
		normDist := avgDistance / float64(keySize)
		candidates = append(candidates, candidateKeySize{keySize: keySize, normDistance: normDist})
	}

	if len(candidates) == 0 {
		return "", "", errors.New("no candidate key sizes found")
	}

	// Sort candidate key sizes by normalized distance (lowest first).
	sort.Sort(ByNormDistance(candidates))
	bestKeySize := candidates[0].keySize

	// Break the ciphertext into blocks of bestKeySize and then transpose them.
	transposed := make([][]byte, bestKeySize)
	for i := 0; i < bestKeySize; i++ {
		transposed[i] = []byte{}
		for j := i; j < len(ciphertext); j += bestKeySize {
			transposed[i] = append(transposed[i], ciphertext[j])
		}
	}

	// For each transposed block, solve the single-byte XOR to determine the key byte.
	keyBytes := make([]byte, bestKeySize)
	for i, block := range transposed {
		keyByte, _, _ := SingleByteXORBytes(block)
		keyBytes[i] = keyByte
	}
	key := string(keyBytes)

	// Decrypt the ciphertext using the discovered repeating key.
	plaintextBytes := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		plaintextBytes[i] = b ^ keyBytes[i%bestKeySize]
	}
	plaintext := string(plaintextBytes)

	return key, plaintext, nil
}