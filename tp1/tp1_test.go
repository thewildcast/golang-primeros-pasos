package tp1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularPrecios(t *testing.T) {
	carrito := Carrito{}
	result := carrito.CalcularPrecios([]int{1, 7})

	assert.IsType(t, Carrito{}, result[0], "Returned object is not a Carrito!")
	assert.True(t, true, result[1], "Result is not successful!")
}

func TestPromedio(t *testing.T) {
	carrito := Carrito{}
	result := carrito.Promedio(1)

	assert.Equal(t, 5045, result, "Expected promedio is wrong")
}

func TestBuscarMasBarato(t *testing.T) {
	carrito := Carrito{}
	result, ok := carrito.BuscarMasBarato(1)

	assert.IsType(t, Carrito{}, result, "Returned object is not a Carrito!")
	assert.True(t, true, ok, "Result is not successful!")
}
