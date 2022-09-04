package response

import "explore/mongodb/features/users"

type User struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromCore(core users.Core) User {
	return User{
		ID:    core.ID,
		Name:  core.Name,
		Email: core.Email,
	}
}
