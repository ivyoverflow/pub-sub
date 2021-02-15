// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
	log        *logger.Logger
}

// New returns a new configured Server object.
func New(cfg *config.Config, log *logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		},
		log: log,
	}
}

// Run configures routes and starts the server.
func (server *Server) Run() error {
	svc := service.NewPublisherSubscriber()
	publisherHandler := handler.NewPublisher(svc, server.log)
	subscriberHandler := handler.NewSubscriber(svc, server.log)

	mux := http.NewServeMux()
	mux.HandleFunc("/publish", publisherHandler.Publish)
	mux.Handle("/subscribe", websocket.Handler(subscriberHandler.Subscribe))

	server.httpServer.Handler = mux

	return server.httpServer.ListenAndServe()
}
