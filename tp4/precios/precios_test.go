package precios

import (
	"reflect"
	"testing"
)

func TestCalcularPrecios(t *testing.T) {
	casos := []struct {
		nombre  string
		pIDs    []int
		sIDs    []string
		precios map[string]int
		errores bool
	}{
		{
			nombre: "calcula la suma de precios correcta",
			pIDs:   []int{1, 2},
			sIDs:   []string{"target", "coto", "dia", "disco", "jumbo"},
			precios: map[string]int{
				"target": 8536,
				"coto":   3923,
				"dia":    15734,
				"disco":  8866,
				"jumbo":  10194,
			},
			errores: false,
		},
		{
			nombre:  "da error si el supermercado no existe",
			pIDs:    []int{1, 2},
			sIDs:    []string{"sssss", "coto", "dia", "disco", "jumbo"},
			precios: nil,
			errores: true,
		},
		{
			nombre:  "da error si no existe el producto",
			pIDs:    []int{-1, 2},
			sIDs:    []string{"coto", "dia", "disco", "jumbo"},
			precios: nil,
			errores: true,
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			service := NewService()
			carritos, err := service.CalcularPrecios(test.pIDs, test.sIDs)
			if err != nil && !test.errores {
				t.Errorf("Se esperaba que el calculo de precios retornara con exito pero dio error: %v", err)
			}
			if err == nil && !test.errores {
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
			}
		})
	}
}
