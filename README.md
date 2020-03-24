# Primeros pasos en Golang

## TP1
El objetivo de este TP es que practiques como iterar slices, utilizar mapas, definir metodos, correr tests y mucho mas!

Los ejercicios para el TP1 se encuentran en `tp1/tp1.go`. Como input siempre vas a recibir una matriz que contiene la informacion acerca de las ventas de un super mercado.

* Sum: en este ejercicio recibis como input una matriz de informacion acerca de ventas
* Average: en este 

#### Como probar tu solucion
Cada una de las funciones definidas en `tp1.go` tiene una function escrita en `tp1_test.go` que actua como *test* de esa funcion. Ya tenemos un conjunto de casos identificados y escritos en los tests que validan que tu funcion ejecute como corresponde.  
Para validar tus soluciones, podes ejecutar los siguientes comandos estando parado/a en la carpeta de `tp1`:
```
# correr los tests de la funcion sum. Aca podrian cambiar
# `Sum` por el nombre de la funcion que estan probando en el momento
go test -run=TestStore_Sum

# correr los tests de todas las funciones
go test tp1_test.go
```

Si quieren tener mas informacion de los tests, por ejemplo saber cuales escenarios fallaron, pueden correr el comando con el flag `-v`:
```
go test -v -run=TestStore_Sum
```
