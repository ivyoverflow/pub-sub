// Package model contains the described structures that will be used in the project.
package model

// Request struct represents the publish request body to the server.
type Request struct {
	Topic string `json:"topic"`
}
