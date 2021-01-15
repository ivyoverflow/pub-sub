// Package service contains all service logic.
package service

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateUniqueID generates a random string that will be used for the book ID.
func GenerateUniqueID() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
