package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/internship/pubsub/server/internal/service"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/internship/pubsub/server/internal/model"
)

type PublisherHandler struct {
	pubSub *service.PublisherSubscriber
}

func NewPublisherHandler(pubSub *service.PublisherSubscriber) *PublisherHandler {
	return &PublisherHandler{pubSub}
}

func (handler *PublisherHandler) Publish(ws *websocket.Conn) {
	var request *model.PublishRequest
	if err := websocket.JSON.Receive(ws, &request); err != nil {
		fmt.Printf("ERROR: %s", err.Error())
		fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))
		return
	}

	handler.pubSub.Publish(request.Topic, request.Message)
}
