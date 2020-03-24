# Primeros pasos en Golang

## TP1
El objetivo de este TP es que practiques como iterar slices, utilizar mapas, definir metodos, correr tests y mucho mas!

Los ejercicios para el TP1 se encuentran en `tp1/tp1.go`. Como input siempre vas a recibir un mapa que contiene el nombre de un super mercado como clave y un arreglo de los productos que ese super mercado tiene.

* CalcularPrecios: dada una lista de ids de productos la funcion deberia calcular cual seria el precio total de todos los productos para cada uno de los super mercados,
* Promedio: dado el id de un producto la funcion deberia calcular cual es el precio promedio de ese producto utilizando la data de todos los super mercados,
* BuscarMasBarato: recibe el id de un producto y deberia retornar cual es el super mercado que lo vende mas barato y a cuanto lo vende.

### Como probar tu solucion
Cada una de las funciones definidas en `tp1.go` tiene una function escrita en `tp1_test.go` que actua como *test* de esa funcion. Ya tenemos un conjunto de casos identificados y escritos en los tests que validan que tu funcion ejecute como corresponde.  
Para validar tus soluciones, podes ejecutar los siguientes comandos estando parado/a en la carpeta de `tp1`:
```
# correr los tests de la funcion sumar. Aca podrian cambiar
# `CalcularPrecios` por el nombre de la funcion que estan probando en el momento
go test -run=TestTienda_CalcularPrecios

# correr los tests de todas las funciones
go test tp1_test.go
```

Si quieren tener mas informacion de los tests, por ejemplo saber cuales escenarios fallaron, pueden correr el comando con el flag `-v`:
```
go test -v -run=TestTienda_CalcularPrecios
```
