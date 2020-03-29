package tp2

type emptyListError struct{
}

func (e emptyListError) Error()string{
	return "The list is empty"
}

func concurrentSum(sumFunc sumador, numeros []int, sum chan int) {
	if len(numeros) == 1 {
		sum <- numeros[0]
	}
	if len(numeros) % 2 != 0 {
		numeros = append(numeros, 0)
	}

	results := make(chan int, len(numeros)/2)
	defer close(results)
	arr := make([]int, len(numeros)/2)

	for i := 0; i<len(numeros); i+=2 {
		go func(first int, second int, results chan int){
			results<-sumFunc(first, second)
		}(numeros[i], numeros[i+1], results)
	}

	for i := 0; i<len(numeros)/2; i++ {
		arr[i] = <-results
	}

	concurrentSum(sumFunc, arr, sum)
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	if len(numeros) == 0 {
		return 0, emptyListError{}
	}

	sum := make(chan int)

	go concurrentSum(sumFunc, numeros, sum)

	return <-sum, nil
}
