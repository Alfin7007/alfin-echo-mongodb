package request

import (
	"explore/mongodb/features/users"
)

type User struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (u User) ToCore() users.Core {
	userCore := users.Core{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
	return userCore
}
