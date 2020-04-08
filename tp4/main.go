package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wildcast/golang-primeros-pasos/tp4/precios"
	"log"
	"net/http"
	"strconv"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Precios API\n===========\n\n"))
	w.Write([]byte("/precios?pid=XXX&sid=SSS devuelve los precios para el producto XXX en el supermercado SSS.\n"))
	w.Write([]byte("/precios?pid=XXX&pid=YYY&sid=SSS&sid=TTT devuelve la suma de los precios para los productos XXX y YYY en los supermercados SSS y TTT.\n"))
}

func preciosHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	// Convierto los ids de string a numero y armo un slice
	pIDs := []int{}
	for _, p := range params["pid"] {
		if id, err := strconv.Atoi(p); err == nil {
			pIDs = append(pIDs, id)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "El ID de producto %s no es valido.", p)
			return
		}
	}
	if precios, err := precios.CalcularPrecios(pIDs, params["sid"]); err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(precios)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler)
	r.Path("/precios").
		Queries("pid", "{pid}"). // Id de producto
		Queries("sid", "{sid}"). // Id de supermercado
		HandlerFunc(preciosHandler).
		Name("precios")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
