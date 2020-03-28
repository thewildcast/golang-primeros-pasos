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
	TIENDA_IDX = 0
	PROD_IDX = 1
	PRE_IDX = 2
)

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	// creo un mapa helper para buscar id sin hacer un range
	// para cada producto, no se si es mas eficiente
	// quisiera creer que si lo es en contraste con dos for anidados
	var mapaIds = make(map[string]int)
	for _, prodId := range ids {
		mapaIds[strconv.Itoa(prodId)] = prodId
	}

	// Inicializo el map con el valor zero de int
	var mapa = make(map[string]int)
	for _, entrada := range p {
		if _, found := mapaIds[entrada[PROD_IDX]]; found == true {
			pPrecio, _ := strconv.Atoi(entrada[PRE_IDX])
			// al usar el valor zero de int, no necesito revisar si el elemento
			// en el mapa existe o no
			mapa[entrada[TIENDA_IDX]] = mapa[entrada[TIENDA_IDX]] + pPrecio
		}
	}

	carritos := []Carrito{}
	for tienda, precioSumarizado := range mapa {
		carritos = append(carritos, Carrito{Tienda: tienda, Precio: precioSumarizado})
	}

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	var suma, contador int
	for _, entrada := range p {
		if prodId,_ := strconv.Atoi(entrada[PROD_IDX]); prodId == idProducto {
			precio, _ := strconv.Atoi(entrada[PRE_IDX])
			suma += precio
			contador++
		}
	}

	var promedio float64
	if contador > 0 {
		promedio = float64(suma) / float64(contador)
	}
	return promedio
}

type articulo struct {
	id, precio int
}

func (a articulo) ID() int {
	return a.id
}

func (a articulo) Precio() int {
	return a.precio
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	var art articulo = articulo{id: idProducto}
	var found bool = false
	for _, entrada := range p {
		if prodId,_ := strconv.Atoi(entrada[PROD_IDX]); prodId == idProducto {
			if precio, _ := strconv.Atoi(entrada[PRE_IDX]); precio < art.precio || !found {
				art.id = prodId
				art.precio = precio
			}
			found = true
		}
	}
	return art, found
}
