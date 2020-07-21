package repositories

import (
	"app/src/core/entities"
)

var (
	// Account 提供帳號持久化操作
	Account AccountRepository
)

// AccountRepository 帳號的 Collection-Like 介面定義
type AccountRepository interface {
	// Add 新增帳號到集合中
	Add(account entities.Account) error
	// Exist 是否有指定 username 的帳號在集合中
	Exist(username string) (bool, error)
	// Clear 將集合的所有資料清除
	Clear()
}
