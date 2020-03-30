package tp2

import (
	"errors"
)

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if len(numeros) > 0 {

		input := make(chan int)
		output := make(chan int)

		go func() {

			for _, value := range numeros {
				input <- value
			}
			close(input)
		}()

		go func() {

			var sumTotal int = 0
			for x := range input {

				sumTotal += x
			}
			output <- sumTotal
		}()

		return <-output, nil

	}

	return 0, errors.New("Received empty array to perform sum operation")

}
