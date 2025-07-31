package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken(n int) (string, error) {
	// Create a byte slice of length n
	b := make([]byte, n)

	// Read random bytes into slice
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Convert to hex string
	return hex.EncodeToString(b), nil
}
