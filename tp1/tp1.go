package tp1

type Tiendas map[string][]Producto

type Carrito struct {
	Tienda string
	Precio int
}

func (t Tiendas) CalcularPrecios(ids ...int) []Carrito {
	return nil
}

func (t Tiendas) Promedio(idProducto int) float64 {
	return 0
}

func (t Tiendas) BuscarMasBarato(idProducto int) (Producto, bool) {
	return Producto{}, false
}
