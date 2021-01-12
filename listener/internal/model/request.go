package model

// Request struct represents the publish request body to the server.
type Request struct {
	Topic string `json:"topic"`
}
