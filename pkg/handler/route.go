package handler

import (
	"database/sql"
	echo "github.com/labstack/echo/v4"
	redis "github.com/redis/go-redis/v9"

	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
)

func SetRoute(c config.Config, app *echo.Echo, rCli *redis.Client, db *sql.DB) *echo.Echo {
	app.GET(captiveportal.AuthHandlerURL, captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey, rCli, db))
	app.POST(captiveportal.AuthGetHandlerURL, captiveportal.NewAuthGetHandler(rCli))

	return app
}
