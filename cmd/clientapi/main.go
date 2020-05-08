package main

// Main or entry point for our application
import (
	"log"

	"net/http"

	"ms-client/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
