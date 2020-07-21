package usecases

import (
	"errors"

	"app/src/core/repositories"

	uuid "github.com/satori/go.uuid"
)

// LoginInput 登入帳號所需參數
type LoginInput struct {
	Username string
	Password string
}

// LoginOutput 登入帳號輸出參數
type LoginOutput struct {
	Token string
}

// Login 登入帳號
func Login(input LoginInput) (output LoginOutput, err error) {
	account, err := repositories.Account.Get(input.Username)
	if err != nil {
		return output, err
	}
	if account.Password != input.Password {
		return output, errors.New("password wrong")
	}

	account.Token = uuid.NewV4().String()
	if err = repositories.Account.Put(account); err != nil {
		return output, err
	}

	output.Token = account.Token

	return output, nil
}
