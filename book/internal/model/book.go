// Package model contains the description of the tables, responses and requests.
package model

// Book struct represents books table.
type Book struct {
	ID          string  `json:"id" bson:"id" db:"id"`
	Name        string  `json:"name" bson:"name" db:"name"`
	DateOfIssue string  `json:"dateOfIssue" bson:"dateOfIssue" db:"date_of_issue"`
	Author      string  `json:"author" bson:"author" db:"author"`
	Description string  `json:"description" bson:"description" db:"description"`
	Rating      float64 `json:"rating" bson:"rating" db:"rating"`
	Price       float64 `json:"price" bson:"price" db:"price"`
	InStock     bool    `json:"inStock" bson:"inStock" db:"in_stock"`
}
