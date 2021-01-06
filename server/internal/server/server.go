package server

import (
	"net/http"

	"github.com/ivyoverflow/internship/pubsub/server/internal/handler"
	"github.com/ivyoverflow/internship/pubsub/server/internal/service"
	"golang.org/x/net/websocket"
)

func Run() error {
	pubSub := service.NewPublisherSubscriber()
	publisherHandler := handler.NewPublisherHandler(pubSub)
	subscriberHandler := handler.NewSubscriberHandler(pubSub)

	http.Handle("/publisher/publish", websocket.Handler(publisherHandler.Publish))
	http.Handle("/user/subscribe", websocket.Handler(subscriberHandler.Subscribe))

	return http.ListenAndServe(":8080", nil)
}
