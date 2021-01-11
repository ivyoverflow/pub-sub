package handler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/model"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

// Subscriber struct contains all handler for subscriber.
type Subscriber struct {
	pubSub *service.PublisherSubscriber
	logger *logger.Logger
}

// NewSubscriber returns a new Subscriber object.
func NewSubscriber(pubSub *service.PublisherSubscriber, logger *logger.Logger) *Subscriber {
	return &Subscriber{pubSub, logger}
}

// Subscribe func processes /user/subscribe route.
func (handler *Subscriber) Subscribe(ws *websocket.Conn) {
	request := &model.SubscribeRequest{}
	if err := websocket.JSON.Receive(ws, request); err != nil {
		handler.logger.Error(err.Error())
		fmt.Fprintf(ws, `{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	messageChannel := handler.pubSub.Subscribe(request.Topic)
	handler.logger.Debug(fmt.Sprintf("The user subscribed to the <<< %s >>> topic", request.Topic))
	for message := range messageChannel {
		response := &model.Response{
			Message: message,
		}

		if err := websocket.JSON.Send(ws, response); err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, `{"error": {"statusCode: %d", "message: %s"}}`, http.StatusInternalServerError, err.Error())

			return
		}

		handler.logger.Debug(fmt.Sprintf("The message <<< %s >>> sends to the subscribed users", response.Message))
	}
}
