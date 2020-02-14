package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func main() {

}

func generateKeyPair(size int) (privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, size)
	if err != nil {
		log.Fatal(err)
	}

	publicKey = &privateKey.PublicKey

	return privateKey, publicKey
}
