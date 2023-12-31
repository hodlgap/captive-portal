package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/newrelic/go-agent/v3/integrations/nrredis-v9"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"

	"github.com/hodlgap/captive-portal/pkg"
	"github.com/hodlgap/captive-portal/pkg/auth"
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/models"
)

func main() {
	c, err := config.FromEnv()
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	nr, err := pkg.NewNewrelic(c)
	if err != nil {
		log.Fatalf("%+v", errors.WithStack(err))
	}

	redisOpt := &redis.Options{
		Addr:                  fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
		Password:              c.Redis.Password, // no password set
		DB:                    c.Redis.DB,       // use default Database
		MaxRetries:            5,
		MinRetryBackoff:       3,
		MaxRetryBackoff:       7,
		DialTimeout:           5 * time.Second,
		ReadTimeout:           3 * time.Second,
		WriteTimeout:          3 * time.Second,
		ContextTimeoutEnabled: true,
		MinIdleConns:          3,
		MaxIdleConns:          5,
		ConnMaxIdleTime:       30 * time.Minute,
		ConnMaxLifetime:       0,
	}
	redisCli := redis.NewClient(redisOpt)
	redisCli.AddHook(nrredis.NewHook(redisOpt))

	db, err := models.NewDB(c.Database.URL)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	authProvider := auth.NewProvider(redisCli)

	app := pkg.NewApp(c, nr, authProvider, db)

	// Start server
	go func() {
		if err := app.Start(fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatalf("shutting down the server\n%+v", errors.WithStack(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.App.GracefulTimeoutSec)*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatalf("%+v", errors.WithStack(err))
	}
}
