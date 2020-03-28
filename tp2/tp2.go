package tp2

import "errors"

// SumarLista recibe una lista de numeros y retorna la suma
// de todos esos numeros. Si la suma no se puede realizar
// por algun motivo se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if numeros == nil {
		return 0, errors.New("No hay numeros para sumar")
	}
	res := 0
	if len(numeros)%2 != 0 {
		res = numeros[len(numeros)-1]
		numeros = numeros[:len(numeros)-1]
	}

	for i := 0; i < len(numeros); i += 2 {
		res += sumFunc(numeros[i], numeros[i+1])
	}
	return res, nil
}
