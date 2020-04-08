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

type ListReq struct {
	Tienda []string `json:"tienda"`
	ID     []int    `json:"id"`
}

var productos []Producto

func homeLink(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(productos)
}

func calcular_precios(list_req ListReq) map[string]int {
	carritos := make(map[string]int)

	for _, prod := range productos {
		for _, tienda_req := range list_req.Tienda {
			if tienda_req == prod.Tienda {
				for _, id_req := range list_req.ID {
					if id_req == prod.ID {
						carritos[tienda_req] += prod.Precio
					}
				}
			}
		}
	}
	return carritos
}

func calcular_precios_handler(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(200)
	var reqData ListReq
	json.NewDecoder(req.Body).Decode(&reqData)
	json.NewEncoder(rw).Encode(calcular_precios(reqData))
}

func main() {
	//tiendas_full := []string{"dia", "jumbo", "supervea", "target", "wallmart", "carrefour",
	//	"disco", "macro", "nini", "coto", "whole foods market"}
	tiendas_test := []string{"dia", "jumbo"}
	prods := get_products(tiendas_test)
	fmt.Println(prods)
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", homeLink)
	r.HandleFunc("/calcular", calcular_precios_handler).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func get_products(tiendas []string) []Producto {
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
			productos = append(productos, new_prod)
			i++
		}
	}
	return productos
}
