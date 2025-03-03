package main

import (
	"crypto/qualifying"
	"encoding/hex"
	"fmt"
)

func hexToASCII(hexStr string) (string, error) {
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func main() {

	// // Challenge 3
	// hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	// ascii, err := hexToASCII(hexStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Decrypted hexStr: %s\n", ascii)

	// decrypted, key, score, err := qualifying.SingleByteXOR(hexStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Decrypted Message: %q\n", decrypted)
	// fmt.Printf("Key Used: %q (ASCII: %d)\n", key, key)
	// fmt.Printf("Score: %.2f\n", score)

	// // Challenge 4
	// filename := "qualifying/data/4.txt"

	// decrypted, key, line, score, err := qualifying.DetectSingleCharacterXOR(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Decrypted Message:", decrypted)
	// fmt.Printf("Key Used: %q (ASCII: %d)\n", key, key)
	// fmt.Println("Encrypted Line:", line)
	// fmt.Printf("Score: %.2f\n", score)

	// Challenge 5 - Repeating Key XOR
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"

	encrypted := qualifying.RepeatingKeyXOR(plaintext, key)

	fmt.Println("Encrypted Hex Output:")
	fmt.Println(encrypted)

}
