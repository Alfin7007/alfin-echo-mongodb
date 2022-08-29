package bussiness

import (
	"explore/mongodb/features/users"

	"golang.org/x/crypto/bcrypt"
)

type UserBussiness struct {
	userData users.Data
}

func NewUserBussiness(newUserData users.Data) users.Bussiness {
	return &UserBussiness{
		userData: newUserData,
	}
}

func (b *UserBussiness) Register(core users.Core) error {
	bytePassword := []byte(core.Password)
	hash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	core.Password = string(hash)

	inserErr := b.userData.InsertData(core)
	if inserErr != nil {
		return inserErr
	}
	return nil
}
