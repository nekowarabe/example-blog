package repositories

import (
	"app/src/core/entities"
)

var (
	// Account 提供帳號持久化操作
	Account AccountRepository
	// Post 提供文章持久化操作
	Post PostRepository
)

// AccountRepository 帳號的 Collection-Like 介面定義
type AccountRepository interface {
	// Get 取得指定 username 的帳號
	Get(username string) (entities.Account, error)
	// GetByToken 取得指定 token 的帳號
	GetByToken(token string) (entities.Account, error)
	// Put 將帳號放入集合中，如果已存在就取代
	Put(account entities.Account) error
	// Remove 將指定帳號移除集合
	Remove(account entities.Account) error
	// Exist 是否有指定 username 的帳號在集合中
	Exist(username string) (bool, error)
	// Clear 將集合的所有資料清除
	Clear()
}

// PostRepository 文章的 Collection-Like 介面定義
type PostRepository interface {
	// Put 將文章放入集合中，如果已存在就取代
	Put(post entities.Post) error
}
