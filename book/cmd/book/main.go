package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
	"github.com/ivyoverflow/pub-sub/book/internal/server"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
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

	bookRepo := mongo.NewBookRepository(db)
	gen := service.NewIDGenerator()
	bookSvc := service.NewBook(bookRepo, gen)
	bookMw := handler.NewMiddleware(ctx, log)
	bookHandl := handler.NewBook(ctx, bookSvc, log)
	srv := server.New(&cfg.Server, bookMw, bookHandl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
