package main

import (
	"log"

	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/server"
)

func main() {
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.NewServer()
	if err := server.Run(logger); err != nil {
		log.Fatal(err.Error())
	}
}
