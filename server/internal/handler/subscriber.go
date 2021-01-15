// Package handler contains described handlers.
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
	svc *service.PublisherSubscriber
	log *logger.Logger
}

// NewSubscriber returns a new Subscriber object.
func NewSubscriber(svc *service.PublisherSubscriber, log *logger.Logger) *Subscriber {
	return &Subscriber{svc, log}
}

// Subscribe processes /subscribe route.
func (handler *Subscriber) Subscribe(ws *websocket.Conn) {
	for {
		request := model.SubscribeRequest{}
		if err := websocket.JSON.Receive(ws, &request); err != nil {
			handler.log.Error(err.Error())
			fmt.Fprintf(ws, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

			return
		}

		channel := handler.svc.Subscribe(request.Topic)
		handler.log.Debug(fmt.Sprintf("The user subscribed to the <<< %s >>> topic", request.Topic))
		go func(channel chan interface{}) {
			for message := range channel {
				response := model.SuccessResponse{
					Message: message,
				}

				if err := websocket.JSON.Send(ws, &response); err != nil {
					handler.log.Error(err.Error())
					fmt.Fprintf(ws, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

					return
				}
			}
		}(channel)
	}
}
