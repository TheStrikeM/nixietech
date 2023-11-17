package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log/slog"
	"net/http"
	"nixietech/internal/fetcher"
	"nixietech/internal/storage"
)

type RestApiSrv struct {
	Fetcher fetcher.Fetcher
	Router  *mux.Router
	Port    string
}

func New(fetcher fetcher.Fetcher, port string) *RestApiSrv {
	return &RestApiSrv{
		Fetcher: fetcher,
		Router:  mux.NewRouter().StrictSlash(true),
		Port:    port,
	}
}

func (rest *RestApiSrv) StartServer() error {
	slog.Info("[Server] Starting rest api server...")

	rest.Router.HandleFunc("/clocks", rest.AllClocksHandler)
	rest.Router.HandleFunc("/clock/{id}", rest.ClockHandler)
	rest.Router.HandleFunc("/create-order", rest.CreateOrder).Methods("POST")

	err := http.ListenAndServe(":8080", rest.Router)
	if err != nil {
		return err
	}
	return nil
}

func (rest *RestApiSrv) AllClocksHandler(w http.ResponseWriter, r *http.Request) {
	clocks, err := rest.Fetcher.GetAll()
	if err != nil {
		slog.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clocks)
}

func (rest *RestApiSrv) ClockHandler(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	clock, err := rest.Fetcher.ClockById(ID)
	if err != nil {
		slog.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clock)
}

func (rest *RestApiSrv) CreateOrder(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var order storage.OrderWithoutId[string]
	json.Unmarshal(requestBody, &order)

	result, err := rest.Fetcher.CreateOrder(&order)
	if err != nil {
		slog.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
