package usecases

import (
	"errors"

	"app/src/core/entities"
	"app/src/core/repositories"
)

// UpdateAccountInput 更新帳號所需參數
type UpdateAccountInput struct {
	Token    string
	Username string
	Password string
	Tel      string
	Age      uint8
	Gender   entities.Gender
}

// UpdateAccount 更新帳號
func UpdateAccount(input UpdateAccountInput) error {
	account, err := repositories.Account.GetByToken(input.Token)
	if err != nil {
		return err
	}
	if account.Username != input.Username {
		return errors.New("unauthorization to updayte this account")
	}

	// 如果有給值就設置
	if input.Password != "" {
		account.Password = input.Password
	}
	if input.Tel != "" {
		account.Tel = input.Tel
	}
	if input.Age != 0 {
		account.Age = input.Age
	}
	if input.Gender.IsValid() {
		account.Gender = input.Gender
	}

	if err = repositories.Account.Put(account); err != nil {
		return err
	}

	return nil
}
