package handlers

import (
	"errors"

	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// UpdatePostParam 更新文章所需參數
type UpdatePostParam struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

// UpdatePost 更新文章
func UpdatePost(ctx echo.Context) error {
	token, ok := ctx.Get("token").(string)
	if !ok {
		return errors.New("failed to convert token to string")
	}

	var param UpdatePostParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	input := usecases.UpdatePostInput{
		Token:    token,
		Username: ctx.Param("username"),
		ID:       ctx.Param("id"),
		Title:    param.Title,
		Content:  param.Content,
		Tags:     param.Tags,
	}
	if err := usecases.UpdatePost(input); err != nil {
		return err
	}

	return ctx.JSON(0, nil)
}
