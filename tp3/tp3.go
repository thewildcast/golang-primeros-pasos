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
	canal := make(chan *Resultado)
	go func() {
		for {
			select {
			case s := <-sumas:
				res := Resultado{
					Resultado: float64(s.A + s.B),
					Operacion: SUMA,
				}
				canal <- &res
			case s := <-mults:
				res := Resultado{
					Resultado: float64(s.A * s.B),
					Operacion: MULT,
				}
				canal <- &res
			case s := <-divisiones:
				res := Resultado{
					Resultado: float64(s.A / s.B),
					Operacion: DIVISION,
				}
				canal <- &res
			case s := <-restas:
				res := Resultado{
					Resultado: float64(s.A - s.B),
					Operacion: RESTA,
				}
				canal <- &res
			case <-corte:
				close(canal)
				return
			}
		}
	}()
	return canal
}
