package tp2

import (
	"fmt"
)

type ErrorSumarLista struct {
	mensaje string
}

func (err ErrorSumarLista) Error() string {
	return err.mensaje
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	fmt.Println("Entrada inicial:", numeros)
	if len(numeros) == 0 {
		return 0, ErrorSumarLista{"Lista vacia"}
	}

	resultadoFinal := sumaRecursiva(sumFunc, numeros)
	if len(resultadoFinal) > 1 {
		return 0, ErrorSumarLista{"Fallo la suma recursiva"} // no deberia pasar nunca, just in case..
	}
	return resultadoFinal[0], nil
}

func sincronizar(resultados chan int, listo chan bool) {
	for len(resultados) < cap(resultados) {
		// seguir esperando hasta llenar el buffer
	}
	close(resultados)
	listo <- true
	close(listo)
}

func sumaRecursiva(sumFunc sumador, entrada []int) []int {
	if len(entrada)%2 > 0 {
		entrada = append(entrada, 0) // array de length par
	}

	resultados := make(chan int, len(entrada)/2)
	for i := 0; i < len(entrada); i += 2 {
		go func(idx int) {
			resultados <- sumFunc(entrada[idx], entrada[idx+1])
		}(i)
	}
	listo := make(chan bool)
	go sincronizar(resultados, listo)
	<-listo // esperamos hasta que podamos proceder

	fmt.Println("termine de producir, nro de elems:", len(resultados))
	salida := []int{}
	for r := range resultados {
		salida = append(salida, r)
	}
	fmt.Println("termine de consumir, elementos:", salida)
	if len(salida) == 1 { // condicion de corte, tengo un solo elemento de toda la suma
		return salida
	}
	return sumaRecursiva(sumFunc, salida) // necesito seguir sumando
}
