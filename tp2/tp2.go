package tp2

type customError struct {
}

func (e customError) Error() string {
	return "numeros cant be empty"
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	// si mi lista de numeros no tiene nada, devolver error.
	if len(numeros) == 0 {
		return 0, customError{}
	}

	var contador int

	valorNumero := len(numeros)
	buffer := len(numeros) / 2

	sumas := make(chan bool, buffer)

	// go routines para listas impares
	if len(numeros)%2 != 0 {
		valorNumero--
		buffer++

		go func() {
			contador += sumFunc(numeros[len(numeros)-1], 0)
			sumas <- true
		}()
	}

	// ahora escribo las go routines para listas pares, iterando sobre valorNumero
	for i := 0; i < valorNumero; i += 2 {
		go func(index int) {
			contador += sumFunc(numeros[index], numeros[index+1])
			sumas <- true
		}(i)
	}

	//ahora itero sobre el tamano del buffer para leer en una sola variable cada uno de mis channels
	for i := 0; i < buffer; i++ {
		<-sumas
	}
	close(sumas)

	return contador, nil
}
