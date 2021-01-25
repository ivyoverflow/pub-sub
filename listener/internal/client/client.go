package client

import (
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
	"github.com/ivyoverflow/pub-sub/publisher/internal/model"
)

type Client struct {
	log *logger.Logger
	cfg *config.Config
}

func New(log *logger.Logger, cfg *config.Config) *Client {
	return &Client{
		log: log,
		cfg: cfg,
	}
}

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
