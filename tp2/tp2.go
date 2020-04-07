package tp2

import (
	"fmt"
)

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	//Se retorna error si es nil
	if len(numeros) == 0 {
		return 0, fmt.Errorf("No se puede realizar suma con lista vacia")
	}
	//Se agrega 0 para emparejar slice impar
	if len(numeros)%2 != 0 {
		numeros = append(numeros, 0)
	}
	//Se crea canal con buffer, y se instancian goroutines para calculo concurrente
	n := len(numeros) / 2
	c := make(chan int, n)
	for i := 0; i < len(numeros); i += 2 {
		go func(a, b int) {
			s := sumFunc(a, b)
			c <- s
		}(numeros[i], numeros[i+1])
	}
	//Sumatoria de valores de goroutines
	var sum int
	for i := 0; i < n; i++ {
		sum += <-c
	}
	return sum, nil
}
