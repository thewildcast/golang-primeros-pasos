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
		output  map[string]int
	}{
		{
			nombre:  "calcula la suma de precios correcta",
			tiendas: productos,
			input:   []int{},
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
			if resultado := test.tiendas.CalcularPrecios(1, 2, 3, 4, 5); !reflect.DeepEqual(resultado, test.output) {
				t.Errorf("CalcularPrecios retorno %+v, se esperaba %+v\n", resultado, test.output)
			}
		})
	}
}

func TestTienda_Promedio(t *testing.T) {
	casos := []struct {
		nombre string
		input  map[string][]Producto
		output int
	}{}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
		})
	}
}

func TestTienda_BuscarMasBarato(t *testing.T) {
	casos := []struct {
		nombre    string
		input     map[string][]Producto
		resultado Producto
		existe    bool
	}{}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
		})
	}
}
