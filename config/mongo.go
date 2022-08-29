package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB() *mongo.Database {
	mongoAuth := map[string]string{
		"dbUser": os.Getenv("DB_USER"),
		"dbPass": os.Getenv("DB_PASS"),
		"dbName": os.Getenv("DB_NAME"),
	}

	credential := options.Credential{
		AuthSource: mongoAuth["dbName"],
		Username:   mongoAuth["dbUser"],
		Password:   mongoAuth["dbPass"],
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, dbErr := mongo.Connect(context.TODO(), clientOpts)
	if dbErr != nil {
		panic(dbErr)
	}

	pingErr := client.Ping(context.TODO(), readpref.Primary())
	if pingErr != nil {
		panic(pingErr)
	}

	db := client.Database(mongoAuth["dbName"])

	return db
}
