package tp2

type noNumbersError struct {
	s string
}

func (e *noNumbersError) Error() string {
	return e.s
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	length := len(numeros)
	if length == 0 {
		return 0, &noNumbersError{s: "No numbers to sum"}
	}

	source := make(chan int, length)

	var result chan int
	if length%2 == 0 {
		result = make(chan int, length/2)
	} else {
		result = make(chan int, (length/2)+1)
	}

	total := make(chan int)

	for _, number := range numeros {
		source <- number
	}
	go recursiveSum(source, result, sumFunc, total)
	return <-total, nil
}

func recursiveSum(src chan int, res chan int, sumFun sumador, total chan int) {
	if cap(res) < 2 {
		total <- sumFun(<-src, <-src)
	} else {
		if cap(src)%2 != 0 {
			res <- <-src
		}
		var newRes chan int
		if cap(res)%2 == 0 {
			newRes = make(chan int, cap(res)/2)
		} else {
			newRes = make(chan int, (cap(res)/2)+1)
		}
		for i := 0; i < cap(src)/2; i++ {
			go func() {
				res <- sumFun(<-src, <-src)
			}()
		}
		go recursiveSum(res, newRes, sumFun, total)
	}
}
