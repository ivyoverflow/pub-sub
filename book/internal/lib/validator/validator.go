package validator

import (
	"github.com/go-playground/validator"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// Validator ...
type Validator struct {
	validator *validator.Validate
}

// New ...
func New() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate ...
func (vld *Validator) Validate(book *model.Book) error {
	if err := vld.validator.Struct(book); err != nil {
		return err
	}

	return nil
}
