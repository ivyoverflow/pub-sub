package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/internship/pubsub/server/internal/service"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/internship/pubsub/server/internal/model"
)

// PublisherHandler struct contains all handlers for publisher.
type PublisherHandler struct {
	pubSub *service.PublisherSubscriber
}

// NewPublisherHandler returns a new configured PublisherHandler object.
func NewPublisherHandler(pubSub *service.PublisherSubscriber) *PublisherHandler {
	return &PublisherHandler{pubSub}
}

// Publish func processes publisher/publish route.
func (handler *PublisherHandler) Publish(ws *websocket.Conn) {
	request := &model.PublishRequest{}
	if err := websocket.JSON.Receive(ws, request); err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		fmt.Println("!!!")
		fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))
		return
	}

	fmt.Println(request.Topic, request.Message)

	handler.pubSub.Publish(request.Topic, request.Message)
}
