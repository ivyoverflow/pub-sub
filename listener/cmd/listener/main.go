package main

import (
	"flag"
	"log"

	"github.com/ivyoverflow/pub-sub/publisher/internal/client"
	"github.com/ivyoverflow/pub-sub/publisher/internal/config"
	"github.com/ivyoverflow/pub-sub/publisher/internal/logger"
)

var (
	topic string
)

func init() {
	flag.StringVar(&topic, "topic", "", "sets the topic name for the subscription")
}

func main() {
	flag.Parse()
	if topic == "" {
		log.Fatal("Set the name of the topic to subscribe. For example: \"news\"")
	}

	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
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
