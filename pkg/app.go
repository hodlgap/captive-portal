package pkg

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/hodlgap/captive-portal/pkg/config"
)

func NewApp(c config.Config, nr *newrelic.Application) (*echo.Echo, error) {
	app := echo.New()
	app.Logger.SetLevel(c.App.LogLevel.ToLevel())

	//app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	//app.Use(nrecho.Middleware(nr))

	return app, nil
}
