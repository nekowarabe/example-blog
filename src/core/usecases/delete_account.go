package usecases

import (
	"errors"

	"app/src/core/repositories"
)

// DeleteAccountInput 移除帳號所需參數
type DeleteAccountInput struct {
	Token    string
	Username string
}

// DeleteAcoount 移除帳號
func DeleteAcoount(input DeleteAccountInput) error {
	account, err := repositories.Account.GetByToken(input.Token)
	if err != nil {
		return err
	}
	if account.Username != input.Username {
		return errors.New("unauthorization to delete this account")
	}

	if err = repositories.Account.Remove(account); err != nil {
		return err
	}

	return nil
}
