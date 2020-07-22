package entities

import "time"

// Post 文章定義
type Post struct {
	ID        string // 文章編號，不可重複
	CreatedAt time.Time
	UpdatedAt time.Time
	Author    string // 文章的作者，對應至 Account.Username
	Title     string
	Content   string
	Tags      []string
}
