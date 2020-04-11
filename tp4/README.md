example of request :

http://localhost:8001/?supermarkets=Wallmart&supermarkets=Coto&supermarkets=Disco&productids=1&productids=2

whe should get a response like this:

{
    "Coto": {
        "Precio": 3923
    },
    "Disco": {
        "Precio": 8866
    },
    "Wallmart": {
        "Precio": 10539
    }
}