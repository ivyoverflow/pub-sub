package main

import (
	"log"

	"github.com/ivyoverflow/pub-sub/publisher/internal/client"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err.Error())
	}

	client := client.NewClient()
	if err := client.Run(logger); err != nil {
		logger.Fatal(err.Error())
	}
}
