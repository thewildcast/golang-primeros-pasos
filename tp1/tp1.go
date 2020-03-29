package tp1

import (
	"fmt"
	"strconv"
)

const (
	SUPER_INDEX  = 0
	ID_INDEX     = 1
	PRECIO_INDEX = 2
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type producto struct {
	Id    int
	Price int
}

func (p producto) ID() int {
	return p.Id
}

func (p producto) Precio() int {
	return p.Price
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

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	// check that my list is not empty
	if len(ids) == 0 {
		return nil
	}

	var carritos []Carrito
	var supermercadoByPrice = make(map[string]int)

	for _, id := range ids {
		for _, productos := range p {
			productoId, err := strconv.Atoi(productos[ID_INDEX])
			if err != nil {
				fmt.Println("not a number")
			}

			if productoId == id {
				precioDelproducto, err := strconv.Atoi(productos[PRECIO_INDEX])
				if err != nil {
					fmt.Println("not a number")
				}
				supermercadoByPrice[productos[SUPER_INDEX]] += precioDelproducto
			}
		}
	}

	for key, value := range supermercadoByPrice {
		carritos = append(carritos, Carrito{Tienda: key, Precio: value})
	}

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	var cantidadProductos int
	var sumaDeProductos int

	for _, productos := range p {
		productoId, err := strconv.Atoi(productos[ID_INDEX])
		if err != nil {
			fmt.Println("not a number")
		}
		precioDelproducto, err := strconv.Atoi(productos[PRECIO_INDEX])
		if err != nil {
			fmt.Println("not a number")
		}

		if productoId == idProducto {
			sumaDeProductos += precioDelproducto
			cantidadProductos++
		}
	}

	if cantidadProductos == 0 {
		return 0.0
	}

	return float64(sumaDeProductos) / float64(cantidadProductos)

}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	//flags
	var masBarato = false
	var precioMasBarato producto = producto{Id: idProducto, Price: 0}

	for _, productos := range p {
		productoId, err := strconv.Atoi(productos[ID_INDEX])
		if err != nil {
			fmt.Println("not a number")
		}
		precioDelproducto, err := strconv.Atoi(productos[PRECIO_INDEX])
		if err != nil {
			fmt.Println("not a number")
		}

		if productoId == idProducto && (precioDelproducto <= precioMasBarato.Precio() || !masBarato) {
			masBarato = true
			precioMasBarato = producto{Id: productoId, Price: precioDelproducto}
		}
	}

	return precioMasBarato, masBarato

}
