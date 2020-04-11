package service

import (
	"strings"

	"github.com/wildcast/golang-primeros-pasos/tp4/model"
)

//Client generic operation
type Client interface {
	Call(url string) (model.SupermarketResponse, error)
}

//Execute Execute function
func Execute(client Client, url string, names []string, ids []string) (map[string]model.Carrito, error) {

	result := make(map[string]model.Carrito)
	var foundError error
	for _, storeName := range names {

		tmpURL := url
		tmpURL = strings.Replace(tmpURL, "name", storeName, 1)

		for _, id := range ids {

			tmpURL = strings.Replace(tmpURL, "id", id, 1)
			apiResponse, error := client.Call(tmpURL)
			foundError = error

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

	return result, foundError
}
