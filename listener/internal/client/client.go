// Package client contains all method to run and configure the application client.
package client

import (
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/listenter/internal/config"
	"github.com/ivyoverflow/pub-sub/listenter/internal/logger"
	"github.com/ivyoverflow/pub-sub/listenter/internal/model"
)

// Client represents application client.
type Client struct {
	log *logger.Logger
	cfg *config.Config
}

// New returns a new configured Client object.
func New(log *logger.Logger, cfg *config.Config) *Client {
	return &Client{
		log: log,
		cfg: cfg,
	}
}

// Run runs application client.
func (client *Client) Run(topic string) error {
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/subscribe", client.cfg.Addr, client.cfg.Port), "",
		fmt.Sprintf("http://%s:%s", client.cfg.Addr, client.cfg.Port))
	if err != nil {
		return err
	}

	defer ws.Close()

	request := &model.Request{
		Topic: topic,
	}

	if err := websocket.JSON.Send(ws, request); err != nil {
		return err
	}

	for {
		response := &model.Response{}
		if err := websocket.JSON.Receive(ws, response); err != nil {
			return err
		}

		client.log.Info(fmt.Sprintf("Client received <<< %s >> message", response.Message))
	}
}
