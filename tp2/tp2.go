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

	wg.Add(2)

	valuesCh := gen(numeros)
	wg.Done()
	output := process(valuesCh, sumFunc)
	wg.Done()

	wg.Wait()

	return <-output, nil

}

func gen(nums []int) <-chan int {

	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out

}

func process(in <-chan int, sumFunc sumador) <-chan int {
	out := make(chan int)
	var t int = 0
	go func() {
		for n := range in {
			t = sumFunc(t, n)
		}
		out <- t
		close(out)
	}()
	return out
}
