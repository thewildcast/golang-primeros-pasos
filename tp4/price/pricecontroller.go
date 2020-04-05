package price

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const urlProductosTemplate = "https://productos-p6pdsjmljq-uc.a.run.app/%s/productos/%d"

type product struct {
	Tienda string `json:"tienda"`
	ID     int    `json:"id,omitempty"`
	Precio int    `json:"precio"`
}

// PricesHandler manages request on /prices
func PricesHandler(response http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	marketIDs := params["market"]
	productIDs := []int{}
	for _, id := range params["product"] {
		value, err := strconv.Atoi(id)
		if err != nil {
			log.Printf("Products IDs: [%s]", params["productID"])
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		productIDs = append(productIDs, value)
	}

	products, err := getProducts(marketIDs, productIDs)
	if err != nil {
		if pErr, ok := err.(*productosError); ok {
			response.WriteHeader(pErr.StatusCode)
		} else {
			response.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(err)
		return
	}

	totals := products.calcularPrecios()

	jsonResponse, err := json.Marshal(totals)
	if err != nil {
		log.Printf("Error: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResponse)
}
