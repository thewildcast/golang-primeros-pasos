package tp1

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	IDX_TIENDA = iota
	IDX_PRODUCTO
	IDX_PRECIO
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type producto struct {
	id, precio   int
}

func (p producto) ID() int {
	return p.id
}

func (p producto) Precio() int {
	return p.precio
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

func existeTienda(nombre string, carritos []Carrito) int {
	for indice, carrito := range carritos {
		if carrito.Tienda == nombre {
			return indice
		}
	}
	return -1
}

func castToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return result
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	carritos := []Carrito{}
	mapaCarritos := make(map[string]int)

	for _, producto := range p {
		idProducto := castToInt(producto[IDX_PRODUCTO])
		if exists(idProducto, ids) {
			if mapaCarritos[producto[IDX_TIENDA]] == 0 {
				mapaCarritos[producto[IDX_TIENDA]] = castToInt(producto[IDX_PRECIO])
			} else {
				mapaCarritos[producto[IDX_TIENDA]] += castToInt(producto[IDX_PRECIO])
			}
		}
	}
	// No se si existe una mejor forma de hacer el cast de map a array
	for tienda, total := range mapaCarritos {
		carritos = append(carritos, Carrito{tienda, total})
	}
	return carritos
}

func exists(idProduct int, ids []int) bool {
	for _, id := range ids {
		if id == idProduct {
			return true
		}
	}
	return false
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var precios, total float64
	for _, producto := range p {
		if castToInt(producto[IDX_PRODUCTO]) == idProducto {
			precios += float64(castToInt(producto[IDX_PRECIO]))
			total++
		}
	}
	if total > 0 {
		total = precios/total
	}
	return total
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	productoMasBarato := producto{idProducto, math.MaxInt64}
	masBarato := false
	for _, producto := range p {
		if castToInt(producto[1]) == idProducto && productoMasBarato.precio > castToInt(producto[2]){
			productoMasBarato.precio = castToInt(producto[2])
			masBarato = true
		}
	}
	if !masBarato {
		return producto{idProducto, 0}, false
	}
	return productoMasBarato, masBarato
}
