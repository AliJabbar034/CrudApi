package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var tasksCollection *mongo.Collection

func init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variable: ", err)
	}
	url := os.Getenv("MONGO_URL")

	if url == "" {
		log.Fatal("Inavlid Url", url)
	}

	client, ok := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if ok != nil {
		log.Fatal("Cannot connect to database")
	}
	userCollection = client.Database("TODO").Collection("User")
	tasksCollection = client.Database("TODO").Collection("Tasks")

	fmt.Println("Connection to database successfully")

}

func GetCollection() (*mongo.Collection, *mongo.Collection) {
	return userCollection, tasksCollection
}
