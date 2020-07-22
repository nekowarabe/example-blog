package server

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
)

// 檢查 Header Authorization ，如果有值就將 Token 的部分取出
func extractToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		auth := ctx.Request().Header.Get(echo.HeaderAuthorization)
		if auth == "" {
			return next(ctx)
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 {
			return errors.New("token header format wrong")
		}

		ctx.Set("token", parts[1])

		return next(ctx)
	}
}
