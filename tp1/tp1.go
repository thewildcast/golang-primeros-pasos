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

type unProducto struct {
	id     int
	precio int
}

func (p unProducto) ID() int {
	return p.id
}

func (p unProducto) Precio() int {
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
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return nil
	}
	mCarrito := map[string]int{}

	for _, valor := range productos {
		v, found := mCarrito[valor[0]]
		idProd, _ := strconv.Atoi(valor[1])
		precio, _ := strconv.Atoi(valor[2])
		if existe(idProd, ids) {
			if found {
				mCarrito[valor[0]] = v + precio
			} else {
				mCarrito[valor[0]] = precio
			}
		}

	}

	listCarrito := []Carrito{}
	for key, valor := range mCarrito {
		unCarrito := Carrito{Tienda: key, Precio: valor}
		listCarrito = append(listCarrito, unCarrito)
	}

	return listCarrito
}

func existe(valorBuscado int, ids []int) bool {
	for _, valor := range ids {
		if valorBuscado == valor {
			return true
		}
	}
	return false
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return 0
	}
	suma := 0.0
	cantSuper := 0.0
	for _, valor := range productos {
		idProd, _ := strconv.Atoi(valor[1])
		if idProd == idProducto {
			precio, _ := strconv.ParseFloat(valor[2], 64)
			suma += precio
			cantSuper++
		}
	}
	if cantSuper > 0 {
		return float64(suma / cantSuper)
	}

	return 0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return nil, false
	}
	var masBarato int
	loEncontro := false

	for _, valor := range productos {
		if idProd, _ := strconv.Atoi(valor[1]); idProd == idProducto {
			precio, _ := strconv.Atoi(valor[2])
			if (!loEncontro) {
				masBarato = precio
			} else if precio < masBarato {
				masBarato = precio
			}
			loEncontro = true
		}
	}
	if loEncontro {
		return unProducto{id: idProducto, precio: masBarato}, true
	} else {
		return unProducto{id: idProducto, precio: 0}, false
	}

}
