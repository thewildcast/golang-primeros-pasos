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

func calculator(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}, resultados chan<- *Resultado) {
	for {
		select {
		case o := <-sumas:
			resultados <- &Resultado{SUMA, float64(o.A + o.B)}
		case o := <-mults:
			resultados <- &Resultado{MULT, float64(o.A * o.B)}
		case o := <-divisiones:
			if o.B != 0 {
				resultados <- &Resultado{DIVISION, float64(o.A / o.B)}
			}
		case o := <-restas:
			resultados <- &Resultado{RESTA, float64(o.A - o.B)}
		case <-corte:
			close(resultados)
			return
		}
	}
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
	resultados := make(chan *Resultado)
	go calculator(sumas, mults, divisiones, restas, corte, resultados)
	return resultados
}
