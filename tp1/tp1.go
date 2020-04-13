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

type Alfajor struct {
	Id    int
	Super string
	Monto int
}

func (a Alfajor) ID() int {
	return a.Id
}

func (a Alfajor) Precio() int {
	return a.Monto
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	var precios []Carrito
	preciosPorSuper := make(map[string]int)
	for _, id := range ids {
		for _, product := range p {
			idProduct, err := strconv.Atoi(product[1])
			if err == nil && id == idProduct {
				precio, err := strconv.Atoi(product[2])
				if err == nil {
					preciosPorSuper[product[0]] += precio
				}
			}
		}
	}

	for super, precio := range preciosPorSuper {
		result := Carrito{Tienda: super, Precio: precio}
		precios = append(precios, result)
	}

	return precios
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	suma := 0.0
	count := 0.0

	for _, producto := range p {
		idProduct, err := strconv.Atoi(producto[1])
		if err == nil && idProduct == idProducto {
			precio, err := strconv.Atoi(producto[2])
			if err == nil {
				count++
				suma += float64(precio)
			}
		}
	}
	if count == 0.0 {
		return 0
	}

	return suma / count
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	producto := Alfajor{Id: idProducto}
	encontrado := false

	for _, prod := range p {
		idProduct, err := strconv.Atoi(prod[1])
		if err == nil && idProduct == idProducto {
			precio, err := strconv.Atoi(prod[2])
			if err == nil && (producto.Monto == 0 || precio < producto.Monto) {
				producto = Alfajor{Id: idProduct, Monto: precio}
				encontrado = true
			}
		}
	}
	return producto, encontrado
}
