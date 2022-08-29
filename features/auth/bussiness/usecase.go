package bussiness

import (
	"errors"
	"explore/mongodb/features/auth"
	"explore/mongodb/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type AuthBussiness struct {
	authData auth.Data
}

func NewAuthBussiness(userData auth.Data) auth.Bussiness {
	return &AuthBussiness{
		authData: userData,
	}
}

func (b *AuthBussiness) Login(core auth.Core) (string, string, error) {
	res, err := b.authData.FindData(core.Email)
	if err != nil {
		return "", "", err
	}
	if res.Password == "" {
		return "", "", errors.New("user not found")
	}
	compareErr := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(core.Password))
	if compareErr != nil {
		return "", "", errors.New("wrong password")
	}

	token, tokenErr := middlewares.CreateToken(res.ID)
	if tokenErr != nil {
		return "", "", tokenErr
	}

	return res.ID, token, nil

}
