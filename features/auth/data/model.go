package data

import (
	"explore/mongodb/features/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

func (u User) toCore() auth.Core {
	return auth.Core{
		ID:       u.ID.Hex(),
		Email:    u.Email,
		Password: u.Password,
	}
}
