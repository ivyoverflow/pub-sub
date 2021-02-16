// Package handler contains all implemented application handlers.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/notifier/internal/model"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// Notification implements all notifier controllers/handlers.
type Notification struct {
	svc service.Notifier
	log *logger.Logger
}

// NewNotification returns a new configured Notification object.
func NewNotification(svc service.Notifier, log *logger.Logger) *Notification {
	return &Notification{svc, log}
}

// Publish processes /publish route.
func (h *Notification) Publish(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	request := model.PublishRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.svc.Publish(r.Context(), request.Book, request.Message); err != nil {
		h.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	h.log.Debug(fmt.Sprintf("The publisher sends a new message <<< %s >>> to the <<< %s >>> topic", request.Message, request.Book))
}

// Subscribe processes /subscribe route.
func (h *Notification) Subscribe(ws *websocket.Conn) {
	for {
		request := model.SubscribeRequest{}
		if err := websocket.JSON.Receive(ws, &request); err != nil {
			h.log.Error(err.Error())
			fmt.Fprintf(ws, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

			return
		}

		channel := h.svc.Subscribe(ws.Request().Context(), request.Book)
		h.log.Debug(fmt.Sprintf("The user subscribed to the <<< %s >>> topic", request.Book))
		go func(channel <-chan *redis.Message) {
			for message := range channel {
				response := model.SuccessResponse{
					Message: message.Payload,
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
