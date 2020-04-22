package tp1

import "strconv"

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}

type product struct{
	id int
	price int
}

func (p product) ID() int {
	return p.id
}

func (p product) Precio() int {
	return p.price
}

func (products Productos) getByIDAndMarketName(id int, supermarketName string) (product, bool) {
	for _, p := range products {
		supermarket_name := p[0]
		product_id, err := strconv.Atoi(p[1])
		price, err := strconv.Atoi(p[2])
		if err != nil {
			panic(err)
		}
		if id == product_id && supermarket_name == supermarketName {
			return product{price: price, id: product_id}, true
		}
	}
	return product{}, false
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
	var carrito []Carrito
	supermarkets := make(map[string]int)

	for _, supermarket_name := range Supermercados {
		for _, product_id := range ids {
			if product, found := p.getByIDAndMarketName(product_id, supermarket_name); found {
				supermarkets[supermarket_name] += product.Precio()
			}
		}
	}

	for supermarket_name, price := range supermarkets {
		carrito = append(carrito, Carrito{
			Tienda: supermarket_name,
			Precio: price,
		})
	}

	return carrito
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	count := 0.0
	sum := 0
	var average float64 = 0.0

	for _, supermarket_name := range Supermercados {
		if product, found := p.getByIDAndMarketName(idProducto, supermarket_name); found {
			count += 1.0
			sum += product.Precio()
		}
	}

	if sum > 0 {
		average = float64(sum) / count
	}

	return average
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	cheaper_product := product{price:0, id:idProducto}
	product_found := false

	for _, supermarket_name := range Supermercados {
		if product, found := p.getByIDAndMarketName(idProducto, supermarket_name); found {
			product_found = true

			if cheaper_product.Precio() == 0 {
				cheaper_product = product
				continue
			}

			if product.Precio() < cheaper_product.Precio(){
				cheaper_product = product
			}
		}
	}

	return cheaper_product, product_found
}
