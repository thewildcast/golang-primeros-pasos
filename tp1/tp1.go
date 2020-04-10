package tp1

import "strconv"

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

const (
	SHOP_NAME_INDEX = 0
	PROD_ID_INDEX   = 1
	PRICE_INDEX     = 2
)

type Article struct {
	Id    int
	Price int
}

func (a Article) ID() int {
	return a.Id
}

func (a Article) Precio() int {
	return a.Price
}

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	var carrito []Carrito
	var m = make(map[string]int)
	for _, productId := range ids {
		for _, productEntry := range p {
			productEntryId, err := strconv.Atoi(productEntry[PROD_ID_INDEX])
			if err != nil {
				// error handling, not a number?
			}
			if productEntryId == productId {
				productEntryPrice, err := strconv.Atoi(productEntry[PRICE_INDEX])
				if err != nil {
					// error handling, not a number?
				}
				m[productEntry[SHOP_NAME_INDEX]] += productEntryPrice
			}
		}
	}
	for key, value := range m {
		carrito = append(carrito, Carrito{Tienda: key, Precio: value})
	}
	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var counter int
	var sum int
	for _, productEntry := range p {
		productEntryId, err := strconv.Atoi(productEntry[PROD_ID_INDEX])
		if err != nil {
			// error handling, not a number?
		}
		productEntryPrice, err := strconv.Atoi(productEntry[PRICE_INDEX])
		if err != nil {
			// error handling, not a number?
		}
		if productEntryId == idProducto {
			sum += productEntryPrice
			counter++
		}
	}
	if counter == 0 {
		return 0.0
	}
	return float64(sum) / float64(counter)
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var minPriceProduct Article = Article{Id: idProducto, Price: 0}
	found := false
	for _, productEntry := range p {
		productEntryId, err := strconv.Atoi(productEntry[PROD_ID_INDEX])
		if err != nil {
			// error handling, not a number?
		}
		productEntryPrice, err := strconv.Atoi(productEntry[PRICE_INDEX])
		if err != nil {
			// error handling, not a number?
		}
		if productEntryId == idProducto && (productEntryPrice <= minPriceProduct.Precio() || !found) {
			minPriceProduct = Article{Id: productEntryId, Price: productEntryPrice}
			found = true
		}
	}
	return minPriceProduct, found
}
