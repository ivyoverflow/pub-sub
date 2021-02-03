// Package types contains all custom errors and types.
package types

import "errors"

// Defines all custom errors that we will use.
var (
	// Returned if the item was not found in the database table.
	// For example: we have an element id, which is a string, when we pass
	// this id to the repository's Get () method, and the element with the
	// passed id is not found, we will get an ErrorNotFound error.
	ErrorNotFound = errors.New("not found")
	// Returned if the received JSON body is invalid.
	// For example: we have a structure that has the following fields:
	// username, email and password.
	// If our user sends a JSON body without any of these fields,
	// we will receive an ErrorBadRequest error.
	ErrorBadRequest = errors.New("bad request")
	// Returned if internal server logic throws an unknown error.
	ErrorInternalServerError = errors.New("internal server error")
	// Returned if the received JSON body has a duplicate value.
	// For example: we have a structure that has the following fields:
	// username, email and password.
	// If our user send a JSON body with username that has already been
	// inserted into the database table, we will get an ErrorDuplicateValue error.
	ErrorDuplicateValue            = errors.New("duplicate value")
	ErrorMongoConnectionRefused    = errors.New("mongodb connection refused")
	ErrorPostgresConnectionRefused = errors.New("postgres connection refused")
	ErrorMigrate                   = errors.New("migrations cannot start")
	ErrorConfigInitialization      = errors.New("config initialization failed")
)
