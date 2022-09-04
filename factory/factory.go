package factory

import (
	userBussiness "explore/mongodb/features/users/bussiness"
	userData "explore/mongodb/features/users/data"
	userPresent "explore/mongodb/features/users/presentation"

	authBussiness "explore/mongodb/features/auth/bussiness"
	authData "explore/mongodb/features/auth/data"
	authPresent "explore/mongodb/features/auth/presentation"

	checklistBussiness "explore/mongodb/features/checklist/bussiness"
	checklistData "explore/mongodb/features/checklist/data"
	checklistPresent "explore/mongodb/features/checklist/presentation"

	"go.mongodb.org/mongo-driver/mongo"
)

type Presenter struct {
	UserPresenter      *userPresent.UserHandler
	AuthPresenter      *authPresent.AuthHandler
	ChecklistPresenter *checklistPresent.ChecklistBussiness
}

func InitFactory(db *mongo.Database) Presenter {
	newUserData := userData.NewuserMongo(db)
	newUserBussiness := userBussiness.NewUserBussiness(newUserData)
	newUserPresentation := userPresent.NewUseHandler(newUserBussiness)

	newAuthData := authData.AuthMongo(db)
	newAuthBussiness := authBussiness.NewAuthBussiness(newAuthData)
	newAuthPresentation := authPresent.NewAuthHandler(newAuthBussiness)

	newChecklistData := checklistData.NewChecklistRepo(db)
	newChecklistBussiness := checklistBussiness.ChecklistBussiness(newChecklistData)
	newChecklistPresentation := checklistPresent.ChecklistHandler(newChecklistBussiness)

	return Presenter{
		UserPresenter:      newUserPresentation,
		AuthPresenter:      newAuthPresentation,
		ChecklistPresenter: newChecklistPresentation,
	}
}
