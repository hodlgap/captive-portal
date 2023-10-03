package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"

	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
)

func SetRoute(c config.Config, app *echo.Echo, rCli *redis.Client) *echo.Echo {
	app.GET(captiveportal.AuthHandlerURL, captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey, rCli))
	app.POST(captiveportal.AuthGetHandlerURL, captiveportal.NewAuthGetHandler(rCli))

	return app
}
