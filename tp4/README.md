API usage :

http://localhost:8001/?supermarkets=Wallmart&supermarkets=Coto&supermarkets=Disco&productids=1&productids=2

According to last request, we should get following response;

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