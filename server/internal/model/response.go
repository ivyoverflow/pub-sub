// Package model contains the described structures that will be used in the project.
package model

// SuccessResponse struct represents the response body from the server.
type SuccessResponse struct {
	Message interface{} `json:"message"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Error Error `json:"error"`
}

// Error ...
type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
