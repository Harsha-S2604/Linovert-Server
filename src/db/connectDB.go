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
	clientOptions := options.Client().ApplyURI("mongodb+srv://arix2406:6Dwy4jq8h4s26us20@linovert.ul3hz.mongodb.net/LinovertDB?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Sorry something went wrong. Couldn't connect to database")
	} else {
		log.Println("Successfully connected to the database :)")
	}
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatal("Sorry something went wrong. Couldn't connect to database")
	// } else {
	// 	log.Println("Successfully connected to database :)")
	// }

	Database = client.Database("LinovertDB")
}
