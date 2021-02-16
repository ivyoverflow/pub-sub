// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
}

// New returns a new configured Server object.
func New(cfg *config.ServerConfig) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		},
	}
}

// Run configures routes and starts the server.
func (server *Server) Run(notificationCtrl *handler.Notification) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/publish", notificationCtrl.Publish)
	mux.Handle("/subscribe", websocket.Handler(notificationCtrl.Subscribe))

	server.httpServer.Handler = mux

	return server.httpServer.ListenAndServe()
}
