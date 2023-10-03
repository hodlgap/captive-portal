package handler

import (
	"database/sql"

	echo "github.com/labstack/echo/v4"

	"github.com/hodlgap/captive-portal/pkg/auth"
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
)

func SetRoute(c config.Config, app *echo.Echo, authProvider auth.Provider, db *sql.DB) *echo.Echo {
	app.GET(captiveportal.AuthHandlerURL, captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey, authProvider, db))
	app.POST(captiveportal.AuthGetHandlerURL, captiveportal.NewAuthGetHandler(authProvider))

	return app
}
