package handler

import (
	"github.com/hodlgap/captive-portal/pkg/config"
	"github.com/hodlgap/captive-portal/pkg/handler/captiveportal"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetRoute(c config.Config, app *echo.Echo) *echo.Echo {
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	app.GET("/captive-portal", captiveportal.NewAuthHandler(c.Openwrt.EncryptionKey))
	app.POST("/captive-portal", captiveportal.NewAuthGetHandler())

	return app
}
