package memory

import "app/src/core/entities"

// Account 記憶體實作的 repositories.Account
type Account map[string]entities.Account

// NewAccountRepository 初始化帳號記憶體儲存庫
func NewAccountRepository() Account {
	return make(map[string]entities.Account)
}

// Add 新增帳號
func (repo Account) Add(account entities.Account) error {
	repo[account.Username] = account
	return nil
}

// Exist 指定 username 的帳號是否存在
func (repo Account) Exist(username string) (bool, error) {
	_, exist := repo[username]
	return exist, nil
}

// Clear 清除所有資料
func (repo Account) Clear() {
	for key := range repo {
		delete(repo, key)
	}
}
