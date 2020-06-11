package tp1

import (
	"fmt"
	"strconv"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

// Productos es una lista de productos donde para cada producto
// se sabe el nombre del super mercado, el id y su precio.
// Esta estructura se puede cargar usando la funcion LeerProductos
// que carga informacion guardada en `productos.json`.
type Productos [][]string

const (
	SUPERMERCADO_INDEX = 0
	PRODUCTO_INDEX     = 1
	PRECIO_INDEX       = 2
)

type UnProducto struct {
	id     int
	precio int
}

func (p UnProducto) ID() int {
	return p.id
}

func (p UnProducto) Precio() int {
	return p.precio
}

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
	//return nil

	var carrito []Carrito
	mapaProductos := make(map[string]int)

	for _, prodInput := range ids {
		for _, prodLista := range p {
			productoActual, err := strconv.Atoi(prodLista[PRODUCTO_INDEX])
			if err != nil {
				fmt.Println("Error handling string conversion", productoActual)
				return nil
			}
			if prodInput == productoActual {
				productoActualPrecio, err := strconv.Atoi(prodLista[PRECIO_INDEX])
				if err != nil {
					fmt.Println("Error handling string conversion", productoActual)
					return nil
				}
				mapaProductos[prodLista[SUPERMERCADO_INDEX]] += productoActualPrecio
			}

		}
	}
	for supermercado, precio := range mapaProductos {
		carrito = append(carrito, Carrito{Tienda: supermercado, Precio: precio})
	}

	return carrito
}


// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var precioTotal float64
	var superCant int

	for _, prodLista := range p {
		productoActual, err := strconv.Atoi(prodLista[PRODUCTO_INDEX])
		if err != nil {
			fmt.Println("Error handling string conversion")
			return 0
		}
		if idProducto == productoActual {
			productoActualPrecio, err := strconv.Atoi(prodLista[PRECIO_INDEX])
			if err != nil {
				fmt.Println("Error handling product price in Productos")
				return 0
			}
			precioTotal += float64(productoActualPrecio)
			superCant++
		}
	}
	if superCant > 0 {
		return float64(precioTotal) / float64(superCant)
	}
	return 0.0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	//return nil, false
	menorPrecio := 0
	existe := false
	for _, prodLista := range p {
		productoActual, err := strconv.Atoi(prodLista[PRODUCTO_INDEX])
		precioActual, err := strconv.Atoi(prodLista[PRECIO_INDEX])
		if err != nil {
			fmt.Println("Error handling producto price in Productos")
			return nil, false
		}
		if idProducto == productoActual {
			if existe == false || precioActual < menorPrecio {
				menorPrecio = precioActual
				existe = true
			}
		}
	}
	return UnProducto{id: idProducto, precio: menorPrecio}, existe

}
