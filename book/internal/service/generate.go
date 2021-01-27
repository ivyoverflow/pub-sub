package service

import "github.com/google/uuid"

// IDGenerator contains Generate() method.
type IDGenerator struct{}

// NewIDGenerator returns a new configured IDGenerator object.
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// Generate generates a new UUID ID.
func (gen *IDGenerator) Generate() uuid.UUID {
	return uuid.New()
}
