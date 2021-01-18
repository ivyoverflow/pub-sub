// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
	handl      *handler.Book
}

// New returns a new configured Server object.
func New(cfg *config.Config, handl *handler.Book) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port),
		},
		handl: handl,
	}
}

// Run configures routes and starts the server.
func (srv *Server) Run() error {
	router := mux.NewRouter()
	booksSubrouter := router.PathPrefix("/books").Subrouter()
	booksSubrouter.HandleFunc("/book/new", srv.handl.Insert).Methods("POST")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Get).Methods("GET")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Update).Methods("PUT")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Delete).Methods("DELETE")

	srv.httpServer.Handler = router

	return srv.httpServer.ListenAndServe()
}
