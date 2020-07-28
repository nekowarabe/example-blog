package handlers

import (
	"errors"

	"app/src/core/entities"
	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// UpdateAccountParam 更新帳號訊息所需參數
type UpdateAccountParam struct {
	Password string          `json:"password"`
	Tel      string          `json:"tel"`
	Age      uint8           `json:"age"`
	Gender   entities.Gender `json:"gender"`
}

// UpdateAccount 更新帳號訊息
func UpdateAccount(ctx echo.Context) error {
	token, ok := ctx.Get("token").(string)
	if !ok {
		return errors.New("failed to convert token to string")
	}

	var param UpdateAccountParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	input := usecases.UpdateAccountInput{
		Token:    token,
		Username: ctx.Param("username"),
		Password: param.Password,
		Tel:      param.Tel,
		Age:      param.Age,
		Gender:   param.Gender,
	}
	if err := usecases.UpdateAccount(input); err != nil {
		return err
	}

	return ctx.JSON(0, nil)
}
