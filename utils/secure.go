package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecureString(length int) (string, error) {

	// Create a byte slice of the desired length
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the bytes in base64 to make it human-readable
	secureString := base64.URLEncoding.EncodeToString(randomBytes)

	// Return a substring to meet the requested length
	return secureString[:length], nil
}
