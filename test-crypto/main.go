package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	plainText := []byte("This is 16 bytes")

	key := []byte("passw0rdpassw0rdpassw0rdpassw0rd")

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	// Encrypt
	cipherText := make([]byte, len(plainText))
	block.Encrypt(cipherText, plainText)
	fmt.Printf("Cipher text: %x\n", cipherText)

	// Decrypt
	decryptedText := make([]byte, len(cipherText))
	block.Decrypt(decryptedText, cipherText)
	fmt.Printf("Decrypted text: %s\n", string(decryptedText))
}
