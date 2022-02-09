package tp4

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
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

// ProductoTienda contiene el nombre de la tienda y el id del producto que quiero consultar
type ProductoTienda struct {
	ID     string
	Tienda string
}

//
// {"tienda":"dia","id":1,"precio":7887}
func GetProductoTienda(idTienda string, idProducto int, carritoMap map[string]int, wg *sync.WaitGroup, someMapMutex *sync.RWMutex) {
	defer wg.Done()
	time.Sleep(time.Second * 3)
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
	someMapMutex.Lock()
	carritoMap[idTienda] += int(j["precio"].(float64))
	someMapMutex.Unlock()
	//carrito := Carrito{Tienda: idTienda, Precio: int(j["precio"].(float64))}
	//*carritos = append(*carritos, carrito)
}

func calcPreciosBulk(idsProductos []int, idTiendas []string) []Carrito {
	var (
		carritoMap   = map[string]int{}
		someMapMutex = sync.RWMutex{}
	)
	carritoMap = make(map[string]int)
	wg := &sync.WaitGroup{}
	for _, idTienda := range idTiendas {
		for _, idProducto := range idsProductos {
			wg.Add(1)
			// remove go and it wil take very long
			go GetProductoTienda(idTienda, idProducto, carritoMap, wg, &someMapMutex)
		}
	}
	wg.Wait()
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
	return calcPreciosBulk(idsProductos, idTiendas)
}
