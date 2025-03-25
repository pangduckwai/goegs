package main

import (
	"crypto/rsa"
	"fmt"
	"log"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

func run(
	prvkeys, prvkeyr *rsa.PrivateKey,
	pubkeys, pubkeyr *rsa.PublicKey,
) {
	// Sign
	token, err := signJws(prvkeys)
	if err != nil {
		log.Fatalf("[SIGN] %v", err)
	}
	fmt.Printf("JWS:\n%v\n\n", token)

	buff := []byte(token)

	encrypted, err := jwe.Encrypt(buff, jwe.WithKey(jwa.RSA_OAEP_256, &prvkeyr.PublicKey), jwe.WithContentEncryption(jwa.A128CBC_HS256))
	if err != nil {
		log.Fatalf("[ENC] %v", err)
	}
	fmt.Printf("JWE:\n%s\n\n", encrypted)

	decrypted, err := jwe.Decrypt(encrypted, jwe.WithKey(jwa.RSA_OAEP_256, prvkeyr))
	if err != nil {
		log.Fatalf("[DEC] %v", err)
	}
	fmt.Printf("MSG:\n%s\n\n", decrypted)
}
