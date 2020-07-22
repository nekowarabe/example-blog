package handlers

import (
	"errors"

	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// DeleteAccount 刪除帳號
func DeleteAccount(ctx echo.Context) error {
	token, ok := ctx.Get("token").(string)
	if !ok {
		return errors.New("failed to convert token to string")
	}

	input := usecases.DeleteAccountInput{
		Token:    token,
		Username: ctx.Param("username"),
	}

	if err := usecases.DeleteAcoount(input); err != nil {
		return err
	}

	return ctx.JSON(0, nil)
}
