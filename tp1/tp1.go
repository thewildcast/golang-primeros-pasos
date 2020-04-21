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

// Carrito contiene el nombre de la tienda y el precio final luego
// de sumar todos los productos.
type Carrito struct {
	Tienda string
	Precio int
}

type ProductoStruct struct {
	id     int
	precio int
}

func (p ProductoStruct) ID() int {
	return p.id
}

func (p ProductoStruct) Precio() int {
	return p.precio
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {

	var carritos []Carrito
	auxMap := make(map[string]int)

	for _, v := range p {
		if idProductLista, err := strconv.Atoi(v[1]); err == nil {
			for _, idProd := range ids {
				if idProductLista == idProd {
					if precioProducto, err := strconv.Atoi(v[2]); err == nil {
						auxMap[v[0]] += precioProducto
					}
				}
			}
		}
	}
	for index, value := range auxMap {
		carritos = append(carritos, Carrito{Tienda: index, Precio: value})
	}
	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var count float64 = 0
	var acumulador float64
	for _, v := range p {
		idProductLista, _ := strconv.Atoi(v[1])
		if idProducto == idProductLista {
			precioProducto, _ := strconv.ParseFloat(v[2], 64)
			acumulador = acumulador + precioProducto
			count++
		}
	}
	if count == 0 {
		return acumulador
	}
	return acumulador / count
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var found bool = false
	var producto ProductoStruct = ProductoStruct{id: idProducto, precio: 0}
	for _, v := range p {
		idProductLista, _ := strconv.Atoi(v[1])
		if idProducto == idProductLista {
			precioProducto, _ := strconv.Atoi(v[2])
			if !found || precioProducto < producto.precio {
				producto.precio = precioProducto
			}
			found = true
		}
	}
	return producto, found
}
