package repositoryimpl

import (
	"context"
	"fmt"
	"log"
	"ms-client/domain/entity"
	"ms-client/domain/repository"
	"ms-client/infrastructure/shared/customerror"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type cRepository struct {
	db *mongo.Client
}

// Connect retunrs a new connection to storage target
func Connect(addr string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(addr))
	if err != nil {
		fmt.Println("Error1!")
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("Error2!")
		log.Println(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Error3!", err)
		log.Println(err)
	}
	fmt.Println("Ping!")
	return client
}

// NewRepository creates a mongo repository with the necessary dependencies
func NewRepository(db *mongo.Client) repository.ClientRepository {
	return cRepository{db: db}
}

// Create saves a given client
func (r cRepository) Create(ctx context.Context, c *entity.Client) error {
	collection := r.db.Database("test").Collection("clients")
	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), c)
	if err != nil {
		fmt.Println("Error insertando", err)
		log.Fatal(err)
	}
	fmt.Println("insertResult", insertResult.InsertedID)
	return nil
}

// Fetch return all clients saved in storage
func (r cRepository) Fetch() ([]*entity.Client, error) {
	return nil, nil
}

// Delete remove a client with given ID
func (r cRepository) Delete(ID string) error {
	return nil
}

// Update modify client with given ID and given new data
func (r cRepository) Update(c *entity.Client) error {
	collection := r.db.Database("test").Collection("clients")
	updateResult, err := collection.UpdateOne(context.TODO(), bson.M{"id": c.IDNumber}, bson.D{
		{"$set", bson.D{{"firtsname", c.FirstName}}},
	})

	if err != nil {
		fmt.Println("Update Error", err)
		log.Panic(err)
	}
	fmt.Println("updateResult", updateResult.UpsertedID)

	return nil
}

// FetchByID returns the client with given ID
func (r cRepository) FetchByID(ID string) (*entity.Client, error) {
	collection := r.db.Database("test").Collection("clients")
	resultStruct := &entity.Client{}
	result := collection.FindOne(context.TODO(), bson.M{"idnumber": ID[1:len(ID)], "idtype": ID[0:1]})
	if result.Err() == mongo.ErrNoDocuments {
		return nil, customerror.ErrRecordNotFound
	}
	result.Decode(&resultStruct)
	return resultStruct, nil

}

func (r cRepository) FetchByEmail(email string) (bool, error) {
	collection := r.db.Database("test").Collection("clients")
	resultStruct := &entity.Client{}
	result := collection.FindOne(context.TODO(), bson.M{"contact.email": email})
	if result.Err() == mongo.ErrNoDocuments {
		return false, customerror.ErrRecordNotFound
	}
	result.Decode(&resultStruct)
	return true, nil

}
