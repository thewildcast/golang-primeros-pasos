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

type producto struct {
	id     int
	precio int
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

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	carrito := map[string]Carrito{}
	for _, param1 := range ids {
		for _, rawProducto := range p {
			super := rawProducto[0]
			id, _ := strconv.Atoi(rawProducto[1])
			precio, _ := strconv.Atoi(rawProducto[2])
			if id == param1 {
				itemCarrito, found := carrito[super]
				if found {
					itemCarrito.Precio = itemCarrito.Precio + precio
					carrito[super] = itemCarrito
				} else {
					newCarrito := Carrito{Tienda: super, Precio: precio}
					carrito[super] = newCarrito
				}

			}

		}
	}

	result := []Carrito{}
	for _, value := range carrito {
		result = append(result, value)
	}
	return result
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	carritosProducto := p.CalcularPrecios(idProducto)
	var sumPrecio int

	for _, carrito := range carritosProducto {

		sumPrecio = sumPrecio + carrito.Precio

	}

	numCarritos := len(carritosProducto)

	var promedio float64

	if numCarritos > 0 {
		promedio = float64(sumPrecio) / float64(numCarritos)
	}
	return promedio
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	var minPrice int
	var product producto

	carritosProducto := p.CalcularPrecios(idProducto)

	if len(carritosProducto) == 0 {
		product := producto{id: idProducto, precio: 0}
		return product, false
	}

	for _, carrito := range carritosProducto {
		if carrito.Precio < minPrice || minPrice == 0 {
			minPrice = carrito.Precio
			product = producto{id: idProducto, precio: carrito.Precio}
		}

	}
	return product, true

}
