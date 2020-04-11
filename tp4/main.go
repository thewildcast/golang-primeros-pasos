package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	supermarket "github.com/wildcast/golang-primeros-pasos/tp4/handler"
)

//curl -XGET http://localhost:8000/\?supermarkets\=dia\&productids\=22
func main() {
	r := mux.NewRouter()
	r.Queries("supermarkets", "productids")

	r.HandleFunc("/", supermarket.PriceHandler)

	log.Fatal(http.ListenAndServe(":8001", r))
}
