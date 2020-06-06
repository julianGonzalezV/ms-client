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
		defaultDatabase = os.Getenv("CLIENT_DEFAULT_DATABASE")
		databaseName    = os.Getenv("CLIENT_DATABASE_NAME")
	)

	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	database := flag.String("database", defaultDatabase, "initialize the api using the given db engine")

	fmt.Println("CLIENTAPI_SERVER_HOST", defaultHost)
	fmt.Println("CLIENTAPI_SERVER_PORT", *port)
	fmt.Println("CLIENT_DEFAULT_DATABASE", defaultDatabase)
	fmt.Println("CLIENT_DATABASE_NAME", databaseName)
	/*
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			fmt.Println(pair)
		}*/
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
	fmt.Println("initializeRepo", *database)
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
	fmt.Println("mongoAddr", mongoAddr)
	client := repositoryimpl.Connect(mongoAddr)
	return repositoryimpl.NewRepository(client)
}
