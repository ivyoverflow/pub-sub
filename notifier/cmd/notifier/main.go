package main

import (
	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/server"
	"github.com/ivyoverflow/pub-sub/platform/logger"
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

	svr := server.New(cfg, log)
	if err := svr.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
