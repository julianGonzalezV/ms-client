package main

// Main or entry point for our application
import (
	"flag"
	"fmt"
	"log"
	"ms-client/application"
	"ms-client/domain/repository"
	"ms-client/domain/service"
	"ms-client/infrastructure/controller"
	"ms-client/infrastructure/repositoryimpl"
	"net/http"
	"os"
	"strconv"
)

/// initializeRepo return a cliente repository based on database type name
func initializeRepo(database *string) repository.ClientRepository {
	switch *database {
	case "mongo":
		return newClientMongoRepository()
	default:
		return nil // we can have several implementation like in memory, postgress etc
	}
}

/// newClientMongoRepository returns the mongoDB implementation
func newClientMongoRepository() repository.ClientRepository {
	mongoAddr := os.Getenv("DATABASE_CONN")
	fmt.Println("mongoAddr => ", mongoAddr)
	client := repositoryimpl.Connect(mongoAddr)
	return repositoryimpl.NewRepository(client)
}

func ClientHandler() {
	var (
		defaultHost    = os.Getenv("CLIENTAPI_SERVER_HOST")
		defaultPort, _ = strconv.Atoi(os.Getenv("CLIENTAPI_SERVER_PORT"))
		dbDriver       = os.Getenv("DATABASE_DRIVER")
	)
	host := flag.String("host", defaultHost, "define host of the server")
	port := flag.Int("port", defaultPort, "define port of the server")
	database := flag.String("database", dbDriver, "initialize the api using the given db engine")
	fmt.Println(defaultHost)
	fmt.Println(defaultPort)
	fmt.Println(dbDriver)
	// Injecting service and repo to Application Layer
	applicationL := application.NewClientApp(service.NewClientService(initializeRepo(database)))

	httpAddr := fmt.Sprintf("%s:%d", *host, *port)

	// Injecting server configuration
	server := controller.New(applicationL)

	// Next two linea are for AWS Conf
	/*http.Handle("/", server.Router())
	log.Fatal(gateway.ListenAndServe(httpAddr, nil))*/

	// Next line is for Local conf
	log.Fatal(http.ListenAndServe(httpAddr, server.Router()))
	fmt.Println("The client server is running", httpAddr)

}

func main() {
	fmt.Println("V1.3.0")
	ClientHandler()
}
