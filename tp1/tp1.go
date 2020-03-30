package tp1

import "strconv"

type Producto interface {
	BuscarMasBarato()
}

type producto struct {
	id     int
	precio float64
}

type Carrito struct {
	supermercado string
	productos    []producto
}

func (c Carrito) CalcularPrecios(productosIds []int) []Carrito {
	productos, _ := LeerProductos("productos.json")
	carrito := []Carrito{}

	for _, paramId := range productosIds {
		for _, rawProducto := range productos {
			id, _ := strconv.Atoi(rawProducto[1])
			if paramId == id {
				precio, _ := strconv.ParseFloat(rawProducto[2], 64)
				p := producto{id: id, precio: precio}

				c.productos = append(c.productos, p)
				c.supermercado = rawProducto[0]

				carrito = append(carrito, c)
			}
		}
	}

	return carrito
}

func (c Carrito) Promedio(productoId int) int {
	productos, _ := LeerProductos("productos.json")
	precios := []float64{}
	var sum float64

	for _, producto := range productos {
		id, _ := strconv.Atoi(producto[1])
		if productoId == id {
			precio, _ := strconv.ParseFloat(producto[2], 64)
			precios = append(precios, precio)
		}
	}

	for i := range precios {
		sum += precios[i]
	}

	return int(int(sum) / len(precios))
}

func (c Carrito) BuscarMasBarato(ProductoId int) (Carrito, bool) {
	productos, _ := LeerProductos("productos.json")
	carrito := []Carrito{}

	// Armo lista de productos que matchean con el ID
	for _, pr := range productos {
		id, _ := strconv.Atoi(pr[1])
		if ProductoId == id {
			precio, _ := strconv.ParseFloat(pr[2], 64)
			p := producto{id: id, precio: precio}

			c.supermercado = pr[0]
			c.productos = append(c.productos, p)

			carrito = append(carrito, c)
		}
	}

	// De un carrito con productos de diferentes supermercados,
	// Buscamos precio más bajo
	productoMasBarato := carrito[0] // Tomamos un precio inicial
	for _, item := range carrito {
		for _, producto := range item.productos {
			if producto.precio < productoMasBarato.productos[0].precio {
				productoMasBarato.supermercado = item.supermercado
			}
		}
	}

	return productoMasBarato, true
}
