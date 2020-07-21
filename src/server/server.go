package server

import (
	"app/src/config"

	"github.com/labstack/echo/v4"
)

// Run 執行 Server
func Run() error {
	server := echo.New()
	server.HideBanner = true

	return server.Start(":" + config.HTTP().Port)
}
