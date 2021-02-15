// Package handler contains described handlers.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivyoverflow/pub-sub/notifier/internal/model"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// Publisher struct contains all handlers for publisher.
type Publisher struct {
	svc *service.PublisherSubscriber
	log *logger.Logger
}

// NewPublisher returns a new configured Publisher object.
func NewPublisher(svc *service.PublisherSubscriber, log *logger.Logger) *Publisher {
	return &Publisher{svc, log}
}

// Publish processes /publish route.
func (h *Publisher) Publish(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	request := model.PublishRequest{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	h.svc.Publish(request.Topic, request.Message)
	h.log.Debug(fmt.Sprintf("The publisher sends a new message <<< %s >>> to the <<< %s >>> topic", request.Message, request.Topic))
}
