package tp1

import (
	"strconv"
)

const (
	SUPER_INDEX = iota
	ID_INDEX
	PRECIO_INDEX
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type producto struct {
	tienda string
	id     int
	precio int
}

func (p producto) Tienda() string {
	return p.tienda
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
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	carritos := make(map[string]int)
	for _, prod := range p {
		p_id, _ := strconv.Atoi(prod[ID_INDEX])
		for _, id := range ids {
			if id == p_id {
				precio, _ := strconv.Atoi(prod[PRECIO_INDEX])
				carritos[prod[SUPER_INDEX]] += precio
			}
		}
	}
	var carritos_return []Carrito

	for k, v := range carritos {
		new_carrito := Carrito{k, v}
		carritos_return = append(carritos_return, new_carrito)
	}
	return carritos_return
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var sum_precio int
	var cant int
	for _, prod := range p {
		p_id, _ := strconv.Atoi(prod[ID_INDEX])
		if idProducto == p_id {
			precio, _ := strconv.Atoi(prod[PRECIO_INDEX])
			sum_precio += precio
			cant++
		}
	}
	var precio_prom float64
	if cant > 0 {
		precio_prom = float64(sum_precio) / float64(cant)
	}
	return precio_prom
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var min_price int
	var existe bool
	var tienda string
	for _, prods := range p {
		p_id, _ := strconv.Atoi(prods[ID_INDEX])
		if idProducto == p_id {
			existe = true
			precio, _ := strconv.Atoi(prods[PRECIO_INDEX])
			if precio < min_price || min_price == 0 {
				min_price = precio
				tienda = prods[SUPER_INDEX]
			}
		}
	}
	pr := producto{tienda, idProducto, min_price}
	return pr, existe
}
