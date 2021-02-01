package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/book/internal/handler"
	"github.com/ivyoverflow/pub-sub/book/internal/repository/mongo"
	"github.com/ivyoverflow/pub-sub/book/internal/server"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

func main() {
	ctx := context.Background()
	log, err := logger.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := mongo.New(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	bookRepo := mongo.NewBookRepository(db)
	gen := service.NewIDGenerator()
	bookSvc := service.NewBook(bookRepo, gen)
	bookHandl := handler.NewBook(ctx, bookSvc, log)
	srv := server.New(bookHandl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
