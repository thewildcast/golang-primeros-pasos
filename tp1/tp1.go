package tp1

import (
	"strconv"
)

const (
	POSICION_TIENDA = 0
	POSICION_ID     = 1
	POSICION_PRECIO = 2
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

// Retorna el Id del producto
func (p producto) ID() int {
	return p.id
}

// Retorna el Precio del producto
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

	mapaCarritos := map[string]int{}

	for _, id := range ids {
		for _, producto := range p {

			if producto[POSICION_ID] == strconv.Itoa(id) {
				mapaCarritos[producto[POSICION_TIENDA]] =
					mapaCarritos[producto[POSICION_TIENDA]] + parseInt(producto[POSICION_PRECIO])
			}
		}
	}

	var carritos []Carrito

	for tienda, precio := range mapaCarritos {
		carritos = append(carritos, Carrito{Tienda: tienda, Precio: precio})
	}

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	tiendas := p.buscarPorId(idProducto)

	cantidad := float64(len(tiendas))
	sumaTotal := 0.0
	var promedio float64

	if cantidad > 0 {
		for _, precioStr := range tiendas {
			sumaTotal = sumaTotal + parseFloat(precioStr)
		}

		promedio = sumaTotal / cantidad
	}

	return promedio
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	tiendas := p.buscarPorId(idProducto)

	productoMasBarato := producto{id: idProducto}
	var encontrado bool

	if tiendas != nil && len(tiendas) > 0 {
		for _, precioStr := range tiendas {

			precio, _ := strconv.Atoi(precioStr)

			if !encontrado || precio < productoMasBarato.Precio() {
				productoMasBarato = producto{id: idProducto, precio: precio}
				encontrado = true
			}
		}
	}

	return productoMasBarato, encontrado
}

// Retorna la conversion de un string a float64
func parseFloat(value string) float64 {
	precio, _ := strconv.ParseFloat(value, 64)

	return precio
}

// Retorna la conversion de un string a int
func parseInt(value string) int {
	precio, _ := strconv.Atoi(value)

	return precio
}

// buscarPorId busca entre los productos todos los que tengan el id
// crea un mapa que guarda la tienda y el precio
// Retorna un mapa con todas lstiendas y el precio por cada una
func (p Productos) buscarPorId(idProducto int) map[string]string {
	tiendas := map[string]string{}

	for _, producto := range p {
		if producto[POSICION_ID] == strconv.Itoa(idProducto) {
			tiendas[producto[POSICION_TIENDA]] = producto[POSICION_PRECIO]
		}
	}

	return tiendas
}
