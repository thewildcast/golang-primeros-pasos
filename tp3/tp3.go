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

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) <-chan *Resultado {
	res := make(chan *Resultado)
	go func() {
		defer close(res)
		for {
			select {
			case suma := <-sumas:
				res <- &Resultado{Operacion: SUMA, Resultado: float64(suma.A + suma.B)}
			case resta := <-restas:
				res <- &Resultado{Operacion: RESTA, Resultado: float64(resta.A - resta.B)}
			case division := <-divisiones:
				res <- &Resultado{Operacion: DIVISION, Resultado: float64(division.A / division.B)}
			case mult := <-mults:
				res <- &Resultado{Operacion: MULT, Resultado: float64(mult.A * mult.B)}
			case <-corte:
				return
			}
		}
	}()
	return res
}
