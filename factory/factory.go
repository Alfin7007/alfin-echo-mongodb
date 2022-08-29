package factory

import (
	userBussiness "explore/mongodb/features/users/bussiness"
	userData "explore/mongodb/features/users/data"
	userPresent "explore/mongodb/features/users/presentation"

	authBussiness "explore/mongodb/features/auth/bussiness"
	authData "explore/mongodb/features/auth/data"
	authPresent "explore/mongodb/features/auth/presentation"

	"go.mongodb.org/mongo-driver/mongo"
)

type Presenter struct {
	UserPresenter *userPresent.UserHandler
	AuthPresenter *authPresent.AuthHandler
}

func InitFactory(db *mongo.Database) Presenter {
	newUserData := userData.NewuserMongo(db)
	newUserBussiness := userBussiness.NewUserBussiness(newUserData)
	newUserPresentation := userPresent.NewUseHandler(newUserBussiness)

	newAuthData := authData.AuthMongo(db)
	newAuthBussiness := authBussiness.NewAuthBussiness(newAuthData)
	newAuthPresentation := authPresent.NewAuthHandler(newAuthBussiness)

	return Presenter{
		UserPresenter: newUserPresentation,
		AuthPresenter: newAuthPresentation,
	}
}
