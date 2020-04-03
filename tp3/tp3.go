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

		for i := 0; i < 4; i++ {

			select {

			case s := <-sumas:
				result <- &Resultado{Resultado: float64(s.A) + float64(s.B), Operacion: (SUMA)}

			case m := <-mults:
				result <- &Resultado{Resultado: float64(m.A) * float64(m.B), Operacion: (MULT)}

			case d := <-divisiones:
				result <- &Resultado{Resultado: float64(d.A) / float64(d.B), Operacion: (DIVISION)}

			case r := <-restas:
				result <- &Resultado{Resultado: float64(r.A) - float64(r.B), Operacion: (RESTA)}

			}
		}
		close(result)

	}()

	return result
}
