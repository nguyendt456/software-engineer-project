package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:project231@mongo:27017/"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB !")
	return client
}

var DatabaseClient = GetDatabase(ConnectToDatabase(), "Project231")

func GetDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	database := client.Database(databaseName)
	return database
}

var UserCollection = GetCollection(DatabaseClient, "user")

func GetCollection(database *mongo.Database, collectionName string) *mongo.Collection {
	collection := database.Collection(collectionName)
	return collection
}
