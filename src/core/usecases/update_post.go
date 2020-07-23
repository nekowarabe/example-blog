package usecases

import (
	"errors"

	"app/src/core/repositories"
)

// UpdatePostInput 更新文章所需參數
type UpdatePostInput struct {
	Token    string
	Username string
	ID       string
	Title    string
	Content  string
	Tags     []string
}

// UpdatePost 更新文章
func UpdatePost(input UpdatePostInput) error {
	account, err := repositories.Account.GetByToken(input.Token)
	if err != nil {
		return err
	}
	if account.Username != input.Username {
		return errors.New("unauthorization to update this post")
	}

	post, err := repositories.Post.Get(input.ID)
	if err != nil {
		return err
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Tags = input.Tags

	if err = repositories.Post.Put(post); err != nil {
		return err
	}

	return nil
}
