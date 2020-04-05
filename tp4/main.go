package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type producto struct {
	Id     int    `json:"id"`
	Precio int    `json:"precio"`
	Tienda string `json:"tienda"`
}

type tiendaPrecio struct {
	Tienda string `json:"tienda"`
	Precio int    `json:"precio"`
}

type calcularPreciosResponse struct {
	tiendaPrecios []tiendaPrecio `json:"tiendaPrecios"`
}

func CalcularPrecios(ids []int, tiendas []string, resultadoChan chan<- []tiendaPrecio) {
	productosChan := make(chan producto)
	tasks := 0
	for _, tienda := range tiendas {
		for _, id := range ids {
			go GetProducto(tienda, id, productosChan)
			tasks++
		}
	}
	var productos []producto
	for i := 0; i < tasks; i++ {
		p := <-productosChan
		productos = append(productos, p)
	}
	close(productosChan)

	tiendaPrecios := CalcularTiendas(ids, productos)

	resultadoChan <- tiendaPrecios
}

func CalcularTiendas(ids []int, productos []producto) []tiendaPrecio {
	var tiendaPrecios map[string]tiendaPrecio
	tiendaPrecios = make(map[string]tiendaPrecio)

	uniqueIds := make(map[int]struct{})
	for _, id := range ids {
		uniqueIds[id] = struct{}{}
	}

	for _, v := range productos {
		if _, ok := uniqueIds[v.Id]; ok {
			if _, ok := tiendaPrecios[v.Tienda]; !ok {
				tiendaPrecios[v.Tienda] = tiendaPrecio{v.Tienda, 0}
			}
			carrito := tiendaPrecios[v.Tienda]
			carrito.Precio += v.Precio
			tiendaPrecios[v.Tienda] = carrito
		}
	}
	return valuesofmap(tiendaPrecios)
}

func valuesofmap(carritos map[string]tiendaPrecio) []tiendaPrecio {
	values := make([]tiendaPrecio, 0, len(carritos))
	for _, v := range carritos {
		values = append(values, v)
	}
	return values
}

func GetProducto(tienda string, id int, productosChan chan<- producto) {
	res, err := http.Get("https://productos-p6pdsjmljq-uc.a.run.app/" + tienda + "/productos/" + strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Hubo un error: %d", res.StatusCode)
	}
	defer res.Body.Close()
	var p producto
	json.NewDecoder(res.Body).Decode(&p)
	productosChan <- p
}

func CalcularPreciosHandler(w http.ResponseWriter, r *http.Request) {
	idsParam := r.FormValue("ids")
	tiendasParam := r.FormValue("tiendas")

	if idsParam == "" || tiendasParam == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ids []int
	err := json.Unmarshal([]byte(idsParam), &ids)
	if err != nil {
		log.Fatal(err)
	}

	var tiendas []string
	err2 := json.Unmarshal([]byte(tiendasParam), &tiendas)
	if err2 != nil {
		log.Fatal("asd ", err2)
	}

	if len(ids) == 0 || len(tiendas) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resultadosChan := make(chan []tiendaPrecio)
	go CalcularPrecios(ids, tiendas, resultadosChan)
	resultado := <-resultadosChan
	close(resultadosChan)

	resultadoJson, err := json.Marshal(resultado)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resultadoJson))
}

func main() {
	fmt.Println(`Example url: http://localhost:8080/calcular-precios?ids=[1,2]&tiendas=["Dia","Target","Coto","Disco","Jumbo","Macro","Nini"]`)
	r := mux.NewRouter()
	r.Path("/calcular-precios").Queries("ids", "tiendas").HandlerFunc(CalcularPreciosHandler).Name("CalcularPreciosHandler")
	r.Path("/calcular-precios").HandlerFunc(CalcularPreciosHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
