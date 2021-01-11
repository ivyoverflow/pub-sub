package client

import (
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/model"
)

// Client contains addr and port fields that will be used to connect to the server.
type Client struct {
	addr string
	port string
}

// New returns a new configured Client object.
func New(cfg *config.Config) *Client {
	return &Client{
		addr: cfg.Addr,
		port: cfg.Port,
	}
}

// Dial connects to the server and sends a filled request.
func (client *Client) Dial() error {
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/publish", client.addr, client.port), "", fmt.Sprintf("http://%s:%s", client.addr, client.port))
	if err != nil {
		return err
	}

	defer ws.Close()

	request := &model.PublishRequest{
		Topic:   "news",
		Message: "As the world begins its vaccination push, delayed rollouts draw criticism and concern",
	}

	if err = websocket.JSON.Send(ws, request); err != nil {
		return err
	}

	return nil
}
