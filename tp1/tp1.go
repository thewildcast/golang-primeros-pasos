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

const (
	TIENDA = iota
	IDPRODUCTO
	PRECIO
)

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	mapCarrito := make(map[string]int)
	var sliceCarrito []Carrito

	for _, prod := range p {
		for _, idProductoBuscado := range ids {
			if idProductoArray, err := strconv.Atoi(prod[IDPRODUCTO]); idProductoArray == idProductoBuscado && err == nil {
				precio, _ := strconv.Atoi(prod[PRECIO])
				mapCarrito[prod[TIENDA]] += precio
			}
		}
	}

	if len(mapCarrito) > 0 {
		for key, value := range mapCarrito {
			sliceCarrito = append(sliceCarrito, Carrito{Tienda: key, Precio: value})
		}
	}

	return sliceCarrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var total float64
	var ocurrencias int

	for _, productoArray := range p {
		if idProductoArray, err := strconv.Atoi(productoArray[IDPRODUCTO]); idProductoArray == idProducto && err == nil {
			precio, _ := strconv.ParseFloat(productoArray[PRECIO], 64)
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
	var masBarato item
	masBarato.id = idProducto
	var productoEncontrado bool

	for _, productoArray := range p {
		idProductoArray, errIDProducto := strconv.Atoi(productoArray[IDPRODUCTO])

		if idProductoArray == idProducto && errIDProducto == nil {
			precio, errPrecio := strconv.Atoi(productoArray[PRECIO])
			if errPrecio == nil {

				if precio < masBarato.Precio() || !productoEncontrado {
					masBarato.supermercado, masBarato.precio = productoArray[TIENDA], precio
				}

				productoEncontrado = true
			}
		}
	}

	return masBarato, productoEncontrado
}
