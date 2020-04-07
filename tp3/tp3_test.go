package tp3

import (
	"testing"
	"time"
    "log"
)

func TestCalcular(t *testing.T) {
	sumas := make(chan *Operandos, 1)
	mults := make(chan *Operandos, 1)
	divisiones := make(chan *Operandos, 1)
	restas := make(chan *Operandos, 1)
	corte := make(chan struct{})

	sumas <- &Operandos{A: 2, B: 4}
	mults <- &Operandos{A: 3, B: 2}
	divisiones <- &Operandos{A: 12, B: 2}
	restas <- &Operandos{A: 10, B: 4}

	time.AfterFunc(1*time.Second, func() {
		panic("La funcion `Calcular` no cerro el channel `resultados` luego de haber recibido la informacion de `corte`. O esta tomando demasiado tiempo en procesar las cuatro operaciones.")
	})

	var i int

    log.Print("MANDOOOOO inicializacion de i:",i)

	for res := range Calcular(sumas, mults, divisiones, restas, corte) {
		i++
        log.Print("MANDOOOOO i:",i)

		if res.Resultado != 6 {
			t.Errorf("operacion de %s fallo. Se esperaba 6.0 pero retorno %.2f", res.Operacion, res.Resultado)
		}

		if i == 4 {
            log.Print("MANDOOOOO CORTE")
			close(corte)
		}
	}

	if i != 4 {
		t.Errorf("se deberian haber procesado 4 operaciones, pero se procesaron %d", i)
	}
}
