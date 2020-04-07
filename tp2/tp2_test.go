package tp2

import (
	"testing"
)

func TestSumarLista(t *testing.T) {

	cases := []struct {
		nombre  string
		input   []int
		withErr bool
	}{
		{
			nombre:  "suma numeros correctamente",
			input:   []int{213, 355, 433, 566, 7789, 1243, 4356, 123450, 1235, 31455},
			withErr: false,
		},
		{
			nombre:  "maneja cantidad impar de numeros",
			input:   []int{213, 355, 433, 566, 7789, 1243, 4356, 123450, 1235},
			withErr: false,
		},
		{
			nombre:  "suma muchos numeros",
			input:   RandomNumbers(10_000, 1000),
			withErr: false,
		},
		{
			nombre:  "falla cuando no hay numeros para sumar",
			input:   nil,
			withErr: true,
		},
	}

	for _, test := range cases {
		t.Run(test.nombre, func(t *testing.T) {
			var sum int
			for _, d := range test.input {
				sum += d
			}
			resultado, err := SumarLista(Sumar, test.input...)
			switch {
			case resultado != sum:
				t.Errorf("SumarLista retorno %d, se esperaba %d", resultado, sum)
			case test.withErr && err == nil:
				t.Error("SumarLista debia retornar un error, pero se recibio nil")
			}
		})
	}
}
