package tp1

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func carritoInSlice(a Carrito, list []Carrito) bool {
	for _, b := range list {
		if b.Tienda == a.Tienda {
			return true
		}
	}
	return false
}
func findCarritoInSlice(a *Carrito, list *[]Carrito) int {
	for index, b := range *(list) {
		if b.Tienda == a.Tienda {
			return index
		}
	}
	return -1
}	

// comparador
type Comparer func(a interface{}, b interface{}) bool

func intComparer(a interface{}, b interface{}) bool {
	return a == b
}

func elementInSlice(a interface{}, list interface{}, comparer Comparer) bool {
	for _, b := range list.([]interface{}) {
		if comparer(a, b) {
			return true
		}
	}
	return false

}
