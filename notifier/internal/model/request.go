// Package model contains the described structures that will be used in the project.
package model

// PublishRequest struct represents the publish request body to the server.
type PublishRequest struct {
	Book    string      `json:"book"`
	Message interface{} `json:"message"`
}

// SubscribeRequest struct represents the subscribe request body to the server.
type SubscribeRequest struct {
	Book string `json:"book"`
}
