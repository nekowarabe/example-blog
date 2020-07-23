package usecases

import (
	"errors"
	"time"

	"app/src/core/entities"
	"app/src/core/repositories"

	uuid "github.com/satori/go.uuid"
)

// AddPostInput 新增文章所需的參數
type AddPostInput struct {
	Token    string
	Username string
	Title    string
	Content  string
	Tags     []string
}

// AddPost 新增文章
func AddPost(input AddPostInput) error {
	account, err := repositories.Account.GetByToken(input.Token)
	if err != nil {
		return err
	}
	if account.Username != input.Username {
		return errors.New("unauthorization to add post")
	}

	now := time.Now()
	post := entities.Post{
		ID:        uuid.NewV4().String(),
		CreatedAt: now,
		UpdatedAt: now,
		Author:    account.Username,
		Title:     input.Title,
		Content:   input.Content,
		Tags:      input.Tags,
	}

	if err = repositories.Post.Put(post); err != nil {
		return err
	}

	return nil
}
