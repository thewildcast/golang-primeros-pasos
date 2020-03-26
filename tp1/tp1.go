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

type ProductoMasBarato struct {
	Id             int
	PrecioProducto int
}

func (p ProductoMasBarato) ID() int {
	return p.Id
}

func (p ProductoMasBarato) Precio() int {
	return p.PrecioProducto
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

	carritos := []Carrito{}

	listadoDeProductos, _ := LeerProductos("productos.json")

	for _, supermercado := range Supermercados {

		carrito := Carrito{Tienda: supermercado}

		precioTotal := 0

		for _, producto := range listadoDeProductos {

			for _, id := range ids {

				if supermercado == producto[0] && strconv.FormatInt(int64(id), 10) == producto[1] {

					precioProducto, _ := strconv.Atoi(producto[2])
					precioTotal += precioProducto
				}
			}
		}
		carrito.Precio = precioTotal

		if precioTotal > 0 {
			carritos = append(carritos, carrito)
		}
	}
	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	listadoDeProductos, _ := LeerProductos("productos.json")

	var precioPromedio float64

	for _, producto := range listadoDeProductos {

		if strconv.FormatInt(int64(idProducto), 10) == producto[1] {
			precio, _ := strconv.ParseFloat(producto[2], 64)
			precioPromedio += precio
		}

	}
	return precioPromedio / float64(len(Supermercados))
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	listadoDeProductos, _ := LeerProductos("productos.json")

	productoMasBarato := ProductoMasBarato{Id: idProducto, PrecioProducto: 0}

	for _, producto := range listadoDeProductos {

		if strconv.FormatInt(int64(idProducto), 10) == producto[1] {

			precio, _ := strconv.Atoi(producto[2])

			if productoMasBarato.Precio() == 0 || precio < productoMasBarato.Precio() {
				productoMasBarato.PrecioProducto = precio
			}
		}

	}
	return productoMasBarato, productoMasBarato.Precio() != 0
}
