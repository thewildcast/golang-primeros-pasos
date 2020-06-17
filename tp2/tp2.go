package tp2

import "fmt"

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.

func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	//return sumFunc(numeros[0], numeros[1]), nil

	if len(numeros) < 2 {
		return 0, fmt.Errorf("Error. Se requieren al menos 2 numeros para sumar.")
	}
	subtotals := make(chan int)
	suma := 0
	for i := 0; i < len(numeros); i += 2 {
		go func(i int) {
			if i+1 >= len(numeros) {
				subtotals <- numeros[i]
			} else {
				subtotals <- Sumar(numeros[i], numeros[i+1])
			}
		}(i)
		suma += <-subtotals
	}
	return suma, nil
}
