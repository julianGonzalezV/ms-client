package main

// Main or entry point for our application
import (
	"gopherapi/pkg/adding"
	"log"
	"ms-client/domain/repository"
	"ms-client/infrastructure/repositoryimpl"
	"ms-client/infrastructure/resource"
	"net/http"
	"os"
)

func main() {
	s := resource.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
	repo := initializeRepo(database, trc, gophers)
	// Services initialization, injecting despendencies
	addingService := adding.NewService(repo)
}

func initializeRepo(database *string) repository.ClientRepository {
	var repo repository.ClientRepository
	switch *database {
	case "mongo":
		repo = newClientMongoRepository()
	default:
		repo = nil // we can have an InMemory implementation
	}
	return repo
}

func newClientMongoRepository() repository.ClientRepository {
	mongoAddr := os.Getenv("MONGO_ADDR")
	client := repositoryimpl.Connect(mongoAddr)
	return repositoryimpl.NewRepository(client)
}
