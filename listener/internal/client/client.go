package client

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
	"github.com/ivyoverflow/pub-sub/publisher/internal/model"
)

// Client contains addr and port fields that will be used to connect to the server.
type Client struct {
	addr   string
	port   string
	logger *logger.Logger
}

// New returns a new configured Client object.
func New(cfg *config.Config, logger *logger.Logger) *Client {
	return &Client{
		addr:   cfg.Addr,
		port:   cfg.Port,
		logger: logger,
	}
}

// Dial connects to the server and sends a filled request.
func (client *Client) Dial() error {
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/subscribe", client.addr, client.port), "", fmt.Sprintf("http://%s:%s", client.addr, client.port))
	if err != nil {
		return err
	}

	defer ws.Close()

	request := &model.Request{
		Topic: "news",
	}

	if err := websocket.JSON.Send(ws, request); err != nil {
		log.Println(err.Error())

		return err
	}
	response := &model.Response{}
	if err := websocket.JSON.Receive(ws, response); err != nil {
		log.Println(err.Error())

		return err
	}

	client.logger.Info(fmt.Sprintf("Client received <<< %s >> message", response.Message))

	return nil
}
