package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wildcast/golang-primeros-pasos/tp4"
	"log"
	"net/http"
)

const SERVER_HTTP = ":8000"

type request struct {
	Ids    []int32  `json:"ids"`
	Tiendas []string `json:"tiendas"`
}

func serveHandle(w http.ResponseWriter, r *http.Request) {
	request := &request{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		log.Printf("No se pudo decodificar el request, error fue: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(request.Ids) <= 0 || len(request.Tiendas) <= 0 {
		log.Printf("No se pudo procesar el request")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	carritos := tp4.CalcularPorTiendas(request.Ids,request.Tiendas)

	json.NewEncoder(w).Encode(carritos)
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/tiendas", serveHandle).Methods("POST")

	log.Fatal(http.ListenAndServe(SERVER_HTTP, r))
}
