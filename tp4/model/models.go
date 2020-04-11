package model

//Carrito struct using for calculation
type Carrito struct {
	Precio int
}

//SupermarketResponse use for client response
type SupermarketResponse struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id"`
	Precio int    `json:"precio"`
}

//Response represents an error if there is not product for a given id
type Response struct {
	StatusCode string
	Message    string
}
