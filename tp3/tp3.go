package tp3

import "log"

type Op int

const (
	SUMA Op = iota
	RESTA
	DIVISION
	MULT
)

func (op Op) String() string {
	switch op {
	case SUMA:
		return "suma"
	case RESTA:
		return "resta"
	case DIVISION:
		return "division"
	case MULT:
		return "mult"
	}
	return "invalida"
}

type Operandos struct {
	A, B int
}

type Resultado struct {
	Operacion Op
	Resultado float64
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {

	chanelResultados := make(chan *Resultado)
	go func() {
		for {
			select {
			case operationSum := <-sumas:
				result := Resultado{Operacion: SUMA, Resultado: float64(operationSum.A) + float64(operationSum.B)}
				chanelResultados <- &result
			case operationMult := <-mults:
				result := Resultado{Operacion: MULT, Resultado: float64(operationMult.A) * float64(operationMult.B)}
				chanelResultados <- &result
			case operationDiv := <-divisiones:
				result := Resultado{Operacion: DIVISION, Resultado: float64(operationDiv.A) / float64(operationDiv.B)}
				chanelResultados <- &result
			case operationRest := <-restas:
				result := Resultado{Operacion: RESTA, Resultado: float64(operationRest.A) - float64(operationRest.B)}
				log.Println("restas:", result)
				chanelResultados <- &result
			case <-corte:
				close(chanelResultados)
				return

			}
		}
	}()
	return chanelResultados
}
