package tp3

type Op int

const (
	SUMA Op = iota
	RESTA
	DIVISION
	MULT
)

func (op Op) Calcular(opChan <-chan *Operandos, resultados chan *Resultado, corte chan <- struct{}) {

	defer func() {
		corte <- struct{}{}
	}()

	resultado := &Resultado{}
	switch op {
		case SUMA:
			operandos := <-opChan

			resultado.Operacion = SUMA
			resultado.Resultado = float64(operandos.A + operandos.B)

			resultados <- resultado
		case RESTA:
			operandos := <-opChan

			resultado.Operacion = RESTA
			resultado.Resultado = float64(operandos.A - operandos.B)

			resultados <- resultado
		case DIVISION:
			operandos := <-opChan

			resultado.Operacion = DIVISION
			resultado.Resultado = float64(operandos.A / operandos.B)

			resultados <- resultado
		case MULT:
			operandos := <-opChan

			resultado.Operacion = MULT
			resultado.Resultado = float64(operandos.A * operandos.B)

			resultados <- resultado
	}
}

type Operandos struct {
	A, B int
}

type Resultado struct {
	Operacion Op
	Resultado float64
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte chan struct{}) chan *Resultado {

	resultados := make(chan *Resultado)

	for i := 0; i < 4; i++ {
		op := Op(i)
		switch op {
		case SUMA:
			go SUMA.Calcular(sumas, resultados, corte)
		case RESTA:
			go RESTA.Calcular(restas, resultados, corte)
		case MULT:
			go MULT.Calcular(mults, resultados, corte)
		case DIVISION:
			go DIVISION.Calcular(divisiones, resultados, corte)
		}
	}

	go func() {
		opened := true

		for opened {
			_, opened = <-corte

			if !opened {
				close(resultados)
			}
		}
	}()

	return resultados
}
