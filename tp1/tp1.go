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

type producto struct {
	id, precio int
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
	precioXSuper := map[string]int{}
	prods := make(map[string]bool)
	for _, id := range ids {
		prods[strconv.Itoa(id)] = true
	}
	for _, prod := range p {
		if _, present := prods[prod[1]]; present {
			if precio, err := strconv.Atoi(prod[2]); err == nil {
				precioXSuper[prod[0]] += precio
			}
		}
	}
	var res []Carrito
	for k, v := range precioXSuper {
		res = append(res, Carrito{k, v})
	}
	return res
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	suma := 0
	cantSuper := 0
	for _, prod := range p {
		if id, err := strconv.Atoi(prod[1]); err == nil && id == idProducto {
			if precio, err := strconv.Atoi(prod[2]); err == nil {
				cantSuper++
				suma += precio
			}
		}
	}
	if cantSuper == 0 {
		return 0.0
	}
	return float64(suma) / float64(cantSuper)
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	masBarato := producto{idProducto, 0}
	for _, prod := range p {
		if id, err := strconv.Atoi(prod[1]); err == nil && id == idProducto {
			if precio, err := strconv.Atoi(prod[2]); err == nil &&
				(masBarato.precio == 0 || masBarato.precio > precio) {
				masBarato.precio = precio
			}
		}
	}
	return masBarato, masBarato.precio > 0
}
