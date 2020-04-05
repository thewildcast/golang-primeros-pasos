package tp3

import "time"

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

// Calcular deberia llevar comentarios pq es explortable
func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {

	results := make(chan *Resultado)

	to := time.After(time.Millisecond * 900)

	go func() {
		for {
			select {
			case ops := <-sumas:
				go func(ops Operandos) {
					results <- &Resultado{SUMA, float64(ops.A + ops.B)}
				}(*ops)
			case ops := <-mults:
				go func(ops Operandos) {
					results <- &Resultado{MULT, float64(ops.A * ops.B)}
				}(*ops)
			case ops := <-divisiones:
				go func(ops Operandos) {
					results <- &Resultado{DIVISION, float64(ops.A / ops.B)}
				}(*ops)
			case ops := <-restas:
				go func(ops Operandos) {
					results <- &Resultado{RESTA, float64(ops.A - ops.B)}
				}(*ops)
			case <-to:
				close(results)
			}
		}
	}()

	return results
}
