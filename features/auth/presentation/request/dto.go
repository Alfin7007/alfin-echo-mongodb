package request

import "explore/mongodb/features/auth"

type Auth struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (a Auth) ToCore() auth.Core {
	return auth.Core{
		Email:    a.Email,
		Password: a.Password,
	}
}
