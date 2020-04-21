package tp2

type sinNumerosError struct {
}

func (i sinNumerosError) Error() string {
	return "No se ingresaron números"
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	//Si no me ingresan nada, devuelvo el error
	if len(numeros) == 0 {
		return 0, sinNumerosError{}
	}

	var res int

	numbersQty := len(numeros)
	bufferSize := len(numeros) / 2

	//Me fijo la cantidad de iteraciones que tengo que hacer. Esto es porque si tengo cantidades pares o impares tengo que hacer distintas cosas
	if len(numeros)%2 != 0 {
		numbersQty = numbersQty - 1
		//Aumento el tamaño para que en el caso de tener 9 valores, tener 5 channels
		bufferSize++
	}

	//Creamos el canal para la cantidad de buffers que necesito.
	status := make(chan bool, bufferSize)

	//Voy lanzando goroutines para cada uno de los pares. En el caso de tener un impar, lo proceso mas adelante
	for i := 0; i < numbersQty; i = i + 2 {
		go func(index int) {
			res += sumFunc(numeros[index], numeros[index+1])
			status <- true
		}(i)
	}

	//Aca tengo en cuenta el último caso que queda sin par. Al mismo le hago el par con 0, asi no cambia su valor
	if len(numeros)%2 != 0 {
		go func() {
			res += sumFunc(numeros[len(numeros)-1], 0)
			status <- true
		}()

	}

	//Esperamos que terminen todas las goroutines
	for i := 0; i < bufferSize; i++ {
		<-status
	}

	return res, nil
}
