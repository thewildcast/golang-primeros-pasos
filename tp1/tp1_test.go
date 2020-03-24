package tp1

import "testing"

func TestTienda_Sumar(t *testing.T) {
	casos := []struct {
		nombre string
		input  [][]int
		output int
	}{}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
		})
	}
}
