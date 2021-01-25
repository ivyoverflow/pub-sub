package service

import "github.com/google/uuid"

// IDGenerator ...
type IDGenerator struct{}

// NewIDGenerator ...
func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

// Generate ...
func (gen *IDGenerator) Generate() uuid.UUID {
	return uuid.New()
}
