package supermarket

import (
	"encoding/json"
	"github.com/wildcast/golang-primeros-pasos/tp4/client"
	"net/http"
	"strings"
	"github.com/wildcast/golang-primeros-pasos/tp4/model"
)

//PriceHandler with will be on charge of perform price operation
func PriceHandler(w http.ResponseWriter, r *http.Request) {

	requestData := r.URL.Query()
	url := "https://productos-p6pdsjmljq-uc.a.run.app/name/productos/id"

	supermarketsNames := requestData["supermarkets"]
	productIds := requestData["productids"]
    result:= make(map[string]model.Carrito)

	for _, storeName := range supermarketsNames {

		tmpURL := url
		tmpURL = strings.Replace(tmpURL, "name", storeName, 1)

		for _, id := range productIds {

			tmpURL = strings.Replace(tmpURL, "id", id, 1)
			apiResponse := client.Call(tmpURL, w)

			if apiResponse.Precio == 0 || apiResponse.Tienda == "" {

				result = nil
				break
			}
				
				_, present := result[storeName]

				if present {
	
					newPrice := result[storeName].Precio + apiResponse.Precio
					result[storeName] = model.Carrito{Precio: newPrice}
				} else {
			
					result[storeName] = model.Carrito{Precio: apiResponse.Precio}
				}
	
				tmpURL = strings.Replace(tmpURL, string(id), "id", 1)
	
		}
	}

	if result != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result) 
	}
	
}
