package handlers

import (
	"errors"

	"app/src/core/usecases"

	"github.com/labstack/echo/v4"
)

// AddPostParam 新增文章所需參數
type AddPostParam struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

// AddPost 新增文章
func AddPost(ctx echo.Context) error {
	token, ok := ctx.Get("token").(string)
	if !ok {
		return errors.New("failed to convert token to string")
	}

	var param AddPostParam
	if err := ctx.Bind(&param); err != nil {
		return err
	}

	input := usecases.AddPostInput{
		Token:    token,
		Username: ctx.Param("username"),
		Title:    param.Title,
		Content:  param.Content,
		Tags:     param.Tags,
	}
	if err := usecases.AddPost(input); err != nil {
		return err
	}

	return ctx.JSON(0, nil)
}
