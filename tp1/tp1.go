package tp1

import (
	"strconv"
)

const INDEX_TIENDA = 0
const INDEX_ID_PRODUCTO = 1
const INDEX_PRECIO = 2

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

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

type producto struct {
	id     int
	precio int
	tinda  string
}

func (p producto) ID() int {
	return p.id
}

func (p producto) Precio() int {
	return p.precio
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	mapCarritos := map[string]int{}

	for _, id := range ids {
		for _, producto := range p {
			if producto[INDEX_ID_PRODUCTO] == strconv.Itoa(id) {
				value, err := strconv.Atoi(producto[INDEX_PRECIO])
				if err == nil {
					mapCarritos[producto[INDEX_TIENDA]] += value
				}
			}
		}
	}

	carritos := []Carrito{}

	for super, precio := range mapCarritos {
		carritos = append(carritos, Carrito{Tienda: super, Precio: precio})
	}

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	precioTotal := 0
	cantidadSuper := 0

	for _, producto := range p {
		if producto[INDEX_ID_PRODUCTO] == strconv.Itoa(idProducto) {
			value, err := strconv.Atoi(producto[INDEX_PRECIO])
			if err == nil {
				precioTotal += value
				cantidadSuper++
			}
		}
	}

	if cantidadSuper > 0 {
		return float64(precioTotal) / float64(cantidadSuper)
	}
	return 0.0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var precioMasBarato int
	encontroProducto := false

	for _, producto := range p {
		if producto[INDEX_ID_PRODUCTO] == strconv.Itoa(idProducto) {
			value, err := strconv.Atoi(producto[INDEX_PRECIO])
			if err == nil && (!encontroProducto || value < precioMasBarato) {
				encontroProducto = true
				precioMasBarato = value
			}
		}
	}

	return producto{id: idProducto, precio: precioMasBarato}, encontroProducto
}
