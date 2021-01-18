package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/config"
	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/postgres"
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

	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	bookRepo := postgres.NewBookRepository(db)
	bookSvc := service.NewBook(bookRepo)
	bookHandl := handler.NewBook(ctx, bookSvc, log)
	srv := server.New(cfg, bookHandl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
