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

type producto struct {
	id     string
	precio string
}

func (p producto) ID() int {
	i, _ := strconv.Atoi(p.id)
	return i
}

func (p producto) Precio() int {
	i, _ := strconv.Atoi(p.precio)
	return i
}

func idInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	// CalcularPrecios es un metodo que le puedo aplica a Productos.
	var lista []Carrito
	productos, _ := LeerProductos("productos.json")
	// indice, valor  (en un for - range). No me interesa en QUE indice estoy parado
	for _, super := range Supermercados {
		// ["Jumbo","0","4804"] ejemplo de un producto
		sumaTotal := 0
		for _, product := range productos {
			if product[0] == super {
				i, _ := strconv.Atoi(product[1])
				if idInSlice(i, ids) {
					price, _ := strconv.Atoi(product[2])
					sumaTotal += price
				}
			}
		}
		if sumaTotal > 0 {
			lista = append(lista, Carrito{super, sumaTotal})
		}
	}
	return lista
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var sumaTotal int
	var cantSuper int
	productos, _ := LeerProductos("productos.json")
	for _, product := range productos {
		id, _ := strconv.Atoi(product[1])
		if id == idProducto {
			precio, _ := strconv.Atoi(product[2])
			sumaTotal += precio
			cantSuper++
		}
	}
	if cantSuper == 0 {
		return 0.0
	}
	return float64(sumaTotal) / float64(cantSuper)
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	stringID := strconv.Itoa(idProducto)
	barato := producto{stringID, "0"}
	productos, _ := LeerProductos("productos.json")
	for _, product := range productos {
		id, _ := strconv.Atoi(product[1])
		precio, _ := strconv.Atoi(product[2])

		if id == idProducto {
			if barato.Precio() > precio || barato.Precio() == 0 {
				barato = producto{product[1], product[2]}
			}
		}
	}
	found := barato.Precio() > 0
	return barato, found
}
