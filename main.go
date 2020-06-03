package main

// Main or entry point for our application
import (
	"flag"
	"fmt"
	"log"
	"ms-client/application/adding"
	"ms-client/domain/repository"
	"ms-client/infrastructure/repositoryimpl"
	"ms-client/infrastructure/resource"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var (
		defaultHost     = os.Getenv("CLIENTAPI_SERVER_HOST")
		defaultPort, _  = strconv.Atoi(os.Getenv("CLIENTAPI_SERVER_PORT"))
		defaultDatabase = os.Getenv("CLIENTAPI_SERVER_PORT")
	)

	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	database := flag.String("database", defaultDatabase, "initialize the api using the given db engine")

	/*
		s := resource.New()
	*/
	repo := initializeRepo(database)
	// Services initialization, injecting despendencies
	addingService := adding.NewService(repo)
	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	s := resource.New(
		addingService,
	)
	fmt.Println("The client server is running", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, s.Router()))
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
