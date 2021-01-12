// Package handler contains described handlers.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/model"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

// Publisher struct contains all handlers for publisher.
type Publisher struct {
	pubSub *service.PublisherSubscriber
	log    *logger.Logger
}

// NewPublisher returns a new configured Publisher object.
func NewPublisher(pubSub *service.PublisherSubscriber, log *logger.Logger) *Publisher {
	return &Publisher{pubSub, log}
}

// Publish processes /publish route.
func (handler *Publisher) Publish(rw http.ResponseWriter, r *http.Request) {
	request := &model.PublishRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	handler.pubSub.Publish(request.Topic, request.Message)
	handler.log.Debug(fmt.Sprintf("The publisher sends a new message <<< %s >>> to the <<< %s >>> topic", request.Message, request.Topic))
}
