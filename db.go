// db.go
package main

import (
	"context"
	"fmt"
	"github.com/JanSolo1/sw-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"reflect"
	"strings"
)

// Global database variable
var db *mongo.Database

// Simplify environment variable retrieval
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// InitMongoDB Initialize MongoDB connection and assign to global `db`
func init() {
	mongoHost := getMongoDBHost()
	mongoPort := getMongoDBPort()
	mongoDBName := getMongoDBName()
	mongoUser := getMongoUser()
	mongoPassword := getMongoPassword()

	// Include the username and password in the connection URI
	connectionURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", mongoUser, mongoPassword, mongoHost, mongoPort)

	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB")

	// Assign the connected database to the global `db`
	db = client.Database(mongoDBName)

	// Auto-provision the database with the StarWarsCharacter model
	AutoProvisionDB(context.Background(), db, []interface{}{models.StarWarsCharacter{}})
}

type MongoDBService struct {
	collection *mongo.Collection
}

func getMongoUser() string {
	return getEnv("MONGO_USER", "root")
}
func getMongoPassword() string {
	return getEnv("MONGO_PASSWORD", "password")

}
func getMongoDBHost() string {
	return getEnv("MONGO_HOST", "localhost")
}

func getMongoDBPort() string {
	return getEnv("MONGO_PORT", "27017")
}

func getMongoDBName() string {
	return getEnv("MONGO_DB", "test")
}

func NewMongoDBService(client *mongo.Client, dbName string, collectionName string) *MongoDBService {
	db = client.Database(dbName) // Assign the global db variable
	collection := db.Collection(collectionName)
	return &MongoDBService{collection: collection}
}

func (s *MongoDBService) Create(item interface{}) (*mongo.InsertOneResult, error) {
	res, err := s.collection.InsertOne(context.Background(), item)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (s *MongoDBService) Read(filter interface{}) (*mongo.SingleResult, error) {
	res := s.collection.FindOne(context.Background(), filter)
	return res, nil
}

func (s *MongoDBService) Update(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	res, err := s.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func (s *MongoDBService) Delete(filter interface{}) (*mongo.DeleteResult, error) {
	res, err := s.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, nil
}

func AutoProvisionDB(ctx context.Context, db *mongo.Database, models []interface{}) {
	for _, model := range models {
		modelType := reflect.TypeOf(model)
		if modelType.Kind() == reflect.Ptr {
			modelType = modelType.Elem() // Dereference pointer to get the struct
		}

		if modelType.Kind() != reflect.Struct {
			log.Printf("Provided model is not a struct: %v", model)
			continue // Skip non-struct types
		}

		// Use the struct name to lowercase as the collection name
		collectionName := strings.ToLower(modelType.Name())

		// Create the collection
		err := db.CreateCollection(ctx, collectionName, options.CreateCollection().SetCapped(false))
		if err != nil {
			log.Printf("Failed to auto-provision collection %s: %v", collectionName, err)
		} else {
			log.Printf("Successfully created collection: %s", collectionName)
		}
	}
}
