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
	tienda string
	id     int
	precio int
}

func (p *producto) Tienda() string {
	return p.tienda
}

func (p *producto) ID() int {
	return p.id
}

func (p *producto) Precio() int {
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

func newProducto(prod []string) producto {
	pTienda := prod[0]
	pID, _ := strconv.Atoi(prod[1])
	pPrecio, _ := strconv.Atoi(prod[2])
	return producto{pTienda, pID, pPrecio}
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	idMap := make(map[int][]producto)
	for i := 0; i < len(p); i++ {
		prod := newProducto(p[i])
		idMap[prod.ID()] = append(idMap[prod.ID()], prod)
	}
	tiendasMap := make(map[string]int)
	for i := 0; i < len(ids); i++ {
		prodByTienda := idMap[ids[i]]
		for j := 0; j < len(prodByTienda); j++ {
			tiendasMap[prodByTienda[j].Tienda()] = tiendasMap[prodByTienda[j].Tienda()] + prodByTienda[j].Precio()
		}
	}
	var carrito []Carrito
	for key, value := range tiendasMap {
		carrito = append(carrito, Carrito{key, value})
	}
	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	return 0
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	return nil, false
}
