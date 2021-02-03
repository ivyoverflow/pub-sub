package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
	"github.com/ivyoverflow/pub-sub/book/internal/server"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func main() {
	ctx := context.Background()
	cfg := config.New()
	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := mongo.New(ctx, &cfg.Mongo)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := mongo.NewBookRepository(db)
	gen := service.NewUUIDGenerator()
	svc := service.NewBookController(repo, gen)
	ctrl := handler.NewBookController(ctx, svc, log)
	srv := server.New(&cfg.Server, ctrl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
