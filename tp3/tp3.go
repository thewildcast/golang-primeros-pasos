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

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
	res := make(chan *Resultado, 4)

	for {
		select {
		case op := <-sumas:
			res <- &Resultado{Operacion: SUMA, Resultado: float64(op.A + op.B)}
		case op := <-mults:
			res <- &Resultado{Operacion: MULT, Resultado: float64(op.A * op.B)}
		case op := <-divisiones:
			res <- &Resultado{Operacion: DIVISION, Resultado: float64(op.A / op.B)}
		case op := <-restas:
			res <- &Resultado{Operacion: RESTA, Resultado: float64(op.A - op.B)}
		default:
			if len(res) == cap(res) {
				close(res)
				return res
			}
		}
	}
}
