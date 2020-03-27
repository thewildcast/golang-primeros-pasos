# Primeros pasos en Golang

## TP1
El objetivo de este TP es que practiques cómo iterar slices, mapas, definir métodos, correr tests, user interfaces  y más!

Para los ejercicios de este TP vamos a estar usando un dataset que contiene un conjunto de productos formados de la siguiente manera:
```
[
	["dia", "1", "234"]
]
```

Cada ejercicio tiene una funcion definida en `tp1/tp1.go` que vas a tener que implementar para que sus tests, definidos en `tp1/tp1_test.go`, pasen exitosamente. Estas son las consignas:
* **CalcularPrecios**: Dada una lista de ids de productos la función debería calcular cuál sería el precio total de todos los productos para cada uno de los supermercados. Basicamente, armar un "carrito" para cada super mercado que se encuentra e indicar cuanto saldria comprar esos items en ese super mercado,
* **Promedio**: Dado el id de un producto la función debería calcular cuál es el precio promedio de ese producto utilizando la data de todos los supermercados,
* **BuscarMasBarato**: Recibe el ID de un producto y debería retornar cuál es el super mercado que lo vende más barato y a cuánto lo vende.

La funcion `BuscarMasBarato` retorna una interfaz llamada `Producto`. Vas a tener que definir algun tipo que cumpla con la definicion de esa interfaz para poder resolver el ejercicio.

![](images/tp1.jpeg)

### ¿Cómo probar tu solución?
Cada una de las funciones definidas en `tp1.go` tiene una `func` escrita en `tp1_test.go` que actúa como *test* de esa función. Ya tenemos un conjunto de casos identificados y escritos en los tests que validan que tu codigo haga lo que corresponde.  

Para validar tus soluciones, podés ejecutar los siguientes comandos parándote en la carpeta de `tp1`:

```
# correr los tests de la funcion sumar. Aca podrian cambiar
# `CalcularPrecios` por el nombre de la funcion que estan probando en el momento
go test -run=TestProductos_CalcularPrecios

# Correr los tests de todas las funciones
go test
```

Si quieren tener más información de los tests, por ejemplo saber cuáles escenarios fallaron, pueden correr el comando con el flag `-v`:
```
go test -v -run=TestProductos_CalcularPrecios
```

## TP2
El objetivo de este TP es que aprendas como usar las capacidades de concurrencia que nos ofrece Go. Vas a usar channels y goroutines para poder ejecutar y coordinar eficientemente una suma de muchos numeros.

Para realizar la suma de dos numeros ya tenes una funcion a tu disposicion que se llama `Sumar` y recibe dos enteros. El problema aca es que `Sumar` tarda una cantidad de tiempo indeterminado (pero no mayor a medio segundo) en retornarte la suma, entonces hacer una suma secuencial de muchos numeros nos puede llegar a tardar mucho tiempo. Para poder hacerla mas eficiente vas a usar goroutines que paralelicen estas operaciones y reduazcan el tiempo que toma sumar todo. Una forma de encarar esto seria usando semaforos y mantener una variable global de suma que cada goroutine incremente, el problema es que esto no esta acorde a uno de los [proverbios mas importantes de Go](https://go-proverbs.github.io/):
> [Don't communicate by sharing memory, share memory by communicating](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s).

Ese proverbio simplemente indica que usar operadores de exclusion mutua y variables compartidas no es lo mejor que se puede hacer. No es una mala herramienta, y a veces es totalmente necesaria, pero para este caso se puede encarar la solucion usando canales para la comunicacion y reducir el overhead de tener que proteger la memoria manualmente.

La consiga es:
* **SumarLista**: Dada una lista de enteros la funcion deberia calcular la suma de todos ellos y retornar un error en caso de que no se reciban numeros.

### ¿Cómo probar tu solución?
Al igual que en el TP1, definimos tests para la funcion `SumarLista` que estan escritos en `tp2/tp2_test.go`. Una vez que termines tu solucion podes correr el siguiente comando estan parado/a en la carpeta `tp2`:
```
go test -run=TestSumarLista
```
