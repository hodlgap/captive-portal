package main

import (
	"context"
	"fmt"
	"github.com/hodlgap/captive-portal/pkg"
	"github.com/hodlgap/captive-portal/pkg/handler"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pkg/errors"

	"github.com/hodlgap/captive-portal/pkg/config"

	"github.com/labstack/gommon/log"
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

	app = handler.SetRoute(c, app)

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
