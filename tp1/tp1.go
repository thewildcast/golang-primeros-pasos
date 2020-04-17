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

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

type producto struct {
	id, precio int
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

	var carrito = []Carrito{}

	// No hay ningun producto
	if len(ids) == 0 {
		return nil
	}

	pm := map[string]int{}

	for _, producto := range p {
		for _, id := range ids {
			if pID := strToInt(producto[1]); pID == id {
				pm[producto[0]] += strToInt(producto[2])
				break
			}
		}
	}

	for t, p := range pm {
		carrito = append(carrito, Carrito{Tienda: t, Precio: p})
	}

	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	var totalPrecio, cantProducto float64

	for _, producto := range p {
		if pID := strToInt(producto[1]); pID != idProducto {
			continue
		}
		pPrecio := strToFloat(producto[2], 64)
		totalPrecio += pPrecio
		cantProducto++

	}

	// El producto ingresado no existe
	if totalPrecio == 0 {
		return totalPrecio
	}

	return totalPrecio / cantProducto
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	prodBarato := producto{idProducto, 0}
	esPrimeraIteracion := true

	for _, prod := range p {
		if pID := strToInt(prod[1]); pID != idProducto {
			continue
		}
		if pPrecio := strToInt(prod[2]); esPrimeraIteracion || pPrecio < prodBarato.precio {
			if esPrimeraIteracion {
				esPrimeraIteracion = false
			}
			prodBarato.precio = pPrecio
		}
	}

	return prodBarato, prodBarato.precio > 0
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Errorf("error converting %s to int: %s", s, err)
	}

	return i
}

func strToFloat(s string, bitSize int) float64 {
	f, err := strconv.ParseFloat(s, bitSize)

	if err != nil {
		fmt.Errorf("error converting %s to float%d: %s", s, bitSize, err)
	}

	return f
}
