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

func preciosWorker(preciosAObtener <-chan precioTienda, precios chan<- precioTienda, errores chan<- error) {
	for p := range preciosAObtener {
		fmt.Printf("Getting price of %d from %s\n", p.ID, p.Tienda)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(apiURL, p.Tienda, p.ID), nil)
		if err != nil {
			errores <- fmt.Errorf("No se pudo obtener el precio para el producto %d en el supermercado %s", p.ID, p.Tienda)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil || res.StatusCode != http.StatusOK {
			errores <- fmt.Errorf("No se pudo obtener el precio para el producto %d en el supermercado %s", p.ID, p.Tienda)
		}
		var precio precioTienda
		json.NewDecoder(res.Body).Decode(&precio)
		precios <- precio
		res.Body.Close()
	}
}

func preciosCollector(precios <-chan precioTienda, errores <-chan error, numFaltantes int, done chan<- error) {
	numPrecios := 0
	for {
		select {
		case p := <-precios:
			if _, ok := cache[p.ID]; !ok {
				cache[p.ID] = make(map[string]int)
			}
			cache[p.ID][p.Tienda] = p.Precio
			numPrecios++
			if numPrecios == numFaltantes {
				done <- nil
				return
			}
		case err := <-errores:
			done <- err
			return
		}
	}
}

func calcularCarritos(pIDs []int, sIDs []string) []Carrito {
	precios := make(map[string]int) // Mapeo de supermercado a precio total
	for _, pID := range pIDs {
		for _, sID := range sIDs {
			precios[sID] += cache[pID][sID]
		}
	}
	var res []Carrito
	for k, v := range precios {
		res = append(res, Carrito{k, v})
	}
	return res
}

func listaPreciosFaltantes(pIDs []int, sIDs []string) []precioTienda {
	preciosFaltantes := []precioTienda{}
	for _, pID := range pIDs {
		for _, sID := range sIDs {
			_, prodInCache := cache[pID]
			priceInCache := false
			if prodInCache {
				_, priceInCache = cache[pID][sID]
			}
			if !priceInCache {
				preciosFaltantes = append(preciosFaltantes, precioTienda{sID, pID, 0})
			}
		}
	}
	return preciosFaltantes
}

func obtenerPreciosFaltantesDeAPI(pIDs []int, sIDs []string) error {
	preciosFaltantes := listaPreciosFaltantes(pIDs, sIDs)

	if len(preciosFaltantes) > 0 {
		// Inicializacion de channels
		preciosAObtener := make(chan precioTienda, len(preciosFaltantes))
		precios := make(chan precioTienda)
		errores := make(chan error)
		done := make(chan error)

		// Precarga de precios a obtener
		for i := range preciosFaltantes {
			preciosAObtener <- preciosFaltantes[i]
		}
		close(preciosAObtener)

		// Inicializacion de workers
		w := 0
		for w < numWorkers {
			go preciosWorker(preciosAObtener, precios, errores)
			w++
		}

		go preciosCollector(precios, errores, len(preciosFaltantes), done)
		err := <-done
		if err != nil {
			return err
		}
	}
	return nil
}

// CalcularPrecios calcula los precios para la lista de IDs de productos pIDs en
// los supermercados indicados en sIDs y devuelve una lista de Carrito
func CalcularPrecios(pIDs []int, sIDs []string) ([]Carrito, error) {
	err := obtenerPreciosFaltantesDeAPI(pIDs, sIDs)
	if err != nil {
		return nil, err
	}
	return calcularCarritos(pIDs, sIDs), nil
}
