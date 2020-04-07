package main

import (
	"net/http"

	"github.com/wildcast/golang-primeros-pasos/tp4/price"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Path("/prices").
		Methods(http.MethodGet).
		Queries("market", "{marketID}").
		Queries("product", "{productID}").
		HandlerFunc(price.PricesHandler)

	http.ListenAndServe(":8080", router)
}
