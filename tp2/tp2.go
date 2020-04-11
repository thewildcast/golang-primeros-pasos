package tp2

import (
	"errors"
	"sync"
)

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	var wg sync.WaitGroup
	if len(numeros) == 0 {

		return 0, errors.New("Received empty array to perform sum operation")
	}

	if len(numeros)%2 != 0 {

		numeros = append(numeros, 0)
	}

	output := make(chan int, len(numeros)/2)
	wg.Add(len(numeros) / 2)

	var total int = 0

	for i := 0; i < len(numeros); i = i + 2 {

		go func(x int) {

			j := sumFunc(numeros[x], numeros[x+1])
			output <- j

		}(i)

		wg.Done()

	}

	for k := 0; k < len(numeros)/2; k++ {
		total += <-output
	}

	wg.Wait()

	return total, nil
}
