package data

import (
	"context"
	"explore/mongodb/features/auth"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDB struct {
	db *mongo.Database
}

func AuthMongo(newDB *mongo.Database) auth.Data {
	return &mongoDB{
		db: newDB,
	}
}

func (mongo *mongoDB) FindData(email string) (auth.Core, error) {
	userModel := User{}
	filter := bson.M{"email": email}
	res := mongo.db.Collection("user").FindOne(context.Background(), filter)

	decodeErr := res.Decode(&userModel)
	if decodeErr != nil {
		return auth.Core{}, decodeErr
	}

	return userModel.toCore(), nil
}
