package tp4

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	URL_PRODUCTOS = "https://productos-p6pdsjmljq-uc.a.run.app/%s/productos/%d"
)

type Producto struct {
	Tienda string `json:"tienda"`
	Id     int32  `json:"id"`
	Precio int32  `json:"precio"`
}

type Carrito struct {
	Tienda string `json:"tienda"`
	Total  int32  `json:"total"`
}

func CalcularPorTiendas(ids []int32, tiendas []string) []Carrito {

	tiendasChan := make(chan Carrito, len(tiendas)+1)

	for _, tienda := range tiendas {
		go CalcularPorTienda(ids, tienda, tiendasChan)
	}

	var carritos []Carrito

	for carrito := range tiendasChan {
		carritos = append(carritos, carrito)

		if len(carritos) == len(tiendas) {
			close(tiendasChan)
		}
	}

	return carritos
}

func CalcularPorTienda(ids []int32, tienda string, tiendasChan chan Carrito) {

	productos := make(chan Producto, len(ids)+1)

	for _, id := range ids {
		go ObtenerProducto(id, tienda, productos)
	}

	carritoTienda := Carrito{Tienda: tienda}

	count := 0
	for prod := range productos {
		carritoTienda.Total += prod.Precio
		if count++; count == len(ids) {
			close(productos)
		}
	}

	tiendasChan <- carritoTienda
}

func ObtenerProducto(id int32, tienda string, productos chan Producto) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(URL_PRODUCTOS, tienda, id), nil)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("Error al obtener el producto con id: %d, el error fue: %s", id, err.Error())
		productos <- Producto{Tienda:tienda, Precio:0}
		return
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("No se encontro el producto con id: %d, statusCode: %d", id, response.StatusCode)
		productos <- Producto{Tienda:tienda, Precio:0}
		return
	}

	defer response.Body.Close()

	prod := &Producto{}
	err = json.NewDecoder(response.Body).Decode(prod)

	if err != nil {
		log.Printf("No se pudo decodificar el producto con id: %d, el error fue: %s", id, err.Error())
		productos <- Producto{Tienda:tienda, Precio:0}
		return
	}

	productos <- *prod
}
