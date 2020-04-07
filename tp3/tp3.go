package tp3

import ( 
       // "time"
        "log"
        )
        
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

func Procesar(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{},chanResultado chan *Resultado){
   // defer close(chanResultado)
	//timeout := time.After(time.Second)

    for{
        select{
            case s:= <- sumas:
                log.Println("Entro Suma!")
                resAux := float64(s.A+s.B)
                chanResultado <- &Resultado {SUMA,resAux}
            case s:= <- mults:
                log.Println("Entro Mults!")
                resAux := float64(s.A*s.B)
                chanResultado <- &Resultado {MULT,resAux}
            case s:= <- divisiones:
                log.Println("Entro Divisiones!")
                if s.B != 0{
                  resAux := float64(s.A/s.B)
                  chanResultado <- &Resultado {DIVISION,resAux}
                }
            case s:= <- restas:
                log.Println("Entro Resta!")
                resAux := float64(s.A - s.B)
                chanResultado <- &Resultado {RESTA,resAux}
            case <- corte:
                log.Println("CORTEE!")
                close(chanResultado)
                //return
           // case <- timeout:
            //    log.Println("Timeout!")
            //    close(chanResultado)

               //default:
                 //   log.Println("Default!")
            }
            log.Println("SALGO DEL SELECT!")

     }
    close(chanResultado)
}

func Calcular(sumas, mults, divisiones, restas <-chan *Operandos, corte <-chan struct{}) chan *Resultado {
    
	chanResultado := make(chan *Resultado, 4)
    
    go Procesar(sumas,mults,divisiones,restas,corte, chanResultado)

    return chanResultado
}
