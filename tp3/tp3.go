package tp3

import "fmt"

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

func CalculoAuxiliar(sumas, mults, divisiones, restas <-chan *Operandos, corte <- chan struct{}, resultados chan <- *Resultado){


	for{
		select {
		case s:=<- sumas:
			
			resSuma := float64(s.A + s.B)

			resultados <- &Resultado{SUMA,resSuma}

		case r:=<- restas:

			resResta := float64(r.A - r.B)

			resultados <- &Resultado{RESTA,resResta}

		case m:=<- mults:
			
			resMulti:= float64(m.A * m.B)

			resultados <- &Resultado{MULT, resMulti}

		case d:=<-divisiones:
			
			if d.B != 0{

				resDiv := float64(d.A / d.B)

				resultados <- &Resultado{DIVISION,resDiv}
			}

		case <-corte:
			fmt.Println("Fin de las operaciones")
			close(resultados)
			
		}
	}
	close(resultados)
	
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {

	resutlados := make(chan *Resultado, 4)

	go CalculoAuxiliar(sumas, mults, divisiones, restas, corte, resutlados)

	return resutlados

}