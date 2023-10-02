package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4/middleware"

	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"

	"github.com/pkg/errors"

	"github.com/hodlgap/captive-portal/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

const (
	configFilepath = "./config.yml"
)

func main() {
	c, err := config.Parse(configFilepath)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}

	// Setup
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)
	app.Use(middleware.HTTPSRedirect())
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/captive-portal", captiveportal.NewHandler("1234567890abcdef1234567890abcdef"))
	app.POST("/captive-portal", captiveportal.NewHandler("1234567890abcdef1234567890abcdef"))

	// Start server
	go func() {
		if err := app.Start(fmt.Sprintf("%s:%d", c.App.Host, c.App.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatal("shutting down the server\n%+v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.App.GracefulTimeoutSec)*time.Second)
	defer cancel()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger.Fatal(err)
	}
}
