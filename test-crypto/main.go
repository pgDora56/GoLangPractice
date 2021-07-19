package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
)

type Config struct {
	Public   string   `json:"public"`
	PublicEx int      `json:"public-ex"`
	D        string   `json:"privateD"`
	Primes   []string `json:"primes"`
}

func main() {
	bytes, err := ioutil.ReadFile("key.json")
	if err != nil {
		log.Fatal("Raise error reading JSON")
	}
	var config Config
	if err := json.Unmarshal(bytes, &config); err != nil {
		log.Fatalf("Raise error unmarshal: %v\n", err)
	}
	// log.Printf("E: %d", config.PublicEx)
	// log.Print(config)
	// rsaCryptoTest()
	rsaCryptoJSON(config)
}

func aesCryptoTest() {
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

func rsaCryptoJSON(cfg Config) {
	plainText := []byte("葉加瀬冬雪")
	pubn := new(big.Int)
	pubn.SetString(cfg.Public, 10)
	publicKey := rsa.PublicKey{N: pubn, E: cfg.PublicEx}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &publicKey, plainText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	fmt.Printf("Cipher text: %x\n", cipherText)

	prvD := new(big.Int)
	private1 := new(big.Int)
	private2 := new(big.Int)
	prvD.SetString(cfg.D, 10)
	private1.SetString(cfg.Primes[0], 10)
	private2.SetString(cfg.Primes[1], 10)

	privateKey := rsa.PrivateKey{
		PublicKey: publicKey,
		D:         prvD,
		Primes:    []*big.Int{private1, private2},
	}
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, &privateKey, cipherText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}

func rsaCryptoTest() {
	plainText := []byte("葉加瀬冬雪")

	// size of key (bits)
	size := 4096

	// Generate private and public key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}
	log.Printf("Public Key: %v\n\nD:%v\n\nPrimes:%v\n\n", &privateKey.PublicKey, &privateKey.D, &privateKey.Primes)

	// Get public key from private key and encrypt
	publicKey := &privateKey.PublicKey

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}
	fmt.Printf("Cipher text: %x\n", cipherText)

	// Decrypt with private key
	decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		fmt.Printf("Err: %s\n", err)
		return
	}

	fmt.Printf("Decrypted text: %s\n", decryptedText)
}
