package main

import (
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// signJws sign jws with the given private key
func signJws(privateKey *rsa.PrivateKey) (token string, err error) {
	now := time.Now()

	// Prepare payload
	var payload = map[string]interface{}{
		"type":   "Consent",
		"id":     "37ce697c-83cd-4e67-8e78-2cc00b76cbe2",
		"status": "ConsentSubmitted",
		"updateTime": map[string]interface{}{
			"t":      now.Unix(),
			"humanT": now.Format("2006-01-02 15:04:05"),
		},
	}

	// Prepare token
	tkn := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"jti":       "2975dd35-6baa-47bc-9c29-be547b0b4cf8",
			"iat":       now.Unix(),
			"aud":       "JWTest",
			"issuer":    "CDI-PUSH-001",
			"issueTime": now.Format("2006-01-02T15:04:05.000"),
			"Payload":   payload,
		},
	)

	// Sign token
	token, err = tkn.SignedString(privateKey)
	if err != nil {
		err = fmt.Errorf("[JWS] %v", err)
	}
	return
}

func verifyJws(token string, publicKey *rsa.PublicKey) (header, claims map[string]interface{}, err error) {
	jws, err := jwt.Parse(
		token,
		func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		},
		jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}),
	)
	if err != nil {
		err = fmt.Errorf("[PARSE] %v", err)
		return
	}
	header = jws.Header
	if clms, ok := jws.Claims.(jwt.MapClaims); ok {
		claims = clms
	} else {
		claims = make(map[string]interface{})
		claims["claim"] = jws.Claims
	}
	return
}
