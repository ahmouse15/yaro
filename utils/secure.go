package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"math/rand"
)

// Generates a one-time use number. Outputs between 299,999,999 and -300,000,000 (inclusive)
func GenerateNonce() int {
	return rand.Intn(600000000) - 300000000
}

// Generates SHA256 hash of given string and returns Base64URL-encoded result
func GenerateSha256(input string) string {

	hash := sha256.Sum256([]byte(input))
	return ToBase64Url(string(hash[:]))
}

// Generates a valid code challenge from the given verifier
func GenerateCodeChallenge(verifier string) string {
	return GenerateSha256(verifier)
}

// Generates a code for OAuth verification
func GenerateCodeVerifier() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-._~")
	length := 64

	code := make([]rune, length)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	result := ToBase64Url(string(code))

	return result
}

// Generate HMAC signature for given payload in Base64 form.
func GenerateHmac(key string, payload string) string {
	signer := hmac.New(sha256.New, []byte(key))
	signer.Reset()
	signer.Write([]byte(payload))

	hmac := string(signer.Sum(nil))

	b64 := ToBase64Url(hmac)
	return b64
}
