package data

import (
	"context"
	"explore/mongodb/features/users"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	db *mongo.Database
}

func NewuserMongo(dbConn *mongo.Database) users.Data {
	return &mongoDB{
		db: dbConn,
	}
}

func (mongo mongoDB) InsertData(core users.Core) error {
	userModel := fromCore(core)
	ctx := context.Background()
	_, err := mongo.db.Collection("user").InsertOne(ctx, userModel)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (mongo mongoDB) GetData(id string) (users.Core, error) {
	userModel := User{}
	objectID, objectIDErr := primitive.ObjectIDFromHex(id)
	if objectIDErr != nil {
		return users.Core{}, objectIDErr
	}
	filter := bson.M{"_id": objectID}
	res := mongo.db.Collection("user").FindOne(context.Background(), filter)
	decodeErr := res.Decode(&userModel)

	if decodeErr != nil {
		return users.Core{}, decodeErr
	}
	return userModel.toCore(), nil
}
