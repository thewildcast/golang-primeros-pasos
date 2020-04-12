package tp2

import (
	"fmt"
)

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	nums := make(chan int, len(numeros))
	result := make(chan int)
	if len(numeros) < 2 {
		return 0, fmt.Errorf("No se puede")
	}
	for _, num := range numeros {
		nums <- num
	}
	close(nums)

	go func() {
		var total int
		for n := range nums {
			total = sumFunc(total, n)
		}
		result <- total
	}()

	total := <-result
	close(result)

	return total, nil
}
