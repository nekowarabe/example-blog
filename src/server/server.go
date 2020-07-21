package server

import (
	"app/src/configs"
	"app/src/handlers"

	"github.com/labstack/echo/v4"
)

// Run 執行 Server
func Run() error {
	server := echo.New()
	server.HideBanner = true

	server.POST("/accounts", handlers.Register)

	return server.Start(":" + configs.HTTP().Port)
}
