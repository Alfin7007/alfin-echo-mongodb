package data

import (
	"context"
	"errors"
	"explore/mongodb/features/items"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (mongo *itemMongo) GetOneData(checklistID, id string) (items.Core, error) {
	itemModel := Item{}
	objectID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return items.Core{}, idErr
	}
	filter := bson.M{"checklist_id": checklistID, "_id": objectID}
	result := mongo.db.Collection("items").FindOne(context.TODO(), filter)
	err := result.Decode(&itemModel)
	if err != nil {
		return items.Core{}, err
	}
	return itemModel.toCore(), nil
}

func (mongo *itemMongo) DeleteOneData(checklist, id string) error {
	objectID, idErr := primitive.ObjectIDFromHex(id)
	if idErr != nil {
		return idErr
	}
	filter := bson.M{"checklist_id": checklist, "_id": objectID}
	result, err := mongo.db.Collection("items").DeleteOne(context.TODO(), filter)
	if err != nil || result.DeletedCount == 0 {
		return errors.New("failed delete")
	}
	return nil
}

func (mongo *itemMongo) UpdateOneData(core items.Core) error {
	objectID, idErr := primitive.ObjectIDFromHex(core.ID)
	if idErr != nil {
		return idErr
	}
	filter := bson.M{"checklist_id": core.ChecklistID, "_id": objectID}

	update := bson.M{"$set": bson.M{"status": core.Status}}
	_, err := mongo.db.Collection("items").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
