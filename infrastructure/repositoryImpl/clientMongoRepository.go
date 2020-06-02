package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-client/domain/model"
	"ms-client/domain/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type cRepository struct {
	db *mongo.Client
}

func Connect(addr string) *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(addr)

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

// NewRepository creates a mongo repository with the necessary dependencies
func NewRepository(db *mongo.Client) repository.ClientRepository {
	return cRepository{db: db}
}

// Create saves a given client
func (r cRepository) Create(ctx context.Context, c *model.Client) error {
	//var client = getConnection()
	// Get a handle for your collection
	collection := r.db.Database("salud-digital-dllo").Collection("clients")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

// Fetch return all clients saved in storage
func (r cRepository) Fetch() ([]*model.Client, error) {
	return nil, nil
}

// Delete remove a client with given ID
func (r cRepository) Delete(ID string) error {
	return nil
}

// Update modify client with given ID and given new data
func (r cRepository) Update(ID string, c *model.Client) error {
	return nil
}

// FetchByID returns the client with given ID
func (r cRepository) FetchByID(ID string) (*model.Client, error) {
	return nil, nil
}
