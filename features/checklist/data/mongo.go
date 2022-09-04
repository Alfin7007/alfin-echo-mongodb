package data

import (
	"context"
	"explore/mongodb/features/checklist"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type checklistMongo struct {
	db *mongo.Database
}

func NewChecklistRepo(mongoDB *mongo.Database) checklist.Data {
	return &checklistMongo{
		db: mongoDB,
	}
}

func (mongo *checklistMongo) InsertData(core checklist.Core) error {
	checklistModel := fromCore(core)
	_, err := mongo.db.Collection("checklist").InsertOne(context.Background(), checklistModel)
	if err != nil {
		return err
	}
	return nil
}

func (mongo *checklistMongo) FindData(userID string) ([]checklist.Core, error) {
	filter := bson.D{{Key: "user_id", Value: userID}}
	cur, err := mongo.db.Collection("checklist").Find(context.TODO(), filter)

	if err != nil {
		return []checklist.Core{}, err
	}
	checklistModel := []Checklist{}
	for cur.Next(context.TODO()) {
		var temp Checklist
		if err := cur.Decode(&temp); err != nil {
			log.Fatal(err)
		}
		checklistModel = append(checklistModel, temp)
	}

	return toCoreList(checklistModel), nil
}
