package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"ms-client/application"
	"ms-client/domain/entity"
	"ms-client/infrastructure/shared/request"

	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
	app    application.ClientAppInterface
}

// Server ...
type Server interface {
	Router() http.Handler
	addClient(w http.ResponseWriter, r *http.Request)
	searchClient(w http.ResponseWriter, r *http.Request)
	searchClientByEmail(w http.ResponseWriter, r *http.Request)
	updateClient(w http.ResponseWriter, r *http.Request)
}

// New ...
func New(
	clientApplication application.ClientAppInterface,
) Server {
	api := &api{app: clientApplication}
	router(api)
	return api
}

// routes settings
func router(a *api) {
	r := mux.NewRouter()
	r.HandleFunc("/clients", a.addClient).Methods(http.MethodPost)
	r.HandleFunc("/clients/{ID:[a-zA-Z0-9_]+}", a.searchClient).Methods(http.MethodGet)
	r.HandleFunc("/clients", a.updateClient).Methods(http.MethodPut)
	r.HandleFunc("/clients/email/{email}", a.searchClientByEmail).Methods(http.MethodGet)
	r.HandleFunc("/ms-client/health", healthCheck).Methods(http.MethodGet)
	a.router = r
}

func (a *api) Router() http.Handler {
	return a.router
}

//Clients ...
type Clients []entity.Client

func healthCheck(w http.ResponseWriter, r *http.Request) {
	var outp struct{ Result string }
	outp.Result = "Ok"
	json.NewEncoder(w).Encode(outp)

}

// AddClient function saves a new client
func (api *api) addClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cl request.ClientRequest
	err := decoder.Decode(&cl)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	if err := api.app.AddClient(r.Context(), cl); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't create the client")
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// SearchClient get a record by id
func (api *api) searchClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json") //ver que pasa sin esto
	if client, error := api.app.GetClient(r.Context(), vars["ID"]); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Client not found")
		return
	} else {
		_ = json.NewEncoder(w).Encode(client)
	}

}

// SearchClient by email
func (api *api) searchClientByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(vars["email"])
	if result, error := api.app.ValidateClientByEmail(r.Context(), vars["email"]); error != nil {
		log.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Email not found")
		return
	} else {
		_ = json.NewEncoder(w).Encode(result)
	}

}

// UpdateClient changes client record properties
func (api *api) updateClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cl request.ClientRequest
	err := decoder.Decode(&cl)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	if err := api.app.SaveClient(r.Context(), cl); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't update the client")
		return
	}
	w.WriteHeader(http.StatusCreated)
}
