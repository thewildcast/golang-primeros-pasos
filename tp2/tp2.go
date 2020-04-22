package tp2

import (
	"fmt"
)

func sumar(c chan <- int, sumFunc sumador, num1 int, num2 int){
	c <- sumFunc(num1, num2)
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
    var result int

	if len(numeros) == 0 {
		return 0, fmt.Errorf("No hay numeros para sumar..")
	}

	if len(numeros)%2 != 0 {
		numeros = append(numeros, 0)
	}

	chResult := make(chan int, (len(numeros)/2))

	for i := 0; i < len(numeros); i+=2 {
        go sumar(chResult, sumFunc, numeros[i], numeros[i+1])
	}

	for i := 0; i < (len(numeros) / 2); i++ {
		result += <-chResult
	}

	return result, nil
}
