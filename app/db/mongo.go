package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDatabase = os.Getenv("MONGO_DB")
var MongoClient *mongo.Client

func SetupMongo() {
	host := os.Getenv("MONGO_HOSTNAME")
	port := os.Getenv("MONGO_PORT")
	username := os.Getenv("MONGO_USERNAME")
	pass := url.QueryEscape(os.Getenv("MONGO_PASSWORD"))
	// DB := os.Getenv("MONGO_DB")

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", username, pass, host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = client
	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
}
