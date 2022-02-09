package tp4

import (
	"reflect"
	"testing"
)

func TestProductos_CalcularPreciosPorSuper(t *testing.T) {
	productos, err := LeerProductos("productos.json")
	if err != nil {
		t.Fatalf("no se puedo leer archivo de datos: %s", err)
	}

	casos := []struct {
		nombre      string
		productos   Productos
		inputProd   []int
		inputTienda []string
		precios     map[string]int
	}{
		{
			nombre:      "calcula la suma de precios correcta",
			productos:   productos,
			inputProd:   []int{1, 2, 2, 1},
			inputTienda: []string{"Dia", "SuperVea"},
			precios: map[string]int{
				"Dia":      7887 + 7847 + 7847 + 7887,
				"SuperVea": 1351 + 6844 + 6844 + 1351 ,
			},
		},
		{
			nombre:      "da cero cuando no hay productos",
			productos:   make(Productos, 0),
			inputProd:   []int{},
			inputTienda: []string{},
			precios:     map[string]int{},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			carritos := test.productos.CalcularPrecios(test.inputProd, test.inputTienda)
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
