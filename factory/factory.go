package factory

import (
	userBussiness "explore/mongodb/features/users/bussiness"
	userData "explore/mongodb/features/users/data"
	userPresent "explore/mongodb/features/users/presentation"

	"go.mongodb.org/mongo-driver/mongo"
)

type Presenter struct {
	UserPresenter *userPresent.UserHandler
}

func InitFactory(db *mongo.Database) Presenter {
	newUserData := userData.NewuserMongo(db)
	newUserBussiness := userBussiness.NewUserBussiness(newUserData)
	newUserPresentation := userPresent.NewUseHandler(newUserBussiness)

	return Presenter{
		UserPresenter: newUserPresentation,
	}
}
