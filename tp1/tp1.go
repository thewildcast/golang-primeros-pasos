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

type ItemProduct struct {
	Name string
	stringId string 
	precio string
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

func (ip ItemProduct) ID() int {
	id, err	 :=	strconv.Atoi(ip.stringId)
	if err != nil {
		return 0 
	}
	return id
}

func (ip ItemProduct) Precio() int {
	precio, err	 :=	strconv.Atoi(ip.precio)
	if err != nil {
		return 0
	}
	return precio
}


// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {

	productos := []ItemProduct{}
	for _, pitem := range p{
		productos = append(productos, ItemProduct{pitem[0],pitem[1],pitem[2]})
	}
	carritos := map[string]Carrito{}
	returnCarritos := []Carrito{}
	for _, itemID := range ids{
		for _, prod := range productos{
			if prod.ID() == itemID{
				c , found := carritos[prod.Name]
				if found{
					carritos[prod.Name] = Carrito{prod.Name, c.Precio + prod.Precio()}
				}else{
					carritos[prod.Name] = Carrito{prod.Name, prod.Precio()}
				}
			     
			}
		}
	}
	for _, v := range carritos{
		returnCarritos = append(returnCarritos, v)
	}
	
	return returnCarritos
}

// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
	productos := []ItemProduct{}
	for _, pitem := range p{
		productos = append(productos, ItemProduct{pitem[0],pitem[1],pitem[2]})
	} 
	suma, count := 0, 0
	for _, prod := range productos{
		if prod.ID() == idProducto{
			count++
			suma += prod.Precio()
		}
	}
	if count == 0{
		return 0.0
	}

	return float64(suma)/float64(count)
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
	productos := []ItemProduct{}
	for _, pitem := range p{
		productos = append(productos, ItemProduct{pitem[0],pitem[1],pitem[2]})
	}
	var minorPrice ItemProduct
	found := false
	for _, prod := range productos{
		if prod.ID() == idProducto{
			if !found{
				minorPrice = prod
				found = true
			}
			if found && prod.Precio() < minorPrice.Precio(){
				minorPrice = prod
			}
		}
	}
	if found{
		return minorPrice, found
	}
	return ItemProduct{stringId: strconv.Itoa(idProducto)}, found

}
