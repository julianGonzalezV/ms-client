package server

import (
	"encoding/json"
	"ms-client/pkg/domain/model"

	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

// Server ...
type Server interface {
	Router() http.Handler
}

// New ...
func New() Server {
	a := &api{}

	r := mux.NewRouter()
	//Configurando las rutas que debe resolver
	r.HandleFunc("/clients", a.fetchAllClients).Methods(http.MethodGet)
	// Note como Gorilla mux permite colocar expresione regulares para establecer reglas en los parámetros que
	// se pasen
	r.HandleFunc("/clients/{ID:[a-zA-Z0-9_]+}", a.fetchClient).Methods(http.MethodGet)
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

//Clients ...
type Clients []model.Client

func (a *api) fetchAllClients(w http.ResponseWriter, r *http.Request) {
	/*
		clients, _ := a.repository.FetchGophers()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(clients)*/
	clients := Clients{
		model.Client{ID: "C66708", FirstName: "Clari", Age: 32},
		model.Client{ID: "C1116235", FirstName: "Juli", Age: 32},
	}
	json.NewEncoder(w).Encode(clients)
}

func (a *api) fetchClient(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	gopher, err := a.repository.FetchGopherByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Gopher Not found")
		return
	}*/

	json.NewEncoder(w).Encode(model.Client{ID: "c111523", FirstName: "Juliano", Age: 32})
}
