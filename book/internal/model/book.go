// Package model contains the description of the tables, responses and requests.
package model

// Book struct represents books table.
type Book struct {
	ID          string  `json:"id" bson:"id" db:"id" validate:"-"`
	Name        string  `json:"name" bson:"name" db:"name" validate:"required"`
	DateOfIssue string  `json:"dateOfIssue" bson:"dateOfIssue" db:"date_of_issue" validate:"required"`
	Author      string  `json:"author" bson:"author" db:"author" validate:"required"`
	Description string  `json:"description" bson:"description" db:"description" validate:"required"`
	Rating      float64 `json:"rating" bson:"rating" db:"rating" validate:"required"`
	Price       float64 `json:"price" bson:"price" db:"price" validate:"required"`
	InStock     bool    `json:"inStock" bson:"inStock" db:"in_stock" validate:"required"`
}
