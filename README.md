# Primeros pasos en Golang

## TP1
El objetivo de este TP es que practiques cómo iterar slices, utilizar mapas, definir métodos, correr tests y mucho más!

Los ejercicios para el TP1 se encuentran en `tp1/tp1.go`. Como input siempre vas a recibir un mapa que contiene el nombre de un supemercado como clave y un arreglo de los productos que ése supermercado tiene.

* CalcularPrecios: Dada una lista de IDs de productos la función debería calcular cuál sería el precio total de todos los productos para cada uno de los supermercados.
* Promedio: Dado el ID de un producto la función debería calcular cuál es el precio promedio de ese producto utilizando la data de todos los supermercados.
* BuscarMasBarato: Recibe el ID de un producto y debería retornar cuál es el supermercado que lo vende más barato y a cuánto lo vende.

#### ¿Cómo probar tu solución?
Cada una de las funciones definidas en `tp1.go` tiene una `func` escrita en `tp1_test.go` que actúa como *test* de esa función. Ya tenemos un conjunto de casos identificados y escritos en los tests que validan que tu función se ejecute como corresponde.  

Para validar tus soluciones, podés ejecutar los siguientes comandos parándote en la carpeta de `tp1`:

```
# correr los tests de la funcion sumar. Aca podrian cambiar
# `CalcularPrecios` por el nombre de la funcion que estan probando en el momento
go test -run=TestTienda_CalcularPrecios

# Correr los tests de todas las funciones
go test tp1_test.go
```

Si quieren tener más información de los tests, por ejemplo saber cuáles escenarios fallaron, pueden correr el comando con el flag `-v`:
```
go test -v -run=TestTienda_CalcularPrecios
```
