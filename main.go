package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wildcast/golang-primeros-pasos/tp4"
)

func serve(w http.ResponseWriter, r *http.Request) {
	type rJson struct {
		IdProductos []int `json:"idProductos"`
		IdTiendas   []string `json:"idTiendas"`
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rJsonVal := rJson{}
	err = json.Unmarshal(body, &rJsonVal)
	fmt.Println(rJsonVal)
	if err != nil {
		panic(err)
	}
	rs, err := tp4.CalcPreciosResponse(rJsonVal.IdProductos, rJsonVal.IdTiendas)
	
	if err != nil {
		panic(err)
	}
	w.Write(rs)
}
func main() {
	http.HandleFunc("/", serve)
	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		panic(err)
	}
}
