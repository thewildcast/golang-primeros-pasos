package tp1

import (
	"fmt"
	"os"
	"strconv"
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

	for _, producto := range p {
		idProducto := castToInt(producto[1])
		if exists(idProducto, ids) {
			existe := existeTienda(producto[0], carritos)
			if existe < 0  {
				carritos = append(carritos, Carrito{producto[0], castToInt(producto[2])})
			} else {
				carritos[existe].Precio += castToInt(producto[2])
			}
		}
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
	precios, total := 0.0, 0.0
	for _, producto := range p {
		if castToInt(producto[1]) == idProducto {
			precios += float64(castToInt(producto[2]))
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
	productoMasBarato := producto{idProducto, 0}
	masBarato := false
	for indice, producto := range p {
		if indice == 0 {
			productoMasBarato.precio = castToInt(producto[2])
			masBarato = false
		} else {
			if castToInt(producto[1]) == idProducto && productoMasBarato.precio > castToInt(producto[2]){
				productoMasBarato.precio = castToInt(producto[2])
				masBarato = true
			}
		}
	}
	if !masBarato {
		return producto{idProducto, 0}, false
	}
	return productoMasBarato, masBarato
}
