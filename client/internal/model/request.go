package model

// PublishRequest struct represents the publish request body to the server.
type PublishRequest struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}

// SubscribeRequest struct represents the subscribe request body to the server.
type SubscribeRequest struct {
	Topic string `json:"topic"`
}
