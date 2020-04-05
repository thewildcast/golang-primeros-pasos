package price

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type productos []product

type productosError struct {
	StatusCode int
	Message    string
}

func (pErr productosError) Error() string {
	return pErr.Message
}

func getProducts(marketIDs []string, productIDs []int) (*productos, error) {
	products := productos{}

	responses := make(chan *product, len(marketIDs)*len(productIDs))
	responsesErr := make(chan *productosError)

	for _, marketID := range marketIDs {
		for _, productID := range productIDs {
			getProduct(marketID, productID, responses, responsesErr)
		}
	}

	requestsDone := 0
	for {
		select {
		case p := <-responses:
			requestsDone++
			products = append(products, *p)
			if requestsDone == len(marketIDs)*len(productIDs) {
				return &products, nil
			}

		case pErr := <-responsesErr:
			return nil, pErr
		}
	}
}

func getProduct(marketID string, productID int, responses chan<- *product, responsesErr chan<- *productosError) {
	go func() {
		p, err := get(marketID, productID)
		if err != nil {
			responsesErr <- err
		} else {
			responses <- p
		}
	}()
}

func get(marketID string, productID int) (*product, *productosError) {
	url := fmt.Sprintf(urlProductosTemplate, marketID, productID)

	productosResponse, err := http.Get(url)
	if err != nil {
		return nil, &productosError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Error URL: [%s] - %s", url, err),
		}
	}

	if productosResponse.StatusCode != http.StatusOK {
		statusCode := productosResponse.StatusCode
		if statusCode != http.StatusNotFound {
			statusCode = http.StatusServiceUnavailable
		}
		return nil, &productosError{
			StatusCode: statusCode,
			Message:    fmt.Sprintf("[%d] - URL: [%s]\n", statusCode, url),
		}
	}

	p := product{}

	defer productosResponse.Body.Close()

	if err := json.NewDecoder(productosResponse.Body).Decode(&p); err != nil {
		return nil, &productosError{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Error: %s", err),
		}
	}

	return &p, nil
}

func (p *productos) calcularPrecios() *productos {
	subtotales := map[string]int{}

	for _, producto := range *p {
		subtotales[producto.Tienda] += producto.Precio
	}

	totales := productos{}

	for market, totalPrice := range subtotales {
		totales = append(totales, product{Tienda: market, Precio: totalPrice})
	}

	return &totales
}
