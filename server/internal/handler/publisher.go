package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/model"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
	"golang.org/x/net/websocket"
)

// PublisherHandler struct contains all handlers for publisher.
type PublisherHandler struct {
	pubSub *service.PublisherSubscriber
	logger *logger.Logger
}

// NewPublisherHandler returns a new configured PublisherHandler object.
func NewPublisherHandler(pubSub *service.PublisherSubscriber, logger *logger.Logger) *PublisherHandler {
	return &PublisherHandler{pubSub, logger}
}

// Publish func processes publisher/publish route.
func (handler *PublisherHandler) Publish(ws *websocket.Conn) {
	for {
		request := &model.PublishRequest{}
		if err := websocket.JSON.Receive(ws, request); err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))

			return
		}

		handler.pubSub.Publish(request.Topic, request.Message)
		handler.logger.Debug(fmt.Sprintf("The publisher sends a new message <<< %s >>> to the <<< %s >>> topic", request.Message, request.Topic))
	}
}
