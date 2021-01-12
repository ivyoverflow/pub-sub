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
	logger *logger.Logger
}

// NewPublisher returns a new configured Publisher object.
func NewPublisher(pubSub *service.PublisherSubscriber, logger *logger.Logger) *Publisher {
	return &Publisher{pubSub, logger}
}

// Publish processes /publish route.
func (handler *Publisher) Publish(rw http.ResponseWriter, r *http.Request) {
	request := &model.PublishRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		handler.logger.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	handler.pubSub.Publish(request.Topic, request.Message)
	handler.logger.Debug(fmt.Sprintf("The publisher sends a new message <<< %s >>> to the <<< %s >>> topic", request.Message, request.Topic))
}
