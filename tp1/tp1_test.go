package tp1

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTienda_CalcularPrecios(t *testing.T) {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre    string
		productos Productos
		input     []int
		precios   map[string]int
	}{
		{
			nombre:    "calcula la suma de precios correcta",
			productos: productos,
			input:     []int{1, 2},
			precios: map[string]int{
				"Target":             8536,
				"Coto":               3923,
				"Dia":                15734,
				"Disco":              8866,
				"Jumbo":              10194,
				"Macro":              20559,
				"Nini":               12053,
				"SuperVea":           8195,
				"Wallmart":           10539,
				"Whole Foods Market": 12785,
				"Carrefour":          6910,
			},
		},
		{
			nombre:    "da cero cuando no hay productos",
			productos: make(Productos, 0),
			input:     []int{},
			precios:   map[string]int{},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			carritos := test.productos.CalcularPrecios(test.input...)
			if len(carritos) != len(test.precios) {
				t.Errorf("CalcularPrecios retorno %d supermercados, se esperaban %d", len(carritos), len(test.precios))
			}

			resultado := map[string]int{}
			for _, carrito := range carritos {
				resultado[carrito.Tienda] = carrito.Precio
			}

			if !reflect.DeepEqual(resultado, test.precios) {
				t.Errorf("CalcularPrecios retorna precios incorrectos %+v, se esperaban %+v", resultado, test.precios)
			}
		})
	}
}

func TestTienda_Promedio(t *testing.T) {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre     string
		productos  Productos
		idProducto int
		output     int
	}{
		{
			nombre:     "calcula el promedio correcto",
			productos:  productos,
			idProducto: 3,
			output:     4912,
		},
		{
			nombre:     "da cero cuando el producto no existe",
			productos:  productos,
			idProducto: 101,
			output:     0,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			if resultado := test.productos.Promedio(test.idProducto); resultado != test.output {
				fmt.Println(resultado)
				t.Errorf("Promedio retorno %d, se esperaba %d", resultado, test.output)
			}
		})
	}
}

func TestTienda_BuscarMasBarato(t *testing.T) {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre    string
		productos Productos
		id        int
		precio    int
		existe    bool
	}{
		{
			nombre:    "retorna el producto correcto cuando existe",
			productos: productos,
			id:        2,
			precio:    1509,
			existe:    true,
		},
		{
			nombre:    "retorna falso cuando el producto no existe",
			productos: productos,
			id:        101,
			precio:    0,
			existe:    false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			resultado, existe := test.productos.BuscarMasBarato(test.id)
			fmt.Println(resultado)
			switch {
			case resultado.Precio() != test.precio:
				t.Errorf("BuscarMasBarato retorno precio %d, se esperaba %d", resultado.Precio(), test.precio)
			case resultado.ID() != test.id:
				t.Errorf("BuscarMasBarato retorno id %d, se esperaba %d", resultado.ID(), test.id)
			case existe != test.existe:
				t.Errorf("BuscarMasBarato retorno existe en %v, se esperaba %v\n", existe, test.existe)
			}
		})
	}
}
