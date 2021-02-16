// Package app contains a Run function that is used to configure all application modules and start the server.
package app

import (
	"context"

	"github.com/ivyoverflow/pub-sub/notifier/internal/config"
	"github.com/ivyoverflow/pub-sub/notifier/internal/handler"
	"github.com/ivyoverflow/pub-sub/notifier/internal/redis"
	"github.com/ivyoverflow/pub-sub/notifier/internal/repository"
	"github.com/ivyoverflow/pub-sub/notifier/internal/server"
	"github.com/ivyoverflow/pub-sub/notifier/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// Run configures all application modules and starts the server.
func Run() error {
	ctx := context.Background()
	serverCfg := config.NewServer()
	redisCfg := config.NewRedis()

	db, err := redis.NewDB(ctx, redisCfg)
	if err != nil {
		return err
	}

	log, err := logger.New()
	if err != nil {
		return err
	}

	repo := repository.NewNotification(db)
	svc := service.NewNotification(repo)
	notificationCtrl := handler.NewNotification(svc, log)
	srv := server.New(serverCfg)
	if err := srv.Run(notificationCtrl); err != nil {
		return err
	}

	return nil
}
