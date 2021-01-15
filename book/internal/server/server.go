// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	"github.com/ivyoverflow/pub-sub/book/internal/store"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
	log        *logger.Logger
	cfg        *config.Config
}

// New returns a new configured Server object.
func New(cfg *config.Config, log *logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port),
		},
		log: log,
		cfg: cfg,
	}
}

// Run configures routes and starts the server.
func (server *Server) Run() error {
	str, err := store.New(server.cfg)
	if err != nil {
		return err
	}

	svc, err := service.NewManager(str)
	if err != nil {
		return err
	}

	bookHandler := handler.NewBook(svc, server.log)
	router := mux.NewRouter()
	bookRouter := router.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/book/new", bookHandler.Add).Methods("POST")
	bookRouter.HandleFunc("/book/{id}", bookHandler.Get).Methods("GET")
	bookRouter.HandleFunc("/book/{id}", bookHandler.Update).Methods("PUT")
	bookRouter.HandleFunc("/book/{id}", bookHandler.Delete).Methods("DELETE")

	server.httpServer.Handler = router

	return server.httpServer.ListenAndServe()
}
