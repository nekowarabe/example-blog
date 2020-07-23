package memory

import (
	"errors"

	"app/src/core/entities"
)

// Post 記憶體實作的 repositories.Post
type Post map[string]entities.Post

// NewPostRepository 初始化文章記憶體儲存庫
func NewPostRepository() Post {
	return make(map[string]entities.Post)
}

// Get 取得指定 ID 的文章
func (repo Post) Get(ID string) (entities.Post, error) {
	if post, exist := repo[ID]; exist {
		return post, nil
	}
	return entities.Post{}, errors.New("can't find post")
}

// Put 將文章放入
func (repo Post) Put(post entities.Post) error {
	repo[post.ID] = post
	return nil
}
