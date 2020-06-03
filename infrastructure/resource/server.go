package resource

import (
	"encoding/json"
	"log"
	"ms-client/application/adding"
	"ms-client/domain/model/client"

	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
	adding adding.Service
}

// Server ...
type Server interface {
	Router() http.Handler
	AddClient(w http.ResponseWriter, r *http.Request)
}

// New ...
func New(
	aS adding.Service,
) Server {
	a := &api{adding: aS}
	router(a)
	return a
}

func router(a *api) {
	r := mux.NewRouter()
	// Configurando las rutas que debe resolver
	r.HandleFunc("/clients", a.AddClient).Methods(http.MethodPost)
	/*r.HandleFunc("/clients", a.fetchAllClients).Methods(http.MethodGet)
	Note como Gorilla mux permite colocar expresione regulares para establecer reglas en los parámetros que
	// se pasen
	r.HandleFunc("/clients/{ID:[a-zA-Z0-9_]+}", a.fetchClient).Methods(http.MethodGet)*/
	a.router = r
}

func (a *api) Router() http.Handler {
	return a.router
}

//Clients ...
type Clients []client.Client

func (a *api) fetchAllClients(w http.ResponseWriter, r *http.Request) {
	clients := Clients{
		client.Client{ID: "C66708", FirstName: "Clari", Age: 32},
		client.Client{ID: "C1116235", FirstName: "Juli", Age: 32},
	}
	json.NewEncoder(w).Encode(clients)
}

func (a *api) fetchClient(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(client.Client{ID: "c111523", FirstName: "Juliano", Age: 32})
}

type addClientRequest struct {
	ID     string `json:"ID"`
	IDType string `json:"IDType"`
	Name   string `json:"name"`
}

// AddClient function saves a new client
func (a *api) AddClient(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cl addClientRequest
	err := decoder.Decode(&cl)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
		return
	}
	if err := a.adding.AddClient(r.Context(), cl.ID, cl.IDType, cl.Name); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode("Can't create the client")
		return
	}
	w.WriteHeader(http.StatusCreated)
}
