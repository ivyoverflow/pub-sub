package server

import (
	"net/http"

	"github.com/ivyoverflow/pub-sub/server/internal/handler"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/server/internal/service"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: ":" + "8080",
		},
	}
}

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
