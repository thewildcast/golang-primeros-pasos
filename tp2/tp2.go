package tp2

import (
	"fmt"
	"log"
)

func SumaArray(sumandos chan int, result chan int) {

	var total int

	for elem := range sumandos {

		total = total + elem
	}

	result <- total

}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if len(numeros) < 2 {
		return 0, fmt.Errorf("No hay numeros suficientes para sumar")
	}

	sumandos := make(chan int, len(numeros))

	result := make(chan int)

	go SumaArray(sumandos, result)

	for _, num := range numeros {

		sumandos <- num
	}

	close(sumandos)

	total := <-result

	log.Println(":😄 Total:", total)

	return total, nil
}
