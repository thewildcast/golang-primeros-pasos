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
	result := make(chan *Resultado)
	go func() {
		for {
			select {
			case op := <-sumas:
				go func() {
					r := Resultado{SUMA, float64(op.A + op.B)}
					result <- &r
				}()
			case op := <-mults:
				go func() {
					r := Resultado{MULT, float64(op.A * op.B)}
					result <- &r
				}()
			case op := <-divisiones:
				go func() {
					r := Resultado{DIVISION, float64(op.A / op.B)}
					result <- &r
				}()
			case op := <-restas:
				go func() {
					r := Resultado{RESTA, float64(op.A - op.B)}
					result <- &r
				}()
			case <-corte:
				close(result)
				return
			}
		}
	}()
	return result
}
