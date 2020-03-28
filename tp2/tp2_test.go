package tp2

import (
	"testing"
)

func TestSumarLista(t *testing.T) {
	numeros, err := LeerNumeros("numeros.json")
	if err != nil {
		t.Fatalf("no se pudo leer los numeros del archivo. Asegurate de correr los tests estando parado/a dentro de la carpeta tp2. Error: %s", err)
	}

	cases := []struct {
		nombre    string
		input     []int
		resultado int
		withErr   bool
	}{
		{
			nombre:    "suma numeros correctamente",
			input:     numeros,
			resultado: 2480947,
			withErr:   false,
		},
		{
			nombre:    "falla cuando no hay numeros para sumar",
			input:     nil,
			resultado: 0,
			withErr:   true,
		},
	}

	for _, test := range cases {
		t.Run(test.nombre, func(t *testing.T) {
			resultado, err := SumarLista(Sumar, test.input...)
			switch {
			case resultado != test.resultado:
				t.Errorf("SumarLista retorno %d, se esperaba %d", resultado, test.resultado)
			case test.withErr && err == nil:
				t.Error("SumarLista debia retornar un error, pero se recibio nil")
			}
		})
	}
}
