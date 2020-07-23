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
	repositories.Post = memory.NewPostRepository()

	server := echo.New()
	server.HideBanner = true
	server.Debug = true

	server.Use(
		extractToken,
	)

	server.POST("/login", handlers.Login)
	server.POST("/accounts", handlers.Register)
	server.GET("/accounts/:username", handlers.GetAccountInfo)
	server.DELETE("/accounts/:username", handlers.DeleteAccount)

	server.POST("/posts/:username", handlers.AddPost)
	server.POST("/posts/:username/:id", handlers.UpdatePost)

	return server.Start(":" + configs.HTTP().Port)
}
