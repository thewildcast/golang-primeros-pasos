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

type product struct {
	id, price int
}

func (p product) ID() int {
	return p.id
}

func (p product) Precio() int {
	return p.price
}

const MarketIdx = 0
const ProdIdIdx = 1
const PriceIdx = 2

func getProducto(producto []string) (string, int, int, bool) {
	super := producto[MarketIdx]
	id, errId := strconv.Atoi(producto[ProdIdIdx])
	precio, errPrecio := strconv.Atoi(producto[PriceIdx])
	if errId != nil || errPrecio != nil {
		return "", 0, 0, true
	}
	return super, id, precio, false
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return nil
	}
	precios := map[string]int{}
	for _, producto := range productos {
		super, id, precio, err := getProducto(producto)
		if err {
			return nil
		}
		for _, inputId := range ids {
			if inputId == id {
				precios[super] += precio
			}
		}
	}
	var carrito []Carrito
	for super, total := range precios {
		carrito = append(carrito, Carrito{
			Tienda: super,
			Precio: total,
		})
	}
	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		return -1
	}
	var totalPrecios float64
	var countPrecios float64
	for _, producto := range productos {
		_, id, precio, err := getProducto(producto)
		if err {
			return -1
		}
		if id == idProducto {
			totalPrecios += float64(precio)
			countPrecios++
		}
	}
	if countPrecios > 0 {
		return totalPrecios / countPrecios
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
	var min int
	exists := false
	for _, producto := range productos {
		_, id, precio, err := getProducto(producto)
		if err {
			return nil, false
		}
		if id == idProducto && (!exists || precio < min) {
			min = precio
			exists = true
		}
	}
	return product{
		id:    idProducto,
		price: min,
	}, exists
}
