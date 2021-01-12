package server

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/server/internal/config"
	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
	logger     *logger.Logger
}

// New returns a new configured Server object.
func New(cfg *config.Config, logger *logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		},
		logger: logger,
	}
}

// Run configures routes and starts the server.
func (server *Server) Run() error {
	publisherSubscriber := service.NewPublisherSubscriber()
	publisherHandler := handler.NewPublisher(publisherSubscriber, server.logger)
	subscriberHandler := handler.NewSubscriber(publisherSubscriber, server.logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/publish", publisherHandler.Publish)
	mux.Handle("/subscribe", websocket.Handler(subscriberHandler.Subscribe))

	server.httpServer.Handler = mux

	return server.httpServer.ListenAndServe()
}
