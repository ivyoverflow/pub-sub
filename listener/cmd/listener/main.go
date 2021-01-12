package main

import (
	"fmt"

	"golang.org/x/net/websocket"

	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
	"github.com/ivyoverflow/pub-sub/publisher/internal/model"
)

func main() {
	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := config.New()
	if cfg.Addr == "" || cfg.Port == "" {
		log.Fatal("Environment variables ADDR and PORT not found")
		return
	}

	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%s/subscribe", cfg.Addr, cfg.Port), "", fmt.Sprintf("http://%s:%s", cfg.Addr, cfg.Port))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer ws.Close()

	request := &model.Request{
		Topic: "news",
	}

	if err := websocket.JSON.Send(ws, request); err != nil {
		log.Error(err.Error())
	}

	for {
		response := &model.Response{}
		if err := websocket.JSON.Receive(ws, response); err != nil {
			log.Error(err.Error())
		}

		log.Info(fmt.Sprintf("Client received <<< %s >> message", response.Message))
	}
}
