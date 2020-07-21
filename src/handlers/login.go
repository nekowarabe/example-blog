package handlers

import (
	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// LoginParam 登入帳號所需參數
type LoginParam struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login 登入帳號
func Login(ctx echo.Context) error {
	var param LoginParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	input := usecases.LoginInput{
		Username: param.Username,
		Password: param.Password,
	}
	output, err := usecases.Login(input)
	if err != nil {
		return err
	}

	return ctx.JSON(0, output)
}
