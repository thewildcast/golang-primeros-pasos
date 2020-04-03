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

func computar(sumas, mults, divisiones, restas <-chan *Operandos, resultados chan<- *Resultado) {
	for {
		select {
		case op := <-sumas:
			resultados <- &Resultado{Operacion: SUMA, Resultado: float64(op.A + op.B)}
		case op := <-mults:
			resultados <- &Resultado{Operacion: MULT, Resultado: float64(op.A * op.B)}
		case op := <-divisiones:
			resultados <- &Resultado{Operacion: DIVISION, Resultado: float64(op.A / op.B)}
		case op := <-restas:
			resultados <- &Resultado{Operacion: RESTA, Resultado: float64(op.A - op.B)}
		}
	}
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
	resultados := make(chan *Resultado)
	go computar(sumas, mults, divisiones, restas, resultados)
	go func() {
		<-corte
		close(resultados)
	}()
	return resultados
}
