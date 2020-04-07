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

//Calcular ...
func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
	rs := make(chan *Resultado, 4)
	go func() {
		for {
			select {
			case <-corte:
				close(rs)
				return
			case op := <-sumas:
				rs <- &Resultado{SUMA, float64(op.A + op.B)}
			case op := <-mults:
				rs <- &Resultado{MULT, float64(op.A * op.B)}
			case op := <-divisiones:
				if op.B != 0 {
					rs <- &Resultado{DIVISION, float64(op.A / op.B)}
				}
			case op := <-restas:
				rs <- &Resultado{RESTA, float64(op.A - op.B)}
			}
		}
	}()
	return rs
}
