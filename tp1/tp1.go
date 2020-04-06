package tp1

import (
	"sort"
	"strconv"
)

type producto struct {
	supermercado string
	id           int
	precio       int
}

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

func (p producto) ID() int {
	return p.id
}

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
func (p Productos) CalcularPrecios(ids []int) []Carrito {

	chango := map[string]int{}

	var id string

	for i := 0; i < len(p); i++ {

		for _, valor := range ids {

			id = strconv.Itoa(valor)

			if id == p[i][1] {

				cadenaPrecio, _ := strconv.Atoi(p[i][2])

				chango[p[i][0]] = cadenaPrecio

			}
		}

	}

	dimension := len(chango)

	carrito := make([]Carrito, 0, dimension)

	for clave, valor := range chango {

		carr := Carrito{clave, valor}

		carrito = append(carrito, carr)

	}

	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	var prom float64
	var suma int
	

	for i := 0; i < len(p); i++ {


		if strconv.Itoa(idProducto) == p[i][1] {

			cadenaPrecio, _ := strconv.Atoi(p[i][2])

			suma += cadenaPrecio
		}

	}

	prom = float64(suma) / float64(len(p))

	return prom
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	produc := map[int]string{}

	flag := false

	var productoBarato producto

	for i := 0; i < len(p); i++ {

		if strconv.Itoa(idProducto) == p[i][1] {

			produc[i] = p[i][2]

		}

	}

	var precios []string
	for _, valor := range produc {

		precios = append(precios, valor)

	}
	sort.Strings(precios)

	for key, valor := range p {

		if valor[1] == strconv.Itoa(idProducto) && valor[2] == precios[0] {

			productoBarato.id, _ = strconv.Atoi(p[key][1])
			productoBarato.precio, _ = strconv.Atoi(p[key][2])

			flag = true

		}
	}
	return productoBarato, flag
}
