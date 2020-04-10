package tp2

import (
	"errors"
)

const NUMBER_OF_GOROUTINES = 10

type operands struct {
	left  int
	right int
}

func reduce(toSum chan operands, results chan int) {
	for t := range toSum {
		results <- (t.left + t.right)
	}
}

func produce(cantidadSumas int, toSum chan operands, results chan int) {
	for i := 0; i < cantidadSumas; i++ {
		toSum <- operands{left: <-results, right: <-results}
	}
	close(toSum)
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	if numeros == nil {
		return 0, errors.New("No hay numeros a sumar")
	}

	toSum := make(chan operands, NUMBER_OF_GOROUTINES)
	results := make(chan int, NUMBER_OF_GOROUTINES*2)
	done := make(chan bool)

	for i := 0; i < NUMBER_OF_GOROUTINES; i++ {
		go func() {
			reduce(toSum, results)
		}()
	}

	go func() {
		produce(len(numeros)-1, toSum, results)
		done <- true
	}()

	for i := 0; i < len(numeros); i++ {
		results <- numeros[i]
	}

	<-done

	return <-results, nil
}
