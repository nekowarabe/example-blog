package handlers

import (
	"app/src/core/entities"
	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// RegisterParam 註冊帳號會收到的參數
type RegisterParam struct {
	Username string          `json:"username" validate:"required"`
	Password string          `json:"password" validate:"required"`
	Tel      string          `json:"tel"`
	Age      uint8           `json:"age"`
	Gender   entities.Gender `json:"gender"`
}

// Register 註冊帳號
func Register(ctx echo.Context) error {
	var param RegisterParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	input := usecases.RegisterInput{
		Username: param.Username,
		Password: param.Password,
		Tel:      param.Tel,
		Age:      param.Age,
		Gender:   param.Gender,
	}

	if err := usecases.Register(input); err != nil {
		return err
	}

	return ctx.JSON(0, nil)
}
