package tp2

import (
    "fmt"
    "time"
)
// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {
    cantNum := len(numeros)
    if cantNum == 0{
        return 0 , fmt.Errorf("no se puede sumar una lista vacia")
    }
    enteros:= make (chan int, cantNum) // definir el buffer
    resultado:= make (chan int, cantNum) // definir el buffer
    
    //la idea es que a medida que una funcion va alojando la suma en un canal desde otro en paraleto voy
    //tomando esa suma y voy generando un resultado final.
    go sumFucCon (enteros, resultado, sumFunc)
    
    for i:=0; i < cantNum; i++ {
	   enteros <- numeros[i]
    }
    
    close(enteros)//importante cerrar el channel
    
    var resultadoFinal int
    for r := range resultado {
       resultadoFinal = sumFunc(resultadoFinal,r)
   }
   return resultadoFinal, nil
}


func sumFucCon(enteros chan int, resultado chan int,sumFunc sumador) int {
	time.Sleep(time.Second*1)
    var res int
	for p := range enteros {
        resultado <- sumFunc(res,p)
	}
	close(resultado) //importante cerrar el channel despues de no escribir mas
    return res
}
 
   