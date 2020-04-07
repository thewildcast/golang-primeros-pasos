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
	go resolve(sumas, mults, divisiones, restas, corte, result)
	return result
}

func resolve(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}, result chan *Resultado) {
	for {
		select {
		case sum := <-sumas:
			result <- &Resultado{
				Operacion: SUMA,
				Resultado: float64(sum.A + sum.B),
			}
		case mult := <-mults:
			result <- &Resultado{
				Operacion: MULT,
				Resultado: float64(mult.A * mult.B),
			}
		case div := <-divisiones:
			result <- &Resultado{
				Operacion: DIVISION,
				Resultado: float64(div.A / div.B),
			}
		case dif := <-restas:
			result <- &Resultado{
				Operacion: RESTA,
				Resultado: float64(dif.A - dif.B),
			}
		case _, open := <-corte:
			if !open {
				close(result)
				return
			}
		}
	}
}
