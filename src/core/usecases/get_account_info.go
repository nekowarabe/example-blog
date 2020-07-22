package usecases

import (
	"errors"

	"app/src/core/entities"
	"app/src/core/repositories"
)

// GetAccountInfoInput 取得帳號資訊所需參數
type GetAccountInfoInput struct {
	Token    string
	Username string
}

// GetAccountInfoOutput 取得帳號資訊輸出參數
type GetAccountInfoOutput struct {
	Username string
	Tel      string
	Age      uint8
	Gender   entities.Gender
}

// GetAccountInfo 取得帳號的資訊
func GetAccountInfo(input GetAccountInfoInput) (output GetAccountInfoOutput, err error) {
	account, err := repositories.Account.GetByToken(input.Token)
	if err != nil {
		return output, err
	}
	if account.Username != input.Username {
		return output, errors.New("unauthorization to access this account")
	}

	output.Username = account.Username
	output.Tel = account.Tel
	output.Age = account.Age
	output.Gender = account.Gender

	return output, nil
}
