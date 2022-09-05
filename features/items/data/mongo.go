package data

import (
	"context"
	"explore/mongodb/features/items"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemMongo struct {
	db *mongo.Database
}

func RepoItem(mongoDB *mongo.Database) items.Data {
	return &itemMongo{
		db: mongoDB,
	}
}

func (mongo *itemMongo) InsertData(core items.Core) error {
	itemModel := fromCore(core)
	_, err := mongo.db.Collection("items").InsertOne(context.TODO(), &itemModel)
	if err != nil {
		return err
	}
	return nil
}

func (mongo *itemMongo) GetData(checklistID string) ([]items.Core, error) {
	itemModel := []Item{}
	filter := bson.M{"checklist_id": checklistID}
	result, err := mongo.db.Collection("items").Find(context.TODO(), filter)
	if err != nil {
		return []items.Core{}, err
	}
	for result.Next(context.TODO()) {
		var itemTemp Item
		decodeErr := result.Decode(&itemTemp)
		if decodeErr != nil {
			return []items.Core{}, decodeErr
		}
		itemModel = append(itemModel, itemTemp)
	}
	return toCOreList(itemModel), nil
}
