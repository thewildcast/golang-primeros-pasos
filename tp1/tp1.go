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

// CalcularPrecios recibe un arreglo de los IDs de productos y calcula,
// para cada super mercado, cuanto saldria comprar esos productos ahi.
// Retorna un slice de carritos, donde se tiene uno para cada super mercado.
func (p Productos) CalcularPrecios(ids ...int) []Carrito {
    //fmt.Println(ids)
    var carrito []Carrito
    
    //recorro los id de productos.
    for _,idValorProducto := range ids {
        //por cada id de producto recorro los sumermercados
        for _,valorProducto := range p {
        	idValorProductoToString := strconv.Itoa(idValorProducto)
            //cuando encuentro el producto en el supermercado cargo un slice 
            if idValorProductoToString == valorProducto[1]{
                var indiceAux int
                var existe bool
                //valido si el super ya esta cargado en el slice
                indiceAux, existe = existeTiendaEnSlice(carrito, valorProducto[0])
                var valorProductoIntAux int
                valorProductoIntAux,_ = strconv.Atoi(valorProducto[2]) 
                //si no existe lo cargo e caso contrario actualizo el precio.
                if (!existe) {
                    //cuando es nueva tienda en el slice
                    carritoAux := Carrito{valorProducto[0],valorProductoIntAux}
                    carrito = append(carrito, carritoAux) 
                }else{
                    //si la tienda ya existe en el slice actualizo el valor.
                    var PrecioAux int 
                    PrecioAux = carrito[indiceAux].Precio + valorProductoIntAux
                    var carritoAux Carrito
                    carritoAux = Carrito{valorProducto[0],PrecioAux}
                    carrito[indiceAux] = carritoAux            
                }  
            }
                
        }
    }
	return carrito
}

//retorna indice y valor
func existeTiendaEnSlice(slice []Carrito, busqueda string ) (int , bool) {
	for indice, numero := range slice {
		if numero.Tienda == busqueda {
			return indice,true
		}
	}
	return 0,false
}
// Promedio recibe el id de un producto y retorna el precio promedio
// de ese producto usando los precios de todos los supermercados.
func (p Productos) Promedio(idProducto int) float64 {
    var suma int = 0
    var cant int = 0
    for _,valorProducto := range p {
    	idValorProductoToString := strconv.Itoa(idProducto)
        if idValorProductoToString == valorProducto[1]{
            var valor int
            valor,_ = strconv.Atoi(valorProducto[2])
            suma = suma + valor
            cant++
        }
    }
    //me pelee con el tipo de dato!
    var resultado float64
    if(suma != 0 && cant !=0){
        resultado = float64(suma)/float64(cant)
    }else {resultado=0.0}
	return resultado
}

func (id ProductoBarato) ID() int {
    return id.IDProdu
}

func (precio ProductoBarato) Precio() int {
    return precio.PrecioProdu
    
}

type ProductoBarato struct {
	IDProdu int
    PrecioProdu int
}

// BuscarMasBarato recibe un id de producto a buscar y te retorna
// el producto mas barato que haya encontrado.
func (p Productos) BuscarMasBarato(idProducto int) (Producto, bool) {
    var productoMasBarato ProductoBarato
    var siExiste bool = false
    var precioProducto int
    
    for _,valorProducto := range p {
    	idValorProductoToString := strconv.Itoa(idProducto)
        if idValorProductoToString == valorProducto[1]{
            precioProducto,_ = strconv.Atoi(valorProducto[2])          
            if(productoMasBarato.PrecioProdu == 0 || productoMasBarato.PrecioProdu > precioProducto){
                siExiste=true
                productoMasBarato=ProductoBarato{idProducto,precioProducto}
            }   
        }
    }
    if(!siExiste) { productoMasBarato.IDProdu = idProducto}
	return productoMasBarato, siExiste
}