package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-client/domain/model/client"
	"ms-client/domain/repository"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type cRepository struct {
	db *mongo.Client
}

// Connect retunrs a new connection to storage target
func Connect(addr string) *mongo.Client {
	fmt.Println("Connect to MongoDB!", addr)
	// Set client options
	//clientOptions := options.Client().ApplyURI(addr)
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		fmt.Println("Error1!")
		log.Fatal(err)
	}
	fmt.Println("voy a contexto")
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	fmt.Println("conectando")
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error2!")
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error3!", err)
		log.Fatal(err)
	}
	fmt.Println("Ping!")
	return client
}

// NewRepository creates a mongo repository with the necessary dependencies
func NewRepository(db *mongo.Client) repository.ClientRepository {
	return cRepository{db: db}
}

// Create saves a given client
func (r cRepository) Create(ctx context.Context, c *client.Client) error {
	//var client = getConnection()
	// Get a handle for your collection
	fmt.Println("Create", c)
	collection := r.db.Database("test").Collection("clients")
	fmt.Println("collection")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), c)
	fmt.Println("insertResult")
	if err != nil {
		fmt.Println("Error insertando", err)
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

// Fetch return all clients saved in storage
func (r cRepository) Fetch() ([]*client.Client, error) {
	return nil, nil
}

// Delete remove a client with given ID
func (r cRepository) Delete(ID string) error {
	return nil
}

// Update modify client with given ID and given new data
func (r cRepository) Update(ID string, c *client.Client) error {
	return nil
}

// FetchByID returns the client with given ID
func (r cRepository) FetchByID(ID string) (*client.Client, error) {
	return nil, nil
}
