package tp3

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

func sumar(op *Operandos, resultado chan<- *Resultado) {
	resultado <- &Resultado{Operacion: SUMA, Resultado: float64(op.A) + float64(op.B)}
}

func multiplicar(op *Operandos, resultado chan<- *Resultado) {
	resultado <- &Resultado{Operacion: MULT, Resultado: float64(op.A) * float64(op.B)}
}

func dividir(op *Operandos, resultado chan<- *Resultado) {
	resultado <- &Resultado{Operacion: DIVISION, Resultado: float64(op.A) / float64(op.B)}
}

func restar(op *Operandos, resultado chan<- *Resultado) {
	resultado <- &Resultado{Operacion: RESTA, Resultado: float64(op.A) - float64(op.B)}
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
	resultados := make(chan *Resultado)

	go func() {
		for {
			select {
			case operandos := <-sumas:
				go sumar(operandos, resultados)
			case operandos := <-mults:
				go multiplicar(operandos, resultados)
			case operandos := <-divisiones:
				go dividir(operandos, resultados)
			case operandos := <-restas:
				go restar(operandos, resultados)
			case <-corte:
				close(resultados)
				return
			}
		}
	}()

	return resultados
}
