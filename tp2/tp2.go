package tp2

import "fmt"

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	var resultado, longNumeros int

	longNumeros = len(numeros)

	if longNumeros == 0 {
		return 0, fmt.Errorf("No se puede realizar la suma")
	}

	numerosASumar := make(chan int, longNumeros)
	resultadoSuma := make(chan int, longNumeros)

	go GenerarSuma(numerosASumar, resultadoSuma, sumFunc)

	for i := 0; i < longNumeros; i++ {

		numerosASumar <- numeros[i]
	}

	close(numerosASumar)

	for valor := range resultadoSuma {

		resultado = sumFunc(resultado, valor)
	}

	return resultado, nil

}

//GenerarSuma recibe ambos canales, realiza la suma
//el resultado de la suma lo guarda en resultadoSuma.
func GenerarSuma(numerosASumar <-chan int, resultadoSuma chan<- int, sumFunc sumador) int {

	var resultado int

	for valor := range numerosASumar {

		resultadoSuma <- sumFunc(valor, resultado)
	}

	close(resultadoSuma)

	return resultado
}
