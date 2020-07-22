package memory

import (
	"errors"

	"app/src/core/entities"
)

// Account 記憶體實作的 repositories.Account
type Account map[string]entities.Account

// NewAccountRepository 初始化帳號記憶體儲存庫
func NewAccountRepository() Account {
	return make(map[string]entities.Account)
}

// Get 取得指定 username 的帳號
func (repo Account) Get(username string) (entities.Account, error) {
	account, exist := repo[username]
	if !exist {
		return entities.Account{}, errors.New("can't find this account")
	}

	return account, nil
}

// GetByToken 取得指定 token 的帳號
func (repo Account) GetByToken(token string) (entities.Account, error) {
	return repo.Get(token)
}

// Put 將帳號放入
func (repo Account) Put(account entities.Account) error {
	repo[account.Username] = account
	repo[account.Token] = account
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
