package main

import (
	"flag"

	"github.com/ivyoverflow/pub-sub/publisher/internal/client"
	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
)

func main() {
	var topic string
	flag.StringVar(&topic, "t", "", "sets the topic name for the subscription")
	flag.Parse()

	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	if topic == "" {
		log.Fatal("Set the name of the topic to subscribe. For example: \"news\"")
	}

	cfg := config.New()
	if cfg.Addr == "" || cfg.Port == "" {
		log.Fatal("Environment variables ADDR and PORT not found")
		return
	}

	clt := client.New(log, cfg)
	if err := clt.Run(topic); err != nil {
		log.Fatal(err.Error())
	}
}
