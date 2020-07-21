package usecases

import (
	"errors"

	"app/src/core/entities"
	"app/src/core/repositories"
)

// RegisterInput 註冊帳號所需參數
type RegisterInput struct {
	Username string
	Password string
	Tel      string
	Age      uint8
	Gender   entities.Gender
}

// Register 註冊帳號
func Register(input RegisterInput) error {
	isExist, err := repositories.Account.Exist(input.Username)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("account already exists")
	}

	account := entities.Account{
		Username: input.Username,
		Password: input.Password,
		Tel:      input.Tel,
		Age:      input.Age,
		Gender:   input.Gender,
	}

	if err = repositories.Account.Add(account); err != nil {
		return err
	}

	return nil
}
