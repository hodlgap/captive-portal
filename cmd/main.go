package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hodlgap/captive-portal/pkg"
	"github.com/hodlgap/captive-portal/pkg/handler"

	"github.com/pkg/errors"

	"github.com/hodlgap/captive-portal/pkg/config"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

const (
	configFilepath = "./config.yml"
)

func main() {
	c, err := config.Parse(configFilepath)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	nr, err := pkg.NewNewrelic(c)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	app, err := pkg.NewApp(c, nr)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	redisCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password: c.Redis.Password, // no password set
		DB:       c.Redis.DB,       // use default DB
	})

	app = handler.SetRoute(c, app, redisCli)

	// Start server
	go func() {
		if err := app.Start(fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatalf("shutting down the server\n%+v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.App.GracefulTimeoutSec)*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatalf("%+v", err)
	}
}
