package tp4

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type ProductoType struct {
	ID     int
	Precio int
}

// Productos es una lista de productos donde para cada producto
// se sabe el nombre del super mercado, el id y su precio.
// Esta estructura se puede cargar usando la funcion LeerProductos
// que carga informacion guardada en `productos.json`.
type Productos [][]string

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

type TiendaProducto struct {
	Tienda string
	ID     int
}

//
// {"tienda":"dia","id":1,"precio":7887}
func GetProductoTienda(idTienda string, idProducto int, requests chan<- Carrito) {
	//time.Sleep(time.Second * 2)
	url := fmt.Sprintf("https://productos-p6pdsjmljq-uc.a.run.app/%s/productos/%s", idTienda, fmt.Sprint(idProducto))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var j map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		panic(err)
	}
	precio := int(j["precio"].(float64))
	requests <- Carrito{Tienda: idTienda, Precio: precio}
}

func calcPreciosBulk(idsProductos []int, idTiendas []string) []Carrito {
	iterations := len(idTiendas) * len(idsProductos)
	requests := make(chan Carrito, iterations)
	for _, idTienda := range idTiendas {
		for _, idProducto := range idsProductos {
			// remove go and it wil take very long
			go GetProductoTienda(idTienda, idProducto, requests)
		}
	}
	carritoMap := make(map[string]int)
	for i := 0; i < iterations; i++ {
		carr := <-requests
		carritoMap[carr.Tienda] += carr.Precio
	}
	var carritoUltimo []Carrito
	for tienda, precio := range carritoMap {
		carritoUltimo = append(carritoUltimo, Carrito{Tienda: tienda, Precio: precio})
	}
	return carritoUltimo

}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(idsProductos []int, idTiendas []string) []Carrito {
	carritos := calcPreciosBulk(idsProductos, idTiendas)
	fmt.Println(carritos)
	repr, err := json.Marshal(carritos)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(repr))

	carritoMap := make(map[string]int)
	for _, carrito := range carritos {
		carritoMap[carrito.Tienda] = carrito.Precio
	}
	repr, err = json.Marshal(carritoMap)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(repr))
	return carritos
}

func CalcPreciosResponse(idsProductos []int, idTiendas []string) ([]byte, error) {
	carritos := calcPreciosBulk(idsProductos, idTiendas)
	return json.Marshal(carritos)
}
