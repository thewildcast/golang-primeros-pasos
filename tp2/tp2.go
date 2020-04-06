package tp2


// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.

type err struct{
	
}

func (e err) Error()string{

	return "Hubo un error al querer realizar la operación."
}

func SumarLista(sumFunc sumador, numeros []int) (int, error) {
		
	if len(numeros) == 0{
		return 0, err{}
	}

	var restulado int

	for i:= 0; i<len(numeros); i++{

		resultado = sumFunc(numeros[i], numeros[i+1])
	}
	
	return resultado, nil

}


