// Package server implements server logic: routes initialization and server configuration.
package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivyoverflow/pub-sub/api/internal/handler"
)

// Server represents application server.
type Server struct {
	httpServer *http.Server
	handl      *handler.BookController
}

// New returns a new configured Server object.
func New(handl *handler.BookController) *Server {
	cfg := NewConfig()

	return &Server{
		httpServer: &http.Server{
			Addr: cfg.GetConnectionURI(),
		},
		handl: handl,
	}
}

// Run configures routes and starts the server.
func (srv *Server) Run() error {
	router := mux.NewRouter()
	booksSubrouter := router.PathPrefix("/v1").Subrouter()
	booksSubrouter.HandleFunc("/book/", srv.handl.Insert).Methods("POST")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Get).Methods("GET")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Update).Methods("PUT")
	booksSubrouter.HandleFunc("/book/{id}", srv.handl.Delete).Methods("DELETE")

	srv.httpServer.Handler = router

	return srv.httpServer.ListenAndServe()
}
