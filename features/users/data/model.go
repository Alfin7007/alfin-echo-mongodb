package data

import (
	"explore/mongodb/features/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func fromCore(core users.Core) User {
	return User{
		Name:     core.Name,
		Email:    core.Email,
		Password: core.Password,
	}
}

func (u User) toCore() users.Core {
	return users.Core{
		ID:       u.ID.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
