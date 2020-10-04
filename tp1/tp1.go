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

// MyProducto struct que implementa Producto
type MyProducto struct {
	id     int
	precio int
}

// ID devuelve id
func (p MyProducto) ID() int {
	return p.id
}

// Precio devuelve precio
func (p MyProducto) Precio() int {
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
	var carrito []Carrito
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return nil
	}

	for _, producto := range productos {
		var supermercado string
		var idProducto int
		var precio int

		supermercado = producto[0]
		idProducto, _ = strconv.Atoi(producto[1])
		precio, _ = strconv.Atoi(producto[2])

		if intInSlice(idProducto, ids) {
			item := Carrito{Tienda: supermercado, Precio: precio}
			index := findCarritoInSlice(&item, &carrito)
			if index != -1 {
				carrito[index].Precio += item.Precio
			} else {
				carrito = append(carrito, item)
			}
		}
	}

	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var promedio float64
	var sumatoriaPrecios float64
	var qty float64

	productos, err := LeerProductos("productos.json")
	if err != nil {
		return 0
	}

	sumatoriaPrecios = 0
	qty = 0
	promedio = 0

	for _, producto := range productos {
		var _idProducto int
		var _precio float64

		_idProducto, _ = strconv.Atoi(producto[1])
		_precio, _ = strconv.ParseFloat(producto[2], 10)

		if _idProducto == idProducto {
			qty++
			sumatoriaPrecios += _precio
		}
	}

	if qty > 0 {
		promedio = sumatoriaPrecios / qty
	}

	return promedio
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var encontro bool
	var producto MyProducto
	var menorPrecio int

	productos, err := LeerProductos("productos.json")
	if err != nil {
		return nil, false
	}

	encontro = false
	menorPrecio = int(^uint(0) >> 1)

	for _, producto := range productos {
		var _idProducto int
		var _precio int

		_idProducto, _ = strconv.Atoi(producto[1])
		_precio, _ = strconv.Atoi(producto[2])

		if _idProducto == idProducto {
			encontro = true
			if _precio < menorPrecio {
				menorPrecio = _precio
			}
		}
	}

	if !encontro {
		menorPrecio = 0
	}

	producto = MyProducto{id: idProducto, precio: menorPrecio}

	return producto, encontro
}
