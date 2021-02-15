package main

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/ivyoverflow/pub-sub/api/internal/handler"
	"github.com/ivyoverflow/pub-sub/api/internal/server"
	"github.com/ivyoverflow/pub-sub/api/internal/service"
	"github.com/ivyoverflow/pub-sub/api/internal/storage/mongo"
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
	gen := service.NewUUIDGenerator()
	bookSvc := service.NewBookController(bookRepo, gen)
	bookHandl := handler.NewBookController(ctx, bookSvc, log)
	srv := server.New(bookHandl)
	if err = srv.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
