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

//Tipo para utilizar []string como interfaz Producto
type productoSlice []string

func (s productoSlice) ID() int {
	id, err := strconv.Atoi(s[1])
	if err != nil {
		panic(err)
	}
	return id
}

func (s productoSlice) Precio() int {
	pc, err := strconv.Atoi(s[2])
	if err != nil {
		fmt.Println("Error accediendo a precio!")
		panic(err)
	}
	return pc
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	//map para acumular los precios por supermercado
	m := make(map[string]int)
	c := []Carrito{}
	//Se iteran los productos
	for _, s := range p {
		//Por cada producto, si coincide con un id de ids, se acumula en map
		for _, pID := range ids {
			if productoSlice(s).ID() == pID {
				m[s[0]] = m[s[0]] + productoSlice(s).Precio()

			}
		}
	}
	//Se convierte map en slice de Carrito
	for k, v := range m {
		c = append(c, Carrito{k, v})
	}
	return c
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var t, c int = 0, 0
	for _, s := range p {
		if productoSlice(s).ID() == idProducto {
			t += productoSlice(s).Precio()
			c++
		}
	}
	if c > 0 {
		return float64(t) / float64(c)
	}
	return 0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var pb []string
	for _, s := range p {
		if productoSlice(s).ID() == idProducto {
			if pb == nil || productoSlice(s).Precio() < productoSlice(pb).Precio() {
				pb = s
			}
		}
	}
	if pb != nil {
		return productoSlice(pb), true
	}
	return productoSlice([]string{"", "101", "0"}), false
}
