package tp2

import (
	"errors"
	"fmt"
)

type sumar struct {
	num1, num2 int
}

func sumarSum(sumFunc sumador, jobs <-chan sumar, result chan<- int) {
	for {
		sum, ok := <-jobs
		if !ok {
			return
		}
		res := sumFunc(sum.num1, sum.num2)
		result <- res
	}
}

func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	if len(numeros) == 0 {
		return 0, errors.New("No hay numeros para sumar")
	}
	if len(numeros) == 1 {
		return numeros[0], nil
	}
	jobs := make(chan sumar, len(numeros)/2)
	result := make(chan int, len(numeros)/2)
	defer close(jobs)
	defer close(result)
	for i := 0; i < len(numeros)/2; i++ {
		go sumarSum(sumFunc, jobs, result)
	}
	for i := 0; i < len(numeros); i += 2 {
		if i < len(numeros)-1 {
			jobs <- sumar{numeros[i], numeros[i+1]}
		} else {
			result <- numeros[i]
		}
	}
	var receivedResults int
	for {
		fmt.Println("Waiting for result")
		num1 := <-result
		fmt.Println("result received")
		receivedResults++
		if receivedResults >= len(numeros)-1 {
			return num1, nil
		}
		num2 := <-result
		receivedResults++
		jobs <- sumar{num1, num2}
	}
}
