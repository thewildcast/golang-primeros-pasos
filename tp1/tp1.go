package tp1

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

// Estructura que implementara ambos metodos ID y Precio
type productImplementation []int

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

	var pricesResult []Carrito
	var tienda string
	collector := map[string]int{}
	var collectedResult map[string]int

	for _, singleProduct := range p {

		for _, marketID := range ids {
			foundID := convertStringToInt(singleProduct[1])
			if foundID == marketID {

				tienda = singleProduct[0]

				price := convertStringToInt(singleProduct[2])

				collectedResult = verifyAndProcess(tienda, price, collector)
			}
		}
	}

	for store, price := range collectedResult {

		pricesResult = append(pricesResult, Carrito{store, price})
	}

	return pricesResult
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {

	var sumAllPrices int = 0
	var dataLeng int = 0
	var found bool

	for _, singleProduct := range p {

		foundID := convertStringToInt(singleProduct[1])

		if foundID == idProducto {
			found = true
			dataLeng++

			sumAllPrices += convertStringToInt(singleProduct[2])

		}
	}

	if !found {
		return 0
	}

	return float64(sumAllPrices) / float64(dataLeng)
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {

	var tmpValue int = 4294967295
	var found bool

	for _, value := range p {

		foundPrice := convertStringToInt(value[1])
		if foundPrice == idProducto {

			found = true
			priceConverted := convertStringToInt(value[2])

			if priceConverted < tmpValue {
				tmpValue = priceConverted
			}
		}
	}

	if !found {
		return productImplementation{idProducto, 0}, found
	}

	return productImplementation{idProducto, tmpValue}, true
}

//Id implementation
func (r productImplementation) ID() int {
	return r[0]
}

//price implementation
func (r productImplementation) Precio() int {
	return r[1]
}
