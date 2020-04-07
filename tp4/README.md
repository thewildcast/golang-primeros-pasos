## TP4

### ¿Cómo probar la solución?
Cada una de las funciones definidas en `tp4.go` tiene una `func` escrita en `tp4_test.go` que actúa como *test* de esa función.  

```
# `ObtenerProducto` por el nombre de la funcion que estan probando en el momento
go test -run=TestObtenerProducto

# Correr los tests de todas las funciones
go test
```


### Para correr el server y probar con un request HTTP

```
go run main.go

curl http://localhost:8000/tiendas -X POST -i -v -d '{"ids": [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30], "tiendas": ["Dia","Carrefour","Disco","SuperVea","Coto","Jumbo","Target","Whole Foods Market","Wallmart","Macro","Nini","lalala"]}'

Se recibira una respuesta como la siguiente:

[{"tienda":"Disco","total":146483},{"tienda":"SuperVea","total":107260},{"tienda":"Whole Foods Market","total":125912},{"tienda":"Wallmart","total":96700},{"tienda":"lalala","total":0},{"tienda":"Jumbo","total":145914},{"tienda":"Macro","total":91259},{"tienda":"Nini","total":107385},{"tienda":"Dia","total":105860},{"tienda":"Carrefour","total":165685},{"tienda":"Target","total":132210},{"tienda":"Coto","total":174635}]
```
