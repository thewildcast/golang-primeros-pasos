package tp4

import (
	"reflect"
	"testing"
)

func TestObtenerProducto(t *testing.T) {
	casos := []struct {
		nombre    string
		tienda    string
		id        int32
		precio    int32
		productos chan Producto
	}{
		{
			nombre:    "obtiene el producto 1 de la tienda dia",
			tienda:    "dia",
			id:        1,
			precio:    7887,
			productos: make(chan Producto, 1),
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			ObtenerProducto(test.id, test.tienda, test.productos)
			prod := <-test.productos

			if !reflect.DeepEqual(prod.Precio, test.precio) {
				t.Errorf("ObtenerProductos retorna el precio incorrecto %+v, se esperaban %+v", prod.Precio, test.precio)
			}
		})
	}
}

func TestCalcularPorTienda(t *testing.T) {
	casos := []struct {
		nombre      string
		tienda      string
		ids         []int32
		total       int32
		tiendasChan chan Carrito
	}{
		{
			nombre:      "obtiene un carrito con la suma de ids",
			tienda:      "dia",
			ids:         []int32{1, 2, 3},
			total:       15793,
			tiendasChan: make(chan Carrito, 1),
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			CalcularPorTienda(test.ids, test.tienda, test.tiendasChan)

			carrito := <-test.tiendasChan

			if !reflect.DeepEqual(carrito.Total, test.total) {
				t.Errorf("CalcularPorTienda retorna el total incorrecto %+v, se esperaban %+v", carrito.Total, test.total)
			}
		})
	}
}

func TestCalcularPorTiendas(t *testing.T) {
	casos := []struct {
		nombre  string
		tiendas []string
		ids     []int32
		totales map[string]int32
	}{
		{
			nombre:  "obtiene un carrito con la suma de ids",
			tiendas: []string{"dia", "carrefour", "disco"},
			ids:     []int32{1, 2, 3},
			totales: map[string]int32{"dia": 15793, "carrefour":9109, "disco":10623},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			carritos := CalcularPorTiendas(test.ids, test.tiendas)

			if len(carritos) != len(test.tiendas) {
				t.Errorf("CalcularPorTiendas retorna la cantidad de carritos incorrecto %+v, se esperaban %+v", len(carritos), len(test.tiendas))
			}

			for _, carrito := range carritos {
				if !reflect.DeepEqual(carrito.Total, test.totales[carrito.Tienda]) {
					t.Errorf("CalcularPorTiendas retorna el total incorrecto %+v, se esperaban %+v", carrito.Total, test.totales[carrito.Tienda])
				}
			}
		})
	}
}
