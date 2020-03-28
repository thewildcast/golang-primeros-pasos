package tp2

import "fmt"

type pares struct {
	a, b int
}

func sumarThread(p pares, resultadosChan chan int, sumFunc sumador) {
	if p.a != 0 && p.b != 0 {
		resultadosChan <- sumFunc(p.a, p.b)
	} else if p.a != 0 {
		resultadosChan <- p.a
	} else if p.b != 0 {
		resultadosChan <- p.b
	}
}

func sumarLista(numeros []int, resultadoChan chan int, sumFunc sumador) {
	if len(numeros)%2 != 0 {
		numeros = append(numeros, 0)
	}

	resultadosChan := make(chan int)

	for i := 0; i < len(numeros); i = i + 2 {
		go sumarThread(pares{a: numeros[i], b: numeros[i+1]}, resultadosChan, sumFunc)
	}

	var resultados []int
	for i := 0; i < len(numeros)-1; i++ {
		r := <-resultadosChan
		resultados = append(resultados, r)
		if len(resultados) == 2 {
			go sumarThread(pares{a: resultados[0], b: resultados[1]}, resultadosChan, sumFunc)
			resultados = resultados[:0]
		}
	}
	close(resultadosChan)

	resultadoChan <- resultados[0]
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("No hay numeros para sumar")
	}

	resultadoChan := make(chan int)
	go sumarLista(numeros, resultadoChan, sumFunc)
	resultado := <-resultadoChan

	return resultado, nil
}
