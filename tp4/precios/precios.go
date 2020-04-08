package precios

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type precioTienda struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id"`
	Precio int    `json:"precio"`
}

const numWorkers = 8

// A modo de cache uso un mapa de ID de producto a mapa de ID de Super a Precio
var cache map[int](map[string]int)

const apiURL string = "https://productos-p6pdsjmljq-uc.a.run.app/%s/productos/%d"

// Carrito contiene una suma de precios para un supermercado dado
type Carrito struct {
	Tienda string `json: "tienda"`
	Precio int    `json: "precio"`
}

func init() {
	log.Println("Initializing Precios...")
	cache = make(map[int](map[string]int))
}

// func preciosWorker(preciosAObtener <-chan precioTienda, precios chan<- precioTienda, errores chan<- error) {

// }

// func preciosCollector(precios <-chan precioTienda, errores <-chan error, done chan<- bool) {

// }

// CalcularPrecios calcula los precios para la lista de IDs de productos pIDs en
// los supermercados indicados en sIDs y devuelve una lista de Carrito
func CalcularPrecios(pIDs []int, sIDs []string) ([]Carrito, error) {
	precios := make(map[string]int) // Mapeo de supermercado a precio total
	fmt.Println("CalcularPrecios", pIDs, sIDs)
	for _, pID := range pIDs {
		for _, sID := range sIDs {
			// fmt.Printf(apiURL, sID, pID)
			_, prodInCache := cache[pID]
			priceInCache := false
			if prodInCache {
				_, priceInCache = cache[pID][sID]
			}
			if !priceInCache {
				fmt.Printf("Getting price of %d from %s\n", pID, sID)

				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(apiURL, sID, pID), nil)
				if err != nil {
					return nil, fmt.Errorf("No se pudo obtener el precio para el producto %d en el supermercado %s", pID, sID)
				}
				res, err := http.DefaultClient.Do(req)
				if err != nil || res.StatusCode != http.StatusOK {
					return nil, fmt.Errorf("No se pudo obtener el precio para el producto %d en el supermercado %s", pID, sID)
				}
				var precio precioTienda
				json.NewDecoder(res.Body).Decode(&precio)
				if _, ok := cache[precio.ID]; !ok {
					cache[precio.ID] = make(map[string]int)
				}
				cache[precio.ID][precio.Tienda] = precio.Precio
				res.Body.Close()
			}
			precios[sID] += cache[pID][sID]
		}
	}
	// preciosAObtener := make(chan precioTienda)
	// precios := make(chan precioTienda)
	// suma := make(chan int)
	// w := 0
	// for w < numWorkers {
	// 	go preciosWorker(preciosAObtener, precios, errores)
	// 	w++
	// }
	// Transformo el map en la respuesta
	var res []Carrito
	for k, v := range precios {
		res = append(res, Carrito{k, v})
	}
	return res, nil
}
