# Solucion TP4

Se implemento una API que contiene un endpoint `/precios` que puede usarse de la siguiente forma:

```
/precios?pid=1&pid=2&sid=target&sid=dia&sid=jumbo
```

Esa llamada devuelve un JSON como este:

```
[{"Tienda":"target","Precio":8536},{"Tienda":"dia","Precio":15734},{"Tienda":"jumbo","Precio":10194}]
```

que contiene la suma de los precios de los productos indicados mediante los parametros `pid` en los supermercados indicados mediante los parametros `sid`.
