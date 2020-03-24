package tp1

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

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
	productos := [][]string{}

	for _, nombre := range Supermercados {
		for i := 0; i < 50; i++ {
			productos = append(productos, []string{nombre, fmt.Sprintf("%d", i), fmt.Sprintf("%d", rand.Intn(12000))})
		}

	}

	f, err := os.Create(archivo)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := json.Marshal(productos)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write(b); err != nil {
		panic(err)
	}
}

func LeerProductos(archivo string) ([][]string, error) {
	f, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	productos := [][]string{}
	if err = json.NewDecoder(f).Decode(&productos); err != nil {
		return nil, err
	}

	return productos, nil
}
