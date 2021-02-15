package service

import (
	"github.com/go-playground/validator"

	"github.com/ivyoverflow/pub-sub/api/internal/model"
)

// Validate checks if received struct is valid.
func Validate(book *model.Book) error {
	vld := validator.New()
	if err := vld.Struct(book); err != nil {
		return err
	}

	return nil
}
