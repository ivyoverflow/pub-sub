package client

import (
	"fmt"

	"github.com/ivyoverflow/internship/pubsub/client/internal/model"
	"golang.org/x/net/websocket"
)

func Run() error {
	ws, err := websocket.Dial("ws://localhost:8080/user/subscribe", "", "https://localhost:8080")
	if err != nil {
		return err
	}

	response := &model.Response{}
	if err := websocket.JSON.Receive(ws, response); err != nil {
		return err
	}

	fmt.Println(response.Message)

	return nil
}
