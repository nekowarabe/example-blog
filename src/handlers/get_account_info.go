package handlers

import (
	"errors"

	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// GetAccountInfo 取得帳號資訊
func GetAccountInfo(ctx echo.Context) error {
	token, ok := ctx.Get("token").(string)
	if !ok {
		return errors.New("failed to convert token to string")
	}

	input := usecases.GetAccountInfoInput{
		Token:    token,
		Username: ctx.Param("username"),
	}
	output, err := usecases.GetAccountInfo(input)
	if err != nil {
		return err
	}

	return ctx.JSON(0, output)
}
