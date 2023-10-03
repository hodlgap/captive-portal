package pkg

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/hodlgap/captive-portal/pkg/auth"
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
)

func NewApp(c config.Config, nr *newrelic.Application, authProvider auth.Provider, db *sql.DB) *echo.Echo {
	app := echo.New()
	app.Logger.SetLevel(c.App.LogLevel.ToLevel())

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(nrecho.Middleware(nr))

	app.GET(captiveportal.AuthHandlerURL, captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey, authProvider, db))
	app.POST(captiveportal.AuthGetHandlerURL, captiveportal.NewAuthGetHandler(authProvider))

	return app
}
