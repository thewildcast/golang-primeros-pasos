package tp2

import (
	"errors"
)

type par struct {
	x, y int
}

const numWorkers = 8

func sumWorker(sumFunc sumador, pares <-chan par, sumas chan<- int) {
	for p := range pares {
		s := sumFunc(p.x, p.y)
		sumas <- s
	}
}

func sumCollector(sumas <-chan int, maxSumas int, pares chan<- par, suma chan<- int) {
	numSumas := 0
	for {
		s1 := <-sumas
		numSumas++
		if numSumas == maxSumas {
			suma <- s1
			break
		}
		s2 := <-sumas
		numSumas++
		pares <- par{s1, s2}
	}
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	N := len(numeros)
	if N == 0 {
		return 0, errors.New("No hay numeros para sumar")
	}
	pares := make(chan par)
	sumas := make(chan int, N)
	suma := make(chan int)
	for i := range numeros {
		sumas <- numeros[i]
	}
	w := 0
	for w < numWorkers {
		go sumWorker(sumFunc, pares, sumas)
		w++
	}
	go sumCollector(sumas, 2*N-1, pares, suma)
	return <-suma, nil
}
