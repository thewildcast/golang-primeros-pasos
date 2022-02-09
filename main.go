package main

import (
	"github.com/wildcast/golang-primeros-pasos/tp4"
	"fmt"
)

func main() {
	carrito := tp4.GetProductoTienda(tp4.ProductoTienda{ID: "1", Tienda: "dia"})
	fmt.Println(carrito)
}
