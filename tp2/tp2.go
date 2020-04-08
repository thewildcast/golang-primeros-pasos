package tp2

import (
	"errors"
	"fmt"
)

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	defer func() { fmt.Println("Gracias por el curso! xD")}()

	if len(numeros) <= 0 {
		return 0, errors.New("No hay numeros para sumar")
	}

	total := sumarLista(numeros)

	return total, nil
}

func sumarLista(numeros []int) int {
	tamañoBuffer := len(numeros) / 2

	if len(numeros) % 2 != 0 {
		tamañoBuffer += 1
	}

	resultados := make(chan int, tamañoBuffer)
	defer close(resultados)

	for i := 0; i < len(numeros); i += 2 {
		a := numeros[i]
		var b int
		if i+1 < len(numeros) {
			b = numeros[i+1]
		}

		go func(a, b int) {
			resultados <- a + b
		}(a, b)
	}

	var totales []int
	for i := 0; i < tamañoBuffer; i++ {
		totales = append(totales,<-resultados)
	}

	if len(totales) > 1{
		return sumarLista(totales)
	}

	return totales[0]
}