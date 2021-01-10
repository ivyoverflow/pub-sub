package server

import (
	"net/http"

	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
}

// NewServer returns a new configured Server object.
func NewServer() *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: ":" + "8080",
		},
	}
}

// Run configures routes and starts the server.
func (server *Server) Run(logger *logger.Logger) error {
	publisherSubscriber := service.NewPublisherSubscriber()
	publisherHandler := handler.NewPublisherHandler(publisherSubscriber, logger)
	subscriberHandler := handler.NewSubscriberHandler(publisherSubscriber, logger)

	mux := http.NewServeMux()
	mux.Handle("/publish", websocket.Handler(publisherHandler.Publish))
	mux.Handle("/subscribe", websocket.Handler(subscriberHandler.Subscribe))

	server.httpServer.Handler = mux

	return server.httpServer.ListenAndServe()
}
