package client

import (
	"net/http"

	"github.com/ivyoverflow/pub-sub/publisher/internal/handler"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
	"golang.org/x/net/websocket"
)

// Client represents application client.
type Client struct {
	httpServer *http.Server
}

// NewClient returns a new configured Client object.
func NewClient() *Client {
	return &Client{
		httpServer: &http.Server{
			Addr: ":" + "3030",
		},
	}
}

// Run configures routes and starts the server.
func (client *Client) Run(logger *logger.Logger) error {
	publisherHandler := handler.NewPublisherHandler(logger)

	mux := http.NewServeMux()
	mux.Handle("/publish", websocket.Handler(publisherHandler.Publish))

	client.httpServer.Handler = mux

	return client.httpServer.ListenAndServe()
}
