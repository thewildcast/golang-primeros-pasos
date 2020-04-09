package tp2

import "fmt"

type pares struct {
	a, b int
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if len(numeros) == 0 {
		return 0, fmt.Errorf("Lista de numeros vacia")
	}
	var scheduled int
	results := make(chan int, len(numeros)/2)
	toSum := make(chan pares, len(numeros)/2)
	// creamos una gorutine nueva por cada funcion sumar a ejecutar
	for i := 0; i < len(numeros)/2; i++ {
		go SumIndep(toSum, results, sumFunc)
	}
	// juntamos los pares de numeros y los transmitimos al canal
	for i := 0; i < len(numeros)/2; i++ {
		toSum <- pares{numeros[i], numeros[i+len(numeros)/2]}
		scheduled++
	}
	var resultadotemp pares
	if len(numeros)%2 != 0 {
		resultadotemp.a = numeros[len(numeros)-1]
	}

	// esperamos los resultados

	var finished bool
	var return_val int
	for r := range results {
		scheduled--
		if finished {
			return_val = r
			close(toSum)
			break
		}
		if scheduled == 0 {
			finished = true
		}
		if resultadotemp.a == 0 {
			resultadotemp.a = r
		} else {
			resultadotemp.b = r
			toSum <- resultadotemp
			resultadotemp.a = 0
			resultadotemp.b = 0
			scheduled++
		}
	}
	if finished {
		close(results)
	}
	return return_val, nil
}

func SumIndep(valores chan pares, results chan int, sumFunc sumador) {
	for v := range valores {
		r := sumFunc(v.a, v.b)
		results <- r
	}
}
