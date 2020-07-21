package server

import (
	"app/src/configs"
	"app/src/core/repositories"
	"app/src/handlers"
	"app/src/store/memory"

	"github.com/labstack/echo/v4"
)

// Run 執行 Server
func Run() error {
	repositories.Account = memory.NewAccountRepository()

	server := echo.New()
	server.HideBanner = true
	server.Debug = true

	server.POST("/accounts", handlers.Register)

	return server.Start(":" + configs.HTTP().Port)
}
