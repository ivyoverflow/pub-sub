package main

import (
	"log"

	"github.com/ivyoverflow/internship/pubsub/server/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
