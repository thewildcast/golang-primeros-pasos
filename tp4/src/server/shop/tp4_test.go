package shop

import (
	"reflect"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	casos := []struct {
		nombre      string
		inputId     []int
		inputMarket []string
		precios     map[string]int
	}{
		{
			nombre:      "calcula la suma de precios correcta",
			inputId:     []int{1, 2},
			inputMarket: []string{"Target", "Coto", "Dia", "Disco", "Jumbo", "Macro", "Nini", "SuperVea", "Wallmart", "Whole Foods Market", "Carrefour"},
			precios: map[string]int{
				"Target":             8536,
				"Coto":               3923,
				"Dia":                15734,
				"Disco":              8866,
				"Jumbo":              10194,
				"Macro":              20559,
				"Nini":               12053,
				"SuperVea":           8195,
				"Wallmart":           10539,
				"Whole Foods Market": 12785,
				"Carrefour":          6910,
			},
		},
		{
			nombre:      "da cero cuando no hay productos",
			inputId:     []int{},
			inputMarket: []string{},
			precios:     map[string]int{},
		},
	}

	for _, test := range casos {
		t.Run(test.nombre, func(t *testing.T) {
			prices, _ := GetTotalPrice(test.inputId, test.inputMarket)
			if len(prices) != len(test.precios) {
				t.Errorf("CalcularPrecios retorno %d supermercados, se esperaban %d", len(prices), len(test.precios))
			}

			result := map[string]int{}
			for _, totalPrice := range prices {
				result[totalPrice.Shop] = totalPrice.Total
			}

			if !reflect.DeepEqual(result, test.precios) {
				t.Errorf("CalcularPrecios retorna precios incorrectos %+v, se esperaban %+v", result, test.precios)
			}
		})
	}
}
