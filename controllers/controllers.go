package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"automator/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "automation"
const CollectionName = "Tasks"

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongo_uri := os.Getenv("MONGO_URL")
	clientOption := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Successful")
	collection = client.Database(dbName).Collection(CollectionName)
	fmt.Println("Database is ready..")

}

func GetTask(taskAddress string) *models.AutoTask {
	var task *models.AutoTask
	data := collection.FindOne(context.Background(), bson.M{"address": taskAddress})
	data.Decode(&task)
	return task
}
