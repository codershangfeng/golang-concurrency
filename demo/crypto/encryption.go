package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

func encrypt() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PublicKey

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("abcd"),
		nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("encrypted bytes: ", encryptedBytes)

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}

	// We get back the original information in the form of bytes, which we
	// the cast to a string and print
	fmt.Println("decrypted message: ", string(decryptedBytes))
}
