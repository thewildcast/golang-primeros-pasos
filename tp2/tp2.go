package tp2

import "fmt"

// Error es una interfaz (especial pq empieza en minusc). Solo tenes q implementar el method "error" q returnea string
// type userError interface {
// 	Error() string
// }

// Se puede usar fmt.Errorf("String q puedo formatear y retornear como error %d y otro var %d", var1, var2)

type tupla struct {
	x, y int
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
	lenNum := len(numeros)
	var numTuplas int

	if lenNum == 0 {
		return 0, fmt.Errorf("No llego ningun numero")
	}

	tuplasRestantes := make(chan tupla, lenNum)
	resultados := make(chan int)

	var proxNum int
	for idx, num := range numeros {
		if idx%2 == 0 {
			if idx+1 == lenNum {
				proxNum = 0
			} else {
				proxNum = numeros[idx+1]
			}
			miTupla := tupla{num, proxNum}
			tuplasRestantes <- miTupla
			numTuplas++
		}

	}

	var miContador int
	for {
		go rutinaSumadora(tuplasRestantes, resultados, sumFunc)
		miContador++
		if miContador == 12 {
			break
		}
	}
	res1 := rutinaAgrupadora(resultados, tuplasRestantes, numTuplas)
	return res1, nil

}

func rutinaSumadora(tuplas <-chan tupla, res chan<- int, funcionCustomSuma sumador) {
	for miTupla := range tuplas {
		res <- funcionCustomSuma(miTupla.x, miTupla.y)
	}
}

func rutinaAgrupadora(results chan int, tuplas chan<- tupla, N int) int {
	cont := 0
	for {
		res0 := <-results
		if cont != N-1 {
			res1 := <-results
			tuplas <- tupla{res0, res1}
		} else {
			return res0
		}
		cont++
	}
}
