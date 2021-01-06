package model

type PublishRequest struct {
	Topic   string      `json:"topic"`
	Message interface{} `json:"message"`
}

type SubscribeRequest struct {
	Topic string `json:"topic"`
}
