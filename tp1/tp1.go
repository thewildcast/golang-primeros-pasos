package tp1

import (
	"strconv"
	"sort"
)

// Producto contiene metodos que nos permiten acceder
// a atributos que esperamos de un Producto.
type Producto interface {
	ID() int
	Precio() int
}


type producto struct {
	mercado string
	id int
	precio int
}

func (p *producto) Mercado() string {
	return p.mercado
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

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	var carritos []Carrito
	carritosIndex := make(map[string]*Carrito)

	sort.Ints(ids)

	for _, array := range p {
		productoArray, err := convertirArrayAProducto(array)
		if err != nil {
			panic(err)
		}

		i := sort.SearchInts(ids, productoArray.ID())

		if(i != len(ids) && productoArray.ID() == ids[i]){
			if(carritosIndex[productoArray.Mercado()] == nil) {
				carritosIndex[productoArray.Mercado()] = &Carrito{ Tienda: productoArray.Mercado(), Precio: 0 };
			}

			carritosIndex[productoArray.Mercado()].Precio += productoArray.Precio()
		}
	}

	for _, carrito := range carritosIndex {
        carritos = append(carritos, *carrito)
    }

	return carritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	precioTotal := 0.0
	cantidadProducto := 0.0

	for _, array := range p {
		productoArray, err := convertirArrayAProducto(array)
		if err != nil {
			panic(err)
		}
		if ( idProducto == productoArray.ID())  {
			precioTotal += float64(productoArray.Precio())
			cantidadProducto++
		} 
	}

	if(cantidadProducto == 0){
		return 0.0
	}

	return precioTotal/cantidadProducto
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	encontrado := false
	masBarato := &producto{ id: idProducto, precio: 0 }
	for _, array := range p {
		productoArray, err := convertirArrayAProducto(array)
		if err != nil {
			panic(err)
		}
		if ( idProducto == productoArray.ID() && (masBarato.Precio() == 0 || productoArray.Precio() < masBarato.Precio()) )  {
			masBarato = productoArray
			encontrado = true
		} 
	}

	return masBarato, encontrado
}

func convertirArrayAProducto(array []string) (*producto, error) {
	mercado := array[0]


	id, err := strconv.Atoi(array[1])
	if err != nil {
		return nil, err
	}

	precio, err := strconv.Atoi(array[2])
	if err != nil {
		return nil, err
	}

	return &producto{ mercado: mercado, id: id, precio: precio }, nil
}