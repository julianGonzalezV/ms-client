package main

// Main or entry point for our application
import (
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
	mongoAddr := os.Getenv("COCKROACH_ADDR")
	client := repositoryimpl.getConnection(mongoAddr)
	return repositoryimpl.NewRepository(client)
}
