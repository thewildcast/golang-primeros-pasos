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
	resultados := make(chan *Resultado)
	go func() {
		for {
			select {
			case s := <-sumas:
				resultados <- &Resultado{Operacion: SUMA, Resultado: float64(s.A + s.B)}
			case m := <-mults:
				resultados <- &Resultado{Operacion: MULT, Resultado: float64(m.A * m.B)}
			case d := <-divisiones:
				resultados <- &Resultado{Operacion: DIVISION, Resultado: float64(d.A / d.B)}
			case r := <-restas:
				resultados <- &Resultado{Operacion: RESTA, Resultado: float64(r.A - r.B)}
			case <-corte:
				close(resultados)
				return
			}
		}
	}()
	return resultados
}
