package tp2

import "fmt"

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	suma := 0

	if len(numeros) == 0 {
		return suma, fmt.Errorf("Error: lista de numeros vacía")
	}

	c := make(chan int)

	for i := 0; i < len(numeros); i += 2 {
		go func() {
			if i+1 >= len(numeros) {
				c <- numeros[i]
			} else {
				c <- sumFunc(numeros[i], numeros[i+1])
			}
		}()
		suma += <-c
	}

	return suma, nil
}
