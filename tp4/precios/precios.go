package precios

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type precioTienda struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id"`
	Precio int    `json:"precio"`
}

const numWorkers = 8

const apiURL string = "https://productos-p6pdsjmljq-uc.a.run.app/%s/productos/%d"

// Carrito contiene una suma de precios para un supermercado dado
type Carrito struct {
	Tienda string `json: "tienda"`
	Precio int    `json: "precio"`
}

// Service es la estructura que usamos para persistir el estado del
// servicio
type Service struct {
	// A modo de cache uso un mapa de ID de producto a mapa de ID de Super a Precio
	cache map[int](map[string]int)
	mux   sync.Mutex
}

// NewService inicializa una estructura que permite utilizar los
// servicios de este paquete
func NewService() Service {
	return Service{cache: make(map[int](map[string]int))}
}

func (s *Service) updateCache(precio precioTienda) {
	s.mux.Lock()
	if _, ok := s.cache[precio.ID]; !ok {
		s.cache[precio.ID] = make(map[string]int)
	}
	s.cache[precio.ID][precio.Tienda] = precio.Precio
	s.mux.Unlock()
}

func (s *Service) getFromCache(pID int, sID string) (int, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	_, prodInCache := s.cache[pID]
	if !prodInCache {
		return 0, false
	}
	price, priceInCache := s.cache[pID][sID]
	return price, priceInCache
}

func (s *Service) preciosWorker(preciosAObtener <-chan precioTienda, precios chan<- precioTienda, errores chan<- error, abort *bool) {
	for p := range preciosAObtener {
		fmt.Printf("Getting price of %d from %s\n", p.ID, p.Tienda)
		// abort indica que tenemos que abortar la ejecucion
		if *abort {
			return
		}
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

func (s *Service) preciosCollector(precios <-chan precioTienda, errores <-chan error, numFaltantes int, done chan<- error) {
	numPrecios := 0
	for {
		select {
		case p := <-precios:
			s.updateCache(p)
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

func (s *Service) calcularCarritos(pIDs []int, sIDs []string) []Carrito {
	precios := make(map[string]int) // Mapeo de supermercado a precio total
	for _, pID := range pIDs {
		for _, sID := range sIDs {
			precio, _ := s.getFromCache(pID, sID)
			precios[sID] += precio
		}
	}
	var res []Carrito
	for k, v := range precios {
		res = append(res, Carrito{k, v})
	}
	return res
}

func (s *Service) listaPreciosFaltantes(pIDs []int, sIDs []string) []precioTienda {
	preciosFaltantes := []precioTienda{}
	for _, pID := range pIDs {
		for _, sID := range sIDs {
			_, inCache := s.getFromCache(pID, sID)
			if !inCache {
				preciosFaltantes = append(preciosFaltantes, precioTienda{sID, pID, 0})
			}
		}
	}
	return preciosFaltantes
}

func (s *Service) obtenerPreciosFaltantesDeAPI(pIDs []int, sIDs []string) error {
	preciosFaltantes := s.listaPreciosFaltantes(pIDs, sIDs)

	if len(preciosFaltantes) > 0 {
		// Inicializacion de channels
		preciosAObtener := make(chan precioTienda, len(preciosFaltantes))
		precios := make(chan precioTienda)
		errores := make(chan error)
		done := make(chan error)
		abort := false

		// Precarga de precios a obtener
		for i := range preciosFaltantes {
			preciosAObtener <- preciosFaltantes[i]
		}
		close(preciosAObtener)

		// Inicializacion de workers
		w := 0
		for w < numWorkers {
			go s.preciosWorker(preciosAObtener, precios, errores, &abort)
			w++
		}

		go s.preciosCollector(precios, errores, len(preciosFaltantes), done)
		err := <-done
		abort = true
		if err != nil {
			return err
		}
	}
	return nil
}

// CalcularPrecios calcula los precios para la lista de IDs de productos pIDs en
// los supermercados indicados en sIDs y devuelve una lista de Carrito
func (s *Service) CalcularPrecios(pIDs []int, sIDs []string) ([]Carrito, error) {
	err := s.obtenerPreciosFaltantesDeAPI(pIDs, sIDs)
	if err != nil {
		return nil, err
	}
	return s.calcularCarritos(pIDs, sIDs), nil
}
