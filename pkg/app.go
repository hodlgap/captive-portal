package pkg

import (
	"context"
	"fmt"
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewApp(c config.Config, nr *newrelic.Application) (*echo.Echo, error) {
	app := echo.New()
	app.Logger.SetLevel(c.App.LogLevel.ToLevel())

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(nrecho.Middleware(nr))

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/captive-portal", captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey))
	app.POST("/captive-portal", captiveportal.NewAuthGetHandler())

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

	return app, nil
}
