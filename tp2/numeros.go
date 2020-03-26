package tp2

import (
	"encoding/json"
	"math/rand"
	"os"
	"time"
)

func LeerNumeros(archivo string) ([]int, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}

	numeros := []int{}
	if err := json.NewDecoder(f).Decode(&numeros); err != nil {
		return nil, err
	}

	return numeros, nil
}

func RandomNumbers(ceiling, total int) []int {
	numbers := []int{}
	for i := 0; i < total; i++ {
		numbers = append(numbers, rand.Intn(ceiling))
	}

	return numbers
}

// Sumar suma dos numeros pero agregar un random sleep
// para simular operaciones que tomen una cantidad de
// tiempo indeterminada.
func Sumar(a, b int) int {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	return a + b
}
