package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-client/domain/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://juigove:<pwd>@db-salud-digital-dllo-5r2lz.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

// Create saves a given client
func Create(c *model.Client) error {
	var client = getConnection()
	// Get a handle for your collection
	collection := client.Database("test").Collection("trainers")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

// Fetch return all clients saved in storage
func Fetch() ([]*model.Client, error) {
	return nil, nil
}

// Delete remove a client with given ID
func Delete(ID string) error {
	return nil
}

// Update modify client with given ID and given new data
func Update(ID string, c *model.Client) error {
	return nil
}

// FetchByID returns the client with given ID
func FetchByID(ID string) (*model.Client, error) {
	return nil, nil
}
