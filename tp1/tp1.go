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

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	totales := map[string]int{}

	for _, id := range ids {
		for _, prod := range p {
			super := prod[0]
			prodID, err := strconv.Atoi(prod[1])
			if err != nil {
				continue
			}

			if id == prodID {
				precio, err := strconv.Atoi(prod[2])
				if err != nil {
					continue
				}
				totales[super] += precio
			}
		}
	}

	carritos := []Carrito{}

	for super, total := range totales {
		carrito := Carrito{
			Tienda: super,
			Precio: total,
		}
		carritos = append(carritos, carrito)
	}

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	carritos := p.CalcularPrecios(idProducto)
	if len(carritos) == 0 {
		return 0
	}

	promedio := 0

	for _, carrito := range carritos {
		promedio += carrito.Precio
	}

	return float64(promedio) / float64(len(carritos))
}

type Item struct {
	Codigo int
	Costo  int
}

func (p Item) ID() int {
	return p.Codigo
}

func (p Item) Precio() int {
	return p.Costo
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	precioMasBarato := 0
	existe := false

	for _, prod := range p {
		prodID, err := strconv.Atoi(prod[1])
		if err != nil {
			continue
		}

		if idProducto == prodID {
			precio, err := strconv.Atoi(prod[2])
			if err != nil {
				continue
			}

			if precio < precioMasBarato || precioMasBarato == 0 {
				precioMasBarato = precio
			}

			existe = true
		}
	}

	productoMasBarato := Item{
		Codigo: idProducto,
		Costo:  precioMasBarato,
	}

	return productoMasBarato, existe
}
