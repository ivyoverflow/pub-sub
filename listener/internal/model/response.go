// Package model contains the described structures that will be used in the project.
package model

// Response struct represents the response body from the server.
type Response struct {
	Message interface{} `json:"message"`
}
