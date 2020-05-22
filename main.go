package main

// Main or entry point for our application
import (
	"log"
	"ms-client/infrastructure/resource"
	"net/http"
)

func main() {
	s := resource.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
