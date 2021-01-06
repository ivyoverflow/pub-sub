package main

import (
	"log"

	"github.com/ivyoverflow/internship/pubsub/client/internal/client"
)

func main() {
	if err := client.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
