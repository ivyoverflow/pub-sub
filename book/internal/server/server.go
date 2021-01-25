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
	mw         *handler.Middleware
	handl      *handler.Book
}

// New returns a new configured Server object.
func New(cfg *config.Config, mw *handler.Middleware, handl *handler.Book) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf("%s:%s", cfg.Server.Addr, cfg.Server.Port),
		},
		handl: handl,
		mw:    mw,
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
	booksSubrouter.Use(srv.mw.AbortWithContext) // middleware component.

	srv.httpServer.Handler = router

	return srv.httpServer.ListenAndServe()
}
