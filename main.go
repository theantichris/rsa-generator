package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
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

func exportToPEM(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (privatePEM, publicPEM string) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEMBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	privatePEM = string(privatePEMBytes)

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	publicPEMBytes := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})

	publicPEM = string(publicPEMBytes)

	return privatePEM, publicPEM
}

func writeFile(content []byte, fileName string) {
	filePath := fmt.Sprintf("%s", fileName)

	err := ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
