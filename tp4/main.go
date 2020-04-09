package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

type Producto struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id"`
	Precio int    `json:"precio"`
}

type ProductoMap struct {
	Tienda string
	ID     int
	Precio int
}

type ListReq struct {
	Tienda []string `json:"tienda"`
	ID     []int    `json:"id"`
}

type key_prod struct {
	Tienda string
	ID     int
}

var productos map[key_prod]int

func homeLinkHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	defer req.Body.Close()
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(productos)
}

func calcularPrecios(list_req ListReq) (map[string]int, error) {
	carritos := make(map[string]int)
	var tiendasIds []key_prod
	var missingIds []key_prod
	var err error

	// Build a list of Tienda ID
	for _, tienda := range list_req.Tienda {
		for _, id_req := range list_req.ID {
			tiendasIds = append(tiendasIds, key_prod{tienda, id_req})
		}
	}

	// Get the price of the Tienda-ID if not Tienda-ID it's stack on a list
	for _, tiendaId := range tiendasIds {
		if val, ok := productos[tiendaId]; ok {
			carritos[tiendaId.Tienda] += val
		} else {
			missingIds = append(missingIds, tiendaId)
		}
	}

	//If the missing Tienda-ID is not empty throw error
	if len(missingIds) > 0 {
		err = fmt.Errorf("No se encotraron los valores %v", missingIds)
	}

	return carritos, err
}

func calcularPreciosHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(200)
	defer req.Body.Close()
	var reqData ListReq
	json.NewDecoder(req.Body).Decode(&reqData)
	retData, err := calcularPrecios(reqData)
	if err != nil {
		json.NewEncoder(rw).Encode(map[string]string{
			"Error": err.Error(),
		})
		return
	}
	json.NewEncoder(rw).Encode(retData)
}

func notFoundHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusNotFound)
	json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": "request not found",
	})
}

func main() {
	productos = make(map[key_prod]int)
	//tiendas_full := []string{"dia", "jumbo", "supervea", "target", "wallmart", "carrefour",
	//	"disco", "macro", "nini", "coto", "whole foods market"}
	tiendas_test := []string{"dia", "jumbo"}
	prods := get_products(tiendas_test)
	fmt.Println(prods)
	r := mux.NewRouter()
	r.HandleFunc("/", homeLinkHandler)
	r.HandleFunc("/calcular", calcularPreciosHandler).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	http.ListenAndServe(":8080", r)
}

func get_products(tiendas []string) map[key_prod]int {
	URL := "productos-p6pdsjmljq-uc.a.run.app"

	for _, tienda := range tiendas {
		var i int
		base_url := path.Join(URL, tienda, "productos")
		for {
			url_string := path.Join(base_url, strconv.Itoa(i))
			https := "https://"
			url_string = fmt.Sprint(https, url_string)
			fmt.Println(url_string)
			resp, err := http.Get(url_string)
			if err != nil {
				fmt.Println(err)
				break
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Println(resp.StatusCode)
				break
			}
			var new_prod Producto
			json.NewDecoder(resp.Body).Decode(&new_prod)

			//Decode de Tienda-ID on as a key of product
			kp := key_prod{Tienda: new_prod.Tienda, ID: new_prod.ID}
			productos[kp] = new_prod.Precio
			i++
		}
	}
	return productos
}
