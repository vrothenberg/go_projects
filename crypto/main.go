package main

import (
	"crypto/qualifying"
	"encoding/hex"
	"fmt"
	"log"
)

func hexToASCII(hexStr string) (string, error) {
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func main() {
	hexStr := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	ascii, err := hexToASCII(hexStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted hexStr: %s\n", ascii)

	decrypted, key, score, err := qualifying.SingleByteXOR(hexStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Decrypted Message: %q\n", decrypted)
	fmt.Printf("Key Used: %q (ASCII: %d)\n", key, key)
	fmt.Printf("Score: %.2f\n", score)
}
