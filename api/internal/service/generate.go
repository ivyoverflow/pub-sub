package service

import "github.com/google/uuid"

// UUIDGenerator contains GenerateUUID() method.
type UUIDGenerator struct{}

// NewUUIDGenerator returns a new configured UUIDGenerator object.
func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

// GenerateUUID Generates a new UUID ID.
func (gen *UUIDGenerator) GenerateUUID() uuid.UUID {
	return uuid.New()
}
