package storage

import (
	"context"
	"fmt"
	"os"
	"quote-me/config"
	"quote-me/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// QuoteStorage struct for retrieving quotes from db
type QuoteStorage struct {
	db     *mongo.Database
	client *mongo.Client
}

var storage *QuoteStorage

// Init function initialises the storage and connects to it
func Init() {
	storage = new(QuoteStorage)
	// Setup a mongodb connection
	client, err := mongo.NewClient(options.Client().ApplyURI(config.GetConfig().GetString("mongo.url")))
	if err != nil {
		fmt.Println("Cannot create the mongodb client", err)
		os.Exit(1)
	}
	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("Cannot connect to mongodb", err)
		os.Exit(1)
	}
	storage.client = client
	//Save the connected db in memory
	storage.db = storage.client.Database(config.GetConfig().GetString("mongo.db"))
	fmt.Println("Connected to db", config.GetConfig().GetString("mongo.url"))
}

// Quote method returns a random quote from the database
func (s *QuoteStorage) Quote() *models.QuoteModel {
	var results []models.QuoteModel

	quotes := s.db.Collection(config.GetConfig().GetString("mongo.quotesCollection"))
	//Use a sample function to find a random sample of size 1, ref: https://docs.mongodb.com/manual/reference/operator/aggregation/sample/#pipe._S_sample
	samplePipeline := bson.D{primitive.E{Key: "$sample", Value: bson.D{primitive.E{Key: "size", Value: 1}}}}

	cursor, err := quotes.Aggregate(context.Background(), mongo.Pipeline{samplePipeline})
	if err != nil {
		fmt.Println("Could not retrieve any documents", err)
		os.Exit(1)
	}
	for cursor.Next(context.Background()) {
		var result models.QuoteModel
		cursor.Decode(&result)
		results = append(results, result)
	}

	if len(results) == 0 {
		fmt.Println("No results available")
		return nil
	}

	return &results[0]
}

// Photo method returns a random photo from the database
func (s *QuoteStorage) Photo() *models.PhotoModel {
	var results []models.PhotoModel
	photos := s.db.Collection(config.GetConfig().GetString("mongo.photosCollection"))
	samplePipeline := bson.D{primitive.E{Key: "$sample", Value: bson.D{primitive.E{Key: "size", Value: 1}}}}

	cursor, err := photos.Aggregate(context.Background(), mongo.Pipeline{samplePipeline})
	if err != nil {
		fmt.Println("Could not retrieve any documents", err)
		os.Exit(1)
	}
	for cursor.Next(context.Background()) {
		var result models.PhotoModel
		cursor.Decode(&result)
		results = append(results, result)
	}

	if len(results) == 0 {
		fmt.Println("No results available")
		return nil
	}
	return &results[0]
}

// GetStorage method returns the current storage
func GetStorage() *QuoteStorage {
	return storage
}
