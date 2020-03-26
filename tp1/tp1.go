package tp1

import (
	"fmt"
	"os"
	"strconv"
)

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

const (
	TIENDA = iota
	PRODUCT_ID
	PRECIO
)

func stringtoint(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func stringtofloat(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func valuesofmap(carritos map[string]Carrito) []Carrito {
	values := make([]Carrito, 0, len(carritos))
	for _, v := range carritos {
		values = append(values, v)
	}
	return values
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	var carritos map[string]Carrito
	carritos = make(map[string]Carrito)

	uniqueIds := make(map[int]struct{})
	for _, id := range ids {
		uniqueIds[id] = struct{}{}
	}

	for _, v := range p {
		if _, ok := uniqueIds[stringtoint(v[PRODUCT_ID])]; ok {
			if _, ok := carritos[v[TIENDA]]; !ok {
				carritos[v[TIENDA]] = Carrito{v[TIENDA], 0}
			}
			carrito := carritos[v[TIENDA]]
			carrito.Precio += stringtoint(v[PRECIO])
			carritos[v[TIENDA]] = carrito
		}
	}
	return valuesofmap(carritos)
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	sumaPrecios := 0
	cantidadProductos := 0.0
	for _, v := range p {
		if stringtoint(v[PRODUCT_ID]) == idProducto {
			sumaPrecios += stringtoint(v[PRECIO])
			cantidadProductos++
		}
	}
	if cantidadProductos > 0 {
		return float64(sumaPrecios) / cantidadProductos
	}
	return 0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var productoMasBarato = producto{idProducto, 0}
	for _, v := range p {
		if stringtoint(v[PRODUCT_ID]) == idProducto && productoMasBarato.Precio() == 0 {
			productoMasBarato = producto{stringtoint(v[PRODUCT_ID]), stringtoint(v[PRECIO])}
		} else if stringtoint(v[PRODUCT_ID]) == idProducto && productoMasBarato.Precio() > stringtoint(v[PRECIO]) {
			productoMasBarato = producto{stringtoint(v[PRODUCT_ID]), stringtoint(v[PRECIO])}
		}
	}
	if productoMasBarato.Precio() == 0 {
		return productoMasBarato, false
	}
	return productoMasBarato, true
}
