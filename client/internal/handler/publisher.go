package handler

import (
	"fmt"
	"net/http"

	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
	"github.com/ivyoverflow/pub-sub/publisher/internal/model"
	"golang.org/x/net/websocket"
)

// PublisherHandler struct contains all handlers for publisher.
type PublisherHandler struct {
	logger *logger.Logger
}

// NewPublisherHandler returns a new configured PublisherHandler object.
func NewPublisherHandler(logger *logger.Logger) *PublisherHandler {
	return &PublisherHandler{logger}
}

// Publish connects to the server and sends a request.
func (handler *PublisherHandler) Publish(ws *websocket.Conn) {
	for {
		request := &model.PublishRequest{}
		if err := websocket.JSON.Receive(ws, request); err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusBadRequest, err.Error()))

			return
		}

		// connect to our server.
		serverWs, err := websocket.Dial("ws://localhost:8080/publish", "", "http://localhost:8080")
		if err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusInternalServerError, err.Error()))

			return
		}

		// sending data to the server.
		if err = websocket.JSON.Send(serverWs, request); err != nil {
			handler.logger.Error(err.Error())
			fmt.Fprintf(ws, fmt.Sprintf(`{"error": {"statusCode: %d", "message: %s"}}`, http.StatusInternalServerError, err.Error()))

			return
		}

		handler.logger.Debug("The publisher sends a new message request to the server")
	}
}
