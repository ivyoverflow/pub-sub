package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/model"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
	"golang.org/x/net/websocket"
)

// SubscriberHandler struct contains all handler for subscriber.
type SubscriberHandler struct {
	pubSub *service.PublisherSubscriber
	logger *logger.Logger
}

// NewSubscriberHandler returns a new SubscriberHandler object.
func NewSubscriberHandler(pubSub *service.PublisherSubscriber, logger *logger.Logger) *SubscriberHandler {
	return &SubscriberHandler{pubSub, logger}
}

// Subscribe func processes /user/subscribe route.
func (handler *SubscriberHandler) Subscribe(ws *websocket.Conn) {
	request := &model.SubscribeRequest{}
	if err := websocket.JSON.Receive(ws, request); err != nil {
		handler.logger.Error(err.Error())
		fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))

		return
	}

	messageChannel := handler.pubSub.Subscribe(request.Topic)
	for message := range messageChannel {
		response := &model.Response{
			Message: message,
		}

		handler.logger.Debug(fmt.Sprintf("The user subscribed to the <<< %s >>> topic", request.Topic))

		if err := websocket.JSON.Send(ws, response); err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusInternalServerError, err.Error()))

			return
		}

		handler.logger.Debug(fmt.Sprintf("The message <<< %s >>> sends to the subscribed users", response.Message))
	}
}
