package tp1

type Producto interface {
	ID() int
	Precio() int
}

type Productos [][]string

type Carrito struct {
	Tienda string
	Precio int
}

func (p Productos) CalcularPrecios(ids ...int) []Carrito {
	return nil
}

func (p Productos) Promedio(idProducto int) int {
	return 0
}

func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	return nil, false
}
