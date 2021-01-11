package main

import (
	"log"

	"github.com/ivyoverflow/pub-sub/publisher/internal/client"
	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
)

func main() {
	logger, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := config.New()
	if cfg.Addr == "" || cfg.Port == "" {
		logger.Fatal("Environment variables ADDR and PORT not found")
		return
	}

	client := client.New(cfg)
	if err = client.Dial(); err != nil {
		logger.Fatal(err.Error())
	}
}
