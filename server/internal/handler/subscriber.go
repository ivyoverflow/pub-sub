package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/internship/pubsub/server/internal/service"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/internship/pubsub/server/internal/model"
)

// SubscriberHandler struct contains all handler for subscriber.
type SubscriberHandler struct {
	pubSub *service.PublisherSubscriber
}

// NewSubscriberHandler returns a new SubscriberHandler object.
func NewSubscriberHandler(pubSub *service.PublisherSubscriber) *SubscriberHandler {
	return &SubscriberHandler{pubSub}
}

// Subscribe func processes /user/subscribe route.
func (handler *SubscriberHandler) Subscribe(ws *websocket.Conn) {
	var request *model.SubscribeRequest
	if err := websocket.JSON.Receive(ws, &request); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))
		return
	}

	messageChannel := handler.pubSub.Subscribe(request.Topic)
	response := &model.Response{
		Message: <-messageChannel,
	}

	if err := websocket.JSON.Send(ws, response); err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusInternalServerError, err.Error()))
		return
	}
}
