package tp1

import (
	"encoding/json"
	"math/rand"
	"os"
)

type Producto struct {
	ID     int `json:"id"`
	Precio int `json:"precio"`
}

var (
	Supermercados = []string{
		"Dia",
		"Carrefour",
		"Disco",
		"SuperVea",
		"Coto",
		"Jumbo",
		"Target",
		"Whole Foods Market",
		"Wallmart",
		"Macro",
		"Nini",
	}
)

func GenerarProductos(archivo string) {
	productos := map[string][]Producto{}

	for _, nombre := range Supermercados {
		productos[nombre] = []Producto{}
		for i := 0; i < 50; i++ {
			productos[nombre] = append(productos[nombre], Producto{
				ID:     i,
				Precio: rand.Intn(12000),
			})
		}

	}

	f, err := os.Create(archivo)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(productos)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

func LeerProductos(archivo string) (map[string][]Producto, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}

	productos := map[string][]Producto{}
	if err = json.NewDecoder(f).Decode(&productos); err != nil {
		return nil, err
	}

	return productos, nil
}
