package validator

import (
	"github.com/go-playground/validator"

	"github.com/ivyoverflow/pub-sub/book/internal/model"
)

// Validator contains all Validate struct methods inside the validator field.
type Validator struct {
	validator *validator.Validate
}

// New returns a new configured Validator object.
func New() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate checks if received struct is valid.
func (vld *Validator) Validate(book *model.Book) error {
	if err := vld.validator.Struct(book); err != nil {
		return err
	}

	return nil
}
