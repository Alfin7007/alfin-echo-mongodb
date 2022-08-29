package data

import (
	"context"
	"explore/mongodb/features/users"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongodb struct {
	db *mongo.Database
}

func NewuserMongo(dbConn *mongo.Database) users.Data {
	return &mongodb{
		db: dbConn,
	}
}

func (mongo mongodb) InsertData(core users.Core) error {
	userModel := fromCore(core)
	ctx := context.Background()
	_, err := mongo.db.Collection("user").InsertOne(ctx, userModel)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
