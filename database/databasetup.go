package database

import (
	"log"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client{
	client, err := mongo.BSONAppenderNewClient(options.Client().ApplyURI("mongodb+srv://siddharth-tricon:Password@cluster0.ccgmbiv.mongodb.net/"))
	if err != nil {
		log.Fatal(err)
	}

	context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {	
		log.Fatal("Failed to connect to the database")
		return nil
	}
	fmt.Println("Connected to the database")
	return client

}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return collection
	
}

func ProductData(client *mongo.Client, collectioName string) *mongo.Collection{
	var productCollection *mongo.Client = client.Database("Ecommerce").Collection(collectionName)
	return produCtcollection
}
