package factory

import (
	userBussiness "explore/mongodb/features/users/bussiness"
	userData "explore/mongodb/features/users/data"
	userPresent "explore/mongodb/features/users/presentation"

	authBussiness "explore/mongodb/features/auth/bussiness"
	authData "explore/mongodb/features/auth/data"
	authPresent "explore/mongodb/features/auth/presentation"

	checklistBussiness "explore/mongodb/features/checklists/bussiness"
	checklistData "explore/mongodb/features/checklists/data"
	checklistPresent "explore/mongodb/features/checklists/presentation"

	itemBussiness "explore/mongodb/features/items/bussiness"
	itemData "explore/mongodb/features/items/data"
	itemPresent "explore/mongodb/features/items/presentation"

	"go.mongodb.org/mongo-driver/mongo"
)

type Presenter struct {
	UserPresenter      *userPresent.UserHandler
	AuthPresenter      *authPresent.AuthHandler
	ChecklistPresenter *checklistPresent.ChecklistBussiness
	ItemPresenter      *itemPresent.ItemHandler
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

	newItemData := itemData.RepoItem(db)
	newItemBussiness := itemBussiness.NewItemBussiness(newItemData)
	newItemPresentation := itemPresent.NewItemHandler(newItemBussiness)

	return Presenter{
		UserPresenter:      newUserPresentation,
		AuthPresenter:      newAuthPresentation,
		ChecklistPresenter: newChecklistPresentation,
		ItemPresenter:      &newItemPresentation,
	}
}
