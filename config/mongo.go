package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() *mongo.Client {
	mongoAuth := map[string]string{
		"dbUser": os.Getenv("DB_USER"),
		"dbPass": os.Getenv("DB_PASS"),
		"dbName": os.Getenv("DB_NAME"),
	}

	ctx := context.Background()
	connectionString := fmt.Sprintf("mongodb://%s:%s@localhost:27017",
		mongoAuth["dbUser"],
		mongoAuth["dbPass"])

	client, dbErr := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

	if dbErr != nil {
		panic(dbErr.Error())
	}

	return client
}
