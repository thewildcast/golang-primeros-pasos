package tp1

import (
	"fmt"
	"strconv"
	"math"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type ProductoType struct {
	id int
	precio int
}

func (p ProductoType) ID() int {
	return p.id
}

func (p ProductoType) Precio() int {
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
	var mapResult = make(map[string]int)
	var result []Carrito
	for _, id := range ids {
		for _, producto := range p {
			if producto[1] == fmt.Sprintf("%d", id) {
				price, err := strconv.Atoi(producto[2])
				if err != nil {
					panic(err)
				}
				mapResult[producto[0]] += price
			}
		}
	}
	for key, value := range mapResult {
		result = append(result, Carrito{key, value})
	}
	return result
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var sum , quantity int
	for _, producto := range p {
		if producto[1] == fmt.Sprintf("%d", idProducto) {
			price, err := strconv.Atoi(producto[2])
			if err != nil {
				panic(err)
			}
			sum += price
			quantity++
		}
	}
	if quantity == 0 {
		return 0
	}
	return float64(sum) / float64(quantity)
}

func stringToProductoType(producto []string) ProductoType {
	id, err := strconv.Atoi(producto[1])
	if err != nil {
		panic(err)
	}
	precio, err := strconv.Atoi(producto[2])
	if err != nil {
		panic(err)
	}
	return ProductoType{
		id: id,
		precio: precio,
	}
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var productoMasBarato ProductoType = ProductoType{
		id: idProducto,
	}
	var minPrice int = math.MaxInt32
	var found bool
	for _, producto := range p {
		newProductoType := stringToProductoType(producto)
		if newProductoType.id == idProducto && newProductoType.precio < minPrice {
			productoMasBarato = ProductoType{
				id: newProductoType.id,
				precio: newProductoType.precio,
			}
			minPrice = newProductoType.precio
			found = true
		}
	}
	return productoMasBarato, found
} 
