// Package handler contains described handlers.
package handler

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/notifier/internal/model"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// Subscriber struct contains all handler for subscriber.
type Subscriber struct {
	svc *service.Notifier
	log *logger.Logger
}

// NewSubscriber returns a new Subscriber object.
func NewSubscriber(svc *service.Notifier, log *logger.Logger) *Subscriber {
	return &Subscriber{svc, log}
}

// Subscribe processes /subscribe route.
func (h *Subscriber) Subscribe(ws *websocket.Conn) {
	for {
		request := model.SubscribeRequest{}
		if err := websocket.JSON.Receive(ws, &request); err != nil {
			h.log.Error(err.Error())
			fmt.Fprintf(ws, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

			return
		}

		channel := h.svc.Subscribe(request.Topic)
		h.log.Debug(fmt.Sprintf("The user subscribed to the <<< %s >>> topic", request.Topic))
		go func(channel chan interface{}) {
			for message := range channel {
				response := model.SuccessResponse{
					Message: message,
				}

				if err := websocket.JSON.Send(ws, &response); err != nil {
					h.log.Error(err.Error())
					fmt.Fprintf(ws, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

					return
				}
			}
		}(channel)
	}
}
