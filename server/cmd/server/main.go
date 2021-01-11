package main

import (
	"log"

	"github.com/ivyoverflow/pub-sub/server/internal/config"
	"github.com/ivyoverflow/pub-sub/server/internal/logger"
	"github.com/ivyoverflow/pub-sub/server/internal/server"
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

	server := server.New(cfg, logger)
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
