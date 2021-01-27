// Package model contains the description of the tables, responses and requests.
package model

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Book struct represents books table.
type Book struct {
	ID          uuid.UUID `json:"id" bson:"id" db:"id" validate:"-"`
	Name        string    `json:"name" bson:"name" db:"name" validate:"required"`
	DateOfIssue string    `json:"dateOfIssue" bson:"dateOfIssue" db:"date_of_issue" validate:"required"`
	Author      string    `json:"author" bson:"author" db:"author" validate:"required"`
	Description string    `json:"description" bson:"description" db:"description" validate:"required"`
	Rating      Decimal   `json:"rating" bson:"rating" db:"rating" validate:"required"`
	Price       Decimal   `json:"price" bson:"price" db:"price" validate:"required"`
	InStock     bool      `json:"inStock" bson:"inStock" db:"in_stock" validate:"required"`
}

// Decimal inherits all decimal.Decimal methods and contains Marshaler and Unmarshaler implementations.
type Decimal struct {
	decimal.Decimal
}

// UnmarshalBSONValue is a custom Unmarshaler interface method implementation.
func (d *Decimal) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	val := bson.M{}
	err := bson.Unmarshal(data, &val)
	if err != nil {
		return err
	}
	if _, ok := val["decimal"]; !ok {
		return errors.New("failed to unmarshall BSON, no decimal field")
	}
	return d.Decimal.Scan(val["decimal"])
}

// MarshalBSON is a custom Marshal interface method implementation.
func (d Decimal) MarshalBSON() (data []byte, err error) {
	bts, err := d.Decimal.MarshalText()
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal decimal value")
	}
	return bson.Marshal(bson.M{"decimal": string(bts)})
}
