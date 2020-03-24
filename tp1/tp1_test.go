package tp1

import (
	"reflect"
	"testing"
)

func TestTienda_CalcularPrecios(t *testing.T) {
	productos, err := LeerProductos("data.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre  string
		tiendas Tiendas
		input   []int
		output  []Carrito
	}{
		{
			nombre:  "calcula la suma de precios correcta",
			tiendas: productos,
			input:   []int{},
			output:  []Carrito{},
		},
		{
			nombre:  "da cero cuando no hay productos",
			tiendas: make(Tiendas, 0),
			input:   []int{},
			output:  nil,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			if resultado := test.tiendas.CalcularPrecios(test.input...); !reflect.DeepEqual(resultado, test.output) {
				t.Errorf("CalcularPrecios retorno %+v, se esperaba %+v\n", resultado, test.output)
			}
		})
	}
}

func TestTienda_Promedio(t *testing.T) {
	productos, err := LeerProductos("data.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre     string
		tiendas    Tiendas
		idProducto int
		output     float64
	}{
		{
			nombre:     "calcula el promedio correcto",
			tiendas:    productos,
			idProducto: 3,
			output:     0,
		},
		{
			nombre:     "da cero cuando el producto no existe",
			tiendas:    productos,
			idProducto: 101,
			output:     0,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			if resultado := test.tiendas.Promedio(test.idProducto); resultado != test.output {
				t.Errorf("Promedio retorno %f, se esperaba %f\n", resultado, test.output)
			}
		})
	}
}

func TestTienda_BuscarMasBarato(t *testing.T) {
	productos, err := LeerProductos("data.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre    string
		tiendas   Tiendas
		id        int
		resultado Producto
		existe    bool
	}{
		{
			nombre:    "retorna el producto correcto cuando existe",
			tiendas:   productos,
			id:        2,
			resultado: Producto{},
			existe:    true,
		},
		{
			nombre:    "retorna falso cuando el producto no existe",
			tiendas:   productos,
			id:        101,
			resultado: Producto{},
			existe:    false,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			resultado, existe := test.tiendas.BuscarMasBarato(test.id)
			switch {
			case !reflect.DeepEqual(resultado, test.resultado):
				t.Errorf("BuscarMasBarato retorno producto %+v, se esperaba %+v\n", resultado, test.resultado)
			case existe != test.existe:
				t.Errorf("BuscarMasBarato retorno existe en %v, se esperaba %v\n", existe, test.existe)
			}
		})
	}
}
