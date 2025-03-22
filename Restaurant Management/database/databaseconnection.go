package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client = DBInstances()

func DBInstances() *mongo.Client {
	MongoDb := "mongodb://localhost:27017"
	fmt.Println(MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.database("restaurant").Collection(collectionName)
	return collection
}
