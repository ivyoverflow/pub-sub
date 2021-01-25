// Package model contains the description of the tables, responses and requests.
package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Book struct represents books table.
type Book struct {
	ID          uuid.UUID       `json:"id" bson:"id" db:"id" validate:"-"`
	Name        string          `json:"name" bson:"name" db:"name" validate:"required"`
	DateOfIssue string          `json:"dateOfIssue" bson:"dateOfIssue" db:"date_of_issue" validate:"required"`
	Author      string          `json:"author" bson:"author" db:"author" validate:"required"`
	Description string          `json:"description" bson:"description" db:"description" validate:"required"`
	Rating      decimal.Decimal `json:"rating" bson:"rating" db:"rating" validate:"required"`
	Price       decimal.Decimal `json:"price" bson:"price" db:"price" validate:"required"`
	InStock     bool            `json:"inStock" bson:"inStock" db:"in_stock" validate:"required"`
}
