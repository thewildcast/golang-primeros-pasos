package tp1

import (
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

type item struct {
	id, precio   int
	supermercado string
}

func (i item) ID() int {
	return i.id
}

func (i item) Precio() int {
	return i.precio
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
	mapCarrito := make(map[string]int)
	var sliceCarrito []Carrito

	for _, prod := range p {
		for _, buscado := range ids {
			if aux, _ := strconv.Atoi(prod[1]); aux == buscado {
				precio, _ := strconv.Atoi(prod[2])
				mapCarrito[prod[0]] += precio
			}
		}
	}

	if mapCarrito != nil {
		for key, value := range mapCarrito {
			sliceCarrito = append(sliceCarrito, Carrito{Tienda: key, Precio: value})
		}
	}

	return sliceCarrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	total := 0.00
	ocurrencias := 0

	for _, prod := range p {
		if aux, _ := strconv.Atoi(prod[1]); aux == idProducto {
			precio, _ := strconv.ParseFloat(prod[2], 64)
			total += precio
			ocurrencias++
		}
	}

	if ocurrencias == 0 {
		ocurrencias = 1
	}

	return (total / float64(ocurrencias))
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var minProducto item
	minProducto.id = idProducto
	existe := false

	for _, prod := range p {
		if idProd, _ := strconv.Atoi(prod[1]); idProd == idProducto {
			precio, _ := strconv.Atoi(prod[2])

			if precio < minProducto.Precio() || !existe {
				minProducto.supermercado, minProducto.precio = prod[0], precio
			}

			existe = true
		}
	}

	return minProducto, existe
}
