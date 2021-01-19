package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/lib/validator"
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

	db, err := mongo.New(ctx, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	bookRepo := mongo.NewBookRepository(db)
	vld := validator.New()
	bookSvc := service.NewBook(bookRepo, vld)
	bookMw := handler.NewMiddleware(ctx, log)
	bookHandl := handler.NewBook(ctx, bookSvc, log)
	srv := server.New(cfg, bookMw, bookHandl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
