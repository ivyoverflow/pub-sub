package main

import (
	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/server"
)

func main() {
	cfg := config.New()
	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	srv := server.New(cfg, log)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
