package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Database *mongo.Database
	ctx      = context.TODO()
)

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Sorry something went wrong. Couldn't connect to database")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Sorry something went wrong. Couldn't connect to database")
	} else {
		log.Println("Successfully connected to database :)")
	}

	Database = client.Database("LinovertDB")
}
